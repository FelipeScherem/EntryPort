package rotas

import (
	"github.com/gin-gonic/gin"
	controllerUsuarios "projeto404/Api/Controllers/ControllerUsuarios"
)

// ConfiguraRotas Define os endpoints
func ConfiguraRotas(router *gin.Engine) *gin.Engine {
	// API routes
	api := router.Group("/api/v1")
	{
		// Usuarios
		api.GET("/usuarios", controllerUsuarios.ListarUsuarios)
		api.GET("/usuario/:id", controllerUsuarios.BuscarUsuarios)
		api.POST("/usuario", controllerUsuarios.CriarUsuarios)
		api.PUT("/usuario/:id", controllerUsuarios.AtualizarUsuarios)
		api.DELETE("/usuario/:id", controllerUsuarios.DeletarUsuarios)
	}

	return router
}
