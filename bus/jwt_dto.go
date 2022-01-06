package bus

type JwtDTO struct {
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    int      `json:"expires_in"`
}
