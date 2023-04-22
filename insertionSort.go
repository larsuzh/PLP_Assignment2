package main

import (
	"runtime"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

func InsertionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, beep beep) {
	var n = len(data)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1

			beep.wg.Add(1)
			go func() {
				defer beep.wg.Done()
				p := play(beep.c, mapToFeq(int(data[j]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, beep.f)
				beep.m.Lock()
				beep.players = append(beep.players, p)
				beep.m.Unlock()
				Sleep(info.delay)
			}()
			info.comparisons = int(i/2*i) + i - j
			Visualize(win, bars, barWidth, data, j, j-1, info)
			Sleep(info.delay)
		}
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
	VisualizeSorted(win, bars, barWidth, data, beep)
}
