package api

import (
	"context"

	"gorm.io/gorm"
)

func (s *ApiHandler) GetAllAcconts(ctx context.Context, db gorm.DB) ([]*Account, error) {
	rows, err := s.db.WithContext(ctx).
		Select("*").
		Table("accounts").
		Rows()
	if err != nil {
		return []*Account{}, err
	}

	response := []*Account{}
	for rows.Next() {
		account := Account{}
		err := rows.Scan(
			&account.ID,
			&account.UserName,
			&account.Password,
		)
		if err != nil {
			return []*Account{}, err
		}
		response = append(response, &account)
	}
	return response, nil
}
