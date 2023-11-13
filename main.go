package main

import (
	"github.com/subosito/gotenv"
	"golang-assignment/database"
	"golang-assignment/migration"
	"golang-assignment/provider"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
	database.ConnectToDatabase()
	migration.MigrateAll()
	// Initialize routes
	provider.Run()
}
