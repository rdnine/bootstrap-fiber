package routes

import(
	"github.com/gofiber/fiber"
)

func Web(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"HelloWorld": "Welcome to Fiber!",
		})
	})

	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})

	app.Get("/fiber", func(c *fiber.Ctx) {
		c.Send("Hello Fiber!")
	})
}