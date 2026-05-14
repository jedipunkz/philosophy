---
source: "https://arxiv.org/abs/1910.02400v1"
title: "A Linear LMP Model for Active and Reactive Power with Power Loss"
author: "Yanghao Yu, Qingchun Hou, Yi Ge, Guojing Liu, Ning Zhang"
year: "2019"
publication: "arXiv preprint / eess.SY"
download: "https://arxiv.org/pdf/1910.02400v1"
pdf: "https://arxiv.org/pdf/1910.02400v1"
captured_at: "2026-05-09T12:16:43Z"
updated_at: "2026-05-09T12:16:43Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche will to power"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# A Linear LMP Model for Active and Reactive Power with Power Loss

- 著者: Yanghao Yu, Qingchun Hou, Yi Ge, Guojing Liu, Ning Zhang
- 年: 2019
- 掲載情報: arXiv preprint / eess.SY
- 情報源: [arxiv](https://arxiv.org/abs/1910.02400v1)
- ダウンロード: https://arxiv.org/pdf/1910.02400v1
- PDF: https://arxiv.org/pdf/1910.02400v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Pricing the reactive power is more necessary than ever before because of the increasing challenge of renewable energy integration on reactive power balance and voltage control. However, reactive power price is hard to be efficiently calculated because of the non-linear nature of optimal AC power flow equation. This paper proposes a linear model to calculate active and reactive power LMP simultaneously considering power loss. Firstly, a linearized AC power flow equation is proposed based on an augmented Generation Shift Distribution Factors (GSDF) matrix. Secondly, a linearized LMP model is derived using GSDF and loss factors. The formulation of LMP is further decomposed into four components: energy, congestion, voltage limitation and power loss. Finally, an iterate algorithm is proposed for calculating LMP with the proposed model. The performance of the proposed model is validated by the IEEE-118 bus system.

## PDF Text

1
A Linear LMP Model for Active and Reactive Power
with Power Loss

Yanghao Yu, Qingchun Hou,
Student Member, IEEE
, Yi Ge, Guojing Liu, and Ning Zhang,
Senior Member, IEEE




Abstract
— Pricing the reactive power is more necessary than
ever before because of the increasing challenge of renewable
energy integration on reactive power balance and voltage control.
However, reactive power price is hard to be efficiently calculated
because of the non-linear nature of optimal AC power flow
equation. This paper proposes a linear model to calculate active
and reactive power LMP simultaneously considering power loss.
Firstly, a linearized AC power flow equation is proposed based on
an augmented Generation Shift Distribution Factors (GSDF)
matrix. Secondly, a linearized LMP model is derived using GSDF
and loss factors. The formulation of LMP is further decomposed
into four components: energy, congestion, voltage limitation and
power loss. Finally, an iterate algorithm is proposed for
calculating LMP with the proposed model. The performance of
the proposed model is validated by the IEEE-118 bus system.

Index Terms
— Reactive power pricing, optimal power flow,
Generation Shift Distribution Factor, LMP, power loss.
N
OMENCLATURE

P,Q

Vectors of bus injected active and reactive power

, ii
PQ

Injected active and reactive power at bus i

j
Y=G+B

Admittance matrix of the power system
j

Y=G+B

Admittance matrix without shunt elements
ijijij
YGjB


Element in the
i th row and
j th column of
Y

ijij
YGjB



Element in the
i th row and
j th column of

Y

V,
θ

Vectors of Magnitudes and phase angles of bus
voltages

, ii
V


Voltage magnitude and phase angle at bus i

,
MN

The number of branches and buses
ijijij ygjb


Admittance of branch
(,)
ij

iiiiii ygjb


Shunt admittance at bus
i

,, lll mmm
PQI

Power flow and current through branch
m

,
GD
ii
PP

The
active power generation and active load at bus
i

,
GD
ii
QQ

The
reactive power generation and reactive load at
bus
i

This work was supported by the Major Smart Grid Joint Project of National
Natural Science Foundation of China and State Grid (No. U1766212) and
Scientific & technical project of State Grid "Research and application of
large-scale energy storage planning for receiving-end power system" (No.
5102-201918309A-0-0-00).
Y. Yu, Q. Hou, N. Zhang are with the State Key Lab of Power Systems,
Department of Electrical Engineering, Tsinghua University, Beijing 100084,
China

