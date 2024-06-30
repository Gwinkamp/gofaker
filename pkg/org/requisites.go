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

// OGRN represents main state registration number
type OGRN struct {
	// Template for legal entity - S YY NN XXXXXXX C
	// Template for individual entrepreneur - S YY NN XXXXXXXXX C

	// Sign of attribution of state registration number (1 digit)
	// (is equal to 1 or 5 for legal entities)
	// (is equal to 3 for individual entrepreneurs)
	Sign uint

	// YearEnd is last two digits of year of entry in state register (2 digits)
	YearEnd uint

	// Region code (2 digits)
	RegionCode uint

	// EntryNumber in the Unified State Register of Legal Entities during the year
	// (7 digits - for legal entities)
	// (9 digits - for individual entrepreneurs)
	EntryNumber uint

	// Checksum number (1 digit)
	Checksum uint

	// Value of OGRN
	Value uint64
}

// String returns OGRN value as string
func (ogrn *OGRN) String() string {
	return strconv.FormatUint(ogrn.Value, 10)
}
