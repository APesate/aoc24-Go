package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
)

type DaySix struct{}

type d6_direction byte

const (
	d6_up    = 0x1
	d6_right = 0x2
	d6_down  = 0x4
	d6_left  = 0x8
)

func addDirection(cell byte, direction byte) byte {
	return cell | direction
}

func hasDirection(cell byte, direction byte) bool {
	return (cell & direction) != 0
}

func signForCell(cell byte) string {
	if hasDirection(cell, d6_up|d6_down) && hasDirection(cell, d6_right|d6_left) {
		return "+"
	}
	if hasDirection(cell, d6_up|d6_down) {
		return "|"
	}
	if hasDirection(cell, d6_right|d6_left) {
		return "-"
	}

	return "."
}

const UP_SIGN byte = 94
const STEP_SIGN byte = 43
const WALL_SIGN byte = 35

func (d *DaySix) processInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	result := [][]byte{}

	for scanner.Scan() {
		result = append(result, append([]byte(nil), scanner.Bytes()...))
	}

	return result
}

func (d *DaySix) startingPoint(input [][]byte) (int, int) {
	stp := [2]int{-1, -1}
	for i, l := range input {
		r := slices.Index(l, UP_SIGN)
		if r >= 0 {
			stp = [2]int{i, r}
			break
		}
	}

	return stp[0], stp[1]
}

func (d *DaySix) walkSim(input [][]byte) ([][]byte, int, [][]byte, bool) {
	w, h := len(input), len(input[0])
	p1_result := utils.CopyGrid(input)
	p2_result := make([][]byte, w)
	for i := range p2_result {
		p2_result[i] = make([]byte, h)
	}
	row, col := d.startingPoint(input)
	dir := d6_up
	stop := false
	found_loop := false
	count := 0

	for !stop && !found_loop {
		dR, dC := 0, 0
		if p1_result[row][col] != STEP_SIGN {
			count += 1
		}

		p1_result[row][col] = STEP_SIGN

		if hasDirection(p2_result[row][col], byte(dir)) {
			found_loop = true
			break
		}

		p2_result[row][col] = addDirection(p2_result[row][col], byte(dir))

		switch dir {
		case d6_up:
			if !utils.IsRowInBounds(input, row-1) {
				stop = true
				break
			}

			if input[row-1][col] != WALL_SIGN {
				dR = -1
			} else {
				dir = d6_right
			}
		case d6_right:
			if !utils.IsColInBounds(input, col+1) {
				stop = true
				break
			}

			if input[row][col+1] != WALL_SIGN {
				dC = 1
			} else {
				dir = d6_down
			}
		case d6_down:
			if !utils.IsRowInBounds(input, row+1) {
				stop = true
				break
			}

			if input[row+1][col] != WALL_SIGN {
				dR = 1
			} else {
				dir = d6_left
			}
		case d6_left:
			if !utils.IsColInBounds(input, col-1) {
				stop = true
				break
			}

			if input[row][col-1] != WALL_SIGN {
				dC = -1
			} else {
				dir = d6_up
			}
		}

		row, col = row+dR, col+dC
	}

	return p1_result, count, p2_result, found_loop
}

func (d *DaySix) wallPlacement(input [][]byte, row, col, ogRow, ogCol int) int {
	if row == -1 || col == -1 {
		return 0
	}

	nR, nC := d.nextPos(input, row, col)

	if input[row][col] == WALL_SIGN || (row == ogRow && col == ogCol) {
		return 0 + d.wallPlacement(input, nR, nC, ogRow, ogCol)
	}

	variation := utils.CopyGrid(input)
	variation[row][col] = WALL_SIGN

	_, _, _, loop := d.walkSim(variation)
	delta := 0
	if loop {
		delta = 1
	}

	return delta + d.wallPlacement(input, nR, nC, ogRow, ogCol)
}

func (d *DaySix) nextPos(input [][]byte, row, col int) (int, int) {
	if col+1 < len(input[0]) {
		return row, col + 1
	} else {
		if row+1 >= len(input) {
			return -1, -1
		}
		return row + 1, 0
	}
}

func (d *DaySix) visualizeLoop(grid [][]byte) {
	for i := range grid {
		for _, c := range grid[i] {
			fmt.Printf(signForCell(c))
		}
		fmt.Println("")
	}
	fmt.Print("\n\n")
}

func (d *DaySix) Run() {
	file := utils.ReadInput(6, 1)
	defer file.Close()
	input := d.processInput(file)
	_, p1_loc, _, _ := d.walkSim(input)
	ogR, ogC := d.startingPoint(input)
	loops := d.wallPlacement(input, 0, 0, ogR, ogC)

	fmt.Println("Loc: ", p1_loc)
	fmt.Println("Loops: ", loops)
}
