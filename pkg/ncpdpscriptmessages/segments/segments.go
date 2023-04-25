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

type RequestClaim struct {
	SegmentID                             *string                       `ncpdp:"07AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	RxServiceQualifier                    *string                       `ncpdp:"07EM" json:"rxServiceQualifier" jsonWithId:"455_EM_rxServiceQualifier"`
	RxServiceReferenceNumber              *string                       `ncpdp:"07D2" json:"rxServiceReferenceNumber" jsonWithId:"406_D2_rxServiceReferenceNumber"`
	ProductServiceQualifier               *string                       `ncpdp:"07E1" json:"productServiceQualifier" jsonWithId:"438_E1_productServiceQualifier"`
	ProductServiceID                      *string                       `ncpdp:"07D7" json:"productServiceID" jsonWithId:"407_D7_productServiceID"`
	AssociatedRxServiceRefNum             *string                       `ncpdp:"07EN" json:"associatedRxServiceRefNum" jsonWithId:"456_EN_associatedRxServiceRefNum"`
	AssociatedRxServiceDate               *string                       `ncpdp:"07EP" json:"associatedRxServiceDate" jsonWithId:"457_EP_associatedRxServiceDate"`
	ProcedureModifierCount                *int                          `ncpdp:"07SE" json:"procedureModifierCount" jsonWithId:"458_SE_procedureModifierCount"`
	ProcedureModifierCode                 []*string                     `ncpdp:"07ER" json:"procedureModifierCode" jsonWithId:"459_ER_procedureModifierCode"`
	QuantityDispensed                     *scriptparser.UnsignedNumeric `ncpdp:"07E7" json:"quantityDispensed" jsonWithId:"442_E7_quantityDispensed"`
	FillNumber                            *int                          `ncpdp:"07D3" json:"fillNumber" jsonWithId:"403_D3_fillNumber"`
	DaysSupply                            *int                          `ncpdp:"07D5" json:"daysSupply" jsonWithId:"405_D5_daysSupply"`
	CompoundCode                          *int                          `ncpdp:"07D6" json:"compoundCode" jsonWithId:"406_D6_compoundCode"`
	DispenseAsWritten                     *int                          `ncpdp:"07D8" json:"dispenseAsWritten" jsonWithId:"408_D8_dispenseAsWritten"`
	DateWritten                           *string                       `ncpdp:"07DE" json:"dateWritten" jsonWithId:"414_DE_dateWritten"`
	NumberOfRefillsAuthorized             *int                          `ncpdp:"07DF" json:"numberOfRefillsAuthorized" jsonWithId:"415_DF_numberOfRefillsAuthorized"`
	RxOriginCode                          *int                          `ncpdp:"07DJ" json:"rxOriginCode" jsonWithId:"419_DJ_rxOriginCode"`
	SubmissionClarificationCodeCount      *int                          `ncpdp:"07NX" json:"submissionClarificationCodeCount" jsonWithId:"354_NX_submissionClarificationCodeCount"`
	SubmissionClarificationCode           []*int                        `ncpdp:"07DK" json:"submissionClarificationCode" jsonWithId:"420_DK_submissionClarificationCode"`
	QuantityPrescriber                    *scriptparser.UnsignedNumeric `ncpdp:"07ET" json:"quantityPrescriber" jsonWithId:"460_ET_quantityPrescriber"`
	OtherCoverageCode                     *int                          `ncpdp:"07C8" json:"otherCoverageCode" jsonWithId:"308_C8_otherCoverageCode"`
	UnitDoseIndicator                     *int                          `ncpdp:"07DT" json:"specialPackageIndicator" jsonWithId:"429_DT_specialPackageIndicator"`
	OrigPrescribedProductServiceQualifier *int                          `ncpdp:"07EJ" json:"origPrescribedProductServiceQualifier" jsonWithId:"453_EJ_origPrescribedProductServiceQualifier"`
	OrigPrescribedProductServiceCode      *string                       `ncpdp:"07EA" json:"origPrescribedProductServiceCode" jsonWithId:"445_EA_origPrescribedProductServiceCode"`
	OrigPrescribedQuantity                *scriptparser.UnsignedNumeric `ncpdp:"07EB" json:"origPrescribedQuantity" jsonWithId:"446_EB_origPrescribedQuantity"`
	AlternateID                           *string                       `ncpdp:"07CW" json:"alternateID" jsonWithId:"330_CW_alternateID"`
	ScheduledRxID                         *string                       `ncpdp:"07EK" json:"scheduledRxID" jsonWithId:"454_EK_scheduledRxID"`
	UnitOfMeasure                         *string                       `ncpdp:"0728" json:"unitOfMeasure" jsonWithId:"600_28_unitOfMeasure"`
	LevelOfService                        *string                       `ncpdp:"07DI" json:"levelOfService" jsonWithId:"418_DI_levelOfService"`
	PriorAuthorizationTypeCode            *int                          `ncpdp:"07EU" json:"priorAuthorizationTypeCode" jsonWithId:"461_EU_priorAuthorizationTypeCode"`
	PriorAuthorizationNumberSubmitted     *string                       `ncpdp:"07EV" json:"priorAuthorizationNumberSubmitted" jsonWithId:"462_EV_priorAuthorizationNumberSubmitted"`
	IntermediaryAuthorizationTypeID       *string                       `ncpdp:"07EW" json:"intermediaryAuthorizationTypeID" jsonWithId:"463_EW_intermediaryAuthorizionTypeID"`
	IntermediaryAuthorizationID           *string                       `ncpdp:"07EX" json:"intermediaryAuthorizationID" jsonWithId:"464_EX_intermediaryAuthorizationID"`
	DispensingStatus                      *string                       `ncpdp:"07HD" json:"dispensingStatus" jsonWithId:"343_HD_dispensingStatus"`
	QuantityIntendedToBeDispersed         *scriptparser.UnsignedNumeric `ncpdp:"07HF" json:"quantityIntendedToBeDispersed" jsonWithId:"344_HF_quantityIntendedToBeDispersed"`
	DaysSupplyIntendedToBeDispersed       *int                          `ncpdp:"07HG" json:"daysSupplyIntendedToBeDispersed" jsonWithId:"345_HG_daysSupplyIntendedToBeDispersed"`
	DelayReasonCode                       *string                       `ncpdp:"07NV" json:"delayReasonCode" jsonWithId:"357_NV_delayReasonCode"`
	TransactionReferenceNumber            *string                       `ncpdp:"07K5" json:"transactionReferenceNumber" jsonWithId:"880_K5_transactionReferenceNumber"`
	PatientAssignmentIndicator            *string                       `ncpdp:"07MT" json:"patientAssignmentIndicator" jsonWithId:"391_MT_patientAssignmentIndicator"`
	RouteOfAdministration                 *string                       `ncpdp:"07E2" json:"routeOfAdministration" jsonWithId:"995_E2_routeOfAdministration"`
	CompoundType                          *string                       `ncpdp:"07G1" json:"compoundType" jsonWithId:"996_G1_compoundType"`
	TransactionCtrlNumber                 *string                       `ncpdp:"07N4" json:"transactionCtrlNumber" jsonWithId:"114_N4_transactionCtrlNumber"`
	PharmacyServiceType                   *string                       `ncpdp:"07U7" json:"pharmacyServiceType" jsonWithId:"147_U7_pharmacyServiceType"`
}

type Pricing struct {
	SegmentID                            *string                    `ncpdp:"11AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	IngredientCostSubmitted              *scriptparser.Currency64   `ncpdp:"11D9" json:"ingredientCostSubmitted" jsonWithId:"409_D9_ingredientCostSubmitted"`
	DispensingFeeSubmitted               *scriptparser.Currency64   `ncpdp:"11DC" json:"dispensingFeeSubmitted" jsonWithId:"412_DC_dispensingFeeSubmitted"`
	ProfessionServiceFeeSubmitted        *scriptparser.Currency64   `ncpdp:"11BE" json:"professionServiceFeeSubmitted" jsonWithId:"477_BE_professionServiceFeeSubmitted"`
	PatientPaidAmount                    *scriptparser.Currency64   `ncpdp:"11DX" json:"patientPaidAmount" jsonWithId:"433_DX_patientPaidAmount"`
	IncentiveAmountSubmitted             *scriptparser.Currency64   `ncpdp:"11E3" json:"incentiveAmountSubmitted" jsonWithId:"438_E3_incentiveAmountSubmitted"`
	OtherAmountClaimedSubmittedCount     *int64                     `ncpdp:"11H7" json:"otherAmountClaimedSubmittedCount" jsonWithId:"478_H7_otherAmountClaimedSubmittedCount"`
	OtherAmountClaimedSubmittedQualifier []*string                  `ncpdp:"11H8" json:"otherAmountClaimedSubmittedQualifier" jsonWithId:"479_H8_otherAmountClaimedSubmittedQualifier"`
	OtherAmountClaimedSubmitted          []*scriptparser.Currency64 `ncpdp:"11H9" json:"otherAmountClaimedSubmitted" jsonWithId:"480_H9_otherAmountClaimedSubmitted"`
	FlatSalesTaxAmountSubmitted          *scriptparser.Currency64   `ncpdp:"11HA" json:"flatSalesTaxAmountSubmitted" jsonWithId:"481_HA_flatSalesTaxAmountSubmitted"`
	PercentageSalesTaxAmountSubmitted    *scriptparser.Currency64   `ncpdp:"11GE" json:"percentageSalesTaxAmountSubmitted" jsonWithId:"482_GE_percentageSalesTaxAmountSubmitted"`
	PercentageSalesTaxRateSubmitted      *scriptparser.Currency64   `ncpdp:"11HE" json:"percentageSalesTaxRateSubmitted" jsonWithId:"484_HE_percentageSalesTaxRateSubmitted"`
	PercentageSalesTaxBasisSubmitted     *string                    `ncpdp:"11JE" json:"percentageSalesTaxBasisSubmitted" jsonWithId:"483_JE_percentageSalesTaxBasisSubmitted"`
	UsualAndCustomaryCharge              *scriptparser.Currency64   `ncpdp:"11DQ" json:"usualAndCustomaryCharge" jsonWithId:"426_DQ_usualAndCustomaryCharge"`
	GrossAmountDue                       *scriptparser.Currency64   `ncpdp:"11DU" json:"grossAmountDue" jsonWithId:"430_DU_grossAmountDue"`
	BasisOfCostDetermination             *string                    `ncpdp:"11DN" json:"basisOfCostDetermination" jsonWithId:"423_DN_basisOfCostDetermination"`
	MedicaidPaidAmount                   *scriptparser.Currency64   `ncpdp:"11N3" json:"medicaidPaidAmount" jsonWithId:"113_N3_medicaidPaidAmount"`
}

