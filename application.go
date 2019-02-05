package main

import (
	"log"
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
var contractAddress = "0xc7e907a241ebee49a667f2fefb047ff81d3ee86b"

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
		if !fetchNewData() {
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
	firstBlock := uint64(5400000)
	lastBlock := uint64(5900000)
	rollEvents, err := contract.PriceRollFilterer.FilterNewRoll(&bind.FilterOpts{Start: firstBlock, End: &lastBlock, Context: nil})
	defer rollEvents.Close()
	for rollEvents.Next() {
		rollEvents.Event.
	}

	log.Println("DONE")

	return false, atomic.LoadUint32(&saveNeeded) == 1
}
