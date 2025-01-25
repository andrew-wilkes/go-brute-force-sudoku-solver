package main

import "testing"

func TestChecks(t *testing.T) {
	testGrid := make([][]int, SIZE)
	for n := 0; n < SIZE; n++ {
		testGrid[n] = make([]int, SIZE)
	}

	want := true
	got := rowOK(1, 0, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

	testGrid[0][3] = 1
	want = false
	got = rowOK(1, 0, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

	want = true
	got = colOK(1, 0, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

	want = false
	got = colOK(1, 3, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

	testGrid[3][3] = 8
	want = true
	got = boxOk(8, 1, 1, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

	want = false
	got = boxOk(8, 3, 3, testGrid)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}
