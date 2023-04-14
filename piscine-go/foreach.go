package piscine

func ForEach(f func(int), tab []int) {
	for _, valeur := range tab {
		f(valeur)
	}
}
