package main

import (
	"fmt"
)

var LowUrgentionMentalHealth = [20]string{"stres", "cemas", "depresi", "panik", "khawatir", "lelah", "burnout", "kelelahan mental", "gangguan tidur", "insomnia", "sedih", "murung", "trauma", "overthinking", "minder", "kecewa", "kesepian", "takut", "tegang", "mood swing"}
var HighUrgentionMentalHealth = [10]string{"bunuh diri", "menyakiti diri", "tidak sanggup", "putus asa", "ingin mati", "sudah tidak tahan", "hilang harapan", "ingin menghilang", "tidak berguna", "tidak ada yang peduli"}
var ProductivityKey = [13]string{"semangat", "motivasi", "tidak fokus", "malas", "kehilangan arah", "bingung", "distraksi", "manajemen waktu", "target", "tujuan", "goals", "disiplin", "tanggung jawab"}

type history struct {
	date         string
	input, saran string
	keyword      [10]string
}

type arrChat [100]history

// MENU
func menu() {
	var input string
	var chat arrChat
	var access bool = true
	for access {
		fmt.Println("Selamat datang di Chatbot AI Go!")
		fmt.Println("Pilih kategori:")
		fmt.Println("[1] Kesehatan Mental")
		fmt.Println("[2] Produktivitas")
		fmt.Println("Ketik 'exit' untuk keluar.")
		fmt.Scan(&input)
		input = LowerCase(input)
		switch input {
		case "1":
			mentalHealthMode(&chat)
		case "2":
			productivityMode()
		case "exit":
			access = false
		default:
			fmt.Println("Input Tidak Valid")
		}
	}
}

// PRODUKTIVITAS
func productivityMode() {
	/*fmt.Println("\n Mode Produktivitas. Ketik 'menu' untuk kembali.")
	var input, sentence string
	for {
		var kata, sentence string
		var i, nHistory int
		i = 0
		fmt.Println("Anda : ")
		fmt.Scan(&kata)
		if input == "menu" {
			fmt.Println("\nKembali ke menu utama.\n")
			return
		}
		for dotDetector(kata){
			sentence = sentence + kata + " "
			fmt.Scan(&kata)
		}
	}*/
}

// KESEHATAN MENTAL  NOTE : MASIH PERTIMBANGAN UNTUK PEMAKAIAN BREAK
func mentalHealthMode(chat *arrChat) {
	var i, nHistory int
	i = 0
	nHistory = 1
	for {
		fmt.Println("\nMode Kesehatan Mental. Ketik 'menu' untuk kembali.")
		var kata, sentence string
		fmt.Print("Anda : ")
		fmt.Scan(&kata)
		if kata == "menu" {
			fmt.Println("\nKembali ke menu utama.\n")
			break
		}
		for !dotDetector(kata) {
			keyDetector(&chat[i].keyword, kata)
			kata = LowerCase(kata)
			sentence = sentence + kata + " "
			fmt.Scan(&kata)
		}
		chat[i].input = sentence
		fmt.Print(chat[i].input)
		i++
		nHistory++
	}
}

func main() {
	menu()
}

// me-lower case inputan user
func LowerCase(kalimat string) string {
	var i int
	var newWord string
	var letter rune
	for i = 0; i < len(kalimat); i++ {
		letter = rune(kalimat[i])
		if letter >= 'A' && letter <= 'Z' {
			letter = letter + ('a' - 'A')
		}
		newWord = newWord + string(letter)
	}
	return newWord
}

// DOT DETECTOR
func dotDetector(kata string) bool {
	for i := 0; i < len(kata); i++ {
		if kata[i] == '.' {
			return true
		}
	}
	return false
}

// pendeteksi jika ada keyword yang sesuai, maka akan masuk ke array keyword untuk pemberian saran
func keyDetector(keyword *[10]string, word string) {
	var i, k int
	var newWord string
	for i = 0; i < 10; i++ {
		if dotDetector(word) {
			newWord = dotRemover(word)
		}
		if newWord == HighUrgentionMentalHealth[i] {
			keyword[k] = newWord
			k++
		}
	}
}

// menghilangkan tanda titik dalam suatu kata
func dotRemover(kata string) string {
	var kataBaru string
	for i := 0; i < len(kata)-1; i++ {
		kataBaru = kataBaru + string(kata[i])
	}
	return kataBaru
}
