---
source: "https://arxiv.org/abs/1803.01211v2"
title: "Robust Power Flow and Three-Phase Power Flow Analyses"
author: "Amritanshu Pandey, Marko Jereminov, Martin R. Wagner, David M. Bromberg, Gabriela Hug, Larry Pileggi"
year: "2018"
publication: "arXiv preprint / eess.SP"
download: "https://arxiv.org/pdf/1803.01211v2"
pdf: "https://arxiv.org/pdf/1803.01211v2"
captured_at: "2026-05-09T12:04:00Z"
updated_at: "2026-05-09T12:04:00Z"
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

# Robust Power Flow and Three-Phase Power Flow Analyses

- 著者: Amritanshu Pandey, Marko Jereminov, Martin R. Wagner, David M. Bromberg, Gabriela Hug, Larry Pileggi
- 年: 2018
- 掲載情報: arXiv preprint / eess.SP
- 情報源: [arxiv](https://arxiv.org/abs/1803.01211v2)
- ダウンロード: https://arxiv.org/pdf/1803.01211v2
- PDF: https://arxiv.org/pdf/1803.01211v2

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Robust simulation is essential for reliable operation and planning of transmission and distribution power grids. At present, disparate methods exist for steady-state analysis of the transmission (power flow) and distribution power grid (three-phase power flow). Due to the non-linear nature of the problem, it is difficult for alternating current (AC) power flow and three-phase power flow analyses to ensure convergence to the correct physical solution, particularly from arbitrary initial conditions, or when evaluating a change (e.g. contingency) in the grid. In this paper, we describe our equivalent circuit formulation approach with current and voltage variables that models both the positive sequence network of the transmission grid and three-phase network of the distribution grid without loss of generality. The proposed circuit models and formalism enable the extension and application of circuit simulation techniques to solve for the steady-state solution with excellent robustness of convergence. Examples for positive sequence transmission and three-phase distribution systems, including actual 75k+ nodes Eastern Interconnection transmission test cases and 8k+ nodes taxonomy distribution test cases, are solved from arbitrary initial guesses to demonstrate the efficacy of our approach.

## PDF Text

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

1

Abstract
—
Robust

simulation

is essential for reliable operation
and

planning of

transmission and distribution
power
grid s
. At
present, disparate methods exist for steadystate analysis of the
transmission (power flow) and distribution power grid (threephase power flow).
Due to the nonlinear nature of the problem, it
is difficult for
alternating current (AC)

power flow and

threephase power flow analyses to ensure convergence to the correct
physical solution
, particularly

from arbitrary initial
conditions, or
when evaluating a change (e.g. contingency) in the grid
. In this
paper
,

we
descr ibe our

equivalent circuit formulati on
approach

with current and voltage variables that

model s

both the positive
sequence
network
of the transmission grid and threephase
network
of the distribution grid without loss of generality.
Th e
proposed

circuit models and formalism enable the extensi on and
application of

circuit simulation techniques

to

solve for the steadystate
solution with excellent robustness of convergence
. Example s
for

positive sequence transmission and threephase distribution
systems, including
actual

75k+ nodes Eastern Inter connection

tran s
mission

test cases and 8k+ nodes taxonomy distribution test
cases, are solved from arbitrary initial guesses to demonstrate the
efficacy

of our approach.

Index Terms
—

circuit simulation methods, continuation
methods, convergence problems, equivalent circuit approach,
power flow,
robust convergence,
steadystate analysis,
threephase
power flow, Tx
s tepping method

I.
I
NTRODUCTION

n interconnected
electric grid

is a networ k of
synchronized

power
providers and consumers that are
connected via transmission and distribution lines and
operated by one of multiple entities.
Reliable and secure

operation of this electric grid is of

utmost importance for

maintaining

a country’s eco nomy and wellbeing of its citizens
.

To operate the grid
reliably and securely

under all
conditions
,

as well as adequately
plan

for the future
, it is
essential

that one can
robustly
analyze
the grid offline and in
realtime.

At present
,

numerous analysis methods exist for
operation and planning of the grid
. These

can be broadly
categorized into one of the following: i) steadystate analysis

in
the
frequency domain

(power flow
,

three phase
power flow, and
harmonic analyse s), ii)
transient

and steadystate

analysis

in
time domain
, iii)
analysis for
optimal
dispatch of resources
, and
iv) oth er market
dispatchbased

analyses
. Among these
analyses,

fundamental frequency based steadystate analysis

(power flow and threephase power flow)

is ess ential for the

This work was supported in part by the Defense Advanced Research Projects Agency (DARPA) under award no. FA8750
-
17
-
1
-
0059 for the RADICS
program.

1
Authors are with the Electrical and Computer Engineering Department, Carnegie Mellon University, Pittsburgh, PA 15213 USA, (email: {amritanp,
mjeremin,
mwagner1, dbromber,
pileggi}@andrew.cmu.edu)
.

2
Author is

with the
Power System Laboratory, ET
H Zurich
, (email:

hug@eeh.ee.ethz.ch
).

Digital Object Identifier 10.1109/TPWRS.2018.2863042

daytoday operation as well as future planning

of the grid
.

Furthermore,

the solution to the steadystate analysis serves as
the initial state for transient

analysis

as well as
the
optimal
power flow problem
.
Due to
its
critical importance,

research
has produced

significant advances

toward improving the
convergence of these solution
methods

[5]
-
[12]

.

At present, steadystate simulation is divided into
two
domains,
highvolta ge
transmission systems and
substation
level voltage
distribution systems. Disparate methods exist

for
analyzing
these two

(
transmission and distribution
)

systems.
The steadystate solution for the high voltage transmission
system is obtained via positive

sequence

AC

power flow
analysis

(
often
referred to as power flow analysis),

whereas the
steadystate operating point for the distribution system

is
obtained via threephase AC power flow analysis. The industry
standard for solving the
positive sequence AC

power flow
problem is

the

‘PQV’ formulation
[1]
, wherein nonlinear power
mismatch equations are solved for bus voltage
magnitudes and
angles

that further define the steadystate operating point of the
system.
In contrast
,
the

backwardforw ard sweep method
[2]

and

the
current injection method (CIM)
[3]

are
primarily

used
for obtaining the steadystate solution of the

threephase power
flow

problem for the distribution grid.

In

their

existing form s,

the

solution methods for

power flow
and threephase power

can suffer

from lack of

convergence
robustness

[5]
,

[10]
. The ‘PQV’ based formulation

for the
positive sequence power flow problem

is known to diverge or
converge to nonphysical

solutions

for illconditioned

[2]

and
large scale (>50k buses) systems

[20]
, where a

nonphysical
solution

corresponds to a s ystem
that contains

low voltages or
demonstrate s

angular instability.

For distribution system
problems
,
the
backwardforward sweep method

that was
proposed to solve radial and weakly meshed distribution
systems with high R/X ratio
[2]

has difficulties converging for

heavily meshed

systems

with more than a single source

of
generation

[12]
.

The

CIM method based on Dommel’s work in
1970
[4]
, like the equivalent circuit ap proach proposed in this
paper,

represents the currents and voltages in terms of
rectangular coordinates, but is challenged by incorporation of

multiple PV buses in the system

[13]
-
[14]
.

In general
, of the
known
challenges associated with convergence for

existing

