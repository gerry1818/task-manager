package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"task-manager/models"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	var tasks []models.Task
	models.DB.Where("user_id = ?", userID).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task models.Task
	if err := models.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	task.UserID = userID
	if err := models.DB.Create(&task).Error; err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task models.Task
	if err := models.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	models.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := models.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func MarkTasksDone(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	var taskIDs []uint
	json.NewDecoder(r.Body).Decode(&taskIDs)

	var wg sync.WaitGroup
	for _, id := range taskIDs {
		wg.Add(1)
		go func(taskID uint) {
			defer wg.Done()
			models.DB.Model(&models.Task{}).Where("id = ? AND user_id = ?", taskID, userID).Update("status", "done")
		}(id)
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
}
