package datev

import "fmt"

type StammdatenLine struct {
	// Personen-Kontonummer
	// Sachkontennummernlänge + 1 = Personenkontenlänge
	Konto string // 1

	// Beim Import werden die Felder in der Datenbank gefüllt, auch wenn sie
	// nicht dem Adressatentyp aus Feld 7 entsprechen. Das kann zu ungewollten
	// Effekten im Programm führen. Bitte übergeben Sie nur die zum
	// Adressatentyp passenden Felder.
	NameUnternehmen        string // 2
	Unternehmensgegenstand string // 3
	NamePerson             string // 4
	VornamePerson          string // 5
	Name                   string // 6

	// 0 = keine Angabe
	// 1 = natürliche Person
	// 2 = Unternehmen
	// Standardwert = Unternehmen
	Adressatentyp string // 7

	Kurzbezeichnung string // 8

	// Die USt-IdNr. besteht aus
	// 2-stelligen Länderkürzel
	// (siehe Dok.-Nr. 1080169; Ausnahme Griechenland: Das Länderkürzel lautet EL)
	// 13-stelliger USt-IdNr.
	// Beachten Sie, dass kein Leerzeichen zwischen diesen beiden Eingabewerten sein darf.
	EULand    string // 9
	EUUStIdNr string // 10

	Anrede string // 11

	// Nur bei Adressatentyp "natürliche Person" relevant.
	// Wird der Titel/Akad.Grad bei einem Adressatentyp "Unternehmen" übergeben, wird der Wert in den Datenbestand übernommen, ist aber an der Oberfläche nicht sichtbar.
	Titel string // 12

	// Nur bei Adressatentyp "natürliche Person" relevant.
	// Wird der Adelstitel bei einem Adressatentyp "Unternehmen" übergeben, wird der Wert in den Datenbestand übernommen, ist aber an der Oberfläche nicht sichtbar.
	Adelstitel string // 13

	// Nur bei Adressatentyp "natürliche Person" relevant.
	// Wird der Namensvorsatz bei einem Adressatentyp "Unternehmen" übergeben, wird der Wert in den Datenbestand übernommen, ist aber an der Oberfläche nicht sichtbar.
	Namensvorsatz string // 14

	// STR = Straße
	// PF = Postfach
	// GK = Großkunde
	// Wird die Adressart nicht übergeben, wird sie automatisch in Abhängigkeit zu
	// den übergebenen Feldern (Straße oder Postfach) gesetzt.
	Adressart string // 15

	// Wird sowohl eine Straße als auch ein Postfach übergeben, werden beide
	// Werte in den Datenbestand übernommen; auf der Visitenkarte in den
	// Debitoren-/Kreditoren-Stammdaten wird die Postfachadresse angezeigt.
	Strasse  string // 16
	Postfach string // 17

	Postleitzahl string // 18
	Ort          string // 19

	// ISO-Code beachten! (Dok.-Nr. 1080169)
	Land string // 20

	Versandzusatz string // 21

	// Beispiel: z. Hd. Herrn Mustermann
	Adresszusatz string // 22

	// Es kann ein beliebiger individueller Text verwendet werden.
	AbweichendeAnrede string // 23

	AbwZustellbezeichnung1 string // 24
	AbwZustellbezeichnung2 string // 25

	// 1= Kennzeichnung Korrespondenzadresse
	KennzKorrespondenzadresse string // 26

	// Format: TTMMJJJJ
	AdresseGultigVon *Date // 27

	// Format: TTMMJJJJ
	AdresseGultigBis *Date // 28

	// Standard-Telefonnummer
	Telefon string // 29

	BemerkungTelefon string // 30

	// Geschäftsleitungs-Telefonnummer
	TelefonGeschaftsleitung string // 31
	BemerkungTelefonGL      string // 32
	Email                   string // 33
	BemerkungEmail          string // 34
	Internet                string // 35
	BemerkungInternet       string // 36
	Fax                     string // 37
	BemerkungFax            string // 38
	Sonstige                string // 39
	BemerkungSonstige       string // 40
	Bankleitzahl1           string // 41
	Bankbezeichnung1        string // 42
	BankkontoNummer1        string // 43
	Länderkennzeichen1      string // 44

	IBAN1            string // 45
	Leerfeld1        string // 46
	SWIFTCode1       string // 47
	AbwKontoinhaber1 string // 48

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb1 string // 49

	// Format: TTMMJJJJ
	Bankverb1GultigVon *Date // 50

	// Format: TTMMJJJJ
	Bankverb1GultigBis *Date // 51

	Bankleitzahl2    string // 52
	Bankbezeichnung2 string // 53
	BankkontoNummer2 string // 54

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen2 string // 55

	IBAN2            string // 56
	Leerfeld2        string // 57
	SWIFTCode2       string // 58
	AbwKontoinhaber2 string // 59

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb2 string // 60

	// Format: TTMMJJJJ
	Bankverb2GultigVon string // 61

	// Format: TTMMJJJJ
	Bankverb2GultigBis string // 62

	Bankleitzahl3    string // 63
	Bankbezeichnung3 string // 64
	BankkontoNummer3 string // 65

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen3 string // 66

	IBAN3            string // 67
	Leerfeld3        string // 68
	SWIFTCode3       string // 69
	AbwKontoinhaber3 string // 70

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb3 string // 71

	// Format: TTMMJJJJ
	Bankverb3GultigVon *Date // 72

	// Format: TTMMJJJJ
	Bankverb3GultigBis *Date // 73

	Bankleitzahl4      string // 74
	Bankbezeichnung4   string // 75
	BankkontoNummer4   string // 76
	Landerkennzeichen4 string // 77

	IBAN4            string // 78
	Leerfeld4        string // 79
	SWIFTCode4       string // 80
	AbwKontoinhaber4 string // 81

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb4 string // 82

	// Format: TTMMJJJJ
	Bankverb4GultigVon *Date // 83

	// Format: TTMMJJJJ
	Bankverb4GultigBis *Date // 84

	Bankleitzahl5    string // 85
	Bankbezeichnung5 string // 86
	BankkontoNummer5 string // 87

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen5 string // 88

	IBAN5            string // 89
	Leerfeld5        string // 90
	SWIFTCode5       string // 91
	AbwKontoinhaber5 string // 92

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb5 string // 93

	// Format: TTMMJJJJ
	Bankverb5GultigVon *Date // 94

	// Format: TTMMJJJJ
	Bankverb5GultigBis *Date // 95

	Leerfeld11  string // 96
	Briefanrede string // 97
	Grussformel string // 98

	// Kann nicht geändert werden, wenn zentralisierte Geschäftspartner
	// verwendet werden.
	Kundennummer string // 99

	Steuernummer string // 100

	// 1 = Deutsch
	// 4 = Französisch
	// 5 = Englisch
	// 10 = Spanisch
	// 19 = Italienisch
	Sprache string // 101

	Ansprechpartner string // 102
	Vertreter       string // 103
	Sachbearbeiter  string // 104

	// 0 = Nein
	// 1 = Ja
	DiverseKonto string // 105

	// 1 = Druck
	// 2 = Telefax
	// 3 = E-Mail
	Ausgabeziel string // 106

	// 0 = Zahlungen in Eingabewährung
	// 2 = Ausgabe in EUR
	Wahrungssteuerung string // 107

	// Nur für Debitoren gültig
	// Beispiel: 1.123.123.123
	KreditlimitDebitor string // 108

	// Eine gespeicherte Zahlungsbedingung kann hier einem Geschäftspartner
	// zugeordnet werden.
	Zahlungsbedingung string // 109

	// Nur für Debitoren gültig
	FalligkeitInTagenDebitor string // 110

	// Nur für Debitoren gültig
	// Beispiel: 12,12
	SkontoInProzentDebitor string // 111

	// Nur für Kreditoren gültig
	KreditorenZiel1Tage string // 112

	// Nur für Kreditoren gültig
	// Beispiel: 12,12
	KreditorenSkonto1Prozent string // 113

	// Nur für Kreditoren gültig
	KreditorenZiel2Tage string // 114

	// Nur für Kreditoren gültig
	// Beispiel: 12,12
	KreditorenSkonto2Prozent string // 115

	// Nur für Kreditoren gültig
	KreditorenZiel3BruttoTage string // 116

	// Nur für Kreditoren gültig
	KreditorenZiel4Tage string // 117

	// Nur für Kreditoren gültig
	// Beispiel: 12,12
	KreditorenSkonto4Prozent string // 118

	// Nur für Kreditoren gültig
	KreditorenZiel5Tage string // 119

	// Nur für Kreditoren gültig
	// Beispiel: 12,12
	KreditorenSkonto5Prozent string // 120

	// 0 = Keine Angaben
	// 1 = 1. Mahnung
	// 2 = 2. Mahnung
	// 3 = 1. + 2. Mahnung
	// 4 = 3. Mahnung
	// 5 = (nicht vergeben)
	// 6 = 2. + 3. Mahnung
	// 7 = 1., 2. + 3. Mahnung
	// 9 = keine Mahnung
	Mahnung string // 121

	// 1 = Kontoauszug für alle Posten
	// 2 = Auszug nur dann, wenn ein Posten mahnfähig ist
	// 3 = Auszug für alle mahnfälligen Posten
	// 9 = kein Kontoauszug
	Kontoauszug string // 122

	// Leer = keinen Mahntext ausgewählt
	// 1 = Textgruppe 1
	// ...
	// 9 = Textgruppe 9
	Mahntext string // 123

	// Leer = keinen Mahntext ausgewählt
	// 1 = Textgruppe 1
	// ...
	// 9 = Textgruppe 9
	Mahntext2 string // 124

	// Leer = keinen Mahntext ausgewählt
	// 1 = Textgruppe 1
	// ...
	// 9 = Textgruppe 9
	Mahntext3 string // 125

	// Leer = kein Kontoauszugstext ausgewählt
	// 1 = Kontoauszugstext 1
	// ...
	// 8 = Kontoauszugstext 8
	// 9 = Kein Kontoauszugstext
	Kontoauszugstext string // 126

	// Beispiel: 12.123,12
	MahnlimitBetrag string // 127

	// Beispiel: 12,12
	MahnlimitProzent string // 128

	// 0 = Stammdaten-Schlüsselung gilt
	// 1 = Fester Zinssatz
	// 2 = Zinssatz über Staffel
	// 9 = Keine Berechnung für diesen Debitor
	Zinsberechnung string // 129

	// Beispiel: 12,12
	Mahnzinssatz1 string // 130

	// Beispiel: 12,12
	Mahnzinssatz2 string // 131

	// Beispiel: 12,12
	Mahnzinssatz3 string // 132

	// Leer bzw. 0 = keine Angaben, es gilt die Stammdaten-Schlüsselung
	// 7 = SEPA-Lastschrift mit einer Rechnung
	// 8 = SEPA-Lastschrift mit mehreren Rechnungen
	// 9 = kein Lastschriftverfahren bei diesem Debitor
	Lastschrift string // 133

	// 0 = Einzugsermächtigung
	// 1 = Abbuchungsverfahren
	Verfahren string // 134

	// Zuordnung der gespeicherten Mandantenbank, die für das
	// Lastschriftverfahren verwendet werden soll.
	Mandantenbank string // 135

	// Leer bzw. 0 = keine Angaben, es gilt die Stammdaten-Schlüsselung
	// 7 = SEPA-Überweisung mit einer Rechnung
	// 8 = SEPA-Überweisung mit mehreren Rechnungen
	// 9 = keine Überweisungen, Schecks
	Zahlungstrager string // 136

	IndivFeld1  string // 137
	IndivFeld2  string // 138
	IndivFeld3  string // 139
	IndivFeld4  string // 140
	IndivFeld5  string // 141
	IndivFeld6  string // 142
	IndivFeld7  string // 143
	IndivFeld8  string // 144
	IndivFeld9  string // 145
	IndivFeld10 string // 146
	IndivFeld11 string // 147

	// Wird derzeit nicht übernommen
	IndivFeld12 string // 148

	// Wird derzeit nicht übernommen
	IndivFeld13 string // 149

	// Wird derzeit nicht übernommen
	IndivFeld14 string // 150

	// Wird derzeit nicht übernommen
	IndivFeld15 string // 151

	// Es kann ein beliebiger individueller Text verwendet werden.
	AbweichendeAnredeRechnungsadresse string // 152

	// STR = Straße PF = Postfach GK = Großkunde Wird die Adressart nicht
	// übergeben, wird sie automatisch in Abhängigkeit zu den übergebenen
	// Feldern (Straße oder Postfach) gesetzt.
	AdressartRechnungsadresse string // 153

	// Wird sowohl eine Straße als auch ein Postfach übergeben, werden beide
	// Werte in den Datenbestand übernommen; auf der Visitenkarte in den
	// Debitoren-/Kreditoren-Stammdaten wird die Postfachadresse angezeigt.
	StrasseRechnungsadresse string // 154

	PostfachRechnungsadresse     string // 155
	PostleitzahlRechnungsadresse string // 156
	OrtRechnungsadresse          string // 157

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	LandRechnungsadresse string // 158

	VersandzusatzRechnungsadresse string // 159

	// Beispiel: z. Hd. Herrn Mustermann
	AdresszusatzRechnungsadresse string // 160

	AbwZustellbezeichnung1Rechnungsadresse string // 161
	AbwZustellbezeichnung2Rechnungsadresse string // 162

	// Format: TTMMJJJJ
	AdresseGultigVonRechnungsadresse *Date // 163

	// Format: TTMMJJJJ
	AdresseGultigBisRechnungsadresse *Date // 164

	Bankleitzahl6    string // 165
	Bankbezeichnung6 string // 166
	BankkontoNummer6 string // 167

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen6 string // 168

	IBAN6            string // 169
	Leerfeld6         string // 170
	SWIFTCode6       string // 171
	AbwKontoinhaber6 string // 172

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb6 string // 173

	// Format: TTMMJJJJ
	Bankverb6GultigVon *Date // 174

	// Format: TTMMJJJJ
	Bankverb6GultigBis *Date // 175

	Bankleitzahl7    string // 176
	Bankbezeichnung7 string // 177
	BankkontoNummer7 string // 178

	// ISO-Code beachten (oder Dok.-Nr. 1080169)
	Landerkennzeichen7 string // 179

	IBAN7            string // 180
	Leerfeld7        string // 181
	SWIFTCode7       string // 182
	AbwKontoinhaber7 string // 183

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb7 string // 184

	// Format: TTMMJJJJ
	Bankverb7GultigVon *Date // 185
	// Format: TTMMJJJJ
	Bankverb7GultigBis *Date // 186

	Bankleitzahl8    string // 187
	Bankbezeichnung8 string // 188
	BankkontoNummer8 string // 189

	// ISO-Code beachten (siehe Dok.-Nr. 108069)
	Landerkennzeichen8 string // 190

	IBAN8            string // 191
	Leerfeld8        string // 192
	SWIFTCode8       string // 193
	AbwKontoinhaber8 string // 194

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb8 string // 195

	// Format: TTMMJJJJ
	Bankverb8GultigVon *Date // 196

	// Format: TTMMJJJJ
	Bankverb8GultigBis *Date // 197

	Bankleitzahl9    string // 198
	Bankbezeichnung9 string // 199
	BankkontoNummer9 string // 200

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen9 string // 201
	IBAN9              string // 202
	Leerfeld9          string // 203
	SWIFTCode9         string // 204
	AbwKontoinhaber9   string // 205

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb9 string // 206

	// Format: TTMMJJJJ
	Bankverb9GultigVon *Date // 207

	// Format: TTMMJJJJ
	Bankverb9GultigBis *Date // 208

	Bankleitzahl10    string // 209
	Bankbezeichnung10 string // 210
	BankkontoNummer10 string // 211

	// ISO-Code beachten (siehe Dok.-Nr. 1080169)
	Landerkennzeichen10 string // 212

	IBAN10            string // 213
	Leerfeld10        string // 214
	SWIFTCode10       string // 215
	AbwKontoinhaber10 string // 216

	// Kennzeichnung als Haupt-Bankverbindung
	// 1 = Ja
	// 0 = Nein
	// Nur eine Bankverbindung eines Debitoren oder Kreditoren kann als
	// Haupt-Bankverbindung gekennzeichnet werden.
	KennzHauptBankverb10 string // 217

	// Format: TTMMJJJJ
	Bankverb10GultigVon *Date // 218

	// Format: TTMMJJJJ
	Bankverb10GultigBis *Date // 219

	// Achtung: Wird bei Verwendung zentralisierter Geschäftspartner von DATEV
	// überschrieben.
	NummerFremdsystem string // 220

	// 0 = Nein
	// 1 = Ja
	Insolvent string // 221

	// Sie können im Feld Mandatsreferenz dem Geschäftspartner je Bank eine
	// Mandatsreferenz eintragen. Für eine korrekte Verwendung muss in der
	// SEPA-Mandatsverwaltung die Mandatsreferenz für den Lastschriftteilnehmer
	// vorhanden sein.
	SEPAMandatsreferenz1  string // 222
	SEPAMandatsreferenz2  string // 223
	SEPAMandatsreferenz3  string // 224
	SEPAMandatsreferenz4  string // 225
	SEPAMandatsreferenz5  string // 226
	SEPAMandatsreferenz6  string // 227
	SEPAMandatsreferenz7  string // 228
	SEPAMandatsreferenz8  string // 229
	SEPAMandatsreferenz9  string // 230
	SEPAMandatsreferenz10 string // 231

	// Sie können für den Geschäftspartner das korrespondierende Konto (im
	// Kreditorenbereich) erfassen, wenn es sich bei dem Geschäftspartner sowohl
	// um einen Kunden als auch um einen Lieferanten handelt.
	VerknupftesOPOSxKonto string // 232

	// Format: TTMMJJJJ
	MahnsperreBis *Date // 233

	// Format: TTMMJJJJ
	LastschriftsperreBis *Date // 234

	// Format: TTMMJJJJ
	ZahlungssperreBis *Date // 235

	// 0 = nein
	// 1 = ja
	Gebuhrenberechnung string // 236

	// Mahngebühren Angabe in Euro
	// Beispiel: 5,10
	Mahngebuhr1 string // 237

	// Mahngebühren Angabe in Euro
	// Beispiel: 5,10
	Mahngebuhr2 string // 238

	// Mahngebühren Angabe in Euro
	// Beispiel: 5,10
	Mahngebuhr3 string // 239

	// 0 = nein
	// 1 = ja
	Pauschalberechnung string // 240

	// Verzugspauschale, Angabe in Euro
	// Beispiel: 10,50
	Verzugspauschale1 string // 241

	// Verzugspauschale, Angabe in Euro
	// Beispiel: 10,50
	Verzugspauschale2 string // 242

	// Verzugspauschale, Angabe in Euro
	// Beispiel: 10,50
	Verzugspauschale3 string // 243

	AlternativerSuchname string // 244

	// 0 = Aktiv
	// 1 = Inaktiv
	Status string // 245

	// 0 = Nein
	// 1 = Ja
	AnschriftManuellGeandertKorrespondenzadresse string // 246

	AnschriftIndividuellKorrespondenzadresse string // 247
	AnschriftManuellGeändertRechnungsadresse string // 248
	AnschriftIndividuellRechnungsadresse     string // 249

	// 0 = nein
	// 1 = ja
	FristberechnungBeiDebitor string // 250

	// Mahnfristen in Tagen
	Mahnfrist1 string // 251

	// Mahnfristen in Tagen
	Mahnfrist2 string // 252

	// Mahnfristen in Tagen
	Mahnfrist3 string // 253

	// Mahnfristen in Tagen
	LetzteFrist string // 254
}

