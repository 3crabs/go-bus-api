package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

func (b *bus) GetAccount(ctx context.Context, accessToken string) (*UserDTO, error) {
	u := b.createUrl("/v1/account", nil)
	user := &UserDTO{}
	err := requests.GetRequest(
		ctx,
		u,
		user,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
