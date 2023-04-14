package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for range tab {
		length++
	}
	capa := length
	sort := 0
	for i := 0; i < length-1; i++ {
		isSorted := f(tab[i], tab[i+1])
		if isSorted > 0 {
			sort++
		} else if isSorted < 0 {
			sort--
		} else {
			capa--
		}
	}
	capa--

	if sort == capa || sort == -1*capa {
		return true
	}
	return false
}
