package pipefilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("Input data should be string")

// SplitFilter 过滤字符串
type SplitFilter struct {
	delimiter string
}

// NewSplitFilter 快捷构造函数
func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

// Process 处理数据
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