power flow and threephase powe r flow
solution methods
, the
Robust
Power Flow and ThreePhase Power
Flow Analyses

Amritanshu Pandey
1
,
Graduate

Student Member,

IEEE
, Marko Jereminov
1
,
Graduate

Student Member,

IEEE
,
Martin R. Wagner
1
,
Graduate

Student Member,

IEEE
,
David M. Bromberg
1
,
Member,

IEEE
,
Gabriela Hug
2
,
Senior
Member
,

IEEE

and Larry Pileggi
1
,
Fellow
, IEEE

A

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

2

two that are most detrimental are

convergence to
lowvoltage

or unacceptable

solution s

[15]

and divergence
[5]
.

The objective

and contribution

of the approaches described
in this paper is to provide

robust

power flow and threephase
power flow
convergen ce
.
Specifically,
a

generalized

approach
for power flow and threephase power flow
analyses

that can
ensure convergence to correct physical solution
independent of
the choice of initial conditions.

The factors that are the most fundamental toward making
th ese problems challenging
are
the use of nonphysical

representations for modeling the power grid components
,

and
in

the case of
the
‘PQV’ formulation
,
the use of

inherently nonlinear

power mismatch equations to formulate the problem
.

The
nonphysical

representations of the system equipment may not
capture the true behavior of the model in the entire range of
system operation. For
example, an

approximated
macromodel
of a generator that is represented via positive sequence or threephase PV
node
can re sult in convergence to a
lowvoltage
solution or divergence

due to its quadratic voltage
characteristics
. Similarly,
the inherent
nonlinearities in the
‘PQV’ formulation almost always cause divergence for large
(>50k) and illconditioned test cases

[20]

when solved
from an
arbitrary

set of initial conditions. This lack of a physicsbased
formulation
,

along with

the

methods that can constrain the nonphysics based models in their physical space
,

is what render s
the existing power flow

and threephase power flow problem
and solution approaches to be “nonrobust
.
”

To develop a robust solver for the steadystate
solution of
the
power
grid, it is imperati ve that the solver can efficiently
and effectively navigate through
the

aforementioned

challenges
while converging to a solution that is both meaningful and
correct.
Intuit ively

and physically, both the transmission and
distribution electric grid s

correspond to an electric circuit.
O
ur
approach toward solving the power flow and threephase power
flow problem s is to utilize circuit modeling and formalism to
develop new algorithms that will robustly solve them
.
Toward
this
goal,
we propose a twopron ged approach
. First,

the use of
an
equivalent circuit formulation
in terms of

the
true state
variables of currents and voltages
[16]
-
[18]

to model both the
transmission and distribution power grid (
Sect.
III.
). Secondly,
the
adaptation
and application
of circuit simulation methods

[19]
-
[22]

to ensure robust convergence to correct physical
solutions

(
Sect.
IV.
)

for power flow and threephase power flow
problems
.

To demonstrate the interaction between the two,
Sect. V of this paper discusses the general algorithm for solving
the power flow and threephase power flow problems.

Several
examples

are shown which demonstrate the efficacy of our
approach.

II.
B
ACKGROUND

A p ower grid in its simplest form can be represented by a set
of
𝒩

buses
,

where
the sets

of generators

𝒢

and load demands
ℒ

are subsets of

𝒩
, which are further connected by a set of

line
elements,
𝒯
X

and set of transformers

𝑥𝑓𝑚𝑟𝑠
.

Furthermore
,

there
is a set of slack buses

(one for each island in the system)

represented by
ξ
. In addition to these, the power grid may
contain other

elements
,
such as shunts,
flexible alternating
current
transmission system (FACTS), etc. The
objective

of
steadystate analysis of the power grid is to model the
fundamental frequency component of

the

power grid and solve
for the complex voltages at its buses. The high voltage
transmission network of t he grid generally operates under
balanced conditions
,

and therefore
,

the steadystate solution of
the transmission network
is

obtained via

positive sequence
power flow model and analysis. In contrast, the distribution
network of the power grid can operate
under unbalanced
conditions
,

and
therefore

we must
apply

threephase power
flow model and analysis to find the steadystate solution

of the
distribution grid
. In the following subsections, we discuss the
current formulations

used for steadystate analysis

of
transmission and distribution network s and highlight their
limitations
.

A.
‘PQV’ based Formulation for Positive Sequence Power
Flow Problem

The
‘
PQV’ based power flow formulation is
the industry
standard

for solving

for

the steadystate

solution of the high
voltage transmission network. In this formulation, a set of
2
(
𝒩
−
ξ
)
−

𝒢

power mismatch equations are solved for
unknown complex voltage magnitudes and angles of the system
using the Newton Raphson (NR) method. The set of power
mismatch
equations are defined

as follows:

𝑃
𝐺
𝑖
−

𝑃
𝐿
𝑖
=
|
𝑉
𝑖
|
∑
|
𝑉
𝑙
|
(
𝐺
𝑖𝑙
𝑌
𝑐𝑜𝑠

𝑖𝑙
+
𝐵
𝑖𝑙
𝑌
𝑠𝑖𝑛

𝑖𝑙
)
𝒩
𝑙
=
1

(
1
)

𝑄
𝐺
𝑖
−

𝑄
𝐿
𝑖
=
|
𝑉
𝑖
|
∑
|
𝑉
𝑙
|
(
𝐺
𝑖𝑙
𝑌
𝑠𝑖𝑛

𝑖𝑙
−
𝐵
𝑖𝑙
𝑌
𝑐𝑜𝑠

𝑖𝑙
)
𝒩
𝑙
=
1

(
2
)

where,
𝑃
𝐺
𝑖
+
𝑗
𝑄
𝐺
𝑖

and
𝑃
𝐿
𝑖
+
𝑗
𝑄
𝐿
𝑖

are the complex generation and
complex load at the node
𝑖

and
𝐺
𝑖𝑙
𝑌
+
𝑗
𝐵
𝑖𝑙
𝑌

is the complex
admittance
between
the nodes
𝑖

and
𝑙
.

In order to
solve for

unknown complex voltages
𝑉
𝑖

∠

𝑖

in
the system, the real and reactive power mismatch equations
given by

(
1
)
-
(
2
)

are solved for the set of
(
𝒩
−
ξ
−
𝒢

)

buses in
the system
,

whereas only real mismatch equations
(
1
)

are
solved for the set of buses with generators
𝒢

connected to it.

Importantly, this
‘PQV’ formulation is inherently nonlinear
,

since

the set of network constraints
result in nonlinear

pow er
mismatch equations independent of physics of the models used.
For example, in
the
‘PQV’ formulation
,

a
linear network
consisting of linear models for

the

slack bus,

the

transmission
lines and
the
loads

would correspond to a nonlinear set of
power misma tch
equations, a feature that could result in
convergence difficulties for systems even trivial in size.

B.
Current Injection Method for ThreePhase Power Flow
Problem

Until
recently
, the

backward forward sweep method was the
most commonly used method for the

steadystate analysis of the

radial and weakly meshed

distribution system s

[2]
. The method
was
preferred

over
the
‘PQV’ formulation due to the
radial
nature of the distribution grid and high R/X ratio s

of the
distribution lines, both of which
are known to cause

convergence difficulties for
the NR method

[2]

with ‘PQV’
formulation
. However, the backward forward sweep method
itself is pr one to convergence difficulties

