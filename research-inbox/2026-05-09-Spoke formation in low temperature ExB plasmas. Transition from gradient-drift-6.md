---
source: "https://arxiv.org/abs/2211.17051v1"
title: "Spoke formation in low temperature ExB plasmas. Transition from gradient-drift instability to ionization wave"
author: "J P Boeuf"
year: "2022"
publication: "arXiv preprint / physics.plasm-ph"
download: "https://arxiv.org/pdf/2211.17051v1"
pdf: "https://arxiv.org/pdf/2211.17051v1"
captured_at: "2026-05-09T13:04:02Z"
updated_at: "2026-05-09T13:04:02Z"
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

# Spoke formation in low temperature ExB plasmas. Transition from gradient-drift instability to ionization wave

- 著者: J P Boeuf
- 年: 2022
- 掲載情報: arXiv preprint / physics.plasm-ph
- 情報源: [arxiv](https://arxiv.org/abs/2211.17051v1)
- ダウンロード: https://arxiv.org/pdf/2211.17051v1
- PDF: https://arxiv.org/pdf/2211.17051v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Long wavelength plasma non-uniformities rotating in the azimuthal direction ("rotating spokes") have been observed in a number of experiments on Hall thrusters or magnetron discharges. We use a two-dimensional (2D), axial-azimuthal Particle-In-Cell Monte Carlo Collisions (PIC-MCC) model to study the formation of instabilities in a direct current (dc) magnetron discharge under conditions close to recent experiments. In spite of the simplified 2D geometry of the model, the simulations can reproduce the main features of the experimental results. At a given position above the cathode, corresponding to the spoke location, the simulations show large amplitude oscillations of the plasma density and a very sharp increase of the plasma potential and electron temperature at the leading edge of the spoke, as in time resolved probe measurements. Moreover, the simulations show that the instability evolves in time from a gradient-drift type of instability in the linear phase, to an ionization wave in the non-linear phase, with rotation in the +ExB direction in the first phase and in the -ExB direction in the second phase. The number of spokes is found to increase with pressure, as in experiments. The mechanisms of electron heating and the role of the Bx$\nabla$B drift in electron heating and in the coherence and direction of spoke rotation are discussed.

## PDF Text

1

Spoke formation in low temperature

𝑬
×
𝑩

plasmas
.

Transition from gradientdrift instability to ionization wave

J.P. Boeuf
a

LAPLACE, Université de Toulouse, CNRS, INPT, UPS, 118 Route de Narbonne, 31062 Toulouse, France

ABSTRACT

Long wavelength p lasma nonuniformities

rotating in the azimuthal direction

(“rotating spokes”)
have been observed
in a number of experiments

on Hall thrusters or magnetron discharges
. We use a twodimensional

(2D)
, axialazimuthal
ParticleInCell Monte Carlo Collisions (PICMCC)
model

to stud y

the formation of instabilities

in a
direct
current (dc)
magnetron discharge

under conditions close to

recent experiments.

In spite of the simplified 2D geometry
of the model, the simulations

can reproduce the main

features of the

experimental results.
At a given position above
the cathode
,

corresponding to the spoke location, the simulations show large amplitude oscillations of the plasma
density and
a
very sharp increase of the plasma potent ial and electron temperature at the leading edge of the spoke
, as
in
time resolved
probe measurements.

Moreover, the simulations show that
the instability evolves in time from a
gradientdrift type

of instability in the linear phase, to an ionization wave
in the nonlinear phase, with rotation in the
+
𝑬
×
𝑩

direction in the first phase and in the
−
𝑬
×
𝑩

direction in the second phase.

The number of spokes
is found
to increase

with pressure
,

as in experiments.

The mechanisms of electron heating and the role of
the
𝑩
×
𝛁
𝑩

drift
in
electron heating and in the coherence
and
direction

of spoke

rotation

are discu ssed.

I.

INTRODUCTION

Hall thrusters, magnetron
discharges, or Penning
discharges
1
-
8
, are low temperature, partially magnetized
plasma devices where electrons are strongly magnetized
while ions are weakly magnetized.
The electron Larmor
radius is smaller than the characteris tic dimensions of the
plasma while the ion Larmor radius is not.
An external
magnetic field is placed perpendicular to the applied electric
field or to the discharge current to increase the electron
residence time in the device and allow ionization and pl asma
sustainment at low pressures (in the
1
-
50
mtorr range
)

and in
relatively small devices (dimensions of a few cms). Electron
confinement is efficient only if the
𝑬
×
𝑩

current is closed
on itself
,
otherwise
the

confinement would be destroyed by
the Hal l effect. The geometry of these devices is therefore
cylindrical, with
𝑬
×
𝑩

in the azimuthal direction, i.e.
𝑬

axial
and
𝑩

radial or the opposite.

Electrons are trapped by the
magnetic field and are collisional while ions can be freely
accelerated by the electric field.
In regions with large electric
fields, the electron drift velocity in the azimuthal,
𝑬
×
𝑩

direction
,

is E/B while ions
are acceler ated out of these
regions before they can acquire the E/B azimuthal velocity
(this is another way to say that these plasmas are partially
magnetized).

The large difference between the azimuthal drift velocit ies

of electrons and ions can generate charge separation which

a

Electronic mail:
jpb@laplace.univtlse.fr

can

lead to the development of instabilities. One of these
instabilities develops when the electric field and the
plasma
density gradient are in the same direction
9

(
gradientdrift or
SimonHoh instability)
. This situation is
very common in

𝑬
×
𝑩

plasma devices such as
magnetron discharges
3
,
4
,
10

or
in

the channel of a Hall thruster
2
.

Other types of instabilities
can develop in regions where the electric field and plasma
density gradients are in opposite direction
, as in the
exhaust
region of a Hall thruster where particle models have shown
the formation of an Electron Cyclotron Drift Instability
11
-
14
.

In this paper we
study
, with the h elp of a 2D PICMCC
model,

the development of instabilities in a direct current (dc)
magnetron discharge, under conditions similar to those of the
experimental paper s

of Held et al.
15

and of Panjan et al.
16
,
17
.

The inst abilities in magnetron discharges
take the form of

largescale plasma nonuniformities rotating in the azimuthal
direction and associated with rotating
regions of enhanced
light emission and ionization rate.

These rotating ionization
regions are generally called “spokes” and were first
mentioned in the context of Hall thrusters by Janes and
Lowder
18
. They have been more recently evidenced by high
speed camera imaging in Hall thrus ters
19
-
24

as well as in
magnetron discharges
25
-
32
. An interesting feature of the spoke
rotation discussed in several papers
17
,
28
,
31
-
33

is
that the spoke
can rotate in the

−
𝑬
×
𝑩

direction in low power dc
magnetrons and in the
+
𝑬
×
𝑩

direction in high power

2

pulsed magnetrons used in plasma processing applications
(HiPIMS, High Power Impulse Magnetron Sputtering
5
).

ParticleInCell Monte Carlo Collisions (PICMCC) are
now widely used to study the

physics and instabilities in Hall
thrusters
11
-
14
,
34
,
35
, magnetron

discharges
36
-
40
, and Penning
discharges
41
-
43
.

Twodimensional (2D) ParticleInCell
Monte Carlo Collisions (PICMCC) simulations of
magnetron discharges by Boeuf and Takahashi
38
,
39

performed
under conditions of dc magnetron discharges similar to those
of the experiment of Ito et al.
31

were able to

qualitatively
reproduce the
formation of spokes

and the
−
𝑬
×
𝑩

rotation.

In the experiments

of

Held al.
15

that are simulated in the
present paper,

a

Langmuir probe
has been used
to measure
the plasma

density, plasma potential, electron temperature
and
e lectron energy probability function (EEPF) at a given
location above the cathode, on the spoke’s path, and as a
function of time. The measurements show sharp increases of
the plasma potential, electron
density and plasma density
when the spoke front reaches the probe, followed by a slower
decay. The plasma properties present large amplitude
oscillations as a function of time in the spoke region. The tail
of the EEPF presents a sharp increase at the spoke

front,
followed by a quick decay. The large
potential drop at the
spoke front had been
reported
in previous
experimental works
on magnetron discharges
3
,
17
,
32
.

The goal of the present

paper is to complement the work
of Refs.
38
,
39

by more detailed (but still qualitative)
comparisons wi th the experiments

of Held et al.
15

and
Panjan
et al.
16
,
17
. We compare the PICMCC simulations results with
the probe measurements of Held et al., and the effect of
pressure on the number of spokes with the exper imental
results of Panjan et al.. We also describe some interesting
results of the particle simulations, which show
that,

starting
from a 1D PICMCC stable solution of the

dc

magnetron
discharge, the formation of the instability

is consistent with
the fl uid theory and dispersion relation

of gradient drift
instabilities
9

in the linear phase but that

the instability
evolves
into

an ionization wave in the nonlinear
stage
. The phase
velocity of the
instability changes direction from a
+
𝑬
×
𝑩

direction in the linear phase, to a
−
𝑬
×
𝑩

direction in the nonlinear phase.

The PIC MCC model is briefly described in section II
. T
he

results are presented and compared to the experimental
results in section III
.

II.

PICMCC M
ODEL

AND CONDITIONS OF THE
SIMULATIONS

The
PICMCC

model is the same as in Refs.
38
,
39
.

The simulations are performed in a 2D Cartesian
geometry

in the plane defined by the axial and azimuthal
direction s

of the magnetron

(see
Figure
1
).

T
he curvature of
the magnetic field
is not

taken into account

in this simplified
geometry
. The magnetic field is perpendicular to the
simulation domain and is supposed to depend
on the axial
coordinate only. Periodic boundary conditions are used in the
azimuthal direction.

Figure
1
: (a)
Twodimensional, axialazimuthal s imulation domain
.
The magnetic field is perpendicular to the simulation domain
; (b)
axial distribution of th e perpendicular magnetic field

of the form
𝐵
(
𝑥
)
=
𝐵
𝑚𝑎𝑥
exp
⁡
(
−
𝑥
𝑎
)
⁄
,
with

𝐵
𝑚𝑎𝑥
=
120
⁡
mT

and
⁡
𝑎
=
0
.
3
⁡
𝑑
𝑥

unless otherwise mentioned
;
⁡
⁡
𝑑
𝑥
=
2
⁡

cm
is the cathodeanode gap
.

𝑗
𝑒
0
⁡
is

the

net

current density

of secondary electrons emitted by the
cathode due to ion impact

and is 0.03 A/m
2

in all the simulation s

(uniform along the cathode surface).

The applied voltage is
𝑉
=
260
⁡

V in most simulations.

The positive
𝑦

direction is downward.
The applied field
𝑬
⁡
is in the
−
𝑥

direction, the magnetic field
𝑩

in
the
+
𝑧

direction and
𝑬
×
𝑩

is in the
+
𝑦

direction.

The 1D version of this 2D PICMCC model (JPIC code,
see Ref.
44
) has also been used to s tudy the 1D axial solutions
of the same magnetron discharges, and to provide initial
conditions for the 2D simulations.

The cathodeanode gap is
𝑑
𝑥
=
2

cm and the length of the
simulation domain in the periodic azimuthal direction is
𝑑
𝑦
=
8 cm. A voltage o f 260 V is applied between cathode
and anode unless otherwise mentioned. The axial profile of
the perpendicular magnetic field is given by:

𝐵
(
𝑥
)
=
𝐵
𝑚𝑎𝑥
exp
⁡
(
−
𝑥
𝑎
)
⁄
⁡

with
𝐵
𝑚𝑎𝑥

=

120 mT and
𝑎
=
0
.
3
⁡
𝑑
𝑥

in the simulations
presented

in section III.

Secondary emission due to ion impact is simulated by
imposing a uniform net current density
𝑗
𝑒
0

of electrons
emitted by the cathode, as in Ref.
39

(“net” current density
means that
each electron coming back to the cathode is reemitted).

The corresponding net secondary electron emission
coefficient under ion impact,
γ
𝑛𝑒𝑡

, can be
deduced

a
posteriori

(ratio o f

the ion current at the cathode to the
net
electron current emitted by t he cathode). For a given, fixed
current density
𝑗
𝑒
0

the secondary emission coefficient
can

therefore depend on the conditions (pressure, voltage
magnetic field), in a n

unr ealistic way.
It is clear that care must
be taken

in
the analysis of
parameter studies performed with
a constant

𝑗
𝑒
0

instead of a constant
γ
𝑛𝑒𝑡

but this method
allows to limit the plasma density to values t hat can be
handled
by the particle model.

The simulations are performed in argon, at pressure
between 0.6 Pa (
clo se to the pressure in

Held et al.
15
)
, and 2.
Pa.

As in Ref.
44
,

t he electron
–
neutral cross sections used in
the simulations are taken from the siglo database
45

of LXCat
46

(compilation of Phelps where the electron excitation levels

3

are grouped in one single level),
and
t he isotropic and
backward scattering ion
–
neutral crosssections for argon are
taken from Phelps
47
.

The number of (x,y) grid intervals is either (128
×
512) or
(256
×
1024) depending on the plasma density (the grid
spacing must satisfy
𝛿𝑥
,
𝛿
𝑦
<
𝜆
𝐷𝑒
, where
𝜆
𝐷𝑒

is the electron
Debye length). The time step is such that
𝛿𝑡
<
0
.
2
/
𝜔
𝑝𝑒

and
must be sufficiently smaller than the minimum cyclotron
rotation time
⁡
around the magnetic field
.
𝛿𝑡
=
2
×
10
−
11

s
wa s used in most simulations
.

The number
of particles per
cell was between 30 and 70 in the
2D simulations presented
in section III
.

O
pen
MP

(Open MultiProcessing)

parallelization
with particle decomposition is used and the
simulations presented here have been run on 12

or 24

threads.

To study t he formation of instabilit ies
, we first perform a
1D PICMCC simulation in the conditions above. The 1D
solutions
are

stable for the conditions considered in this paper

(small amplitude axial plasma oscillations can be present
under

some conditions, depending on the magnetic field
profile, plasma density, etc…)
. We then use the 1D solution
as an initial condition for the 2D simulation
. This means that
the
initial
axial
po sitions and velocities

of the
charged
particles in the 2D simulations are identical to those of the 1D
steady state

solution
; the initial azimuthal positions are
chosen randomly according to a uniform distribution
. This
allows the study of the linear devel opment of the instability
around the unperturbed
, stable, steady state

1D solutions.

III.

SIMULATION RESULTS AND COMPARISONS WITH
EXPERIMENTS

In the first part of this section (III.A) we show the results
of 1D PICMCC simulations of magnetron discharges in the
conditions described in section II. We then show (section
III.B) the formation of instabilities in 2D axialazimuthal
simulations with initial conditions corresponding to the 1D
stable solutions
, and the transition to the spoke (ionization
wave) regime.
Th e linear phase of the instability is discussed
in section III.C.
The spoke regime is described in more detail
in section III.
D
.
Comparisons with probe measurements are
presented in section III.
E

while consideration s

about spoke
velocity, electron heating,
and the effect of pressure are
discussed in sections III.
F
,
G
, and
H

respectively.

A.

1D PICMCC simulations

Figure
2

shows the axial distributions, at steady state, of
the charged particle densities, electric field, electron
temperature and ionization rate obtained with the 1D PICMCC model in argon at a pr essure of 0.6 Pa, cathodeanode
gap
⁡
𝑑
𝑥
=2

cm, applied voltage
𝑉
=260 V and net current
density of electron emitted by at the cathode

due to ion
impact

𝑗
𝑒
0

=
0.03
A/m
2
. The perpendicular magnetic field
profile is given in
Figure
2
, with
𝐵
𝑚𝑎𝑥
=
120 mT.

The total
electron current density at the anode (close to the discharge
current
density)
de d
uced from the simulation
is
𝑗
𝑒
𝑑
=
5.1
A/m
2
. This gives an electron multiplication
𝑀
=
𝑗
𝑒𝑑
𝑗
𝑒
0
⁄
=
170

and a net secondary electron emission coefficient
γ
𝑛𝑒𝑡
=
1
(
𝑀
−
1
)
⁄
=
6
×
10
−
3
.

Figure
2

: Profiles of the e lectron density, ion density, electric field
and ionization rate
from

the

1D PICMCC
model

in argon at a
pressure
𝑝
=
0.6 Pa and with a cathodeanode gap
⁡
𝑑
𝑥
=2

cm

and an
applied voltage
𝑉
=260 V.

The
net electron density at the cathode
is
𝑗
𝑒
0
=

0.03
A/m
2
.

The axial profile of the
perpendicular
magnetic
field is shown in
Figure
1
b

and

the

maximum magnetic
field

at the
cathode surface

is

𝐵
𝑚𝑎𝑥
=
120 mT
.

Assuming that the probability of reflection of an electron
coming back to the cathode to be reemitted is 0.1, a net
secondary emission coefficient of
6
×
10
−
3

corresponds to a
“real” secondary electron emission by ion impact
γ
=
6
×
10
−
2
, which is a reasonable

order of magnitude for argon
ions on a metal cathode
48
.

We see on
Figure
2

the cathode sheath and quasineutral
plasma. Because of the large perpendicular magnetic field
and low electron mobility there is no negative glow and the
electric field is lar ge in the quasineutral plasma. The electron
temperature is large, on the order of 10 eV over the whole
gap, and the location of maximum ionization rate is closer to
the anode than to the cathode. The plasma density gradient on
the anode side of the maximum

plasma density is negative,
i.e. of the same sign as the electric field. This suggests that
gradient drift
(SimonHoh)
types of instabilities can form in
this region if the azimuthal direction is taken into account.

B.

2D PICMCC simulations
–

Tr ansition from gradient
drift instability to ionization wave

As said above, the 2D simulation s

are

started with initial
conditions c orresponding to the 1D stable solution

described
in the previous subsection
.
This allows to study the
formation of the 2D azimuthal instability around the stable
solution, i.e. in the linear phase, and to follow the evolution
from the linear stage to the
nonlinear, saturated stage.

4

Figure
3
:
Color contour plots of the t im e evolution of

the
ion density

(linear scale)

from the 2D axialazimuthal PICMCC simulation in the
conditions of
Figure
2

and with the magnetic profile of
Figure
1
b ,
𝐵
𝑚𝑎𝑥
=
120 mT; argon, 0.6 Pa, 260 V dc, dimensions (2 cm
×

8 cm),
𝑗
𝑒
0

=
0.03
A/m
2
, (256
×
1024) grid intervals, time step
𝛿𝑡
=
2
×
10
−
11

s; initial conditions from the 1D simulation resu lts of
Figure
2
, with
charged particle uniformly distributed along the azimuthal direction.

The

time is indicated
above

each contour plot, and the
unit

(in 10
16

m
-
3
)
i s given
in

a white
rectangle

inside each plot.

Figure
3

shows such evolution in the conditions of
Figure
2
. We see that the instability forms in about 1

s. T
he
wavelength of the instability in the early stage of its
development is
slightly larger than

1 cm and increases in time
until one single plasma nonuniformity is present in the
azimuthal direction (at about 1
0


s
)
.

During the time interval [1
-
5]

s

the nonuniformities
move in the

+
𝑬
×
𝑩

direction with a velocity around 8 km/s
(the sound speed or argon ion for an electron temperature of
10 eV
–

see
Figure
2
, is

close to 4.
9

km/s). During this time
the growth of the plasma density is not significant. After time
t=
4
-
5

s

a welldefined region of larger plasma density, that
we will call “spoke” in the rest of this paper has formed and
the growth of the plasma densi ty due to ionization becomes
significant till time t=
40

s
. Between t=
4

s

and t=
8

s

the
spoke velocity goes through zero and one can observe a
reversal of the direction of the spoke velocity from the
+
𝑬
×
𝑩

to the
−
𝑬
×
𝑩

direction. After t=
8

s

the spoke
moves in the
−
𝑬
×
𝑩

at a constant velocity slightly smaller
than 1.5 km/s and witho ut significant change in its shape and
maximum plasma density. The motion of the spoke is very
regular and coherent after that time. The simulation has been
performed up to t=
200

s

without any significant change in
the spoke properties and velocity.

The spoke corresponds to the nonlinear, saturated stage of
the instability.
We will see below that the nonlinear stage of
the instability after t=
8

s

corresponds to an ionization wave
.
Before studying in detail the spoke properties

in the
ionization wave regime
, we briefly discuss
, in the next subsection,

the

mechanisms responsible for the instability and the

evolution

of the instability from short wavelengths to larger
wavelengths be fore the spoke formation.

C.

Linear stage of the instability
–

Dispersion relation

F
luid dispersion relation s of low temperature partially
magnetized
𝑬
×
𝑩

plasma are

described in the article by
Smolyakov et al.
9

(see also Refs.

43
,
49
,
50
).

We consider purely
azimuthal perturbations of the density and

potential of the
form
exp
⁡
(
−
𝑖𝜔𝑡
+
𝑖𝑘𝑦
)

, where

𝜔

is the complex wave
frequency

(
𝜔
=
𝜔
𝑅
+
𝑖𝛾
)
,
𝑘

is the wave number in the

5

azimuthal direction.

𝜔
𝑅
, the real part of
𝜔

is the wave
frequency and
𝛾

is the growth rate (the perturbation grows
when
𝛾
>
0
).

T
aking into account
𝑬
×
𝑩

drift,

diamagnetic drift
,

magnetic
field gradient, and including electron inertia, gyroviscosity,
electronneutral collisions and the effect of finite Debye
length
,

the flui d dispersion relation
can be written as
9
,
43
,
49

(
assuming constant electron t emperature
)
:

𝑘
2
𝜆
𝐷𝑒
2
+
𝜔
∗
−
𝜔
D
+
𝑘
2
𝜌
𝑒
2
(
𝜔
−
𝜔
0
+
𝑖
𝜈
𝑒𝑛
)
𝜔
−
𝜔
0
−
𝜔
D
+
𝑘
2
𝜌
𝑒
2
(
𝜔
−
𝜔
0
+
𝑖
𝜈
𝑒𝑛
)
=
𝑘
2
𝑐
𝑠
2
𝜔
2

(1)

where
𝑘

is the wave number in the azimuthal direction,
𝜌
𝑒
=
(
𝑒
T
e
𝑚
⁄
)
1
/
2
𝜔
𝑐𝑒
⁄

is the electron Larmor radius

(
𝜔
𝑐𝑒
=
𝑒𝐵
𝑚
⁡
⁄
,
𝑚

is the electron mass)
,
𝜔
∗
=
𝑘
𝑣
∗
,
𝑣
∗
=
−
𝑇
𝑒
(
𝐵𝐿
𝑛
)
⁄

is
the diamagnetic drift velocity and

𝐿
𝑛
=
[
𝜕
(
ln
𝑛
0
)
/
𝜕𝑥
]
−
1
,
𝜔
0
=
𝑘
𝑣
0
,
𝑣
0
=
−
𝐸
0
𝐵
⁄

is
the

𝑬
×
𝑩

drift
velocity,
⁡
𝑐
𝑠
=
(
𝑒
𝑇
𝑒
𝑀
⁄
)
1
/
2

is the ion sound speed
(
𝑀

is the
ion mass)
and
𝜈
𝑒𝑛

is the electronneutral collision frequency.
𝜔
D

is related to the magnetic field gradient:

𝜔
D
=
𝑘
𝑣
D

with
𝑣
D
=
2
𝑇
𝑒
(
𝐵𝐿
𝐵
)
⁄
⁡
and
𝐿
𝐵
=
[
𝜕
(
ln
⁡
𝐵
)
/
𝜕𝑥
]
−
1
. The index
“zero” in
𝑛
0

and
𝐸
0

indicates equilibrium values (i.e. related
to 1D results used as initial
conditions
).
𝑛
0

and
𝐸
0

depend on
the axial position
𝑥
). Note that with our notations

