package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdpB1Request := "016904D0B1SRRGA     1011114917903     20230119          \u001E\u001CAM04\u001CC28Y07Y99NR52         \u001CCCORVILLE     \u001CCDPARKER         \u001CC90\u001CC1               \u001CC61\u001E\u001CAM01\u001CC419441101\u001CC51\u001CCAORVILLE     \u001CCBPARKER         \u001CCM801 BRIM ST                   \u001CCNDESLOGE             \u001CCOMO\u001CCP63601          \u001D\u001E\u001CAM07\u001CEM1\u001CD2000008339523\u001CE103\u001CD780777028205        \u001CE70000000500\u001CD300\u001CD5001\u001CD61\u001CD80\u001CDE20230119\u001CDF00\u001CDJ5\u001CET0000000500\u001CC800\u001CDT0\u001CEU00\u001CEV00000000000\u001CU701\u001E\u001CAM11\u001CD90000000{\u001CDC0000400{\u001CDQ0000400{\u001CDU0000400{\u001CDN01\u001E\u001CAM03\u001CEZ01\u001CDB1750496246     \u001CDRTURABELIDZE    \u001C2NMO"
	ncpdpBytesB1Request := []byte(ncpdpB1Request)
	parseB1Request(ncpdpBytesB1Request)
	//
	//ncpdpB2Request := "610011D0B2IRX       1011255644357     20180227          \u001E\u001CAM04\u001CC2MEPE01082018\u001CC1GNRX\u001D\u001E\u001CAM07\u001CEM1\u001CD2181791740140028\u001CE103\u001CD758406044504"
	//ncpdpBytesB2Request := []byte(ncpdpB2Request)
	//parseB2Request(ncpdpBytesB2Request)
	//
	ncpdpB1Response := "D0B11A011548217136     20230124\u001E\u001CAM20\u001CF4UNIT OF MEASURE IS REQUIRED\u001D\u001E\u001CAM21\u001CANR\u001CF322653129\u001CFA1\u001CFB26\u001CUF0\u001C7F03\u001C8F8665223386\u001E\u001CAM22\u001CEM1\u001CD2000006133614\u001C9F0\u001E\u001CAM23\u001CJ20"
	ncpdpBytesB1Response := []byte(ncpdpB1Response)
	parseB1Response(ncpdpBytesB1Response)

	//ncpdpB2Response := "D0B21A011548217136     20221017\u001E\u001CAM20\u001CF4CLAIM TO OLD\u001D\u001E\u001CAM21\u001CANR\u001CF321947605\u001CFA1\u001CFB87\u001CUF1\u001CUH1\u001CFQREVERSAL FAILED \u001C7F03\u001C8F8665223386\u001E\u001CAM22\u001CEM1\u001CD2000006130674\u001C9F0\u001E\u001CAM23\u001CJ20"
	//ncpdpBytesB2Response := []byte(ncpdpB2Response)
	//parseB2Response(ncpdpBytesB2Response)

}

func parseB1Response(ncpdpBytes []byte) {

	transactionType, err := scriptparser.DetermineTransactionType(ncpdpBytes)
	if err != nil {
		log.Fatal(err)
	}
	if transactionType != scriptparser.B1ResponseType {
		log.Fatal("Transaction type is not B1")
	}
	log.Printf("Transaction Type:%s", transactionType)

	msg := ncpdpscriptmessages.B1Response{}
	err = scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

}

func parseB1Request(ncpdpBytes []byte) {
	transactionType, err := scriptparser.DetermineTransactionType(ncpdpBytes)
	if err != nil {
		log.Fatal(err)
	}
	if transactionType != scriptparser.B1RequestType {
		log.Fatal("Transaction type is not B1")
	}
	log.Printf("Transaction Type:%s", transactionType)
	msg := ncpdpscriptmessages.B1Request{}
	err = scriptparser.Unmarshal(ncpdpBytes, &msg)

	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	ncpdpJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}
	log.Printf("NCPDP %s", ncpdpJson)

	trxType, err := scriptparser.DetermineTransactionType(ncpdpBytes)
	log.Printf("Transaction Type:%s", trxType)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

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

	trxType, err := scriptparser.DetermineTransactionType(ncpdpBytes)
	log.Printf("Transaction Type:%s", trxType)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

}
