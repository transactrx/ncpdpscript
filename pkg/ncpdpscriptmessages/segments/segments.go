package segments

type ResponseDUR struct {
	SegmentID                *string   `ncpdp:"24AM" json:"segmentID"`
	ResponseCodeCounter      []*int    `ncpdp:"24J6" json:"responseCodeCounter"`
	ReasonForServiceCode     []*string `ncpdp:"24E4" json:"reasonForServiceCode"`
	ClinicalSignificanceCode []*string `ncpdp:"24FS" json:"clinicalSignificanceCode"`
	OtherPharmacyIndicator   []*int    `ncpdp:"24FT" json:"otherPharmacyIndicator"`
	PreviousDateOfFill       []*string `ncpdp:"24FU" json:"previousDateOfFill"`
	PreviousFillQuantity     []*int    `ncpdp:"24FV" json:"previousFillQuantity"`
	DatabaseIndicator        []*string `ncpdp:"24FW" json:"databaseIndicator"`
	OtherPrescriberIndicator []*string `ncpdp:"24FX" json:"otherPrescriberIndicator"`
	Message                  []*string `ncpdp:"24FY" json:"message"`
	AdditionalText           []*string `ncpdp:"24NS" json:"additionalText"`
}