for
systems

that are
highly meshed or have
multiple sources

[12]
.

The c urrent

injection method (CIM) for the
threephase

power flow problem
[3]

was proposed to addre ss challenges
IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

3

associated with
the
‘PQV’ formulation and the backwardforward sweep method
. In the CIM formulation, the nonlinear
current mismatch equations for the system buses are solved via
NR method for each individual phase with complex rectangular
real and imaginary voltages

(
𝑉
𝑅𝑖
𝛺
+

𝑗𝑉
𝐼𝑖
𝛺
)

as the unknown
variables. The current mismatch equations for the threephase
power flow problem are defined as follows
[3]
:

𝛥
𝐼
𝑅𝑖
𝛺
=
(
𝑃
𝑖
𝑠𝑝
)
𝛺
𝑉
𝑅𝑖
𝛺
+
(
𝑄
𝑖
𝑠𝑝
)
𝛺
𝑉
𝐼𝑖
𝛺
(
𝑉
𝑅𝑖
𝛺
)
2
+

(
𝑉
𝐼𝑖
𝛺
)
2
−
∑
∑
(
𝐺
𝑖𝑙
𝛺𝑡
𝑉
𝑅𝑖
𝑡
−
𝐵
𝑖𝑙
𝛺𝑡
𝑉
𝐼𝑖
𝑡
)
𝑡𝜖
𝛺
𝑠𝑒𝑡
𝒩
𝑙
=
1

(
3
)

𝛥
𝐼
𝐼𝑖
𝛺
=
(
𝑃
𝑖
𝑠𝑝
)
𝛺
𝑉
𝐼𝑖
𝛺
−
(
𝑄
𝑖
𝑠𝑝
)
𝛺
𝑉
𝑅𝑖
𝛺
(
𝑉
𝑅𝑖
𝛺
)
2
+

(
𝑉
𝐼𝑖
𝛺
)
2
−
∑
∑
(
𝐺
𝑖𝑙
𝛺𝑡
𝑉
𝐼𝑖
𝑡
−
𝐵
𝑖𝑙
𝛺𝑡
𝑉
𝑅𝑖
𝑡
)
𝑡𝜖
𝛺
𝑠𝑒𝑡
𝒩
𝑙
=
1

(
4
)

where
𝛥
𝐼
𝑅𝑖
𝛺
+
𝑗
𝛥
𝐼
𝐼𝑖
𝛺

is the net current mismatch
in
phase
𝛺

at
node
𝑖

and
(
𝑃
𝑖
𝑠𝑝
)
𝛺
+
𝑗
(
𝑄
𝑖
𝑠𝑝
)
𝛺

is the specified complex power
injection
at node
𝑖
.

The set
Ω
set

includes
phases a
, b
and

c.

Although,
the
CIM method
is known to improve the
convergence properties for heavily and weakly meshed threephase radial

distribution systems

with

high R/X ratio, the
method
is known to diverge for

testcases with high penet ration
of PV buses
[13]
.
Traditionally
, the number of PV buses in the
distribution system were limited to a small number
; h owever,
with the advent
of large scale

installation of

distributed
generation

(DGs)

and voltage control devices

in the distribution
system
this is no longer true. Therefore, it
is essential that
a

robust
threephase power flow formulation can

robustly handle

high penetration of PV buses

and

other

voltage control devices
in the syste m.

III.
E
QUIVALENT
C
IRCUIT
F
ORMULATION

We
extend
the equivalent circuit approach in
[16]
-
[20]

for
steadystate analysis of the transmission and distribution power
grid to
tackle the challenges exhibited by the existing
formulations. This
approach for generalized modeling of the
power system in steadystate (i.e. power flow and threephase
power flow) represents both the transmission and distribution
power grid
elements
in t erms of equivalent circuit elements

without loss of generality
. It was shown that each of the power
system components can be directly mapped to an equivalent
circuit model based on the underlying relationship between
current and voltage state vari ables
. Im portantly, this
formulation can represent any physics based model or
measurement based semiempirical models as a subcircuit
, as
shown in
[24]
,
[25]

and
[26]
,

and these models

can be combined
hierarchically with other circuit abstr actions to build larger
aggregated models
. In the following section
,

we discuss
generic
equivalent circuit representations

of

power system components
for both the positive sequence power flow problem and the
threephase power
flow problem. Note that throu ghout the
paper, the symbol superscript
Ω

in the mathematical
expressions of the equivalent circuit models
represents a phase
from
the
set
Ω
𝑠𝑒𝑡

of three

phases a, b
and

c

for
the
threephase
problem and
represents the
positive sequence

(
p
) component

for
the power flow problem
.

A.
PV Bus or the Generator Model

In
the
equivalent circuit approach, the generator (PV) bus
model is modeled via a complex current source
[19]

and has the
same behavior as of the PV node in power flow and threephase
power flow problems
. To
enable the application of

NR, th is

complex current source is split into

real and imaginary current

sources

(
𝐼
𝑅𝐺
Ω

and

𝐼
𝐼𝐺
Ω
, respectively)
. This is n ecessary

due to the
nonanalyticity of complex conjugate function s

[16]
. The
resulting equations for the PV model in
the
power flow and
threephase power problem a re:

𝐼
𝑅𝐺
Ω
=
𝑃
𝐺
Ω
𝑉
𝑅𝐺
Ω
+
𝑄
𝐺
Ω
𝑉
𝐼𝐺
Ω
(
𝑉
𝑅𝐺
Ω
)
2
+
(
𝑉
𝐼𝐺
Ω
)
2

(
5
)

𝐼
𝐼𝐺
Ω
=
𝑃
𝐺
Ω
𝑉
𝐼𝐺
Ω
−
𝑄
𝐺
Ω
𝑉
𝑅𝐺
Ω
(
𝑉
𝑅𝐺
Ω
)
2
+
(
𝑉
𝐼𝐺
Ω
)
2

(
6
)

Additional constraints that allow the generators to control the
voltage magnitude either at its own node or any other remote
node in the system are modeled by a control circuits, as shown
in the following subsection. In the case of power flow problem,
a si ngle control circuit is needed whereas for the threephase
power flow problem, three such control circuits are
needed for
each PV bus in the system
. The reactive power
𝑄
𝐺
Ω

of the
generator acts as the additional unknown variable for the
additional constr aint that is introduced due to voltage control.

In case of threephase power flow,
three such
additional
variables and constraints

are introduced.

As an
example, the equivalent circuit for the

positive
sequence model for

a
PV bus
used in power flow
is shown in
Fig. 1

for
the
𝑘
+
1
𝑡
ℎ

iteration of NR
. It is constructed
by

linearizing the set of equations
(
5
)
-
(
6
)

for the positive sequence
parameters

and then representing t he
resulting
equations

using
fundamental circuit elements

(detailed procedure for this
provided in
[16]
)
. To

construct the PV bus equivalent circuit for
threephase power flow p roblem, three such circuits are first
constructed and then are connected in groundedwye
configuration.

Figure
1
: Equivalent Circuit Model for PV generator model.

B.
Voltage Regulation of
a
Bus

Numerous power grid elements such as generators, FACTS
devices, transformers, shunts etc., can control a voltage
magnitude at a given node in the system. Moreover, they can
control the voltage magnitude
(
𝑉
𝑠𝑒𝑡
Ω
)

