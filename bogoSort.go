package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/faiface/pixel/pixelgl"
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

func BogoSort(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, info info, beep beep) {
	for {
		beep.wg.Add(1)
		go func() {
			defer beep.wg.Done()
			p := play(beep.c, mapToFeq(int(data[rand.Intn(len(bars))]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, beep.f)
			beep.m.Lock()
			beep.players = append(beep.players, p)
			beep.m.Unlock()
			Sleep(info.delay)
		}()
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
		Visualize(win, bars, barWidth, data, -1, -1, info)
		Sleep(info.delay)
		if checkSortedArray(data) {
			break
		}
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
	VisualizeSorted(win, bars, barWidth, data, beep)
}
