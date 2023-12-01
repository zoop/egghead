package users

import (
	"egghead/app/services"

	"gorm.io/gorm"
)

// UserController represents the controller for user related actions
type UserController struct {
	UserService *services.UserService
}

// NewUserController creates a new instance of UserController.
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		UserService: services.NewUserService(db),
	}
}
