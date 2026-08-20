package httpapi

import (
	"encoding/json"
	"errors"
	messageapp "example.com/maildelivery/internal/message/application"
	"example.com/maildelivery/internal/message/domain"
	"example.com/maildelivery/internal/platform/metrics"
	queueapp "example.com/maildelivery/internal/queue/application"
	queuedomain "example.com/maildelivery/internal/queue/domain"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Messages *messageapp.Service
	Queue    *queueapp.Service
	Metrics  *metrics.Registry
	Log      *slog.Logger
	HTTP     *http.Server
}

func New(addr string, m *messageapp.Service, q *queueapp.Service, reg *metrics.Registry, l *slog.Logger) *Server {
	s := &Server{Messages: m, Queue: q, Metrics: reg, Log: l}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/readyz", s.health)
	mux.HandleFunc("/metrics", reg.Handler)
	mux.HandleFunc("/v1/messages", s.messages)
	mux.HandleFunc("/v1/messages/", s.message)
	mux.HandleFunc("/v1/jobs", s.jobs)
	s.HTTP = &http.Server{Addr: addr, Handler: s.middleware(mux), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	return s
}
func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", requestID(r))
		if r.Method != "GET" && r.Method != "POST" {
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "method not allowed", 405)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func requestID(r *http.Request) string {
	if v := r.Header.Get("X-Request-ID"); v != "" && !strings.ContainsAny(v, "\r\n") {
		return v
	}
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func (s *Server) messages(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		writeJSON(w, 200, s.Messages.List(r.Context(), r.URL.Query().Get("tenant_id")))
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var in struct {
		TenantID       string            `json:"tenant_id"`
		From           string            `json:"from"`
		To             string            `json:"to"`
		ReplyTo        string            `json:"reply_to"`
		Subject        string            `json:"subject"`
		Text           string            `json:"text"`
		HTML           string            `json:"html"`
		IdempotencyKey string            `json:"idempotency_key"`
		Headers        map[string]string `json:"headers"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	m, err := s.Messages.Submit(r.Context(), domain.Message{TenantID: in.TenantID, From: in.From, To: in.To, ReplyTo: in.ReplyTo, Subject: in.Subject, Text: in.Text, HTML: in.HTML, IdempotencyKey: in.IdempotencyKey, Headers: in.Headers})
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	_ = s.Queue.Enqueue(r.Context(), queuedomain.Job{ID: "job-" + m.ID, MessageID: m.ID, Kind: "delivery", Status: "queued"})
	writeJSON(w, 202, m)
}
func (s *Server) message(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/messages/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	m, err := s.Messages.Get(r.Context(), id)
	if errors.Is(err, errors.New("record not found")) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, m)
}
func (s *Server) jobs(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.Queue.List(r.Context()))
}
