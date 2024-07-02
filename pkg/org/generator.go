package org

import (
	"math/rand/v2"

	"github.com/Gwinkamp/gofaker/pkg/csum"
)

const (
	testRegionCode  uint = 96
	numberOfRegions uint = 89
)

// Faker for generating requisites of organizations
type Faker struct {
	rnd *rand.Rand
}

// NewFaker creates a new OrgFaker
func NewFaker(src rand.Source) *Faker {
	return &Faker{rnd: rand.New(src)}
}

// INNLE generates INN of legal entity
func (faker *Faker) INNLE() INN {
	inn := INN{Type: LegalEntity}

	inn.RegionCode = testRegionCode
	inn.Inspection = inn.RegionCode*100 + faker.rnd.UintN(90) + 10
	inn.SerialNumber = faker.rnd.UintN(100000)

	number := (uint64(inn.Inspection) * 100000) + uint64(inn.SerialNumber)

	c, _ := csum.CalcINNLE(number)

	inn.Checksum = uint(c)
	inn.Value = (number * 10) + c

	return inn
}

// INN generates INN of private person
func (faker *Faker) INN() INN {
	inn := INN{Type: IndividualEntrepreneur}

	inn.RegionCode = testRegionCode
	inn.Inspection = inn.RegionCode*100 + faker.rnd.UintN(90) + 10
	inn.SerialNumber = faker.rnd.UintN(1_000_000)

	number := (uint64(inn.Inspection) * 1_000_000) + uint64(inn.SerialNumber)

	c, _ := csum.CalcINN(number)

	inn.Checksum = uint(c)
	inn.Value = (number * 100) + c

	return inn
}

// OGRN generates OGRN of legal entity
func (faker *Faker) OGRN() OGRN {
	ogrn := OGRN{Type: LegalEntity}

	if faker.rnd.IntN(2) == 0 {
		ogrn.Sign = 1
	} else {
		ogrn.Sign = 5
	}

	ogrn.YearEnd = faker.rnd.UintN(100)
	ogrn.RegionCode = faker.rnd.UintN(numberOfRegions) + 1
	ogrn.EntryNumber = faker.rnd.UintN(10_000_000)

	number := uint64(ogrn.Sign)*100_000_000_000 +
		uint64(ogrn.YearEnd)*1_000_000_000 +
		uint64(ogrn.RegionCode)*10_000_000 +
		uint64(ogrn.EntryNumber)

	c, _ := csum.CalcOGRN(number)

	ogrn.Checksum = uint(c)
	ogrn.Value = (number * 10) + c

	return ogrn
}

// OGRNIP generates OGRN of individual entrepreneur
func (faker *Faker) OGRNIP() OGRN {
	ogrn := OGRN{Type: IndividualEntrepreneur}

	ogrn.Sign = 3
	ogrn.YearEnd = faker.rnd.UintN(100)
	ogrn.RegionCode = faker.rnd.UintN(numberOfRegions) + 1
	ogrn.EntryNumber = faker.rnd.UintN(1_000_000_000)

	number := uint64(ogrn.Sign)*10_000_000_000_000 +
		uint64(ogrn.YearEnd)*100_000_000_000 +
		uint64(ogrn.RegionCode)*1_000_000_000 +
		uint64(ogrn.EntryNumber)

	c, _ := csum.CalcOGRNIP(number)

	ogrn.Checksum = uint(c)
	ogrn.Value = (number * 10) + c

	return ogrn
}
