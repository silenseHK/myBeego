package models

type GameConfig struct{
	Id int `json:"id"`
	GameId int `json:"game_id"`
	Name string `json:"name"`
	Type int `json:"type"`
	Odds float64 `json:"odds"`
	GameCId int `json:"game_c_id"`
}
