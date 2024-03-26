package handler

import (
	"fmt"
	"time"
)

type CreateUserRequest struct {
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	CriptPassword string     `json:"criptpassword"`
	Birthday      *time.Time `json:"birthday"`
	Number        int64      `json:"number"`
	Address       string     `json:"address"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" && r.Birthday == nil && r.Number <= 0 && r.Address == "" {
		return fmt.Errorf("corpo do request inválido")
	}
	if r.Name == "" {
		return ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return ErrParamIsRequired("password", "string")
	}
	if r.Birthday == nil {
		return ErrParamIsRequired("birthday", "time*Time")
	}
	if r.Number <= 0 {
		return ErrParamIsRequired("number", "int64")
	}
	if r.Address == "" {
		return ErrParamIsRequired("address", "string")
	}

	return nil
}
