package domain

// AuthRequest represents the request body for authentication.
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response body for authentication.
type AuthResponse struct {
	Token string `json:"token"`
}
