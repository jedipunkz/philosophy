---
source: "https://archive.org/details/mkl-20231218-01-edv-MOS-K1520-Anleitung-fuer-den-Systemprogrammierer-SCPX-1526-01"
title: "robotron MOS K1520_Anleitung_für_den Systemprogrammierer SCPX 1526"
author: "VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt"
year: "1986"
captured_at: "2026-08-29T21:08:21Z"
updated_at: "2026-08-29T21:08:21Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "マルクス"
query: "Karl Marx"
plain_text_url: "https://archive.org/download/mkl-20231218-01-edv-MOS-K1520-Anleitung-fuer-den-Systemprogrammierer-SCPX-1526-01/MOS_K1520_Anleitung_fuer_den Systemprogrammierer SCPX 1526_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "社会思想"
  - "唯物論"
status: raw
---

# robotron MOS K1520_Anleitung_für_den Systemprogrammierer SCPX 1526

- 著者: VEB Robotron Buchungsmaschinenwerk Karl-Marx-Stadt
- 初版: 1986
- 情報源: [archive](https://archive.org/details/mkl-20231218-01-edv-MOS-K1520-Anleitung-fuer-den-Systemprogrammierer-SCPX-1526-01)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[マルクス]]
- 研究動向: [[マルクス-現代研究動向]]

## Full Text

robotron

Systemunterlagendokumentation

MOS K1520

Anwenderdokumentation

Anleitung für den Systemprogrammierer
unter dem Betriebssystem SCPX 1526

Die vorliegende 3. Auflage der - Dokumentation "Anleitung fuer
den Systemprogramsierer” - entspricht dem Stand Januar 1986 und
unterliegt nicht dem Aenderungsdienst.

Nachdruck, jegliche Vervielfaeltigung oder Auszuege daraus
sind unzulaessig.

Die Dokumentation wurde durch Kollektiv des VEB Robotron
Buchungsmaschinenwerk Karl-Marx-Stadt ausgearbeitet.

Im Interesse einer staendigen Weiterentwicklung werden die
Leser gebeten, dem Herausgeber ihre Vorschlaege bzw. Hinweise
zur Verbesserung mitzuteilen.

Anwerkun,
Die Anuenderdokumentation zum SCP - System besteht aus

= Anleitung fuer den Bediener SCP 1520

= Anleitung fuer den Programmierer SCP 1520
Teil I und Teil II (Sprachbeschreibung ASM)

- Anleitung fuer den Systemprogrammierer SCP 1520
- Hardwarebeschreibung

= Anwenderdokumentation BASIC- Interpreter

= Anuenderdokumentat ion BASIC- Compiler

- Anuenderdokumentation C- Compiler

- Anwenderdokumentat ion PASCAL

- Anwenderdokumentation FORTRAN

= Anwenderdokumentation Textverarbeitungssys'

= Anwenderdokumentation Installierungs- Programm
fuer TP

- Anwenderdokumentation KOMBO- Druck

- Schulungshandbuch fuer das Textverarbeitungs-
system TP

VEB Robotron-Buchungsmaschinenuerk
Karıl-Marx-Stadt

Software-Zentrum

9010 Karı-Marx-Stadt

Postschliessfach 129

Inhaltsverzeichn!s

ii Einleitung
2. Systemueberbl ick

2.1. Systenladen

2.2. Automatischer Programmaufruf
2.3. Gesamt- Speicher- Adressplan
2.4. Hinweise zur Interruptarbeit
3. BIOS

3.1. Adressplan, BIOS- Rufe

3.2. Tastatur

3.3. Bildschirm

3.4. Drucker

3.5. Diskettenarbeit

3.5.1. Arbeitsweise

3.5.2. Diskparameterhaeder (DPH)
3.5.3. Diskparameterblock (DPB)

Anlage A: Belegte Toradressen und ihre Bedeutung

1.__Einleltung

Die "Anleitung fuer den Systemprogrammierer” gestattet einen
Veberblick ueber das gesamte System (Groesse von Bereichen,
Aufgabenteilung, markante Speicheradressen) und soll Systen-
programmierer befaehigen, nach Bedarf die BIOS Rufe zu nutzen
und Modifikationen am System vorzunehmen.

Voraussetzung dafuer ist die Kenntnis der “Anleitung fuer den
erer“, da fuer Systemprogrammierung die Assenbler-
sprache zu empfehlen ist .

ds Systsmusbschlick
Aslı__Systeeloden

Nach dem Einschalten der Maschine ("Kaltstart”) oder durch
Auslo n von RESET durch zueimaliges Betaetigen der Ein-
schalttaste wird der sogenannte Lade-PROM der ZRE aktiviert.
Dieser belegt den Speicheradressraum von 8 bis 3FFH
Ab der Adres: ® wird das Programm, das in die: Lade-PROK
gespeichert ist, abgear! tet. Es werden folgende, fuer SCPX
wesentliche Schritte ausgefuehrt:

Loeschen RAM- Bereiche 8...7FFH.

Einstellen Interruptmode 2.

= Programm aus Lade-PROM in den RAM- Bereich einlesen.

Abschalten des Lade-PROM.

- Laden des Sys
ABOH.
Der Systemlader belegt die ersten 512 Byte der Spur 9 (Sek
tor 1 - 4 bzw. 1 - 2, je nach Diskettenformat). Er enthael!
auf den ersten 3 Byte die Kennung -SYL
Fehlt diese Kennung, so werden die Di
den Laufwerken (B, Systemlader abgefragt unc
dieser 9 eingelesen (im Laufwerk A muss aber
zumindest eine formatierte Diskette sein !).

laders von der Diskette im Laufwerk A nach

tten in den folgen-

= Aufruf des Systemladers ab Adresse 8437H.

Vom Programm "Systemlader” werden unter anderem folgende
wesentliche Funktionen realisiert:

- Loeschen des Bildschirmes.
- Anzeige “ROBOTRON LOADER“ auf dem Bildschirm.
- EDC- Zeichen- Kontrolle des Systemladers.

Bei Nichtuebereinstimmung des berechneten mit dem eingelese-

nen EDC- Zeichen wird “NO READ” angezeigt und das Programm
nicht fortgesetzt.

- Von der selben Diskette wird das Betriebssystem, das auf
der Diskette direkt hinter dem Systemlader angeordnet ist,
einge! n. Die Laenge des SCPX, die Ladeadresse und der
Laufuerkstyp sind im vorderen Teil des Systemladers enthal-

ten.

= Ansprung des SCPX im Speicher (Kaltstartroutine im BIOS).

h die Kaltstartroutine werden unter anderem folgende
ntliche Funktionen ausgefuehrt:

- Das I- Register wird auf F7H eingestellt, die fuer SCPX
rvierte Interruptsaeule (F7EOH. ..F7FFH) wird ge-
loescht.

Die untersten 2 K Byte des Speichers werden geloescht.

Die Tastatur wird initialisiert.

Die Meldung "SCPX 1526 - V x.y (52 K)” wird auf dem Bild-
schirm auf der zweiten Zeile angezeigt.

- Die Spruenge an BIOS (Warmstart) und BDOS werden auf den
Bytes 8 - 2 und 5 - 7 eingetragen.

- Die aktuelle DMA- Adresse wird auf BBH (Standard- DMA-
Adresse) eingestellt.

= Sprung zu CCP zur Kommandoeingabeaktivierung

2.2

Normalerweise ruft der Bediener Kommandos oder Programme
1.COM- Datein) nach der Anzeige "A)” durch die Eingabe einer
Zeile in den dafuer vorgesehenen CCP- Kommando- Puffer und der
gleichzeitigen Anzeige dieser Zeile auf dem Bildschirm auf.

utomatischer Programmaufruf

Wird dieser Kommandopuffer im CCP mit einem gueltigen Kommando
voreingestellt, so geht SCPX nach Kalt- und Warmstart(!) nicht
in den Systemgrundzustand mit der Anzeige ”A)”, sondern er
laedt und startet das Kommando/ Programm sofort.

4

Damit kann gleich nach dem Einschalten (od
der BASIC- Interpreter oder das Textprogra:

Varastart) B.
aktiviert werden.

Die Voreinstellung des Kommandopuffers von CCP kann bei der
Systemgenerierung mit SYSG (sh. dazu "Anleitung fuer den
Bediener Pkt.4.2.2.) wie folgt realisiert werden:

- Aufruf SYSG und Abarbeitung, bis das System SCPX von der
"Quell”- Diskette eingelesen ist.
Danach SYSG mit CTRL C (im weiteren Text AC) abbrechen.

- Speichern des SYSG mit nunmehr eingelesenem SCPX als Datel
auf der Diskette mit :

SAVE 52 SYSGSCPX.COM (beliebige Bezeichnung)

- Einlesen und Modifizieren des SYSGSCPX.COM mit Hilfe des
Programms DU (sh."Anleitung fuer den Progra

DU SYSGSCPX.COM

der BDOS-
@ 986H angelegt:

Der Kommandopuffer von CCP, aufgebaut wie bi
Funktion 10, ist jetzt ab der Adr

86H maximale Laenge (7FH)

987H aktuelle Kommandolaenge

9B8H -

AD6H Zeichenspeicher fuer Kommando

Das geuuenschte Kommando ist ab 988H, hinter dem letzten
Zeichen des Kommandos ist eine DBH einzutragen. Aut rdem
ist die aktuelle Laenge des Kommandos auf 987H abzulegen.

Soll z.B. der BASIC- Interpreter nach Kalt- und Warmstart
sofort von der Systemdiskette im Laufwerk A aufgerufen
werden, so ist folgendes im CCP- Kommandopuffer einzutragen:

987H @4H
9B8H 42H (”B”)
89H 41H CA”)
9BAH 53H (785")
98BH 49H (719
98CH 99H

- Nun kann das modifizierte SYSGSCPX mit dem DU- Kommando 9199
gestartet und ein System mit automatischem Programmaufruf
auf Diskette generiert werden. Dabei ist zu beachten, dass
die Abfrage nach dem "SOURCE DRIVE NAME” (Quellaufwerks-
namen? fuer das erneute Einlesen des SCPX von einer Dis-
kette durch (ENTER) zu umgehen ist!

- Ist Kommando mit EXT eingebunden, ist der automatische
Programmaufruf nicht mehr moeglich (kein CCP-Nachladen von
Disk!)

2.3._. _Gesgst-_Speicher- Adressplen

Das SCPX ist fuer einen Gesamtspeicherbereich von 64 K Byte
vorgesehen. Dieser Speicherbereich ist wie folgt aufgeteilt?

0) B000H...OOFFH Systendatenbereich
Die wichtigsten Bytes dieses Bereiches
sind im Pkt.2.1. der "Anleitung fuer den
Programmierer” erlaeutert.

b>  CBBOH...CFFFH ccp
(sh. "Anleitung fuer den Programmierer)
wit:

CBBOH.. .CBB2H Sprung zum automatischen Programmaufruf
(nach Pkt.2.2.) oder, falls die Komman-
dolaenge = 8, Sprung zur System-
Grundwarteschleife.

C803H.. .C8B5H Genereller Sprung zur System-Grundwarte-

tohne Beachtung des eventuell
voreingestellten Kommandos).

CBB6H...CBB7H CCP- Kommandopuffer

<)  DOB6H...DDFFH BDos
(sh. "Anleitung fuer den Programmierer“)
4) DEBBH...F7FFH BIOS (sh.Pkt. 3.1.)
mit:
FED. ..FIFFH Interruptsaeule von SCPX.
e)  FEDBH...FFIFH Wiederholspeicher fuer Bildschirm mi
Format BB x 24
bzu. Wiederholspeicher fuer Bildschirm mit
FCOBH... .FFFFH Format 64 x 16

