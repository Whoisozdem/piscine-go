package piscine

func SplitWhiteSpaces(str string) []string {
	Teextchanger := ""
	t := []string{}
	for i, v := range str {
		if i == lent2(str)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			Teextchanger += string(v)
			t = append(t, Teextchanger)
		} else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			Teextchanger += string(v)
		} else {
			if len(Teextchanger) >= 1 {
				t = append(t, Teextchanger)
			}
			Teextchanger = ""
		}
	}
	return t
}

func lent2(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}
