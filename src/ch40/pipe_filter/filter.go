package pipefilter

// Request is the input filter
type Request interface{}

// Response is the output filter
type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}
