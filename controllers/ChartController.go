package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vbcm/easynvest.focus/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// ChartController represents the controller for operating on the Chart resource
	ChartController struct {
		session *mgo.Session
	}
)

// NewChartController for new chart controller
func NewChartController(s *mgo.Session) *ChartController {
	return &ChartController{s}
}

// Get : get chart data
func (uc ChartController) Get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	c := models.Chart{}

	if err := uc.session.DB("focus").C("chart").FindId(oid).One(&c); err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(c)
}

// Feed : create new chart data
func (uc ChartController) Feed(w http.ResponseWriter, r *http.Request) {

	c := models.Chart{
		ID:           bson.NewObjectId(),
		Labels:       []string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun"},
		Custody:      []float32{310, 630, 900, 1200, 1500, 1800},
		Accomplished: []float32{300, 600, 900, 1200, 1500, 1800},
		Meta:         []float32{300, 600, 900, 1200, 1500, 1800, 2100, 2400, 2700, 3000, 3300, 3600},
	}

	json.NewDecoder(r.Body).Decode(&c)

	uc.session.DB("focus").C("chart").Insert(c)

	json.NewEncoder(w).Encode(c)
}
