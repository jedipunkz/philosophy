---
source: "https://archive.org/details/interactivetable745tind"
title: "An interactive table-driven parser system"
author: "Tindall, Michael Harry"
year: "1975"
captured_at: "2026-10-23T04:17:49Z"
updated_at: "2026-10-23T04:17:49Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato"
plain_text_url: "https://archive.org/download/interactivetable745tind/interactivetable745tind_djvu.txt"
public_domain: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# An interactive table-driven parser system

- 著者: Tindall, Michael Harry
- 初版: 1975
- 情報源: [archive](https://archive.org/details/interactivetable745tind)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

UNIVERSITY  OF

ILLINOIS  LIBRARY

AT  URBANA-CHAMPAIGN

Z-^^^     UTlJCDCS-R-75-7i+5

}u).  7/y

TXrf^^/

IHii  LIUKAKY  OH    iH£

SEP     8  1975

iirjivfriTv  OF  II I  iMOr.

AN    INTERACTIVE   TABLE -DRIVEN   PARSER    SYSTEM

by

Michael  Harry  Tindall

August,  1975

Digitized  by  the  Internet  Archive
in  2013

http://archive.org/details/interactivetable745tind

UTUCDCS-R-75-7^5

AN  INTERACTIVE  TABLE -DRIVE  PARSER  SYSTEM

by

Michael  Harry  Tindall

August,  1975

Department  of  Computer  Science
University  of  Illinois  at  Urbana-Champaign
Urbana,  Illinois  6I8OI

*This  work'  was  supported  in  part  by  the  IBM  Corporation,  Thomas  J.  Watson
Research  Laboratory  and  the  National  Science  Foundation,  Grant  No.
EC -^1511  and  was  submitted  in  partial  fulfillment  of  the  requirements
for  the  degree  of  Master  of  Science  in  Computer  Science,  August  1975*

iii

ACKNOWLEDGEMENT

I  would  like  to  express  my  graditude  to  my  thesis  advisor,
Professor  Thomas  Wilcox.  He  provided  many  s\:;ggestions  and  much  useful
guidance  during  my  work  on  this  project.

Thanks  are  also  in  order  for  Mike  Mj.lner  and  Bob  Barnett  who
worked  on  the  initial  implementations  of  Fortran  IV  and  COBOL.   They  provided
very  useful  feedback  on  the  operations  and  design  of  the  overall  system,
and  their  assistance  is  greatly  appreciated.

Also,  a  word  of  thanks  goes  to  Cinda  Robbins  for  her  excellent
typing  of  this  manuscript.

(.

iv

TABLE  OF  CONTENTS

Page

ACKNOWLEDGMENT iii

1.  INTRODUCTION •  1

1.1  Organization  of  this  Thesis 1

1.2  The  Compiling  Environment  1

1.3  The  Interative  Compiler  2

2.  THE  TRANSITION  DIAGRAM  PARSING  MODEL k

2.1  The  Basic  Model h

2 .2  Extensions  to  the  Model 9

3.  ASSEMBLER  SYNTAX  SOURCE  INSTRUCTION  SPECIFICATION  ...  10

3-1   Introduction  and  Chapter  Organization  10

3.2  Overall  Organization  of  the  Syntax

Specification  ik

3.3  Parser  Action  Instruction  Description  22

3.^   Description  of  Valid  Parameters  Used  in  Action

Instructions 36

h.        THE  PARSER  TABLE  INSTRUCTION  FORMS k^J

U.l   Notation., i+Y

U.2   The  Table  Instructions 48

5.  THE   COMPILER'S  TABLE  MAINTENANCE    SYSTEM 57

5.1  Maintenance   System's  Purpose 57

5 .2  General  Operation  of  the  Maintenance   System    ...  57

5.3  Logging  into  the  Maintenance   System .  58

5-4       Preparing  the   Syntax  Table  for  the  Compiler

System 64

6.  FUTURE  DEA/ELOIMENT 67

LIST  OF  REFERENCES 69

APPENDIX •       70

CHAPTER  1.
1.     INTRODUCTION

1.1  Organization  of  this  Thesis

This  paper  discusses  the  design  and  implementation  of  a  table-
driven  syntactic  parser  with  concurrent  static  semantic  checks  to  he  used
in  an  interactive  compiling  environment.   After  a  brief  introduction  in
this  chapter  to  the  compiling  environment  in  which  the  parser  is  to  be  used,
and  the  the  general  operation  of  the  parser  system.  Chapter  2  will  present
a  model  of  the  selected  transition  diagram  parsing  technique.   Chapter  3
contains  documentation  on  the  parser's  assembler  language  syntax  source
instruction  specification.   Chapter  h   contains  detailed  documentation
on  the  form  of  each  instruction  that  is  used  in  the  actual  parser  table.
Chapter  5  describes  the  ta.ble  maintenance  system  that  is  used  by  the
compiler  system,  and  the  thesis  concludes  with  a  few  comments  about
further  refinements  that  can  be  made  to  this  parsing  system.

1.2  The  Compiling  Environment

Recently  a  project  at  the  University  of  Illinois  at  Urbana-
Champaign  has  been  under  way  to  automate  the  teaching  of  the  basic
Computer  Science  courses  by  utilizing  the  PLA.TO  IV  computer-aided
instructional  system  that  is  being  developed  on  this  campus  [h].      This
computer  system  features  an  excellent  graphical  display  terminal  and
fairly  sophisticated  computer-aided  instructional  software  support  for
the  writing  of  instructional  lessons  and  the  corresponding  course
curriculum  [5].   The  curriculum  that  is  being  implemented  will  teach  new

2

comp-uter  science  programming  language  concepts  and  constructs.   Specific
programming  detail  on  a  variety  of  languages  (e.g.,  FORTRAN  IV,  PL/1,  COBOL,
basic)  is  available;  students  will  progress  at  their  own  speeds  thro-ugh  a
fairly  flexible  course  structure  [2],  An  important  part  of  the  system  is
an  online  compiler  in  which  a  student  can  easily  and  conveniently  try  out
new  programming  constructs  immediately  after  learning  about  them  in  an
instructional  lesson.

The  remainder  of  this  paper  will  discuss  this  compiler  and,
in  particular,  the  parser  system  that  is  used  in  the  compiler.
r

1.3  The  Interative  Compiler

A  number  of  design  criterion  emerge  from  examining  the  environment
for  this  compiler  system.   First,  the  compiler  should  be  as  interactive  as
possible  to  utilize  the  PLATO  IV  system  effectively  and  to  maintain  a  desirable
computer-aided  instructional  environment  for  the  student.   To  accomplish  this,
the  compiler  compiles  character-by-character,  that  is,  each  single  key
press  by  the  student  using  the  compiler  is  examined  immediately  as  the  student
types  it  in;  thus  the  compiler  keeps  up  completely  with  the  student  and  detects
programming  syntax  errors  as  soon  as  possible.   The  student  is  able  to  edit
his  program  by  moving  a  cursor  through  the  program  on  the  screen.   The
compiler  moves  with  the  cursor,  compiling  when  the  cursor  moves  forward  in
the  program,  and  backing-up  ("un compiling",  i.e.,  resetting  the  lexical  and
syntactical  analyzers  to  previous  states)  when  the  cursor  moves  backward  in
the  program.   Thus,  the  compiler  is  highly  interactive  and  easy  to  use.

A  second  design  criterion  for  the  compiler  is  that  it  be  multilingual.
To  accomplish  this,  the  compiler  is  completely  table -driven;  to  allow

3

another  language  to  be  recognized  by  the  compiler  system  and  used

by  students,  a  language  designer  must  merely  fill  in  a  new  set
of  tables  and  provide  an  execution  supervisor  system  for  the  actual
interpretive  execution  of  compiled  programs.   This  paper  is  concerned
with  one  of  these  compiler  tables,  neunely,  the  syntax  parser  table.

A  third  design  criterion  for  the  compiler  is  that  it  provide  a.
a  high  and  sophisticated  level  of  error  diagnostics  for  the  student
when  a  syntactic  or  semantic  error  is  detected  in  the  program  by  the
parser  system.  Since  the  intended  users  of  the  compiler  are  beginning
students,  the  error  messages  must  be  direct  and  to-the-point .   To
accomplish  this  goal,  an  automatic,  interactive  error  diagnostic  system

has  been  designed  and  implemented  [6].   The  important  point  for  this
discussion  is  that  this  automatic  error  system  is  driven  by  the  compiler'.
syntax  tables,  that  is,  it  is  essentially  language-independent.   Thus,
it  is  apparent  that  these  syntax  tables  are  a  very  important  part  of
this  compiler  system.

We  now  examine  a  model  for  the  transition  diagram  parsing
technique  used  by  the  compiler  system.

CHAPTER  2.
2.   THE  TRMSITION  DIAGRAM  PARSING  MODEL
.  2.1  The  Basic  Model

The  syntactic  analyzer  used  in  this  interactive  compiler  is
based  on  the  transition  diagram  systems  first  introduced  by  Conway  [1],
and  recently  formalized  by  Lomet  [3].  A  transition  diagram  system
consists  of  a  set  of  nested  push-down  automata  (NPDA)  that  have  the
capability  of  invoking  one  another.   The  remainder  of  this  section  will
present  first  an  intuitive,  graphical  description  of  transition  diagram
systems,  followed  by  a  slightly  more  formal  description  of  the  transition
diagram  model.

A  key  concept  of  a  transition  diagram  parser  is  that  of  the
parser  "STATE":   the  STATE  is  a  descriptor  that  maintains  information
about  what  input  has  already  been  accepted  and  what  further  inputs  would
be  acceptable  to  the  parser.  While  "STATE"  is  a  very  important  concept,
it  is  a  very  easy  thing  to  visualize  and  implement  in  a  transition  diagram
system.   For  the  rest  of  this  paper,  STATE  refers  to  this  "state  of  the
parse".

"The  action  of  a  transition  diagram  parser  is  to  examine  the
"possible"  or  "acceptable"  parsing  options  (determined  by  the  current  STATE
and  the  (transition  idagrams)  along  with  the  current  input  token;  based  on
this  information,  the  parser  will  accept  the  token  by  updating  the  STATE
information  and  asking  for  a  new  input  token,  or  reject  the  current  input
token  and  signal  a  syntactic/semantic  error  if  the  token  does  not  satisfy
any  of  the  available  options.  Note  that  all  the  parsing  options  are  defined

in  terms  of  "tokens"  only,  that  is,  the  tokens  that  are  acciomulated  and

output  from  the  lexical  analyzer.  This  process  can  be  conveniently  shown
graphically  as  follows :
Let

SI

denote  STATE  "Si";  each  branch  out  of  a  STATE  corresponds  to  a
possible  syntax  option  for  that  STATE:   these  branches  are
labeled  with  their  particular  syntactic  option  requirements
(these  labels  will  be  described  as  the  paper  progresses).
Then,  a  PL/l  "GOTO"  statement  can  be  shown  in  transition  diagram  form  as i

"GOTO"

'GO'

S2

'TO'

[label-name]  /"^)A  ";"   ,  ^^

This  is  interpreted  as:   if  the  parser  is  in  STATE  SI  and  the  current
input  token  is  "GOTO",  make  the  state  transition  to  STATE  S3;  if  the
following  token  is  a  [label-name],  move  to  STATE  Sk;   if  the  token  after

that  is  ";",  accept  it,  and  (in  this  case)  accept  an  entire  "GOTO"  state-
ment.  Note  that  if  none  of  the  branch  options  for  the  current  STATE
satisfies  the  current  input  token,  then  that  token  is  in  error,  and  the
normal  parsing  error  condition  should  be  signaled.
Another  example  is  a  PL/l  "IF"  statement:

SI

'IF'

S2

<conditional"expr>

S3

"THEN"

$k

<statement>

<statement>

ST

In  this  case,  notice  the  references  to  <conditional-expr>  and  <statement>
as  labels  on  option  branches:   this  indicates  that  if,  for  example,  the
parser  is  in  STATE  S2,  then  when  trying  to  accept  the  input  token,  it
should  refer  to  another  transition  diagram  in  the  system,  corresponding
to  <conditional-expr>;  after  a  <conditional-expr>  has  been  parsed  and
accepted,  the  parser  should  then  return  to  STATE  SS,  having  successfully
satisfied  the  <condtional-expr>  option  branch  out  of  STATE  S2.   This  is
an  example  of  one  transition  diagram  "invoking"  ("calling",  "referring  to")
another .

One  more  example  is  needed  to  illustrate  another  important
feature  of  transition  diagrams:   a  transition  diagram  system  which
has  been  invoked  by  another  transition  diagram  has  the  capability  of  return-
ing to  one  of  a  number  of  possible  STATES  in  the  invoking  transition  diagrsim.
A  good  example  of  where  this  is  useful  arises  in  trying  to  parse  a  (simplified)

PL/I  <conditional-expr>  as   follows:

expression  parentheses

(    (    (    I  +   100    )    *  J   )   =  K   )

L

conditional  expression  parentheses

When  the  initial  parentheses  "("  are  examined,  it  is  not  knovm  whether
they  are  part  of  the  overall  conditional  expression  or  are  part  of  the
inside  simple  expression.  The  technique  to  solve  this  ambiguity  is  to
allow  an  "invoked"  transition  diagram  to  return  to  more  than  one  STATE
in  the  "calling"  transition  diagram,  depending  on  what  tokens  are  found
later  on.

The  diagram  for  <conditional-expr>  can  be  drawn  as :

8

In  this  case,  each  time  <conditional-expr>  is  invoked,  2  return  STATES

n
are  specified  (i.e.,  Rl  and  R2);  when  the  parser  reaches

,  it

returns  to  return-state  n.   Note  that  the  way  in  which  this  <conditional-
expr>  has  been  drawn  corresponds  to  assuming  that  the  initial  "("  belongs
to  the  overall  conditional  expression;  if  it  turns  out  they  actually
belonged  to  the  first  simple  expression,  then  the  parse  is  resumed  at
point  5,  which  accepts  the  assixmed  "conditional"  parenthesis  and  continues
parsing  the  simple  expression.

As  a  more  formal  description,  a  transition  diagram  system  consists
of  a  set  of  nested  push-down  automata  (NPDA)  that  have  the  capability  of
invoking  one  another.   Each  NPDA  is  capable  of  reading  a  portion  of  the
input  string  and  accepting  or  rejecting  it.   Lomet  calls  the  NPDA  that
are  capable  of  being  invoked  "submachines";  the  initial  STATE  of  a
submachine  is  known  as  its  "entry"  STATE  and  a  submachine  is  invoked  by
the  use  of  its  entry  STATE  number  by  another  NPDA.   This  results  in  the
invoking  STATE  being  saved  at  the  top  of  a  parser  stack  and  the  parse
resumed  at  the  new  entry  STATE.   Each  submachine  also  contains  one  or
more  "exit"  STATEs;  when  an  exit  STATE  is  reached,  the  top  of  the  stack
together  with  the  particular  exit  STATE  determine  the  STATE  in  the
original  invoking  NPDA  with  which  to  continue  the  parse.  An  error  in
the  parse  is  detected  if  an  NPDA  reads  a  token  from  the  input  string  for
which  there  is  no  corresponding  STATE  transition  that  the  NPDA  can  make.
The  reader  is  referred  to  the  discussion  by  Lomet  [3]  for  further  technical
details.

2.2  Extensions  to  the  Model

The  preceding  discussion  in  this  chapter  describes  a  parser
model  that  is  sufficiently  powerful  to  recognize  all  deterministic
context-free  languages  [3]'     A  few  extensions  to  the  model  have  actually-
been  included  in  the  compiler's  parser  system  to  enable  the  handling
of  context-sensitive  static  semantic  requirements  such  as  proper  and
consistent  declaration  of  attributes  for  an  identifier,  consistent
references  to  declared  array  variables^  etc.   These  extensions  include
allowing  auxiliary  memory  variables  to  be  utilized  by  the  parser  (for
operations  such  as  counting  the  number  of  subscripts  in  an  array  reference),
allowing  the  labels  on  any  branch  in  a  transition  diagram  to  refer  to  any
of  these  auxiliary  variables  or  any  symbol  table  field,  and  allowing  special,
language -dependent  semantic -subroutines  to  be  invoked  at  appropriate
times  during  the  operations  of  the  parser  system.

To  summarize  the  activities  of  a  transition  diagram  parser:   it
must  be  capable  of  requesting  that  a  new  input  token  be  read  in;  it  has  to
be  able  to  test  the  current  input  token  to  decide  which  labeled  branch  to
follow  out  of  the  current  STATE;  it  must  have  facilities  to  manipulate
a  return-state  parser  stack;  and,  finally,  it  must  be  able  to  perform
any  semantic  analysis  that  is  required  by  a  particular  language.

10
CHAPTER  3.
3.    ASSEMBLER  SYNTAX  SOURCE  INSTRUCTION  SPECIFICATION
3.1  Introduction  and  Chapter  Organization

3.1.1  The  Syntax  Parser  Table  Description

A  table-driven  parser  system  has  been  designed  that  incorporates
the  actions  described  in  the  preceding  chapter.   The  syntax  table  is
actually  an  encoding  of  the  programming  language  syntax  productions  in  a
small,  interpretable  instruction  set.   The  parser  of  the  compiler  consists
of  a  routine  that  interprets  this  instruction  set  in  an  appropriate  way
(this  routine  can  be  called  the  "table  interpreter"  or  "table  driven"
routine).

The  parser  has  control  over  and  maintains  certain  tables  and
data  structures  within  the  compiler.  All  of  these  tables  and  structures
are  located  in  a  region  of  memory  that  is  referred  to  as  the  parser
storage  area.   One  particular  variable  that  is  maintained  is  the  parser
STATE  variable.   This  variable  always  points  to  some  instruction  in  the
syntax  table.  As  the  input  string  is  parsed,  this  state  pointer  variable
is  updated  (i.e.,  moved  to  point  to  a  different  sequence)  to  reflect  the
current  state  of  the  parse.

The  parser  maintains  two  stacks.   The  first  is  the  regular
parsing  return-state  stack  which  is  used  when  one  NPDA  invokes  another.
The  other  stack  is  called  the  "variable  stack";  it  is  used  by  a  language
implementor  to  save  miscellaneous  information  as  the  parse  of  an  input
string  progresses;  the  language  implementor  has  control  over  the

11

manipulation  of  entries  in  this  stack:   when  to  put  new  entries  on  the
stack  and  delete  entries  from  the  stack  (synchronized  with  PROC  entry
and  exit),  or  change  the  value  stored  in  an  entry  on  the  stack.

There  are  also  certain  tables  that  are  maintained  by  the
parser.   These  include  the  compiler's  symbol-table  and  the  compiler's
block  structure  tables.   Entries  from  these  tables  can  be  examined  or
modified  by  the  parser  with  the  complete  control  of  the  language  implernentor.

The  action  of  the  parser  system  in  the  compiler  is  to  examine
and  interpret,  beginning  at  some  location  in  the  syntax  table  (determined
by  the  STATE  variable),  the  table  instructions  that  specify  the  acceptable
syntax  for  the  programming  language.   The  instructions  are  interpreted
sequentially  unless  a  particular  instruction  modifies  the  parser's  table
instruction  pointer.

For  convenience  in  specifying  the  instructions  in  the  parser
table,  an  assembler  language  representation  for  each  table  instruction
has  been  designed,  and  an  assembler  program  is  provided  as  part  of  the
compiler  system's  maintanence  utilities  (chapter  5)  that  translates  the
assembler  source  language  representation  into  the  actual  table  instruction
form  that  is  used  by  the  compiler  system..   The  remainder  of  this  chapter
documents  the  assembler  source  language  specification  requirements,  and
chapter  h   documents  the  actual  table  form  of  the  parser  table  instructions.

3.1.2  Chapter  Organization

This  chapter  documents  the  assembler  source  language  that  is
used  to  specify  the  syntactic  requirements  for  a  programming  language.

12

Once  a  syntax  source  representation  for  the  language  has  been  prepared
(using  the  compiler  system's  table  builder  option),  an  assembler  program
will  translate  the  source  representation  into  a  compact  table  form  to  be
used  by  the  actual  compiler.

The  chapter  is  divided  into  three  major  sections  as  follows :
Section  3.2:   Overall  Organization  of  the  Syntax  Specification.

Discusses  the  proper  ordering  of  instructions  for  the  syntax
specification;  discusses  the  function  and  use  of  procedures
(PROC-EKD  blocks)  in  the  specification;  explains  the  form
and  the  use  of  the  mnemonic  definition  (DEFINE)  instruction,
the  storage  allocation  (ALLOCATE)  instruction,  the  different
error  NAME  instructions,  and  the  purpose  and  form  of  the  FINAL
PARSE  STATE  instruction.
Section  3-3'-   Parser  Action  Instructions.

Discusses  the  form  and  purpose  of  the  instructions  that  are
used  to  control  the  actions  of  the  parser:  SCAN,  GOTO,  CALLI,
CALL,  RETI,  RET,  BC,  SEMA,  and  the  auxiliary  environment -
changing  instructions  (ASSIGN,  MASKON,  MASKOFF,  ADDIT,  SUBIT).
Section  3-^:   Description  of  Valid  Parameters  used  in  Action  Instructions,
CLASS,  PDN,  UDN,  PDSTP,  UDSTP,  defined  constants,  ALLOCATEd
variables,  Symbol-table  entries.  Block-table  entries  and  TEMP
variables .

3.1.3  Source  Text  Preparation  Rules

The  rules  for  preparing  the  syntax  source  specification  for  the
assembler  program  are  as  follows:

13

1)  Names:

All  def Ined-names ,  variable-names,  PROC-names  and  label-names

discussed  in  this  paper  are  of  the  form:

Any  ccambination  of  9  letters  and  number,  with  the  first  character
being  a  letter  (note  that  capital  letters  are  acceptable,
but  they  require  2  character  positions  in  the  name),

2)  Form  of  instructions:

a)  All  instructions  in  section  3.2  are  of  the  form:

<instruction>      eol
that  is,  one  instruction  per  line.   The  eol  is  inserted
automatically  by  the  editor  program  when  a  line  is  terminated.
No  explicit  spacing  is  required  within  a  line  (free  form,
extra  blanks  ignored),

b)  All  instructions  in  section  3.3  (Action  Instructions)  are
of  the  form:

<label>    <instruction>    eol
where  the  <label>  is  optional  in  all  cases  (except  on
instructions  following  a  RETI,  RET,  GOTO,  or  unconditional
branch  (BC  TRUE)  instruction).   The  <label>,  if  present,
marks  the  table-location  for  that  instruction:  control  may
then  be  passed  to  these  <labe]>s  via  a  BC,  GOTO,  or  the
multiple-return  form  of  the  CALLI  and  CALL  instructions.   It
is  suggested  for  readability  (not  required)  that  all
<Label>s  begin  at  the  left  margin  and  all  instructions  begin
at  the  normal  tab  position.

lU
c)  The  following  "brackets"  are  used  in  dociimenting  the  valid
forms  of  instructions:
[,..]  :  The  specifications  inside  the  [...]  may  appear

0  or  1  times.
{...}  :  The  specifications  inside  the  (...)  may  appear  any

