package server

import (
	"ConfigApp/model"
	"ConfigApp/storage"
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
}

func NewAPIServer(storage storage.Storage) *APIServer{
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
	}
	return server
}

func (s *APIServer) Run() {

	s.registerRoutes()

	s.router.Run(":"+ os.Getenv("PORT"))
}

func (s *APIServer) registerRoutes() {
	s.router.POST("/devices", s.addDeviceInfo)
	s.router.GET("/devices/:orgId", s.getDeviceInfosByOrgId)
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