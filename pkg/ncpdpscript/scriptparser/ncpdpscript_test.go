package scriptparser

import (
	"fmt"
	"log"
	"testing"
)

type Header struct {
	Bin                        string `ncpdp:"header:bin"`
	VersionRel                 string `ncpdp:"header:versionRel"`
	TransactionCode            string `ncpdp:"header:transactionCode"`
	Pcn                        string `ncpdp:"header:pcn"`
	TransactionCount           int    `ncpdp:"header:transactionCount"`
	ServiceProviderIdQualifier string `ncpdp:"header:serviceProviderIdQualifier"`
	ServiceProviderId          string `ncpdp:"header:serviceProviderId"`
	Dos                        string `ncpdp:"header:dos"`
}

type Insurance struct {
	SegmentId                          string  `ncpdp:"04AM"`
	CardHolderId                       *string `ncpdp:"04C2"`
	CardHolderFirstName                *string `ncpdp:"04CC"`
	CardHolderLastName                 *string `ncpdp:"04CD"`
	HomePlan                           *string `ncpdp:"04CE"`
	PlanId                             *string `ncpdp:"04FO"`
	EligibilityClarificationCode       *string `ncpdp:"04C9"`
	GroupID                            *string `ncpdp:"04C1"`
	PersonCode                         *string `ncpdp:"04C3"`
	RelationshipCode                   *string `ncpdp:"04C6"`
	OtherPayerBin                      *string `ncpdp:"04MG"`
	OtherPayerPCN                      *string `ncpdp:"04MH"`
	OtherPayerCardholderId             *string `ncpdp:"04NU"`
	OtherPayerGroupId                  *string `ncpdp:"04MJ"`
	MedigapId                          *string `ncpdp:"042A"`
	MedicaidIndicator                  *string `ncpdp:"042B"`
	ProviderAcceptsAssignmentIndicator *string `ncpdp:"042D"`
	CmsPartDDefinedQualifiedFacility   *string `ncpdp:"04G2"`
	MedicaidIdNumber                   *string `ncpdp:"04N5"`
	MedicaidAgencyNumber               *string `ncpdp:"04N6"`
}

type Patient struct {
	SegmentID          string  `ncpdp:"01AM"`
	PatientIDQualifier *string `ncpdp:"01CX"`
	PatientID          *string `ncpdp:"01CY"`
	DateOfBirth        *string `ncpdp:"01C4"`
	GenderCode         *string `ncpdp:"01C5"`
	FirstName          *string `ncpdp:"01CA"`
	LastName           *string `ncpdp:"01CB"`
	StreetAddress      *string `ncpdp:"01CM"`
	City               *string `ncpdp:"01CN"`
	State              *string `ncpdp:"01CO"`
	ZipCode            *string `ncpdp:"01CP"`
	PhoneNumber        *string `ncpdp:"01CQ"`
	LocationCode       *string `ncpdp:"01C7"`
	EmployerID         *string `ncpdp:"01CZ"`
	Smoker             *string `ncpdp:"011C"`
	PregnancyIndicator *string `ncpdp:"012C"`
	Email              *string `ncpdp:"01HN"`
	Residence          *string `ncpdp:"014X"`
}
type Claim struct {
	Header    *Header    `ncpdp:"0:header"`
	Insurance *Insurance `ncpdp:"0:04"`
	Patient   *Patient   `ncpdp:"0:01"`
}

func TestUnmarshal(t *testing.T) {
	testNCPDPMessage := "016904D0B1SMAR0     1011326148644     20220722          \u001E\u001CAM04\u001CC24WR9AP4KV57\u001CCCDANNY\u001CCDHAZELWOOD\u001CC90\u001CC3001\u001CC61\u001E\u001CAM01\u001CC419510521\u001CC51\u001CCADANNY\u001CCBHAZELWOOD\u001CCM15 PONDEROSA DR\u001CCNLONOKE\u001CCOAR\u001CCP72086\u001CCQ8702670944\u001C4X0\u001D\u001E\u001CAM07\u001CEM1\u001CD26249119\u001CE103\u001CD780777027310\u001CE7250\u001CD300\u001CD51\u001CD61\u001CD80\u001CDE20220722\u001CDJ3\u001CC81\u001CDT0\u001C28EA\u001CU799\u001E\u001CAM11\u001CD9A\u001CDC400{\u001CE3400{\u001CDQ400A\u001CDU800A\u001CDN07\u001E\u001CAM03\u001CEZ01\u001CDB1942896055\u001CDRRABB\u001C2JNOAH\u001C2MCABOT\u001C2NAR\u001C2P72023\u001E\u001CAM13\u001CVE1\u001CWE02"

	testCLM := Claim{}

	err := Unmarshal([]byte(testNCPDPMessage), &testCLM)
	if err != nil {
		log.Panic(err)
	}

}