at either their own node
𝒪

or a remote node
𝒲

in the system. In equivalent circuit
formulation, we represent
the control of
the
voltage magnitude
by a control circuit
(
Fig. 2
),

which is

governed

by

the

following
expression:

𝐹
𝒲
Ω
≡
(
𝑉
𝑠𝑒𝑡
Ω
)
2
−
(
𝑉
𝑅𝒲

Ω
)
2
−
(
𝑉
𝐼𝒲

Ω
)
2
=
0

(
7
)

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

4

The circuit in
Fig. 2

is derived from the linearized version
of
(
7
)
. For the power flow problem,

it is
stamped

(
i.e. values are
added to the Jacobian

in a modular
way
)

for each node
𝒲

i n
the system whose voltage is being controlled such that there
exists at least one single path between the node
𝒲

and the
equipment’s

node
𝒪

that is controlling it. Similarly, for threephase power flow three of these circuits are stamped for each
node
𝒲
. The additional unknown variables for these additional
constraints are dependent on the power system device that is
controlling the voltage magnitude. For example, the additional
unknown variable for a generator is its reactive power

𝑄
Ω
,
whereas in the
case of transformers, it is the transformer tap
𝑡𝑟
Ω
, and for FACTS devices it
is

the firing angle

𝜑
Ω
. The
previous section
already
described how the additional unknown
variable

𝑄
Ω

for PV buses is integrated in to

the respective
equivalent circuits for generators.

Figure
2
: Voltage magnitude constraint control equivalent circuit
.

C.
ZIP Load Model

In this section, we derive the
positive sequence and threephase
model for
the
ZIP load. The
ZIP load models

the
aggregated load in the system as a mix of constant impedance

(
𝑍

𝛺
+
𝑗
𝑍
𝑄
𝛺
)
, constant current

(
𝐼

𝛺
+
𝑗
𝐼
𝑄
𝛺
)
,
and constant
power

(
𝑆

𝛺
+
𝑗
𝑆
𝑄
𝛺
)

load

behavior
, which can be

mathematically
represented as follows
:

(
𝑃
𝑖
𝑍𝐼
)
𝛺
=
𝑍

𝛺
(
|
𝑉
𝑖
𝛺
|
)
2
+
𝐼

𝛺
(
|
𝑉
𝑖
𝛺
|
)
+
𝑆

𝛺

(
8
)

(
𝑄
𝑖
𝑍𝐼
)
𝛺
=
𝑍
𝑄
𝛺
(
|
𝑉
𝑖
𝛺
|
)
2
+

𝐼
𝑄
𝛺
(
|
𝑉
𝑖
𝛺
|
)
+
𝑆
𝑄
𝛺

(
9
)

In the equivalent circuit approach, the equations for the ZIP
load model can be rewritten as:

(
𝐼
𝑅𝑖
𝑍𝐼
)
𝛺
=

𝑌

𝛺
𝑉
𝑅𝑖
𝛺
−
𝑌
𝑄
𝛺
𝑉
𝐼𝑖
𝛺
+

𝑆

𝛺
𝑉
𝑅𝑖
𝛺
+

𝑆
𝑄
𝛺
𝑉
𝐼𝑖
𝛺
(
𝑉
𝑅𝑖
𝛺
)
2
+

(
𝑉
𝐼𝑖
𝛺
)
2

(
10
)

+

(
√
𝐼

𝛺
2
+

𝐼
𝑄
𝛺
2
)
.
𝑐𝑜𝑠

(

𝑖
𝛺
−
𝐼
𝑝𝑓
𝛺
)

(
𝐼
𝐼𝑖
𝑍𝐼
)
𝛺
=

𝑌

𝛺
𝑉
𝐼𝑖
𝛺
+
𝑌
𝑄
𝛺
𝑉
𝑅𝑖
𝛺
+

𝑆

𝛺
𝑉
𝐼𝑖
𝛺
−

𝑆
𝑄
𝛺
𝑉
𝑅𝑖
𝛺
(
𝑉
𝑅𝑖
𝛺
)
2
+

(
𝑉
𝐼𝑖
𝛺
)
2

(
11
)

+

(
√
𝐼

𝛺
2
+

𝐼
𝑄
𝛺
2
)
.
𝑠𝑖𝑛

(

𝑖
𝛺
−
𝐼
𝑝𝑓
𝛺
)

where:

𝐼
𝑝𝑓
𝛺
=

𝑡𝑎𝑛
−
1
(
𝐼
𝑄
𝛺
𝐼

𝛺
)

(
12
)


𝑖
𝛺
=

𝑡𝑎𝑛
−
1
(
𝑉
𝐼𝑖
𝛺
𝑉
𝑅𝑖
𝛺
)

(
13
)

𝑌
𝑃
𝛺
+
𝑗
𝑌
𝑄
𝛺
=

1
𝑍
𝑃
𝛺
+
𝑗
𝑍
𝑄
𝛺

(
14
)

For the load model given in

(
10
)

through

(
14
)
, the constant
impedance part of the load is linear
,

whereas the constant
current and constant power part of the aggregated load
is

nonlinear.
Once,
(
10
)
-
(
11
)

are linearized, they are used to
construct the equivalent circuit models for both the power flow
and threephase power flow problem. The

cons tructed
threephase model of the

ZIP load mode l

can either be connected in
wye or delt a formation. As

an

example, ZIP load model
connected in wye

and delta

formation is shown in
Fig. 3
.

Figure
3
: Real circuit
for a) wye connected ZIP Load Model (on left) b) delta
(D) connected ZIP load model (on right).

It is important to note that both the ZIP and PQ load models
result in nonlinear network constraints for both the ‘PQV’ and
CIM formulations.
In the ‘PQV’ form ulation the nonlinearities
in the network constraints are due to the
use of power mismatch
equations
whereas in the CIM,

the

nonlinearities are due to PQ
and ZIP model equations
.
These added nonlinearities are one
of the

primary causes of divergence and

convergence to
low
voltage
solutions. To address this problem, we have proposed
an accurate and yet linear

BIG
load model
[25]
-
[27]
.

D.
BIG Linear Load Model

The
BIG

aggregated load model was proposed based on the
circuit theoretic approach in
[25]
-
[27]

and aims to create a
linear load model that can capture the true measure and
sensitivity of t he aggregated load in the system
.
The model is
comprised of a susceptance (B), independent current source (I),
and conductance (G).
The complex governing equation

of the
generalized load current for the
BIG

load model is given by:

(
𝐼
𝑅
𝐵𝐼𝐺
)
Ω
+
𝑗
(
𝐼
𝐼
𝐵𝐼𝐺
)
Ω
=
(

𝑅
𝐵𝐼𝐺
)
Ω
+
𝑗
(

𝐼
𝐵𝐼𝐺
)
Ω

(
15
)

+
(
(
𝑉
𝑅
𝐵𝐼𝐺
)
𝛺
+
𝑗
(
𝑉
𝐼
𝐵𝐼𝐺
)
𝛺
)
(
(
𝐺
𝐵𝐼𝐺
)
𝛺
+
𝑗
(
𝐵
𝐵𝐼𝐺
)
𝛺
)

w here

(

𝑅
𝐵𝐼𝐺
)
Ω
+
𝑗
(

𝐼
𝐵𝐼𝐺
)
Ω

