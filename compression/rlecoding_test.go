package compression_test
import logritelog "log"

import (
	"bytes"
	"testing"

	"github.com/TheAlgorithms/Go/compression"
)

func TestCompressionRLEncode(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCompressionRLEncode(t=%v)", t)
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "test 1",
			data: "WWWWWWWWWWWWBWWWWWWWWWWWWBBB",
			want: "12W1B12W3B",
		},
		{
			name: "test 2",
			data: "AABCCCDEEEE",
			want: "2A1B3C1D4E",
		},
		{
			name: "test 3",
			data: "AAAABBBCCDA",
			want: "4A3B2C1D1A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEncode(tt.data); got != tt.want {
				t.Errorf("RLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCompressionRLEncode(t=%v)", t)
}

func TestCompressionRLEDecode(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCompressionRLEDecode(t=%v)", t)
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "test 1",
			data: "12W1B12W3B",
			want: "WWWWWWWWWWWWBWWWWWWWWWWWWBBB",
		},
		{
			name: "test 2",
			data: "2A1B3C1D4E",
			want: "AABCCCDEEEE",
		},
		{
			name: "test 3",
			data: "4A3B2C1D1A",
			want: "AAAABBBCCDA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEdecode(tt.data); got != tt.want {
				t.Errorf("RLEdecode() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCompressionRLEDecode(t=%v)", t)
}

func TestCompressionRLEncodeBytes(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCompressionRLEncodeBytes(t=%v)", t)
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			name: "test 1",
			data: []byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"),
			want: []byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'},
		},
		{
			name: "test 2",
			data: []byte("AABCCCDEEEE"),
			want: []byte{2, 'A', 1, 'B', 3, 'C', 1, 'D', 4, 'E'},
		},
		{
			name: "test 3",
			data: []byte("AAAABBBCCDA"),
			want: []byte{4, 'A', 3, 'B', 2, 'C', 1, 'D', 1, 'A'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEncodebytes(tt.data); !bytes.Equal(got, tt.want) {
				t.Errorf("RLEncodebytes() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCompressionRLEncodeBytes(t=%v)", t)
}

func TestCompressionRLEDecodeBytes(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCompressionRLEDecodeBytes(t=%v)", t)
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			name: "test 1",
			data: []byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'},
			want: []byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"),
		},
		{
			name: "test 2",
			data: []byte{2, 'A', 1, 'B', 3, 'C', 1, 'D', 4, 'E'},
			want: []byte("AABCCCDEEEE"),
		},
		{
			name: "test 3",
			data: []byte{4, 'A', 3, 'B', 2, 'C', 1, 'D', 1, 'A'},
			want: []byte("AAAABBBCCDA"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEdecodebytes(tt.data); !bytes.Equal(got, tt.want) {
				t.Errorf("RLEdecodebytes() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCompressionRLEDecodeBytes(t=%v)", t)
}

/* --- BENCHMARKS --- */
func BenchmarkRLEncode(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRLEncode(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = compression.RLEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBB")
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRLEncode(b=%v)", b)
}

func BenchmarkRLEDecode(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRLEDecode(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = compression.RLEdecode("12W1B12W3B")
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRLEDecode(b=%v)", b)
}

func BenchmarkRLEncodeBytes(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRLEncodeBytes(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = compression.RLEncodebytes([]byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"))
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRLEncodeBytes(b=%v)", b)
}

func BenchmarkRLEDecodeBytes(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRLEDecodeBytes(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = compression.RLEdecodebytes([]byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'})
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRLEDecodeBytes(b=%v)", b)
}
