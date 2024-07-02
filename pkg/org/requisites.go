package org

import (
	"fmt"
)

const (
	formatINN        string = "%012d"
	formatINNLE      string = "%010d"
	formatOGRN       string = "%013d"
	foramtOGRNIP     string = "%015d"
	formatSNILS      string = "%011d"
	defaultFormatVal string = "undefined"
)

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

	// Type indicates type of organization
	Type OrgType

	// Value of INN
	Value uint64
}

// String returns INN value as string
func (inn *INN) String() string {
	var format string
	switch inn.Type {
	case LegalEntity:
		format = formatINNLE
	case IndividualEntrepreneur:
		format = formatINN
	default:
		return defaultFormatVal
	}
	return fmt.Sprintf(format, inn.Value)
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

	// Type indicates type of organization
	Type OrgType

	// Value of OGRN
	Value uint64
}

// String returns OGRN value as string
func (ogrn *OGRN) String() string {
	var format string
	switch ogrn.Type {
	case LegalEntity:
		format = formatOGRN
	case IndividualEntrepreneur:
		format = foramtOGRNIP
	default:
		return defaultFormatVal
	}
	return fmt.Sprintf(format, ogrn.Value)
}

// SNILS represents insurance number of individual personal account
type SNILS struct {
	SerialNumber uint
	Checksum     uint
	Value        uint64
}

// String returns SNILS value as string
func (snils *SNILS) String() string {
	return fmt.Sprintf(formatSNILS, snils.Value)
}

// Stringf returns SNILS value as formatted string (XXX-XXX-XXX XX)
func (snils *SNILS) Stringf() string {
	var (
		r1 uint = snils.SerialNumber / 1_000_000
		r2 uint = snils.SerialNumber / 1_000 % 1_000
		r3 uint = snils.SerialNumber % 1_000
	)
	return fmt.Sprintf("%03d-%03d-%03d %02d", r1, r2, r3, snils.Checksum)
}
