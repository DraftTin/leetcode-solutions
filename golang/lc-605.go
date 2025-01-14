package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			left := (i == 0) || (flowerbed[i-1] == 0)
			right := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
			if left && right {
				count++
				flowerbed[i] = 1
			}
		}
	}
	fmt.Println(count)
	return count >= n
}
