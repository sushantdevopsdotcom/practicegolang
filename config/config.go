package config

import (
    "fmt"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // Import the SQLite dialect
    "big-web-app/models" // Ensure this import matches your project structure
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and performs migrations
func ConnectDatabase() {
    var err error
    DB, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    
    // AutoMigrate will create the table if it doesn't exist, or update it if needed
    if err := DB.AutoMigrate(&models.Item{}).Error; err != nil {
        log.Fatalf("Error migrating the database: %v", err)
    }

    fmt.Println("Database connection successful")
}

// CloseDatabase closes the database connection
func CloseDatabase() {
    if err := DB.Close(); err != nil {
        log.Fatalf("Error closing the database: %v", err)
    }
}
