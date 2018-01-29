package StringConcatNew

import (
	"bytes"
)

// ConcatOperator
func ConcatOperator(original *string, concat string) {
	// This could be written as 'return *original + concat' but wanted to confirm no special
	// compiler optimizations existed for concatenating a string to itself.
	*original = *original + concat
}

// SelfConcatOperator
func SelfConcatOperator(input string, n int) string {
	output := ""
	for i := 0; i < n; i++ {
		ConcatOperator(&output, input)
	}
	return output
}

// ConcatBuffer
func ConcatBuffer(original *bytes.Buffer, concat string) {
	original.WriteString(concat)
}

// SelfConcatBuffer
func SelfConcatBuffer(input string, n int) string {
	var output bytes.Buffer
	for i := 0; i < n; i++ {
		ConcatBuffer(&output, input)
	}
	return string(output.Bytes())
}
