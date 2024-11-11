package server

import (
	"data_viewer/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
	config config.Config
}

func NewAPIServer(config config.Config) *APIServer{
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
		config: config,
	}
	return server
}

func (s *APIServer) Run() {

	s.registerRoutes()

	s.router.Run(":"+ s.config.Server.Port)
}

func (s *APIServer) registerRoutes() {
	s.router.GET("/fetchData", s.fetchDeviceData)
}

func (s *APIServer) fetchDeviceData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}