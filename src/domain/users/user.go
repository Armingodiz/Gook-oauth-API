package users

type User struct {
	Id          int64  `json:"id"`
	FirsName    string `json:"first_name"`
	LastNAme    string `json:"last_name"`
	Email       string `json:"email"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
