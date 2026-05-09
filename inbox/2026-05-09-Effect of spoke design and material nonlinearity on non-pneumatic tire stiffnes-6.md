---
source: "https://arxiv.org/abs/2103.03637v1"
title: "Effect of spoke design and material nonlinearity on non-pneumatic tire stiffness and durability performance"
author: "Priyankkumar Dhrangdhariya, Soumyadipta Maiti, Beena Rai"
year: "2021"
publication: "arXiv preprint / cond-mat.mtrl-sci"
download: "https://arxiv.org/pdf/2103.03637v1"
pdf: "https://arxiv.org/pdf/2103.03637v1"
captured_at: "2026-05-09T13:03:53Z"
updated_at: "2026-05-09T13:03:53Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Thus Spoke Zarathustra"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Effect of spoke design and material nonlinearity on non-pneumatic tire stiffness and durability performance

- 著者: Priyankkumar Dhrangdhariya, Soumyadipta Maiti, Beena Rai
- 年: 2021
- 掲載情報: arXiv preprint / cond-mat.mtrl-sci
- 情報源: [arxiv](https://arxiv.org/abs/2103.03637v1)
- ダウンロード: https://arxiv.org/pdf/2103.03637v1
- PDF: https://arxiv.org/pdf/2103.03637v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Non-pneumatic tire has been widely used due to their advantages of no run-flat, no need of air maintenance, low rolling resistance, and improvement of passengers comfort due to its better shock absorption. It has variety of application in the military vehicle, earthmovers, lunar rover, stair climbing vehicles etc. Recently UPTIS (Unique Puncture-Proof Tire System) non pneumatic tire has been introduced for passenger vehicles. In this study three different design configuration Tweel, Honeycomb and newly developed UPTIS have been compared. Effect of Polyurethane (PU) material nonlinearity have also been introduced by applying 5 different nonlinear PU material property in the spokes. The combined analysis of the PU material nonlinearity and spoke design configuration on the overall tire stiffness and spoke damage prediction is analysed using 3-Dimensional FEM simulations performed in ANSYS 16.0. It has been observed that Mooney Rivlin 5-parameter model is best to capture all 5 studied PU materials the nonlinearity. Effect of material nonlinearity on various spoke designs have been studied. The best combination of spoke design and the use of nonlinear material have been suggested in terms of riding comfort, tire stiffness and durability performance.

## PDF Text

1

Effect of spoke design and material
non linearity
on
n onpneumatic tire stiffness and durability performance

Priyankkumar Dhrangdhariya
1
, Soumyadipta Maiti
1
*
, Beena Rai
1

1

TCS Research,
Tata Research Development and Design Centre,

54
-
B, Hadapsar
Industrial Estate
,

Pune
-
4
11013,

India

*

Corresponding author email:
soumya.maiti@tcs.com

Abstract

Nonpneumatic tire
has been widely used

due to their advantages of no runflat, no need of air
maintenance, low rolling resistance, and im provement of passenger’s comfort due to

its

better shock
absorption. It has variety of application in the

m ilitary
v ehicle,
e arthmovers,
l unar
r over,
s tair climbing
vehicles etc.

Recently
UPTIS (Unique PuncturePr oof Tire System
) non
pneumatic

tire

has bee n
introduced

for passenger vehicle s
.

In this study
three different design configuration Tweel, Honeycomb
and newly developed UPTIS have been
compared
.

Effect of
Polyurethane (
PU
)

material nonlinearity
have
also
been introduced by applying 5 different
non linear

PU material property in the spokes
.

T
he
combin ed analysis

of the PU material
nonlinearity

and spoke design configuration on the overall tire
stiffness and spoke damage predi ction

is

analysed using

3
-
Dimensional

FEM simulations performed in
ANSYS 16
.0.
It has been observed that
Mooney Rivlin 5
-
parameter model is best to capture all 5
studied
PU material’s the nonlinearity
.

Effect of material non linearity

on various spoke des igns have been
studied.
The best combination of spoke design and
the use of

nonlinear
material have been suggested
in terms of riding comfort, tire stiffness and durability performance.

1

Introduction

For last few year s,

the automotive industries are looking into the development of nonpneumatic tyres
(NPTs)
to attain

the poten tial advantages over pneumatic tyres in terms of no runflat, no need of air
maintenance,

low rolling resistance

and

improvement of passenger’s com fort due to better shock
absorption

[1
–
3]
.
Several tire engineers have attempted to develop nonpneumatic tires by filling an

elastomer into the space for air or by building polygon typed spokes to replace the air of a pneumatic
tire
[3
–
6]
.

2

A typical NPT

usually consists of a hub, flexible spokes, a
shear ring

and a tread
.

The axle of the vehicle
is attached to the rigid hub. Aluminium alloy is used as a material for Hub.
The main task of the tread
is to provide proper cooperation of its blocks with the g round and providing protection of reinforcement
layers against damage r esulting from the impact of road
’
s obstacles. A rubber compound is used as a
material for a tread layer.
The shear ring acts as a supporting structure determines the properties of the
n onpneumatic tire (e.g. directional stiffness, rolling resistance), whi ch consists of composite structure
composed of a shear band with two circumferential reinforcement s (i.e., inner and outer rings)
[3,7,8]
.

Flexible spokes have remarkable impact on the ov erall performance of NPT, which is made from
Pol yurethane
[9]
.
The NPT stiffness properties and contact pressu re during loading is affected by several
geometric and material parameters, including ring thickness, reinforcement type, and spoke geometry

[10]
.
Various shapes of flexible spokes have been introduc ed by Michelin, Bridgestone,

and

Polari s for
the desired application like
Military Vehicle, Earthmovers, Lunar Rover, Stair climbing vehicles,
Passenger vehicles etc.

In 2005, Michelin developed the Non pneumatic tire named as Tweel which is
classified by
Time magazine as “one of the most amazi ng inventions” in that year
[11]
.
Resilient
Technology introduced NPT having Honeycomb spokes in 2012, later tested and used by U.S army in
the military vehicle for i ts excell ent ‘b ulletproof’
capability
[12,13]
.

I
n 2013, Bridgestone revealed its
secondgeneration

air free concept nonpneumatic tire featuring improved loadbearing capabilities,
environmental de sign and driving performance
.

Recently Michelin presented UPTIS (Unique PunctureProof Tire System)
at the
Movin’On Summit in 2017,

which

is
an airless mobility solution for passenger
vehicles with the goal of introduci ng the
model

as early as 2024
[14]
.

Many
literatures

are available in design of spokes and its effect on the

c ellular spo ke

or

variation in
honeycomb spoke design
[7,15
–
17]
.

K
ucewicz et

al

presented the modelling methodology of nonpneumatic tire having different spoke structure and conducted radial defle ction analysis using Finite
ele ment modelling
(FEM)
te c
hniques
[17]
.

Ingrole et al

c ompared different H
oneycomb spoke designs
and suggested
A
uxetic
structure

hav ing

good compressive strength

compared to other design s
[15]
.

The
team has found good agreement of spoke deformation shape
received from F
EM with experiments.

Bezgam et

al

analysed t he effect of cell angles of honeycomb spokes with the same thickness an d
same
load carrying capacity on

the deformation, stress distribution and contac t pressure distribution of

NPT
[18]
.

Ju et al

investigated the validity of NPT with high fatigue resistance by achiev ing low local
stress under vertical stiffness loading using FEM analysis. The elastic limits of several hexagon

honeycomb

structure having different cell angle of spokes were analy s
ed and compared

to suggest
a

better geometrical designs of lower localize d peak stress and fatigue properties

[7]
.

In
literature

t hree
different
configurations

of Michelin Tweel, Resilient technology and Bridgestone NPT
were

investigated by conducti ng 2D FEM analysis, to find

out

comparison
of
these spoke structure on the

3

contact pressure, vertical tire stiffness and
stresses of NPT. The results showed that the shape of the
spokes ha ve

a great effect on tire’s behaviour.
[19]

Macro and local cellular properties of hexagonal honeycombs have

also been extensively
stu died
[10,20
–
23]
.

Rugsaj

et al

have extracted the material property of NPT by preparing the spec imen
using water jet cutting. They have developed 3D FEM model to capture the deformation behaviour of
these NPT component with the help of
hyperelastic constitutive models (i.e. Neo Hookean model,
Ogden model, Mooney Rivlin model)
[24]
.
The most commonly used material models fo r describing
the properties and behavior of hyp erelastic material
in
supporting structure, spokes and tread

are
:
Ogd en, MooneyRivlin, NeoHook ean and Marlow
[7,8,17,25]
.

There are many factors which are useful to understand the durability performance of NPT. The strain
en ergy density

(SED)

has been used as a predictor of fatigue life in elastomers since the development
of the fracture mechanics for this class of material

[26]
. The
SED

has been found to be

proportional to
crack density and size under simple load conditions, so it can be used as a measure of the fatigue life of
the material
[26]
.

Abraham et. al,

has also
point ed out that SED

is a better predictor than
maximum
strain or stre ss

based

predictors

[27]

.

Other than SED, Maximum Principal strain/stretch and octahedral
shear strain can a lso be used as

crack
initiation

predictor for the hyperelastic material.

[28]
. Stress,
apparently

on the other hand

has rarely been used as a fatigue life pa rameter
[29]
.

It has been found that several researchers are taking

the

elastic modulus of the PU as constant (E ~32
MPa) which
was applied
for shear beam and spokes

[7,19]
.
But in reality
,

PU material properties can
become highly nonlinear depending on the ma terial processing routes.
In this study, we have focused
to emphas ize

on
the effect of nonlinearity for PU

stressstrain behaviour

which have

: (1) Same
Elastic

rubber

m odulus
around

standard
300%

strain but
different nonlinear stressstrain behaviou r
(2)
Different
e lastic
modulus at 300% strain but have s i
m ilar

non linearity of stressstrain
;

on the Tire
stiffness

and Durability

analysis.
In this study t he comparative analysis of Mich elin UPTIS

(which is
recent development of Michelin for Passenger Vehicle)
with H
oneyco mb
and Tweel spokes in terms of
Tire stiffness and durability performance

(with the help of SED, Principal strain and Octaherdal shear
strain)

has been
done
, which is ra rely found in the literature. B
ased on the FEM simulation results, we
compared the perf ormance of three different spokes and 5 different nonlinear PU
property and

suggested the optimum combination of PU property an d spoke structure in terms of better durab ility,
s tiffness and riding comfort.

4

2

Modeling & Simulation Details

2.1

CAD Model Details:

A
s mentioned in the introduction section, NPT consists of a Hub, Shear band with having inner and
outer reinforcement rings, Flexible spokes and tread. The detailed description of NPT is mentioned in
the
Figure
1
.

S
ince

none of the companies provide information about materials used and dimensions,
all

executed
numerical models were developed based on articles, photos and press reports which are available
on the
literature

[17]
.
In
this study we selected the dimensions of the
n onpneumatic t yre
’s components
based
on
that

are

available in the lit erature

[7,19]
.

Figure
1
.

Detailed description of Non Pneumatic Tire

We developed

CAD models which depicts three
different NPT

models

i.e (a) Michelin UPTIS (b)
Michelin XTweel and (c
) Res ilient’s Honeycomb

using
FreeCAD

0.18 as shown

in the
Figure
2
.

Then,
the
CAD d r
a wings are
it convert ed

to
‘
.step
’

file
format and import ed

to ANSYS 16.0 for further
analysis.

Here
dimensions of
all the components are similar exce pt spoke st ructure. Spoke
’
s

thickness
and n umber s

have been selected such

that the volume

difference

of all th ree spokes are
less
than 5%

and

the

volume of all spokes
is

almost eq uivalent

and comparable.
Diameter of Hub and outside
dimensions of NPT
are

ta ken as shown in the
Figure
1
. The inplane thickness of NPT is 215 mm.
T
hickness of
the h ub,
i nner and outer reinforcement rings, shear band and treads are 1 mm, 0.5 mm, 15
mm and 5 mm respectively.

5

(a) UPTIS

(b) Twe el

(c)

Honeycomb

2.2

Material propert y and Hyperelastic cu r
ve fitting

Material property for the
h ub,

i nner and outer reinforcement ring, and
tread ha ve

been selected as shown
in
Table
2
.
1
.
[7]
.

Table
2
.
1

Material property of NPT Component s

Components

Materials

Density (k g/m
3
)

E

μ

Hub

7075 T6 Al Alloy

2800

72 G
p a

0.33

Inner and Outer Ring

ANSI 4340 HSS

7800

210 G
p a

0.29

Tread

Rubber

1043

11.9 MP
a

0.49

As a general practice
Polyurethane (PU) is used as a suitable material

for the spokes and

the

shear beam
.
In this study we have ex tracted 5 different PU stressstrain data from Cristina

et al

and Milena et al
.
T
he behaviour of these 5 various PU stressstrain property and their hyper elast ic curve
fitting are as
follows

[30,31]
.

2.2.1

Material property of spokes and shear beam

Case 1)

