package pipefilter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("Input data should be []string")

// ToIntFilter 转换字符串为IntFilter
type ToIntFilter struct {
}

// NewToIntFilter 快捷构造函数
func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

// Process 处理数据
func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parks, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}

	ret := []int{}
	for _, park := range parks {
		s, err := strconv.Atoi(park)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
