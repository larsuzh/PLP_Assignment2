package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/hajimehoshi/oto"
)

func SelectionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, players []oto.Player, f int, c *oto.Context) {
	var n = len(data)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
			info.comparisons = i*len(data) + j - i
			p := play(c, mapToFeq(int(data[j]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, f)
			players = append(players, p)
			Sleep(info.delay)
			Visualize(win, bars, barWidth, data, minIdx, j, info)
			Sleep(info.delay)
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	VisualizeSorted(win, bars, barWidth, data, players, f, c)
}
