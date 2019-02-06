package main

import (
	"encoding/binary"
	"log"
	"math/big"
	"os"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var server *Server

var ops uint64
var refreshRate int64 = 10

var saveNeeded uint32

var concurrency = 10
var sem = make(chan bool, concurrency)
var network = "kovan"
var apiKey = "fc1f68c8cd0e4755b7744e45b124ee9a"
var nodeURL = "https://" + network + ".infura.io/v3/" + apiKey
var nodeWebSocket = "wss://" + network + ".infura.io/ws/v3/" + apiKey
var contractAddress = "0x16a9dcc9098b03e3d5279e762bfcd5a9fcf8854a"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server = NewServer()
	var err error

	if server.data.load() != nil {
		log.Println("Failed to fetch database, creating one : ", err)
		log.Println("fetching new data for the first time")
	}

	log.Println("starting updating loop")
	go updateCache()

	log.Println("starting api on port ", port)
	err = server.Serve(port)
	if err != nil {
		log.Println("Failed to start server : ", err)
	}
}

func updateCache() {
	for true {
		log.Println("fetching new data...")
		changed, err := fetchNewData()
		if err != nil {
			log.Println("Failed to fetch new data, creating one : ", err)
		}
		if !changed {
			log.Println("No changes...")
		} else {
			persist()
		}
		time.Sleep(time.Duration(refreshRate) * time.Second)
	}
}

func persist() {
	server.data.save()
	server.resetCache()
	log.Println("Cache Updated")
}

func fetchNewData() (bool, error) {
	atomic.StoreUint32(&saveNeeded, 0)

	var rolls []RollData
	var err error

	conn, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	} else {
		defer conn.Close()
	}
	contract, err := NewPriceRoll(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to reach contract: %v", err)
	}

	rollEvents, err := contract.PriceRollFilterer.FilterNewRoll(&bind.FilterOpts{Start: 0, End: nil, Context: nil}, nil)
	defer rollEvents.Close()
	for rollEvents.Next() {
		rolls = append(rolls, RollData{RollNumber: rollEvents.Event.Round.Uint64()})
		rolls[len(rolls)-1].Bets = make(map[string]BetData)
	}

	log.Println("ROLLS DONE")

	for _, v := range rolls {

		roll, contains := server.data.rollsData[v.RollNumber]
		if !contains {
			//create it and append
			server.data.mux.Lock()
			server.data.rollsData[v.RollNumber] = v
			server.data.mux.Unlock()
			if err != nil {
				log.Fatal("Error :", err)
			}
			roll = v
		}
		if !roll.AllClaimed {
			//no data, update
			changed, err := updateRollData(&roll, contract)
			if err != nil {
				log.Fatal("Error :", err)
			} else if changed {
				server.data.mux.Lock()
				server.data.rollsData[v.RollNumber] = roll
				server.data.mux.Unlock()
				atomic.StoreUint32(&saveNeeded, 1)
			}
		}
	}

	return atomic.LoadUint32(&saveNeeded) == 1, nil
}

func updateRollData(roll *RollData, contract *PriceRoll) (bool, error) {
	rollIds := []*big.Int{big.NewInt(int64(roll.RollNumber))}
	var bets []BetData
	betEvents, err := contract.PriceRollFilterer.FilterBetPlaced(&bind.FilterOpts{Start: 0, End: nil, Context: nil}, rollIds, nil)
	if err != nil {
		log.Fatalf("Failed to reach contract: %v", err)
	}
	defer betEvents.Close()

	for betEvents.Next() {

		plop1 := binary.BigEndian.Uint64(betEvents.Event.Raw.Data[24:32])
		plop2 := binary.BigEndian.Uint64(betEvents.Event.Raw.Data[56:64])
		plop3 := binary.BigEndian.Uint64(betEvents.Event.Raw.Data[88:96])
		log.Println(plop1)
		log.Println(plop2)
		log.Println(plop3)
		//plop2 := binary.BigEndian.Uint64(betEvents.Event.Raw.Data[32])

		bets = append(bets, BetData{
			RollNumber:    betEvents.Event.Round.Uint64(),
			Amount:        betEvents.Event.Amount.Uint64(),
			Player:        betEvents.Event.Player.Hex(),
			ExpectedValue: uint8(plop2),
			IsUp:          plop3 == 1,
		})
	}

	return true, nil
}
