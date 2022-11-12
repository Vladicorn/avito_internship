package rout

import (
	"github.com/Vladicorn/avito_internship/pkg/controllers"

	_ "github.com/Vladicorn/avito_internship/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	//	_ "github.com/zhashkevych/todo-app/docs"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	api.Post("/balance", controllers.PutBalance)
	api.Get("/balance/:id", controllers.GetBalance)
	api.Post("/StartContract", controllers.StartContract)
	api.Post("/FinishContract", controllers.FinishContract)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	api.Post("/Report", controllers.ReportPrint)

}
