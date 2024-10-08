package strutils

import "strings"

type Strings []string

func (ss Strings) ToStringSlice() []string {
	return ss
}

func (ss Strings) ToInterfaceSlice() []interface{} {
	iface := make([]interface{}, len(ss))
	for i := range ss {
		iface[i] = ss[i]
	}
	return iface
}

func (ss Strings) IsIn(value string) bool {
	if len(value) == 0 {
		return false
	}
	v := strings.ToUpper(value)
	for _, s := range ss {
		if strings.ToUpper(s) == v {
			return true
		}
	}
	return false
}

func (ss Strings) IsInCaseSensitive(value string) bool {
	if len(value) == 0 {
		return false
	}

	for _, s := range ss {
		if s == value {
			return true
		}
	}
	return false
}
