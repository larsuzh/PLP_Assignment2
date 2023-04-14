package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func InsertionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, delay int) {
	var n = len(data)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
		Visualize(win, bars, barWidth, data)
		Sleep(delay)
	}
}