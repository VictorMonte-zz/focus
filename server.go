package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vbcm/easynvest.focus/models"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/chart", func(w http.ResponseWriter, req *http.Request) {
		c := models.Chart{
			Labels:       []string{"jan", "fev", "abr", "maio"},
			Custody:      []float32{1, 2, 3, 4},
			Accomplished: []float32{1, 2, 3, 4},
			Meta:         []float32{1, 2, 3, 4},
		}

		json.NewEncoder(w).Encode(c)

	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
