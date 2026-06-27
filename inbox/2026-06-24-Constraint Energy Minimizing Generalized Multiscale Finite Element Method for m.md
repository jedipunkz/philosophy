---
source: "https://arxiv.org/abs/2205.11294v1"
title: "Constraint Energy Minimizing Generalized Multiscale Finite Element Method for multi-continuum Richards equations"
author: "Tina Mai, Siu Wun Cheung, Jun Sur Richard Park"
year: "2022"
publication: "arXiv preprint / math.NA"
download: "https://arxiv.org/pdf/2205.11294v1"
pdf: "https://arxiv.org/pdf/2205.11294v1"
captured_at: "2026-06-24T11:10:33Z"
updated_at: "2026-06-24T11:10:33Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Richard Rorty"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Constraint Energy Minimizing Generalized Multiscale Finite Element Method for multi-continuum Richards equations

- 著者: Tina Mai, Siu Wun Cheung, Jun Sur Richard Park
- 年: 2022
- 掲載情報: arXiv preprint / math.NA
- 情報源: [arxiv](https://arxiv.org/abs/2205.11294v1)
- ダウンロード: https://arxiv.org/pdf/2205.11294v1
- PDF: https://arxiv.org/pdf/2205.11294v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

In fluid flow simulation, the multi-continuum model is a useful strategy. When the heterogeneity and contrast of coefficients are high, the system becomes multiscale, and some kinds of reduced-order methods are demanded. Combining these techniques with nonlinearity, we will consider in this paper a dual-continuum model which is generalized as a multi-continuum model for a coupled system of nonlinear Richards equations as unsaturated flows, in complex heterogeneous fractured porous media; and we will solve it by a novel multiscale approach utilizing the constraint energy minimizing generalized multiscale finite element method (CEM-GMsFEM). In particular, such a nonlinear system will be discretized in time and then linearized by Picard iteration (whose global convergence is proved theoretically). Subsequently, we tackle the resulting linearized equations by the CEM-GMsFEM and obtain proper offline multiscale basis functions to span the multiscale space (which contains the pressure solution). More specifically, we first introduce two new sources of samples, and the GMsFEM is used over each coarse block to build local auxiliary multiscale basis functions via solving local spectral problems, that are crucial for detecting high-contrast channels. Second, per oversampled coarse region, local multiscale basis functions are created through the CEM as constrainedly minimizing an energy functional. Various numerical tests for our approach reveal that the error converges with the coarse-grid size alone and that only a few oversampling layers, as well as basis functions, are needed.

## PDF Text

CONSTRAINT ENERGY MINIMIZING GENERALIZED MULTISCALE FINITE
ELEMENT METHOD FOR MULTI-CONTINUUM RICHARDS EQUATIONS

arXiv:2205.11294v1 [math.NA] 23 May 2022

TINA MAIa,b , SIU WUN CHEUNGc , AND JUN SUR RICHARD PARKd,∗
Abstract. In fluid flow simulation, the multi-continuum model is a useful strategy. When the heterogeneity and contrast of coefficients are high, the system becomes multiscale, and some kinds of reduced order methods are demanded. Combining these techniques with nonlinearity, we will consider in this paper a dual-continuum model which is generalized as a multi-continuum model for a coupled system of nonlinear Richards equations as unsaturated flows, in complex heterogeneous fractured porous media; and we will solve it by a novel multiscale approach utilizing the constraint energy minimizing generalized multiscale finite element method (CEM-GMsFEM). In particular, such a nonlinear system will be discretized in time and then linearized by
Picard iteration (whose global convergence is proved theoretically). Subsequently, we tackle the resulting linearized equations by the CEM-GMsFEM and obtain proper offline multiscale basis functions to span the multiscale space (which contains the pressure solution). More specifically, we first introduce two new sources of samples, and the GMsFEM is used over each coarse block to build local auxiliary multiscale basis functions via solving local spectral problems, that are crucial for detecting highcontrast channels. Second, per oversampled coarse region, local multiscale basis functions are created through the CEM as constrainedly minimizing an energy functional. Various numerical tests for our approach reveal that the error converges with the coarse-grid size alone and that only few oversampling layers as well as basis functions are needed.
Key words. Heterogeneous fractured porous media; Unsaturated flows; Constraint energy minimizing generalized multiscale method; Model reduction; Multi-continuum; Coupled system of nonlinear Richards equations
AMS subject classifications. 65M60, 65M12

1. Introduction. For any soil sample, the amount of water retained within the gaps between unsaturated soil particles is known as soil moisture. Even being a small portion in many parts of the water cycle, soil moisture is crucial to various procedures of hydrology, biology, and biogeochemistry. For example, soil moisture is a key variable to farming, environmental management, groundwater storage, geotechnics, energy balances, meteorological forecast, and earth system dynamics, etc. Richards equation [77, 12, 11, 46, 40], which features the seepage of water into some porous material filled with water and air [29], is used as an unsaturated flow to quantitatively model the associated processes. Evaporation and precipitation, which are tightly coupled in nonlinear ways, affect moisture near the soil surface the most, prompting us to explore a coupled system of nonlinear Richards equations.
Also, in our considering porous media, there can exist complex heterogeneous rock properties, faults, intricate fracture geometries, multi-continuum background with mass transfer, high contrast, and numerous scales, among other aspects. Especially, the material characteristics of fractures can differ significantly from those of the surrounding media, which can also comprise extremely heterogeneous and large-contrast regions as well as high permeability. These obstacles lead to the fact that they can have a considerable impact on nonlinear fluid flow processes and solutions comprise multiple scales, making traditional numerical simulations much more difficult because extra computing power is needed.
The purpose of this study is to build and examine some reduced models for these types of issues. In the standard upscaling methods through homogenization, the computational domain is first partitioned into coarse-scale blocks, where scales are not necessarily resolved, then effective material property for each coarse block are calculated employing the fine-scale solutions of some local problem [31, 89]. However, it is well understood that one effective coefficient per coarse patch is insufficient to represent all features of the solutions, particularly in the regions holding important modes, fractures, high-contrast heterogeneities, and interaction of continua.
To resolve this disadvantage, we utilize on coarse grid the multi-continuum strategies [8, 6, 88, 54, 91, 75], where a number of effective medium properties are built. Physically, each continuum is treated as a system
(throughout the entire domain) so that the flow between them can be easily characterized. Different continua
Tina Mai; a Institute of Research and Development, Duy Tan University, Da Nang, 550000, Vietnam; b Faculty of Natural
Sciences, Duy Tan University, Da Nang, 550000, Vietnam; maitina@duytan.edu.vn
Siu Wun Cheung; c Center for Applied Scientific Computing, Lawrence Livermore National Laboratory, Livermore, CA
94550, USA; cheung26@llnl.gov
∗ Corresponding author: Jun Sur Richard Park ; d Department of Mathematics, The University of Iowa, Iowa City, IA, USA; junsur-park@uiowa.edu
1

are adjacent in the fine grid. They coexist via mean characteristics [8] at every location of the considering region on the coarse grid, and interactions appear among them. Mathematically, we represent on each coarse block a system of equations, each of which corresponds to one of the fine grid’s multicontinua. In this paper, using dual-continuum model for the unsaturated flows, we construct distinct Richards equations for the flow in natural fractures and the flow within matrix (background), and some specific interaction terms are coupled to such equations, as in [8, 30, 88]. This purpose is achieved by assuming that each continuum is connected to the other (even if it is not topologically connected, across the kind of coupling and the entire domain), provided that it possesses global effects solely.
To illustrate our multi-continuum strategy, we now look at dual-continuum background in further detail.
Barenblatt [8] developed the first dual-porosity model for flow simulation in fissured rock. The proposed two continua in that work are for characterizing low and high porosity continua, namely, a system of natural fractures (so-called small-scale connected, highly developed, or well-developed fractures) and a matrix, both of which are used in our paper. On the basis of [8], there was also an early work using homogenization on dual continua [6]. Intraflow and interflow transfers are together considered per continuum. Essentially, the dual-continuum background can take arbitrary shape and fit any of the above approaches.
Dual-continuum models are also utilized to represent a variety of scientific and engineering applications, such as complicated processes in shale reservoirs [4, 1, 2], where those models are employed to depict a complex interplay of organic and inorganic matter. Also, dual-continuum models can characterize flow through vugs and the rest media in vuggy carbonate reservoirs’ simulations [92, 90, 86].
The classical direct approach to tackle multi-continuum models with fractures is local fine-grid simulation, in a few simple steps [51]. First, a fine grid is built locally to represent the shapes of fractures and heterogeneities of background. Second, the flow equations are discretized on that fine grid, and a global solution is obtained from the collection of local solutions. This technique can be implemented using wellknown frameworks, such as the Finite Element Method (FEM) [7] and the Finite Volume Method (FVM)
[9, 45, 52, 66, 68, 76]. Within the confines of the finite-element framework, there are considerations of the ordinary Galerkin formulation in [7, 50, 53, 57], the mixed finite element approach in [39, 48, 60, 64], the hierarchical FEM in [73], and the discontinuous Galerkin method in [38, 47]. A hybrid strategy has also been studied [44, 65, 67], which combines the FVM for the transport equation with the FEM for the pressure equation. However, even with the aid of supercomputers and parallel computing, direct fine-grid simulation of multiple-scale problems is difficult and expensive, leading to the need of some multiscale methods.
The inspiration for the novel multiscale approach we develop in this article is the generalized multiscale finite element method (GMsFEM) [33, 16, 19, 15], which may be thought of as a generalization of the multiscale finite element method (MsFEM) [36, 32]. We will build coarse-grid multiscale basis functions that can couple multiple continua together with high-contrast channels, to achieve the small-scale impact on the large scales without the urge to solve for all minor intricacies. The GMsFEM’s primary idea is to employ local spectral decomposition in some appropriate snapshot spaces to find local dominant modes.
The resulting dominant eigenfunctions can transmit local to global properties using coarse-grid multiscale basis functions. These concepts are crucial for recognizing the effects of high-contrast regions as well as channels, which must be described separately by distinct basis functions. For instance, if we have n different connected fracture networks within a coarse block, then there are n very small eigenvalues, and the related dominant eigenvectors will reflect these connected fracture networks and can be regarded as lowered degrees of freedom that depict these fracture effects. In this way, the GMsFEM and multi-continuum techniques possess certain commonalities (see [23], for instance). A variety of domain decomposition methods [43, 56, 55]
have leveraged the idea of building local basis functions utilizing spectral problem. Recently, the GMsFEM
has been successfully applied to a variety of problems [2, 86, 23, 79, 81, 1, 3, 5, 18, 27, 26, 59].
In our previous work [71], the GMsFEM was utilized to solve dual-continuum Richards equations within complex heterogeneous fractured porous media. Nevertheless, it is not straightforward to construct a multiscale approach whose convergence is only determined by the coarse-grid size and is unaffected by scales or contrast. Several ways are discussed in the literature to generate multiscale algorithms with mesh-dependent convergence [70, 63, 69, 49, 22, 17]. The use of local spectral problems to capture the impact of high-contrast channels is motivated by the GMsFEM’s theory, and this principle is also applied to mesh-dependent convergence [49, 22, 17].
2

In that spirit, a new multiscale technique for a linear dual-continuum model was built and investigated in [14], following the constraint energy minimizing generalized multiscale finite element method (CEMGMsFEM) [22, 17, 13]. Such technique relies on the coarse-grid size alone for convergence. In our paper, for the case of multi-continuum nonlinear Richards equations, after temporal discretization, at each time step until the halting time, we employ linearization in Picard iteration (with a desired ending indicator), and the
CEM-GMsFEM [14] is applied to find the solution of these linearized equations in two-dimensional porous fractured heterogeneous media. In Appendix A, we shall additionally show that the Picard iteration process converges globally.
The CEM-GMsFEM in this paper is made up of two components: the creation of local basis functions for each coarse element (via employing the GMsFEM to build the auxiliary multiscale basis functions) and then per oversampled coarse domain (by using the CEM to achieve the collection of multiscale basis functions with locally minimal energy). More specifically, after introducing two new sources of samples, we will first establish local auxiliary multiscale basis functions over each coarse block through the GMsFEM. The number of such functions is the same as the number of high-contrast channels. These functions are dominant eigenfunctions (corresponding to the smallest eigenvalues of local spectral problems) and can be considered as the reduced degrees of freedom needed to model channelized effects. Such eigenfunctions are also important in the establishment of localized basis functions, as we point out. Second, multiscale basis functions form the other important component. For each oversampled coarse region, these functions are created by the
CEM as minimizing an energy functional, which is constrained in a manner that its minimizer satisfies a set of orthogonality requirements with respect to the auxiliary functions. It is clear from the numerical results that the error converges with the coarse-grid size only, so our multiscale approach derived by a Galerkin formulation has the mesh-dependent convergence property.
Adaptivity can also be performed in the CEM-GMsFEM, as shown in [25, 20, 21]. Some recent applications of the CEM-GMsFEM can be found in [37, 95, 87]. With the CEM-GMsFEM’s development, there have been important studies on the non-local multi-continuum (NLMC) method [94, 82, 24, 83, 84]. The
NLMC method’s main principle is similar to the CEM-GMsFEM, with an exception that the multiscale basis functions are changed to reflect the average of solutions such that the degrees of freedom possess physical meanings. In multi-continuum fractured media, these techniques are also effective in tackling high-contrast and multiscale components.
The paper is structured in the following way. In Section 2, the multi-continuum model for a coupled system of nonlinear Richards equations is introduced, within fractured heterogeneous porous media. We present in Section 3 the fine-scale discretization and Picard iteration for linearization of such system. In
Section 4, our novel multiscale approach will be provided to solve this linearized system, using a new idea of two sample sources for pressure in constructing multiscale spaces, following the constraint energy minimizing generalized multiscale finite element method (CEM-GMsFEM). In Section 5, various numerical examples will be shown to expose the approach’s mesh-dependent convergence. The paper is concluded in Section 6. We give a proof for the Picard linearization’s global convergence in Appendix A.

2. Multi-continuum Richards equations. Let Ω be a bounded, simply connected, open, Lipschitz, convex computational domain in Rd . The case d = 2 is considered to ease our discussion throughout this paper, however the method can be easily generalized to d = 3 . The subscripts i and l stand for indices
∂
and ∇ respectively represent the of continua, where N denotes the number of continua. The symbols
∂t temporal derivative and spatial gradient. Other notation is as in [72, 71]. Vector fields and matrix fields over Ω are denoted by bold letters (e.g., v and T ) while functions are represented by italic capitals (e.g., f ).
Over Ω , the spaces of functions, vector fields, and matrix fields are respectively expressed by italic capitals
(e.g., L2 (Ω)), boldface Roman capitals (e.g., V ), and special Roman capitals (e.g., S).
3

At first glance, a coupled system of dual-continuum nonlinear Richards equations [71, 79] has the form

(2.1)

∂p1 (t, x)
− div[κ1 (x, p1 (t, x))∇p1 (t, x)] + Q12 (x, p1 (t, x), p2 (t, x))(p1 (t, x) − p2 (t, x))
∂t
= f1 (t, x) in (0, T ] × Ω ,
∂p2 (t, x)
− div[κ2 (x, p2 (t, x))∇p2 (t, x)] + Q21 (x, p2 (t, x), p1 (t, x))(p2 (t, x) − p1 (t, x))
∂t
= f2 (t, x) in (0, T ] × Ω .

In this paper, we consider a general multi-continuum model of such system as follows [79, 85]: for each continuum i = 1, . . . , N ,
N

(2.2)

X
∂pi
− div(κi (pi )∇pi ) +
Qil (x, pi , pl )(pi − pl ) = fi (t, x) in (0, T ] × Ω ,
∂t l=1

where T > 0 is the final time. This system is prescribed with the initial condition pi (0, x) = pi,0 in Ω and the Dirichlet boundary condition pi (t, x) = 0 on (0, T ] × ∂Ω . Basic notation can be found in [71]. Here, pi := pi (t, x) stands for the pressure head, κi (pi ) := κi (x, pi ) denotes the unsaturated hydraulic conductivity, fi represents the source or sink function for the ith continuum, and the term Qil (x, pi , pl )(pi − pl ) describes mass transfer of the liquid which flows from the ith continuum into the lth continuum per unit of media volume as well as per unit of time [8], where we denote Qil = Qil (pi , pl ) := Qil (x, pi , pl ) . When this mass exchange term Qil (x, pi , pl )(pi − pl ) vanishes, the system (2.2) becomes a single-continuum equation (see
Section 5).
The L2 inner product is represented by (·, ·) , and the Sobolev space V := H01 (Ω) = W01,2 (Ω) is equipped with the norm k · kV :

 21
.
kvkV = kvk2L2 (Ω) + k∇vk2L2 (Ω)
Here, k∇vkL2 (Ω) := k|∇v|kL2 (Ω) , where |∇v| indicates the Euclidean norm of the d-component vector-valued function ∇v . We also denote V = V N = [H01 (Ω)]N . With v = (v1 , · · · , vN ) ∈ V , k∇vkL2 (Ω) := k|∇v|kL2 (Ω) , where |∇v| stands for the Frobenius norm of the N × d matrix ∇v .
The hydraulic conductivity together with its spatial gradient as well as the mass transfer coefficient are assumed to be uniformly bounded, that is, positive constants κ, κ and β, β exist so that the following inequalities are satisfied:
(2.3)

κ ≤ κi (x, pi ), |∇κi (x, pi )| ≤ κ ,
β ≤ Qil (x, pi , pl ) ≤ β .

Without loss of generality, the initial condition is assumed to be
(2.4)

p0 = p(0, x) = (p1,0 , . . . , pN,0 ) ∈ V .

Given u = (u1 , . . . , uN ) ∈ V , with i = 1, 2, . . . , N , we define the following bilinear forms: for all p =
(p1 , · · · , pN ) , v = (v1 , · · · , vN ) ∈ V ,
Z
(2.5)
ai (pi , vi ; ui ) =
κi (ui )∇pi · ∇vi dx ,
(2.6)

qi (p, vi ; u) =

Ω
N
XZ
l=1

Qil (ui , ul )(pi − pl )vi dx .

Ω

The variational form of (2.2) reads: find p = (p1 , . . . , pN ) ∈ V such that with i = 1, . . . , N ,


∂pi
(2.7)
, vi + ai (pi , vi ; pi ) + qi (p, vi ; p) = (fi , vi ) ,
∂t for any v = (v1 , · · · , vN ) ∈ V , with a.e. t ∈ (0, T ] , and fi (t, ·) ∈ L2 (Ω) . The initial condition is given in
(2.4).
4

3. Fine-scale discretization and Picard iteration for linearization. To tackle our problem’s nonlinearity, we take advantage of an efficient Picard iterative scheme, as described in [61, 79, 42, 62]. Over this section, such an iteration algorithm is presented for time-dependent multi-continuum systems.
To achieve the system (2.7)’s first goal of temporal discretization (see [61, 79, 12], for instance), we will use the following conventional backward Euler finite-difference algorithm: find ps = (p1,s , . . . , pN,s ) ∈ V
such that for all v = (v1 , . . . , vN ) ∈ V ,


pi,s+1 − pi,s
(3.1)
, vi + ai (pi,s+1 , vi ; pi,s+1 ) + qi (ps+1 , vi ; ps+1 ) = (fi,s+1 , vi ) ,
τ
where we divide the temporal domain [0, T ] equally into S intervals, having τ = T /S > 0 as the size of time step, and the subscript s signifies the value of a function at the time point ts = sτ (with s = 0, 1, · · · , S).
Following that, the nonlinearity in space will be linearized using Picard iteration (see [61, 79, 42], for instance). At the (s + 1)th temporal step, p0s+1 ∈ V is guessed. Given pns+1 ∈ V , with n = 0, 1, 2, · · · , we seek pn+1
s+1 ∈ V such that for all v = (v1 , . . . , vN ) ∈ V ,
(3.2)

pn+1
i,s+1 − pi,s
, vi
τ

!
n+1
n
+ ai (pi,s+1
, vi ; pni,s+1 ) + qi (pn+1
s+1 , vi ; ps+1 ) = (fi,s+1 , vi ) .

As proved in [74], there exists a unique solution pn+1
i,s+1 to this linearized system (3.2).
When n tends to ∞ , the Picard iterative procedure converges to a limit (see Appendix A for a theoretical proof). In simulation, we end this procedure at an αth iteration when it satisfies a specific halting indicator, leading to the previous time data
(3.3)

ps+1 = pα
s+1 ,

in order to move on to the next time step in (3.1). A terminating criterion is proposed over this paper employing the relative successive difference, that is, provided a user-defined tolerance δ0 > 0, if
(3.4)

n kpn+1
i,s+1 − pi,s+1 kL2 (Ω)
≤ δ0 , kpni,s+1 kL2 (Ω)

for i = 1, . . . , N , then the iterative process is stopped.
Now, the fine-grid notation is considered. First, to begin discretizing the variational problem (2.7), we let Th be a fine grid of size h , which is assumed to be very small. With this assumption of h , the fine-grid solution will be sufficiently close to the exact solution. Second, with respect to the rectangular fine grid Th , we define Vh as the H01 (Ω)-conforming finite element basis space of piecewise bilinear functions:
(3.5)

Vh := {v ∈ V : v|K ∈ Q1 (K) ∀K ∈ Th } ,

where the space Q1 (K) consists of all bilinear elements (or multilinear d-elements when d > 2) over K . We let V h = VhN and denote the [L2 (Ω)]N projection operator onto V h by Ph .
On the fine scale, the completely discrete Picard iterative algorithm is as follows: beginning with an initial ph,0 = Ph p0 ∈ V h having p0 from (2.4), at the temporal step (s + 1)th, we make a guess p0h,s+1 ∈ V h and perform iteration from (3.2) in V h :
!
pn+1
i,h,s+1 − pi,h,s n+1
n n
(3.6)
, vi + ai (pn+1
i,h,s+1 , vi ; pi,h,s+1 ) + qi (ph,s+1 , vi ; ph,s+1 ) = (fi,s+1 , vi ) ,
τ
where n = 0, 1, 2, · · · , until reaching (3.4) at an αth Picard step, for all v = (v1 , . . . , vN ) ∈ V h . To proceed to the next time step in (3.1), we utilize (3.3) to set the previous time data
(3.7)

ph,s+1 = pα
h,s+1 .
5

4. CEM-GMsFEM for coupled multi-continuum nonlinear Richards equations. Following
[34, 42, 14, 71], we now establish a new strategy for coupled multi-continuum nonlinear Richards equations
(2.2) in complex heterogeneous fractured porous media, using the constraint energy minimizing generalized multiscale finite element method (CEM-GMsFEM). More specifically, in the pressure computation for the equivalently nonlinear system (2.7), we will show the establishment of auxiliary space (utilizing a novel idea of two sample sources) and multiscale space. To properly construct such CEM-GMsFEM, after temporal discretization of (2.7), the linearized formulation (3.2) can be used to consider the nonlinearity as a constant at each Picard iterative step. Multiscale space can therefore be built according to this nonlinearity.
4.1. Overview. First, we will go over the concepts of coarse and fine grids. The start is partitioning
Ω into finite elements, where multiscale characteristics are not necessarily resolved. The partition is named coarse grid T H , and Th is its refinement. In T H , a generic element K is named a coarse-grid block (also known as coarse element or coarse patch). Moreover, we call H > 0 the coarse-grid size, where H  h . Let
Nc be the number of coarse blocks and Nv be the number of coarse-grid nodes. The collection of all coarse v
nodes (vertices) is denoted by {xk }N
k=1 . Figure 1 illustrates the fine grid, coarse grid, as well as a coarse block K .

K

Figure 1: Illustration of the fine grid, coarse grid, and a coarse block K .

For some subdomain D ⊂ Ω , the restrictions of V and V on D are respectively denoted by V (D)
and V (D) . Furthermore, the subspace of V (D) containing functions with zero trace on ∂D is represented by V 0 (D) . Using this definition on a coarse block Kj , provided u = (u1 , · · · , uN ) ∈ V (Kj ) , for all p =
(p1 , · · · , pN ) , v = (v1 , · · · , vN ) ∈ V (Kj ) , i = 1, 2, · · · , N , we define local bilinear forms by
Z
(j)
ai (pi , vi ; ui ) =
κi (ui )∇pi · ∇vi dx ,
Kj

(4.1)

(j)

qi (p, vi ; u) =

N Z
X
l=1

(j)
ri (pi , vi ; ui ) =

Qil (ui , ul )(pi − pl )vi dx ,

Kj

Z
Kj

κ
e i (ui )pi vi dx .

Here,
κ
e i (ui ) = κi (ui )

Nv
X
k=1

6

|∇χk,i |2 ,

in which for the coarse node xk , each χk,i (with linear boundary conditions for cell problems [35]) is a conventional multiscale finite element basis function within the ith continuum. In the current context,
H
v
{χk,i }N
k=1 is a collection of bilinear partition of unity functions (for T ) supported in the continuum ith. On each coarse block Kj , the coupled local bilinear forms are defined as follows: for any p = (p1 , · · · , pN ), v =
(v1 , · · · , vN ) ∈ V (Kj ), i = 1, · · · , N ,
X (j)
a(j) (p, v; u) =
ai (pi , vi ; ui ), i

q

(j)

X (j)
(p, v; u) =
qi (p, vi ; u),

r

(j)

X (j)
(p, v; u) =
ri (pi , vi ; ui ),

(4.2)

i

i
(j)
(j)
aQ (p, v; u) = a (p, v; u) + q (j) (p, v; u).

We also define the bilinear forms aQ and r by aQ (p, v; u) =

Nc
X
(j)
aQ (p, v; u) , j=1

(4.3)
r(p, v; u) =

Nc
X

r(j) (p, v; u) .

j=1

In multiscale space, our major target is seeking for (2.7) a multiscale solution pms that approximates the fine-scale solution ph better than using the GMsFEM ([10]). For this purpose, the CEM-GMsFEM is utilized to attain the multiscale solution pms . Two levels are required to build the multiscale space. First, through the GMsFEM, an auxiliary space is generated. Second, a multiscale space V ms is established (benefiting from that auxiliary space) and possesses multiscale basis functions with locally minimal energy over some subregions. Ultimately, a multiscale solution can be found using these multiscale basis functions. Note that the obtained V ms is stable throughout this procedure when utilizing either the first or second source of samples us from Subsection 4.2. We refer the readers to [71, 72, 33, 34, 25, 20, 16, 41] and [22, 21, 14, 42] for more information about the GMsFEM and CEM-GMsFEM, respectively. At the first stage, the GMsFEM
will be used to design our auxiliary multiscale basis functions as follows.
4.2. Auxiliary multiscale basis functions. Given a set of samples {ub }B
b=1 ⊂ V and weights
+
{wb }B
⊂
R
.
The following bilinear forms are determined from weighted
Monte
Carlo integration [93]: b=1
(j)

AQ (p, v) =

B
X

(j)

wb aQ (p, v; ub ),

b=1

(4.4)
R(j) (p, v) =

B
X

wb r(j) (p, v; ub ).

b=1

Next, we will create our local auxiliary multiscale basis functions employing the GMsFEM. In particular,
(j)
these coupled functions are identified by a local spectral problem, that is, to seek a real number λk and an
(j)
associated function φk ∈ V (Kj ) such that
(j)

(j)

(j)

(j)

AQ (φk , v) = λk R(j) (φk , v) for any v ∈ V (Kj ) .

(4.5)
(j)

The eigenfunctions φk of (4.5) are normalized in the norm produced by the inner product R(j) as follows:
(4.6)

(j)

(j)

R(j) (φk , φk ) = 1 .
7

(j)

The eigenvalues λk of (4.5) are organized in nondecreasing order over k . Then, using the first corresponding
Lj eigenfunctions, we generate the following local auxiliary multiscale space:
(4.7)

(j)

V (j)
aux = V aux (Kj ) = span{φk : 1 ≤ k ≤ Lj } .

The sum of such local auxiliary multiscale spaces represents the global auxiliary multiscale space:
(j)
c
V aux = ⊕N
j=1 V aux .

(4.8)

Also, the global bilinear forms AQ and R are defined by
Nc
X

AQ (p, v) =

(j)

AQ (p, v) ,

j=1

(4.9)

Nc
X

R(p, v) =

R(j) (p, v) .

j=1

We briefly mention two sources of samples {ub }B
b=1 which will be considered in our numerical experiments.
The first source of samples is simply from choosing a single realization p?h ∈ V h , that is, B = 1 where w1 = 1
in (4.4), and u1 = p?h is the steady-state FEM solution (3.7). This first sample source will be used for steadystate cases.
The second source of samples is from taking realizations of the sink or source function in (2.2) and the backward Euler temporal discretization (3.1) within the fully Picard discrete scheme (3.6) for solving (2.2), to obtain from (3.7) a numerical fine-grid approximation p?h,s ∈ V h at the time step s (where t = sτ ). In this case, we have B = S + 1, and ub = p?h,b−1 .

(4.10)
By setting the weights
(4.11)

(
1
wb =
1/2

if b = 2, 3, . . . , S , if b = 1, S + 1 ,

the weighted integration (4.4) can be regarded as numerical integration of piecewise linear functions using trapezoidal rule in the temporal variable. That is,
Z T
(j)
(j)
AQ (p, v) =
aQ (p, v; p?h ) dt ,
0
(4.12)
Z T
R(j) (p, v) =

r(j) (p, v; p?h ) dt,

0

where p?h is the piecewise linear Lagrange interpolation of {p?h,s }Ss=0 on the temporal grid {sτ }Ss=0 . This second sample source will be applied to the time-dependent cases. In general, one can utilize multiple realizations in a similar manner for both sample sources.
4.3. Multiscale space. To define our multiscale basis functions for spanning the solution space, we
(j)
introduce the concept of φ-orthogonality. Providing an auxiliary basis function φk ∈ V aux within a coarse
(j)
block Kj , we state that ψ ∈ V is φk -orthogonal if for 1 ≤ k 0 ≤ Lj 0 and 1 ≤ j 0 ≤ Nc ,


(j 0 )
(4.13)
R ψ, φk0
= δk,k0 δj,j 0 , equivalently,
(4.14)



(j)
R ψ, φk = 1 ,



(j 0 )
R ψ, φk0
=0
8

with k 0 6= k or j 0 6= j .

(j)

This φk -orthogonality gives rise to the orthogonal projection operator π : [L2 (Ω)]N → V aux proposed by
Nc
X
π=
πj , where πj : [L2 (Kj )]N → V (j)
aux is defined as j=1

(4.15)

πj (v) =

Lj
(j)
X
R(j) (v, φk )

(j)

(j) (φ(j) , φ(j) )
k=1 R
k k

φk for any v ∈ [L2 (Kj )]N .
(j)

Our global multiscale basis functions are now being built. For each auxiliary function φk ∈ V aux , the solution to the following constrained energy minimization problem specifies the global multiscale basis
(j)
function ψ k ∈ V : n
o
(j)
(j)
(4.16)
ψ k = argmin AQ (ψ, ψ) : ψ ∈ V is φk -orthogonal .
(j)

(j)

The variational form of this minimization problem (4.16) is as follows: determine ψ k ∈ V and µk ∈ V aux such that
(j)

(j)

AQ (ψ k , v) + R(v, µk ) = 0 for all v ∈ V ,

(4.17)

(j)

(j)

R(ψ k − φk , ν) = 0 for all ν ∈ V aux .
We identify our localized multiscale basis functions as a result of the establishment of global multiscale basis functions. An oversampled domain is created by extending the coarse grid block Kj by m coarse-grid layers, for each Kj ∈ T H . Figure 2 depicts an example of an oversampled region. The solution of the following
(j)
constrained energy minimization problem defines the localized multiscale basis function ψ k,ms ∈ V 0 (Kj,m ) : n
o
(j)
(j)
ψ k,ms = argmin AQ (ψ, ψ) : ψ ∈ V0 (Kj,m ) is φk -orthogonal .

(4.18)

Now, let Zj be the set of indices such that if z ∈ Zj , then Kz ⊂ Kj,m ; and let n
o
X
(z)
V aux (Kj,m ) =
V aux (Kz ) = span φk : 1 ≤ k ≤ Lz , z ∈ Zj .
z∈Zj
(j)

Then, the following variational problem is equivalent to the minimization problem (4.18): find ψ k,ms ∈
(j)

V 0 (Kj,m ) and µk,ms ∈ V aux (Kj,m ) such that
(j)

(4.19)

(j)

AQ (ψ k,ms , v) + R(v, µk,ms ) = 0 for all v ∈ V 0 (Kj,m ),
(j)

(j)

R(ψ k,ms − φk , ν) = 0 for all ν ∈ V aux (Kj,m ) .
Using the localized multiscale basis functions, we generate the multiscale finite element space
(4.20)

(j)

V ms = span{ψ k,ms : 1 ≤ k ≤ Lj , 1 ≤ j ≤ Nc } .

Remark that outside of some local (oversampled) subdomains, the global multiscale basis functions (4.16)
(which are generally supported in the whole domain) exponentially decay [22]. This characteristic is critical in the CEM-GMsFEM’s convergence investigation, demonstrating the utilization of local multiscale basis functions (4.18) in V ms [14].
4.4. CEM-GMsFEM for coupled system of nonlinear Richard equations. We note first that the space V = H 10 (Ω) is continuous throughout the previous Sections and Subsections 4.1–4.3. The multiscale space V ms ⊂ V therefore requires some finite dimensional analogues for simulations. As a result, the problem under consideration is tackled via the fine grid together with a suitable finite element method
[21, 42] in our numerical computations.
9

Kj,1
Kj

Figure 2: Illustration of an oversampled region Kj,1 established via extending Kj by 1 coarse-grid layer.

Recalling ph,0 = Ph p0 ∈ V h with p0 from (2.4), we now have an initial pms (0, ·) = pms,0 that satisfies
(4.21)

ai (pi,h,0 − pi,ms,0 , vi ; pi,h,0 ) + qi (ph,0 − pms,0 , vi ; ph,0 ) = 0 ,

for any v = (v1 , · · · , vN ) ∈ V ms . Fixing the (s + 1)th temporal step, our strategy (as in [71, 41, 42]) is to solve the problem (2.7) through linearization relied on Picard’s iterative technique. This can be done by employing at each Picard iteration the CEM-GMsFEM (in Subsections 4.2 and 4.3) with the constructed stable offline multiscale space V ms (proposed at the close of Subsection 4.1).
In particular, all along the online stage, the full model reduction approach is as follows: beginning with pms,0 ∈ V ms from (4.21), we pick a guess p0ms,s+1 ∈ V ms at the temporal step (s + 1)th and iterate from
(3.2) in V ms :
!
pn+1
i,ms,s+1 − pi,ms,s n+1
n
, ψi + ai (pi,ms,s+1
(4.22)
, ψi ; pni,ms,s+1 ) + qi (pn+1
ms,s+1 , ψi ; pms,s+1 ) = (fi,s+1 , ψi ) ,
τ
where ψ = (ψ1 , · · · , ψN ) ∈ V ms and n = 0, 1, 2, · · · , until getting to (3.4) at some Picard step αth. To move to the next time step in (3.1), we use (3.3) for selecting the previous temporal data
(4.23)

pms,s+1 = pα
ms,1 .

Remark 4.1. The CEM-GMsFEM for single-continuum cases are treated similarly to the multicontinuum cases in this Section by allowing the transfer term to vanish in (2.7) and in all related expressions.
Moreover, the steady-state cases are handled similarly to the time-dependent cases in this Section by letting the second source of samples (4.10) for Subsection 4.2 to be the first source of samples (3.7) as the steadystate FEM solution. After we choose an appropriate source of samples, the snapshot functions as well as the basis functions are time-independent. Relying on specific case, Eq. (4.22) is the corresponding equation in Section 3 (time-dependent dual-continuum case (3.6), time-dependent single-continuum case, steady-state dual-continum case, and steady-state single-continum case), where the subscript “h” is replaced by “ms”.
5. Numerical examples. We will give various numerical tests in this section to demonstrate our approach’s performance. In each test, the effect of coarse-grid size H is investigated. For all experiments, we used the fine-grid size h = 1/128 and the number of oversampling layers m ≈ 10 log(1/H)/ log(64). Within each experiment, the number of local multiscale basis functions is fixed throughout all coarse elements. We compare the solutions obtained by our strategy employing the constraint energy minimizing generalized
10

multiscale finite element method (CEM-GMsFEM, abbreviated by CEM) with the solutions computed by the finite element method (FEM).
The spatial domain is Ω = [0, 1]2 . For time-dependent equations in the temporal interval [0, T ] , the stopping time is T = Sτ = 2 while the temporal step size is τ = 1/10 , so there are S = 20 time steps.
We consider the channelized media and consequently deal with the high-contrast coefficients defined in the spatial domain Ω . The Picard iteration’s halting indicator is δ0 = 10−5 , which guarantees the convergence of this linearization procedure. The continua are assumed to be isotropic (and the anisotropic case is treated in the same way). Then, hydraulic conductivity tensors can be considered as scalar functions κi (x, pi )
(multiplying with the identity matrix [28]), for i = 1, 2 .
5.1. Experiements for single-continuum Richards equations. In this section, we examine steadystate and time-dependent single-continuum Richards equations from (2.2), respectively: find p ∈ V such that


(5.1)
− div κ(x)ep(x) ∇p(x) = 1 in Ω , and
(5.2)

∂p(t, x)
− div[κ(x)ep(t,x) ∇p(t, x)] = f (t, x) in (0, T ] × Ω ,
∂t

where f (t, x) = sin(πx1 ) sin(πx2 ) , and κ(x, p) = κ(x)ep(t,x) . Both problems have the zero Dirichlet boundary condition, and the initial condition for (5.2) is p(0, x) = 0 . The permeability field κ(x) in (5.1) and (5.2)
is depicted in Figure 3. The value of κ(x) in the yellow regions (channels) is 1000 and in the blue region is
10 .

Figure 3: Permeability field κ(x) for problems (5.1) and (5.2).

We compute the CEM solutions pms from (4.22) as well as their respective FEM references ph in (3.6) for
Eqs. (5.1) and (5.2). Note that the above single-continuum equations do not have the interaction terms, thus the corresponding terms in (3.6) and (4.22) are ignored when following the algorithms to compute the CEM
and FEM solutions. Further, for steady-state case (5.1), we do not need to involve the iteration with respect to time. The initial guess for Picard iteration of problem (5.1) is the identically zero function in the domain and of problem (5.2) is the previous time data ((3.7) for FEM and (4.23) for CEM). To construct multiscale basis functions, the sampling method described in Subsection 4.2 is utilized. We employed the second source of samples (4.10) based on (3.7) for time-dependent case and used the steady-state FEM solution ph (3.7)
as the first source of samples for the steady-state case.
11

In both experiments, the relative L2 and H 1 errors between our CEM solutions pms and the FEM
references ph are defined as
(5.3)

epL2 =

||∇pms − ∇ph ||L2 (Ω)
||pms − ph ||L2 (Ω) p
, eH 1 =
.
||ph ||L2 (Ω)
||∇ph ||L2 (Ω)

We investigate these errors with respect to the coarse-grid size H, by different number of local multiscale basis functions. Table 1 presents the errors for the steady-state equation (5.1), and Table 2 shows the errors in the time-dependent case (5.2). We note that the total number of degrees of freedom for our multiscale method
(dim(Vms )) relates entirely to the coarse-grid size H and the number of local multiscale basis functions. In all tables, we observe that the numerical approximations are very accurate for every choice of coarse-grid size H. Also, H 1 errors less than 3% and L2 errors less than 0.5% even with relatively large coarse-grid size
H = 1/4, where only 96 degrees of freedom are utilized for CEM at maximum throughout the experiments.
This number is much less than 16129, the total number of degrees of freedom used in FEM. It is explicit from those tables that as the sequence of coarse-grid sizes H converges, the sequence of CEM solutions converges.
According to the tables, both H 1 and L2 errors can be further decreased once more local multiscale basis functions and oversampling layers are involved. However, a too large number of multiscale basis functions has a direct impact on the method’s computational complexity. It is unknown whether the contrast has a direct effect on the number of required multiscale basis functions. Figure 4 illustrates the comparison of the plots of solutions to (5.2) at the final time T = 2, computed by the CEM and FEM when H = 1/16. One can see that each solution obtained by CEM almost coincides with its reference solution computed by FEM.
H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
2.0760%
1.3022%
0.7354%
0.2829%

dim(Vms )
64
256
1024
4096

L2 error
0.3446%
0.1188%
0.0389%
0.0073%

(a) Errors with 4 local basis functions for (5.1).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
1.6471%
0.9370%
0.4682%
0.1736%

dim(Vms )
80
320
1280
5120

L2 error
0.2410%
0.0885%
0.0219%
0.0041%

(b) Errors with 5 local basis functions for (5.1).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
1.4018%
0.6085%
0.2779%
0.0707%

dim(Vms )
96
384
1536
6144

L2 error
0.1914%
0.0452%
0.0112%
0.0012%

(c) Errors with 6 local basis functions for (5.1).

Table 1: Relative L2 and H 1 errors for (5.1) with different number of local basis functions; dim(Vh ) = 16129 .

5.2. Experiments for dual-continuum Richards equations. Benefiting from [71], we consider the following steady-state problem of the form (2.2) in the domain Ω : find p = (p1 , p2 ) ∈ [H01 (Ω)]2 such that


κ1 (x)
10
− div
∇p1 +
(p1 − p2 ) = 1 ,
1 + |p1 |
1 + |p1 |


(5.4)
κ2 (x)
10
− div
∇p2 +
(p2 − p1 ) = −1 ,
1 + |p2 |
1 + |p2 |
12

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
2.7544%
1.3024%
0.7187%
0.2687%

dim(Vms )
64
256
1024
4096

L2 error
0.4359%
0.1261%
0.0376%
0.0068%

(a) Errors with 4 local basis functions for (5.2).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
2.0580%
1.0094%
0.4816%
0.1659%

dim(Vms )
80
320
1280
5120

L2 error
0.2887%
0.0962%
0.0226%
0.0038%

(b) Errors with 5 local basis functions for (5.2).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
1.8294%
0.7953%
0.2685%
0.0701%

dim(Vms )
96
384
1536
6144

L2 error
0.2391%
0.0676%
0.0106%
0.0011%

(c) Errors with 6 local basis functions for (5.2).

Table 2: Relative L2 and H 1 errors for (5.2) with different number of local basis functions; dim(Vh ) = 16129 .

a. pms (T, x) by CEM.

b. ph (T, x) by FEM.

Figure 4: Problem (5.2): Solutions p(T, x) obtained by CEM and FEM when H = 1/16 .

where it has zero Dirichlet boundary condition, and the configurations of high-contrast permeability fields
κ1 and κ2 are shown in Figure 5. The values in the yellow regions (channels) are higher than the values in the blue regions, and κ1 (x) ≥ κ2 (x) for all x ∈ Ω .
As a special case of (2.2) (with the given conditions there), the following problem is also considered.
That is, we investigate the Gardner-Basha model, which utilizes more intricate right-hand side functions and includes both sources and sinks (remark that the van Genuchten-Mualem model can work as well)
[58, 78, 71]. Employing the Gardner-Basha model in [58], we solely consider the unsaturated hydraulic conductivity’s nonlinearity, and the volumetric water content function is assumed to be identity. In the
13

given domain [0, T ] × Ω , we seek solution p = (p1 , p2 ) ∈ [H01 (Ω)]2 of the system

(5.5)

∂p1
102
− div (κ1 (x)Kr (p1 )∇p1 ) +
(p1 − p2 ) = f1 (x) in (0, T ] × Ω ,
∂t
1 + |p1 |
∂p2
102
− div (κ2 (x)Kr (p2 )∇p2 ) +
(p2 − p1 ) = f2 (x) in (0, T ] × Ω ,
∂t
1 + |p2 |

with the Dirichlet boundary condition pi (t, x) = 0 on (0, T ] × ∂Ω , and with the initial condition pi (0, x) = 0
in Ω . Here, the expression of relative hydraulic conductivity Kr is
Kr (p) = e−αG |p| ,

(5.6)

where αG is parameter characteristic of the soil pore size distribution. The geometric mean of αG is assumed to be 0.1 . Fig. 6 describes the high-contrast κi (x) for i = 1, 2 . We choose in the blue regions κ1 (x) = 10
and κ2 (x) = 1 , as well as in the yellow regions κ1 (x) = 104 and κ2 (x) = 10 . The specific source and sink functions are respectively provided by f1 (x) = ex1 +x2 , f2 (x) = −ex1 +x2 .

a. κ1 (x) .

b. κ2 (x) .

Figure 5: Problem (5.4): κ1 (x) = 104 , κ2 (x) = 10 in the corresponding yellow regions (channels); κ1 (x) = 10,
κ2 (x) = 0.5 in the associated blue regions.
We compute the numerical CEM solutions pms of (5.4) and (5.5) (based on (4.22)) as well as their respective FEM references ph using (3.6) without having to consider the time-step iterations for (5.4). The following relative L2 and H 1 errors are between the CEM solutions pms and their FEM references ph :
(5.7)

epL2 =

||pms − ph ||L2 (Ω) p
||∇pms − ∇ph ||L2 (Ω)
, eH 1 =
.
||ph ||L2 (Ω)
||∇ph ||L2 (Ω)

Table 3 and 4 present the relative L2 and H 1 errors for problems (5.4) and (5.5), respectively. We note that the total number of degrees of freedom of our multiscale method (dim(Vms )) depends on the coarse-grid size
H and the number of local multiscale basis functions. Those tables show clearly that the errors converge once the coarse-grid size H is refined. Also, according to the tables, increasing the number of local multiscale basis functions and oversampling layers help reduce the errors. With H = 1/4 , the errors are relatively large especially for the time-dependent problem (5.5), but they can be lower once more local multiscale basis functions are used. For relatively small coarse-grid size H, the error convergence tend to stagnate as we already have enough number of total degrees of freedom based on Table 3 and 4. We observe that for small coarse-grid size, only few number of local multiscale basis functions are needed. For the first continuum and at the final time T = 2 , Figure 7 plots the solutions p1 (T, x) of (5.5), obtained by the CEM and FEM when
H = 1/16 . Both solutions are almost identical throughout the entire domain.
14

a. κ1 (x) .

b. κ2 (x) .

Figure 6: Problem (5.5): κ1 (x) = 104 , κ2 (x) = 10 in their yellow regions (channels); κ1 (x) = 10, κ2 (x) = 1
in their blue regions.

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
5.6621%
2.2713%
1.0382%
0.4093%

dim(Vms )
64
256
1024
4096

L2 error
1.2575%
0.2370%
0.0736%
0.0135%

(a) Errors with 4 local basis functions for (5.4).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
5.4959%
2.0411%
1.0442%
0.3896%

dim(Vms )
80
320
1280
5120

L2 error
1.1861%
0.2042%
0.0709%
0.0109%

(b) Errors with 5 local basis functions for (5.4).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
4.6551%
1.9344%
1.0086%
0.3736%

dim(Vms )
96
384
1536
6144

L2 error
0.9401%
0.1860%
0.0647%
0.0092%

(c) Errors with 6 local basis functions for (5.4).

Table 3: Relative L2 and H 1 errors for (5.4) with different number of local basis functions; dim(Vh ) = 16129 .

6. Conclusions. We present in this paper a methodology for handling issues from coupled system of multi-continuum nonlinear Richards equations, in complex fractured heterogeneous porous media, utilizing the constraint energy minimizing generalized multiscale finite element method (CEM-GMsFEM). The basic concept is to discretize this system temporally via the implicit backward Euler method, then linearize it spatially by Picard iteration (with the required convergence indicator) at each time step until the ending time. The CEM-GMsFEM is used in each Picard iteration to systematically create multiscale basis functions
(with locally minimal energy) for pressure. In order to do so, we propose two new sources of samples and solve
15

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
15.7431%
6.4665%
1.7532%
0.6889%

dim(Vms )
64
256
1024
4096

L2 error
5.7543%
1.0124%
0.1586%
0.0285%

(a) Errors with 4 local basis functions for (5.5).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
13.7012%
4.5603%
1.6636%
0.6465%

dim(Vms )
80
320
1280
5120

L2 error
3.7477%
0.6151%
0.1481%
0.0223%

(b) Errors with 5 local basis functions for (5.5).

H
1/4
1/8
1/16
1/32

m
3
5
7
8

H 1 error
10.7440%
2.9378%
1.5865%
0.6125%

dim(Vms )
96
384
1536
6144

L2 error
2.5620%
0.3349%
0.1317%
0.0175%

(c) Errors with 6 local basis functions for (5.5).

Table 4: Relative L2 andf H 1 errors for (5.5) with different number of local basis functions; dim(Vh ) = 16129 .

a. p1,ms (T, x) by CEM.

b. p1,h (T, x) by FEM.

Figure 7: Problem (5.5): Solutions p1 (T, x) obtained by CEM and FEM when T = 2 .

local spectral problems via the GMsFEM to first build the local auxiliary multiscale basis functions, which are crucial for determining high-contrast channels. Second, employing the CEM through some constraints connected to the auxiliary functions, we solve an energy minimizing problem by oversampling technique, to establish localized multiscale basis functions. Our numerical results exhibits that the error converges with the coarse-grid size alone, and the method is very accurate. Appendix A provides a theoretical proof for global convergence of the Picard iteration process.
16

Acknowledgements.
Tina Mai’s research was supported by RFBR and VAST under grant 21-51-54001 and by Duy Tan
University under decision 5390/QD-DHDT.
This work was performed under the auspices of the U.S. Department of Energy by Lawrence Livermore
National Laboratory under Contract DE-AC52-07NA27344 and LLNL-JRNL-833749.
Appendix A. Global convergence of Picard linearization procedure.
We will prove the global convergence of Picard linearization process (given in Section 3) for the system
(3.6) relied on (3.1) and originated from (2.1), with i, l = 1, 2 in this appendix, following [61] (and thanks to J.
Batista and A. Mazzucato, personal communication, January 9, 2022). The generalization to i, l = 1, · · · , N
is proved similarly.
In (2.2) (and thus (2.7), (3.1) and (3.6)), each hydraulic conductivity coefficient κi satisfies the assumption (2.3), that is, 0 < κ1 (x, p1 ), κ2 (x, p2 ) ≤ κ for some positive constant κ . Each function κi is globally Lipschitz continuous with Lipschitz constant Lκi (without any dependence on t and x). Let p = (p1 , p2 ) , pil = (pi , pl ) , then the function Qil (pil ) := Qil (x, pi , pl ) is nonlinear yet globally Lipschitz continuous with the Lipschitz constant LQil for every i, l = 1, 2, i 6= l (without any dependence on t and x).
Furthermore, we suppose that each pi is positive and that Qil is uniformly bounded above by some constant
βQil (≤ β as from (2.3)). The subscripts (s + 1) and h are eliminated from the Picard iteration (3.6) for simplicity.
Then in V h , we obtain the following equality by subtracting (3.1) from (3.6) and picking correct vi =
pn+1
− pi,s+1 : i
(A.1)

1 n+1
kp
− pi,s+1 k2 + ai (pn+1
, pn+1
− pi,s+1 ; pni ) − ai (pi,s+1 , pn+1
− pi,s+1 ; pi,s+1 )
i i
i
τ i
= −qi (pn+1 , pn+1
− pi,s+1 ; pn ) + qi (ps+1 , pn+1
− pi,s+1 ; ps+1 ) .
i i

The left-hand side of (A.1) is indicated by the following LS using (2.5):
1 n+1
kp
− pi,s+1 k2 + (κ(pni )∇pn+1
, ∇(pn+1
− pi,s+1 ))
i i
τ i
− (κ(pi,s+1 )∇pi,s+1 , ∇(pn+1
− pi,s+1 ))
i
1 n+1
= kpi − pi,s+1 k2 + (κi (pni )∇(pn+1
− pi,s+1 ), ∇(pn+1
− pi,s+1 ))
i i
τ
+ ((κi (pni ) − κi (pi,s+1 ))∇pi,s+1 , ∇(pn+1
− pi,s+1 )) .
i

LS =
(A.2)

The right-hand side of (A.1) is represented by RS employing (2.6) as follows:
X
RS =
(−Qil (pnil ) · (pn+1
− pn+1
), pn+1
− pi,s+1 )
i i
l l

(A.3)

+(Qil (pil,s+1 ) · (pi,s+1 − pl,s+1 ), pn+1
− pi,s+1 )
i
X
n+1
n+1
n
=
(−Qil (pil ) · ((pi − pl ) − (pi,s+1 − pl,s+1 )), pn+1
− pi,s+1 )
i l

+((−Qil (pnil ) + Qil (pil,s+1 )) · (pi,s+1 − pl,s+1 ), pn+1
− pi,s+1 ) .
i
For i, l = 1, 2 , i 6= l , one notes that
(A.4)

kpni − pnl k ≤ k(pni , pnl )k = kpn k ,

kpnl − pl,s+1 k ≤ kpn − ps+1 k ≤

X

kpnl − pl,s+1 k .

l

Assume there are respectively sufficiently large and small constants Dl > 0 and Ml > 0 such that Ml ≤
kpn+1
− pl,s+1 k ≤ Ĉk∇(pn+1
− pl,s+1 )k ≤ Dl , (for l = 1, 2), where the constant Ĉ depends only on Ω . Also, l
l assume that each pi,s+1 (t, ·) ∈ Cc∞ (Ω) so that kps+1 k∞ ≤ D̂k∇ps+1 k∞ , where D̂ is the distance between
17

the two parallel hyperplanes bounding Ω . Let U be the exact solution of the problem at hand (2.2). Then, we obtain the following inequalities by employing the error estimate in [80] (Theorem 1.5):
(A.5)
1
kps+1 k∞ ≤ k∇ps+1 k∞ ≤ k∇(U (ts+1 ) − ps+1 )k∞ + k∇U (ts+1 )k∞ ≤ C(U )(h + τ ) + k∇U (ts+1 )k∞ = M ,
D̂
for some constant C(U ) depending on U .
We therefore get from (A.1), (A.2), (A.3), Young’s inequality, and (A.4) that for i, l = 1, 2, i 6= l ,
1 n+1
kp
− pi,s+1 k2 +

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
