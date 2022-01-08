package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
)

func (b *bus) Login(ctx context.Context, phone, password string) (*JwtDTO, error) {
	u := b.createUrl("/login", nil)
	l := LoginDTO{
		Username: phone,
		Password: password,
	}
	jwt := &JwtDTO{}
	err := requests.PostRequest(ctx, u, l, jwt)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}