(e-mall:
ningzhang@tsinghua.edu.cn
).
Y. Ge, G. Liu are with Economics and Technological Research Institute,
State Grid Jiangsu Electric Power Co., Ltd., Nanjing, China.
,
PQ
ii
FF

The fictional nodal demand at bus
i

,
LossLoss
PQ

Total power loss of the system
,
PQ
ii
LFLF

The active and reactive power loss factor at bus
i

,
PQ
ii
DFDF

The active and
reactive power delivery factor at
bus
i

I.

I
NTRODUCTION

OCATIONAL Marginal Price (LMP) is defined as the
marginal expenses of both generators and transmission
system to supply the increment of active power at each bus. A
power market using LMP can provide the right price signal to
relieve the congestion of power system. Therefore, LMP has
been widely adopted in Independent System Operators (ISOs),
such as PJM interconnection, the New York market, and the
New England ISO [1].
The LMP can be divided into active power LMP (ALMP)
and reactive power LMP (RLMP). The later denotes marginal
cost to supply the increment of reactive power demand at each
bus. Currently the LMP adopted by all of the power market is
ALMP only. Most reactive power sources like static VAR
compensators (SVCs) and switched capacitors (SCs) are not
priced by RLMP but compensated as ancillary service in power
market. Nowadays, increasing renewable energy penetration
brings great challenge to power system voltage security. The
adequate reactive power supply may become critical. Since the
reactive power cannot be transmitted by long distance and is
usually locally balanced, the need of reactive power for
different node would be distinct. The value of providing
reactive power would be different even there is no congestions
in the power system. Therefore, introducing RLMP in the
market would provide significant price signal for both
short-term operation and long-term planning.
The value of LMP can be decomposed into several
components. Reference [1] proposes the DCOPF model with
loss to decompose LMP into the costs of marginal energy,
marginal loss and congestion. Reference [2] introduces the
concept of LMP to a distribution system. The distribution LMP
(DLMP) is decomposed into marginal energy cost, loss cost
and congestion components. References [3], [4] further
analyzes both active and reactive power DLMP with
consideration of congestion and voltage support. The DLMP is
divided into five components: marginal costs for active power,
reactive power, congestion, voltage support, and power loss.
The value of different LMP components together reveal the
economic and operation conditions of power system, which
provide price signal to influence the behavior of participants to
mitigate transmission congestion and stimulate reactive power
L

2
supply for voltage support.
LMP can be derived from optimal power flow (OPF) [1]. In
practical application of LMP, AC optimal power flow (ACOPF)
and DC optimal power flow (DCOPF) are two classic models.
ACOPF can simultaneously calculate ALMP and RLMP
accurately. But it is a non-convex problem which may be
computationally intractable even for smaller systems [5].
Additionally, LMP derived from ACOPF are not analytical so
that it cannot be decomposed into different components and is
lack of transparency [6]. Comparatively, DCOPF model is
more concise and efficient, which makes it widely utilized by
LMP-based markets. However, since it neglects the voltage
fluctuation and power loss, it has a relatively low accuracy in
LMP calculation. In addition, DCOPF cannot calculate RLMP
since it ignores the reactive power balance and reactive power
flow.
Recent researches have been conducted to improve the LMP
calculation. Reference [7] uses energy reference bus
independent model to obtain a more accurate calculation on
energy, congestion and loss components of LMP. In Reference
[8], an iterative DCOPF algorithm is proposed to address the
nonlinear marginal loss. Reference [9] proposes a linear
loss-embedded model to calculate LMP with loss components
by dealing with marginal generators under lossless DCOPF
solution. Reference [10] proposes a linear-constrained and
convergence-guaranteed OPF model that explicitly considers
the constraints of reactive power balance and bus voltage.
Reference [11] proposes a linearized power flow model with
decoupling of voltage magnitude and phase angle to achieve
accurate estimation of voltage magnitude. Based on the
linearized power flow model, Reference [6] further proposes
the decoupled OPF model using Generation Shift Distribution
Factors (GSDF) and further applies it to jointly calculate
ALMP and RLMP with respect to voltage constraints, but the
impact of network loss is neglected in this model.
This paper proposes a linearized iterative model to calculate
both ALMP and RLMP considering power loss. A linearized
AC power flow equation is proposed based on the augmented
GSDF matrix. The active and reactive power loss are linearized
by introducing loss factor (LF). An iterate algorithm is
proposed to calculate power loss and the LMP. Finally, the
accuracy and performance of the proposed model is validated
on an IEEE-118 bus case.
The rest of this paper is organized as follows: Section II
constructs the linearized LMP model. Section III validates the
accuracy of the ALMP and RLMP calculation of the model
through the IEEE 118-bus case on MATPOWER. Finally,
Section IV concludes the paper.
II.

