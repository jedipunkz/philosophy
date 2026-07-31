---
source: "https://www.semanticscholar.org/paper/12c54f789744469c136efa0c7bd95263e0877438"
title: "AI-Aristotle: A physics-informed framework for systems biology gray-box identification"
author: "Nazanin Ahmadi Daryakenari, Mario De Florio, K. Shukla, G. Karniadakis"
year: "2023"
publication: "PLoS Comput. Biol."
download: "https://journals.plos.org/ploscompbiol/article/file?id=10.1371/journal.pcbi.1011916&type=printable"
pdf: "https://journals.plos.org/ploscompbiol/article/file?id=10.1371/journal.pcbi.1011916&type=printable"
captured_at: "2026-07-31T03:39:06Z"
updated_at: "2026-07-31T03:39:06Z"
capture_tool: "scrapem"
source_name: "semanticscholar"
keyword: "アリストテレス"
query: "Aristotle"
tags:
  - "古代哲学"
  - "倫理学"
  - "形而上学"
status: raw
---

# AI-Aristotle: A physics-informed framework for systems biology gray-box identification

- 著者: Nazanin Ahmadi Daryakenari, Mario De Florio, K. Shukla, G. Karniadakis
- 年: 2023
- 掲載情報: PLoS Comput. Biol.
- 情報源: [semanticscholar](https://www.semanticscholar.org/paper/12c54f789744469c136efa0c7bd95263e0877438)
- ダウンロード: https://journals.plos.org/ploscompbiol/article/file?id=10.1371/journal.pcbi.1011916&type=printable
- PDF: https://journals.plos.org/ploscompbiol/article/file?id=10.1371/journal.pcbi.1011916&type=printable

## Obsidian Links

- 研究動向: [[アリストテレス-現代研究動向]]
- キーワード: [[アリストテレス]]
- 関連分野: [[古代哲学]]
- 関連分野: [[倫理学]]
- 関連分野: [[形而上学]]
- 関連タグ: #古代哲学 #倫理学 #形而上学

## Abstract

Discovering mathematical equations that govern physical and biological systems from observed data is a fundamental challenge in scientific research. We present a new physics-informed framework for parameter estimation and missing physics identification (gray-box) in the field of Systems Biology. The proposed framework—named AI-Aristotle—combines the eXtreme Theory of Functional Connections (X-TFC) domain-decomposition and Physics-Informed Neural Networks (PINNs) with symbolic regression (SR) techniques for parameter discovery and gray-box identification. We test the accuracy, speed, flexibility, and robustness of AI-Aristotle based on two benchmark problems in Systems Biology: a pharmacokinetics drug absorption model and an ultradian endocrine model for glucose-insulin interactions. We compare the two machine learning methods (X-TFC and PINNs), and moreover, we employ two different symbolic regression techniques to cross-verify our results. To test the performance of AI-Aristotle, we use sparse synthetic data perturbed by uniformly distributed noise. More broadly, our work provides insights into the accuracy, cost, scalability, and robustness of integrating neural networks with symbolic regressors, offering a comprehensive guide for researchers tackling gray-box identification challenges in complex dynamical systems in biomedicine and beyond.

## Citation

DOI: 10.1371/journal.pcbi.1011916

## PDF Text

PLOS COMPUTATIONAL BIOLOGY

RESEARCH ARTICLE

AI-Aristotle: A physics-informed framework for systems biology gray-box identification
Nazanin Ahmadi Daryakenari ID1☯, Mario De Florio ID2☯, Khemraj Shukla2, George
Em Karniadakis ID2*
1 Center for Biomedical Engineering, School of Engineering, Brown University, Providence, Rhode Island,
United States of America, 2 Division of Applied Mathematics, Brown University, Providence, Rhode Island,
United States of America
☯ These authors contributed equally to this work.
* george_karniadakis@brown.edu

a1111111111
a1111111111
a1111111111
a1111111111
a1111111111

OPEN ACCESS
Citation: Ahmadi Daryakenari N, De Florio M,
Shukla K, Karniadakis GE (2024) AI-Aristotle: A
physics-informed framework for systems biology gray-box identification. PLoS Comput Biol 20(3): e1011916. https://doi.org/10.1371/journal.
pcbi.1011916
Editor: Piero Fariselli, Universita degli Studi di
Torino, ITALY
Received: October 4, 2023
Accepted: February 13, 2024
Published: March 12, 2024
Copyright: © 2024 Ahmadi Daryakenari et al. This is an open access article distributed under the terms of the Creative Commons Attribution
License, which permits unrestricted use, distribution, and reproduction in any medium, provided the original author and source are credited.
Data Availability Statement: All data and codes used in this manuscript are available on GitHub at https://github.com/mariodeflorio/AI-Aristotle.
Funding: NAD and GEK gratefully acknowledge the
National Institutes of Health (NIH) Spleen grant
R01HL154150. MD, KS, and GEK gratefully acknowledge the Office of Naval Research (ONR)
Vannevar Bush grant N00014-22-1-2795. The funders had no role in study design, data collection

Abstract
Discovering mathematical equations that govern physical and biological systems from observed data is a fundamental challenge in scientific research. We present a new physicsinformed framework for parameter estimation and missing physics identification (gray-box)
in the field of Systems Biology. The proposed framework—named AI-Aristotle—combines the eXtreme Theory of Functional Connections (X-TFC) domain-decomposition and Physics-Informed Neural Networks (PINNs) with symbolic regression (SR) techniques for parameter discovery and gray-box identification. We test the accuracy, speed, flexibility, and robustness of AI-Aristotle based on two benchmark problems in Systems Biology: a pharmacokinetics drug absorption model and an ultradian endocrine model for glucose-insulin interactions. We compare the two machine learning methods (X-TFC and PINNs), and moreover, we employ two different symbolic regression techniques to cross-verify our results. To test the performance of AI-Aristotle, we use sparse synthetic data perturbed by uniformly distributed noise. More broadly, our work provides insights into the accuracy, cost, scalability, and robustness of integrating neural networks with symbolic regressors, offering a comprehensive guide for researchers tackling gray-box identification challenges in complex dynamical systems in biomedicine and beyond.

Author summary
Our study addresses the fundamental challenge of uncovering mathematical rules governing physical and biological systems from real-world data. We introduce a novel framework, AI-Aristotle, designed for parameter estimation and identifying hidden physics
(gray-box) in Systems Biology. AI-Aristotle combines the powerful eXtreme Theory of
Functional Connections (X-TFC), Physics-Informed Neural Networks (PINNs), and symbolic regression (SR) techniques to discover parameters and uncover hidden relationships.
Our work offers guidance to researchers addressing gray-box identification challenges in complex dynamic systems, including applications in biomedicine and beyond.

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

1 / 33

PLOS COMPUTATIONAL BIOLOGY

and analysis, decision to publish, or preparation of the manuscript.
Competing interests: The authors have declared that no competing interests exist.

AI-Aristotle

1 Introduction
One of the most coveted tasks in Machine Learning is the discovery of new physics laws from observed and experimental data. When dealing with dynamical systems, a classic goal for inverse problems is parameter discovery, where experimental data and systems of differential equations are leveraged to estimate the unknown parameters governing [1]. In some cases, only partial knowledge of the physics may be available, which means one or several terms of the system of equations are unknown. This is the case with the so-called Gray-Box model, where an inversion can be performed to recover the missing terms [2].
One of the first attempts to extrapolate governing equations from observed data is presented in the well-known work by Brunton et al. [3], in which the authors propose a new school of thought for dynamical system discovery problem from the perspective of sparse regression [4] and compressed sensing [5]. In particular, they take advantage of the fact that most physical systems are described by only a few relevant terms governing the dynamics, making the governing equations sparse in a high-dimensional non-linear function space. This method named SINDy—Sparse Identification of Nonlinear Dynamics—depends on the choice of the candidate non-linear functions library and the availability and quality of the data. Thus, it is not a generalized method and works better if guided by the available knowledge in the form of constraints on the functional form of the phenomena under study. For example, given the trend of the observed data, one can approximately understand if it is a trigonometric or polynomial trend and build the library accordingly. SINDy has shown its capability in identifying non-linear dynamical systems from data without previous assumptions of the forms of the differential equations governing the phenomena.
Another method to retrieve governing equations from data has been proposed by Udrescu et al. [6]. In this paper, the authors make use of symbolic regression (SR), which aims to find a symbolic expression that accurately represents an unknown function based on a given dataset.
They developed a novel recursive multidimensional symbolic regression algorithm, named
AI-Feynman, that combines neural network techniques with physics-inspired strategies. The efficiency of this method has been proved by discovering 100 equations from the Feynman Lectures on Physics, outperforming the accuracy of the state-of-the-art publicly available software.
However, despite the groundbreaking capability of this work, there are some drawbacks and areas for improvement. The method currently focuses on equations involving elementary functions but does not handle equations involving derivatives and integrals commonly found in physics. Integrating the capability to discover such equations would be valuable. Also, while the AI-Feynman shows promise, it could further benefit from combining the strengths of genetic algorithms and its approach to generate a more robust and versatile equation discovery tool. Overall, the development and refinement of symbolic regression algorithms continue to evolve, offering exciting possibilities for future discoveries in the realm of physics and beyond.
In this research direction, a new framework named AI-Descartes has been recently published [7]. In this paper, the authors address the challenge of deriving meaningful mathematical models from both axiomatic knowledge and experimental data by combining logical reasoning with SR. The novelty of this method lies in the attempt to generate models that are consistent with general logical axioms. The authors showcase their method’s effectiveness by applying it to three classic scientific laws: Kepler’s third law of planetary motion, Einstein’s relativistic time-dilation law, and Langmuir’s theory of adsorption. They demonstrate the capability to discover governing laws even with limited data points, emphasizing the importance of logical reasoning in distinguishing between candidate formulas with similar data-fit accuracy.
However, this method relies on the correctness and completeness of background theories, which may not always hold, and the development of further techniques such as abductive

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

2 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

reasoning [8] for partially addressing incomplete theories would be needed. Scaling behavior remains a challenge, especially regarding the undecidability of certain logical types and variations in run-time performance.
Another recently developed SR package, named Feyn [9] and based on the symbolic regressor QLattice, is showing great performance and capabilities, especially for small data sets, where traditional machine learning techniques such as gradient boosting and random forests tend to overfit [10]. Christensen et al. [11] efficiently used Feyn on clinical omics datasets to generate high-performing models to predict disease outcomes and to reveal putative disease mechanisms.
Other approaches using particular type of Neural Networks called Random Projection Neural Networks (RPNNs) [12–14] are used in combination with SR. RPNNs demonstrated great efficiency in solving forward problems of stiff ODEs and DAEs, outperforming traditional solvers [15, 16]. In Ref. [17], RPNNs are used for learning PDEs from spatio-temporal data and for the construction of the bifurcation diagram of the learned PDE. In a recent work [18],
RPNNs are used to model a representation for SR called Interaction Transformation [19], showing the capability of this framework in drastically reducing the computational effort. In another work [20], a single-layer NN is combined with SR. In this approach, the SR layer, incorporating mathematical operators and basis functions, is constructed randomly instead of using genetic programming, and the output weighting parameters are optimized through least-squares optimization. The use of least-squares optimization significantly reduced computational time, resulting in system models based on simple analytic expressions that accurately represent the input-output relationship of dynamic systems. Recently, RPNNs and SR were combined in the AI-Lorenz [21] to discover chaotic dynamical systems in a black-box fashion, when the differential equations of the model are totally unknown.
One of the earliest works on addressing “gray-box” identification for nonlinear dynamical systems is the one of Ref. [2]. The gray-box in this paper is composed of a known part, represented by a system of Ordinary Differential Equations (ODEs), and unknown parts, which are approximated using neural networks. The paper illustrates this approach by applying it to model a complex reacting system with nonlinear kinetics for parameter discovery. The authors also highlight the challenges of working with discrete-time models and the advantages of using continuous-time approximations for a more nuanced understanding of system behavior.
Other gray-box identification and parameter estimation methodologies were applied to a wide range of applications, such as phase field systems, biotechnology, and optogenetics [22–26].
More recently in [27], NNs and Gaussian Processes were used to perform gray box identification of PDEs based on stochastic Monte Carlo simulators in biological systems and in particular for the chemotaxis motility.
The PINN frameworks [28] are advancing the state-of-the-art methodologies for inverse problems of parameter discovery. Particularly challenging is the scenario in which we have a highly nonlinear dynamics system with many unknown parameters and very few available experimental data to leverage. This challenge has been addressed in a systems-biologyinformed deep learning algorithm that incorporates the system of ODEs into the neural networks. In the works [29, 30], the authors proved the efficiency of this new algorithm to infer the dynamics of unobserved species using only a few scattered and noisy measurements by testing it for benchmark problems in systems biology.
In this work, we propose a new framework named AI-Aristotle to perform parameter discovery and gray-box identification for problems in Systems Biology. We employ two neural networks based methods for the unknown terms approximation, such as PINNs and X-TFC
[31] with domain decomposition [15], and two symbolic regression algorithms for the mathematical explicitation of the gray-box model, such as PySR [32, 33] and gplearn [34]. Our

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

3 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

framework is tested for two problems. The first one is a three-compartment pharmacokinetics model describing single-dose drug absorption. The second, more challenging problem is an ultradian endocrine model describing the glucose-insulin interaction. PINNs and X-TFC have been previously employed for gray box identification [21, 35, 36]. The novelty of this work lies in its unique integration of these methods and their concatenation with symbolic regression algorithms. This integrated framework allows the user to select the neural network-based module depending on the data availability, using two different symbolic regression algorithms for cross-validation. Unlike the SINDy method, which encounters difficulties with high-dimensional noisy data, the symbolic regression methods in this framework effectively address these challenges.
This paper is organized as follows. In Section 2, we present an introduction of the physicsbased models used for our simulations. In Section 3, we report the two Neural Networks methods for solving the inverse problem with data and physics models and the two SR algorithms used to explicitly identify the previously retrieved gray-boxes. In Section 4, we report the results obtained by the two NN methods and the two SR algorithms for different test cases involving both parameter discovery and gray-box identification. Finally, we summarize conclusions and discussion in Section 4.3.2.

2 Models
In this section, the mathematical models describing the phenomena of our simulations are introduced. These models are designed to capture the dynamic interactions within specific biological processes, such as drug absorption and glucose-insulin interaction, offering physicsbased knowledge of the behavior and characteristics of the systems under study.

2.1 Pharmacokinetics model
The first model we aim to use for our simulations is a single-dose compartmental Pharmacokinetics (PK) model [37], represented by the following system of ODEs:
8
dB
8
>
>
¼ kg G kb B
>
Bð0Þ ¼ 0
>
>
dt
>
>
>
>
>
>
>
>
>
>
>
>
<
<
dG
s:t:
Gð0Þ ¼ 0:1mg
ð1Þ
¼ kg G
>
>
dt
>
>
>
>
>
>
>
>
>
>
>
>
:
> dU
>
>
Uð0Þ ¼ 0
:
¼ kb B
dt
This model evaluates the variation of drug concentration in three compartments, in a time range [0, 10] hours. The drug is initially introduced in the GI-tract (first compartment G), where it dissolves and diffuses into the bloodstream (second compartment B). Finally, the drug is eliminated from the bloodstream through the liver, kidneys, and urinary tract (third compartment U). The parameters kg = 0.72h−1 and kb = 0.15h−1 represent the rates at which the drug diffuses from the GI-tract into the bloodstream, and then eliminated from the bloodstream through the liver, kidneys, and urinary tract, respectively. The intake drug is considered to be 0.1μg of antibiotic tetracycline. In Section 4, we will show our simulations using this model for two test cases: 1) Parameters discovery, and 2) Gray-Box identification. With “GrayBox”, we indicate the missing terms of a model. For this PK model, the missing term considered is the right-hand-side of the first ODE, which we approximate with an unknown function

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

