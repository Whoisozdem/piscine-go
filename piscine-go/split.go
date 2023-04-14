package piscine

func Split(str, cara string) []string {
	x := 0

	for i := range cara {
		x = i + 1
	}

	ligne := 0
	for i := range str {
		ligne = i + 1
	}

	for i := 0; i < ligne-x; i++ {
		if str[i:i+x] == cara {
			str = str[:i] + " " + str[i+x:]
			ligne -= x
		}
	}

	return SplitWhiteSpaces(str)
}
