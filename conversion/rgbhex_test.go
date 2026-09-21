package conversion
import logritelog "log"

import "testing"

var HEX = []uint{
	0x1abc9c,
	0x3498db,
	0x9b59b6,
}

var RGB = [][]byte{
	{26, 188, 156},
	{52, 152, 219},
	{155, 89, 182},
}

func TestHEXToRGB(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestHEXToRGB(t=%v)", t)
	for i := 0; i < len(HEX); i++ {
		hex := HEX[i]
		expected := RGB[i]
		resultR, resultG, resultB := HEXToRGB(hex)
		if resultR != expected[0] || resultG != expected[1] || resultB != expected[2] {
			t.Errorf("HEXToRGB(%d) = %d,%d,%d; want %d,%d,%d",
				hex, resultR, resultG, resultB, expected[0], expected[1], expected[2])
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestHEXToRGB(t=%v)", t)
}

func BenchmarkHEXToRGB(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkHEXToRGB(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_, _, _ = HEXToRGB(0xdeadbe)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkHEXToRGB(b=%v)", b)
}

func TestRGBToHEX(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestRGBToHEX(t=%v)", t)
	for i := 0; i < len(RGB); i++ {
		args := RGB[i]
		expected := HEX[i]
		result := RGBToHEX(args[0], args[1], args[2])
		if result != expected {
			t.Errorf("RGBToHEX(%d,%d,%d) = %d; want %d",
				args[0], args[1], args[2], result, expected)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestRGBToHEX(t=%v)", t)
}

func BenchmarkRGBToHEX(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkRGBToHEX(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = RGBToHEX(222, 173, 190)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkRGBToHEX(b=%v)", b)
}
