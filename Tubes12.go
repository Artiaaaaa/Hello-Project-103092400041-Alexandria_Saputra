package main // Tugas GIT Wawasan Global Alexandria Saputra

import (
	"fmt"
)

const NMAX = 100

type Barang struct {
	id                   string
	nama                 string
	hargabeli, hargajual int
	stok, terjual        int
}

type Transaksi struct {
	idbarang   string
	jumlah     int
	totalharga int
}

var dataBarang [NMAX]Barang
var dataTransaksi [NMAX]Transaksi
var nBarang, nTransaksi int

func main() {
	var pilih int
	for {
		fmt.Println("\n===============================")
		fmt.Println("          MENU UTAMA")
		fmt.Println("===============================")
		fmt.Println("1. Tambah data")
		fmt.Println("2. Edit data")
		fmt.Println("3. Hapus data")
		fmt.Println("4. Tampilkan data")
		fmt.Println("5. Cari barang")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuTambah()
		} else if pilih == 2 {
			menuEdit()
		} else if pilih == 3 {
			menuHapus()
		} else if pilih == 4 {
			menuTampil()
		} else if pilih == 5 {
			menuCari()
		} else if pilih == 6 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTambah() {
	var pilih int
	for {
		fmt.Println("Menu Tambah Data")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tambah Transaksi Penjualan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahBarang()
		} else if pilih == 2 {
			tambahTransaksi()
		} else if pilih == 0 {
			return // Kembali
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuEdit() {
	var pilih int
	for {
		fmt.Println("Menu Edit Data")
		fmt.Println("1. Edit Barang")
		fmt.Println("2. Edit Transaksi")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			editBarang()
		} else if pilih == 2 {
			editTransaksi()
		} else if pilih == 0 {
			return // Kembali
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuHapus() {
	var pilih int
	for {
		fmt.Println("Menu Hapus Data")
		fmt.Println("1. Hapus Barang")
		fmt.Println("2. Hapus Transaksi")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			hapusBarang()
		} else if pilih == 2 {
			hapusTransaksi()
		} else if pilih == 0 {
			return // Kembali
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTampil() {
	var pilih int
	for {
		fmt.Println("Menu Tampilkan Data")
		fmt.Println("1. Tampilkan Data Barang")
		fmt.Println("2. Tampilkan Data Transaksi")
		fmt.Println("3. Tampilkan Barang Terlaris")
		fmt.Println("4. Tampilkan Barang Tersedikit")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			var pilihSort int
			fmt.Println("\n  Mau pilih yang mana maszeh:")
			fmt.Println("  1. Ascending ")
			fmt.Println("  2. Descending ")
			fmt.Print("  Pilih: ")
			fmt.Scan(&pilihSort)
			if pilihSort == 1 {
				urutkanBarangByID(true) // true untuk ascending
			} else if pilihSort == 2 {
				urutkanBarangByID(false) // false untuk descending
			}
			cetakDaftarBarang()
		} else if pilih == 2 {
			var pilihSort int
			fmt.Println("\n  Pilih yang mana bos:")
			fmt.Println("  1. Ascending (Termurah)")
			fmt.Println("  2. Descending (Termahal)")
			fmt.Print("  Pilih: ")
			fmt.Scan(&pilihSort)
			if pilihSort == 1 {
				urutkanTransaksiByTotal(true) // true untuk ascending
			} else if pilihSort == 2 {
				urutkanTransaksiByTotal(false) // false untuk descending
			}
			cetakDaftarTransaksi()
		} else if pilih == 3 {
			cetakBarangTerlaris()
		} else if pilih == 4 {
			stoktersedikit()
		} else if pilih == 0 {
			return // Kembali
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuCari() {
	var pilih int
	for {
		fmt.Println("Menu Cari Barang")
		fmt.Println("1. Cari dengan Sequential Search")
		fmt.Println("2. Cari dengan Binary Search")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuCariBarangSequential()
		} else if pilih == 2 {
			menuCariBarangBinary()
		} else if pilih == 0 {
			return // Kembali
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func newBarang(id, nama string, hargabeli, hargajual, stok int) Barang {
	return Barang{
		id:        id,
		nama:      nama,
		hargabeli: hargabeli,
		hargajual: hargajual,
		stok:      stok,
		terjual:   0,
	}
}

func newTransaksi(idbarang string, jumlah, total int) Transaksi {
	return Transaksi{
		idbarang:   idbarang,
		jumlah:     jumlah,
		totalharga: total,
	}
}

func tambahBarang() {
	if nBarang < NMAX {
		var id, nama string
		var hargabeli, hargajual, stok int
		fmt.Println("Menambah Barang Baru")
		fmt.Print(" ID         : ")
		fmt.Scan(&id)
		fmt.Print(" Nama       : ")
		fmt.Scan(&nama)
		fmt.Print(" Harga Beli : ")
		fmt.Scan(&hargabeli)
		fmt.Print(" Harga Jual : ")
		fmt.Scan(&hargajual)
		fmt.Print(" Stok       : ")
		fmt.Scan(&stok)
		dataBarang[nBarang] = newBarang(id, nama, hargabeli, hargajual, stok)
		nBarang++
		fmt.Println("Barang baru berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas data barang sudah penuh.")
	}
}

func editBarang() {
	var id string
	fmt.Println("Edit Data Barang")
	fmt.Print("Masukkan ID barang yang ingin diedit: ")
	fmt.Scan(&id)
	idx := cariIndexBarangByID(id)
	if idx != -1 {
		fmt.Printf("Mengedit barang: %s\n", dataBarang[idx].nama)
		fmt.Print(" Nama baru       : ")
		fmt.Scan(&dataBarang[idx].nama)
		fmt.Print(" Harga beli baru : ")
		fmt.Scan(&dataBarang[idx].hargabeli)
		fmt.Print(" Harga jual baru : ")
		fmt.Scan(&dataBarang[idx].hargajual)
		fmt.Print(" Stok baru       : ")
		fmt.Scan(&dataBarang[idx].stok)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusBarang() {
	var id string
	fmt.Println("Hapus Data Barang")
	fmt.Print("Masukkan ID barang yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := cariIndexBarangByID(id)
	if idx != -1 {
		for i := idx; i < nBarang-1; i++ {
			dataBarang[i] = dataBarang[i+1]
		}
		nBarang--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cetakDaftarBarang() {
	fmt.Println("Daftar Semua Barang")
	if nBarang == 0 {
		fmt.Println("Belum ada data barang.")
		return
	}
	fmt.Printf("%-10s %-15s %-12s %-12s %-6s %-8s\n", "ID", "Nama", "HargaBeli", "HargaJual", "Stok", "Terjual")
	fmt.Println("--------------------------------------------------------------------")
	for i := 0; i < nBarang; i++ {
		b := dataBarang[i]
		fmt.Printf("%-10s %-15s %-12d %-12d %-6d %-8d\n", b.id, b.nama, b.hargabeli, b.hargajual, b.stok, b.terjual)
	}
	fmt.Println()
}

func tambahTransaksi() {
	var id string
	var jumlah int
	fmt.Println("Tambah Transaksi")
	fmt.Print("Masukkan ID barang yang dibeli: ")
	fmt.Scan(&id)
	idx := cariIndexBarangByID(id)
	if idx == -1 {
		fmt.Println("ID barang tidak ditemukan.")
		return
	}
	fmt.Printf("Masukkan jumlah yang dibeli (stok tersedia: %d): ", dataBarang[idx].stok)
	fmt.Scan(&jumlah)
	if jumlah <= 0 || jumlah > dataBarang[idx].stok {
		fmt.Println("Jumlah tidak valid atau melebihi stok.")
		return
	}
	total := dataBarang[idx].hargajual * jumlah
	dataTransaksi[nTransaksi] = newTransaksi(id, jumlah, total)
	nTransaksi++
	dataBarang[idx].stok -= jumlah
	dataBarang[idx].terjual += jumlah
	fmt.Println("Transaksi berhasil dicatat.")
}

func editTransaksi() {
	var idx int
	fmt.Println("edit Transaksi")
	if nTransaksi == 0 {
		fmt.Println("Belum ada transaksi untuk diedit.")
		return
	}
	cetakDaftarTransaksi()
	fmt.Print("Masukkan nomor transaksi yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx >= 0 && idx < nTransaksi {
		idBarangLama := dataTransaksi[idx].idbarang
		jumlahLama := dataTransaksi[idx].jumlah
		idxBarang := cariIndexBarangByID(idBarangLama)
		if idxBarang == -1 {
			fmt.Println("Error: Barang terkait transaksi tidak ditemukan.")
			return
		}
		dataBarang[idxBarang].stok += jumlahLama
		dataBarang[idxBarang].terjual -= jumlahLama
		var jumlahBaru int
		fmt.Printf("Mengedit transaksi untuk barang '%s'\n", dataBarang[idxBarang].nama)
		fmt.Printf("Masukkan jumlah baru (stok tersedia sekarang: %d): ", dataBarang[idxBarang].stok)
		fmt.Scan(&jumlahBaru)
		if jumlahBaru > 0 && jumlahBaru <= dataBarang[idxBarang].stok {
			dataTransaksi[idx].jumlah = jumlahBaru
			dataTransaksi[idx].totalharga = dataBarang[idxBarang].hargajual * jumlahBaru
			dataBarang[idxBarang].stok -= jumlahBaru
			dataBarang[idxBarang].terjual += jumlahBaru
			fmt.Println("Transaksi berhasil diubah.")
		} else {
			dataBarang[idxBarang].stok -= jumlahLama
			dataBarang[idxBarang].terjual += jumlahLama
			fmt.Println("Jumlah baru tidak valid. Perubahan dibatalkan.")
		}
	} else {
		fmt.Println("Nomor transaksi tidak valid.")
	}
}

func hapusTransaksi() {
	var idx int
	fmt.Println("Hapus Transaksi")
	if nTransaksi == 0 {
		fmt.Println("Belum ada transaksi untuk dihapus.")
		return
	}
	cetakDaftarTransaksi()
	fmt.Print("Masukkan nomor transaksi yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--
	if idx >= 0 && idx < nTransaksi {
		idBarangDihapus := dataTransaksi[idx].idbarang
		jumlahDihapus := dataTransaksi[idx].jumlah
		idxBarang := cariIndexBarangByID(idBarangDihapus)
		if idxBarang != -1 {
			dataBarang[idxBarang].stok += jumlahDihapus
			dataBarang[idxBarang].terjual -= jumlahDihapus
		}
		for i := idx; i < nTransaksi-1; i++ {
			dataTransaksi[i] = dataTransaksi[i+1]
		}
		nTransaksi--
		fmt.Println("Transaksi berhasil dihapus.")
	} else {
		fmt.Println("Nomor transaksi tidak valid.")
	}
}

func cetakDaftarTransaksi() {
	fmt.Println("Daftar Transaksi")
	if nTransaksi == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	fmt.Printf("%-5s %-10s %-15s %-10s %-12s\n", "No.", "IDBarang", "Nama Barang", "Jumlah", "Total Harga")
	fmt.Println("----------------------------------------------------------")
	for i := 0; i < nTransaksi; i++ {
		t := dataTransaksi[i]
		idxBarang := cariIndexBarangByID(t.idbarang)
		namaBarang := "N/A"
		if idxBarang != -1 {
			namaBarang = dataBarang[idxBarang].nama
		}
		fmt.Printf("%-5d %-10s %-15s %-10d %-12d\n", i+1, t.idbarang, namaBarang, t.jumlah, t.totalharga)
	}
	fmt.Println()
}

func cariIndexBarangByID(id string) int {
	for i := 0; i < nBarang; i++ {
		if dataBarang[i].id == id {
			return i
		}
	}
	return -1
}

func menuCariBarangSequential() {
	var nama string
	fmt.Println("Cari Barang")
	fmt.Print("Masukkan Nama Barang yang dicari: ")
	fmt.Scan(&nama)
	fmt.Println("Hasil Pencarian")
	found := false
	fmt.Printf("%-10s %-15s %-12s %-12s %-6s %-8s\n", "ID", "Nama", "HargaBeli", "HargaJual", "Stok", "Terjual")
	fmt.Println("--------------------------------------------------------------------")
	for i := 0; i < nBarang; i++ {
		if dataBarang[i].nama == nama {
			b := dataBarang[i]
			fmt.Printf("%-10s %-15s %-12d %-12d %-6d %-8d\n", b.id, b.nama, b.hargabeli, b.hargajual, b.stok, b.terjual)
			found = true
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}
	fmt.Println()
}

func datasementara(opo *[NMAX]Barang, n int) { // selection
	var temp Barang
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for i := pass + 1; i < n; i++ {
			if opo[i].nama < opo[idx].nama {
				idx = i
			}
		}
		temp = opo[pass]
		opo[pass] = opo[idx]
		opo[idx] = temp
	}
}

func menuCariBarangBinary() {
	var nama string
	fmt.Println("Cari Barang")
	if nBarang == 0 {
		fmt.Println("Data barang kosong, tidak bisa melakukan pencarian.")
		return
	}

	var dataUrut [NMAX]Barang
	for i := 0; i < nBarang; i++ {
		dataUrut[i] = dataBarang[i]
	}

	datasementara(&dataUrut, nBarang)

	fmt.Print("Masukkan Nama Barang yang dicari (nama harus persis): ")
	fmt.Scan(&nama)

	low := 0
	high := nBarang - 1
	found := false
	for low <= high {
		mid := low + (high-low)/2
		if dataUrut[mid].nama == nama {
			fmt.Println("Barang Ditemukan")
			fmt.Printf("%-10s %-15s %-12s %-12s %-6s %-8s\n", "ID", "Nama", "HargaBeli", "HargaJual", "Stok", "Terjual")
			fmt.Println("--------------------------------------------------------------------")
			b := dataUrut[mid]
			fmt.Printf("%-10s %-15s %-12d %-12d %-6d %-8d\n", b.id, b.nama, b.hargabeli, b.hargajual, b.stok, b.terjual)
			found = true
			break
		} else if dataUrut[mid].nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}
	fmt.Println()
}

func urutkanBarangByID(asc bool) { // Insertion
	for i := 1; i < nBarang; i++ {
		temp := dataBarang[i]
		j := i - 1
		if asc { // Ascending
			for j >= 0 && dataBarang[j].id > temp.id {
				dataBarang[j+1] = dataBarang[j]
				j--
			}
		} else { // Descending
			for j >= 0 && dataBarang[j].id < temp.id {
				dataBarang[j+1] = dataBarang[j]
				j--
			}
		}
		dataBarang[j+1] = temp
	}
}

func urutkanTransaksiByTotal(asc bool) {
	// Selection
	for i := 0; i < nTransaksi-1; i++ {
		extIdx := i
		for j := i + 1; j < nTransaksi; j++ {
			if asc {
				if dataTransaksi[j].totalharga < dataTransaksi[extIdx].totalharga {
					extIdx = j
				}
			} else {
				if dataTransaksi[j].totalharga > dataTransaksi[extIdx].totalharga {
					extIdx = j
				}
			}
		}
		temp := dataTransaksi[i]
		dataTransaksi[i] = dataTransaksi[extIdx]
		dataTransaksi[extIdx] = temp
	}
}

func cetakBarangTerlaris() {
	fmt.Println("Laporan Barang Terlaris")
	if nBarang == 0 {
		fmt.Println("Belum ada data barang.")
		return
	}

	var dataUrut [NMAX]Barang
	for i := 0; i < nBarang; i++ {
		dataUrut[i] = dataBarang[i]
	}

	for i := 0; i < nBarang-1; i++ { // selection
		maxIdx := i
		for j := i + 1; j < nBarang; j++ {
			if dataUrut[j].terjual > dataUrut[maxIdx].terjual {
				maxIdx = j
			}
		}
		temp := dataUrut[i]
		dataUrut[i] = dataUrut[maxIdx]
		dataUrut[maxIdx] = temp
	}

	fmt.Printf("%-5s %-10s %-15s %-8s\n", "Rank", "ID", "Nama", "Terjual")
	fmt.Println("-------------------------------------------")

	for i := 0; i < 5; i++ {
		if i < nBarang {
			if dataUrut[i].terjual > 0 {
				b := dataUrut[i]
				fmt.Printf("%-5d %-10s %-15s %-8d\n", i+1, b.id, b.nama, b.terjual)
			}
		}
	}
	fmt.Println()
}

func stoktersedikit() {
	fmt.Println("Laporan Barang Tersedikit")
	if nBarang == 0 {
		fmt.Println("Belum ada data barang.")
		return
	}

	var dataUrut [NMAX]Barang
	for i := 0; i < nBarang; i++ {
		dataUrut[i] = dataBarang[i]
	}

	for i := 0; i < nBarang-1; i++ { // selection
		maxIdx := i
		for j := i + 1; j < nBarang; j++ {
			if dataUrut[j].stok > dataUrut[maxIdx].stok {
				maxIdx = j
			}
		}
		temp := dataUrut[i]
		dataUrut[i] = dataUrut[maxIdx]
		dataUrut[maxIdx] = temp
	}

	fmt.Printf("%-5s %-10s %-15s %-8s\n", "Rank", "ID", "Nama", "Stok")
	fmt.Println("-------------------------------------------")

	for i := 0; i < 5; i++ {
		if i < nBarang {
			if dataUrut[i].terjual > 0 {
				b := dataUrut[i]
				fmt.Printf("%-5d %-10s %-15s %-8d\n", i+1, b.id, b.nama, b.stok)
			}
		}
	}
	fmt.Println()
}
