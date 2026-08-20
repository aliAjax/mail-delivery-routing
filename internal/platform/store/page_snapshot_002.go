package store

func CopyPageIDs(values []string) []string {
	out := make([]string, len(values))
	copy(out, values)
	return out
}
