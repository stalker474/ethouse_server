package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Server Represents our api server
type Server struct {
	data     *PersistObject
	cacheMux sync.Mutex
	cache    map[string]string
}

// NewServer blabla
func NewServer() (s *Server) {
	s = new(Server)
	s.data = NewPersistObject()
	s.resetCache()
	return s
}

func (s *Server) resetCache() {
	s.cacheMux.Lock()
	s.cache = make(map[string]string)
	s.cacheMux.Unlock()
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

func getFromAndTo(r *http.Request) (from uint64, to uint64, err error) {
	keysFrom, okFrom := r.URL.Query()["from"]
	keysTo, okTo := r.URL.Query()["to"]

	from = 0
	to = 9999999999999999999

	if okFrom && (len(keysFrom) > 0) {
		val, err := strconv.ParseInt(keysFrom[0], 10, 64)
		if err != nil {
			return 0, 0, err
		}

		from = uint64(val)
	}
	if okTo && (len(keysTo) > 0) {
		val, err := strconv.ParseInt(keysTo[0], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		to = uint64(val)
	}

	return from, to, nil
}

// Serve start the server on port port
func (s *Server) Serve(port string) error {

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		enableDecoratorsGz(&w)
		enableCors(&w)
		fmt.Fprintln(w, "plop")
	})

	return http.ListenAndServe(":"+port, nil)
}
