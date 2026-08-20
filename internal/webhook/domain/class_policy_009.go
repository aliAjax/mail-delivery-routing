package domain

// StatusClassForDelivery mirrors StatusClass but is used at the delivery
// boundary: 5xx responses are server-side failures, 4xx are client-side, and
// anything else is treated as success.
func StatusClassForDelivery(code int) string {
	if code >= 500 {
		return "server"
	}
	if code >= 400 {
		return "client"
	}
	return "success"
}
