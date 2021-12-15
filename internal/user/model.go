package user

// User model for mongodb. bson tags for mongoDB
type User struct {
	ID           string `json:"id" bson:"_id,omitempty"` // omitempty - can be empty
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"` // json: "-" means field will not be in json
}

// CreateUserDTO user DTO
type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
