package strutils

import (
	"unicode"
	"unicode/utf8"
	"strings"
)

// Skopiowane czesciowo z https://github.com/fatih/camelcase/blob/master/camelcase.go

// Split splits the camelcase word and returns a list of words. It also
// supports digits. Both lower camel case and upper camel case are supported.
// For more info please check: http://en.wikipedia.org/wiki/CamelCase
//
// Examples
//
//   "" =>                     [""]
//   "lowercase" =>            ["lowercase"]
//   "Class" =>                ["Class"]
//   "MyClass" =>              ["My", "Class"]
//   "MyC" =>                  ["My", "C"]
//   "HTML" =>                 ["HTML"]
//   "PDFLoader" =>            ["PDF", "Loader"]
//   "AString" =>              ["A", "String"]
//   "SimpleXMLParser" =>      ["Simple", "XML", "Parser"]
//   "vimRPCPlugin" =>         ["vim", "RPC", "Plugin"]
//   "GL11Version" =>          ["GL", "11", "Version"]
//   "99Bottles" =>            ["99", "Bottles"]
//   "May5" =>                 ["May", "5"]
//   "BFG9000" =>              ["BFG", "9000"]
//   "BöseÜberraschung" =>     ["Böse", "Überraschung"]
//   "Two  spaces" =>          ["Two", "  ", "spaces"]
//   "BadUTF8\xe2\xe2\xa1" =>  ["BadUTF8\xe2\xe2\xa1"]
//
// Splitting rules
//
//  1) If string is not valid UTF-8, return it without splitting as
//     single item array.
//  2) Assign all unicode characters into one of 4 sets: lower case
//     letters, upper case letters, numbers, and all other characters.
//  3) Iterate through characters of string, introducing splits
//     between adjacent characters that belong to different sets.
//  4) Iterate through array of split strings, and if a given string
//     is upper case:
//       if subsequent string is lower case:
//         move last character of upper case string to beginning of
//         lower case string
func Split(src string,skip string) (entries []string) {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return []string{src}
	}
	entries = []string{}
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	// handle upper case -> lower case sequences, e.g.
	// "PDFL", "oader" -> "PDF", "Loader"
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	anySkip := len(skip) > 0

	for _, s := range runes {
		if len(s) > 0 {
			if anySkip && skip == string(s) {
				continue
			}
			entries = append(entries, string(s))
		}
	}
	return
}

// zamiana:
// zmienna -> zmienna
// zmiennaInna -> zmienna_Inna
// ZmiennnaInnaDruga -> Zmienna_Inna_Drug
func UnderscoreCase(s string) string {
	sp:=Split(s, "_")
	if len(sp) > 0 {
		return strings.Join(sp,"_")
	} else {
		return s
	}
}

// zamiana:
// zmienna -> zmienna
// zmiennaInna -> zmienna_inna
// ZmiennnaInnaDruga -> zmienna_inna_drug
func UnderscoreCaseLower(s string) string {
	sp:=Split(s, "_")
	if len(sp) > 0 {
		return strings.ToLower(strings.Join(sp,"_"))
	} else {
		return s
	}
}



// zamiana:
// zmienna -> ZMIENNA
// zmiennaInna -> ZMIENNA_INNA
// ZmiennnaInnaDruga -> ZMIENNA_INNA_DRUGA
func UnderscoreCaseUpper(s string) string {
	return strings.ToUpper(UnderscoreCase(s))
}


// zamiana:
// zmienna -> zmienna
// zmiennaInna -> zmiennaInna
// ZmiennnaInnaDruga -> zmiennaInnaDruga
// Zmienna1 Zmienna2 -> zmienna1Zmienna2
func SnakeCaseFirstLower(s string) string {
	sp:=Split(s, "_")
	if len(sp) > 0 {
		var rc string
		first := false
		for _,str:=range sp {
			if !first {
				if len(str) > 1 {
					str = strings.ToLower(str[:1]) + str[1:]
					first = true
				}
			} else {
				if len(str) > 1 {
					str = strings.ToUpper(str[:1]) + str[1:]

				}
			}

			rc += str
		}
		return rc
	} else {
		return s
	}
}

// zamiana:
// zmienna -> Zmienna
// zmiennaInna -> ZmiennaInna
// ZmiennnaInnaDruga -> ZmiennaInnaDruga
// Zmienna1 Zmienna2 -> Zmienna1Zmienna2
func SnakeCaseFirstUpper(s string) string {
	sp:=Split(s, "_")
	if len(sp) > 0 {
		var rc string
		firstUpper := false
		for _,str:=range sp {
			if !firstUpper {
				if len(str) > 1 {
					str = strings.ToUpper(str[:1]) + str[1:]
					firstUpper = true
				}
			}

			rc += str
		}
		return rc
	} else {
		return s
	}
}


func InArray(array []string, search string) bool {
	if array == nil {
		return false
	}
	for _,v:=range array {
		if v == search {
			return true
		}
	}
	return false
}