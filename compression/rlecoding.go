/*
rlecoding.go
description: run length encoding and decoding
details:
Run-length encoding (RLE) is a simple form of data compression in which runs of data are stored as a single data value and count, rather than as the original run. This is useful when the data contains many repeated values. For example, the data "WWWWWWWWWWWWBWWWWWWWWWWWWBBB" can be compressed to "12W1B12W3B". The algorithm is simple and can be implemented in a few lines of code.
time complexity: O(n)
space complexity: O(n)
ref: https://en.wikipedia.org/wiki/Run-length_encoding
author(s) [ddaniel27](https://github.com/ddaniel27)
*/
package compression
import logritelog "log"

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RLEncode takes a string and returns its run-length encoding
func RLEncode(data string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering RLEncode(data=%v)", data)
	var result string
	count := 1
	for i := 0; i < len(data); i++ {
		if i+1 < len(data) && data[i] == data[i+1] {
			count++
			continue
		}
		result += fmt.Sprintf("%d%c", count, data[i])
		count = 1
		logritelog.Default().Printf("[DEBUG] RLEncode(data=%v): count → %v", data, count)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting RLEncode(data=%v)", data)
	return result
}

// RLEdecode takes a run-length encoded string and returns the original string
func RLEdecode(data string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering RLEdecode(data=%v)", data)
	var result string
	regex := regexp.MustCompile(`(\d+)(\w)`)

	for _, match := range regex.FindAllStringSubmatch(data, -1) {
		num, _ := strconv.Atoi(match[1])
		result += strings.Repeat(match[2], num)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting RLEdecode(data=%v)", data)
	return result
}

// RLEncodebytes takes a byte slice and returns its run-length encoding as a byte slice
func RLEncodebytes(data []byte) []byte {
	logritelog.Default().Printf("[ERROR] >>> Entering RLEncodebytes(data=%v)", data)
	var result []byte
	var count byte = 1

	for i := 0; i < len(data); i++ {
		if i+1 < len(data) && data[i] == data[i+1] {
			count++
			continue
		}
		result = append(result, count, data[i])
		count = 1
		logritelog.Default().Printf("[DEBUG] RLEncodebytes(data=%v): count → %v", data, count)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting RLEncodebytes(data=%v)", data)
	return result
}

// RLEdecodebytes takes a run-length encoded byte slice and returns the original byte slice
func RLEdecodebytes(data []byte) []byte {
	logritelog.Default().Printf("[ERROR] >>> Entering RLEdecodebytes(data=%v)", data)
	var result []byte

	for i := 0; i < len(data); i += 2 {
		count := int(data[i])
		result = append(result, bytes.Repeat([]byte{data[i+1]}, count)...)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting RLEdecodebytes(data=%v)", data)
	return result
}
