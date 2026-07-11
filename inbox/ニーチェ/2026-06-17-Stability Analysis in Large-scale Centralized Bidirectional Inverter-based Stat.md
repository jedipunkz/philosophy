---
source: "https://arxiv.org/abs/2606.17565v1"
title: "Stability Analysis in Large-scale Centralized Bidirectional Inverter-based Stations Connected to Bulk Power Systems through AC and DC Connections"
author: "Qiang Fu, Wenjuan Du, Siqi Bu, Haifeng Wang"
year: "2026"
publication: "arXiv preprint / eess.SY"
download: "https://arxiv.org/pdf/2606.17565v1"
pdf: "https://arxiv.org/pdf/2606.17565v1"
captured_at: "2026-06-17T08:06:30Z"
updated_at: "2026-06-17T08:06:30Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche will to power"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Stability Analysis in Large-scale Centralized Bidirectional Inverter-based Stations Connected to Bulk Power Systems through AC and DC Connections

- 著者: Qiang Fu, Wenjuan Du, Siqi Bu, Haifeng Wang
- 年: 2026
- 掲載情報: arXiv preprint / eess.SY
- 情報源: [arxiv](https://arxiv.org/abs/2606.17565v1)
- ダウンロード: https://arxiv.org/pdf/2606.17565v1
- PDF: https://arxiv.org/pdf/2606.17565v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Massive controlled DC resources (CDCRs), such as battery energy storage systems, are connected to AC power systems through bidirectional inverters for power balance requirements. This study investigates converter-driven stability (CDS) issues in the sub-synchronous frequency range caused by large-scale bidirectional inverter-based stations (IBSs). The impacts of the AC and DC connections of IBSs on subsynchronous oscillations (SSOs) are compared by examining three factors: the number of CDCRs, power flow direction, and control parameters of the inverters. For AC connections, IBSs may induce instability as the number of CDCRs increases, regardless of the power flow direction. To maintain stability, the maximum power amplitude of the IBS is calculated. It is found that switching to DC connections can reduce these instability risks if the DC line resistance is much less than the AC line reactance. Moreover, the method of tuning control parameters is demonstrated to be more effective in improving power-related critical stability under DC connections. Therefore, The DC-IBS is preferred for high-voltage transmission. Finally, the conclusions are validated in power systems connected with both AC- and DC-IBSs under various network topologies and system scales.

## PDF Text

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

Abstract
—
Massive controlled DC resources (CDCRs), such as
battery energy storage systems,
are connected

to

AC power
systems
through

bidirectional inverters
for

power balance
requirements. This study investigates converterdriven stability
(CDS) issues in the subsynchronous frequency range caused by
largescale bidirectional inverterbased stations (
IBS
s).

The
impacts of the AC and DC connections of
IBS
s on subsynchronous
oscillations (SSOs) are compared by examining three factors: the
number of CDCRs, power flow direction, and control parameters
of the inverters.

For AC connections,
IBS
s may induce instability

as the number of CDCRs increases, regardless of the power flow
direction. To maintain stability, the maximum power amplitude of
the
IBS

is calculated. It is found that switching to DC connections
can reduce these instability risks
if

the DC
line resistance is much
less than
the AC line
reactance. Moreover, the method of tuning
control parameters is demonstrated to be more effective in
improving
powerrelated critical
stability under DC connections.
Therefore,
The
DCIBS

i s

preferred
for

highvoltage transm ission.
Finally, the conclusions are validated
in power
system s connected
with

both

ACand DCIBS
s under
various network topologies and
system scales.

Index Terms
—
C
onverterdriven
s tability
,
DC
power systems,
energy storage
system s (ESSs),

HV
DC
,

inverterbased

resource s
,
small signal stability,
subsynchronous oscillations

N
OMENCLATURE

C
k
:
DC capacitance of the
k th

inverter.

E
(
subscript
)
k
(
s
)
:
Transfer function of the
k th

ACIBS
,

expressed in
coordinates specified by the subscript.

I
:

Identity matrix.

I
s ubscript
:
Vector of powersystem currents (components
specified by the subscript).

I
dck
:

DC current
of the
k th

component
.

I
(s ubscript
)
k
:
AC current of the
k th

inverter
(
with coordinates
specified by the subscript
)
.

K
:

Inverter controller parameters.

L
:
Matrix of the AC line lengths.

P
(s ubscript
)
k
:

Active power
related to the
k th

component

specified
by the subscript.

Q
v k
:

Rea ctive power
output from the
k th

inverter.

V
s ubscript
: Vector of power system voltages (components
specified by the subscript)
.

V
dck
: DC voltage across the DC capacitor of the
k th

component
.

V
e
: AC voltage of the external bulk power system.

V
(s ubscript
)
k
: AC voltage

at the PCC of the
k th

inverter

(
with
coordinates specified by the subscript
)
.

V
v
(s ubscript
)
k
:
AC voltage at the terminal of the
k th

inverter

(
with
coordinates specified by the subscript
)
.

X
s ubscript
:
Reactance of the component specified by the
subscript.

Z
s ubscript
:

Network impedance of the power system specified by
the subscript.

θ
k
: Voltage

angle of the
k th

inverter.

ρ
k
:
k th

eigenvalue of
L
.

ω
s ubscript
:
Angular frequency of the coordinate s

specified by the
subscript.

ε
s
: SSO damping.

Δ
:
Variation in a variable.

Diag
(
·
):
Diagonal

matrix.

⊗
: Kronecker product.

Subscript
d
:

daxis
variables
(
in
dq

coordinate s
).

Subscript
q
:

qaxis
variables

(
in
dq

coordinate s
).

Subscript
x
:

xaxis
variables

(
in
xy

coordinate s
).

Subscript
y
:

yaxis
variables

(
in
xy

coordinate s
).

Subscript
dc
:

V
ariable s of the DC power system
.

Subscript
dq
:

V
ariable s

of AC power system
(
dq

coordinate s
).

Subscript
xy
: V
ariable s

of AC power system
(
xy

coordinate s
).

Subscript
v
:
Inverterside variable s
.

Subscript
L
: V
ariable s of

t ransmission

line.

Subscript
k
: Variables of the
k th

component,
k

∈

[1, N].

Subscript
j
: Variables of the
j th

component,
j

∈

[1, N+1].

Subscript
a
×
b
:
The matrix is composed of a rows and b
columns.

Superscript ref
:
Reference value.

PCC
:

Point of common coupling.

I.

I
NTRODUCTION

HE increas ed

penetration of renewable energy sources
(RESs) has led to fluctuations in active power output,
creating challenges for frequency regulation in power systems
[1]. To address this issue, a significant number of bidirectional
inverterbased resources (IBRs),
such as

energy storage systems
(ESSs), have been deployed in power networks [2]. Compared
to distributed integration, centralized inverter integration offers
improved economic efficiency and regulatory performance.
This has promoted the development of lar gescale centralized
bidirectional inverterbased stations (
IBS
s), typically located in
suburban areas and consisting of multiple converters within a
confined region. Consequently, the dynamic interactions among
these power electronics become more pronounced,
necessitating

a thorough investigation to prevent converterdriven stability

Qiang Fu,
Member, IEEE
, Wenjuan Du,
Member, IEEE
,
Siqi Bu,
Senior Member, IEEE
,
and

Haifeng Wang,
Senior Member, IEEE

Stability Analysis
in Largescale Centralized Bidirectional
Inverterbased Stations

Connected

to Bulk Power Systems
through AC and DC

Connections

T

This work is supported by the National Natural Science Foundation of China
(Grant

No.

52407129), and the Sichuan Science and Technology Program
(Grant No. 2025YFHZ0234). We also acknowledge the support from Hong
Kong Scholars Program. C
orresponding author: Wenjuan Du, Email:
ddwenjuan@qq.com.

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

(CDS) issues [3].

In
IBS
s,
controlled
DC resources

(CDCRs)
, typically
represented by battery energy storages (BESs)

[4]
-
[5]
,
are

connected to AC power systems either through AC (ACIBS
s)
or DC connections (DCIBS
s). In the case of DC connections
[
6
], DC transmission lines are
mainly
used to link
CDCRs

to
external bulk power systems.
The centralized bidirectional
inverter serves
for

the power conversion.

In this configuration,
dynamic interactions within the DC power system

are primarily
considered
. In contrast, for ACIBS
s,
distributed local inverters
directly perform power conversion, and the power is then
collected and
transmitted through AC transmission lines to the
external bulk power system

[7]. Dynamic interactions occur
among

the gridside inverters.

The configurations of
IBS
s with
these connections are shown in Fig. 1, and their effects on power
system stability are different and discussed below.

In
a
DCIBS
, a centralized bidirectional inverter regulates the
DC voltage,
and

the power output from CDCRs is maintained
at constant levels. The
stability

of simple DC power systems has
been extensively analyzed in [
8
]
-
[
10
]. When a CDCR draws
power from a DC system, it behaves like a constant power load
(CPL) with negative impedance. This negative impedance
intensifies as the power consumption increases, and if it
surpasses the positive impedance of the centralized inverter
, it
can lead to system oscillati ons. In contrast, when the CDCR
supplies power to the DC system, it functions similarly to a
constant power source (CPS). In this role, it exhibits positive
impedance that increases with the power output, promoting
stability rather than causing instability
.

However,
instabilities

in
such cases may arise from other factors, such as the control
parameters of the centralized inverter [
11
]
-
[12]
. Therefore, in
DCIBS
s
, the impact of bidirectional inverters on
stability

is
closely tied to power flow direction and control settings,
warranting further investigation particularly in largescale DCIBS
s
.

For
an
ACIBS
, each

CDCR

is
typically link ed

to a
gridside
inverter [1
3
], resulting in numerous bidirectional inverters in
a
largescale ACIBS
. The dynamics of these
gridside
inverters
are comparable to those of renewable energy sources when they
supply power to an external bulk power system. In such cases,
the primary focus is on the typical DC voltage control [1
4
].
Theoretically, the
instabilities

linked to RES inverters are
similar to those of
the
ACIBS
, including the synchronization
instability of t he phaselocked loop (PLL) under weak
connections [1
5
]
-
[1
6
] and resonances between inverters and
series compensation [
17
]
-
[18]
. However, practical scenarios
may differ from these theoretical considerations.

First, IBSs are
usually installed in suburban areas for power balancing services
rather than for long distance power delivery.

The connection
between
the ACIBS
s

and the external bulk power system is
typically strong and does not require the installation of series
compensation.
Therefore
,
the instabiliti es

related to weak
connections and series compensation in RES systems may not
be applicable to ACIBS
s. Instead,
the instabilities

caused by
the dynamic interactions among
massive
inverters under strong
connections require attention. Second, unlike RESs, which
primarily supply power to the connected power system, ACIBS
s are capable of both supplying and absorbing power. This
bidirectional capability

may

introduce potential
instabilities

that
require investigation
.

This study addresses these issues by considering bidirectional
power flows and different connection configurations of
IBS
s,
clarifying how the integration of largescale
IBS
s

introduces
instabilities in the
subsynchronous
frequency range,
specifically

subsynchronous oscillations

(SSOs).

The key contributions of
this study are as follows.

1)
A l argescale ACIBS

