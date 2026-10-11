---
source: "https://archive.org/details/bitsavers_cdcterminasersGuideApr81_14411870"
title: "cdc :: terminal :: IST :: 97405900C PLATO Users Guide Apr81"
author: ""
year: "1981"
captured_at: "2026-10-11T05:14:19Z"
updated_at: "2026-10-11T05:14:19Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "プラトン"
query: "Plato"
plain_text_url: "https://archive.org/download/bitsavers_cdcterminasersGuideApr81_14411870/97405900C_PLATO_Users_Guide_Apr81_djvu.txt"
public_domain: true
book_text_truncated: true
subjects:
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# cdc :: terminal :: IST :: 97405900C PLATO Users Guide Apr81

- 著者: -
- 初版: 1981
- 情報源: [archive](https://archive.org/details/bitsavers_cdcterminasersGuideApr81_14411870)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[プラトン]]
- 研究動向: [[プラトン-現代研究動向]]

## Full Text

PLATO

USER’S GUIDE

CONTROL
DATA

PLATO

USER’S GUIDE

REVISION RECORD
DESCRIPTION

0
(07-18-74)

=

Preliminary release.

Manual released. This printing obsoletes the previous edition.

>

(09-19-75)
Manual revised to reflect most currently developed displays and procedures. Because of extensive
(12-31-76) changes to this manual, change bars and dots are not used and all pages reflect the latest revision

level. This edition obsoletes all previous editions.

Manual revised to reflect changes to system software from Cut 5 to Cut 22. Because of extensive

(04-30-81) changes to this manual, change bars and dots are not used and all pages reflect the latest revision level. This

edition obsoletes all previous editions.

Q

GM

97405900

as)
r=
s
_
7
ce)
~
-
=
co)
=]
4
2

REVISION LETTERS |, 0, @ AND X ARE NOT USED
Address comments concerning this
manual to:
Control Data Corporation
HQA02Q
P.O. Box O
© 1974, 1975 1976, 1981 Minneapolis, Minnesota 55440

by Control Data Corporation :
or use Comment Sheet in the back of
All rights reserved this manual,

Printed in the United States of America

ii

LIST OF EFFECTIVE PAGES

New features, as well as changes, deletions, and additions to information in this manual, are indicated by bars in the margins or by a dot
near the page number if the entire page is affected. A bar by the page number indicates pagination rather than content has changed.

OVOUVVOVVOYNVONDYDD |; OOOVVUVDOVVVOVVVVNODV VDD

Cover
Back Cover

Inside Back

OLOLOROLOLOLSLOLOLOROUSLOLOLOLONOLONORONORON OL ONO TOLOLOROLOLOROLOEOROLOR OL OROEOROROR OLOLOROR OL OL OLOROLSTOLOLOTOIORORS)

OOVDOUOODOVVVOOVOVVOVVOVVVVVVOVONVODVVDOVVODNDND , OOVOOVVOOVUVOVVVOVYVD

Front Cover
Title Page

tt

lii/iv
2-3
2-4
2-5
2-6
2-7
2-8
2-9
2-10
2-11
2-12
2-13
2-14
2-15
2-16
2-17
2-18
2-19
Divider
3-i/3-ii
3-1

iii/iv

97405900 C

PREFACE

DEFINITION

The CONTROL DATA® PLATO® system is a computer-based educational development and delivery
system.

AUDIENCE AND ORGANIZATION

The PLATO User's Guide can be used to orient new users to the PLATO system and serves as a
comprehensive reference manual for more experienced system users.

The manual is organized according to the PLATO system's defined user types. Each section defines and
describes the features available to a specific user type and gives step-by-step instructions for using each
feature.

Some information contained in this manual is applicable to all users while other information applies only
to specifie user categories. All users should read the introduction (section 1) for a general introduction to
the PLATO system and user categories and then read the section(s) of the manual designated for their
specific user types.

General reference information is contained in the appendixes located at the end of the manual.

The information contained in this manual is accurate as of Cut 22 of the PLATO system software.

RELATED PUBLICATIONS
Refer to the following publications for additional information related to that contained in this manual.
Control Data Publication Publication Number
PLATO System Overview 97406700
PLATO Terminal User's Guide 97404800
PLATO Author Language Reference Manual 97405100
PLATO Author Language Instruction Formats 97406600
PLATO CMI System Overview 97406100
PLATO CMI Instructor's Guide 97406300
PLATO CMI Author's Guide 97406200

97405900 C Vv

Control Data Publication Publication Number

PLATO Courseware Catalog 76360775
Micro PLATO User's Installation Guide 76368339
Micro PLATO Instructional Disk 76773000

These publications are available through the nearest Control Data Corporation sales office or Literature
Distribution Services Center. :

DISCLAIMER

This product is intended for use only as described in this document. Control Data cannot be responsible
for the proper functioning of undocumented features.

97405900 C

CONTENTS

1. INTRODUCTION

Definition
User Categories
Student
Instructor
Author
Recommended Learning Sequence for New
Users
Using the PLATO Keyboard
New User Registration
The PLATO Terminal
Connecting the PLATO Terminal
Using the Central PLATO System
Using the Micro PLATO System
How All Users Sign On
How All Users Sign Off
How to Change Your Password

2. USING STUDENT FEATURES

Curriculum Structures
Taking PLATO Lessons
Understanding Your Curriculum
Index and "mrouter" Lessons
PLATO Learning Management Lessons
Requesting Help
Help Within a Lesson
Help from Your Instructor
TERM-ask
Student Notes
Personal Notes
Commenting on Lessons (TERM-comment)
Communicating with Group Members
Handling Problems
Communieation Errors
Lesson Execution Errors
Messages You Might Receive
Helpful Tools
Checking the Time
Doing Mathematical Calculations
Additional Student Options
Using the Micro PLATO System

97405900 C

3. USING INSTRUCTOR FEATURES

Introduction
PLATO Facilities Display
Using AIDS
The PLATO Group
Group Data
General Group Information
Associated Files
Security Codes
Group Security
Group Operations
Registering Students
Deleting Students
Deleting Individuai Students
Deleting All Students
Listing All Group Members
Inspecting/Changing Student Records
Leaving a Message
Using Your Account
Reviewing and Studying Lessons
PLATO Courses and Curricula
Designing a Curriculum
Instructional Management Tools
How to Use Index Lessons
How to Use the System-Supported
Router ("mrouter")
How to Use PLATO Learning
Management (PLM)
How to Use Your Own Router
Using Datafiles
Additional Instructor Options
Monitoring Group Members
Specifying Group Data Collection
Templating Records
Copying Records from Another Group
Creating Instructor Records
Managing Personal Notes
Using Notes
Using Interactive Communications
Requesting Prints
Receiving Help
Giving Help

4, USING AUTHOR FEATURES

Introduction
Author Records
The Author Mode Display
AIDS
Catalog of Available Courseware
Notes
Personal Notes
User List
Prints
Understanding File Structure and Use
Getting Help
Using AIDS
Consulting Help for Authors and
Instructors
Using TERM-ask
TERM-talk
Using Communications Features
Using the Talk Feature
TERM-busy
TERM-reject

Commenting on Lessons and Features

Notes
Types of Notes Files
Notes File Access
Using General Notes
Using Personal Notes
Using Lesson Notes
Using Student Notes
Using Intersystem Notes
Writing General Notes
Using Notes File Features
Using Reference Tools
AIDS
Using the Catalog of Available
Courseware
Title Index
Author Index
Subject Index
File Name Index
Using the On-Line Author Listing
Using the PLATO User List
User Statistics and Flag Settings
Time-Saving Features
TERM-spell
Using Documentation Features
Using Documentor
Graphics Utility for Interactive
Documentation Ease (GUIDE)
Requesting Prints
Writing PLATO Lessons

viii

4-1

4-1
4-1
4-3
4-5
4-5
4-5
4-5
4-5
4-6
4-6
4-10
4-10

4-15
4-16
4-19
4-20
4-20
4-21
4-21
4-22
4-22
4-22
4-23
4-23
4-29
4-31
4-32
4-33
4-33
4-36
4-39
4-39

4-39
4-41
4-43
4-44
4-44
4-44
4-47
4-49
4-49
4-50
4-50
4-50

4-52
4-53
4-56

Using the TUTOR Lesson File
Registering Author and Lesson

4-56

Infor mation 4-56
Assigning Security Codewords 4-58
Editing TUTOR Files 4-61
Creating a TUTOR Block 4-61
Using the PLATO System Editor 4-63
Requesting Editing Help 4-66
AIDS Quick Reference (Q-Ref) 4-67
Condensing a Lesson 4-67
Creating Displays 4-68
Sereen Locations 4-68
Inserting and Showing Displays
(ID/SD 4-71
Creating Characters and Line
Drawings 4-71
AIDS Listing of Display Commands 4-84
Suggested Guidelines for Writing
PLATO Lessons 4-85
Understanding ECS/ESM Usage and
Charges 4-85
Referencing PLATO Author Language
Commands 4-87
Using Other Block Types 4-88
Common Block 4-88
Micro Block 4-89
Leslist Block 4-90
Vocabulary Block 4-90
Listing Block 4-90
Text Block 4-90
Copy-a-Block 4-91
Using Advanced Editing Directives 4-91
Using Block Listing Display Options 4-93
Documenting Lesson Changes 4-93
Writing Routers 4-93
Preparing Lessons for Publication 4-94
Using Data Collection Files 4-95
Dataset Files 4-95
Nameset Files 4-95
Code Files 4-96
Using Advanced Author Options 4-96
TERM-pnote 4-96
Setting an Alarm 4-96
Using an Automatic Sign-on 4-97
Using Author Reserved TERMS 4-97
TERM-cursor 4-97
TERM-grid 4-98
TERM-step 4-98
TERM-charset 4-99
Additional Security Options 4-99
Additional SHIFT-DATA Display Options 4-100
97405900 C

Micro PLATO Authoring
Micro PLATO File Security
Micro PLATO Levels
Micro PLATO Lesson Execution
Flexible Disk Preparation
The Micro PLATO Router
Author References

5. USING ACCOUNT OPTIONS

General Account Information

Account Owners and Account Directors

Using Your PLATO Account
Account Main Options Display
General Account Information Display
Maintaining Account Security
Using Security Codewords
Using an Account Access List
Creating an Account Access List
Registering Users in the Account
Access List
Assigning Access Options
Allowing/Disallowing System
Access
Maintaining File Security
Default File Security Codes
Changing File Security
Managing File Space
Creating Files
Destroying Files
Renaming Files

A. KEYBOARD

B. ACCESSING THE PLATO SYSTEM
THROUGH THE DATA SERVICES
NETWORK

97405900 C

4-101 Making Another Copy of a File
4-102 Lengthening and Reorganizing Files
4-102 Shortening Files
4-102 Copying the Contents of a Lesson
4-103 Changing General File Information
4-104 Inspecting File Information
4-104 File Space Management Tools
Displaying File Information
Listing Files in Your Account
Listing Archived Files in the
5-1 Account
Reviewing All File Operations
5-1 Inspecting File Data Displays
5-1 Making a Leslist of Account Files
5-3 Listing the Last Editors of Files
5-3 Lesson Usage Data
5-5 Listing Selected Major Users
5-7 Usage Data by Lesson
5-7 References by Other Accounts
5-8 References to Lessons
5-9 Listing Current Account Users
Report Generator Options
5-9 Group Records Report Generator
5-11 Options
Archiving Files
5-12 Controlling Print Access
5-13 Using Interaccount Options
5-13 Using Network Options
5-15 Preparing the Account to Use the
5-16 Networking Features
5-17 Intersystem File Security
5-18 Understanding Network Transfer
5-18 Units (NTU)
APPENDIXES
A-1 C. GLOSSARY
B-1
INDEX

5-18
5-19
5-19
5-19
5-20
5-21
5-21
5-22
5-23

5-23
5-23
5-24
5-25
5-25
5-26
5-26
5-27
5-27
5-27
5-28
5-28

5-30
5-31
5-31
5-32
5-32

5-32
5-39

5-40

C-1

PLATO Keyboard
IST Terminal
IST-II and IST-If Terminals'
Exterior
Telephone
Acoustic Coupler
Micro PLATO Station
Press NEXT to Begin Message
Welcome Display
Group Name Display
Password Choice Display,
Parts A and B
Password Display
Sample PLM Curriculum Display
Sample Student Index Display
PLATO Facilities Display
Author Mode Display
Curriculum Structure-Example 1
Curriculum Structure-Example 2
Examples of PLM
and Index/"mrouter" Lessons
PLM Course Options Display
PLM Module Options Display
Personal Notes Display
Personal Notes Text Display
PLATO Facilities Display
AIDS Title Display
AIDS Index
AIDS What TUTOR Feature Display
Group DATA (Directory) Display
Security Codewords Display
Group Operations Display
Account Main Options Display
General Account Information Display
Curriculum Structure-Example 1
Curriculum Structure-Example 2
Associated Files Options Display
Instructor File Information Display
Curriculum Options Display
Module Listing Display
Available Instructor Options
Display
Available Author Options
Display
Author Mode Display
SHIFT-DATA Display
Examples of File Indexes
Examples of File Directories
AIDS Title Display

FIGURES

1-5
1-8

1-8

1-11
1-12
1-13
1-14
1-15
1-16

1-17
1-17
1-18
1-19
1-20
1-21
2-1

2-2

2-4
2-6
2-7
2-11
2-12
3-2
3-5
3-6
3-8
3-10
3-13
3-14
3-19
3-20
3-22
3-23
3-27
3-29
3-29
3-32

3-38

4-2
4-3
4-4
4-7
4-8
4-11

4-7

4-8

4-9

4-10
4-11
4-12
4-13
4-14
4-15
4-16
4-17
4-18

4-19
4-20
4-21
4-22

4-23
4-24
4-25
4-26
4-27
4-28
4-29
4-30

4-31
4-32
4-33

4-34
4-35
4-36
4-37
4-38
4-39
4-40
5-1
5-2
5-3
5-4
5-5
5-6
5-7
5-8
5-9
5-10

AIDS Index 4-12
AIDS What TUTOR Feature Display 4-14
PLATO Notes Display 4-24
Notes Options Display 4-25
Notes File Index Display 4-26
Sequencer Editing Options Display 4-28
Personal Notes Display 4-30
Block Listing Display 4-32
Insert Mode Display 4-34
Notes File Options Display 4-37
Courseware Catalog Options Index 4-40
Catalog of Available Courseware

Title Index 4-41
Lesson Information Display 4-42
Lesson Description Display 4-43
Directory of PLATO Authors Display 4-45
Alphabetical List of Authors

Display 4-46
Total Users Display 4-48
Documentor Section Index Display 4-51
Print Requests Display 4-54
Author Information Display 4-57
Block Listing Display 4-57
Security Codewords Display 4-59
Block Creation Options Display 4-62
Example of Forward and Backward

Editing Directives 4-64
Editing Help Index 4-66
Coarse and Fine Grids 4-69
Example of How Letters Occupy

Spaces 4-70
Charset Options Display 4-72
Character Design Display 4-74
Multiple Character Creation Display 4-77
Lineset Normal Grid 4-79
Lineset Large Grid 4-80
Lineset Options Display 4-81
Author Mode Display with ECS Tally 4-86
Account Main Options Display 5-3
General Account Information Display 5-6
Access Options Display 5-10
User Access Options Display 5-11
File Management Options Display 5-15
File Data Display 5-22
Lesson Usage Data Display 5-26
Author Information Display 5-29
General Account Information Display 5-33
Network Options Display 5-34

97405900 C

2-1 Troubleshooting Procedures
3-1 Curriculum Design Options

97405900 C

TABLES

2-14
3-25

4-1 Functions of Editing Keys for
Normal and Large Grids
5-1 File Types

xi

SECTION 1
INTRODUCTION

Definition
User Categories
Student
Instructor
Author
Recommended Learning Sequence for New
Users
Using the PLATO Keyboard

97405900 C

INTRODUCTION

1-1
1-2
1-2
1-3
1-3

1-3
1-5

New User Registration

The PLATO Terminal

Connecting the PLATO Terminal
Using the Central PLATO System
Using the Micro PLATO System

How All Users Sign On

How All Users Sign Off

How to Change Your Password

1-i/1-ii

1-6
1-8
1-9
1-10
1-12
1-14
1-21
1-22

INTRODUCTION 1

(a a a TE Le EE AT ES

This section is a general introduction to the PLATO system and its user types. All users should read this
section before using the PLATO system and before reading other sections of this manual.

DEFINITION

The PLATO system is a computer-based educational development and delivery system. It can be used to
prepare and present instructional material, to summarize student performance, and to evaluate and revise

materials. It can also be used as a communications tool between users working to achieve the same
instructional and training goals.

The system is extremely diverse in its capabilities and has a wide variety of uses in many different
environments. PLATO systems are found in schools, universities, hospitals, and businesses throughout the
world. Students of all ages, business people, medical professionals, airline pilots, engineers, technicians,
secretaries, salespeople, accountants, bankers, and bank tellers all use the PLATO system daily to learn
new concepts, facts, and procedures; collect data; review mastered materials; simulate complex,
dangerous, or expensive laboratory tests; or take qualifying, competency-based, or course completion
examinations.

The PLATO system is a delivery system. It is not based on any one set of training or instructional
principles. It can present instructional materials according to any instructional theory, philosophy, or
methodology, or according to any training method or plan. It is, therefore, a valuable tool for
researchers, as well as a teaching and testing tool for instructors and students which accommodates
change and growth in education, training, and business, as well as state-of-the-art computer technology.
The system provides a special author language which allows users without computer programming
experience to develop instructional materials and tests. The instructional materials can be designed to
meet any training needs.

Individualized instruction (for example, frequent testing, frequent feedback, detailed feedback, alternate
learning paths, mastery learning, objective-based instruction) is most frequently chosen for delivery on
the PLATO system because its presentation capabilities are broad enough to support these demanding
instructional methods. But, the capability to support all methods, old and new, has been and will be
retained.

The PLATO system is constantly being evaluated and upgraded. PLATO system users can rely upon a

delivery system which today meets their needs their way, and which tomorrow will add the additional
capabilities that improved technology and education and training research provide.

97405900 C 1-1

USER CATEGORIES

The PLATO system identifies each user as being one of three basic user types and treats each user on the
system somewhat differently. The three user types are:

e Students (including multiples)
e = Instructors
e Authors

The options and features available to users in each of these user categories is determined according to
each user's individual needs. Some users can access all the features available to users in their specific
user category, while others can access only a subset of those features, depending upon their specific needs.

The sets of features available in each user category overlap. This allows users who are instructors to
access student features as well as instructor features, and allows users who are authors to access student
and instructor features, as well as author features. It is recommended that new users initially start using
the system as students and then gradually progress to instructor and author status. Refer to
Recommended Learning Sequence for New Users later in this section for recommendations on orienting
new users to the system.

Students can access fewer system features and options than other user types. They require minimal
instruction on how to use the PLATO system because the system guides them through their lessons and
activities. Instructors can access more system features and options than students, and therefore require
more time to learn how to use those options and features. Authors can access the greatest number of
system features and options of all user types and therefore require the greatest amount of time to learn
about the system.

The following paragraphs describe each of the three PLATO user types.

STUDENT

A student typically uses the PLATO system to study a lesson or set of lessons called a curriculum. A
PLATO student is not always a student in school but can be anyone using the system for instruction or for
information. The system guides the student user to a specific lesson, program, or index of lessons or
programs. Because careful guidance is provided, many infrequent users utilize student sign-ons (a user
identification which is registered in the PLATO system) which take them directly to what they want to
see or use (executive report generators, data summaries, or instructional evaluation tools). A user can be
assigned a student sign-on when a predetermined set of lessons, features, or programs is needed. Only a
small amount of instruction is needed for students to learn how to study lessons or access programs and
system features.

The PLATO system keeps a student record or file for most students on the system. A student record
stores various kinds of information about the student such as the name of the lesson(s) the student is
studying, the number of hours the student has used the system, test scores, and other information.

A multiple is a special kind of student user. Most often, multiples are users who are not assigned to
specific lessons or activities, but are users who want to see examples of various features, capabilities, and
lessons. Multiples share sign-ons with other multiple users. That is, they share the same identification
required to use the system. Because multiple users share their sign-ons with other users, the system does
not keep records of lessons completed, test scores, and so on, as it does for student users. Each time a
multiple user signs on to the system, the system treats that user as if it is the first time the user is
signing on.

1-2 97405900 C

INSTRUCTOR

An instructor can enroll and remove students from courses, review lessons, write and read notes to and
from students, see information on student performance, and collect student data. Instructors also help
students when problems arise. They can organize lessons into curricular groupings suitable to various
classes of students, and can define the kinds of student information they want collected for each
curriculum.

Instructors can work and communicate with other instructors or authors writing or testing new lessons. In
addition to their own instructor capabilities, instructors can access all the system features students can,
and can view lessons and materials as students do.

“Some instructors are also account directors. Account directors coordinate and manage the use and
allocation of contracted PLATO resources.

AUTHOR

An author develops instructional materials for the PLATO system. Authors communicate directly with
the system by means of a computer language called the PLATO Author Language. Authors have access to
files on the system and can write lessons by inserting, modifying, or deleting information in these files.
They can review completed lessons and lessons they are developing.

Authors can access communication features on the PLATO system which allow them to obtain general
information on system operation, plans, news, and new features; write notes to other system users;
participate in group discussions; obtain help or suggestions from Control Data PLATO consultants; or
communicate privately with other users. Authors can access all the system features students and
instructors can, in addition to their own author capabilities.

Some authors are also account directors. Account directors coordinate and manage the use and allocation
of contracted files used by a given customer and perform file management tasks such as creating,
destroying, renaming, and lengthening files.

RECOMMENDED LEARNING SEQUENCE FOR NEW USERS

The PLATO system contains a large number of features and has many capabilities which can sometimes be
overwhelming to new users, particularly new authors and instructors. A good way for new authors and
instructors to learn about the system and become familiar with its features and capabilities is to initially
start using the system as a student user and then gradually progress to an instructor user, and finally to an
author user. Students can see and use many system features, but always in a controlled setting. As
students, users practice interacting with the system and gradually understand its operation. Some
examples of the typical kinds of things most new users gain from starting as student users are: learning
how to sign on to and off from the system; recognizing when the system requires a typed response;
learning the actions of function keys and becoming accustomed to using them; and understanding the
general structure of most PLATO lessons.

Once new users have achieved a basic level of understanding as students, those users who are designated
to become instructors and authors should be given instructor access to the system. As instructors, users
are introduced to several new sets of features which were not available to them as students. Many of
these features help users see some of the behind-the-scenes actions which controlled what they could see
and do as students. Communications features are available to instructors, as well as documentation and
word processing tools which introduce them to editing features.

97405900 C 1-3

After users are comfortable using the features available to instructors, those users who are designated to
become authors should then be given author access to the system. As authors, users are introduced to
several more sets of features which were not available to instructors. Many editing features are available
to authors which allow them to write lessons for students to study. Authors can access all the features
instructors and students can, in addition to those reserved exclusively for authors. Other features allow
authors to act as consultants to help other users while using the system.

The following is a list of lessons which introduce several basic PLATO system features and provide
examples of some of its capabilities. These lessons provide a good orientation to the PLATO system and
give new users an opportunity to practice interacting with it. If you are a new user and are not initially
given student access, you should contact the person who registered you in the system, ask for student
access, and arrange to see these lessons before progressing to instructor or author access. If you are
résponsible for registering users in the system and assigning user types, you should initially register all
new users as students and assign these lessons as an introduction to the PLATO system.

PLATO File Name PLATO Lesson Title Purpose
Swhatsnext What's NEXT Defines basic terminology; introduces
users to using the keyset.
Sgenintro An Introduction to the Provides a general introduction to the
PLATO Terminal PLATO terminal and keyset.
Stermcomme TERM-comments Describes how to use TERM-comment, a

system feature which allows users to
comment on lessons they are studying.

Stermeonsu TERM-consult Describes how to use TERM-consult, a
system feature which allows users to
receive on-line help from PLATO system

consultants.
Snotesintr An Introduction to Introduces PLATO system note writing and
Notes sending facilities. Covers both general
and personal notes.
#frose Rose Iustrates graphics capabilities.
Scalculate Calculation: a Touch Illustrates touch panel capabilities.
Lesson
Sdarts Darts Tlustrates how to numerically respond

to test questions.

In addition to the above lessons, new users should also be given access to AIDS and the Catalog of
Available Courseware.

1-4 97405900 C

USING THE PLATO KEYBOARD

The primary means of communication with the PLATO system is through the terminal keyboard. The
PLATO terminal keyboard resembles a standard typewriter keyboard (figure 1-1). Like a standard
typewriter, it has character, number, and punctuation keys (unshaded keys in the figure). The PLATO
keyboard, however, also contains function keys (shaded keys in the figure). Funetion keys instruct the
system to move from one screen image to another or add information to your present screen display.
They are used instead of typing an instruction and are always represented with capital letters (for
example, NEXT, BACK, and HELP are all function keys). Some examples of the kinds of things function
keys are used for are: to tell the system you are finished reading the information on one screen display
and are ready to see another; to go back and reread a screen display read previously; to see extra
information to help you understand something; or to do some lab problems or exercises.

Which funetion key you press depends upon the kind of action you want to take. The following describes
the function keys used most frequently by PLATO system users. New users should learn how to use these
keys before using the PLATO system.

For more detailed information on the PLATO keyboard, refer to appendix A.

eee)

JIODOOISOOOFE Se)
SEN ISGMOOOO IIIS)
SOOO IIIDOVOIIDSS Se)
LICE IC) IIJDOODE IE)
L |

x
c

