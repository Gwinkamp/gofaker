package org

// OrgType represents type of organization
// 1 - legal entity
// 2 - individual entrepreneur
type OrgType uint

const (
	LegalEntity OrgType = iota + 1
	IndividualEntrepreneur
)

// String returns string representation of OrgType
func (t OrgType) String() string {
	switch t {
	case LegalEntity:
		return "legal entity"
	case IndividualEntrepreneur:
		return "individual entrepreneur"
	default:
		return "unknown"
	}
}
