package bus

import "time"

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