L
INEARIZED
LMP

M
ODEL
W
ITH
P
OWER
L
OSS

This Section builds the LMP model based on GSDF. The
modeling is divided into five subproblems: 1) Introducing
GSDF-based linear AC power flow model; 2) Linearizing the
power losses; 3) Building GSDF-based LMP model; 4)
Deriving ALMP and RLMP based on the LMP model; 5)
Proposing iterate algorithm to solve LMP. This Section will
present the technical details of each subproblem.
��

�����������������������
A linearized power flow equation with respect to voltage
magnitude and phase angle is proposed by reference [11] as
follows.

-
�
��������
��������
�
��������
�����
�
���
�����

(1)
The matrix
�
in Eq. (1) is a
22
��
�
matrix determined
solely by the element of admittance matrix.
Based on Eq. (1), the active power balance equation is
derived as follows:

111111111
11
0
���������
�������������
���������
��
�����
��
�������
���
qq
=========
==
����
��
=-=-
����
����
=��
���������
��
(2)
The two approximation steps in Eq. (2) are based on the
assumptions that
1
�
�
�
and
0
��
�
�
. Similarly, we have the
reactive power balance equation as follows:

111
���
������
���
����
===
=-�-
���

(3)

According to Eq. (2), the equations Eq. (1) is not
independent. In other words, matrix
�
in (1) is close to
singular. So we remove the row and column of the reference
bus in
�
and inverse the new
(
)
(
)
2121
��
-�-
matrix. We
then fill the row and column of the reference bus in the inversed
matrix with zeros. Hence, we have a new
22
��
�
matrix
�

satisfying:

����
=
����
����
�
�
�
��
(4)
Unfolding the matrix form of Eq. (4), the voltage equation is
acquired as follows:

,,
11
��
����������
��
�����
+++
==
=+
��

(5)

According to Eq. (4), the power flow of branches can be
linearly expressed by the bus power injection both of active and
reactive parts:

(
)
(
)
()()
()
()()
()
,,,,
1
,,,,
1

=
=

����������
�
���������������
�
�
�������������������
�
������
�������
�������
qq
++
=
++++++
=
�---
---
+---
�
�

(6)

(
)
(
)
()()
()
()()
()
,,,,
1
,,,,
1
=


����������
�
���������������
�
�
�������������������
�
������
�������
�������
qq
++
=
++++++
=
�----
=----
+----
�
�

(7)

The coefficients of
�
�
and
�
�
in Eqs. (6)-(7) forms the
GSDF matrix between the branch power flow and power
injection:

3








,,,,
,,,,
,,,,
,,,,
��
m���������������
��
m�������������������
��
m���������������
��
m�������������������
������b��
������b��
������b��
������b��













(8)

Finally, the active and reactive power flow equations (Eqs.
(6)-(7)) can be expressed using GSDF:



