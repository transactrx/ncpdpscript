package ncpdpscriptmessages

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"

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
