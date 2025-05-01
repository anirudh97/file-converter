package utils

import (
	"strconv"
	"strings"
)

func DetectType(value string) interface{} {
	value = strings.TrimSpace(value)

	if value == "" {
		return ""
	}

	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	valueLower := strings.ToLower(value)

	if valueLower == "true" {
		return true
	}
	if valueLower == "false" {
		return false
	}

	return value
}