i s considered rather than singleequipmentbased cases, and a method is proposed to maintain
the

completed

dynamics related to
the
dominant SSO mode s
,
addressing highorder computational challenges. Based on this,
SSOs under strong connections are analyzed,
which differs from
the traditionally studied weak connections.

2)
The impact of dynamic interactions within ACand DCIBS
s on SSOs is analyzed. It is demonstrated that
even under
the strong grid connections
, the dynamic interactions amplify as
the scale of
IBS
s increases, which can lead to
growing SSOs
.
To mitigate these risks, it is crucial to limit the
scale

of ACand
DCIBS
s.

3)
The impacts of ACand DCIBS
s on SSOs are compared
under scenarios with bidirectional power flows. This study
highlights the differences between the two and suggests a
preference for DCIBS
s when highvoltage transmission lines
are used, particularly when the resistance of the DC line is
significantly lower than the reactance of the AC line.

The remainder of this paper is organized as follows: Section
II introduces the linearized models for
IBS
s with both AC and
DC connections. Section III presents a theoretical investigation,
proposing methods to analyze and mitigate
SSOs

caused

by
ACand DCIBS
s, followed by a comparison of

them
. Section IV
validates the mathematical models and theoretical findings
through case studies. Finally, Section V concludes the study by
summarizing the key insights.

