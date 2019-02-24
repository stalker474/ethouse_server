package main

import "sync"

// Db Represents our api server
type Db struct {
	mux          sync.Mutex
	ownerToSlots map[string][]Key
	indexToSlot  map[Key]SlotData
}

// NewDb create a new database
func NewDb() *Db {
	db := new(Db)
	db.ownerToSlots = make(map[string][]Key)
	db.indexToSlot = make(map[Key]SlotData)
	return db
}
