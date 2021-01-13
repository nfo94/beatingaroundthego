package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Philosopher model
type Philosopher struct {
	Name          string
	Location      string
	ThoughtSchool string
	ID            string
}

type philosopherHandlers struct {
	sync.Mutex
	store map[string]Philosopher
}

func (h *philosopherHandlers) philosophers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (h *philosopherHandlers) post(w http.ResponseWriter, r *http.Request) {

}

func (h *philosopherHandlers) get(w http.ResponseWriter, r *http.Request) {
	philosophers := make([]Philosopher, len(h.store))

	i := 0
	for _, philosopher := range h.store {
		philosophers[i] = philosopher
		i++
	}

	jsonBytes, err := json.Marshal(philosophers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newPhilosopherHandlers() *philosopherHandlers {
	return &philosopherHandlers{
		store: map[string]Philosopher{
			"id1": Philosopher{
				Name:          "Friedrich Nietzsche",
				Location:      "Germany",
				ThoughtSchool: "Nihilism",
				ID:            "1",
			},
		},
	}
}

func main() {
	philosopherHandlers := newPhilosopherHandlers()
	http.HandleFunc("/philosophers", philosopherHandlers.philosophers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
