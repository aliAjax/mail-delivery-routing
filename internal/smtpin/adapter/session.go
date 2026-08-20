package adapter

import "strings"

func Command(line string) (verb, arg string) {
	p := strings.SplitN(strings.TrimSpace(line), " ", 2)
	verb = strings.ToUpper(p[0])
	if len(p) == 2 {
		arg = p[1]
	}
	return
}
func ValidMailbox(v string) bool { return strings.Count(strings.Trim(v, "<>"), "@") == 1 }
