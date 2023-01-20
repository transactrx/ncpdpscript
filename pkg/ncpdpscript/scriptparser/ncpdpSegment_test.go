package scriptparser

import (
	"log"
	"testing"
)

var segTestData string = SegmentSeparator + FieldSeparator + "AM04\u001CC24WR9AP4KV57\u001CCCDANNY\u001CCDHAZELWOOD\u001CC90\u001CC3001\u001CC61"
var newSeg, segInitError = newSegment([]byte(segTestData))

func TestSegmentId(t *testing.T) {
	if segInitError != nil {
		log.Panic(segInitError)
	}
	want := "04"
	got := newSeg.id
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSegment_GetFieldById(t *testing.T) {

	if segInitError != nil {
		log.Panic(segInitError)
	}
	c2Field := newSeg.GetFieldById("C2")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want := "4WR9AP4KV57"
	got := c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	c2Field = newSeg.GetFieldById("CC")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want = "DANNY"
	got = c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	c2Field = newSeg.GetFieldById("CD")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want = "HAZELWOOD"
	got = c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	c2Field = newSeg.GetFieldById("C9")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want = "0"
	got = c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	c2Field = newSeg.GetFieldById("C3")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want = "001"
	got = c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	c2Field = newSeg.GetFieldById("C6")
	if c2Field == nil {
		t.Errorf("getFieldById fail, expected firled c2, but got nil")
	}
	want = "1"
	got = c2Field.fieldData
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestSegment_String(t *testing.T) {
	if segInitError != nil {
		log.Panic(segInitError)
	}
	want := segTestData
	got, err := newSeg.String()
	if err != nil {
		log.Panic(err)
	}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
