package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vbcm/easynvest.focus/controllers"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	router := mux.NewRouter()

	cc := controllers.NewChartController(getSession())
	mc := controllers.NewMetaController(getSession())

	router.HandleFunc("/chart/{id}", cc.Get).Methods("GET")
	router.HandleFunc("/chart/feed", cc.Feed).Methods("POST")
	router.HandleFunc("/meta", mc.Post).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://focus:focus@ds153494.mlab.com:53494/focus")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
