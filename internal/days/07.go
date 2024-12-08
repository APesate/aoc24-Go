package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type DaySeven struct {
	part int
}

type d7_Line struct {
	res     int
	strNums []string
}

type d7_Equation struct {
	res        int
	numbers    []int
	operations []d7_Operation
}

type d7_Operation struct {
	symbol string
	apply  func(a, b int) int
}

var d7_Operations = []d7_Operation{
	{"+", func(a, b int) int { return a + b }},
	{"*", func(a, b int) int { return a * b }},
	{
		"||", func(a, b int) int {
			aS := strconv.Itoa(a)
			bS := strconv.Itoa(b)
			n, _ := strconv.Atoi(aS + bS)
			return n
		},
	},
}

var part int

func (d *DaySeven) processInput(file *os.File) int {
	scanner := bufio.NewScanner(file)
	result := 0

	// Parallel Exec
	validResult := make(chan int)
	workers := make(chan struct{}, 10)
	var wg sync.WaitGroup

	go func() {
		for r := range validResult {
			result += r
		}
	}()

	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, ": ")
		res, _ := strconv.Atoi(parts[0])
		strNums := strings.Split(parts[1], " ")

		line := d7_Line{
			res,
			strNums,
		}

		wg.Add(1)
		workers <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-workers }()
			if d.evaluateLine(line) {
				validResult <- line.res
			}
		}()
	}

	wg.Wait()
	close(validResult)
	close(workers)

	return result
}

func (d *DaySeven) evaluateLine(l d7_Line) bool {
	numValidOps := 2
	if d.part == 2 {
		numValidOps = 3
	}

	numOps := len(l.strNums) - 1
	maxComb := int(math.Pow(float64(numValidOps), float64(len(l.strNums)-1)))

	numbers := make([]int, len(l.strNums))
	for i, s := range l.strNums {
		n, _ := strconv.Atoi(s)
		numbers[i] = n
	}

	for i := range maxComb {
		temp := i
		operators := make([]d7_Operation, numOps)

		for j := range numOps {
			switch temp % numValidOps {
			case 0:
				operators[j] = d7_Operations[0]
			case 1:
				operators[j] = d7_Operations[1]
			case 2:
				operators[j] = d7_Operations[2]
			}
			temp /= numValidOps
		}

		// Evaluate Expression
		eq := d7_Equation{
			res:        l.res,
			numbers:    numbers,
			operations: operators,
		}

		if eq.isValid() {
			return true
		}
	}

	return false
}

func (eq *d7_Equation) isValid() bool {
	total := eq.numbers[0]

	for i, op := range eq.operations {
		total = op.apply(total, eq.numbers[i+1])
	}

	return total == eq.res
}

func (d *DaySeven) Run() {
	file := utils.ReadInput(7, 1)
	defer file.Close()

	// d.part = 1
	// r1 := d.processInput(file)
	// fmt.Println("Part 1: ", r1)

	d.part = 2
	r2 := d.processInput(file)
	fmt.Println("Part 2: ", r2)
}