represents the base
value for the
modeled aggregated load

and
the
corresponding
complex
admittance
(
(
𝐺
𝐵𝐼𝐺
)
𝛺
+
𝑗
(
𝐵
𝐵𝐼𝐺
)
𝛺
)

captures the voltage
sensitivities.
For instance
, a negative conductance in
conjunction with complex current
(
(

𝑅
𝐵𝐼𝐺
)
Ω
+
𝑗
(

𝐼
𝐵𝐼𝐺
)
Ω
)

mimics

the inverse current/voltage sensitivity relationship
, similar to
constant power (PQ) load behavior and
positive conductance
in
conjunction with complex current source
will represent the
positively

correlated current/voltage sensitivity relationship
,
similar to the impedance load

behavior
. Both the positive
and
negative impedances capture the change in load with voltage
with respect to the portion of the load that is modeled by the
current source
.
Fig. 4

s hows the positive sequence (
p
)
BIG

load
model.
Similar to

the

ZIP load model, the threephase
BIG

load
model can be constructed by connecting the equivalent circuits
of individual phases in wye or delta formation.

Figure
4
:
Equivalent circuit of a BIG load model

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

5

IV.
C
IRCUIT
S
IMULATION
M
ETHODS

Decades of research in circuit simulation have demonstrated
that circuit simulation methods can be applied for determining
the DC state of highly nonlinear circuits using NR. These
techniqu es have been shown to make NR robust and practical
for largescale circuit problems
[21]
-
[22]

consisting of billions
of nodes. Most notable is the ability to guarantee convergence
to the correct physical solution (i.e. global convergence) and the
capability of finding multiple operating points
[28]
. We have
previously

proposed analogous techniques for ensuring
convergence to the correct physical solution for the
positive
sequence
power flow

problem

[19]
-
[20]
. In this section, we
extend these methods to be use d with positivesequence power
flow and threephase power flow problems alike
.

A.
General Methods

1)
Variable Limiting

The solution space of the system node voltages in a power
flow and threephase power flow problem are well defined.
While solving these problems, a large NR step may step out of
this solution space and result in either divergence or
convergence to a nonph ysical solution.
It is
,

therefore
,

important to limit the NR step before an invalid step out of the
solution space

is made
. In
[19]
,

we proposed
th e
variable
limiting
method
to achieve the postulated goal

for power flow
problem
. In this technique, the state variables that are most
sensitive to initial guesses are damped when the NR algorithm
takes a large step out of the predefined solution space. N
ote
,
however,

that
not all

the
system
variables
are damped for the
variable limiting technique, as is done for traditional damped
NR. C
ircuit simulation research has shown that
damping the
most sensitive variables

provides superior convergence
compared to
damped NR in
general
[21]
.

To apply variable limiting in our prototype simulator for the
power flow and threephase power flow problem, the
mathematical expression s for the PV nodes in the system are
modified as follows:

𝐼
𝐶𝐺
𝛺

1
=

𝜍
𝐼
𝐶𝐺
𝛺
𝑉
𝑅𝐺
𝛺
(
𝑉
𝑅𝐺
𝛺
k

1
−
𝑉
𝑅𝐺
𝛺
k
)
⏟

∆
𝑉
𝑅𝐺
𝛺
+
𝐼
𝐶𝐺
𝛺
k

(
16
)

+

𝜍
𝐼
𝐶𝐺
𝛺
𝑉
𝐼𝐺
𝛺
(
𝑉
𝐼𝐺
𝛺

1
−
𝑉
𝐼𝐺
𝛺

)
⏟

∆
𝑉
𝐼𝐺
𝛺
+
𝐼
𝐶𝐺
𝛺
𝑄
𝐺
𝛺
(
𝑄
𝐺
𝛺

1
−
𝑄
𝐺
𝛺

)

where,

0
≤
ς
≤
1
. The magnitude of

ς

is dynamically varied
through heuristics such that convergence to the correct physical
solution is achieved in the most efficient manner. The
heuristics
depend on the largest delta voltage
(
∆
𝑉
𝑅𝐺
𝛺
,
∆
𝑉
𝐼𝐺
𝛺
)

step during
subsequent NR iterations. If during subsequent NR iterations, a
large step
(
∆
𝑉
𝑅𝐺
𝛺
,
∆
𝑉
𝐼𝐺
𝛺
)

is encountered,

then

the factor
ς

is
decreased. The

factor

ς

is scaled back up if consecutive NR
steps result in monotonically decreasing abs ol ute values for the
largest error
.

2)
Voltage Limiting

An equally simple, yet effective, technique is to limit the
absolute value of the delta step that the real and imaginary
voltage vectors can make during each NR iteration.
This
is

analogous to the voltage limiting technique used for diodes in
circuit simulation, wherein the maximum allowable voltage
step during NR is limited to twice the thermal voltage of the
diode
[22]
. Similarly, for the power flow and threephase power
flow analyses, a hard limit is enforced on the normalized real
and imaginary voltages in the system. The mathematical
implementation of vo ltage limiting in our formulation is as
follows:

(
𝑉
𝐶
Ω
)
k

1
=
𝑚𝑖𝑛
𝑉
𝐶
𝑚𝑖𝑛
𝑚𝑎𝑥
𝑉
𝐶
𝑚𝑎𝑥
(
(
𝑉
𝐶
Ω
)
k
+
𝛿
𝑆
𝑚𝑖𝑛
(
|
∆
(
𝑉
𝐶
Ω
)
k
|
,
∆
𝑉
𝐶
𝑚𝑎𝑥
)
)

(
17
)

𝑚𝑖𝑛
𝑉
𝐶
𝑚𝑖𝑛
𝑚𝑎𝑥
𝑉
𝐶
𝑚𝑎𝑥
=

