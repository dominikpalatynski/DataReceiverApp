package server

import (
	"data_viewer/config"
	"data_viewer/model"
	"data_viewer/storage"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
	storage storage.Storage
	config config.Config
}

func NewAPIServer(storage storage.Storage, config config.Config) *APIServer{
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
	queryParams := new(model.QueryParams)

	if ok := getQueryParamsFromApi(c, queryParams); ok != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": ok.Error()})
		return
	}

	data, err := s.storage.FetchData(queryParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)

}