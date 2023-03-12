package main

import "fmt"

func main() {
	// Buatlah sebuah program go dengan implementasi perulangan for dan kombinasi if-else

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
		if i >= 4 {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					const text = "САШАРВО"
					for index, runeValue := range text {
						fmt.Printf("character %#U starts at byte position %d\n", runeValue, index)
					}
				} else {
					fmt.Println("Nilai j =", j)
				}
			}
		}
	}
}

/*
https://www.tutorialspoint.com/go/go_nested_loops.htm
https://go.dev/blog/strings
*/
