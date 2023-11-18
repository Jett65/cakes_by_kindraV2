//TODO: Create logs for each request

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
    "log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env") 

    porlString := fmt.Sprintf(":%s", os.Getenv("PORT"))

    router := mux.NewRouter()

    server := &http.Server{
        Addr: porlString, 
        WriteTimeout: time.Second * 15, 
        ReadTimeout: time.Second * 15,
        IdleTimeout: time.Second * 60,
        Handler: router,
    }
    
    v1Router := router.PathPrefix("/v1").Subrouter()
    v1Router.HandleFunc("/health", handlerRediness)
    
    log.Printf("Server running on port %s", porlString)
    log.Fatal(server.ListenAndServe())
}   
