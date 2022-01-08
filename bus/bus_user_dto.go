package bus

type UserDTO struct {
	ConfirmEmail bool   `json:"confirmEmail"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	Id           int    `json:"id"`
	LastName     string `json:"lastName"`
	MiddleName   string `json:"middleName"`
	Phone        string `json:"phone"`
}
