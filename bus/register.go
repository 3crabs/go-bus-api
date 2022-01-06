package bus

import (
	"context"
	"github.com/3crabs/go-bus-api/rest"
)

func (b *bus) Register(ctx context.Context, phone string) error {
	u := b.createUrl("/v1/users", nil)
	p := PhoneDTO{Phone: phone}
	return rest.PostRequest(ctx, u, p, nil)
}
