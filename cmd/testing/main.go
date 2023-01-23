package main

import (
	"encoding/json"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscript/scriptparser"
	"github.com/transactrx/ncpdpscript/pkg/ncpdpscriptmessages"
	"log"
)

func main() {

	ncpdpTwoClaims := "016904D0B2SMOR0     1011548217136     20221017          \u001D\u001E\u001CAM07\u001CEM1\u001CD2000006130674\u001CE103\u001CD780777028299\u001CD300"

	ncpdpBytes := []byte(ncpdpTwoClaims)
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
