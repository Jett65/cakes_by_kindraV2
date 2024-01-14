package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jett65/cakes_by_kindraV2/internal/datebase"
)

func (apiCfg *apiconfig) handlerCreateFlavor(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string  `json:"name"`
		AddPrice float64 `json:"add_price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
	}

	flavor, err := apiCfg.DB.CreateFlavor(r.Context(), datebase.CreateFlavorParams{
		ID:       uuid.New(),
		Name:     params.Name,
		AddPrice: strconv.FormatFloat(params.AddPrice, 'f', 5, 64),
	})

	payload, err := databaseFlavorToFlavor(flavor)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 201, payload)
}

func (apiCfg *apiconfig) handlerGetFlavors(w http.ResponseWriter, r *http.Request) {
	flavors, err := apiCfg.DB.GetAllFlavors(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get Flavor: %e", err))
		return
	}

	payload, err := databaseFlavorsToFlavors(flavors)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't Parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerGetFlavor(w http.ResponseWriter, r *http.Request) {
	flavor_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse flavor id: %v", err))
		return
	}

	flavor, err := apiCfg.DB.GetFlavorById(r.Context(), flavor_id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get flavor: %v", err))
		return
	}

	payload, err := databaseFlavorToFlavor(flavor)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerUpdateFlavor(w http.ResponseWriter, r *http.Request) {
	flavor_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse flavor id: %v", err))
		return
	}

	type parameters struct {
		Name     string  `json:"name"`
		AddPrice float64 `json:"add_price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
		return
	}

	flavor, err := apiCfg.DB.UpdateFlavor(r.Context(), datebase.UpdateFlavorParams{
		ID:       flavor_id,
		Name:     params.Name,
		AddPrice: strconv.FormatFloat(params.AddPrice, 'f', 5, 64),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update flavor: %v", err))
		return
	}

	payload, err := databaseFlavorToFlavor(flavor)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerDeleteFlavor(w http.ResponseWriter, r *http.Request) {
	flavor_id, err := uuid.Parse(mux.Vars(r)["id"])
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse flavor id: %v", err))
        return
    }


    err = apiCfg.DB.DeleteFlavor(r.Context(), flavor_id)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't delete flavor: %v", err))
        return
    }

    respondWithJson(w, 200, struct{}{})
}

