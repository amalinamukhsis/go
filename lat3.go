/*buatlah sebuah program yang menerima input nama, umur, status KTP
1. apabila umur < 17 tahun, tidak bisa menggunakan layanan
2. apabila status ktp "t" maka buat pertanyaan apakah ybs mau membuat ktp
3. apabila status ktp "ya" say thanks
4. input NIK harus 16 digit



""" Contoh output 1
Nama            : Lala
Umur            : 90
Punya KTP (y/t) : y
NIK             : 1234567890123456
Terima kasih sudah menggunakan layanan kami
"""
""" Contoh output 2
Nama            : Lili
Umur            : 90
Punya KTP (y/t) : 90
Input status KTP anda tidak tepat
Terima kasih sudah menggunakan layanan kami
"""
""" Contoh output 3
Nama            : Lulu
Umur            : 90
Punya KTP (y/t) : Y
NIK             : 10
NIK yang anda masukkan tidak valid
Terima kasih sudah menggunakan layanan kami
"""
""" Contoh output 4
Nama            : lele
Umur            : 90
Punya KTP (y/t) : t
Apakah anda ingin membuat KTP?(y/t) Y
Silahkan lanjutkan ke layanan pembuatan KTP
Terima kasih sudah menggunakan layanan kami
"""
""" Contoh output 5
Nama            : kaka
Umur            : 90
Punya KTP (y/t) : T
Apakah anda ingin membuat KTP?(y/t) t
Terima kasih sudah menggunakan layanan kami
"""
""" Contoh output 6
Nama            : kiki
Umur            : 10
Layanan ini hanya untuk orang yang berumur 17 keatas
*/

package main

import "fmt"

func main() {
	var nama, status string
	var umur, NIK int
	var pilihan string
	fmt.Print("nama :")
	fmt.Scan(&nama)
	fmt.Print("umur :")
	fmt.Scan(&umur)
	if umur > 17 {
		fmt.Print("punya KTP (y/t) :")
		fmt.Scan(&status)
		if status == "y" || status == "Y" { //apa beda =(assigned nilai) dan ==(kondisi) ?
			fmt.Print("NIK : ")
			fmt.Scan(&NIK)
			if len(fmt.Sprintf("%d", NIK)) == 16 {
				fmt.Println("terimakasih sudah menggunakan layanan ini")
			} else {
				fmt.Println("NIK yang anda masukkan tidak valid")
			}
		} else if status == "t" || status == "T" {
			fmt.Print("Apakah anda ingin membuat KTP?(y/t) :")
			fmt.Scan(&pilihan)
			if pilihan == "y" || pilihan == "Y" {
				fmt.Println("Silahkan lanjutkan ke layanan pembuatan KTP")
			} else if pilihan == "t" || pilihan == "T" {
				fmt.Println("Terima kasih sudah menggunakan layanan kami")
			} else {
				fmt.Println("Input status KTP anda tidak tepat")
			}
		} else {
			fmt.Println("Input status anda tidak tepat")
		}
	} else {
		fmt.Println("Layanan ini hanya untuk orang yang berumur 17 keatas")
	}
}
