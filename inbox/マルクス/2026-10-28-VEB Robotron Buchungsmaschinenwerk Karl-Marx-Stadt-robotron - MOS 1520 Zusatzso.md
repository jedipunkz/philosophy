---
source: "https://archive.org/details/mkl-20231220-02-edv-MOS-1520-Zusatzsoftware-unter-dem-Betriebssystem-SCPX-1526-Teil-II-01"
title: "robotron - MOS 1520 Zusatzsoftware II unter dem Betriebssystem SCPX 1526"
author: "VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt"
year: "1986"
captured_at: "2026-10-28T22:54:58Z"
updated_at: "2026-10-28T22:54:58Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "マルクス"
query: "Karl Marx"
plain_text_url: "https://archive.org/download/mkl-20231220-02-edv-MOS-1520-Zusatzsoftware-unter-dem-Betriebssystem-SCPX-1526-Teil-II-01/MOS_K1520_Zusatzsoftware_II_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "社会思想"
  - "唯物論"
status: raw
---

# robotron - MOS 1520 Zusatzsoftware II unter dem Betriebssystem SCPX 1526

- 著者: VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt
- 初版: 1986
- 情報源: [archive](https://archive.org/details/mkl-20231220-02-edv-MOS-1520-Zusatzsoftware-unter-dem-Betriebssystem-SCPX-1526-Teil-II-01)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[マルクス]]
- 研究動向: [[マルクス-現代研究動向]]

## Full Text

robotron

Systemunterlagendokumentation

MOS K1520

Anwenderdokumentation

Zusatzsoftware Il

unter dem Betriebssystem SCPX 1526

Systemunterlagen- Anwenderdokumentation

dokumentation

Stand: 30.6.85

Anleitung fuer den Bediener
fuer die
"Zusatzsoftware II unter dem

Betriebssystem SCPX 1526”

VEB Robotron Buchungsmaschinenwerk

Karl-Marx-Stadt 1985

nos

K1520

Die vorliegende 1. Auflage der - Dokumentation "Zusatzsoftuore
II unter dem Betriebssystem SCPX 13526” - entspricht dem Stand
Junı 1985 und unterliegt nicht dem Aenderungsdienst.

Nachdruck, jegliche Vervielfaeltigung oder Auszuege daraus
sind unzulaessig.

Die Dokumentation wurde durch ein Kollektiv des VEB Robotron
Buchungsmaschinenwerk Karl-Marx-Stadt ausgearbeitet.

Im Interesse einer staendigen Weiterentwicklung werden die
Leser gebeten, dem Herausgeber Ihre Vorschlaege bzw. Hinuelse
zur Verbesserung mitzuteilen.

Anmerkung:

Die Anuenderdokumentation zum SCP-System basteht aus:

= Anleitung fuer den Bediener SCP 1520

= Anleitung fuer den Programmiarer SCP 1520
Teil I und Teil II (Sprachbeschreibung ASM;

= Anleitung fuer den Systemprogrammierer SCP 1520
- Hardwarebeschreibung

= Anwanderdokumentation BASIC-Interprater

- Anwenderdokuwantation BASIC-Conpiler

- Anuenderdokumentation C-Compiler

- Anwanderdokumentat ion PASCAL

- Anuanderdokumentat ion FORTRAN

= Anuenderdokumentation Textverarbeltungssystam TP

= Anwenderdokumentation Install ierungs-Programm
fuer TP

- Anuenderdokumentation KOMBD-Druck

= Schulungshandbuch fuer das Textverarbeltungs-
system TP

- Zusatzsoftware I
- Zusatzsoftwore II
VEB Robotron-Buchungsmasch inenwerk
Karl-Marx-Stadt
Software-Zentrum
9010 Karl-Marx-Stadt/PSF 129

INHALISVERZEICHNIS

Treiber-Unterprogramme

Treiber fuer den Anschluss sahrerer

Handleseeinheiten HLE4 1520 (SCPX)
Hardware
Handleseinheit und Anschlussteuerung

Technische Daten der Handleseeinheit

Fehlererkennung

Die Plastkarte mit Magnetstreifen
Aufbau der Karten

Behandlung der Karten

Anwanderschnittstelle

Allgemeines

Funkt lonsbeschre bung

A-Bit-Code nach 150 3554 fuar Spur 2

spiel fuer Verwendung das
.

Treiber fuer 1/2°-Magnetbandgeraet
an AS130 MBG 1520 (SCPX)

Harduare
Anschlussteuerung AMB K5025
1/2°-Magnetbandgeraet CM 5300.01

180 1863-1976
Anuanderschnittste
Einstellen der Para
Aufruf des Traibars
Real sierung der Funktionen
Rusckmaldung des Traibars
Statusbyte

1/2°-MB-Treib
er

Badien- und Anzeigeele:
MBG CM 5300.01

Treiber Zusatzdrucker mit IFSS-
Interface ZUDRU 1520 (SCPX)
Sortieren und Mischen von Dateien
SORT 1520 (SCPX)

Uebersicht
Programmaufruf/Kommandoe ingabe
Kommando-Format

Reihenfolge der Kommandoeingaben
Aendern/Fortsetzen von Kommandos
Kommando-Zeilen
Kommando-Notat ion
Kommandobeschreibung
Eingabedatei (en)
INPUT-ATTRIBUTES

SORT-FILES

7
7
7
7?
8
8
8
7
2
[2
9
1
ı

2.3.1.3.
2.3.2.
2.3.3.
2.3.4.
2.4.
2.5.
2.6.
2.7.
2.8.
2.9.
2.9.1.
2.9.2.
2.9.3.
2.9.4.
2.9.4.1.
2.9.4.2.
2.9.4.3.
2.9.5.
2.10.
2.11.

3.

3.1.
3.2.
3.2.1.
3.2.2.
3.2.3.
3.3.
3.3.1.
3.3.2.
3.4.
3.4.1.
3.4.2.
3.4.3.
EI TE
3.4.5.
EIE FIT
3.4.7.
3.4.8.
3.4.9.
3.4.10.
3.4.11.
3.4.12.
3.4.13.
3.4.14.
3.4.15.1.
3.4.15.2.
3.4.16.
3.4.17.

MERGE-FILES
Ausgabedatei
Schluessel

Arbeitsdatei
TAGSORT
Ergebnisausschr ift
Weitere Kommandos

Selektion
Einfuehrung

Bereichstest

Weitere Selektionsmoegl ichkeiten
Selektion bei Nicht-ASCII-Daten
Spezielle Testattribute
Numerische Konstanten
Konstantenl iste

Veraenderung der Rangfolge
Anhang A: Formate

Anhang Bi Fehlerausschr iften

Universelles Dienstprogramm
DIENST 1520 (SCPX)

Einleitung
Kommandoe ingabesyntax des

DIENST

Das DIENST-Numsernmenue
Kommandosyntax ohne Nummernsenue
(sit Dateinauen)

Kommandosyntax ohne Nummernmenue
und ohne Dateinamen

Stauertasten

Steuertasten fuer Bildschirm
Stevertasten fuer Kommandoe ingabe
Kommandos des DIENST

? (Hilfsfunkt ion)

DUMP/DUMPX/DUMPH/DUMPA
ERA
EXIT
FILL
co
GROUP
JP/EX
LOAD
SAVE
LOG
MOVE
READ
WRITE

n

DND,ONNN
RUSBUnrOo

27.1.
27.2.
27.3.

un
so

30.

ui
Be

33.

Wassssshsu hass na
“uw
eNsas

3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3
3

READGR

WRITEGR

RECLAIM

REN

RESET

RUN

SEARCH

SET L-character]
SET

SET L+character]
SETDIR/SETSYS
SETRO/SETUR

SORT

SIZE

SPEED

STAT

TEST
TYPE/TYPEX/TYPEH/TYPEA
USER

XUSER

UR
Fehlermeldungen

Komprimieren einer SCP-Datei
CONF 1520 (SCPX)

Einleitung

Aufrufen des Programmes

Eingabe der Parameter der Qualldate
Eingabe der Parameter der Zieldatai
Programmende

Fehlerausschr iften

Zusammenhoenge des Dateiaufbaus
S10S-Datei und SCP-Datei bei dar
Bearbeitung mit UTRI und CONF

Drucken einer SCP-Datei
PRIF 1520 (SCPX)

Einleitung

Aufrufen des Programmes
Eingeben der Datelparanetar
Eingeben der Parameter fuer die
Druckl iste

Datelende

Fehlermeldungen

Mandeln einer SIOS-Textdatei fuer
die TP-Bearbeitung TRTX 1520 (SCPX)

Einleitung

Aufrufen des Programmes

Eingabe der Parameter der Quelldatei
Eingabe der Parameter der Zieldatei
Dateiende

Fehlernachr ichten

=.

8
90
90
2
92
93
94
95
96
96
96
97
97
98
98
9
4
100
101
101
102
102

105

105
105
105
106
109
109

110

112

112
112
112

113
115
115

117

117
117
118
118
120
120

Diskettenanalyseprogramm
LESD 1520 (SCPX)

Transformat ionsprogramm fuer
Text 20 - Dateien in Text 30 - Dateien
UTR3 1520 (SCPX)

Einleitung”

Parametereingabe fuer die

Text 20 - Diskette
Parametereingabe fuer die

Text 30 - Diskette

Fehlernachr Ichten und Meldungen
Fehlernachr ichten

Meldungen

122

125
125
125
127
129

1297
131

HLE4 1520 (SCPX)

1.___Iceiber-Unterprogrguug

1.1._Ireiber_fver_den_Anschluss_sehrerer _Handleseeinheiten
HLE4_1529_(SCPX)

1.1.1. Hardware

Aulelsle Handl Loheit_und_Anschlussteverung

Die Handleseeinhelt (HLE) K 6503 Ist ein Geraet zum Ver-
arbeiten von Plastkarten mit Magnetstreifen und an den

BC A5120/30 bzw. an Terminals anschl iessbar. s
Der Anschluss kann mit folgender Anschlussteuerung erfolgen:

K 6003.10 fuer 1-4 HLE Anschlusslaenge max.Sm
K 6003.11 fuer 1-2 HLE Anschlusslaenge max.5n
K 6003.20 fuer 1-2 HLE Anschlusslaenge 25m

Die Anschlussteuerung fuer den Handleser K 6003.10/.11/.20
dient der Kopplung der HLE K 4503.10 an die ZRE 2526 ueber
das Linieninterface BUS K 1520 und realisiert dabei den Da-
tanaustausch.

Es sind ueber je einen 10-poligen Steckverbinder 4 Handleser
anschl iessbar, die simultan bedient werden.

Aslals2s_Insbolssbe_Doten_dec_Handlesesinbelt

Datentraagar: Plastkarte mit Magnetstreifen nach
TGL 42091,42092,42093,150 3554,
DIN 9781,9785

Kopf: 1-Spur-Wiedergabekopf
Funktionsgruppen: -Kartenfuehrung =it Kopfhalterung
-Verstaerker

-Takt-Datentrennung
Schreibkapazitaet: Spur2: 37 numerische Zeichen

Schreibverfahren: Wechseltaktschrift nach DIN 66010

Auf ze Ichnungscod 4-Bit-Code (siehe Anlage 1)

Schrelbdicht,

3,0 Bit/mm

Kartengeschwind.: 10...100 cm/s

Abmassungen: 56 Hoche
75 mm Breite
180 mm Tiefe

Masse: 1,0 kg

Leistungsaufnahme: max. 5 U

HLE4 1520 (SCPX)

Speisespannung: +5 V#/- 5%
+12 U #+/- 5%

Die Geraete werden als Beistellbaugruppen (ohne eigene
Stromversorgung) ueber ein Kabel (Zubehoer) an die jeweilige
Anschlussteuerung angeschlossen.

Asls21.3._Feblererkennung

Beim Lesen der Karten erfolgt die Kontrolle durch die An-
schlussteuerung K 6003. Jedes aufgezeichnete Datenzeichen wird
durch ein Paritaetsbit auf ungerade Paritaet ergaenzt.

Jeder Datensatz wird durch ein LRC-Zeichen abgeschlossen,
das alle Bits gleicher Wertigkeit, ausser dem Paritaetsbit,
auf gerade Paritaet ergaenzt.

Das LRC-Zeichen selbst wird auf ungerade Paritaet ergaenzt.

Aululıdı_Dle_Plastkarte_wlt_Magnetztceife
Arleledalı_ Aufbau_der Karten

Als Dataentraeger wird die Plastkarte mit Magnetstreifen nach
150 3554 verwendet.

Die Karte besteht aus Kunststoff oder kunststofflaminiertem
Werkstoff und Ist einseltig mit einem Magnetstreifen varsehen.
Mit der HLE kann nur die Spur 2 gelesen werden.

Durchzugsr Ichtung

Spur 1)
Spur 2) JMagnetband
Spur 3)

Die Spur 2 wird im 4-Bit-Code nach ISO 3554 aufgezeichnet
(d.h. ein Zeichen entspricht einem 4-Bit-Element). Die Spur
Ist folgendermassen aufgebaut:

1 Startzeichen

37 numerische Zeichen
1 Stopzeichen
1 Pruefzeichen

HLE4 1520 (SCPX)

1z1:1,4.2. Behandlung_der_Karten

Sichtbare Verschmutzungen der Kartenaberflaeche sind vor der
Eingabe in die HLE zu entfernen.

Verletzungen der Magnetstreifenoberflaeche sind auf alle
Faelle zu vermeiden.

Zum Schutz vor Verschmutzung und mechanischer Beschaedigung
koennen die Magnetsreifen mit einer mehrere Mikrometer dicken
Deckschicht aus durchsichtigen, nicht magnetisierbarem Mate-
rial versehen sein.

Vor fremden Magnetfeldern sind die Karten zu schuetzen.

Aulu2.-Anwenderschnlttstelle

Aulı2ıla-Allogeeines

Der Treiber wird als REL-Datei bereitgestellt (verschlebliche
Ob Jektcodadatel).

Der Treibername ist HLE. Dieses Symbol Ist vom Anwender als
externes Symbol zu def Inieren.

Der Treiber Ist mit Call HLE aufzurufen.

Der Treiber rettet keine Register. Dies ist bei Bedarf vor
Aufruf des Treibers vom aufrufenden Modul zu geuaahrlaisten.
Der Treiber arbeitet Im Interruptbetrieb (INT-Mode 2) und
belegt Innerhalb der Interruptsaeule (OF740H - OF7DFH) 16
Byte.

Aala2ı2._Eunktlonsbsashrelbung

Der Aufruf der Funktionen des Treibers erfolgt nach folgendem
Schema:

Register C - Funktion 0...n

(Betr iebsbeginn mit Funktion 0)
Register DE - Paramater ( HLE NR. )
Register HL - Parameter ( Einleseadre
CALL HLE

fuer Fkt.1 )

Es sind folgende 4 Funktionen aufrufbar :

Funktion 0 - Initialisieren der HLE

Funktion 1 - Lesen

Funktion 2 - Statusabfrage

Funktion 3 - Loeschen der LED-Anzeige
.

HLE4 1520 (SCPX)

Fuer jede Funktion werden nachfolgende Registerzuuelsungen
vereinbart:

HLE in Grundstellung
(Initlialisierung)

ung an ALE

HLE NR.

I

I

I

I

HLE NR. I
|

I

I

Einleseadresse |
I

Statusadr

HLE NR.

I

I

I

I

D als auit- |

tierung des Lesevorganges j
I

I

j

durch das Anwenderprogramn)
HLE NR.

A=FF HLE nicht initiblisiert bzu. kaine Leseanfor-
A=FE Lesen initialisiert ‚aber noch nichts gelesen
A=FD Lesen Initialisiert ‚angefangen mit Lesen
Am 7F fertig mit Lesen / Fehler

A= 00 fertig mit Lesen / kein Fehler

I
I
I
l
| derung initlalisiert
I
1
|
I
I

(=== 2. Steckplatz = HLE 2

= HLE 1

HLE 0

HLE4 1520 (SCPX)

Die Funktionen 0,1,2 koennen in beliebiger Reihenfolge auf-
gerufen werden. Ist die HLE bei Aufruf C=1 noch nicht
initlalisiert, so wird dies vos Treiber automatisch aus-
gefuchrt.

Mit der Funktion 3 hat der aufrufende Modul die Hoeglichkeit,
die Lesequittierung selbst zu steuern. r
Wird die Funktion 3 nicht aufgerufen, so bleibt die LED nach
Initiallsierung staendig aktiv (LED leuchtet).

Mit Funktion 3 kann der Anwender nach Pruefungen (o.ae.) die
LED lo@schen.

In der Initiallsierungsfunktion wird die SCPX-Version 1.1. ge-
testet. Ist der Test negativ, so erscheint am Bildschirm die
Maldung: “Sorry SCPX 1.1.Cor upward) Is required |“
Gleichzeitig erfolgt ein Programmabbruch (Warmstart).

sn I=Blt-Code_nash_150_3554_Luec_Seuc_2

I

I

!

{I}
0000 o I o
o001 1 I 1
o010 2 I 2
oo11 3 I 3
0100 4 l 4
0101 5 j E
o110 % I 6
o1ıı 7 I 7
1000 8 1 8
1001 , f ,
1010 10 | wird nicht genutzt
1011 11 |  Startzeichen
1100 12 | wird nicht genutzt
1104 13 |  Leerfeldzeichen
1110 14 | wird nicht genutzt
1111 15 |  Endezeichen

1

Asls4ı___Eroscgmebslsplel_fver_Versenduns_des_HLE4-Icelbacs

#Programsbeispiel:

#Das nachfolgende Programm zeigt den von einer vorher aus-
jgewaehlten Handleseeinheit (0...3) gelesenen Inhalt einer
Karte auf.

Beim Lesen einer Karte wird nacheinander angezeigt:

; STATUS HLECNR.):
; DATEN HLECNR.IE

-11-

aseg
org 100h
bdos equ 0005h
wstart equ Oh
extrn hie
wied: 1d de,texo
1d <,
call bdos
wh! 1d de,texi
1d <,F
call  bdos
1d <,10
1d de,mkp
call bdos
14 a, (mkp+2)

5 Pruefung auf gueltige Eingabe

HLE4 1520 (SCPX)

;Warmstartadresse

jAusgabe Kopf

jAnweisung Taste druecken

;HLE-NR. auf Merkplatz mkp

fuer HLE NR. (0...3)

keine gueltige Eingabe

14 b,rar
<p b
Jr nc,wied
1d b,ror
cp b
Jr <,wiad
» Funktion Initialisieren einstellen
xor 30h
1d a
1d e),a
1d <,o
call hie

» Funktion Lesen einstellen

1d hi,nke
1d e,m

1d hl,burfx
1d c,1

call hie

Abbruch des Testprogrammes ?

1d <,6

1d e,orfh
call  bdos

<p 3

Jp z,ustart

14 hi,ake
1d e,n
14 <,2
call hie

42;

zAufruf HLE Treiber

FE S=HLE NR.
#HL:=Lasepufferadresse

sAufruf HLE Treiber

Direkt Console Input

st Control-C Taste gedrueckt?
a--)Warmstart / Abbruch

ILE NR.

zAufruf HLE Treiber

HLE4 1520 (SCPX)

Lesen beendet ?

bit 7,a ;Bit7=0?
jr z,aus
je rs jLesen nicht beendet/Uarten

; Ausgabe der gewaehlten HLE NR.

aus:  1d hl ,mke
14 e,m
push de
push af
1d a,ror
add a,e
1d (shle),a NR der HLE F.Status
1d dhle),a #NR der HLE f.Daten

; Ausgabe der Statusmeldung

1d de,tex2 STATUS?’

14 <,9

call  bdos

pop ar

ld de,teok zroor

or a

Jr z,out

id de,ter TFRnnFEHLERMuM"
out: 1 <,

call  bdos

5 Ausgabe "Daten:’

1d de,taxg "DATEN: ”
1d <,
call bdos

Funktion LED loeschen einstellen

pop de

1d <,3

call hie zAufruf HLE Treiber
; Ausgabe der Daten aus dem Lesepuffer

1d de,buffx

1d <,9

call bdos

ip wh

a

texo:

db
db
db
db
db
db
db
db
db

db
db
db
db
db
db

HLE4 1520 (SCPX)

12, 10,13, RR

10,13,  ’* * ar
10,13, ’* “
10,13, "= x TESTPROGRAMM fuer HLE 4"
10,13, '* “
10,13, '* “ Pr)
10,13, RABATT
10,13,24h

10,10,10,13,”==) GEBEN SIE DIE NR.’
"DER HLE EIN! (0...3),24h
1,0,0

o

37,0h

24h

10,10,10,13, "STATUS HLE’
"12 7,24

10,13, "DATEN HLE?

»1: ?,24h
"IERMMFEHLERRMMN?

voon’

wg

MBG 1520 (SCPX)
1.2. Treiber fuer 1/2"-Magnetbandgeraet an 45130

MEG_ 1520 (SCPX)

1.2.1. Hardware

1.2.1.1. Anschlussteuerung AMB K 5025

Die Anschluss-Steuerung AIB X 5025 dient zum Anschluss des
2'-iaznetbandgeraetes C!! 5300.01 ar. den Buerocomputer A5130

Bestan:
Anschlussplatte: 045-5071/81 DI{A-Steckeinheit
Steuereinheit: 012-7102 ZRE-K 2522
045-1070 Dekodiersteckeinheit
045-5069 CRC-Steckeinheit
045-8068 Koppelsteckeinheit

‚netband;

Dan Ilagnetbandgoraet CIi 5300.01 ist ein lansenspeicher fuer
sequenticllen Zugriff. Es kann auf Grund seiner Bigenschaften
ale externer Speicher fuer grosse Datenmengen dienen oder ala
Zwischendatentraeger fuer Grossrechner eingesetzt werden.
Durch den Anschluss des llagnetbandgeraetes an den BC ist die
Konvertierung von “inidisketten 5,25", Standarddisketten 8"
einfacher oder doppelter Bitdichte und lagnetbandkessetten auf
1/2"-jagnetoand gewaehrleistet.

("echnische Date

Auf zeichnungverfahren: 9 Spuren (150 1863-1976)
Aufzeichnungsdichte: 32 Bit/mm (800 bp1)
Vebertragungsgeschwind.: 10 köyte/s
Bandgeschwindigkeit: 0,3175 m/s
Nueckspulgeschwindigkeit: 1,5 m/s
ignetband; -Laenge: 370 m

Breite: 12,7 mm

Grund funktione:
Blockochrei ben
Dlocklesen

Schreiben von 2 Bandmarken u. Ruecksetzen in die Luecke dazwi-
schen

Vor-/Ruecksetzen um Bloecke
Vor-/Ruecksetzen ım Bandmarken (Dateien)
Rueckspulen

-5-

MBG 1520 (SCPX)
1.2.1.3. ISO 1863-1976

Dieser Internationale Standard trifft ?estlegungen ueber das
Format und die Aufzeichnung eines 1/2"-Hagnetbandes.

Magnetband - Spurendarstellung

Blockzwischen- Auf zeichnungs-
raum Block beginnzone
1-2 I >|

! I |
| Position I
| IRC-Zeichen |
>11 — I
I | Position |

>|-I<- Spuren/

ı

I-

| 118-2048 Informations- Spurnummer
l

Izeilen

I —

Bezugs-
kante |

!
I I
! I
I I
I I
[| CRC-Zeichen 1 |
| I
I I
| I
I [

> Vorlaufrichtung des Bandes
— Ruechklöufrichtung des Bandes

"I1l...11" - Datenbits

- Identifizierung der Spuren:

Spur des „.agnetbandes: 1 34560709
Nortigkeit: PS Zur Dur Br 7 Eu Gr Ber Br

p ist das Pruefbit. Die Zahlenparitaet ist ungeradzahlig.
Nie anderen 0 Bit bilden eine Informationzzeile.

cstlezu

Breite einer aufzezeichneten Spur:
Zeilenn.ustand:

Spurenadstand:
Dlockzwischenraun:

Auf zeichnunfsbeginnzone:

AvAv

= 416=

MBG 1520 (SCPX)

Auf zeichnungsdichte: 32 Zeilen/rm
Auf zeichnungsgeschwincdigkeit: 31,75 em/s
Reflektormarken

Jede Bandspule besitzt 2 lichtreflektierende liarken.

Bandanfangsmarke BOT : - Abstand 4,9 +0,6 m vom Bandanfang
-0,5 m
Bandendemarke BOT : - Abstand 7,6 +1,5 m vom Bandende
°
Bandmarke

Zur Trennung der Dateien wird ein einzelner Kontrollblock (Bi)
geschrieben. Dieser Block wird von einem LRC-Zeichen beglei-
‘ tet, Die Bandmarke wird durch "1" in den Spuren 2, 3 und 8
dargestellt (alle anderen Spuren 0).

Blookaufbau

Jeder Block besteht aus 18 bis 2048 (2 K Byte) Infornations-
zeilen. Am Ende jedes Blockes wird das CRC-Zeichen (Kontroll-
zeile der zyklischen Redundanz) zur Fehlererkennung zuf einer
Spur geschrieben.

Im Anschluss an das CRC-Zeichen wird das LRC-Zeichen (Laenes-
pruefzeile) zur spaeteren Pehlererkennung beim Lesen aufge-
zeichnet. Ein Laengopruefzeichen wird auf den Spuren markiert,
deren Zeichenanzahl ungerade ist.

Int dem LRC-Zeichen ist jedoch nur eine geringe Fehlerrate or-
ennbar.

1.2.2. Anwenderschnittstelle 1/2"-MB-Treiber

Festlegungen:

- Der Treiber wird als verschiebliche Datei (Typ .REL) bereit-
gestellt. Er ist im absoluten Modus mit ORG-Anweisung pro-
grammiert worden.

ASEG
ORG 103H

Es ist keine Interruptarbeit erforderlich. Die AllB arbeitet
im Batch-Betrieb.

- Dem Treiber muessen vom Anwender folgende Parameter in den
CPU-Registern bereitgestellt werden:

in Register C - Funktionscode FC
in Register DE - Blockadresse BA
in Register HL — Blocklaeng: BL

— Rueckmeldungen des Treibers an das aufrufende Programm wer-
den

Pr

MBG 1520 (scpx)

in Register A - Statusbyte
in Register HL — Brgebnisvektor bereitgestellt.

- Der Name des Treibers !BG ist fuer das aufrufende Programm
das externe Symbol.

Das aufrufende Programm muss mit END <Progr.-anf.-marke> be-
endet werden, damit der Programmverbinder auf die Adr. 100H
einen Sprung schreibt, der auf den logischen Programmanfang
zeigt (auch mit G-Schalter moeglich - siehe LINK).

1.2.2.1. Binstellen der Parameter

Vor Aufruf der einzelnen Funktionen deg Treibers sind folgende
Parameter einzustellen:

in Register C - Funktionscode FC

= FC fuer die einzelnen Purktionen:

Funktion

-1-

MBG 1520 (SCPX)
- Bitbelezung des Funktionscodebytes:
T T T T T T T

! I
I0o-vor 10- 10- 10- 10- 10- 10- 10- |
setzen |1- \11- 1- 11- 11-BH  11- I1- l
I1- lohne |rueck-IBi |Block Ischr. |Block |Block |
Ieeen Ischr. |

|

Irueck-|Bedeu-|spulen|Steu- |Steu- |

[setzen | tung | lerbit jersit j
SE BEER Fat

| Bit7 | Bit6 | Bit5 | Bit4 | Bit3 |

! } ı \

in Reg; r Blookadresse BA

Die Blockadresse BA muss bei folgenden Funktionen vom Anwender
bereitgestellt werden:

—_ Lesen:

BA ist die Adresse im Speicher, auf die die vom Magnetband ge-
lesenen Daten uebernommen werden (E/A-Bereich).

> Schreiben:

BA ist die Adresse im Speicher, von der die Daten auf das Nag-
netband geschrieben werden (E/A-Bereich).

in Register HL - Blocklaenge BL

Die Blocklaenge BL wird bei folgenden Funktionen benoetigt:

Lese:
BL ist die Anzahl der zu lesenden Bytes.
- Schreiben:
BL ist die Anzahl der zu schreibenden Bytes.

= Block vor/rueck:

BL ist die Anzahl der Bloecke, die vor- oder rueckgesetzt wer-
den sollen.

Datei vor/rueck:

BL ist die Anzahl der Dateien, die vor- oder rueckgesetzt wer-
den sollen.

1.2.2.2. Aufruf des Treibers

Der Aufruf des Treibers erfolgt nach Parametereingabe durch
CALL MBG.

-19 -

NBG 1520 (SCPX)

ıD 0 ,rc
ID DE,BA
ID HL,BL
CALL MBG

1.2.2.3. Realisierung der Funktionen

- Block schreiben

Der durch die Blockadresse BA gekennzeichnete Speicherbereich
wird ohne Zwischenspeicherung mit Datensicherungsmassnahmen
auf Band geschrieben.

Die Aufzeichnungskon‘ lle erfolgt durch Bildung des CRC-Zei-
chens fucr Lesedaten. Werden Aufzeichnungsfehler erkannt, wer-
den 3 weitere Aufzeichnungsversuche auf der gleichen Bandstel-
le und je 4 auf 3 um ca. 100 mm vorgezogenen und geloeschten
Bandstellen unternommen. Bei negativem Ergebnis wird Bit7 im
Statusbyte gesetzt.

Das Aufzeichnungssystem befindet sich nach Abbruch vor dem
fchlerhaft aufgezeichneten Block, so dass dieser anachliessend
ueberschrieben werden kann.

Befindet sich die BOT-llarke unter der Abtesteinrichtung, er-
folst durch die AlIB automatisch die Aufzeichnung der Initial-
luocke (400 mm) vor dem Aufzeichnen des Datenblocks,

Bei Eriennen der EOT-Harke wird Bit4 im Statusbyte gesetzt.
" Ist keine ilagnetbandspule eingelegt oder der Rings zum Auf-
weichnungsschutz abgenommen, wird im Statusbyte Bit6 gesetzt,

= Bloe! lesen

Nie Bytes verden vom Band gelesen und in den durch die Block-

advense BA gekennzeichneten Speicherbereich eingetragen. Die

gelesene Dytezahl wird im Ergebnievektor angegeben,

Ict der Block auf dem llagnetband Inenger oder kuerzer ala der

mugewiesene Speicherbereich, werden die restlichen Byteo ze-

mehlt und Bit2 in Statusbyte markiert. Im Ergebniovektor wird

die tatsnechlich gelesene Byteanzahl von der AllB eingetragen.

Die Naten sind erst nach gegebener Quittung gueltig.

Die Datensicherung bei der Hiedesgabe erfolgt durch Bildung

und Kontwolle des CRC- und des LNF-Zeichens.

Neil iHecevgnvefehler worden bis zu @ \iedergabeversuche unter-

nommen. Bei negativem Ergebnis wird Bit7 im Statusbyte ge-
t besystem befindet sich dabei vor dem feh-

ock, d.i. bei erneuten Lesekommando vird der Lese-

wiederholt. Der gelesene Block steht im
poicherbersieit zur Yerfu ann durch dus Kommando
"3loe': vorsetzen" ueberleser

erkannt, dann wird vor
usbytc Bit5 belest. Die
"Block vorsetzen" oder "Da-

sositioniert
senn duren das X
tzen" ueberlesen :

eichnungsvorgeng behandelt.

MBG 1520 (SCPX)

- Bandmarke schreiben

Es werden 2 Bandmarken nach ISO 1863 aufgezeichnet. Danach
wird das Band zwischen beide Bandmarken positioniert, um die
Aufzeichnung’ einer weiteren Datei zu ermoeglichen.

Die Fehlerbehandlung erfolgt wie beim Block schreiben.

Da beim Ruecksetzen um 1 Bandmarke diese zur Kontrolle gelesen
wird, ist im Statusbyte bei korrekter Ausfuehrung Bit5 ce-
setzt.

Ist keine Nagnetbandspule eingelegt oder der Aufzeichnungsring
abgenommen, wird Bit 5 im Statusbyte gesetzt.

- Vor-/Ruecksetzen um Bloecke

Das Kommandobyte "Block vor-/ruecksetzen" bewirkt die Funktion
der Blockpositionierung mit Arbeitsgeschwindigkeit. Entspre-
chend der vorgegebenen Blockanzahl BL wird das Hagnetband

- um n Bloecke vorgesetzt, wenn Bit7 von BL (HIGH) Null ist
oder
- um n Bloecke rueckgesetzt, wenn Bit7 von BL (HIGH) 1 ist.

Der Lesevorgang wird zeitlich ueberwacht. Bei Veberschreiten
der zulaessigen Maximalluecke (>7500 mm) wird Bit7 im Status-
byte gesetzt.

Nach Funktionsausfuehrung wird der Ergebnisvektor versorgt und
das Statusbyte aktualisiert. Entspricht die Anzahl der vor-/
rueokgesetzten Bloecke nicht der durch BL geforderderten, dann
wird Bit2 im Statusbyte gesetzt und die tatsaechlich vor-/
rueckgesetzte Blockanzahl in den Ergebnisvektor eingetragen.

a) Bedingungen beim Block vorsetzen:

Wird eine Bandmarke erkannt, wird die Funktion abgebro-
chen. Das lagnetband steht vor dem 1. Datenblock der
naechsten Datei in Vorwaertsrichtung. Die aktuelle Anzahl
dor vorgesetzten Bloecke wird im Ergebnisvektor angegeben.
Im Statusbyte wird Bit5 belegt.

Bei Erkennen der EOT-Marke wird die Funktion fortgefuehrt
und im Statusbyte Bit4 belegt.

b) Bedingungen beim Blockruecksetzen:

Wird eine Bandmarke erkannt, wird die Funktion abgebrochen.
Das MNagnetband steht vor der Bandmarke in Vorwaertsrich-
tung. Bei Erkennen der BOT-Marke wird die Funktion eben-
falls abgebrochen.

In beiden Faellen wird die aktuelle zurueckgesetzte Block-
anzahl in den Ergebnisvektor eingetragen und im Statusbyte
Bit3 bzw. Bit5 gesetzt.

-21-

MBG 1520 (SCPX)

Das Magnetband wird um n Bandmarken vor- bzw. rueckgesetzt.
Wird beim Vorsetzen das logische Bandende erkannt (zwei durch
Blockluecke getrennte Bandmarken), wird die Funktion abgebro-
chen, die Anzahl der vorgesetzten Dateien im Ergebnisvektor
bereitgestellt und im Statusbyte Bit5 belegt. Danach wird das
Magnetband zwischen beide Bandmarken positioniert.

Wird beim Ruecksetzen die BOT-Harke erkannt, wird die Funktion
abgebrochen und im Ergebnisvektor die Anzahl der rueckgesetz-
ten Dateien bereitgestellt. Im Statusbyte wird Bit3 gesetzt.
Nach Dateiruecksetzen steht das Magnetband in Vorwaertsrich-
tung vor dem 1. Block der Datei.

Wird bei Operationen die BOT-Harke erreicht, wird in Status-
byte zusaetzlich Bit4 gesetzt.

Bei Nichtuebereinstimmung von geforderter Dateianzahl und tat-
saechlicher Anzahl wird Bit2 im Statusbyte gesetzt und die ak-
tuelle vor-/rueckgesetzte Dateianzahl in den Ergebnisvektor
eingetragen.

Da die Funktion im Normalfall nit Erkennen einer Bandnarke ab-
geschlossen wird, ist im. Statusbyte Bit5 belegt. B

 Rueckspulen

Das Magnetband wird bis zur BOT-Harke zurueckgespult. Daa Er-
gebnis wird im Statuspyte bereitgestellt.

Der Umspulvorgang wird durch die AI'B zeitlich weberwacht. Bei
Veberschreiten der Umspulzeit wird Bit? im Statuebyto ge-
setzt. Das Umspulen wird nach Brreichen des Bandanfangs quit-
tiert.

1 Rueokmeldungen des T: 2

Nach Ausfuehrung der Punktion gibt der Treiber folgende Rueck-
meldungen:

- in Register A - Statusbyte (siehe Pkt. 2.2.)

- in Register HL - Brgebnisvektor (Anzahl der gelesenen Eytee
bzw. Anzahl der vor-/rueckgesetzten Bloccke/Dateien)

-22 =

MBG 1520 (SCPX)
1.2.2.4.1. Rueckmeldung im Statusbyte

Bitbelegung des Statusbyte:

T T T T = T I
0- 10- 10- 10- 10o- 1o-BL  lo-MBG l0-1BG |
4-P/U-|1-Auf-11-BM  |1-BOR |1-BOT |=Brgeblin Ru-|bereit|
Irehlerlzeich-| e- | er- | er- Inisveklhe 1-1BG |
jerae-Inungs-| kannt| kannt! kanntitor I1-NBG Inicht |
tefeh-Iver- | I I larbei-|bereit|
Iler Iletz. | I I I1-un- Itet | I
j ! I I I Igleichl ! |

| I- I I

! I I I

| ! \ 1}

Jie ersten beiden Bits sind nach Punktionsausfuehrung immer
nit Null belegt.

Bit 0=-0
Bit 1=0

der Treiber wartet intern immer auf Funktionsbeendigung, d.h.
Sit! = 0 ist nach Rueckkehr aus dem Treiber immer eingestellt.
Sollte das IB-Gerzet nicht bereit sein (nicht eingeschalten,
ein Band eingelegt), so ist BitO = 1 gesetzt.

Stotusmeldungen der einzelnen Punktionen

- Schreiben:

Code im Statusbyte 1) Bedeutung
|

