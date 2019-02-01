package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"time"
)

var tempDbFile = "mle_db.json"

// ByRaceNumber sorter
type ByRaceNumber []RollData

func (a ByRaceNumber) Len() int           { return len(a) }
func (a ByRaceNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaceNumber) Less(i, j int) bool { return a[i].RaceNumber > a[j].RaceNumber }

// PersistObject blabla
type PersistObject struct {
	rollsData map[uint32]RollData
	mux       sync.Mutex
}

// NewPersistObject Create new database
func NewPersistObject() (p *PersistObject) {
	p = new(PersistObject)
	p.rollsData = make(map[uint32]RollData)
	return p
}

// Save blabla
func (p *PersistObject) save() error {
	p.mux.Lock()
	cache := NewCache(p.rollsData, 0, 99999)
	p.mux.Unlock()
	resp, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tempDbFile, resp, 0644)
}

// Load blabla
func (p *PersistObject) load() error {
	jsonText, err := ioutil.ReadFile(tempDbFile)
	if err != nil {
		return err
	}
	var cache Cache
	err = json.Unmarshal(jsonText, &cache)
	p.mux.Lock()
	p.rollsData = cache.toMap()
	p.mux.Unlock()
	return nil
}

func (p *PersistObject) contains(rollNumber uint32) bool {
	_, exists := p.rollsData[rollNumber]
	return exists
}

func (p *PersistObject) toJSON(from uint32, to uint32) (s string, err error) {
	p.mux.Lock()
	data, err := json.Marshal(NewCache(p.racesData, from, to))
	p.mux.Unlock()

	s = string(data[:])
	return s, err
}

func (p *PersistObject) toZJSON(from uint32, to uint32) (s string, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	data, err := p.toJSON(from, to)

	_, err = zw.Write([]byte(data))
	if err != nil {
		return s, err
	}

	if err := zw.Close(); err != nil {
		return s, err
	}

	return buf.String(), err
}

// NewCache blabla
func NewCache(m map[uint32]RollData, from uint32, to uint32) (cache *Cache) {
	cache = new(Cache)
	for _, v := range m {
		if (v.RollNumber >= from) && (v.RollNumber <= to) {
			cache.List = append(cache.List, v)
		}
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Sort(ByRaceNumber(cache.List))

	return cache
}

// NewCacheLight blabla
func NewCacheLight(m map[uint32]RaceData, from uint32, to uint32) (cache *CacheLight) {
	cache = new(CacheLight)
	for _, v := range m {
		if (v.RaceNumber >= from) && (v.RaceNumber <= to) {
			cache.List = append(cache.List, RaceDataLight{
				ContractID:      v.ContractID,
				Date:            v.Date,
				RaceDuration:    v.RaceDuration,
				BettingDuration: v.BettingDuration,
				EndTime:         v.EndTime,
				RaceNumber:      v.RaceNumber,
				Version:         v.Version,
				WinnerHorses:    v.WinnerHorses,
				Odds:            v.Odds,
				Volume:          v.Volume,
				Refunded:        v.Refunded,
				Active:          v.Active,
				Complete:        v.Complete})
		}
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Slice(cache.List, func(i, j int) bool {
		return cache.List[i].RaceNumber > cache.List[j].RaceNumber
	})

	return cache
}

// ToMap blabla
func (c *Cache) toMap() map[uint32]RaceData {
	m := make(map[uint32]RaceData)
	for _, v := range c.List[:] {
		m[v.RaceNumber] = v
	}
	return m
}

// Contains blabla
func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

// Contains 2 blabla
func Contains2(s []Withdraw, e string) bool {
	for _, a := range s {
		if strings.Compare(a.To, e) == 0 {
			return true
		}
	}
	return false
}