II.

L
INEARIZED
M
ODEL
S

OF
ACAND
DCIBS
S

CONNECTED
TO
E
XTERNAL
B
ULK
P
OWER
S
YSTEMS

(a) ACIBS

(b)
D
CIBS

Fig. 1. Configuration of
the ACand DCIBS
s
.

Fig. 1 illustrates the configurations of the ACand DCIBS
s
based on

practical

centralized energy storage station s

in Jiangsu
Province, China.

I
n
the
ACIBS
,
N

CDCRs
are integrated
by

an
AC network and
then
connected to an external bulk power
system. Each
CDCR

corresponds to a
bidirectional inverter

for
©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

the power conversion
.

For
the
DCIBS
,
N

CDCRs

are integrated
through a DC network and

then

connected to an external bulk
power system via a centralized bidirectional inverter. In contrast
to ACIBS
s, DCIBS
s typically require
only one

bidirectional
inverter to
control the

DC voltage.

Taking

the Jiangsu
ACIBS

as

an example
, the rated

voltages for
station modules (SMs)
, AC
network, and
the external
bulk power system are 35
kV
, 220
kV
,
and 500
kV
, respectively. The
ACIBS

is
located less than 50
km from the external bulk power system,
and thus they are
under strong connections
.

Additional examples can be found on
the official website of Sungrow Power Supply Co., Ltd., such as
the implementation of DCIBSs [
19
].

A.

Transfer Function of
th e
CDCRs

Fig. 2. Configuration of the
CDCRk
.

The
constant

power control used by the DCDC converter

of
CDCRk

is shown in Fig. 2.
c b
k

is the
DC capacitance of the
k th

DCDC converter
,
l b
k

is the inductance between the battery and
the
DCDC converter
, and
v b
0

is the
DC voltage
of the battery
.

The transfer function of
CDCRk

can be
obtained

as (
detailed
derivations can be found in [2
0
])

(1)

where
H
dck
(
s
) denotes the output admittance of
CDCRk
.

It should be noted that dynamic of DCDC converter control
is normally faster than subsynchronous dynamics, thus in the
subsynchronous frequency range,
s

=
jω
sub
, yields

(2)

Given that (2) is extensively
used

in prior research

[2
1
]
-
[2
2
],
we do not delve into its details here.

B.

Transfer Function of
th e

Bidirectional Inverters

Fig. 3.
DC voltage control loop

of

the

b idirectional
i nverter

of the
SMk
.

B
ased on the control loops illustrated in Fig. 3,
the transfer
function of the bidirectional inverter is derived as (detailed
derivations can be found in
[2
3
]
-
[2
4
]
)

(3)

where

K
p k

and

K
i k

are the DC voltage o uter control loop
parameters
.

K
q k

and

K
qi k

are

the reactive power o uter control
loop
parameters
.

K
cp k

and
K
ci k

are the inner

control loop
parameters.
X
k

is the r eactance of ACside filter of the
k th

inverter.

ω
0

is the fundamental angular frequency.

F
or
a
DCIBS
,
the central ized
bidirectional
in verter regulates
the DC voltage, and its control configuration is similar to that
shown in Fig. 3. Therefore,
its

transfer function is obtained as
(detailed derivations can be found in [2
5
])

(4)

where
H
dc
(
N
+1)
(
s
) is output admittance of the
in verter,
and

.

C.

Model of the
Power System Connected with Either
A LargeScale
ACor DCIBS

The model of
DC

power
system can be obtained by directly
connecting the transfer functions of
the CDCRs,
central ized
bidirectional
in verter
, and

DC

network impedance matrix
because they are in the same coordinate, as shown in Fig. 4 (a).
The
generalized
DC impedance matrix
Z
dc

can be written as

(5)

From (1), (4), and (5),
the
obtained model
for
a
D
CIBS

is

(6)

Developing a model for a power system that incorporates ACIBS

requires a different methodology than that used in (6). This
is because the AC network impedance matrix is formed in the

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

(a) DC coordinates (b) AC
dq

and
xy

coordinates

Fig. 4.

Relationship between

the different coordinates.

xy

coordinates, whereas (3) is expressed in the

dq

coordinates.

Fig
.

4

(b) illustrates the correlation between these two
coordinate systems. To bridge this gap, the transfer functions
in

(3)
should be transferred

from
dq

to
xy

coordinate s

using

(7)

Thus, combining (3) and (7) yields (detailed derivations can
be found in
[2
6
]
)

(8)

The AC network impedance matrix can be represented as

(9)

where
I
xy
= [
I
x
1
,
I
y
1
, …,
I
x
N
,
I
y
N
]
T
,
V
xy
= [
V
x
1
,
V
y
1
, …,
V
x
N
,
V
y
N
]
T
,
and
there is no limitation on the topology of the AC network.

From (8) and (9),
the
obtained model
including

an
A
CIBS

is

(10)

To evaluate the stability of power systems connected with
DCand ACIBS
s, the eigenvalues of the system, determined
by solving (6) and (10), serve as key indicators. However, as the
number of CDCRs increases,
power
system models become
significantly more complex and higherorder. Given this
challenge, the focus of this
study

narrows to subsynchronous
frequency range.
In the following sections, a targeted method is
developed to efficiently identify and compute the
dominant

SSO
modes, enabling a precise stability assessment.

III.

A
NALYSIS OF AND
COMPARISON

BETWEEN THE IMPACT OF
AC

AND
DC

CONNECTIONS ON
SSO
S

A.

