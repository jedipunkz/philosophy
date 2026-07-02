---
source: "https://archive.org/details/IJLA008"
title: "Relating Plato to the Pythagoreans-The Development of a Software System to Explore the Influence of the Pythagoreans on Plato’s Work"
author: ""
year: ""
captured_at: "2026-07-02T01:10:27Z"
updated_at: "2026-07-02T01:10:27Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato Symposium"
plain_text_url: "https://archive.org/download/IJLA008/IJLA008_djvu.txt"
public_domain: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# Relating Plato to the Pythagoreans-The Development of a Software System to Explore the Influence of the Pythagoreans on Plato’s Work

- 著者: -
- 情報源: [archive](https://archive.org/details/IJLA008)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

www.seipub.org/ijla

Relating Plato to the Pythagoreans

The Development of a Software System to Explore the Influence of the
Pythagoreans on Plato's Work

Evangelos C. Papakitsos* 1 , Maria I. Tsoumaki 2

Department of Linguistics, Faculty of Philology, National & Kapodistrian University of Athens, Greece
Panepistimioupolis, 157 84 Zografou, Athens, Greece

^papakitsevOsch.gr; 2 mtsoumaki@hotmail.com

Abstract

This paper gives a brief description of a software system
which has been developed in order to facilitate the research
that aims to reveal the relation between Plato and the
Pythagoreans. There are evidence presented supporting the
idea that the influence of the latter on the renowned
philosopher impelled him to give a mathematical structure
to his writings. The effort to explore such an underlying
mathematical structure is complemented in many ways,
with the assistance of a computerized tool which is briefly
described here, along with some exemplary results from
demonstrating the proposed methods of usage to a famous
work of Plato, the Symposium.

Keywords

Plato; Pythagoreans; Symposium; Computational Stichometry
Introduction

The influence of the Pythagoreans in thinking and
philosophical theory of Plato has been considered
particularly catalytic, already since the antiquity, in
such a degree that both Aristotle and other
contemporaries labelled him as basically a Pythagorean
philosopher.

Plato's acquaintance with the Pythagoreans began in
Taranto, Italy, in 387 BC, during the trips made by the
philosopher after the death of Socrates. Diogenes
Laertius (1994) specifically mentioned Plato's
acquaintance with Filolaus and Euritus, while in the
Seventh Letter, attributed to Plato himself, we read
about the close friendship developed between himself
and Archytas, a still prominent Pythagorean, who has
even been argued to be a "new paradigm philosopher for
Plato" (Vlastos 1991) after Socrates.

The Pythagorean School asserted that the numbers are
the components of the substance of the universe. For
the Pythagorean philosophers, the numbers, the
mathematical relationships and the geometric shapes
that are depicted in paper, are the representation of
their original models which constitute unique and

immutable standards in the human mind. This is the
famous Pythagorean teaching of "imitation",
according to the saying that everything noticeable is
an imperfect imitation of the perfect imaginary
world. The influence of this theory in the mind of
Plato, for the subsequent development of his own
theory about the world of ideas, is obvious. In a
further analysis of the Pythagorean philosophy, the
basis of this is the so-called "Tetraktys" (see Fig.
1). This is the sum of the first four natural numbers (1
+ 2 + 3 + 4 = 10), by which the Pythagoreans argued
that the ratios of 4th, 5th and 8th harmonic could be
constructed, which in turn created the so-called
"Harmony of the Spheres", a tune that governed the
universe and could be understood only by
philosophers (Kennedy 2011).

AAA

FIG. 1 TETRACTYS

Due to these fundamental theories in the philosophy
of the Pythagoreans, it has been argued that they gave
a mathematical structure in their writings, as it was
believed by Vitruvius. Thus, there are many scholars,
so far, who have investigated the potential impact of
Pythagorean theories even in the way in which Plato
would structure his works, seeking some latent
mathematical organization and musicality in his
dialogues.

The most recent study of stichometry about Plato was
conducted by J. B. Kennedy (2011, 2010), who,
considering the dialogues Symposium and Efthyfron,
claims to reveal the musical structure that governs
these dialogues. Specifically, the stichometric analysis
of Kennedy reveals that every dialogue can he divided

7

www.seipub.org/ijla

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

into twelve -parts, each of which is a symbolic representation
of a musical scale of twelve notes (Kennedy 2011).

All the above, combined with the widespread practice
of stichometry by the same authors and still for
practical purposes, since the period of antiquity as
shown by Ohly (1928), was the impetus for the
preparation of this paper. The aim is to further
investigate the possibility of mathematical
organization of the Symposium of Plato, and by
extension of the Pythagorean influence, with a method
of automated stichometry, after producing a relevant
computational tool in the programming environment
of Visual C# computer language (see Foxall 2008).

Given the large extent of these texts needed to be
verified in every such study, the manual method of
stichometry, which has been applied until today (with
the exception of Kennedy's work, see Kennedy 2010,
2011), namely, the counting of syllables in order to
make the arrangement of the initial text in lines of a
predetermined number of syllables, as in this case,
apart from being very time consuming, poses a high
probability of error. Thus, it seemed appropriate to
automate this process through the computational
implementation of a kind of syllabifier, which will
provide the user with the ability to make immediate
and effective arrangement of the original text into lines
of a specified number of syllables, through a user-
friendly graphical user interface (GUI), and then proceed
to conclusions thereon. The core of this computational
implementation was based on a number of previous
relevant works (Papakitsos 2011, 2009, 2000, 1992). The
final software tool has been tested for its overall
efficiency to the Symposium of Plato, as the "target"
text. The results of this application will be presented,
later on, in this paper.

The Target-Text

Looking at platonic dialogues from a literary-
philological point of view, we can distinguish two
major categories:

The first one includes dialogues, in which the
discussion is transferred directly from the
participants-protagonists themselves (see Bury R.G.),
talking successively (usually Socrates and whoever his
main interlocutor is, as in Phaedrus, Cratylus, Efthyfron,
etc). In this case, each time, the speaker's name
precedes his speech.

In contrast, the second category includes those
dialogues in which the debate occurs indirectly by a

narrator, who either participated himself in the
discussion, as attending, or not. So, if the narrator was
not an eye witness (or even better a hearsay witness)
of those dialogues, this kind of narration includes an
additional introductory and a concluding explanatory
paragraph. Dialogues like Theaetetus, Parmenides and
of course the Symposium belong to this category,
whose narrative levels of the latter are presented in
Table 1 (Tzouma 1997).

TABLE 1 THE MULTILEVEL NARRATION OF THE SYMPOSIUM

Level of Narration

Narrator

Auditor

Outer-narrative

ApoUodorus

Hetaerus

Inner-narrative

Aristodemus

ApoUodorus

Post-narrative

SYMPOSIUM IN THE HOUSE OF

AGATHON:
Phaedrus, Pausanias, Eryximachus,
Aristophanes, Agathon, Alcibiades
Socrates (including Diotima)

Thus, of particular concern at this point is to look into
the structure of the Symposium that causes the
specificity of this dialogue about the narrative
technique used by Plato to present the speeches of the
different participants. To understand what will follow,
the external structure of the Symposium should be
described, at least broadly.

The structure of the dialogue is masterfully designed.
The narration, being layered so to speak, is
constructed essentially based on a pair of question-
and-answer. Consequently, the prologue begins "in
medias res", in the middle of a conversation, after a
request that has already been preceded by a company
of friends of ApoUodorus, to narrate the discussions
held in the house of Agathon, about love. Although
ApoUodorus was not himself present at the
symposium, Aristodemus, who had attended as a
listener, had narrated in details this event to him
before. So ApoUodorus accepted willingly to transfer
the conversation to his friends, just like he had
recently done to Glaucon, when he met the latter on
the road from Faliro to Athens. Below are the speeches
of all those who attended the symposium:

a) The speech of Phaedrus,

