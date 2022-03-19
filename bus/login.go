package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
)

type JwtDTO struct {
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    int      `json:"expires_in"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *bus) Login(ctx context.Context, phone, password string) (*JwtDTO, error) {
	u := b.createUrl("/login", nil)
	l := LoginDTO{
		Username: phone,
		Password: password,
	}
	jwt := &JwtDTO{}
	if err := requests.PostRequest(ctx, u, l, jwt); err != nil {
		return nil, err
	}
	return jwt, nil
}
