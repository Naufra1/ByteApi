package handler

import "fmt"

func ErrParamIsRequired(p, t string) error {
	return fmt.Errorf("param: %s (type: %s) é requerido", p, t)
}

type CreateComputerRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Case        string `json:"case"`
	Cpu         string `json:"cpu"`
	Cooler      string `json:"cooler"`
	Graphics    string `json:"gpu"`
	Motherboard string `json:"motherboard"`
	Ram         string `json:"ram"`
	Storage     string `json:"storage"`
}

func (r *CreateComputerRequest) Validate() error {
	if r.Name == "" && r.Image == "" && r.Price == "" && r.Case == "" && r.Cpu == "" && r.Cooler == "" && r.Graphics == "" && r.Motherboard == "" && r.Ram == "" && r.Storage == "" {
		return fmt.Errorf("corpo do request inválido ou vazio")
	}
	if r.Name == "" {
		return ErrParamIsRequired("name", "string")
	}
	if r.Image == "" {
		return ErrParamIsRequired("image", "string")
	}
	if r.Price == "" {
		return ErrParamIsRequired("price", "string")
	}
	if r.Case == "" {
		return ErrParamIsRequired("case", "string")
	}
	if r.Cpu == "" {
		return ErrParamIsRequired("cpu", "string")
	}
	if r.Cooler == "" {
		return ErrParamIsRequired("cooler", "string")
	}
	if r.Graphics == "" {
		return ErrParamIsRequired("graphics", "string")
	}
	if r.Motherboard == "" {
		return ErrParamIsRequired("motherboard", "string")
	}
	if r.Ram == "" {
		return ErrParamIsRequired("ram", "string")
	}
	if r.Storage == "" {
		return ErrParamIsRequired("storage", "string")
	}

	return nil
}