Modeequivalent Method for
A
Larges cale DCIBS

I
t has been demonstrated in [
25
] that the SSO mode in a DC
power system is primarily formed by the DC voltage control
converter and is affected by the
remaining

active powercontrolled converters.
Combining

the dynamic transfer function
of CDCRs listed in (2) with the DC network
impedance

matrix
listed in (5), the
input impedance of the remaining DC
power
system
, excluding the centralized inverter can be
represented

by
R
eq
, which is derived by

(11)

where
Z
1×
N

is
a

row vector consisting of the entries in row
N
+1
and columns 1 through
N

of
Z
dc
.

Z
N
×
1

is
a

column vector
consisting of the elements in column
N
+1 and rows 1 through N
of
Z
dc
. Z
N
×
N

is
a

submatrix of
Z
dc

formed by rows and columns
1 through N.
z
(
N
+1)(
N
+1)

is
an

element of
Z
dc

at position (
N
+1,
N
+1).

F
ocusing on the
central ized
bidirectional
in verter
, (6) can be
further expressed as

(1
2
)

The damping of
SSO mode (denoted as
λ
s

=

-
ε
s

+
jω
s
)

can
be
calculated from (4) and (11) as

(
Detailed derivation s

can be
found in [2
7
]
)

(1
3
)

B.

Impact of
DC Connections

on SSOs

Several factors influence the damping, but this study focuses
primarily on three key factors due to their significant roles in
SSO analysis:

1)
Power Amplitude and Direction.

As shown in (13), when
CDCRs
operate as

CPLs,
the negative impedances introduced
by

these CPLs
can cause
the
value

of
R
eq

to
become

negative,
thereby reducing the SSO damping. However
,

the value ofI
dc
(
N
+1)

is positive, which

counteracts this by improving the SSO
damping
.
I
n simple DC power systems, an analytical solution
for
R
eq

can be derived so that the
combined
impact
of
V
dc
(
N
+1)
R
eq
-
1
andI
dc
(
N
+1)

on the SSO damping can be numerically
evaluated

as the power amplitude and flow direction change
.
However,
obtaining an analytical solution for
R
eq

is complex and even
impossible for
a
largescale DCIBS
. It is challenging to
determine whether the combined impact of
V
dc
(
N
+1)
R
eq
-
1
andI
dc
(
N
+1)

on

the
SSO damping is positive or negative. Therefore,
the following derivations are conducted to assess how the power
amplitude and flow direction of
a
largescale DCIBS

affect
SSOs.

Let
i m

denote the current flowing through the
m th

branch line
(with resistance
r m
), and the power injected into the centralized
inverter is given by

(1
4
)

By linearizing (14), we get

(1
5
)

From (1
5
), the analytical solution
of
V
dc
(
N
+1)
R
eq
-
1
can be
expressed

as

(1
6
)

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

w here

.

The damping from (1
3
) can be rewritten as

(17)

I
t can be
concluded

from (17) that w hen the power flow is
from the CDCRs to the centralized inverter

(positive direction)
,
the DC currents are positive, with
α
N
+1

< 1
,
such that the SSO
damping is
reduced. Conversely, when the power flow reverses

(negative direction)
, the DC currents become negative,
with

α
N
+1

>

1,
the
damping continues to decrease as the power
amplitude increases.

However, under the same power amplitude,
the
damping is consistently better when power flow

direction

is
positive than that is

negative,

which is due to

(1
8
)

T
herefore,
for
a
largescale DCIBS
, it is confirmed that the
SSO damping decreases as the power amplitude increases.
Nevertheless,
the
damping is always improved when the power
flow direction reverses from negative to positive at the same
power amplitude. Consequently,
the
DCIBS

hold s

the stability
when SSOs under the negative power direction are
mitigated
.

2)
Control Parameters of the DC Voltage Control Loop.

I
t
can be seen from (13) that the control parameter
K
p
(
N
+1)
directly
affects the SSO damping,
independent of the power amplitude
and direction of the DCIBS
.
A higher

K
p
(
N
+1)

increases the
damping and enhances stability, while a lower
one

may cause
instability due to insufficient damping.

3) Number

of
CDCRs
.

Increasing the number of CDCRs
affects the SSO damping in two ways. First, as the number of
CDCRs increases, the total power amplitude of the DCIBRs
increases,
reducing
the SSO damping
,
as demonstrated by (1
7
).
Second, the addition of more branch lines from the integration
of additional CDCRs increases the amplitude of
κ
, which
implies that reversing the power flow direction creates a more
significant difference in SSO damping between the positive and
negative power flow directions.

C.

Modeequivalent Method for

A

Larges cale
A
CIBS

I
t is difficult to directly analyze the dynamics of an
ACIBS

because of the
numerous
inverters

with highorder calculations.

To
address this issue,
theoretical verifications are conducted to
remine the dominant
oscillation

mode,
λ
s
, of
ACIBS
,
ensuring
that the simplification does not affect the
SSO

analysis results.
the transfer function of the
power

system connected with
N

SM
s
can be derived as

(1
9
)

where
,
X
0

is the

AC transmission line r eactance per unit length
,
and
L

is a

m atrix composed of the lengths of
the AC
lines.

Based on similarity transfer theory [2
8
], Multiplying
P

and
P
–
1

to (1
9
), yields

(
20
)

where

P

and
P
–
1
are matrices composed of the left and right
eigenvectors of
L
, respectively
, and
the transfer function s

of
the
SM
s

are

identical, considering that they are from the same
manufacturer.

B
ased on (20), the power system connected
with a

largescale
ACIBS (including
N

SM
s
)

is

simplified into
N

subsystems,
each consisting of an

SM
(
E
xy k

(
s
)) and an
equivalent
impedance
(
ρ
k
X
0
).

T
he subsystem with the maximum
ρ
k

corresponds to the
dominant SSO mode, and the
transfer function model is

(
21
)

Based on (
21
),
the issue of
highorder calculation
is addressed
by focusing on the dominant SSO mode formed by aggregated
dynamics

of the largescale ACIBS
. T
he damping of
dominant

