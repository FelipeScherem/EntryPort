package server

import (
	"log"
	"projeto404/Api/Server/Routes"

	"github.com/gin-gonic/gin"
)

// Server Define atributos do servidor
type Server struct {
	port   string
	server *gin.Engine
}

// RodarServidor Instancia o servidor
func ServerParameters() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

// Run abre o servidor e configura rotas
func (s *Server) Run() {
	router := rotas.ConfiguraRotas(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
