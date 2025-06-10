Chatbot Kesehatan Mental & Laporan Produktivitas (Go)
Program ini adalah chatbot berbasis konsol yang ditulis dalam bahasa Go, dirancang untuk memberikan dukungan kesehatan mental serta laporan produktivitas sederhana berdasarkan input pengguna. Chatbot ini mampu mendeteksi kata kunci dari percakapan pengguna dan memberikan saran aktivitas berdasarkan tingkat urgensi emosional yang terdeteksi.

Fitur Utama
1. Mode Kesehatan Mental
   - Mendeteksi kata kunci yang terkait dengan masalah mental dengan tingkat urgensi rendah maupun tinggi.
   - Memberikan saran aktivitas secara acak berdasarkan kata kunci yang terdeteksi.
   - Menyimpan percakapan untuk referensi di masa depan.
2. Laporan Produktivitas
   - Menyusun semua saran yang pernah diberikan selama sesi konsultasi.
   - Menghindari duplikasi saran yang sama.
3. Riwayat Chat
   - Melihat semua riwayat percakapan dan saran.
   - Mengurutkan riwayat berdasarkan urgensi atau waktu (ID percakapan).
   - Menghapus dan mengedit riwayat percakapan tertentu berdasarkan ID.


Struktur Data
  - history: Menyimpan informasi percakapan, termasuk input pengguna, kata kunci yang ditemukan, saran aktivitas, ID percakapan, dan tingkat urgensi.
  - arrChat: Array dari 100 elemen history untuk menyimpan riwayat percakapan.

Logika Deteksi
  - Kata-kata yang dianggap berurgensi tinggi (contoh: "bunuh_diri", "putus_asa") akan memberikan saran tindakan segera dan serius.
  - Kata-kata berurgensi rendah (contoh: "stres", "cemas") akan memberikan saran aktivitas santai dan perawatan diri.

Cara Menggunakan
  1. Jalankan program.
  2. Pilih salah satu dari tiga menu utama:
      [1] Kesehatan Mental
      [2] Laporan Produktivitas
      [3] Riwayat Chat

  3. Ketik kalimat (diakhiri dengan titik .) saat konsultasi kesehatan mental.
  4. Gunakan menu Riwayat untuk melihat, menyortir, menghapus, atau mengedit percakapan.

Contoh Percakapan
Selamat datang di Chatbot AI Go!
Pilih kategori:
[1] Kesehatan Mental
[2] Laporan Produktivitas
[3] Riwayat chat
Ketik 'exit' untuk keluar.
  > 1
Mode Kesehatan Mental. Ketik 'menu' untuk kembali.
jika ingin melanjutkan ketik apa saja
  > lanjut
Anda : saya merasa stres dan insomnia.
[Bot memberi saran berdasarkan kata 'stres' dan 'insomnia']

Fitur Tambahan
  - Sorting ID (ascending) menggunakan Selection Sort.
  - Sorting urgensi (descending) menggunakan Insertion Sort.
  - Binary Search digunakan untuk mencari ID dalam operasi edit/hapus.

Catatan Teknis
  - Input kalimat harus diakhiri dengan titik (.) agar dikenali sebagai akhir dari satu percakapan.
  - Saat mengedit riwayat, sistem akan menghapus keyword dan saran sebelumnya sebelum mengisi yang baru.
  - Semua input dikonversi ke huruf kecil untuk mempermudah pencocokan kata kunci.

