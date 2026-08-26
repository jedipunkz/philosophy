---
source: "https://archive.org/details/micro_IA41153522_0109"
title: "The PLATO V Terminal."
author: "Stifle, J. E., Illinois Univ., Urbana. Computer-Based Education Research Lab."
year: "1978"
captured_at: "2026-08-26T19:57:39Z"
updated_at: "2026-08-26T19:57:39Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato"
plain_text_url: "https://archive.org/download/micro_IA41153522_0109/micro_IA41153522_0109_djvu.txt"
public_domain: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# The PLATO V Terminal.

- 著者: Stifle, J. E., Illinois Univ., Urbana. Computer-Based Education Research Lab.
- 初版: 1978
- 情報源: [archive](https://archive.org/details/micro_IA41153522_0109)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

ED 200 244
AUTHOR
TITLE
INSTITUTION
PUB DATE

NOTE
AVAILABLE FROM

EDES PRICE
DESCBIPTORS
IDENTIFIERS

ABSTRACT

architecture and programmins of the PLATO V terminal,

DOCUME’T RESUME
IR 009 287

Stifle, J. E.

The PLATO V Teraginal.
Illinois Univ., ‘rbana.
Lab.

apr 78

o6p.

Computer-Based Education Research Laboratory,
University of Illinois, urhana, IL 61801 (33.00).

Computer-Based Rducation

MFO1/PCO3 Plus Postage.
Computers; Flow Charts;
*Erograming

*PLATO

*Input Output Devices;
p P

This report provides a detailed description of «he
which contains

an 8080 microprocessor and is capable of being cperaiezd by programs

located in a host computer.
storing local prograas,
terminal operation,
(Author)

memory.

The terminal contains 3k ot memory for
a 4k ROM resident program which supervises
a 2k ROM character set and 2k of spare ROM

SRESERSSKEKCKRKCKEKREREKCKKEKCEKKEKREKEKCKKEKKCREKRKEKEKRKEA KKK EKER KEKE KKK KK
* Reproductions supplied by EDRS are the pest that can be made

”

from the original document.

a
x

SEEKER SK ESKER EKKRKSKEERARERERERKEKEKEREK EKER EEK KEKE EE ERE EEK KE KH

US DEPARTMENT OF HE aLTH
EOUCATION & WELFARE
NATIONAL INSTITUTE OF

EDUCATION

ED200244

TERMINAI

AB CPD A “T

oi MAN

termina

rog

er
2
J
,]
2
er
4
“t.

>

Tonsole

L Addre

Subrout.nes..
SNES Bee
Var AD LE éeces

PrOGram. . 2.6%

S

Se eweeaeeen 4%
Device Sneri f
vevice of 1t

—

~

large
ut-output (1

sma structure and loadin
memory will mak system behave as
let

memor y

- - ? _ . -
ammable feature
l oO

PLATO IV terminal

—O RS232 INPUT

SERIAL

INPUT/OUTPUT
PORT

=

RCV XMIT
INT INT

IO

DATA BUS BUFFER
(8 BITS)

EXTERNAL EXTERNAL
DEVICE DEVICE

(OPT) (OPT )|

INTERRUPT
CONTROL

&

|

Vv

DISPLAY
IN “ERFACE prcememrss 2 Be. ae

CONSOLE INT

-H———— EXT INT g

SOR INT

SIR INT
DATA

(16 BITS) PLASMA CARRIER INT
PANEL KEYSET

Figure 1.0. Yerminal Block Diagram

service requests t .e 8080 are mad

A device ‘questing service

request
“ne f{ sen mntent

memory; and

4 word

accepted

iterrupt)

erm

Ailabil

ave

ine

Sp aul

ete

wv

E

r+

rT}

m( BYTE )<-DATA)
|BYTED BYTE +!)
Rncnncterstpenieienil | Ricieninetgpiasinaniaal

|
BYTE a1

Meyer

itil ERR-
iTSs<sir YES
‘phhdetteity satttovas @ ABORT <> —-(sn)
NO

yes
| oa
qiomnnianslllpesisnene ASSEMBLE CLEAR ASORT

word LOAD WORD
BITS 2 FNZB |

|
_—————————— 7
}
s
al

a
ITY ERROR ERR-
OR LOST ves

DATA

SIR Flow Chart

Y

DISABLE
INTERRUPTS

{

DISABLE
INTERRUPTS

DISABLE
INTERRUPTS

oe

a
|
“

DISARM
XMIT INTR

|
|
}
|
|

DISARM ALL
BUT XMIT, RCV

RESTORE

Y

ENABLE
ARMED

INTERRUPTS

|
|

|

EXIT TO
INTERRUPTED
PROGRAM

SOR, KST, TP, EXT Interrupts

terminal will refuse
Formation e cept fe 1 LEM instruction (deser ibe
of an LDM instruccion with bit 14, "one" wil
terminal to normal Operating mode.
prevents the terminal from processing data

ent an erroneous mode change word is received

verflowing the job stceck will also cause the terminal
mode. In this case the offending word is treated as though it

rived containing an error.

Keyset
——s___.

t inu.-ates that a key has been pushed on the key-

set. The layout and coding of the keyset are shown in Figure 1.
Touch Panel (TP)
The Touch Panel is an input device which allows the terminal user

Ouch the display surface and input positional information directly

=z

to the microprocessor. The touch Sensitive surface is a 16 x 16 array

Squares and the TP interrupt is generated whenever any square is touched.
Short audio tone is generated each time the TP interrupt is accepted

>

the microprocessor.

External Ing it-Output (EXTS)

The EXT# interrupt indicates that a device attached to
lata ‘us is requesting service. The terminal provides for t}
up to 32 input and 32 Output devices. Such
ional, clude a random-access slide projector for the »rojecti
images on the rear of the plasma panel, 2 random-acc>ss audi
to the terminal user pre-recorded 1 dic
System. Other user-defired
Data rate

