package config

import "fmt"

func ReloadGeneration(r Runtime, err error) (Runtime, error) {
	if err != nil {
		return r, err
	}
	if r.Generation < 0 {
		return r, fmt.Errorf("invalid generation: %d", r.Generation)
	}
	return r, nil
}