11
11
11
11
,for=1,2
,for=1,2
��
�����
mm��m��
��
��
��������
m���m���
��
��
�����
mm��m��
��
��
��������
m���m���
��
���������
����������m�
���������
����������m�




















�
�

(9)
Different from classic DC power flow model, the proposed
model considers both active and reactive power, which provide
foundations for deriving ALMP and RLMP.
��

��ns�der�n� ���er ��ss �n �����based ���er ����
���at��ns
It should be noted that in the derivation of last subsection we
do not consider power losses. In other words, the total active
power generation and reactive power generation in Eq. (2)
equals the active load and reactive load respectively (plus the
load on the shunt branch). Such lossless consumption does not
coincide with practical power system. Therefore, we need to
consider the power losses in power balance equations and
allocate the power losses to each node as virtual load.
According to Joule’s Law, both active and reactive power
losses, can be calculated using branch power flow (assuming
that
1
�
�

):




22
11
2
111
22
11
2
111
��
��
��ssmmmm mm
���
��������
m���m���m m��
��
��
��ssmmmm mm
���
��������
m���m���m m�m
�I���
�����������
�I���
�����������






















(10)

In order to consider power losses, the concept of loss factor
(LF), delivery factor (DF) and fictional nodal demand (FND) [8]
are introduced in this paper. LF indicates the marginal power
losses with respect to an incremental power injection on a bus,
while DF indicate the active and reactive power that is
equivalently injected to the grid. LF and DF for active and
reactive power can be stated as Eq. (11).

1
1
=2,1for=1,2
=2,1for=1,2
�
������
��ss
�mm�m��
�
m
�
�
������
��ss
�mm�m��
�
m
�
�
�������������
�
�
�������������
�














�
�

(11)

It should be noted that
�
�
��
and
�
�
��
may be negative for
some node, denoting that the power injection would reduce the
power losses. We introduce the DF into the power balance Eqs.
(2)-(3) so that the power losses can be explicitly formulated in
the equation:



1
11
0
�
���
�����ss
�
��
���
�����ss��
��
�����
�����b






(12)
FND is introduced to evenly allocate the loss of each branch
to its two end buses, mathematically:


22
11
,,for=1,2
22
����
�mm�mm
�m�m
��������



�

(13)

Each bus takes the
�
�
�
and
�
�
�
on their connected branches
as a virtual injection. Therefore, the power injection at each bus
is restated as:

for=1,2
for=1,2
���
����
���
����
������
������


�
�

(14)
By introducing the FND in each bus, the total active power
generation no longer equals to the active load since the
generation need to balance the extra virtual injection. It
suggests that both the active and reactive power losses are
considered in the proposed GSDF-based power flow equations.
��

��� ��de� �ased �n ����
According to the derivations above, especially with Eqs. (5),
(9), (12), and (14), we draw the GSDF-based LMP model as
shown in Eq. (15).
,,
���
abc
represent the cost coefficients of
generator at bus
�
, the cost functions of active and reactive
power generation are both quadratic while reactive function
does not have primary term;
,,,
����
��

denote the Lagrange
multiplier of the constraints of active power balance, reactive
power balance, power flow congestion and voltage respectively;
,
��
��
����
denote the DF at bus
�
;
,
��ss��ss
��
are the total
system active and reactive power loss;
,
��
��
��
are active and
reactive FND at bus
�
;
lim
�
�
is the active power limits of branch
�
;
minmax
,
��
��
denote the minimum and maximum voltage at
bus
�
respectively;
minmaxminmax
,,,
����
����
����
denote the active
and reactive power generation limits at bus
�
respectively.

4






22
1
1
11
11
,,
1
..0
,for=1,2
�
���
������
�
�
���
������ss�
�
��
���
������ss���
��
��
�����������
mm����m����mm
��
�
����
���������
�
��na�b�c�
st������
������b
��������������m��
������




















�
�
��
�

1
limlim minmax minmax minmax
,for=1,2
,for=1,2
,for=1,2
,for1,2
,for1,2
�
���
�������
�
�
mmm
���
���
���
���
���
������
���m�
�����
�����
�����









