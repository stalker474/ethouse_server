package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

var server *Server
var network = "kovan"
var apiKey = "fc1f68c8cd0e4755b7744e45b124ee9a"
var nodeURL = "https://" + network + ".infura.io/v3/" + apiKey
var nodeWebSocket = "wss://" + network + ".infura.io/ws/v3/" + apiKey
var contractAddress = "0x63E91B45dcBB01eD5CD1eAb2AcfD658AD2994d51"

// InternalError blabla
type InternalError struct {
	Path string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Path)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server = NewServer()
	var err error
	conn, err := ethclient.Dial(nodeWebSocket)
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	} else {
		defer conn.Close()
	}
	contract, err := NewCoinEmpire(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to reach contract: %v", err)
	}

	log.Println("starting...")
	go update(contract, server.Database)

	log.Println("starting api on port ", port)
	err = server.Serve(port)
	if err != nil {
		log.Println("Failed to start server : ", err)
	}
	log.Println("exiting...")
}

func update(contract *CoinEmpire, db *Db) {
	purchaseEvents, err := contract.CoinEmpireFilterer.FilterSlotPurchased(&bind.FilterOpts{Start: 0, End: nil, Context: nil})
	if err != nil {
		log.Println("Failed to fetch init events : ", err)
	}

	for purchaseEvents.Next() {
		handleNewPurchase(purchaseEvents.Event, db)
	}

	purchaseEvents.Close()

	logs := make(chan *CoinEmpireSlotPurchased)

	err = new(InternalError)
	var sub event.Subscription
	for err != nil {
		sub, err = contract.CoinEmpireFilterer.WatchSlotPurchased(&bind.WatchOpts{Start: nil, Context: nil}, logs)
		if err != nil {
			log.Println("Failed to subscribe to future events : ", err)
			conn, err := ethclient.Dial(nodeWebSocket)
			if err != nil {
				log.Fatalf("Failed to init node: %v", err)
			} else {
				defer conn.Close()
			}
			contract, err = NewCoinEmpire(common.HexToAddress(contractAddress), conn)
			if err != nil {
				log.Fatalf("Failed to reach contract: %v", err)
			}
		}
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			handleNewPurchase(vLog, db)
		}
	}
}

func remove(s []Key, i int) []Key {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func handleContractValuesUpdate(contract *CoinEmpire, db *Db) {
	db.mux.Lock()

	db.mux.Unlock()
}

func handleNewPurchase(event *CoinEmpireSlotPurchased, db *Db) {
	db.mux.Lock()
	log.Println("New event!")
	index := event.Price.Uint64()
	tier := event.Tier
	owner := event.By.Hex()
	prevOwner := event.From.Hex()
	fmt.Printf("Slot %d, new data \n", index)
	key := Key{index, tier}
	_, found := db.indexToSlot[key]
	if found {
		log.Println("Slot already exists in DB")
		//find index of the key
		slots := db.ownerToSlots[prevOwner]
		if len(slots) > 0 {
			var i int
			var element Key
			for i, element = range slots {
				if element.Index == index && element.Tier == tier {
					break
				}
			}
			db.ownerToSlots[prevOwner] = remove(db.ownerToSlots[prevOwner], i)
		}

	} else {
		db.indexToSlot[key] = SlotData{index, tier, 0, owner, event.From.Hex()}
		if db.ownerToSlots[owner] == nil {
			db.ownerToSlots[owner] = make([]Key, 0)
		}
	}

	db.ownerToSlots[owner] = append(db.ownerToSlots[owner], key)
	db.mux.Unlock()
}
