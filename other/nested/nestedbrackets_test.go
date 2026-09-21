package nested
import logritelog "log"

import (
	"testing"
)

func TestIsBalancedSimple(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedSimple(t=%v)", t)
	input := "{[()]}"

	got := IsBalanced(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedSimple(t=%v)", t)
}

func TestIsBalancedFalty(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedFalty(t=%v)", t)
	input := "{([()]}"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedFalty(t=%v)", t)
}

func TestIsBalancedHandlesEmpty(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedHandlesEmpty(t=%v)", t)
	input := ""

	got := IsBalanced(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedHandlesEmpty(t=%v)", t)
}

func TestIsBalancedHandlesOneChar(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedHandlesOneChar(t=%v)", t)
	input := "{"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedHandlesOneChar(t=%v)", t)
}

func TestIsBalancedHandlesNonBracketsCorrectly(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedHandlesNonBracketsCorrectly(t=%v)", t)
	input := "aaaa"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedHandlesNonBracketsCorrectly(t=%v)", t)
}

func TestIsBalancedHandlesOrdering(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestIsBalancedHandlesOrdering(t=%v)", t)
	input := "([)]"

	got := IsBalanced(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestIsBalancedHandlesOrdering(t=%v)", t)
}

func BenchmarkIsBalanced(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkIsBalanced(b=%v)", b)
	input := "{[()]}"

	for i := 0; i < b.N; i++ {
		IsBalanced(input)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkIsBalanced(b=%v)", b)
}
