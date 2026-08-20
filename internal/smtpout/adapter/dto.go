package adapter

type DeliveryRequest struct{ From, To, Subject, Body string }
type DeliveryResult struct {
	Accepted bool
	Code     int
	Reason   string
}

func Success() DeliveryResult { return DeliveryResult{Accepted: true, Code: 250, Reason: "accepted"} }
func Temporary(reason string) DeliveryResult {
	return DeliveryResult{Accepted: false, Code: 451, Reason: reason}
}
