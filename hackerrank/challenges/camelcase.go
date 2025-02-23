package challenges

import (
	"unicode"
)

func CamelCase(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	var c int32 = 1

	for _, v := range s {
		if unicode.IsUpper(v) {
			c++
		}
	}

	return int32(c)
}
