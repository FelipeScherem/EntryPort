package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projeto404/Api/Controllers/ControllerUteis"
	"projeto404/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/Api/Repositorys/RepositorysUsuarios"
	"time"
)

// AtualizarUsuarios Valida e atualiza dados dos usuarios
func AtualizarUsuarios(c *gin.Context) {

	var usuarioRequest AtualizarUsuarioRequest

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

	if usuarioRequest.NovaSenha != "" {
		mensagem, statusSenha := ControllerUteis.ValidarSenha(usuarioRequest.NovaSenha)
		if !statusSenha {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagem})
			return
		}
		usuarioRequest.Senha = usuarioRequest.NovaSenha
	}

	usuarioModel := modelUsuario.UsuarioStruct{
		Nome:             usuarioRequest.Nome,
		Email:            usuarioRequest.Email,
		Telefone:         usuarioRequest.Telefone,
		Senha:            usuarioRequest.Senha,
		DataDeNascimento: usuarioRequest.DataDeNascimento,
		Foto:             usuarioRequest.Foto,
	}

	mensagemDeErro, err := repositoryUsuarios.AtualizarUsuario(usuarioModel, *idDoUsuario)
	if err != nil && mensagemDeErro != "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagemDeErro, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário atualizado com sucesso"})
}

// TODO:AtualizarUsuarioRequest Docs swagger
type AtualizarUsuarioRequest struct {
	Nome             string    `json:"nome"`
	Email            string    `json:"email"`
	Telefone         string    `json:"telefone"`
	Senha            string    `json:"senha"`
	NovaSenha        string    `json:"novaSenha"`
	DataDeNascimento time.Time `json:"dataDeNascimento"`
	Foto             string    `json:"foto"`
}
