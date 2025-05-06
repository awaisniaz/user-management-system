package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/awaisniaz/user-management/internal/models"
	"github.com/awaisniaz/user-management/internal/repository"
	"github.com/awaisniaz/user-management/internal/utils"
	"github.com/google/uuid"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Failed to hash Password", http.StatusInternalServerError)
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		http.Error(w, "Failed to hash Password", http.StatusInternalServerError)
	}
	user := models.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  hashPassword,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	err = repository.CreateUser(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