b) of Pausanias,

c) of Eryximachus,

d) of Aristophanes,

e) of Agathon,

f) the speech of Socrates (which includes the
words of Diotima as well) and

g) the praises of Socrates by Alcibiades.

8

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

www.seipub.org/ijla

The text displays a great complexity in terms of
narrative techniques, as shown in Table 1 above,
which may hinder the attempt of exploring some of
the variations of the text that we thought possible at
the computational stage of analysis.

Related Work

The research about the quantitative text analysis of
Plato's work has started in the 19 th century. Nowadays,
before presenting the results obtained from the
arrangement of the text of the Symposium in lines of
specified number of syllables, using the software
developed for this purpose (Tsoumaki 2012), it would
be appropriate to refer to the most recent research
conducted by J. B. Kennedy (2011, 2010), in an effort to
reveal a possible stichometric organization of Plato's
dialogues.

Kennedy argued that not only the extent of the
speeches in the dialogues of Plato and their position
within the current text, but also the location of
important turns in the arguments of speakers reveals
an underlying stichometric organization. In particular,
he stressed the importance of a 12-parts structure that
the Symposium and other dialogues of Plato are shown,
as he claimed by even presenting particularly
convincing evidence.

In details, a plea of Kennedy is that in dialogues like
Menexenus, Phaedrus and of course the Symposium, the
speeches of the persons involved are clearly
demarcated from the rest of the text. Considering in
particular the Symposium, he noted that the speeches of
Pausanias, Eryximachus and Aristophanes occupy -
each one - the 1/12 of the total size of the text, while
the speech of Socrates (including the discussions he
had with Agathon and Diotima) and that of Alcibiades
constitute 3/12 and 2/12 respectively. Thus, he
concluded that a ratio of 1:12 was probably of main
importance in this work. This is also evidenced by the
relative position of these speeches in the text, since for
example the beginning of the speech of Pausanias is
placed at the 2 nd /12 of the dialogue, the beginning of
the words of Eryximachus at the 3 rd /12 and of
Aristophanes at the 4 fh /12. What is noteworthy is also
the fact that the rhetorical fireworks with the praises of
Eros within the speech of Agathon, occurs in the 6 th /12
of the dialogue, perhaps not a coincidence, since love
is the central theme of the Symposium.

