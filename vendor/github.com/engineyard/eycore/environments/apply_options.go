package environments

import (
	"fmt"
)

type ApplyOptions struct {
	Type string
}

func (options *ApplyOptions) SetType(runType string) {
	options.Type = runType
}

func (options *ApplyOptions) Body() []byte {
	if len(options.Type) > 0 {
		return []byte(fmt.Sprintf(`{"type":"%s"}`, options.Type))
	}

	return nil
}