[J

Figure 1-1. PLATO Keyboard

97405900 C 1-5

Key Description/Function

NEXT The NEXT key is the most frequently used key on the keyboard. It is located
with the function keys on the right side of the keyboard. Because it is the most
frequently used key, it is designed to be easy to find; it is the only key that is a
different color from all other keys on the keyboard. Pressing the NEXT key tells
the PLATO system you are finished typing an answer to a question, or you are
finished reading the information on the screen display and are ready to see some
additional information on the next display. Pressing the NEXT key instructs the
system to respond to an answer you typed, or to add or erase information on the
screen. Whenever you are in doubt about what to do next when using the PLATO
system, press NEXT.

SHIFT The SHIFT key is used to type the capital letters of the alphabetic characters.
To type a capital letter, hold the SHIFT key down and press the desired
character key. (For example, to type the letter B, hold the SHIFT key down and
press the b key).

The SHIFT key is also used to allow the numeric, punctuation, and function keys
to have two characters or functions. Keys which have two characters or
functions printed on them can produce two characters or perform two functions.
For example, the number five key (5) types the % sign if the SHIFT key is held
down while pressing the 5 key.

There are several keypresses which require using the SHIFT key and another
function key at the same time. These keypresses are always marked by a hyphen
following the SHIFT notation. For example, SHIFT-NEXT means hold the SHIFT
key down and then, while continuing to hold down the SHIFT key, press the NEXT
key.

SHIFT-STOP The SHIFT-STOP keys are used to stop a lesson or sign off from the system.
They are also used during the sign-on sequence. To press SHIFT-STOP, hold the
SHIFT key down and then, while continuing to hold down the SHIFT key, press
the STOP key.

ERASE The ERASE key erases all or part of the information a user has typed on the
screen. Each time the ERASE key is pressed, one character is removed from the
typed response. To remove a complete word, press the ERASE key while holding
down the SHIFT key (SHIFT-ERASE). Erasing always begins with the last word
or character typed.

NEW USER REGISTRATION

All users must be registered in the PLATO system before using the PLATO terminal and system for the
first time. To register, you must give the system two identifiers: a name to identify you; and a group
name to identify either your course of study, the material you should see on the sereen (if you are a
student), or other applicable information such as the group of authors you are working with, the company,
school, ageney, or organization you work for, and so on. Some examples of PLATO names and group
names are: sally/music, john p/pilots, and mary smith/caleulus. PLATO names and group identifiers are
always written in this paired fashion with a slash separating the two identifiers. Either your instructor or
another author or instructor in your group registers this information in the system for you.

1-6 97405900 C

After you are registered, you must give your password! before you can see any lessons or information on
the terminal. Your password is a secret word which allows you to use the PLATO system. It ensures the
system that you are the same person whose identifiers are registered in the system. You choose your
password the first time you use the PLATO terminal. These three identifiers — your name, group, and
password — are called your sign-on. The following defines and describes the three parts of the sign-on.

Name

Group

Password

Your PLATO name is the name you and the person who registers you on the PLATO
system select for you to use when signing on to the system. It can be your full name,
first name, last name, or nickname; or any combination of letters, numbers, or spaces
up to Ps characters. Capital letters are never required in PLATO names and group
identifiers.

Your PLATO group is the name of a file on the PLATO system. It is assigned to you
by your instructor or another author or instructor in your group. Your PLATO group
contains the names of a set of people who have something in common on the PLATO
system, such as taking the same course of study or writing the same lesson.

Each person listed in a group is automatically assigned a user record. Your user record
contains information about you. It tells the system if you are a student, instructor, or
author. If you are a student, your record can tell the system what lessons you are
assigned to study, store information about lessons you have already studied, and save
test scores for your instructor to see. If you are an instructor or an author, your user
record can specify whether or not you can change students' test scores, assign or study
lessons from a catalog, create or destroy user records, write and receive notes, create
curricula, and so on. Generally, your record within the group tells the system what to
do with you.

Your PLATO password is your personal identification to the PLATO system. It is a
secret word that you select or create which makes your sign-on uniquely yours. It tells
the PLATO system that you are the same person who is registered on the system.
Your password should consist of a series of letters, numbers, and/or spaces up to 10
characters. It should be unusual so no one can guess it, and should never be told to
anyone. Do not choose obvious passwords such as your spouse’s name; the name of
your group, employer, or location; your pet's name; your telephone number; a period,
comma, abe, or other words or character strings which can be easily guessed. Choose
something with which only you can identify. Your password is stored in your user
record. No one, however, can see the actual password. It is displayed on the screen as
a random series of x's whenever you type it. Only you and the PLATO system know
what your password is.

Each time you use the PLATO system, you must type your sign-on. You cannot use the system without
typing your sign-on because this identifies you as a registered user of the system. The following are
examples of sign-ons.

Name Group Password
ann miller / algebra / XXxx
jane p / logic / XXXXx
butch /  musie / XXxx

The users’ passwords are not given because users should never tell their passwords to others. Remember,
the PLATO notation for a user's sign-on is name/group (for example, maria/math).

T Not all users are required to have passwords. Students for whom no record keeping is necessary (for
example, young children or special students) are seldom required to have passwords. Instructors decide
whether or not to require students to have passwords on an individual basis.

97405900 C

1-7

THE PLATO TERMINAL
The PLATO terminal you are using is one of three types of PLATO terminals. These are the Information

Systems Terminal
Terminal II (IST-III) (figures 1-2 and 1-3). Each terminal has four major controls.

(IST), the Information Systems Terminal I (IST-I), and the Information Systems
You should know the

loeation and function of each of these controls before using your PLATO terminal. You will not use these
controls every time you use the terminal, but it is important for you to understand how these controls
work in order to correct minor problems that may occur while you are using the terminal.

° ! ERROR INDICATOR

a ON/OFF
@l__ BRIGHTNESS CONTROL

MASTER CLEAR BUTTON
(BEHIND KEYBOARD)

TALK/DATA SWITCH

ERR INDICATOR

all pL a trance CONTROL

ON/OFF

Figure 1-3. IST-II and IST-I Terminais' Exterior

97405900 C

ON/OFF

This switch must be set to ON in order for the terminal to operate. It is not necessary to turn
the switch to OFF when finished using the terminal.

ERROR (ERR) Indicator

This indicator lights during a loss of communication between the terminal and the PLATO
system. Usually, it clears automatically. If the indicator stays lit, press the STOP key or either
the MASTER CLEAR button or RESET switch (depending upon the type of terminal you are
using). If the light still does not go off, press SHIFT-STOP (hold the SHIFT key down while
pressing the STOP key).

MASTER CLEAR Button (IST Terminals)

Pressing this button clears the communication lines. Press the MASTER CLEAR button if the
ERROR indicator lights and the terminal ignores all keyboard and touch panel input. After the
ERROR light goes off, press NEXT to continue the lesson you were working on. If this does not
work, press SHIFT-STOP (hold the SHIFT key down while pressing the STOP key) and then resume
your lesson.

RESET Switch (IST-II and IST-0I Terminals)

The RESET switch performs the same functions as the MASTER CLEAR button. To clear the
communication lines, press the RESET switch momentarily (less than 3 seconds). If this does not
work, press SHIFT-STOP (hold the SHIFT key down while pressing the STOP key) and then resume
your lesson.

BRIGHTNESS Control

This knob adjusts the brightness of the screen display to a comfortable viewing level. To
inerease the intensity of the display, rotate clockwise; to decrease the intensity of the display,

rotate counterclockwise.

If the brightness control is set too high, the display
will be out of focus and the life of some internal
hardware may be shortened unnecessarily.

CONNECTING THE PLATO TERMINAL

The PLATO terminal can be connected to a central PLATO computer or to a local delivery device
attached to the terminal. Terminals connected to a central PLATO computer transmit information
through telephone lines. Terminals connected to a local delivery device transmit information from a
flexible disk drive attached to the PLATO terminal (Micro PLATO). The PLATO IST-II and IST-III
terminals can be used for either the central PLATO system or local (Micro PLATO) delivery.

The following describes how to connect the PLATO terminal to the central PLATO system and the Micro
PLATO system.

97405900 C 1-9

USING THE CENTRAL PLATO SYSTEM

There are two methods of connecting the PLATO terminal to the central computer. These are the direct
connection method and the dial-in method.

With the direct connection method, the terminal is always connected to the central computer. The
terminal has a direct connection if communications lines and equipment directly connect the terminal to a
central PLATO system. You can usually tell that your terminal is directly connected to the system if
there is not a telephone near or next to the terminal. If you are using a terminal directly connected to
the computer, turn the terminal on (if it is off) and proceed with the sign-on sequence. (Refer to How All
Users Sign On, later in this section, to learn how to sign on to the PLATO system.)

With the dial-in connection method, the terminal is connected to the PLATO system by a telephone. You
ean usually tell that your terminal has a dial-in connection if there is a telephone near or next to your
terminal. If you are using a dial-in terminal, turn the terminal on (if it is off) and connect the terminal
according to the following procedure before proceeding with the sign-on sequence. (Refer to How All
Users Sign On, later in this section, to learn how to sign on to the PLATO system.)

1. Do one of the following depending upon the type of terminal you are using.

a. If you are using an IST-II or IST-II terminal, set the TALK/DATA switch to TALK. Dial the
telephone number that connects the terminal to the central computer.

b. If you are using an IST terminal, simply dial the telephone number that connects the
terminal to the central computer. (No TALK/DATA switch exists.)

2. When you hear a constant high-pitched tone, do one of the following (depending upon the type of
terminal you are using).

a. If you are using an IST-II or IST-II terminal, set the TALK/DATA switch to DATA and hang
up the telephone. Each time an IST-II or IST-III terminal is turned on and connected to the
PLATO system, a 2- to 3-minute waiting period occurs as the terminal receives and stores
information from the central system. Messages appear on the terminal screen as the central
system sends information to the terminal. When this process is complete, the Welcome
display appears on the screen.

The 2- to 3-minute waiting period occurs whenever
the terminal is turned on. To avoid this
communieation process from occuring more than
once a day, leave the terminal on all day, turning it
off only in the evening.

1-10 97405900 C

b. If you are using an IST terminal, do one of the following.

1) If your terminal is equipped with a telephone, pull upward on the white key on the left
of the telephone (figure 1-4). This disconnects the telephone handset and connects the
terminal to the PLATO system. Set the handset aside, but do not hang up the telephone.

2) If your terminal is equipped with an acoustic coupler, insert the telephone handset into
the acoustic coupler (figure 1-5). This connects the terminal to the PLATO system.

3. If the red ERROR (ERR) indicator lights, press SHIFT-STOP several times (hold down the SHIFT
key while pressing the STOP key). If the light does not go off after several attempts, press the
MASTER CLEAR button or RESET switch. If the light still remains on, refer to Communication
Errors in section 2.

4. If you hear a busy signal after dialing, all dial-in lines to the PLATO system are busy. Hang up
the receiver, check the number, wait awhile, and dial again.

Refer to the PLATO Terminal User's Guide for more detailed information on connecting specific types of
terminals to the PLATO system.

Figure 1-4. Telephone

97405900 C 1-11

Figure 1-5. Acoustic Coupler

USING THE MICRO PLATO SYSTEM

