package main

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/fatih/color"
	"github.com/monaco-io/request"
	"github.com/tidwall/gjson"
)

// Defining colors
var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var cyan = color.New(color.FgCyan)

// validateGoogleMapsApiKey will validate if the
// provided API key match with a Google Maps
// regular expression
func validateGoogleMapsApiKey(apiKey string) {
	match, _ := regexp.MatchString(`AIza[0-9A-Za-z\-_]{35}`, apiKey)
	if !match || len(apiKey) != 39 {
		fmt.Printf("🔑 %s is not a valid Google Maps API key.\n", yellow.Sprintf(apiKey))
		os.Exit(0)
	}
}
func validateProxyUrl(proxyUrl string) {
	_, err := url.ParseRequestURI(proxyUrl)
	if err != nil {
		fmt.Printf("🔗 %s is not a valid Proxy address.\n", yellow.Sprintf(proxyUrl))
		os.Exit(0)
	}
}

// The API Checks calls are made from this
// function. Comment to remove checks
func ApiChecks(api, proxy string, poc bool) {
	fmt.Printf("ℹ️  Performing checks for %v\n", yellow.Sprintf(api))

	CustomSearchAPI(api, proxy, poc)
	StaticMapAPI(api, proxy, poc)
	StreetViewAPI(api, proxy, poc)
	EmbedBasicAPI(api, proxy, poc)
	EmbedAdvancedAPI(api, proxy, poc)
	DirectionsAPI(api, proxy, poc)
	GeocodeAPI(api, proxy, poc)
	DistanceMatrixAPI(api, proxy, poc)
	FindPlaceFromTextAPI(api, proxy, poc)
	AutocompleteAPI(api, proxy, poc)
	ElevationAPI(api, proxy, poc)
	TimezoneAPI(api, proxy, poc)
	NearestRoadsAPI(api, proxy, poc)
	GeolocationAPI(api, proxy, poc)
	RouteToTraveledAPI(api, proxy, poc)
	SpeedLimitRoadsAPI(api, proxy, poc)
	PlaceDetailsAPI(api, proxy, poc)
	NearbySearchPlacesAPI(api, proxy, poc)
	TextSearchPlacesAPI(api, proxy, poc)
	PlacesPhotoAPI(api, proxy, poc)
	PlayableLocationsAPI(api, proxy, poc)
	FCMAPI(api, proxy, poc)
	QueryAutocompletePlaces(api, proxy, poc)

}

func CustomSearchAPI(api, proxy string, poc bool) {
	//println(proxy)
	var url = `https://www.googleapis.com/customsearch/v1?cx=017576662512468239146:omuauf_lfve&q=lectures&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")
	if (resp.Code() == 403 && value.String() == "PERMISSION_DENIED") || (resp.Code() == 400 && value.String() == "INVALID_ARGUMENT") {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to CustomSearchAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to CustomSearchAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	}
}

func StaticMapAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/staticmap?center=45%2C10&zoom=7&size=400x400&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	if resp.Code() == 200 {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to StaticMapAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to StaticMapAPI"))
	}
}

func StreetViewAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/streetview?size=400x400&location=40.720032,-73.988354&fov=90&heading=235&pitch=10&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()

	if resp.Code() == 200 {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to StreetViewAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to StreetViewAPI"))
	}
}

func EmbedBasicAPI(api, proxy string, poc bool) {
	var url = `https://www.google.com/maps/embed/v1/place?q=Seattle&key=` + api
	var iframe = fmt.Sprintf(`<iframe width="600" height="450" frameborder="0" style="border:0" src="%s" allowfullscreen></iframe>`, url)
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	if resp.Code() == 200 {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to EmbedBasicAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC iframe:"), iframe)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to EmbedBasicAPI"))
	}
}

func EmbedAdvancedAPI(api, proxy string, poc bool) {
	var url = `https://www.google.com/maps/embed/v1/search?q=record+stores+in+Seattle&key=` + api
	var iframe = fmt.Sprintf(`<iframe width="600" height="450" frameborder="0" style="border:0" src="%s" allowfullscreen></iframe>`, url)
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()

	if resp.Code() == 200 {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to EmbedAdvancedAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC iframe:"), iframe)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to EmbedAdvancedAPI"))
	}

}

func DirectionsAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/directions/json?origin=Disneyland&destination=Universal+Studios+Hollywood4&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to DirectionsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to DirectionsAPI"))
	}
}

func GeocodeAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/geocode/json?latlng=40,30&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to GeocodeAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to GeocodeAPI"))
	}
}

func DistanceMatrixAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to DistanceMatrixAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to DistanceMatrixAPI"))
	}
}

func FindPlaceFromTextAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=photos,formatted_address,name,rating,opening_hours,geometry&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to FindPlaceFromTextAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to FindPlaceFromTextAPI"))
	}
}

func AutocompleteAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/autocomplete/json?input=Bingh&types=%28cities%29&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to AutocompleteAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to AutocompleteAPI"))
	}
}

func ElevationAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/elevation/json?locations=39.7391536,-104.9847034&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to ElevationAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to ElevationAPI"))
	}
}

func TimezoneAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/timezone/json?location=39.6034810,-119.6822510&timestamp=1331161200&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to TimezoneAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to TimezoneAPI"))
	}
}

