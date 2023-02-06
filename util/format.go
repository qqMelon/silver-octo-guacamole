package util

import (
	"fmt"
	"math"
)

func HumanFormat(bytes int64) string {
	n := float64(bytes)

	for _, unit := range []string{"", "Ki", "Mi", "Gi"} {
		if math.Abs(n) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", n, unit)
		}
		n /= 1024.0
	}

	return fmt.Sprintf("%.1fTiB", n)
}