"ver den Anwender steht im Allgemeinen der Bereich 199H bis
Deshalb ist

{JP_9) zu prograi
3008 von der Systemdiskette im Laufwerk A (1) nachgeladen.
Ausserdem werden bei Warmstart die Spruenge an BIOS
(Warmstart) und BDOS auf den Adressen 9-2 und 5-7
zingetragen, die Standard- DMA- Adrı uf 9980H eingestellt
ınd CCP (Adresse CB99H) aufgerufen.

In Systemgrundzustand wird das zuletzt aktuelle logische
Laufwerk angezeigt ("A)”, "BP", "C)”, "DI" oder "EI”).

Fuer spezielle Programme, in denen weder BDOS- Funktionen noch
BIOS- Rufe verwendet werden, kann der gesante cher genutzt
werden. Ua nach Abarbeitung des Programss wieder in das SCPX
zu kommen, ist ein Kaltstart (RESET) erforderlich.

4

2.4._Hinvelss_zur_Interruptarbeit

Bei der Einbindung neuer Geraete, Adapter oder Anschluesse
sollte die Informat ionsuebertragung im Polling- trieb
erfolgen.

st jedoch eine Interruptarbeit erforderlich, so ist folgendes
zu beachten:

= Der Interruptmode 2 ist eingestellt und darf nicht geaendert
werden.

