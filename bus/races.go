package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"net/url"
	"strconv"
	"time"
)

type RaceDTO struct {
	ArrivalDate        time.Time `json:"arrivalDate"`
	ArrivalPointId     int       `json:"arrivalPointId"`
	ArrivalStationName string    `json:"arrivalStationName"`
	BusInfo            string    `json:"busInfo"`
	Carrier            string    `json:"carrier"`
	CarrierInn         string    `json:"carrierInn"`
	Clazz              struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"clazz"`
	DataRequired        bool      `json:"dataRequired"`
	DepotId             int       `json:"depotId"`
	DispatchDate        time.Time `json:"dispatchDate"`
	DispatchPointId     int       `json:"dispatchPointId"`
	DispatchStationName string    `json:"dispatchStationName"`
	FreeSeatCount       int       `json:"freeSeatCount"`
	FreeSeatEstimation  string    `json:"freeSeatEstimation"`
	FromCache           bool      `json:"fromCache"`
	Name                string    `json:"name"`
	Num                 string    `json:"num"`
	Price               float64   `json:"price"`
	Status              struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	SupplierPrice float64 `json:"supplierPrice"`
	Type          struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Uid string `json:"uid"`
}

type RaceSummaryDTO struct {
	Depot struct {
		Address          string `json:"address"`
		BookingTimeLimit int    `json:"bookingTimeLimit"`
		Engine           string `json:"engine"`
		Features         string `json:"features"`
		Id               int    `json:"id"`
		Info             string `json:"info"`
		Latitude         string `json:"latitude"`
		Longitude        string `json:"longitude"`
		Name             string `json:"name"`
		Online           bool   `json:"online"`
		PhoneRequired    bool   `json:"phoneRequired"`
		Phones           string `json:"phones"`
		PrintInfo        string `json:"printInfo"`
		ReturnInfo       string `json:"returnInfo"`
		Site             string `json:"site"`
		TicketLimit      int    `json:"ticketLimit"`
		Timezone         string `json:"timezone"`
		Version          string `json:"version"`
		WorkingTime      string `json:"workingTime"`
	} `json:"depot"`
	DocTypes []struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"docTypes"`
	Race struct {
		ArrivalDate        time.Time `json:"arrivalDate"`
		ArrivalPointId     int       `json:"arrivalPointId"`
		ArrivalStationName string    `json:"arrivalStationName"`
		BusInfo            string    `json:"busInfo"`
		Carrier            string    `json:"carrier"`
		CarrierInn         string    `json:"carrierInn"`
		Clazz              struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"clazz"`
		DataRequired        bool      `json:"dataRequired"`
		DepotId             int       `json:"depotId"`
		DispatchDate        time.Time `json:"dispatchDate"`
		DispatchPointId     int       `json:"dispatchPointId"`
		DispatchStationName string    `json:"dispatchStationName"`
		FreeSeatCount       int       `json:"freeSeatCount"`
		FreeSeatEstimation  string    `json:"freeSeatEstimation"`
		FromCache           bool      `json:"fromCache"`
		Name                string    `json:"name"`
		Num                 string    `json:"num"`
		Price               float64   `json:"price"`
		Status              struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"status"`
		SupplierPrice float64 `json:"supplierPrice"`
		Type          struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Uid string `json:"uid"`
	} `json:"race"`
	Seats []struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"seats"`
	Stops []struct {
		ArrivalDate  time.Time `json:"arrivalDate"`
		Code         string    `json:"code"`
		DispatchDate time.Time `json:"dispatchDate"`
		Distance     int       `json:"distance"`
		Name         string    `json:"name"`
		RegionName   string    `json:"regionName"`
		StopTime     int       `json:"stopTime"`
	} `json:"stops"`
	TicketTypes []struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		TicketClass string  `json:"ticketClass"`
	} `json:"ticketTypes"`
}

type SaleDTO struct {
	Birthday       time.Time `json:"birthday"`
	Citizenship    string    `json:"citizenship"`
	DocNum         string    `json:"docNum"`
	DocSeries      string    `json:"docSeries"`
	DocTypeCode    string    `json:"docTypeCode"`
	FirstName      string    `json:"firstName"`
	Gender         string    `json:"gender"`
	LastName       string    `json:"lastName"`
	MiddleName     string    `json:"middleName"`
	Phone          string    `json:"phone"`
	SeatCode       string    `json:"seatCode"`
	TicketTypeCode string    `json:"ticketTypeCode"`
}

type BusOrderDTO struct {
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

type TourDTO struct {
	First struct {
		RaceUid string `json:"raceUid"`
		Sales   []struct {
			Birthday       time.Time `json:"birthday"`
			Citizenship    string    `json:"citizenship"`
			DocNum         string    `json:"docNum"`
			DocSeries      string    `json:"docSeries"`
			DocTypeCode    string    `json:"docTypeCode"`
			FirstName      string    `json:"firstName"`
			Gender         string    `json:"gender"`
			LastName       string    `json:"lastName"`
			MiddleName     string    `json:"middleName"`
			Phone          string    `json:"phone"`
			SeatCode       string    `json:"seatCode"`
			TicketTypeCode string    `json:"ticketTypeCode"`
		} `json:"sales"`
	} `json:"first"`
	Second struct {
		RaceUid string `json:"raceUid"`
		Sales   []struct {
			Birthday       time.Time `json:"birthday"`
			Citizenship    string    `json:"citizenship"`
			DocNum         string    `json:"docNum"`
			DocSeries      string    `json:"docSeries"`
			DocTypeCode    string    `json:"docTypeCode"`
			FirstName      string    `json:"firstName"`
			Gender         string    `json:"gender"`
			LastName       string    `json:"lastName"`
			MiddleName     string    `json:"middleName"`
			Phone          string    `json:"phone"`
			SeatCode       string    `json:"seatCode"`
			TicketTypeCode string    `json:"ticketTypeCode"`
		} `json:"sales"`
	} `json:"second"`
}

func (b *bus) GetRaces(ctx context.Context, fromID, toID int, date string, count int) (*[]RaceDTO, error) {
	v := url.Values{
		"from":  []string{strconv.Itoa(fromID)},
		"to":    []string{strconv.Itoa(toID)},
		"date":  []string{date},
		"count": []string{strconv.Itoa(count)},
	}
	u := b.createUrl("/v1/races", v)
	races := &[]RaceDTO{}
	if err := requests.GetRequest(ctx, u, races); err != nil {
		return nil, err
	}
	return races, nil
}

func (b *bus) GetRaceSummary(ctx context.Context, raceUID string) (*RaceSummaryDTO, error) {
	u := b.createUrl(fmt.Sprintf("/v1/races/%s/summary", raceUID), nil)
	raceSummary := &RaceSummaryDTO{}
	if err := requests.GetRequest(ctx, u, raceSummary); err != nil {
		return nil, err
	}
	return raceSummary, nil
}

func (b *bus) OrderOnOneRace(ctx context.Context, raceUID string, sales []SaleDTO, accessToken string) (*BusOrderDTO, error) {
	u := b.createUrl(fmt.Sprintf("/v1/races/%s/orders", raceUID), nil)
	busOrder := &BusOrderDTO{}
	if err := requests.PostRequest(ctx, u, sales, busOrder, auth(accessToken)); err != nil {
		return nil, err
	}
	return busOrder, nil
}

func (b *bus) OrderOnTwoRace(ctx context.Context, tourDTO TourDTO, accessToken string) (*BusOrderDTO, error) {
	u := b.createUrl("/v1/races/orders", nil)
	busOrder := &BusOrderDTO{}
	if err := requests.PostRequest(ctx, u, tourDTO, busOrder, auth(accessToken)); err != nil {
		return nil, err
	}
	return busOrder, nil
}
