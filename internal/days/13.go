package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DayThirteen struct{}

type d13_Prize struct {
	x int
	y int
}

type d13_Machine struct {
	btn1  d13_Btn
	btn2  d13_Btn
	prize d13_Prize
}

type d13_Btn struct {
	x int
	y int
}

var Rgx = regexp.MustCompile(`.+?(\d+).+?(\d+)`)

func (d *DayThirteen) processRgx(l string) (int, int) {
	var c1, c2 int
	for _, match := range Rgx.FindAllStringSubmatch(l, 1) {
		c1, _ = strconv.Atoi(match[1])
		c2, _ = strconv.Atoi(match[2])
	}

	return c1, c2
}

func (d *DayThirteen) processInput(file *os.File) []d13_Machine {
	scanner := bufio.NewScanner(file)
	machines := []d13_Machine{{}}

	for scanner.Scan() {
		l := scanner.Text()

		if strings.Contains(l, "Button") {
			c1, c2 := d.processRgx(l)
			if (machines[len(machines)-1].btn1 == d13_Btn{}) {
				machines[len(machines)-1].btn1.x = c1
				machines[len(machines)-1].btn1.y = c2
			} else {
				machines[len(machines)-1].btn2.x = c1
				machines[len(machines)-1].btn2.y = c2
			}
		} else if strings.Contains(l, "Prize") {
			x, y := d.processRgx(l)
			machines[len(machines)-1].prize.x = x
			machines[len(machines)-1].prize.y = y
		} else {
			machines = append(machines, d13_Machine{})
		}
	}

	return machines
}

func (m d13_Machine) winingCombo() (int, int) {
	c1 := m.btn1.x
	c2 := m.btn2.x
	d1 := m.btn1.y
	d2 := m.btn2.y
	A := m.prize.x
	B := m.prize.y

	b := (c1*B - d1*A) / (c1*d2 - d1*c2)
	a := (A - c2*b) / c1
	return a, b
}

func (d *DayThirteen) partOne(machines []d13_Machine) int {
	aCost, bCost, total := 3, 1, 0
	for _, m := range machines {
		a, b := m.winingCombo()

		if a > 100 || b > 100 {
			continue
		}

		if m.btn1.x*a+m.btn2.x*b != m.prize.x ||
			m.btn1.y*a+m.btn2.y*b != m.prize.y {
			continue
		}

		total += aCost*a + bCost*b
	}

	return total
}

func (d *DayThirteen) partTwo(machines []d13_Machine) int {
	aCost, bCost, total := 3, 1, 0
	for _, m := range machines {
		m.prize.x += 10000000000000
		m.prize.y += 10000000000000

		a, b := m.winingCombo()

		if m.btn1.x*a+m.btn2.x*b != m.prize.x ||
			m.btn1.y*a+m.btn2.y*b != m.prize.y {
			continue
		}

		total += aCost*a + bCost*b
	}

	return total
}

func (d *DayThirteen) Run() {
	file := utils.ReadInput(13, 1)
	defer file.Close()
	machines := d.processInput(file)
	fmt.Println("Part One:", d.partOne(machines))
	fmt.Println("Part Two:", d.partTwo(machines))
}
