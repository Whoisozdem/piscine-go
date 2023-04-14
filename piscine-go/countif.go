package piscine

func CountIf(f func(string) bool, tab []string) int {
	compteur := 0
	for _, stri := range tab {
		if f(stri) {
			compteur++
		}
	}
	return compteur
}
