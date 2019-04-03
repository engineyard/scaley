package scaley

type Direction int

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

func (d Direction) DesiredState() string {
	if d == Up {
		return "stopped"
	}

	return "running"
}

const (
	None Direction = iota
	Down
	Up
)
