package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projeto404/Api/Controllers/ControllerUteis"
	"projeto404/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/Api/Repositorys/RepositorysUsuarios"
)

func DeletarUsuarios(c *gin.Context) {
	var usuarioRequest DeletarUsuarioRequest

	idDoUsuario := ControllerUteis.UtilIdUser(c)

	ControllerUteis.BindJson(c, &usuarioRequest)

	validarSenha, err := repositoryUsuarios.ValidarSenha(*idDoUsuario, usuarioRequest.Senha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	}
	if !validarSenha {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Senha não confere"})
		return
	}

	usuarioModel := modelUsuario.UsuarioStruct{
		Senha: usuarioRequest.Senha,
	}

	mensagemDeErro, err := repositoryUsuarios.DeletarUsuario(usuarioModel, *idDoUsuario)
	if err != nil && mensagemDeErro != "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagemDeErro, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário deletado com sucesso"})
}

// TODO:UsuarioRequest Docs swagger
type DeletarUsuarioRequest struct {
	Senha string `json:"senha"`
}
