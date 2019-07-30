package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Alamat: ")
	text2, _ := reader.ReadString('\n')
	fmt.Print("Jenis Barang : ")
	text3, _ := reader.ReadString('\n')
	fmt.Print("Harga : Rp.")
	text4, _ := reader.ReadString('\n')
	// print
	fmt.Printf("\"%s\", \"%s\", \"%s\", \"%s\"\n",
		strings.TrimSpace(text), strings.TrimSpace(text2), strings.TrimSpace(text3), strings.TrimSpace(text4))
}
