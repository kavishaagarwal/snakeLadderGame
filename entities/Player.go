package entities

type Player struct {
	ID              string
	CurrentPosition int
}

// Constructor for Player
func NewPlayer(id string, currentPosition int) *Player {
	return &Player{
		ID:              id,
		CurrentPosition: currentPosition,
	}
}

// Getter for ID
func (p *Player) GetID() string {
	return p.ID
}

// Setter for ID
func (p *Player) SetID(id string) {
	p.ID = id
}

// Getter for CurrentPosition
func (p *Player) GetCurrentPosition() int {
	return p.CurrentPosition
}

// Setter for CurrentPosition
func (p *Player) SetCurrentPosition(position int) {
	p.CurrentPosition = position
}