The Micro PLATO system is a delivery system which can present PLATO lessons on a standard PLATO
terminal without any connection to a central PLATO system. Rather than being connected by a telephone
to a central PLATO computer, the Micro PLATO terminal (either an IST-II or IST-III terminal) is attached
to a flexible disk drive (Control Data's PLATO Flexible Disk Subsystem), figure 1-6. Lessons are stored on
flexible disks resembling 45 rpm phonograph records in their paper covers. To take a lesson, one simply
loads a disk into the flexible disk drive and follows the instructions displayed on the terminal screen. The
Micro PLATO system, like the central PLATO system, instantaneously interacts with users and guides
them through their chosen lesson. The flexible disk (a magnetic medium) operates much like a record. As
it spins in the disk drive, a lesson is transferred to the Micro PLATO terminal.

The following steps describe how to start and use the Micro PLATO system.

1.
2.

3.

Turn on the disk drive.
Push the black button to open the door where the flexible disk will be inserted.
Remove the flexible disk from its protective envelope. |

Insert the flexible disk (label facing you; it should be just to the right of your thumb). Push the
disk until you hear a click.

Close the door by pressing down on the door handle until the door latches.
Turn on the terminal and follow the instructions displayed on the terminal sereen.

Press the RESET switch on the terminal.

97405900 C

Figure 1-6. Micro PLATO Station

The following steps describe how to remove the flexible disk.

CAUTION

Flexible disks should be removed from the disk
drive before turning off the equipment. Turning
off the equipment with a flexible disk inserted
could damage the disk.

1. Open the door.

2. Pull out the disk.

3. Replace the disk in its protective envelope.

97405900 C

Although the actual recording surface of the flexible disk is enclosed in a plastic jacket and then packaged
in a protective envelope, disks must be handled with care. The following procedures will help prolong the
life and ensure the performance of your disks.

e Donot write on labels once they are on the disk jacket.

e Do not attach anything to the disk jacket (especially staples or paper clips).
e Donot touch the disk surface exposed by the jacket slot.

e Do not twist, fold, or bend the disk.

e Donot attempt to clean the disk.

e Keep the disk away from magnetic fields and magnetized materials.

e Protect the disk from liquids, dust, smoke, ashes, and metallic substances.
e Do not eat, drink, or smoke while using the disk drive or handling disks.

e Store the disk in its protective envelope when not in use. (It also is a good practice to store disks
in a closed box or cabinet if they will not be used for more than a short period.)

Published courseware for Micro PLATO delivery will use both the central PLATO and Micro PLATO
delivery methods together. Central PLATO, usually through PLATO Learning Management, will provide
testing and instructional guidance and management through a curriculum. Micro PLATO will be used to
deliver instructional materials. Instructions on when to use Micro PLATO and what disks to use will be
provided through the central PLATO system. An example of how both delivery methods could effectively
be used together is: out of a set of five Micro PLATO stations, four could be used to deliver instructional
materials using the Micro PLATO system while one of the five could be connected to the central system
for instructional management purposes.

HOW ALL USERS SIGN ON

You must sign on to the PLATO system before you can see any lessons, files, or programs or interact with
the PLATO system. The sign-on sequence is the identification exchange between you and the PLATO
system. This exchange determines whether or not you can use the system as well as what you can do once
you are signed on. You must repeat the sign-on sequence each time you use the system. The following
steps describe how to sign on to the PLATO system.

1. When the terminal connects properly to the system, the following message appears on the
screen: Press NEXT to begin (figure 1-7). If the sereen shows anything else, press the BACK key
or the SHIFT-STOP keys (hold the SHIFT key down while pressing the STOP key) several times
until Press NEXT to begin appears on the screen.

Press NEXT to begin

Figure 1-7. Press NEXT to Begin Message

1-14 97405900 C

2. Press the NEXT key. The Welcome display (figure 1-8) appears.

3:58 pm CDT

Monday, September 15, 1988

Welcome to the “mirnc” PLATO System,
a Control Data Customer Service System.

Type your PLATO name. Then press NEXT.

pLato® ie a trademark of Control Data Corporation.

Figure 1-8. Welcome Display

3. Type your PLATO name. As you type, each character appears to the right of the arrow on the
Welcome display. Whenever you see an arrow on a PLATO display, the system is waiting for you
to type a response. Any response you type appears to the right of the arrow. If you make a
mistake, press the ERASE key to erase each letter back to and including the mistake, and retype
your response correctly from that point. When finished, press the NEXT key. The Group Name
display appears (figure 1-9).

4. Type the name of your PLATO group. These characters also appear to the right of the arrow.
When finished, press SHIFT-STOP (hold the SHIFT key down while pressing the STOP key).

97405900 C 1-15

1-16

5.

6.

Type the name of your PLATO group. Then, while
holding down the SHIFT key, press the STOP key.

When you are ready to leave, you should press
these same keys (SHIFT-STOP) to "sign off". -

>

Figure 1-9. Group Name Display

The next display allows you to either select or enter a password, depending upon whether or not
you have previously chosen or been given a password. If you have not chosen or been given a
password, read step 5a. If the person who created your sign-on provided you with a password
(along with your PLATO name and group), or if you previously selected a password, read step 5b.

a.

b.

If you are required to have a password and this is the first time you are signing on, the
Password Choice display appears (figure 1-10, part A). Select your password and type it
carefully. A random number of x's appear to the right of the arrow as you type so no one
ean read your password. When you finish typing your password, press NEXT. Part B of

figure 1-10 appears. Type your password again. Press NEXT and go to step 6.

Not all student and multiple users are required to
have passwords. If this is the first time you are
signing on and the Password Choice display does
not appear, you are not required to have a
password. Go to step 6.

If you have previously selected or been assigned a password, the screen shows the Password

display (figure 1-11). Type your password. Press NEXT. Go to step 6.

If you have forgotten your password, check with
your instructor or the person who registered you in
the system. Your instructor can clear your
password from your user record and allow you to
select another password.

A display appropriate for your user type appears on the screen. If you are a student, either an
index, a lesson, or another planned activity appears on your screen (figures 1-12 and 1-13). If you
are an instructor, the PLATO Facilities display appears (figure 1-14). If you are an author, the
Author Mode display (figure 1-15) appears.

97405900 C

97405900 C

Choose a secret PASSWORD that you will
remember. Do not tell anyone what it is.

Ais you type your password, several X's will
appear so that nobody can see what you are
typing.

Type your password, then press NEXT.

& KKKKAKKAM

Try it again to make sure.

Type your password, then press NEXT.

>

Figure 1-10. Password Choice Display, Parts A and B

Type your password, ther press NEXT.

S

OR... Press the LAB key for additional options.

Figure 1-11. Password Display

PLATO Learning Management

Contemporary Biol ogy

Session # 6
bioi#ia

November 6, 1969
November 6, 1989

Press NEXT to continue

Figure 1-12. Sample PLM Curriculum Display

1-18 97405900 C

Introduction to Arithmetic
Addition
Subtraction
Multiplication

Division

Choose a letter, or press one of these keys:
SHIFT-STOP to sign off

HELP for explanation

Figure 1-13. Sample Student Index

97405900 C 1-19

PLATO Facilities

Group operations (roster, statistics, etc.)
b. Datafiles
Account transactions

Choose a lesson to study

Notes

Interactive communications

Request 4 print

AIDS (information about PLATO and TUTOR)
PLMAIDS (information about the PLM system)

the letter (a-i) of one of the options above.

Press HELP for more information.

Press SHIFT-STOP to leave.

Figure 1-14. PLATO Facilities Display

1-20 97405900 C

AUTHOR MODE

Choose a lesson:

HELP available

Figure 1-15. Author Mode Display

HOW ALL USERS SIGN OFF

You must sign off from the system after completing a session on the terminal to prevent unauthorized use
of your sign-on and of records. To sign off from the system, press SHIFT-STOP (hold the SHIFT key down
while pressing the STOP key) several times until the screen displays Press NEXT to begin. This is the only
way to sign off from the system.

If you are using an IST-II or IST-III terminal, press SHIFT-STOP until Press NEXT to begin appears on the
screen and then set the TALK/DATA switch to TALK.

If you have a dial-in connected terminal, hang up the telephone after you see the Press NEXT to begin
message. Hanging up the telephone disconnects the terminal from the central computer but does not sign
you off from the system.

Turning off the terminal or hanging up the
telephone does not sign you off from the system.
Do not turn off the terminal or hang up the
telephone until you see Press NEXT to begin.

97405900 C 1-21

The security of your records may be jeopardized if you do not sign off from the system before hanging up
the telephone or setting the TALK/DATA switeh to TALK. The PLATO system remembers what lesson or
activity you were working on and connects your sign-on to the next person who dials into the system.
(Some systems automatically sign you off from the system if you forget to sign off before disconnecting
communication lines. However, because not all systems can do this, you should always make sure you
properly sign off from the system before disconnecting communication lines.) If you forget and hang up
the telephone or set the TALK/DATA switch to TALK before signing off from the system, quickly dial
into the system and sign on again. If your lesson or activity appears on the sereen, sign off and then hang
up the telephone or set the TALK/DATA switch to TALK. If you see the message, Sorry, Records Already
Being Used, press SHIFT-HELP (hold down the SHIFT key while pressing the HELP key). Then, press
SHIFT-STOP (hold the SHIFT key down while pressing the STOP key) until you see Press NEXT to begin.

If you dial in to the terminal and another user's lesson or activity appears on your screen, that user has
forgotten to sign off before hanging up the telephone. Press SHIFT-STOP (hold the SHIFT key down while

pressing the STOP key) until you see Press NEXT to begin to sign the user off from the system. Then sign
yourself on to the system by proceeding with the sign-on sequence.

HOW TO CHANGE YOUR PASSWORD
You can change your password any time you sign on to the PLATO system. You should change your
password frequently to avoid the possibility of someone guessing your password and gaining access to your
records. The following steps describe how to change your password.

1. Sign on to the PLATO system as you normally do, typing your PLATO name and PLATO group.

2. ute your old password and press LAB when the system displays the Password display (figure
1-11).

3. Press LAB again. The system displays the Password Choice display (figure 1-10).
4. Type your new password. Press NEXT.

5. Type your new password again to verify it and help you remember it. Press NEXT.

1-22 97405900 C

SECTION 2
USING STUDENT FEATURES

USING STUDENT FEATURES

Curriculum Structures
Taking PLATO Lessons
Understanding Your Curriculum
Index and "mrouter" Lessons
PLATO Learning Management Lessons
Requesting Help
Help Within a Lesson
Help from Your Instructor
TER M-ask
Student Notes
Personal Notes

97405900 C

2-1
2-3
2-3
2-5
2-5
2-7
2-7
2-8
2-8
2-10
2-10

Commenting on Lessons (TERM-comment)
Communicating with Group Members
Handling Problems

Communication Errors

Lesson Execution Errors

Messages You Might Receive
Helpful Tools

Checking the Time

Doing Mathematical Calculations
Additional Student Options
Using the Micro PLATO System

2-i/2-ii

2-12
2-13
2-13
2-15
2-15
2-16
2-16
2-17
2-17
2-18
2-18

USING STUDENT FEATURES 2

a I TE

This seetion presents an overview of the different types of curriculum structures available on the PLATO
system, explains the kinds of activities students are frequently involved in, and defines and describes how
to use the PLATO system features which are usually available to all user types.

All users should read the Introduction (section 1) before reading this section.

CURRICULUM STRUCTURES

PLATO lessons can be presented to students in various ways. Your instructor determines how your lessons
are presented and also how much flexibility to give you while taking lessons. Most students are assigned a
curriculum to study. A curriculum is a study plan which concentrates on a specifie topic or subject.
Curricula usually cover a broad subject area such as English, biology, and so on.

Curricula can be presented to students in different ways, depending upon the method of instructional
delivery selected by your instructor. Some curricula are composed of several modules. A module is a
group of lessons which relate to the same basic subject. Each lesson in the module presents instructional

materials which concentrate on a different area or aspect of the module topic. The eombined lessons and
modules compose the curriculum (figure 2-1).

CURRICULUM

MODULE MODULE

MODULE
LESSON LESSON LESSON LESSON

LESSON LESSON
LESSON LESSON

LESSON

Figure 2-1. Curriculum Structure - Example 1

97405900 C 2-1

Other curricula are composed of several courses. A course is a complete learning package which
concentrates on a specific topic or subject area of the curriculum. It contains modules which can present
objectives, provide instructional lessons, administer tests, and suggest study materials related to the

subject of the course (figure 2-2).

CURRICULUM

COURSE

COURSE

MODULE MODULE MODULE MODULE

@ Objectives e@ Objectives e Objectives @ Objectives

@ Lessons e Lessons e Lessons @ Lessons

e@ Tests @ Tests @ Tests @ Tests

e Study materials e Study materials e Study materials e@ Study materials

Figure 2-2. Curriculum Structure - Example 2

Most curricula are designed to meet the individual needs of students. Curricula can be designed to allow
differing degrees of flexibility to students studying the curricula. Some curricula are designed to allow
you to choose the order of the lessons you want to study, while others require you to study lessons
according to a specified sequence. Some curricula allow you to take a test before studying a lesson.
Usually, if you pass the test, you are not required to study the lesson. Taking a test before studying the
lesson materials helps you know which parts of the lesson you should concentrate on more than others, and
lets you preview the test questions. Many curricula include a statement of the objectives for the modules
and lessons. These objectives help you get a better understanding of the purpose of the lesson and the
information you are expected to know once you complete the course of study.

2-2 97405900 C

TAKING PLATO LESSONS

Before you begin studying lessons on the PLATO system, it is helpful to know what kinds of things you
might be asked to do, or the kinds of things you might want to do.

Press NEXT when you finish reading the information displayed on your screen and are ready to see more
information. Usually, the PLATO system reminds you to press NEXT when finished reading by displaying
Press NEXT or NEXT at the bottom of the screen. Remember, whenever you are in doubt about what to
do on the system, press NEXT.

Sometimes, the PLATO system asks you to answer a question by typing a response. The system tells you
when it expects a typed response by displaying an arrow (>) on the screen. You can respond to a question
in one of two ways. If the system provides a list of options to choose from, type the letter or number in
front of the option. (Occasionally, the system requires you to press NEXT after you make your selection.
If nothing happens after you type your selection from a list of options, press NEXT.) If the system does
not provide a list of options to choose from, type your answer and press NEXT. You should always press
NEXT after you type a response to a question that is not selected from a list of options.

Pressing BACK usually allows you to see displays you read previously in the lesson. Press BACK until you
reach the display you want to see. To return to your original display, press NEXT until you reaeh the
desired display.

Many PLATO lessons contain extra reading material to help you understand different parts of a lesson.
Pressing HELP often displays useful reading material which can further explain a point or procedure
described in the lesson. Press HELP if you need help understanding part of a lesson, if you are interested
in seeing more detailed information about part of a lesson, or if you are unsure of what to do in a lesson.
Usually, the system displays HELP or HELP available at the bottom of the screen to remind you help is
available and to use the HELP key.

Some keys are used only occasionally in PLATO lessons. These keys are ealled function or branching keys
and are used only if the author of the lesson programs them to work. Function keys ean do a variety of
things, depending upon what the lesson author programs them to do. Some keys provide lab exercises or
problems for you to solve, some cause the display to partially or totally erase and add new information,
and others branch you to a new series of displays which give more information about a specifie subject.
Whenever you use a function key, the PLATO system always returns you to the same place in your lesson
from where you first pressed the function key. This prevents you from getting lost while taking a PLATO
lesson. Most lessons tell you which function keys are available for you to use during your PLATO lesson.
This information is usually given at the beginning of the lesson. Other times, individual displays state
which keys are available for you to use. Traditional places to look to find which keys are available are the
bottom two lines of the screen, often in the corners. The following keys are often used as function keys:
LAB, DATA, SHIFT-NEXT, SHIFT-BACK, SHIFT-LAB, and SHIFT-DATA. Remember, keypresses which
have a hyphen following the SHIFT notation (such as SHIFT-NEXT) require you to hold down the SHIFT
key while pressing the function key.

UNDERSTANDING YOUR CURRICULUM

After you sign on to the PLATO system, your display resembles one of the displays in figure 2-3. Each
curriculum has its own set of instructions on how to proceed through the lessons in it. Common locations
for these instructions are the bottom corners of the screen. Find the display in the figure which most
closely resembles your display and follow the instructions printed below it.

97405900 C 2-3

PLATO Learning Management

Contemporary Biol ogy

Name ........ student
Group ....... biolgla

Today's date . November 6, 1988
Last date on ...... November 6, 1989

Welcome back!

Press NEXT to continue

Introduction to Arithmetic
Addition
Subtraction
Multiplication

Division

Choose a letter, or press one of these keys:
SHIFT-STOP to sign off

HELP for explanation

Figure 2-3. Examples of PLM and Index/"mrouter" Lessons

97405900 C

INDEX AND “mrouter” LESSONS

An index display (figure 2-3) lists the lessons in your curriculum which your instructor selected for you to
study. Occasionally, new lessons are added to the index as you complete other lessons. Depending upon
how your instructor designs your curriculum, you can randomly choose lessons to study, choose lessons
according to a specified sequence set by your instructor, or choose lessons according to a specified
sequence and also review lessons previously studied.

PLATO LEARNING MANAGEMENT LESSONS

The PLM Curriculum Welcome display (figure 2-3) appears the first time you sign on to a curriculum using
PLATO Learning Management (PLM). PLM is a system capability designed to direct you through the
curriculum in an individualized manner. It presents tests covering learning objectives, selects learning
resources for you to study, and keeps records of your performance.

PLM gives your instructor several options to choose from in designing your curriculum and also gives you,
the student, a range of options to choose from while studying the curriculum. For example, you can see
the objectives for each module before you begin working in it. You should read the objectives to see what
material the test and learning materials cover. PLM allows you to decide when you want to study the
materials and when you want to take a test. You should take a test before you study to find out which
areas you need to study and receive an assignment for those areas. You can take tests both before and
after you study the learning materials.

PLM curricula frequently give you a choice of learning materials to study. For example, you can choose
to read a textbook, listen to an audiotape, view a videotape, or study a PLATO lesson to achieve one or
more of the module objectives. Many curricula combine these and other types of learning materials for
you to study. The types of learning materials available for you to study and the flexibility to choose one
in favor of another depends upon your instructor and the design of your curriculum.

The PLATO system tells you how to proceed from one part of the curriculum to the next and explains
what you have to do to master each part of the curriculum. You are shown how to see the objectives,
take a test, begin a PLATO lesson, and see how well you are progressing through the curriculum.

The Course Options display (figure 2-4) shows those courses that make up your curriculum. You can
choose to work in any course listed under the heading Courses You Can Work On Now. You pick a course
to work in by typing its number and pressing NEXT.

The Module Options display (figure 2-5) shows the modules in your current course. You can work on any of
the modules listed under Modules You Can Work On Now. Usually there is a small arrow pointing to the
module recommended for study. Some modules may not be available initially, but will become available
after you master one or more of the other modules in the course.

Some modules may be designated as optional modules. You may be required to master one or more of
these modules in order to master the course, but you can choose which optional modules you want to work
on. To work on a module, type its letter or type a number for any of the other options listed at the
bottom of the page.

97405900 C 2-5

2-6

Contemporary Biol ogy

COURSES YOU CAN WORK ON:

[ 1.] Introduction to Biclogical Concepts
2. Survey of Plant Life
3. Survey of Animal Life

THERE ARE 4 COURSES FOR YOLI TO WORK ON LATER.

WHICH COURSE DO YOU WANT TO WORK ON NOW? 2 |

&. See how well you're doing c. Sign off

b. Review instructions

Figure 2-4. PLM Course Options Display

97405900 C

Survey of Plant Life :

WHICH MODULE DO YOU WANT TO WORK ON NOW?

MODULES YOU CAN WORK ON NOU:

Demonstration
Populations ITI
attributes

Bill's Module

1. See how well you're doing 3. Review instructions

2. Work in a different course

Figure 2-5. PLM Module Options Display

REQUESTING HELP

If you have questions or do not understand a procedure while using the PLATO system, you can request
help. Two kinds of help are available to students using the PLATO system. These are: help within a
lesson and help from your instructor.

HELP WITHIN A LESSON

Most lessons contain a standard set of helpful information which is available to all students studying the
lesson. This information can describe a point or procedure in more detail than was explained in the main
part of the lesson, give how-to information, provide definitions and formulas, review previous information,
or give information on how to proceed in the lesson.

97405900 C 2-7

This type of helpful information is usually accessed by pressing HELP. The HELP key is the first key you
should press when you have a question during a lesson. The HELP key is a programmable key, however,
and only works if the author of the lesson programs it to work. (Most authors design their lessons to
include the HELP key). The screen usually displays HELP available or HELP when the HELP key is
working. Pressing HELP during your lesson usually causes one of three things to happen:

e Information can be added to your current display.
e Your current display can be erased and show new information.

e New information can be added to your current display as well as give you the option to see more
information.

After you press HELP and additional information is displayed on your screen, the HELP sequence usually
tells you which keys to press to proceed through the HELP sequence and return to the lesson. If no
instructions are given, press NEXT to proceed. The lesson usually returns you to the display in the lesson
where you initially requested help.

HELP FROM YOUR INSTRUCTOR

You can communicate with your instructor about questions or problems you have while using the PLATO
system. You should request help if you do not understand what you are doing in a lesson or if you feel lost
and do not know what to do next. Three ways to receive help from your instructor are: TERM-ask,
student notes, and personal notes. The following describes each of these features and gives information
on how to use them.

TERM-ask

TERM-ask allows you to request assistance from an instructor about a lesson you are studying while you
are actually studying the lesson. It allows you to communicate with your instructor by typing messages
back and forth on the bottom of your screen. The TERM-ask feature is not automatically available to all
students. If your instructor arranges for your PLATO group to use TERM-ask, you can access the feature.

The following steps describe how to use TERM-ask.

Contacting Your Instructor

1. Press TERM (hold down the SHIFT key while pressing the TERM/ANS key). The system displays
what term? >.

2. Type ask and press NEXT. One of the following messages is displayed.
a. If you see the message Someone Has Been Notified, your instructor is available to answer
your question and has been notified you called. You can continue your lesson while you wait

for a instructor to contact you (it usually takes a few minutes for your instructor to
reply).

2-8 97405900 C

b. If you see the message Sorry, No One Is Available, your instructor is not currently available
to answer your question. In some cases, you are given the option to write a note to your
instructor. If you want to write a note, follow the instructions in Commenting on Lessons,
later in this section, for more information on how to write a note.

c. If you see the message Sorry, Your Group is not Prepared for TERM-ask, your instructor has
not prepared your group to use the TERM-ask feature.

3. When your instructor contacts you, you see a message such as sally jones/teacher/minna also sees
this display.

Communicating with Your Instructor

When your instructor contacts you, you can communicate by typing messages back and forth on the
bottom of your screens. Your instructor can also monitor your screen and see the same information on
his/her screen that you see on yours. (This eliminates the need for you to describe in detail where you are
in the lesson and what problems you are having.)

The following steps describe how to communicate with your instructor.

1. When your instructor contacts you and you see a message such as sally jones/teacher/minna also
sees this display, an arrow appears in the lower left corner of your screen. Any message your
instructor types appears to the right of this arrow.

2. To communicate with your instructor, press TERM (hold down the SHIFT key and press the
TERM/ANS key). A second arrow appears. Any message you type appears to the right of this
arrow.

3. Type your message. Your instructor sees the message as you type it. If your message requires
more than one line of typing, press LAB to clear the line and continue typing. The LAB key is
the only key that allows you to continue typing. If you press a function key other than LAB (for

example, BACK), the arrow disappears. If you accidentally press another function key and the
arrow disappears, press TERM to recall the arrow and resume typing.

Showing Your Instructor Your Screen Display
Although your instructor can monitor your screen, your instructor cannot see the display you are looking
at when she/he initially contacts you. If the display you want your instructor to see is the one you are
looking at when your instructor contacts you, you need to replot your screen (replot means to move from
one display to another).
The following steps describe how to replot your screen.

1. Do one of the following, depending upon the type of lesson you are using.

a. If you are using PLM, press DATA.

b. If you are taking an index lesson, either press BACK and then NEXT, or press HELP and then
NEXT.

2. Press TERM to talk to your instructor while she/he monitors your screen.
3. To show your instructor a different display, press BACK to discontinue typing. Go to the new

display and repeat step 2.

97405900 C 2-9

Ending TER M-ask

When your questions are answered and you do not want further help, type thanks or bye. (Remember to
press the TERM key to type messages.) Your instructor ends the communication. The system displays a
message telling you TERM-ask is over.

Student Notes

Student notes are part of a special notes system which allows you to read and write notes to your
instructor, group members, and other system users. Your instructor determines whether you can use the
notes feature and also which options you can use. Your instructor allows you to write and receive notes,
write notes only, or receive notes only. Your instructor also determines with whom you can communicate
(for example, group members, all system users, or just your instructor). To find out if you have access to
the student notes feature, look at your lesson index. You have access if notes is one of the options listed,
or if a function key is designated as a notes option. If nothing is listed, your instructor has not arranged
for you to use student notes.

You should write a student note to your instructor if you have a comment or a question about the lesson
you are taking, or to respond to a note written to you. The following steps describe how to read and write
student notes.

1. If notes are listed as an option, either type the letter in front of the option or press the
designated function key associated with the option.

2. Doone of the following, depending upon the kind of activity you want to do.
a. Press NEXT to read your notes.
b. Press SHIFT-LAB to write a note. Go to step 3.
ec. Press BACK to return to the index.
If your instructor sends you a note, the system tells you there are notes you have not read.

3. After you press SHIFT-LAB, the system displays a rectangular box with an arrow in the upper
left corner. Directions for writing a note are listed under the box. Press HELP for information
on how to write notes.

Personal Notes

Personal notes are private notes between individuals on the system. You can write personal notes to or
receive personal notes from any user whose group is prepared to receive notes, provided your instructor
has granted you a personal notes option. To find out if you can use the personal notes feature, look at
your lesson index. You can access personal notes if personal notes is one of the options listed, or if a
function key is designated as a personal notes option. If nothing is listed, your instructor has not arranged
for you to use personal notes.

2-10 97405900 C

The following steps describe how to write personal notes.

1.

3.
4,

5.

Follow the instructions on your lesson index to access personal notes. (Function keys and lesson
oo vary with different curricula.) The system displays the Personal Notes display (figure
2-6).

Type the name of the person to whom you want to send a personal note. Press NEXT.
Type the name of the group in which the user is registered. Press NEXT.

Type the name of the PLATO system in which the user is registered (for example, minna, minnb,
minne, and so on). Press NEXT. If the user is registered in the same PLATO system as you,
press NEXT before typing the system name. Your system is automatically recorded. (The name
of your PLATO system is on the Welcome display.) The system displays the Personal Notes Text
display (figure 2-7).

Read the instructions printed at the bottom of the display. Press HELP for more information on
how to write notes.

Refer to Using Personal Notes in section 4 for more information on reading and writing personal notes.

PERSONAL NOTES

Press: LAB to read your notes
DATA for other options
HELP for explanation and policy

To whom do you wish to send a note:

Name >
Group

System

Figure 2-6. Personal Notes Display

97405900 C 2-11

Personal note to mary smith / medicine /” minne

for the next line SHIFT-NEXT when finished

for the previous line SHIFT-BACK to exit and not send
to change the line SHIFT-LAB to insert a line

for more directions SHIFT-HELP to delete lines

Figure 2-7. Personal Notes Text Display

COMMENTING ON LESSONS (TERM-comment)

You can comment on a lesson you are taking by using the TERM-comment feature. TERM-comment
allows you to comment on a lesson and send the comment to the lesson author (your instructor can also
receive the comment). Some reasons to comment on a lesson are: unclear instructions, confusing
explanations, incorrect answers or information, and so on. Write a comment about areas in a lesson which
you do not understand or find confusing, to ask for a clarification of a technical point, or to start a
discussion. Keep your comments brief but be specific. The following steps describe how to use
TERM-comment.

1. Press TERM (hold the SHIFT key down while pressing the TERM/ANS key). The system displays
what term? >.

2. Type comment and press NEXT. The system responds with a message saying your comment will

be sent to the lesson author or your instructor. A line appears at the bottom of the screen with
an arrow below it on the left side. Some instructions are below the line.

2-12 97405900 C

3. Press HELP for more information on using TERM-comment.

4. Type your message. Press NEXT at the end of each line to continue typing. The maximum
length of a comment is usually 20 lines.

5. Press BACK to read or correct lines previously typed. Continue pressing BACK until the line you
want to read or change appears to the right of the arrow. Use the ERASE and EDIT keys to
ey your corrections (refer to appendix A for information on how to use the ERASE and EDIT
keys).

a. If your comment is completed after making your corrections, go to step 6.

b. If you want to continue typing after making your corrections, press NEXT repeatedly until
you return to the last line of your comment. Finish typing your comment and go to step 6.

6. Press SHIFT-NEXT to send the comment to the lesson author, or press SHIFT-BACK to cancel
the comment.

COMMUNICATING WITH GROUP MEMBERS

As a student, you can communicate with other members in your group, other groups, or all users on the
system through general notes. General notes are a collection of notes written by members of a defined
community about a particular subject. The notes are grouped into sets. Each set concentrates on a
specific topic and is identified by a name. General notes allow group members to share and exchange
ideas and comments about specific topics or subjects of interest.

Not all general notes can be read or responded to by all users, however. Many directors of notes restrict
the access to notes to certain PLATO groups only. Some directors allow you to read and write notes, but
others allow you to do only one or the other. The system allows you to see only the notes you can access.

Your instructor can tell you which general notes, if any, you can access. The notes option appears on your
lesson index if you can access one or more general notes. To learn how to access and participate in
general notes, refer to Using General Notes in section 4.

HANDLING PROBLEMS

Occasionally, problems might occur while you are using the PLATO system. Common problems are
distorted screen displays, lines across the screen, or random letters or symbols displayed on the screen.
Most of these problems are minor and can be easily corrected by you. The following describes the
different types of problems you might experience (infrequently) while using the PLATO system. Table 2-1
provides a list of suggested ways to solve problems.

97405900 C 2-13

TABLE 2-1. TROUBLESHOOTING PROCEDURES

Telephone rings, but you get a busy signal
or no answer.

Hang up and try egain.

Call PLATO hot line.

1. Press either MASTER-CLEAR button or

Terminal writing appears upside down,
RESET switch briefly.

backwards, and so on.

Press STOP key until light goes off.

Red ERROR (ERR) light goes on and you get
no response from typing or touching the
screen.

2. Press either MASTER-CLEAR button or
RESET switch briefly.

Hang up and redial.

4. Call the PLATO hot line.!

Red ERROR (ERR) light goes on frequently, 1. Hang up, wait 5 minutes, and redial.

causing screen display errors.

2. Call the PLATO hot line.!

Wait for one of these messages:

PLATO OFF message appears on screen.
a. Press NEXT to begin.

b. PLATO will return in 5-10 minutes.

e. PLATO will return at_—~ hours
Central Time.

1. This message is occasionally seen at
night and indicates that routine
maintenance is being done.

PLATO not available message appears on
sereen.

tThe PLATO hot line number is 1-800-328-9114 or 612-482-2006 (Minnesota only).

If you call the hot line, give the staff person on duty the following information:

e Your name and location.
e Your dial up number.

eA description of your problem.

2-14 97405900 C

COMMUNICATION ERRORS

Communication errors result when there is interference on the communication lines between the central
computer and your PLATO terminal. These are usually minor problems which you can correct. When a
communication error occurs, the display on your screen is often distorted. Sometimes lines may be drawn
across the screen, sentences may be upside down or sideways, or text may be written backwards. Usually,
the ERROR (ERR) indicator lights when there is a communication error. To correct the error and clear
the communication lines, press STOP. If the problem persists after you press STOP one or two times, do
the following.

e Press the MASTER CLEAR button or RESET switch. This clears both the communication lines
and your screen. When your screen clears, it is blank. Press DATA or BACK to return to your
lesson.

e Press SHIFT-STOP if the problem persists after you press MASTER CLEAR or RESET. In
addition to clearing the communication lines, SHIFT-STOP can also sign you out of your lesson or
off from the system. If you sign off while clearing the communication lines, sign on again and
continue your lesson.

Communication errors can also cause graphic displays to be distorted. Graphic displays are usually
pictures, animated characters, boxes or lines around text, or extra large or small printing. If a graphics
display does not look right (lines appear across text, figures drawn over text, upside down letters, and so
on), press the TERM key (hold the SHIFT key down and press the TERM/ANS key), type charset, and press
NEXT.

LESSON EXECUTION ERRORS

Lesson execution errors occur when the author of a lesson does not design the lesson to accommodate
every response or keypress you might select. Because PLATO lessons are so flexible and offer many
options to you as a student, the author of your lesson might not anticipate or plan for all the responses or
keypresses you might choose. When you ask the system to do something which the author of the lesson did
not think you might do, a lesson execution error might occur.

When a lesson execution error occurs, the system displays a lesson execution error message. This message
tells you a lesson execution error has occurred and gives some information on the place in the lesson
where the error occurred. When the system displays the lesson execution error message, it gives you the
option of writing a note to the lesson author. You can help the lesson author determine how to correct
the problem by recalling what happened right before the lesson execution error occurred. For example,
state which key(s) you pressed, the response you typed, or anything unusual you did or noticed.

97405900 C 2-15

This information helps the author find the problem and correct it quickly. The following steps describe
how to write a note to the lesson author after a lesson execution error occurs.

1. At the bottom of the lesson execution error message, find the line with an arrow below it on the
left side. Read the instructions printed below the line.

2. Press HELP for additional instructions.
3. Type your note. Press NEXT at the end of each line to continue typing.

4. Press BACK to reread or correct lines previously typed. Continue pressing BACK until the line
you want to read or change appears to the right of the arrow. Use the ERASE and EDIT keys to
make your corrections.

a. If your note is finished after making your corrections, go to step 5.

b. If you want to continue typing after making your corrections, press NEXT repeatedly until
you return to the last line of your note. Finish typing your note and go to step 5.

5. Press SHIFT-NEXT to send the note to the lesson author, or press SHIFT-BACK to cancel the
note.

MESSAGES YOU MIGHT RECEIVE

Sometimes problems can occur within the PLATO system which cause the system to stop working. These
system problems are called crashes. When the system crashes, the system stops working and displays a
message indicating the system is off. If you are using a terminal when the system crashes, the lesson or
activity you are working on stops and the system ignores all keyboard input. A PLATO Off sign appears at
the top of your screen. A few minutes later, the screen erases and the PLATO Down display appears. The
PLATO Down display usually tells you what time the system is expected to be working again.

Whenever the PLATO system is down, you do not have to sign off from the system. The system
automatically signs off all users when it crashes. This is the only time you do not have to sign off from
the system after you have signed on.

If the PLATO system is temporarily unavailable, a message is displayed which usually indicates what time
PLATO services are scheduled to resume.

HELPFUL TOOLS

Some PLATO system features are similar to reference materials. These features can perform
mathematical calculations or tell you the correct time of day and the current date. Students can use
these features any time they are using the PLATO system. Occasionally, an author may inhibit these
features from working in a particular lesson. For example, the feature that calculates mathematical
expressions is usually turned off in math lessons. The following paragraphs describe these features and
how to use them.

2-16 97405900 C

CHECKING THE TIME

You can request to see the current time and date by using the TERM-time feature. TERM-time displays
the current time and the day, month, and year. To use TERM-time, follow these steps.

1. Press TERM (hold the SHIFT key down while pressing the TERM/ANS key). The system displays
What term? >.

2. Type time and press NEXT. The system responds by displaying the current time and date at the
bottom of the screen.

DOING MATHEMATICAL CALCULATIONS

You can use the PLATO system to do mathematical calculations by using the TERM-cale feature.
TERM-cale allows you to present mathematical problems or equations for the system to solve.

You need to understand how the PLATO system solves mathematical equations before using the
TERM-cale feature in order to set up your equations correctly and prevent the system from
misunderstanding your equation. The rules for mathematical equations are basically those of ordinary
arithmetic. The order of operations from first to last is:

1. Exponentiation.

2. Multiplication.

3. Division.

4, Addition and subtraction.
All equations are solved in this order.
Use parentheses liberally when writing your equations to make your expressions clear. For example, 6-2x3
is read by the PLATO system as 6-(2x3) instead of (6-2)x3, since multiplication is done before division.
Parentheses ensure your equation is interpreted correctly.
The function keys on the left of the keyboard contain the four operations: x, /, +, - (multiply, divide, add,
and subtract, respectively). An asterisk (*) is equivalent to x and a slash (/) is equivalent to +.
Exponentiation is indicated with two asterisks (3**2=9) or superscripts (22#4=16). Use the SUPER key
to type one superscript at a time and SHIFT-SUPER to type more than one superscript at a time. Press
SHIFT-SUB to return to the normal typing line.
To use TERM-ceale, follow these steps.

1. Press TERM (hold down the SHIFT key while pressing the TERM/ANS key). The system displays
What term? >.

2. Type cale and press NEXT. The system displays an arrow at the bottom of the sereen.

3. Type your calculation but do not include an equal (=) sign. Press NEXT. For example, (250 +
250) - 250 NEXT. The system gives the correct answer (for example, 250).

4. Press NEXT to enter another expression.
5. Press BACK to return to your previous activity.

To learn more about TERM-cealc, study the on-line PLATO lesson "Stermcale".

97405900 C 2-17

ADDITIONAL STUDENT OPTIONS

Not all users with student sign-ons study curricular materials. Some users are assigned student sign-ons
for direct access to specific system features. The following defines some of the features available to this
type of student user and references the sections of the manual which contain detailed information about
the feature.

Doeumentor

Doeumentor is a utility on the PLATO system which can be used as a tool for organizing, editing,
and presenting text-oriented information. Refer to Using Documentor in section 4 for more
information about documentor.

Catalog of Available Courseware

The Catalog of Available Courseware is a reference catalog which contains a listing of all
published PLATO lessons. Refer to Using the Catalog of Available Courseware in section 4 for
more information on this feature.

Print Requests

This feature allows you to request a print of a file on the PLATO system, check the status of a
request previously made, and check the availability of the line printer. Refer to Requesting
Prints in section 4 for more information on this feature.

On-Line Author Listing

This feature allows you to see a listing of all authors on the PLATO system and biographical
information about them. Refer to Using the On-Line Author Listing in section 4 for more
information about this feature.

USING THE MICRO PLATO SYSTEM

The following information is directed toward
students who are studying lessons using the Micro
PLATO system.

The Micro PLATO system is an easy and uncomplicated system to use. It is similar to using the central
PLATO system in that the keyboard functions the same and lesson use is the same (directions are provided
to guide you through the lessons). Although the two systems have some similarities, they also have some
differences. Unlike the central PLATO system, the Micro PLATO system does not always require you to
sign on or identify yourself before seeing lessons on the terminal. Lessons are contained on flexible disks
and different lessons often have different requirements. Some lessons may require you to sign on before
using the system, while others will simply present the lessons. (Usually, if you are required to sign on, it
means information is being collected about your performance or use of the lesson; for example, if you are
taking a test).

2-18 97405900 C

An index will usually be displayed which contains titles of lessons and possibly several options to choose
from. Because each lesson is unique, its structure or form of presentation may differ from other lessons.
However, as with using the central PLATO system, the lesson instructs you on how to proceed through the
lesson.

Some lessons may require you to take a test on the material presented. Usually, testing is done using the
central PLATO system since that system can collect and store data (such as test scores for large numbers
of students). Your lesson or other instructional materials will instruct you as to whether or not you need
to sign on to the central PLATO system to take a test or participate in another activity. (Refer to How
All Users Sign On in section 1 for information on how to sign on to the central PLATO system.)

Because the Micro PLATO system is not connected to the central PLATO system and uses lessons which
are contained on flexible disks, several features which are available on the central PLATO system are not
available on the Micro PLATO system. Some of these features include TERMS (TERM-ask, TERM-cale,
and so on) and notes.

97405900 C 2-19

SECTION 3
USING INSTRUCTOR FEATURES

USING INSTRUCTOR FEATURES

Introduction
PLATO Facilities Display
Using AIDS
The PLATO Group
Group Data
General Group Information
Associated Files
Security Codes
Group Security
Group Operations
Registering Students
Deleting Students
Deleting Individual Students
Deleting All Students
Listing All Group Members
Inspecting/Changing Student Records
Leaving a Message
Using Your Account
Reviewing and Studying Lessons
PLATO Courses and Curricula
Designing a Curriculum

97405900 C

Instructional Management Tools
How to Use Index Lessons
How to Use the System-Supported
Router ("mrouter")
How to Use PLATO Learning
Management (PLM)
How to Use Your Own Router
Using Datafiles
Additional Instructor Options
Monitoring Group Members
Specifying Group Data Collection
Templating Records
Copying Records from Another Group
Creating Instructor Records
Managing Personal Notes
Using Notes
Using Interactive Communications
Requesting Prints
Receiving Help
Giving Help

3-i/3-ii

3-24
3-26

3-26

3-33
3-33
3-33
3-35
3-35
3-36
3-36
3-37
3-37
3-39
3-39
3-39
3-40
3-40
3-41

USING INSTRUCTOR FEATURES 3

a ag NS IE IT II a

This section presents an overview of the functions of an instructor and defines and describes the PLATO
system features available primarily to instructors. All users should read the Introduction (section 1) and
Using Student Features (section 2) before reading this section.

INTRODUCTION

An instructor using the PLATO system has responsibilities similar to those of an instructor in a
conventional school setting. Instructors using the PLATO system create class rosters by registering
students in a group, review PLATO lessons for applicability and content, choose and assign lessons for
students to study, arrange lessons into curricula or assign published curricula, select appropriate
instructional management tools, monitor the progress of each student and the class as a whole, and answer
students questions as they proceed through the learning material.

PLATO FACILITIES DISPLAY

As an instructor, the first display you see after signing on to the PLATO sytem is the PLATO Facilities
display. This display is your navigational tool on the system. All the system resources you need can be
reached from this display. The PLATO Facilities display routes you to displays where administrative tasks
can be performed. v

The complete PLATO Facilities display consists of nine options. Not all instructors have all nine options
available to them. Only the options which are assigned to you appear on your PLATO Facilities display.
Your account director determines which options are most applicable to the tasks you need to perform on
the system and assigns those options to you. Some users have subsets of or limited access to some
selected instructor features and options. This type of user is typically a training coordinator,
administrator, teaching assistant, or someone assisting an instructor or training director in some
capacity. Consult your account director to add needed options to your display. The nine PLATO
Facilities display options are shown in figure 3-1.

97405900 C 3-1

PLATO Facilities

a. Group operations (roster, statistics, etc.)
b. Datafiles
Recount transactions

Choose a lesson to study
Notes
Interactive communications

Request a print

AIDS (information about PLATO and TUTOR)
PLMAIDS (information about the PLM system)

the letter (a-i) of one of the options above.

Press HELP for more information.

Press SHIFT-STOP to leave.

Figure 3-1. PLATO Facilities Display

Each of the nine options on the PLATO Facilities display allows you to access different system features
and do different things. The following briefly describes the general types of functions associated with
each option.

Group Operations

This option allows you to perform group management tasks. You can enroll students in or delete
students from a group, see and change individual student and group records, leave messages for
students, and assign lessons to students with this option.

Datafiles

This option allows you to collect and examine supplemental data (usually formative in nature) on
lessons in your curriculum which have been coded to allow data collection. Some examples of the
kinds of data you can examine are: student requests for help; answers students gave to questions in a
lesson (both correct and incorrect); and information on lesson execution errors.

97405900 C

Account Transactions

This option allows you to see what files are in an account (a named file on a PLATO system through
which all contracted PLATO resources are managed by a designated account owner or designee), add
files to or delete files from the account, look at statistics on lessons, inspect the log of account
transactions, see what users in the account are signed on to the PLATO system, and create files.

Choose a Lesson to Study

This option allows you to study a PLATO lesson as a student, as well as access on-line reference
materials.

Notes

This option allows you to communicate with other users on the PLATO system through notes. You
can write personal notes to and receive notes from other PLATO system users. You can also read
notes from systems personnel regarding the current status of the PLATO system, as well as read,
write, and respond to notes in general notes files. This option also allows you to write notes to and
receive notes from students in your group.

Interactive Communications

This option allows you to see a list of people currently using the PLATO system, talk to someone by
using the TERM-talk feature, set your TERM-talk feature options, and respond to requests for help
from students or others using TERM-ask.

Request a Print

This option allows you to request prints of files (lessons, notes, and so on). You can also check the
status of a print request previously made and check the status of the printer.

AIDS

This option allows you to access the on-line reference manual for the PLATO system. AIDS contains
helpful reference information on PLATO system features for both authors and instructors.

PLM AIDS

This option allows you to access the on-line reference manual for PLATO Learning Management
(PLM). PLM AIDS contains helpful reference information on PLM features for both authors and
instructors.

To select an option on the PLATO Facilities display, type the letter in front of the desired option (for
example, to look at the first option type the letter a). The system then shows you a new display giving
you more detail about that option or a list of other options. Whenever you have selected an option from
the PLATO Facilities display, you ean usually return to the PLATO Facilities display by pressing BACK.

97405900 C 3-3

USING AIDS

AIDS is an on-line reference manual for authors and instructors. It contains definitions and explanations
of most of the PLATO system features and all of the PLATO Author Language commands. AIDS contains
more than 100 lessons which collectively form a complete reference manual for the PLATO Author

Language and system features. Authors and instructors frequently use AIDS as a reference tool when
using the PLATO system. The following are examples of the kind of information available in AIDS.

e List of indexes in AIDS.
e Definitions of PLATO system terminology.
e Descriptions of PLATO system features and lessons.
e@ Names and descriptions of useful lessons.
e Information on PLATO publications
e Suggestions on writing, testing, and evaluating lessons.
Users should refer to AIDS whenever they need more information about a specifie feature. To use AIDS,
choose the AIDS option from the PLATO Facilities display by typing the letter in front of AIDS. The
system displays the AIDS Title display (figure 3-2). From the AIDS Title display, you can do one of three
things, depending upon your needs.
e Press HELP for more information on AIDS and how to use it.
e Press NEXT for the AIDS Index (figure 3-3). The AIDS Index consists of two displays (press
NEXT for the second display, BACK to return to the first) that function like a table of contents.
It presents a general overview of the information covered in AIDS. Choose an option which
generally covers the information you want to see by typing the letter in front of the option. The
system either displays the information or a second, more detailed, index.
Press HELP for more information on how to use the AIDS Index.
e Press DATA to bypass the index and see a display which allows you to request information on a
specific command or feature (figure 3-4). Type the name of the command or system feature on
which you want information and press NEXT. The system displays the information.

For a quick reference, you can press DATA from anywhere in AIDS, see the What TUTOR
Feature display, and request information.

Press HELP from the What TUTOR Feature display for more information on how to use this
display.

3-4 97405900 C

97405900 C

Pi. da. ds

Elaine Avner, Darlene Chirolas,
Celia Davis, Jim Ghesquiere,
Tina Gunsalus, Jim Kraatz
and Judy Sherwood

PSO Author Group -- CERL
Univ of Illinois, Urbana

Press HELP if this is your first time in lesson AIDS

(©) Copyright, 1973, 1974, 1975, 1976, 1977, 1978, 1979, 1988
Board of Trustees of the University of Illinois

NO portion of the AIDS lessons may be reproduced
in a form without permission from the authors.
1241 features requested per day for the last 698 days

Figure 3-2. AIDS Title Display

3-5

Press a letter; or press NEXT for page 2

a Aids for new authors
How to use AIDS

Author Resources

Alphabetical list of TUTOR commands
Functional lists of TUTOR commands
List of Indexes in AIDS

Lists of System Defined Variables, Keynames,
Functions, Logical & Bit Operators, -specs- Tags

Making Displays

Making Graphs & Charts
Calculations and Variables
Conditional Operations
Sequencing

Judging

Execution of TUTOR

SHIFT-BACK always returns you to an index display.
HELP, DATA, BACK, SHIFT-NEXT are always available.

Figure 3-3. AIDS Index Display (Sheet 1 of 2)

97405900 C

Press a letter; or press NEXT for page 1

The PLATO Computer

Special Characters: ACCESS Characters, Linesets,
FONT Characters (Character Sets), & MICRO Keys

>

Student Data, Instructor Options, & Routers

Keynames, Keycodes, and Internal Codes

Programming Errors and Condense, Lesson, &
Execution Errors

Library of Author Routines
Microfiche and Photographing the Plasma Panel
Systematic Lesson Design

The Programmable Terminal (PTs & ISTs)

SHIFT-BACK always returns you to an index display.
HELP, DATA, BACK, SHIFT-NEXT are always available.

Figure 3-3. AIDS Index Display (Sheet 2 of 2)

97405900 C

What TUTOR feature? >

HELP for help on how to use AIDS

SHIFT-BACK for Main AIDS Index

SHIFT-DATA to make a comment about AIDS

Figure 3-4. AIDS What TUTOR Feature Display

The following are some suggested topics for new instructors to refer to in AIDS to familiarize themselves
with some system features and instructor responsibilities.

® groups

e curriculum design
e courseware catalog
e notes

e  term-ask

THE PLATO GROUP

The primary working tool of an instructor using the PLATO system is the PLATO group. This section
defines the PLATO group, explains its functions and the information it contains, lists the operations that
can be performed within the group, and gives step-by-step instructions on group operations.

3-8 97405900 C

A PLATO group consists of a set of users who share something in common in their involvement with the
PLATO system. For example, users who are studying the same course or curriculum, writing lessons for
the same course or organization, or involved in the same project might be members of the same group.
Each of these users is registered with the system in the same file. (A file is a finite set of data. The
space within the computer system which stores this data is called file space. A file space is a defined,
finite subset of a file.) Instructors use the group file to perform two major functions: student enrollment
and monitoring; and curriculum design and management. The group file contains general information
about the group and records for each of the users in the group.

A group file must be created before users can be enrolled in a group. To create a group file, an instructor
contacts the account director of his/her account and requests that a group file be created. (Some
instructors can edit their accounts and therefore can create the files they need themselves.) Creating a
group file involves acquiring the needed file space and imposing safeguards against misuse of the space
and its eventual contents. Once this is done, the instructor can begin enrolling the group members in the
group file.

As an instructor, there are a number of administrative procedures you can perform concerning your
group. You can register someone in the group, delete someone from the group, see a list of all group
members, leave a message for one or more group members, see a list of all group members currently using
the system, change or clear a student's password if she/he has forgotten it, and see statistical information
of group members on an individual or group basis.

To see information on an individual student, you need to refer to that individual group member's user
record. The system maintains a user record for each person enrolled in the group. Each record varies
according to which user category the group member belongs (for example, student, author, instructor,
multiple), although some information is the same for all user eategories. User records contain such
information as: user's name, date user last used the system, number of days user has signed on, and user
category. In addition to providing information, the user record allows the instructor to change the record,
leave a message to the user, change a course of study, or see the lessons the user is completing.

GROUP DATA

Information about your group is contained in the Group DATA (Directory) display (figure 3-5). From the
Group DATA (Directory) display, you can see general information about the group file, a listing of other
files associated with your group file, and information about the kinds of security codes required to use the
group file. This information can be accessed from one of three options on the Group DATA (Directory)
display - Group Information, Associated Files, or Security Codes.

97405900 C 3-9

medicine (1 part) Disk pack -- caemast
Starting date --- 89/19/88 Account ---- cbhedd

99719788 12.57.26
jean price of adev
1-15

Change group codeword

Type the appropriate letter: >

Group information
Associated files

Security codewords

HELP available

Figure 3-5. Group DATA (Directory) Display

To reach the Group DATA (Directory) display, choose the Group Operations option from the PLATO
Facilities display by typing the letter in front of the option and then press DATA. The following describes
the three options available on the Group DATA (Directory) display.

General Group Information

This option contains general information relating to the group. The information ineludes the name of the
person responsible for the group file, the general kinds of users who are registered in the group, and a
short description of the purpose of the group.

To enter or change general group information, type the number in front of the information to be changed.
Type the information and press NEXT. Identifying the group owner and audience is particularly important
whenever more than one person ean edit an account. Proper identification of the group owner, audience,
and purpose of the file can prevent accidental deletion of the file by an account director.

3-10 97405900 C

Associated Files

This option allows you to identify other files your group uses or relies upon. Associated files can set up
routers for your curriculum, allow students to write comments about lessons, allow students to write
personal notes, and so on. Some examples of the kinds of files you might use with your group file are:

Student notes file

Datafile

TERM-ask group

Processor lesson

Router

Instructor file

A file in which students write notes to the instructor or comment on
lessons they are studying.

A file instructors ean use to see data collected on lessons and the
curriculum. It is usually used for formative evaluation purposes.

A group identified as containing sign-ons of authors and/or instructors
who are available to answer questions of students whose records reside in
the group being edited/inspected.

A user-written editor to be used instead of the standard PLATO system
group editor. (An editor is simply a lesson used to insert, change, or
inspect information. The PLATO group is an editor because it is a lesson
used to create, inspect, or change information about students. The
options in the PLATO group editor are a general set, anticipated to be
needed by most system users. When more flexibility than the standard
PLATO group editor can provide is required, a special editor can be
written. This lesson is identified as the processor lesson because it is used
to process the group's data. Most instructors will never need to designate
a processor lesson.)

A file that controls lesson sequencing and lesson selection and generally
direets decisions regarding a student's progress through a curriculum.
Either the PLATO system router, "mrouter", or PLM (PLATO Learning
Management), can be used or a new router can be developed to meet a
particular set of curriculum needs.

A file that contains curriculum design and course catalog information. An
instructor file is used only if the group uses the PLATO system router,
"mrouter"; otherwise, the designed curriculum and catalog must be
incorporated in the router lesson (unless PLATO Learning Management
curriculum files are used). An instructor file ean be used by more than
one group at a time.

To attach a file to your group file, type the number in front of the desired file, type the name of the file
you want to attach, and press NEXT.

Security Codes

This option allows you to set the security codes for your group file. You can choose to allow only group
members to see and/or change the file, allow only account members to see and/or change the file, or limit
access to you or a set of people who know a typeable security code. This option also allows you to choose
whether or not to allow systems personnel access to the file. Refer to the following paragraphs to learn
more about security codes and group security.

97405900 C

3-11

GROUP SECURITY

Instructors and account directors are responsible for group security. The group file contains both general
and specific information about the group as well as student records. Because of the confidential nature of
student data, it is important to control which users are allowed to see the file. Just as instructors in
conventional instructional and training settings do not want students or others to see their gradebooks or
student records, neither do instructors on the PLATO system.

Group information and student records can be kept secure and confidential by using codewords.
Codewords are similar to passwords in that the system cheeks the codeword assigned to a group file
before allowing users to see or change the file. The person who creates the group file is initially
responsible for assigning codewords to the file. If someone creates a group file for you, you should change
the codewords the first time you use the file so only you know the codewords to the file. Codewords can
be set to allow some users or specific user types different kinds of access to the group file. For example,
you ean set the security codes of a group file to allow only you to see or change the file, or allow only
authors and instructors within your group or account to see and/or change the file. The following are
examples of different types of security codes you can set.

Typed code Requires all users (except instructors in the group) to type the security
codeword in order to see and/or change the group file.

GROUP code Allows all authors and instructors listed within the group to see and/or
change the group file without typing a codeword.

ACCOUNT code Allows all authors and instructors in groups listed within an account to see
and/or change the file without typing a codeword.

Unmatchable code Prevents all users (except instructors in the group) from seeing or
changing the file.

It is important to be creative when assigning codewords. If you use a typed code, be sure it is something
no one can guess. Do not use obvious codes like your spouse's name; the name of your group, account, or
file; your pet's name; your password; your telephone number; a period; a, b, ¢, and so on. Choose
something with which only you can identify. Change your codewords frequently to prevent the possibility
of some unauthorized or unanticipated person gaining access to your group. Examples of good typed
codewords are misspelled words of at least seven characters or words with numbers inserted in them,

words spelled backward with one or two numbers inserted, or a combination of several short words.

Since curriculum design information and student data are contained in the group file, instructors should be
very selective about which users are allowed to access the group file. When you allow a user access to the
file, you give that user the right to see, change, and/or destroy your records. Before assigning GROUP or
ACCOUNT codes, carefully determine whether or not you want all users within the group or account to
have access to group information.

One of the security options in the group file is the System Access option. This option lets you choose
whether or not to give systems personnel (users responsible for maintaining the PLATO system) access to
the group file. Systems personnel occasionally need access to files to check for errors if problems occur
on the system. If you choose to allow systems personnel access to your file, authorized users can access
the file in inspect (read) mode, without typing a security codeword.

3-12 97405900 C

All security codeword settings are made from the Security Codewords display (figure 3-6). The following
steps describe how to reach the Security Codewords display and how to set and change codewords and the
Systems Personnel Access option for the group file.

Press the associated number to change an entry.

SECURITY CODES:
1. To change group ----  *ektexeane
2. To inspect group ---  #kxKKKK HEE

3. To write records --- No match permitted
4. To read records ---- No match permitted

Access to group by system personnel:
5. System Access

HELP available

Figure 3-6. Security Codewords Display

1. From the PLATO Facilities display, choose the Group Operations option by typing the letter in
front of the option. The system displays the Group Operations display (figure 3-7).

2. From the Group Operations display, press DATA. The system displays the Group DATA
(Directory) display (figure 3-5).

3. From the Group DATA (Directory) display, choose the Security Codewords option by typing the
letter in front of the option. The system displays the Security Codewords display.

97405900 C 3-13

(authors)

Group “medicine” 2 people
33% full

Choose an option (or press HELP):

1 SEE or change someone's record

2 ROSTER operations
(ist, add, delete, messages, who's running)

STATISTICS on records
CURRICULUM design

SPECIAL options

Press BACK to leave.
Press LAB for space usage information.
Press DATA for group description.

Figure 3-7. Group Operations Display

From the Security Codewords display, you can do any of the following, depending upon the kind of access
(inspect, edit) you want to set for the group file.

@ The To Change Group option allows you to determine which users or groups of users can have
access to change (edit) the group file. To choose this option, type the number in front of the
option. Do one of the following depending upon the type of user access you want to allow.

- To allow editing access to yourself only, type a codeword only you know and press NEXT. A
random number of X's appear to the right of the arrow as you type. The system asks you to
retype it to verify it and help you remember it. Press NEXT.

- To allow editing access to all authors and instructors in your group who have the group

access option set to yes in their user records, press LAB. The system responds by displaying
a GROUP option and an ACCOUNT option. Type the number in front of the GROUP option.

3-14 97405900 C

- To allow editing access to all authors and instructors in groups in your account who have the
group access option set to yes in their user records, press LAB. The system responds by
displaying a GROUP option and an ACCOUNT option. Type the number in front of the
ACCOUNT option.

~ To prevent any user from seeing or changing the file, press LAB. Type the number in front
of the unmatchable code option.

e@ The To Inspect Group option allows you to determine which users or groups of users have access
to read information in the group file. To choose this option, type the number in front of the
option. Do one of the following depending upon the type of user access you want to allow.

- To allow inspect access to yourself only, type a codeword only you know and press NEXT. A
random number of X's appear to the right of the arrow as you type. The system asks you to
retype the codeword to verify it and help you remember it. Press NEXT.

- To allow inspect access to all authors and instructors in your group, press LAB. The system
responds by displaying a GROUP option and an ACCOUNT option. Type the number in front
of the GROUP option.

- To allow inspect access to all authors and instructors in groups in your account, press LAB.
The system responds by displaying a GROUP option and an ACCOUNT option. Type the
number in front of the ACCOUNT option.

- To prevent any user from seeing or changing information in the file, press LAB. Type the
number in front of the unmatchable code option.

@ To change either the inspect or change codewords, do the following.

- Type the number in front of the option you want changed. The system clears the previous
codeword.

- Type anew codeword and press NEXT, or press LAB to allow either group or account access.
~ Retype the codeword to verify it and help you remember it. Press NEXT.

@ The System Access to Group option allows you to choose whether or not to give systems
personnel access to your group file. Type the number in front of the option to either select or

change the option. For example, to disallow systems personnel access, type the number in front
of the option. To change the access, type the number again.

GROUP OPERATIONS

The most frequently used option on the PLATO Facilities display is the Group Operations option. The
Group Operations option allows you to add students to or delete students from your group file, see a list of
all group members, change student records, and do other administrative procedures,

The group operations described in this section are those instructors use most often. Both a description of

the option and how-to information is given. Refer to Additional Instructor Options later in this section
for information on Group Operations options used less frequently.

97405900 C 3-15

REGISTERING STUDENTS

As an instructor, you can register (roster) students or other types of users in your group. The following
steps describe how to add a user to your group.

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

2. Select the Roster Operations option.
3. Select the Add Someone to the Roster option.

4. Follow the display instructions.

DELETING STUDENTS
This option allows you to remove students from your group, either one at a time or all students at once.
When you delete a user from your group, you permanently delete that user's file or user record. (To

temporarily turn off a record, without destroying it, refer to Inspecting/Changing Student Records later in
this section.) The following steps describe how to delete users from your group.

Deleting Individual Students

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

2. Select the Roster Operations option.
3. Select the Delete Someone from the Roster option.

4. Follow the display instructions.

Deleting All Students

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

2. Select the Special Options option.
3. Select the Delete All Records option.

4. Follow the display instructions.

3-16 97405900 C

LISTING ALL GROUP MEMBERS

This option allows you to see a listing of all the users who are registered in your group. From this listing,
you can choose to see an individual user's record to review or delete it. The following steps describe how
to see a listing of all group members.

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

2. Select the Roster Operations option.
3. Select the See the Roster of people option.

4. Type the name or number of the user's record you want to see. Press NEXT to see the record or
press SHIFT-HELP to delete it.

INSPECTING/CHANGING STUDENT RECORDS
This option allows you to see or change a user's record. Use this option to check a student's progress in a
lesson or curriculum, change a user's password, change an author's system privileges, or to make other
changes as needed. The following steps describe how to access a user record.

1. Choose the Group Operations option from the PLATO Facilities display.

2. Select the See or Change Someone's Record option.

3. Type the PLATO name of the user whose record you want to see and press NEXT.

4. Select the option that most appropriately describes the desired record change.

5. Follow the display instructions.

LEAVING A MESSAGE

As an instructor, you can leave a message to the members of your group. The group members see the
message when they sign on to the PLATO system (immediately after typing their passwords). Messages
can be displayed to one user, all group members, or specific user types (student, instructor, author). The
following steps describe how to leave a message for your group.

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

2. Select the Roster Operations option.
3. Select the Leave a Message for Someone option.

4. Follow the display instructions.

97405900 C 3-17

USING YOUR ACCOUNT

Your PLATO account contains a definition of the services your organization purchased with the PLATO
system. It contains information on the number of people who can use the system simultaneously, and also
keeps a running record of the amount of file space purchased and used to date. [ A file is a finite set of
data. The space within the computer system which stores this information is called file space (file space
is a defined, finite subset of a file). Group records, notes files, lessons, and so on are examples of files. |
Think of a file as a book, a file space as a chapter, and a record as a section within a chapter.

Each account is allotted a specific number of file spaces when the account is established. This file space
is allocated to users in the account. The main person responsible for the account is the account owner.
The account owner is designated and identified by his/her PLATO name within the account when the
account is created. The account owner manages, creates, and lengthens files; controls which users can
see account information; and communicates with Control Data when additional file space is needed.

Account owners can delegate their responsibilities to other users in the account. Users who have been
delegated account authority are often called account directors. (There are several different levels of
account authority; not all users with access to account information are account directors. Refer to Using
an Account Aecess List in section 5 for more information on account access levels.) As an instructor,
your account owner might give you the authority to access the account. This might enable you to allocate
file space to yourself and other users in the account, as well as access other account information. If you
are an account director, or are assigned account director responsibilities by your account owner, refer to
Using Account Options (section 5) to learn more about account director funetions and responsibilities. If
you are not an account director or do not have account director authority, contact your account owner or
director if you need to create new files and increase the length of a file you use.

As an instructor, you can see information about your account if your account owner or director gives you
access. The following steps describe how to access your account.

1. Choose the Account Transactions option from the PLATO Facilities display by typing the letter
in front of the option.

2. Type the name of your account and press NEXT. Type the security code (if required) and press
NEXT. The system displays the Account Main Options display (figure 3-8).

3. Press DATA from the Account Main Options display. The system displays the General Account
Information display (figure 3-9).

3-18 97405900 C

wSaceeeocleeses testa on minne

Disk parts remaining -- 254

NEXT for PLATO and PLM file management options

Display file data

Lesson usage data

Current users in this account

Report generator options

Group records report generator

Archive options

Print access control options

Inter-accournt options

Network options

DATA for General Account Information

HELP available

Figure 3-8. Account Main Options Display

97405900 C 3-19

wore e nee testa on minne

woe renee lavalley 7 s

Disk parts remaining - 262

Disk parts allotted -- 482% HELP is available.
Files in account ----- 47

Subscriptions --------

1
2. Data change code ------------

3. File change code ------------ No code--owner only
4

5

Access by system personnel -- ALLOWED
Recount access list --------- This account file

a. Lesson Notes File -----------

b. Default file change code ---- GROUP 5

c. Default file inspect code --- ACCOUNT system
d. Network log datafile -------- rrl2

Network alternate log file -- rrii

Press the number or letter to change an item.
Press DATA for lesson access classes for this account.
Press SHIFT-NEXT to inspect or edit the account access list.

Account last changed on 11/83/88 at 9:88:33 am
by renee lavalley 7 5 at station 8-4

destroy file zrrlact

Figure 3-9. General Account Information Display

The General Account Information display gives general information about the account. It provides the
name of the account owner and stores the security codes selected by the account owner which control
user access to the account. It also contains information on the date the account was last changed, the
name of the person who changed it, and the change that was made.

To learn more about accounts, refer to Using Account Options (section 5).

REVIEWING AND STUDYING LESSONS

As an instructor, one of your responsibilities is to assign lessons for students in your group to study. You
should review lessons to examine their content and determine their applicability to your curriculum. You
can also review lessons to familiarize yourself with the learning materials, to write test questions from,
or to determine whether or not to assign additional learning resources to accompany the lessons.

3-20 97405900 C

You ean review any published lesson on the PLATO system for which your account contracts. Published
courseware is copyrighted and is available on all PLATO systems. Before the courseware is published, it
is tested and reviewed to ensure that the lessons operate properly, that all function keys work as
described, and that there are no coding errors which could cause the lesson to work incorrectly. Published
courseware is well maintained and reliable. It is never unexpectedly revised or deleted from the system.

All published lessons are included in a PLATO courseware library. Each PLATO account contracts for
access to specific courseware libraries. You can see any lesson in the library(s) for which your account
contracts. A list of the libraries available to your account is available from the General Account
Information display in your account. The Catalog of Available Courseware contains a listing of all
published PLATO lessons. Refer to Using the Catalog of Available Courseware in section 4 for a
description of the differences between published and proprietary courseware and for information on using
the catalog.

To review lessons or curricula and see them as a student, either use the Catalog of Available Courseware
to see an individual lesson (refer to Using the Catalog of Available Courseware in section 4) or create a
student sign-on for yourself to see a PLM or "mrouter" curriculum.

The following steps describe how to review published PLATO lessons.

1. Select the Choose a Lesson to Study option from the PLATO Facilities display by typing the
letter in front of that option.

2. Type the name of the lesson you want to see and press NEXT. The system displays your lesson.

PLATO COURSES AND CURRICULA

PLATO lessons ean be presented to students in various ways. As an instructor, you determine how your
lessons are presented and also how much flexibility to give your students while taking lessons. Most
students are assigned a curriculum to study. A curriculum is a study plan which concentrates on a specific
topic or subject, or set of topics.

Curricula can be presented to students in different ways, depending upon which method of instructional
delivery you select. Some curricula are composed of several modules. A module is a group of lessons
which relate to the same basic subject. Each lesson in the module presents instructional materials which
concentrate on a different area or aspect of the module-topic. The combined lessons and modules
compose the curriculum (figure 3-10).

97405900 C 3-21

CURRICULUM

MODULE MODULE

MODULE
LESSON LESSON LESSON LESSON

LESSON LESSON
LESSON LESSON

LESSON

Figure 3-10. Curriculum Structure - Example i

Other curricula are composed of several courses. A course is a complete learning package which
concentrates on a specific topic or subject area of the curriculum. It contains modules which can present
objectives, provide instructional lessons, administer tests, and suggest study materials related to the
subject of the course (figure 3-11).

Most curricula are designed to meet the individual needs of students. Curricula can be designed to allow
differing degrees of flexibility to students studying the curricula. Some curricula are designed to allow
students to choose the order of the lessons they want to study, while others require them to study lessons
according to a specified sequence or establish hierarchies of prerequisite and more advanced lessons.
Some curricula give students the option to take a test before studying a lesson. Usually, if they pass the
test, they are not required to study the lesson. Taking a test before studying the lesson materials helps
them know which parts of the lesson they should concentrate on more than others, and lets them preview
the test questions. Many curricula include a statement of the objectives for the modules and lessons.
These objectives help the students get a better understanding of the purpose of the lesson and the
information they are expected to know once they complete the course of study.

3-22 97405900 C

CURRICULUM

COURSE

COURSE

MODULE MODULE MODULE

@ Objectives @ Objectives e@ Objectives @ Objectives

e@ Lessons e Lessons e@ Lessons e Lessons

e@ Tests e Tests e Tests e Tests

e@ Study materials @ Study materials e Study materials e@ Study materials

Figure 3-11. Curriculum Structure - Example 2

DESIGNING A CURRICULUM

Planning a course of study for students is called designing a curriculum. Several factors are involved in
designing a curriculum. First, decide what the curriculum should accomplish. This determines the
instructor's curriculum needs and also helps the instructor select the best instructional management tool
for these needs. Some examples of things to consider when determining the curriculum objectives are:
the kind of student population studying the curriculum, the difficulty of the subject, the time frame to
work within, and so on.

There are several ways to design a curriculum. One way is to choose a published curriculum from the
Catalog of Available Courseware (section 4), and assign it to your students. A published curriculum
contains a preselected series of lessons, usually arranged in modules, which relate to the same subject
area. As an instructor, you simply assign the published curriculum for your students to study. Another
way to design a curriculum is to choose either individual lessons or sets of individual lessons from the
Catalog of Available Courseware or other sources to include in a curriculum of your own design. You
select the lessons you want included in your curriculum and either arrange them into modules, or index
them in a list for students to study.

Curricula are assigned to specific groups of students. The PLATO system, however, allows you, the
instructor, to individualize the curriculum for your students. For example, even though all students in the
group are registered for the same curriculum, not all students must study all the lessons in that
curriculum. You can specify which lessons should be studied by which students; vary the order or
sequence of lesson presentation; allow the students to choose which lessons they want to study first,
second, and so on; and have the system test students on lessons studied and record their progress.

97405900 C 3-23

INSTRUCTIONAL MANAGEMENT TOOLS

Designing and individualizing curricula is called instructional management. The PLATO system provides
four instructional management tools which assist instructors in individualizing a curriculum for all
students in the group. Each of these four instructional management tools has its own set of features and
capabilities associated with it. Many of the features overlap and are available with more than one
instructional management tool. This provides a wide range of options for the instructor to choose from
when designing curricula. The four instructional management tools are: index lessons, the PLATO system
router ("mrouter"), PLATO Learning Management (PLM), and routers. The following briefly describes
these tools.

Index lessons

An index lesson is a lesson written by an author using the PLATO Author Language. The index lesson
presents a set of lesson choices on an index display to students. They are usually used for curricula
which are very straightforward and when little to no student data collection is desired. Index lessons
do not collect student data unless they are coded to do so.

The system supported router ("mrouter")

The PLATO system router, "mrouter", is a lesson delivery system which contains the mechanics for
presenting a list of lessons to students. Since "mrouter" is a delivery system, it does not contain
specific information about which lessons to present, the order of lesson presentation, or what criteria
are required to master the curriculum. This information is supplied by the instructor through an
instructor file and is then delivered to students by "mrouter". The PLATO system router also collects
student data for instructors to evaluate student progress and performance.

PLATO Learning Management (PLM)

PLATO Learning Management (PLM) is similar to "mrouter" in that it directs the student through a
curriculum. PLM, however, has many additional features. As its name implies, PLM has management
capabilities. It allows the organization of both on-line and off-line instructional materials into
modules and courses which comprise a PLM curriculum. Some of PLM's management capabilities
allow users to:

e Provide an introduction to the curriculum for students.
e Enter learning objectives associated with both on-line and off-line materials for students.

e Enter test questions and instructions for the presentation of the test questions. Test
preparation requires no programming knowledge on the part of the instructor.

Router lessons

A router lesson is a lesson written by an author using the PLATO Author Language. A router lesson is
written when an instructor has some special curriculum design requirements which are not included in
‘mrouter" or PLM. A router lesson contains the code which executes the special features which are
needed to fulfill the requirements of the desired curriculum design.

Table 3-1 summarizes the features and capabilities of the instructional management tools. Refer to the
table to select the instructional management tool which best meets your curriculum design needs. Then
refer to the appropriate following section to learn how to design your curriculum using that particular
instructional management tool.

3-24 97405900 C

TABLE 3-1. CURRICULUM DESIGN OPTIONS

User
Index Written
Curriculum Design Options Lesson "mrouter" Router
Curriculum introduction

Total touch input; no keyset input
beyond sign-on

Touch input for test questions

Can display objectives to students
Shows students their scores

Requires testing for student placement
Allows specification of prerequisites
Sequence of lessons; no index available
Allows review of completed material

Module indexes

ba > ns ° a © ns ~ a ~ a ° a” ©

Allows reference to material not on
PLATO system

Suggests study assignment options

Allows an introduction to an assignment
on or off PLATO system

Maximum number of lessons per curriculum
Maximum number of modules per curriculum
Maximum number of modules per course
Maximum number of courses per curriculum
Student statistics by module

Student statistics by curriculum

t More extensive student data than "mrouter".

S Standard option.
P Programmable option.
- Not available.

97405900 C 3-25

How to Use Index Lessons

An index lesson is a PLATO Author Language coded lesson in a TUTOR file which presents a list of lessons
for students to study.

Index lessons do not store any information beyond what is ordinarily kept in student variables. Index
lessons are usually used when instructors do not want to keep student data or lesson data. They are used
to provide an index of lessons not requiring the module structure of "mrouter" or PLM.

Before you can correctly write and code an index lesson, you need to have an author sign-on and a working
knowledge of the PLATO Author Language. If you do not know the PLATO Author Language, contact
your account owner for help or for information on how to learn the PLATO Author Language.

When you are ready to begin coding your index lesson, you should contact your account owner to create a
file in which to store your code, and to give you an author sign-on to access the file.

How to Use the System Supported Router (“mrouter”)

The system supported lesson delivery system "mrouter" contains the mechanics for presenting lists of
lessons (modules) for students to study. In addition to presenting lessons, "mrouter" collects data on
individual student progress and performance, as well as group data. You can see information on the
number of days, hours, and sessions each student used the system, or the average time for all students in
the group. Information on lessons students completed and their test scores is also available, as well as the
average score for all students in a lesson.

Since "mrouter" is a lesson delivery system, it does not contain specific information about any one
curriculum design. It is a mechanism which presents many curricula. Curriculum specific information is
stored in an instructor file. An instructor file is a file which contains specifie information about each
curriculum design which "mrouter" delivers. This information includes the file names and titles of all the
lessons in the curriculum, the order of lesson presentation, the number and structure of the modules, and
the criteria for mastering the modules in the curriculum. Essentially, the instructor file tells the system
what lessons to present to the student, and when and how to present them.

Each instructor file contains a curriculum catalog which stores the file names and titles of the individual
lessons which comprise the curriculum. The lessons contained in this curriculum catalog are selected
from either the Catalog of Available Courseware, or from other sources (such as unpublished lessons
written by other authors). From the lessons listed in the curriculum catalog, modules are created by
selecting lessons and listing them in specific modules.

Basically, there are two ways to design a curriculum using "mrouter". One way is to create your own
curriculum by selecting individual lessons, listing them in the curriculum catalog, and inserting them into
modules. Another way is to select a published curricululm from the Catalog of Available Courseware.
Generally, all published curricula are organized in instructor files. A published instructor file contains a
completed curriculum catalog, definitions of the number and composition of the curriculum's modules, and
mastery criteria for each module. When an instructor selects a published curriculum, the instructor uses
that curriculum's published instructor file instead of creating one of his/her own. The only time an
instructor needs to create an instructor file and insert information in it is if the instructor is creating a
curriculum by choosing individual lessons from the Catalog of Available Courseware, arranging the lessons
into modules, and establishing criteria for completing the modules. Your account owner can create an
instructor file for you if you do not have account director capabilities.

The following sections describe how to use "mrouter" with a published curriculum and when designing your
own curriculum.

3-26 97405900 C

Using "mrouter" with a Published Curriculum

After you have selected a published curriculum from the Catalog of Available Courseware, do the
following steps to attach the curriculum to your group file.

1. Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option. The system displays the Group Operations display (figure 3-7).

2. Press DATA. The system displays the Group DATA (Directory) display (figure 3-5).

3. Choose the Associated Files option by typing the letter in front of the option. The system
displays the Associated Files Options display (figure 3-12).

medicine

Press the associated number to change an entry.

Student notes:
1.

Data collection:
2.

5. Access privileges --

Routers:
6. Student router mrouter
7. Instructor router -- imode

Instructor file:
8. Instructor file ---- med!

Figure 3-12. Associated Files Options Display

4. Locate the router section of the display. Choose the student router option by typing the number
in front of the option.

5. Type mrouter and press NEXT. The system asks for the name of your instructor file.

97405900 C 3-27

6.

7.

e the name of the instructor file (from the Catalog of Available Courseware) and press

NEXT. The system asks for the use codeword. The use codeword for all published instructor
files is the same as the name of the instructor file. Type the instructor file name again and
press NEXT.

Access the instructor file and complete the author information section in the file.

Using "mrouter" with Your Own Curriculum

There are five steps involved in using "mrouter" with a curriculum of your own design.

3-28

1.

Before you begin these steps, it is a good idea to
read the reference materials in AIDS on "mrouter"
to get a thorough understanding of its intended
uses.

Attach "mrouter" and an instructor file to the group file.

a.

b.

Cc.

d.

Follow steps 1 through 5 in Using "mrouter" with a Published Curriculum (directly preceding
this section).

Type the name of your instructor file and press NEXT. (Remember to contact your account
owner to create an instructor file for you if you are not an account director.)

Access the instruetor file from the Curriculum Design option on the Group Operations
display (figure 3-7) of the file.

If no eodewords were assigned to your instructor file when it was created, you will be
brought to the Instructor File Information display (figure 3-13), also reached by pressing
DATA from the Curriculum Options display (figure 3-14). Assign codewords to the
instructor file and enter the descriptive information to prevent accidental deletion of your
file during account cleanup.

Insert lesson information in the curriculum catalog.

a.

After you attach an instructor file to your group file, the Curriculum Design option appears
on the Group Operations display. This option appears only after an instructor file is
attached to your group file. Access the Curriculum Options display (figure 3-14) by typing
the letter in front of the Curriculum Design option on the Group Operations display. From
the Curriculum Options display, choose the See Catalog of Lessons option to insert your
lesson list in the curriculum catalog. This list consists of all lesson file names and lesson
titles to be used in all modules of your "mrouter" curriculum. Follow the instructions at the
bottom of the screen to insert lesson information. To copy an instructor file or curriculum
catalog from another curriculum, or to delete a lesson or change lesson names, use Special
Curriculum Options on the Curriculum Options display.

Press HELP from the Curriculum Options display and read the section titled The Curriculum
Catalog for more detailed information on the curriculum catalog and how to use it.

97405900 C

97405900 C

Instructor File Information:
medi

“change” codeword: SETKRLERED
“use only” codeword: sasssssees

- Name of Owner: jane doe

Group(s) for which this file is used:
medicine

Type the number of the entry you want to charge.
(This information must be filled in before the
file can be used.)

Figure 3-13. Instructor File Information Display

CURRICULUM OPTIONS

see/design curriculum MODULES
seerconstruct SEQUENCES of lessons
see CATALOG of lessons

SPECIAL curriculum options

Press DATA to see instructor file information.

Press HELP for a discussion of modules, sequences,
and the curriculum catalog.

Figure 3-14. Curriculum Options Display

3-29

3-30

3.

4.

Design modules.

There are three types of modules you can use in your curriculum.

Index module - Contains a list of lessons and/or sequences from which a student makes a
choice. The lessons may be selected in any order by the student and any lesson may be
reviewed.

Sequence module - Contains a list of lessons for students to study in a specified order. The
instructor specifies which lessons the students will study and the order in which the lessons
are studied. Each lesson is presented immediately after the previous lesson is completed.
The students must study the lessons according to the specified sequence. Students do not
see an index from which they can choose lessons.

Sequence with review module - Contains an index of lessons which expands as the student
completes lessons in that sequence. Students see an index listing only completed lessons.
They can review previously studied lessons or continue with the next lesson in the sequence.
This module is similar to the sequence module except the student can review any completed
lessons.

After you have determined the kind(s) of module(s) you want to use in your curriculum, you can
begin designing modules by inserting individual lessons in the modules or by assigning sequences
of lessons to them. The following steps describe how to design modules.

a.

b.

Choose the See/Design Curriculum Modules option from the Curriculum Options display.
The system asks you to select the type of module you want to create. Type the number in
front of the desired module.

Type a name for the module and press NEXT. The system displays a Module Description
display related to the type of module you are creating. Press HELP for information on how
to insert lessons in the module or assign a sequence of lessons.

Create sequences of lessons.

If you plan to use sequence modules or sequence with review modules in your curriculum, you
need to establish sequences of lessons to be used with those module types. The lessons in the
sequences are selected from the lessons in the curriculum catalog. You select the lessons and
determine the order in which you want the lessons presented. You can create up to 10 different
sequences per curriculum but only one sequence can be assigned per module.

The following steps describe how to create sequences of lessons.

If you prefer, you can create sequences of lessons

when you create sequence modules or sequence
with review modules.

97405900 C

a. Choose the See/Construct Sequences of Lessons option from the Curriculum Options
display. The system asks you to number the sequence you want to create.

b. Assign a number to the sequence. Press NEXT. The system displays a list of numbers.

e. Type a to insert lessons in the sequence. The system displays the list of lessons in the
curriculum catalog. From this list, select the lessons you want included in the sequence you
are creating by typing the number in front of the desired lesson(s) and pressing NEXT. The
order in which you select the lessons will be the order in which the lessons will appear in the

sequence.

In addition to assigning sequences of lessons to individual modules in the curriculum, you can also
assign sequences of lessons to individual students in your group. The sequences can be
individualized to meet the specifie needs of selected students. When a sequence is assigned to a
specific student and is tailored to meet the individual needs of that student, it does not affect
any other sequences defined in the curriculum's modules. Student sequences are separate from
general module definitions.

To assign a specific sequence to an individual student, access the student's user record in the
group file, choose the Curriculum Status option, and follow the instructions.

5. Assign completion criteria to modules.

After you have created your modules and sequences, you should establish the completion criteria
for the modules. There are three types of completion criteria you can assign to your modules.

e Score criterion - Allows you to specify a minimum score for a lesson or a minimum score for
all lessons in the module.

e Item criterion - Allows you to specify a specific lesson to be eompleted or a minimum
number of lessons to be completed for that module.

e Time criterion - Allows you to specify either a specific date or time frame by which the
students must complete the module.

To establish completion criteria for the module, press DATA from the Module Descriptions
display and then press EDIT. The system asks you to select the type of criterion you want to
use. Type the number in front of the kind of criterion you want to use and then follow the
system instructions. (For more detailed information, press HELP from the Curriculum Options
display and read the section titled Completion Criteria for Modules.)

After you have created a module and pressed BACK, the system displays the Module Listing
display (figure 3-15). This display lists all the modules in your curriculum and provides
summarized information about each (type of module, number of lessons in the module, and
completion criteria for the module). From this display, you can create new modules, change the
sequence of modules, and change the titles of the modules. Press HELP from this display for
information on how to do these operations.

97405900 C 3-31

Design Curriculum Modules medicine"

4 in use 4 maximum “medi ®

NAME, CTYPE) * ITEMS CRITERIA

cobol Move ahead after 3 days
(index) Next: 2 Back: 1

two Completion of sequence
(sequence) Next: 3 Back: 1

cybol Completion of sequence
(seq. with review) Next: 1 Back: 2

four Completion of sequence
(seq. with review) Next: 4 Back: 3

See/Revise Module Number: >
HELP available
DATA to start a new module
LAB to change module progression
shift-LAB to change module titles

Figure 3-15. Module Listing Display

Special Instructor File Options

There are several additional options available from the Special Options entry on the Curriculum Options
display. These options allow you to copy lessons from another curriculum catalog, copy another instructor
file, revise the curriculum catalog, increase the number of modules in the curriculum, delete all modules
and sequences, delete individual lessons from the curriculum catalog, and destroy the curriculum catalog.

3-32 97405900 C

How to Use PLATO Learning Management (PLM)

Refer to the following manuals for more information about PLM and to learn how to use this feature.
e PLATO CMI t System Overview
e PLATO CMI Instructor's Guide
e@ PLATO CMI Author's Guide

Users should also refer to the PLM On-Line Reference Manual in AIDS. To use PLM AIDS, choose the
PLM AIDS option from the PLATO Facilities display by typing the letter in front of the option.

How to Use Your Own Router

To create your own router, you need either an author sign-on or the assistance of someone with an author
sign-on. For information on how to design your own router, refer to AIDS and the PLATO Author
Language Reference Manual.

USING DATAFILES

As an instructor, you can collect information on lessons in your curriculum with a datafile. A datafile is a
special file which collects and stores information on lessons which have been coded by an author to collect
data. As an instructor, you ean specify the kinds of data you want collected about the lesson; the lesson
author, however, must have coded the lesson to allow data collection in order for data to actually be
collected. Each datafile has a set of options which allow you to choose the kind of information you want
collected. You should use a datafile to collect student data relating to a specific lesson or set of lessons.
Most instructors use datafiles for summative data collection (for example, to see a student's performance
for the semester, or see how the class as a whole is performing). Most authors use a datafile for
formulative evaluation purposes during the development of a lesson. Datafiles are rarely used with
published lessons or curricula.

Datafiles are also used to collect area summaries for students in a group. Area summaries inelude such
things as unanticipated responses from students to questions in a lesson, lesson HELP requests made but
not received, and the amount of time a student spent in a particular part of a lesson. Output data is
information which the author of the lesson has coded the lesson to collect. Output data can include
information which is not covered in an area summary. The area summaries can be analyzed for several
areas of lessons.
The following are some examples of the kinds of information a datafile can coliect.

e Student answers to questions in a lesson (correct answers and unanticipated incorrect answers).

e Ratio of right to wrong answers to questions.

e@ Student requests for HELP, both answered and unanswered.

e Number of TERM requests made and completed (answered).

e Searches of other datafiles for specific types of data on students, lessons, modules, and so on.

These search options are available in any combination.
tT PLATO Learning Management (PLM) was originally entitled Computer Managed Instruction (CMI).

Some documentation still exists which refers to PLM in this manner.

97405900 C 3-33

There are three steps involved in preparing a datafile to collect student data. These steps must be
performed before you can begin collecting information in the datafile.

1.

2.

3.

Create a datafile and attach it to the group file of the students for whom you want data
collected.

Set the security codewords in the datafile to prevent unauthorized users from seeing or changing
the contents of the file.

Set the data collection options within the group file to determine the kinds of student data to be
collected in the datafile. :

To create and attach a datafile to your group file, do the following steps.

1.

2.

3.

6.

Create a datafile through your account (if you have account director capabilities) or ask your
account director to create a datafile for you.

Choose the Group Operations option on the PLATO Facilities display by typing the letter in front
of the option. The system displays the Group Operations display.

Press DATA for the Group DATA (Directory) display.
Choose the Associated Files option by typing the letter in front of the option.
Type the number next to the Data Collection option. The system displays an arrow.

Type the name of your datafile. Press NEXT.

To set the datafile security codewords, do the following steps.

1.

2.

3.

Choose the Datafiles option on the PLATO Facilities display by typing the letter in front of the
option.

Type the name of the datafile. Press NEXT.

Press DATA. The system displays four options. Select an option by typing the number in front
of the option. The options are:

e Change code Allows you to restrict which users can see and change information
in the datafile. You can assign either a typed, group, account, or
unmatehable security code for the datafile. (Refer to Group
Security earlier in this section for more information on file
security codes and how to set them.)

e Inspect code Allows you to restrict which users can see information in the
datafile. You can assign either a typed, GROUP, ACCOUNT, or
unmatechable security code for the datafile. (Refer to Group
Security earlier in this section for more information on file
security codes and how to set them.)

e System access Allows you to choose whether or not to allow systems personnel
access to the datafile. Type the number in front of the option to
change the setting.

e Print information Allows you to enter your name and mailing address to ensure that
if you request a print of the datafile, it will be sent to you. (Refer
to Requesting Prints in section 4 for information on requesting
prints.)

97405900 C

To select the data collection options, do the following steps.

1

3.

From the Group Operations display, choose the Special Options option by typing the number in
front of the option.

Choose the Change Group-wide Data Collection option by typing the letter in front of the option.

Do one of the following, depending upon the type of data you want to collect.

a

b.

Select the Change Data Collection option if you want to set group-wide data collection
options for all students in the group (as a whole). The system displays a list of options. The
options you select must either match the lesson's ~dataon- tags or the lesson must have a
blank -dataon- command. Choose the options which relate to the type of data you want
collected. After you make your selections, press NEXT. These options will then be set for
all new students in the group (that is, any new students added to the group). To set these
options for students already registered in the group, press SHIFT-HELP.

Select the Specify Data Collection option if you want a specific lesson to collect extra data
beyond what is already specified for the group (in the Specify Data Collection option). (The
lesson's -dataon- tags must match the options you select in order for the data to be
collected.)

Press HELP for more information on data collection or refer to AIDS for a more detailed
explanation of datafiles.

ADDITIONAL INSTRUCTOR OPTIONS

The following options are the remaining Group Operations options available to instructors from the
PLATO Facilities display. These options are usually used less frequently than those described earlier in

this section.

MONITORING GROUP MEMBERS

From this option, you can see a list of all the group members who are currently signed on to PLATO
terminals. You can also see additional information such as the number of hours an individual user has
been signed on, the name of the lesson(s) being studied, the user category in which the person belongs, and
so on. This option also allows you to monitor (see) another user's screen. The following steps describe
how to access this option.

97405900 C

Choose the Group Operations option from the PLATO Facilities display by typing the letter in
front of the option.

Select the Roster Operations option.
Select the See Who is Now Running option.

Follow the display instructions.

3-35

SPECIFYING GROUP DATA COLLECTION
As an instructor, you can see selected statistical information on students in your group. Formative
evaluation information can be collected for individual students or for all students in your group in a
datafile. The following steps describe how to collect group data.

1. Choose the Group Operations option from the PLATO Facilities display.

2. Select the Special Options option.

3. Select the Change Group-wide Data Collection option.

4. Follow the display instructions.

Refer to Using Datafiles earlier in this section for information on how to create and use datafiles.

TEMPLATING RECORDS

A template is a student record that is used as a model or pattern for other student records in your group.
A templated record standardizes one or more areas of a student's record and then duplicates that area on
other students' records. The student record which contains the original standardization is called the
template. Passwords, lesson names, unit names, student variables, and curriculum options can be
templated so that they are or will be the same on all records. Templates can be created for students who
are currently registered in the group, new students to be added to the group, or both.

After you decide which student record you want to use as the template, aecess the student record and set
the options in the record according to how you want the other records in the group to be set.

If lesson names and unit names are templated, all
restart information kept by the system is lost. If
variables are templated, information on students'
work may be lost.

The following steps describe how to set a template.
1. Choose the Special Options option from the PLATO Facilities display.
2. Choose the Set Up a Template Record option.

3. Select the first option if you want to set a template for all new students added to the group, or
select the second option if you want to set a template for all existing students in the group.

4. Type the name of the student whose record is to be used as the template. All student records

affected by the template will indicate such and identify the student user record used as the
template.

3-36 97405900 C

COPYING RECORDS FROM ANOTHER GROUP

This option allows you to copy the user record of a student registered in a group file other than your own
to your group file. Choosing this option transfers a copy of the student user record to your group file
without deleting the record from the original file. The following steps describe how to copy a user record.

1.
2.
3.
4.
5.

6.

Choose the Group Operations option from the PLATO Facilities display.

Choose the Special Options option.

Choose the Copy a Record from Another Group option.

Type the name of the group from which the user record is to be copied and press NEXT.
Type the codeword for the group. Press NEXT.

Type the name of the user whose record you want copied and press NEXT. A message indicates
completion of the copy.

CREATING INSTRUCTOR RECORDS

When you create an instructor record, the system displays a number of options which can be available to
the new user. It is your responsibility to designate which of these options you want the new instructor to
have access to. When you are deciding which options to allow, it is important to consider what kind of
position the new instructor holds and how much responsibility you want to delegate to that individual. For
example, if you are creating an instructor record for a new teaching assistant, you may not want to give
that person the option to change students’ scores, or if the new instructor is not involved with PLM, you
may choose not to include the PLM options. Use your judgement to determine how much responsibility
(that is, how many options) to initially delegate to the new user. An example of some of the available
instructor options is shown in figure 3-16.

97405900 C 3-37

Change options that mary smith can use.
These options will hold true at all times.

' Primary Instructor Options:

yes a choose a lesson from an instr. file CATALOG

yes b choose ANY lesson (by lesson name)
yes c see who is running at the SITE
see system-wide list of USERS
access PUBLIC notes and system announcements

Receive TERM-ask requests

Press HELP for information.

Figure 3-16. Available Instructor Options Display

The following steps describe how to create an instructor record and select instructor options.

1.
2.
3.
4.

5.

3-38

Choose the Group Operations option from the PLATO Facilities display.
Choose the Roster Operations option.

Choose the Add Someone to the Roster option.

Type the number in front of the kind of record you want to create.

Type the name of the person you want to add to the group and press NEXT.
Press DATA to see the new record.

Choose the Allowable Instructor Options option and do one of the following steps.

97405900 C

a. Press DATA to give the new instructor the same options you have.
b. Type the letter in front of the option you want to set.
ce. Press HELP for more information and a complete set of instructions.

8. Press SHIFT-BACK to return to the options index.

MANAGING PERSONAL NOTES

You ean see statistics regarding the use of personal notes by members of your group. This option allows
you to turn individual users' notes options on or off, see the total number of personal notes each user has
received, and set a maximum number of notes each user is allowed to receive. You can also delete notes
addressed to a user after that user has been deleted from the group. The following steps describe how to
manage the personal notes activities for your group.

1. Choose the Group Operations option from the PLATO Facilities display.
2. Choose the Special Options option.
3. Choose the Manage Personal Notes Activity option.

4. Follow the display instructions.

USING NOTES ft

Notes are messages stored in files on the PLATO system. They allow users to privately or publicly
communicate with each other and allow system messages to reach large numbers of users.

As an instructor, you can use a variety of notes features. You can write personal notes to or receive
personal notes from other PLATO system users; read or participate in general notes files and public notes
files; and read system announcements from PLATO systems personnel. Refer to Notes in section 4 for
information on the different kinds of notes and notes files and how to use them.

Instructors can also use notes to communicate with students in their group. Many instructors use the
student notes feature to do this. Student notes differ from other notes in that they are contained in a
special notes file created by instructors for students in a specific group. Student notes are unique because
they have a variety of functions. Refer to Using Student Notes in section 4 for more information on
student notes and to learn how to use them.

USING INTERACTIVE COMMUNICATIONS ‘

As an instructor, you can communicate with students and other users on the system by typing messages on
your screen. The person you are communicating with reads the message as you type it. You can use this
feature to receive answers to questions or solutions to problems from PLATO consultants, to help students
who have questions about their lessons while they are studying the material, or to converse with another
user about PLATO-related topics.

1 Refer to inside back cover for important regulatory notice concerning the use of communications
features.

97405900 C 3-39

The following interactive communications features are available to instructors.

TERM-talk Allows users to communicate with other authors or instructors by typing
messages on the screen. Refer to Using the Talk Feature in section 4 to learn
how to use this feature.

TERM-ask Allows authors and instructors to help users (students and other authors and
instructors) solve problems or to answer questions about their lesson materials
while the students are studying the lesson. Refer to Using TERM-ask in section
4 to learn how to answer users' TERM-ask requests, and to TERM-ask in section
2 to learn how to ask a question using TERM-ask.

TERM-consult Allows authors and instructors to receive on-line help from PLATO
consultants. Consultants and users communicate by typing messages on their
sereen. Refer to Consulting Help for Authors and Instructors in section 4 to
learn how to use TERM-consult.

REQUESTING PRINTS

As an instructor, you can request prints of any file (which can be printed) on the PLATO system for which
you know the security code. Refer to Requesting Prints in section 4 for more information on requesting
prints and how to use this feature. :

RECEIVING HELP

At some time, almost all users have questions or need help while using the PLATO system. There are
basically two kinds of help available to instructors: programmed help sequences and personal help. The
kind of help you request depends upon the type and extent of help you need, as well as the type of activity
you are engaged in at the time you request help. Many features and lessons contain programmed help
sequences which provide additional information about a feature or lesson. These programmed help
sequences are programmed by the author of the lesson or feature and are only available if the author has
programmed the help sequences into the lesson. Some examples of the kinds of information contained in
these help sequences are: descriptions on how to use specific features, detailed information about a topic
presented in a lesson, and information on how to proceed though the lesson. Programmed help is usually
accessed by pressing HELP while using a lesson or feature.

Occasionally, you might have questions which require more help than you are able to receive from
programmed help sequences. As an instructor, you can request personal help from other authors or
instruetors, or from PLATO system consultants to answer your questions. Personal help allows you to
communicate with another user and receive help while on-line.

The following describes some of the PLATO system features which instructors can use to receive help.
AIDS AIDS is an on-line reference manual for authors and instructors. It contains
definitions and explanations of most of the PLATO system features and all of

the PLATO Author Language commands. Refer to Using AIDS earlier in this
section for more information on AIDS and how to use this feature.

3-40 97405900 C

TERM-consult Authors and instructors can receive help while using the PLATO system by
using the TERM-consult feature. TERM-consult allows you to communicate
with a PLATO consultant about questions or problems you have while using the
PLATO system. Consultants are Control Data systems personnel who are
available to answer questions and solve problems for users on the system.
Refer to Consulting Help for Authors and Instructors in section 4 for more
information on TERM-consult and how to use this feature.

TERM-ask TERM-ask is a PLATO system feature which gives users (usually students and
multiples) the opportunity to ask authors and instructors questions about
PLATO materials while the materials are being presented to them. It also gives
authors and instruetors the opportunity to discuss with other authors and
instructors questions or problems they might have while using the PLATO
system. Refer to Using TERM-ask in section 4 to learn how authors and
instructors use TERM-ask to give help to other users and to TERM-ask in
section 2 to learn how to use the feature to receive help from other authors and
instructors.

TERM-talk The TERM-talk feature allows you to communicate with another user who is
currently signed on to the PLATO system by typing messages back and forth on
the bottom two lines of your sereens. Refer to Using the Talk Feature in
section 4 for more information on TERM-talk and how to use this feature.

GIVING HELP

Part of your responsibility as an instructor is to provide assistance to the students in your group when they
need help. The TERM-ask feature allows you to do this. TERM-ask allows students in your group to
request help from you (or other authors or instructors in your absence) about questions or problems they
have while using the PLATO system. You and your students can communicate by typing messages back
and forth on the screen and seeing the messages as they are being typed. It also lets you monitor (see) the
screen of the student requesting help.

Before your students can use the TERM-ask feature, there are some initial administrative tasks which you
need to complete. Refer to Using TERM-ask in section 4 for complete information on the TERM-ask
feature and how to use it.

97405900 C 3-41

SECTION 4
USING AUTHOR FEATURES

USING AUTHOR FEATURES

Introduction
Author Records
The Author Mode Display
AIDS
Catalog of Available Courseware
Notes
Personal Notes
User List
Prints
Understanding File Structure and Use
Getting Help
Using AIDS
Consulting Help for Authors and
Instructors
Using TER M-ask
TERM-talk
Using Communications Features
Using the Talk Feature
TER M-busy
TERM-reject
Commenting on Lessons and Features
Notes
Types of Notes Files
Notes File Access
Using General Notes
Using Personal Notes
Using Lesson Notes
Using Student Notes
Using Intersystem Notes
Writing General Notes
Using Notes File Features
Using Reference Tools
AIDS
Using the Catalog of Available
Courseware
Title Index
Author Index
Subject Index
File Name Index
Using the On-Line Author Listing
Using the PLATO User List
User Statistics and Flag Settings
Time-Saving Features
TERM-spell
Using Documentation Features
Using Documentor
Graphies Utility for Interactive
Documentation Ease (GUIDE)
Requesting Prints

97405900 C

4-1
4-1
4-3
4-5
4-5
4-5
4-5
4-5
4-6
4-6
4-10
4-10

4-15
4-16
4-19
4-20
4-20
4-21
4-21
4-22
4-22
4-22
4-23
4-23
4-29
4-31
4-32
4-33
4-33
4-36
4-39
4-39

4-39
4-41
4-43
4-44
4-44
4-44
4-47
4-49
4-49
4-50
4-50
4-50

4-52
4-53

Writing PLATO Lessons
Using the TUTOR Lesson File
Registering Author and Lesson
Information
Assigning Security Codewords
Editing TUTOR Files
Creating a TUTOR Block
Using the PLATO System Editor
Requesting Editing Help
AIDS Quick Reference (Q-Ref)
Condensing a Lesson
Creating Displays
Screen Locations
Inserting and Showing Displays
(ID/SD)
Creating Characters and Line
Drawings

AIDS Listing of Display Commands

Suggested Guidelines for Writing
PLATO Lessons

Understanding ECS/ESM Usage and
Charges

Referencing PLATO Author Language

Commands
Using Other Block Types
Common Block
Micro Block
Leslist Block
Vocabulary Block
Listing Block
Text Block
Copy-a-Block
Using Advanced Editing Directives
Using Block Listing Display Options
Documenting Lesson Changes
Writing Routers
Preparing Lessons for Publication
Using Data Collection Files
Dataset Files
Nameset Files
Code Files
Using Advanced Author Options
TERM-pnote
Setting an Alarm
Using an Automatic Sign-on
Using Author Reserved TERMS
TERM-cursor
TERM-grid
TERM-step

4-56
4-56

4-56
4-58
4-61
4-61
4-63
4-66
4-67
4-67
4-68
4-68

4-71

4-71
4-84

4-85
4-85

4-87
4-88
4-88
4-89
4-90
4-90
4-90
4-90
4-91
4-91
4-93
4-93
4-93
4-94
4-95
4-95
4-95
4-96
4-96
4-96
4-96
4-97
4-97
4-97
4-98
4-98

TERM-charset
Additional Security Options
Additional SHIFT-DATA Display Options

Micro PLATO Authoring

4-ii

Micro PLATO File Security

4-99
4-99
4-100
4-101
4-102

Micro PLATO Levels

Micro PLATO Lesson Execution
Flexible Disk Preparation

The Micro PLATO Router
Author References

4-102
4-102
4-103
4-104
4-104

97405900 C

USING AUTHOR FEATURES 4

This section presents an overview of the functions of an author and defines and describes the PLATO
system features available primarily to authors.

All users should read the Introduction (section 1) before reading this section. In addition, it is
recommended that new authors also read sections 2 (Using Student Features) and 3 (Using Instructor
Features), and practice using the system as a student and as an instructor before reading this section to
get a general understanding of the features available to those user types. (Some system features
described in this section are also available to instructors. The features which are available to both user
types contain instructions for both users.)

INTRODUCTION

The primary function of an author using the PLATO system is to write lessons for students to study.
Authors write lessons using a computer language called the PLATO Author Language. However, not all
users with author sign-ons know and use the PLATO Author Language to write lessons. Some users have
author sign-ons for the wide range of features and options available to authors.

An author sign-on is unique in that it allows authors to interact with the PLATO system on two levels.
Since most authors write lessons, the PLATO system is designed to allow authors not only to write and
edit lessons, but also to use the lessons exactly as students do. It allows authors to move between student
mode and author mode without signing on to and off from the system.

Because of the complexity and diversity of the activities authors are involved in, the PLATO system
provides authors with a large number of system features to support their work. Some of these features
provide reference information about system capabilities, some allow authors to receive on-line help from
other users or system resources, and some facilitate writing and editing lessons. Some of these resources
are available only to authors, while others are available to all users or specific user types.

A good way for new authors to learn about and practice using the PLATO system features available to
authors is to use the reference tools. Refer to Getting Help and Using Reference Tools later in this
section for information on how to use these features.

AUTHOR RECORDS

The PLATO system contains a user record for each author on the system. The user record contains a
complete listing of all the system options available to authors. From this listing, specific options are
selected for individual authors to access. These options control what each author is allowed to do on the
system. As an author, the options you can access on the system are determined from the options
designated as allowed in your user record. The person who registers you on the system is responsible for
assigning these options to you. An example of some of the available author options is shown in figure 4-1.

97405900 C 4-1

Press DATA if you want this person to have
the same options you have.

or enter the letter of the type of individual
options you want to set:
a. Primary Instructor Options
b. File Editing & Printing
c. Roster Options
. General Record Editing Options
. Student Record Editing Options
» Active User Options
Messages and Notes
» Data Collection Options
i. “mrouter” options

j- PLM options

Figure 4-1. Available Author Options Display

Some authors are responsible for registering other users on the system and assigning allowable author
options. The following steps describe how to create an author record and assign author options.

1. From the Author Mode display, type the name of the group in which you want to register the
author and press NEXT. The system displays the Group Operations display.

2. Choose the Roster Operations option by typing the letter in front of the option.
3. Choose the Add Someone to the Group option.

4. Type the number in front of the kind of record (author) you want to create.

5. Type the name of the person you want to add to the group and press NEXT.

6. Press DATA to see the new user record.

4-2 97405900 C

7. Choose Allowable Author Options and do one of the following steps.

a. Press DATA to give the new author the same options you have in your user record. You can
only assign options which are assigned to you in your own user record. The options you are
allowed to set are marked yes and no (not capitalized). Options you are not allowed to set
or change are marked YES and NO (capitalized).

b. Type the letter(s) in front of the specific option(s) you want to assign. Yes indicates the
option is on or allowed, and no indicates the option is off or not allowed.

c. Press HELP for more information and instructions.

8. Press SHIFT-BACK to return to the options index.

THE AUTHOR MODE DISPLAY

As an author, the first display you see after signing on to the PLATO system is the Author Mode display
(figure 4-2). This display is your navigational tool on the system. All the system resources you need can
be reached from this display. Unlike the PLATO Facilities display used by instructors, the Author Mode
display does not provide any initial options for you to choose from. Although this might be confusing for
new users, it allows more experienced authors to directly access the lesson they want to see or the
feature they want to use without selecting several options to do so.

HELP available

Figure 4-2. Author Mode Display

97405900 C 4-3

You can access a list of options frequently used by authors from the Author Mode display. This list of
options is the SHIFT-DATA display (figure 4-3) and is reached by pressing SHIFT-DATA from the Author
Mode display. The SHIFT-DATA display lists options frequently used by authors, gives a short description
of the options, and gives directions on how to access the options. All the options on the SHIFT-DATA
display can be accessed by typing the shifted letter of the option from the Author Mode display. For
example, to access the AIDS option, press SHIFT and type A from the Author Mode display.

OPTIONS

aids

bulletin board

charset

desk calculator

ECS usage

catalog of lessons

info on records/talk flags

HNMOODD

j-stack (not available here)
lineset

micro

notes

personal notes

questions (aids)

request prints

security code

time

user list

version

write common to disk
lesson X-search

zzz (alarm service)

J
L
M
N
P
Q
R
S
T
U
V
W
x
Zz

Press HELP for more information

On the AUTHOR MODE display, press the SHIFTed letter
to access the corresponding option immediately.

Figure 4-3. SHIFT-DATA Display

Not all users can access all options on the
SHIFT-DATA display. Accessibility of some
options is determined by the group the user is
registered in and the individual author options
selected for that user. Contact your account
director to gain access to unavailable options.

4-4 97405900 C

The SHIFT-DATA display functions as a reference tool which reminds you of what key to press to access a
specific feature. After a short time, you will learn which keys are associated with specific features and
will not need to refer to the SHIFT-DATA display for that information.

The following paragraphs briefly describe some of the most frequently used options on the SHIFT-DATA
display. The remaining options are described in Additional SHIFT-DATA Display Options later in this
section.

AIDS

AIDS is an on-line reference manual for authors and instructors which contains definitions and
explanations of most of the PLATO system features and all of the PLATO Author Language commands.
Authors frequently use AIDS as a reference tool when using the PLATO system. To access AIDS, press
SHIFT and type A from the Author Mode display. Refer to Using AIDS later in this section for more
information on AIDS and how to use this feature.

CATALOG OF AVAILABLE COURSEWARE

The Catalog of Available Courseware is a reference catalog which contains a listing of all published
courseware on the PLATO system. The catalog is used as a tool for locating courseware materials and for
providing information about courseware materials. Authors often refer to the Catalog of Available
Courseware as the F Catalog because it is accessed by pressing SHIFT and typing F from the Author Mode
display. Refer to Using the Catalog of Available Courseware later in this section for more information on
the catalog and how to use it.

NOTES

Notes are messages which allow users to communicate with each other and allow system announcements
to reach large numbers of users. Authors can access different kinds of notes files, depending upon the
notes options selected for their use and the access status of individual notes files. To reach the PLATO
Notes Options display, press SHIFT and type N from the Author Mode display. Refer to Notes later in this
section for more information on notes and how to use them.

PERSONAL NOTES

Personal notes are private messages between two users on the PLATO system. They allow users to
communicate on an individual and personal basis. To reach the Personal Notes display, press SHIFT and
type P from the Author Mode display. Refer to Using Personal Notes later in this section for more
information on personal notes and how to use them.

USER LIST

The PLATO User List displays the names of all authors and instructors who are currently signed on to the
PLATO system and who have voluntarily included their names in the list. Authors can see this
information and either include or remove their names from the list. To reach the PLATO User List, press
SHIFT and type U from the Author Mode display. Refer to Using the PLATO User List later in this
section for more information on the User List and how to use it.

97405900 C 4-5

PRINTS

The print feature allows users to request prints of files from the central site (central prints), or to print a
file or make copies of displays on the PLATO terminal screen using a printer attached to a PLATO
terminal (local prints). In order to request central prints, a user's account must contract to use the print
feature. No contractual agreements are necessary for local prints. Users whose accounts are on Control
Data Services Systems should have the print option in their author record turned on and contact their
account director to request central prints. The print feature also allows users to check the status of a
central print request previously made and the status of the central printer. To reach the print feature,
press SHIFT and type R from the Author Mode display. Refer to Requesting Prints later in this section
for more information on prints and how to request them.

UNDERSTANDING FILE STRUCTURE AND USE

The PLATO system can store large amounts of information for all users. Because it can store vast
amounts of information, the information must be organized in such a way that it can be easily retrieved
and used. All information on the PLATO system is contained in files. A file is simply a delegated amount
of space in the computer's memory in which information can be stored. It allows users to easily insert,
organize, and retrieve information. An easy way to understand the concept of files is to think of the
PLATO system as a huge filing cabinet. Each drawer in the cabinet can be compared to a file. Each
drawer (file) has a certain amount of space in which information can be stored or from which it can be
retrieved.

Most users use files in some way while using the PLATO system. Students use files when they study
lessons. The lessons they study are files containing code which presents the lesson. Instructors use files
to register students in curricula and to keep records of students’ progress. Authors use files to write
lessons as well as to collect data and communicate with other users.

There are several types of files available on the PLATO system. Different types of files are used to store
different kinds of information. The types of files you use is determined by the kind of information you
want to store and the kinds of things you want to do with the information.

Although file types differ as far as the kinds of information they are designed to store, most files are
structured similarly. Each file is assigned a certain number of parts. Each file part can store a specific
amount of data. The amount of data each part can store is determined by the maximum number of
computer words that can be stored in the part (usually 320 computer words can be stored in one part). A
computer word consists of a string of 10 characters. A character can be any typed letter, number, or
symbol. Spaces between characters count as a character, but blank characters at the end of a line do
not. Capital letters count as two characters. The number of parts assigned to the file determines the
file's length or size. The length of the file (the number of parts) is assigned by the account director when
the oa is created. (There is, however, a maximum number of parts which can be assigned for each type
of file.

Each type of file differs slightly due to the varying functions of the files, but all have some kind of
understandable structure which allows you to easily use the files. Most files have a central index which
allows you to organize information in a logical manner, as well as easily locate information and move
around in the file. The structures of the indexes for each type of file differ slightly. Some files are
indexed by blocks (as in TUTOR lesson files, figure 4-4a), while others might be divided into numbered
sections (as in documentor files, figure 4-4b). Whatever the structure, directions are always provided to
help you understand and effectively use the indexes.

4-6 97405900 C

LESSON--engl ish HELP available

Document dintro INSPECT ONLY

Preface/Introduction
Introduction to Documentor
What is Documentor?

What is a Document?

What is a Major Section?
What is a Subsection?
Document Size Limits
Accessing a Document

-1 Reading a Document

.2 Editing a Document

a

FRoHAO TH

bat
ee
—_—ee SAU kh WN oe

.
uw

ENO OF MAJOR SECTION ««

What command >
HELP available

Part B

Figure 4-4. Examples of File Indexes

97405900 C

4-7

Each type of file also has a directory which contains information about the author of the file; briefly
describes its contents; and contains basic information about the file, its security eodewords, and
associated files. This information is entered by the author when the file is used for the first time. This
directory display is reached by pressing DATA from the main index of most file types (SHIFT-DATA is the
keypress for notes file directories). Some examples of directory displays are in figure 4-5.

Document name -- usersged Disk pack -- epmast
Starting date -- 11724788 Account ---- support

Last accessed -- 12/84/68 13.43.48.
by ---- jean price of sinspect at 35-1
Last action ---- Inspect 7 edit section !

Type the appropriate letter:

Author information
Associated files
Codewords

Editing specifications

Document information

Press LAB for space usage information.

HELP available

Figure 4-5. Examples of File Directories (Sheet 1 of 2)

4-8 97405900 C

Lesson name ---- english Disk pack -- eamast
Starting date -- 89/15/88 Recount ---- chedd

Last edited ---- 11/86/68 18.41.22
by ---- jean price of adev at

Last action ---- Single block deletion

Type the appropriate letter:

futhor information
fisesociated files
Codeworde

Editing specifications

Micro PLATO level

HELP available

Figure 4-5. Examples of File Directories (Sheet 2 of 2)

Some files have one or more editors which allow you to insert or change information in the file. Editing
directives can vary for different types of files, but directions are always provided to help you use the
editors. Table 5-1 (refer to section 5) lists the different types of PLATO system files, their primary
functions, and maximum allowable sizes.

97405900 C 4-9

GETTING HELP

As an author, you can request and receive help from several sources while using the PLATO system. The
kind of help you request depends upon the type and extent of help you need, as well as the type of activity
you are involved in at the time you request help.

Authors can interact with the PLATO system on two levels — as students (seeing lessons as students see
them) and as authors (writing and editing lesson code and using the system resources provided for
authors). When you are using the PLATO system as a student (executing and seeing lessons as students see
them), there are usually programmed help sequences within the lessons which display helpful information
about the lesson. These sequences are programmed by the author of the lesson and often provide more
detailed information on topics presented in the lesson, information on how to proceed through the lesson,
or other related information. Programmed help is usually accessed by pressing HELP while using a
lesson. Many lessons display HELP or HELP available on the bottom lines of the screen to remind you to
use the HELP key.

When you are using the PLATO system as an author (using system resources provided for authors, editing
lessons, and so on), you can request either system-supported help or help from other users.
System-supported help includes features which the PLATO system provides to aid authors while using the
PLATO system. Some of these features, such as AIDS, allow you to refer to information about PLATO
Author Language commands or system features. Other features, such as TERM-ask, TERM-talk, and
TERM-consult, allow you to contact another author, instructor, or a PLATO system consultant and
communicate on-line to discuss questions or problems you have.

The following paragraphs describe the help features available to authors.

USING AIDS
AIDS is an on-line reference manual for authors and instructors. It contains definitions and explanations
of most of the PLATO system features and all of the PLATO Author Language commands. AIDS contains
more than 100 lessons which collectively form a complete reference manual for the PLATO Author
Language and system features. Authors and instructors frequently use AIDS as a reference too] when
using the PLATO system. The following list contains examples of the kinds of information available in
AIDS.

e Overview of the areas of the PLATO Author Language.

e Complete descriptions of all PLATO Author Language commands and system-defined variables.

e Author resources (lists of publications, names of PLATO systems personnel).

e Descriptions of helpful system features and lessons.

e Definitions of PLATO system terminology.

e Libraries of useful PLATO Author Language routines and character sets.

e Lists of commands (alphabetically and by function).

e Lists of indexes in AIDS.

4-10 97405900 C

Authors and instructors reach AIDS from different points in the system. As an author,

ou can reach AIDS

from either the Author Mode display or the Block display (if you are editing a lesson). From the Author
Mode display, you can either press SHIFT and type A, or type aids and press NEXT. The system displays
the AIDS Title display (figure 4-6). From the Block display, press SHIFT and type Q and then press NEXT
(this action takes you to the AIDS What TUTOR Feature display, described below). As an instructor, you
ean access AIDS from the PLATO Facilities display. Choose the AIDS option by typing the letter in front

of the option. The system displays the AIDS Title display (figure 4-6).

Fy cle, te ee

Elaine Avner, Darlene Chirolas,
Celia Davis, Jim Ghesquiere,
Tina Gunsalus, Jim Kraatz
and Judy Sherwood

PSO Author Group -- CERL
Univ of Illinois, Urbana

Press HELP if this is your first time in lesson AIDS

() Copyright, 1973, 1974, 1975, 1976, 1977, 1978, 1979, 1988
Board of Trustees of the University of Illinois

NO portion of the AIDS lessons may be reproduced
in a form without permission from the authors.
1241 features requested per day for the last 628 days

Figure 4-6. AIDS Title Display

From the AIDS Title display, you can do one of three things depending upon your needs.

@ Press HELP for more information on AIDS and how to use the feature.

e Press NEXT for the AIDS Index (figure 4-7). The AIDS Index consists of two displays (press
NEXT for the second display, BACK to return to the first) which are like a table of contents. It
presents a general overview of the information contained in AIDS. Choose an option which
generally covers the type of information you want to see by typing the letter in front of the
option. The system either displays the information or a second, more detailed index.

Press HELP for more information on how to use the AIDS Index.

97405900 C

Press a letter; or press NEXT for page 2

a Aids for new authors
How to use AIDS

Author Resources

Alphabetical list of TUTOR commands
Functional lists of TUTOR commands
List of Indexes in AIDS

Lists of System Defined Variables, Keynames,
Functions, Logical & Bit Operators, -specs- Tags

Making Displays

Making Graphs & Charts
Calculations and Variables
Conditional Operations
Sequencing

Judging

Execution of TUTOR

SHIFT-BACK always returns you to an index display.
HELP, DATA, BACK, SHIFT-NEXT are always available.

Figure 4-7. AIDS Index (Sheet 1 of 2)

4-12 97405900 C

Press a letter; or press NEXT for page 1

The PLATO Computer

Special Characters: ACCESS Characters, Linesets,
FONT Characters (Character Sets), & MICRO Keys

Student Data, Instructor Options, & Routers

Keynames, Keycodes, and Internal Codes

Programming Errors and Condense, Lesson, &
Execution Errors

Library of Author Routines
Microfiche and Photographing the Plasma Panel
Systematic Lesson Design

The Programmable Terminal (PPTs & ISTs)

SHIFT-BACK always returns you to an index display.
HELP, DATA, BACK, SHIFT-NEXT are always available.

Figure 4-7. AIDS Index (Sheet 2 of 2)

97405900 C 4-13

e Press DATA to bypass the index and see a display which allows you to type a specific command
or feature on which you want information. This display is the What TUTOR Feature display
(figure 4-8). Type the name of the command or system feature on which you want information
and press NEXT. The system displays the information.

What TUTOR feature? >

HELP for help on how to use AIDS

SHIFT-BACK for Main AIDS Index

SHIFT-DATA to make a comment about AIDS

Figure 4-8. AIDS What TUTOR Feature Display

For a quick reference, you can press DATA from anywhere in AIDS, see the What TUTOR Feature display,
and request information.

Press HELP from the What TUTOR Feature display for more information on how to use this display.

Refer to Quick Reference (Q-Ref) later in this section for information on how to access AIDS while
editing a TUTOR file.

4-14 97405900 C

CONSULTING HELP FOR AUTHORS AND INSTRUCTORS

Authors and instructors ean receive help while using the PLATO system by using the TERM-consult
feature. TERM-consult allows you to communicate with a PLATO consultant about questions or problems
you have while using the PLATO system. Consultants are Control Data systems personnel who are
available to answer questions and solve problems for users on the system.

The following paragraphs describe how to use the TERM-consult feature.
To contact a consultant, do the following steps.
1. Press TERM (hold the SHIFT key down while pressing the TERM/ANS key).

2. Type consult and press NEXT. The system responds by displaying either a message that says a
consultant has been notified of your request and will respond as soon as possible, or that no one is
available at this time but your name has been placed on a waiting list.

3. When a consultant answers your call, you see a message at the bottom of the screen indicating
the name of the consultant and that the consultant sees the same display that is on your screen
(for example, mary jones/pso sees this display).

If you solve your problem before a consultant
reaches you, you can cancel the request. To cancel
a request, press TERM and type consult again. The
system asks if you want to reaffirm your request or
cancel it. Press NEXT to repeat your request for
help, or press SHIFT-HELP to cancel it.

To talk with the consultant, do the following steps.

1. When the consultant answers your call and a message similar to mary jones/pso sees this display
appears, an arrow also appears at the bottom of the screen. Any message your consultant types
to you appears to the right of that arrow. You can communicate with the consultant by pressing
TERM. When you press TERM, a second arrow appears on the screen. This means you are in talk
mode. Talk mode allows both you and the consultant to type and respond to messages on the
sereen. As you type, your message appears to the right of the second arrow.

2. If your message requires more than one line of typing, press LAB to clear the line and continue
typing. The LAB key is the only key that allows you to continue typing. If you press any
function key other than LAB, the arrow disappears. You can only communicate with the
consultant when the second arrow is visible and you are in talk mode. If the arrow disappears,
press TERM to recall it and resume typing.

Sometimes it is helpful for a consultant to see what is on your screen in order to answer your question or
solve your problem. Showing the consultant your screen eliminates the need for you to describe in detail
where you are on the system and what is giving you problems.

In order for the consultant to see the same thing on his/her screen that you see on yours, you must replot

your sereen. Replotting your screen means to move from the display you are currently looking at to a new
display (you can always return to your original display again if that is the display you need help with).

97405900 C 4-15

To show the consultant your screen, do the following steps.
1. Tell the consultant you are replotting your screen.
2. Press BACK to leave talk mode, then go to the first display you want the consultant to see.
3. Press TERM to get into talk mode again and ask your question or discuss the display.

4. To show the consultant a different display, press BACK to leave talk mode. Go to the new
display and repeat step 3.

To end the consultation, do the following steps.

1. When your question has been answered and you do not need any further help, type thanks and the
consultant ends your communication.

2. The system displays a message saying the consultation is over.

USING TERM-ask

TERM-ask is a PLATO system feature which gives users (usually students and multiples) the opportunity
to ask authors and instructors questions about instructional materials while the materials are being
presented to them. It also gives authors and instructors the opportunity to discuss with other authors and
instructors questions or problems they might have while using the PLATO system. TERM-ask can be used
as an instructional aid by authors and instructors who want to discuss questions with students as they
arise, or for courses requiring dialog to reach objectives. It is particularly useful for students who are
having difficulty with some concepts in their lessons. It is also useful within author groups in which more
experienced authors would like to answer new authors’ questions about procedures, policies, specific
development projects, or programming in general.

TERM-ask is available to any user whose group is prepared to use the TERM-ask feature. The author or
instructor responsible for the group determines whether or not to allow users in the group to use the
feature, and also determines which authors and instructors in the group can respond to TERM-ask requests.

The following paragraphs describe how to prepare a group to use the TERM-ask feature, how to request
help from other authors and instructors, and how to respond to TERM-ask requests from other users.

Preparing the Group to Use TERM-ask

1. Doone of the following depending upon your user type.

a. From the Author Mode display (for authors), type the group name of the users who you want
to use TERM-ask. Press NEXT. The system displays the Group Operations display.

b. From the PLATO Facilities display (for instructors), choose the Group Operations option by

typing the number in front of the option. Type the group name of the users who you want to
use TERM-ask. Press NEXT. The system displays the Group Operations display.

4-16 97405900 C

2. Press DATA. The system displays the Group DATA (Directory) display.

3. Choose the Associated Files option by typing the letter in front of the option.

4. Type the number below the TERM-ask option. The system displays an arrow.

5. Type the group name of the users (authors and instructors) who you want to answer the
TERM-ask requests.

You can only enter one group name and it must be
a group for which you know the security code.

After you prepare the student group to use TERM-ask, you need to designate which users in your group
(the group in which you are registered) are available to receive TERM-ask requests and answer student
questions. Usually, this includes you and occasionally another author or instructor in your group who is
familiar with your curriculum. The following steps describe how to enable your group to receive
TERM-ask requests and answer student questions.

1. Doone of the following depending upon your user type.

a

b.

From the Author Mode display (for authors), type the group name of the user(s) who you
want to receive student TERM-ask questions. Press NEXT. (Remember, you ean only enter
one group name. Be sure it is the group in which you are registered and for which you know
the security code if you want to answer your students' TERM-ask questions.) The system

displays the Group Operations display.

From the PLATO Facilities display (for instructors), choose the Group Operations option by
typing the letter in front of the option. Type the group name of the user(s) who you want to
receive student TERM-ask questions. Press NEXT. (Remember, you can only enter one
group name. Be sure it is the group in which you are registered and for which you know the
security code if you want to answer your students’ TERM-ask questions.) The system
displays the Group Operations display.

2. Choose the See or Change Someone's Record option by typing the letter in front of the option.

97405900 C

Before you can set another user's record to allow
that user to receive TERM-ask requests, your user
record must have the Receive TERM~-ask Requests
option set to YES. Ask your account owner or an
account director to assign this option to you.

4-17

3. Type the name of a person in your group who you want to receive and answer TERM-ask
questions. Press NEXT.

4. Select Choose Allowable Instructor Options.
5. Choose Primary Instructor Options.

6. Locate the Receive TERM-ask Requests option. Set the option to yes by typing the letter in
front of the option.

7. Repeat steps 2 through 6 from the Group Operations display to enable other authors and
instructors in your group to answer TERM-ask requests.

Each author and instructor in the group who wants
to receive TERM-ask requests must set his/her own
TERM-ask user flag setting to YES in order to
receive TERM-ask requests. (This is in addition to
having the TERM-ask option turned on in his/her
user record.) Refer to User Statistics and Flag
Settings later in this section for information on
how to set your own user flags.

Sometimes you might not be signed on to the PLATO system when a student requests help. You can help
students at a later time if you have a student notes file attached to your student group. A student notes
file records information on which students request help while you are signed off the system. It collects
students’ unanswered questions and allows you to answer them at a later date by writing the student a
note. To use a student notes file to collect information while you are signed off from the system, attach
a student notes file to the student group which requests help using TERM-ask. The following steps
describe how to attach a student notes file.

1. Create a student notes file for your group through your account. Your account director can
create a student notes file for you if you do not have account director capabilities.

9. Go to the Group Operations display of the student up which uses TERM-ask to ask questions.
Press DATA. The system displays the Group DATA Directory) display.

3. Choose the Associated Files option by typing the letter in front of the option.
4. Type the number below the Student Notes option. The system displays an arrow.
5. Type the name of the student notes file to the right of the arrow.

4-18 97405900 C

Using TERM-ask to Request Help

To use TERM-ask to request and receive help from another author or instructor, follow the instructions
given for students in section 2, TERM-ask.

Responding to TERM-ask Requests for Help

When a student or another author or instructor uses TERM-ask to request help, the request is shown to all
authors or instructors who have been designated to receive TERM-ask requests for that group. The
request appears as a message on the bottom of the author's or instructor's screen indicating the name of
the user requesting help; the group in which the user is registered; and the site number, station number,
and system the person is using.

To respond to a user's TERM-ask request for help, authors should type ask on the Author Mode display and
press NEXT, and instructors should choose the Interactive Communications option on the PLATO
Facilities display by typing the letter in front of the option and then selecting the Respond to TERM-ask
Requests option on the Interactive Communications display. The system displays the TERM-ask Options
display. From this display, you can choose any of the following options. To use any of these options, type
the number in front of the desired option.

e See which users in your group are signed on and are registered as TERM-ask consultants.

e@ See which users are currently signed on in any group for which your group is the consulting
group. (This option allows you to see which users might request help).

e Do any of the following:
- See a list of pending requests for help from users in your group.

- Send a message to any user who has requested help. (For example, this option could be used
to say, "Tll be with you in a moment.")

- Monitor the screen of any user requesting help.

- Delete a request for help, indicating to other consultants that this user's problem has been
solved.

The list of users waiting for help is circular. Older requests are automatically overwritten, usually only

after several hours have passed. Remember that users who request help using TERM-ask are given an
option to write a note to the student notes file for their group when no one is available to help them.

TERM-talk

The TERM-talk feature allows you to communicate with another user who is currently signed on to the
PLATO system by typing messages back and forth on the bottom two lines of your screens.

To learn how to use TERM-talk, refer to Using the Talk Feature later in this section.

97405900 C 4-19

USING COMMUNICATIONS FEATURES *

The PLATO system provides authors and instructors with a wide range of communications features to
use. As an author or instructor, you can communicate with other system users in a variety of ways. You
can privately communicate with another user by typing messages back and forth on your screens and see
the messages as they are being typed, or you can write personal notes which can only be read by the
person to whom the note is sent. You can publicly communicate with several users by writing general
notes which several users can read and respond to. PLATO systems personnel use this type of feature to
communicate with all system users at one time. The communications features also allow you to
communicate with the author of a lesson you are using by writing a note to the lesson author while you are
using the lesson.

The following paragraphs describe the PLATO system communieations features which are available to
authors and instructors and how to use them.

USING THE TALK FEATURE

The TERM-talk feature allows you to communicate with another user who is currently signed on to the
PLATO system by typing messages back and forth on the bottom two lines of the screen.

The following steps describe how to access and use the TERM-talk feature.
1. Access TERM-talk in either of the following ways.
e Press TERM (hold SHIFT key down while pressing TERM/ANS key) from any location.
e Press DATA from the Total Users Display (from the PLATO User List).

2. An arrow appears at the bottom of your screen. Type the name of the user you want to talk to
and press NEXT.

3. Type the user's group name and press NEXT.

4. The system responds by paging (flashing a message at the bottom of the screen, such as
TERM-talk: mary smith/biology) the person you want to talk to (if that user is signed on), or it
informs you that the person is unavailable (not signed on) or is busy.

5. When the person you are paging answers, two arrows appear at the bottom of your screen. Type
your message and press NEXT at the end of the line of typing.

6. Before ending a TERM-talk conversation, PLATO etiquette suggests ending with good bye or
another indication to the other party that the conversation has ended. Press SHIFT-BACK to end
TERM-talk.

The TERM-talk feature is not always available. Authors sometimes code lessons to inhibit specific TERM
features from working in some of their lessons. Generally, this coding technique is not recommended
unless the use of a TERM would defeat the lesson objectives (for example, use of TERM-cale could defeat
the objectives of an arithmetic test).

t+ Refer to inside back cover for important regulatory notice concerning the use of communications features.

4-20 97405900 C_

If you are taking a lesson which does not allow a specific TERM feature, the TERM feature is
unavailable. If someone tries to TERM-talk with you while you are using a lesson that is programmed not
to allow TERM-talk, you will see a message indicating that someone wants to talk with you but, when you
press SHIFT-TERM, the message What Term? and the arrow will not appear. In order to talk to that
person, you must leave the lesson (SHIFT-STOP).

There may be times when you are using the PLATO system when you do not want your work to be

interrupted by a user paging you to TERM-talk. TERM-busy and TERM-reject are two features you can
use to either make yourself unavailable for TERM-talk or to reject a TERM-talk.

TERM-busy

TERM-busy allows you to turn off your TERM-talk feature. When users try to TERM-talk with you, they
receive 4 message saying you are busy but have been told they called. The system notifies you when
someone wants to TERM-talk with you.

The following steps describe how to use TERM-busy.
1. Press TERM (hold down SHIFT key while pressing TERM/ANS key) from any display.

2. Type busy and press NEXT. System responds with a message saying you are unavailable for
TERM-talk.

3. When someone tries to TERM-talk with you, the system tells the user you are busy but have been
notified they called. The system also tells you who wants to talk with you.

4. To clear your TERM-busy status, press TERM, type busy, and press DATA. Signing off also
clears TERM-busy.

You can also set your TERM-busy status from the User Flags display. Refer to User Statistics and Flag
Settings later in this section for information on how to use this display.

TERM-reject

TERM-reject allows you to reject a TERM-talk from another user, but gives you the option of leaving a
message for the user.

The following steps describe how to use TERM-reject.

1. When the system pages you to TERM-talk with another user, press TERM, type reject, and press
NEXT.

2. The system responds by giving you the option of typing a message (40 characters, maximum) to
the user. Type your message and press NEXT.

3. The system notifies the user that you are busy and displays your message.

97405900 C 4-21

COMMENTING ON LESSONS AND FEATURES

You can comment on lessons or features on the PLATO system by using TERM-comment.
TERM-comment allows you to write a message about a lesson or feature on which you have questions or
comments and to send it to the lesson author.

Depending upon what you are commenting on, your comments may be seen by people other than the lesson
author. Lessons and features are divided into three categories: published lessons, privately-owned
courseware, and system features. Published lessons are listed in the Catalog of Available Courseware and
their file names usually begin with a 8 System features include lessons like notes, your account, the
Catalog of Available Courseware, the User List, your group file, and so on. If you comment on a published
lesson, your comment goes to the lesson author or to Control Data personnel maintaining published
courseware. If you comment on privately-owned courseware, the author receives the comment if he/she
has a notes file attached to her/his lessons or files. If you comment on a system feature, your comment
will first be read by consultants in group p or pso. Often a consultant contacts you to answer your
question. If your comment reports a problem, the consultant forwards the information to the personnel
who maintain the system.

To learn how to use TERM-comment, refer to Commenting on Lessons in section 2.

NoTES T

Notes are indexed messages stored in files in the PLATO system. They allow users to communicate with
each other and allow system messages to reach large numbers of users. Each note can have a number of
responses. Notes can contain questions or informative material on any topic.

Types of Notes Files
There are several types of notes files on the PLATO system. The different types of notes files are:

Personal notes Private notes between two users on the system. Only the addressee of a
personal note ean read the note.

General notes Notes between members of a defined user community. General notes allow
members to read or participate in a group discussion. General notes are also
used to relay system announcements to all users.

Intersystem notes Notes that are part of a connected notes file which enables the notes and
responses written in one notes file to appear in other connected notes files on
other PLATO systems.

Student notes Notes between students and instructors. Student notes can be similar to
personal notes or general notes. Students and instructors can communicate
privately or all students can participate in a group discussion or receive
messages from their instructor.

Lesson notes Notes about a lesson written by users while they are studying the lesson.

+ Refer to the inside back cover for important regulatory notice concerning the use of communications
features.

4-22 97405900 C

Notes File Access

Not all users can participate in all notes files. Only users who are allowed access to a given notes file can
read or write notes. Within each notes file is a table of groups and names of users within the groups from
which the access status of users is determined. When you request access to a notes file, the system
searches for your group name, your PLATO name, or your user type, depending upon what criteria are
established for user access. If the system finds your name, group, or user type in the allowed access
column, you are granted access; if not, the system tells you you are not allowed access to the notes file.

Using General Notes

General notes are notes between members of a defined user community. General notes can be discussion
files on topics of general interest, or they can be announcements to users on the system. Access to
general notes can be open to all students, authors, and instructors, or can be restricted to a smaller group
of users. The following general notes files can be accessed by all authors and instructors.

System Announcements

System announcements are notes to all users from PLATO systems personnel (users responsible for
maintaining the PLATO system). They frequently describe new commands and editing features. Only
PLATO systems personnel can write notes in this file. All other users can only read the notes in the file.

It is important for you to check the system announcements notes file periodically, as changes to the
system can affect your work. All authors and instructors should include this notes file in their notes file
sequencer. (Refer to Using the Notes File Sequencer later in this section for information on how to use
the sequencer.)

Three times a year, the PLATO system software is updated on all Control Data service systems. These
changes are referred to as cuts of the PLATO system software. This manual is accurate as of Cut 22.
Before a new cut is installed, an announcement of the new cut will be placed in this notes file along witha
description of new features and additions to existing features. All authors and instructors should check
the system announcements notes file at least once a week (if not daily) for new system changes.

Sometimes, when a new cut is installed, changes to the structure, nature, and function of existing files is
necessary. These changes are usually automatic and are referred to as file conversions.

97405900 C 4-23.

The following steps describe how to access the system announcements notes file.
Authors (do one of the following steps):
® Type announce on the Author Mode display and press NEXT.

e@ Hold down the SHIFT key and type N from the Author Mode display. The system displays
the PLATO Notes display (figure 4-9). Choose the System Announcements option.

--PLATO NOTES--

99729 16.33

Choose an option >

System Rrmouncements
Publ ic Notes

Other Notes

Personal Notes

Notes Files Sequencer

Press HELP for information.

Figure 4-9. PLATO Notes Display

Instructors:
1. Select the Notes option from the PLATO Facilities display. Type the number in front of the
System Announcements and Public Notes option. The system displays the Notes Options
display (figure 4-10).

2. Type the number in front of the System Announcements option.

4-24 97405900 C

Notes Options

1. Public notes and system announcements

2. Personal notes

3. Student notes

Figure 4-10. Notes Options Display

Follow the procedure for reading general notes described in How to Read General Notes later in this
section.

Public Notes
Public notes are a public discussion or forum on subjects of general interest to PLATO users. Public notes
are often used to discuss problems encountered while programming, to report errors in the system, or to
make suggestions for improving the system. The name of the public notes file is "pbnotes".
The following steps describe how to access the public notes file.
Authors (do one of the following steps):
@ Type pbnotes on the Author Mode display and press NEXT.

e Hold down the SHIFT key and type N from the Author Mode display. The system displays
the PLATO Notes display (figure 4-9). Choose the Publie Notes option.

Instructors:

1. Seleet the Notes option from the PLATO Facilities display. Type the number in front of the
System Announcements and Public Notes option. The system displays the Notes Options
display (figure 4-10). Type the letter in front of the Public Notes option.

Follow the procedure described in How to Read General Notes.

97405900 C 4-25.

How to Read General Notes

The first display you see after you gain access to a general notes file is the Notes File Index display
(figure 4-11). The notes file index contains a complete list of notes in the file and directions on how to
read, write, and respond to notes.

Control Data PLATO Public Notes
9/26 Welcome!
Wel come! xs intersystem notesfile «**
no “error”
ist2 zttype?
bitzer lecture
9/728 FOREST
new notes file What note? >
9/29 filestatus

axe End of Notes ax

Press LAB for file policy

SHIFT LAB to write a note

SHIFT BACK to exit

Press HELP for information

Figure 4-11. Notes File Index Display

The following information is included in the index: the identifying number of each note, the date the note
was written, the title of the note, and the number of responses to the note. The note number is the
identifying number you type to request to read a note. The date of the note is the date the note was
written. The title of the note gives a general idea of the subject of the note. The number of responses is
the number of written responses to the note.

4-26 97405900 C

The following steps describe how to read a general note.
1. The Notes File Index display contains the names of the nine most recently written notes.
a. To select a particular note to read, type the number of that note and press NEXT.
b. To see an index of the nine notes written previously, press BACK. To read a particular note,
type the number of that note and press NEXT. Repeated pressing of the BACK key displays

the nine notes written before the last notes you see in the index.

ce. To see the index of the first notes in the notes file (without seeing the index information for
all notes in the file), press SHIFT -.

d. To return to the end of the notes file index (the most recently written notes), press SHIFT
+. You can always tell when you are at the end of the notes file because the last note in the
file is followed by an End of Notes message.

2. After you type the number of the note you want to read and press NEXT, the system displays the

note for you to read. You can either continue to press LAB to read each response to the note, or
press NEXT to read the next note in the file.

Responses to notes are not regarded as new notes
(titled, numbered, and dated) because they all
relate to the same topic.

How to Read Archived Notes

Notes files have a maximum length; therefore, it is sometimes necessary to store past-dated notes to
make room for more recent notes. Stored notes are called archived notes. The following steps describe
how to read archived notes.

1. From the Notes File Index display, type a and press BACK. The system takes you to the most
recently archived notes file. Follow the procedure for reading general notes to read archived
notes.

2.  Tosee another archived file, repeat step 1.

3. To leave the archived files and return to the most recent notes file, type new and press NEXT.
Remember to look for the End of Notes message following the last note in the file.

Although archived notes can continue to be read,
they cannot be responded to.

97405900 C 4-27

Using the Notes File Sequencer

You ean access all of your notes files by using the notes file sequencer. The notes file sequencer takes
you from one notes file to another without returning to the PLATO Notes display or the PLATO Facilities
or Author Mode display. It also directs your attention only to those notes and responses which have been
written since the last time you read notes. For example, if you frequently read or participate in notes
files A, B, and C, you can use the notes file sequencer to see the new notes and responses in file A, file B,
and file C, and then return to the PLATO Notes display. The notes file sequencer also allows you to
change the order of the notes files or skip from one notes file to another.

The following steps describe how to register files in the notes file sequencer.

1. To reach the notes file sequencer, choose the Notes File Sequencer option from the PLATO
Notes display. The system displays the Sequencer Editing Options display (figure 4-12).

2. Type the name of a notes file you want to include in the sequencer. Press NEXT.

-- Sequencer Editing Options --

Enter file name >

HELP available
BACK for sequencer
SHIFT-LAB for bypass options

Figure 4-12. Sequencer Editing Options Display

3. The system asks in which numerical position you want the file listed (if this is not the first file
you entered in the sequencer). You can either type the number of the position in which you want
the file listed, or press NEXT to add the file to the end of your list.

4, Repeat steps 2 and 3 until you have included the names of all the files you want in the
sequencer. The maximum number of files allowed is 60.

5. Press HELP for more information on the notes file sequencer.

4-28 97405900 C

You can also delete files from the notes file sequencer, or move files to a different position in the list.
The following steps describe how to do these procedures from the Sequencer Editing Options display.

1. To delete a file, type the name or number of the file and press SHIFT-HELP to delete the file
from the sequencer.

2. To change a file's position in the sequencer, type the name or number of the file you want
changed and press NEXT. Press LAB. Type the number of the new position you want the file to
appear and press NEXT.

3. Press HELP for more information.

4. The notes file sequencer also provides several other options which help you access new notes
very quickly. These options are accessed by pressing SHIFT-LAB from the Sequencer Editing
Options display.

After you have registered a list of notes files in the sequencer, you ean access them one at a time by
pressing DATA from the PLATO Notes display. Continually pressing DATA takes you to each new note
and response. When a response has been written to an old note, you are first shown the note. The next
DATA keypress takes you to the new response.

Additional General Notes Features

There are several features available to you while you are reading a general note. Some of these features
allow you to send a personal note to the author of the note you are reading, talk (TERM-talk) to the
author of a note, and see the current time and date.

To see a complete list of the features available to you while reading a general note, press HELP while
reading a general note.

Using Personal Notes

Personal notes are private notes between two users on the PLATO system. They allow two PLATO system
users to communicate on an individual and personal basis. Only the individual who wrote the note and the
person to whom the note is addressed can read the personal note.

To write a personal note, you need to access the Personal Notes display (figure 4-13). Authors and

instructors access this display in different ways. The following steps describe how authors and instructors
access the Personal Notes display.

97405900 C 4-29

PERSONAL NOTES

Press: LAB to read your notes
DATA for other options
HELP for explanation and policy

To whom do you wish to send a note:

Name >

Group

System

Figure 4-13. Personal Notes Display

Authors (do one of the following steps):

e From the Author Mode display, press SHIFT and type P. The system takes you to the
Personal Notes display.

e From the Author Mode display, press SHIFT and type N. The system takes you to the
PLATO Notes display. Type the letter in front of the Personal Notes option. The system
takes you to the Personal Notes display.

Instructors:

1. From the PLATO Facilities display, choose the Notes option by typing the number in front
of the option. The system displays the Notes Options display (figure 4-10).

2. Type the number in front of the Personal Notes option. The system displays the Personal
Notes display.

From the Personal Notes display, you can read, write, and respond to personal notes. You can also save,
copy, and forward personal notes. The following describes how to read and write personal notes.

4-30 97405900 C

How to Read Personal Notes

When someone sends you a personal note, the Author Mode display (for authors) or the PLATO Facilities
display (for instructors) displays a message indicating you have a personal note to read.

To read your notes, go to the Personal Notes display and press LAB. The system displays your first new
personal note.

After you read your note, you can do one or more of the following steps.
e Press NEXT to read the next personal note addressed to you.
@ Press SHIFT-LAB to respond to the note.
e Press SHIFT-HELP to delete the note.
e Press BACK to read previous notes.

@ Press SHIFT-BACK to return to the Personal Notes display.

How to Write Personal Notes

You can write a personal note from the Personal Notes display, or respond to a note written to you while
looking at the display of the note to which you want to respond.

To write a personal note, go to the Personal Notes display and type the name, group, and system of the
person to whom you are writing the note. Press NEXT. The system displays either the Insert Mode
display (for authors) or the Easy Editor display (for instructors). Type your note using the same procedure
described for general notes in Writing General Notes later in this section. Press SHIFT-NEXT to send
your note or press SHIFT-BACK to cancel the note.

Using Lesson Notes

Lesson notes are comments about a lesson written by users studying the lesson. They are stored in a
lesson notes file which is attached to the lesson by the lesson author.

Notes reach the lesson notes file through TERM-comment. When a student makes a comment using the
TERM-comment feature, it is stored in the lesson notes file. TERM-comments are usually written to
point out problems within a lesson such as unclear instructions, confusing explanations, incorrect answers,
and so forth. From these comments, authors can rewrite their lesson to make it more understandable.
Lesson notes files are usually used only during the early development and testing of a lesson or set of
lessons.

As an author, you can attach a lesson notes file to your lesson from the lesson Block Listing display (figure

4-14). From the lesson Block Listing display, press DATA and choose the Associated Files option by
typing the letter in front of the option. Choose the Lesson Notes File option.

97405900 C 4-31 —

LESSON--engl ish HELP available

BLOCK NAME
- a sdirectory
b grammar
c spelling
d sentences
e letters
f

chars a*x one

Figure 4-14. Block Listing Display

Although instuctors do not write lessons and therefore cannot set up lesson notes files, they can still
receive TERM-comments from users about a lesson which is ineluded in their curiculum. If an instructor
creates a student notes file, all TERM-comments made by students in the group containing the student
notes file about lessons in that group's curriculum are forwarded to the student notes file instead of to the
author's lesson notes file. Authors and instructors should refer to the following section, Using Student
Notes, to learn how to create a student notes file and read and respond to student notes.

Using Student Notes

A student notes file has notes features specifically for student sign ons. It can provide private notes
between the student and the instructor or general notes for all students in the group. This allows a
student to send a comment or question to his/her instructor and receive a personal response. In the
context of student notes files, an instructor is any author or instructor who has access to the student
notes file.

When a student is working on a lesson and writes a TERM-comment, that comment goes into the student
notes file. The system router, "mrouter", provides a direct entry into student notes for the student, and
other routers can do so if the author codes them to do so.

The system router, "mrouter", and other routers can also allow students to enter the student notes file and
read all of the notes, as if it were a general notes file. This, however, defeats the privacy aspect of the
student notes file. The only advantage of using a student notes file in this manner is that it combines the
funetions of two files: the file behaves much like a general notes file, and TERM-comments are routed
into this file instead of into the lesson notes file. If space is available, it is better to use two distinct
files: a student notes file for private notes, and a general notes file for public discussions.

For information on how to read, set up, and use a student notes file (using "mrouter" or your own router),
refer to AIDS and type student notes on the What TUTOR Feature display.
4~32 97405900 C

Using Intersystem Notes

Authorized users ean connect notes files between and among PLATO systems, and thereby transfer notes
and responses written in each notes file to other connected notes files on other systems. The result is a
maximum of 34 general notes files, each on a different PLATO system, that combine the communications
of all user communities. All notes and responses written in any connected files on any system appear in
all connected files on all systems. Refer to Using Network Options in section 5 for more information on
intersystem notes.

Writing General Notes

You can write a general note to start a discussion about a specific topic, or to respond to a note already
entered in the general notes file. Authors and instructors usually use different editors to write and
respond to notes. Instructors use the easy editor and authors can use either the easy editor or the
standard editor.

The standard editor is automatically assigned to
authors to use. To switch to the easy editor, press
BACK from the Insert Mode display (described
below in step 2 of Using the Standard Editor), and
press SHIFT-DATA.

The following sections describe how to use each of these editors.

Using the Standard Editor (Authors)
1. Do one of the following steps, depending upon the kind of note you want to write.

e To write a new note (a note on a new topic), press SHIFT-LAB from the Notes File Index
display.

@ To respond to a note already written, press SHIFT-LAB from the display of the note to
which you want to respond. For example, to respond to note 10, your screen must show the
text of note 10 or an existing resonse to note 10. Press SHIFT-LAB.

97405900 C 4-33 -

4-34

2.

After you press SHIFT-LAB (from either location), the system displays the Insert Mode display
(figure 4-15). The Insert Mode display is marked by the words INSERT MODE in the upper right
corner of the screen and an arrow on the left side of the screen. The upper left corner of the
display tells how much space is available for you to insert text. Most general notes are
restricted to 20 lines (or 120 computer words). Each time you press NEXT after typing a line of
text, the space available numbers change to reflect how much space is left for your note. Insert
mode means the system is ready for you to insert (type) lines of text.

Space left: 128 words or 28 Lines INSERT MODE

Press NEXT at the end of every line
BACK to end insert mode

Figure 4-15. Insert Mode Display

Press HELP for more information on writing notes.

Type your note. Press NEXT at the end of each line to continue typing. Pressing NEXT
advances the arrow one line at a time.

Press BACK when you are finished Br: your note. Pressing BACK takes you out of insert
mode and allows you to proofread and edit (change or correct) your note.

97405900 C

6. The following editing directives help you edit your note.

R

The R directive is the replace directive. It is used to replace or change text within a
line. To use the replace directive, type r (you do not need to capitalize it), type the
number of the line you want changed, and press NEXT. For example, to correct an
error in line 3, type r3 and press NEXT. The system is now in replace mode. In replace
mode, the system displays the line you want to edit with an arrow directly under it.
You ean either retype the entire line correctly, or use the COPY and EDIT keys to edit
the line. (Refer to appendix A to learn how to use the COPY and EDIT keys.) After
you complete your corrections, press BACK.

The I directive is the insert directive. It is used to insert new material into the text of
your note. It puts the system into insert mode. To use the I directive, type i (you do
not need to capitalize it), type the number of the line you want the new material to
follow, and press NEXT. For example, to add text after line 8, type i8 and press
NEXT. To add text to the beginning of a note (before line 1), type i0 and press NEXT.
Type your additional lines and press BACK.

The F directive is the forward directive. It is used to move lines of text forward on
your screen. To use the F directive, type f (you do not need to capitalize it), type the
number of lines you want the screen to advance, and press NEXT. For example, if lines
1 through 30 are displayed on your screen and you want to see the four lines which
follow line 30, type f4 and press NEXT. The screen now displays the next four lines
(and the text which follows those lines).

The B directive is the backward directive. It is used to move lines of text backward on
the screen. To use the B directive, type b (you do not need to capitalize it), type the
number of lines you want the screen to back up, and press NEXT. For example, if lines
1 through 30 are displayed on your screen and you want to see the three lines which
precede line 1, type b3 and press NEXT. The system now displays the three lines
preceding the previous line 1. To see the beginning of your note, type b20 and press
NEXT.

The D directive is the delete directive. It is used to delete lines of text from the top of
your note. To use the D directive, type d (you do not need to capitalize it), type the
number of lines you want deleted (from the top of the displayed text), and press
SHIFT-HELP. For example, to delete the top two lines of your note, type d2 and press
SHIFT-HELP. Since the D directive only deletes text from the top of your screen, you
must use the F and B directives to position the line(s) you want deleted at the top of
your screen. For example, if lines 1 through 30 are visible on your screen and you want
to delete line 3, type f2 (to move the screen up two lines) and press NEXT. Line 3 is
now at the top of the screen (it is renumbered line 1, however, since the first line on
the screen is always line 1), and can be deleted by typing d or dl (both have the same
effect) and pressing SHIFT-HELP.

SHIFT-HELP is a special keypress used for deleting text. It is used for deletions
because it is not likely to be pressed accidentally. A way to be sure the system deleted
the lines you requested is to check the Space Available entry at the top of the note.
When text is deleted, the space available increases.

7. After you write and proofread your note, do one of two things.

97405900 C

Press SHIFT-NEXT to send the note and include it in the general notes file.

Press SHIFT-BACK to cancel the note.

4-35

Using the Easy Editor (Instructors and Authors)

1.

To write a general note, do one of two things.

e To write a new note (a note on a new topic), press SHIFT-LAB from the Notes File Index
display.

e To respond to a note already written, press SHIFT-LAB from the display of the note or
response to which you want to respond. For example, to respond to note 10, your screen
must show the text of note 10. Press SHIFT-LAB.

After you press SHIFT-LAB (from either location), the system displays a rectangular box with an
arrow in the upper left corner. Some instructions and editing directives are listed below the box.

Press HELP for more information on how to use the editing directives.

Type your message. The text appears to the right of the arrow. Press NEXT at the end of each
line to move the arrow to a new line. Your note can be up to 20 lines long.

If you make a mistake and need to change a line, move the arrow to the line you want to change
by pressing NEXT or BACK. (BACK moves the arrow up, NEXT moves the arrow down.)

When the arrow is pointing to the line you want to change, press EDIT. Pressing EDIT erases the
entire line. You can bring back the sentence one word at a time by pressing EDIT again, one
press for each word. Make your corrections by using the ERASE key or inserting new words.

You can insert a new line by pressing SHIFT-LAB. Position the arrow at the sentence directly
beneath the line after which you want the space inserted. Press SHIFT-LAB.

To delete a line, move the arrow to the first line you want deleted and press SHIFT-HELP. The
system then numbers the remaining lines and asks you how many lines you want deleted. Type
how many lines you want deleted and press NEXT.

When you finish writing and correcting your note, do one of two things.

e Press SHIFT-NEXT to send the note and include it in the general notes file.

e Press SHIFT-BACK to cancel the note.

Using Notes File Features

When you are reading general, personal, or student notes, there are a number of options available to you.
Press HELP while reading a note to see a list of these options (figure 4-16). Most of the options are
self-explanatory and easy to use. The following paragraphs describe some of these options in more detail.

4-36

97405900 C

While reading notes, you may press:

NEXT to go on to the next note

BACK to go to the previous note

LAB to go to the next note or response

DATA to skip to next note or response written

since a certain date and time
SHIFT-DATA to skip to next note since a certain
date and time
SHIF T-BACK to return to the directory
SHIFT-LAB to respond to a note
SHIFT-EDIT to edit or delete a note
CYou may only edit notes that you
have written, and may not edit a
note after responses have been made.)
SHIFT-COPY to copy a note to another notes file
“n* to replot the note
SHIFT-"x" to search for a certain note title
SHIFT-"s" to Save a note in your save buffer
SHIFT-“a" to Append a note to your save buffer
SHIFT-“d" to display current date and time
SHIFT-"t* to Talk to the note author
SHIFT-"p” to send a Personal note to the author
SHIFT-"f" to Forward a note via personal notes

Several other keys are available to facilitate
moving through long chains of responses to a note:
A number key (e.g., 4°) moves forward that number.
A shifted number moves forward that number plus 18.
"+" moves forward 1.

"-" moves backward 1.

SHIFT-"+" jumps to the last response.
SHIFT-"-" jumps to the base note.

Figure 4-16. Notes File Options Display

Saving Notes

You can save a note in one notes file and move a copy of it to another file or to another position in the
same file by using the save directive. From the display of the note you want saved, press the SHIFT key
and type S.. The system stores your note and displays the number of words and lines saved. Go to the file
where you want the note inserted, press SHIFT-LAB to create a new note, and press BACK to end insert
mode, then type is and the number of the line directly above where you want the saved material inserted
(for example, to insert saved material after line 8, type is8 and press NEXT). The system displays your
note in a new location.

The save directive can only be used once before
you transfer the saved material to a new file. If
you save additional material before transferring
the original saved material, the system deletes the
original material and replaces it with the new
information. A maximum of 320 computer words
can be saved at one time.
97405900 C 4-37 |

You can, however, add more information to the save buffer without destroying the original information by
using the append directive. The following section describes how to use the append directive.

Appending Notes

You can save additional material, up to a total of 320 computer words, by using the append directive. The
append directive allows you to add more material without deleting the original information. From the
display of the material you want added, press SHIFT, type A, and type the number of lines you want

saved. Press NEXT. For example, to append (save) lines 1 through 4, type A4 and press NEXT. The
system includes the added information with the saved material.

Copying Notes
You can copy notes from your personal notes file or a general notes file to another notes file with the
copy directive. From the display of the note you want to copy, press SHIFT and COPY. The system

responds with copy to what file? >. Type the name of the notes file you want the note copied to, and
press NEXT. The system copies the note to that file.

Forwarding Notes

You can forward a note from your personal notes file or a general notes file to your own or another user's
personal notes file by using the forward directive. From the note you want to forward, press SHIFT and
type F. The system takes you to the Personal Notes display. Type the name, group, and system of the
user to whom you want the note forwarded. The system displays three options. Press SHIFT-NEXT to

forward the note, SHIFT-LAB to edit the note before you send it, or DATA to write a new note (and
cancel the forwarded note).

Using Notes File Director Options
Every notes file on the PLATO system is under the direction of a notes file director. The notes file
director is responsible for the general management of the notes file. Some of the notes file director
responsibilities are:

e@ Maintain user access list for the notes file.

e Establish notes file use and policy.

e Lengthen and shorten file space as necessary.

e Delete inappropriate notes in the notes file.

e  Allow/disallow connection to notes files on other systems.

e Archive the notes file.

4-38 97405900 C

USING REFERENCE TOOLS

The PLATO system contains several reference tools which authors can use to access on-line reference
materials. Some of these tools are available only to authors, while others are available to all user types.
Some examples of the kinds of reference materials available to authors are: a directory of PLATO system
authors, a list of users currently signed on, personal system use statistics, the Catalog of Available
Courseware, the current time and date, and information about PLATO Author Language commands and
system features (AIDS).

The following paragraphs describe the PLATO system's reference tools.

AIDS

AIDS is an on-line reference manual for authors and instructors which contains definitions and
explanations of most of the PLATO system features and all of the PLATO Author Language and Micro
PLATO Language commands. Authors frequently use AIDS as a reference tool when using the PLATO
system. To aecess AIDS, press SHIFT and type A from the Author Mode display. Refer to Using AIDS
earlier in this section for more information on AIDS and how to use it.

USING THE CATALOG OF AVAILABLE COURSEWARE

The Catalog of Available Courseware is a reference catalog which contains a listing of all published
PLATO-based courseware on the PLATO system. The catalog is used as a tool for reviewing courseware
materials and for providing information about courseware materials. Authors often refer to the Catalog
of Available Courseware as the F Catalog because it can be accessed from the Author Mode display by
pressing SHIFT and typing F.

The Catalog of Available Courseware is similar to the standard ecard eatalogs used in libraries. The
Catalog of Available Courseware, however, contains its information on-line rather than on actual cards.
The Catalog of Available Courseware contains four indexes of courseware materials arranged by title,
author, subject, and file name. From each index, you can see detailed deseriptions of courseware
materials and, in many cases, try the actual courseware materials.

The Catalog of Available Courseware contains a list of all published lessons and curricula on the PLATO
systems. Generally, there are two kinds of lessons and curricula available to users: published courseware
and proprietary courseware. Published courseware is courseware which is copyrighted and available on all
PLATO sytems. Before publication, the courseware is tested and reviewed to ensure the lessons operate
properly, that all function keys work as described, and that there are no coding errors which could cause
the lesson to work incorrectly. Published courseware is well maintained and reliable. It is never
unexpectedly revised or deleted from the system. Proprietary courseware, on the other hand, is
courseware which an author owns and has made available to other users. Proprietary courseware is not
copyrighted and has not been tested or reviewed by lesson publishing personnel. It is available only on
PLATO systems on which the author of the courseware resides. Proprietary courseware can be rewritten,
changed, or deleted without warning. With proprietary courseware, you have no guarantee that the lesson
is stable. Instructors should be extremely cautious about using proprietary courseware in their lessons,
are recommended to use lessons for a short period of time (1 or 2 weeks) only, and to contact the lesson
author before doing so.

97405900 C 4-39

Authors and instructors reach the Catalog of Available Courseware in different ways. Authors can access
the catalog from the Author Mode display by pressing SHIFT and typing F. The system displays the
Courseware Catalog Options Index (figure 4-17). Instructors can access the Catalog of Available
Courseware from the PLATO Facilities display. Select the Choose a Lesson to Study option by typing the
letter in front of the option. At the next display, press LAB to see the Courseware Catalog Options Index.

From the Courseware Catalog Options Index, you ean see an overall description of the eatalog and its
features and see courseware materials arranged by title, author, subject, and file name.

To see the overall description of the catalog and updated courseware information, type a from the
Courseware Catalog Options Index. The system displays a list of options describing features and aspects
of the Catalog of Available Courseware. Select an option by typing the number in front of it.

Catalog of Available Courseware

(a) How to Use This Catalog

Materials Arranged by:

Title
Authors
Subj ects

Filenames

Press LAB for New Titles
Press SHIFT-LAB for Courseware Queries
Press SHIFT-NEXT for Catalog Display Options

Figure 4-17. Courseware Catalog Options Index

To see courseware materials, type the letter in front of the type of index you want to see. For example,
to see a listing of courseware materials arranged by title, select the title index option. The following
paragraphs describe how to use the title, author, subject, and file name indexes.

4-40 97405900 C

Title Index

The title index (figure 4-18) alphabetically lists the titles of all the PLATO-based courseware materials.
Although there is one main alphabetical title index, lists of titles also appear throughout the catalog under

subject headings and author names.

Catalog of Available Courseware

Al phabetized Title Index

Type in first few letters of Title

>

Press NEXT for beginning of list

Figure 4-18. Catalog of Available Courseware Title Index

From the title index, you can see descriptions of courseware materials and, in some cases, try the actual
lessons. The following steps describe how to move through the title index.

97405900 C 4-41

1. You can page through the index in sequential alphabetical order (a, b, c, and so on), or you can
quickly move to any part of the index (for example, from b to s). Press NEXT to advance the
index alphabetically one page at a time and press BACK to reverse the index one page at a time.

2. To move quickly to a different alphabetical point in the index, type a few letters of the part of
the alphabet you want to see and press NEXT. For example, if you want to see titles that begin
with ple, type ple and press NEXT.

3. You ean type a specific lesson title and press NEXT to see that specific lesson.

To see the description of courseware materials or to have the option to try a lesson, type the number in
front of the title on which you want information and press NEXT. The system displays the Lesson
Information display (figure 4-19).

Neuron structure and function

BY:. Stephen H. Boggs
University of Illinois

COPYRIGHT DATE: 1975 FILENAME: Sneurons
LIBRARY TYPE: Academic

This learning activity uses graphics to explain the structure

and function of neurons and the way in which they transmit
information.

(a) Further Information fo) Authors