on the IO

04! Ose 043 044 045 047 050 O51 173 060 O61 062 073

CTL Sal es aa

000 00! 002 gee 004 020 O21 O22
063 064 065 O74

54 056 161 16 a 162 |
ILICIC JIL OIC Je3 Gs tele)
ERAS HELP
cats
0! 120 023 034

\27 105 =; Te 13 117

163 144 146 1\47 150 154 174 O06 O 075
2

Sdcocuscaacescae)

E
106 110 112 134 O26 O27 030 035

=

175 07 O72

AD SOOOGOOOQOC Eek

103 126 102 1:6 137 | a6
140

BS
SPACE

| Ol2

Keyset Coding

lL r
} a)
i

) >
eS

10
I r 1 interrupt 1 ised ternally by the terminal.
mnsole Mode (CONSOLE
yntained within the resident program is a routine which permits

be used

“ONSOLE

In CONSOLE mode the user may di

the contents of memory,

Switch

as program debugging aids.

must be placed in

splay the

alter the contents of memory,

the

CONSOLE

contents of the 8080 regis-

step through

jump to other

message

programs one instruction at a step, set a breakpoint, and
programs. This feature is described in detail in section 3.3.
carrier Interrupt (CARRIER)

This signal is used to indicate an interruption of communication
with the central computer. The resident program will generate a

on the display indicating loss

rejyuires a modem with carrier detect

Interface Unit (DIU)

Display

The DIU contains the registers

ficiently attach

1.4, contains the 9-bit x and

Figure

<-h
it s

it parallel data register

the write-erase control circ

‘control information

y registers are bidirectional counte

trolled by the 8080.

arallel data upper (PDU) and parail

DU/L registers are used only when o

of communication.

a plasma panel to the data bus.

(PDL/U) ,

capability.)

and contro!

(Use of this

The DIU,

feature

circuits required to

shown in

y display address registers, a

uits.

is supplied to the DIU from the data bus.

rs which can be

el data lower (PDL)

perating parallel

a 3-bit display mode

registers.

(PDM) regis-

Data for the registers and

The x and

independently con-

The parallel data register consists of the 8-bit

The

input display de-

vices.
The format of the PDM register is shown in Figure 1.5. Bits 0 and
l specify the write/erase mode, and bit 2, the panel operating mode.
f pit is "zero," the display is operating in the serial mode; and the
ntent of the x and y registers specify the address of a point to be

9 BITS 9 BITS
x ADDRESS Y ADDRESS

|
a

ERASE
16 BITS DATA a
Ls > DISPLAY 4 BULK

READY

Figure 1.4 Display Interface Unit

written or erased. If bit 2 is "one," the display is operating in the
parallel mode; and the contents of PDL/U will be written or erased on the
panel at the address specified by the contents of the x and y registers.
The data will be displayed in a vertical column with bit 0 of PDL at

the bottom and bit 7 of PDU at the top.

Information is written on the display if WE, = 1 or erased of WE, = 0.

The use of these bits is explained in section 2.5 in the discussion of

Mode 3.

I oO
WE, |WEo

Figure 1.5, Pa Register

}

ne

memory
character set.
process PLATO data

ca. and IO routines in

user programs located

resident in

Word Format

The data to be

OPERATING

processed by the terminal consists of 20-bit words

(with start bit removed) with the format shown in Figure 2.0.
i9 18 oO! OO
c/D DATA Pp
— — — -
Figure 2.0. Terminal Word Format
Bit 900 Parity bit - even parity
Bits 01 - 18 Data
Bit 19 Control bit - O = control word
- 1 = data word
: Terminal words may be of two types: control words and/or data words.
E Data words ‘c = 1) contain the data to be processed by the terminal while
‘rontrol words (c = 0) are instructions used to establish operating con-
lJitions within the terminal.
\
Mi
\f 2x3 ‘ontrol Word Formats
\ The PLATO control word format is shown in Figure ..1.
19 18 17 16 15 Oo! OO
e) | D CONTROL INFORMATION P
eee

Bits Re ee So
Bits 16 - 18

)_(NOP)

Figure

2.1. Control Word Format

Control Information

Type of Control Word

peration ins There

hose generated by the PLATO communications har

}) and those generated by the PLATO software

e hardware NOP is generated automatically when the central

no data to be transmitted to the terminal. The software
system software to insert timing delays in the output
The software NOP will cause the terminal word count to be incremented
the hardware NOP will not affect the terminal status in any way.

e

ither of the NOPs is stored in the job stack.

D = 001 (LDM) Load Mode (267us)

19 §633

MODE WORD

WORD COUNT

instruction establishes the operating mode »f the terminal
each mode of terminal operation, there is an associated mode wor«

16) which directs the processing of incoming data.
in a given mode, the terminal remains in that mode until receipt

a new LDM instruction.

Eight different processing modes are available, five of which are

incorporated in the terminal resident and three of which are reserved

ir

for local programs. The processing modes are described later

section.
If bit 14 (WC) of the LDM word is "one," the word count register

be set to the value specified by bits 07 - 13. It.is the receipt

the terminal to

instruction with bit 14 set which will restore

of this
normal mode if it is in the ABORT mode. This is the only instruction
the terminal will accept if it is in ABORT mode. Receipt of the
struction while in the ABORT mode will clear the ABORT flag
initialize the word count, but will not alter the terminal processing.
used to actuate or inhibit external

>

devices attached to the ae >ipt f an LDM word with

will disable the

terminal mode word format

06 O05 O§ OO C2

