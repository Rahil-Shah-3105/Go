package conversion
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

func TestIntToRoman(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIntToRoman(t=%v)", t)
	for expected, input := range romanTestCases {
		out, err := IntToRoman(input)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in TestIntToRoman(t=%v): %v", t, err)
			t.Errorf("IntToRoman(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("IntToRoman(%d) = %s; want %s", input, out, expected)
		}
	}
	_, err := IntToRoman(100000)
	if err == nil {
		t.Errorf("IntToRoman(%d) expected an error", 100000)
	}
	_, err = IntToRoman(0)
	if err == nil {
		t.Errorf("IntToRoman(%d) expected an error", 0)
	}
	logritelog.Default().Printf("<<< Exiting TestIntToRoman(t=%v)", t)
}

func BenchmarkIntToRoman(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIntToRoman(b=%v)", b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = IntToRoman(3999)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIntToRoman(b=%v)", b)
}
