package bus

import (
	"context"
	"github.com/3crabs/go-bus-api/rest"
)

func (b *bus) SendFeedback(ctx context.Context, feedback FeedbackDTO) error {
	u := b.createUrl("/v1/feedbacks", nil)
	return rest.PostRequest(ctx, u, feedback, nil)
}
