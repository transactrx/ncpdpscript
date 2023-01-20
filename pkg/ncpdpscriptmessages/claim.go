package ncpdpscriptmessages

import "gihub.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"

type B1 struct {
	Header    *segments.Header    `ncpdp:"0:header"`
	Insurance *segments.Insurance `ncpdp:"0:04"`
	Patient   *segments.Patient   `ncpdp:"0:01"`
	Claims    []*B1Claim          `ncpdp:""`
}

type B1Claim struct {
	Claim      *segments.Claim      `ncpdp:":07"`
	Pricing    *segments.Pricing    `ncpdp:":11"`
	Prescriber *segments.Prescriber `ncpdp:":03"`
}