MODE WE, WE,

Figure 2.2. Mode Word Format
Screen Command. If this
the entire display is

Select write or erase function in the
DIU as follows:

write character background.
erase character background.
suppress character background

suppress character background

terminal processing mode.

fr

(LDC) _Load Co-ordinate 225us

5 14

Wl

instruction loads the x register it 2 = 0) or the y regis-

data as follows:

Load register with bits Ol - 09.
Load register with the arithmetic
sum of its present value and bits

Load register with the arithmetic

difference between its present value
and bits 01 - 09.

nt velue 1s
‘eturns performed

plotting vertically

Request

VL, (FAS fd PG STATUS REQUEST

; y a 4 y , ae ee
eee: * Se Meee ee ey te se Ae AD a

is used to request information from the

terminal or to
operations to be performed. Presently used codes are:

Iperation

Request terminal type. Termiral
responds with code 73.

Generate 1 second audible tune
(Touch panel must be attached for
this operation.)

Request terminal status information.
The lower seven bits of the memory
location specified by the Memory
address register (MAR) will be re-
turned to the best computer 1n a
Status Response word (see section

hoe? «

| Memory Address (229s )

INITIAL MEMORY STORE ADDRESS

Memory Address Register (MAR). This

orage address to be used upon entry into

2300 (HEX) are reserved for use
me

Attempting to write below this
terminal behavior.

5 i 10 QO9 QO8

aint heeled
IO ADORESS | | I
seiieieiataai ARE eS i i

This instruction is used to read and write da

to the external bus and to enable interrupts.

ii = 25 specify the device address.

10 specifies a read (input) operation
if a "1," a weite (output) if a “0.

09 inhibits the actual read or write
function, but permits the device ad-
dress to be saved by the terminal.
The inhibit write function can be used
to establish a write address for
later use by the EXT command. The
inhibit read function can be used to
establish a read address to be used
upon the occurrence of an external
interrupt. If location m.extpa con-
tains 0, the resident will perform
a read from the selected device and
transmit the data to PLATO via an
External input word. PLATO may use
a local program to process the in-
terrupt by previously loading the
program and storing the program ad-
dress in m.extpa.

Write addresses 0 and 1 are special cases of the SSF instruction.

is assigned to the slide selector and for this device only
the data field is 10 bits long. The SSF word format for this device

waa —s =

10wn below.

1! lO Q9 O8 00

Pa ° |

- 08 are sent to address 00 and bits 9 - 10 to address Ol.

8 select one of 256 slides for display on the plasma panel.
09 controls the projector shutter. For normal operation this bit

lways QO. However, if this bit is a "l," the shutter will be closed

closed until receipt of a load slide command with bit 09 =
controls the project lamp. The lamp will be turned on if

ve ONS OEE 2F

assigned to the Interrupt Mask register located
word format for this address is shown below. Mem-

will be loaded with a copy of the interrupt mask

it =610)6F€OS O07 O06 O05 O04 O3 Oo! OO

Bit 06 enables the Touch Panel if a "1," disables it
14 enables the External bus interrupt if a "1," disables it if a
(This data is not actually transmitted to address 0l because the

Interrupt Mask Register is internal to the terminal.)

Following execution of this instruction, memory location M.enab

copy of the mask data.

Load Fy‘ernal (275us)

instruction transfers two 8-bit bytes, byte 0 first, to the

external device selected by a previous SSF instruction.

Processing Modes - Mode 0
normal operation, the terminal is assigi:ed an operating mode by
it a LDM instruction followed by all of the data to be processed

mode.

terminal resident program contains the programs for processing
five modes. In addition, up to three additional user-defined-

grams can loaded into RAM.

a point-plotting mode. Each mode O data word, Figure

address of a point on the panel to be written or erased.

19

The W/E, bit in the mode word determines which operation is performed.

is #18 io O9 oO! OO

Figure 2.3. Mode O Data Word
The processing time for a Mode 0 word is 238us.

Mode 1
Mode 1 is a line drawing mode. Each data word, Figure 2.4, speci-
fies the terminal coordinates of a line, tie origin of which is contained

in the x and y registers.

ig ~=18 10 O9 Ol OO

Figure 2.4. Mode 1 Data Format

The terminal point of a given line is also interpreted as the ori-

the next line. Line origins may be relocated, however, by the use

“4
»-
a |
rr

2f the LDC command without exiting from Mode 0l. A line will be drawn if

“nr "

—. a8 a “k;" erased if thie bét is a ).
“ The processing time for a Mode 1 word ranges from lms for a line

length of one dot to 1ll.lms for the maximum line length of 512 dots.

Mode 2 is a load memory mode. Each Mode 2 data word, Figure 2.5,

contains two eight-bit bytes to be stored in “AM memory. These bytes
are stored, lower first, in two successive locations starting with the
present contents of the memory address register (MAR). After each byte

ls stored, the MAR is automatically incremented by l.

9 6 I7 . o9 08 O1 OO

a tee,

C Co} BYTE | BYTE O

Figure 2.5. Made 2 Data Word

20

As each byte is stored, a longitudinal parity check is performed by
exclusive "oring" each byte with a check word and left shifting the re-
Sult. This chesk should be all zeros at the conclusion of the transmis-

sion of a block of data.

Bits 17 and 18 (Co» C,) activate the block error check as follows:

Ci Co

0 O Store data.

0 ] Byte 0 is the last byte of this trans-
mission and contains the block check
character to make the longitudinal
check word all "zeros."

l Byte 1 is the last byte of this trans-
mission and contains the block check
character.

1 l Not used.

After receipt of the block check byte, if the check word is not zero,
the terminal automatically transmits the STATUS REPORT code (05H) to
the central computer. If the data being loaded is character data, it will
appear when displayed as a vertical column with bit O01 of byte O at the