Having
similar
300% modulus a nd Different
nonlinearity
:

The engineering stressstrain data for 4,
4’
-
diphenyl methane diisocyanate (MDI) and 4,4’
-
dibenzyl
diisocyanate (DBDI) are available in the
Cristina
et al

as shown in the
Figure
3
a
.

The difference in the
MDI and DBDI is degree of crystallinity,
MDI is noncrystallizing and DBDI is crystallizing

(14%
degree of crystallinity)
. Here MDI and DBDI, both have equal ~300 %

rubber

modulus but the curvature
of the engineering s tressstrain curve is different

[31]
.

Figure
2
.

CAD Models of NonPneumatic tire having different
spoke d esign

6

(a)

(b)

(c)

(
d
)

(All Dimensions are in mm)

Figure
3
.

(a,b) 5 different nonlinear PU material property

stress strain Curve
(c)

MDI Curve fitting
using Hyperelastic models (d) FEM simula tion of Uniaxial tensile specimen

7

Cas e 2)

Having Same Linearity and Different 300% Modulus
:

The engineering stressstrain data have been plotted for three different PU

(PU
-
0, PUA12
-
H0 and PUA7
-
H1)

having almost similar curvature (or linearity) but diff erent 300% elastic modulus as shown in

Figure
3
b
. These have different
size of filler particle,

