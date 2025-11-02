package notify

type IObserver interface {
	OnNotify(Event)
}