4 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

h(t) as follows:
8
dB
>
>
¼ hðtÞ
>
>
dt
>
>
>
>
>
>
>
>
>
<
dG
¼ kg G
>
dt
>
>
>
>
>
>
>
>
>
>
>
>
: dU ¼ k B
b dt

s:t:

8
Bð0Þ ¼ 0
>
>
>
>
>
>
>
>
>
>
<
Gð0Þ ¼ 0:1mg
>
>
>
>
>
>
>
>
>
>
:
Uð0Þ ¼ 0

ð2Þ

which we aim to obtain by using available data for B, G, and U.

2.2 Ultradian Endocrine model
The second model used in our simulations is an ultradian model for the glucose-insulin interaction [38], which is modeled by 6 state variables and 30 parameters [29]. This model describes the existence of rhythmic oscillations in both glucose and insulin levels within the body that occur on a relatively short timescale, typically less than 24 hours. In particular, in our simulation, we will use a time range [0, 1800] minutes. It results in the following system of ODEs:
!
8
dI
I
>
I
p p
>
i
>
¼ f1 ðGÞ E
>
>
>
dt
Vp Vi
>
>
>
>
>
>
!
>
>
>
Ip dIi
Ii
Ii
>
>
>
¼E
>
>
dt
V
V
ti
>
p i
>
>
>
>
>
>
>
dG
>
>
>
< dt ¼ f4 ðh3 Þ þ IG ðtÞ f2 ðGÞ
>
>
>
>
dh1 1
>
>
¼ ðIp h1 Þ
>
>
>
dt td
>
>
>
>
>
>
>
>
dh2 1
>
>
¼ ðh1 h2 Þ
>
>
td dt
>
>
>
>
>
>
>
>
>
> dh3 ¼ 1 ðh
: h3 Þ
td 2
dt

