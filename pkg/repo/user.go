package repo

import (
	"time"
)

// @Description User account information
// @Description with user id and updatemoney
type User struct {
	Id          uint64  `json:"id" example:"1"`
	Balance     float64 `json:"balance"`
	UpdateMoney float64 `json:"updatemoney" example:"100.0"`
}

// @Description AvitoBalance account information

type AvitoBalance struct {
	Id        uint64     `json:"id"`
	IdUser    uint64     `json:"iduser" example:"1"`
	IdService uint64     `json:"idservice" example:"1"`
	IdOrder   uint64     `json:"idorder" example:"1"`
	Price     float64    `json:"price" example:"10.0"`
	Start     bool       `json:"-" `
	Finish    bool       `json:"-" `
	Error     bool       `json:"-" `
	Valid     bool       `json:"valid" example:"true"` //Если ошибка пришла с другого сервиса
	Time      *time.Time `json:"-" `
}

// @Description Report sum for service

type Report struct {
	IdService   uint64 `json:"idservice" example:"1"`
	YearStart   string `json:"yearstart" example:"2022"`
	MonthStart  string `json:"monthstart" example:"10"`
	YearFinish  string `json:"yearfinish" example:"2022"`
	MonthFinish string `json:"monthfinish" example:"12"`
}
