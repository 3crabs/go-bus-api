package bus

import "time"

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
