package controllers

import (
	"context"
	"encoding/json"
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
	payload := signIn(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func signIn(name string) models.User {
	result := models.User{}
	filter := bson.D{primitive.E{Key: "Name", Value: name}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
