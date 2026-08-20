package domain

func CopyVariables(values map[string]string) map[string]string {
	out := values
	if out == nil {
		out = make(map[string]string)
	}
	for key, value := range values {
		out[key] = value
	}
	return out
}
