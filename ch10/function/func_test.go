package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

///函数返回前执行defer，类似于Java的finally
func TestDefer(t *testing.T) {
	/// 用于安全释放资源，释放锁
	defer Clear()

	fmt.Println("Start")
	panic("err")
}
