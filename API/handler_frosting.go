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

func (apiCfg *apiconfig) handlerCreateFrosting(w http.ResponseWriter, r *http.Request) {
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

	frosting, err := apiCfg.DB.CreateFrosting(r.Context(), datebase.CreateFrostingParams{
		ID:       uuid.New(),
		Name:     params.Name,
		AddPrice: strconv.FormatFloat(params.AddPrice, 'f', 5, 64),
	})

	payload, err := databaseFrostingToFrosting(frosting)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 201, payload)
}

func (apiCfg *apiconfig) handlerGetFrostings(w http.ResponseWriter, r *http.Request) {
	frostings, err := apiCfg.DB.GetAllFrostings(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get Frosting: %e", err))
		return
	}

	payload, err := databaseFrostingsToFrostngs(frostings)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't Parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerGetFrosting(w http.ResponseWriter, r *http.Request) {
	frosting_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse frosting id: %v", err))
		return
	}

	frosting, err := apiCfg.DB.GetFrostingById(r.Context(), frosting_id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get frosting: %v", err))
		return
	}

	payload, err := databaseFrostingToFrosting(frosting)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerUpdateFrosting(w http.ResponseWriter, r *http.Request) {
	frosting_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse frosting id: %v", err))
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

	frosting, err := apiCfg.DB.UpdateFrosting(r.Context(), datebase.UpdateFrostingParams{
		ID:       frosting_id,
		Name:     params.Name,
		AddPrice: strconv.FormatFloat(params.AddPrice, 'f', 5, 64),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update frosting: %v", err))
		return
	}

	payload, err := databaseFrostingToFrosting(frosting)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerDeleteFrosting(w http.ResponseWriter, r *http.Request) {
	frosting_id, err := uuid.Parse(mux.Vars(r)["id"])
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse frosting id: %v", err))
        return
    }


    err = apiCfg.DB.DeleteFrosting(r.Context(), frosting_id)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't delete frosting: %v", err))
        return
    }

    respondWithJson(w, 200, struct{}{})
}


