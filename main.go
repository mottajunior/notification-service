package main

import (
	"fmt"
	"github.com/gorilla/mux"
	notificationRouter "github.com/mottajunior/notification-service/router"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Stay Alive")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/races", notificationRouter.SendNotification).Methods("GET")

	var port = ":4000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

