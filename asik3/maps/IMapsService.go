package maps

import "context"

type IMapsService interface {
	DistanceKm(ctx context.Context, from LatLng, to LatLng) (float64, error)
}