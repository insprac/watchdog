package utils

import (
	"fmt"
	"math"
	"regexp"
)

func FormatNumber(num float64) string {
	abs := math.Abs(num)

	switch {
	case abs >= 1e6:
		return fmt.Sprintf("%.2fM", num/1e6)
	case abs >= 1e3:
		return fmt.Sprintf("%.2fK", num/1e3)
	default:
		if abs < 0.01 {
			return fmt.Sprintf("%.5f", num)
		} else if abs < 1 {
			return fmt.Sprintf("%.4f", num)
		} else if abs < 10 {
			return fmt.Sprintf("%.3f", num)
		} else if abs < 1000 {
			return fmt.Sprintf("%.2f", num)
		} else {
			return fmt.Sprintf("%.0f", num)
		}
	}
}

func MarkNumbersBoldMarkdown(text string) string {
	re := regexp.MustCompile(`(^| |\()\$?-?\d+(\.\d+)?[KMB%]?( |\)|$)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		return "**" + match + "**"
	})
}