Similar regularities seem to arise from the comparative
study of Plato's dialogues, concerning the relative
position within the text, of ideas with positive or

negative content. For example, negative ideas like
dishonesty, disease, denial, etc. tend to have a certain
size and are usually placed between the 10 th /12 and the
ll th /12 of the total text. In contrast, positive ideas such
as virtue, justice, and kindness are shown roughly
between the 8 th /12 and the 9 th /12 of the text.

Besides all the above, the main motivation for the
preparation of this work is the assumption (Ohly 1928)
that the extent of classical narrative compositions is
usually measured by the number of "defined" lines,
each of which has a length equivalent to that of a verse
of the epic poetry, written in Hexameter. Herein our
research work starts. This is because, as it was said, an
investigation of possible Pythagorean structure at the
Symposium will take place after the arrangement of the
text in lines of 12 to 17 syllables. The choice of these
numbers, however, is not random but directly related
to the number of syllables that a verse in Hexameter
potentially has (see Table 2).

The Hexameter consists of six feet, each of which has
duration of two times (Papakitsos 2011, Halporn,
Ostwald & Rosenmeyer 1980). These times are given
either by a long syllable and two short-ones (dactyl:
dactylic foot) or by two long ones (spondee: spondaic
foot). The last foot, however, is always a spondee,
consisting of two long syllables. From the above, it is
calculated that the minimum number of syllables of a
verse in Hexameter is 12, while the maximum is 17
(see Table 2).

TABLE 2 THE RANGE OF METRICAL PATTERN OF HEXAMETER

Order of Foot

Type of Foot

1st

Spondaic

Dactylic

2nd

Spondaic

Dactylic

3rd

Spondaic

Dactylic

4th

Spondaic

Dactylic

5th

Spondaic

Dactylic

6th

Spondaic

Spondaic

12

17

Total of Syllables

(=6x2)

(= [5x3] +2)

Minimum

Maximum

Data and Preparation

Given our interest in the spatial organization of the
ancient scrolls, other parameters in the context of this
analysis should be explored, which may lead to
alternative forms of the original version of this
Platonic work. As claimed elsewhere (Kennedy 2010),
the ancient scribes were paid per line and their wages
could even be legally safeguarded. Also the cost of a
papyrus was partly proportional to its length, while
some scholars argued that the ancient literary papyri
were of predetermined length and that for this reason

9

www.seipub.org/ijla

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

the authors premeditated their compositions to match
the scrolls, thus avoiding large gaps (W.A. Johnson
did not find evidence for such a tendency in the later
papyri of Oxyrhynchus, see Johnson 2004). Therefore,
even the way of declaring successive speakers in the
Symposium is a factor that could affect the final
outcome, in terms of overall layout of the papyrus.
The exploration of these parameters will be attempted
next.

