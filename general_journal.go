package datev

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

func NewCSVHeaderLine() *CSVHeaderLine {
	return &CSVHeaderLine{
		DATEVFormatKZ:   "EXTF",
		Versionsnummer:  510,
		Datenkategorie:  21,
		Formatname:      "Buchungsstapel",
		Formatversion:   7,
		ErzeugtAm:       &Time{time.Now()},
		Sachkontenlange: 4,
		Bezeichnung:     "Buchungen",
		Buchungstyp:     1,
		WKZ:             "EUR",
	}
}

// http://www.datev.de/dnlexom/client/app/index.html#/document/1036228/D103622800010
type CSVHeaderLine struct {
	// vom Datev angegeben
	// EXTF = für Datei-Formate, die von externen Programmen erstellt wurden
	DATEVFormatKZ string // 1

	// entspricht der zugrundeliegenden Versionsnummer des Scnittstellen-Entwicklungsleitfadens
	Versionsnummer int // 2

	// vom Datev angegeben
	Datenkategorie int // 3

	// vom Datev angegeben
	Formatname string // 4

	// vom Datev angegeben
	Formatversion int // 5

	ErzeugtAm *Time // 6

	// Darf nicht gefüllt werden, durch Import gesetzt.
	Importiert *Date // 7

	// Herkunfts-Kennzeichen
	// Beim Import wird das Herkunfts-Kennzeichen durch „SV“ (= Stapelverarbeitung) ersetzt.
	Herkunft string // 8

	// Benutzername
	ExportiertVon string // 9

	// Darf nicht gefüllt werden, durch Import gesetzt.
	ImportiertVon string // 10

	// Beraternummer
	Berater int // 11

	// Mandantennummer
	Mandant int // 12

	// Wirtschaftsjahresbeginn
	WJBeginn *Date // 13

	// Kleinste Sachkontenlänge = 4, Grösste Sachkontenlänge = 8
	Sachkontenlange int // 14

	DatumVom *Date // 15

	DatumBis *Date // 16

	// Bezeichnung des Buchungsstapels
	Bezeichnung string // 17

	// Diktatkürzel Beispiel MM = Max Mustermann
	Diktatkurzel string // 18

	// 1 = Finanzbuchhaltung, 2 = Jahresabschluss
	Buchungstyp int // 19

	// 0 oder leer = Rechnungslegungszweckunabhängig
	Rechnungslegungszweck int // 20

	// leer = nicht definiert; wird ab Jahreswechselversion 2016/2017 automatisch festgeschrieben
	// 0 = keine Festschreibung
	// 1 = Festschreibung
	Festschreibung *Bool // 21

	// Währungskennzeichen
	WKZ string // 22

	Reserviert string // 23

	Derivatskennzeichen string // 24

	Reserviert2 string // 25

	Reserviert3 string // 26

	SKR string // 27

	BranchenlosungID int // 28

	Reserviert4 string // 29
	Reserviert5 string // 30

	// Verarbeitungskennzeichen der abgebenden Anwendung => Bsp. '9/2016'
	Anwendungsinformation string
}

