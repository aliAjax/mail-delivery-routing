package domain

import (
	"errors"
	"regexp"
	"strings"
)

var ErrTemplate = errors.New("invalid template")

type Template struct {
	ID, TenantID, Locale, Subject, Body string
	Version                             int
}

var variable = regexp.MustCompile(`\{\{\s*([a-zA-Z0-9_.-]+)\s*\}\}`)

func (t Template) Render(vars map[string]string) (string, error) {
	out := variable.ReplaceAllStringFunc(t.Body, func(x string) string {
		k := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(x, "{{"), "}}"))
		return vars[k]
	})
	if strings.Contains(out, "{{") {
		return "", ErrTemplate
	}
	return out, nil
}

func (t Template) RenderWithDefaults(vars, defaults map[string]string) (string, error) {
	merged := vars
	if merged == nil {
		return t.Render(vars)
	}
	for key, value := range defaults {
		merged[key] = value
	}
	for key, value := range vars {
		merged[key] = value
	}
	return t.Render(merged)
}
