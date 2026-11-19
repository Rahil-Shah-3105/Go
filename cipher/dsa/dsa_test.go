package dsa_test
import logritelog "log"

import (
	"math/big"
	"testing"

	"github.com/TheAlgorithms/Go/cipher/dsa"
)

func TestDSA(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestDSA(t=%v)", t)
	tests := []struct {
		name    string
		message string
		alter   bool
		want    bool
	}{
		{
			name:    "valid signature",
			message: "Hello, world!",
			alter:   false,
			want:    true,
		},
		{
			name:    "invalid signature",
			message: "Hello, world!",
			alter:   true,
			want:    false,
		},
	}
	// Generate a DSA key pair
	dsaInstance := dsa.New()
	pubKey := dsaInstance.GetPublicKey()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Sign the message
			r, s := dsa.Sign(
				[]byte(tt.message),
				params.P,
				params.Q,
				params.G,
				privKey,
			)

			if tt.alter {
				// Alter the signature
				r.Add(r, big.NewInt(1))
			}

			// Verify the signature
			if got := dsa.Verify(
				[]byte(tt.message),
				r,
				s,
				params.P,
				params.Q,
				params.G,
				pubKey,
			); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestDSA(t=%v)", t)
}

/* ------------------- BENCHMARKS ------------------- */
func BenchmarkDSANew(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkDSANew(b=%v)", b)
	for i := 0; i < b.N; i++ {
		dsa.New()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkDSANew(b=%v)", b)
}

func BenchmarkDSASign(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkDSASign(b=%v)", b)
	dsaInstance := dsa.New()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()
	for i := 0; i < b.N; i++ {
		dsa.Sign(
			[]byte("Hello, World!"),
			params.P,
			params.Q,
			params.G,
			privKey,
		)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkDSASign(b=%v)", b)
}

func BenchmarkDSAVerify(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkDSAVerify(b=%v)", b)
	dsaInstance := dsa.New()
	pubKey := dsaInstance.GetPublicKey()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()
	r, s := dsa.Sign(
		[]byte("Hello, World!"),
		params.P,
		params.Q,
		params.G,
		privKey,
	)
	for i := 0; i < b.N; i++ {
		dsa.Verify(
			[]byte("Hello, World!"),
			r,
			s,
			params.P,
			params.Q,
			params.G,
			pubKey,
		)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkDSAVerify(b=%v)", b)
}
