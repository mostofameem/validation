package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required,min=5,max=20,alpha"`
	Email string `json:"email" validate:"required,email"`
}
