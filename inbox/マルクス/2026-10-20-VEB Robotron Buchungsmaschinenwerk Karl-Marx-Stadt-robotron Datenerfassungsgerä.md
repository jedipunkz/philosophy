---
source: "https://archive.org/details/mkl-20250602-05-edv-Datenerfassungsgeraet-robotron-daro-1370-1255-Magnetbandko"
title: "robotron Datenerfassungsgerät robotron daro 1370-1255 Magnetbandkonverter Handbuch Mikroprogrammierung"
author: "VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt"
year: "1983"
captured_at: "2026-10-20T11:51:48Z"
updated_at: "2026-10-20T11:51:48Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "マルクス"
query: "Karl Marx"
plain_text_url: "https://archive.org/download/mkl-20250602-05-edv-Datenerfassungsgeraet-robotron-daro-1370-1255-Magnetbandko/Datenerfassungsgerät robotron daro 1370-1255 Magnetbandkonverter Handbuch Mikroprogrammierung_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "社会思想"
  - "唯物論"
status: raw
---

# robotron Datenerfassungsgerät robotron daro 1370-1255 Magnetbandkonverter Handbuch Mikroprogrammierung

- 著者: VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt
- 初版: 1983
- 情報源: [archive](https://archive.org/details/mkl-20250602-05-edv-Datenerfassungsgeraet-robotron-daro-1370-1255-Magnetbandko)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[マルクス]]
- 研究動向: [[マルクス-現代研究動向]]

## Full Text

robotron

1370/1255

Datenerfassungsgeräte
Magneibandkonverter

Handbuch der

Mikroprogrammierung

Handbuch der Mikroprogrammierung

robotron 1370 / 1255

Vorbemerkungen

; Begriffe der Mikroelektronik/Abkuerzungen/Anglismen
‘ Struktur und Funktion der Nikrorechnersysteme

Informationsdarstellung im Mikroprozessor

; Programmierung
; Befehlsbeschreibung

Belegung der Bin- und Ausgabetore
Vebersicht ueber die Belegung besonderer Speicherzellen
RAM

; Funktionszuordnung im ROM

Hilfsprogramme bei der Arbeit mit Mikroprogrammen

im

Q, Vorbemerkungen

Das "Handbuch der Nikroprogrammierung robotron 1370/1255" soll
als Lehrbuch und Nachschlagewerk fuer dienen.
Technische Einzelheiten sind nur soweit rueoksichtigt, wie
sie zum Verstaendnis erforderlich sind und werden teilweise
vereinfacht dargestellt.

x

Der gleichartige Grundaufbau des Daienerfassungsgeraetes
robotron 1370 und des Konverters robotron 1255 gestatten es,
trotz unterschiedlicher ee und unterschied-
licher Mikroprogramme ein Hand

uch zu nutzen‘.

Baszersehnessuszen robotron 1370,
1255, 1720

I

l

j hf I LB-Stanzer

periphaere Soft- System- | Drüoker Programme,
j

Geraete ware unter- Nakrospeicher
lagen 1370, 1255
; I_ RAM, ROM
|
Speicher I E/A-Noduln I

Elektronik sorgung

Hakzenrazenser U 808 in 1370
1255

- ist eine digitale integrierte Schaltung mit komplexem
Schaltungsinhalt

- auf einem Halbleiterohip befinden sich Rechen- und Steuer-
werk als Funktionseinheit; man spricht von einer zentralen
Verarbeitungseinheit

- ist eine frei programmierbare Recheneinheit mit festem Be-
fehlssatz oder mikroprogrammierbar (z.B. U 808);

I
I

externe Stronver-
I

1.2, Abkuerzungsverzeichnis

ALE Arithmetik- und Logik-Einheit
CPU central-processing-unit (ZVE)
ZVE Zentrale Verarbeitungseinheit
LSI-Technik large-scale-integration (Grossintegration)
ROM Lese-Speicher (read only menory)
RAN Lese- und Schreibspeiocher
(random access memory)
E/A-Modul Ein- und Ausgabemodul
MOS metal-oxid-semiconduotor-Technik
DIL dual - in - line (Gehaeuseart)
TIL transistor-transistor logio
FF Flip-Flop
U 808 Mikroprozessor vom Funkwerk Erfurt
Ari [)
Pointer-Register Zeiger-Register
Pins Kontaktstifte, Zufuehrungsdraehte
arithmetic-logio-unit Arithmetik-Logik-Einheit
stack and soratch pad dynamischer Stapel-Speicher
memory (Kellerspeicher)
adress stack Adresstellenspeicher
f = Unterprogramm - Adresspeicher
scratch pad memory Arbeitsspeicher (Notizblock=
speicher
stack pointer Stapelzeiger
flop,flag ' Flip-Flop
interrupt Unterbrechung
chip Bauelement der Mikroelektronik

Flaeche ca. 1 om

Defehkdecocuerung
Befehissieuerung

bei13%011Kk: 02 03 beiRss:02 15
Ir: 02 07

ar 2 unktionseinheiten des Mikroprozessors U_ 808

252,1, Ein- und Ausgabetore

Die (gepufferten) Ein- und Ausgabetore verbinden den Mikro-
prozessor mit den externen Geraeten.
Verarbeitungsbreite: 1 Byte = 8 bit.
Die Bereitstellung der auszugebenden Bitkombinationen
(Daten- oder Steuerformationen) erfolgt im Arbeitsregister

E)

Bitkombinationen (Daten- und Steuerinformationen) von
externen Geraeten werden im Arbeitsregister A zur Auswertung
bereitgestellt.

2.2.2, Befehlscodier Befehlssteueru;

= Befehlslaenge: ein Byte,
zwei Byte (Direktoperationen),
i drei Byte (Sprungbefehle).
- jeder anliegende Befehl wird von der "Befehlscodierung"
entschluesselt.
- entsprechend Decodierungsergebnis erfolgt
— die Steuerung der Prozessablaeufe,
- das Abrufen des naechsten Befehls.

2.2.3, Befehlszaehler/Unterprogramm-Adresspeicher

= Der Mikroprozessor verfuegt ueber die "echte" Unterpro-
grammtechnik". i

- Unterprogramme koennen maximal siebenmal geschachtelt
werden‘.

- Der Unterprogramm-Adresspeicher enthaelt acht Register:

- den Befehlszaehler,
= sieben Adresspeicher zur Aufnahme der Unterprogramm-
Ruecksprungaädressen

- Der Befehlszaehler erlaubt es, einen 16 K Speicher (ROM und
RAM) zu adressieren.

- Der Befehlszaehler speichert die Adresse des (gerade) abzu-
arbeitenden Befehls‘. Nach der Befehlsausfuehrung erhaelt der
Befehlszaehler die Adresse des naechsten abzuarbeitenden
Befehls. Befehlszaehler ist die Adresse des naechsten abzu-
arbeitenden Befehls‘. $

- Bei Aufruf eines Unterprogramms (CALL-Befehl) rueckt der
Befehlszaehler um eins nach unten (siehe Blockschaltbild).
Die so erreichte Ebene im Unterprogramm-Adresspeicher dient
im Unterprogramm als Befehlszaehler.

Im "alten" Befehlszaehler wird die Adresse des naechsten im
uebergeordneten Programm abzuarbeitenden Befehls
gespeichert. Mit Rueckkehr aus dem Unterprogramm (RETURN-Be-
fehl) wird der Befehlszaehler um eins nach oben gerueckt.

Das Programm wird mit jenem Befehl fortgesetzt, dessen
Adresse im erreichten Befehlszaehler gespeichert ist. .

- Wird die Kapazitaet des Speichers ueberschritten (mehr als
sieben geschachtelte Unterprogramm-Aufrufe),. so werden ent-
sprechende Speicherebenen zerstoert; das Programm wird un-
definiert fortgesetzt.

2.2.4. te

- Der Nikroprosessor verfuegt ueber sieben Arbeitsregister &
4 Byte (= Verarbeitungsbreite).

Bezeichnungen: A, B, GC, D, E, HE, L.

= Die Arbeitsregister bilden quasi den "internen Speicher” des
Mikroprozessors.

= Einige Register erfuellen Sonderfunktione::.
IL ‚to

- Register A erfuellt im Mikroprosessor die Akkumulator-
funktion.

- Register A enthaelt bei allen logischen und arithmetischen
Operationen den ersten Operanden.
Das Ergebnis dieser Operationen steht in A (Einadress-
prinzip),

98.707 0)

=» In H (Hoeherer Adressteil = Gruppe) und L (niederer Adress-
teil = Byteadresse) wird bei Operationen, die mit dem RON-
oder RAN-Speicher zusammenwirken, die Adresse eines Bytes
gespeichert. i

= Der Speicherplatz, dessen Adresse in H und L zwiscohenge-
speichert ist, wird allgemein als "M" bezeichnet‘.

- M ist quasi ein "externes" Register der Laenge 1 Byte, mit
dem der Mikroprozessor zusammenwirken kann (ROM-Bereich nur
lesend!).

- H und L koennen auch als Arbeitsspeicher genutzt werden.

D [5

-B, 0, D und E sind Arbeitsregister, die vom Programmierer
beliebig verwendet werderi koennen.

= In den Registern B und C wird ausserdem nach Ausfuehrungen
des Makrobefehls "Sprung ins Mikroprogram" bei robotron
1370/1255 die Adresse des naechsten im Makroprogram abzu-
arbeitenden Befehls im Register zwischengespeichert.

Zed.de Arithmetik- und Logikeinheit

- Die logischen und arithmetischen Operationen werden von
einem 8-bit-Binaerrechneiwerk mit paralleler Uebertrags-
bildung ausgefuehrt. y

— Abhaengig vom Ergebnis der arithmetischen/logischen/Ver-
schiebe - Operationeh werden von der ALE bis zu vier
Kennzeichen in sogenannten Flip-Flops eingestellt.

2.2.8. Bedingungen des NMikroprozessors fuer die Programmverzweizung

Die als Ergebnis der Operationen eingestellten 1..4 Kenn-
zeichen bilden im Befehlsvorrat des Mikroprozessors die Be-
dingungen fuer die Programmverzweigung,

Die Kennzeichen koennen den Wert O0 oder 1 haben,

Sie koennen einzeln ausgewertet und damit fuer die Program-
verzweigung genutzt werden.

2.2.6.1. Kennzeichen C (Vebertrags-FF)

- C <—— carry (Uebertrag)

- 6 wird 1 gesetzt, wenn ein Ueberlauf aus den hoechstwer-
tigen Bit des Arbeitsregisters A auftritt.
(siehe Zeichendarstellung, Pkt. 3.1.).

2.2.6.2. Kennzeic I-FF

- Z <——- zero (Null)

- Z wird bei einer logischen oder arithmetischen Operation 1
sesetzt, wenn das Resultat gleich Null ist.

6,3, Kennz V =)