// pharmacy provider segment
type PharmacyProvider struct {
	SegmentID           *string `ncpdp:"02AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	ProviderIdQualifier *string `ncpdp:"01EY" json:"providerIdQualifier" jsonWithId:"465_EY_providerIdQualifier"`
	ProviderId          *string `ncpdp:"01E9" json:"providerId" jsonWithId:"444_E9_providerId"`
}

// Prescriber segment
type Prescriber struct {
	SegmentID                    *string `ncpdp:"03AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	PrescriberIdQualifier        *string `ncpdp:"03EZ" json:"prescriberIdQualifier" jsonWithId:"466_EZ_prescriberIdQualifier"`
	PrescriberId                 *string `ncpdp:"03DB" json:"prescriberId" jsonWithId:"411_DB_prescriberId"`
	PrescriberLastname           *string `ncpdp:"03DR" json:"prescriberLastname" jsonWithId:"427_DR_prescriberLastname"`
	PrescriberPhoneNumber        *string `ncpdp:"03PM" json:"prescriberPhoneNumber" jsonWithId:"498_PM_prescriberPhoneNumber"`
	PrimaryCareProviderQualifier *string `ncpdp:"032E" json:"primaryCareProviderQualifier" jsonWithId:"468_2E_primaryCareProviderQualifier"`
	PrimaryCareProvider          *string `ncpdp:"03DL" json:"primaryCareProvider" jsonWithId:"421_DL_primaryCareProvider"`
	PrimaryCareProviderLastname  *string `ncpdp:"034E" json:"primaryCareProviderLastname" jsonWithId:"470_4E_primaryCareProviderLastname"`
	PrescriberFirstname          *string `ncpdp:"032J" json:"prescriberFirstname" jsonWithId:"364_2J_prescriberFirstname"`
	PrescriberStreetAddress      *string `ncpdp:"032k" json:"prescriberStreetAddress" jsonWithId:"365_2k_prescriberStreetAddress"`
	PrescriberCity               *string `ncpdp:"032M" json:"prescriberCity" jsonWithId:"366_2M_prescriberCity"`
	PrescriberState              *string `ncpdp:"032N" json:"prescriberState" jsonWithId:"367_2N_prescriberState"`
	PrescriberZip                *string `ncpdp:"032P" json:"prescriberZip" jsonWithId:"368_2P_prescriberZip"`
}

