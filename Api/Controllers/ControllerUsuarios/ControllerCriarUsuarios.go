package controllerUsuarios

import (
	"net/http"
	"projeto404/Api/Controllers/ControllerUteis"
	"projeto404/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/Api/Repositorys/RepositorysUsuarios"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CriarUsuarios Valida e insere dados dos usuarios
func CriarUsuarios(c *gin.Context) {
	var usuarioRequest UsuarioRequest

	ControllerUteis.BindJson(c, &usuarioRequest)

	if validaSeCamposVazio(c, usuarioRequest) {
		return
	}

	var usuarioDeletado bool
	usuarioDeletado, err := repositoryUsuarios.VerificarSoftdelete(usuarioRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	}

	if usuarioDeletado {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Email já cadastrado, deseja reativar seu usuario?"})
		return
	}

	mensagem, statusSenha := ControllerUteis.ValidarSenha(usuarioRequest.Senha)
	if !statusSenha {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagem})
		return
	}

	usuarioModel := modelUsuario.UsuarioStruct{
		Nome:             usuarioRequest.Nome,
		Email:            usuarioRequest.Email,
		Telefone:         usuarioRequest.Telefone,
		Senha:            usuarioRequest.Senha,
		DataDeNascimento: usuarioRequest.DataDeNascimento,
		Foto:             usuarioRequest.Foto,
	}

	mensagemDeErro, err := repositoryUsuarios.CriarUsuario(usuarioModel)
	if err != nil && mensagemDeErro != "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagemDeErro, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário criado com sucesso"})
}

// True se houver campos vazios
//
//	False se estiverem todos preenchidos
func validaSeCamposVazio(c *gin.Context, usuarioRequest UsuarioRequest) bool {

	var camposVazios []string
	usuarioRequestReflect := reflect.ValueOf(usuarioRequest)

	// Percorre os campos do struct
	for i := 0; i < usuarioRequestReflect.NumField(); i++ {
		valorDoCampo := usuarioRequestReflect.Field(i)
		nomeDoCampo := usuarioRequestReflect.Type().Field(i).Tag.Get("json")

		if valorDoCampo.Kind() == reflect.String && valorDoCampo.Len() == 0 {
			// Ignorar CPF e CNPJ nesta etapa, pois um deles pode ser vazio quando o outro estiver preenchido
			if nomeDoCampo != "telefone" {
				camposVazios = append(camposVazios, nomeDoCampo)
			}
		}
	}

	// Construir mensagem de erro se houver campos vazios
	if len(camposVazios) > 0 {
		camposStr := strings.Join(camposVazios, ", ")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos " + camposStr + " não podem estar vazios"})
		return true
	}

	return false
}

// TODO:UsuarioRequest Docs swagger
type UsuarioRequest struct {
	Nome             string    `json:"nome"`
	Email            string    `json:"email"`
	Telefone         string    `json:"telefone"`
	Senha            string    `json:"senha"`
	DataDeNascimento time.Time `json:"dataDeNascimento"`
	Foto             string    `json:"foto"`
}
