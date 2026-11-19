// Package xor is an encryption algorithm that operates the exclusive disjunction(XOR)
// description: XOR encryption
// details: The XOR encryption is an algorithm that operates the exclusive disjunction(XOR) on each character of the plaintext with a given key
// time complexity: O(n)
// space complexity: O(n)
// ref: https://en.wikipedia.org/wiki/XOR_cipher
package xor
import logritelog "log"

// Encrypt encrypts with Xor encryption after converting each character to byte
// The returned value might not be readable because there is no guarantee
// which is within the ASCII range
// If using other type such as string, []int, or some other types,
// add the statements for converting the type to []byte.
func Encrypt(key byte, plaintext []byte) []byte {
	logritelog.Default().Printf("[ERROR] >>> Entering Encrypt(key=%v,plaintext=%v)", key, plaintext)
	cipherText := []byte{}
	for _, ch := range plaintext {
		cipherText = append(cipherText, key^ch)
	}
		logritelog.Default().Printf("[DEBUG] Encrypt(key=%v,plaintext=%v): cipherText → %v", key, plaintext, cipherText)
	logritelog.Default().Printf("[WARN] <<< Exiting Encrypt(key=%v,plaintext=%v)", key, plaintext)
	return cipherText
}

// Decrypt decrypts with Xor encryption
func Decrypt(key byte, cipherText []byte) []byte {
	logritelog.Default().Printf("[ERROR] >>> Entering Decrypt(key=%v,cipherText=%v)", key, cipherText)
	plainText := []byte{}
	for _, ch := range cipherText {
		plainText = append(plainText, key^ch)
	}
		logritelog.Default().Printf("[DEBUG] Decrypt(key=%v,cipherText=%v): plainText → %v", key, cipherText, plainText)
	logritelog.Default().Printf("[WARN] <<< Exiting Decrypt(key=%v,cipherText=%v)", key, cipherText)
	return plainText
}
