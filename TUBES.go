package main

import (
	"fmt"
)

var LowUrgentionMentalHealth = [20]string{"stres", "cemas", "depresi", "panik", "khawatir", "lelah", "burnout", "kelelahan_mental", "gangguan_tidur", "insomnia", "sedih", "murung", "trauma", "overthinking", "minder", "kecewa", "kesepian", "takut", "tegang", "mood swing"}
var HighUrgentionMentalHealth = [10]string{"bunuh_diri", "menyakiti_diri", "tidak_sanggup", "putus_asa", "ingin_mati", "sudah_tidak_tahan", "hilang_harapan", "ingin_menghilang", "tidak_berguna", "tidak_ada_yang_peduli"}

//var ProductivityKey = [13]string{"semangat", "motivasi", "tidak fokus", "malas", "kehilangan arah", "bingung", "distraksi", "manajemen waktu", "target", "tujuan", "goals", "disiplin", "tanggung jawab"}

type history struct {
	id           int //mewakili urutan chatting
	input, saran string
	keyword      []string
	urgensi      int
}

type arrChat [100]history

// MENU
func menu() {
	var input string
	var chat arrChat
	var nHistory int
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
			mentalHealthMode(&chat, &nHistory)
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
func mentalHealthMode(chat *arrChat, nHistory *int) {
	for {
		fmt.Println("\nMode Kesehatan Mental. Ketik 'menu' untuk kembali.")
		var toMenu string
		fmt.Print("\njika ingin melanjutkan ketik apa saja\n")
		fmt.Scan(&toMenu)
		if toMenu == "menu" {
			fmt.Println("\nKembali ke menu utama.\n")
			break
		} else {
			fmt.Print("Anda : ")
			chatsession(&*chat, *nHistory)
			chat[*nHistory].id = *nHistory + 1
			*nHistory = *nHistory + 1
			fmt.Print(*chat)
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
/*func keyDetectorHighUrgention(keyword *[10]string, word string, k *int) {
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
}*/

// menghilangkan tanda titik dalam suatu kata
func dotRemover(kata string) string {
	var kataBaru string
	for i := 0; i < len(kata)-1; i++ {
		kataBaru = kataBaru + string(kata[i])
	}
	return kataBaru
}

func rangkaiKalimat(chat *arrChat, i int, listkata *[]string) {
	var kata, sentence string
	for {
		fmt.Scan(&kata)
		kata = LowerCase(kata)
		if dotDetector(kata) {
			kata = dotRemover(kata)
			sentence = sentence + kata
			*listkata = append(*listkata, kata)
			break
		}
		sentence = sentence + kata + " "
		*listkata = append(*listkata, kata)
	}
	chat[i].input = sentence
	fmt.Println(*listkata)
}

func chatsession(chat *arrChat, i int) {
	var listkata []string
	rangkaiKalimat(&*chat, i, &listkata)
	keywordinput(&*chat, i, listkata)
}

/*sorting riwayat
func selectionSort(arr []int) {
	n := ....
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func insertionSort(arr []int) {
	n := ....
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}*/
//mencari keyword dengan urgensi tinggi dan rendah menggunakan function sequential search
func keywordinput(chat *arrChat, i int, listkata []string) {
	var j, nKeyword int
	for j = 0; j < len(listkata); j++ {
		usersWord := listkata[j]
		for k := 0; k < len(HighUrgentionMentalHealth); k++ {
			if usersWord == HighUrgentionMentalHealth[k] {
				chat[i].keyword[nKeyword] = usersWord
				nKeyword++
				if chat[i].urgensi == 0 {
					chat[i].urgensi = 3
				}
			}
		}
	}
	for j = 0; j < len(listkata); j++ {
		usersWord := listkata[j]
		for k := 0; k < len(LowUrgentionMentalHealth); k++ {
			if usersWord == LowUrgentionMentalHealth[k] {
				chat[i].keyword[nKeyword] = usersWord
				nKeyword++
				if chat[i].urgensi == 0 {
					chat[i].urgensi = 2
				}
			}
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
