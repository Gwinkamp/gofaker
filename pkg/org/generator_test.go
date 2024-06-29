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
			Value:        9624744818,
		},
		{
			RegionCode:   96,
			Inspection:   9674,
			SerialNumber: 87818,
			Checksum:     0,
			Value:        9674878180,
		},
		{
			RegionCode:   96,
			Inspection:   9656,
			SerialNumber: 49778,
			Checksum:     4,
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
			Value:        962474481606,
		},
		{
			RegionCode:   96,
			Inspection:   9674,
			SerialNumber: 878184,
			Checksum:     61,
			Value:        967487818461,
		},
		{
			RegionCode:   96,
			Inspection:   9656,
			SerialNumber: 497780,
			Checksum:     03,
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
