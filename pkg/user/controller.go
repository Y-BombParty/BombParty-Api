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
// @Router /user/register [post]
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

	token, err := authentication.GenerateToken(config.JwtKey, userEntry.Email, userEntry.UserName)

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
// @Router /user/login [post]
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
		return
	}

	token, err := authentication.GenerateToken(config.JwtKey, user.Email, user.UserName)

	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the token generation", "error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

func (config *UserConfig) GetOneUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id_user")
	if id == "" {
		render.JSON(w, r, map[string]string{"message": "error with query parameters"})
		return
	}

	user, err := config.UserRepository.FindOne("id_user", id)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during fetching", "error": err.Error()})
		return
	}

	response := convertToResponse(user)

	render.JSON(w, r, response)

}

func (config *UserConfig) Update(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	req := &model.UserUpdatePayload{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the payload", "error": err.Error()})
		return
	}
	userEntry := createUserEntryFromUpdate(req)
	user, err := config.UserRepository.Update(userEntry, email)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during the update", "error": err.Error()})
		return
	}

	token, err := authentication.GenerateToken(config.JwtKey, user.Email, user.UserName)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error with the token generation", "error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

func (config *UserConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id_user")
	if id == "" {
		render.JSON(w, r, map[string]string{"message": "error with query parameters"})
		return
	}

	err := config.UserRepository.Delete(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during fetching", "error": err.Error()})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Deleted user succesfuly"})

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

func createUserEntryFromUpdate(user *model.UserUpdatePayload) *dbmodel.UserEntry {
	return &dbmodel.UserEntry{
		Email:    user.Email,
		Password: user.Password,
		UserName: user.UserName,
	}
}

func convertToResponse(user *dbmodel.UserEntry) model.UserResponse {
	response := model.UserResponse{
		IdUser:   user.IDUser,
		UserName: user.UserName,
		Email:    user.Email,
	}
	if user.IDTeam != nil {
		response.IdTeam = *user.IDTeam
	}
	return response
}
