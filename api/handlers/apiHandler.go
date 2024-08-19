package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiHandler struct {
	db gorm.DB
}

type Account struct {
	ID       uuid.UUID `gorm:"primarykey" json:"id"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
}

func New() *ApiHandler {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.AutoMigrate(&Account{})
	return &ApiHandler{
		db: *db,
	}
}

func (s *ApiHandler) WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiHandler) NewAccount(username string, password string) *Account {
	id := uuid.New()
	return &Account{
		ID:       id,
		UserName: username,
		Password: password,
	}
}

func (s *ApiHandler) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	if r.Method == "GET" {
		return s.handleGetALLAccount(ctx, w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(ctx, w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
