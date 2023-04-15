package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func SelectionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, delay int) {
	var n = len(data)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
			Visualize(win, bars, barWidth, data)
			Sleep(delay)
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}
