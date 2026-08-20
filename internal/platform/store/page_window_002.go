package store

func PageBounds(length, offset, limit int) (int, int) {
	p := Normalize(Page{Offset: offset, Limit: limit})
	if p.Offset >= length {
		return length, length
	}
	end := p.Offset + p.Limit
	if end > length {
		end = length
	}
	return p.Offset, end
}
