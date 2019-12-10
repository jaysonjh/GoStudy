package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 4
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int:
		fmt.Println("Integer")
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	checkType(f)
}

/// 放射的具体用法

type Employ struct {
	EmployID string
	Name string `format:"normal"`
	Age int
}

func (e *Employ)
