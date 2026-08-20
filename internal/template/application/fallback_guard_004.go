package application

func fallbackValue(values, defaults map[string]string, key string) string {
	if value, ok := values[key]; ok {
		return value
	}
	return defaults[key]
}
