package checks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httputil"

	"github.com/fatih/color"
)

// Defining colors
var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var cyan = color.New(color.FgCyan)

func ApiChecks(api string, poc bool) {
	fmt.Printf("%v Performing checks using %v\n", cyan.Sprintf("[i]"), yellow.Sprintf(api))
	/*
		CustomSearchAPI(api, poc)
		StaticMapAPI(api, poc)
		StreetViewAPI(api, poc)
		EmbedBasicAPI(api, poc)
		EmbedAdvancedAPI(api, poc)
		DirectionsAPI(api, poc)
		GeocodeAPI(api, poc)
		DistanceMatrixAPI(api, poc)
		FindPlaceFromTextAPI(api, poc)
		AutocompleteAPI(api, poc)
		ElevationAPI(api, poc)
		TimezoneAPI(api, poc)
		NearestRoadsAPI(api, poc)
		GeolocationAPI(api, poc)
		RouteToTraveledAPI(api, poc)
		SpeedLimitRoadsAPI(api, poc)
		PlaceDetailsAPI(api, poc)
		NearbySearchPlacesAPI(api, poc)
		TextSearchPlacesAPI(api, poc)
		PlacesPhotoAPI(api, poc)
		PlayableLocationsAPI(api, poc)
		FCMAPI(api, poc)
	*/
	QueryAutocompletePlaces(api, poc)

}

func CustomSearchAPI(api string, poc bool) {
	var url = `https://www.googleapis.com/customsearch/v1?cx=017576662512468239146:omuauf_lfve&q=lectures&key=` + api
	response := MakeRequest(url)
	var data CustomSearchAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to DirectionsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to DirectionsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	}
}

func StaticMapAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/staticmap?center=45%2C10&zoom=7&size=400x400&key=` + api
	response := MakeRequest(url)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to StaticMapAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to StaticMapAPI"))
	}
}

func StreetViewAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/streetview?size=400x400&location=40.720032,-73.988354&fov=90&heading=235&pitch=10&key=` + api
	response := MakeRequest(url)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to StreetViewAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to StreetViewAPI"))
	}
	defer response.Body.Close()
}

func EmbedBasicAPI(api string, poc bool) {
	var url = `https://www.google.com/maps/embed/v1/place?q=Seattle&key=` + api
	var iframe = fmt.Sprintf(`<iframe width="600" height="450" frameborder="0" style="border:0" src="%s" allowfullscreen></iframe>`, url)
	response := MakeRequest(url)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to EmbedBasicAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC iframe:"), iframe)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to EmbedBasicAPI"))
	}
}

func EmbedAdvancedAPI(api string, poc bool) {
	var url = `https://www.google.com/maps/embed/v1/search?q=record+stores+in+Seattle&key=` + api
	var iframe = fmt.Sprintf(`<iframe width="600" height="450" frameborder="0" style="border:0" src="%s" allowfullscreen></iframe>`, url)
	response := MakeRequest(url)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to EmbedAdvancedAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC iframe:"), iframe)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to EmbedAdvancedAPI"))
	}

}

func DirectionsAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/directions/json?origin=Disneyland&destination=Universal+Studios+Hollywood4&key=` + api
	response := MakeRequest(url)

	var data DirectionsAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to DirectionsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to DirectionsAPI"))
	}
}

func GeocodeAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/geocode/json?latlng=40,30&key=` + api
	response := MakeRequest(url)
	var data GeocodeAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to GeocodeAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to GeocodeAPI"))
	}
}

func DistanceMatrixAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626&key=` + api
	response := MakeRequest(url)
	var data DistanceMatrixAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to DistanceMatrixAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to DistanceMatrixAPI"))
	}
}

func FindPlaceFromTextAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=photos,formatted_address,name,rating,opening_hours,geometry&key=` + api
	response := MakeRequest(url)
	var data FindPlaceFromTextAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to FindPlaceFromTextAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to FindPlaceFromTextAPI"))
	}
}

func AutocompleteAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/autocomplete/json?input=Bingh&types=%28cities%29&key=` + api
	response := MakeRequest(url)
	var data AutocompleteAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to AutocompleteAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to AutocompleteAPI"))
	}
}

func ElevationAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/elevation/json?locations=39.7391536,-104.9847034&key=` + api
	response := MakeRequest(url)
	var data ElevationAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to ElevationAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to ElevationAPI"))
	}
}

func TimezoneAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/timezone/json?location=39.6034810,-119.6822510&timestamp=1331161200&key=` + api
	response := MakeRequest(url)
	var data TimezoneAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to TimezoneAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to TimezoneAPI"))
	}
}

func NearestRoadsAPI(api string, poc bool) {
	var url = `https://roads.googleapis.com/v1/nearestRoads?points=60.170880,24.942795|60.170879,24.942796|60.170877,24.942796&key=` + api
	response := MakeRequest(url)
	var data NearestRoadsAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to NearestRoadsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to NearestRoadsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	}
}

