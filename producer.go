package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
File Format

Harvey;9.6
San Juanito;-57.8
Chebba;55.1
Longwan;-70.5
*/

func ReadMeasurements(filename string) chan string {

	lines := make(chan string, 100)

	go func() {
		readFile, err := os.Open(filename)

		if err != nil {
			fmt.Println(err)
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			lines <- fileScanner.Text()
		}

		readFile.Close()
		close(lines)
	}()
	return lines
}
