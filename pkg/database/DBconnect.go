package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ShowAccount godoc
// @Summary Deposit
// @Tags deposit
// @Description положить деньги на аккаунт
// @ID deposit
// @Accept json
// @Produce json
// @Param input body repo.User true "account info"
// @Success      200  {integer}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /balance [post]
const (
	host     = "127.0.0.1"
	port     = 5432 // Default port
	user     = "postgres"
	password = "qwerty"
	dbname   = "postgres"
)

func ConnectDB() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