- Punktion ausgefuehrt

- EOT-Harke erkannt

Fehlermeldungen:

I
|
|
|
[|
h
4OH | - Aufzeichnungsverletzung
|  (Aufzeichnungsschutzring abgenom-
l men oder keine IB-Spule eingelegt)
I
I
|)

- falscher Funktionscode
- Aufzeichnungsfehler
- Geraetefehler

Bedeutung

Funktion ausgefuehrt

-23-

MBG 1520 (SCPX)

= Wiedergabefehler

- Geraetefehler

- Veberschreiten zulaessige Maxi-
malluecke (7500 mm)

Bandaacke_ (Bl) schreiben:

I
ı
04H | - vorgegebene Byteanzahl_ BL ungleich
|  tatsaechlich gelesene Byteanzahl
l
14H | - EOT-Marke erkannt,
| vorgegebene Byteanzahl BL ungleich
|  tatsaechlich gelesene Byteanzahl
l
24H | - BM erkannt,
| vorgegebene Byteanzahl BL ungleich
|  tatsaechlich gelesene Byteanzahl
I
80H | - falscher Funkt ionscode
I
I
1
I
I

=
Blesk_versetzen:

Code Im Statusbyte Bedeutung

Code Im Statusbyte I Bedeutung
bi ---|- -

20H | = Funktion ausgefuehrt, BM erkannt
l
I

Eehlecmeldungen: l

10H | - EOT-Narke erkannt
I

AOH | - Aufzelchnungsverletzung
| (Aufzeichnungsr Ing abgenommen oder
| keine MB-Spule eingelegt)

soH | - falscher Funkt ionscode
| - Aufzeichnungsfehler
| - Geraetefahlar
!
I
I
I
I
I

00H = Funktion ausgefuehrt

Esblermeldungenz

'

EOT-Marke erkannt,

vorgegebene Blockanzahl BL un-
‚gleich tatsaechlich vorgesetzte
Blockanzahl

I
l
I
14H I
j

- 24 -

MBG 1520 (SCPX)

- Geraetefehler

Veberschreiten zulaessige Naxi-
|  malluecke (7500 mm)

Bensenanuneennansnnennen | auenennnenn

- Block ruecksetze:

24H | - BH erkannt,
| vorgegebene Blockanzahl BL un-
| gleich tatsaechlich vorgesetzte
| Blockanzahl
| (Band steht nach BN, am Anfang
N der neuen Datei)

80H | - falscher Funktionseode
|
|

Oode im Statusbyte

- Funktion ausgefuehrt

00H

20H - Funktion ausgefuehrt, BM erkannt

Fehlermeldungen:

OCH - BOT-Harke erkannt,
vorgegebene Blockanzahl BL un-
gleich tatsaechlich rueokgesetzte

Blockanzahl

- BM erkannt,
vorgegebene Blookanzahl BL un-
glsieh tatsaechlich rueokgesetzte
Band steht an Dateianfang)

80H

- falscher Funktionsoode
-6 tefehl

Datei vorsetzen

Saas im Stataabyte Bedeutung

- Punktion ausgefuehrt, BM erkannt

Fehlermeldungen:

14H - BOT-Marke erkannt,
vorgegebene Dateianzahl BL un-
gleich tatsaechlich vorgesetzte

Dateianzahl

-35-

NBG 1520 (SCPX)

24H | - BM erkannt,
| vorgegebene Dateianzahl BL un-
| gleich tatsaechlich vorgesetzte
H Dateianzahl (logisches Bandende)
80H | - falscher Funktionscode

| - Geraetefehler

I
Code im Statuabyte I Bedeutung

— Datei ruecksetzen:

Fehlermeldungen:

OCH - BOT-Marke erkannt,
gleich tatsaechlich rueckgesetzte
Dateianzahl

80H falscher Punktionscode

- Geraetefehler

= Rueckspulen:
Code Am Statusbyte

|
I
I
1
I
I vorgegebene Dateianzahl BL un-
1
|
I
|
|
I
!

Bedeutung

- Punktion ausgefuehrt, BoT- Marke
erkannt

Fehlermeldungen:

80H - falscher Funktionscode
- Geraetefehler
I

Rueokmeldun;

— Sohreiben
Es erfolgt keine Rueckmeldung im Ergebnisvektor.
— Lesen

Von der AMB wird die Anzahl der gelesenen Bytes in den Ergeb-
nisvektor eingetragen.

-26 -

MBG 1520 (SCPX)
- BM schreiben
Es erfolgt keine Rueckmeldung im Ergebnisvektor.
- Block vor-/ruecksetzen

Von der AMB wird die Anzahl der vor-/rueckgesetzten Bloecke
eingetragen.
- Datei vor-/ruecksetzen

Von der AMB wird die Anzahl der vor-/rueckgesetzten Dateien
eingetragen.

— Rueckspulen
Es erfolgt keine Rueckmeldung im Ergebnisvektor.

lemente des MBG CM 5300

bei Einschalten wird MB-Einheit an-

- Netztaste:
geschlossen

leuchtet bei eingeschalteter Strom-
Versorgung

„Anzeige "Netz":

- Taste "Laden - bringt den Pufferhebel in den Ar-
beitsbereich und atellt "on-line"-
f Zustand her
- Band wird nach vorn gespult bis zum
3 AFPY3KAı Erkennen der BOT-Marke
bei Erkennen dieser Marke leuchtet
diese Tante auf
- ist bei Netzeinschaltung Magnetband
eingelegt, so ist nach "Laden" der
Vorwaertslauf des NB durch "STOP"
abzubrechen und durch "Rueckopulen"
auf die Harke "BOT" zu positionieren

'

=

TDUCTAHL |

=_-Anzeige.
nungssci

1 3RauTa |
SANMCH

=_Iaste_/Rueckspulen”:

MBG 1520 (SCPX)

durch Druecken dieser Taste wird
logische Verbindung zwischen Rechner
und Magnetbandeinheit hergestellt
die Steuerung erfolgt ueber Inter-
faceleitungen, die Taste leuchtet
auf

Umschaltung von “on-line” auf “off-
line” durch Taste “STOP”

leuchtet auf, wenn keine MB-Spule
eingelegt ist oder wenn der Ring zum
Aufzeichnungsschutz abgenommen ist

es erfolgt
zur BOT-Narke
danach wird
ausgefuchrt
Pufferhebel gehen in Endlage zurueck
und Band wird ganz abgespult
alle Operationen, ausser
werden blockiert

schneller Ruecklauf bis

Operation "Entladen”

"STop”,

Bandtransport wird unterbrochen, In-
terfaceleitungen blockiert
Geraet wird In Betriebsart
line” gebracht

“off-

Bandvorlauf, bis zum
EOT-Marke

waehrend der Betriebsart "on-line”
ist diese Taste wirkungslos

Erkennen der

Bandruecklauf bis
BOT-Marke
waehrend der Betriebsart "on-Line“
ist diese Taste wirkungslos

zum Erkennen der

-28 -

ZUDRU 1520 (SCPX)

tzdruoker mit IPSS-Interfaoe

1.3. Treiber fuer 2:
u ZUDRU_ 1520 (SCPX)

Dieses Unterprogramm dient zum Aufruf eines Zusatzdruckers mit
IPSS-Anschluss und muss in Anwenderprogramme eingebunden
werden. Mit Hilfe dieses Programmes kann eine Zeichenkette
ausgedruckt werden. Innerhalb dieser Zeichenkette sind alle
druckbaren Zeichen und Steuerzeichen des Druckers zugelassen.

Folgende Register muessen im Anwenderprogramm voreingestellt
werden:

HL --> Adresse der auszugebenden Zeichenkette
BC --> Anzahl der auszugebenden Zeichen (0 kein Zeichen)
Der Aufruf des Unterprogrammes erfolgt durch:

CALL ZUDRU.

Rueckmeldung im Register A:

<a> = 00 (HEX)
<A> = FF CHEX)

> kein Fehler
> Druckerfehler

Im Fehlerfali leuchtet LED auf. urch Betaetigen einer
beliebigen Taste ist die Weiterarbeit fast immer moeglich
(z.B. falsches Steuerzeichen fuer Drucker!). Bei Papierende-
meldung und technischen Stoerungen des Druckers muss der Feh-
ler erst behoben werden. Sonst erfolgt die Meldung:

<A> = PF CHEX).

Das Unterprogramm wird unter der Bezeichnung ZUDRU .REL
bereitgestellt.

-29 -

SORT 1520 (SCPX)
Sortieren und M

CPX

SORT 1520 (SCPX) ist ein Dienstprogramm zum Sortieren und
Mischen von Dateien, die auf Disketten gespeichert sind. Das
Programm hat eine Laenge von 20 K Byte.

2.1. Ubersioht
Sortieren/Mischen

- Sortieren von maximal 32 Eingabedateien in eine Ausgabeda-
tei;

- Mischen maximal 32 vorsortierter Dateien;

- gerhinzen und Mischen von Dateien kann im gleichen Lauf er-
olgen;

- unabhaengig vom Sortier-/Mischvorgang koennen eine Satzse-
lektion, eine Dateikonvertierung und weitere Operationen
ausgefuehrt werden.

Datei- und Satzformate

- Verarbeitung von SCP-Dieketten-Dateien;

- Dateien koennen ASCII-, BCD- und/oder Binaerdaten enthal-
ten; f

- maximale logische Satzlaenge betraegt 4096 Zeichen;

- Dateien koennen Saetze mit fester oder variabler logischer
Satzlaenge enthalten;

- die als Schluessel (Sortiermerkmal) festgelegten Felder
koennen eine feste Position im Satz besitzen oder eine
variable, wenn der Satz aus durch Kommas getrennten Feldern

eht.

Sortierschluessel

- 1 bie 32 Felder koennen als Sortierschluessel angegeben
werden; der Schluessel mit der hoecheten Prioritaet wird zu-
erst angegeben;

- Sohlue ldaten koennen als ASCII-Text, ASCII-numerisch
oder binaer interpretiert werden;

- Zusaetzliche Moeglichkeiten bei der Sohluesselinterpreta-
tion sind das Umwandeln von Klein- in Grossbuchstaben, das
Ignorieren des hoeohstwertigen Bits eines Bytes und andere,

Satzselektion

- Saetze koennen nach vorgegebenen Kriterien aus Eingabeda-
teien selektiert werden;

- äie Selaktion wird durch Vergleich eines Feldes nit einen
konstanten Wert, einer Folge konstanter Werte oder einem
anderen Feld im gleichen Satz spezifiziert;

- Vergleichsoperatoren sind KLEINER GLEICH, GLEICH usw., aber
auch ENTHALTEN und NICHT ENTHALTEN. Die Vergleiche koennen
wahlweise mit den logischen Operatoren UND, ODER und
EXCLUSIV-ODER kombiniert werden;

- Zur Satzselektion koennen Felder fester oder variabler
Position verwendet werden.

- 30 -

SORT 1520 (SCPX)
weitere Noeglichkeiten

- Teildateien:
fuer jede Sortiereingabedatei koennen Anfangs- und End-Satz-
nummer angegeben werden;

= Uechsel Ausgabediskette:
ermoeglicht Sortieren in einem 2-Laufwerks-System ohne die
Eingabedatei zu ueberschreiben;

- TAGSORT:

Abhaengig von der zu saortierenden Datenmenge und der
Schluessellaenge wird eine temporaere Disketten-Arbeitsdatel
angelegt. Durch TAGSORT wird ein anderer Sortieralgor Ithmus
verwendet, der zwar die Sortierzeit erhoeht, aber keine Ar-
beitsdatel erfordert bzw. den fuer die Arbeitsdatel benoe-
tigten Diskettenbereich reduziert (2- bzw. 1-Laufwerkskonf i-
gurat onen);

- Konsolen-Ausgabe:
Auswahl unter sechs Moeglichkeiten der Ausschrift von
Informationen ueber den Sort ier-/Mischlaufz;

- Zusatzfunktionen:
Konvertieren einer Datei eines bestimmten Satztyps in eine
Date! anderen Typ
Veraenderung der Satzlaenge, Ausblenden bzw. Umpositionlaren
von Feldern.

Zusaetzliche Ausgab

oegl ichkeiten

- Nur-Schluassel:
Ausgabedatei enthaelt nur die Sortierbegriffez
kann zum Aufbau einer Indexdatel, zur Summenbildung oder
zur Selektion oder zum Umpasitionieren von Saetzen genutzt
werden;

- Satznummern:
Ausgabedate| enthaelt nur die Satznummern der Einyabedatelen
Im ASCII- oder Binaerformatz
die entsprechenden Schluessel koennen wahlweise mit gespel-
chert werden;

- Zeiger:
Ausgabedate| enthaelt Sektornummern und relativen Satzbaginn
(Offset), mit oder ohne Schluesselspeicherung.

2224 _Prgsrammaufruf/Komsandgelnsahg

Die von SORT 1520 (SCPX) ausfuehrbaren Funktionen werden durch
Kommandoeingaben aktiviert. Die Kommandos koennen einzeln,
zusammengefasst zu einer Kommandodatei oder in der SORT aufru-
fenden Systemzeile eingegeben werden.

Jede Kommandozeile wird durch (ENTER) beendet.

==

SORT 1520 (SCPX)
Im einfachen Fall wird SORT in der Systemzeile durch
SORT

geladen und zeigt durch Anzeige eines "*" seine Bereitschaft
zur Kommandoeingabe an.

Einige oder alle Kommandos koennen in der Systemzeile nach dem
Programmnamen "SORT" eingegeben werden, z.B.:

A>SORT INP=80,CR; SO=EIN. DAT; OUT=AUS.DAT;KEY=f#2,4;G0

Waehrend des Sortier-/Hischlaufes ist keine Tastatureingabe
moeglich.

2.2.1 ormat

Jedes SORT-Kommando besteht aus einem Schluesselwort, das den
Kommandotyp festlegt, gefolgt von Parametern. Abhaengig vom
Kommandotyp bestehen die Parameter aus Dateinamen, Zahlen und
Schlue sselwoertern.

Schluesselwoerter bestehen aus Buchstaben und Bindestrich "-"
und besitzen eine Lang- und mehrere Kurzformen.

Buchstaben koennen als Klein- oder Grossbuchstaben eingegeben
werden. Im folgenden werden die Schluesselworte vorrangig in
ihrer selbstdokumentierbaren Langform dargestellt. llaehrend
diese Form leicht lesbar ist, benoetigt die Kurzform weniger
Eingabeaufwand. Die Kurzform wird gebildet, indem Buchstaben
und/oder der Bindestrich am Ende des Schluesaelwortes wegge-
lassen werden.

Ale minimale Anzahl Buchstaben eines Schluesselwortes muessen
eingegeben werden: ein oder zwei Buchstaben bei den meisten
Kommandos; drei, wenn das Wort mit "NO" beginnt.

Drei oder mehr Buchstaben werden stets akzeptiert.
2.B. sind folgende Bezeichnungen aequivalent:

INPUT-ATTRIBUTZ
INPUT

ım
1

Leerzeichen koennen anstelle von Kommas oder dem Cleichheits-
zeichen nach dem Kommandoschluesselwort verwendet werden.
Leerzeichen koennen beliebig zwischen den \lorten eingefuegt
werden.

2.2.2. Reihenfolge der Kommand:

schlauf koennen
gibt zei Aus-

Die einzelnen Kommandos fuer einen Sortier-/if:
in beliebiger Reihenfolge eingegeben werden.

Pe

SORT 1520 (SCPX)

nahmen: Scluessel messen in der gewuenschten Folge eingegeben
werden, das Kommando G0 zum Start mss zuletzt angegeben
werden.

2.2.3. Aendern/Fortsetzen von Kommandos

Werden Kommandoeingaben wiederholt, gilt die letzte Bingabe.
Dies ermoeglicht das spaetere Aendern bereits eingegebener
Kommandos.

Kommandos mit vielen Eintragungen (z.B. KEY=) koennen
erveitert werden, wenn man anstelle des Gleichheitszeichens
=" das Pluszeichen "+" angibt.

2 Kommando-Zeilen

SORT-Kommandos werden ueblicherweise einzeln und in einer
Zeile eingegeben. Jedoch sind auch folgende Varianten erlaubt:

- mehrere Kommandos in einer Zeile:
Kommandos muessen durch Semikolon getrennt werden;

Fortsetzung eines Kommandos in einer Folgezeile:
Eingabe eines Ampersand "&" vor <ENTER>;

- Kommentare :

ein Vertikalstrich "|" definiert den Rest der Zeile als
Kommentar.

5. Kommando-No' jon

Zur Definition der Kommandos gilt folgende Notation:

INPUT Grossbuchstaben fuer Kommandoachluessel-
woerter

dateiname Kleinbuchstaben fuer zu ersetzende Be-
griffe

ce» Bingabezwang fuer eine der uebereinan-
derstehenden Iloeglichkeiten

E)] Eingabe kann weggelassen werden

GH keine Eingabe oder Eingabe der ange-

fuehrten Alternativen in beliebiger
Anzahl und Reihenfolge, getrennt durch
Komma oder Leerzeichen

beliebige Anzahl von iliederholungen der
vorangestellten Parameter

= 33.-

SORT 1520 (SCPX)

2.3. Kommandobeschreibung

Line rliche Erlaeuterung der Datei- und Datentypen
enthaelt Anhang A.

2.3.1. Eingabedatei(en)

Drei Kommandos werden zur Definition der Eingabedateien
verwendet:

INPUT-ATTRIBUTES spezifiziert die Satzlaenge, den Satztyp
und andere fuer alle Eingabedateien ge-
meinsanen Attribute;
musa stets angegeben erden

SORT-FILES Namen der zu sortierenden Dateien, wahl-
weise Angabe einer Anfangs- und Endsatz-
nummer bei Sortierung einer Teildatei

MERGE-FILES Namen der zu mischenden Dateien
Es muss entweder ein SORT-FILZS- oder ein MERGE-FILES-Komman-

do angegeben werden. Beide koennen benutzt werden, wenn im
gleichen Lauf sortiert und gemischt werden soll.

1.1. INPUT-ATTRIBUTBS

