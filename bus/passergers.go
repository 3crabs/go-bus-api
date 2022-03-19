package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
	"time"
)

type PassengerDTO struct {
	Birthday     time.Time `json:"birthday"`
	Citizenship  string    `json:"citizenship"`
	ConfirmEmail bool      `json:"confirmEmail"`
	DocNum       string    `json:"docNum"`
	DocSeries    string    `json:"docSeries"`
	DocTypeCode  string    `json:"docTypeCode"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	Gender       string    `json:"gender"`
	Id           int       `json:"id"`
	LastName     string    `json:"lastName"`
	MiddleName   string    `json:"middleName"`
	Owner        bool      `json:"owner"`
	Phone        string    `json:"phone"`
}

type PassengerCreateDTO struct {
	Birthday    time.Time `json:"birthday"`
	Citizenship string    `json:"citizenship"`
	DocNum      string    `json:"docNum"`
	DocSeries   string    `json:"docSeries"`
	DocTypeCode string    `json:"docTypeCode"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	Gender      string    `json:"gender"`
	LastName    string    `json:"lastName"`
	MiddleName  string    `json:"middleName"`
	Owner       bool      `json:"owner"`
	Phone       string    `json:"phone"`
}

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

// TODO: редактирование пассажира

// TODO: подтверждение почты пассажира
