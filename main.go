package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	ensureDataDir()

	var k int
	fmt.Print("Masukkan jumlah perbandingan: ")
	fmt.Scan(&k)

	ns := make([]int, k)
	fmt.Println("Masukkan nilai n (contoh: 5 10 100 1000):")
	for i := 0; i < k; i++ {
		fmt.Scan(&ns[i])
	}

	for _, n := range ns {
		fmt.Println("Menjalankan eksperimen untuk n =", n)
		experiment(n)
	}
}

func minMaxIterative(arr []int) (int, int) {
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

func minMaxRecursive(arr []int, n int) (int, int) {
	if n == 1 {
		return arr[0], arr[0]
	}
	minPrev, maxPrev := minMaxRecursive(arr, n-1)

	if arr[n-1] < minPrev {
		minPrev = arr[n-1]
	}
	if arr[n-1] > maxPrev {
		maxPrev = arr[n-1]
	}
	return minPrev, maxPrev
}

func ensureDataDir() {
	if _, err := os.Stat("../data"); os.IsNotExist(err) {
		err := os.Mkdir("../data", 0755)
		if err != nil {
			panic(err)
		}
	}
}

func experiment(n int) {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000)
	}

	loop := 10000

	var tBothIter, tBothRec time.Duration

	for i := 0; i < loop; i++ {
		start := time.Now()
		minMaxIterative(arr)
		tBothIter += time.Since(start)

		start = time.Now()
		minMaxRecursive(arr, len(arr))
		tBothRec += time.Since(start)
	}

	minI, maxI := minMaxIterative(arr)
	minR, maxR := minMaxRecursive(arr, len(arr))

	file, err := os.Create(fmt.Sprintf("../data/result_n%d.csv", n))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"N",
		"Min_Iter", "Min_Recur",
		"Max_Iter", "Max_Recur",
		"Time_Iter", "Time_Recur",
	})

	writer.Write([]string{
		strconv.Itoa(n),
		strconv.Itoa(minI),
		strconv.Itoa(minR),
		strconv.Itoa(maxI),
		strconv.Itoa(maxR),
		fmt.Sprintf("%.8f", tBothIter.Seconds()/float64(loop)),
		fmt.Sprintf("%.8f", tBothRec.Seconds()/float64(loop)),
	})
}

