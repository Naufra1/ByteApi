package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Name        string
	Image       string
	Price       string
	Case        string
	Cpu         string
	Cooler      string
	Graphics    string
	Motherboard string
	Ram         string
	Storage     string
}

type ComputerSpecs struct {
	Case        string `json:"case"`
	Cpu         string `json:"cpu"`
	Cooler      string `json:"cooler"`
	Graphics    string `json:"gpu"`
	Motherboard string `json:"motherboard"`
	Ram         string `json:"ram"`
	Storage     string `json:"storage"`
}

type ComputerResponse struct {
	ID        uint          `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt time.Time     `json:"deletedAt,omitempty"`
	Name      string        `json:"name"`
	Image     string        `json:"image"`
	Price     string        `json:"price"`
	Specs     ComputerSpecs `json:"Specs"`
}
