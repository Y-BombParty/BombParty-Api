package team

import (
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type TeamConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TeamConfig {
	return &TeamConfig{configuration}
}

// CreateTeamHandler godoc
// @Summary      Create a team
// @Description  Create a new team linked to a game
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param        team  body      model.TeamRequest  true  "Team payload"
// @Success      201   {object}  dbmodel.TeamEntry
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/teams [post]
func (config *TeamConfig) CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.TeamRequest{}

	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	team := &dbmodel.TeamEntry{
		Score:  req.Score,
		Name:   req.Name,
		Color:  req.Color,
		IDGame: req.IDGame,
	}

	savedTeam, err := config.TeamRepository.Create(team)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to save team",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, savedTeam)
}

// GetAllTeamsHandler godoc
// @Summary      Get all teams
// @Description  Retrieve all teams
// @Tags         Teams
// @Produce      json
// @Success      200  {array}   dbmodel.TeamEntry
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/teams [get]
func (config *TeamConfig) GetAllTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := config.TeamRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch team",
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, teams)
}

// GetTeamByIDHandler godoc
// @Summary      Get team by ID
// @Description  Retrieve a team by its UUID
// @Tags         Teams
// @Produce      json
// @Param        id   path      string  true  "Team ID (UUID)"
// @Success      200  {object}  dbmodel.TeamEntry
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/v1/teams/{id} [get]
func (config *TeamConfig) GetTeamByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	teamID, err := uuid.Parse(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid UUID format",
		})
		return
	}

	team, err := config.TeamRepository.FindById(teamID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "team not found",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, team)

}

// UpdateTeamHandler godoc
// @Summary      Update a team
// @Description  Update an existing team
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param        id    path      string               true  "Team ID (UUID)"
// @Param        team  body      model.TeamRequest   true  "Updated team payload"
// @Success      200   {object}  dbmodel.TeamEntry
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/teams/{id} [put]
func (config *TeamConfig) UpdateTeamHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	teamID, err := uuid.Parse(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid UUID format",
		})
		return
	}

	req := &model.TeamRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	existing, err := config.TeamRepository.FindById(teamID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "team not found",
		})
		return
	}

	existing.Score = req.Score
	existing.Name = req.Name
	existing.Color = req.Color
	existing.IDGame = req.IDGame

	updatedTeam, err := config.TeamRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to update team",
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, updatedTeam)
}

// DeleteTeamHandler godoc
// @Summary      Delete a team
// @Description  Delete a team by ID
// @Tags         Teams
// @Produce      json
// @Param        id   path      string  true  "Team ID (UUID)"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/teams/{id} [delete]
func (config *TeamConfig) DeleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	teamID, err := uuid.Parse(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid UUID format",
		})
		return
	}
	team, err := config.TeamRepository.FindById(teamID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "team not found",
		})
		return
	}

	if err := config.TeamRepository.Delete(teamID, team); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to delete team",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": "team deleted",
	})
}
