package main

import (
	"os"
	"todo-list/database"
	"todo-list/routers"
)

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = ":3030"
	}

	database.StartDB()
	routers.StartApp().Run(apiURL)
}
