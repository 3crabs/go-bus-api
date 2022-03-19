package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"time"
)

type OrderDTO struct {
	Agent struct {
		Extra    int    `json:"extra"`
		FullName string `json:"fullName"`
		Id       int    `json:"id"`
		Markup   int    `json:"markup"`
	} `json:"agent"`
	Created       time.Time `json:"created"`
	Expired       time.Time `json:"expired"`
	Finished      time.Time `json:"finished"`
	Id            int       `json:"id"`
	PaymentMethod string    `json:"paymentMethod"`
	Repayment     int       `json:"repayment"`
	ReserveCode   string    `json:"reserveCode"`
	Status        string    `json:"status"`
	Tickets       []struct {
		ArrivalAddress    string    `json:"arrivalAddress"`
		ArrivalDate       time.Time `json:"arrivalDate"`
		ArrivalStation    string    `json:"arrivalStation"`
		Barcode           string    `json:"barcode"`
		Birthday          time.Time `json:"birthday"`
		BusInfo           string    `json:"busInfo"`
		Carrier           string    `json:"carrier"`
		CarrierInn        string    `json:"carrierInn"`
		Citizenship       string    `json:"citizenship"`
		DispatchAddress   string    `json:"dispatchAddress"`
		DispatchDate      time.Time `json:"dispatchDate"`
		DispatchStation   string    `json:"dispatchStation"`
		DocNum            string    `json:"docNum"`
		DocSeries         string    `json:"docSeries"`
		DocType           string    `json:"docType"`
		DownloadUrl       string    `json:"downloadUrl"`
		Dues              int       `json:"dues"`
		FirstName         string    `json:"firstName"`
		Gender            string    `json:"gender"`
		Id                int       `json:"id"`
		LastName          string    `json:"lastName"`
		MiddleName        string    `json:"middleName"`
		Phone             string    `json:"phone"`
		Platform          string    `json:"platform"`
		Price             int       `json:"price"`
		RaceClassId       int       `json:"raceClassId"`
		RaceName          string    `json:"raceName"`
		RaceNum           string    `json:"raceNum"`
		RaceUid           string    `json:"raceUid"`
		Repayment         int       `json:"repayment"`
		Returned          time.Time `json:"returned"`
		Seat              string    `json:"seat"`
		Status            string    `json:"status"`
		SupplierDues      int       `json:"supplierDues"`
		SupplierFare      int       `json:"supplierFare"`
		SupplierPrice     int       `json:"supplierPrice"`
		SupplierRepayment int       `json:"supplierRepayment"`
		TicketClass       string    `json:"ticketClass"`
		TicketCode        string    `json:"ticketCode"`
		TicketNum         string    `json:"ticketNum"`
		TicketSeries      string    `json:"ticketSeries"`
		TicketType        string    `json:"ticketType"`
		Updatable         bool      `json:"updatable"`
		Vat               int       `json:"vat"`
	} `json:"tickets"`
	Total int `json:"total"`
}

func (b *bus) GetOrders(ctx context.Context, accessToken string) (*[]OrderDTO, error) {
	u := b.createUrl("/v1/orders", nil)
	orders := &[]OrderDTO{}
	if err := requests.GetRequest(ctx, u, orders, auth(accessToken)); err != nil {
		return nil, err
	}
	return orders, nil
}

func (b *bus) CancelOrder(ctx context.Context, id int, accessToken string) (*OrderDTO, error) {
	u := b.createUrl(fmt.Sprintf("/v1/orders/%d", id), nil)
	order := &OrderDTO{}
	if err := requests.DeleteRequest(ctx, u, nil, order, auth(accessToken)); err != nil {
		return nil, err
	}
	return order, nil
}
