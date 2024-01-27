package piscine

func PodiumPosition(podium [][]string) [][]string {
	lenP := 0
	for range podium {
		lenP++
	}
	for range podium {
		for i := 0; i < lenP-1; i++ {
			if podium[i][0] > podium[i+1][0] {
				podium[i], podium[i+1] = podium[i+1], podium[i]
			}
		}
	}
	return podium
}
