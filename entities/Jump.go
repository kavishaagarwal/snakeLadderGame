package entities

type Jump struct {
	Start int
	End   int
}

func NewJump(start, end int) *Jump {
	return &Jump{start, end}
}

// Getter for Start
func (j *Jump) GetStart() int {
	return j.Start
}

// Setter for Start
func (j *Jump) SetStart(start int) {
	j.Start = start
}

// Getter for End
func (j *Jump) GetEnd() int {
	return j.End
}

// Setter for End
func (j *Jump) SetEnd(end int) {
	j.End = end
}
