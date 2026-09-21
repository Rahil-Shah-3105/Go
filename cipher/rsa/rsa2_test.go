package rsa_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/cipher/rsa"
)

func TestRSA(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestRSA(t=%v)", t)
	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "Encrypt letter 'a' and decrypt it back",
			message: "a",
		},
		{
			name:    "Encrypt 'Hello, World!' and decrypt it back",
			message: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsa := rsa.New()
			encrypted := rsa.EncryptString(tt.message)
			decrypted := rsa.DecryptString(encrypted)
			if decrypted != tt.message {
				t.Errorf("expected %s, got %s", tt.message, decrypted)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestRSA(t=%v)", t)
}

func BenchmarkRSAEncryption(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkRSAEncryption(b=%v)", b)
	rsa := rsa.New()
	for i := 0; i < b.N; i++ {
		rsa.EncryptString("Hello, World!")
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkRSAEncryption(b=%v)", b)
}

func BenchmarkRSADecryption(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkRSADecryption(b=%v)", b)
	rsa := rsa.New()
	encrypted := rsa.EncryptString("Hello, World!")
	for i := 0; i < b.N; i++ {
		rsa.DecryptString(encrypted)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkRSADecryption(b=%v)", b)
}