SSO
mode can be
calculated

by

(
22
)

where
E
xy k
(s)

is simplified by assuming
θ
k

≈
0

and
I
qk

= 0,
which
is valid under the highpower factor

condition
.

If the subsystem

in (22)

is stable, the
power

system is stable. Otherwise,
SSOs
occur in the power systems
.

D.

Impact of
AC Connection s on SSOs

Similar to the three key factors analyzed for
DC connections
,
these factors are also examined
for AC connections
.

1)
Power Amplitude and Direction.

Equation
(2
2
) indicates
that the SSO damping depends on the power amplitude
associated with
I
0
. This indicates that system stability is closely
linked to the active power output of
the
ACIBS
. When the
active power increases in the positive direction
,
the power flows
from the ACIBS

to the bulk power system
,
the value of
I
0

increases,

and the
SSO damping

decreases
. If the power flow is
reversed, changing the AC current from |
I
0
| to
-
|
I
0
|, the impact
on SSO damping remains unchanged as long as the amplitude
remains the same. Thus, the influence of
the
ACIBS

on SSOs
primarily depends on their
power
amplitude

rather

than their
direction.

2)
Control Parameters of the DC Voltage Control Loop.

I
f the
AC system is stable and has positive damping, increasing the
control parameter can further enhance the damping, leading to
better stability. However, it cannot improve the critical stability

related to the output AC
current. Regardless of the parameters
used, the
maximum value

of
the AC
current can always be
expressed as

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

(
23
)

Therefore, changing the parameter
K
pk

of the DC voltage
control loop
cannot increase the maximum
transmitted

power of
the ACIBS
, which is limited by the AC network
.

Increasing
K
pk

can make a stable AC system more stable but is
less
effective in
expanding the
power
scale of the ACIBS.

This conclusion
differs from that of the DCIBS, where
the scale can be extended
by improving the control parameter
K
p
(
N
+1)
.

3) Number

of
CDCRs
.

The transmitted power is typically
small for
a power system connected with an SM (an SM
includes a CDCR)
.

Instability occurs only when the AC line is
extremely long,
which
is theoretically correct but not practically
relevant. T
herefore
, the instability is analyzed under the shortdistance (strong) connections of
multiple
SMs
.

The
SSO damping

of
the

power system decreases when
ρ
max

increases
. If a

SM
is newly connected to the power system, the
impedance matrix changes from (1
9
) as

(
24
)

where
L
1
×
N

is a vector
with

one row and
N

columns,
L
N
×1

is a
transposition of
L
1
×
N
, and

L
N
+1

is a number that
corresponds

to
the connection line of the new
SM
.

We denote that the arrangement of eigenvalues of
Z
xy

is

(
25
)

Simllarly, the arrangement of eigenvalues of
Z
new

is

(
2
6
)

Accoring to “Cauchy's Interlace Theorem for Eigenvalues of
Hermitian Matrices” [
29
], yields

(2
7
)

Therefore, as the
number

of
CDCRs

increases, the dimension
of
the
impedance matrix
in (24)
increases, and the maximum
eigenvalue also increases. In this case, if
a
largescale ACIBS

output/absorb large power to/from the connected

bulk

power
system,
SSO

may occur even under strong connections.

E.

Stability Improvement Methods for
IBS
s.

To avoid
SSOs

caused by

largescale
IBS
s
, methods
for
improving SSO damping
are proposed as follows:

1)

Adjusting the Power Amplitude and Direction.

If the
number of CDCRs

and
the control parameters of the inverter s
are determined,
the SSO damping
can be improved by reducing
the power amplitude for
the
IBS
s.

The

s tability

of an A
CIBS

is maintained when the maximum
power amplitude of each
CDCR, i.e.,
P
dck
,

is limited
in the
following range

(2
8
)

F
or
the
D
CIBS
,
only the maximum power amplitude of the
centralized inverter under the
negative

power flow needs to be
limited by

(2
9
)

2)

Reducing the
Number of CDCRs
.

If the rated power
of the
CDCRs

and
the control parameters

of
the inverter s
have

been
determined,
the SSO damping

can be improved by reducing the
number of
CDCRs
.

F
or
the
A
CIBS
,
t he maximum number of
CDCRs

need to

comply with the following condition

(
30
)

F
or
the
D
CIBS
, the maximum
number

of
CDCRs can be
calculated

by (
14
)

and (29) as

(
31
)

3)

Tuning
the
Control Parameters.

If the above conditions
have been determined, the control parameters can be adjusted.

For
the
DCIBS
, increasing the proportional coefficient
K
p
(
N
+1)

c an improve stability,

the

parameter
should
satisf y

(
32
)

However,
changing the proportional coefficient
K
pk

is less
effective

in extending the power scale
, but it can improve the
damping ratio

of
the ACIBS
. If the
critical

damping ratio is
assumed to be
d lim
,
K
pk

should be limited by

(
33
)

F.

Compared Results between ACand DCIBS
s

It is demonstrated that many factors can result in instability
of AC and DC

power

systems, whether
a largescale
ACIBS

or
DCIBS

is
connected. However, the types of instability and their
causes differ from each other.
A summary

is provided in Table
I

to clarify these differences, ensuring accurate identification of
instabilities and announcing preferred scenarios for different
connections.

Additionally, Fig. 5 illustrates the flowchart for the
stability analysis
for connecting
largescale
ACand
D
CIBS
s.

The DCIBS

i s preferred when higher capacity of CDCRs can
be integrated under the condition of critical stability, i.e.,

(34)

where the condition
V
e

≈

V
0

=

V
dc
(
N
+1)

= 1

p.u.

is

considered.

Specifically, for the single transmission line configuration,
where the power of CDCRs is transmitted through a single DC
line with resistance
R
L

for
the
DCIBS
, or through a single AC
line with reactance
X
L

for
the
ACIBS
, (3
4
) can be specified as

(35)

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

From (35),

