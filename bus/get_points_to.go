package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"net/url"
)

func (b *bus) GetPointsTo(ctx context.Context, fromID int, pattern string) (*[]PointDTO, error) {
	v := url.Values{
		"pattern": []string{pattern},
	}
	u := b.createUrl(fmt.Sprintf("/v1/points/%d/to", fromID), v)
	points := &[]PointDTO{}
	if err := requests.GetRequest(ctx, u, points); err != nil {
		return nil, err
	}
	return points, nil
}
