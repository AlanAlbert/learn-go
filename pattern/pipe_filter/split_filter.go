package pipe_filter

import (
	"errors"
	"strconv"
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
		items := strings.Split(v, ",")
		res := make([]int, len(items))
		var err error
		for i, v := range items {
			if res[i], err = strconv.Atoi(v); err != nil {
				return nil, err
			}
		}
		return res, nil
	}
}