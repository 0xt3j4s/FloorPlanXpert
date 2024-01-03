package main

import (
	"FloorPlanXpert/api"
	"FloorPlanXpert/internal/db"
)

func main() {
    // Establish database connection
    if err := db.Connect(); err != nil {
        panic(err)
    }
    defer db.Close() // Close the database connection when the application exits

    // Create a Gin router
    router := api.SetupRouter()

    // Run the server
    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}
