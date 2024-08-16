package main

import (
    "big-web-app/config"
    "big-web-app/routes"
)

func main() {
    config.ConnectDatabase()
    defer config.CloseDatabase() // Ensure the database is closed when the application exits
    
    r := routes.SetupRouter()
    r.Run(":8080")
}
