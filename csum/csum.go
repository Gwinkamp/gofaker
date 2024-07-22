package csum

import (
	"errors"
	"fmt"
)

const (
	minBitDepthINNLE = 8
	maxBitDepthINNLE = 9
	minBitDepthINN   = 9
	maxBitDepthINN   = 10
	bitDepthOGRN     = 12
	bitDepthOGRNIP   = 14
	maxBitDepthSNILS = 9

	delimiterOGRN   = bitDepthOGRN - 1
	delimiterOGRNIP = bitDepthOGRNIP - 1
)

type Number interface {
	int | int32 | int64 | uint | uint32 | uint64
}

// Coefficients
var (
	cfcINNLE = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2}
	cfcINN1  = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2, 7}
	cfcINN2  = []uint64{8, 6, 4, 9, 5, 3, 10, 4, 2, 7, 3}
	cfcSNILS = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

// CalcINNLE calculates checksum of INN legal entity
func CalcINNLE[V Number](inn V) (uint64, error) {
	const operation = "csum.CalcINNLE"

	val, err := prepareForCalc(inn, minBitDepthINNLE, maxBitDepthINNLE)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}
	return calcINNCSum(val, cfcINNLE), nil
}

// CalcINN calculates checksum of INN private person
func CalcINN[V Number](inn V) (uint64, error) {
	const operation = "csum.CalcINN"

	val, err := prepareForCalc(inn, minBitDepthINN, maxBitDepthINN)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}

	csum1 := calcINNCSum(val, cfcINN1)
	csum1 %= 10 // if csum == 10 then csum must be 0

	val = val*10 + csum1
	csum2 := calcINNCSum(val, cfcINN2)
	csum2 %= 10

	return csum1*10 + csum2, nil
}

// CalcOGRN calculates checksum of OGRN legal entity
func CalcOGRN[V Number](ogrn V) (uint64, error) {
	const operation = "csum.CalcOGRN"

	val, err := prepareForCalc(ogrn, bitDepthOGRN, bitDepthOGRN)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}
	return (val % delimiterOGRN) % 10, nil
}

// CalcOGRNIP calculates checksum of OGRN individual entrepreneur
func CalcOGRNIP[V Number](ogrnip V) (uint64, error) {
	const operation = "csum.CalcOGRNIP"

	val, err := prepareForCalc(ogrnip, bitDepthOGRNIP, bitDepthOGRNIP)
	if err != nil {
		return 0, fmt.Errorf("%s: %v", operation, err)
	}
	return (val % delimiterOGRNIP) % 10, nil
}

// CalcSNILS calculates checksum of SNILS
func CalcSNILS[V Number](snils V) uint64 {
	val := uint64(snils)
	bd := getBitDepth(val)
	if bd > maxBitDepthSNILS {
    val /= pow10(bd - maxBitDepthSNILS)
  }

	csum := calcWithCfc(val, cfcSNILS)
	switch {
	case csum < 100:
		return csum
	case csum == 100 || csum == 101:
		return 0
	default:
		return csum % 101
	}
}

// prepareForCalc prepares number for calculating checksum
func prepareForCalc[V Number](num V, minBitDepth, maxBitDepth int) (uint64, error) {
	val := uint64(num)
	bd := getBitDepth(val)

	if bd < minBitDepth {
		return 0, errors.New("value is too short")
	}
	if bd > maxBitDepth {
		val /= pow10(bd - maxBitDepth)
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
	return (calcWithCfc(inn, cfc) % 11) % 10
}

// calcWithCfc calculates checksum of number with coefficients
func calcWithCfc(val uint64, cfc []uint64) uint64 {
	var (
		cnt int    = 0
		res uint64 = 0
	)
	for val != 0 {
		res += val % 10 * cfc[cnt]
		val /= 10
		cnt++
	}
	return res
}
