package main

import "testing"

func BenchmarkBRC(b *testing.B) {

	for i := 0; i < b.N; i++ {
		lines := ReadMeasurements("10k.txt")

		SplitLine(lines)
	}
}
