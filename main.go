package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/MassouAnas/ChiBackEnd/application"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	mongodbURI := os.Getenv("MONGODB_URI")
	// Initialize application
	app, err := application.New(mongodbURI)
	if err != nil {
		fmt.Println("failed to initialize app:", err)
		return
	}

	// Create a context that cancels with SIGINT (Ctrl+C)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Start the application with the created context
	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}



