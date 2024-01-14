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

func (apiCfg *apiconfig) handlerCreateFilling(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string  `json:"name"`
		Price float64 `json:"price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
	}

	filling, err := apiCfg.DB.CreateFilling(r.Context(), datebase.CreateFillingParams{
		ID:       uuid.New(),
		Name:     params.Name,
		Price: strconv.FormatFloat(params.Price, 'f', 5, 64),
	})

	payload, err := databaseFillingToFilling(filling)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 201, payload)
}

func (apiCfg *apiconfig) handlerGetFillings(w http.ResponseWriter, r *http.Request) {
	filling, err := apiCfg.DB.GetAllFillings(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get Frosting: %e", err))
		return
	}

	payload, err := databaseFillingsToFilligs(filling)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't Parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerGetFilling(w http.ResponseWriter, r *http.Request) {
	filling_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse filling id: %v", err))
		return
	}

	filling, err := apiCfg.DB.GetFillingById(r.Context(), filling_id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get filling: %v", err))
		return
	}

	payload, err := databaseFillingToFilling(filling)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerUpdateFilling(w http.ResponseWriter, r *http.Request) {
	filling_id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse filling id: %v", err))
		return
	}

	type parameters struct {
		Name     string  `json:"name"`
		Price float64 `json:"add_price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
		return
	}

	filling, err := apiCfg.DB.UpdateFilling(r.Context(), datebase.UpdateFillingParams{
		ID:       filling_id,
		Name:     params.Name,
		Price: strconv.FormatFloat(params.Price, 'f', 5, 64),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update filling: %v", err))
		return
	}

	payload, err := databaseFillingToFilling(filling)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
		return
	}

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerDeleteFilling(w http.ResponseWriter, r *http.Request) {
	filling_id, err := uuid.Parse(mux.Vars(r)["id"])
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse filling id: %v", err))
        return
    }


    err = apiCfg.DB.DeleteFilling(r.Context(), filling_id)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't delete filling: %v", err))
        return
    }

    respondWithJson(w, 200, struct{}{})
}
