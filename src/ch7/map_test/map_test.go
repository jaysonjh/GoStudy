package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Logf("len m1 %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3 %d", len(m3))
}

func TestMapValue(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 is %d", v)
	} else {
		t.Log("Key 3 is not existing")
	}
}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
