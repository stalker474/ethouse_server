package main

// RollData blabla
type RollData struct {
	RollNumber uint64             `json:"number"`
	Timestamp  uint64             `json:"timestamp"`
	Pool       uint64             `json:"pool"`
	State      uint8              `json:"state"`
	Coin       uint8              `json:"coin"`
	ResultRng  uint8              `json:"result_rng"`
	IsUp       bool               `json:"is_up"`
	AllClaimed bool               `json:"all_claimed"`
	Bets       map[string]BetData `json:"bets"`
}

// BetData blabla
type BetData struct {
	RollNumber    uint64 `json:"number"`
	Amount        uint64 `json:"amount"`
	ExpectedValue uint8  `json:"expected_value"`
	IsUp          bool   `json:"is_up"`
	Player        string `json:"player"`
}

// Cache blabla
type Cache struct {
	List       []RollData `json:"list"`
	LastUpdate int64      `json:"last_update"`
}
