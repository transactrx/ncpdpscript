package segments

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

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

type PharmacyProvider struct {
	SegmentID           *string `ncpdp:"02AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	ProviderIdQualifier *string `ncpdp:"01EY" json:"providerIdQualifier" jsonWithId:"465_EY_providerIdQualifier"`
	ProviderId          *string `ncpdp:"01E9" json:"providerId" jsonWithId:"444_E9_providerId"`
}

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

type WorkersCompensation struct {
	SegmentID                           *string `ncpdp:"06AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	DateOfInjury                        *string `ncpdp:"06DY" json:"dateOfInjury" jsonWithId:"434_DY_dateOfInjury"`
	EmployerName                        *string `ncpdp:"06CF" json:"employerName" jsonWithId:"315_CF_employerName"`
	EmployerStreetAddress               *string `ncpdp:"06CG" json:"employerStreetAddress" jsonWithId:"316_CG_employerStreetAddress"`
	EmployerCityAddress                 *string `ncpdp:"06CH" json:"employerCityAddress" jsonWithId:"317_CH_employerCityAddress"`
	EmployerStateProvinceAddress        *string `ncpdp:"06CI" json:"employerStateProvinceAddress" jsonWithId:"318_CI_employerStateProvinceAddress"`
	EmployerZipPostalZone               *string `ncpdp:"06CJ" json:"employerZipPostalZone" jsonWithId:"319_CJ_employerZipPostalZone"`
	EmployerPhoneNumber                 *string `ncpdp:"06CK" json:"employerPhoneNumber" jsonWithId:"320_CK_employerPhoneNumber"`
	EmployerContactName                 *string `ncpdp:"06CL" json:"employerContactName" jsonWithId:"321_CL_employerContactName"`
	CarrierID                           *string `ncpdp:"06CR" json:"carrierID" jsonWithId:"327_CR_carrierID"`
	ClaimReferenceID                    *string `ncpdp:"06DZ" json:"claimReferenceID" jsonWithId:"435_DZ_claimReferenceID"`
	BillingEntityTypeIndicator          *string `ncpdp:"06TR" json:"billingEntityTypeIndicator" jsonWithId:"117_TR_billingEntityTypeIndicator"`
	PayToQualifier                      *string `ncpdp:"06TS" json:"payToQualifier" jsonWithId:"118_TS_payToQualifier"`
	PayToID                             *string `ncpdp:"06TT" json:"payToID" jsonWithId:"119_TT_payToID"`
	PayToName                           *string `ncpdp:"06TU" json:"payToName" jsonWithId:"120_TU_payToName"`
	PayToStreetAddress                  *string `ncpdp:"06TV" json:"payToStreetAddress" jsonWithId:"121_TV_payToStreetAddress"`
	PayToCityAddress                    *string `ncpdp:"06TW" json:"payToCityAddress" jsonWithId:"122_TW_payToCityAddress"`
	PayToStateProvinceAddress           *string `ncpdp:"06TX" json:"payToStateProvinceAddress" jsonWithId:"123_TX_payToStateProvinceAddress"`
	PayToZipPostalZone                  *string `ncpdp:"06TY" json:"payToZipPostalZone" jsonWithId:"124_TY_payToZipPostalZone"`
	GenericEquivalentProductIDQualifier *string `ncpdp:"06TZ" json:"genericEquivalentProductIDQualifier" jsonWithId:"125_TZ_genericEquivalentProductIDQualifier"`
	GenericEquivalentProductID          *string `ncpdp:"06UA" json:"genericEquivalentProductID" jsonWithId:"126_UA_genericEquivalentProductID"`
}

