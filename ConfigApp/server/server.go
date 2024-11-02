package server

import (
	"ConfigApp/model"
	"ConfigApp/storage"
	"log"
	"net/http"
	"os"

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
		AllowOrigins:     []string{"http://localhost:4242"},
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
	s.router.GET("/devices", s.test)
	s.router.POST("/devices", s.addDeviceInfo)
}

func (s *APIServer) test(c *gin.Context) {
	
    c.JSON(http.StatusOK, "server response test")
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