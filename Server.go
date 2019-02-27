package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

// Server Represents our api server
type Server struct {
	Database *Db
}

// NewServer blabla
func NewServer() (s *Server) {
	s = new(Server)
	s.Database = NewDb()
	return s
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableDecorators(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/plain")
}

func enableDecoratorsGz(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Encoding", "gzip")
	(*w).Header().Set("Content-Type", "text/javascript")
}

func (s *Server) getMinMax() (uint64, uint64) {
	spread := s.Database.config.Spread.Uint64() / 2
	precision := uint64(1000000)
	price := s.Database.currentPrice.Uint64()
	min := price * (precision - spread) / precision
	max := price * (precision + spread) / precision
	return min,max
}

func getHotnessModifier(price uint64, hotness uint64) uint64 {
	min, max := getMinMax()
	med := (max - min) / 2
	value = price - min - med
	absValue = math.Abs(value);
	precision = 1000000;
	
	result := absValue * precision / med * hotness / precision;
	//clamp
	result := math.Min(result, hotness);

	return precision - result;
}

// Serve start the server on port port
func (s *Server) Serve(port string) error {

	http.HandleFunc("/owned", func(w http.ResponseWriter, r *http.Request) {
		//enableDecoratorsGz(&w)
		enableCors(&w)
		keysWho, okWho := r.URL.Query()["who"]
		if !okWho {
			fmt.Fprintln(w, "Missing who param")
			return
		}
		s.Database.mux.Lock()
		slots, exists := s.Database.ownerToSlots[keysWho[0]]
		if !exists {
			slots = make([]Key, 0)
		} else {

		}

		data, err := json.Marshal(slots)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			str := string(data[:])
			fmt.Fprintln(w, str)
		}
		s.Database.mux.Unlock()
	})

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		s.Database.mux.Lock()

		data, err := json.Marshal(s.Database.config)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			str := string(data[:])
			fmt.Fprintln(w, str)
		}
		s.Database.mux.Unlock()
	})

	http.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		s.Database.mux.Lock()
		spread := s.Database.config.Spread.Uint64() / 2
		precision := uint64(1000000)
		price := s.Database.currentPrice.Uint64()
		min := price * (precision - spread) / precision
		max := price * (precision + spread) / precision

		minCity := min / 100
		maxCity := max / 100

		citiesList := make([]SlotData, maxCity-minCity+1)

		for i := minCity; i <= maxCity; i++ {
			searchKey := Key{i * 100, 0}
			val := s.Database.indexToSlot[searchKey]
			val.Index = searchKey.Index
			if val.Price == 0 {
				price := big.NewInt(0).Mul(s.Database.config.Tier1Price, s.Database.config.HotnessMod)
				val.Price = 
			}
			citiesList[i-minCity] = val
		}

		data, err := json.Marshal(citiesList)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			str := string(data[:])
			fmt.Fprintln(w, str)
		}
		s.Database.mux.Unlock()
	})

	http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		keysIndex, ok := r.URL.Query()["index"]
		if !ok {
			fmt.Fprintln(w, "Missing index param")
			return
		}
		index, err := strconv.Atoi(keysIndex[0])
		s.Database.mux.Lock()
		min := uint64(index / 100 * 10)
		max := uint64(index/100*10 + 9)

		list := make([]SlotData, max-min+1)

		for i := min; i <= max; i++ {
			searchKey := Key{i * 10, 0}
			val := s.Database.indexToSlot[searchKey]
			val.Index = searchKey.Index
			val.Tier = 1
			list[i-min] = val
		}

		data, err := json.Marshal(list)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			str := string(data[:])
			fmt.Fprintln(w, str)
		}
		s.Database.mux.Unlock()
	})

	http.HandleFunc("/lots", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		keysIndex, ok := r.URL.Query()["index"]
		if !ok {
			fmt.Fprintln(w, "Missing index param")
			return
		}
		index, err := strconv.Atoi(keysIndex[0])
		s.Database.mux.Lock()
		min := uint64(index / 100 * 100)
		max := uint64(index/100*100 + 9)

		list := make([]SlotData, max-min+1)

		for i := min; i <= max; i++ {
			searchKey := Key{i, 0}
			val := s.Database.indexToSlot[searchKey]
			val.Index = searchKey.Index
			val.Tier = 2
			list[i-min] = val
		}

		data, err := json.Marshal(list)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			str := string(data[:])
			fmt.Fprintln(w, str)
		}
		s.Database.mux.Unlock()
	})

	return http.ListenAndServe(":"+port, nil)
}
