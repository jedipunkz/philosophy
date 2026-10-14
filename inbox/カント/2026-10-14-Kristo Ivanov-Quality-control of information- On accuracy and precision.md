---
source: "https://archive.org/details/diss-avh"
title: "Quality-control of information: On accuracy and precision"
author: "Kristo Ivanov"
year: "1972"
captured_at: "2026-10-14T17:30:59Z"
updated_at: "2026-10-14T17:30:59Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "カント"
query: "Critique of Practical Reason"
plain_text_url: "https://archive.org/download/diss-avh/IvanovDissQualityControlInfo copy_djvu.txt"
public_domain: true
book_text_truncated: true
subjects:
tags:
  - "近代哲学"
  - "認識論"
  - "倫理学"
status: raw
---

# Quality-control of information: On accuracy and precision

- 著者: Kristo Ivanov
- 初版: 1972
- 情報源: [archive](https://archive.org/details/diss-avh)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[カント]]
- 研究動向: [[カント-現代研究動向]]

## Full Text

THE ROYAL INSTITUTE OF TECHNOLOGY
AND

THE UNIVERSITY OF STOCKHOLM
SWEDEN

DEPARTMENT OF
INFORMATION PROCESSING
COMPUTER SCIENCE

QUALITY-CONTROL OF INFORMATION
On the Concept of Accuracy of

Information in Data-Banks and in
Management Information Systems

by Kristo Ivanov

for,

The Royal Institute of Technology Title page
Department of Information Processing
Computer Science

Fack
S$ - 104 05 Stockholm 50
Sweden

QUALITY-CONTROL OF INFORMATION

On the Concept of Accuracy of
Information in Data-Banks and in
Management Information Systems

by Kristo Ivanov

A Dissertation Presented to the Faculty of
the Royal Institute of Technology

and disputed on December 11, 1972

in candidacy for Teknologie Doktor Degree

This publication is available as order code PB-219297
from the National Technical Information Service NTIS,
Springfield, Virginia, 22151 (USA).

©) 1972 by Kristo Ivanov

The Royal Institute of Technology
Department of Information Processing
Computer Science

Fack
S - 104 05 Stockholm 50
Sweden

QUALITY-CONTROL OF INFORMATION

On the Concept of Accuracy of
Information in Data-Banks and in
Management Information Systems

by Kristo Ivanov

BIBLIOGRAPHIC DATA

Ivanov, Kristo (1972). Quality-control of information: On
the concept of accuracy of information in data-banks and
in management information systems. Stockholm: The
Royal Institute of Technology KTH. (Doctoral dissertation,
256 pages). Copies can be obtained from

<htip:/Avww. informatik.umu.se/~kivanoy/diss. html> or
from the USA's National Technical Information Service
NTIS, order # PB-219297, (about 40 US$), fax +1 703
6056900, <http:/Avww tis, cov/help/ordermethods.aspx>,

Sweden-Library Information System LIBRIS-ID:256993
<bupu/bris.kb.se-bib/256003?\W=full>, and Dissertation
Abstracts International 1974, Vol 35A, 3, p. J611-A.

(GC) 1972 by Kristo Ivanov

Title page

QUALITY-CONTROL,_OF INPORMATTON:
On the concept of accuracy of information
in Data-Banks and in Management Information
Systems.

ABSTRACT:

This paper is intended to assist those who
develop, use, maintain, audit, or in general
may be affected by so-called Data-Sanks and
Management Information Systems.

One_purpose of the paper is to recognize the
importance of accuracy, or more generally of
quality of information. Data-Banks and Manage-
ment Information Systems may typically imply
some processing performed on externally
obtained measurements and pre-processed inputs,
while their outputs may be stored and used by
people in unknown contexts,

To the extent that this happens it becomes
more difficult to expect that the quality of
information can be represented by a measure of
effectiveness of system and subsystems in
relation to operational goals. Thus,

a second purpose of this paper is to suggest
Some possibilities of attaching a measure of
quality to discrete items of information, such
as coded observations and intermediate compu-
tational results.

The paper consists of five chapters supporting
five sets of statements regarding the conse-
quences of present practices,and what can be
done to implement the most necessary improve-
ments. Illustrative examples emphasize adminis-~
trative applications such as in public planning
and in industrial manufacturing.

KEY-WORDS

Accuracy, Integrity, Privacy, Secrecy, Quality,
EDP-Auditing, System-Management, Data-Management

FOREWORD

I started the study reported in this paper with a
feeling of curiosity and persofal challenge origina-
ted from particular problems experienced during my
professional activity in industry.

After many months of work I felt disapynointment and
amazement for not belng able to frame a scientific
statement of the problem, and of course, much less
a solution to it. The problem apparently "did not
exist" according to the available literature and
reports on current research,

My fortuitous contact with the writings by C.W.
Churchman initiated a period of deep satisfaction
and allowed me to organize my subsequent work with
a feeling of being on the right way.

I terminate this study in a fourth mood: strong
appréhension,because of the implications of my
coriclusions, with respect to the possible social
impact of information systems for public planning
and administration, The same applies with respect
to the possible social impact of certain directions
of current sociological and psychological research.

IT hope that I will be proved to have beeh wrong,
In the meantime my strongest desire is to stimulate
others to further study of these issues.

T want to thank all the numerous people who in many
different ways helped and encouraged me to accomplish
this work, An attempt to enumerate them would proba-
bly result in neglecting unintentionally somebody.

Therefore,I will explicitly thank only Boérje Langefors
who first showed to me the need and possibility of a
scientific systems thinking, and whose intellectual
courage and open-mindedness made this work possible.

Secondly, IT want to acknowledge my intellectual debt
to C. West Churchman whose writings opened my way
towards a scientific and human understanding of the
issues related to this study.

March 1972

Kristo Ivanov

R,. Almstroemsgatan 3
$-113 36 Stockholm
Sweden

CONTENTS

Abstract
Foreword
Introduction

3.

Quality in the EDP - literature

1.1. On Accuracy

2. On Accuracy and Quality

3. On the Thirty-six Proposed Attributes
of Information

4, On the Importance of Quality

5, Some comments on the contents of this
Chapter. Summary

6, Conclusions from this chapter

Empirical quantitative results on Error-

Rates and Quality

2.1. Wanted: A Practical, Realistic, Empi-
rical Approach

2.2. Literature with Empirical Quantitati-
ve results

2.3. What does the Quantitative Literature
contain ?

2.4. Questions that are raised by the lite-
rature

2.5. What can be stated on the basis of the
results ?

2.6, Comments on the statements obtained
from the reviewed literature

2.7. The general setting of the empirical
quantitative results

2.8. The Commuinication-Approach to the
Accuracy problem

2.9. The reviewed literature gives practi-
cal examples of important unsolved
quality problems

2,.10.Some general considerations on the
material in this chapter.Summary

2.11.Conclusions from this chapter

Aggregation and Coding, and examples of
consequences of a limited quality concent
3.1. Aggregation and Coding: two contexts
which are less obvious
3.2. Aggregation
3.2.1. Aggregation and Errors
3.2.2. Aggregation and Urrors in
Hoonomics
3.2.3. Aggregation and the accuracy of
inventory records: a case study
3.3. Coding

43.3.1. Porcing reality to fit the model
3.3.2. A Cybernetic interpretation, and

other interpretations
3.4, General comments on the contents of
this chapter
3.5. Conclusions from this chapter

4. The definition of Quality of Information

5.

awh,

43,

Ah,

Attempts to extend the Communication-

Approach

4.1.1. "Review" in administrative pro-
cesses

4.1.2, Quality as Value and Ffficiency

Towards Accuracy and Precision

4.2.1, The concept of "Judgement"

4.2.2, Quality and Judgement in manu-
facturing and in physics

4.2.3. The role of Physics in descri-
bing controlled systems

4.2.4, Scientific method

Quality and Judgement in Data-Banks

and in Information Systems

4.3.1. The Criterion of Measurable
Error: redefining Accuracy and
Precision

4.3.2. The definition of Accuracy and
Precision

An overview on the contents of this

chapter

4.5. Conclusions from this chapter

The implementation of Quality-Control: towards
a "Handbook" for quality-control of informa-
tion

5.1. A conventional handbook for quality-

5423

5.3.

control of information
The "Conventional" handbook is not an

alternative: the role and limitations of

Statistics

5.2.1, Statement of the Problem, defi-
ning the population, illustra-
tion from Economics

5.2.2, Censuses and Surveys, Statistical
Intervals, “Rejection of Outliers

and Historical Research
5.2.3. Summary on the role and limita-
tions of Statistics

Design for Quality-Control of Informa-—

tion: Scientifically justified princi-

ples of design

5.3.1. Overview

3.2. Refining the definitions of

Accuracy and Precision
3.3. Tllustrative examples
.3.4. Mathematical formalization for
quantitative applications

3.5, FPormalization in languages for
problem-statement and automated
systems design

5.3.6, Economic aspects

General considerations on the contents

of this chapter: Summary

Conclusion from this chanter

Conclusions from this study

FoF
~

H WOAITN

Bw

FFE FF FEE
NE
EA N

N
\o

£
Le
Ww

4,39

4,50
ALBL

5.6

5.11

5.16

(of 43)

APPENDIXES:

Al

A2

A3

AY

A5

A6

AT

A8

Ag

Al1O

ALL

AL2

Conceptualization of Quality of Infor-
mation in the Electronic-Data~Processing

(EDP) ~ oriented literature

Enpirical-Quantitative results related

to input error-rates

Case-study on differences between
Perpetual Inventory Records and
Rotating Inventory Counts in a manu-
facturing plant's stock

History of Quality in Manufacturing

Basic concepts of Quality in Manu-
facturing

Basic concepts of Quality in Physics

Origin and Meaning of Accuracy and
Precision

Review of empirical results from the
reviewed literature on input quality
(refers to Appendix A2)

Statistics and the "Rejection of
Outiiers"

Historical Criticism

Suggestions for further Action and
Research

Methods for Systems Analysis

Human Thinking and Manipulation of
symbols

Information Quality and Law

Some possible implications of "Com-
munication" thinking

Some notes on the Method for this
study

REFERENCES

0.5

A6.1 (of 1)

AQ.1 (of 3)

A1O.1 (of 6)

ALL.1L (of 19)
All.1

AL1L.6
A11.10
Al1.13

A12.1 (of 2)

RL.1L (of 9)

0.6

INTRODUCTION

The motivation to start this study originated from
the results of an investigation led by the author
at the time he had managerial responsibilities in
the engineering department of a manufacturing plant.

The investigation was directed towards the analysis
of errors in the data-base describing the manufactu-
red products. Many of the errors turned out to be
other than the conventional "input" errors like
transposition, substitution of digits, ete. As a mat-
ter of fact we felt that many of these errors had at
some time to be committed in order to keep the system
going, and they should perhaps not be called "errors"
in the conventional meaning of the word. A proper ap-
preciation of their nature led us to the domains of
systems design, integration, data identification, etc.

This implies that our study is CRITICAL, that is,it
presupposes that things are NOT going well in the
area of systems design and operation, Thus, our expe-
rience has determined the general orientation of our
work and it has furnished rich unstructured empirical
material which was not explicitly utilized in this
paper.

The graph on the next page gives an overview of the
structure of this paper. Chapter 1 ~ based on

our presuppositions, experience and observations re-
sults in summaries of what the conventional literatu-
re on electronic data-processing (EDP) says about
errors and quality of information. In a similar way,
chapter 2 results in summaries of empirical quantita-
tive results on error rates. With the assistance of
some more theoretical and scientific literature, chap-
ter 3 integrates the results of the two earlier chap-
ters suggesting the typical consequences and nature

of a limited understanding of the error-quality issue
as evidenced by the reviewed literature,

Chapter 4 draws heavily upon scientific literature in
order to allow a scientifically justified definition
of some aspects of quality of information in a way
that is consistent with the suggestions set forth in
chapter 3. Finally, chapter 5 uses the newly defined
aspects of quality, refines them, and evaluates them
in the light of the earlier practical-empirical results
of chapter 2. The chapter results in particular recom-
mendations on how and where to concentrate the quali-
ty effort of an organization, and may therefore be
seen as the core of a "handbook for quality-control

of information" assisting the designers and users of

a data~bank or management information system.

More detailed contents of this paper may be found in
the previous list of "contents".

=
ria
Chee ch.it
ee os
fempirical /

/Quantita~

0.7

i Implicit

Presuppo-
sitions

Scienti-
fie meth-
od & the

Quality
‘in EDP

tive re- litera-
sults / ture
— ees
i oa a ni
oN i ae
!Conse-
quences,
examples
1
ne
a
/ nitions
log quaLli-
ity
Chase
iY,
4
ee eee
i tuaae oo)

i quality

of infor-
mation

Information-precedence graph illustrating
a rough averview of this paper in terms
of relations between the presuppositions
of this study and the conclusions from
chapters 1 to 5,

0.8

This paper contains also several apnendixes containing
both material originated by us and material written by
other authors which was selected and sometimes heavily
edited by us, This should be kept in mind when evalua-
ting the material of others, since our editings are
out of context and can never do full justice to the
authors of the original text,

Exact citations are always enclosed between quotation
marks. Both the extensive citations and the apnendixes
are judged by us as necessary for a proper understan-
ding of this paper, which spans over a very wide ran-
‘ge of professional literature, most of it not readily
available at minor locations. Our references to

" (Casual ) Documents ", sometimes abbreviated "CD",
refer to the corresponding items in appendix Al. They
originate from personal notes that we wrote in the cour-
se of the years, based on literature which we are not
able to identify. We included them because they are
valuable as testimony of thinking found in the business
arid administrative community,

A discussion of the method for our work is presented

in appendix Al2. We feel that the full implications of
the discussion are better realized after having read

the main body of this paper. At this time it will suffi-
ce to remark that we did not judge convenient to

attempt the use of a precise terminology. Our under-
standing of what is meant by information systems corres-
ponds to the ideas set forth in Sweden by B. Langefors
and in the USA e.g. by C.W.Churchman: information is
used for decisions. For the rest, the reader should

not assign any particular importance to the shifting

use of terms except for what may be inferred from the
context: the meaning of used words will emerge in the
course of the arguments in the paper.

For example, we use alternatively the words Data-RBanks,
Information Systems, Management Information Systems;
Accuracy, Quality; Model, Theory; Measurement, Obser-
vation; Administration, Organization, ete.

More explicit statement in the text are emphasized with
the mark D> at the left hand margin. Such statements
are often the basis for the specific conclusions in

the corresvonding chapter.

QUALITY IN THE EDP-LITERATURE

ON ACCURACY

There is one concept in the litterature on elec-
tronic data-processing which appears to be of
fundamental importance; especially in the context
of information systems for administrative appli-
cations.

It is the concept of ACCURACY,

We say that it appears to be of fundamental impor-
tance because the word is found whenever somebody
wants to declare the importance or value of a
so-called information system to be developed, as
well as of such a system that is already installed
and operational. Futthermore, the word is also
found in the context of emphasizing the importance
of correct input to an already developed and
installed system:

In order to determine the desirability of further
research on the natute of accuraty, a review was
made of the professional literature dealing speci-
fically with electronic data-processing. The
review included books, periodicals, research re-
ports, instructional booklets of computer manufac-—
turers and internal company reports from places

to which the author had access in the course of
his professional activity.

No intentional "a priori" selection was made of
which literature out of the above would be more
closely examined, Through browsing the focus of
attention was put on those publications that had
something stated about the nature of accuracy or
about concept intuitively related with the accuracy
issue.

ON ACCURACY AND QUALITY

Appendix Al displays an edited selection of such

a review with a view towards answering the question
“what is accuracy ?", and "is it in some sense
important - justifying further research ?".

The appendix was created for the convenience of
the readers, bringing together some material that
was spread out in many different sources. The text
had to be taken out of context and edited, which
should be kept in mind because of the danger of
misunderstandings and of not doing justice to the
authors of the original text.

3

1.2

Consideration of appendix Al introduces a multitude

of new concepts intuitively related to accuracy. They
are listed here below, We have completed the list with
those terms which are known from other occasions, in-
cluding those which denote aspects intentionally ex-
cluded from our main study, like security.

Accuracy Usefulness Trueness
Value Confidentinxnlity Relevance
Validity Consistency Reasonableness
Dependability Authenticity Pertinence
Integrity Completeness Acceptability
Correctness Reliability Refinement
Precision Degree of Detail Aporoximation
Timeliness Recency Currency
Freedom from Error Controllability Rightness
Fxactness Goodness Accessibility
Quality Availability Security
Secrecy Privacy Coverage

For the purpose of further reference in this paper,

we will often choose the word QUALTTY for representing
roughly the set of all above words. In this sense,
Quality stands for a generic attribute of informavion,.

ON THE THIRTY-SIX PROPOSED

ATTRIBUTES OF

NEORMATION

A closer analysis of the material in the appendix Al
may be performed in attempting to answer the following
questions.

lL. Is the particular concept defined ?

2. Is any justification given on whether it is, in
some sense, important ?

3. Are any recommendations given about what can be
done in order to improve the quality ?

Out of the about twenty sources in appendix AL, less
than ten appear as having attempted to define quality.
The attempts appear done in terms of conceptual rather
than operational or functional definitions: i.e. the
definition relates the concept being defined to one

or more other concepts and generally takes the form
similar to that of dictionary definitions.

