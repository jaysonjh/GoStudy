package pipefilter

import "testing"

func TestStraighPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	intFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	sp := NewStraighPipeline("p1", splitFilter, intFilter, sumFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}

	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}

func TestSplitFilter(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	sp := NewStraighPipeline("p1", splitFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
