package authentication

import (
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/model"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	*config.Config
}

func New(config *config.Config) *AuthConfig {
	return &AuthConfig{config}
}

// Register godoc
// @Summary Créer un nouveau compte utilisateur
// @Description Enregistre un nouvel utilisateur avec un nom d'utilisateur, email et mot de passe
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.UserCreatePayload true "Informations d'inscription"
// @Success 200 {object} map[string]string "Token JWT généré"
// @Failure 400 {object} map[string]string "Erreur avec le payload"
// @Failure 500 {object} map[string]string "Erreur serveur"
// @Router /api/v1/auth/register [post]
func (config *AuthConfig) Register(w http.ResponseWriter, r *http.Request) {
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

	token, err := GenerateToken(config.JwtKey, userEntry.Email, userEntry.UserName)

	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the token generation", "error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

// Login godoc
// @Summary Connexion utilisateur
// @Description Authentifie un utilisateur avec son email et mot de passe et retourne un token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.UserLoginPayload true "Identifiants de connexion"
// @Success 200 {object} map[string]string "Token JWT généré"
// @Failure 400 {object} map[string]string "Erreur avec le payload"
// @Failure 401 {object} map[string]string "Identifiants invalides"
// @Failure 500 {object} map[string]string "Erreur serveur"
// @Router /api/v1/auth/login [post]
func (config *AuthConfig) Login(w http.ResponseWriter, r *http.Request) {
	req := &model.UserLoginPayload{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the payload", "error": err.Error()})
		return
	}

	userEntry := createUserEntryFromLogin(req)
	user, err := config.UserRepository.Login(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "error during login", "error": err.Error()})
		return
	}

	token, err := GenerateToken(config.JwtKey, user.Email, user.UserName)

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
