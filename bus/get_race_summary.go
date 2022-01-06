package bus

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/rest"
)

func (b *bus) GetRaceSummary(ctx context.Context, raceUID string) (*RaceSummaryDTO, error) {
	u := b.createUrl(fmt.Sprintf("/v1/races/%s/summary", raceUID), nil)
	raceSummary := &RaceSummaryDTO{}
	if err := rest.GetRequest(ctx, u, raceSummary); err != nil {
		return nil, err
	}
	return raceSummary, nil
}
