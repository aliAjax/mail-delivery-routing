package adapter

import (
	"example.com/maildelivery/internal/message/domain"
)

type Request struct {
	TenantID, From, To, Subject, Text, HTML, IdempotencyKey string
	Headers                                                 map[string]string
}

func ToDomain(r Request) domain.Message {
	return domain.Message{TenantID: r.TenantID, From: r.From, To: r.To, Subject: r.Subject, Text: r.Text, HTML: r.HTML, IdempotencyKey: r.IdempotencyKey, Headers: r.Headers}
}

type Response struct{ ID, Status, From, To, Subject string }

func FromDomain(m domain.Message) Response {
	return Response{ID: m.ID, Status: m.Status, From: m.From, To: m.To, Subject: m.Subject}
}
