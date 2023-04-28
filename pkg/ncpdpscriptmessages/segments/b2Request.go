package segments

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

type B2RequestHeader struct {
	VersionRel                 *string `ncpdp:"header:versionRel" json:"versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode" json:"transactionCode"`
	TransactionCount           *int    `ncpdp:"header:transactionCount" json:"transactionCount"`
	TransactionResponseStatus  *string `ncpdp:"header:transactionResponseStatus" json:"transactionResponseStatus"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier" json:"serviceProviderIdQualifier"`
	ServiceProviderId          *string `ncpdp:"header:serviceProviderId" json:"serviceProviderId"`
	Dos                        *string `ncpdp:"header:dos" json:"dos"`
}
type B2RequestInsurance struct {
	SegmentId                          *string `ncpdp:"04AM" json:"segmentId" jsonWithId:"111_AM_segmentId"`
	CardHolderId                       *string `ncpdp:"04C2" json:"cardHolderId" jsonWithId:"302_C2_cardHolderId"`
	CardHolderFirstName                *string `ncpdp:"04CC" json:"cardHolderFirstName" jsonWithId:"312_CC_cardHolderFirstName"`
	CardHolderLastName                 *string `ncpdp:"04CD" json:"cardHolderLastName" jsonWithId:"313_CD_cardHolderLastName"`
	HomePlan                           *string `ncpdp:"04CE" json:"homePlan" jsonWithId:"314_CE_homePlan"`
	PlanId                             *string `ncpdp:"04FO" json:"planId" jsonWithId:"524_FO_planId"`
	EligibilityClarificationCode       *string `ncpdp:"04C9" json:"eligibilityClarificationCode" jsonWithId:"309_C9_eligibilityClarificationCode"`
	GroupID                            *string `ncpdp:"04C1" json:"groupID" jsonWithId:"301_C1_groupID"`
	PersonCode                         *string `ncpdp:"04C3" json:"personCode" jsonWithId:"303_C3_personCode"`
	PatientRelationshipCode            *string `ncpdp:"04C6" json:"patientRelationshipCode" jsonWithId:"306_C6_patientRelationshipCode"`
	OtherPayerBinNumber                *string `ncpdp:"04MG" json:"otherPayerBinNumber" jsonWithId:"990_MG_otherPayerBinNumber"`
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

type B2RequestClaim struct {
	SegmentID                                     *string                       `ncpdp:"07AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	RxServiceRefNumberQualifier                   *string                       `ncpdp:"07EM" json:"rxServiceRefNumberQualifier" jsonWithId:"455_EM_rxServiceRefNumberQualifier"`
	RxServiceRefNumber                            *string                       `ncpdp:"07D2" json:"rxServiceReferenceNumber" jsonWithId:"402_D2_rxServiceReferenceNumber"`
	ProductServiceIDQualifier                     *string                       `ncpdp:"07E1" json:"productServiceIDQualifier" jsonWithId:"436_E1_productServiceIDQualifier"`
	ProductServiceID                              *string                       `ncpdp:"07D7" json:"productServiceID" jsonWithId:"407_D7_productServiceID"`
	AssociatedRxServiceRefNumber                  *string                       `ncpdp:"07EN" json:"associatedRxServiceRefNumber" jsonWithId:"456_EN_associatedRxServiceRefNumber"`
	AssociatedRxServiceDate                       *string                       `ncpdp:"07EP" json:"associatedRxServiceDate" jsonWithId:"457_EP_associatedRxServiceDate"`
	ProcedureModifierCodeCount                    *int                          `ncpdp:"07SE" json:"procedureModifierCodeCount" jsonWithId:"458_SE_procedureModifierCodeCount"`
	ProcedureModifierCode                         []*string                     `ncpdp:"07ER" json:"procedureModifierCode" jsonWithId:"459_ER_procedureModifierCode"`
	QuantityDispensed                             *scriptparser.UnsignedNumeric `ncpdp:"07E7" json:"quantityDispensed" jsonWithId:"442_E7_quantityDispensed"`
	FillNumber                                    *int                          `ncpdp:"07D3" json:"fillNumber" jsonWithId:"403_D3_fillNumber"`
	DaysSupply                                    *int                          `ncpdp:"07D5" json:"daysSupply" jsonWithId:"405_D5_daysSupply"`
	CompoundCode                                  *int                          `ncpdp:"07D6" json:"compoundCode" jsonWithId:"406_D6_compoundCode"`
	DispenseAsWritten                             *int                          `ncpdp:"07D8" json:"dispenseAsWritten" jsonWithId:"408_D8_dispenseAsWritten"`
	DateRxWritten                                 *string                       `ncpdp:"07DE" json:"dateRxWritten" jsonWithId:"414_DE_dateRxWritten"`
	NumberOfRefillsAuthorized                     *int                          `ncpdp:"07DF" json:"numberOfRefillsAuthorized" jsonWithId:"415_DF_numberOfRefillsAuthorized"`
	RxOriginCode                                  *int                          `ncpdp:"07DJ" json:"rxOriginCode" jsonWithId:"419_DJ_rxOriginCode"`
	SubmissionClarificationCodeCount              *int                          `ncpdp:"07NX" json:"submissionClarificationCodeCount" jsonWithId:"354_NX_submissionClarificationCodeCount"`
	SubmissionClarificationCode                   []*int                        `ncpdp:"07DK" json:"submissionClarificationCode" jsonWithId:"420_DK_submissionClarificationCode"`
	QuantityPrescriber                            *scriptparser.UnsignedNumeric `ncpdp:"07ET" json:"quantityPrescriber" jsonWithId:"460_ET_quantityPrescriber"`
	OtherCoverageCode                             *int                          `ncpdp:"07C8" json:"otherCoverageCode" jsonWithId:"308_C8_otherCoverageCode"`
	SpecialPackageIndicator                       *int                          `ncpdp:"07DT" json:"specialPackageIndicator" jsonWithId:"429_DT_specialPackageIndicator"`
	OriginallyPrescribedProductServiceIdQualifier *int                          `ncpdp:"07EJ" json:"originallyPrescribedProductServiceIdQualifier" jsonWithId:"453_EJ_originallyPrescribedProductServiceIdQualifier"`
	OriginallyPrescribedProductServiceCode        *string                       `ncpdp:"07EA" json:"originallyPrescribedProductServiceCode" jsonWithId:"445_EA_originallyPrescribedProductServiceCode"`
	OriginallyPrescribedQuantity                  *scriptparser.UnsignedNumeric `ncpdp:"07EB" json:"originallyPrescribedQuantity" jsonWithId:"446_EB_originallyPrescribedQuantity"`
	AlternateID                                   *string                       `ncpdp:"07CW" json:"alternateID" jsonWithId:"330_CW_alternateID"`
	ScheduledRxIdNumber                           *string                       `ncpdp:"07EK" json:"scheduledRxIdNumber" jsonWithId:"454_EK_scheduledRxIdNumber"`
	UnitOfMeasure                                 *string                       `ncpdp:"0728" json:"unitOfMeasure" jsonWithId:"600_28_unitOfMeasure"`
	LevelOfService                                *string                       `ncpdp:"07DI" json:"levelOfService" jsonWithId:"418_DI_levelOfService"`
	PriorAuthorizationTypeCode                    *int                          `ncpdp:"07EU" json:"priorAuthorizationTypeCode" jsonWithId:"461_EU_priorAuthorizationTypeCode"`
	PriorAuthorizationNumberSubmitted             *string                       `ncpdp:"07EV" json:"priorAuthorizationNumberSubmitted" jsonWithId:"462_EV_priorAuthorizationNumberSubmitted"`
	IntermediaryAuthorizationTypeID               *string                       `ncpdp:"07EW" json:"intermediaryAuthorizationTypeID" jsonWithId:"463_EW_intermediaryAuthorizionTypeID"`
	IntermediaryAuthorizationID                   *string                       `ncpdp:"07EX" json:"intermediaryAuthorizationID" jsonWithId:"464_EX_intermediaryAuthorizationID"`
	DispensingStatus                              *string                       `ncpdp:"07HD" json:"dispensingStatus" jsonWithId:"343_HD_dispensingStatus"`
	QuantityIntendedToBeDispensed                 *scriptparser.UnsignedNumeric `ncpdp:"07HF" json:"quantityIntendedToBeDispensed" jsonWithId:"344_HF_quantityIntendedToBeDispensed"`
	DaysSupplyIntendedToBeDispensed               *int                          `ncpdp:"07HG" json:"daysSupplyIntendedToBeDispensed" jsonWithId:"345_HG_daysSupplyIntendedToBeDispensed"`
	DelayReasonCode                               *string                       `ncpdp:"07NV" json:"delayReasonCode" jsonWithId:"357_NV_delayReasonCode"`
	TransactionReferenceNumber                    *string                       `ncpdp:"07K5" json:"transactionReferenceNumber" jsonWithId:"880_K5_transactionReferenceNumber"`
	PatientAssignmentIndicator                    *string                       `ncpdp:"07MT" json:"patientAssignmentIndicator" jsonWithId:"391_MT_patientAssignmentIndicator"`
	RouteOfAdministration                         *string                       `ncpdp:"07E2" json:"routeOfAdministration" jsonWithId:"995_E2_routeOfAdministration"`
	CompoundType                                  *string                       `ncpdp:"07G1" json:"compoundType" jsonWithId:"996_G1_compoundType"`
	TransactionCtrlNumber                         *string                       `ncpdp:"07N4" json:"transactionCtrlNumber" jsonWithId:"114_N4_transactionCtrlNumber"`
	PharmacyServiceType                           *string                       `ncpdp:"07U7" json:"pharmacyServiceType" jsonWithId:"147_U7_pharmacyServiceType"`
}

type B2RequestDUR struct {
	SegmentID               *string   `ncpdp:"08AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	DURPPSCodeCounter       *int      `ncpdp:"087C" json:"DURPPSCodeCounter" jsonWithId:"473_7E_DURPPSCodeCounter"`
	ReasonForServiceCode    *[]string `ncpdp:"08E4" json:"reasonForServiceCode" jsonWithId:"439_E4_reasonForServiceCode"`
	ProfessionalServiceCode *[]string `ncpdp:"08E5" json:"professionalServiceCode" jsonWithId:"440_E5_professionalServiceCode"`
	ResultOfServiceCode     *[]string `ncpdp:"08E6" json:"resultOfServiceCode" jsonWithId:"441_E6_resultOfServiceCode"`
	DURPPSLevelOfEffortCode *[]string `ncpdp:"088E" json:"DURPPSLevelOfEffortCode" jsonWithId:"474_8E_DURPPSLevelOfEffortCode"`
	DURCoAgentIDQualifier   *[]string `ncpdp:"08J9" json:"DURCoAgentIDQualifier" jsonWithId:"475_J9_DURCoAgentIDQualifier"`
	DURCoAgentID            *[]string `ncpdp:"08H6" json:"DURCoAgentID" jsonWithId:"476_H6_DURCoAgentID"`
}
type B2RequestPricing struct {
	SegmentID                            *string                    `ncpdp:"11AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	IngredientCostSubmitted              *scriptparser.Currency64   `ncpdp:"11D9" json:"ingredientCostSubmitted" jsonWithId:"409_D9_ingredientCostSubmitted"`
	DispensingFeeSubmitted               *scriptparser.Currency64   `ncpdp:"11DC" json:"dispensingFeeSubmitted" jsonWithId:"412_DC_dispensingFeeSubmitted"`
	ProfessionServiceFeeSubmitted        *scriptparser.Currency64   `ncpdp:"11BE" json:"professionServiceFeeSubmitted" jsonWithId:"477_BE_professionServiceFeeSubmitted"`
	PatientPaidAmountSubmitted           *scriptparser.Currency64   `ncpdp:"11DX" json:"patientPaidAmountSubmitted" jsonWithId:"433_DX_patientPaidAmountSubmitted"`
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
type B2RequestCOB struct {
	SegmentID                                      *string                    `ncpdp:"05AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	OtherPaymentsCount                             *int                       `ncpdp:"054C" json:"otherPaymentsCount" jsonWithId:"337_4C_otherPaymentsCount"`
	OtherPayerCoverageType                         *[]string                  `ncpdp:"054C" json:"otherPayerCoverageType" jsonWithId:"338_4C_otherPayerCoverageType"`
	OtherPayerIdQualifier                          *[]string                  `ncpdp:"056C" json:"otherPayerIdQualifier" jsonWithId:"339_6C_otherPayerIdQualifier"`
	OtherPayerId                                   *[]string                  `ncpdp:"057C" json:"otherPayerId" jsonWithId:"340_7C_otherPayerId"`
	OtherPayerDate                                 *[]string                  `ncpdp:"05E8" json:"otherPayerDate" jsonWithId:"443_E8_otherPayerDate"`
	InternalControlNumber                          *[]string                  `ncpdp:"05A7" json:"internalControlNumber" jsonWithId:"993_A7_internalControlNumber"`
	OtherPayerAmountPaidCount                      *int                       `ncpdp:"05HB" json:"otherPayerAmountPaidCount" jsonWithId:"341_HB_otherPayerAmountPaidCount"`
	OtherPayerAmountPaidQualifier                  *[]string                  `ncpdp:"05HC" json:"otherPayerAmountPaidQualifier" jsonWithId:"342_HC_otherPayerAmmountPaidQualifier"`
	OtherPayerAmountPaid                           *[]scriptparser.Currency64 `ncpdp:"05DV" json:"otherPayerAmountPaid" jsonWithId:"431_DV_otherPayerAmountPaid"`
	OtherPayerRejectCount                          *int                       `ncpdp:"055E" json:"otherPayerRejectCount" jsonWithId:"471_5E_otherPayerRejectCount"`
	OtherPayerRejectCode                           *[]string                  `ncpdp:"056E" json:"otherPayerRejectCode" jsonWithId:"472_6E_otherPayerRejectCode"`
	OtherPayerPatientResponsibilityAmountCount     *int                       `ncpdp:"05NR" json:"otherPayerPatientResponsibilityAmountCount" jsonWithId:"353_NR_otherPayerPatientResponsibilityAmountCount"`
	OtherPayerPatientResponsibilityAmountQualifier *[]string                  `ncpdp:"05NP" json:"otherPayerPatientResponsibilityAmountQualifier" jsonWithId:"351_NP_otherPayerPatientResponsibilityAmountQualifier"`
	OtherPayerPatientResponsibilityAmount          *[]scriptparser.Currency64 `ncpdp:"05NQ" json:"otherPayerPatientResponsibilityAmount" jsonWithId:"352_NQ_otherPayerPatientResponsibilityAmount"`
	BenefitStageCount                              *int                       `ncpdp:"05MU" json:"benefitStageCount" jsonWithId:"392_MU_benefitStageCount"`
	BenefitStageQualifier                          *[]string                  `ncpdp:"05MV" json:"benefitStageQualifier" jsonWithId:"393_MV_benefitStageQualifier"`
	BenefitStageAmount                             *[]scriptparser.Currency64 `ncpdp:"05MW" json:"benefitStageAmount" jsonWithId:"394_MW_benefitStageAmount"`
}
