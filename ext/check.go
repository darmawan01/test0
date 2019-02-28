package ext

import "strings"

// Check the string
func Check(s string) (int, int) {
	newVal := strings.ToLower(s)
	aHidup, bHidup := strings.Count(newVal, "a"), strings.Count(newVal, "o")
	mati := (len(newVal) - 1) - (aHidup + bHidup)

	if aHidup > 1 {
		aHidup = 1
	} else if bHidup > 0 {
		bHidup = 1
	}

	hidup := aHidup + bHidup

	return mati, hidup
}
