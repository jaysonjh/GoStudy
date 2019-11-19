package series

import "errors"

var LessThanError = errors.New("n shoulb be greater than 2")
var LargeThanError = errors.New("n shoulb be less than 100")

func GetFibonacciSeries(n int) ([]int, error) {

	if n < 2 {
		return nil, LessThanError
	}

	if n > 100 {
		return nil, LargeThanError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
