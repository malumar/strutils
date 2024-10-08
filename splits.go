package strutils

import "strings"

// zamiana:
// zmienna -> zmienna
// zmiennaInna -> zmienna_Inna
// ZmiennnaInnaDruga -> Zmienna_Inna_Drug
func UnderscoreCase(s string) string {
	sp := Split(s, "_")
	if len(sp) > 0 {
		//var r string
		//for _, v := range sp {
		//	if len(v) == 0 {
		//		continue
		//	}
		//	if len(r) > 0 {
		//		r += "_" + v
		//	} else {
		//		r += v
		//	}
		//}
		//return r
		return strings.Join(sp, "_")
	} else {
		return s
	}
}

// zamiana:
// zmienna -> zmienna
// zmiennaInna -> zmienna_inna
// ZmiennnaInnaDruga -> zmienna_inna_drug
func UnderscoreCaseLower(s string) string {
	sp := Split(s, "_")
	if len(sp) > 0 {
		return strings.ToLower(strings.Join(sp, "_"))
	} else {
		return s
	}
}

// zamiana
// zmienna --> zmienna
// zmienna_inna ZmiennaInna
func PascalCase(s string) string {
	items := strings.Split(s, "_")
	if len(items) > 0 {

	}
	return FirstUpper(s)
}

// zmienna abc -> zmiennaabc
// zmiena ABC -> zmiennaABC
// zmiena_abc -> zmienaabc
func FlatCase(s string) string {
	items := Split(s, "")
	if len(items) > 0 {
		return strings.Join(items, "")
	}
	return s
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
	sp := Split(s, "_")
	if len(sp) > 0 {
		var rc string
		first := false
		for _, str := range sp {
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

func PascalSnakeCaseAllFirstUpper(s string) string {
	return strings.Replace(SnakeCaseAllFirstUpper(strings.Replace(s, "_", " ", -1)), " ", "", -1)
}

func SnakeCaseAllFirstUpper(s string) string {
	sp := Split(s, "")
	if len(sp) > 0 {
		var rc string
		for _, str := range sp {
			if len(str) > 1 {
				str = strings.ToUpper(str[:1]) + str[1:]
			}

			rc += str
		}
		return rc
	}
	return s
}

// zamiana:
// zmienna -> Zmienna
// zmiennaInna -> ZmiennaInna
// ZmiennnaInnaDruga -> ZmiennaInnaDruga
// Zmienna1 Zmienna2 -> Zmienna1Zmienna2
func SnakeCaseFirstUpper(s string) string {
	sp := Split(s, "_")
	if len(sp) > 0 {
		var rc string
		firstUpper := false
		for _, str := range sp {
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
	for _, v := range array {
		if v == search {
			return true
		}
	}
	return false
}

func FirstLower(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToLower(s[:1]) + s[1:]
}
func FirstUpper(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
