package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Get CPU profile.
	f, err := os.Create("cpu_profile")
	if err != nil {
		fmt.Println("cannot create cpu_profile")
		return
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("get cpu profile fail")
		return
	}
	defer pprof.StopCPUProfile()

	// code
	data := fillData()
	res := sum(data)
	fmt.Printf("res is %d\n", res)


	// Get memory profile.
	mf, err := os.Create("memory_profile")
	if err != nil {
		fmt.Println("cannot create memory_profile")
		return
	}
	if err := pprof.WriteHeapProfile(mf); err != nil {
		fmt.Println("cannot write to memory_profile")
		return
	}
	_ = mf.Close()


	// Get goroutine profile.
	gf, err := os.Create("goroutine_profile")
	if err != nil {
		fmt.Println("cannot create goroutine_profile")
		return
	}
	if gProfile := pprof.Lookup("goroutine"); gProfile == nil {
		fmt.Println("cannot get goroutine profile")
		return
	} else {
		_ = gProfile.WriteTo(gf, 0)
	}
	_ = gf.Close()

}

func fillData() [1000][1000]int {
	res := [1000][1000]int{}

	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range res {
		for j := range res[i] {
			res[i][j] = s.Intn(10000)
		}
	}

	return res
}

func sum(data [1000][1000]int) int {
	tmp := 0
	for _, row := range data {
		for _, v := range row {
			tmp += v
		}
	}

	return tmp
}