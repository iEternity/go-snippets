package sort

import (
	"reflect"
	"testing"
)

func TestQsort(t *testing.T){
	testData := [][]int{
		{},
		{0},
		{3,1},
		{1,2,3},
		{1,5,3,9,28,90},
	}
	right := [][]int{
		{},
		{0},
		{1,3},
		{1,2,3},
		{1,3,5,9,28,90},
	}

	for i, v := range testData{
		qsort(v)
		if !reflect.DeepEqual(v, right[i]){
			t.Fatal(v, right[i])
		}
	}

	if !reflect.DeepEqual(testData, right){
		t.Fatal(testData, right)
	}
}