package org

import "strconv"

// INN represents individual taxpayer number
type INN struct {
	// Template for legal entity - NNNN XXXXX C
	// Template for private person - NNNN XXXXXX CC

	// Region code (2 digits)
	RegionCode uint

	// Inspection code of tax authority that assigned INN (4 digits)
	Inspection uint

	// Serial number of entry about the person in
	// the territorial section of the Unified State Register of Taxpayers
	// of the tax authority that assigned the INN
	// (5 digits - for legal entities)
	// (6 digits - for private persons)
	SerialNumber uint

	// Checksum number
	// (1 digits - for legal entities)
	// (2 digits - for private persons)
	// calculated according to a special algorithm established by the Federal Tax Service
	Checksum uint

	// Value of INN
	Value uint64
}

// String returns INN value as string
func (inn *INN) String() string {
	return strconv.FormatUint(inn.Value, 10)
}
