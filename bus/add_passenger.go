package bus

import (
	"context"
	"github.com/3crabs/go-bus-api/rest"
	"net/http"
)

func (b *bus) AddPassenger(ctx context.Context, accessToken string, passenger PassengerCreateDTO) (*PassengerDTO, error) {
	u := b.createUrl("/v1/passengers", nil)
	newPassenger := &PassengerDTO{}
	err := rest.PostRequest(
		ctx,
		u,
		passenger,
		newPassenger,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
	if err != nil {
		return nil, err
	}
	return newPassenger, nil
}
