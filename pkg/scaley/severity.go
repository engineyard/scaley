package scaley

type Severity int

const (
	Okay Severity = iota
	Warning
	Failure
)