Ip tp

8
I ð0Þ ¼ 36mU=ml
>
>
> p
>
>
>
>
>
>
>
>
>
Ii ð0Þ ¼ 44mU=ml
>
>
>
>
>
>
>
>
>
>
>
>
Gð0Þ ¼ 110mg=dl
>
<

f3 ðIi ÞG
s:t:

>
>
>
>
h1 ð0Þ ¼ 0
>
>
>
>
>
>
>
>
>
>
>
>
h2 ð0Þ ¼ 0
>
>
>
>
>
>
>
>
>
: h3 ð0Þ ¼ 0

ð3Þ

The three main variables of this model are the plasma insulin concentration Ip, the interstitial insulin concentration Ii, and the glucose concentration G. The last three variables h1, h2, and h3—a three-stage linear filter—represent the delay process between insulin and glucose production [38]. The functions f1, f2, f3, and f4, represent the insulin secretion, the insulinindependent glucose utilization, the insulin-dependent glucose utilization, and insulinPLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

5 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

dependent glucose utilization, respectively [39], and they are expressed as follows:
R
� m

�;
þ a1
!!
G
f2 ðGÞ ¼ Ub 1 exp
;
C2 Vg
!
1
Um f3 ðIi Þ ¼
U0 þ
; b
C3 Vg
1 þ ðkIi Þ

