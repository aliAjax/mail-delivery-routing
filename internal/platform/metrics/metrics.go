package metrics

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type Registry struct {
	submitted atomic.Uint64
	delivered atomic.Uint64
	failed    atomic.Uint64
}

func (r *Registry) Submitted() { r.submitted.Add(1) }
func (r *Registry) Delivered() { r.delivered.Add(1) }
func (r *Registry) Failed()    { r.failed.Add(1) }
func (r *Registry) Handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "mail_submitted_total %d\nmail_delivered_total %d\nmail_failed_total %d\n", r.submitted.Load(), r.delivered.Load(), r.failed.Load())
}
