package strutils

import (
	"strings"
	"unicode"
)

func ReplaceAll(s string, newString string, searches ...string) string {
	for _, ss := range searches {
		s = strings.ReplaceAll(s, ss, newString)
	}
	return s
}

type Pair struct {
	Search  string
	Replace string
}

type PairOfRunes struct {
	Search  rune
	Replace rune
}

func ReplaceAllPairs(s string, pairs ...Pair) string {
	for _, p := range pairs {
		s = strings.ReplaceAll(s, p.Search, p.Replace)
	}
	return s
}

func RemoveDoubleWhiteSpace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := range str {
		if !(str[i] == 32 && (i+1 < len(str) && str[i+1] == 32)) {
			b.WriteRune(rune(str[i]))
		}
	}
	return b.String()
}

func RemoveUtf8PlAccentsFromQuery(text string) string {
	return ReplacePairOfRunes(text, PL_UTF8_CHARSET_PAIRS...)
}

func ReplacePairOfRunes(s string, pairs ...PairOfRunes) string {
	var sb strings.Builder
	for _, r := range s {
		for _, p := range pairs {
			if r == p.Search {
				r = p.Replace
				break
			}
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func RemoveWhiteSpaces(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if !unicode.IsSpace(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func RemoveNonAlphaNumeric(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

var PL_UTF8_CHARSET_PAIRS = []PairOfRunes{
	PairOfRunes{'ę', 'e'},
	PairOfRunes{'ó', 'o'},
	PairOfRunes{'ą', 'a'},
	PairOfRunes{'ś', 's'},
	PairOfRunes{'ł', 'l'},
	PairOfRunes{'ż', 'z'},
	PairOfRunes{'ź', 'z'},
	PairOfRunes{'ć', 'c'},
	PairOfRunes{'ń', 'N'},
	PairOfRunes{'Ę', 'E'},
	PairOfRunes{'Ó', 'O'},
	PairOfRunes{'Ą', 'A'},
	PairOfRunes{'Ś', 'S'},
	PairOfRunes{'Ł', 'L'},
	PairOfRunes{'Ż', 'Z'},
	PairOfRunes{'Ź', 'Ć'},
	PairOfRunes{'Ń', 'N'},
}
