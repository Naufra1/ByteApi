package user

import (
	"fmt"
	"time"

	"github.com/Naufra1/ByteApi/handler/computer"
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
		return computer.ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return computer.ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return computer.ErrParamIsRequired("password", "string")
	}
	if r.Birthday == nil {
		return computer.ErrParamIsRequired("birthday", "time*Time")
	}
	if r.Number <= 0 {
		return computer.ErrParamIsRequired("number", "int64")
	}
	if r.Address == "" {
		return computer.ErrParamIsRequired("address", "string")
	}

	return nil
}
