package env

import (
	"fmt"
	"os"
)

func Validate(variables []string) error {
	for _, variable := range variables {
		value, exists := os.LookupEnv(variable)
		if !exists || value == "" {
			return fmt.Errorf("environment variable %s is mandatory", variable)
		}
	}
	return nil
}
