package csum_test

import (
	"fmt"
	"testing"

	"github.com/Gwinkamp/gofaker/pkg/csum"
)

func TestCalcINNLE__Positive(t *testing.T) {
	testcases := []struct {
		inn  uint64
		csum uint64
	}{
		{inn: 963522201, csum: 0},
		{inn: 968572434, csum: 3},
		{inn: 963727193, csum: 6},
		{inn: 9602866361, csum: 1},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("inn:%d", tc.inn)
		t.Run(testname, func(t *testing.T) {
			result, err := csum.CalcINNLE(tc.inn)
			if err != nil {
				t.Fatalf("CalcINNLE returned error: %v", err)
			}
			if result != tc.csum {
				t.Fatalf("CalcINNLE returned %d, expected %d", result, tc.csum)
			}
		})
	}
}

func TestCalcINNLE__ValidateINN(t *testing.T) {
	result, err := csum.CalcINNLE(123)
	if err == nil {
		t.Fatal("CalcINNLE completed without error, but error is expected")
	}

	expectedErrMsg := "csum.CalcINNLE: value is too short"
	if err.Error() != expectedErrMsg {
		t.Errorf("CalcINNLE returned error with message '%s', expected '%s'", err.Error(), expectedErrMsg)
	}
	if result != 0 {
		t.Errorf("CalcINNLE returned %d, expected 0", result)
	}
}

func TestCalcINN__Positive(t *testing.T) {
	testcases := []struct {
		inn  uint64
		csum uint64
	}{
		{inn: 9650163434, csum: 14},
		{inn: 9667961824, csum: 39},
		{inn: 962089855560, csum: 60},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("inn:%d", tc.inn)
		t.Run(testname, func(t *testing.T) {
			result, err := csum.CalcINN(tc.inn)
			if err != nil {
				t.Fatalf("CalcINN returned error: %v", err)
			}
			if result != tc.csum {
				t.Fatalf("CalcINN returned %d, expected %d", result, tc.csum)
			}
		})
	}
}

func TestCalcINN__ValidateINN(t *testing.T) {
	result, err := csum.CalcINN(123)
	if err == nil {
		t.Fatal("CalcINN completed without error, but error is expected")
	}

	expectedErrMsg := "csum.CalcINN: value is too short"
	if err.Error() != expectedErrMsg {
		t.Errorf("CalcINN returned error with message '%s', expected '%s'", err.Error(), expectedErrMsg)
	}
	if result != 0 {
		t.Errorf("CalcINN returned %d, expected 0", result)
	}
}
