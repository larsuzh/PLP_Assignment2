package main

import (
	"runtime"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

func SelectionSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, beep beep) {
	var n = len(data)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
			beep.wg.Add(1)
			go func() {
				defer beep.wg.Done()
				p := play(beep.c, mapToFeq(int(data[j]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, beep.f)
				beep.m.Lock()
				beep.players = append(beep.players, p)
				beep.m.Unlock()
				Sleep(info.delay)
			}()
			info.comparisons = i*len(data) + j - i
			Visualize(win, bars, barWidth, data, minIdx, j, info)
			Sleep(info.delay)
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
	VisualizeSorted(win, bars, barWidth, data, beep)
}
