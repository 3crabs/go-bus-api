package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

func (b *bus) AddPassenger(ctx context.Context, accessToken string, passenger PassengerCreateDTO) (*PassengerDTO, error) {
	u := b.createUrl("/v1/passengers", nil)
	newPassenger := &PassengerDTO{}
	err := requests.PostRequest(
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
