package main

import (
	"asik7/pricing/notify"
	"asik7/pricing/notify/concrete/observer"
	"asik7/pricing/notify/concrete/subject"
)

func main() {
	pricePub := subject.NewPriceSubject()

	alerts := observer.NewPriceAlertListener("EmailCampaign")
	audit := observer.NewAuditLogger("PricingService")

	pricePub.Subscribe(alerts)
	pricePub.Subscribe(audit)

	pricePub.Notify(notify.Event{
		Type:    "price-change",
		Message: "SUV weekend promo: -10%",
		Data: map[string]string{
			"vehicle": "SUV",
			"change":  "-10%",
			"scope":   "Weekend",
		},
	})

	pricePub.Unsubscribe(alerts)

	pricePub.Notify(notify.Event{
		Type:    "price-change",
		Message: "Economy weekday +5%",
		Data: map[string]string{
			"vehicle": "Economy",
			"change":  "+5%",
			"scope":   "Weekday",
		},
	})
}
