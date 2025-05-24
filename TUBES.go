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
	var i, j, nHistory int
	i = 0
	j = 0
	nHistory = 1
	for {
		fmt.Println("\nMode Kesehatan Mental. Ketik 'menu' untuk kembali.")
		var kata, sentence, toMenu string
		fmt.Print("Ketik 'menu' jika ingin kembali ke menu\njika ingin melanjutkan enter saja")
		fmt.Scan(&toMenu)
		if toMenu == "menu" {
			fmt.Println("\nKembali ke menu utama.\n")
			break
		} else {
			fmt.Print("Anda : ")
			chat[i].input = rangkaiKalimat()
			keyDetectorMentalHealth(kata, &chat[i].keyword, &nkata, HighUrgentionMentalHealth[:])
			keyDetectorMentalHealth(kata, &chat[i].keyword, &nkata, LowUrgentionMentalHealth[:])
			i++
			nHistory++
		}
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
func keyDetectorHighUrgention(keyword *[10]string, word string, k *int) {
	var i int
	if dotDetector(word) {
		word = dotRemover(word)
	}
	for i = 0; i < len(HighUrgentionMentalHealth); i++ {
		if word == HighUrgentionMentalHealth[i] {
			*keyword[*k] = word
			*k++
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

func rangkaiKalimat() string {
	var kata, sentence string
	for {
		fmt.Scan(&kata)
		kata = LowerCase(kata)
		if dotDetector(kata) {
			sentence = sentence + dotRemover(kata)
			break
		}
		sentence = sentence + kata + " "
	}
	return sentence
}

// memecah kalimat menjadi beberapa kata
func splitKalimat(kalimat string, listkata *[100]string, nkata *int) {
	var kata string
	var j int = 0
	*nkata = 1
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] != ' ' {
			kata += string(kalimat[i])
		} else {
			if kata != "" {
				listkata[j] = kata
				j++
				*nkata = *nkata + 1
				kata = ""
			}
		}
	}
	// simpan kata terakhir kalau ada
	if kata != "" {
		listkata[j] = kata
	}
}

func keyDetectorMentalHealth(kalimat string, keyword *[10]string, idx *int, Basis []string) {
	var i, j int
	var cocok bool
	for j = 0; j < len(Basis); j++ {
		if *idx >= 10 {
			break
		}
		cocok = false
		kata := Basis[j]
		// pencarian kata yang sesuai dengan membandingkan perhurufnya
		for i = 0; i <= len(kalimat)-len(kata); i++ {
			match := true
			for k := 0; k < len(kata); k++ {
				if kalimat[i+k] != kata[k] {
					match = false
					break
				}
			}
			// cocokkan juga jika karakter setelah frasa adalah spasi atau akhir kalimat
			if match {
				// frasa cocok — pastikan batas kata
				// cek apakah karakter sebelumnya spasi atau awal kalimat
				if (i == 0 || kalimat[i-1] == ' ') &&
					(i+len(kata) == len(kalimat) || kalimat[i+len(kata)] == ' ') {
					cocok = true
					break
				}
			}
		}
		if cocok {
			keyword[*idx] = kata
			*idx++
		}
	}
}
