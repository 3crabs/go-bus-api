package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-requests/go-requests"
	"net/url"
)

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

func (b *bus) GetPointsFrom(ctx context.Context, pattern string) (*[]PointDTO, error) {
	v := url.Values{
		"pattern": []string{pattern},
	}
	u := b.createUrl("/v1/points/from", v)
	points := &[]PointDTO{}
	if err := requests.GetRequest(ctx, u, points); err != nil {
		return nil, err
	}
	return points, nil
}

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
