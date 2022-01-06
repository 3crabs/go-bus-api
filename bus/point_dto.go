package bus

type PointDTO struct {
	Address   string `json:"address"`
	Details   string `json:"details"`
	Id        int    `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Okato     string `json:"okato"`
	Place     bool   `json:"place"`
	Region    string `json:"region"`
}
