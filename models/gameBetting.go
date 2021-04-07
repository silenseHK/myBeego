package models

type GameBetting struct {
	Id             int `json:"id"`
	BettingNum     int64 `json:"betting_num"`
	UserId         int `json:"user_id"`
	GameId         int `json:"game_id"`
	GamePId        int `json:"game_p_id"`
	GameCXId       int `json:"game_c_x_id"`
	Money          float64 `json:"money"`
	Odds           float64 `json:"odds"`
	WinMoney       int `json:"win_money"`
	BettingTime    int64 `json:"betting_time"`
	SettlementTime int `json:"settlement_time"`
	Status         int `json:"status"`
	Type           int `json:"type"`
	ServiceCharge  float64 `json:"service_charge"`
}
