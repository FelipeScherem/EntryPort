package rotas

import (
	"github.com/gin-gonic/gin"
	"projeto404/Api/Controllers"
	controllerUsuarios "projeto404/Api/Controllers/ControllerUsuarios"
	"projeto404/Api/Middleware"
)

// ConfiguraRotas Define os endpoints
func ConfiguraRotas(router *gin.Engine) *gin.Engine {

	router.POST("/login", Controllers.UsuarioLogin)

	api := router.Group("/api/v1")
	api.Use(Middleware.Autenticar())
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
