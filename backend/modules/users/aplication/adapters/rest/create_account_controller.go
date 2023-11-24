package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountData struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *UserController) CreateAccount(ctx *gin.Context) {
	var createAccountForm CreateAccountData

	// Vincular o corpo da solicitação JSON à struct LoginForm
	if err := ctx.ShouldBindJSON(&createAccountForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := c.createUserUseCase.Execute(
		createAccountForm.Name,
		createAccountForm.Email,
		createAccountForm.Phone,
		createAccountForm.Password,
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email.Email,
		"phone": user.Phone.Phone,
		"type":  user.Type.Id,
	})
}
