/*buatlah sebuah program yang menerima input bilangan bulat, tentukan bilangan tersebut ganjil atau genap

contoh 1 :

masukkan input : 65
bilangan 65 adalah ganjil

contoh 2 :
masukkan input : -45
input harus bilangan positive

contoh 3 :
masukkan input : 0
bilangan 0 tidak ganjil dan tidak genap

*/

package main

import "fmt"

//nested iif
func main() {
	var input int
	fmt.Println("masukkan input")
	fmt.Scan(&input)
	//cek input positive negative atau 0
	if input > 0 { // bilangan positive
		if input%2 == 0 { // jika input dibagi 2 sisa 0 maka genap
			fmt.Printf("bilangan %d adalah genap\n", input) // string printf
		} else {
			fmt.Printf("bilangan %d adalah ganjil\n", input)
		}
	} else if input == 0 { //kondisi input 0
		fmt.Printf("ilangan 0 tidak ganjil dan tidak genap")
	} else { //kondisi input negative
		fmt.Printf("input harus bilangan positif")
	}
}

//type data
// %d ---> diskrit = integer
// %f --> float
// %s ---> string

// operator ( ka ba ta ku)
/*
 + * / - %
 beda /(pembagian hasilnya float) dan //(integer)
 5/2 = 2.5     5//2 = 2, 5 div 2
 6 % 5 = 1
 78 % 10 = 8 ,78 mod 10

*/
