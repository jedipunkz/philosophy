---
source: "https://archive.org/details/micro_IA41152922_0276"
title: "PLATO IV: First Year Report, Computerized Training System, Project ABACUS."
author: "Hinkle, Lawrence R., Army Signal Center and School, Fort Monmouth, NJ. Computerized Training System."
year: "1974"
captured_at: "2026-08-11T13:20:37Z"
updated_at: "2026-08-11T13:20:37Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato"
plain_text_url: "https://archive.org/download/micro_IA41152922_0276/micro_IA41152922_0276_djvu.txt"
public_domain: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# PLATO IV: First Year Report, Computerized Training System, Project ABACUS.

- 著者: Hinkle, Lawrence R., Army Signal Center and School, Fort Monmouth, NJ. Computerized Training System.
- 初版: 1974
- 情報源: [archive](https://archive.org/details/micro_IA41152922_0276)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

ED 095 926

AUTHOR
TITLE

INSTITUTION
SPONS AGENCY
REPORT WO
PUB DATE
NOTE

EDRS PRICE
DESCRIPTORS

IDENTIFIERS

ABSTRACT

DOCUMENT RESUME

IR 001 113

Hinkle, Lawrence R.

PLATO IV: First Year Report, Computerized Training
System, Project ABACUS.

Army Signal Center and School, Fort Monmouth, N.J.
Computerized Training Systen.

Army Trainin and Doctrine Command, Fort Monroe,
Va.

CTS-TR-74-2

1 Apr 74

18p.

MF-$0.75 HC-$1.50 PLUS POSTAGE
*Annual Reports; *Computer Assisted Instruction;

Computer Programs; Electronics; Equipment Evaluation;

Material Development; Military Training

Advanced Research Projects Agency; Army; ARPA;
*Computerized Training Systems; PLATO IV; Project
ABACUS

This report covers the PLATO IV activities during

calendar year 1973 at the Computerized Training System. The work
reported herein is supported by a program sponsored by the Advanced
Research Projects Agency (ARPA) to evaluate the PLATO IV system for
use in training, oriented to the needs of the Armed Services. The
report presents a synopsis of each lesson written, a discussion of
support routines for these lessons, a summary of demonstrations of
PLATO IV given by the Computerized Training System (CTS), and the
future PLATO IV plans of CTS. (Author)

TRADOC BEST COPY AVAILABLE

ile CTS - TR74-2
crs

ED 095926—

PLATO IV
First Year Report
Computerized Training System
Project ABACUS

Captain Lawrence R. Hinkle

Office of the Product Manager
Computerized Training System

US Army Signal Center and School
Fort Monmouth, New Jersey 07703

1 April 1974

Approved for public release
Distribution unlimited

Prepared for
US ARMY TRAINING AND DOCTRINE COMMAND

Fort Monroe, Virginia 23651

BEST COPY AVAILABLE

NOTICES
This report has been reviewed and is appreyed.

—trnah Chucate

Frank E. Giunti
Technical Director |

Disclaimer

The contents:of the report are not to be construed as
an official Department of the position unless
so designated by other authorized documents.

» Disposition

Destroy this report when it is no longer needed.
Do Wot return it to the originator.

UNCLASSIFIED

Security Classification

DOCUMENT CONTROL DATA-R&D

(Security classification of title, body of abstract and indexing annotation must be entered when the overall report is claseilied)

’ 6 INA TIN Bradunt | orate euthor) 2a. REPORT SECURITY CLASSIFICATION
ice o uc anager :
. Unclassified

Cuamnanens Training System aS
US Army Signal Center & School, Fort Monmouth, NJ

32. REPORT TITLE
PLATO IV First Year Report
Computerized Training System
| __ Project ABACUS

4. CESCRIP TIVE NOTES (Type of report end inclusive dates)

Technical Report TR74-2

8. AUTHOR(S) (Firat name, middie initial, last name)

Captain Lawrence R. Hinkle

6 REPORT CATE 7a. TOTAL NO. OF PAGES 7b. NO. OF REFS

1 April 74 14

ta. CONTRACT OR GRANT NO. 9a. CRIGINATOR’S REPORT NUMBER(S)
CTS-TR-74-2

