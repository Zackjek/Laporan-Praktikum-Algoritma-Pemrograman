package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    var nama, nim, kelas string

    // Membaca input pengguna
    fmt.Print("Masukkan Nama: ")
    nama, _ = reader.ReadString('\n')
    fmt.Print("Masukkan NIM: ")
    nim, _ = reader.ReadString('\n')
    fmt.Print("Masukkan Kelas: ")
    kelas, _ = reader.ReadString('\n')

    // Menghapus karakter newline dari input
    nama = nama[:len(nama)-1]
    nim = nim[:len(nim)-1]
    kelas = kelas[:len(kelas)-1]

    // Menampilkan resume singkat
    fmt.Println("\n--- Resume Mahasiswa ---")
    fmt.Println("Nama  :", nama)
    fmt.Println("NIM   :", nim)
    fmt.Println("Kelas :", kelas)
}
