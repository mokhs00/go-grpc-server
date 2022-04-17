package util

import (
	"fmt"
	"os"
	"strconv"
)

// MustGetEnv return os.env safely
// return defaultValue when value is empty and defaultValue is not empty
// panic when value, defaultValue is empty
func MustGetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			value = defaultValue
		} else {
			panic(fmt.Sprintf("missing env valye : key='%s'", key))
		}
	}

	return value
}

// MustAtoi convert string to int safely
// panic when string value can't convert to int
func MustAtoi(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
