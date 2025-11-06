package observer

import (
	"fmt"

	"asik7/pricing/notify"
)

type PriceAlertListener struct {
	channelName string
}

func NewPriceAlertListener(channelName string) PriceAlertListener {
	return PriceAlertListener{channelName: channelName}
}

func (l PriceAlertListener) OnNotify(e notify.Event) {
	if e.Type == "price-change" {
		fmt.Printf("[ALERT:%s] %s | data=%v\n", l.channelName, e.Message, e.Data)
	}
}