INPUT-ATTRIBUTES = satzlaenge {PIXED-LENGTH}
{CR-DELIIITED}
{VARIABLE }
{RELATIVE  }
{NO-SINGLE-Z }
{n0-222 !
{PP222

satzlaenge:
ist eine Zahl zwischen 1 und 4096, die die Laenge von Saetzen
mit fester Laenge in Bytes angibt oder die maximale Laenge von
Saetzen des Typs VARIABLE oder CR-DELINITED.

Eines der folgenden vier Attribute kann zur Spezifizierung des
Satztyps fuer alle Bingabedateien angegeben werden. Wird kei-
nes angegeben, wird als Standard FIXED-LENGTH angenommen.

FIXED-LENGTH die Bingabedatei(len) bestehen aus
Saetzen fester Laenge
CR-DELIMITED die Bingabedatei(en) enthalten Saetze,

normalerweise von variabler Laenge, die
durch carriage retum (ODH) und/oder
line feed (OAH) getrennt sind

VARIABLE die Eingabedatei(en) enthalten Saetze
variabler Laenge mit der Angabe der
Laenge in den ersten zwei Bytes jedes
Satzes

ze Aue

SORT 1520 (SCPX)

Beachte:

Die Satzlaenge muss wm 2 groesser ange-
geben werden als die maximale Anzahl der
eigentlichen Datenzeichen des Satzes.

RELATIVE die Bingabedatei(en) bestehen aus Saet-
zen des Typs RELATIVE

Beispiele:

INPUT-ATTRIBUTES
INPUT-ATTRIBUTBS
INPUT-ATTRIBUTBS
INPUT-ATTRIBUTBS
INPUT = '256, FIXED
1 256, FI (minimale Kurzform)

28 (feste Laenge angenommen)
80 ,RELATIVB

32, CR-DELINITED
256 , FIXED-LENGTH

Die letzten drei Beispiele sind aequivalent.

Die folgenden drei INPUT-ATTRIBUTES definieren die Datei-Ende-
Bedingung in FIXED-LENGTH-Dateien, die Binserdaten enthalten.
Fuer andere Datei- bzw. Datentypen sind diese Attribute ohne
Bedeutung.

NO-SINGLE-Z die PIXED-LENGTH-Datei ist durch ein 1AH
(CTRL Z) beendet

No-222 die PIXED-LENGTH-Datei wird durch einen
“ Satz von Zeichen 1AH beendet (schlieost
das Attribut N0-SINGLE-Z ein)

PFZZZ die Datei wird durch einen Satz abge-
echlossen, der mit zwei Zeichen FFH be-
einnt und mit 1AH aufgefuellt ist

Beispie

INPUT-ATTRIBUTBS = 80 ,FIXED-LENGTH,NO-SINGLE-2
INPUT = 128, FIX,N022,PPZ2Z
I 128 FI NOZ FF (minimale Kurzform)

2.

. SORT-FILBS

Das SORT-FILES-Kommando spezifiziert die Namen der zu sortie-
renden Dateien und wahlweise eine Anfangs- und End-Satznummer
je Datei.
SORT-FILES (eydatsinana[ (onfängasäbs ‚endsatz)],...

+

- . noch kein SORT-FILES spezifiziert bzw.
eingegebenes SORT-FILES-Kommando soll
ersetzt werden

-35-

SORT 1520 (SCPX)

* die angegebene(n) Datei(en) ergaenzen
die schon vorher spezifisierte(n):
bis 32 Bingabedateien koennen in

einen Lauf sortiert werden

dateiname SCP-Dateiname mit wahlweiser Laufwerks-
und Typangabe

anfangssatz eine Zahl von 1 bis 65535 als Nummer des
ersten Satzes einer Datei, der in die
Sortierung einbezogen werden soll;
0 oder 1 bedeuten Dateibeginn

endsatz Nummer des letzten zu sortierenden
Satzes einer Datei; 0 oder 65535 bedeu-
ten bis Dateiende

Beispiele:

SORT-PILE =EIN.DAT

SORT-FILE =B: BIN. DAT(10, 120)
SORT-FILES=E1.DAT,E2.DAT, B3.TXT

SORT = B:E1.DAT(1,119) ,C: B2.DAT(120,65535)

SORT-FILES + B4. DAT, E5.DS1(50,255)

S BIN.DAT (mininale Kurzform)

3.1.3. MBRGB-FILBS

Das MNERGE-FILES-Kommando spezifiziert die Namen der zu mi-
schenden Dateien.

HERGB-PILESC"Adateinans,dateinans, ...
+

dateiname SCP-Dateiname
- ersetzt vorher spezifizierte Nisch-Da-
' teinamen durch in diesem Kommando ange-
gebene
+ die in diesem Kommando spezifizierten

Dateien ergaenzen die vorher angegebe-
nen; maximal 32 zu mischende Dateien
sind moeglich 2

Beispiele:
MERGE-PILB = STAHM.DAT

M STAMM-DAT (minimale Kurzform)
MERGB + A:STAMM1.DAT,B:STAHM2.DAT

Wird ein SORT-FILES- oder ein MERGE-FILES-Kommando eingegeben,
prueft SORT 1520 (SCPX), ob die angegebenen Dateien vorhanden

- 36 -

SORT 1520 (SCPX)

sind und gibt gegebenenfalls eine Fehleranzeige. Das betref-
fende Kommando kann dann korrigiert werden.

2.3.2._Ausgabedatel

Im DUTPUT-FILE-Kommando wird der Name der Ausgabedatei ange-
geben, wahlweise ergaenzt durch Spezifikationen.

OUTPUT-FILE = dateinameL/CIL,satzlaenge) {FIXED-LENGTH)
{CR-DELINITED)
{VARIABLE
{RELATIVE
{K-OUTPUT
{R-OUTPUT
{P-OUTPUT
{KR-OUTPUT
{KP-OUTPUT
(FFZZZ
{NO-FFZZZ

datelname Name der Ausgabedatel; muss sich von den
Namen aller angegebenen Misch-Eingabe-
dateien unterscheiden, kann aber gleich
den Namen einer Sortier-Eingabedatei
sein (letzteres nur In Ausnahmefaellen
anuenden, da Eingabedatel ueberschr ieben
wird!)

„ Diskette Im angegebenen Laufwerk fuer
die Ausgabedatel soll vor dam Schreiben
der Ausgabedatei gewechselt werden.
Durch SORT 1520 (SCPX) erfolgt aine
Aufforderung zum Wechseln der Diskette,
die nach erfolgtem Austausch mit (ENTER)
zu beantworten ist.

Soll die Ausgabedatei vom Typ und der Spezifikation der Einga-
gabedatelen INPUT-FILE-Kommando) nicht abweichen, koennen
weitere OUTPUT-FILE-Spezifikationen weggelassen werden.

satzlaenge 1 - 4096, spezifiziert die Satzlaenge in
Bytes bzu. die maximale Satzlaenge fuer
CR-DELIMITED- oder VARIABLE-Saetze.
Angabe wird ignoriert, wenn eine
-OUTPUT-Option (K-QUTPUT usw.) spezifi-
ziert ist, da diese Formen eine ganz
bestimmte Laenge besitzen (s. Anhang A).

FIXED-LENGTH Bedeutung wıe entsprechende INPUT-Attri-
CR-DELIMITED bute

VARIABLE

RELATIVE

Die folgenden Bemerkungen sind nur wichtig, wenn der Typ der
Ausgabedatei und/oder die Satzlaenge unterschiedlich zum Typ
bzu. der Satzloenge der Eingabedatei angegeben wird, d.h. wenn

- 37 -

SORT 1520 (SCPX)
SORT 1520 (SCPX) eine Dateikonvertierung durchfuehren soll:

1. Weicht der Typ der Ausgabedatei von dem der Bingabeda-
tei(en) ab und wird die (maximale) Satzlaenge nicht angege-
ben, erfolgt standardmaessig:

a. ist die Eingabedatei VARIABLE, aber nicht die Ausgabeda-
tei, so ist die Satzlaenge der Ausgabedatei um 2 Byte
kleiner als die der Eingabedatei, da die zwei Laengen-
bytes geloescht werden

h. iet die Ausgabedatei VARIABLE, aber nicht die Eingabeda-
tei, so ist die Satzlaenge der Ausgabedatei ım 2 Byte
groesser als die der Eingabedatei, aber nicht groesser
als 4096.

2. Ist die angegebene Satzlaenge der Ausgabedatei kleiner als
die Satzlaenge der Eingabedatei, wird der Satz nach rechts
abgeschritten; die Laengenbytes bei VARIABLE werden ent-
sprechend veraendert.

vo

Ist die Ausgabedatei PIXED oder RELATIVE, ao werden gegebe-
nenfalls die Saetze mit Leerzeichen (20H) bis zur spezifi-
zierten Laenge aufgefuellt.

Die folgenden OUTPUT-Options spezifizieren Ausgabedateien:

K-OUTPUT Ausgabedatei enthaelt nur die Schluessel
als Teil der Snetze der Eingabeda-
teilen), die Ausgabe-Satzlaenge ent-
spricht der Schluessellaenge

R-OUTPUT Ausgabedatei enthaelt nur die sortierten
Satznummern; ist die Ausgabedatei vom
Typ CR-DELIMITED, so werden die Satz-
nummern als ASCII-Text mit je 6 Ziffern
angegeben, andere Dateitypen erzeugen
3-Byte-Binaernunnern

P-OUTPUT Ausgabedatei enthaelt nur Zeiger zu den
Bingabesaetzen; jeder Zeiger besteht aus
der Nummer des Sektors, in dem der Satz
beginnt und der relativen Position
(Byte-Offset) des Satzbeginns zum Beginn
des Sektors; die Zahlenangaben sind ent-
weder ASCII-Zeichen, durch ein Komma ge-
trennt (bei CR-DELINITED-Dateien) oder
binaer in zwei Bytes fuer jede Zahl
(andere Datei-Typen.

KR-OUTPUT Ausgabedatei enthaelt Schluessel und
Satznummern

KP-OUTPUT Ausgabedatei enthaelt Zeiger und Satz-
nummmern

-38 -

SORT 1520 (SCPX)

RECORD-NUMBERS aequivalent zu R-OUTPUT, minimale Kurz-
form ist REC

KEYS-AND-NUMBERS aequivalent zu KR-OUTPUT, minimale Kurz-
form ist KEY-A

Wenn eine der obigen OUTPUT-Options verwendet wird, ausser X-
OUTPUT, ist nur eine Eingabedatei erlaubt.

Die folgenden Attribute sind nur fuer Dateien vom Typ FIXED-
IBNGTH mit binaeren Daten signifikant:

FFZ2Z beendet die Ausgabedatei durch Auffuel-
len des freien Bereiches im letzten Sek-
tor mit 2 Bytes PFH und nachfolgenden
Zeichen 1AH (CTRL 2)

NO-FFZZZ beendet die Ausgabedatei durch Auffuel-
len des letzten Sektors der Datei mit
1AH; Standard, wenn nicht FPZZZ fucr (ic
Eingabe spezifiziert ist

Beispiele:

OUTPUT-FILE = IN.SUM , PIXED-LENGTH,KP-OUTPUT
OUTPUT = AUS.STM/C,2000,FIX,NO-FFZ

Sohlu:

Das KEY-Kommando spezifiziert einen oder mehrere Sortier-
schluessel. Es koennen bis zu 32 Schluessel angegeben werden,
jeder mit eigenen Attributen. Mehrfachschluessel koennen in
einem einfachen Kommando (wenn noetig, durch "&" Fortsetzung
auf der naechsten Zeile) oder in mehreren Kommandos (mit
KEY+) spezifiziert werden.

KEY (=) (hes,end) {ASCENDING }
(+) (#feld-nr,max-1) {DESCENDING }
(NWIERIO-ASCIT }

(UPPER-CASE }

{RIGHT-JUSTIPY 1

{LOHI
{MASK-PARITY-BIT}

{ALTSBQ }
{TWOS-CONPLEMENT}
{COMPUTATIONAL }
{INTEGER }
jELORTTNG. Born? }

PACKED-BCD }

{COMPUTATIONAL-3}

Brsetzen vorher spezifizierter Schlues-
sel

=39.:=

SORT 1520 (SCPX)

+ Hinzufuegen zu vorher spezifizierten
Schluesseln

Fuer ein Positionsfeld werden zuei Zahlen angegeben:

beg erste Position des Feldes

end letzte (enthaltene) Position

Das erste Zeichen eines Satzes hat die Position 1 (bei Saetzen
vom Typ VARIABLE befindet sich das erste Datenzeichen auf Po-
sition D.

Bei einer Satzstruktur, bei der die einzelnen Felder durch

Komma getrennt sind, werden zur Bezeichnung eines Feldes
(Kommafeld) das Zeichen "#” und zwei Zahlen angegeben:

feld-nr Feldnummer, z.B. Feld #3 befindet sich
im Satz zuischen dam zweiten und dritten
Komma .
max-ı maximale Anzahl der vor dem Feld zu be-

ruecksichtigenden Zeichen, wobel Anfueh-
rungszeichen, fuehrende und nach rechts
auffuellende Leerzeichen nicht gezahlt
werden;

veber max-] hinausgehende Zeichen werden
ignoriert; besitzt das Feld weniger Zai-
chen, so erfolgt ein Auffuellen mit
Leerzeichen

Die folgenden Angaben spezifizieren die Sortierreihanfolge,
ASCENDING ist Standard:

ASCENDING Feld aufsteigend sortieren
DESCENDING absteigend sortieren
Beispiele:

KEY = 2,10 ASCENDING (Bytes 2 bis 10)

KEY = 62,10 ASCENDING (2. Feld, maximal 10 Zeichen)
KEY = 62,10 DESCENDING (wie oben, aber absteigend)
K #2 10 DE (Minimalform)

Mehrere Schluessel koennen in einem Kommando spezifiziert
werden:

KEY = 2,10,ASC, 20,H5,DESC
KEY = #11,40,ASC, #3,5,DESC, #6,8,ASC
KEY = 1,10 ASC, #4,20 DESC

-%00-

SORT 1520 (SCPX)
Im KEY-Kommando angebbare Testattribute sind:

NUMERIC-ASCII interpretiert das Pelä als Text (im
Gegensatz zu binzer), bestehend aus nu-
merischen Zeichen

UPPER-CASE interpretiert im Feld enthaltene Klein-
tuchstaben als Grossbuchstaben

RIGHT-JUSTIPY fuellt den Feldanfang mit Leerzeichen
bei Kommandofeldern, deren Laenge klei-
ner ist als die spezifizierte maximale
Laenge (nax-1)

LOHI testet dieses Feld von rechts nach
links, vorgesehen fuer Binzerdaten, bei
denen der LOW-Teil zuerst gespeichert
ist, kann aber bei jedem Positionsfeld
verwendet werden; bei Kommafeldern nicht
erlaubt

HASK-PARITY-BIT hoechstwertiges Bit eines jeden Bytes
wird ignoriert

ALTSEQ sortiert das Peld entsprechend der im
COLLATIIIG-SBQUENCB-Kommando spezifizier-
ten Rangfolge (s.u.)

weitere Beispiele:
KEY = #17, 10,NUMERIC-ASCII,DESCENDING

KRY + #5,10 DESC,NW!-ASC, #8,20,ASCEID,RIGKT,UPPER
KEY = 1,10 LOHI,DESCEND, 11,20 ALTSEQ,ASCEND

Die folgenden Testattribute bezichen sich auf nicht-ASCII-
Datentypen:

TWOS-COMPLENENT jedes dicser aequivalenten Attribute
COMPUTATIONAL bedeutet, se die Daten binaere Fest-
kommazahlen sind, tobei negative Zahlen

als Zweierkonj;lenent gecpeichert sind.
Ds wird angenommen, dass der hoeherwer-
tige Teil der Zehl zmuerot gespeichert
ist (2-Bytc-Pe

rn).

INTEGER binnere Pestiommszahl; _niederwertiger
Teil zuerst gespeichert (2-Byte-Felder)

FLOATING-POINT Daten im binaeren Cleit'onma-Pormat;
(4-Byte-Felder:einfache Senauigkeit oder
F-Byte-Felder:dopielte Tenauiskeit)

PACKED-BCD jedes diesor aequivalenten Attribute

COMPUTATIONAL-3 bezeichnet geyacktc binzereodierte Dezi-
malzahlen (ECD) mit oder ohne Vorzei-
chen;

-4M-

SORT 1520 (SCPX)

wm PFehlinterpretationen zu vermeiden,
sollte COMPUTATZONAL-3 nur mit COMP-3
oder C3 abgekuerzt werden

Fuer vorzeichenlose binaere Festkommazahlen (hoechstwertiges
Bit repraesentiert Ziffernstelle), die in der Reihenfolge
I-Teil, H-Teil gespeichert sind, sollte LOHI benutzt werden,
bei Speicherung H-Teil, L-Teil wird kein Attribut benoetist.

weitere Beispiele:

KEY = 1,1 TWOS-COMPLEMENT, 3,4, TWOS-CONPLEIIENT

KEY = 1,2 INTEGER,DESC, 3,6 FLOATING-POINT, 11,14 FLOAT

KEY + 11,12 TWOS-COMPLEMENT,5,10 PACKED-BCD

KEY + 11,12 COMPUTATIONAL,S,10 COMP-3 (aequivalent vorherge-
hender Zeile)

2.3 Start

Sind alle gewuenschten Kommandos eingegeben, wird die Aus-
fuehrung durch die Bingabe des Kommandos

so
gestartet.

Das GO-Kommando (oder einfach G) bringt bei unvollstaendiger
Kommandoeingabe (z.B. keine INPUT-ATTRIBUTES spezifiziert)
eine entsprechende Fehlerausschrift (siehe Anhang B). Eine
Fehlerausschrift erfolgt auch bei sich widersprechenden Spe-
zifikationen (z.B. Schluesselposition groesser als definierte
Eingabenatzlaenge).

Nach einer Fehlerausschrift eind die vorher eingegebenen
Kommandos noch gueltig, so dass man nur die betreifenden
fehlerhaften Kommandos zu aendern hat und mit 60 erneut star-
ten kann.

Nachfolgend einige Beispiele fuer komplette Kommandofolgen:

INPUT-ATTRIBUTES = 62,CR-DELINITED
SORT-FILE = BIN.DAT
OUTPUT-FILE = AUS.DA
KEY = #2,4 ASCENDING  roHT-JusTLrr

INPUT = &0,FIXED
SORT LIST.NBU (1,1000)

MERGE LIST.STA

KEY = 64,70 ASC,31-40 ASC,1,19 ASC UPPER
OUTPUT = LIST.STA/C

co

-2-

SORT 1520 (SCPX)

INPUT = 1024, CR-DEL

MERGE LAG1. ART, LAG2.ART,LAG3. ART
KEY = #3,10,NUMERIC-ASCII,DESC
OUTPUT = LAGER.ART, 1000, FIXED

co

I = 64,CR-D;S=EIN.DAT;0=AUS;K=#2,8,DE;G

2.4. Kommandodatei

Einige oder alle Kommandos koennen zu einer Kommandodatei
zusammengefasst werden, dadurch vereinfacht sich die Komman-
doeingabe bei wiederholt durchzufuehrenden gleichartigen
Sortier- und/oder Hischvorgaengen wesentlich.

Die Kommandos in dieser Datei haben genau das gleiche Aussehen
wie bei einer einzelnen Eingabe an der Konsole. Die Datei kann
mit einem Texteditor (z.B. TP 1520 (SCPX)) angelegt werden.

Die Kommandodatei wird mit dem Kommando
CFILB=dateiname

aufgerufen.

Beispiele:

GFILE = SORT.TXT
CF B:UNSORT.STA

Eine Kommandodatei mss nicht die komplette Kommandofolge
enthalten.

Kommandos koennen vor dem Aufruf der Kommandodatei eingegeben
werden und „ach dem CFILE-Kommando, wenn die Kommandodatei ein
GO enthaelt.

Zur Anzeige der in der Kommandodatei enthaltenen Kommandos
dient das Kommando

LIST

‚ das entweder Bestandteil der Kommandodatei sein kann oder
mittels Tastatur vor Aufruf der Kommandodatei eingegeben wird.
Die Anzeige beginnt mit der nach LIST folgenden Kommandozeile.

Fuer alle Sortierungen, ausser bei sehr kleinem Datenbestand,
benutzt SORT 1520 (SCPX) eine Arbeitsdatei auf Diskette.
Standardmaessig befindet sich diese Datei auf der im aktuellen
Laufwerk vorhandenen Diskette.

= >

SORT 1520 (SCPX)
Diese Standardzuordnung kann durch das Kommando:
WORK-DRIVE = laufwerk

veraendert werden. Dabei ist "laufwerk" ein Buchntabe A bis E
entsprechend der Systemkonfiguration, wahlweise gefolgt von
einem Doppelpunkt ":".

Beispiele:

or
\ORK E

RIVT =

Bei umfangreichen zu sortierenden Datenmengen sollte ein ent-
sprechend freier Bereich auf einer Diskette reserviert (etwa
in Groesse der Ausgabedatei) und das WORK-DRIVE-Kommando be-
nutzt werden.

Die Arbeitsdatei erhaelt den Namen SORT.oou und wird nach
erfolgter Sortierung automatisch geloescht.

2.6. TAGSORT

Das Kommando
TAGSORT

veranlasst SORT 1520 (SCPX), einen speziellen Algorithmus zum
Erzeugen der Ausgabedatei anzuwenden.

Dieses Verfahren bewirkt eine Reduzierung bzw. einen Wegfall
des fuer die Arbeitsdatei benoetigten Diskettenbereiches (2-
Laufwerksanlagen).

TAGSORT darf nur verwendet werden, wenn nur eine zu sortie-
rende Eingabedatei und keine zu mischenden Eingabedateien vor-
liegen.

Ergebnisausschrift

SORT 1520 (SCPX) besitzt sechs Varianten der Ausschrift nach
Abschluss des Sortier- bzw. Nischvorganges. Die Auswahl der
gewuenschten Variante erfolgt durch das Kommando

PRINT-LBVEL = n
wobei n eine Zahl von 0 bis 5 repraesentiert (2 ist Standard).
Beispiele:

PRINT-LEVEL = 4
PR 1

- Variante 0: keine Anzeige

- bb -

- Variante 1:

- Variante 2:

- Variante 3:

- Variante 4:

- Variante 5:

SORT 1520 (SCPX)

„Anzahl der surtierten Saetze, wenn £ 0
‚Anzahl der gemischten Saetze, wenn # 0°
‚Anzahl der Ausgabesaetze, wenn Angabe
nicht redundant

.Groesse der Arbeitsdatei in KByte,

wenn £ 0

(Standard)
wie Variante 1, plu
Anzahl der SORT-Eingabesaetze, wenn

nicht redundant
‚Anzahl der MERGE-Eingabesaetze, wenn
nicht redundant
“Anzahl der bei Selektinn nicht
behandelten Saetze:

der SORT-Bingabesaetze, wenn / 0

der !IERGE-Bingabesaetze, wenn / 0

wie Variante 2, plus:
‚Anzahl der Sortierlaeufe
‚Anzahl der Mischlaeufe

wie Variante 3, plus Text:
SORTING.

MERGING.

ADDITIONAL MERGE...

FINAL MERGE...

wie Variante 4, plus:

„Anzahl der Bytes des Arbeitsspeichers
‚Anzahl der Eingabesaetze bei jedem Sor-
tierlauf

Anzahl der sortierten Eingabebloecke
bei jedem Mischlauf

2.8. Weitere Kommandos

CANCEL

BYE

NO-ERROR-MESSAGES

loescht alle vorher eingegebenen Kom-
mandos (ausser COLLATING-SEQUENCE-Kom-
mandos)

verlaesst SORT, aegivalent zu CTRL C
(WHarastart)

tritt ein Fehler waehrend der Abarbei-
tung auf, so wird nur die Pehlernummer,
nicht der gesamte Text angezeigt; der
fuer den Fehlertext benoetigte Speicher-
bereich ist zusaetzlich als Arbeitsspei
cher verfuegbar;

Kurzform: NO-E

-45 -

SORT 1520 (SCPX)

ERROR-MESSAGE zahl zeigt fuer den Fehler nit der Nummer
"zahl" den gesamten Fehlertext an; An-
gabe nur nach einem NO-BRROR-MESSAGB-

B Kommando sinnvoll

RETURN-TO-CONSOLE kehrt nach Abarbeitung nicht zum Be-
triebssystem zurueck, sondern ermoeg-
licht die Pingabe von Kommandos fuer
einen naechsten Sortier-/Nischlauf;
Kommando reduziert den verfuegbaren Ar
beitsspeicherbereich; ist nicht zu em-
pfehlen bei grossen Datenmengen oder
maximaler Verarbeitungsgeschwindigkeit

2.9. Selektion

Es besteht die Moeglichkeit, nur mit ausgewaehlten Saetzen
einen Sortier- bzw. Mischvorgang durchzufuehren.

2.9.1. Rinfuehrung

SELECT und EXCLUDE definieren Äriterien fuer die Selektion von
Snetzen auf der Basis von Datentests. Bei SELECT werden alle
die Testkriterien erfuellenden Saetze in den Sortier-/Misch-
prozess einbezogen. Bei EXCLUDE dagegen werden in diesen
Prozess alle die Saetze einbezogen, bei denen die Testkri-
terien nicht erfuellt sind.

SELECT und BXCLUDE koennen bei Sortier- und Misohvorgaengen
verwendet werden.

Die SELECT- und EXCLUDE-Kommandos haben eine gemeinsame Syntax
mit vielen Optionen und Varianten. Die gebraeuchlichste Form
ist:

(SELECT) (=) (FIELD beg,end) operator"text" (NUMERIC-ASCII }
(EXCLUDE)C(+) ([LFIELD]#feld-nr) IURZEn-CHER
RIGHT-JUSTIFY

{LoHI
{HASK-PARITY-BIT}
{ALTSEQ }

- ersetzen ein vorher eingegebenes SELEOT/
EXCLUDE-Kommando bzw. Ersteingabe

+ ergaenzen vorher eingegebene gueltige
SELECT/ BXCLUDE-Kommandos
FIELD bes,end beg und end bezeichnen ein Positionsfeld

durch Angabe der Nummer des ersten und
letzten Bytes, wie beim KEY-Kommando

-46 -

SORT 1520 (SCPX)

FIELD #feld-nr bezeichnet ein variables Komnafeld, wie
beim KEY-Kommando, ausser dass hier
keine aximallaenge angegeben wird; das
vollstaendige Feld wird in den Test ein-

bezogen
operator einer der nachfolgenden Vergleichsopera-
toren:
< oder LT kleiner
<= IE kleiner oder gleich
= EQ_ gleich
> NE ungleich
>= GE groesser oder gleich
> GT groesser
"text" ist der Text, mit dem das spezifizierte
Peld verglichen wird, eingeschlossen in
Anfuehrungsstrichen
NUNERIC-ASCIT
UPPER-CASE
RIGHT-JUSTIFY analoge Bedeutung wie beim KEY-Kommand
H beziehen sich auf beide, Feld und "text"

LOHI
MASK-PARITY-BIT
ALTSEQ

Beispiele:

SELECT = FIELD #2>="84101.4"

SELECT = FIELD #3="UEBWS"

SE #3 = "UBBIS" (aequivalent vorhergehender Zeile)
SELECT = FIELD #7>"300>NWIERIC-ASCII

SELECT + #6LE"400"NWM-ASC

Verden MNehrfach-SELECT- und/oder EXCLUDE-Kommandos angegeben
(a.0., letzte beide Zeilen), so werden all jene Saetze in das
Sortieren/lischen einbezogen, die bei SELECT alle Be-dingungen
erfuellen bzw. bei EXCLUDE nicht erfuellen.

2. Bereich

t

Der Test, ob der Wert eines Feldes in einem besfimnten Bereich
liegt, kann auch mit den Operatoren "between" und "not-
between" erfolgen.

(SELECT)(=) Feldangabe wie oben (BT) "text1","text2" Attribute

wie oben
(EXCLUDE)C+) (nB)
Br zwischen "text1" und "text2" (ein-
schliesslich)
NB nicht zwischen "text1" und "text2"
"texti" untere Bereichsgrenze

==

SORT 1520 (SCPX)
"text2" obere Bereichsgrenze
Beispiele:

SELECT #2 NB"-200","200" NUMBRIC-ASCII
SELECT = PI 33,41 BT "100","999"

2.9.3. Weitere Selektionsmoeglichkeiten

- Vertauschen von Feldern und Konstanten:
Angabe der Konstanten vor und des Feldes nach dem Operator;
Vergleich eines Feldes mit einem oder mehreren Feldern im
gleichen Satz:

SELECT = "UEBWS" = FIELD #3
SELECT = PIELD #1 > PIELD #2
SELECT = "128" BT PIELD #1, FIELD #2

Der letzte Test ist erfuellt, wenn Feld 1 kleiner oder
gleich 128 und Feld 2 groesser oder gleich 128 ist.

- Logische Operatoren:
komplexe Bedingungen koennen durch Verwenden der logischen
Operatoren AND, OR und XOR angegeben werden:

SELECT = #2 > "200"NnUM OR #1 = "SUN"
SELECT = #2 GT "200"NUM AND NOT #1 = "SUR"

AND, OR, XOR und.NOT koennen nicht abgekuerzt werden.

- Klammerung:
Klammern koennen zur Zusammenfassung logischer Operationen
verwendet werden; ohne Klammerung erfolgt die Verarbeitung
in der Reihenfolge: alle Vergleiche zueret, dann NOT, AND,
XOR und OR:

SELECT = #3 = "A" AND (#4 >="615" OR #1 = "XY")
2.9.4. Selektion bei nioht-ASCII-Daten
Eine Selektion von Saetzen mit BCD-Daten und verschiedenen

Typen von Binaerdaten ist durch die Angabe spezieller Testat-
+ribute und Konstantendarstellungen ebenfalls moeglich.

2.9.4.1. Spezielle Testattribute

Die folgenden Testattribute koennen ebenfalls zur Selektion
von Saetzen benutzt werden. Ihre Bedeutung ist wie die der
entsprechenden KEY-Attribute.

TUOS=CONPLEIEIIT oder COMNPUTATIONAL
INTEGER

FLOATING POINT

PACKED-BCD oder COMPUTATIONAL-3

-18-

SORT 1520 (SCPX)

Zuerst einige Beispiele zur Anwendung obiger Testattribute
beim Vergleich von Feldern:

SELECT = FIELD 1,2 > FIELD 3,4 INTEGER

SELECT = FIELD 1,2 < FIELD 3,4 LOHI

SELECT = FIELD 1,8 BT FIELD 9,16, FIELD 17,24 FLOATIIC
SELECT = FIELD 1,10 < 21,30 COMPUTATIONAL-3

2.9 Numerische Konstanten

Die gegeneinander zu testenden lierte koennen statt eines in
Anfuehrungszeichen eingeschlossenen Textes auch numerische
Konstanten sein. Die numerischen Konstanten repraesentieren
die im Buerocomputer gespeicherten binaeren \erte, ausge-
drueckt als Oktal-, Dezinal-, Hexadezimal- oder BCD-Zehlen. Es
kann die Reihenfolge der internen Speicherung angegeben vwer-
den, ob der hoeher- oder niederwertige Teil zuerst gespeichert
ist.

Numerische Konstanten werden versendet, wenn Binaer- oder ECD-
Daten getestet werden, «mehrend der "ext bei ASCII-Daten ge-
nutzt wird. (Nwmerische Konstanten koennen auch bei ASCII-
Daten verwendet werden, dazu muss die ASCII-Codierung eines
jeden Zeichens eingegeben werden.)

SORT 1520 (SCPX) erlaubt die explizite Angabe der Konstan-
tenlaenge in Bytes. Enthaelt die Konstante weniger
signifikante Ziffern, so werden die fehlenden hoeherwertigen
Bytes durch SORT hinzugefuegt, bei negativen Zahlen (ausser
bei BOD-Konstanten) Bits mit dem Vert 1, in allen anderen
Faellen Bite mit dem Wert 0

Bine numerische Konstante besteht aus einen optionalen
Vorzeiohen, einer Anzahl von Ziffern, die in einer Zeile
eingegeben werden koennen. Optional ist die Angabe der Zahlen-
daratellung und, wenn eine Angabe der Zeichendarstellung er-
folgt, optional eine Dezimalzahl, die die Laenge der Konstan-
ten in Bytes darstellt. Vor dem Buchstaben, der die Zahlendar-
stellung spezifiziert bzw. zwischen diesem und einer nachfol-
genden Laengenangabe, darf kein Leerzeichen stehen.

Die Zahlendarstellung wird wie folgt definiert:
T dezimal, Zahl gewandelt in Binaerzahl

I dezimal, gewandelt in Binaerzahl, niederwertiger Teil zuerst
gespeichert

P dezimal, gewandelt in BCD-Zahl
H hexadezimal

Q oktal

-19 -

SORT 1520 (SCPX)

Wird keine Zehlendarstellung spezifiziert, wird dezimal (T)
angenommen, ausser, wenn die Konstante mindestens ein Zeichen
A- P enthaelt, dann hexadezinal.

Alle numerischen Konstanten werden von werthoechster (an
adresseniederer Stelle) zu wertniedrigster Stelle im Speicher
abgelegt, ausser denen, die als I-Konstanten definiert wurden,
wo das niederwertige Byte des Feldes ebenfalls zuerst gespei-
chert ist, d.h. wo das LOHI- oder INTEGER-Testattribut verwen-
det wird. Wird eine hexadezimale Konstante bei einem solchen
Test verwendet, dann muss sie so angegeben werden, dass das
niederwertige Byte zuerst geschrieben wird.

Beispiele:

Konstante | Byteanzahl | Speicherdarstellung (hexadezimal)
! |
! j

12 ! 1 ! oc

12T | 1 | oc

1272 I 2 | 00 0C

1212 I 2 j oc 00 “

12H I 1 | 12

12P | 1 I 12

-2 ! 1 I FE

214 I 4 ! 2 PF FF FF

1234 P | 2 ! 12 3%

Konstante | Byteanzahl | Speicherdarstellung (hexadezimal)

|

MZIP ! 3 I 01 23 a0

-123up6 | 4 ! 00 01 23 4D

1029AH | 3 ! 01 02 9A

1029A ! 3 ı 01 02 9A (ala hexa angenommen)

0004H H 1 I 01 (fuehrende Nullen nicht

I

abgespeichert)

1

Einige Beispiele fuer die Verwendung numerischer Konstanten in,
SELECT-Kommandos:

SELECT = FIELD 17,17 = 4AH
Test auf 4A hexa in Position 17

SELECT = PIELD 1,8 < 1024PB CONPUTATIONAL-3
nimm: Saetze, bei denen in Positionen 1-&, interpretiert als
gepackte Dezimalzahl, ein ert kleiner ala 1024 enthalten ist.

2.9 3. Konstantenliste

In Ergaenzung zu Peldern und ‘einfachen Konstanten kann noch
eine andere Form von Werten angegeben werden: die Konstanten-
liste. Eine Konstantenliste ist eine Folge von Konstanten,
getrennt durch Komma oder Leerzeichen und eingeschlossen in
eckige Klammern. Die Konstanten brauchen nicht von gleichem

Par Ben

SORT 1520 (SCPX)

Typ zu sein. Bin Teil kann als Text und ein anderer als nume-
rische Konstante angegeben werden:

SELECT = FIELD 1,6 = ["suMit",ODOAH]

2.9.5. Veraenderung der Rangfolge

lit dem COLLATING-SEQUENCE-Kommando kann die Rangfolge der
Zeichen spezifiziert oder modifiziert werden. Diese veraen-
derte Rangfolge wird beim Feldtest im KEY-, SELECT- und
EXCLUDE-Kommando durch das Attribut ALTSEQ wirksam.

COLLATING-SEQUENCE (2 arsunent,argunent,...
+

(+ und = besitzen gleiche Wirkung).

Standardmaessig ist die ASCII-Zeichenfolge (8-Bit-Binaervert-
folge) als Rangfolge definiert, d.h. auf Position 0 befindet
sich das Zeichen mit dem Binaerwert OOH oder auf Position 49
die Ziffer "1" als Zeichen mit dem Bimmervert 31H bzw. auf
Position 65 das ASCII-Zeichen "A" mit dem Binaermert 41H.

Damit gilt beim Vergleich "1" < "a",

Bine Veraenderung dieser Rangfolge durch das COLLATING-
SBQUENCE-Kommando kann in zwei Formen erfolgen: implizit und
explizit.

- Implizite Positionsargumente

Die Position ergibt sich aus der Reihenfolge der angegobenen
Zeiohen, die allgemein in Anfuehrungszeichen eingeschlossen
oder als Code dezimal, hexadezimal oder oktal dargestellt
werden koennen.

Beispiele:
Co = "m, mazu, "0-9", 7FH

Das erste Zeichen, hier das Leerzeichen, wird stets der
Position 0 zugeordnet, das zweite Zeichen der Position 1
usw.

Im angefuehrten Beispiel ist das zweite Zeichen der Beginn
eines Bereiches von Zeichen, naemlich der Buchstaben "A bis
2". Bei der Angabe von Bereichen werden erstes und letztes
Zeichen durch einen Bindestrich getrennt. Nicht angefuehrte
Zeichen behalten ihre Standardposition.

- Explizite Positionsargumente

Eine Position wird durch eine numerische Konstante oder ein
in Anfuehrungszeichen eingeschlossenes Zeichen direkt ange-
geben, gefolgt von einem "+" fuer steigende, einem "-"
fuer fallende oder keinem Zeichen fuer gleiche Positionen.

- 51 -

SORT 1520 (SCPX)

Beispiele:
a) COL = "at - "zu = 554, "al — nZU = 654

d.h. "a" = "A" (Position 65)

"bp" = "B" (Position 56)

Klein- und Grossbuchsteben besitzen gleiche Positionen.
5) COL = "or - "gt = 50-

ist gleich

COL = non non = 404
e) COL = "on - "9

50
alle Ziffern haben die gleiche Position
- Gemischte Angaben

Folgt einen expliziten Positionsargwment ein implizites
Argument, so wird ihm die Position des letzten expliziten
Argumentes +1 zugeordnet, unabhaengig davon, ob die
vorhergehenden Argumente steigend, fallend oder von
konstanten Positionen sind.

Ein Argwent setzt ein links von, ihm stehendes Argument
ausser Kraft.

Beispiel:
coL 0 - 255 = 255, "." = 0, "0" -

alle Zeichen (0-255) erhalten explizit die Position 255,
implizit der Punkt die Position O und die Ziffern "0 bis 9"
die Positionen 1 - 10.

- liehrfach-Kommandos

Werden mehrere COLLATING-SEQUENCE-Kommandos angegeben, so
beginnt das zuletzt eingegebene Kommando implizit bei
Position 0, d.h. ein "+" nach dem Kommandonamen hat ergaen-
zende Wirkung wie bei anderen Kommandos (z.B. KEY+...).

- Reinitialisieren
Es erfolgt kein automatisches Ruecksetzen einer veraenderten
Rangfolge zur Standardrangfolge durch das CANCEL-Kommando.
Durch
coL 0 - 255

vird aber die Standardrangfolge wiederhergestellt.

SORT 1520 (SCPX)

INHANG_A:_Foı

Vebersicht

Dateiformate
1. _ CR-DELIMITED
2. _FIXED-LENGSTH
.3. VARIABLE

4. RELATIVE

= Feldformate

“1. Positionsfelder
“2. Kommafelder

3. Datenformate
3.1.  NUMERIC-ASCII
3.2. FACKED-BCD (COMPUTATIONAL-3)

4. Ausgabedateiformate

4.1. K-OUTPUT (nur Schluessel)

4.2. _R-DUTPUT (Satznummern)

4.3. P-OUTPUT (Zeiger)

4.4. _KR-OUTPUT (Schluessel und Satznummern)
4.5. _KP-OUTPUT (Schlussel und Zeiger)

4.__Datellormate

SORT 1520 (SCPX) kann Dateien behandeln, die Saetze eines der
vier Typen CR-DELIMITED, FIXED-LENGTH, VARIABLE oder RELATIVE
enthaelt.

As12_CR-DELINITED

Eine SCP-Datel besteht normalerweise aus Textzeilen (=Saetze),
die durch carriage return (ODH, Wagenrichtung und Zeilen-
schaltung auf Zellenanfang) und line feed (OAH, Zeilenschal-
tung) abgeschlossen werden. Endet der letzte Satz der Datei
nicht genau am Sektorende, folgt dem letzten Satz ein CTRL Z
“AH. Endet der letzte Satz genau am Sektorende, so kann
abhaengig vom dateierzeugenden Programm oder der verwendeten
Sprache ein zusaetzlicher Sektor mit einem CTRL Z die Datei
abschliessen.

SORT 1520 (SCPX) erzeugt dieses Format fuer die Ausgabedatel.
Fuer Eingabedateien sind zum oben beschriebenen Format Var ia-
tionen moeglich.

Wird eine CR-DELIMITED-Datei gelesen, so wird auch ein einzel-
nes Zeichen carriage return oder line feed als Satztrennzei-
chen akzeptiert. Zusaetzliche carriage return oder line feed
vor dem Beginn des naechsten Satzes werden ignoriert.

Eine Datei benoetigt also nur ein Trennzeichen zwischen den
Saetzen; Leerzeilen, die nicht mindestens ein Leerzeichen
enthalten, werden bei der Eingabe ignoriert.

55.4

SORT 1520 (SCPX)

Bin CTRL Z wird auch als Satztrennzeichen angesehen, so dass
der letzte Satz der Datei auch dann korrekt gelesen wird, wenn
das carriage retum nach ihm fehlt.

SORT 1520 (SCPX) erkennt das Dateiende, wenn das physische
Dateiende erreicht ist oder ein CTRL Z gelesen wird. Die
Eingabeattribute NO-SINGLE-Z und NO-2ZZZ wirken bei CR-DELIII-
TED-Eingabedateien nicht.

Wird eine CR-DELINITED-Ausgabedatei erzeugt, so werden ein
carriage return und ein line feed nach jedem Satz geschrieben,
ünabhaengig davon, wie der entsprechende Ringabesatz abge-
sehlonsen war. Die Datei wird abgeschlossen, indem der freie
Bereich des letzten Sektors mit CTRL Z's oder zwei Zeichen
FF, gefolgt von CTAL Z's wenn das FFZZZ-Ausgabeattxibut epe-
zifiziert wurde, aufgefuellt wird. Reichen die Daten bis zum
Sektorende, werden keine EOF-Zeichen ausgegeben.

OR-DELINITED-Dateien enthalten normalerweise nur ASCII-Text-
daten (20H - 7FH), bei anderen Daten mies sicher sein, dass
sich keine Codes 14H, ODH oder OAH innerhalb des Satzea
befinden.

1.2. FIXBD-LENGTH

PIXED-LENGTH-Dateien bestehen aus Saetzen fester Laenge und
koennen alle Zeichen ausser 1AH enthalten. Die Satzlaenge ist
unabhaengig von der Sektorlaenge. Sie kann 1 - 4096 Bytos
betragen.

Bei Bingabedateien wird das Dateiende erkannt, wenn bia zum
physischen Dateiende weniger Bytes ale die spezifizierte Satz-
laenge bleiben. Wurde entweder NO-SINGLE-Z oder N0-22Z apezi-
fiziert, dann wird das Ende der Datei duroh ein CTRL 2 (NAH)
als erstes Zeichen eines Satzes erkannt.

Bei Ausgabedateien wird der ungenutzte Bereich des letzten
Sektors der Datei entweder mit CTRL Z's (1AH) oder zwei FFH,
gefolgt von CTRL Z's aufgefuellt (Attribu. FFZ22). Endet der
letzte Satz am Sektorende, werden keine CTRL Z'sausgegeben,
FF222 wird standardmaessig fuer die Ausgabedatei angenommen,
wenn es fuer die Bingabedatei spezifiziert wurde.

1 VAR!

ILB

Dateien vom Satztyp VARIABLE enthalten Saetze variabler
Laenge, wobei die Laenge jedes Satzes binaer in den ersten
beiden Bytes des Satzes enthalten ist. Im ersten Laengenbyte
ist der H-Teil der Satzlaenge gespeichert.

Diese Satzlaenge kann im Bereich von 1 - 4094 liegen, wobei
die zwei Laengenbytes nicht einbezogen sind. Es ist zu beach-
ten, dass bei der Spezifikation der llaximalsatzlaenge der
Datei diese zwei Bytes einzubeziehen sind. Das erste Datenzei-
chen befindet sich immer auf Position 3.

54 -

SORT 1520 (SCPX)

Das Dateiende bei Eingebedateien wird erkannt, wenn bis zum
physischen Dateiende weniger als zwei Bytes verbleiben oder
wenn die Laengenbytes zwei CTRL Z's (14H) enthalten.

Bine Ausgabedatei wird durch zwei Bytes 00H beendet (Mull-
Laengensatz), eventuell verbleibender freier Bereich bis zum
Sektorende wird mit CTRL Z's aufgefuellt (bzw. 2FFH), gefolgt
von 1AH bei spezifiziertem FFZZZ).

VARIABLE-Dateieh koennen beliebige Datenzeichen ab Position 3
eines jeden Satzes enthalten.

1 RELATIVE

RELATIVB-Dateien sind ein Spezialtyp von Dateien fester
Satzlaenge. Die Saetze der Datei sind von 1 ab durchnumeriert
anzusehen (Satznummer nicht gespeichert). Geloeschte Saetze
bestehen aus Binaernullen (00H) in der spezifizierten
Satzlaenge von 1 - 4096 Bytes

In die Ausgabedatei werden geloesohte Saotze nioht ueber-
nommen. Jede RELATIVE-Datei besitzt am Dateianfang ein
"Header"-Feld aus 6 Bytes, das folgende Angaben enthaelt:

Byte Inhalt

0 als Kennzeichen fuer
RELATIVE-Datei

4,2 Satzlaenge binaer,
3 frei
4,5 hoechste Satznummer der

in der Datei enthaltenen

|
|
|
|
|
|
| I-Teil zuerst gespeichert
|
|
|
|
|
| Sastze

Der erste Satz beginnt ab Byte 6 des ersten Sektors, Die
Saetze koennen beliebige Daten enthalten.

2. Feldformate
Bei KEY, SELECT bzw. EXCLUDE erfolgt die Angabe eines oder

mehrerer Felder, die als Positionsfelder oder Kommafelder
definiert werden koennen.

SORT 1520 (SCPX)

Positionsfelder

Bin Positions! £ i Nureh eine Anfongs- und
Endposition. Die Positionen (Spalten) beginnen mit 1 beim
ersten Byte eines Satzes oder 3 beim ersten Datenbyte einer
VARTABLE-Datei. Kommas, Anfuehrungszeichen oder andere Zeichen
sind Bestandteil des Satzes und erfahren keine Sonderbehand-
lung.

Geht die Endeposition eines Feldes ueber das Ende eines CR-
DELINITED-Satzes hinaus, werden den restlichen Positionen
Leerzeichen (20H) zugeordnet. Befindet sich das gesamte Feld
ausserhalb des Satzes, erfolgt zusaetzlich eine
Fehlerausschrift.

2.2. Kommafelder

Bin Kommafeld wird durch seine Feldnummer spezifiziert.

Feld n ist dabei das Feld, das nach dem (n-1)-ten Komma
(ausser in Anfuehrungszeichen eingeschlossene Kommas) beginnt.
Ist n=1, beginnt das Feld am Satzbeginn (Position 3 bei
VARIABLE-Saetzen). Das naechste Komma (ausser in Anfuehrungo-
zeichen eingeschlossene) bzw. das Satzende kennzeichnen das
Ende des Peldes.

Leerzeichen (20H) am Beginn oder Ende eines Feldes werden
ignoriert, ebenso alle Anfuehrungszeichen im Feld.

Ein Schluesselfeld, das nach dem Eliminieren dieser Zeichen
noch groesser ist als die spezifizierte MNaximallaenge, wird
nach rechts abgeschnitten.

Ein Kommafeld kann bis zu 4096 Zeichen lang sein, die
Feldnummer muss im Bereich 1 - 255 liegen.

3. Datenförmate

Das Basisdatenformat von SORT 1520 (SCPX) ist ein Strom vor-
zeichenloser (unsigned) Bytes. Werden zwei Schluessel waehrend
des Sortierens oder Mischens verglichen oder ein Satzselek-
tionstest durchgefuehrt, werden die zwei Bytestroeme von links
nach rechts verglichen bis eine Position gefunden ist, an der
die Bytes ungleich sind. So werden ASCII-Daten und vorzeichen-
lose binaere Festkommazahlen (unsigned binary fixed point),
bei denen der H-Teil (hoeherwertige Teil) zuerst gespeichert
ist, sortiert. Dabei gibt es zwischen beiden Typen keinen
Unterschied. Nachfolgend werden weitere Datenformate erlaeu-
tert.

3 NUMERIC-ASCII

Bin NUMBRIC-ASCII-Datım ist eine Zahl, repraesentiert durch
eine Folge von ASCII-Zeichen. Eine in untenstehender Reihen-
folge angeordnete Zeichenfolge eines Feldes wird dabei als

56.

SORT 1520 (SCPX)
korrekt interpretiert:
- eine beliebige Anzahl von Leerzeichen

- wahlweise "+" oder "-", gefolgt von einer Anzahl Leerzeichen

fuehrende Nullen

Ziffern

\

wahlweise Dezimalpunkt

zusaetzliche Ziffern bei vorhandenem Dezimalpunkt

Leerzeichen

-E,  e, D oder d als Kennzeichen eines nachfolgenden
Brponenten

wenn E, e, D oder d:

. Leerzeichen

« yahlveise "+" oder "-", wihlweise gefolgt von Leerzeichen
bis zu 3 Ziffern zwischen -255 und +255

Maximal 14 signifikante Ziffern sind zugelassen.

Kann nur ein Teil eines als NUMERIC-ASCII deklarierten Feldes
numerisch interpretiert werden, so wird der so erhaltene Wert
benutzt. Ist das gesamte Feld nicht numerisch, wird der Wert 0
angenommen.

Ein BExponent ausserhalb des Bereiches von -255 bis +255 oder
mehr als 3 Ziffern fuer den Bxponenten ergeben einen undefi-
nierten Wert. Bine Fehlerausschrift erfolgt nicht.

3. PACKED-BCD (COMPUTATIONAL-3)

Das PACKED-BCD (Binary Coded Decimal)-Attribut definiert
gepackte Dezimalzahlen. Bei diesen Daten werden zwei Ziffern
in einem Byte gespeichert, die hoeherwertige Ziffer in den
hoechsten Bits; die hoehemirertigen Bytes zuerst (an adressnie-
derer Speicherstelle). Jede Ziffer wird in 4 Bits dargestellt,
das optionale Vorzeichen befindet sich in den niederen (rech-
ten) 4 Bits des letzten Bytes. Die \lerte in diesen 4 Bits
besitzen folgende Bedeutung:

0-9 Ziffer: Zahl ist ohne Vorzeichen
A, C, E oder ? hexa + Vorzeichen
B oder D hexa - Vorzeichen

SoRT 1520 (SCPX) vergleicht korrekt PACKED-BCD-Zahlen mit und
ohne Vorzeichen miteinander, auch Zahlen unterschiedlicher
Vorzeichendarstellung.

=:

SORT 1520 (SCPX)

4 Ausgabedateiformate

Es gibt verschiedene Ausgabemoeglichkeiten. In allen Faellen
beginnen die Saetze mit zwei Bytes Satzlaenge, wenn VANIABLE-
Ausgabedatei spezifiziert ist und enden mit carriage veturm
und line feed bei CR-DELIMITED-Ausgabedateien.

Wurde eine der Ausgabemoeglichkeiten spezifiziert, ist eine
angegebene Ausgabesatzlaenge ohne Bedeutung. Die Ausgabesatz-
laenge fuer die verschiedenen Optionen sind die folgenden:

Dateityp K-OUTPUT R-OUTPUT P-OUTPUT KR-OUTPUT KP-OUNPUT
CR-DELIMITED Keys 6 13 Keys+6 Keys+13
FIXED-LENGTH Keys 3 4 Keys+3  Koyatı
VARIABLE Keys+2 5 6  Keys+5  Keys+
RELATIVE Keys 3 4 Keys+3 Keys+4

"Keys" bezeichnet die Schluessellaenge, wie unten bei K-OUTPUT
beschrieben. Zur Satzlaenge bei CR-DELIIIITED-Dateien werden
earriage return und line feed nicht mitgezaehlt.

4.1

UTPUT (nur Schlu: 1

Jeder "Ausgabesatz" einer K-OUTPUT-Datei enthaelt nur die
Schluesselbegriffe. Die Satzlaenge ist die Gesantlzenge de
modifizierten Schluessel. TACKED-BCD-Felder besitzen ein Byte
mehr, alle NWMERIC-ASCII-Schluessel haben 9 Bytes.

Die Schluessel sind "verkettet", der Schluessel mit der
hoechsten Prioritaet zuerst. Bei Kommafeldern werden, wie oben
beschrieben, Leerzeichen und Ausfuehrungszeichen eliminiert;
sie werden bis zu ihrer llaximallaenge mit Leerzeichen aufge-
fuellt. Spezifizierte Attribute werden wie folgt verwendet:

NUMERIC-ASCLI Daten in 9-Byto-Pormat konvertiert
(siehe unten,

UPPER-CASE Kleinbuchstaben in Grossbuchataben unge-
wandelt

RIGHT-JUSTIPY fuehrende Leerzeichen werden bei Komma-

feldern bis zum Erreichen der Naximal-
laenge hinzugefuegt

LOHI umgekehrte Bytereihenfolge
MASK-PARITY-BIT in allen Bytes Bit 7 gleich 0 gesetzt
ALTSBQ Position jedes Bytes entsprechend

COLLATING-SEQUENCE-Kommando veraendert

TWO S-COMPLEMENT Zweierkomplement (Komplement Bit 7 des
ersten Bytes)

-58 -

SORT 1520 (SCPX)

INTBGER wngekehrte Bytereihenfolge und Komple-
ment Bit 7 des ersten Bytes

FLOATING-POINT siehe unten

PACKED-BCD ein Byte FFH wird dem Schluessel voran-

gestellt; 4st ein Vorzeichen vorhanden,
so wird dies durch eine Rechtsverschie-
bung wm 4 Bits geloescht; war das Vor-
zeichen negativ, so wird das Binerkon-
plement des gesamten Schluessels, ein-
schliesslich des hinzugefuegten Bytes,
gebildet

DESCENDING Einerkomplement des gesamten Schlues-
sels, nachdem alle anderen Attribute be-
ruecksichtigt wurden

Bine K-OUTPUT-Datei, die ohne Angabe von Attributen gebildet
wird, enthmelt die entsprechenden Daten unveraendert - so wie
sie in der Eingabedatei enthalten sind. Attribute wie NUMERIC-
ASCII und DESCENDING veraendern die Ursprungsdaten.

Das interne 9-Byte-Format, erzeugt duroh das NUNERIC-ASCII-
Attribut, ist eine eortierbare BCD-Gleitkommazahl:

Bytes/Bits Definitionen

Byte 0: Bit 7 1

Byte 0: Bit 6-0 normalisierter Exponent (Basis 10),
berechnet aus: Wert der Ziffer nach E,
wenn vorhanden, plus der Zahl der Zif-
fern vor dem Dezimalpunkt, plus 512;
Null, wenn keine signifikanten Zif-
fern vorhanden sind

Bytes 2-6 signifikante Ziffern, jede Ziffer

4 Bit, nach rechts mit 0 aufgefuellt

War die Zahl negativ, wird das Binerkomplement gebildet.

Das FLOATING-POINT-Attribut traneformiert den Schluessel wie
folgt:

. umgekehrte Bytereihenfolge
. wenn Byte 0 (nach der Umkehrung) 0 ist, werden alle Bytes
auf 0 gesetzt
. Bit 0 von Byte 0 erhaelt Wert von Bit 7 Byte 3
Byte 0 um eine Position nach rechts verschoben
Bit 7 von Byte 0 wird 1 gesetzt
wenn Bit 7 von Byte 1 vor Schritt 3 1 war, Einerkomplement
aller Bytes

ausw na

-59 -

SORT 1520 (SCPX)
2. R-OUTPUT (Satznummern) .

Eine R-OUTPUT-Datei enthzelt nur die Eingabesatznunmern. Die
Satznumnern beginnen bei 1, die maximale Satznunmer ist 65535.
\ird die Ausgabedatei ale CR-DELINITED spezifiziert, werden
die HNummern als ASCII-Zeichen nit Suehrenden lullen zum Auf-
fuellen auf 6 Ziffern ausgegebe

Ist die Ausgabedatei als PIXED-LENCTH, VARIABLE oder RELATIVE
deklariert, werden die Satznunmern binaer ausgegeben, je 3
Byte lang. Das erste Byte ist stets 0, das zweite Byte der L-
Teil der Satznuzmer und das dritte Byte der H-Teil.

Wird eine FIXED-LENGTH R-OUTPUT-Datei mit einem anderen
Programm gelesen, wird das Dateiende durch das physische Da-
teiende oder durch 1AH im ersten Byte (oder FFH beim FFZ2Z-
Ausgabeattribut) angezeigt.

4.3 UTPUT (Zeiger)

Eine P-OUTPUN-Datei enthaelt "Zeiger" zum Beginn eines jeden
Eingabesatzes. Jeder Zeiger besteht aus zwei Zahlen: der Sek-
tornummer (beginnend bei 1), die den Sektor bezeichnet, in dem
der Satz beginnt und einen Byte-Offset (Zahl zwischen 1 und
128), das anzeigt, mit wlchen Byte im Sektor der Satz be-
sinnt. Ist die Ausgabedatei CA-DELINITED, werden die Zahlen
als ASCII-Zeichen ausgegeben, 5 Ziffern fuer jeden Satz, durch
Kommas getrennt.

Ist die Ausgebedatei von anderen Typ, so werden die Zahlen als
2-Byte-Verte binaer ausgegeben, ier L-Teil zuerst.

Wird eine FIXED-LENGTH-P-OUTPUV-Datei mit einem Programm
gelesen, kann das Dateiende entweder durch das physische Da-
teiende oder durch einen Satz, dessen beide Zahlen aus 1AH
bestehen, erkannt werden (ist die Datei mit FPZZZ geschrieben,
dann -1 als erete Zahl).

4.4. KR-OUTPUT (Schlu: 1 und Satznummern)

Eine KR-OUTPUT-Datei enthaelt die Schluesael, gefolgt von den
Satznummern, wie oben beschrieben.

4.5. KP-OUTPUT (Schlu: 1 und Zeiger,

Eine KP-OUTPUT-Datei enthaelt die Schluessel, gefolgt von den
Zeigern zu den Eingabesaetzen, wie oben beschrieben.

-60-

SORT 1520 (SCPX)
2.11. ANHANG B: Fehlerausschriften

ein,;abe

Bei einem Fehler waehrend der Kommandoeingabe wird die
erkannte fehlerhafte Position durch ein Zeichen "(5EH) gekenn-
zeichnet, eine Fehlermeldung ausgeschrieben und die
Bereitschaft zur erneuten Kommandoeingabe angezeigt.

Tritt ein Fehler in einer Zeile nit mehreren Kommandos auf, so
sind die rechts vom fehlerhaften Kommando stehenden Kommandos
zu wiederholen, die links stehenden bleiben gueltig.

o
Bei Fehlerausschriften nach dem GO-Kommando erfolgt die Ruek-
kehr zur Kommandoeingabe, wo fehlende Kommandos ergaenzt beu.
fehlerhaft eingegebene Kommandos korrigiert werden koennen.
Verher eingegebene korrekte Kommandos bleiben gueltig.

Sortier-/l!ischlauf

fehler waehrend der Programmausfuehrung fuehren zun Abbruch.
Die benutzten Dateien werden nicht geschlossen, die Arbeite-
datei name.nun wird nicht geloescht.

Normalerweise wird nach einem solchen Fehler zum Betriebsays-
tem zurueckgekehrt. Wurde das RETURN-TO-CONSOLE-Kommando gege-
ben, wird mit einer Kommandoeingabe weitergearbeitet. Vorher
eingegebene Kommandos sind ungueltig.

Im folgenden werden die vollstaendigen Pehlerausschriften an-
gegeben; die Fehlernumner ist Bestandteil der Ausschrift.
Wurde das NO-ERROR-IIESSAGES-Kommando eingegeben, werden nur
"ERROR S" und die Tehlernummer ausgeschrieben.

ERROR 529: NOT BNOUGH WORKING STORAGE

erscheint mach 60; die Groesse des verfuegbaren Arbeitanpei-
chers reicht fuer die spezielle Kombination von Satzlaengen
und Optionen nicht aus

ERROR S30: COMMAND INPUT WORKSPACR FULL,
TRY MORB MEMORY OR SIMPLER COMMANDS

tritt ımehrend der Kommandoeingabe auf; das eingegebene Kom-
mando ist zu komplex fuer den verfuegbaren Speicherbereich, um
ausgefuehrt werden zu koennen

ERROR 537: INVALID FILE NAME
fehlerhafter Dateiname wmehrend der Kommandoeingabe: enthaelt

keine Duchstaben oder Ziffern vor der Typangabe (wenn
vorhanden) 'bzr. die Zeichen "*" oder "on

'
n
\

. SORT 1520 (SCPX)
ERROR S38: INVALID DRIVB NAME

unkorrekte Angabe des Laufwerks in einen Kommando

ERROR S39: UNRBCOGNIZED COMMAND

Das erste Wort des Kommandos ist lin Schluesseluort bzw.
keine zulaessige Abkuerzung eines Schluesselwortes.

ERROR S40: INVALID ARGUMENT

falsches Argument bei Kommandoeingabe

ERROR S41: MISSING ARGUMENT

fehlendes Argument bei Kommandoeingabe

ERROR S42: INVALID LINB TERMINATION

Kommandozeile fehlerhaft abgeschlossen

ERROR S43: NUMBER REQUIRED

fehlende Ziffer bei erforderlicher Zchlenangabe im Kommando
BRROR S44: NUMBER TOO LARGE

Zahl im Kommando groesser als 65535 (ausser numerische
Konstanten in SELECT- und EXCLUDE-Argumenten)

BRROR S45: RBCORD LENGTH MISSING

Kommandoeingabe: fehlende Satzlaenge im INPUT-ATTRIBUTES-Kom-
mando

ERROR S46: THAT COMMAND MUST BB LAST ON LINE

Kommandoeingabe: "60" und "CFILE" duerfen kein "
weiteres Kommando folgen

ERROR S50: NOT INPUT PILE SPRCIFIED

nach G0-Kommando: kein SORT-PILES- oder NERGE-FILES-Kommando
angegeben; nach der Spezifikationen der fehlenden Eingabedatei
kann erneut das G0O-Kommando eingegeben werden

ERROR S51: MORE THAN 32 SORT INPUT FILBS

ein

mehr als 32 Eingabedateien
ERROR S52: NO KEYS SPECIPIED

nach GO-Kommando: kein gueltiges KBY-Kommando eingegeben

2 SORT 1520 (SCPx)
ERROR S53: MORE THAN 32 KEYS
" mehr als 32 Schluessel vorgegeben
ERROR S54: MORB THAN 32 SELECT CRITBRIA

Kommandoeingabe: Anzahl der SELECT-Kommandos plus Anzahl der
EXCLUDE-Kommandos ist groesser als 32

ERROR S55: NO OUTPUT PILB SPRCIPIBD
nach G0-Kommando: kein OUTPUT-PILE-Kommando eingegeben
ERROR S56: PRINT LEVEL NOT 0 TO 5

waehrend Kommandoeingabe; keine exakte Vorgabe der Aus-
schriftsvartante

ERROR 557: RBCORD LENGTH NOT 1 TO 4096

waehrend Kommandoeingabe: falsche Ein- oder Ausgabedateilaen-
ge; bei VARIABLE-Dateien ist die minimale Satzlaenge 3

ERROR S58: FIBLD END COLUMN < START COLUMN

Kommandoeingabe: bei Positionsfeldspezifikationen in KEY-,
SELECT- und EXCLUDE-Kommandos, Feldende < Feldanfang

ERROR S59: #-FIELD NUMBER NOT 1 TO 255

Kommandoeingabe: bei KEY-, SELECT- und EXCLUDE-Kommafeldern,
Peldnummer nicht im Bereich 1 bis 255 .

ERROR S61: FIBLD END COLUMN > RECORD LENGTH

Kommandoeingabe oder nach 60: Positionsfeldangabe bei KEY,
SELECT und EXCLUDE, Peldende > Satzlaenge

ERROR S62: FIRLD START COLUMN < 1 OR > RECORD LENGTH

Kommandoeingabe oder G0: Positionsfelder bei KEY, SELECT und
EXCLUDE, Feldanfang < 1 oder > Satzlaenge

ERROR S63: FIRLD LENGTH NOT 1 TO 4096

Eingabe eines KBY-, SELECT- oder EXCLUDE-Kommandos,
Feldlaenge nicht im Bereich 1 bis 4096

BRROR S65: NO 'GO' AFTER ERROR IN COMMAND
FROM COMMAND PILE —
CORRECT ERROR, ENTER 'GO' AGAIN

erscheint nach Eingabe eines GO-Kommandos, wenn in einem

vorher eingegebenen Kommando ein Pehler enthalten ist (z.B,
bei einer CFILE-Kommandodatei)

-63.-

SORT 1520 (SCPX)
ERROR S67: TAGSORT WITH INPUT FILE SAME AS OUTPUT

GO-Kommando: TAGSORT angegeben und als Bin- und Ausgabedatei
gleiche Datei spezifiziert

ERROR S68: MORB THAN 32 MERGE-ONLY INPUT FILBS
waehrend Kommandoeingabe: mehr als 32 Nischeingabedateien

ERROR S69: NO INPUT RRCORD LENGTH SPECIFIED

nach GO-Kommando: kein INPUT-ATTRIBUTBS-Kommando eingegeben
ERROR S70: ) MISSING

bei Eingabe eines SELECT- oder EXCLUDE-Kommandos: Klammern
fehlen dort, wo sie entsprechend der Syntax erforderlich sind

BRROR S71: COMPARISON OPERATOR RBQUIRED

bei Eingabe eines SELBCT- oder BXCLUDE-Kommandos: nach der
Angabe eines Wertes (Feldspezifikation, Textkonstante oder
numerische Konstante) fehlt ein Vergleichsoperator

ERROR 572: FIBLD OR CONSTANT RBQUIRBD

bei Eingabe eines SELECT- oder EXCLUDE-Kommandos: vor oder
nach Vergleichsoperator fehlt Wertangabe (Feld, Text, Zahl)

ERROR 873: SBCOND " MISSING

Kommandoeingabe: ein abschliessendes Anfuehrungszeiohen (")
fehlt bei einer Textkonstanten

ERROR S74: ILLEGAL DIGIT POR NUMBER BASE

bei Bingabe eines SELECT- oder BXCLUDE-Kommandos: Zeichen 8
und 9 oder A bis P sind in einer Oktalkonstanten bzw. A bis F
in einer Dezimalkonstanten enthalten

ERROR S75: SIZB IN BYTBS IS TOO SMALL TO HOLD THB VALUR GIVEN

bei Eingabe eines SELECT- oder EXCLUDE-Kommandos: die signi-
fikanten Ziffern einer Konstanten erfordern zur Speicherung
mehr Eytes als nach der Basisangabe spezifiziert sind, z.B.
sei 12345H2 spezifiziert, so erfolgt eine Fehlerauasohrift, da
12345 hexadezimal nicht in 2 Bytes gespeichert werden kann

BRROR S76: ] MISSING "

bei Eingabe eines SELECT- oder EXCLUDE-Kommandos: weder eine
andere gueltige Konstante noch die schliessende Klammer (]) in
einer in eckigen Klammern [ ] einzuschliessenden Konstantenli-
ste gefunden

-66-

SORT 1520 (SCPX)
ERROR S78: 1-BYTR CONSTANT RBQUIRBD
bei Eingabe eines COLLATING-SEQUENCE-Kommandos : als Zeichen

mehr in Anfuehrungszeichen
) angegeben oder als Zahl,

eingeschlossene Zeichen (bzw. nur "
deren Wert groesser als 255 ist

ERROR S79: TOO MANY NBSTBD ()'s OR INTERNAL BRROR

bei Eingabe eines SELECT- oder EXCLUDB-Kommandos: mehr als 3
verschachtelte Klammern oder interner Fehler

ERROR S81: WORKING STORAGE FULL:
TRY USING MORB WORKING STORAGE, OR TRY TAGSORT.

waehrend Programmausfuehrung: eine relativ grosse Anzahl zu
sortierender oder zu mischender Eingabesaetze bei einem klei-
nen verfuegbaren Arbeitsspeicher fuehrt zum Veberlauf des
Arbeitsspeichers

ERROR S85: MERGR-ONLY INPUT PILE NAME SANB AS OUTPUT FILE

nach GO-Kommando: Ausgabedatei bei Mischen darf keine der
Eingabedateien sein

ERROR S86: TAGSORT, AND INPUT PILR DRIVE SAME AS OUTPUT /c
DRIVE

nach GO-Kommando, wenn bei TAGSURT und Wechsel der Ausgabedis-
kette fuer Ein- und Ausgabedatei das gleiche Laufwerk
angegeben wurde

ERROR S87: WORK FILE DRIVE SAME AS OUTPUT /C DRIVE

bei GO-Kommando: wenn Wechsel der Ausgabediskette spezifiziert
wurde, muss die Arbeitsdatei einem anderen Laufwerk zugeord-
net sein (die Arbeitsdatei ist dem aktuellen Laufwerk
zugeordnet, wenn nicht mit dem WORK-DRIVE-Kommando anderwei-
tig spezifiziert)

ERROR S88: INVALID COMBINATION OF FIELD TBST ATTRIBUTBS

eine unzulaessige Kombination von KEY- oder SELBCT-/BXCLUDE-
Feldtestattributen wurde spezifiziert, so z.B. NUMSERIC-ASCII
und INTEGER zusammen oder UPPER-CASB und PLOATING-POINT; auch
wenn LOHI fuer ein Kommafeld angegeben wurde;

Fehlerausschrift erscheint nach Kommandoeingabe

BRROR 589: MORB THAN ONB OF CR-DRLIMITED, FIXED, VARIABLE, AND
RBLATIVE

zwei verschiedene Dateitypen fuer Bingabe (oder Ausgabe)

spezifiziert (Eingabe und Ausgabe koennen aber verschiedenen
Typs sein)

=:65.=

SORT 1520 (SCPX)

ERROR S91: MORE THAN ONB INPUT FILE WITH OUTPUT OPTION OR
TAGSORT

R-OUTPUT, P-OUTPUT, KR-OUTPUT KP-OUTPUT duerfen nur mit
einer (Sortier- oder lisch-) ngabedatei verwendet werden;
TAGSORT darf nur mit einer Sortiereingabedatei benutzt werden.
set, K-OUTPUT koennen mehrere Eingabedateien angegeben wer-
den.

ERROR S92: MORB THAN ONB OUTPUT OPTION, OR AN OUTPUT OPTION
AND TAGSORT

nach G0: mehr als eine der Ausgabemoeglichkeiten K-,
KP-, KR-OUTPUT oder TAGSORT spezifiziert

ERROR S94: MERGE-ONLY INPUT FILE WITH TAGSORT

nach 60: TAGSORT darf nur bei Angabe einer Sortiereingabeda-
tei verwendet werden

ERROR S128: FILE name.typ NOT POUND

bei SORT-FILES- oder NERGE-FILBS-Kommando: angegebene Datei im
spezifizierten Laufwerk nicht gefunden

BRROR S129: DISK d: PULL WHILE WRITING PILB name.typ

Erecheint weahrend der Programmausfuehrung, wenn fuer die
gesamte Datei nicht genuegend freier Speicherplatz vorhanden
ist. Ist diese Datei die Arbeitsdatei (name.typ als SORT. nun
angezeigt), kann mit dem WORK-DRIVE-Kommando ein anderes Lauf-
werk zugewiesen bzw. TAGSORT zur Reduzierung der Groesse der
Arbeitsdatei benutzt werden.

ERROR S130: DIRBCTORY OP DISK d: FULL WHILE
WRITING PILE name.typ

waehrend Sortierlauf: Diekettenverzeichnis gefuellt

BRROR S131: CLOSB PAILURB PILE name.typ

interner Fehler: Diskette waehrend des Laufes gewechselt
BRROR S133: OVRRLONG RECORD (LENGTH nnn) IN PILR name.typ

erkannte Satzlaenge in einer VARIABLE-Eingabedatei laenger als
die spezifizierte maximale Satzlaenge :

ERROR S134: NOT A RRLATIVE FILE : FILB name.typ

Eingabedatei als RELATIVE spezifiziert, aber kein Dateikenn-
satz (mit einer Null im ersten Byte) vorhanden .

-66 -

SORT 1520 (SCPX)

ERROR S135: INCORRECT RBCORD LENGTH GIVEN
FOR RELATIVE PILE name.typ

waehrend Programmausfuehrung: spezifizierte Bingabesatzlaenge
entspricht nicht der im Dateikennsatz angegebenen

ERROR S140: COMMAND DECODING WORKSPACE FULL,
TRY MORB MEMORY OR SIMPLER COMMANDS

waehrend Kommandoeingabe: entweder ein sehr kompliziertes
Kommando (wahrscheinlich bei SELECT oder EXCLUDE) wurde bei
kleinem Arbeitsspeicherbereich eingegeben oder ein interner
Fehler

ERROR S160: INTBRNAL ERROR

ERROR S253: INTERNAL ERROR

interner Fehler

ERROR Snn: NO SUCH ERROR

ERROR-MESSAGE-Kommando mit undefinierter Fehlernummer einge-
geben oder interner Fehler.

67:2

DIENST 1520 (SCPX)

DIENST_1520_(SCPX)

DIENST ist eın kompaktes Programmpaket der im System enthal-
tenen Dienst- und Manıpulationsfunktionen. Neben auch in
anderen Dienst- und Hilfsprogrammen enthaltenen Funktionen
werden eine Reihe zusaetzlicher Moeglichkeiten (z.B.
Lesen/Schreiben ab Spur/Scktor, Diskettentest, Wiederzugaeng-
lichmachung von geloeschten Dateien) realisıert bzw. ist durch
das Nummernmenue cın effektiveres Arbeiten mit bereits vor-
handenen Funktionen moeglich. Das Programm belegt 16K.

Kommandauebersicht:

CHECK Kontrollsummenberechnung von Dateıen

cH Byteweiser Vergleich von Hauptspeiicherbereichen

copy Kopieren von Dateien

DIR Auflisten der Dateibezeichnungen des Verzeichnis-
ses (directory) einer Diskette auf dem Bildschirm

DISK Anzeige von Diskettensystemporametern

Ds Byteweıse Anzeige und Aenderung von Hauptspeicher-
Inhalten

DUMP/DUMPX/

DUMPH/DUMPA Anzeige von Hauptspeicher inhalten (in verschiede-
nen Codes)

ERA Loeschen von Datcıon

EXIT Verlassen des DIENST-Programmes und Rueckkehr in
den SCPX-Grundzustand

FILL Speicherbereich mit vorgegeben Zeichen auffuellen

co Laden und Starten eines Programmes (mit Adressan-
gabe)

GROUP Auflisten der durchnumerierten 2K-Bloecke der

Diskettendateien

IP/EX Starten von Programmen

LORD” Laden von Dateien

SAVE Speicher inhalt als Datei auf Diskette aufzeichnen
LOG Anzeige der aktuellen DIENST-Standardparameterein-

stellungen

- 68 -

MOVE

READ

WRITE

READGR

URITEGR

RECLAIM

REN

RESET

RUN

SEARCH

SET

SETDIR/
SETSYS

SET RO/
SETRO
SORT
SITZE
SPEED

STAT

TEST

TYPE/TYPEX/
TYPEH/TYPEA

USER

XUSER

DIENST 1520 (SCPX)

Verschiebung von zusammenhaengenden Speicherberei-
chen

Lesen von Diskette (ab Spur/Sektor) in den Speı-
cher oder auf den Bildschirm

Aufzeichnen von im Speicher stehenden Daten auf
Diskette (ab Spur/Sektor)

Lesen von Diskette (ab Blocknummer) in den Spei-
cher oder auf den Bildschirm

Aufzeichnen von im Speicher stehenden Daten auf
Diskette (ab Blocknummer)

Wırderzugaengl ıchmachung von aus dem Verzeichnis
geloeschten Dateien

Umbenennung von Dateien

Ruecksetzen einzelner Laufwerke in den Grundzu-
stand :

Laden und Starten von Programmen

Suchen von definierter Rytefolge im angegebenen
Hauptspeicherbereich

Setzen von Merkmalen in Dateibezeichnung

Wandlung DIR-Dateien (--) SYS-Dateien

Wandlung schreibgeschuetzte (--) nichtschreibge-
schuetzte Dateien

Sortieren der Verzeichniseintragungen

Meldung ueber die Groesse von Dateien

Festlegung der Bildschirmausgabegeschwindigkeit

Kapazitaetsmeldung ueber freien bzu. belegten
Speicherbereich cıner Diskette

Test von Disketten

Auflisten von Dateien auf den Bildschirm (in
verschiedenen Codes)

Einstellen Benutzerbereich

Festlegung Zielbenutzerbereich beim Kopieren

6

DIENST 1520 (SCPX)

UR Anspringen spezieller zusaetzlich vom Anwender in
DIENST eingebundener Routinen

Der Aufruf von DIENST erfolgt durch Eingabe von:
AYDIENST
Das Programm wird geladen, und es erscheint die Ausschr ift:
DIENST 1520 (SCPX) V 1.0
WESENTLICHE SCP-DIENSTFUNKTIONEN
AO=
AO = ist die Bereitschaftsmeldung zur Kommandoeingabe (A Ist
das Laufwerk, von dem DIENST aufgerufen wurde, 0 Ist das User-
Feld). Wird ein anderes Laufwerk (LU) gewuenscht, kann dies
mit der Eingabe der LU-Bezeichnung (A-E) angegeben werden

(z.B AO=bi, as erscheint BO=. B ist somit das aktuelle
Laufwerk).

Das Kommando zur Rusckkehr in das SCP-Kommandoe Ingaben iveau
ist EXIT.

Werden waehrend der Arbeit mit DIENST in einem Laufwerk die
Disketten geuechselt, dann ist ein Warmstart auszufuchren, um
die Diskette R/U zu setzen. Andernfalls erscheint bei Schraib-
zugriff eine entsprechende Aufforderung auf dem Bildschirm.

Zum Abbruch eines Kommandos in DIENST wird AC oder die ESCABE-
Taste verwendet.

Alle DIENST-Kommandos koennen durch Eingabe eines Fragezei-
chens (?) angezeigt werden.

z.B.: (A0=?)

will man nur ein DIENST-Kommando abarbeiten und dann In das
SCP-Kommandoeingabeniveau zurueckkehren, gibt man ein:

A>DIENST Kommando
2.B.: A)DIENST CHECK

Bei einigen Kommandos koennen zusaetzliche Parameter, in ecki-
gen Klammern eingeschlossen, angegeben werden.

2.B.: A0=sCOPY [0] @ - Umbenennen des Dateinamens (hier
waehrend des Kopierens)

Fehlen die Parameter, gelten die mit dem DIENST-Kommando LOB
anzeigbaren Standardeinstellungen.

Achtung: In Laufwerk A muss immer e

e Systendiskette liegen!

-%-

DIENST 1520 (SCPX)

10

3

Einige Kommandos zeigen bei ihrem Aufruf das Dateiverzeichnis
oder Teile daraus (Dateigruppen) an, wobei die Dateien mit
einer laufenden Nummer versehen sind. Diese Nummer kennzeich-
net die Datei Fuer die weitere Arbeit mit Kommandos.

Nach der Frage "WAEHLE?” koennen eine, mehrere oder Gruppen
von Datelen angegeben werden.

Einzelne Nummern sind durch Leerzeichen zu trennen; Bereiche
werden durch Anfang und Ende mit Minus (-) verbunden angega-
ben. Fehlt die Angabe des Endes, wird bis zur letzten Nummer
gewaehlt.

Belspiel: WAEHLE? 1 3 4-8 10-

Die zugewiesenen Nummern haben keine bleibende Verbindung zu
den Datelen, sie koennen veraendert werden (z.B: durch Loe-
schen oder Umsortieren von Dateien).

Boi der Eingabe der Kommandos, dic das Nummernmenue aufrufen,
albt es folgende Moeglichkeiten (als Beispiel dient REN):

(1) AO=REN CENTER) - Anzeige Nummernmenue aktuelles LU A
(2) AO=REN B: XENTER) Anzeige Nummernmenue von LU B
aktuelles LU bleibt A

Umschalten aktuelles LU von A auf B
Anzeige Nummernmenue von LW B
Umschalten aktuelles LU von A auf B
Anzeige Nummernsenue von LU C
Anzeige Nummernmenue laut Parametern
(z.B: Spaltenangabe)--) siehe LOG-
Kommando!

Anzeige Nummernmenue der Dateien,
deren Namen mit T beginnt

(w.® - allgemeine SCP-Dateibezeich-
nung»

wie (6), aber die DIENST-Datelbe-
zeichnung nutzend

Anzeige Nummernmenue aller Dateien
des LU A

(##* -DIENST-Bezeichnung fuer belie-
bige Dateien)

(3) AO=B:REN (ENTER)

(4) AO=B:REN Ci (ENTER)

(5) AO=REN L I (ENTER)

(6) AO=REN T#.# (ENTER)

(7) AO=REN T#® (ENTER)

(8) AO=REN #u# (ENTER)

(9) AO=REN 5727.COM
KENTER) - Anzeige aller .COM-Dateien, deren
Namen 4 Zeichen lang sind und an
erster Stelle eın S und an dritter
Stelle ein Z haben

Komma: _Dateinamen)

Bei dieser Kommandoeingabe wird die exakte Dateibezeichnung
verwendet.

= 71%

DIENST 1520 (SCPX)
Es gibt folgende Moeglichkeiten (als Beispiel dient TYPE:

(1) AO=TYPE XXX.DOX (ENTER) - Kommando zur Ausgabe der Datei
xXxx.DOK von LU A
(2) AO=TYPE B:XXX.DOK

XENTER) - wie (1), aber von LU B, A bleibt

aktuelles LU
(3) AO=B:TYPE XXX.DOK

ENTER) - Umschalten aktuelles LU von A
auf B, Ausgabe der Datei von B

Bei einigen Kommandos ist nur die genaue Komnandobezeichnung
anzugeben. Es gibt Folgende Moeglichkeiten (als Beispiel dien
TEST):

(1) AO=TEST (ENTER) - Test Diskette in LU A

(2) AO=TEST B: CENTER) - Test Diskette in LU B, aktuelles LW
bleibt A

(3) AO=B:TEST (ENTER) - Umschaltung aktuelles LU von A auf

B, Test Diskette in LU B

sofortiges Beenden aller Ausgubeoperationen

- es kann folgen:

Leertaste noch einmal --) Zeilenvorschub
ENTER) --) Anzeige neue Bildschirmseite
andere Taste --) fortlaufende Anzeige

