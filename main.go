package main

import (
	"flag"
	"fmt"
)

var (
	apiFlag     string
	proxyFlag   string
	pocFlag     bool
	quietFlag   bool
	versionFlag bool
	version     string
)

func init() {
	flag.StringVar(&apiFlag, "api", "", "Google Maps API key")
	flag.StringVar(&proxyFlag, "x", "", "Proxy URL. Ex: http://127.0.0.1:8080")
	flag.BoolVar(&pocFlag, "poc", false, "Generate PoC for vulnerable ones")
	flag.BoolVar(&quietFlag, "quiet", false, "Print only vulnerable APIs")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Printf("\n     🗺  GAP %v\n\n", version)
		return
	}
	if proxyFlag != "" {
		validateProxyUrl(proxyFlag)
	}
	if apiFlag == "" {
		flag.PrintDefaults()
		return
	} else {
		validateGoogleMapsApiKey(apiFlag)
		ApiChecks(apiFlag, proxyFlag, pocFlag, quietFlag)
	}
}
