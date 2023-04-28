package segments

import "github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"

type B2ResponseHeader struct {
	VersionRel                 *string `ncpdp:"header:versionRel" json:"versionRel"`
	TransactionCode            *string `ncpdp:"header:transactionCode" json:"transactionCode"`
	TransactionCount           *int    `ncpdp:"header:transactionCount" json:"transactionCount"`
	TransactionResponseStatus  *string `ncpdp:"header:transactionResponseStatus" json:"transactionResponseStatus"`
	ServiceProviderIdQualifier *string `ncpdp:"header:serviceProviderIdQualifier" json:"serviceProviderIdQualifier"`
	ServiceProviderId          *string `ncpdp:"header:serviceProviderId" json:"serviceProviderId"`
	Dos                        *string `ncpdp:"header:dos" json:"dos"`
}

type B2ResponseMessage struct {
	SegmentID *string `ncpdp:"20AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	Message   *string `ncpdp:"20F4" json:"message" jsonWithId:"504_F4_message"`
}

type B2ResponseStatus struct {
	SegmentID                              *string   `ncpdp:"21AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	TransactionResponseStatus              *string   `ncpdp:"21AN" json:"transactionResponseStatus" jsonWithId:"112_AN_transactionResponseStatus"`
	AuthorizationNumber                    *string   `ncpdp:"21F3" json:"authorizationNumber" jsonWithId:"503_F3_authorizationNumber"`
	RejectCount                            *int      `ncpdp:"21FA" json:"rejectCount" jsonWithId:"510_FA_rejectCount"`
	RejectCode                             []*string `ncpdp:"21FB" json:"rejectCode" jsonWithId:"511_FB_rejectCode"`
	RejectFieldOccurrenceIndicator         *string   `ncpdp:"214F" json:"rejectFieldOccurrenceIndicator" jsonWithId:"546_4F_rejectFieldOccurrenceIndicator"`
	ApprovedMessageCodeCount               *int      `ncpdp:"215F" json:"approvedMessageCodeCount" jsonWithId:"547_5F_approvedMessageCodeCount"`
	ApprovedMessageCode                    []*string `ncpdp:"216F" json:"approvedMessageCode" jsonWithId:"548_6F_approvedMessageCode"`
	AdditionalMessageInformationQualifier  []*string `ncpdp:"21UH" json:"additionalMessageInformationQualifier" jsonWithId:"132_UH_additionalMessageInformationQualifier"`
	AdditionalMessageInformation           []*string `ncpdp:"21FQ" json:"additionalMessageInformation" jsonWithId:"526_FQ_additionalMessageInformation"`
	AdditionalMessageInformationCount      *int      `ncpdp:"21UF" json:"additionalMessageInformationCount" jsonWithId:"130_UF_additionalMessageInformationCount"`
	AdditionalMessageInformationContinuity []*string `ncpdp:"21UG" json:"additionalMessageInformationContinuity" jsonWithId:"131_UG_additionalMessageInformationContinuity"`
	HelpDeskPhoneNumberQualifier           *string   `ncpdp:"217F" json:"helpDeskPhoneNumberQualifier" jsonWithId:"549_7F_helpDeskPhoneNumberQualifier"`
	HelpDeskPhoneNumber                    *string   `ncpdp:"218F" json:"helpDeskPhoneNumber" jsonWithId:"550_8F_helpDeskPhoneNumber"`
	TransactionReferenceNumber             *string   `ncpdp:"21K5" json:"transactionReferenceNumber" jsonWithId:"880_K5_transactionReferenceNumber"`
	InternalControlNumber                  *string   `ncpdp:"21A7" json:"internalControlNumber" jsonWithId:"993_A7_internalControlNumber"`
	URL                                    *string   `ncpdp:"21MA" json:"url" jsonWithId:"987_MA_url"`
}
type B2ResponseClaim struct {
	SegmentID                                   *string                    `ncpdp:"22AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	PrescriptionServiceReferenceNumberQualifier *string                    `ncpdp:"22EM" json:"prescriptionServiceReferenceNumberQualifier" jsonWithId:"455_EM_prescriptionServiceReferenceNumberQualifier"`
	PrescriptionServiceReferenceNumber          *string                    `ncpdp:"22D2" json:"prescriptionServiceReferenceNumber" jsonWithId:"402_D2_prescriptionServiceReferenceNumber"`
	PreferredProductCount                       *int                       `ncpdp:"229F" json:"preferredProductCount" jsonWithId:"551_9F_preferredProductCount"`
	PreferredProductIDQualifier                 []*string                  `ncpdp:"22AP" json:"preferredProductIDQualifier" jsonWithId:"552_AP_preferredProductIDQualifier"`
	PreferredProductID                          []*string                  `ncpdp:"22AR" json:"preferredProductID" jsonWithId:"553_AR_preferredProductID"`
	PreferredProductIncentive                   []*scriptparser.Currency64 `ncpdp:"22AS" json:"preferredProductIncentive" jsonWithId:"554_AS_preferredProductIncentive"`
	PreferredProductCostShareIncentive          []*scriptparser.Currency64 `ncpdp:"22AT" json:"preferredProductCostShareIncentive" jsonWithId:"555_AT_preferredProductCostShareIncentive"`
	PreferredProductDescription                 []*string                  `ncpdp:"22AU" json:"preferredProductDescription" jsonWithId:"556_AU_preferredProductDescription"`
	MedicaidSubrogationInternalControlNumber    *string                    `ncpdp:"22N4" json:"medicaidSubrogationInternalControlNumber" jsonWithId:"114_N4_medicaidSubrogationInternalControlNumber"`
}
type B2ResponsePricing struct {
	SegmentID                                                    *string                  `ncpdp:"23AM" json:"segmentID" jsonWithId:"111_AM_segmentID"`
	PatientPayAmount                                             *scriptparser.Currency64 `ncpdp:"23F5" json:"patientPayAmount" jsonWithId:"505_F5_patientPayAmount"`
	IngredientCostPaid                                           *scriptparser.Currency64 `ncpdp:"23F6" json:"ingredientCostPaid" jsonWithId:"506_F6_ingredientCostPaid"`
	DispensingFeePaid                                            *scriptparser.Currency64 `ncpdp:"23F7" json:"dispensingFeePaid" jsonWithId:"507_F7_dispensingFeePaid"`
	TaxExemptIndicator                                           *string                  `ncpdp:"23AV" json:"taxExemptIndicator" jsonWithId:"557_AV_taxExemptIndicator"`
	FlatSalesTaxAmountPaid                                       *scriptparser.Currency64 `ncpdp:"23AW" json:"flatSalesTaxAmountPaid" jsonWithId:"558_AW_flatSalesTaxAmountPaid"`
	PercentageSalesTaxAmountPaid                                 *scriptparser.Currency64 `ncpdp:"23AX" json:"percentageSalesTaxAmountPaid" jsonWithId:"559_AX_percentageSalesTaxAmountPaid"`
	PercentageSalesTaxRatePaid                                   *scriptparser.Currency64 `ncpdp:"23AY" json:"percentageSalesTaxRatePaid" jsonWithId:"560_AY_percentageSalesTaxRatePaid"`
	PercentageSalesTaxBasisPaid                                  *scriptparser.Currency64 `ncpdp:"23AZ" json:"percentageSalesTaxBasisPaid" jsonWithId:"561_AZ_percentageSalesTaxBasisPaid"`
	IncentiveAmountPaid                                          *scriptparser.Currency64 `ncpdp:"23FL" json:"incentiveAmountPaid" jsonWithId:"521_FL_incentiveAmountPaid"`
	ProfessionalServiceFeePaid                                   *string                  `ncpdp:"23J1" json:"professionalServiceFeePaid" jsonWithId:"562_J1_professionalServiceFeePaid"`
	OtherAmountPaidCount                                         *int                     `ncpdp:"23J2" json:"otherAmountPaidCount" jsonWithId:"563_J2_otherAmountPaidCount"`
	OtherAmountPaidQualifier                                     *string                  `ncpdp:"23J3" json:"otherAmountPaidQualifier" jsonWithId:"564_J3_otherAmountPaidQualifier"`
	OtherAmountPaid                                              *scriptparser.Currency64 `ncpdp:"23J4" json:"otherAmountPaid" jsonWithId:"565_J4_otherAmountPaid"`
	OtherPayerAmountRecognized                                   *scriptparser.Currency64 `ncpdp:"23J5" json:"otherPayerAmountRecognized" jsonWithId:"566_J5_otherPayerAmountRecognized"`
	TotalAmountPaid                                              *scriptparser.Currency64 `ncpdp:"23F9" json:"totalAmountPaid" jsonWithId:"509_F9_totalAmountPaid"`
	BasisOfReimbursementDetermination                            *string                  `ncpdp:"23FM" json:"basisOfReimbursementDetermination" jsonWithId:"522_FM_basisOfReimbursementDetermination"`
	AmountAttributedToSalesTax                                   *scriptparser.Currency64 `ncpdp:"23FN" json:"amountAttributedToSalesTax" jsonWithId:"523_FN_amountAttributedToSalesTax"`
	AccumulatedDeductibleAmount                                  *scriptparser.Currency64 `ncpdp:"23FC" json:"accumulatedDeductibleAmount" jsonWithId:"512_FC_accumulatedDeductibleAmount"`
	RemainingDeductibleAmount                                    *scriptparser.Currency64 `ncpdp:"23FD" json:"remainingDeductibleAmount" jsonWithId:"513_FD_remainingDeductibleAmount"`
	RemainingBenefitAmount                                       *scriptparser.Currency64 `ncpdp:"23FE" json:"remainingBenefitAmount" jsonWithId:"514_FE_remainingBenefitAmount"`
	AmountAppliedToPeriodicDeductible                            *scriptparser.Currency64 `ncpdp:"23FH" json:"amountAppliedToPeriodicDeductible" jsonWithId:"517_FH_amountAppliedToPeriodicDeductible"`
	AmountOfCopay                                                *scriptparser.Currency64 `ncpdp:"23FI" json:"amountOfCopay" jsonWithId:"518_FI_amountOfCopay"`
	AmountExceedingPeriodicBenefitMax                            *scriptparser.Currency64 `ncpdp:"23FK" json:"amountExceedingPeriodicBenefitMax" jsonWithId:"520_FK_amountExceedingPeriodicBenefitMax"`
	BasisOfCalculationDispensingFee                              *string                  `ncpdp:"23HH" json:"basisOfCalculationDispensingFee" jsonWithId:"346_HH_basisOfCalculationDispensingFee"`
	BasisOfCalculationCopay                                      *string                  `ncpdp:"23HJ" json:"basisOfCalculationCopay" jsonWithId:"347_HJ_basisOfCalculationCopay"`
	BasisOfCalculationFlatSalesTax                               *string                  `ncpdp:"23HK" json:"basisOfCalculationFlatSalesTax" jsonWithId:"348_HK_basisOfCalculationFlatSalesTax"`
	BasisOfCalculationPercentageSalesTax                         *string                  `ncpdp:"23HM" json:"basisOfCalculationPercentageSalesTax" jsonWithId:"349_HM_basisOfCalculationPercentageSalesTax"`
	AmountAttributedToProcessorFee                               *scriptparser.Currency64 `ncpdp:"23NZ" json:"amountAttributedToProcessorFee" jsonWithId:"571_NZ_amountAttributedToProcessorFee"`
	PatientSalesTaxAmount                                        *scriptparser.Currency64 `ncpdp:"23EQ" json:"patientSalesTaxAmount" jsonWithId:"575_EQ_patientSalesTaxAmount"`
	PlanSalesTaxAmount                                           *scriptparser.Currency64 `ncpdp:"232Y" json:"planSalesTaxAmount" jsonWithId:"574_2Y_planSalesTaxAmount"`
	AmountOfCoinsurance                                          *scriptparser.Currency64 `ncpdp:"234U" json:"amountOfCoinsurance" jsonWithId:"572_4U_amountOfCoinsurance"`
	BasisOfCalculationCoinsurance                                *string                  `ncpdp:"234V" json:"basisOfCalculationCoinsurance" jsonWithId:"573_4V_basisOfCalculationCoinsurance"`
	BenefitStageCount                                            *int                     `ncpdp:"23MU" json:"benefitStageCount" jsonWithId:"392_MU_benefitStageCount"`
	BenefitStageQualifier                                        *string                  `ncpdp:"23MV" json:"benefitStageQualifier" jsonWithId:"393_MV_benefitStageQualifier"`
	BenefitStageAmount                                           *scriptparser.Currency64 `ncpdp:"23MW" json:"benefitStageAmount" jsonWithId:"394_MW_benefitStageAmount"`
	EstimatedGenericSavings                                      *scriptparser.Currency64 `ncpdp:"23G3" json:"estimatedGenericSavings" jsonWithId:"577_G3_estimatedGenericSavings"`
	SpendingAccountAmountRemaining                               *scriptparser.Currency64 `ncpdp:"23UC" json:"spendingAccountAmountRemaining" jsonWithId:"128_UC_spendingAccountAmountRemaining"`
	HealthPlanFundedAssistanceAmount                             *scriptparser.Currency64 `ncpdp:"23UD" json:"healthPlanFundedAssistanceAmount" jsonWithId:"129_UD_healthPlanFundedAssistanceAmount"`
	AmountAttributedToProviderNetworkSelection                   *scriptparser.Currency64 `ncpdp:"23UJ" json:"amountAttributedToProviderNetworkSelection" jsonWithId:"133_UJ_amountAttributedToProviderNetworkSelection"`
	AmountAttributedToProductSelectionBrandDrug                  *scriptparser.Currency64 `ncpdp:"23UK" json:"amountAttributedToProductSelectionBrandDrug" jsonWithId:"134_UK_amountAttributedToProductSelectionBrandDrug"`
	AmountAttributedToProductSelectionNonPreferredFormulary      *scriptparser.Currency64 `ncpdp:"23UM" json:"amountAttributedToProductSelectionNonPreferredFormulary" jsonWithId:"135_UM_amountAttributedToProductSelectionNonPreferredFormulary"`
	AmountAttributedToProductSelectionBrandNonPreferredFormulary *scriptparser.Currency64 `ncpdp:"23UN" json:"amountAttributedToProductSelectionBrandNonPreferredFormulary" jsonWithId:"136_UN_amountAttributedToProductSelectionBrandNonPreferredFormulary"`
	AmountAttributedToCoverageGap                                *scriptparser.Currency64 `ncpdp:"23UP" json:"amountAttributedToCoverageGap" jsonWithId:"137_UP_amountAttributedToCoverageGap"`
	IngredientCostContractedReimbursableAmount                   *scriptparser.Currency64 `ncpdp:"23U8" json:"ingredientCostContractedReimbursableAmount" jsonWithId:"148_U8_ingredientCostContractedReimbursableAmount"`
	DispensingFeeContractedReimbursableAmount                    *scriptparser.Currency64 `ncpdp:"23U9" json:"dispensingFeeContractedReimbursableAmount" jsonWithId:"149_U9_dispensingFeeContractedReimbursableAmount"`
}

func (b2RespStatus *B2ResponseStatus) GetAdditionalMessageInformation() string {
	//combine all message arrays into on string
	var message string
	if b2RespStatus.AdditionalMessageInformation != nil {
		msgInfos := (*b2RespStatus).AdditionalMessageInformation
		for _, msg := range msgInfos {
			message += *msg + " "
		}
		return message
	}
	return ""
}
