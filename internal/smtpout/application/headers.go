package application

import "strings"

func BuildHeaders(subject, messageID string) string {
	var b strings.Builder
	b.WriteString("Subject: ")
	b.WriteString(strings.ReplaceAll(subject, "\n", ""))
	b.WriteString("\r\n")
	if messageID != "" {
		b.WriteString("Message-ID: <")
		b.WriteString(messageID)
		b.WriteString(">\r\n")
	}
	return b.String()
}
func NormalizeDomain(address string) string {
	p := strings.Split(strings.Trim(address, "<>"), "@")
	if len(p) != 2 {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(p[1]))
}