func (l CSVHeaderLine) Validate() []error {
	var errs []error

	if utf8.RuneCountInString(l.DATEVFormatKZ) > 4 {
		errs = append(errs, errors.New("DATEVFormatKZ has a maximum length of 4"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Versionsnummer)) > 3 {
		errs = append(errs, errors.New("Versionsnummer has a maximum length of 3"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Datenkategorie)) > 2 {
		errs = append(errs, errors.New("Datenkategorie has a maximum length of 2"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Formatversion)) > 3 {
		errs = append(errs, errors.New("Formatversion has a maximum length of 3"))
	}

	if utf8.RuneCountInString(l.Herkunft) > 2 {
		errs = append(errs, errors.New("Herkunft has a maximum length of 2"))
	}

	if utf8.RuneCountInString(l.ExportiertVon) > 25 {
		errs = append(errs, errors.New("ExportiertVon has a maximum length of 25"))
	}

	if utf8.RuneCountInString(l.ImportiertVon) > 10 {
		errs = append(errs, errors.New("ImportiertVon has a maximum length of 10"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Berater)) > 7 {
		errs = append(errs, errors.New("Berater has a maximum length of 7"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Mandant)) > 5 {
		errs = append(errs, errors.New("Mandant has a maximum length of 5"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Sachkontenlange)) > 1 {
		errs = append(errs, errors.New("Sachkontenlange has a maximum length of 1"))
	}

	if utf8.RuneCountInString(l.Bezeichnung) > 30 {
		errs = append(errs, errors.New("Bezeichnung has a maximum length of 30"))
	}

	if utf8.RuneCountInString(l.Diktatkurzel) > 2 {
		errs = append(errs, errors.New("Diktatkurzel has a maximum length of 2"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Buchungstyp)) > 1 {
		errs = append(errs, errors.New("Buchungstyp has a maximum length of 1"))
	}

	if utf8.RuneCountInString(fmt.Sprint(l.Rechnungslegungszweck)) > 2 {
		errs = append(errs, errors.New("Rechnungslegungszweck has a maximum length of 2"))
	}

	if utf8.RuneCountInString(l.WKZ) > 3 {
		errs = append(errs, errors.New("WKZ has a maximum length of 3"))
	}

	if utf8.RuneCountInString(l.Anwendungsinformation) > 16 {
		errs = append(errs, errors.New("Anwendungsinformation a maximum length of 16"))
	}

	return errs
}

func (e CSVHeaderLine) Headers() []string {
	return []string{
		"DATEV-Format-KZ",
		"Versionsnummer",
		"Datenkategorie",
		"Formatname",
		"Formatversion",
		"Erzeugt am",
		"Importiert",
		"Herkunft",
		"Exportiert von",
		"Importiert von",
		"Berater",
		"Mandant",
		"WJ-Beginn",
		"Sachkontenlänge",
		"Datum vom",
		"Datum bis",
		"Bezeichnung",
		"Diktatkürzel",
		"Buchungstyp",
		"Rechnungslegungszweck",
		"Festschreibung",
		"WKZ",
		"reserviert",
		"Derivatskennzeichen",
		"reserviert 2",
		"reserviert 3",
		"SKR",
		"Branchenlösung-Id",
		"reserviert 4",
		"reserviert 5",
		"Anwendungsinformation",
	}
}

func (l CSVHeaderLine) Values() []interface{} {
	return []interface{}{
		l.DATEVFormatKZ,
		l.Versionsnummer,
		l.Datenkategorie,
		l.Formatname,
		l.Formatversion,
		l.ErzeugtAm,
		l.Importiert,
		l.Herkunft,
		l.ExportiertVon,
		l.ImportiertVon,
		l.Berater,
		l.Mandant,
		l.WJBeginn,
		l.Sachkontenlange,
		l.DatumVom,
		l.DatumBis,
		l.Bezeichnung,
		l.Diktatkurzel,
		l.Buchungstyp,
		l.Rechnungslegungszweck,
		l.Festschreibung,
		l.WKZ,
		l.Reserviert,
		l.Derivatskennzeichen,
		l.Reserviert2,
		l.Reserviert3,
		l.SKR,
		l.BranchenlosungID,
		l.Reserviert4,
		l.Reserviert5,
		l.Anwendungsinformation,
	}
}

func (l CSVHeaderLine) StringValues() []string {
	values := l.Values()
	st := make([]string, len(values))
	for k, v := range values {
		st[k] = fmt.Sprint(v)
	}
	return st
}

type CSVBookingLine struct {
	// Beispiel: 1234567890,12
	// Muss immer ein positiver Wert sein
	// required
	Umsatz Decimal // 1

	// Soll-/Haben-Kennzeichnung des Umsatzes bezieht sich auf das Konto, das im Feld Konto angegeben wird:
	// S = Soll
	// H = Haben
	// required
	SollHaben string // 2

	// ISO-Code der Währung (Dok.-Nr. 1080170); gibt an, welche Währung dem Betrag zugrunde liegt.
	// Wenn kein Wert angegeben ist, wird das WKZ aus dem Header übernommen.
	WKZUmsatz string // 3

	// Der Fremdwährungskurs bestimmt, wie der angegebene Umsatz, der in Fremdwährung übergeben wird, in die Basiswährung umzurechnen ist, wenn es sich um ein Nicht-EWU-Land handelt.
	// Beispiel: 1123,123456
	// Achtung: Der Wert 0 ist unzulässig.
	Kurs Decimal // 4

	// Wenn das Feld Basisumsatz verwendet wird, muss auch das Feld WKZ Basisumsatz gefüllt werden.
	// Beispiel: 1123123123,12
	Basisumsatz Decimal // 5

	// Währungskennzeichen der hinterlegten Basiswährung. Wenn das Feld WKZ Basisumsatz verwendet wird, muss auch das Feld Basisumsatz verwendet werden.
	// ISO-Code beachten (siehe Dok.-Nr.1080170)
	WKZBasisumsatz string // 6

	// Sach- oder Personen-Kontonummer Darf max. 8- bzw. max. 9-stellig sein
	// (abhängig von der Information im Header) Die Personenkontenlänge darf nur
	// 1 Stelle länger sein als die definierte Sachkontennummernlänge.
	Konto Int // 7

	// Sach- oder Personen-Kontonummer Darf max. 8- bzw. max. 9-stellig sein
	// (abhängig von der Information im Header) Die Personenkontenlänge darf nur
	// 1 Stelle länger sein als die definierte Sachkontennummernlänge.
	Gegenkonto Int // 8

	// Steuerschlüssel und/oder Berichtigungsschlüssel
	BUSchlussel string // 9

	// Achtung: Auch bei individueller Feldformatierung mit vierstelliger
	// Jahreszahl wird immer in das aktuelle Wirtschaftsjahr importiert, wenn
	// Tag und Monat des Datums im bebuchbaren Zeitraum liegen, da die
	// Jahreszahl nicht berücksichtigt wird.
	Belegdatum ShortDate // 10

	// Rechnungs-/Belegnummer
	// Das Belegfeld 1 ist der "Schlüssel" für die Verwaltung von Offenen Posten.
	// Bei einer Zahlung oder Gutschrift erfolgt nur dann ein OP-Ausgleich, wenn die Belegnummer mit dem Belegfeld 1 identisch ist.
	Belegfeld1 string // 11

	// Belegnummer oder OPOS-Verarbeitungsinformationen
	Belegfeld2 string // 12

	// Skonto-Betrag/-Abzug
	// Nur bei Zahlungen zulässig.
	// Beispiel 12123123,12
	// Achtung: Der Wert 0 ist unzulässig.
	Skonto Decimal // 13

	Buchungstext string // 14

	// Mahn-/Zahl-Sperre
	// Die Rechnung kann aus dem Mahnwesen / Zahlungsvorschlag ausgeschlossen werden.
	// true = Postensperre
	// false/keine Angabe = keine Sperre
	// Nur in Verbindung mit einer Rechnungsbuchung und Personenkonto (OPOS) relevant.
	Postensperre *Bool // 15

	// Adressnummer einer diversen Adresse
	// Nur in Verbindung mit OPOS relevant.
	DiverseAdressnummer string // 16

	// Wenn für eine Lastschrift oder Überweisung eine bestimmte Bank des Geschäftspartners genutzt werden soll.
	// Nur in Verbindung mit OPOS relevant.
	Geschaftspartnerbank *Int // 17

	// Der Sachverhalt wird in Rechnungswesen pro verwendet, um Buchungen/Posten als Mahnzins/Mahngebühr zu identifizieren.
	// Für diese Posten werden z. B. beim Erstellen von Mahnungen keine Mahnzinsen berechnet.
	// 31 = Mahnzins
	// 40 = Mahngebühr
	// Nur in Verbindung mit OPOS relevant.
	Sachverhalt *Int // 18

	// Hier kann eine Zinssperre übergeben werden; dadurch werden für diesen Posten bei Erstellung einer Mahnung keine Mahnzinsen ermittelt.
	// Nur in Verbindung mit OPOS relevant.
	// keine Angabe und 0 = keine Sperre
	// 1 = Zinssperre
	Zinssperre *Int // 19

	// Link auf den Buchungsbeleg, der digital in einem Dokumenten-Management-System (z. B. DATEV Dokumentenablage, DATEV DMS classic) abgelegt wurde.
	// Der Beleglink hat folgenden Aufbau:
	//   4-stelliges Kürzel für Dokumentenmanagementsystem (siehe unten)
	//   Leerzeichen
	//   Anführungszeichen
	//   Beleglink (GUID, Dateiname des Belegs), max. 36 Zeichen
	//   Schlusszeichen
	// Beispiel für einen Beleglink aus Belege online:
	// BEDI “CB6A8F8F-099A-B3A9-2BAA-0CB64E299BA”
	// Das Kürzel bezeichnet das Quellsystem (Dokumentenmanagement), indem die digitalen Belege abgelegt sind.
	// DATEV verwendet für seine Dokumentenmanagement-Systeme folgende Kürzel:
	//   Belegverwaltung online → BEDI
	//   DATEV DMS → DDMS
	//   Dokumentenablage → DDMS (vormals DORG)
	Beleglink string // 20

	// Bei einem ASCII-Format, das aus einem DATEV pro-Rechnungswesen-Programm erstellt wurde, können diese Felder Informationen aus einem Beleg (z. B. einem elektronischen Kontoumsatz) enthalten.
	// Wenn die Feldlänge eines Beleginfo-Inhalts-Felds überschritten wird, wird die Information im nächsten Beleginfo-Feld weitergeführt.
	// Wichtiger Hinweis
	// Eine Beleginfo besteht immer aus den Bestandteilen Beleginfo-Art und Beleginfo-Inhalt. Wenn Sie die Beleginfo nutzen möchten, befüllen Sie immer beide Felder.
	// Beispiel:
	// Beleginfo-Art:
	// Kontoumsätze der jeweiligen Bank
	// Beleginfo-Inhalt:
	// Buchungsspezifische Inhalte zu den oben genannten Informationsarten
	BeleginfoArt1    string // 21
	BeleginfoInhalt1 string // 22
	BeleginfoArt2    string // 23
	BeleginfoInhalt2 string // 24
	BeleginfoArt3    string // 25
	BeleginfoInhalt3 string // 26
	BeleginfoArt4    string // 27
	BeleginfoInhalt4 string // 28
	BeleginfoArt5    string // 29
	BeleginfoInhalt5 string // 30
	BeleginfoArt6    string // 31
	BeleginfoInhalt6 string // 32
	BeleginfoArt7    string // 33
	BeleginfoInhalt7 string // 34
	BeleginfoArt8    string // 35
	BeleginfoInhalt8 string // 36

	// Über KOST1 erfolgt die Zuordnung des Geschäftsvorfalls für die anschließende Kostenrechnung.
	KOST1Kostenstelle string // 37

	// Über KOST2 erfolgt die Zuordnung des Geschäftsvorfalls für die anschließende Kostenrechnung.
	KOST2Kostenstelle string // 38

	// Im KOST-Mengenfeld wird die Wertgabe zu einer bestimmten Bezugsgröße für eine Kostenstelle erfasst. Diese Bezugsgröße kann z. B. kg, g, cm, m, % sein. Die Bezugsgröße ist definiert in den Kostenrechnungs-Stammdaten.
	// Beispiel: 123123123,12
	KostMenge Decimal // 39

	// Die USt-IdNr. besteht aus:
	// 2-stelligen Länderkürzel (siehe Dok.-Nr. 1080169; Ausnahme Griechenland: Das Länderkürzel lautet EL)
	// 13-stelliger USt-IdNr.
	// VAT Number
	UStIDNR string // 40

	// Nur für entsprechende EU-Buchungen:
	// Der im EU-Bestimmungsland gültige Steuersatz.
	// Beispiel: 12,12
	EUSteuersatz Decimal // 41

	// Für Buchungen, die in einer von der Mandantenstammdaten-Schlüsselung abweichenden Umsatzsteuerart verarbeitet werden sollen, kann die abweichende Versteuerungsart im Buchungssatz übergeben werden:
	// I = Ist-Versteuerung
	// K = keine Umsatzsteuerrechnung
	// P = Pauschalierung (z. B. für Land- und Forstwirtschaft)
	// S = Soll-Versteuerung
	ABBVersteuerungsart string // 42

	// Sachverhalte gem. § 13b Abs. 1 Satz 1 Nrn. 1.ff UStG
	// Achtung: Der Wert 0 ist unzulässig.
	// (siehe Dok.-Nr. 1034915)
	SachverhaltLL *Int // 43

	// Steuersatz/Funktion zum L+L-Sachverhalt
	// Achtung: Der Wert 0 ist unzulässig.
	// (siehe Dok.-Nr. 1034915)
	Funktionserganzung *Int // 44

	// Bei Verwendung des BU-Schlüssels 49 für "andere Steuersätze" muss der steuerliche Sachverhalt mitgegeben werden.
	BU49Hauptfunktionstyp *Int // 45

	BU49Hauptfunktionsnummer *Int // 46

	BU49Funktionsergänzung *Int // 47

	// Zusatzinformationen, die zu Buchungssätzen erfasst werden können.
	// Diese Zusatzinformationen besitzen den Charakter eines Notizzettels und können frei erfasst werden.
	// Wichtiger Hinweis
	// Eine Zusatzinformation besteht immer aus den Bestandteilen Informationsart und Informationsinhalt. Wenn Sie die Zusatzinformation nutzen möchten, füllen Sie immer beide Felder.
	// Beispiel:
	// Informationsart, z. B. Filiale oder Mengengrößen (qm)
	// Informationsinhalt: Buchungsspezifische Inhalte zu den oben genannten Informationsarten.
	ZusatzinformationArt1     string // 48
	ZusatzinformationInhalt1  string // 49
	ZusatzinformationArt2     string // 49
	ZusatzinformationInhalt2  string // 50
	ZusatzinformationArt3     string // 50
	ZusatzinformationInhalt3  string // 51
	ZusatzinformationArt4     string // 51
	ZusatzinformationInhalt4  string // 52
	ZusatzinformationArt5     string // 52
	ZusatzinformationInhalt5  string // 53
	ZusatzinformationArt6     string // 53
	ZusatzinformationInhalt6  string // 54
	ZusatzinformationArt7     string // 54
	ZusatzinformationInhalt7  string // 55
	ZusatzinformationArt8     string // 55
	ZusatzinformationInhalt8  string // 56
	ZusatzinformationArt9     string // 56
	ZusatzinformationInhalt9  string // 57
	ZusatzinformationArt10    string // 57
	ZusatzinformationInhalt10 string // 58
	ZusatzinformationArt11    string // 58
	ZusatzinformationInhalt11 string // 59
	ZusatzinformationArt12    string // 59
	ZusatzinformationInhalt12 string // 60
	ZusatzinformationArt13    string // 60
	ZusatzinformationInhalt13 string // 61
	ZusatzinformationArt14    string // 61
	ZusatzinformationInhalt14 string // 62
	ZusatzinformationArt15    string // 62
	ZusatzinformationInhalt15 string // 63
	ZusatzinformationArt16    string // 63
	ZusatzinformationInhalt16 string // 64
	ZusatzinformationArt17    string // 64
	ZusatzinformationInhalt17 string // 65
	ZusatzinformationArt18    string // 65
	ZusatzinformationInhalt18 string // 66
	ZusatzinformationArt19    string // 66
	ZusatzinformationInhalt19 string // 67
	ZusatzinformationArt20    string // 67
	ZusatzinformationInhalt20 string // 68
	ZusatzinformationArt21    string // 68
	ZusatzinformationInhalt21 string // 69
	ZusatzinformationArt22    string // 69
	ZusatzinformationInhalt22 string // 70
	ZusatzinformationArt23    string // 70
	ZusatzinformationInhalt23 string // 71
	ZusatzinformationArt24    string // 71
	ZusatzinformationInhalt24 string // 72
	ZusatzinformationArt25    string // 72
	ZusatzinformationInhalt25 string // 73
	ZusatzinformationArt26    string // 73
	ZusatzinformationInhalt26 string // 74
	ZusatzinformationArt27    string // 74
	ZusatzinformationInhalt27 string // 75
	ZusatzinformationArt28    string // 75
	ZusatzinformationInhalt28 string // 76
	ZusatzinformationArt29    string // 76
	ZusatzinformationInhalt29 string // 77
	ZusatzinformationArt30    string // 77
	ZusatzinformationInhalt30 string // 78
	ZusatzinformationArt31    string // 78
	ZusatzinformationInhalt31 string // 79
	ZusatzinformationArt32    string // 79
	ZusatzinformationInhalt32 string // 80
	ZusatzinformationArt33    string // 80
	ZusatzinformationInhalt33 string // 81
	ZusatzinformationArt34    string // 81
	ZusatzinformationInhalt34 string // 82
	ZusatzinformationArt35    string // 82
	ZusatzinformationInhalt35 string // 83
	ZusatzinformationArt36    string // 83
	ZusatzinformationInhalt36 string // 84
	ZusatzinformationArt37    string // 84
	ZusatzinformationInhalt37 string // 85
	ZusatzinformationArt38    string // 85
	ZusatzinformationInhalt38 string // 86
	ZusatzinformationArt39    string // 86
	ZusatzinformationInhalt39 string // 87

	// Wirkt sich nur bei Sachverhalt mit SKR14 Land- und Forstwirtschaft aus, für andere SKR werden die Felder beim Import/Export überlesen bzw. leer exportiert.
	Stuck *Int // 88

	Gewicht Decimal // 89

	// OPOS-Informationen kommunal
	// 1 = Lastschrift
	// 2 = Mahnung
	// 3 = Zahlung
	Zahlweise *Int // 90

	// OPOS-Informationen kommunal
	Forderungsart string // 91

	// OPOS-Informationen kommunal
	Veranlagungsjahr Year // 92

	// OPOS-Informationen kommunal
	ZugeordneteFalligkeit Date // 93

	// 1 = Einkauf von Waren
	// 2 = Erwerb von Roh-Hilfs- und Betriebsstoffen
	Skontotyp *Int // 94

	// Allgemeine Bezeichnung, des Auftrags/Projekts
	Auftragsnummer string // 95

	// AA = Angeforderte Anzahlung/Abschlagsrechnung
	// AG = Erhaltene Anzahlung (Geldeingang)
	// AV = Erhaltene Anzahlung (Verbindlichkeit)
	// SR = Schlussrechnung
	// SU = Schlussrechnung (Umbuchung)
	// SG = Schlussrechnung (Geldeingang)
	// SO = Sonstige
	Buchungstyp string // 96

	// USt-Schlüssel der späteren Schlussrechnung
	UStSchlüsselAnzahlungen *Int // 97

	// EU-Mitgliedstaat der späteren Schlussrechnung
	// (siehe Dok.-Nr. 1080169)
	EUMitgliedstaatAnzahlungen string // 98

	// L+L-Sachverhalt der späteren Schlussrechnung
	// Sachverhalte gem. § 13b Abs. 1 Satz 1 Nrn. 1.-5. UStG
	// Achtung: Der Wert 0 ist unzulässig.
	SachverhaltLLAnzahlungen *Int // 99

	// EU-Steuersatz der späteren Schlussrechnung
	// Nur für entsprechende EU-Buchungen: Der im EU-Bestimmungsland gültige Steuersatz.
	// Beispiel: 12,12
	EUSteuersatzAnzahlungen Decimal // 100

	// Erlöskonto der späteren Schlussrechnung
	ErloskontoAnzahlungen *Int // 101

	// Wird beim Import durch SV (Stapelverarbeitung) ersetzt.
	HerkunftKz string // 102

	// Wird von DATEV verwendet.
	Leerfeld string // 103

	// Format TTMMJJJJ
	KostDatum Date // 104

	// Vom Zahlungsempfänger individuell vergebenes Kennzeichen eines Mandats (z. B. Rechnungs- oder Kundennummer).
	SEPAMandatsreferenz string // 105

	// 0 = keine Skontosperre
	// 1 = Skontosperre
	Skontosperre Bool // 106

	Gesellschaftername string //107

	Beteiligtennummer *Int // 108

	Identifikationsnummer string // 109

	Zeichnernummer string // 110

	Postensperreis Date // 111

	// SoBil-Sachverhalt
	Bezeichnung string // 112

	// SoBil-Buchung
	Kennzeichen *Int // 113

	// leer = nicht definiert; wird ab Jahreswechselversion 2016/2017 automatisch festgeschrieben
	// 0 = keine Festschreibung
	// 1 = Festschreibung
	// Hat ein Buchungssatz in diesem Feld den Inhalt 1, so wird der gesamte Stapel nach dem Import festgeschrieben.
	// Ab Jahreswechselversion 2016/2017 gilt das auch bei Inhalt = leer.
	Festschreibung Bool // 114

	Leistungsdatum Date // 115

	// Steuerperiode
	DatumZuord Date // 116
}

func (l CSVBookingLine) Validate() []error {
	var errs []error
	return errs
}

func (e CSVBookingLine) Headers() []string {
	return []string{
		"Umsatz (ohne Soll/Haben-Kz)",
		"Soll/Haben-Kennzeichen",
		"WKZ Umsatz",
		"Kurs",
		"Basisumsatz",
		"WKZ Basisumsatz",
		"Konto",
		"Gegenkonto (ohne BU-Schlüssel)",
		"BU-Schlüssel",
		"Belegdatum",
		"Belegfeld 1",
		"Belegfeld 2",
		"Skonto",
		"Buchungstext",
		"Postensperre",
		"Diverse Adressnummer",
		"Geschäftspartnerbank",
		"Sachverhalt",
		"Zinssperre",
		"Beleglink",
		"Beleginfo Art 1",
		"Beleginfo Inhalt 1",
		"Beleginfo Art 2",
		"Beleginfo Inhalt 2",
		"Beleginfo Art 3",
		"Beleginfo Inhalt 3",
		"Beleginfo Art 4",
		"Beleginfo Inhalt 4",
		"Beleginfo Art 5",
		"Beleginfo Inhalt 5",
		"Beleginfo Art 6",
		"Beleginfo Inhalt 6",
		"Beleginfo Art 7",
		"Beleginfo Inhalt 7",
		"Beleginfo Art 8",
		"Beleginfo Inhalt 8",
		"KOST1 Kostenstelle",
		"KOST2 Kostenstelle",
		"Kost Menge",
		"EU-Land u. USt-IdNr.",
		"EU-Steuersatz",
		"Abw. Versteuerungsart",
		"Sachverhalt L+L",
		"Funktionsergänzung L+L",
		"BU 49 Hauptfunktionstyp",
		"BU 49 Hauptfunktionsnummer",
		"BU 49 Funktionsergänzung",
		"Zusatzinformation Art 1",
		"Zusatzinformation Inhalt 1",
		"Zusatzinformation Art 2",
		"Zusatzinformation Inhalt 2",
		"Zusatzinformation Art 3",
		"Zusatzinformation Inhalt 3",
		"Zusatzinformation Art 4",
		"Zusatzinformation Inhalt 4",
		"Zusatzinformation Art 5",
		"Zusatzinformation Inhalt 5",
		"Zusatzinformation Art 6",
		"Zusatzinformation Inhalt 6",
		"Zusatzinformation Art 7",
		"Zusatzinformation Inhalt 7",
		"Zusatzinformation Art 8",
		"Zusatzinformation Inhalt 8",
		"Zusatzinformation Art 9",
		"Zusatzinformation Inhalt 9",
		"Zusatzinformation Art 10",
		"Zusatzinformation Inhalt 10",
		"Zusatzinformation Art 11",
		"Zusatzinformation Inhalt 11",
		"Zusatzinformation Art 12",
		"Zusatzinformation Inhalt 12",
		"Zusatzinformation Art 13",
		"Zusatzinformation Inhalt 13",
		"Zusatzinformation Art 14",
		"Zusatzinformation Inhalt 14",
		"Zusatzinformation Art 15",
		"Zusatzinformation Inhalt 15",
		"Zusatzinformation Art 16",
		"Zusatzinformation Inhalt 16",
		"Zusatzinformation Art 17",
		"Zusatzinformation Inhalt 17",
		"Zusatzinformation Art 18",
		"Zusatzinformation Inhalt 18",
		"Zusatzinformation Art 19",
		"Zusatzinformation Inhalt 19",
		"Zusatzinformation Art 20",
		"Zusatzinformation Inhalt 20",
		"Stück",
		"Gewicht",
		"Zahlweise",
		"Forderungsart",
		"Veranlagungsjahr",
		"Zugeordnete Fälligkeit",
		"Skontotyp",
		"Auftragsnummer",
		"Buchungstyp",
		"USt-Schlüssel (Anzahlungen)",
		"EU-Mitgliedstaat (Anzahlungen)",
		"Sachverhalt L+L (Anzahlungen)",
		"EU-Steuersatz (Anzahlungen)",
		"Erlöskonto (Anzahlungen)",
		"Herkunft-Kz",
		"Leerfeld",
		"KOST-Datum",
		"SEPA-Mandatsreferenz",
		"Skontosperre",
		"Gesellschaftername",
		"Beteiligtennummer",
		"Identifikationsnummer",
		"Zeichnernummer",
		"Postensperre bis",
		"Bezeichnung",
		"Kennzeichen",
		"Festschreibung",
		"Leistungsdatum",
		"Datum Zuord.",
	}
}

func (l CSVBookingLine) Values() []interface{} {
	return []interface{}{
		l.Umsatz,
		l.SollHaben,
		l.WKZUmsatz,
		l.Kurs,
		l.Basisumsatz,
		l.WKZBasisumsatz,
		l.Konto,
		l.Gegenkonto,
		l.BUSchlussel,
		l.Belegdatum,
		l.Belegfeld1,
		l.Belegfeld2,
		l.Skonto,
		l.Buchungstext,
		l.Postensperre,
		l.DiverseAdressnummer,
		l.Geschaftspartnerbank,
		l.Sachverhalt,
		l.Zinssperre,
		l.Beleglink,
		l.BeleginfoArt1,
		l.BeleginfoInhalt1,
		l.BeleginfoArt2,
		l.BeleginfoInhalt2,
		l.BeleginfoArt3,
		l.BeleginfoInhalt3,
		l.BeleginfoArt4,
		l.BeleginfoInhalt4,
		l.BeleginfoArt5,
		l.BeleginfoInhalt5,
		l.BeleginfoArt6,
		l.BeleginfoInhalt6,
		l.BeleginfoArt7,
		l.BeleginfoInhalt7,
		l.BeleginfoArt8,
		l.BeleginfoInhalt8,
		l.KOST1Kostenstelle,
		l.KOST2Kostenstelle,
		l.KostMenge,
		l.UStIDNR,
		l.EUSteuersatz,
		l.ABBVersteuerungsart,
		l.SachverhaltLL,
		l.Funktionserganzung,
		l.BU49Hauptfunktionstyp,
		l.BU49Hauptfunktionsnummer,
		l.BU49Funktionsergänzung,
		l.ZusatzinformationArt1,
		l.ZusatzinformationInhalt1,
		l.ZusatzinformationArt2,
		l.ZusatzinformationInhalt2,
		l.ZusatzinformationArt3,
		l.ZusatzinformationInhalt3,
		l.ZusatzinformationArt4,
		l.ZusatzinformationInhalt4,
		l.ZusatzinformationArt5,
		l.ZusatzinformationInhalt5,
		l.ZusatzinformationArt6,
		l.ZusatzinformationInhalt6,
		l.ZusatzinformationArt7,
		l.ZusatzinformationInhalt7,
		l.ZusatzinformationArt8,
		l.ZusatzinformationInhalt8,
		l.ZusatzinformationArt9,
		l.ZusatzinformationInhalt9,
		l.ZusatzinformationArt10,
		l.ZusatzinformationInhalt10,
		l.ZusatzinformationArt11,
		l.ZusatzinformationInhalt11,
		l.ZusatzinformationArt12,
		l.ZusatzinformationInhalt12,
		l.ZusatzinformationArt13,
		l.ZusatzinformationInhalt13,
		l.ZusatzinformationArt14,
		l.ZusatzinformationInhalt14,
		l.ZusatzinformationArt15,
		l.ZusatzinformationInhalt15,
		l.ZusatzinformationArt16,
		l.ZusatzinformationInhalt16,
		l.ZusatzinformationArt17,
		l.ZusatzinformationInhalt17,
		l.ZusatzinformationArt18,
		l.ZusatzinformationInhalt18,
		l.ZusatzinformationArt19,
		l.ZusatzinformationInhalt19,
		l.ZusatzinformationArt20,
		l.ZusatzinformationInhalt20,
		l.ZusatzinformationArt21,
		l.ZusatzinformationInhalt21,
		l.ZusatzinformationArt22,
		l.ZusatzinformationInhalt22,
		l.ZusatzinformationArt23,
		l.ZusatzinformationInhalt23,
		l.ZusatzinformationArt24,
		l.ZusatzinformationInhalt24,
		l.ZusatzinformationArt25,
		l.ZusatzinformationInhalt25,
		l.ZusatzinformationArt26,
		l.ZusatzinformationInhalt26,
		l.ZusatzinformationArt27,
		l.ZusatzinformationInhalt27,
		l.ZusatzinformationArt28,
		l.ZusatzinformationInhalt28,
		l.ZusatzinformationArt29,
		l.ZusatzinformationInhalt29,
		l.ZusatzinformationArt30,
		l.ZusatzinformationInhalt30,
		l.ZusatzinformationArt31,
		l.ZusatzinformationInhalt31,
		l.ZusatzinformationArt32,
		l.ZusatzinformationInhalt32,
		l.ZusatzinformationArt33,
		l.ZusatzinformationInhalt33,
		l.ZusatzinformationArt34,
		l.ZusatzinformationInhalt34,
		l.ZusatzinformationArt35,
		l.ZusatzinformationInhalt35,
		l.ZusatzinformationArt36,
		l.ZusatzinformationInhalt36,
		l.ZusatzinformationArt37,
		l.ZusatzinformationInhalt37,
		l.ZusatzinformationArt38,
		l.ZusatzinformationInhalt38,
		l.ZusatzinformationArt39,
		l.ZusatzinformationInhalt39,
		l.Stuck,
		l.Gewicht,
		l.Zahlweise,
		l.Forderungsart,
		l.Veranlagungsjahr,
		&l.ZugeordneteFalligkeit,
		l.Skontotyp,
		l.Auftragsnummer,
		l.Buchungstyp,
		l.UStSchlüsselAnzahlungen,
		l.EUMitgliedstaatAnzahlungen,
		l.SachverhaltLLAnzahlungen,
		l.EUSteuersatzAnzahlungen,
		l.ErloskontoAnzahlungen,
		l.HerkunftKz,
		l.Leerfeld,
		&l.KostDatum,
		l.SEPAMandatsreferenz,
		l.Skontosperre,
		l.Gesellschaftername,
		l.Beteiligtennummer,
		l.Identifikationsnummer,
		l.Zeichnernummer,
		&l.Postensperreis,
		l.Bezeichnung,
		l.Kennzeichen,
		l.Festschreibung,
		&l.Leistungsdatum,
		&l.DatumZuord,
	}
}

func (l CSVBookingLine) StringValues() []string {
	values := l.Values()
	st := make([]string, len(values))
	for k, v := range values {
		st[k] = fmt.Sprint(v)
	}
	return st
}
