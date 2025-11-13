package Testing

import (
	"testing"
)

// Unit Test
func TestCalculate(t *testing.T) {
	if got := Calculate(3); got != "Foo" {
		t.Errorf("Calculate(3) = %s; want Foo", got)
	}
	if got := Calculate(1); got != "1" {
		t.Errorf("Calculate(1) = %s; want 1", got)
	}
}

// Table-Driven Tests with Parallel Test
func TestCalculateTable(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Multiples of 3 (Foo)
		{"3 is Foo", 3, "Foo"},
		{"9 is Foo", 9, "Foo"},
		{"33 is Foo", 33, "Foo"},

		// Multiples of 5 (Bar) - Assumed Logic
		{"5 is Bar", 5, "Bar"},
		{"10 is Bar", 10, "Bar"},
		{"50 is Bar", 50, "Bar"},

		// Multiples of 15 (FooBar) - Assumed Logic
		{"15 is FooBar", 15, "FooBar"},
		{"30 is FooBar", 30, "FooBar"},
		{"150 is FooBar", 150, "FooBar"},

		// Non-Multiples (Number itself)
		{"1 is 1", 1, "1"},
		{"2 is 2", 2, "2"},
		{"4 is 4", 4, "4"},
		{"11 is 11", 11, "11"},

		// Edge Cases
		{"0 is 0", 0, "0"},
	}

	for _, tt := range tests {
		// **CRITICAL for Parallelism:** Capture the loop variable locally
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			// **Parallel Execution:** Run this subtest concurrently with others
			t.Parallel()

			if got := Calculate(tt.input); got != tt.want {
				t.Errorf("Calculate(%d) got %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Benchmark Test
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}

// Fuzz Test
func FuzzCalculate(f *testing.F) {
	// Seed corpus: starting inputs that the fuzzer will expand upon.
	f.Add(1)
	f.Add(3)
	f.Add(15)
	f.Add(100)

	// The fuzzer continuously generates random 'n' values of type int.
	f.Fuzz(func(t *testing.T, n int) {
		// Run the function with the generated input
		result := Calculate(n)

		// Add invariants (properties that must always be true)

		// Invariant 1: Result must not be an empty string unless input is unexpected (e.g., negative).
		if n >= 0 && result == "" {
			t.Errorf("Calculate(%d) returned empty string", n)
		}

		// Invariant 2: If the result is "FooBar", the number must be a multiple of 15.
		// This is a powerful check to ensure logic holds up for random inputs.
		if result == "FooBar" && n%15 != 0 {
			t.Errorf("Calculate(%d) returned 'FooBar', but is not a multiple of 15", n)
		}

		// Note: The main goal of fuzzing is finding panics, but invariants improve bug detection.
	})
}
