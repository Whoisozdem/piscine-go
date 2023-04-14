package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, x := range table {
		mot := x
		for _, j := range mot {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