Jede Interruptroutine hat einen eigenen Stack einzurichten.
Der ehemalige Stackpointer ist zu retten, der eigene Stack-
pointer ist einzustellen und dann sollte erst der CPU-Befehl
EI (Interruptfreigabe) folgen. Das Rueckschreiben des alten
Stackpointers vor RETI ist mit DI und EI einzuschliessen
Dabei sind fuer Unterbrechungen durch Interrupts hoeherer
Prioritaet 4 Byte in der Groessenbemessung des Stacks zu be-
Fuecksichtigen! A

- Das Interruptregister ist auf den Wert @F7H eingestellt.

- Die Interruptsaeule fuer SCPX
LBE,CTC) belegt die Bytes OF7EO

lbst (u.a. fuer FD, Drucker,
ÖFTFFH und OF740 - OF7ALH.

Fuer die Interruptarbeit weiterer Treiber, die nicht In das
SCPX ab Version 1.2 gehoeren, wird der Interruptsaeulenbereich
0F742...0F7DFH reserviert. Die Vergabe d Bereiches er-
folgt dynamisch (von adresshoechster zur adressniedrigsten
Stelle), je nach Anordnung der Treiber beim Programmbinden.
Auf Zelle BF7EBH steht der L-Adressteil des adresshoechsten
freien Bytes (Zeiger auf Beginn des freien Interruptsaeulenbe-
reiches). «Bei Kalt- und Warmstart ist dort stets der Wert
SDFH eingetragen. Eine Loeschung der Interruptsaeule erfolgt
Jedoch nur bei Kaltstart!)

