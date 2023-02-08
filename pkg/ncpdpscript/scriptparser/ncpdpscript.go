// Package scriptparser contains the lower level code that can marshall into
// NCPDP script format and is able unmarshall from the NCPDP script format
package scriptparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type StringerWithErr interface {
	String() (string, error)
}
type ByterWithErr interface {
	AsBytes() ([]byte, error)
}

type Float interface {
	float64 | float32 | Currency32 | Currency64
}

const (
	B1RequestType = iota
	B1ResponseType
	B2RequestType
	B2ResponseType
)

var NCPDPFormatError = errors.New("NCPDP format error")
var NCPDPSegmentNotFound = errors.New("segment not found by given id")
var NCPDPFieldNotFound = errors.New("field not found in segment by given id")
var RequestStructInvalid = errors.New("the structure passed for marshalling or unmarshalling is invalid")
var NCPDPMessageInvalidOrUnsupported = errors.New("the NCPDP message is invalid or unsupported.  Currently support only exist for b1 and b2 request and response messages")

const SegmentSeparator = "\u001E"
const GroupSeparator = "\u001D"
const FieldSeparator = "\u001C"
const SegmentSeparatorByte = 30
const GroupSeparatorByte = 29
const FieldSeparatorByte = 28

type Currency32 float32
type Currency64 float64
type UnsignedNumeric float32 //9(7)v999 Format=9999999.999

// metric decimal units
func parseUnsignedNumericFromString(value string) (*UnsignedNumeric, error) {

	if value == "" {
		t := UnsignedNumeric(0.0)
		return &t, nil
	}

	tResult, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil, err
	}
	result := UnsignedNumeric(tResult / 1000)

	return &result, nil
}

func parseNCPDPCurrencyString(value string, bit int) (*Currency64, error) {

	if bit != 32 && bit != 64 {
		return nil, fmt.Errorf("only 32 or 64 bit are supported, %d is not supported", bit)
	}
	if value == "" {

		t := Currency64(0.0)

		return &t, nil
	}

	lastChar := string(value[len(value)-1])
	v := ""
	negative := false

	switch {

	case lastChar == "}" || lastChar == "{":
		negative = lastChar == "}"
		v = "0"
		break
	case lastChar == "A" || lastChar == "J":
		negative = lastChar == "J"
		v = "1"
		break
	case lastChar == "B" || lastChar == "K":
		negative = lastChar == "K"
		v = "2"
		break
	case lastChar == "C" || lastChar == "L":
		negative = lastChar == "L"
		v = "3"
		break
	case lastChar == "D" || lastChar == "M":
		negative = lastChar == "M"
		v = "4"
		break

	case lastChar == "N" || lastChar == "E":
		negative = lastChar == "N"
		v = "5"
		break

	case lastChar == "F" || lastChar == "O":
		negative = lastChar == "O"
		v = "6"
		break
	case lastChar == "G" || lastChar == "P":
		negative = lastChar == "P"
		v = "7"
		break
	case lastChar == "H" || lastChar == "Q":
		negative = lastChar == "Q"
		v = "8"
		break
	case lastChar == "I" || lastChar == "R":
		negative = lastChar == "R"
		v = "9"
		break
	}
	value = strings.Replace(value, lastChar, v, 1)

	tResult, err := strconv.ParseFloat(value, bit)
	if err != nil {
		return nil, err
	}
	result := Currency64(tResult / 100)
	if negative {
		result = Currency64(-tResult / 100)
	}

	return &result, nil
}
func renderNCPDPCurrencyString[T Float](value T) string {

	sVal := fmt.Sprintf("%.0f", value*100)

	lastCharIndex := len(sVal) - 1
	lastChar := string(sVal[lastCharIndex])
	replacementChar := ""

	switch {
	case lastChar == "0":
		if value >= 0 {
			replacementChar = "{"
		} else {
			replacementChar = "}"
		}
	case lastChar == "1":
		if value >= 0 {
			replacementChar = "A"
		} else {
			replacementChar = "J"
		}

	case lastChar == "2":
		if value >= 0 {
			replacementChar = "B"
		} else {
			replacementChar = "K"
		}
	case lastChar == "3":
		if value >= 0 {
			replacementChar = "C"
		} else {
			replacementChar = "L"
		}
	case lastChar == "4":
		if value >= 0 {
			replacementChar = "D"
		} else {
			replacementChar = "M"
		}
	case lastChar == "5":
		if value >= 0 {
			replacementChar = "E"
		} else {
			replacementChar = "N"
		}
	case lastChar == "6":
		if value >= 0 {
			replacementChar = "F"
		} else {
			replacementChar = "O"
		}
	case lastChar == "7":
		if value >= 0 {
			replacementChar = "G"
		} else {
			replacementChar = "P"
		}
	case lastChar == "8":
		if value >= 0 {
			replacementChar = "H"
		} else {
			replacementChar = "Q"
		}
	case lastChar == "9":
		if value >= 0 {
			replacementChar = "I"
		} else {
			replacementChar = "R"
		}
	}

	res := strings.TrimSuffix(sVal, lastChar) + replacementChar
	//remove negatives
	return strings.TrimPrefix(res, "-")

}

func DetermineTransactionType(data []byte) (int, error) {

	//check if response object
	headerInfo := string(data[0:4])
	if headerInfo == "D0B1" || headerInfo == "D0S1" {
		return B1ResponseType, nil
	}
	if headerInfo == "D0B2" || headerInfo == "D0S2" {
		return B2ResponseType, nil
	}
	headerInfo = string(data[6:10])
	if headerInfo == "D0B1" || headerInfo == "D0S1" {
		return B1RequestType, nil
	}
	if headerInfo == "D0B2" || headerInfo == "D0S2" {
		return B2RequestType, nil
	}
	fmt.Printf("Unable to parse transactions. NCPDP message is invalid or unsupported -> %s", string(data))
	return 0, NCPDPMessageInvalidOrUnsupported
}
