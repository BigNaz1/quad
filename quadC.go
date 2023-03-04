package main

import "fmt"

func QuadC(x int, y int) int {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if h == 1 || h == y {
				if w == 1 && h == 1 {
					fmt.Print("A")
				} else if w == x && h == y && h != 1 {
					fmt.Print("C")
				} else if w == x && h == 1 {
					fmt.Print("A")
				} else if w == 1 && h == y {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			} else if w == 1 || w == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return x + y
}

func main() {
	QuadC(5, 3)
}
