package user

import (
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/authentication"
	"bombparty.com/bombparty-api/pkg/model"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	*config.Config
}

func New(config *config.Config) *UserConfig {
	return &UserConfig{config}
}

func (config *UserConfig) Register(w http.ResponseWriter, r *http.Request) {
	req := &model.UserCreatePayload{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the payload", "error": err.Error()})
		return
	}

	password := []byte(req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the hashing", "error": err.Error()})
		return
	}
	req.Password = string(hashedPassword)

	userEntry := createUserEntryFromRegister(req)
	userEntry, err = config.UserRepository.Register(userEntry)

	token, err := authentication.GenerateToken(config.JwtKey, userEntry.UserName, userEntry.Email)

	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the token generation", "error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

func (config *UserConfig) Login(w http.ResponseWriter, r *http.Request) {
	req := &model.UserLoginPayload{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the payload", "error": err.Error()})
		return
	}

	userEntry := createUserEntryFromLogin(req)
	user, err := config.UserRepository.Login(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "error during login", "error": err.Error()})
	}

	token, err := authentication.GenerateToken(config.JwtKey, user.UserName, user.Email)

	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the token generation", "error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": token})

}

func createUserEntryFromRegister(user *model.UserCreatePayload) *dbmodel.UserEntry {
	return &dbmodel.UserEntry{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func createUserEntryFromLogin(user *model.UserLoginPayload) *dbmodel.UserEntry {
	return &dbmodel.UserEntry{
		Email:    user.Email,
		Password: user.Password,
	}
}
