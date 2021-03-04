# GAP

Google API checker.

Based on the study [Unauthorized Google Maps API Key Usage Cases, and Why You Need to Care](https://ozguralp.medium.com/unauthorized-google-maps-api-key-usage-cases-and-why-you-need-to-care-1ccb28bf21e) and [Google Maps API (Not the Key) Bugs That I Found Over the Years](https://ozguralp.medium.com/google-maps-api-not-the-key-bugs-that-i-found-over-the-years-781840fc82aa).

## Checks performed

+ `CustomSearchAPI`
+ `StaticMapAPI`
+ `StreetViewAPI`
+ `EmbedBasicAPI`
+ `EmbedAdvancedAPI`
+ `DirectionsAPI`
+ `GeocodeAPI`
+ `DistanceMatrixAPI`
+ `FindPlaceFromTextAPI`
+ `AutocompleteAPI`
+ `ElevationAPI`
+ `TimezoneAPI`
+ `NearestRoadsAPI`
+ `GeolocationAPI`
+ `RouteToTraveledAPI`
+ `SpeedLimitRoadsAPI`
+ `PlaceDetailsAPI`
+ `NearbySearchPlacesAPI`
+ `TextSearchPlacesAPI`
+ `PlacesPhotoAPI`
+ `PlayableLocationsAPI`
+ `FCMAPI`

## USAGE

```py
$> gap -api "AIza[CONFIDENTIAL]" -poc

[i] Performing checks using AIza[CONFIDENTIAL]
[+] Not vulnerable to DirectionsAPI
[+] Not vulnerable to StaticMapAPI
[+] Not vulnerable to StreetViewAPI
[+] Not vulnerable to EmbedBasicAPI
[+] Not vulnerable to EmbedAdvancedAPI
[+] Not vulnerable to DirectionsAPI
[-] Vulnerable to GeocodeAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/geocode/json?latlng=40,30&key=AIza[CONFIDENTIAL]

[-] Vulnerable to DistanceMatrixAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626&key=AIza[CONFIDENTIAL]

[-] Vulnerable to FindPlaceFromTextAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=photos,formatted_address,name,rating,opening_hours,geometry&key=AIza[CONFIDENTIAL]

[-] Vulnerable to AutocompleteAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/place/autocomplete/json?input=Bingh&types=%28cities%29&key=AIza[CONFIDENTIAL]

[+] Not vulnerable to ElevationAPI
[+] Not vulnerable to TimezoneAPI
[+] Not vulnerable to NearestRoadsAPI
[-] Vulnerable to GeolocationAPI
[!] PoC Request:
POST /geolocation/v1/geolocate?key=AIza[CONFIDENTIAL] HTTP/1.1
Host: www.googleapis.com
Content-Type: application/json

{"considerIp": true}

[+] Not vulnerable to RouteToTraveledAPI
[+] Not vulnerable to SpeedLimitRoadsAPI
[-] Vulnerable to PlaceDetailsAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/place/details/json?place_id=ChIJN1t_tDeuEmsRUsoyG83frY4&fields=name,rating,formatted_phone_number&key=AIza[CONFIDENTIAL]

[-] Vulnerable to NearbySearchPlacesAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=100&types=food&name=harbour&key=AIza[CONFIDENTIAL]

[-] Vulnerable to TextSearchPlacesAPI
[!] PoC URL: https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=AIza[CONFIDENTIAL]

[+] Not vulnerable to PlacesPhotoAPI
[+] Not vulnerable to PlayableLocationsAPI
[+] Not vulnerable to FCMAPI
```