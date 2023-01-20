package scriptparser

import (
	"fmt"
	"strconv"
)

type field struct {
	fieldId     string
	fieldData   string
	fieldLength *int
	header      bool
}

func newField(data []byte) (*field, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("invalid NCPDP field, it must be at least 2 characters long %w", NCPDPFormatError)
	}
	f := field{header: false}
	f.fieldId = string(data[:2])
	f.fieldData = string(data[2:])
	return &f, nil
}
func newHeaderField(data []byte, fieldId string, fieldLength int) (*field, error) {
	f := field{
		header:      true,
		fieldLength: &fieldLength,
		fieldId:     fieldId,
		fieldData:   string(data),
	}
	return &f, nil
}

//func (f *field) fixField() error {
//	if f.fieldLength != nil {
//		if len(f.fieldData) < *f.fieldLength {
//			f.fieldData = fmt.Sprintf("%-*s", *f.fieldLength, string(f.data))
//		}
//		if len(f.fieldData) > *f.fieldLength {
//			return fmt.Errorf("header field Id:%s is fixed length, but the data in it is too long. Max length is %d, current length is %d, %w", f.fieldId, f.fieldLength, len(f.data), NCPDPFormatError)
//		}
//	}
//	return nil
//}

func (f *field) AsBytes() ([]byte, error) {
	str, err := f.String()

	if err != nil {
		return nil, err
	}
	return []byte(str), err
}

func (f *field) String() (string, error) {

	if !f.header {
		return FieldSeparator + f.fieldId + f.fieldData, nil
	} else {
		return f.fieldData, nil
	}

}

func (f *field) int32() (*int32, error) {
	var result int32
	tempResult, err := strconv.ParseInt(f.fieldData, 10, 32)

	if err != nil {
		return nil, err
	}
	result = int32(tempResult)
	return &result, err
}
func (f *field) int64() (*int64, error) {

	result, err := strconv.ParseInt(f.fieldData, 10, 64)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (f *field) int() (*int, error) {

	result, err := strconv.Atoi(f.fieldData)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (f *field) Currency32() (*Currency32, error) {
	if f.fieldData == "" {
		return nil, nil
	}
	fl32, err := parseNCPDPCurrencyString(f.fieldData, 32)
	if err != nil {
		return nil, err
	}
	result := Currency32(*fl32)

	return &result, nil
}
func (f *field) Currency64() (*Currency64, error) {
	if f.fieldData == "" {
		return nil, nil
	}
	result, err := parseNCPDPCurrencyString(f.fieldData, 64)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (f *field) Float32() (*float32, error) {
	flotVal64, err := strconv.ParseFloat(f.fieldData, 32)
	flotVal32 := float32(flotVal64)
	if err != nil {
		return nil, err
	} else {
		var flVal float32 = flotVal32 / 100
		return &flVal, nil
	}
}
func (f *field) Float64() (*float64, error) {
	flotVal64, err := strconv.ParseFloat(f.fieldData, 64)
	if err != nil {
		return nil, err
	} else {

		return &flotVal64, nil
	}
}
