package maps

import (
	"asik3/vendors/twogis"
)

type TwoGisAdapter struct {
	client *twogis.Client
}

func NewTwoGisAdapter(c *twogis.Client) *TwoGisAdapter {
	return &TwoGisAdapter{client: c}
}
