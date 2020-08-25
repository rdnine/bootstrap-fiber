package main

import(
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {

	//Declare the GO html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"HelloWorld": "Welcome to Fiber!",
		})
	})

	app.Listen(3000)
}