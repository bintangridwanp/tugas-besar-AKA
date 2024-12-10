package main

import (
    "fmt"
    "time"
    "tugas-besar-AKA/algoritma_iteratif"   // Nama paket yang benar
    "tugas-besar-AKA/algoritma_rekursif"   // Nama paket yang benar
)

var kabupaten_kota = []string{
    "Kabupaten Aceh Barat", "Kabupaten Aceh Barat Daya", "Kabupaten Aceh Besar", "Kabupaten Aceh Jaya", "Kabupaten Aceh Selatan", 
    "Kabupaten Aceh Singkil", "Kabupaten Aceh Tamiang", "Kabupaten Aceh Tengah", "Kabupaten Aceh Tenggara", "Kabupaten Aceh Timur", 
    "Kabupaten Aceh Utara", "Kabupaten Bener Meriah", "Kabupaten Bireuen", "Kabupaten Gayo Lues", "Kabupaten Nagan Raya", 
    "Kabupaten Pidie", "Kabupaten Pidie Jaya", "Kabupaten Simeulue", "Kota Banda Aceh", "Kota Langsa", "Kota Lhokseumawe", 
    "Kota Sabang", "Kota Subulussalam", "Kabupaten Asahan", "Kabupaten Batu Bara", "Kabupaten Dairi", "Kabupaten Deli Serdang", 
    "Kabupaten Humbang Hasundutan", "Kabupaten Karo", "Kabupaten Labuhanbatu", "Kabupaten Labuhanbatu Selatan", "Kabupaten Labuhanbatu Utara", 
    "Kabupaten Langkat", "Kabupaten Mandailing Natal", "Kabupaten Nias", "Kabupaten Nias Barat", "Kabupaten Nias Selatan", 
    "Kabupaten Nias Utara", "Kabupaten Padang Lawas", "Kabupaten Padang Lawas Utara", "Kabupaten Pakpak Bharat", "Kabupaten Samosir", 
    "Kabupaten Serdang Bedagai", "Kabupaten Simalungun", "Kabupaten Tapanuli Selatan", "Kabupaten Tapanuli Utara", "Kabupaten Toba Samosir", 
    "Kota Medan", "Kota Padang", "Kota Pekanbaru", "Kota Palembang", "Kota Bandar Lampung", "Kota Jakarta", "Kota Surabaya",
}

func main() {
    start := time.Now()
    
    // Memanggil algoritma iteratif dan rekursif
    indexIteratif := algoritma_iteratif.IterativeSequential(kabupaten_kota, "Kabupaten Konawe")
    indexRekursif := algoritma_rekursif.RecursiveSequential(kabupaten_kota, "Kabupaten Konawe", 0)

    // Output hasil pencarian
    fmt.Printf("Index (Iteratif): %d\n", indexIteratif)
    fmt.Printf("Index (Rekursif): %d\n", indexRekursif)

    fmt.Println("Waktu Algoritma:", time.Since(start))
}
