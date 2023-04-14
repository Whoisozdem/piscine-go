package piscine

func Unmatch(tab []int) int {
	for i, y := range tab {
		count := 1
		for j, x := range tab {
			if y == x && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return y
		}
	}
	return -1
}
