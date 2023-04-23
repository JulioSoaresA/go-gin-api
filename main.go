package main

import (
	"github.com/JulioSoaresA/go-gin-api/database"
	"github.com/JulioSoaresA/go-gin-api/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()
}
