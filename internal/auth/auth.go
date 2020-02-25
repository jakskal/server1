package auth

// Auth represent params to authenticate user.
type Auth struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
