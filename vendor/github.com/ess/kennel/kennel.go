package kennel

import (
	"github.com/DATA-DOG/godog"
)

type Stepper interface {
	StepUp(*godog.Suite)
}

var steppers []Stepper

func Register(s Stepper) {
	steppers = append(steppers, s)
}

func StepUp(s *godog.Suite) {
	for _, stepper := range steppers {
		stepper.StepUp(s)
	}
}
