package logy

import (
	"fmt"
	"strconv"
	"time"
)

const (
	red         = "\033[31m"
	green       = "\033[32m"
	yellow      = "\033[33m"
	blue        = "\033[34m"
	magenta     = "\033[35m"
	brightGreen = "\033[92m"
	brightBlue  = "\033[94m"
	cyan        = "\033[36m"
	reset       = "\033[0m"
)

func colorizeString(color string, input any) string {
	switch i := input.(type) {
	case string:
		return color + i + reset
	case int:
		return color + strconv.Itoa(i) + reset
	case time.Duration:
		return color + i.String() + reset
	default:
		return fmt.Sprintf("%v", i)
	}
}
