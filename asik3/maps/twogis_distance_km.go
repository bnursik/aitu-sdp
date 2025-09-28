package maps

import "context"

func (a *TwoGisAdapter) DistanceKm(ctx context.Context, from LatLng, to LatLng) (float64, error) {
	reply, err := a.client.RouteDistance(ctx, from.Latitude, from.Longitude, to.Latitude, to.Longitude)
	if err != nil {
		return 0, ErrUpstream
	}
	if reply.Meters < 0 {
		return 0, ErrNegativeDistance
	}
	return float64(reply.Meters) / 1000.0, nil
}

