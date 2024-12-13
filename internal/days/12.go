package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
)

type DayTwelve struct{}

type d12_Region struct {
	name      rune
	area      int
	perimeter int
	side      int
}

func (r d12_Region) price(part int) int {
	if part == 1 {
		return r.area * r.perimeter
	} else {
		return r.area * r.side
	}
}

func (d *DayTwelve) processInput(file *os.File) [][]rune {
	scanner := bufio.NewScanner(file)
	grid := [][]rune{}

	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, []rune{})
		for _, r := range l {
			grid[len(grid)-1] = append(grid[len(grid)-1], r)
		}
	}

	return grid
}

func (d *DayTwelve) restOfTheSide(str [2]int, grid [][]rune, dir [2]int, look [2]int) [][2]int {
	i, j := str[0], str[1]
	match := [][2]int{}
	r := grid[i][j]
	curr := str

	for curr[0] >= 0 && curr[0] < len(grid) && curr[1] >= 0 && curr[1] < len(grid[0]) {
		if grid[curr[0]][curr[1]] != r ||
			curr[0]+look[0] >= 0 && curr[0]+look[0] < len(grid) && curr[1]+look[1] >= 0 && curr[1]+look[1] < len(grid[0]) && grid[curr[0]+look[0]][curr[1]+look[1]] == r {
			break
		}

		match = append(match, curr)
		curr = [2]int{curr[0] + dir[0], curr[1] + dir[1]}
	}

	return match
}

func (d *DayTwelve) calculateRegiosnsPrice(grid [][]rune) (int, int) {
	regions := []d12_Region{}
	visited := [][2]int{}

	for row := 0; row < len(grid); row += 1 {
		for col := 0; col < len(grid); col += 1 {
			if slices.Contains(visited, [2]int{row, col}) {
				continue
			}

			queue := [][2]int{}
			queue = append(queue, [2]int{row, col})
			r := grid[row][col]
			region := d12_Region{name: r}
			sides := map[rune]map[[2]int]bool{}

			for _, l := range []rune{'u', 'l', 'd', 'r'} {
				sides[l] = map[[2]int]bool{}
			}

			for len(queue) > 0 {
				var curr [2]int
				curr, queue = queue[0], queue[1:]
				visited = append(visited, curr)
				i, j := curr[0], curr[1]

				region.area += 1

				if i-1 < 0 || grid[i-1][j] != r {
					region.perimeter += 1

					if !sides['u'][curr] {
						region.side += 1
						for _, v := range d.restOfTheSide(curr, grid, [2]int{0, 1}, [2]int{-1, 0}) {
							sides['u'][v] = true
						}
						for _, v := range d.restOfTheSide(curr, grid, [2]int{0, -1}, [2]int{-1, 0}) {
							sides['u'][v] = true
						}
					}
				}
				if j-1 < 0 || grid[i][j-1] != r {
					region.perimeter += 1

					if !sides['l'][curr] {
						region.side += 1
						for _, v := range d.restOfTheSide(curr, grid, [2]int{1, 0}, [2]int{0, -1}) {
							sides['l'][v] = true
						}
						for _, v := range d.restOfTheSide(curr, grid, [2]int{-1, 0}, [2]int{0, -1}) {
							sides['l'][v] = true
						}
					}
				}
				if i+1 >= len(grid) || grid[i+1][j] != r {
					region.perimeter += 1

					if !sides['d'][curr] {
						region.side += 1
						for _, v := range d.restOfTheSide(curr, grid, [2]int{0, 1}, [2]int{1, 0}) {
							sides['d'][v] = true
						}
						for _, v := range d.restOfTheSide(curr, grid, [2]int{0, -1}, [2]int{1, 0}) {
							sides['d'][v] = true
						}
					}
				}
				if j+1 >= len(grid[0]) || grid[i][j+1] != r {
					region.perimeter += 1

					if !sides['r'][curr] {
						region.side += 1
						for _, v := range d.restOfTheSide(curr, grid, [2]int{1, 0}, [2]int{0, 1}) {
							sides['r'][v] = true
						}
						for _, v := range d.restOfTheSide(curr, grid, [2]int{-1, 0}, [2]int{0, 1}) {
							sides['r'][v] = true
						}
					}
				}

				for _, delta := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
					if slices.Contains(visited, [2]int{i + delta[0], j + delta[1]}) ||
						slices.Contains(queue, [2]int{i + delta[0], j + delta[1]}) ||
						i+delta[0] < 0 || i+delta[0] >= len(grid) ||
						j+delta[1] < 0 || j+delta[1] >= len(grid[0]) ||
						grid[i+delta[0]][j+delta[1]] != r {
						continue
					}

					queue = append(queue, [2]int{i + delta[0], j + delta[1]})
				}
			}

			regions = append(regions, region)
		}
	}

	t1, t2 := 0, 0
	for _, r := range regions {
		t1 += r.price(1)
		t2 += r.price(2)
	}
	return t1, t2
}

func (d *DayTwelve) Run() {
	file := utils.ReadInput(12, 1)
	defer file.Close()
	grid := d.processInput(file)
	p1, p2 := d.calculateRegiosnsPrice(grid)
	fmt.Println("Part One:", p1)
	fmt.Println("Part Two:", p2)
}
