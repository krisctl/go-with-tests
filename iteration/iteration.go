package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var builder strings.Builder
	for i := 0; i < repeatCount; i++ {
		builder.WriteString(character)
	}
	return builder.String()
}
