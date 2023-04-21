package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/hajimehoshi/oto"
)

func InsertionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, players []oto.Player, f int, c *oto.Context) {
	var n = len(data)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
			info.comparisons = int(i/2*i) + i - j
			p := play(c, mapToFeq(int(data[j]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, f)
			players = append(players, p)
			Sleep(info.delay)
			Visualize(win, bars, barWidth, data, j, j-1, info)
			Sleep(info.delay)
		}
	}
	VisualizeSorted(win, bars, barWidth, data, players, f, c)
}
