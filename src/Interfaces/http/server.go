package http

import (
	"log"
	config "voucher-redeem-api/src/Commons"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServer() *Server {
	cnfg := config.LoadConfig()

	port := cnfg["app_port"]
	if port == "" {
		port = "8080"
	}

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Length", "Content-Type"},
		AllowAllOrigins: true,
	}))

	return &Server{
		Engine: engine,
		Port:   port,
	}
}

func (s *Server) Run() {
	log.Printf("Server running on port %s", s.Port)
	err := s.Engine.Run(":" + s.Port)
	if err != nil {
		return
	}
}
