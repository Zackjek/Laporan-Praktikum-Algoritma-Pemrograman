package main

import "fmt"

func main() {
    var fahrenheit float64

    // Membaca input pengguna
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&fahrenheit)

    // Menghitung suhu dalam Celsius
    celsius := (fahrenheit - 32) * 5 / 9

    // Menampilkan hasil konversi
    fmt.Printf("Suhu dalam Celsius: %.2f\n", celsius)
}
