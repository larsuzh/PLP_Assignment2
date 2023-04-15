package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func BubbleSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, delay int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
			Visualize(win, bars, barWidth, data)
			Sleep(delay)
		}
	}
}
