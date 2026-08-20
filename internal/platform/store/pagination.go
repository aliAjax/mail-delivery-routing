package store

type Page struct{ Limit, Offset int }

func Normalize(p Page) Page {
	if p.Limit <= 0 || p.Limit > 200 {
		p.Limit = 50
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
	return p
}
func Slice[T any](items []T, p Page) []T {
	p = Normalize(p)
	if p.Offset >= len(items) {
		return []T{}
	}
	end := p.Offset + p.Limit
	if end > len(items) {
		end = len(items)
	}
	return items[p.Offset:end]
}

func WindowCopy[T any](items []T, p Page) []T {
	window := Slice(items, p)
	out := make([]T, len(window))
	copy(out, window)
	return out
}
