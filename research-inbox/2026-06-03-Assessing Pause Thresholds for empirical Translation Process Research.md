---
source: "https://arxiv.org/abs/2604.01410v2"
title: "Assessing Pause Thresholds for empirical Translation Process Research"
author: "Devi Sri Bandaru, Michael Carl, Xinyue Ren"
year: "2026"
publication: "arXiv preprint / cs.CL"
download: "https://arxiv.org/pdf/2604.01410v2"
pdf: "https://arxiv.org/pdf/2604.01410v2"
captured_at: "2026-06-03T19:30:58Z"
updated_at: "2026-06-03T19:30:58Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# Assessing Pause Thresholds for empirical Translation Process Research

- 著者: Devi Sri Bandaru, Michael Carl, Xinyue Ren
- 年: 2026
- 掲載情報: arXiv preprint / cs.CL
- 情報源: [arxiv](https://arxiv.org/abs/2604.01410v2)
- ダウンロード: https://arxiv.org/pdf/2604.01410v2
- PDF: https://arxiv.org/pdf/2604.01410v2

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

Text production (and translations) proceeds in the form of stretches of typing, interrupted by keystroke pauses. It is often assumed that fast typing reflects unchallenged/automated translation production while long(er) typing pauses are indicative of translation problems, hurdles or difficulties. Building on a long discussion concerning the determination of pause thresholds that separate automated from presumably reflective translation processes (O'Brien, 2006; Alves and Vale, 2009; Timarova et al., 2011; Dragsted and Carl, 2013; Lacruz et al., 2014; Kumpulainen, 2015; Heilmann and Neumann 2016), this paper compares five approaches for computing these pause thresholds, and suggest and evaluate a novel method for computing Production Unit Breaks.

## PDF Text

Proceedings of
Translation in Transition 8, 2026

Assessing

Pause

Thresholds

for

empirical

Translation

Process

Research

Devi

Sri

Bandaru,

Michael

Carl,

Xinyue

REN
CRITT, Kent State University

Translation production proceeds in the form of stretches of typing, interrupted by
keystroke pauses. It is often assumed that fast typing reflects unchallenged/automated
processes while long(er) typing pauses are indicative of translation problems, hurdles

or difficulties. While there has been a long discussion

concerning

the

determination

of

pause

thresholds

that

separate

automated from

presumably

reflective translation

processes (O’Brien

2006;

Alves

and

Vale

2009; Timarová

et

al.

2011;

Dragsted

and

Carl

2013;

Lacruz

et

al.

2014;

Kumpulainen

2015; Heilmann and Neumann 2016), this
paper compares three recent approaches for computing these pause thresholds.

Methods and Metrics

We draw

on

a set of

478

fromscratch

translation sessions from

the CRITT

TPRDB that
cover 10 different language pairs to assess
two

pause thresholds that have recently
been proposed and a novel method. In contrast to previous approaches (e.g., Carl and
Kay 2011; CoutoVale 2017) these recent approaches take into account translatorspecific typing speed based on the distinction between

median withinword (WW) interkeystroke intervals (IKIs
1
) and median betweenword (BW) IKIs (Muñoz & Apfelthaler
2022, Miljanović et al., 2025, Carl et al 2025). WWIKIs are shorter

than BWIKIs
,
presumably

because
automated typing processes
are dominant
when a typist is
completing a word/thought that they have in mind

(Crotti et al 2026)
.

BWIKIs are longer
than

WWIKIs

as

they

“are

more

likely

to be linked to lexical retrieval” (Miljanović et al.
2025).

To

compute

WW

and

BW,

we

adopt

the

method

suggested

by

Miljanović

et

al.

(2025):

•

WW:

median

delay

between

two

successive

alphanumeric

keystrokes

•

BW:

median

delay

between

an

alphanumeric

and

a

nonalphanumeric

keystroke

Two

cases

can

be

distinguished
for

computing

BWs
:

WS

is
a

keystroke

bigram

that
occurs

at the

end

of

a

word

(in alphabetical languages)

when an

alphanumeric

keystroke (W,

i.e.,
last character of a word) is followed by a nonalphanumeric keystroke
(S, e.g., a blank space).

SW

occurs

at

the

beginning

of

a new

word,

when

a

nonalphanumeric

keystroke

(S) is followed by an alphanumeric keystroke (W
)
, first
character of the next word. Surprisingly,

the

durations

of

these

two

types

of

BWs

are

quite

different

(see

Figure

2): median WSIKIs are much shorter
,

and their distribution
more similar to WW
,

than median

SWIKIs, suggesting

that the
whitespace

follow ing

a
word

is more tightly linked to the preceding word than to the following word.

Accordingly,
we base the BW computation on the longer SWIKI. We compare three approaches:

1

Translation sessions were collected with TranslogII (Carl 2012). TranslogII has a builtin keylogg er (key
&
text
logging for Chinese
/Japanese
)
where an

IKI is the lag of time between two successive keydown events
. Note
that
Crotti et al (2026) define an IKI as “
the lapse between the keyup of one key and the keydown

of the next
”
.

Proceedings of
Translation in Transition 8, 2026

1.

Muñoz

&

Apfelthaler

(2022)

distinguish

between:

❖

RSP
:

Respite:

Unintentional

short

pause,

defined

as

2

×

median

WWIKIs.

❖

TSP:

TaskSegment

Pause:

Intentional

breaks,

defined

as

3

×

median

BWIKIs.

2.

Miljanović

et

al.

(2025)

suggest

the

following

two

measures:

❖

BW
:

Processing

effort:

i.e.,

median

BWIKIs.

(
same as
SW
)

❖

MUD
:

MicroUnit

Delimiter:

defined

as

median

WWIKI

+

2

×

standard

deviation

3.

This

paper:

❖

KU
I
:

Keystroke

Interruption
,

defined

as

2 ×

median

WWIKIs

(same

as

RSP
).

❖

PUB
:

Production

Unit

Break,

partially

defined

based

on

product

data,

as

follows:

We suggest and evaluate a novel method for computing Production Unit Breaks
(
PUBs
)
1
. The rationale for
PUB
s is grounded in the observation that text production
proceeds in chunks of a few words (say 3 words) rather than wordbyword. That is, we
are more likely to see

an extended lexical retrieval (or reflection)

pause every third word
rather

than

every

word.

Accordingly,

we

take

the

IKI

duration of

the quantile

for

this

ratio
(text
length/#chunks) as a threshold that separates automated from reflective
processes.

For

instance,

assuming

that

a

word

(in

an
English

text
)

has

on

average

6

characters

and a chunk length of 3 words, we take it that the upper 1/(6*3=18) IKI
quantile
—
that

is,
~
5.5% of the
longest
IKIs
—
is related to intentional/reflective/lexical
retrieval pausing, while
the remaining
94.5% of the shorter IKIs are related to
automated processes.

Since the average length of words changes from text to text, also
this threshold ch anges slightly.

Table

1 shows

the

Pearson
(upper triangle)
and
Spearman
correlation s

between

the

four

measures

(the measures
3GL

and
3GU

are discussed below)
.

Note

that

there

is

a

perfect

correlation

between
SW

and
TSP
as the latter is just a multiple of the former.

Table

1

(left)
:

Correlation

matrix

of

pause

thresholds

and

GMM

clusters;

Upper triangle: Pearson correlation in italics;
lower triangle: Spearman correlation.
All Pearson correlation for
3G
L

(not shown)
are insignificant.
All
shown Spearman
correlations are highly significant (
p < 0.001
)
.

Figure

1

(right)
:

distribution

of

quantiles

for the four pause thresholds
.

Figure 2 (left) plots the distribution of the two BW measures and
,

on the right
,

the 4
pause

measures

for

the

478

translation

sessions.

Clearly,

the

distribution

of

SWIKI

is

1

This

version

of

PUB

is

different

from

the

one

described
in

(
Carl

et

al

2025
)
.

I
t is the basis for
the segmentation of

Production and Translation Units in
the TPRDB 3.0
, see
https://crittkent.github.io/TPRDBdocumentation/
.

RSP

TSP

MUD

PUB

3G
U

RSP

1

0.85

0.44

0.63

0.45

TSP

0.75

1

0.44

0.72

0.48

MUD

0.55

0.49

1

0.51

0.36

PUB

0.64

0.7

0.61

1

0.58

3GL

0.31

0.28

0.23

0.33

1

3GU

0.59

0.57

0.51

0.68

0.65

Proceedings of
Translation in Transition 8, 2026

much

flatter

than

the

WSIKI
.

TSP

and

PUB

are

strongly

correlated,

while

MUD

is much
more stretched
and

less strongly correlated with the other measures.

Validating

Pausing
Methods

While this analysis gives us a picture of how these
methods

distribute IKI
s

into two

categories (
pre sumably automated vs. reflective processes), it is unclear whether these
thresholds
/measures

actually do their job, and which of them in the best way.
To

address this question, we investigated the distributions of IKIs per translation session
with Gaussian Mixture Models (GMMs). GMMs identify clusters in the data without
assuming theoretically imposed cutoffs (as our four methods do). If a (e.g. 2
-
componen t) GMM consistently separates IKIs (or fixation durations
,
see
below
) into two
distinct distributions,
this suggests that there may be two underlying cognitive
processes

generating

these

distinct

timing

patterns,

potentially

corresponding

to

automated

versus reflective processing.

We applied the

Akaike Information Criterion (AIC) to determine the optimal number of
pause
categories

based

on

the

empirical

distribution

of

the

data.

The

AIC

analysis

consistently favored a threecomponent model for IKIs
—
suggesting that the IKI
distributions are actually composed of three independent processes
—
which
are
separated by two pause thresholds,
provid ing

strong statistical support for the
theoretical framework proposed by Muñoz and

Apfelthaler (2022)
.

They
posit three
components:

1.

Automated

Translation

Routines:

fastresponse

component,
short IKIs,
corresponding to motor execution and fluent text production.

2.

Unintentional Halts (Respites):
micropauses

where

cognitive

resources

are

momentarily

drawn

away from typing
;

yet
no
deep problemsolving

involved
.

Figure 2:
Left: d istribution of withinword (WW) and two types of betweenword BWIKIs
.

Note

that

the

distribution

of

WW

is

very

similar

to

WS,

suggesting that

the

separator following the

word

is still

part

of

the word

just typed
,

while SWIKIs are

indicative of

different
mental
process ing
.

Right: four different

pause

thresholds:

Note that
RSPs and TSP
are from

the same distributions as
WWs and SWs (left) multiplied by a factor of 2 and 3, respectively.
Horizontal axes in b oth Figures
are on a log scale.

Proceedings of
Translation in Transition 8, 2026

3.

Intentional

Stops

(Reflective

Pauses):

a
slowresponse component, probably
corresponding to higherorder cognitive processing and
decisionmaking.

The tripartite structure suggests two pause thresholds that
temporarily
separate the
three processes
.

W
e

label

3GL

the

lower

threshold

and

3GU

for

the

upper

threshold

in
Table 1
. In our ‘theorydriven’

analysis we have so far one candidate (
RSPs
) for the
lower threshold but three candidates
(
TSP
,
MUD
,
PUB
)

for the upper threshold
which

is
supposed to separate unintentional halts from intentional stops.

As shown in Table 1,
there is virtually no correlation between the lower threshold (
3GL
) and the ‘theoretical’
metrics
(not even
RSP
)
but a moderate to strong positive correlation for the three upper
threshold and the
3GU
.

This

suggests

all

three

analytic

thresholds

(
TSP
,

MUD
,

PUB
)

are

not

merely

arbitrary
but align with a datadriven separation of processing modes; the
PUB
method
,
albeit,
in
the best way.

Keystroke
Pausing
and
Fixation Duration

We validate whether the two kinds of IKI thresholds identified in our keystroke analysis
also correspond to different kinds of cognitive processing load as measured through
fixation duration.

It is generally assumed that longer fixations are indicative of
higher
cognitive load, which leads us to hypothesize that translators with longer average
fixation durations might also show longer keystroke pauses.

We distinguished between visual attention allocated to the

source text (ST, input
processing), and the

target text (TT, output monitoring).

Furthermore, we separated the
data into

first fixations (FF) duration, indicative of processing effort for initial l exical
access, and

all fixations (AF), which captures the total dwelling time including rereading and regressions.

We applied the GMM method to logtransformed fixation durations (LogDur), separately
for FF and AF. In striking contrast to the tripartite structure observed in IKIs, the AIC
analysis for ST and TT fixation durations consistently favored a

twocomponent mo del

(2G)
.

This suggests that visual attention in translation operates more likely under a
binary cognitive architecture, probably distinguishing between

automated
scanning

(short fixations for navigation) and

reflective processing

(long fixations
associate d with active decisionmaking).

We then determined the point where the two fixation duration probability curves cross
—

that is, the duration at which a fixation is equally likely to belong to either
the
automated or reflexive
component. This crossover point was calculated numerically
using Brent's method, a standard rootfinding algorithm. Fixations longer than this
threshold are assumed to reflect more likely deliberate, conscious/reflective processing
rather than automat ed

reading behavior.

Proceedings of
Translation in Transition 8, 2026

Metric

2G_ST_FF

2G_ST_AF

2G_TT_FF

PUB

RSP

TSP

MUD

3GL

3GU

2G_ST_FF

1

0.89

0.75

0.35

0.36

0.31

0.18

0.14

0.28

2G_ST_AF

0.89

1

0.81

0.35

0.4

0.37

0.2

0.13

0.29

2G_TT_FF

0.75

0.81

1

0.28

0.39

0.37

0.23

0.12

0.28

2G_TT_AF

0.81

0.89

0.9

0.31

0.43

0.4

0.21

0.12

0.28

Table 2
: Spearman correlations for fixation and pause thresholds.
Rows: f ixation 2
-
component GMM
threshold s for ST/TT and FF/AF
.
The
correlations

with

3GL

(GMM lower keystroke threshold)

are very
weak but significant p < 0.05. All other correlations are highly significant p < 0.01.

We
compute
this
2G

threshold
for
the
ST and TT
and for FF and AF and correlate the
threshold values

with keystroke pause metrics
—
PUB
,
3GL
,
3GU
,
TSP
,
RSP
, and
MUD

—

over

all
478

translation sessions. As shown in

Table 2,
the four
2G

thresholds
strongly correlate with each other, indicating that these measures are not independent.
The
analysis
also
reveals a consistent but moderate correlation between most
production pauses and fixation durations. Specifically, the
PUB
,
RSP
,
TSP

and
3GU

aligns more closely with AF metrics (for both ST and TT) than with FF, suggesting that
intentional pauses are more frequently associated with late gaze measures (rereading,
regressions, verificatio n) rather than initial lexical access.

Table 2 provides the Spearman correlation, but
Pearson correlation

consistently
provided lower values (not shown). The divergence between these two correlations
indicates that the relationship between gaze durations and keystroke pauses is
primarily monotonic rather than strictly linear; suggesting that visual (cognitive
) load
increases as keystroke pauses increase. N
otably, the
RSP

and
TSP

thresholds exhibit
the strongest monotonic (Spearman) correlations with TT fixations
,

whereas the lower
keystroke thre shold

3GL
,

show s

very
weak

correlation s

across all gaze measures. This
weak

correlation
in Tables 1 and 2
suggests that
3GL

likely captures motor execution
noise rather than cognitively driven pauses.
However,

the significant
correlation

across
the higher thresholds
suggests

that the thresholding approach captures a robust and
reliable signal of shifts between automatic and reflective processing modes.

Data Availability:

The data and the python scripts for this study can be downloaded from:

https://github.com/CrittKent/KeyGazeCorrelation

References

Alves, F., D. Vale (2009).
Probing the unit of translation in time: aspects of the design and development of a web
application

for

storing,

annotating,

and

querying

translation

process

data.

Across

Lang.

Cult.,

10

(2)

(2009),

pp.
251
-
273

Carl,

M.,

M.

Kay

(2011).

Gazing

and

typing

activities

during

translation

:

a

comparative

study

of

translation

units

of
professional and student translators. Meta, 56 (4) (2011), pp. 952
-
975

Carl,

M;

Takanori

Mizowaki,

Aishvarya

Raj,

Masaru

Yamada,

Devi

Sri

Bandaru,

Xinyue

REN.

(2025).

The

behavioural
translation

style

space:

Towards simulating

the temporal dynamics

of affect, behaviour, and

cognition in human
translation production. SKASE Journal of Translation and Interpretation, 2025; 18(2): 212
–
39 doi:
Proceedings of
Translation in Transition 8, 2026

10.33542/JTI2025
-
S
-
11 (also on
https://arxiv.org/abs/2507.12208)

CoutoVale,

D.

(2017)

What

does

a

translator

do

when

not

writing?

S.

Hansen

Schirra,

O.

Czulo,

S.

Hofmann

(Eds.),
Empirical

Modelling

of

Translation

and

Interpreting,

ume

10,

Language

Science

Press,

Berlin

(2017).

page

209
–
237Dragsted, B., M. Carl (2013) Towards a classification of translator profiles based on eyetracking and
keylogging data. J. Writ. Res., 5 (1) (2013), pp. 133
-
158,
10.17239/jowr
-
2013.05.01.6

Crotti, Roberto
,

Giovanni Denaro, Zhiqiang Du, Ricardo Muñoz Martín

(2026)
Hylog: A Hybrid Approach to Logging
Text Production in Nonalphabetic Scripts
.
https://arxiv.org/abs/2601.17753

Heilmann, A., S. Neumann (2016). Dynamic pause assessment of keystroke logged data for the detection of
complexity

in

translation

and

monolingual

text

production.

Conference:

Workshop

on

Computational

Linguistics
for Linguistic Complexity (CL4LC), Coling (2016)

Kumpulainen,

M.

(2015)

On

the

operationalisation

of

‘pauses’

in

translation

process

research.

The

International
Journal for Translation and Interpreting Research, 7 (2015), pp. 47
-
58. ti.106201.2015.a04

Lacruz,

I.,

Denkowski,

M.,

&

Lavie,

A.

(2014).

Cognitive

demand

and

cognitive

effort

in

postediting.

In

Proceedings

of
the 11th Conference of the

Association for Machine Translation in the

Americas
(pp. 73
-
84).

Miljanović,

Z,

F

Alves
,

C

Brost,

S

Neumann

(2025)

Directionality

in

translation:

Throwing

new

light

on

an

old

question
.

Muñoz,

R.,

M.

Apfelthaler

(2022)

A

Task

Segment

Framework

to

study

keylogged

translation

processes.

Translation
and Interpreting, 14 (2022)

O'Brien, S.

(2006). Pauses as indicators of cognitive effort in postediting machine translation output.

Across Lang.
Cult.,

7

(1)

(2006),

pp.

1
-
21,

10.1556/Acr.7.2006.1.1

https://akjournals.com/view/journals/084/7/1/articlep1.xml

SKASE

Journal

of

Translation

and

Interpretation.

Timarová, Š., Dragsted, B., & Hansen, I. G. (2011).
Time lag in translation and interpreting: A methodological
exploration. In
Methods and strategies of process research
(pp. 121
-
146). John Benjamins Publishing
Company

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
