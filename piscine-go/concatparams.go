package piscine

func ConcatParams(args []string) string {
	sti := ""
	x := 0
	for i := range args {
		x = i
	}
	for i := 0; i <= x; i++ {
		if i != x {
			sti = sti + args[i] + "\n"
		} else {
			sti = sti + args[i]
		}
	}
	return sti
}
