package bus

import (
	"context"
	"github.com/3crabs/go-bus-api/rest"
	"net/http"
)

func (b *bus) GetPassengers(ctx context.Context, accessToken string) (*[]PassengerDTO, error) {
	u := b.createUrl("/v1/passengers", nil)
	passengers := &[]PassengerDTO{}
	err := rest.GetRequest(
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
