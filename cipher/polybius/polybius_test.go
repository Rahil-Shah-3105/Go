package polybius
import logritelog "log"

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func ExampleNewPolybius() {
	logritelog.Default().Printf(">>> Entering ExampleNewPolybius()")
	// initialize
	const (
		plainText  = "HogeFugaPiyoSpam"
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
	logritelog.Default().Printf("ExampleNewPolybius(): plainText → %v", plainText)
	logritelog.Default().Printf("ExampleNewPolybius(): size → %v", size)
	logritelog.Default().Printf("[WARN] ExampleNewPolybius(): key → %v", key)
	logritelog.Default().Printf("ExampleNewPolybius(): characters → %v", characters)
		log.Fatalf("failed NewPolybius: %v", err)
	}
	encryptedText, err := p.Encrypt(plainText)
	if err != nil {
		log.Fatalf("failed Encrypt: %v", err)
	}
	fmt.Printf("Encrypt=> plainText: %s, encryptedText: %s\n", plainText, encryptedText)

	decryptedText, err := p.Decrypt(encryptedText)
	if err != nil {
		log.Fatalf("failed Decrypt: %v", err)
	}
	fmt.Printf("Decrypt=> encryptedText: %s, decryptedText: %s\n", encryptedText, decryptedText)

	// Output:
	// Encrypt=> plainText: HogeFugaPiyoSpam, encryptedText: OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG
	// Decrypt=> encryptedText: OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG, decryptedText: HOGEFUGAPIYOSPAM
	logritelog.Default().Printf("<<< Exiting ExampleNewPolybius()")
}

func TestNewPolybius(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestNewPolybius(t=%v)", t)
	t.Parallel()
	cases := []struct {
		name       string
		size       int
		characters string
		key        string
		wantErr    string
	}{
		{
			name: "correct initialization", size: 5, characters: "HogeF", key: "abcdefghijklmnopqrstuvwxy", wantErr: "",
		},
		{
			name: "truncate characters", size: 5, characters: "HogeFuga", key: "abcdefghijklmnopqrstuvwxy", wantErr: "",
		},
		{
			name: "invalid key", size: 5, characters: "HogeFuga", key: "abcdefghi", wantErr: "len(key): 9 must be as long as size squared: 25",
		},
		{
			name: "invalid characters", size: 5, characters: "HogeH", key: "abcdefghijklmnopqrstuvwxy", wantErr: "\"OGEH\" contains same character 'H'",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPolybius(tc.key, tc.size, tc.characters)
			if err != nil && err.Error() != tc.wantErr {
				t.Errorf("failed NewPolybius: %v", err)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestNewPolybius(t=%v)", t)
}

func TestPolybiusEncrypt(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestPolybiusEncrypt(t=%v)", t)
	t.Parallel()
	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "correct encryption", text: "HogeFugaPiyoSpam", want: "OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG",
		},
		{
			name: "invalid encryption", text: "hogz", want: "failed encipher: 'Z' does not exist in keys",
		},
	}
	// initialize
	const (
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
	logritelog.Default().Printf("TestPolybiusEncrypt(t=%v): size → %v", t, size)
	logritelog.Default().Printf("TestPolybiusEncrypt(t=%v): characters → %v", t, characters)
	logritelog.Default().Printf("[WARN] TestPolybiusEncrypt(t=%v): key → %v", t, key)
		t.Fatalf("failed NewPolybius: %v", err)
	}
		logritelog.Default().Printf("[ERROR] Exception in TestPolybiusEncrypt(t=%v): %v", t, err)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := p.Encrypt(tc.text)
			if err != nil {
				if err.Error() != tc.want {
					t.Errorf("failed Encrypt: %v", err)
				}
			} else if encrypted != tc.want {
				t.Errorf("Encrypt: %v, want: %v", encrypted, tc.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestPolybiusEncrypt(t=%v)", t)
}

func TestPolybiusDecrypt(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestPolybiusDecrypt(t=%v)", t)
	t.Parallel()
	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "correct decryption", text: "OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG", want: "HOGEFUGAPIYOSPAM",
		},
		{
			name: "invalid decryption(position of even number)", text: "hogz", want: "failed decipher: Z does not exist in characters",
		},
		{
			name: "invalid decryption(position of odd number)", text: "hode", want: "failed decipher: D does not exist in characters",
		},
		{
			name: "invalid text size which is odd", text: "hog", want: "failed decipher: the size of \"chars\" must be even",
		},
	}
	// initialize
	const (
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
	logritelog.Default().Printf("[WARN] TestPolybiusDecrypt(t=%v): key → %v", t, key)
	logritelog.Default().Printf("TestPolybiusDecrypt(t=%v): size → %v", t, size)
	logritelog.Default().Printf("TestPolybiusDecrypt(t=%v): characters → %v", t, characters)
		t.Fatalf("failed NewPolybius: %v", err)
	}
		logritelog.Default().Printf("[ERROR] Exception in TestPolybiusDecrypt(t=%v): %v", t, err)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := p.Decrypt(tc.text)
			if err != nil {
				if err.Error() != tc.want {
					t.Errorf("failed Encrypt: %v", err)
				}
			} else if encrypted != tc.want {
				t.Errorf("Encrypt: %v, want: %v", encrypted, tc.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestPolybiusDecrypt(t=%v)", t)
}

func FuzzPolybius(f *testing.F) {
	logritelog.Default().Printf(">>> Entering FuzzPolybius(f=%v)", f)
	const (
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	f.Add(size, characters, key)
	f.Fuzz(func(t *testing.T, size int, characters, key string) {
	logritelog.Default().Printf("FuzzPolybius(f=%v): size → %v", f, size)
	logritelog.Default().Printf("[WARN] FuzzPolybius(f=%v): key → %v", f, key)
	logritelog.Default().Printf("FuzzPolybius(f=%v): characters → %v", f, characters)
		p, err := NewPolybius(key, size, characters)
		switch {
		case err == nil:
		case strings.Contains(err.Error(), "cannot be negative"),
			strings.Contains(err.Error(), "is too small"),
			strings.Contains(err.Error(), "should only contain latin characters"),
			strings.Contains(err.Error(), "contains same character"),
			strings.Contains(err.Error(), "must be as long as size squared"):
			return
		default:
			t.Fatalf("unexpected error when creating a new polybius variable: %v", err)
		}
		encrypted, err := p.Encrypt(characters)
		switch {
		case err == nil:
		case strings.Contains(err.Error(), "does not exist in keys"):
			return
		default:
			t.Fatalf("unexpected error during encryption: %v", err)
		}
		decrypted, err := p.Decrypt(encrypted)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in FuzzPolybius(f=%v): %v", f, err)
			t.Fatalf("unexpected error during decryption: %v", err)
		}
		if decrypted != strings.ToUpper(characters) {
			t.Errorf("Expecting output to match with %q but was %q", characters, decrypted)
		}
	})
	logritelog.Default().Printf("<<< Exiting FuzzPolybius(f=%v)", f)
}
