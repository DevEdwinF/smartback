package entity

type User struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
}
