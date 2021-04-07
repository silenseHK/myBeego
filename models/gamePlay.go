package models

type GamePlay struct{
	Id int `json:"id"`
	Number int64 `json:"number"`
	GameId int `json:"game_id"`
	PrizeNumber string `json:"prize_number"`
	Status int `json:"status"`
	PrizeTime int `json:"prize_time"`
	EndTime int64 `json:"end_time"`
	StartTime int `json:"start_time"`
	Type int `json:"type"`
	Winmoney int `json:"winmoney"`
	BMoney int `json:"b_money"`
	Lostmoney int `json:"lostmoney"`
	PtMoney int `json:"pt_money"`
	IsQueue int `json:"is_queue"`
	IsStatus int `json:"is_status"`
}
