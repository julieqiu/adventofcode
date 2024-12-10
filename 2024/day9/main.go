package main

import (
	"log"
	"strconv"

	"github.com/julieqiu/adventofcode/2024/internal/runner"
)

func main() {
	if err := runner.Run(problem1, problem2); err != nil {
		log.Fatal(err)
	}
}

func problem1(lines []string) (int, error) {
	blocks := removeFreeSpace(lines[0])
	return calculateChecksum(blocks), nil
}

func problem2(lines []string) (int, error) {
	freed := removeFreeSpacePart2(lines[0])
	return calculateChecksum(freed), nil
}

func calculateChecksum(blocks []string) int {
	var result int
	for i, b := range blocks {
		if b == "." {
			continue
		}
		n, _ := strconv.Atoi(b)
		result += i * n
	}
	return result
}

func removeFreeSpace(diskmap string) []string {
	blocks := diskmapToBlocks(diskmap)
	j := len(blocks) - 1
	for i := 0; i < len(blocks); i++ {
		if j <= i {
			break
		}
		if blocks[i] == "." {
			blocks[i] = blocks[j]
			blocks[j] = "."
			for blocks[j] == "." && j > i {
				j--
			}
		}
	}
	return blocks
}

func diskmapToBlocks(diskmap string) []string {
	var (
		blocks []string
		count  = 0
		block  = true
	)
	for _, digit := range diskmap {
		num, _ := strconv.Atoi(string(digit))
		c := "."
		if block {
			c = strconv.Itoa(count)
		}

		for i := 0; i < num; i++ {
			blocks = append(blocks, c)
		}

		if block {
			count += 1
		}
		block = !block
	}
	return blocks
}

func removeFreeSpacePart2(diskmap string) []string {
	blocks := diskmapToBlocks(diskmap)

	start := len(blocks)
	val := -1
	var end, fstart, fend int
	for {
		start, end, val = findBlock(val-1, start, blocks)
		if start == -1 {
			break
		}
		// fmt.Println(start, end, val)
		for {
			if fstart == -1 {
				break
			}
			if fend >= start {
				break
			}
			fstart, fend = findSpace(fend, blocks)
			// fmt.Println(start, end, val, blocks[start:end], fstart, fend, blocks[fstart:fend])
			if (fend - fstart) >= (end - start) {
				for i := 0; i < (end - start); i++ {
					// Replace "." with character.
					blocks[fstart+i] = blocks[start]
				}
				for i := start; i < end; i++ {
					blocks[i] = "."
				}
				// fmt.Println(blocks)
				break
			}
		}
		fstart = 0
		fend = 0
	}
	return blocks
}

func findBlock(val, end int, blocks []string) (int, int, int) {
	start := end - 1

	if val == -2 {
		for {
			if blocks[start] != "." {
				break
			}
			start--
		}
		val, _ = strconv.Atoi(blocks[start])
	}

	// Keep going until we have found val.
	for {
		if start < 0 {
			// We have reached the end of the line, so there is nothing to
			// place.
			return -1, -1, -1
		}
		c, _ := strconv.Atoi(blocks[start])
		if c == val {
			break
		}
		start--
	}

	// Place end right after that character.
	end = start + 1

	// Keep going until we see something different.
	for {
		if start < 0 {
			break
		}
		c, _ := strconv.Atoi(blocks[start])
		if c != val {
			break
		}
		start--
	}
	return start + 1, end, val
}

func findSpace(start int, blocks []string) (int, int) {
	for {
		if start >= len(blocks) {
			return -1, -1
		}
		if blocks[start] == "." {
			break
		}
		start += 1
	}

	end := start + 1
	for {
		if end >= len(blocks) {
			return start, end
		}
		if blocks[end] != "." {
			break
		}
		end += 1
	}
	return start, end
}
