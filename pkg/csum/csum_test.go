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

func TestCalcOGRN__Positive(t *testing.T) {
	testcases := []struct {
		ogrn uint64
		csum uint64
	}{
		{ogrn: 139230112701, csum: 9},
		{ogrn: 151070729700, csum: 0},
		{ogrn: 1685537354292, csum: 2},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("ogrn:%d", tc.ogrn)
		t.Run(testname, func(t *testing.T) {
			result, err := csum.CalcOGRN(tc.ogrn)
			if err != nil {
				t.Fatalf("CalcOGRN returned error: %v", err)
			}
			if result != tc.csum {
				t.Fatalf("CalcOGRN returned %d, expected %d", result, tc.csum)
			}
		})
	}
}

func TestCalcOGRN__ValidateOGRN(t *testing.T) {
	result, err := csum.CalcOGRN(123)
	if err == nil {
		t.Fatal("CalcOGRN completed without error, but error is expected")
	}

	expectedErrMsg := "csum.CalcOGRN: value is too short"
	if err.Error() != expectedErrMsg {
		t.Errorf("CalcOGRN returned error with message '%s', expected '%s'", err.Error(), expectedErrMsg)
	}
	if result != 0 {
		t.Errorf("CalcOGRN returned %d, expected 0", result)
	}
}

func TestCalcOGRNIP__Positive(t *testing.T) {
	testcases := []struct {
		ogrn uint64
		csum uint64
	}{
		{ogrn: 44702273543581, csum: 4},
		{ogrn: 31888609105041, csum: 0},
		{ogrn: 370392935189995, csum: 5},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("ogrn:%d", tc.ogrn)
		t.Run(testname, func(t *testing.T) {
			result, err := csum.CalcOGRNIP(tc.ogrn)
			if err != nil {
				t.Fatalf("CalcOGRNIP returned error: %v", err)
			}
			if result != tc.csum {
				t.Fatalf("CalcOGRNIP returned %d, expected %d", result, tc.csum)
			}
		})
	}
}

func TestCalcOGRN__ValidateOGRNIP(t *testing.T) {
	result, err := csum.CalcOGRN(123)
	if err == nil {
		t.Fatal("CalcOGRNIP completed without error, but error is expected")
	}

	expectedErrMsg := "csum.CalcOGRN: value is too short"
	if err.Error() != expectedErrMsg {
		t.Errorf("CalcOGRNIP returned error with message '%s', expected '%s'", err.Error(), expectedErrMsg)
	}
	if result != 0 {
		t.Errorf("CalcOGRNIP returned %d, expected 0", result)
	}
}
