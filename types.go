package datev

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"
)

type Int int

func (i Int) String() string {
	return fmt.Sprint(int(i))
}

type Decimal float64

func (d Decimal) String() string {
	return fmt.Sprint(float64(d))
}

type Date struct {
	time.Time
}

// used for csv export
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) String() string {
	return d.Time.Format("20060102")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try fivaldi date format
	d.Time, err = time.Parse("20060102", value)
	return err
}

type Time struct {
	time.Time
}

type Amount float64

func (a Amount) String() string {
	s := fmt.Sprintf("%.2f", a)
	return s
}

func (a Amount) Abs() Amount {
	aa := math.Abs(float64(a))
	return Amount(aa)
}

// used for csv export
func (a Amount) MarshalText() ([]byte, error) {
	s := a.String()
	s = strings.Replace(a.String(), ".", ",", -1)
	return []byte(s), nil
}

type Bool bool

type Year int
