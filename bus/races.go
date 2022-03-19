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
