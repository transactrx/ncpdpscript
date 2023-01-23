package scriptparser

import (
	"bytes"
	"fmt"
)

type segment struct {
	id        string
	fields    []*field
	fieldsMap map[string]*field
	header    bool
}

func DetermineTransactionType(data []byte) (string, error) {
	seg, err := newSegmentHeader(data)
	if err != nil {
		return "", err
	}
	fld := seg.GetFieldById("transactionCode")
	if len(fld.fieldData) < 1 || fld.fieldData != "B1" {
		return "", fmt.Errorf("NCPDP transaction type %s not supported", fld.fieldData)
	}

	return fld.fieldData, nil
}
func newSegment(data []byte) (*segment, error) {
	seg := segment{}
	seg.fieldsMap = make(map[string]*field)
	fieldsBytes := bytes.Split(data, []byte(FieldSeparator))

	for i, fieldsByte := range fieldsBytes {
		if i == 0 {
			//empty field
			continue
		}
		newField, err := newField(fieldsByte)
		if err != nil {
			return nil, err
		}
		seg.addField(newField)

	}
	if len(seg.fields) == 0 {
		return nil, fmt.Errorf("invalid empty segment, %w", NCPDPFormatError)
	}
	seg.id = seg.fields[0].fieldData

	return &seg, nil
}
func newSegmentHeader(data []byte) (*segment, error) {
	seg := segment{header: true}
	seg.id = "header"
	seg.fieldsMap = make(map[string]*field)
	bin := data[0:6]
	versionRel := data[6:8]
	transactionCode := data[8:10]
	pcn := data[10:20]
	transactionCount := data[20:21]
	serviceProviderIdQualifier := data[21:23]
	serviceProviderId := data[23:38]
	dos := data[38:46]
	softwareCertId := data[46:]

	field1, err := newHeaderField(bin, "bin", 6)
	if err != nil {
		return nil, err
	}
	seg.addField(field1)
	field2, err := newHeaderField(versionRel, "versionRel", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field2)
	field3, err := newHeaderField(transactionCode, "transactionCode", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field3)
	field4, err := newHeaderField(pcn, "pcn", 10)
	if err != nil {
		return nil, err
	}
	seg.addField(field4)
	field5, err := newHeaderField(transactionCount, "transactionCount", 1)
	if err != nil {
		return nil, err
	}
	seg.addField(field5)
	field6, err := newHeaderField(serviceProviderIdQualifier, "serviceProviderIdQualifier", 2)
	if err != nil {
		return nil, err
	}
	seg.addField(field6)
	field7, err := newHeaderField(serviceProviderId, "serviceProviderId", 15)
	if err != nil {
		return nil, err
	}
	seg.addField(field7)
	field8, err := newHeaderField(dos, "dos", 8)
	if err != nil {
		return nil, err
	}
	seg.addField(field8)
	field9, err := newHeaderField(softwareCertId, "softwareCertId", 10)
	if err != nil {
		return nil, err
	}
	seg.addField(field9)

	return &seg, nil
}

func (seg *segment) String() (string, error) {
	result := ""
	if !seg.header {
		result = SegmentSeparator
	}

	for i := 0; i < len(seg.fields); i++ {
		fieldValue, err := seg.fields[i].String()
		if err != nil {
			return "", err
		}

		result += fieldValue

	}
	return result, nil
}

func (seg *segment) GetFieldById(fieldId string) *field {
	return seg.fieldsMap[fieldId]
}
func (seg *segment) GetFieldArrayById(fieldId string) []*field {

	var result []*field = make([]*field, 0)
	i := 0

	for true {

		iStr := fmt.Sprintf(":%d", i)
		if i == 0 {
			iStr = ""
		}
		if seg.fieldsMap[fieldId+iStr] != nil {
			result = append(result, seg.fieldsMap[fieldId+iStr])

		} else {
			break
		}
		i++
	}
	return result

}

func (seg *segment) addField(f *field) {
	seg.fields = append(seg.fields, f)
	if seg.fieldsMap[f.fieldId] != nil {

		i := 1
		for true {
			istr := fmt.Sprintf("%d", i)
			if seg.fieldsMap[f.fieldId+":"+istr] == nil {
				seg.fieldsMap[f.fieldId+":"+istr] = f
				break
			}
			i++
		}

	} else {
		seg.fieldsMap[f.fieldId] = f
	}

}
