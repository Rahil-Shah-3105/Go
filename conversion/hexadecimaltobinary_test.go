package conversion
import logritelog "log"

import (
	"testing"
)

func TestHexToBinary(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestHexToBinary(t=%v)", t)
	tests := []struct {
		hex     string
		want    string
		wantErr bool
	}{
		{"", "", true},
		{"G123", "", true},
		{"12XZ", "", true},
		{"1", "1", false},
		{"A", "1010", false},
		{"10", "10000", false},
		{"1A", "11010", false},
		{"aB", "10101011", false},
		{"0Ff", "11111111", false},
		{"  1A ", "11010", false},
		{"0001A", "11010", false},
		{"7FFFFFFFFFFFFFFF", "111111111111111111111111111111111111111111111111111111111111111", false},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			got, err := hexToBinary(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToBinary(%q) error = %v, wantErr %v", tt.hex, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hexToBinary(%q) = %v, want %v", tt.hex, got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestHexToBinary(t=%v)", t)
}

func BenchmarkHexToBinary(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkHexToBinary(b=%v)", b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = hexToBinary("7FFFFFFFFFFFFFFF")
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkHexToBinary(b=%v)", b)
}
