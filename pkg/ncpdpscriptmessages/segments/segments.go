package segments

import "gihub.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

type Header struct {
	Bin                        *string `ncpdp:"header:bin"`
	VersionRel                 *string `ncpdp:"header:versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode"`
	Pcn                        *string `ncpdp:"header:pcn"`
	TransactionCount           *int    `ncpdp:"header:transactionCount"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier"`
	ServiceProviderId          *string `ncpdp:"header:serviceProviderId"`
	Dos                        *string `ncpdp:"header:dos"`
}

type Insurance struct {
	SegmentId                          *string `ncpdp:"04AM"`
	CardHolderId                       *string `ncpdp:"04C2"`
	CardHolderFirstName                *string `ncpdp:"04CC"`
	CardHolderLastName                 *string `ncpdp:"04CD"`
	HomePlan                           *string `ncpdp:"04CE"`
	PlanId                             *string `ncpdp:"04FO"`
	EligibilityClarificationCode       *string `ncpdp:"04C9"`
	GroupID                            *string `ncpdp:"04C1"`
	PersonCode                         *string `ncpdp:"04C3"`
	RelationshipCode                   *string `ncpdp:"04C6"`
	OtherPayerBin                      *string `ncpdp:"04MG"`
	OtherPayerPCN                      *string `ncpdp:"04MH"`
	OtherPayerCardholderId             *string `ncpdp:"04NU"`
	OtherPayerGroupId                  *string `ncpdp:"04MJ"`
	MedigapId                          *string `ncpdp:"042A"`
	MedicaidIndicator                  *string `ncpdp:"042B"`
	ProviderAcceptsAssignmentIndicator *string `ncpdp:"042D"`
	CmsPartDDefinedQualifiedFacility   *string `ncpdp:"04G2"`
	MedicaidIdNumber                   *string `ncpdp:"04N5"`
	MedicaidAgencyNumber               *string `ncpdp:"04N6"`
}

type Patient struct {
	SegmentID          *string `ncpdp:"01AM"`
	PatientIDQualifier *string `ncpdp:"01CX"`
	PatientID          *string `ncpdp:"01CY"`
	DateOfBirth        *string `ncpdp:"01C4"`
	GenderCode         *string `ncpdp:"01C5"`
	FirstName          *string `ncpdp:"01CA"`
	LastName           *string `ncpdp:"01CB"`
	StreetAddress      *string `ncpdp:"01CM"`
	City               *string `ncpdp:"01CN"`
	State              *string `ncpdp:"01CO"`
	ZipCode            *string `ncpdp:"01CP"`
	PhoneNumber        *string `ncpdp:"01CQ"`
	LocationCode       *string `ncpdp:"01C7"`
	EmployerID         *string `ncpdp:"01CZ"`
	Smoker             *string `ncpdp:"011C"`
	PregnancyIndicator *string `ncpdp:"012C"`
	Email              *string `ncpdp:"01HN"`
	Residence          *string `ncpdp:"014X"`
}