ENIER_oderAH,_.

- zum Vorschub (z.B. bei Auflisten eines grossen Diskettenver-
zeichnisses)

- es wird jeweils eine neue Bildschirmseite ausgegeben

AP:

- Druckereinschaltung (nur fortlaufend, nicht seitenwelse
moeglich)

- eine im CCP-Grundzustand erfolgte Druckerzuschaltung kann
innerhalb DIENST nicht rueckgaengig gemacht werden

(Achtung? Bei nochmaligem AP erfolgt Zeichandruck dop-
pelt!)

- 72 -

DIENST 1520 (SCPX)

- Eingabe einer dieser Nummern waehrend der Ausgabe macht
seitenuelse Ausgabe ungueltig und setzt Geschwindigkeit
der zeilenweisen Ausgabe (0-schnell, 9-langsam:

- siehe SPEED-Komsando

Beli t

- macht seitenweise Ausgabe ungueltig
- zeilenueise fortlaufende Ausgabe

3.3.2._Steuertasten_|
DIENST besitzt einen Mini-Editor, der es erlaubt, Eingabefeh-

le zu loeschen und zu korrigieren. Ausserdem arbeiten die
herkoemml Ichen SCP-Steuerzeichen auch im DIENST-Programm.

ingabe

DELETE/BACKSPACE/AH - Loeschen letztes eingegebenes Zei-
chen, danach Neueingabe moeglich

ENTER AM/AJ - Abschluss Kommandoeingabe und Aus-
fuchrungsbeginn

E - Kursor auf neue Zeile setzen (bel
langen Kommandoanwe sungen)

ap - Druckerzuschaltung

AR - nochmal ige Anzeige einer eingege-
benen und korrigierten Eingabezeile

au - Ungueltigmachen der Kommandoe inga-
be, Kursor auf Anfang neue Zeile
setzen

ax - Ungueltigwachen der Kommandoe inga-
be, Kursor auf Anfang gleiche Zeile
setzen

3.4._Kommandos

3:4:12_?_(H
Ao=?

Es wird eine Liste aller DIENST-Kommandos auf dem Bildschirm
angezeigt. Als Hinweis auf dieses Kommando erscheint bei feh-

lerhaften Kommandoeingaben (Syntaxfehler) die Bildschirmaus-
schrift "incorrect, for list of commands enter (?)”.

-73 -

DIENST 1520 (SCPX)

CHECK dient zum Errechnen von Kontrollssummen einzelner Da-
teien. Der Anwender kann schnell pruefen, ob zwei Dateien
identisch sind. Fuer weitere Untersuchungen im Falle Nichti-
dentitaet kann das Speichervergleichskommando CM zum Finden
unterschledlicher Bytes genutzt werden, oder man listet mit
dem TYPE-Kommando die Dateien komplett auf Bildschirm und
Drucker aus. Die laufende Summe der Dateikontrollsummen kann
fuer den Vergleich von Dateigruppen auf Identitaet herangezo-
gen werden. Fuer die Dateiauswahl kann die Syntax des Nummern-
menues verwendet werden. Die berechneten Kontrollsummen zur
jeweiligen Date) werden auf dem Bildschirm angezeigt.

Beispiel:

AOSCHECK (ENTER)

AO: 1= DOK „DOK | 2= TEST .COM | 3= TKO „DOL
AO: 4= DIENST .COM | S= DIENST .DOK

WAEHLE? 2 3

A: TEST „COM -checksum: DIDB total: DIDB dec: 53723
A: TKO .DOK -checksum: 368B total: 0866 dec: 2150
A0=
checksum - aktuelle Kontrollsumme der Diskettendatel in HEX-
Darstellung
total - laufende Summe der Kontrollsummen In HEX
dec - total-Uert in DEZ

32423._CH
AO=CH (aadri) (eadri) (aadr2)

taadr1) - Anfangsadresse erster Speicherbereich In HEX
teadr1) - Endadrasse erster Speicherbereich In HEX
(aadr2) - Anfangsadresse zweiter Speicherbereich In HEX

