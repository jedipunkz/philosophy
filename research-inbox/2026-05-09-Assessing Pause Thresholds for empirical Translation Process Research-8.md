---
source: "https://arxiv.org/abs/2604.01410v1"
title: "Assessing Pause Thresholds for empirical Translation Process Research"
author: "Devi Sri Bandaru, Michael Carl, Xinyue Ren"
year: "2026"
publication: "arXiv preprint / cs.CL"
download: "https://arxiv.org/pdf/2604.01410v1"
pdf: "https://arxiv.org/pdf/2604.01410v1"
captured_at: "2026-05-09T13:00:52Z"
updated_at: "2026-05-09T13:00:52Z"
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
- 情報源: [arxiv](https://arxiv.org/abs/2604.01410v1)
- ダウンロード: https://arxiv.org/pdf/2604.01410v1
- PDF: https://arxiv.org/pdf/2604.01410v1

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

Text production (and translations) proceeds in the form of stretches of typing, interrupted by keystroke pauses. It is often assumed that fast typing reflects unchallenged/automated translation production while long(er) typing pauses are indicative of translation problems, hurdles or difficulties. Building on a long discussion concerning the determination of pause thresholds that separate automated from presumably reflective translation processes (O'Brien, 2006; Alves and Vale, 2009; Timarova et al., 2011; Dragsted and Carl, 2013; Lacruz et al., 2014; Kumpulainen, 2015; Heilmann and Neumann 2016), this paper compares three recent approaches for computing these pause thresholds, and suggest and evaluate a novel method for computing Production Unit Breaks.

## PDF Text

A
ccepted for
Presentation
at

"Translation in Transition 8, September 2026"

Assessing Pause Thresholds for empirical Translation Process Research

Devi

S
ri

Bandaru
, Michael Carl, Xinyue

REN

CRITT, Kent State University

Text production (and translations) proceeds in the form of
stretches of
typing
,

interrupted by keystroke

paus es
. It is often

assumed that fast
typing
reflects
unchallenged
/automated t ranslation production while long(er) typing pauses are
indicative of translation problems, hurdles or difficulties.
While there has been a long
discussion concerning the determination of paus e

thresholds that separate automated
from
presumably
reflective
translation
processes

(O’Brien 2006; Alves and Vale 2009;
Timarová

et al. 2011;

Dragsted

and Carl

201
3
; Lacruz et al. 2014; Kumpulainen

2015;

Heilmann and Neumann 2016
)
,
this paper
compares

three
recent
approaches

for

computing
these
pause
thres holds
.

We draw on a set of 490
fromscratch
translation sessions from the
CRITT
TPRDB
that
cover 10
different
language pairs
to assess
three
differen t

pause thresholds

that have
recently been proposed and a novel method
.
In contrast to previous
approaches

(
e.g.,
Carl and Kay 2011
;

CoutoVale 2017) these recent approaches
take into account
translatorspecific typing speed
,

based on the distinction between
median

withinword
(WW)
interkeystroke intervals (IKIs)
and
median

betweenword (BW)
IKI
s

(
Muñoz &
Apfelthaler 2022
,
Miljanović

et al., 2025, Carl et al 2025)
. WWIKIs

are
comparatively
shorter
,

likely related to automated typing processes

when

a
typist is

completing a
word
/thought

that they have in mind
.
Average BWIKIs are longer and believed to be
related to reflective processes.

T
hey are longer than WWIKIs as they “
are more likely to
be linked to lexical retrieval
” (
Miljanović

et al
.

2025).

To compute WW and BW, w e

adopt

the

method suggested by
Miljanović

et al
.

(2025)
:

•

WW:
median

delay between

two successive

alphanumeric key strokes

•

BW:
median

delay
between

an
alphanumeric
and a
nonalphanumeric key stroke

Two cases can be distinguished
in computing

BW
s
.

WS
:
this character bigram
occurs at
the end of a word (in
alphabetical

languages) when

an alphanumeric keystroke

(W
, last
character of a word
)
is
followed by a nonalphanumeric keystroke (
S
, e.g., a blank
space
)
.

SW

occurs
at the beginning of a
new word
, when

a nonalphanumeric keystroke
(S)
is
followed by an alphanumeric keystroke

(W
, first character of the next word
)
.
Surprisingly,
the duration s

of
these two types of
BW
s

are quite different (see Figure
2
):
median
WSIKI
s

are

much shorterand their distribution more similar to WWthan
median
SWIKI
s
, suggesting that
WSi.e.,
the separator

follow s

a wordis (mentally)
more tightly linked to
the preceding word than to the following word.
Accordingly, we
base the BW computation on the longer SWIKI.
We co mpare three approaches
:

A
ccepted for
Presentation
at

"Translation in Transition 8, September 2026"

1.

Muñoz
&

Apfelthaler (2022)

distinguish

between
:

❖

RSP
:
Respite:
Unintentional
short
pause
,

d efined as 2

×

median

WWIKIs.

❖

TSP
:

TaskSegment Pause:
Intentional breaks
,

d efined as 3

×

median
BWIKIs.

2.

Miljanović

et al
.

(2025)
suggest the following two measures
:

❖

MU
D
:
MicroUnit

Delimiter:

d efined as median

WWIKI

+ 2
×

s tandard deviation

❖

B
W
: P
rocessing effort
:

i.e.,
median
BWIKIs.

(
SW
)

3.

T
his paper
:

❖

KUB
:
Keystroke Unit Break, d efined as 2

×

median

WWIKIs

(same as
RSP
)
.

❖

PUB
:
Production Unit Break,
partially
defined
based
on product data, as follows:

We suggest and evaluate a novel method for computing
Production Unit Break s

(
PUB
s
)
1
. The rational e

for
PUB
s

is grounded in the observation that text production
proceeds in chunks of a few words (say 3

words
)

rather than wordbyword
. Th at is, we
are more likely to see a n extended
lexical retrieval

(
or reflection
)

pause
every third word

rather than

every word
.
Accordingly, we take the IKI duration
of

the

quantile
for

th is ratio

(
text length/
#
chunk s
)
as a temporal
threshold

that

separat es

automated
from

reflective
processes. For instance, a ssuming that a word (in English) has on average 6 characters

and a chunk length of 3 words
, we take it that
the upper
1/(6*3 =18)

IKI
quantile
—
that
is
,

5.5% of the IKIs
—
is
related to
intentional/
reflective
/
lexical retrieval

paus ing
, while
94.5% of the
shorter
IKIs are related to automated processes
.

Table
1
shows the correlation between the four measures

and the 3
-
component
3G
L

/
3
G
U

(see
explanation
below)
.
Note that there is a perfect correlation between
SW

and
TSP

as the latter is just a multiple of the former.

RSP

TSP

MUD

PUB

3G
L

3G
U

RSP

1.00

0.85

0.44

0.63

-
0.10

0.45

TSP

0.85

1.00

0.44

0.72

-
0.09

0.48

MUD

0.44

0.44

1.00

0.51

-
0.02

0.36

PUB

0.63

0.72

0.51

1.00

0.01

0.58

3G
L

-
0.10

-
0.09

-
0.02

0.01

1.00

0.34

3G
U

0.45

0.48

0.36

0.58

0.34

1.00

Table 1: Correlation matrix of pause thresholds and GMM clusters; Figure 1: distribution of their quantiles.

Figure
2

(left)
plots the
distribution of
the
two BW measures and

on the right the

4

pause
measures

for the

490 translation sessions.

Clearly, the
distribution of
SWIKI is

1

This version of
PUB

is
different from the one in Carl et al 2025.

A
ccepted for
Presentation
at

"Translation in Transition 8, September 2026"

much
flatter

than the WSIKI, and while
TSP

and
PUB

are strongly correlated,
MUD

is
much more stretched
but

less strongly correlated with the other measures.

While th is analysis

give s

us a
picture of

how these measures
distribut e IKI into two
categories

(
assumably
automated vs. reflective processes)
, it is unclear whether these
thresholds

actually

do

their job
, and which of them in the
best

way
.
I
n order to address
this question
,

we investigate d

the distributions of IKIs and Fixation per translation
session with

Gaussian Mixture Models (GMM
s
).
GMMs identify
clusters

in the data
without
assuming theoretically
impos ed

cutoffs

(as our fo u
r methods do)
.
If
an

(e.g.
2
-
component
)

GMM consistently separates IKIs
(
or fixation durations)
into
two
distinct
distributions, this suggests
that
there may be two underlying cognitive modes
generating these
distinct
timing patterns
,
potentially corresponding to automated versus
reflective processing.

We
applied the Akaike Information Criterion

(AIC)

to determine the optimal number of
categories

based on the empirical distribution of the data.
The AIC analysis consistently
favored a threecomponent model for
IKIs
—
suggesting that

the IKI distribution s

are

actually composed of three independent processes
—
which provides strong statistical
support for the theoretical framework proposed by Muñoz and Apfelthaler (2022), who
posit
three components
:

1.

Automated Translation Routines:

Represented by
a

fastresponse component,
corresponding to motor execution and fluent text production.

2.

Unintentional Halts (Respites):

Represented by
an

intermediate component,
reflecting micropauses where cognitive resources are momentarily drawn away
from typing, yet deep problemsolving is not yet engaged.

Figure 2: Distribution of
withinword (WW) and
two types of
betweenword
BWIKIs (left)
,

and
four
different
pause thresholds
(right). Note that
the distribution of WWs is very similar to WS, suggesting
that the separator following the word is still part of the word, while SWIKIs are indicative of different
process. Note that RSPs and
T
SP (right)
constitute the same distributions as WWs and
SW
s

(left)
multipli ed

by
a factor of 2 and
3
, respectively
.

Both Figures are on a log scale.

A
ccepted for
Presentation
at

"Translation in Transition 8, September 2026"

3.

Intentional Stops (Reflective Pauses):

Represented by
a

third, slowresponse
component,
probably
corresponding to higherorder cognitive processing and
decisionmaking.

The tripartite structure
suggests
two pause thresholds

that separate the three
processes
,

which we label
3G
L
for the lower threshold and
3G
U

for the upper threshold.
In our
‘theorydriven’
analysis we have
so far
only one candidate (
RSP
s
) for
the
lower
threshold but three candidates {
TSP
,
MUD
,
PUB
} for
the upper

threshold that
is
supposed

to
separate
u nintentional
h alts from
i ntentional
s tops.

A
s shown in Table 1,
the re is virtually no

correlation between the
lower
threshold
(
3GL
) and the ‘theoretical’
metrics but a
moderate to strong positive correlation

for the
three
upper threshold
and
the
3GU
.

This suggests all three
analytic
thresholds

{
TSP
,
MUD
,
PUB
}

are not

merely arbitrary
but align with a datadriven separation of processing modes
;

the
PUB

method
in

the
best
way
.

In a full version of this paper
,

we’ll add a section on
the
GMM methods for fixation data
and dive more into detail s

of the GMM computation.

References

Alves, F., D. Vale (2009).
Probing the unit of translation in time: aspects of the design and development of a web
application for storing, annotating, and querying translation process data. Across Lang. Cult., 10 (2) (2009), pp.
251
-
273

Carl,

M., M.

Kay (2011). Gazing and typing activities during translation : a comparative study of translation units of
professional and student translators. Meta,

56

(4)

(2011), pp.

952
-
975

Carl, M; Takanori Mizowaki, Aishvarya Raj, Masaru Yamada, Devi Sri Bandaru, Xinyue REN. (2025). The behavioural
translation style space: Towards simulating the temporal dynamics of affect, behaviour, and cognition in human
translation production. SKASE Jou rnal of Translation and Interpretation, 2025; 18(2): 212
–
39 doi:
10.33542/JTI2025
-
S
-
11 (also on https://arxiv.org/abs/2507.12208)

CoutoVale, D. (2017) What does a translator do when not writing
? S. Hansen Schirra, O. Czulo, S. Hofmann (Eds.),
Empirical Modelling of Translation and Interpreting, ume 10, Language Science Press, Berlin (2017). page 209
–
237
Dragsted,

B., M.

Carl (2013) Towards a classification of translator profiles based on eyetracking and
keylogging data. J. Writ. Res.,

5

(1)

(2013), pp.

133
-
158,

10.17239/jowr
-
2013.05.01.6

Heilmann,

A., S.

Neumann (2016). Dynamic pause assessment of keystroke logged data for the detection of
complexity in translation and monolingual text production. Conference: Workshop on Computational Linguistics
for Linguistic Complexity (CL4LC),

Coling

(
2016)

Kumpulainen, M.

(2015) On the operationalisation of ‘pauses’ in translation process research. The International
Journal for Translation and Interpreting Research,

7

(2015), pp.

47
-
58. ti.106201.2015.a04

Lacruz, I., Denkowski, M., & Lavie, A. (2014). Cognitive demand and cognitive effort in postediting. In

Proceedings of
the 11th Conference of the Association for Machine Translation in the Americas

(pp. 73
-
84).

Miljanović,

Z,
F Alves
, C Brost, S Neumann (2025)
Directionality in translation: Throwing new light on an old question
.
SKASE Journal of Translation and Interpretation.

Muñoz, R., M. Apfelthaler (2022) A Task Segment Framework to study keylogged translation processes. Translation
and Interpreting, 14 (2022)

O'Brien, S.

(2006). Pauses as indicators of cognitive effort in postediting machine translation output. Across Lang.
Cult.,

7

(1)

(2006), pp.

1
-
21,

10.1556/Acr.7.2006.1.1

https://akjournals.com/view/journals/084/7/1/articlep1.xml

A
ccepted for
Presentation
at

"Translation in Transition 8, September 2026"

Timarová, Š., Dragsted, B., & Hansen, I. G. (2011).
Time lag in translation and interpreting: A methodological
exploration. In

Methods and strategies of process research

(pp. 121
-
146). John Benjamins Publishing
Company.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