(see
Figure
1

and
Figure
2
),

𝑬
𝟎
=
𝐸
0
𝒙
̂
,

𝑩
=
𝐵
𝒛
̂
,
𝑬
0
×
𝑩
𝐵
2
=
−
𝐸
0
𝐵
𝒚
̂

, so

𝐸
0

is
negative
,
𝐵

is positive
,

𝐿
𝐵

is negative,
and

𝑣
0
,
𝑣
D
,
𝜔
0
and
𝜔
D

are
positive
;
𝐿
𝑛

is negative
and
𝑣
∗

and
𝜔
∗

are
positive
on
the right
side
of the maximum plasma density in
Figure
2
,
and
change sign

on the left

side
.

The

classical collisionless

SimonHoh instability is
obtained in the limit of small values of
𝑘𝜌
𝑒

(i.e. long
wavelength)

and small Debye length s
.
This gives:

𝜔
∗
−
𝜔
D
𝜔
−
𝜔
0
=
𝑘
2
𝑐
𝑠
2
𝜔
2

(2)

For small wave frequencies,
ω
≪
𝜔
0
⁡
, this equation has
positive complex solutions, i.e. the instability is growing,
⁡
when
(
𝜔
∗
−
𝜔
D
)
𝜔
0
⁄
>
0
.
⁡

