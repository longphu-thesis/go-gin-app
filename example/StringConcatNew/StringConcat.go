package StringConcatNew

import (
	"bytes"
)

// Function Concat Operator
func ConcatOperator(original *string, concat string) {
	// This could be written as 'return *original + concat' but wanted to confirm no special
	// compiler optimizations existed for concatenating a string to itself.
	*original = *original + concat
}

// Function Self Concat Operator
func SelfConcatOperator(input string, n int) string {
	output := ""
	for i := 0; i < n; i++ {
		ConcatOperator(&output, input)
	}
	return output
}

// Function Concat Buffer
func ConcatBuffer(original *bytes.Buffer, concat string) {
	original.WriteString(concat)
}

// Function Self Concat Buffer
func SelfConcatBuffer(input string, n int) string {
	var output bytes.Buffer
	for i := 0; i < n; i++ {
		ConcatBuffer(&output, input)
	}
	return string(output.Bytes())
}