Das Kommando dient zum byteweisen Vergleich zueler Speicher-
bereiche. Die Laenge der zu vergleichenden Bereiche wird auto-
matisch aus der Differenz von End- und Anfangsadresse des er-
sten Speicherbereiches berechnet. Deshalb erfolgt keine Ein-
gabe der Endadresse des zueiten Speicherbereiches. Das Komman-
do ermoeglicht die Kontrolle zweier Dateien auf Identitaet,
wenn vorher mit dem DIENST-Kommando LOAD diese Dateien auf die
entsprechenden Bereiche gebracht wurden. Es werden nur die
unterschiedlichen Bytes mit Adressen und Inhalt angezeigt.

-74 -

DIENST 1520 (SCPX)

AO=CH 4000 4003 5000 (ENTER)
4000 43 54 5000
4002 44 55 5002
tAdr. Inh.Inh. Adr.)

AO=COPY Lpar]

coPY dient zum Kopieren einer beliebigen Anzahl von Dateien
auf ein anderes Laufwerk oder auf die Quelldiskette unter
anderem Naman (Parameter Q angeben). In Verbindung mit dem
Kommando XUSER koennen auf der Quelldiskette In andaren
Nutzerbereichen Kopien abgelegt werden (siehe 3.4.37.). COPY
kann mit dem Nummernmenue arbeiten, so dass nur die Nummern
der zu kopierenden Dateien und das Ziellaufwerk eingeben war-
den muassen. Zur Kontrolle wird der Name der im Kopiervorgang
befindlichen Datei auf dem Bildschirm angezeigt. Erscheint
eine Datei, die nicht kopiert werden soll, so drueckt man die
ESCAPE-Taste, wodurch der Kopiervorgang abgebrochen wird. Auf
der Zieldiskette befindet sich dann eine Pseudodatel
"Name.xx#”, die geloescht werden muss. Bei diesem Kommando Ist
die Angabe von Parametern In L ] moeglich. Vor dieser Eingabe
Ist deshalb das LOG-Kommando zu aktivieren, um zu kontrollie-
ren, ule der Standard gesetzt ist. Vor dem Kopieren kann mit
SIZE und STAT herausgefunden werden, ob die zu kopierenden
Dateien noch auf der Zieldiskette Platz haben. Wird beim
Kopieren BAD SECTOR festgestellt, so erscheint die Abfrage
"bad sector on read, abort (Y/N):” (Abbruch des Koplarvorgang
JA/NEIN). Bei Eingabe N wird das Kopiaren fortgesetzt, wobel
der Anwender anschliessend mit TYPE die kopierte Datei "auf der
Zieldiskette kontrollieren sollte.

Wieht ige_COl

A - Existiert die Datei auf dem Ziellaufwerk bereits,
dann Ueberschreiben dieser (overlay)

B - Existiert die Datei auf dem Ziellaufwerk bereits,
dann automatisches Mandeln in eine BAK-Datei und
Kopieren der neuen Datei (back up)

c - Ist die Datei auf dem Ziellaufwerk bereits vorhan-
den, dann Frage ob: (0) verlay, (B) ack up or
(Skip:
(Eingabe 0, B oder S --) siehe A-, B- und D-
Parameter)

D - Existiert die Datei auf dem Ziellaufwerk bereits,

dann kein Kopieren (skip)

R con)

(Y/N)-Entscheidungsfrage vor Kopieren jeder ein
zelnen Datei

- 75 -

DIENST 1520 (SCPX)

v con Lesen der kopierten Datei zur Kontrolle auf

Schreibfehler

T con

Abbruch des COPY beim Kopieren einer Dateiserie,
bei jener Datei in der eingegebenen Reihenfolge,
die nicht auf die Zieldiskette passt

T OFF) - Passt die Dateiserie nicht komplett auf die Ziel-
diskette, werden soviel Dateien wie moeglich
uebernommen (d.h. der zu grossen Datei folgende
kleinere Dateien, die noch Platz haben, werden
kopiert)

M con)

Markieren der Dateien im Verzeichnis nach dem Ko-

pieren --) «€ Driglinaldatei (auf Quelldiskette)
(Marke kann durch Kommando SET [+21
geloescht werden)

--).) kopierte Datei (auf Zieldiskette)
(Marke kann durch Kommando SET [+1]
geloescht werden)

a - erlaubt Umbenennen der Dateien, die kopiert werden
damit z.B. sofortiges Anlegen eines Duplikates
einer Datei auf der gleichen Diskette moeglich ist

Die Parameter mit (ON/OFF) schalten Ihre Wirkung um, wenn sie
ein zueites Mal angegeben werden.

Es koennen mehrere Parameter gleichzeitig angegeben werden
(z.B: AO=COPY LCROMI).

Die Parameter A, B, C und D schliessen einander aus.

324.5._DIB
AO=DIR LUXII

Das Kommando dient zur Anzeige von Bibliotheksverzeichnissen.
Es besteht die Moeglichkeit, die Verzeichnisse aller Laufwer-
ke und USER-Felder aufzulisten.

L - alle USER-Felder (wenn nicht angegeben, dann nur aktuelles
USER-Feld)

X - alle Laufwerke (wenn nicht angegeben, dann nur aktuelles
Laufwerk)

1 - Anzahl der Anzeigespalten (wenn nicht angegeben, dann
entsprechend Spaltenangabe bei LOG)

AO=DIR LUX22

Parameter koennen auch einzeln oder gar nicht angegeben

Durch Eingabe mehrdeutiger Dateibezeichnungen koennen auch
bestimmte Dateigruppen angezeigt werden. Dabei wird die SCP-
Bezeichnung (#.*) ebenso akzeptiert, wie die spezielle DIENST-
Bezeichnung (##).

-7 -

DIENST 1520 (SCPX)

"  (SCP: DIR Name.x)
: DIR #.Typ)

- gleicher Dateityp: DIR Na;
- gleicher Dateiname: DIR ##Typ (SCI

Die angezeigten Dateien koennen folgende Kennzeichnung be-
sitzen:

® - Datei ist R/O gesetzt
€) - Systemdatei
(® - R/O und Systemdatel

AO=DISK

Dieses Kommando zeigt die vichtinsten Systemparameter der DI
ketten (Diskettenformat, Kapazitaet, Dichte, Informationen
ueber Systemspur) an.

Beispie

Disk: A: ;Disketten-LU A

Disk capacity: 296K ;Diskettenkapazitaet

Trackst 77 3 System 377 Spuren,
davon 3 Systemspuren

Sectors/tracks: 32 32 last ;32 Sektoren/Spur, letzte
Spur hat 32 Sektoren

Sectors/system: 96 16 dir ;96 Sektoren fuer System
und 16 Sektoren fuer
Verzeichnis

Dir entries: 64 2x Verzeichnis kann 64
Eintragungen aufnehmen
(auf 2K)

Sectors/group: 16 2K 93H group z;Block aus 16 Sektoren be-

legt 2K5  groesstmoegl i-
che Blocknummer (in HEX)
auf der Diskette = 93H

KRyte/extent: 32K jein Extent belegt 32K
(gemeint ıst die maximale
Dateigroesse, auf die
sich eine Verzeichnisein-
tragung bezieht)

3

AO=DS adr

Das Kommando dient zur byteweisen Anzeige und zur Aenderung
von Speicher inhalten. Die Anzeıge erfolgt ab der vorgegebenen
Adresse in HEX-, DEZIMAL-, BINARY- und ASCII-Code. In der
Spalte Enter wird die Codeart angezeigt, in dem das angezeigte
Byte bzw. die folgenden Bytes geaendert werden koennen CH),
‘D), <B> oder <A»). Dieser Eingabemodus kann durch
pachfolgende Eingabe der gewuenschten Codeart wumdefiniert
werden (.H, .D, .B oder .A). Es koennen bis zu 120 Bytes, die

72 =

DIENST 1520 (SCPxX)

geaendert werden sollen, in gewuenschter Eingabemoduskombina-
tion aufeinanderfolgend und durch Leerzeichen getrennt
eingegeben werden. Die Byteanzeige kann adresseitig vor- und
Fueckwaertsbewegend erfolgen (Umschaltung mit TILDE-Taste).

pi

AO=DS 5000 (ENTER)
addr: hex dec binary ascii Enter

C3 195 11000011 -C <A) KENTER)

10 16 00010000 AP XA).H 13 63 CENTER)

1A 26 00011010 AZ “HI 7»

63 99 01100011 < <H> <ENTER)

13 19 00010011 AS “BB cm)

63 99 01100011 < CH) CENTER)

1A 26 00011010 az <H).A a bc .D 10 11 CENTER)

2E 46 00101110 = XD)... KENTER)

Bei ASCII-Code bedeutet das Minuszeichen (bei Adresse 5000:
-C), dass das adresshoschste Bit gasetzt Ist, das A-Zeichen
(z.B. AP) verweist auf den Steuerfunktionscharakter eines
Codes. x

Ecleeuterung_der_ Ich

.A - Eingabemodus ASCII-Code

=) os werden nur ASCII-Zeichen akzept lert

=) Zeichen sind durch ein Leerzeichen zu trennen

=) Loerzeichen kann nicht durch ein ASCII-Zei-
chen angegeben warden (.H20 oder .D32 verwan-
den)

=) eingegebene Steuerzeichen werden entweder aus-
gefuehrt (z.B. AP) oder angezeigt (z.B. 18)

ingabez

-B - Eingabemodus Binaercode
=) os werden nur Binaerwoerter akzeptiert
-) Binaerwoerter sind durch Leerzeichen zu trennen
=) alle 8-Bitpositlonen muessen angegeben werden

.D - Eingabemodus Dezimalcode
=) 25 werden nur Dezimalzahlen akzeptiert
=) Dezimalzahlen sind durch Leerzeichen zu trennen

«H - Eingabemodus HEX-Code
=) es wird nur der HEX-Code akzeptiert
=) Hex-Zahlen sind durch Leerzeichen zu trennen

ENTER) - Vorruecken auf neue Adresszeile
=) wird nur der Eingabemodus geaendert, dann An-
zeige der gleichen Adresszeile mit neuem Modus
=) zu beachten: AM und AJ erzielen die gleiche
Wirkung wie (ENTER)

-7-

DIENST 1520 (SCPX)

“> - Umkehr der Richtung der Speicheranzeıge
(TILDE) =) wenn Speicher vorwaerts durchschritten wird,
erscheint Doppelpunkt nach der Adressanzeige
=) wenn Speicher rueckwaerts durchschritten wird,
erscheint Minus (-) nach der Adressanzeige

„„(ENTER) - beendet DS-Kommando

Ausserdem kann die Funktion genutzt werden, um von einem
Zeichen alle Codedarstellungsarten anzuzeigen (z.B. den
unbekannten Binaercode eines bestimmten ASCII-Zeichens), indem
man das Zeichen auf einen unbelegten Speicherplatz eing|bt und
mit TILDE zu diesem Speicherplatz zurueckkehrt.

Diese Kommandos dienen zur Inhaltsanzeige angegebener Haupt-
spelcharbereiche auf dem Bildschirm bzw., falls vorher AP
eingegeban wird, zum gleichzeitigen Protokollieren auf dem
Drucker.

Das Kommando kann In Verbindung mit DS und SAVE bel
Aenderungsarbelten herangezogen werden (z.B. Wiederherstellen
zerstoorter Textdateien). Die entsprechenden Dateien koennen
mit LOAD, READ oder READGR in den Speicher gebracht werden.
DUMPP und DUMPA werden automatisch beendet, wenn dia
Dateiendemarke 1AH (AZ) gafunden wird.

Falls mahr Informationen angozaigt werden sollen als mit einem
Mal auf dem Bildschirm Platz haben, erfolgt bei gafuelltem
Bildschirm aln Stopp.

Weiterarbeitsmoegl Ichkelten:

ENTER) - Anzalge naechste Selte
“Leorzeichen) = naechste Zeile wird "hineingerollt”
beliebige andere Taste - Durchrollen der restlichen Informa-
tionen
- Abbruch ist mit ESCAPE-Taste moag-
lich
DUMP - Anzeige des Speicherbereiches im ASCII-Code (ist der
Datei mit Tabulatoren formatiert, dann formatierte
Ausgabe)
DUMPX - Anzeige des Speicherbereiches ım HEX- und ASCII -CLode
“in numerierten Zeilen zu je 16 Byte)
DUMPH - Anzeige des Speicherbereiches im HEX-Code (In nume-
rierten Zeilen zu je 1& Byte)
DUMPA - Anzeige des Speicherbereiches im ASCII-Code (in nume-

rierten Zeilen zu je 16 Byte)

-9-

DIENST 1520 (SCPX)

(DUMPX als Beispiel)

(1) AO=DUMPX (ENTER) - Ausgabe von 128 Byte des Speichers
- Start bei Endadresse des letzten
DUMP-Kommandos

(2)AO=DUMPX ‚6 (ENTER) - Ausgabe von n Byte ab zuletzt
verwendeter Adresse (im Beispiel 6
Byte)

(3)AO=DUNPX ,, (ENTER)  - Ausgabe gesamter Speicher ab zu-
letzt verwendeter Adrosse

(4) AO=DUMPX 100 (ENTER) - Ausgabe des Bytes, das auf der
angegebenen Adresse steht (im Bei-
spiel 100H)

