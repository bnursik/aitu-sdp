package main

import (
	"context"
	"fmt"

	"asik3/booking"
	"asik3/maps"
	"asik3/vendors/twogis"
)

func main() {
	cli := &twogis.Client{Token: "demo"}
	var ms maps.IMapsService = maps.NewTwoGisAdapter(cli)

	ok, d, err := booking.IsWithinRadius(
		context.Background(),
		ms,
		maps.LatLng{Latitude: 51.128, Longitude: 71.430},
		maps.LatLng{Latitude: 51.160, Longitude: 71.470},
		15.0,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Within radius? %v (%.2f km)\n", ok, d)
}