This corresponds to conditions
where the
electric field and the plasma density gradient are in
the same direction

and

𝜔
∗
>
𝜔
D
. This is the case in the stable
1D solution of
Figure
2
,
sufficiently close to the anode where
we
have

1
|
𝐿
𝑛
|
>
2
|
𝐿
𝐵
|

(
|
𝐿
𝑛
|
⁡
decreases close to the anode and
|
𝐿
𝐵
|
⁡
is con stant).

Further away from the anode
|
𝐿
𝑛
|

increases,
(
𝜔
∗
−
𝜔
D
)

becomes negative and the SimonHoh
mode disappears

(the magnetic field gradient stabilizes the
SimonHoh mode)
.

To illustrate this,
w e can solve
the dispersion relation

(1)

for the plasma paramet ers corresponding to a position close
to the anode,
⁡
𝑥
=
1
.
8
⁡
cm
.

T
he parameters corresponding to

𝑥
=
1
.
8
⁡
cm

can be obtained from the data of

Figure
2
:

𝑛
0
=
0
.
8
×
10
15
⁡
m
−
3
,
𝐸
0
=
−
2200
⁡
V
m
⁄
,
𝐵
=
6
⁡
mT
,
𝐿
𝑛
=
−
2
.
2
⁡
mm
,
𝐿
𝐵
=
−
6
⁡
mm
,
𝑇
𝑒
=
11
.
4
⁡
eV
,
𝜌
𝑒
=
1
.
35
⁡
mm
,
𝜆
𝐷𝑒
=
0
.
88
⁡
mm
,
⁡
𝑐
𝑠
=
5
.
2
⁡
km
/
s
,
𝜈
𝑒𝑛
=
2
×
10
7
𝑠
−
1
.

