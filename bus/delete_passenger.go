package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/rest"
	"net/http"
)

func (b *bus) DeletePassenger(ctx context.Context, accessToken string, passengerID int) error {
	u := b.createUrl(fmt.Sprintf("/v1/passengers/%d", passengerID), nil)
	return rest.DeleteRequest(
		ctx,
		u,
		nil,
		nil,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
}
