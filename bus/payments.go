package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/http"
)

type PaymentTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (b *bus) GetPayments(ctx context.Context, accessToken string) (*[]PaymentTypeDTO, error) {
	u := b.createUrl("/v1/payments", nil)
	paymentTypeDTOs := &[]PaymentTypeDTO{}
	if err := requests.GetRequest(
		ctx,
		u,
		paymentTypeDTOs,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	); err != nil {
		return nil, err
	}
	return paymentTypeDTOs, nil
}

// TODO: получение url для оплаты

// TODO: возврат билетов
