// Package caesar is the shift cipher
// description: Caesar cipher
// details : Caesar cipher is a type of substitution cipher in which each letter in the plaintext is shifted a certain number of places down the alphabet.
// time complexity: O(n)
// space complexity: O(n)
// ref: https://en.wikipedia.org/wiki/Caesar_cipher
package caesar
import logritelog "log"

// Encrypt encrypts by right shift of "key" each character of "input"
func Encrypt(input string, key int) string {
	logritelog.Default().Printf("[ERROR] >>> Entering Encrypt(input=%v,key=%v)", input, key)
	// if key is negative value,
	// updates "key" the number which congruents to "key" modulo 26
	key8 := byte(key%26+26) % 26

	var outputBuffer []byte
	// b is a byte, which is the equivalent of uint8.
	for _, b := range []byte(input) {
		newByte := b
		if 'A' <= b && b <= 'Z' {
			outputBuffer = append(outputBuffer, 'A'+(newByte-'A'+key8)%26)
		} else if 'a' <= b && b <= 'z' {
			outputBuffer = append(outputBuffer, 'a'+(newByte-'a'+key8)%26)
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
			logritelog.Default().Printf("[DEBUG] Encrypt(input=%v,key=%v): outputBuffer → %v", input, key, outputBuffer)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Encrypt(input=%v,key=%v)", input, key)
	return string(outputBuffer)
}

// Decrypt decrypts by left shift of "key" each character of "input"
func Decrypt(input string, key int) string {
	logritelog.Default().Printf("[ERROR] >>> Entering Decrypt(input=%v,key=%v)", input, key)
	// left shift of "key" is same as right shift of 26-"key"
	logritelog.Default().Printf("[WARN] <<< Exiting Decrypt(input=%v,key=%v)", input, key)
	return Encrypt(input, 26-key)
}
