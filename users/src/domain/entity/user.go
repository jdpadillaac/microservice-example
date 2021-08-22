package entity

type User struct {
	ID       string `json:"id"`
	Names    string `json:"names"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RolCode  string `json:"rol_code"`
}