bottom and bit 16 of byte 1 at the top.
The processing time for a Mode 2 word is 288s.

Mode 3

Mode 3 is a character-plotting mode. The data words in this mode

‘ontain three 6-bit character codes as shown in Figure 2.6. Each code

selects one of 63 characters from one of the character sets contained

in the terminal.

9 8 3s O7 06 O!| OO

-
a CHAR | | CHAR 2 CHAR 3 P

s

—_"

“igure 2.6. Mode 3 Data Word

we

The terminal provides for up to eight character sets of 63 charac-
ters each. Character sets MO and Ml are contained within ROM memory
and hold the characters shown in Table 2. The other character sets con-

tain user-defined characters and are stored in RAM.

The contents of the character memories are processed in Mode 3 as
63 arrays of sixteen 8-bit words. The contents of 16 consecutive addresses
a displayed as one character within a matrix as shown in Figure 2.7.
{ . top three rows and bottom row of all character matrices from MO and

Ml are always unfilled.

“
ort 0 00100000
al TTT] ' 00001110
geereag 2 9000/0000
BRRESe 3 90001000
3/5]? | 9/113} 4 000! 0000
| | 5 0001000!
oe 6 000! 0000
1 | | ?7 00010001
t+ ++ ++4 8 50010000
1 | 9 000! 0001
iO 00010000
|a|6 joke 2 99948086
0 2/4/68 |;Oli2\4|
word —t | | | 13 00001000
no [| | 14 00060000
ee 15 QOO0O0000
te a BE AS ots Se TE
MATRIX STORED RESULTING
ORGANIZATION DATA DISPLAY
choracter choracter 4
origin if origin if °
forward plot reverse plot
Figure 2.7. Character Matrix
The data for the characters stored in ROM is shown in Figures 2.8.0 -
4:0c2

The contents of any character memory can be enlarged via selection
of character size 2 (size 0 is normal size). Selection of size 2 will
result in a 2X magnification of the characters. Figure 2.8.4 illustrates
‘-haracters drawn in size 2. All character format operations will be auto-

matically adjusted when using size 2 characters. The method of specifying

character size is described later in this section.

LD
Sat

Tharacter write/erase is specified by the write/erase bits WEO,
WE1l in the mode word. (See LDM instruction.) If WEO = l, character:
are writte.; 1f WEO = 0, characters are erased. The inverse of the opera-
tion called for by WEO will be performed on the background or unfilled
portion of the character matrix if WE] = 0; while if WEl = 1, the back-

jround remains unaltered.

Character plotting speed ranges from a minimum of 355 characters
per second using a serial plasma panel and up to 3080 characters per second

ising a parallel plasma panel.

There are several non-plotting control characters available for

rmatting the display of data in Mode 3. These control characters may

be accessed via the use of the “uncover” code (77). Upon receipt of a

/7 code, the terminal interprets the next character code as a control char-
acter instead of a character to be plotted. Following execution of the
‘ontrol character, normal plotting mode is resumed. If several uncover
(odes are sent in sequence, the first non-uncover code will be treated

as the control character.

,

he operations performed by each of the control characters are shown
in Table 3. In the case of some characters, the operation performed is
i function of the character memory, the plotting mode, horizontal or

vertical, and the plotting direction, forward or reverse, which is being

forizontal plotting mode is set by the 30 code; vertical plotting,
by the 31 code. Forward plotting, set by the 32 code, is from left to
right in horizontal mode and from bottom to top in vertical mode. Reverse

plotting direction, set by the 33 code, is from right to left in horizon-

tal mode and from top to bottom in vertical mode.

Boidface characters are selected by the 35 code, normal size by

The terminator code (00) is used to terminate character string plotting

from |] i] rograms. See the discussion of r.chars in section 3.l.

teat
+4

2, @

+

i

ee

+

eae
Li,
=

+4
beep tee

+ tee

tse ft

e *
4-4+=$2.} - oR en oe — 4— = a
eCOCCOCOC@ mm ce - aia!
+ bps is +++ Cae + - 4 .

oe 6 b> -b 4-+- + ~~ 5-444 —

us !

: aie dilinaniiatn diiaammaiin aan amet ae

aa

=+ 4

haracters 00

nq
—

7

'e/6,

MO

Jae!

+
+
4

+4
+
+
,

+4
++
be
an
++
++
+

on owe. JU
B0/0010.0 4

thnendhonih

8.0.

2

>

‘
.

+ & be
+

Figures

a ap
ele, | | le

{!

26

- +4

|
bb pe be ey

f

ae

i
Bane anea

ee ee oe |

re
b+ +p + ++ +4

f

se
T}

vote rey
=

ee cans aa Se See Saw ee aw Wee Cae ee ae Oe
I ee

7-4 - 4-4-4
om A OS NE As SS a Oe ee Oe

c

im *

Y

bpp eo ye
| | |

}
r-
++
t
f
L

++
' rH

,
aE +
pS Se SS ee ar A Ee —

SSSR Ss
rT

ee 7
sagimmyniprmeciigee,

Characters

desde,

° sien

i . @
Ll #e Se @eee |

;
;
‘

dis

+ + b+

iio je ne

‘s
J

wT

ro |

:

saapgeeses
++
A I . isigie ielejeisjeie, | +H

LM
+
4
‘

4
‘
4

4
+

——)

Sedge
+++

I

 & fp ee
~ b>

: BES « tt
oe ~

t ae
eas sar
+e

i

$+++++++
~9- & +e

b + ++ ey

X

fs

-

indiana aan
be

+t pp
eee ty
@eesese

*

i

cre
| anda cea |

aor
ee |

r

eens is

