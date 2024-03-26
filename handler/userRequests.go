package handler

import (
	"fmt"
)

type CreateUserRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Birthday        string `json:"birthday"`
	Number          int64  `json:"number"`
	Address         string `json:"address"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" && r.Birthday == "" && r.Number <= 0 && r.Address == "" && r.ConfirmPassword == "" {
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
	if r.ConfirmPassword == "" {
		return ErrParamIsRequired("confirm password", "string")
	}
	if r.ConfirmPassword != r.Password {
		return fmt.Errorf("digite senhas iguais")
	}
	if r.Birthday == "" {
		return ErrParamIsRequired("birthday", "string")
	}
	if r.Number <= 0 {
		return ErrParamIsRequired("number", "int64")
	}
	if r.Address == "" {
		return ErrParamIsRequired("address", "string")
	}

	return nil
}
