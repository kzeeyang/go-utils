package utils

import (
	"fmt"
	"strconv"
)

// DecimalFixed2 - format float or string  with 2 decimal
func DecimalFixed2(value interface{}) float64 {
	data, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return data
}
