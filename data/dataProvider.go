package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/BlockscapeLab/goz-data-backend/data/types"
	rest "github.com/BlockscapeLab/goz-data-backend/rest/types"
)

type DataProvider struct {
	beginTime       time.Time
	endTime         time.Time
	lastSyncedBlock int
	teams           map[string]types.TeamDetails
	scoreboard      *types.Scoreboard
	chartData       map[string]types.TeamChart
}

func NewDataProvider(startTime, endTime string) (*DataProvider, error) {
	begin, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		return nil, err
	}
	dp := DataProvider{
		beginTime:       begin,
		endTime:         end,
		lastSyncedBlock: 0,
		teams:           make(map[string]types.TeamDetails),
		chartData:       make(map[string]types.TeamChart),
		scoreboard:      &types.Scoreboard{},
	}

	return &dp, nil
}

func (dp *DataProvider) GetScoreboardJSON() ([]byte, error) {
	return json.MarshalIndent(dp.scoreboard, "", " ")
}

func (dp *DataProvider) GetTeamDetailsJSON(chainID string) ([]byte, error) {
	t, ok := dp.teams[chainID]
	if !ok {
		return nil, rest.NewErrorNotFound(fmt.Sprintf("Couldn't find team with chain ID '%s'", chainID))
	}
	return json.MarshalIndent(t, "", " ")
}

func (dp *DataProvider) GetTeamChartDataJSON(chainID string) ([]byte, error) {
	return nil, errors.New("Not implemented")
}
