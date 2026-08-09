---
source: "https://archive.org/details/calc2platoivless611schr"
title: "calc2 - a PLATO IV lesson on differentiation"
author: "Schreiner, Axel T"
year: "1973"
captured_at: "2026-08-09T18:53:29Z"
updated_at: "2026-08-09T18:53:29Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato"
plain_text_url: "https://archive.org/download/calc2platoivless611schr/calc2platoivless611schr_djvu.txt"
public_domain: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# calc2 - a PLATO IV lesson on differentiation

- 著者: Schreiner, Axel T
- 初版: 1973
- 情報源: [archive](https://archive.org/details/calc2platoivless611schr)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

mBMSm

HN18ttmfiM

RBDSDBHRffiMflUffi

ImBIHSI

B

NMI      hRm

H

H  13

■■■■HI

'»w }        B

HI
H  BB

BHBfl

t%®     m  m  h

■

hOubHI

HP

I

mam

LIBRARY  OF  THE

UNIVERSITY  OF  ILLINOIS

AT  URBANA-CHAMPAIGN

510. 84

no.  Q>01- 612
cop.  2

CEMRAl  CIKUIMION  AMD  BOOKSTACKS

responsible    or  its  J™™a    below.  You

^X^»?TTi0f  $75°°

be  causes  for  student  <W""g '        Library  are  the
-"^^te'STfVnols  and  ore  projected  by

Univers£^^

APR

When  renewing  by  phone,  write  new  due  date
below  previous  due  date.

'«:

%lfi-

Report  No.   UIUCDCS-R-73-611
t  ill

/TX^L^L^

calc2  -  A  PLATO  IV  Lesson  on  Differentiation

by

Axel  T.   Schreiner

December  1973

THE  LIBRARY  OF  THE

FEB  1  5  1974

AT  >'°r

DEPARTMENT  OF  COMPUTER  SCIENCE
UNIVERSITY  OF  ILLINOIS  AT  URBANA-CHAMPAIGN

URBANA,  ILLINOIS

Report  No.  UIUCDCS-R-73-611

calc2  -  A  PLATO  IV  Lesson  on  Differentiation

by
Axel  T.  Schreiner

December  1973

Departments -of  Computer  Science  and  Mathematics
University  of  Illinois  at  Urbana-Champaign
Urbana,  Illinois  61801

Digitized  by  the  Internet  Archive
in  2013

http://archive.org/details/calc2platoivless611schr

Ill

ACKNOWLEDGMENT

The  PLATO  IV  consulting  staff,  especially  Judy  Sherwood  and  Jim
Kraatz,  helped  me  many  times  when  TUTOR  and  I  were  on  different  wave  lengths.
After  I  was  done  programming,  Rick  Blomme  and  his  staff  further  improved  the
PLATO  system.  Professor  H.  George  Friedman,  Jr.  (as  usual)  read  this  paper
and  removed  my  hyphens.  And  Vivian  Alsip  read  my  handwriting  and  produced  a
perfect  manuscript.

Without  all   this  help,   I  would  still   be  reading  in  Knuth's  book  [1]
on  how  to  differentiate...

IV

TABLE  OF  CONTENTS

Page

1.  Introduction 1

2.  Formula  Management 2

3.  Formula  Representation 3

4.  The  Various  Types 5

5.  Input  (readin,  and  postfix)  7

6.  The  Choice  Page  (options) 10

7.  The  Differentiation  Algorithm  (derive) 12

8.  Interfacing  with  the  Student  (infix,  and  output) 18

9.  Wrap  up 20

10.     Conclusions 21

REFERENCES 23

A  Listing  of  calc2 24

LIST  OF  FIGURES

Page

1.  Storage  Layout  for  the  Formulas 4

2.  Example  of  a  Differentiation  Situation  12

3.  Example  of  the  Differentiation  Process  13

4.  Layout  of  an  Entry  in  the  Driving  Table 17

LIST  OF  TABLES

Page

1.  Types  of  a  Node  in  the  Formula 6

2.  Codes  in  the  Expanded  Input  String  8

3.  The  Input  Parser  Stack  Mechanism  9

4.  The  Differentiation  Situations  16

5.  Displayed  Information  in  a  Differentiation  Situation  18

6.  Use  of  the  Student  Bank 22

calc2  -  A  PLATO  IV  Lesson  on  Differentiation

1.  Introduction

In  the  summer  of  1971,  I  wrote  a  symbolic  differentiator  as  a
demonstration  program  for  a  course  on  computer  applications  in  calculus  [2].
This  program  was,  in  its  central  part,  a  direct  and  unmodified  implementation
of  the  algorithm  described  by  Knuth  [1],  but  it  was  able  to  differentiate  all
trigonometric,  inverse  trigonometric,  and  hyperbolic  functions,  as  well  as  ex,
and  various  logarithms.  The  program  was  written  in  BASIC.

In  fall  1972,  I  decided  to  recreate  this  program  on  the  PLATO  IV
system  [3],  originally  as  a  demonstration  for  the  various  stack-oriented  algo-
rithms used  during  input,  differentiation,  and  output  of  a  formula.  BASIC  had
not  been  a  good  language  for  string-  and  tree-manipulation  (neither  is  TUTOR,  the
language  of  the  PLATO  IV  system),  and  the  recreation  was  also  intended  to  improve
on  the  implementation  of  the  algorithm.

During  the  design  of  this  PLATO  IV  lesson,  calc2,  however,  I  decided
to  make  the  student  a  participant  rather  than  a  spectator  to  the  differentiation
process.  calc2  was  to  control  a  student's  attempts  at  differentiation  by
guiding  him  through  the  steps  and  controlling  each  of  his  operations.  To  keep
this  guidance  deterministic  from  the  machine's  point  of  view,  the  student  must
differentiate  in  postorder  (i.e.,  for  each  operator,  he  must  differentiate  the
operands  before  he  is  allowed  to  construct  the  derivative  of  the  phrase  involving
the  operator).  An  adapting  control  allows  him  to  dynamically  bypass  some  or  all
of  the  simpler  steps  in  differentiation.

Calc2  is  table-driven;  it  only  provides  the  routines  to  manipulate  the
data  structures  which  represent  the  formula  and  the  piece-wise  derivatives;  the
actual  differentiation  is  under  control  of  a  table  of  differentiation  algorithms.
It  was  thus  possible  during  the  design  of  this  lesson  to  expand  the  set  of
differentiation  situations—rules  may  be  easily  added  or  deleted.  It  is  quite

conceivable  that  a  particular  table  might  be  introduced  (then  together  with
suitably  chosen  problems!)  which  only  practices  the  chain  rule,  etc.  Similar
effects  can  be  obtained  by  manipulating  the  dynamic  bypass  mechanism.

In  this  report  I  intend  to  describe  how  calc2  operates.  TUTOR,
PLATO'S  authoring  language,  is  ever-expanding,  and  some  coding  in  calc2  might
now  be  inefficient;  however,  I  believe  the  basic  techniques  still  are  acceptable.

2.  Formula  Management

Knuth  [1]  presents  a  formula  to  his  differentiator  as  a  binary
threaded  tree.  If  one  intends  to  perform  a  lot  of  (for  the  mathematician-
user  necessarily  infixed)  formula  fragment  display,  and  if  the  implementation
language  is  not  suited  to  tree  manipulation,  such  a  tree  is  not  the  best  data
structure.

Instead,  I  decided  to  employ  one  queue  and  two  stacks  with  variable
sized  entries.

Initially,  the  queue  is  loaded  with  the  given  formula  in  post  order
(i.e.,  each  operator  follows  its  operands),  each  element  of  this  queue  containing
one  fragment  (operator  or  operand)  of  the  formula.  Both  stacks  are  empty.

As  the  differentiation  progresses,  one  stack  is  pushed  with  phrases
of  the  formula,  the  other  with  the  derivatives  of  these  phrases;  e.g.,  if  a
binary  operator  of  the  formula  is  processed,  the  top  two  phrases  of  each  stack
are  combined  to  form  a  new  top  phrase  of  the  derivative  stack,  and  then  the  top
two  phrases  of  the  formula  stack  and  the  binary  operator  are  combined  to  form  the
new  top  phrase  of  the  formula  stack.  For  display  and  judging  purposes,  these
processed  phrases  are  maintained  with  binary  operators  infixed  between  their
operands  and  parentheses  are  employed.

Constructing  an  infixed  formula  'bottom  up'  allows  one  to  insert  the
minimal  number  of  parentheses  without  any  effort,  provided  the  actual  data

structure  underlying  the  formula  is  flexible  enough  to  accommodate  these
parentheses  without  expanding  in  size.

3.  Formula  Representation

PLATO  IV  operates  on  a  Control  Data  Cyber  70  computer;  the  word
size  is  60  bits,  and  at  the  time  calc2  was  written,  TUTOR  did  not  provide  for
the  subdivision  of  a  word.  Bit  operations  (logical  and,  or,  and  circular
shift)  were  available,  as  was  a  one-dimensional  array  facility.  A  student  in
the  PLATO  IV  system  has  150  "student  variables"  available,  and  it  is  the  lesson
author's  responsibility  to  manage  those  variables;  no  dynamic  variable  allocation
is  provided  in  TUTOR.

I  decided  to  allocate  the  queue  and  both  stacks  in  one  array.  The
queue  occupies  the  beginning  of  the  array.  As  elements  are  disposed  of  from
the  beginning  of  the  queue,  the  formula  stack  overlays  these  elements.  If  we
accommodate  the  parenthesized  representation  at  no  increase  in  space  require-
ments, the  queue  will  not  have  to  be  moved.  The  derivative  stack  follows  the
queue  in  the  array.  This  method  is  destructive,  but  once  differentiation  is
complete,  formula  and  derivative  exist  infixed  and  either  one  could  then  again
be  converted  to  post-order  to  repeat  the  process.

n75  (a  TUTOR  variable)

i

formula
stack

tformel

* 0

item

queue

deriv

nl45

tderiv

derivative  stack
this  slack  results  from  simplification

tformel:  pointer  to  top  of  formula  stack

tderiv:   pointer  to  top  of  derivative  stack

deriv:    pointer  to  bottom  of  derivative  stack
(one  element  beyond  queue)

item:    pointer  to  current  element  in  queue

Figure  1.  Storage  Layout  for  the  Formulas

Elements  in  the  queue  are  one  word  long,  and  these  "nodes"  have  the
following  fields:

M° type   prin    leng    Ipar   rpar   wher

info:  bits  1...8
type:  bits  9... 14
wher:  bits  39... 60

(pointer  to)  auxiliary  information
type  of  this  element

begin  position  of  this  element  in  the  master
display  line

These  elements  are  essentially  copied  from  the  queue  into  the  stack,  and
then  the  following  fields  are  also  used:

rpar:  bits  34... 38:  number  of  )  following  this  node

lpar:  bits  29... 33  number  of  (  preceding  this  node

leng:  bits  21... 28:  length  field

poin:  bits  15. ..20:  principal  type  field

The  length  and  principal  type  field  implement  the  variable  sized  stack  entries
and  allow  simplification:  the  last  node  of  a  phrase  contains  in  its  leng  field
the  length  of  that  phrase,  i.e.,  the  number  of  nodes  constituting  it,  and  the
prin  field  indicates  the  principal  type  of  the  phrase.  This  principal  type
is  defined  to  be  the  type  of  the  root  node  of  the  formula  tree  into  which  the
phrase  could  be  expanded.  If  the  phrase  is  a  function  evaluation,  the  node
representing  the  function  contains  in  its  prin  field  the  principal  type  of  the
function  argument.  It  is  thus  possible  to  "collapse"  the  composition  of  a
function  and  its  inverse—like  two  successive  unary  minus  operators.

4.  The  Various  Types

calc2  allows  its  user  to  specify  the  formula  which  he  wishes  to
differentiate.  The  user  may  use  26  variables  a...z,  the  binary  operators  +,-,
*,/  and  **  for  exponentiation,  the  unary  minus  operator  -,  and  the  functions
sin,  cos,  tan,  cot,  sec,  esc  (trigonometric),  sinh,  cosh,  tanh,  (hyperbolic),
asin,  acos,  atan  (inverse  trigonometric),  and  In,  exp,  sqrt.  Parentheses  may
be  used.  Following  TUTOR  conventions,  multiplication  takes  precedence  over
division.  At  most  points  in  calc2,  the  standard  TUTOR  formula  evaluation
mechanisms  can  be  employed;  these  permit  juxtaposition  of  variable  names  to
denote  multiplication,  superscript  to  denote  exponentiation,  and  omission  of
(most)  parentheses  in  function  application;  however,  in  order  not  to  overly
complicate  the  input  scanner  for  the  given  formula,  it  must  be  specified  in  a
more  rigorous,  conventional  syntax;  we  assign  the  unary  minus  the  highest
precedence  so  that  the  phrase  a**-b  would  be  legitimate.  Upon  input  we  have
two  types  of  terminal  operands:  numerical  values  which  for  simplification
purposes  are  assigned  the  types  0,  1,  and  v,  and  variables  which  are  assigned
the  type  c.  Only  during  differentiation  do  we  make  the  further  distinction
between  type  c  for  constant  variables  and  type  x  for  the  differentiation  variable,

type        type  name

info

1

x(variable)

poi

2

c(constant)

poi

3

0  (zero)

poi

4

1  (one)

poi

5

v  (value)

poi

6

+

7

-

8

*

9

/

10

**

11

-(unary)

12

In

13

exp

14

sin

15

as  in

16

cos

17

acos

18

tan

19

atan

20

cot

21

sec

22

esc

23

sinh

24

cosh

25

tanh

26

sqrt

nter  to  variable  storage  at  random  value  for  x

nter  to  variable  storage  at  random  value

nter  to  constant  storage  at  value  0

nter  to  constant  storage  at  value  1

pointer  to  constant  storage  at  value

Table  1.  Types  of  a  Node  in  the  Formula

(to  be  constructed  by  the  input-parser)

In  TUTOR,  characters  a  through  z  have  the  internal  codes  1  through  26;  we
allocate  variable  storage  (used  for  randomization  during  answer  judgment)  into
variables  vl  through  v26  and  thus  have  the  info  fields  of  types  x  and  c  also
store  the  character  associated  with  the  node.

Values  pose  a  bigger  problem.  When  judging  phrase  answers  I
randomize  all  variables  (types  c  and  x),  evaluate  the  correct  phrase  known  to
the  lesson  and  the  answer  phrase  given  by  the  student,  and  compare.  In  order
to  evaluate  efficiently,  I  intended  to  use  TUTOR'S  "compute"  command  which  takes
a  character  string,  compiles  it  (it  must  be  a  formula),  and  evaluates  it.  Such
a  formula  may  contain  the  usual  operators,  numbers,  and  "defined  names".  Names
are  defined  with  a  "define"  statement  which  assigns  at  compile  time  either  a
variable  location  or  a  constant  as  a  string  replacement  for  the  name.  Treating
the  variables  a...z  is  simple:  they  are  defined  to  be  variable  locations  vl
through  v26,  and  before  using  compute,  random  numbers  can  be  placed  into  those
locations.  How  should  I  maintain  values?  There  exists,  of  course,  a  table  of
values,  and  the  info  field  of  types  0,  1,  and  v,  points  to  an  entry  in  this
table.  But  TUTOR  does  not  provide  a  conversion  from  floating  point  to  character
string  in  memory.  I  could  have  stored  values  as  character  strings,  but  that
would  have  severely  limited  the  simplification  techniques  which  may  introduce
new  values.  I,  therefore,  "defined"  the  names  A,  B,  ...  for  the  locations  of
the  value  table,  and  in  phrases  presented  to  the  "compute"  command,  these  names
are  used  instead  of  values.  This  meant  then  that  output  of  these  phrases  to  the
screen  had  to  be  accomplished  interpretively:  strings  between  values  could  be
shown  directly,  values  had  to  replace  their  names  in  the  phrase  strings.

5.  Input  (readin,  and  postfix)

The  input  parser  accepts  a  formula  and  constructs  the  queue  from  it;

8

the  formula  is  in  infix  notation  and  in  a  conventional  syntax,  the  queue  is  in
postorder,  parentheses-free,  and  encoded  into  nodes.  The  value  table  must  also
be  constructed.  The  input  parser  has  two  phases:  first  the  formula  string  is
converted  from  a  character  string  into  a  sequence  of  codes,  one  per  computer
word.  Blanks  are  eliminated  and  illegal  characters  are  rejected;  function  names
including  left  parentheses  are  verified.  The  codes  are  outlined  in  Table  2.
They  overlay  the  area  which  is  then  occupied  by  the  (more  compact)  queue.

Code Meaning

1  ...

26

variable  a  . . .  z

27

+

28

-

29

*

30

/

31

**

32

the  next  compute

33

the  next  compute

(this  always  indicates  a  left  parenthesis)
34 1

subcode meaning

1         (

2  ...  16       functions  of  type  =  subcode  +  10

Table  2.  Codes  in  the  Expanded  Input  String

The  second  phase  of  the  input  parser  is  a  syntax  analyzer,  driven  by  an  operator
precedence  grammar.  A  stack  is  used  and  the  coded  formula  is  converted  into  the

postorder  nodes  in  the  queue.  This  part  of  the  parser  is  straightforward  and
I  therefore  just  present  it  as  Table  3.

state

next  symbol

code

action

next  state

expect
operand

variable

1  ...

26

emit

expect  operator

value

32

emit  and  enter
into  value  table

expect  operator

function  or
left  paren

33,  si

bcode

stack,  increase  #
of  parens

expect  operand

+

27

ignore  unary  plus

expect  operand

-

28

stack  or  collapse

unary  minus

expect  operand

other

SYNTAX  error

retrieve  new
formula

expect
operator

T  J      >      5   /  >

27

...31

emit  from  stack  and
push  by  precedence

expect  operand

)

34

emit  from  stack
decrease  #  of
parens

expect  operator

end  of  formula

if  no  parens  open

exit

other

SYNTAX  ERR0R

retrieve  new
formula

Table  3.  The  Input  Parser  Stack  Mechanism

TUTOR  is  not  overly  cooperative  in  constructing  either  phase  of  the
input  scanner:  the  formula  is  read  as  a  student  response  to  a  TUTOR  arrow,  and
at  least  conversion  to  the  first  internal  representation  should  have  been  possible
by  using  TUTOR'S  rather  eye-pleasing  array  of  answer-judging  techniques.  Un-
fortunately, no  coding  could  be  devised  to  overcome  certain  time  slicing  problems

10

which  still  exist:  answer  judging  may  not  take  more  than  25ms  of  computer  time
or  the  user  will  be  swapped  and  some  of  the  judging  material  is  lost.  I  was
unable  to  determine  whether  my  code  was  troubled  by  bugs  in  TUTOR  or  whether  I
was  violating  a  principle  of  the  language.  Consequently,  the  first  phase  is
accomplished  by  "brute  force",  the  characters  of  the  given  formula  are  spread
out,  one  to  a  computer  word,  and  then  analyzed.  The  coding  is  certainly  in-
efficient, FORTRAN-style,  and  TUTOR'S  limits  on  the  size  of  a  program  segment
("unit")  were  felt  to  be  a  severe  restriction.

The  second  phase  of  the  scanner  charged  with  creating  postorder  is
simple,  except  for  the  treatment  of  the  wher  field  in  each  output  node:  even-
tually our  lesson  will  present  on  each  page  a  differentiation  situation  taken
from  the  given  formula.  This  formula  is  also  shown  on  each  page,  and  an  arrow
indicates  the  location  of  the  presented  situation  in  context.  We  thus  must
maintain  with  each  node  an  indication  where  it  is  displayed  in  the  formula.
This  information  is  maintained  in  the  wher  field,  and  it  must  be  stacked  for
operators  during  the  postorder  conversion.  The  information  is  obtained  by
writing  the  element  onto  the  screen  (with  a  standard-length  format  for  constants)
and  observing  how  much  space  is  needed  for  writing.  (Theoretically,  TUTOR'S
screen  location  variable  "where"  could  be  used,  but  "where"  had  not  yet  been
completely  debugged.)

In  order  to  generate  the  given  formula  display  efficiently  on  each
new  page  (since  a  selective  screen  erase  conserving  the  formula  turned  out  to  be
too  slow),  it  is  generated  and  saved  in  infixed  form  in  memory,  with  values  re-
placed by  capital  letter-defined  names,  while  the  endorder  queue  is  constructed.

6.  The  Choice  Page  (options)

Having  successfully  presented  a  formula  to  the  system,  the  user  is

11

given  four  choices:  he  may  read  a  new  formula  (implemented  as  a  return  to  the
reinitialized  input  scanner),  replace  a  variable  by  a  value,  evaluate  or  dif-
ferentiate the  given  formula.

The  replacement  is  implemented  as  a  simple  substitution:  an  entry
in  the  value  table  is  made  and  in  all  nodes  referencing  the  variable,  the  type
and  info  fields  are  changed  to  indicate  a  value  and  to  point  to  the  entry  in
the  value-table.  All  wher  fields  to  the  right  of  the  value  must  be  updated  to
allow  for  the  display  of  the  value,  and  in  the  infix  representation  of  the
formula,  the  letter  variable  must  be  replaced  by  a  capital  letter  value
reference.  (Since  in  TUTOR  such  a  capital  letter  takes  two  six-bit  codes,
the  first  of  which  indicates  capitalization,  the  replacement  would  be  rather
cumbersome.  Instead,  each  letter  variable  in  infix  is  preceded  by  a  blank
character,  and  this  blank-code  then  can  be  overwritten  with  a  shift-code  during
replacement.)

Evaluation  is  implemented  by  presenting  the  infix  string  to  the  compute
command.  This,  of  course,  must  be  preceded  by  assigning  user-supplied  values  to
the  variables  in  the  given  formulas;  these  values  are  stored  into  those  locations
in  the  student  bank  which  the  defined  names  for  the  variables  reference.

The  differentiation  driver  requests  the  name  of  the  differentiation
variable  and  changes  the  type  field  of  all  nodes  referencing  this  variable  to
indicate  type  "x".  Control  then  passes  to  the  differentiation  driver.  A
special,  non-published  choice  is  also  available;  differentiation  can  be  requested
to  be  automatic,  bypassing  any  questioning  of  the  user.  This  is  accomplished
by  setting  the  request  control  for  the  questioning  to  a  special  value  as  will
be  outlined  below.

7.  The  Differentiation  Algorithm  (derive)

Consider  the  differentiation  situation  in  Figure  2:

12

leng      ^_

prin

**

**

1

*

tformel

item

tderiv

V

DU

DV

Figure  2.     Example  of  a  Differentiation  Situation

item  selects  a  multiplication  operator,  the  corresponding  operands  are  (x+1)
and  (x**2),  both  have  non-zero  derivatives  (assuming  a  differentiation  with
respect  to  x),  the  corresponding  differentiation  rule  is  the  product  rule

DX(U*V)  =  DU*V+U*DV   .

With  the  aid  of  the  leng  fields,  the  phrases  U,V,DU,DV  can  be  located  from  the
top  of  the  respective  stacks;  the  relevant  pointers  have  been  added  to  the
figure.

The  'algorithm*  described  in  Figure  3  will  produce  the  correct  de-
rivative (the  figure  shows  the  stepwise  results  in  the  derivative  stack  only).
To  complete  processing  the  item  '*',  we  now  have  to  infix  it  into  the  formula
stack.  This,  however,  can  be  accomplished  by  the  same  routines  which  perform
multiplication  in  the  derivative  stack  (an  overflow  of  the  formula  stack  is
impossible).

13

insert  top  formula  phrase
below  top  of  derivative

multiply  below  top  of  derivative

1,2

insert  below  top  formula  phrase
on  top  of  derivative

multiply  on  top  derivative

1,3

1

X

**

2

2

*   X

1

**

*
+

tderiv

X

**

2

2

*

X

**

*
tderi

V

X

**

2

2

*

X

X

+

1

**

*

+
tderiv

X

**

2

2

*

x  *

(x

+   1)

**

+
tderiv

X

**

2

+

2

*   X

*

(x  +  1)

add  on  top_  derivative

+
t
tderiv

1  any  operation  is  done  by  infixing  the  operator  between  the  relevant  phrases

2  a  multiplicative  factor  of  '1'  can  be  discarded.

3  since  the  precedence  of  the  second  phrase  (principal  type  '+')  is  lower
than  the  precedence  of  the  operator  to  be  infixed,  '*',  we  add  parentheses
to  the  second  phrase.

Figure  3.  Example  of  the  Differentiation  Process

14

The  differentiation  machine  is  built  from  eight  processors  for
machine  operations:

perform  the  indicated  binary  operation  to  the  indicated
stack  element  and  the  stack  element  below  as  right  and
left  operand  respectively.

perform  the  indicated  function  (unary  operator)  on  the
indicated  stack  element.

copy  the  indicated  formula  stack  element  to  the  indicated
position  in  the  derivative  stack.

create  the  indicated  node  into  the  indicated  position  in  the
derivative  stack.

The  first  six  processors  can  operate  on  either  the  function  or  the  derivative
stack--they  don't  know  the  difference.  Each  will  first  attempt  an  appropriate
simplification  (discard  additive  zero,  multiplicative  or  power  one,  evaluate
any  operation  on  two  constants,  collapse  inverse  function  composition)  and  if
it  fails,  will  infix  the  operator  and  correct  the  result  by  parenthesization  if
necessary.

These  processors  are  slightly  more  generalized  than  is  needed  for
their  application  in  the  differentiation  machine.  This  differentiation  machine
therefore  is  implemented  with  a  restricted  machine  language  to  access  the
processors:  (it  should  become  clear  below,  why  this  machine  only  seems  to
operate  on  the  derivative  stack.)

15

add
sub
mul
div
pow

top
below  top

fun

node

copy

top

below  top  _

top

below  top  .

top
below  top

of  derivative  stack

kind

top
below  top

to  derivative   from  formula
stack        stack

The  machine  selects  (by  inspection  of  the  item  node  and  DU  and  DV,
whichever  is  appropriate)  one  of  28  differentiation  situations  (compare  the
table  below).  The  machine's  driving  table  then  contains  a  differentiation
algorithm,  written  in  'machine  language',  for  each  differentiation  situation.
This  algorithm  is  executed  by  driving  the  processors  to  produce  the  derivate
phrase  corresponding  to  item.  Finally,  the  node  item  is  passed  to  one  of  the
processors  to  update  the  formula  stack,  item  is  advanced  to  the  next  node  in
the  queue,  and  the  algorithm  repeats.

Actually,  the  first  two  situations  'constant'  and  'variable'  are  recognized  by  the
machine  as  special  cases  and  are  not  handled  through  the  table--it  seemed  silly  to
add  a  special  machine  instruction  for  these  two  cases.

(A

■»->

CO  O

I—  E
3

S-  4-

E  4->

O  C

•r-  CO

+->  T3

03  E

•r-  CD

+->  Q.

C  <D

a;  T3

s-

O)  I

4-  I

4-  I

<U
+->
e  -t->

<u  e
E  ro
<u  +J
■ —  to

a.  e
E  o

•i-    u

i
i
o

CD

Q.

O  S-

•i-  <L)

