package main

/*

3. Tipe Data String
	-- String adalah tipe data kumpulan karakter
	-- Direpresentasikan dengan kata kunci string
	-- Diawali dengan karakter “ (petik dua) dan diakhiri dengan karakter “ (petik dua)

	Function yang terdapat pada String:
	-- len("string") = Menghitung jumlah karakter di String
	-- “string”[number] = Mengambil karakter pada posisi yang ditentukan
*/

func main() {
	println("Ahmad")
	println("Ahmad Mufied")
	println("Ahmad Mufied Nugroho")

	println(len("Ahmad"))
	println("Ahmad Mufied"[0])
	println("Ahmad Mufied Nugroho"[1])
}