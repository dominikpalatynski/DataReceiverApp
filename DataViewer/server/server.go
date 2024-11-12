package server

import (
	"data_viewer/config"
	"data_viewer/model"
	"data_viewer/storage"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type APIServer struct {
	router *gin.Engine
	storage storage.Storage
	config config.Config
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
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
	s.router.GET("/ws", s.wsHandler)
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

func (s *APIServer) wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("WebSocket upgrade error:", err)
        return
    }
    // defer conn.Close()

    // Odbieranie wiadomości od klienta z konfiguracją zapytania
    _, message, err := conn.ReadMessage()
    if err != nil {
        log.Println("ReadMessage error:", err)
        return
    }

	queryParams := new(model.QueryParams)

	if err := json.Unmarshal(message, &queryParams); err != nil {
        log.Println("JSON Unmarshal error:", err)
        conn.WriteMessage(websocket.TextMessage, []byte("Invalid request format"))
        return
    }

	for {
        time.Sleep(5 * time.Second)

		data, err := s.storage.FetchData(queryParams)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

        if err := conn.WriteJSON(data); err != nil {
            log.Println("WriteJSON error:", err)
            break
        }
    }

}