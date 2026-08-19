package hamming
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

var testCases = []struct {
	name     string
	string1  string
	string2  string
	expected int
}{
	{
		"empty strings",
		"",
		"",
		0,
	},
	{
		"single character strings",
		"A",
		"A",
		0,
	},
	{
		"two different strings with a same length",
		"TestString 1",
		"TestString 2",
		1,
	},
	{
		"two different strings with a different length",
		"TestString1",
		"TestString",
		-1,
	},
	{
		"two same strings with a same length",
		"TestString",
		"TestString",
		0,
	},
}

func TestHammingDistance(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestHammingDistance(t=%v)", t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Distance(tc.string1, tc.string2)
			if err != nil {
				logritelog.Default().Printf("[ERROR] Exception in TestHammingDistance(t=%v): %v", t, err)
				if tc.expected != -1 {
					t.Fatalf("Expected no error, but got %v", err)
				}
			} else if actual != tc.expected {
				t.Errorf("Expected Hamming distance between strings: '%s' and '%s' is %v, but got: %v", tc.string1, tc.string2, tc.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestHammingDistance(t=%v)", t)
}
