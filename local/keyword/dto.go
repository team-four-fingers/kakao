package keyword

type Response struct {
	Meta struct {
		SameName struct {
			Region         []interface{} `json:"region"`
			Keyword        string        `json:"keyword"`
			SelectedRegion string        `json:"selected_region"`
		} `json:"same_name"`
		PageableCount int  `json:"pageable_count"`
		TotalCount    int  `json:"total_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []struct {
		PlaceName         string `json:"place_name"`
		Distance          string `json:"distance"`
		PlaceUrl          string `json:"place_url"`
		CategoryName      string `json:"category_name"`
		AddressName       string `json:"address_name"`
		RoadAddressName   string `json:"road_address_name"`
		Id                string `json:"id"`
		Phone             string `json:"phone"`
		CategoryGroupCode string `json:"category_group_code"`
		CategoryGroupName string `json:"category_group_name"`
		X                 string `json:"x"`
		Y                 string `json:"y"`
	} `json:"documents"`
}
