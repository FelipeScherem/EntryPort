package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "projeto404/Api/Controllers/ControllerUteis"
	repositoryUsuarios "projeto404/Api/Repositorys/RepositorysUsuarios"
	util2 "projeto404/Api/Uteis"
)

// UsuarioLogin
func UsuarioLogin(c *gin.Context) {
	var userLoginRequest UserLoginRequest
	util.BindJson(c, &userLoginRequest)

	// Verifica se o usuário existe no banco
	usuario, err := repositoryUsuarios.BuscarUsuario(loginRequest.Email)
	if err != nil || usuario.Senha != loginRequest.Senha {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Email ou senha inválidos"})
		return
	}

	// Gera o token JWT
	token, err := util2.GerarTokenJWT(usuario.ID, usuario.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type UserLoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
