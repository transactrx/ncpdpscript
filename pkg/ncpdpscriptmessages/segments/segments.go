package segments

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

type RequestHeader struct {
	Bin                        *string `ncpdp:"header:bin" json:"bin" jsonWithId:"101_A1_bin"`
	VersionRel                 *string `ncpdp:"header:versionRel" json:"versionRel" jsonWithId:"101_A2_versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode" json:"transactionCode" jsonWithId:"101_A3_transactionCode"`
	Pcn                        *string `ncpdp:"header:pcn" json:"pcn" jsonWithId:"101_A4_pcn"`
	TransactionCount           *int    `ncpdp:"header:transactionCount" json:"transactionCount" jsonWithId:"101_A9_transactionCount"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier" json:"serviceProviderIdQualifier" jsonWithId:"101_B2_serviceProviderIdQualifier"`
	ServiceProviderId          *string `ncpdp:"header:serviceProviderId" json:"serviceProviderId"`
	Dos                        *string `ncpdp:"header:dos" json:"dos"`
}

type ResponseHeader struct {
	VersionRel                 *string `ncpdp:"header:versionRel" json:"versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode" json:"transactionCode"`
	TransactionCount           *int    `ncpdp:"header:transactionCount" json:"transactionCount"`
	TransactionResponseStatus  *string `ncpdp:"header:transactionResponseStatus" json:"transactionResponseStatus"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier" json:"serviceProviderIdQualifier"`
	ServiceProviderId          *string `ncpdp:"header:serviceProviderId" json:"serviceProviderId"`
	Dos                        *string `ncpdp:"header:dos" json:"dos"`
}

type Message struct {
	SegmentID *string `ncpdp:"20AM" json:"segmentID"`
	Message   *string `ncpdp:"20F4" json:"message"`
}

type ResponseStatus struct {
	SegmentID                       *string   `ncpdp:"21AM" json:"segmentID"`
	TransactionResponseStatus       *string   `ncpdp:"21AN" json:"transactionResponseStatus"`
	AuthorizationNumber             *string   `ncpdp:"21F3" json:"authorizationNumber"`
	RejectCount                     *int      `ncpdp:"21FA" json:"rejectCount"`
	RejectCode                      []*string `ncpdp:"21FB" json:"rejectCode"`
	RejectFieldOccurrenceIndicator  *string   `ncpdp:"214F" json:"rejectFieldOccurrenceIndicator"`
	ApprovedMessageCodeCount        *int      `ncpdp:"215F" json:"approvedMessageCodeCount"`
	ApprovedMessageCode             *string   `ncpdp:"216F" json:"approvedMessageCode"`
	AdditionalMessageInfoCount      *int      `ncpdp:"21UF" json:"additionalMessageInfoCount"`
	AdditionalMessageInfoQualifier  *[]string `ncpdp:"21UH" json:"additionalMessageInfoQualifier"`
	AdditionalMessageInfo           *[]string `ncpdp:"21FQ" json:"additionalMessageInfo"`
	AdditionalMessageInfoContinuity *[]string `ncpdp:"21UG" json:"additionalMessageInfoContinuity"`
	HelpDeskPhoneNumberQualifier    *string   `ncpdp:"217F" json:"helpDeskPhoneNumberQualifier"`
	HelpDeskPhoneNumber             *string   `ncpdp:"218F" json:"helpDeskPhoneNumber"`
	TransactionReferenceNumber      *string   `ncpdp:"21K5" json:"transactionReferenceNumber"`
	InternalControlNumber           *string   `ncpdp:"21A7" json:"internalControlNumber"`
	Url                             *string   `ncpdp:"21MA" json:"url"`
}

type ResponseClaim struct {
	SegmentID                          *string                  `ncpdp:"22AM" json:"segmentID"`
	RxServiceQualifier                 *string                  `ncpdp:"22EM" json:"rxServiceQualifier"`
	RxServiceReferenceNumber           *string                  `ncpdp:"22D2" json:"rxServiceReferenceNumber"`
	PreferredProductCount              *int                     `ncpdp:"229F" json:"preferredProductCount"`
	PreferredProductIdQualifier        []*string                `ncpdp:"22AP" json:"preferredProductIdQualifier"`
	PreferredProductId                 []*string                `ncpdp:"22AR" json:"preferredProductId"`
	PreferredProductIncentive          *scriptparser.Currency64 `ncpdp:"22AS" json:"preferredProductIncentive"`
	PreferredProductCostShareIncentive *scriptparser.Currency64 `ncpdp:"22AT" json:"preferredProductCostShareIncentive"`
	PreferredProductDescription        []*string                `ncpdp:"22AU" json:"preferredProductDescription"`
	MedicaidInternalCtrlNumber         *string                  `ncpdp:"22E7" json:"medicaidInternalCtrlNumber"`
}

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

type ResponseInsurance struct {
	SegmentId              *string `ncpdp:"25AM" json:"segmentId"`
	GroupID                *string `ncpdp:"25C1" json:"groupID"`
	PlanId                 *string `ncpdp:"25FO" json:"planId"`
	NetworkReimbursementId *string `ncpdp:"252F" json:"networkReimbursementId"`
	PayerQualifier         *string `ncpdp:"25J7" json:"payerQualifier"`
	PayerId                *string `ncpdp:"25J8" json:"payerId"`
	MedicaidIdNumber       *string `ncpdp:"25N5" json:"medicaidIdNumber"`
	MedicaidAgencyNumber   *string `ncpdp:"25N6" json:"medicaidAgencyNumber"`
	CardHolderId           *string `ncpdp:"25C2" json:"cardHolderId"`
}