{
𝑉
𝑐
𝑚𝑎𝑥
,
𝑖𝑓

𝑥
>

𝑉
𝑐
𝑚𝑎𝑥

𝑉
𝑐
𝑚𝑖𝑛
,
𝑖𝑓

𝑥
<

𝑉
𝑐
𝑚𝑖𝑛

𝑥
,
𝑜𝑡
ℎ
𝑒𝑟𝑤𝑖𝑠𝑒

(
18
)

and
𝛿
𝑆
=
𝑠𝑖𝑔𝑛
(
∆
(
𝑉
𝐶
Ω
)
k
)

and
𝐶
∈
{
𝑅
,
𝐼
}

represents the

placeholder for real and imaginary parts.

Analogously,
other
system
variables such as the reactive
power
𝑄
𝐺

of the PV buses,
can
be limited by limiting the
calculated

currents
𝐼
𝐶
Ω
+

∆
(
𝐼
𝐶
Ω
)
k

at NR step
k
+
1

and then
finding the new
𝑄
𝐺

1

from

inverse function

(
𝑓
−
1
)

of limited
(
𝐼
𝐶
Ω
+

∆
(
𝐼
𝐶
Ω
)
k
̅
̅
̅
̅
̅
̅
̅
)
.

B.
Homotopy Methods

Limiting methods may fail to ensure

convergence for certain
illconditi oned and large test systems when solved from an
arbitrary set of initial guesses. To

ensure convergence for

these
network models

to the correct physical solutions independent
of the choice of initial conditions
, we propose the use of
homotopy methods.

Homotopy

methods in past have been used
to study the voltage collapse

of a given network

or to determine
maximum available transfer capability

[8]
-
[9]
. They have also
been researched
for locating
all solutions to a power flow

problem

[11]
,
[30]
. However,
their usage

for

enabling

convergence for hard to solve

positive sequence and threephase power flow problems

has been limited

at bes t.
Of the
proposed methods for better convergence

[5]
,
[23]
, most

have
suffered f rom convergence to low voltage solutions
or
divergence.
On the other hand, some of them have been
developed for formulations that are not standard for both
positive sequence as well as threephase power flow
[6]
.
Furthermore, none of the previously proposed homotopy
methods are known to scale up to test systems

that are of the
scale of
the
European or the US
grids and

in general are not
extendable to
the
threephase power flow problem.

In the homotopy approach, the original problem is replaced
with a set of subproblems that are sequentially solved. The set
of subproblems exhibit certain properties, namely, the first
subproblem has a trivial solution an d each incrementally
subsequent problem has a solution very close to the solution of
the prior subproblem. Mathematically this can be described via
the following expression:

ℋ
(
𝑥
,
𝜆
)
=
(
1
−
𝜆
)
ℱ
(
𝑥
)
+

𝜆𝒢
(
𝑥
)

(
19
)

where

𝜆



[
0
,
1
]
.

The method begins by replacing the original problem
ℱ
(
𝑥
)
=
0

with

ℋ
(
𝑥
,
𝜆
)
=
0
.
The equation set
𝒢
(
𝑥
)

is a
representation of the system that

has a trivial solution. The
homotopy factor
𝜆

has the value of 1 for the first subproblem
and therefore the initial so lution is equal to trivial solution
of

𝒢
(
𝑥
)
. For the final subproblem that corresponds to the
original problem, the homotopy factor
𝜆

has the value of zero.
To generate sequential subproblems, the homotopy factor is
dynamically decreased in small steps
until it has reached the
value of zero.

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

6

In this paper, we discuss two homotopy methods that
are
specifically developed for the
power flow and threephase
power flow analyses
:

1)
Tx Stepping

We proposed the “Tx Stepping” method

in

[20]

specifically
for the power flow problem. In this section, the method is
further extended for the threephase power flow problem.

a)
General Approa ch

In Tx stepping method, the

series elements in the system
(transmission lines, transformers etc.) are first “virtually”
shorted to

solve the initial problem that
has

a trivial solution.
Specifically, a large conductance
(
≫
𝐺
𝑖𝑙
)

and
a large
susceptanc e (
≫
𝐵
𝑖𝑙

) are added in parallel to each transmission
line and transformer model in the system
.
In case of threephase
power flow, a large selfimpedance
(
≫
Y
ΩΩ
𝑖𝑙
) is added in parallel
to each phase of the transmission line and transformer model.
Furthe rmore, the shunts in the system, are opencircuited by
modifying the original shunt conductance and susceptance

values.
Importantly, the solution to this initial problem results
in high system voltages

(magnitudes),

as they are essentially
driven
by

the slack bus

complex voltages and the PV bus
voltage magnitudes

due to the low voltage drops in the lines and
transformers (
as expected with virtually shorted systems
).

Similarly, the solution
for the
bus voltage angles will lie within
a n


-
small radius
around the slack bus angle.
Subsequently,
like other continuation methods, the formulated system
proble m is then gradually relaxed

to represent the original
system by taking small increment steps of the homotopy factor
(
𝜆
) until convergence
to the solutio n of

the original problem is
achieved. Math ematically,

the line and transformer
impedances during homotopy

for the power flow is expressed
by
:


𝑖𝑙

∈

{
𝒯
𝑋

,
𝑥𝑓𝑚𝑟𝑠
}
∶
𝐺
̂
𝑖𝑙
+
𝑗
𝐵
̂
𝑖𝑙
=

(
𝐺
𝑖𝑙
+
𝑗
𝐵
𝑖𝑙
)
(
1
+
𝜆𝛾
)

(
20
)

and for

the threephase problem:

[
𝑌
̂
𝑎𝑎
𝑖𝑙
𝑌
̂
𝑎𝑏
𝑖𝑙
𝑌
̂
𝑎𝑐
𝑖𝑙
𝑌
̂
𝑏𝑎
𝑖𝑙
𝑌
̂
𝑏𝑏
𝑖𝑙
𝑌
̂
𝑏𝑐
𝑖𝑙
𝑌
̂
𝑐𝑎
𝑖𝑙
𝑌
̂
𝑐𝑏
𝑖𝑙
𝑌
̂
𝑐𝑐
𝑖𝑙
]
=

[
Y
𝑎𝑎
𝑖𝑙
(
1
+
𝛾𝜆
)
Y
𝑎𝑏
𝑖𝑙
Y
𝑎𝑐
𝑖𝑙
Y
𝑏𝑎
𝑖𝑙
Y
𝑏𝑏
𝑖𝑙
(
1
+
𝛾𝜆
)
Y
𝑏𝑐
𝑖𝑙
Y
𝑐𝑎
𝑖𝑙
Y
𝑐𝑏
𝑖𝑙
Y
𝑐𝑐
𝑖𝑙
(
1
+
𝛾𝜆
)
]

(
21
)

where
,

𝐺
𝑖𝑙
,

𝐵
𝑖𝑙

, and

𝑌
ΩΩ
𝑖𝑙

are the original system impedances
and
𝐺
̂
𝑖𝑙
,
𝐵
̂
𝑖𝑙
, and

𝑌
̂
ΩΩ
𝑖𝑙

are the system impedances used while
iterating from

the

trivial problem to the original problem. The
parameter

𝛾

is used as a

scaling factor for the
conductances and
susceptances.

If the

homotopy factor
(
𝜆
)

takes the value one,
the system has a trivial solution and if its takes the value zero,
the original system is represented
.

Along with ensuring convergence
for a problem
,

Tx
s tepping
avoids

th e undesirable low voltage solutions for the positive
sequence power flow and threephase power flow problem

since the initial problem results in a solution with high system
voltages, and
each
subsequent step of the homotopy approach
c ontinues and deviates ever so slightly

from this initial solution,
thereby guaranteeing convergence to the high voltage solution
for the original problem.

b)
Handling of Transformer Phase Shifters and Taps

To “virtually short” a power system, we must also acc ount
for transformer taps
𝑡𝑟
Ω

and phase shift ing angles

𝛩
Ω
. In a
“virtually” shorted condition, all the nodes in the system must
have complex voltages that are near the slack bus or PV bus
complex voltages, which can be intuitively defined by a small
ep silon norm ball around these voltages. Therefore, to achieve
the following form, we must modify the transformer taps and
phase shifter angles such that at

𝜆
=
1
,

their turns ratios and
phase shift angles correspond to a magnitude of
1

pu

and
0
°
,
respectively. Subsequently, the homotopy factor
𝜆

is varied
such that the original problem is solved with original
transformer tap and phase shifter settings. This can be
mathematically expressed as follows:


𝑖

∈

𝑥𝑓𝑚𝑟𝑠
∶
𝑡𝑟
̂
𝑖
𝛺
=
𝑡𝑟
𝑖
𝛺
+
𝜆
(
1
−
𝑡𝑟
𝑖
𝛺
)

(
22
)


𝑖

∈

