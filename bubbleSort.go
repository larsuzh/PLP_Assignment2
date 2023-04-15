package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func BubbleSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
			info.comparisons = i*len(data) + j
			Visualize(win, bars, barWidth, data, j, j+1, info)
			Sleep(info.delay)
		}
	}
}