number  of  times.  (0,1,2,  ,..).
{.,.]  :  The  specifications  may  appear  at  most  "n"  times
(0,1,2,  ...,  n).
3)   Throughout  the  remainder  of  this  chapter  all  instruction
"^  keywords  that  are  discussed  will  be  CAPITALIZED  for  emphasis.

However,  as  shown  in  the  various  examples  given,  these  key-
words are  accepted  in  lower  case  only  by  the  assembler  program.
k)     An  asterisk  appearing  anywhere  on  a  line  causes  the  rest  of
the  line  to  be  treated  as  a  comm^ent.   Comments  may  be  used
liberally.

3.2  Overall  Organization  of  the  Syntax  Specification

3.2.1  Syntax  Source  Text  Organization

The  normal  form  for  the  syntax  source  text  is  to  have  the  main
procedure  text  first,  followed  by  a  sequence  of  PROC  -  END  blocks.

The  main  procedure  text  contains  all  DEFINE  instructions,  follow-
ed by  all  global  variable  ALLOCATE  instructions,  followed  by  all  error  system
instructions  (CLASS  NAME,  IlASK  NAME,  FIELDPDN,  and  ERROR  MESSAGE),  followed  by
the  main  parsing  procedure  (generally  containing  references  to  some  PROCs).

15

The  final  FROC  -  END  block  is  followed  by  END  SYNA,  which
signals  the  physical  end  of  the  Syntax  Source  Representation.
SUMMARY;

[DEFINE  instructions]

{Global  ALLOCATE  instructions]

(CLASS  NAME  instructions]

(MASK  NAME   instructions]

(ERROR  MESSAGE  instructions]

(FIELDFDN  instructions]

• « main  parsing  procedure

(PROC  -  END  blocks]

END  SYM

3.2.2  PROC  -  END  blocks

The  purpose  of  PROC  -  EKD  blocks  is  to  make  transition  diagram
submachines  well  defined  constructs.
FORM:

PROC     <procname>     [(<ar^     f,<arg>]^)]     [RETURN  <#rets>]

[NAME     (<error  print  name>) ]  eol

(local  variable  ALLOCATE  instructions]

proc  instructions  J  including  at  least  1  RET  instruction

END     PROC        eol
ACTION:
<procname>  is  the  procedure  name.

The  arg\iment  list  is  optional however,  a  PROC  must  be

consistently  specified.  The  maximum  number  of  arguments  for  a  procedure
is  5.  The  <arg>s  (if  present)  are  considered  to  be  local  va.riables  to  the

16
procediare  (call  by  value,  no  result  returned).   They  do  not  need  to  be
ALLOCATEd  (the  assembler  will  automatically  allocate  space  for  any  procedure
<arg>s),  however  they  may  be  included  in  an  ALLOCATE  statement  inside  the
PROC  if  desired  (the  assembler  accepts  either  implicit  or  explicit  alloca-
tion in  this  case)  (see  section  3*2.6  for  more  information  about  local
variable  ALLOCATE  instructions).

The  MULTIPLE  RETURN  option  is  available  to  allow  the  parser  to
return  to  more  than  one  state  in  the  Syntax  Specification  after  the  procedure
has  been  executed  (see  section  3«3.3-  3.3.6  for  more  information  on
CALLing  and  RETurning  from  procedures).   The  MULTIPLE  RETURN  option  must  be
specified  only  if  the  PROC  uses  multiple  returns.   If  the  PROC  uses  multiple
returns,  the  <#rets>  parameter  (which  must  be  a  numeric  constant)  specifies
the  number  of  locations  the  procedure  may  return  to  (the  maximum  number  is

31).

The  NAME  option  is  for  use  by  the  compiler's  autmatic  sjmtax
error  system.   The  <error  print  name>  chosen  for  a  PROC  should  be  a
short  logical  name  that  describes  the  function  of  the  PROC  (examples:
"statement",  "expression",  "declaration  list",  "array  bound",  etc.)  (see
section  3.2.7  for  more  information  about  <error  print  names>).
EXAMPLE:

proc    expr  (type)  return  2  name  (expression)    eol

end   proc     eol

17

3.2.3  ENTRY  Instruction
FORM:

ENTRY    <Procnajne>  rN/VME(<error  print  namo)]    eol
ACTION:

Defines  an  alternate  entry  point  into  the  containing  PROC.
The  <procname>  is  treated  as  a  regular  procedure  name  in  CALL  instructions.
However,  all  of  the  attributes  of  the  ENTRY  <procname>  are  the  same  as  those
of  the  outer  PROC:

number  and  order  of  parameter  arguments;

niimber  of  multiple  return  points;

ni:unber  of  local  variables.

Furthermore,  none  of  these  attributes  can  be  specified  at  the  ENTRY  instruc-
tion    only  at  the  containing  PROC.  The  only  unique  attribute  of  an

ENTRY  <procname>  is  the  (possibly)  unique  <error  print  name>,
EXAMPLE:

entry  operand  name  (operand)       eol

3o2.i+  END  SYNA  Instruction
FORM:

END   SYNA       eol
ACTION:

Signals  the  physical  end  of  the  syntax  source  text  to  the
assembler.

18

3.2.5  DEFINE  Instruction
FORM:

DEFINE   <name>  =  <constant>  {,  <3!ianie>  =  <constant>}   eol
ACTION:

This  instruction  is  used  to  define  mnemonic  constants  for  the
assembler  to  use.   No  table  code  is  actually  generated  for  this  instruc-
tion.

Note  that  <constant>  can  be  either  a  numeric  constant  or  a
previously  defined  mnemonic  constant,
EXAMPLE:

define   ifx  =  0,  colonx  =  lOU,  dclvar  =  10      eol
ALL  DEFINE  instructions  must  preceed  all  other  statements  of  the  syntax
source  specification,

ALTERMATE  FORM:

Instead  of  <constant>  above,

MASK  (<12-bit  mask  of  O's  and  1 ' s>)
can  be  used.   There  must  be  exactly  12  bits  specified  between  the
parentheses.

This  is  a  convenient  way  to  define  a  particular  mask  bit

pattern.

EXAMPLE:

define   n-jxaeric  =  mask  ( 000001011111 )       eol
It  is  also  legal  to  combine  the  two  forms  of  the  DEFINE  instruction  on  the
same  line.

19

EXAMPLE:

define   ifx  =  0,  numeric  =  mask  ( 000001011111 ) ,  dclvar  =  10   eol

3.2.6  ALLOCATE  Instructions
FORM:

ALLOCATE   <variable>   (,  <variable>}        eol
ACTION:

Causes  variable  storage  to  be  allocated  on  a  stack  in  the
compiler.  These  storage  locations  are  referenced  by  using  the  name
<variable>.

Note  that  the  allocated  storage  will  be  a  GLOBAL  allocation  if
the  ALLOCATE  instruction  comes  at  the  beginning  of  the  syntax  program
(i.e.,  before  the  first  PROC  -  END  block  definition)  and  a  LOCAL  allocation
(implying  possibly  recursive  allocation)  if  the  ALLOCATE  instruction  is
within  a  PROC  -  END  construct.

Global  variables  can  be  referenced  from  anywhere  within  the
syntax  specification,  whereas  local  variables  can  be  referenced  only
from  within  the  PROC  -  END  block  in  which  they  were  allocated.

3.2.7  Error  System  Instructions
FORMS:

CLASS  NAME  <c lass -numb er>  (<error  print  name>)        eol

MASK  NAME  <mask-pattem>(<error  print  name>)"        eol

ERROR  MESSAGE  <error-number>  (<overide  error  message  text>)   eol

FIELDPDN    <pdn-number>  {,<pdn-number>}       eol

20
ACTION:

The  purpose  of  the  CLASS  NAME  and  mSK  NAME  instructions,  as
well  as  the  procedure  NAME  instruction  (see  documentation  on  PROC  -  END
blocks)  is  to  provide  the  compiler's  automatic  error  analysis  system
with  text  to  refer  to  CLASSes,  MASKs  and  specific  errors  that  appear  in
the  parser  environment.

The  automatic  error  analysis  system  interacts  with  the  user  by
suggesting  different  modifications  that  can  be  made  to  correct  an  error
in  the  program.   Many  of  these  suggestions  need  to  be  made  in  the  termin-
ology of  the  programming  language  involved;  these  NAME  instructions
provide  the  correlation  between  things  in  the  parser  environment  and  the
language  terminology  that  describes  these  things.

Typical  messages  using  these  NAMES  are:
"Replace  |    |  with  a  relational  operator."  Class  Najne

"Insert  an  array  bound  in  front  of  I    1  ."  Proc  Name

"Replace  I    I  with  a  declared  variable  (numeric) ."     Class,  Mask  Names
Each  Mask  pattern  that  is  used  in  the  mask  form  of  the  conditional  branch
instruction  (see  section  3.3.7.3)  and  each  CLASS  number  should  be  given  an
associated  error  name.

The  ERROR  MESSAGE  instruction  is  used  to  override  the  operation
and  analysis  of  the  automatic  error  system.   If  the  error  number  signalled
by  the  parser  (via  a  conditional  branch  (be)  instruction)  matches  the
<error-number>  given  in  an  ERROR  MESSAGE  instruction,  then  the  overide
text  is  displayed  to  the  user  and  all  subsequent  error  analysis  processing
is  aborted.

21
Finally,  the  FIELDPDN  instruction  is  used  to  inform  the  error
system  a.bout  which  pdn  numbers  (section  3.^.3.3)  represent  "field  tokens";
each  field  pdn  number  should  be  included  in  a  FIELDPDN  instruction.
EXAMPLES:

class  name    punct  (punctuation)      eol

mask  name     numeric  (numeric  variable)       eol

fieldpdn      label,  stmt      eol

3.2.8  FimL  PARSE  STATE  Instruction
FORM:

FINAL  PARSE  STATE       eol
ACTION:

Before  the  "execution"  of  a  compiled  user  program  can  be
attempted,  there  must  be  some  way  for  the  compiler  supervisor  system
to  verify  that  the  program  is  indeed  "complete"  (since  in  the  interactive
compiling  environment  a  user  could  request  execution  of  an  unfinished
program).

This  instruction  indicates  to  the  parser  that  it  is  in  the
program- accepting  state;  that  is,  the  program  can  be  executed  if  and
only  if  the  parser  is  in  this  state.

There  must  be  exactly  one  accepting  state  specified  in  the
Syntax  Specification  (i.e.,  the  FINAL  PARSE  STATE  instruction  must  occur
exactly  once).

22
3- 3  Parser  Action  Instruction  Description

This  section  describes  the  form  and  use  of  the  assembler  instruc-
tions that  allow  a  language  designer /implement or  to  fully  specify  the
syntactic  and  semantic  requirements  of  the  language  being  implemented.
These  instructions  constitute  an  implementation  of  the  augmented  parser
transition  diagram  model  discussed  in  Chapter  2  of  this  paper.   The
instructions  allow  invoking  and  returning  from  "subma chines"  (PROCs  in
this  assembler  language),  examining  the  current  input  token  (from  the
lexical  analyzer)  for  validity,  both  syntactically  and  semantically
(i.e.,  context-sensitive  requirements),  accepting  the  current  token  and
requesting  that  a  new  token  be  input  from  the  lexical  analyzer,  and
finally,  a  few  instructions  allow  changes  to  be  made  to  the  parser
environment  (i.e.,  symbol-table  modifications  or  parser  ALLOCATEd  variable
modifications).

One  further  note:   many  of  the  instructions  to  be  described
refer  to  general  "parameters",  which  are  the  operand  (s)  used  by  the
instructions.   These  parameters  will  be  referred  to  as  <parra>,  or  <parml>
and  <parm2>  in  the  form  of  the  instructions.   In  all  cases,  these  <parm>s
will  resolve  to  a  memory  location  in  the  parser  environment  (like  a  symbol-
table  reference).   These  <parra>s  are  discussed  in  detail  in  section
3.^  of  this  paper.

The  remainder  of  this  section  will  discuss  the  Action  Instructions:
SCAN,  GOTO,  CALLI,  CALL,  RETI,  RET,  BC,  SEMA,  and  the  aiixiliary  environment-
changing  instructions  (ASSIGN,  MASKON,  MASKOFF,  ADDIT,  SUBIT).

23

3.3.1  SCAJJ  Instruction
FORM:

SCAN         eol
ACTION:

Causes  a  return  to  LEXI  for  another  token.  When  the  next
token  comes  in,  the  parser  resumes  parsing  at  the  State  following  the
SCAN  instruction  (i.e.,  at  the  table-location  following  the  SCAN
instruction's  table-location),

3.3.2  GOTO  Instruction
FORM:

GOTO  <label>    eol
ACTION:

Causes  a  SCAN  instruction  to  be  executed,  followed  by  an
unconditional  branch  to  the  instruction  corresponding  to  <label>  in  the
table.
EXAMPLE:

goto  stmntl       eol

3.3.3  CALLI  Instruction
FORM:

CALLI  <procname>  [(<parm>  {,  <parm>})]  [,  THEN<Label>  {,<Label>}]  eol
ACTION:

This  instruction  is  used  to  invoke  a  EROC  with  the  name  <procname>.

(see  section  3.2.2  and  3.2.3.)

2k

Causes  a  return  address  (table-location)  to  be  saved  on  the
parser  stack,  passes  the  argument  values  to  the  proper  local  variables
in  <procname>j  and  resumes  parsing  at  <procname>  table-location.

When  a  return  instruction  (RETT,  RET)  is  executed,  the
proper  return  label  location  is  selected,  the  parser  stack  is  popped,
and  parsing  is  resumed  at  the  new  location.   If  the  PROC  multiple-return
option  is  not  used,  parsing  is  resumed  at  the  location  of  the  instruction
following  the  CALLI  (  or  CALL)  instruction.

Note  that  the  parameter  argument  values  are  passed  into  the
PROC  only,  and  that  the  final  values  are  not  passed  back  to  the  parajneter
argument  upon  returning  from  the  PROC  (i.e.,  call  -  by  -  value  only).

Note  also  that  all  instances  of  the  multiple  return  option  for
a  PROC  must  be  consistently  specified  (both  in  the  CALLI  (CALL)  instruc-
tions and  in  the  PROC  -  END  definition).
EXAMPLE:

calli   var  (m,n),  then  labl,  lab2     eol

3.3.^  CALL  Instruction
FORM;

Same  as  CALLI  instruction

ACTION:

Causes  a  SCAN  instruction  to  be  executed,  follcswed  by  a  CALLI
instruction.

EXAMPLE;

call  subscr        eol

25

3.3.5  RETT  Instruction
FORM;

RETI  [<retum-nuinber>]      eol
ACTION;

Causes  a  return  from  a  previously  invoked  PROC  (see  section

3.3.3).

If  the  PROC  has  any  locally-ALLOCATEd  variables,  the  space  is
deallocated  from  the  variable-storage  stack  in  the  parser  environment.
Then  the  return  address  (table  location)  is  popped  off  of  the  regular
parser  stack.

If  the  multiple-return  option  is  not  used  by  this  FROC,  parsing
resumes  at  the  popped  return  address  location.

If  the  multiple -re turn  option  is  used  by  the  PROC,  the
"re turn -number "th  <label>  given  in  the  original  CALLI  (CALL)  instruction
is  selected  and  parsing  resumes  at  the  table -location  corresponding  to
this  selected  <label>.

Note  that  <return-number>  must  be  a  constant  (or  a  mnemonic
defined  constant)  in  the  assembler.
EXAMPLE:

reti   2    eol

3.3.6  RET  Instruction
FORM;

Same  as  RETI  Instruction

26

ACTION;

Causes  a  SCAIT  instruction  to  be  executed,  followed  by  a  RETI
instruction.
EXAMPLE:

ret      eol

3.3.7  BC  Instruction

f  3.3.7.1  Normal  BC  Instruction

FORM;

BC  <relation-type>,  <parml>,  <parm2>,  <true-option>      eol
where

<relation-type>  : ;=BQ  J  NE  |  GT  |  GE  |  LT  |  LE
<true-option>   ;  :=<label>  |  ERROR  [<parm>]
ACTION:

Causes  the  2  <parm>s  to  be  compared  according  to  <relation-
type>.

If  the  comparison  is  false,  the  <true-option>  is  ignored  and
control  falls  through  to  the  next  instruction  in  the  table.

If  the  comparison  is  true,  the  <true-option>  is  taken;
If  <true-option>  is  a  <Label>,  then  parsing  resumes  immediately

at  the  table -location  corresponding  to  that  label.
Otherwise,  <true-option>  is  a  syntax  error  indicator;  this  causes
the  parser  to  halt  its  operation  and  compiler  control  is  pass-
ed to  the  compiler's  error  analysis  system,  with  an  Error  Number
equal  to  the  value  of  the  <parm>  (if  specified).

27

In  the  compiler's  automatic  error  analyzer,  the  Error
Numbers  are  ignored  unless  an  ERROR  MESSAGE  instruction  for
the  particular  Error  Number  has  been  included  with  the  Syntax
Specification  (see  section  3.2.7  of  this  paper).

In  a  hand-coded  compiler  error  system,  the  Error  Number
can  be  used  to  display  a  unique  error  message  to  the  user.
The  BC  instruction  is  the  only  instruction  that  is  available  to  signal
that  a  syntactic/ semantic  error  has  occurred.

Note  that  <parml>  cannot  be  a  constant  or  mnemonic  defined
constant.
EXAMPLES :

be  ne,  pdn,  colonx,  not lab  eol
be  eq,  class,  dclvar,  assign  eol
be   ne,  pdn,  thenx,  error  10        eol

3.3.7.2  Unconditional  Branch  Instruction

FORM:

BC   TRUE,  <true-option>      eol
ACTION:

Causes  the  <true-option>  to  be  executed  exactly  as  if  the
instruction  was  a  normal  BC  instruction  whose  parameter  comparison
was  TRUE.
EXAMPLE:

be  true,  looplab  eol

28
3.3.7.3  Attribute-Checking  BC  Instruction
FORM:

Same  as  the  normal  BC  instruction,  except  that

<relation-type>  ;  :=MA.SK,  <mask-type>

<mask-type>      :  :=NOTANY  |  NOTALL  |  AM   \   ALL
ACTION:

Most  context-sensitive  language  requirements  can  be  viewed  as
"attributes"  of  the  particular  tokens  (both  pre-defined  and  user-defined
tokens)  used  in  a  user's  program.  In  this  compiler  system,  each  symbol-
table  field  contains  12  bits;  although  some  of  the  symbol-table  fields
have  very  specific  builtin  uses  (i.e.,  PDN  for  pre-defined  symbols,  and
CLASS  for  both  pre-defined  and  user-defined  tokens),  some  of  the  remaining
fields  have  no  builtin  use  (for  example,  the  UDN  field  for  user-defined
tokens).  It  is  suggested  that  the  language  designer  select  an  unused
symbol-table  field  (such  as  UDN)  and  let  each  of  the  12  bits  in  the  field
represent  a  different  attribute  that  a  user-defined  token  may  have.  Specific
attribute  bits  may  be  turned  on  or  off  using  the  MASKON  and  MASKOFF
instructions  (see  section  3.3.9).  The  existence  of  an  attribute  for  a  token
can  then  be  checked  using  the  MASK  form  of  the  BC  instruction.

The  2  parameters  are  compared  according  to  the  specified
<mask-type>.   For  example,  if  <jnask-type>  is  NOTANY,  then  the  comparison
is  TRUE  if  NOTANY  of  the  bits  that  are  1  in  <parm2>  are  also  1  in  <parml>,
and  FAISE  otherwise.

29

Note  that  it  is  possible  to  check  for  1  attribute  bit  being
on  or  off,  or  any  combination  of  attribute  bits  being  on  or  off.  This
allows  for  example,  the  grouping  of  two  attributes  like  "FIXED  and  "FLOAT"
together  for  certain  types  of  tests  (like  the  attempted  declaration  of  an
attribute  "CHARACTER");  it  may  not  be  important  which  of  the  grouped
attributes  conflicts,  but  simply  that  a  conflict  exists.

The  compiler's  automatic  error  analysis  system  views  the  MASK
form  of  the  BC  instruction  as  specifying  an  attribute  check  of  the  ciarrent
token  being  examined  by  the  parser.   It  is  possible  to  give  each  particular
bit-mask  that  is  used  as  a  <parm2>  in  a  BC  MASK  instruction  a  unique  error
print  name  (see  section  3.2.7);  this  print  name  will  then  be  used  in  any
generated  diagnostic  messages  that  involve  the  bit  mask.

If  a  hand-coded  error  system  is  used,  appropriate  unique  Error
Numbers  must  be  used  as  in  a  normal  BC  instruction.
EXAMPLES:

Assume  that  at  some  point  in  the  syntax  specification,  the  only
valid  syntax  option  is  a  numeric  declared  variable.   Then  if

dclvar   : :  =  a  constant  whose  value  is  the  CLASS  for  a  user  declared

variable,  and
num     : :  =  a  constant  whose  value  is  the  bit-mask  for  the  numeric
attribute (s ) ,

30
the  following  instructions  perform  the  required  check:

be  ne,  class,  dclvar,  error  1  Mclvar  required  here...

be  mask,  notany,  udn,  num,  error  2  ■'^numeric-type  required.

(
As  another  example,  assume  that  a  new  user  identifier  is  being  declared.

Let

varattrib   : :  =  accumulated  attribute  bit-mask  for  new  identifier,  and
conflict   : :  =  the  conflicting  attribute  bit-mask  for  a  new  attribute
the  user  is  trying  to  add  for  the  identifier.

The  following  instruction  check  the  validity  of  the  new  attribute:

be  mask,  any,  varattrib,  conflict,  error  3  "'^attributes  conflict,

31

3-3-B  SEMA  Instruction
FORM:

SEMA  <seina-nuin'ber>   [(<pann>  {,   <pann>})]       eol
ACTION:

Occasionally  certain  unique,  non-standard  operations  must  be
performed  by  the  parser.   The  SEMA  instruction  allows  a  language  implementor
to  write  a  regular  TUTOR  unit  to  perform  these  operations,  and  then  have  the
parser  execute  these  TUTOR  units  at  appropriate  times.

<sema-number>  must  be  a  constant  or  a  mnemonic  defined  constant,
between  1  and  15-   The  language  implementor  supplies  TUTOR  units  named
SMI  to  Sml5  for  the  compiler.   Upon  execution  of  the  SEMA  instruction  in  the
syntax  table,  the  parser  will  do  the  correspondingly  numbered  TUTOR  unit.
After  the  TUTOR  unit  is  finished,  the  parser  resumes  parsing  with  the  next
instruction  in  the  table.

Up  to  5  parameters  may  be  passed  to  semantic  routines  (note  that
the  assembler  performs  no  check  for  inconsistent  number  of  arguments  for
different  uses  of  a  particular  semantic  routine  number).

To  use  the  parameters  in  the  TUTOR  unit:

The  parameters  are  passed  as  addresses  into  the  parser  storage

environment  through  the  use  of  5  specially  located  variables  in  the  parser

storage:   they  are  located  at  parser  storage  locations  (ps_prm  +  i),  where

i  is  the  parameter's  number  (1-5)  and  ps_j)rm  is  a  compiler  system  defined

constant.

Therefore,  to  use  the  address  of  the  argument  passed

through  parameter  i,  reference  :  ps  (ps_prm  +  i).   This  address  is

some  location  in  the  parser  storage  environment  (for  example,  a

particular  symbol-table  field  for  the  current  token  in  the  parser).

32

To  use  the  value  of  the  argument  passed  through  parameter  i,  refer-
ence: parm(i)  which  is  defined  in  the  compiler  as  ps(ps(ps_prm  +  i)).
The  only  exception  to  the  above  is  that  constants  (or  mnemonic
.  defined  constants)  passed  as  parameters  are  passed  as  just  the
value  of  the  constant  (i.e.,  reference  this  value  as  ps(ps_prm  +  i)).
It  is  up  to  the  language  implementor  to  know  and  keep  track  of  which
parameters  in  a  semantic  routine  are  passed  as  constants  and  which
are  passed  as  addresses,  and  to  specify  the  parameters  consistently
[   for  each  routine.
EXAMPLES:

sema  getdopev  (udstp,  s t_dvl ( uds tp ) )         eol
sema  opendblk         eol
sema  k  eol

RESTRICTIONS  ON  THE  USE  OF  SEMANTIC  ROUTINES:

There  are  a  few  restrictions  that  must  be  placed  on  the  use  of  semantic

routines:

(l)   Tracing  of  any  changes  made  to  parameters  passed  as  addresses:
In  order  to  allow  the  compiler  to  properly  "backup"  if  the  user
edits  the  program  being  written  (since  this  is  an  interactive  compil-
ing environment),  anytime  that  the  value  of  an  address  in  the  parser
storage  environment  is  changed,  the  old  value  must  first  be  traced
using  the  compiler  unit  TRACE,  which  has  as  its  one  argument  the
address  to  be  changed.   For  example,  if  semantic  routine  parameter
number  2  is  passed  as  an  address,  and  the  semantic  routine  decides
to  change  the  value  at  that  address,  the  following  TUTOR  code  is  needed:

33

do   trace  (ps(ps_prm  +2))   $$  trace  old  value
calc  parm  (2)  ♦-  'the  new  value'

•  •  •
(2)  Handling  of  syntactic/semantic  errors  detected  within  a  semantic
routine :

If  an  error  is  detected  in  a  semantic  routine,  it  is  not  permitted
for  the  semantic  routine  TUTOR  unit  itself  to  execute  a  transfer
of  control  from  the  parser  to  the  error  analysis  system.   Instead,
all  exits  to  the  error  system  must  come  through  having  the  parser
execute  a  BC  instruction  that  has  a  <true-option>  of  the  EEIROR  form
(see  section  3 .3. 7  of  this  paper).

The  easiest  way  to  do  this  for  an  error  that  is  detected  with-
in a  TUTOR  semantic  unit  (note  that  this  type  of  error  checking  in
a  TUTOR  semantic  unit  is  very  non-standard--nearly  all  detectable
syntactic/semantic  errors  can  be  detected  through  the  use  of  appropri-
ate BC  testing  instructions)  is  to  use  a  temporary  variable  TEMPI
(see  section  3.^.3.5  of  this  paper)  that  is  returned  from  the  TUTOR
unit  as  either  0  (everything  is  ok),  or  non-zero  (error  detected —
the  non-zero  value  can  be  an  appropriate  Error  Number),  and  then  the
TEMP  variable  can  be  checked  in  the  instruction  following  the  SEMA
instruction:

•  •  •

sema     some -number  (some -parameters)         eol

be  ne,  tempi,  0,  error  tempi     Ogives  proper  error  number

•^for  hand-coded  error  system.

31+

3)   Variables  that  may  be  referenced  within  a  TUTOR  semantic  routine:

Any  parser  storage  address  that  must  be  referenced  within  a  TUTOR
semantic  routine  unit  must  be  passed  through  the  parameter  list.
Although  this  is  the  only  way  to  reference  an  ALLOCATED  variable
in  the  Syntax  Specification,  this  restriction  also  includes  the
special  parser  storage  locations  like  PDN,  CLASS,  UDN,  PDSTP  and
UDSTP,  even  though  these  locations  are  also  defined  directly
within  the  compiler  system  itself  (see  section  3.^.3  of  this  paper
for  a  description  of  these  special  locations).

This  restriction  is  imposed  by  the  compiler's  automatic
error  analysis  system.
k)       Modifying  a  special  parser  storage  location  (section  3.^.3)  in  a
semantic  routine  requires  some  care  by  the  language  implementor.
Since  the  special  locations  are  actually  just  duplicate,  easily
referencable  copies  of  some  of  the  fields  in  the  symbol  table
entry  for  a  token,  any  changes  to  either  of  the  two  corresponding
locations  should  be  accompanied  by  a  change  to  the  other  location
also.   Note  that  only  the  symbol-table  fields'  values  need  to  be
traced,  and  not  the  special  locations.

3.3.9  ASSIGN,  MASKON,  MASKOFF,  ADDIT,  SUBIT  Instructions
FORMS:

ASSIGN  <parml>,  <parm2>  eol

MASKON  <parml>,  <parm2>  eol

MASKOFF  <parml>,  <parm2>  eol

ADDIT  <parml>,  <parm2>  eol

SUBIT  <parml>,  <parm2>  eol

35
ACTION:

Each  of  these  instructions  is  used  to  change  the  value  of  the
parser  storage  location  for  <p£Lrml>  to  <parm2>,  or  some  function  of
<parnil>  and  <parm2>.

The  <parm>s  are  described  in  section  3.^  of  this  thesis.
ASSIGN:

Sets  the  value  of  <parml>  to  the  value  of  <parm2>.
MASKON,  MASKOFF:

Sets  (on,  off)  all  bits  in  <parml>  corresponding  to  I's  in
<parm2>;  does  not  change  bits  in  <parml>  corresponding  to  O's  in  <parm2>.
ADDIT,  SUBIT:

(Adds,  subtracts)  the  value  of  <parm2>  (to,  from)  the  value  of
<parml>.

Each  of  these  instructions  will  trace  the  old  value  of  <parml>
before  executing  the  instruction  (except  for  the  TEMP  variable,  and  also
PDN,  CLASS,  UDN,  UDSTP,  PDSTP  variables,  which  do  not  need  to  be  traced
(see  sections  3- ^.3.1  -  3«^.3.^  for  more  details)).
IMPORTANT  NOTE;

If  <parml>  is  CLASS,  UDN,  PDN,  then  both  the  special  parser
storage  location  and  the  corresponding  symbol-table  field  are  modified
(see  section  3.^.3.2  -  3.^.3.^  for  more  details).
EXAMPLES:

assign  class,  dclvar      eol
addit  numsubs,  1  eol

36

3.^  Description  of  Valid  Parauneters  Used  in  Action  Instructions
Most  of  the  Parser  Action  Instructions  have  one  or  more
<parm>s,  that  is  they  have  "parameters"  or  "operands"  associated  with
them  (specifically,  BC,  CALLI,  CALL,  SEMA,  ASSIGN,  MASKON,  MASKOFF,  ADDIT
SUBIT).  This  section  documents  the  form  and  uses  of  the  different
<parm>s  that  are  available  in  the  paxser  environment.

3.^.1  Numeric  Constants

Numeric  (integer)  constants  or  mnemonic  DEFINEd  constants  are
legal  parameters  (except  as  <parml>  of  a  BC  (see  section  3.3.7)  or  an
auxiliary  environment -changing  (section  3.3.9)  instruction).   The  constant's
value  is  packed  directly  into  the  parser  table.   The  only  illegal  constant
value  is  ^095  (octal  o7777)j  which  is  reserved  for  use  by  the  compiler's
automatic  error  analysis  system.

3.)4.2  ALLOCATED  (Global  and  Local)  Variables

