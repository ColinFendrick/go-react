package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SignIn confirms that the user exists and returns the user
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	payload := signIn(params["name"])
	json.NewEncoder(w).Encode(payload)
}

func signIn(name string) models.User {
	var result models.User
	filter := bson.D{primitive.E{Key: "Name", Value: name}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// CreateUser makes a user in the db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&user)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	_, existingErr := checkForDuplicate(user.Name)
	if existingErr != nil {
		log.Fatal(existingErr)
	}

	addUser(user)
	json.NewEncoder(w).Encode(user)
}

func checkForDuplicate(field interface{}) (models.User, error) {
	var result models.User
	filter := bson.D{primitive.E{Key: "Name", Value: field}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func addUser(user models.User) {
	insertResult, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