weight

percentage

of hard segment content and
degree of crystallinity as shown in the
Table
2
.
2

[30]
.

Table
2
.
2

Characteristics of PU Nanocomposites

Code

Silic a Nanofiller

Filler size

(nm)

Hard segment
content (
%)

Degree of
crystallinity (%)

PU
-
0

No ne

-

17.8

7.1

PUA12
-
H0

Aerosil 380 (A12)

12

17.4

9.3

PUA7
-
HI

Aerosil R974 (A7)

07

17.7

8.8

2.2.2

Hyperelastic model curve fitting

To capture th ese

nonlinear behavio ur of 5 different PU property, we need to find the hyperelastic
constants of the hyperelastic models as mentioned below.

[32]

1.

MooneyRivlin Hyperelasticity
:

The strain energy function (
𝜓
)

of a MooneyRivlin hyperelastic model can be expressed
as mentioned
in the equation 1.

𝜓
=
𝐶
10

(
𝐼
̅
1
−
3
)
+
𝐶
01

(
𝐼
̅
2
−
3
)
+
𝐶
11

(
𝐼
̅
1
−
3
)
(
𝐼
̅
2
−
3
)
+
𝐶
20

(
𝐼
̅
1
−
3
)
2
+

𝐶
02

(
𝐼
̅
2
−
3
)
2
+
1
𝑑

(
𝐽
−
1
)
2

(1)

Where
𝐼
̅
1
,
2

is the deviatoric first principal invariant, J is the Jacobian and the required input parameters
are C
10
, C
01
, C
11
, C
20
, C
02

(material constants) and d (material incompressibility parameter). The initial
shear modulus is defined as (
𝜇
=
2

(
𝐶
10
+

𝐶
01

)
)

and the initial bulk modulus is defined as
(K = 2/d).

T
he

equation 1
have 5 mat erial constants

and

