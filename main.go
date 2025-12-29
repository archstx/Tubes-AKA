package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotutil"
)

// ALGORITMA ITERATIF
func findMinMaxIterative(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// ALGORITMA REKURSIF
func findMinMaxRecursive(arr []int, n int) (int, int) {
	if n == 1 {
		return arr[0], arr[0]
	}

	minPrev, maxPrev := findMinMaxRecursive(arr, n-1)

	if arr[n-1] < minPrev {
		minPrev = arr[n-1]
	}
	if arr[n-1] > maxPrev {
		maxPrev = arr[n-1]
	}

	return minPrev, maxPrev
}

// MEMBUAT GRAFIK BATANG
func createBarChart(timeIter, timeRecur float64) {
	p := plot.New()
	p.Title.Text = "Perbandingan Waktu Eksekusi"
	p.Y.Label.Text = "Waktu (detik)"

	values := plotter.Values{timeIter, timeRecur}

	bar, err := plotter.NewBarChart(values, vg.Points(40))
	if err != nil {
		panic(err)
	}

	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(1)

	p.Add(bar)
	p.NominalX("Iteratif", "Rekursif")

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "perbandingan_waktu.png"); err != nil {
		panic(err)
	}

	fmt.Println("Grafik disimpan sebagai perbandingan_waktu.png")
}

// FUNGSI COMPARE
func compare(arr []int) {
	loop := 10000

	var timeIter, timeRecur time.Duration

	for i := 0; i < loop; i++ {
		start := time.Now()
		findMinMaxIterative(arr)
		timeIter += time.Since(start)
	}

	for i := 0; i < loop; i++ {
		start := time.Now()
		findMinMaxRecursive(arr, len(arr))
		timeRecur += time.Since(start)
	}

	avgIter := timeIter.Seconds() / float64(loop)
	avgRecur := timeRecur.Seconds() / float64(loop)

	minI, maxI := findMinMaxIterative(arr)
	minR, maxR := findMinMaxRecursive(arr, len(arr))

	fmt.Println("=== HASIL PERBANDINGAN ===")
	fmt.Printf("Iteratif  -> Min: %d, Max: %d\n", minI, maxI)
	fmt.Printf("Waktu rata-rata: %.8f detik\n\n", avgIter)

	fmt.Printf("Rekursif -> Min: %d, Max: %d\n", minR, maxR)
	fmt.Printf("Waktu rata-rata: %.8f detik\n", avgRecur)

	createBarChart(avgIter, avgRecur)
}

// MAIN
func main() {
	rand.Seed(time.Now().UnixNano())

	n := 1000
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10000)
	}

	compare(arr)
}
