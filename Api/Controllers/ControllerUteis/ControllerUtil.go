package ControllerUteis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UtilIdUser(c *gin.Context) *int {
	idDoUsuario, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "ID inválido", "erro": err})
		return nil
	}
	return &idDoUsuario
}

// BindJson recebe um ponteiro para a struct que será populada com os dados JSON
func BindJson(c *gin.Context, request interface{}) {
	err := c.ShouldBindJSON(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao processar requisição JSON", "erro": err})
		return
	}
}
