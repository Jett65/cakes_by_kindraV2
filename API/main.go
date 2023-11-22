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

	porlString := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if porlString == "" {
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
		Addr:         porlString,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	v1Router := router.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/health", handlerRediness).Methods("GET")
	v1Router.HandleFunc("/error", handlerErr).Methods("GET")

	v1Router.HandleFunc("/cakes", apiCfg.handlerCreateCake).Methods("POST")
    v1Router.HandleFunc("/cakes", apiCfg.handlerGetCakes).Methods("GET")
    // Update cake
    // Delete cake

	log.Printf("Server running on port %s", porlString)
	log.Fatal(server.ListenAndServe())
}
