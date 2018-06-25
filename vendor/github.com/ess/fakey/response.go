package fakey

import (
	"errors"
)

type ResponseCollection struct {
	responses map[string][]string
}

func (rc *ResponseCollection) Add(method, path, response string) {
	identifier := rc.identify(method, path)

	rc.setup(identifier)

	rc.responses[identifier] = append(rc.responses[identifier], response)
}

func (rc *ResponseCollection) Remove(method, path string) {
	identifier := rc.identify(method, path)

	rc.responses[identifier] = nil

	rc.setup(identifier)
}

func (rc *ResponseCollection) Consume(method, path string) (string, error) {
	var response string

	rc.setup("")

	identifier := rc.identify(method, path)

	if len(rc.responses[identifier]) == 0 {
		return response, errors.New("No response")
	}

	response = rc.responses[identifier][0]
	rc.trim(identifier)

	return response, nil
}

func (rc *ResponseCollection) trim(identifier string) {
	if len(rc.responses[identifier]) == 1 {
		rc.responses[identifier] = nil
	} else {
		rc.responses[identifier] = rc.responses[identifier][1:]
	}
}

func (rc *ResponseCollection) identify(method, path string) string {
	return method + ":" + path
}

func (rc *ResponseCollection) setup(scope string) {
	if rc.responses == nil {
		rc.responses = make(map[string][]string)
	}

	if len(scope) > 0 && rc.responses[scope] == nil {
		rc.responses[scope] = make([]string, 0)
	}
}
