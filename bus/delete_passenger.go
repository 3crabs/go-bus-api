package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

func (b *bus) DeletePassenger(ctx context.Context, accessToken string, passengerID int) error {
	u := b.createUrl(fmt.Sprintf("/v1/passengers/%d", passengerID), nil)
	return requests.DeleteRequest(
		ctx,
		u,
		nil,
		nil,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
}
