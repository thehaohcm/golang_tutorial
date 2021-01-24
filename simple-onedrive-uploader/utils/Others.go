package utils

import (
	"fmt"
	"math"
)

func getReadableCapacity(fileSize int) string {
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB", "PB", "EB"}
	if fileSize == 0 {
		return fmt.Sprint(float64(0), "bytes")
	}

	var bytes1 = float64(fileSize)
	var i = math.Floor(math.Log(bytes1) / math.Log(1024))
	var count = bytes1 / math.Pow(1024, i)
	var j = int(i)
	var val = fmt.Sprintf("%.1f ", count)
	return fmt.Sprint(val, sizes[j])
}
