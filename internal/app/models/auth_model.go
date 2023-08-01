package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Password string `json:"password"`
	FkRoleId string `json:"role"`
	RoleName string `json:"role_name"`
}
