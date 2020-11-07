package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Request data interface
type Request interface {}

// Response data interface
type Response interface{}

// Filter interface
type Filter interface {
	// GetName get name of filter.
	GetName() string

	// Process process request and return response.
	Process(data Request) (Response, error)
}

// Pipe interface
type Pipe interface {
	// GetName get name of pipe.
	GetName() string

	// Run start run filters.
	Run(filters ...Filter) (Response, error)
}

type examplePipe struct {}

func (e *examplePipe) GetName() string {
	return "examplePipe"
}

func (e *examplePipe) Run(filters ...Filter) (Response, error) {
	var (
		request Request
		response Response
		err error
	)
	request = "1,2,3,4,5"
	for _, filter := range filters {
		if response, err = filter.Process(request); err != nil {
			return nil, err
		}
		fmt.Println(filter.GetName(), response)
		request = response
	}

	return response, nil
}

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

func main()  {
	splitFilter := SplitFilter{}
	sumFilter := SumFilter{}
	pipe := examplePipe{}

	if response, err := pipe.Run(splitFilter, sumFilter); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