GLOBAL:  Global  variables  are  legal  parameters  anyplace  in  the

Syntax  Specification.   The  parser  storage  direct  address
of  the  Global  variable  (known  at  assembly  time)  is  packed
into  the  parser  table.

ALL  GLOBAL  variables  must  be  ALLOCATEd  at  the  beginning
of  the  Syntax  Specification  (directly  following  any

DEFINE  instructions) (see  section  3.2.6).

LOCAL:   Local  variables  are  legal  parameters  any  place  within
the  FROC  -  END  block  in  which  they  are  ALLOCATEd.  An

37

indexed  parser  storage  address  (relative  to  the

parser's  variable  storage  stack)  is  packed  into

the  parser  table.   LOCAL  variables  must  be  ALLOCATED

at  the  beginning  of  the  particular  PROC  -  END  block

in  which  they  are  active;  note  that  a  PROC's  parameters

are  implicity  ALLOCATEd  Local  variables  (see  section

3.2.6).

3.^.3  Pre-defined  Parser  Variables

There  are  a  number  of  pre-defined  parser  storage  variables  that
can  be  tested  or  otherwise  used  as  legal  pareuneters  anywhere  in  the  Syntax
Specification.  In  all  cases,  the  parser  storage  direct  address  (known
prior  to  and  during  assembly  time)  is  packed  into  the  parser  table.

3.^.3.1  PDSTP,  UDSTP:   Pre-defined  and  User-defined  Symbol
Table  Pointers

In  this  compiler  system,  the  symbol-table  is  logically  divided
into  2  parts:  one  part  consists  of  all  of  the  pre-defined  tokens  that
belong  to  the  language  being  implemented  (i.e.,  the  reserved(or  unreserved)
keywords,  punctuation  symbols,  operators,  etc.);  the  second  part  of  the
symbol  table  consists  of  any  tokens  that  a  user  may  have  used  in  writing
a  program  and  that  either  do  not  have  a  corresponding  pre-defined  entry,
or  else  are  used  in  a  context  that  is  different  from  that  of  the  correspond-
ing pre-defined  entry,  (for  example,  a  declared  variable  "IF"  in  PL/1).

38

The  PDSTP  and  UDSTP  variables  always  contain  the  values  of  the
appropriate  symbol-table  pointers  for  the  current  token  that  the  parser  has
received  from  the  lexical  analyzer.  If  the  current  token  has  a  pre-defined
and  a  user-defined  symbol  table  entry,  then  PDSTP  and  UDSTP  point  to  these
entries,  respectively.  If  the  current  token  has  only  a  pre-defined  entry,
then  both  PDSTP  and  UDSTP  point  to  this  pre-defined  entry.  If  the  current
token  has  only  a  user-defined  symbol  entry,  then  UDSTP  points  to  this  entry
and  PDSTP  is  essentially  null  (it  points  to  an  empty  pre-defined  symbol-
table  entry  that  no  token  can  ever  resolve  to).

Note  that  the  only  way  for  a  token  to  have  both  a  pre-defined
and  a  user-defined  symbol-table  entry  is  for  the  language  implementor  to
specifically  provide  an  appropriate  SEMAntic  routine  (see  section  3.3.8)
that  actually  creates  the  user-defined  entry  from  the  pre-defined  entry;
the  compiler's  symbol-table  manager  will  not  create  a  user-defined  entry
automatically  for  a  token  that  resolves  to  a  pre-defined  location.

See  section  3.^.^.1  for  a  description  of  how  to  reference
particular  symbol-table  fields,  given  UDSTP  or  PDSTP.

3.^.3.2  CLASS:   The  syntactic/semantic  class  of  the
current  token

The  CLASS  variable  is  defined  as  ST_TYP  (UDSTP)  (see  section

3.^.4.1).  For  pre-defined  tokens,  CLASSes  will  usually  be  things  like

"statement  keywords",  "relational  operators",  "attribute  keywords",  etc.

For  user-defined  tokens,  CLASSes  will  be  things  like  "declared  variable",

"labels",  "undeclared  variable",  etc.   These  class  values  are  so  commonly

referred  to  in  parsing  a  language  that  the  parser  maintains  a  special

39

parser  storeige  location  that  contains  the  CIASS  of  the  current  token  that
is  being  examined.

If  the  current  token  has  only  a  pre-defined  symbol-table  entry,
then  CIASS  is  set  to  the  ST_TYP  field  of  this  pre-defined  entry.   If
the  current  token  has  a  user-defined  symbol-table  entry  (or  both  a  pre-
defined and  a  user-defined  entry),  then  CLASS  is  set  to  the  ST_TYP  field
of  this  user-defined  entry.

Note  that  the  CLASS  of  a  user-defined  token  may  be  changed
and  updated  through  the  use  of  any  of  the  auxiliary  instructions  (see
section  3.3.9) •   If  the  lexical  analyzer  accumulates  a  "new"  token,
that  is,  a  token  that  has  no  symbol-table  entry,  then  the  symbol-table
manager  will  automatically  create  a  new  user-defined  symbol-table  entry
for  the  new  token,  and  the  CLASS  (ST_TYP)  field  of  the  new  entry  will  be
set  to  the  default  CLASS  value  specified  by  the  lexical  analyzer.

Note  that  all  CLASS  values  used  in  a  particular  language
implementation  should  be  given  CLASS  NAMEs  for  the  compiler's  automatic
error  analysis  system  (see  section  3.2,7).

3.i+.3.3  FDN:   Pre-defined  Number
PDN  is  a  unique  identification  number  for  a  pre-defined  token.
The  FDN  variable  is  defined  as  ST_IDN  (PDSTP)  (see  section  3.U.i+.l  for  a
description  of  symbol-table  fields).  Each  token  that  is  entered  in  the
pre-defined  symbol-table  by  a  language  implementor  should  be  given  an
identification  number  that  can  be  checked  by  the  parser  (using  BC  instruct
tion,  section  3.3.7)  when  a  particular  pre-defined  token  is  a  valid  parse
option  for  the  current  parser  State  and  parser  environment.   The  same

ko

identification  number  may  be  given  to  more  than  one  pre-defined  token  to
allow  for  pre-defined  synonyms  or  abbreviations  (such  as  "DECLAEE"  and
"DCL"  in  PL/I).

Note  that  EDN  is  always  based  on  PDSTP,  so  if  no  pre-defined
symbol-table  entry  exists  for  the  current  token  being  examined,  the
value  of  EDN  is  null,  that  is,  it  will  match  nothing.
*

3.^.3.^  UDN:  User -de fined  Number

(

UDN  is  the  value  of  ST_IDN  (UDSTP),  which  does  not  have  a  pre-

reserved  meaning  in  the  symbol-table  (that  is,  it  is  available  for  use

by  a  language  impleraentor  in  whatever  way  is  desired).

ITie  UDN  variable  (i.e.,  the  ST_IDN  (UDSTP)  symbol-table  field)

is  best  used  as  an  attribute  field  for  user-defined  tokens.   Each  bit  in

the  field  (there  are  12  of  them)  can  be  used  to  denote  a  particular

attribute  in  the  programming  language  that  a  user -defined  token  may  assume.

These  attribute  bits  can  be  turned  on  or  off  using  the  MA.SKON  and  MA.SKOFF

auxiliary  instructions;  also  the  attributes  of  a  particular  token  can  be

tested  for  consistency  using  the  BC  MASK  form  of  the  BC  instruction.  See

section  3.3.7  for  a  more  complete  discussion  of  the  handling  of  attributes

by  the  parser.

3.^.3.5  TEMP:  Used  as  temporary  computation  variables  only
There  are  5  temporary  variables  available  for  use  by  the
language  implementor:   TEMPI,  TEMP2,  TEMP3,  TEMPi+,  TEMP5.   These  variables
can  be  used  to  temporarily  save  any  value  in  the  parser  environment.   The
most  important  restriction  on  the  use  of  these  TEMP  variables,  however.

is  that  they  may  NOT  be  used  to  save  a  value  if  the  parser  retvirns  to
the  lexical  analyzer  for  a  new  token:  these  variables  are  only  valid
between  returns  to  the  lexical  analyzer  by  the  parser.

The  reason  for  this  is  that  the  value  changes  of  these
variables  are  not  traced  by  the  parser,  so  that  program  editing  perform-
ed by  the  user  will  not  properly  restore  the  values  to  the  variables.
The  variables  are  most  useful  for  returning  error  indicators  from  SEMAantic
routines .

3.i+.3.6  BLOCK:
Contains  the  current  symbol-table  block  number  that  is  being
used  by  the  symbol-table  and  the  parser.

3.4.3.7  ITPTP :   Intermediate  Text  Parser  Token  Pointer
Contains  the  current  location  of  the  intermediate  text  pointer.
This  value  is  useful  for  saving  information  about  locations  in  the  inter-
mediate text  that  will  be  needed  when  the  user  executes  the  program
(like  "label"  locations,  or  Subprogram  entry  points,  etc.).

3.4.4  Table  References

There  are  two  tables  in  the  parser  environment  that  may  be
accessed  as  general  parameters:   the  symbol-table  and  the  block-structure
table.  Assembler  references  to  these  tables  all  follow  the  same  general
form:

h2

FIELD-NAME    (<parmx>)
where

FIELD-NAME  is  the  particular  table  field  name,  and

<parmx>  is  the  table  Index  pointer.
The  following  restrictions  apply  to  <parmx>:

a)  <parmx>  may  be  ALLOCATEd  variable  (section  3.2.6);

b)  <parmx>  may  be  pre-defined  Parser  Variable  (section  3.^.3);

c)  <parmx>  may  be  a  table  reference  again,  but  then  the  para-
r

meter  for  the  new  table  reference  must  be  either  a  Global
ALLOCATEd  variable  or  a  pre-defined  Parser  Variable  only
(i.e.,  at  most  2  levels  of  indexing  are  allowed).
3.i+.U.l  Symbol  Table  Fields
Each  symbol-table'  entry  in  the  symbol-table  contains  10  different
fields;  in  addition,  a  dope-vector  symbol-table  entry  may  be  associated
with  a  regular  symbol-table  entry  through  the  use  of  the  ST_DVL  fields
in  the  regular  entry  (see  the  discussion  of  the  ST_DVL  field  below);
these  dope-vector  entries  contain  5  fields.

ST_TYP:   This  field  is  used  to  contain  the  syntactic  type  for
the  token.   For  user-defined  tokens,  that  is,  those  tokens  which  are
accumulated  by  the  lexical  analyzer  and  that  initially  have  no  symbol-
table  entry,  the  symbol-table  manager  in  the  compiler  will  assign  a
symbol-table  entry  and  will  initialize  this  ST_TYP  field  to  a  value
supplied  by  the  lexical  analyzer.

This  is  so  important  to  the  operation  of  the  parser  that  a
special  pre-defined  parser  storage  variable  has  been  provided  to  hold
the  ST_TYP  of  the  current  token  being  examined  by  the  parser.   This
special  field  is  the  CLASS  field;  see  section  3«^'3.2  of  this  paper  for

1*3

more  information  about  the  CLASS  pre-defined  location.

ST_IDN:   For  pre-defined  symbol-table  entries,  this  field
contains  a  unique  identification  number  for  the  pre-defined  token.  This
ST_IDN  number  is  so  important  that  a  special  pre-defined  parser  storage
location  has  been  provided  to  hold  the  ST_IDN  number  for  the  current  token
being  exajnined  by  the  parser.   This  special  field  is  the  PDN  field;  see
3.^.3.3  of  this  paper  for  more  information  about  the  PDN  pre-defined
location.

For  user-defined  symbol-table  entries,  the  ST_IDN  field  does
not  have  a  pre-reserved  meaning  in  the  symbol-table;  therefore,  a  language
iraplementor  may  use  the  field  in  whatever  way  it  desired.  However,  it  is
strongly  suggested  that  this  field  be  used  as  an  attribute  field  for  the
user-defined  tokens  (section  3 ♦3. 7  of  this  paper  discusses  more  completely
the  handling  of  attributes  by  the  parser  system).  The  special  pre-defined
parser  storage  variable  UDN  has  been  provided  to  hold  the  ST_IDN  value  for
user-defined  tokens  (section  3. ^•3.^).

ST_OFF:

ST_LEN:  These  2  fields  in  the  symbol-table  have  no  pre-reserved
meaning,  and  a  language  implementor  may  utilize  them  in  whatever  way  is
desired  (for  appropriate  tokens,  these  fields  should  be  used  for  the
offset  and  length  of  storage  at  runtime).

ST_BLK:  This  field  is  unused  for  pre-defined  symbol-table
entries.  For  user -defined  entries,  this  field  contains  the  block  n^umber
of  the  inner-most  block  where  the  corresponding  token  was  first  used;  this
number  can  be  used  to  access  the  parser's  Block  Tables  (section  3.^.^.2).

ST_SIB:  This  field  is  used  to  link  together  all  user -defined
symhol-table  entries  in  the  same  Block.

ST_LNK:  The  field  is  used  to  link  together  all  symbol-table
entries,  regardless  of  what  Block  they  are  in.

ST_PDE:   This  field  points  to  the  pre-defined  symbol-table
entry  for  a  token;  if  no  pre-defined  entry  exists,  it  points  to  a  special
"null"  pre-defined  symbol-table  entry.

ST_NTP:   This  field  points  to  the  NAME  table  entry  for  the
token.  Note,  that  the  NAME  table  itself  is  not  considered  part  of  the
parser  environment,  and  is  thus  not  accessible  directly  as  a  parameter
(however,  a  SEM/Vntic  routine  may  reference  the  NAME  table  fields).

ST_DVL:   This  field  is  unused  for  pre-defined  symbol-table
entries,  and  also  for  user-defined  entries  that  have  no  corresponding
dope  vector  entry.

If  a  user -defined  symbol-table  entry  requires  a  dope  vector,
the  language  iraplementor  must  provide  a  SEMAntic  routine  (section  3.3.8)
that  will  get  a  dope  vector  entry  from  the  symbol-table  manager  and  then
set  ST_DVL  to  point  to  this  dope  vector  entry.   Then  any  dope  vector
field  can  be  referenced  indirectly  through  this  ST_DVL  field.

Dope  Vectors:  A  Dope  Vector  contains  these  fields:  ST_DIM,
ST_LB1,  ST_LB2,  ST_UB1,  ST_UB2.   It  is  very  important  to  remember  that  a
Dope  Vector  is  only  available  for  a  token  if  it  has  been  explicitly
requested  (see  ST_DVL  above).   Then  the  Dope  Vector  fields  must  be
referenced  indirectly  through  the  regular  symbol-table  entry's  ST_DVL
field.   Note  that  the  following  definitions  of  the  Dope  Vector  fields

k3

are  actually  only  suggested;  a  language  implementor  may  use  the  fields
in  any  way  that  is  desired.

ST_DIM:  Contains  the  number  of  dimensions  for  an  array.  The
niimber  can  be  either  1  or  2.

ST_LB1:

Contains  the  lower  bounds  of  the  dimensions  of  an  array.

ST_LB2
ST_UB1
ST  IIB2

Contains  the  upper  bounds  of  the  dimensions  of  an  array.

3.U.i^.2  Block  Table  Fields

The  compiler ' s  symbol-table  and  symbol-table  manager  were
designed  to  allow  normal  block-structuring  with  respect  to  user-defined
tokens  to  be  possible.  ■  For  this  purpose,  two  block  structure  tables,
referenced  as

BT_LNK  and  BT_KID
are  included  as  part  of  the  parser  environment.  Both  of  these  tables
axe   indexed  by  an  appropriate  Block  number,  usually  contained  in  the
parser  variable  BLOCK  (section  3.^.3.6).

