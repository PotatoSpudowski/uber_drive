package controllers

import (
	"encoding/json"
	"../models"
	u "../utils"
	"net/http"
)

var CreateRides = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	ride := &models.Ride{}

	err := json.NewDecoder(r.Body).Decode(ride)
	if err != nil {
		u.Respond(w, u.Message(false, "Error whle decoding request body"))
		return
	}
}

var GetRidesFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetRides(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

