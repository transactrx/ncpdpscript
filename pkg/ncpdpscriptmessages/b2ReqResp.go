package ncpdpscriptmessages

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"
)

type B2Request struct {
	Header    *segments.B2RequestHeader    `ncpdp:"0:header" json:"header"`
	Insurance *segments.B2RequestInsurance `ncpdp:"0:04" json:"insurance"`
	Claims    []*B2RequestClaim            `ncpdp:"" json:"claims"`
}

type B2Response struct {
	Header   *segments.B2ResponseHeader  `ncpdp:"0:header" json:"header"`
	Message  *segments.B2ResponseMessage `ncpdp:"0:20" json:"message"`
	Statuses []*B2ResponseStatus         `ncpdp:"" json:"statuses"`
}

type B2ResponseStatus struct {
	ResponseStatus  *segments.B2ResponseStatus  `ncpdp:":21" json:"responseStatus"`
	ResponseClaim   *segments.B2ResponseClaim   `ncpdp:":22" json:"responseClaim"`
	ResponsePricing *segments.B2ResponsePricing `ncpdp:":23" json:"responsePricing"`
}

type B2RequestClaim struct {
	Claim   *segments.B2RequestClaim   `ncpdp:":07" json:"claim"`
	DUR     *segments.B2RequestDUR     `ncpdp:":08" json:"DUR"`
	Pricing *segments.B2RequestPricing `ncpdp:":11" json:"pricing"`
	COB     *segments.B2RequestCOB     `ncpdp:":05" json:"COB"`
}

func (b2Response *B2Response) GetResponseStatus(rxNumber string) (*B2ResponseStatus, error) {

	for _, status := range b2Response.Statuses {
		if status.ResponseClaim != nil && status.ResponseClaim.PrescriptionServiceReferenceNumber != nil && *status.ResponseClaim.PrescriptionServiceReferenceNumber == rxNumber {
			return status, nil
		}
	}
	return nil, fmt.Errorf("no status found for rx %s", rxNumber)
}

func (b2Response *B2Response) GetFormattedFullMessage(rxNumber string) (string, error) {
	b2RespStatus, err := b2Response.GetResponseStatus(rxNumber)
	if err != nil {

	}
	additionalMessage := b2RespStatus.ResponseStatus.GetAdditionalMessageInformation()

	fullMsg := fmt.Sprintf("%s", additionalMessage)
	return fullMsg, nil
}

func (b2Response *B2Response) IsTransactionAccepted(rxNumber string) (bool, error) {

	if *b2Response.Header.TransactionResponseStatus != "A" {
		return false, nil
	}
	b1RespStatus, err := b2Response.GetResponseStatus(rxNumber)
	if err != nil {
		return false, err
	}

	if b1RespStatus.ResponseStatus.TransactionResponseStatus == nil {
		return false, fmt.Errorf("no transaction response status found for rx %s", rxNumber)
	}

	switch *b1RespStatus.ResponseStatus.TransactionResponseStatus {
	case "A":
		return true, err
	default:
		return false, nil
	}
}

func (b2Response *B2Response) IsTransactionStatusText(rxNumber string) (string, error) {

	b1RespStatus, err := b2Response.GetResponseStatus(rxNumber)
	if err != nil {
		return "", err
	}

	if b1RespStatus.ResponseStatus.TransactionResponseStatus == nil {
		return "", fmt.Errorf("no transaction response status found for rx %s", rxNumber)
	}

	switch *b1RespStatus.ResponseStatus.TransactionResponseStatus {
	case "A":
		return "A-ACCEPTED", err
	case "R":
		return "R-REJECTED", err
	default:
		return "", nil
	}
}
func (b2Response *B2Request) GetJsonWithIds() ([]byte, error) {
	var json1 = jsoniter.Config{
		TagKey: "jsonWithId",
	}.Froze()
	jsonWithName, err := json1.Marshal(b2Response)
	return jsonWithName, err
}
