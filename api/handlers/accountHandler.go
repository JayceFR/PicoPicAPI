package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *ApiHandler) handleGetALLAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resoponse, err := s.GetAllAcconts(ctx, s.db)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s.WriteJson(w, http.StatusOK, resoponse)
}

type CreateAccount struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

const null_uuid = "00000000-0000-0000-0000-000000000000"

func (s *ApiHandler) handleCreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	createAccount := CreateAccount{}
	err := json.NewDecoder(r.Body).Decode(&createAccount)
	if err != nil {
		fmt.Println(err.Error())
	}
	var check_account Account
	s.db.First(&check_account, "username = ?", createAccount.UserName)
	if check_account.ID.String() == null_uuid {
		account := s.NewAccount(createAccount.UserName, createAccount.Password)
		s.db.Create(account)
		return s.WriteJson(w, http.StatusOK, createAccount)
	}
	return s.WriteJson(w, http.StatusBadRequest, "Found another user with the same username")
}
