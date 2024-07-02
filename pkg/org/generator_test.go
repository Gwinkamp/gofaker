package org_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/Gwinkamp/gofaker/pkg/org"
)

var (
	seed = [32]byte{'a', 'e', '5', '1', '6', 'f', '6', '2', 'd', '2', '0', '4', '4', '7', 'a', 'c', '9', '2', '5', '0', '8', 'a', '0', 'c', '0', '5', '8', 'e', '7', 'e', '7', '8'}
)

func TestINNLE(t *testing.T) {
	cc8 := rand.NewChaCha8(seed)
	faker := org.NewFaker(cc8)

	testcases := []org.INN{
		{
			RegionCode:   96,
			Inspection:   9624,
			SerialNumber: 74481,
			Checksum:     8,
			Type:         org.LegalEntity,
			Value:        9624744818,
		},
		{
			RegionCode:   96,
			Inspection:   9674,
			SerialNumber: 87818,
			Checksum:     0,
			Type:         org.LegalEntity,
			Value:        9674878180,
		},
		{
			RegionCode:   96,
			Inspection:   9656,
			SerialNumber: 49778,
			Checksum:     4,
			Type:         org.LegalEntity,
			Value:        9656497784,
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("inn:%d", tc.Value)
		t.Run(testname, func(t *testing.T) {
			result := faker.INNLE()
			if result != tc {
				t.Fatalf("INNLE returned %d, expected %d", result, tc)
			}
		})
	}
}

func TestINN(t *testing.T) {
	cc8 := rand.NewChaCha8(seed)
	faker := org.NewFaker(cc8)

	testcases := []org.INN{
		{
			RegionCode:   96,
			Inspection:   9624,
			SerialNumber: 744816,
			Checksum:     06,
			Type:         org.IndividualEntrepreneur,
			Value:        962474481606,
		},
		{
			RegionCode:   96,
			Inspection:   9674,
			SerialNumber: 878184,
			Checksum:     61,
			Type:         org.IndividualEntrepreneur,
			Value:        967487818461,
		},
		{
			RegionCode:   96,
			Inspection:   9656,
			SerialNumber: 497780,
			Checksum:     03,
			Type:         org.IndividualEntrepreneur,
			Value:        965649778003,
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("inn:%d", tc.Value)
		t.Run(testname, func(t *testing.T) {
			result := faker.INN()
			if result != tc {
				t.Fatalf("INN returned %d, expected %d", result, tc)
			}
		})
	}
}

func TestOGRN(t *testing.T) {
	cc8 := rand.NewChaCha8(seed)
	faker := org.NewFaker(cc8)

	testcases := []org.OGRN{
		{
			Sign:        1,
			YearEnd:     74,
			RegionCode:  65,
			EntryNumber: 8781849,
			Checksum:    2,
			Type:        org.LegalEntity,
			Value:       1746587818492,
		},
		{
			Sign:        5,
			YearEnd:     49,
			RegionCode:  39,
			EntryNumber: 1357938,
			Checksum:    5,
			Type:        org.LegalEntity,
			Value:       5493913579385,
		},
		{
			Sign:        1,
			YearEnd:     22,
			RegionCode:  29,
			EntryNumber: 5575851,
			Checksum:    9,
			Type:        org.LegalEntity,
			Value:       1222955758519,
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("ogrn:%d", tc.Value)
		t.Run(testname, func(t *testing.T) {
			result := faker.OGRN()
			if result != tc {
				t.Fatalf("OGRN returned %d, expected %d", result, tc)
			}
		})
	}
}

func TestOGRNIP(t *testing.T) {
	cc8 := rand.NewChaCha8(seed)
	faker := org.NewFaker(cc8)

	testcases := []org.OGRN{
		{
			Sign:        3,
			YearEnd:     16,
			RegionCode:  67,
			EntryNumber: 721135191,
			Checksum:    1,
			Type:        org.IndividualEntrepreneur,
			Value:       316677211351911,
		},
		{
			Sign:        3,
			YearEnd:     87,
			RegionCode:  46,
			EntryNumber: 497780578,
			Checksum:    3,
			Type:        org.IndividualEntrepreneur,
			Value:       387464977805783,
		},
		{
			Sign:        3,
			YearEnd:     43,
			RegionCode:  13,
			EntryNumber: 729878391,
			Checksum:    2,
			Type:        org.IndividualEntrepreneur,
			Value:       343137298783912,
		},
	}

	for _, tc := range testcases {
		testname := fmt.Sprintf("ogrn:%d", tc.Value)
		t.Run(testname, func(t *testing.T) {
			result := faker.OGRNIP()
			if result != tc {
				t.Fatalf("OGRNIP returned %d, expected %d", result, tc)
			}
		})
	}
}
