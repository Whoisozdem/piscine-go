package piscine

func mon_StrLen(stri string) int {
	astr := []rune(stri)
	a := 1
	for index := range astr {
		a = index
	}
	return a + 1
}

func StrRev(s string) string {
	a := mon_StrLen(s)
	a = a - 1
	a_s := []rune(s)
	b_s := []rune(s)
	for index := range a_s {
		b_s[index] = a_s[a-index]
	}
	return string(b_s)
}
