package piscine

func Compact(x *[]string) int {
	avoir := *x
	count := 0
	for _, val := range *x {
		if val != "" {
			avoir[count] = val
			count++
		}
	}
	*x = avoir[0:count]
	return count
}
