package main

import "fmt"

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	fmt.Println("Soal 1")
	paskal := generate(15)
	soal1(paskal)

	fmt.Println()
	fmt.Println("Soal 2")
	soal2(paskal)

	fmt.Println()
	fmt.Println("Soal 3")
	soal3(paskal)
}

func printValue(idx, val, rowLen int) {
	if idx == 0 || idx == rowLen-1 {
		fmt.Printf("%d ", val)
		return
	}

	if val%2 == 0 {
		fmt.Printf(Green+"%d "+Reset, val)
	} else {
		fmt.Printf(Red+"%d "+Reset, val)
	}
}

func soal1(paskal [][]int) {
	for _, row := range paskal {
		for a, b := range row {
			printValue(a, b, len(row))
		}

		fmt.Println()
	}
}

func soal2(paskal [][]int) {
	for _, row := range paskal {
		sum := 0

		for a, b := range row {
			printValue(a, b, len(row))

			if a > 0 && a < len(row)-1 {
				sum += b
			}
		}

		if len(row) > 2 {
			fmt.Printf("== %d", sum)
		} else {
			fmt.Printf("== %d", sum)
		}

		fmt.Println()
	}
}

func soal3(paskal [][]int) {
	for i := len(paskal) - 1; i >= 0; i-- {
		sum := 0
		for a, b := range paskal[i] {
			printValue(a, b, len(paskal[i]))

			if a > 0 && a < len(paskal[i])-1 {
				sum += b
			}
		}

		if len(paskal[i]) > 2 {
			fmt.Printf("== %d", sum)
		} else {
			fmt.Printf("== %d", sum)
		}
		fmt.Println()
	}
}
