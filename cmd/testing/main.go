package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdp := "016904D0B1PARTD     1011962953034     20221130          \u001E\u001CAM04\u001CC231736950\u001CCCRICHARD\u001CCDKARASEK\u001CFOS4802_146\u001CC90\u001CC1788257\u001CC3001\u001CC61\u001CNU1KD8GP8UN61\u001E\u001CAM01\u001CC419560527\u001CC51\u001CCARICHARD\u001CCBKARASEK\u001CCM2575 CASCADE CREEK\u001CCNJACKSON\u001CCOMI\u001CCP49203\u001D\u001E\u001CAM07\u001CEM1\u001CD29700077\u001CE103\u001CD758160084211\u001CENPB523229460\u001CEP20221130\u001CSE1\u001CERGA\u001CE71000\u001CD301\u001E\u001CAM11\u001CD91020{\u001CE3350{\u001CDQ1370{\u001CDU1370{\u001E\u001CAM03\u001CEZ01\u001CDB1851820518\u001CDRRILEY\u001C2JPHILIP\u001E\u001CAM08\u001C7E1\u001CE5MA\u001C8E11\u001E\u001CAM13\u001CVE1\u001CWE01\u001CDOZ23"
	//ncpdp := "610011D0B2IRX       1011255644357     20180227          \u001E\u001CAM04\u001CC2MEPE01082018\u001CC1GNRX\u001D\u001E\u001CAM07\u001CEM1\u001CD2181791740140028\u001CE103\u001CD758406044504"
	//ncpdp := "016904D0B2SMOR0     1011548217136     20221017          \u001D\u001E\u001CAM07\u001CEM1\u001CD2000006130674\u001CE103\u001CD780777028299\u001CD300"
	//B1 response with GroupSeparatorByte,SegmentSeparatorByte,FieldSeparatorByte
	//ncpdp := "D0B11A011821181355     20220422\u001D\u001E\u001CAM21\u001CANR\u001CF322513327\u001CFA1\u001CFB07\u001CUF0\u001C7F03\u001C8F8665223386\u001E\u001CAM22\u001CEM1\u001CD2000006479563\u001C9F0\u001E\u001CAM23\u001CJ20\nD0B11A011548217"
	//B1 response with SegmentSeparatorByte,FieldSeparatorByte
	//ncpdp := "D0B11A011548217136     20230124\u001E\u001CAM20\u001CF4UNIT OF MEASURE IS REQUIRED\u001D\u001E\u001CAM21\u001CANR\u001CF322653129\u001CFA1\u001CFB26\u001CUF0\u001C7F03\u001C8F8665223386\u001E\u001CAM22\u001CEM1\u001CD2000006133614\u001C9F0\u001E\u001CAM23\u001CJ20"
	//ncpdp := "D0B11A011669688842     20230130\u001E\u001CAM25\u001CFOEGWP\u001C2FIRX7MP\u001CC2M0024718700\u001E\u001CAM29\u001CCAFRANK\u001CCBVROOM\u001CC419421224\u001D\u001E\u001CAM21\u001CANR\u001CF3230301644105018999\u001CFA2\u001CFB75\u001CFB569\u001CUF2\u001CUH01\u001CUH02\u001CFQIF LOC CHANGE CALL 1-855-577-6517 DRUG REQUIRES PRIOR AUTHORIZATION \u001E\u001CAM22\u001CEM1\u001CD2000009972710\u001C9F0\u001E\u001CAM23\u001CJ20\u001E\u001CAM26\u001CPR20230130"
	//ncpdp := "D0B21A011548217136     20221017\u001E\u001CAM20\u001CF4CLAIM TO OLD\u001D\u001E\u001CAM21\u001CANR\u001CF321947605\u001CFA1\u001CFB87\u001CUF1\u001CUH1\u001CFQREVERSAL FAILED \u001C7F03\u001C8F8665223386\u001E\u001CAM22\u001CEM1\u001CD2000006130674\u001C9F0\u001E\u001CAM23\u001CJ20"
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

	isAccepted := ncpdpscriptmessages.IsTransactionB2ResponseAccepted(msg)
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

	isAccepted := ncpdpscriptmessages.IsTransactionB1ResponseAccepted(msg)
	log.Printf("IsTransactionB1ResponseAccepted: %t", isAccepted)

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
