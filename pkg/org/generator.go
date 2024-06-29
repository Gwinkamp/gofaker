package org

import (
	"math/rand/v2"

	"github.com/Gwinkamp/gofaker/pkg/csum"
)

const (
	testRegionCode uint = 96
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
	inn := INN{}

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
	inn := INN{}

  inn.RegionCode = testRegionCode
  inn.Inspection = inn.RegionCode*100 + faker.rnd.UintN(90) + 10
  inn.SerialNumber = faker.rnd.UintN(1000000)

  number := (uint64(inn.Inspection) * 1000000) + uint64(inn.SerialNumber)

  c, _ := csum.CalcINN(number)

  inn.Checksum = uint(c)
  inn.Value = (number * 100) + c

  return inn
}
