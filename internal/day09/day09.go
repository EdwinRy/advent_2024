package day09

import (
	"fmt"
	"strconv"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

func strToIntArr(input string) []int {
	arr := make([]int, 0)
	for _, r := range input {
		num, _ := strconv.Atoi(string(r))
		arr = append(arr, num)
	}
	return arr
}

func genBlock(size int, value int) []int {
	block := make([]int, size)
	for i := 0; i < size; i++ {
		block[i] = value
	}
	return block
}

func generateDisk(description []int) []int {
	disk := make([]int, 0)

	readingFile := true
	fileIdx := 0
	for i := 0; i < len(description); i++ {

		blockSize := description[i]
		var writeChar int

		if readingFile {
			writeChar = fileIdx
			fileIdx++
		} else {
			writeChar = -1
		}

		disk = append(disk, genBlock(blockSize, writeChar)...)
		readingFile = !readingFile
	}

	return disk
}

func calculateDiskChecksum(disk []int) int {
	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}
		checksum += disk[i] * i
	}
	return checksum
}

func task1(input string) (int, error) {
	inputArr := strToIntArr(input)
	disk := generateDisk(inputArr)

	ptr1, ptr2 := 0, len(disk)-1

	for ptr1 < ptr2 {
		if disk[ptr2] == -1 {
			ptr2--
			continue
		}
		if disk[ptr1] == -1 {
			disk[ptr1] = disk[ptr2]
			disk[ptr2] = -1
			ptr1++
			ptr2--
			continue
		}
		ptr1++
	}

	return calculateDiskChecksum(disk), nil
}

type Block struct {
	diskIdx int
	size    int
	fileId  int
}

func generateDiskBlocks(description []int) []Block {
	disk := make([]Block, 0)

	readingFile := true
	fileIdx := 0
	position := 0
	for i := 0; i < len(description); i++ {

		blockSize := description[i]
		var writeChar int

		if readingFile {
			writeChar = fileIdx
			fileIdx++
		} else {
			writeChar = -1
		}

		disk = append(disk, Block{size: blockSize, fileId: writeChar, diskIdx: position})
		position += blockSize
		readingFile = !readingFile
	}

	return disk
}

func blocksChecksum(blocks []Block) int {
	checksum := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i].fileId == -1 {
			continue
		}
		for j := 0; j < blocks[i].size; j++ {
			checksum += blocks[i].fileId * (blocks[i].diskIdx + j)
		}
	}
	return checksum
}

func blocksToString(blocks []Block) string {
	str := ""
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].size; j++ {
			if blocks[i].fileId == -1 {
				str += "."
			} else {
				str += strconv.Itoa(blocks[i].fileId)
			}
		}
	}
	return str
}

func findEmptyBlock(blocks []Block, size int, maxIdx int) int {
	for i := 0; i < maxIdx; i++ {
		if blocks[i].fileId == -1 && blocks[i].size >= size {
			return i
		}
	}
	return -1
}

func task2(input string) (int, error) {
	inputArr := strToIntArr(input)
	blocks := generateDiskBlocks(inputArr)

	for ptr2 := len(blocks) - 1; ptr2 > 0; ptr2-- {
		if blocks[ptr2].fileId == -1 {
			continue
		}

		ptr1 := findEmptyBlock(blocks, blocks[ptr2].size, ptr2)
		if ptr1 == -1 {
			continue
		}

		blocks[ptr1].fileId = blocks[ptr2].fileId
		blocks[ptr2].fileId = -1

		if blocks[ptr1].size != blocks[ptr2].size {
			diff := blocks[ptr1].size - blocks[ptr2].size
			blocks[ptr1].size = blocks[ptr2].size
			proceedingEmptySpace := Block{
				size:    diff,
				fileId:  -1,
				diskIdx: blocks[ptr1].diskIdx + blocks[ptr1].size,
			}
			blocks = utils.SliceInsert(blocks, ptr1+1, proceedingEmptySpace)
		}

	}

	return blocksChecksum(blocks), nil
}

func Day09() {
	input, _ := utils.ReadFile("inputs/day09/input.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 09 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 09 task 2: ", task2Result)
}
