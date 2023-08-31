package waypoints

import "github.com/team-four-fingers/kakao/mobility/common"

type NavigateRouteThroughWaypointsRequest struct {
	Origin       common.Location   `json:"origin"`
	Destination  common.Location   `json:"destination"`
	Waypoints    []common.Location `json:"waypoints,omitempty"`
	Priority     *string           `json:"priority,omitempty"`
	Avoid        []string          `json:"avoid,omitempty"`
	CarFuel      *string           `json:"car_fuel,omitempty"`
	CarType      *int              `json:"car_type,omitempty"`
	CarHipass    *bool             `json:"car_hipass,omitempty"`
	Alternatives *bool             `json:"alternatives,omitempty"`
	RoadDetails  *bool             `json:"road_details,omitempty"`
	Summary      *bool             `json:"summary,omitempty"`
}

type NavigateRouteThroughWaypointsResponse struct {
	TransId string  `json:"trans_id"`
	Routes  []Route `json:"routes"`
}

type Route struct {
	ResultCode int       `json:"result_code"`
	ResultMsg  string    `json:"result_msg"`
	Summary    Summary   `json:"summary"`
	Sections   []Section `json:"sections"`
}

type Section struct {
	Distance int           `json:"distance"`
	Duration int           `json:"duration"`
	Bound    *common.Bound `json:"bound,omitempty"`
	Roads    []Road        `json:"roads,omitempty"`
	Guides   []Guide       `json:"guides,omitempty"`
}

type Summary struct {
	Origin      common.Location   `json:"origin"`
	Destination common.Location   `json:"destination"`
	Waypoints   []common.Location `json:"waypoints"`
	Priority    string            `json:"priority"`
	Bound       *common.Bound     `json:"bound,omitempty"`
	Fare        common.Fare       `json:"fare"`
	Distance    int               `json:"distance"`
	Duration    int               `json:"duration"`
}

type Road struct {
	Name         string    `json:"name"`
	Distance     int       `json:"distance"`
	Duration     int       `json:"duration"`
	TrafficSpeed float64   `json:"traffic_speed"`
	TrafficState int       `json:"traffic_state"`
	Vertexes     []float64 `json:"vertexes"`
}

type Guide struct {
	Name      string  `json:"name"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Distance  int     `json:"distance"`
	Duration  int     `json:"duration"`
	Type      int     `json:"type"`
	Guidance  string  `json:"guidance"`
	RoadIndex int     `json:"road_index"`
}
