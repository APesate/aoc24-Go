package utils

import (
	"fmt"
	"log/slog"
	"math"
	"os"
)

func ReadInput(day, part int) *os.File {
	suffix := "input"

	if os.Getenv("TEST") == "1" {
		suffix = "test"
	}

	fileName := fmt.Sprintf("./inputs/%02d-%d.%s", day, part, suffix)
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error(err.Error())
	}

	return file
}

func Distance(lhs, rhs int) int {
	return int(math.Abs(float64(lhs) - float64(rhs)))
}

// Reducer is a type for reduce function signatures
type Reducer[T, R any] func(R, T) R

// Reduce is a generic function that reduces a slice using a reducer function
func Reduce[T, R any](slice []T, initial R, reducer Reducer[T, R]) R {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func IsRowInBounds[T any](grid [][]T, row int) bool {
	return row >= 0 && row < len(grid)
}

func IsColInBounds[T any](grid [][]T, col int) bool {
	return col >= 0 && col < len(grid[0])
}

func CopyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