- 5 <—- sign (Vorzeichen)

- 5 wird bei einer logischen oder arithmetischen Operation 1
gesetzt, wenn das hoechste bit gleich 1 ist (siehe Zeichen-
darstellung Pkt. 3.1.).

2.2.6.4. Kennzeichen ar 8-

- P<—- Parity

- P wird bei gerader Anzahl von Null verschiedener Bitstellen

(2, 4, 6° oder 8 Bitstellen) nach arithmetischen oder
logischen Operationen gleich 1 gesetzt.

3. Informationsdarstellung im Mikroprozessor

Alle Informationen des Nikroprozessors werden dual in
einem Datenwort fester Laenge dargestellt

8 bit = 1 Byte.

Innerhalb des Mikroprozessors werden die Informationen von
einem

8-bit-Datenbus
uebertragen.
Fuer die Uebertragung von Speicheradressen nach den Regis-
tern H und L (vgl. 2.2.4.2.) bzw. nach dem Adresspeicher des
Mikroprozessors sind deshalb zwei aufeinanderfolgende Ueber-
tragungen erforderlich.

Die Verbindung des Mikroprozessors zu den angeschlossenen
Speichern und externen Geraeten wird durch einen

E/A-Puffer
hergestellt.
Auch der E/A-Puffer hat eine Laenge von 8 bit
(Anschluesse I ... D,).

Veber diesen Puffer erfolgt der gesamte Befehls- und Datenaus-
tausch.
Datenwort:

D; Do Binaerstelle
br ah ler ae her ste
1

2 De a a a a 5 - Stellenwert

Das Datenwort kann demnach intern die Werte 0...255 annehmen.

‘1, Zahlendarstellu;

Die interne Darstellung kann durch Vereinbarung der Lage eines
Kommas unterschiedlich interpretiert werden:

Ganze Zahl I T T T 1 T 1 T I

(positiv) t ab ze Se es ee fe el
ij as 29 08x#255
Festkommazahl
(einfache e F
Genauigkeit) rt ige D 0=x=0, 99609375

Muessen Zahlen groesserer Laenge verarbeitet werden, wird dazu
intern das Vebertragungszeichen C (vgl. 2.2.6.1.) genutzt;

Mehrbytaddition (Beispiel)

Byte n PR Byte 2 Byte 1
Zahl x Xn ... x2 xa
Zahl y Yn .. Yy2 ya
Vebertrag C Cn-4 ..... C =
xty=z Zn ... 22 2,

BSEENSSRSERSTSSSmEeESeSmeommesenmmmmemeeesemmesereeesmmcnnn

Negative Zahlen werden durch das Zweierkomplement dargestellt.
Die Stelle D,; wird als Vorzeichen gedeutet:

D)}=0 positive Zahl

D) = 1 negative Zahl
Vom Reohnerwerk wird D als "normale" Binaerstelle behandelt.
Als Zahlenbereich ergibt sich infolge der Vorzeichendarstel-
lung in;

-2lax=2l-ı

Fuer die Bildung des Zweierkomplements gibt es zwei Moeglich-
keiten:

4. Invertierung jeder Binaerstelle und Addition von 1
2. Pruefung aller Binaerstellen ab Du auf # 0,

Invertierung aller Binaerstellen links von der ersten von 0
verschiedenen Stelle

Ammeorkungi

Invertierung ist die Umkehrung des Zustandes jeder Bitstelle
(0 > 1, 1 -> 0)

Durch „die Bildung des Zweierkomplements wird die Subtraktion
im Mikroprozessor auf eine Addition zurueckgefuehrt und nur so
ausgefuehrt.

Boachtei

Das Uebertragungskennzeichen C ist unterschiedlich

2 nach der Addition von Komplementzahlen bzw.
b) nach der Subtraktion "normaler" Zahlen,

‚obwhl die 8 Stellen des Datenwortes das gleiche Ergebnis
zeigen:

4. Bildung einer Komplementzahl
+12 0000L100

LLLLOOL
+0000000

LLLLOLOO Komplement von 12

2a. Adäition der Komplementzahl

+10 0000L0L0
+ (- 12) LLLLLLLO
nn

Ka En,
6=0 LLLLLLLO

10

2b. Subtraktion der normalen Zahl

+10 0000L0L0
+ +12) 0000L1L00

— |mununnuuummun

0=L LLLLLLLO

WSZWERSSSERETEKETZunzmnum
(Daes Do), = (Dres Do)y
4% %
Ask Befshlsdarstellung
Der Mikroprozessor wird durch die intern entschluesselten
Mikrobefehle sowie durch interne Signale gesteuert.

Ergebnis:

Ein NMikrobefehl enthaelt die Art der auszufuehrenden
Operationen und die Operanden-Adressen oder bei Programmver-
zweigungen die Befehle-Adressen fuer die auszufuehrende
Sprungoperation.

de nach Art der auszufuehrenden Operation haben die Kikrobe-
fehle eine Laenge von

41, 2 oder 3 Byte:

rBxte-Befehle
I 1 T T T T 17 I
\ l l | l l l |
o
sind dann moeglich, wenn
a) eine Register-Adresse (vgl. 2.2.4.) angegeben werden muss
(3 bit). ’ Dann sind 5 bit fuer den Operationscode
verfuegbar, oder

b) zwei Register-Adressen angegeben werden muessen (6 bit) und
zwei bit fuer den Operationscode ausreichen.

Z-örte-Beichle

I 1 Ixil T 1 T T I
LI _L_L 0 L00Lj Direktoperand

werden benoetigt, wenn ,

im ersten Byte der Operationscode (6% 5 bit) und eine Register-
Adresse (3 bit) angegeben ist und

im zweiten Byte ein b als sogenannter Direkt-
operand (vgl. 3.3.3.) angegeben Ist.

11

3-Byte-Befehle

I 1 1 T T 1 T T I LI-Teil

Befehls-ADR

> RE « | | | | I H-Teil

werden benoetigt, wenn bei Programmverzweigungen.

Im ersten Byte wird die Art der Sprungoperation, im zweiten
Byte der niedere Adressteil (L) und im dritten Byte der hohe
Adressteil (H) des Sprungziels angegeben.

Bei ausgefuehrtem Sprung werden die 14 bit der neuen Befehls-
Adresse in den Adresspeicher des Mikroprozessors in die Byte L
(8 bit) und H (6 bit) entsprechend eingeschrieben. R

er Adressierungsarten

Bei den Mikroprozessoren ist die direkte und die indirekte
Adressierung moeglich.

Ausserdem wird die unmittelbare Definition von Operanden im
Befehl (Direktoperanden) angewendet.

3.331. Direkte Adressierung

Die direkte Adressierung wird bei den Arbeitsregistern A«...L
sowie bei Programmverzweigungen angewendet.

Fuer die Adressierung der Arbeitsregister werden 93 bit
benoetigt:

Duebsyoaw»
Hermoooo
HoormHoo
oHororo

Fuer die komplette Adressierung des Speichers werden 14 bit
benoetigt (Sprung an jedes Byte des Speichers ist damit
prinzipiell moeglich).

12

2% Indirekte Adress

‚Die endgueltige Adresse fuer die Ausfuehrung einer Operation

wird duroh Modifikation einer vorgegebenen Adresse ermittelt.
Dafuer kennt der Mikroprozessor mehrere Moeglichkeiten:

tet direkte B eber Speicheradress-Register

Wie unter 3.3.1. dargestellt, werden fuer die komplette Adres-
sierung des Speichers 14 bit benoetigt.

Das gilt fuer die direkte Darstellung der Sprung-Adressen
ebenso wie fuer die Adressierung des Speichers mit dem Ziel
arithmetischer oder logischer Operationen.

Die Adresse des Operanden mit der allgemeinen Bezeichnung "N"
wird in die Speicheradress-Register L (8 bit fuer niederen
Teil) und H (6 bit fuer den hoeheren Teil) eingeschrieben.

Die Ausfuehrung der Operation erfolgt in zwei Schritten:

1. Laden der Speicheradress-Register mit der Adresse des
Operanden "M"

2 Ausfuehrung der Operation mit "M"._
3.2,2,. Adressensubstitutio

Die Substitution von Adressen kann ebenfalls nur ueber die
Speicheraäress-Register H und L erfolgen:

1% Lade Register H, L mit Operanden-ADR "M"
(M selbst enthaelt Adressdaten!)

2; Laden H und/oder L mit M
3. Ausfuehrung der Operation mit M.

326 3,

Der Mikroprozessor verwendet fuer die Indizierung keine
speziellen Befehle. Ausser dem Arbeitsregister A koennen aber
praktisch alle Arbeitsregister als Indexregister benutzt
werden, da direkte Befehle fuer Addition bzw. Subtraktion
von 1 vorhanden sind;

Wird die Basis-Adresse in die Speicheradress-Register H und L
geladen, kann die Aktualisierung des als Index eingesetzten
Registers direkt erfolgen (z.B. "L" - Weiterzaehlen um + 1).

13

Der Mikroprozessor kennt die Moeglichkeit, unter Umgehung
seiner gesonderten Anfangawertbildung und damit verbundener
Adressierung direkt im Befehl einen Operanden zu definieren.

Bei entsprechendem Operationsoode wird das zweite Byte des Be-
fehls als Datenbyte erkannt und direkt fuer die Operation ver-
wendet.

Die Operation kann sowohl mit einem Arbeitsregister (direkte
Adressierung) als auch mit dem durch H und L definierten
Speicherplatz M (indirekte Adressierung) erfolgen.

Hi  Broszammierung robotron 1370/1255

Die Verarbeitungsbreite (8 bit) und die Adressierungsmoeglich-
keit (14 bit) des Mikroprozessors U 808 bedingen eine ele-
mentare Gestaltung des Befehlsvorrates. Diese Tatsache ist
vorteilhaft fuer die Anwendung des Hikroprozessors in einem
Mikroreohner, Aus den elementaren Mikrobefehlen des WHikro-
prozessors lassen nich die gewuenschten Funktionen des NMikro-
reohners zusammensetzen. Jede dieser Funktionen (Makrobefehle
des Mikrorechners) loest demnach eine festprogrammierte Folge
von Nikrobefehlen aus. Diese NMikrobefehlsfolgen und -zyklen
sind im ROH gespeichert. In den Makrobefehlsvorrat werden
solche Befehle aufgenommen, die den anwendungstechnischen
Forderungen des Mikrorechnersystems entsprechen.

