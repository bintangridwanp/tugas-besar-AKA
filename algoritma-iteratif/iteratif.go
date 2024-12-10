package algoritma_iteratif

func iterativeSequential(data []string, target string) int {
	var i int
	for i = 0; i < len(data); i++ {
		if data[i] == target {
			println(data[i])
			return i
		}
	}
	return -1
}