package _struct

type User struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Username  string `json:"username"`
}
