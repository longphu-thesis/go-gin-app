// https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go
package fib

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}