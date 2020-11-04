package pipe_filter

import (
	"errors"
	"strings"
)

// SplitFilter 分割Filter
type SplitFilter struct {}

// GetName get name of filter.
func (s SplitFilter) GetName() string {
	return "SplitFilter"
}

// Process process
func (s SplitFilter) Process(data Request) (Response, error) {
	if v, ok := data.(string); !ok {
		return nil, errors.New("data error")
	} else {
		return strings.Split(v, ","), nil
	}
}

