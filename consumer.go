package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var SM sync.Map

type Temperature struct {
	City        string
	Temperature float64
}

type TempStats struct {
	Min float64
	Avg float64
	Max float64
}

func SplitLine(lines <-chan string) error {

	wg := sync.WaitGroup{}
	for l := range lines {
		acc := make([]string, 0, 1000)
		for i := 0; i < 1000; i++ {
			acc = append(acc, l)
			l = <-lines
		}

		wg.Add(1)
		go func(acc []string) {
			defer wg.Done()

			tps := make([]Temperature, 1000)
			for _, l := range acc {
				if len(l) == 0 {
					continue
				}
				vals := strings.Split(l, ";")

				if len(vals) != 2 {
					panic(fmt.Sprintf("%s, %+v", l, vals))
				}

				temp, err := strconv.ParseFloat(vals[1], 64)
				if err != nil {
					panic(fmt.Sprintf("%+v, %s", vals, err))
				}

				tps = append(tps, Temperature{City: vals[0], Temperature: temp})
			}

			go ProcessAndStore(tps, &SM)
		}(acc)
	}

	wg.Wait()

	fmt.Printf("Ranging over sync.Map\n")
	SM.Range(func(key, value interface{}) bool {
		fmt.Printf("%s %+v \n", key, value)
		return true
	})

	return nil
}

func ProcessAndStore(tps []Temperature, sm *sync.Map) {

	for _, t := range tps {
		actual, loaded := sm.LoadOrStore(t.City, TempStats{Min: t.Temperature, Avg: t.Temperature, Max: t.Temperature})
		if loaded {
			val, ok := actual.(TempStats)
			if ok {
				if val.Min > t.Temperature {
					val.Min = t.Temperature
				}

				if val.Max < t.Temperature {
					val.Max = t.Temperature
				}

				val.Avg = (val.Avg + t.Temperature) / 2
				sm.Store(t.City, val)
			}
		}
	}

}
