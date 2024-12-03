package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type DayThree struct{}

type instruction struct {
	lhs int
	rhs int
}

// var r = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)                  // Part One
var r = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`) // Part Two

func (d *DayThree) processInput(file *os.File) []instruction {
	scanner := bufio.NewScanner(file)
	instructions := make([]instruction, 0)
	enabled := true

	for scanner.Scan() {
		for _, match := range r.FindAllStringSubmatch(scanner.Text(), -1) {

			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			}

			if !enabled {
				continue
			}

			lhs, _ := strconv.Atoi(match[1])
			rhs, _ := strconv.Atoi(match[2])

			instructions = append(instructions, instruction{
				lhs,
				rhs,
			})
		}

	}

	return instructions
}

func (d *DayThree) calculate(input []instruction) int {
	total := 0
	for _, inst := range input {
		total += inst.lhs * inst.rhs
	}
	return total
}

func (d *DayThree) Run() {
	file := utils.ReadInput(3, 1)
	defer file.Close()
	input := d.processInput(file)

	fmt.Println(d.calculate(input))
}
