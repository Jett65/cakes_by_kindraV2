//TODO: Create logs for each request

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jett65/cakes_by_kindraV2/internal/datebase"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiconfig struct {
	DB *datebase.Queries
}

func main() {
	godotenv.Load(".env")

	portString := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := apiconfig{
		DB: datebase.New(conn),
	}

	router := mux.NewRouter()

	server := &http.Server{
		Addr:         portString,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	v1Router := router.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/health", handlerRediness).Methods("GET")
	v1Router.HandleFunc("/error", handlerErr).Methods("GET") 

    // TODO: Chnage all PUT request to PATCH requests

    // /cakes
	v1Router.HandleFunc("/cakes", apiCfg.handlerCreateCake).Methods("POST")
    v1Router.HandleFunc("/cakes", apiCfg.handlerGetCakes).Methods("GET")
    v1Router.HandleFunc("/cakes/{id}", apiCfg.handlerGetCake).Methods("GET")
    v1Router.HandleFunc("/cakes/{id}", apiCfg.handlerUpdateCake).Methods("PUT")
    v1Router.HandleFunc("/cakes/{id}", apiCfg.handlerDeleteCake).Methods("DELETE")

    // /flavors
    v1Router.HandleFunc("/flavors", apiCfg.handlerCreateFlavor).Methods("POST")
    v1Router.HandleFunc("/flavors", apiCfg.handlerGetFlavors).Methods("GET")
    v1Router.HandleFunc("/flavors/{id}", apiCfg.handlerGetFlavor).Methods("GET")
    v1Router.HandleFunc("/flavors/{id}", apiCfg.handlerUpdateFlavor).Methods("PUT")
    v1Router.HandleFunc("/flavors/{id}", apiCfg.handlerDeleteFlavor).Methods("DELETE")

    // /frosting 
    v1Router.HandleFunc("/frosting", apiCfg.handlerCreateFrosting).Methods("POST") 
    v1Router.HandleFunc("/frosting", apiCfg.handlerGetFrostings).Methods("GET")
    v1Router.HandleFunc("/frosting/{id}", apiCfg.handlerGetFrosting).Methods("GET") 
    v1Router.HandleFunc("/frosting/{id}", apiCfg.handlerUpdateFrosting).Methods("PUT")
    v1Router.HandleFunc("/frosting/{id}", apiCfg.handlerDeleteFrosting).Methods("DELETE")

    // /filling    
    v1Router.HandleFunc("/filling", apiCfg.handlerCreateFilling).Methods("POST")   
    v1Router.HandleFunc("/filling", apiCfg.handlerGetFillings).Methods("GET")   
    v1Router.HandleFunc("/filling/{id}", apiCfg.handlerGetFilling).Methods("GET")  
    v1Router.HandleFunc("/filling/{id}", apiCfg.handlerUpdateFilling).Methods("PUT") 
    v1Router.HandleFunc("/filling/{id}", apiCfg.handlerDeleteFilling).Methods("DELETE")

	log.Printf("Server running on port %s", portString)
	log.Fatal(server.ListenAndServe())
}
