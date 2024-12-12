package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DayEleven struct{}
type d11_Stone struct {
	engraving string
}

func (s *d11_Stone) blink(count int) map[string]int {
	result := make(map[string]int)

	if s.engraving == "0" {
		result["1"] += count
	} else if len(s.engraving)%2 == 0 {
		h := len(s.engraving) / 2
		lhs, rhs := s.engraving[:h], s.engraving[h:]
		lhn, _ := strconv.Atoi(lhs)
		rhn, _ := strconv.Atoi(rhs)

		result[strconv.Itoa(lhn)] += count
		result[strconv.Itoa(rhn)] += count
	} else {
		n, _ := strconv.Atoi(s.engraving)
		newEngraving := strconv.Itoa(n * 2024)
		result[newEngraving] += count
	}

	return result
}

func (d *DayEleven) processInput(file *os.File) []d11_Stone {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stones := []d11_Stone{}

	for _, s := range strings.Split(scanner.Text(), " ") {
		stones = append(stones, d11_Stone{engraving: s})
	}

	return stones
}

func (d *DayEleven) performBlinks(blinks int, input []d11_Stone) int {
	stoneCount := make(map[string]int)

	for _, stone := range input {
		stoneCount[stone.engraving]++
	}

	for i := 0; i < blinks; i++ {
		nextStoneCount := make(map[string]int)

		for engraving, count := range stoneCount {
			s := d11_Stone{engraving: engraving}
			blinksResult := s.blink(count)
			for newEngraving, newCount := range blinksResult {
				nextStoneCount[newEngraving] += newCount
			}
		}

		stoneCount = nextStoneCount
	}

	totalStones := 0
	for _, count := range stoneCount {
		totalStones += count
	}

	return totalStones
}

func (d *DayEleven) Run() {
	file := utils.ReadInput(11, 1)
	defer file.Close()
	input := d.processInput(file)
	fmt.Println("Part One: ", d.performBlinks(25, input))
	fmt.Println("Part Two: ", d.performBlinks(75, input))
}