The digital variant of the Symposium, as the source of
the present work, is the text of the Perseus Digital
Library, being accessible to everyone through internet
(www.perseus.tufts.edu/hopper/). The first intervention
made is the conversion of the polytonal system of
writing characters to mono-tonal, for the needs of the
software application. This should be done in order to
reduce the complexity of the program. This variant,
when used without any other changes, includes
additional names of speakers in their full within curly
brackets (e.g. {Apollodorus}). Also the passages are
declared in square brackets (e.g. [172a]). As
understood, because the declaration of the passages
contains only numbers and letters of the Latin
alphabet, it does not affect the counting of the syllables
and so it is not necessary to remove these statements
from the text. Regarding the names in curly brackets
though, they had to be erased as not being part of the
original text. After all, by combining two narrative
techniques in the text (Tzouma 1997), that of 3 rd -
person narration ("... Aristophanes said ...") and that
of dialogue ("Well ... Eryximachus indeed ..."), Plato
managed both to inform us about the identity of the
speaker and to retain the vitality and immediacy of the
narration.

TABLE 3 SYLLABIC PATTERNS OF THE ALCIBIADES PAPYRUS

Line

Line

Number of syllables per line

order
without
change
of

Number
of

syllables
per line

order
with
change
of

not including

gaps or
abbreviated
names

aggregation
of gaps and
abbreviated
names

speaker

speaker

2nd

15

5th

10

12

3rd

15

6th

11

13

7th

14

8th

11

15

9th

15

10th

11

15

11th

16

13th

13

15

12th

13

14th

12

14

16th

15

15th

10

14

17th

15

The final issue to be addressed is the number of
syllables per line of papyri, concerning also the change
of speaker. Because the Symposium is not particularly
helpful in this respect, another work has been

examined, that of Alcibiades, which will give us an
initial indication. The objective of the count is to
discover any kind of regularity in the number of
syllables per line, whether the lines contain a change
of speaker or not. From the sample data of Table 3, it
becomes apparent that the particular papyrus does not
follow a strict compliance to a specified number of
syllables per line. However, it has been observed that
the lines not containing a change of speaker have a
range of 13 to 16 syllables.

Respectively, it can be seen that the lines containing a
change of speaker via a space gap (or an abbreviated
name) exhibit a range of 10 to 13 syllables. Conversely,
if the count is included in the space gaps as
conventional syllables, then the total number of
syllables per line comes back again in a range of 12 to
15 syllables. Therefore, it may be concluded that the
range of 13 to 16 syllables is the standard pattern that
we shall explore further, which is also compatible to
the standard pattern of the Hexameter.

The Software Tool

This entire project has been initiated by an original
idea of J. B. Kennedy, who suggested the development
of such a software system, in order to explore further
the issue of Pythagorean structure in the works of
Plato. It may be proved whether some ancient
philosophers counted the number of syllables per line
or not, in order to give a mathematical structure in
their work.

The software requirements, as defined by the user of
this computational system (Tsoumaki 2012), are listed
below:

1) The application should be compatible with the
operating system of Windows (versions XP,
Vista, 7).

2) The recovery of the digital file of the text of the
Symposium will be from the hard disk of a
user's computer.

3) The users should be able to choose the number
of syllables per line in which they wish to split
each text.

4) The tool should perform the arrangement of an
original text in lines of 12 to 17 syllables and
should be able to display the results both on
the application window and in a text file that is
stored on the hard disk of the user.

5) To facilitate inference on the results, it would
be very useful to enumerate the lines of the
edited text. In this way, every line can be

10

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

www.seipub.org/ijla

absolutely referenced.

Following the previous requirements (1-5), the
computer program developed in this work receives a
text-file as input (see Fig. 2) plus the desired number
of syllables per line of the arranged text after
processing.

C,djoiKai%eXevTr]oaoiv. [180c] OcuSqov \xev toioutov

