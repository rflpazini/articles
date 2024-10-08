package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Em um cenário real, nunca armazenar senhas em texto :D
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
