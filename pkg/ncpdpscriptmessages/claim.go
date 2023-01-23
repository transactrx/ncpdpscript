package ncpdpscriptmessages

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"

type B1 struct {
	Header    *segments.Header    `ncpdp:"0:header" json:"header"`
	Insurance *segments.Insurance `ncpdp:"0:04" json:"insurance"`
	Patient   *segments.Patient   `ncpdp:"0:01" json:"patient"`
	Claims    []*B1Claim          `ncpdp:"" json:"claims"`
}

type B1Claim struct {
	Claim      *segments.Claim      `ncpdp:":07" json:"claim"`
	Pricing    *segments.Pricing    `ncpdp:":11" json:"pricing"`
	Prescriber *segments.Prescriber `ncpdp:":03" json:"prescriber"`
}