T
+
t+ bh eee

ppb +> +4
> +> + +++

PL it

wave ww

ttretet

rreTrrTy.

q

+--+ +

aracters

--+

‘egaeces

I
Ch

T '
‘Sewer e wR e

>

Figures

o

This terminal 1s a prototype
for a PLATO V terminal. It is
operated under the supervision
and control of an INTEL 8988

microprocessor. It 1s actually
a miniature — driven
time-shared computer system

with a | adie panel attached

to the 1“o bus. The terminal
contains 12k bytes of memory,
8k of which are RAM which can
be used to store data or

s which can be executed
. terminal.
ABCDEFGHIJELMNMOPORSTUVIUxMYS 54?!
G6123456769<2 (]) S%_'* )+-+"¢,.¢37

ProLr ali:
=

Figure 2.8.4 "Boldface" Character Set

| ADDRESS ADDRESS

(OCTAL) G (OCTAL)

40
41
42
43
44
45
46
47

50
51
52
53
54
55
56
57
60
61
62
63
64
65
66
67
70

71

Se wee ee Se ee Oe Oe 2 eS ee ee Se ee He ne ee

72

Q

73

74

75
76

77

Table 2. ROM Characters

HORIZONTAL

VERTICAL

SIZE

Terminate

Backspace

Line
Vertical Tab

Feed

Fiorm

Carriage Retu

n

iy~-16

Terminate

x-8 x+8

x+8 x-8
y-16

y+16 y+16

x>504

x>f

x (MARGIN)

y-8
y+8
x+16
x-16

y>8
x15

y*— (MARGIN)

x+16

select

select

select

- 7 of
seiect

Character String Plotting

x+16

”
yte

y-8 x-16

x+16 y-32

y-32

x-16 y+32 y+32 x-32

x>496
y>480

y>2 yo49¢

x31

x>9
y>480

y>5£4

x<-15 x3]

x*—(MARGIN)
y-32 y~32

y*— (MARGIN)

x+16 x+32 x+32

y+10 x-10

x+10

y-10

character memory

character memory

character memory

character memory (RAM)

character memory (RAM)

character memory (RAM)

character memory 6 (RAM)

character memory 7 (RAM)

horizontal plot mode

vertical plot mode

forward plot direction

reverse plot direction

normal-size characters

large-size characters

Characters

served as address regis

*haracter sets These

by PLATO <& the character

WY

Function

“ . oo
Address E

> 30¢€ Character set O2 origin.

Character set origin.
Character set origin.
Character set | origin
Character set origin

Character set origin

Table 4. Character set Address Registers

4 is a block erase mode. In this mode each pair of data words

of an area to be erased. The area erased is that

tne corners

,
\ : > |ly.-y = and y. = Yj.

closed by \x| = A \ \
Y,, a vertical line

horizontal line is erased.
lescription assumes = 0; if WEp = 1, the area is written.

A

00
P |WORD 2

Mode 4 Word Format

mode, the terminal leaves

ee This is the address

address registers set to x) and y,7!

writing characters in the erased area.

The word format for any of these modes is defined by the user. When
operating in any of these modes, the resident places the PLATO data in
D, and E registers as shown in Figure 2.10, and transfers control

’

local program.

E

Dig Orr | De Dis Dra [Ors = |2o9| or Do4|Oos oa| Day |

Figure 2.10. Modes 5,6,7 Word Format

memory locations are reserved for use as address registers
origins of the local programs. These addresses are

a

Address Function

300 Mode 5 program origin.
Mode 6 proyram origin.

Mode 7 program origin.
Table 5. Local Program Address Registers
he resident has transferred control to a iecal program,
‘ontrol of the terminal remains with that program until a
return instruction is executed (or until the clear switch

‘% ie interrupts are left enabled the res

lent wil rontinue to perform all IO functions.

4

Jutput Data Format

o£ 32

Data transmitted from the te

Nh

2
> |

rminal to the computer center consists

-bit words with the format shown below.

iI 10 OS Os

i alae
Bit, — J i Identifies

Ol

Start Pority Bit,
Always "|" Type of Data Odd Pority
The six types of terminal data words are shown below:
i! 8) 0S O08 o7 O!| OO
| fe) KEYBOARD DATA P |
i 09 O08 oO7 Oi OO
! @ STATUS RESPONSE P
li 0 O09 o@ OS 04 QO! 0O
TOUCH PANEL
4 .@) ! x | Y | P
L
—
E | 8) WORD COUNT | P
fatus Request code 70H is used to request terminal type. A re-
‘ f 73H indicates an 8 80-type terminal with 8k of RAM memory.

The Unsolicited Status Report codes are used to inform the central
occurrence of some special event within the terminal.

of the presently used codes is shown in Table 6.

Status Report

(hex) Event Reported

02 Reset (clear switch has been de-
pressed).

05 Longitudinal parity error occurred
in Mode 2.

Table 6. Unsolicited Status Report Codes

RESIDENT PROGRAM

eneral

The terminal memory allocation is shown in Figure 3.0. All of ROM
and that portion of RAM below address 2300 (hex) is reserved for use
by the resident. The resid2nt program contains those programs required
to process PLATO data plus routines for operating the serial communication
port, the parallel IO bus, servicing interrupts and communicating with

the Display Interface Unit.

The push .*own stack is used by the resident to store the status of
the termina] duiing the processing of interrupts. The job stack provides
temporary s‘orage for incoming jobs in the event the processor is busy.
The resident and PLATO variable sections of memory contain terminal status

information which may be used by both the resident and user programs.

Resident Subroutines

The resident program provides several callable subroutines which
may be referenced by user programs. In using these subroutines the follow-
ing convention should be observed:
Single argument subroutines will have the argument passed in
the HL register pai.

