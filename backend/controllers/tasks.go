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

// GetAllTasks Gets all tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	var user models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&user)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	_ = json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(user, task)
	json.NewEncoder(w).Encode(task)
}

func insertOneTask(user models.User, task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func confirmTaskBelongsToUser(user models.User, taskID string) {
	var belongs bool

	id, _ := primitive.ObjectIDFromHex(taskID)
	for _, a := range user.ToDoList.ToDos {
		if a.ID == id {
			belongs = true
			break
		}
		belongs = false
	}

	if !belongs {
		log.Fatal("This task does not belong to you")
	}
}

// DeleteTask delets a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&user)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	params := mux.Vars(r)

	confirmTaskBelongsToUser(user, params["id"])

	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	return d.DeletedCount
}