�
�
�
�
�

(15)

��

A��� and ���� �ec�m��s�t��n
When the optimal solution of the model is reached, LMP can
be calculated by the Lagrange Function:


22
1
11
�
���
����������
�
��
�
mm��
m�
a�b�c���
��











(16)
ALMP is defined as the marginal cost of an incremental
change of active power demand at a bus. ALMP at bus
�
can be
calculated as the first-order partial derivative of the Lagrange
Function with respect to the active power load at bus
�
:

11
,
11
,
11
=
=
=
=
�
�
�
�
��
�
�m
�m�
���
m�
���
��
���
��mm�����
m�
��
���
�mm�������
m�
A���
�
�
��
���
������
������






















(17)
Similarly, RLMP is defined as marginal cost of an
incremental change of reactive power demand at a bus. The
calculation of RLMP at bus
�
is:

11
,
11
,
11
�
�
�
�
��
��
m
�m�
���
m�
���
��
���
��mm������
m�
��
���
�mm��������
m�
����
�
��
�
���
������
������























(18)
Eqs (17) and (18) derives the analytical form different
components of ALMP and RLMP. The first term of LMP is the
energy component derived from the power balance constraints.
The second term is the congestion component which is related
to the power flow equations. The third term is the voltage
constraint component. The last term is loss component. If loss
is not considered in a model, LF is zero. When the constraints
on power flow or voltage magnitudes are active, their
corresponding components of LMP are non-zero.
��

�����n� A���r�t�m
An iteration algorithm is proposed to solve the LMP model
with power loss, as shown in Fig. 1. First, we solve the LMP
model without power loss. Then the power loss is calculated
using Eqs. (10)-(11) and is allocated using Eqs. (14). The LMP
model is later updated with the calculated DF and FND. After
one iteration, we solve the updated LMP model and re-calculate
the power losses. When the difference of power losses in two
iterations is smaller than a certain threshold, iteration stops and
a final LMP is obtained from the newest LMP model.

Figure 1 Flowchart of the iteration algorithm
III.
�
C
ASE
S
TUDY

In this section, the GSDF-based LMP model is applied to an
IEEE 118-bus case on MATPOWER 6.0 to validate its

5
accuracy in calculating ALMP and RLMP. The model is
implemented in Matlab R2016a with the solver Gurobi8.0. The
simulation platform is a personal computer with an Intel Core
i7 CPU @2.50GHz and 8GB RAM.
The IEEE 118-bus test case is from MATPOWER 6.0
toolbox of MATLAB [12]. Meanwhile, the original case does
not contain limits on branch power flow. The branch MVA
rating data are taken from Index of Data Illinois Institute of
Technology [13]. Moreover, the case does not provide cost for
reactive power. We add the reactive cost function of generators
and let
i c
in Eq. (15) be the same as
i b
.
The LMP results of our model (denoted as LMP-L) are
compared with the LMP model without loss (denoted as
LMP-O), and the LMP results derived from ACOPF and
DCOPF solver in MATPOWER (denoted as LMP-AC and
LMP-DC). The result obtained from ACOPF is set as the
benchmark result to calculate the accuracy of other results.
A.

ALMP Accuracy Evaluation.
Since the ALMP is mostly concerned in the market, we
compare the accuracy of ALMP obtained from the proposed
method and that from result derived from LMP-O and DCOPF,
taking ACOPF as the standard result. The index of the average
error of ALMP (AEA) is introduced as:

()
1
/
N
ACAC
iii i
ALMPALMPALMP
AEA
N
=
-
=
�

(19)
Here
AC
i
ALMP
is the ALMP result at bus
i
from ACOPF.
In the ACOPF model, when the voltage constraint is active,
the voltage component influences the LMP. In order to
demonstrate the effectiveness of our proposed model in
considering the voltage constraint, three scenarios with
different voltage constraints are carried out in the case study:
1)

Loose voltage constraint: minmax
0.90,1.10
VV
==
;
2)

