package main

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	Message string `json:"message"`
}

func main() {
	// hello worldを返す
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := Res{Message: "hello world!"}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
