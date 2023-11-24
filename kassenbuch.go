package datev

import "fmt"

type Kassenbuch []KassenbuchLine

type KassenbuchLine struct {
	// Währungskennzeichen
	// Nur EUR zulässig
	Wahrung string // 1

	// Vorzeichen und Betrag
	// Vorzeichen: + oder –
	// Betrag maximale Länge: 10,2
	// Numerisch
	// Trennzeichen Komma (,)
	// Keine Tausendertrennzeichen
	// Wert >0
	VorzBetrag NilDecimal // 2

	// Beleg-Nr.
	// Maximale Länge: 36 Zeichen
	// Alphanumerisch
	// Sonderzeichen: $ & % * + - /
	RechNr string // 3

	// Datum der Kassenbewegungen
	// Format: TTMM
	// Maximale Länge: 4 Zeichen
	// Numerisch
	// Keine Trennzeichen
	BelegDatum string // 4

	// Belegtext
	// Maximale Länge: 30 Zeichen
	// Alphanumerisch & Sonderzeichen
	Belegtext string // 5

	// Umsatzsteuersatz in %
	// Maximale Länge: 2 Vorkomma-, 2
	// Nachkommastellen
	// Numerisch
	// Dezimalpunkt = Komma (,)
	UStSatz NilDecimal // 6

	// Berichtigungsschlüssel (B)
	// Umsatzsteuerschlüssel (U)
	// Maximale Länge: 4 Zeichen
	// Numerisch
	// Gültige BU-Schlüssel finden Sie im
	// Dokument: DATEV Unternehmen
	// online - Übersicht Steuerschlüssel
	// (BU)
	BU string // 7

	// Gegenkonto
	// Maximale Länge: Sachkontolänge +1
	// (zum Beispiel bei SKL 4 maximal 5
	// Zeichen)
	// Numerisch
	Gegenkonto *Int // 8

	// KOST1, Informationen für die Kostenrechnung, zum Beispiel Kostenstellen
	// Maximale Länge: 36 Zeichen
	// Alphanumerisch & Sonderzeichen
	Kost1 string // 9

	// KOST2, Informationen für die Kostenrechnung, zum Beispiel Kostenstellen
	// Maximale Länge: 36 Zeichen
	// Alphanumerisch & Sonderzeichen
	Kost2 string // 10

	// Wertangabe zu Bezugsgröße in Kostenrechnung, zum Beispiel kg, cm, %
	// Maximale Länge: 12,4
	// Numerische
	// Dezimalpunkt = Komma (,)
	// Keine Tausendertrennzeichen
	// Wert >0
	Kostmenge NilDecimal // 11

	// Skontobetrag in EUR
	// Maximale Länge: 8,2
	// Numerisch
	// Dezimalpunkt = Komma (,)
	// Wert >0
	Skonto NilDecimal // 12

	// Nachricht
	// Maximale Länge: 120 Zeichen
	// Alphanumerisch & Sonderzeichen
	Nachricht string
}

func (l KassenbuchLine) Validate() []error {
	var errs []error
	return errs
}

func (e KassenbuchLine) Headers() []string {
	return []string{
		"Währung",
		"VorzBetrag",
		"RechNr",
		"BelegDatum",
		"Belegtext",
		"UStSatz",
		"BU",
		"Gegenkonto",
		"Kost1",
		"Kost2",
		"Kostmenge",
		"Skonto",
		"Nachricht",
	}
}

func (l KassenbuchLine) Values() []interface{} {
	return []interface{}{
		l.Wahrung,    // 1
		l.VorzBetrag, // 2
		l.RechNr,     // 3
		l.BelegDatum, // 4
		l.Belegtext,  // 5
		l.UStSatz,    // 6
		l.BU,         // 7
		l.Gegenkonto, // 8
		l.Kost1,      // 9
		l.Kost2,      // 10
		l.Kostmenge,  // 11
		l.Skonto,     // 12
		l.Nachricht,  // 13
	}
}

func (l KassenbuchLine) StringValues() []string {
	values := l.Values()
	st := make([]string, len(values))
	for k, v := range values {
		st[k] = fmt.Sprint(v)
	}
	return st
}