𝑥𝑓𝑚𝑟𝑠
∶
𝛩
̂
𝑖
𝛺
=
𝛩
𝑖
𝛺
−
𝜆
𝛩
𝑖
𝛺

(
23
)

c)
Handling of Voltage Control for Remote Buses

To achieve a trivial solution during the first step of Tx
stepping it is essential that we also
handle remote voltage
control appropriately. Remote voltage control refers to a device
on node

𝒪

in the system controlling the voltage of another node
𝒲

in the system. This behavior is highly nonlinear and if not
handled correctly can result in divergen ce or converge nce

to
low voltage solution. Existing commercial tools for power flow
and threephase power flow analyses
have difficulties dealing
with this problem and suffer from lack of robust convergence
when modeling remote voltage control in general.
With Tx
stepping we can handle this problem efficiently and effectively.
We first incorporate a “virtually short path” between the
controlling node (
𝒪
) and the controlled node (
𝒲
) at

𝜆
=
1
,
such that the device at the controlling node can easily supply
t he current needed for node
𝒲

to control its voltage. Then
following the homotopy progression, we gradually relax the
system such that additional line connecting the controlling
node
(
𝒪
)

and
controlled
node
(
𝒲
)

is open at

𝜆
=
0
.

d)
Implementation of Tx Stepp ing in Equivalent Circuit
Formulation

Unlike traditional implementations of
homotopy methods, in
equivalent circuit formulation we do not directly

modify the
nonlinear set of mathematical equations, but instead embed a
homotopy factor in each of the equiv alent circuit models for the
power grid components. In doing so we allow for incorporation
of any power system equipment into the Tx stepping approach
within
the
equivalent circuit formulation framework, without
loss of generality. Furthermore, we ensure,
that the physics of
the system is preserved while modifying it for the homotopy
method.
Fig. 5

demonstrates

how the homotopy factor is
embedded into t he equivalent circuit of the transformer.

Figure
5
: Homotopy factor embedded in transformer equivalent circuit
.

2)
Dynamic Power Stepping

Another homotopy technique that can ensure robust
convergence for systems that have a low percentage of constant
voltage nodes in the system is the dynamic power stepping
method. Existing distribution systems tend to
belong to these
types of systems

and th erefore, dynamic power stepping can be
applied to robustly obtain the steadysolution of the distribution
grid by solving the threephase power flow problem.

This
method has been previously described for the positivesequence
IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

7

power flow problem in
[19]

and

is analogous to the

source
stepping and gmin stepping approaches in standard circuit
simulation solvers.

In the power stepping method, the system
loads

and
generat ion are

scaled back by a factor of

𝛽

until the
convergence is achieved
. If these loads and generation s

are
scaled down all the way to zero, then the
constraints for the
PQ
buses in the system
result in linear network constraints
.
Similarly, current

source

nonlinearities

of the PV buses
that
are due to the constant real power
are also eliminated
.
Therefore, by applying the power stepping factor
, the nonlinearities in the system are greatly eased and convergence is
easily achieved. Upon convergence,

the factor is gradually
scaled back up to unity to solve the original problem.
In this
method, as in all continuation methods, the solution from the
prior step is used as the initial condition for the next step. The
mathematical representation of dynamic
power stepping for the
threephase power flow and positive sequence power flow
problem is as follows:


𝐺
∈

𝑃𝑉
:
𝑃
̂
𝐺
Ω
=
𝛽
𝑃
𝐺
Ω

(
24
)


𝐿
∈

𝑃𝑄
:
𝑃
̂
𝐿
Ω
=
𝛽
𝑃
𝐿
Ω

𝑎𝑛𝑑

𝑄
̂
𝐿
Ω
=
𝛽
𝑄
𝐿
Ω

(
25
)

w here
,

PQ

are all

load

nodes and
PV

are all
generator

nodes.

V.
P
OWER
F
LOW AND
T
HREEP
HASE
P
OWER
F
LOW
A
LGORITHM

Algorithm 1
: Simulation algorithm for Positive Sequence and ThreePhase
Power Flow

Solver

Algorithm 1 shows the recipe for the solving the positivesequence
as
well as

threephase power flow problem in
equivalent circuit approach with

the use of

circuit simulation
methods.
In this framework, t he solver
starts with building

the
system models based on the input file supplied. Linear models
(
𝑌
𝐿
,
𝐽
𝐿
)

are then stam ped in the Jacobian matrix. Input state
variable s

and other continuation parameters (
𝑥
0
,
𝛿
,
𝜁
,
𝜆
)

are then
initialized. Nonlinear models are then stamped (
𝑌
𝑁𝐿
,
𝐽
𝑁𝐿
)

and
NR
is

applied with limiting

methods

to calculate

the

next iterate
for voltage s and generator reactive powers
(
𝑋
̂

1
)
.
Continuation and limiting

parameters are then dynamically
updated and homotopy models

(
𝑌
𝐻
,
𝐽
𝐻
)

are stamped

or restamped

if
required

to ensure convergence. Upon convergence
of inner loop generator limits, switch ed shunts and transformer
taps are adjusted and inner loop is repeated until final solution
is achieved.

VI.
R
ESULTS

Example cases were simulated in our prototype solver
SUGAR (
S
imulation with
U
nified
G
rid
A
nalyses and
R
enewables) to
demonstrate that the equiv alent circuit
approach along with circuit simulation techniques facilitates a
robust framework for positive sequence power flow and the
threephase power flow analyses
.

T
he first set of results
compare
the
solution
of

contingency analyses for two hard to
solve cases

with and without the use of circuit simulation
methods

to demonstrate the
efficacy
for these methods
. All the
further results compare the results of SUGAR (with circuit
simulation methods) with other industry tools.
The example
cases
for positi ve sequence power flow analyses
include known
illconditioned
test cases
and large
network models

that
represent different operating

and loading

conditions
for

the
eastern interconnection

network

of the US grid.
For
the
threephase power flow analysis, exa mple cases include a set of
standard distribution taxonomy cases
[29]
, high density urban
test cases
[31]
,

and
a
meshed transmission grid test case that
was
modified from

a

positive sequence to

a

threephase
network model. The results

that follow demonstrate
that the
proposed framework

along with the

use of

circuit simulation

methods

can
ensure

convergence to

a

correct physical solution
for
all
the

power flow

and threephase power flow

cases,
independent of the choice of the initial guess

and thus
overcomes
the challenges faced by existing formulations
.

A.
Circuit Simulation Methods

The purpose of
following

set of results is to demonstrate
the
robustness of the solver
that is enabled
due to the use of circuit
simulation methods
. To
show

this,

contingencies

were
simulated

on

two (2) hard to solve testcases
that represent

a
real network for the
subset of
the
US power grid
.

The base case
for both simulations
is

first solved via Txstepping method and
then used as an

initial condition for
the set of contingencies
.
The contingencies
in the

contingency set
repres ent

the loss of
largest 10% of online generators and highest capacity lines

and
transformers

dropped one at a time
.

T
ABLE
1
:

C
OMPARISON OF
SUGAR

WITH AND
W
ITHOUT
C
IRCUIT
S
IMULATION
T
ECHNIQUES

Case
Id

#

Bus

#
Total
Cases

SUGAR w/o
Circuit
Simulation Methods

SUGAR with Circuit
Simulation Methods

Converge

Diverge

