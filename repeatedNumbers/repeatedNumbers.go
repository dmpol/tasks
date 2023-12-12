package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100
	fmt.Println("myFunc1: ", myFunc1(num))
	fmt.Println("myFunc2: ", myFunc2(num))

}

func myFunc1(num int) int {
	count := 0

	for i := 1; i <= num; i++ {
		m := make(map[string]int)
		str := strconv.Itoa(i)

		for j := 0; j < len(str); j++ {

			if _, ok := m[string(str[j])]; ok {
				count++
				break
			}
			m[string(str[j])]++
		}
	}
	return count
}

func myFunc2(num int) int {

	var count, k int
	var n byte

	for i := 1; i <= num; i++ {
		m := make(map[byte]byte)
		k = i
		n = 0

		for k != 0 {
			n = (byte)(k % 10)
			k = (k - (int)(n)) / 10

			if _, ok := m[n]; ok {
				count++
				break
			}
			m[n]++
		}
	}
	return count
}