(5) AO=DUMPX 100,0F «ENTER)- Ausgabe von n Byte in HEX (im
Beispiel OFH=15 dezimal) ab ange-

gebener Adresse (im Beispiel LOÖH)

(6) AO=DUMPX 100, CENTER? - Ausgabe von 128 Byte ab angegebe-
ner Adresse (im Beispiel 100H)
Ausgabe des gesamten Speicher in-
haltes ab angegebener Adresse (im
Beispiel 1004)

'

(7)AO=DUMPX 100,, (ENTER)

Fuer die Bildschirmausgabe sind die Bildschirmsteuertasten
13.3.1.) zu beachten.

Das Kommando dient zum Loeschen von Dateien. Es kann mit dom
Nummernmenue gearbeitet werden. Nach Auswahl der zu loeschen-
den Dateien erscheint die Abfrage "ERASE (Y/N)?”. Wachrend des
Loeschens wird der Name jeder ausgewachlten Datei angezeigt.
Versehentliches Loeschen von Dateıen kann mit dem RECLAIN-
Kommando rueckgaengig gemacht werden. Allerdings ist dies nur
bis zum naechsten Schreibzugriff auf dıe Diskette sinnvoll,
weil dann die freigegebenen Bereiche bereits wieder ver-
faelscht sein koennen. Mit ERA kann der [RJ-Parameter angege-
ben werden, wodurch bei jeder zu loeschenden Datei eine noch-
malige Entscheidung ueber die Ausfuehrung moeglich ist. Da
dadurch die DIENST-Parametereinstellung R [Forderung einer
(Y/N Antwort vor Kommandbausfuehrung --) siehe LOG-Kommandol
auf ON) gebracht wird, sollte nach Beendigen des Loeschens
noch einmal ERA LRI zum Rueckstellen auf (OFF) eingegeben
werden.

- 89 -

DIENST 1520 (SCPX>

AO=ERA [RI (ENTER)

AO DOK .DOK | 2= ASM „COM | 3=EDIT „COM
a0: LINK „DOK
WAEHLE? 1 3

ERASE (Y/N)? <Y) yes

DOK .DOK (Y/N (Y) yes
EDIT .COM Y/N) <NO> no

ao;
Das Nummernmenue kann beim Loeschen nur einer Datei durch
direkte Angabe des Dateinamens nach ERA uebergangen werden.

Bei Angabe eines mehrdeutigen Dateinamens werden ım Nummernme-
nue nur diesem Dateinamen entsprechende Dateren angezeigt.

324.10._EXIT
AO=EXIT

Mit diesem Kommando wird ın den SCPX-Grundzustand
Zurveckgekehrt, wobei die USER-Einstellung erhalten bleibt und
das vor dem Aufruf von DIENST aktuelle Laufwerk eingestellt
wird. Genau wie bei AC, mit dem ın den DIENST-Grundzustand
Zurueckgekehrt werden kann, erfolgt ein automat scher
Speichertest. Die Anzeige eines "?"-Zeichens bedeutet Fehler,
worauf ein umfassender Speichertest eingeleitet werden sollt-.

AO=FILL (aadr) (eadr) (byte)

(andr) - Anfangsadresse Speicherbereich in HEX
(eadr) - Endadresse Speicherbereich ın HEX
(byte) - Angabe eines Bytes in HEX

Mit diesem Kommando kann eın Speicherbereich mit einem vorge-
gebenen Zeichen aufgefuellt werden.

B

Dielz: AO=FILL 4000 5000 20 (ENTER)

Es wird genutzt, um vor Ueberarbeitung eınes Programmes den
Speicher mit Hexanullen aufzufuellen. Es ıst guenstig, vor
Eingaben oder Einlesen den Speicher mıt Space zu fucllen, um
bei anschliessender Nutzung von DUMPX klar dıe Grenzen der
eingegebenen Daten zu erkennen. Weiterhin hilft es bei der
Manipulierung zerstoerter Disketteninhalte, indem der Speicher
mit ES helegt wird und mit URITE auf der Diskette die funk-
tionstuecht igen Teıle isoliert werden.

-8-

DIENST 1520 (SCPX)

3.4.12._60

AO=60 (Dateibezeichnung) (adr)
adr) - Speicheranfangsadresse in HEX

60 ist eine spezifische Form des RUN-Kommandos, bei der die
gewuenschte Programmdatei auf die angegebene Speicheradresse
geladen und anschliessend das Programm gestartet wird.

Belspiel: AO=G0 BEI.COM 2000 CENTER»

Dadurch Ist es moeglich, spezielle Programme abzuarbeiten, die
nicht die vebliche Anfangsadresse 100H haben. Wenn diese Pro-
gramme mit einem RET abschliessen, wird zu DIENST zurueckge-
kehrt. Ausserdem koennen mit GO ueber Tastatur eingegebene
Parameter an das abzuarbeitende Programm uebergeben werden. Im
Programm muessen sie vom DIENST-Eingabepuffer geholt warden.
Die Parameteradresse wird Im Register HL bereitgestellt.

A0=G0 BEI.CON 2000 B:ZUSA.COM DATEI.NEW

AO=GROUP

Das Kommando listet die durchnumer ierten 2K-Bloecke auf @iner
Diskette auf, die jeweils eine Datei bilden. Damit kann he-
rausgefunden werden, wie die Dateien physisch auf der Diskette
vertellt sind. Bei der Anzeige erscheint nach dem Dateinamen
als erstes die Extent-Angabe (ein Extent = 16 K). Es folgen
die Blocknummern. Das GROUP-Kommando kann mit dam Nummernmenua
arbeiten.

Baisplel:

AO=GROUP B: (ENTER)

BO: 1= MAT .DOK | 2= TEST „COM | 3= TKO „Doc
BO: 4= DIENST .COM | 5= DIENST .DOK

WAEHLE? 1-3

B:MAT .DOK 00 01 02 .

„COM 00 03 04 05 06
.DDK 01 09 OA OB OC OD OE OF 10 11 12 13 14 15 16

02 19 1A ıB iC ıD

Das Nummernmenue kann bei Auflisten nur einer Datel durch
direkte Angabe des Dateinamens nach GROUP uebergangen werden.
Bei Angabe cines mehrdeutigen Dateinamens werden Im
Nummernmenue nur diesem Dateinamen entsprechende Dateien
angezeigt.

- 82 -

DIENST 1520 (SCPX)

Falls eine grosse Datei lange Zugriffszeiten hat, kann mit
GROUP eine eventuell vorhandene unguenstige Verteilung auf der
Diskette nachgeprueft und durch Kopieren auf eine andere Dis-
kette beseitigt werden. Zum Lesen einzelner 2K-Bloecke einer
Datei kann das READGR-Kommando genutzt werden.

3.4.142_JP/EX

AO=JP (adr) (argument)
AQ=EX (adr) (argument)

Kadr) - Sprungadresse
(argument) - moegliche Uebergabeparameter

Mit diesen beiden Kommandos koennen im Speicher befindliche
Programme gestartet werden, indem die Programmanfangsadresse
angegeben wird. Ab der angegegebenen Sprungadresse muss ein
abarbeltungsfachiges Programm stehen, das mit RET (oder bei
gewuenschtem Warmstart mit BDOS-Funktionsaufruf O oder JP_ 0)
endet. Beim Kommando JP wird nach Programmausfuehrung zum SCP-
Warmstart (zu Adresse 0) gesprungen. Beim Kommando EX wird
nach Programmausfuehrung In den DIENST-Grundzustand zurueckga-
kehrt.

Beispieli AO=JP 2000

Boi beiden Kommandos koennen ueber Tastatur eingegebene Para-
meter an das abzuarbeitende Programm uebergaben warden. Im
Programm muessen sie vom DIENST-Eingabepuffer geholt warden.
Die Pufferadresse wird im Register HL bereitgestellt.

elsplalz AO=JP 2000 MAT.DAT

3:4.15.1._LOAD
AO=LOAD (dateibezeichnung) (adr)

(datelbazeichnung) - eindeutiger Bezeichner der zu ladenden
Datei
tadr) - Speicheranfangsadresse In HEX

Mit dem LOAD-Kommando kann eine Datei von der Diskette in den
Speicher geladen werden. In Verbindung mit dem DS-Kommando
(Anzeige/Aenderung vom Speicher inhalt) und dem SAVE-Kommando
(Schreiben Speicherinhalt auf Diskette) wird es zum schnellen
Aendern von Dateien genutzt.

Vor Ausfuehrung von LOAD muss analysiert warden (mit Hilfe des
SIZE-Kommandos), wie gross die zu ladende Datei und der zur
Verfuegung stehende TPA-Anwenderbereich (mit Hilfe des LOB-
Kommandos) sind, um eıu versehentliches Ueberschreiben des
Betriebssystems bzw. DIENST-Programmes zu verhindern.

- 83 -

DIENST 1520 (SCPX)

Zu grosse Dateien muessen in Sektionen zerlegt und diese
einzeln zur Bearbeitung geladen werden (mit Hilfe von GROUP
und READGR)I.

Nach Ausfuehrung von LOAD wird automatisch die beim Laden
erreichte Speicherendadresse und die durch die Datei belegte
Anzahl von Sektoren (als noetige Information fuer SAVE
vermerken!) angezeigt. Werden .COM-Dateien geladen, die
normalerweise ab 1004 beginnen, so sollte zur rechner ischen
Vereinfachung eine Anfangsadresse K+100H (z.B. 51004) gewachlt
werden. LOAD bietet die Moeglichkeit, Programme in den
Speicher zu laden, wenn deren normaler Speicherbelegungsbe-
reich andersueitig benoetigt wird.

Beispiel: AO=LOAD TEST.COM 5100 (ENTER)

3.4.19.2._SAVE
AO=SAVE (dateibezeichnung) (adr) (sekt)

(dateibezeichnung) - eindeutiger Bezeichner der aufzuzeichnen-

den Datei

Kadr) - Speicheranfangsadresse In HEX

(sekt) - Anzahl der zu rettendon 128 Ryte-Sektoren
tAchtun Unterschied zum residenten

Kommmando SAVE!)

Mit SAVE kann der Speicher inhalt ab der angegebenen Adresse
auf Diskette aufgezeichnet werden.

Wurde eine Datei von der gleichen Diskette geladen und
bei Aenderungen nicht vergroessert, so braucht bei Beibehal-
ten des urspruenglichen Dateibezeichners keine Sektoranzahl
angegeben werden. Die urspruengliche Dateı wird ueberschr ie-
ben. Somit koennen in Verbindung mit LOAD und DS schnell
Dataiaenderungen ausgefuehrt werden.

In allen anderen Faellen, wie z.B. bei Nutzung eines anderen
Dateibezeichners als der der geladenen Datei, muss die Anzahl
der auszugebenden 128-Byte-Sektoren angegeben werden. Dabei
kann auf die bei LOAD erfolgte Meldung ueber die eingelesene
Sektoranzahl zurueckgegr iffen werden.

Beispiel: AO=SAVE TEST.COM 5100 40 (ENTER)

Es ıst zu beachten, dass bei DIENST-Kommandos, die das
Nummernmenue nutzen, der Speicherinhalt verfaelscht werden
kann. So sollten zwischen einem LOAD-SAVE-Vorgang diese
Kommandos nicht genutzt werden. Anwendbar sind aber die
DIENST-Kommandos DUMP, DS, FILL, MOVE und SEARCH.

-9 -

DIENST 1520 (SCPX).

AO=LOG

Mit diesem Kommando koennen die aktuellen DIENST-Standardpara-
metereinstellungen angezeigt werden.

A

c If file exists? A-overlay, B-back up, C-ask, D-skip
3 colums

P_ ON paging

R (OFF) request Y/N on current file (O-request new Name)
v con) read after write

s (oN) show system Files

T (OFF) stop if disk Full

M (OFF) mark copied Files

x list drives A-P if on-line

u list users 0-15

“ cONd submit =mw.SUB - IENST

#/= 1-8 or (R)ead/urite, (S)ystem/dir, e(X)tra

DIENST O100h - taddr)

TPA taddr) - (addr) (number) sectors

Ao=

Das Beispiel zeigt die uebliche Ausgangseinstellung.

Die Aenderung der Parametereinstellung kann entweder direkt ın
der DIENST-Grundstellung oder als Abschluss einer DIENST-
Kommandoeingabe erfolgen.

Die anzugebenden Parameter muessen In eckigen Klammern LI
eingeschlossen warden (abschliessende Klammer kann weggelassen
worden). Die Angabe mehrerer Parameter gleichzeitig ist mong-
Lich.

Beispleli AO=LA) CENTER)
Die Spaltenanzahl bei Folgenden Bildschirman-
zeigen (z.B. Nummernmenue) ist 4. Der Ruack-
sprung nach Programmausfuchrung In die DIENST-
Grundstellung wird ein- oder ausgestellt (Zei-
chen "x").

Einige der Parameter haben Schalterchorakter, die ihre Wertung
bei Angabe des Parameters jeweils umschalten. Um sie wieder ın
den Grundzustand zu bringen, muessen sie ein zueites Mal
angegeben werden.

Andere Parameter sind nur fuer das augenblickliche Kommando

wirksan und schalten ihre Wirkung nach Ausfuehrung des
Kommandos selbst ab.

ie

DIENST 1520 (SCPX)

A,B,C,D,V,T,M,Q- sind COPY-Parameter (siche 3.4.4.)

3 = Spaltenanzahl (colums) bei Anzeige (moeg-
lich: 1-9)

P = Bildschirmanzeige (ON) - seitenueise

\ (OFF) - fortlaufend rollend

R com = Forderung einer (Y/N) - Antwort vor Ausfuch-
rung jeder Operation (siehe 3.4.4.)

Ss «om = Anzeige von Systemdateien wird nicht unter-
drueckt

x,U = sind DIR-Parameter (siehe 3.4.5.)

" cON) = nach Programmausfuchrung erfolgt Ruecksprung

in die DIENST-Grundstellung (A0=)
#/- = sind SET-Parameter (siehe 3.4.27.)

DIENST 0100H-(addr)
= zeigt den Speicherbereich an, der durch
DIENST.COM belegt ist

TPA  «addr)-(addr) number) sactors
= zeigt den Speicherbereich an, der fuer das
Erstellen und Bearbeiten von Dateien verfurg-
bar ist.

Die gewachlte Standardparametereinstellung kann geaandert
werden und mit SAVE als neue DIENST-Version abgespeichert
werden (dadurch neue Ausgangseinstellung).

3.4.12._MOVE
AO=MOVE (aadri) (eadri) (aadr2)

taadr1) - Anfangsadresse des Quellbereiches in HEX
(eadr1) - Endadresse des Quellbereiches in HEX
taadr2) - Anfangsadresse des Zielbereiches in HEX

Mit dem Kommando MOVE koennen zusammenhaengende Speicherberei-
che beliebig verschoben werden. Das Kommando transportiert den
Inhalt des Speicherbereiches mit der Anfangsadresse (aadri)
und Endadresse (eadri1) byteweise nach dem ab Adresse (aadr2)
beginnenden Bereich.

Beispiel: AO=MOVE 1000 IFFF 4000 (ENTER)

ne

DIENST 1520 (SCPX)

Die Endadresse des Zielbereiches wird nicht angegeben, da
DIENST aus der Anfangs- und Endadresse des Quellbereiches die
Laenge des zu verschiebenden Bereiches ermittelt. Der Inhalt
des Quellbereiches bleibt erhalten, vorausgesetzt, dass er
sich nicht mit dem Zielbereich ueberlappt (wenn die Ziel-
adresse niedriger ist als die Endeadresse des Quellbereiches).

Es ist zu beachten, dass MOVE keine Speichergrenzen erkennt,
so dass ein Ueberschreiten abgesichert werden muss.
3.4.18. _READ
1._Lesen_in_den_Spai:

AO=READ (track) (sector) (adr) (n sectors)

(track) - Spur auf Diskette, ab der gelesen wird in DEZ
(sector) - Sektor dieser Spur in DEZ
tadr) - Adresse, ab der das Gelesene Im Speicher abge-

legt werden soll
(n sectors) - Anzahl der zu lesenden Sektoren

Mit diesem Kommando koennen Daten einer Diskette (ab Spur/Sek-
tor) ab angegabener Adresse In den Speicher eingelesen werden.
In Verbindung mit den Kommandos DS und WRITE kann schnell eine
Datei geaendert werden. Wachrend dieses Vorganges sollten
keine Kommandos genutzt werden, die das Nummernmonue aufrufen,
da dadurch eingelesene Daten ueberschrieben werden koennen.
Fahlt die Angabe der Sektoranzahl (n sectors), so wird nur ein
Sektor eingelesen. Wird die Speicheradresse (adr) nicht
angegeben, wird der Sektor ab Adresse BOH eingelesen.

Beispiel:

AO=READ 10 4 XENTER)
AO=READ 10 4 1000 ENTER)
AO=READ 10 4 1000 8 (ENTER)

Waahrend Ausfuehrung des Kommandos erscheinen auf dem Bild-
schirm eine Reihe von Informationen:
G=XX:YY  T=XXK S=XRX PS=ARX  ALSKKXKKUXK

6 - XX - Blocknummer
YY - Ordnungsnummer im Block

T - Spurnummer

S_- Sektornummer

PS - physische Sektornummer

At - Speicherbereich, in den DIENST die eingelesenen
Sektoren eingelesen hat

-87 -

DIENST 1520 (SCPX)

AO=READ (track) (sector) (kind of read) (n sectors)

(track) - Spur, ab der gelesen wird in DEZ
(sector) - Sektor dieser Spur
(kind of read) - Anzeigeform:
x% - Anzeige in HEX und ASCII
XA - Anzeige in ASCII
XH - Anzeige in HEX
*% - gleiche Wirkung wie XX
(n sectors) - Anzahl der zu lesenden Sektoren
(soll nur ein Sektor gelesen werden, kann die
Angabe entfallen)

Mit diesem Kommando koennen Daten einer Diskette gelesen und
auf Bildschirm dargestellt werden. Es wird aber immer nur ain
Sektor in den Speicherbereich BOH-OFFH abgelegt. Analog wie
bei 3.4.18.1. werden auf dem Bildschirm zusaetzliche Infor-
mationen (6,T,5,PS,At - siehe 3.4.18.1.) angezeigt.

Bi

el: AO=READ 10 2 XX 3

3.4.19._URITE
AO=URITE (track) (sector) (adr) (n sactors)

track) - Spur, ab der geschrieben wird In DEZ

(sector) - Sektor dieser Spur In DEZ

tadr) - Speicheradresse, ab der die Daten ausgegeban
werden

in sectors) - Anzahl der auszugebenden Sektoren

Mit diesem Kommando koennen im Speicher ab Adresse (adr)
stehende Daten auf Diskette (Spur, Sektor) geschrieben werden.
In Verbindung mit don Kommandos READ und DS kann schnell eine
Datei geaendert werden. Uachrend dieses Vorganges sollten
keine Kommandos genutzt werden, die das Nummernmenue aufrufen,
da dadurch eingelesene Daten ueberschrieben werden koennen.
Fehlt die Angabe (n sectors), wird nur ein Sektor geschrieben.
Wird die Speicheradresse (adr) nicht angegeben, wird der Sak-
tor ab Adresse BOH (bis Adresse OFFH) ausgegeben.

Beispiel:
AO=URITE 10 4 ENTER)

AO=WRITE 10 4 1000 «ENTER»
AO=URITE 10 4 1000 8 CENTER»

Es ist zu beachten, dass der read only-Schutz bei WRITE nicht
wirkt und so jeglicher Bereich (auch der Systembereich
einschliesslich Verzeichnis) ueberschrieben werden kann. Falls
auf den Verzeichnisbereich geschrieben wird, muss ein AC fol-
gen.

Im Kommando ist auch eine Ziellaufwerksangabe moegl ich#

DIENST 1520 (SCPX)

AO=URITE B: 10 4 1000 8 (ENTER)

AO=READGR (start group) (adr) (n sectors)

(start group) - Blocknummer des zu lesenden Blockes in HEX

Kadr) - Adresse des Speichers, auf den Daten gelesen
werden

in sectors) - Anzahl der zu lesenden Sektoren des Blockes in
DEZ

Dieses Kommando dient zum Lesen von Datenbloecken in den Spei-
char. Vor dem Kommando kann mit dem DISK-Kommando (Ermittlung
der Groesse und Anzahl der Bloccke) und dem GROUP-Kommando
(Ermittlung der Verteilung der Bloecke einer Datei auf Disket-
te) der zu lesende Block analysiert werden. Nach dem Ein-
lesen sollten keine Kommandos genutzt werden, die das
Nummernmenue aufrufen, da dadurch eingelesene Daten
veberschriaben werden koennen.

Fohlt die Angabe der Sektoranzahl (n sectors), wird nur ein
Sektor eingelesen. Ist die Speicheradresse (adr) nicht angega-
ben, wird der Sektor ab Adresse 80H aingegeben.

Belsplel: AO=READGR 1 2000 4 <ENTER)

Waehrend der Ausfuchrung des Kommandos erscheinen auf dem
Bildschirm eine Reihe von Informationen (6,T,5,PS,At - siahe
3.4.18.1.).

3..Lasen_auf_Blldschir®
AO=READGR (start group) (kind of raad) (n sectors)

(start group) - Blocknummer des zu lesenden Blockes In HEX
kind of read) - Anzeigefor
xX - Anzeige in HEX und ASCII
xA - Anzeige in ASCII
XH - Anzeige in HEX
X - gleiche Wirkung wie XX
(n sectors) - Anzahl der zu lesenden Sektoren in DEZ

Mit diesam Kommando koennen Datenbloecke von der Diskette auf
den Bildschirm gebracht werden.

Sektorweise werden die Daten in den Spelicherbereich BOH-OFFH
uebernommen.

AO=READGR 2 XX 4

-8 -

DIENST 1520 (SCPX)

3:

AO=URITEGR (group) (adr) (n sectors)

(group) - Blocknumser des auf Diskette auszugebenden
Blockes in HEX

Kadr) - Speicheradresse, ab der die Ausgabe erfolgt

in sectors) - Anzahl der auszugebenden Sektoren In HEX

Das Kommando dient zum Schreiben von im Speicher stehender
Daten als Block auf die Diskette.

Blocknummer 0 ist fuer das Verzeichnis reserviert. Falls auf
den Verzeichnisbereich geschrieben wird, muss anschliessend
ein AC erfolgen. Fehlt die Angabe (n sectors), wird nur ein
Sektor ausgegeben. Wird die Speicheradresse (adr) nicht ange-
geben, wird der Sektor ab Adresse BOH ausgegeben (Achtung:
zwischen READGR und WRITEGR keine Diskettenverzeichn |akomman-
dos ausfuehren, da sonst der Bereich ab B0H varfaelscht
wird).

Das Kommando URITEGR beruecksichtigt keine Diskettenmerkmale
oder Verzeichnis-Dateischutzmerkmale. Deshalb grosse Vorsicht!

AO=WRITEGR 1 2000 4 CENTER)

324222. _REC!
AO=RECLAIN

LAIM

Das Kommando dient dazu, geloeschte Dateien wieder zugaangl ich
zu machen. Nach Aufruf von RECLAIN werden nacheinander die
Namen gelooschtor Dateien mit einor (Y/N)-Entscheidungsabfor-
derung angezeigt. Bei Eingabe von Y wird die antsprechande
Dateiverzeichniseintragung reganer iort-

Wurde der zu einer geloeschten Datei gehoerige Diskettenbe-
reich bereits mit anderen Daten uaberschrieben, erscheint bei
Anzeige des Dateinamens eine entsprechende Mitteilung. Die
Datei kann nicht regeneriert werden.

Wurde die Datei gefunden und reorganisiert, kann mit der
ESCAPE-Taste das weltere Anzeigen der zu loeschenden Datei
abgebrochen werden.

Ist die Dateibezeichnung der wiederherzustellenden Datei
bekannt, kann die Regener ierung in der Form:

AO=RECLAIM (dateibezeichnung) (ENTER)
erfolgen.

Jeder wiederhergestellten Datei wird aus Sicherheitsgruenden
R/O-Status zugewiesen.

x##.SUB-Dateien sollten niemals regeneriert werden (dies kann
zur Zerstoerung des Disketteninhaltes fuehren). Wird dies
versehentlich getan, ist ===.SUB in den Status R/W zu setzen
und mit ERA zu loeschen.

-9-

DIENST 1520 (SCPX)

Soll eine Datei wiederhergestellt werden, deren Dateibezeich-
nung bereits auf dieser Diskette existiert, muss die existie-
rende Datei umbenannt werden (Zuweisung eines neuen Namens
fuer eine geloeschte Datei innerhalb RECLAIM ist nicht moeg-
lich).

<ENTER>

„DOC  recover (Y/ND? (n) no

„DOC can not recover, file with bad extent
-DOK _ recover (Y/N)? (y) yes

AO=REN

Das Kommando dient zum Umbenennen von Dateien. Da mit dem
Nummernmenue gearbeitet wird, ist ein effektiveres Arbeiten
als mit dem rasidenten Kommando REN moeglich, wenn mehrere
Dateien umbenannt werden sollen.

Nach der Auswahl der umzubenennenden Datelen werden diese
einzeln angezeigt, und es erscheint Jeweils die Aufforderung
zur Eingabe des neuen Namens (neu NAME:). Nach dessen Eingabe
erscheint auf der Folgezeile noch einmal alter und neuer Name
(alter Name = neuer Name). Das Umbenennen erfolgt nur, wenn
die Diskette R/U-Status hat. Andernfalls erfolgt eine Fehlor-
mittellung, dass nicht auf die Diskette geschrieben werden
kann und es muss ein AC ausgefuehrt werden. Soll eine ausge-
waehlte Datei nicht umbenannt werden, wird bei der Aufforde-
rung zur Eingabe des neuen Namens (new NAME: nur (ENTER?
eingegeben.

Zur schnellen Eingabe der neuen Dateibezeichner kann fuer
gleichbleibende Teile das Sternsymbol (=) verwendet werden.

Belspigl:
AO=REN CENTER)
AO:1= DOK „DOK | pIP „DOK | 3= MIN „DOK

2
AO:4= LINK .DOK | 5:
WAEHLE? 1 3 (ENTER)

DIENST .MAC

AO:DOK „DOK=new NAME:NEU.TXT (ENTER)
AO=DOK „DOK=NEU „TAT

AO:MIN ew NAME:K#.® (ENTER)
AO:MIN IN „DOK

Ao=

Durch Eingabe des Parameters [RI ist eine zusaetzliche Namens-
aenderungsentscheidung (Y/N) nach Eingabe des neuen Namens
moeglich (Achtung: Wirkung bleibt auch bei darauffolgenden
Kommandos erhalten. Deshalb durch nochmalige Eingabe [LRJ dan
Standard zuruecksetzen).

-91-

DIENST 1520 (SCPX)

DOK „DOK | 2= PIP „COM | 3= MIN .DOK
WAEHLE? 1 (ENTER)

OK „DOK=new NAMEZNEU.TXT (ENTER)
:DOK „DOK=NEU .TXT rename (Y/ND? (y)
AO=CRI (ENTER)

AO=

Soll nur eine Datei umbenannt werden, kann unter Umgehung des
Nummernmenues der alte Name sofort eingegeben werden.

Beispi

AO=REN PIP.COM (ENTER)
A:PIP „COM=new NAME:PROG.COM (ENTER)
IP „COM=PROG „COM

Generell kann auch jedes andere Laufwerk direkt angesprochen
werden (z.B. AO=REN B:).

Es besteht die Moeglichkeit, durch Angabe eines Doppelpunktos
(2) nach new NAME:: bei der Eingabe des neuen Dateinamens
mehrere ausgewachite Dateien ohne Angabe jedes einzelnen neuen
Dateinamens umzubenennen, Indem beim neuen Dateibezeichner der
Teil, der unveraendert bleibt, mit ® angegeben wird.

AO=REN Bi (ENTER)

BO: 1= DOK .DOK | 2= PIP „COM | 3= MIN „DOK
BO: 4= LINK .DOK | 5= DIENST .COM

WAEHLE? 1 3 4

B:DOK „DOK=now NAME::m.TXT (ENTER)
DOK „DOK=DOK -TXT

MIN „DOK=MIN „TRT

B:LINK „DOK=LINK “TXT

A0=

Ausserdem soll auf die Moeglichkeit hingewiesen werden, dass
beim COPY-Kommando durch Angabe des LQJ-Parameters ein
sofortiges Umbenennen der Zieldatei bein Kopieren moeglich
ist, wodurch auch auf der gleichen Diskette ein Duplikat der
Datei erstellt werden kann.

AO=RESET

Dieses Kommando bietet die Moeglichkeit, eın einzelnes Lauf-
werk ohne Nutzung des Warmstartes (AC) wieder in den Grundzu-

-92 -

DIENST 1520 (SCPX)

stand zu setzen (on-line). Somit brauchen nicht wie beim
Warmstart alle Laufwerke in den Grundzustand gesetzt werden,
wodurch Zeit gewonnen wird und andere R/O-gesetzte Laufwerke
nicht ihren Schreibschutz verlieren.

Dieses Kommando kann auch genutzt werden, um Disketten mit
zerstoertem Verzeichnis zu retten.

Das aktuelle Laufwerk ist A. Es soll auf das Laufwerk B
kopiert werden, auf dem die Diskette geuechselt wurde
(nach dem letzten Warmstort). Um die R/O-Fehlermeldung
beim Kopierversuch zu vermeiden, muss vorher AO=RESET
(ENTER) eingegeben werden.

Das Kommando dient zum Abarbeiten und Laden von .COM-Dateien
(obarbeitbare Programmdateien). Es kann mit dem Nummernmenue
gearbeitet werden, das auf der Diskette alle vorhandenen
Dateien vom Typ .COM auswachlt und anzeigt. Nach Auswahl
einer Programmdatei und deren Abarbeitung wird wieder in den
DIENST-Grundzustand zurueckgekehrt.

AO=RUN Bi: (ENTER)
BO: 1= INIT COM |
Bo: 4= PIP „COM |
WAEHLE? 1 CENTER)

sYS6 „COM | 3= SEFR „com
S= STAT „COM

Das zu ladende und abzuarbeitende Programm kann sofort
angegeben werden. In diesem Fall kann als zusaetzlicher
Parameter der Nane jener Datei angegeben warden, die an-
schliessend durch das geladene Programm bearbeitet werden
soll.

Beisp.
Verwendung beim Textprozessor TP
AO=RUN TP B:TEST.DOK

Der Textprozessor wird vom Laufwerk A geladen, gestartet und
es kann die Textdatei TEST.DOK auf Laufwerk B bearbeitet

werden.

DIENST ist mit einem automatischen Restart ausge-
stattet (Submit-Funktion =(ONI/(OFF) siehe LOG).

Ist dieser eingeschaltet (=(0N)), wird vor dem Star-
ten des geuoehlten Rrogramms (RUN) eine Submit-Dateı
auf der Diskette geschaffen zur Kontrolle ueber SCP.

-93-

DIENST 1520 (SCPX)

Am Programmende wird diese Datei durch DIENST ge-
loescht und DIENST neu gestartet.

Somit muss gewaehrleistet sein, dass die Diskette,von
der man Programme mit RUN startet, nicht R/D gesetzt
sind, wenn der Parameter [x] (ON) ist.

Bei SCPX U 1.2 erscheint nach Ausfuehrung des mit RUN
gestarteten Programms die Fehlerlampe und undef inier-
te Zeichen auf dem Bildschirm. Durch Druecken einer
beliebigen Taste kann zum Programm DIENST zurueckga-
kehrt werden, wobei die Datei DIENST.COM zum erneuten
Einlesen im aktuellen Laufwerk zur Verfuegung stehen

muss.

3:42

AO=SEARCH (aadr) (eadr) (byte) Libyte)]

taadr) = Anfangsadresse des Speicherbereiches

Keadr) - Endadresse des Speicherbereiches

(byto) - Byte bzw. Bytefolge (Laenge 1 nis 128 Zeichen),
die Im angegebenen Speicherbereich gesucht uerdan
sollen

Das Kommando dient zum Suchen von Bytefolgen, dargestollt Im
HEX- oder ASCII-Code. Aufeinanderfolgende Zeichen im HEX-
Darstellung sind durch Leerzeichen voneinander zu trennen.
Aufeinanderfolgende Zeichen In ASCII-Darstalleung werden
lueckenlos hintereinander eingegeben und die Folge am Anfang
und am Ende mit “-Zeichen eingeschlossen. Die Eingabe einer
gemischten Codefolge (HEX und ASCIT) Ist moeglich, wobel zwi-
schen HEX- und ASCII-Zeichen ein Leerzeichen stehen muss.

Sind ainzelne Zeichen einer Bytefolge nicht bekannt bzu.
beliebig, so kann an den entsprechenden Stellen ein
Fragezeichen (?) stehen.

Nach Ausfushrung von SEARCH werden auf dem Bildschirm alle
erkannten Codefolgen mit Speicheradresse angezeigt. Wenn keine
antsprechende Codefolge Im angegebenen Speicherbereich
existiert, kehrt DIENST ohne zusaetzliche Meldung In den
Grundzustand zurueck (A0Q=).

Belspie

AO=SEARCH 1000 1400 "???WORT“ C3 ? ? CENTER)
1004 "SCPWORT” C3 64 70

12F4 "(0D)KOA)LOA) WORT“ C3 C4 78

1309 "NEUWORT” C3 64 70

AO=

Es ist zu beachten, dass bei Bytefolgen, die Befehle dar-
stellen, HEX-Speicheradressen so einzugeben sind, wie sie im
Speicher stehen (erst LOW-, dann HIGH-Teil).

-9 -

DIENST 1520 (SCPX)

Es sollen alle Sprungbefehle (unbedingte Absolutspruenge)
im Speicherbereich 81004 bis B000H gefunden werden, die
als Zieladresse #FOSH haben:

AO=SEARCH 6100 8000 C3 06 6F <ENTER)

