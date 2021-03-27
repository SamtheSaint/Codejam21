package main

import (
	"codejam21/moonsandumbrellas"
	"codejam21/reversesort"
	"fmt"
	"log"
	"os"
)

func main() {
	problem := os.Args[1]

	inputFile, err := os.Open(fmt.Sprintf("./%s/example.in", problem))
	if err != nil {
		log.Fatalln(err)
	}
	
	outputFile, err := os.Create(fmt.Sprintf("./%s/example.out", problem))
	if err != nil {
		log.Fatalln(err)
	}

	switch problem {
	case "reversesort":
		reverseSortDriver(inputFile, outputFile)
	case "moonsandumbrellas":
		moonsAndUmbrellasDriver(inputFile, outputFile)
	default:
		log.Fatalln("Unknown problem")
	}
}

func reverseSortDriver(in, out *os.File) {
	var T, N, _L int
	var L []int

	fmt.Fscanf(in, "%d", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(in, "%d", &N)
		L = []int{}
		for j := 0; j < N; j++ {
			fmt.Fscanf(in,"%d", &_L)
			L = append(L, _L)
		}
		fmt.Fprintf(out, "Case #%d: %d\n", i+1, reversesort.ReverseSort(&L))
	}
}

func moonsAndUmbrellasDriver(in, out *os.File) {
	var T, X, Y int
	var S string
	fmt.Fscanf(in, "%d", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(in, "%d %d %s", &X, &Y, &S)
		moonsandumbrellas.ClearCache()
		fmt.Fprintf(out, "Case #%d: %d\n", i+1,moonsandumbrellas.MinimumCost(X, Y, S))
	}
}

