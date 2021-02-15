package responses

// LoginResponse model
type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `token:"token"`
}
