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

type Insurance struct {
	SegmentId                          *string `ncpdp:"04AM" json:"segmentId" jsonWithId:"111_AM_segmentId"`
	CardHolderId                       *string `ncpdp:"04C2" json:"cardHolderId" jsonWithId:"302_C2_cardHolderId"`
	CardHolderFirstName                *string `ncpdp:"04CC" json:"cardHolderFirstName" jsonWithId:"312_CC_cardHolderFirstName"`
	CardHolderLastName                 *string `ncpdp:"04CD" json:"cardHolderLastName" jsonWithId:"313_CD_cardHolderLastName"`
	HomePlan                           *string `ncpdp:"04CE" json:"homePlan" jsonWithId:"314_CE_homePlan"`
	PlanId                             *string `ncpdp:"04FO" json:"planId" jsonWithId:"524_FO_planId"`
	EligibilityClarificationCode       *string `ncpdp:"04C9" json:"eligibilityClarificationCode" jsonWithId:"309_C9_eligibilityClarificationCode"`
	GroupID                            *string `ncpdp:"04C1" json:"groupID" jsonWithId:"301_C1_groupID"`
	PersonCode                         *string `ncpdp:"04C3" json:"personCode" jsonWithId:"303_C3_personCode"`
	RelationshipCode                   *string `ncpdp:"04C6" json:"relationshipCode" jsonWithId:"306_C6_relationshipCode"`
	OtherPayerBin                      *string `ncpdp:"04MG" json:"otherPayerBin" jsonWithId:"990_MG_otherPayerBin"`
	OtherPayerPCN                      *string `ncpdp:"04MH" json:"otherPayerPCN" jsonWithId:"991_MH_otherPayerPCN"`
	OtherPayerCardholderId             *string `ncpdp:"04NU" json:"otherPayerCardholderId" jsonWithId:"356_NU_otherPayerCardholderId"`
	OtherPayerGroupId                  *string `ncpdp:"04MJ" json:"otherPayerGroupId" jsonWithId:"992_MJ_otherPayerGroupId"`
	MedigapId                          *string `ncpdp:"042A" json:"medigapId" jsonWithId:"359_2A_medigapId"`
	MedicaidIndicator                  *string `ncpdp:"042B" json:"medicaidIndicator" jsonWithId:"360_2B_medicaidIndicator"`
	ProviderAcceptsAssignmentIndicator *string `ncpdp:"042D" json:"providerAcceptsAssignmentIndicator" jsonWithId:"361_2D_providerAcceptsAssignmentIndicator"`
	CmsPartDDefinedQualifiedFacility   *string `ncpdp:"04G2" json:"cmsPartDDefinedQualifiedFacility" jsonWithId:"997_G2_cmsPartDDefinedQualifiedFacility"`
	MedicaidIdNumber                   *string `ncpdp:"04N5" json:"medicaidIdNumber" jsonWithId:"115_N5_medicaidIdNumber"`
	MedicaidAgencyNumber               *string `ncpdp:"04N6" json:"medicaidAgencyNumber" jsonWithId:"116_N6_medicaidAgencyNumber"`
}

type Patient struct {
	SegmentID          *string `ncpdp:"01AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	PatientIDQualifier *string `ncpdp:"01CX" json:"patientIDQualifier" jsonWithId:"331_CX_patientIDQualifier"`
	PatientID          *string `ncpdp:"01CY" json:"patientID" jsonWithId:"332_CY_patientID"`
	DateOfBirth        *string `ncpdp:"01C4" json:"dateOfBirth" jsonWithId:"304_C4_dateOfBirth"`
	GenderCode         *string `ncpdp:"01C5" json:"genderCode" jsonWithId:"305_C5_genderCode"`
	FirstName          *string `ncpdp:"01CA" json:"firstName" jsonWithId:"310_CA_firstName"`
	LastName           *string `ncpdp:"01CB" json:"lastName" jsonWithId:"311_CB_lastName"`
	StreetAddress      *string `ncpdp:"01CM" json:"streetAddress" jsonWithId:"322_CM_streetAddress"`
	City               *string `ncpdp:"01CN" json:"city" jsonWithId:"323_CN_city"`
	State              *string `ncpdp:"01CO" json:"state" jsonWithId:"324_CO_state"`
	ZipCode            *string `ncpdp:"01CP" json:"zipCode" jsonWithId:"325_CP_zipCode"`
	PhoneNumber        *string `ncpdp:"01CQ" json:"phoneNumber" jsonWithId:"326_CQ_phoneNumber"`
	LocationCode       *string `ncpdp:"01C7" json:"placeOfService" jsonWithId:"307_C7_placeOfService"`
	EmployerID         *string `ncpdp:"01CZ" json:"employerID" jsonWithId:"333_CZ_employerID"`
	Smoker             *string `ncpdp:"011C" json:"smoker" jsonWithId:"334_1C_smoker"`
	PregnancyIndicator *string `ncpdp:"012C" json:"pregnancyIndicator" jsonWithId:"335_2C_pregnancyIndicator"`
	Email              *string `ncpdp:"01HN" json:"email" jsnoWithId:"350_HN_email"`
	Residence          *string `ncpdp:"014X" json:"residence" jsonWithId:"384_4X_residence"`
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
