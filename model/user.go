package model

// User represents a user in the application
type User struct {
	ID       interface{} `bson:"_id,omitempty"` // Use interface{} to handle different types of ID
	Name     string      `bson:"name"`
	Email    string      `bson:"email"`
	Password string      `bson:"password"`
}
