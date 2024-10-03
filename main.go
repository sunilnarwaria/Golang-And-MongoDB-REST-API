package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunilnarwaria/Golang-And-MongoDB-REST-API/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Create a new router
	r := httprouter.New()

	// Initialize a new UserController and pass the MongoDB session
	uc := controllers.NewUserController(getSession())

	// Define routes
	r.GET("/user/:id", uc.GetUser)       // Get user by ID
	r.POST("/user", uc.CreateUser)       // Create a new user
	r.DELETE("/user/:id", uc.DeleteUser) // Delete a user by ID

	// Start the server on port 9000
	http.ListenAndServe("localhost:9000", r)
}

// getSession creates a new MongoDB session and connects to localhost MongoDB instance
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017") // Replace with your MongoDB URI if needed
	if err != nil {
		panic(err)
	}
	return s
}
