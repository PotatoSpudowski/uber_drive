package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"./app"
	// "./controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/drivers/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/drivers/login", controllers.Authenticate).Methods("POST")
	// router.HandleFunc("/api/rides/new", controllers.CreateRide).Methods("POST")
	// router.HandleFunc("/api/me/rides", controllers.GetRidesFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port == "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}