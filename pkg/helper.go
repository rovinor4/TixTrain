package pkg

import (
	"fmt"
	"strings"
)

type Helper struct {
}

// IntToRupiah : convert int to rupiah format
func (h *Helper) IntToRupiah(amount int) string {
	rupiah := "Rp " + fmt.Sprintf("%s", h.FormatNumber(amount))
	return rupiah
}

// FormatNumber : format number with dot as a thousand separator ||  example: 1000000 -> 1.000.000
func (h *Helper) FormatNumber(amount int) string {
	result := fmt.Sprintf("%d", amount)
	n := len(result)
	if n <= 3 {
		return result
	}

	var formatted string
	for i, digit := range result {
		if (n-i)%3 == 0 && i != 0 {
			formatted += "."
		}
		formatted += string(digit)
	}
	return formatted
}

// SensorString : 08123456789 -> 0812****6789
func (h *Helper) SensorString(number string) string {
	n := len(number)
	if n <= 6 {
		return number
	}

	sensorLength := n - 6
	sensor := ""
	for i := 0; i < sensorLength; i++ {
		sensor += "*"
	}

	return number[:4] + sensor + number[n-2:]
}

// TitleCase : convert string from "WONOKROMO" to "Wonokromo"
func (h *Helper) TitleCase(name string) string {
	if len(name) == 0 {
		return name
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(string(name[0])), strings.ToLower(name[1:]))
}
