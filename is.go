package strutils

import "unicode"

func IsLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsAlpha(s string) bool {
	for _, c := range s {
		if !IsLetter(c) {
			return false
		}
	}
	return true
}

func IsUnicodeLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsDigit(c rune) bool {
	return ('0' <= c && c <= '9')
}

func IsNumber(s string) bool {
	any := false
	for _, c := range s {
		if !IsDigit(c) {
			return false
		}
		any = true
	}
	return any
}
func IsNumberOrNumberWithSign(s string) bool {
	havesign := false
	havedigits := false
	for i, c := range s {
		if !IsDigit(c) {
			if i == 0 {
				if s == "-" || s == "+" {
					havesign = true
					continue
				}
			}
			return false
		} else {
			havedigits = true
		}
	}
	if havesign && !havedigits {
		return false
	}
	return true
}
