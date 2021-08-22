package models

type RqRegisterUser struct {
	Names    string `json:"names" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RolCode  string `json:"rol_code" validate:"required"`
}