b, PROJECT NO.

9b. OTHER REPORT NOS) (Any other numbere thet may be escigned
thie report)

d.

10. DISTRIBUTION STATEMENT

Distribution of document is unlimited

11. SUPPLEMENTARY NOTES 12. SPONSORING MILITARY ACTIVITY
US Army Training and Doctrine Command
Fort Monroe, Virginia 23651

13. ABSTRACT

This report covers the PLATO IV activities which have transpired during calendar year
1973 at the Computerized Training System.
t»
The work reported herein is supported by a program sponsored by the Advanced Research
Projects Agency (ARPA) to evaluate the PLATO IV system for use in training, oriented to the
needs of the Armed Services.

The report presents a synopsis of each lesson written, a discussion of support routines
for these lessons, a summary of demonstrations of PLATO IV givex by the Computerized
Training System, (CTS) and the future PLATO IV plans of CTS.

DD 147s ta UNC LASSIFIED
Security Clessification

UNCLASSIFIED

Security Classification

KEY woRODS

Instructional Programmer Training
Lesson Support Routines

PLATO IV Lessons

PLATO IV Graphics

Plasma Panel Efficacy

UNCLASSIFIED

Security Classification

FOREWORD

This report covers the PLATO IV activities which have trans-
pired during calendar year 1973 at the Computerized Training
System.

The work reported herein is supported by a program sponsored
by the Advanced Research Projects Agency (ARPA) to evaluate the

PLATO IV system for use in training, oriented to the needs of the
Armed Services.

The report presents a synopsis of each lesson written, a dis-
cussion of support routines for these lessons, a summary of demon-
strations of PLATO IV given by the Computerized Training System
(CTS), and the future PLATO IV plans of CTS.

SE /CPu

G. B. HOWA
COL, SigC
Product Manager, CTS

i'able of Contents

Foreword . . .--.

. . 2 . * . . 2 ° .

Int roduction . ° oa . . . . e ° . om 7 . . . * 7 . * .

Course Material Prepared
Ohm's LAW s..:0 sis + 2 + ss
Series Circuits. .....
Parallel Circuits. ......
Series-Parallel Circuits. .
DO Power sic cc tae ee
Troubleshooting Procedures
First Aid and Safety....
BADIM 0:0 0a 808s 08%
Recruiting . «:. oes «tee.
DELVOTR yo 0 6 6 00 6 OHO 8
Support Routines. ......ee-s

aQOoanw# kt PP | WH WD b& DH

for)

Instructional Programmer Training .

Technique Development
~ectronic Symbol Set. .....
Multiple Ceoice Driver .....
Numeric Driver... s s.« «i« «ss
Comment Collection Routine ..
Data Manipulation. .......

Terminal Operations... .

System and Terminal Reliability .

Teenie... 6 5 + 0 oe 8 Hs
Future Plans oe — — - 2 . . . a om os . oe — ae a 7 e -_ = . . * se 7 a — * — . — — es

eg ee ere ee Ge ee oe eee a eg ee a ee eae eee er
Table 1. Average PLATO IV Terminal Usage .......2-2eecc06 «©
Table 2. Percentage of Authoring and Downtime. ........2.e. -
Table 3. Pemometrations 1078 os 6.5.3 66 oe Ta oe ee eS 8 8
Table 4. On-Gite Demonstrations.1973 26. 6 6 be 6 68 6 6 6 Oe 0

ii

Purpose

This report documents the Computerized Training System's accomplishments,
findings and progress in the Joint Service/Advanced Research Projects Agency
(ARPA) PLATO IV evaluation program. Site activities up to 31 December 1973
will be covered.

Introduction

The Computerized Training System's (CTS) involvement with the PLATO IV
project began in early January 1973 with the delivery and installation of two PLATO
IV student consoles. Two additional consoles were received in March giving the
CTS project a total of four. These four student consoles operate over one vuice
grade commerciat telephone leased line using an experimental multiplexer developed
by the Computer-based Education Research Laboratory (CERL) at the University of
Illinois, Urbana, Dlinois. The four terminals were purchased by the Advanced
Research Project Agency for use by CTS at the US Army Signal Center and School
(USASCS), the leased line and maintenance cost are paid by ARPA,

