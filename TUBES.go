package main

import (
	"fmt"
	"math/rand"
	"time"
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

type history struct {
	id      int //mewakili urutan chatting
	input   string
	keyword []string
	saran   []string
	urgensi int
}

type arrChat [100]history

// MENU
func menu() {
	var input string
	var chat arrChat
	var nID int = 1
	var nHistory int
	var access bool = true
	for access {
		fmt.Println("Selamat datang di Chatbot AI Go!")
		fmt.Println("Pilih kategori:")
		fmt.Println("[1] Kesehatan Mental")
		fmt.Println("[2] Laporan Produktivitas")
		fmt.Println("[3] Riwayat chat")
		fmt.Println("Ketik 'exit' untuk keluar.")
		fmt.Scan(&input)
		input = LowerCase(input)
		switch input {
		case "1":
			ClearScreen()
			mentalHealthMode(&chat, &nHistory, &nID)
			ClearScreen()
		case "2":
			laporanProduktivitas(chat, nHistory)
		case "3":
			ClearScreen()
			Riwayat(&chat, &nHistory)
			ClearScreen()
		case "exit":
			access = false
		default:
			fmt.Println("Input Tidak Valid")
		}
	}
}

// KESEHATAN MENTAL
func mentalHealthMode(chat *arrChat, nHistory *int, nID *int) {
	for {
		fmt.Println("\nMode Kesehatan Mental. Ketik 'menu' untuk kembali.")
		var toMenu string
		fmt.Print("jika ingin melanjutkan ketik apa saja\n")
		fmt.Scan(&toMenu)
		if toMenu == "menu" {
			fmt.Println("\nKembali ke menu utama.")
			break
		} else {
			fmt.Print("Anda : ")
			chatsession(&*chat, *nHistory)
			chat[*nHistory].id = *nID
			*nHistory = *nHistory + 1
			*nID = *nID + 1
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
	daftarSaran(&*chat, i)
	fmt.Println(chat[i].keyword, chat[i].saran)
	cetakSaran(*chat, i)
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

func daftarSaran(chat *arrChat, i int) {
	rand.Seed(time.Now().UnixNano()) // Agar hasil acaknya beda setiap dijalankan
	var j, k int
	var index int
	for j = 0; j < len(chat[i].keyword); j++ {
		for k = 0; k < len(HighUrgentionMentalHealth); k++ {
			if chat[i].keyword[j] == HighUrgentionMentalHealth[k] {
				index = rand.Intn(len(ActivitySuggestionsHigh))
				chat[i].saran = append(chat[i].saran, ActivitySuggestionsHigh[index])
			}
		}
		for k = 0; k < len(LowUrgentionMentalHealth); k++ {
			if chat[i].keyword[j] == LowUrgentionMentalHealth[k] {
				index = rand.Intn(len(ActivitySuggestionsLow))
				chat[i].saran = append(chat[i].saran, ActivitySuggestionsLow[index])
			}
		}
	}
}

// fungsi untuk mencetak aktivitas apa saja yang disarankan untuk user
func cetakSaran(chat arrChat, i int) {
	fmt.Print("berikut list aktivitas yang disarankan:\n")
	for j := 0; j < len(chat[i].saran); j++ {
		fmt.Printf("%d. %s\n", j+1, chat[i].saran[j])
	}
	fmt.Println()
}

func laporanProduktivitas(chat arrChat, n int) {
	var i, j, nSaran int
	var laporanSaran [20]string
	if n > 0 {
		for i = 0; i < n; i++ {
			for j = 0; j < len(chat[i].saran); j++ {
				if isntExist(chat[i].saran[j], laporanSaran, nSaran) {
					if nSaran < 20 {
						laporanSaran[nSaran] = chat[i].saran[j]
						nSaran++
					}
				}
			}
		}
		fmt.Println("berikut daftar saran yang harus dilakukan dari chat yang anda lakukan")
		for i = 0; i < nSaran; i++ {
			fmt.Printf("%d. %s\n", i+1, laporanSaran[i])
		}
	} else {
		fmt.Println("data percakapan kosong, lakukan konsultasi terlebih dahulu")
	}
	fmt.Println()
}

func isntExist(kata string, arrSaran [20]string, n int) bool {
	for i := 0; i < n; i++ {
		if kata == arrSaran[i] {
			return false
		}
	}
	return true
}

func ClearScreen() { //
	fmt.Print("\033[H\033[2J")
}

// sorting dari urgensi, yang mana berurgensi tinggi bernilai 3 dan urgensi rendah bernilai 2
func sortingChatByUrgency(chat *arrChat, n int) {
	var temp history
	for i := 1; i < n; i++ {
		temp = chat[i]
		j := i
		for j > 0 && chat[j-1].urgensi < temp.urgensi {
			chat[j] = chat[j-1]
			j--
		}
		chat[j] = temp
	}
}

func sortingChatByID(chat *arrChat, n int) {
	var pass, i, idx int
	var temp history
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if chat[idx].id > chat[i].id {
				idx = i
			}
		}
		temp = chat[pass-1]
		chat[pass-1] = chat[idx]
		chat[idx] = temp
	}
}

// mencari ID yang sesuai dengan binary search
func binarySearchID(chat arrChat, n int, targetID int) int {
	var low, high, mid int
	low = 0
	high = n - 1
	for low <= high {
		mid = (high + low) / 2
		if chat[mid].id > targetID {
			high = mid - 1
		} else if chat[mid].id < targetID {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func cetakriwayat(chat arrChat, n int) {
	fmt.Println("-------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Println("ID :", chat[i].id)
		fmt.Printf("Anda :\n%s\n", chat[i].input)
		fmt.Println("Bot :")
		cetakSaran(chat, i)
		fmt.Println("-------------------------------------------------------")
	}
}

func Riwayat(chat *arrChat, n *int) {
	var pilih int
	for {
		cetakriwayat(*chat, *n)
		fmt.Println("[1] urut berdasarkan urgensi")
		fmt.Println("[2] urut berdasarkan waktu chat dilakukan")
		fmt.Println("[3] hapus riwayat")
		fmt.Println("[4] edit riwayat")
		fmt.Println("[5] keluar")
		fmt.Print("pilih dari 1 sampai 5 :")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			sortingChatByUrgency(&*chat, *n)
		case 2:
			sortingChatByID(&*chat, *n)
		case 3:
			hapusRiwayat(&*chat, &*n)
		case 4:
			editRiwayat(&*chat, &*n)
		}
		if pilih == 5 {
			break
		}
	}
}

func hapusRiwayat(chat *arrChat, n *int) {
	var ID, index int
	if *n == 0 {
		fmt.Println("Belum ada riwayat yang bisa dihapus.")
		return
	}
	sortingChatByID(chat, *n)
	fmt.Print("Masukkan ID riwayat yang ingin dihapus: ")
	fmt.Scan(&ID)
	index = binarySearchID(*chat, *n, ID)
	if index != -1 {
		for i := index; i < *n-1; i++ {
			chat[i] = chat[i+1]
		}
		*n = *n - 1
		fmt.Println("Riwayat berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func editRiwayat(chat *arrChat, n *int) {
	var ID, target int
	sortingChatByID(&*chat, *n)
	fmt.Print("ID riwayat mana yang ingin diedit:")
	fmt.Scan(&ID)
	target = binarySearchID(*chat, *n, ID)
	if target != -1 {
		chat[target].keyword = nil
		chat[target].saran = nil
		chat[target].urgensi = 0
		fmt.Println("Silakan masukkan ulang chat Anda:")
		chatsession(&*chat, target)
		chat[target].id = ID
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}
