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

var tempDbFile = "db.json"

// ByRollNumber sorter
type ByRollNumber []RollData

func (a ByRollNumber) Len() int           { return len(a) }
func (a ByRollNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRollNumber) Less(i, j int) bool { return a[i].RollNumber > a[j].RollNumber }

// PersistObject blabla
type PersistObject struct {
	rollsData map[uint64]RollData
	mux       sync.Mutex
}

// NewPersistObject Create new database
func NewPersistObject() (p *PersistObject) {
	p = new(PersistObject)
	p.rollsData = make(map[uint64]RollData)
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

func (p *PersistObject) contains(rollNumber uint64) bool {
	_, exists := p.rollsData[rollNumber]
	return exists
}

func (p *PersistObject) toJSON(from uint64, to uint64) (s string, err error) {
	p.mux.Lock()
	data, err := json.Marshal(NewCache(p.rollsData, from, to))
	p.mux.Unlock()

	s = string(data[:])
	return s, err
}

func (p *PersistObject) toZJSON(from uint64, to uint64) (s string, err error) {
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
func NewCache(m map[uint64]RollData, from uint64, to uint64) (cache *Cache) {
	cache = new(Cache)
	for _, v := range m {
		if (v.RollNumber >= from) && (v.RollNumber <= to) {
			cache.List = append(cache.List, v)
		}
	}
	cache.LastUpdate = time.Now().Unix()
	sort.Sort(ByRollNumber(cache.List))

	return cache
}

// ToMap blabla
func (c *Cache) toMap() map[uint64]RollData {
	m := make(map[uint64]RollData)
	for _, v := range c.List[:] {
		m[v.RollNumber] = v
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