Press the letter of the option you wish to select.
>

LAB to try this item
SHIFT -NEXT/SHIFT-BACK to move BACK to exit

Figure 4-19. Lesson Information Display
4-42 97405900 C

The Lesson Information display contains a brief description of the lesson; the name of the primary
author(s) of the lesson; the copyright date; the file name; the library type; and options to select to see
more information on the lesson, the names of all the lesson authors, and general ordering information.
Press LAB to try the lesson.

When you select the further information option from the Lesson Information display, the system displays
the Lesson Description display (figure 4-20). The Lesson Description display contains a more detailed
description of the lesson, the estimated length of time it takes to complete the lesson, the type of
learning materials used, the type of audience the lesson is intended for and, where applicable, a list of

contents and a goal statement.

Further Information

Author Index

ESTIMATED LENGTH: 38 minutes
198% CAI

INTENDED AUDIENCE:

Beginning biology or medical science students,
accelerated high school or college level.

DESCRIPTION:

The lesson is divided into five parts. Each part presents a
Simulation, with which the student may interact, of the dif-
ferent parts of a neuron and how they function.

CONTENTS:

Neuron Structure - The five major parts of a neuron are pre-
sented and reviewed.

Action Potentials - The manner in which a neuron transmits
information is demonstrated.

Threshold Experiment - A simulation in which the student learns