indicates MooneyR
ivlin 5 parameter model, while
M
ooneyR
ivlin 2 parameter model doesn’t have 3
rd
, 4
th

and 5
th

term, and MooneyRivlin 3 Parameter
m odel doesn’t count 4
th

and 5
th

term.

2.

Yeoh Hyperelastici t
y:

The Yeoh hyperelastic strain energy function
(
𝜓
)
is similar to the MooneyRivlin models described
above except that it is only based on the first deviatoric strain invariant (
𝐼
̅
1
)
. It has th e general form
as mentioned in the equation 2.

8

𝜓
=

∑
𝐶
𝑖
0

(

𝐼
̅
1

−
3
)
𝑖
𝑁
𝑖
=
1
+

∑
1
𝑑
𝑘
𝑁
𝑘
=
1

(

𝐽
−
1
)
2
𝑘

(2)

Where N= order of Yeoh Hyperelastic models (i.e 1,

2,

3).

In

Yeoh 3
rd

order hyperelastic model, required input parameters are

three

material constants (C
10
, C
20
,
C
30
) and three incompressib ility parameters (d
1
, d
2
, d
3
). While, In Yeoh 2
nd

order

hyperelastic model,
required input parameters are two material constant s (C
10
, C
20
) and two incompressibility parameters
(d
1
, d
2
).

The initial shear modulus is defined as µ = 2 C
10

and the initial bulk modulus is defined as K =
2/d
1
.

3.

N
eo Hookean Hyperelasticiy:

The strain energy function (
𝜓
)

for the NeoHookean hyperelastic model is shown in the equation 3.
Where the required input parameters are initial shear modulus of the mater ial (μ) and incompressibility
parameter (d).
T
he initial bulk modulus (K) is defined as K = 2/d.

𝜓
=

𝜇
2

(
𝐼
̅
1
−
3
)
+

1
𝑑

(
𝐽
−
1
)
2

(3)

4.

ArrudaBoyce
Hyperelasticiy:

The strain energy function (
𝜓
) of ArrudaBoyce hyperelastic model is shown in equation 4.

𝜓
=

𝜇

[
1
2

(

𝐼
̅
1
−
3
)
+

1
20

𝜆
𝐿
2

(
𝐼
̅
1

2
−
9
)
+

11
1050

𝜆
𝐿
4

(
𝐼
̅
1

3
−
27
)
+

19
7050

𝜆
𝐿
6

(
𝐼
̅
1

4
−
81
)
+

519
673750

𝜆
𝐿
8

(
𝐼
̅
1

5
−
243
)
]
+

1
𝑑

(
𝐽
2
−
1
2
−
ln
𝐽
)

(4)

Where the required input parameters are
initial shear modulus of the material (μ), limiting network
stretch (λ
L
), and incompressibility parameter (d). The initial bulk modulus (K) is defined sam e as
mentioned in the NeoHookea n hyperelasticity
.

To check the best hyperelastic model curve fitting, we selected above hyper elastic models and applied
to MDI

type PU

stressstrain data.
The required in put parameters
have been calculated using ANSYS
16.0 curve fitting
tool as shown in
the table. As the material is incompressible, the incompressibility
parameter selected as tends to zero value. The curve fitting of these hyperelastic models to the MDI
experimental data
has been plotted in the

figure 3
c
.

9

Table
2
.
3

Value of required input parameters calculated from ANSYS 16.0 curve fitting tool for
MDI experimental data

To
identify

the q uality of the curve fitting, Sum squared error
(or Residual) value and R
2

value have
been calculated as mentioned in the
Error! Reference source not found.
. R
2
value is alm ost e q
ual and
tends to 0.99 for Mooney Rivlin 5 parameter, 3 paramet er, Yeoh
-
3
rd

order, Yeoh 2
nd

order, and ArrudaBoyce model. But If we compare the Sum square error or Residual value then we can identify that
MooneyRivlin 5 parameter model has lowest va lue of residual and hence we have selected the MooneyRivlin 5 param eter model for the rest of the Polyurethane (i.e DBDI, PU
-
0, PUA12
-
H0, and PUA7
-
H1)

curve fitting
. The Material constant s

are mentioned in the
Table
2
.
4
.

To verify the curve fitting
quality of Mooney Rivlin 5 para meter model, FEM uniaxial tensile specimen

(dogbone shape)

has
been designed as shown in fig 3d. Thickness of the specimen has been taken as 2 mm. All the dimension
are shown in the figure
.

SOLID 186 (3D 20 node Hexahedral and tetrahedral element) are use d

to mesh
the model.
Left hand side of the specimen (mentio n as A) has applied fixed support and around 150%
strain is applied to the right side of the specimen (mentioned as B)

as shown in fig 3d. v o
n

M
ises stress
distribution at 150% strain is shown in
the figure.

It can be observed from the
von M
ises stress co ntour

that the stress distribution is very uniform throughout the gauge length of the tensile test specimen.

The
FEM output of uniaxial tension has also been presented in the fig 3c, and it is givi ng exact match with
Model/ Constants

C10

C01

C20

C11

C02

C30

μ

λ
L

Mooney Rivlin 5

-
3.883

9.3616

0.0626

0.0578

0.5533

-

-

Mooney Rivlin 3

-
4.9677

11.306

-

4.9994

-

-

-

-

Mooney Rivlin 2

4.1205

-
5.1963

-

-

-

-

-

-

Yeoh 3
rd

Order

1.8131

-

0.00107

-

-

-
0.011

-

-

Yeoh 2
nd