Figure
4
: Angular frequency

ω
𝑅

and growth rate
γ

of the instability
normalized to the lower hybrid frequency
ω
𝐿𝐻

calculated from
the
dispersion relation

(1),
calculate
with the

plasma and magnetic field

parameters corresponding to the
position
𝑥
=
1
.
8
⁡
cm

of
Figure
2

(
𝜔
𝐿𝐻
≈
4
×
10
6
⁡
rd s
⁄
). The dashed line represents the
corresponding wavelength of the instability
𝜆
=
2
𝜋
𝑘
⁄
.

The solutions of the dispersion relation

at this position
,
normalized to the lower hyb rid frequency
𝜔
𝐿𝐻
=
(
𝑒𝐵
𝑚𝑀
⁄
)
1
/
2

are shown in
Figure
4

(
𝜔
𝐿𝐻
≈
4
×
10
6
⁡
rd s
⁄
⁡
⁡
at
𝑥
=
1
.
8
⁡
cm
).
We see on
Figure
4

that the
instability grows for low values of
𝑘
𝜌
𝑒

with a maximum
growth rate

𝛾
𝑚𝑎𝑥

(
𝛾
𝑚𝑎𝑥
≈
2
.
7
𝜔
𝐿𝐻
≈
11
.
×
10
6
⁡
s
−
1
),
corresponding to a wavenumber
𝑘
𝛾
𝑚𝑎𝑥
≈
400
⁡
m
−
1

and to a
wavelength
𝜆
𝛾
𝑚𝑎𝑥
=
2
𝜋
𝑘
𝛾
𝑚𝑎𝑥
⁄

of about 1.6 cm, consistent
with the first phase of
Figure
3
.

At

the maximum growth rate

𝛾
𝑚𝑎𝑥
, the wave angular frequency
𝜔
𝑅
,
𝛾
𝑚𝑎𝑥

and the phase
velocity

