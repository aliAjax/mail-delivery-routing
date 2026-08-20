package store

import (
	"context"
	"runtime/debug"
	"testing"
)

func TestListPageReturnsIndependentRecords(t *testing.T) {
	m := NewMemory()
	for i := 0; i < 4; i++ {
		id := string(rune('a' + i))
		if err := m.Put(context.Background(), Record{ID: id, TenantID: "tenant", Subject: "original"}); err != nil {
			t.Fatal(err)
		}
	}
	page := m.ListPage(context.Background(), "tenant", Page{Limit: 20})
	if len(page) != 4 {
		debug.PrintStack()
		t.Fatalf("got %d records", len(page))
	}
	page[0].Subject = "mutated outside store"
	again := m.ListPage(context.Background(), "tenant", Page{Limit: 20})
	for _, record := range again {
		if record.Subject != "original" {
			t.Fatalf("store record changed through page alias: %+v", record)
		}
	}
}