// CoordinationOfBenefitsOtherPayments segment
type COB struct {
	SegmentID                                *string                    `ncpdp:"05AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	OtherPaymentsCount                       *int                       `ncpdp:"054C" json:"otherPaymentsCount" jsonWithId:"337_4C_otherPaymentsCount"`
	OtherPayerCoverageType                   *[]string                  `ncpdp:"054C" json:"otherPayerCoverageType" jsonWithId:"338_4C_otherPayerCoverageType"`
	OtherPayerIdQualifier                    *[]string                  `ncpdp:"056C" json:"otherPayerIdQualifier" jsonWithId:"339_6C_otherPayerIdQualifier"`
	OtherPayerId                             *[]string                  `ncpdp:"057C" json:"otherPayerId" jsonWithId:"340_7C_otherPayerId"`
	OtherPayerDate                           *[]string                  `ncpdp:"05E8" json:"otherPayerDate" jsonWithId:"443_E8_otherPayerDate"`
	InternalControlNumber                    *[]string                  `ncpdp:"05A7" json:"internalControlNumber" jsonWithId:"993_A7_internalControlNumber"`
	OtherPayerAmountPaidCount                *int                       `ncpdp:"05HB" json:"otherPayerAmountPaidCount" jsonWithId:"341_HB_otherPayerAmountPaidCount"`
	OtherPayerAmountPaidQualifier            *[]string                  `ncpdp:"05HC" json:"otherPayerAmountPaidQualifier" jsonWithId:"342_HC_otherPayerAmmountPaidQualifier"`
	OtherPayerAmountPaid                     *[]scriptparser.Currency64 `ncpdp:"05DV" json:"otherPayerAmountPaid" jsonWithId:"431_DV_otherPayerAmountPaid"`
	OtherPayerRejectCount                    *int                       `ncpdp:"055E" json:"otherPayerRejectCount" jsonWithId:"471_5E_otherPayerRejectCount"`
	OtherPayerRejectCode                     *[]string                  `ncpdp:"056E" json:"otherPayerRejectCode" jsonWithId:"472_6E_otherPayerRejectCode"`
	OtherPayerPatientResponsibilityCount     *int                       `ncpdp:"05NR" json:"otherPayerPatientResponsibilityCount" jsonWithId:"353_NR_otherPayerPatientResponsibilityCount"`
	OtherPayerPatientResponsibilityQualifier *[]string                  `ncpdp:"05NP" json:"otherPayerPatientResponsibilityQualifier" jsonWithId:"351_NP_otherPayerPatientResponsibilityQualifier"`
	OtherPayerPatientResponsibility          *[]scriptparser.Currency64 `ncpdp:"05NQ" json:"otherPayerPatientResponsibility" jsonWithId:"352_NQ_otherPayerPatientResponsibility"`
	BenefitStageCount                        *int                       `ncpdp:"05MU" json:"benefitStageCount" jsonWithId:"392_MU_benefitStageCount"`
	BenefitStageQualifier                    *[]string                  `ncpdp:"05MV" json:"benefitStageQualifier" jsonWithId:"393_MV_benefitStageQualifier"`
	BenefitStageAmount                       *[]scriptparser.Currency64 `ncpdp:"05MW" json:"benefitStageAmount" jsonWithId:"394_MW_benefitStageAmount"`
}

type DUR struct {
	SegmentID               *string   `ncpdp:"08AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	ResponseCodeCounter     *int      `ncpdp:"087C" json:"responseCodeCounter" jsonWithId:"473_7E_responseCodeCounter"`
	ReasonForServiceCode    *[]string `ncpdp:"08E4" json:"reasonForServiceCode" jsonWithId:"439_E4_reasonForServiceCode"`
	ProfessionalServiceCode *[]string `ncpdp:"08E5" json:"professionalServiceCode" jsonWithId:"440_E5_professionalServiceCode"`
	ResultOfServiceCode     *[]string `ncpdp:"08E6" json:"resultOfServiceCode" jsonWithId:"441_E6_resultOfServiceCode"`
	LevelOfEffort           *[]string `ncpdp:"088E" json:"levelOfEffortCode" jsonWithId:"474_8E_levelOfEffortCode"`
	CoAgentIDQualifier      *[]string `ncpdp:"08E7" json:"coAgentIDQualifier" jsonWithId:"442_E7_coAgentIDQualifier"`
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
