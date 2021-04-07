package models

type Users struct{
	Id int `json:"id"`
	Phone string `json:"phone"`
	Balance float64 `json:"balance"`
}
