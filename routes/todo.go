package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmagdiel/goTodo/controllers"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
