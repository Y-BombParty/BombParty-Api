package bomb

import (
	"net/http"
	"strconv"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type BombConfig struct {
	*config.Config
}

func New(configuration *config.Config) *BombConfig {
	return &BombConfig{configuration}
}

// CreateBomb godoc
// @Summary      Create a new bomb
// @Description  Create a new bomb with the provided data
// @Tags         Bombs
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        bomb body     model.BombRequest true "Bomb data"
// @Success      201  {object} model.BombResponse
// @Failure      400  {object} map[string]string
// @Failure      500  {object} map[string]string
// @Router       /api/v1/bombs [post]
func (c *BombConfig) CreateBomb(w http.ResponseWriter, r *http.Request) {
	req := &model.BombRequest{}
	if err := render.Bind(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	bombEntry := dbmodel.BombEntry{
		Lat:      req.Lat,
		Long:     req.Long,
		TypeBomb: req.TypeBomb,
		//IdUser:   req.IdUser,
	}

	bomb, err := c.BombRepository.Create(&bombEntry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error creating bomb"})
		return
	}

	res := &model.BombResponse{
		BombId:   int(bomb.BombID),
		Lat:      bombEntry.Lat,
		Long:     bombEntry.Long,
		TypeBomb: bombEntry.TypeBomb,
		//IdUser:   bombEntry.IdUser,
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, res)
}

// GetBomb godoc
// @Summary Get a bomb by ID
// @Description Get a single bomb by its ID
// @Tags Bombs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Bomb ID"
// @Success 200 {object} model.BombResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/bombs/{id} [get]
func (c *BombConfig) GetBomb(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	bomb, err := c.BombRepository.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "Bomb not found"})
		return
	}

	res := &model.BombResponse{
		Lat:      bomb.Lat,
		Long:     bomb.Long,
		TypeBomb: bomb.TypeBomb,
		//IdUser:   bomb.IdUser,
	}
	render.JSON(w, r, res)
}

// GetAllBombs godoc
// @Summary List all bombs
// @Description Get all bombs
// @Tags Bombs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} model.BombResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/bombs [get]
func (c *BombConfig) GetAllBombs(w http.ResponseWriter, r *http.Request) {
	bombs, err := c.BombRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error fetching bombs"})
		return
	}

	responses := make([]model.BombResponse, len(bombs))
	for i, bomb := range bombs {
		responses[i] = model.BombResponse{
			Lat:      bomb.Lat,
			Long:     bomb.Long,
			TypeBomb: bomb.TypeBomb,
			//IdUser:   bomb.IdUser,
		}
	}
	render.JSON(w, r, responses)
}

// GetBombsByUserId godoc
// @Summary List bombs by user
// @Description Get all bombs created by a specific user
// @Tags Bombs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} model.BombResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/bombs/user/{userId} [get]
func (c *BombConfig) GetBombsByUserId(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(strId)
	if err != nil || userId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user id parameter"})
		return
	}

	bombs, err := c.BombRepository.FindAllByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error fetching bombs"})
		return
	}

	responses := make([]model.BombResponse, len(bombs))
	for i, bomb := range bombs {
		responses[i] = model.BombResponse{
			Lat:      bomb.Lat,
			Long:     bomb.Long,
			TypeBomb: bomb.TypeBomb,
			//IdUser:   bomb.IdUser,
		}
	}
	render.JSON(w, r, responses)
}

// UpdateBomb godoc
// @Summary Update a bomb
// @Description Update an existing bomb
// @Tags Bombs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Bomb ID"
// @Param bomb body model.BombUpdateRequest true "Bomb update data"
// @Success 200 {object} model.BombResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/bombs/{id} [put]
func (c *BombConfig) UpdateBomb(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	bomb, err := c.BombRepository.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "Bomb not found"})
		return
	}

	req := &model.BombUpdateRequest{}
	if err := render.Bind(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if req.Lat != nil {
		bomb.Lat = *req.Lat
	}
	if req.Long != nil {
		bomb.Long = *req.Long
	}
	if req.TypeBomb != nil {
		bomb.TypeBomb = *req.TypeBomb
	}

	bomb, err = c.BombRepository.Update(bomb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error updating bomb"})
		return
	}

	res := &model.BombResponse{
		Lat:      bomb.Lat,
		Long:     bomb.Long,
		TypeBomb: bomb.TypeBomb,
		//IdUser:   bomb.IdUser,
	}
	render.JSON(w, r, res)
}

// DeleteBomb godoc
// @Summary Delete a bomb
// @Description Delete a bomb by ID
// @Tags Bombs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Bomb ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/bombs/{id} [delete]
func (c *BombConfig) DeleteBomb(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	err = c.BombRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error deleting bomb"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
