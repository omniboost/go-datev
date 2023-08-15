package datev

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Int int

func (i *Int) String() string {
	if i == nil {
		return ""
	}

	return fmt.Sprint(int(*i))
}

type Decimal float64

func (d Decimal) String() string {
	if float64(d) == 0 {
		return ""
	}
	s := fmt.Sprintf("%.2f", float64(d))
	return strings.Replace(s, ".", ",", -1)
}

type Date struct {
	time.Time
}

func (d *Date) String() string {
	if d == nil {
		return ""
	}
	if d.IsZero() {
		return ""
	}
	return d.Time.Format("20060102")
}

// used for csv export
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
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

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try datev date format
	d.Time, err = time.Parse("20060102", value)
	return err
}

type ShortDate struct {
	time.Time
}

func (d ShortDate) String() string {
	return d.Time.Format("0201")
}

func (d *ShortDate) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
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

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try datev date format
	d.Time, err = time.Parse("0201", value)
	return err
}

type MediumDate struct {
	time.Time
}

func (d MediumDate) String() string {
	return d.Time.Format("02012006")
}

func (d *MediumDate) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
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

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try datev date format
	d.Time, err = time.Parse("02012006", value)
	if err == nil {
		return nil
	}
	return err
}

type Time struct {
	time.Time
}

func (t *Time) String() string {
	if t == nil {
		return ""
	}

	return t.Format("20060102150405000")
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	t.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try datev time format
	t.Time, err = time.Parse("20060102150405000", value)
	return err
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

func (b *Bool) String() string {
	if b == nil {
		return ""
	}

	if *b {
		return "1"
	}
	return "0"
}

type Year int

func (y *Year) String() string {
	if y == nil {
		return ""
	}
	return strconv.Itoa(int(*y))
}
