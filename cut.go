package strutils

func CutStringToLen(val string, maxLen int) string {
	if maxLen <= 0 || len(val) < maxLen {
		return val
	}
	return val[:maxLen]
}