The objectives of CTS in participating in this program are given below. Each
of these objectives will be addressed in this report, as wil! the future PLATO IV
plans of the CTS project. Also covered will be the PLATO IV demonstrations con-
ducted by CTS.

a. Prepare course materials aimed at the development of technical and
occupational skills.

b. Train instructional programmers in the development of computer-based
instructional materials.

c, Develop techniques and procedures for developing and presenting
computer-based course materials.

d. Evaluate the effectiveness and efficiency of the PLATO IV Plasma Ter-
minals from the instructional and human factors points of view.

e. Evaluate system and terminal reliability.

Course Material Prepared

To date a total of 11 lessons have been developed by CTS. They each will be
discussed in this section along with the lesson support routines that have been
developed. Five lessons have been written in the electronic fundamentals area.
They are as follows:

a. Ohm's Law

(1) This lesson was the initial CTS lesson and precipitated the develop-
ment of a multiple choice test scoring driver for use in later lessons.

(2) The objective is to teach the relationship of current, voltage, and
resistance as stated by Ohm's Law.

(3) Lesson was completed in April 1973.
(4) The author is Mr. Joseph Rich.

(5) One part lesson utilizing approximately 4,000 words External Core
Storage (ECS).

(6) Material presented by tutorial mode.
(7) Average completion time is 30 minutes.
b. Series Circuits
(1) The objective is to teach the student to recognize schematic repre-
sentation of short, simple and series circuits and to determine the
polarity of a voltage source.

The author is Mr. Tom Button.

Lesson completed in July 1973.

One part lesson using approximately 3,400 words ECS.

Material presented by tutorial mode.

Average completion time is 25 minutes.

c. Parallel Circuits

2

(1) The.objective is to teach the student to identify parallel circuits and
compute values of voltage, current, and resistance.

(2) The author is Mr. Joseph Rich.
(3) Lesson completed in June 1973.
(4) Three part lesson using approximately 10, 600 words ECS.
(5) Material presented by tutorial mode.
(6) Average completion time is 80 minutes.
d. Series-Parallel Circuits

(1) The objective is to teach the student to identify a series-parallel
circuit and compute values of voltage, current and resistance.

(2) The authors are CPT John Lohr and Mr. Joseph Rich.

(3) Lesson completed in December 1973.

Two part lesson approximately 6,500 words ECS.

Material presented in tutorial mode.
Average completion time 50 minutes.
e. DC Power

(1) The objective is to teach the definition of DC power to compute values
using Ohm's Law.

(2) The authors are Mr. Ed Colliton and Mrs. Janet Lamb.
(3) Lesson completed in June 1973.
(4) Two part lesson approximately 5,500 words ECS.

(5) Average completion time 35 minutes.

f. Troubleshooting Procedures
One lesson has been developed to introduce the student to a systematic
approach for troubleshooting electronic equipment. This lesson was to test the

feasibility of using multimedia with PLATO IV.

(1) The objective is to teach the student a set pattern for locating prob-
lems in electronic equipment.

(2) The authors are Mr. Tom Button, Mr. Ed Colliton, Mrs. Janet Lamb,
SP5 Joseph Martinez, Mr. Joseph Rich and Mr. Kermit Van Pelt.

(3) Lesson completed in August 1973

(4) One part lesson approximately 2,100 words ECS.

(5) Material presented in syncronized tutorial and sound on slide mode.

(6) Average completion time 45 minutes.
g. Two lessons were written to present Army common subjects.
(1) First Aid and Safety

(a) The objective is to determine the student's level of first aid
knowledge and to provide instruction if needed.

(b) The author is SP4 Tom Plunk.
(c) Lesson completed in August 1973.
(d) One part lesson approximately 3,000 words ECS.

(e) Material presented in on-line test and instruction in closed cir-
cuit television cassette.

(f) Average completion time 40 minutes.
(2) Subversion and Espionage Directed Against the Army
(a) The objective is to create awareness of SAEDA in student.

(b) The author is SP4 James Samuelson.

(c) Lesson completed in August 1973.
(d) One part lesson approximately 2,000 words ECS.

