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

func (apiCfg *apiconfig) handlerCreateCake(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Type         string  `json:"type"`
		Layer_number int     `json:"layer_number"`
		Tiere_number int     `json:"tiere_number"`
		Size         string  `json:"size"`
		Price        float64 `json:"price"` }

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
	}

	cake, err := apiCfg.DB.CreateCake(r.Context(), datebase.CreateCakeParams{
		ID:          uuid.New(),
		Type:        params.Type,
		LayerNumber: int16(params.Layer_number),
		TiereNumber: int16(params.Tiere_number),
		Size:        params.Size,
		Price:       strconv.FormatFloat(params.Price, 'f', 5, 64),
	})

    payload, err := databaseCakeToCake(cake)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Can't Parse float %v", err))
    }

	respondWithJson(w, 201, payload)
}

func (apiCfg *apiconfig) handlerGetCakes(w http.ResponseWriter, r *http.Request) {
	cakes, err := apiCfg.DB.GetAllCakes(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get Cakes: %e", err))
		return
	}

    payload, err := databaseCakesToCakes(cakes)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Can't Parse float: %v", err))
    }

	respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerGetCake(w http.ResponseWriter, r *http.Request) {
    cake_id, err := uuid.Parse(mux.Vars(r)["id"])
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse cake id: %v", err))
        return
    }

    cake, err := apiCfg.DB.GetCakeById(r.Context(), cake_id)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't get cake: %v", err))
        return
    }

    payload, err := databaseCakeToCake(cake)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
    }

    respondWithJson(w, 200, payload)
}

func (apiCfg *apiconfig) handlerUpdateCake(w http.ResponseWriter, r *http.Request) {
    cake_id, err := uuid.Parse(mux.Vars(r)["id"])   
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse cake id: %v", err))
        return
    }

    type parameters struct {
		Type         string  `json:"type"`
		Layer_number int     `json:"layer_number"`
		Tiere_number int     `json:"tiere_number"`
		Size         string  `json:"size"`
		Price        float64 `json:"price"`
	} 

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %e", err))
        return
	}

    cake, err := apiCfg.DB.UpdateCake(r.Context(), datebase.UpdateCakeParams{
        ID: cake_id,
        Type: params.Type,
        LayerNumber: int16(params.Layer_number),
        TiereNumber: int16(params.Tiere_number),
        Size: params.Size,
        Price: strconv.FormatFloat(params.Price, 'f', 5, 64),
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't update cake: %v", err))
        return
    }

    payload, err := databaseCakeToCake(cake)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Can't parse float: %v", err))
        return
    }

    respondWithJson(w, 200, payload) 
} 

func (apiCfg *apiconfig) handlerDeleteCake(w http.ResponseWriter, r *http.Request) {
	cake_id, err := uuid.Parse(mux.Vars(r)["id"])
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse cake id: %v", err))
        return
    }


    err = apiCfg.DB.DeleteCake(r.Context(), cake_id) 
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't delete cake: %v", err))
        return
    }

    respondWithJson(w, 200, struct{}{})
}


