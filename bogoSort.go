package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/hajimehoshi/oto"
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

func BogoSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, players []oto.Player, f int, c *oto.Context) {
	for {
		p := play(c, mapToFeq(int(data[rand.Intn(len(bars))]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, f)
		players = append(players, p)
		Sleep(info.delay)
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
		Visualize(win, bars, barWidth, data, -1, -1, info)
		Sleep(info.delay)
		if checkSortedArray(data) {
			break
		}
	}
	VisualizeSorted(win, bars, barWidth, data, players, f, c)
}
