package datev_test

import (
	"testing"
	"time"

	datev "github.com/omniboost/go-datev"
)

func TestMediumDateString(t *testing.T) {
	if got := (datev.MediumDate{}).String(); got != "" {
		t.Errorf("zero date: got %q, want empty", got)
	}
	d := datev.MediumDate{Time: time.Date(2026, 10, 12, 0, 0, 0, 0, time.UTC)}
	if got := d.String(); got != "12102026" {
		t.Errorf("got %q, want 12102026", got)
	}
}