about threshold values needed to fire a neuron.

The Synapses - A simulation of the means by which neurons com-

municate and the various means (i.e. neurological poisons)

Press NEXT to continue

BACK-go to previous page SHIFT-BACK go to options index

LAB to try this item

Figure 4-20. Lesson Description Display

The author index alphabetically lists all the authors of published PLATO-based courseware. From the
author index, you can see biographical information about authors and the titles of lessons written by
individual authors.

97405900 C

4-43

You can page through the author index in sequential alphabetical order or you can quickly move to any
part of the index by following the procedure described in the title index.

To see the titles of lessons written by an individual author, type the number in front of the author's name
and press NEXT. The system displays a list of titles of lessons written by that author. This title display
functions the same as the title index described previously.

To see biographical information about an author, type the number in front of the author's name and press
DATA.

Subject Index

The subject index alphabetically lists keywords which relate to the subjects of PLATO-based courseware.
Keywords summarize the subject of a lesson. To find materials in the subject index, think of a keyword
which summarizes the topic you are interested in. Type that keyword at the arrow on the Subject Index
display and press NEXT. You can also page through the index alphabetically to see the entire listing of
courseware subjects. Press NEXT to advance the listing one display at a time, or press BACK to reverse
the listing one display at a time. You can also jump to any part of the index (for example, from b to g) by
typing the letter(s) of the part of the index you want to see.

