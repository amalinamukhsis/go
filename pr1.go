/*buatlah sebuah program untuk menetukan apakah tahun input user kabisat atau tidak

input : 2000
output : kabisat

input : 2012
output : 2012 tahun kabisat

input : 2100
output : 2100  bukan tahun kabisat
*/

package main

import "fmt"

//program mengecek input tahun kabisat
func main() {
	var input int
	fmt.Scan(&input)
	if input%400 == 0 {
		fmt.Println("tahun kabisat")
	} else if input%100 == 0 {
		fmt.Println("tidak kabisat")
	} else if input%4 == 0 {
		fmt.Println("tahun kabisat")
	} else {
		fmt.Println("tidak kabisat")
	}
}
