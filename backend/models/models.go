package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDo is the basic item in a ToDoList
type ToDo struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

// ToDoList is simply a collection of Todos
type ToDoList struct {
	ID   primitive.ObjectID
	ToDo []ToDo
}

// User is the current user and their todo list
type User struct {
	ID       primitive.ObjectID
	Name     string
	Email    string
	ToDoList ToDoList
}
