package utils

import "strings"

func BuildString(separator string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteString(separator)
	}

	return sb.String()[:sb.Len()-len(separator)]
}