DC connections can provide a greater capacity for
CDCRs when the resistance of
the
DC line is significantly lower
than the reactance of the AC line. This condition is more likely
to be satisfied in highvoltage power transmission scenarios.
Therefore, considering that largescale centralized
IBS
s

are

typically

with highvoltage transmission
lines
, DC

connections
may
become more pronounced

in these cases.

Table I

A summary of compared results between ACand DCIBS
s

Terms

ACIBS

DCIBS

C
auses

Control interactions in
largescale
inverter
clusters
.

Negative damping
provided from CDCRs to

DC voltage

control loop.

Power
a mplitude and
d irection

Instability occurs under high power amplitude
conditions
.

Control
parameters

Increasing
K
pk

can
improve
damping ratio but
is lesseffective to extend
the scale of ACIBS
.

Increasing
K
p
(
N
+1)

can
improve stability and
extend the scale of DCIBS
.

Stability
improvement
methods

Cannot change the critical
stability
related to
the
output AC current but can
improve
damping
by
increasing
K
pk
.

Can change the critical
stability and can improve
positive damping by
increasing
K
p
(
N
+1)
.

Reducing the
maximum
eigenvalue of
AC
network
impedance

matrix.

Reducing the
resistances
of DC
network
impedance

matrix.

Reducing the maximum rated power

and number of
CDCRs.

Fig.
5
.

Flowchart

for the stability analysis of largescale DCand ACIBS
s.

IV.

C
ASE
S
TUDY

A.

Model Verifications

Fig.
6
.

A

bulk

power system connected with an
ACIBS
.

Fig
. 6
illustrates the setup
of the power system incorporating
an ACIBS
. The control schemes for the ACIBS

follow the
configurations
presented

in Figs. 2 and 3. To verify the results,
a

detailed model is employed for validation
, as given by (10).

The
p roportional and integral coefficients

of the DC voltage
control are 0.1 p.u. and 100 rad/s, respectively.
X
L
1

= 0.22 p.u.,

and

C
1

= 5 p.u.
The consistency of the results between the
established model and the
electromagnetic transient
(EMT)

model
for

DC connections has been verified in [
3
0
], and the
verifications for AC connections are conducted as follows.

Fig.
7
.

Trajectories of the
subsynchronous

osci ll ation mode
when

the output
power varies.

In Fig
.

7
, an SSO mode primarily influenced by the DC
voltage outer control loop is identified. The trajectory of this
mode is examined under varying output power conditions, with
P
dc
1

shifting from 0 p.u. to ±

1 p.u. The fullorder model
in (10)
and

the simplified model in
(2
1
) yield corresponding trajectories,
which are depicted in blue and red, respectively. At
P
dc
1

= 0 p.u.,
the computed results from both models are closely matched,
with values of
-
3.81 +
j
87.38 and
-
3.81 +
j
87.23 rad/s,
respectively, m arked by a black circle. As
P
dc
1

increases to 1
p.u., the
SSO
mode shifts rightward, with computed values of
-
3.75 +
j
84.38 and
-
3.52 +
j
83.85 rad/s, respectively. Conversely,
when
P
dc
1

decreases to
-
1 p.u., the values are
-
3.28 +
j
83.46 and
-
3.52 +
j
83.85 rad/s, respectively.

To further validate the accuracy of the established models,
two specific operational states are examined, including the fullorder mathematical model in (10), the simplified model in (21),
and the electromagnetic transient (EMT) model within the
FPGAbased

realtime simulation experiment platform from
Modeling Tech [3
1
]. The platform's configuration is depicted in
Fig. 8 (a), where the MT 6040 Simulator serves as an industrialgrade realtime simulator, featuring a robust Intel Xeon CPU
and Xilinx UltraScal e FPGA, while the MT 1070 RCP functions
as a rapid control prototype for power electronic system control.

The detailed results are presented in Table I
I

and Fig.
8
.

Table II indicates a strong agreement between the fullorder
and simplified models, with errors of 0.4 % and 0.2 % for output
power levels of 0.5 p.u. and
-
0.5 p.u., respectively.

To further
verify the accuracy,

Fig
.

8

extends the validation by comparing
©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

the timedomain results across different modeling approaches

based on
the experimental platform from Modeling Tech.

The
results demonstrate that the mathematical

model is accurate
,

thereby supporting its use in further analyses.

Therefore, the simplified model demonstrates a strong ability
to predict the trajectory of the
SSO

mode
at different steady
states
. The analysis consistently shows that the

SSO

damping is
reduced as the amplitude of the active power output from the
CDCR increases. This relationship suggests that
a
larger power
increase
, whether
in
positive or
in
negative, contribute s

to a
diminished stability.
This

highlights the critical role of power
amplitude in governing system stability.

Table I
I

Calculated results of fullorder and simplified models

Output power of
CDCRs

Models

Eigenvalue (rad/s)

0.5 p.u.

Fullorder model

-
3.86 ±
j
86.80

Simplified model

-
3.74 ±
j
86.44

-
0.5 p.u.

Fullorder model

-
3.63 ±
j
86.34

Simplified model

-
3.74 ±
j
86.44

(a)
C
onfiguration
of the e xperimental platform from Modeling Tech

(
b
) Output power of
CDCRs

is 0.5 p.u.

(
c
) Output power of
CDCRs

is
-
0.5 p.u.

Fig.
8
.

Simulation
platform and
results
of fullorder and
experimental

models
.

B.

Impact of
A Largescale
ACIBS

on SSOs

Fig.
9
.

The configuration of
a bulk power system connected with

an
ACIBS
.

The configuration of a bulk power system incorporating
multiple
SM
s
(each SM includes a CDCR)
via AC connections
is illustrated in Fig.
9
. In this configuration, four
SM
s are
connected to a common node as a group, and subsequently, four
such groups are integrated into a collection node through AC
lines with a reactance of 0.006 p.u. These are then connected to
the external power system via an AC line with a reactance of
0
.05 p.u. Each
S
M

is identical in terms of output power and
control parameters to ensure

consistency within the ACIBS
.
The rated power of each
CDCR

is specified as 1 p.u.,
and the

total output power of the ACIBS

is

16 p.u.

