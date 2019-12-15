package pipefilter

import "errors"

var SumFilterWrongFormatError = errors.New("Input data should be []int")

// SumFilter 统计Filter
type SumFilter struct {
}

// NewSumIntFilter 快捷构造函数
func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// Process 处理数据
func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}

	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