After you find a subject of interest, you can see the titles of lessons relating to that subject. Type the
number in front of the subject (keyword) and press NEXT. The system displays a list of titles relating to
that subject. This title display functions the same as the title index described previously.

File Name Index

The file name index alphabetically lists the file names of all the published PLATO-based courseware
materials.

From the file name index, you can see the title, author, and library type of the lesson. From this index,
you ean also request to see more information on a specific lesson by typing its file name.

To see more information on a specific lesson, type the number in front of the desired lesson and press
NEXT. To move to another part of the index, type the letter(s) of the alphabet from where you want the
listing to start and press NEXT.

USING THE ON-LINE AUTHOR LISTING

As an author or instructor, you can see a listing of all authors on the PLATO system who have voluntarily
ineluded their names in the listing and include your name in the listing if you choose. Information
available in the author listing includes: each author's full name, principal sign-on(s), office and home
phone numbers, mailing address, and authored subjects.

4-44 97405900 C

The following steps describe how to access and use the on-line listing of PLATO authors.
1. Do one of the following steps, depending upon your user type.
e From the Author Mode display (for authors), type authors and press NEXT.

e From the PLATO Facilities display (for instructors), select the Choose a Lesson to Study
option. Type authors at the What Lesson arrow and press NEXT.

2. The system displays the Directory of PLATO Authors display (figure 4-21). From this display,
you ean do any of the following steps.

