package controllers

import (
	"log"

	"github.com/Vladicorn/avito_internship/pkg/database"
	"github.com/Vladicorn/avito_internship/pkg/repo"
	"github.com/gofiber/fiber/v2"
)

// @Summary Положить деньги
// @Tags Пользователь
// @Description Метод начисления средств на баланс.
// @ID deposit
// @Accept json
// @Produce json
// @Param input body repo.User true "send money"
// @Router       /api/balance [post]
func PutBalance(c *fiber.Ctx) error {

	user := new(repo.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return err
	}

	check, err := database.CheckUser(user.Id)
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": "Error"})
	}

	if check == true {
		err = database.UpdateBalanceUser(user)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{"status": "Error"})
		}
	} else {
		err = database.PutBalanceUser(user)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{"status": "Error"})
		}
	}

	return c.JSON(fiber.Map{"status": "Success", "Ваш Id": user.Id})
}

// @Summary Получить баланс
// @Tags Пользователь
// @Description Метод получения баланса пользователя
// @ID balance
// @Accept json
// @Produce json
// @Param input body repo.User true "check for id"
// @Router       /api/balance/{id} [get]
func GetBalance(c *fiber.Ctx) error {

	user, err := database.GetBalanceUser(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": "Error"})
	}

	return c.JSON(fiber.Map{"status": "Success", "Ваш баланс:": user.Balance})
}
