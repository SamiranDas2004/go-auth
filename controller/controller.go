package controller

import (
	"context"
	"log"
	"time"

	"github.com/SamiranDas2004/go-auth/dbconnect"
	"github.com/SamiranDas2004/go-auth/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot hash password"})
	}

	user := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	collection := dbconnect.ConnectMongoDB()

	// Check if a user with the same email already exists
	filter := bson.M{"email": data["email"]}
	var existingUser model.User
	err = collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user already exists with this email"})
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot insert user"})
	}

	log.Println("User registered:", user) // Logging for debugging

	return c.JSON(user)
}
func Login(c *fiber.Ctx) error {
	// Parse request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Check if email and password fields are present
	email, ok := data["email"]
	if !ok || email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email field is missing"})
	}
	password, ok := data["password"]
	if !ok || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password field is missing"})
	}

	// Connect to MongoDB
	collection := dbconnect.ConnectMongoDB()

	// Find user by email
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		// If the user is not found, return the same error message for security reasons
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// If the password does not match, return the same error message for security reasons
		log.Println("Password comparison failed:", err) // Logging for debugging
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiry time
	})

	// Sign the token with a secret key and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("whatislovebabaydonthurtme"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	// Set the JWT token as a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	// If the credentials are correct, return the JWT token
	return c.JSON(fiber.Map{"message": "Logged in", "token": tokenString})
}

func Logout(c *fiber.Ctx) error {
	// Set the JWT token as an expired cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Logged out"})
}
