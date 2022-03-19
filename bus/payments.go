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

type PaymentsUrlDTO struct {
	URL string `json:"url"`
}

type PaymentsUrlRequest struct {
	OrderIds      []int `json:"orderIds"`
	PaymentTypeId int   `json:"paymentTypeId"`
}

type RefundTicketsRequest struct {
	TicketIds []int `json:"ticketIds"`
}

type RefundTicketsDTO struct {
	Data string `json:"data"`
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

func (b *bus) GetPaymentsURL(ctx context.Context, paymentsUrlRequest PaymentsUrlRequest, accessToken string) (*PaymentsUrlDTO, error) {
	u := b.createUrl("/v1/payments", nil)
	paymentsUrl := &PaymentsUrlDTO{}
	if err := requests.PostRequest(
		ctx,
		u,
		paymentsUrlRequest,
		paymentsUrl,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	); err != nil {
		return nil, err
	}
	return paymentsUrl, nil
}

func (b *bus) GetPaymentsRefundTickets(ctx context.Context, refundTicketsRequest RefundTicketsRequest, accessToken string) (*RefundTicketsDTO, error) {
	u := b.createUrl("/v1/payments/refund", nil)
	refundTicketsDTO := &RefundTicketsDTO{}
	if err := requests.PostRequest(
		ctx,
		u,
		refundTicketsRequest,
		refundTicketsDTO,
		func(req *http.Request) {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		},
	); err != nil {
		return nil, err
	}
	return refundTicketsDTO, nil
}
