package algoritma_iteratif

//import "fmt"

// Fungsi untuk mencari data secara iteratif
func IterativeSequential(data []string, target string) int {
    for i := 0; i < len(data); i++ {
        if data[i] == target {
        //fmt.Println(data[i])
            return i // Jika ditemukan, kembalikan index
        }
    }
    return -1 // Jika tidak ditemukan
}
