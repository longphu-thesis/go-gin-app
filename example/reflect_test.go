package example

import "testing"

// Test Function Sum
func TestRun(t *testing.T) {
	f := &Foo{
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       50,
	}

	f.reflect()
}