Double argument subroutines will have one argument passed in
HL and the other in DE.

Results are returned from subroutines in HL.

The user must provide for saving and restoring any register
or status he wants preserved.

Reference to a subroutine should always be by symbol and not
memory address.

RESIDENT

Li Le

M@, MI

CHARACTER
DATA

JOB STACK

MODE O05
PUSH DOWN

STACK MODE 06

MODE 07

RESIDENT
VARIABLES

CHARSET 2

a CHARSET 3
VARIABLES

CHARSET .

USER CHARSET 5

STORAGE

CHARSET 6

CHARSET 7

Figure 3.0. Memory Allocation

Following 1s ¢ ist of the subroutines.

r.init (40)
This routine will initialize the terminal operating conditions as

> ,
follows:

The screen will be erased.

Memory locations m.margin, m.ksw, and m.extpe will set

to zero.

M.enab will set to enable the serial input port (SIR,
SOR, CARRIER), and the keyset (KST) interrupts.

Select character memory m0; select normal character size;
select horizontal left to right plotting mode.

Initialize other pointers required by the resident.

Remove ABORT condition if it exists.

A jump must be made to this routine in which case control
will be returned to the resident after execution.

sutine will write (or erase) the point on the display screen

ecified by the contents of the HL and DE registers.

YM es
| XE YHA SASSAS

?

location m.mode specifies a write operation if a

erase operation of a "0. After execution of this routine

y registers i » DIU will contain the values entered

ine (46)
routine will write (or erase) a line on the display screen
sciginating at the current address given by the x and y registers

und terminating at he address contained in HL and DE.

WEE
TUE \ Vole cf afer

tt

37

The WE, bit in m.mode specifies a write operation if a "l," an
erase operation if a "0." After execution of this routine the x
and y registers in the DIU will contain the values entered in HL
and DE.

Note: This routine will enable interrupts via execution of an EI
instruction.

r.chars (49)

This routine will write (or erase) a string of characters on the
display. Register pair HL specifies the string origin address.
The string must be terminated with an uncover code followed by the
terminator code (7700). Character coding is the same as shown in
Tables 2 and 3. Character write/erase is specified by the WE, and

WE) bits in location m.mode as described in the discussion of mode

a.

r.block (4c)

This routine will erase (or write) an area of the display screen
specified by a list of coordinates stored in memory. Register pair
HL contains the origin address of the list. The coordinates are

stored in the following order:

hl Fiala a x0 x) lower
hl+l 0000000x8 x, upper
h1+2 VY] =-s><< yO y, lower
} } l
hl+3 O009000y8 Yy upper
h1+4 Xie = xO x. lower
h1+5 2000000x8 x. upper
h1+eé y?7------ yo y. lower
hl+7 O000000y8 Y. upper

9rdinates x,y, and x.y. are any two corners of the area involved.

7 7 5 )
~ » — =
Tr area will be erased if the WE. in location m.mode is "0," written
{J
ec . ' Al
i it 3 “ees
4 ~

r.inpx (4f)

This routine will return the current display x address in HL.

H

Co[oTololo [olor beolFoshosbed edo Pod

r.inpy (52)

This routine will return the current display y address in HL.

4

CeTeToTe[eToTo fea] FefeslesPoses Peel Po

r.outx (55)

This routine will load the display x address register with the con-

tents (lower 9 bits) of HL.

HOODOOOZ E222 22242

r.outy (58)

This routine will load the display y address register with the con-

tents (lower 9 bits) of HL.

CoTeTeTe [ole o Pe) faked Paes fela Fe

r.xmit (5b)

This routine will transmit to PLATO the contents (lower 10 hits) of

HL.

Cofofofofofofrala) [

source of data as follows:

keyset or status response word.
touch panel word.
external input word.

word count or unsolicited status word.

The output word formats are described in more detail in section

(5e)

This ro“tine will load location m.mode with the terminal and dis-

lay opecating mode as specified by the contents of HL.

H

YiIA”’/"J0 222

specifies a full screen erase (the screen will
be erased when this routine is called if s is
"1," the screen is not affected if this bit is

the write erase mode as follows:

erase, rewrite mode.

write, rewrite mode.
overstrike mode.
overstrike mode.

processing mode.

he screen

-mode with

ll increment or decrement the display address
contents

it OF

location m.dir specify the operation

performed as follows:

x0 | YO

i

specifies x direction as follows:

forward (increment)
reverse (decrenent)

specifies y direction.

l increment or decrement the display

Vv

y address regis-

contents

s of location m.dir specify the operation to be

a

as in r.stepx.

(or erase) the current xy address.

setup the x and y direction flags for later use

Location m.dir will be loaded

with contents

specifies x direction as follows:

forward (increment)
reverse (decrement)

direction.

r.input (6d)

This routine will return in HL the last keyset, touch panel or ex

ternal word received by the terminal. If location m.ksw is 0O, the

word will also have been sent to PLATO. If m.ksw = 3, the word is

returned only to the users program.

The format of the data is shown below.

H

Lo }ofo}o}o}o ftofte} |

"O" if data present, "l" of no data.
specify source of data as follows:

0 keyset word.
O 3 touch panel word.
1 0 external word.

r.sst t70

This routine is used to read and write data to devices on the ex-

ternal bus and to enable interrupts.

[© |aca}Acs|Ace|Aon Aco € | |

specify the device address.

specifies a read (input) if "l," a write (outpu
 & nyt ti

if this bit is "l," the device address will be
saved by the resident, but the actual read or
operation will be inhibited.

ssf write operation, L contains the data
be sent to the external device. nan ssf read
operation, r.ssf will return in L the data read
from the external device.

=

t)

to