Tivd Aoyov ecfin Eircdv, lietcx be Ocu&qov aAAou? Tivd? elvai,
cov ov navv &i£Lrvrm6v£i> ev ovc, nageic, tov Ylavaaviov
Aoyov 5i.r|Y£LTO. eitieIv 5' ainov oti- ov KaAcoc. liol 6okei, a)
OcuSqe, 7iQo|3£|3Afja9air||-UV o Aoyoc., to anAojc,

FIG. 2 THE LAYOUT OF THE INPUT FILE

The output of the program is another text-file,
containing the original text of the input-file, arranged
in enumerated lines of the predefined number of
syllables, within the range of 12 to 17, as desired by
the user (see Fig. 3, for 14 syllables per line).

403

404 CdjOl KM TEAEUTTjo" CtOlV. [180c] OctlSQOV LiEV TOIOUTOV

405 Tivd Aoyov Ecfin eitceiv, LiETd be ©ai&oov dA-

406 Aoui; Tivd? eivai, cov ou navv 6i£fivrm6v£u-

407 ev ovc, nageic, tov Ylavaaviov Aoyov &ir)Y£i-

408 to. EL7TE lv 5' cuTtov 6tl' ov KaAcoc, |uoi 6okeL, co

409 OaL&QE, 7iQo[3£|3Arja0ai. r\yiiv o Aoyoi;, to anAa)?
410

FIG. 3 THE LAYOUT OF THE OUTPUT FILE

The above output is achieved by the main algorithm of
the system, which performs counting and spelling,
according to the relevant rules.

In counting, the goal is the correct identification of any
syllable. For this task, the discovering of vowels is
sufficient enough, while the consonants and other
marks can be ignored. There are seven vowels in
Greek Alphabet (Ancient or Modern alike):
[a, e, n, i, o, u, co].

Each one can form a syllable alone. Yet, there are also
ten combinations of them in Ancient Greek, the so
called diphthongs, consisting of two vowels:
[at, au, el, eu, ni, nu, 01, ou, in, an].

Each combination counts for only one syllable. Thus,
the task here is not counting two syllables, where there
is only one. So the main steps of the algorithm are the
following (as expressed in a usual pseudo-language
for computing purposes):

■ Check a/next character.

■ IF vowel, THEN increase the number of
syllables by one and GO TO the next step,
ELSE ignore and RETURN to previous step.

■ IF vowel in previous step, THEN check the
next character.

■ IF the next character is vowel (as well), THEN
check for diphthong, ELSE RETURN to the first
step.

■ IF diphthong in previous step, THEN RETURN
to the first step (we don't want to count two
syllables here), ELSE increase the number of
syllables by one (we have two adjacent vowels
not forming a diphthong) and RETURN to first
step.

When the desired number of syllables per line is
computed, the tool forms a line in the output file and
then repeatedly continues for the next line, until the
end of the input text-file. But here we may have the
case of a word having to be split in two lines. This
hyphenation must be performed correctly, which is
the task for the spelling module of the tool. This
module improves the theoretically correct appearance
of the output, but it is not essential for the overall task
of the software system, which is the precise counting
of the syllables per line.

Other advantages of the system are the following:

■ Consistency and simplicity.

■ Minimization of the user's activities and
immediate feedback.

■ Minimization of computing storage.

■ Alignment with user's perceptions.

■ Flexibility and adaptability.

The performance of the software system reaches a
success rate of 100% in the correct counting of
syllables and an average of 99.93% in spelling, since
only two out of 2996 lines of text (in the 12 syllables
per line version) display an error. This performance
has been verified manually. The previous error rate is
insignificant both compared to the size of the text and
the desired goal, since it cannot affect any research
conclusions.

Results and Conclusions

Since a mathematical (numerical) structure was
determined to be implemented in an author's work,
without considering various quantitative patterns of
linguistic nature, there are at least four possible ways
to do it:

i) by counting the number of written characters,

ii) by counting the number of written syllables,

iii) by counting the number of written words, and

iv) by arranging the text in lines of a particularly
chosen pattern.

11

www.seipub.org/ijla

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

For example, if Plato, as influenced by the
Pythagoreans, desired to give a mathematical
structure to his works, based on the Pythagorean
sacred number 12, he would have to do it via one of
the above ways.

Regarding the first way (i), having Plato's work
Symposium as the examined text, J. B. Kennedy has
presented an underlying mathematical structure,
which is summarized in the next table (Table 4),
constituting the initial hypothesis of our research.

