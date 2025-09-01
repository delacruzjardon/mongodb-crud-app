package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"mongodb-crud-app/handlers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace with your MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:passwordone@cluster0.r1nn65n.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Test connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// Get collection
	collection := client.Database("crudapp").Collection("users")

	// Initialize handler
	userHandler := handlers.NewUserHandler(collection)

	// Setup routes
	router := mux.NewRouter()

	// Web routes
	router.HandleFunc("/", userHandler.IndexHandler).Methods("GET")
	router.HandleFunc("/create", userHandler.CreateFormHandler).Methods("GET")
	router.HandleFunc("/create", userHandler.CreateHandler).Methods("POST")
	router.HandleFunc("/edit/{id}", userHandler.EditFormHandler).Methods("GET")
	router.HandleFunc("/update/{id}", userHandler.UpdateHandler).Methods("POST")
	router.HandleFunc("/delete/{id}", userHandler.DeleteHandler).Methods("GET")
	// Add these routes to your existing router setup in main.go

	// Web routes (add after existing routes)
	router.HandleFunc("/search", userHandler.SearchHandler).Methods("GET")

	// API routes (add after existing API routes)
	router.HandleFunc("/api/search", userHandler.SearchUsersAPI).Methods("GET")

	// API routes
	router.HandleFunc("/api/users", userHandler.GetUsersAPI).Methods("GET")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
