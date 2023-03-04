package main

import "fmt"

func QuadB(x int, y int) int {
	for h := 1; h <= x; h++ {
		for w := 1; w <= y; w++ {
			if h == 1 || h == x {
				if w == 1 && h == 1 {
					fmt.Print("/")
				} else if w == y && h == x {
					fmt.Print("/")
				} else if w == y && h == 1 {
					fmt.Print("\\")
				} else if w == 1 && h == x {
					fmt.Print("\\")
				} else {
					fmt.Print("*")
				}
			} else if w == 1 || w == y {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return x + y
}

func main() {
	QuadB(3, 5)
}