Ein Anuenderprogramm kann fuer moegliche Interruptarbeit den
freien Teil der Interruptsaeule nutzen. Werden mit dem Pro-
gramm Treiberunterprogramme aus dem SCP-Angebot gebunden, die
ie_ Interruptbetrieb arbeiten, so ist auf die untere ren
OF742H zu achten. Au ist nach de:
Zelle BF7EBH der sich ergebende neue L-Adressteil des adress-
hoschsten freien Bytes der Saeule eingetragen.

Interruptsaeu-
} le SCPX ab
Version 1.2

Fr Zeiger auf naechste
e Adresse (Low-Teil)

) lenbereich
fuer Anwender

Die zum SCP vertriebenen Treiberunterprogramme tra:
Interruptverzwelgungsadressen unter Behandlung des Zeigers auf
Zelle F7EBH ein.

Treiber UP I _traegt n Bytes in Saeule ein

Neben der Erweiterung d
und RES) enthaelt die
10 ms-CTC-Rout ine.

Mit dieser Routine (Interruptroutine) erfolgt alle 10 ms unab-
haengig vom laufenden Programm eine Tastaturabfrage Die
Tastaturcodes werden in einem 15 Byte grossen Tastaturpuffer
abgelegt und fuer die CONIN- und CONST-Routinen im BIOS be-
reitgestellt. Dadurch ist eine gepufferte Eingabe (Vorein-
tastung) von maximal 15 Zeichen moeglich.

