package piscine

func BasicAtoi(s string) int {
	num := 0
	a := 0
	a_s := []rune(s)
	for _, popo := range a_s { // On opte pour le blank identifier
		for i := '0'; i < popo; i++ {
			a++
		}
		num = num*10 + a
		a = 0
	}
	return num
}
