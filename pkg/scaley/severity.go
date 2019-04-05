package scaley

// Severity describes the importance of a log message.
type Severity int

const (
	// Okay is used for success logs.
	Okay Severity = iota
	// Warning is used for information logs.
	Warning
	// Failure is used for failure logs.
	Failure
)
