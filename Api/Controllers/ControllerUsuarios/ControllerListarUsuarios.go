package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	repositoryUsuarios "projeto404/Api/Repositorys/RepositorysUsuarios"
)

func ListarUsuarios(c *gin.Context) {

	usuarios, err := repositoryUsuarios.ListarUsuarios()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Houve um erro ao buscar usuarios", "error": err})
		return
	}

	c.JSON(http.StatusOK, usuarios)
	return
}
