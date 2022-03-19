package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
)

type FeedbackRequest struct {
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

func (b *bus) SendFeedback(ctx context.Context, feedback FeedbackRequest) error {
	u := b.createUrl("/v1/feedbacks", nil)
	return requests.PostRequest(ctx, u, feedback, nil)
}
