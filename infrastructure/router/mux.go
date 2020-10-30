package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Mux struct {
	router *mux.Router
}

func NewMux() *Mux {
	return &Mux{
		router: mux.NewRouter(),
	}
}

func (m *Mux) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.router.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *Mux) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.router.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (m *Mux) SERVE(port string) {
	log.Printf("Mux HTTP server running on port %v", port)

	server := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      m.router,
	}

	log.Fatal(server.ListenAndServe())
}
