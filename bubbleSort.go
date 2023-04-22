package main

import (
	"runtime"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

func BubbleSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, beep beep) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
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
			info.comparisons = i*len(data) + j
			Visualize(win, bars, barWidth, data, j, j+1, info)
			Sleep(info.delay)
		}
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
	VisualizeSorted(win, bars, barWidth, data, beep)
}
