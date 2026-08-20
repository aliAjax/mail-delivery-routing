package metrics

type Snapshot struct{ Submitted, Delivered, Failed uint64 }

func (r *Registry) Snapshot() Snapshot {
	return Snapshot{Submitted: r.submitted.Load(), Delivered: r.delivered.Load(), Failed: r.failed.Load()}
}
func (s Snapshot) Healthy() bool { return s.Failed == 0 || s.Delivered >= s.Failed }
func (s Snapshot) AsMap() map[string]uint64 {
	return map[string]uint64{"submitted": s.Submitted, "delivered": s.Delivered, "failed": s.Failed}
}
func (s Snapshot) Total() uint64 { return s.Submitted }
