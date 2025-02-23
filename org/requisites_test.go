package org_test

import (
	"fmt"
	"testing"

	"github.com/Gwinkamp/gofaker/org"
)

func TestStringINNLE(t *testing.T) {
	testcases := []struct {
		inn      org.INN
		expected string
	}{
		{
			inn: org.INN{
				RegionCode:   96,
				Inspection:   9624,
				SerialNumber: 74481,
				Checksum:     8,
				Type:         org.LegalEntity,
				Value:        9624744818,
			},
			expected: "9624744818",
		},
		{
			inn: org.INN{
				RegionCode:   1,
				Inspection:   100,
				SerialNumber: 692,
				Checksum:     0,
				Type:         org.LegalEntity,
				Value:        100006920,
			},
			expected: "0100006920",
		},
		{
			inn: org.INN{
				RegionCode:   96,
				Inspection:   9624,
				SerialNumber: 744816,
				Checksum:     06,
				Type:         org.IndividualEntrepreneur,
				Value:        962474481606,
			},
			expected: "962474481606",
		},
		{
			inn: org.INN{
				RegionCode:   2,
				Inspection:   201,
				SerialNumber: 7218,
				Checksum:     10,
				Type:         org.IndividualEntrepreneur,
				Value:        20100721810,
			},
			expected: "020100721810",
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("inn:%s", tc.expected)
		t.Run(testname, func(t *testing.T) {
			result := tc.inn.String()
			if result != tc.expected {
				t.Fatalf("INN.String returned %s, expected %s", result, tc.expected)
			}
		})
	}
}

func TestStringOGRN(t *testing.T) {
	testcases := []struct {
		ogrn     org.OGRN
		expected string
	}{
		{
			ogrn: org.OGRN{
				Sign:        1,
				YearEnd:     74,
				RegionCode:  65,
				EntryNumber: 8781849,
				Checksum:    2,
				Type:        org.LegalEntity,
				Value:       1746587818492,
			},
			expected: "1746587818492",
		},
		{
			ogrn: org.OGRN{
				Sign:        0,
				YearEnd:     74,
				RegionCode:  65,
				EntryNumber: 8781849,
				Checksum:    2,
				Type:        org.LegalEntity,
				Value:       746587818492,
			},
			expected: "0746587818492",
		},
		{
			ogrn: org.OGRN{
				Sign:        3,
				YearEnd:     16,
				RegionCode:  67,
				EntryNumber: 721135191,
				Checksum:    1,
				Type:        org.IndividualEntrepreneur,
				Value:       316677211351911,
			},
			expected: "316677211351911",
		},
		{
			ogrn: org.OGRN{
				Sign:        0,
				YearEnd:     16,
				RegionCode:  67,
				EntryNumber: 721135191,
				Checksum:    1,
				Type:        org.IndividualEntrepreneur,
				Value:       16677211351911,
			},
			expected: "016677211351911",
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("ogrn:%s", tc.expected)
		t.Run(testname, func(t *testing.T) {
			result := tc.ogrn.String()
			if result != tc.expected {
				t.Fatalf("OGRN.String returned %s, expected %s", result, tc.expected)
			}
		})
	}
}

func TestStringSNILS(t *testing.T) {
	testcases := []struct {
		snils     org.SNILS
		snilsStr  string
		snilsStrf string
	}{
		{
			snils: org.SNILS{
				SerialNumber: 135273663,
				Checksum:     59,
				Value:        13527366359,
			},
			snilsStr:  "13527366359",
			snilsStrf: "135-273-663 59",
		},
		{
			snils: org.SNILS{
				SerialNumber: 57033028,
				Checksum:     27,
				Value:        5703302827,
			},
			snilsStr:  "05703302827",
			snilsStrf: "057-033-028 27",
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("snils:%s", tc.snilsStrf)
		t.Run(testname, func(t *testing.T) {
			snilsStr := tc.snils.String()
			snilsStrf := tc.snils.Stringf()
			if snilsStr != tc.snilsStr {
				t.Fatalf("OGRN.String returned %s, expected %s", snilsStr, tc.snilsStr)
			}
			if snilsStrf != tc.snilsStrf {
				t.Fatalf("OGRN.Stringf returned %s, expected %s", snilsStrf, tc.snilsStrf)
			}
		})
	}
}
