package common

type Location struct {
	Name *string `json:"name,omitempty"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

type Bound struct {
	MinX float64 `json:"min_x"`
	MinY float64 `json:"min_y"`
	MaxX float64 `json:"max_x"`
	MaxY float64 `json:"max_y"`
}

type Fare struct {
	Taxi int `json:"taxi"`
	Toll int `json:"toll"`
}
