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
	Header  *segments.ResponseHeader `ncpdp:"0:header" json:"header"`
	Message *segments.Message        `ncpdp:"0:20" json:"message"`
	Claims  []*ResponseClaim         `ncpdp:"" json:"claims"`
}

type B2Response struct {
	Header  *segments.ResponseHeader `ncpdp:"0:header" json:"header"`
	Message *segments.Message        `ncpdp:"0:20" json:"message"`
	Claims  []*ResponseClaim         `ncpdp:"" json:"claims"`
}

type B1RequestClaim struct {
	Claim      *segments.RequestClaim `ncpdp:":07" json:"claim"`
	Pricing    *segments.Pricing      `ncpdp:":11" json:"pricing"`
	Prescriber *segments.Prescriber   `ncpdp:":03" json:"prescriber"`
}

type B2RequestClaim struct {
	Claim   *segments.RequestClaim `ncpdp:":07" json:"claim"`
	Pricing *segments.Pricing      `ncpdp:":11" json:"pricing"`
}

type ResponseClaim struct {
	ResponseStatus *segments.ResponseStatus `ncpdp:":21"json:"responseStatus"`
	ResponseClaim  *segments.ResponseClaim  `ncpdp:":22"json:"responseClaim"`
}

func IsTransactionB1ResponseAccepted(msg B1Response) bool {
	var result = false
	header := msg.Header
	if msg.Claims != nil && len(msg.Claims) > 0 && header != nil {
		if msg.Claims[0].ResponseStatus.TransactionResponseStatus != nil && header.TransactionResponseStatus != nil {
			claimStatus := *msg.Claims[0].ResponseStatus.TransactionResponseStatus
			headerReponseStatus := *header.TransactionResponseStatus
			//log.Printf("-> claim.Status: %s", claimStatus)
			//log.Printf("-> header.Status: %s", headerReponseStatus)
			if (claimStatus == "A" || claimStatus == "C" || claimStatus == "D" || claimStatus == "P" || claimStatus == "Q") && headerReponseStatus == "A" {
				result = true
			}
		}
	}
	return result
}
func IsTransactionB2ResponseAccepted(msg B2Response) bool {
	var result = false
	header := msg.Header
	if msg.Claims != nil && len(msg.Claims) > 0 && header != nil {
		if msg.Claims[0].ResponseStatus.TransactionResponseStatus != nil && header.TransactionResponseStatus != nil {
			claimStatus := *msg.Claims[0].ResponseStatus.TransactionResponseStatus
			headerReponseStatus := *header.TransactionResponseStatus
			//log.Printf("-> claim.Status: %s", claimStatus)
			//log.Printf("-> header.Status: %s", headerReponseStatus)
			if claimStatus == "A" && headerReponseStatus == "A" {
				result = true
			}
		}
	}
	return result
}
