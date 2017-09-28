package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vbcm/easynvest.focus/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// MetaController represents the controller for operating on the meta resource
	MetaController struct {
		session *mgo.Session
	}
)

// NewMetaController for new chart controller
func NewMetaController(s *mgo.Session) *MetaController {
	return &MetaController{s}
}

// Post : save the meta of customer
func (uc MetaController) Post(w http.ResponseWriter, r *http.Request) {
	m := models.Meta{}

	json.NewDecoder(r.Body).Decode(&m)

	m.ID = bson.NewObjectId()
	m.Start = time.Now()

	uc.session.DB("focus").C("meta").Insert(m)

	json.NewEncoder(w).Encode(m)
}
