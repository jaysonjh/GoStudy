package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const row = 10000
const col = 10000

func fillMatrix(matrix *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			matrix[i][j] = s.Intn(10000)
		}
	}
}

func calculate(matrix *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += matrix[i][j]
		}
	}
}

func main() {
	// 创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create cpu profile", err)
	}
	// 获取信息
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start cpu profile", err)
	}
	defer pprof.StopCPUProfile()

	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile", err)
	}
	runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("could not write memory profile", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write goroutine profile", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}
