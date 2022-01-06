package bus

import "time"

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
