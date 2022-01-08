package bus

type FeedbackDTO struct {
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}
