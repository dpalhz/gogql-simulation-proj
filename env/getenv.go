package env

import "os"

// GetEnv reads an environment variable and returns its value.
// If the environment variable is not set, it returns a specified default value.
// This function encapsulates the standard library's [os.LookupEnv] to provide defaults,
// following the common Philosophy Go idiom of "make the zero value useful".
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
