package segments

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

type RequestHeader struct {
	Bin                        *string `ncpdp:"header:bin" json:"bin"`
	VersionRel                 *string `ncpdp:"header:versionRel" json:"versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode" json:"transactionCode"`
	Pcn                        *string `ncpdp:"header:pcn" json:"pcn"`
	TransactionCount           *int    `ncpdp:"header:transactionCount" json:"transactionCount"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier" json:"serviceProviderIdQualifier"`
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
	SegmentId                          *string `ncpdp:"04AM" json:"segmentId"`
	CardHolderId                       *string `ncpdp:"04C2" json:"cardHolderId"`
	CardHolderFirstName                *string `ncpdp:"04CC" json:"cardHolderFirstName"`
	CardHolderLastName                 *string `ncpdp:"04CD" json:"cardHolderLastName"`
	HomePlan                           *string `ncpdp:"04CE" json:"homePlan"`
	PlanId                             *string `ncpdp:"04FO" json:"planId"`
	EligibilityClarificationCode       *string `ncpdp:"04C9" json:"eligibilityClarificationCode"`
	GroupID                            *string `ncpdp:"04C1" json:"groupID"`
	PersonCode                         *string `ncpdp:"04C3" json:"personCode"`
	RelationshipCode                   *string `ncpdp:"04C6" json:"relationshipCode"`
	OtherPayerBin                      *string `ncpdp:"04MG" json:"otherPayerBin"`
	OtherPayerPCN                      *string `ncpdp:"04MH" json:"otherPayerPCN"`
	OtherPayerCardholderId             *string `ncpdp:"04NU" json:"otherPayerCardholderId"`
	OtherPayerGroupId                  *string `ncpdp:"04MJ" json:"otherPayerGroupId"`
	MedigapId                          *string `ncpdp:"042A" json:"medigapId"`
	MedicaidIndicator                  *string `ncpdp:"042B" json:"medicaidIndicator"`
	ProviderAcceptsAssignmentIndicator *string `ncpdp:"042D" json:"providerAcceptsAssignmentIndicator"`
	CmsPartDDefinedQualifiedFacility   *string `ncpdp:"04G2" json:"cmsPartDDefinedQualifiedFacility"`
	MedicaidIdNumber                   *string `ncpdp:"04N5" json:"medicaidIdNumber"`
	MedicaidAgencyNumber               *string `ncpdp:"04N6" json:"medicaidAgencyNumber"`
}

type Patient struct {
	SegmentID          *string `ncpdp:"01AM" json:"segmentID"`
	PatientIDQualifier *string `ncpdp:"01CX" json:"patientIDQualifier"`
	PatientID          *string `ncpdp:"01CY" json:"patientID"`
	DateOfBirth        *string `ncpdp:"01C4" json:"dateOfBirth"`
	GenderCode         *string `ncpdp:"01C5" json:"genderCode"`
	FirstName          *string `ncpdp:"01CA" json:"firstName"`
	LastName           *string `ncpdp:"01CB" json:"lastName"`
	StreetAddress      *string `ncpdp:"01CM" json:"streetAddress"`
	City               *string `ncpdp:"01CN" json:"city"`
	State              *string `ncpdp:"01CO" json:"state"`
	ZipCode            *string `ncpdp:"01CP" json:"zipCode"`
	PhoneNumber        *string `ncpdp:"01CQ" json:"phoneNumber"`
	LocationCode       *string `ncpdp:"01C7" json:"locationCode"`
	EmployerID         *string `ncpdp:"01CZ" json:"employerID"`
	Smoker             *string `ncpdp:"011C" json:"smoker"`
	PregnancyIndicator *string `ncpdp:"012C" json:"pregnancyIndicator"`
	Email              *string `ncpdp:"01HN" json:"email"`
	Residence          *string `ncpdp:"014X" json:"residence"`
}