func TestParseNCPDPCurrencyStringNegatives(t *testing.T) {
	//N
	gott, err := parseNCPDPCurrencyString("999N", 32)
	if err != nil {
		log.Panic(err)
	}
	got := *gott
	want := Currency64(-99.95)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//}
	gott, err = parseNCPDPCurrencyString("999}", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.90)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//J
	gott, err = parseNCPDPCurrencyString("999J", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.91)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//K
	gott, err = parseNCPDPCurrencyString("999K", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.92)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//L
	gott, err = parseNCPDPCurrencyString("999L", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.93)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//M
	gott, err = parseNCPDPCurrencyString("999M", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.94)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//O
	gott, err = parseNCPDPCurrencyString("999O", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.96)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//P
	gott, err = parseNCPDPCurrencyString("999P", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.97)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//Q
	gott, err = parseNCPDPCurrencyString("999Q", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.98)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//R
	gott, err = parseNCPDPCurrencyString("999R", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(-99.99)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRenderingCurrency(t *testing.T) {
	originalNCPDP := "999E"
	original, err := parseNCPDPCurrencyString("999E", 32)
	if err != nil {
		log.Panic(err)
	}
	newVal := renderNCPDPCurrencyString(*original)

	if originalNCPDP != newVal {
		t.Errorf("Expected %s but got %s", originalNCPDP, newVal)
	}

	originalNCPDP = "999R"
	original, err = parseNCPDPCurrencyString(originalNCPDP, 32)
	if err != nil {
		log.Panic(err)
	}
	newVal = renderNCPDPCurrencyString(*original)

	if originalNCPDP != newVal {
		t.Errorf("Expected %s but got %s", originalNCPDP, newVal)
	}

}

func TestCurrencyParsingAndRenderingAllPossibilities(t *testing.T) {

	for i := 0; i < 10; i++ {

		val := 10.00 + float64(i)*0.01
		log.Printf("Testing %.2f", val)
		err := checkRenderToAndParseFromNCPDPCurrency(Currency64(val))
		if err != nil {
			t.Errorf("error while testing currencyRenderParse %v", err)
		}

		val = val * -1
		log.Printf("Testing %.2f", val)
		err = checkRenderToAndParseFromNCPDPCurrency(Currency64(val))
		if err != nil {
			t.Errorf("error while testing currencyRenderParse %v", err)
		}

	}

}

func checkRenderToAndParseFromNCPDPCurrency(val Currency64) error {

	ncpdpFormat := renderNCPDPCurrencyString(val)
	parsedVal, err := parseNCPDPCurrencyString(ncpdpFormat, 64)
	if err != nil {
		return err
	}
	if val != *parsedVal {
		return fmt.Errorf("unable to transform original val %.2f to ncpdpFormat %s and the when we transformed back we got %.2f ", val, ncpdpFormat, *parsedVal)
	}

	return nil

}

func TestParseNCPDPCurrencyStringPositives(t *testing.T) {
	//N
	gott, err := parseNCPDPCurrencyString("999E", 32)
	if err != nil {
		log.Panic(err)
	}
	got := *gott
	want := Currency64(99.95)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//}
	gott, err = parseNCPDPCurrencyString("999{", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.90)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//J
	gott, err = parseNCPDPCurrencyString("999A", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.91)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//K
	gott, err = parseNCPDPCurrencyString("999B", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.92)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//L
	gott, err = parseNCPDPCurrencyString("999C", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.93)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//M
	gott, err = parseNCPDPCurrencyString("999D", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.94)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//O
	gott, err = parseNCPDPCurrencyString("999F", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.96)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//P
	gott, err = parseNCPDPCurrencyString("999G", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.97)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//Q
	gott, err = parseNCPDPCurrencyString("999H", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.98)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	//R
	gott, err = parseNCPDPCurrencyString("999I", 32)
	if err != nil {
		log.Panic(err)
	}
	got = *gott
	want = Currency64(99.99)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
