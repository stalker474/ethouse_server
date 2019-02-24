package main

import "math/big"

// Key blabla
type Key struct {
	Index uint64
	Tier  uint8
}

// SlotData blabla
type SlotData struct {
	Index         uint64 `json:"index"`
	Tier          uint8  `json:"tier"`
	Price         uint64 `json:"price"`
	Owner         string `json:"owner"`
	PreviousOwner string `json:"previous_owner"`
}

// ContractData blabla
type ContractData struct {
	PriceCheckDelay big.Int `json:"pricecheck_delay"`
	Tier1Payout     big.Int `json:"tier1_payout"`
	Tier2Payout     big.Int `json:"tier2_payout"`
	Tier3Payout     big.Int `json:"tier3_payout"`
	Tier1Price      big.Int `json:"tier1_price"`
	Tier2Price      big.Int `json:"tier2_parice"`
	Tier3Price      big.Int `json:"tier3_price"`
	RebuyMult       big.Int `json:"rebuy_mult"`
	RebuyFee        big.Int `json:"rebuy_fee"`
	ResellFee       big.Int `json:"resell_fee"`
	HotnessMod      big.Int `json:"hotness_mod"`
	Spread          big.Int `json:"spread"`
	HouseCut        big.Int `json:"house_cut"`
}
