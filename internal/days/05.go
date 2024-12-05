package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DayFive struct{}

func (d *DayFive) processInput(file *os.File) (map[string][]string, [][]string) {
	scanner := bufio.NewScanner(file)
	section := 0
	rules := map[string][]string{}
	pages := [][]string{}

	s0 := func(arg string) {
		parts := strings.Split(arg, "|")
		r0, r1 := parts[0], parts[1]
		rules[r0] = append(rules[r0], r1)

	}

	s1 := func(arg string) {
		pages = append(pages, strings.Split(arg, ","))
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			section = 1
			continue
		}

		switch section {
		case 0:
			s0(line)
		case 1:
			s1(line)
		}
	}

	return rules, pages
}

func (d *DayFive) isValidPage(page []string, rules map[string][]string) bool {
	isValid := true

	for i, v := range page {
		for _, o := range page[i:] {
			if slices.Contains(rules[o], v) {
				isValid = false
				break
			}
		}

		if !isValid {
			break
		}
	}

	return isValid
}

func (d *DayFive) partOne(rules map[string][]string, pages [][]string) int {
	validPages := utils.Reduce(pages, [][]string{}, func(acc [][]string, page []string) [][]string {
		if d.isValidPage(page, rules) {
			return append(acc, page)
		}

		return acc
	})

	return utils.Reduce(validPages, 0, func(acc int, slice []string) int {
		n, _ := strconv.Atoi(slice[len(slice)/2])
		return acc + n
	})
}

func (d *DayFive) partTwo(rules map[string][]string, pages [][]string) int {
	invalidPages := utils.Reduce(pages, [][]string{}, func(acc [][]string, page []string) [][]string {
		if !d.isValidPage(page, rules) {
			return append(acc, page)
		}

		return acc
	})

	for _, p := range invalidPages {
		slices.SortFunc(p, func(lhs, rhs string) int {
			if slices.Contains(rules[rhs], lhs) {
				return 1
			}

			return -1
		})
	}

	return utils.Reduce(invalidPages, 0, func(acc int, slice []string) int {
		n, _ := strconv.Atoi(slice[len(slice)/2])
		return acc + n
	})
}

func (d *DayFive) Run() {
	file := utils.ReadInput(5, 1)
	defer file.Close()
	r, p := d.processInput(file)
	// fmt.Printf("P1: %d\n", d.partOne(r, p))
	fmt.Printf("P2: %d\n", d.partTwo(r, p))
}
