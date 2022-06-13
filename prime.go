package main

import "fmt"

func main() {

	fmt.Println("prime number between 1 to 100 is :")
	for j := 1; j <= 100; j++ {
		k := 0
		for i := 1; i <= j; i++ {

			if j%i == 0 {
				k++
			}
		}
		if k == 2 {
			fmt.Println(j)
		}

	}
}
