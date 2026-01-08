package inventory

import (
	"net/http"

	"bombparty.com/bombparty-api/config"
	"bombparty.com/bombparty-api/database/dbmodel"
	"bombparty.com/bombparty-api/pkg/models"
	"github.com/go-chi/render"
)

type InventoryConfig struct {
	*config.Config
}

func New(config *config.Config) *InventoryConfig {
	return &InventoryConfig{config}
}

func (config *InventoryConfig) GetUserInventory(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		render.JSON(w, r, map[string]string{"message": "Email cannot be null"})
		return
	}

	user, err := config.UserRepository.FindOne("email", userEmail)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "User not found", "error": err.Error()})
		return
	}

	inventory, err := config.InventoryRepository.FindByUser(*user)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during fetch", "error": err.Error()})
		return
	}

	response := convertToUserInventoryResponse(*user, inventory)

	render.JSON(w, r, response)

}

func (config *InventoryConfig) InitUserInventory(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		render.JSON(w, r, map[string]string{"message": "Email cannot be null"})
		return
	}

	user, err := config.UserRepository.FindOne("email", userEmail)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "User not found", "error": err.Error()})
		return
	}

	inventories, err := config.InventoryRepository.InitUserInventory(*user)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during the initialisation of the inventory", "error": err.Error()})
		return
	}

	response := convertToUserInventoryResponse(*user, inventories)

	render.JSON(w, r, response)
}

func (config *InventoryConfig) ChangeBombsAmount(w http.ResponseWriter, r *http.Request) {
	req := &models.InventoryBombAmountChangePayload{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"message": "Invalid Payload", "error": err.Error()})
		return
	}

	user, err := config.UserRepository.FindOne("email", req.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "User not found", "error": err.Error()})
		return
	}

	amount, err := config.InventoryRepository.ChangeBombsAmount(*user, req.TypeBomb, req.Amount)
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "Error during request", "error": err.Error()})
		return
	}

	response := convertToResponse(amount)
	render.JSON(w, r, response)

}

func convertToResponse(inventory *dbmodel.InventoryEntry) models.InventoryElement {
	return models.InventoryElement{
		TypeBomb: inventory.TypeBomb,
		Amount:   inventory.Amount,
	}
}

func convertToUserInventoryResponse(user dbmodel.UserEntry, inventories []*dbmodel.InventoryEntry) models.InventoryResponse {
	var elements []models.InventoryElement
	for _, ele := range inventories {
		elements = append(elements, convertToResponse(ele))
	}

	return models.InventoryResponse{
		IDUser:   user.IDUser,
		Elements: elements,
	}
}
