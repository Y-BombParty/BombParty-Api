package game

import (
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type GameConfig struct {
	*config.Config
}

func New(configuration *config.Config) *GameConfig {
	return &GameConfig{configuration}
}

// PostHandler godoc
// @Summary      Create a new Game
// @Description  Creates a new Game entry in the database
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        Game  body      model.GameRequest  true  "Game creation payload"
// @Security     BearerAuth
// @Success      200    {object}  model.GameResponse
// @Failure      400    {object}  map[string]string  "Invalid request payload"
// @Failure      500    {object}  map[string]string  "Failed to create Game"
// @Router       /api/v1/game [post]
func (config *GameConfig) PostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request
	req := &model.GameRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"Error": "Invalid Game Post request payload. " + err.Error()})
		return
	}

	// Check if there is nil value
	if err := req.ValidateCreate(); err != nil {
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}

	// Convert the requested data into dbmodel.GameEntry type for the "Create" function
	gameEntry := &dbmodel.GameEntry{
		CenterLatitude:  *req.CenterLatitude,
		CenterLongitude: *req.CenterLongitude,
		Size:            *req.Size,
		StartingDate:    *req.StartingDate,
		EndingDate:      *req.EndingDate}

	// Request the DB to Create the informations
	entries, err := config.GameRepository.Create(gameEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to Create Game"})
		return
	}

	// Set up to a dedicated type for the response
	res := &model.GameResponse{
		CenterLatitude:  entries.CenterLatitude,
		CenterLongitude: entries.CenterLongitude,
		Size:            entries.Size,
		StartingDate:    entries.StartingDate,
		EndingDate:      entries.EndingDate,
		Teams:           []*model.TeamResponse{}}

	render.JSON(w, r, res)
}

// GetByIdHandler godoc
// @Summary      Get game by ID
// @Description  Retrieves a specific game from the database by its ID, including associated teams
// @Tags         games
// @Produce      json
// @Param        id   path      string  true  "game ID"
// @Security     BearerAuth
// @Success      200  {object}  model.GameResponse
// @Failure      404  {object}  map[string]string  "Game not found"
// @Failure      500  {object}  map[string]string  "Failed to find specific game"
// @Router       /api/v1/game/{id} [get]
func (config *GameConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id"})
		return
	}

	uuid, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id"})
		return
	}

	// Request the DB to Get the needed informations
	entries, err := config.GameRepository.FindById(uuid)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to Find specific game"})
		return
	}

	// Set up to a dedicated type for the response
	var teams []*model.TeamResponse

	for _, team := range entries.Teams {
		teams = append(teams, &model.TeamResponse{
			Score:  team.Score,
			Name:   team.Name,
			Color:  team.Color,
			//IDGame: team.IDGame
			})
	}

	res := &model.GameResponse{
		CenterLatitude:  entries.CenterLatitude,
		CenterLongitude: entries.CenterLongitude,
		Size:            entries.Size,
		StartingDate:    entries.StartingDate,
		EndingDate:      entries.EndingDate,
		Teams:           teams}

	render.JSON(w, r, res)
}

// GetAlldHandler godoc
// @Summary      Get all games
// @Description  Retrieves a list of all games from the database
// @Tags         games
// @Produce      json
// @Security     BearerAuth
// @Success      200     {array}   model.GameResponse
// @Failure      500     {object}  map[string]string  "Failed to retrieve games"
// @Router       /api/v1/games [get]
func (config *GameConfig) GetAlldHandler(w http.ResponseWriter, r *http.Request) {

	entries, err := config.GameRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to Find Games"})
		return
	}

	// Set up to a dedicated type for the response
	var res []*model.GameResponse

	for _, game := range entries {
		var teams []*model.TeamResponse
		for _, team := range game.Teams {
			teams = append(teams, &model.TeamResponse{
				Score:  team.Score,
				Name:   team.Name,
				Color:  team.Color,
				//IDGame: team.IDGame
			})
		}

		res = append(res,
			&model.GameResponse{
				CenterLatitude:  game.CenterLatitude,
				CenterLongitude: game.CenterLongitude,
				Size:            game.Size,
				StartingDate:    game.StartingDate,
				EndingDate:      game.EndingDate,
				Teams:           teams})
	}

	render.JSON(w, r, res)
}

// UpdateHandler godoc
// @Summary      Update a game
// @Description  Updates an existing game's information in the database
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        id     path      string                  true  "Game ID"
// @Param        game  body      model.GameRequest  true  "Game update payload"
// @Security     BearerAuth
// @Success      200    {object}  model.GameResponse
// @Failure      400    {object}  map[string]string  "Invalid request payload"
// @Failure      404    {object}  map[string]string  "Game not found"
// @Failure      500    {object}  map[string]string  "Failed to update game"
// @Router       /api/v1/games/{id} [patch]
func (config *GameConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id"})
		return
	}

	uuid, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id"})
		return
	}

	// Get the request
	req := &model.GameRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"Error": "Invalid Game Update request payload. " + err.Error()})
		return
	}

	// Check if there is at least one new value
	if err := req.ValidateUpdate(); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Check if the linked cat id existe
	existingGame, err := config.GameRepository.FindById(uuid)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Game not found in the DB"})
		return
	}

	// Convert the requested data into dbmodel.GameEntry type for the "Update" function
	gameEntry := existingGame

	// Fill the gameEntry with the good values
	if req.CenterLatitude != nil {
		gameEntry.CenterLatitude = *req.CenterLatitude
	}
	if req.CenterLongitude != nil {
		gameEntry.CenterLongitude = *req.CenterLongitude
	}
	if req.Size != nil {
		gameEntry.Size = *req.Size
	}
	if req.StartingDate != nil {
		gameEntry.StartingDate = *req.StartingDate
	}
	if req.EndingDate != nil {
		gameEntry.EndingDate = *req.EndingDate
	}

	// Request the DB to Update the informations
	entries, err := config.GameRepository.Update(gameEntry, uuid)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to Update Game"})
		return
	}

	// Set up to a dedicated type for the response
	var teams []*model.TeamResponse
	for _, team := range entries.Teams {
		teams = append(teams,
			&model.TeamResponse{
				Score:  team.Score,
				Name:   team.Name,
				Color:  team.Color,
				//IDGame: team.IDGame,
			})
	}

	res := &model.GameResponse{
		CenterLatitude:  entries.CenterLatitude,
		CenterLongitude: entries.CenterLongitude,
		Size:            entries.Size,
		StartingDate:    entries.StartingDate,
		EndingDate:      entries.EndingDate,
		Teams:           teams}

	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete a game
// @Description  Deletes a game from the database by its ID
// @Tags         games
// @Produce      json
// @Param        id   path      string  true  "Game ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "Game deleted successfully"
// @Failure      404  {object}  map[string]string  "Game not found"
// @Failure      500  {object}  map[string]string  "Failed to delete game"
// @Router       /api/v1/games/{id} [delete]
func (config *GameConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id"})
		return
	}

	// Request the DB to Delete the informations
	errDelete := config.GameRepository.DeleteById(id)
	if errDelete != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to Delete Game"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Game deleted successfully"})
}