type DUR struct {
	SegmentID               *string   `ncpdp:"08AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	ResponseCodeCounter     *int      `ncpdp:"087C" json:"responseCodeCounter" jsonWithId:"473_7E_responseCodeCounter"`
	ReasonForServiceCode    *[]string `ncpdp:"08E4" json:"reasonForServiceCode" jsonWithId:"439_E4_reasonForServiceCode"`
	ProfessionalServiceCode *[]string `ncpdp:"08E5" json:"professionalServiceCode" jsonWithId:"440_E5_professionalServiceCode"`
	ResultOfServiceCode     *[]string `ncpdp:"08E6" json:"resultOfServiceCode" jsonWithId:"441_E6_resultOfServiceCode"`
	LevelOfEffort           *[]string `ncpdp:"088E" json:"levelOfEffortCode" jsonWithId:"474_8E_levelOfEffortCode"`
	CoAgentIDQualifier      *[]string `ncpdp:"08J9" json:"coAgentIDQualifier" jsonWithId:"475_J9_coAgentIDQualifier"`
	CoAgentID               *[]string `ncpdp:"08H6" json:"coAgentID" jsonWithId:"476_H6_coAgentID"`
}

type Coupon struct {
	SegmentID         *string `ncpdp:"09AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	CouponType        *string `ncpdp:"09KE" json:"couponType" jsonWithId:"485_KE_couponType"`
	CouponNumber      *string `ncpdp:"09ME" json:"couponNumber" jsonWithId:"486_ME_couponNumber"`
	CouponValueAmount *string `ncpdp:"09NE" json:"couponValueAmount" jsonWithId:"487_NE_couponValueAmount"`
}
type Compound struct {
	SegmentID                                  *string   `ncpdp:"10AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	CompoundDosageFormDescriptionCode          *string   `ncpdp:"104C" json:"compoundDosageFormDescriptionCode" jsonWithId:"407_4C_compoundDosageFormDescriptionCode"`
	CompoundDispensingUnitFormIndicator        *string   `ncpdp:"10U7" json:"compoundDispensingUnitFormIndicator" jsonWithId:"456_U7_compoundDispensingUnitFormIndicator"`
	CompoundIngredientComponentCount           *int      `ncpdp:"10EC" json:"compoundIngredientComponentCount" jsonWithId:"447_EC_compoundIngredientComponentCount"`
	CompoundProductIDQualifier                 *[]string `ncpdp:"10RE" json:"compoundProductIDQualifier" jsonWithId:"488_RE_compoundProductIDQualifier"`
	CompoundProductID                          *[]string `ncpdp:"10TE" json:"compoundProductID" jsonWithId:"489_TE_compoundProductID"`
	CompoundIngredientQuantity                 *[]string `ncpdp:"10ED" json:"compoundIngredientQuantity" jsonWithId:"448_ED_compoundIngredientQuantity"`
	CompoundIngredientDrugCost                 *[]string `ncpdp:"10EE" json:"compoundIngredientDrugCost" jsonWithId:"449_EE_compoundIngredientDrugCost"`
	CompoundIngredientBasisOfCostDetermination *[]string `ncpdp:"10UE" json:"compoundIngredientBasisOfCostDetermination" jsonWithId:"490_UE_compoundIngredientBasisOfCostDetermination"`
	CompoundIngredientModifierCodeCount        *int      `ncpdp:"102G" json:"compoundIngredientModifierCodeCount" jsonWithId:"362_2G_compoundIngredientModifierCodeCount"`
	CompoundIngredientModifierCode             *[]string `ncpdp:"102H" json:"compoundIngredientModifierCode" jsonWithId:"363_2H_compoundIngredientModifierCode"`
}
type Clinical struct {
	SegmentID                  *string   `ncpdp:"13AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	DiagnosisCodeCount         *int      `ncpdp:"13VE" json:"diagnosisCodeCount" jsonWithId:"491_VE_diagnosisCodeCount"`
	DiagnosisCodeQualifier     *[]string `ncpdp:"13WE" json:"diagnosisCodeQualifier" jsonWithId:"492_WE_diagnosisCodeQualifier"`
	DiagnosisCode              *[]string `ncpdp:"13DO" json:"diagnosisCode" jsonWithId:"424_DO_diagnosisCode"`
	ClinicalInformationCounter *int      `ncpdp:"13XE" json:"clinicalInformationCounter" jsonWithId:"493_XE_clinicalInformationCounter"`
	MeasurementDate            *[]string `ncpdp:"13ZE" json:"measurementDate" jsonWithId:"494_ZE_measurementDate"`
	MeasurementTime            *[]string `ncpdp:"13H1" json:"measurementTime" jsonWithId:"495_H1_measurementTime"`
	MeasurementDimension       *[]string `ncpdp:"13H2" json:"measurementDimension" jsonWithId:"496_H2_measurementDimension"`
	MeasurementUnit            *[]string `ncpdp:"13H3" json:"measurementUnit" jsonWithId:"497_H3_measurementUnit"`
	MeasurementValue           *[]string `ncpdp:"13H4" json:"measurementValue" jsonWithId:"499_H4_measurementValue"`
}
type AdditionalDocumentation struct {
	SegmentID                      *string   `ncpdp:"14AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	AdditionalDocumentationTypeID  *string   `ncpdp:"142Q" json:"additionalDocumentationTypeID" jsonWithId:"369_2Q_additionalDocumentationTypeID"`
	RequestPeriodBeginDate         *string   `ncpdp:"142V" json:"requestPeriodBeginDate" jsonWithId:"374_2V_requestPeriodBeginDate"`
	RequestStatus                  *string   `ncpdp:"142U" json:"requestStatus" jsonWithId:"373_2U_requestStatus"`
	RequestPeriodRecertRevisedDate *string   `ncpdp:"142W" json:"requestPeriodRecertRevisedDate" jsonWithId:"375_2W_requestPeriodRecertRevisedDate"`
	LengthOfNeedQualifier          *string   `ncpdp:"142S" json:"lengthOfNeedQualifier" jsonWithId:"371_2S_lengthOfNeedQualifier"`
	LengthOfNeed                   *string   `ncpdp:"142R" json:"lengthOfNeed" jsonWithId:"370_2R_lengthOfNeed"`
	PrescriberSupplierDateSigned   *string   `ncpdp:"142T" json:"prescriberSupplierDateSigned" jsonWithId:"372_2T_prescriberSupplierDateSigned"`
	SupportingDocumentation        *string   `ncpdp:"142X" json:"supportingDocumentation" jsonWithId:"376_2X_supportingDocumentation"`
	QuestionNumberLetterCount      *int      `ncpdp:"142Z" json:"questionNumberLetterCount" jsonWithId:"377_2Z_questionNumberLetterCount"`
	QuestionNumberLetter           *[]string `ncpdp:"144B" json:"questionNumberLetter" jsonWithId:"378_4B_questionNumberLetter"`
	QuestionPercentResponse        *[]string `ncpdp:"144D" json:"questionPercentResponse" jsonWithId:"379_4D_questionPercentResponse"`
	QuestionDateResponse           *[]string `ncpdp:"144G" json:"questionDateResponse" jsonWithId:"380_4G_questionDateResponse"`
	QuestionDollarAmountResponse   *[]string `ncpdp:"144H" json:"questionDollarAmountResponse" jsonWithId:"381_4H_questionDollarAmountResponse"`
	QuestionNumericResponse        *[]string `ncpdp:"144J" json:"questionNumericResponse" jsonWithId:"382_4J_questionNumericResponse"`
	QuestionAlphanumericResponse   *[]string `ncpdp:"144K" json:"questionAlphanumericResponse" jsonWithId:"383_4K_questionAlphanumericResponse"`
}
type Facility struct {
	SegmentID                    *string `ncpdp:"15AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	FacilityID                   *string `ncpdp:"158C" json:"facilityID" jsonWithId:"336_8C_facilityID"`
	FacilityName                 *string `ncpdp:"153Q" json:"facilityName" jsonWithId:"385_3Q_facilityName"`
	FacilityStreetAddress        *string `ncpdp:"153U" json:"facilityStreetAddress" jsonWithId:"386_3U_facilityStreetAddress"`
	FacilityCityAddress          *string `ncpdp:"155J" json:"facilityCityAddress" jsonWithId:"388_5J_facilityCityAddress"`
	FacilityStateProvinceAddress *string `ncpdp:"153V" json:"facilityStateProvinceAddress" jsonWithId:"387_3V_facilityStateProvinceAddress"`
	FacilityZipPostalZone        *string `ncpdp:"156D" json:"facilityZipPostalZone" jsonWithId:"389_6D_facilityZipPostalZone"`
}

type Narrative struct {
	SegmentID        *string `ncpdp:"16AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	NarrativeMessage *string `ncpdp:"16BM" json:"narrativeMessage" jsonWithId:"390_BM_narrativeMessage"`
}
