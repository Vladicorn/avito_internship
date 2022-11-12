package database

import (
	"github.com/Vladicorn/avito_internship/pkg/repo"

	"log"
	"time"
	//"strconv"
)

func CheckUser(id uint64) (bool, error) {
	user := repo.User{}

	rows, err := DB.Query("SELECT id,balance FROM avito_users WHERE id= $1", id)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Balance)

	}

	if user.Id != 0 {
		return true, nil
	}

	return false, nil
}

func PutBalanceUser(user *repo.User) error {
	rows, err := DB.Query("INSERT INTO avito_users (balance) VALUES ($1) RETURNING id ", user.UpdateMoney)
	for rows.Next() {
		rows.Scan(&user.Id)
	}
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateBalanceUser(user *repo.User) error {
	_, err := DB.Query("UPDATE  avito_users SET balance = balance + $1 WHERE id = $2 ", user.UpdateMoney, user.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetBalanceUser(id string) (repo.User, error) {
	user := repo.User{}
	rows, err := DB.Query("SELECT id,balance FROM avito_users WHERE id= $1", id)
	if err != nil {
		log.Println(err)
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Balance)

	}
	return user, nil
}

func StartContractBalance(balance *repo.AvitoBalance) error {
	var date_layout = "2006-01-02"
	today := time.Now()
	//	t, _ := time.Parse(date_layout, today)
	t := today.Format(date_layout)
	log.Println(t)
	balance.Start = true
	balance.Finish = false
	balance.Error = false
	_, err := DB.Query("INSERT INTO avito_balance (id_user,id_service,id_order,price,start,finish,error,time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", balance.IdUser, balance.IdService, balance.IdOrder, balance.Price, balance.Start, balance.Finish, balance.Error, today)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FinishContractBalance(balance *repo.AvitoBalance) error {
	balance.Start = false
	balance.Finish = true
	_, err := DB.Query("UPDATE  avito_balance SET start = $1, finish = $2 WHERE id_user = $3 AND id_service = $4 AND id_order = $5 AND price = $6 ", balance.Start, balance.Finish, balance.IdUser, balance.IdService, balance.IdOrder, balance.Price)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ErrorContractBalance(balance *repo.AvitoBalance) error {
	balance.Start = false
	balance.Finish = false
	balance.Error = true
	_, err := DB.Query("UPDATE  avito_balance SET start = $1, error = $2 WHERE id_user = $3 AND id_service = $4 AND id_order = $5 AND price = $6 ", balance.Start, balance.Error, balance.IdUser, balance.IdService, balance.IdOrder, balance.Price)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetSumService(id uint64, dateStart, dateFinish string) (float64, error) {
	var sum float64

	row := DB.QueryRow("SELECT SUM(price) FROM avito_balance WHERE id_service= $1 AND finish = 'true' AND time >=$2 AND time < $3", id, dateStart, dateFinish)

	err := row.Scan(&sum)
	if err != nil {
		log.Println(err)
		return sum, err
	}
	return sum, nil
}
