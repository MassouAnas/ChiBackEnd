package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MassouAnas/ChiBackEnd/model"
	"github.com/MassouAnas/ChiBackEnd/repository/todo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct{
	router http.Handler
	client *mongo.Client
}

func New(mongodbURI string) (*App, error){

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client , err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w",err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	app := &App{
		router: listRoutes(),
		client: client,
	}
	repo := todo.NewMongoRepo(client)

	now := time.Now()
	t := model.Todo{
	TodoID: 1,   
	TodoTitle: "First todo", 
	TodoDescr: "This just might be our first todo",
	CreatedAt:&now,
	FinishedAt: nil,  
	}
	
	err = repo.Insert(context.Background(), t)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert todo %w", err)
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3333",
		Handler: a.router,
	}
	fmt.Printf("Listening for http requests on localhost%s \n", server.Addr)

	// Start the HTTP server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to listen to server: %s\n", err)
		}
	}()

	// Wait for termination signal
	<-ctx.Done()

	// Shutdown the server
	fmt.Println("Shutting down server...")
	err := server.Shutdown(context.Background())
	if err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
	}

	// Close MongoDB connection
	if err := a.client.Disconnect(ctx); err != nil {
		fmt.Printf("Error disconnecting from MongoDB: %s\n", err)
	}

	return nil
}

// func (a *App) Start(ctx context.Context) error{
// 	server := &http.Server{
// 		Addr: ":3333",
// 		Handler: a.router,
// 	}
// 	fmt.Printf("Listening for http requests on localhost%s \n", server.Addr)
	
// 	ch := make(chan error, 1)

// 	go func(){
// 		err := server.ListenAndServe()
// 	if err != nil {
// 		ch <- fmt.Errorf("Failed to listen to server: %w", err)
// 	}
// 	close(ch)
// 	}()

// 	go func() {
// 		<-ctx.Done()
// 		fmt.Println("Shutting down server...")
// 		err := server.Shutdown(context.Background())
// 		if err != nil {
// 			fmt.Printf("Error shutting down server: %s\n", err)
// 		}

// 		// Close MongoDB connection
// 		if err := a.client.Disconnect(ctx); err != nil {
// 			fmt.Printf("Error disconnecting from MongoDB: %s\n", err)
// 		}
// 	}()
	
// 	return nil
// }
