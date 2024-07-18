/*buatlah sebuah program yang menerima 3 buah input bilangan bulat, lalu bandingkaan bilangan tersebut dan keluarkan angka tertinggi. gunakan nested if

input : 23 56 78
output : digit-3 adalah nilai maximum

input : 45 90 88
output : digit-2 adalah nilai maximum

NB : input user ke samping
*/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > c {
		if a > c {
			fmt.Println(a)
		} else {
			fmt.Println(c)
		}
	} else {
		if b > c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
	}

}
