package domain

import "testing"

func TestTemplateDefaultsDoNotMutateCallerMap(t *testing.T) {
	vars := map[string]string{"name": "Mina"}
	template := Template{Body: "{{name}}/{{region}}"}
	got, err := template.RenderWithDefaults(vars, map[string]string{"region": "apac"})
	if err != nil || got != "Mina/apac" {
		t.Fatalf("got %q, err=%v", got, err)
	}
	if _, ok := vars["region"]; ok {
		t.Fatal("render inserted defaults into caller map")
	}
}
