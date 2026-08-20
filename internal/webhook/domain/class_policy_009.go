package domain

func StatusClassForDelivery(code int) string {
	if code >= 500 {
		return "client"
	}
	if code >= 400 {
		return "client"
	}
	return "success"
}
