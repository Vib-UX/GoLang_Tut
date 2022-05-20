package math

import (
	"testing"
)

type AddData struct {
	x, y, result int
}

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Test for 1+3 failed. Expected %d, got %d\n", 4, result)
	} else {
		t.Logf("Test passed. Expected %d, got %d\n", 4, result)
	}

	testData := []AddData{
		{1, 2, 3},
		{2, 4, 6},
		{-1, 3, 2},
	}

	for _, dataNum := range testData {
		result := Add(dataNum.x, dataNum.y)
		if result != dataNum.result {
			t.Errorf("Test for %d, + %d failed. Expected %d, got %d\n", dataNum.x, dataNum.y, dataNum.result, result)
		} else {
			t.Logf("Test passed. Expected %d, got %d\n", dataNum.result, result)
		}
	}
}
