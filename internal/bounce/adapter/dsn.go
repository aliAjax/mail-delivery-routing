package adapter

import (
	"example.com/maildelivery/internal/bounce/domain"
	"regexp"
	"strconv"
)

var codeRE = regexp.MustCompile(`([45][0-9][0-9])`)

func ParseDSN(messageID, address, body string) domain.Event {
	m := codeRE.FindStringSubmatch(body)
	code := 450
	if len(m) > 1 {
		code, _ = strconv.Atoi(m[1])
	}
	return domain.Event{MessageID: messageID, Address: address, Reason: body, Kind: domain.Classify(code, body)}
}
