package ncpdpscriptmessages

import (
	"fmt"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages/segments"
	"strconv"
	"strings"
)

func getPayerResponse(isAccepted bool, claims []*ResponseClaim, message *segments.Message) string {
	var msg = ""
	if isAccepted {

		msg += getFormattedPayerRejectedResponse(claims, message)
		dur := getFormattedDUR(claims)
		if len(strings.TrimSpace(dur)) > 0 {
			if len(strings.TrimSpace(msg)) > 0 {
				msg += "\n"
			}
			msg += dur
		}
	} else {
		additionalMessageInfo := getStatusAdditionalMessageInfo(claims)
		if len(strings.TrimSpace(additionalMessageInfo)) > 0 {
			msg += "Response Message: " + strings.TrimSpace(additionalMessageInfo)
			msg += getMessageEmptyIfNull(message)
		} else {
			msg += getMessageEmptyIfNull(message)
		}

		dur := getFormattedDUR(claims)
		if len(strings.TrimSpace(dur)) > 0 {
			if len(strings.TrimSpace(msg)) > 0 {
				msg += "\n"
			}
			msg += dur
		}
	}
	return msg
}

func GetFormattedPayerResponse(obj interface{}) (string, error) {
	if response, ok := any(obj).(B1Response); ok == true {
		return getPayerResponse(!response.IsTransactionResponseAccepted(), response.Claims, response.Message), nil
	} else if response, ok := any(obj).(B2Response); ok == true {
		return getPayerResponse(!response.IsTransactionResponseAccepted(), response.Claims, response.Message), nil
	} else {
		return "", scriptparser.NCPDPMessageInvalidOrUnsupported
	}
}

func getFormattedPayerRejectedResponse(claims []*ResponseClaim, message *segments.Message) string {
	result := ""

	if claims != nil && len(claims) > 0 && claims[0].ResponseStatus != nil && claims[0].ResponseStatus.AdditionalMessageInfo != nil && len(strings.TrimSpace(*claims[0].ResponseStatus.AdditionalMessageInfo)) > 0 {
		additionalMessageInfo := strings.TrimSpace(*claims[0].ResponseStatus.AdditionalMessageInfo)
		result += fmt.Sprintf("Response Message: %s %s", strings.TrimSpace(additionalMessageInfo), getMessageEmptyIfNull(message))
	} else {
		result += getMessageEmptyIfNull(message)
	}

	return result
}

func getFormattedDUR(claims []*ResponseClaim) string {
	result := ""
	if claims != nil && len(claims) > 0 && claims[0].ResponseDUR != nil {
		responseDUR := claims[0].ResponseDUR
		if responseDUR != nil {

			rfsc := responseDUR.ReasonForServiceCode
			if rfsc != nil {
				for _, value := range rfsc {
					result += "Reason For Service Code: " + getReasonForServiceCodeDescription(*value) + "\n"
				}
			}

			clinicalSignificanceCode := responseDUR.ClinicalSignificanceCode
			if clinicalSignificanceCode != nil {
				for _, value := range clinicalSignificanceCode {
					if value != nil && len(*value) > 0 {
						csc, err := strconv.ParseInt(*value, 0, 8)
						if err != nil {
							switch csc {
							case 1:
								result += "Clinical Significance Code: Major\n"
							case 2:
								result += "Clinical Significance Code: Moderate\n"
							case 3:
								result += "Clinical Significance Code: Minor\n"
							default:
								result += ""
							}
						}
					}
				}
			}

			otherPharmacyIndicator := responseDUR.OtherPharmacyIndicator
			if otherPharmacyIndicator != nil {
				for _, value := range otherPharmacyIndicator {
					if value != nil {
						switch *value {
						case 0:
							result += "Other Pharmacy Indicator: Not Specified\n"
						case 1:
							result += "Other Pharmacy Indicator: Your Pharmacy\n"
						case 2:
							result += "Other Pharmacy Indicator: Other Pharmacy in Same Chain\n"
						case 3:
							result += "Other Pharmacy Indicator: Other Pharmacy\n"
						default:
							result += ""
						}
					}
				}
			}

			previousDateOfFill := responseDUR.PreviousDateOfFill
			if previousDateOfFill != nil {
				for _, value := range previousDateOfFill {
					if value != nil {
						result += fmt.Sprintf("Previous Date Of Fill: %s\n", *value)
					}
				}
			}

			qty := responseDUR.PreviousFillQuantity
			if qty != nil {
				for _, value := range qty {
					if value != nil {
						result += fmt.Sprintf("Previous Fill Quantity: %d\n", *value/1000)
					}
				}
			}

			dbI := responseDUR.DatabaseIndicator
			if dbI != nil {
				for _, value := range dbI {
					if value != nil {
						result += fmt.Sprintf("Database Indicator: %s\n", *value)
					}
				}
			}

			opi := responseDUR.OtherPrescriberIndicator
			if opi != nil {
				for _, value := range opi {
					if value != nil && len(*value) > 0 {
						v, err := strconv.ParseInt(*value, 0, 8)
						if err != nil {
							switch v {
							case 0:
								result += "Other Prescriber Indicator: Not Specified\n"
							case 2:
								result += "Other Prescriber Indicator: Same Prescriber\n"
							case 3:
								result += "Other Prescriber Indicator: Other Prescriber\n"
							default:
								result += ""
							}
						}
					}
				}
			}

			msg := responseDUR.Message
			if msg != nil {
				for _, value := range msg {
					if value != nil {
						result += fmt.Sprintf("Message: %s\n", *value)
					}
				}
			}

			if len(strings.TrimSpace(result)) > 0 {
				result = fmt.Sprintf("%s\n%s", "DUR Messages:", result)
			}

		}

	}
	return result
}