type RequestClaim struct {
	SegmentID                             *string                       `ncpdp:"07AM" json:"segmentID"`
	RxServiceQualifier                    *string                       `ncpdp:"07EM" json:"rxServiceQualifier"`
	RxServiceReferenceNumber              *string                       `ncpdp:"07D2" json:"rxServiceReferenceNumber"`
	ProductServiceQualifier               *string                       `ncpdp:"07E1" json:"productServiceQualifier"`
	ProductServiceID                      *string                       `ncpdp:"07D7" json:"productServiceID"`
	AssociatedRxServiceRefNum             *string                       `ncpdp:"07EN" json:"associatedRxServiceRefNum"`
	AssociatedRxServiceDate               *string                       `ncpdp:"07EP" json:"associatedRxServiceDate"`
	ProcedureModifierCount                *int                          `ncpdp:"07SE" json:"procedureModifierCount"`
	ProcedureModifierCode                 []*string                     `ncpdp:"07ER" json:"procedureModifierCode"`
	QuantityDispensed                     *scriptparser.UnsignedNumeric `ncpdp:"07E7" json:"quantityDispensed"`
	FillNumber                            *int                          `ncpdp:"07D3" json:"fillNumber"`
	DaysSupply                            *int                          `ncpdp:"07D5" json:"daysSupply"`
	CompoundCode                          *int                          `ncpdp:"07D6" json:"compoundCode"`
	DispenseAsWritten                     *int                          `ncpdp:"07D8" json:"dispenseAsWritten"`
	DateWritten                           *string                       `ncpdp:"07DE" json:"dateWritten"`
	NumberOfRefillsAuthorized             *int                          `ncpdp:"07DF" json:"numberOfRefillsAuthorized"`
	RxOriginCode                          *int                          `ncpdp:"07DJ" json:"rxOriginCode"`
	SubmissionClarificationCodeCount      *int                          `ncpdp:"07NX" json:"submissionClarificationCodeCount"`
	SubmissionClarificationCode           []*int                        `ncpdp:"07DK" json:"submissionClarificationCode"`
	QuantityPrescriber                    *scriptparser.UnsignedNumeric `ncpdp:"07ET" json:"quantityPrescriber"`
	OtherCoverageCode                     *int                          `ncpdp:"07C8" json:"otherCoverageCode"`
	UnitDoseIndicator                     *int                          `ncpdp:"07DT" json:"unitDoseIndicator"`
	OrigPrescribedProductServiceQualifier *int                          `ncpdp:"07EJ" json:"origPrescribedProductServiceQualifier"`
	OrigPrescribedProductServiceCode      *string                       `ncpdp:"07EA" json:"origPrescribedProductServiceCode"`
	OrigPrescribedQuantity                *scriptparser.UnsignedNumeric `ncpdp:"07EB" json:"origPrescribedQuantity"`
	AlternateID                           *string                       `ncpdp:"07CW" json:"alternateID"`
	ScheduledRxID                         *string                       `ncpdp:"07EK" json:"scheduledRxID"`
	UnitOfMeasure                         *string                       `ncpdp:"0728" json:"unitOfMeasure"`
	LevelOfService                        *string                       `ncpdp:"07DI" json:"levelOfService"`
	PriorAuthorizationTypeCode            *int                          `ncpdp:"07EU" json:"priorAuthorizationTypeCode"`
	PriorAuthorizationNumberSubmitted     *string                       `ncpdp:"07EV" json:"priorAuthorizationNumberSubmitted"`
	IntermediaryAuthorizionTypeID         *string                       `ncpdp:"07EW" json:"intermediaryAuthorizionTypeID"`
	IntermediaryAuthorizionID             *string                       `ncpdp:"07EX" json:"intermediaryAuthorizionID"`
	DispensingStatus                      *string                       `ncpdp:"07HD" json:"dispensingStatus"`
	QuantityIntendedToBeDispersed         *scriptparser.UnsignedNumeric `ncpdp:"07HF" json:"quantityIntendedToBeDispersed"`
	DaysSupplyIntendedToBeDispersed       *int                          `ncpdp:"07HG" json:"daysSupplyIntendedToBeDispersed"`
	DelayReasonCode                       *string                       `ncpdp:"07NV" json:"delayReasonCode"`
	TransactionReferenceNumber            *string                       `ncpdp:"07K5" json:"transactionReferenceNumber"`
	PatientAssignmentIndicator            *string                       `ncpdp:"07MT" json:"patientAssignmentIndicator"`
	RouteOfAdministration                 *string                       `ncpdp:"07E2" json:"routeOfAdministration"`
	CompoundType                          *string                       `ncpdp:"07G1" json:"compoundType"`
	TransactionCtrlNumber                 *string                       `ncpdp:"07N4" json:"transactionCtrlNumber"`
	PharmacyServiceType                   *string                       `ncpdp:"07U7" json:"pharmacyServiceType"`
}

