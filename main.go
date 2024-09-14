package main

import (
	"FORUM/config"
	"FORUM/routes"
	"FORUM/utilis"
	"fmt"
	"log"
	"net/http"
)


func main() {
	// http.HandleFunc("/",)

	var err error
	
	utilis.DB, err = utilis.OpenDatabase(config.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	
	defer utilis.DB.Close()
	utilis.CreateTables()
	
	routes.SetupRoutes()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
