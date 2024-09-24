package controllers

import (
	"encoding/json"
	"net/http"
	"task-manager/models"
	"task-manager/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.HashPassword(user.Password); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	if err := models.DB.Create(&user).Error; err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var dbUser models.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := models.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := dbUser.CheckPassword(user.Password); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
