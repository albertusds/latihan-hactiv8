package models

type Person struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
