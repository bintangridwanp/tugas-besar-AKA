package algoritma_rekursif

func recursiveSequential(data []string, target string, index int) int {
	if index >= len(data) {
		return -1 // Tidak ditemukan
	}

	if data[index] == target {
		return index // Ditemukan
	}

	return recursiveSequential(data, target, index+1) // Pindah ke elemen berikutnya
}