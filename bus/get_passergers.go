package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

func (b *bus) GetPassengers(ctx context.Context, accessToken string) (*[]PassengerDTO, error) {
	u := b.createUrl("/v1/passengers", nil)
	passengers := &[]PassengerDTO{}
	err := requests.GetRequest(
		ctx,
		u,
		passengers,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
	if err != nil {
		return nil, err
	}
	return passengers, nil
}
