package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	//todo error: Unamarshal err: strconv.Atoi: parsing "OA": invalid syntax
	//ncpdp := "016904D0B1SMMS0     1011447260393     20230206          \u001E\u001CAM04\u001CC2IHT3H4OUG88\u001CC301\u001CC61\u001E\u001CAM01\u001CCX01\u001CCY1HT3H40UG\u001CC419521025\u001CC51\u001CCARICHARD\u001CCBHICKS\u001CCM4830 NORTH STATE STREET\u001CCNJACKSON\u001CCOMS\u001CCP39206\u001CCQ6019427689\u001C4X04\u001D\u001E\u001CAM07\u001CEM1\u001CD27789295\u001CE103\u001CD700904530680\u001CE70000030000\u001CD300\u001CD530\u001CD61\u001CD80\u001CDE20230206\u001CDF00\u001CDJ2\u001CC8OA\u001CDT2\u001C28EA\u001CU701\u001E\u001CAM03\u001CEZ01\u001CDB1750833984\u001CDRTREST\u001CPM6019414688\u001C2E01\u001CDL1750833984\u001C4ETREST\u001C2JCLINT\u001C2K275 N BRANDON BLVD\u001C2MBRANDON\u001C2NMS\u001C2P39042\u001E\u001CAM05\u001C4C1\u001C5C01\u001CE820230206\u001C5E2\u001C6E70\u001C6EA5\u001E\u001CAM11\u001CD95D\u001CDC59A\u001CDQ419I\u001CDU64E\u001CDN01\u001E\u001CAM15\u001C3QOASIS HAVEN\u001C3U1458 MONCURE MARBLE ROAD\u001C5JTERRY\u001C3VMS\u001C6D39170"

	//todo error: Unamarshal err: strconv.Atoi: parsing "SI": invalid syntax
	//ncpdp := "016904D0B1SMMS0     1011417236894     20230105          \u001E\u001CAM04\u001CC23QF7Q58PN63\u001CC1TRRX\u001CC61\u001E\u001CAM01\u001CCX01\u001CCY587982037\u001CC419560516\u001CC52\u001CCALUTRICIA\u001CCBDUPREE\u001CCM5911 HELMS COURT\u001CCNJACKSON\u001CCOMS\u001CCP39213\u001CCQ6013219653\u001C2C1\u001C4X0\u001D\u001E\u001CAM07\u001CEM1\u001CD26172261\u001CE103\u001CD700169633910\u001CE70000015000\u001CD313\u001CD55\u001CD61\u001CD89\u001CDE20220121\u001CDF03\u001CDJ2\u001CC8SI\u001C28ML\u001CU701\u001E\u001CAM03\u001CEZ01\u001CDB1336535053\u001CDRGLOVER\u001CPM6019845610\u001C2E01\u001CDL1336535053\u001C4EGLOVER\u001C2JWENDELL\u001C2K1410 E WOODROW WILSON AVE\u001C2MJACKSON\u001C2NMS\u001C2P392137681\u001E\u001CAM05\u001C4C1\u001C5C01\u001CE820230105\u001C5E2\u001C6E569\u001C6E88\u001E\u001CAM11\u001CD96706{\u001CDC59A\u001CDQ7329I\u001CDU6765A\u001CDN01\u001E\u001CAM15\u001C3QSINCERE HOME CARE\u001C3U5911 HOLMES COURT\u001C5JJACKSON\u001C3VMS\u001C6D39213"

	//todo error: Unamarshal err: strconv.Atoi: parsing "HE": invalid syntax
	//ncpdp := "016904D0B1SMMS0     1011447260393     20221222          \u001E\u001CAM04\u001CC2426191044A\u001CC301\u001CC61\u001E\u001CAM01\u001CCX01\u001CCY426191044\u001CC419620423\u001CC51\u001CCAHOWARD\u001CCBELMORE\u001CCM5008 HWY.80 WEST\u001CCNJACKSON\u001CCOMS\u001CCP39209\u001CCQ7777000000\u001C4X0\u001D\u001E\u001CAM07\u001CEM1\u001CD27779447\u001CE103\u001CD733342007215\u001CE70000030000\u001CD300\u001CD530\u001CD61\u001CD80\u001CDE20221222\u001CDF00\u001CDJ2\u001CC8HE\u001C28EA\u001CU701\u001E\u001CAM03\u001CEZ01\u001CDB1053753194\u001CDRCALHOUN\u001CPM6019859355\u001C2E01\u001CDL1053753194\u001C4ECALHOUN\u001C2JTANYA\u001C2K203 REMINGTON DR\u001C2MBRANDON\u001C2NMS\u001C2P390422829\u001E\u001CAM05\u001C4C1\u001C5C01\u001CE820221222\u001C5E1\u001C6E88\u001E\u001CAM11\u001CD911115{\u001CDC122I\u001CDQ3582I\u001CDU11237I\u001CDN01\u001E\u001CAM15\u001C3QPARADISE PERSONAL CARE HOME\u001C3U20114 HWY 18\u001C5JHERMANVILLE\u001C3VMS\u001C6D39086"

	ncpdpB1Response := ``
	ncpdpBytesRequest := []byte(ncpdpB1Response)
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
