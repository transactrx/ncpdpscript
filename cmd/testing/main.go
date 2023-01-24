package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdpTwoClaims := "016904D0B1SRRGA     1011114917903     20230119          \u001E\u001CAM04\u001CC28Y07Y99NR52         \u001CCCORVILLE     \u001CCDPARKER         \u001CC90\u001CC1               \u001CC61\u001E\u001CAM01\u001CC419441101\u001CC51\u001CCAORVILLE     \u001CCBPARKER         \u001CCM801 BRIM ST                   \u001CCNDESLOGE             \u001CCOMO\u001CCP63601          \u001D\u001E\u001CAM07\u001CEM1\u001CD2000008339523\u001CE103\u001CD780777028205        \u001CE70000000500\u001CD300\u001CD5001\u001CD61\u001CD80\u001CDE20230119\u001CDF00\u001CDJ5\u001CET0000000500\u001CC800\u001CDT0\u001CEU00\u001CEV00000000000\u001CU701\u001E\u001CAM11\u001CD90000000{\u001CDC0000400{\u001CDQ0000400{\u001CDU0000400{\u001CDN01\u001E\u001CAM03\u001CEZ01\u001CDB1750496246     \u001CDRTURABELIDZE    \u001C2NMO"
	//ncpdpTwoClaims := "610011D0B2IRX       1011255644357     20180227          \u001E\u001CAM04\u001CC2MEPE01082018\u001CC1GNRX\u001D\u001E\u001CAM07\u001CEM1\u001CD2181791740140028\u001CE103\u001CD758406044504"

	ncpdpBytes := []byte(ncpdpTwoClaims)

	trxType, err := scriptparser.DetermineTransactionType(ncpdpBytes)
	log.Printf("Transaction Type:%s", trxType)
	if err != nil {
		log.Printf("Unamarshal err: %v", err)
	}

	if trxType == "B1" {
		parseB1(ncpdpBytes)
	} else if trxType == "B2" {
		parseB2(ncpdpBytes)
	}

}

func parseB1(ncpdpBytes []byte) {
	msg := ncpdpscriptmessages.B1{}
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

func parseB2(ncpdpBytes []byte) {
	msg := ncpdpscriptmessages.B2{}
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
