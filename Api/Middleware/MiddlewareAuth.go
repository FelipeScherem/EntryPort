package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "projeto404/Api/Uteis"
)

func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token não informado"})
			c.Abort()
			return
		}

		// Remove "Bearer " caso esteja presente
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := util.ValidarToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err, "mensagem": "Token Invalido"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
