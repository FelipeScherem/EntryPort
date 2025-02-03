package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projeto404/Api/Controllers/ControllerUteis"
	"projeto404/Api/Repositorys/RepositorysUsuarios"
)

func BuscarUsuarios(c *gin.Context) {
	idDoUsuario := ControllerUteis.UtilIdUser(c)

	usuario, err := RepositoryUsuarios.BuscarUsuario(*idDoUsuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Houve um erro ao buscar o usuario", "error": err})
		return
	}

	c.JSON(http.StatusOK, usuario)
	return
}
