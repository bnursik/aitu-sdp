package notify

type ISubject interface {
	Subscribe(IObserver)
	Unsubscribe(IObserver)
	Notify(Event)
}
