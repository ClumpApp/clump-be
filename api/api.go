package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/service"

	"github.com/gofiber/fiber/v2"
)

const prefix = "/api/v1"
const id = "id"

type API struct {
	service *service.Service
}

func New(service *service.Service) *API {
	return &API{service}
}

func (obj *API) Run() {
	app := fiber.New()

	app.Use(middleware.GetCORSMiddleware())
	//app.Use(middleware.GetLimiterMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Hello": "World!"})
	})

	app.Post("/login", obj.login)
	app.Post("/signup/register", obj.signup)

	app.Use(middleware.GetJWTMiddleware())

	api := app.Group(prefix)

	api.Get("/messages", obj.getGroupMessages)
	api.Get("/users", obj.getGroupUsers)
	api.Get("/users/me", obj.getUser)

	api.Post("/messages", obj.postMessage)
	api.Post("/messages/image", obj.postImage)
	api.Post("/messages/video", obj.postVideo)
	api.Post("/messages/other", obj.postOther)

	api.Put("/users/:"+id, obj.putUser)

	api.Delete("/users/:"+id, obj.deleteUser)
	api.Delete("/messages/:"+id, obj.deletemessage)

	app.Use(obj.notFound)

	app.Listen(":8080")
}

func (obj *API) notFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Hey, there are no friends to make here.")
}

func (obj *API) getIDFromParam(c *fiber.Ctx) string {
	return c.Params(id)
}

func (obj *API) getGroupIDFromToken(c *fiber.Ctx) float64 {
	return middleware.GetGroupID(c.Locals(middleware.ContextKey))
}

func (obj *API) getUserIDFromToken(c *fiber.Ctx) float64 {
	return middleware.GetUserID(c.Locals(middleware.ContextKey))
}
