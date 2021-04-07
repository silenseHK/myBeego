package models

type UserBalanceLogs struct{
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Type int8 `json:"type"`
	DqBalance float64 `json:"dq_balance"`
	WcBalance float64 `json:"wc_balance"`
	Time int64 `json:"time"`
	Msg string `json:"msg"`
	Money float64 `json:"money"`
	IsFirstRecharge int8 `json:"is_first_recharge"`
	AdminId int `json:"admin_id"`
}
