package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatNumber(number int) string {

	s := strconv.Itoa(number)
	r1 := ""
	idx := 0

	for i := len(s) - 1; i >= 0; i-- {
		idx++
		if idx == 4 {
			idx = 1
			r1 = r1 + "."
		}
		r1 = r1 + string(s[i])
	}

	r2 := ""
	for i := len(r1) - 1; i >= 0; i-- {
		r2 = r2 + string(r1[i])
	}

	return r2

}

func FormatNumbering(number int, maxNumber int) string {

	overallLen := countDigits(maxNumber)
	var padCountInt = overallLen - countDigits(number)
	var retStr = fmt.Sprintf("%s%d", strings.Repeat("0", padCountInt), number)
	return retStr
}

func countDigits(number int) int {
	var digits int
	for number != 0 {
		number /= 10
		digits = digits + 1
	}
	return digits
}