func (l StammdatenLine) Validate() []error {
	var errs []error
	return errs
}

func (e StammdatenLine) Headers() []string {
	return []string{
		"Konto",
		"Name (Adressatentyp Unternehmen)",
		"Unternehmensgegenstand",
		"Name (Adressatentyp natürl. Person)",
		"Vorname (Adressatentyp natürl. Person)",
		"Name (Adressatentyp keine Angabe)",
		"Adressatentyp",
		"Kurzbezeichnung",
		"EU-Land",
		"EU-USt-IdNr.",
		"Anrede",
		"Titel/Akad. Grad",
		"Adelstitel",
		"Namensvorsatz",
		"Adressart",
		"Straße",
		"Postfach",
		"Postleitzahl",
		"Ort",
		"Land",
		"Versandzusatz",
		"Adresszusatz",
		"Abweichende Anrede",
		"Abw. Zustellbezeichnung 1",
		"Abw. Zustellbezeichnung 2",
		"Kennz. Korrespondenzadresse",
		"Adresse Gültig von",
		"Adresse Gültig bis",
		"Telefon",
		"Bemerkung (Telefon)",
		"Telefon Geschäftsleitung",
		"Bemerkung (Telefon GL)",
		"E-Mail",
		"Bemerkung (E-Mail)",
		"Internet",
		"Bemerkung (Internet)",
		"Fax",
		"Bemerkung (Fax)",
		"Sonstige",
		"Bemerkung (Sonstige)",
		"Bankleitzahl 1",
		"Bankbezeichnung 1",
		"Bankkonto-Nummer 1",
		"Länderkennzeichen 1",
		"IBAN 1",
		"Leerfeld 1",
		"SWIFT-Code 1",
		"Abw. Kontoinhaber 1",
		"Kennz. Haupt-Bankverb. 1",
		"Bankverb. 1 Gültig von",
		"Bankverb. 1 Gültig bis",
		"Bankleitzahl 2",
		"Bankbezeichnung 2",
		"Bankkonto-Nummer 2",
		"Länderkennzeichen 2",
		"IBAN 2",
		"Leerfeld 2",
		"SWIFT-Code 2",
		"Abw. Kontoinhaber 2",
		"Kennz. Haupt-Bankverb. 2",
		"Bankverb. 2 Gültig von",
		"Bankverb. 2 Gültig bis",
		"Bankleitzahl 3",
		"Bankbezeichnung 3",
		"Bankkonto-Nummer 3",
		"Länderkennzeichen 3",
		"IBAN 3",
		"Leerfeld 3",
		"SWIFT-Code 3",
		"Abw. Kontoinhaber 3",
		"Kennz. Haupt-Bankverb. 3",
		"Bankverb. 3 Gültig von",
		"Bankverb. 3 Gültig bis",
		"Bankleitzahl 4",
		"Bankbezeichnung 4",
		"Bankkonto-Nummer 4",
		"Länderkennzeichen 4",
		"IBAN 4",
		"Leerfeld 4",
		"SWIFT-Code 4",
		"Abw. Kontoinhaber 4",
		"Kennz. Haupt-Bankverb. 4",
		"Bankverb. 4 Gültig von",
		"Bankverb. 4 Gültig bis",
		"Bankleitzahl 5",
		"Bankbezeichnung 5",
		"Bankkonto-Nummer 5",
		"Länderkennzeichen 5",
		"IBAN 5",
		"Leerfeld 5",
		"SWIFT-Code 5",
		"Abw. Kontoinhaber 5",
		"Kennz. Haupt-Bankverb. 5",
		"Bankverb. 5 Gültig von",
		"Bankverb. 5 Gültig bis",
		"Leerfeld 11",
		"Briefanrede",
		"Grußformel",
		"Kundennummer",
		"Steuernummer",
		"Sprache",
		"Ansprechpartner",
		"Vertreter",
		"Sachbearbeiter",
		"Diverse-Konto",
		"Ausgabeziel",
		"Währungssteuerung",
		"Kreditlimit (Debitor)",
		"Zahlungsbedingung",
		"Fälligkeit in Tagen (Debitor)",
		"Skonto in Prozent (Debitor)",
		"Kreditoren-Ziel 1 (Tage)",
		"Kreditoren-Skonto 1 (%)",
		"Kreditoren-Ziel 2 (Tage)",
		"Kreditoren-Skonto 2 (%)",
		"Kreditoren-Ziel 3 Brutto (Tage)",
		"Kreditoren-Ziel 4 (Tage)",
		"Kreditoren-Skonto 4 (%)",
		"Kreditoren-Ziel 5 (Tage)",
		"Kreditoren-Skonto 5 (%)",
		"Mahnung",
		"Kontoauszug",
		"Mahntext",
		"Mahntext 2",
		"Mahntext 3",
		"Kontoauszugstext",
		"Mahnlimit Betrag",
		"Mahnlimit %",
		"Zinsberechnung",
		"Mahnzinssatz 1",
		"Mahnzinssatz 2",
		"Mahnzinssatz 3",
		"Lastschrift",
		"Verfahren",
		"Mandantenbank",
		"Zahlungsträger",
		"Indiv. Feld 1",
		"Indiv. Feld 2",
		"Indiv. Feld 3",
		"Indiv. Feld 4",
		"Indiv. Feld 5",
		"Indiv. Feld 6",
		"Indiv. Feld 7",
		"Indiv. Feld 8",
		"Indiv. Feld 9",
		"Indiv. Feld 10",
		"Indiv. Feld 11",
		"Indiv. Feld 12",
		"Indiv. Feld 13",
		"Indiv. Feld 14",
		"Indiv. Feld 15",
		"Abweichende Anrede (Rechnungsadresse)",
		"Adressart (Rechnungsadresse)",
		"Straße (Rechnungsadresse)",
		"Postfach (Rechnungsadresse)",
		"Postleitzahl (Rechnungsadresse)",
		"Ort (Rechnungsadresse)",
		"Land (Rechnungsadresse)",
		"Versandzusatz (Rechnungsadresse)",
		"Adresszusatz (Rechnungsadresse)",
		"Abw. Zustellbezeichnung 1 (Rechnungsadresse)",
		"Abw. Zustellbezeichnung 2 (Rechnungsadresse)",
		"Adresse Gültig von (Rechnungsadresse)",
		"Adresse Gültig bis (Rechnungsadresse)",
		"Bankleitzahl 6",
		"Bankbezeichnung 6",
		"Bankkonto-Nummer 6",
		"Länderkennzeichen 6",
		"IBAN 6",
		"Leerfeld 6",
		"SWIFT-Code 6",
		"Abw. Kontoinhaber 6",
		"Kennz. Haupt-Bankverb. 6",
		"Bankverb. 6 Gültig von",
		"Bankverb. 6 Gültig bis",
		"Bankleitzahl 7",
		"Bankbezeichnung 7",
		"Bankkonto-Nummer 7",
		"Länderkennzeichen 7",
		"IBAN 7",
		"Leerfeld 7",
		"SWIFT-Code 7",
		"Abw. Kontoinhaber 7",
		"Kennz. Haupt-Bankverb. 7",
		"Bankverb. 7 Gültig von",
		"Bankverb. 7 Gültig bis",
		"Bankleitzahl 8",
		"Bankbezeichnung 8",
		"Bankkonto-Nummer 8",
		"Länderkennzeichen 8",
		"IBAN 8",
		"Leerfeld 8",
		"SWIFT-Code 8",
		"Abw. Kontoinhaber 8",
		"Kennz. Haupt-Bankverb. 8",
		"Bankverb. 8 Gültig von",
		"Bankverb. 8 Gültig bis",
		"Bankleitzahl 9",
		"Bankbezeichnung 9",
		"Bankkonto-Nummer 9",
		"Länderkennzeichen 9",
		"IBAN 9",
		"Leerfeld 9",
		"SWIFT-Code 9",
		"Abw. Kontoinhaber 9",
		"Kennz. Haupt-Bankverb. 9",
		"Bankverb. 9 Gültig von",
		"Bankverb. 9 Gültig bis",
		"Bankleitzahl 10",
		"Bankbezeichnung 10",
		"Bankkonto-Nummer 10",
		"Länderkennzeichen 10",
		"IBAN 10",
		"Leerfeld 10",
		"SWIFT-Code 10",
		"Abw. Kontoinhaber 10",
		"Kennz. Haupt-Bankverb. 10",
		"Bankverb. 10 Gültig von",
		"Bankverb. 10 Gültig bis",
		"Nummer Fremdsystem",
		"Insolvent",
		"SEPA-Mandatsreferenz 1",
		"SEPA-Mandatsreferenz 2",
		"SEPA-Mandatsreferenz 3",
		"SEPA-Mandatsreferenz 4",
		"SEPA-Mandatsreferenz 5",
		"SEPA-Mandatsreferenz 6",
		"SEPA-Mandatsreferenz 7",
		"SEPA-Mandatsreferenz 8",
		"SEPA-Mandatsreferenz 9",
		"SEPA-Mandatsreferenz 10",
		"Verknüpftes OPOS-Konto",
		"Mahnsperre bis",
		"Lastschriftsperre bis",
		"Zahlungssperre bis",
		"Gebührenberechnung",
		"Mahngebühr 1",
		"Mahngebühr 2",
		"Mahngebühr 3",
		"Pauschalberechnung",
		"Verzugspauschale 1",
		"Verzugspauschale 2",
		"Verzugspauschale 3",
		"Alternativer Suchname",
		"Status",
		"Anschrift manuell geändert (Korrespondenzadresse)",
		"Anschrift individuell (Korrespondenzadresse)",
		"Anschrift manuell geändert (Rechnungsadresse)",
		"Anschrift individuell (Rechnungsadresse)",
		"Fristberechnung bei Debitor",
		"Mahnfrist 1",
		"Mahnfrist 2",
		"Mahnfrist 3",
		"Letzte Frist",
	}
}

