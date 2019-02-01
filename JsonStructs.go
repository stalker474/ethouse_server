package main

// RollData blabla
type RollData struct {
	RollNumber uint32 `json:"number"`
	Timestamp  uint64 `json:"timestamp"`
	Pool       uint64 `json:"pool"`
	State      uint8  `json:"state"`
	Coin       uint8  `json:"coin"`
	ResultRng  uint8  `json:"result_rng"`
	IsUp       bool   `json:"is_up"`
}

// Cache blabla
type Cache struct {
	List       []RollData `json:"list"`
	LastUpdate int64      `json:"last_update"`
}
