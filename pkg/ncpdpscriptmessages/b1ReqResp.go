package ncpdpscriptmessages

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"
)

type B1Request struct {
	Header    *segments.B1RequestHeader    `ncpdp:"0:header" json:"header" jsonWithId:"header"`
	Insurance *segments.B1RequestInsurance `ncpdp:"0:04" json:"insurance" jsonWithId:"AM04_insurance"`
	Patient   *segments.B1RequestPatient   `ncpdp:"0:01" json:"patient" jsonWithId:"AM01_patient"`
	Claims    []*B1RequestClaim            `ncpdp:"" json:"claims" jsonWithId:"claims"`
}

type B1Response struct {
	Header    *segments.B1ResponseHeader    `ncpdp:"0:header" json:"header"`
	Message   *segments.B1ResponseMessage   `ncpdp:"0:20" json:"message"`
	Insurance *segments.B1ResponseInsurance `ncpdp:"0:25" json:"insurance"`
	Patient   *segments.B1ResponsePatient   `ncpdp:"0:29" json:"patient"`
	Statuses  []*B1ResponseStatus           `ncpdp:"" json:"statuses"`
}

type B1RequestClaim struct {
	Claim                   *segments.B1RequestClaim                   `ncpdp:":07" json:"claim" jsonWithId:"AM07_claim"`
	Pricing                 *segments.B1RequestPricing                 `ncpdp:":11" json:"pricing" jsonWithId:"AM11_pricing"`
	PharmacyProvider        *segments.B1RequestPharmacyProvider        `ncpdp:":02" json:"pharmacyProvider" jsonWithId:"AM02_pharmacyProvider"`
	Prescriber              *segments.B1RequestPrescriber              `ncpdp:":03" json:"prescriber" jsonWithId:"AM03_prescriber"`
	COB                     *segments.B1RequestCOB                     `ncpdp:":05" json:"COB" jsonWithId:"AM05_COB"`
	WorkersComp             *segments.B1RequestWorkersCompensation     `ncpdp:":06" json:"workersComp" jsonWithId:"AM06_workersComp"`
	DUR                     *segments.B1RequestDUR                     `ncpdp:":08" json:"DUR" jsonWithId:"AM08_DUR"`
	Coupon                  *segments.B1RequestCoupon                  `ncpdp:":09" json:"coupon" jsonWithId:"AM09_coupon"`
	Compound                *segments.B1RequestCompound                `ncpdp:":10" json:"compound" jsonWithId:"AM10_compound"`
	Clinical                *segments.B1RequestClinical                `ncpdp:":13" json:"clinical" jsonWithId:"AM13_clinical"`
	AdditionalDocumentation *segments.B1RequestAdditionalDocumentation `ncpdp:":14" json:"additionalDocumentation" jsonWithId:"AM14_additionalDocumentation"`
	Facility                *segments.B1RequestFacility                `ncpdp:":15" json:"facility" jsonWithId:"AM15_facility"`
	Narrative               *segments.B1RequestNarrative               `ncpdp:":16" json:"narrative" jsonWithId:"AM16_narrative"`
}

type B1ResponseStatus struct {
	ResponseStatus  *segments.B1ResponseStatus  `ncpdp:":21" json:"responseStatus"`
	ResponseClaim   *segments.B1ResponseClaim   `ncpdp:":22" json:"responseClaim"`
	ResponsePricing *segments.B1ResponsePricing `ncpdp:":23" json:"responsePricing"`
	ResponseDUR     *segments.B1ResponseDUR     `ncpdp:":24" json:"responseDUR"`
	COB             *segments.B1ResponseCOB     `ncpdp:":28" json:"responseCOB"`
}

func (b1Response *B1Response) GetResponseStatus(rxNumber string) (*B1ResponseStatus, error) {

	for _, status := range b1Response.Statuses {
		if status.ResponseClaim != nil && status.ResponseClaim.PrescriptionServiceReferenceNumber != nil && *status.ResponseClaim.PrescriptionServiceReferenceNumber == rxNumber {
			return status, nil
		}
	}
	return nil, fmt.Errorf("no status found for rx %s", rxNumber)
}

func (b1Response *B1Response) GetFormattedFullMessage(rxNumber string) (string, error) {
	b1RespStatus, err := b1Response.GetResponseStatus(rxNumber)
	if err != nil {

	}
	additionalMessage := b1RespStatus.ResponseStatus.GetAdditionalMessageInformation()
	dur := b1RespStatus.ResponseDUR.GetFormattedDUR()
	fullMsg := fmt.Sprintf("%s\n%s", additionalMessage, dur)
	return fullMsg, nil
}

func (b1Response *B1Response) IsTransactionAccepted(rxNumber string) (bool, error) {

	if *b1Response.Header.TransactionResponseStatus != "A" {
		return false, nil
	}
	b1RespStatus, err := b1Response.GetResponseStatus(rxNumber)
	if err != nil {
		return false, err
	}

	if b1RespStatus.ResponseStatus.TransactionResponseStatus == nil {
		return false, fmt.Errorf("no transaction response status found for rx %s", rxNumber)
	}

	switch *b1RespStatus.ResponseStatus.TransactionResponseStatus {
	case "A":
		return true, err
	case "C":
		return true, err
	case "D":
		return true, err
	case "P":
		return true, err
	case "Q":
		return true, err
	default:
		return false, nil
	}
}

func (b1Response *B1Response) IsTransactionStatusText(rxNumber string) (string, error) {

	b1RespStatus, err := b1Response.GetResponseStatus(rxNumber)
	if err != nil {
		return "", err
	}

	if b1RespStatus.ResponseStatus.TransactionResponseStatus == nil {
		return "", fmt.Errorf("no transaction response status found for rx %s", rxNumber)
	}

	switch *b1RespStatus.ResponseStatus.TransactionResponseStatus {
	case "A":
		return "A-ACCEPTED", err
	case "C":
		return "C-CAPTURED", err
	case "D":
		return "D-DUPLICATE", err
	case "P":
		return "P-PAID", err
	case "Q":
		return "Q-QUEUED", err
	case "R":
		return "R-REJECTED", err
	default:
		return "", nil
	}
}

func (b1Response *B1Request) GetJsonWithIds() ([]byte, error) {
	var json1 = jsoniter.Config{
		TagKey: "jsonWithId",
	}.Froze()
	jsonWithName, err := json1.Marshal(b1Response)
	return jsonWithName, err
}
