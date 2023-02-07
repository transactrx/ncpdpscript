package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdp := "D0B11A011669688842     20230203\u001E\u001CAM25\u001CC1HAPMDD\u001C2FDSIVC\u001D\u001E\u001CAM21\u001CANP\u001CF3590171942594114101\u001C5F2\u001C6F038\u001C6F036\u001CUF1\u001CUH13\u001CFQMAXIMUM DAY SUPPLY OF 030 ALLOWED\u001E\u001CAM22\u001CEM1\u001CD29724828\u001E\u001CAM23\u001CF500{\u001CF6407F\u001CF737E\u001CAW00{\u001CAX00{\u001CFL200{\u001CF9645A\u001CFM9\u001CMU1\u001CMV02\u001CMW645A\u001E\u001CAM24\u001CJ61\u001CE4LD\u001CFS0\u001CFU20230203\u001CFV500\u001CFW1\u001CFYMIN DAILY DOSAGE    0.50"
	//ncpdp := "D0B11A011669688842     20230201\u001E\u001CAM25\u001C2FMDR0S100P3\u001E\u001CAM29\u001CCACAROLYN\u001CCBMILLER\u001CC419570207\u001D\u001E\u001CAM21\u001CANR\u001CF3230381886082321999\u001CFA3\u001CFB75\u001CFBA6\u001CFB569\u001CUF3\u001CUH01\u001CUH02\u001CUH03\u001CFQBVD PA REQ RPH/PRVDR 1-855-344-0930 DRUG REQUIRES PRIOR AUTHORIZATION (PHARMACY HELP DESK 1-866-281-0635) \u001C7F03\u001C8F8662810635\u001E\u001CAM22\u001CEM1\u001CD29724808\u001C9F0\u001E\u001CAM23\u001CJ20"

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