type Claim struct {
	SegmentID                             *string   `ncpdp:"07AM"`
	RxServiceQualifier                    *string   `ncpdp:"07EM"`
	RxServiceReferenceNumber              *string   `ncpdp:"07D2"`
	ProductServiceQualifier               *string   `ncpdp:"07E1"`
	ProductServiceID                      *string   `ncpdp:"07D7"`
	AssociatedRxServiceRefNum             *string   `ncpdp:"07EN"`
	AssociatedRxServiceDate               *string   `ncpdp:"07EP"`
	ProcedureModifierCount                *int      `ncpdp:"07SE"`
	ProcedureModifierCode                 []*string `ncpdp:"07ER"`
	QuantityDispensed                     *float64  `ncpdp:"07E7"`
	FillNumber                            *int      `ncpdp:"07D3"`
	DaysSupply                            *int      `ncpdp:"07D5"`
	CompoundCode                          *int      `ncpdp:"07D6"`
	DispenseAsWritten                     *int      `ncpdp:"07D8"`
	DateWritten                           *string   `ncpdp:"07DE"`
	NumberOfRefillsAuthorized             *int      `ncpdp:"07DF"`
	RxOriginCode                          *int      `ncpdp:"07DJ"`
	SubmissionClarificationCodeCount      *int      `ncpdp:"07NX"`
	SubmissionClarificationCode           []*int    `ncpdp:"07DK"`
	QuantityPrescriber                    *float64  `ncpdp:"07ET"`
	OtherCoverageCode                     *int      `ncpdp:"07C8"`
	UnitDoseIndicator                     *int      `ncpdp:"07DT"`
	OrigPrescribedProductServiceQualifier *int      `ncpdp:"07EJ"`
	OrigPrescribedProductServiceCode      *string   `ncpdp:"07EA"`
	OrigPrescribedQuantity                *float64  `ncpdp:"07EB"`
	AlternateID                           *string   `ncpdp:"07CW"`
	ScheduledRxID                         *string   `ncpdp:"07EK"`
	UnitOfMeasure                         *string   `ncpdp:"0728"`
	LevelOfService                        *string   `ncpdp:"07DI"`
	PriorAuthorizationTypeCode            *int      `ncpdp:"07EU"`
	PriorAuthorizationNumberSubmitted     *string   `ncpdp:"07EV"`
	IntermediaryAuthorizionTypeID         *string   `ncpdp:"07EW"`
	IntermediaryAuthorizionID             *string   `ncpdp:"07EX"`
	DispensingStatus                      *string   `ncpdp:"07HD"`
	QuantityIntendedToBeDispersed         *int      `ncpdp:"07HF"`
	DaysSupplyIntendedToBeDispersed       *int      `ncpdp:"07HG"`
	DelayReasonCode                       *string   `ncpdp:"07NV"`
	TransactionReferenceNumber            *string   `ncpdp:"07K5"`
	PatientAssignmentIndicator            *string   `ncpdp:"07MT"`
	RouteOfAdministration                 *string   `ncpdp:"07E2"`
	CompoundType                          *string   `ncpdp:"07G1"`
	TransactionCtrlNumber                 *string   `ncpdp:"07N4"`
	PharmacyServiceType                   *string   `ncpdp:"07U7"`
}

type Pricing struct {
	SegmentID                            *string                    `ncpdp:"11AM"`
	IngredientCostSubmitted              *scriptparser.Currency64   `ncpdp:"11D9"`
	DispensingFeeSubmitted               *scriptparser.Currency64   `ncpdp:"11DC"`
	ProfessionServiceFeeSubmitted        *scriptparser.Currency64   `ncpdp:"11BE"`
	PatientPaidAmount                    *scriptparser.Currency64   `ncpdp:"11DX"`
	IncentiveAmountSubmitted             *scriptparser.Currency64   `ncpdp:"11E3"`
	FlatSalesTaxAmountSubmitted          *scriptparser.Currency64   `ncpdp:"11HA"`
	OtherAmountClaimedSubmittedCount     *int64                     `ncpdp:"11H7"`
	OtherAmountClaimedSubmittedQualifier []*string                  `ncpdp:"11H8"`
	OtherAmountClaimedSubmitted          []*scriptparser.Currency64 `ncpdp:"11H9"`
	PercentageSalesTaxAmountSubmitted    *scriptparser.Currency64   `ncpdp:"11GE"`
	PercentageSalesTaxBasisSubmitted     *string                    `ncpdp:"11JE"`
	PercentageSalesTaxRateSubmitted      *string                    `ncpdp:"11HE"`
	UsualAndCustomaryCharge              *scriptparser.Currency64   `ncpdp:"11DQ"`
	GrossAmountDue                       *scriptparser.Currency64   `ncpdp:"11DU"`
	BasisOfCostDetermination             *string                    `ncpdp:"11DN"`
	MedicaidPaidAmount                   *scriptparser.Currency64   `ncpdp:"11N3"`
}

// Prescriber segment
type Prescriber struct {
	SegmentID                    *string `ncpdp:"03AM"`
	PrescriberIdQualifier        *string `ncpdp:"03EZ"`
	PrescriberId                 *string `ncpdp:"03DB"`
	PrescriberLastname           *string `ncpdp:"03DR"`
	PrescriberPhoneNumber        *string `ncpdp:"03PM"`
	PrimaryCareProviderQualifier *string `ncpdp:"032E"`
	PrimaryCareProvider          *string `ncpdp:"03DL"`
	PrimaryCareProviderLastname  *string `ncpdp:"034E"`
	PrescriberFirstname          *string `ncpdp:"032J"`
	PrescriberStreetAddress      *string `ncpdp:"032k"`
	PrescriberCity               *string `ncpdp:"032M"`
	PrescriberState              *string `ncpdp:"032N"`
	PrescriberZip                *string `ncpdp:"032P"`
}