func GeolocationAPI(api string, poc bool) {
	var url = `https://www.googleapis.com/geolocation/v1/geolocate?key=`
	var postData = []byte(`{"considerIp": true}`)
	request, response := MakePostRequest(url, postData, api)
	var data GeolocationAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to GeolocationAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to GeolocationAPI"))
		if poc {
			req, err := httputil.DumpRequest(request, false)
			CheckErr(err)
			fmt.Printf("%v\n%s%s\n\n", yellow.Sprintf("[!] PoC Request:"), string(req), string(postData))
		}
	}

}

func RouteToTraveledAPI(api string, poc bool) {
	var url = `https://roads.googleapis.com/v1/snapToRoads?path=-35.27801,149.12958|-35.28032,149.12907&interpolate=true&key=` + api
	response := MakeRequest(url)
	var data RouteToTraveledAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to RouteToTraveledAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to RouteToTraveledAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	}
}

func SpeedLimitRoadsAPI(api string, poc bool) {
	var url = `https://roads.googleapis.com/v1/speedLimits?path=38.75807927603043,-9.03741754643809&key=` + api
	response := MakeRequest(url)
	var data SpeedLimitRoadsAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to SpeedLimitRoadsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to SpeedLimitRoadsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	}
}

func PlaceDetailsAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/details/json?place_id=ChIJN1t_tDeuEmsRUsoyG83frY4&fields=name,rating,formatted_phone_number&key=` + api
	response := MakeRequest(url)
	var data PlaceDetailsAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to PlaceDetailsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to PlaceDetailsAPI"))
	}
}

func NearbySearchPlacesAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=100&types=food&name=harbour&key=` + api
	response := MakeRequest(url)
	var data NearbySearchPlacesAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to NearbySearchPlacesAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to NearbySearchPlacesAPI"))
	}
}

func TextSearchPlacesAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=` + api
	response := MakeRequest(url)
	var data TextSearchPlacesAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Status == "OK" {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to TextSearchPlacesAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to TextSearchPlacesAPI"))
	}
}

func PlacesPhotoAPI(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=CnRtAAAATLZNl354RwP_9UKbQ_5Psy40texXePv4oAlgP4qNEkdIrkyse7rPXYGd9D_Uj1rVsQdWT4oRz4QrYAJNpFX7rzqqMlZw2h2E2y5IKMUZ7ouD_SlcHxYq1yL4KbKUv3qtWgTK0A6QbGh87GB3sscrHRIQiG2RrmU_jF4tENr9wGS_YxoUSSDrYjWmrNfeEHSGSc3FyhNLlBU&key=` + api
	response := MakeRequest(url)
	defer response.Body.Close()

	if response.StatusCode == 302 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to PlacesPhotoAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to PlacesPhotoAPI"))
	}
}

func PlayableLocationsAPI(api string, poc bool) {
	var url = `https://playablelocations.googleapis.com/v3:samplePlayableLocations?key=`
	var postData = []byte(`{"area_filter":{"s2_cell_id":7715420662885515264},"criteria":[{"gameObjectType":1,"filter":{"maxLocationCount":4,"includedTypes":["food_and_drink"]},"fields_to_return": {"paths": ["name"]}},{"gameObjectType":2,"filter":{"maxLocationCount":4},"fields_to_return": {"paths": ["types", "snapped_point"]}}]}`)
	request, response := MakePostRequest(url, postData, api)
	var data PlayableLocationsAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to PlayableLocationsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to PlayableLocationsAPI"))
		if poc {
			req, err := httputil.DumpRequest(request, false)
			CheckErr(err)
			fmt.Printf("%v\n%s%s\n\n", yellow.Sprintf("[!] PoC Request:"), string(req), string(postData))
		}
	}

}

func FCMAPI(api string, poc bool) {
	var url = `https://fcm.googleapis.com/fcm/send`
	var postData = []byte(`{"registration_ids":["ABC"]}`)
	request, response := MakePostRequest(url, postData, api)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to FCMAPI"))
		if poc {
			req, err := httputil.DumpRequest(request, false)
			CheckErr(err)
			fmt.Printf("%v\n%s%s\n\n", yellow.Sprintf("[!] PoC Request:"), string(req), string(postData))
		}
	} else {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to FCMAPI"))
	}
}

func QueryAutocompletePlaces(api string, poc bool) {
	var url = `https://maps.googleapis.com/maps/api/place/queryautocomplete/json?input=pizza+near%20par&key=` + api
	response := MakeRequest(url)
	var data CustomSearchAPIStruct
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &data)
	CheckErr(err)

	if data.Error.Status == "PERMISSION_DENIED" {
		fmt.Printf("%v\n", green.Sprintf("[+] Not vulnerable to DirectionsAPI"))
	} else {
		fmt.Printf("%v\n", red.Sprintf("[-] Vulnerable to DirectionsAPI"))
		if poc {
			fmt.Printf("%v %s\n\n", yellow.Sprintf("[!] PoC URL:"), url)
		}
	}
}
