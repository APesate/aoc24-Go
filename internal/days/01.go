package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DayOne struct{}

func (d *DayOne) makeListsFromFile(file *os.File) ([]string, []string) {
	scanner := bufio.NewScanner(file)
	lhs := make([]string, 0, 1000)
	rhs := make([]string, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		lhs = append(lhs, values[0])
		rhs = append(rhs, values[1])
	}

	if err := scanner.Err(); err != nil {
		slog.Error(err.Error())
	}

	slices.Sort(lhs)
	slices.Sort(rhs)

	return lhs, rhs
}

func (d *DayOne) makeMapFromFile(file *os.File) ([]string, map[string]int) {
	scanner := bufio.NewScanner(file)
	lhs := make([]string, 0, 1000)
	entriesMap := make(map[string]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		lhs = append(lhs, values[0])

		entriesMap[values[1]] += 1
	}

	if err := scanner.Err(); err != nil {
		slog.Error(err.Error())
	}

	return lhs, entriesMap
}

func (d *DayOne) calcDistances(lhs, rhs []string) int {
	total := 0
	for i, s := range lhs {
		l, _ := strconv.Atoi(s)
		r, _ := strconv.Atoi(rhs[i])
		total += int(math.Abs(float64(l) - float64(r)))
	}

	return total
}

func (d *DayOne) calcSimilarity(l []string, m map[string]int) int {
	total := 0

	for _, v := range l {
		iv, _ := strconv.Atoi(v)
		total += iv * m[v]
	}

	return total
}

func (d *DayOne) Run() {
	file := utils.ReadInput(1, 1)
	defer file.Close()

	// Part 1
	// lhs, rhs := makeListsFromFile(file)
	// fmt.Printf("Total distance: %d", calcDistances(lhs, rhs))

	// Part 2
	lhs, rhsm := d.makeMapFromFile(file)
	fmt.Printf("Part 2: %d", d.calcSimilarity(lhs, rhsm))
}
