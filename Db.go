package main

import (
	"math/big"
	"sync"
)

// MinMax blabla
type MinMax struct {
	Min *big.Int
	Max *big.Int
}

// Db Represents our api server
type Db struct {
	mux          sync.Mutex
	ownerToSlots map[string][]Key
	indexToSlot  map[Key]SlotData
	config       ConfigData
	configMap    map[string]**big.Int
	currentPrice *big.Int
}

// NewDb create a new database
func NewDb() *Db {
	db := new(Db)
	db.ownerToSlots = make(map[string][]Key)
	db.indexToSlot = make(map[Key]SlotData)
	db.configMap = make(map[string]**big.Int)

	db.configMap["PRICECHECK_DELAY"] = &db.config.PriceCheckDelay
	db.configMap["TIER1_PAYOUT"] = &db.config.Tier1Payout
	db.configMap["TIER2_PAYOUT"] = &db.config.Tier2Payout
	db.configMap["TIER3_PAYOUT"] = &db.config.Tier3Payout
	db.configMap["TIER1_PRICE"] = &db.config.Tier1Price
	db.configMap["TIER2_PRICE"] = &db.config.Tier2Price
	db.configMap["TIER3_PRICE"] = &db.config.Tier3Price
	db.configMap["REBUY_MULT"] = &db.config.RebuyMult
	db.configMap["REBUY_FEE"] = &db.config.RebuyFee
	db.configMap["RESELL_FEE"] = &db.config.ResellFee
	db.configMap["HOTNESS_MOD"] = &db.config.HotnessMod
	db.configMap["SPREAD"] = &db.config.Spread
	db.configMap["HOUSE_CUT"] = &db.config.HouseCut

	return db
}
