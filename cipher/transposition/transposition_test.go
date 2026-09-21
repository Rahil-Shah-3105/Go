// transposition_test.go
// description: Transposition cipher
// author(s) [red_byte](https://github.com/i-redbyte)
// see transposition.go

package transposition
import logritelog "log"

import (
	"errors"
	"math/rand"
	"reflect"
	"testing"
)

const enAlphabet = "abcdefghijklmnopqrstuvwxyz"

var texts = []string{
	"Ilya Sokolov",
	"A slice literal is declared just like an array literal, except you leave out the element count",
	"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
	"Go’s treatment of errors as values has served us well over the last decade. Although the standard library’s support for errors has been minimal—just the errors.New and fmt.Errorf functions, which produce errors that contain only a message—the built-in error interface allows Go programmers to add whatever information they desire. All it requires is a type that implements an Error method:",
	"А тут для примера русский текст",
}

func getRandomString() string {
	logritelog.Default().Printf("[ERROR] >>> Entering getRandomString()")
	enRunes := []rune(enAlphabet)
	b := make([]rune, rand.Intn(100))
	for i := range b {
		b[i] = enRunes[rand.Intn(len(enRunes))]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting getRandomString()")
	return string(b)
}

func TestEncrypt(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestEncrypt(t=%v)", t)
	fn := func(text string, keyWord string) (bool, error) {
		encrypt, err := Encrypt([]rune(text), keyWord)
		if err != nil && !errors.Is(err, ErrNoTextToEncrypt) && !errors.Is(err, ErrKeyMissing) {
			t.Error("Unexpected error ", err)
		}
		return text == string(encrypt), err
	}
	for _, s := range texts {
		if check, err := fn(s, getRandomString()); check || err != nil {
			t.Error("String ", s, " not encrypted")
		}
	}
	if _, err := fn(getRandomString(), ""); err == nil {
		t.Error("Error! empty string encryption")
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestEncrypt(t=%v)", t)
}

func TestDecrypt(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestDecrypt(t=%v)", t)
	for _, s := range texts {
		keyWord := getRandomString()
		encrypt, errEncrypt := Encrypt([]rune(s), keyWord)
		if errEncrypt != nil &&
			!errors.Is(errEncrypt, ErrNoTextToEncrypt) &&
			!errors.Is(errEncrypt, ErrKeyMissing) {
			t.Error("Unexpected error ", errEncrypt)
		}
		if errEncrypt != nil {
			t.Error(errEncrypt)
		}
		decrypt, errDecrypt := Decrypt([]rune(encrypt), keyWord)
		if errDecrypt != nil &&
			!errors.Is(errDecrypt, ErrNoTextToEncrypt) &&
			!errors.Is(errDecrypt, ErrKeyMissing) {
			t.Error("Unexpected error ", errDecrypt)
		}
		if errDecrypt != nil {
			t.Error(errDecrypt)
		}
		if reflect.DeepEqual(encrypt, decrypt) {
			t.Error("String ", s, " not encrypted")
		}
		if reflect.DeepEqual(encrypt, s) {
			t.Error("String ", s, " not encrypted")
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestDecrypt(t=%v)", t)
}

func TestEncryptDecrypt(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestEncryptDecrypt(t=%v)", t)
	text := []rune("Test text for checking the algorithm")
	key1 := "testKey"
	key2 := "Test Key2"
	encrypt, errEncrypt := Encrypt(text, key1)
	if errEncrypt != nil {
		t.Error(errEncrypt)
	}
	decrypt, errDecrypt := Decrypt(encrypt, key1)
	if errDecrypt != nil {
		t.Error(errDecrypt)
	}
	if !reflect.DeepEqual(decrypt, text) {
		t.Errorf("The string was not decrypted correctly %q %q", decrypt, text)
	}
	decrypt, _ = Decrypt([]rune(encrypt), key2)
	if reflect.DeepEqual(decrypt, text) {
		t.Errorf("The string was decrypted with a different key: %q %q", decrypt, text)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestEncryptDecrypt(t=%v)", t)
}

func FuzzTransposition(f *testing.F) {
	logritelog.Default().Printf("[ERROR] >>> Entering FuzzTransposition(f=%v)", f)
	for _, transpositionTestInput := range texts {
		f.Add(transpositionTestInput)
	}
	f.Fuzz(func(t *testing.T, input string) {
		keyword := getRandomString()
		message := []rune(input)
		encrypted, err := Encrypt(message, keyword)
		switch {
		case err == nil:
		case errors.Is(err, ErrKeyMissing),
			errors.Is(err, ErrNoTextToEncrypt):
			return
		default:
			t.Fatalf("unexpected error when encrypting string %q: %v", input, err)
		}
		decrypted, err := Decrypt([]rune(encrypted), keyword)
		switch {
		case err == nil:
		case errors.Is(err, ErrKeyMissing),
			errors.Is(err, ErrNoTextToEncrypt):
			return
		default:
			t.Fatalf("unexpected error when decrypting string %q: %v", encrypted, err)
		}

		if !reflect.DeepEqual(message, decrypted) {
			t.Fatalf("expected: %+v, got: %+v", message, []rune(decrypted))
		}
	})
	logritelog.Default().Printf("[WARN] <<< Exiting FuzzTransposition(f=%v)", f)
}
