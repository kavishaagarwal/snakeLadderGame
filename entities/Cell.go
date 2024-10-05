package entities

type Cell struct {
	Jump *Jump
}

// Getter for Jump
func (c *Cell) GetJump() *Jump {
	return c.Jump
}

// Setter for Jump
func (c *Cell) SetJump(jump *Jump) {
	c.Jump = jump
}
