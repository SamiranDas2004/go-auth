// controller/controller.go

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SamiranDas2004/go-auth/dbconnect"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register is a handler for user registration
func Register(w http.ResponseWriter, r *http.Request) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request fields
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	// Hash password with bcrypt
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update request struct with hashed password
	req.Password = string(HashedPassword)

	// Get the MongoDB collection
	collection := dbconnect.ConnectMongoDB()

	var existingUser struct {
		Email string `json:"email"`
	}
	filter := bson.M{"email": req.Email}

	existUser := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if existUser != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	// else if existUser != mongo.ErrNoDocuments {
	// 	http.Error(w, existUser.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Insert the new user into the collection
	_, err = collection.InsertOne(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and Password is Required", http.StatusBadRequest)
		return
	}

	collection := dbconnect.ConnectMongoDB()

	filter := bson.M{"email": req.Email}
	var user bson.M
	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		// Handle error or return appropriate response
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Extract hashed password from the user object
	hashedPassword, ok := user["password"].(string)
	fmt.Println(hashedPassword)
	if !ok {
		http.Error(w, "Invalid password format in the database", http.StatusInternalServerError)
		return
	}

	// Compare hashed passwords
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Passwords match, user is authenticated
	fmt.Println("User authenticated successfully")

	// Here you could compare passwords, but for simplicity, let's assume the user exists

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		// You can add more claims here such as user ID, roles, etc.
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with a secret
	tokenString, err := token.SignedString([]byte("juigfuweq89qwcur8cqiwutwhtuwheuthwurhthwthw45tuhw utvwue"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with token
	json.NewEncoder(w).Encode(map[string]string{"user logedin ": tokenString})
}
