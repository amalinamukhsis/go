/*buatlah sebuah program menentukan apakah input bisa dibagi 2 dan dibagi 3, bisa dibagi 3 tapi tidak bisa dibagi 2, dan tidak bisa dibagi 2 dan 3

contoh 1 :
input : 15
output :bilangan 15 bisa dibagi 3 tapi tidak bisa dibagi 2

input : 8
output : bilangan 8 bisa dibagi 2 tapi tidak bisa dibagi 3

input : 12
output bilangan 12 bisa dibagi 2 dan 3

input : 5
output : bilangan tidak bisa dibagi 2 dan 3*/

package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	if input%2 == 0 {
		if input%3 == 0 {
			fmt.Printf("bilangan %d bisa dibagi 2 dan 3\n", input)
		} else {
			fmt.Printf("bilangan %d bisa dibagi 2 tapi tidak bisa dibagi 3\n", input)
		}
	} else {
		if input%3 == 0 {
			fmt.Printf("bilangan %d bisa dibagi 3 tapi tidak bisa dibagi 2\n", input)
		} else {
			fmt.Printf("bilangan %d tidak bisa dibagi 2 dan 3\n", input)
		}
	}
}
