package piscine

func CollatzCountdown(debut int) int {
	if debut <= 0 {
		return -1
	}
	count := 0
	for debut != 1 {
		if debut%2 == 0 {
			debut = debut / 2
		} else {
			debut = (debut * 3) + 1
		}
		count++
	}
	return count
}
