package pipefilter

func NewStraighPipeline(name string, filters ...Filter) *StraighPipeline {
	return &StraighPipeline{Name: name, Filters: &filters}
}

type StraighPipeline struct {
	Name    string
	Filters *[]Filter
}

// Process 处理数据
func (sp *StraighPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *sp.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
