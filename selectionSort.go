package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func SelectionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info) {
	var n = len(data)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
			info.comparisons = i*len(data) + j - i
			Visualize(win, bars, barWidth, data, minIdx, j, info)
			Sleep(info.delay)
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	VisualizeSorted(win, bars, barWidth, data)
}