Auer : den Fall robotron 1370 bedeutet es, dans die Makrobe-
fehlsliste fuer Probleme der Datenerfassung aufgestellt wurde
und die fuer robotron 1255 zur Loesung von Aufgaben der Kon-
vertierung. .
Der Ablauf der Arbeiten an den Anlagen robotron 1370/1255 ist
fentgelegt durch die im RAU befindlichen Anwendungsprogramme..
Diese Programme bestehen im allgemeinen aus Makrobsefehlen. Der
Anwender kann Standardprogramme oder auch selbst erarbeitete
Programme benutzen.

Nil Tunktion des Nikroprogrammen im ROM

Mach dem Start des Anwenderprogrammes im RAU uebernimmt das
Anwenderprosraim gemaess der darin programmierten Befehle die
Ablaufsteuerung der Arbeiten an den Anlagen robotron
1370/1255. Jeder Nakrobefehl wird durch fest zugeordnete
Mikrobefehlsfolgen und »zyklen verwirklicht. Die logische
Verarbeitung der Befehle nacheinander (linear oder verzweigt
durch Sprungbefehle) fuehrt der spezielle Zyklus "Hakrode-
eodierung" durch, Dieser Hikroprogrammeyklus ist ebenfalls im
ROM enthalten.

14

Makrodecodierung (vereinfachtes Schema)

———u—>

Transport des Makrobefehlszaehlers
aus dem RAY in den CPU

Transport des Hakrobefehls aus dem
RAM in den CPU 1)

Weiterzaehlen des Makrobefehlszaehlers
im RAM 1)

Entsohluesselung und Durohfuehrung

)
!
+
]
+
I
|
+
I
l
+
des Makrobefehls

1) Dabei berueoksichtigt die Makrodecodierung die unter-

sohiedlichen Befehlslaengen (Ein- und Zweibytebefehle bei
robotron 1370, Bin-, Zwei- und Dreibytebefehle bei
robotron 1255).
Die Durohfuehr des Makrobefehls geschieht in den ent»
sprechenden Befehlsfolgen und -zyklen, die im ROM fest pro-
srammiert vorhanden sind. Am Ende jeder Ausfuehrung eines
Makrobefehls wird im allgemeinen an den Anfang der Makrode-
oodierung zurueckgesprungen und der logisch naschstfolgende
Makrobefehl abgearbeitet, Nach den Befehlen "Sprung ins
Vorwahlprogramm" und MR ins Mikroprogramm” erfolgt
keine Fortsetzung in der Hakrodecodierung.

Hebı _ Zussmmennirkung der Hakro- und Nikroprossammierung im RAN

Die Anlagen "robotron 1370/1255 geben dem Anwender die

Moeglichkeit, Programme abzuarbeiten, die aus Mikrobefehlen

des Mikroprozessors geschrieben wurden. Dazu muessen diese

Nikroprogramme in den RAM eingelagert werden. Der et

dieser Mikroprogramme erfolgt mit Hilfe des Hakrobefehls

"Sprung ins Wikroprogramm" - JINP. .

Mach der Abarbeitung dieses Nakrobefehla kehrt man nicht zur

Makrodeoodierung des naechsten Makrobefehls zurusck, sondern

os wird die Abarbeitung der angesprungenen Mikrobefehle begon-

nen.

Der Makrobefohlszaehler enthaelt nach dem Absprung die Adresse
des auf dem MWakrobefehls "JIMP" nachfolgenden Byte.

Die Adressierung der abzuarbeitenden Mikrobefehle ist in den
Anlagen robotron 1370/1255 unterschiedlich vorsunehmen,

15

Vebergang vom Makro- zum Mikroprogramm in der Anlage robotron
1370 im RAM-Bereich:

LAGZ oder LAG 1

JIMP

RST 1
Makroprogramm Mikroprogramm
(Makrobefehle) (Nikrobefehle)

Vor dem Befehl "Sprung ins Mikroprogramm" muss mit Hilfe des
Befehls "Laden Indexregister mit Adresse Gruppe 0" oder "Laden
Indexregister mit Gruppe 1" in das Indexregister 3 die
Gruppenadresse. des RAM-Platzes geladen werden, auf dem der
erste Mikrobefehl gespeichert ist. Die ee Byte-
adresse steht im zweiten Byte des Sprungs "JIMP".

Codierung der Befehle: LAGZ 09 11 xx xx oder
LAG 1 09 15 x xx

(xx xx = Gruppenadresse 02 00, 02 01, 02 02, 02 03)
und bei 2 K-Byte-RAM auch
02 04, 02 05, 02 06, 02 07
JIMP 13° 15 yy yy

ww Byteadresse von 00 00 bis 15 15)

Die Befehle LAGZ bzw. LAG 1 und JIMP muessen nicht unmittel-
bar aufeinanderfolgen. Es koennen andere Mikrobefehle da-
zwischen stehen, sofoern sie nicht das Indexregister 3 be-
einflussen.

Der Mikrobefehl
RST 1, Codierung 00 13

bewirkt die Rueckkehr ins Makroprogrammi.

Er gehoert zur Gruppe der Unterprogrammspruenge mit fester
Adresse und ruft das Mikroprogramm "Makrodecodierung" auf. Da
zum Zeitpunkt des Ueberganges in das Mikroprogramm der Makro-
befehlszaehler die Adresse desjenigen Bytes enthaelt, das auf
den Sprung "JIMP" unmittelbar folgt, wird das Makroprogramm
auch an jener Stelle fortgesetzt. Eine Abweichung von dieser
Fortsetzungsart ist moeglich, wenn in dem durch "JIMP" auge-
rufenen Mikroprogramm der Inhalt des Makrobefehlszaehlers ver-
aendert wurde.

16

Vebergang vom Makro- zum Mikroprogramm in der Anlage
robotron 1255 im RAM-Bereich:

JIMP
RST 1
Makroprogramm Mikroprogramm
(Makrobefehle) (Mikrobefehle)

Da in der Anlage robotron 1255 der Befehl
"Sprung ins Mikroprogramm"

ein Dreibytebefehl ist, kann die komplette Adresse des anzu-
springenden Mikroprogramms in dem Sprungbefehl "JIMP" codiert
werden:

JIMP 12 15 x x w ww
xx xx = Gruppenadresse von 02 00 bis 02 15
yy yy = PByteadresse von 00 00 bis 15 15
Die Rueckkehr in das Makroprogramm geschieht auf die gleiche
Art wie bei der Anlage robotron 1370:

Der Sprung "JIMP" gestattet es, in den Anlagen robotron 1370
und 1255 auch Mikroprogramme im ROM-Bereich anzuspringen‘,
Erlaeuterungen dazu stehen unter Punkt 8.

Befehlsbeschreib;
1, meine Hinweise

Die Befehlscodierung ist aus der Spalte Operationscode aus der
entsprechenden Matrix direkt waehlbar.

17

32, Symbolerlaeuterung

Register N (äkkumulator)
n R
" D
" B ki
u. H rs pen eg
" L (Speicheradressregister

wird durch H und L dargestellt

feste Adresse

steht fuer den Inhalt der Bedingungs-Flip-Flop C/2/S/P
Operationsbodes beliebig; Vorzugsoode bei x =
Register allgemein 5
Zielregister (alter Inhalt wird verdraengt)
Quellregister (Inhalt bleibt erhalten)

n beliebiger Wert (Direktoperand)

ladr niedere 8 Bit der Adresse

nadr hohe 6 Bit der Adresse

(4 Veberrags-FF; 1 bei Uebertrag

2 Null-FF; 1 bei Null

5 Vorzeichen-FF; 1 bei Vorzeichenstelle = 1

P Paarigkeits-FF, i bei paarig.

BUNNMoRr EumSunwW>

eo

3 Er. st = u bi
'

- gestrichene Felder sind nicht belegt
- offene Felder nur fuer Techniker von Bedeutung

18

: Inkrement-u.Dekrenentbefehle
6 Ir r

8 Der r

ri

Additions-u, Subtraktionsbefehle

ri

1

9 ADr r

10 DD Kim + M

11 ADn n

12 ACr r

13 ACH ee re er;
1 ACn “

15 SüUr r

16 Sum Kim - MU

17 sUn n

18 SBr r

19 SEN A: A- N -e
20 SBn n

Vergleich

21 Pr r
22 CPU Cie A< Mizinie
23 CPn n n

Q 1 2 L3 &

1200 1201 1202 1203 1204 1205 1206
1208 1209 1210 1241 1212 1213 1214
1300 1301 1302 1303 1304 1305 1306
1308 1309 1310 1311 1312 1313 1314
1400 1401 1402 1403 1404 1405 1406
1408 1409 1410 1411 1412 1413 1414
1500 1501 1502 1503 1504 1505 1506
1508 1509 1510 1511 1512 1513 1514
0006 0014 0106 0114 0206 0214 0306

A B c D BE H L
x 0008 0100 0108 0200 0208 0300
x 0009 0101 0109 0201 0209 0301

A B c D B H L
0800 0801 0802 0803 0804 0805 0806

0004
0808 0809 0810 0811 0812 0813-0814

0012
0900 0901 0902 0905 0904 0905 0906

0104
0908 0909 0910 0911 0912 0913 0914

0112

a B c D E H LI
1108 1109 1110 1411 1112 1113 1114
11

0312 s

19

u
0807

0815

0907

M

2.Byte, B

robotron 1570 -Likroprogrammierung-
Befehlsliste

kuvec';kehr unbedingt

x: 0 4 2 5 4 5 6 7.
HU Kt unbedingter iueckenrung 0007 0015 0107 0115 0207 0215 0307 0315
zur letzten Abi. im

Oyerationscodes Bytel 3.Byt
I 1
Logische Operationen (UID; Ex,ODER;ODER) N l
ri B c D E H L M l 1
24 Nu r 1000 1001 1002 1005 1004 1005 1006 ı N
25 uDH Keic u 1007 | |
26,30 n n 0204 ki
27 kur r 1006 1009 1010 1014 1012 1013 1014 I I
Ep A ı= P3 1015 I I}
23 Kun n 0212 ee wu!
50 0Rr r 1100 1101: 1102 1103 1104 1105 1106 ı l
51 08h Kim A l | ı
32 0in n 0504 . l n |
R T
kotationsbefehls (links/rechts) l f
I I
=> RIC C = A ‚Aislh. 0002 ! I
Fi! C ie A ‚Ar-(A 0102 | I
5 AG Am (A Acc :C:=A 0040 | I
36 04 A := (CHA 2een,A ),C:A 0110 \ 1
l l
I '
&; eng unbedingt I N
x: 0 4 2 5 ORDER RER RE j N

ER Pe unbedingter Sprung 0404 0412 0504 0512 0604 0612 0704 0712 | ladr | hadr
BE a En Be a a 1
! ı
Aufruf unbedingt ! l
0 1 2 3 4 Fi 6 % N! !
Pe #7 unbedingter Sprung a 0005 0015 0105 0113 0205 0213 0305 0313 | I
feste AUR nit Huec! I I
absicht ı ı
x: 0 1 2 3 4 5 6 7 I I

55 CaL Ur-Lprung mit jiuec UhLls Guns 0506 0514 Ocu6 0614 0706 0714 | ladr | hadr
ebsicht l }
! l
} I
I I
| !
I 1
k !
| 1

