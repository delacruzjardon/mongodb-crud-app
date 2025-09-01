package handlers

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"mongodb-crud-app/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	Collection *mongo.Collection
}

func NewUserHandler(collection *mongo.Collection) *UserHandler {
	return &UserHandler{
		Collection: collection,
	}
}

// Home page - displays all users
func (h *UserHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.Collection.Find(ctx, bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	tmpl.Execute(w, users)
}

// Create user form
func (h *UserHandler) CreateFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/create.html"))
	tmpl.Execute(w, nil)
}

// Create user
func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	ageStr := r.FormValue("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = h.Collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Edit user form
func (h *UserHandler) EditFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = h.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tmpl.Execute(w, user)
}

// Update user
func (h *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	ageStr := r.FormValue("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":  name,
			"email": email,
			"age":   age,
		},
	}

	_, err = h.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete user
func (h *UserHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = h.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// API endpoints for JSON responses
func (h *UserHandler) GetUsersAPI(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.Collection.Find(ctx, bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Add this to your existing UserHandler struct methods

// SearchHandler handles user search requests
func (h *UserHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		// If no search query, redirect to home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a case-insensitive regex search across name and email fields
	searchRegex := primitive.Regex{
		Pattern: query,
		Options: "i", // Case-insensitive
	}

	filter := bson.M{
		"$or": []bson.M{
			{"name": searchRegex},
			{"email": searchRegex},
		},
	}

	cursor, err := h.Collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error searching users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		log.Printf("Error decoding search results: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create search results data
	searchData := struct {
		Users []models.User
		Query string
		Count int
	}{
		Users: users,
		Query: query,
		Count: len(users),
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/search.html"))
	tmpl.Execute(w, searchData)
}

// API endpoint for search
func (h *UserHandler) SearchUsersAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.User{})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	searchRegex := primitive.Regex{
		Pattern: query,
		Options: "i",
	}

	filter := bson.M{
		"$or": []bson.M{
			{"name": searchRegex},
			{"email": searchRegex},
		},
	}

	cursor, err := h.Collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
