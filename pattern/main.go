package main

import (
	"fmt"
	"pattern/pipe_filter"
)

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

func testFilter()  {
	splitFilter := SplitFilter{}
	sumFilter := SumFilter{}
	pipe := examplePipe{}

	if response, err := pipe.Run(splitFilter, sumFilter); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