20

robotron 1370 -Mikroprogrammierung-
Befehlsliste

rd! Üperationscodes 2.Byte] 3.Byte

1 1
! 1
” Bedingte Operationen 1 I
o:cC 2 Ss pP I I}
wenn :c=0 I I
41 IFo Sprung 0400 0408 0500 0508 I ladr I hadr
#2 CFo Sprung mit Rueokkehrabs. 0402 0410 0502 0510 | l1adr | hadr
43 RFo Rueoksprung zur letzten 0300. 0011 0103 0111 ı j
Adre: im Stapel | 1
wenn : o= 4 eo: cC zZ S 2 | I
44 IT Sprung 0600 0608 0700 0708 | ladr | hadr
45 CTo Sprung mit Rueckkehrabs. 0602 0610 0702 0710 | ladr | hadr
46 RTo Rueoksprung zur letzten 0203 0211 0303 0311 I I
Adresse im Stapel | |
! !
Halt-Befehl l l
!
47 HLT Stop der 2VE 0000 } }
0001 1 !
1515 ı !
eh ae
I j
Bingabe-Befehl
48 INP A: = <Bingabetor>
ee
Befehle- Akkumulatorbits Bitdarstellung des Arbeitsregisters
oodierung nach INP
mess ll 2
7 6 5 4 3 2 4 0
m [0000010000000
0401 AUB 5 .1: Sel.4 Sel.3 Sel.2 Sel.i
0403 Fehler-
selbst-
halte-
kreis
ee UBS FESESSSS SEES EEE SEES
0405 Kursor- Kursor- Kursor- Taste
taste taste taste "Netz"
"Rueck- "Spalte" "Zeile" aus
sprung”

21

0409
0411
LBS
\ 0413
KB Sohreib- Bandan- KUB Kassette
nase fang - betriebs- eingelegt
” ende bereit
Ta Sicht- e Papier- letz: teohn..
barke Zeichen kontrolle Zeichen Fehler
"letzt: x
0445 Zeichen"
Ausgabefehl
49 OUT Ausgabetor : = <A >
befells- Bitdarstellung des Arbeitsregisters vor OUT
codierung E
0500
0503
mm 11 ln nn
f Markierg bit 7 Zeichen- bit 5 2 bit 4 bit 3 bit 2 bit 1
0503 Schreiben speicher
Sperren
#
0505 .

j
!

0505
ode Supr Stelle hei slnabnmmerischez.Anaeige u...
Programmanseigelampen
0507
1a 8 4 2 1a 1 Gueltig-
keit
mm nein U
0509 Fehlerlampe einschalter
-——— [00
0541 Pehlerlampe loeschen
ms des mn
Setzen der Selektoren Loesachen de: Selektoren
0513 % 3 2 1 4 3 2 1
Dez IE
os15 ”
3 mur Bedeutung fuer Techniker
0715

23

“2. Ladebefehle (Transportbefehle

Die Ladebefehle transportieren Daten zwischen den CPU-
Registern und Speichern. Die Befehle muessen eine Ausgangs-
adresse, von der die Daten zu entnehmen sind und eine Ziel-
adresse enthalten. i
Der Inhalt des Quellenspeicherplatzes wird durch den
Ladebefehl nicht veraendert. Beim Zielspeicherplatz wird das
Ergebnis eingeschrieben.

5.2.1. Laden zwischen Arbeitsregistern

Hier handelt es sich um 1-Byte-Befehle. Die Verarbeitung in
den Arbeitsregistern A, B, C, D, E, H, L (siehe Punkt 3.3.1.)
erfolgt byteweise. Die Arbeitsregister dienen der Zwischen-
speicherung beliebiger Informationen. Die Register H und L
sind speziell als Adressregister fuer die Ansteuerung be-
liebiger Adressen (ROM und RAN) gedacht, sie koennen jedoch
auch als normale Arbeitsregister Verwendung finden.

5.2.2. Laden mit "NM"

Jeder beliebige Speicherplatz (M) im RAM und ROM wird durch H
und L adressiert (siehe Punkt 3.3.2.1.).

5.2.23. Laden eines Direktoperanden "n"

Bei der Gruppe Ladebefehle gibt es eine Moeglichkeit, sowohl
Arbeitsregister als auch "HM" mit einem Direktoperanden (n) zu
laden. Der Direktoperand ist dabei immer das dem Operations-
code folgende Byte (siehe Punkt 3.3.3.).

Seie Arithmetische Befehle und Vergleich

Dazu zaehlen die Additions- und Subtraktionsbefehle sowie der
Inkrement- und der Dekrementbefehl.

Der Vergleich basiert ebenso wie die unter Punkt 5.4. be-
schriebenen logischen Operationen auf vorangegangenen
arithmetischen Operationen.

5.3.1. Inkrement- und Dekrementbefehl (INr, DCr)

Der Inhalt der Arbeitsregister B, C, D, E, H, L (siehe Pkt.

3.3%1.) wird beim Inkrementbefehl (Dekrementbefehl) um 1 dual

erhoeht (vermindert).

Das Arbeitsregister A kann nicht verwendet werden. Die

Bedingungs-Flip-Flop's P, S, Z werden eingestellt.

Beim Inkrementieren eines Registers mit dem Inhalt 255 erhaelt
man als Ergebnis $.

Beim Dekrementieren eines Registers mit dem Inhalt 9 erhaelt

man als Ergebnis. 255.

24

22.2, Additions- und Subtraktionsbefehle

Der Inhalt des Arbeitsregisters A wird entsprechend der
Rechenregeln der Arithmetik mit einem Operanden dual
verknuepft. Das Ergebnis steht im Arbeitsregister A.

Der Operand bleibt erhalten (Ausnahme: A wurde selbst als
Operand verwendet).

Die Bedingungs-Flip-Flop's (C, Z, S, P) werden entsprechend
der anliegenden Bedingungen eingestellt (siehe Punkt 2.2.6.).

„2.1. Addition und Subtraktion eines Arbeitsregisters
ADr, ACr, SUr, SBr

Als Operand wird ein Arbeitsregister (A, B, C, D, H, L) an-
gegeben. Bei den Befehlen ACr bzw. SUr wird zusaetzlich ein
vorangegangener Uebertrag beruecksichtigt und dual addiert
bzw. subtrahiert (vgl. 3.1.).

„2.2. Addition und Subtraktion eines Direktoperanden
ADn, ACn, SUn, SBM

Als Operand wird ein Direktoperand (Konstante) angegeben. Bei
den Befehlen ACn bzw. SBn wird zusaetzlich ein vorangegangener
re REN und dual addiert bzw. subtrahiert
V8L: 3.4)

„2.3. Addition und Subraktion eines Speicherplatzes M
ADM, ACM, SUM, SBH

Als Operand wird ein durch die Register H und L adressierter
Speicherplatz (M) im RAM oder ROM angegeben.

Bei den Befehlen ACH bzw. SBM wird zusaetzlich ein vorange-
gangener Uebertrag beruecksichtigt und dual addiert bzw.
subtrahiert (vgl. 3.1.). N

52323: Vergleich

Die Bedingungs-Flip-Flop 2, C geben in der nachfolgend aufge-
fuehrten Weise Aufschluss ueber das Ergebnis.

Der Inhalt der Operanden bleibt erhalten.

CPr: Vergleiche A und r

CPM: Vergleiche A und M
CPn: Vergleiche A und n

Ergebnis:
2 Bed.-Flip-Flop
c

2
(A) > (r), (M), (n) 0 0
(A) = (r), (M), (n) 4 0)
(A) < (r), (N), (n 0 4

25

Sid, Logische Operationen

Der Akkumulatorinhalt (A, 1. Operand) wird mit dem Inhalt
eines anderen Registers (r), mit dem Inhalt (M) eines durch
die Register H und L adressierten Speicherplatzes oder mit
einer Konstanten (n, 2. Byte des Befehls), bitweise
verknuepft. Das Ergebnis steht im Akkumulator.

Die Bedingungs-Flip-Flop's Z, S, P werden entsprechend dem
Bitmuster des Ergebnisses aktualisiert.

Das C-Flip-Flop wird nicht gesetzt.

ur4h

Es erfolgt bitweise Verknuepfung der Operanden nach logischer
UND-Funktion.

Ar = Aar/WY/n

0 0 0)

0 1 I) %
0074

1 4 4

Nr: Konjunktion von A und r
(A): = (A)A (vn)

