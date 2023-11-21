package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	loginUseCase      *usecases.LoginUseCase
	createUserUseCase *usecases.CreateUserUseCase
}

func NewUserController(db *mongo.Database) *UserController {
	usersRepository := repositories.NewUsersRepository(db)
	loginUseCase := usecases.NewLoginUseCase(usersRepository)
	createUserUseCase := usecases.NewCreateUserUseCase(usersRepository)

	return &UserController{loginUseCase: loginUseCase, createUserUseCase: createUserUseCase}
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Lista de usuários
// @Description Obtém uma lista de usuários
// @Tags Auth
// @Produce  json
// @Param   username     body   string     true   "Nome de usuário"
// @Param   password     body   string     true   "Senha"
// @Success 200 {object} object
// @Router /users [get]
func (c *UserController) Login(ctx *gin.Context) {
	var loginForm LoginData

	// Vincular o corpo da solicitação JSON à struct LoginForm
	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, err := c.loginUseCase.Execute(loginForm.Email, loginForm.Password)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"token": login.Token,
		"id":    login.User.Id,
		"name":  login.User.Name,
		"email": login.User.Email.Email,
		"phone": login.User.Phone.Phone,
		"type":  login.User.Type.Id,
	})
}
