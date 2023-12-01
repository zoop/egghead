package health

import (
	"egghead/app/services"

	"gorm.io/gorm"
)

// HealthController represents the controller for health-related actions.
type HealthController struct {
	HealthService *services.HealthService
}

// NewHealthController creates a new instance of HealthController.
func NewHealthController(db *gorm.DB) *HealthController {
	return &HealthController{
		HealthService: services.NewHealthService(db),
	}
}
