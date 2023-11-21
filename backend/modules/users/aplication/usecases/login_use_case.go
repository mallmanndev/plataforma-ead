package usecases

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
)

type LoginUseCase struct {
	UsersRepository ports.UsersRepository
}

func NewLoginUseCase(UsersRepository ports.UsersRepository) *LoginUseCase {
	return &LoginUseCase{UsersRepository: UsersRepository}
}

type LoginUseCaseOutput struct {
	User  *entities.User
	Token string
}

func (u *LoginUseCase) Execute(Email string, Password string) (*LoginUseCaseOutput, error) {
	email, err := value_objects.NewEmailAddress(Email)
	if err != nil {
		return nil, err
	}

	user, err := u.UsersRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Error on find user!")
	}

	if user == nil {
		return nil, errors.New("User not found!")
	}

	if err := user.ComparePassword(Password); err != nil {
		return nil, errors.New("Invalid password!")
	}

	var sampleSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["id"] = user.Id
	claims["email"] = user.Email.Email

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return nil, err
	}

	return &LoginUseCaseOutput{User: user, Token: tokenString}, nil
}
