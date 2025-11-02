package notify

type Event struct {
	Type    string
	Message string
	Data    map[string]string
}