Ausserdem wird mit der 10 ms-CTC-Routine eine Uhr auf den
Speicherzellen 40H...42H (Stunden, Minuten, Sekunden) Im BCD-
Format realisiert.

CCP um zwei residente Kommandos (EXT
triebs-Version 1.2 von SCPX eine

3.1. _Adressolan, BIOS- Rufe

Das BIOS des SCPX beginnt mit dem Sprung- Vektor, einer An-
einonderreihung von 17 Sprungbefehlen, die geraetespezifisch:
BIOS- Unterprogramme aufrufen.

Erlaeuterung

DEBEH Jjp boot #Kaltstartroutine (sh.”An-
leitung f.d.Programmierer).
DEBSH wboote:jp wboot Warmstartroutine (sh.”An-

leitung f.d.Programmierer).

DEBEH jp const sAbfrage, ob eine Taste be-
taetigt wurde.
<A)= FFH = Taste liegt an,
“A)= BOH = keine Taste

DEDPH jp conin sAktuellen Tastencode nach
A abholen. Liegt keine
Taste an, wird auf Taste
gewartet.

DESCH Jp conout Ausgabe eines Zeichens an
"den Bildschirm auf Kursor-
position. Auch Steuerzei-
chen sind moeglich. Das
Zeichen ist In Register C
bereitzustellen.

DEBFH Jp List JAusgabe eines Zeichens an
den Drucker. Auch Stever-
zeichen sind moeglich. Das
Zeichen ist In Register C
bereitzustellen.

DE12H Jjp punch reserviert fuer Lochband-
stanzer

DEISH Jp reader jreserviert fuer Lochband-
leser

DEIBH Jp home }Kopfpositionierung des akt.
Laufwerkes auf Spur 9.

DE1BH jp seldsk jAusuaehlen des Laufwerke:
(LW) Fuer eine folgende FD-
Operation

(@= LuA, 1= LUB...).
Rueckneldung: <HL)= Adresse
ausgewaehlten 16 Byte
langen DPH (sh. 3.5)
Bei fehlendem LU: CHLI=0BHPH

Bezeichnung

jp settrk
DEZIH jp setsec
DE24H jp setdma
DE27H Jp read
DEZAH jp write
DE2DH jp Listst
DE3BH jp sectran
Kenstantenbereiche
DE33H sign: db rBur
DE3SH vers: db cr 12”)
DESEH <posi dw...
DESAH errad: dw erradr

10

Erlaeuterung

Einstellen der Spurnummer
des ausgeuaehlten Laufwer-
kes fuer nachfolgende FD-
Operation. (BC)=Spurnummer.

Einstellen der Sektornummer
des ausgewaehlten Laufu:
kes fuer nachfolgende FI
Operation. (BC)= Sektornum-
“er.

‚Einstellen der DMA-Adresse
<BC)= DMA- Adresse

jLesen eines Records (128
Byte) nach entspr. einge-
stellten Parametern (BELDSK,
SETTRK, SETSEC, SETDHA).
“A)= 8 fehlerfrei

“A)= 1 Fehler

Schreiben eines Recordsz

Parameter, Rueckmeldung und
rholung wie bei READ.
Daten werden als "nicht

geloeschte Daten” gekenn-
zeichnet.
Status vom Drucker abfra-

wobei:
@OH -Drucker nicht be-

reit
«“A)= FFH -Drucker bereit.

Zuordnung log.-phys. Sektor

Vorgabe:

DE) - Adresse der Zuord-
nungstabelle

<BC) - log. Sektornummer

Rueckmeldung:

{HL) = phys. Sektornumser.

Signum BUK.

»Versions-Nr.
triebssysteus.

x.y des

jAktuelle Kursorposition.

ıSpezifische Adresse bei der
DISK- 1/0- Routine

DE3EH

DEABH

DES4H
DE6SSH

idadr: dw

wedb: dw

inter: dw

dskio: dw

treiwba: dw

- DEAFH db
tranad:dw
cdb: ds
bufad: dw
dmaad: ds
vor: db
hd: db

Bezeichnung

cdbwi

intab

diskio

dpbase

trand

13

hstburf

(AaH)

co)

[C))

11