𝜔
𝑅
,
𝛾
𝑚𝑎𝑥
𝑘
𝛾
𝑚𝑎𝑥
⁄

are
⁡
respectively
on the order of

7
×
10
6
⁡
rd
⁡
s
−
1

and

17

km/s
(note the sharp variations of
𝜔
𝑅
⁡
⁡
for values of
𝑘

around
𝑘
𝛾
𝑚𝑎𝑥
in
⁡
Figure
⁡
4
)
. T
he phase
velocity estimated from
Figure
3

was 8 km/s.

The increase of
the wavelength with time in
Figure
3

is due to the deformation
of the plasma density and electric field during the
development of the instability.

As said above, the instabil it y triggered at

𝑥
=
1
.
8
⁡
cm

corresponds to a SimonHoh mode that disappears further
away from the anode because
|
𝐿
𝑛
|

increases
.

(
𝜔
∗
−
𝜔
D
)

goes
through zero and changes sign
at

𝑥
≲
1
.
7
⁡
cm
. Around this
point we observe, in the solutions of the dispersion relation
(1), a transition from a SimonHoh
mode

to a lowerhybrid
mode.
For small values of
(
𝜔
∗
−
𝜔
D
)
⁡
t he dispersion relati on
can be written
9
,
assuming

small Debye length and

ω
,
𝜔
D
⁡
≪
𝜔
0

, neglecting the

𝑘
2
𝜌
𝑒
2

term in the denominator of eq.
(1), and noting that
𝑐
𝑠
2
𝜌
𝑒
2
=
𝜔
𝐿𝐻
2
⁄

:

(
𝜔
0
−
𝑖
𝜈
𝑒𝑛
)
𝜔
0
⁡
=
𝜔
𝐿𝐻
2
𝜔
2

(
3
)

For small

𝜈
𝑒𝑛

with respect to
𝜔
0
,

this
equation has
solutions with positive growth rate
9
:

𝜔
=
𝜔
𝐿𝐻
(
1
+
𝑖
𝜈
𝑒𝑛
2
𝜔
0
)

(4)

This dispersion relation corresponds to a lower hybridmode
(
𝜔
𝑅
=
𝜔
𝐿𝐻
)

destabilized by collisions and
𝑬
×
𝑩

flow.

6

Numerical solutions of the dispersion relation (1) for
𝑥
=
1
.
7
⁡
cm

(not shown here) confirm the transition from
a
SimonHoh

or gradientdrift mode

to
a
lowerhybrid mode, with
maximum
growth rates significantly smaller tha n

in
Figure
4
.

In conclusion we can say that, in the conditions of
Figure
3
, the instability develops close to
the anode surface through

a

SimonHoh o r gradientdrift mode. The triggering of the
instability is however strongly dependent on the relative
values of the characteristic length of plasma density and
magnetic field gradients. For values of the magnetic fie ld
gradient in the anode region

smaller

than that used in the
conditions of
Figure
3
, the instability can be triggered by a
lowerhybrid mode destabilized by collisio ns and
𝑬
×
𝑩

flow. We checked this numerically (not described here) by
decreasing the magnetic field gradient

in the anode region
(keeping the same profile in the cathode region) and found
that even in that case the instability evolves into an ionization
w ave
similar to that
of
Figure
3
,
on a long enough time scale.

D.

R
otating spoke
regime

We discuss here the steady state spoke regime that takes
place after time
t=
10

s

in
Figure
3
.

The electric potential,
electron temperature and ionization rate

at time
t=
80

s

are
shown in
Figure
5
. The potential wave is correlated with the
spoke motion shown on
Figure
3
.
As found in Ref.
39
, the
spoke is associated with strong distortions of the electric
potenti al, electron temperature and ionization rate.

Figure
5

: 2D color contour

plots

of plasma
potential,
𝜙
, electron
temperature,
𝑇
𝑒
,
ionization rate,
𝑛
𝑒
𝜈
𝑖
,

and ion density,
𝑛
𝑖
, at time
t=
80

s
. in the conditions of
Figure
3
.
𝜙
,

𝑇
𝑒
,
and
𝑛
𝑒
𝜈
𝑖

are plotted on
a linear scale

(color bar on the left side), while
𝑛
𝑖

is plotted on a
twodecade logscale (c olor bar on the right side).

The potential
contours (black lines) are plotted every 20 V from 0 to 260 V. The
unit is 20 eV for the electron temperature,
1
.
4
×
10
22

m
-
3
s
-
1

for the
ionization rate
,
and

2
×
10
16

m
-
3

for the ion density. Three contours
of constant ionization rate (black lines) corresponding to 0.2, 0.4,
and
0
.
8
×
10
22

m
-
3
s
-
1

are superimposed on the ion density color
contour plot. T
he axial
profiles

of the charged particle densities and
electric field of
Figure
6

are plotted at the azimuthal positions y
1

and
y
2

indicated by the dashed lines on the potential plot.

Figure
6

: Axial
profiles of th e charged particle densities,
𝑛
𝑒
,
𝑛
𝑖
,

and
electric field
𝐸
, at two different azimuthal locations in the conditions
of
Figure
5

(time
t=
80

s

of
Figure
3
) corresponding to (a),
𝑦
1
=5.56
cm, and (b),
𝑦
2
=1.36 cm. These two azimuthal positions are
indicated by dashed lines on the electric potential plot of
Figure
5
.
Note that the scales of the plots in (a) and (b) are

different.

The word “spoke”, in the context of low temperature
𝑬
×
𝑩

plasmas, is generally used to refer to the region of
intense light emission rotating azimuthally, which could be
due to a local enhancement of the plasma density or electron
temperatu re (and ionization and light emission) or both.
In the
magnetron discharge simulations of
Figure
3

and
Figure
5
, it
appears that the increase of the plasma density is the result of
the local increase of the ionization rate and to the low electric
field region behind the ionization region. In the rest o f this
paper we use the word “spoke” in a more general sense, to
designate the
ensemble formed by the ionization region
(spoke front) plus the
region of large plasma density
generated behind

the

spoke front
, and that is clearly visible in
Figure
3

after time t=
10

s
, and in
Figure
5
.
The
spoke is
characterized by a large

plasma

density

(
about

2
×
10
16

m
-
3

in the conditions of
Figure
5
) and a quasiconstant potential
,

close to the anode potential

(electrons are trapped in this low
elec tric field region, hence the high plasma density).

Figure
6

shows the axial profiles of the electric field and
charged particle densities at two different azimuthal locations
noted
𝑦
1

and
𝑦
2

on the plasma potential plot of
Figure
5
. On e

can see in
Figure
6
a t he large electric field in the sheath
between the cathode and the spoke at the azimuthal location
𝑦
1

and the very sharp increase o f the plasma density at the
spoke boundary.

Outside the spoke in the azimuthal
direction, the electric field penetrates in the quasineutral
plasma

as can be seen in the potential plot of
Figure
5
. The
large electric field in the quasineutral plasma outside the
spoke can also be seen in
Figure
6
b where the axial profiles
of the ch arge d

particles densities and electric field are plotted
at the location
𝑦
2
.

The distribution of electron temperature shown in
Figure
5

presents an important maximum at the boundary between the
low electric field region in the spoke and the large field
outside the spoke

and is

large r

at

the spoke front
, i.e. on the
upper side

