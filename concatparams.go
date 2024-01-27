package piscine

func ConcatParams(args []string) string {
	lenS := 0
	for range args {
		lenS++
	}
	var s string
	for i := 0; i < lenS; i++ {
		s += args[i]
		if i != lenS-1 {
			s += "\n"
		}
	}
	return s
}
