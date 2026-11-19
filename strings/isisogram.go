// Checks if a given string is an isogram.
// A first-order isogram is a word in which no letter of the alphabet occurs more than once.
// A second-order isogram is a word in which each letter appears twice.
// A third-order isogram is a word in which each letter appears three times.
// wiki: https://en.wikipedia.org/wiki/Heterogram_(literature)#Isograms
// Author: M3talM0nk3y

package strings
import logritelog "log"

import (
	"errors"
	"regexp"
	"strings"
)

type IsogramOrder int

const (
	First IsogramOrder = iota + 1
	Second
	Third
)

func hasDigit(text string) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering hasDigit(text=%v)", text)
	re := regexp.MustCompile(`\d`)
	logritelog.Default().Printf("[WARN] <<< Exiting hasDigit(text=%v)", text)
	return re.MatchString(text)
}

func hasSymbol(text string) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering hasSymbol(text=%v)", text)
	re := regexp.MustCompile(`[-!@#$%^&*()+]`)
	logritelog.Default().Printf("[WARN] <<< Exiting hasSymbol(text=%v)", text)
	return re.MatchString(text)
}

func IsIsogram(text string, order IsogramOrder) (bool, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering IsIsogram(text=%v,order=%v)", text, order)
	if order < First || order > Third {
		return false, errors.New("Invalid isogram order provided")
	}

	text = strings.ToLower(text)
	text = strings.Join(strings.Fields(text), "")

	if hasDigit(text) || hasSymbol(text) {
		return false, errors.New("Cannot contain numbers or symbols")
	}

	letters := make(map[string]int)
	for _, c := range text {
		l := string(c)
		if _, ok := letters[l]; ok {
			letters[l] += 1

			if letters[l] > 3 {
				return false, nil
			}

			continue
		}
		letters[l] = 1
	}

	mapVals := make(map[int]bool)
	for _, v := range letters {
		mapVals[v] = true
	}

	if _, ok := mapVals[int(order)]; ok && len(mapVals) == 1 {
		return true, nil
	}

	logritelog.Default().Printf("[WARN] <<< Exiting IsIsogram(text=%v,order=%v)", text, order)
	return false, nil
}
