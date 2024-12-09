package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
)

type DayEight struct{}

type d8_Antenna struct {
	i      int
	j      int
	symbol rune
}

type d8_Node struct {
	i int
	j int
}

func (n *d8_Node) isValidInGrid(gridSize [2]int) bool {
	vi := n.i >= 0 && n.i < gridSize[0]
	vj := n.j >= 0 && n.j < gridSize[1]
	return vi && vj
}

func (d *DayEight) processInput(file *os.File) ([2]int, map[rune][]d8_Antenna) {
	scanner := bufio.NewScanner(file)
	antennas := make(map[rune][]d8_Antenna, 0)
	i := 0
	gridSize := [2]int{}

	for scanner.Scan() {
		l := scanner.Text()

		if gridSize[1] == 0 {
			gridSize[1] = len(l)
		}

		for j, r := range l {
			if r == '.' {
				continue
			}

			antennas[r] = append(antennas[r], d8_Antenna{
				i:      i,
				j:      j,
				symbol: r,
			})
		}

		i += 1
	}

	gridSize[0] = i

	return gridSize, antennas
}

func (d *DayEight) findMatches(gridSize [2]int, antennas map[rune][]d8_Antenna) int {
	nodes := map[d8_Node]bool{}

	for _, v := range antennas {
		for i, lhs := range v {
			for _, rhs := range v[i+1:] {
				di := lhs.i - rhs.i
				dj := lhs.j - rhs.j
				nodes[d8_Node{lhs.i + di, lhs.j + dj}] = true
				nodes[d8_Node{rhs.i - di, rhs.j - dj}] = true
			}
		}
	}

	total := 0
	for n := range nodes {
		if n.isValidInGrid(gridSize) {
			total += 1
		}
	}

	return total
}

func (d *DayEight) part2FindMatches(gridSize [2]int, antennas map[rune][]d8_Antenna) int {
	nodes := map[d8_Node]int{}

	for _, v := range antennas {
		for i, lhs := range v {
			if len(v) > 1 {
				nodes[d8_Node{lhs.i, lhs.j}] += 1
			}

			for _, rhs := range v[i+1:] {
				di := lhs.i - rhs.i
				dj := lhs.j - rhs.j

				n1 := d8_Node{lhs.i + di, lhs.j + dj}
				for n1.isValidInGrid(gridSize) {
					nodes[n1] += 1
					n1 = d8_Node{n1.i + di, n1.j + dj}
				}

				n2 := d8_Node{rhs.i - di, rhs.j - dj}
				for n2.isValidInGrid(gridSize) {
					nodes[n2] += 1
					n2 = d8_Node{n2.i - di, n2.j - dj}
				}
			}
		}
	}

	return len(nodes)
}

func (d *DayEight) Run() {
	file := utils.ReadInput(8, 1)
	defer file.Close()
	gridSize, antennas := d.processInput(file)
	fmt.Println("Part One:", d.findMatches(gridSize, antennas))
	fmt.Println("Part Two:", d.part2FindMatches(gridSize, antennas))
}