Fig.
10.
Trajectories of the
SSO

mode
as

the

number of CDCRs increases
.

Fig
.

10

illustrates the trajectories of
the dominant

SSO mode
as the number of CDCRs increases from
N

= 1 to 16. The SSO
mode consistently shifts to the right, irrespective of whether the
output power increases from 0 p.u. to 1 p.u. in either the positive
or negative direction. However, this movement is minimal when
the number of CDCRs is small under the
rated power conditions.
Notably, the damping of the SSO mode decreases significantly,
reaching a damping value of
-
0.38 rad/s when
N

= 16 and
P
dck

=
-
1 p.u
. This observation confirms that an increase in the number
of CDCRs reduces
the
system stability. At a fixed number of
CDCRs, the primary factor influencing oscillation damping is
the amplitude of the output active power. As the output power
increases, there is a consistent reduction in the damping, which
is particularly pronounced wh en a larger number of CDCRs is
involved. This suggests that the interaction among multiple
SM
s
can intensify oscillations, further highlighting the crucial role of
power amplit ude in maintaining system stability.

Table I
II

Calculated results with different number of CDCRs

Number
Damping
Oscillation
Number
Damping
Oscillation
©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

of CDCRs

ra t
io

mode (rad/s)

of CDCRs

ra t
io

mode (rad/s)

1

4.86

%

-
4.96 +
j
101.87

9

4.26

%

-
3.86 +
j
90.47

2

4.83

%

-
4.90 +
j
101.29

1
0

4.12

%

-
3.62 +
j
87.80

3

4.79

%

-
4.82 +
j
100.48

11

3.9
4
%

-
3.33 +
j
84.49

4

4.73

%

-
4.71 +
j
99.42

12

3.71 %

-
2.98 +
j
80.30

5

4.6
7
%

-
4.59 +
j
98.23

13

3.43 %

-
2.58 +
j
75.19

6

4.
60
%

-
4.45 +
j
96.73

14

3.05 %

-
2.08 +
j
68.25

7

4.
50
%

-
4.27 +
j
94.89

15

2.45 %

-
1.42 +
j
57.83

8

4.3
8
%

-
4.06 +
j
92.62

16

1.16 %

-
0.43 +
j
37.13

The detailed relationship between the number of CDCRs,
their corresponding damping ratios, and the SSO modes is
summarized in Table I
II, where the output power of each
CDCR

is
-
1 p.u
.
T
he damping ratio decreases steadily as the number of
CDCRs increases. For instance, with 12 CDCRs, the damping
ratio is 3.71

%, whereas with 16 CDCRs, it drops to 1.16

%. If
we set the critical damping ratio threshold at 3

%, it becomes
clear that the maximum number of CDCRs should be limited to
fewer than 15 to maintain syst em stability and prevent the onset
of excessive oscillations in
the
ACIBS
.

Fig.
1
1. Timedomain

simulation results for different numbers of CDCRs.

Fig. 1
1

provides timedomain simulation results for different
numbers

of CDCRs, where the active power of each CDCR is
fixed at
-
1 p.u

and a disturbance occurs at 0.1 s
. These results
highlight a key observation: when the number of CDCRs is 14,
the oscillations are well damped, and the system remains stable.
However, with 16 CDCRs, the system starts to experience
continuous oscillations. This further strengthens the conc lusion
that an excessive number of CDCRs can worsen system stability,
even when the grid

connection is
strong
.

From a practical perspective, these findings have significant
implications for the design and operation of
largescale
ACIBS.

They indicate that an excess of CDCRs might lead to reduced
damping and increased
SSOs
.
Therefore, it is essential to
carefully consider the number of CDCRs included in
the
ACIBS and ensure that this number is limited to a level that
prevents oscillatory instability.

C.

Impact of
A Largescale DCIBS on SSOs

Fig.
12
.

The configuration of
a bulk power system connected with
a
DCIBS
.

The configuration of a bulk power system connected with
multiple
SM
s
(each SM includes a CDCR)

via DC connections
is illustrated in Fig. 1
2
. The network topology
is the same as

that
of ACIBS
, with the DC
main

line resistance specified as
R
L

=
0.002 p.u. and the DC
branch
line

resistance as 0.0002 p.u. The
rated power of each
CDCR

is specified as 1 p.u., and all CDCRs
are identical in terms of output power and control parameters to
ensure consistency within the DCIBS
.

Fig.
13.
Trajectories of the
SSO

mode
as

the

power flow is reversed
.

In Fig.
13
, the number of CDCRs is maintained at
N

= 16, and
the output power of each CDCR varies from 0 p.u. to 1 p.u. in
both the positive and negative directions. It is observed that the
SSO mode shifts to the right as the power amplitude increases
in both directions. However, the SSO damping in the ne gative
direction consistently exhibits poorer performance than that in
the positive direction. Consequently, the results confirm that the
SSO damping decreases as the power amplitude increases. As
R
L

increases from 0.001 p.u. to 0.002 p.u., the SSO damping

enters the unstable region when output power
of
each CDCR
reaches
-
1 p.u., indicating the system instability. Therefore, for
the
DCIBS
, a smaller resistance is preferable for stability,
making
it

more suitable for highvoltage power transmission,
where the transmission resistance can be significantly reduced.

D.

Comparison

between ACand DCIBSs

under

Different
R
esistance s

of the

DC

T
ransmission
L
ine

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

Fig.
1
4
.

Trajectories of the
SSO

mode s under AC and DC connections as the
number of CDCRs increases (the resistance of the transmission line varies).

Fig. 1
4

illustrates the trajectories of the SSO modes with the
worst damping as the number of CDCRs increases. The output
power of each CDCR
is
-
1 p.u., and the SSO modes are plotted
for both ACand DCIBS
s, with the trajectories represented by
blue and red lines, respectively. As the number of CDCRs
increases from
N

= 1 to 16
, the SSO damping consistently
decreases
for

both connections. This observation is consistent
with previous findings, where an increase in the number of
CDCRs led to reduced damping, regardl ess of the connection
type.

