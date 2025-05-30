package main

import (
	"fmt"
)

var LowUrgentionMentalHealth = [20]string{"stress", "cemas", "depresi", "panik", "khawatir", "lelah", "burnout", "kelelahan_mental", "gangguan_tidur", "insomnia", "sedih", "murung", "trauma", "overthinking", "minder", "kecewa", "kesepian", "takut", "tegang", "mood swing"}
var HighUrgentionMentalHealth = [10]string{"bunuh_diri", "menyakiti_diri", "tidak_sanggup", "putus_asa", "ingin_mati", "sudah_tidak_tahan", "hilang_harapan", "ingin_menghilang", "tidak_berguna", "tidak_ada_yang_peduli"}
var ActivitySuggestionsLow = [13]string{
	"Bercerita ke teman dekat",
	"Berjalan santai di pagi atau sore hari",
	"Mendengarkan musik yang menenangkan",
	"Meditasi ringan selama 5-10 menit",
	"Menggambar atau mewarnai",
	"Menonton film favorit yang menyenangkan",
	"Membaca buku yang disukai",
	"Melakukan peregangan atau olahraga ringan",
	"Tidur lebih awal dan cukup istirahat",
	"Membuat to-do list atau rencana mingguan",
	"Mendengarkan podcast inspiratif atau edukatif",
	"Mencoba teknik pernapasan dalam (deep breathing)",
	"Membatasi waktu media sosial dan screen time",
}

var ActivitySuggestionsHigh = [13]string{
	"Segera hubungi layanan bantuan profesional",
	"Ceritakan perasaanmu pada orang yang kamu percaya",
	"Temui guru, konselor, atau tenaga kesehatan terdekat",
	"Jauhkan diri dari benda atau hal yang membahayakan diri",
	"Ambil waktu untuk istirahat di tempat yang tenang",
	"Tuliskan isi pikiranmu di jurnal atau kertas",
	"Coba hubungi sahabat lama atau anggota keluarga",
	"Ingatkan diri bahwa kamu berharga dan tidak sendirian",
	"Berada di dekat orang lain walau hanya duduk bersama",
	"Cari bantuan dari layanan konseling online atau hotline",
	"Tonton video atau konten yang membuatmu tersenyum atau tertawa ringan",
	"Mengucapkan doa atau kalimat spiritual yang menenangkan",
	"Dekatkan diri kepada tuhan dan senantiasa meningatnya",
}

//var ProductivityKey = [13]string{"semangat", "motivasi", "tidak fokus", "malas", "kehilangan arah", "bingung", "distraksi", "manajemen waktu", "target", "tujuan", "goals", "disiplin", "tanggung jawab"}

type history struct {
	id      int //mewakili urutan chatting
	input   string
	keyword [10]string
	saran   []string
	urgensi int
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
		fmt.Println("[3] Riwayat chat")
		fmt.Println("Ketik 'exit' untuk keluar.")
		fmt.Scan(&input)
		input = LowerCase(input)
		switch input {
		case "1":
			mentalHealthMode(&chat, &nHistory)
			ClearScreen()
		case "2":
			productivityMode()
			ClearScreen()
		case "3":

		case "exit":
			access = false
		default:
			fmt.Println("Input Tidak Valid")
		}
	}
}

// PRODUKTIVITAS
func productivityMode() {

}

// KESEHATAN MENTAL  NOTE : MASIH PERTIMBANGAN UNTUK PEMAKAIAN BREAK
func mentalHealthMode(chat *arrChat, nHistory *int) {
	for {
		fmt.Println("\nMode Kesehatan Mental. Ketik 'menu' untuk kembali.")
		var toMenu string
		fmt.Print("jika ingin melanjutkan ketik apa saja\n")
		fmt.Scan(&toMenu)
		if toMenu == "menu" {
			fmt.Println("\nKembali ke menu utama.\n")
			break
		} else {
			fmt.Print("Anda : ")
			chatsession(&*chat, *nHistory)
			chat[*nHistory].id = *nHistory + 1
			*nHistory = *nHistory + 1
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
}

func chatsession(chat *arrChat, i int) {
	var listkata []string
	rangkaiKalimat(&*chat, i, &listkata)
	keywordinput(&*chat, i, listkata)
}

// mencari keyword dengan urgensi tinggi dan rendah menggunakan function sequential search
func keywordinput(chat *arrChat, i int, listkata []string) {
	var j int
	for j = 0; j < len(listkata); j++ {
		usersWord := listkata[j]
		for k := 0; k < len(HighUrgentionMentalHealth); k++ {
			if usersWord == HighUrgentionMentalHealth[k] {
				chat[i].keyword = append(chat[i].keyword, usersWord)
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
				chat[i].keyword = append(chat[i].keyword, usersWord)
				if chat[i].urgensi == 0 {
					chat[i].urgensi = 2
				}
			}
		}
	}
}

func daftarSaran(chat *arrChat)

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
}

func binarySearch(arr []int, target int) int {
	n :=
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}*/

func ClearScreen() { //
	fmt.Print("\033[H\033[2J")
}
