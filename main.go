package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dadebulba/marisabooking/routes"
	"github.com/dadebulba/marisabooking/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	dsn := os.Getenv("DATABASE_CONNECTION")
	utils.InitDB(dsn)
	defer utils.CloseDB()

	// Initialize Gin router
	router := routes.SetupRouter()

	// Graceful shutdown handling
	go func() {
		log.Println("Server is running on port 8080")
		if err := router.Run("127.0.0.1:8080"); err != nil {
			log.Fatalf("Server stopped with error: %v", err)
		}
	}()

	// Handle shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
