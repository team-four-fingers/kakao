package mobility

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/mobility/common"
	"github.com/team-four-fingers/kakao/mobility/waypoints"
	"os"
)

func ExampleNewClient() {
	kakaoRESTAPIKey := os.Getenv("KAKAO_REST_API_KEY")

	if err := core.InitializeSDK(kakaoRESTAPIKey); err != nil {
		fmt.Println("err:", err)
		return
	}

	cli, err := NewClient()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	resp, err := cli.NavigateRouteThroughWaypoints(&waypoints.NavigateRouteThroughWaypointsRequest{
		//	"origin": {
		//		"x": 126.946362033068,         "y": 37.5404741779088
		//	},
		//	"destination": {
		//		"x": 127.1101250888609,         "y": 37.39407843730005
		//	},
		//	"waypoints": [
		//{
		//	"coordinate": {
		//	"x": 126.92716700037366,         "y": 37.5266641708316
		//},

		Origin: common.Location{
			X: 126.946362033068,
			Y: 37.5404741779088,
		},
		Destination: common.Location{
			X: 127.1101250888609,
			Y: 37.39407843730005,
		},
		Waypoints: []common.Location{
			{
				X: 126.92716700037366,
				Y: 37.5266641708316,
			},
		},
		Priority:     lo.ToPtr(common.PathPriority.Recommend),
		CarFuel:      lo.ToPtr(common.CarFuel.Gasoline),
		CarHipass:    lo.ToPtr(false),
		Alternatives: lo.ToPtr(false),
		RoadDetails:  lo.ToPtr(false),
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received: ", resp)
	// Output: response received:  0
}