of the potential wave. The ionization rate is also

7

maximum along the spoke boundary and is slightly shifte d

downward with respect to the electron temperature.

The maximum of ionization rate is located on the upper
side of the maximum of plasma density (i
.e. in the

−
𝑬
×
𝑩

direction
) as can be seen on the right plot of
Figure
5

where
the ion density is represented as color

contours and the
ionization rate as contour lines. This shows that the spoke
motion is associated with an ionization wave moving in the
−
𝑬
×
𝑩

direction.
One can infer that the velocity of the spoke
motion is related to the position of the ionization rate

with
respect to the maximum plasma density. The mechanisms of
electron heating, which control in a large part the space
distribution of the ionization rate, are therefore a determining
factor in establishing the speed of the spoke rotation.

The conditions

of the 1D simulation of
Figure
2

and of the
2D simulation of
Figure
3
,
Figure
5
, and
Figure
6

are identical
(same current density of electrons emitted at the cathode,
applied voltage, gas pressure and electrode gap). It is
interesting to note that

despite

these identical conditions,
t he
maximum plasma density

is about ten times larger in the
spoke of the 2D simulation than in the 1D simulation. Outside
the spoke in the azimuthal direction,
the 1D and 2D
simulation results present similar values of the plasma
density, and an important p enetration of the axial electric field
in the quasineutral plasma (compare
Figure
2

and
Figure
6
b).
The

much larger plasma density and very small electric field
in the spoke

of the 2D simulation
are

due to

the large
ionization rate at the spoke front
, which

is responsible for the
plasma generation behind the front. Note that
despite

the
much larger maximum plasma density in the 2D result, the
calculated total current density and the electron multiplication
in the 2D simulation (respectively 4.5 A/m
2

and 150) are not
much different from (and even slightly smaller than) those of
the 1D

simulation (respectively 5.1 A/m
2

and 170, see above).

Figure
7

shows the 2D distributions of the charge density
[
𝑛
𝑖
−
𝑛
𝑒
]
(
𝑥
,
𝑦
)
, of the electron current density
𝑗
𝑒
(
𝑥
,
𝑦
)

and
current lines in the gap, and of the azimuthal profile
𝑗
𝑒
(
𝑑
𝑥
,
𝑦
)
⁡
of the electron current density on the anode surface
in the conditions of
Figure
3
,
at time
t=
80

s
. The plot of the
charge density shows that a double layer forms at the interface
between the low electric field region in the spoke and the
higher electric field around it. The ion sheath associated with
this double layer is clearly visible

on
Figure
7
, while the
negative charge density is more diffuse in the spoke.

The
electron current density is maximum in the
ion sheath region
of the
double layer. Th is corresponds to maximum values of
the
𝑬
×
𝑩

and diamagnetic current density. Note that the
electric field is perpendicular to the equipotential lines so the

𝑬
×
𝑩

current density follows the equipotential lines.

It appears on
Figure
7
a that the deformation of the plasma
potential associated with the spoke creates a kind of shortcircuit between cathode and anode and is responsible for the
transport of a large electron current from the cathode region
to the anode side. This can be seen on the electron current
density profile on the anode surface, plotted in
Figure
7
c

(n ote
the different scales of

the

𝑗
𝑒
(
𝑥
,
𝑦
)

and
𝑗
𝑒
(
𝑑
𝑥
,
𝑦
)

plots).

Figure
7
: 2D distribution s

of
(a)
the charge density
𝑛
𝑖
−
𝑛
𝑒
,
⁡
⁡
⁡
(
units:
10
15

m
-
3
),
and (b)
the

total electron
current

density

𝑗
𝑒

(
unit: 500
A/m
2
)
;
the colour contour s

show the intensity of the current density,
the white lines show the current lines
; (
c
) electron current density
profile on the anode surface.
Conditions of
Figure
3

at
time
t=
80

s
.

E.

Comparisons with probe measurements

Several experimental papers report probe measurements
of the plasma properties in magnetron discharges. In this
section we consider the re cent measurements of Held et al.
15

who used Lan gmuir probes to measure the time evolution of
plasma potential, electron density, electron temperature and
electron energy probability functions (EEPF) in dc and pulsed
magnetron discharges in the spoke region. The conditions of
the measurements

in the dc
case

(argon, 0.56 Pa
, 2
55

V dc
voltage
) are close to those of the
particle
simulations
described
above. The magnetic field distribution used in the
simulation (
Figure
1
b) is only

a

rough approximation of the
magnetic field in the experiment, based on Figure 1 of Held
et al.
15
.

The Langmuir probe in the experiment was place d

6
mm above the cathode surface
. The azimuthal period of the
simulations (8 cm) i s close to the length of the azimuthal line
at the probe location in the experiment.

The probe measurements of Held et al. are shown in
Figure
8
. The measurements exhibit large amplitude oscillations of
the electron density, electron temperature and plasma
potential, with a very sharp increase of the plasma potential
and electron temp erature when the spoke reaches the probe.
The modulation of the electron density is as large as 70 % and
the plasma potential and electron temperature jumps at the
spoke front are respectively on the order of 20 eV and 60 V.

The results of the 2D PICMCC s imulations are displayed
in
Figure
9
. An important difference between experiments
and simulations that can be seen by comparing
Figure
8

a nd
Figure
9

is that the plasma density is about ten times smaller
in the simulations. This is because the current density of
electrons emitted at the cathode in the si mulations has been
imposed (0.03 A/m
2
) in order to limit the maximum plasma
density to numerically tractable values. Despite this
difference, the time evolution of the plasma parameters in the
simulation are qualitatively similar to those of the
experiment
.

8

Figure
8

: Probe measurement s

by Held et al.
15

of the time evolution
of electron density,
𝑛
𝑒
, electron temperature,
𝑇
𝑒
, and plasma
potential
𝜙

in a dc magnetron discharge in argon at 0.5
6

Pa and with
a dc voltage of 2
55

V.
The probe is located 6 mm above the cathode
surface on the spoke path.
The gray areas indicate positions where
the validity of the electron temperature and plasma potential are in
question (see Held et al. for a complete description of these
Langmuir prob e

measurements).

The simulation results of
Figure
9

exhibit large amplitude
oscillations of the plasma properties with abrupt increase of
the electron temperature,

plasma potential, and electron
density when the spoke reaches the location of the numerical
probe. The potential jump is larger in the simulation and,
although the minimum and maximum electron temperature
are similar in the experiment and simulation, the
detailed time
variations of the electron temperature present some
discrepancy with those of the measurements.

For example, in the experiment, the electron temperature
increases very sharply at the spoke front and decreases more
slowly after the front. This

is the opposite in the simulation
results where the decay of electron temperature after the
spoke (marked “B” in the electron temperature plot of
Figure
9
) is much faster than the temperature rise just at the spoke
front (“A” in the temperature plot). One possible reason for
the sharper electron temperature rise at the spoke front is the
larger

plasma density (and smaller Debye length
, i.e. thinner
double layer
) in the experiment. The temperature after the
spoke front quickly drops to about 3 eV in the simulation
(“C” in the electron temperature plot of
Figure
9
), and
increases again at the trailing edge of the spoke (“D” in the
electron temperature plot). This is consistent with the 2D
plots of the electron temperature in
Figure
5
).

