package checks

type DirectionsAPIStruct struct {
	ErrorMessage string        `json:"error_message"`
	Routes       []interface{} `json:"routes,omitempty"`
	Status       string        `json:"status"`
}
type GeocodeAPIStruct struct {
	PlusCode interface{}   `json:"plus_code"`
	Results  []interface{} `json:"results"`
	Status   string        `json:"status"`
}

type CustomSearchAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Errors  []interface{} `json:"errors"`
		Status  string        `json:"status"`
	} `json:"error"`
}

type CustomSearchAPIStructError struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
	Status  string        `json:"status"`
}

type DistanceMatrixAPIStruct struct {
	DestinationAddresses []string      `json:"destination_addresses"`
	OriginAddresses      []string      `json:"origin_addresses"`
	Rows                 []interface{} `json:"rows"`
	Status               string        `json:"status"`
}

type FindPlaceFromTextAPIStruct struct {
	Candidates []interface{} `json:"candidates"`
	Status     string        `json:"status"`
}

type AutocompleteAPIStruct struct {
	Predictions []interface{} `json:"predictions"`
	Status      string        `json:"status"`
}

type ElevationAPIStruct struct {
	Results []interface{} `json:"results"`
	Status  string        `json:"status"`
}

type TimezoneAPIStruct struct {
	DstOffset    int    `json:"dstOffset"`
	RawOffset    int    `json:"rawOffset"`
	Status       string `json:"status"`
	TimeZoneID   string `json:"timeZoneId"`
	TimeZoneName string `json:"timeZoneName"`
}

type NearestRoadsAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Status  string        `json:"status"`
		Details []interface{} `json:"details"`
	} `json:"error"`
}

type GeolocationAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Errors  []interface{} `json:"errors"`
		Status  string        `json:"status"`
	} `json:"error"`
}

type RouteToTraveledAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Status  string        `json:"status"`
		Details []interface{} `json:"details"`
	} `json:"error"`
}

type SpeedLimitRoadsAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Status  string        `json:"status"`
		Details []interface{} `json:"details"`
	} `json:"error"`
}

type PlaceDetailsAPIStruct struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result           interface{}   `json:"result"`
	Status           string        `json:"status"`
}

type NearbySearchPlacesAPIStruct struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Results          []interface{} `json:"results"`
	Status           string        `json:"status"`
}

type TextSearchPlacesAPIStruct struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Results          []interface{} `json:"results"`
	Status           string        `json:"status"`
}

type PlayableLocationsAPIStruct struct {
	Error struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Status  string        `json:"status"`
		Details []interface{} `json:"details"`
	} `json:"error"`
}
