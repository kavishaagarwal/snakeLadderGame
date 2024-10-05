package entities

import (
	"math/rand"
	"time"
)

type Dice struct {
	diceCount int
	//min       int
	//max       int
}

func NewDice(diceCount int) *Dice {
	return &Dice{
		diceCount: diceCount,
		//min:       1,
		//max:       6,
	}
}

func (d *Dice) RollDice() int {
	totalSum := 0
	rand.Seed(time.Now().UnixNano())
	for diceUsed := 0; diceUsed < d.diceCount; diceUsed++ {
		totalSum += rand.Intn(6) + 1
	}
	return totalSum
}
