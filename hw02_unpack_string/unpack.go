package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var ShieldSlash = `\`

type iteratorData struct {
	prevRune      int32
	currentRune   int32
	nextRune      int32
	slashSequence int
}

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
	var slashSequence int
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

		isDigitNext := unicode.IsDigit(nextRune)
		data := iteratorData{
			prevRune,
			currentRune,
			nextRune,
			slashSequence,
		}
		if isDigitNext {
			bFormedByNext(data, &b)
		} else {
			bFormedByCurrent(data, &b)
		}
	}

	return b.String(), nil
}

func bFormedByNext(data iteratorData, b *strings.Builder) {
	countOfRepeat, _ := strconv.Atoi(string(data.nextRune))
	switch string(data.currentRune) != ShieldSlash {
	case true:
		if string(data.prevRune) == ShieldSlash {
			countOfRepeat--
		}
		b.WriteString(copySign(data.currentRune, countOfRepeat))
	case false:
		if data.slashSequence > 2 {
			b.WriteString(ShieldSlash)
			b.WriteRune(data.nextRune)
		}

		if data.slashSequence <= 2 && string(data.prevRune) == ShieldSlash {
			b.WriteString(copySign(data.currentRune, countOfRepeat))
		}

		if data.slashSequence <= 2 && string(data.prevRune) != ShieldSlash {
			b.WriteRune(data.nextRune)
		}
	}
}

func bFormedByCurrent(data iteratorData, b *strings.Builder) {
	if !unicode.IsDigit(data.currentRune) && string(data.currentRune) != ShieldSlash {
		b.WriteRune(data.currentRune)
	}
}