Normal voltage constraint: minmax
0.95,1.05
VV
==
;
3)

Tight voltage constraint: minmax
0.97,1.03
VV
==
�
Fig.2 shows the AEA of the loss-embedded LMP model
(ALMP-L), LMP model without loss (ALMP-O) and DCOPF
model (ALMP-DC) when load level changes in different
voltage constraint scenarios. ALMP-O performs almost the
same as ALMP-DC in loose and normal scenarios and slightly
better than ALMP-DC in tight scenario. This indicates that the
linearized model without power loss is comparable with
DCOPF. In addition, the ALMP-L results outperform both
ALMP-O and ALMP-DC in all scenarios and under all load
levels. At 0.95 p.u. load level, the AEA of ALMP-DC is above
3% in all scenarios, while the AEA of ALMP-L is around 1.5%.
LOPF performs 50% better than DCOPF. As load level rises,
AEA of ALMP-L has a trend of increasing, but it still
outperforms DCOPF on ALMP. The results show that the
proposed linearize LMP model with power loss shows better
accuracy in calculating ALMP when comparing with DCOPF,
and power loss component plays an important role in ALMP
accuracy.

Figure 2 AEA in three voltage scenarios
B.

ALMP and RLMP Analysis
The case in tight voltage constraint is further analyzed. Fig.3
presents the ALMP and RLMP at each bus from LMP-L,
DCOPF and ACOPF respectively (DCOPF does not have
RLMP). LMP-L tracks the fluctuation of ALMP derived from
ACOPF at most buses. However, DCOPF gives a fixed ALMP
value at all buses, thus it has massive error compared with
ACOPF. One of the reasons is that DCOPF does not consider
the limit of bus voltage and reactive power. In the meanwhile,
the RLMP results from LMP-L model is almost the same as that
from ACOPF, which indicates that model also performs well in
RLMP calculation.
($/MW)
($/Mvar)

Figure 3 ALMP and RLMP in tight voltage constraint
Fig. 4 illustrates that voltage magnitudes calculated from
LMP-L are also accurate comparing to LMP-AC. Moreover, it
indicates whether the voltage magnitude at certain bus reaches
its boundary or not, which is clearly reflected on the LMP value
as voltage component. For example, the voltage magnitudes at
bus 10, 25, 66 and 69 reach their upper limit, which means they
already have abundant reactive power. Consequently, RLMP in
Fig. 3 are relatively low at these buses. On the contrary, the
voltage magnitudes at bus 58 and 117 reach their lower limit,
making their RLMP relatively higher. In fact, ALMP follows
the same rule as RLMP but it’s not as apparent as RLMP shown

6
in Fig. 3. As a result, LMP especially RLMP can work as an
indicator of reactive power demand. It further provides
incentives for the reactive power compensators to join the
market.

Figure 4 Bus voltage magnitude at each bus
Moreover, Fig. 5 shows the active and reactive power
generated by all generators. The active power generation result
of LMP-L is closer to ACOPF than DCOPF. As for the reactive
power generation, LMP-L is nearly the same as that from
ACOPF at the marginal generators such as generator NO. 23,
24, and 35. The marginal generators altogether determine the
energy term of LMP which is the majority of LMP. Therefore,
locating the right marginal generators greatly improves the
LMP results.
5
101520253035404550
Generator
0
200
400
600
800
(MW)
Active power generation
Pmax
LMP-DC
LMP-AC
LMP-L
5
101520253035404550
Generator
-50
0
50
70
(Mvar)
Reactive power generation
Qmax
Qmin
LMP-AC
LMP-L

Figure 5 Generators output
(MW)
(Mvar)

Figure 6 Active and reactive power flows
Fig. 6 presents the active and reactive power flow through all
186 branches. It indicates that none of the branch limits is
reached. Thus, the congestion component does not have
contribution to LMP in this case. That’s another reason why the
DCOPF gives a fixed ALMP value at all buses.
Above all, through the case study on IEEE-118 bus system,
this model is proved to be more accurate than DCOPF when
calculating ALMP. This model also shows high accuracy in
RLMP calculation compared with ACOPF model.
IV.

