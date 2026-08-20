package config

import "fmt"

func WrapConfigError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("config reload: %w", err)
}
