package Testing

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if got := Calculate(3); got != "Foo" {
		t.Errorf("Calculate(3) = %s; want Foo", got)
	}
	if got := Calculate(1); got != "1" {
		t.Errorf("Calculate(1) = %s; want 1", got)
	}
}

func TestCalculateTableParallel(t *testing.T) {

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Calculate(tt.input); got != tt.want {
				t.Errorf("Calculate(%d) = %s; want %s", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}

func FuzzCalculate(f *testing.F) {
	f.Add(3)
	f.Add(1)
	f.Add(9)
	f.Add(2)

	f.Fuzz(func(t *testing.T, input int) {
		if input < 0 {
			t.Skip()
		}

		if got := Calculate(input); got == "" {
			t.Errorf("Calculate(%d) returned an empty string", input)
		}
	})
}