f1 ðGÞ

¼

f4 ðh3 Þ

¼

1 þ exp

G
Vg C1

R
� �g
1 þ exp a C5hV3 p

�� ;
1

where
�
�
1 1
1
; k¼
þ
C4 Vi Eti and IG(t) is the exogenous (externally driven) glucose delivery rate. In our simulations, we define it over N = 3 nutrition events, at time tj (minutes) with a carbohydrate quantity mj
(grams):
N
X

mj k expðkðtj

IG ðtÞ ¼

tÞÞ;

ð4Þ

j¼1

where (tj, mj) = [(300, 60)(650, 40)(1100, 50)](min, g), and the parameters governing this system of ODEs are listed in Table 1. Fig 1 shows the flow diagram of the glucose-insulin model, where the circles represent the three main state variables (Ip, Ii, G), the solid arrows represent the input and output flows and rate of exchange, and dashed arrows represent functional relationships. The delay arrow denotes the delay process of h1, h2, h3 state variables.
Also for this second model, we aim to pursue parameter discovery and graybox identification. For the latter case, the missing terms we approximate with two unknown functions, f(t) and g(t), which are in the first two ODEs, as follows:
8
dIp
>
>
¼ f1 ðGÞ þ f ðtÞ
>
>
dt
>
>
>
>
>
>
dIi
>
>
>
¼ gðtÞ
>
> dt
>
>
>
>
>
>
dG
>
>
>
< dt ¼ f4 ðh3 Þ þ IG ðtÞ f2 ðGÞ f3 ðIi ÞG
ð5Þ
>
dh1 1
>
>
>
¼ ðIp h1 Þ
>
>
td dt
>
>
>
>
>
>
dh2 1
>
>
¼ ðh1 h2 Þ
>
>
>
td
> dt
>
>
>
>
>
dh
1
>
: 3 ¼ ðh2 h3 Þ: td dt

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

6 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Table 1. Ultradian Endocrine model: List of parameters for the model. The search ranges are listed only for the five parameters used for the parameter discovery in our simulations.
Parameter

Nominal value

Unit

Search range

Vp

3

lit

–
–

Vi

11

lit

Vg

10

lit

–

E

0.2

lit min−1

(0.1, 0.3)

tp

6

min

(4, 8)

ti

100

min

(60, 140)
–

td

12

min

k

0.0083

min−1

–

Rm

209

mU min−1

(41.8, 376.2)

a1

6.6

–

(1.32, 11.88)

C1

300

mg lit−1

–

C2

144

mg lit−1

–

C3

100

mg lit−1

–

C4

80

mU lit−1

–

C5

26

mU lit−1

–

Ub

72

mg min−1

–

U0

4

mgmin−1

–

Um

90

mg min

−1

–

Rg

180

mg min−1

–

α

7.5

–

–

β

1.772

–

–

https://doi.org/10.1371/journal.pcbi.1011916.t001

Fig 1. Ultradian Endocrine model: Flow diagram. The circles represent the three main state variables (Ip, Ii, G), the solid arrows represent the input and output flows and rate of exchange, and the dashed arrows represent functional relationships.
https://doi.org/10.1371/journal.pcbi.1011916.g001

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

7 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Fig 2. AI-Aristotle framework for gray-box identification: 1. The observed data and the partial knowledge of physics are used to train the selected neural networkbased module. 2. The selection of the neural networks-based module needs to be done between (a) X-TFC, recommended for high-resolution data and missing terms discovery, and (b) PINN, recommended for sparse data and parameter estimation. The neural network outputs are the time-dependent representations of the missing terms of the dynamical systems, which are fed into the symbolic regression algorithm. 3. The selected Symbolic Regression module identifies the mathematical expressions of the missing terms. It is recommended to use both symbolic regressors for cross-validation. 4. The full knowledge of physics is now available, allowing forward modeling performance.
https://doi.org/10.1371/journal.pcbi.1011916.g002

3 Methodology
As mentioned in the Introduction section, the parameter discovery and approximation of the unknown terms in the systems of ODEs are performed by two NN-based methods, while the symbolic regression is performed by two different algorithms, to cross-verify the mathematical expressions obtained. In this section, we present some details of these methods that are included in the AI-Aristotle framework, whose overall schematic is shown in Fig 2.

3.1 X-TFC
The first NN-based method presented uses a single-layer random projection neural network.
For the sake of simplicity, we will show its implementation for the gray-box identification in the pharmacokinetics model only, since the implementation for the ultradian endocrine model is similar.
Different techniques are combined to build this algorithm for solving both forward and inverse problems involving differential equations. The first one is a functional interpolation technique named the Theory of Functional Connections (TFC) [40, 41]. According to TFC
[42], we can approximate the unknown solutions of our system of ODEs, taking into consideration the initial conditions, with the so-called constrained expressions (CE) as follows:
T

ð6aÞ

T

ð6bÞ

T

ð6cÞ

B ¼ ðσðtÞ

σð0ÞÞ βB þ Bð0Þ

G ¼ ðσðtÞ

σð0ÞÞ βG þ Gð0Þ

U ¼ ðσðtÞ

σð0ÞÞ βU þ Uð0Þ

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

8 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

whose derivatives can be analytically expressed: dB
¼ cσ 0 T βB
dt

ð7aÞ

dG
¼ cσ 0 T βG
dt

ð7bÞ

dU
dt

ð7cÞ

¼ cσ 0 T βU

The parameter c represents a mapping coefficient that maps the time domain t into the activation function domain. To these systems, we need to add the NN approximation of the unknown term h(t), which is hðtÞ ¼ σðtÞβh :

ð8Þ