AO=SET L-charakter] (ENTER) character - 1-8,5,R,X

character - Eingabe des Merkmals in den Dateinamen durch Anga-
be einer Ziffer von 1-8 (Setzen eines Bits an eine
der 8 Stellen des Dateinamens)
- Eingabe des Merkmals in den Dateityp durch Eingabe
von S - [SIystem (1.Byte des Dateityps = Byte 9 in
Verzeichniseintragung)
(Datei wird Systemdatei)
R - ERlead/Urite (2.Byte = Byte 10 in Ver-
zeichniseintragung)
(Datei wird R/U gesetzt)
X - ElxItra (3.Byte = Byte 11 In Verzeichnis-
eintragung)
“ohne weitere Bedeutung)

Diesas Kommando dient zum Setzen von Merkmalen an eine be-
stimmte Stelle der Dateibezeichnung in der Verzeichniseintra-
gung. Dies arlaubt die Gruppierung einer Anzahl von Dateien
durch Setzen des gleichen Merkmals in ihren Dateibezeichnern.

Die markierten Dateien werden angezeigt. Auch in der Var-
zelchniseintragung nicht genutzte Namansstellen koennen mar-
klart worden.

Beisele:!

AO=SET [-51 (ENTER)

AO: 1= TESTI „MAC | 2= DEMO „DOK | 3= TESTI COM
AO: 4= TESTI  .REL | 5= XXX „COM | 6= TESTI  .DOK
WAEHLE? 1 3 4 (ENTER)

AO: TESTI  .MAC
AO: TESTI  .COM
AO: TESTI  .REL
AO=

in_der_Dateibe:

zeichnung
SET kann in Verbindung mit dem Aufruf anderer DIENST-Kommandos

verwendet werden. Es wird das Kommando dann Jeweils nur mit
den Dateien ausgefuehrt, die das geforderte Merkmal besitzen.

-95 -

5 DIENST 1520 (SCPX)

oPy L+5I (ENTER?
TESTI  .MAC | 2= TESTI .COM | 3= TESTI | .REL

no:

WAEHLE?

Es werden im Nummernmenue nur die Dateien angezeigt, be
denen das 5. Merkmal gesetzt ist.

3.4.22.2._SET
AO=SET (ENTER?

Mit SET kann geprueft werden, welche Merkmale in den Dateibe-
zeichnungen gesetzt sind. Sollen alle Dateien geprueft werden,
ist sofort nach SET #®* oder bei WAEHLE? i- anzugeben.

Bi‘
AO=SET CENTER?

AO: 1= TESTI „MAC | 2= DEMO .DOK | 3= TESTI .COM
AO: 4= TESTI  .REL | 5= XXX „COM | = TEST .DOK
WAEHLE? 1-

TESTL „MAC
DEMO „DOK
TESTI „COM
TESTI  .REL
xxx „COM
TESTı  .DOK

324.2223._SEI_
AO=SET L+charakter)

chorakter - 1-8,5,R,X

Dieses Kommando dient zum Loeschen der mit SET L[-character]
gesetzten Merkmale. Es werden im Nummernmenue nur die Dateien
angezeigt, die das angegebene Merkmal besitzen. Nach Datelaus-
wahl wird das Merkmal der gewachlten Dateien zuruackgesetzt.

Beispiel:
AO=SET L+5J (ENTER)
AO: 1= TESTI  .MAC | 2= TESTI „COM | 3= TESTI .REL
WAEHLE? 1-
ao: TESTI  .MAC

AO: TESTI COM
a0: TESTI -REL
Ao=

Achtung: SET arbeitet ohne Ruecksicht auf R/0O-Status (deshalb
vorher AC ausfuehren!).

-9 -

DIENST 1520 (SCPX)

AO=SETSYS
AO=SETDIR

SETSYS wandelt Dateien zu Systemdateien, die bei “Aus” (OFF)
gestellten LSJ-Schalter (siehe LOG-Kommando) nicht beim DIR”
Kommando bzw. im Nummernmenue mit aufgelistet werden. Die SYS-
Dateien koennen mit dem SORT-Kommando (als zuletzt angezeigte
Dateien) sichtbar gemacht werden (oder [SI-Schalter auf ON)
setzen --> dann sind diese Dateien durch runde Klammern um den
Dateibezeichner speziell markiert).

SETDIR setzt das Systemmerkmal wieder zurueck.
Beide Kommandos koennen mit dem Nummernmenue (nur wandelbare

Dateien werden angezeigt) oder mit speziellen Dateibezeichnun-
gen arbeiten.

Achtung: Es wird ke Ruecksicht auf R/O-Status genommen
(deshalb vorher AC ausfuehren!).

3.4.22._SEIRO/SETUR

AO=SETRO
AO=SETUR

SETRO setzt Dateien zum Schutz R/O, so dass sie nicht be-
schrieben werden koennen.

SETUR macht Dateien wieder beschreibbar.

Beide Kommandos koennen mit dem Nummernmenue arbeiten. Bei
Eingabe von SETRO werden nur die Dateien angezeigt, die R/U
gesetzt sind. Bei Eingabe von SETUR werden nur die Dateien an-
gezeigt, die R/O gesetzt sind.

Die Datei "“w.SUB darf niemals R/O gesetzt werden, da dies zur
Zerstoerung des Disketteninhaltes fuchren kann.

Achtung: Es wird keine Ruecksicht auf R/0-Status genommen
“deshalb vorher AC ausfuehren!).

3.4.30._SORT
AO=SORT n (n=1-4)

Dieses Kommando sortiert die Verzeichniseintragungen alphabe-
tisch nach 4 verschiedenen Moegl ichkeiten!

- Sortieren nach Dateinamen

- Sortieren nach Dateinamen mıt Systemdateien zuletzt
- Sortieren nach Dateityp

- Sortieren nach Dateityp mit Systemdateien zuletzt

-97 -

DIENST 1520 (SCPX)

Bei darauffolgender Eingabe von DIR oder eines Kommandos, das
das Nummernmenue benutzt, werden die Dateien dann entsprechend
sortiert angezeigt.

Mit SORT 8 erhaelt man ein unsortiertes Verzeichnis. Auf der
Diskette bleiben generell die unsortierten Verzeichniseintra-
gungen erhalten.
Um die benutzte SORT-Einstellung permanent als Standard zu
setzen, muss das modifizierte DIENST mit SAVE auf Diskette
gerettet werden.

3.4.31.
AO=SIZE

ZE

SIZE liefert eine Statistik ueber die Groesse von Diskettenda-
teien. Das Kommando kann mit dem Nummernmenue oder mit spe-
ziellen Dateinamen arbeiten.

Nach Auswahl der entsprechenden Dateien werden die Anzahl der
Sektoren, die die Datei belegt, die Anzahl der freion Saktoran
Im lotzten dazugehoerigen Block, die Groesse der Datal In
KByte und die laufende Summe der Groessen der gewanhitan
Dateien in KByt» or ozeigt.

Beisele)
AO-SIZE (hu ic
1= IFsrı = 1 2= TESTI MAC | 3= DEMO .DOK
AO“ A= TAN „uor
WAFHLE? 1 4 CENTER
TOTı 00m 46 ectors 15 empty 6K  &K
TKO Wu 53 vsctors 12 ampty 55K 61K

AO=

Das Kommando kann sr 17t werd.n, um vor dem Kommando COPY die

Groesse der zu kopıorını .. Datei zu bestimmen (bei verschlo-
denen Diskettenformaten dabei auch das DISK-Kommando mit ha-
ranziohen).

SPEED ist ein Kommando, mit dem man die Geschwindigkeit der
Bildschirmanzeige fuer alle Ausgaben des DIENST-Programmes
einstellen kann.

m=0 - schnellste Geschwindigkeit

I
j
m=9 - langsamste Geschwindigkeit

-98 -

DIENST 1520 (SCPX)

Das Kommando ist besonders beim Auflisten grosser Verzeich-
nisse sowie bei den Kommandos TYPE und DUMP nuetzlich.

Waehrend der Ausgabe (durch diese Kommandos) kann durch Ein-
gabe von O0 bis 9 die Geschwindigkeit geaendert werden (ohne
Rueckkehr in Grundzustand [AO=J und Eingabe SPEED).

Durch vorherige Eingabe von AP kann die Ausgabe auch
gleichzeitig auf den Drucker erfolgen (ohne dass Seitenvor-
schub erzeugt werden kann).

3.4.33._SIAI
AO=STAT

STAT liefert statistische Informationen ueber alle Laufwerke,
die on-line sind. Es werden der R/O- oder R/U-Status dar
Diskette, der belegte Speicherbereich in KByte, der freie
Speicherbereich In KByte und die Diskettenkapazitaet In KByte
angezeigt. Bei Angabe R/O wurde die Diskette gewechselt, und
die gelieferten Informationen gelten nach fuer die ursprueng“
liche Diskette (deshalb vorher AC eingeben!).

Balseie
AO=STAT B: (ENTER)

AO: R/U Used: 292K Frae: 2K Capacity: 296K
BO: R/O Used: 170K Free: 124K Tapazity: 296K
Die Differenz aus (Gesamtkapazitoct “eier Spei-

cherbereich)) ergibt sich daraus, dass Ins Dateiver-

zeichnis genutzt wird.

3.4.34._IESI
AO=TEST

Dieses Kommando dient zum Test von Distcitun aut det.“to Sak-
toren. Es werden alle Spuren der Diskot' ,.) | au. ueber-
pruaft, ausser den Systemspuren.

Nach Kommandoaufruf erscheint die Ausschrift “- disk chacksum
& test”. Anschliessend wird getestet. Ist dıe Diskette frei
von defekten Sektoren, erscheinen durchweg die Zeichen
wann... und am Ende wird die Kontrollsumme der gelesenen
Bytes In HEX angezeigt (checksum ___.

Wird ein defekter Sektor auf der Diskette erkannt, so werden
Blocknummer, Spurnummer, Sektornummer und physische Sektorpn-
sition dieses Sektors sofort angezeigt. Danach wird die Dis-
katte weiter getestet. Ist der Test beendet, fragt
"Save?(Y/N)”, ob man die Dateien, die defekte Sektoren enthal-
ten, retten will. Bei Antwort mit "Yes" werden die defekten
Sektoren in einer speziellen Datei ısoliert (DIENST schafft

-9-

DIENST 1520 (SCPX)

eine bad-sector-Datei im Verzeichnis, in der die defekten
Sektoren untergebracht werden). Anschliessend wird gefragt,
ob man die Dateien sehen will, die die defekten Sektoren
enthalten:

"Show bad files (Y/N)”. Beı Eingabe “Yes” werden diese Dateien
angezeigt. Die Datei = == ist die bad-sector-Datei.
Danach erscheint die Frage "Repair (Y/N)” zur Ausbesserung
defekter Sektoren. Bei "Yes" werden die Dateien neu eingelesen
und wenn die defekten Sektoren noch lesbar waren, werden sie
mit ihrem Inhalt neu ueberschrieben. Sind sie nicht mehr
lesbar, werden sie mit 0 aufgefuellt. Nies fuchrt zum Verlust
von 128 Byte/Sektor, jedoch werden die Dateien wieder in
lesbarer Form verfuegbar.

Als letzte Information wird die Kontrollsumme (total checksum)
fuer die Diskette angezeigt.

eielz

AO=TEST (ENTER) - disk checksum & test
HERE
BLSLLLELELTETTETETEITESTIZESSFITSTEITPOPEPEUPETTETTTTTET
- bad sector on G=003:02 T=003 S=034 PS=120

DEREN
DELLEEELTEITETSITETETESSTETPETDTTUTITTETLILTLITLTETTTEIT
BERLELLLLELLETETEETEITEZSIZTETSTTETPETTUUUTITTEOTEITTT
BELELEELITZTEITETEISIIPSSSITUTTPPETETTELTITLTLTTLTLITLIT
BELLELELTEZEITEIZZESZTEITETEITUTPTTETTEPETTEUPETTELTTITT
BELELLELTIZEETTETEIZEITETSIESZESTETTETUUPPUTUTTLLTEELLIT
BELLELEITTZTEITESSTTISESTSETETTTPITPETEUTEOTETTTLTTETLTT
ELLELTEITEITIITUPTT Tee

Save?(Y/N) Yes

Show bad files(Y/N’Yes

TKO „DOC

Repair (Y/N? Yes

6=003:02 T=003 S=034 PS=120
Disk has 1 bad block
checksum S4BB

Die Testfunktion wird sofort abgebrochen, wenn sich das Ver-
zeichnis auf defekten Sektoren befinden.

IESI_S

AO=TEST S

Dieses Kommando ist identisch mit dem Kommando TEST, jedoch
werden zusaetzlich die Systemspuren getestet.

- 100 -

DIENST 1520 (SCPX)

Dieses Kommando dient zur Ausgabe von Dateien auf den Rild-
schirm. Mit TYPEH koennen auch .COM Dateien in HEX-Code
sichtbar gemacht werden. Es kann mit dem Nummernmenue gearbei-
tet werden. Fuer die Bildschirmausgabe sind die Bildschirm
steuertasten (siehe 3.3.1.) zu beachten.

TYPE - nur fuer Textdateien, die mit TP erfasst wurden an-
wendbar
- wenn Text formatiert ist, dann erfolgt fFormatierte
Ausgabe
TYPEA - zeigt Textdateien in numerierten Zeilen zu je 16

Zeichen sektorweise an (unformatiert) mit Angabe von
Blocknummer,  Spurnummer, Sektornummer, physischer
Sektornummer, Dateinome und Extentnummer

TYPEH - zeigt Dateien ın HEX-Code in numerierten Zeilen
zu je 16 Byte sektorweise an (auch „COM-Dateien); mit
Angaben wie TYPEA

TYPEX - zeigt Dateien In HEX- und ASCIHI-Code jeweils in

numerierten Zeilen zu je 16 Byte sektorweise an
(auch „COM-Dat.); mit Angaben wie TYPEA

Nach der Datelauswahl werden die gewuenschten Datelen ausge-
geben. Wird eine Datelserie hintereinander ausgegeben, so Ist
AK alnzugeben, wenn die Ausgabe der aktuellen Dateı abgebro-
chen und zur naechsten uebergegangen werden soll.

324.36. _USER

AO=USER | I-Benutzerbereichsnummer (0-15)
Mit diosem Kommando wird der Benutzerbereich fuer die Diskat-
tendatelen fuer laufende Operationen festgelegt.

Als Parameter | kann cine Zahl von 0-15 eingegeben werden. Der
User O0 ist nach dem Kaltstart eingestellt.

Beispiel:

AO=USER 1 (ENTER)
Als

3.4.32._XUSER
AO=XUSER j j-Zielnutzerbereichsnummer
Dieses Kommando erlaubt dıe Uebertragung von Dateien von einem

Nutzerbereich zu einem anderen (auf derselben kette oder
auf ein anderes Laufwerk --) unter Nutzung von COPY).

- 101 -

DIENST 1520 (SCPX)

AI=XUSER 2 (ENTER)

A1-2=COPY (ENTER)

Alz 1= TESTI .MAC | 2= TKO „DOC | 3= TESTI .CoM
Al: A= XXX „DOK | 5= TEST „REL

WAEHLE? 2 4

destination drive:a
“C)opy or (Move:c
TKO „DOC
xxx „DOK

(Copy - Dateien werden auf neuen Nutzerbereich kopiert,
bleiben jedoch im alten erhalten
- auf derselben Diskette, aber auch auf eine andare
moeglich

-)(Move

\

Datelen werden auf neuen Nutzerbereich verschoben,
bleiben im alten nicht erhalten
= nur auf derseiben Diskette moeglich

Jede kopierte bzw. ‚cr chobene Datei wird angezeigt.
Durch Eingabe won USER D kommt man wieder In den Nutzerba-
reich 0 (A0=).

3.4.38._UR

AO=URn (ENTER) 81

Dieses Kommanco r'nnı ‚prucnge zu speziellen Routinen, die
der Anwender zu NEN susaetzlich uuenscht, Sind diasa In-
stalliort, so kocn. 10 aufgerufen werden.

Der Speicherbereich fuer die Eingobe der Anwenderspruenge be-
ginnt bei der DIENST-Startadresse +40H (d.h. 140W).

Fuer jedes UR-Kommando werden 8 Bytes bereitgestellt.

URi -> 0140 - 0147H
UR2 -> 0148 - O14FH
UR3 -) 0150 - 0157H
UR4 -) 0158 - O15FH

Die Sprungadressen koennen mit DS eingegeben werden.
Anschliessend muss die neue DIENST-Version mit SAVE gerettet
werden.

Standardmaessig ist URI ein RETURN (69).

- 102 -

DIENST 1520 (SCPX)

1. Can not LOG on Bi
Diek log ERROR on B:, logging to drive only

1. keine Diskette in Laufwerk B eingelegt

2. Laufwerk B nicht korrekt verriegelt

3. defekte Diskette oder Diskettenverzeichnis in Lauf-
werk B

4. beim Schreiben - Zieldiskette in Laufwerk B kann phy-
sIschen Schreibschutz haben

5. beim Lesen - kein AC betaetigt, na “ = Diskette in
Laufwerk B eingelegt wurde

6. beim Kopieren von einer schreibge- Ni
mit dem Parameter [MItON) konn DIENST dı
nicht setzen

Yun Tiskett
marwen 0,0

7. beim Starten des Programms auf einer sinr«.Lyuschuetz-
ten Diskette mit RUN mit dem Parameter | Dmit (ON)
(automatischer Restart) konn die Submit Funktion keine
Datei DIENST.==“ auf der Diskette anlc,cı

2. Incorrect, for list of commands enter: ?

1. Eingabe eines falschen DIENS] -hommandos

2. Kommandoe ingabefehler

3. Eingabe von RUN vor Angabe e ns st whrbaren Datel-
namans vergessen

3. Invalid argument

1. notwendige  Adresseingabe oder andere Informat Ionen
nach einem Kommando vergessen

2. Eingabe eines Buchstabens statt einer geforderten Zahl

3. bei SAVE-Kommando Dateilaengenangabe vergessen

4. Press AC, disk Is =R/O on B:

1. Diskettenwechsel ohne nachfolgende Eingabe von NL --)
Versuch, auf schreibgeschuetzte Diskette bzw. Laufwerk
zu schreiben

5. Disk full on Bi

1. beim Kopieren einer Dateiserie kann die Zieldiskette
restlos gefuellt werden
- wenn Parameter LTIKOFF) versucht COPY, den Rest der
Serie noch auf die Zieldiskette in Laufwerk B zu
bekommen, so dass mon bei zu grossen Dateien die
Fehlermeldung erhaelt, jedoch danach eine kleinere
Datei, die noch darauf passt, kopiert werden kann
- es kann eine Serie dieser Fehlermeldungen erscheinen

- 103 -

DIENST 1520 (SCPX)
2. Beziehen des RUN-Komsandos auf eine gefuellte Disket-
te, wenn Parameter [xIsubmit (DN) ist, so dass keine
Submit-Datei geschaffen werden kann
&. file is xR/O

Versuch der Aenderung des Dateiinhaltes einer R/D-ge-
setzten Datei (erst mit SETUR R/U setzen)

7. Bad sector on READ Abort (Y/N)
wenn defekter Sektor auf Datei beim Lesen erkannt wird
8. No files ?7227922.727? on Ar

1. Frage nach Verzeichnis auf einer leeren Diskette
2. Frage nach Verzeichnis In einem leeren USER-Feld

9. Bad name

bei Benennen oder Umbenennen von Dateien mit Bezeich-
nungen, die nicht von SCP akzeptiert werden.

Dorueberhinaus existieren bei einzelnen Kommandos noch
spezielle Fehlermeldungen, deren Bedeutung aus dem angezeigten
Fehlertext hervorgeht.

Beispiele:

Out of disk bounds z.B. bei Kommando READER ist
eine zu grosse Blocknummer
angegeben, die die Disketten-
grenze ueberschreitet.

Can nat copy to same drive Es wird versucht, ohne neue Na-
menszuueisung “Parameter 0
Fehlt) auf die Quelldiskette In
den selben Nutzerbereich zu
kopieren.

- 104 -

CONF 1520 (SCPX)
ER Komprimieren einer SCP-Datei CONP 1520 (SCPX)

Einleitung

Das Programm CONF st n Zusammenhang mit dem Trans-
fornationsprogramm UTR1, -Anvenderdateien unter Regie
des Betriebssystems SCP in SCP-Dateien transformiert.

Bei dieser Transformation werden äfc Tnetze bzw. Dloecke der
SIOS-Datei stets mit einem neuen SCP-Satz begonnen. Dabei
koennen im Anschluss an den tatsaechlichen logischen Satz auf
der Diskette unbenutzte und nit dem Zeichen 00h aufgefuellte
Bereiche entstehen, die bei der Weiterverarbeitung dem voran-
gehenden oder dem nachfolgenden Satz zugerechnet werden.

Da diese Bereiche nicht benoetist werden und in bestimmten
Faellen bei der weiteren Dateibearbeitung stoeren, wurde das
Progrann COHF geschaffen. Es konpriniert die Datei und besei-
tigt so derartige bei der Dateitransformation entstandene
Bereiche durch lueckenloses Ketten der Saetze (Vergl. auch die
Beschrei bung zu UTR1 1520 (SCPX)).

Erlaeuterung

- Programmaufruf durch Tastatur
eingabe;

- Abschliessen der Eingabe mit

<EITER>;

- Programm CONF.CON wird damit
geladen;

CONF 1520 (SCPX) V x.x - Neldeausschrift des Programmes

- x.x steht fuer die Versions-
nummer;

Anzeige/Ausschrift

- Eingabeaufforderung fuer das
Laufwerk, in dem sich die mit
UTR1 erzeugte Quelldatei be-
findet.

SOURCE DRIVE NAN

- Eingabe ist nicht mit <ENTER>
abzuschliessen.

- Laufwerksname wird auf Guel-
tigkeit geprueft;

-165 -

CoNF 1520 (SCPX)

Anzeige/Ausschrift Brlaeuterung

- Wurde nur <BENTBER> bedient,
wird das Program mit Warn-
start verlassen.

SOURCB FILE NAME: _

Eingabe des Namens der Quell-
datei;

- Bingabe ist mit <ENTER> abzu-
schliessen, wenn weniger ala
17 Zeichen eingegeben werden.

- Wurde nur <ENTER> bedient,
verzweigt das Programm zur
Eingabe des Quellaufwerkana-

mens.
4.4. Bing der Parameter der Zieldi
Anzeige/Ausschrift Brlaeuterung

DESTINATION DRIVE NAMB

Bingabe des Laufwerksnamens,
in dem eich die Zieldatei be-
findet;

Quellaufwerk und Ziellaufwerk
duerfen gleich sein, wobei
durch haeufige Kopfpositionie-
rung erhoehte Programmlauf-
zeiten auftreten

- Eingabe ist nicht mit <ENTER>
abzuschliessen.

Der Laufwerksname wird auf
Gueltigkeit geprueft.

- Wurde nur <BNTER> bedient,
verzweigt das Programm zur
erneuten Eingabe des Namens
der Quelldatei.

DESTINATION FILE NAHE:_ - Eingabe des Namens der zu er-
zeugenden Datei;

Vorangestellte Laufwerksbe-
zeichnung sollte nicht angege-
ben werden. Wenn sie eingege-
ben war, wird sie ignoriert.

- 106 -

Anzeige/Ausschrift

CONF 1520 (SCPX)

Brlaeuterung

FILE EXISTS ALKBADY
DELETE FILE (Y OR N):_

RECORDS END IN 0D/OA
(r oR N:_

Es wird geprueft, ob die Datei
schon vorhanden ist. Existiert
sie noch nicht, wird sie ange-
legt.

Datei soll geloescht werden;

Anzeige erscheint, wenn eine
Datei mit diesem Namen bereits
existiert.

Bingabe <ENTER> oder "N":
Eingabe verhindert das Loe-
schen. Das Programm fordert
einen neuen Dateinanen.

Eingabe "Y":
Die Datei wird geloescht und
neu angelegt.

Vor Eingabe des Laufwerks kann
die Diskette gewechselt wer-
den.

Ziellaufwerk und Quellaufwerk
äuerfen gleich sein.

Satzendekennung;

Durch UTR1 ist es moeglich,
die uebertragenen SIOS-Saetze
mit einer _Satzendekennung
0D/OA abzuschliessen oder auch
nur in der logischen Satzlaen-
ge zu uebernehmen.

Bingabe "Y":
Die Dateikomprimierung wird
Gestartet.

Wurde nur <ENTER> bedient,
wird die Bingabeaufforderung
fuer den Namen der zu erzeu-
genden Datei wiederholt.

Zingabe
Es wird zur naechsten Eingabe
auf gefordert.

- 107 -

CONF 1520 (SCPX)

Bei urspruenglich geblockten Dateien, die bei UTRI in der
Vebertragungsart 2 behandelt wurden, ist an dieser Stelle auch
"N" einzugeben.

Wurde "N" eingegeben obwohl eine Endekennung vorhanden ist,
bleibt diese unberuecksichtigt und erhalten. Durch die Angabe
der logischen Satzlaenge bei der nachfolgenden Bingabe wird
die bestehende Satzstruktur durch das Einfuegen einer weiteren
Endekennung verfaelscht.

Anzeige/Ausachrift Brlaeuterung

LOGICAL RECORD LENGTH

(USE BLOCK LENGTH FOR

BLOCKED FILES): _ - Bingabe der logischen Satz-
laenge;

- Bei weniger’ als 5 Stellen ist
mit <ENTER> zu beenden.

Handelt es sich bei der Quelldatei um eine urspruenglich
geblockte SIOS-Datei und wurde bei UTR! mit der Vebertragungs-
art 0 gearbeitet (Blockung blieb erhalten ohne Binfuegen einer
atzendekennung), dann ist an dieser Stelle die logische
Blocklaenge einzugeben. Die Blockung bleibt dann auch in der
komprimierten (Ziel-) Datei erhalten. Die Bloecoke werden durch
die Satzendekennung (hier in der Funktion als Blockendeken-
nung) getrennt und lueckenlos gekettet.

Handelt es sich bei der Quelldatei um eine urspruenglich
geblookte SIOS-Datei und wurde bei UTR! mit der Uebertra-
gungsart 2 gearbeitet (Blockung blieb erhalten, jeder Satz im
Block mit Satzendekennung), dann ist durch die folgende Be-
rechnung die einzugebende Blocklaenge zu ermitteln:

BL=nx SL + 2(n-1)

BL = Blooklaenge fuer Eingabe
n = Anzahl der Saetze im Block
SL = logische Satzlaenge

Durch das Komprimieren wird auch in diesem Falle ein luecken-
loses Ketten aller Saetze erreicht, wobei die Blockung aufge-
hoben wird.

Soll aus einer geblockten SIOS-Datei eine gekettete SCP-Datei
unter Verwendung von UTR1 und CONF erzeugt werden, dann ist
bei UTRI die Transformationsart 1 als die guenstigste Variante
zu waehlen (es wird entblockt) und bei CONF mit der logischen
Satzlaenge zu arbeiten.

Anzeige/Ausschrift Erlaeuterung
CONDENSE IS WORKING - Das Komprimieren ist gestar-
tet (Arbeitszustand);

- 108 -

CONF 1520 (SCPX)

Anze sge/Aus schrift

ird die Diakette waehrend der
Arbeit voll, wird durch eine
Neldung inforziert.

4 Programmende

Anzeige/Ausschrift Erlaeuterung

END OF FILE - Ende der Datei ist erreicht;

- Anschliessend erfolgt erneuter
Sprung zum, „ Progranmanfang
(entsprechend Pkt.

- Die erzeugte komprimierte Du-
tei wird zueaetzlich mit der
Dateiendekennung 1A versehen.
Dadurch wird auch ein Ausli-
ster mit dem Kommando TYPE
moeglich.

Tmvarın DRIVE NANE - Unzulaeasiger Laufwerksname;

INWALID FILE NANE - Datei dieses Namens ist auf
der Diskette als Quelldatei
nicht vorhanden.

- Es ist ein anderer Name einzu-
geben.

DIRECTORY IS FULL - Das Verzeichnis der Diskette
ist voll. Es besteht die Iloeg-
lichkeit, die Diskette zu
wechseln, bevor der Laufwerks-
name eingegeben wird.

R/W FRROR - Schreibfehler beim Anlegen der
Datei. Es kann vor Eingabe des
Laufwerksnamens die Diskette
gewechselt werden.

=109-=

CONP 1520 (SCPX)
Anzeige/Ausschrift Brlaenterung

ERROR WRITB DISK - Fehler bein Schreiben der

Zieldatei (oder gefuellte
Zielkette). Das Programm

bricht mit Bearbeitung ab und
kehrt zum Anfang zurueck.

Fehler beim Lesen der Quelldatei werden direkt vom Betriebs-
system gemeldet. Es werden nach Bedienen von <ENTER>
Lesewiederholungen durchgefuehrt. Bleibt das erfolglos, bricht
das Programm die Bearbeitung ab und kehrt zum System zurueck.

ITrans-| Verarbei-|
SI0S- |forna-| SCP-Datei tungsart | SCP-Datei
Datei |tions-| I

lart | I

I}
| I lin logi- ISaetze werden ge

0) Ineuer SIOS-Satz = Ischer Ikettet und erhal
ee) Ineuer SCP-Satz ISatzlaen-|ten eine Satzen-
1 | Ige „|dekennung
8 =
blockt Ineuer SIOS-Satz = Imit Ende-ISaetze werden ge
2 Ineuer SCP-Satz mit Ikennung Ikettet,EK bleibt
|Endekennung OD/OA I lornalten
I
I

