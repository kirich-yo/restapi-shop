package auth

type LoginRequest struct {
	Username string
	Password string
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