The  BT_KID  table  entry  for  a  Block  consists  of  a  pointer  to
a  symbol-table  entry  for  a  token  that  has  been  declared  by  the  user
in  that  block;  then  the  rest  of  the  symbol-table  entries  also  in  the  same
block  are  chained  together  through  the  ST_SIB  field  in  the  normal  symbol-
table  entries.  A  zero  ST_SIB  field  ends  the  chain.

The  BT_LNK  table  entry  for  a  Block  contains  the  number  of
the  block  containing  the  current  block.  This  information  can  be  used
to  find  the  outer  block  declarations  of  a  najne.  It  is  also  useful  for

maintaining  correct  nesting  of  variable  definitions  at  runtime.

The  language  implementor  has  control  over  the  opening  and
closing  of  blocks.   To  (open,  close)  a  Block,  it  is  necessary  to
provide  compiler  units  (blkbgn,  blkend)  to  be  executed  in  TUTOR.  These
units  will  change  the  current  value  of  BLOCK  (section  3.^.3.6)  and
perform  other  appropriate  modifications  for  the  desired  action.
EXAMPLES :

The  following  are  all  legal  table  references;  the  references  to
the  dope  vector  fields  assumes  that  a  dope  vector  has  been  allocated  in
the  symbol-table  "^or  the  token  (see  discussion  above  on  ST_DVL  field).

(udstp)   , corresponds  to  CLASS.

(pdstp)   , corresponds  to  PDN.

(udstp)   , corresponds  to  UDN.

(udstp)

(udstp)

(udstp)

(udstp)

(st_dvl( udstp) )

(st_dvl( udstp) )

(block)

(outerblk)   ,  where  outerblk  is  an  ALLOCATEd  variable

whose  value  is  a  valid  block  number,
(var)   , where  var  is  an  ALLOCATEd  variable  whose  value
is  the  symbol-table  pointer  (udstp)  for  a  token,

st_