C
ONCLUSION

In this paper, a GSDF-based LMP model for both ALMP and
RLMP calculation with power loss is proposed. The necessity
to price the reactive power is first discussed. Then a linearized
LMP model is derived from an augmented GSDF matrix. Loss
equations and loss-related factors are added to enrich the model.
Through derivation of Lagrange function, LMP can be
decomposed to four components (energy, congestion, voltage
limitation and power loss). A case study on the IEEE-118 bus
case validates the accuracy of our model in LMP calculation.
The case demonstrates that ACOPF constraints can be clearly
reflected on LMP components. Therefore, with the congestion
and voltage components of LMP, the operator can effectively
manage the branch congestion and voltage support respectively.
It is also indicated that the power loss component in our linear
LMP model is effective to improve the LMP results.
R
EFERENCES

[1] Yong Fu and Zuyi Li, “Different models and properties on LMP
calculations,” in 2006 IEEE Power Engineering Society General Meeting,
2006, pp. 1-11.
[2] G. T. Heydt et al., “Pricing and Control in the Next Generation Power
Distribution System,” IEEE Trans. Smart Grid, vol. 3, no. 2, pp. 907–914,
Jun. 2012.
[3] H. Yuan, F. Li, Y. Wei, and J. Zhu, “Novel linearized power flow and
linearized OPF models for active distribution networks with application in
distribution LMP,” IEEE Trans. Smart Grid, vol. 9, no. 1, pp. 438–448,
2018.
[4] L. Bai, J. Wang, C. Wang, C. Chen, and F. Li, “Distribution Locational
Marginal Pricing (DLMP) for Congestion Management and Voltage
Support,” IEEE Trans. Power Syst., vol. 33, no. 4, pp. 4061–4073, 2018.
[5] P. Pareek and A. Verma, “Linear OPF with linearization of quadratic
branch flow limits,” in 2018 IEEMA Engineer Infinite Conference
(eTechNxT), 2018, pp. 1–6.
[6] Q. Hou, N. Zhang, J. Yang, C. Kang, Q. Xia, and M. Miao, “Linearized
Model for Active and Reactive LMP Considering Bus Voltage
Constraints,” IEEE Power Energy Soc. Gen. Meet., vol. 2018-Augus, pp.
1–5, 2018.
[7] X. Cheng and T. J. Overbye, “An Energy Reference Bus Independent
LMP Decomposition Algorithm,” IEEE Trans. Power Syst., vol. 21, no. 3,
pp. 1041–1049, Aug. 2006.
[8] F. Li and R. Bo, “DCOPF-Based LMP Simulation : Algorithm ,
Comparison With ACOPF, and Sensitivity,” Power, vol. 22, no. 4, pp.
1475–1485, 2007.
[9] Z. Yang et al., “LMP Revisited: A Linear Model for the Loss-Embedded
LMP,” IEEE Trans. Power Syst., vol. 32, no. 5, pp. 4080–4090, Sep. 2017.
[10] Z. Yang, H. Zhong, A. Bose, T. Zheng, Q. Xia, and C. Kang, “A
Linearized OPF Model With Reactive Power and Voltage Magnitude: A
Pathway to Improve the MW-Only DC OPF,” IEEE Trans. Power Syst.,
vol. 33, no. 2, pp. 1734–1745, Mar. 2018.
[11] J. Yang, N. Zhang, C. Kang, and Q. Xia, “A State-Independent Linear
Power Flow Model with Accurate Estimation of Voltage Magnitude,”
IEEE Trans. Power Syst., vol. 32, no. 5, pp. 3607–3617, 2017.
[12] MATPOWER, “Case118.” [Online]. Available:
https://matpower.org/docs/ref/matpower6.0/case118.html.
[13] IIT, “‘Index of data Illinois Institute of Technology’ 118_UMP.” [Online].
Available: http://motor.ece.iit.edu/data/.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
