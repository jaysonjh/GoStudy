package easyjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info": {
		"name":"Jayson",
		"age":35
	},
	"job_info": {
		"skills": ["Java","Go","C"]
	}
}
`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		// fmt.Println(*e)
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		e.UnmarshalJSON([]byte(jsonStr))
		// fmt.Println(e)
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