Here, σ is the free-chosen function of the CE. No matter what free-chosen function will be selected, the CE will always satisfy the initial conditions exactly. According to the X-TFC
framework [31], we select a single-layer NN as a free-chosen function, such as
2
3T
sðw1 t þ b1 Þ
6
7
L
X
6
7
.
6
7 β ¼ σT β
ð9Þ
.
gðtÞ ¼
bj sðwj t þ bj Þ ¼ 6
.
7
4
5
j¼1
sðwL t þ bL Þ
where L is the number of neurons, wj 2 R is the jth input weight connecting the input node with the jth neuron, bj 2 R with j = 1, . . ., L is the jth output weight connecting the output node with the jth neuron, bj is the bias of the jth neuron, and σj(�) is the NN’s activation function, which is selected by the user (for all the simulations in this work, we select a tanh activation function. The motivation for this choice is reported in the first section of S1 Text. In the extreme learning machine algorithm [43], input weights and biases are randomly pre-selected
(uniform random distribution), thus the only unknown parameters that need to be computed are the output weights β = [β1, . . ., βL]T. Once the CEs are built, they can be replaced in the system of ODEs of Eq (2), to obtain the loss functions
LB ¼
�
LG ¼
LU ¼

cσ 0 ðtÞ

T

cσ 0 ðtÞ βU þ σðtÞβh

T

kg ðσðtÞ

T

cσ 0 ðtÞ βU þ kb ðσðtÞ

σð0ÞÞ

�T

ð10aÞ
βG þ Gð0Þ

ð10bÞ

σð0ÞÞ βB þ kb Bð0Þ

ð10cÞ

T

T

ð10dÞ

T

ð10eÞ

T

ð10f Þ

^
LdataB ¼ B

ðσðtÞ

σð0ÞÞ βB þ Bð0Þ

^
LdataG ¼ G

ðσðtÞ

σð0ÞÞ βG þ Gð0Þ

^
LdataU ¼ U

ðσðtÞ

σð0ÞÞ βU þ Uð0Þ;

^ and U
^ G,
^ are the available observed data of the three variables. As we can see, now we where B;

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

9 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

have reduced the problem into a system of linear equations of the type Ax = b, where the unknown x is the vector of output weights β. However, here we show the procedure to solve it as a system of non-linear equations (that will be the case of the Ultradian Endocrine model).
When dealing with a system of non-linear ODEs, the next step is to build the Jacobian matrix, by deriving the six previous losses with respect to βB, βG, βU, and βf. For the pharmacokinetics model, the Jacobian is
2
3
@LB
@LB
@LB
0
6 βB
βG
βh 7
6
7
6
7
6
7
@LG
6
7
0
0 7
6 0
6
7
βG
6
7
6
7
6 @L
7
@LU
U
6
0
0 7
6 β
7
βU
6
7
B
6
7
J ¼6
ð11Þ
7
6 @Ldata
7
B
6
7
0
0
0 7
6 β
6
7
B
6
7
6
7
6
7
@L
dataG
6 0
0
0 7
6
7
βG
6
7
6
7
6
7
4
5
@LdataU
0
0
0:
βU
The unknown vector β is computed by iteratively solve the linear system J Dbk ¼ L. Each k-th iteration corresponds to an update of the output weights βk+1 = βk + Δβk, where
1
Dbk ¼ ðJ T ðbk ÞJ ðbk ÞÞ J T ðbk ÞLðbk Þ. If the Jacobian is rank-deficient, it is good practice to minimize the value of the Euclidean norm to achieve better performance or compute the
Moore-Penrose pseudoinverse of the Jacobian as proposed in Refs. [16, 17]. Once all the output weights β are computed, they will be replaced into the CEs of Eqs (6a) to (6c) and (8) to find our sought solutions. In this work, X-TFC is used in a domain-decomposition fashion
[15, 44], where the time-domain is decomposed into several sub-domains with equispaced time steps, and the algorithm is applied to each sub-domain subsequently, such that the solution found at the interface becomes the new initial condition for the subsequent iteration of the algorithm in the next sub-domain. A schematic of the X-TFC algorithm to solve the graybox inverse problem for the pharmacokinetics model is shown in Fig 3.

3.2 Physics-Informed Neural Networks (PINNs)
The second NN-based approach is known as Physics-Informed Neural Networks (PINNs).
This method has the capability to address both forward and inverse problems associated with differential equations by using a deep, fully connected neural network.
3.2.1 PINNs for Pharmacokinetics model. Building upon the concept of PINNs as originally proposed in reference [28], we introduce a deep learning framework that incorporates the differential equations governing the single-dose compartmental Pharmacokinetics model.
In this framework, a neural network characterized by parameters θ1 takes time t as input and generates an output vector representing the state variables u(t; θ1) = (uB(t, θ1), uG(t; θ1), uU(t;
θ1)) which serves as an approximation of the ODE solution u^ ðtÞ. To solve the gray-box inverse problem, in addition to the unknown parameters, we have an unknown component of the equation. Thus, we introduce another neural network with a different design to approximate

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

10 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Fig 3. Pharmacokinetics model: Schematic of the X-TFC algorithm. Input weights and biases are randomly selected. The last step solves iteratively a least squares system, thus no back-propagation is involved in the training, allowing fast computational times.
https://doi.org/10.1371/journal.pcbi.1011916.g003

the unknown term h(t). The system of ODEs for this model is as follows:
8
dB
>
>
¼ hðt; y2 Þ
>
>
dt
>
>
>
>
>
>
<
dG
¼ kg G
>
dt
>
>
>
>
>
>
>
>
>
: dU ¼ kb B
dt

s:t:

8
Bð0Þ ¼ 0
>
>
>
>
>
>
>
<
Gð0Þ ¼ 0:1mg
>
>
>
>
>
>
>
:
Uð0Þ ¼ 0:

ð12Þ

Here, the parameters θ2 characterize the second neural network, which takes t as input and generates an output h(t; θ2).
The next crucial step involves constraining the neural network to satisfy both the scattered observations of u^ ðtÞ and the system of ODEs (12). This is achieved by constructing the loss function that takes into account terms corresponding to the observations and the ODE system.
To be more specific, let us assume that we have measurements of u^ data ¼ f^
u 1 ; u^ 2 ; . . . ; u^ M g at various time instances t1, t2, . . ., tMdata. We want to ensure that the neural network satisfies the
ODE system at specific time points t1, t2, . . ., tNode. It is important to note that the time instants t1, t2, . . ., tMdata, and t1, t2, . . ., tNode may not necessarily be on a uniform grid and can be chosen arbitrarily. Here, N is the number of collocation points, and M is the number of data points.
For computing the total loss, we employ the Self-Adaptive Loss Balanced method [45, 46].
The total loss function is defined as a function of θ1, θ2, p, λode, where p represents the unknown parameters of the ODEs, and λode is a vector representing the individual loss weights for all the state variables, i.e., λode = (λ1, λ2, . . ., λS), where S is the number of state variables.
Note that λdata and λIC are constant values equal to 1 in this study and are not trainable variables in our neural network [46]. The total loss function is defined as a function of θ1, θ2, p,
λode, where p represents the unknown parameters of the ODEs, and λode is a vector representing the individual loss weights for all the state variables, i.e., λode = (λ1, λ2, . . ., λS), where S is the number of state variables. Note that λdata and λIC are constant values, equal to 1 in this

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

11 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

study, and are not trainable variables in our neural network. The total loss is computed as follows:
Lðy1 ; y2 ; p; lode Þ ¼ lIC LIC ðy1 Þ þ ldata Ldata ðy1 Þ þ lode Lode ðy1 ; y2 ; pÞ;

ð13Þ

where
LIC ðy1 Þ ¼ ð^
u ðt0 Þ

uðt0 ; y1 ÞÞ

data
X
1 M
Ldata ðy1 Þ ¼ data
ð^
u ðtm Þ
M m¼1

0 �
�
N ode
1 X
@du ��
Lode ðy1 ; y2 ; pÞ ¼ ode
N n¼1 dt �

2

uðtm ; y1 ÞÞ

ð14Þ

2

ð15Þ

12
Fðtn ; uðtn ; y1 Þ; hðtn ; y2 Þ; pÞA

ð16Þ

tn

We emphasize that Ldata and LIC represent the discrepancies between the neural network predictions and the measured data, making them supervised losses. Conversely, Lode is derived from the ODE system and, therefore, qualifies as an unsupervised loss. In the final step, we simultaneously determine the parameters y∗1 , y∗2 of both neural networks and the unknown
ODE parameters p� by minimizing the loss function using gradient-based optimization methods, such as the Adam optimizer [47] and L-BFGS optimizer [48]. Additionally, we determine the l∗ode vector by updating adaptive weights in each epoch by solving: y∗1 ; y∗2 ; p∗ ; l∗ode ¼ arg max min Lðy1 ; y2 ; p; lode Þ
lode y1 ;y2 ;p

ð17Þ

For the training process, where our goal is to predict the unknown term h(t; θ2) and the values of parameters simultaneously, we employ the Adam optimizer with default hyperparameters and a learning rate of 10−4. Training is performed on the entire dataset. Since our total loss comprises two supervised losses and one unsupervised loss, we adopt a two-stage training strategy as follows:
1. Recognizing that supervised training typically yields faster convergence than unsupervised training, we initially train the network using the two supervised losses, Ldata and LIC , for a set number of iterations. This initial training phase enables the network to quickly align with the observed data points.
2. Subsequently, we continue the training process, incorporating all three losses.
Empirical observations demonstrate that this two-stage training approach expedites network convergence. The specific number of iterations for each stage and parameters for the implementation are detailed in Section 4.1. A schematic of the PINNs algorithm for solving the gray-box inverse problem in the pharmacokinetics model is shown in Fig 4.

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

12 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Fig 4. Pharmacokinetics model: Schematic of the PINNs algorithm for predicting the unknown term h(t; θ2) and the values of parameters simultaneously. Here, u(t; θ1) is a vector that contains all three output states. Unlike the
X-TFC network, PINN requires back-propagation, which is the expensive computational component.
https://doi.org/10.1371/journal.pcbi.1011916.g004

3.2.2 PINNs for Ultradian Endocrine model. The system of ODEs for this model is as follows:
8
dI
>
>
> p ¼ f1 ðGÞ þ f ðt; y2 Þ
>
>
dt
>
>
>
>
>
>
>
>
> dIi
>
>
¼ gðt; y2 Þ
>
>
dt
>
>
>
>
>
>
>
>
dG
>
>
¼ f4 ðh3 Þ þ IG ðtÞ f2 ðGÞ
>
>
< dt
>
>
dh1 1
>
>
>
¼ ðIp
>
>
td dt
>
>
>
>
>
>
>
>
dh2 1
>
>
¼ ðh1
>
>
>
td dt
>
>
>
>
>
>
>
>
dh
1
>
> 3 ¼ ðh2
: td dt

f3 ðIi ÞG
ð18Þ

h1 Þ

h2 Þ

h3 Þ

Here, parameters θ2 characterize the second neural network, which takes t as input and generates two outputs f(t; θ2) and g(t; θ2).
In accordance with the pharmacokinetics model, this study adopts a self-adaptive loss-balanced method and a two-stage training strategy. To expedite the neural network training process, extending the discussion from the previous section on Fully connected Neural Networks, we introduce supplementary layers following the workflow presented in [29].

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

13 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

• Input Scaling Layer: In cases where the time domain exhibits significant variation spanning multiple orders of magnitude, which can detrimentally affect neural network training, we employ a linear scaling function on the time variable t, using a value in the time domain T to obtain ~t ¼ Tt , which approximates values to be *O(1). In this study, for the time interval ranging from 0 to 1800, we have adopted a value of T = 100.
• Feature Layer: Frequently, solutions to ordinary differential equations (ODEs) display patterns such as periodicity or exponential decay. To enhance the neural network’s ability to learn these patterns, especially in multimodal solutions with multiple levels of frequencies, we incorporate a dedicated feature layer. This layer is key in capturing the complexity of multimodal solutions. The general framework remains consistent across different problems.
We utilize the set of functions e1(θ), e2(θ), . . ., eL(θ) to construct L features e1 ð~t Þ; e2 ð~t Þ; . . . ; eL ð~t Þ, as illustrated in Fig 5. If discerning a clear pattern proves challenging, it is advisable to omit the feature layer rather than introducing inaccurate information. This feature layer is a training aid and not a mandatory component for the success of the PINNs for system biology identification problems.
• Output Scaling Layer: The predicted outputs, denoted as u~ Ip ; u~ Ii ; . . . ; u~ h3 , may exhibit variations in magnitudes. To address this, we can normalize the network outputs. To standardize these outputs, we employ a normalization procedure, expressed as follows: uIp

¼ kIp u~ Ip

uIi

¼ kIi u~ Ii
..
.

uh3

¼ kh3 u~ h3 :

Here, kIp ; kIi ; . . . ; kh3 represent the magnitudes of the corresponding ODE solutions u^ Ip ; u^ Ii ; . . . ; u^ h3 . This normalization ensures that the predicted outputs are scaled consistently

Fig 5. Ultradian Endocrine model: Schematic of the PINNs algorithm for solving a gray-box identification problem.
https://doi.org/10.1371/journal.pcbi.1011916.g005

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

14 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

with the characteristics of the underlying ODE solutions. Furthermore, we introduce an additional component to this layer to facilitate the alignment of the state variables with a linear trajectory connecting the initial and final data points. This linear transformation facilitates interpreting and visualizing the model’s outputs, ensuring their alignment with meaningful data trends. In summary, the Output Scaling Layer standardizes predicted outputs while integrating a linear transformation component. This integration enhances the interpretability and relevance of the model’s results, expediting the neural network’s convergence towards an accurate solution. We observed that without the output scaling layer, the model tended to get stuck in local minima.
The list of parameters of this model can be found in Section 4.2. A schematic of the PINNs algorithm for solving the gray-box identification problem in the Ultradian Endocrine model is shown in Fig 5.

3.3 Symbolic regression
Symbolic regression is a powerful method used in machine learning, designed to discover a mathematical expression or equation that provides the optimal fit for a provided dataset.
Unlike traditional regression methods (e.g., linear regression, polynomial regression), symbolic regression seeks to discover the underlying mathematical relationship between input variables and the target variable without making assumptions about the form of the equation.
Two popular symbolic regression algorithms commonly used in this context are PySR (Python
Symbolic Regression) [32] and gplearn (Genetic Programming for Symbolic Regression) [34].
These algorithms employ different techniques to discover symbolic expressions from data, and their processes are very similar to each other.
They are SR libraries that combine genetic programming with machine learning techniques to discover mathematical expressions. The first step of their processes is creating an initial population of candidate equations represented by mathematical expressions composed of simple mathematical operations (+, −, ×, �), functions (e.g., sine, cosine, exponential), and variables.
Subsequently, each candidate equation is evaluated against the given dataset, and its performance is assessed using a fitness function, that measures how well the equation fits the data, typically by calculating the mean squared error (MSE) or a similar metric. A genetic algorithm is used to select the best-performing candidate equations for the next generation. Equations that fit the data well are more likely to be selected, while less fit equations may be removed.
Genetic operations like crossover (combining parts of two equations) and mutation (making small changes to an equation) are applied to the selected equations to create a new generation of candidate equations. This process iterates through multiple generations, continually improving the equations’ fitness until a termination condition, such as a maximum number of generations, or a threshold fitness level, is met.

4 Results
In this section, the results of our simulations are reported and discussed. The first two subsections 4.1 and 4.2 show the performance of X-TFC and PINNs in parameter discovery and gray-box identification for both the Pharmacokinetics and Ultradian Endocrine model. The synthetic data are generated by solving the forward problems with Runge-Kutta method for
PINNs, and RPNNs for X-TFC. The outputs of the gray-box identification are used as input in the symbolic regression algorithms for the symbolic distillation of both NN-based methods, whose results and performance are shown in subsections 4.3.1 and 4.3.2.

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

15 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

4.1 Pharmacokinetics
In the parameter discovery test case, we aim to infer the value of the parameters kg = 0.72h−1
and kb = 0.15h−1 of the system of ODEs in Eq (1), given a certain number of available data points of B, G, and U. The results and performance for both X-TFC and PINNs are reported in
Table 2, simulating the variation of drug concentration in the three compartments for a time domain of 50 hours. The number of data points used varies from 10 to 100, and both methods show great accuracy in retrieving both the parameters governing the ODEs. The accuracy of the methods is evaluated with the absolute difference between the nominal value of the parameters and their inference. As expected, we see an increase in accuracy while increasing the number of data points, but one can see that both methods can give great precision even for a meager dataset (10 data points—one every five hours). To substantiate this claim, particularly for PINNs, we executed the model 10 times, each with a distinct random seed. We then computed the average relative error (%) of the inferred parameter values over these 10 runs and reported this average alongside the corresponding average computational time in Table 2. For the pharmacokinetics inverse problem, in PINNs, we utilized the Adam optimization with Nc
= 500, learning rate (lr) of 1×10−4, and we conducted training for 50,000 iterations. Notably, in this context, the application of self-adaptive loss balancing weights was deemed unnecessary, and the two-phase training method was not employed. We perform the computational experiments for PINNs on NVIDIA’s GeForce RTX 3090 GPUs, which are powered by NVIDIA’s
2nd generation RTX Ampere architecture. The GPU has 10496 core and is endowed with 24
GB of GDDR-6X memory. PINNs parameters setup is shown in Table 3.
Since X-TFC uses a domain decomposition technique, we report the number of iterations needed from the iterative least-squares for each sub-domain, with an iteration tolerance set equal to 1e-06. The X-TFC results reported in Tables 2 and 4 are obtained with certain neural networks hyperparameters setup, which are specified so that they can be readily reproducible.
With a proper ablation study and domain decomposition, we can reduce these errors by several orders of magnitude, as shown in Tables A and B in S1 Text. The tuning hyperparameters are N number of points per sub-domain, L number of neurons, and tstep the length of each subdomain. These setups for each simulation are reported in Table 5, made with an Intel(R) Xeon
(R) W-2255 CPU @ 3.70GHz machine.
Table 2. Pharmacokinetics model: Performance of X-TFC and PINNs for parameter discovery for time range [0,
50] hours. Refer to Table 1 for X-TFC hyperparameters.
X-TFC
# data points

relative error (%)

# of iter.

comp. time [s]

6.96

5

0.07

2.57

5

0.07

0.55

0.17

5

0.07

0.38

0.11

5

0.07

# of iter.

comp. time [s]

kg

kb

10

41.66

20

9.23

50
100

PINNs
# data points

relative error (%)
kg

kb

10

1.25

0.11

5e04

48.54

20

0.21

0.03

5e04

48.74

50

0.12

0.02

5e04

48.86

100

0.09

0.02

5e04

48.38

https://doi.org/10.1371/journal.pcbi.1011916.t002

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

16 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Table 3. Pharmacokinetics model: PINN parameters setup for the discovery of unknown terms over a time range of [0, 50] hours. The initial and second numbers in the ‘Number of Iterations’ Row represent the iterations during the primary and secondary training stages using Adam optimization. The third number corresponds to the training stage utilizing L-BFGS. The first and second numbers in the ‘Architecture of Neural Networks’ indicate the width and depth, respectively.
PINNs parameters
Optimizer

Adam, LBFGS

Activation Function

Tanh

Number of Iterations

5000, 25000, 100

Architecture of main NN

50, 7

Architecture of second NN

20, 5

Learning Rate for main NN

0.001

Learning Rate for second NN

0.0001

Number of Collocation Points

500

https://doi.org/10.1371/journal.pcbi.1011916.t003

GPUs, renowned for their inherently parallel architecture, excel in efficiently distributing specific computations across a multitude of cores. As the volume of data points grows, the potential for enhanced parallelization efficiency becomes evident, potentially resulting in reduced computation times. It is worth highlighting that computational times may decrease when employing GPUs as the number of data points increases, as illustrated in Table 2 depicting the results of the PINNs method. This phenomenon is particularly noticeable due to our utilization of GPUs for this method.
In the gray-box identification test case for the Pharmacokinetics model, we aim to obtain the right-hand-side unknown term h(t) of the first ODE of the system (2). X-TFC and PINN
results and performance for a simulation of 50 hours are shown in Table 4. Performance is evaluated via Mean Absolute Error (MAE):
PN ^
jh ðtÞ hi ðtÞj
MAE ¼ i¼1 i
;
N

Table 4. Pharmacokinetics model: Unknown term discovery for time range [0, 50] hours. Comparison between X-TFC and PINNs performance via MAE, RMSE, RE, and computational time for different numbers of data points. The initial number in the ‘# of Iter.’ column for PINNs represents the iterations during the primary training stages using Adam optimization while the second number corresponds to the training stage utilizing L-BFGS.
X-TFC
h(t)

# data points
MAE

# of iter.

RMSE

RE (%)

comp. time [sec.]

10

2.57e-03

7.88e-03

36.00

2

0.003

20

1.24e-04

4.59e-04

2.89

1,1

0.015

50

3.75e-07

2.24e-06

1.99e-02

1,1,1,1

0.05

100

1.41e-08

8.69e-08

9.18e-04

1,1,1,1

0.05

# of iter.

comp. time [sec.]

PINNs h(t)

# data points
MAE

RMSE

RE (%)

10

1.26e-04

5.57e-04

6.92

3e04, 1e02

141.68

20

1.09e-04

4.82e-04

5.99

3e04, 1e02

145.97

50

6.59e-05

2.26e-04

2.80

3e04, 1e02

140.81

100

6.54e-05

1.84e-04

2.29

3e04, 1e02

143.37

https://doi.org/10.1371/journal.pcbi.1011916.t004

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

17 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Aristotle

Table 5. Pharmacokinetics model: X-TFC hyperparameters setup for parameter discovery and unknown term discovery, for time range [0, 50] hours.
Parameter discovery

Unknown terms

# data points

N

L

tstep

N

L

tstep

10

100

100

50

11

100

50

20

100

100

50

11

100

25

50

100

100

50

13

100

12.5

100

100

100

50

26

100

12.5

https://doi.org/10.1371/journal.pcbi.1011916.t005

Root Mean Squared Error (RMSE): sffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffi
PN ^
2
hi ðtÞÞ
i¼1 ðh i ðtÞ
RMSE ¼
;
N
and Relative Error (RE): qffiP
ffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffi
N
2
^
hi ðtÞÞ
i¼1 ðh i ðtÞ
ffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffiffi qffiP
RE ¼
N ^
2
i¼1 h i ðtÞ
^ and h(t) are the exact and learned solutions, respectively. Also, for these test cases, where hðtÞ
we can see how both methods can perform a good inversion of the unknown term h(t) given a few data samples. Fig 6A shows the learned concentrations in time of the three state variable B,
G, and U for X-TFC and PINNs solutions vs. the exact solution (given by 50 data points), while the learned function h(t) is plotted in Fig 6B.
As presented in Tables 2 and 4, our comparative analysis reveals valuable insights into the performance of the X-TFC and PINNs methods when applied to the same problem with varying data sizes within the same time range. For smaller sizes of the dataset (e.g., 10 data points), the PINNs method can achieve better performance in accuracy, especially for the gray-box test case, showing its inherent performance in handling sparse datasets for approximating complex functions due to the high expressivity of the deep neural network. Conversely, as the dataset size increases, the performance of the X-TFC method in terms of accuracy improves substantially. Its computational speed, a distinct advantage, allows it to effectively capitalize on larger datasets. With more data points, the X-TFC method can produce increasingly accurate results, eventually surpassing the accuracy achieved by the PINNs method. Despite the initial accuracy advantage of PINNs, it reaches a point where further increasing the dataset size does not significantly improve accuracy with the same setup while still keeping great performance. This is probably due to the optimization error, and overcoming this limitation may involve architectural enhancements, such as increasing the neural network’s depth, employing different optimization algorithms, or implementing alternative techniques. In contrast, the X-TFC method continues to benefit from additional data, showcasing its scalability and adaptability. In summary, for problems with small datasets, the PINN method excels in providing accurate solutions. For larger datasets the X-TFC method becomes increasingly competitive, offering the potential for superior accuracy with adequate computational resources.
Finally, we evaluate the performance of the two NN-based models for noisy data, simulating a more realistic scenario. We perturb 100 synthetic data points with a Uniform random distribution noise at four different levels of noise n = [1%, 2%, 3%, 4%, 5%, 10%] as follows:
^y noise ¼ ^y � ð1 þ nξÞ

PLOS Computational Biology | https://doi.org/10.1371/journal.pcbi.1011916 March 12, 2024

ð19Þ

18 / 33

PLOS COMPUTATIONAL BIOLOGY

AI-Ar

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
