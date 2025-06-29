package auth

type LoginRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type RegisterRequest struct {
	Username string
	FirstName string
	LastName string
	DateOfBirth string
	PhotoURL string
	Password string
}

type RefreshRequest struct {
	RefreshToken string
}
