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
	Winmoney float64 `json:"winmoney"`
	BMoney float64 `json:"b_money"`
	Lostmoney float64 `json:"lostmoney"`
	PtMoney float64 `json:"pt_money"`
	IsQueue int `json:"is_queue"`
	IsStatus int `json:"is_status"`
}
