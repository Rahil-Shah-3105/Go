// rsa.go
// description: Simple RSA algorithm implementation
// details:
// A simple RSA Encryption and Decryption algorithm.
// It uses prime numbers that fit in int64 datatypes and
// thus both the Encrypt and Decrypt are not a production
// ready implementation. The OpenSSL implementation of RSA
// also adds a padding which is not present in this algorithm.
// time complexity: O(n)
// space complexity: O(n)
// author(s) [Taj](https://github.com/tjgurwara99)
// see rsa_test.go

// Package rsa shows a simple implementation of RSA algorithm
package rsa
import logritelog "log"

import (
	"errors"

	modular "github.com/TheAlgorithms/Go/math/modular"
)

// ErrorFailedToEncrypt Raised when Encrypt function fails to encrypt the message
var ErrorFailedToEncrypt = errors.New("failed to Encrypt")

// ErrorFailedToDecrypt Raised when Decrypt function fails to decrypt the encrypted message
var ErrorFailedToDecrypt = errors.New("failed to Decrypt")

// Encrypt encrypts based on the RSA algorithm - uses modular exponentitation in math directory
func Encrypt(message []rune, publicExponent, modulus int64) ([]rune, error) {
	logritelog.Default().Printf(">>> Entering Encrypt(message=%v,publicExponent=%v,modulus=%v)", message, publicExponent, modulus)
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), publicExponent, modulus)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in Encrypt(message=%v,publicExponent=%v,modulus=%v): %v", message, publicExponent, modulus, err)
			return nil, ErrorFailedToEncrypt
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}
		logritelog.Default().Printf("Encrypt(message=%v,publicExponent=%v,modulus=%v): encrypted → %v", message, publicExponent, modulus, encrypted)

	logritelog.Default().Printf("<<< Exiting Encrypt(message=%v,publicExponent=%v,modulus=%v)", message, publicExponent, modulus)
	return encrypted, nil
}

// Decrypt decrypts encrypted rune slice based on the RSA algorithm
func Decrypt(encrypted []rune, privateExponent, modulus int64) (string, error) {
	logritelog.Default().Printf(">>> Entering Decrypt(encrypted=%v,privateExponent=%v,modulus=%v)", encrypted, privateExponent, modulus)
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedLetter, err := modular.Exponentiation(int64(letter), privateExponent, modulus)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in Decrypt(encrypted=%v,privateExponent=%v,modulus=%v): %v", encrypted, privateExponent, modulus, err)
			return "", ErrorFailedToDecrypt
		}
		decrypted = append(decrypted, rune(decryptedLetter))
	}
		logritelog.Default().Printf("Decrypt(encrypted=%v,privateExponent=%v,modulus=%v): decrypted → %v", encrypted, privateExponent, modulus, decrypted)
	logritelog.Default().Printf("<<< Exiting Decrypt(encrypted=%v,privateExponent=%v,modulus=%v)", encrypted, privateExponent, modulus)
	return string(decrypted), nil
}