Order

1.21

-

0.04192

-

-

-

-

-

ArrudaBoyce

-

-

-

-

-

-

3.0412

3.5724

Neo Hookean

-

-

-

-

-

-

5.787

-

10

input MDI engineering stressstrain cur ve. Hence
it can be analysed that the

MooneyRivlin 5 parameter

hyperelastic model

is giving best

curve fit

to capture

PU material

nonlinearity.

Table 2.4 Residual (SSE) and R2 calculation of hy perelastic curve fitting

Hyperelastic Model s

Residual

(SSE
) (
M
Pa)

R
2

Mooney Rivlin 5

6.077 E+06

0.99979

Mooney Rivlin 3

1.729 E+07

0.99941

Mooney Rivlin 2

1.447 E+09

0.95098

Yeoh 3
rd

Order

2.721 E+08

0.99078

Yeoh 2
nd

Order

1.127 E+08

0.99618

ArrudaBoyce

1.421 E+08

0.99518

NeoHook ean

1.972 E+09

0.93471

Table
2
.
4
MooneyRivlin 5 Parameter model's material constants for PU

M
aterial

C10

C01

C20

C11

C02

MDI

-
3.883

9.3616

0.062
6

0.0578

0.5533

DBDI

-
16.
565

30.572

-
0.0281

0.3005

3.7316

PU
-
0

-
24.831

34.049

0.1064

-
0.8767

9.0086

PUA12
-
H0

-
17.014

24.208

0.0431

-
0.3908

5.8038

PUA7
-
H1

-
11.817

17.440

0.0253

-
0.2277

3.9123

2.3

FEM Simulation
Details:

CAD models of NPTUPTIS
,
Tweel and Honeycomb structures have

been imported to ANSYS 16.0
Workbench for the FEM Simulation. All interfaces, the spokes, hub, inner

and

outer reinforcement ring,
shear band, and tread are connected to each other using

bonded contact
.
Road surface and Tread of the
tire has been tied to each

other with rough contact w hi ch won’t allow sliding between tread and road
surface.

F
ixed support constraint

is

applied to the inner face of

the

h ub, which restricts all translational

and
rotation al movement
.
A
50 mm of displacement has been applied normal to

the road surface as shown
in

the
Figure 4
. Th e

5
0

mm of road displacement is applied with

2 mm

increment

with

total 25
gradual

steps to reduce the element distortion and convergence error. SOLID 187 (3D 10 node Tetrahedral
element) used

to model thin reinforcement ri ngs

and SOLID 186 (
3D 20 node

Hexahedral and

11

tetrahedral

element)

are

used in the
spokes, tread, shear band, hub and road surfaces.

TARGE170 and
CONTA174 have been used to identify bonded target and contact interfaces, (Speciality of bonded
contact)
.
Due t o large
deformation and material nonlinearity it was necessary to use

nonlinear method
to solve FEM equations. Therefore, iterative NewtonRaphson

solution was used

in this FEM
simulation
.

Figure
4
.

Boundary conditions description

3

Results & Discussion

3.1

S
train

distribution for three types of spokes

Figure 4 shows vonMises strain distribution of Honeycomb, Tweel and UPTIS spokes having MDI
as
a spoke’s

material

property

and deformed by
50 mm of road displacement. It is observed that the lower
portion of
the spo kes is more strained, indicated by red box.

How ever, the region of maximum von
M
i ses str ain

depends much on the spoke design. For example, for Honeycomb and Tweel, this region
is close to the Tread part, whereas it is close to the rim for UPTIS.

It
has als o been
analysed

that at the
same
degree of road
displacement, the vonMises strain is higher in the UPTIS spoke design compared
to Honeycomb and Tweel structure.
This has also been the general tendency observed among different
spoke designs when loaded wi th vari ous other linear and nonlinear PU properties.

12

(a) Honeycomb

(b) Tweel

(c) UPTIS

Figure
4
.

VonMises strain distribution in the various spoke structure

3.2

Effect of PU material nonLinearity

and spoke design

on Tir e stiffness

and Damage
prediction.

3.2.1

Tire Stiffness:

Fig 5 shows the static load displacement analysis for (a) Linear elastic PU (Having constant elastic
modulus of 32 MPa)
[7,19]

and (b) MDI (Highly nonlinear PU, having different elastic mod ulus at
different percen tage of strain) as discussed earlier in the
F
ig 3a. It is observed that
both
l inear elastic
PU and Nonlinear MDI both have quite
similar
amount of stiffness
.
Although all three spoke design s

have equivalent spoke volume,
UPTIS spoke
design is
less stiff

or
having more cushioning effect
compared to Tweel and Honeycomb spokes.
By comparing the nature of the static load displacement
curve, it can be seen that
T
weel and Honeycomb have almost linear or constant stiffness slope at
h igh er

and lower loading, while U
PTIS has different degree of stiffness at lower and higher loading value
.

T
he
percentage difference in tire stiffness

for the different spoke designs

goes down with load increment.
For

the first 10 mm of displacement
, the overall

tyre stiffness for UPTIS i s
around
4
1
% less than the
Honeycomb and Tweel counterpart.
The stiffness difference goes down to

20
-
30
%
range
among the
spoke designs at hi gher displacement
of 40 mm
.

13

Figure
5

Static load displacement curves
of various spoke des igns
having

(a) Linear Elastic Poly
Urethane (E=32MPa) (b)
NonLinear

Poly Urethane (MDI)

material property for spokes

