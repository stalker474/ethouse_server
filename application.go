package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

var network = "kovan"
var apiKey = "fc1f68c8cd0e4755b7744e45b124ee9a"
var nodeURL = "https://" + network + ".infura.io/v3/" + apiKey
var nodeWebSocket = "wss://" + network + ".infura.io/ws/v3/" + apiKey
var contractAddress = "0x3a86e26B30AC48F5ec06aF6BC8E3F4898F569886"

var server *Server

// InternalError blabla
type InternalError struct {
	Path string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Path)
}

func initConn() (*ethclient.Client, *CoinEmpire, error) {
	conn, err := ethclient.Dial(nodeWebSocket)
	if err != nil {
		return nil, nil, err
	}

	contract, err := NewCoinEmpire(common.HexToAddress(contractAddress), conn)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}
	return conn, contract, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server = NewServer()

	conn, contract, err := initConn()
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}
	defer conn.Close()

	log.Println("starting...")

	log.Println("updating slots...")
	err = updateSlots(contract)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	log.Println("updating config...")
	err = updateConfig(contract)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	log.Println("updating other data...")
	err = updateAdditional(contract)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	log.Println("starting watching config...")
	go watchConfig()
	log.Println("starting watching slots...")
	go watchSlots()
	log.Println("starting watching price updates...")
	go watchSamplingPriceEnded()

	log.Println("starting api on port ", port)
	err = server.Serve(port)
	if err != nil {
		log.Println("Failed to start server : ", err)
	}
	log.Println("exiting...")
}

func updateSlots(contract *CoinEmpire) error {
	purchaseEvents, err := contract.CoinEmpireFilterer.FilterSlotPurchased(&bind.FilterOpts{Start: 0, End: nil, Context: nil})
	if err != nil {
		return err
	}

	for purchaseEvents.Next() {
		handleNewPurchase(purchaseEvents.Event)
	}

	purchaseEvents.Close()

	return nil
}

func watchSamplingPriceEnded() {
	var conn *ethclient.Client
	var contract *CoinEmpire
	var err error

	logs := make(chan *CoinEmpireSamplingPriceEnded)

	fmt.Println("Subscribing SamplingPriceEnded")

	var sub event.Subscription
	err = retry(func() error {
		conn, contract, err = initConn()
		if err != nil {
			log.Fatalf("Critical : %v", err)
		}

		sub, err = contract.CoinEmpireFilterer.WatchSamplingPriceEnded(&bind.WatchOpts{Start: nil, Context: nil}, logs)
		if err != nil {
			log.Fatalf("Critical : %v", err)
		}
		return err
	}, 10)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	defer conn.Close()

	fmt.Println("Subscribed")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			handleSamplingPriceEnded(vLog)
		}
	}
}

func watchSlots() {
	var conn *ethclient.Client
	var contract *CoinEmpire
	var err error

	logs := make(chan *CoinEmpireSlotPurchased)

	fmt.Println("Subscribing WatchSlotPurchased...")

	var sub event.Subscription
	err = retry(func() error {
		conn, contract, err = initConn()
		if err != nil {
			log.Fatalf("Critical : %v", err)
		}

		sub, err = contract.CoinEmpireFilterer.WatchSlotPurchased(&bind.WatchOpts{Start: nil, Context: nil}, logs)
		if err != nil {
			log.Fatalf("Critical : %v", err)
		}
		return err
	}, 10)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	defer conn.Close()

	fmt.Println("Subscribed")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			handleNewPurchase(vLog)
		}
	}
}

func updateConfig(contract *CoinEmpire) error {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	for _, name := range ConfigNames {
		result, err := contract.Config(nil, stringToBytes32(name))
		if err != nil {
			return err
		}
		*server.Database.configMap[name] = result
	}

	return nil
}

func updateAdditional(contract *CoinEmpire) error {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	result, err := contract.CurrentPrice(nil)
	if err != nil {
		return err
	}
	server.Database.currentPrice = result
	return nil
}

func watchConfig() {
	var conn *ethclient.Client
	var contract *CoinEmpire
	var err error
	logs := make(chan *CoinEmpireConfigChanged)

	fmt.Println("Subscribing WatchConfigChanged...")
	var sub event.Subscription
	err = retry(func() error {
		conn, contract, err = initConn()
		if err != nil {
			log.Fatalf("Critical : %v", err)
		}

		sub, err = contract.CoinEmpireFilterer.WatchConfigChanged(&bind.WatchOpts{Start: nil, Context: nil}, logs)
		if err != nil {
			log.Printf("Error : %v", err)
		}
		return err
	}, 10)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	defer conn.Close()

	fmt.Println("Subscribed")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			handleConfigChanged(vLog)
		}
	}
}

type fetchFunc func() error

func retry(myFunc fetchFunc, maxRetry uint) error {
	var counter uint
	for (myFunc() != nil) && counter < maxRetry {
		counter++
	}
	if counter >= maxRetry {
		return errors.New("Used all retry credits")
	}
	return nil
}

func remove(s []Key, i int) []Key {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func stringToBytes32(s string) [32]byte {
	var bytes [32]byte
	for i, c := range s {
		bytes[i] = byte(c)
	}
	return bytes
}

func bytes32ToString(bytes [32]byte) string {
	return string(bytes[:32])
}

func handleConfigChanged(event *CoinEmpireConfigChanged) {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	log.Println("New handleConfigChanged!")
	name := bytes32ToString(event.Name)
	*server.Database.configMap[name] = event.NewValue
}

func handleNewPurchase(event *CoinEmpireSlotPurchased) {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	log.Println("New handleNewPurchase!")
	index := event.Price.Uint64()
	tier := event.Tier
	owner := event.By.Hex()
	prevOwner := event.From.Hex()
	fmt.Printf("Slot %d, new data \n", index)
	key := Key{index, tier}
	_, found := server.Database.indexToSlot[key]
	if found {
		log.Println("Slot already exists in DB")
		//find index of the key
		slots := server.Database.ownerToSlots[prevOwner]
		if len(slots) > 0 {
			var i int
			var element Key
			for i, element = range slots {
				if element.Index == index && element.Tier == tier {
					break
				}
			}
			server.Database.ownerToSlots[prevOwner] = remove(server.Database.ownerToSlots[prevOwner], i)
		}

	} else {
		server.Database.indexToSlot[key] = SlotData{index, tier, 0, owner, event.From.Hex()}
		if server.Database.ownerToSlots[owner] == nil {
			server.Database.ownerToSlots[owner] = make([]Key, 0)
		}
	}

	server.Database.ownerToSlots[owner] = append(server.Database.ownerToSlots[owner], key)
}

func handleSamplingPriceEnded(event *CoinEmpireSamplingPriceEnded) {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	log.Println("New handleSamplingPriceEnded!")
	//update min max and stuff
	server.Database.currentPrice = event.Price
}
