package main

import (
	"flag"
	"fmt"

	"github.com/joanbono/gap/utils/checks"
)

var (
	apiFlag     string
	pocFlag     bool
	versionFlag bool
	version     string
)

func init() {
	flag.StringVar(&apiFlag, "api", "", "Google Maps API key")
	flag.BoolVar(&pocFlag, "poc", false, "Generate PoC for vulnerable ones")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Printf("\nGAP %v\n\n", version)
		return
	}
	if apiFlag == "" {
		flag.PrintDefaults()
		return
	} else {
		checks.ApiChecks(apiFlag, pocFlag)
	}
}