NDN: Konjunktion von A und M
(A: = (A)ı (WM)

NDn: Konjunktion von A und n
(A): = (A)A (n)

8 VES_ODER

Es erfolgt bitweise Verknuepfung der Operanden nach logischer
Funktion EXKLUSIVES ODER.

4 4 0
4 () 4
0 4 1

XÄRr: Antivalenz von A und r
(Mi = (A) ® m

XRN: Antivalenz von A und M
()i=(M)® 2)

XRn: Antivalenz von A und n

(A): = (A) ® (n)

26

ie 38 tion (OD

Es erfolgt bitweise Verknuepfung der Operanden Bach logischer
Funktion ODER.

4 1 0
1 0 1
4 1 1

ORr: Disjunktion von A und r
(A) : = (A) v (r)

ORM: Disjunktion von A und M
(A) : = (A) v (M)

ORn: Disjunktion von A und n
(A) : = (A) v (n)

er befehle
51% b Spru:
581 Unbedi er Spru IMP

Das Mikroprogramm wird bei dem als Zieladresse programmierten
Befehl fortgesetzt.

Das Programm wird bei dem als ADR programmierten Befehl in
einem mehrfach benoetigten Unterprogramm fortgesetzt.

Vor der Ausfuehrung dieser Sprungoperationen wird die
Absprungaädresse + 1 in Stapelregister in Fi, und Fi, zwischen-
gespeichert‘;

In einem Unterprogramm usw. kann ein weiterer CAL-Sprung er-
folgen, dabei wird der gesamte Stapelinhalt um einen Schritt
nach unten verschoben. Das ist bis zu 7 x moeglich.

27

„5.1.3. Unbedingter Spr an eine feste Adresse mit
Rueckkehrabsicht

Es erfolgt ein Sprung in eines der nachstehenden Unter-
programme:

Codierung Unterprogramm

00/05 Vorwahlprogramm

00/13 Makrodecodierung

01/05 Warteprogramm

01/13 Tastaturabfrage

02/05 Druckkopfkontrolle

02/13 Registerrettung bei Tastaturabfrage
03/05 Protokollierprogramm

03/13 frei

Die Ansprungadresse + 41 wird in Pi L und PH zwischenge-
speichert.

Der gesamte Stapelinhalt wird um einen Schritt nach unten ver-
schoben.

5r5.1.4. Unbedingter Sprung zur letzten ADR im Stapel <RET).

Dieser Befehl erfuellt die Funktion des unbedingten Rueck-
sprungs aus einem Unterprogramm.
Der Sprung erfolgt entsprechend der aktuellen Adresse in
Stapelregister. Nach Ausfuehrung wird der Stapelinhalt um
einen Schritt nach oben verschoben‘.

bi. Bed Sprue
Sprung zur Speicheradresse, wenn das Bedingungs-Flip-Flop c
die Bedingung erfuellt. Andernfalls wird der naechste Befehl
ausgefuehrt.
ce steht fuer den Inhalt C, 2, S, P (vgl. 2.2.6.)

521,5 wenn c_= 0 (JFe

W522 S von c = 1 (JTec
5.542.3. Sprung mit Rueckkehrabsicht wenn c = 0_(CFo)

52.4. Sprung mit Rueckkehrabsicht wenn c_ = 1 (CTe

2: Rueckspr' zur letzten ADR im Stapel wenn c = 0 _(RFc

52.6, Ruecksprung zur letzten ADR im Stapel wenn c = 1 (RTc

28

Sb Sonstige Befehle
PAAR Rotationsbefehle

Vergleichbar mit Verschiebefehlen.

Rotationsbefehle haben die Aufgabe, Bitkombinationen innerhalb
eines Bytes zu verschieben (z.B. beim Packen). Die Bitkom-
bination muss im Arbeitsregister stehen.

Rotationsbefehle stellen das C-Flip-Flop (Uebertrag) ein.

5.6,1,1. Linksrotation = RLC

Linksrotation in A

Bedingungs-Flip-Flop C schaltet parallel bei Vebertrag (A 7
geht nach A 0)

5.6.1.2. Rechnerrotation = RRC

Rechtsrotation in A
Bedingungs-Flip-Flop c schaltet parallel bei Vebertrag

(Ao geht nach A 7)
AT, Ao

Ge nn nn u u m
} I
j I

Summen nn nn nn nun .—-

5.6.1.3. Linksrotation RAL

Linksrotation des Inhaltes von A bei Einschaltung des C-Flip-
Flop als 9. Bit (A7 —> C, C —> Ao)

DS ee re ee ee ee
<—-|_C] IA Ao

5.6.1.4. Rechtsrotation RAR

Rechtsrotation des Inhaltes von A bei Einschaltung des C-Flip-
Flop als 9. Bit (C —> A7, Ao --> C)

ea v2 Er om iS ba ame 5 cn

Summen nnnnnn nk in nn nn nn nn nn nn nn nn nn

29

5.6.2. Halt-Befehl HLT

Der HLT-Befehl ist ein maschineninterner Mikrobefehl, der Ver-
wendung findet, wenn die Peripherie langsamer ist als die ZE.

3.7. __Ein- und Ausgabebefehle

Veber den Ein- und Ausgabebefehl werden die Ein- und Ausgabe-
tore angesprochen.

Die Ein- und Ausgabetore stellen die Verbindung zwischen dem
Mikroprozessor und der Peripherie dar. ;
re der Ein- und Ausgabetore siehe Pkt. 6. und Befehls-
liste).

ti: = ‚be,

INP heisst: Lies den Inhalt des ausgewaehlten Eingabetores in
das Arbeitsregister A. .

Die Peripherie sendet ihre Daten- und Steuerinformationen -
ueber das angesteuerte Eingabetor in das Arbeitsregister A.
Vom Programm kann anschliessend die Auswertung der belegten
Bits im Arbeitsregister erfolgen.

„lu2 OUT_= b

OUT heisst: Schreibe den Inhalt des Arbeitsregisters A in das
ausgewaehlte Ausgabetor.

Der Mikroprozessor sendet den Inhalt des Arbeitsregisters A
(Daten- und Steuerinformationen) in das angesteuerte Ausgabe-
tor“ Die belegten Bits des Tores rufen die entsprechende
Wirkung in der Peripherie hervor.

6; Belegung der Ein- und Ausgabetore

Es erfolgt nur die Beschreibung der fuer den Programmierer
wichtigen Tore (manuelle Beeinflussung des Programmes,
Anzeigen Lampen). Die Tore werden dabei durch ihre ent-
spreohende Befehlscodierung gekennzeichnet (siehe Punkt 5.1.)

Sr;

0404 Dieses Tor dient der Selektorenabfrage.
Bit 4 = Selektor 4
Bit 3 = Selektor 3
: Bit 2 = Selektor 2
Bit 1 = Selektor 1

Sind die Bits belegt, so wurde der entsprechende
Selektor gesetzt._

Bedingungi

Vor dem Eingabebefehl muss das 5. Bit im Arbeitsre-
re mit 1 belegt werden, alle anderen Bits sind zu
oeschen. j

\ 30

0413

0413

Dieses Tor dient der Abfrage des Fehlerselbsthalte-
kreises

Bit 6 = Fehlerselbsthaltekreis

Ist dieses Bit belegt, so wurde die Fehlerlampe
eingeschaltet.

Dieses Tor dient zur Abfrage der Tastenbedienung

Bit 7 = Kursortaste "Ruecksprung"
Bit6= “ "Zeile"
Bit5= ” "Spalte"

Bit 4 = Taste "Netz aus” ö

Sind diese Bits belegt, dann wurden im Augenblick der
Abarbeitung des Ringabebefehls die entsprechenden
Tasten gedrueckt.

Dieses Tor dient der Arbeit mit dem KMB.

Bit 6 = Schreibnase

Ist dieses Bit belegt, befindet sich keine Schreib-
nase in der Kassette, die Kassette ist schreibge-
schuetzt.

Bit 4 = Taste "Loeschen Ausgabesperre"

Ist dieses Bit belegt, wurde im Augenblick der Ab-.
arbeitung des Eingabebefehls diese Taste gedrueckt.

Bit 3 = Bandanfang - ende

Ist dieses Bit belegt, dann steht die Kassette auf
dem Klarsichtband am Bandanfang oder -ende.

Bit 2 = KMB betriebsbereit

Ist dieses Bit belegt, so ist die Betriebsbereit-
schaft des KMB hergestellt‘,

Bit 1 = Kassette eingelegt

Ist dieses Bit belegt, so wurde eine Kassette einge-
legt.

Dieses Tor dient der Arbeit mit dem Drucker
Bit 7 = Taste Siohtbarkeit "Letztes Zeichen"

Ist dieses Bit belegt, so wurde im Augenblick der Ab-
arbeitung des Eingabebefehls diese Taste gedrueckt.

31

Bit 6 = erstes Zeichen

Ist dieses Bit belegt, ‚steht der Drucker in Grund-
stellung. x

Bit 5 = Papierkontrolle

Ist dieses Bit belegt, befindet sich kein Papier im
der Anlage.

Bit 2 = letztes Zeichen

Ist dieses Bit belegt, steht der Drucker in End-
stellung. \

Bit 1 = Technischer Fehler

Ist dieses Bit belegt, trat ein technischer Fehler am
Drucker auf.

622% Ausgabetore

0503
#

0505
#

Dieses Tor dient der Arbeit mit dem Bildschirm.
In Verbindung mit Tor 0505 kann es zum Loeschen des
Bildschirms verwandt werden.

Bit 7 = Markierung schreiben
Ist dieses Bit belegt, so soll der Kursor ein "_" auf

die im Tor 0505 angegebene Stelle des Bildschirmes
schreiben. ’

Bit 6 = bit 7 )
Bit 7 = bit 5 ) Spalte
Bit 3 = bit 4 )
Bit 2 = bit 3 ) Zeile
Bit 1 = bit 2 )
Bit 0 = bit 1 )

In diesen Bits wird das Zeichen verschluesselt,
welches auf der im Tor 0505 angegebenen Stelle auf
dem Bildschirm erscheinen soll.

Bit 5 = Zeichenspeicher sperren

Ist dieses Bit belegt, dann kann nicht geschrieben
werden.

In den Bit 7-0 wird die Stelle angegeben, auf welcher
ein im Tor 0503 angegebenes Zeichen angezeigt werden
soll.
Die Stelle kann 1-256 sein und wird dual in A
angegeben.
1=0
256 =111141 111149

32

0507 Dieses Tor dient der Arbeit mit den Programmanzeige-
lampen.

Bit
Bit
Bit
Bit

= Lampe
= Lampe
= Lampe
= Lampe

FUIDS
-nFo

Bei Belegung dieser Bits koennen vom Mikroprogramm
her die entsprechenden Programmanzeigelampen einge-
schaltet werden.

Bit 3 = Gueltigkeit

Ist dieses Bit belegt, dann kann vom Programm her ein
Setzen der Programmanzeigelampen erfolgen.

0509 Bei Aufruf dieses Tores erfolgt das Einschalten der
Fehlerlampe.
Die Belegung von A ist dabei ohne Bedeutung.
0511 Bei Aufruf dieses Tores erfolgt das Ausschalten der
Fehlerlampe.
Die Belegung von A ist dabei ohne Bedeutung.
0513 Dieses Tor dient der Arbeit mit den Selektoren.
Bit 7 = Selektor 4
Bit 6 = Selektor 3 )
Bit 5 = Selektor 2_ ) Selektoren setzen
Bit 4 = Selektor i 2
Bit 3 = Selektor 4 )
Bit 2 = Selektor 3 )
Bit 1 = Selektor 2_ ) Selektoren loeschen
Bit 0 = Selektor 1 )