4->  _Q

03

3  3

4J  E

to

0)

-Q
as

•i —

s-

03

>

X

Q

e
o

+->  CD
03  i—

•i-  JD
+->  03
E
CD  S-
S-  n
CD    >

4->

e

03
4->
l/>

E

o
u

o
II

Cvl

4->

c

03

+->

+->

to

c

e

a>

o

c

o

o

a>

4->

Q.

to

+->

E

03

X

03

c

fO

V

-O

03

+->

cn

■»->

</)

c

+J

+->

■»->

to

e

»r»

E

c

c

c

o

s-

03

03

03

o

u

o

•4->

to

to

to

u

03

u

03

E

o

o

C

o

03

+J

4-

o

CD

o

o

4J

3

a>

1—

=3

O

r^

#1

03

3

•*

A

o

a>

3

CD

S-

<u

a)

ai

a»

o

cn

S-

^—

Ol

1 —

i —

r^

O"

r—

e

e

3

c

4->

3

3

3

c

3

<u

•i—

-M

S-

•^

C

S-

S-

S-

•  r—

J-

s-

s-

O

s-

ai

J-

CD

o

3

E

o

•f—

C

c

C

o

E

4-

+->

"O

•r—

+->

+->

•r-

•r~

•  r—

+->

•  r—

E

4-

o

O

03

o

o

03

03

03

o

03

3

*i —

03

$-

JE

03

3

J=

JZ

-C

03

-E

to

-a

4-

a.

>

Q

ZD

<->

C\J

4-

CVJ

==»

>•

o

O

o
o

o

ZD

c

>
zo

+

4-

O

>

>

>

1

1

c

1

ZD

CD

CD

1

>

>

o

,—

>
ZD

CD

+

1

ZD
CD

>

o

O

ZD

ZD

>

>

4-_

zd

zd

zd

o

ID

Q

Q

Z3

o"

CD

Q

a

O

1

Q

o

>

ZD
CZl

Q
1

II

II

II

n

II

II

II

II

II

II

II

II

^

>

>

zd

>

>

O

>

ZD

+

1

*

*

*\

*^^

v —

o

>

>

ZD

*. — ■

ZD

zd

o

ZD

o

ZD

Z>

ZD

o

ZD

1

4-

D

Q

Q

CD

o

Q

O

Q

Q

Q

Q

E
O
•r—

+

1

*

O

o

*

O

O

03      1

C

3

O

E
3
4-

>-

>>

CO

"tts.

>,

•K

-««.

-14^

1

l

E

E

Ct

e

O

rv

C

0%

ft

1

1

<a

03

n)

>•
e

03
O

03

O

c

03
O

>i

C
03

C

o

>•

>1

>>

1U

"Us.

TW

>,

1^

>>

>,

e

E

o  e

0k

o

«*

fl

C

O

•V

C

E

03

fC

ro

>-

C
03

>5

e

03

>1
C
03

03

c

03

03

03

cc

CM

CO

«!»■

ld

<X>

r^.

00

cr.

O

1^

CO

CO

*d-

16

to

E
O

■M
03
3
+->
*r—
<S)

E
O

+->
03

E
CD
i-
<L)
4-
4-

Ct)

O)

JD
03

17

The  driving  table  is  currently  somewhat  wasteful ly  implemented:  it
takes  up  550  words  of  common  storage  (which  is  about  8%  of  the  entire  lesson)
and  is  read-only.  The  table  head  contains  a  list  of  28  words  pointing  to  the
beginning  of  the  table  entry  for  the  respective  differentiation  type.  An  entry
has  the  format  outlined  in  Figure  4.

code  for  display  (see  below)

#  of  words  used  to  store  structure  formula

#  of  characters  in  structure  formula
structure  formula

#  of  commands  to  follow

commands  for  this  differentiation  rule

Figure  4.  Layout  of  an  Entry  in  the  Driving  Table

A  command  consists  of  two  or  three  words  containing  the  command  type  and  the
application  information;  this  could  be  very   much  compressed,  but  at  the  time  the
above  implementation  was  considered  to  be  the  most  efficient  in  terms  of  speed.

If,  for  example,  one  would  like  to  practice  only  with  polynomials,
the  driver  routine  could  be  modified  to  recognize  fewer  types,  and  a  table
could  be  substituted  to  analyze  only,  e.g.,  xn  rather  than  f(x)  (where  the
latter  requires  the  chain  rule  in  the  structure  formula).  This  is  where  the
real  flexibility  of  this  approach  lies.  Of  course,  in  this  case  one  would  have
to  restrict  the  original  formula  inputs,  e.g.,  to  preloaded  exercises.

18

Alternatively,  the  driver  and  table  can  be  expanded  to  make  further
cases,  distinguishing,  e.g.,  sin(x)  and  sin(f(x))  as  different  differentiation
situations.  This  expansion  was  actually  (and  without  any  problem)  performed
while  the  lesson  was  written,  to  have  2  multiplication,  and  3  division,  and
exponentiation  situations.

8.  Interfacing  with  the  Student  (infix,  and  output)

The  algorithm  described  above  actually  operates  in  the  background  of
the  lesson.  In  each  differentiation  situation,  it  is  used  to  generate  internally
the  correct  (although  probably  incompletely  simplified)  solution.  The  student
is  then  provided  with  the  environment  of  the  situation  (generated  from  the  in-
fixed phrases  in  the  stacks  before  differentiation)  and  must  answer  some  ques-
tions. If  he  is  unable  or  unwilling  to  answer,  hints  or  the  complete  answer
can  be  given  from  the  differentiation  algorithm.

situation environment

x                U  =  x  or  constant
or  constant            DU  =  ?
recognized  and  handled  as  a  special  case

unary  operator  U  =  argument

DU  =  derivative  of  argument
question  for  structure
question  for  this  solution

binary  operator         U,V  =  arguments

DU,DV  =  their  derivatives
question  for  structure
question  for  this  solution

Table  5.  Displayed  Information  in  a  Differentiation  Situation

How  many  phrases  constitute  the  environment  is  controlled  by  the  dis-
play code,  the  first  entry  in  the  command  table  (this  proved  to  be  an  efficient
implementation;  it  is  not  strictly  a  necessary  table  entry).  The  correct

19

structure  formula  (e.g.,  in  situation  6  it  would  be  D(U*V)  =  DU*V  +  U*DV,  and
commutativity  must  be  observed)  is  available  from  the  command  table.  The
correct  solution  in  the  individual  case  is  available  from  the  differentiation
algorithm.

The  student  is  first  questioned  for  the  correct  structure  formula.
His  answer  is  to  involve  only  the  quantities  c  (for  a  constant),  U,  V,  DU,  and
DV  (whichever  is  appropriate).  The  answer  is  judged  by  assigning  random  values
to  c  ...  DV,  and  evaluating  the  student  answer  and  the  system-known  formula.
If  the  two  do  not  match,  in  turn  DU  and  DV  (if  appropriate)  are  set  equal  to
one  to  test  by  reevaluation  whether  the  student  may  have  forgotten  DU  or  DV.
Finally,  DU  and  DV  are  set  equal  to  zero  to  test  for  a  missing  term  in  the
structure  formula.  If  no  guess  matches,  the  student  is  invited  to  key  -help-
to  see  the  correct  answer  (if  he  does  not  wish  to  retry  his  response).

Once  a  correct  structure  is  established,  the  student  must  substitute
the  values  at  hand.  This  algebraic  answer  is  tested  with  all  randomized
variables,  and  if  it  does  not  match  the  known  correct  answer,  -help-  is
available  to  see  this  correct  answer.

The  above  seemed  to  be  a  reasonable  interface:  because  of  the  -help-
facility  the  student  cannot  get  stuck,  and  his  progress  could  be  accounted  for.
The  algorithm  is  completely  independent  of  the  presented  situation  and  is
relatively  efficient  in  execution  speed.  Suggested  improvements  include:

1)  count  the  number  of  operators  in  the  formulas  to  be
matched  to  force  a  minimal  (reasonable)  simplifica-
tion on  the  student;

2)  if  the  student's  answer  is  more  simplified  than  the
system's,  substitute;

20

3)  supply  'wrong'  answers  in  the  command  table,  and
aid-information  with  them.  This,  however,  can  only
improve  the  response  on  the  structure  formula  and
the  implemented  correction  method  seems  already  very
satisfactory;

4)  there  should  be  a  way  to  check  on  the  student's
algebra  in  substitution  ...  .

1)  is  a  minor  extension  (thanks  to  TUTOR)  and  as  such  already  planned.
2)  is  fairly  prohibitive  due  to  the  scanning  problems:  such  a  formula  then
would  have  to  follow  the  input  parser's  syntax  rather  than  the  far  more  lenient
TUTOR  syntax  in  order  to  be  readily  converted  into  the  stack.  Also,  execution
of  the  scanner  is  not  fast  enough  to  allow  this  as  a  rather  frequent  operation.

There  is  one  more  control  mechanism  to  be  described:  with  each  dif-
ferentiation situation  a  counter  is  associated.  These  counters  can  be  used  to
prohibit  the  situation  (it  may  be  unknown  to  the  student),  or  to  bypass  ques-
tioning. Currently,  questioning  is  bypassed  once  the  situation  has  been
answered  correctly  one  or  more  times,  depending  on  the  situation.  With  a
suitable  control  strategy,  rather  elaborate  learning  paths  could  be  established,
leading  (within  prescribed  examples)  from  simple  to  more  complex  differentiation
situations,  etc.  The  counters  very   definitely  eliminate  boring  repetition.

9.  Wrap  up

Once  the  differentiation  is  complete,  the  user  can  again  evaluate
formula  and  derivative  for  various  arguments,  and  he  can  retain  either  one  for
another  round  in  differentiation.  Formulas  of  medium  complexity  can  undergo
about  three  differentiations  before  the  available  stack  and  queue  space  is  ex-
hausted. The  formulas  are  retained  by  expansion  into  a  form  that  is  acceptable
to  the  second  phase  of  the  input  parser.  Formula  evaluation  is  accomplished  by
the  same  routine  as  described  in  section  6.

21

10.  Conclusions

PLATO  IV  and  TUTOR  do  allow  for  inefficient  but  rather  complex
algebraic  symbol  manipulation.  The  presented  processors,  representations,
and  techniques  are  reasonably  compact  and  versatile  in  application.  The  table
implementation  of  the  differentiator  should  be  easily  adaptable  to  a  variety  of
levels  of  experience  and  problems  for  students—it  is  at  least  being  contemplated
to  have  a  calculus  exercise  system  on  PLATO  IV  in  which  calc2  might  play  a  large
role.  Currently  under  investigation  is  the  implementation  of  a  similarly  table
driven  integration  teacher.

variable
number

first  usage

second  usage

22

third  usage

1

26

random  values  for  the
variables  a  ...  z  in
algebra-answer  judging

random  values  for
the  variables  U,V,c,
DU,DV  in  structure-
answer  judging

storage  of  the  char-
acter form  of  the
structure  answer

scratch  locations
for  subroutines,
and  some  display
controls

27

deriv

28

cpt(top  of  constant  table)

29

tderiv

30

tformel

31     dopt  (operator  to  be
differentiated)

48

56

stack  for  postorder
conversion

infix  form  of  formula
undergoing  differenti-
ation

semi -global  variables
for  differentiation:
e.g.,

item

typDU

typDV
infix  form  of  algebra-
answer,  and  argument
formulas  for  display

57

74

constant

table

0

1

2

A

15

values

75

145

node  li

sts

(see  Figure

1)

146

algebra

hel

p  request  counter

147

structui
counter

re  h

lelp  request

148     question  control  (counters
for  each  differentiation
!      situation,  and  their  re-
150 store  values)

Table  6.  Use  of  the  Student  Bank

23

REFERENCES

[1]      D.  Knuth,  "The  Art  of  Computer  Programming,"  Vol.  1,  Section  2.3.2,
1968,  Addison  Wesley,  Reading,  Mass.

[2]      A.  Schreiner,  "Computer  Calculus,"  Chapter  7,  1972,  Stipes,
Champaign.

[3]      D.  L.  Bitzer,  B.  Sherwood,  and  B.  Tenczar,  Computer-based  Science
Education,  CERL  Report  X-37,  Dec.  1972  (Computer-based  Edu-
cation Research  Laboratory,  University  of  Illinois  at  Urbana-
Champaign

24

A  Listing  of  calc2

25

ui

«VI

aJ          O

a      uj

_l

v  ui  i/)

o

_>

»-  3   <

X

o

-1  Jj

«-

a

2  <  a

r

z

►-

O  >   <J

c

)

o

2

—         UJ

<_>

O

1-fflQ

t-

o

<  2

_<  _  1/1

u

u

i

z
o

2

••

►-»-•-

1  >-

3

!J              M  *-

2   2

:

i

*—  >—

M

►-t           — i  r>-

ui  3  uj

«■    CO

»-

