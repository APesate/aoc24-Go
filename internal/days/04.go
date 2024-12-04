package days

import (
	"aoc24/internal/utils"
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type DayFour struct{}

var xmas []byte = []byte{88, 77, 65, 83}

type direction int

const (
	up direction = iota
	down
	left
	right
	upleft
	upright
	downleft
	downright
	start
)

type point struct {
	x     int
	y     int
	ltrId int
	dir   direction
}

func (d *DayFour) processInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	input := make([][]byte, 0)

	for scanner.Scan() {
		// Crazy gotcha from scanner.Bytes()
		// scanner.Bytes() returns a slice that points to an internal buffer that gets reused.
		// This means multiple rows in your input might end up pointing to the same data.
		input = append(input, append([]byte(nil), scanner.Bytes()...))
	}

	return input
}

func (d *DayFour) wordSearch(graph [][]byte, word []byte) int {
	valid := 0

	for i, l := range graph {
		for j, c := range l {
			if c != word[0] {
				continue
			}

			valid += d.bfs(graph, point{i, j, 0, start}, word)
		}
	}

	return valid
}

func (d *DayFour) bfs(graph [][]byte, strt point, word []byte) int {
	queue := make([]point, 0)

	queue = append(queue, strt)
	valid := 0

	for len(queue) > 0 {
		var curr point
		curr, queue = queue[0], queue[1:]

		if graph[curr.x][curr.y] != word[curr.ltrId] {
			continue
		}

		if curr.ltrId == len(word)-1 {
			valid += 1
			continue
		}

		for _, dir := range []point{{0, 1, curr.ltrId + 1, right}, {0, -1, curr.ltrId + 1, left}, {1, 0, curr.ltrId + 1, down}, {-1, 0, curr.ltrId + 1, up}, {-1, -1, curr.ltrId + 1, upleft}, {-1, 1, curr.ltrId + 1, upright}, {1, -1, curr.ltrId + 1, downleft}, {1, 1, curr.ltrId + 1, downright}} {
			neighbour := point{curr.x + dir.x, curr.y + dir.y, dir.ltrId, dir.dir}
			if neighbour.x < 0 ||
				neighbour.x >= len(graph) ||
				neighbour.y < 0 ||
				neighbour.y >= len(graph[0]) ||
				graph[neighbour.x][neighbour.y] != word[neighbour.ltrId] ||
				(curr.dir != start && curr.dir != neighbour.dir) {
				continue
			}

			queue = append(queue, neighbour)
		}
	}

	return valid
}

func (d *DayFour) Run() {
	file := utils.ReadInput(4, 1)
	input := d.processInput(file)
	// fmt.Println(d.wordSearch(input, xmas)) // Part 1
	fmt.Println(d.findXmas(input)) // part 2
}

type vct struct {
	i int
	j int
}

var deltas = []vct{
	{-1, -1}, {0, 0}, {1, 1}, {1, -1}, {0, 0}, {-1, 1},
}
var mas []byte = []byte{77, 65, 83}
var sam []byte = []byte{83, 65, 77}

func (d *DayFour) findXmas(grid [][]byte) int {
	valid := 0

	bounds := func(row, col int) bool {
		return row >= 0 &&
			row < len(grid) &&
			col >= 0 &&
			col < len(grid[0])
	}

	for i, l := range grid {
		for j, c := range l {
			if c != mas[1] {
				continue
			}

			word1 := make([]byte, 3)
			word2 := make([]byte, 3)

			isValid := true

			for s, d := range deltas[:3] {
				row, col := i+d.i, j+d.j
				if !bounds(row, col) {
					isValid = false
					break
				}
				word1[s] = grid[row][col]
			}

			if !isValid || (!bytes.Equal(word1, mas) && !bytes.Equal(word1, sam)) {
				continue
			}

			for s, d := range deltas[3:] {
				row, col := i+d.i, j+d.j
				if !bounds(row, col) {
					isValid = false
					break
				}
				word2[s] = grid[row][col]
			}

			if isValid && (bytes.Equal(word2, mas) || bytes.Equal(word2, sam)) {
				valid += 1
			}
		}
	}

	return valid
}