Erlaeuterung

Adresse des ID- Narkenver-
gleichsfeldes des zuletzt
gelesenen Sektors.

scdb- Adresse von Warmstart
“Spur ®).
<db -control device block

Adresse der Interrupt-Saeu-
le der DISK-I/O-Routine.

sAdresse der DISK-I/0-
Rout ine.

sAdresse des ersten Diskpa-
rameterheader.

Anzahl der Wiederholungen
bei Schreib -/ Lesefehlern
des FD.

sdr iver-wboote-adre:

»7 Byte reserviert.

jUnrechnungstabelle 1t. Ver-
schiebefaktor zwischen log.
und phys. Sektornummer.

sControl device block der
DISK-I/0-Rout ine.

Adresse des phys. Schreib-/
Lesepuffers (1 K) fuer FD.

ı2 Byte zur Abspeicherung
der DMA- Adresse.

s»Das bier abgespeicherte
By gibt dual mit seiner
oberen Tetrade die Anzahl
der 8”-Laufwerke und mit
seiner unteren Tetrade die
Anzahl der 5,25”- Laufwerke

s an.
OH = 4x8”, kein 5,25”
jreserviert.

Anzahl der log. FD- Lauf-
werke (max.5).

3.2. Tastaturen__K_7606(04), K_7636(34), K_2632

Die Tastaturen werden im SCP mittels zweier Routinen bedient:

1. CONST: Abfrage der Tastatur, ob eine gueltige Taste anliegt
(sh. 3.1.3.

2. CONIN: Warteschleife, bis ein queltiger Tastencode anliegt
und Abholen des Tastencodes (sh.3.1.).

Es bleibt dem Anuender ueberlassen, wie er die eingegebenen
Zeichen weiterverarbeitet ( unter Nutzung der entsprechenden
BDOS- Funktionen bzu. BIOS- Rufe >.

Es besteht kein Dialog- Regime wie z.B. im Betriebssystem

Die Tastaturabfrage selbst wird im Polling- Betrieb realisiert
(kein Interrupt!). Es ist ein 3- Tasten- roll- over moeglich.

Tabelle der SCP- Tastencodes sh. "Anleitung fuer den Program-
mierer” Anlage D.

3232__Blldsshice

Die Bildschirmanzeige besteht aus Bildschirm und entsprechen-
der Anschlusssteuerung mit Zeichengenerator.

Es sind die Bildschirme BAB 2 (Zeichenkapazitaet 1920) oder
BAB 1 (Zeichenkapazitaet 1024) verwendbar.

Der Bildwiederholspeicher beginnt bei BAB 2 ab Adresse FACH
und bei BAB 1 ab Adresse FC9GH.

Die gesamte Bildschirmarbeit wird ueber die CONOUT- Routine

innerhalb des BIOS realisiert. Ausser der Anzeige darstellba-
Zeichen werden noch verschiedene Steuerkommandos ausge-

(sh. den Programmierer”, Anlage A).

Die Codes im Intervall 88H bis GAFH werden ignoriert.

Alle Codes von GBOH bis OFFH werden mit geloeschtem Bit 7 zur

Anzeige gebracht.

Bei Bildschirmueberlauf wird durch BIOS eine zeilenweise Roll-
funktion ausgefuehrt.

Besondere Bedeutung hat der Code 88H. Er bewirkt die LED-
Anzeige auf der Tastatur (neben bzw. ueber der Taste INS-
Mode), das Warten auf eine Tastenbetaetigung und dann das
Ausschalten der LED. War die betaetigte Taste CTRL C, so
erfolgt ein Sprung zum Warmstart. Bei allen anderen Tasten
wird in das aufrufende Programm zurueckgekehrt. Damit kann
z.B. ein Fehlerzustand an den Bediener signalisiert werden,
der mit entsprechender Tastenbetaetigung die Fortsetzung fest-
legen kann.

12

Fuer die Ausgabe an den Drucker ist der BIOS- Ruf LIST”
Vorgesehen. Es wird das im Register C befindliche Druck- oder
Steuerzeichen (sh.”Anleitung fuer den Programmierer” Anlage C)
ausgegeben. Vor der ersten Ausgabe erfolgt das Initialisieren
der Schaltkreıse, die zur Realisierung der physischen Schnitt-
stelle vorgesehen sind

Um den geraetespezifischen Teil der LIST- Routine variabel
einbinden zu koennen (z.B. mit dem Dienstprogramm SEPR ), ist
direkt vor der LIST- Routine die Adresse der geraetespezifi-
schen Programmteile abgespeichert.

Wird ein Fehler des Druckers erkannt (z.B. durch Fehlermeldung
des Druckers oder fehlende Bereitschaft des Druckers), das
durch Einschalten der ERROR-LED dem Bediener signalisiert.

Mit der BIOS- Routine ”LISTST” kann die Bereitschaft das
Druckers abgefragt werden. Wird Im Register A eine 90H zu-
rueckgemeldet, ist der Drucker nicht bereit, ein Zeichen zu
vebernehmen.

Die im SCPX vorgesehenen Drucker- Interfaces sind in der "An-
leitung fuer den Programmierer” Anlage B beschr ieben.

3.3.___Diskettenarbeit

323.1._Arbeitsuelse

Mit Folgenden BIOS- Rufen wird die Diskettenarb;
home, seldsk, settrk, setsec, setdma, read, url
Waehrend die ersten fuenf Rufe reine Voreinstellungen zur
Laufwerks-, Spur-, Sektor- und DNA- Adressen- Auswahl sind
Ohne sofortige Auswirkungen auf das Laufwerk (home stellt
Spur 8 ein), wird nur bei “read” und "urite” das ausgewaehlte
Laufwerk angesprochen.

Zur wesentlichen Senkung der Zugriffszeiten ist die geblockt:
Arbeitsweise realisiert. Es wird stets ein Block der Groi
1 K Byte von der Diskette gelesen bzw. auf die Diskette ge-
schrieben. bei sequentieller Arbeit nur alle acht
Saetze ein Zugriff auf die rforderlich fuer die
anderen sieben Saetze erfolgt die wesentlich schnellere Arbeit
mit dem Blockpuffer. Dieser ist im BIOS angelegt, seine
* im Konstantenbereich unter "bufad” eingetragen.

r Arbeit mit d zusaetzlichen logischen Laufwerk
et. Bei jedem zu lesenden oder zu
ette.

it realisiert:

4
wird nicht geblockt gearb:
schreibenden Satz erfolgt ein Zugriff auf die Di

Nur

13

323.2. _Diskpargseterhegder_(DPH)

Hit dem BIOS- Ruf "seldsk” wird das logische Laufwerk aus;
waehlt und ausserdem die Adresse des zugehoerigen Diskpar
terheader im Doppelregister HL zurueckgemeldet.

Jedes logische Laufwerk hat seinen eigenen DPH. Dieser Ist 16
Byte lang, von denen 10 Byte fuer Adressen und die restlichen
& Byte fuer SCPX- Interne Zwecke verwendet werden.

Folgende Adressen sind im DPH gespeichert:

tran (Byte 8 und 1 des DPH)_ Es wird die Zuordnungstabelle

logischer - physischer Sektor

adressiert. Enthaelt diese Ta-

belle den Wert BBBBH, bedeutet

das, die 1og.Sektoradre,

gleich der phys.Sektoradre

Das gilt fuer alle log. Lauf-

werke ausser dem zusaetzlichen

log.Laufuerk. Fuer dieses gibt

eine Zuordnungstabelle, mit

r_ der Verschlebefaktor 6
realisiert wird.

dirbf (Byte 8 und 9 des DPH)_ Diese Adresse verweist auf
einen 128-Byte-Puffer fuer das
Einlesen der directory durch
SCPX. Alle DPH enthalten die
gleiche Adresse.

dpbn (Byte 18 und 11 des DPH)_ Adresse des Diskparameterblock
(DPB). Jedes logische Laufwerk
(n = 8...4) hat einen eigenen
Diskparameterblock. Er ist im
Pkt. 3.5.3. erlaeutert.

chkn (Byte 12 und 13 des DPH_ Adresse eines Puffers, der
fuer das Speichern eines
Checkvectors (zur Pruefung auf
Diskettenuechsel) erforderlich
Jedes logische Laufwerk
„.4) benoetigt einen sol-
chen Puffer.

alln (Byte 14 und 15 des DPH) Adresse des Allocat ionvectors,
der die Diskettenspeicherbele-
t

Vectors
bedeutet, dass der Block n auf
der Diskette von einer Datei
belegt ist. Eine "0" bedeutet,
dass der Block unbelegt ist.
Die ersten Bloecke und damit
die ersten Bits sind durch die
directory belegt. Jedes log.
Laufwerk (n=8...4) hat einen
eigenen Allocationvector.

14

Die Adresse des ersten Diskparameterheader ist im Konstanten-
bereich vom BIOS unter "dskb” eingetragen. Alle fuenf DPH
(auch bei weniger angeschlossenen Laufwerken) stehen ab di rn
Adresse lueckenlos hintereinander.

Be3s3. Diskoaranstschlock <DPB)

Die Diskparameterbloecke (je 33 Byte lang) enthalten Parameter
zur umfassenden Beschreibung des logischen Laufwerks.

Byte

Anzahl der Saetze / Spur

2 Blockgroessenfaktor
s bedeutet z.B.: 3
4
EJ Blockmaske (Anzahl der Saetze/Block minus 1)
Es bedeutet z 7
15

4 Extentmaske (abhaengig von der Blockgroesse und
der Anzahl Bloecke/Diskette)
Es bedeutet z.B.: 8 1 Extent/directory- Ein-
tragung
1 2 Extents/directory-Ein-
tragung
5,6 Anzahl Bloecke/Diskette minus 1

Anzahl der directory- Eintragungen minus 1
z.B.: K 5600.10 = 63
K 5602 :127

9,10 Erste 16 Bits des Allocationvectors (Bit 9...15)
dient zum Setzen der Bits im Allocationvector
fuer die durch die directory belegten Bloecke)