(e) Material presented in a test and instruction in closed circuit
television cassette.

(f) Average completion time 65 minutes.
a. Recruiting
A career guidance lesson was written for use by recruiters.

(1) The objective is to inform prospective enlistees of the career oppor-
tunities existing in the Army.

(2) The author is CPT Larry Hinkle.

(3) Lesson completed in November 1973.

(4) Two part lesson approximately 4,000 words ECS,

(5) Material presented in an informative reference mode.
(6) Average completion time 50 minutes.
b. Drivers

A lesson was written to explain operation of the multiple choice and nu-
meric drivers.

(1) The objective was to teach instructional programmers how to use
CTS drivers.

(2) The authors are SP4 Peggy McClintock and CPT Larry Hinkle.
Lesson completed in June 1973.
One part lesson approximately 4,000 words ECS.
Material presented in a tutorial and drill mode.

Average completion time 30 minutes.

c. Index
An index lesson was written to provide a list of all available CTS lessons.
The index includes the electronic symbol set, schematics, the ABACUS symbol,
and a slide program.
(1) The objective is to provide easy reference to CTS lessons.

(2) The author is SP4 Peggy McClintock.

(3) Lesson completed October 1973.

(4) One part lesson approximately 2,500 words ECS.

(5) Material presented in demonstrative mode.
d. Short Demonstrations and Administrative Routines

Several short demonstrations and administrative routines have been written.
Details on these routines are available upon request from CTS. The demos and
routines are listed below:

(1) Multiple choice drivers

(2) Numeric judging drivers

(3) Student comment collection routine

(4) Electronic symbol character set

(5) Student date file manipulator

(6) Career guidance charset

(7) Move the tile game

(8) Electron flow demonstration

Instructional Programmer Training

Eight instructional programmers received their initial training in PLATO IV
authoring at the University of Illinois. Six received one week of training while
two received an additional week of instruction. After this initial period, further
training was conducted at Fort Monmouth using material provided by CERL.

A total of 15 individuals became proficient enough in authoring to develop
material suitable for student use. Six of the instructional programmers are no
longer involved with CTS. Currently three instructional programmers “re under-
going instruction at CTS in authoring techniques.

The self-paced material provided by the CERL was found to be a very good
workbook for new instructional programmers. After the new author had com-
pleted most of the CERL material he was given an assignment and then as he
developed the material he could use the more experienced instructional program-
mers to assist him through difficulties. We have found the best results in tiaining
new authors are obtained when they are working on aterminal. Developing material
on-line will rapidly train an author.

Technique Development

In developing lesson material for PLATO IV, experience has shown that a
minimal team concept will ‘produce the fastest results and therefore, CTS uses
basically a three man lesson development team. The instructional programmer
is responsible for the total lesson and as such writes and codes his lesson off-
line, on coding sheets. He can be assisted in special coding problems by a system
programmer from the Systems Operations and Programming Division when re-
quired. Prior to on-line entry the lesson is reviewed by the Course Development
Chief. The third person in the team is the entry specialist. The coding sheets of
a completed lesson are turned over to the entry specialist, this individual enters
the lesson into the computer making minor editing changes as needed. The in-
structional programmer with the assistance of the entry specialist and a systems
programmer then debugs the lesson. Finally, the completed lesson is subjected
to peer review.

Under this concept the Systems Operations and Programming Division de-
veloped several routines and drivers to assist the instructional programmer.
They are:

a. On-line electronic graphics symbol set. This is described in more detail
in CTS Technical Report 73-3.

b. Multiple choice answer judging driver. This driver judges multiple choice
answers with 3, 4, or 5 distractors, prohibits use of the same wrong answer
twice, provides and unrecognized answer message, and after n wrong attempts
(n specified by author) locks the keyboard thus forcing the student to seek help
from the classroom proctor or instructor.

c. Numeric answer judging driver. This driver allow the instructional pro-
grammer to specify tolerance, high remediation, low remediation, provides an
unrecognized answer message, and after n wrong attempts the keyboard will lock.

d. Student comment collection routine. This routine was copied from an
cxisting routine created by an author at CERL. It was modifed to suit CTS require-
ments and allows the instructional programmer to collect student comments on his
lesson, to aid in improving instruction.