TABLE 4 STRUCTURE BY CHARACTERS

Speech of

Proportion/Size

Initiating
position

Pausanias

1/12

2"<yi2

Eryximachus

1/12

3 rd /12

Aristophanes

1/12

4*/12

Agathon

3/12

Socrates
(including Diotima's words)

Alcibiades

2/12

Regarding the second way (ii), by using the developed
computational tool, it is measured that the Symposium
consists of 35,941 syllables. According to this total, the
Symposium does not demonstrate any significant
mathematical structure compared to the previous one,
as presented in the following table (see Table 4:
Proportion/Size vs. Table 5: Ratio).

TABLE 5 STRUCTURE BY SYLLABLES

Speech of

Number of syllables

Ratio

Pausanias

3,317

1/10.8

Eryximachus

2,371

1 / 15.2

Aristophanes

3,214

1/11.2

Alcibiades

5,105

1/7.0

Regarding the third way (iii), the Symposium consists
of a total of 17,991 words. Considering this, the ratios
of the above speeches in number of words are
presented in the next table (Table 6).

TABLE 6 STRUCTURE BY WORDS

Speech of

Number of words

Ratio

Pausanias

1,579

1/11.4

Eryximachus

1,149

1 / 15.7

Aristophanes

1,536

1/11.7

Alcibiades

2,601

1/6.9

These ratios are very close to the ones of the previous
table (Table 5). Except for the ratio's proximity
between the speeches of Pausanias and Aristophanes
in both tables, only the speech of Aristophanes is close
to the proportion of 1/12 (see Table 4). Yet, the overall
picture is that neither this way conforms to the initial
hypothesis. Finally, the outcome of the last way (iv) is
presented in the following table (Table 7) and

commented below.

The first observation is that the ratio of each speech
(Table 7: Ratio per S/L) is approximately the same, no
matter how many syllables per line the text is
arranged (including the structures of 12 and 17
syllables per line, which are not presented in Table 7),
as expected. This ratio is almost the same to the one of
the second way (see Table 5: Ratio) and it is presented
only for comparison purposes to the other ways. So far,
the initial hypothesis is not verified by this way either.

TABLE 7 STRUCTURE BY LINES OF FIXED NUMBER OF SYLLABLES

Syllables/Line
(S/L)

13

14

15

16

Ratio

Total lines
per case

2765

2568

2397

2247

per

S/L

Speech of

Number of lines per speech

Pausanias

255

237

221

207

1/ 10.8

Eryximachus

183

170

158

148

1/ 15.1

Aristophanes

247

230

214

201

1/11.2

Alcibiades

393

365

340

319

1/ 7.0

Another possible way to examine the structure of the
Symposium, regarding the fourth way, would be not
having a fixed number of syllables per line, for the
entire extent of the text. Namely, the average number
of syllables per line would differ from one speech to
another, but still within the range of 12 to 17, aiming to
conform to the initial hypothesis. This notion is
consistent with the one presented in Table 3. In this
manner, the speeches of Pausanias, of Eryximachus
and of Aristophanes, should occupy the same number
of lines (the 1/12 of the total lines), while the speech of
Alcibiades should occupy double this number. Let us
call this number "NoL". In such a way, the desired
outcome is depicted in the next table (Table 8).

TABLE 8 STRUCTURE BY LINES OF VARIED NUMBER OF SYLLABLES

Speech of

Pausanias

Eryximachus

Total syllables
per speech

3317

2371

Average number of
syllables per line

Ni

N2

NoL

3317/ Ni

2371/ N2

Speech of

Aristophanes

Alcibiades

Total syllables
per speech

3214

2552
(= 5105/2)

Average number of
syllables per line

N3

N4

NoL

3214/ N 3

2552/ N 4

To examine the mathematical possibility of the above
proposal, the NoL between the highest and the lowest
values of the ratio should be computed, which are
those of Pausanias and Eryximachus:

3317/ Ni = 2371/ N2 3317/2371 = Ni/ N2 ->

12

International Journal of Literature and Art (IJLA) Volume 1 Issue 1, August 2013

www.seipub.org/ijla

N1/N2 = 1.40 .

The above result (1.40) could have been verified only if
the speech of Pausanias consists of 17 syllables per line
(Ni = 17), while the one of Eryximachus consists of 12
syllables per line (N2 = 12), since {17/12 = 1.42}. The
normal range of values, though, according to Table 3,
is from 13 to 16, giving a result of 1.23 (=16/13). Thus, it
is suggested here that the previous proposal is highly
unlikely.

In conclusion, none of these ways of text structuring,
except the first one (i), does show interesting results
for the Symposium. It is doubtful though if any
particular text could have been structured in many
ways simultaneously, so the appeared failure of the
ways (ii) - (iv) is rather expected. There is also the
possibility that an author may follow a different way
to structure a different work. To reveal whether an
author proceeded to similar structures or not, it is
essential to explore all of them. The main purpose of
this paper was to suggest more than one way to
conduct this exploration, by demonstrating the
proposed methodology having the Symposium as the
test example. The Symposium was chosen because one
method (i) had already been applied to (before the
present work), consequently, the demonstration and
the comparison to the rest of them would have been
easier and more clear. In this respect, the contribution
of a computerized tool is invaluable, as it was
hopefully demonstrated in the present paper, along
with the essential (the authors claims also most
enjoyable) interdisciplinary collaboration between
Software Engineering and the Arts & Humanities.

ACKNOWLEDGEMENT

The authors would like to thank Prof. J.B. Kennedy, of
The University of Manchester, UK, both for his
original idea and for providing us with valuable
textual material. They would also like to thank the
reviewers for their useful suggestions.

REFERENCES

Bury, R.G. "The symposium of Plato, (A) The Method of

Narration and the Preface". Retrieved from: www.

perseus.tufts.edu.
Diogenes Laertius (In Greek). "Lives of the Philosophers 2,

Books 3-5". Athens: Cactus, 1994.
Foxall, J. "Sams Teach Yourself Visual C# 2008 in 24 Hours:

Complete Starter Kit". USA: Pearson Education Inc., 2008.

Halporn, W.J., Ostwald, M. and Rosenmeyer T.G. "The

Meters of Greek and Latin Poetry", (Revised Edition).

USA: University of Oklahoma Press, 1980.
Johnson, WA. "Bookrolls and Scribes in Oxyrhynchus".

Toronto: University of Toronto Press, 2004.
Kennedy, J.B. "The Musical Structure of Plato's Dialogues".

Durham: Acumen, 2011.
Kennedy, J.B. "Plato's Forms, Pythagorean Mathematics,

and Stichometry". Retrieved from: http://personalpages.

manchester.ac.uk/staff/jay.kennedy/Kennedy_Apeiron_p

roofs.pdf, 2010.
Ohly, K. "Stichometrische Untersuchungen". Leipzig: O.

Harrassowitz, 1928.
Papakitsos, E. "Computerized Scansion of Ancient Greek

Hexameter". Literary and Linguistic Computing, Vol. 26,

Issue l.Oxford University Press, 2011,doi: 10.1093/llc/fqq

015.

Papakitsos, E. (In Greek). "Scansion of Hexameter by Object-
Oriented Programming". Documentation manual of a
program in Visual Basic 6.0. Greece: PDSA 567/E. K.
Thessalou/Themistokleous 42, Athens, 2009.

Papakitsos, E. (In Greek). "Automated Scansion of Ancient
Greek Epic Poetry". Documentation manual of a
program in Turbo Pascal 6. Greece: National Library of
Greece 4825/6-12-2000, Athens.

Papakitsos, E. "Automated Scansion of Classical Greek
Verse". Postgraduate Dissertation (MSc in Information
Systems). Liverpool: The University of Liverpool/
Department of Computer Science, 1992.

Tsoumaki, M. (In Greek). "Automated Stichometry: The
relation between Plato & Pythagoras". Postgraduate
Dissertation (Interdisciplinary Postgraduate Programme
"Technoglossia"). Athens: National & Kapodistrian
University of Athens/ Faculty of Philology - National
Technical University of Athens/ School of Electrical &
Computer Engineering, 2012.

Tzouma, A. (In Greek). "Introduction to 'Narratology',
Theory and Application of Narrative Typology of G.
Genette". Athens: Symmetry, 1997.

Vlastos, G. "Socrates: Ironist and Moral Philosopher". Ithaca:
Cornell University Press, 1991.

13

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