/Infeasible

Converge

Diverge

/Infeasible

Case
1

5944

754

735

19

750

4

Case
2

7029

801

706

95

793

8

The results in
th e
Table
1

confirm that
the
circuit simulation
methods when applied to equivalent

circuit formulation can
significantly increase the robustness of the power flow solver.

IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

8

B.
Positive Sequence Power Flow Re sults

1)
IllConditioned and Large Test cases

A convergence sweep was run on
the

illconditioned 13659
bus PEGASE test case
using the

SUGAR solver and a standard
commercial tool and their results were compared
.
Fig. 6

shows
that SUGAR was

able to robustly converge

to

the
correct
physical solution independent of the choice of
the
initial
condition s
, whereas the standard tool was highly sensitive to
the choi ce of

the

initial guess and could converge to the correct
physical solution

only

from a few samples
for the
initial guess.

Figure
6
: Convergence sweep comparison for 1365
9

node PEGASE testcase
between SUGAR and Standard tool
.

Red

indicates

divergence

and
green

indicates

convergence

A similar convergence sweep was performed for larger test
cases

(> 75k+ nodes)

that represent different loading and
operating scenarios for eastern interconnection of the US grid.

Simulations were performed on three
different
test cases

for 15
different initial conditions each. Results are shown in
Table
2
.

The set of initial con ditions

for all buses were identical and

were uniformly sampled from:

𝑉
𝑎𝑛𝑔

∈

[
−
40
,
40
]

,
𝑉
𝑚𝑎𝑔

∈

[
0
.
9
,
1
.
1
]
.

(
26
)

T
ABLE
2
:

C
ONVERGENCE
P
ERFORMANCE FOR
L
ARGE
E
ASTERN
I
NTERCONNECTION
T
EST
C
ASES

Case
Name

# Nodes

Standard Tool

SUGAR

#
Converge

#
Diverge

#
Converge

#
Diverge

Case 1

80778

0

15

15

0

Case 2

76228

0

15

15

0

Case 3

81904

0

15

15

0

For the larger eastern interconnection test cases, the runtime
per iteration is less than 0.4 seconds and is comparable to other
simulation tools out in the market. The total
computation time

in general is dependent on the choice of initial conditions. A
s ufficiently close initial condition may result in convergence
within 7 iterations whereas a totally random set of initial
guesses may take up to 100 iterations with Tx stepping method.

2)
Contingency Analysis

In the next set of results, we performed a set of contingency
analyses
with

SUGAR and

a
standard commercial tool
for
two
test cases
that represent

different network configuration of the

eastern interconnection of the US grid. The initial guess for

solving

the contingency cases was chosen to be the operating
point prior to the contingency. The set of contingencies in the
experiment include s

loss of generation (
ℒ
𝐺
) and loss of
branch es

(
ℒ
𝐵
). The results are summarized in
Table
3

and

highlight the need for continuation methods to solve

such
problems robustly.

T
ABLE
3
:

C
OMPARISON OF
C
ONTINGENCIES OF
L
ARGE
T
EST
C
ASES

Case

# Nodes

Contingency
*

Standard Tool

SUGAR

Case 1

76228

2ℒ
𝐺

Diverge

Converge

2ℒ
𝐺
+

2ℒ
𝐵

Diverge

Converge

Case 2

78201

2ℒ
𝐺

Diverge

Converge

2ℒ
𝐺
+

2ℒ
𝐵

Diverge

Converge

*
The number in
front of

ℒ
𝐺

and
ℒ
𝐵

represents the
equipment outage count.
(For
e.g. 2
ℒ
𝐺

represents that two generators were lost during this contingency).

C.
ThreePhase Power Flow Results

1)
Taxonomical Test Cases

Table
4

documents the result s obtained from
SUGAR threephase solver for standard

taxonomical cases
and

three

large
meshed test case s
.
The standard taxonomical cases
include

both
balanced and unbalan ced threephase test cases.
The first two
meshed test cases

are

342
-
Node Low Voltage Network Test
Systems
[31]

that represent high density

urban

meshed low
voltage network s
.
The third meshed test system is a high
voltage 9241 node PEGASE transmission system

that was
extended

to a

balanced

threephase

model. All these cases were
simulated in SUGAR threephase solver to

validate
its
accuracy
by comp aring the
obtained results

against a standard
distribution power flow tool GridLABD. Slight difference s
(less than 1e
-
2
)

in
the
results

were observed

for
cases

between
SUGAR and GridLABD
and are

due to default values used for
unspecified parameters

(e.g.

neutral conductor resistance)

in
GridLABD.

T
ABLE
4
:

SUGAR

T
HREEP
HASE RESULTS FOR
T
AXONOMICAL
C
ASES

Cases

#Nodes

Iter.
Count

Deviation from GridLABD

Max. ΔV
mag

[pu]

Max. ΔV
ang

[°]

R1
-
12.47
-
1

2455

5

8.73E
-
04

9.94E
-
03

R2
-
12.47
-
3

2311

5

6.56E
-
04

1.32E
-
02

R3
-
12.47
-
3

7096

5

1.94E
-
03

3.89E
-
02

R4
-
12.47
-
1

2157

5

6.81E
-
04

9.61E
-
03

R5
-
12.47
-
5

2216

5

5.44E
-
05

4.20E
-
03

NetworkModel

1

1420

3

3.38E
-
03

2.14E
-
03

NetworkModel
2

1420

3

3.83E
-
03

6.00E
-
03

case9241pegase
*

12528

5

NA
#

NA
#

* 9241 bus PEGASE transmission test case was
extended to threephase model

#The following case did not run in GridLABD

2)
IllConditioned Test Cases

To solve certain hard to solve illconditioned threephase
test cases, we made use of homotopy methods. To demonstrate
one such example, we
extended
the standard 145 bus
transmission system model into a
balanced
threephase network
model.

Figure
7
: Converge nce

of 145 bus test case for threephase power flow with

(bottom)
and without

(top)

power stepping.
For

the power stepping case
, the

green

dotted line
represent s

the change
in

continuation factor λ

Fig. 7

plots the convergence results for th is

test case with and
without the use

of

dynamic power stepping. Without the use

of
IEEE TRANSACTIONS ON POWER SYSTEMS

0885
-
8950 © 2018 IEEE. Personal use is permitted, but republication/redistribution requires IEEE
permission.

See http://www.ieee.org/publications standards/publications/rights/index.html for more information.

9

dynamic power stepping,

the test system did not converge
within maximum number of allowable iterations; however, with
the use of dynamic power stepping, the system robustly
converged to correct physical solution.

VII.
C
ONCLUSIONS

In this paper
,

we have demonstrated that
the

equivalen t
circuit approach with the use of novel circuit simulation
methods can robustly solve for the steadystate solution of the
transmission and distribution grid without loss of generality.
Th is

proposed formulation and the

analogous

circuit simulation
method s can be generically applied to both the positive
sequence power flow problem and the threephase power flow
problem. Importantly, our approach toward steadystate
analyses of transmission and distribution grid ensures robust
convergence to correct physica l solution s
, and in doing so
enables robust contingency analyses, statistical analyses, and
security constrained optimal power flow analyses.

Furthermore,
the proposed
generic framework

for transmission and
distribution grid analys e
s
can be extended for

jo int simulation
of transmission and distribution cir

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
