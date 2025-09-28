package booking

import (
	"context"

	"asik3/maps"
)

func IsWithinRadius(ctx context.Context, mapsService maps.IMapsService, pickup, branch maps.LatLng, maxKm float64) (bool, float64, error) {
	d, err := mapsService.DistanceKm(ctx, pickup, branch)
	if err != nil {
		return false, 0, err
	}
	return d <= maxKm, d, nil
}


