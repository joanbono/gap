package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	apiFlag     string
	proxyFlag   string
	pocFlag     bool
	quietFlag   bool
	versionFlag bool
	version     string
	headerFlags headerList
)

// headerList collects repeated -X "Name: Value" flags
type headerList []string

func (h *headerList) String() string     { return strings.Join(*h, ", ") }
func (h *headerList) Set(v string) error { *h = append(*h, v); return nil }

func init() {
	flag.StringVar(&apiFlag, "api", "", "Google Maps API key")
	flag.StringVar(&proxyFlag, "x", "", "Proxy URL. Ex: http://127.0.0.1:8080")
	flag.BoolVar(&pocFlag, "poc", false, "Generate PoC for vulnerable ones")
	flag.BoolVar(&quietFlag, "quiet", false, "Print only vulnerable APIs")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.Var(&headerFlags, "X", "Custom header, repeatable. Ex: -X \"Header1: Value\" -X \"Header2: Value\"")
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
	parseCustomHeaders(headerFlags)
	if apiFlag == "" {
		flag.PrintDefaults()
		return
	} else {
		validateGoogleMapsApiKey(apiFlag)
		ApiChecks(apiFlag, proxyFlag, pocFlag, quietFlag)
	}
}