11,12 Laenge des Checkvectors
(aus Anzahl der directory- Eintragungen
dividiert durch 4)

13,14 Anzahl der Systenspuren
Die folgenden 17 Byte sind fuer die interne Steuerung der

verschiedenen physischen Laufwerke erforderlich und duerfen
nicht geaendert werden.

Anlage_A

Belegte Toradressen und ihre Bedeutung:

90H @BH - von Zentraleinheit (ZRE) und Tastatur belegt
DCH ».. BFH - CTC auf ZRE

10H ... 1BH - fuer FD-Adapter belegt

iCH - reserviert fuer 1/2” MB mit Anschluss AMB

K5025 (Treiber-UP: MBG.REL)

30H. ..37H - reserviert fuer KMBG K5200 mit Anschluss AKB
K5020 (Treiber-UP: KMBG.REL)

40H - bei Adapter K7025 (BAB) belegt

MBH ... FH - belegt, falls Adapter K6028.48 fuer serielle
Tastatur K7637 verwendet wird

SOH ... SFH_ - belegt Adapter K6028 bzw. K8025 fuer Drucker-
bedienung

60H ... 6FH  - viert

PH... 97H - irviert fuer LBE K6200 (EML.COM 5 Treiber-
UP: LBU.REL)

MH 22. 9BH - viert fuer SLE K6001 (Treiber-UP:SLE.REL)

90H ... FH “reserviert fuer HLE K6003 (Treiber-UP:HLE4.REL)

EEH...EFH - 088, Speichererweiterung 48 K Byte RAM
(Treiber-UP: ZSPAB.REL)

FOH ... FFH -p rviert fuer Programmier- und Pruefein-
richtungen

Die anderen E/A-Tore sind frei und koennen in Anwender-
n

verwendet werden. Bei
innerhalb SCP ist jedoch mit der Belegung weiter Toradressen
zu rechnen.

Verwendung der CTC-Kanaele auf ZRE K2526:

Kanal
Kanal 1 N BDH | durch KMBG.REL
Kanal 2 l BEH | Zeitlimit fuer 10 ms-Takt

16

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
