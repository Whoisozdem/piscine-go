package piscine

func Join(strs []string, x string) string {
	res := strs[0]
	for _, y := range strs[1:] {
		res = res + x + y
	}
	return res
}
