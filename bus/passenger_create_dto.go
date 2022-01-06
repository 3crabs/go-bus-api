package bus

import "time"

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
