package server

import (
	"ConfigApp/cache"
	"ConfigApp/config"
	"ConfigApp/logging"
	"ConfigApp/model"
	"ConfigApp/storage"
	"ConfigApp/user"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err !=nil {
		log.Fatal("Error loading .env")
	}
}

type APIServer struct {
	router *gin.Engine
	storage storage.Storage
	userHandler user.UserHandler
	config config.Config
	cache cache.Cache
}

func NewAPIServer(storage storage.Storage, userHandler user.UserHandler, config config.Config, cache cache.Cache) *APIServer{
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4242"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server := &APIServer{
		router: r,
		storage: storage,
		userHandler: userHandler,
		config: config,
		cache: cache,
	}
	return server
}

func (s *APIServer) Run() {

	s.registerRoutes()
	fmt.Println("Server is running on port: " + s.config.Server.Port)
	s.router.Run(":"+ s.config.Server.Port)
}

func (s *APIServer) registerRoutes() {
	s.router.POST("/device/assign", s.assignDeviceInfo)
	s.router.GET("/devices/:orgId", s.getDeviceInfosByOrgId)
	s.router.POST("/org/create", s.createOrg)
	s.router.GET("/org/connected", s.getOrganizationsConnectedToUser)
	s.router.GET("/org/devices/:deviceId", s.getDeviceInfosByDeviceId)
	s.router.GET("/deviceData/:deviceId", s.getDeviceDataByDeviceId)
	s.router.POST("/deviceData/get_unique_id", s.getOrCreateDeviceID)
	s.router.POST("/update_sensor", s.updateSensor)
}

func (s *APIServer) assignDeviceInfo(c *gin.Context) {
	
	deviceInfoRequest := new(model.AddDeviceInfo)

	if ok := getDeviceInfoFromAPI(c, deviceInfoRequest); ok != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": ok.Error()})
		return
	}

	deviceInfo, err := s.storage.AssignDeviceToOrganization(*deviceInfoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, deviceInfo)
}

func (s *APIServer) createOrg(c *gin.Context) {

	token, err := c.Cookie(s.config.Server.AuthCookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	user, ok := s.userHandler.GetUserData(token)
	if ok != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	organizationDataReq := new(model.OrganizationDataRequest)

	if ok := getOrganizationDataFromAPI(c, organizationDataReq); ok != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": ok.Error()})
		return
	}

	organizationDataResponse, err := s.storage.CreateOrganization(*organizationDataReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := s.storage.CreateUserOrganizationConnection(organizationDataResponse.ID, user.ID, "owner"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, organizationDataResponse)
}

func (s *APIServer) updateSensor(c *gin.Context) {

	sensorUpdateRequest := new(model.SensorUpdate)  

	if err := c.ShouldBindJSON(&sensorUpdateRequest); err != nil {
		logging.Log.Warnf("Error during binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
	updatedSensor, err := s.storage.UpdateSensor(*sensorUpdateRequest)

	if err != nil {
		logging.Log.Warnf("Error during updating: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logging.Log.Infof("Sensor with id: %v updated successfully", updatedSensor.Id)
    c.JSON(http.StatusOK, updatedSensor)
}

func (s *APIServer) getDeviceInfosByOrgId(c *gin.Context) {
	orgId, err := strconv.Atoi(c.Param("orgId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse orgId"})
		return
	}

	deviceInfos, err := s.storage.GetDeviceInfoByOrgId(orgId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, deviceInfos)
}

func (s *APIServer) getDeviceInfosByDeviceId(c *gin.Context) {
	deviceId, err := strconv.Atoi(c.Param("deviceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse deviceId"})
		return
	}

	deviceInfo, err := s.storage.GetDeviceInfoByDeviceId(deviceId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, deviceInfo)
}

func (s *APIServer) getDeviceDataByDeviceId(c *gin.Context) {
	deviceId := c.Param("deviceId")

	deviceData, err := s.getDeviceData(deviceId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logging.Log.Errorf("Error getting device data: %v", err)
		return
	}

    c.JSON(http.StatusOK, deviceData)
}

func (s *APIServer) getOrCreateDeviceID(c *gin.Context) {
    var request model.DeviceInitDataRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if request.Token != "sample_token" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
        return
    }

    device, _ := s.storage.GetDeviceInfoByMAC(request.MAC)

    if device == nil {
		newDeviceRequest := model.AddDeviceInfo{
			Interval: 1000,
			OrgId: 0,
			MAC:      request.MAC,
		}

		createdDevice, err := s.storage.CreateDeviceInfo(newDeviceRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new device: " + err.Error()})
			return
		}

		if err := s.storage.CreateInitialSensorsForDevice(createdDevice.Id, 4); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	
		c.JSON(http.StatusOK, model.DeviceInitDataResponse{Id: strconv.Itoa(createdDevice.Id)})
		return
	}

    c.JSON(http.StatusOK, model.DeviceInitDataResponse{Id: strconv.Itoa(device.Id)})
}


func (s *APIServer) getOrganizationsConnectedToUser(c *gin.Context) {

	token, err := c.Cookie(s.config.Server.AuthCookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	user, ok := s.userHandler.GetUserData(token)
	if ok != nil {
		fmt.Println("error here 2")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	organizationsConnectedToUser, err := s.storage.GetOrganizationsConnectedToUser(user.ID)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(organizationsConnectedToUser)

    c.JSON(http.StatusOK, organizationsConnectedToUser)
}