Bei Belegung dieser Bits werden die entsprechenden
Selektoren gesetzt bzw. geloescht.

Te Uebersicht ueber die Belegung besonderer Speicherzeilen

im RAM

Die Zeilen 00 bzw. 01 der RAM-Gruppe 0 werden vom ROM-Mikro-
programm benutzt fuer die Einstellung und Auswertung von Merk-
malen.

7.1. Gruppe 0, Zeile 00

Die Speicherzeile 00 wird von den ROM-Programmteilen spezi-
fisch belegt. Die Belegung der einzelnen Bits ist nur aus den
Unterlagen des jeweiligen Mikroprogrammteiles ersichtlich.

7.2. Gruppe 0, Zeile 01

Diese Speicherzeile darf nur von speziellen Programmen fuer
Zwischenspeicherungen genutzt werden.

Die Belegung der einzelnen Bits wird in der folgenden Veber-
sicht dargestellt und in den nachfolgenden Punkten naeher er-
laeutert.

33

Byte I 7 1 6 15 l “| a 1 | 141 0.
00 1 Makrobefehlszaehler (L-Teil)

0 | neus l

02 IMerkmal| Sel Sel | Sel | Bel | Sel I Bel I 8Sel

3 IGm- larım I Gr“

Ei

I 0r.- I minus | plus |IKAo 1-0

“yn
[or er) © Q
14 I x | Start- | I & I 3 ı2 4 iu‘
I I taste | 1 I l l [

| Eingabezaehler: zachlt die eingetasteten Ziffern und das

x) ‚ Die Bytes 09, 10, 11, 12 und 13 werden im Notstromfall
mit anderen Informationen belegt.

Ze2s1s 00-04-00

Dieses Byt enthaelt den L-Teil (niedere B bit) des Makrobe-
fehlszashlers. Dieser L-Teil gibt die aktuelle Byteadresse an.

IRT-IRL- TI

- bit 0-5 enthaelt den H-Teil (obere 6 bit) des Makrobefehls-
zashlers.

- bit 6 und 7 enthaelt die neue Gruppenadresse fuer Gruppen-
umsohaltungen.

bit ! 7 6 5 4 3 2 4 ()
N

0r.0 Io 0 4 0 0 () 0) ()

ar 1 110 4 4 0 0 0 0 4

a2 I lı _0o| A 0 0 0 4 ()

03 1A 4 4 0 0 0 4 4

Dis 3u_00-01-02

Diese bits enthalten. den Wert "1" wenn der entsprechende
Belektor genetzt ist. \

Taste 00-01=03

- bit 0-3 enthaelt eine "1", wenn der entsprechende Ergebris-
ood anliegt:

- bit 4-7 enthaelt die Gruppenadresse der Indexregister.
Die Arbeit mit den Fe Hg Fe kann nur mit der RAl-Gruppe O0
und 1 erfolgen, da fuer die Gruppendarstellung der IR nur je 1
bit zur Verfuegung steht.

’ -01-04

Auf diesem Byte wird die Byteadresse des IRO abgespeichert.
Die Adressierung der Indexregister benoetigt 8 bit fuer die
Bytedarstellung, und fuer die Gruppendarstellung ist bei jedem
Indexregister nur 1 bit vorgesehen.

2,6, 00-01

Byteaädresse des IR 1
siehe. Punkt 7.2.5. ö

Byteadrense des IR 2
siehe Punkt 7.2.5.

.

35

MARIN 0900-01-07

Byteadresse des IR 3
siehe Punkt 7.2.5.

73269: 00-01-08

- Das "bit 0” wird durch den Befehl "KONV" auf "1" gestellt.
Das VWP stellt DEG-Status (bit 0=0) ein.

- bit 1, 2, 3 ist nicht belegt.

- bit 4, 5, 6 und 7 sind entsprechend dem Lampencode gesetzt.

7,210... 00-01-09
Dieses Byte wird im Notstromfall belegt.

7, 2,115 00-01-10 bis 00-01-15

Diese Byte werden zur Zwischenspeicherung nach der Tastatur-
abfrage benoetigt. e
Im Notstromfall haben diese Bytes eine besondere Belegung.

Funktionszuordnu; im ROM

Site Allgemeine Hinweise

Der Gliederungspunkt gibt Auskunft ueber alle vorhandenen
Mikroprogrammteile von robotron 1370 und 1255.

Die sich anschliessende Uebersicht enthaelt die Angaben,
welohe Mikroprogrammteile existieren und an welchen Adressen
der ROM-Gruppen die verschiedenen Mikroprogrammteile stehen.
Teile dieser Mikroprogramme koennen auch in Anwenderprogrammen
genutzt werden. Dabei sind jedoch die Anfangs- und Anschluss-
bedingungen (Parameter) der einzelnen Nikroprogramme zu be-
achten sowie die Verschachtelung der NMikroprogrammteile.

Die scheinbaren Doppelbelegungen im Punkt 8.1.1. resultieren
aus den verschiedenen Varianten der Anlage robotron 1370. Fuer
die einzelnen Varianten ist der ROM unterschiedlich belegt, so
dass bei der Benutzung dieser Mikroprogramme darauf geachtet
werden muss, welches Mikroprogramm verwendet werden soll.

Einige ausgewaehlte Mikroprogramme werden unter Punkt 8.2.
naeher beschrieben.

36

8,1,1, Zuordnung zu ROM-Gruppen 1370

Kurzbe- Name des Nikroprogrammteiles ROM-Adressen
zeichng‘. von bis
lt.techn.

Unterlgl'

g

SEN = 5 ERSS

ee:
“ Tei orwahlprogramm
Makrobef.-decodierung

1. Teil des UP-Verzoegerung
durch Warteschleife
Tastaturabfrage - Anfang -
Druckwerkskontrollprogr.

UP num‘. Tastaturabfrage mit
Registerrettung
Tastaturmatrixabfrage
Speicherplatz mit Codezeichen
Laden (ohne Weiterzaehlen ADR)
Speicherplatz mit Codezeichen
Laden (mit Weiterzaehlen ADR)
Arithmetische Befehle Addition
Sonstige Befehle - Gruppenum-
schaltung

Abwaertszaehlen Indexregister
Fortsetzung der Makrobefehls-
codierung

UP Loeschen

UP Makrobeft-zaehler weiterzaehlen

UP weiterzaehlen ADR
UP Nulltest

Arithmetische Befehle - Subtraktion

keine Operation®- Loeschen
Ergebnisoode (EC = 0)

Laden des IR mit ADR in Gruppe O0
Laden des IR mit ADR in Gruppe 1