Figure
9
: Time variations of

electron density,
𝑛
𝑒
, electron
temperature,
𝑇
𝑒
, and plasma potential
𝜙

from the 2D PICMCC
simulation, 6 mm above the
cathode surface

(
𝑥
=
6
⁡
mm
)
, in the
conditions of
Figure
3

(argon, 0.6 Pa, dc voltage 260 V).

The slower decay of the electron temperature in the
experiments could be

due to Coulomb collisions, which are
not taken into account in the simulations.

It is interesting to look at the electron energy probability
function (EEPF) in the region of the spoke front and to
compare the EEPF deduced from the simulations, with the
EE
PF measured by Held et al.

15

(Figure 9 of Ref.
15
).
Figure
10

shows the EEPF deduced from the PICMCC simulations
at three different locations around the spoke front. The tail of
the EEPF is strongly modulated in this region. The population
of

high energy electrons increases abruptly in the double layer
region and decays in the high plasma density and low field
region of the spoke. This is consistent with the heating of
electrons flowing along the double layer in the spoke front.
The residence
time of electrons in the highdensity plasma
behind the spoke front is long because these electrons are
trapped in this low electric field region. This explains the fast
decay of the electron temperature in the simulations and the
large population of low e nergy electrons in that region
(Coulomb collisions could reduce this effect).

Finally, we note that the period of variations of the plasma
parameters is shorter in the experiments of Held et al. (about
32

s, see
Figure
8
) than in the simulation (60

s,
Figure
9
).
This indicates that the spoke velocity in the experiments (8
cm in 32


s, i.e. 2.5 km/s) is about twice the velocity in the
simulation (about 1.3 km/s according to
Figure
9
).

9

Figure
10
:
(a)
Normalized electron energy probability function
(EEPF) at
four

different

azimuthal

locations around the spoke front
,
8
.8 mm above the cathode surface
(
𝑥
=
8
.
8
⁡
mm
)

obtained from the
PICMCC simulation at a given time in the conditions of
Figure
3
.
The EEPF are integrated over
squares of 80

m side (indicated on
(c))

in space
, and over 10 ns

in time
;

(b) azimuthal
profile

of the
electron temperature at
𝑥
=
8
.
8
⁡
mm

in the conditions of (a)

(showing only 4 cm of the 8 cm azimuthal length)
; (c) 2D
distribution of electron temperature, and (d) of ion density. The
azimuthal locations corresponding to the EEPF of (a) are indicated

by
colored squares

and numbers

on (c), and by lines on

(b), with the
same colors

and numbers

as the EEPF plots on (a).
The three white
lines on (d) correspond to contour of constant potential 240, 2
5
0,
and 2
55

V.

F.

Spoke velocity

The spoke velocity
in
dc
m agnetron discharge
experiments in argon is

in the
−
𝑬
×
𝑩

direction,
typically
between 0 and
10

km/s
,

and is found to increase with
increasing current
or power
(see, e.g.,
Figure 6 of Hecimovic
and von Keudell
32
)
. In the high power, pulsed regime the
spokes rotate in the
+
𝑬
×
𝑩

direction
with velocities up to
15 km/s
32
.

We did not perform a systematic parameter study of the
spoke velocity
,

but we noticed that the profile of the magnetic
field has an in fluence on the spoke velocity in the simulation.
For example in the simulations of Boeuf and Takahashi
38
,
39

where the magnetic field was supposed to have a Gaussian
shape (with a maximum of 100 mT at the cathode surface
)
,
the spoke ve locity

(in argon)

was on the order of 10 km/s.
In
the present work we also performed simulations at 0.6 Pa with
a magnetic field profile as in Refs.

38
,
39

and a maximum
magnetic field of 120 mT (same co nditions as
Figure
3

except
for the magnetic field profile). The spoke velocity was about
4 km/s in that case (instead of 1.3 km/s for the case of
Figure
3
). Finally, the spoke velocity at a pressure of 2 Pa in the same
conditions as
Figure
3

(this case is d escribed in section III.
H

below)

was 1.5 km/s, slightly larger

than
the

0.6 Pa
case of
Figure
3
.

We can conclude that the spoke velocity depends in a
complex way on ma ny parameters (pressure, magnetic field
intensity and profile
, power,

etc…). This is because the spoke
motion is

due to

an ionization wave and its velocity depends
on the relative distributions of ionization rate and plasma
density. In any case, we can exp ect that the order of
magnitude of the spoke velocity is also related to the ion mass
(through the ion velocity) since the motion of ions accelerated
outside the spoke front is also a parameter controlling the
spoke velocity.

More work is necessary to bett er understand
the parameters controlling the spoke velocity.

G.

Electron

heating

and role of
the
𝑩
×
𝛁
𝑩

drift

As can be seen in
Figure
5
, the electron temperature and
ionization rate are maximum in the region of the double layer
around the spoke. Electron heating is therefor e

larger
along
the double layer
. We showed in Refs.
38
,
39

that the electron
drift due to the magnetic field gradient

(“
𝛁
𝐵

drift
”)

can play
an important role in electron heating in magnetron discharges.

The
𝛁
𝐵

electron drift velocity is given by
51

𝒗
𝛁
𝐵
=
−
1
2
⁄

𝜌
𝑒
𝑣
⊥
𝑩
×
𝛁
𝐵
/
𝐵
2
=
−
𝘀
⊥
𝑩
×
𝛁
𝐵
/
𝐵
3

where
𝑣
⊥
⁡
is the
electron velocity and
𝘀
⊥

the electron energy in eV,
perpendicular to the magnetic field.
In the conditions of the
simulations considered here (see
Figure
1
)
𝒗
𝛁
𝐵

is directed in
the
azimuthal

direction
,

in the positive
𝑦

direction
.

The electron heating or cooling

per electron due to the
𝛁
𝐵

drift

is given by
:

𝜃
𝒗
𝛁
𝐵
=
−
𝒗
𝛁
𝐵
.
𝑬
=
𝘀
⊥
𝐸
𝑦
(
𝐵
𝐿
𝐵
)
⁄

(5)

where
𝐿
𝐵
=
𝐵
𝜕
𝑥
𝐵
⁄
.

Considering the direction of the
azimuthal electric field

𝐸
𝑦

around the spoke that can be
deduced from the equipotential contours of
Figure
5
, we can
conclude that the
𝛁
𝐵

drift

contributes to electron heating on
the upper side of the potential wave

(
𝐸
𝑦

and
𝐿
𝐵

are negative
in eq. (5))

and to electron
cooling
on the lower side

(
𝐸
𝑦

positive,
𝐿
𝐵

negative)
.
Note that in the absence of collisions,
electron heating and cooling would cancel

eachother

because
of the adiabaticity of the electron trajectories. The fact that
there is a net heating in our conditions is due to collisio ns
:
electrons can lose energy through inelastic collisions before

𝛁
𝐵

drift

cooling becomes important.

Nonadiabatic trajectories leading to electron heating is
possible when
sharp spatial variations of the electric field
such t ha t

|
𝛁
𝐸
|
𝐵
⁄
>
𝜔
𝑐𝑒

are present
52
. This situation occurs
in collisionless shocks in space plasmas
53
-
55
.
We checked that
t he

condition
|
𝛁
𝐸
|
𝐵
⁄
>
𝜔
𝑐𝑒

