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