Directory of PLATO Authors

minne system

Enter partial name
or name/group

>

for alphabetical listing
DATR for another system

LAB for lesson statistics
for information

Figure 4-21. Directory of PLATO Authors Display

97405900 C 4-45

Type either the name of the author on which you want to see information and press NEXT,
or type the letter of the alphabetical listing you want to see and press NEXT.

Press NEXT to see the Alphabetical List of Authors display (figure 4-22).

Press DATA to see the author listing for another PLATO system.

Press HELP for more information on how to use the On-Line Author Listing.

Alphabetical List of Authors

minne system

Full Name of Author

allen, michael w.

aspnes, gordon g

auld, warren e.

bartlett, judith a

battin, jim

berigan-pirro, denise

beringer, dennis b.

boggs, stephen h.

Enter partial name

or namegroup >

DATA for another system
HELP for information

Figure 4-22. Alphabetical List of Authors Display

You ean see an alphabetical listing of PLATO courseware subjects from the Directory of PLATO Authors
display. To see the subject listing, press SHIFT, type X, and press NEXT to see the listing starting at the
beginning of the alphabet. Press HELP for more information on how to use this feature.

4-46

97405900 C

To inelude your name in the author listing or to change your biographical information, press SHIFT-NEXT
from the Directory of PLATO Authors display. Type the letter in front of the information you want to
include or change, type the information at the arrow, and press NEXT. Press LAB to change information
not preceded by a letter.

USING THE PLATO USER LIST

The PLATO User List displays the names of all users who are currently signed on to the PLATO system
and have voluntarily included their name in the list.

The following steps describe how to see the list.
1. Do one of the following steps, depending upon your user type.

