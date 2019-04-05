package scaley

// Direction describes the desired direction (up, down, none) for a scaling
// event.
type Direction int

// String returns a text representation of a direction.
func (d Direction) String() string {
	switch d {
	case Up:
		return "UP"
	case Down:
		return "DOWN"
	default:
		return "NONE"
	}
}

// DesiredState returns the state that a server should be in for scaling to
// be considered for a given direction.
func (d Direction) DesiredState() string {
	if d == Up {
		return "stopped"
	}

	return "running"
}

const (
	// None indicates that no scaling event should occur.
	None Direction = iota
	// Down indicates that the group should be scaled down.
	Down
	// Up indicates that the group should be scaled up.
	Up
)