Fig 6 shows the static load displacement analysis of NPTTweel for (a) MDI, DBDI (ha ving
around
same 300% elastic modulus but different
nonlinearity) and (b) PU
-
0, PUA12
-
H0, and PUA7
-
H1
(Having different 300% elastic modulus
but

similar
nonlinearity) material property applied for spokes
and shear beam while keeping all other material
property as constant. The similar
tire stiffness
beh aviour for
these 5
-
material properties

has also been observed for NPTHoneycomb and NPTUPTIS
as well
.

F
or the sake of simplicity NPTTweel load displacement curve has been
shown

as the stiffness
differen ce in the MDI and DBDI

is more pronounced
. From fig
6(a) it has been illustrated that the nonlinearity of PU property has significant effect on the tire stiffness
. In Fig 6 (b) it can be seen that, As
PU
-
0, PUA12
-
H0 and PUA7
-
H1 have similar linearity bu t having different 300% elastic modulus,
the stiffne ss trend of these three material properties are similar to

th eir material stressstrain nature.

From Fig. 6(a) it is evident that the tyre stiffness difference due to material nonlinearity is almost
1
1
0%, whereas there is virtually no difference of the ru bber modulus at a standard 300% stretch.

In Fig.
6(b) the
difference in tyre stiffness at

under

10 mm displacement

is around

17
%, which is
even little
smaller

to the 2
1
% difference found in the rubber modulus

at a standard 300% stretch shown

in Fig.
3a2.

14

Figure 6 St atic load displacement analysis for NPT having Tweel spokes and having material property
of (a)

MDI, DBDI and (b) PU
-
0, PUA12
-
H0, and PUA7
-
H1
. Effect of tyre stiffness depending on
material properties.

3.2.2

Damage Prediction and Durability performance:

As di scussed in section
2
, SED,
p rincipal strain and
o ctahedral shear strains are used to predict the crack
initiation in the elastomers, which is helpful to analyse the damage prediction of
n on pneumatic tire in
spokes.
[26,28]

SED and
o ct ahedral shear strain can be calculated as described in the following equation 5 & 6
respectively.

SED =
1
2

[σ
x
*
ε
x

+ σ
y
*
ε
y

+ σ
z
*
ε
z

+ σ
xy
*
ε
xy

+ σ
yz
*
ε
yz
+ σ
xz
*

ε
xz
]

(5)

Octahedral shear st rain =
2
3

√
[
(
𝜀
1
−

𝜀
2
)
2
+

(
𝜀
2
−

𝜀
3
)
2
+

(
𝜀
1
−

𝜀
3
)
2
]

(6)

In Fig
7
(a) and
7
(c) SED variation with respect to reaction force in spokes is illustrated for LE
Polyurethane and MDI. Similarly, maximum pr incipal strain va riation in the spokes is

illustrated in the
Fig
7
(b) and Fig
7
(d).

Fig
7

shows that the value of strain energy density (SED) an d max principal strain is lowest

in the Tweel
spoke design and highest in the UPTIS. There is quick rise in the

SED and Max prin cipal strain for the
Honeycomb spoke after around 400 kg of static load,
which

is due to the highly localized elastic
deformation in the honeycomb spokes due to its structur al

design. By comparing SED curve for the
l inear
e lastic PU property

(Fig
7
a) and
n on linear PU (MDI) (fig
7
c), it can be seen
that irrespective of
the material nonlinearity, there can be a large deviation of important damage parameter like SED and
Tweel seems to be the better design compared to the other spoke designs.
W
e

may

infer

that Tweel has
more crack initiation resistance than Honeycomb and UPTIS spoke structure.

15

Having identified Tweel as a design of choice from a damage

initiation

perspective, we
have
investigated the effect of various material properties and selection of da mage param eters like SED,
m a
x imum

p rincipal
strain and
o cta hedral shear strain
.
From Fig 8(a), it can be illustrated that the value
of SED is almost equal till 1000 kg of static load

for NPTTweel

fo r all the 5 material properties
,

and

a t higher
load DBDI

has l e
ss

amount

of SED value than all other

material

properties.

Maximum
principal strain and
m aximum octahedral shear strain variation are shown in the

Fig 8(b) and fig 8(c)
.
MDI and PUA7
-
H1, both have almost similar value of SED,
m ax imum

p rincipal strain and
m ax imum

octahedr al shear s train. If we look at all the crack initiation parameter, we can observe that the DBDI
has more crack initiation resistance and hence good durability performance.
As SED predicts similar
values of damage for all materials until a load of 1000 N, t he other d amage parameters like
m ax imum

s train and
maximum o ctahedral strains appear to be useful to distinguish the material with least damage
factors. Here DBDI becomes the favourable candidate
material
according to this study.

Figure
7

SED &
m ax imum

p rincipal strain vs
r eacti on
f orce for (a, b) LE Polyurethane and (c, d)
MDI
. Effect of tyre spoke designs on damage related factors.

16

Figure
8

Variation

of SED,
m ax imum

principal strain and octahedral shear strain with increasing static
load for
Tweel spokes

In
Fig s.

9

and

10 the effec t of
s poke structure on the damage parameters (i.e. SED and
m ax imum

principal strain) have been compared. From Fig 9c it can be seen that the value of SED for Honeycomb
spoke is almost equal for all fivematerial property.
But there is marginal difference in t he maximum
principal strain for spokes in NPTHoneycomb spoke structure.
Although there is significant variation
in the material property of MDI & PUA7
-
H1, The curves of SED and Max principal strain for MDI
and PUA7
-
H
1 are overlapping to each other.

