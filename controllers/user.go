package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunilnarwaria/Golang-And-MongoDB-REST-API/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserController defines the controller for user operations
type UserController struct {
	session *mgo.Session
}

// NewUserController returns a new UserController
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves a user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Verify if the ID is a valid ObjectID
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Create an empty User model to hold the fetched result
	u := models.User{}

	// Query the database
	if err := uc.session.DB("userdb").C("users").FindId(oid).One(&u); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Marshal the result into JSON and send as response
	uJson, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// Decode the incoming JSON payload
	json.NewDecoder(r.Body).Decode(&u)

	// Assign a new ObjectId to the user
	u.Id = bson.NewObjectId()

	// Insert the new user into the database
	uc.session.DB("userdb").C("users").Insert(u)

	// Marshal the result into JSON and send as response
	uJson, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uJson)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Verify if the ID is a valid ObjectID
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Remove the user from the database
	if err := uc.session.DB("userdb").C("users").RemoveId(oid); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
}