normally used to establi
for later use by the r.extout routine (described
hibit ssf is used to establish a read addres
an external interrupt processing program. In thi

must have previously loaded location m.extpa with

tnterrupt program. If m.extpa contains 0, the res

the read operation when the external interrupt occ

memory and may be retrieved via the
the resident interrupt program is

a 3 to prevent the data from also

Write addresses 0 and 1 are special cases of the
re address O is assigned to the slide selector
the data bits are sent to address 0O and bit
to address l.
address 1 is assigned to the Interrupt Mask

located within the terminal. The data format

is described below.

serial input port interrupt
keyset interrupt
touch panel interrupt
external IO interrupt
console switch interrupt
modem carrier interrupt
interrupt is enabled if tl associated bit i
.-" Memory location m.enable will be loaded

interrupt mask data.

routine is used >) establish character plotting condition:

by the r.chars routine.

WILLA, [| ekeo Fa] ]

specifies that the last character plotted was the
“uncover™ code (77). This bit should normally

i

be set to "0."

specifies the character plotting direction. A
"O" = forward, a "1" = reverse. Forward direc-
tion is left to right in horizental plotting and
bottom to top in vertical plotting.

specifies character size, 0 = normal, 1
specifies character memory.

specifies horizontal plot i )," vertical

if att hae
m.ccr will be loaded with the data

.extout (76)

ine will transmit the data stored in a buffer t

selected device on the external IO bus. The device

used » execute the
a maximum of 1.48 second
sutput (excluding NOPs). > number of jobs awaiting

contained in location m.

retrieve tne

is returned in

L

Poe|020|0o] [© | © [Poe|0=|Ova)Pvs)Pv2)Por

the resident. Location

decremented ) I Ob ‘trieved.

this routine when m.jobs is zero (emtpy
cause erroneous data to be returned in DE

routine may the resident for execution the

sontained in gis E L. The format of the job is

of memory contains information used by the resident

scal programs. This data may be read direct-
exception « m.dir and m.ksw, should be

routines Ss ibed earlier.

and

The lower series
location are used to
of
program.
3 (bit) is
resident program
all flags

dd to the

Address Function

fies present Margin setting

used for carriage return:

specifies number of job:
maining in job stack. See
r.gjob and r.exec for use
this data.

specifies character mode plotting
conditions. See r.chars for
definition of data.

specifies terminal operating
mode. See r.mode for defini-
tion of data.

specifies directional infor-
mation for display address
registers. See r.dir, r.stepx,
and r.stepy.

controls transmission of data

to central computer. If this
location contains 0, data is
transmitted to central computer;
if this location contains 3,

data is retained at terminal.
See r.input for more details.
m.enab specifies interrupt selection.

See r.ssf for more details.

r

Program

program is provided in the resident for use as a program diag-
It permits the keyset and display to be used as a computer
program, the RUN-CONSOLE switch must
The contents of the 8080 registers
display screen and an arrow
where directives may be entered from

their function are described below.

starting

to memory loca-

ion xXxXXxX ‘
16

Set breakpoint register

return to th “ONSOLE program.
the switch is in RUN position,
exit from CONSOLE mode will oc

to address

‘all program at address xxxx
called program must contain

(
instruction to return

routine).

step. Execute one instruction at
the present value of the program

rounter (PC).

program starting at
PC and continue
oint address is encoun
the 8080 register

g program exe

above, but do not uy
ay. The program wil

executed faster than above.

beginning at

for nn.
See
ine

6
assumed

3.4 Input-Output Addresses
Input and output operations are performed by the IN and OUT instruc-

tions and involve the transmission of 8-bit data words between the accu-
mulator (A) register and devices external to the 8080 processor. In some
‘ases the OUT instruction is used to set control flags in the Display
Interface Unit, and no data is actually transferred. In these cases the
contents of A are immaterial.

The format of the I/O address is shown in Figure 3.l.

ee... 5 3 1)

oe 3 YO ADDRESS

Figure 3.1. 1/0 Address Format

1

Addresses through 1F (hex) are used internally by the terminal

while 20 through 3F are used externally. INiy external addresses may

be ecified by the SSF instruction.
The address assignments are tabulated on the following pages. They

appear here for information purposes only; the user should not write

Zr

programs uSing these addresses, but instead should make use of the callable

Function

Input serial data byte

Dos Dor Dog Dos Dog Dox Do2 Do

Input communication port status

ee x x L x x x

| | L— |= lost RCV dato
ae Recieve Data Ready

|
L__ |= Tronsmitter Ready
D= Not Ready

Input interrupt vector and display

D xX x

L Display Type
D = Poralli! Interrupt Device Address
1 = Serial as Follows:

000 SIR
00 | KST
O10 TP
Ol} SOR
100 EXT p
101 EXT 1
110 CONSOLE
ae CARRIER

Input keyset word

E Kog Kos Koq Koz Koo Koy

Input touch panel word

[%os Xo2 Xo1 Xoo Yos Yo2 Yo: Yoo

en

Horizontal Vertical Position

Position of of Touch
Touch

Lype

Input Address
1eEMONnLC

: Function
(hex)

10 XL Input lower 8 bits of x register

Xo7 Xog Xos Xoa Xos “oe %o X00)

Input most significant bit (x
the x direction flag

x

|

ay = REVERSE
O = FORWARD

Input lower 8 bits

Input
the ba

flags

ABORT FLAG ee

| = ABORT
O = NORMAL

DIRECTION
REVERSE
FORWARD

Available for

Load the transmitter with the contents
of A and transmit to the central c mputer

Dor Dog Dos Dog Doz3 Dog Oa, Dag

Load the serial IO port with the following
Status word. This word conditions the

port to receive PLATO data.

interrupt mask

CN CRI

)
[ ae Carrier
Consoie

External
Device |

External
Device O

ig vector in the DIU
“ontents of A are unused. See OUT CLOCKL

iction for use of this flag.

he DIU to y.
are unused. See
of this flag.

Mnemonic

51

Function

OD

OF

CLRABT

XU

ae}

