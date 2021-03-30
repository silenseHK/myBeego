package models

type GameBetting struct {
	Id             int `json:"id"`
	BettingNum     int `json:"betting_num"`
	UserId         int `json:"user_id"`
	GameId         int `json:"game_id"`
	GamePId        int `json:"game_p_id"`
	GameCXId       int `json:"game_c_x_id"`
	Money          int `json:"money"`
	Odds           int `json:"odds"`
	WinMoney       int `json:"win_money"`
	BettingTime    int `json:"betting_time"`
	SettlementTime int `json:"settlement_time"`
	Status         int `json:"status"`
	Type           int `json:"type"`
	ServiceCharge  int `json:"service_charge"`
}