e. Data manipulation. This routine was written to facilitate collection of
student data and to allow the semi-permanent saving of desired information.

Terminal Operations

To evaluate the effectiveness and efficiency of the PLATO IV Plasma Terminals
from the instructional and human factors points of views, CTS surveyed each CTS
instructional programmer and system programmer for their comments about the
terminals. The major responses are listed below:

Insufficient lesson and core space
Inadequate microfiche production techniques
Projector bulb short lifetime

Premature requirement for slide projector
Spacers in screen not lighting

Glare from panel

System crashes

Terminal malfunction

Documentation off-line non-existent

Excessive parity

k. Key arrangement on keyboard

Many of these complaints are in the process of being satisfied and all others
have been at least discussed with CERL during the PLATO User Group meetings.

System and Terminal Reliability

To evaluate system and terminal reliability, CTS initiated a log book policy.
Each terminal was given a log book and terminal users were required to log into
the book when they used the terminal. The logs were then used to report both
system and terminal problems encountered by the terminal user. The icllowing
data are summarized from the logs up to 31 December 1973.

During this reporting period terminal usage averaged 6 hrs a day of the avail-

able time. Data shown in Table 1 below are the average time for the four ter-
minals in each category and are reported in hours.

Table 1. Average PLATO IV Terminal Usage

Authoring Available
Month Time Downtime Time

March 16.50 2.75

April 53.25 10.25
May 51.00 8.75

June 32.25 2.75
July 31.00 1.5
August* _

September * _

October 79.25

November 63.75

December** 53.00 1.50

*In August and September CTS was relocating and logging procedures
were suspended.

**In December three terminals were inoperative due to the failure of
the A9 modem in each.

Two problems arose with the log on policy. First there was no way to insure
the logs were filled in so the actual times recorded are less for both authoring and
downtime, and the second was that during the actual authoring process, system
crasher, ECS shortages and parity errors were often not recorded when they
occurred.

A conclusion drawn from the data collected is that the system reliability
improved over the reporting period as shown below in Table 2. Percentages
shown in the first column are the ratios of authoring time to available time and
the nercentages shown in the second column are the ratio of downtime to avail-
able time.

Table 2. Percentage of Authoring and Downtime

Authoring
Month Time Downtime

March 55% 9.17%
April 42% 8.53%
May 39% 6.63%
June 2.18%
July 1.19%
August*

September*

October 5.16%
November "0.44%

December j 1.56%

*In August and September CTS was relocating and logging pro-
cedures were suspended.

Demonstrations

Table 3 lists the PLATO IV demonstrations given by the Product Manasrer's
Office, CTS, during the reporting period. For each of these demonstrations it was
necessary to request a dial up telephone and modem from CERL. The Product °
Manager, CTS, recommends that a dial up telephone and modem be given to CTS to
facilitate future demonstrations.

Table 3. Demonstrations 1973

Location Recipient

Fort Monmouth, N. J. Association of US Army Armed
Forces Communication Electronics
Associates

7-8 May Pentagon, Wash, D.C. MG Fair

17-20 Sep Fort Gordon, GA Fourth Communications System
Program Review

T;bie 4 reflects the number of individual requesters, end gives on-site demon-
strations of PLATO IV during the reporting period.

Table 4. On-Site Demonstrations 1973

Category Number

General Officers

Other U. S. Military
Foreign Military

ROTC

Civilian Educators /Students

Industry

Future Plans

In order to refine existing lessons and to provide input for satisfying the objec-
tives of CTS in participation in the PLATO IV program, plans are being made to
run student trials of the CTS PLATO IV lessons at Fort Monmouth.

The CTS P», ;ect has recommended to ARPA that consideration be given to re-
assigning two of CTS's terminals as follows: One terminal to Army Research In-

stitute (ARI) and one to Brook Army Hospital. The remaining two terminals to re-

main with the PMO, CTS for continued participation in the Joint Service ARPA
PLATO IV evaluation program.

Consistent with this recommendation CTS is re-
vising its 'Plan for Use and Evaluation of PLATO [V Plasma Terminals at tne US
Army Signal Center and School,"

Misa-F’ 19-77

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