For instance, Carr (1970) apparently equates RELIART-
LITY with CONSISTENCY. Lauren (1970) suggests that
RELTABILITY is the same as ACCURACY, On the other
hand I3M (#20-0006) sugpests that RELIABILITY and
ACCURACY are two distinct concepts.

1.3

Orlicky (1969) implies that QUALITY is FREEDOM PROM
ERROR but he does not offer a further definition of
error. Since INTEGRITY is mentioned by him as freedom
of error, completeness, and timeliness, one could con-
clude that quality is only one of the aspects of inte-
erity.

Rodin (1971) relates quality to the concept of IDEAL.
The ideal value corresponds to the COMPLETELY EXACT
AND CURRENT value, Since he defines quality also by
its components COMPLETENESS, PRECISION, CORRECTNESS,
AND CURRENCY, we could conclude that his concept of
exactness is equivalent to a synthesis of the three
concepts of precision, completeness, and correctness.

Montelius et al. (1970), Rodin, and a Casual
Document (1964) make use of statistical terms such as
RANDOM, ERROR LIMITS, and STANDARD DEVIATION, They do
not, however, develop the meaning of these terms in
the particular application. Since such words refer to
very elusive and misused ideas, their use by the au-
thors should be submitted to a critical evaluation.
It should have been necessary to have, for instance,
a reference to scientific-statistical literature or

a closer specification on how to obtain the relevant
observations.

Blumenthal (1969) in a book wholly dedicated to plan-
ning and development of management information systems
does not make any reference to the problem of quality,
or we were not able to find any such reference, unless
it is considered as implied by a successful design.
Quality is not ineluded neither in the input data de-
finition nor in the analysis of user requirements,
The author apparently considers quality specification
of data as a meaningless question since data are by
him defined as “uninterpreted raw statements of facts",

Carr (1970) implies that legal and administrative
applications of data are not decision making and that
their data requirements do not generate data with good
quality. From his formulation one is led to think that
bad quality in terms of observation errors results

to a large extent from the implied applications of

such data. In spite of the vagueness of the statements
this suggests important objections against Blumenthal's
conceptualization of DATA, without however assisting

in the definition of the terms.

J.C, Emery makes quality dependent on ACCURACY. Accu-
racy is seen as a QUALITATIVE characteristic of infor-
mation which attempts to substitute the quantitative
estimates of information value at lower levels of de-
cision-making. Emery seems to imply that at high le-
vels of decision-making neither accuracy nor informa-
tion value can guide design decisions for development
of information systems. The author apparently differen—
tiates between accuracy of input data and RUFINEMENT

1.4

of the estimates of input variables that are critical
in determining payoff. For the former, PERFECTION or
absolute accuracy is only a question of costs and not
limited by the nature of human knowledge; this is ap-
parently what Emery implies. For the latter he sug-

which by the way may also be limited by the INHERENT
STATISTICAL VARTABILITY of reality. Emery, however,
does not define accuracy, errors and other used terms,

W.Edwards et al. propose that quality of information
be substituted by quantity but,as far as we could see,
do not define quality. This is particularly troubling
when one knows that there are cases in which the pro-
posed Bayesian probabilistic models are being used in
military information systems. Tt is legitimate to won-
der what do the assurances mean that, for instance,

a nuclear attack carinot be triggered BY MISTAKE or

BY ERROR!

Sundgren & Lundin do not either define quality but
they attempt to consider it as one among other goals
of a public data-bank, and then they proceed showing
implicitly its nature by means of its relationships
to the other goais. The authors, however, do not jus-
tify their allocation of the quality goal to the go-
vernment: it could be conceived as being. originated
also by the citizen or by the organizations.

Montelius et al. (1970) state that the input elements
must be regarded as NEUTRAL from the VIEWPOINT of the
information process, where the process is chosen on
the basis of experience and error-controls will be
based on CHALLENGING in some way the the PRESCRIBED
STANDARD PROCESSES. The authors, however, do not deve-
lop the ideas of neutral, control of standards etc.
Therefore,their definition of error is also indeter~
mined, vague.

Owsowita & Sweetland (1965) in spite of adopting an
ambitious approach in terms of PREVENTION of errors,
apparently consider it possible to limit their study
on INPUT errors and disregard the correctness of the
information processes. Their definitions of accuracy
validity, consistency are vague in the sense that
for instance they do not explicitly state what proce-
dures should be followed in practice,to determine the
validity of a recording mechanism.

’

Vagueness and circularity of definitions is, in our
opinion also characteristic of Weinmeister's approach
(1971) and also in N.P.Edwards' approach,The latter,
for instance refers to the ACCURACY of a cost estimate,
ACCURACY of the command and control system and of its
subsystems, ACCURACY of the raw data, ACCURACY of the

1.4

1.5

value of timeliness, accuracy, reliability,
ACCURACY of the knowledge of the exact present
location of the target, and ACCURACY and age-qua-
lity of the knowledge of the target's Last posi-
tion.

ON THE IMPORTANCE OF QUALITY

Out of the reviewed literature, the Casual
Document ( Casual - Documents will be referred
to,by the abbreviation "CD") from 1966 states, for
example that accuracy is the fundamental objecti-
ve of information systems,

IBM (F20-0006) states accurate processing of data
means that the processing, besides of being perfor-
med without undetected errors and in accordance
with management's policies and instructions, FULLY
ACCOMPLISHES ITS PURPOSES.

CD (1964) seems to imply that accuracy as well
as other attributes of information such as time~-

liness and dependability,is a component of its
value,

The above thoughts lead us to the more general
and interesting matter of the relation between
the quality of information, its value and the
goals of a system. Emery touches this by stating
that it is our inability to make quantitative
estimates of information value,that forces us to
use the concept of quality in developing organi-
gational information systems,

It is apparent that from these points of view,

the quality of information is of fundamental im-
portance for information systems, This statement
is made even more interesting by the possibility
that the value-impact, or more specifically the
economic impact, of quality problems may rapidly
increase because of the proliferation of so-called
data-banks and management information systems.

Especially to the extent that the sources of
information are not the same as the users of

such information after its processing by some
system, and to the extent that the user or affec-
ted population itself cannot be limited and defi-
ned, no "data-management" will be possible. The
impact of the quality problem may have serious
consequences: this may turn out to be the case
with many public information systems unless some
scientific control is established proving the
contrary.

1.6

Physics, as a science, enjoys high status and
reputation, As an illustration,the importance

of quality of information may also be appreciated
by referring to the issue as it appears in the
physical science's information system:

Assume an engineer retrieving from a data bank
some technical data to be used in the construc~
tion of a bridge: if he gets for instance the
tensile strength of a certain kind of steel ,with-
out any indication on the accuracy and precision
of the figure, he will not be able to use such
steel in his work, Or alternatively, if he uses
the steel anyway, say nine out of every ten brid-
ges he builds will prove to not bear the load for
which they were designed !

Thus it is apparent that e.g. in "general" data-
banks, the quality of information cannot be assu-
med to be less important, We would rather say
that,unless somebody proves the contrary, the
quality of general, organizational or social,
information is still more important that in the
physical sciences since the weaker theory building
prevents testing the consequences of the use of
information with inadequate quality. It is diffi-
cult to show the collapse of a social or business
"bridge" and to put it in relation to its cause
or "steel". Nor can weaker theory be compensated
always through more “pure or raw facts" or direct
observations: a country's unemployment figures
stored in a public data bank are not more direct
or basic facts than the physical properties of
steel, stored in a technical data-bank,

There are indications that quality is in bad

shape even in the physical sciences: Branscomh
(1968) now director at the USA's National Bureau
of Standards makes this very clear when at the
same time giving a hint about the importance of
the issue. He refers to research on a particular
physical problem, cross sections for electron
eollisions, and he suggests a method for saving

a substantial part of 44 million dollars in the
course of a four-year period: "Simply by not doing
the work at all unless it is written up in such

a way that it can be evaluated and therefore
become useful" (1968). If applied toa data banks
and information systems the same statement would
read: "Do not generate or store information for
information systems at all unless its quality is
specified in such a way that it can be evaluated",

In spite of its importance, then, the quality
problem is not properly understood or is ignored
in the context of well established sciences,

Also Eisenhart (1968) and Hallert (1968, 1970)
show through their attempts to explain quality

to natural scientists and technicians, that such
explanations are badly needed in broad areas out-—
side of our immediate concern with ADMINISTRATIVE
data-banks and information systems. Their emphasis,
however, is directly relevant to design of data-
banks containing, for instance, information about
physical quantities. Since much of the experimen-
tal and theoretical work in so-called ARTIFICIAL
INTELLIGENCE and fact-answering or fact-deducing
systems is aimed initially at the simpler and
better known physical reality, one may wonder
whether such projects make allowance for storing
and processing quality specifications.

It comes, therefore, eventually as quite naturally
to learn that the situation is much worse in natio-
nal and business economic statistics. This comes
very close to the emphasis which we have given to
this study. It may be only of question of time
before in all industrialized western countries

such economic statistics is regarded as information
processing of "facts" stored in public data-banks
and information systems. A whole book by O. Morgen-
stern "On the Accuracy of Economic Cbservations$1963,
may be regarded as a qualified massive document on
the immense importance of the quality of informa-
tion.

1.5

SOME COMMENTS ON THE CONTENTS OF
THIS CHAPTER, SUMMARY

In general, the above kind of reasoning is what we
think that can be accomplished from an analysis of
the EDP literature and its definitions regarded as
CONCEPTUAL definitions, sometimes also called consti-
tutive or contextual. It is apparent that there is

no agreement among the various authors: each one of
them brings his own particular experience and intui-
tion without framing his ideas in basic consideration
of scientific method,

It is difficult to see that a further analysis along
the same lines as above would be fruitful for our
purposes. We could go on showing that some literature,
like EDP Analyzer (Feb.1968) goes about by listing
major causes of poor data, other like Casu-

al ~ Document (1970) or the auditing literature re-
presented by IBM (F20-0006) and Davis (1968) ° just
propose what should be done to improve quality in tetms
of detailed EDP validation techniques or principles of
organization, The implied scope of quality thinking
ranges from trivial keypunching errors to the almost
"“everythine" of the broad and vague concept of DATA-
MANAGEMENT, One wonders how such an ambitious and
vague data-~management as suggested by Casual
Document (1970) can be enforced on an universal social
basis for the purposes of public data-banks !

"Self-evident" truths turn out to be no self-evident
at all. For example the elimination of the human ele-
ment from the input data stream is often assumed to
result in better accuracy of the input. This is sugges-—
ted, for example, by Blumenthal (1969, p.175) and

by J.C. Emery (1969,p.38). J.P.McNerney (1961), on the
other hand, in avery well justified and interesting
study suggests that the opposite may be indeed true

in certain circumstances. How to define "the circun-
stances" ? C.W,Churchman (1968b,p.189) suggests some
of the deep implications of this issue: "objectivity"
obtained by putting more and more of the act of obser-
vation into hardware such as computers and physical
instruments greatly limits what can be observed,to the
realm of PHYSICAL reality.

vv

1.9

After a review of the EDP literature we find ourselves
in a really bad shape. Nowhere is told us how to mea-
sure quality and for what purposes, in an explicit
manner. We are not able to use the implicit definitions
in their present form as a basis for binding negotia-
tions on desired and committed quality levels between
a "buyer" anda "seller" of information. To the ex-
tent that the authors offer recommendations on what
should be done in order to improve quality, we do not
know why we should place confidence in their advice;
and even if we placed confidence and implemented their
advice we would not be able to evaluate the results

of their recommendations.

We state, therefore, that the available FDP literatu-
re does not define QUALITY OF ITNFORMATION,in the sense
that it does not explicitly support the formulation

of operational definitions of the concept. The review
gives at best some kind of insight: there appears to
be some consistency among the authors in identifying

a TIME-RELATED aspect of quality that goes under the
denominations of timeliness, recency, currency; other
aspects are not explicitly time-related, Furthermore
it appears that quality may be either associated with
the information itself or with the system generating
such information. We are not capable, however,to use
these insights in their present form.

In face of the discouraging results of our review,

we turn to the Literature on scientific method in or-
der to see what is said about definitions and opera-
tional definitions.

In the context of discussing what the CONTENT of con-
ceptual and operational definitions in science should
be Ackoff (1962,p.146) states: "In the newer branches
of science, in particular, it has become increasingly
common to define one concept in terms of others which,
if anything, are less well understood than the one
being defined and whose operational significance is
even more obscure." Later, (p.150), Ackoff suggests
five instructions for the build-up of definitions,
which we shall roughly follow in the spirit of this
paper. The basic idea, as we see it, is that defini-
tions cannot be created out of thin air; they must be
anchored in some established scientific knowledge, theo
ries.

As Churchman (1948,p.159) summarizes it:"traditional
empiricism has misread the significance of conceptions
or general ideas; it has connected them with experien-
ce of the actual world; it has connected the origin -
and validity of general ideas with antecedent expe-
rience. According to it, coneepts are formed by com-
paring particular objects, already perceived, with one
another, and then eliminating the elements in which

1.6

1,10

they disagree and retaining that which they have in
common. Concepts are thus simply memoranda of identi-
cal features in objects already perceived" (cited
from J.Dewey's "The Quest for certainty"- 1929).
Traditional empiricism has thus failed to realize the
important role of generalizations; its"ideas are dead,
incapable of performing a regulative office in new
situations." (same source).

Continuing his integrating discussion of empiricism
versus rationalism Churchman continues citing Dewey:
"the basic error of traditional theories of knowled-
ge resides in the isolation and fixation of some pha-
se of the whole process of inquiry in resolving pro-
blematic situations. Sometimes sense-data are taken:
sometimes, conceptions; sometimes, objects previously
known. An episode in a series of operational acts is
fastened upon, and then in its isolation and conse-
quent fragmentary character is made the foundation of
a theory of knowledge".

We think that no comments are necessary except for
putting the question whether what we witness in the
EDP literature is a variant of traditional empiricism
or positivism,extensively criticized by Churchman
(1968b). If this is true then we have a basis for ex-
plaining why we felt that we come nowhere up to now,
and a basis for expecting that a "practical" approach
as attempted in the next chapter will also raise pro-
blems of interpretation and generalization. We termi-
nate this chapter by consolidating earlier statements
in this chapter into the following.

CONCLUSIONS FROM THIS CHAPTER

1. The reviewed EDP literature does not offer defini-
tions of quality of information, in the sense that
no explicit supnort is found for the formulation of
operational definitions of the concept,

2. The quality of information is of fundamental im-
portance for the development and use of data banks
or information systems; this is the opinion implied
in the available EDP literature and it is also im-
plied by the lack of a scientifically justified me-
thod for cost-benefit analysis of data-banks and
information systems,

This motivates an extension of our study into the next
chapter. We will attempt there to bypass the theoreti-
cal issues by inferring on quality from what has been
and is practically done.

WANTED: A PRACTICAL, REALISTIC, EMPIRICAL APPROACH

The statements, good advices,"theories" and defini-
tions found in the previously reviewed EDP Literatu~-
re were shown to be based on shaky scientific foun-
dations. However they presumably have originated
from human experience with concrete problems,

After all, everybody will agree that there are "err-
ors" in the inputs to an EDP system, ei:g, a wrong
address of a customer, wrong quantity to be shipped
ete.

The EDP practitioner may, therefore, in a specific
situation ask for advices ot investigations on how
to improve the accuracy of inputs to the system,
Something HAS TO BE DONE and CAN BE DONE, even wi+
thout "understanding" the whole issue or being able
to define what errors are.

In the context of our research it is therefore
tempting to hop a plane and invade some business
firm having accuracy problems with some installed
information system. We can take an army of statisti-
cians with us, who will gather lots of hard data

on the problem, talk with the people who developed
and use the system, and finally apply statistical
techniques and common sense to the data in order to
suggest improvements. The object of investigation
could be the accuracy of card punching and verifica-
tion. In more sophisticated installations the object
could be the accuracy of procedures leading to the
keying of input data into on-line direct entry ter-
minals etc.

IT TURNS OUT THAT MANY SUCH INVESTIGATIONS HAVE
ALREADY BEEN DONE. The results are however spread
out in publications ranging from the subject of

EDP to applied psychology and human factors. We have
made a review of such literature which may be rele-
vant to our purposes and an overview is presented

in appendix A2 for the convenience of the readers,

If the literature shows in some sense reliable and
valuable material, we will be able to consolidate it
obtaining a set of guidelines for improving the
quality of information, obtaining implicitly-at least-—
some theoretical understanding of the quality issue,
and in any case concluding about the desirability

and nature of further study of the quality issue.

LITERATURE WITH EMPIRICAL QUANTITATIVE RESULTS

The basic selection criterium for the literature
reviewed in appendix A2 was that something should
be stated on specific ERROR RATES in the context of
information. This would hopefully take us to some
implicit concept of ERROR ahd of QUALITY. FPurther-
more we vaguely expect, departing from the familiar
context of quality control in industrial manufactu-
ving, that we might establish some "normal" error
rates which will assist us in the search of methods
for decreasing such ratesi

The appendix consists of edited selections from the
referenced papers, The selection was made with em-
phasis on the ERROR RATES rather than on abstracting
the whole paper. Although not always consistent,

we attempted to keep our own comments ahd heavier
editings aligned at the most left hand side of the
page! To the extent that the authors applied advan-
ced statistical techniques, ovr comments do not im-
ply that we have critically analyzed the calculations
and found them to be correct.

Sdrcte the edited text is taken out of context, no
guafantee can be givet that we make justice to the
authors! the readets must refer to the given sources
in order to evaluate the papers,

The review reached beyond the area of literature on
EDP, including more general and scientific literatu-
re from such areas as theoretical analysis of infor-
mation systems, applied psychology, ergonomics and
human factors, statistical journals and research

in education, As a self~imposed limitation to the
scope of our work we have not included the area of
statistics applied to censuses, surveys, validity
and reliability of psychological tests etc.

We will later attempt to show that this does not
detract from the conclusions of this chapter.

The reviewed papers and our overview may be appre-
ciated in terms of e.g.

- the reference to quantitatively specified error
rates (the basic necessary condition for being
considered in the review)

- the level of ambition, ranging from keypunch errors
as in Blirotechnische Sammlung (1956) to the consi-
deration of subtle environmental influences as for
instance in Smith (1966)

- the depth of the eventual theoretical approach,
related to the level of ambition above and to the
attempt to classify errors, discussing their na-
ture, as in Langefors (1968a), Smith (1966), Root
& Sadacca (1967), Owsowitz & Sweetland Gee.

To the extent that such theoretical approach is
found in EDP literature, it could be included in
appendix Al, as we in fact did with the Owsowitz &
Sweetland's discussion of approaches to error,

253

2.3

- originality of the approach, in considering in-
fluences which were ignored by most other reviewed
investigations, e.g. the Berglund & Larson's study
of punched card layout, Smith's or Root & Sadacca's
study of so-called content or omission errors.
Another aspect of originality of approach may be
the use of original methods in detecting or correc-
ting errors, as for instance the development of
predicting routines by Carlson (1963) based on the
decision-tree heuristics suggested by Newell, Shaw
and Simon.

- generality of the approach; in covering many possi-
ble aspects of the error or quality problem, as
done by Smith (1966) or by EDP Analyzer (1971la,
1971b). EDP Analyzer, ‘however, obtains generality
thanks to its overview approach, mostly referring
to relevant sources of literature,

- clarity in the explanation of used concepts or
pérformed investigations, preventing ambiguity in
the mind of the readeri Anh excellent example of
desirable clarity is given in the Bergiund & Larson
paper,

All the above modes of appreciation were determinant
of the selective abstracting in appendix A2.

WHAT DOES THE
QUANTITATIVE LITERATURE CONTAIN ?

Many of the reports result from the application of
statistical methods.

Variables are generally related to

- types of entry devices, types of keyboards

- use of punch verification, check digits etc.

- skills or experience of operators at entry devices

- grouping, length, composition (alpha content, etc.)
of messages, punched card layout etc,

- aural versus visual presentation of stimulus (ori-
final)

~ rate of presentation of stimulus or time-pressure
on entry

- use of mnemonic codes or Letter-pattern familiar
codes

- management or supervision emphasis on accuracy or
speed of entry

- allocation of entry functions between the creator
of source document and operator of entry device

- use of pre-assigned media such as pre-punched cards,
badges for personal identification or identifica-
tion of remote terminal,

2.4

2.4

Performance of the entry process or of the handling
of information is generally expressed in terms of
ERROR RATES which relate to the degree of identity
between stimulus or original message and the output
from the human subject or from the entry device
activated the the human operator. Sometimes the
check of identity is extended to the output from
some editing routines in the computer system.

Whenever the nature of the information handling
process prevents a simple one-to-one correspondence
between input and output, new performance measures
are proposed either in terms of communication the-~
oty applied e.g. by Van Gigch to models of "integra-
tive behaviour", or in terms of especially develo-
ped error-classification schemes as by Berglund &
Larson.

Smith offers an interesting list of alternative

etiteria for data collection performance:

- time per entry, would be meaningful only in those
cases when a substantial portion of the subject's
time is devoted to the entry process

- rate of information flow (as proposed by e.g,
Cardozo & Leopold and by Van Gigch) has no frame
of referenoe for inclusion of omitted or incomple-
te messages, but it is interesting for its combi-
ning of speed and accuracy in orie measure,

- number of consecutive good entries between mista-
kes would be of no practical utility because, Smith
says, the computer system normaily has to analyze
all input messages. Martin's and Norman's discus-
sion of accuracy in communication networks and
Langefors' reference to the importance of many
small transactions for the impact of errors on
administrative EDP systems, however, suggest that
such measure may be meaningful in some respects,

- ratio of volume and time of supervisory (adminis-
trative)messages to system input messages is said
to be too dependent on many environmental charac-
teristics. It is however interesting since it seems
to imply the important concept of error-correction,

Smith finally chooses the PERCENTAGE OF INACCURATE
OR INAPPROPRIATE ENTRIES as the most UNIVERSAL CRIT=-
RIUM OF DATA COLLECTION PERFORMANCE.

QUESTIONS THAT ARE RAISED BY THE LITERATURE

While reviewing the literature, several questions

are raised beyond the above discussion of the meaning
of performance and error rates. The questions reside
in how to compare and use error rates in face of
differences and ambiguities in the nature of the
reported figures,

2.5

Error rates are either in terms of errorless entries
(i.e. all entries except those with AT LEAST ONE
error), where entries consist of different amounts
of symbols (message lengths), or they may be also
expressed in terms of individual symbol errors.
Symbols for message syntax (such as field separation
or field and record identification) may or may not
be included in the error statistics,

Error figures may include errors that were detected
and possibly corrected by the operator himself du-~
ring the entry process, but such figures may also
refer only the undetected or residual errors.
Uncorrectable errors may designate the same’ thing
as so-called residual errors, i.e. those errors
whieh were not corrected at the last step, before
entry into systen computation, Uhcortectable errors,
however, sometimes designate those errors which are
detected by checks at the entry device but are amena-
bie to error in the source document: the etror is
not caused by the operator and therefore is not
correctable by him without heavy loss of go-called
etfictency in the entry process.

Error rates after detection and correctioti by opera-
tor himself at entry, should not be equated to error
rates at input to the computer editing and validation
routines since sometimes entry verification (e.g.
punch verification) is done by another operator in

an independent entry procedure, and/or by verfica-
tion-validation checks by software incorporated to
the entry device.

Comparison of error rates for messages of different
lengths is furthermore complicated by the use of e.g.
prepunched sections in the messages and by many
ambiguities in the terminology. DIGITS may commonly
denote arabic numerals but sometimes they are used
in expressions like "10-digit numeric data words"

in which case the term is understood to be used also
for denoting alphabetic characters. In such case is
“digit' equivalent to ALPHA-NUMERIC character or
SYMBOL, but “symbol" may rather be used to include
special signs and letters from foreign languages,

not belonging to a particular alphabet, LETTER is
often used as synonym to CHARACTER. Finally one
meets ambiguities in the. meaning of terms like DATA
which may stand for all the previous concepts of
digit, character, symbol etc., but also for CODE,
MESSAGE, ITEM, and in general the ENTRY'S DATA REPRE-
SENTATION,

The most serious difficulties of interpretation of
results, in the sense of being able to compare and
use the reported error rates, however, stem from
the environment in which they were obtained. It may
be e.g. FIELD or LABORATORY. If field, it can still
be field trial-or field experiment, and field opera-—
tional {as in Kramer, 1970). If laboratory, it may

2.5

2.6

still simulate field inputs (as in Root & Sadacca,
1967). Eventually some results are a flat statement
of experience, presumably based on field or laborato-
ry reports (as in Orlicky, 1969). Carlson (19635 pre-
sents a study of historical field data.

Such different environmental conditions may explain
the appearance of error rates such as percentages

in terms of types of inputs (e.g. percent of megsa-
ges of 10-digits length which were in error)
terms of persons (e.g. percent of entries made by

Different envirormental settings imply also sa:
special handlings and exceptions in the processing

of original error information: for instance scmetimes
errors in the "cents" positions of dollar amounts
were not counted as errors ~ the seme happening to
those original errors which could be ascribed to

have been caused by poor handwriting on the oripi-
nal source document. In other cases some symbols

were not used which could be visually or aurally
confused with other symbols (e.g. M can be aurally
confused with N). In one case the investigators
report that they did not count as errors those which
would conceivably have been prevented or detected

in an operational field environment, by means e.g.

of better training or programmed validity checks,

WHAT CAN BE STATED ON THE BASIS OF THE RESULTS ?

It is apparent that any advices based on the reviewed
literature must be qualified by "if", "possibly"
etc., including recommendation of careful evaluation
of the original literature.

At the level of general advice we could gather
guidelines like the following:

1. Errors increase as the number of characters in
the data code (code Lengtt ) increases. Longer
codes should be avoided, if not possible, they
should be devided in smaller units of three or
four characters, e.g. 123-4567 instead of 1224567,

2.. The characters used in data codes should avoid
digits or letters that can easily be confused
with each other, such as I versus 1,2 versus @,
slash (/) or virgule C3) versus number 1, letter
OQ versus Q, O versus 6, U versus V. :

3. Nonsignificant codes should avoid characters that
when pronounced sound alike, such as M versus N,

B versus P,

4, Significant or meaningful data codes are prefer-

red over non-significant since this factilitat

oO

recall by the human coder and reduces errors,
For example M and F are expected to be mere re-
liable for MALE and FEMALE than 1 and 2.

6.

10,

11.

125

13.

14,

15.

16,

17.

2.7

In the cases when the code is structured of both
alpha and numeric characters, similar types of
characters (alpha, respectively numeric) should
be grouped and not dispersed throughout the code.
For example, fewer errors eccur in a three cha-
racter code where the structure is alpha-alpha-
numeric (e.g: HW5) than alpha-numeric-alpha

(e.g. HSW),

When designing a code number system, try to avoid
the chance of double occurrences of a character,
Repeating characters are a major source of trans-~-
etiption error: the chance of error is greater

in transcribing 31146 than it is when transcri-
bing 31046.

Use check digits whenever possible and appropria-
te.

Avoid the use of variable length, fixed order
puneh card layout unless the higher probability
of errors are offset by other advantages.

In the design of number check routines in verifi-
cation consider that most digit manipulation er-
rors are caused by single digit substitution,
followed by omissions,

In general, use sight verification when data is
of language type, i.e. in terms of words and
phrases, and key verification when the data must
be compared on a character-by-character basis.
Consider that there are limits to the accuracy

of human sight-verification capability: the lower
the frequency of errors to be detected, the less
percentage of them will be in fact detected by
the human sight verifier,

In selecting punch machine operators, consider
that the fastest operators are also those who tend
to make the less mistakes. (In addition there

are psychological testsfor selecting such person-
nel).

Easy correction of operator mistakes at entry
devices tends to enhance both the speed and accu-
racy of input. The same is true for easy detection
in terms e.g. of answer-back tones at direct en-
try devices.

Confirmatory answer-back tones should not be too
long since they can lead to other kinds of erra-
tic behaviour by operators who get impatient,

The profitability of punch verification should

be continuously questioned since it deletes a
very limited propwtion of punch errors.

Consider source errors, sometimes called content,
event, omission, procedural, misidentification,
miscount,etc., generally more important in per-
cent and seriousness of consequences than other
entry operator errors and hardware or communica-—
tion-links errors.

No preference, in general,can be stated for the
use of alpha or numeric codes in a particular
system.

18.

19.

20,

21.

22.

23.

24,

25.

26.

27.

28.

No preference, in general, can be stated on whe-~
ther the person making the data-entry as operator
of an entry device should be the same as the
person creating the original source information,
No statement, in general, can be done about the
effect of using pre+assigned media such as pre-
punched cards on the accuracy of input.

Coding errors can be reduced at the entry stage

by providing keypunch operators with knowledge

on the set of the possible codes. This effect is
greatest with menemonic codes.

There seems to be a substantial advantage in
accuracy by copying a code by hand immediately
beneath the original. Forms, dockets,etc,, should
be designed in such a way that this is possible.
Ten-key keyboards yield a significantly lower
errot rate and are preferred by operators, compa-
red to other devices such as Levers, matrix key-
boards, rotary knobs and telephone sets,

Speed of human sight-ckeck of errors is highest
for groupings of 3 to 4 digits in numeric mate-
rial, and it is inversely related to the frequen-
ey of errors to be detected. The percentage of
undetected errors increases with the higher speed
of checking but it is not influenced by variation
in grouping,

Por several tasks including keyboard entry and
telephone dialing, grouping of digits by 3's and
4's is consistently best in speed with no tenden-
cy to differences in error rate. Users often state
preferences for larger groups than those producing
best performance. .
For codes of a given length (number of characters)
coding errors tend to be proportional to the alpha
content.

It cannot be stated that the use of mnemonic codes
reduces coding errors. However, Letter~pattern
familiarity affects coding errors: codes contai-
ning letter pairs in familiar sequences (e.g. AT,
BY, OK) have lower error rates. (Example of mnemo-
nic code: OVH for “overhead"),

Time pressure on making data entries does not need
to affect the rate of initial original errors of
entry, but it may contribute to higher rates of
residual errors by affecting the rate of both
detection and correction of mistakes by the ope-
rator at the point of entry.

The rate of correct information that is retrieva-
ble from coded information depends not only upon
the error rate of the coding process but also

upon the detectability of errors. This latter con-
cept includes consideration of the ratio of the
number of codes used to the total number of possi-
ble codes which may obtained from all combinations
of the allowable character set,

2.6

2.9

COMMENTS ON THE STATEMENTS OBTAINED FROM
THE REVIEWED LITERATURE

The guidelines which were suggested in the previous
section appear much more useful. than any speculations
about proper definitions of quality, accuracy etc.
However this does not imply that we have bypassed

the conceptual difficulties of the quality problem.
Maybe the guidelines cannot be applied to the parti-
cular installation for which one wants to improve
accuracy. Maybe they are not so useful as we wish.

Agreement among different authors terminates at a
very low level of ambition indeed, For instance,
Biirotechnische Sammlung (1956) , Carlson (1963), and
Smith (1966) agree quite well on such a simple mattor
as the proportion of digit manipulation errors which
may be expected to consist of single digit substitu-
tion, say more than 60 %. But concerning omissions,
Bliirotechnische Sammlung gives the figure 7 % while
the other two give about 20 %. Conrad & Hull (1967)
on their part show with their wide variation of per-~
cent figures that they require mcuh closer analysis
for appro priate interpretation.

Wright (1952) suggests that 0,3;6 and 7 are those
digits which lead to most unreadable and ambiguous
readings (combined). Owsowitz and Sweetland (1965)
suggest instead that 2 is the most incorrectly re-
produced digit, Upon closer analysis it will be found
that Owsowitz and Sweetland included even letters

in their investigation, leading to the 2 being very
often confused with the letter 2, and this explains
the differences between the two findings.

Concerning the use of either alpha or numeric charac-—
ters in the construction of codes, EDP Analyzer
(1971b) refers to Davidson who advises the use of
numeric codes only. On the other hand Conrad & Hull
(1967) in the context of manual copying of codes,
state that the conclusion that digit codes are pre-
ferable to letter codes "...must be thickly surroun-
ded by qualifications." mainly because of the possi-
bility of utilizing language habits. Furthermore,
Owsowitz & Sweetland explicitly state that the fact
that error rate for alpha codes is generally several
times greater than for numeric codes, this does not
mean that they should be avoided; the decision will
depend upon several other considerations since alpha
codes can transmit a good deal more information per
character than numeric codes can,

The ambiguity of advices and guidelines does not
decrease but obviously rather increase when reaching
more subtle problems. Let's just illustrate the case
of whether the operator making the data-entry at the
entry device should be the same as the person who
creates the original document or codes the event-—
observation:

A reviewer of Root & Sadacca's paper concludes that
their findings seem to justify the following.

The direct entry method (same person doing both jobs)
seems to be recommended where the best total sveed
and accuracy are needed, where there is no reason %o
save the message generator's time by delegating the
data entry task to another, and where he could be
taught typing efficiency (e.g. more than 35 w.p.m,
i.e: words per minute). Examples of this type of si-
tuation are mainly on the military field but could
also be, for instance, the air traffic controller's
task, However, where the message generator is a cos-~
tly specialist (e.g. a hospital doctor), or where he
is not and cannot be taught to be a fast typist, then
his time could be saved by having a clerk to do the
data entry. But in such a case, when ERRORS might
sometimes be vital (e.g. drug prespriptions in hospi-
tal), it could well be advisable for the specialist
to enter certain details directly, especially since
the experiment showed significantly worse errors when
transcription was by another person.".

Thus, the reviewer concludes, it seems clear that
any decision on the method must depend on a DETAILED
. AND THOROUGH ANALYSIS OF THE DATA ENTRY TASK AND OF
THE SITUATION REQUIREMENTS. CAUTION IS NEEDED IN THE
INTERPRETATION OF THIS EXPERIMENT, More research
would be desirable to enable better guidance.
(Shackel, 1969, p.159).

Smith (1966) on his hand states that although his
study"showed no clear preference for clerical, group
or individual production worker reporting, the choice
might be dictated by the NATURE OF THE PRODUCTION
CYCLE OR PLOW. All other things appearing to be equal,
it would be preferable for the person recording events
to be the one most affected by the ACCURACY and TIMB-~
LINESS of the entries. The complexity of messages
transmitted and the variety of types of transactions
made by an individual can be limited by assigning

him uniquely to a device at a single work station
where his primary duties are related to a production
task rather than data collection, If personnel are
required to make only occasional entries in a varie-
ty of message construction forms, individual diffe-
rences in performance can play a dominant role in es-
tablishing the expected accuracy. In these cases whe-
re procedural mistakes might be caused by Low volume
reporting from a work station or by message complexi-
ty, a clerk with primary emphasis on recording the
data could be the best choice. Reporting events by
groups compromises these issues and reduces the

ot

required quantity of relatively expensive data col-
lection terminals, but increases the non~productive
travel time to an input terminal,"

By means of comments we may now realize some of the
manyfold implications of the above problem. From
what is said it looks like if ACCURACY were some
composite function of MOTIVATION (for high accuracy),
and VARTABILITY & FREQUENCY - say FAMILIARITY - of
certain tasks.

Familiarity may be seen as referring to the perfor-
mance of the original "object" task as well as its
original observation and coding, but it may also re-
fer to the task of entering the coded observation,
directly or by transeription from e.g. an original
form, into the system, Smith taiks about both tasks
as"recording" probably because in his production en
vironment the entry was directly made by the workers-
observers into the remote terminals, In his work
Smith keeps anyway the distinction between the two
tasks by means of classification of errors in diffe-
rent types; the distinction, however, cannot be con-
sidered so clear as in Root & Sadacca's study.

Concerning the choice itself between direct and in-
direct entry, one criterion appears to be the maxi-
mization of familiarity, but at the same time a
trade-off is envisaged against motivation,

It is difficult to find support for the suggestion
that direct entry is recommended when best total
speed and accuracy are needed. Indirect entry, by
saving the time of the observer-coder might be prefe-
rable, not so much because the time of costly specia-
lists is expensive, but rather because of lower rates
of certain kinds of errors in the performance, obser-
vation and coding of the original "object" task, The
lower rates of such errors might well compensate

and more than compensate for an increase of the rate
of other less important transcription errors.

The above comments are concerned with allocation of
data-entry and observation-coding tasks. A similar
discussion could be done, but is left outside the
scope of this paper, concerning the use of pre-punched
cards and other computer-prepared turnaround docu~
ments, In that case we would have Davidson's sugges-
tion, as referred by EDP Analyzer (1971b, p-9), to

be qualified by the empirical findings of Smith
(1966, p.16,66) and Kramer (1970, p.246). These last
two authors suggest that the use of pre-punched
cards, badges for individual or machine identifica-
tion etc. may have a negative effect on accuracy
because of increased opportunities of certain procedu-
ral mistakes which are not offset by the system's
detection and correction features.

2.7

What conclusions can we draw from the above comments
on the statements obtained from the reviewed Litera-
ture ? We do not know how to use the reported fi-
gures of "hard" research on error rates. We do not
know how much confidence to place on general advices
not even in those cases where they are based on
experimental confirmation of common-sense fuesses.
We do not know what ACCURACY really is: we are rather
told what it might depend on, in certain circumstan-—
ces.

In order to formulate the only conclusion which
appears to be safe, we are tempted to borrow the
words from some of the reviewed papers and formula-
te the following:

"In any specific situation, the decisions for impro-
ving accuracy will have to be based on an analysis
of the task and of the situation requirements, of
the nature of the task cycle and flow."

And this is about the same as saying nothing, a
meager result indeed, considering the scope and
statistical ambitions of the reviewed papers and the
ambitions of our own study! At this point we feel
that it is also doubtful whether some support can be
found for Shackel's statement that "more research
would be desirable to enable better guidance", if
one thinks of the research being done along the same
lines as the one we have reviewed.

THE GENERAL SETTING OF THE
EMPIRICAL QUANTITATIVE RESULTS

In order to come out of the impasse, let's go back
to statement No. 16 in the earlier list of state~
ments of the section where we asked ourselves:
“what can be stated on the basis of the results

(of the review of literature with empirical quantita-
tive results) ?",

Statement No. 16 is of our own make, and it was sug~
gested by some of the literature, It states the fol-
lowing,

"Consider source errors, sometimes called content,
event, omission, procedural, misidentification, mis-—
count, ete., generally more important in percent and
seriousness of consequences than other entry-operator,
hardware or communication-links errors,"

A review of the literature indicates that errors

in EDP hardware and communication links are often
associated with figures of about 1:100,000 or less.
Similarly, entry-operator e.g. punch-machine operator
errors are about at 1:100 in order of magnitude.

on

2,13

(Let us for the moment forget the problem of inter-
preting the units of such figures. The reader is
referred for this purpose to the earlier discussion
in this chapter. )

However, as soon as the literature touches on the
subject of what we in statement 16 called "source"
errors, error rates seem to soar up to 1:10 or
1:5 without difficulties.

Figure 2.1 is an attempt to visualize the general
experimental setting which in the reviewed literatu-
re conducted to the mentioned rates above.

Figure 2.1 shows a source with an ongoing series of
activities which are observed and coded in a coding
process (2). Such codes may generally be registered
on an original document like a form which is subse-
quently used in a data entry process (3). The data
entry may, as for instance in the case of keypunching
of cards, be followed by a correction process (4)
that in the example would be a keypunch verification
(and correction). The verified inputs so obtained
are then, possibly after being transmitted through

a communication channel, be submitted to so-called
editing, validation or diagnostic preparatory progr-
ams of the computer, (6) prior to their input and
use in the normal information processing programs.

The source, which also could be visualized as a set
of processes, is also designated by a number, (1)
in spite of standing for more than proper informa-
tion processes as the following ones.

Figure 2,1 suggests that the source might begin by
contributing to the total error rate with what we
call in this context "source errors". The coding
process results in the information set that we label
ORIGINAL INPUTS or SOURCE DOCUMENTS, possibly supple-
mented with check digits or control totals. This is
the first existent information set in the sense of
the reviewed literature (information related to the
EDP system) and it will, besides the previously men-
tioned undefined source errors, include CODING ERRORS
added in the course of the coding process.

The data-entry process transcribes the original in-
puts or source documents to e.g. card, tape or disc,
i.e, to INPUTS IN MACHINE READABLE FORM, and it will
contribute with TRANSCRIPTION ERRORS. This data-entry
process may use devices with built-in programmed
verification and validation (in the sense explained
by e.g. EDP Analyzer, 1971a) and correction,

2,14

\ (1)

Y Source

Noe aid
Observations with oe ae
SOURCE ERRORS !

Coding (2)

Original inputs or |
Source documents with added |
CODING ERRORS |
Data Entry or
Transcription

(3)

To tard,tape,

dis

(Self) Verification
Validation

(Self) Correction

Machine-readable
inputs with added

transcription Correction 7 (4)
errors

Sight ,Key

verification

(detection)

Correction
Verified inputs —

less corrected errors |

| Conmunication (5)
Communicated inputs |
with transmission Computer Input (6)
errors | Processing
Computer Editing

Run of diagnostic

programs
System inputs with
residual errors
Information
processing (7)

(Normal run)

Figure 2,1

The general setting for experiments
and measures leading to the reviewed
empirical quantitative results,

The prefix "self" stands for the feature being in-
corporated to the entry device, rather than being
performed by the human operator,

To the extent that verification and correction are
not performed or sufficient at the data-entry process
stage, they will be performed separately at the
following correction stage. Correction is seen to
include the detection sub-process (e.g. sight of

key as in keypunch verification) and the correction
itself, leading to what we labeled VERIFIED INPUTS.
Verification, validation and correction in the data-
entry process (3) and in the correction process (4)
will delete some of the errors previously introduced
in the chain, but - at least theoretically - may
introduce own errors which we label CORRECTION ERRORS
(e.g. correcting an input which is actually right,

to become wrong - Klemmer, 1959, is one author who
considers this problein).

The vetified inputs may be submitted to a transmission
process by a communication system resulting in what

we label COMMUNICATED INPUTS which include undetec-
ted TRANSMISSION ERRORS (we delete here the detailed
breakdown of the communication problem - that is
considered e.g. by Norman, LO71). Such commuhicated
inputs are finally used in the computer input process
(6) leading to the final INPUTS which include what is
usually labeled as RESIDUAL ERRORS.

What does this visualized experimental setting tell
us 7? In the first place it calls our attention on the
possibility of placing emphasis on different stages
of the overall process. Before going any further

let us associate figure 2.1 with another similar
figure that is found in the scientific literature.

THE COMMUNICATION-APPROACH
TO THE ACCURACY PROBLEM

In discussing the case of a "discrete channel with
noise" in the context of his mathematical theory of
communication, C.E. Shannon (1949) considers the
problem of a signal that is perturbed by a chance
variable - called NOISE - during trausmission or at
one or the other of the terminals. He considers the
case in which the received signal is not the sane as
that sent out by the transmitter, and when it does
not always undergo the same change in transmission
(distortion), i.e. most generally he considers the
case when the RECEIVED SIGNAL TS NOT A DEFINITE FUNC-
PION OF THE TRANSMITTED SIGNAL.

In order to develop a theorem that gives a direct
intuitive interpretation of the average uncertainty
of the correctness of the received signal, Shannon

considers a communication system and an observer

(or auxiliary device). THE OBSERVER CAN SEE BOTH
WHAT IS SENT AND WHAT IS RECOVERED (WITH ERRORS DUE
to NOISE), THIS OBSERVER NOTES THE ERRORS IN THE
RECOVERED MESSAGE AND TRANSMITS DATA TO THE RECEIVING
POINT OVER A "CORRECTION CHANNEL" TO ENABLE THE
RECEIVER TO CORRECT THE ERRORS, The situation is
indicated schematically by Shannon in the figure 2,2
below which was slightly changed by us for the pur-
pose of clarity in the following discussion;

Source of
information

"True" message
ee
Message

Transmitter

| Si enal Observer

Receiver

Message

Device < Correction data

Destination

Pigure 2.2

Schematic diagram of a general communication
and correction system

It is now apparent that the communication approach
to the accuracy problem, as illustrated in figure
2.2, can only be applied to the process steps (3),
(4), (5), (6) of the earlier figure 2.1 where we
visualized the general setting of the empirical
quantitative results. To the extent that one is able
to consider the output of an EDP program as a
FUNCTION of the input information, there is a possi~
bility to apply the communication approach also

to step (7). In any case this appears to be the
implicit basis for present thinking in AUDITING OF
EDP SYSTEMS.

The important thing to note in the context of apply-
ing the communication approach is that in all cases
one assumes the existence of an "objective" OBSERVER
WHO "KNOWS" THE TRUTH OR CORRECTNESS OF TWO OUT OTP
THE THREE ELEMENTS -INPUT, ~FUNCTION, -OUTPUT and

is therefore in position of"authority"for "CORRECTING"
THE THIRD ONE. For example, if one knows that the
customer address printed by the computer-printer on
the invoice is not true (i.e. wrong), and also knows
that the program updating the customer file is true
(1.e, right), then one can deduce that the input to
the program was not true (i.e. wrong). The "one who
knows the truth" is what we labeled as the "objective"
observer. In specific situations, the objective obser
ver appears sometimes disguised:under other labels
such as system analyst, manager, decision maker, in-
vestigator, researcher, verifier.

From the above it is apparent that it will be trou-
blesome to apply the communication approach to those
steps of figure 2.1 which include truths of doubt-
ful observability, such as steps (1) + that is
dealing with events outside the frame of the review-—
ed literature-,(2), and (7).

Let us consider the process (2) ~ coding -.

What is a RIGHT or ACCURATE input to the coding pro-
cess, since such input is appearing prior to our
formalization in terms of information ? Whenever the
reviewed literature has touched on related problems,
e.g. Owsowitz & Sweetland (1965) and Van Gigch (1970a
and 1970b), it has assumed the existence of a certain
set of right inputs; this is a particularly impor-
tant assumption as remarked by Weaver (Shannon & Wea~
ver, 1949), since it emphasizes that the general
setting of the analytical communication studies deals
with only the first level, A, out of three possible
levels of communication problems:

A. How ACCURATELY can the symbols of communication
be transmitted ? ( The technical problem).

B. How PRECISELY do the transmitted symbols convey
the desired meaning ? (The semantic problem)

C. How EFFECTIVELY does the received meaning affect
conduct in the desired way ? (The effectiveness
problem).

2.9

Shannon & Weaver's mathematical theory of communica-
tion then deals only with level A. It is therefore
left unsaid whether the subdivision in the other

two levels (semantic and effectiveness), as well

as Weaver's use of the words ACCURACY versus PRECT-
SION and EFFECTIVENESS,are in some sense scientifi-
cally justified. In our opinion, the distinction
among these words as suggested by Weaver does not
assist our research on the issue of quality of in-~
formation,

In any case it is now clear that most reviewed em—
pirical results, as suggested by appendix A2, adopt
the communication approach and as such deal with

all processes of figure 2.1 except (1), (2) and (7).
An analysis of the quality of information in these
terms apparently disregards the most important as~
pects of quality relative to data banks and informa-
tion systems for administrative control.

Furthermore, we do not know of any proof showing
that such most important aspects are intractable in
terms of other approaches,other than the communica-~
tion approach. On the contrary, the physical scien-
ces make extensive use of the concepts of accuracy
and precision in situationswhere no "observer" is
idealized who can compare a supposedly "true" input
to the output etc. and where the inputs are conside-
red to be members of a set of possible true inputs.
The example of quality concepts from the physical
sciences suggests alternative approaches to the pro~
blem.

THE REVIEWED LITERATURE GIVES PRACTICAL EXAMPLES
OF IMPORTANT UNSOLVED QUALITY PROBLEMS

EDP Analyzer (1968) ,referrea in appendix Al, in our
opinion touches on some symptoms of the most impor-
tant quality problems when Listing EVENTS THAT po
NOT CONFORM TO POLICY among one of the major causes
of poor data. At the same time it differentiates
such cause from INCORRECT CODING OF CLASSIFICATION
FIELDS, suggesting that in terms of our figure 2.1
both causes may correspond to the source and coding
processes (1) and (2)..

Smith (1966) classifies mistakes in FORMAT errors,
CONTENT errors, and EVENT errors. Format errors are
by him defined as items that can be detected and
screened from system input (such as wrong message
length, illegal characters or malfunction of data
entry equipment), Content errors are items that have
correct form, but can be detected as logically in-
consistent (such as shop status contradictions,
unusual quantities, wrong machine or operator desig-
nations), Event errors are those items that have cor-
rect form and are logically processable, but prove
inconsistent after subsequent entries or upon use,

(such as omitted entries, failure to correct detec-—
ted mistakes).

Smith furthermore points out that some comparisons
between error rates in the field and in the labora-—
tory experiment must make allowance for the fact
that CONTENT MISTAKES WOULD BE FEWER IN THE EXPERI-
MENT BECAUSE MISIDENTIFICATION AND MISCOUNTS WERE
NOT INVOLVED.

Referring to the card-verification procedure used
to check the accuracy of card punching operations,
Smith states that such card verification procedure
can only identify mechanical and copying mistakes,
HAVING NO FRAME OF REFERENCE to analyze event des-—
cription; misidentification or miscount, and many
format inconsistencies,

What is this "frame of reference" that Smith is re-
ferring to ? We think that it has much to do with

the theoretical understanding of the quality of in-
formation that we ourselves are looking for. The
expression "frame of reference" unhappily belongs to
the class of heavily misused words ("concept" is
another such misused word) but it appears that Smith
considers his classification scheme as the frame of
reference appropriate for the object of his study.

We cannot accept such frame of reference for our pur-
poses since Smith does not motivate it with conside-
rations of scientific method which assure its genera-
lity and indicate how it will be used,

For instance, why does Smith consider corrections to
good entries as CONTENT mistakes, while failure to
correct detected mistakes is considered as an EVENT
mistake ? (See p. 5,6,39,40 of Smith, 1966)(It is
difficult to conclude whether some inconsistency in
the allocation of mistakes to different classes is
an unintended print error.)

More important, however, we see the problem of eva-
luating Smith's frame of reference or classification
with, for instance Root & Sadaccas (1967) classifi-
cation in SPELLING, OMISSION, CONTENT and SEQUENCE.
By omission,they mean any failure to enter a required
item of information; by content,they mean wrong in-
formation such as wrong identification of the nature
of an event or object (e.g. "tank" instead of
"truck"); by sequence, they mean information itens
in a message not being in the proper sequence,

We are then led to believe that Root & Sadacca's
omission, content and sequence all are included in
Smith's event type.

Similar comparisons may be done with Kramer's
PROCEDURAL and OMTSSION errors; Berglund & Larson's
errors due to the NATURE OF MATERIAL (source wuncer-
tainties) and OMISSIONS; EDP Analyzer (1971a)
classification of data (ana implicitly-errors ?) in
TEXT, JARGON, and NON SENSE.

This means to us that without a general theoretical
understanding of the quality issue one will not be
able to compare own error rates with Smith's rest-
dual rate at about 4 % of entries or Root & Sadacca's
approximate rate of 2 %. And we have now seen that
the difficulty to compare is caused by much more

deep reasons than any ambiguity on what is meant by
digit-character-symbol, or ambiguity on the nature

of the message in terms of number of digits-including
or not including pre-punched sections etc. (See the
discussions ih the earlier sections of this chapter).

Finally we see that even practical, empirical approa-
ches to the problem of quality of information raise
unavoidable important theoretical questions. Such
questions appear in spite of using a communication-
naive setting; because this setting is being applied
to some complex aspects of the information systems
problem,

SOME GENERAL CONSIDERATIONS ON THE MATERIAL IN

THIS CHAPTER. SUMMARY,

In the previous chapter we had met difficulties in
defining and measuring the quality of information.
We raised the question whether such difficulties
could be by-passed, avoided by applying a so-called
practical, realistic approach to the problem.

An extensive preview of literature containing empiri-
cal quantitative results disclosed a great number

of figures on error rates which proved difficult to
interprete and apply in practical situations, The
same appeared as a result of analysis of statements
containing advices on what to do in order to improve
quality, where the statements were explicitly or im-
plicitly obtained from the reviewed literature,

The remarkably higher rate of certain types of errors
reported in the literature, suggested that they refer
red to certain steps of a general information-proces~
sing sequence, This sequence was visualized in terms
of a figure which encompassed the measurement setting
of most reported figures on error rates. This set-
ting was seen to be the same as the one used to il-
lustrate the quality-accuracy issue in communication
systems,

The communication approach to the quality of infor-
mation was seen to be too limited for the purposes
of application to data banks and information systems,
Attempts to apply this approach to such environment
raise many more questions than are able to answer,
put they suggest that the unanswered questions are
the most important ones justifying our further study
in that direction.

CONCLUSIONS FROM THIS CHAPTER

1. Most available measures of information quality
in quantitative terms assume a concept of quality
in terms of communication theory (theory of signal
transmission),

2. The utilization of above measures in a particular
information system, and the development of other
necessary measures require a broader concept of
quality which can be made operational,

The above two statements were formulated from the
material contained in the sections of this chapter,
specifically: questions that are raised by the Lli-
terature, comments on the statements obtained from
the reviewed literature, the communication approach
to the accuracy problem, and - the reviewed litera-
ture gives practical examples of important unsolved
quality problems.

Before attempting the development of a broader conh-
cept of quality we will dedicate the next chapter to
illustrate two possible consequences of lacking

such a broader concept. This illustration is inten-
ded as an additional support to the conclusion of
the previous chapter regarding the importance of the
quality issue, and it will at the same time supply

a concrete feeling for the implications of the theo-
retical developments of the broader concept,

AGGREGATION AND CODING

A LIMITED QUALITY CONCEP

AGGREGATION AND CODING:
TWO CONTEXTS WHICH ARE LESS OBVIOUS

In the attempt of illustrating the implications of

a narrow understanding of the information quality,
it is easy to think about the waste of research and
activities which are to be based on false informa-
tion premises. Alternatively one may think about the
damage inflicted to business and society resulting
from the implementation of false conclusions derived
from false premises.

Within the more limited scope of this paper we in-
stead intend to illustrate the way in which the
narrow understanding of the quality issue hides
important exposures in the context of two quite fa-
miliar and supposedly non-controversial activities
of the data-bank and information system environment,

AGGREGATION

Aggregation, in the context of control systems, is
described by. one author as being the description of
a system by a lower order model, lower in the sense
that the model variables in a given sense represent
"averages" of the system variables. This given sense
may be a mathematical function defining an "index"
of the original variables.

Emery (1969) expresses the function of aggregation
in the context of design and implementation of or-
ganizational planning and control systems, as being
one way of obtaining DATA COMPRESSION. The purpose
of data compression in an organization is said to

be the reduction of the VOLIME OF AVATLABLE DATA

in order not to swamp the organization with trivial
information and in order not to reduce too severely
their information content, The aggregation of data
over unwanted CLASSIFICATION DIMENSIONS, IRRELEVANT
FOR THE PURPOSE AT HAND, attains reduction of volume.
For instance, sales transactions might be aggregated
along the dimensions of customer, salesman, industry,
and geographic region, leaving the data classified
in terms of the remaining dimensions - item and time
period. ‘

What is said above has a strong intuitive appeal,

it recalls obvious experiences we all have had in

the context of simple EDP applications, and is clear-—
ly related to much traditional thinking in statistics
where one talks about SUPPICTENT STATISTICS or con-
tractions of observations, sufficient for the PURPO-
SES TO WHICH THE OBSERVATIONS MAY BE PUT, and espe-
cially providing a SIGNIFICANT SAVING IN THE MECHANI-
CAL LABOR OF STORING AND PRESENTING DATA.

3.2.41

The same view on aggregation may be held in many con-
texts of applied research and operations analysis.

A good illustration of such contexts is given for in-
stance by Ackoff (1962, p.126) who, in the context

of omitting uncontrolled variables in the building
and use of models states that the aggregation of se~
veral variables does not exactly omit any of the va-
riables, but it does reduce the number that have to
be considered. Ackoff also gives some examples of
aggregations from business applications.

AGGREGATION AND ERRORS

Up to now everything seems OK; our interest in ag-
gregation appeared the first time because of what.
is satd on ERRORS in the context of aggregation:
this is what we will cover next with a question in
our mitid - "does aggregation help to attain better
accuracy ?",

Emery (1969) in discussing qualitative aspects of the
value of ACCURACY of information states that in the
case of decision processes dealing with unaggregated
data, the VALUE of information may be highly SENSITI-
VE TO ERRORS. When data are aggregated for high-level
decisions, Emery says, THE VALUE OF GREAT ACCURACY
drops off sharply. The author illustrates this point
with the case of an error in a bank account balance,
which may be very expensive indeed, while its possi-
ble impact on high level decisions using aggregate
bank-deposits by state, is much weaker.

While Emery makes his statements in the same context
as ours, i.e, data-banks and information systems, 1+
is interesting to note that his views seem to be ana-
log to those expressed e.g. by Ackoff (1962, p.126)
in the much more constrained context of a well struc-
tured applied research. Ackoff then states that where
variables are aggregated, the ERROR (in the estimate
of the outcome) which is introduced is ROUGHLY PRO-
PORTIONAL TO THE RATIO OF THE WITHIN-AGGREGATION VA-
RIANCE TO THE BETWEEN-AGGREGATION VARIANCE, Put in
another way, he says, it is desirable to make the va-~
riables aggregated as homogeneous as possible and the
aggregations as heterogeneous as possible.

The above makes us believe that an interpretation of
such view on aggregation and related error-accuracy
problems, is that the variables refer to the compo-
nents of a so-called NEARLY DECOMPOSABLE HIERARCHIC
SYSTEM. Such near-decomposability implies that the
short-run behavior of each of the component subsys-
tems is APPROXIMATELY INDEPENDENT of the short-run
behavior of the other components, and that in the
long run, the behavior of any one of the components
depends IN ONLY AN AGGREGATE WAY on the behavior of
the other components. (Simon, 1969, p.100).

34212

3.3

The striking implication of this proposal is that

one knows the implications of aggregation and rela-
ted errors, if the system-problem is dssumed to have
been already solved in the sense suggested by Simon
(1969) or Langefors (1968b). One of the serious dif-

ficulties of such an assumption, however, is the com-

mon knowledge that information must be used and er-
ros estimated in business and social contexts where
obviously the assumption does not hold, since nobody
claims to have designed the system or defined its
goals ete. Furthermore, many data-banks will use and
store information which has been generated and which
will be used in unknown contexts, certainly not de-
signed nor understood in Simon's or Langefors' sys-
tem terms.

AGGREGATION AND ERRORS IN ECONOMICS

Applications of economic seience in business and na-
tiorial planning makes use of an enormous quantity

and variety of data or statistics which can very well

be immagined to be stored in data banks, In most in-
dustrialized western countries such implementations
of data-bank are to some extent already being done,
and it might be only a-question of time before it
becomes common-place.

Applications of economics to business and national

planning are much closer to our context of data-banks

and information systems for administrative control,
than the trivial applications to bank accounts
or assumed well structured problems of applied re-

search mentioned in the last section of this chapter.

ait is therefore important what O.Morgenstern has
to report from an extensive experience in the sub-
ject matter, in his book "On the Accuracy of Econo-
mic Observations" (1963). We edited the following.

A whole economy is entirely inaccessible for com-
putation, unless drastic simplifications are in-
troduced, This leads to the process of aggregation,
i.e. the formation of larger entities from myriads
of components, which presents one of the most im-
portant but also most troublesome problems of eco-
nomics. Too much aggregation mixes the unmixable
and gives us models that are easy to handle but
with low, if any, power of resolution, By aggrega-
ting, errors of a new kind are introduced, (p.101)

It is possible that the influence of one error

which drives a number in one direction is exactly
offset by the influence of another errors doing

the opposite, leading to a "true" figure for our
observation. But we have not MADE a true observa-~
tion ! The notion that errors cancel out is wide-
spread and when not explicitly stated, it appears

as the almost inevitable argument of investigators
when they are pressed to say why their statistics
should be acceptable. YET ANY STATEMENT THAT
ERRORS "CANCEL", NEUTRALIZE EACH OTHER'S INPLU--
ENCE, HAS TO BE PROVED. Such proofs are difficult
and whether a "proof" is acceptable of not is not
easy to decide. (p.53)

The mere repeated "checking" of the transcription
of figures from some source and their correct
transfer to other papers is no substitute for the
determination of errors of observation and their
significance for deductions and inferences.

It is also necessary that WORTHLESS STATISTICS

BE COMPLETELY AND MERCILESSLY REJECTED ON THE
GROUND THAT IT IS USUALLY BETTER TO SAY NOTHING
THAN TO GIVE WRONG INFORMATION WHICH ~ QUITE APART
FROM ITS PRACTICAL, POLITICAL ABUSE - in turn mis-—
leads hosts of later investigators who are not al-
ways able to check the quality of the data pro-
cessed by earlier investigators. THIS IS ESPECIAL-~
LY IMPORTANT IF DATA ARE TO BE USED IN EXTENSIVE
AGGREGATIONS. When elaborate calculations are
needed that are difficult to set up, this mislea-
ding information may make the use of high-speed
computing machines meaningless.(p.54)

How can one evaluate what Morgenstemsays in the con-
text of economics against for instance Fmery's much
more optimistic view of the matter ? Maybe the an-
swer lies in the assumptions. Maybe the answer is
suggested by what Morgenstern says on the success of
modern physics:
In physics errors were recognized for a very long
time; but they were held to be a secondary nuisan-
ce, to be neglected and to be ignored by the
THEORY. Or as Brillouin expresses it:" The assump-
tion was that errors could be made ‘as small as
might be desired', by careful instrumentation, and
played no essential role. Modern physics had to
get rid of these unrealistic schemes, and it was
indispensable to recognize the fundamental impor-
tance of errors, together with the unpleasant fact
that they cannot be made ‘as small as desired! and
must be included in the theory." (p.61)

This implies that aggregations will not imply any
difficulties when they are performed in an informa-
tion system dealing with problems which are well ex--
plained by available theories, like physics,

The situation will become much worse in the
context of social events such as found in business
and government where no established theory exists.

Such insight on the problems of errors and accuracy
in the context of aggregation is impossible within
the much narrower frame of accuracy suggested by the
literature reviewed in the earlier chapters.

3.2.3

3.5

AGGREGATION AND THE ACCURACY OF INVENTORY RECORDS:
A CASE STUDY

Appendix A3 presents some details of a case study

on so-called inventory differences as signs of the
inaccuracies of inventory records of the stock of

completed parts in a plant manufacturing electro-

mechanical machines.

The results of the case study are not fully exploi-
ted in this paper, but some of them can be used to
illustrate the vagueness of the implications of ag
gregation in a situation in which the system pro-
blem has not been solved, as well to illustrate the
complexity of the quality-accuracy issue in terms
of the vague SOURCE and CODING errors mentioned ih
the last chapter,

Plant management and auditors consider the accuracy
of the inventory records to be a very important mat-
ter. Does this imply that they do not care of the
differences since they will in some sense "cancel
out" in aggregations over time and over items ?
Certainly not, to judge from the existence itself

of a rotating inventory count and from the tecurring
investigations on the nature of the differences found
through these counts. Also, certainly not, to judge
from the richness of the number of reports and va-
riables in the follow-up statistics on inventory dif-
ferences; most of them not usable for low-level deci-
sion making.

It actually appears that’ higher levels of manage-
ment are very dependent on detailed knowledge of
differences, They are not interested in the possi-
bility that positive differences "cancel out" nega~-
tive differences. They must keep negative differen-
ces down to some minimum because of e.g,

1. Danger of running in Line-stop leading to delays
in delivery of products and waste of idle resour-
ees.

2. Incurrence in extra costs for placement of addi-
tional emergency orders,

3. Requirement to protect stockholders.

4. Losses leading to charges on the product price,

Positive differences must be kept down to a minimum
because of e.g.

1. Losses from interest on investment on too high
stock,

2. Losses from not being able to take advantage of
the maximum allowable write-down of stock value.

3. Protection of the public from an over-evaluation
of assets

3.6

In the appendix A3 it is possible to see that no easy
conclusions can be drawn on the aggregate effects

of the errors listed in the summary list of errors
leading to inventory differences. If anything, it

is for instance possible to notice from the summa-

ry table of the first investigation (1964) that

there are two kinds of causes that only contribute

to negative differences and never can cancel each
other.

The most interesting insight, however, from the case
study is that even if the differences cancel-out,
the problem is to know HOW they cancel out, and to
what extent the way of cancelling is acceptable in
face of management's above listed seven objectives
in keepirig differences to a minimum: for instance,
what is the amplitude and frequency of fluctuations
around the "true" value,that would be considered

to be acceptable by certain particular stockholders ?
An evaluation of aggregation and its errors is thus
seen to require an understanding of the total sys-
tems

Outside of the particular subject of aggregation,

the case study also shows the nature of many source,
coding or observation errors, as they were labeled
in the review of literature on quality, Ighoring
such kinds of errors appears to be equivalent to ig-
riorihg the larger system in which the purely tech-
nital EDP system is contained. It is not surprising
that, to the extent that error rates can be measured
in some way, the larger contribution to suth rates
in a complex social system will originate outside

the strictly defined technical EDP subsystem. Con-
centrating the quality effort on the technical aspects
may thus be a grave suboptimization: it is something
like avoiding the cause of difference listed under
number 9 in the appendix A3, wrong punching, when
the other 29 listed causes are not considered at all.

The list of causes of differences also shows how ma-
ny so-called "human factor" errors may be in their
turn considered as caused by the inflexibility of
the EDP program itself (for instance see points 13,
23, 26). Such facts should have far reaching organi-~
gational implications in future complex systems.

Finally we can very concretely notice the absence of
the "objective" observer of the "true" inputs. The
rotating inventory clerk is the verifier-observer of
the stock clerks' activities; the three reported in-
vestigations were performed by verifiers of the veri-
fiers, i.e. by objective observers of both the rota-
ting inventory and stock clerks; and our own present
study can be seen as a further step of verification
or "objectiveness" - we are discussing the meaning
of the accuracy of those who checked the accuracy of
the rotating and perpetual inventory system. A dis-
cussion of the accuracy of the follow-up statistics
summarized in appendix A3 would be a concrete docu-
ment of the vagueness of the complex accuracy issue.

343

3.7

A superficial examination of the summary of the con-
tents of follow-up statistics on inventory differen-
ces, as displayed in appendix A3, discloses the cre-
ation of a great number of "aggregation variables",
out of the basic original observations of differen-
ces, These extensive statistics and tabulations re-
late only indirectly to the basic problems as illus~
trated by the list of causes of differences, This
suggests to us the applicability of Ackoff's state-
ment originated from a number of experiences in the
field of operations research: "The less we understand
a phenomenon, the more variables we require to exp-
lain it. Hence the manager who does not fully under-~
stand the phenomenon that Ke tontrols plays it
'safe"™ and wants as much information as he can get."

This suggésts that the vagueness of the complex ac-
curacy issue leads to the use of aggregations whose
aim is not datascompression for preventing the need
to commtuhicate Large volumes of trivial ihformation,
Aggregations may then rather be used in the attempt
to remediate Lack of knowledge on the nature of er-
rors or lack of control on them, by massive data-
processing of the information that happens to be
avallable on them. Such perspective is just one al-~
ternative to the image of tomorrow's sales manager
who, when confronted with an unfavorable trend of
sales, sits down at an on-line terminal requesting
all possible aggregations and statistical tests to

‘be performed on past sales transactions, "searching

for patterns in the data". We obviously question

the belief that such a procedure will substitute the
direct understanding of the original object system;
the available resources might better be applied to
such an understanding.

CODING

There is some evidence that the broad subject of
coding in the context of information systems and
data-banks is not completely understood.

In the EDP literature, coding has at its best been
considered as a communication tool, and it has been
evaluated in technical terms: communication~economy
through a channel, economy of identification in the
storage and retrieval of information etc, Codes have
been developed with primary attention given to machi-
ne processes, in order to facilitate machine opera
tions, such as the "tight" coding on many 80-column
punched cards.

In more recent years, as suggested by some of the

literature reviewed in the previous chapters, some
people have realized the need to "design human fac~
tors into" the code structure in order to minimize

for instance transcription or digit-manipulation
errors by humans who are then also considered as
"communication channels",

In view of coming ambitious projects for implementa-
tion of complex data-banks and information systems,
we think that the time has come to enlarge the above
view on the meaning of coding.

One possibility is to integrate the communication-—
approach into the body of modern organization theory.
An organization may create CATEGORTES for classify-
ing situations and events. Such classification sche-
ines are the basis for the program-evoking aspects of
communication: once the event has been classified,
the appropriate program can be executed. (March &
Siniony 1958, p.162)

The alove can be illustrated as follows. As soon
one knows in a manufacturing plant that a particular
item is not a detail part but rather an assembly,

to be bought from a local vendor to whom the plant
will have to consign all of the detail parts to be
assembled, - then the particular item is to be coded
CQ-509 in the perpetual inventory file, This file
will later be used e.g. in the requirement genera-
tion program.

Another illustration may be taken from the EDP appli-
cation for updating perpetual inventory records of
the manufacturing plant. If a particular item was
previously requisitioned from stock in order to be
quality-inspected, and it is found that it is no
usable, and it cannot be reworked but it must rather
be scrapped, then the transaction to the EDP applica-
tion program must be coded 5119 08. The transaction
will then be processed updating stock status and la-
ter it will also be used e.g. in accounting applica-
tions.

A second possibility to enlarge the view on the
meaning of coding is to regard it as one method of
expressing measurements: one attempts to assign e.g.
objects to classes, while in others one tries to es-
tablish a specific relationship or attempts to assign
numbers, A coding system will then be a language for
class assignment, whose rules are the means by which
a decision-maker uses the information expressed in
the language. A perfectly "adequate" Language of
class assignments must meet all potential informatio-
nal requirements, that is, must provide an exhaustive
classification. (Churchman, 1961, p. 106)

The striking consequence of this enlarged view of
coding is that it becomes much more than a question
of economy in communications, storage, and retrieval.
It becomes an information problem that reaches beyond
hardware, software or human-factors considerations.

3.3.1

3.9
FORCING REALITY TO FIT THE MODEL

CODING MAY THEN BE REGARDED. AS THE COUPLING, INTER-
FACE, OR MEASUREMENT PROCESS LINKING THE REAL WORLD
(OBJECT SYSTEM WHICH IS TO BE CONTROLLED) To THE
MODEL REPRESENTED BY THE INFORMATION SYSTEM OR BY
THE INFORMAL HANDLING OF THE CODED INFORMATION,

The importance of this insight for our inquiry in
errors and quality of information derives from the
possibility to regard CODING ERRORS not only as cau-
sed by the “human factor" or by non-understanding of
the model by the human coder. The coding errors may
also be seen as caused by NON-ADEQUATENESS OF THE
MODEL itself, i.e. by MODEL ERRORS in not taking
into account relevant aspects of the teal world,
including the social system and humans who are sup-
posed to work with the model.

As a simple illustration, consider the research re-
ported by Cardozo & Leopold (1963) and its extension
py Van Gigceh (1970a and 1970b). Theit results sug-
gest the existence of a maxiinum humatt communication
load, above which human communication error rates
are expected to increase steeply. Codes belonging to
a code=stheme can be interpreted in terms of such
comtiunication load. If the load is too high many co-
ding errors will be committed, Which is the conclu-
sion % Before we had available the referenced re-
search, or to the extent that it is not accepted as
part of "established" psychological theory, we would
have claimed that the errors were OBSERVATION or
CODING errors, requiring e.g. better discipline and
training of the human subjects. To the extent that
the research is accepted and perhaps incorporated
in a theory, we would instead claim that the errors
were MODEL errors: the system designer has allowed
the disorganized growth of interdependent EDP pro-
grams which impose their own coding scheme without
consideration to the known "facts" on human constra-
ints. The system designer will have to improve his
training and discipline.

The above could have been reached by sheer common
sense, What the enlarged view of coding enables us

to do is, hopefully, to integrate and evaluate many
different concurrent interpretations of coding errors
in a particular situation,in terms of scientific
method,

As a more complex illustration of the implications

of the enlarged view on coding, reconsider the case
of coding of items in the perpetual inventory file of
the manufacturing plant. What would happen in a case
when some but not all detail parts are to be consi-
gned to the vendor from whom the plant buys the com-
pleted assembly ? Or in the case when some of the
detall parts turn out to be also assemblies in their
turn ? What will the coder do, which will his "infor-
mation load" be if many such sometimes unique excep-
tions appear every day, and which will the consequen-
ces of his coding decision be in terms of the system
designer's or programmer's understanding of the co-
der's environment,code scheme, and program logic ?

343.2

3.10

Inaccuracies may in such situations arise from the
coder's attempts to FORCE REALITY TO FIT THE MODEL.
In our case study on inventory differences this could
be illustrated by the stock clerk reporting a stock
location as being 999 whenever he has to store a
certain item in a "third" stock location, He knows
that in this way he will prevent errors of the type
listed under number 26 of the list on causes of dif-
ferences: parts not found because the EDP program
allows a reporting of at most two stock locations

for the same item, and deletes the record of the
first upon reporting of the third. The stock clerk
knows that each time he reads 999 as second stock lo-
cation of an item, he must refer to his own manual
records or to a common stock location where many
such items are stapled leading perhaps to errors of
the type listed under number 14: parts are not found
because too many different parts are stored at the
same stock location, being easy to overlook them.

How to evaluate such errors in other environments
(object systems) which may be much more complex than
the stockroom of a manufacturing plant, especially
whenever there are no resources for adapting inflexi-
ble information systems, coding systems and EDP pro-
grams, to a changing environment ?

A striking cybernetic interpretation of the deep im-
plications of what has been said above, for the pos-
sibilities to control organizations, is given by
Beer (1966). It is reviewed here below in terms of
edited abstracts. (p.310)

A CYBERNETIC INTERPRETATION, AND OTHER INTERPRETATIONS

On the shop floor, one can always find an example
of a machine-loading arrangement which "controls
the flow and allocation of material around the
shop". What it actually does is to make desperate
attempts to keep the job cards posted as they are
returned - to provide something like an accurate
reflection of what is going on. To the objective
cybernetician, then, the shop floor is a control
system generating variety for the purpose of con-
trolling the planning office, and not vice versa.
The reasons for this unhappy example have been
formally uncovered. They are: lack of requisite
variety, disobedience of the theorems of communi-
cation about channel capacity and so on, AND ABOVE
ALL, A STATIC, INADEQUATE, UNADAPTIVE MODEL OF
WHAT THE WORLD SITUATION USED TO BE LIKE SEVERAL
YEARS AGO. aes

Fortunately, however, control procedures have a

way of keeping themselves viable and of rectifying
their mistakes: by means of "ad hoc" comparisons

3.11

of real events with their predictors, the control
subsystem struggles in a horribly inefficient way
to acquire a certain adaptability. GIVEN THAT THE
PROPORTION OF NEW EVENTS IS QUITE LOW (new events
is namely what this kind of control is very bad
at handling), and given the capability to organi-
ze the feedback information, everything usually
can run fairly smoothly.

The trouble, however, is that in the course of
time, BECAUSE OF THE VITAL NECESSITY FOR CREATING
CONTINUOUS AND DETAILED FEEDBACK, the control
organization must be allowed to grow and become
prohibitively expensive in terms of personnel,
facilities, and equipment including large-scale
electronic data processing equipment,

But nobody notices that this 1s a fault in the
state of affairs, because it is too familiar,

and because the energies of all concerned are to-
tally absorbed in arguing the merits of alterna-
tive computers. Typically, the absurdities inhe-
reht in the situation are obscured by the APPEA-
RANCE of modernity and technical competence which
all this activity betokens.

Beer's cybernetic interpretation is paralleled by
Blumenthal's system-positivistic interpretation and
description of troubles at higher organizational le-
vels. (Blumenthal, 1969, p.197). Disorganized incre-
mental growth of so-called systems where
system is piled on system, or a continuing series
of minor enhancements is made to existing systems,
in an attempt to generate relevant management in-
formation as‘a hopefully serendipitous by-product
accompanying the production of increasingly vaster
quantities of irrelevant, unused, or merely his-
toric data...

As the information pile-up occurs another system
modification or addition is created to produce
ostensibly only that which is wanted in the situ-
ation...

This process is a form of change, true; but it is
only marginally and fitfully adaptable change.
Ultimately the patchwork collapses, Systems beco-
me moribund, and, like dead horses, no longer res-
pond to the whip,

Blumenthal goes on giving, in a positivistic mood,
the answer to this problem: a new dynamics of adap-
table systoms growth. This appears to us,a new
aspect and name for tho ovon-pervading problem of
model building - in this case infoimahion oyotomes.
What bothers us, however, is the implication of all
what was said above for our issue of quality-accura~
cy of information, as well for the related "errors"
made at distinct organizational levels.

3.12

It is easy to immagine that the terrible descrip-
tions of serious problems by Beer and Rlumenthal

must mean also serious things happening to the er-
rors and accuracy at different levels of the infor-
mation system. We tried to illustrate this by means
of the case study on inventory differences but the
illustration is obviously incomplete in many respects,

The reviewed literature does not offer any example
of the problem. We could guess about an example by
reading"between the lines" of Orlicky's discussion
of so-called integrity of an average manufacturing
routing file, and its maintenance. (Orlicky,1969,p,153)

Such file will consist of several tens of thou-
sands of records encompassing active, inactive,
semiobsolete, and obsolete parts. Each of these
records carries the prescribed sequence of opera-
tions, their descriptions, the routing to the va-
rious departments and machine tools within these,
job standards, and the required tooling, not to
mention part master data in the record header.

This file is constantly affected by so many chan-~-
ges in manufacturing method, standards, tooling,
engineering changes, machine tool procurement,
downgrading and retirement, shop reobganizations,
etc., that true file maintenance becomes a night-
mare. This is so BECAUSE MANY TYPES OF CHANGE LT-
TERALLY EXPLODE THROUGHOUT THE FILE (such as in
case of adoption of a new class of cutting tools,
changes in departmental boundaries, or the acqui-
sition of new productive equipment). A single
such change MAY CALL FOR HUNDREDS, OR EVEN THOU-
SANDS, OF PARTS TO BE REROUTED, operations to be
added or deleted, and methods, standards, and too-
ling revised accordingly.

It is possible to guess what such requirements mean
in terms of impact on the ACCURACY OF CODING. Orli-
cky goes on stating that the key to this problem is
the staffing and budget provided for file maintenan-—
ce, Our broad concept of the nature of the coding
process allows us to frame this statement in concert
with e.g. Beer's and Blumenthal'ts: when things begin
to look like as in Orlicky's description,a better
contribution to overall accuracy might come from an
improved system design (with built-in human factors
considerations) rather than from increased staffing
and budget for file maintenance,

We hold that the quality of information, particular~
ly as expressed in the nature and rates of coding er-
rors, may be an important indicator of the adequacy
of system design or of the model. Up to now it has
been regarded as an indicator mainly of the coding
and observation process itself.

3.4

3.5

3.13

GENERAL COMMENTS ON THE CONTENTS OF THIS CHAPTER

After showing in the previous chapter that the
empirical approach to quality of information assu-
mes a narrow coricept of quality, and that it does
not dispense a sound theoretical understanding of
the issue, we attempted in this chapter to show how
a too narrow understanding of the issue misses
important problems arising eig. during use of data
banks and information systems.

Two such less obvious problems are the considera-
tions of accuracy in the context of coding and ag-
gregation. The optimistic view on the aggregation
of data assumes that the system-problem is already
solved, and this was seen to be not motivated as
suggested from problems in economic science and
in a case study on inventory differences in a manu-
facturing plant. The optimistic view on coding re-
gards it as a communication tool for efficient ma-
chine processing and misses the possibility of re-
garding it as a measurement process where "errors"
may indicate model inadequacies.

We may eventually note that the issues of coding and
ageregation appear to be closely related. When one
aggregates for example sales transactions along the
dimensions of customer, salesman, industry and geo-
graphic region, this corresponds to the creation of
a new set of sales data where the above dimensions
do not make difference and are therefore coded as
belonging to one same class. The assumption is that
the new code defines a class of information that
will be useful for some particular decision.

CONCLUSIONS FROM THIS CHAPTER

1. Without an understanding of the information-qua-
lity issue, aggregation of data may be uncritical-
ly accepted as being error-free in the context of
high-level decision making.

2. Without an understanding of the information-—qua-
lity issue, it is possible to miss the evaluation
of coding errors and coding difficulties as
symptoms of inadequate model building or system
design.

As a contribution to improved system design, we

shall try in the next chapter to develop the concepts
of ACCURACY and PRECISION as two aspects of the broa-
der understanding of quality of information, to be
operationalized and "built into" data-banks or in-
formation systems for administrative control.

THE DEFINITION OF QUALITY OF INFORMATION

ATTEMPTS TO EXTEND THE COMMUNICATION-APPROACH

Before developing our proposal for the concepts of
accuracy and precision, let us recall that the quality
problem in the reviewed literature was shown to be con-
ceived in terms of the degree of identity between input
and output of a communication channel. We call this

the "communication-naive" approach which regards data-
banks in terms of telephone or telegraph system, where
the output is an "'identity function" of the input.

At the next higher level of sophistication we can place
the stiggested extension of the communication approach,
where the output of the channel is a general functicn of
the input. The channel can then be seen as a "data-pro-
cessing" channel, and the function can be regarded as

a data-process, program. The quality of information at
the output is now related to the degree to which it cor-
responds to what would have been expected if the right
program had been applied to the input. Particular pro-
blems arise when evaluating quantitatively the quality
of output, if the one-to-one correspondence is lost
between input and output, as suggested e.g. by Van Gigch
(1970) in trying to apply the information~-load idea to
the quality issue.

In both the above cases, we have necessity of the presen-
ce of an observer, manager, auditor, client, supplier of
input, or the like, who has the AUTHORITY TO STATE THE
TRUTH or quality of some out of the three elements: in-
put, program, and output. Knowing the truth-status of

two of them, it is possible to infer the need to correct
the third one. For instance, the output is stated to be
wrong (the client complains), the program is stated to

be right (the programmer or the auditor of the EDP system
states this), and therefore the conclusion follows that
the input by the clerk must have been wrong. A special
case occurs when the input is declared wrong and is re-
jected to the system's environment, Two basic concepts
involved in this thinking appear to be the DETECTION and
CORRECTION of ERRORS, and the quality control system may
be visualized as below

Quality Control System

Error Detection. | | Error Correction|
Subsystem | | ' Subsystem |

|
|

Fig. 4.1

hoe

The identity-function, communication approach and the
general-function, data-processing approaches were shown
to be most problematic if applied to the context of
data-banks and information systems, outside the limited
and highly structured situation depicted in the most of
the reviewed literature. It is therefore motivated to
question the applicability of the approach suggested by
Montelius et ali (1970) to the theoretical analysis of
errors in integrated information systems. As our edited
abstract from their work in appendix Al shows, the au-
thors state that "we" must commit ourselves to a desired,
so~-to~say right process on the basis of experience, Such
assumption on its rightness is seen as a kind of compro-
mise on truth in terms of a prescribed standard which
does not dispense of an error-control, Furthermore, they
state that the input elements must be regarded as "nenu-
tral" from the viewpoint of the considered process,

It has not been possible here to evaluate the meaning of
their statements or applicability of their approach since
the authors do not develop their idea of standard, con-
trol of the standard, and neutrality of the input ele-
ments. We guess, however, that intuitively their thin-
king is in terms of what we called above the data-proces-
sing, general function approach, It is conceivable that
such an approach is fruitful in a highly structured,
self-contained, optimally designed system, Consider, for
example the case of a customer who complains that he has
been bilied a wrong quantity of merchandise, If the sys-
tem is designed optimally in the sense that it follows
Langefors' theoretical analysis of information systems
(1968b), a precedence analysis would assist in the de-
termination of the causal error-chain, possibly "untrue"
values of relevant variables. This could lead to the
identification of a wrong input or of a wrong process,

A succedence analysis would likewise assist in the deter-
mination of e.g. errors propagated by the discovered
wrong input or process to other parts of the system..

In a similar way, a detection process may be performed
through the use of a succedence analysis upon the dis-
covery of a clerical input error in setting the unit pri-
ce of a merchandise. Correction processes would follow
along the same pathways. The ideas are somewhat explored
in the paper by Montelius et al. and in another under-
graduate paper based on their approach (Danielsson &
Helin, LOZL1). They also take up the question whether

the error should be corrected through the system itself
or outside of it, e.g. by apologizing for a misspelled
address or name,without mailing a duplicate of the whole
invoicey or e.g. requesting an authorized correction of
input delivered to the system, such as a wrong bill from
the vendor of a detail part that was assembled into a
shipped product.

The above approach has an intuitive appeal which suggests
that it might be useful in some situations. At the same
time it makes some questionable assumptions such as in
the context of choice of correction method, which is

made on the basis of the incremental value and cost of

a message or transaction. This assumes, however, that

gay

an individual transaction can be costed by itself, and
it is exactly the difficulty to do this that Leads to
alternative approaches in terms of systems theory. The
issue is discussed by Churchma+ (1961, p.321) in the
context of assets and transactions where he convincingly
criticizes what he calls the "transaction theory of va-
dues",

Furthermore, we can name just one more assumption of

the extended communication approach to information sys-
tems as illustrated above: THAT THE CUSTOMER COMPLAINS
or IS EXPECTED TO COMPLAIN UPON THE PREDICTED CONSTQUEN--
CES OF THE DISCOVERED PRICING ERROR BY THE CLERK. ia
either case this amounts again to assuming the truth

of the output or of the process and the behavior of the
customer, illustrating at the same time the well-known
relativity of output and process relative to the assumed
environment of the system, This also depicts the central
importance of the issue of value, in the above example
represented by the COMPLAINT OR EXPECTATION OF COMPLAZNT
BY A WELL IDENTIFIED CUSTOMER. It is apparent that such
issues could be disregarded or could be handled intuitive
ly in system design up to now, together with the issue
of QUALITY, because of the very limited scope of the
systems. The situation is well different e.g. in the
case of public data-banks serving complex values in the
sense of unknown, unidentified customers requiring un-
predictable processing of information. The situation is
further complicated in the case the customers are repre-
sented by decision-makers, public officials in agencies
which use—-up this information. This also obscures the
issue of the impact of customer complaint: how much at--
tention will be given to it by the decision maker(s)
supplying the information, e.g. in the context of choo-
sing a "fair" correction method ?

Error detection and correction becomes then extremely
complicated. It is therefore natural to make a desperate
attempt to extend the communication approach to the
third next higher level of sophistication, above the
just covered data-processing, general function Level.

We will then say that the best thing is to avoid the
need of error detection and correction by means of a
PREVENTION activity. The quality control system might
then be visualized as below, in terms of a further de-
velopment of fig. 4.1

Quality Control System

Error Preventio Error Detection) | Error Correction
1
|

Subsystem | Subsystem Subsystem |
a L —-
t
{

hat

The figure suggests that errors will be approached in
terms of the earlier quality control system of fig. 4.1
to the extent that they are not "caught" or prevented
by the prevention subsystem.

We may now illustrate this last prevention approach of
fig. 4.2 by immagining how a traditional system-designer
could intuitively attack the problem of designing such

a system for "total control of the quality of information"
in the context of a particular information system. The

following could be the result of an initial attempt of
breakdown:

Total (information)

System
! are Auton
| |
Quality Control !
(Prevention of !
Error Effects)
| Pia
|
Prevention Correction of
(Avoidance of Error Effects
Errors)
eae en Dees . Peer Tice aaa SET ORE Te
| at
| | | |
Cause-Fffect Error Correction Operation
Analysis Statistics of Errors with Error
i enna , Releases ees
| | |
System Human Repair & Human
Logic Factors Replacement Standby
of Subsystems
nts = ae
Detection Classification Correction
of Errors (Screening & Procedures

Evaluation)

1
Detection of Detection of

Error Source Propagated
Errors
Figure 4.3

Tentative breakdown of an advanced
information system with an own sub~
system for total control of the
quality of information.

AAs

Obviously many questions come up into the mind of who
looks and tries to work with a figure like fig. 4.3,

In particular one might ask whether it is possible to
associate with each subsystem a measure of performance
which is consistent with the goals of the overall sys-
tem, a basic requirement for the system thinking (see
for, instance Churchman, 1968a). What will be the impli-
cations for the above, of the fact that for example de-
tection of errors is the basis of error statistics; and
that repair & replacement of system is also an aspect
of the correction of errors ?

Since the whole reasénitig, However, is based intuitively
on the concept of ERROR, we might rightly ask ourselves
what is an error, how should it be defined or what is
its meaning. To say that its meaning depends on how we
apply the concept of error would lead us to circularity
in reasoning since we pose the question exactly in or-
der to be able to apply it. We may instead drop this
question for the moment and pick up another one by re-
marking that the introduction of the concept of PREVEN-
TION most explicitly forces the recognition of the need
of PREDICTION. In order to prevent we must do certain
things today which will prevent their predicted conse-
quences tomorrow. This may be seen as looking for causes,
as suggested by the cybernetic idea of going from error-
controlled to cause-controlled regulators: they imply
the need of prediction; and prediction is the fundamen-
tal problem of scientific method,

On a closer thought, however, it appears that DETECTION

as seen in the simpler model of figure 4,1 also required
prediction: we must know what to detect in order to set

up detecting procedures and in this sense the detecting

procedures are also prevention. .

We are then led to believe that"objective arbiters of
truth" of the communication approach to quality, cannot
anymore in the extended version just "see" the truth,
as the observer, auditor, manager etce., who look at the
input of a telephone or telegraph channel. The problem
of prediction in science is much more than to postulate
a general mathematical function or algorithm on the ba-
sis of so~called experience or sound judgement; and an
information system is much more than a telephone or te~
legraph system, The"objective arbiters of truth" must
now start to predict and in order to do that they must
seek assistance in the context of scientific method

and various "theories", And this makes indeed sense if,
as we expect, no ERRORS exist without prediction, since
errors are deviations between predicted and observed
values. Things will not become easier if, as we also
expect, observations imply predictions too since they
are based on assumptions and measurements made possible
through theories and respective predictions.

The above questions are at any rate enough for leaving
figure 4.3 and the attempts to extend the communication
approach to quality of information, and plunge instead
in scientific literature with a view towards "error",
"prediction", "“aceuracy", "precision", etc,

416

"REVIEW" IN ADMINISTRATIVE PROCESSES

In the context of a study of decision-making processes
in administration, H.A. Simon (1957) proposes a THEORY
of human choice or decision-making, The author defines
one function of REVIEW as DIAGNOSTS OF THE QUALITY OF
DECISIONS being made by subordinates. It is followed by
the function of MODIFICATION through influence on sub-
sequent decisions, the CORRECTION of incorrect decisions
that have already been made, and the enforcement of san-
ctions. Review is then, among other things, THE MEANS
WHEREBY THE ADMINISTRATIVE HIERARCHY LEARNS WHETHER DE-
CISIONS ARE BEING MADE CORRECTLY or incorrectly, and it
is a fundamental source of information with the help of
which, improvements can be introduced into the decision
making process, (Simon, 1957 - and.ed. p. 232)

To the extent that we regard information systems as a
formalization and possibly computerization of adminis-
trative decision-making, it appears that review then
includes our previously mentioned concepts of error
detection, correction, and prevention. As such it might
be relevant for our study.

A search for what Simon means by "correctness" does not,
however, assist our investigation. Upon making dis-
tinction between ethical and factual elements in a deci-
sion, and stating that criteria of correctness have no
meaning in relation to the purely valuational (ethical)
elements, he argues that "correctness" as applied to
factual propositions means objective, empirical TRUTH.
Furthermore, Simon argues that in the factual aspects of
decision-making, the administrator must be guided by the
criterion of efficiency. In order to determine in advan-
ce (PREDICT 7?) whether some statement is TRUE or false
one must use JUDGEMENT, not to be confused with the
ethical element above. Furthermore one must be careful
in order not to allow that CONFIDENCE IN THE CORRECTNESS
of judgements shall take the place of any SERIOUS ATTEMPL
TO EVALUATE THEM SYSTEMATICALLY ON THE BASTS OF SUBSE-
QUENT RESULTS. (p.50-53,197)

Simon does not develop his concepts of objectivity,
empirical truth, systematic evaluation, etc,, and this
is the reason we were not able to use his results in our
investigation about quality of information. Simon refers,
however,to "logical positivism"to which we shall return.
Using what apparently constitutes Simon's extension of
his "review" concept to performance programs in organi-
gations, A, Danielsson (1963) makes an interesting ana-
lysis of the relationships between programs, actions
(activities), and output (product) ,in the context of
organizational control. Danielsson suggests (p. 45) that
independently on whether programs consist of specifica-—
tions of actions or of QUALITY and quantity of output,
the RELATIONS BETWEEN ACTIONS AND OUTPUT MUST BE ASSUMED
"given", known within the company, either by management
or by the subordinates,if programs are to be utilized as
a basis for control. This suggests that the application
of this approach to quality is also "communication" orien
ted.

47

QUALITY AS VALUE AND EFFICIENCY

Modern administration and organization theory, as repre-
sented for example by Simon, seemingly attempts the
reduction of FACTUAL questions (related to the lower le-
vels of the means-ends hierarchy) to an evaluation

of their truth or falsity on the basis of the criterion
of EFFICIENCY.

On one hand, however, the idea of CORRECTNESS as applied
to final, end goals or values is often not considered to
be reducible to factual terms. Such premises must be
taken as "given" (by the highest levels of the hierarchy)
and they are said to have meaning only in terms of
"subjective" human values. Democratic institutions are
in this context mentioned, since the principal justifi-
cation for their existence is exactly that they are a
procedure for the validation of value judgements.

If, on the other hand, intermediate goals ate expressi-
ble in concrete terms so that the correctness of deci-
sions can be factually tested, NO ASSURANCE TS GIVEN ON
HOW THEY AFFECT THE HIGHER, FINAL, END GOALS OR VALUES.
This may be expressed by saying that no methods exist
for a scientific breakdown of the highest levels of the
means-ends hierarchy to concrete, factually testable
lower intermediate goals, relatable to the criterion of
efficiency. In this context it is explicitly declared
that the process of valuation lies outside the scope of
science.

Furthermore, it is recognized that little knowledge exists
on how decisions affect goals, even when they are expres-
sed in concrete terms ("production functions" of adminis-—
trative activities), and even assuming compromise

and proper weighing of multiple conflicting goals.

We see then that the "subjective", scientifically uncon-
trollable element enters at various important stages in
such administrative-organization theory: at the deter-
mination of concrete intermediate goals, and in the deci-

sion processes leading to such goals ~ to the extent that
the administrative production functions are not known
because of the fact that concrete, empirical investi-

gations have not yet been made of the way in which re-
sults change when the extra-administrative and adminis-
trative variables are altered. Furthermore we may have
a subjective element also in the establishment of what
is to be considered as abjective, concrete empirical
truth of the results of an investigation.

If we add what was said above to the previously mentio-
ned difficulties of making reviews, we conclude that the
reference to values and to efficiency in administrative
situations does not solve our problem of determining the
quality of the information used and produced by adminis-
trative decisions. In this sense, as suggested by one
statement of Emery in appendix Al, reference to value
does not dispense the need of the concept of ACCURACY.

4,2

48

TOWARDS ACCURACY AND PRECISION

Let us return for a moment to the case of an engineer

who retrieves from a technical data-bank the tensile
strength of a certain kind of steel to be used in the
construction of a bridge, As indicated by Eisenhart (1968)
and by Churchman (1961,p.335), we can safely say that if
the engineer gets his figure without an estimate of accu-
racy and precision, the figure will be WORTHLESS and
MEANINGLESS. More concretely this implies that the engi-
neer will not be able to use the steel in the design

and construction of the bridge.

Would anybody argue in defense of the use of the steel
anyway, oh the ground that no specification of quality
of this item of information on the steel is required
sihce such specification will be substituted by a mea-
sure of the improvement of bridge construction that

the information makes possible ? This argument may

be seen as an attempt to bypass the problem of accuracy
ahtl precision of information by referring its use %o
accrued value of the object system.

Such argument would raise serious objections, since even
supposing that the value of the bridge is measurable,

and that it is very great (for example in terms of net
savings due to higher traffic thruput), we cannot know
whether such net savings will be really net, in the sen-~
se that maybe the first nine bridges will collapse be-
fore the tenth proves to function as intended; this may
result, say, in a ten-fold increase of costs as compared
with the original estimate of net savings.

Appendix A4 indicates that in the context of mass-produc-
tion, it was until about year 1925-common to consi-
der "efficiency" in terms of output quantity in manufac-—
turing without due regard to scrap and rework costs.
Modern manufacturing knows better, as witnessed by de-
partments for quality assurance and quality engineering
in industrial firms. Do scientific researchers and for
instance designers of data-banks or information systems
know also better ? Do engineers always realize the im-
portance of quality of information ?

With regard to laboratory technicians, researchers, and
engineers, the papers of Branscomb, Eisenhart, and
Hallert, to name a few, are witnessing the fact that
many people today would be ready to, so-to-say, buiid
ten bridges in order to have one usable. Maybe the si-
tuation is far from being satisfactory even in such
'tsuccessful" fields as those of natural science. Does
such "success" in some sense imply that quality of in-
formation, after all, is not so important ? Churchman
(1961,p.342) suggests the answer to this apparent para-
dox: "The success of physical science may be largely at-
tributable to the amount of time and resources put into
the effort and not to the methods used; an analysis of
the methods might vastly reduce the need for such large
expenditures."

4.9

The concrete immlications of using bad methods in the
context of quality of information might be inefficient
use of resources in the form of duplication of research,
useless experiments caused by uncritical acceptance of
false results reported by previous researchers, meaning~
less talk about "random" errors cancelling out in the
course of the computations, creation of new undefined
concepts,like "confidence" and "usefulness" of data,
which add to the general confusion, ete.

We recognize that no argument is available against the
possibility that the same risks will be incurred in the
context of coming data-banks and information systems:
possible indiscriminate use of great masses of "data" or
"facts" stored in big, costly data-banks, which will be
used to "deduce" new "facts" to be in their turn the
input to decision makers and to other information sys-
tems.

Recall the engineer who retrieved the tensile strength
of the steel and is sophisticated enough to ask about
the accuracy and precision of the figure. The problem is
now to whom will he submit the question. Neither he

nor the vendor, nor the programmer - system designer

can go to the input of some channel to observe "“objecti-
vely" the true value which would dispense knowledge on
the accuracy and precision. Guidelines on "validity check
of input" in traditional EDP system handbooks would not
help because it is not a question of checking that the
field will be all-numeric and have a value range between
35 and 85,for example.

Let's leave the engineer and go to an administrative
decision maker who has just retrieved from a data-bank
the numbers of unemployed in two major cities, say res-
pectively 1,036 and 15,000, or the standard cost of two
sub-assemblies manufactured by a plant, say 37 and 700
dollars, or the amounts stolen once upon a time by two
ex-convicts, say 100 and 500,000 dollars. Why should the
decision-maker dispense specification of the quality of
such items of information ? He cannot be assumed to be
better served by his own "judgement" than the engineer
was; the figures cannot be said to be more "basic" or
“direct" or "raw" observations, they are not more "fac-
tual" or empirical, the observers who made the original
input cannot be said to have been more reliable or care-
ful; the consequences of his decision cannot be said to
be less important than the construction of a bridge or
the manufacture of a piece of machinery.

The feeling sometimes invades naive scientists and ad-
ministrators, that there has been some original INPUT
based on a very direct, “obvious" observation and that
later on the rest was taken care of by means of so-called
established statistical techniques or sound systems de-
sign. Perhaps these very same people like to think of

the sense apparatus of a human as being the analog to

the input device of a computer. Churchman (1968b,p.39)
poses then a very simple and puzzling question which we
believe is worth long meditation:

x

4,10

"The rational doubt about empiricism is based on the
very simple idea that the senses could te1l us false
things. What is the basis on which we believe that which
our senses tell us ? One analogy of the sense apparatus
of a human is the input device of a computer. But we all
know that a computer can accept falsity as readily as it
accepts truth. If our senses tell us that this is light
and not dark, how are we to know whether the input from
nature is not a complete falsity 7"

It is now important to note that the "review" which was
illustrated in an earlier section of this chapter appears
to be understood by its proponents as a review of the
so-called correctness of decisions and their measured
results, seen as specifications of actions and output.
We have not been able to find a discussion of the review
of INPUTS, Seen against the background of what has been
said in this section, we think that this is a remarkable
situation which requires clarification. We have investi-
gated this matter atid come to the conclusion that the
review of inputs is included in the review of actions,
since such actions include those which constitute OPE-
RATIONAL DEFINITIONS of the input variables in terms

of operations which must be performed in order to measu-
re them,

We have thus identified the "review" attitude towards
the problem of quality of information as subscribing to
the so-called schools of OPERATIONALISM and LOGICAL POSI-~
TIVISM. Following this matter further we have become con-
vinced that this view does not support our purpose of
specifying the quality of information, i.e. of finding

a guarantee against falsity of observation, or a guaran-
tee of value of the particular item of information.

A discussion of operationalism and logical positivism
would take us outside the scope of this paper, but the
interested reader may find for instance in Churchman
(1948), Ackoff (1962), and Northrop (1947) an illustra-
tion of the problems raised by operationalism, Such pro-
blems are mostly related to the ambiguity of the word
"operation", to the impossibility to find ultimately
simple operations,to that whether or not a specific set
of operations provides PERTINENT DATA depends upon what
kind of natural world we presuppose, and to that the
positivist finds meaning in a series of propositions the
confirmation of which cannot be a part of scientific
method.

If we dare to put it in more simple words, it appears
that what characterizes the positivist and operationalist
approach is their dependence upon UNCHALLENGEABLE ASSUMP-
TIONS. We think that such assumptions were clearly seen
to be dictated by higher management in the context of
administrative review, and e.g. by the observer in what
we called "the communication approach". The unchallenged
assumptions may correspond to the "non-systematically e-
valuated" management-" judgement" dictating the allocation
of deviations between predictions and observations to the
method of measurement (inputs), method of processing the
information (model),and method of measurement (output).
This amounts to state what is TRUE, i.e.not to be changed.

ALL

THE CONCEPT OF "JUDGEMENT"

For the sake of having a short summary over the previous

' gections of this chapter, let us recall our purpose of
developing in this chapter two aspects of the quality of
information,which can be used in the context of data-
banks and information systems. We are looking for a brcoa-
der meaning of quality than the offered by what we called
the "communication" approach, in order to take care of
the problems considered in the earlier chapters.

We started this chapter by reviewing the simplest
communication-quality. When attempting to extend it tn
order to cover the general-function "data-processing"
approach, we suggested several of the important assump-
tions - many kindof "given" things, like knowledge on
the. behavior of the customer ete, An attempt to bypass
such difficulties by means of error-prevention, required
a knowledge on the nature of error and introduced us ta
the concept of prediction, Since prediction is a funde-
mental problem of scientific method we recurred to

some scientific literature covering a theory of adminis-
trative behavior. It was seen that both administrative
review and the following of the criterion of efficiency
fall short of offering a guarantee of quality of a par-
ticular item of information.

Together with the earlier "communication" approach, re-

view and efficiency as a measure of correctness of in-

formation appeared related to the schools of operatio-
nalist and logical-positivist thought. We thought to ha-.
ve recognized some of the strong unchallengeable agsum-
ptions of such schools of thought,in the role given to
judgement by managers and observers in the context, for
example of

1. Validating the highest, final values or ends, br
judgement of the democratic character of the perti-
nent institutions,

2, Establishing through judgement the intermediate goals
corresponding to the highest values above.

3. Determining in advance the truth or falsity of a sta-
tement about the observable world to the extent that
no empirical results are available in the form of
production functions relating administrative activi
ties to results.

4. By means of implicit or explicit reference to opera.
tionalism and to logical positivism, determining in
part by judgement what is to be considered factual
result of empirical research, i.e. "empirical truth",

We feel that the above roles given to judgement are so
important that they justify a more detailed analysis o
it. The reader should recall that particularly in the

context of public data-banks, but also in private pre:
jects extending far into the future, final values may

not be identified, and much less related to intermediate
eoncrete goals. This obscures further the role of

ment in such cases, and consequently also its poss
contribution to the quality of information. Let's ge
started by illustrating judgement in the context cx

manufacturing and physics.

f

4,12

QUALITY AND JUDGEMENT IN MANUFACTURING
AND IN PHYSICS

In the same way as the processing of information is re-
garded by some people as the "production" of new infor-
mation, it is natural that in the search of methods for
controlling the quality of information we intuitively
think about the methods for controlling the quality of
manufactured products. The reader should not feel parti-
cularly distressed because of the confusion of concepts:
the confusion is well motivated indeed ! We ARE dealing
with paradoxical questions,

It appears that W.A. Shewhart is regarded as the "father"
of quality control in modern manufacturing. While his
"Economic Control of Quality of Manufactured Product"
written in 1931 is mostly dedicated to ways of expressing
quality of product, to the basis for specification of
quality control, and quality control in practice, the
SCIENTIFIC basis of his work is presented in a later
book: "Statistical Method from the Viewpoint of Quality
Control" of the year 1939. Appendix A4 is an account
(edited by us - out of a paper by S.R. _Littauer) of the
history of quality control, while in appendix A5 we have
edited some statements by Shewhart himself (1939).

The first thing to note is then that the"father" of one
of the most important activities in the most "down-to the
~-earth" contexts of the world, manufacturing, had to be-
come one of the most outstanding theorists of statisti-
cal method in its relation to scientific method, in order
to develop and apply new methods for quality control.

A review of the appendixes and of the referred literature
reveals that while borrowing from the "operational" schao-
ol and to logical positivism, the important accomplish-
ment of Shewhart was to develop scientific-statistical
CRITERIA OF ACCEPTANCE OF PHYSICAL HYPOTHESES which had
until that time been the JUDGEMENT OF THE INDIVIDUAL
ENGINEER OR SCIENTIST. In order to do this, Shewhart
recognized that manufacturing was to be regarded as a
scientific problem, and not as the tendency had been

up to that time ~- to regard it as a mathematical-arith-
metical "efficiency" problem in terms of counting units
of produced output and used input resources.

The question comes to our mind whether in the context

of EDP we are today in the same position as industry was
in the context of manufacturing before Shewhart: are we
only counting number of transactions processed per unit
of time, measuring “output" in the "production" of in-
formation, leaving the problem of acceptance to the
judgement of the individual decision-maker ?

In any case the lesson to be learned from manufacturing
is that one does not produce unless one produces UP TO
SPECIFICATIONS. If not, the subsequent test - if at all
possible - on the completed product may just prove to
be destructive for the product itself or for the produ-
cing company:bankruptey preventing further manufacture,

4.13

Thus, if somebody wants to consider the "production" of
information in analogy to physical production he must
also give specifications for such produced information:
it will indeed be used - as in the case of data-banks -
in further processing, in an analog way to the physical
piece of product which must meet certain specifications
in order to fit in some further mechanism, If the infor-
thation system just "produces", i.e; measures and proces-~
ses information, without regard to producing up to spe-
cifications, the information system itself and its spon-
sor may go into bahkruptcy.(Recall:the bridge collapses),

Thus, the use of a eccded observation, of the result of

a measurement, or of an intermediate computational re-
sult which is stored in a data bank is analog to the

use of a detail part stored in the stock of a manufac-
turing plant. The trouble is that when dealing with a
manufactured part we know that it "works" to the extent
that the customer buying the preduct in which it was
assembled does not complain; or to the extent that such
product in which it is used works in terms of verifiable
physical functions; or at least to the extent that it
satisfies operationally verifiable tolerance limits for
its physical dimensions ON THE BASIS OF A PHYSICAL AND
MATHEMATICAL THEORY that encompasses the specification
(e.g. the drawing), the measurement (its accuracy and
precision, and related quality evaluation in terms of
tolerance limits), and the physical manufacturing pro-
cess itself. Such comprehensiveness is also what allows
the relating of a customer complaint, or final product
disfunction to a "failure" of the particular detail part.

In pushing the physical~production, manufacturing analo-
ey so far as it can go, we would then like to obviate
the possible objections to present carelessness in eva-
luating quality of information by specifying for each
"kind" of information, each variable, some tolerance li-
mits which are to be verifiable and satisfied in order
to consider and accept a particular variable-value as a
"good" value.

We think that it is at this point that the paradoxical
aspects of the whole question of quality of information
undergo the most difficult scrutiny. For instance, will
the tolerance limits relate as they should to the values
of the information system (as opposed to the object, e.g.
physical system), and to the accuracy and precision of
the pertinent measurement process ? Are previous proces-
sings of information to be considered as the"measurement
process"? In such case how shall we operationalize such
process in order to obtain verifiable meaning for its
precision and accuracy ?

The above questions make it difficult to pursue the ques-
tion in terms of considering an information processing
system as analog to a physical production system. Infor-
mation is not "produced" but it 1s rather created by
means of MEASURMENTS embedded in theories on the vague
"reality" (which is not the particular and limited phy-
sical reality - corresponding to physics). The USE of the
measurements will also have to be made in terms of theory.

4.14

tn other words, QUALITY IN MANUFACTURING means

the attainement of somebody's values which are related
to manufacturing activities as described by the theory
of physics. It is the theory of physics that allows the
creation of information, by means of measurements, which
will be used in specifying and attaining quality, i.e.
indirectly the values.

The right analogy then appears to be that QUALITY IN
OTHER ACTIVITIES (not those which are today described

as manufacturing), such as those assisted by general
data-banks or information systems, means also the attai-
nement of people's values as related to those activities
as described by other pertinent theories. Such other
theories, as we wish that for example psychology, socio-
logy, and political science could, should be able to
describe and measure such activities - i.e. they should
allow specification and evaluation of attained quality.

In both cases, however, we have the basic notion of
measurement that was defined in the case of manufactu-
ring in terms of Shewhart's concepts of ACCURACY and
PRECISION. We are theri looking for a general meaning of
accuracy and precision, Such general meaning will be the
meaning of measurements leading to the general informa-
tion stored in general data-banks, "general" in the sen-
se that the use of such information is not known in ad-
vance, or if known it is not covered by any theory.

In this context it is interesting to note that knowledge
(empirical knowledge) of manufacturing production func-—
tions does not dispense the accuracy and precision of

the related measurements. WHY SHOULD THEN EMPIRICAL KNOW-
LEDGE OF "ADMINISTRATIVE PRODUCTION FUNCTIONS" DISPENSE
THE NEED OF ACCURACY AND PRECISION IN THE CREATION OF
PERTINENT INFORMATION ? Shewhart and Eisenhart make it
clear that accuracy and precision have a PREDICTIVE fun-
ection, as a guarantee of an item of measured information:
they attain this by CONCENTRATING ON THE MEASUREMENT PRO-
CESS which generates such kind of information, rather
than referring to the particular vaiue itself. This pre-
dictive, guaranteeing character of accuracy and precision
could lead us to believe that the function of these con-
cepts in administrative contexts is performed by JUDGE-
MENT (See Simon, 1957, p.51).

It is then important to note that Shewhart also requires
the concept of judgement in quality control of manufactu~
ring, and still does not dispense the need of accuracy
and precision. A review of Shewhart's work (1931, and
1939) leads us to the conclusion that THE FUNCTION OF
ACCURACY AND PRECISION IS TO ALLOW THE SYSTEMATIC EVA-~
LUATION OF JUDGEMENTS (in advance, of the truth or fal-
sity of statements about the obsesrvable world) ON THE
BASIS OF SUBSEQUENT RESULTS. This is indeed the same
thing Simon was looking for (1957,p.51) in order to pre-
vent unwarranted confidence in the correctness of jud-
gements, while recognizing that the process by which

4.15

It is clear that we might be wrong in our concluding that
accuracy and precision have the purpose of evaluating
judgements as above understood by Simon. We might then
assign them to the function of determining empirically
the factual content, the objective truth of administra-
tive production functions.

In any case, the concepts of accuracy and precision raise
difficult problems to the operationalist and logical-~po-
sitivist approach to administrative decision-making,

This approach makes no reference to the accuracy and pre-
cision of the measurement protesses leading to informa-
tion to be stored and used in the context of data-banks
and management information systems. However, as we illus-
trate in appendix A4, A5,A6 the work of Shewhart, as
well as the paper of Eisenhart on the concepts in physics
show the following:

1. TRUTH of reported values is a function of the accura-
ey and precision of the measurement process. The re-
quired accuracy and precision depends on the uses and
VALUE of the information.

2. In the context of ECONOMIC values, JUDGEMENT has a

role,for example

2a. For establishing BCONOMIC specifications in terms
of tolerance limits which must be based on ECONO-
MICALLY assignable causes of variation.

2b. For making an ECONOMIC choice among many different
practically verifiable criteria (criteria with
operational meaning) of attainement of specifica-
tions, i.e. criteria of TRUTH and of ERROR.

2c. In evaluating the QUALITATIVE, as opposed to the
quantitative aspects of measurement. (Not specific
for economic values).

3. ACCURACY AND PRECISION may be seen as a measure of
the DEGREE OF TRUTH since the OBJECTIVITY of a guali-
ty characteristic exists only in the CONSISTENCY be-
tween the indefinitely large number of potentially
infinite sequences constituting the numerical aspects
of several different methods of measurement. Precision
is a measure of disagreement or consistency for ONE
method while accuracy encompasses: disagreewenk across
several different METHODS, or between them and a me-
thod chosen as a STANDARD, ~

Churchman (1961, p.196) refers to several of the above
ideas in the following way: "... the assignment of a
length to an object enables one to predict how the ob-
ject would compare with other objects in various environ-
ments.. What number is assigned is determined by the eco-
nomic conditions entailed in any construction of stan-
dards.. These economic conditions depend on the actual
utilization that is made of information about lengths,
namely, certain kinds of comparisons."

In summary, FACTS appear to be a matter of degree intima-
tely related to VALUES. The problems that this raises for
the operationalist approach to quality are expressed by
Shewhart's analysis of the relations between EVIDENCE,
BELIEF, PREDICTION, KNOWLEDGE, and VALIDITY OF JUDGEMENT,

4.2.3

4.2.3.2

ML N\ (>
haces

416

THE ROLE OF PHYSICS IN DESCRIBING CONTROLLED SYSTEMS

In an attempt to expand the scope of our analysis in
order to evaluate the more complex aspects of quality
of information we met the concept of JUDGEMENT in the
context of the operationalist and logical-positivist
approach to administrative decision-making. In the pre-
vious section we searched for the role given to judge-
ment in the best known, most concrete field of physical
manufacturing, with the purpose of better understanding
the eventual possibility of using it as an indicator of
quality of information. We found that judgement did not
dispense, but rather completed the concepts of ACCURACY
and PRECISION of measurement which were met for the first
time in the referenced literature and in appendixes A4
to AG.

The most disturbing implication for the logic-positivist
approach was that even in the context of the most conere-
te production functions of industrial manufacturing, as
well as in physical research, FACT and TRUTH appeared to
be a matter of degree and were intimately related to
VALUES and JUDGEMENT. We shall now explore how this in-
sight may be illustrated in connection with some common
concretizations of information problems. We feel that
the illustration will assist in appreciating later our
attempt to generalize the concept of quality of informa-
tion.

FIGURES ILLUSTRATING ACCURACY AND PRECISION

A relatively common and appreciated method of illustra-
ting the meaning of accuracy and precision, as well as
several concept of related errors is by means of the
following figure

P ee yo
Hee,
Pe

\

Figure 4.4
Target patterns of shots fired by two riflemen.
The left pattern exhibits low precision and high
accuracy with large random errors, while the
right pattern exhibits low accuracy and high
precision with large bias (systematic error).
(Adapted from A.Chapanis, 1951)

4.2.3.2

4,17

A.Chapanis uses the similar figures in a paper dedicated
to the "Theory and Methods for Analyzing Error in Man-
Machine Systems" (1951). He mentions "naval information
systems" but his concern more closely specified appears
to be the accuracy, in some sense, of naval radar equip-
ment, The idea of information comes from the statement
of a research program including the objective of " The
evaluation of naval radar equipment in terms of the
ACCURACY, KIND, and AMOUNT of information an operator
can extract from it", and from seing radar systems as
dealing "with a rather nebulous product - information",

Since, most SETS OF ERRORS, in both physical and biolo-~
gical phenomena, appear to be normally distributed,
Chapanis suggests that the statistician may apply the
standard statistical methods for the analysis of varian-
ce.

The figures have also been used in illustrating human
variability, ahd related nature, frequency, and effects
of human ERRORS on defects, failures, and accidents

in the context of industrial product manufacturing.

It appears to us that great care must be taken in apply-
ing the thinking ahove outside the limited field of

purely physical systems. The application of such thinking
to the ahalysis of human error already raises important
questions, and many more may appear in the context of
data~banks and information systems for administrative con-
trol. The most important unwarranted assumption is the
self-evident knowledge of the OBJECTIVE or TRUE VALUE,
which allows for measurment of deviations leading to

also self-evident concepts of error.

ILLUSTRATING CONTROL SYSTEMS

In the context of decision-making, the concept of deci-
sion and control may be illustrated in the following way

Se,
2 a
~ a“
Patt
Trajectory
a
z eg, a ee
Transmit a ae
a si
Receive

: Computation ©)
Center

-<[}
5
=

Zs L ates)
Tracking Stations

Human
operators © ey

Figure 4.5

4.18

The figure is taken from A.Kaufman (1968) who also sug-

gests the following analogies for the concepts numbered
1 to 5:

T

IN GENERAL

VEHICLE BUSINESS FIRM
1.Car and its ‘1. Objectives, 1.0bject and trajec-
driver. \ tory. |
2.Centers of con- |2.Centers of ace- |2,Controls,.
trol and infor- ountancy, sta-
mation inside tistics, and |
and outside the control.
car,at disposal
of driver.
3.Driver's brain. {3.Management 3.Calculation.
computer.
4.Centers of per- 4. Executive 4 Methods of execu-
ception and con+ levels. tion and receptions
trol of the dri ‘
ver. i ‘ i £
5.Free will. [5.Responsibility (5.Command. H
| for decisions. H

4.2.3.3

If we have captured the intent of the illustration,
Kaufman wants to convey a "feeling" about the meahing
of decision and control. However it is clear that the
analogy fails in several aspects, the most important
again being related to the idea of OBJECTIVES. The ana-
logies have the advantage of raising the important
question of who is the driver, who is in command, whose
objectives (ar at all definable) are being served, and
what is the role of free will and responsibility for de-
cisions.

By ignoring such issues one begs the question of the es-
tablishment and evaluation of "facts", and it may be said
that it is equivalent to bypassing all the most important
and difficuit aspects in the development and operation
of information systems for administrative control.

Such aspects are considered for example by Churchman

in the book "The systems approach" (1968a).

THE SYNTHESIS OF RELIABLE ORGANISMS FROM
UNRELIABLE COMPONENTS

Five lectures given by J.von Neumann in 1952 were publi-
shed in 1956 under the title of "Probabilistic Logics
and the synthesis of reliable organisms from unreliable
components" (See "Automata Studies" edited by C.E.Shan-
non and J.Mc Carthy, 1956, p.43-~98),

In spite of Von Neumann himself stressing that the sub-
ject-matter is the ROLE OF ERROR IN LOGICS, OR IN THE
PHYSICAL IMPLEMENTATION OF LOGICS, it has been recently
suggested (G.Montelius et al., 1970) that the apnroach
is generally relevant to the study of errors and the ef-
fect of errors in information systems for administrative
control. We have not found support for this suggestion.
Von Neumann was actually concentrated on the logical-
physical aspects of computation, especially as related
to the mathematical ones. In another paper, however, he

4.19

together with H.H, Goldstine (1947) present a much more
complex understanding of what they call the "sources of
errots in a computation",

As they state it "When a problem in pure or in applied
mathematics is "solved" by numerical computation, errors,
that is, deviations of the numerical "solution" obtained
from the true, rigorous one, are unavoidable. Such a
"solution" is therefore meaningless, unless there is an
estimate of the total error in the ahove sense. "(p.1023)

In an attempt to enumerate and classify the sources of
errors they present the following:

1. The model or mathematical formulation of the problem,
representing only a (more or less explicit) theory of
some phase of reality:errors due to theory

2. Parameters in the model above, the values of which
have to be derived directly or indirectly (that is,
through other theories or calculations) from observa-
tions! observation errors

3. The approximations of the mathematical statement as in
1. above; iH replacing it by elementary arithmetical
processes which the computer cah handle directly, and
by explicit definitions, which correspond to a finite,
constructive procedure that resolves itself into a li-
near sequence of steps.approximation~truncation errors.

4, The "hardware" ~ the computing procedure or device
performing the operations which are its "elementary"
operations as specified by the results of the numerical
analysis in point 3. above: "random noise" of the com-
puting instrument, that is, errors and imverfections
inherent in any PHYSICAL, engineering embodiment of
a mathematical principle,

In the spirit of the earlier figures 4.1 to 4.3 one could
then essay to"illustrate" the error-control program for
an information system by means of the following figure:

Quality Control System

1 i :
Control of iControl of lecnaeet-s of | Control
model. observation approxim./ of phy-
errors errors | truncation sical
H | t errors eel i erro rs iu

Figure 4.6
A tentative illustration of Von Neumann-Goldstine's
approach to the sources of errors in a computation

Von Neumann and Goldstine's work dealt mostly with errors
originating under point 3., while the earlier mentioned.
work by Von Neumann alone dealt with those related to
point 4., together with errors of logic which may be seen
as a link to the other mentioned issues under inquiry.
The figure, however, by itself raises well mativated
doubts about the soundness of a partial approach to the

4.2.3.4

4,20

“information errors", as well as about the soundness of
an approach along the ideas illustrated in figures 4,1
to 4.3, prior to having obtained a deep scientific un-
derstanding of the nature of information, of quality,
and of error. Furthermore the figure sets us in guard
against some naive thinking in the context of human fac-
tors in information systems, as represented for example
by the statement that increased "reliability", and
"accuracy" of information systems may be obtained by
eliminating the human "link", putting more of the act

of observation into the computer, avoiding duplicate in-
puts,etc.

What Von Neumann and Goldstine do not discuss in depth
is the meaning of the "true, rigorous! solution, and
particularly the meaning of logic and errors in logic,
The analysis made by Churctimian in several of his works,
howevet, (see for example 1968b, p.41) shows that the
analysis of physical and logic errors, as advanced by
Von Neumann (1956) leaves untouched the most important
questions about truth, error, and quality of information,
The importance of the Von Neumann-Goldstine approach
in their work of 1947 is for our purposes the insight
that "facts", especially after some computation, but
even if derived from what they call a "direct" observa-
tion must be evaluated for errors.

THE "UNDERLYING PHYSICAL PROCESSES", AND THE
MULTILEVEL STRUCTURE OF ORGANIZATIONS

The most common way to visualize organizations is today
in terms of multilevel hierarchies with an underlying
system of PHYSICAL processes which may be described by
the laws of physics and chemistry (See for example J.C,
Emery, 1969, p.36; M.D. Mesarovic, 1970). The higher
levels consist of programmed and non-programmed decision
processes which may be described by signals and informa-
tion in terms of"pure" symbol manipulation and data-pro-
cessing (in some sense), or - at the highest Levels ~-
for example in economic terms.

The development of a "theory" for the control of organi-
zations on the above basis has apparently required the
creation of new words like STRATA for levels of desc1lp-
tion, LAYERS for levels of so-called decisiom vomploalty,
and ECHELONS for levels of organizational hicvarchy. The
analysis for control of organizations seems Later to re-
quire the study of relations among these different types
of levels,

For our analysis, what is extremely interesting in the
above approach is that it appears in some sense wholly
grounded on the "factuality" of the underlying physical
processes. It is from there that "facts" or "events"

are described or observed in terms of some sort of "co-
ding scheme" as a means of entering into the information
system (INPUTS) . Input data on events and performance,
and information feedback flow upwards in the hierarchy,
while coordination and control in terms of constraining
decisions are transmitted downwards.

oo

x 4.21
va \
\
/ \
/ \
/ \
/ oe
Yo jes S -
j — [Reodeaon| —\,_Droteton-malsing
/ Unit v0
/ Hage ct some echelons
\
\
\
\
Decision *

ieee as

Decision | Decision | Deckeion|
Unit | Unit Unit | \

inpus i "Natural" Process Output —

Resources

e.€@. physical, chemical, biological Preddet

Figure 4.7
One of the multilevel descriptions
of the overall control problem.
(Adapted from Mesarovic, 1970)

It appears to be the above concept of relation between
information and underlying physical processes, that ori-
ginates the understanding that the "facts" are the infor-
mation inputs to the information system, in terms of
coded observed events in, say, physical processes.

The idea apparently recurs in case of distinctions which
sometimes are made between physical and information pro-
cesses, or between material system and information sys-
tem. This is the conceptual framework which apparently
explains,for example Emery's view of data-collection as
consisting of sensing and recording of data where

" A human senses information primarily through sight,

as in the reading of a meter or observing boxcar serial
numbers." (1969, p.38). This may also be the background
of Blumenthal'ts statement, as seen in appendix Al, that
" A datum is an uninterpreted raw statement of fact."
(1969, p.30). Furthermore J.Porrester when discussing
inputs to decision functions apparently assumes a simi-
lar framework since he refers to "the distinction betwe-
en the TRUE value of a variable and the value of informa-
tion ABOUT the variable..." (1961, p.103).

The same approach would be implicit in the following
first tentative conceptualization of inventory difference.

1A 4,22

(True)
quantity

Reported
delivery
and req's
from stoc

6A

True,
(correct-,

ly comput
balance

Computed
differen~
ce :

ener |

;

Figure 4.8
A tentative conceptualizalLion of inventory
difference, as relating to situation des-
eribed in appendix A3, using the concept
of "true vatnues" as opposed to reported
(that is observed and coded) and computed
values which may be in "error",

The diagram is drawn according to the method
of documentation by M.Lundeberg for informa-
tion-analysis according to B.Langefors.
(See - M. Lundeberg, 1970, p.180)

4.23

The reader may recognize the relation between the approach
of figure 4,8, and figure 4.7, We have the input/output

of the physical process in terms of incoming (deliveries)
and outgoing (result of stock requisitions) parts from
stock. The data, facts on this process are the reports,
coded observations which are input to the information sys-
tem but in our conceptualization they are distinct from
the "true" values in order to account for observation and
other errors (see appendix A3). The figure 1s simplified:
for instance the information set 1A stands for both

true quantity in stock and for truly delivered quantities,
and several other relations are not shown,

Most information processes, 2, 3 , 4, 5 and 6 are not
specified. Observe that process 2 generating the informa-
tion set 2A (which could be obtained by direct interviewing
of stock clerk upon completed search of the part in stock)
may depend e.g. upon information on time which is availa-
ble for search. The part may be urgently needed anc if not
found within one hour it might be better to request a new
one from the vendor "across the street", Process 2 is ob-
viously also depending in a more traditional way on infor-
mation about the stock location, inventory bin where such
parts are expected to be found. Such information itself
may be obtained from the information system, and may be
wrong.

Information set 5A may be wrong according to the concept of
error advanced by Von Neumann-Goldstine, because of logic,
physical, model or numeric errors,

What we called "true found difference"7A, is less true
than another information set which is not shown in the fi-
gure but which would correspond to the difference between
1A (instead of 2A) and 5A or 6A. Observe that our "true
found difference"7A may itself be wrong because of possi-
bly wrong computation of stock balance 5A,

What is the ERROR ? Will the correction of 54 (and there-
fore implicitly our conception of which is the TRUE value)
be based on 6A, 1A, 7A or 8A? What is the role of a control
of the difference by a rotating inventory clerk, and how
will it be incorporated in the analysis ? It is interes-
ting to question how "statistical methods" would help the
solution of the problem,

We think that the above illustrates the vagueness and pro-
blems of the TRUE VALUE, even in the most simple, self-
evident physical reality, the most simple logic and arith-
metic related to the stock of a manufacturing plant.

We see then that the underlying physical process, as sug-
gested. by figure 4.7, for all PRACTICAL purposes (and the-
refore theoretical as weil) does not generate facts but ra-
ther only information with a certain error content.

We can now examine more closely figure 4.7 and ask oursel-
ves if the "natural" processes, physical, chemical, and
biological might be completed with psychological, social,
and economic. Where, how, and why goes the limit ?

4.2.4

SCIENTIFIC METHOD

Does the scientific literature help in unraveling the many
questions raised on the role of physics in describing con-
trol and controlled systems 7 Does such a role really dis-
pense from a meaningful discussion on the truth of the in-
puts to an information system, or on the truth of informa-
tion stored in a data~bank?

We have found that some literature apparently touches on
the very same problems that we raised. For example Ackoff
(1962, p-170) in the context of searching for a definition
of information, and a general meaning for PRECEDENCE, and
PRODUCTION, states: "It may be very simple to determine
whether an object is red where the consequences of error
are trivial. But if the observer's life depends on the
color determination, the problem becomes as complicated as
possible,"

Churchman (1959, p.90) states: "In effect, the "cost" of
adjusting data rises as more precision is attained, just

as the cost of the absence of precision goes up as we
attempt to find "simpler" data. Experience has shown that
it is possible to be naive with respect to precision in an
attempt to be simple in procedures, All of the supposedly
"simple" instances...-a report of a witness; of a labora-
tory technician, of a stock clerk - are not simple at all
if the decision on which they are based has any importance."

Will information stored in data-banks be used for decisions
of "any importance"? If so, how to reconcile the talk

seen about facts to the problems of measurement ?

As a further illustration let us consider the measurement
of birth-dates of citizens to be stored in public data-
banks. The measurement of birth dates appears to be so sim-
ple to the point of sometimes being declared that they are
just facts, and that as discrete (as opposed to continuous)
variables, they are just right or wrong and that there is
no meaning in talking about the accuracy of such measure-
ment. We think,however, that the intent of Ackoff's and
Churchman's statements above can be concretized in

part by immagining that legal and economic advantages are
instituted for those being born on one date rather than
another, What if the children are usually born at home ra-
ther than at a public hospital ? Will the date be made de-
pendent upon the minutes, seconds, and tenths of seconds

of "birth" ? How would one reach agreement on which event
would then correspond to "birth" ? How would one control
the process of measurement of time ? How would one adjust
birth dates already stored in the data-bank, related to
people who are retro-actively affected by such institution
of legal-economic advantages ?

In an analog way, counting of number of parts in stock, is
simple because we can ask the observer to repeat the count
one, two, ten times and everybody agrees that after, say,
the second count the counts converge towards the "true"
value. But what if deliveries to and from stock are made

425

while the counts are proceeding 7 Let's hire two, three,
ten observers depending on the frequency of deliveries,
and the space available for their simultaneous observa-
tions. But we cannot do for all the 10,000 different part
numbers in the stock of a manufacturing plant, at the sa-
me time, in any case we could not afford that, Then we
have to draw samples and make inferences from the sample.
It may appear similar to measurements of continuous varia~
bles in physics, where each determination or reported va-
lue is idealized as an individual of a population to which
we try to apply statistical theory.

We would however deal with a very illdefined population in
deed if the observers had own interests and judgements,
and if they were observing unwanted attributes of people
rather than of parts in stock ! Then we reach outside of
the realm of physics and of statistical theoryi The same
may be true if starving observers were counting units of
food in stock upon which the life of other starving plant
employees was depending upon, Even if the example is ex-
treme it is easy to immagine that the issue is a matter of
degree.

The unwarranted supremacy of physics in the description of

the control problem, information systems etc., has been
discussed in detail by several authors. Ackoff (1964,p.53)
summarizes in the most impressive way the criticism

against the school of logical positivism as supporter of

the unwarranted role of physics as expressed in much con-
temporary thinking about information systems, artificial

intelligence, etc. He concludes that

1. Scientific concepts are NOT reducible to a set of ulti-
mately irreducible concepts provided by direct observa-
tion or as undefined concepts of a formal system.

2. IT I8 NOT possible to synthetize all other meaningful
concepts in chemistry, biology, psychology and social
science, through manipulation of "physical thing pre-
dicates" i.e. physical properties of things derivable
from physical attributes.

3. Consequently, physics is NOT the one only discipline
that is conceptually independent of other empirical
disciplines, and it CANNOT assume a position at the
head of a hierarchy of scientific disciplines such as
chemistry, biology, psychology, and social science, in
that order.

4, In general, it is not possible to pose the problem of
unifying science by interrelating disciplinary output
either in the form of FACTS or CONCEPTS (i.e. logical
positivism), or laws or theories (i.e. so-calied gene-
ral systems theory).

Then, it appears that it was the Logical positivist approa-
ch that conditioned the earlier presented ways of illustra-
ting accuracy and precision, control, reliability, etc.

4.26

In particular this may explain how it could happen that
VALUES and JUDGHMENT could disappear in the context of
FACTS and TRUTH, allowing the relatively common statement
that "the problems do not lie in the computer and data-
bank, since they only store FACTS; the problemslie with
the people who are going to use the facts or be affected
by them",

Ackoff's discussion also gives a hint oh why many of us

may have felt perplexed when trying to apply the idea of
the "underlying physical processes" to the design of an
information system for a purely administrative organiza-
tion, for the limited scope of an engineering department,
for a hospital. It might have been difficult indeed to

find the "basic facts" if the criticism against logical po-
sitivism is well motivated.

Kaplan (1964, p.254) writes: "...the distinction between
facts and values cannot be drawn so sharply and so simply
as is commonly supposed. Any conclusion as to what the
facts are in a given case is the outcome of a process in
which certain valuations also play an essential role."

Northrop (1947, p.36) writes: "It cannot be too strongly
emphasized that if one wants pure fact, apart from all
theory, then one must keep completely silent, never repor-
ting, either verbally or in writing one's observations,,."
And later (p.177):"It is usual for the popular mind and
eccasional uncritical, scientific minds to assert that
science is concerned only with fact in the sense of what
can be observed and that it has nothing to do with theory.
»e-ITf it is pure fact, apart from all theory, which one
wants, then it is not to science but to the arts when they
function in and for themselves that one must go." Further-
more Northrop offers an extremely interesting discussion
‘of "facts" and "truth of inputs" in discussing operationa-
lism (p.125-128).

Morgenstern (1963, p.133, 88) distinguishes between"data"
and "information" that is SCIENTIFIC FACT , or measuremen-—
t . He writes: "The data by themselves tell us no story
whatsoever, neither a true nor a false one, They are silent"
And "..,data as such tell no story, or they tell many dif-
ferent and conflictning stories simultaneously; either con-
dition is equivalent to the Lack of a theory" The author
illustrates his point from the following figure, .slightly
adapted by us. He distinguishes between OBSERVATIONS that
are deliberately designed, and other DATA that are merely
obtained:

SCIENTIFIC INFORMATION is regarded as made up of
1. QUANTITATIVE OBSERVATION, i.e. body of data consisting
of gathered (numerical) statistics, but encompassed by

theory
2. DESCRIPTION, i.e. other data, such as historical events
or (now) non—-measurable data, e.g. “expectations” ~—

but which are alsa ecompassed by theory,

427

eee
a

* CC is the theory
/pased partly on A and
/3: as well as on deduc-

/ tively obtained facts
(perhaps not accessible to
direct observations)

fs \ac and AB is Sea en- va c \
/ Ng pipe infgrmay Pa <
/ Se tion / a
iA is the body of en ameee Bis data such as \
data consisting of | historical events or |}
gathered (numerical) ! (mow) non-measurable i
statistics i data, e.g."expecta- |
i tions" }
\ ae. iy
. \ 4
=e a @
s a SS, *
\ Ne a“ ™~

Figure 4.9
Adapted from Morgenstern and illustrating
the author's understanding of the truth
content of facts in economics;
Intersection AC is QUANTITATIVE OBSERVATION,
Intersection BC is DESCRIPTION,
Intersection AC and BC is SCYENTLFEIC ITNFORMATLON,
Most economic quantitative (statistical) data are
of the class A minus ¢

We may now pause for a moment. If "facts" are not self-
evident and given how does this reflect in the context
of data-banks and information systems, outside the liwi.—
ted scope of our simple case of inventory differences ?
Churchman, who in almost all his referenced work, has
been explaining the relativity of facts to values and
theory, gives what we feel is a pertinent example,

(1968b, p.153):

"A manager may ask: Given these sales Last year, what
will the sales be next year ? Another and far more in-
teresting question is: To what degree is this a sale ? ...
To learn that a customer is sold in degrees of conviction
is to Learn why he appears to be someone we sold to last
year... To ask why a customer aapears to be sold is also
the start of an inquiry in which forecasts of next year's
sales based on this year's sales are irrelevant. It is to

4.28

understand that recording a sale is a delicate decision.
To record some transaction as a sale when the customer is
truly dissatisfied, or truly erratic, or truly dead, is
to make a foolish decision."

We can, after this self-explanatory citation continue

by asking ourselves what are the values, or the theory
which guarantees the factuality of the transactions on
events or facts, that are stored for example in a public
data-bank. Will it be physics ? Or mathematics and logic ?
Or will it be in some sense a "THEORY OF DATA-PROCESSING",
or "THEORY OF INFORMATION SYSTEMS" ? Or will the problem
in some sense be taken care of by some governmental agen-
cy for "DATA MANAGEMENT" ?

Thus, we come into the deep but extremely important waters
of VERIFIABILITY, TESTS OF VALIDITY, and the like, which
we had left after illustrating quality and judgement in
manufacturing and physics in the previous section, We. em-
barked into analyzing the role of physics in describing
the control problem, since it appeared that no values or
judgements were required there in order to evaluate the
facts about the underlying physical processes. We see now
that we are back there. What does the scientific Literatu-
re suggest for testing the validity of information ?

Morgenstern, who appears to be quite statistically orien-
ted in his approach, is however one of the few who has
seriously considered this problem in the broad and impor-
tant context of economics. For instance in CHECKING THE
ACCURACY OF production statistics a method which is well
suited is the following: "If two or more processes are
known to be interrelated in a rigid manner, say technolo-
@ically, and the data for one process are trustworthy, then
the measurements of those other processes may be estima-
ted on the basis of this interrelationship."(1963,p. 52)
Furthermore, in discussing the TNTSRNAL CONSISTYNCY of
statistical data and other qualitative information, espe-
cially if AGGREGATES are formed, the author recommends the
establishment of CONSTSTENCY TESTS, the safest consisten-
cies being always TECHNOLOGICAL,He notes, however, that
whatever "consistency" is tested, IT CAN ONLY BE ESTABLI-
SHED ON THE BASIS OF SOME MODEL. (1963, p.132)

We feel,then » that there is a disadvantage in limiting
us to technological consistencies in testing validity or
truth in the context of information systems: it might be
like allowing the logical positivists returning through

the back-door. It limits what CAN be verified and therefo-
re what can be changed. If a biologist observes some un-
explainable phenomenon through a microscope, he may easi-~
ly verify through the theory of physics whether the instru-
ment is well adjusted, but this does not legitimate the

use of the microscope for that particular observation.

43

Y

V

os.

4,29

QUALITY AND JUDGEMENT IN DATA BANKS
AND IN INFORMATION SYSTEMS

Our search for a guarantee of quality of information
in information systems and data-banks took us to the
concept of JUDGEMENT. It was seen, however,that judge-
ment in the control of physical manufacturing processes
and of physical research had to be complemented by

the specification of ACCURACY and PRECTSION. The split
between judgement on one side, and accuracy and preci-
sion on the other was seen to be not justified: first
because physical processes require judgement for esta-
blishment of their factuality, secondly because physi-
cal processes cannot be separated from any other pro-
cesses by the criterion of factuality or truth.

Both reasons may be two aspects of the basic nature of
scientific method, that is our way of "knowing",

In appendixes AH to A6 we saw that accuracy and preci-
sion could be seen as a formalization of some of the
valuational aspects of judgement: for example economic
values in manufacturing and potential uses of results
in physical research. Appendix A7 is our edited inter-
pretation of what is written in some scientific litera-
ture on the concepts of accuracy and precision seen as
two relevant aspects of the quality of scientific in-
formation, in general. The findings in such literature
confirm that accuracy and precision can be seen as

a partial formalization of judgement. Such partial for-
malization aims at GUARANTEEING IN TERMS OF A MEASURE
FULURE ATTAINMENT OF GOALS WHICH CANNOT BE SPECIFIED
IN DETAIL.

Appendix A7 and the referenced literature furthermore
suggests that such guarantee of value without reference
to detailed goals is made possible BY RELATING DISAGREE-
MENT TO THE OSJECTS AND TO THE HUMANS WHO MAY BE DIPFFE-
RENTLY AFFECTED BY FUTURE USE OF THE INFORMATION,

For detailed alternative definitions of accuracy and
precision the reader is referred to the appendixes A5
to A7. We will return to the problem of defining them,
later in this chapter. For the moment it will suffice
to emphasize the fundamental role of ACCURACY as an in-

PRECISION appears in some sense to be an indicator of
repeatability in the course of time.

We conclude then that quality and judgement in the ge-
neral context of science may be reduced to formal terms
and quantified in the form of accuracy and precision.

oO

4.30

If what was said refers to SCIENCE, what is its relation-~
ship to our original problem of data~banks and informa~~
tion systems ? Since they are designed and used directly
or indirectly for the purpose of managing or doing, it

is relevant to observe that Churchman shows how scien-

ce is a kind of management, and management is a kind of
science. (1968b, p.29,36,43,144) This implies that was

is said about quality of scientific information should

be relevant also for the quality of management informa-
tion,

Another way to arrive at the same conclusion is to re-
fer to the earlier conclusion that every "fact" in
terms of a recorded item of information, implies a
theory. Conseqtiently, since theory is a concept of
science, if we record and store or use these facts,

we are at least implicitly assuming a scientific theo-~
by. And such theory will have to correspond to the
formal processing of information by the information sys-
tem for to the so-called symbol-manipulating, fact-de-
ducting systems) and to the informal use of such infor-
mation by people. This amounts to say that data-banks
and information systems may be regarded as theories,

or formal statements of beliefs in predictions aimed

at certain goals.

Such implicit "theory" will obviously be an integration,
in some sense, of several kinds of disciplinary theories
(physics, geometry, arithmetics, psychology, economics,
etc.), since human knowledge is organized along such
"information subsystems".

The important point to note, then, is that to the ex-
tent that we look at information systems as if they
were communication or storage-and-retrieval systems,
not only will the CODING ASPECTS be purely physical-~
technological ones, but the whole system will be desi-
gned and evaluated in physical terms. A case of purely
physical-economic design is renorted, for example, by
Churchman, as related to a case study. (1968a, p.126)

What we mean, then is that the technological interpre-
tation of computer programs misses the point that such
programs when applied to e.g. business control, rather
than to control of purely physical processes are in-
deed integrating natural science models with much less
established models and "ad hoc" hunches on psyohotogi-
eal and social behavior. In the field of physical sci-
ences, where there has been a successful theory-buil-
ding, most"errors" may be classified and assigned to
the class of OSSWRVATION ERRORS, If a machine does not
"work", we are more inclined to think in a "human error"
in the operation or assembly of the machine, than to
question the laws of physics according to which the ma-
chine was designed.

Not so with "errors" in the context of information sys-~
tems, An observation which does not "fit", that is,has
been "wrongly" coded into such an integrating program

4,31

should not be "a priori" rejected but rather regarded
as an ELEMENT IN THE TEST of such integrated model or
tentative "theory" about the object system. In the same
way, an observation should not be "a priori" accepted
just because it happens to be made by an authoritative
observer with "good judgement",

The logic and the economy of the integrated model, as
well as for example the physics of the hardware can be
perfect and still the model may at the end fail because
the psychology in it was very poor; one can name this
as an OBSENVATION EkROR, but it could rather be named
as a PSYCHOLOGICAL MODEL-ERROR. This is another way of
concluding that it is not motivated to see the problem
of misusing information stored in data-banks,in terms
USE of the information upon retrieval from the bank,
under the pretext that there is no alternative to the
"simple" storing of "pure facts". Concretizations ta
this point were seen earlier in this paper, in the con-
text,for example,of CODING and of the meaning of FACTS,
and will not be repeated here.

Anything, however, can happen to the extent that we have
no TESTS for soiving the ahove problems. We have already
touched upon such tests at the end of the previous sec-
tion when we referred to Morgenstern's recommendation

of internal consistency tests based, if possible, on
technological relations which are the safest ones,

Most tests presently performed in administrative EDP
applications are extremely naive: typical programmed
checks are e.g. record counts, file totals (amounts or
hash-totals), limit checks, cross-footing balance checks,
zero balancing, internal file labeling, file restrictions
etc. They have usually the objective to detect loss or
non-processing of data, to determine that arithmetic
operations are performed correctly, to determine that all
transactions are posted to the proper file record, to
ensure proper handling of error-conditions (by bypassing
of erroneous records as implicit above), etc.

Although for instance Orlicky (summarized in appendix Al)
and literature on auditing of internal control of EDP
systems show a higher degree of sophistication in terms
of recommending consistency tests between files, espe-
cial design of test data, etc., they really seem to subh-
scribe to the communication-review approach and cannot
come into question in this context,

It is however known that EDP apnlications for scientific
computations, such as found in nuclear physics, structu-
ral analysis, and numerical-analysis applications allow
for a wide range of controls or test procedures which
guarantee the accuracy of the results, Ts it possible to
learn something about the nature of such tests in order
to broaden the limited scope of the present naive con-
trols in EDP, to suit the problems of information systems?

o

4.32

A review of the nature of scientific method indicates
that there are very specific reasons why so-called sci-
entific computations, for example applied to analysis

of force-systems in space (such as found in aerodynamic
problems), aliow the design of mathematical programmed
checks which may detect errors. Such detections of errors
in the course of an EDP-computed structural analysis may
indeed assure a desired level of ACCURACY, for example by
relating aspects of the problem expressed in both STATICS
and GEOMETRY,

The reason why this is possible,however, is that the

IDENTIFY OBJECTS IN THE NATURAL WORLD, for the purposes
related to the use of physics today. In other words,

they specify for an observer HOW AN OBS"RVATION IS TO

BE MADE in order to have meaning, i.e. in order to be
PERTINENT to the answer of certain types of questions,
Being so, it is possible in the context of a computerized
structural analysis to make pertinent observations (col-
lect input data) in order to perform INTERNAL CONSISTEN-
CY CHECKS, as in the Morgenstern sense, based on the
integrated ~ interrelated models or theories.

The matter is comprehensively discussed by Churchman
(1948, p.117), who proceeds showing that IN GENERAL, i.e.
for examvle in studying phenomena more complex than just
moving particles -(as found in administration) ,geometry,
kinematics, and mechanics are indeed NECESSARY, but by
far not SUFFICIENT to guarantee the PERTINENCE OF OBSER-
VATIONS in answering questions about the natural world
(object system). In particular concerning PROBABILITY,

on how to know something about the universe (copulation)
from which the observations are drawn when it is not pos-
sible to make all the observations, it can be said that
presuppositions must be considerably extended beyond the
purely statistical in order to define PERTINENT observa-
tions.

In light of the above problems, we get once more confir-
mation of the relativity of "facts", and of the difficul-
ty but also of the necessity to find some method for
VALIDATION or verifiability of information systems.
Instead of searching for such verifiability in terms of
meaning and TRUTH based on values, efficiency, or facts,
as suggested by our discussion up to now, and by appendi-
xes AL to AJ, we will attempt the following. We will sug-
gest the development of a CRITERION OF MEASURABLE ERROR,
in terms of redefined concepts of ACCURACY and PRECTSION,

4.3.1

4.33

THE CRITERTON OF MEASURABLE ERROR:
REDEFINING ACCURACY AND PRECISION

A criterion of measurable error implies an understan-
ding of what FACT is, that is, it leads to a defini-
tion of what is to be meant by "a question of fact",

As expressed by Churchman (1948, p.217), under such a
criterion a question of fact is said to have meaning if
(in our own words)

1. We can express an answer

2, Measure the error of the answer

3. Reduce the error

Under such postulation, one may ask what "answer",
‘error”, "reduction" etc, mean and still the answers to
such questions may be given and their errors measured,
"The true nature of reality can become a meaningful pro-
blem for discussion, despite the fact that reality is
never directly observed; for we may define the "real"
world as a limiting concept, toward which all experimen-
tal effort is proceeding", Furthermore, it can be seen
that this formulation has an advantage over the vositi-
vistie one in that it does not make any one science ba-
sic to all experimental method,

The misuses of illustrative figures discussed under the
topic of the role of physics in the description of con-
trol problems has probably already justified our "verba-
lism" and restrain from drawing figures in this paper.
Figures may be seen as a kind of language, and it was
seen to imply in turn some theory. In particular we meet
the paradox of not being able to discuss truth in one
same language, as illustrated by our figure 4,8, and we
are not sure of what are the implications of illustra-
ting Morgenstern's concept of information, as in figure
4.9 in terms of a theory of geometry. Such paradoxical
aspects of language and logic are discussed, for exam-
ple by Churchman (1968b, p.108) and in another more
vague cybernetic-oriented sense by S.Beer (1967, p.69)

It is apparent that such problems of illustration, re-
presentation, and expression hide an important dependen-
dence on the basic concept of "truth", as discussed in
our paper, which may be of the utmost sienificance also
in the context of so-called artificial intelligence.

We can, for example, read M.E,. Maron stating: "In order
for an artifact to exhibit indications of knowing, gai-
ning information, etce., it would have to embody a model
of its world"! Furthermore he cites:"In order to display
behavior indicating a comprehension of the difference be-
tween language and what language describes (and also how
language is used), an artifact would have to embody a mo-
del of both the communication process itself and the ori-
ginator of a message as a goal-directed entity who uses
messages to update the internal state of the receiver,"
(Maron, 1964)

4.34
With such reservations about the possibilities for gra-
phic illustration, we suggest the following illustration
for the purpose of stimulating the thought on the issue,

/ Assump- /
tions

'
eee

Lt

~, i
ea
' “. Data
rie aoa peal beck.
AN 2A 34
i F bh 7 =
| iSpecif. of
[pres ess, Tnout
| Parameter ( routine)
| fe nie ei "Data"
| ‘
| i ie ‘ re ; "
! "Controlling
a, Laer 4 Information Sys.
5 Problem-solving
Information System
| {
SAL Niza >
: . Purposeful
t :
i eas j "Contro1"/
Pim ee jobserva-
i tions irks |
fa pee , a
: Pa 7 a
~
6 | "Independent"
i Error Computation
| |
6A t
! Error a
Disagree A
ment | :
t a ‘
v ee i
as | Pe ae aS :
ol Pe se eee
7 ;
TEI oo
‘Complete
output =
input to Figure 4,10

j next Tentative visualization of "fact"& error

hh

4.35

Information process 1 stands for those psychological

and social processes leading to the ASSUMPTIONS 1A.
Information set 1A represents for example human langua~
ge and law, (by which the highest values and goals may
be expressed, or agreement reached in the context of

a debate). Furthermore, 1A stands for the theory of
physics which describes e.g, the techniques for the
manufacture of computer hardware, or the technologies
relating input resources to output products in physical
processes, The assumptions 14 include also economic the-
ory, which indicates what is going to be considered as
costs of resources or development effort, or what is

the expected relation between sales and profit, or ru-
les for calculating profit or "soundness" of the busi-
ness operations: 1A will include also logics and arith-
metics determining e.g. that two different quantities of
the sare product cannot be produced at the same time.
Logic will also be the basis for developing computer pro-
grams in process 21 The assumptions 1A may also include
the formalization of attitudes towards risk as expressed
by constraints on resources, as well as "intangibles"
such as product sales price (or demand for output), and
the estimated opportunity costs of the investors.

The assumptions LA are first used in the process 2 of
designing the methods of processing the information
later derived by the process 3, as "inputs" to the in-
formation system. :

The information set 2A and 3A (describing the METHODS

OR PROGRAMS for processing the INPUTS STORED IN THE DA-
TA BANK) constitute together a description of the INFOR-
MATION SYSTEM. It may be thought as a complete descrip-
tion in the sense of including manual procedures, des-
cription of EDP programs as well as description of the
hardware, All this will be in terms of Language, logic,
mathematics (e.g. for numerical comouting procedures),
physics (for the hardware), etc.
Process 5 describes the actual computation on the basis
of the specifications in 2A and 3A and it was the focus
of the earlier seen Von-Neumann & Goldstine's paper.

It result in 5A is the OBSERVATIONAL REPORT IN
CODED FORM, THE OUTPUT DATA from the operation of the
information system, Such output, a criterion variable
or more generally an intermediate computational result
is controlled by means of the observation 4A, This in-
formation set is actually obtained from a measurement

coding at process 3 and the subsequent processing by

the special-purpose information system, The purposeful
CONTROL OBSERVATION 4A may, if seen in greater detail,
have been obtained by a method similar to 2A and 3A, and
it may be different but not necessarily more TRUE than 5A,

4.36

As a matter of fact, the important thing to note now is
that TRUTH will be a function of the ERROR 6A obtained

by comparing, in some sense 6, the information sets 5A

and 4A and expressing their DISAGREEMENT in the infor-

mation set 6A,

The disagreement 6A may then be seen as a measure of

the differences between the two methods of observing,
measuring, L.e. more generally of predicting since as
Shewhart and Churchman show, every measurement involves

a prediction, THE MOST IMPORTANT ELEMENT OF THE DIFFE-
RENCE BETWEEN THE TWO METHODS, HOWEVER, MAY BE THE ASSUM-
PTIONS 1A, AND THE MOST IMPORTANT ELEMENT IN THESE ASSUM-
PTIONS MAY BE THE IMPLICIT VALUES OR GOALS. This is es-
pecially possible if we note that in LA we should in

fact have included e.g. psychological and socinlogical
theories. Since such established theories do not exist,
or at least ar not considered in the design and operation
of information systems, they are indeed substituted by
implicit uttwarranted hunches on psychological and social
behaviors It is therefore possible that the difference

IN PERSONS performing the processes 2; 3 and 4, that is,
INTERPERSONAL DIFFERENCE is the most important aspect

of disagreement for detecting differences in assumptions
and allowing an iterative revision of them.

We conclude the overview of the figure, observing that
process 7 combines the specification of the measurement
result with its error, leading to the final OUTPUT infor-
mation from our information system,.information set 7A
which may be regarded as INPUT to the next system desi-
ring to use it. We see now why we did not until now dis-
cuss the difference in the problem of quality of input

or output information, The same principles for specifying
the quality of our output, should be used for requesting
specification of input 3A. If this had been done for the
input 3A, then we could at the process step 5 compare the
reported disgreement (quantitatively or qualitatively
defined) with our own QUALITY REQUIREMENTS, for instance
in terms of MAXIMUM ALLOWABLE DISAGREEMENT, We could then
reject a particular result of process 3, that is an in-
put value right away and refuse to process it further

in the routine programs of 2A. This would be tantamount
to creating general criteria of "pertinence" of observa-
tions,.

For the sake of completeness, it should be noted that
"errors" could be also defined at e.g. levels 2A and 5A,
It is possible to check the "soundness" of a design on
paper of an electronic circuit, made at the stage 2.

In such a case it is easier to allocate the error, than
aif it is allowed to combine with other errors and to re-
sult in the later deviation 6A, Deviation or error, or
disagreement 6A may in fact, to the extent that we have
no "total" theory and criteria of pertinence, be alloca-~
ted ("fed back") to any one or several out of all infor-
mation processes 1 to 6, implying a statement of "cause"

437

It is now apparent that the above mentioned hunches on
psychological and social behavior, in 1A, such as as-
sumptions on the political effects of the information
system or assumption on human behavior in the measure-
ment situations (e.g. his cooperativeness in following
the operational instructions, or his sensing-coding ca-
pabilities), will originate deviations which cannot be
detected at early stages 2, 3 or 4. The deviations may
therefore sum up at the level 6A, and the final alloca-
tion may happen to be made by the “authoritative judge-
ment" of the controlling observer or analyst who perfor-
med the process 4, It is believable that he will not
assign the deviation to himself not to his colleagues
analysts who performed the process 2; not either to

his own managers who performed the process 1, It might
therefore be in the nature of the situation that devia-
tions ate assigned to the process 3 performed by clerks,
(and not including input design-parameters who belong

to process 2).

to ihe communieatdonzapproach Figures w2el and 2, 5 seen

isthe ae we feel that figure “hy 10 may be reduced

to Von-Neumann's and Goldstine's approaches (1947, 1956)
by abstracting the physical, logical, and numerical-ma-
thematical aspects from the elements of the figure, (see
figure 4,6), Finally, figure 4.10 also ecompasses figure
4.2 in the sense that fig.10 allows for prediction and
definition of error, which are the background for the
idea of prevention and detection. Correction has not
been represented as such in fig.4.10 since it is an ac-—
tion in the natural worid and not information, that is,
a description of it. It should be noted, however, that
SPECIFICATIONS of actions are contained in the operatio-~
nal definitions of measurements such as those occurring
in processes 3 and 4 of fig.4.10. To the extent that
errors are allocated to 3 we would then expect changes
of the operational definitions of the measurement of
routine inputs to the information system (i.e. CODING)
in the direction of making them more detailed; this
amounts to attempting to constrain the actions of clerks

It is possible to see how this eould he illustrated in
the case study of our appendix A3, where most errors
in the summary list might be prevented by means of mo-
re detailed operational instructions for the measure-
ment of e.g. the quantity of parts in a bin.

However, to the extent that the operational instructions
for the measurements cannot be followed, i.e. are NOT
followed, the error will subsist and it will require
either a relaxation of the allowable error limits (tole-
rance limits), a reallocation of the error to other ele-+
ments, in particular a change in the assumptions,
because of a detected constraint in the natural world.
Increased tolerances means abandoning scientific method.

4,38

This follows from our initial definition of factual
question in terms of the criterion of measurable error:
point 3 stated that it must be possible to reduce error.

In order to limit the scope of the paper at this point
we have only some cursory further comments about figure
4,10, We think that its implications are in line with
the spirit of the literature referenced in appendixes

Alu to A7. The concept of ERROR that it illustrates re-
presents a partial systematic evaluation of judgements
in terms of a measure of DISAGREEMENT. As such it is

an anticipated indication, a guarantee of possible value
of the information for a decision-maker, but without
necessarily referring directly to values, and in this
sense indicates a degree of truth or factuality.

Such measure of error may be seen as an overall ACCURA-
CY-PRECISION which characterizes both the information
process leading to an observation, and the particular
observation as related to the process, The error defi-
ned in figure 4.10 is a measure at a more general or
"later",less detailed level than analog errors that
could be defined through the breakdown of figure 4.10

in more elementary problem-solving steps (subsystems

of the information system 2A and 3A) .At each level such
errors aliow the possibility of raising the question
"WHY ?" for the disagreement and in this way they may
detect e.g. problems of "pertinence" and of time synchron-
nization, i.e. "timeliness" where time is seen as a tool
for individuation and identification,

Furthermore, it should be noted that the error concept
illustrated by figure 4.10 does NOT by itself imply
control, but rather only the nossibility for it. Control
is the long-run aspect of accuracy, and the problem of
control is the problem of determining when and where to
test for accuracy, i.e, at what points of the overall
process,error should be measured and what should the
maximum allowed error (tolerance limits) be. To say
that one cannot afford to measure error at any point,
any time in the process, is equivalent to allow an in-
creasing unknown tolerance of error, i.e. to give up
control, or as already seen, to abandon scientific me-
thod. In this sense we touch also upon the scientific
meaning of OBJECTIVITY versus SUBJECTIVITY, since a
“subjective answer" may be seen simply as Lacking a
(long-run) control. (Churchman,1948,p.165; 1968b, p.118
and 123). To search for disagreement and to explain it
through reduced error, is to strive for objectivity.

Finally it appears that means-—ends analysis (Simon,1969,
p. 66-69) as commonly understood in present research on
computerized problem-solving or "artificial intelligencd
may be seen as a special case of the more general means-
ends analysis, and general concepts of "production" and
"precedence" related to fig. 4.10 as in part suggested
by Ackoff (1962,p.172), Churchman (1948,p.164;1968b,
p-72,102; 1961,especially criticism on p.376 , and p.99),

“~

4.3.

4.39

THE DEFINITION OF AUCURACY AND PRECISION

Up to this section we have mostly talked about ERROR in
terms of disagreement or deviation without closer spe-
cification of how it should be defined in an administra-
tive context. The starting point for this section will
be the statement reproduced in appendix A7:

"If scientific method is to be extended to decision-
making in general, the ideals of accuracy and control
will also have to be redefined."

We will be aware of the danger of falling into the

naive fallacy of idoking for some "true" definition,

We will instead apply the criterion of measurable error
to this definition problem, and expect that such error
will in some sense be measurable in terms of results or
eventual debate about it. With this in mind we may
recall what was said in the tontext of control of mass
manufacturing; to paraphrase Shewhart:"Disagreement of
results among themselves" is itself not very definite
because there is obviously and indefinitely large number
of senses in which results might be said to disagree
among themselves. We might, for example, think of their
disagreement in terms of the way they cluster around the
observed average, or in terms of the magnitude of some
one or more of the indefinitely Large number of symmetric
functions of these data. Or again we might concer our-
selves with the order in which the observations appear:

For example, a special commission of the International
Society for Photogrammetry dedicates a whole chapter

of a paper on "Quality Problems in Photogrammetry" pu-
blished in 1967, to the analysis of basic concepts and

terminology including accuracy, precision, deviation,

error, and weight, It states e.g, that precision may be
expressed as standard deviation of a single observation
or of the mean (or other funetions) of observations.
Accuracy may be expressed as root mean square value of
errors or discrepancies from the given true value, or as
standard error of other functions of observations.

In administrative situations the theoretical foundations
for such definitions cannot be expected to hold except
for possibly the most trivial routine data~processing.
The universe of observations is not defined, their dis-
tributions are not known, in particular REPEATABILITY

is not found, and the traditional notions of error - in
the statistical sense - do not hold. Many aspects of
this problem have already been considered in our paper,

o>

Returning to figure 4,10 we begin by noting that in
discussing the information set 6A, error, we made referen-
ce to the difference between TWO METHODS of observing,
measuring, predicting, and we mentioned that INTERPERSONAL
difference might be the most important element of such
difference.

“aA4O

This appears to be consistent with what Kaplan calls
INTERSUBJECTIVITY, in appendix A7. We feel that this
has to do with the fact that the absence of a psycholo-
gical-sociological theory prevents us from immagining
some "objective" impersonal meaning of the vague wor-
king concepts of "goals" or "values", This warrants
that we stick in first place to PEOPLE, to OBSERVERS
and OBSERVED, 2

For the "practical" mind the above cannot be over-empha-
sized in the context of posing the question: " WHO will
pay ?" In connection with the material referenced in
appendix A2 one may discuss for example reject rates

and error rates of OCR equipment. In connection with

the general issue of so~called validation one may dis-
cuss verification costs versus error costs. Sometimes

it is stated that "a relatively high error rate may BE
TOLERATED...". In discussing the figure 4,10 as well

as in chapter 3 we discussed the assignment of coding
errors to the input clerks versus assignment e.g. to
system design. In some literature on computer-aided
medical diagnosis (outside the scope of appendix A2)
sometimes reference is made to the "patient's satisfac-
tion" and to the "physician's decision" with due consi~
deration of "the problem of dollar cost", to the "utili-
ties of death and cure" relative to the dollar costs of
tests, etc.

The practical mind will probably not refuse to consider
the questions of who will pay for the rejects respecti-
vely the costs above: the customer of a telegraph com-
pany may receive an illegible text (see appendix A2,on
accuracy of communication links) and the company may

be happy in requesting a retransmission rather than
preventing such event, whenever the customer complains.
Would such policy be accepted in computations of sala-
ry payments ? The question is who will pay for verifica-
tion respectively error costs in more complex contexts
of large, say, public data~banks. Will the clerk or
system designer pay for the error in the final result ?
"High error rate may BE tolerated" ~ the question is
tolerated by WHOM ? It is a very important practical,
and therefore also scientific question to investigate
who will decide what is to be tolerate by whom. And fi-
nally in the case of computer aided medical diagnosis
we meet the most important question of the world: "WHO
WILL DIE ?", Who will pay for the diagnostic tests and
estimate their marginal utility versus maximizing the
patient's satisfaction ? We have seen at least one pa-
per where an interviewable patient was not questioned
at all about his preferences for alternative disabili-
ties following physician's alternative decisions. The
patient was not represented in the decision model since
the physician made ail the estimations for the patient's
best satisfaction!

Furthermore the physician's estimates may be formalized
in terms of certain models for formalization of utili-

AaL

ties or valties, Such models are based on “rational rules
of behavior" and "game theory" which are scientifically
highly questionable. Churchman, (1968b, p.98) summarizes
an extensive criticism against such thinking

The above few examples are intended to suggest the extre-
me importance of WHOSE goals and observations as rela-
ted to WHAT goals and observations, If the intent has
been attained then one gets less surprised for example
in noticing a great number of"errors" being"discove-—
red"suddenly in a EDP file as soon as it begins to be
used in an application that serves other people than
than those who create the input. One might also get less
surprised in front of the difficulties of standardizing
so-called data-elements or elementary items of informa-
tion across geogtaphically dispersed units of a corpora-
tion. It may be more than a question of goodwill in sol-
ving misunderstandings: our own experience supports what
we referred to in appendix Al - as an example one "date
of transaction" may not SATISFY ALL USERS,

There ate, however, much deeper reasons for considering
the primacy of the WHO question in the context of truth
and disagreement. Many of us have sometimes felt puzzled
by the vagueness of the problem of validating SIMULATION
results, as well as the vagueness of the literature dea-
ling with this problem. The reason for this, obviously
is that one must SIMULATE SOMETHING and this something
should conceivably be TRUTH... We may, therefore expect

to meet all the truth problems discussed up to now in
our paper. From the only paper which we know discusses
such aspects of simulation we find the following of im-
portance for our study. (Churchman, 1963)

The concept of REALITY is meaningful only when there are
at least two minds. A single mind, receiving "inputs",
has no way of recognizing what is simulation and what is
real, The second mind observes the ENVIRONMENT of the
first, recognizes the sources of the inputs, recognizes
how the first mind responds. The observing mind has a
purpose in making the observations. What it should cons-
true as the REALITY OF THE OBSERVED MIND is based in
part on this PURPOSE.

Reality is then a mode used by the observing mind to
describe an observed mind, and the observing mind has

a choice as to what it should assign as the reality of
the first observed mind. Whether or not the choice is
correct depends on a third mind, one that judges the pur-
poses of the second. The second mind cannot know the re-
ality of the first until all observing minds are content,
and this contentment is an unattainable ideal.

A practical organizational implication of the above is
that a system that approximates reality must include both
rules by which data are collected (responsibility for au-
thenticating them) and construction of model for proper
assignment of causes (by tests) if trouble occurs.

4,42

In summary, the concept of reality is basically inter-
personal, or to use Kaplan's word, intersubjective,
prior to be anything like "purposeful", Indeed the
concept of purpose appears very soon in the above pro-
posal, but already as an attribute of a human, Further-
more, it appears to us very promising that the proposed
concept of reality on one hand has a deep philosophical
justification in terms of the criterion of measurable
error, and on the other hand it is consistent with
recent trerds in social psychology which are emerging
after several years of strong debate,

This supports, then,our generat discussion on alloca-
tion of errors in the figure 4.10, and in particular
our statement that the control-observation 4A may be
different but not necessarily more true than 5A. On the
contrary, the proposed concept of reality makes truth
itself dependent on the relation between 4A and 5A,
Furthermore, the proposed concept of reality shows that
the NUMBER of controlling-observers is a relevant va-
tiable in the test of the input information and of the
results from the infotmation system,

Churchman (1968b, p.86) summarizes some of the points
above in the following wotds: "A researcher is not a
special kind of person; rather every person is a special
kind of researcher... One of the most absurd myths of
the social sciences is the "objectivity" that is alleged
to occur in the relation between the scientist-as-an-
observer and the people he observes... Instead of the
silly and empty claim that an observation is objective
if it resides in the brain of an unbiased observer, one
should say that an observation is objective if it is

the creation of many inquirers with many different points
of view." And further: "The real expert is still Every-
man, stupid, humorous, serious, and comprehensive all at
the same time. The public always knows more than any of
the "experts", be they economists, behavioral scientists,
or whoever; the problem of the systems apnornach is to
learn what "everybody" knows."(1968a, p.231)

On the basis of what we have developed up to now in this
study we cannot but agree with the above statements,

They are also consistent with our own experience. The
problem then becomes for us the Lastly mentioned of
incorporating the ideas as they relate to specifying the
quality of information to the methodology of systems
design. Without pushing much farther the use of the fi-
gure 4,10, we ask ourselves how to design the process

6, that is, how to compute the error, In a subtle way;
through the feedback of error to different processes

we are also asking for the optimal design of 4's or the
proper selection of the 4A'ts. We are looking for the most
severe test, generating the largest disagreement within
the constraint of a limited number of control-observations.

44g

We urge the reader to notice that this step of inquiry
is dedicated to the generation of DISAGREEMENT, and not
of the more intuitive-common concept of agreement.

From the most successful science of physics, and from
literature on scientific method it can be Learnt that
agreement by itself does not have a definite meaning.
Agreements reached about, for example, observations of
physical events must be reached in the context of CARE-
FUL CONTROL. And control of observation means that

"the scientist is capable of judging whether or not ex-
traneous causes have influenced the observations; it
means that he can judge the extent to which the observa-
tions have been influenced by unforeseen or unknown
events.",Agreement is in science considered to be a
dangerous basis for rational conclusions: it can rather
be regarded as a kind of evidence of danger ahead, We
have in appendix A7 also touched upon the fact that no
scientist seeks to obtain absolute agreement of obser-
vational reports, because such agreement contains no
information about the nature of the system he is gstu-
dying. Disagreement is the way of Aiscovering hidden
unchallenged assumptions. Each time a scientist obtains
agreement in his instrument's reading, he will try to
push them to the next decimal place. Or, as Ackoff
expresses it (1962, p.251), the scieritist may suspect
that his instrument is jammed or has not sufficient
sensitivity: he will investigate the cause of CONFORMITY
and "correct" it so that he gets variation among obser-
vations. This process yields ever-increasing ACCURACY
of observations !

We see then that the real problem is not to obtain agree-
ment: it may obtained by jamming the instrument or by
silencing those who disagree: the problem is rather to
PROVIDE BY MEANS OF RATTONAL DESIGN THE STRONGEST POSSI-~
BLE KIND OF DEBATE, This might be the meaning of forma-
lizing at least a part of the judgement process, and
this is what,for example Shewhart did in the context of
manufacturing quality control, when he avoided the need
to rely on the subjective judgement of the "experts"
engineers or scientists (See appendix A4), If this is
so in manufacturing, then what to say about judgement
in the context of complex social-technical problems
where we are constantly asked to rely on, to trust, or
to have faith in this or that "expert"? In a recent
paper, I.T. Mitroff (1971) summarizes many of these
points. In an age where many important social issues
cut across expertise and fields of study, and where the
consequences of believing in experts may be deadly, it
is foolish to just trust in experts. "WOULD IT NOT 3E
BETTER TO SPEND THE TIME REMOVING THE CONDITIONS THAT
MAKE TRUST NECESSARY, RATHER THAN DEVELOPING THE CONDI-
TIONS FOR BUILDING TRUST ?" What we need is the capa-
bility to maximally challenge an expert, because if we
ean do this, then we have less need to "trust" him.

If we want to regard truth as a kind of agreement, the
latter must coneern the method of resolving disagreements.

aay

We will, for the purposes of our work, propose the defi-
nition of truth, as being agreement estahlished in the

agreement.

If we think of judgement as a result (an information set)
rather than the process generating it, we will say
that agreement is a judgement in the form of an "output"
final value, for example as expressed by the average of
a set of pointer readings. (Sound) judgement will be the
result of establishing agreement, for example by some
kind of negotiation, in the context of the strongest
possible disagreement. The latter may be expressed, for
example,by the standard deviation of the set of pointer
readings; it represents the degree of doubt (or belief)
in the judgement. |

In the light of the earlier expressed doubts about the
graphic representability of the above language descrip-
tion, we will attempt to complete the lower part of fi-
gure 4,10 in order to illustrate the above ideas.

TAL 7A2
a Oa he j
Output Information / Contract,/
ub jectiv olerance
Value { Error iLimits
£ wt rm Po, | |
x y |

Negotiation 9

. Sy
os |
——_ !

/ \ |

|

|

|

a ee ae OB Ne  -
/Agreed "objective" /
7 output | New
Buen) ered Deeeee of} A contract
value doubt / | |

y sine —

v Sistas ea ot
"sold/bought"

output
Figure 4.11

ACS

In the relation between figure 4.10 and 4.11 we recogni-
ze that while process 6 of figure 4,10 was the first
step of control (measurement of disagreement = error)
such step was necessary but not sufficient for control,
It is possible that the nature of disagreement and error
6A is such that the "right" 7A, and automatic allocation
of 6A to pertinent processes cannot be told. To the ex-
tent that negotiations must be anyway set-up for alloca-
tion of the causes of the error, they may also influence
the generator of the output 7A to revise 7Al to 9Al.

He will, in other words, be in position of choosing whe-
ther to keep 9Al close to 7A1 and having to declare a
great error 9A2, or alternatively get influenced by tho-
se who disagree and revise substantially 7Al to a quite
different 9Al, in which case he will be "premiated" by
being allowed to declare a smaller error(collective de-
gree of doubt)9A2, We see then that the generator or
responsible for the conputation of 7Al is "free"to render
the account he wishes, but he is bound to account for
his error, His freedom, however; is limited to the extent
that he has a contract 8A to follows

In the case of Shewhart's control of mass-manufacturing,
the contract could be seen as signed with the buyer of
the produced product, who was then authorized to perform
the control-observation 4 (fig. 4.10) in order to check
whether the tolerance limits were satisfied, The contract,
however, at early stages of manufacturing could be seen
as signed by the manufacturer (running the information
system 2A & 34 for his product), so-to-say with himself
in order to stay in business. If the manufacturer did

not respect the tolerance limits at early stages of manu-
facturing, then his information system based on the the-
ory, say, of mechanics for his mechanical product, may
predict that the final product will not satisfy the to-
lerance limits on the contract with the buyer: if he goes
to court he will be imposed to keep his product, refund
the presumptive buyer, and perhaps (also legally) imposed
to stay out of business - an outcome which perhaps would
already be economically determined,

At a more general level than physical manufacturing, nego
tiations according to figure 4,11 will have to be conduc-—
ted whenever there is a contract 8A specifying e.g. tole-
rance limits that somebody reports as not being satisfied.
Analyzing figure 4,11 again at a general level, we will
consider 7A as composed of the unchanged value 5A = 7AL +
the measured error 6A = 7A2 (compare with figure 4.10),
The value 5A may be seen as the subjective report of the
decision maker running the process 5. The contract 8A

may be seen as a kind of group goals, attained through
earlier negotiations, including rules for negotiation,
and in this respect it is one meaning of the "agreement"
associated with the result 9A of the negotiations, The
eontract includes also some kind of specification of the
"object"-identity, and stability,

Sard

446

We shall now say that 7Al and 7A2 together constitute
the"evidence" 7A on which negotiations will be conducted
in the light of the contract 8A which is an aspect of
the assumptions 1A in figure 4.10.

On this basis, the following process 9 may be seen as
taking place at the input of an information system, such
as the case would have been at process 3 of figure 4.10,
in case the description of desired processes (programs)
2A had furnished the contract terms at 3.

The negotiation 9, then, is the second step of control.
The first step 6 determined which is the maximum possi-
ble disagreement (error). The step 9 determines whether
this disagreement is greater than the specified in the
tolerance Limits of the contract, Sometimes we find that
the term "error" is reserved to the event when the magni-
tude of the disagreement is larger than the allowed by
the tolerance limits, We do not follow this usage. Step
9 summarizes also value; e.g. economic, considerations

as implied in the setting of the tolerance limits. The
step 9 may be seen as determining the answer to"WHY ?"
(the error), and "WHAT TO DECIDE" (the output, objective,
predicted value for the overall computation). As mentio-
ned earlier there may be possibilities of trade-off,
within the tolerance range, between the prediction and
its degree of doubt (9Al, respectively 9A2), The predic-~
tion is "sold" at the input of the next information sys-
tem, which is then certain to accept it as objective and
true. The degree of doubt (or belief) is then fed back
to the agreed-upon processes,in the form of specified
changes in the resulting information sets, The informa-
tion set 9A represent the "agreement",

Another result from the negotiations 9 may be a revised
contract 93, which, to be consistent with our understan-
ding of scientific method in terms of the criterion of
measurable error, should in the long run lead to decrea-
sed tolerance range.

Tt should be noted that tolerance ranges are idealized

as being tied to fixed (true ) value . In a general ca-
se where we have no theory, it can be apnoroximated by
a function of the -observations, such as a maxi-

mum standard deviation. between 5A, and all 4A Ss, to

be compared with the same function's result in the parti-
cular case (6A). In order to permit the described trade-
off between 9A1 and 9A2, we could furthermore compute

942 as a root mean square function of the discrepancies
between the 4A's and the chosen 9Al.

We can eventually summarize with an overview of figure
4,11 in the following terms: The evidence 7A is submitted
to a judgement process 9 which making use of values and
assumptions in 8A leads to an agreement unon what is ta
be considered as a sound judgement of the predicted value
9A and of what should be done for future improvement.

N47

In the language used by Shewhart, then, a judgement pro-
cess always involves a specified evidence Cae nest
and a specified prediction (sound judgement). The jud-
gement may be valid, and still the prediction may be
false, since a sound judgement is incorporating a desree
of rational belief, for example in the nature and origin
of disagreement, on the fairness of the rules for the
judgement process, and other assumptions, Or, to para-
phrase Churchman, in societies with powerful ruling clas
ses it is easy to define rational planning, reason, ru-
les for sound judgement and overall fairness of assunp-
tions; much as reason in any patriarchal household ts
the principle that "Father knows best", reason in such
societies is taken to be the set of principles that

keep the ruling class in power. (Churchman ,1958b, p.98)
It is apparent that the falsity of a prediction based
ona "valid" judgement in such a social setting, may be
"proved" in terms of the results of, say, a rebellion.

As Shewhart understood it, knowledge or truth may be

seen in terms of its fundamental components:

1. Original data (evidence)

2. Prediction, with an operationally verifiable meaning
which can turn out to be false even if the judgement
is valid in terms of valid assumptions.

3. Degree of (rational) belief in the prediction, based
on the evidence. ;

Knowledge begins in the original data and ends in the
data predicted, these future data being the(operational-
ly verifiable)meaning of the original data. (Shewhart,

1939, p.86,122,143).

In the context of our attempt, now, to define accuracy
and precision in a social environment, such as data-banks
and information systems used in business and in public
planning, the above problems of "knowledge", "judgement",
etc. reappear in paradoxical questions. For example, in
order that the predicted objective value 9A in figure
4,11 be "true" in our proposed sense, the disagreement
7A2 must be the strongest possible, i.e. the error must
be the largest possible. Possible FOR WHOM ? Disagree-
ment BY WHOM ? Error computed by whom ? Maximum disarree
ment requires that the controlling "indenendent" ohser-_
vers be"free" to report their readings or judgements,
that is, they must NOT BE UNDER THE CONTROL of the deci-
sion-maker who generates 5A. Who will determine whether
they are or are not under such control ? In some sense
such questions have a judicial character.

Within the scope of this paper, we shall propose a tenta-
tive definition of accuracy and precision as two aspects
of error,:- We expect that they will be object for the
"strongest possible" debate leading to their gradual
refinement. They will be based on the fundamentally im-
portant ideas of IDHNTITY or SUBJECTIVITY, and INTER-
SUBJECTIVITY.

~

4,48

ACCURACY = Is a measure of the reproducibility of an
observed, cnmputed value, of a prediction,
of a judgement, TO THE EXTENT THAT IT IS
APFECTED BY WHAT IS NOT UNDER THE CONTROL of
the particular observer, computer, predictor,
or judge, i,e, humans to whom we will refer
as DECISTON-MAKERS,

PRECISION- Is a measure of the reproducibility of the
same as above, TO THE EXTENT THAT IT 15
AFFECTED BY WHAT IS UNDER THE CONTROL of the
particular decision-maker,

By means of the above definitions we attempt to cavture
the nature of the alternative definitions found in appen-
dixes A4 to AJ, as well as to meet the criticism and
ideas presented in this chapter up to now. In some sub-
tle sense, our concept of precision aims at guaranteeing
the identity of the observer or of the observed, which
is a necessary condition for the more meaningful discus-
sion of intersubjectivity in terms of accuracy and truth.
We regard then accuracy as the most important concept,

a measure of truth, while precision is a necessary con-
dition for the measurement of accuracy, Accuracy, in so-
me sense aims at generality of application in the inter-
personal dimension, while precision aims at generality
of application in the time dimension,

A starting point for a refinement of the above ideas is
provided e.g. by Ackoff (1962, p.210,251,11), Churchman
(1961, p.216; 1968b, p.34; 1948, p.141).

Two distinctive features of our definitions are the lack
of emphasis on REPETITIVITY and on METHODS of measure-
ment. We justify the first on the basis that repetitivi-
ty is usually required as a means of substantiating jud-
gements in terms of objective probability. We feel, how-
ever, convinced that such means of substantiating judge-
ment has no primacy over other ways as proposed here,
since "objective" probabilities and counting of relative
frequencies makes strong assumptions on the judgements
themselves. (Churchman, 1961, p.137, 169) This is also
the reason why we do not consider Savage's criticism of
accuracy,as relevant to our proposal, while our proposal
should hopefully take into account his emphasis on the
issue of "multipersonal problems". (Savage,1954,p.257,154)

Concerning our lack of emphasis on METHODS, we would like
to propose that methods have not primacy either over
intersubjectivity. In the same way as repetitivity was
tacitly implied in the success of the scientific method,
because of the repeated verification obtained by

rosie

44g

DIFFERENT SCIENTISTS, we expect that relevant differen-
ces of the natural world will be tacitly imolied in the
fundamental difference on which reality itself is hased:
the interpersonal difference, Differences in purposes
to be partially served by common observations and com-
putations, may be the source of the differences in me~
thods, Reference to the theory of physics, for examples
of "impersonal" methods which determine accuracy, would
incur in the earlier seen criticism against the "“under-
lying physical processes" and the role of logical posi-
tivism.,. It is clear that to the extent that we abstract
human elements out of the studied field, and to the ex-
tent that we build a theory of what is left, then such
theory will not be dependent on the interpersonal on
intersubjective differences,

Other important problems raised by our proposal will,
within the limited scepe of this paper,be touched upon
in the next chapter. ith the purpose of stimulate thin-
king in sur proposal, and with no claim of scientific
value, we would like to present the following "flip-
chart illustration" of our concepts of accuracy and pre-
cision, as applied to a business organization.

Organization structure (Decision—-makers) |

eo

TN gets aa :
ee ~ Mee “a |

"PACT" as Q )
"event" on \ .—> ‘Wy iY }
"object" ‘ \ Ny j

SP ge A

. Figure 4,12
"Plip-chart" illustration of accuracy and precision.

4d

4.50

In figure 4.12, decision maker A corresponds to the
decision-maker responsible for the accuracy of the in-
formation set 5A in figure 4.10, while the independent
controlling observers B,C, and D perform the control
observations of the type 4A. Precision is a measure of
A's stability in time, disregarding B to D, in terms of
changes in what was assumed to be'constant in relation
to A, Such precision is used in the computation of accu-
racy which is then fed back to all the decision-makers!
processes."Facts do not exist but are rather represen-
ted by the accuracy. The inclusion of more controllers,
possibly as different as conceivable from A,increases
the accuracy: such difference could be obtained by
substituting perhaps D by one of his subordinates, or
by including somebody from outside the organization,

The concept of accuracy allows to consider as D's suhor-
dinate,professional specialists including "operative"
people such as clerks and machine-shop personnel.

In considering figure 4,12 it should be recalled that
accuracy should be measured at different stages of the
organizational activities. We have not shown, for exam-
ple, the determination of the accuracy and precision con-
cerning the questions or events that usually are the con-
cern of the top-manager of the organization, The pvrinci-
ples for such determination would be analog to the illus-
trated in figure 4,12. In this kind of settings, it is

a relative matter who should be called observed and ob-
server, controller and controlled; agreement may then

be used to determine whether one is capturing the intent
of those who work with a concept.

AN OVERVIEW ON THE CONTENTS OF THIS CHAPTER

After attemtvoing, initially, a traditional systems ap-
prach to the quality problem in terms of prevention,
detection, and correction subsystems, we were confron~
ted with the need of a much deeper understanding of what
quality and error could mean, With this purpose in mind
we turned to more scientific Literature. Administration
and organization theory introduced us to the concepts of
value, efficiency, and judgement, the latter referring
to factual questions and empirical truth.

Judgement, however, was seen to rely on the need for

its systematic evaluation on the basic of subsequent re-~
sults of its application, the same being true of the
factual-empirical questions of administrative and physi-
eal production functions. The most factual-empirical
matters of physical mass-manufacturing did not dispense
systematic evaluation of judzements in terms of accura-
cy and precision, We illustrated theoretically and prac~
tically the untenable division of problems in factual
versus value issues, physical versus administrative

45

AL5L

or organizational-policy issues, including the case of
physical science itself. The analysis of the history of
scientific method offered to us the idea of the criterion
of measurable error. We applied it to the redefinition

of accuracy and precision in information systems which
aim at the control of general activities, in analogy to
the quality control system which is applied to the con-
trol of industrial manufacturing activities. Only under
such circumstances can the creation and use of information
be conceived as a "production" of information without fal-
ling in some of the fallacies of the logical-positivistic
thinking. Such concept we have proposed for accuracy and
precision as related to information systems does not ma-
ke direct reference to values and outcomes and is apparen
tly well suited to general business data=banks aimed at
future unknown needs, as well as to public data-banks.

CONCLUSIONS FROM THIS CHAPTER

1, Information systems and data-banks can be regarded
as integrating different theories or models at diffe-
rent levels of maturity, which require an overall
concept of truth or quality.

2. It is possible to redefine accuracy and precision as
two aspects of overall quality of information, with
the purpose of allowing inferences on the reproduci-
bility of the computational results,

On the basis of the above conclusions, the next chapter
will present the frame for a "handbook of quality control
of information" to be developed in the context of a par~
ticular information system,for use,for instance by the
system.designers, The frame will be presented in terms

of illustrative examples, a discussion of the difficul-
ties associated with the application of our concepts,

and evaluation of available helpful knowledge such as
found in the statistical literature.

5.1
THE IMPLHMENTATION OF QUALITY-CONTROL:

A CONVENTIONAL HANDBOOK FOR
QUALITY CONTROL OF TINPOKMATION

Prior to suggesting any guidelines for the development
of a handbook on the basis of our pronosal in the
last chapter, we will show a conceivable alternative.
We ask the reader to immagine that we take up this
task in the course of our exposition in chapter Pane
that is, after the section which was dedicated to
listing twenty-eight statements based on the results
from our review of the empirical literature.

In such a case we wilt start by referring to appendix
Al and create a definition of quality of information
that in some way, The task would not be easy but still
it would be manageable,for example in terms of combi-
ning the most reasonable definitions and thoughts offe-
red by, say, J.C,Emery, and G,Rodin. We can then state
that some aspects of the quality of stored information
will be taken care of by, for example, stating the
point in time (date) when a specific item of information
was created (coded), updated, computed, changed or

used the latest time. To the extent that we store phy-
sical dimensions stich as width of highways, or weight
of objects, other aspects of quality can be considered
by storing together with the measured values also an
indication of the level of uncertainty, in some sense,
of such measures, say plus/minus something.

To the extent that we deal with information which is
the concern of higher levels of hierarchy, we cannot,
according Emery's implication, expect to measure the
quality in terms of such detailed accuracy but we will
rather look for an authorized statement on its value.

The next step in developing the conventional handbook
may be related to the material presented in appendix
A2, We shall surely note that there is a kind of "gap"
between the theoretical framework supposedly represented
by the earlier definitions. We state, however, that
obviously some hints are required in order to attain
quality of information. From a practical point of view
we see that the empirical literature offers a series
of statements, most of which we attempted to summarize
in the mentioned list of chapter 2. Since several of
the empirical results are apnarently contradictory or
not clear enough for the occasional reader, we analy-
ze them more carefully in order to consolidate them
ina final"set of principles to be followed by the
designer of information systems,"

For example, we start by observing that some statements
are obviously true on the basis of sheer common sense,
to the point of not even having required a costly re-
search for the purpose of confirmation. Perhaps state-

Dee

metit No, 3 belongs to such class of statements, (that is
"avoid characters which pronounced sound alike, eigs

M and N", )Furthermore we notice that statement No, 4
may be not true in its simple form since it apnears

to be questioned by statement No.26: we should clarify
what is meant by significance, meaningfulness, mnemo-
nic, and lotter-pattern familiarity, The next step in
the consolidation of the set of principles, may consist
in noticing that statements l, 24, and 25 have some-
thing in common, and their meaning may possibly be con-
veyed by one same statement, Going further, recalling
what we have read in EDP Analyzer of October 1971 we
fiotice that it refers to an author who questions sta-
ment 28 obtained from Owsowitz &Sweetland: he advises
that "if possible" one should stick to numeric codes
anid avoid alphanumeric ones, This was the reason why
when writing down point 17 of the list, sugsested by
the author referenced by EDP Analyzer, we mitigated

its content for accounting of the conflict with the
later point 28.

This last consideration makes us recall that many other
similar ambiguties exist as implied in the formulation

of points 18 and 19 the subject of which was discussed

in the text of chapter 2,

We conclude that in order to allow the system designer
to use the proposed set of principles, we must refer
him to the literature which originated the statements.
With this purpose in mind we create an overview table
shown in appendix A8. The vague principle for its or-
ganization i to have at the vertical-axis of the ma-
trix several groupings of "independent" variables or .
attributes of situation which may vary in different cir-
cumstances for different information systems, At the
horizontal axis we put an identification of the particu-
lar paper that in some way considers a particular va-
riable.

With the help of the overview table, the system desi-
gner will be able to qualify statement 18, for example,
by referring to Smith and hopefully evaluating other
vague aspects of the issue such as motivational factors,
message complexity, volume of reporting,. cost of entry
devices as well as walking distance to them, time re-
guired for rewrding entries, possibilities of interrup-
ting the primary job, etc.

The following step in developing the conventional hand-
book may be the adaptation of the empirical results to
the particular information system and its environment
by means of specific computations or additional empiri-
cal studies at the local level. As an example the sys-
tem designer may feel that it is relevant for his work
to answer the question:"What is the volume (number) of
errors in the input stream of my EDP system ?"

5.3

One item of the reviewed literature was seen to sug-
gest that a typical job shop with 1,000 employees could
inject into the EDP system about 100 to 200 errors eve-
ry day. In this figure are included several types of
errors other than pure punching errors. If the system
designer rightly feels that such a "standard" figure
will not be applicable for his installation, and wants
to limit his attention to punching errors, he may as-
sume, against the background of the reviewed investiga-~
tions, (overviewed in appendix A8) a typical punch
error rate of 0.1 % after verification. If he calcula-
tes with an average of 50 columns per card punched with
fresh digits (not reproduced automatically from other
cards), and assuming a card reader reading at a speed
of 1,000 cards per minute, the result is an input of

50 errors per ninute into the system during the opera-
tion of the reader, where errors are understood as er-
reneous digits, and prior to any validation or editing
procedures at the system,

A more optimistic estimate could assume a punch error
rate after verification,of 0.01 %, and 10 columns per
card giving an input error rate of | error ner minute
of operation of the same card reader.

Another way of approaching the estimation is by star-
ting with the average number of strokes per day of key-
punch operators, say 70,000, that is ahout 10,000 per
effective hour of work. This implies, with an error
rate of 0.01 % that each keypunch operator contributes
with one punch error per hour into the system.

It may be felt that a more realistic feeling is obtai-
ned if we look at the estimate from the point of view
of "transaction" error. For a digit error rate of 0.01 %
that we look at as an error-probability of 1/10,000,

and for a 10 digits-transaction, the probability that
the transaction will be completely error~free is
(9,999/10,000)exp 10 = 0.99907, where we have accented
the usual necessary assumptions of a constant, indepen-
dent probability of error. This all means that 93 tran-
sactions out of 100,000, or about 9 out of 10,090 will
be in error. With a quite more pessimistic error rate
that may be seen as including certain errors in source
documents, say 1 %, the corresponding transaction er-
ror rate would be calculated at about 10 % for ten-digit
transactions, and 18% for twenty-digits transactions.

ft is now difficult to say where we go from here, after
having made such estimates. It is however conceivable
that they may be useful in certain circumstances, Diffi-
culties will, however, be compounded by the necessity
of considering the effects of validity checks, or for
example clustering of errors, which was seen to be so
important in the analysis of errors in communication
systems (appendix A2, Martin and Norman).This relates
too to the meaning of error "probabilities",

ier
fo

To these mentioned difficulties one could add many of
those implicit in our discussions in chapter 2, In any
case there are reports of much more elaborate probabili-
ty thinking than the applied in the examples seen abo-
ve, which has provided valuable results in structured
military and industrial situations, We have left out

of the scope of chanter 2 the review of literature re-
porting how human-factors specialists use human-error-
rate data and make certain gross behavioral assumptions
in order to estimate human error-rates in the context
of a particular man-machine system,

The interested reader may find a description of a pro-
cedure and some assumptions for estimating error-rates
in a report by A.D. Swain (1963). It is conceivable

that the reported techniques may be adapted to the eva-
luation of the overall turn-around reliability of alter-
native combinations of EDP input-output media and devi-
ces, This implies the evaluation of the reliability,
e.g: in terms of failure and error rates, in the chain
of components of an EDP input-output system. Such com-
ponents may be input-output MEDIA such as punched cards,
OCR (optical character recognition) documents, MICR (ma~
gnetic ink character recognition)cards, magnetic tape,
etc., as well as input-output DEVICES such as card
read/punches, direct entry keyboards (e.g. to tape or

to disc), MICK card reader/printers, OCR readers, high-
speed paper printers, etc.

Besides these special-purpose calculations of particu-
lar error-rates using the "basic error-rate data" re-
ferred in appendix A2, the referred material may pro-
bably be used in order to avoid many "traps" in the
definition and evaluation of errors and error rates.
Definitions and guidelines for evaluation would have

to be contained in the conventional handbook for quali-
ty control of information: a review of appendix A2, to-
gether with the discussion in chapter 2, for example

on the problems of terminology met in reviewing the em-
pirical literature, will enable the avoidance of vari-
ous ambiguities, They were seen to appear, for instan-
ce in the dimensions of errors percent of digits or
of characters, or of entries), In the context of OCR
error rates one could, for example, refer to the LOWER
error rate of an entry procedure compared with another,
but the LOWEK referred to lower rate of wrongly identi-
fied characters, thanks to an earlier stage of typing
where transcription errors were introduced: the overall
error rate in the considered stages could actually turn
out to be HIGHER, not Lower.

The next step in develoning the conventional handhook
may, on the basis of the developed terminology attempt
a classification of errors on the basis of their vague
nature and their relative rates, We suggested in chapter
2, and expressely stated in statement No. 16 of the list

5.5

of statements that certain kinds of errors at certain
stages of the system operation, namely "source" errors
could be more important in percent and seriousness of
consequences,than other entry-operator errors and hard-
ware or communication failures. Error rates for such
type, could soar up to about 1:5 compared with typical
hardware and communication errors of 1:100,000 or entry
operator errors of 1:100. In the setting of the conven-
tional handbook one may feel that the only thing to do
is to assure adherence to managerial practices, to so-
called sound principles of system design and work,

to set up of appropriate validity checks at the input
of the system as well as adequate controls for proper
processing and check of output, to insure adequate pro-
fessional Level and training of personnel, to establish
appropriate division of responsibilities within an
adequate organizational structure, etc, It is conceiva-
ble that such set of activities will minimize all kinds
of errors, in particular source errors including those
illistrated in appendix A3 for the case study on inven-
tory differences,

An overview of the above "right" activities and pro-
cedures constitute the object of much literature on

EDP and auditing of EDP, and it was teferenced in chap-
ter 2 and app. Al, A2. The corresponding section of the
handbook may be conceived as a kind of consolidation

of such literature, e.g. G.O,Davis (1968), IBM (Porm
F20-0006), Oriicky (1969), etc. In this context it
may also be appropriate to include economic considera-
tions such as those referred by EDP Analyzer, (October
1971, p.10), in the more limited context of trade-offs
and "efficiencies" of alternative data-entry systems,
The broader economics of overall quality of information
will be considered to fall within the realm of cost-be-
nefits evaluation of the total information system par~
tially considered by Orlicky in a qualitative way (1969
p-63), and partially by Blumenthal (1969,p.144) in a
more quantitative way. Eventually, the handbook may
attempt relating the quality of information to the cost-
benefit analysis of the total information system, in
terms of the overall complete approach suggested by
Langefors (1968b,p.184). It is probable that special
developements will be required to adapt the above audi-
ting ideas, recommended EDP procedures, and economic eva
luation to the case of a data-bank which is not self-
contained and embedded in the the information system of
one only organization; this would be the case with nu-
blic data-banks.

We stop here in discussing the conventional handbook. Tt
amounts to setting up quantitative standards of error ra-
tes and qualitative procedural standards. Tt apnears

that the main scientific basis for the handbook is STA-

TISTICS as implied in the empirically determined error
rates, and in the validation of judgements on procedures,

5.6

THE "CONVENTIONAL" HANDBOOK IS NOT AN ALTERNATIVE:
THE ROLE AND LIMITATIONS OF STATISTICS,

By means of the previous section's exercise in desi-
gning a conventional handbook for quality control of
information we wanted to prepare the stage for an illus
tration of the role and limitations of statistics.

It will be recalled that we emnrehended the develop-
ment of the conventional handbook well before the dis-
cussions and conclusions in the second half of chanter
2. We shall now show that the same conclusions may

be obtained by an analysis of such a handbook; at the
same time we will show what we mentioned at the he-
ginning of chapter 2, namely that deleting of statis-
tical literature on censuses, surveys, etc, from the
review does not detract from the conclusions of that
chapter. This is particularly important for convincing
those Laymen and uncritical scientists who have a va-
gue feeling that "errors, reliability, and such" can
always be accounted for, by means of some fancy sta-
tistical analysis of "data". We hope then, that after
this section, ALL readers will be highly motivated to
make the best out of the illustrations of our tentati-
ve proposal as they will be presented in the next sec-
tion of this chapter.

An overview of the conventional handbook may be obtai-
ned by the following figure:

Figure 5.1L here

We can now ask ourselves: what is the SCT™NTIPIC basis
of such a handbook ? In other words, what is the justi-
fication for our confidence that it will "work"? As in
the case of the engineer designing a bridge, the pro-
blem is of knowing IN ADVANCE what are our chances of
success: "even a broken watch is right - twice per
day", or "if a flip a coin to determine the answer to
all my yes-no questions, TE will, after all, be right
about half the time" ! What is the basis on which

to evaluate this intuitive development of a handbook
compared with the approaches illustrated in figures
4,1 to 4.3 in the earlier chapter, in terms of preven-
tion, detection, and correction of errors ?

Looking at figure 5.1, and recalling our comments on
administration or organization theory in relation to
judgement etc., it appears that the basis for confi-
dence is to be sought in the use of statistics. We shall
therefore try to illustrate what may be said ahout the
scientific nature of statistics, and related problems.

1A 2A
[Probabi- / EDP & |
/ lity the eta

tistics

lory & Stal tific li
/ terature

itions of /
ee

|
| /Defini- /
|
i
|

: ee
Oe

j ene
Quantita~/
tive er-

fror data

ag fot 5 Sheen eed
aad

/Special- iSound prot
; purpose joedureas &

/ error- /managen.

| data [practices

‘mal qua- f{
hity hand:
i book

figure 5.1
Overview of the design of the conventional handbook 7A
in the past section,based on statistics and reviewed
EDP literature.

Walter A Shewhart was orie of the few who had to under-
stand deeply the role and limitations of statistics in
order to apply it to the practical problems of indus-
trial mass-manufacturing. In the context of discussing
the results of measurements presented as "knowledge",
he notes that the degree of belief that a scientist
holds in a prediction made upon the basis of measure-
ments of some physical constant or property DEPENDS A
LOT MORE ON THE CONSISTENCY BETWEEN RESULTS OBTALNED
UNDER SLIGHTLY DIFFERENT CONDITIONS, AND RY DIFRERENT
METHODS OF MEASUREMENT than it depends unon the numher
of repetitions made under what HE CONSIDERS TO BE THE
SAME ESSENTIAL CONDITIONS, Shewhart states also that
THE STATISTICIAN MAY COVTEHT [UTE TO THE EFFORTS OF THE
SCIENTIST IN DISCOVERING GRABLE DIFPERENCES RETYEEN

TWO OK MORE SETS OF OBSERVATIONS) "(19397 plLiZy ~~ 77777

Later, Shewhart adds! From the viewpoint of scientific
inquiry, the validity attainable in predictions depends
so much upon the skill of the experimentalist IN SELECT-
ING APPROPRIATE SENSE DATA on the one side and connect-

unless this process is carried out successfully ALMOST
NOTHING THAT THE STATISTICTAN CONTRIBUTES IS SIGNIFI-
CANT, One must not place too much reliance upon the
existence or non-existence of so-called significant

In another paper recently published, thirty years af-
ter Shewhart's warnings, R.E. Strauch discusses the
extensive abuses of techniques of statistical inferen-~-
ce caused by increasing pressure for "hard" quantita-
tive analysis in the military and civil fields such

as criminal statistics, in order to "objectively"
support "rational" policy and decision-making,
Strauch points out that statistical inference, in prin-
ciple, NEVER TNVOLVES DINUCT INFERENCE PROM THE DATA
OBSERVED TO THE PROCESS CAUSING THE DATA (a.%,from the
sample to the population in the case of sampling). It
consists, instead, of comparing the observed data with
that expected from various members of a collection of
predictive models which ARE ASSUMED TO BE ADEQUATE MO-—
DELS of possible alternative versions of the process
being observed, (Strauch, 1970) The basic principle
underlying all statistical inference, then, is that we
attempt to distinguish the process actually being ob-
served from alternative possible versions of that pro-
cess on the basis of expected differences in the out-—

An important point that Strauch makes is that the ana-
lyst in any case at least IMPLICITLY makes use of the
predictive models whenever he explicitly uses the tech-
niques of statistical inference. THE MOST SERIOUS ASPECT
of all this, however, is that the tmplicit models are
NOT self-verifying. If they were, then whenever a model

did not fit the process producing the data, this would
be evident from the data and would prevent future in-
correct inferences from being drawn. Unfortunately, this
is seldom the case. THE COLLECTION OF PREDICTIVE MODELS
CONTAINED IN THE STATISTICAL MODELS OF MANY COMMON STA-
TISTICAL PROBLEMS IS LARGE ENOUGH TO EXPLAIN ALMOST ANY

As amatter of fact, Strauch suggests to us that statis-
tical inference can also be seen in terms of what we

in this paper have called "the communication approach”
to quality of information, He reminds that given any
two of the three elements of the ideal problem, the

urn composition (balls to be drawn), the sampling pro-
cedure, and the resulting sample, it is possible to
make meaningful statements or to draw inferences about
the third, Tf we know only one of the three, however,
there is little we can say about the other two. We ask
the reader to recall our discussion of figures 2,1, 2.2
and 4.10 }

errors in the statistical inference, is the most trou-
bling in the context of our study of quality of informa-
tion. This emphasizes Churchman's statements on the
importance of having theories of factual evidence, and
on the nature of statistical tests: To test an hypothe-
sis by one or more "statistics", it is essential that

we are able to make estimates about the probability of
erroneous rejection or acceptance, and that we know HOW
LOW THE PROBABILITIES OF SUCH ERRORS SHOULD BE, The re-
quired probabilities of error turn out to be theories

in the sense that they are multiple hypotheses concer-
ning the samples that will occur under various possible
"states of Nature", (1961, p.86,168)

Against this background it makes sense indeed that a care
ful scientist as Ackoff in discussing scientific method
only takes up statistics AFTER several chanters dedica-
ted to problem definition, model building, measurement,
meaning of "optimal solutions", ete. (1962, p.218),

And that is consistent with Churchman's statement that
"The function of the statistician is not to nrovide cri-
teria for the best test, but rather to present a method
for determining the chances of error associated with
any given test, under any permissible hypothesis concer-
ning the natural world", (1948, p-283). If the reader,
then, is amazed for not finding in R.A. Fisher's "The
Design of Experiments" (1951) a complete discussion of
the limitations of statistics as suggested above, it
will be important to note together with Churchman (1948,
p.22) that Fisher's meaning of design has nothing to do
with the technique of making observations, or the formal
presuppositions we bring to bear on an experiment:

5.10

Fisher "presupposes that certain observations can be
made, that they are pertinent in general to the ques-
tion asked, and that the observations obey certain
probability laws. He then attempts to solve the statis-
tical problem: how to group the observations so that

we obtain the "maximum information" for a given number
of observations, "Maximum information" is an ambiguous
term....". Furthermore Churchman emphasizes that "..,in
order that statistical procedures be experimentally
sound, it is necessary to postulate that the statisti-
cian's hypotheses are "pertinent"; that is, we must
know why randomness can be assumed; or why a continuous
distribution function can be posited, And the answers
to these questions lie in the meaning of the original
question and the techniques for gathering data; but
this meaning and these techniques must be given within
‘a theory of the science in terms of which the original
question is posed. Hence, statistical hypotheses should
be consequences of some such theory of nature."(1948,
p.224, 218)

We feel that the above is enough for us to realize how
delicate the use of statistics really is. How many of
the statistical hypotheses tested in the literature re-
ferenced in chapter 2 and appendix AZ were "consequen-
ces of a formal theory of Nature" ? In such case were
they consequences of the physical nature or, say, of
the psychological nature ? Once again we note tHe dan-
ger of logical-positivistic influence leading us to

tie down everything to physical science. We think that
in physics it is easy to talk about "data" and to diffe-
rentiate between ohservation and other errors. But
those "data" may be submitted to statistical techniques
and disentangled from the observer only because physi-
cal science has succeded in identifying what part of
the output from instrumental observation is to be re-
garded as a description of PHYSICAL reality, indepen-
dent of the instrument and of the observer, for the
purposes to which physics is intended for, As Churchman
puts it "The disinterested observer thus becomes a de-
sign part of the system, a design based on the best
available theory of instrumentation. The effectiveness
of the design is measured by our ability to infer the
non-instrumental properties of the observing system's
output." (1968b, p.188) In our understanding the abo-
ve raises the most important questions about the appli-
cability of statistical techniques for investigating
"errors" in information systems other than those inten-
ded for the control of physical reality.

The above appears to us as being another way of apnroa-
ching the findings in chapter 3 and 4, from the view-
point of statistical theory. If statistical theory is
going to be applied to other than physical reality, then
one must consider Savage's criticism and his view of
statistics as, for example, was referred by Kaplan in
our appendix A7.This implies getting close to chapter 4,

5.11

STATEMENT OF TH PROBLEM, DEFINING THE POPULATION,
ILLUSTRATION FROM HCONOMICS

Tf, then, somebody still wants to apply statistical
methods in the analysis of "general" information sys-
tem problems, we suggest that the following seven ques-
tions be first answered (See Churchman, 1951, p.26)

1. Are you confident that the data are really pertinent
with respect to the problem ?

2. Has all pertinent information been applied to the
problem ?

3. Are the alternative hypotheses real with respect to
action ?

4, Do the data suggest any new avenues of inquiry ?

5. What statistical assumptions can legitimately he
made about the data ?

6. Is a statistical analysis necessary ?

7. How should the probability of error be set ?

We shall now go over and see how the difficulties im-
plied above practically appear in concrete situations,
i.e. in terms of difficulties at particular steps of
investigations.

We think that most of the above difficulties are hidden
in the definition or characterization of POPULATION,
OBJECT, EVENT, PROPERTIES Of ATTREGUTES, CONDITIONS,
ELEMENT, PHENOMENON, CLASSIFICATION: From this point

of view we could, for the purposes of our study, define
ERROR as an INCOMPLETTNESS of a DESCRIPTION.

Observe, for instance, that sampling may be seen as
being concerned with what subset of the set of possible
relevant observation should actually be made, when it

is not possible or practical to make ALL observations
that are ideally desirable. Which are the all possible
observations ? Observations of what ? Possible in econo-
mic or other terms ?

To express errors of estinates yielded by alternative
sample designs it is, among other things, necessary to
know a great deal about the distribution of the vroper-
ty in question among the elements of the population to
be sampled. How much can be known ? What has to be as-
sumed ?

. In order to determine the nature of observer errors,

it is necessary to know a lot about the nature of the
object or event observed. Vhenever the "true value" is
not known, observers are usually checked by using a
standard object or event under specified conditions.
What is the basis for assuming such a true value ?

If the thing observed is destroyed or significantly
changed with respect of the relevant property by the
observation process,then the method with the standard
cannot be used. How to determine whether a change was
significant ? What to do in such a case ? In snite of
all the doubts the discussions about observational

ies
e
pel

versus sampling errors is, in statistics, usually done
in terms of an assumed well defined population of ele-
ments having a particular well defined and measurable
property that is to be estimated, and it may be assumed
that the true values of the elements' properties are
normally distributed; etc. The assumptions are common
to the related discussion of bias. A general feature of
the discussions is the acceptance of indiscutable ob-
jects and attributes. The possession of an attribute
such as blue-eyedness might, however, present the same
difficulties that were suggested for the determination
of red color, in case the life of a person would depend
on such a determination (recall chapter 4),

No, the question of defining objects and attributes

is by no means simple, and it is a basic scientific
problem prior to any statistical computations. Consider
for example what Ackoff, who also discuss many of the
above questions,says on the concent of "object" that
was made necessary in quantum mechanics: "This seems
to offend our feeling that all "objects" can be loca-
ted at some specific place at some specific time. But
the new physics requires that we reinterpret the con-
cept Yobject" in terms dealing with the way it is
obsetved. In effect, an object in the new mechanics is
a "state of nature" which is described statistically;
it is not a ‘particle of matter." (1962, p.210)

The above makes us understand why the "object" having
"attributes" in, say, a public data~bank is perhaps not
at all properly characterized and identified by means
of only the name, birth date, and social-security num-
ber. Compare what Ackoff said above with the following:
"What is needed is a system of legal controls, so that
the user of the (information) center cannot simply re-
trieve the datum "Jones was convicted of burglary." The
information, instead, would contain something like an
abbreviated model of Jone's life, so that one under-
stands the implications of the assertion about the con-
viction relative to decision making." (Churchman, 1968b,
p.196). What this implies is the need of redefining the
concept of "person" in the context of public data-banks
and social decision-making.

It is interesting to note that such need is really com-
mon-plLace in the context of modern manufacturing of
technically advanced products, Such manufacturing re-
quires that the final-assembly be described in terms

of a breakdown, a "bill-of-material" structure of sub-
assemblies and components, where each sub-assembly or
component part at each level is identified by a part
number PLUS AN "ENGINEERING CHANGE" NUMBER providing a
cross-reference to engineering documentation that des-
eribes the "story" of the changes to the drawing. When-
ever a decision affecting a part is of any importance,
it is necessary to have both the part number and the
latest engineering-change number that affected the part.
The data-files are often designed to provide and to vro-

5.13

cess both simultaneously. People working with the con-
cepts often require that the "part-number" concept be
enlarged in some way to include the "engineering-chan-
ge number" concept resulting in a kind of composite
identification number that changes with the course of
events.

From a scientific point of view, therefore, it annears
dangerously naive and unjustified to expect that data-
banks can be developed and operated in the much more
delicate context of social systems, without having
submitted the whole problem of object, attributes, etc,
to an exhausting analysis.

Continuing our review of difficulties in concrete si-
tuations we may recall the problems of definition, and
classification that we met in the coritext of chapter 2
and appendix A2. It is obvious that we can barely expe-
ct to be able to consolidate most of the reviewed re-
search to the extent that its hypotheses were not the
results of some formal theories or to the extent that
the information system itself does not represent a for-
mal theory of the controlled system; such as the case
was for the quality control of manufacturing. One can-
not just go on creating "concepts" such as CHARACTERS
or RESIDUAL ERROKS for every particular investigation
and then expect that they will be integrated in an over-
all"theory" for a general information system. Maybe the
nature itself of information systems is such as to pre-
vent a meaningful discussion of errors in these terms,
and this can be one of the imnlications of our proposal
in chapter 4.

Next, statistics in economics also shows many of the
basic difficulties and limitations of statistical meth-
ods. Morgenstern presents many examples which
may be perfect analogies of troubles to be met in fu-
ture complex data-banks and information systems. Dis-
ecrepancies between revorts of the same event are not
considered "errors" in the statistical sense, but are
merely differences in definition - differences in em-
phasis in which components of a statistics are imnpor-
tant. One is therefore faced with alternative sets of
data which aim to describe the same phenomenon but which
appear quite different, One has to deal with incompara-~
bility due to definitional kinds of errors which are un-
known to physicists who work with carefully defined
terms in a field where there cannot be alternative non-
equivalent descriptions of the same phenomenon,

And that is the result of Lack of thenry,where border~

line cases occur which do not fit properly in a parti-

cular category (recall chapter 3) because of changes in
the property of the object measured, In census of manu-
facturers uncertainties of classification may arise

Rides

5.14

because of the appearance of new commodities, new in-
dustries, because of changes in the quality and appea-
rance of products. The difficulties are comnounded
when some widely used statistics are produced by means
of an inappropriate procedure, neglecting the change
in the framework into which the concepts must be em-
pedded. For those who are more familiar with physics,
it is easy to be misled by the fact that physical nro-
cesses not only have more "stability" (e.g. astronomy)
but also the classification of phenomena is much less
in doubt thanks to a well developed instrumentation
and theory.

Morgenstern (1963, p.92) raises an extremely important
point, when he emphasizes that the quality of the data
themselves on the basis of which econometric models
are established, may preclude the successful testing
and improvement of such models, Neither changes of
parameters nor inclusion of "not earlier considered
hidden variables" with the help of sensitivity analy-
sis, fancy statistical techniques, or sheer intuition,
will substitute a scientific analysis of the nature

of used basic data. The word "randomness" should not
be used, but rather the concept of error should be
applied in the build-up of theories which separate er-
rors of observation from failure to account for factors
which should enter in the models. This appears to be
consistent with earlier material in this chapter and
with the spirit of our chapter 4.

Another very important point that Morgenstern raises

is the increased indeterminacy and vagueness of measu-
rement of a concept in pace with its increased scope

of application or importance (1963, p.44). It is appa-
rent that the statistics dealing with an object ina
very varied and illdefined environment or conditions
must to an increasing degree "sample" the relevant ele-
ments with the relevant attributes in the relevant con-
ditions, for some purpose, The case was made concrete
in chapter 4 when discussing the case of the determina~
tion of red color, of the birth-date, or of the true
stock level in the case study of appendix A2, This may
be a new way of conceiving the difficulties in measu-~
ring final or high goals: the "state" of the nation's
economy, as well as its correlate the "“eoal" of the
economy cannot be described or measured because they
are indeed attributes of the concept ~ object "economy"
which is so complex and broad in its scope, The "con-
cept" then gots indoterminate, and its attributes as
well, invalidating any talk about a statistical approach
to its measurement.

In the context of the last paragraphs we shall also men-
tion that the so~called Bayesian attitute towards facts
and information systems as for instance advanced by

J. Marschak (1959, 1964), and by J.C, Emery must meet
all the objections implicit above, and in the referen-
ced literature, In particular,the approaches by Mar-

f

schak and Emery assume a set of all possible "states of
nature" - external and internal environment, assume in
the argumentation the existence of "faults" in the des-
cription of "actual" states of nature, and assume pro-
babilties being assigned to "events" and to the "outco-
mes" of the actions of "consistent ~- rational" men.
Bayesian thinking then comes into the picture in the
context that the receipt of a message may alter the de-
cision-maker's "view of the world" and cause him to
revise his estimates of state probabilities.

To the extent that, as Marschak suggests (1964, p, 38),
such foundations are considered to be relevant to the
future of macro~economics of information seen as an
extension of the theory of welfare economics, or public
policy, we would like to add our objections to those
expressed by Churchman (1961, p.167,1968b, p.100). The
reader is urged to note that these are serious matters:
Marschak suggests attempting 't.to characterize a social-
ly optimal allocation of channels, given the distribu-
tion of tastes and beliefs, and given the society's to-
tal resources and their initial distribution." And this
is far indeed from Pmery's illustrative example of ap-
plication of Marschak's concepts to defective pieces

in a manufacturing environment, where he concludes that
"Quite apart from any theoretical limitations of the
model, it is obviously difficult to apply it in prac-
tice... Nevertheless, a theoretical discussion of the
value of information has considerable usefulness, First
of all, a substantial formalization is now possible, par
ticularly in lower-level processes that deal with routi-~
ne operations." (Smery, 1969, p.90)

We agree, then, than non-problematic apnlication of
statistics, probabilities, and simple concepts is possi-
ble when a good theory exists, such as in physical manu-
facturing, or when the importance of applying the con-
cepts is little or none (routine applications). 2ut

not further: a completely different apvroach may be re-
quired. If we do not do this,it may well happen in the
above Bayesian apolications, as well as in the milita-
ry applications suggested by W.Edwards et al. (1968)
which were referenced in appendix Al, that we fulfil
the prophecy implicit in another statement by Church—
man: ",.,.the basis for a decision about the "next event’
may vory well have been already inherently established
in decisions about the relevance and accuracy of the
data." (1961, p.167). Kecall also our reference to the
problem of forecasting sales, based on past sales ver-
sus based on analysis of causes and nature of sales,

in chapter 4: if one just STANTS with the registered
past sales as "facts" then the problem may turn out

to be just to develop a forecast formula based on the
best available statistical techniques !

as

5.16

CENSUSES AND SURVEYS, STATISTICAL INTERVALS,
"REJECTION OF OUTLIERS", AND HISTORICAL RESRARCH,

Next, we can observe the symptoms of the limitations
of statistical methods also in the context of censuses
and surveys. A paper by M.H, Hansen et al. (1961)
shows that the obtained observations refer to attribu-
tes such as age, income, but also other more vague
characteristics such as buying performance and attitu-
de on a particular question, Such characteristics are
regarded as belonging to “objects" such as a person,
household, farm, business, area, or other "unit",

The "true" value of the statistics is idealized as
being that proportion of the population of elements,
having some "valuc" which represents a specified cha-
racteristic. In order to insure ADEQUATE QUALITY of
the estimates it is necessary to attempt to impose
such"conditions"(under the control of the survey de-
signer or sponsor) that"specify various aspects" of
the conduct of the survey. Some examples of conditions
under which the samples may be taken are questionnaire
design, publicity in connection with the survey, the
type of organization and job assignments in connection
with the survey, qualifications and training of the per
sonnel to be selected, pay system, inspection and con-
trol procedures,

In the text of the referenced paper we could find the
following three statements which we feel are symntoma-
tic for the purposes of our study.

"We...shall use the root mean square error of any esti-
mate as a measure of its accuracy. Although in practi-
ce we cannot know the...mean square error of ...(the
estimate) , we may be able to obtain an approximation
or a useful over-estimate or under-estimate."(p.361)

"There are a number of ways of designing experiments

to obtain approximate estimates of the response varian-
ce or of specified components of the response variance,
although we know of no way of obtaining unbiased or
consistent estimates of them." (p.367)

"We have no reasonably satisfactory approach for mea-
surement of response bias, although there are some
helpful methods," (p. 370)

in the course of developing the last citation above,
the authors explain the following. "The monthly Cur-
rent Population Survey (cPs) taken by the Bureau of the
Census is carried out under much more rigorous controls
than is feasible for the complete decennial census, and
there are reasons ta believe (and the Census Bureau has
adopted this position) that the results of the CPS are
more nearly accurate on the average, than those of the
census. Consequently, approximate measures of resnonse
bias in the census are obtained by using the CPS measu-
rements as standard" (p.372)

5417

We see, then, that the reviewed most refined sta-
tistical techniques as they are used in official sur-
veys and censuses, make recourse to vague conditions,
reasonably satisfactory approximations, helpful methods,
and eventual comparison against a standard. We are

thus back to chapter 2 and chanter 4: what is done

may also be seen in terms of the communication approa-
ch to quality of information, to the extent that some-~
body, who "knows" and has authnrity, tells us which is
the "right" procedure or program to be followed. The
problem is then that the right procedure cannot be
enforced on a large scale because for instance the in-
terviewers introduce the "bias" of their own judgements
and therefore such response deviations must be detected
by means of comparison with a more structured situa~
tion, the standard situation (as the CPS above) where
it is possible to enforce the only authorized, expert
judgements. This leads us back to chapter 4, and our
struggle to disentangie the origins and the systematic
evaluation of judgements.

Next, against the background of so many conceptual
difficulties, we should not get surprised about the
unclear meaning of the concepts of accuracy, precision,
confidence intervals, tolerance intervals, etc. as
used in many statistical investigations. In the same
way as precision and accuracy are often vaguely asso-
ciated with sampling and respectively observation
errors (to be detected and corrected through compari-
sons with the standard, such as detailed interviews
in depth), both tolerance and confidence are associa-
ted with truth.

What is often not realized is that confidence inter-
vals, such as the Student range discussed by Shewhart
(1939, p.97) tell only to us the probability that a
certain range of numbers constructed out of cbserva-
tions on one same well defined population, will inclu-
de the "true" value. On the other hand, if a system

is known to have been in control, the tolerance Limits
tell us the probability of making an error of a cer-
tain magnitude, that is of deviating from the true
measurement by a snecific amount. In neither case it
is purely statistical problem for the decision maker
to see how he can use the confidence and talerance ran-
ges resulting from a statistical investigation. (See
also Churchman, 1961,p.128). This was also seen in the
context of chapter 4, and appendix A5.

In the course of illustrating the role and limitation
of statistics, we shall next refer the reader to annen-
dix AQ where we made an overview presentation of what
statisticians say about a particular problem: rejection
of outliers, As we have earlier seen in this paper, and
as can be inferred for example from the paper by Hansen

5.18

et al, (1961), repeatability is a basic requirement in
many experimental approaches to truth. How do statis-

ticians proceed when one value obtained by a particu-

lar measurement process of a supposedly constant mag-

nitude turns out to deviate "too much" from the other

values in a series of repeated measurements ?

The appendix is, after our discussions, self—explana-~
tory. It is interesting to note that suddenly new con-
cepts appear in the context of statistical investiga-
tions: inherent variability, execution error (recall
our "source" errors and appendix A3, The basic cri-
teria for rejection of deviating observations is said
to depend on the purposes of the investigation and on
the nature of the statistical material, and eventual-
ly an approach is suggested that in much reminds Chur-
chman's seven questions to be answered before initia-~
ting a statistical irvestigation. It appears to us
obvious that statisticians recur ih these cases to
discussing the basic problems of scientific method and
theory of science. But this eorrespondence appears to
be seldom recognized,

We feel that it is remarkable that statisticians do not
explicitly seem to reengnize that an enlabgement of

the scope of statistical applications, encompassing
more and more of social and nsychological phenomena,
amounts to turning statistics into sheer scientific
method, When reviewing much of the statistically orien-
ted literature, however, we folt that a picture was
growing into us, conveyed by the literature, and which
may be summarized in the following terms:

"What we need is well-developed techniques for put-
ting together into a meaningful and objective pictu-
re the items of information contained in various com-
ponents of knowledge and observations. We need a uni-
versal statistical error-theory which supplies us
with quantitative estimates of error in any field of
application, in order to prevent the effects of
misunderstandings, carelessness, and of people intro-
ducing their own judgements in the context, for in-
stance,of interviewing somebody for the purneses of

a survey. Such a statistical theory would allow,

for example, to recognize the direction and extent

of wilful distortion of information and to eliminate
its influence."

The reader should note the important implications of
Morgenstern's statement about problems "...in a Large
povulation sampling with living beings having attribu-
tes that are difficult to describe and often not wan-
ted by those questioned..." (1963, p-218) Observe the
implications if somebody qualified slightly the state-
ment as follows "...with living beings to whom somebo-
dy has assigned attributes which are not wanted by the
questioned since they have motives to expect that such

5.19

attributes will be used against what they consider as
their legitimate interests...". Or, consider the im-
plications of stating that interviewers (and inter-
viewed 1) also have legitimate judgements that per-

or ilegitimate judgements of the sponsor or of the
designer of the survey ! Refer also to Morgenstern's
comments on the relation between the concepts of "lies"
versus "wrong judgements" (1963, p. 25,81) and see their
applicability in analyzing lies of respondents versus
judgements of sponsors of surveys.

Next, we shall finally explore whether all the above
problems do not, as they intuitively should, appear in
the context of historic research. If a nuclear war
erased several nations from the face of the earth and
left just a few well protected data-banks, how would
survivors proceed in order to infer about the vast ?
It is obvious that such a question may be relevant for
our study of quality of information, We prepared, the-
refore, ap QO which in our opinion clearly
shows the 1 difficulties being multiplied in
such complex context, There appear a host of poorly
defined concepts such as consistency, relevance, cre-
dibility, fitness for use ete,

Furthermore,the overview supports many of the findings
presented by Morgenstern, who in fact covered also si-
milar material to the contained in the historical case
studies. A deep analysis of the material would proba-
biy help in predicting analog problems or errors that
will appear in future ambitious information systems,
especially in connection with the concept of genesis: -
original data, raw material, primary versus secondary
statistics, first versus second-hand source, and cre-
dibility.

Since the referenced work by Schiller & Odén is writ-
ten in swedish, our readers may find an excellent al-
ternative in S.Rokkan et al, (1969) where interested
researchers can read $.Verba's contribution on "The
Uses of Survey Research in the Study of Comparative
Politics." In our opinion, Verba succeeds in covering
many of the deep and complex problems which were not
considered in another book by R. Naroll on reliabi-
Lity of ethnographic data, with the rather misleading
title “Data Quality Control - A New Research Techni-
que", (Naroll, 1962). Naroll, however, also presents
some interesting case studies.

In the context of accuracy of measurements, Verba
talks about problems of comparability in multi-contex-
tual research, and he differentiates the technical
problem of measurement from problems of so-called con-
ceptualization, Comparisons based on survey research
MUST take into account the so-called context (social

Yu

Rico:

5.20

structure and culture) within which the individual me-
asurements were taken. Only then can one talk on ac~
curate information and meaningful information within
different social settings, and compare the same "thing"
word, act or attitudes with the same "label", for e-
xample "votes", Norimes", "suicides" or in general
"answers to the same question?

Ways in which context of the individual measure can

be taken into account is, for example, by means of
proper selection of variables, or by breaking them
into component parts (disaggregate them) and there

one meets the all-important problem of objective ver-
sus subjective definition of terms. The problem turns
then out to be HOW tn disaggregate. What is campared

is not the absolute frequencies of attributes, say
voting, between two systems, nor even between compa-
rable subgroups in two systems. One rather compares
systems in terms of ways in which voting rates DIFFER
among subgroups within the several systems. In this
way statistics applied to historic research attempts to
obviate the problem presented by the insight that the
"fact" that an individual voted can mean at least

five different things (and some more may be immagined).
(See Verba on voting, 1969)p.70)

The work of Morgenstern, Schiller & Odén, and Verba
exemplify the enormous complexity of the error con-
cept. We feel that it must, at the general level, be
analyzed in terms of scientific method, and not by
piccemeal attacks on "source" errors whose high rates
and magnitudes may rather express the inadequacy of
statistical methods, and not any increased understan-
ding of the nature of errors and of the system, or of
statistics itself. It is then unfortunate that histo-
rical statistics also appears divorced from scientific
method: "The decision for accepting facts about the
past is based on a predictive theory about the futu-
re, for example, repetition of the same observer re-
ports in various circumstances..... the theory that
underlies a fact also predicts the future; it predicts
continuing acceptance of the evidence, for example,"
(Churehman, 1961, p.167). We feel, therefore, that it
may be fruitful to relate our study to historical re-
search. Some direct implications may be derived, e.g.
in relation to coding in content analysis, as touched
upon e.g. by S.Rokkan in the mentioned work (Rokkan
et al. 1969): coding could obviously be seen in terms
of some functional definition of measurement
(Churchman, 1961,p.93). See also Ackoff (1962,p.174).

SUMMARY ON THE ROLE AND LIMITATIONS OF STATISTICS

We conclude that a conventional handbook for quality
control of information is not really an: alternative

to a handbook based on our approach in chapter t, It
does not appear meaningful to ciscuss errors on the ba-
sis of statistics alone. Therefore we,are not able to

Seek

utilize the findings reviewed in chapter 2, nor to
implement the idea of figure 5.1. All this may also
explain why we were not able to find any statistical
approach to the overall problem of quality of informa-
tion in data-banks, in the context of the literature
reviewed in chapters 1 and 2, and appendixes Al and A2,.

As Churchman expresses it (1970, p.B-41):

"Though it is obviously difficult to assess the se-
riousness of ignoring the systemic judgement impli-
eit in operations-research data, I'd estimate that
it is a far more serious error than the typical
errors associated with statistical analysis to which
formal education does devote a great deal of its
time. IT IS TO 3K NOTED THAT THE PROBLEM OF THE
CORRECT SYSTEMIC JUDGEMENT IS NOT HANDLED BY STATIS-
TICAL THEORY, WHICH, IN EFFECT, PRESUPPOSES THAT IT
HAS BEEN SOLVED." (Our emphasis)

Tenoring the problem of system

<!-- Book text truncated by scrapem max_book_chars. -->

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
