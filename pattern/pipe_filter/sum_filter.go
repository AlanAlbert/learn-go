package pipe_filter

import "errors"

// SumFilter get num of array or slice.
type SumFilter struct {}

// GetName get name of SumFilter.
func (s SumFilter) GetName() string {
	return "SumFilter"
}

// Process process
func (s SumFilter) Process(data Request) (Response, error) {
	if v, ok := data.([]int); !ok {
		return nil, errors.New("data error")
	} else {
		res := 0
		for _, item := range v {
			res += item
		}
		return res, nil
	}
}
