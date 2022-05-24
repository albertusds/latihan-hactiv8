package params

type CreatePerson struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdatePerson struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
