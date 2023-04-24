package ncpdpscriptmessages

import (
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"
)

type B1Request struct {
	Header    *segments.RequestHeader `ncpdp:"0:header" json:"header"`
	Insurance *segments.Insurance     `ncpdp:"0:04" json:"insurance"`
	Patient   *segments.Patient       `ncpdp:"0:01" json:"patient"`
	Claims    []*B1RequestClaim       `ncpdp:"" json:"claims"`
}

type B2Request struct {
	Header    *segments.RequestHeader `ncpdp:"0:header" json:"header"`
	Insurance *segments.Insurance     `ncpdp:"0:04" json:"insurance"`
	Claims    []*B2RequestClaim       `ncpdp:"" json:"claims"`
}

type B1Response struct {
	Header    *segments.ResponseHeader    `ncpdp:"0:header" json:"header"`
	Message   *segments.Message           `ncpdp:"0:20" json:"message"`
	Insurance *segments.ResponseInsurance `ncpdp:"0:25" json:"insurance"`
	Claims    []*ResponseClaim            `ncpdp:"" json:"claims"`
}

type B2Response struct {
	Header    *segments.ResponseHeader    `ncpdp:"0:header" json:"header"`
	Message   *segments.Message           `ncpdp:"0:20" json:"message"`
	Insurance *segments.ResponseInsurance `ncpdp:"0:25" json:"insurance"`
	Claims    []*ResponseClaim            `ncpdp:"" json:"claims"`
}

type B1RequestClaim struct {
	Claim      *segments.RequestClaim `ncpdp:":07" json:"claim"`
	Pricing    *segments.Pricing      `ncpdp:":11" json:"pricing"`
	Prescriber *segments.Prescriber   `ncpdp:":03" json:"prescriber"`
	Clinical   *segments.Clinical     `ncpdp:":13" json:"prescriber"`
}

type B2RequestClaim struct {
	Claim   *segments.RequestClaim `ncpdp:":07" json:"claim"`
	Pricing *segments.Pricing      `ncpdp:":11" json:"pricing"`
}

type ResponseClaim struct {
	ResponseStatus *segments.ResponseStatus `ncpdp:":21"json:"responseStatus"`
	ResponseClaim  *segments.ResponseClaim  `ncpdp:":22"json:"responseClaim"`
	ResponseDUR    *segments.ResponseDUR    `ncpdp:":24"json:"responseDUR"`
}

func (ncpdp B1Response) GetFullFormattedPayerResponse() (string, error) {
	return GetFormattedPayerResponse(ncpdp)
}

func (ncpdp B2Response) GetFullFormattedPayerResponse() (string, error) {
	return GetFormattedPayerResponse(ncpdp)
}

//func (ncpdp B1Response) IsTransactionResponseAccepted() bool {
//	var result = false
//	header := ncpdp.Header
//	if ncpdp.Claims != nil && len(ncpdp.Claims) > 0 && header != nil {
//		if ncpdp.Claims[0].ResponseStatus != nil && ncpdp.Claims[0].ResponseStatus.TransactionResponseStatus != nil && header.TransactionResponseStatus != nil {
//			claimStatus := *ncpdp.Claims[0].ResponseStatus.TransactionResponseStatus
//			headerReponseStatus := *header.TransactionResponseStatus
//			if (claimStatus == "A" || claimStatus == "C" || claimStatus == "D" || claimStatus == "P" || claimStatus == "Q") && headerReponseStatus == "A" {
//				result = true
//			}
//		}
//	}
//	return result
//}

// todo Only rejected is completed rejected
func (ncpdp B1Response) IsTransactionResponseAccepted() bool {

	header := ncpdp.Header
	if ncpdp.Claims != nil && len(ncpdp.Claims) > 0 && header != nil {
		if header.TransactionResponseStatus != nil {
			headerReponseStatus := *header.TransactionResponseStatus
			if headerReponseStatus == "A" {
				for _, claim := range ncpdp.Claims {
					if claim.IsClaimResponseAccepted("B1") == false {
						return false
					}
				}
			}
		}
	}
	return false
}

func (claim ResponseClaim) IsClaimResponseAccepted(tp string) bool {
	if claim.ResponseStatus != nil && claim.ResponseStatus.TransactionResponseStatus != nil {
		claimStatus := *claim.ResponseStatus.TransactionResponseStatus
		if tp == "B1" {
			if claimStatus == "A" || claimStatus == "C" || claimStatus == "D" || claimStatus == "P" || claimStatus == "Q" {
				return true
			}
		} else if tp == "B2" {
			if claimStatus == "A" {
				return true
			}
		}
	}
	return false
}

func (claim ResponseClaim) getClaimResponseStatus() string {
	if claim.ResponseStatus != nil && claim.ResponseStatus.TransactionResponseStatus != nil {
		claimStatus := *claim.ResponseStatus.TransactionResponseStatus

		switch claimStatus {
		case "A":
			return "A-ACCEPTED"
		case "R":
			return "R-REJECTED"
		case "D":
			return "D-DUPLICATED"
		case "P":
			return "p-PAID"
		case "Q":
			return "Q-QUEUE"
		case "C":
			return "C-CAPTURED"
		default:
			return ""
		}
	}

	return ""
}

//func (ncpdp B2Response) IsTransactionResponseAccepted() bool {
//	var result = false
//	header := ncpdp.Header
//	if ncpdp.Claims != nil && len(ncpdp.Claims) > 0 && header != nil {
//		if ncpdp.Claims[0].ResponseStatus != nil && ncpdp.Claims[0].ResponseStatus.TransactionResponseStatus != nil && header.TransactionResponseStatus != nil {
//			claimStatus := *ncpdp.Claims[0].ResponseStatus.TransactionResponseStatus
//			headerReponseStatus := *header.TransactionResponseStatus
//			//log.Printf("-> claim.Status: %s", claimStatus)
//			//log.Printf("-> header.Status: %s", headerReponseStatus)
//			if claimStatus == "A" && headerReponseStatus == "A" {
//				result = true
//			}
//		}
//	}
//	return result
//}

func (ncpdp B2Response) IsTransactionResponseAccepted() bool {

	header := ncpdp.Header
	if ncpdp.Claims != nil && len(ncpdp.Claims) > 0 && header != nil {
		if header.TransactionResponseStatus != nil {
			headerReponseStatus := *header.TransactionResponseStatus
			if headerReponseStatus == "A" {
				for _, claim := range ncpdp.Claims {
					if claim.IsClaimResponseAccepted("B2") == false {
						return false
					}
				}
			}
		}
	}
	return false
}