func (l StammdatenLine) Values() []interface{} {
	return []interface{}{
		l.Konto,                                  // 1
		l.NameUnternehmen,                        // 2
		l.Unternehmensgegenstand,                 // 3
		l.NamePerson,                             // 4
		l.VornamePerson,                          // 5
		l.Name,                                   // 6
		l.Adressatentyp,                          // 7
		l.Kurzbezeichnung,                        // 8
		l.EULand,                                 // 9
		l.EUUStIdNr,                              // 10
		l.Anrede,                                 // 11
		l.Titel,                                  // 12
		l.Adelstitel,                             // 13
		l.Namensvorsatz,                          // 14
		l.Adressart,                              // 15
		l.Strasse,                                // 16
		l.Postfach,                               // 17
		l.Postleitzahl,                           // 18
		l.Ort,                                    // 19
		l.Land,                                   // 20
		l.Versandzusatz,                          // 21
		l.Adresszusatz,                           // 22
		l.AbweichendeAnrede,                      // 23
		l.AbwZustellbezeichnung1,                 // 24
		l.AbwZustellbezeichnung2,                 // 25
		l.KennzKorrespondenzadresse,              // 26
		l.AdresseGultigVon,                       // 27
		l.AdresseGultigBis,                       // 28
		l.Telefon,                                // 29
		l.BemerkungTelefon,                       // 30
		l.TelefonGeschaftsleitung,                // 31
		l.BemerkungTelefonGL,                     // 32
		l.Email,                                  // 33
		l.BemerkungEmail,                         // 34
		l.Internet,                               // 35
		l.BemerkungInternet,                      // 36
		l.Fax,                                    // 37
		l.BemerkungFax,                           // 38
		l.Sonstige,                               // 39
		l.BemerkungSonstige,                      // 40
		l.Bankleitzahl1,                          // 41
		l.Bankbezeichnung1,                       // 42
		l.BankkontoNummer1,                       // 43
		l.Länderkennzeichen1,                     // 44
		l.IBAN1,                                  // 45
		l.Leerfeld1,                              // 46
		l.SWIFTCode1,                             // 47
		l.AbwKontoinhaber1,                       // 48
		l.KennzHauptBankverb1,                    // 49
		l.Bankverb1GultigVon,                     // 50
		l.Bankverb1GultigBis,                     // 51
		l.Bankleitzahl2,                          // 52
		l.Bankbezeichnung2,                       // 53
		l.BankkontoNummer2,                       // 54
		l.Landerkennzeichen2,                     // 55
		l.IBAN2,                                  // 56
		l.Leerfeld2,                              // 57
		l.SWIFTCode2,                             // 58
		l.AbwKontoinhaber2,                       // 59
		l.KennzHauptBankverb2,                    // 60
		l.Bankverb2GultigVon,                     // 61
		l.Bankverb2GultigBis,                     // 62
		l.Bankleitzahl3,                          // 63
		l.Bankbezeichnung3,                       // 64
		l.BankkontoNummer3,                       // 65
		l.Landerkennzeichen3,                     // 66
		l.IBAN3,                                  // 67
		l.Leerfeld3,                              // 68
		l.SWIFTCode3,                             // 69
		l.AbwKontoinhaber3,                       // 70
		l.KennzHauptBankverb3,                    // 71
		l.Bankverb3GultigVon,                     // 72
		l.Bankverb3GultigBis,                     // 73
		l.Bankleitzahl4,                          // 74
		l.Bankbezeichnung4,                       // 75
		l.BankkontoNummer4,                       // 76
		l.Landerkennzeichen4,                     // 77
		l.IBAN4,                                  // 78
		l.Leerfeld4,                              // 79
		l.SWIFTCode4,                             // 80
		l.AbwKontoinhaber4,                       // 81
		l.KennzHauptBankverb4,                     // 82
		l.Bankverb4GultigVon,                     // 83
		l.Bankverb4GultigBis,                     // 84
		l.Bankleitzahl5,                          // 85
		l.Bankbezeichnung5,                       // 86
		l.BankkontoNummer5,                       // 87
		l.Landerkennzeichen5,                     // 88
		l.IBAN5,                                  // 89
		l.Leerfeld5,                              // 90
		l.SWIFTCode5,                             // 91
		l.AbwKontoinhaber5,                       // 92
		l.KennzHauptBankverb5,                    // 93
		l.Bankverb5GultigVon,                     // 94
		l.Bankverb5GultigBis,                     // 95
		l.Leerfeld11,                             // 96
		l.Briefanrede,                            // 97
		l.Grussformel,                            // 98
		l.Kundennummer,                           // 99
		l.Steuernummer,                           // 100
		l.Sprache,                                // 101
		l.Ansprechpartner,                        // 102
		l.Vertreter,                              // 103
		l.Sachbearbeiter,                         // 104
		l.DiverseKonto,                           // 105
		l.Ausgabeziel,                            // 106
		l.Wahrungssteuerung,                      // 107
		l.KreditlimitDebitor,                     // 108
		l.Zahlungsbedingung,                      // 109
		l.FalligkeitInTagenDebitor,               // 110
		l.SkontoInProzentDebitor,                 // 111
		l.KreditorenZiel1Tage,                    // 112
		l.KreditorenSkonto1Prozent,               // 113
		l.KreditorenZiel2Tage,                    // 114
		l.KreditorenSkonto2Prozent,               // 115
		l.KreditorenZiel3BruttoTage,              // 116
		l.KreditorenZiel4Tage,                    // 117
		l.KreditorenSkonto4Prozent,               // 118
		l.KreditorenZiel5Tage,                    // 119
		l.KreditorenSkonto5Prozent,               // 120
		l.Mahnung,                                // 121
		l.Kontoauszug,                            // 122
		l.Mahntext,                               // 123
		l.Mahntext2,                              // 124
		l.Mahntext3,                              // 125
		l.Kontoauszugstext,                       // 126
		l.MahnlimitBetrag,                        // 127
		l.MahnlimitProzent,                       // 128
		l.Zinsberechnung,                         // 129
		l.Mahnzinssatz1,                          // 130
		l.Mahnzinssatz2,                          // 131
		l.Mahnzinssatz3,                          // 132
		l.Lastschrift,                            // 133
		l.Verfahren,                              // 134
		l.Mandantenbank,                          // 135
		l.Zahlungstrager,                         // 136
		l.IndivFeld1,                             // 137
		l.IndivFeld2,                             // 138
		l.IndivFeld3,                             // 139
		l.IndivFeld4,                             // 140
		l.IndivFeld5,                             // 141
		l.IndivFeld6,                             // 142
		l.IndivFeld7,                             // 143
		l.IndivFeld8,                             // 144
		l.IndivFeld9,                             // 145
		l.IndivFeld10,                            // 146
		l.IndivFeld11,                            // 147
		l.IndivFeld12,                            // 148
		l.IndivFeld13,                            // 149
		l.IndivFeld14,                            // 150
		l.IndivFeld15,                            // 151
		l.AbweichendeAnredeRechnungsadresse,      // 152
		l.AdressartRechnungsadresse,              // 153
		l.StrasseRechnungsadresse,                // 154
		l.PostfachRechnungsadresse,               // 155
		l.PostleitzahlRechnungsadresse,           // 156
		l.OrtRechnungsadresse,                    // 157
		l.LandRechnungsadresse,                   // 158
		l.VersandzusatzRechnungsadresse,          // 159
		l.AdresszusatzRechnungsadresse,           // 160
		l.AbwZustellbezeichnung1Rechnungsadresse, // 161
		l.AbwZustellbezeichnung2Rechnungsadresse, // 162
		l.AdresseGultigVonRechnungsadresse,       // 163
		l.AdresseGultigBisRechnungsadresse,       // 164
		l.Bankleitzahl6,                          // 165
		l.Bankbezeichnung6,                       // 166
		l.BankkontoNummer6,                       // 167
		l.Landerkennzeichen6,                     // 168
		l.IBAN6,                                  // 169
		l.Leerfeld6,                               // 170
		l.SWIFTCode6,                             // 171
		l.AbwKontoinhaber6,                       // 172
		l.KennzHauptBankverb6,                    // 173
		l.Bankverb6GultigVon,                     // 174
		l.Bankverb6GultigBis,                     // 175
		l.Bankleitzahl7,                          // 176
		l.Bankbezeichnung7,                       // 177
		l.BankkontoNummer7,                       // 178
		l.Landerkennzeichen7,                     // 179
		l.IBAN7,                                  // 180
		l.Leerfeld7,                              // 181
		l.SWIFTCode7,                             // 182
		l.AbwKontoinhaber7,                       // 183
		l.KennzHauptBankverb7,                    // 184
		l.Bankverb7GultigVon,                     // 185
		l.Bankverb7GultigBis,                     // 186
		l.Bankleitzahl8,                          // 187
		l.Bankbezeichnung8,                       // 188
		l.BankkontoNummer8,                       // 189
		l.Landerkennzeichen8,                     // 190
		l.IBAN8,                                  // 191
		l.Leerfeld8,                              // 192
		l.SWIFTCode8,                             // 193
		l.AbwKontoinhaber8,                       // 194
		l.KennzHauptBankverb8,                    // 195
		l.Bankverb8GultigVon,                     // 196
		l.Bankverb8GultigBis,                     // 197
		l.Bankleitzahl9,                          // 198
		l.Bankbezeichnung9,                       // 199
		l.BankkontoNummer9,                       // 200
		l.Landerkennzeichen9,                     // 201
		l.IBAN9,                                  // 202
		l.Leerfeld9,                              // 203
		l.SWIFTCode9,                             // 204
		l.AbwKontoinhaber9,                       // 205
		l.KennzHauptBankverb9,                    // 206
		l.Bankverb9GultigVon,                     // 207
		l.Bankverb9GultigBis,                     // 208
		l.Bankleitzahl10,                         // 209
		l.Bankbezeichnung10,                      // 210
		l.BankkontoNummer10,                      // 211
		l.Landerkennzeichen10,                    // 212
		l.IBAN10,                                 // 213
		l.Leerfeld10,                             // 214
		l.SWIFTCode10,                            // 215
		l.AbwKontoinhaber10,                      // 216
		l.KennzHauptBankverb10,                   // 217
		l.Bankverb10GultigVon,                    // 218
		l.Bankverb10GultigBis,                    // 219
		l.NummerFremdsystem,                      // 220
		l.Insolvent,                              // 221
		l.SEPAMandatsreferenz1,                   // 222
		l.SEPAMandatsreferenz2,                   // 223
		l.SEPAMandatsreferenz3,                   // 224
		l.SEPAMandatsreferenz4,                   // 225
		l.SEPAMandatsreferenz5,                   // 226
		l.SEPAMandatsreferenz6,                   // 227
		l.SEPAMandatsreferenz7,                   // 228
		l.SEPAMandatsreferenz8,                   // 229
		l.SEPAMandatsreferenz9,                   // 230
		l.SEPAMandatsreferenz10,                  // 231
		l.VerknupftesOPOSxKonto,                  // 232
		l.MahnsperreBis,                          // 233
		l.LastschriftsperreBis,                   // 234
		l.ZahlungssperreBis,                      // 235
		l.Gebuhrenberechnung,                     // 236
		l.Mahngebuhr1,                            // 237
		l.Mahngebuhr2,                            // 238
		l.Mahngebuhr3,                            // 239
		l.Pauschalberechnung,                     // 240
		l.Verzugspauschale1,                      // 241
		l.Verzugspauschale2,                      // 242
		l.Verzugspauschale3,                      // 243
		l.AlternativerSuchname,                   // 244
		l.Status,                                 // 245
		l.AnschriftManuellGeandertKorrespondenzadresse, // 246
		l.AnschriftIndividuellKorrespondenzadresse,     // 247
		l.AnschriftManuellGeändertRechnungsadresse,     // 248
		l.AnschriftIndividuellRechnungsadresse,         // 249
		l.FristberechnungBeiDebitor,                    // 250
		l.Mahnfrist1,                                   // 251
		l.Mahnfrist2,                                   // 252
		l.Mahnfrist3,                                   // 253
		l.LetzteFrist,                                  // 254
	}
}

func (l StammdatenLine) StringValues() []string {
	values := l.Values()
	st := make([]string, len(values))
	for k, v := range values {
		st[k] = fmt.Sprint(v)
	}
	return st
}
