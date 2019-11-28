package cancle_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(ch chan struct{}) bool {

	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancelChannel(t *testing.T) {
	cancelCh := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelCh)
	}
	cancel_1(cancelCh)
	time.Sleep(time.Second * 1)

}
