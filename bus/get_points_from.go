package bus

import (
	"context"
	"github.com/3crabs/go-bus-api/rest"
	"net/url"
)

func (b *bus) GetPointsFrom(ctx context.Context, pattern string) (*[]PointDTO, error) {
	v := url.Values{
		"pattern": []string{pattern},
	}
	u := b.createUrl("/v1/points/from", v)
	points := &[]PointDTO{}
	if err := rest.GetRequest(ctx, u, points); err != nil {
		return nil, err
	}
	return points, nil
}
