package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Vladicorn/avito_internship/pkg/database"
	"github.com/Vladicorn/avito_internship/pkg/repo"
	"github.com/gofiber/fiber/v2"
)

// @Summary Начать услугу
// @Tags Услуга
// @Description Метод резервирования средств с основного баланса на отдельном счете
// @ID start
// @Accept json
// @Produce json
// @Param input body repo.AvitoBalance true "start"
// @Router       /api/StartContract [post]
func StartContract(c *fiber.Ctx) error {

	balance := new(repo.AvitoBalance)

	if err := c.BodyParser(balance); err != nil {
		log.Println(err, 1)
		return err
	}

	user, err := database.GetBalanceUser(strconv.Itoa(int(balance.IdUser)))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": "Error"})
	}

	if user.Balance-balance.Price >= 0 {

		err := database.StartContractBalance(balance)
		if err != nil {
			log.Println(err, 1)
			return c.JSON(fiber.Map{"status": "Error"})
		}

		user.UpdateMoney = (balance.Price) * -1
		err = database.UpdateBalanceUser(&user)
		if err != nil {
			log.Println(err, 1)
			return c.JSON(fiber.Map{"status": "Error"})
		}

		return c.JSON(fiber.Map{"status": "Success", "message": "Ваш платеж обрабатывается"})
	} else {
		return c.JSON(fiber.Map{"status": "Error", "message": "У вас не хватает денег"})
	}

	return c.JSON(fiber.Map{"status": "Error", "message": "Попробуйте снова"})
}

// @Summary Закончить услугу
// @Tags Услуга
// @Description Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. если valid =false, то деньги вернутся пользователю
// @ID finish
// @Accept json
// @Produce json
// @Param input body repo.AvitoBalance true "start"
// @Router       /api/FinishContract [post]
func FinishContract(c *fiber.Ctx) error {

	balance := new(repo.AvitoBalance)

	if err := c.BodyParser(balance); err != nil {
		log.Println(err, 1)
		return err
	}

	if balance.Valid == true {
		err := database.FinishContractBalance(balance)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{"status": "Error"})
		}
		return c.JSON(fiber.Map{"status": "Success", "message": "Ваш платеж прошел!"})
	} else {
		err := database.ErrorContractBalance(balance)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{"status": "Error"})
		}
		user := repo.User{
			Id:          balance.IdUser,
			UpdateMoney: balance.Price,
		}

		err = database.UpdateBalanceUser(&user)
		if err != nil {
			log.Println(err, 1)
			return c.JSON(fiber.Map{"status": "Error"})
		}
		return c.JSON(fiber.Map{"status": "Error", "message": "Услуга не прошла, возвращаем деньги"})
	}

	return c.JSON(fiber.Map{"status": "Error", "message": "Попробуйте снова"})
}

// @Summary Получить отчет
// @Tags Бухгалтерия
// @Description Бухгалтерия раз в месяц просит предоставить сводный отчет по всем пользователем, с указанием сумм выручки по каждой из предоставленной услуги для расчета и уплаты налогов.
// @ID report
// @Accept json
// @Produce json
// @Param input body repo.Report true "start"
// @Router       /api/Report [post]
func ReportPrint(c *fiber.Ctx) error {
	date := repo.Report{}

	if err := c.BodyParser(&date); err != nil {
		log.Println(err)
		return err
	}

	dateStart := date.YearStart + "-" + date.MonthStart + "-" + "01"
	dateFinish := date.YearFinish + "-" + date.MonthFinish + "-" + "01"
	sum, err := database.GetSumService(date.IdService, dateStart, dateFinish)
	if err != nil {
		log.Println(err)
		return err
	}
	sumS := fmt.Sprint(sum)
	csvfile := [][]string{
		{"Id service -  ", strconv.Itoa(int(date.IdService))},
		{"Price sum service -  ", sumS},
		{"Start period -  ", dateStart},
		{"End period -  ", dateFinish},
	}

	f, err := os.Create("./data/balance.csv")
	defer f.Close()
	if err != nil {
		log.Println("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range csvfile {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": "Error"})
	}

	return c.JSON(fiber.Map{"status": "Success", "link": "./data/balance.csv"})
	//return c.Download("./data/balance.csv") -работа через фронт

}
