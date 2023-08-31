package mobility

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/mobility/common"
	"github.com/team-four-fingers/kakao/mobility/waypoints"
)

const (
	baseURL = "https://apis-navi.kakaomobility.com"
)

type Client interface {
	NavigateRouteThroughWaypoints(req *waypoints.NavigateRouteThroughWaypointsRequest) (
		*waypoints.NavigateRouteThroughWaypointsResponse,
		error,
	)
}

type defaultClient struct {
	*resty.Client
}

func NewClient() (Client, error) {
	restyCli := resty.New()

	if err := core.ConfigureClient(restyCli, core.WithBaseURL(baseURL)); err != nil {
		return nil, err
	}

	return &defaultClient{Client: restyCli}, nil
}

func (d *defaultClient) NavigateRouteThroughWaypoints(
	req *waypoints.NavigateRouteThroughWaypointsRequest,
) (*waypoints.NavigateRouteThroughWaypointsResponse, error) {
	respJSON, err := d.NewRequest().
		SetBody(req).
		Post(waypoints.Path)
	if err != nil {
		return nil, err
	}

	var resp *waypoints.NavigateRouteThroughWaypointsResponse
	if err := json.Unmarshal(respJSON.Body(), &resp); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := string(respJSON.Body())[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
		return nil, err
	}

	foundRoutes := lo.Filter(resp.Routes, func(item waypoints.Route, _ int) bool {
		return item.ResultCode == common.NavigationResultCode.Success
	})

	resp.Routes = foundRoutes

	return resp, nil
}
