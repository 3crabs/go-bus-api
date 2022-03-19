package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
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

type ConfirmPassengerEmailResponse struct {
	Data string `json:"data"`
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
	if err := requests.PostRequest(ctx, u, passenger, newPassenger, auth(accessToken)); err != nil {
		return nil, err
	}
	return newPassenger, nil
}

func (b *bus) GetPassengers(ctx context.Context, accessToken string) (*[]PassengerDTO, error) {
	u := b.createUrl("/v1/passengers", nil)
	passengers := &[]PassengerDTO{}
	if err := requests.GetRequest(ctx, u, passengers, auth(accessToken)); err != nil {
		return nil, err
	}
	return passengers, nil
}

func (b *bus) PutPassengers(ctx context.Context, id int, newPassenger PassengerDTO, accessToken string) (*PassengerDTO, error) {
	u := b.createUrl(fmt.Sprintf("/v1/passengers/%d", id), nil)
	passenger := &PassengerDTO{}
	if err := requests.PutRequest(ctx, u, newPassenger, passenger, auth(accessToken)); err != nil {
		return nil, err
	}
	return passenger, nil
}

func (b *bus) DeletePassenger(ctx context.Context, accessToken string, passengerID int) error {
	u := b.createUrl(fmt.Sprintf("/v1/passengers/%d", passengerID), nil)
	return requests.DeleteRequest(ctx, u, nil, nil, auth(accessToken))
}

func (b *bus) ConfirmPassengerEmail(ctx context.Context, id int, accessToken string) (*ConfirmPassengerEmailResponse, error) {
	u := b.createUrl(fmt.Sprintf("/v1/passengers/%d/confirmEmailRequest", id), nil)
	confirmPassengerEmailResponse := &ConfirmPassengerEmailResponse{}
	if err := requests.GetRequest(ctx, u, confirmPassengerEmailResponse, auth(accessToken)); err != nil {
		return nil, err
	}
	return confirmPassengerEmailResponse, nil
}
