package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Datas struct {
	id        int64
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	data := getDatas()

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <nama/absen>")
		return
	}

	arg := os.Args[1]
	if absen, err := strconv.ParseInt(arg, 10, 64); err == nil {
		// Jika input berupa nomor absen
		findByAbsen(data, absen)
	} else {
		// Jika input berupa nama
		findByName(data, arg)
	}
}

func getDatas() []Datas {
	return []Datas{
		{id: 1, nama: "daniela", alamat: "banyuwangi", pekerjaan: "mahasiswa", alasan: "menambah pengalaman"},
		{id: 2, nama: "zoro", alamat: "surabaya", pekerjaan: "fullstack", alasan: "menambah jam terbang"},
		{id: 3, nama: "nami", alamat: "pamekasan", pekerjaan: "backend", alasan: "menambah relasi"},
		{id: 4, nama: "luffy", alamat: "sidoarjo", pekerjaan: "mahasiswa", alasan: "Meningkatkan Pengetahuan"},
	}
}

func findByAbsen(data []Datas, absen int64) {
	found := false
	for _, v := range data {
		if v.id == absen {
			fmt.Printf("Nama : %s\n", v.nama)
			fmt.Printf("Alamat : %s\n", v.alamat)
			fmt.Printf("Pekerjaan : %s\n", v.pekerjaan)
			fmt.Printf("Alasan : %s\n", v.alasan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Data peserta tidak ditemukan.")
	}
}

func findByName(data []Datas, name string) {
	name = strings.ToLower(name) // ngubah biar lowercase, mau dibikin seperti equals ignore case di java
	found := false
	for _, v := range data {
		if strings.EqualFold(strings.ToLower(v.nama), name) {
			fmt.Printf("Nama : %s\n", v.nama)
			fmt.Printf("Alamat : %s\n", v.alamat)
			fmt.Printf("Pekerjaan : %s\n", v.pekerjaan)
			fmt.Printf("Alasan : %s\n", v.alasan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Data peserta tidak ditemukan.")
	}
}
