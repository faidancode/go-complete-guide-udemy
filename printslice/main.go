package main

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range number {
		if number%2 == 0 {
			println(number, "is even")
		} else {
			println(number, "is odd")
		}
	}
}
