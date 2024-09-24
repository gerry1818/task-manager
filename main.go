package main

import (
	"log"
	"net/http"
	"task-manager/controllers"
	"task-manager/middlewares"
	"task-manager/models"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the DB
	models.ConnectDatabase()

	// Initialize Router
	r := mux.NewRouter()

	// Middlewares
	r.Use(middlewares.Logger)

	// Auth routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Task routes (protected)
	s := r.PathPrefix("/tasks").Subrouter()
	s.Use(middlewares.Auth)
	s.HandleFunc("", controllers.GetTasks).Methods("GET")
	s.HandleFunc("/{id}", controllers.GetTask).Methods("GET")
	s.HandleFunc("", controllers.CreateTask).Methods("POST")
	s.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
	s.HandleFunc("/{id}", controllers.DeleteTask).Methods("DELETE")
	s.HandleFunc("/done", controllers.MarkTasksDone).Methods("POST") // Mark tasks as done concurrently

	log.Fatal(http.ListenAndServe(":8000", r))
}