_typ  (

st_

idn   (

st_

_idn    (

st_

_off   (

st_

_len   (

st_

_blk   (

st_

_dvl  (

st_

dim   (

st_

_ub2    (

bt_

_kid   (

bt

kid    (

st  typ

U7

CHAPTER  k.
k,      THE  PARSER  TABLE  INSTRUCTION  FORMS
k.l     Notation

This  chapter  will  describe  the  actual  instruction  forms  that
appear  in  the  syntax  parser  table.  It  is  these  instruction  forms  that
are  exsunined  and  interpreted  by  the  parser  in  the  compiler  to  determine
whether  a  user's  program  is  syntactically  and  semantically  correct.

The  basic  instruction  size,  like  the  word  size  in  the  parser
storage  area,  is  12  bits;  this  chapter  will  refer  to  these  12 -bit  packages
as  "words".  Some  instructions  require  exactly  1  word  in  the  table;  others
require  exactly  N  words,  where  N>1;  and  some  instructions  require  a
variable  number  of  words,  depending  on  the  value  of  fields  within  the
initial  words  of  the  particular  instruction.

The  following  notation  will  be  used  on  the  following  pages:
[state]:  state  number,  points  to  an  instruction  word  location

somewhere  in  the  syntax  table  (examples  are  [call-state],
[return-state],  [true-state]  ).
<parm>:   an  instruction  parameter  or  operand.  A  <parm>  is  either
a  number  (numeric  constant)  packed  directly  in  the
table,  or  else  the  address  of  a  variable  in  the  parser
storage  area.  Thus  a  <parm>  can  take  1  word  (for
numeric  constant  packed  in  the  syntax  table,  or  a
direct  address  of  a  number  in  the  parser  storage
area),  or  2  words  (an  indexed  address  of  a  number  in
parser  storage),  or  3  words  (for  a  doubly-indexed
address  of  a  variable  in  the  parser  storage  area) .

k8
The  forms  of  the  <parm>s  used  in  particular  instructions
are  determined  by  the  actual  instruction  numbers  (op  codes),

^.2  The  Table  Instructions

U.2.1  SCM  Table  Instruction
TABLE  FORM:

/OOOQOOOQQOOO/
ACTION:

Sets  the  parser's  STATE  variable  to  point  to  the  following
instruction  in  the  table,  and  returns  to  the  lexical  analyzer  for  another
token.  When  another  token  has  been  input  to  the  lexical  analyzer,  the
symbol-table  manager  determines  the  symbol-table  location(s)  for  the  new
token  and  parsing  resumes  with  the  saved  STATE  instruction.

1+.2.2  ALLOCATE  Table  Instruction
TABLE  FORM:

/dddddddO-OOOl/
ACTION:

Pushes  ddddddd  new  entries  on  the  parser's  variable  stack.  No
initialization  of  the  values  of  these  entries  is  performed;  the  stack  pointer
is  merely  incremented.

1^.2. 3  DEALLOCATE  Table  Instruction
TABLE  FORM:

/dddddddOOOlO/

^9

ACTION;

Pops  ddddddd  entries  off  of  the  parser's  variable  stack.  The
value  of  each  entry  that  is  popped  off  is  traced,  and  the  stack  pointer  is
decremented.

k,2.k     Call  Table  Instruction

TABLE  FORM:

/dddddddOOOll/
/xxxxxxxxxxxx/
/yyyyyyyyyyyy/

word  1
word  2
word  3

to

^

:      Optional

/yyyyyyyyyyyy/  word  3  +  ddddddd  -  iJ

ACTION:

xxxxxxxxxxxx  :  [call-state]  table  location.

ddddddd  :  number  of  multiple  return  points  for  the  called
procedure.  If  ddddddd  =  0,  then  the  procedure
is  normal  (not  multiple  return).

yyyyyyyyyyyy  :  the  multiple  return  table  locations  (if  any).
Pushes  the  table  location  for  word  3  onto  the  parser's  return-state
stack.   Then  resumes  parsing  immediately  at  the  table  location  [call-state].
See  section  U.2.5  for  a  discussion  of  the  actions  that  occur  when  a
procedure  is  returned  from  via  a  RET  instruction.

Note  that  any  arguments  that  are  passed  into  the  call-procedure  are
passed  by  using  ASSIGN  instructions  immediately  proceeding  the  CALL
instruction  (see  section  i|.2.8.3  for  a  discussion  of  the  ASSIGN  instruction).

50
i^.2.5     RET  Ta."ble  Instruction
TABLE  FORM:

/dddddddOOlOO/

ACTION:

Causes  a  return  from  the  current  procedure  to  where  ever  it  was
called  from  (see  section  k.2.k   for  a  description  of  the  CALL  instruction).

Pops  the  return-address  table  location  off  of  the  parser's  stack;
this  location  corresponds  to  the  work  in  the  CALL  instruction  that  follows
the  [call-state]  number.

If  ddddddd  =  0,  then  the  procedure  does  not  use  the  multiple-
return  feature,  so  the  popped  address  corresponds  to  the  location  of
the  instruction  that  follows  the  original  CALL  instruction.   Therefore,
parsing  is  resumed  immediately  at  this  location.

If  ddddddd>  0,  the  procedure  does  use  the  multiple  return  feature,
and  ddddddd  is  the  number  of  the  multiple  return  address  to  use.   For  this
case,  the  multiple  return  addresses  are  all  packed  directly  in  the  parser
table  following  the  CALL  instruction's  [call-state!  number;  the  return
address  table  location  that  was  popped  off  of  the  parser's  stack  points
to  the  first  of  these  multiple  return  addresses.   Therefore,  the  ddddddd 'th
multiple  return  address  is  selected  from  the  parser  table,  and  parsing
resumes  immediately  at  this  address.

Note  that  when  the  parser's  stack  is  popped,  the  old  entry
at  the  top  of  the  stack,  as  well  as  the  stack  pointer  itself,  are  traced.

i

51

k.2.6     SEMA  Table  Instruction
TABLE  FORM:

/dddddddOOlOl/
ACTION:

Causes  the  ddddddd'th  TUTOR-coded  semantic  routine  (supplied  by
the  language  Implementor  for  each  language)  to  be  executed.  The  semantic
routines  are  named  Sml,  Sm2,  ...,  Sml5.

Any  parameters  for  the  semantic  routine  are  passed  In  the  "parm"
parser  storage  variables  (see  section  3. 3 .8)  by  using  the  appropriate
PASSION  Instruction  (section  ^+.2.8.3)  Immediately  preceding  the  SEMA
instruction.

k.2.7     Unconditional  Branch  Table  Instruction
TABLE  FORM;

/mmOOOOOOOllO/        word  1

/  [true -option]    /         word  2  (and  possibly  word  3)

ACTION:

Causes  an  unconditional  interpreting  of  the  [true-option]  field,
based  on  the  value  of  mm:

mm  =  0  :   [true-option]  is  an  instruction  table  location;  for  this

case,  parsing  resumes  immediately  at  this  new  location.
mm  =  1  :   [true -opt ion]  is  an  error  number;  for  this  case,  control
passes  from  the  parser  to  the  error  system,  with  this
error  number.

52

mm  =  2  :   [true-option]  is  the  direct  address  in  parser  storage

of  an  error  number;  control  passes  to  the  error  system

with  the  value  of  this  error  number.
mm  =  3  :   [true-option]  is  the  indexed  address  in  parser  storage

of  an  error  niimber;  control  passes  to  the  error  system

with  the  value  of  this  error  number.

U.2.8  General  Parameter  Table  Instructions

i+.2.8.1  General  Instruction  Form
TABLE  FORM;

/mmcccccppppp/      word  1

/  table  parameter  1  .    /       1,  2,  or  3  words

/  table  parameter  2    /       1,  2,  or  3  words

/  [true-option] j  1,  or  2  words,  only  for

conditional  branch  instructions,
ACTION;

This  instruction  form  is  used  for  all  instructions  that  use
2  parameters  as  operands.

PPPPP    tells  what  the  2  parsuneter  table  forms  look  like:
aa  =  7 J  ai  =  8,  a(ii)  =  9?  ac  =  10,
ia  =  11,  ii  =  12,  i(ii)  =  13,  ic  =  lU,
(ii)a  =  15,  (ii)i  =  l6,  (ii)(ii)  =  17,  (ii)c  =  l8.

53

where

a  :  direct  address  is  parser  storage  (l  word)

i  :  indexed  address  in  parser  storage  (2  words)
ii  :  doubly-indexed  address  in  parser  storage  (3  words)

c  :  table  constant  packed  in  the  peirser  table  itself  (l  word).
Based  on  the  value  of  ppppp,  the  2  parameters  are  decoded  into  both  their
addresses  and  their  values;  for  this  discussion  let  the  2  parajneters  be
denoted  as  pi  and  p2,  and  the  decoding  yields:

val(pl)  and  addr(pl)       and

val(p2)  and  addr(p2).
These  2  parameters  are  then  used  according  to  the  value  of  ccccc:

bceq  =  0,  bene  =  1,  bcgt  =  2,  bcge  =  3?  belt  =  k,   bcle  =  5?

bcnotall  =  6,-  bcnotany  =  t,  bcall  =  8,  beany  =  8,

assign  =  10,  passign  =  11,  maskon  =  12,  maskoff  =  13,

addit  =  ik,   subit  =15.
The  BC  table  instructions  are  described  in  section  ^.2.8.2,  and  the
environment -changing  instructions  are  described  in  section  ^+,2,8.3.

^.2.8.2  C^onditional  Branch  BC  Table  Instructions
ACTION;

Causes  the  comparison  of  the  values  of  2  parameters,  X  =  val(pl),
Y  =  val(p2).
This  chart  describes  the  TRUE  condition  requirements  for  each  ccccc:

51^

ccccc

condition

0

bceq

1

bene

2

bcgt

3

bcge

1+

belt

5

bele

6

bcnotall

7

bcnotany

8

be  all

9

beany

Condition  TRUE  if
X  =  Y
X  ^  Y
X  >  Y
X  >  Y
X  <  Y
X  <  Y

((  X  $mask$  Y)  ^  Y  )
((X  $mask$  Y)  =  0  )
((  X  $raask$  Y)  =  Y  )
((  X  $mask$  Y)  /  0  )
where  "$mask$"  performs  the  bit-wise  "AND"  of  X  and  Y.
If  the  comparison  of  the  2  parameters  yields  FALSE,  then  the  [true-option]
is  ignored,  and  parsing  resumes  with  the  next  instruction  in  the  table.
If  the  comparison  of  the  2  parameters  yields  TRUE,  the  [true-option]
is  interpreted  as  follows:

mm  =  0  :   [true-option]  is  an  instruction  table  location;  for  this

case,  parsing  resumes  immediately  at  this  new  location,
mm  =  1  :   [true-option]  is  an  error  number;  for  this  ease,  control
passes  from  the  parser  to  the  error  system,  with  this
error  number,
mm  =  2  :   [true-option]  is  the  direct  address  in  parser  storage  of
an  error  number;  control  passes  to  the  error  system,  with
the  value  of  this  error  number,
mm  =  3  :   [true-option]  is  the  indexed  address  on  parser  storage  of
an  error  number;  control  passes  to  the  error  system  with
the  value  of  this  error  number.

55

U.2.8.3     Parser   Environment -Chajiglng  Table   Instructions

ACTION:

All  of  these  instruction  forms  involve  changing  or  modifying
the  value  of  the  first  parameter  in  the  parser  storage  area.   The  follow-
ing chart  summarizes  the  actions  of  the  different  instructions:

ccccc

instruction

10

assign

11

pas sign

12

maskon

13

maskoff

ll+

addit

15

sub  it

action

val(pl)  -val(p2)

val(pl)  ^addr(p2)

val(pl)  ^val(pl)  $union$  val(p2)

val(pl)  *-val(pl)  $mask$  (-val(p2))

val(pl)  *-val(pl)  +  val(p2)

val(pl)  ^val(pl)  -  val(p2)

where  "$mask$"  is  the  bit-wise  "AND",  "$union$"  is  the  bit-wise  "OR",  '
and  the  (-val(p2))  for  the  maskoff  instruction  is  the  bit-wise  complement
of  val  (p2).

Before  any  of  these  instructions  are  executed  by  the  parser,  the
old  value  of  parameter  1  is  traced  if  necessary  (see  section  3.3.9  for
exceptions).

Note  that  the  only  current  use  of  the  PASSION  table  instruction
is  for  passing  parameters  to  semantic  routines;  there  is  no  assembler
source  form  of  the  PASSION  instruction.  •

56

The  following  are  some  illegal  table  references:
st_idn  (st_dvl(var) )   ,where  var  is  a  local  ALLOC AT Ed  variable.
st_typ(l)   , constants  are  Not  allowed  in  current  version.
st_ub3  (udstp)   , there  is  no  st_ub3I

57
CHAPTER  5.
5.    THE  COMPILER'S  TABLE  MAINTENANCE  SYSTEM

5. 1  Maintenance  System's  Purpose

Included  as  part  of  the  compiler  system  that  has  been  implemented
on  PLATO  rv  is  a  general  compiler's  table  maintenance  system.   This
maintenance  system  is  designed  to  allow  language  implementors  to  completely
specify  all  of  the  tables  that  are  required  for  a  particular  language  to
be  recognized  and  used  as  part  of  the  computer  science  compiler  system.
The  maintenance  system  also  allows  changes  to  be  made  and  tested  to  exist-
ing "stable"  versions  of  the  compiler  system  tables  without  disturbing  the
"stable"  version  until  the  modifications  have  been  throughly  debugged.

The  remainder  of  this  chapter  will  discuss  the  utilization  of  this
table  maintenance  system.

5.2  General  Operation  of  the  Maintenance  System

The  compiler's  table  maintenance  system  maintains  a  large  dataset
file  that  PLATO  IV  stores  on  a  disk.   Each  "used"  block  on  the  dataset  is
associated  with  some  particular  language  implementation  for  the  compiler;
all  of  the  compiler  tables  for  a  given  language  are  stored  within  the  blocks
that  are  allocated  for  that  language.

A  language  implementor  is  allowed  to  modify  any  of  the  tables
that  are  stored  on  the  dataset  for  the  language  that  is  being  implemented.
Experimental  versions  of  an  existing  language  are  easily  created;  these
new  versions  are  allocated  their  own  disk  space,  which  allows  modifications
to  be  made  without  disturbing  the  original  version.   Figure  5.1  shows  all
of  the  dataset  FILE  MANAGEMENT  options  that  are  available  to  a  language
implementor.   Figure  5.2  shows  an  example  of  the  dataset  directory.

58
The  actual  compiler  itself  utilizes  a  set  of  tables  that  are
stored  in  "common",  which  is  an  Extended  Core  storage  file  associated
directly  with  the  compiler  lesson.  The  table  maintenance  system  allows
a  language  implementor  to  update  the  compiler's  "common"  tables  from
the  copy  of  the  tables  that  are  stored  in  the  dataset.

Thus,  the  maintenance  system  allows  a  language  implementor  to
completely  maintain  the  tables  that  are  used  by  the  compiler  system.

5.3  Logging  into  the  Maintenance  System

The  table  maintenance  system  requires  that  each  language  in
the  dataset  be  code-word  protected.  When  signing  into  the  maintenance  system,
the  proper  code-word  must  be  typed  in  before  the  tables  for  a  language  may
be  modified  (figure  5*3  illustrates  this  process).

After  the  proper  code-word  has  been  entered,  the  language  name  ^
must  be  specified  (see  figure  5.^).

Once  the  language  name  has  been  correctly  typed  in,  the
language  ijuplementor  is  officially  logged  into  the  table  maintenance
system.  At  this  time,  a  number  of  options  are  available  (see  figure  5.5) j
including  editing  or  assembling  the  syntax  language  source  text.  Access  is
also  allowed  to  any  of  the  FILE  MANAGEMENT  options  mentioned  above  (figure
5»1).

Note  that  the  table  maintenance  system  also  allows  someone  to
sign  into  the  system  without  requiring  that  a  code-word  be  entered;  this
puts  the  person  in  an  INSPECT  ONLY  mode,  in  which  the  person  may  examine
any  of  the  tables  for  a  language,  but  is  not  allowed  to  make  any  change
to  any  language.

59

u
u

CE

c

0

IZ

'U

l3

Q

0

[_

Q

^

1.0

(.)

•It

0

•...'

r—t

V

■— 1

!-

■t-t

mi

jQ

0

■Vh

4-'

"■

'Z

Lft

ai

in

•13

4-'

-•■■■.

X

.

1.0

'1'

in

0

i_

•  f— 1

iHi

•^-1

11*1

1

4-'

0

<4-i

r|_

I'"*!

mi

1

■4^'

ai

111

4^'

i'  ^1

n-

1  1

r<

.'1'

HI

1 — 1

V

. — 1

'I'

l.\

mi

mi

V

iTi

' — '

1 — (

t~t

1 — t

■  »-H

r-H

Q_

'J

C.I

CZ

'&-

+•'

■  r-t

■  f-H

'V-i

•  1-H

''-H

•  f-l

C.>

'— '

0

•  r-l

'^

■■.-.

'-*~l

Vh

'^H

n

V

\t—

r]_

T5

-■•\

'■Ci

iri

N

Zi

•-^H

I.M

ID

illi

iTi

iri

mi

•  r-H

-1-'

.»—

X

1 1

4-'

4-"

n

n:i

a.'

7j

'i.'

;l'

n

II"

'l'

mi

HI

!-

a.i

-4-'

.»-

4-'

lZ

C

0.0

4-'

•  rH

t.n

V

a

4-'

iri

1].'

1-

mi

+-'

1

1*—

mi

4-'

mi

^;.

1.0

•  t-l

>i'

4-'

p

!-

Ifl

iri

•I"

•  r~^

4^

s:

;::

T5

U

\w\

A-

'1.1

'I'

^

U

c

113

0

HH

LJ

(.  ;i

u

01

Qi

LJ

0

u

M

Q

u

r*

Q

«:..

f"4

i;"'i

h

'-•.'

11.

[1.1

i;.'

111

mi

Q-:

L«L

1'

1

^•'

Ti

H.i

1. 1 .;

j.i.i

...J

r""

ilj

iT'i

ri

1.  j

F--

'X

''::-C

J._

('  1

<L

tf

LJ

1

ilJ

1

l.l'l

11'!

ifi

Il"l

'V

ai

o

H

H
(in

nj   _Q     Q   771     'U   "^     ij.O

LI.    U

6o

•

IT*

LA

LTi

LTl

\

c^

\

LTi

U-'i

r-.

ITi

ITi

r--

lA

lA

ITi

CvJ

u^

U"i

CJ

U"i

ITi

r--

r--

\

r-

r--

\

r^

r-w

r>.

\

f--

r-u

\

f-

f-

\

\

cr.

\

\

as

\

\

\

in

\

\

if\

ITi

\

\

rv.

1^

•.N

fv.

CO

CJ

cri

LTi

r>.

IV.

<sj

rv.

r--

iS!

r--

K

«.D

(\j

(M

s

CxJ

O-i

s.

C-J

r-

CvJ

f^j

CnI

CSJ

\

C^J

C  J

\

\

U^

\

\.

LT.

\

\

\

\

\

\

T

\

\

.'— ..

>H

LTi

Wt

!S!

U^

u^

G!

ir»

cr.

u^

U'l

lil

lA

uS

f-

C\l

U-1

U'l

1.0

s

>

!S-

iS)

t^

is.

55

oj

iSJ

i:sj

0

iSj

'S!

0

\

!S!

iS!

'l'

g

£D

c:

fl

\

is*

ij.i:i  1-.
ii:i  CO  ■

H

'    ^

Q

■;:

5=

0

1^

f:

Cv

s:

iS!

i~

■;:

>-i

c:

£1

-r—l

f:

C

_i

U

0

0

0

0

0

0

0

a-.

0

0

a-"

0

0

ilMJ     II.

o

h-

/— s

/ — ^

-r

^

c:

C

s

hH

.-^

._^

'15

^

^^

-li

^_^

c

^^

^_

1.0

^_

^ .

1.0

0

.' — ^

.—^

mi  yli

Q

iTi

ITJ

1.0

iri

iri

lO

(j

0

iTi

iT:i'

(J

>ii

n:i'

0

03

■'T1

' — *      x"*

CO

UJ

ui

UJ

Q,

in

10

^q

g

10

1.0

•• — ■•

10

1.0

•• — •■

,— ..

10

1.0

I-H

1

^^

0

0^

O

- 1— 1

^_^.

U

0

4-'

0

Q

iT5

■• /'

0

0

r--   Z

h-

■s '

.— <

V '

.— <

li

cr

V '

^w_/

t-

^^

^-'

:-

T5

""—'

'—^

—1    HH

o

in

i-

u

.—1

f

1*-

u.

1 — 1

'^■^

i-

t-

t-

'l'

J-

S-

ri

f-

i-

u

■"-■■  |X

<D

^j-

'li

a.'

m

•1'

V

'Hi

U

-— ■■

'l'

a.i

X!

'1.1

V

t:

in

'1.1

'1.1

.i_

^

_l

s_

^.

Tis

!-

y-

"0

1'

l~

SI

!Z

li

*;:

A-

iTi

f^

c

d

LJ

+3

-— <

t-

l"*!

1.0

1 — 1

,^

• — 1

M

n

1 — 1

i— 1

fil'

Ch

-  f-<

•  *— •

•  rH

■  .-1

•  1— 1

•  r-l

d

"1

■  i-H

•  1-4

()

■  f-t

■  .-H

'Hi

iTi

■  I-H

•  r-l

O

p

E

4-'

lZ

£

+-'

(.)

u

cz

a

1.0

cz

I—
1_

!-

I'j

£

1 —
1—

"^

0
H

X

1^)

<L

lO

I-H

t:

^i:

1—

§

>

u
o

u-^

^

^

iSi

^

CO

_j

CO

Ch

a^

■r-H

T-l

a^

VD

.,—1

*— 1

^

r-^

T-H

-r

cr-

r-l

CO

Cr^

I-H     -^

^

q:

CD

■_J  ■  C-J

#

1— t

6

h-       II

CO

•

K

u

.J

CO

Q

W

u

GD

^ci

(_.'  ifi

&i

(T

u

\St

'l-

j->

Ij-^

t

u^

1-

U"^

^

u-\

1.1 1

U"»

u-^

^

U-1

U1

t-

<r  ~\

O

u

H

o
_1

Q-

co

H

q:

#

CD

h-i

Q

C

Ol

3

•D

-^-r

0

"S
o

a;

1)3

0

r— f

1—

y

r—»

i:«o

ijO

l'

0

•D

— 1

0

!^

O

'I'

»--•

<<-i

0

■ — 1

■  r-»

n5

0

C-J

i)

Iri

u

•

o

■  r-t

4-'

0

^

0

•

SI

0

Cl

CL

1)

v.:

Q_

4'

!-

■  r-4

4-'

iii

1/)

J-

»-•

•  r-t

Ai

3

0

X!

r— «

1.0

1.0

»-<

1.0

i-

i-

m

!-

CL

iTi

0

.— •

E

0

a.i

$1

0

X!

•  r-l

iTi

. — 1

•  r-<

0

irj

'P

.0

_l

ji

Cm

a.

•-•—1

u

!-

m

0

y

1 — H

Cl

Cl

1 — 1

•-n

-0

ijj

■+->

61

_l

:;^

(D

-.i

^

V

■s-

w~i

"^

_r\

#

Tlj

«-

<•    «
•*     <>

<-  *  :>

-•^  ---

*  J-  *

*f  V  <■

■>■     •:'i    ■>

^

c>

^

«■

O

V

*

■«•

•s-

m

.•^

T=r

LO

>-

*J

•

-*

•«•

U

>

*

-:{.

■i^

*

•S-

:■<.

*

•<■

^

1

X

u

"Z.

m

1

ii)

'V

10

m

U7

-0

V

v

i-

CL

>

•  f— •

c

u

V

Q.

a.i

( >

0

t-

0

U

U

x

CL

4-'

CO

•  f-l

23

H<

1

1  ■)

f

■^

v

0

u:i

<*-t

'1.1

1

r— f

H

•  r-t

v

Li_

y

1

o

.'0

:\

10

,— •

•l'

1

L

■r-t

Q..

(If

Q-

X

\~

I

CO
I

i.n
'1.1

i

w

Q
O
O

fc>D
CO

H

I-    4-'

62

<f     C'

rn

O
h- 1

h-
CL
O'

UJ
Q
_J
I— I
Zl
iD

U

_J
lQ
(E

•U

t=

1

■TJ

a

h-

C

£

X

0

Ld

1.1

y

—T

CiO

1—1

1

ITS

■— «

IS

Q.

10

£^

1.0

•li

1— «

iT5

c

V

C

x:

rH

y-

"U

a

>v

UJ

i.i:i

"r"-'

iri

r^

Cl

'11.

E

^

'I'

i

ilLl)

0

irj

_'\

I-

iij

'I'

!=

ij.d

k-

'I'

•Ci

« — 1

o

cr

1

j-

irj

'X

(Z

Q

1

1

10

10

lO

1.0

'U

•l"

S-

1*

1*—

CL

LL

63

CO

o

H-
CL
O

Of
Ul
Q

Z)
(D

Ld
-I
CQ

O

I)

\
\

0

111

iri

-P

:\

'i'

•i"

V

4-'

u?

10

lii

iri

-i

J

Iji)

•ti

•   \fi

V    V

o  — .

i-  Jii

•

li     irj

■

'IJ

0   4-'

£

(.)

1.0

m

t-

^

u

li

'1.1    V

m

O

m  '—

ni

10

ii5     ■-<

■  <-i

Zi     Q.

.

Xi

'l'

i'lii    £

uO

t:   •:•

a.i

1.0

115

ii:i    u

. — 1

•  r-l

Zi

.— <

■  ft

10

Ci.O

r

o

>.

~r^

iTi  a:

. — (

aj

. — 1

.•-'  -i_.

()

!Z

a.i

!^      0

•15

y

A

-•*"\

'l'

(^i

iVi

1.0   —1

i-

.— H

C

-1-^'

,— 1

4-'

iTl

o_

l~

'l.i    mi

0

_.:'.^

r— 1

i~'l

■  —t

0

ijj

J  "J    Zx.

4^'

X

4-'

£   '■<-<

V

~i~'

Ij.l        -r-.

Q.

iz

r-<

ii;i  i;?

0

r  'i

1.0     0

Zi

iri

•  1— 1

if

05     £

•r— .

-f-'

t-

Cl

o

0    0

o

O

i;ji

4^'

4-'    4-'

4^'

'■*-,

C

iTIi

-■^v

,— ,.    ,—..

_.

.

'U

1-H

i-j   '>:•-.'

V

U-i^

1.0

0

0

^

o

'I'

III

ilill
il5

0

J

-r,

•  *"•

m

C

-p

f-

0

!?-

iTli

3:

r  j

. — 1

iJi

1.0

4-'

i-

iri

w:

1

Q

^U.

■+•'

u

v

l,'\

zi

M

!.-

(  'l

ITJ

il'i

r"

' '.'

-.•'■..

'I'

V

CI

1.0

ijj.'l

mi

'II

I  I

i:  a:  i

'w'  I-  OD

c  'X  '.r

iTj  q  j

I  I  I

lO  1.0  lO

ii!i  ifi  in

'l'  V  'IJ

1-  t-  L

fl-  Q-  Q.

o

H

EH

LPv

H

6k

3-h     Preparing  the  Syntax  Table  for  the  Compiler  System

Once  a  new  language  has  been  created  on  the  dataset,  the
language  implementor  then  proceeds  with  the  "edit  the  syntax  language
source"  table  builder  option;  the  initial  version  of  the  syntax  source
specification  is  then  typed  into  the  syntax  blocks  for  the  language.

When  the  syntax  specification  is  complete,  the  table  option
"to  assemble  the  syntax  language  source"  should  be  attempted.   This  option
assembles  the  source  form  into  the  compiler's  table  form;  as  the  assembler
is  running,  the  current  label  in  the  syntax  source  specification  is  displayed
to  allow  the  language  implementor  to  follow  the  progi^ess  of  the  assembler
(see  figure  5*6) .

Any  errors  in  the  syntax  specification  will  be  flagged  as  they
are  detected;  no  attempt  is  made  by  the  assembler  to  correct  the  error--
the  language  implementor  must  return  to  the  table  builder,  fix  the  indicated
error,  and  attempt  to  reassemble  the  syntax  source  until  no  errors  are
detected  by  the  assembler.

When  a  correctly  assembled  table  has  been  prepared  by  the
assembler,  upon  returning  to  the  table  maintenance  system  the  language
implementor  should  update  the  copy  of  the  tables  on  the  disk  (see  figure  5* 7)

Once  the  syntax  table  is  prepared,  as  well  as  all  of  the  other
compiler  tables,  the  table  builder  option  "to  jump  to  the  compiler"  should
be  taken  to  initialize  or  update  the  compiler's  actual  "common"  table
version  (note  that  this  version  of.  the  compiler  will  also  contain  the
appropriate  set  of 'TUT(t)R  semantic  routines  for  the  given  language  (see
section  3.3.8)).   Then  any  logic  errors  in  the  syntax  specification  must
b^  discovered  and  corrected  by  the  language  implementor.

OJil

i-

(6

«

65

U
_J

a:]

>

LJ

m
m

(_9

h-H

Ql

•  t-l

4J
0

H

0)
M

a

•H

c

C5

H

•13

GG

M

10

•f*

"0

V

^

•

-M

10

(0

c

<ll

0

"0

■r-(

0)

(U

JC

■— <

u

j:i

d

■13

(0

-p

-tJ

t-

d

d

0

0

Xl

>.

Id

<iJ

+»

4-'

Oi

Id

^

TJ

S-

a

0

d

<M

0

0

4J

+»

1

,

a!

H

_i

X

u

LJ

X

^

1

1

10

10

10

10

m

lU

i-

J-

Q_

Q.

to

CD

0)
-P

o

&

o

M

-p

CS

5^

6?

CHAPTER  6.
6.     FUTURE  DEVELOPMENT

The  table-driven  parser  system  that  has  been  described  in  this
paper  works  fairly  well  for  the  languages  that  have  been  implemented  thus
far  (PL/I,  FORTRAN,  COBOL,  BASIC).  However,  there  are  a  few  areas  in
which  improvements  could  be  made  within  both  the  actual  table  system,
and  the  corresponding  assembler  language  specifications.

One  area  that  could  be  greatly  improved  is  the  handling  of  .
semantic  routines  in  the  system  (section  3.3.8).  When  the  table-driven
system  was  first  designed,  very  few  semantic-type  instructions  were
included  (examples  are  ASSIGN,  ADDIT,  SUBIT,  etc.)  because  it  was  not
known  exactly  what  instructions  would  be  needed.  Now  that  a  number  of
languages  have  been  implemented,  the  semantic  routines  that  are  used  by
these  languages  need  to  be  surveyed  very  carefully  so  that  the  commonly-
used  routines  can  be  included  directly  as  new  instructions  in  the  system
(examples  might  be  instructions  to  open  (close)  a  block,  or  to  request
that  a  dope  vector  be  allocated  and  linked  to  a  particular  symbol  table
entry).  Ultimately,  the  hope  would  be  to  eliminate  nearly  all  actual
uses  of  the  SEMA  instruction  in  the  system,  with  the  appropriate  functions
being  accomplished  more  directly.

A  second  improvement  to  the  system  would  be  the  development  of
a  slightly  higher-level  form  of  the  assembler  language  to  be  used  in
specifying  the  syntax/ semantics  of  a  programming  language.   For  example,
a  simple  looping -type  construct  would  be  very  useful  in  the  assembler.

68

Another  useful  addition  would  be  to  allow  more  complicated  data-structures
to  be  declared  within  the  Parser's  variable  stack;  for  example,  it  would
be  useful  to  be  able  to  easily  create  and  manipulate  linked-lists  of
ALLOCATEd  variables  on  the  parser's  stack.

A  final  improvement,  and  by  far  the  most  difficult  one,  is  to
develop  a  program  to  convert  a  BNF-like  representation  of  the  syntactic/
semantic  requirements  for  a  language  into  the  table-form  that  is  required
by  the  compiler  system.  Although  the  most  difficult  improvement  suggested
here,  this  would  also  be  the  most  useful  because  it  would  allow  language
designers  to  implement  new  languages  without  having  a  detailed  knowledge
of  the  actual  parser  table  system  that  is  used  in  the  compiler.

69

LIST  OF  REFERENCES

[1]   Conway,  M.E.,  "Design  of  a  Separable  Transition-Diagram  Compiler",
CACM,  Vol.  6,  (July  I963),  PP.  396-^+08.

[2]   Eland,  David,  Forthcoming  Ph.D.  thesis  on  the  GUIDE  information-
retrieval  system,  to  be  published  in  the  summer  of  (1975).

[3]   Lomet,  D.  B. ,  "A  Formalization  of  Transition-diagram  Systems",
Journal  of  the  ACM,  Vol.  20,  No.  2,  (April  1973),  PP-  235-257.

[k]     Nievergelt,  J.,  Reingold,  E,  M. ,  and  Wilcox,  T.R.,  "The  Automation
of  Introductory  Computer  Science  Courses,"  A.  Gunther,  et  al.
(editors).  International  Computing  Symposium  1973?  North-Holland
Publishing  Co.,  197^.

[5]  Sherwood,  B.  A.,  The  TUTOR  Language,  Computer-based  Education
Research  Laboratory  and  Department  of  Physics,  University  of
Illinois,  Urbana.

[6]   Tindall,  M.  H. ,  Forthcoming  Ph.D.  thesis  on  an  interactive,

compile-time  table-driven  error  analysis  system,  to  be  published
fall  (1975).   •

70

APPENDIX

This  appendix  contains  a  sample  assembler  source  language  listing  for  a
version  of  a  subset  of  the  PL/I  programming  language.

71

or

o

Uu

II

a

»— 1

a

>»

II

u.

UJ

»— t

K

>♦

*-t

»-^

21

o

II

*~

u.

aJ

z

OD

z

•-«

l/J

\xi

<

fn

xl

—

9»

»

M

-1

»^

t>

3J

■n

o

— «

^

II

•— t

UJ

<v

•—1

II

f«-

a:

«

_i

••>

/—»

II

II

X

^^

<

OJ

II

o  ■

>

Ui

•— 1

X

X  ro

a

II

^

"•— f

u.

c

cr

a

z

^.  •"*

t— •

X

»

II

•->

c

I-

>-

UJ

»-i

UJ  II

^

_l

(VJ

tr

o

a

z

•»

z

»-

a

<

X  X

(/)

•a

-*

If)

«

II

1— •

a

•

«■

UJ

>

2

»-   U.'

»

•5

II

a

«

»-

it

o

fvj  r^

(/)

II

o

♦-

»

-  _l

^

>— •

X

UJ

4t

z

1—

»

f^    »-H

^

r

•— •

'VI

*  •-•

II

u

>-

X

UJUJ

IT

II        II

c

••■

a

•

fV'

'-  I

X

UJ

a

c

oe

II

o  >

_l

a

UJ

a

II

II    IE

1-

o

1-

m  »-

•

_l

3t   <1

o

_J

UJ

2

<j

X

X     ••

UJ

»

z

o

>

II

D

«

2

UJ

CO  ct

tvi

o

>-

3

X

in

O  IP

'vS

vO

UJ

•— 1

CL

1—

O

O

a

a.  a:

II

u.

u

<:.

u

z

\-  -•

•>

•—1

»

II

<l

_»

•

•— )

•-•

a

•  <i

^

«a

«

*

o

•  II

ir.

II

•— »

X

^~

u.

r^

•— t

1-

o

t-»  »

o

■y)

ac

H-t

O  X

II

X

-♦

-I

3

o

II

II

t-%

•■

^  <o

L>

=t

«t

1

II  >-

X

►—

II

_l

a:

>—

>

z

>*

II  —

at

1

X

.>^

.-»

a

X  re

1-

<l

X

<

or

1-

z

c

•-«

II

h-     H

t-

1

u

—t

o

o

a.    •

3

D

X

(-)

o

H-

UJ  t-

u.

o

yi  u

(/)

1

II

—*

o

»

•-•  — <

X

_J

<

•

u.

<i

Q  oJ

Ui

z

"ii  >-

»

■J)

I— i

r-K

o

CO

•    II

»

u.

X

OC:

Xt

•*

■2,

o

o

o

O  i

0^

t—

*»

o

»-•

t-^

(\J  K

•— t

»

o

II

n

*

fVJ

•— t

»

t;

tJ  K-

r--

C/)

-J

o

o

II

II    O

c

>*

»

X

II

II

»

o

-^

»-

a

»  z

II

t-<

<

o

o

X

X  o

I— t

o

<r

2

X

»c

1-

>C

f— t

»— 1

2

c

a

c  u.'

u.

_j

s

o

o

o

c   •

II

f— 1

>*

or

a

4^

UJ

■1

II

11

*1

»

II

»-<    •

t— 1

•— 1

o

o

UJ

^-  »c

a

II

II

D  K

II

(Tj

tr

c

^

»-

rr

X

II    LT

a

UJ

•

t->

o

o

u

C    II

1-

X

X

1-

(/)

t/)

X

<s

z

X

(/)

II

•— •

a  -^

►-

►-

•

UJ

o

o

o

<Si  X

X

<

c

UJ  OD

ra

1-

_J

CD

u

z

»-

ct

<5    II

IT

3

•

o

o

o

tr

•  a

e;

i.

UJ

cr

-3

-ti

1-

(/I

e>

O)

•

o

C  »-

>  u.

»

a.

t—

II

er

c

a

r^  c

f— »

r

X

»

lA,

►—

<r

C

^

<

•

u

ir

1-

_'  u

CT

K-

->

~-

«-

•

It  1-

»>

c

»-

f»^

•

^

•

G

I

;

c

O"

a

<;

o  a

f-^

Q:

y:

••

ic

ic

If;

X    l/^

»

o

r,

u

>*c:

♦•

#—

»-

<j

»

c

a

u

o

♦

CL    •

II

»-

-5

f-

o;

cr<

c

o  ♦

r-

»

II

»

II

h-

•■

II

Ifi

»

r?

Ul

0-

•-4

»

h-

•   -t

:

1-

►-<

<!

<

<2

r-.

e  —

c:

c

X

c

X

II

-;»

ir

H

II

cr

,—

<A

II

z.

r-

r

K

C    -

c'

V

<"

^

c

',

7

Il

•  .-

.-■

n

L

r

f

>

'^

si

1/

L

ii

II

j:

c

c

\r

M

fi

rr

II     II

!j'

f'

^1

II

■X'

"-  II

II

II

rv

II

t—

II

II

(A  u

~

r;-

"^

rv

^

LU

"T*

-?

-y

»-    _l

Ci

~

•»

U.

u.

II

'V

"

X  X

a.

X

<

X

'V

-■i

X

LO

2:

»-

1— •

-?

UJ

a:

UJ

— /

^

r^

fr— '

►— :

O  UJ

u.:

_t

<■

\^

II

■T

<

c/>

c

II    LU

(-

1-

1

u.

'^t

o

>

'

U.I

JL

_l

-r"

_i

o

IT

U.'

_i

:x

T»

_^

—     »J

►—

-:!

^

.  '

c

JT

i

uJ

X   O",

u.

j-.

Z)

vi

^—

~-

a

_J

X

»-

/;,

H-

t —

^

>-

-r

<3

c

a

h-

~   <

X

>

_i

e^

•»■

Zz.

O

_.

o

C    _J

U.'

*—

u,

a

UJ

c

<i

*>

^

H

Uu

U

u.>

a

lA,

^

>

o

ir

1  _l

vjj

UJ

X

Cj

<-

<-

_iO

UJ  UJ

-I

_J

D

a

a

u

■>

4?

2j

<

-^

IT

U",

jj

♦-

z

*— 4

<i

UJ

<

Z

u

\r.

UJ

c

u.

>-j

X

>•

2

o

\r

z

u;

UJ

UJ  UJ

UJ

t— t

II

»— 1

UJ

UJ

UJ

UJ  U.(

UJ

UJ

UJ

UJ

UJ  U)

UJ  UJ

<S

UJ

UJ  UJ

UJ

U.

UJ

<s

»— .

';

~

2-    -?

— r

vf

"Z

2

u.

2"

2

?

■7

^    *c

iT-

z

^

z

~

~"

"^

z

JCL

;^

z

^

■^

1-

?

r;

„_

u.

K

»—

t~   *-

»-

«

c

ir

II

►-

»-

c

»-

»-  »-

»-

»-

►-

»—

1^

t~-

»—

»--

u

*-

»—

►—

»-■

u

*-

z

CJ

u.

li.

u.

u.  u.

U-

H

— ■

—1

u.

U.

a

u.

U.   U-

u.

u

u.

U-

u.

\±.

u.

u

lA

U.

u.

u

u.

>-

u.

u

c

U-

u.

U.1  u

u;

t/

U-i

U-i

u.-

u.  u

u.

u

u.

u

u

u

u_

u

U

u

u

UJ

1/

U.I

*

«

c

c

Cj  C-

c  «

*

«

.c

«

/~

c   *   *

c

c  c

c

r^

c

c

c

C'

c

c

«

C:

c

c

c    *   *

c

«  «

72

o

s:

»-•

z

UJ

1-

c

K

o

»-i

«/1

■».

—  z

^»-

>-

a:

QC  3

«•«

UJ  u

(/)

o

OU.

.>«,

.-.

UJ

-I  z

1-

1-

-1  UJ

_i

□D  3

oe

4

<1  z

o

_!'-

3)

<1  U.

»«

o

X

X  •-•

OD

3)»-

<t

t-«

_l

cr

UJ

UJ    1

s:

<tz

i-i

.^>.

a  z

<l

a:

Ol

Q.  »-

>-

•-•  <

tx  ~

.-«

>

d  •-«

a:

UJ

o

Al         •-•

in

act-  '^
<to  o

>  Ui

UJ

o  —

4

>  1

1-

UJ

z

o

tJ

O  _1  o

.*«.

z

>  z  a.

OD

ZUJ

(£

-1  _l

>-*

o

»-

♦-I

Al   <I  OD

UJ  o

o  o

c  <

UJX

4

4  •-•

_J

.««

•-t

3

1-

z

»-

»-i

o  o:*

UJ  _l

ct<

Z  3

q:

1-

<

\ki

q:  O  u

—  3

»-

UJ

a:

U;Z

o

a  oc

C

UJ

..■«

4

r

o  »-*  ►-•

C  CD

<

q:  u  o

<l  Q

U.

UJ

UJ

z

»-

o

O  X

cc

..-.

X

»-  a

z  *-*

3

<I   t-H   UJ

-lUJ

UJ>

oc

»-  CD

•-I

o

>-l

o  •-•

o

1

»-

—

Al  <  UJ

*-*oc

1-

_ia:  >

oz

tret

<t  X  2

a:

<t

a:

_IO

u.

»-t

Al  U.   _J  i

cr  h-

O

OUJ  QC

UJ  •-«

•-

_J

UJ  »-•

1-

ex:

UJ

-J  •

•

cc

_»

Al  ujr;

»-  K

z

UJ  7  UJ

c  u.

_JZ

o

—  QC

«/)

<T

5

<  Q

tr

♦

<

Al

—  a  z

(/-t   «3

3

03  01

Z  UJ

UJUJ

UJ

_fK

«-

</:

I

3  «/■'

<

U.'

— -

•—

•~-

•.»  •«»

•Wr     'W

a

•-  Z  UJ

3  O

oc  —

c

<  (/>

UJ

•

o

Z  UI

UJ  >

to

2

a

e  CD

~-

cr  —  (X

■w     «_

<IO

»"

z  «-

z

X

•

'»-'

—  x

_i    »

^

<i

Z"

z

•-

I"  _i  u.

^  »-'

K

<♦-  "-

-J  _1

_ll>

>

cr  u.

c-

<

•

<

a  :=

c

2

~

►-

C

C     li    .-

•-  c

C

;  «/  c

C    L

-^  D

«

u    »-

c

^

V

D

^

<I  o

»— 1

a

T

z-

o  q:  X

rr  ♦-

<c

_r  z  "?

C  T

U.K

or

1-    T

tt

\JJ

<r

■5-

»-•  •

»-

l/l

o  CL  a.

a.  a.  CL

»-  »-

32

UO  I/)

O  <

UJZ

a

X  »-

H-

lO

X

X

3  it:

a.  a.

«i

cr

o  o

o  c>  o

(/■<  <

a

Q  O  £K

S  _J

(XUJ

«t

UJ  to

KA

l/>

<

O

z  1/5

^  I

o

•~i

'aJ

<1

z

<-

4-    <

>  u-

c

_i

5"

UJ

uJ  UJ   UJ  UlI

UJ  UJ

Uj

Uj  UJ  UJ

U;  UJ

UJUJ

UJ

UJ  UJ

UJ

_l

X

_»

tj

<

5-

x

z  2:  ^

s:  2

r

X  X  T

2:  X

X  X

X

X  X

X

o

-^

UJ

UJ

-I  lij

_i

Z

<

<t

«a  <*  <t

«a  <r

«5

<s  <s  <t

<i  <

<x  <t

<t

<I  <

<

U>  X

X  U-

<  1-

<I

or

Z

z

7^  ^  ^

z  z

z

Z  Z  T

z  z

zz

z

z  z

z

u.

4

<t

<o

X  <

o

to

o

X  2

z

o  u

c

or

u-

lA^

(n

Kf)  KTi  KfS

\r,  ifi

{T

ir,  \r  in

(/)  to

irt/>

U)

«/J  t^/

iT

c

_i  c

~

a

<*

</

<r

V    it   If:

if  if

«/

w  tr  V

tA  (/^

Ol  V

u

v;  V

V

c

z

it

ic   Z

C'  _J

UJ

U.'

_J

<:

<r

<5    «;  <

<■  <

<

<:  <i  <i

<:  <r

•c  <r

«?

c  <r

*z

z

c

tAi

to  UJ

_J

e.

_J

r

-i  _«  _•

_)  _i

_'

_'  _'  —1

_J  _J

_J_t

-^

_» _l

•

U.'

~'

<i

•5

«  <

«  *

*

«

O

u

o  o  u

o  u  o

oo  o

OO

oo

o  o  o

O

«  «

«

X

X  ♦

73

o

_

«-

z

u.>

«-

r

ut

4

4-

Z

UJ

o

V

oi

JJ

4-

Z

r

Jj

X

D

►-

ct

U.I

«■

z

o

UJ

z

cc

U'

c

X

z

4-

a

u
o

UJ

UJ

<-

•»

«-

UJ

4-

a

>

X

q:

UJ

»-H

a

UJ

a

4

u.<

♦

a

or

«-

U.'

a:

X

r

r

a.

o

UJ

r

<
<-

r

o
u.

UJ

r

UJ

»-

••

_

J.

,..

^-

(/)

«-

o

z

>—

«-

)C

■V

(/)

o

Uj

»— 1

3

t>

z

z

o

q:

<r

o

3

UJ

UJ

o

l-l

*

Ui

«

s

UJ

UJ

a

a

T

UJ

a

►-

«n

«

i/i

^

«

UJ

X

z

X

in

^*

<^

vO

2
o

X
Ui

«

«

iT

X
UJ

X

U.I

«

JJ
I/)

»

UJ

fV

UJ
X

cr

a

►-•

u

>*

(i

X

o

Ci

.— <

X

UJ

o

X

o

>—

o

X

X

2

o

-I

r\i

o

r

a:

o

a:

a.

a

a

X

c

•»

o

Ui

X

X

_I

X

X

i

JL

JC

o

a.

o

o

X

X

u

X

o

<

u

X

UJ

UJ

o

ex

UJ

•

a

a

a

o

X

o

*-

X

c

UJ

Z

»

>-

Uj

•

c

X

a

tx

UJ

UJ

UJ

a

<_)

a.

«

or

»

»-t

_i  OL

»

X

UJ

(/)

UJ

u

»-

X

•

a

>-

Uj

u^

a

_J

»-

Ol-

X

Q

ct

z

»

••

a

X

»

X

X

»>

»—

vC

»

u

Zj

o

^'

-^

UJ

o

Q.

X

\-

*

H^

X

o

1-

X

<

r^

a

o

o

o

UJ

o

u

♦

•— t

»-

Z

I

z

o

o

2

►-I

»-

»-

o

X

^

»

_J

c

>-

U.

•-t

iD

U.'

^

-I

UJ

>

«/j

X

u.

^

»

(T

o

or

r-

Q

U.'

<?

»— 1

•—1

tr

u.

U.'

o

»

u

c

u.

»

o

0^

U"

l_

O"

a

•

c

i

Zi

X

»

»

*~

•

v~

c

t/"

w

U-

a

_J

ir

t—

u-

<

»

<l

»

a

»•

•

»

»

Q

z

*—

z

z

Cr

(/^

#

iT,

X

»

if)

t-

n  <«

_l

2.

_;

^

a

z

^

2L

^—

a

c

_j

c

2.

Q-

■c

rr

X

uJ

^

4

«_

<•    '

C-

c

CA

c

u

c

a

C

c

u

a.

2'

a

K

♦

.

C

<3

•

c

___t

<3

c     t_

c

c

c

c

c

Q

c

•

L

•

U

L>

c

L

c

L

c

C

r

f/     •

i"

•■

a

•-

~"

•■

»

»

»

^

Ui

~

ry

"T^

»

•

— .

»

»

U'

6  uj

C5

UJ

UJ

'J3

UJ

U_'

UJ

UJ

o

z

UJ

I— •

cr

UJ

UJ

_l

&

UJ

UJ

c/)

r  .

»— 1

.'■!!   "Z

<

~     ~7

iT.

H^

z

n  n

.'    ^

r.  7"

•— •

~

<-

■.^

'

f—

~

^

z  z

*'    «

t~

~'

^.

"^

^

a

yi

1^

^■

1

-i

iTi

^

a

^~^

yi

/  o

^■

-^    L;

_J

^'

-3.

^      4?

ci

>-  ~c

o  tr

<-.  (^

u.

u  o  u

t/~

o

u  u

u  o

uej

t/.

U    X

U-

U   X

<3.

o  o

(J-

o  o

o  •-.

<_

o

c

CJ

_J  o

_  X

%i.

(/)   X

;/!

i/>   r

V)

=3

X

y)    X"

i/j  X-

1/5  CC

<

i/)

i/1

vO

u

X^T

X

iff  r

in  u.

;X

X-

m

iX

-J  ;/)

<r

•-H

a

<r

!-•

z

c

o

z

»-

3

►-

UJ

o

z

UJ

X

#-H

»-

k-)     ^

</■

_'

U.I

«/■

u

X

^

^      •-

c

c

c

c

c

t

U

*~    'S

c-

c

c

c-

c

C

■>

u

c

c

V

c

L

r>

u.

*  a:

a

a.

a

a

0.

CL

«

Ih

(/)

(T

•

<.

2

>

o

•^

-J

»-

U

<

o

.->

1^

i-i

-1
-i

UJ

CD

<

•
•

a

t-

X
1   Ui

cr  cr

O  UJ

a  X

tr

UJ  —

•
z  o

<IUJ

2:  cr

u.

«o

tij

^

•-4

•-•

0.

*-«

tn

cr

OJ

IT

»-

•-•

►-

♦  m

•-«

L>

>*

o

(\1  z

1-  X

1— 1

X  ct

o

cc

a:

X

>-  X

z  <

nj

»

00  ac

o  o

a

a

cr

o

o

—  cr  »-

2  _l

1-

„_

OOO

o  cc

OL

o

ixi

a:

cc

h-  <  (/)

►-  •-

z

X

iT

cr  a

»

ex

»

a

tr

^  3   •

lO  (/)

2

h-

a  a:

a  UJ

X

a

cc

UJ

UJ

UJ  O  (T

»    »

»-

_i

•

V)

O  UJ

•   •

{/)

UJ

l-«

•

»

2     •  <a

>-  _»

tr

UJ

•

o

a  •

X  0.

z

»

a

X

a.

•-(

21

UJ  o  >

<;  o

••

I£)

•

z>

a  X

<  1-

a:

a

>~

>-

»—

2

«

»-  X  _1

tr  o

U.

<

U)

~»

uiz

o

2:  I

o

►-

\-

cr

X

UJ

(T

<!(/)<_>

cr  o

UJ

_i  cr

U.

•  o

o

2.   V3

^^

»-

u.

<s.

»-

e>

v>

O

»-  a  o

<  z

tr

»

UI

U.

X  _I

►-

C  ^

<r

UJ

U,'

»

^

»-i

e^

c

(T     •     »

»  •

»

irx

o

c  o

z

Kj  a

•>

a

_»

O)

u.

a

c

cr

~--  ir  (/>

U)  Vi

</?

tr

>c

U-'  e>

OL

2

»   »

a

»

»

tr

»

»

ct

a

UJ  t/i  tr

IT  I/)

(/)

<x

tr

<-

»  »

<r

»-

z  z

a:

z

z

<

*~-

z

Q.

2    <I  <I

<  <

<i

_l

_i  <

►-

z  z

_i

V)

c  c

UJ

c

c

-i

c

C'

•

_i

<I    _J  _'

_i  _J

_l  UJ

o

1

t/)

c  c

IT

r-

C    CI

D

Q

c

D

o

u

<r

^     «_  t

c  c

c

a

y

c  a

C

U

»     »

-?

•>

••

•

•

•

— .

►-I

•     •

•      »

«

<r

^

a

7^

»  »

a

■^

<3  UJ

J5

UJ

UJ

UJ

O

UJ

a:

1-

K    3  C2

o  o

UJ

_J

a

1

<b

a  UJ

cr

UJ  z

■z

•—1

-7

7T

^

z

—r

UJ  z

z

z

»-

»-«

Z   UJ  UJ

Ui  UJ

z

»-^

21

»-•  *'

UlZ

<.

z

K

■i

UI

^

=1

<

<x

?"

;r

<

tr

v:  <

■£.

<

u  u

o

CA»

O

U

o

O  tj

o  u

o

o

e;

t-H

♦-    U  U

O  CJ

u

IT  <I

fcC  o

ou

U)  o

O

X  'Ss

t/>

<

X

I/)

X

I/)  CD

X

t/)

X

(/)

X

tn  X  X

X  X

X  UJ

<

r

<  !/)

X  X

{/)

(/)

X

u.

tr

UI

o

o

U)

tv

o

o

»-

X

o

1-

1-

«/>

tr

o

C

z

O  2

z

— 1

<

^

C

UJ

C    5

■5

^

_i

c

D    1-

♦-

K

:«

a

«

*

♦  a  t/^

LO

♦

(/)

«

75

ac

>-

•
•

4

2

_i

■

>

«

jj

Jl

-J

£

s

^

u

<

»-

Q

O

-J

:d

1

c

t;

cc

a:

H-

UJ

•-N

•-•

a

or

c

»-

z

a

a

a:

a:

l/)

UJ

Z

<t

o

1-

Ui

UJ

rt

<

a

U.1

o

u

»-

£

H-

tn

i/>

if)

UJ

X

Ui

a

•-I

►-^

a

UJ

<t

>

>

a:

X

UJ

to

_J

t

z

»

Vi

«r

""*

U-

a

X

I/}

a

\-

X

t-

«-

<■

I

a

X

••

II

Ui

a

3
o

o)

»-

a  -J

C  3

oa

o  a

UJ

II

oJ  O

a

»-

o

<I

rHC

(V.  c

(/)

UJ

o

X

«

z

z

—■  u.

•-  a

•-  a

«

c

*

a

U>  Uj

5.

«

UJ

Z  X

z  a:

Jj

a

o

»—

*  o

.^i*.

s  ,u

s:  ^

-»

a:

fVj  «

-<■

■■Zi

X

y)

.».

a

•-   •

1-   •

OC

o^

CC«

?■

^

(T  »-

♦-

J^  i.

Ul   i

«

^•^

1— 1

a

o

o

5-J

*-4     •-^    jj

(/)

-  D

••  3

■X

I

a

a

o

<i

sD

o

-I          CO

o

a:  z

a  z

o

fVi

o

a

O  fV

f-

a

T

»-

t-

_J

u  a  *

o

<   4-

<a.  «-

a:

i:

o

a

1/5

a

z

(/)

Xi

o  o

■>-

X    •

I    •■

a

a

a

a

a  a

UJ

»

z

»

•'-'

t-  a

_J

uz

o  z

iij

o

UJ  a

UJ  c

Z

•>

»-

X

o

t/)  a  ~

>

♦  c

*  c

•

a

»

u.

••  a

<I

_J

1/1

UJ

»— 4

••  UJ  a

c

•  ZJ

••  z>

X

OC

a

»

X  a

o

UJ

o-

a

^

X  »  a

vT

2:   •

(VI

^  •

in

UJ

»-

X

c^

Z  UJ

t/1

n

QD

OC

u.

<a

OJ

«  X    UJ

4-

c  >

1— t

o  >

O

•>

X

t-t

rn

o   -^

*

•— t

<.

o

_l

<■

^  »-i  —

r—

i:  ^:

a

»-

z>  ^'

_l

>c

o

2.

a

•-

_f  X

1-

_J

a

o

»-

>.   2.    »-

KT

•  <

_)

2^

•■  <I

o

C

»— I

U'

U-

_i

2r

c  c

T"

•

c

c

UJ

a

C   U-    _

u.

»

>  »-

o

5

>  K

c

U."

at

(/.

^

Q

j_

e-  u.-

a

?

IT

a

a

c

•

H-

o  c/  u

z

J"

z  c

c^

h-

Z"  c

»

»

•>

»

•— 1

O

»-

•  »

a

1-

(T.

a

o

»

a

U.'

•■  •■  c

H^

K

•— 1

<i  Z-

t/)  (/■.

«a  z

z

e

a

z

z

_i

C/)

tr

~  ^

X

»/)

<

UJ

3r

z

a

^

z  z  a

_J

-^-

C

•    •

a

•

»    •■

c

c

a

/—

c

;«

iX

»

c  c

u.

»

_l

•

c.

u.

L

C  C    K

>

i

i    y

3

t

y  N

Q

c^

>

c

c

L

~

u

C    0

»-

u

t

L.

»•

c

c-

Q    C     h

u

►-

'.''

lA    (A

»r

"~  ■

IT  ir

»

•

b

»

»

'y

ir

— ■

»     »

tr

o

V

;^

KT

»

»-

-^

Q

•    •   "er

tC

C/5

o

«1    <I

a.

<3-   -C

f^

UJ

3

UJ

a

C  U.'

a

O

a

a

UJ

£C

cs

3  UJ

»-<

?"  2"

[

»-

5.     ..

Z    U:

T"

_]

UJ

2r

<r

I

K

—

u.  z

_l

I-

UI

H-

I— 1

z

"T"

t— ■

_i

U!  .~  -a

4

1—

yi

_j

■i'

_J

>.    1 —

;

•C

_J

a.

►—

l/l

_1

-L      h-

^

u-

e.  o

c

<_

CJ  O

O  CJ

O

<i

(_>

o-

UJ  u

<3.

e

o

o  u

<

o

u

O

c

a;

»/.

d

<_     O    LL.

u   u.

uJ

<

X    X

(_>

r

■C  'C

■u/l  "E

X

O

X

X

LO  a

C_>

X

in

^    -y>

O

--*-'

r

X

t—

.X

^

O

X  T  I/)  t/>  a

S

<.

J

^

»—

CM

n

u

u

o

«-

c

«-^

•—1

»-«

f^

•— t

M

m

(V

5t:

tv

o

•— t

►—

t-

»-

1-

»-

1-

K

♦-

1-

c

K

♦

_!

_i

ir

^

^

~"

A

~-

<^

2.

mL

c

o

c

C

<_

w

5

5

J

2

Zi

i

ii

2

_1

<r

«-

C'

o

<:

k-

»-

t-

H-

V-

l-

»-

1-

^"

K

K

«

tr.

V)

tr.

(/)

*  1/1

</1

tr

4e  (/)

♦

*  «

c

*

I/!

to

76

UJ

X

<

z

a

>-

a.

a

X

Q

4-
»-«

4-

•

-1'
UI
CD

z

UJ

o

«-
1-
«-
o

UJ

o

z
o

z

1-1

z

UI

X

UJ

-J

«-

.^i^

z

a

o

4-

UJ

UJ

O

4-

Q

>-

UJ

I

UJ

•->

-1

UJ

Z

o

z

4-
X

o

X

<

Z  UJ'

o

se:

4-

«■

>

<  X

»-

»-

•»

z

u

z

o

4-

«-

C2»

«

-J

«  UJ

t-t

o

1-

O

a

UJ

.

CD

a  a:

UJ

lij

a

«*

r^  <i

a  a

*

QC

(T

oo  ^

oo  -i

UJ  X

(VI

♦

n

oo

J.  UJ

r-

«

«

o

o

OL

a  <

•-I

cr

♦-

•-«

»-

n

o  cr

o

»-  >-

o»

oo

o

o

o

a:  o

cr  o

(/)  _i  a:

>-4

o

a

O  1-

ct  oc

a:  Ui

•  z

o

or

t-4

o

u.

»-

o

»-

o

UJ  a

UJ  a:

a.

{VI  o

q:

o

U.

•-

QC

¥■*

(/>

ec

v>  o

•

Ui

»

c

u

Qc:

a

•—1

o

O

\-

»

(T

•>

»-

oc

»

o  «

»-

•I  *

Ui  QC

1-

©

a

i/1

-1

Ui

_J  c/)

<

>

>-

(Tl

»-

•

Ui

(A

t-

q:

»

ou.

»

UJ

»

>

<I

a  r^

t~

u)  a

a

•■

•

W)

UJ

X

o

UJ

•X

GDU.

-ICE

K    r-l

(/)

o»

►-

X

X

»

»

o

c  tr

»-«

<  Ui

o  a

z

•

z

X

Z

UJ

X

X

►-

2

»

2.

_i  or

o  <c

UJ  a:

X

UI  DC

o

Ui

(/)

o

c

•-H

o

•■

\r.

Ui

UJ

m-

»

»

•>

•  o

U

X  C

»-•

X

_J

e

1-

c

e

<r  iT

</)

z

(/>

IT/

(T

V)

<r  a

»--

1-  a

cr

»-

u

»

»

»-

»

V)

<t

»

i-l

jr-

to

tr

XTt

(o  a

»

•  a

»

»

1-

»

»-

z

"ZL

o

2^

<i  _J

z

_J

^  <

<r

<

<f  UJ

z

O  UI

z

z

z

z

z

c-

O

e

C

i-

f

ij

c

:i

_i

_l

;

_1

-J  •

c

K

z    *

c

c

2-

c

:r

o

c.

K

0

^

t

o

u

<--

C:

C

C

t  u

Q

^

c  u

Q

Q

»-

o

K

»

»

(T

»

?'

•

7

9

3-

»

»

»

•

•  ^

»

UJ  o  r>

»

»

t/>

••

U>

LJ

UJ

UJ

»-

(jj

'v5

UJ

o

o

C

o

a  a:

UI

"s:

cc

C

UI

UI

z:  »-

*-.   <t^

o

~

in  -

z

►— 1

z

"T*

<i

Ui  UI

UI  Ui

UJ  1-

z

UJ  _l  1-

Ui

z

_l

z

_!•-«

7^

».^

►-

<

IX)

J

r

H-

»-

_J

-I

_l»-

O  2

o  u

c

o

o  «-;

O  IT/

o

o

UJ

Ut  o  o

o  o

u  o

o

<t

<  o

u

o

<

o

<a.  Ui

■XiK-

(/>  en

o

X

«-  (/)

X

<

I/)

I

lO

a:  03

X

oc

X

OD  n

OD

t-

U  X

X

cc

o

X

CJ  QC

(/5

<-

•-I

n

(VI

cn

QC

O

o

o

c

o

o

u

a

«-

»-

<-

»-

t-

1-

,

<-

Uj

fV

^-.

c  c

c

c

c

c

o

u

^

u

u

u.

c^  •«-

CI

<-

ej

e:

c

»—

4-

♦-■

»—

»-(

1-

»-

K

1-

1-

1-

K

K

1-

(/)  *

«  tn

*

(/)

cn

(/>

«

*■   V)

*

(/)

I/)

i/)

77

UJ

q:

a

Ul

UJ

I

I

cr

a:

<

■"■

,»

•"~

UJ

-^

a*-

>

>

z

u.

>-
z

cr

uJ

>-
z

I

•

4

D

4

r

4

xl

D

•

r

UJ

z

,»

r

X
UI

UJ

3

z

o

«-

c

o

X

4-

^.

o

_)

c

ct

c

_J

c

K

4-

H-

c

»-

«»

c

«-

O

UJ

t— t

z

UJ

1-^

UJ

or

(/)

4-

(/)

UJ

■yj

a:

UJ

o

T

♦•

cr

:»

<

I

<r

4-

<i

o

ct

o

UJ

X

«

uJ

r

4-

T

_J

I

UJ

r

-4

X

u.

a

4-

Ct

4-

or

UJ

4-

X

z

»— •

*

<n

q:

X

CL

X

•
>-

a

X

4-
I

a.

X

-1
4-

"^^

a:
Qc  a:

o

1-

tx

o

II

UJ

4-

UJ

4-

Ui

•— •

o

UJ  a

Ul

z

o

a

X

«

2

4-

•— I

UJ

a  X

cc

u

a

oj

_l

a
a:

a

uJ

o

UJ

*

4-

fVi

4-
»

a

X
4-

UJ

a

O  tu

♦

X

O)

I

aj

a

•

a:

cr

O

4-

r>

>-
4-

cr

4-

c

»

1—  >-
if)  _J

•-4

o

3.

2:  H-

Zj

*

»—

at

X

o

-♦

•  z

o

3  O

D  in

2

cr

4-

o

4-    _l

a

3  »-

,^  o

IX

UJ

\~

o

z

^

<-

tVi

o

a:

>-»

o

uoo

cr

I

o

o

ar

\f\

»-

<-

3

»

!X.

O

>-

j:

e  X

cr

a:

»

o

2    «

c

ct

»-

»

(/-,

»-

«—

X

Z-

a

a

U.  OL

U!

Uj  3£

a

X

•X

C

t-

a:

3

a

»

2

_l

UJ

Q

o

UJ

a

O

ct  o

U!

Ui

(X.

C  flt

tr

UJ

a

*

<i

>

n

>

o

3

a

»

c

•>

^*

•

«

t

UJ

»-  o-

•

»

H-

>

<

Z)

C

z

»

a

a

«

1-

a

«  >-

a

»— <

•

(/)

X

X

tr

_i

cr

»

^

»— 4

>-

UJ

y-

(/)

»-

tr

»—

X

a

ct

a

•-I

•

(V:

u  cc

>-

4-

z

»

I

m

»

X

»

X

IT

S-

H-

z  c

o

2

X

.— •

o

<

_^

»—

a

X

«3

X

^z>

t— <

X

o

>♦    X

cv

r-

»

U.

UJ  cr

\-

Ll.

u.

c

•

»

»

<t

t/^

_>

»-

»-

C

•— i

c

»-♦

>-  V

►-I

z

U^i

X  a

V

tr

z

/—

Q

lA

tr

K

»

a-

»—

c

UJ

a

»

♦-

or

•  tl

0-

»

c

1

»-   UJ

*

»

»— <

»

Cl

tr

IT

C

^

u

3

li

»

»

a

»

»

a    •

^

a

Q.

»

»   »

z.

T"

_j

^

u;

<i

<.

z

•— <

t/)

»

~z.

q:

z

a

z

a

-^

Q:  Z

Ct

z

ct

»

z

a  UJ

c

H-

c.

X

f^

^^

•  '

C

a  a

^

c

a

c

u.  c

a

c

U-  c

Q

c

b

u

c

"^    "~

a

<■

f

t

r

c

c

V

""

c

V

c

V

c

c

>

Q

r.

>

D

^

C!

c  c

•

?

r

i-

»

T'

»

•

cr

■"-■

tr-

c

<•

•■

U.

1    •

z

•

U'

•■

,*'    •

u

»-

7"

•

1_>  h-

UJ

w-

U.)

UJ

o

o

UJ

<

iT

-J

tL.

UJ

f^

cr

UJ

C

O  UJ

r*

rr

u

UJ

7-,"

tn

^    '-

<3

"^

K

^  •— '

U.!

-5»

«'-

f-^

_J

z

Jl_

_J  Ui

t— 1

T-

_j

UJ

>-i  -^

_)

U-

»— ■

a

r^  r^

_l  u

1

7    k—

a  tr

Cl

_i

iXJ

xj  -;

_J

CO

-J

uo

_)

(//

<i;

_»  X

o

Q.

c  o

U.'   U^

c

2.

o  u.

o  o

<J

l/>

<i

^

X  o

o

<.

o

t/:

c;

^

e

tr  o

<I

O

tr

c;  (_

<3

T;

<-
o

(/>.r.

•JO  :i

X

u^,  <

x

t;

X

^

u

1— 1

1--

o

UI

X

u

X

<

X

u

X

«    X

u

X

<

»-•

LO   X

a:
ct

a

K

O

ct

a

c

■>

X

UI

c

«-

«■

u

t/'

y-

X-

X

*•

♦-

t/^

c

c

»-

c

c

c

C

c

a

«-

c

«-

►—

^

c

c

c

Cl

»-

H

H

»-■

1-

H

t-

ir

«

*

to  *

*

(Al

1/5

KT.

tr

tXi

78

tr

ui

X

a

»-

a

«■

X

(/)

«p«

Ui

4-

o

»-•

>■

^z

<-

z.

o

-1

«

c_>

«-

4j.

z

z

XX.

UJ

o

►-I

c

X

o

Ul

,

UJ

»-

X

or

.••«

bJ

1-

UJ

UJ

Ui

Ui

«-

«-

X

(n

r

UJ

I

>
z

UJ

Ul

4--

Ui

X

<
r

^

or

UJ

'"•

<

z

*-*

X

<

1-4

<-

^.^

X

X

X

o

_J

■»

a

J

a

o

*»

UJ

o

3t

«■

4-

o

X

tr

o

1-

UJ

2

o

UJ

XX

Ul

a:

UJ

o

UJ

♦

«  u.

tr

a:

O

a

•

o

«

•

</1

4-

X

CE

St

CM

^

z

Q

*

4-

UJ

»

»-

in

(V

(/)

4^  3

X
UJ

*

a.

X

Z

w^

<-

o

•

Q.

_J

_1  O

>*

o

a:  —

o  t-

UI

lo

a.

o

X

X

»-

4-  »-

X

a.

o

(T  Z

o

X

«-

h-

•»-'

1-

o

o

3

3

o

X

ex.

ac  o

X  UJ

o

o

>-

i>

3

X

X

X

3  X

X

o

Ui

OL  u.

UJ

c

K

(T

Ul

o

a

o

a

a.

X

1-

UJI-

X

X

•

uja

»

z

(/)

q:

o

UJ

1-

UJ  1-

UJ

UJ

tn

X  \n

UJ

X

a

»

X

►-I

1-

»

UJ

>-

cr

(/)

Qc:

(/)

»

»

••

»

'»

UJ

1-

a  ♦

»— 1

2

X

Q.

•>

u>

•■

»

a

X

a

X

♦   X

X

•

X

t~

2

2:

o

O

X

»

«

X

9

a

H-

H-

o

Ui

t-i

•-

a

o

X

UJ  «

H-

z

o

•-4

X

Q.

»-

I

X

2

o

-«2:

(/)

»-

♦-•

ts  ^

</)

V)

UJ

_J

X

»-

o

»-i

t/J

C

u.

O

cr>

»-

«a

-i  UJ

►-I

u.

a

»— «

»

••

o

u.

K

r>

»— t

:^

a

cv

UJ

•-I

•^

3

a

•  </)

_J

Uj

a   •

'2.

K

C

2

1-

c

IT

2

a

»

\r

U.'

»

_J

a

a

a

»

x    •

»

_J

m

•  cr

Cj  -z

^

c

z  »-

' »

U!

»

or

»

♦-

a

»

»

»

1-

2

X  z

z

»

2

z-  cr

Q.

UJ

2

CL

r

t/)

2

c:

z

X

2

Ul

cr

z  x

z

z

1/5

O

UJ  Q

o

z

c

c  u

»

c

«-

»

►-

»

C

»-

c

u.

C

s;

u

e  a

c

o

•

a

a

a

c

C

a

li

^

U

c

u

b

C

~

Q

Q

<

D    >

Q

Q

U

•

<:      •

•■

o

m-

••  .^

-y

*-'

«-

u-

*3'

»

o

»

K

?

»

X

?■

•  U-l

»

»

— '

u^-

cr  c

U'

»

UJ

ecD'

•— 1

a

lU

UJ

zo

UJ

<

c

UJ

UJ

O

X

2

t-i  UJ

2

Ui

^  p*

UJ*—

O

«

Z  _l  o

_i

»-

^

*^

<r

2

s:

»-(

2  2

a

2    HH

2   _»  2

2  UJ

v~

2  W

2  2

et

I/)

a:

<^

.-J  M

X

_i

<

T  (-

►-  (/)

<

<  </)

_»

«I

c_>  <

to  o

O

<t

o  o

UlA

U.'

o  »-

<i

u

u

(_>

UJ  u

U

t/^  (At

U  O

a

o  tn

O  <

o

o  u

o

X  o

«a  X

X

CJ  CJ

X  tn

a^<

I/)

to  z

ux

(/>

X

t/l  CE

X

<I

tn  OD

♦•

to  <x

CD  O

CD

lO  X

OD

O)

</)  X

»-«

3

Q.

>-

*~*

UJ

a  ♦-

»-

f*

51

C

o

«-

<-

CD

ox

(/)

I

UJ

Ql

O

z

3

ic

a

z  o

_J

3-

</)

C

_l

Ui

»-

«■

*

K

K  »-

\-

C

c

c

c

c

^

a

\r

"

3  3

^

c

c

_J

o

c

a

♦

<-

a

a  a

X

t-

»-

>~

»-

*-

1-

»-  ►-

►-

(/)

(/)

«

t/1

tr

♦  {/)

«

•

(/i

(/)  t/>

tn

79

^

X

UJ

r

•~

UJ

a

a

UJ

c

;±J

a

-I

r

UJ

a

»

r

1-

••

3

o

••

a

UJ

C!

1-

a

UJ

o

{/)

a

UJ

»

«

a

_i

«

5«

V,

<  --

y)

-J

3

^

a»

rvj

.^

vO

«

*■

lij  a

»<

o

1-

a

f— •

^H

«*H

»-

a

a

•-*

V

h-  t-

—I

>-

»-

^

_i

r-»

■4-

uJ

a

►—

o

_i

-^

«-

Z3  CO

a

D

13

a

H-  a

a

CS

o

in

»^

»-

a

V

a  C!

K

a

O

»—

o

o

_J

a

a

»-

a

o

1-

oj

/-^

a

«-

K  ^

3

►—

»-

(/>

a

X

a

_J

o

o

1/5

a

Z)

li.'

;5

a

o

>-

u>  —

a

tf)

{/-■

»-  a

a

«a

Q

a

»

UJ

—'

e

t-

a

a

<x

»  _j

H-

»

»

z

I/)  iij

UJ

u

a

a

a

••

_i

»-

1/5

U.I

a

ex

>-  >

tn

z

>z

u,

»

»

»

»-

Ui

li.

<

V

>

(/)

»

»

u;

cr

<  c

•

~

►—

I

X

a

X

IT

•

»

>

<5

o

»

X

a

*

•a

a:  vo

a

a

cc

»-

<  h-

»— t

»

X

a

_j

a

«c

a

c.

<.

►-

X

a  «-

♦—

K  o  a

»

z:

I

z

X

1-

t-

u

a

<-

►-

>-

r-4

s:

I

•-•

•

o   »-

u.

a    »

c;

a

i

to

UJ

»-

O)

u.

a

<i

»-

u.

a

»—

>

^

3.

Q

•  tr

u

•-  u-

V

a

c

»-<

iT

u

\u

»—

u.

•

••

tr

u.

t— i

u.

o

•— 1

L.  U

«

ir   •

a  (A

if

X

c^

a

•>

^

c

_!

)

IT

lA

•

_j

a

11

c.

Q

{/^  Z

T

IT;  r

»

O  <:

\r

li.

e

•■

T"

^—

•

»

»

KT

1/5

r

»

o

K

»

»

•  »-«

O  ««   «-•

^

cr-  _»

<i

t-

z  z

c

_J

-7

Z

z

<

<

•— •

z

1/5

tr.

2^

Z

Z    _J

_;    C

c.

a  (_

"-

c

c

a

3^

c

c

c

_

;

c

c

X

•

c.

c

c  s

»"

c

L

_     •

c

c

D

c

•

L

c

r

c

C

t

c

L

ID

c

C    L

— ^

•       "5

•

(/■'  '^

»

»

»

U'

7>

»

>~

V

»

•

•

?■

•■

tr

"~

»

•

»     y'

o

li.  'S>

aJ

uJ

UJ

►— <

c

'Jj

z

UJ

T'

LU

U.I

^^

UJ

o

Ul

a

!3:

UJ

UJ

~  rr

"^  •— •  I

"    —

_J

";*

<-

_;

ll.'

-?

<

~^

T

■;?•

-  z

^

\x:

A

1—*

*;

z

_J

►-

^

U.'

"~~

z

"  <

4         -A

/,

I

_t    sJ

■^

-J

■'-*  ^

2^

►—

k—        ■;

-^

_2

lA

a

_1

a

c

;/  K

d

c  lA.  <_  e

«     a

O

<-;

=1

u

o

c;  X

U

UJ  o

tr,  o

O

o  <_

u

o

o

(/^

o

c

<a

u

o

o

o

o

o  u,  u

{/)     5

X    c  in  I

;_>

r

i/>

U

X

ou

y)

1/1

or  X

yi

X

t/)   X

1/5

X

r":

■i

in

X

u

X

;/>  X

X

\r.

X.  to  a

C

2

»-

1-

a  X

e

2

cr.

X

-J

13

4-

a

o

_i»

*-

c

X

liJ

c

O

UJ

_j

a

r^

»-  «

~

►-

»-

K

►-

>-

\-

4-

•-

1-

\-

r  «

c

"^

~

~

—

~~

u

C

u

u

u

a  *

K

a

a

Q

a

Q

tr

«-

e

ere

»-  «

(/

t-

»—

K-

»—

y-^

K

»-

K

1-

U)  «

«

U)

(/)

u~.

t/)

in

tr.

«

OJ

VI  \n

80

O

a»

z

»-«

a

</l

o

u>

in

t-t

*—

!»-

z

ir

«

«

>*

«  so

QC

tM

s

O

in

vO

a:

cr

t\t

•F4

(V

cr

<— •

>-

o

a:  o

a

m

.-t

r>

o

1— •

•■4

a

-1

o  a:

UJ

^H

at

1-

(M

a

>-

OC

1-

_<

_i

a:  tr

•

cvio:

-I  o

oc

UJ

V  a:

cr

cr

o

oc

z.

<i

<

IX  Ul

-I

-J

o

_l

a

o

OL

oc  o

UJ

»-

cr

o

Ui

i->

-^

lU    •

«3

_» a

<

a

cr

t-

»-  a

»

z  tr

a

\-

9-

oc

•  a

z

<.  cr

o

UJ

q:

U)

z  a

XI

UJ  UJ

cr

(/)

_i

Ui

_i  <

a.

u

UJ

»

»

UJ

»

UJ  UJ

t-^

»

•

UJ

•

u

K

UJ  >

UJ

»

»

X

Q.

»

X

»    •

cr

X

(3l

»

X

c

X

□D  _»

H-

X

Q.

<

1-

X

>-

X   Q.

►-

<t  t-

X

-1

c

UJ

•-•

<  O

K

»-^

h-

3:

X

•—1

cc

•-•  »-

1-

T

I

•-H

-J

^

•>

c

-1  a

UJ

s:

Ul

>.

o

i

»-

H-

2.  U.

<

2l.

<S

2.

<r

»

O)

-J

•  »

*

u.

UJ

o

»-l

u.

UJ

z

^

U.I  UJ

»

o

»-•

UJ

lU

O  1-

IT

w

t

IT  (T

tr

t/v

_J

o

a

(/)

^

UJ  UJ

«/l    _J

tr

tj  a

t/)

2

*  z-

t/)

<

<

tr  cr

U5

•

•■

»

»

•

»-l

■>

S

»     »

</>

»

»

»

»-4

-z  UJ

<

_J  o

<  <r

<

2

T

ct

z

z

z

_J

2

UJ

z  z

<

z

z

2

-J

c  r

;

o

»

_'  _J

_J

C

C

a

c

o

Cj

3?

c

K

c  c

_f

c

c

C?

3

o  u

c.

u

c  c

c

c

o

>

0

o

o

u

Q

<"

0     D

e

D

c

Q

U

»  •-

»

7-

'

•  •

»

»

v-

U.I

•■

»

»

:^

•

»-

»     »

•>

»

»

•

2

\±i  <

liJ

r^  a:

C5  C?

UJ

o

UJ

c

UJ

UJ

UJ  V)

'3  UJ

UJ

t3

UJ

UJ

z-  »-

z

3

»— 1

•-

Ll'  UJ

Z"

rf_

UJ

^

_)  UJ

2!

27

;^

<3

~?

7"

UJ  z

z

IZ

z

LlI

z

■y

z

<r

U1

-i

tr

'a

_J

1

r  »—

V

<I

-■3

■d

<3

s:  h-

o

{J

u

tr

o

o  a

o

u

o

o

•d

o

O

tJ

o

Uj  UJ

o

<-

O

o  u  u

o  o

t-*

u

tJ

(_>

UJ  li.

X  _J

t/)

OD

<

X

03  I

CD

t/)

u

X

u

X

Tj

lA)  03

(/)  cc

a  a:

(/I  CD   Xi  to

XI

(/)

CD

2D

t/)  X

in  ac

4-

<-

_l

>

1-

_l  «-

•-«

or

<-

r-»

(VI

-J  <I

O

c

•-(

fV!

1-

2-

>

>

<:   «

!

.^

^

_;

21

«-

a

Q

o  o

_l

_l

_J

_)

u.

U-1

H-

»-

t-  ♦

c

<I

<r

<-

1-

<-

ii

^__

C/1   «

t>

u

o

O

</)

«

■

UJ

UJ

81

<
I/)

2"

Ul

X

UJ

►-

<

»-

^1.

m

2
0

c

1-4

<

»-

cr

<.

«

«

or

-J

a^

u

)£.

>0

ac

;m

tv

IT;

0

0

0

>"VJ

>-
<

IT

<
>

IT

JC

m

-J

a:

a:

_;

»-

o

r->

a:

X

03

»—

0

ac

z  a:

0

tn

tr

a

Z

a

c

»-

Ul

•«—

(/)

a

<

•-I

*-•  0

C

_»

0

a:

a

c

oc.

z.

•-«

Cl

_J  x

q:   •

S

Cl  a:

»

_i

q:

ii.

»-

a

a

2

u.

a

t— t

_/  u.

;  „-.

»-i

♦  a:

.OS

^

q:

»

Ui

or

iij

►-

•^

t-

^

u

•

oa

«-«

j2

^  UJ

ac

a

u;

X

a.

iij

•-

{/5

•-

1-

»r

Q

_i

(r,<i

^

»

>    •

<

•

»

z

9

»

a.

*

z

<x

<■

»

0

a

•  >

CL

X

0  a

>

X

a

cr

►-

X

a

»-

Ui

»

»-

a

0

»-

a  —

»

•a

—  »-

f— •  «^

.—

«a

H-

D

A'-

•— t

K

X

c

»-

cc

•-

0

I/)

»-a

or

■^

2:  X

»-  a

1-

2

I

►-

U.I

2.

U.

o

•-«

»-«

c

•

u.

-?

c

u  >

<3

2

»-i  0

u-  >-

U!

>

tr

U

2

L^

li.

»-

u

■»-

?

•

1-

u

•

""

u  l-

>

r— '

C

cv

C  f^

^     K-

^

0

1— '

a

li.:

IT

!

c

~^

Ul

»-'

Q

►—

_f

IT

*

_J  «£

■w.

»

C^

»

>f  a

e;  .£

<J

♦-  U

cr

»

K

»

»

»

t-i

2

_i

V-

T

»

iT.

a-

»  <-

>

•5-

c

»

y

c

♦    ♦

<r  <-

<

u   •

»

2"

<I

zr

-^

or

2*

-i

«5

y-

•-^

z

<J

<

Z  1-

D

I— •

z

z

t-l

z*

»-  z

Q.   t-

a

^  z

z

C

»-

c

c

a.

c

"y

^

u

<

_'

c

_;

>

C  1/

K

c

~

c

c

~

c/-  C

»   W

•

c  c

c

k.i

V

c

c

>

c

L

\-

c

e

c

Li

c

c

c

c

'xi

L

c    D

c

»

m

»

u.-

»

>

t—

0

'

^'

•

#

7

•  ^

tr

">^

T

•

■p-

T

->■     •

~>

Q     •

»

UJ

7^

3

UJ

UJ

U!

0

IS

iD

c

Ul

'fv

UJC

vD

u;

a

tD  U)

cc  i

d

e

UJ

.r

<-

-j£*

Ui

«r*

_i

"■^^

2-

<.

U

yr

0

»— 1

»— 1

U'

71

t— 1

j_

7'  •— 1

<i

H-

_j

T*

HH

_J

►— '  ^*

'

I-  —t

»-

_1UJ

z  ^

.^

-a

_f

"-X

i\

^—  C'

■>_>

_J

LT

y:

a>

^

in

-i

t/;

_,

ir>

_l

<r

ii.

w-

_j

^

c-

<-

c

u

C--

«j

u

«_>U'

u  ct

<a

_i

tr  IT

c;

0

IT  0

otr

U

IT

c

u

IT'

<I

I/)  0

u

c_>  a-

0

<i  0

(_  u

E

4-

■r

-T

.X

o

-

(/>  i/;^

a  X

CL

<

<i

<i

cc

X

1

IT/

;I3  H

in

<

0

"C

<a

u

^  .x

:/)

x  <

r

u  r.

T    J".

<-

i-H

»-

1-

u.

"

cr

U

♦

Q

<_

IT-.

c

_

T

a

f-

c

C

T

t

^^'

►-

*

U

4:

c

»-

c

c

tr.

♦

lT

UJ

*   Q.

Cl

ir

/^l

82

Vi

U)

t-

o

¥-1

a.

cr

u.

'

¥-

u.

_;

^

o

-1

«-  —

J.

^fD

z

C/)  t-i

