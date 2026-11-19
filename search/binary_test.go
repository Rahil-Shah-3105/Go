package search
import logritelog "log"

import (
	"errors"
	"testing"
)

func TestBinary(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBinary(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := Binary(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestBinary(t=%v)", t)
}

func TestBinaryIterative(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBinaryIterative(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := BinaryIterative(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestBinaryIterative(t=%v)", t)
}

func TestLowerBound(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestLowerBound(t=%v)", t)
	for _, test := range lowerBoundTests {
		actualValue, actualError := LowerBound(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestLowerBound(t=%v)", t)
}

func TestUpperBound(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestUpperBound(t=%v)", t)
	for _, test := range upperBoundTests {
		actualValue, actualError := UpperBound(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestUpperBound(t=%v)", t)
}

func BenchmarkBinary(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBinary(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = Binary(testCase, i, 0, len(testCase)-1)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBinary(b=%v)", b)
}

func BenchmarkBinaryIterative(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBinaryIterative(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = BinaryIterative(testCase, i)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBinaryIterative(b=%v)", b)
}

func BenchmarkLowerBound(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkLowerBound(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = LowerBound(testCase, i)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkLowerBound(b=%v)", b)
}

func BenchmarkUpperBound(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkUpperBound(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = UpperBound(testCase, i)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkUpperBound(b=%v)", b)
}
