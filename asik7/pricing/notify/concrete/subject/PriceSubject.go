package subject

import "asik7/pricing/notify"

type PriceSubject struct {
	subscribers []notify.IObserver
}

func NewPriceSubject() *PriceSubject {
	return &PriceSubject{subscribers: make([]notify.IObserver, 0, 4)}
}

func (s *PriceSubject) Subscribe(o notify.IObserver) {
	s.subscribers = append(s.subscribers, o)
}

func (s *PriceSubject) Unsubscribe(o notify.IObserver) {
	next := s.subscribers[:0]
	for _, sub := range s.subscribers {
		if sub != o {
			next = append(next, sub)
		}
	}
	s.subscribers = next
}

func (s *PriceSubject) Notify(e notify.Event) {
	for _, sub := range s.subscribers {
		sub.OnNotify(e)
	}
}