is not satisfied in our
simulations where the plasma density is
limited because of
numerical constraints.

10

Figure
11

: Comparisons of the 2D distributions of the heating per
electron due to
, (a),
(b)

𝛁
𝑩

drift
,
𝜃
𝒗
𝛁
𝐵

, eq. (5)
, with
positive
contribution in (a) and negative contribution in (b)
,

and

(c),

to
collisional heating,
𝜃
𝑬

,
eq. (6)
; (d),
azimuthal profiles

at
⁡
𝑥
=
1

cm
(dashed

line of the 2D contour plots

(a), (b), and (c)
) of the
𝛁
𝑩

drift

and collisional el e
ctron heating.

Note the different scales,
10
9

in (a),
(b),
and
2
×
10
8
⁡
eV/s

in

(c).

This condition could however be reached under the

much

higher plasma densities and smaller Debye lengths

(i.e.
thinner double layer)

of
high power
magnetrons discharges,
leading to larger values of
|
𝛁
𝐸
|
𝐵
⁄

in the double layer.

This
question is left for further studies.

Since the electric field is large in the
double layer around
the spoke, another possible contribution to electron heating is
the classical, collisional heating. Neglecting the diamagnetic
term, the contribution of collisional heating to the total
electron heating can be written as:

𝜃
𝑬
=
−
𝒗
∥
𝐄
.
𝑬
=
𝝂
𝛀
𝒄𝒆
𝑬
𝟐
𝑩

(6)

where
⁡
𝒗
∥
𝐄

is the electron drift velocity parallel to the electric
field.

The contribution of the
𝒗
𝛁
𝐵

and
𝒗
∥
𝐄

(collisional) heating to
the total electron heating in the conditions of
Figure
5

are
shown on
Figure
11
.

Figure
11
a

and
Figure
11
b
confirm

that
the
𝛁
𝐵

drift

contributes to electron heating on one side of the
potenti al wave and to electron cooling on the other side.
Collisional electron heating is present all along the double
layer
,

but its magnitude is significantly smaller than the
heating due to
𝛁
𝐵

drift
. This is also true in the larger pressure
condition (2 Pa)

of section III.H. This confirms that the
𝛁
𝐵

drift

plays an essential role in the spoke formation and
sustainment in low power dc magnetron discharges.

Electron
heating on the upper side of the potential wave due to
𝛁
𝐵

drift

and electron cooling
on the

lower side tend to limit the
ionization region to the upper side, i.e. above the region of
large plasma density. This clearly explains the rotation of the
spoke in the
−
𝑬
×
𝑩

direction and the coherence of the
spoke rotation.

To confirm this conclusion,
we performed simulations with
uniform magnetic field s
, in order to remove the
𝛁
𝐵

drift

contribution to electron heating. In
a
previous work
36
,
37

we
also

performed
2D
PICMCC simulation s

of magnetron
discharges under uniform magnetic field, but this was done in
cylindrical geometry, i.e. with axial magnetic field and
cylindrical cathode and anode.
In the cylindrical
configuration as well as in the planar
magnetron
configuration (the present work)

with uniform magnetic field,

the particle simulation predicts the rotation of plasma nonuniformities in the
+
𝑬
×
𝑩

direction. These nonuniformities
are

relatively coherent or much more chaotic at
steady state, depending on the conditions, but they
are

never
as

regular and
coherent as the spoke

in a magnetic field
gradient

described in the present work and in Refs.
38
,
39
.
Moreover, in the simulations in a uniform magnetic field,
we
never found conditions where most of the ionization took
place in the spoke (ionization in the cathode sheath was
always dominant). For example, one can see in Fig ure 9 of
Ref.
37

or in Figure 2 of Ref.
36

that ionization
in the spoke is
a very small fraction of the total ionization.

H.

Influence of pressure

We have seen above that electron heating is larger
in the ion
sheath
region

of the double layer

around the spoke
,

on the
upper side of the potential wave
,

due to
a combination of
𝛁
𝐵

electron drift

and
(to a lesser extent)
collisional crossfield transport
.

Electrons heated in this region continue
flowing

in the
𝑬
×
𝑩

direction along the
ion sheath around
the spok e. They lose energy along this
line due to inelastic
collisions

and of
𝛁
𝐵

drift

cooling
. The
mean free path for
ionization in the
𝑬
×
𝑩

direction can be estimated
as
𝜆
𝑖
=
𝑣
𝑑
𝜈
𝑖
⁄

w here
𝑣
𝑑
=
𝐸
𝐵
⁄
⁡
is the
𝑬
×
𝑩

drift velocity and
𝜈
𝑖

the
ionization

frequency.
The simulation in the conditions of
Figure
3

gives a maximum of
𝜈
𝑖

on the order of
10
7

s
-
1
in the
electron heating region and a drift velocity

𝑣
𝑑

between
5
×
10
5

and
10
6
⁡
m/s. This gives a value of the ionization
mean free path between
5

and 1
0

cm

which is on the order of
the azimuthal length in the conditions of
Figure
3
. It seems
reasonable to expect an increase of the number of spokes as
the ionization mean free path decreases, i.e. as the pressure
increases, as discussed by Panjan et al.
16
.

This is confirmed by the particle model.
PICMCC
simulations performed at a pressure of 2 Pa

and applied
voltage
𝑉
=280 V

(other conditions
being

the same as in
Figure
3
) predict the formation of two spokes in the azimuthal
direction, as can be seen in

Figure
12
. At 2 Pa, we can expect
an ioniza tion mean free path smaller than at 0.6 Pa by a factor
of about 3. This is consistent with the length of the ionization
region in
Figure
12

(2 Pa) compared with
Figure
3

(0.6 Pa).

The consequence is that the number of spokes increases with
gas pressure. This has been found in experiments as discussed
in detail in the article by Panjan et al.
16
.

The electron current density at the anode in these conditions
(2 Pa, 280 V) is
𝑗
𝑒
𝑑
=
7.25 A/m
2
, slightly larger than at 0.6
Pa, 5.1 A/m
2
. This corresponds
to an electron multiplication
𝑜𝑓
⁡
242

and a net secondary electron emission coefficient
of
⁡
4
×
10
−
3
.

11

Figure
12
: 2D color contour plots of plasma potential,
𝜙
, electron
temperature,
𝑇
𝑒
, ionization rate,
𝑛
𝑒
𝜈
𝑖
, and ion density,
𝑛
𝑖

in argon at
a pressure
𝑝
=
2 Pa and an applied voltage
𝑉
=280 V
,

all the other
conditions being the same as in
Figure
3
.
𝜙
,
𝑇
𝑒
, and
𝑛
𝑒
𝜈
𝑖

are plotted
on a linear scale

(color bar on the left side) while
𝑛
𝑖

is plotted on a
twodecade logscale (color bar on the right side).

The potential
contours are
plotted every 20 V from 0 to 280 V. The unit is 20 eV
for the electron temperature,
3
×
10
22

m
-
3
s
-
1

for the ionization rate
,
and

1
.
5
×
10
16

m
-
3

for the ion density.
The spokes move in the

−
𝑬
×
𝑩

direction, as in the 0.6 Pa case.

If the pressure is further increased, the number of spokes
continues to increase. The spokes become closer to each
other, and the modulation of the plasm

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
