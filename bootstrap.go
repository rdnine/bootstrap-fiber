package main

import(
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"./routes"
)

func main() {

	//Declare the GO html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	routes.Web(app)

	app.Listen(3000)
}