Sprung ins Vorwahlprograma
Tastaturabfrage
Tastaturmatrixspaltenabfrage
Zeichentransport von R (ohne
Weiterzaehlen ADR)
Zeichentransport von R (mit
Weiterzaehlen ADR

Arithmetische Befehle - Addition 1

2. Teil Vorwahlprogramm
Zahlenpruefung
Zeichentransport nach R (mit
Weiterzaehlen d. ADR)
Zeichentransport nach R (ohne
Weiterzaehlen d.' ADR)
Arithmetische Befehle -
Subtraktion 1

Tabelle fuer Makrobefehlsdeood.

UP Programmeingabe manuell
UP Vorzeichenaufbereitung

37

00/00/00-00/00/07
00/00/08-00/00/15

00/01 /00-00/01/05
00/01/08=-00/01/15
00/02/00-00/02/06

00/02/08-00/02/15
00/03/00-00/03/10

00/03/11-00/04/04

00/04/05-00/04/13
00/04/14=00/05/14

00/05/12-00/06/09
00/06/10-00/07/10

00/07/11-00/43/05
00/13/06-00/14/01
00/14/02=00/15/07
00/15/08=-01/00/03
01/00/04=01/01/0%
01/01/05-04/01/08

01/01/09-01/02/05

. 01/02/06-01/02/13

01/02/14-01/04/03
01/04/04-01/15/05
01/04/05-01/15/05
01/15/06-02/03/07

02/03/14-02/07/11

02/07/12-02/08/03
02/08/04-02/09/11
02/09/12-02/14/15
02/15/00-03/03/07

03/03/08-03/03/15
03/04/00-03/05/04
03/05/05-03/05/09
03/05/10-03/10/03

03/10/04=03/14/03
03/14/04-03/14/13

ux UP Addition -

w Arithmetische Befehle - Transport

vI Be befshle

us Subtraktion

un UP Wortlaenge

U UP Addition

49 Makroprograma im P ROM

ar Unterprogramm fuer eine Wartezeit
von 4,5 ms

UB up Bel-iktualisierung

WL Arithmetische Befehle - Loeschen

FZ Netsabschaltung

FY Setzen des Selektors X

Spr.nach

pr> 4 r

Zeilenschaltung einzeilig
vorwaerts

FV re bis zum progr.

op -
FU EEHEETSIEHBER bis zur naechsten
eite

Fr Loeschen der Bildschirmanzeige

FS Setzen Selektor Y

FR Lampen setzen, Selektor 1. XY,
Loesohen, Komma fuer num. Anzeige

FE Steuerbefchle - num. Ergebnis-
anzeige

BI Sprung ins Wikroprogramm

Spru nach

”®

ch Programmeingabe manuell

vo UP Subtraktion, komplett

uc UP Brgebnisoode - Aufbereitung

HZ Wortweise Ringabe

HR Wortweise Ausgabe OWd 8...15

w Arithmetische Befehle -
Verschiebung

Spr.. nach

HZ mu>

@W Bingabe ueber Alphatastatur

AA num. LED-Anzeige

zB‘ Tabelle fuer num. LED-Anzeige

GN Ausdruck eines Speicherbereiches

HH Lochen eines Zeichens

HS KMusterdruok zugehoerig zur wort-
weisen Ausg.

HT UP su Druokprogrammen

FQ Druokkopfbewegung rechts/links

GR Druck eines Oodierzeichens

42 Tastaturtabelle lat. (Standard)

Hr Sichtbarmachung des letzten

Druckseichens

38

03/14/14-04/02/07
04/02/08=-04/03/10

04/14/03-04/15/03
04/15/04=04/15/11
04/15/12=05/00/09
05/00/10-05/01/08
05/01 /09-05/02/05
05/02/06-05/02/11
05/02/12-05/03/01
05/03/02-05/03/11
05/03/12-05/04/00
05/04/01=05/04/05
05/04/06-05/05/06
05/08/07-05/05/140
05/05/11-05/10/04
05/10/05-05/11/04
05/11/05-05/11/11

05/11/12-05/11/14

07/00/00=-07/06/12
07/06/13=07/10/06
07/10/07-07/10/09

07/10/10/=09/01/13 &
+

09/01 /14-09/05/09
09/03/10-09/06/08
09/06/09-09/09/0
09/09/10-09/12/

09/12/09-10/01/10
40/01/11-40/05/00
10/05/01-10/07/06
410/07/00=10/04/15
10/08/00-10/14/15
410/15/00-14/00/43

lpha
ast.

1. UP zur Loohbandeingabe 11/00/14=11/11/09

mu Lochbandausgabe sugehoerig sur L
wortweisen Ausgabe 11/11 /10=11/44/00 | ©
vw Arithmetische Befehle - Vergleich 12/01/13-12/03/03 | ©
Spes nach => -|h
[) 12/03/04-12/03/06 R
I: nach “> 2a
[; 12/03/07-12/03/09 .
Pal nach > ' t
() . 12/03/10-12/03/12 =
Spri nach > 5
) 12/03/13-12/04/02 e
GL Ausgabe eines Speicherbereiches
sum Lochbandstanzer 12/04/03-12/06/06
[1 Lochen eines Üodierseichens 12/06/07-12/06/13
zr Setzen und Loeschen der Belektoren
1, 2, 3 und & 12/11/02-12/12/04
2» Anfangsloeschung der Bildachirm*
anzeige 12/12/10=12/13/05 I
BV dentsraete im ROM fuer Loohband- 1700/00-13/43/
anzer 13/00/00=13/15/15 m. Be
v0 Programie von der Karts 14/00/00-14/05/12 Prager
0) Erogrammeinge auf Karte 14/05/13=14/13/11_| auf Karte
av Bingabe weber Alphatastätur 07/10/10-09/01/13
(teilweise
4 mum:; Anseige auf Bildschire 09/01/14-09/03/15 | mit Bild-
“x Anzeige eines Speicherbereiches sohirn
auf Bildschirm 09/04/00=09/05/15
au Ausdruck eines Speicherbereiches 09/06/09-09/09/09
Eu lochen eines Zeichens 09/09/10/09/12/08
1: NMusterdruek zugehoerig zur wort-
weisen Ausg 09/12/09-10/09/10
m UP m Druskprogramen 10/01/41=10/05/00 | Drucker,
nn kkorfben gung rechts/liuks 10/05/01=10/07/06 |. Tastatur-
aR ee 10/07/07=10/07/45 | tab,
| & Tastaturtabelle kyrill (0) 10/08/00-10/14/15 | Kyrill.
Bichtbarmachung des letsten Druck- (50)
eisen 10/15/00=11/00/13
BU sur Loohbapdausgabe 11/00 /14&11/11/09
Geirwin) UPMBKG 1/03/00-12/04/°
insaue 4 au 01/12 2:
(geitweise) ig gumeinsene
Arithmetische Befehle - Vergleich 12/01/13-12/03/03 Progranne
Yan nach
ri bin 12/03/04-12/03/06
“ nach
> 12/03/07-12/03/09
“nach
— ... 412/03/10-12/03/12
Spr. nach
ww 12/03/13-12/04/02

BE Wortweise Ausgabe Inhalt Arbeits-

speicher nach Datentraeger

12/04/03-12/09/11

GU Befehl zum Einlesen eines Daten-

blocks vom KMBG auf Speicher 12/09/14-12/11/01
ZzF Setzen und Loeschen der Selektoren

1, 2, 3 und 4 12/11/02-12/12/04
ZE Anfangsloescohung der Bildsohirmanz. 12/12/10-42/13/05
Spr. nach
ZF > 12/13/06-12/13/08
FP Masch.'Steuerbef\. KMBG-Decodierung 12/13/09-13/03/09
BG Spezielle UP fuer KMBG Pki 13/03/10-13/11/02
co CRC-Verrechnung fuer KMBG 13/11/03-13/14/03
(teilweise)

En Umspulprogramm fuer KMG 13/14/04-13/15/02

Gemeinsame UP MBKG 13/15/03-13/15/15
(feitveise)

Programmeingabe von KMBG 14/00/00-14/04/03
co CRC-Verrechnung fuer KMGB 44/04/04=-14/05/12
ee)

Progr.-ausgabe auf KMBG 14/05/13-14/14/11
co CRC-Verrechnung fuer KMBG 14/14/12=-14/15/15
(teilweise)

Notstrom 15/00/00-15/15/13

8.1.2. Vebersicht Funktionszuord.

zu den ROM-Gruppen 12

} rss

Kurzbe- Name des Mikroprogrammteils ROM-Adressen.
zeichng. von is
Unterlg. z
AL 4. Teil Vorwahlprogramm 0/00/00/00-0/00/00/07
w Makrobefehlscodierung (RST 1) 0/00/00/08-0/00/00/45
BR 1. Teil .UP Verzoegerung durch

Warteschleife (RST 2) 0/00/01 /00-0/00/01/07
YZ Tastaturabfrage (RST 3) 0/00/01/08-0/00/01/15
AS Druckwerkskontrollprogramm 0/00/02/00-0/00/02/07
UN UP num. Tastaturabfrage m. j

Registerrettung (RST 5) 0/00/02/08-0/00/02/15
XE Tastaturabfrage 1. Fortsetzung 0/00/03/00-0/00/03/10
WA Arithmetische Befehls Addition 0/00/03/11-0/00/04/08
WU Sonstige Befehle Gruppenum-

schaltung DEG 0/00/04/09-0/00/05/06
BV Abwaertszaehlen des IR/DEG 0/00/05/07-0/00/06/02
ZT Speicherplatz m. Codezeichen

laden 0/00/06/03-0/00/10/12
AV Fortsetzung Makrobefehlsde-

codierung 0/00/10/13-0/01/00/11
VS Arithmetische Befehle

Subtraktion 0/01/00/12-0/01/00/15
NZ keine Operation, Loe Merkmal

EC=0, Rueckspr. aus UP 0/01 7/01/00-0/01/03/04
NV Sprung in WP; Setzen Merkmal

DEG/Konv..

40

0/01/03/05-0/01/04/01

Laden d. IR Bee ne
Gr. 0 (DEG) KoNV.F.

Laden d. IR m. Adresse in

Gr% 1 (DEG) KONV:FıGW
UL Loeschen

UP Makrobefehlszashler weiter-

zaehlen

UP Weiterzaehlen Adresse

UP Nulltest

UP Vorzeichenaufbereitung

UP Addition

UP Subtraktion

Arithmetischen Befehle
Addition 1

Laden IR fi; Konverter,
direkter Steuerbefehl .

Transport Zeichen von AR

UP Wortlaenge
UP Addition

Arithmetischg Befehle
Subtraktion
Addition/Subtraktion 1,
Beginn Blocktransport,
Zeichenvergl‘.
Adressreohnung

Transport Zeichen nach A
UP .„ Wartezeit w 4,5 ms
UP E-, L Aktualisierung

Tastaturabfrage 1.Teil
Fortsetzung/Schluss
Arithmetische Befehle
Transport ADR --> R
Sprung

Blooktransport

Tabelle fr Makrobefehlsde-

codierung

Arithmetische Befehle Loeschen

Netzabschaltung
Setzen Selektor X

Zeilenschaltung einzeilig

vorwaerts

Formularvorschub bis zum

programmierten Stop

Formularvorschub bis zur

naechsten Seite
Setzen Selektor Y

Lampen Setzen, Sel. 1,X,Y,

0/01/04/02-0/01/04/09

0/01 /04/10-0/01/05/15
0/01/06/00-0/01706/11

0/01/06/12-0/01/08/04
0/01/08/0270/01/08/13
0/01 /08/14=0/01/09/14
0/01/09/15=0/01/10/08
0/01 /10/09-0/01/14/02
0/01/14/03-0/02/01/08

0/02/01/09-0/02/03/00

0/02/03/01-0/02/07/02
0/02/07/03-0/02/13/14
0/02/13/15-0/02/15/08
0/02/15/09=-0/03/01/00

0/03/01/01-0/03/01/05

0/03/01/06-0/03/04/04
0/03/04/05-0/03/08/07
0/03/08/08-0/03/12/03
0/03/12/04-0/03/12/11
0/03/12/12-0/03/13/09

0/03/13/10-0/04/01/11
0/04/01/12=0/04/02/14
0/04/02/15=-0/04/10/02
0/04/10/03-0/04/13/04
0/04/13/02-0/05/01/41

0/05/01/12-0/05/02/10

0/05/02/11-0/05/03/07

0/05/03/08-0/05/03/13
0/05/03/14-0/05/04/07
0/05/04/08-0/05/04/12

0/05/04/13-0/05/05/01
0/05/05/02-0/05/05/05

Loeschen, Komma fi. numwAnzeige 0/05/05/06-0/05/09/45
Druckkopfbewegung rechts/links 0/05/10/00-0/05/12/05

Sprung an FP
41

'Setzen und Loeschen der

Selektoren 1-4
Numerische Anzeige

#

Sprung ins Mikroprogramm
Steuerbefehle Magnetband

Befehl Zahlenpruefung
Codewandel

Aus- und Eingabe Nagnetband

Arithmetische Befehle
Verschiebung

Wortweise Ausgabe 0W 0.415

Wortweise Eingabe

Befehl Zeichenvergleich
Tastaturabfrage Fortsetzung/

Sohluss

Progr'.-eingabe manuell

0/05/12/09-0/05/13/11
0/05/13/12-0/05/14/11
0/05/14/12-0/05/15/02
0/05/15/03-0/06/03/11

0/06/03/12-0/06/12/13
0/08/12/14=-0/07/00/00

0/07/00/08-0/07/03/07
0/07/03/08-0/07/07/01
0/07/07/02-0/07/13/1&,
0/07/13/15-0/08/09/12
0/08/09/ :3-0/08/13/09
0/08/13/10-0/09/08/10

0/09/08/11-0/09/14/10

UP Progr.-eingabe manuell 0/09/14/11-0/10/02/10
Tastaturtabelle lat. (Stand“) 0/10/02/11-0/10/09/10
UP Subtraktion komplett 0/10/09/11=-0/10/10/14
UP Ergebniscodeaufbereitung 0/10/10/12-0/10/11/13
UP zur LB-Ausgabe 0/10/11/14=-0710/13/15

Sichtbarmackung d. letzten

Druckzeichens

Musterdruck zugehoerig zur

wortweisen Ausgabe

UP zu Druckprogrammen
Numerische LED-Anzeige

Tabelle f. num. LED-Anzeige

2. Teil VWP

