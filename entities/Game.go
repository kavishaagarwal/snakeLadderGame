package entities

import (
	"container/list"
	"encoding/json"
	"fmt"
)

type Game struct {
	board       *Board
	dice        *Dice
	playersList *list.List
	winner      *Player
}

func NewGame() *Game {
	game := &Game{}
	game.initializeGame()
	return game
}

func (g *Game) initializeGame() {
	g.board = NewBoard(10, 5, 4)
	for _, row := range g.board.Cells {
		ss, _ := json.Marshal(row)
		fmt.Println(string(ss))
	}
	g.dice = NewDice(2)
	g.winner = nil
	g.addPlayers()
}

func (g *Game) addPlayers() {
	player1 := NewPlayer("p1", 0)
	player2 := NewPlayer("p2", 0)
	g.playersList = list.New()
	g.playersList.PushBack(player1)
	g.playersList.PushBack(player2)
}

func (g *Game) StartGame() {
	for g.winner == nil {
		playerTurn := g.findPlayerTurn()
		fmt.Printf("Player turn is: %s, current position is: %d\n", playerTurn.GetID(), playerTurn.GetCurrentPosition())

		// Roll the dice
		diceNumbers := g.dice.RollDice()

		// Get the new position
		playerNewPosition := playerTurn.GetCurrentPosition() + diceNumbers
		playerNewPosition = g.jumpCheck(playerNewPosition)
		playerTurn.SetCurrentPosition(playerNewPosition)

		fmt.Printf("Player turn is: %s, new position is: %d\n", playerTurn.GetID(), playerNewPosition)

		// Check for winning condition
		if playerNewPosition >= len(g.board.Cells)*len(g.board.Cells[0])-1 {
			g.winner = playerTurn
		}
	}

	fmt.Printf("WINNER IS: %s\n", g.winner.GetID())
}

func (g *Game) findPlayerTurn() *Player {
	element := g.playersList.Remove(g.playersList.Front())
	g.playersList.PushBack(element)
	return element.(*Player)
}

func (g *Game) jumpCheck(playerNewPosition int) int {
	if playerNewPosition >= len(g.board.Cells)*len(g.board.Cells[0]) {
		return playerNewPosition
	}
	fmt.Println("playerNewPosition:", playerNewPosition)
	cell := g.board.getCell(playerNewPosition)
	if cell.GetJump() != nil && cell.Jump.GetStart() == playerNewPosition {
		jumpBy := "snake"
		if cell.Jump.GetStart() < cell.Jump.GetEnd() {
			jumpBy = "ladder"
		}
		fmt.Printf("Jump done by: %s\n", jumpBy)
		return cell.Jump.GetStart()
	}
	return playerNewPosition
}
