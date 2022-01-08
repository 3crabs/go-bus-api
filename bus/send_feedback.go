package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
)

func (b *bus) SendFeedback(ctx context.Context, feedback FeedbackDTO) error {
	u := b.createUrl("/v1/feedbacks", nil)
	return requests.PostRequest(ctx, u, feedback, nil)
}
