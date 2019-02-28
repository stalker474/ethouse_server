package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var network = "kovan"
var apiKey = "fc1f68c8cd0e4755b7744e45b124ee9a"
var nodeURL = "https://" + network + ".infura.io/v3/" + apiKey
var nodeWebSocket = "wss://" + network + ".infura.io/ws"
var contractAddress = "0xe909b6925286f6774f783649788e12a2414a5c56"

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

/*
type myContext struct {
}

func (*myContext) Deadline() (time.Time, bool) {
	return time.Now(), false
}

func (myContext) Done() <-chan struct{} {
	return nil
}

func (myContext) Err() error {
	return nil
}

func (myContext) Value(key interface{}) interface{} {
	return nil
}
*/
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

	log.Println("updating payouts...")
	err = updatePayouts(contract)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	log.Println("updating other data...")
	err = updateAdditional(contract)
	if err != nil {
		log.Fatalf("Critical : %v", err)
	}

	log.Println("starting watching config...")
	go watchEvent(func(contract *CoinEmpire) error {
		logs := make(chan *CoinEmpireConfigChanged)
		sub, err := contract.CoinEmpireFilterer.WatchConfigChanged(nil, logs)
		if err != nil {
			return err
		}
		fmt.Println("Success...")
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("Watcher errored")
				return err
			case vLog := <-logs:
				handleConfigChanged(vLog)
			}
		}
	})
	log.Println("starting watching slots...")
	go watchEvent(func(contract *CoinEmpire) error {
		logs := make(chan *CoinEmpireSlotPurchased)
		sub, err := contract.CoinEmpireFilterer.WatchSlotPurchased(nil, logs)
		if err != nil {
			return err
		}
		fmt.Println("Success...")
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("Watcher errored")
				return err
			case vLog := <-logs:
				handleNewPurchase(vLog)
			}
		}
	})
	log.Println("starting watching price updates...")
	go watchEvent(func(contract *CoinEmpire) error {
		logs := make(chan *CoinEmpireSamplingPriceEnded)
		sub, err := contract.CoinEmpireFilterer.WatchSamplingPriceEnded(nil, logs)
		if err != nil {
			return err
		}
		fmt.Println("Success...")
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("Watcher errored")
				return err
			case vLog := <-logs:
				handleSamplingPriceEnded(vLog)
			}
		}
	})
	log.Println("starting watching payouts...")
	go watchEvent(func(contract *CoinEmpire) error {
		logs := make(chan *CoinEmpirePayout)
		sub, err := contract.CoinEmpireFilterer.WatchPayout(nil, logs)
		if err != nil {
			return err
		}
		fmt.Println("Success...")
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("Watcher errored")
				return err
			case vLog := <-logs:
				handlePayout(vLog)
			}
		}
	})

	log.Println("starting api on port ", port)
	err = server.Serve(port)
	if err != nil {
		log.Println("Failed to start server : ", err)
	}
	log.Println("exiting...")
}

func updatePayouts(contract *CoinEmpire) error {
	events, err := contract.CoinEmpireFilterer.FilterPayout(nil)
	if err != nil {
		return err
	}

	for events.Next() {
		handlePayout(events.Event)
	}
	events.Close()

	return nil
}

func updateSlots(contract *CoinEmpire) error {
	events, err := contract.CoinEmpireFilterer.FilterSlotPurchased(nil)
	if err != nil {
		return err
	}

	for events.Next() {
		handleNewPurchase(events.Event)
	}
	events.Close()

	return nil
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

type subscriberF func(contract *CoinEmpire) error

func watchEvent(subscriberFunc subscriberF) {

	var err error
	fmt.Println("Subscribing...")

	err = retry(func() error {
		conn, contract, err2 := initConn()
		if err2 != nil {
			log.Fatalf("Critical conn: %v", err2)
		}
		defer conn.Close()

		err2 = subscriberFunc(contract)
		if err2 != nil {
			log.Printf("Error : %v", err2)
		}
		return err2
	}, 500)
	if err != nil {
		log.Fatalf("Critical : %v", err)
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

func handlePayout(event *CoinEmpirePayout) {
	server.Database.mux.Lock()
	defer server.Database.mux.Unlock()
	log.Println("New handlePayout!")
	key := Key{event.Index.Uint64(), event.Tier}
	slot, found := server.Database.indexToSlot[key]
	if !found {
		server.Database.indexToSlot[key] = SlotData{}
	}
	slot.Payout = big.NewInt(0).Add(server.Database.indexToSlot[key].Payout, event.Amount)
	server.Database.indexToSlot[key] = slot
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
	index := event.Index.Uint64()
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
		server.Database.indexToSlot[key] = SlotData{index, tier, big.NewInt(0), big.NewInt(0), big.NewInt(0), owner}
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