Set the y direction flag in the DIU. The
y register will be decremented by all sub-
sequent clock y signals. The contents of
A are unused.

Reset (CLEAR) the y direction flag in the
DIU. The y register will be incremented
by all subsequent clock y signals. The
contents of A are unused.

Set the ABORT flag. This instruction
Places the terminal in the ABORT mode.
The contents of A are unused.

Reset (CLEAR) the ABORT flag. This instruc-
tion places the terminal in the normal
operating mode. The contents of A are
unused.

Load the lower 8 bits of the x register
with the contents of A.

Load the most significant bit (x_) of the
x register with bit A. of the accumulator
The other bits of A ake unused.

Load the lower 8 bits of the y register.

Load the most significant bit (y.) of the
y register with bith A_ of the accumulator.
The other bits of A are unused.

Load the lower 8 bits of the panel parallel
lata register with the contents of A.

Load the upper 8 bits of the panel parallel
data register with the contents of A and
write (WE =1) or erase (WE,=4) the contents
of the parallel data veniales (16 bits)

on the display.

Load the DIU mode register with the lower

i SE wr? Oe ee WE
b
—_—_—_—_—-

Operating Mode ae

O = Serial
| = Parallel

Write /Erase Mode

Yutput Address

(hex)

Function

,

CLOCKX

CLOCKY

CLOCKXY

Load both PDL and PDU with the contents

of A and write (WE,=9) or erase (WE =)

the contents of the parallel data register

on the display.

NOTE: The write (erase) operation is
reversed in this operation.

Clock the x register. The x direction
flag specifies the direction; forward
if reset (f), reverse if set (1). The
contents of A are unused.

Clock the y register. The y direction
flag specifies the direction; forward
if reset (fg), reverse if set (1). The
contents of A are unused.

Clock both the x and y registers and

write (WE.=1) or erase (WE =) the resulting
address of the panel. The’ x and y direction
flags specify direction of change. The
contents of A are unused.

Clock the long vector, x or y, (as specified
by the long vector flag) in the DIU if

A_=$; clock both the x and y registers if
A’=1. The resulting address is then written
(WE _=1) or erased (WE,=) on the display.

The direction flags specify direction of

xy change.

Clock the y register and write (or erase)
the resulting address on the panel. The
contents of A, and the WE bits in the
Panel Mode register specify the operation
performed as follows:

WE FUNCTION
4

write
erase
nop

nop

erase
write
erase

write

‘lock the x register and write (or erase)
the resulting address on the panel as
shown in the table for OUT HCHAR instruction.

Mnemonic

WE Write if WE .=1 or erase if WE g
— x

address specified by the content

x and y registers. The contents

are unused.

SCREEN Erase the entire display. The

contents
of A are unused.

Load the lower 8 bits of the slide pro-

jJector register with the contents of A.

Xo3 Xo2 Xo: Xoo Yos Yor Yo: Yoo

“ x

—~p—

Shde X Shde Y
Address Address

Load the upper 2 bits of the slide |

jector with the lower 2 bits of A.

« L

Lamp J

Shutter ——~

information needed

external IO bus.

communicate with the terminal via the IO

& DATA (8)

de)
ADDRESS (6)

EXT WR

TERMINAL

EXT RD

WRO
EXxT®
EXT RDY

escribed below.

of bidirectional data.

address specifyin

equipment to receive

> ic i

(WRITE) address.

indicates that

shown

EXT RDY High level indicates that external
equipment specified by address bus can
perform requested operation. Addressed
device should take this line low if it
is unable to perform requested opera-
tion. ‘Terminal will then halt until
device is ready.

[IO Timing

The timing diagram for the input operation is shown in Figure 4.1.
All times are in nanoseconds (NS). No earlier than 90ONS after placing
the IO address on the address bus, the terminal issues the EXT RD signal.
The selected device places the data on the bus within 350NS after receipt
»f the EXT RD signal. If the device cannot supply data within this time,
it must take the EXT RDY line low within 50NS after receipt of the EXT RD

signal or within 140ONS after decoding the address.

10 ADDRESS

639 — 874

—
xT RDY
° 50 max

INPUT DATA

!
U
|
I
|
|
'
i
!
|
!
!
|

le— 350 max Senn!

Figure 4.1. Input Timing
terminal will then wait until the device has data ready. When ready,
ce places the data on the bus and raises the EXT RDY

lata must remain stable until the EXT RD is termi-

terminal.

The timing diagram for the output operation hown in Figure 4.2
~)4 ! ry |} <a Y } T - _- » *} 1. }
No earlier than 9ONS after placing the 1 addres: n 1 addr bus,

the terminal io, 1s the EXT WR signal indicating t t he elect
that data is forthcoming. The terminal places the data on the data bu
st least 200NS before issuing the WRO signal. The selected device can us

+} : = } . ie asi on : . -
nis ignal to read the data. If the selected devic: l iot ready t

receive data, it must take the EXT RDY line low within 5 NS after receipt
Ff siaqnal or within 140NS after decoding the iddré ‘ The

11 still issue the WRO signal, but will halt with thi signal

t
rt
3
~
0
—
2
~

low until the selected device raises the EXT RDY Signal. After receipt

-<
~

f the EXT RDY signal, the terminal will hold the WRO sign

tnals stable for SOONS before terminating the output operation.

+

10 ADDRESS

EXT WR

OUTPUT DATA

utput Tin

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
