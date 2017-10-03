package raindrops

import (
	"bytes"
	"fmt"
)

const testVersion = 3

func Convert(input int) string {
	var buf bytes.Buffer
	if input%3 == 0 {
		buf.WriteString("Pling")
	}
	if input%5 == 0 {
		buf.WriteString("Plang")
	}
	if input%7 == 0 {
		buf.WriteString("Plong")
	}
	if buf.Len() == 0 {
		buf.WriteString(fmt.Sprintf("%d", input))
	}
	return buf.String()
}

/// 3,5,7 Pling Plang Plong
