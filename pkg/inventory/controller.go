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

	var elements []models.InventoryElement
	for _, ele := range inventories {
		elements = append(elements, convertToResponse(ele))
	}

	response := models.InventoryResponse{
		IDUser:   user.IDUser,
		Elements: elements,
	}

	render.JSON(w, r, response)
}

func convertToResponse(inventory *dbmodel.InventoryEntry) models.InventoryElement {
	return models.InventoryElement{
		TypeBomb: inventory.TypeBomb,
		Amount:   inventory.Amount,
	}
}
