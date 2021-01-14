package authorization

type AuthorizationResponse struct {
	Token string
}

func NewAuthorizationResponse(token string) *AuthorizationResponse {
	return &AuthorizationResponse{Token: token}
}
