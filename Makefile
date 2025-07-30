# menjalankan aplikasi
run: 
	go run main.go

# Mengupdate dependensi
tidy:
	go mod tidy

# membuild aplikasi
# -s: Menghapus tabel simbol.
# -w: Menghapus informasi DWARF (debug).
compile:
	go build -ldflags="-s -w" -o build/go_contact_management