type Pricing struct {
	SegmentID                            *string                    `ncpdp:"11AM" json:"segmentID"`
	IngredientCostSubmitted              *scriptparser.Currency64   `ncpdp:"11D9" json:"ingredientCostSubmitted"`
	DispensingFeeSubmitted               *scriptparser.Currency64   `ncpdp:"11DC" json:"dispensingFeeSubmitted"`
	ProfessionServiceFeeSubmitted        *scriptparser.Currency64   `ncpdp:"11BE" json:"professionServiceFeeSubmitted"`
	PatientPaidAmount                    *scriptparser.Currency64   `ncpdp:"11DX" json:"patientPaidAmount"`
	IncentiveAmountSubmitted             *scriptparser.Currency64   `ncpdp:"11E3" json:"incentiveAmountSubmitted"`
	FlatSalesTaxAmountSubmitted          *scriptparser.Currency64   `ncpdp:"11HA" json:"flatSalesTaxAmountSubmitted"`
	OtherAmountClaimedSubmittedCount     *int64                     `ncpdp:"11H7" json:"otherAmountClaimedSubmittedCount"`
	OtherAmountClaimedSubmittedQualifier []*string                  `ncpdp:"11H8" json:"otherAmountClaimedSubmittedQualifier"`
	OtherAmountClaimedSubmitted          []*scriptparser.Currency64 `ncpdp:"11H9" json:"otherAmountClaimedSubmitted"`
	PercentageSalesTaxAmountSubmitted    *scriptparser.Currency64   `ncpdp:"11GE" json:"percentageSalesTaxAmountSubmitted"`
	PercentageSalesTaxBasisSubmitted     *string                    `ncpdp:"11JE" json:"percentageSalesTaxBasisSubmitted"`
	PercentageSalesTaxRateSubmitted      *scriptparser.Currency64   `ncpdp:"11HE" json:"percentageSalesTaxRateSubmitted"`
	UsualAndCustomaryCharge              *scriptparser.Currency64   `ncpdp:"11DQ" json:"usualAndCustomaryCharge"`
	GrossAmountDue                       *scriptparser.Currency64   `ncpdp:"11DU" json:"grossAmountDue"`
	BasisOfCostDetermination             *string                    `ncpdp:"11DN" json:"basisOfCostDetermination"`
	MedicaidPaidAmount                   *scriptparser.Currency64   `ncpdp:"11N3" json:"medicaidPaidAmount"`
}

// Prescriber segment
type Prescriber struct {
	SegmentID                    *string `ncpdp:"03AM" json:"segmentID"`
	PrescriberIdQualifier        *string `ncpdp:"03EZ" json:"prescriberIdQualifier"`
	PrescriberId                 *string `ncpdp:"03DB" json:"prescriberId"`
	PrescriberLastname           *string `ncpdp:"03DR" json:"prescriberLastname"`
	PrescriberPhoneNumber        *string `ncpdp:"03PM" json:"prescriberPhoneNumber"`
	PrimaryCareProviderQualifier *string `ncpdp:"032E" json:"primaryCareProviderQualifier"`
	PrimaryCareProvider          *string `ncpdp:"03DL" json:"primaryCareProvider"`
	PrimaryCareProviderLastname  *string `ncpdp:"034E" json:"primaryCareProviderLastname"`
	PrescriberFirstname          *string `ncpdp:"032J" json:"prescriberFirstname"`
	PrescriberStreetAddress      *string `ncpdp:"032k" json:"prescriberStreetAddress"`
	PrescriberCity               *string `ncpdp:"032M" json:"prescriberCity"`
	PrescriberState              *string `ncpdp:"032N" json:"prescriberState"`
	PrescriberZip                *string `ncpdp:"032P" json:"prescriberZip"`
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
	RejectFieldOccuranceIndicator   *string   `ncpdp:"214F" json:"rejectFieldOccuranceIndicator"`
	ApprovedMessageCodeCount        *int      `ncpdp:"215F" json:"approvedMessageCodeCount"`
	ApprovedMessageCode             *string   `ncpdp:"216F" json:"approvedMessageCode"`
	AdditionalMessageInfoCount      *int      `ncpdp:"21UF" json:"additionalMessageInfoCount"`
	AdditionalMessageInfoQualifier  *string   `ncpdp:"21UH" json:"additionalMessageInfoQualifier"`
	AdditionalMessageInfo           *string   `ncpdp:"21FQ" json:"additionalMessageInfo"`
	AdditionalMessageInfoContinuity *string   `ncpdp:"21UG" json:"additionalMessageInfoContinuity"`
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

type Clinical struct {
	SegmentID                  *string   `ncpdp:"13AM" json:"segmentID"`
	DiagnosisCodeCount         *int      `ncpdp:"13VE" json:"diagnosisCodeCount"`
	DiagnosisCodeQualifier     []*string `ncpdp:"13WE" json:"diagnosisCodeQualifier"`
	DiagnosisCode              []*string `ncpdp:"13DO" json:"diagnosisCode"`
	ClinicalInformationCounter *int      `ncpdp:"13XE" json:"clinicalInformationCounter"`
	MeasurementDate            *string   `ncpdp:"13ZE" json:"measurementDate"`
	MeasurementTime            *string   `ncpdp:"13H1" json:"measurementTime"`
	MeasurementDimension       []*string `ncpdp:"13H2" json:"measurementDimension"`
	MeasurementUnit            []*string `ncpdp:"13H3" json:"measurementUnit"`
	MeasurementValue           []*string `ncpdp:"13H4" json:"measurementValue"`
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