|Bloskung | bleibt er- lin logi- |Bloeoke werden
0 _ Ihalten; neuer SIOS- lecher Igekettet und er-

|Block = neuer SCP- |BL=(n SL)|halten eine EK
[| I

|
I
!
|
I
| \a K)
|
|
I
| ISatz

BESHNERER, (s>==

ge- ; neuer lin logi- |Saetze werden

blockt| 1 ISIOS-Satz = neuer Ischer SL Igekettet mit EK
I ISCP-Sat I I

-|---

|Blockung bleibt

!

| eue lin log. |Saetze werden ge
Il 2 ISIOS-Block = neuer ISL/BL Ikettet;EK bleibt
| ISCP-Satz; jeder Satz |BL=nSl+ lerhalten

j Imit EK I(n-1) j

- 110 -

CONF 1520 (SCPX)
Beachte: .

Soll aus einer geblockten SIOS-Datei eine gekettete SCP-Datei
durch die Programme UTR1 und CONF erzeugt werden, dann ist bei
UTR1 die Transformationsart 1 als die guenstigste Variante zu
waehlen, obwohl die durch UTR1 erzeugte Datei unter Umstaen-
den groesser ist.

Die Weiterverarbeitung durch CONF ist mit den dargestellten
Ergebnissen fuer alle UTR1-Transformationsarten moeglich.

PRIF 1520 (SCPX)

5 Drucken einer SC jatei PRIP 1520 (SCPX)
5.1. Binleitung

Durch das Programm wird die Hoeglichkeit geschaffen, beliebige
SCP-Dateien auf Papier unterschiedlicher Breite zu drucken.

5.2. Aufrufen des Programmes

Anseige/Ausschrift Brlaeuterung
PRIF - Progranmaufruf durch Tastatur-
eingabe;

'

Abschliessen der Eingabe mit
<ENTER>;

- Programm PRIF.COM wird damit
geladen;

PRIF 1520 (SCPX) V x.x

Meldeausschrift des Programmes

-x,x steht fuer die Versions-
nummer;

Bingeben der

Anzeige/Aussc hrift

DRIVE NAMB:_ - Bingabeaufforderung fuer das
Laufwerk, in dem sich die Dis-
kette mit der zu druckenden
Datei befindet.

— Eingabe ist nicht mit <ENTER>
abzuschliessen;

- Laufwerksname wird auf Guel-
tigkeit geprueft.

- Wurde nur <ENTER> bedient,
wird das Programm mit Warm-
start verlassen.

FILE NAME: _ - Bingabe des Dateinamens (ein-

deutiger SCP- gueltiger Datei-
bezeichner)

\

Eingabe ist mit <ENTER> abzu-
schliessen, wenn weniger als
17 Zeichen eingeben werden.

- 112 -

PRIP 1520 (SCPX)

Anzeige/Ausschrift Erlaeuterung

- Wurde nur <BENTER> bedient,
verzweigt das Programm zur
Laufwerkseingabe.

- Der eingegebene Dateiname wird
auf der Diskette gesucht.
Wird die Datei nicht gefunden,
erscheint eine Fehlermeldung
und erneute Bingabe des Datei-
namens wird gefordert.

RECORDS END IN OD/OA
TOR N:_ - Satzendekennung;

= Bingabe "Y":
Der Drucklistenparameter wird
fortgesetzt.

Wurde nur <BNTER> bedient,
kehrt das Programm zur Eingabe
des Dateinameng zurueck.

- Bingabe "N":
Aufforderung zur naechsten
Bingabe;

LOGICAL RECORD LENGTH:_ - Binga der logischen Satz-

laenge;

- Bei weniger a} Stellen ist
mit <ENTBR> zu beenden;

Wurde nur <ENTER> bedient,
kehrt das Programm zur Frage
nach der Satzendekennung zu-
rueck.

PRINT RECORDS IN FULL LENGTH

WORM:_ - Be besteht die Moeglichkeit,
die Saetze der Datei nur anzu-
arucken (Teilduck).

- Wurde nur <ENTER> bedient,
kehrt das Programm zur Frage
nach der Satzendekennung zu-
rueck.

= 4:2

Anzeige/Ausschrift

RECORD LENGTH TO PRIN

LENGTH OP PRINTING LINE
(CHARACTERS/LINB):_

1 LIND DISTANCR BEÜNEEI
PRINTED RBCORDS (Y OR 1):_

PRINTER READY?

Cr or mL

PRIF 1520 (SCPX)

Brlaeuterung

- Eingabe "Y'
Eingabe der Druckzeilenlaenge;

- Bingabe "N":
Durch die Anzeige:

wird zur Eingabe der zu druk-
kenden Satzlaenge aufgefordert.

- Werden weniger als 5 Ziffern-
stellen eingegeben, ist Bin-
gabe mit <ENTER> abzuschlies-
sen.

- Yurde nur <ENTER> bedient,

vird dic vorherige Frage wie-
derholt.

Eingabe der Druckzeilenlaenge;

Es ist ein lert im Bereich
1...128 einzugeben.

- Bei Eingabe von 0 oder >128
wird diese Frage wiederholt.

- lurde nur <EITER> bedient,
kehrt das Programm zur vorhe-
rigen Lingabe zurueck.

- Es besteht die lioeglichkeit,
nach jedem gedruckten Satz
eine Leerzeile einzufuegen.

Wurde nur <ENTER> bedient,
kehrt das Programm zur vorhe-
rigen Eingabe zurueck.

- Druckerbereitschaft;

Wiederholung der Frage;

zur <ENDEN> bedient,
das Program zur vorhe-
Tingabe zurueck.

PRIF 1520 (SCPX)

Anzeige/Ausschrift Brlaeuterung

- Bingabe :
Kontrolle der Druckerbereit-
schaft;

- Ist der Drucker nicht bereit,
erfolgt gef. eine Pehlermel-
dung in Abhaengigkeit vom
Druckertyp und von dessen
Binschaltzustand und die Frage
nach der Druckerbereitschaft
wird erneut gestellt.

PRINT IS WORKING - Druck ist gestartet (Arbeits-
zustand);

Auf dem Papier wird zunaechst ein Kopftext folgenden Inhalts
gedruckt:

PRIP 1520 (SCPX) V x.x

PRINT SCP PILE FILE NAME :nnnnnannnnnnnnnnn

wobei x.x fuer die Versionsnummer von PRIP steht und nnn...nnn
den Dateinamen darstellt.

Anschliessend wird die Datei in der gewuenschten Form voll-
staendig gedruckt.

5.5, Dateiende

Anzeige/Ausschrift Erlaeuterung

END OF FILE - Ende der Datei ist erreicht;

- Dieser Text wird als Abschluss
auf die Liste gedruckt.

- Anschliessend kehrt das Pro-
gramm zum Anfang zurueck.

5 Fehlermeldungen

Anzeige/Ausschrift Brlaeuterung

INVALID DRIVE NAHE = Unzulaessiger Laufuerksnane;

INVALID FILE NAHE - Datei dieses Namens jet
der Diskette nicht vorhanden.

PRIF 1520 (SCPX)

Anzeige/Ausschrift Erlaeuterung

- Es ist ein anderer Name einzu-
geben.

PRINTER IS IIOT RBADY - Keine Druckerbereitschaft;

Fehler beim Lesen der Diskette werden direkt vom Betriebs-
system gemeldet. Es werden nach Bedienen von <ENTER> Lesewie-
derholungen durohgefuehrt. Bleibt das erfolglos, bricht das
Programm die Bearbeitung ab und kehrt zum System zurueck.

- 116 -

TRTX 1520 (SCPX)

6. Wandeln einer SIOS-Textdatei fuer die TP-Bearbeitung
RTX 1520 (SCPX)

6.1. Einleitung

Bei einer Vielzahl von Anwendern besteht der Wunsch, bisher
mit dem SIOS-Dienstprogramm TRBXT 1520 (SIEX) bearbeitete Texte
auch mit dem SCP-Programm TP 1520 (SCPX) bearbeiten zu koen-
nen.

Dazu genuest es wegen der internen Struktur der von TEXT
41520 (SIEX) verwendeten Datei "TEXT" nicht, nur eine einfache
Dateitransformation mittels UTR1 vorzunehmen.

Es ist vielmehr folgender Ablauf einzuhalten:

- Mit Hilfe des SIOS-Programmes TEXT ist die Datei "TEXT"
einer REORGANISATION (Kdo. 9) zu unterziehen.

- Es sollte moeglichst nur ein Text in der Datei stehen, da
spaeter unter SCP alles als ein Text betrachtet wird. Die
uebrigen Texte sind in geeigneter Form zu retten und in der
Datei anschliessend zu loeschen.

- Die Vebertragung der Datei "TEXT" erfolgt mittels UTR1 und
in der Vebertragungsart 0. Die entstandene SCP-Datei ist
zunaechst ein Abbild 1:1 der SIOS-Datei. Es sind neben dem
lextverzeichnis auch noch die satzinternen Zeiger und
Steuerzeichen vorhanden.

= Die so entstandene Datei wird nun mit dem Programm TRTX

1520 (SCPX) in eine Form gebracht, die die Verarbeitung mit
TP 1520 (SCPX) erlaubt.

6.2. Aufrufen des Programm
Anzeige/Ausschrift Brlaeuterung

Programmaufruf durch Tastatur
eingabe;

- Abschliessen der Eingabe mit

<ENTER>;
- Programm TRTX.CON wird damit
geladen;
TRTX 1520 (SCPX) V x.x - lleldeausschirft des Programmes

-x.x steht fuer die Versions-
nummer;

- 117 -

TRTX 1520 (SCPX)

6.3. Bingabe der Parameter der Quelldatei

Anzeige/Ausschrift

SOURCE DRIVE NAME:_

SOURCE

Brlaeuterung

FILE NAME: _

- Eingabeaufforderung fuer das
Laufwerk, in dem sich die mit
UTR1 erzeugte Quelldatei be-
findet;

- Eingabe ist nicht mit <ENTER>
abzuschliessen;

Laufwerksname wird auf Guel-
tigkeit geprueft;

Würde nur <ENTER> bedient,
wird das Programm mit Warm-
start verlassen.

- Eingabe des Namens der Quell-
datei;

Eingabe ist mit <ENTER> abzu-
schliessen, wenn weniger als
17 Zeichen eingeben werden.

Wurde nur <ENTER> bedient
verzweigt das Programm zur
Eingabe des Quellaufwerkna-
mens.

6.5. Bingabe der Parameter der Zieldatei

Anzeige/Ausschrift

DESTINATION DRI

vB NANE

Brlaeuterung

Bingabe des Laufwerksnamens,
in dem sich die Zieldatei be-
findet;

- Quellaufwerk und Ziellaufwerk
duerfen gleich sein, wobei
durch haeufige Kopfpositionie-
rung erhoehte Programmlaufzei-
ten auftreten koennen.

- Bingabe ist nicht mit <ENTER>
abzuschliessen.

- Laufwerksname wird auf Guel-
tigkeit geprueft.

- 118 -

DESTINATION PILB NAMB:_

FILE BXISTS
DELETE FILE

TRANSFORM

ALRBADY
WOoR m:

IS WORKING

'

TRTX 1520 (SCPX)

Wurde nur <ENTER> bedient,
verzweigt das Programm zur
erneuten Eingabe des Namens
der Quelldatei.

Bingabe des Namens der zu er-
zeugenden Datei;

Vorangestellte Laufwerksbe-
zeichnung sollte nicht angege-
ben werden. War sie eingege-
ben, wird sie ignoriert.

Bs wird geprueft, ob die Datei
schon vorhanden ist. Bxistiert
sie noch nicht, wird sie ange-
lest.

Datei soll geloescht werden;

Anzeige erscheint, wenn eine
Datei mit diesem Namen bereits
existiert.

Eingabe <ENTER> oder "N":
Eingabe verhindert das Loe-
schen. Das Programm fordert
einen neuen Dateinamen.

Bingabe "Y":
Die Datei wird geloescht und
neu angelegt.

Vor Bingabe des Laufwerks kann
die Diskette gewechselt wer-
den.

Ziellaufwerk und Quellaufwerk
duerfen gleich sein.

Textunformung ist gestartet
(Arbeitszustand);

Wird die Diskette waehrend der
Arbeit gefuellt, wird durch
eine Meldung informiert.

6.5. Dateiende
Anzeige/Ausschrift

END OF FILB

DELETE SOURCE FILE NOW ?
TOR N:_

6.6. Fehlernaohriohten
Anzeige/Ausschrift

TRTX 1520 (SCPX)

Erlseuterung

- Ende der Dateı ist erreicht;

Abfrage, ob die alte Quellda-
tei sofort geloescht werden
soll;

\

Bingabe "N":
Die Datei wird nicht ge-
loescht.

Wurde nur <ENTER> bedient,
wird die Frage wiederholt.

- Bingabe "Y":
Die Quelldatei wird geloescht,
und das Programm kehrt an den
Anfang zurueck.

Brlaeuterung

INVALID DrI VE NAME

INVALID FILE NAMB

DIRECTORY IS FULL

R/W ERROR

= Unzulaessiger Laufwerkename;

Datei dieses Namens ist auf
der Diskette als Quelldatei
nicht vorhanden.

Be {at ein anderer Name einzu-
geben,

- Das Verzeiohnis der Z2ieldis-
kette ist gefuellt. Es besteht
die Moeglichkeit, die Diskette
zu wechseln, bevor der Lauf-
werksname eingegeben wird.

- Schreibfehler beim Anlegen der
Datei. Es kann vor Eingabe des
Laufwerkenamens die Diskette
gewechselt werden.

- 120 -

Anzeige/Ausschrift

TRTX 1520 (SCPX)

Brlaeuterung

ERROR WRITB DISK

R/W ERROR SOURCE DISK

beim Lesen der
Nach

Fehler
system gemeldet.

holungen durohgefuehrt.
gramm die Bearbeitung ab und kehrt zum

- Pehler beim

Schreiben der
Zieldatei ( oder gefuellte
Zieldiskette). Das Programm

bricht die Bearbeitung ab und
kehrt zum System zurueok.

- Sohreibfehler beim Loeschen
der Quelldatei. Loeschen ist
nicht moeglich. Quelldatei ist
nicht geloescht (SOURCE FILE
WAS NOT DELETED).

Quelldatei werden direkt vom Betriebs-
Bedienen von <ENTER> werden Lesewieder-

Bleibt das erfolglos, bricht das Pro-
System zurueck.

- 121 -

LESD 1520 (SCPX)
vA Diskettenanalyseprogramm LESD 1520 (SCPx)

Das Programm stellt äie physiechen Parameter einer unbekennten
Diskette fest und prueft, ob sich alle Sektoren lesen lassen.
Der Disketteninhalt wird nicht zerstoert!

IESD - Programnaufruf durch Tastatun
eingabe;

Abschliessen der Lingabe mit
<ENTZF>;

- Programm LESD,.CO:; wird damit
geladen;

LESD 1520 (SCPX) V x.x

lleldeausschrift den Programmes

x.x steht fuer die Vorsions-
nummer;

PRINT (Y OR N):_ = Eingabe Y:
Protokolldruck wird gewuenscht

Eingabe N:
kein Protokollädruck;

- Eingabe ist nicht mit <ENTER>
abzuschliessen;

- Wurde nur <ENTER> bedient,
wird das Programm beendet.

LOAD SCP SYSTEM DISKETTE

INTO DRIVE A PRESS <ENTER> - SCP-Systemdiskette in Laufwerk
A einlegen und mit <ENTER> ab-
schliessen;

- Damit wird das Programm ord-
nungsgemaess mit \ärmstart
beendet.

Beachte:

Es sollte keine Bingabeaufforderung zum Zweck des Programmab-
bruchs mit “C beantwortet werden, da dann unter Umstaenden der
Warmstart des Systems nicht ausgefuehrt werden kann!

- R2-

LBSD 1520 (SCPX)

Anzeige/Ausschrift Erlaeuterung

Zingabeaufforderung ‚fuer das
Laufwerk, in dem sich die zu
pruefende Diskette befindet.

- Bingebe ist nicht mit <EITER>
abzuschliessen;

- Wurde nur <INTER> bedient,
verzweigt das Program zur
Druckauswahl;

Eine fehlerhafte Eingabe
fuehrt zu der lieläung: .

INVALID DRIVE NAME - Bingabewiederholung richtiges
Laufwerk;
Anzeige der phy: Diskettenparam

Fuer eine Diskette, die dem Format IBi 3740, KROS 5108 bzw.
KROS 5110/01 entspricht:

IDENTIPICATION: 8" [5 1/4"] SIOS DISKETTE
SINGLE DOUBLE] SIDE, SOFT SECTOR: nnnn
SINGLE [DOUBLE] DENSITY
DERBOTIVE TRACKS: 11, 33

Fuer Disketten, die nicht diesem Standard entepreohen:

{DENTIPICATION: 8" [5 1/4"] DISKETTE, SINGLE [DOUBLE] SIDE
TRACK 0: SOFT SECTOR: nnnn,

SINGLE [DOUBLE] DENSITY

OTHER TRACKS: SOFT SECTOR: nnnn,

SINGLE [DOUBLE] DENSITY

nnnn ist die Sektorlaenge und kann sein: 128, 256, 512 oder
1024 Byte.
41 und 3$ sind die Nummern der defekten Spuren.

Anzeige/Ausschrift Erlaeuterung

UNDEFINED DISK FORMAT I11 - Be konnte keines dieser Forma-
te festgestellt werden;

- Das Programm geht zurueok zur
Druckauswahl.

Nachdem die Diskette identifiziert ist, wird sie in der fol-
genden Art sektorweise sequentiell gelesen.

Die Arbeit des Programms kann waehrend des Lesens zu jeder
Zeit unterbrochen werden - ausser beim Lesen der Spur 0 -
durch das Betaetigen einer beliebigen Taste.

- 123 =

Anzeige/Ausschrift

LESD 1520 (SOPX)

Brlaeuterung

TEST CONTINUB (C)
OR END (B):_

= Bingabe 'C':
bewirkt, dass das Programm an
der Stelle fortgesetzt wird,
wo es unterbrochen wurd

Bingabe 'E':
beendet den Diskettentest.

- Das Programm verzweigt zur
Druckauswahl.

Bs ist der Test der naechsten
Diskette moeglich, bzw. das
Beenden des Programmes.

finformation wird auf dem Bildschirm die Nummer der
gerade bearbeiteten Spur angezeigt.
jeden nicht ordnungsgemaess gelesenen Sektor wird eine

Fehlermeldung ausgegeben:

TRACK: nn SECTOR: nn Fehlerart

- nn steht fuer die entspreohen-
de Nummer.

- Die Fehlerart kann in:

MARKED SECTOR
INDEX ERROR
ID FIBLD ERROR
CRC ERROR

Ausserdem kann als Fehlermeldung:

DRIVE NOT READY
REPBAT (Y OR N):_

auftreten.

Das Laufwerk war nicht bereit.

Bingabe 'Y'!:
Die Leseoperation wird wieder-
holt;

- Eingabe 'N':
Das Programm wird beendet.

- 124 -

UTR3 1520 (SCPX)

in

Das Programm UTR3 1520 (SCPX) transformiert Textdateien, die
mit dem Textsystem Text 20 (unter dem Betriebssystem SIOS)
erstellt wurden, in Textdateien die mit dem SCP- Textsystem TP
weiterbearbeitet werden koennen.

Damit wurue die Moeglichkeit geschaffen, eine problemlose
Weiterbearbeitung der Text 20 - Dateien unter TP (SCPX) vor-
zunehmen.

Alle transformierten Text 20 - Dateien werden fuer TP im Al
Hochformat ohne Begrenzung des reohten Randes abgelegt, unab-
haengig, in welchem Format die Dateien im Text 20 - Syatem
erfasst wurden. Es wurde der Grundsatz gewahrt:

eine Zeile im Text 20 = eine Zeile in TP
eine Seite im Text 20 = eine Seite in TP

Transformiert wird pro Durchlauf eine komplette Text 20 -
Datei. Da Text 20 mit einem deutschen Zeichensatz arbeitet, TP
aber allgemein mit dem Latein-Standard-Zeichensatz, ist eine
Auswahlmoeglichkeit im UTR3 vorhanden, sine Umwandlung in
a6,oe,ue und as zu bestimmen.

Die Text 20 - Steuerzeichen Fettäruck, Schattendruck, Unter-
streichen, Index, Exponent, Rueokschritt, Trennzeiehen und
Zeilonendezeiohen werden in analoge TP - Steuerzeichen umge-
wandelt. Geschuetzte Leerzeichen werden als Space (20H) aus-
gegeben. Alle anderen Zeichen werden 1 : 1 uebernommen.

Bine transformierte Datei wird in den haeufigsten Faellen
nochmals ueberarbeitet werden muessen.

Die Zusammenarbeit zwischen UTR3 und dem Bediener ist wie
folgt gestaltet:

Anzeige/Ausschrift Brlaeuterung

UTR3 <ENTER>

vom Bediener.
UTR3.CON ist damit geladen.

UTR3 1520 (SCPX) V x.x
- - UTR3 Meldeausschrift
x.x ist die Versionsnummer

8.2. Parametereingabe fuer die Text 20 - Disket
Anzeige/Ausschrift Brlaeuterung

TRANSFER TEXT 20
PRuBeie - ==-----  mationsrichtung meldet
das Programm;

sich

- 125 -

UTR3 1520 (SCPX)

Anzeige/Ausschrift Erlaeuterun

SOURCE DRIVE (TEXT 20 -

DISKETTB):_ - Eingabeaufforderung fuer das
Quellaufwerk;
Buchstabe B bis E fuer das
Laufwerk, in dem die Text 20 -
Diskette liegt, eingeben.
Bingabe nicht mit <ENTER> ab-
schliessen.

hte:

Fuer die Text 20 - Diskette darf niemals das Laufwerk A ver-
wendet werden, da UTR3 mit Warmstart (“C) abgebrochen bzw.
beendet wird. Damit wird im Laufwerk A immer eine SCP-Diskette
mit Betriebssystem verlangt!

Der Laufwerksnane wird auf
Gueltigkeit geprueft. Wurde
nur <ENTER> eingegeben, wird
das Programm mit Warmstart
verlassen.

VOLUME ID (TEXT 20 -
DISKETTE)

Bingabe des VOL-Name der Text
20 - Diskette;

Bingabe ist nur dann mit
<ENTER> abzuschliessen, wenn
der Name kuerzer als 6 Zeichen
ist. Datentraegerkennsatz wird
gelesen und der eingegebene
Name ueberprueft.

Wird nur <BENTER> eingegeben,
erscheint wieder die Eingabe-
aufforderung fuer das Quell-
laufwerk.

Hat Text 20 - Diskette einen
anderen Namen, erfolgt Feh-
lermeldung und erneute Ein-
gabe des Datentraegernamens
wird gefordert.

Text 20 - Diskette kann vorher
bei Bedarf in diesem Laufwerk
gewechselt werden.

- 126 -

UTR3 1520 (SCPX)
8.3. Parametereingabe fuer die Text TP_- Diskette

Anzeige/Ausschrift Erlaeuterun;

DESTINATION DRIVE (TBXT TP -
DISKETTE):_ - Eingabe des Laufwerksnamens
fuer die Text TP - Diskette;

Beacht:

unt dem Betriebssyst SCPX ini-
mittels SYSG ein Betriebssys darauf
gebracht worden sein, wenn sie im Laufwerk A bi ohrieben

werden soll.

- Danach ist der Name der zu er-
zeugenden TP - Datei nach
den in der "Anleitung fuer den
Bediener" beschriebenen Regeln
einzugeben.

Bine vorangestellte Laufwerks-
bezeichnung wird ignoriert.

DESTINATION FILE NANB

(TP(SCPX) - DISKETTE):_ - Bingabe Dateiname fuer die
TP - Datei;
Es wird geprueft, ob die Datei
schon vorhanden ist. Existiert
sie noch nicht, wird die Mel-
dung "NEW FILE" angezeigt und
die Datei angelegt.

- Bxistiert bereits eine Datei
dieses Namens, wird gefragt,
ob sie geloescht werden soll.

FILE BXISTS ALRBADY

DELETE FILE (Y OR N):_ - Bingabe "Y":
Datei wird geloescht und neu
angelegt.

Eingabe "N" oder <ENTER>:
Loeschen wird verhindert;
Programm fordert Eingabe eines
neuen Dateinamens;

- 127 -

Anzeige/Auss:

UTR3 1520 (SCPX)

DESTINANION

TILE IAIIE
(TEXT 20 - DISKETTE):_

PILE NAUE TEXT 20 NOT FOUND !

DEUTSCH ---> LATEIN (Y/N):_

TRANSPER IS WORKING

'

Ist des
nient
?ehlermeldung
zur Laufwerkseingabe
gekehrt.

Vor der Eingabe der Laufwerks-

Anlegen Datei
meglich, wird eine
angezeigt und

zurueck-

angabe kann die Diskette ge-
wechselt werden.

Bingabe des Dateinamen dex
Text 20 - Datei wie er im Ver-
zeichnis bei Text 20 abgelegt
ist (max. 30 Zeichen).
Ist der Name kuerzer, ist mit
<ENTER> abzuschliessen.

Leerzeichen sind mit einzuge-
ben!

Bs erfolgt eine Suche des Da-
teinamens im Text 20 - Ver-
zeichnis.

Wird der Name nicht gefunden,

erfolgt die Anzeige:

und es erfolgt der Abbruch der
Transformation.
Weiterarbeit ab Grundstellung
des Programms;

Wurde der Dateiname gefunden,
erfolgt die Anzeige:

Eingabe "N":
Text wird 1:1 uebernommen;

Bingabe "Y":
Zeichen ae, ue, oe und sa wer-
den eingesetzt;

Dateiuebertragung findet

statt;

- 128 -

Anseigo/kus schrift

UTR3 1520 (SCPX)

Brlaeuterung®

DESTINATION DISKBTTB IS FULL
NEW DISKETTB (Y OR N):_

END OF TRANSFER

- TP - Diskette ist gefuellt;
Entscheidung, ob die Datei auf
einer neuen Diskette fortge-
setzt werden soll.

- Dateiuebertragung
beendet;

erfolgreich

- Programm haelt in Pro-

granmgrundstellung;

der

- Die naechste Datei kann ueber-
tragen werden.

8.4, Pehlernachrichten und Meldungen

8,4,1 Fehlernaohriohten
Anseige/kus ohrift

Erlaeuterung

InvALID DRIVB NAME

INVALID VOLUME ID

INVALID SOURCE FILE NAME

DRIVE NOT RBADY
REPBAT (Y OR N):

- Unzulaessiger Laufwerksname;
Programm verlangt erneut die
Bingabe des Laufwerkanamens;

- Text 20 - Diskette
anderen Nanen.
Es wird zur Eingabe des Daten-
traegernamens verzweigt. Dor
neue Name isf einzugeben. Vor-
her kann bei Bedarf eine ande-
re Text 20 - Diskette in
dieses Laufwerk eingelegt wer-
en.

hat einen

wurde eine Diskette einge-

die keine Text 20 -
enthaelt. Bine andere
Diskette muss verwendet wer-
den.

- Laufwerk, in dem die Text 20 -
Diskette liegt, ist nicht be-
reit.

-129 -

Anzeige/Ausschrift

UTR3 1520 (SCPX)

Erlaeuterung

READ ERROR SIOS-DISKETTE

DIRECTORY IS FULL

WRITE ERROR SCP-DISKETTE

READ ERROR SCP-DISKETTE

MEMORY TOO SMALL FOR
TRANSFERING

Laufwerk bereitmachen und die
Bingabeaufforderung mit "Y"
oder <BNTBR> beantworten. Dann
wird die Zugriffsoperation
wiederholt. B

Wird mit "N" beantwortet,
bricht der Transformationspro-
zess ab und das Programm steht
wieder am Anfang.

35 kann das Quellaufwerk ein-
gegeben werden. UTR3 kann
nur durch <BNTER> verlassen
werden.

Lesefehler auf der Text 20 -
Diskette;

Verarbeitung wird abgebrochen,
das Programm steht wieder am
Anfang.

Verzeichnis der TP - Diskette
ist gefuellt;

Bas besteht die Moeglichkeit,
die Diskette zu wechseln, be-
vor der Laufwerksname eingege-
ben wird.

Schreibfehler auf der TP- Dis-
kette;

Abbruch! Bs wird zur Eingabe
des Laufwerksnamens fuer die
TP - Diskette zurueckgegangen.

Lesefehler auf der TP-Dis-
kette;

Abbruch! Bs wird zur Eingabe
des Laufwerksnamens fuer die
TP - Diskette zurueckgegangen.

Der zur Verfuegung stehende
freie Hauptspeicher ist klei-
ner als die Blocklaenge der
Text 20 - Datei. Es ist keine
Transformation moeglich.

- 130 -

UTR3 1520 (SCPX)

Anzeige/Ausschrift Erlaeuterung
INVALID FILE NAME - Der eingegebene TP-Dateiname
ist unzulaessig und ent-

spricht nicht dem gefor-
derten SCP-Pormat.

Der Dateiname ist erneut ein-
zugeben.

FILE NAME TEXT 20 NOT FOUND ! - Dateiname fuer die Text 20 -
Datei wurde nicht gefunden.
Abbruch der Transformation und
Halt in der Grundstellung des

Programms .
8 Meldung:
Anzeige/Ausschrift Brlaeuterung

NEW FILE

Bine Datei dieses Namens gibt
es noch nicht auf der TP-Dis-
kette. Sie wird angelegt.

LOAD SCP-DISKETTE AND

PRESS <ENTER> - Aufforderung, eine neue TP -
Diskette in das Laufwerk ein-
zulegen.

FILB BXISTS ALRBADY

DELETE FILB (Y OR N):_ - Die Datei ist bereits vorhan-
den.
Entscheidung, wie weitergear-
beitet werden soll.

DESTINATION DISKBTTE IS FULL

NEW DISKETTE (Y OR N):_ - Auf der TP- Diskette ist
kein Platz mehr vorhanden.
Entscheidung, ob die Transfor-
mation auf einer neuen TP -
Diskette fortzusetzen ist.

Beachte:

Diskette erst wechseln, wenn die Aufforderung dasu erscheint,
da die Datei erst abgenchlonsen werden muss.

- 11 -

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
