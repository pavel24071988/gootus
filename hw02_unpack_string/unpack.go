package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func copySign(asRune int32, count int) string {
	return strings.Repeat(string(asRune), count)
}

func shieldingAvailable(currentRune int32, nextRune int32) error {
	if string(currentRune) == `\` && !unicode.IsDigit(nextRune) && string(nextRune) != `\` {
		return ErrInvalidString
	}

	return nil
}

func multiplyAvailable(prevRune int32, currentRune int32, nextRune int32) error {
	if string(prevRune) != `\` && unicode.IsDigit(currentRune) && unicode.IsDigit(nextRune) {
		return ErrInvalidString
	}

	return nil
}

func checkFirstSign(asRune []rune) error {
	if len(asRune) > 0 && unicode.IsDigit(asRune[0]) {
		return ErrInvalidString
	}

	return nil
}

func Unpack(phrase string) (string, error) {
	var b strings.Builder
	var err error
	var slashSequence int8
	shieldSlash := `\`
	asRune := []rune(phrase)

	err = checkFirstSign(asRune)
	if err != nil {
		return "", err
	}

	for i, currentRune := range asRune {
		var prevRune int32
		if i > 0 {
			prevRune = asRune[i-1]
		}
		var nextRune int32
		if i+1 != len(asRune) {
			nextRune = asRune[i+1]
		}

		err = shieldingAvailable(currentRune, nextRune)
		if err != nil {
			return "", err
		}

		err = multiplyAvailable(prevRune, currentRune, nextRune)
		if err != nil {
			return "", err
		}

		if string(currentRune) == shieldSlash {
			slashSequence++
		} else {
			slashSequence = 0
		}

		if unicode.IsDigit(nextRune) {
			countOfRepeat, _ := strconv.Atoi(string(nextRune))
			if string(currentRune) != shieldSlash {
				if string(prevRune) == shieldSlash {
					countOfRepeat--
				}
				b.WriteString(copySign(currentRune, countOfRepeat))
			} else if slashSequence > 2 {
				b.WriteString(shieldSlash)
				b.WriteRune(nextRune)
			} else if string(prevRune) == shieldSlash {
				b.WriteString(copySign(currentRune, countOfRepeat))
			} else {
				b.WriteRune(nextRune)
			}
		} else if !unicode.IsDigit(currentRune) && string(currentRune) != shieldSlash {
			b.WriteRune(currentRune)
		}
	}

	return b.String(), nil
}
