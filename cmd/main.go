package main

import (
	"github.com/Vladicorn/avito_internship/pkg/database"
	"github.com/Vladicorn/avito_internship/pkg/rout"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"log"
)

// @title          Тестовое задание на позицию стажёра-бэкендера 2022
// @verison        1.0
// @description    Микросервис для работы с балансом пользователей
// @termsOfService http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  fiber@swagger.io
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           127.0.0.1:8000
// @BasePath       /

var IP string = "http://127.0.0.1"

func main() {

	err := database.ConnectDB()
	if err != nil {
		log.Println(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     IP,
		AllowHeaders:     "Origin, Content-Type, Accept,Access-Control-Allow-Origin",
	}))

	rout.Setup(app)

	app.Listen(":8000")
}
