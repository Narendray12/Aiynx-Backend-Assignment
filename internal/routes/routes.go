package routes

import (
	"github.com/gofiber/fiber/v2"

	"aiynx/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	users := app.Group("/users")

	users.Post("/", h.CreateUser)      // POST /users
	users.Get("/", h.ListUsers)        // GET /users
	users.Get("/:id", h.GetUser)       // GET /users/:id
	users.Put("/:id", h.UpdateUser)    // PUT /users/:id
	users.Delete("/:id", h.DeleteUser) // DELETE /users/:id
}
