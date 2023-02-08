package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdp := "016904D0S1SMFL0     1011912086687     20221214          \u001E\u001CAM01\u001CC419391127\u001CC52\u001CCAROSE\u001CCBSANDERS\u001CCM2000 LOWSON BLVD\u001CCNDELRAY BEACH\u001CCOFL\u001CCP33445\u001CC712\u001C2C1\u001E\u001CAM04\u001CC27E68KG2VM54\u001CCCROSE\u001CCDSANDERS\u001CC90\u001CC3   \u001CC61\u001D\u001E\u001CAM07\u001CEM2\u001CD21588541\u001CE107\u001CD7M0201\u001CE71000\u001CD300\u001CD51\u001CDE20221115\u001CDF00\u001CET1000\u001CC800\u001CDI00\u001CU71\u001E\u001CAM08\u001C7E1\u001CE4PH\u001CE5MA\u001CE63N\u001E\u001CAM11\u001CBE350{\u001CDQ0000350{\u001CDU0000350{"

	ncpdpBytesRequest := []byte(ncpdp)
	transactionType, err := scriptparser.DetermineTransactionType(ncpdpBytesRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction Type:%v", transactionType)
	if transactionType == scriptparser.B1RequestType {
		parseB1Request(ncpdpBytesRequest)
	} else if transactionType == scriptparser.B2RequestType {
		parseB2Request(ncpdpBytesRequest)
	} else if transactionType == scriptparser.B1ResponseType {
		parseB1Response(ncpdpBytesRequest)
	} else if transactionType == scriptparser.B2ResponseType {
		parseB2Response(ncpdpBytesRequest)
	}

}

func parseB2Response(ncpdpBytes []byte) {

	msg := ncpdpscriptmessages.B2Response{}
	err := scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

	isAccepted := msg.IsTransactionResponseAccepted()
	log.Printf("IsTransactionB2ResponseAccepted: %t", isAccepted)
	ncpdpJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	log.Printf("NCPDP %s", ncpdpJson)
}

func parseB1Response(ncpdpBytes []byte) {

	msg := ncpdpscriptmessages.B1Response{}
	err := scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

	log.Printf("IsTransactionB1ResponseAccepted: %t", msg.IsTransactionResponseAccepted())
	if fullFormattedPayerB1Response, err := msg.GetFullFormattedPayerResponse(); err == nil {
		log.Printf("fullFormattedPayerB1Response: %s", fullFormattedPayerB1Response)
	}

	ncpdpJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	log.Printf("NCPDP %s", ncpdpJson)

}

func parseB1Request(ncpdpBytes []byte) {

	msg := ncpdpscriptmessages.B1Request{}
	err := scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	ncpdpJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	log.Printf("NCPDP %s", ncpdpJson)

}

func parseB2Request(ncpdpBytes []byte) {
	msg := ncpdpscriptmessages.B2Request{}
	err := scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	ncpdpJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	log.Printf("NCPDP %s", ncpdpJson)

}