o

•>  _l

u

•

1-  c

CD

if)  o

t-t

o

a

a.    •

h-

-^  aj

>~

Z  •-H

<

UJ  a.

t/1

>0  H-

....

•-•

1-

nj  z

t/)

n

•

H-l

<>«

»-  ^

•>    r^

n

-1

I/)

oo

UJ  _l

a:  0.

f— 1

IT

^^

f^

\-

tv

5t:  >-

t-   21

>—t

sf

a.

(X.

.^^

»-

(VJ

Q.  -~

o  m

h-   Ui

»-

o

Ct

>-

o

a.

•-«  s:

tr

<

<I  •-

UJ

q:

a.

o

1-

2

c^

!-•

oc

cr  •-•

0

a   •

»

:<:

o

a:

a

<l

:d

0  —

X

0

0  3

cr

•  ►-

t-  a.

a

tr

UJ

ct

»

o

«~

2  z

.».

CJ

a

(/)  ^

or.

CD  •-»

•-I  o

<

a

»

^•^

UJ

•-

cc

2:

CD  *-*

3E

U)

cr

CE  X

Ul

•-^  oc

2  a

a

UJ

»-

h-

»

»-*

i-i

»  0

t-i

(C

UJ

r>  •-•

»

a

->  a.

»

»

I/)

t-l

a

z

>

0

X    »

0

r)

