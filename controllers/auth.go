package controllers

import (
	"api-produto/entity"
	"api-produto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, service service.ProdutoServiceInterface) {

	var user *entity.User
	var token entity.Token

	var userDB entity.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userDB = *service.GetUser(&user.Username)

	if user.Username != userDB.Username || user.Senha != userDB.Senha {
		c.JSON(401, gin.H{"status": "unauthorized"})
		return
	}

	token.Token = entity.USER_TOKEN

	c.JSON(200, gin.H{
		"token":  token.Token,
		"status": "success",
	})

}

func isAuth(c *gin.Context) bool {

	token := c.GetHeader("Authorization")

	if (len(token) > 0) && (token[7:] == entity.USER_TOKEN) {
		return true
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return false
	}
}
