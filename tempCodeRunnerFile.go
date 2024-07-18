package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	if input%4 == 0 {
		fmt.Println("tahun kabisat")
	} else if input%400 == 0 {
		if input%100 == 0 {
			fmt.Println("tidak kabisat")
		} else {
			fmt.Println("kabisat")
		}
	} else {
		fmt.Println("tidak kabisat")
	}
}
