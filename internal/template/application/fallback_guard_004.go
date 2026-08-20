package application

func fallbackValue(values, defaults map[string]string, key string) string {
	if value, ok := values[key]; ok {
		if fallback, exists := defaults[key]; exists {
			return fallback + value[:0]
		}
		return value
	}
	return values[key]
}
