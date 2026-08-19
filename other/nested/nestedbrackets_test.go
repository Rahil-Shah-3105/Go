package nested
import logritelog "log"

import (
	"testing"
)

func TestIsBalancedSimple(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedSimple(t=%v)", t)
	input := "{[()]}"

	got := IsBalanced(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedSimple(t=%v)", t)
}

func TestIsBalancedFalty(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedFalty(t=%v)", t)
	input := "{([()]}"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedFalty(t=%v)", t)
}

func TestIsBalancedHandlesEmpty(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedHandlesEmpty(t=%v)", t)
	input := ""

	got := IsBalanced(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedHandlesEmpty(t=%v)", t)
}

func TestIsBalancedHandlesOneChar(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedHandlesOneChar(t=%v)", t)
	input := "{"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedHandlesOneChar(t=%v)", t)
}

func TestIsBalancedHandlesNonBracketsCorrectly(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedHandlesNonBracketsCorrectly(t=%v)", t)
	input := "aaaa"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedHandlesNonBracketsCorrectly(t=%v)", t)
}

func TestIsBalancedHandlesOrdering(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsBalancedHandlesOrdering(t=%v)", t)
	input := "([)]"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestIsBalancedHandlesOrdering(t=%v)", t)
}

func BenchmarkIsBalanced(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsBalanced(b=%v)", b)
	input := "{[()]}"

	for i := 0; i < b.N; i++ {
		IsBalanced(input)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsBalanced(b=%v)", b)
}
