package config

import (
	"fmt"
	"os"
)

type EnvMap map[string]func(string)

func findEnv(keys []string) (string, bool) {
	for _, key := range keys {
		if value, ok := os.LookupEnv(key); ok {
			return value, true
		}
	}

	return "", false
}

func loadProcessEnv(prefix string, envMap EnvMap) {
	for envKey, mapper := range envMap {
		if value, ok := findEnv([]string{
			fmt.Sprintf("%s_%s", prefix, envKey),
		}); ok {
			mapper(value)
		}
	}
}
