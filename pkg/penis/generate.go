package penis

import "strings"

func Generate(length int) string {
	var b strings.Builder

	b.WriteRune('8')

	for i := 0; i < length; i++ {
		b.WriteRune('=')
	}

	b.WriteRune('Э')

	return b.String()
}
