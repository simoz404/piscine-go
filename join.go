package piscine

func Join(strs []string, sep string) string {
	var out string
	lenS := 0
	for range strs {
		lenS++
	}
	for i := 0; i < lenS; i++ {
		out = out + strs[i]
		if i != lenS-1 {
			out += sep
		}
	}
	return out
}
