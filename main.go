package main

import (
	"fmt"
	"net/http"
	"os"

	"./app"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//attach JWT auth Middleware
	router.Use(app.JwtAuthentication)
	//get the port from the env file
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	//Launch the app, visit localhost:8000/api
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
