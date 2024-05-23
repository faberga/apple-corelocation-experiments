package main

import (
	"fmt"
	"geostalk/lib"
	"log"
	"math"

	"github.com/gptlang/oui/ouidb"
	"github.com/leaanthony/clir"
)

func main() {
	cli := clir.NewCli("geostalk", "Retrieve BSSID geolocation using Apple's API", "v0.0.1")
	getCmd := cli.NewSubCommand("get", "Gets and displays adjacent BSSID locations given an existing BSSID")
	var bssids []string
	getCmd.StringsFlag("bssid", "One or more known bssid strings", &bssids)
	getCmd.Action(func() error {
		if len(bssids) == 0 {
			log.Fatalln("BSSIDs cannot be empty")
		}
		blocks := lib.QueryBssid(bssids, false)
		for _, wifi := range blocks.GetWifiDevices() {
			man, err := ouidb.Lookup(wifi.GetBssid())
			if err != nil {
				man = "Unknown"
			}
			fmt.Printf("BSSID: %s (%s) found at Lat: %f Long: %f\n", wifi.GetBssid(), man, float64(*wifi.GetLocation().Latitude)*math.Pow10(-8), float64(*wifi.GetLocation().Longitude)*math.Pow10(-8))
		}
		fmt.Println(len(blocks.GetWifiDevices()), "number of devices found in area")
		return nil
	})
	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}