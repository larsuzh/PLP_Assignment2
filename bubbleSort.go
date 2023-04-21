package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/hajimehoshi/oto"
)

func BubbleSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, players []oto.Player, f int, c *oto.Context) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
			p := play(c, mapToFeq(int(data[j]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, f)
			players = append(players, p)
			Sleep(info.delay)
			info.comparisons = i*len(data) + j
			Visualize(win, bars, barWidth, data, j, j+1, info)
			Sleep(info.delay)
		}
	}
	VisualizeSorted(win, bars, barWidth, data, players, f, c)
}