func NearestRoadsAPI(api, proxy string, poc bool) {
	var url = `https://roads.googleapis.com/v1/nearestRoads?points=60.170880,24.942795|60.170879,24.942796|60.170877,24.942796&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")

	if (resp.Code() == 403 && value.String() == "PERMISSION_DENIED") || (resp.Code() == 400 && value.String() == "INVALID_ARGUMENT") {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to NearestRoadsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to NearestRoadsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	}
}

func GeolocationAPI(api, proxy string, poc bool) {
	var url = `https://www.googleapis.com/geolocation/v1/geolocate?key=` + api
	var body = struct {
		considerIp bool
	}{considerIp: true}
	c := request.Client{
		URL:    url,
		Method: "POST",
		Header: (map[string]string{
			"User-Agent":    "GAP - The Google Maps API Checker",
			"Authorization": "key=" + api,
			"Content-Type":  "application/json",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		JSON:      body,
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")

	if (resp.Code() == 403 && value.String() == "PERMISSION_DENIED") || (resp.Code() == 400 && value.String() == "INVALID_ARGUMENT") {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to GeolocationAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to GeolocationAPI"))
		if poc {
			command := fmt.Sprintf("curl -X POST -k -s -H \"Authorization: key=%s\" -H \"Content-Type: application/json\" -A \"GAP - The Google Maps API Checker\" -d '{\"considerIp\": true}' \"%s\"", api, url)
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC COMMAND:"), command)
		}
	}

}

func RouteToTraveledAPI(api, proxy string, poc bool) {
	var url = `https://roads.googleapis.com/v1/snapToRoads?path=-35.27801,149.12958|-35.28032,149.12907&interpolate=true&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")
	if (resp.Code() == 403 && value.String() == "PERMISSION_DENIED") || (resp.Code() == 400 && value.String() == "INVALID_ARGUMENT") {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to RouteToTraveledAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to RouteToTraveledAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	}
}

func SpeedLimitRoadsAPI(api, proxy string, poc bool) {
	var url = `https://roads.googleapis.com/v1/speedLimits?path=38.75807927603043,-9.03741754643809&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")

	if (resp.Code() == 403 && value.String() == "PERMISSION_DENIED") || (resp.Code() == 400 && value.String() == "INVALID_ARGUMENT") {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to SpeedLimitRoadsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to SpeedLimitRoadsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	}
}

func PlaceDetailsAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/details/json?place_id=ChIJN1t_tDeuEmsRUsoyG83frY4&fields=name,rating,formatted_phone_number&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to PlaceDetailsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to PlaceDetailsAPI"))
	}
}

func NearbySearchPlacesAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=100&types=food&name=harbour&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to NearbySearchPlacesAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to NearbySearchPlacesAPI"))
	}
}

func TextSearchPlacesAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to TextSearchPlacesAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to TextSearchPlacesAPI"))
	}
}

func PlacesPhotoAPI(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=CnRtAAAATLZNl354RwP_9UKbQ_5Psy40texXePv4oAlgP4qNEkdIrkyse7rPXYGd9D_Uj1rVsQdWT4oRz4QrYAJNpFX7rzqqMlZw2h2E2y5IKMUZ7ouD_SlcHxYq1yL4KbKUv3qtWgTK0A6QbGh87GB3sscrHRIQiG2RrmU_jF4tENr9wGS_YxoUSSDrYjWmrNfeEHSGSc3FyhNLlBU&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()

	if resp.Code() == 200 {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to PlacesPhotoAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to PlacesPhotoAPI"))
	}
}

func PlayableLocationsAPI(api, proxy string, poc bool) {
	var url = `https://playablelocations.googleapis.com/v3:samplePlayableLocations?key=` + api
	c := request.Client{
		URL:    url,
		Method: "POST",
		Header: (map[string]string{
			"User-Agent":    "GAP - The Google Maps API Checker",
			"Authorization": "key=" + api,
			"Content-Type":  "application/json",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		JSON:      []byte(`{"area_filter":{"s2_cell_id":7715420662885515264},"criteria":[{"gameObjectType":1,"filter":{"maxLocationCount":4,"includedTypes":["food_and_drink"]},"fields_to_return": {"paths": ["name"]}},{"gameObjectType":2,"filter":{"maxLocationCount":4},"fields_to_return": {"paths": ["types", "snapped_point"]}}]}`),
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "error.status")

	if (resp.Code() != 403 && value.String() == "PERMISSION_DENIED") || resp.Code() != 404 {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to PlayableLocationsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to PlayableLocationsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	}
}

func FCMAPI(api, proxy string, poc bool) {
	var url = `https://fcm.googleapis.com/fcm/send`
	c := request.Client{
		URL:    url,
		Method: "POST",
		Header: (map[string]string{
			"User-Agent":    "GAP - The Google Maps API Checker",
			"Authorization": "key=" + api,
			"Content-Type":  "application/json",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		JSON:      []byte(`{"registration_ids":["ABC"]}`),
	}

	resp := c.Send()

	if resp.Code() != 200 {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to FCMAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to FCMAPI"))
		if poc {
			command := fmt.Sprintf("curl -k -s -X POST -H \"Content-Type: application/json\" -H \"Authorization: key=%s\" -A \"GAP - The Google Maps API Checker\" -d '{\"registration_ids\":[\"ABC\"]}' \"%s\"", api, url)
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC Command: %s"), command)
		}
	}
}

func QueryAutocompletePlaces(api, proxy string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/queryautocomplete/json?input=pizza+near%20par&key=` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "GAP - The Google Maps API Checker",
		}),
		ProxyURL:  proxy,
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	value := gjson.Get(resp.String(), "status")

	if resp.Code() == 200 && value.String() == "OK" {
		fmt.Printf("%v\n", red.Sprintf("❌ Vulnerable to QueryAutocompletePlaces"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("⚠️  PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("✅ Not vulnerable to QueryAutocompletePlaces"))
	}
}
