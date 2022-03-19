package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

type UserDTO struct {
	ConfirmEmail bool   `json:"confirmEmail"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	Id           int    `json:"id"`
	LastName     string `json:"lastName"`
	MiddleName   string `json:"middleName"`
	Phone        string `json:"phone"`
}

type UpdateUserDTO struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}

func (b *bus) GetAccount(ctx context.Context, accessToken string) (*UserDTO, error) {
	u := b.createUrl("/v1/account", nil)
	user := &UserDTO{}
	err := requests.GetRequest(
		ctx,
		u,
		user,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *bus) PutAccount(ctx context.Context, updateUserDTO UpdateUserDTO, accessToken string) (*UserDTO, error) {
	u := b.createUrl("/v1/account", nil)
	user := &UserDTO{}
	err := requests.PutRequest(
		ctx,
		u,
		updateUserDTO,
		user,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
