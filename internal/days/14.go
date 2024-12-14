package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type DayFourteen struct{}

type d14_Robot struct {
	position [2]int
	velocity [2]int
}

func (r *d14_Robot) move(gSize [2]int) {
	r.position[0] += r.velocity[0]
	r.position[1] += r.velocity[1]

	if r.position[0] < 0 {
		r.position[0] += gSize[0]
	}

	if r.position[0] >= gSize[0] {
		r.position[0] -= gSize[0]
	}

	if r.position[1] < 0 {
		r.position[1] += gSize[1]
	}

	if r.position[1] >= gSize[1] {
		r.position[1] -= gSize[1]
	}
}

var R = regexp.MustCompile(`(\d+,\d+)\sv=([-]?\d+,[-]?\d+)`)

func (d *DayFourteen) processInput(file *os.File) []*d14_Robot {
	scanner := bufio.NewScanner(file)
	robots := []*d14_Robot{}

	for scanner.Scan() {
		l := scanner.Text()

		for _, match := range R.FindAllStringSubmatch(l, -1) {
			pos, vel := strings.Split(match[1], ","), strings.Split(match[2], ",")
			robots = append(robots, &d14_Robot{
				position: [2]int{utils.StringToInt(pos[0]), utils.StringToInt(pos[1])},
				velocity: [2]int{utils.StringToInt(vel[0]), utils.StringToInt(vel[1])},
			})
		}
	}

	return robots
}

func (d *DayFourteen) simulateTime(gSize [2]int, sec int, robots []*d14_Robot) {
	for s := range sec {
		locations := map[[2]int]bool{}
		for _, r := range robots {
			r.move(gSize)
			locations[r.position] = true
		}

		if len(locations) == len(robots) {
			fmt.Println("Sec", s+1)
			d.visualize(robots, gSize)
			fmt.Println()
			break
		}
	}
}

func (d *DayFourteen) qtyPerQuadrant(robots []*d14_Robot, gSize [2]int) [4]int {
	hW, hH := gSize[0]/2, gSize[1]/2
	qty := [4]int{}

	for _, r := range robots {
		if r.position[0] < hW && r.position[1] < hH {
			qty[0] += 1
		} else if r.position[0] > hW && r.position[1] < hH {
			qty[1] += 1
		} else if r.position[0] < hW && r.position[1] > hH {
			qty[2] += 1
		} else if r.position[0] > hW && r.position[1] > hH {
			qty[3] += 1
		}
	}

	return qty
}

func (d *DayFourteen) safeFactor(robots []*d14_Robot, gSize [2]int, seconds int) int {
	d.simulateTime(gSize, seconds, robots)
	qdts := d.qtyPerQuadrant(robots, gSize)
	return qdts[0] * qdts[1] * qdts[2] * qdts[3]
}

func (d *DayFourteen) visualize(robots []*d14_Robot, gSize [2]int) {
	g := [][]int{}
	for range gSize[1] {
		g = append(g, make([]int, gSize[0]))
	}

	for _, r := range robots {
		g[r.position[1]][r.position[0]] += 1
	}

	for i := range gSize[1] {
		for j := range gSize[0] {
			if g[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func (d *DayFourteen) Run() {
	gSize := [2]int{101, 103}
	if os.Getenv("TEST") == "1" {
		gSize = [2]int{11, 7}
	}
	file := utils.ReadInput(14, 1)
	defer file.Close()
	robots := d.processInput(file)
	fmt.Println("Safe Factor:", d.safeFactor(robots, gSize, 100))
	fmt.Println("Tree", d.safeFactor(robots, gSize, 10000))
}