Fig.

9

shows that the damage parameter SED alone cannot
distinguish between
the
different material properties for damage behaviour for any particular NPT
spoke design.
So
,

in Fig. 10 the other distinguishing damage parameter
m ax imum

p rincipal
s train has
been plotted fo r all 5 material properties applied for all the tyres. This
m ax
p rincipal
s train has much
variability
to
capture

for damage aspects for all spoke design s

than the previously discussed parameter
SED.
From
F
ig
.

10

it can be
seen that the honeycomb spokes have
“
Stype
”

v ariation in the maximum
principal strain with increasing load
,

which is due to their
complex
design
and

is quite different than
Tweel and UPTIS spokes. If we compare Fig 9 and Fig 10, we can say that combination of Tweel with

17

DB
DI has lowest value of SED

and
m a
ximum principal strain,
h ence this is the most suitable design
for
better

durability performance
. However, from damage perspective, the material property of DBDI
appears to be the best

material

candidate

among the 5
materials

for all the s poke de signs involved.

Figure
9

SED variation with increasing static load having
variety

of PU property for (a) Tweel (b)
UPTIS and (c) Honeycomb spokes

18

Figure 10 Max principal strain variation with increasing static load having variety of PU property
for
(a) Tweel (b) UPTIS and (c) Honeycomb spokes

4

Conclusion:

FEM models of three different spokes structures (Tweel, Honeycomb and UPTIS) have been developed.
Effects of the
n onlinearity of PU materials (Linear elastic, MDI, DBDI, PU
-
0, PUA12
-
H0,
PUA7
-
H
1) properties of spokes have been studied with
a
series of

FEM

simulation s

carried out in ANSYS
16.0. The major findings of th is

stud y

are as follows.

1.

Mooney Rivlin 5 parameter model is
best

to capture all
five
types of PU hyperelastic
behaviour

lik e the n onlinearity and rubber modulus changes based on the 5 d i
fferent

non linear

stressstrain graphs

from

the

materials
related to

this study
.

2.

The tyre spoke designs have some degree of effects on the overall tyre stiffness found in the
forcedisplacement curves. But

the di fference of overall tyre stiffness goes down with
increasing load.

19

3.

NPTUPTIS

has softer cushioning effects than the
NPTHoneycomb

and
NPTTweel for both
linear and nonlinear material properties

for spoke structure.

4.

Nonlinearity

of PUstress strain curv e
has a larger impact than the
300% r ubber modulus

(E)

on

the overall tyre stiffness
which is

observed f or

all

three

spoke designs
.

5.

Damage p redictors

like max principal strain and SED are
significantly

lower in
NPTTweel
than
NPTUP
TIS

and
NPTHoneycomb. DB
DI

material

has lowest value of SED and Max
principal strain in all there spoke structure
.

H
ence

it is the best choice for having Tweel with
DBDI spoke material, if tire designer wants good tire stiffness and durability performance
.

However, if the importa nce is given more on the comfort of driving and more cushioning
effect

needed, then NPTUPTIS
can

be a better choice
.

6.

Damage pr edictors

like max imum

principal strain and octahedral strain have more
differentiating effect than
strain energy density
. S
ED has very similar effec t for all nonlinear
material property for loading
up to

1000 kg.
But m ax imum

principal strain

has more effect
a s
a damage initiation factor
on

all

nonlinear

material s
.

B
ased on that it can be concluded that
DBDI has highest resistance to crack initiation in the
NPT
s poke structure.

7.

M
axprinc ipal strain and max Octahedral shear strains have larger dependence on the material
properties. So
,

these parameters rather may be taken as more important damage parameters for
damage related performance prediction.

5

References

1.

Ma, J., Summers, J., and Joseph, P., Dynamic Impact Simulation of Interaction between NonPneumatic Tire and Sand with Obstacle, 2011, doi:10.4271/2011
-
01
-
0184.

2.

Cho, J.R., Kim, K.W., Yoo, W.S.,
and Hong, S.I., “Mesh generation cons idering detailed tread
blocks for reliable 3D tire analysis,”
Adv. Eng. Softw.

35(2):105
–
113, 2004,
doi:10.1016/j.advengsoft.2003.10.002.

3.

Rhyne, T.B. and Cron, S.M., “Development of a nonpneumatic wheel,”
Tire Sci.
Technol.

34(3):150
–
169, 2006.

4.

Rogu e, A., “U.S. Patent No. 3329192,” 3
–
4, 1965.

5.

Kubica, W. ; Schmidt, O., “U.S. Patent No. 4169494,” 2
–
5, 1979.

6.

Manesh, A., Terchea, M., Anderson, B., Meliska, B., and Ceranski, F., “Tensionbased nonpneumatic tire,
”
World Intellect. Prop. Organ. WO

118983:A1, 2008.

7.

Ju, J., Kim, D.M., and Kim,

K., “Flexible cellular solid spokes of a nonpneumatic tire,”
Compos. Struct.

94(8):2285
–
2295, 2012, doi:10.1016/j.compstruct.2011.12.022.

8.

Kim, K., Heo, H., Uddin, M.S.,
Ju, J., and Kim, D.M., “Optimization of Nonpneumatic Tire

20

with Hexagonal Lattice S
pokes for Reducing Rolling Resistance,”
SAE Tech. Pap.

2015
–
April(April), 2015, doi:10.4271/2015
-
01
-
1515.

