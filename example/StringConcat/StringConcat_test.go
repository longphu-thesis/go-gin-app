package StringConcat

import "testing"

// Test Benchmark
func BenchmarkSelfConcatOperator1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelfConcatOperator("test", 1000)
	}
}

// Test Function ConcatOperator
func TestConcatOperator(t *testing.T) {
	output := ""

	ConcatOperator(&output, "test")

	if output != "test" {
		t.Errorf("Concat was incorrect, got: %s, want: %s.", output, "test")
	}
}

// Test Function SelfConcatOperator
func TestSelfConcatOperator(t *testing.T) {

	output := SelfConcatOperator("test", 1)

	if output != "test" {
		t.Errorf("Concat was incorrect, got: %s, want: %s.", output, "test")
	}
}
