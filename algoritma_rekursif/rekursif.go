package algoritma_rekursif

// Fungsi untuk mencari data secara rekursif
func RecursiveSequential(data []string, target string, index int) int {
    if index >= len(data) {
        return -1 // Tidak ditemukan
    }

    if data[index] == target {
        return index // Ditemukan
    }

    return RecursiveSequential(data, target, index+1) // Pindah ke elemen berikutnya
}
