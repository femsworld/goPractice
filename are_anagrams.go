package sprint

import (
	"sort"
	"strings"
)

func normalizeString(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}

func AreAnagrams(str1, str2 string) bool {
	normalizedStr1 := normalizeString(str1)
	normalizedStr2 := normalizeString(str2)

	return sortString(normalizedStr1) == sortString(normalizedStr2)
}

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
