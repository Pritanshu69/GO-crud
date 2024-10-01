package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "sync"
)

// Define the data model
type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// In-memory store
var store = make(map[string]Item)
var mu sync.Mutex

func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    mu.Lock()
    store[item.ID] = item
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func readItem(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    mu.Lock()
    item, exists := store[id]
    mu.Unlock()

    if !exists {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    var updatedItem Item
    err := json.NewDecoder(r.Body).Decode(&updatedItem)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
//check
    mu.Lock()
    _, exists := store[id]
    if !exists {
        mu.Unlock()
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }
    store[id] = updatedItem
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    mu.Lock()
    _, exists := store[id]
    if !exists {
        mu.Unlock()
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }
    delete(store, id)
    mu.Unlock()

    w.WriteHeader(http.StatusNoContent)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/items", createItem).Methods("POST")
    r.HandleFunc("/items/{id}", readItem).Methods("GET")
    r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
