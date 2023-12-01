package routes

import (
	"egghead/app/controllers/users/v1"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRouter(route fiber.Router, db *gorm.DB) {
	//Initialize the user controller
	userController := users.NewUserController(db)

	// Define rotues for the user controller
	// /product/:productID - removing this
	route.Post("/user/register", userController.RegisterUser)
	route.Delete("/user/:userID/archieve", userController.ArchiveUser)
}
