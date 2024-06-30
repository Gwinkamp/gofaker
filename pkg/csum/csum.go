package csum

import (
	"errors"
	"fmt"
)

const (
	// minBitDepthInnLE is minimum bit depth of inn legal entity for calculating checksum
	minBitDepthInnLE = 9
	// minBitDepthInn is minimum bit depth of inn private person for calculating checksum
	minBitDepthInn = 10
	// minBitDepthOGRN is minimum bit depth of ogrn legal entity for calculating checksum
	minBitDepthOGRN = 12
	// minBitDepthOGRNIP is minimum bit depth of ogrn individual entrepreneur for calculating checksum
	minBitDepthOGRNIP = 14

	delimiterOGRN = minBitDepthOGRN - 1
	delimiterOGRNIP = minBitDepthOGRNIP - 1
)

type Number interface {
	int | int32 | int64 | uint | uint32 | uint64
}

// Coefficients
var (
	cfcInnLE = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2}
	cfcInn1  = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2, 7}
	cfcInn2  = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2, 7, 3}
)

// CalcINNLE calculates checksum of INN legal entity
func CalcINNLE[V Number](inn V) (uint64, error) {
	const operation = "csum.CalcINNLE"

	val, err := prepareForCalc(inn, minBitDepthInnLE)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}
	return calcINNCSum(val, cfcInnLE), nil
}

// CalcINN calculates checksum of INN private person
func CalcINN[V Number](inn V) (uint64, error) {
	const operation = "csum.CalcINN"

	val, err := prepareForCalc(inn, minBitDepthInn)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}

	csum1 := calcINNCSum(val, cfcInn1)
	csum1 %= 10 // if csum == 10 then csum must be 0

	val = val*10 + csum1
	csum2 := calcINNCSum(val, cfcInn2)
	csum2 %= 10

	return csum1*10 + csum2, nil
}

// CalcOGRN calculates checksum of OGRN legal entity
func CalcOGRN[V Number](ogrn V) (uint64, error) {
	const operation = "csum.CalcOGRN"

	val, err := prepareForCalc(ogrn, minBitDepthOGRN)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}
	return (val % delimiterOGRN) % 10, nil
}

// CalcOGRNIP calculates checksum of OGRN individual entrepreneur
func CalcOGRNIP[V Number](ogrnip V) (uint64, error) {
	const operation = "csum.CalcOGRNIP"

  val, err := prepareForCalc(ogrnip, minBitDepthOGRNIP)
  if err!= nil {
    return 0, fmt.Errorf("%s: %v", operation, err)
  }
  return (val % delimiterOGRNIP) % 10, nil
}

// prepareForCalc prepares number for calculating checksum
func prepareForCalc[V Number](num V, minBitDepth int) (uint64, error) {
	val := uint64(num)
	bd := getBitDepth(val)

	if bd < minBitDepth {
		return 0, errors.New("value is too short")
	}
	if bd > minBitDepth {
		val /= pow10(bd - minBitDepth)
	}
	return val, nil
}

// getBitDepth calculates bit depth of a number
func getBitDepth(num uint64) int {
	if num == 0 {
		return 1
	}
	cnt := 0
	for num != 0 {
		num /= 10
		cnt++
	}
	return cnt
}

// pow10 raising 10 to power of num
func pow10[V Number](num V) uint64 {
	if num == 0 {
		return 1
	}

	var res uint64 = 10
	for i := num - 1; i != 0; i-- {
		res *= 10
	}
	return res
}

// calcINNCSum calculates checksum of INN
func calcINNCSum(inn uint64, cfc []uint64) uint64 {
	var (
		cnt int    = 0
		res uint64 = 0
	)
	for inn != 0 {
		res += inn % 10 * cfc[cnt]
		inn /= 10
		cnt++
	}
	return (res % 11) % 10
}