Further analysis is conducted on the impact of varying
transmission line resistance values (
R
L

= 0.001 p.u. and 0.002
p.u.) on the system stability. Notably,
the resistance of
DC
transmission lines
is
typically much lower than
the reactance of
AC transmission lines, which can have significant implications
for the system performance. For
R
L

=

0.001

p.u., the SSO mode
at
N

= 16 for
the
DCIBS

is
-
17.76

+

j
84.55 rad/s, showing
improved damping compared with the ACIBS
, where the SSO
mode exhibits poorer stability.
However, for
R
L

=

0.002

p.u., the
SSO mode of
the
DCIBS shifts to 3.54

+

j
86.76 rad/s, reflecting
a deterioration in stability compared to that of the AC
connections.

This

indicates

that
the

advantage
of the
DCIBS

is
not always guaranteed, particularly when the transmission line
resistance increases significantly.

Fig.
1
5. Timedomain

simulation results
for different resistances of the DC
transmission line
.

Fig. 1
5

shows the timedomain simulation results that validate
th is conclusion.

The figure demonstrates that when
R
L
=
0.001
p.u.
, the DCIBS

offers improved stability, whereas for

R
L
=
0.002 p.u.
, the ACIBS

provides better stability.

This suggests
that the selection of ACor DCIBS should depend on the
specific attributes of the transmission line, particularly its
relative resistance and reactance. DC connections offer a clear
advantage when the resistance of the DC transmission line is
significantly lower than the reactance of the AC transmission
line. However, if the resistance in creases considerably, AC
connections may prove to be more stable.

E.

Comparison

between ACand DCIBSs
under Different
P
roportional
C
oefficient s

of
I
nverter s

Fig.
1
6
.

Trajectories of the SSO modes under AC and DC connections as the
number of CDCRs increases
(proportional coefficient of inverter varies).

Fig. 1
6

shows the impact

of
different

proportional coefficient s

of the DC voltage control loop on the SSO

damping. In this case,
the maximum number of CDCRs is maintained at
N

= 16
, and
the
DC
transmission line resistance is set
to

R
L
=

0.001

p.u. The
results show that increasing the proportional coefficient
improves the SSO damping for both the AC and DC connections.
For example, when
N

= 1, the initial SSO damping improves
from
-
18.77 rad/s to
-
37.62 rad/s as
K
p
(
N
+1)

of
the
DCIBS

increases from 0.5 p.u. to 1 p.u. Similarly, the damping
of

the
ACIBS

improves as
K
pk

increases, although the improvements
are less pronounced than those of the DCIBS
.

Notably, the SSO damping varie s

based on the connection
type when
N

=16. Specifically, with DC connections, there is a
marked improvement in SSO damping, increasing from 1.09
rad/s to
-
17.76 rad/s, which helps in reducing instability. In
contrast, with AC connections, no significant improvement in
damping is observed. Thi s clearly highlights that the critical
stability of the system is greatly enhanced with DC connections
by adjusting proportional coefficient

K
p
(
N
+1)
. Additionally, this
adjustment allows for the integration of a larger capacity of
CDCRs into the system wit hout affecting stability,
which

is not
easily achieved with AC connections.

The timedomain simulation results
in Fig. 17

further confirm
that increasing the proportional coefficient
K
p
(
N
+1)

of the DCIBS
can enhance critical stability, allowing a greater number of
CDCRs to be connected with in
the system stability.

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

Fig.
1
7
.

Timedomain

simulation results
for

d ifferent
p roportional
c oefficients
.

F.

Comparison between ACand DCIBS
s under Different
Network Topologies

Fig. 1
8
.

Network topologies of ACand DCIBS
s
.

As shown in Fig. 1
8
, three topologies are considered: radial,
ring, and series topologies. For the ACIBS
, the line impedance
between any two connected CDCRs is 0.007 p.u. For the DCIBS
, the line
resistance

between any two connected CDCRs is
0.0006 p.u. The centralized inverter of the DCIBS

is installed
close to the infinite bus
, as that in
Fig. 1
2
.

Fig. 1
9

presents the trajectories of the SSO modes with the
worst damping under different AC and DC network topologies
as the number of CDCRs increases. As
N

increases from 1 to
16, the SSO damping under
both

AC and DC

connections
consistently decreases across all topologies, confirming that
a
larger

integration of CDCRs reduces the damping, regardless of

Fig. 1
9
.

Trajectories of SSO modes under various network topologies as the
number of CDCRs increases
.

the topology.

However, the specific impact of topology on
stability varies. For ACand DCIBS
s with radial topology,
SSO damping is superior among the three. In contrast, the series
topology results in the poorest damping. Under the series
topology, both ACand DCIBS
s face instability risks, whereas
they remain stable in radial and ring topologies. The timedomain simulations in the first two figures of Fig.
20

support
these findings: sustained SSOs appear in the series topology,
whereas oscillations are effectively d amped in the radial and
ring topologies.

In a fixed network configuration, the choice between AC and
DC connections
relies

on a comparison of the DC line resistance
with the AC line reactance. For instance, in a ring topology, a
DC line
resistance

of 0.0006 p.u. results in better damping
(14.70 rad/s) than an AC connection (3.23 rad/s).
Conversely, if
the DC line
resistance

increases to 0.0012 p.u., the DCIBS
becomes

unstable, with the damping decreasing to
-
3.46 rad/s,
as shown by the dotted path in Fig. 1
9
, making the AC
connection more stable.

The timedomain analysis in the third
figure of Fig.
20

supports this finding: a higher DC line
resistance in the DCIBS

leads to
growing
SSOs. Thus, the
general conclusion across various topologies is that DCIBS

is

more

suitable for highvoltage transmission
,

whe re

the DC line
resistance is low compared to the AC line reactance.

Fig.
20
. Timedomain

simulation results
for

different

network topologies.

G.

Integration of

Hybrid
ACand DCIBS
s into

a Largescale
Power System

©

20
26

IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current
or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collec tive
wo rks, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works.

Fig.
2
1
.

New England power system wi

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
