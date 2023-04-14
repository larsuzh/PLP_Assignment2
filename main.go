package main

import (
	"bufio"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type bar struct {
	rect  pixel.Rect
	color color.Color
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

func Visualize(win *pixelgl.Window, bars []bar, barWidth float64, data []float64) {
	win.Update()
	win.Clear(colornames.Lightslategray)
	for i := 0; i < len(bars); i++ {
		bars[i].rect = pixel.R(barWidth*float64(i), data[i], barWidth*float64(i)+barWidth, 0)
		bars[i].color = colornames.Lightblue
	}
	for b := 0; b < len(bars); b++ {
		imd := imdraw.New(nil)
		imd.Color = bars[b].color
		imd.Push(bars[b].rect.Min, bars[b].rect.Max)
		imd.Rectangle(0)
		imd.Draw(win)
	}
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	algo := readAlgo(reader)
	dataSize := readSize(reader)
	delay := readDelay(reader)

	barWidth = float64(WIDTH) / float64(dataSize)
	bars := make([]bar, dataSize)
	data := generateRandomData(dataSize)

	win := createWindow()

	for !win.Closed() {
		switch algo {
		case "1":
			BubbleSort(win, bars, barWidth, data, delay)
			return
		case "2":
			InsertionSort(win, bars, barWidth, data, delay)
			return
		case "3":
			SelectionSort(win, bars, barWidth, data, delay)
			return
		case "4":
			minHeap := NewMinHeap(data)
			minHeap.Sort(win, bars, barWidth, len(data), delay)
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