»

(TO

»— 1

X

»-   iC

-i  UJ

K

a

<«;

2.

t-

r-t

•—1

<a

»

2>

»

(/)

a

—  >o

a

<

1-  (/)

•—   •■

q:

»-

o

•— t

X

K

_l

cr

>

0  CD

>

~-

t-

UI  <-

3

X

<3.    <!.

f-    o

«3.

u.

u

_l

t5

UJ

•—■

tr

0

_!-'

0

UJ

u.

X  H-

10

X

•  s

(T     •

T

Ul

»

•«-

»-t

it

h-

<r

--'

C  C

^^

2

3-

u>

<i  to

»

0

(T

U_   i-'

C

_i

o^

:?

or

u

Uj

«^

CD

<_  2:

C

<r

""*

_j

2     •

^1

0

C/1     •

1-  a

»

»

tr

UJ

»

<

c/^

UJ  2

t3r

►  CC

2

2

2

»

•s

q:

»

»

<I  cc

or  r

21

^

<r

_j

:?

a

cr:

X

CD

CL

2  0

CD  a

2

»-  ID

a  X

2

_)  »-'■

»-  u.

C

c

_'

»-

c

9-

K

<l

1-

x^

c  z

1-  Q

a

UjC

a  -2

X

■"^

c;

c  c

K   1-

C

c

c

L

o

U

1-

^

t

I

0   z

U    >

■

1-

0

*-

ll

i"

Q

»  K

<r     •■

»

•■

•

cr

»

^

<r

tT'

»  0

(/••  U'

a

<»

»

□r  ?•

»

»

UJ  H-

UJ

'jj

UJ

UJ

UJ

;?  or

o

»— t

UJ

u

OUJ

0  0

>— »

e

Ul

r"  <t

<  z

zr  -F

7-

■y

•^

<

;t»

Z"

<

»-

<r

t-H

u

4^-

<t

-J

2  <

<r  _j

•-•

0

10

0

2

-!i

(/)    HH

_J  UJ

"3

*-

■5

C

T

Sr

{_>

^

>—

o

D

21

_J

^

S   _J

»—

0

•D

_l

<

OL  to

_j

u  u-

U-  o

o  o

o  o

o

UJ

<J

o

(/>  u

Uj

UJ

a

O  Ui

<;

OUJ

Uj  «a

UJ

a

3

_J

u

u

3  V.'

<i

0

U

X  t-i

to  iC

X  (/)

X!

I/)

x>

to

to

X

M^

OQ

VI

ct

a.

no

I/)

u

xto

(/I  0

cr

CL

to

<

CD

I/)

IT  <  0

OD

OD

^ll_

•-4

(VJ

»-  li.

»-

1-

Li.   C

u.

\l:

>-

i    K

ic

^

c

0

C

Q

O  IT

r--

CJ

«_)

c

c

0

c

c

»-

<r

•x

<•

«r

4^.

D

-c

i'

c

^

CL    «

Q.

a

a

UJ  «

a

OD

UJ  >»

a

«

Ui

83

o

0

»-

z

<

4

ac

a

Ui

UJ

a.

(L

o

O

•

u

(K

•-«

UJ

(t

Z

Ui

o

X     ■

z

z

(/)

•

o

UJ

xC

>c

UJ

r>

n

z

^m.

„—

q:

0:

2

2

«

0

0

c

C

a

q:

H4

»-»

>/)

•-«

a:

tr

;n

a>

fn

O

\iX  ^

LjJ

>o

^«

IT

UJ

^Hi

q:

^^'

z
<I

•  Q.

5.  »-

»

<»5

-x

X

3:

X

o

cc

ac

<-«

3  C/)

D

X

(/)

o

a

a.

UJ  cr

UJ

UJ  a

0

{V>

Z  Ci

2

(^

0

^

0

a:

X

X

a.  X

a

a:

0.

0

a.  Q.

«-  3

«•

a.  -•

JL

a

a

q:

UJ

UJ

c

Uj

c

a

o

»

0  0

»  ^i'

•■

C  fvj

a

IT

c

•«^

Uj

»

»

»

>c

a

»

»

2  _J

Z

•  Q.

UJ

a

»

u.

»

c;

o

2.

X

2.

Uj

u

<

»-

>-

o>

c

2  C

»

0  u.

u.

a

•-4

»— (

ZJ

»-

»— (

*

•-H

>

(/)

<l

z:  0

^

►-    •■

(T

»

»— 1

0

h-

a:

a

a

<»

cc

a

_j

2

CC

•  >c

»

CD  U)

0

a

(T

>c

X

UJ

UJ

o  o  a

UJ  u  Oct

>-  «-

>-

Q.  3

2

>-

a

«■

«;

TS.

X

^

•■

z

o

>.

0

0  <x

2»-

a

2

0  _l

»-4

u.

0

»-

•-•

;^

C

:;^

iro

»-

^.

»

»

•

<r  (A

_i

<■

•  C

5.

u.

•

(T  >-

Q

?

U.i

7

cr

c

tr

«r.

tr

cr  tr

H-     ♦

Q

»-

tr  ♦

♦

_j  i/"

»u-

»

c  a

•^

<

•

t/)

■*^

(/I

tr  ir

c  ?

0

C

</)   •

»

>—

»

tr<

2  »-.

a

Z

UJ

2^  O

c

U.I

_j

z

<

U.C

<t

ij

<r

z  »-

(/>

2

<  z

2

/-%

^

<

»-«  _J

a

c.

T

a

•

^

^

c

r;,

_!

■5"

I

^

_,'

•  c

a.

*

^  c

C

z

c

■•

c.  e:

>

c

<^

r

b

r

<

♦

c

c

<

c.

c

c

^

~

ii

c  r

G

c

0

e

D

u

•

~

—

/—

~'

c

»

»

~'

»

•

•

IT   -^

tr

tr

•   •

»

0

»

•

?•  <

UJ

t— t

CL

UJ

o

c

O"

nr

Ul

<r  c

<

UJ  0

UJ

C-H

Uj

C

_>

^

u

3r

_J»-

_)

3'

U.'

UJ  ►-*

O

^-^

UI

ll

~

5:  —

_J  •-

T

r  u,

P'

c

UJ

T^

•—1  1   •— «

_j

t—

^-^

■s

_l

_)

i^'

u

r—

r

-?

J-.

-I  H-

1—

h-

i/!   _J    >-

<i

o  u..

a

X

<i

O

<j

a

or

e

c  u

(>

a

0

C

0

U  (/>

<3.    U.

u

U.'

0  0

c_>

0

u

u

tr^  <  UJ

o

X  l:

X

^

(_>

o

c

a

X

X  :r

Ix

^^

■rs

T

x

:r:  <i

0  a:

ja

cr

L.I..    X

<v

T

3  c_)  a:

^-

t.

c

»

c

U.  Q

c

c

r-

a

o

c.

u.

«

►-

a

U'

^-■,

0

^'

r—

c

fV'

(V

f*

~ .

^

e

0

«

^

C

Q

*.

D'

a

a

D

Q

Q

Q

00

liJ

*  a

o

»

u

4i

C

UJ  >>

a

c

0

0

C

c

C

81+

z
<

o
o

tf)

m  ~
a

a:  I-

o  tn

a  o

ex:  ql

••  u.

«-•  c

CD  O

I-  4-

V)  »-

•  tr

tr  »

</!  r

<i  H^

_i  c

c

»  ^

UJ  c

T'  t-i

tr,

tj  (/)

r  <

Q
O

m

o
cr

C£

•

X
CK
•-
V)

<r  H-

tr  _j
••  e

c:  <i

n  K
•  «r

C3

UJ  _l  t-<

t-  <3  Ll;
:d  u  a

in

X

o

a:

(X

UJ

3:  z
a  c
>  c

Ui     •

UJ

a.

o

_J  r~     Z>

<i  <_'  u.  a
u  X  cr  X

IT
Q
O

a  ^

O  UJ  *

tn

z

U.'

:d

ID

cr

<i

Ui  T.

5  rr

<t  z

z

U.
»-     K

tr  <^
I-.  o

_i

.D

a  _*
4  <i

o

r

o
m

a
o
ct

UJ

•-•  U.
C  U.'

•  _J
2:   •

ZJ  z
z  c

D

•o

a.
o
a:

ct

UJ

IT

O

a.

Ui

ac

Z)

tn

a:  r
a  z>

7  •  U.!  •
O  UJ  C?
•-  Z   _JUJ

in       _)
in  o

«a  o

4  CD  OOD

X  a

<i  I-

z  X

o  •-•

z  a:  z:

Q    >  C

»  U.I     »

UJ  UJ

Z  _I  z        u
_l         I-  o

u  <i  u  iLj  a

33  u  X  a:  a

2
O

tn
tn

UJ

ac
a.

X
uJ

Z

o

(X

<

o  I-

z
o

UJ

6
u

cr       c
tn       UJ

c
c

Q

«  a

u.

UJ

_J
»  or
z  a
c  X
c   b

»

UJ-J

_»

<_■  <:

cv

z
c-
o

z

3
I— »

tn
(vjcn  0.

0.  UJO

X  cr  _i
UJ  a  UJ

•  X  (t

-t  UJ
UJ  ♦    (C
QC         Z
Q-         •-«
O
»

tn
tn

<

3

o

o

n

Q.

UJ  •->  X
H-   UJ
U  UJ

n  a:  0.

3

O

o

z      z

c      o

z        •-•

o        •-
o       >-•

c  z

z  o

c  o

u  *
Q.

c

»
IT

tn

<
a:  _i  oj
a  tj  r  ^
>    •  c

U'  U"  c>

z

_J  c  •-•

—I  O  I-  *—

<l  CD  O  UJ

U  O  1

UJ  it         V  o

tt  tJi        «-  z

C  •-   fV  c  c

z  a  a  z  z

C  X   o  O

O  ♦  UJ  o  o

a

X
lU

a
X  a.

o  X

UJ

a  CL

2    •

o  o
o  z

o
a  u
o  a
u.

z
:i^  UJ
o  I
O  I-

_l  •■

♦  c

z

•  c

*-  c

3.   <f

o

tn

z

l-t

tl>
u

cr  or

<

«    CL

o

2
O

o

ac
o

X

ac

O  UJ

z  •
=)  a.
o  »-
u,  X
o

o  •-•
z  a  V
c    •  4-
O  2  O
C   2

c  o   c

UJ    •  o
M  UJ

»-<  z  o

in       t-

UJ  tJ  o

X   QD  (D

»-
z

UJ

D    C
<:  2

a  o

c

«  <c  a.

85

c

2

<

^

ft
UJ

>-

Q.

z

O

<t

s:

X

o

X

o

u

K

3

VI

UJ

<x

X

r

«

X

..^

r^

r^

a

2"

r

r-

X

O

UJ

»-4

t/5

c

a
o

ft
O

*

UJ

z

<

a
a

a

a

—t

X

X

UJ

UJ

o

3.
K

a

UJ

xl

a

o

a

o
a

(V

9-

X

4

a

»

X

<r

ro

o

n

00

(V

tr

IT

X

a.

o

o

a

X

•j\

I

a

X

O  X

X

X

o

X

o

a

y-

K-

o

U  Q

o

o

a

o

a  c

o

a

o

a:

o

UJ

»-

UJ

U)

(/)  ►-

«-

3

4-

1-

o

X

o  a

X

o

a

a

i?

»-

\r\

»-

»

•

IT

••

»^

»

KT)

a

a

1-  a

a

a

X

Q  UJ

o

U

»

u

a

z-

»

z

_J

Z

»

a

UJ

KTi  Ui

^

a

UJ

2      •

{_

<I

X

<

o

c

>

c

>

c

X

U.'

»

»         9-

»

UJ

»

3  Q.

«

cr

»-

a

>

o

<i

_<

O

3

a

•>

X

a  X

a

»

a

O  *-

_l

<3

«a

<

-I

a

a

•

vD

»

1-

a

<I

»-  <i

h-

a

»-

U.  I

UJ

X

a

o

I

(J

»-  a

>

«-

>

«/)

»-

5-

X  X

I

»-

X

e?

or

o

UJ

z

o

O

tn

<i

z

»-

a

'»

tr

u.

5.

ID   2.

e?

u.

o

Q   •-<

c

—

C

G

o

•«-

•■

»

•

<r

IT

r

<.

^

u

C

K-  O

►— '

u

»— t

c

c   cr

-f  .

u

>

C

c

u

Uj

a

I/-

\r-

K

»

a

>~

tr

_j

a

CJ

o   c:

'^

_i

C

a

^

X     »

c

^-

a

»-

^

•^

2"

(A<

tr

\r

C

^

c;

c

»

•>

a

»

•-    •

»

*-

a

»

a

UJ  -z'

a

o

<

<pa.

^

<f

<r

<3

<

z

•— '

(Tr

z

-y

-^

X

z

a  2  z:

a

2-

2:

X

z

c

c

u.

»

r

P

•

c

c

2'

_j

_l

_i

•

C

ar

•

C

c

u.

c

a

(-.  c

a

c

C:

u.

c

»-

r    c

D

t

1/

L

H

c

c

C

t

1

~

5i

c

0

»-

c

>

0    n

>

c

D

»-

c

tA

U!     »

c

—

n

~~

'^

»

/-■

•

*

•

IT'

■'

(A

lA

•

•■

'y^

♦

U'

•■       »

U'

»

••

tr

»

»

(SI  UJ

or

fl

•— i

cc

c

Z

;^

<^

UJ

<3.

o

<

UJ

U.'

UJ

e  U.I

Ul

IkJ

UJ

►-1   i  ■

_J

»-

u

j<

_<

K

1

Lj

*—

o

a

u>

u.

r^

r

»— »

_J

•— 1

5

~

^

Tt"

_j

1'

_JUJ  -

:

^"^

■p"

_»

'^

O

•-r.

-J

-~,

'a»

_j

_l

h-

~

D

o^

_J

t—

►—

-t

»

1

1

t—

_J

r—

o

u.'  o

<c

o

cL

»—

<i

(_/  "3

<_;

U.

.-%

h-

o

o

o

<J

a.

<i

U.'

e

Uj

o

CJ

u

<x

o

•a

u  c

<a

u  u

O

<;

CJ  U.'

a

T    I.

o

X

a

y)

u

X

o

^

a

ut-

'Ti

r

I

X

CI?

<

U

X

X

a;

X

i/)

X

o

T;

O

X    X

o

.T  a

•T

o

I  a

a

»-

Z

.—

UJ

a

a

tt  a

u.

u.

r—

c.

(V

r

<  a

c

c

a

C

a

&

c

c

a  X

o

c

o

c

c

c

c

o

c

c

c

u

<L

c

»-

»-

*'

r:

t—

K

1-

•-

£

♦  a

UJ

*  Q.

(A;

t/>

Ul

*  a

tA)

lA

(/^

lA

UJ

86

tf)

UJ  >-

ac

a  4-

UJ

uia

Q.

'

X  «■

o

r

l/)Z

_l  «-

Ui  UJ

4-

o
z  tr

o

u>  z
c

UJ  o

o

r>

or
Z  Ql
QL  X
O  UJ

•-•

UJ

X
3
Z

in

o

UJ
UJ

z

z
o

l-t

(n

UJ

a

Q.

X
UJ

(VI

X
X
UJ

(/)

>-

IT:

«■

X

«^

O  U)

I-  Q.

z

rHO

1-

■•a
u.

UJ    »

X  •-

r^

-*iO

r-^

2

X

t-  X

en  X

l-«

3)

3  Q

tn

»-

o  tn

D

1-^

>—

.^^

X  UJ

_l

O

fVJ

in

UJ

»-  z

za.

z

X  QC

Z

X

X

UJ
X   '-*

3  tn

X

o  •
4  z

»-

X

cr

«-  UJ

O  X

a

X  o

X

in

X

Z  X

(M  (n

••3

z

a

o

o

o»-

H4  UJ

X

X  a

X

^o

X  o

X

o

X  Z

UJ

o

l£

£X

«"  X

cn  a

UJ

UJX

UJ

X

Q  <

o  <t

o

3

<  «■

X

a

^QC

a

-J  o

</)

a

a  UJ

a

«-

z    •■

<    »

<x

«—

I    •

z>

q:

•— 1  UJ

UJ

4- 1

UJZ

»

•>  •>

»

o

<ta

•  z

»

_J

oz

o

lij

X    •

»

(/)

a  UJ

X

Z  X

z

«-

X  <I

h-  o

>

>

♦^  o

a;

»

3  X

Q.

«a

a  X

»-

r>  »-

i— (

UJ  >

(/)  o

<o

•  z>

<

a

V)  <i

»-

>

X  f-

<

X  >«

X

o

X  _J

Z  X

X

^

z    •

^—

Z

h-

1-  2:

X

K

UJ_I

UJ     v'

o

o  u

a

»-l

O  O  O  1-

X

4-

Q  >-

U.)

»-•

u.

WJ  i^

o

Ul

>-  y-

—  o

z

•-  z

o

X

«-  o

U  O)

<x  *-

3  Z  X

2

5

c

UJ

•  c

»-H

_J

<S.    <i

UJZ

c

o

v>  o

•>

u>

UJ     »

•     »

9-

tr

•  «a  _J

<I

■"^

•

1

G

-H     C

a

Q-

e

tS  UJ

5  a

z

o

(/>  o

(T

^

^  O)

<r  <r

w

»

>  ►-  a

Z

Z

r

»

Q

V     »

a

»

a

*-*  a

<t  o

a

»

c

<  »

(/5

3

<  U)

w  t/>

1/5

r

ZOO

3

2

X  s  z

K

Z

<

_«ir

z  >-

o

z

z

-J  z

<

z

z  «a

<r  <i

«x

»-«

<  Z  V)

►-

liJ

T

c

U)

r  c

U

C

K

a

^

K

o

CL

o  c

_1

«

_J

_j  _J

_l  c

»   ♦  cc

{/

h

o

f-

7    C

»-

Q

(/

z  u

c    «

V

c

C

•  c

C

c  c

t  t

c

y  i  r

!

•a

7'

•

tr

»     »

U^

»

*

C   X

o

»

.— "

c    ♦

»

a

?    ♦

»       »

•-

T^

IT  (T   (A   <\

CT

U  C'l

UJ

c;  UJ

UJ

•

X  ►-•

_IC?

-I

UJ  e

c

a  o

e  o

UJ  o

<t  <t

r>'

o

t— •

<.

_J

UJ  i'

_J

z

<_)

tr  _J

G

UJ  _J

_l

UI

<-•  _J

UJ

UJ

t~*

o

O  UJ

UJ  UJ

z

•-«

T    T   _J  "-I

<

-J  CT)

_l

_)

►-

C

O  _f

a.1

H-   _J

<

H-  <

u

♦—

o

>-

(A,

_i  •—

»-

_i  tr>

o

<i

o  o

<3

ou.

a

c  ^

1-

z>  <.

u

o

U-'  O

X  o

CJ

u

a

z  o

o  o

O  (A,

O  <J  <1   U-

I/)

■ti

<3"

■x;

a

li  X

o

a.

a.  3c
a

>

Uj

2

a.  o

»-

X

>-

a:

X
Z  3

X

X

a

<  I

X  X

X

<a

T  X  O  X

»-l

UJU.I

UJ

tr>  (/I

z

X  z

•-♦

tr

(T  T

_l

a

CL

a

a.  or

c

c

~

U.  1-

0.

c

o

a

Q

>  o

c  z

o

IT/

o

X  «-

s

c

X

X

X

U'  X

o

C  X

c

»-

^

»-

»-^

c

u

U'

u

D    UJ

^

a  c

a

t/>

UJ  «    «    «

«

Q.

CL

a

a

♦  a

UJ  «

a  <

87

ct

cc

u^

in

>-

a:

c

»-

a

1^

•y)

a:

in

X

X

jj

kO

— .

tv

D

3)

o

o

»

CM

a

tv

o

OJ

tv

in

<    IT

in

=a

z

t-

a

y-

r^i

»

»

3

m

3C

■j)

^

o

\n

a:

t—

tr

a:

>o  X

X

cc

2

a

^O

Cs

a

a

a

a

o

U)

o

o

a  o

c

<i

<■

o  fvj  a:

J"ii

o

a:

o

o

X

i

ct

£t

O  X

X

I

»

<3

a  cr

□.   ^

<i

UJ

<i

a

CL

c

tr

Ct

<i  a

X

o

2

••

C  U-'

o  u.

»

»

»

or

liJ

<a

UJ

UJ

UJ

UJ

<-

C

2

<I    •

<I  u.

u.

u.

X

u;

•-

»

»

••

z   •

•"

»

23

»— <

•  (T

•  o

•—1

»-♦

o:

»

X

a

X

Q.

Uu  a

X

z

»

X

t/> :::;

a  >c

QD

CE

»-

a

<»

\-

«a

1-

X  ►-

»-

a

>

Q.

3  z

H-  <-

Q.

H

y)

1-

X

X

r

X

»-  I

X

a

ZJ

^

C

—••-I

U-  »-

C

IT

cc

H-

u.

5.

o

5

o

••e>

ei

f

»

<r

•

a  5

Ll:  a

»  »-

•

~"

ir

u

C

1— •

O

»-l

a  •-

t-H

n

>

»-

(/

<-  ♦

__'     •

y^  tr

u-

cr

—r

_f  or

o

cr

o

a

a  a

a

ts

i?

c

lA,

»   •-

»  ZL

(/)  •—

ir

•-

c

•

a

»■

»

•

»

X    »

•■

(/>

•sr

2

<t

Z  3

C

-^  t-/

<  _t

<

"Z

or

Z

X

z

ct

A-

z

<!:

z

UJZ

•^

a

•■

•-

c  c

"^

c.  c

-J  IT

:

c

<

fl

L.i

^—

a

c

c

G.

i—

»-  c

c

~

:^

Nr

t

r    (:

cl

o

C    C

c

c

f-

c

»-

n

>

c

r

>

c

Z   D

c

tr

w—

(/

cr

•

»      »

C7

r>

•  r^

r  <-•

f\

»■

V

(/"

a

•

{f

»

ll

1    »

»

U

»

a    »

•

^

<

a

t— 1

UJ

e  u-i

C7  CD

liJ

UJ

-^

U-'

UJ

^

UJ

UJ

F-H

UJ

1— t

U.'

ru

_I

»— (

~

*~

"^

LU  ~

_J

•— "

U'  •-•

z  _s

t-H

^

uJ

-J

»— B

-T*

"Z

_J

^

_j

'  UJ

*■"

_j

-7

-I  z

~

u

<x

_l

f—

1—

f—

_>

t—

lT

_J

1—

_j

t—

<r

1

_j

li

f—

i

H-

\-

■**"

-?

«

ti.

o

O

u.

u.

CJ

c  u

<I

U.

t_  (A,

O  <I

Li_

<-;

4J

<i

Ll

u

o

<:

t-

<[

u

C-

«.

c

U.-

<^  O

u

c;

U.

a

>

o

J.

r

T

X

a

r?

T:   X

o

.r

C  <I

:X!  O

X

XI

X'

o

r--'

00

X

o

X

U  X

X

tJ

X

X

U  X

X

r

'V

X

'Si

(V!

c

r-

r  c^  fv

(V

r

<

a

a

a

a  a  Q

a

a

a

a

f_.

c

c  c    c

c_

c

c

c

<t

<f

<  <  <,

<r

<i

<

<

a  a       a
c    ;,         c

U.'   «    UJ

BIBLIOGRAPHIC  DATA
SHEET

1.   Report  No.

UTUCDCS-R-7'5-7US

4.   Tiilf  and  Subt  it  Ir

An  Interactive  Table -Driven  Parser  System

3.  Recipient's  Acceaaion  No.

5.    Report   Datr

Aijf/ust  197!?

7.  Author(s)

Michael  Harry  Tindall

8-    Performing  Organization  Rept.

No-   IJTnCDC3-R-75-7H5

9.   Performing  Organization  Name  and  Address

University  of  Illinois  at  Urbana-Champaign
Department  of  Computer  Science
Urbana,  Illinois  618OI

10.   Proiect/Taslc/Work  Unit  No.

11.  Contract/Grant  No.

NSF  Grant  EC-U15II

12.  Sponsoring  Organization  Name  and  Address

IBM  Corporation

Thomas  J.  Watson  Research  Laboratory

Yorktca-^n  Heights,  New  York  10598

13.   Type  of  Report  &  Period
Covered

Master   of   Science  Thes

14.

15.  Supplementary  Notes

National  Science  Foundation,  Wasington,  D.C  20550

16.  Abstracts

The  design  and  implementation  of  the  parser  system  for  a  multi-lingual  compiler
on  the  PLATO  IV  computer-aided  instructional  system  is  described.  A  brief
discussion  of  the  formal  transition  diagram  parsing  method  is  followed  by  the
description  of  the  assembler  language  used  to  specify  the  syntax  of  a  prograjmning
language  and  the  description  of  the  actual  parser  tables.

17.  Key  Words  and  Document  Analysis.     17a.  Descriptors

Compilers,  Parser  Systems,  Table  Driven,  Interactive  Compiling,  PLATO  IV

17b.   Identifiers/Open-Ended  Terms

17c.  COSATI  Field/Group

18.  Availability  Statement

Release  Unlimited

19.  Security  Class  (This
Report)

IJ>^ri,ASSIFIF.D

20.  Security  Class  (This
Page

UNCLASSIFIED

21.  No.  of  Pages

90

22.  Price

FORM   NTIS-35    (10-70)

USCOMM-DC    40329-P71

:<^

\r

.^'

^

m

2  0  ^9*^

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
