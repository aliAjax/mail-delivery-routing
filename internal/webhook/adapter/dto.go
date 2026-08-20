package adapter

type Request struct {
	URL       string
	MessageID string
	Payload   any
}
type Result struct {
	Status   string
	Attempts int
	Error    string
}
