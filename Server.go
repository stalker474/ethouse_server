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

func (s *Server) getMinMax() (*big.Int, *big.Int) {
	spread := big.NewInt(0).Div(s.Database.config.Spread, big.NewInt(2))
	precision := big.NewInt(1000000)
	m := big.NewInt(0).Sub(precision, spread)
	m2 := big.NewInt(0).Add(precision, spread)
	min := big.NewInt(0).Div(big.NewInt(0).Mul(s.Database.currentPrice, m), precision)
	max := big.NewInt(0).Div(big.NewInt(0).Mul(s.Database.currentPrice, m2), precision)
	return min, max
}

func (s *Server) getHotnessModifier(price *big.Int, hotness *big.Int) *big.Int {
	min, max := s.getMinMax()
	med := big.NewInt(0).Div(big.NewInt(0).Sub(max, min), big.NewInt(2))
	value := big.NewInt(0).Sub(big.NewInt(0).Sub(price, min), med)
	absValue := big.NewInt(0).Abs(value)
	precision := big.NewInt(1000000)

	a := big.NewInt(0).Mul(absValue, precision)
	b := big.NewInt(0).Div(a, med)
	c := big.NewInt(0).Mul(b, hotness)
	result := big.NewInt(0).Div(c, precision)
	//clamp
	if result.Cmp(hotness) == 1 {
		result = hotness
	}
	return big.NewInt(0).Sub(precision, result)
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
		precision := big.NewInt(1000000)
		basePrice := s.Database.config.Tier1Price
		mult := s.Database.config.RebuyMult
		min, max := s.getMinMax()

		minCity := min.Uint64() / 100
		maxCity := max.Uint64() / 100

		citiesList := make([]SlotData, maxCity-minCity+1)

		for i := minCity; i <= maxCity; i++ {
			searchKey := Key{i * 100, 0}
			val := s.Database.indexToSlot[searchKey]
			val.Index = searchKey.Index
			if val.Price == nil {
				mod := s.getHotnessModifier(big.NewInt(0).SetUint64(searchKey.Index), s.Database.config.HotnessMod)
				price := big.NewInt(0)
				if mod.Cmp(big.NewInt(0)) != 0 {
					price.Mul(big.NewInt(0).Div(basePrice, precision), mod)
				}
				val.Purchase = big.NewInt(0).Add(basePrice, price)
			} else {
				val.Purchase = big.NewInt(0).Mul(big.NewInt(0).Div(val.Price, precision), mult)
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

	http.HandleFunc("/slots", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		keysIndex, ok := r.URL.Query()["index"]
		if !ok {
			fmt.Fprintln(w, "Missing index param")
			return
		}
		keysTier, ok := r.URL.Query()["tier"]
		if !ok {
			fmt.Fprintln(w, "Missing tier param")
			return
		}
		index, err := strconv.Atoi(keysIndex[0])
		if err != nil {
			fmt.Fprintln(w, "Invalid index param")
			return
		}
		tier, err := strconv.Atoi(keysTier[0])
		if err != nil || tier > 2 || tier < 1 {
			fmt.Fprintln(w, "Invalid tier param")
			return
		}
		s.Database.mux.Lock()
		min := uint64(index / 100 * 10)
		max := uint64(index/100*10 + 9)
		precision := big.NewInt(1000000)
		basePrice := s.Database.config.Tier2Price
		if tier == 2 {
			basePrice = s.Database.config.Tier3Price
		}
		mult := s.Database.config.RebuyMult

		list := make([]SlotData, max-min+1)

		for i := min; i <= max; i++ {
			searchKey := Key{i * 10, 0}
			val := s.Database.indexToSlot[searchKey]
			val.Index = searchKey.Index
			val.Tier = uint8(tier)
			if val.Price == nil {
				mod := s.getHotnessModifier(big.NewInt(0).SetUint64(searchKey.Index), s.Database.config.HotnessMod)
				price := big.NewInt(0)
				if mod.Cmp(big.NewInt(0)) != 0 {
					price.Mul(big.NewInt(0).Div(basePrice, precision), mod)
				}
				val.Purchase = big.NewInt(0).Add(basePrice, price)
			} else {
				val.Purchase = big.NewInt(0).Mul(big.NewInt(0).Div(val.Price, precision), mult)
			}

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
