package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main() {

	f, perr := os.Create("cpu.pprof")
	if perr != nil {
		panic(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	lines := ReadMeasurements("1mil.txt")

	SplitLine(lines)

	//ProcessAndStore(temps, &SM)
}
