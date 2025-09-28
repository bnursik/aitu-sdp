package twogis

import "context"

type Client struct{ Token string }

type DistanceReply struct{ Meters int }

func (c *Client) RouteDistance(ctx context.Context, fromLat, fromLng, toLat, toLng float64) (DistanceReply, error) {
	// Imagine a real HTTP call here.
	return DistanceReply{Meters: 12800}, nil 
}
