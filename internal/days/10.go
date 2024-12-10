package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type DayTen struct{}

func (d *DayTen) processInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	lineNum := 0

	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, []int{})

		for _, r := range l {
			n, _ := strconv.Atoi(string(r))
			grid[lineNum] = append(grid[lineNum], n)
		}

		lineNum += 1
	}

	return grid
}

func (d *DayTen) findTrailheads(grid [][]int) [][2]int {
	trailheads := [][2]int{}

	for i, l := range grid {
		for j, n := range l {
			if n == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	return trailheads
}

func (d *DayTen) findPaths(grid [][]int, str [2]int) (int, int) {
	queue := [][2]int{}
	queue = append(queue, str)
	destinations := map[[2]int]int{}

	for len(queue) > 0 {
		var curr [2]int
		curr, queue = queue[0], queue[1:]

		if grid[curr[0]][curr[1]] == 9 {
			destinations[curr] += 1
			continue
		}

		// Add Neighbours
		for _, nh := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if curr[0]+nh[0] < 0 ||
				curr[0]+nh[0] >= len(grid) ||
				curr[1]+nh[1] < 0 ||
				curr[1]+nh[1] >= len(grid[0]) ||
				grid[curr[0]+nh[0]][curr[1]+nh[1]]-grid[curr[0]][curr[1]] != 1 {
				continue
			}

			queue = append(queue, [2]int{curr[0] + nh[0], curr[1] + nh[1]})
		}
	}

	total := 0
	for _, v := range destinations {
		total += v
	}
	return len(destinations), total
}

func (d *DayTen) partOne(grid [][]int) (int, int) {
	trailheads := d.findTrailheads(grid)
	t1 := 0
	t2 := 0

	var wg sync.WaitGroup
	r1 := make(chan int)
	r2 := make(chan int)

	var processorWg sync.WaitGroup
	processorWg.Add(1)
	go func() {
		defer processorWg.Done()
		for r := range r1 {
			t1 += r
		}
	}()

	processorWg.Add(1)
	go func() {
		defer processorWg.Done()
		for r := range r2 {
			t2 += r
		}
	}()

	for _, th := range trailheads {
		wg.Add(1)
		go func(th [2]int) {
			defer wg.Done()
			p1, p2 := d.findPaths(grid, th)
			r1 <- p1
			r2 <- p2
		}(th)
	}

	wg.Wait()
	close(r1)
	close(r2)
	processorWg.Wait()

	return t1, t2
}

func (d *DayTen) Run() {
	file := utils.ReadInput(10, 1)
	defer file.Close()
	grid := d.processInput(file)
	r1, r2 := d.partOne(grid)
	fmt.Println("Part One:", r1)
	fmt.Println("Part Two:", r2)
}
