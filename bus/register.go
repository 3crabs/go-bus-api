package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
)

func (b *bus) Register(ctx context.Context, phone string) error {
	u := b.createUrl("/v1/users", nil)
	p := PhoneDTO{Phone: phone}
	return requests.PostRequest(ctx, u, p, nil)
}
