package scriptparser

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type ncpdpMessage struct {
	groups []*group
}

func NewNCPDPMessage(data []byte) (*ncpdpMessage, error) {
	msg := ncpdpMessage{}

	ok, errorMessage := msg.Validate(data)
	if !ok {
		return nil, fmt.Errorf("%s. %w", errorMessage, NCPDPFormatError)
	}

	err := msg.Parse(data)

	//msg.GetTransactionAsBytes()
	return &msg, err

}
func (msg *ncpdpMessage) Parse(data []byte) error {

	trxType, err := DetermineTransactionType(data)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

	groups := bytes.Split(data, []byte(GroupSeparator))

	for i, groupData := range groups {
		grp, err := newGroup(groupData, i, trxType)
		if err != nil {
			return err
		}
		msg.groups = append(msg.groups, grp)
	}
	return nil
}

func (msg *ncpdpMessage) String() (string, error) {
	result := ""

	for _, grp := range msg.groups {

		grpStr, err := grp.String()
		if err != nil {
			return "", err
		}

		result += grpStr

	}

	return result, nil
}

func (msg *ncpdpMessage) ValidateRequest(data []byte) (bool, string) {
	okMessage := true
	errorMessage := ""
	if data[56] != SegmentSeparatorByte || data[57] != FieldSeparatorByte {
		errorMessage = "Bad structure, first segment is expected at char 57, but it is not"
		okMessage = false
	}
	bin := data[0:6]

	if !isNumber(bin) {
		okMessage = false
		errorMessage = "bin number must be numeric"
	}

	transactionCount := data[20:21]
	if !isNumber(transactionCount) {
		okMessage = false
		errorMessage = "transaction count must be numeric"
	}

	dos := data[38:46]
	if !isDate(dos) {
		okMessage = false
		errorMessage = "Date of Service must be a date"
	}
	return okMessage, errorMessage
}

func (msg *ncpdpMessage) ValidateResponse(data []byte) (bool, string) {
	okMessage := true
	errorMessage := ""
	if data[31] != SegmentSeparatorByte || data[32] != FieldSeparatorByte {
		errorMessage = "Bad structure, first segment is expected at char 57, but it is not"
		okMessage = false
	}

	transactionCount := data[4:5]
	if !isNumber(transactionCount) {
		okMessage = false
		errorMessage = "transaction count must be numeric"
	}

	dos := data[23:31]
	if !isDate(dos) {
		okMessage = false
		errorMessage = "Date of Service must be a date"
	}
	return okMessage, errorMessage
}

func (msg *ncpdpMessage) Validate(data []byte) (bool, string) {

	transactionType, err := DetermineTransactionType(data)
	if err != nil {
		return false, err.Error()
	}
	if transactionType == B1RequestType || transactionType == B2RequestType {
		return msg.ValidateRequest(data)
	} else {
		return msg.ValidateResponse(data)
	}
}

func (msg *ncpdpMessage) isGroupPresent(group int) bool {
	if len(msg.groups)-1 < group {
		return false
	} else {
		return true
	}

}
func (msg *ncpdpMessage) isSegmentPresent(group int, segmentId string) bool {

	if msg.isGroupPresent(group) {
		return msg.groups[group].IsSegmentPresent(segmentId)
	}

	return false
}

func (msg *ncpdpMessage) isSegmentPresentByTag(identifier string) bool {
	ids := strings.Split(identifier, ":")
	if len(ids) < 2 {
		log.Panic(errors.New("SegmentsIdentifiers must have groupnumber and segment separated by \":\" exampled \"0:HEADER\" or \"1:04\""))
	}
	groupId, err := strconv.Atoi(ids[0])
	if err != nil {
		log.Panic(err)
	}
	return msg.isSegmentPresent(groupId, ids[1])
}

func isNumber(v []byte) bool {
	_, err := strconv.Atoi(string(v))
	if err != nil {
		return false
	}
	return true
}
func isDate(v []byte) bool {
	_, err := time.Parse("20060102", string(v))
	if err != nil {
		return false
	}
	return true
}
