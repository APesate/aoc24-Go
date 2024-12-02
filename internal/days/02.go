package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mapInput(line []string) []int {
	result := make([]int, len(line))

	for i, s := range line {
		num, _ := strconv.Atoi(s)
		result[i] = num
	}

	return result
}

func processInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	input := make([]([]int), 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		input = append(input, mapInput(values))
	}

	return input
}

func validDistance(lhs, rhs int) bool {
	distance := utils.Distance(lhs, rhs)
	return distance >= 1 && distance <= 3
}

func isValidSequence(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	isInc := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		if !validDistance(nums[i-1], nums[i]) {
			return false
		}
		if (isInc && nums[i-1] > nums[i]) || (!isInc && nums[i-1] < nums[i]) {
			return false
		}
	}
	return true
}

func partOne(input [][]int) int {
	count := 0

	for _, report := range input {
		if !isValidSequence(report) {
			continue
		}

		count++
	}

	return count
}

func partTwo(input [][]int) int {
	count := 0

	for _, report := range input {
		if isValidSequence(report) {
			count++
			continue
		}

		isValid := false
		for i := 0; i < len(report); i++ {
			newReport := make([]int, 0, len(report)-1)
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)

			if isValidSequence(newReport) {
				isValid = true
				break
			}
		}

		if isValid {
			count++
		}
	}

	return count
}

func DayTwo() {
	file := utils.ReadInput(2, 1)
	input := processInput(file)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
