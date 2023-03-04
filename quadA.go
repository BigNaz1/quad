package main

import "fmt"

func QuadA(height int, width int) int {
	for h := 1; h <= height; h++ {
		for w := 1; w <= width; w++ {
			if h == 1 || h == height {
				if w == 1 || w == width {
					fmt.Print("o")
				} else {
					fmt.Print("-")
				}

			} else if w == 1 || w == width {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return height + width
}

func main() {
	QuadA(3, 5)
}