e From the Author Mode display (for authors), either press SHIFT and type U, or type users
and press NEXT.

@ From the PLATO Facilities display (for instructors), select the Interactive Communications
option and then select the See Users on the System option.

The system displays the Total Users display (figure 4-23).

97405900 C 4-47

4-48

2.

Total users on system = 65

Enter: a physical site number,
a group name,

a user’s name/group,

a site-station number.

Or press: NEXT (alone) to see all sites
DATA to see users at your logical site
LAB to display recordsvtalk flag information

SHIFT-DATA to see a list of operations personnel.

HELP available

Figure 4-23. Total Users Display

Choose any of the following options from the Total Users display.

Press HELP for more information.
Press LAB to set your user list options.
Press DATA to see the users at your logical site (those users with whom you share ECS/ESM).

Type a physical site number and press NEXT to see a list of the users at your physical site (a
grouping of terminals related by communications hardware).

Type a group name to see the users signed on from that group.

Type a site station number (the number assigned to a specific terminal) to see who is using it.

97405900 C

USER STATISTICS AND FLAG SETTINGS

The PLATO system records system use statistics for each author and instructor on the PLATO system.
The type of statistical information recorded includes the amount of time you have spent using the system
(both total time and individual sessions); information about your user record, such as your user type,
account name, logical site, and so on; and use of the central processing unit (CPU) and disk resources.

In addition to keeping statistics on system use, the PLATO system also keeps a record of your user flag
settings. These flag settings allow you to indicate to the system whether or not you want to use the
TERM-talk feature, appear in the User List, or receive TERM-ask calls. You can set or change your user
flags at any time.

To see information about your system use and flag settings, either press LAB from the Total Users display
(figure 4-23), or type I (hold the SHIFT key down and type i) from the Author Mode display or the PLATO
Facilities display (Choose a Lesson to Study option). Press HELP for more information about these
displays.

TIME-SAVING FEATURES

There are a number of features available on the PLATO system which serve as useful tools and
time-saving conveniences. Many of these features are called TERMS because the TERM key is used to
access them. Some TERMS tell the time of day, perform mathematical calculations, allow you to
comment on a lesson, or provide the correct spelling of a word. These TERMS are: TERM-time,
TERM-cale, TERM-comment, and TERM-spell, respectively. Refer to Checking the Time, Doing
Mathematical Calculations, and Commenting on Lessons in section 2 to learn how to use TERM-time,
TERM-cale, and TERM-comment. The following paragraphs describe how to use TERM-spell.

As an author, you ean code your lessons to prohibit
a specific TERM from working in your lesson.
Generally, this coding technique is used only when
use of the TERM would defeat the objectives of
the lesson. TERMS ean be prohibited from working
in either the entire lesson or certain parts of the
lesson.

Refer to AIDS for more information on how to
prohibit TERMS using the -termop- command.

97405900 C 4-49 .

TERM -spell

You can check the correct spelling of a word by using the TERM-spell feature. TERM-spell asks you to
type the word you need the correct spelling of, as close as possible to the correct spelling. The system
then shows you three words from its listing which are closest to that spelling. If none of these words
match, you can see a list of words which precede or follow those displayed. To use TERM-spell, follow
these steps.

1. Press TERM (hold SHIFT key down while pressing TERM/ANS key). The system displays What
term? >.

2. Type spell and press NEXT. The system responds with word: ( > ).

3. Type the first few letters of the word you want the correct spelling of and press NEXT. For
example, if you need the correct spelling of the word horizon, type hori and press NEXT. The
system displays three words from its listing which most closely match your entry.

4. If none of the words match, you can advance or reverse the listing to see words preceding or
following those displayed.

a. Press NEXT to advance the listing to the three words which alphabetically appear after the
words displayed. Continue pressing NEXT to advance the listing.

b. Press BACK to reverse the listing and see the three words which alphabetically appear
before the words displayed. Continue pressing BACK to reverse the listing.

5. Press SHIFT-BACK to return to your previous activity.

USING DOCUMENTATION FEATURES

The PLATO system provides three features which assist authors in writing and printing text materials and
creating displays. These features are the documentor file, the Graphics Utility for Interactive
Documentation Ease (GUIDE), and the print request feature. A documentor file allows users to write and
store text in the file. Documentor files, as well as other types of files, can be printed out on hard copy
using the print request feature. The GUIDE feature allows users to create graphic displays quickly and
easily without requiring the user to know the PLATO Author Language to do so.

The following paragraphs describe these features and how to use them.

USING DOCUMENTOR

Documentor is a file on the PLATO system which can be used as a tool for organizing, editing, and
presenting text-oriented material. As a text organizing and editing facility, it significantly reduces the
amount of time required to write, review, and polish a document. Documents that change frequently can
be kept current on a documentor file with the latest version available to anyone with access to the file.

As an author or instructor, your account director must create the documentor file for you (unless you have
account director authority). After the file is created, you can access documentor one of two ways,
depending upon your user type. [ As an author, you can access the file from the Author Mode display.
Type the name of the file and press NEXT. Type the security codeword (if required) and press NEXT. As
an instructor, select the Choose a Lesson to Study option from the PLATO Facilities display, type the
name of the documentor file, and press NEXT. ]

4-50 97405900 C

After you gain access to the documentor file, the system displays the Section Index display (figure 4-24).
From the Section Index display, you can create sections within your documentor file as well as read or
edit these sections. Press HELP for information on how to do this, as well as for a complete list and
description of options available in documentor.

Document dintro INSPECT ONLY

a4 Major Section/Subsection Indexes
b Format of Commands
Section Commands
Adding a Single Section ("a")
Adding Multiple Sections ("A")
Copying a Section ("c")
Deleting a Single Section ("d")
Deleting Multiple Sections ("D")
Editing a Section ("“e")
Listing Sections ("1")
Renumbering a Section ("n")

B Renumbering Multiple Secs. ("N")

1 Saving a Section ("5s")

2 Retitling a Section ("t")

3 Reading a Section ("r")

. .
meee DOONA VM 2 ON =~

4.
c
d
e
f
&
h
i
j
k
|
m
n
°

a>arADA A AA HA RAKED HAH

«x CONTINUED **

What command >
HELP available

Figure 4-24. Documentor Section Index Display

The editing directives used to edit a documentor file are similar to those associated with the standard
editor, with some exceptions. One major difference is the delete directive. In a documentor file, you can
delete lines by specifying the number of the line or lines you want deleted and press SHIFT-HELP. Lines
are not deleted from the top of the display as they are when using the standard editor. For example, if
lines 1 through 30 are visible on your screen and you want to delete lines 15 and 16, type d15-16 and press
SHIFT-HELP. For more information on documentor editing directives, press HELP while editing a section.

A complete description of documentor and how to use it is available in the PLATO lesson "dintro". Study
the lesson to learn more about documentor and its uses.

97405900 C 4-51 |

GRAPHICS UTILITY FOR INTERACTIVE DOCUMENTATION EASE (GUIDE)

As an author, you can create and edit graphic displays using the Graphics Utility for Interactive
Documentation Ease (GUIDE). Displays created using GUIDE can be linked together for interactive
documentation. The documentation produced using GUIDE functions in much the same manner as AIDS
does.

Using GUIDE, you ean easily create and edit graphic displays using an editor similar to ID/SD and the PLM
graphics editor. Displays ean be linked together for interactive use by specifying where to branch when a
key is pressed. GUIDE can also be used to quickly produce mock lesson displays or displays to be copied
for use as overhead or hardcopy graphic documents. No knowledge of the PLATO Author Language is
required to create complex graphics or to specify branching.

The GUIDE package provides both an editor and a driver. Execution of the editor provides an authoring
mode, while execution of the driver provides student mode.

The editor is similar to the TUTOR file graphics editor, but has many human interface and convenience
enhancements. Use of only the editor component is valuable for easy creation and modification of
displays to be copied for use as overhead transparencies or hardcopy graphic documents.

GUIDE provides complete freedom in the creation and formatting of all displays, including indexes. No
existing visual structures or formats are imposed on the user.

GUIDE's driver capabilities allow branching from one display to another. Jumpouts to PLATO lessons are
also possible. Special index displays allow presentation of menus when use of function keys (for example,
DATA, BACK, LAB) for branching is not sufficient.

The GUIDE driver provides a mechanism that automatically tracks each user's path through a display
sequence. With this mechanism, successive presses of BACK can automatically route the user through
displays previously seen — whatever branching options he/she may have chosen. As a result, review
sequences are very easy to provide. Display replotting is also an automatic feature.

GUIDE contains its own extensive HELP sequence, providing detailed how-to information within the
editor, similar to the detailed HELPs that authors find in the TUTOR file editor. Unlike the TUTOR file
editor, users will find more help and will need less prerequisite knowledge to begin using the GUIDE
system.

To use the GUIDE system to produce documentation interactively, a user needs a nameset with 20
character names and 64 word records. One can set lesson "guide" as the processor lesson of a given
nameset. Authors can also type guide on the Author Mode display, press DATA, and enter the name of a
nameset. Instructors can choose the Choose a Lesson to Study option on the PLATO Facilities display,
and, at the What lesson arrow, enter guide and press NEXT.

The GUIDE system allows both typeable and ACCOUNT and GROUP codewords. Setting lesson "guide" as
the processor lesson is recommended when using ACCOUNT and GROUP codewords; in such cases, Special
read/write access must be set.

For more information on GUIDE and to learn how to use it, study the PLATO on-line lesson "guideaids"

(type guideaids on the Author Mode display and press DATA or select the Choose a Lesson to Study option
on the PLATO Facilities display and type guideaids).

4-52 97405900 C

REQUESTING PRINTS

The print feature allows users to request prints of files from the central site (central prints) or to print a
file or make copies of displays on the PLATO terminal screen using a printer attached to a PLATO
terminal (local prints). In order to request central prints, a user's account must contract to use the print

feature. No contractual agreements are necessary for local prints.

The following steps describe how to request central prints and make local prints.

If your account is part of Control Data Services
Systems, the prints option on your author or
instructor record should be turned on and you
should contact your account director to request
prints.

To request a central print, do the following steps.
1. Do one of the following steps, depending upon your user type.

e@ From the Author Mode displey (for authors), either type prints and press DATA, or press
SHIFT and type R.

e From the PLATO Facilities display (for instructors), choose the Request a Print option by
typing the letter in front of the option.

The system displays the Print Requests display (figure 4-25).

97405900 C 4-53

4-54

--- PRINT REQUESTS ---

Choose an Option...

REQUEST a Printout

Check STATUS of Print Request

Check Printer STATUS

Print using a printer attached
to your terminal (local printer)

HELP is available

Figure 4-25. Print Requests Display

Press HELP for more information on how to request prints.

Choose the Request a Printout option by typing the letter in front of that option. The system
asks what file you want printed.

Type the name of the file you want printed and press NEXT. The system asks for the security
code of the file (if required).

Type the security code (if required) and press NEXT. The system displays a print selection
display. This display allows you to choose which sections of your file you want printed. You can
choose to print the entire file, including the title page, outline, and text (if printing a
documentor file), or choose any combination of these and selected sections of the file. Follow
the instructions on the display to choose the parts of the file you want printed. Press
SHIFT-NEXT when finished.

Type your name and mailing address. Press SHIFT-NEXT when finished. The system tells you
when the file will be printed.

97405900 C

To check the status of a print request made previously, do the following steps.
1. Choose the Check Status of Print Request option from the Print Request display.
2. Type the name of the file on which you want a status report and press NEXT.

3. The system displays a message stating whether or not the print has been made. The system also
provides the option to cancel a request if the file has not been printed as of that time.

You can also check the current availability of the line printer by choosing the Check Printer Status option
from the Print Request display.

To request a local print, do the following steps.
1. Do one of the following steps, depending upon your user type.
e From the Author Mode display (for authors), type print and press DATA.

e From the PLATO Facilities display (for instructors), select the Choose a Lesson to Study
option, type print, and press DATA.

Authors and instructors can also request local
prints by choosing the Print Using a Printer
Attached to your Terminal (Local Printer) option
from the Print Requests display (figure 4-25).

2. Do one of the following steps, depending upon the type of print you want.

e To print an entire file, choose the Print a File Using a Printer Attached to your Terminal
option by typing the letter in front of that option. The system displays several options.
Press HELP for an explanation of these options.

e@ To make copies of screen displays, choose the Make Copies of the Screen Using a Printer
Attached to your Terminal option by typing the letter in front of that option. The system
displays two screen copy options. Choose the option which corresponds to the type of
printer attached to your terminal by typing the number in front of the option.

97405900 C 4-55,

WRITING PLATO LESSONS

Authors write lessons for central PLATO system delivery using a computer language called the PLATO
Author Language (lessons written for Micro PLATO delivery use the Micro PLATO Language). The lessons
for both central and Micro PLATO delivery are written and stored in files on the PLATO system. A file
contains a collection of data which is stored in a reserved space in the PLATO system. File lengths vary
according to the amount of computer memory space reserved for each file.

The PLATO system has several types of files for authors to use, depending upon their needs. The type of
file used to write lessons is a TUTOR file. A TUTOR file is composed of parts and blocks. Parts and
blocks are subdivisions within a TUTOR file. Each file contains at least one part. Each part contains up
to seven blocks. Each block can store 320 computer words. The number of parts in a file determines the
length of the file. A one-part file contains 7 blocks, a two-part file contains 14 blocks, and so on.
TUTOR files ean contain a maximum of 10 parts or 70 blocks.

The block is the part of the file in which authors write the code for the lesson. There are several types of
blocks authors can use for their lessons. The most common block type used for writing lessons is the
TUTOR block. The TUTOR block holds the lesson code. Other block types, which the author uses in
conjunction with the TUTOR block in writing a lesson, perform specifie functions. Some blocks allow
authors to create graphic characters (charset blocks, lineset blocks) and store documentation for the
lesson (text block). Some blocks have time-saving functions and others have creative functions. These
bloek types are described in Using Other Block Types later in this section.

When a file is created, the person creating the file determines the length of the file, or how many parts to
assign to the file. If you, as an author, have account director capabilities, you can create your own
TUTOR file and determine its length. If you do not have account director capabilities, your account
director can create a TUTOR file for you. Remember to tell your account director how many parts to
include in the file. Refer to Creating Files in section 5 to learn how to create your own files if you have
account director capabilities.

USING THE TUTOR LESSON FILE

After you or your account director create a TUTOR file for you to use, there are some things you should
do to protect the file. These procedures are similar to those an instructor does when using a group file.
They involve registering information about yourself and your lesson and assigning security codewords to
the file. The following paragraphs describe how to do these procedures.

Registering Author and Lesson Information

Each TUTOR file has a directory which allows you to register and store information about you and your
lesson. This display is the Author Information display (figure 4-26). The first time you access your
TUTOR file by typing the name of your TUTOR lesson file from the Author Mode display and pressing
NEXT, the system displays the Author Information display. Enter information for all the entries on this
display, particularly the one-line description. If you do not, the system returns you to this display each
time you access the file until all information is entered. Completing all entries on this display also
prevents your file from being inadvertently destroyed by your account director during a routine file
cleanup. If you are uncertain about the kind of information to include in the one-line description of the
Lesson Information section, simply type something which indicates the purpose of the file. To enter
information on the Author Information display, type the number in front of the entry you want to access,
type your information, and press NEXT.

After information has been entered on the Author Information display, you will no longer be brought
directly to this display each time you enter the file. From thereon, the Block Listing display (figure 4-27)
appears each time you access the file. To reach the Author Information display from the Block Listing
display, press DATA or choose the *directory option on the display.

4-56 97405900 C

97405900 C

Lesson name ---- english

Account --------

Press the associated number to change an entry.

Ruthor Information:

1. Name ----------9---- jane doe
2. Dept. /Affiliation -- english
3. Telephone number --- 221-3221

Lesson Information:

4. Subject Matter ----- literature
5. Intended Audience -- high school seniors
6. One line description ----------

> examines literature during the early 1888's

Figure 4-26. Author Information Display

LESSON--engl ish HELP available

BLOCK NAME

rdirectory
grammar
spelling
sentences
letters

Figure 4-27. Block Listing Display

4-57

Assigning Security Codewords

You can protect your file from access by unauthorized users by assigning security codewords to the file.
Security codewords control which users can see or change the contents of the file. As an author, it is your
responsibility to create and set codewords to your files. Codewords are similar to passwords in that they
control which users can see or change the lesson code in the TUTOR file. Codewords can be set to allow
specific users or specific user types access to the file. For example, you can set the codewords so only
you can see or change the file, so only specific users you choose can see or change the file, or so only
authors in your group or account can see or change the file. Examples of the different types of security
codes you can set are:

Typed code Requires all users to type the file's security codeword to see and/or change
the contents of the file.

GROUP code Allows all authors within the group to see and/or change the contents of the
file without typing the codeword first.

ACCOUNT code Allows all authors whose groups are listed within an account to see and/or
change the contents of the file without typing the codeword first.

Unmatchable code No security code of any sort can match this; no access by any other file is
possible. This codeword is automatically assigned to newly created files for
all codes other than the change and inspect codewords. These security
codes are assigned by the system to prevent accidental accesses to
databases before authors have had an opportunity to assign appropriate
security controls. In addition, an author or account director can assign
unmatchable codes to any codeword of any file; this would prevent any
author from making any changes to that file. If used for the common code
of a lesson, that lesson and only that lesson can connect to commons and
leslists stored within that file.

It is important to be creative when assigning codewords. If you use a typed code, be sure it is something
no one can guess. Do not use obvious codes like your spouse's name; the name of your group, account, or
file; your pet's name; your password; your telephone number; a period; a, b, ¢, and so on. Choose
something with which only you can identify. Change your typed codewords frequently to prevent the
possibility of someone guessing them. Examples of good codewords are misspelled words of at least seven
characters, or words which have numbers inserted in them.

If you use a GROUP or ACCOUNT code, access to the file is limited only to people in your group or
account. Although this limits the number of people who can access your file, it makes your sign-on
password extremely important because access to the file is based on who you are rather than what you
know. Therefore, it is very important to protect your sign-on password by making sure it is a creative
password which no one can guess. Remember, do not use obvious passwords which are familiar to anyone,
and change your password frequently.

Each TUTOR file contains a display which allows you to register and store security information. This
display is the Security Codewords display (figure 4-28). When a TUTOR file is created, the Security
Codewords display may be blank. Account owners may choose to assign default codewords in their
accounts. Default codewords are placed automatically on all files created within an account. Usually,
default codewords only provide a limited security level because they are placed on all files. The default
codewords of all newly created files should be evaluated to assure each file has a security level matching
its function. You should set or evaluate the codewords on the Security Codewords display as soon as the
file is created to prevent other users from seeing or changing the file. The following steps describe how
to access the Security Codewords display and assign codewords in the TUTOR file.

4-58 97405900 C

Lesson name ---- english

Account --------

Press the associated number to change an entry.

SECURITY CODES:

1. To change lesson ----  ##KkeKKKRKH
2. To inspect lesson ---  ##kekeE EH
3. To access common ---- No match permitted
4. To -use- lesson ----- No match permitted
5. To -jumpout- to ----- No match permitted
6. To -attach- a file -- No match permitted

Access to file by system personnel :
7. System Access -------

Figure 4-28. Security Codewords Display

1. From the Author Mode display, type the name of your TUTOR file and press NEXT. The system
displays the Block Listing display (figure 4-27). This means some author information, and
possibly account or group security codes, have already been entered in the file's directory.

a. Press DATA to see directory information.

b. Choose the Security Codewords option. If no security codes exist (Blank - open to all),
immediately assign codewords. If general security codes, such as ACCOUNT or GROUP
codes, have been assigned, consider their appropriateness for the file. Ask yourself whether
or not such a large group of authors should have access to change or inspect the file.

To assign security codewords, do any of the following.

@ The To Change Lesson option allows you to determine which users or group of users can access
and change (edit) your lesson code. To choose this option, type the number in front of it. Do one
of the following steps, depending upon the type of user access you want to allow.

- To allow access to yourself only or to a small group of people who know the codeword, type
a codeword and press NEXT. A random number of X's appear to the right of the arrow as
you type. The system asks you to retype the codeword to verify it and help you remember
it. Press NEXT.

- To allow access to all authors in your group, press LAB. The system responds by displaying a
GROUP option and an ACCOUNT option. Type the number in front of the GROUP option.

97405900 C 4-59

4-60

- To allow access to all authors in your account, press LAB. The system responds by
displaying a GROUP option and an ACCOUNT option. Type the number in front of the
ACCOUNT option.

- To restrict access to all users, press LAB. Type the number in front of the unmatchable
code option.

The To Inspect Lesson option allows you to determine which users or group of users can read but
not change your lesson code. To choose this option, type the number in front of the option.
Follow the instructions in step 2 for assigning codewords.

Choose a different codeword than the change

codeword if you want users to see but not change
your lesson code.

The To Access Common option allows you to access any common block of this or another TUTOR
file and use the data contained in that common block in your lesson. To do this, the security
codewords for your access common option and the access common option of the file you want to
access must match. To choose this option, type the number in front of it. Follow the
instructions in step 2 for assigning codewords. The option also controls access to leslist blocks.

The To -use- Lesson option allows you to use the code contained in another TUTOR lesson file in
your lesson. To do this, the -use- lesson codewords for the files must match. To choose this
option, type the number in front of it. Follow the instructions in step 2 for assigning codewords.

The To -jumpout- To option allows you to control which lessons can be reached from your lesson
and also control whether other authors can ~jumpout- to your lesson. To do this, the -jumpout-
codewords for both files must match. To choose this option, type the number in front of it.
Follow the instructions in step 2 for assigning codewords.

The To -attach- to a File option allows you to manipulate or change information in one file by
executing code residing in another file. To do this, the codewords for the files must match. To
choose this option, type the number in front of the option. Follow the instructions in step 2 for
assigning user access.

The System Access option allows you to choose whether or not to give systems personnel access
to your TUTOR file. Systems personnel (users responsible for maintaining the PLATO system)
occasionally need access to files to check for errors if hardware problems occur on the system.
If you choose to allow systems personnel access to your file, authorized users can access the file
in inspect mode, without typing a security code. Type the number in front of this option to
change the option.

97405900 C

EDITING TUTOR FILES

Before you begin entering lesson code into your TUTOR file, you should know how to create a block in
which to write and store your lesson code, how to use the PLATO system editor to enter and edit code,
and how to receive help while editing your lesson. The following sections describe how to create a
TUTOR block for writing and storing code, how to use the PLATO system editor, and how to get editing

help.

Creating a TUTOR Block

The TUTOR bloek is the part of the TUTOR file in which you write lesson code. Each part of a TUTOR
file contains space for you to create seven blocks. Each TUTOR block stores 320 computer words or
approximately 50 lines of code.

The following steps describe how to create a TUTOR block.

1. From the Author Mode display, type the name of your TUTOR file and press NEXT. The system
asks for the file security code (if you assigned typed inspect or change codes to the file).

2. Type the security codeword (if required) and press NEXT. The system displays the Block Listing
display (figure 4-27). Notice that the Block Listing display lists only one block (a directory).
The block directory contains author and lesson information, file security information, and
associated files and editing specifications information. This is the only block which is not used
for writing lesson code. The block directory is reached by typing a or pressing DATA from the
Block Listing display.

3. Create a block for writing and storing lesson code by typing the capital letter of the block which
precedes the block you want to create. For example, to create your first block, type A since
block a precedes the block you want to create. The system displays the Block Creation Options
display (figure 4-29).

97405900 C 4-61

for a normal TUTOR block
for copy-a-bl ock

a common block

a charset block

a micro block

a leslist block

a vocabulary block
a lineset block

@ listing block

a text block

1
2
3
4
5
6
?
8

Figure 4-29. Block Creation Options Display

4, The Block Creation Options display gives you several options to choose from for block creation.
Create a normal TUTOR block for entering lesson code by pressing NEXT. The system asks you
to name the block you created.

5. Type a name for your block and press NEXT. If you are unsure of what to title your block, use
test, workspace, or another name which indicates the purpose or subject of the block. The
system displays your TUTOR Block display.

6. Do one of the following steps.

e If you are familiar with the PLATO Author Language and editing directives, and are
prepared to enter lesson code, enter your code using the correct editing directives.

e If you are unfamiliar with the PLATO Author Language and editing directives, or are not
prepared to enter lesson code at this time, type i and press NEXT. The system is now in
insert mode. When the system is in insert mode, it is ready for you to enter code. Type an
asterisk and press BACK.

You must type something in a new block which the
system can store or else the block is deleted from
the Block Listing display.

The system displays the Block Listing display. Two asterisks indicate the last block you edited before
returning to the Block Listing display.

4-62 97405900 C

Using the PLATO System Editor

The PLATO system editor allows you to insert and change the code you write in the blocks of your file. It
allows you to enter, revise, read, and transfer your lesson code. The editor also tells you how much space
is available in each block for you to enter code.

Special instructions are required to use the PLATO system editor. These instructions are called editing
directives. These editing directives are similar to those authors use for writing notes. The following
paragraphs describe the basic editing directives you should know before entering code in your TUTOR
block. Refer to Requesting Editing Help later in this section for information on how to see detailed
information on all the editing directives.

Inserting Code

The insert directive allows you to type lesson eode into the TUTOR block. To use the insert directive,
type i and the number of the line you want the new material to follow. Press NEXT. For example, to
insert code after line 3, type i3 and press NEXT. To add code before line 1, type i0 and press NEXT. If
the block is empty and this is the first time you are inserting any code, type i and press NEXT. You do
not need to type a line number. When you are in insert mode, the system displays INSERT MODE at the
top of the screen. The upper right corner of the display tells how much space is available for you to insert
lines of code. After you finish inserting code, press BACK.

Replacing Code

The replace directive allows you to replace or change lines of code already entered in the block. To use
the replace directive, type r and the number of the line you want changed. For example, to correct an
error in line 3, type r3 and press NEXT. The system is now in replace mode. In replace mode, the system
displays the line you want to replace with an arrow directly under it. You can either retype the entire
line correctly or use the COPY and EDIT keys to edit the line. (Refer to appendix A to learn how to use
the COPY and EDIT keys.) After you complete your corrections, press BACK.

Positioning Code on the Screen Display

The PLATO system displays 31 lines of text at one time. Since most lessons are longer than 31 lines, the
system allows you to page through your code so you can see different parts of it. Each TUTOR block can
store approximately 50 to 60 lines of code. You can think of these lines as being on a scroll, the scroll
being viewed through a 31-line screen. The scroll can be rolled forward to show lines toward the end of
the scroll, or it ean be rolled backward to show lines toward the beginning of the scroll. An example of
use of the forward and backward editing directives is in figure 4-30. There are several ways to move code
on the screen. Two of the easiest ways are with the forward and backward editing directives.

97405900 C 4-63.

4-64

SAME LINES

Figure 4-30. Example of Forward and Backward

BLOCK i-b = one SPACE = 158

1 uni raven!
2 at sis
3 write Once upon a midnight dreary,
4 while I pondered, weak and weary,
5 Over many a quaint and curious volume
6 of forgotten lore--
? While I nodded, nearly nappirg,
8 suddenly there came a tapping,
9 fe of someone gently rapping,
is repping at my chamber door.
ro "'Tie some visitor,” I muttered,
12 “tapping at my chamber door--
13 Only this and nothing more.”
14
1s Ah, distinctly I remember
16 it was in the bleak December;
7 Pind each separate dying eber
16 wrought its ghost upon the floor.
19 Eagerly I wished the morrow; --
2a vainly I had sought to borrow
at From my books surcease of sorrow--
22 sorrow for the lost Lenore--
23 Fore the rare and radiant maiden
24 whom the angels name Lenore--
25 Nameless here forevermore. SAME LINE
26
27 at 3a48
28 write Press NEXT
29 unit ravenz
3g at sis
31 write find the silken, sad, uncertain rustling
FORWARD 12
BLOCK 1-b = one SPACE = iss
T Only this end nothing more.”
2
3 Ah, distinctly I remember
4 it was in the bleak December;
s find each separate dying eber
6 wrought its ghoet upon the floor.
7 Eagerly I wished the morrow; --
6 vainly IT had sought to borrow
9 From my books surcease of sorrow--
is sorrow for the lost Lenore--
it Fore the rare and radiant maiden
12 whom the angels name Lenore--
13 Nemeless here forevermore.
14
15 at 3948
16 write Presse NEXT
17 unit raven2
18 at sis
19 write fnd the silken, sad, uncertain rust] ing
20 of each purple curtain
a Thrilled me--filled me with SAME LINE
22 fantastic terrors never felt before;
23 So that now, to still the beating of my heart,
24 I stood repeating,
25 “‘Tis some visitor entreating entrance
26 at my chamber door--
27 Some late visitor entreating entrance
26 at my chamber door; --
29 This is it and nothing more.”
KJ
3h Presently my soul grew stronger;
1
"2
3
4 rapping at my chamber door.
5 *'Tis some visitor,” I muttered,
6 “tapping at my chember door--
7? Only this and nothing more.“
6
9 Ah, distinctly I resember
is it wee in the bleak Decesber;
11 find each separate dying eber
12 wrought its ghost upon the floor.
13 Eagerly I wished the morrow; --
14 vainly I had sought to borrow
15 From my books surcease of sorrow--
16 sorrow for the loet Lenore--
7 Fore the rare and radiant maiden
18 whom the angels name Lenore--
19 Nemelees here forevermore.
an
21 at 3548
22 write Presse NEXT
23) unit raven2
24 at sig
25 write find the silken, sad, uncertain rust] ing
26 of each purple curtain
2? Thrilled me--filled me with
28 fantastic terrors never felt before;
29 So that

<!-- Book text truncated by scrapem max_book_chars. -->

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
