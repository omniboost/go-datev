package datev_test

import (
	"log"
	"testing"

	datev "github.com/omniboost/go-datev"
)

func TestCSVBookingLine(t *testing.T) {
	var errs []error

	h := datev.CSVHeaderLine{
		WKZ: "OO",
	}

	errs = h.Validate()
	if len(errs) > 0 {
		t.Error(errs[0])
	}
	log.Println(h.StringValues())

	l := datev.CSVBookingLine{
		Umsatz: 123.456789,
	}

	errs = l.Validate()
	if len(errs) > 0 {
		t.Error(errs[0])
	}
	log.Println(l.StringValues())
}
