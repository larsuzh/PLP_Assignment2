package main

import (
	"bufio"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/hajimehoshi/oto"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type bar struct {
	rect  pixel.Rect
	color color.Color
}

type info struct {
	algo        string
	delay       int
	comparisons int
}

type beep struct {
	players []oto.Player
	f       int
	c       *oto.Context
	wg      sync.WaitGroup
	m       sync.Mutex
}

var (
	barWidth float64
)

const (
	WIDTH  = 1024
	HEIGHT = 700
)

func readAlgo(reader *bufio.Reader) string {
	fmt.Println("Select a sorting algorithm:")
	fmt.Println("1. Bubble sort")
	fmt.Println("2. Insertion sort")
	fmt.Println("3. Selection sort")
	fmt.Println("4. Heap sort")
	fmt.Println("5. Bogo sort")
	fmt.Print("Enter option number: ")
	algo, _ := reader.ReadString('\n')
	return strings.TrimSpace(algo)
}

func readSize(reader *bufio.Reader) int {
	fmt.Print("> n: ")
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	dataSize, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Invalid Input: Provided Key is not a valid number!", err)
	}
	return dataSize
}

func readDelay(reader *bufio.Reader) int {
	fmt.Print("> delay [ms]: ")
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	delay, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Invalid Input: Provided Key is not a valid number!", err)
	}
	return delay
}

func generateRandomData(size int) []float64 {
	data := make([]float64, size)
	for i := 0; i < len(data); i++ {
		data[i] = float64(HEIGHT / size * (size - i - 1))
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}

func Sleep(duration int) {
	delay := time.Duration(duration) * time.Millisecond
	time.Sleep(delay)
}

func createWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Visualizing Algos",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func mapToFeq(num, n int) float64 {
	return float64(num*1140/n + 60)
}

func Visualize(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, j int, k int, info info) {
	win.Update()
	win.Clear(colornames.Lightslategray)
	for i := 0; i < len(bars); i++ {
		bars[i].rect = pixel.R(barWidth*float64(i), data[i], barWidth*float64(i)+barWidth, 0)
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
	infoTxt := info.algo + ", delay: " + fmt.Sprintf("%d", info.delay) + ", comparisons: " + fmt.Sprintf("%d", info.comparisons)
	txt.WriteString(infoTxt)
	txt.Draw(win, pixel.IM)
}

func VisualizeSorted(win *pixelgl.Window, bars []bar, barWidth float64, data []float64, beep beep) {

	for j := 0; j < len(bars); j++ {
		beep.wg.Add(1)
		go func() {
			defer beep.wg.Done()
			p := play(beep.c, mapToFeq(j, len(bars)), time.Duration(600/len(data))*time.Millisecond, *channelCount, beep.f)
			beep.m.Lock()
			beep.players = append(beep.players, p)
			beep.m.Unlock()
			Sleep(600 / len(data))
		}()
		win.Update()
		win.Clear(colornames.Lightslategray)
		for i := 0; i < len(bars); i++ {
			bars[i].rect = pixel.R(barWidth*float64(i), data[i], barWidth*float64(i)+barWidth, 0)
			if i > j {
				bars[i].color = colornames.Lightblue
			} else {
				bars[i].color = colornames.Green
			}

		}
		for b := 0; b < len(bars); b++ {
			imd := imdraw.New(nil)
			imd.Color = bars[b].color
			imd.Push(bars[b].rect.Min, bars[b].rect.Max)
			imd.Rectangle(0)
			imd.Draw(win)
		}
	}
	beep.wg.Wait()
	runtime.KeepAlive(beep.players)
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	var info info
	var beep beep
	info.algo = readAlgo(reader)
	dataSize := readSize(reader)
	info.delay = readDelay(reader)

	barWidth = float64(WIDTH) / float64(dataSize)
	bars := make([]bar, dataSize)
	data := generateRandomData(dataSize)

	win := createWindow()

	switch *format {
	case "f32le":
		beep.f = oto.FormatFloat32LE
	case "u8":
		beep.f = oto.FormatUnsignedInt8
	case "s16le":
		beep.f = oto.FormatSignedInt16LE
	default:
		return
	}
	c, ready, err := oto.NewContext(*sampleRate, *channelCount, beep.f)
	if err != nil {
		return
	}
	beep.c = c
	<-ready

	for !win.Closed() {
		switch info.algo {
		case "1":
			info.algo = "Bubble sort"
			BubbleSort(win, bars, barWidth, data, info, beep)
			return
		case "2":
			info.algo = "Insertion sort"
			InsertionSort(win, bars, barWidth, data, info, beep)
			return
		case "3":
			info.algo = "Selection sort"
			SelectionSort(win, bars, barWidth, data, info, beep)
			return
		case "4":
			info.algo = "Heap sort"
			minHeap := NewMinHeap(data)
			minHeap.Sort(win, bars, barWidth, len(data), info, beep)
			return
		case "5":
			info.algo = "Bogo sort"
			info.comparisons = 0
			BogoSort(win, bars, barWidth, data, info, beep)
			return
		default:
			fmt.Println("Invalid option, please try again")
		}
		break
	}
}

func main() {
	pixelgl.Run(run)
}
