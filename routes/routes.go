package routes

import(
	"github.com/gofiber/fiber"
)

func web(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"HelloWorld": "Welcome to Fiber!",
		})
	})
}