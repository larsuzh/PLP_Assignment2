package main

import (
	"math/rand"

	"github.com/faiface/pixel/pixelgl"
)

func checkSortedArray(arr []float64) bool {
	sortedArray := true
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				sortedArray = false
				break
			}
		}
	}
	return sortedArray
}

func BogoSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, delay int) {
	for {
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
		Visualize(win, bars, barWidth, data)
		Sleep(delay)
		if checkSortedArray(data) {
			return
		}
	}
}