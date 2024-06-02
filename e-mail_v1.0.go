package main

import "fmt"

const maxUserAndInbox int = 5 // agar mudah didemonstrasikan diset jangan banyak banyak

type user struct {
	alamat, passUser string
	approved         bool
	inbox            [maxUserAndInbox]emails
	jumlahEmail      int
}

type emails struct {
	from, to, subject, body string
}

var tabUser [maxUserAndInbox]user
var jumlahUser int
var passAdmin string = "pAdmin" // ini passwordnya memang dihardcode saja

func main() {
	var opsi string // milih gini mending pake string biar kalo keinput karakter apa saja akan keluar tidak valid
	jumlahUser = 0
	var keluarAplikasi bool = false // kenapa gak boleh pake break jadinya harus define boolean gini terus (T_T)

	for !keluarAplikasi {
		fmt.Println("--- SELAMAT DATANG DI APLIKASI EMAIL-EMAILAN! ---")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login sebagai user")
		fmt.Println("3. Login sebagai admin")
		fmt.Println("4. Keluar aplikasi")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&opsi)

		switch opsi {
		case "1":
			registrasiUser()
		case "2":
			loginUser()
		case "3":
			loginAdmin()
		case "4":
			fmt.Println("Berhasil keluar dari aplikasi. Bye bye!")
			keluarAplikasi = true
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

// FUNGSI-FUNGSI LOGIN/REGISTRASI USER

func loginUser() {
	var emailnya, passwordnya string
	var loopingLoginUser bool
	var x int
	loopingLoginUser = true

	for loopingLoginUser != false {
		fmt.Println("--- HALAMAN LOGIN ---")
		fmt.Print("Masukkan alamat email: ")
		fmt.Scanln(&emailnya)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&passwordnya)
		x = cariEmail(emailnya)

		if x >= 0 {
			if tabUser[x].approved != true { // ini ngecek kalo emailnya udah di approve atau belum sama adminnya
				fmt.Println("Alamat email belum disetujui admin")
				loopingLoginUser = false
			} else {
				if tabUser[x].passUser == passwordnya {
					fmt.Println("Berhasil login.")
					menuUser(&tabUser[x], x)
					loopingLoginUser = false
				} else {
					fmt.Println("Password salah")
					loopingLoginUser = false
				}
			}
		} else {
			fmt.Println("Alamat email tidak ditemukan.")
			loopingLoginUser = false
		}
	}
}

func menuUser(pengguna *user, nomor int) {
	var pilihan string
	var loopingMenuUser bool = true

	for loopingMenuUser != false {
		fmt.Println("1. Kirim Email")
		fmt.Println("2. Lihat Inbox")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			kirimEmail(pengguna)
		case "2":
			menuInbox(&*pengguna, nomor)
		case "3":
			loopingMenuUser = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func registrasiUser() {
	var alamatEmail, password string
	var loopingRegistrasiUser bool = true
	var loopingLuar bool = true

	for loopingLuar != false {
		if jumlahUser >= maxUserAndInbox {
			fmt.Println("Kapasitas user sudah penuh. Coba lagi untuk registrasi nanti")
			loopingLuar = false // nah ini kapasitasnya itu yang jumlahUser
		} else {
			fmt.Println("--- HALAMAN REGISTRASI ---")

			for loopingRegistrasiUser {
				fmt.Print("Masukkan alamat email: ")
				fmt.Scanln(&alamatEmail)
				if cariEmail(alamatEmail) >= 0 {
					fmt.Println("Alamat email ini sudah digunakan, silakan coba yang lain.")
				} else {
					loopingRegistrasiUser = false
				}
			}

			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			userBaru := user{alamat: alamatEmail, passUser: password, approved: false}
			tabUser[jumlahUser] = userBaru
			jumlahUser++
			sortTabUser()
			fmt.Println("Registrasi berhasil. Tunggu persetujuan admin.")
			loopingLuar = false
		}
	}
}

// FUNGSI-FUNGSI EMAIL

func kirimEmail(pengirim *user) {
	var ke, subjek string
	var x int = -1

	fmt.Println("--- HALAMAN PENGIRIMAN EMAIL ---")

	for x < 0 {
		fmt.Print("Masukkan alamat email penerima: ")
		fmt.Scanln(&ke)
		x = cariEmail(ke)
		if x < 0 {
			fmt.Println("Alamat email penerima tidak ditemukan. Coba lagi.")
		}
	}

	fmt.Print("Masukkan subjek email: ")
	fmt.Scanln(&subjek)
	buatDanBalasEmail(&*pengirim, x, ke, subjek)
}

func menuInbox(pengguna *user, x int) {
	var pilihanEmail int
	var loopingMenuInbox bool = true

	for loopingMenuInbox {
		fmt.Println("--- INBOX ---")
		for i := 0; i < pengguna.jumlahEmail; i++ {
			fmt.Print(i+1, ". ", pengguna.inbox[i].subject, " dari ", pengguna.inbox[i].from)
			fmt.Println()
		}
		fmt.Println("(Pilih 0 untuk keluar dan pilih -1 untuk sort inbox dengan subjek)")
		fmt.Print("Pilih email: ")
		fmt.Scanln(&pilihanEmail)

		switch pilihanEmail {
		case 0:
			loopingMenuInbox = false
		case -1:
			sortInbox(&*pengguna)
		default:
			if pilihanEmail < 0 {
				fmt.Println("Tidak ada email tersebut.")
			} else if pilihanEmail > pengguna.jumlahEmail {
				fmt.Println("Tidak ada email tersebut.")
			} else {
				manipulasiEmail(&*pengguna, pilihanEmail-1)
			}
		}
	}
}

func manipulasiEmail(pengguna *user, nomor int) {
	var pilih string
	var loopingManinpulasiEmail bool = true

	for loopingManinpulasiEmail != false {
		fmt.Println("--- MELIHAT EMAIL ---")
		fmt.Print("Email dengan subjek: ", pengguna.inbox[nomor].subject, ", dari: ", pengguna.inbox[nomor].from)
		fmt.Println()
		fmt.Println("1. Baca")
		fmt.Println("2. Balas")
		fmt.Println("3. Hapus")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			fmt.Println("Subjek:", pengguna.inbox[nomor].subject)
			fmt.Println("Dari:", pengguna.inbox[nomor].from)
			fmt.Println("Isi:", pengguna.inbox[nomor].body)
		case "2":
			ke := pengguna.inbox[nomor].from
			subjek := "reply"
			buatDanBalasEmail(&*pengguna, nomor, ke, subjek)
		case "3":
			var kosongkan emails
			pengguna.inbox[nomor] = kosongkan
			pengguna.jumlahEmail -= 1
			for i := nomor; i < pengguna.jumlahEmail; i++ {
				pengguna.inbox[i] = pengguna.inbox[i+1]
			}
			fmt.Println("Email berhasil dihapus. Kembali ke list inbox...")
			loopingManinpulasiEmail = false
		case "4":
			fmt.Println("Kembali ke list inbox...")
			loopingManinpulasiEmail = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func buatDanBalasEmail(pengirim *user, nomor int, ke, subjek string) { // untuk membalas dan ngirim email
	var isi string

	fmt.Print("Masukkan isi email: ")
	fmt.Scanln(&isi)
	email := emails{from: pengirim.alamat, to: ke, subject: subjek, body: isi}

	if tabUser[nomor].jumlahEmail < maxUserAndInbox {
		tabUser[nomor].jumlahEmail += 1
	}
	for i := maxUserAndInbox - 1; i > 0; i-- { // ini jadi yang email sebelumnya dishift ke nomor selanjutnya
		tabUser[nomor].inbox[i] = tabUser[nomor].inbox[i-1]
	}
	tabUser[nomor].inbox[0] = email
}

// FUNGSI-FUNGSI ADMIN

func loginAdmin() {
	var password string

	fmt.Print("Masukkan password admin: ")
	fmt.Scanln(&password)
	if password == passAdmin /*ini password yg dihardcode*/ {
		menuAdmin()
	} else {
		fmt.Println("Password salah. Kembali ke halaman...")
	}
}

func menuAdmin() {
	var pilihannya string
	var kembali bool = false

	for kembali != true {
		fmt.Println("--- HALAMAN ADMIN ---")
		fmt.Println("1. List pengguna")
		fmt.Println("2. Setujui/Tolak Registrasi Pengguna")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihannya)
		switch pilihannya {
		case "1":
			listPenggunaApproved()
		case "2":
			approveRegistration()
		case "3":
			kembali = true
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func listPenggunaApproved() {
	fmt.Println("--- HALAMAN LIST PENGGUNA ---")
	i := 0
	nomor := 0

	fmt.Println("Pengguna sudah disetujui:")
	for i < jumlahUser {
		if tabUser[i].approved == true {
			nomor += 1
			fmt.Print(nomor, ". ", tabUser[i].alamat)
			fmt.Println()
		}
		i++
	}
	if nomor == 0 {
		fmt.Println("Tidak ada pengguna sudah disetujui")
	}

	j := 0
	nomor = 0

	fmt.Println("Pengguna belum disetujui:")
	for j < jumlahUser {
		if tabUser[j].approved == false {
			nomor += 1
			fmt.Print(nomor, ". ", tabUser[j].alamat)
			fmt.Println()
		}
		j++
	}
	if nomor == 0 {
		fmt.Println("Tidak ada pengguna belum disetujui")
	}

	fmt.Println("Jumlah user (sudah dan belum di approve):", jumlahUser)
	fmt.Println("Kembali ke halaman admin...")
}

func approveRegistration() {
	fmt.Println("--- HALAMAN PERSETUJUAN REGISTRASI ---")
	fmt.Println("List pengguna belum disetujui:")

	nomor := 0
	for j := 0; j < jumlahUser; j++ {
		if tabUser[j].approved == false {
			nomor += 1
			fmt.Print(nomor, ". ", tabUser[j].alamat)
			fmt.Println()
		}
	}

	if nomor == 0 {
		fmt.Println("Tidak ada pendaftar baru. Kembali ke halaman admin...")
	} else {
		i := 0
		for i < jumlahUser {
			if tabUser[i].approved != true {
				var persetujuan string

				fmt.Print("Setujui pengguna ", tabUser[i].alamat, "?")
				fmt.Println()
				fmt.Println("1. Setuju, 2. Tolak, 3. Skip")
				fmt.Print("Pilih opsi: ")
				fmt.Scanln(&persetujuan)

				switch persetujuan {
				case "1":
					tabUser[i].approved = true
					fmt.Println("Pengguna disetujui.")
				case "2":
					var kosongkan user
					tabUser[i] = kosongkan
					if jumlahUser > 0 {
						jumlahUser--
					}
					for j := i; j < jumlahUser; j++ {
						tabUser[j] = tabUser[j+1]
					}
					i = i - 1
					fmt.Println("Pengguna ditolak.")
				case "3":
					fmt.Println("Pengguna diskip dahulu.")
				default:
					fmt.Println("Opsi tidak valid.")
				}
			}
			i++
		}
	}
}

// FUNGSI-FUNGSI LAIN RANDOM

func sortTabUser() { // sorting alamat email pengguna secara alfabet
	for i := 0; i < jumlahUser-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahUser; j++ {
			if tabUser[j].alamat < tabUser[minIndex].alamat {
				minIndex = j
			}
		}
		if minIndex != i {
			tabUser[i], tabUser[minIndex] = tabUser[minIndex], tabUser[i]
		}
	}
}

func sortInbox(pengguna *user) { // sorting inbox secara alfabet berdasarkan subjek
	for i := 1; i < pengguna.jumlahEmail; i++ {
		tumbal := pengguna.inbox[i]
		j := i - 1
		for j >= 0 && pengguna.inbox[j].subject > tumbal.subject {
			pengguna.inbox[j+1] = pengguna.inbox[j]
			j = j - 1
		}
		pengguna.inbox[j+1] = tumbal
	}
}

func cariEmail(alamatEmail string) int { // mengeluarkan angka untuk arraynya
	var x int = -1

	for i := 0; i < jumlahUser; i++ {
		if tabUser[i].alamat == alamatEmail {
			x = i
		}
	}
	return x
}
