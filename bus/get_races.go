package bus

import (
	"context"
	"github.com/3crabs/go-requests/go-requests"
	"net/url"
	"strconv"
)

func (b *bus) GetRaces(ctx context.Context, fromID, toID int, date string, count int) (*[]RaceDTO, error) {
	v := url.Values{
		"from":  []string{strconv.Itoa(fromID)},
		"to":    []string{strconv.Itoa(toID)},
		"date":  []string{date},
		"count": []string{strconv.Itoa(count)},
	}
	u := b.createUrl("/v1/races", v)
	races := &[]RaceDTO{}
	if err := requests.GetRequest(ctx, u, races); err != nil {
		return nil, err
	}
	return races, nil
}
