package server

import (
	"ConfigApp/model"
	"ConfigApp/storage"
	"ConfigApp/user"
	"fmt"
	"log"
	"net/http"
	"os"
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
}

func NewAPIServer(storage storage.Storage, userHandler user.UserHandler) *APIServer{
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
	}
	return server
}

func (s *APIServer) Run() {

	s.registerRoutes()

	s.router.Run(":"+ os.Getenv("PORT"))
}

func (s *APIServer) registerRoutes() {
	s.router.POST("/devices", s.addDeviceInfo)
	s.router.POST("/org/create", s.createOrg)
	s.router.GET("/devices/:orgId", s.getDeviceInfosByOrgId)
	s.router.GET("/org/devices/:deviceId", s.getDeviceInfosByDeviceId)
	s.router.GET("/deviceData/:deviceId", s.getDeviceDataByDeviceId)
}

func (s *APIServer) addDeviceInfo(c *gin.Context) {
	
	deviceInfoRequest := new(model.DeviceInfo)

	if ok := getDeviceInfoFromAPI(c, deviceInfoRequest); ok != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": ok.Error()})
		return
	}

	deviceInfo, err := s.storage.CreateDeviceInfo(*deviceInfoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, deviceInfo)
}

func (s *APIServer) createOrg(c *gin.Context) {
	
	token, err := c.Cookie("sb-zsgthpzpbdkcdcdyzbkt-auth-token")
	if err != nil {
		fmt.Println("error here")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	fmt.Println(token)

	user, ok := s.userHandler.GetUserData(token)
	if ok != nil {
		fmt.Println("error here 2")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	fmt.Println(user.ID)

	organizationDataReq := new(model.OrganizationData)

	if ok := getOrganizationDataFromAPI(c, organizationDataReq); ok != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": ok.Error()})
		return
	}

	organizationDataResponse, err := s.storage.CreateOrganization(*organizationDataReq)

	if err != nil {
		fmt.Println("error here 3")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, organizationDataResponse)
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
	deviceId, err := strconv.Atoi(c.Param("deviceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse deviceId"})
		return
	}

	deviceData, err := s.storage.GetDeviceDataByDeviceId(deviceId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, deviceData)
}