package web

import(
	"github.com/gofiber/fiber"
)

func router(app *fiber) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"HelloWorld": "Welcome to Fiber!",
		})
	})
}