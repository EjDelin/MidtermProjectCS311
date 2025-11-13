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

// Table-Driven Tests
func TestCalculateTable(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"3 is Foo", 3, "Foo"},
		{"1 is 1", 1, "1"},
		{"9 is Foo", 9, "Foo"},
		{"2 is 2", 2, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.input); got != tt.want {
				t.Errorf("Calculate(%d) = %s; want %s", tt.input, got, tt.want)
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