Arithmetischer Befehl Vergl.

Druck eines Codezeichens

Ausdruck eines Speicherber.;

Sprung an GU
Sprung an BE

Eingabe ueber Alphatenstatur

Steuerbefehl KMBG
Lesen KMBG

Ausgabe von Daten auf KMBG

Progr.-ausgabe auf KMBG
Progr',-eingabe vom KMBG

UP fi: KNBG

UP BF f‘; KMBG (Schreiben)

CRO=-Zeicher-Berechnung
Umspulprogramm f'. KNBG

42

0/10/14/00-0/:0/15/13
0/10/15/ 14-0711 704/15

0/11 7/05/00-0/11/08/05
0/11 /08/ 05-9/:1,12/01
0/11/12/02-0/11/13/00
0/11/13/01-0/12/02/04

0/12/02/05-0/12/03/11
0/12/03/12-0/12/04/04
0/12/04/05-0/12/08/15

0/12/09/00-0/12/09/02

0/12/09/03-0/12/09/05
0/12/09/06-0/14/00/09
0/14/00 /10-0/14/09/15
0/14/10/00-0/15/05/14
0/15/05/15-1/00/00/04

1/00/00/05-1/00/10/02
1/00/10/03-1/00/13/10
1/00/13/11-1/01/12/04
1/01/12/05-1/03/01/09

1/03/01 /10-1/03/07/05
1/03/07/06-1/03/09/02

usgewaeh

Programmeingabe
von PKi

Verriegeln einer
MBE:

Zahlenpruefungi

Rusekspulen ohne
EM-Aufzeichnung:

Verspulen ohne
EM-Aufseichnungi

ro teile 1370

Program CC

ROM 14 Byte 0000
Ansprung durch CAL
nach P E und T+ im VWP

Programm B F
ROM 11 Byte 1209
Ansprung durch CAL
anschliessend muss LAI 0006
1413
CUT 0615
RST 1 0013 programmiert
werden

Programm AF

ROM 04 Byte "504

Ansprung durch CAL 2
anschliessend muss RST 1 0013 programmiert

werden

Programm AX, ZP
ROM 02 Byte 1500
Ansprung durch CAL
vorher muss LLI 0306

0000

LMI 0314
Nrs die Bingabetabelle fuer
Fehlerwiederholung

Programm BF

ROM 11 Byte 0808

Ansprung durch CAL

vorher muss LAI 3 programmiert werden
01

nach UP-Abarbeitung muss RST 1 programmiert
werden

Programm BF

RON 11 Byte 0808
Ansprung duroh CAL
vorher muss LAI 0006

nach UP-Abarbeitung muss RST 1 programmiert
werden.

43

Loesohen des Programm UL
Rechenregistersi ROM 00 Byte 1308
Ansprung durch CAL
vorher muss LEI 0206
0015 (15 St.) programiert

werden
LLI 0306 programmiert
werden
0500
nach UP=-Abarbeitung muss RST 1 programmiert

werden

9% Hilfsprogramme bei der Arbeit mit Wikroprogrammen

Die nachfolgend aufgefuehrten Hilfsprogramme sind Neuerer-
leistungen und sind auf der Basis von Nachnutzungsvertraegen
verfuegbar'.

9.1. Testhilfe TEMIK
Kurzbeschreibung

TEMIK ist ein Testhilfsprogramm zum Test von Programmen bzw.
Programmteilen, die in der Maschinensprache des Mikropro-
zessors U808 bzw. INTEL 8008 geschrieben sind (Mikroprgramme)',
Durch TEMIK werden folgende Funktionen realisiert:

1. Schrittweise Abarbeitung des zu testenden Programmes mit
Druck und Stop nach jedem Schritt
2. Schrittweise Abarbeitung des zu testenden Programmes mit
Druck und Stop auf einer vorgegebenen Befehlsadresse
3. Abarbeitung des zu testenden Programmes ohne Druck mit Stop
auf einer vorgegebenen Befehlsadresse
4. Nach jedem Stop sind moeglich:
- beliebiger Wechsel der Abarbeitungsart
- Beginn mit neuer Anfangsadresse
- Veraenderung von Speicherinhalten (RAM-Inhalt, CPU-Re-
gister)
- Veraenderung der Flag-Einstellung, der Stellung des
%  Staock-Pointers und des Stack-Inhaltes
- visuelle Kontrolle (tetradendezimale numerische Anzeige)
von Speicherinhalten (RAM, ROM, CPU)

Damit kann eine dynamische Programmtestung bei einfacher Be-
dienung durchgefuehrt werden.

44

Bei Abarbeitung mit Druck erfolgt die Ausgabe je Abarbeitungs-
schritt von

- Befehlsadresse (RAM, ROM)

= Befehlsmnemonik Befehlscodierung

= Inhalt der CPU-Register A, B, C, D, E,H, L

nach Abarbeitung des Befehls

Inhalt des durch H L adressierten Speicherplatzes a)
Flag-Binstellung nach Abarbeitung des Befehls

Die zu testenden Programme koennen im RAM (00/0000..09/1115)
oder im ROM (gesamter Speicherbereich) des Magnetbandkon-
verters robotron 1255 gespeichert sein’;

AAN ntation von Mikroprogrammteilen DOKMIK

Das Programm Dokmik ermoeglicht die maschinelle Dokumen-
tierung von Programmen bzw. Programmteilen, die in der
Maschinensprache des Mikroprozessors U808 bzw. Intel 8008 ge-
schrieben sind und sioh im RAM (00/0600.'..10/1515) oder ROM
(gesanter Speicherbereich) des Magnet bandkonverters
robotron 1255 befinden‘.

Die Dokumentierung erfolgt in der Form

Befehlsmnemonik - Befehlscode - Befehlsadresse

Beispieli INH 0208 02/0000
IHI 0214 02/0004
- 0203 02/0002

Bei 2- und 3-Byte-Befehlen wird fuer das zweite bzw. dritte .
Byte in der Spalte "Befehlsmnemonik" das Minus-Zeichen (=)
gedruckt. Ebenso werden Konstantenbereiche dokumentiert.

Mit DOKMIK kann auch eine Dokumentation in der Form

Speicherinhalt - Speicheradresse
erfolgen (zB. Dokumentieren von Programmen ohne Nnemonik)'.

Beispielei 0906 03/1214
1311 03/1215
0001 03/1300

Der Druck erfolgt seitenweise. Die Zeilenanzahl/Seite und der
Abstand zwischen zwei Druckseiten sind beliebig vorgebbar. Der
Druck kann fortlaufend oder mit Stop nach jeder Seite
erfolgen; Die Abstaende zwischen den Druckspalten sind eben-
falls veraenderbar‘;

45

Su

71
HT |
HT |
107 |
TaT |
10 |
1ax |
Tas |
TV |
!
|
|
wo |
j
!

INT
IHT

IET |

1474

HHT
HaT
HET
H4d

HaS
HIV
Ln0
Ino
LnO
anI
LSU
Isa
ISY
SSU

zu/
anT |
HT |
gaT |
sat |
add |
aux |
aas |
aov |
l
l
I
!
l
I
t

ddy
218

04/
ORT
9HT
9qaT
Dal
9409
DuX
913
29V
du)
219
dd)
249

uvu
DuN

3poosnzxog-inyFeIydg !FTqarıeq
Zunpuezrun "yıauur Funzsrpoy-X
atyezag-ar4qryay-£'Z

yorstdzaı "0
zustearyuy *'X

ayazsqa+teH
aryazsqgdunıdg

a1yszeqsuorze4og

usuorzeıxedg ayosr307

tapo °°0 2

pun "MM atyozsqapeı

atTyaysqyususınu]

aTyszFsqsqedsny pun -urg

sogeırzıaqen saura "syoenzeg ru "as/'oaV

an sours "sxdanzsg auyo "ng/"qy STUPFSgsUoTPIeIIgng "n-SUOTITPPY

Tgunısertag
1 60/7 I 80/ I 20/1 90/1 ti. #0/ 1 £0/ I 20/ I 80/ 1 00/ |

I aut | VMI I MT TI I 377 | aTT | 97T I TI I VITI 5
| HT I HT | RITI TITl | a97 | aa | 99T I SdT I vadT | M
| gar I vari mri TOT! | MIT I MT I DOT IMTI WII &
| agı Ivami WII WI I WI IAvI I OovI I TI wii ©
I sd I vaı moi To I I guo | auo I duo I auo | vNo | bi
I aux | vux I man | Tan! ı ac I aan | dan | san | van | Or
jegasıvaı msı Tas | ans | ans | ons I ans I vas I 60
ı gaov I vov | Navı Tal {| gdav | aav | Day | aav | vav I '80
1 200 I due 2100 | \ I I ano I sad I ıno | SIE 20
I ıno | zi£ I 200 | ! ! I 200 I 0m I ano | DOLL I 90
I 200 I ddp I 200 | I I I 2no I sad Fan I Sar I So
I anı | Zap | anzi TWOl |  dHe f ANT I 040 I anı I Dat I #0
l I ! Ik. SIE I Ivo | sıu | | 9a I Inı I <o
| H0a I ANI | I IT I Ian I ouu | I a0a | aNI I 20
I ada I aNT | 420% I Ins | sau I vu | Dam I OnT I 10
Igoa tan | mi WII sei Tavı oa om 13H I 33H | 00
DLEL/GGEL uozzogoLt Iunzsrunreıgordoasit; ap yongqpuey ımz oFeTuy

’ Anlage zum Handbuch der Mikroprogrammierung robotron 1370/1255

oo
au

U ABLHEIHE

RBuNN sense IE.

Barren
ER SIIISEIEIENTTE
zac SESEIESEIETEIETEN
I [af no u [an u u
E14 Posten aealn [nm Dr ma u ao senfa Fa am
EI NTERFIETINNTINTITTRIENTTETEN
| [oo Por [7o2[7osfaos [rs ro rer Von os Inolam aa 13 zu zo |

ses Sala | fee

on
=
5
oc
Sale
mas)

=
un

rot

je E} + ri ad ohne Berlicks. eines Ü
Additions- u. Subtrak ionsbefehle< 0./SB, mit ® eines Übertr.
Ein- u. Ausgabebefehle

Inkrementbefehle ä
Ladebefehle N... Und
S| Logische Operationen —— 0: AieIvelene
FERN Rotationsbefehle —c.! Vergleich
Sprungbefehle 2,3- Mehrbyte-Befehle
Haltbefehles x- Codierung innerh. Umrandung -

beliebig; Schraffur=Vorzugscoo«

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
