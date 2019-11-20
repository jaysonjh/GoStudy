package async_service

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "Done"
}

func otherTask() {
	fmt.Println("Working Task Runing")
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Task Done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("Returned Result")
		retCh <- ret
		fmt.Println("Service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}
