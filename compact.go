package piscine

func Compact(ptr *[]string) int {
	var s []string
	i := 0
	for _, v := range *ptr {
		if v != "" {
			s = append(s, v)
			i++
		}
	}
	*ptr = s
	return i
}
