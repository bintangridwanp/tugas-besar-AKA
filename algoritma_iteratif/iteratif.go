package algoritma_iteratif

// Fungsi untuk mencari data secara iteratif
func IterativeSequential(data []string, target string) int {
    for i := 0; i < len(data); i++ {
        if data[i] == target {
            return i // Jika ditemukan, kembalikan index
        }
    }
    return -1 // Jika tidak ditemukan
}
