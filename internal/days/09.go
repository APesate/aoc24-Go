package days

import (
	"aoc24/internal/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

type DayNine struct{}

type d9_MemoryLayout struct {
	id  int
	qty int
}

func (d *DayNine) processInput(file *os.File) ([]d9_MemoryLayout, []d9_MemoryLayout) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	files := []d9_MemoryLayout{}
	space := []d9_MemoryLayout{}

	for i, c := range scanner.Text() {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			files = append(files, d9_MemoryLayout{id: i - len(space), qty: n})
		} else {
			space = append(space, d9_MemoryLayout{id: 0, qty: n})
		}
	}

	return files, space
}

func (d *DayNine) rearrange(files []d9_MemoryLayout, spaces []d9_MemoryLayout) []d9_MemoryLayout {
	finalLayout := make([]d9_MemoryLayout, 0)
	finalLayout = append(finalLayout, files[0])
	fileCursor := 1

	for _, space := range spaces {
		spaceAvailable := space.qty

		for spaceAvailable > 0 {
			lastFile := files[len(files)-1]
			requiredSpace := min(spaceAvailable, lastFile.qty)
			spaceAvailable -= requiredSpace
			finalLayout = append(finalLayout, d9_MemoryLayout{id: lastFile.id, qty: requiredSpace})
			lastFile.qty -= requiredSpace

			if lastFile.qty == 0 {
				files = files[:len(files)-1]
			} else {
				files[len(files)-1] = lastFile
			}
		}

		if len(files) > fileCursor {
			finalLayout = append(finalLayout, files[fileCursor])
		}

		fileCursor += 1

		if fileCursor >= len(files) {
			break
		}
	}

	return finalLayout
}

func (d *DayNine) rearrangeFullSize(files []d9_MemoryLayout, spaces []d9_MemoryLayout) []d9_MemoryLayout {
	finalLayout := make([]d9_MemoryLayout, len(files)+len(spaces))

	for i := range len(files) + len(spaces) {
		if i%2 == 0 {
			finalLayout[i] = files[i/2]
		} else {
			finalLayout[i] = spaces[int(math.Ceil(float64(i/2)))]
		}
	}

	findFile := func(availableSpace int, files []d9_MemoryLayout, fromIndex int) (int, d9_MemoryLayout, bool) {
		for i := len(files) - 1; i > fromIndex; i-- {
			if files[i].qty <= availableSpace && files[i].id != 0 {
				return i, files[i], true
			}
		}

		return -1, d9_MemoryLayout{}, false
	}

	for i := 1; i < len(finalLayout); i += 1 {
		cursor := finalLayout[i]
		if cursor.id != 0 || cursor.qty == 0 {
			continue
		}

		if index, fileToMove, ok := findFile(cursor.qty, finalLayout, i); ok {
			remainingSpace := cursor.qty - fileToMove.qty
			cursor.id = fileToMove.id
			cursor.qty = fileToMove.qty
			finalLayout[i] = cursor
			finalLayout[index].id = 0

			if remainingSpace > 0 {
				finalLayout = slices.Insert(finalLayout, i+1, d9_MemoryLayout{id: 0, qty: remainingSpace})
			}
		}
	}

	return finalLayout
}

func (d *DayNine) checksum(files []d9_MemoryLayout) int {
	total := 0
	index := 0

	for _, file := range files {
		for range file.qty {
			total += file.id * index
			index += 1
		}
	}

	return total
}

func (d *DayNine) Run() {
	file := utils.ReadInput(9, 1)
	defer file.Close()
	files, spaces := d.processInput(file)
	// fmt.Println("Part One:", d.checksum(d.rearrange(files, spaces)))
	fmt.Println("Part One:", d.checksum(d.rearrangeFullSize(files, spaces)))
}