func getReasonForServiceCodeDescription(reasonForServiceCode string) string {
	result := ""
	if len(strings.TrimSpace(reasonForServiceCode)) > 0 {
		result = getReasonForServiceCodeMap()[strings.ToUpper(reasonForServiceCode)]
	}
	return result
}

func getMessageEmptyIfNull(message *segments.Message) string {
	result := ""
	if message != nil && len(strings.TrimSpace(*message.Message)) > 0 {
		result = *message.Message
	}
	return result
}

func getStatusAdditionalMessageInfo(claims []*ResponseClaim) string {
	result := ""
	if claims != nil && len(claims) > 0 && claims[0].ResponseStatus != nil && claims[0].ResponseStatus.AdditionalMessageInfo != nil && len(strings.TrimSpace(*claims[0].ResponseStatus.AdditionalMessageInfo)) > 0 {
		result = strings.TrimSpace(*claims[0].ResponseStatus.AdditionalMessageInfo)
	}
	return result
}

func getReasonForServiceCodeMap() map[string]string {
	REASON_FOR_SERVICE_CODE_MAP := make(map[string]string)
	REASON_FOR_SERVICE_CODE_MAP["AD"] = "Additional Drug Needed"
	REASON_FOR_SERVICE_CODE_MAP["AN"] = "Prescription Authentication"
	REASON_FOR_SERVICE_CODE_MAP["AR"] = "Adverse Drug Reaction"
	REASON_FOR_SERVICE_CODE_MAP["AT"] = "Additive Toxicity"
	REASON_FOR_SERVICE_CODE_MAP["CD"] = "Chronic Disease Management"
	REASON_FOR_SERVICE_CODE_MAP["CH"] = "Call Help Desk"
	REASON_FOR_SERVICE_CODE_MAP["CS"] = "Patient Complaint/Symptom"
	REASON_FOR_SERVICE_CODE_MAP["DA"] = "Drug-Allergy"
	REASON_FOR_SERVICE_CODE_MAP["DC"] = "Drug-Disease (Inferred)"
	REASON_FOR_SERVICE_CODE_MAP["DD"] = "Drug-Drug Interaction"
	REASON_FOR_SERVICE_CODE_MAP["DF"] = "Drug-Food interaction"
	REASON_FOR_SERVICE_CODE_MAP["DI"] = "Drug Incompatibility"
	REASON_FOR_SERVICE_CODE_MAP["DL"] = "Drug-Lab Conflict"
	REASON_FOR_SERVICE_CODE_MAP["DM"] = "Apparent Drug Misuse"
	REASON_FOR_SERVICE_CODE_MAP["DS"] = "Tobacco Use"
	REASON_FOR_SERVICE_CODE_MAP["ED"] = "Patient Education/Instruction"
	REASON_FOR_SERVICE_CODE_MAP["ER"] = "Overuse"
	REASON_FOR_SERVICE_CODE_MAP["EX"] = "Excessive Quantity"
	REASON_FOR_SERVICE_CODE_MAP["HD"] = "High Dose"
	REASON_FOR_SERVICE_CODE_MAP["IC"] = "Iatrogenic Condition"
	REASON_FOR_SERVICE_CODE_MAP["ID"] = "Ingredient Duplication"
	REASON_FOR_SERVICE_CODE_MAP["LD"] = "Low Dose"
	REASON_FOR_SERVICE_CODE_MAP["LK"] = "Lock In Recipient"
	REASON_FOR_SERVICE_CODE_MAP["LR"] = "Underuse"
	REASON_FOR_SERVICE_CODE_MAP["MC"] = "Drug-Disease (Reported)"
	REASON_FOR_SERVICE_CODE_MAP["MN"] = "Insufficeint Duration"
	REASON_FOR_SERVICE_CODE_MAP["MS"] = "Missing Information/Clarification"
	REASON_FOR_SERVICE_CODE_MAP["MX"] = "Excessive Duration"
	REASON_FOR_SERVICE_CODE_MAP["NA"] = "Drug Not Available"
	REASON_FOR_SERVICE_CODE_MAP["NC"] = "Non-covered Drug Purchase"
	REASON_FOR_SERVICE_CODE_MAP["ND"] = "New Disease/Diagnosis"
	REASON_FOR_SERVICE_CODE_MAP["NF"] = "Non-Formulary Drug"
	REASON_FOR_SERVICE_CODE_MAP["NN"] = "Unnecessary Drug"
	REASON_FOR_SERVICE_CODE_MAP["NP"] = "New Patient Processing"
	REASON_FOR_SERVICE_CODE_MAP["NR"] = "Lactation/Nursing Interaction"
	REASON_FOR_SERVICE_CODE_MAP["NS"] = "Insufficient Quantity"
	REASON_FOR_SERVICE_CODE_MAP["OH"] = "Alcohol Conflict"
	REASON_FOR_SERVICE_CODE_MAP["PA"] = "Drug-Age"
	REASON_FOR_SERVICE_CODE_MAP["PC"] = "Patient Question/Concern"
	REASON_FOR_SERVICE_CODE_MAP["PG"] = "Drug-Pregnancy"
	REASON_FOR_SERVICE_CODE_MAP["PH"] = "Preventive Health Care"
	REASON_FOR_SERVICE_CODE_MAP["PN"] = "Prescriber Consultation"
	REASON_FOR_SERVICE_CODE_MAP["PP"] = "Plan Protocol"
	REASON_FOR_SERVICE_CODE_MAP["PR"] = "Prior Adverse Reaction"
	REASON_FOR_SERVICE_CODE_MAP["PS"] = "Product Selection Opportunity"
	REASON_FOR_SERVICE_CODE_MAP["RE"] = "Suspected Environmental Risk"
	REASON_FOR_SERVICE_CODE_MAP["RF"] = "Health Provider Referral"
	REASON_FOR_SERVICE_CODE_MAP["SC"] = "Suboptimal Compliance"
	REASON_FOR_SERVICE_CODE_MAP["SD"] = "Suboptimal Drug/Indication"
	REASON_FOR_SERVICE_CODE_MAP["SE"] = "Side Effect"
	REASON_FOR_SERVICE_CODE_MAP["SF"] = "Suboptimal Dosage Form"
	REASON_FOR_SERVICE_CODE_MAP["SR"] = "Suboptimal Regimen"
	REASON_FOR_SERVICE_CODE_MAP["SX"] = "Drug-Gender"
	REASON_FOR_SERVICE_CODE_MAP["TD"] = "Therapeutic"
	REASON_FOR_SERVICE_CODE_MAP["TN"] = "Laboratory Test Needed"
	REASON_FOR_SERVICE_CODE_MAP["TP"] = "Payer/Processor Question"
	return REASON_FOR_SERVICE_CODE_MAP
}