9.

Zmuda, M., Jackowski, J., and Hryciów, Z., “Numerical research of

selected features of the nonpneumatic tire,”
AIP Conf. Proc.

2078(March), 2019,
doi:10.1063/1.5092030.

10.

Narasimhan, A., Ziegert, J., and Thompson, L., “Effects of Material Properties on Static LoadDeflection and Vibration of a NonPneumatic Tire During HighSpeed Rolling,”
SAE Int. J.
Passeng. CarsMech. Syst.

4(1):59
–
72, 2011,
doi:10.4271/2011
-
01
-
0101.

1
1.

Sassi, S., Ebrahemi, M., and AlMozien, M., “New design of flatproof nonpneumatic tire,”
Int
J Mech Syst Eng

2:114, 2016.

12.

Bridgestone Corporation, “Air Free Concept (NonPneumatic) Tire,”
Bridgestone
, 2013.

13.

Boyer, M., “Bridgestone AirFree Con cept Tyre,” 2011.

14.

Golson, J., “Michelin develops new airless tyre,”
Auto Express
, 2019.

15.

Ingrole, A., Hao, A., and Liang, R., “Design and modeling of auxetic and hybrid honeycomb
structures for inplane proper ty enhancement,”
Mater. Des.

117:72
–
83,
2017,
doi:10.1016/j.matdes.2016.12.067.

16.

Jin, X., Hou, C., Fan, X., Sun, Y., Lv, J., and Lu, C., “Investigation on the static and dynamic
behaviors of nonpneumatic tires with honeycomb spokes,”
Compos. Struct.

18
7(April
2017):27
–
35, 2018, doi:10.1016/j
.compstruct.2017.12.044.

17.

Małachowski, M.K.P.B.J., “Proceedings of the 13th International Scientific Conference:
Computer Aided Engineering,”
Lect. Notes Mech. Eng.

293
–
301, 2017, doi:10.1007/978
-
3
-
319
-
50938
-
9.

18
.

Bezgam, S., “Design and Analysis of Alternating Spoke Pair Concepts

for a NonPneumatic
Tire with Reduced Vibration at High Speed Rolling,” 2009.

19.

AboulYazid, A.M., Emam, M.A.A., Shaaban, S., and ElNashar, M.A., “Effect of spokes
structures on chara cteristics performance of nonpneumatic tires,”
Int. J. Automot. Mech
. Eng.

11(1):2212
–
2223, 2015, doi:10.15282/ijame.11.2015.4.0185.

20.

Berglind, L., Ju, J., and Summers, J.D., “Method to design honeycombs for a shear flexible
structure,”
SAE Int. J. Pas seng. CarsMechanical Syst.

3(2010
-
1
–
762):588
–
597, 2010.

21.

Abd ElS
ayed, F.K., Jones, R., and Burgess, I.W., “A theoretical approach to the deformation
of honeycomb based composite materials,”
Composites

10(4):209
–
214, 1979.

22.

Gibson, L.J., Ashby, M.F.
, Schajer, G.S., and Robertson, C.I., “The mechanics of twodimension al cellular materials,”
Proc. R. Soc. London. A. Math. Phys. Sci.

382(1782):25
–
42,
1982.

23.

Masters, I.G. and Evans, K.E., “Models for the elastic deformation of honeycombs,”
Compos.
Str uct.

35(4):403
–
422, 1996.

24.

Rugsaj, R. and Suvanjumrat, C., “Proper

radial spokes of nonpneumatic tire for vertical load
supporting by finite element analysis,”
Int. J. Automot. Technol.

20(4):801
–
812, 2019,
doi:10.1007/s12239
-
019
-
0075
-
y.

25.

Veeramurth y, M., Ju, J., Thompson, L.L., and Summers, J.D., “Optimisation of ge ometry and
material properties of a nonpneumatic tyre for reducing rolling resistance,”
Int. J. Veh. Des.

66(2):193
–
216, 2014, doi:10.1504/IJVD.2014.064567.

26.

Mars, W. V. and Fatemi, A
., “A literature survey on fatigue analysis approaches for rubber,”
I
nt. J. Fatigue

24(9):949
–
961, 2002, doi:10.1016/S0142
-
1123(02)00008
-
7.

21

27.

Abraham, F., Alshuth, T., and Jerrams, S., “The effect of minimum stress and stress amplitude
on the fatigue lif e of non strain crystallising elastomers,”
Mater. Des.

26(3):239
–
245,

2005,
doi:10.1016/j.matdes.2004.02.020.

28.

Previati, G. and Kaliske, M., “Crack propagation in pneumatic tires: Continuum mechanics and
fracture mechanics approaches,”
Int. J. Fatigue

3
7:69
–
78, 2012,
doi:10.1016/j.ijfatigue.2011.10.002.

29.

André, N., Ca illetaud, G., and Piques, R., “Haigh diagram for fatigue crack initiation prediction
of natural rubber components,”
Kautschuk Gummi Kunststoffe

52(2):120
–
123, 1999.

30.

Špírková, M., Porę
ba, R., and Brožová, L., “The Influence of the Size , Shape and Chara cter of
Nanofillers on Functional Properties of Polyurethane Elastomers,” 2013.

31.

Prisacariu, C., “Polyurethane elastomers: from morphology to mechanical aspects,” Springer
Science & Bu siness Media, ISBN 3709105145, 2011.

32.

Ansys

16.0
, “Ansys Documenta tion, Mechanical User’s Guide.”

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