*-        (\  r»>  r»

a  o  3

L

j

X    UJ

ua

tu      •-  <\i  r>-

UJ  U  J

(

5

•  *  3

UJ

x       n  (v  (vi

u.       <

2

r

>-

z

O   WO

3

i       ru  -« -•

u.  >c  >

"

I

■

21

o

z  «•

O

►-        —  n>>  tvj

►-.  Ufl

a

H- 1

O  2  O

•-•  1/1  ^*  (M  f-i

O  -*■  ©

c

>

o

»-

-1   «•   2

UJ

x  ui  <\i  rvj  (M

V         Z

2

<

<   1  •->

M

<    (J    rt    rt    rt

(/>    *  t-i

►

O

i

o

t-

-i

►-*

3  m  <v  <\i

»-  U- ■  »-

1

►m

•

<                             y-

=:»-•-

<

_»

I  »-  (VI  —•  -<

-•3  2

<

K

•

Ul                                        2

O  •-•  UJ

»-•

4

1-   <J-<  AJ  (VI

XI  _l  3

3

:      <

X                                       UJ.

>■  r  ui

X

-1

—•  o  '-^  — ♦  — #

4  3

c

i      -•

*—  >■

•                                       X

oi

o

»-

j  a  -  -im

>C  >  O

C

J          V-

UJ

O                                          UJ

Ul    •  X

»-

*-■           11          *—    — «   — 4  -H
2         ■*  Q.  1/5  "<  ft  (VI

2

50

in                          u.

>  1>1

3

(/>  ►-     •

s

J         UJ

<t

in

u.

1          o  z  o

»-

i-«          »-  _)         •-.   (V  •—

C   UJ  X

i      a  _i

I

»

HN

O                     3T    UJ  >-i

II   UJ    •  •■<  (V  o

_l  l/l  UJ

'

C          UJ  4

a

UJ

Q

f^

a  o

z

«*•          K    I    «•    O  .-•  (\)          CllJJ    O

u.  «->

_j

>

V

_i  a  >

o

*         l/>               — i  O  O      .    X  X  (/>   UJ

U.    X

<  >-

>-■

O

n

_i  <a  <t

•—I

Ul  :*    rf»    O    *     *    e*                 Z    t-

)       ►-<  o

UJ

•-

z

M

-  i

►-

tlf    *    C    • t    —   Q    <    •-

a          t.  »-■  (v  it  •  _j      x

o  »-

UJ  (/I

<

•-•

V

3          O

«

3

U                 Z>

x  a

>

o

ir

II   »-

*-»

UJ  ^;  r»   —   «    ♦    QL     •  O  ►—   >-h
»    4^1-KI-Jll-TU    I

1         <S  K-

m

»-  <

»— •

»— 1

—    1    2

»-

r>

2

r»-

i

a

1—

c

«-  (_  ►-,  p.

2

f-

V  ~«  —  it.  tr.  tr  u.  tr       in  z

a

(-       ►-  (/:

V

IS  o

uj  uj          (/

c

«-

X    «-    I    bP

it<

>

Hi?  n-  'i)  lo  if  u.  e*  a   •—

-

u  ~

^

■^      •

c  ?           ~

c

r

LU*      <      f—

I

B

J         U.

j.          /(ll    II    D  D  D    i  D  N  I

•-<  _i

v

•-  o

*  *■*          _

a

z

33   <-    O   Z

V

Li-

r

_1

_JI-lJJ<C/,;CCBC«fcC.-0     1

u.u;±aau,uiuxuj      <->   i
ia  irajjairijQ:*-        i

>-  aJ  lii  al  UJ  -  —   —   ud  ■—  — iU-
ZXjtlIZZZX2    E  -h   f>-

1

k;         h-   3

IT

<V           O   _l

fV   (V  _l

a

«■

2    U.   Z   •-

-J

lt.

Lu

U

J          C~  O

<_>   IT  <   <

u  u  uo        «_

c        o

r

O  LU    «-           G

o

u.

u

l/l

<

J           <S    _J

IT

JMOOI-      1

J  JI--«

r-     —     —

2.

>-i  i.    z  u.    <

r

—4

1-

>

r        X  <

t—

<-«_••-•  _l

<t  <t   *  a        4

r>  OJ  XI

"r,

CO  ill    «-    -*    UJ

X

>

3

«

.       x

C

J        X  u

<

o  "i  «-  r>  <    ;

U  <->  O  -•  rvj  <_

O    —I   ^-.  —

rvi  x    i    «-  i

i/)

►-

»-  •-    1

t-

»-

a

<

A

>

Uj  2     \
UJ  (/!  O  M

s:       t-  a  u  u

a  2

1  <I  O                     Ul

UJ

Ul  u

UJ

v

►-

3

'

\-  7.    ~Z         U'  H

U.'          »-

1-   u

t-           t-

■^

y>

"•   X  Ul               -I

r      «  4  >-  4

JljM             -SI   — <

IVJ             ^H

-•  3

-i           <

J

Ul

2   xl  ul             <                ,

1

ii  t-  x  E  _i  q

ui  o  o  i—  "h  a

i-.^(X  —

X  <

i-  X                jj

*  *

*  *

*  r>

4

0

-   c

;i-it          u           '*««*

J  *  *   »  ♦
i  t   n  o

•j

^  -j  -+  \J  *>  j

1  X  U  ~)<   W3
JD   O   "-  JD   >    3

(/)    <I    3  <

-t  m  n  *

3    Q

n  o

.  <  3                Z   «

<-  a  >   3  -t  m

~*  ~1

■-4     -•     -»      -•     -<      -^

-»'-«  \J  \l

1

1

1    MM   M    M

"W

T"*

*  *  *  *

*  *

*  *  *   Jl  n  n

UI

Z

U-
Uj

O

UJ

z

tr
o

xl
z1

<

o

c

2

a>

4

z

■

Z

ru

O

UJ

1/1

t-

t/>

«a

Ui

o

r.

u.  a

-  >  ?  tu

— >  <  I       ii  »-

-J  J  I       lf«-H                ^«

uj  a.  i  uj  ui  a       ui
i-  ►»  *  <~  »-

o  o  Ci  o  ac  h  a

Z>   «      •    'i

n  o   ^  t

a

0  •-■<  _l

1  <  o

n   o  o  a   o   0

I    i

z
oil  ui

'   *
1/1  UJ

♦  u.

ft!

JO

3  15

♦  Z
u  -•

♦  u

c  jm         o
—  *   t-     I    «-<

•  <t  o        •

*■!    ♦     <3  (V

uj  o  a       uj
»-  4  a  ir i  f-

<  l<>  <*  a

O  -hot:

■1    ll   U.J         4
1    I-  »-  T    L/p

O    k4  •-•  O    3

«/>  »  s  <.  </>

n  a  -  a  >

o  a  o  o  o  *

:    o  >»  «

c^:

CM

UJ
2
»-<
U.

UJ

D

27

2-

o

<x

2

in

x
uj

UJ

UJ

UJ

a

**
m

o

f-  r»

■»>  •*

OS

r-  r-
f»-  r>-

k  r-

*-  r-
r-  r>-

M     II

a  a
i  i
a  a

M  X

tf>  IT

in          o

*»  un  r>  r>

,  IT  —i  —  O

n.  >c  ><

iw  (o  no

0«M«

M    MWO

o  *  *.  *

*:  *:  *:  o

O  U>  IT  </

l/>   (/)(/>  o

c  <a  <i  <
o  X  I   2

«  <s  <:  o

2:  3:  x  o

o  a*  ^  *

rt   4  MO

3  3>   T>  -i

v  #  t  o

r-       — -v

rvj  '■n  -»>  iff

f-    *    ■/»    #

*>•*>/»  *:

M/l  ITl/

</>  in  in  i/)

r^  -J  -  -

j  _i  -i  _i  <

r«-  u  o  <_

u  o  u  2:

r>-  Iff  **■   <A

^   +  J.  if.

o  c  o  c  o

U.   U-  u.  u.  u.
25    2  2  2    2

«-■«-«-«-*•«-«-♦>

./><_>  o  —••_>  ♦    ii  h.

a  •{  *■  *■  *  «-  4-  «J  «-

o

ZljUU-'UJUJIiIUIIillJ

aaaaaaaa

2>    >->>>>->•>-

N  N  N  NN    fg  (VI  IN

s  r»j  n.  n»  rv.   rsjivirsi  z  M  *-  tff  iff  iff  iff

r*  —  —  —  —    —  ~~   C  *    t#    iff  tff  tff  ^
f-  2   2  2  2    Z  Z  2   ~

l~-       II       II       II       II         II       II       II      I-  -£

N  —  —  —  — <;  II

r"v~''x,v^.-v,si^    ■+  rv  <*i  •*  if .  O

BNMSNNNN3:il     II     II    II  no

a  —  ~-  —  — ex  <_,  e  ~  >  «

UiCu.^C(X0:2li.|j;liJ|iJliJlt;Uj

iua«-2<«3u,za  c.  a  a  a  a

J2>IUlllM>.>>.y).>.

'si'-ii-Qi-J-iari       1—  »^  »—  1 —  • —  k-

>•        1

<         !        r«-       r-
«-  r>  s  4  «■  2  «■  */"<  *  2  «-  «-

r>~  3  ~  «-  C.  2»-«i/)C2<  ►-  U
«-«-         2K>-il/)00<ll-f0uJ

hi  r  1  _i  nii^>  ■*  o  <  »-•«  L>'io
r>-  r»-  nn  f.  r^  r»-  r*-  h~  r-  K  s  r-

4-«-«-«-«*«-«-«-«-«-*«-«-

u;  ii.  uiujujLl;ujujujuju.'U-Uj
a  a  aaq.aaa.aa.aaa
>>->>->-  >>->>v  »->>-

l-l->-»-H-»-»-»-H-r-hAl-l—

O— •Mro-frlfl.Or—    CDO*    O  — «

o  #— ••-«#— 1*4<#— »*-^i— tf-*«-i^»fvifv
11   n   11   11   ■    it  11   ii   n  11    11   11   h

-•  o

o  a

UJ  UJ

a  a

>-  >-

o  A  -*  0

uii'i/;<i04i-<i  o  in

U'  UJ  U.'  UJ  UJ  U.'  U.1  ut  UJ   UJ  UJ

aaaaaaaa a aa
>-v>*->->->->->-    >->->-

a

3   3

o>o-i\j-n*/i    o^-

tj  t<  t>   -.  \j  -m
-<  -t  \j  \j  \i  \j

T)      »)     T)      »J      Y)      TJ      ^Y)

\J      \J      M      \J      \J      >J

4  U.  u

4

s  s  r-

X  1  I  ►-  c
o  i  </>  2  a  uj
in  -  o  ■<  o  a
sjr  'U  >— i/)  >-
r^  r  ^-  h-  r>-  »-
*-*■*■*■*■    ll

li.    L.'  L.'    U  U.    N

a.  a  a  o.  a  nj

(V)  <

•  _i

-4  =>

•I  X

a  a

o

uj  u       <

o

z

UJ

a
o

u.

H      >       2       >i

in      «->  uj  •-•

<i  1-  >       _j  <i|

3|  j  4  h  r  j

ID  3>  I-    O   X    >-H

■  x  ►-<  «  n-  a  a

u,  c  a  >  4  o  u.

OIIiJh        u.  c

u.  a  a  1-

2j  UJ  2   2    2

0  2  2  Qalwi-i

t—  « >-i      a

h  00  c:  a  a  a

dlillii  ZDOC

ac|  teen  iijui-i-
3  aaoo no

r   r  r  r   r   r   r

l/1     1/^     t/3     if      f\     tf     |/*i     iy>     if

(\j  rvli

w»  v>.  w  a

et    tf     tf      Q

-*  in  -o  11
<\j  rv;  rvi  —

II     II      II    IV

u  k  3  —  r>

Illf  UjC
UJ  uj  UJ  ^  >  5

a  a  a  _i  <  o
>->-><  _i  c

!-(-»->  U

J»

-.    \|  -r)

*  ■*  #

o  *  r>  £>  ^  o
n  jn  n  n  n  />

<
Id
o

o
o

X
ui

z

2

x  Lil
lii  (/I  I

O   <[  l

z  a  :

w  ul  '

ai
i  i

aJ   2  I

r\j  r\j  not

29

o

rs

33

UJ
Q

t3

-)
33
UJ
Q

O

2

(t

a:

UJ

U

«

»-  O
«-   -.

I

r
c

3

o
o
©

AJ
V

UJ
(/>

Z»
£

►-  *•
o  x

-1  Uj

_l  Q

<  Z

X  i-i

—  r>-
«-

*  •
t^  o  r^
*  z  «■
«j  <r  t,

-j  j°  _i  ii
«      <x<J

I-  a:   ug

—  C  —  2

-  x.  ~  c

*-    UJ    *■  h-

o

O

4

Cl

u.

C>

(II

UJ

\'

t-

Ul.

o

1/

z

2

Z

UJ

1

UJ

I/-

(J>

z

3:

a:

UJ
'l'U

o  o
z  z

k  <  —  *  a

O    31   >E    Uj  t—   if!

z  uj  ■vj  or  «■  — <

•+  •-  <  t-  a
pzuu  J
(V  •-•  or  o  =*

fifixir.  iifivcjl:

Z  Z>

uj  uj  ■-  c

*-  irif  ;  at;

—  'X   A  —  T  i-

J*  uj  _l  C  ~>15   *

*>  »  /i  a  —  cl  >
<vi  f\)  im  rvi  (\i  oj  t\j

z       z>

U.I  111  w     ,    o

in  ir.  c  o.  c
-•  ir  <  fi  o  •-•  <  .n  -*  z  —
z  j'j  ji-ioiijjqDo
:zi-z»-M<*5:j*uj_ici>-j<r*

►M  UJ  |

Lu

*

uj      uy      uj  jk  uj
«-l       uj-<si/>       k       f-ca  •-

Tiuni-^:       >4       -•  x  3        !  -t

-3H-_JG0OUJ<:*«:*<*(/jO«i3t

:3

o<3

r\j  st

1/1

3

o

3

*-'E  i-  r  -it  i-  c

<W<l/)<lfl<Ut2*    *    *

I

=  ->\j»)*   n  x>  --  o  j|  •=  •-•  m      -o   {  no  ■-  a  j>  3i\j,)*/id^B>  r>-»\J-o*-jr)o>-t»
rvj  c\j  t\j  cv  (\j  r\j(\j{M(vioii(\if\j{\j      <\j  (|j(\jru<M(\i<viryojrurvjr\irvi^irvjrvi(vjrvif\i(\)cvif\jfvir\j(\j(v

3^
U

T  -3  —  \J
J3  XJ  X>  XJ
rvi  c\;  rvi  M

■M

a

c  c  a        i

3   3    O

no:  -
UJ  UJ  at   02

Q   O   W   M

<\l  O   ■"-•

~>  s.  z

1   z

UJ         mi
if)  UJ  »-   >a  U-        H-it-i

^(/><t-U  -1  _  o  _i  -3  "-
i4  d  <  4  i-3  1-  I  r  3TCX
UJtrZC    _J<3:<3tin3t<n:!:

n  a  «•  a  y
000000-

■V  (M  flj  N  rg  r\i

rvj  rsi  r/vj  rvi  r\j  <v

if'

if  a
1/1  c
I   ^   o  a.  —

O    Wl/I  UJ

r  ar't-n
ID  31  in  <  x  <:

--    X)   >    J  -»  X
*-    X>   X>    X

fvj  (\i  rvi  r\t  rvi  r\

U         II  •

r      >        >

at  if  •—  un  •— •  irt

c  if  a  if  ct  u,

U.   (*)  UJ  -thJX

»—  -*  3  ^  O

II  •

J   Kil

U.    Lf    Ui
*-    fit-

<
»-
<

c
«-

«-

X

«-  ;

if  uj  a

if  «-  1

in  z  c

<V  *•  <-

UJ

15

3  >-■  '  3  •-«  o  >-•
1:  i-i'M  r  ho:  1-  r  hi
tfl  <  ar  <ifl<:j<(/)«ij

*>   *  n    o  »-

D    O    D      D    D

rvj  oj  oj  oj  oj

3»  Ol  -•
JO  31  3>
<\l  (\|  (\j

£

l\J  «V

>
UJ

a
<«

o
3:

•3      •

a  <r

o  uj

U.  Q

3

n  o I  *■  o  j>  -
j>  >•  j»  >  >  =

ru  rvi   oj  rvi  rvj

UJ

c
o

z

<

UJ

-J
<r

z
<

4

*s  e

'     .-

f\j   l/t>

3
QT'
UJ

3

O
i

1-1      1    D

Z    Kit   I

3  «*   in

^3S

f*)  rt  ro

3

3
CD

UJ

3
»
>

3
GC
UJ

a

>
3

■!<_>■•<

()BO
C>3  O
,*(T  —
*  UJ  ■*
-J,  3  —
•*    •  •#

Q    —  3

—  1

4<uj  -*<
a  c
>-   £

X

•  If
Q.O

X  U
UJ    »

•  r
z  z
_J*»

•  in
3  •
+•  o

1  m

•  u

:j

-  UJ

111  <n

■h  »-

•  a
1  u

•  ♦

•  z

•  <

>  h-

•  Z!
o  <

•  I— I
o  *

;£

o  <*;   •

O     •    A

•  l/l     A

-*  a    ac

—  2  Q    -

3  "-I  C    C

T  in  in  5
~  <r   .  -

ui  «  x  u;
a.  z  2  a
>->-<<>
•—  c/1  1—  t-

u

UJ

o  >-'Or

3  <    3C

:.  o   p   3  a  -|.H

n>  c>  r»)  (*>(*)  o  c

ac  -4

-  ^  3*1

f)  C)  (*)  f)  (*)

31

o

4

UJ

a

■>

<_>

o

13

3

3

■a

K

Ul

UJ

3

a

'  -3  ->  \j  •>  at
'•  M  \i  \i  \i  >l

ii  1  m  f)  n  r«>

3
<r

t/)

Q

or

D

r

UJ   UJ

i/>  u.

3  to

Z    2
I    4

»-   X
3   — >
-I
-I   UJ

4  X
2   »-

•-    >-

O  X  >-

r  »-  4

—  UJ  *-
a  irt
>-
a  c  o

4

5
C

2
Ul

-I

4
-)

Z

a

o

u.

I    «

r  id

UJ
2

I    Z

4    >—

x  a  o      x  t-  a

O  3  — '  C  UJ  4  <
O  ll>  I  Z  Z  3  uJ
»-  X   «V  «-     I     IX

^4

3   *

4

z

4
»-
l/l

»-       z
z       o

4  t-  O

t-  2
(S)  4  U.
2  »-  O
O  (/>
O  2  Q.

o  o

5  O  K
O

>—  .►-  <tt
— 1<  *»

o  .u
n  z

Srt    V    (V

rv

— '      ♦

*    UJ
o  —  _l

*      -HID

~   •»    4

iJUJl-

_l  -I  *

X   G   t-

44a
o

z

4
•-
(/I

Z

o
u

I
a

UJ

IT*

4

3
Si

a.

Q

u_

_J

4

~t      a:

jj   _}  UJ

l/!U)ct-
4  *4  ~>   Z

X  <a   o  uj

•>  >  >  UJ  3

x
c

H

<5

a
u
a

c

>

a

3  4

S     5

4  «-
X

If)

U)   If!   U

_» »—  <a

X  z
4  4  *

X  </l

4  2   S
■>  O
u  *

a

O  _f  II

a.  4

M  >-•

•  «->    I

•  UJ
4   a

-1    M

ro  ftci  ci  o  rt  (*>

■>■  n  o  «-

4       y~>~>\)-r\tno

UJ

(A  UJ
4  71

r  4

UJ  x  *

tori's   -•  m  ■n  *  n

t  **/i  nn/i  n/i
(*t  f")  o  c|  w  rtf)  fn  i*>

X      I  UJ

f-  2         t-
c  >-)        ■-«

z  o!  h-  r

•-.  "f  4    3

n  n  rt  n  a  -d
(*)  r*>  n  f>  f*>

vi  *i  t  n  x>  ~-

o  o  o  a  o  £

f1!  fl  Cl  H  O  (*

<
UJ:
X

uJ

X

<3
o

in
o

o

z

tn

z  ■«
<

<  X

•  o
z  in

<  •

►-  x

in  <i

o»-
u   •

<  r

•  to

in  o

o  u
o  •

•  X

z  z

•-I  •-<

in  in

«   ••
•■  u

Z  l/>

►-  CJ

</   •■

•  u

Q.  UJ
>.   \T)

Ui    «■

•  t—

z  o

II

_i  a)

3

~  r  c  <

ox  —  *

O  O  -*    ■*

rvj  u  (\j  *

•-f  ui    ♦

>
4  o

«-

•4  UJ  N

o
z  x

UJ  o

a
ui   II

3|  O  —

Z

a  2

•4  a  o

*

n  in  in

«■

*  >

<i  •->

>-  x  <-a

4]     <]

x

<t

X

:>
o

Q

<I

o
z

UI

o
z   H

O  it

M  O      •

o  o  O    «■

■vj  in  z   •

3       in
c  o  o  c

X   Z  UJ  •-
X  3  a  3

•jjuia

y     3  -.   \J     »>*|no~-T>»>3

!  •*     'X

♦  UJ

I-  o

m  a

»

Z  I  lA
UJ

A  —

t-  '    r-t

z  ♦

3  h-

o  ;  z

'O  |3

->  o

—  o

I  ~5

O  Cj

II  Ui  7

i-  qui

Z  Z  \.

— '  •"*  ^-

C  I-  o

<_  ?    2.

->  o  —

—  uz

UJ

a      z  uz

i        -•    3uJ
3   *  O    3  X    K

!E  O  ">    ~?  O   C

-  v»  i    t-  n   x
o  o   o  xi  o   z

fO  f*)  f*>   <*)  f*>

-3

x  m

UJ'i-f

z

•-•  ai

o  o

a.  u

r4>

I

•-      I

a

£3

in  mi
t  a
r\j  <*

o  ui

Z    X

o  a

_»  s

u  T>  *

-  ol  >
o  a  n

o  r*>  r*i

*   *

1

z

o

13
in

z

xi
uJ

i

i

♦
•t

H>        —
(VUI  (\

q  in  o

X  —

ui  z

Oi-

*  *

>

t   CVJ  «\J

—  41  ♦

-«   -H  «V  ■

♦  o  o  ;

(VjS  I
O   *    «  I

i  A  e\.

»31i

3

■

~3

n  n

—  —

— •  -*

1

<*-*

<
x

33

2
N

«■
X

o

z

X  X

2

*>    X    X

Q

a;

llJ
_l
(VJ

2    U.  "J  _»U
«   t\J   M  '\J  M

•     »      »    •     »

«-«■«-  *  «-

II      II     H     II      II

a.

U_

*-
z

<-
ii
r

rvj  <v  oj  (v

O  O

3:  x

2  2

*»  trt

Q  o

2  2

<  <J

»  «

*■  *•

o  o

»  »

«-  ♦

II  II

O  0

Z  X

—  ^»

2  2

O  Ci

2  2

<I  <

«-  *•

O  1/1

»   *

II  II

t\J

<\J

♦

a  •»

^*

ijj  #— •

o

1-  o

5

i-  z

—

Lu    —

2

_»  2

<A

f\J  tA

X

Q

•  a

•  2

2

oc  o

O   <

•3

(VI  «*

rvi  »-

■*

•>  *

•

p

»  «-

•  X

«-

«-  <

«•  o

x:

2    -

r  u.

»

•  +■

»

«-

<-     II

♦  *>

ii

ii   r

ii  *>

h    it    ii   ii  ^rt__|^,^^on»

— ,-4.~        ♦■»♦♦■+   +    +♦«-

c  c

Z  X

2  2

x  r

2  2

I  <    <  <

J  -E    1  TC

'-    u    L  X

?!*

X

-  »•  u

-  C  _l

3  ►-  O

c  c  c  uj
x  a  x  2

o

2  2  2   f

t

r  r  -  i

2  2f  O  2
<  <t  a.  <«

X  X  —  X
I  •*  2  X"

X
LlJ

CCCQOCCC     •

xszazzzx*

2    222J2222    —

xirirrxx*

(JUUU)U(JUU(V.
22ZZZ2Z2C
<<acr<t<<<J<iZ

xxxxxxxx—

1IIUII1    332

O  UJ

»->

rv»

aj

r\j

»

•

•

•i

«

»

«-

«-

«-

t/>

2

z

•>

•

•>

«-

«-

*-

ii

11

II

<n

m

ro

♦

♦

♦

-^

—4

~-t

o

e

O

X

Z

z

2

2

z

t«

feP

&

o  o

a

2

z

z

4

d

«

VI

if-

tf

w

»

■

•4    «-

«•

<-

o

»-#

<

X    -

»

•

UJ  «-

«-

«-

t-  11

II

ii

1—      M.

—

—*

UJ  <\i

<\l

f>j

_J  *

♦

••

•— <

•— •

•— «

a  o

o

o

o  X

X

X

^*

—-

— r

to  z

z

Z

Z  trt

tfl

.«

O  0  Q

1

»-  2

z

t-  <t

<I

<

CJ  <*

*»

t^l

T      -

»

•1

5  «-

«-

♦■

X.  o

/)

t—

•

«

•

(#•  «-

*-

*■

w  ii

II

n

x*i+i*r*

<J  (V  o  (V  CJ  (V  «->  (V
202C20ZO
<X=IZ'CZ<I3

i  —  a-^x—x  —

x  2  x  2  x  2|

r  ■♦  i  oj  o  '
a  o.  o   I  ■£

2  C   2    t    *

<  X  <I  •-  — '  i

X  —  X  o  o  <

X  2   X  X  Z

♦    ■»•    ♦■  —

o  o  c  c  _,

X   X   X  X  2

— C

2   2    2   2  (•

♦

x  r  r  —  t

f*  UJ  u  o  <_>  «\  c.  >

iVIZ    2    2  O

<i  <  >a  <t  Z  <

xxx  X.  —  a

iaJ  X   X   X  2  1

111

_l         _)

C\J  -«

a
>  r  ■  <r
-•ox

U.    u.    _l

1*1  fl  (*1

(V

ft.

a  uj  a  a

,-«\i'»*no~-o>^>-(vj-04-nDUaj>->-«\j'i*n-o^t)»o-»M'o<-jri

(Vi

x  z  q

D    —    O    W    "3  -i

n  i>  r»  n  o  o

-j  *  «   «

m  ■n  *   n

O    O  O    £1

X

— i  o

E    _l

z  <

3  «  O

lu

UJ
_l
<\J

»-  n  >  ■>  -«
o  t>  o  •*-  •*•
•*  <t  •*  ■*  ■*

<*»

o

34

z

X

o

O  •-!

>    m  >  •

li.    t.  u.  *.     IT

ro  —  ei  ->

♦  i

I    *  I    ♦  o  c

<_>  (VI  O   (V  S3

Z  O  Z   O  *

<3  2.   «S    X  —

x  —  a  —  oc

izxz  i:

3  — •     i.

rg       rsi   r>

K

X

UJ

-  z;

*-    3
—  l±.

*-  a

II   c

r  u.

-•  M

c  en  o
X  e>  Q

—  *  «t  —  —
Z  —  rr  ♦  ♦
— •  f\
r    iiivool?

U\U    I    2    2

z  o  z  *  *   «  o

«3  X   <  —  ^-c  (VJ  <l

x  -aaoqaj
x  z  x  x  x  x  x

»   n  o  ■>-    ©  *
♦  *■<»♦    *  -a

a

x  z  o

uj  o  «■

■>  -»  m  n  »  n
o  o  o  a  o  d
*  ■*  ■*  4  *  •*

c  a

O

<V<

m
a

-■J

IX

:  x

(VI

a
f\j

4-  «•

O

X

o

o

2

■•v
<  <s

uJ  .1

a  uj

«  r
z  z

3  —

z      z

M   >C  M  X

O    »  Q    •■

z  o  z  *:

<*   (VI    <(\J

zzzo

»  •  »  m
«-«-«■«-
II    II     II    II

(VI  (VI  (VI  (Vjl
♦     •»     ♦     ♦

odod
x  xi  x  X

~   ~1   ~-   ~4

Z  Z  zz\

oq  od
zz  zz
<  <r   <«x

«     «    «k    »

«■«-«■*■

M  Hi    O  UJ

m  •  •  «
«-«■«-  *

II    II     II    li

♦    +     *    ♦  —

c  q  oO

I    X   I  X

z  z  z  z

r  r  n

o  u  <_>u
z  z  z  z

<.<!<!<

x  x  x  x

n  3D    XCD

o  u.

X

—  c

Z  C

*

—  X

o

2  «a

—  a:

z  cc

IT

_l
<V<

3  -•  vj  i    *  n

>    >     >  >    >  7*

*     t    *    ♦    <t  *

I
Z  I  O  |  z
►_        UJ     :  •-■

Ice     'X'
I  o      o

1  u_         U.

«-  *■         «■

O  *         X

*3   '3  *  1
«-  q  «-  q  «-  »-

*.  u,  t  ui  t

—  f)  —  &>  —  f

♦  r  ♦  I  ♦  x

<V   <J  «V    U  «V   <_'

o  z  o  z  c
i   <i  <i

—  x  —  x  —  a
z  op  Z  X  Z  J

(VI

(VJ

O  ~-  O  T(  ->  ->
>  7<  7>  »  O  3
<»    J-    <»     ^  lf>  IT

a
a

35

:>
o

z

UJ

a

i

c

S

u-

*  \u

-  o

v.  »

-  3
!  I

w

C   «*   U

o      z
^^otioo^ooo

*       ti        *        *
x  ~  i  -»    X  —  X—  X  -.   I   *
oivi^oi.ucv.orv  «_>  n.  <_>  <v
z  c  zd   ZOZOZOZO

X   Z   BZ   I  Z    X

1    <r  i-i

a   -h

3     O
UJ    O

Q

UJ

O

jJ

■uJ

X

UJ

•a

UJ

a

X)

<

i
i
i

«-
or

«-

«-

<
«-
U.

«-

r
«-

UJ

«■
i-
«-
(/>

«-
>-
*-

r  >  -j

\i  "O   *    o    o  ~.  D

*>  i  »i  i  »>  -n  »>

I   M   \J    ^    11      .

ijfl  mininiflj">ininininj>in  j)  min  in

7<     7>    -•    M

>-  X

a  —  a  a  a  a  a

i-  r  t  x  %  x  s

z  z   r  3  Z>  D  a

UJJ  •-.  "5  ~>  -}  ->  ~)

T    1   -t  M   *)    *

*  n  n  n  n  n

irtininininini/iminininin

a  a

u  I

*  *

n  o  »»  ti  r

(in/inn

»-  uu      un

ZO.t-  X    X

r   Z    <  UJ  3   UU

3   -»  -y    •>    *

o  £>  jd  o  o  c

ji  in  if)  /i  in

a
o
a
x

a>  a
a  ~5  *

o  •»  o  >  r>
o  o  o  .o  *-
in  in  in  sjtv  in

z

I/)
I/)
UJ

a
a

1/1

4
JJ

I

I
I

a

«-
o

UJ  «

x  •-  a

<  >->  «-

»-  a  —>  a

Z  3  o  *

>  y  o  uj  i

in  a  •-,  o*  *■  -1

UJ  UJ    u
t-    UJ  (/>  t-     l>>

-•  zs       <&  —  <x
z  o  t-  x  <a  ■  a

r>  s  <  uj  3  u;

->  M  •)   *  r>   c

■»-      Vf   >~     n_    ^      >.

in  in  m  in  in  ir

o

<r

UI

a

LlJ

z
o

o
a

a

o
a
cr

III
if  a

ui

r
■d

GL  ~>  •    •

~    13    >   3    -«
M  «-    »-     O  »

A  m  in  in  in

I
I

I

«-

a

^  u.>  *

ae

«-

14.1  a  -«      rt.

?o      «-

CJ)  Jj    3         Ui  O

2:  a  ~«o*  «-  -*

UJ  IU  UJ  u.

UJ  l/IH    ITil/

"3         4  -•  <    Z>

o  y  r  <x  a  -t
2  <  uj  3  ui  a

">   »  n  o  --   o

D    O    O    O    O     S

in  ir>  in  in  in  1/1

UI

37

o

a

IS)

o
a.

<\i

u.

Q.

J-     -J

-  >  UJ
«-•  (J

3      K  «

a.     4  X  X

-  >  <

3     —  X  ►-

a  o  x

*      Ui  UJ

*   o  zz

Q

2

<r
a

a

o

i
<s

X

(/I

X

J  u.       o

•

t  O  UJ  t-

X

M

c

•   I  >-  C

>-

ewe

2

*—  t

.►-

._,— ,3

IT

J  O  UJ  2

::

I  3   S>  <t

■  •

*■   *  if-

••

■  *   <*  *

i

-

2

^          1

^

1  t-  X

a

.  a  t

*

J  UI  1/

H

■    t    «g    ^

5

-  >  .-  c

—  tr.  —

<i

a  x  >-

•i

J31

.-J

'   3   U.  C5

1      5-<\l1l     t
■0  ■©   *  ■£)   X)

(A)  Q

(/)

UJ  rf»

(_>  <A

C

X  X

a      <

t-

O  X  2
2  <  V

<»-in

2  •
»-  >  IT
O  (/>  ^
al  •  U.'
X  <\J  X
X  ^*  •
Jj  u.  *

a  ~
*  •  u.
¥»  -<  a

u.  m
x  ^
•  u.
x  a

<\*  r«-
•-c  rv.

X      t       I

t— t  *~  —
U.   »-  *-

t-  a  a

If  J5
O  —  —
X    2  2

n  o  -   n  >  -a
■nil  t  *)  *

.O   >£>    £>    -O   -D    -O

UJ

J.

3  »

-•  M

*    *

2    <_>  2   2   H-

—  _|  _,  >-«  _t

O   4  D   3   <S

tp-JO  -}  ->  uj  »    *

*  n  o  •»  n
*****

X  ^  X   >c  ^o

3    -«    "VJ

r>  /i  ji

%©  ^  ,£)

*l  n  a
n  i)  t\  s]  r>  i\  n

*0   ^t1  >0   •£•  sO   "X)  ■&

OJ

1/  t.

2?  ~
C    <->

.J   D
_J

<v  ■-<
■fl

c  >

2    *-•
C     I

o

D  -•
I  —
I/I    2

>  o

si

2  —

u.    2  <J!

a

UJ

a

q

a

al.

3'
uj

X

3  o
*   *

i  a

JOJ

<  r  *

-i  v(  *>  *  n  o

**.  ^*  >~  .**  ^. .  *

sD    sO    -£)  *£>    >£)   s£

«M

U.

u.

a.

m

u_
a.

fM

o

r

M

1
C
I-
<

a:

UJ

a.

o

UJ

a

x

v1

J  UJ

••  x

M"3  O

-H      I      *  t

-H     _J   UJ  —

u..  «  i  o

a.  >  x  x

3    3*  O

r>  >  ^>    -<  v

«■  -  o    o  c
oca    i.  4

u.

o  a

a  uj

HI  I

jf

X>   n

£  .0

O

«A  u.

ui

a
c

u

f  Ui

—  a  *

U.    <

a  ui

z      x

3*3

no-

ODD

err

UJ

a

UJ

a

i/»  x

Uli  i—i

*»  o

a

2

^S

•-  X

a  t   1

X  •-    H

—  a.  a
Z  x  _i

O  (VJ  »-

~-  «  a

•  LJ   _'

—  I     *

»-  3  *~

a  t  a

I  #£

z  »  z

<

3   _l

r  <

ID  >    Ji>
<L   <C   <G     •£  <C

(A
O

a

x
o
cr

u.

o

3  <

~>  uj  *  *

<c  4

O  —  T>  *
>  >  >  >
<C  vO   •£    >£

t

.*

a  x

3  2  3

zoa

3  --H    J.
O     I
U.   tft  *>

.1"

a    i

5  !

a

5

111

4,  a  K

U,    X   Oi
Q|    UISN

t-         Ot-

I

<«    X
*    CJ  UJ  «

n  a

£3

!

•u

i

H

i/>  -    I

as  a  -    p

u)  s  o  ♦

>UJC

a  a  -i

a

►4   UJ

-J  3

^X1

32

*!   °

39

x

o

X

C

<\J

u.
a

e>

z

K

•— *

M

Q

3

a

-1

UJ

O

J

z

JPM

4

•— .

Q

'-U

Z  <A

3

l

r>  <a

Z

o

«*

u

U_  X

«

o

W

•  X

»-

M

><

WD

o
a

x
o

(v  a:

o

-i  *-

UJ  4
>  X
UJ  uJ
-I  X

o

z

X

c

X
O

a

IA

3
X
a

lUJ

i/.  a

in  «  «  *

■t)     1«     "i    -«     \J

"vi  \<  n  o   ~
f«»  k  k  r>-  r^

—  uj
»v  a

U.   X

X   UJ

o  —
x  t
<  a
x  a
_i  o

<

H  <A

X  u_
£   X

'iJ

I   .0
3    X

*.  ►*
UJ   U.

r  z

o

(V  LU

a'  a

u.  x
X  uJ

<Vi
X

a

X

o

■—   ♦

x  a  x
c  c  s:

h-  y-  t
</;  tn  i-
a

2  Z  £

UJ

—  a
m  o

ii  <_>

n  x

a  uj

x  >

a  o     —

—  .       u.

I  --     ►-

•HIIA       in
»  X        o

<a  h      a

3  o  x

x  -•   •  s:

»  A    OO
«  X

►-  —  i—  u.
~Q.1T  W

it  x   i  a  *■■

I  —   II    UJ  (A

II  Z  X   >
11}   X    «  X    o

4  x  x  o  ►

■♦*  ox  —  o

♦  —  c  i  c

I  I    1  -<  _J

4  -•  -*  ♦  a

■»  ♦  m  v  c

4'x'    •  2  u

-1  r  >-  x  a

l/J  r-i  q;  _
t    i  a.  3  w  t    f

x  t  c   tMi/ia
o  -*   i   uj  _j  z>  c

»-  O  (\l  I   2  X   I—

in  r

J-  U    2

♦  *  *

i  *■  n    o  -  q
n  i   ■"    *)  ■»>  i

t--  r-  r«-  r«-  r-r->

2    O

2  D   <X

3  *   ~>  O

2   *-  O  2    V-

»-•»—(  _J  »—••— •  i

O    X    <i  3    X  !

-5  3  O  ~>   UJ  *

1    O

1    ♦

-«   vj  -n  *   n
*  ♦  *  *   *

r»  r»-  h  r-  r-

r^  ^■

3    -•

jn  rt

"VJ    *>     J

4

r>  q  ^  u  >
n  n  D  n  n

t^  t*  t-  r»  r^

3  —i  —  (/)  rg

<  a  o:o  x

o  a  o   o  o
t^  r*  t^  p»-  r^

in
u.
x

<

3
UJ
X
X

o

X

X
UJ

UJ

<_>

z

UJ

Q
UJ

u

UJ
X
X

*

z
<x
o

a

•  i  >-
v  -+  x
f-       «»

X  »-f  z

1  u  •-<

UJ  UJ  XI

a

W  X  M

tA   UJ  tA

X   tA  X

O         O

•  <,  »

«J  X    A

<.  a  —
t-a  (v

(A    ^H      I

V     -  X

r  o  o

X     A    t-

o  x  to

►-  X  ~

IA  C  2

XXX

U  U  <J

2  2  2

<    <  <
XXX

X   Xi  CD

UJ

o

z

UJ

a
uj  >

0»i

UJ  o
X

X  <A

UJ'

ia  a

*  UJ

<_>

x  uj

*  X

a  a

ru_i

N  3  -

~  z  r»

-►  I

rvj  ia  ii

•  w  —

X  r\

a  x  i

Hi  «  x

(A   C  C

2

-*

ts

I

1

It

z

•4

. ,

Ma

^

(VJ

♦

►4

1

r-

2  X

I

>-t>

o

C

A

t—

Q

^1

y>

=*

(\J

—

UJ

>s

2

X

«%

(A

>

X

D

f-

X

Z

♦

o

<3

~~

II

IA

a

^^

n

•

•4-

i

a

^t-

n

c

1-

xH

3

X

o

V

z

I  X

i

O  u

i

2  2

a

<*

<a

a

X

x  a

JD

a

_

3    «   *

X

I  <_>  c

!  -J  ^

■*  -•

-~     X)    T    3

r>-  r+  f^  r-

^  n  o

r^  r>*  h-  r>-  »>-  i*

3
a

3

o

3

3
O

3

a

3

o

3
a.

o

o  i  a  uj*-
►   a. »-  a  a

->-  —   »-     —     ^  Q

O     *    ~3    -*     "\J

-  »»  o  n    o

f-  r-  r-  »*-    t^

*  n  o  -  n  t>
one  o  d  o
r^  r«  r^  r»  f^

i/i

3
Q

3    »    *    *    «
>    >    J»   >    >

r-  f-  r>i  r--  r-

H  U

OH

*   *  *

l)>c

o  *

r>  o1  >»  t>  *•

>   7»    7>  >    J»

r-  t-  r-

1

(Ml

u.  Z

4*

*   «

3   4
I?

:J

—  a

iLl    4       *     4

0  H    O    J«    ^>  -

3   3     3  3    -»•-
(D  OQ    3D  30    00   0(

3**

CJI-f  l±J  t-

H  — '  "-f  "*  **

UUJ  X

*  n  o

oo  a>  cd  a

or

<r
z
u

Vr

z

»-
in

a.

(A

,.

r1  c
a  t
I-  r»
3  O
O  X

z  «
3  «_

cooc

*  •  *

41

n

■

*
o

(V

*  Z

-I

V

2

0    m.  _

•  r>  —
o  o  u

•  ♦  -I

O  (VI  OJ

3  ♦  O

XuX

•  _l  —

—  (E  >
«  «t    »

-•l

o  ~  n

X  >  O

♦  w     J

lf»  (/)   «

O  13  rvt  (V

z  «  ♦   «

■^        •  ~   iiJ    «

-  Z  J  o

:    ►-  _i  i  r

o,  —  <*  *

♦     X  »-  (-    .-

3    -  Z  -  o

O    Z  -i  >  n

3

in

U   (_>  3  O

>  J3  J

o  <  r  4

x  «  x  x  a

*  n
■»>'»i

x  co

x
— «
u.

z

X

q  «  «   «   «  «

■^    1   1    ♦    *     *

x  co  co  x  x  x

T  *    *
.  •— i  —

a  z  a

3.  —  X

ku  cp  <vj  a

I-,  a.  x  a
-iuj  a

i.Q.     •

r+   >-    ^4
-     *

Z    —    I     —

K-.    C>    —  O

i/>  3   t-  x

..  a    •

•->  <  <\l  CO   x
X  X    —  X

mh   3:    h*  La-  i— <

u.  c  u-  a.  u.
Z  r  z  >-  z

<—   IS)   m<  H-   i-«

2         C    O    3
D  *    £i  ~>  C)  *

x  x

*  *

X

or
o

w        u

or
x       uj

0  >^  >

1  •»  o
V)  (/)   •»

3  «-•

n°  x  i/l

1    t«     •  X

~  *

»-  o

a  -•

|       Z  ~

I        *  -■
rvi>-  x

1    X  if,    *

|     H    MU)

U.    J3
Z   2  X

►-   UJ  z

UJ   ■— c

UJ  W    UJ

X  <*  _J

*X  X

»    *     «

/>    X   ^    O    7>

jn  ft  n  is  n

CO    Cf)   CD   CD  CD

2>-i  M

o  o  o   q

X  X  X   X

o  c
x  a

z  z

<  <

r-  _l

(/I  CO

z

Q  *»

u  t*

X  w

to  JD

X

w    •

(/>  —

t-

a.

X

♦

C/t    X

3    •
X   UJ

4a

—  o
K  o

o.  or
x  uj

— »  •>
z  o

-*  M-l
♦  if)

—  X

»-  -H
0,  O

CD  **
->    A

a  ~

u-  (VI

z  ♦

l-l  (/

o
u

or
o

u.

n  t-t-  x

X     (/)(/)    t

M     MM    1^

U.    JJD
Z    Z  Z  X

M     UJ  UJ  Z

»— «     >     X     M«

z  o  a  o;

D   X  Z  ~>  «(

O   »-  T>    >

o  a  o  o   ~~
CO   X  X    X  X

X  X

5?      -

a,

u.

or  t-t  mi
a  i

V)   (/

o»      :

:  or  x
moo

'  u.

<*
«  a

UitA    X

»■  —

a  x  z

X    *

f-

»«  x   »-■

tfl    M-l    Q

u_  a
z  -
'  ►*  c
-  u
<v

,    A    H

M-     Cl

0.         '■

4  X  >-
X   —  U)

mC^
U.  U.  _l
Z   Z    2

mZIi!

m-<    MH     i

r  o .a
D  ->  a

W  -*   r

ep  x  c{i

z
o

42

in

X

x
•-•
u.

z

UJ
Or
O
O

or

O:

a
o

■r      r.  -r.

O  1

»n

A

<  rt  ■*  <\»

ujf  i  ^  r*

cm  tr  e>

i   -*  i    •  _i  3:

«»-■■♦►»

VI    1    *   />    O    «•

a  x>  o  xj  x>  z

CO  (X)   CO  CO   CO   3

o  •

X  (A

(/>  o

X

M  -

»-

a
x

o

X   «/

X

U  _I3

z  c

^  3

-•  v<
co  a

U)2

»— *

a

(35   CD  CQ

O1  O    O-

I   W

Z>  4     *

-.0>   «    -«

UJ  UJ    UJ    Ul

i/)h  ift  u;a

ac  oq  cr  ^  o

UJ

-5  *     *

s

Z

o

UJ

UJ
Nl
<-iC

«/) .
rr  c
o  c
u.  i

*  il

ub  •*»'

4    OKlt
3  lit  X  >

<    [f.

X  4  O  M  <

ui  a>  ~>  <  o

O   0>   Of  0>  O    CM

43

a
o

O

x

I
U.

c

3)

a
c

a

a?
cc

im

*  -

X-

z

4

x-

u.

IT

>

Z

o

o

a

o

4

UJ

X

>

4

UJ

o

_l

X-

CD

3

X-

<t

£

UJ

1

a:

_1

1

4

o

1

4

3

u.

4

SI

J

£

jj

o

z

3

cr

X

X-

o

£

o

fc—

•-•

X

u.

UJ

►-

C

UJ

~ J

<

u.

UJ

X-

_»

a;

X

4

4

uJ

3

1-

H*

=►

a

UJ

X"

c

z

UJ

X-

2
Ul

4

UJ

4

4

a

Z

-1

z>

UJ

(5

I

c

_J

u.

x-

<r

«3

1

u.

VI

_l

UJ

>

»— 1

(/)

►—

a

Uj

c

<

<

•r  ^

>

f— '

r-~

p— *

=1

.x  <

X

u

Q

•

ID

>-

£

D-  .-

a;  —

"><   3   4    3

UJ

u
»— «
o

X

o

z

4

I

«-
«-        —1  xt

Ul

a
L

-L.    I
1/

w
ki

2'         O
Ci         »-

•4      a
o

Ql  »

c       o

c  r  or
o  «-

3$

uj

a
o

X

Q.

a

c

4
UJ

f\J  iT.  tV»

cr

c

X
X

Z    4    llilb

x:    t  i/i  ♦

C  •+  <1  •"
O  0   X  O

-«z  a

c

UJ

i/i

3

WD     -

a

h-   O

r-   O

r»  _i

r-    -

r«-  rvi
r-  v
(*-•  r

x-  w*

f-  cr

f-    4
t»  Q.

r»  t

h-  ^

r-  a

x»  4  »-

o  a  x

*  —  <

x«    ,H    X    l/)

taut

—  4    Z   —

a  a  4  a.

4    —  X   =1

x  2  x  a

3

£
a
o

z

3

o

o

z

X

a
x  c

uj   •

x-  i»-

r-
o  f»
Z  x-

f-

w  r«-

%•>  r»

x-

x  r-

»  i
rvj )
x  r
>—  if
a  r»
o  *■

*  r-
or-
uj  c
a  ii
>-  r
x-  —

A   -

r  »i
—  a

^  <t
a  o
<j  -
a  q

UJ

a  i-l

^

x   X-  x
34   i

^\IM    \|1
(>  Cr   O  Cr

T

Cf  0»

1*     />X)"t)>r>
1»)     1  »)   -Tl    »»  *1    *

UJ   3         V)  X   X

u)  c  iiuu  *-a
4  x  z  nix  r  s:

x-XXDX^Z3

4UJ4   JI/1J   hT»    A

I

OCJf   OCfOCJ'OCf^O^CI*

a
_i  a

J|   4   _l

n  i\  n  n  r>  s

O  0*   O"  O  ^   0'

.-  UJ
X  V)
x-   <

x  a
3  UJ

>-  x

x  •-  u>'  a  oi

x.    £  CO  £   S
Z    Z  4  O   3

UJ  >-hX  "J  -i

->  -•  \l  »»  *

ox)  d  o  o
i>  (7>  cj»  cr  cf

o
a

>

o       x

o       -

Z         X  x

x*  x-  a

a  a  4

o  ox

o  «

_l  ->    UJ

•-•  N

•f         Cti-I

*x  UJ  V

C   Q

v  o

♦  <
x  -

<■   2

a  rt
t  a
tv

x

t-   4  <

x  x  a  <a

OlU-3

$  X

x-  U

1-4  x

>-   X  -

a  —  a  a

—    C  E    £

Z    Z  3  O

Ui  •-«  ")

£>   *-  O    T
O    O   O   O

CT>  0>  0>    CI

a

x<
■H
(VJ

or     I  M     i

4  ♦

X         UJ
IX

IT         U.  J
(VI         X  4

ir-  ii   3  ;>|

*

X

<-
u

o  «-

rr  z
=>  «-

UJ  C/l  u

x-'X  —  i£   x-  X
4   3  4  1/743

o->  \il*  n

"  ^  ^  -I  *-  *-

&■  Cfl  (3>  o\  o  o>

a
o

Q.
O

a.
o

z
<

uj

a.

at

3     2

a
o

U.

a
»  c

(/)  c

2  «i
o  *
l-l  f\

I-  □
a
o  ol

(V
a

<a
0.  <*>

*   «

»-     St

2    ■<  .-

D  (t<:

>    ^

X

O

a

o

UJ

M

a1

<n  »-

xi  2

•I  ~

•-*  c

oc  a.

3    2

'  a

V    Z>

an  »-

j   uj

z  a

3

U

CD

<
o  •
O    CV

•c   I

■H-  —

(\J  —

a:  a

a  «*

a  a

UJI

o*  o  2
ii  *  m  i-<

UJ   UJ  UJ  UJ

t-  1/1   *-  (/)  I-

-•  <t  --  a:  -<

l  U      f  'J  -I      «

vj  *>  *   n  o  --

ti  o  »   a  o  o

z

l/V

a

X

2

i,

-t  o  I

O    I

U-     !

_  •»  XI
<  f\j  «-|
a  a  a
—■  <t  *-'

at  a  or;

U.     ♦    «"!

2  oa;

~  «1

=

m  >

~-»  ~*  uJ

o  a
x  o
x  M
<  «/}

UJ

*-<x|  *
i  <  3j  o  •

j  a  >  >  >  >
o»  o  o*>  O  O  O

X'
UJ

I  X

(/J

2

UJ       o

Q  (/■  I-

I-  <

ax      a

o  u.1      O

h  hii)  k:
•-•  E  ;/)  o

2  2. •<!<■<

3  *4  od  tr

n  a1  »-  o

O  ON  <3>    0>

o

u.

2

UJ

u

X

a

u
c

G

U.

Z

UJ

X'

u.

X

~«  -J

Z.  4.

V:

Ml
i  o

1^1

*JJ

n

45

a
o

a

a.
o

a

o

o

UJ

j

u

z

UJ

i

X

*^

a

>■*

•*

a

;

iv

0

3

H

z

u

«T

-»

0

1

C

l/l

0

a  o

ui3

u.

»-  X

UJ

x  z  a

<  —

z

H    H     IfcaJ

oa

0

1

U.         1

lv

a  uj

■X

X  X

—                 X

3  X

X

:<J

trt  (/l

to               .-«

X

•-H

a

»•-•►"

u.

h-  •-•  <?

u.

T     !

:«  •

_i ;

1 0           »-

z*   •

z

■4

x  ma

zxiuin

0  z  a

»-•

u,

:T

►  <

•-«  •  a  0

000

1

A

UK.

1  z  a  >  a

.-.  0

z

*  1

*:   V)JU   <  l-

*>  z  _l  —

1    ►-•

3          !

;M  '

o  uj  <j  cl  oj  a  d       2

O          jOWkll         C
-It**           *-    >»-           •-
••   *»  «4  _l     »  O                 H

<A  3  (VI  UJ

X

w  •  X

UJ

X

M

>  x  a*  >  3

u

«9

•  •

»   UJ    •—    —

<  -

■  •-  3

W-*

•  3

~             1  «  >          2!          <

uj  r  x  x

-1  «

li-X

m

■  •

n  ii  j

X  'Uj

i-<    •

l/)

»*:   Ul

*  IM  O  I

Ul   N

1  I  <

3

3    X

a   •-  I  3   >-    3d  3   J
-J  Z  M         1-3

O    <fl    V     X

X   -

y>  _»

;\j

X

X-  O

Z  *  —  •-<

u

i      3

«

«-  I

—   C  IN  **    M          «»>*

(VW  ^*

*«  a

:  »*  2

a

<3

*4

ao^^-ww'-"^

•  4  *  z

W  C

1  w  a

•

_J

H-  C

u>  ->4      »-  *»      a

u.  I    IT   O

u

0

■*

3

0.     2

1*1/10.               UJ
—  I    O^M  J           rtQ-

>  t    11  2   1   ~       *  a  ir

I   <A  O  •-•

1

•       u.

0

2

_|i

*   —  X  Z

(/

1  x    •

£

X

4:   *

vir  t3a:

■

I     •  <VJ

•

O

C-  <*■

•-  u,    r  «t  2  u.      »-ra

f  C   IT.  »   <   ►

'                           I

t  ^»  •♦

in

u.

4.

a  X  w-f  uj  a      iirr

to  x  0  (M  uj  a

1

a  m

3   ~

•>

2     —

ujxoM-ii-       jo-

3  «■*  X   O  1/1  u

u

>   r-   C

X     -H

#— 1

-       1

c  --  2  a  2  »-       *  2  a

x    2    ~  X  —   C

~

-  a  2

-     O

«•)

"

A    -           _i    T                   if           U

<       «        /

,                                                            X

a  »

3    T

♦

x:  o

r  0  i  -

01      0  r  r

I   —  I    —  X     1

4

>  «  ►-

_l    II

1-

C;  I

M  u;

u.  u  ■—  2  L.  3

O  X    <J  K  <J   »-

cv  r

!  a  (/1

^  3  <»

^  (/

2:   t

a  <v  2  1

*  2  *  *  z   1

z  c  z  a  2  Q

c  c

3    A    >-<

020

«'   ^^

<5:   -r

J  C   <fl

i-*<*K«l/<*i     <3    _J   <    _,

I                                 t—    H

-    (V    _l

k  a  x

H-    _l

*     O

-SIC
1    -  1  2

0  x  0  a  x  3

X  —  X   —  X    ~*

a  a.

0  z

a  0  —

*s

to  J

z  x  x  _i  x  x

X   2    X  Z   X     1

)   X  —

O  a.    1

J

a       a

UJ

I  »-

•

*-  2  0

<      c

X    2           H

00       -1

»-

Z   U.'

*-  UJ  1-

t-   U.I

i

•-1  •-•  _l

^i      3

»— 1

•-1  ^

—<>■•-•

►H    •>

<  D:<*

-n      j

2    3          *
1             (VI  —        u

2  o  00

z  0  x

z  0

U

UJ  3   O

-      rx

1  *    4    *          3   £

3  3  X    #    4    *

3  2   UJ  «    «    4

3  2    «

*    4

*3  >

0  -•    \J  *1

*  j>  0  >-  x>  >

3  -«  \i  *i  ♦   /I

0  *>•  r>       ?>  -

-•    \|  "O    *    />

O  «-   D    >    3   -

Vi*
0  0  0

/»   JJ

Ml  \4

11    11
O  O  O  s

1  1  n  i  1  i

*****  ■  ■»

*  *  *       *  r

ji  />  n  />  n

J1J1  Jl   Jl   «{

O   £

es  o

00000  a

O   O  O   O  O   G

000       0  c

1  0  0  0  0  0

O  O  O   O    O   C

,        000

O  O

a
c

0

0

0

r-

0

r-

0

r>-

0

r^

0

(S

r>-

z

O

0

>-4

r^-

f^-

z

0

0

0

r^

x  r^

•-I

0

•■  0

»-

r^-

a  k

y>

0

o  01

i

1^

O  r^

0

_J  O

a

f^

^   M

0

•  *:

-I

h-

^  Wl

-I

0

•»    4

<

0

1/1    tfl

w

r^

UJ  -~

on

a

3   CV1

«s

f-

O    ♦

X

0

UJt-

>-

•^

X    U1

'XI

O

V    UJ

&

r  3

w

*

—  0

tA

W1

*,  UJ

<

♦  a

^

^-1  ~-

*a

0  z

*- »

X     t.

t  — .

»-  C

--  (V

(/I  2

0  ♦

UJ

^j

2  »-

3

2

—  »/l

c?

t

uj

UJ

^-

X  3

O

a

0  a

<-

*

C

Z   Ul

1—

^~

s

<  X

X

0

X  —

0

22

x>  z

a

h-

u

c

— <

_i  0

2

<t

_i

3

0

1-

1  »  *  ♦

1

I>

>

3

■H    Ml  4IJ1

0

£i

—

"  ^3  *"  ^"^
0  0!  0  a  0

0

0

c

u_

i

uj

X

in

z
o

z

Ui

or  "u

uj  ♦

U.    UJ

•h  o  u_  a

t-  ^"  M     111

0-  >»  Q   I

o  <\j  «-   x

<  41-ir  ia

*  »>  "»  n  »  x
o  a  o  a  o

x

o
z

o

»

rvj
v

r
a

<

a

<

</)  <*

a  •

<i  «-

u  I

o  v

z  Uj

•  *
o  *
z  a

*  «j
o  >

O  Q

Z  «-

(/)  UJ  UJ

U  ID   I/)  O

jj  O    <t  _l

a.  *:  o  ac  ■*

</)  O  ~>  uj  o

»>  *•  n  3  ^  o
d  a  o  sos
o  o  o  a  o

47

UI

>

<r
ui
a

UI

a.
o

x

ui
o

UI

z

5

e

2

M
(/)

I  <*

i  *

V       c
*      I-

o

<

Ui       c

at      •-

ui  in  o

•  a  <  >-

«*  a:  x

|J*  uio

a  z

2  •-<  </i

O  _J   U

If  Ui    >  ■*  »•)  *

t-  »-    IT  m  C!  C .

Oj  <t    2  <0  ^  XiC

CJ  3    <*  <VJ  C)  <\J  2

U)  X

t-  -a

12  3

*>  »  n  o  *4  x>
i  n  *)  i  »i  i
-<-•-<  -H  -t  ^

ui

>

i/)  a:

up

U,    Uj  ~

in  <t  »-

2  a:  x  <*

4U  O  01  ~  <\l   «-    *■

1

•—  UI    Srf

»h  E  I/)

2  2  <

3  -•  CC

2  a

UI    IS  l/i

l/i  Oil
-•  _i  <  x  u
o  J  t-  a  x  a

~5   I  <  UI    <  Vt

a
ui
3  a

10   2

2  o  a

<  "*>  2    »

UI

>

X

at

1  z

i   2

»

a   i

O  H-

U.   «•

X

r»  «-

«-   U>

t/1  «•

UI  2

>-  «•

^  I

«-
I

f.  >■  </)

—  (2  —

ooi-  a

-•  -g  »>
.n  n  .n  .n

VI  Ul         UJ

ff  mUI  I-         t/)

t-  r  i/)  x       <t

2   Z  <  UJ  >-  '«•►-  X

Uji->X2<;uj<2

»  \j  ri

*  *  n  n  n  n  n       ADflniiodoii

r^4     f^     ^      H     >-4     rH     r4t  ~4f— lr-*~-4~-j**-4^.-«— •

r>  &  —  o  t
j>  /)  n  n  j-

(M

UJ

>

49

<
_i
a

CO

a.
«o

«-
c
«■
a.

3

UJ

_l

CD

<
I

a

J  x  r        ;ol     .       Lu

«  —  a.  if  «-iy|i  of

i  u.  i-  ©       •  =j  -.  i

*  23  i^3Q>  Mf

<  "O   -f  4Qfl    I

UJ

tO  f-

a

3  x

«-  o

Q  a

«-  a.

3    3

i  r

to  to

(V   —
>    >-

<  <S

uj  a

O  2

f-

—  UJ

o  a

z  >4

UJ

_l

I  -

■>  *■

t-k  C

<z  «-

u

O  »-

f-  U

>

X  M
UJ

o  -•

f-

«  Q

f-  <

a.  -

<  2

•  »■+  x

^  a

i  a

~  1

>-
<
_i
a.
to

x

UJ
Q

Q

«-
a

c
«-
a
>

UJ

>

«-

1

*•  O  f-

d  s  x  r

*  *  •-  x

«/)(/i-H  —  ai>-u.f-

—<  —>  ~>  o

3  3  -.  Z

CO   Z    3

X      HH     C

►1-H

O

33h(IHI-il
~  3    «    »<(/><»    3

J    — >

->  ~)  »   «    4
_  o  *  n  xi  -  t3

lO*)**.*.*-*^**.  *»Jl/)J1J1flJ1/)/lli

<VM(\j<vjrvJtvj(\jrufM(\jfv|       (vj(vi(vi(vj(v)(vj(M(Vjrvj(vjfli

"*   **    "^    *^^   #-•   »^    **  »■*  ^*    *^  _«^H~^^~-4.-«^~-«r^~-i_-«

»>    3-«\J«>*-f>0>-

f-  Z  U  <J

•-.  -.  _J  _J

z  o^  <  3

3  ~>  <  U  CJ

©  >    :>  -t

*  *    /I  n

(VI  (VI  (VI  (VJ

<  (V,

>•  >

<J  <

_J  _1

a  a

to  </)

_j  x

UJ  4

x  s

X  *

O  X

u.  O

f-  3    x4

II    o  ^

II         I-   3  xi        :>  o     ►

t  •-a    ui  «•  rv>  a  ot

»—U.f-C  •<3^-f-

Hrf^>oiJiz:  cr  3  o  >  a*  a

f   rt<   X"Ci   -f   4Qh<i

a

<t

z

•-  X

X  <

a.  r

>  X

*-  a

C  3t

«-  *

Q.  1-.

>-  to  Z   3

I-  ><  hh  O

X  3

Z   3    f+X'f«-<         3  3]    f-X  >-  rf-'<

<  wt  <  i.

3  3  <[  3  «t  O

3    -<

(\1  (VI

UJ

f-         O

>-•         _l

Z  3

3  "I   <  3t

i  *  n  o^  i

O    O    O    O  O   £

f\i  t\>  (vi  ru  (vj  op

4

3         _l

j>  a  -<  V  n  *

(VI  (VI  (VI  (VI  (VI  (V

Z  Z

3  a

3  ->*

/»  a  <-  »  >

^.  ^- . "»»  ^h  ■•*-   x

(VI  (VI  (VJ  (VI  (VJ  (V

p

VI  Ol  4-

b  r>  Xi
(VI  (VJ  (VI  (V

— 1-

50

UJ

z

M

o

•3

itr
uj

IQ

a  uj

UJ  i

«A  UJ

m  c

i

o
it-
er i
uj  :

:>
o

o

z

uj
3-

El

»-

a:

►-  ■>■

o  >

o  <

—    _l   2

*•  a.  *

#(U  or  Q

3  O

C!   h-  5    3

*     t.  t    1

O

o  a  a.  </

_l    <     X    w

U

a>

~  a.  z
r  x  ■-
z  r>o  <*

•-  -j  ->  <_i

—  a  j»  a  -.
a  a  a  >  >
rvi  rvi  rvi  rvj  rvi  rv

N

>«    (/!
~  XL

u.  cr
z  o

-•  J*

z  u

o  <

->  u  ♦   «    *

•n  *•  n  o  ~-
>  ?•  7>  r  t
rvj  rvj  oj  rvi  rvi

</)

O  d  —  -<

»»4

a.

a  >  3  -•  \j  n

J>    >  3    t>     3  3

rvj  rgro  rl  nm

*  *

■»  ii   u  ~~  */  ^'

3  31    3  31    3    3
**)  CH   r*>  f*1  r*)  r*)

a  s:

z  uj  a

«j  -«  u

IJC

z  a  t-
a  a  4

ujx

<r

H-    (_>

~      O

rv  -   o

(*)   f*  f)  r*

i,

:4

z

4

ii

o

IS

-1
m

!  «
»— •

a

i  *
>

I  z
!  o

<
is
o
a.
a.

z
•-»

M

ul        I

Z
-

r
i

cr

«-
c

Ui.     •

ouj  -

en  a  t.  <J|>o  uj

C    tt

~i  2;     •   3*-^   «-

3^i

o  o  c

3'Ht-ii:
a.  ca<

a  a]

^      t

rvj  \i  \j  \  m  \J  m  vj  v,  "i

en  (•)rocir*)»"»«r)f*jr,>  f

1
UJ

sd

i  uj  o    .o  i  >

>  a
vis:

->  "VI

n  -o

en  en

^

b

a

t-  <
a

o  r
c  c

>

z
o
cr

»>  *  n  x
n  5*)

(*>  do  c

51

a

J
uj

r

(VI

UJ

U.

x

UJ

2

V
«-
V

4-
I

a
«■
_i
«•

lu

«■
I

C

eu.
—  u.
•oz

-« «-

hi

-  a
<  s

O  C

a

_i

u.-  ^

I  UJ    .

*     l-t-

(/■  •-   Q

a  o  a;  c

_i  «i  _r  c

ji  n  u

i-<a.niJi««*  *-  z

►-  UJ   <J  lu   3*

mi/)   J  3  3

Z   <  <    H  O   I

z  o

i

h-

«-  Q

x  u

*-  c

LU  V

■x.

a

3

LU

a.
a

>•
at

UJ

o
cr
a.

UJ

">  *  *  *

"t  tn  <t  n

*>  ■»-  n  x>  »»

>   3    -•    VI    1    4

♦  n  />  n/ir

f*>  (*)  (*>  w  ro

*  «

CVJ    O     «-  M
<*>  (Vi     I    •-»

>    1'  LU    UJ

2M    lu  (/Di-

ll/) <t     -•

a  z  =<  »-  'oc  r  i

LiJ  «->   CO  <t  UJ    3

rvs

h-    UJ

i/i  </:

Lu   <I

13  or

o  uj

H-    >->  Ul    0-

/>  x>  «-  a
ji  /)  n  i\

E  yi

3     -»  \i    1     *

•o  o  o  a  o
n  (*)  (^>  i*>  <*)  co

o  a

4  .o  «-

<s
a.

>

2

z

UJ
(X
UJ

UJ

r

z
o

2
3

«v       </■    *

o.  c •  i-i  a
_i  —     <

UJ  J1  D  >
T  -<  *  3

l/i

a    i

UJ  <f   UJ
_  O  -•

cr

UJ
2
UJ

£2

-<  M  -n  #  n  -o
c*)  (*>  r*>  **i  ro

a

>->  —  a

Q    O  >

«-    >*■  C

3
O

»*)  c>  o

-l    M
X)    S

3    -

«-   —

—  3
</>   «-

O   —

u  u
•  (/)

2  X
*-•  *-
(/)  o
<  UJ

•*«/)  —

_    •  3

3  —   «-

«J  =>   -
T  «■  h-

Z  ~  aq
Hi-a

IflOW

^l

X   2    Z

UJ    f.    <

-§  —  ad

•»  2  «rt

3  <  q

f  t-U

I    •  •♦

— t  «-    4-i
I    —   —

er>  H  W)  i  lt
o  a  o  z  --<  i

O  O  <J  M  O
<VJ  Q  <  t/^  fVJ   >-

x>  x>

x>  x>i  ■»•  v

B    »    D     D

<*)  n  ro

z
o

52

in

UJ

o

M

a

r

CM

a  i

x  cm

o  ■>

-t  —

4   -4

53

UJ

r

o

•*  in

i  -

4.  Ji  vf

■a.  -  n  -    I

or-  »i

'Z  f\J  "  t\j    |

UJ

:*

(/)

a

«  <  <i

o  w  II

ao~  2q>«    ii  tas

~»  z  _i  q  o  3  >

UJ      ?       </>
►"      •-      i-      ccoa  oaaa  a-  at      »-      oa

^•-I-xq-TOaO    3  3  O  D    DO0D0:DO3i-<H<l-JC*:<i    d

1  *  <  »  <  jqo  -j»/>  ocaoi  aaQ  a  cc  at  m  a  o.  a  v>  «/b  <i  3  oo  o
***   <t*«»finnnj)Aiinnji££/j£ii/jj)ooo

->  *  #  «

*  n  o  ~-

c  n   c  n

r—  ci  r^  u  ao  ir

$

ui

Ul

■r-od  i-  ii-iac

3  <  UJ   <  CO    «t

7>    3    -•

z
o

M

UJ

Z

o
o

z
<

V)

«-

UJ

o
a

3

l/)

UJ

a

x

UJ

I

</>

Ul

>l

•

>-*

— •

t— 1

<  i

4.

•

>

<

•■*

>

a

«-

2

ii

jj

>

i

♦

X  3

*-  O

«n  o

•""*    UJ

cr
Z  O
z>  *-
in  u

4

4u

U.    UJ

o  or

uj  z

>  4    •

»-       -1

<  t-    3

>  z  x
•-•  <i

a  i-  f-

ui  </>  u

mora

I  O   Q

t-  ui  «_>  a
a  i        a:

O  t-   <    0_
O    «-     *     *■

no   —  o  >    3
>  >  >  >  >  3

♦  •♦  -«  *  ■♦  in

UJ

«n

z> ;

c

o

z  c

4  U

a
i-  c
z  »-

4  u

z
ou
u  a

UJ  2
X  <

D  2

O  <

T  t/

O  2

U  L

«-    «

Z

o
a

X

UJ

3

«— *  •

•  x  •
i-  «*  •
z  > »-

UJ         <
•Z    -  I

aiOuJt-

ja.m
3  x  «*  t-

(X  U!  CC  ZJ

o
t-  ►-  »-  aa
z  z  z  4

u  -at  <r

>-•  t-  i-  >-
»-  «/5  i/i  a

o  z  ?  a:

3  c  c  o

3  u  o  in

«-*•  +  «-

uj>       c
</)       u

3    3    3    3   3

u>  in  in  m  in

*  *

o  •»

3    3

>n  i/i

x

CO

►-.  UJ

r  i/>

z  <

►-»  CC

3-1

-•  -•  i

irtm

«■

*   Zj    O  «■

(T   4  —  li

a.  *   -  a  *

_l  r*   ~  r*  2

ui  a  o  (vj  o  *■

i  j;  ifiw  i

1

i

bi

ui

U)      1    IT

>     I  •*  t-
o  k  x  a

Z     i  ui

m  irj  tn  in  in  u  i

UJ

in

J.

-t  -\J  -\J

in  it  in

z  <\

o  c

Q  <

z  c.

4

X  ^

-    <*

3  c:

\j  »l
\j  \
in  n

x

CO

jj

ID

_!
<

e
i  v>  e

M

It

5

a

c

<  -
i-   i

l/l    c
u  -

3  <

*  *

in  iri

a

J

r

55

Jtx

z>
a

.11

o

o

•h
»•>

V

a

c

O    H    c

r«       — '01  m

w        ui      J       u

»jt«     k     a     w

I  -IOC  -<i|  .-X

:.

:'   >-D5«   a  -j\j  •>  *  n  o

•ji"*"*****^.*^  -.  _.  _

£

i/i

3   X
3   uj

Nl
•-•

o

o

z

^  o

►  o

<6  r»

e\i  *

-•  m

%  ~

(V  v

as  ►-

<  a

a  c

»  o

s  -

C     I

O    IP

UJ  •a

h-  a

acujoa— '■«  rviouo  <
Jifiira  <vi  r»-  z

in
a
<i
<_>

2
O
i— •

to
(/)

UJ

X

a.
x

UJ

as

Z>
C

x

3

I    _l

41

X  «A

a  m
u

(\J

uj  a

u  <

as  a.

O    •■  x

U.   >l

tft

«-

«•

O    "V

I/)  Up  O    ♦

a:  i-t-  -h  a.
o  H  —  «•

x  yi

x

o

■i  ro  3
(Nl  -*  O

>    DC

i  m  w  uja

Z    2<iOjJH
UJ  mOD   I  I

UJ  3         (/)     I    UJ         UJ  Uj

y>  ceuai      k  i/ia

«  I   2JH  O         — i  ."O   Z

x  o  X  3.Q.  zj  »-  —  'X  *  D

/1  n  n  jj  n  £   o  a  o.oodo>-^^*v-*1>~'',*^s-^x   n  n  r>  o   r>  n

un  m  in  un  in  uh  in  un  ununinuniniftinuhinun  in  in  m  m  in  ji  in  in  in  in  m  ur

1

or
o  o
t  3

<v   •  <

a  _)!■
<<  <t  c

X    >•    <
UJ

uji »-»  *

V    UJ

h-  *
Q-  X
C  «
C   I

m

<
Q-
-t
Ui

X

c

I  c

in  uj

2!

1

uJ

rvj  -1

UJ  3  UJ  UJ

i/ia  o  a  a       Kva       ut

■X  T  -|  "^  S       H  3  SB       <

X  D  <S  ^<D"3    ►-  K    «*:>»-  X

uj->c  <_>  u  ->  <  3f  a  ~>  <i

->  *

x>  a

>  ■

in  in  in  in       if

•4^  H>^  f-

2    d

2    3

0.

a

a

UJ    1

en  m

<a  i-f  o
X

V  N   C

Hj  •-•  O

a.  <n

a  *.  ¥
_t  x  a
o  o

•I   3     3

<\l  m  (\i
«vl  ^  (V)

<n  x

UJ      I

cn    I  o

<     I  -i

X»-l'<

m  *>i  *  n

>  >J  y  >
in  irj  m  in

JS,J>

in  ur

1

UJ

z
*-*
<*

56

5
-J

a  >-
»-  x

3   UJ

o  z.

-z  tt.

It,

z

57

UJ
(X

2

(V

•H

Ul

a

ul

c

uJ
I/)

z

3

a

r»  o
r>-  o

P-  (V

r~  v

«-  r

r--  ♦

r-  ~

r-  a

r-  *

f-  q-

f>-  <  >

O  Q.  ►-•

*  —  a

-  u.»

-•  ~x  a

*  cr  <j  t

~  <*  z  -*

iu  <  ~  :r  <«

:c  a  z  x  a.

a

1

Ul»         «)  x

""  c  c  u  <_>  ►+  a

«*  X   Z  -J  t-  Xj  *

*  V  *>   *    r>

■o  *  -o  ,o  -c

*-    33    J>  TS

ji  ji  n  «

(J  ~«

\j  n  *   n  o  H  r

■0     £1    O    O    O    £1     O

^)    ^}  £    O  £   4   ^)

X
<
Z
tt
O

u

o

_*
a
o

UJ
li
O

Oj

-i

ic

cr

>-      o

z      •

o      —

X        — .

x       m
i

<A         —
^        O

Zj
—|

♦

a

o
-I
a
o

u

o  k  a
cr  —t  x  i
or  o  jj  ~

<X    -I  Z  2

<r  3!  io  <c
•-  o  </>  q  a  2
r  t  uj  < '  z  mj
z  x,  a.  o  r  a.

Md   Q.  JCD  C

>     3|    -.

O    >i  J.

•o  •c  -o

<0   sC

**    r+     •-«    «^  f-4     ^M

z         •-

o
u.

>»
o

<

UJ

or

*
o

5

UJ

CD

X
«

ul

Sj

u4

>

I

t-  UJ  UJ

a  _j^

«l  cr  *-i  rv
t-  <  2  z

*    *   O  O
fVI  i-  U   <

o  a  "J  uj
E  u  —  a

a

<

■°i '

J>  JV>'JM   >
>C  >D  *0  >0    vO

3
Z

O

o

<3

O

Q

3

O
O

J

s

►-

d

(/>

2  »-i

Z
CD

«s
or

3.
3

o       3
i       71

3  *   •    •  ♦

3    3

a

UJ  -~
i-  i-

Z  Q.
•-•  «t
O    —

a  a  v

z  «-

QUO
Z  _J

o    »

Ut-

■J   Q

II

uj  a

UJ  z

<  a

i  o  i   z
t   j  i  -•

1*13

♦    *    «    O  *    ~5

O   «N  X)    *  -»  -•

3    3  3    :>

—•  — ■     — •      ,-«    ^-4    ^-»     *-4  — •    ^- t  rH

2  _)  *  _J

<-i  -~  aj  »-

I  2    I  O  Z

l  «    I  J~

•  O  •  <JO

*  3  •  <->->

m  -n  t  r>  o  -■

<_        2 1-

*

*    *    c

\J  M|   M   \i  \J    \

*-•  ~i  ^*  .->  ^  •-

a

CD

S
©

31        u.

*-    I  a

(<  s

*-     !    CD

Z

UJ

z
o

4

a
u.

3

M  O

uia  o  a

>  Cf   ui

o  *  >

i  It  o

ujcti  z  a

as  o

r^  is-  r^

•—4    #-•   ^*   ~-*  »-  I

3

or

u.

</)

Z    t)l-

3   <

*  r  .o

t/

»-        o        —I

I       i.

£

<
i/i

z

»/l

;,    o    0   *     «     *    »

J*  q  -«  M

*»  #■■#■*
r  r-  i>4  »^  r^

,- 1  ^*  ^4  ^^  ^m

Qi
Z        =

3

Q

r

a:

x  u

<  I

X  <

»-  a

CD  K

3  a

en  3

IT

Ui

>  I

a

z  ►-

UJ  b

\  I'

W»  Mil

*-

2  0-

>-•<!
o  -i
a  c

z;

OUJ

z
c
u

UJO.I

</>  <;

>UI

J—

a
o
<

«-  o

ouj

a.

n  >-

t-

o  II

z  —

UJ  H

i  al

uia.

c

u

v  a

*  >

o  ►

n
ii  -

K

c  a
<  <  z  a

a  —  uj  -

1-232
CO'-'  Z   ►■

3  jc  -•  a
i/>  a  z  a

i  q  i  z  i

i  <  '  i    a  i
•  (J  •  3|  •

*  *

o

a

>
c

X

3
l/>

a
u

3

:   C

J

-  c

~3

CD  Z
3  UJ
»/>  3

—  a

(Mtfl

-ICO

UJ  3
>   71

UJ

-J  2

M
*    O

z
ul

z
<
a

H

CO

i/>  i-

i  a

tr

~  z  t

x  rr  i-
303.

tfl  u.  z

a

z

3
Z

a.

x

3

cc

o

z

UJ

X  CD
<  3

a  v>

CO  X
DO

i/i  a:

Uh
I

z  w

cc

o

r

— <  M
X  *»
UJ
Q.   O

>-  o  a

2^

Z  X
3  <

3  (\J

»    «     *         3

3   *   O

z  u  z  >-

3  <

*  !

<

I/)

Z

c

CD

o

z
<*

(/I
a      z

O         UJ

t-         X

(3

(/)         <f

i-  or

UJ  u.
m

uj  o

X  X

c

z
<
a

cc

UJ

3
X

a.

x

3
CO

a

3
X  14

—  II

2

X   3

UJ  o

>—  a
t-       JO.
a

U.         •-<  0

c

3

o
o

cc
a

5"       3
z       cc

3  ♦-■
X    UJ

Q

«-    U>

r  4

i
i
i

K

a

N  ^  V

«■   UJ

•~  a

i  >

ii

UJI  I-

-i  <

Of    —Qj

1-4  2

Q    —

a
x

»-i  z

UJ  ~
I-  t-

z  a

— .  <
o  —
a  o

z

O  UJ

Z  -I
c    I

uj  a

U)  <

t.

uj  a

_t  x

3  a.

x  M

cc

a  c

O   M

O

O  UJ

ii  a.

a  >
z  i-
<t  ii

o  ~
•*  »-
_ji  a
a.  x

_j  ►-<

3   X

s  a

3<J  3  UJ    *    *    *

3    *

*    *   *

I  <_>  I  Z
I  _l  I  -•
1*13

*    U  *    3

n  n  n

r-  k.  ^

—   O    >    3

/>  />  .n  o

f-  r^  r-  r»

-«  \i  t  *  n
a  &  o  o  a
*^  i~\  *~  r-  f^

a  *»

o  >  ^  -*  \j

O  O    ^-    »•  N-

r>-  r--  r^  r-  f-

*    *  Jl  J3  -»  o    T

r--  i»-  r-  r>-  rn-  r-  r-

o

z

3

'X

!  a-
z

3
CD

;1

r  >i
a.
*g
cc

c  cp

M   O

,  °  4
•-*  UJ

a.  »-^
>  o

ii

a

4     II    3  3

<4  —  o  z
i-  a  uj
a  a  a
<r       v

4  -  c

_l  —i  —  ct
3    CC  3  X

5  a  x  o

u

3

a

s

U|1

X

c

->

►— »

3
CD

z   i  <j  e

—•   I  _J  -^

3    1  <3

3   «  O  3

-3    -•  M    1

O     O     O    D

to
a
o

*

<
t-

»-i  z

Ul

3
X

3  »    4  *  *

jl  r>  «-.  J3  >

til  so

a,

-4

a

°i

z

3
CO

UJ

3

UJ

>      o.
»-•       >

n  o

H

Qi  X
OOh
Ifl  X
(^  >-•  Q.
kfl  >   W

a      »-«o:

Ui   >OC

z  o.  a  c

•-i  ^  O  u.

I    (J>   I   2

I
I

•  *  #  o  *  ~

3   -t\nin
j>  >  >  7*  >  >

f>-  r^  r»-  h*  r^  r^

z

UJ

3
3
3

I-      !  3

z     jm

UJ       '

—>  3t  C

»-  <-ll-

O   3

DUW

3  a  +>

>i
C  I-  c

_J  tl  _»

i— t  QI  •— i
3  QJ  3

x  a  a>

I  U2

I  _J-

I  ^3

»  tjJ3  «j    <

O  *-   O    >'  3

"1'

f:

z
3

o
a

o

a

z

UJ

I

o

<1

X

u.

u.  !

O      1

to

Q.

a

O

UJ

h-

■*

o

«/>

a

»-

UJ

o

J1

h-  i

z

UJ

X

o

X

u

UJ

<

T

it

u

<3
>-
tO

z

in

UJ

1/5

<  H

3)

1/

I-  u.

z  a

UJ

z  c
o
a.
x  »-

UJ  Ul

l/l

o
«-  (/
r  f

i

i
i
t-

.  *    »   o     «    «

vj  *>  t  n   o

3    3    3     3     3

X)    CO   CO   1]     CO  ct

i  u   i  z  i

I  _l  I  -t  I

•  II  I

*   *  o  *  -j  *

o   >  3  -•  \j  o

3     3    -•-«-•     -1

si x  ns i

Z  I  u  Z
-•  I  _l  M
O   I     <*  3

t    (1    O   v  D     J

_    -d    _.  _4,_

cc  xj  xj   a  co

rv

UJ     I        u

7

1.

(?»-  c

UJO.   c

—  t.    3

3     K  I-    C

o  uja.  i

a.  i/>z  c

CJ   2

Z
3

CO   tJ

•n
\j

CO    0(.

z

or

I  o

i

lot

I  uj

]  >

o

5

l/l  '3

uj  a

x  u

C  Mi

_J  V

»-.

3  I
CD

I  &      !

•4  o    !
*  it    I

3t    -I   W

!

I

^3

1

z

3

z

a

o

UJ

Z
3

r>  *

*-   -JtX

uj  to  a

12°

^3

fa1

Ul
Q  H  Q.

«-    Uj  >-

I

o

*    «    «      4   *    *

i  *  n  d  -

n  o  o    1  i

0;      >     =>

COXCOCTCDcqcDC

i  4  z  Z  z  z

|  .— |     — I    .— 4,     _    .-.

I  31   D  O   3  3

*  f  "*  ~?   "*  ™

-•  vi  »>    *|  n  0

*  4  *  *i  -»  *

X)  CO  CO   CD    XX

~*  »-»^*  ^  —  —<

-
UJ

>

61

>-
x

z
3

x
<
z

!

i

i

^t

UJ

»4

1-

•

<

j

1

3

1

»-

-J

a

•*

1

I

<  «*

>

»  •*■

UJ

2  a

J

Ul   2   -

M

a.  uj  —

M

4    J    1

1      1    ~

X

I/)    1-  *-

•»

•  a  a

m

>  «  «j

>-

jj  —  —

X

ex  z  tt

<s

^*

L>-    ->   Z

z  w

X

—  X

«*-  a  ui

3   U

X  X

-a    a.  j

•  VI

Z

X  o

j— «  *~  «~

>  tu

UJ

O   "*

— .

►  -•z  a

UJ   X

X

-  z

«■«

_l

;♦  •-  z

a  »-

<I

u  •—•

,   X

«->

<J

j~  u

v    Z

X

X   X

1  o

M

>

c-  a.  j

~                t-    UJ

<

>   X

-*  t-

<

V

•x  •-•  •*•

>          v  a

t-  l-f

>  <t

**

^^

(<    If    If

a           r  «s

>

Z

t-l  tf>

<  Ul

—       t-

•»                UJ

w  z  z

1          ~a

UJ

o

w  z

—         X

_i    ]       3

cu>  o  o

X

a

•-4

z  c

Z  X

■  M      |  _l

<           Ji

;  2    *-i  ►-«

•-•       >■

a         a  c

>

»-

o  •-<  >-

2  °

<       —

(VI

>           ^  >

3ul  Z  Z

U.           X

—          <  c

»-

<

►-•  Z  X

!  —       >

v  tv

w          af  a

0-J   33

UJ         <

—  <

A

o

Z  3  «*

>-

1-       1

~  N

N.               1  <

I    *  *

x      z

nj          z

m

o

3  *•  z

oa.       —

x    :  —

*-»   —

f  ~               3f   Z

«»-  z  ©

X         3

>-t  &

K

-J

»«   Z  3

«»       1-

_l    '  —

«^   «

i-  -1               Ul  3

•a.  •-•  z

j          a  *

a

4~.        ►— «

-»  Z        X

1  —      t-  -

-,                t-    -

k  <                2J

«     X  UJ

UJ         X

UJ               X

<

UJ

►  ~  at  x

UJ3          _J

>        X    -

—               XI-

*                    X

t    X  _l

r      o

>■  X         -A    <

*-•

Z       -h

•-'X'O

X                 —

1         3-f-

l        H- <               Jl

'V           co

Z    NN

•-      a

Ul  o       o    <

■  z

o

-  X    MX

UJ  U-           >

•^      c/)  x

X                —  _

1   —               »-*  X

i-«    *    *

-L.

_JI-        z  z

tH

<  /»  u.

JOlU    1

—         V   _l

_l                >    -

1             »-*  u

IX    St    St

a

4         <X   U

a

Z         2

—    St

»« i-  3  —i

1-       —  ~

—                   1    31

'=    %

i  (a.   it  it.  —

I       t»

m  a       if  a

a

u      u

J  Z  •/)  W

X         1-3

»         >             i    «-   -

•tA    «    <    t-

3        *»

Mill        o   o

<A

X        J

•-<<««

W  Z  <   <4

a      x  -

r       -             XX

~                JW

lo  x  x  a.  «<

-  X

a      uj  a

O

C         -

•  X   X

Ul  1-  >

«/J       Jtf

Z              x  >

£         5

x           >4

12    M   *    <f  -

o      x  «

z

*

X    M

X   *    *

\      —a

—  •-»               Ul  u

«* .

o

x>-    «

4

3  -*^

—      =»  u

»(/>—-.!      H

_i           u

;«*  ►- »-  isc

It-

>-    <•  *-    2

rfs

Z  K

Oh    1

1-       —  >

-2i  2

'2   a  a  z<

a  ui  a  ►<

•  z

"          Q

»-  X

X  X  z

a  -*  z  «

»  X  ~   1-  1-   —    -

Jm  <  <  uju

m

<  _i  r  x

1-4

if           <3

X   <

««£:

.  _  ~  _J  _  <l  -

_J  -.  X  X   1-  H

lZ _,_

M

za~u

X

if

-a  -

.  1—  1-  —  t-  1-  1-

—  i—  _i  -J  x  a

X   X

lU    2    2      1   -

3   <   H-    Q

UJ

•-  2

Ul  C   UJ  H

a.  x  >  x  «j  a

>  X '  _J  _

1   X    _J           tf

X     t     t    t-    i

O  I-

►-  a    >

a

t—

X   1-   1

3   X   -

t   ft

i-  u  a  a

J  J- JI  _

—  _!>>    —    -

■  uj  -r »—  if
ft   >  X

►> a.  _

Q  X

(VII-  <  4  t-

>

X

(*"<    Z   V    _

—  —  z  —  <  -

•  Z  — >   3

t-     ►-   K     <J  H

-  U'  X

>■  a  r>  —  i

t-

<I

—  U   X  H

-  »—  *-

>-  3  •"•!-   «

-  >  >  <5  >   i    :

►  <  ■>   U".  Z   —   -

-  _i  •**  <s  O

ii   x  a   t  a

>   ft

a  u  _j  z    i

ii

t

ft  >  *  a

IT.    X

a  j  t  I    :

»• i-    —  (V    -

•  »-  <->   O  *+  X   Q

4   K    t    4

•  X     <    1    »-  3

c  1-

>*  it.  <  •—  a

a

y-

!-•   O    t-    <

.   2    <

•<t   <    t-    X      »

-  X   z  U  Ul  V   2

<_)  IT   (J   IT.  X   >

>   3   h-   O

:x x  -

Z    0.

z  z  ;>  x  a

X

X

(/>  X.  x  -

'  —    "*•

Z   >  X  X    z

xhi  a  x  •-

X   O    WliJn

—  3  X   Q

'3    Z    Z    2  2

a  <t  <\<

3-iija  a

o

^

X    O  <*  2

z  z  m

3   U>  _l  O   _

ui/)4  U  r  i/

<t  u  ^  -•  —   -

■  —  if)  _J  _j  r5

|

•2    O

z  u  >-

K          1   Z    2

z

<_:

z  c;              i-

1/1

1-         O  «J

O  2^  l-

*-«  _l

H     JM

»— •         |    — f  •—

1— •

_l

-•  _j              -■

-       -4  Jl

Z         i  «

j4h

-->    *

3    -X   X

z:       i  o  d

o

*

o  4

<

<  q .  <

n  <->

->  <j  ui  *    *   ;

1       z>  »    *  ->  -

»  "3

U

">  u

Ui  *     *    <

3  *    O  U

UlliJ*

'»    /l   3    -O

j>   -a  -•  \j   -n   *-

Jl  o  ■>-  o  3>

3

_

g  n  ♦  r.

-C   >-    0    J"    ^>    -

\l  -o    *   1>  o

*-    3J  >     -3  -t    \|

n  *'  n  -6  -~  v

>    3(    -1    vl    »)    ^

'fl  ji  /i  />  r

/l     O    O    D     D     £

i  x  x  xi  t;  a

o  -c  o  a  c

*-.

^>

•*w  n^  ^»  ^

^   >.    ^   *,    D    t

X  X  X  X  3

O    O  3)    >    >    7>

■>>   >■>    I>   7

(3)    X    X    X  3

X  X   X   X   3

X

X

X   X  X  3

X  X   X  X  X   3

X   XX    X   X   3

X  X  X  X  X  3

X  »  O  »  (T  o

z

UJ

<
a

63

3

o

a:

</J

z

UJ

•X
4

x

4

Z
UJ

■x

a.

d

Ui

X
©

4
x

UJ        U.

*  «

D   >

z

UJ
(X

<
a.

4

2
3  *

2
u.

<
1

a

«-
r

?."!

•-

f\J

r-4

rl

z

*

•

I

UJ

a

— .

-—

X

z

»-

»-

l»-

ID

4

X

X

o

4

4

4

X

X

— -

w

*~

**

u

»~

©

a

X

X

z

4

4

UJ

4

Ul

X

X

a.

—

_l

X

-I

>-

Z

r

*-»

—

►-

»-

X

X

•

5

a

4

<x

z

a

4

X

X

o

X

—

X

-I

•-<

x

z

1-4

»— •

K

««

W

u

q

z

z

z

Ul

o

o

O

>

•—I

»-*

x

o

z
3

z
3

4

UJ
X

X

X

X

4

4

3

ui

X

X

X'

X

X

-I

u.

4

ISJ

fsi

t—

W

-v

rfi

l/>

z

I/)

*

♦

•XL

UJ

UJ

UJ

</l

(/)

(/>

X  10

4

t-

4

UJ

©

UJ

2:

X

X

I

4

X

tA

4

w

\-

x  *-

*^

—

—

z

u.

z

»-

(9

r-

UJ

UJ

X

z

c

a

a

X

4

Uj

X

4

«-

<a

~-

_J

~~

a.

r

a

Z
*

1
1—

z

z

UJ

i

~-

X

~

UJ

>

i

►-

<3

r*

a

o

i

X

1

c

<

X

i—

<;

f~-

x

X

Ul

X

■■*

o

~—

i/i

X

4

z

X

2

4
*  u

4     4    *   4    *

*  *  i

VJIJ-nO^OJtSHMIt/i
"_-   "3  --    -—--«-    4  D    X)    »   JJ   D    i

O   JJ  f  J    0>  ^   O  0*  ^  O   0'

Ul

a

X

X

o

o

X

a

>

t-

o

t/>

>         X

X

z

3   X

J—

4

J       »-  c

1 — ■

X

_J

©

X

1-

31  UJ

z

O

O  2)

Ul

1—

— .         ui

J

-I

-1         IB

UIUJ

1

X

UJ

>

i      tr  5c

_!

o

X        c

— *

4

*J

-J

X         <-

X

I         X  t-

X

Ul

o

UJ

t         «/>

a

X

X         (/)

O

3  —

o

-

»-         UJ

K

X   *

X

X  w                      !  —        Q

l               o

t-

«s

o      o

k-

•*

•w

»

z      z

Sf

«•     •■

(5

Ul  -4.

UJZ

1       o  x

z

U  i-l

_IUI  X

ug

O

i   UJ

4  V)

1  *.  o

X

y        •  »-

_J

XX         UJ

_l  4

X

»

V)    ♦          >

ui»-  a

-J

3  X

^*

:>       —i

X         Ul

_1

i-'3

-1

2    H<               ■—  1            •<

ldliD>-

!-•

X

UJ

U)  X   4  4          •

CCD  X  X

X

Z  a.

X

a.  Jnj>       x

X         3  O

O

X

O  Q  X  »-i         X

»-  o  z  u

(A

UJ

c

»-  4a      c

•  K

w

X  :*

X

tfi    t        UJ        X

_J        WW

u

t         l/>  4

1-

v  ■>  -z.  c      a

Ul  4   W  W

X

3   1-

^— '

>-)  Ul                X

X  Ui

-J

X

o

X  X  X  X         >

at  ex ■

a

t                -

z

3  Ul  O  O         O

O  4  1-  1-

—    *-

^r^

UJ  c

X         XX

u

*•     r-l

>  I-  u.  X       Q

1-  X   X  XI

—      1

t.  •-

o  -♦  c •  o      z

*  a  —  —  -•  —

<

1

•-"  Q

X  t      »-      u

1-         ©  Z    1      ■

1

tr,  J.

1  C   -t  X    t         /

X  X   Z    t.  »h   1-

t—

1        1

x  c

•  1/1  o  >        >

x  3  ui  — to  a

X

1        1

•  h

>t   X  l~  H>  X  »-

•  I-  _l  "J    x

UJ

J          1    X

X    t

••    •  t  x  i/>  a

x    *    *  to  *    t

to

t/)  c

o  t-

i/;  i/i  m  ijj  4  il

O  >-    "  X  >-    t-

z

3  X

x  a

z  z:  l/>  3  X  3

xx  o  — '  to  a.

■— •

i          m  u.

u.  x

-•  •—  X   t-^  XI  i-

x  a  ►-•

X  X   X   Z  X    X

a

t/i

to  ui  ui  a

uiu       x

t-

o  c

>  z  o  </)  i/5  •-.  a

Ol/l   U  3

>-t

_1  J

-*_i  /iJiri

_i  y>  _j  3

J2

4    4

D   4    1'iZ.l
>  ~>U  4    4   —  ~.

44    4   -J

3    4

1

4    4     4

*    CJ   t

O  4    O  •-»

■S    I)

7

3H     \J

*»  -t    /I

O    •-  »   »   3  H

VI  ni  ■*•  ni  15   «-

D  C

a

>    J«    >

3*   >    3>

.  >    >     J»   J*  •  O    3

J)    31    3-3i3    3

o  q(  o  a  o  a
(VJ  rvj  c\j  (Ml  rvj  (\

o  0*

a

CT>  0»   (3*

(*  <J>  O

ff»    O*    ^  O*    O   C

UJ

o

o

o
o

at
±j

<

a
c

c

a

A

o

z
*

o

2

I
«.
z
«s
a:
a

U-

o

UJ

in

UJ

a

*
rt

c
o

z

z

G   «

f  ->  —
J  —  -+■
o  o  o
(VI  rvi  (VI

o

rvi  i

c

UJ

CO

I
</)

g>
a

«-
o  a

-B

a.  to
o  G

t-  a.

z  u

0  CD

1  o
I/)  ►-

G

a  uj

£3
•  =>

i  i

i  i

i  i

t/>  _j

z  -■*

•-<  >

*  *  t.

vi*   »  n  ©  •-  o  >

o  oo  o  o  a
r\j    w  ru  ivi  rvi  rvi

;>
C

c
z

h.
I

>

_i  a  uj  -
*  •  ►  t  in  a.
t-  i/)  i/;  h  au
a.  z  z  a,  oc  a

X  -•  «  _J

a  •->

</>  U,'  CO

~t  _i  (/>  r  x

G   4    4  <T  G

•a  -•  vj  *i  *  n

.M  VI    VJ    V)  VI  \

o  o  o  o  o

rvi  rvi  rvi  rvi  rvi  rv

Ul  Cl-   (J

a    -  a  uj

I-    M

">   <J     «    «i

O  —  fl  T>

vi  vJ  vi  v

o  a  o  c£

rvi  (vi  (vi  <v

-•  vii  ■*>  *|  r>  c

o  <3  o  O   O  Ci
rvi  rvi  aj  (Vf  (VI  rv

it

a
u.

o

!   G   *

M  «!
i  />  XI
i     o

;  (vi  (\i

CO

G
CO

UJ  I

UJ  IT  I

n   n  n

*     fV  I)

X>

O  (

rvi  rvi  rvj  o

rvi

I

«  '

X

o
►-
(/>

65

o

a.

UJ

■>
o

i

Q

a

UJ

>•
o

a
o
a
x

UJ

UJ

t/r  a

«>  T
It  D
Ul  ->   *    «

=  •   M  1  ♦    />

CO  6  AdO

CD   O  O  O    O

»\j  oj  oj  rvi  cvi

a:
o

o

>

X

UJ

c

>
u.
•— »

Q
O

2

a

u

>

O

z

I/)

3
Z
X

o

a.

o

X

z

UJ

X

a
o  <

UJ  a
>  u.
c

X  u.
o

hi

Oj  Z         <
O         UJ

oi-«      a

I-  J-  o  <

4  UJ
»-  U  Z   Z
Z  O  UJ  UJ

uj  -J  a  a

SCO

J}  *

4  uJ  .J<X
X  Z  Q    O

U.  t-

o  o
»-  i-  t-  o

(/>  «-

ujt-i/i  r

oi/iat-

_»  Ul  O  UJ

a  3  c/>

a

r  •-  r  >-

UJ  4

z.  —

a  *-

o  •-<

r

S)

x

i
i
i

y)

tv

X

Jl
aJ

3»
O

X

UJ

>
c

UJ

l/>  U.  X

z

•—I

o

ooooopooo

T

UJ

N

•-  Wl

4

>

H      -J

UJ

CD

IN

««  O  >    =1

d  <C  O   M

oo  o  o

oj  rvj  oj  oj  oj  (VI  CM  (M  (VI

3    -•

X  x>

o  o  o
(vi  !M  (vi  m  cm

*  /»  a  •»

OXH  I

3  o  o  o

oj  oj  oj  oj  rvj

>    3  —  M  1  .*

»  >  J>  J>  >  J

o  a>  O  Q  O

oj  oj  rvj  rvj  oj

o.

Jl  JJ  —    H  >

>■>■>    3*  J>   C

o  o  o  o  o

oj  rvj  rvj  rvj  ru  a

-«  vl  o

3    3    3

*    4

o  m  t>  Ti  3  -h

3-3    3    3)  •-*■-*

cm  rvj  rvj  rvj  rvj  a

a.
3
ui

>•
o

x

a
o

ac

UI

(A

or
UI

o
o

Ui

>
o
x
o

Ui

<r

3

m

>-

tr

k

M

u

m

4

-1

X

t/>

p

o

Z

<

H-l  X

•—  ui

X  »

o

UJ  •

>  X

O  UJ

X  »-

•-•  _l

UJ         «/)  u
M  I   2

•-.  ♦     Q

fV  _l  JC

a.  <       ui  u

=;  •-  ui  j
ui  t-  in  a  1
:>  ~  <  c
o  z  tr  u.  c

x  •-  ui  i-  2.

do

*-      -i  a  <J

--•      r  t

z       z  o  «:

31    m    Tti         «    *    *

rJ       (V|(vi(Virvjfy(vi(vi(vi(vi

Ui

(A

a

Ui

Ui

<_>

X
Ui

a

«-

3

X
«-
Ui

«-

LlJ

t-  a  — •

i/i  a  o

>-

(A

<-
*

ui  o      to  m

a  ^  a>  «-  ~

UJ  Ui  u
Ui  (/}  h-

£3>         «*►-•<

at-iadx'O

4  Ui  2

•   X

n  *  n  o  o~  t>

VJ   M    M    \J    M   \i

rvi  (vi  cu  (VI  oj  (vi

5

UJ

(/)  a

3  X

<  3

a  ->  *

J
(\j  (Vi(v|  (VI

-«  "M  *>

"0  -m  »>i

c

o  u,

<3         UJ

ui   U!  —

a.  h  a
ui  t/|  a

>  >J  UJ  '
O  l/}

UJ

«-
t-
f

UI

>

4-

or  -»  »  *

c

uj  u
i-       ui       l/l

•-•  O  4    •"-!

Z'        O.i-HCC'X
3  *     3    <  UJ

I
(VI  (\|  (VI  (M  (vt  (V

a.
o

QC

a

UJ

67

a

c<
■    ►•

'     3
¥  -

>l

u  '  •  •
>  a-

0  ■•-

1  :«-

:Cl

*  at

*  i-

—9

-  —Z

-  Cj
i/i    *J

:1I

>-

AL
IE  f-

:~-  a

i     M
i'  t-*

a:  a

ic  x

»-
a

z

e  a

z  o

Ui  O

-I  -J

r  «•
»-

a  I

*  z

>-  <t

a  x

L

3  U
>  J

J  <<

*  U

7'  n   *  *-

«>  £    O  o

*
a
o

*   «    *   V)

X>  >    3  -•

J  M  rvi  nj  rvi  rvj    (Vjcvi

(XI

at

in

z

3

U

k.
Ul

tr

z
o

u

o

z

3

If

o

o

a  so

in  r-t-

c*  ft

"1

tn  p

o  c*l
f>-  if
eo

o»

©

(M

o»

o»

a

CM

O*  «M
OD  ♦
lf(  ^

cm

ir*

o»

«VJ

x>
cm

,£>  <y  K
(VI   -»   -D

xi  x>  cm

2.     :

2     i

2

■♦•  -•«>  ©
Mr  \j«-  -n;

CD    oifi  -->
-4  CM        (M

0*  X    t>
X:  »•  ~«

h  09  h  a>

*-l    ^  ^H    ^*    H

-•  ©o»  in  o  at
©  (\i  <m  in  xi  a

r-  r--    r^    .-i   y    O

UJ

Z
J  J  jn  u
3  3  3  in  a

X   T  I    Z    «*

o  a  a

o  a  c

O  O  O  O-

<  <t  <  o

►-   c  c:

«/i  z

UJ  UJ

~  cv  >-  cr

o  o  o  ./>  <  •-'

3  o  o  z  n

«  1  1   1  1

©

— «  CO  *

x»

4

CM

Xt  o  X)

o

tfl

O*

<■»>  X>  <M

(•>

■4

C*>    »H    if.

-*CM

m  <\i<n      en

<n'(Mtf>        fH'lM

»  cmo

z
nj

«

a  uj  uj

c  i-  »-

i  »-^  • — •

til  K  t-

C  (V  _l

JUIU

3   <  <I
CD  O  O

X  X

UJ  Uj

O  C

z

o  e
c  c

o  c

X)  h-  3D    XI  ©  (•)

©  c» »  ot  ©

C>  CM  <\l    OM

in  ro  cm  moo

i\IX>  ©   X»0»    -"

co  cmo  -*«m  n

xxx

UJ  UJ  UJ

O   O   O      rt(\i

Z  Z  Z    UJ  UJ  u.

t— •»—.»— I   ->  :>   >

rv

c  o  t  a  a  a
a  m  rr  uj  uj  u.
coo  Q  o  c

>  ~  cm  uj  <  <

OQO>Z2

3  3   3   Hi  3   =

jj  uluU  jJUJ  >xl

o  a  o  o  o  c

o  -•
•♦  X)
(Ml  CM

CO  CM

o  D>  o    rj*  XI  -Of
©  Of  0>    0>  ©  cv

<\j  _  _    J.OK

«-•  0*1  3D  I  OUT)  x>
00  AJ  *  XI  30  *
AJ  M  (M    CMC-  -

*-t  *^  —•   — <•-*  fM

a

x  a

O0       -I      — .       (MO     t~

UJ  UJ  UJ   UQ.  1/

>  ->  >    ■>

a  as  a  a  >  3

UJ  UJ  UJ   UJ  •—•  c

o  a  C/  3  c  2

»*  CVJ    ■*

>•  >■  >
a  4  <x  <

«  -I  -I  -J        u.

>  a  a  a  ?
rtifl  m  ^i  >  o
t^j  ,>-«  ►— •  -h  •— •  x
O  Q  o  doc

x»r^  (M

•»

O  X)  XI

00

o>  if"  o

IT

r-l           (VI

r^

M     i

X*
<M

ID
(V

Ifl
f>

eo

in
ifi

or

in  x)  ifi

»<M  X»

(n

tn

<n  x>  jo  o  -»
r-  n  «V'  »^  x

ac  (M  CM  0"  Ifi

-4  -iaa  --•  a  *  r«-

r-f»>(Viirx  r.'Kcvfvi'-o

<r  (M  (\i  a  in  r»  ■  o  -»  s»  cm  (M

r^  if«  -f  -+e

X)  r>-  ■*■  -o>  •*

O  0>   IT

(M  r-

*~

-♦  if*  if>

x>

m

in  in  ifi

kft

(M

to  co  co

9

xi  co  o  i~-  r>-  rv
-•  ao  o  o  r»>«fl

-*  -*  in  cyao  a

— i  ^»  _    I

(M

©

M

co

o

CO
CM

Iff*

!=0         O

I  ^H  CO

!e>'      ifi

ifl

*>-  iniji|K4tMio»-*T>
x*oi  •♦uji  -*h  «»-c-*h-r^

x>X)Xicof^oo(vjrn^x  oc
^  ^.  _,        (\j  <vi

oxi  m.a  ^  r>

X)  ■  x>  r*-  oO  rj>  (7>
co  ao  co  co  co  ac

x  x  <

-  -1-
U.  Ui  u.

z  z  z

u.  u

Z    2

(M   rt  -4  Ijft   X>
XXX    Xj   X    X

I-*    M   •-♦    l-t    HH     •-

u.  uj  u_  u.  u.  u.

Z  Z)  z  z!  Z   2

30  I>"   ifl  O   •♦  (*)

-•co  i\in»o

(M  0>    If)  O  —»  —

-■—•—•  (\i  rvj  rvi

►-         «/l         — •  (V

»-i  H    0-      '  0.   Q

tr  tt  uju  3  -

HUCKili
X  !/>  O  Q>

r  z  3  qo  3

D  O   f>fM  (M

4>i      in-*  co

►»     ,  roa  in

—      |    -<(Mj

U.   »-

UJ  K  UJ  UJ

cm  x  o  z

3   V-   UJO  O.  1'
Z.  2t    Z  Z;  Z    C

<\J

co

CD

•c

69

(VJ

CD

(VI

<C

o

rvj

CC

(VI

in

IT
IT

CO

3

CD

CD

>

co

•VI

co  (Vi  r-

>    O  O

«  OD  --•

-«  (VI

#

o

in  >D

^  n

in

s»

-#

CD

f-  rvi

(VI   -O

(VI

<o

in

Cl-

— *

<c  r-

c

(VI

r-

K

CO

in

■o

P   J    H   O    ^H

o  n  s>  o-iic|«HOHH!v

*  <  <  IT  C)  II   h  (V  If  ■  IT  ^

r»»ff   -•  o*  o  o  o  o  •£

3  O  Cf   D  i  CMT
0   o  r-  ^«  (»>  0>

>  »  o>  o  —•  a

■*  &•  -O  i  (VI  CO

(vi  -*•  in  *•-<

o  o  o  o  o»

(VI  »^

=  <0  >   J>  •#  .fl
in  (V  0>   C  ©

o>  o>  in  ht-  ec

i
_  r~-  (T-  — ■     j  in       •♦

in  k  coo       r»-      r>-
(vi    !

=P

co  o

r»)(Vi

CO  CD   CC

tfll/)l/>l/>   l/)(/t(/>l/)l/)(/)l/)
Z  Z  2  2    2  Z  2  2  2   2  2

o.  a  a  a  a  a
o  o  c  o  o

o.  a  a  a  a
c  o  o  o  c

V)

2

2  I-
C  3
.-.  1
»-  »-

a  3

o  c

■O  P)  SI  (*)  CO    JD

c  ec  <c  a  a
x  o  >©  cc  (vi

.-t  (Vi

r-  o  m      m

MOO  M

co  m  o*      •-•

(VI   (V

H4)S   ON

«C  •*  (VI  C   ■*   s-

*C    O  .£    sD    -£)    s£

(VI

(vi  (V)  r-  --•  o  cc

,0  «».  r».  f»-  J)  s

tr  f.  o  «»•  Ci

>c  ■£  -o  ~c  r-  h-

(vi  s\  x> ■■* "-  sliflnj'O-om'ir

NON>0    >0    ^O^^

co

4.

in

(vi  r-

f-co
♦  in

-»    O  (VI

on  o
in  o»  ■»

CO  f-   CO
o  ?*>  **

in  r>  -*

,r»  in  co
■*  o-o

I*  (VI  -♦'

M  r>0    >    O  O

IT     W    r-t    r—   (\.

LTi  O   ■£   ■£>  <£  r

•*■   h«  D    -»•-*•»

^-  or-  ^-i  ^i

c  s»m  (vi  (vi  (vj

«(vn       —.^-(vo^o  loo

.iiiuudo  ccc»-  +  a

-»-►-»—  H-H»—  »-»—  r-  >-  >— •—

Uiiiaiiia.ua.  xo

oooocqooocc  00

rom  or«-(>(v|  -«f~r»rriLn(vi
>  -•  <»j  f»i  (vi  a

IflO'ON  <D  CD

(V»  (VJ  ^rt  ^  (VJ  ~3

u.  u,  u.  u.  u.  u.  u.

(/></!>  tn  ir  it  a
o  o  o  c  o
a  1  a  a  a  c

©  —  (VI  (T.

X    X    X     X     X

.  U-  u.  u  u.

H(Vl-ttl-ht-

!\>  (v«  (/)  -n  in  in  ./)  *

'jJ'O  a  o  o  o  o

a  a  a  a  1  a

-*  (»jr>-  -^co  in
<vi  .*in  (V  ^o  rvi
ao  mm  nn  -*

f«- ■  («- .   ^.  H  a
*  Z  Z    2  Z  Z

o  o  o  o  o  c

-<  h  co  ^  n  ^
a  m  a  in  m  if
(*)      o  m  m  m

i/|i/>  i/)i/)i/i

>     LU   U_     till  UJ    ll.

c  a  e  o  c  c

<  <     r*(V)  -^
H-  I-     I-  I-    »-

•-  1/5  «/)  l/J  (/",  1/
X    LU  UJ   uJ  'jJ  aJ

a  0  c  o  e  c

(vi  0*  co  (vii  rvi  o

M  (^'-*Ol  -^  (VI

m  o  m  1*1  ■♦  ■*■

(V

O   —t  (*)  -H   •-•  (V

^-  2    2  2!  2  2

►——*»— 1  t-^  h^  »-i

in  O  Q  Q  Q  Q

b_   cr  <  <i  <i  <

13  uj  uj  uj  uj  u.

c  or  or  ct  a  a

(VI

S.-Z.  -z.

o  <-<•-•  nj  (vi  c

0  a  c  q  q  c

Z  <   <   <   <  <*

<l  JJ  Jj.'jj

1  u  a  at  or  a

O

iJi

in

in

CO

x>

in

in

•0

z

3

u

5,

cr

z
o

o

z
3

o
o

~>;         1  X»
X»»         Nf-

oto  IT
ru-  -*  -»•
<vir»)  oo

o
oo

in      in  ^
x»      am

-*■  -o  x>  *-  x>
irmr-in
m  ml  ^  vo  -c  -c

■o  x>  x  x<  x*  oo  r»)  k  i

r-4     ^*  ^1    »*  -+   *-<   *^   ^*

00  J3«(\l  — t   CJ»  •*

•#  in  x>  x»  x>  x

z  z  z  z  z

•— »    .— «   »—   •— I  *—  I

Q  O  «*   «t  <  <

•a  <r  i-  *-  ♦-

UJ  UJ  UJ  UJ  iu

a  cr  a:  cr  a  a

—  c.
u  o  <-

«M         Z  Z  Z  2

fl   ■*  i-i  ft  —

3  O  «a  <  «»  «

<t  <»  t-  »-  t-  «-

jJ   UJ   aJ  UJ     UJ  xl

a  cr  a  cr   a  a

c-  n.  n  &  o  o
in  uf-  ir.  ©  in  ir

DIMIDN    -

xi  >q  o>  fN  r^  if

z

UJ  _i  _i

a  3D

Z  2  «a  Z  I   2

<*  <j  a

»-  »-  c  o  c  <

UjU  jQOli

cr  a  cr  <  ■«  a

z  z  z

hmU

<  <j  cr       -

HI-    <I    Dt

cr  a  tr  i/>  i/>  i/

— « 00  00    -*■    £1

*  —  ~*  •*  ■*

<»>  r>  jr«  x>  s>
c*j  oo  ao  oo  a

2

•-«  >-  >-  v  v

a  t  x  at  a

<i  <1    <t   <z  <

UJ  2   Z   Z  2

CH  3  3  3  Z

I

O  «*<  (VI  ft/

3  >-  >-   V   >

Z  cr  cr  cr  a

3  -a  <  <  <

C3Z2

►-  3

<r,  «t  •*

3>3   =

i  I

I  I

I  !

•*  (v:  x  m  in  — <ff  mo  -<
^(vir-winc^-^eimoo
(vi^r^opccoDCO'Ooo*

cc  *>  (vi  *

x  ■*  —>  (VI

CO  f-    (3D  o»

x  r»

O  f)

•*  in

IT

in  4  «  m  o

»-•

(V

«\i

•*

o>  m  -♦  o>  ■*

o

a

c*

c

-•  o  m  m  ■#

e

«c

—•

»■«  ••*  ^H    «-4  ^*

(VI

•-•

71

IT  -i*  <»  it  orcein  o
x>  ^»  »-  « in  n  o  (vi  m  cr
cvir^r^ccececocro^cr

XI    --    O  LP

o  r»

ITC)  i-ftl

o  n

cor-  iff

tn  in

(vi  ■*  -*tn  n  —
•#  o>  m  •-•  o>  e>
o      «-•  o  m  m  ■♦

CD

o

x

(»)

x>

(vio>x(vimcc-*Kino
in  ©  r-  -hui  ^oMinx
(vi^r^ooxxxooo

mn  m«

O  (VI

x

CD

-«x  (V)  m  (vi

•♦

m  (vi  o  (vi

oo>

-*

n

o  -*  ^  cc  cc

0>

co  r-  r-  o

C)  ■♦

O

o

nfim  mn

x

**  ^«  »-•

1—   r-»

(VI

p*

f— •     ^-t    -^      r-!     ~^

*-*

*<cj>(vioinor--x-*o
.n  :>  x>-».n  o  f>-  m  />  x

(\jr^t~-cocococDcra«o

r»)st-coo-*-r>-f»-in-*r^
rn^-»^Hininf^(vij>^

(V.'X^XXXXO>   (J|  0>

•&  a  rt  o

^  (VI

(H

/I  -«   >AI

il  »•

♦

JIKh-ff

(Vi  -*

o

^-t       ~-«     r-«

l-i  »*

(VI

-«  cm  x>  in  -*  r-  x>

o-*-*-iofiinr^-(viinlx

rvi(»>fMr-xxxo>ONO

rvi  in

7

o>  ■»  ir  r-*  (Vj  <■

(H  i

c      f»o(n(r>o*m(v<<
«      -*o>(»io>-*tnr>-(vj

«D       ccrvjNNoctccDff

X)
x

a*

a.  a

CO

co  0>
co

»   -«  Xf>-  S>  r»Jj

o

^  «-»*>-  .\J

n

*      co  (Vir-

r^

(VI  X  <#  O  9   X

■»  XI  Is-  ~+  r*l  X)

x  x  x  o>  on  cr

»-«  (*•  o*  p»-  cf
X  X  X  X  X
0>         X

c  u,  e  u

Z  Z  3  Z

—  »-*  k—  l-t

u-  j  a  j

ir.
j>  o  •-
n  oa
«  <*  <*

z      z  z  z

UJ  Ul  UJ

DU'OCCIi
Z    O  3    3    2
t—  >->
V>  T

i/:j  mi/

N(f   ON

4OC0M

oc  k  ko»

**  (•)  X  o
0>  CO  I—  (VI

«0    KO

—*  —t   O  o  J

a  a  i*-  (v  *•

«C   K?    0'

e>  in  o>  n  <#
o>  r^  m  (vi

*  CO    M    C>

ro  o

(V  4

(VI

o

(VI

o  m  xt  — <  ^»

X

"XT'

?io  x  ^

X

■>   >

•-•  w  •*  m  o

(VI

X

-^     ^-4    •—*     ^^    ^-«

-•

p-4

«  in  r«-  o  ■♦  ■*  (vi

IM  *  I  OKSf^

c  xi  ^<  m  -*  in  (vi

#-*   •— «    - — t    . — t   ^-«    ^4   ^^

en

-*  X

■*

^  o>

(VI

X

(T

-*  o

IT

-i

-*r«-

■#

(VI  ^

X

o

•-i  •-«

^-

r

xt  -*  o  rvi
•♦  r^  X»  -*

X)  X  r»-  X

o  o  X>  4

o  (v  cc  c

cm  <*       a

0>   O  If.  Ml
O  (\l  cc  o

^  •♦      (#

r-  ji  go  x>  in

»>-  !•)  *  o  s  J

in  •-•*       ai

a

X)

x*  nn  •♦  '•"•(*)

0(V

•*

o  x

•*

-■ocons

f^  •*■

in

0>  0*

X

on  ~« (*> >•&

in  rvi

in

X

r>-

^4  •— 1  r-*   *-i  t-^    •-*  ^^

""•

3>  On^l  "M

x  x  «-"  a  — ■
o  (vi  r-<  (v  -a

>»  o>  in  m  if'  -a
iC  (v  <f  x  o  a

«OmIM

m  ■*

X  —

in  •-•

(V  x

x>  o
if  -*

(VI

in

in

in  i
cc

o  in  »«MioKN  ^xxixin  »<n  ok
0»  oo  ■  J>  X  '^-  «p  X)  'co  »  X  ■»  .nl  ■  •*  «  (VI  ("-  X  r\J
m  (Vi

x  x  -y  si  *  x  (vi  i
lOHrtrt  m  o  xi

^H    *-*    r~*    ^4     r-f    r-<     r-«

C)  C>(M

(V

a  a      ~

cd       r^

o»  o>  r^
(vi  co  a
(VI        x

x>

X)

x  ot  c-  rv  r--  (vi

^    XI    X    NO-J

X         XX

O  UJ  UJ

3   Z    2

IT,  2.   3.

I-         I-  t~         *~  *-     \
Z         Z  Z         Z  Z     I
UJ     I    UJ  UJ         UJ  UJ  (A  1/5
C.|jL/CCUJCClilU.UJ        Iv

K>-ll-hHKh4    4H  w

ir4  «i  if,  i  ki  v.  o  c  i      a

z  z

<r  <j
►-  t-  z
o  o  -•
.r  q:  to

<I  <    <3.

2

3.
3

o  u  u

T

■X>  i—  >-  O        I

oa.airt'
ouuuc

(M  >

Li)    UJ   —

z

5

o

z

Q

z

Lii  Uj

2J        Z

UJ  UJ

Uj  C  UJ   Ci  UJ

z  a  z  3  z

t— •  ^-  l-l    H-  *-»

S    l/II    U)S

or  ,  o

C  _l

or  a  a.      or

X'Xl   ■<       JO
LJluUJU.li.

(VI

o>  m

(VI

CO

n

CO  (VI

o>

-*

co

o>  o>

CD

o

(VI

•-«

(VI

«-•

-*X>  (V

CO

O  IT,  O

•#

n«c  o

o

«-«  .-«  (VI

(VI

O  CD  IT  0"

it  ec

«|

X  (Vl  ■*   X

O   (VI

e

KCOS

ec  o»

(W

■♦

r-  ri

i  i>

<*

e>

•»

CD  (VI

cc

IT

m

o

»  »

CD

0»

(VI

(VI

—

*4

•-1

-•in  ■*
in  it  a

wx  o-

ok  o  o>  ■*  •—

o*  (vi  ■#  ec  o>  o

^cogcir

(VI

x>
x>

IT

■#  (VI

in

in  r>-

1    O

cc

«1

(•>.*•*

r^

in

X  (VI

r-

■#■  *-

1    •"•

-♦

»

mm  o>

•»

o*

O  0»

CD

CD  CC

(VI

o

3

(vice  o>

o

•-•

-•

rH      —

1  t-l

(VI

r-*    «^     f-t

(VI

irn^h  i  -#-  o
x  o>  (\j  m  x  o>  o
i-  X)  oe>  co  co  o>

c
(VJ

o

CD

to

in
u
o

z

in

UJ

z

o

T

CJ*

m  o
x>  in
o  o

ir  xa

o  »d

■+  cj  r>  -4
e>  in  rvio*
od  e»oi

1

n  o  om
x>o»   x>  m

2  2

lulu    U.UJ
?     S    Z  Z

C

3  Z>

Z    Z     2   Z     1/  0-

Z  «  U-

UJ  N

_i  zj  .-

z  x  i/-  z

r  a  n  o

3  o  o  a

in

in    j  in

~    I  — .

rorux
■  >

in     in

m

r-4

in

H  in

-»  ■*

.Tin  x

in

!    "

in  (vi  x)

CO

(vi  ■*■  x>

r^

(VI    *

<»

m  o  •*

x>  -*

inouiKK  n»

n

■-   M    M

0

\j    ^    p^

3

n  •*

>

>,_•■>

in  ■♦-

r>  -*■»>»  >■>

o

B  0*   O*

co

O  CD   CO

(VI

0>  a

X

— <  CO  o

~>  O

r^  x>  o  o  co  co  co

(VI

-.   -

(VI   —•  ~+

~H

•^    (M

••

»*  »-•  »h

(VI  (VJ

•-■  ~*  — *  — *  ~-«

cc
m
ec

-»  -*  xi  r-
r-  00  (VI  x>
cc  cc  on  ec

-o      co  (vi  in

in      x»r-  (vj
co      x  x>  o>

4-  oi  ru  ao!  n
^  cc  CC'  (vl  o>

cw     i  **    I

(M

o  o>

s

cc  d  k

XfU  (V!

c*  m  o

•4    <C  (V!  (V

(V)  ~*  If    •-<

i>~  c.n

in  co  o1  o> I'O
a»  >  >o  (vi

Ho

in      in       z

U_ ■  Zi   U    3    LU    C

u     UJ
>-l    1-4    I

arcczul^zzs

•

<j  <  i/i  it

a  a  a  a  g

-i-i  _i  _i  _i

iJ  UJ  jj  LlJUJ

i  I  i  x  i  H

~3

O  0  T

U.  Z  <I  o

z  iaj  a  u.

-•  -1  J  z

—•Kin  ma*  n  •*

(virgin  omr--»  •©

c   -i  n  (V)  — »  x   a  ^j

(Vt^^«-«  ^  (V)  •— <  <-*  •"!

o>  ro  ^*|  ■♦  — t  <o

1>  T   X    -0   ♦  CO

•-•  K  O*    •-•  O  CD

rt  — .  «    (VI  (VI  -H

(vi  ^

CO    •♦

x>      o»  it  o>

co '     o*  ao  o

ru  rn  x>
-C  vC  o-
K  do*

o  o
ec  o

I*-  0>

>0  >  o  -i
M»a  it      a

^  O  l»-    hi  C

<c  o

«J-H

-•■r^i

2^

X>

x>  j-   -.  f>-

t^  a

UJ  UJ  UJ

?  zi  r  ?

z  z:

z  a

►-.  «
j»i-  i  i

x  x>  *  -<

K  -*  0>  ■  •#

(VI

o>  in  o  o

CC  <£    0>

o  r>  <x>

a

-^  ■*  xi

(VI

-(viq

5s

co  n

J

l-
i  z

UJ
UJ  Uv  Q

r  3  -z
—  —**—+-

Uic.J

UI  ox

t      a  uu  r

UJ         >   I  o

»-            u_,»  U

•-<                l-H     «-"  ~>    _l

•«■         X>  XI  Is-
CO  (8N«0

Ift         .^K  0>

T

X)  J>
<  -*  if

in      eo  k  o>

*T  I  •-•  »H

cd  0>  IT

•Jt  o  ir
O  k  o>

oo  m  co  (vi  (vi  k

f>-  ar>-  —t  «-t  in
^(H  co

1

2      1  2
Ul  UJ

C  Uji  C  Uj

I—   •-!   I—   >-l
1/1  Z   IT  X

>-

i

UJ

UJ

*  M

_l  _

X)  o  in  xi  (vi  co
co!-*  •#cotod:>,o»
o>ir^  x>kooox-x

I
<*o       ooaoionrvtoin
in-*       rnxx«xi-*coi*)'ao

(VI  (VI

T

(MX  X

OM-0>  -♦

x>  x  o

-f  'i  (VI

*-  o
o  »*
o  o
(vi  (vi

I

o>  oi
ec  n

(VI  (VI

O    O    41

nair
X  r-  a

•■*  X  Ci  ^  ^>

r

*  X  X   ~*  ^  CI

o  o  o  ^3  r-
(VI  (VI  (VI

u.  u  u
Z.    Z  Z

S  z  z

=  7*  rvi

X    (V  cc

x  X^  i>

rvi  rux

X  x  r-

n  x
o  x

O  if/

r>  (Vi

o  x

r-  x
o-  •*
X  o

'i  -•  (vi

o  x  ad  x  o

— •   -+    — i  -H  f\J

in  -*  "^  x  ■*  o

X)  ^0  r-  o  rvi  x

i

Z  2

ox

u

or

<  O.  X

J

I

-*

0>  -*

rvi  1 1

•o  rvi  x

O  -Mi  r«"
x  o

Clulu

-  z  z

KM   II

IT  Z   Z

1-

z  £  z

ip  it  o  o  c>  ©  o  x  n-  •*  &■  it  ^~  •—  it  in  c*  r-

*  .»  4-  -tifiiriririiroo-Htvjf^or-.

^_,    _   _   _    ^H    _    fvj    (M

IP  O  «  IT  MM(1  CI  (V  K

OvtvDccor>— .<»— .cr

444«ITITC   ON  o

^«  ^  (VI

f)  \C(Ji  o  CO

Hitni  oo

4>  OC  CD  *-*  <x>
~*  (VI

If)  (VI  •-«  «-t  h-  o  41

in  ^w  n  o  in*

«£    ^    f*-    CD  O    ^*  ^-«

^  ^  (VI  (VI  (VI

73

•*  o  o

CT-    o  —
^  (VI  (VI

-*  •*  cy  ofno>(viinx«JXM4)oinr~xx
(vcmcc©— *>»iro»r'>inc»'r^^"-'-*o©

■4  -a  -*  *miririr^«Dooi-iwi/><iOf->

,-,,-.,*,-.-  ^*  <V   (V

•tor-^inr-omor-

C4U>fflO^-rir<t»

<r  ^  ^»  -»  ir  ir  <r  orvio

~  —i  (V

-*c»>in>#r»eoinc*
o  co  rvj  o  o  in  ©  «n
•*  co  o  o  cc  ©  x  o

(viir.(vi(vioiriO'>o  4>  in  ©

X4C(VI<£C|f1^  -*    o  c

O^^NNtCO^H  (X    O  ■-<

•-<               ~  —  (V  (V.  (V  •-  (VI  (V

nnir.  cc^och'-**  (viN-tviino^-r^xc*'
i\j  o  *  r^o-<(vnri(f>fo»vicr»'^.-»*oo>

■*«  «4«inminMsoooMin«Qo

(VI

o  o  o  rvj  it  4>  a

^H    *-t   ^-*    f*    •-•    w+   (VI

^•vDinoifiip-orvjco  43

-»^-*-*i/iincoo(\4o
^  -i  (VI

■*  ^ifl  ■*  in  i-  ■*  oo

oniMiroinon

40QO0O0O

r-4   (VI  ^  ^

(Vlin(VIOX<#in-«  IT   N>C

co-^-oxinotnx  -*  o  o-

cr-vO^f^coo— »— i  o  o  o

*-•                •-•(virvirvi  ^  --« rvi

(\i(viinx»Ti©r^-*->*-cr,X'->ino%-*4)

■VJ    ■»)    .*•    «->_«    -VJ>    >  .  \|    _#>.>■  3    -t   #■

©  -n-M|>  ■  »-♦  4>  *>'
-*coooxo4>o
(VI         —•         --•

1

<vjr»)«Mr-N-cvj>#.#  nr-in

■o-*>nn©nx>  ■*■>>

9  ^  4MI)ohh  o  o  o

—•  — »  rvj  rvj  rvj  ,*^(vj

«rfrtMomn-i'vjcoo*ii(V£-tifl
i-»7^-h-o-H(viinx(vi-^3>cJvo^^-»^i!j>

—  (VI  If!  4>   -4    O
(VI

■*<#•«  .*  .*  in  in  in  4>  x

-

1

o^ir|f-or»>in(vi^cD
H>«  ni

iO  o  >  o  H»  •*  (VI
Mon  r-  c*  ©  rvi  in
-*  *  *  -*  -*  ld  n  n  -c

-*  m  mi  —  •
f-  »■*  oi  t-  i

n  3>  >  >  o
,  —  (v  e  to
•*•  .»  -4  -*•<

lo  cvi  «  4:  cc
-»  <*  -*  44

*  if)   4>    4)  (V

X)   ■»>  -•  »»    4-

c  iv  ir  «v  ~

m  m  ir  o  i  j

ONiTrtC^

m  m  m  4>  x  o

|9M?I    4>Xi'©   (VI  4-    ©    ©  Jl

*  *  ■*•  -*-4|.ninm4>xx

-t  ♦  t>  in  m  ^
r~  c  n-  if  -#  c

O   O  ^  1*1   i)   4

OOMflulP

n  o>  n-  in  •»

o  ©  ~*  it  4>

©  <t  r-  **  -*■
r»  t«  r>-  .n  -*  .ri
©  ©  -*  o  4>  4

— <

in  r>--f>

tin  go  ^  '^  ci  o*

^)  4

(vi  ao  r-t
O  (VI -<

r>

o>  -^  <

r>-  »  rvi  >*  -•  n-

T   J>

o

o  -#  4

4f    ♦   SI  i/1    o»h

J)  -•

**^.

(VI

^^

«   ^  (VI

(VI  (V

a-  ~
it  ■*

Mf)OJ34ir^p~    X"*>

0((VI*4
IMPhffH
©  »-•  >0  (•)  ■♦
(VI  cv,

<t    4-  IT  iTi  ac

r    —  ^  cv

-in4io
p-i  ■-«  rv  <-h

r*    4    ^«   <B    ^-

t>-  Cf   «-■  ■*   CD

-♦  -^  m  ir>  cc

ii>

o^)i»-ih-m(v)-^m-<af  .o

30/1  »~  CO  ©   ■*   -O    D  "ifl'H

©-»«c)-*^-*^mmco
(vi  rvi

(Vlc^•^fvl-♦<o^cc^
orvKvioom^tvi

4-aoo©cc©>ccc

4m       '-

©  o>  •♦  m  it

0^  rC^  (*)   ^O  (VI

ON  04}   0D

(VI  — 1
I

©CJ»!*)-*'©'"-:*>-»Mrv

O-    (Vi-f^-(<t(Vi'ifV

-Hr-  cvtr^inmo>cn
co^c^mfvic^^^

o*o^^oflcr©^H      c>

•-♦  »-t  (V  (VI

o>  r»

•^  ^    (VI

CO  >o  ©  It  l*\

*>  It  <J>  ■#  (VI
1>  41  >o  S  X

en  d  (vi  i  1*  1  d      c>  £

2f
co  4«

^  ot  (vi  r*   a  4)  «v

m  o  rvj a>  ©  4)  rvi
co  c*>  cc  ci   rvj  41

CO   »£■  CD  ■

(»)  N  ©  -O  —  -

-t  3»  :fl  'Cr>   -'MX

(vi  ©  co  it  00  x

~.  01

o»  rvi  (vi  m  •*

m  xi  r\a  x>  ->

©  41

4>  r»-  ^  h-  rvi  f»

»h  4)  — •  m  — •  **i

COOX  d  o  o>  0>0  K^

2    J

o

O

_l

—  rvi

2L   I

UJ  Uj
Z  3

r  I

O

II

UJ

a

X

it

x  -•  d  n  rv
4  (v  a  <*
4t  r~

CO  ON  4  0

41  (\j  \T  n
o>  4>  4S  t>-  ^

iLlllii

3  3  2
•— •  •— *  < — 1
^   3   Z

41  r-  1-

o  o  a.
r  a  z

(»)  m  (vi  04

in  *  4t  00

y  ©  —  o  >

-<  (VI  (VI  ^

►>.  J«  M

f;  r>  4
»  e>  ^
^  fM  (VI

x  rvi
n  u

C*   CT

N-  ©  C'  ;   4!

cvi  (\i  4>  r»

J    Ort  CO

^  0(1  (VI  I  ■-•

■tNooN  a

N-   ©4>   4>  r-    *

X  ©  <-*  ^  X  C^
-1  (V|  (VI  (V)  -»   —

2

O      j

•-•     I

©    I

4>i
o
©
rvi

X

—

IT

c  — i^4>f»cM4ifi-<Mrs4(Mfi4irir.  co  x       .-«  ru

•-■  **  co  HtDHonaio  ^*.«>.-«eo«-"«"><*>xx  -♦•  o  m
©  a  <v  <xco©o4t^mo>r^eDXS-*o,,co-*inin)coo>>

tvj  >t  •*  -»-<»■*  in  in  in  x  ^>r-r^cooof\j<occi(?'0-  o

Pt  x
o  e

k  i*>  cc  <o  a

rvj  in  k  in  o-

X  ®   0*   ^  f*T

r-  oj  r-  cc  a

a;  •*  r>  •*       n

x  »-  r-  x       ©

0*  l*>  ♦  X         Pi

co  •*  o  .*

X  ~<  «■*  x
O*  Pi  ■»  x

a  x

co

IT

OJ

II?

Oj

OJ

in r-  x  x  r-  ©  co  04  ►-  in  .-in^i^rtCinxinojiooj
^c*ojr>maic*oji^me*^ttXl*cifceo-*inii>cco>
~*H<#  ■♦■♦■*■-♦  in  in  x  xr-r-cc«oo(\jxcoo,'Oe>

I

•#  •#  x  mmN  x  A  in  r>-  ©ojo— i<*>oj-«inmr^eDai

-taH  m  ■*»■/»  r—  »  ■vn'n  ■»>'>  x  » 1 01  !•>'»■-•  ♦  n  H  O1*

i  x   x  !*•  t*-  CO

Hr)**<  +  *mi/>

o

X

>■

UJ

o

z

U-l

l/>

u

liJ

c

o
s

r

oooixcoaolo^a

*-«.-+.— I     —4    -^     -^    -*    (\J

Olf

*-*  CC

in  <♦  <*5  in  -♦  o  <»>  x  ©  x  o»  o>

ojf«)in^ovoj!f»>tncf>tnr^oj

f}-»  ■*-»•♦-*■  in"  in  x  x  r-  r^  a

oj  tr

—•  c

-•01

,0   ^
CVJ

X       —>

(MJJ   IT     I     O

P;

c*-  -o  -♦  co

o  •*»  -*  r«

o  cv

o  a

hOJ

UJ

i/>  Z

UJ    UJ

•*  n  >-•  <♦>  •♦  o
-*-**-»••#  m

in  :\|  y  «-•  rn  30
<v  rr  -*  r*  a

-  i  c\i  irt
ii  m  in

(v>  j<  h  rf  iv  r-

OJ  <t>  ■♦  I**  0>   C

0(1  •*  <4  ■*•*■*  "

in  e»  ♦  at  ^

—*  pi  •-»  r4  •>

ui

71

r  z  z   z

-i  jj  xin  a  w
(»>  m  eo  m  r-  m
in  mi  x  *»-  r-  co

j  rJ  jj  *>
oj  ul  x  r*>  r-
in  in!  x  r-  r^  a

x  -*  foxxxiojcixol'tvif
oj  tn  *  Ami  o»  i*  o"  <■+  r^  r*

mm  *Mhh  cc  o  o  m  x  <t

m  -■*> '  oj  ■el  ■  x  m
mini  x  ^  r-  »
m  m  oH  r^^

•*  oj  m  m  ■♦  k
x»  i>  o  *  tj>  ^-

O    o   (M  X   X  CD

co  co

0>    O
-<  OJ

»  'vl  -•  ml  oj  n

0>  14  ©  •+  O"  f-
X    Ct  •—  SI   X)  I

-<m*>aJ0>q      ojm

x>  m

X)  ©

O    X)  »»!

©  ©  ri|
<"o  -*m

•-I    •-!  (Hi

^  f-  H
n  -»<n
oj  nm

'1

(*)  x>  OJJ

n     -HO1

oj  n  -*

at  ■*  o

m  oj

n  »1

=1

o

X)  CM    -i  AJ

>  cm  -«^  n  ■*

1

-*  — >  oj  >»  r\i
x>  co  co  «  oj
r^  -h  r-  co  a*

Oh    O   J^
OS    O    »-  D

f*-  »h  ^  co  co

4«ec  n
X>  *->  O  x»

o«  «n  ■*  xt

x  c
co  ^
o  in

m  co

OJ  >

M  f-

m  r>-

m  ■<

A  ^

f^  -<

f^  c

g  cc

m  •->  f>-  ojr-  x  ft)  :
X»-«-»o»~l)r^n
u  n  ci  o  x  c*  m

r^  ro  r~  m  X)  •-•  ■♦  .

XI  ©  -«  Pi  *~  CO  X*  1 1

&  Pi  pi  \T\  x  o>  m

OJ

a

in

in

r-

OJ

OJ

~4

)OJ  r»)

J

U>  u<  3  3

*  i

Z  2   Z   2

oj  m  m  Pi  xt

<t  h-  — •  4-  r~-

r~-  f-   f>-  u  -D

m  ©  H^

K  COO'*"

xi  >  m  ■  «Ht

*-•   i-4   w+   w+

■a  ?>  r>-

-*  oil  -*  r?
n  k  «-<
i^  h-  m  a

O   CO  CO    **!    X

X»"SOJ   X    =  <t

'NIMfl'S^'flJ)
X  h-    n  CX    CO  9>

;     w-»    r-  I    -H    -H

9     -*

X  IT.
CO  0>

m  co
o»  tr

p

■*  rvj  co  ■*■>  in
n  rt  ^  m  x  -h

CD  0*         d>  O   f

:

MOW

r~  f~  oji  «|
x  o>  m

A©
x  o*

Ullfllfl^)

fw  xj'OjcH-l

•*   X  9-  P}'*

:ts

■"r

ZCC

3

a  a

a.

r>  cc

rvj  on
m  o

o  -*  o>  r--       •*  X)

*-•  <c  h-  f--       co  (vi
r-  ^  r-  ■*       co  o>

N<r-ff  xi  x>  x>  id

cr  co  &  o  od^  fvj  ^

hNo  h  (vi  •-•  o  o>

~+  — —i  <\l  —•  (VI  -<  — •

x>

CO

75

IT
o

—  xt  m  cc  cc  m  r-  in
M  -•  x>  it  r-  f«-  o>  (vi
00       x>  r--  r~  ■*       xic

»  •♦i/i  ec  «»5  (\|  n  co

cc  cc  o>  o  ceo  x>  e>

•-•  (VI  O*  •-«  NO  0*  0*

»*    •-  »»  (Vi  r-   (Ml  —

tn
co

(V  cc  r~.  <t

^*in

o>  ♦  (vi  r-

tn  t\i

f»»

(VI

vO  IT  f-  h-

nw

co  t-  •—  o

co  o

co

(VJ

x>  r-  r-  -*■

0DO>

— <  <\i  m  — i

A   ~  r*    (VI

(VI  o

co

0>

p4

ip^i

-«vj

■-I

in

co

m  w

X   — •    XI    X»  XI

♦  xi  ^  b-  rt

o  x»  r>-  »-  r-
ni

o*

coo»

CO  —  (VI  IT

-4  oj  m  o
— *  — •  rvj

T

in  n      —•
r«  ssf      so

^

♦
co

x»  o  o

O   ■*  CO

c  x«  o»

x»  (vi  m  xi
-»*--*  x>

(VI  !*)  rvi  o»  IT  XI  X)

— •  •-»  -*  in  it  r~-  co

(T   to  ^SNh
i    «M

O  1/

■♦  o>

—  X)

o

Pi

->   -.  o
f^  -•  o  /l

CM

X)  X)  -*
—  CC  0
CO         3

*  ir
c  cc

cc

o
f>  X

o  xi  n
o  cc  c
o>>  t-  o

a  —  —  r.

ec  o»  r-  o

ec  i»»  o

■*  o>  xi  <n  in
-•  f-  »  ji  »
in

Ci  ^

o>  h*c*  n  co

CO    I  D  3    ?•   SC

i  co  x>  (vi

Kl-K  ►-  I-         »-!»-»-

Z  Z   1/5  2Z         Z         ZZ

Ul  UJ   l*J  jj   aJ          UJ          uJ    LU

cc;iu  ,  itccii,cii,ccu.ii.uj

►-   »-    LkJ    >-«  iMK»-l-lt-l-l|-hHHM

if  fa  J  2:tr(/sio3i/itrzz2

co  t>  in  in  in

o*  in  «-t  t--  co

X)   o   r^  r-  r-

XI  D    -»   -*   •*

a  ir  «-«  r-  cc
xi  xi  i-  f~  r-

o  ir,  ■«  n  ■«
o>  m  —•  r»-  x

o  x>  r--  r*-  r>-

0*  (*)(")•-«  <—
ao  si  -*  t>»  »  »«■■

in  <»

-•(VI

o>  (»>  r^-
in  (v)  ♦
cj  *  •©

r«  (VJ

x  m  -«•

IT  (V)  -*

-n  (VI

r-  (V  cc

in  evi  en
co  a  o

X)  ^(Vl  (VI

cp  in  a  (vi

n  -4

x)  in  —

x>

l»)  &

»n  m  r-

o

-»  cvi  (»)  c

(V  o

— •  CD    X>

r>

r4    -«r-«    fVI

•-•  —

»<VI           ^

■— »

r'

•^   r^    (VI

ro  4

X)    -•    *

o

in  a  (vi

O  9

in  -*  <s>

ao

(vi  m  o

(VI  »

'^  ^j if)

ao

«-<  — •  (VI

t(V!           ■"-!

i/i  -in  oo  n  -J
a    c  CMC  c;
-i  (vi  n  o  ^  (vi

- 1  ^  Hi  (VI  (V  -<

rf  (*>  o  •->  co  cc  -!■

o>  o>  o  un  cc

ei  -+  (Vi  o  »«  o

•■ «  ri  «<  (VI  (V\)  •-•

rvi  CVl  (VI   30   T  0»

in  in
U;  uj

a

3   3  «  O-i  Z   2

c  a  a  a    |  ac  i/ii/ii/i  i/iu(

I  *~  e-  o  a

-•    3  3  »-   I-

«/)  in  in  </5  «n

ix,

»— «

a

3

(*1  00  P»  3»   (*>  if'

>  SO  ^  3  X

•-•  (VI  o>  ■-*  (VI  f

-^  -»  -*  (VI

(VI

•»
cc

(Vt

-*

SO

X    SO    X

H     C     »^

(VI  -<

ec  *  —

r-i  c  a

•-I  (*>  o

(VI  ^<

cc

CD

X

i
^  o  in  ro  (vi  i

(VI         •-•

z  z

Ul    aJ

O  C  Ui

(T,   IT  2

I  Q

z    I

2   2    2

a

Of)  O  2

O  Q

>«  <t

UJ

a.  a.  a
v  >-  >
i-  »-

U)

UJ  LJ

■♦  a

cc  x

in  so  (vi  >o  in  o

"»>  -•  r>  ;(•>  (VI

U.  U/  U.  U ,  U.I  U.   UJ  U.  UJ  U )   UJ  UJ         UJ  U.1   L  J  Ui

?  ZZ'ZZZzZt'Z't  2  2   2'  2

Z  1/1  K>Q

K  O  ©  ~

<t  U  O  <J  c

U.'  UJ  UJ  UJ  UJ

a.  x  a.  oi  a.  u
v  >>  >-  >  >-

ec  ■♦  in  rn  <x

r^  — «  cc  cc  •-•
h^  (1>  CO  h-  cc

•-H      f-^     ^—      •— *

SHOE'S
M  (Vl  (^  (Vi  (VI  C

hmo4  — •  -i

3     2   2

2   2    3     2

I  u.

z

o

CD

CC

f-

o

t^

■*

in

-*

o

r-

o

fM

#— •

CM

*  od  cc   cc  o"  o  o  in  cr

•*  ^»  tn  cm  o>  oS  cm  o>
oneccriMOOKO

-4-  cC  CO   CO  O-   o  0*  ■♦  CO

•-*   ^>  t-i  CM         •—  —>

a  net  o>  NKNhco

•*  COCO   COOO   *  ♦  CD

^■t    »-l    ^*   ^H         |    ^*   9-4

o

CD
2
>-
(/>

O

UJ

u

z

UJ
CC
UJ

Ul

if)

u

UJ

c

O

B
i

O  CO  ^  CM

(f.  ^  o  CO

o»  en  e»  cm

r-i  — 4     "\J     ■-*

a  »-  o>  co

•-«  ^-«  i-M  •— I

i

o         —•<

ao       f*- 1

(J>  CD  C>

CM         if  i  IT

h-        o  CO

(M    r»»  :\1  -*   — i

J\

x  I  z  a   i

CC    2

Z3    <

LiJ    Uj  u     Uj

a.  o.  a.  x

oco-*o<ocMr-<o  n
■n-o: >  '»  •V-'H  3  M'  >
co  co  co  o^  o*  <d  *♦  co       •— •

cMr--  n(7,if)H,°'^,,'(VI

r-cocforvjf-oco-*^

X  CO    XcoOsD-^-CO*— t

^••-•«-*^h         ~*  •— i  CM   ^*

-^  <o  cm  ca  in  r»

f>-  CO   C»>  0>  CM
■O  CO    CO  CO  O   -D

O  -*  tM

IOC
r^  o  r>

it

co  fir  in
r^  in  r^
r^  CM  CM

O  <\»  CM  CW  3»  *C

."M  CO  CO  i*-  *-   M

-*cq  a  a  cv
c  a-  x  cd  cr

^-<     #-•     *■*    ^H

**cm  a  c*  cm

■4  in  cc  co  cr

fM  -*

z  z

UJ  UJ
UJ  Uj  UJ  c  c
2  3T   Z  3  3  Z

i-i  _«-.(-  t-  _
X    X    3"   VI  if   2

3  >

♦  «-

~  O  O

Ul   ♦  «-

a  a  x

>->•>-:

»-  f-  ►-  3  >  5

mo

»i  r«4

-*C0

C*   T

*o

co  a

»*  —

«.  r-  ai  co

co  <t

»»»  ^.   p«.   f«-

X    X  CO  CO

if.  ir  (jo  if  if  o:  if  if  i/>  «

J  ji  o

D    ,0

!

CD    XI

x«  r>-

h-  (V  oj  <*>  cm
x>  ©  <ioh
X  M       x  (vj  x>

f\|  CO  «V  r+

•O  x>  cy  ■*

N    Iff   O

CD    OCj    O

o1  o  ir

o  •-*  -h

77

o  *

CO  IP

•"-•*#        x  en  X)
«h-       •#  r*-  <f

OD   »*•        CD  •— •  l/>

*-  r>-  r^  o

r-  cc  cd  o

XI  0»   -h

eo  a  o

«04

a  —  ~

o  xt
co  -*
x>  r»

ox  ©      ■*  a  •*       r«cvi1oo>
JVJO        — •  O  a*        Hi  <oo  w

CD«-*         (COW         N    too   o

x>  <r  —

CO   CC   9

CO  »   •*
0*  O  •-•

(*>  .c*  -^  >c

«|nk

(\i

*»•♦  #  x

>    >  ■!«-

M

■c  r--  o  (\j

-o  nj  -c
— »  — «  — t

CO

fw  ^

i

CO   f*)   fVI  Xi    f\J  X*  C^    **)

3M>       -i  o  a  -•  o

acKvi       hixooo

tn  co  o

X    »■>

r*  <r  n

>   O    -•

--  x  r^

X  r»  rv  ~

r^  —  X)  r~  rvj  r^  ru

X  F*-  (\i  ~*  *©  IV   X

-4  x>  <")
(V  x

3  a  ■

©  o  -«  ©

■X   —  -»   IT   <V

vfl   N   O    OX

fyrvjcvtvimr^ooo*

XXTXXr^r^r-

I

-<  n  o  o      r-  43 :  ,o

—t  ■♦  ej  co      05  x>  o»

r--  ©  ©  xi      i/i  (\i  in

in

-h
cc  IX

C\J

x»

m

©  -h  .*  .-•  r»-

uti  o~  t*  (vj  cc  h-

»t>    X>   O    f*1    O    X

o»  -*  m  :m  o»  tvi  <»i  «c  -a-  r-

O   D    H   VM     ©   ©<•-    -Hi  O

IT  f-
CVJ  it)

x<  ■*

ro  0»

-D

r^  at  in  r^  cc

x
x

x

-*  o  o  x
(\i  o  o  r>-
x»  r-  x  cv

n  t*i  n  x  i
— •  X  o  -o
x  f>»  co  (M

r»-  n»  X  f)

•t   xt  r»  o>  in

X  CC  CD  cc    o

x  x  a*  oa
■*f-  »  m

n  x  x  o

x  n  x  m

•*  rt-  a  ir

x  x  a.   o

•*

in  od  o

h  -• n

m

X    X  ■   0*

31    ©    -Hi

O"

.—    -H

n«-iir<c<£i*>«>

r-c  f  MA  (VI
C)  IVJ  <£  cc  <v  <o

A't  r-

1'0  0  3hc|    o  x>  t>   n   X

-Hi  »-l    •— t    — 4^-4  I— H  t-#»-^  — t

in  N  O

x  x  c^

xt  »-•  CU

»     ©     -H

!    — <    -HI

f:  ef  «»>  ©

vc  i>  o*  in

XX©

•  p-  K  »-■  K  I-
2Z2ZZZ2Z22

tUJoJUillJUIlLilL)     U-    UJ    LlJ

cacacacccc
v  v  ir  of  ir

UJ
C    ||]

(/•(/■   W  (/■   2

>    >    >    >    >

•»  ».    00   (T
>  :>  >  >

J>   fVJ  1%  ©

ih  ec  ir  x

<>  ^  oc  ff>

$

X   tVi  •€.
<C  tf1  <c

co  ^  o  o*

— T    «-H  «44    a-H

©  ©  o  •©  o  -a

-H   S\

a  a

it  ut ■  r>-  a  ^

(*^  X    X   X   Cf

o

x  in  xi  o>  -h  in

©  H   X    »'3»'»

b.'   U    li-i  UJ   li.  lu  C  UJ  ti.

~Z  3  ~  Z    Z.

t-(  M  ^  »-H    »— I

5  J  W  I  Z

?    2    3    2

—  >_  ^_  _

o

UJ  u.

ct  a  ^

UJ  iil  1
r.  E'O  3
2  3  2

I

><  ISJ

i  <  t-i
JT   if)

I  1  t-

.3   C  X

■    X    ~X   X    XX

n

x

z  2

UJ

c

/,  lr/ 1  I  I  J  I  I  I  tf

UJ  U.'  Uj  U.  UJ

Z    7    T  H    7

ois  ct2  tr
u.  2  «»  ►-  <i

2  jiio.>
h4:j£iz  »-'

V  Ifi  M|0  N  N    N  r.1

2
UJ

UJ  c
?  3

3:

UJ

©  M

©  —

©   *¥
■—i   — i

a>  -4  t/i

•c^  -♦  -<

•4  -x .

3<

00

in

■C

in

in

t>   a  x>  n

f.  rr  r.  «

in  in  in  -*-

^->     — •    —1    — 1

-1

o

X

2

>-

'

(/>

<£   <0  <C  O

o

»-

■♦-*•*-*

u

z

J.

UJ

X

UJ

c

o

X

z

m
«n

rvi

Hi

O    J1  o   — <     <»■

o  o  o  o   m  r

z

UJ
UJ  U.  UJ    c

Z     Z  Z    Z     V>2

N

C    t-  i-i  «t    —  I
J-    1   /)    >     4

♦       4-    ♦     4-      4-    4

IT

in
m

Sfll    U  »  33

in

2  Z  ZZZ

UJ         UJ  UJ  UJ  u.

C  U.    C  C  C  C

3  2  3  3  —
t-  «  »-  f*  ►-

>  >
«•   «-

J  3  aJ   i.  'i  I
4     4-     4-     «•     4-    «

4

tvj  <\j  rvj  <vj  (vi  (\  i

B    D   D    XI    D   I

I-  H   K  H-  t-

ZZ  2  2  Z  2;

UJ  UJ  UJ  uli  U-»

c  c  c<  c  c

3  3  ~  ~  r

-•  1  *  A  *  -

3

o
n

(VI

-♦•■*  CO  f»

ao  xi  ry  -<

:  tvi  rv

2       2  2

UJ         UJ  UJ

c  u.  c  c  u.  u

-)      3<      — .  ^     T      "»

\f y.  <s  ir.  iru:  i^s  i/i  wz  z

2  x

U  U

_l  c

h            <a  z
a.           o

<    ♦  <    *-   <  ♦

rvj

(v;

om  (\)  (\)  co  -
r>  -i  j»  :»  ao  T
(VI  oj

2  2    Z

,   UJ  uJ  UJ
U_  UJ  c   c   c   u.

222332

z   z  in  ur  <r,  :.

11

r--  rd  <  ij   a  a
«-  *    r  1    r

JI.IOGRAPHIC  DATA
,rET

1.   Report  No.

UIUCDCS-R-73-611

■  ^untitle

calc2  -  A  PLATO  IV  Lesson  on  Differentiation

UlhlH    s

Axel  T.   Schreiner

"i-rlotming  Organizat  ion  Name  and  Address

University  of  Illinois  at  Urbana-Champaign
Departments  of  Computer  Science  and  Mathematics
Urbana,   Illinois     61801

3.  K

ec  '|M<-,, i

5.    Kcpurr   I J  j  t  < ■

December  1973

6.

8.    Performing  Organization   Rept.
No.

10.   Project/Task/Work   linn  No.

11.   Contract    ('rani    No.

•>ring  Organization  Name  and  Address

University  of  Illinois  at  Urbana-Champaign
Departments  of  Computer  Science  and  Mathematics
Urbana,   Illinois     61801

13.   Type  of   I'm- port   &  Period
Covered

14.

Supplementary  Notes

(Abstracts

calc2  is  a  lesson  differentiation  operating  on  the  PLATO  IV  computer-aided
instruction  system.  A  symbolic  differentiator  guides  the  student  through  all
steps  of  the  differentiation  of  a  student-chosen  formula.  The  report  describes
the  details  of  the  table-driven  symbolic  differentiator.

1  Key  *orls  and  Document  Analysis.     17a.   Descriptors

calculus

computer-aided  instruction

symbolic  differentiation

derivative

Identifiers    Open-Ended  Terms

II  Field/l

■  iiemi-nt

RELEASE  UNLIMITED

T»M  NTlS-35   II0-7C

19.  Security  Class   (Thi-
Report  i

_, UNCLASSIFIED

20.  Security  Class  (This
Page

UNCLASSIFIED

I?}.

>o.     >[    I'M

84

??.

USCOMM-DC    4G329-F-

"*.

'

in    iww

H

UNIVERSITY  Of  ILLINOIS  URBANA

3  0112  064441527

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
