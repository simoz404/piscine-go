package piscine

func BasicJoin(elems []string) string {
	var out string
	for _, v := range elems {
		out = out + v
	}
	return out
}
