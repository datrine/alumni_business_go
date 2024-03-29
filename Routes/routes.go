package routes

import (
	handlers "github.com/datrine/alumni_business/Handlers"
	middleware "github.com/datrine/alumni_business/Middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) fiber.Router {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
	api := app.Group("/api")
	api.Post("/users/register", handlers.Register)
	api.Post("/login/basic", handlers.Login)
	api.Get("/generate_payment_link", handlers.GeneratePaystackLink)
	authApi := api.Group("/auth", middleware.Auth)
	authApi.Post("/me/edit", handlers.UpdateUserProfile)
	authApi.Get("/me", handlers.GetMyProfile)
	authApi.Post("/password/change", handlers.ChangePassword)
	postsApi := api.Group("/posts")
	postsApi.Post("/", middleware.Auth, handlers.CreatePost)
	postsApi.Put("/:id", middleware.Auth, handlers.EditPost)
	postsApi.Delete("/:id", middleware.Auth, handlers.DeletePost)
	postsApi.Get("/", handlers.GetPosts)
	return api
}
