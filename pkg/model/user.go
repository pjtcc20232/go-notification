package model

import (
	"encoding/json"
	"time"

	"github.com/notification/back-end/internal/config/logger"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint64 `json:"_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password,omitempty"`
	HashedPassword string `json:"-"`
	Enable         bool   `json:"enable"`
	IsLocked       bool   `json:"isLocked"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

func (u *User) passwordToHash() {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			logger.Error("Erro to SetPassWord:"+err.Error(), err)

		}

		u.HashedPassword = string(hashedPassword)
	}
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	if err != nil {
		logger.Error("Erro to CheckPassword:"+err.Error(), err)
		return false
	}
	return true
}

func (u *User) PrepareToSave() {
	dt := time.Now().Format(time.RFC3339)
	u.passwordToHash()

	if u.ID == 0 {
		u.CreatedAt = dt
		u.UpdatedAt = dt
	} else {
		u.UpdatedAt = dt
	}
}

type UsertList struct {
	List []*User `json:"list"`
}

func (Ul UsertList) String() string {
	data, err := json.Marshal(Ul)

	if err != nil {
		logger.Error("error to convert ProductList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
