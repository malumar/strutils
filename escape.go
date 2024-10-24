package strutils

import "strings"

func EscWildcard(source string, addPfx bool) string {
	// \* -> {WILDCARD_STAR}
	output := strings.Replace(source, "\\*", "{WILDCARD_STAR}", -1)
	// \? -> {WILDCARD_QM}
	output = strings.Replace(output, "\\?", "{WILDCARD_QM}", -1)
	// % -> \%
	output = strings.Replace(output, "%", "\\%", -1)
	// _ -> \_
	output = strings.Replace(output, "_", "\\_", -1)
	// * -> %
	output = strings.Replace(output, "*", "%", -1)
	// ? -> _
	output = strings.Replace(output, "?", "_", -1)
	// We're coming back
	// \* -> {WILDCARD_STAR}
	output = strings.Replace(output, "{WILDCARD_STAR}", "*", -1)
	// \? -> {WILDCARD_QM}
	output = strings.Replace(output, "{WILDCARD_QM}", "?", -1)

	if addPfx {
		if !strings.HasPrefix(output, "%") {
			output = "%" + output
		}
		if !strings.HasSuffix(output, "%") {
			output = output + "%"
		}
	}
	return output
}
