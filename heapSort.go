package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type minheap struct {
	arr []float64
}

func NewMinHeap(arr []float64) *minheap {
	minheap := &minheap{
		arr: arr,
	}
	return minheap
}

func (m *minheap) leftchildIndex(index int) int {
	return 2*index + 1
}

func (m *minheap) rightchildIndex(index int) int {
	return 2*index + 2
}

func (m *minheap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *minheap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (m *minheap) downHeapify(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchildIndex(current)
	rightRightIndex := m.rightchildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest, size)
	}
	return
}

func (m *minheap) buildMinHeap(size int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

func (m *minheap) Sort(win *pixelgl.Window, bars []bar, barWidth float64, size int, info info, beep beep) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		m.swap(0, i)
		m.downHeapify(0, i)
		beep.wg.Add(1)
		go func() {
			defer beep.wg.Done()
			p := play(beep.c, mapToFeq(int(m.arr[i]), len(bars)), time.Duration(info.delay)*time.Millisecond, *channelCount, beep.f)
			beep.m.Lock()
			beep.players = append(beep.players, p)
			beep.m.Unlock()
			Sleep(info.delay)
		}()
		info.comparisons = (size - i) * int(math.Log2(float64(size)))
		m.Visualize(win, bars, barWidth, 0, i, info)
		Sleep(info.delay)
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
	VisualizeSorted(win, bars, barWidth, m.arr, beep)
}

func (m *minheap) Visualize(win *pixelgl.Window, bars []bar, barWidth float64, j int, k int, info info) {
	win.Update()
	win.Clear(colornames.Lightslategray)
	for i := 0; i < len(bars); i++ {
		bars[i].rect = pixel.R(barWidth*float64(i), m.arr[i], barWidth*float64(i)+barWidth, 0)
		bars[i].color = colornames.Lightblue
	}
	if j >= 0 {
		bars[j].color = colornames.Red
	}
	if k >= 0 {
		bars[k].color = colornames.Red
	}
	for b := 0; b < len(bars); b++ {
		imd := imdraw.New(nil)
		imd.Color = bars[b].color
		imd.Push(bars[b].rect.Min, bars[b].rect.Max)
		imd.Rectangle(0)
		imd.Draw(win)
	}
	atlas := text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)
	txt := text.New(pixel.V(30, 680), atlas)
	txt.Color = colornames.Black
	infoTxt := info.algo + ", delay [ms]: " + fmt.Sprintf("%d", info.delay) + ", comparisons: " + fmt.Sprintf("%d", info.comparisons)
	txt.WriteString(infoTxt)
	txt.Draw(win, pixel.IM)
}
