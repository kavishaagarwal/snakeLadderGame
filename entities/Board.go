package entities

import (
	"math/rand"
	"time"
)

type Board struct {
	Cells [][]Cell
}

func NewBoard(boardSize int, numberOfSnakes int, numberOfLadders int) *Board {
	board := &Board{}
	board.initializeCells(boardSize)
	board.addSnakesLadders(numberOfSnakes, numberOfLadders)
	return board
}

func (b *Board) initializeCells(boardSize int) {
	b.Cells = make([][]Cell, boardSize)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, boardSize)
		for j := range b.Cells[i] {
			b.Cells[i][j] = Cell{}
		}
	}
}

func (b *Board) addSnakesLadders(numberOfSnakes int, numberOfLadders int) {
	rand.Seed(time.Now().UnixNano())

	// Add snakes
	for numberOfSnakes > 0 {
		snakeHead := rand.Intn(len(b.Cells)*len(b.Cells)-1) + 1
		snakeTail := rand.Intn(len(b.Cells)*len(b.Cells)-1) + 1
		if snakeTail >= snakeHead {
			continue
		}

		snakeObj := NewJump(snakeHead, snakeTail)

		cell := b.getCell(snakeHead)
		cell.SetJump(snakeObj)

		numberOfSnakes--
	}

	// Add ladders
	for numberOfLadders > 0 {
		ladderStart := rand.Intn(len(b.Cells)*len(b.Cells)-1) + 1
		ladderEnd := rand.Intn(len(b.Cells)*len(b.Cells)-1) + 1
		if ladderStart >= ladderEnd {
			continue
		}

		ladderObj := NewJump(ladderStart, ladderEnd)

		cell := b.getCell(ladderStart)
		cell.SetJump(ladderObj)

		numberOfLadders--
	}
}

func (b *Board) getCell(playerPosition int) *Cell {
	boardRow := playerPosition / len(b.Cells)
	boardColumn := playerPosition % len(b.Cells)
	return &b.Cells[boardRow][boardColumn]
}
