---
source: "https://arxiv.org/abs/2010.09181v2"
title: "Multiscale simulations for multi-continuum Richards equations"
author: "Jun Sur Richard Park, Siu Wun Cheung, Tina Mai"
year: "2020"
publication: "arXiv preprint / math.NA"
download: "https://arxiv.org/pdf/2010.09181v2"
pdf: "https://arxiv.org/pdf/2010.09181v2"
captured_at: "2026-06-24T11:10:31Z"
updated_at: "2026-06-24T11:10:31Z"
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

# Multiscale simulations for multi-continuum Richards equations

- 著者: Jun Sur Richard Park, Siu Wun Cheung, Tina Mai
- 年: 2020
- 掲載情報: arXiv preprint / math.NA
- 情報源: [arxiv](https://arxiv.org/abs/2010.09181v2)
- ダウンロード: https://arxiv.org/pdf/2010.09181v2
- PDF: https://arxiv.org/pdf/2010.09181v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

In this paper, we study a multiscale method for simulating a dual-continuum unsaturated flow problem within complex heterogeneous fractured porous media. Mathematically, each of the dual continua is modeled by a multiscale Richards equation (for pressure head), and these equations are coupled to one another by transfer terms. On its own, Richards equation is already a nonlinear partial differential equation, and it is exceedingly difficult to solve numerically due to the extra nonlinear dependencies involving the soil water. To deal with multiple scales, our strategy is that starting from a microscopic scale, we upscale the coupled system of dual-continuum Richards equations via homogenization by the two-scale asymptotic expansion, to obtain a homogenized system, at an intermediate scale (level). Based on a hierarchical approach, the homogenization's effective coefficients are computed through solving the arising cell problems. To tackle the nonlinearity, after time discretization, we use Picard iteration procedure for linearization of the homogenized Richards equations. At each Picard iteration, some degree of multiscale still remains from the intermediate level, so we utilize the generalized multiscale finite element method (GMsFEM) combining with a multi-continuum approach, to upscale the homogenized system to a macroscopic (coarse-grid) level. This scheme involves building uncoupled and coupled multiscale basis functions, which are used not only to construct coarse-grid solution approximation with high accuracy but also (with the coupled multiscale basis) to capture the interactions among continua. These prospects and convergence are demonstrated by several numerical results for the proposed method.

## PDF Text

Multiscale simulations for multi-continuum
Richards equations
1

arXiv:2010.09181v2 [math.NA] 2 Jun 2021

Jun Sur Richard Parka , Siu Wun Cheungb , Tina Maic,d,∗
Abstract. In this paper, we study a multiscale method for simulating a dual-continuum unsaturated flow problem within complex heterogeneous fractured porous media. Mathematically, each of the dual continua is modeled by a multiscale Richards equation (for pressure head), and these equations are coupled to one another by transfer terms. On its own, Richards equation is already a nonlinear partial differential equation, and it is exceedingly difficult to solve numerically due to the extra nonlinear dependencies involving the soil water. To deal with multiple scales, our strategy is that starting from a microscopic scale, we upscale the coupled system of dual-continuum Richards equations via homogenization by the two-scale asymptotic expansion, to obtain a homogenized system, at an intermediate scale (level). Based on a hierarchical approach, the homogenization’s effective coefficients are computed through solving the arising cell problems.
To tackle the nonlinearity, after time discretization, we use Picard iteration procedure for linearization of the homogenized Richards equations. At each Picard iteration, some degree of multiscale still remains from the intermediate level, so we utilize the generalized multiscale finite element method (GMsFEM) combining with a multi-continuum approach, to upscale the homogenized system to a macroscopic (coarse-grid) level. This scheme involves building uncoupled and coupled multiscale basis functions, which are used not only to construct coarse-grid solution approximation with high accuracy but also (with the coupled multiscale basis) to capture the interactions among continua.
These prospects and convergence are demonstrated by several numerical results for the proposed method.

Keywords. Nonlinear Richards equations; Unsaturated flow; Coupled system; Multicontinuum; Upscaling; Hierarchical finite element; Heterogeneous fractured porous media;
Multiscale method; GMsFEM
Mathematics Subject Classification. 65N30, 65N99
a

Jun Sur Richard Park ; Department of Mathematics, The University of Iowa, Iowa City,
IA, USA; junsur-park@uiowa.edu b

Siu Wun Cheung; Center for Applied Scientific Computing, Lawrence Livermore National
Laboratory, Livermore, CA 94550, USA; cheung26@llnl.gov
∗

Corresponding author: Tina Mai ; c Institute of Research and Development, Duy Tan
University, Da Nang, 550000, Vietnam; d Faculty of Natural Sciences, Duy Tan University,
Da Nang, 550000, Vietnam; maitina@duytan.edu.vn
Date: June 3, 2021.
1
Accepted manuscript by Journal of Computational and Applied Mathematics (2021).
doi of published journal article: https://doi.org/10.1016/j.cam.2021.113648.
1

The

2

J.S.R. Park, S.W. Cheung, T. Mai

1. Introduction
Soil moisture predictions have imperatively drawn attention not only in agriculture but also in hydrology, environment, energy balances and global climate forecast, etc. The involved processes are mathematically modeled by unsaturated flow, namely, Richards equation [79, 18, 17, 55, 49], which delineates the seepage of water into some porous media having pores filled with water and air [36]. Moisture close to the soil surface is mainly affected by precipitation together with evaporation that are strongly coupled in nonlinear manners, which leads to our considering coupled Richards equations. It is noticed from [49] that analytical solutions of Richards equation can be found only with restrictive assumptions, so most applied problems demand a numerical solution in one or two or three dimensions, especially in our setting of coupled equations. One of the most interesting features of the Richards equation is that regardless of its simple derivation, it is arguably one of the hardest equations to find reliable and accurate numerical solution in all of hydrosciences [49]. Moreover, the porous media can possess complex heterogeneous rock characteristics, intricate geography of fractures, multi-continuum background, high contrast and multiple scales, etc.
Among various challenges, the fractures’ high permeability can significantly influence the fluid flow processes and raise a requirement for a specific strategy in order to build mathematical model and computational approaches. One of the earliest strategies is the hierarchical model, which is employed to characterize the multiscale fractures interacting with porous materials [28]. Now, given at hand a multi-continuum strategy [12, 91,
59, 92, 78], we can represent and analyze complicated processes with multiple scales of heterogeneity or fractures. Generally, in each continuum, different fluid flow models can be used. For example, in [60], the mathematical model allows coupling Darcy–Forchheimer flow in the fractures with Darcy flow in the matrix.
In this work, we focus on a multi-continuum model for unsaturated flow, which arises from the coupled system of nonlinear Richards equations, in complex multiscale fractured porous media. Our considering dual-continuum background thus consists of a system of small scale highly connected fractures (the so-called natural ones or well-developed ones or highly developed ones) as the first continuum and a matrix as the second continuum
[12]. In such heterogeneous multi-continuum media, the simulation of nonlinear fluid flow as coupled Richards equations is even more challenging, primarily because of the nonlinearity, different properties of continua, high contrast, multiple scales, and mass transfer among continua and a variety of scales in many forms.
To solve those difficult problems, one needs some kind of model reduction for flow simulation. Conventional methods involve dividing the considered domain into coarse-scale grid blocks, where effective properties in each coarse block are computed [37]. This procedure in common upscaling methods based on homogenization employs the fine-scale solutions of local problems in every coarse block or representative volume. Such computation, however, may not reveal multiple important modes in each coarse block including the continua’s interaction.
That downside urged the multi-continuum strategies [12, 9, 91, 59, 92, 78] on coarse grid.
Physically, each continuum is considered as a system (over the whole domain) so that the flow between different continua can be conveniently described. In the fine grid, different

Multiscale simulations for multi-continuum Richards equations

3

continua are neighboring. In the coarse grid, they co-exist (through mean characteristics
[12]) at every point of the region, and they interact with one another. Mathematically, a couple of equations are formed for each coarse block, and an individual equation corresponds to one of the dual continua on the fine grid. For example, in fractured media, the flow equations for the system of natural fractures and the matrix are written distinctly with some exchange terms. Such interaction terms are coupled toward a system of coupled multiscale equations following the mass conservation law. In order to reach that goal, even when each continuum is not connected topologically, one can assume that it is connected to the other (throughout the type of the coupling and the whole domain), given that it has only global (non-local) effects.
In these contexts, we now examine (in detail) our chosen dual-continuum background.
For simulating flow in naturally fissured rock, Barenblatt introduced the first dual-porosity model [12]. He suggested in that work two continua to describe low and high porosity continua, that is, respectively system of small connected fractures and matrix, which are also used in our paper. An example of some beginning work on dual continua based on
[12] is [9] (1990), where homogenization theory was employed, within a naturally fractured medium, to obtain a general form of the double-porosity model representing single-phase flow. For each continuum, both inter and intraflow exchanges are taken into consideration.
Basically, the dual-continuum background can be in any form, where the above schemes can be utilized.
For such homogenization theory but now with multiple scales, we contribute in this paper the detailed results (thanks to [76, 75]) on homogenization of the two-scale dualcontinuum system of nonlinear Richards equations (which are linear in [76, 75]), with optimal computational cost. As in [75], the interaction between the continua here is scaled as 1/ (which was 1/2 in [76]), and  stands for the periodically microscopic scale of the material. This homogenized problem arrives from the two-scale asymptotic expansion [13, 11, 58]. For the dual continua, we show that given this scale 1/ (as in [75])
of the interaction term, the coupled limits (which became only one equation in [76] for the scale 1/2 ) occur when letting the microscopic scale  vanish in the homogenization procedure. The homogenization’s effective coefficients are computed from solutions of four cell problems (as in [75]). Solving these local cell problems at every macroscopic point using the same fine mesh is expensive because there are a huge number of such points.
Therefore, our additional contribution is a development of the hierarchical technique from [76] (which has not been investigated in [75]) for solving that four cell problems using a reasonable number of degrees of freedom while the accuracy essentially remains.
More specifically, we solve the four cell problems using a dense network of macroscopic points. This scheme takes advantage of the fact that there are similar characteristics of neighboring representative volume elements (RVEs) [76, 75, 24, 74], which leads to close effective properties of the neighboring RVEs. Such a hierarchical technique achieves optimal computational complexity by employing different levels of resolution of FE spaces at different macroscopic points, for the four cell problems. In particular, regarding the cell problems, solutions at the points belonging to lower levels in the hierarchy, which are solved from higher levels of resolution, are employed to correct each solution at a nearby macroscopic point lying on a higher level in the hierarchy, which is obtained via a lower

4

J.S.R. Park, S.W. Cheung, T. Mai

level of accuracy (resolution). This hierarchy of macrogrids of points corresponds to a nest of approximation spaces having different levels of resolution. More specifically, each level of macroscopic points is appointed to an approximation finite element (FE) space, then these points as well as space are where one solves the cell problems. We theoretically prove that this hierarchical FE approach reaches the same level of accuracy as that of the full solve where cell problems at every macroscopic point are solved by FE space with the highest level of resolution, but the required number of degrees of freedom is optimal as expected.
In the literature, for other multiscale equations, the hierarchical finite element (FE)
scheme has been advanced to solve the cell problems and calculate the effective coefficients. In [15], the approach was developed for the case of deterministic two-scale Stokes–
Darcy systems within a slowly varying porous material. Later in [16], the hierarchical strategy was applied to a two-scale ergodic random homogenization problem without any assumption on microscopic periodicity. Based on these achievements, in our paper, we extend the hierarchical approach to the two-scale dual-continuum nonlinear system where the interaction terms are scaled as O(1/), with respect to the computation of homogenized coefficients. The interaction terms lead to interesting four cell problems (as in [75]), that is, a system of coupled four equations.
Starting from a microscopic scale , the resulting homogenized coupled Richards equations are still nonlinear at an intermediate scale. To tackle the challenge from the nonlinearity, after discretization by time, we apply linearization via traditional Picard iteration
(with a reasonable termination criterion) for each time step until the stopped time. At the current iteration of the linearization, note that (as in [74]) some degree of multiscale still remains in the homogenized equations, so direct method (to be discussed later) is not effective. To overcome the difficulties from such multiple scales as well as high contrast, fractures, heterogeneity, and to reduce computational cost, we utilize our recent numerical strategy [74] based on the GMsFEM [43, 22, 21] combining with the dual-continuum approach, to attain coarse-grid (macroscopic) level of the coupled dual-continuum linearized homogenized equations. Note that the fine-grid scale in our paper (as in [74]) is the intermediate scale resulting from the homogenization procedure, so it is different from the fine-grid scale of the GMsFEM in [28] and [81].
Normally, the GMsFEM helps us systematically create multiple multiscale basis functions, by including new basis functions (degrees of freedom) in every coarse block. Such new basis functions are computed by establishing local snapshots and operating local spectral decomposition in the snapshot space. Hence, the obtained eigenfunctions can convey the local properties to the global ones, through the coarse-grid multiscale basis functions.
It is important to note that the GMsFEM’s convergence is related to the eigenvalue decay of the local spectral problems. To construct the global multiscale basis functions, we select certain number of eigenfunctions that correspond to the smallest eigenvalues from each coarse neighborhood. Then the error converges with a rate correlated with 1/Λ, where Λ is the minimum of the excluded eigenvalues over all the coarse neighborhoods.
This suggests that one needs to include all the multiscale basis functions that correspond to the smallest eigenvalues for a good solution’s accuracy (see [22], for instance).

Multiscale simulations for multi-continuum Richards equations

5

Efficiently, the GMsFEM has been utilized for a variety of multi-continuum problems.
Recently, there has been an example [2] regarding shale gas transport in dual-continuum background comprising organic and inorganic media. With this motivation, a third continuum can be included in dual continua toward triple continua (see [90], for instance).
On the whole, flow simulation was investigated in heterogeneously varying multicontinua
[28, 81, 84] and in fractured porous media [28, 1, 3].
Before the advent of GMsFEM, several multiscale methods are developed to solve problems with such heterogeneous characteristics, for example, multiscale finite element method (MsFEM) [45, 42] (which the GMsFEM directly based on), multiscale finite volume method (MsFVM) [54], and heterogeneous multiscale methods (HMM) [38]. More specifically, in [57, 56, 53], the authors establish the coarse-grid approximation thanks to the MsFEM for solving the unsaturated flow problems possessing heterogeneous coefficients. In [19], upscaling method is utilized for the Richards equation. Especially, multiscale methods for flow problems in fractured porous media are addressed in [1].
In light of the GMsFEM’s success, there are numerous and vital studies on new model reduction techniques consisting of constraint energy minimizing (CEM) GMsFEM [27, 26,
52, 23] and involved numerical methods for multi-continuum systems with fractures [20]
comprising non-local multi-continuum method (NLMC) [86, 30, 87, 88]. In NLMC approach, one builds multiscale basis functions from the solutions of some local constrained energy minimizing problems as in the CEM-GMsFEM. These strategies also efficaciously deal with multiscale as well as high-contrast components in multi-continuum fractured media.
Besides such advanced multiscale methods, as mentioned above, traditional direct approach can tackle those difficulties in multi-continuum flow simulation, via fine-grid simulation, in a couple of steps. First, a locally fine grid is created. Second, one discretizes the flow equations on that fine grid and derives a global solution from the set of local solutions. This scheme can be operated under popular frameworks, for example, the Finite Element Method (FEM) in [10] and Finite Volume Method (FVM). However, due to enormous size of computation, the direct numerical simulation of multiple-scale problems is hard even with the assistance of supercomputers. By parallel computing (which utilizes domain decomposition methods), this hardness can be eased a little [45]. Classical parallel computing methods, nevertheless, demand rigorous interaction among processes, whereas the size of the discrete problem still remains. Hence, multiscale methods are needed, to obtain the small-scale effect on the large scales (without the desire of solving all the smallscale details) via constructing coarse-grid approximations through offline (precomputed)
multiscale basis functions.
Given a time step and a Picard iteration of the linearization (of the above homogenized nonlinear equations), we present the GMsFEM [43, 32] for solution of the linearized homogenized equations from the nonlinear unsaturated multi-continuum flow problem in heterogeneous porous fractured materials. We base on the previous works for linear case
[74] and for saturated flow problem [28]. Herein, we respectively extend those work to fractured domain and study solution of the unsaturated flow problem in two-dimensional case (while our technique can also work for three-dimensional case).

6

J.S.R. Park, S.W. Cheung, T. Mai

Within the GMsFEM, we investigate two basis types: uncoupled and coupled multiscale basis. In the first case (called uncoupled GMsFEM), multiscale basis functions will be established for each continuum independently, by taking into consideration only the permeability and ignoring the transfer terms. We then employ the GMsFEM presented above. In the second case (called coupled GMsFEM, which is focused in our paper), multiscale basis functions will be built by solving a coupled problem for snapshot space and operating a spectral decomposition. From this step, the GMsFEM is also applied. Generally, in the GMsFEM, multiscale basis functions are constructed in order to automatically locate each continuum through solution of the local spectral problems [31, 42, 43]. Now, for coupled system of equations in multi-continuum models, we solve local coupled system of equations in order to establish highly accurate coupled multiscale basis functions that describe complex interaction among continua in the coarse-grid level.
In numerical simulations, within each Picard iteration, we aim at coupling the GMsFEM with the multi-continuum approach. To demonstrate accuracy and robustness of the proposed coupled GMsFEM, we consider a dual-continuum background (connected fractures and matrix) model as above. Several numerical examples are presented for twodimensional test problems. At the last time step and at the end of the Picard iteration procedure, the multiscale solution is compared with the reference fine-scale solution. Our numerical results (after benefiting both coupled and uncoupled GMsFEM) prove that the GMsFEM solutions converge when we increase the number of local basis functions, and that coupled GMsFEM (for large interaction coefficients) has higher accuracy than uncoupled GMsFEM. Also, our numerical results show that the GMsFEM is able to incorporate with the dual-continuum approach to reach an accurate solution via only few basis functions, and that the GMsFEM is robust respecting high-contrast coefficients. Regarding further results about Picard iteration procedure for linearization of the coupled dual-continuum systems of Richards equations, numerically, convergence is guaranteed by a very small termination criterion number. Theoretically, we also prove the global convergence of this Picard linearization process in Appendix D.
The paper is organized as follows. In Section 2, preliminaries are presented. Section 3 is about formulating the system of dual-continuum coupled nonlinear multiscale
Richards equations and the corresponding homogenized system (that we will work with), in heterogeneous fractured porous media. We provide in Section 4 fine-scale finite element discretization and Picard iteration procedure for linearization of the homogenized system. In Section 5, the GMsFEM is presented for our homogenized coupled nonlinear system, making use of both uncoupled and coupled GMsFEM. We show in Section 6
several numerical results for this homogenized system. Section 7 is regarding open discussion. Conclusions are congregated in Section 8. In Appendix A, we give a derivation of the homogenized equations (introduced in Section 3). Appendix B is devoted to analyzing hierarchical numerical solutions of the cell problems from the homogenization. Appendix
C gives a proof of the main Theorem B.3 (about convergence results for the hierarchical solve). In the last Appendix D, we present a proof for the global convergence of Picard linearization process.

Multiscale simulations for multi-continuum Richards equations

7

2. Preliminaries
Let Ω be our bounded computational domain in Rd . To ease our discussion, we consider d = 2 in the remaining of the paper, but the method can be generalized to d = 3. We refer the readers to [33, 51, 52, 74] for the basic preliminaries. Latin indices i, j are in the set {1, 2}. Functions are denoted by italic capitals (e.g., f ), vector fields in R2 and 2 × 2
matrix fields over Ω are denoted by bold letters (e.g., v and T ). The spaces of functions, vector fields in R2 , and 2 × 2 matrix fields defined over Ω are respectively represented by italic capitals (e.g., L2 (Ω)), boldface Roman capitals (e.g., V ), and special Roman capitals (e.g., S).
Throughout this paper, the symbol ∇ stands for the gradient with respect to x of a function which only depends on the variable x (or the variables x and t). By ∇x , we represent the partial gradient with respect to x of a function which depends on x, t and other variables as well. Einstein summation is reflected by repeated indices. The class C 0
consists of all continuous functions. Spaces of periodic functions come with subscript # .
Consider the space V := H01 (Ω) = W01,2 (Ω). Its dual space (also called the adjoint space) is denoted by H −1 (Ω) , which consists of continuous linear functionals on H01 (Ω).
The value of a functional f ∈ H −1 (Ω) at a point v ∈ H01 (Ω) is denoted by the inner product hf, vi. Whereas, the notation (·, ·) stands for the standard L2 inner product.
The Sobolev norm k · kW 1,2 (Ω) is of the form
0


 21
kvkW 1,2 (Ω) = kvk2L2 (Ω) + k∇vk2L2 (Ω)
.
0

Here, k∇vkL2 (Ω) := k|∇v|kL2 (Ω) , where |∇v| denotes the Euclidean norm of the 2component vector-valued function ∇v; and for v = (v1 , v2 ), k∇vkL2 (Ω) := k|∇v|kL2 (Ω) , where |∇v| denotes the Frobenius norm of the 2 × 2 matrix ∇v. We recall that the
Frobenius norm on L2 (Ω) is defined by |A|2 := A · A = tr(AT A) .
The dual norm to k · kH01 (Ω) is k · kH −1 (Ω) , that is,
|hf, vi|
.
v∈H01 (Ω) kvkH01 (Ω)

kf kH −1 (Ω) = sup

For every 1 ≤ r < ∞, we use Lr (0, T ; X) to represent the Bochner space [48] with the norms
Z T
1/r r
kφkLr (0,T ;X) :=
kφkX dt
< +∞ ,
0

kφkL∞ (0,T ;X) := sup kφkX < +∞ ,
0≤t≤T

where (X, k · kX ) is a Banach space, for example X = H01 (Ω) . Also, we define

H 1 (0, T ; X) := φ ∈ L2 (0, T ; X) : ∂t φ ∈ L2 (0, T ; X) .
To reduce notation [50], instead of L2 (0, T ; H01 (Ω)), we denote by V = H01 (Ω) the space for the pressure head pi (t, ·) and by V = V × V = H01 (Ω) × H01 (Ω) the space for p(t, ·) =
(p1 (t, ·), p2 (t, ·)), where i = 1, 2 and t ∈ [0, T ], T > 0 .

8

J.S.R. Park, S.W. Cheung, T. Mai

3. System of coupled Richards equations and its homogenization
As in Section 1, let  be the characteristic length representing the periodically small scale variability of the media. In the current section, the system of Richards equations is based on [81, 19, 89]. We first give an overview of the system then focus on the 1/-scale case in Subsection 3.1.
The following system of Richards equations is considered (by combining the mass conservation law with Darcy’s law, respectively):
∂Θ1 (p1 (t, x))
+ div(v1 ) + L12 = f1 (t, x) in (0, t) × Ω ,
∂t
∂Θ2 (p2 (t, x))
+ div(v2 ) − L21 = f2 (t, x) in (0, t) × Ω ,
∂t v1 = −k1 (x, p1 )∇p1 (t, x)
in (0, t) × Ω ,

(3.1)

v2 = −k2 (x, p2 )∇p2 (t, x)

in (0, t) × Ω ,

where each pi = pi (t, x) [m] denotes the pressure head, for i = 1, 2 (each i corresponds to a continuum); vi [ms−1 ] represents the Darcy velocity; Θi (pi (t, x)) is volumetric soil water content; fi (t, x)[s−1 ] refers to source or sink term; Ki (x, pi )[ms−1 ] stands for unsaturated hydraulic conductivity tensor [89, 81], which is bounded, symmetric and positive definite.
In our paper, the media are simply assumed to be isotropic so that each hydraulic conductivity Ki (x, pi ) becomes function ki (x, pi ) multiplying with the identity matrix I [34].
The techniques here can be extended to anisotropic media. We use L12 and L21 to denote the transfer terms between fracture-matrix and matrix-fracture, respectively, where
Z
Z

L12 dx − L21 dx = 0 , and L12 = L21 ≈ Q12 (x, p1 , p2 )(p1 − p2 ) ,
Ω

Ω

with the mass transfer (exchange) term [81]
(3.2)

Q12 = Q21 = Q1 = Q2 = Q ,

that are calculated via the lower unsaturated hydraulic conductivity
(3.3)

Q1 = Q2 = ζk2 (x, p2 ) ,

having ζ as the shape factor [91, 12].
After substituting the Darcy’s law (the last two equations of (3.1)) and the mass transfer term Q12 at (3.2) into the mass conservation equations (the first two equations of (3.1)), we get the following system for the pressure heads:
∂Θ1 (p1 )
− div(k1 (x, p1 )∇p1 ) + Q1 (x, p1 , p2 )(p1 − p2 ) = f1 in (0, t) × Ω ,
∂t
(3.4)
∂Θ2 (p2 )
− div(k2 (x, p2 )∇p2 ) + Q2 (x, p1 , p2 )(p2 − p1 ) = f2 in (0, t) × Ω .
∂t
Note that we use only one domain Ω here, as in [84, 74, 76, 75].
In Section 6, the unsaturated hydraulic conductivity is written as a product of permeability (depending on the spatial variable) and some function (depending on the pressure head variable) [81, 89]. Note that hydraulic conductivity [8] is the fluid’s ability to pass through material, while permeability [8] is the material’s ability allowing fluid to pass

Multiscale simulations for multi-continuum Richards equations

9

through. Permeability is a feature of the medium itself, whereas hydraulic conductivity is the characteristic of both medium and fluid [8]. Also note that generally, intrinsic permeability or simply permeability k (temperature independent) does not coincide with saturated hydraulic conductivity Ks (temperature dependent).
The source or sink term fi (t, x) can be potential recharge flux at the ground surface, evaporation, (plant) transpiration, precipitation rate [61], leakage through a confining layer, rainfall, pumpage [35], surface infiltration, evapotranspiration, and runoff [61], etc.
Positive values of fi represent a source function, while negative values of fi imply a sink function [35]. Some physical meanings of these fi terms [41] are as follows. The unsaturated flow (Richards equation) needs to be equipped with a term representing water uptake by plant roots whenever soil-water-plant system is considered. This term
(as a function) is referred to as the sink term because water in the soil is withdrawn by the root system. When water is supplemented to the soil, by an aquifer for instance, then the equipped term (as a function) is referred to as the source term.
The form of fi can depend on the pressure heads pi [61]. For example, f1 = k1 p1 − k2 p2 , f2 = −k1 p1 +k2 p2 , such that f1 +f2 = 0 . Both fi can be simultaneously positive as in [82], also see [61] for coupled surface and subsurface flows. Each fi can be either positive or negative in a region. For example, in the vadose zone, we consider a source term possessing both positive and negative values given by f = f (x, z) = 0.006cos(4/3πz)sin(2πx) on
Ωvad , whereas it holds that f ≡ 0 in the saturated zone Ωgw [67].
Some explicit examples of source or sink term are from [65, 85] as follows. The first form of sink term is from [65]:
(3.5)

f1 (p1 ) = α(p1 )Smax ,

where α(p1 ) is a dimensionless soil water availability factor of pressure head p1 , and Smax is the maximum possible root-water extraction when soil water is not restricting (that is, the potential transpiration). The second example of sink term is from [85], as follows:
(3.6)

f1 (x, p1 ) = α1 (x, p1 )α2 (x, p1 )g(x)S ,

where S is the potential transpiration rate, g(x) is the root density function, α1 (x, p1 ) represents the compensation mechanism, α2 (x, p1 ) stands for water stress (for interpretation of each component, see [85], for instance).

3.1. The 1/-scale. This Subsection is based on [75]. Recall from Section 2 that Ω is the bounded computational domain in R2 and Latin indices i, j vary in the set {1, 2}. We let Y be a unit cube in R2 and [a, b] be an interval in R . With pi ∈ V = H01 (Ω) having values in [a, b], we define the following two-scale coefficients:
 x 
 x

(3.7)
ki (x, pi ) = ki x, , pi , Qi (x, p1 , p2 ) = Qi x, , p1 , p2 .


Note that these two-scale coefficients can be written in the following form (without x):
x 
x

(3.8)
ki (x, pi ) = ki
, pi , Qi (x, p1 , p2 ) = Qi
, p1 , p2 ,



10

J.S.R. Park, S.W. Cheung, T. Mai

whenever our discussion does not involve x alone. Here, we assume that Qi (y, p1 , p2 ) ∈
C 0 (Y × [a, b] × [a, b]) and ki (y, pi ) ∈ C 0 (Y × [a, b]) are continuous functions, which are x
Y -periodic with respect to y = . Also, each fi (t, ·) is assumed to be in L2 (Ω) .

Now, we consider the following problem:
(3.9)
∂p1 (t, x)
1
− div[k1 (x, p1 (t, x))∇p1 (t, x)] + Q1 (x, p1 (t, x), p2 (t, x))(p1 (t, x) − p2 (t, x))
∂t

= f1 (t, x) in (0, T ) × Ω ,
∂p2 (t, x)
1
− div[k2 (x, p2 (t, x))∇p2 (t, x)] + Q2 (x, p1 (t, x), p2 (t, x))(p2 (t, x) − p1 (t, x))
∂t

= f2 (t, x) in (0, T ) × Ω .
For i, j = 1, 2 , this system (3.9) is equivalent to the following equations:
(3.10)
∂pi (t, x)
1X 
− div(ki (x, pi )∇pi ) +
Qi (x, p1 , p2 )(pi − pj ) = fi (t, x) in (0, T ) × Ω .
∂t
 j
The system of equations (3.10) is equipped with the initial condition p1 (0, x) = 0, p2 (0, x) = 0 in Ω , and with the Dirichlet boundary condition p1 (t, x) = p2 (t, x) = 0
on (0, T ) × ∂Ω .
Note that in (3.4), the volumetric water content Θi (pi ) is usually a nonlinear function of the pressure head pi , and the following form is also valid [18]:
∂p
∂Θi (pi )
= C(pi ) i .
∂t
∂t
In our analysis, the homogenization much involves the nonlinear hydraulic conductivity ki (x, pi ) . Whereas, the term C(pi ) is not important and can be ignored so that Θi is the identity function, that is, Θi (pi ) = pi in (3.9).
Toward homogenization of the original system (3.10), we postulate that the system’s solutions (that is, the pressure heads p1 and p2 ) admit the following two-scale asymptotic expansion (see [40], for instance):



x
x
x

2
p1 (t, x) = p10 t, x,
+ p11 t, x,
+  p12 t, x,
+ ··· ,






(3.11)
x x
x p2 (t, x) = p20 t, x,
+ p21 t, x,
+ 2 p22 t, x,
+ ··· ,



in Ω . For i = 1, 2, we define

x
x


pi = pi (t, x) := p̂i t, x,
= p̂i (t, x, y) , with y = .


In order to process further, we denote
(3.12) p1 = p1 (t, x) := p10 (t, x) = p10 (t, x, y) ,

p2 = p2 (t, x) := p20 (t, x) = p20 (t, x, y) ,

where the fact that p10 (t, x, y) and p20 (t, x, y) are independent of y can be easily verified as in [75].

Multiscale simulations for multi-continuum Richards equations

11

Assume that the mass transfer term, the hydraulic conductivity and its spatial gradient are uniformly bounded, that is, there exist positive constants k, k and m, m such that the following inequalities hold:
(3.13)

k ≤ k1 (y, p1 ), k2 (y, p2 ), |∇y k1 (y, p1 )|, |∇y k2 (y, p2 )| ≤ k , m ≤ Q1 (y, p1 , p2 ), Q2 (y, p1 , p2 ) ≤ m .

Resulting from the original system (3.10), our homogenized system is (A.3). The details of the homogenization’s derivation and the hierarchical algorithm for the cell problems can be found in Appendices A and B (with a proof C), respectively.
Now, we consider p1 = p1 (t, x) = p10 (t, x), p2 = p2 (t, x) = p20 (t, x) in V = H01 (Ω) as from (3.12), the homogenized conductivities κ1 (p1 ), κ2 (p2 ) abbreviated for κ1 (x, p1 ), κ2 (x, p2 )
in (A.5), and f1 (t, ·), f2 (t, ·) in L2 (Ω) as from (3.10).
Thanks to [6], one needs to solve the reduced form of the system (A.3) for p1 (t, ·), p2 (t, ·) ∈
V , that is, in the domain (0, T ) × Ω as follows:
(3.14)
∂p1
− div(κ1 (p1 ) ∇p1 ) + b11 (p1 , p2 ) · ∇p1 + b12 (p1 , p2 ) · ∇p2 + c1 (p1 , p2 ) (p1 − p2 ) = f1 ,
∂t
∂p2
− div(κ2 (p2 ) ∇p2 ) + b21 (p1 , p2 ) · ∇p1 + b22 (p1 , p2 ) · ∇p2 + c2 (p1 , p2 ) (p2 − p1 ) = f2 ,
∂t equivalently (for i, j = 1, 2),
(3.15)

X
X
∂pi
− div(κi (pi ) ∇pi ) +
[bij (p1 , p2 ) · ∇pj ] +
[ci (p1 , p2 ) (pi − pj )] = fi ,
∂t j
j

with the initial condition pi (0, x) = 0 in Ω , and with the Dirichlet boundary condition pi (t, x) = 0 on (0, T ) × ∂Ω . Note that in the system (3.14) and (3.15), each ci (p1 , p2 )
(abbreviated for ci (x, p1 , p2 )) is mass transfer term and each bij (p1 , p2 ) (abbreviated for bij (x, p1 , p2 )) is nonlinear velocity (see [6], for instance), which are explicitly defined in
(A.3) in terms of correctors being solutions of the so-called cell problems (A.6a)–(A.6b).

4. Fine-scale discretization and Picard iteration for linearization
This section is based on [71, 81, 52]. Note that there are a variety of recently developed linearization techniques for Richards equation, such as the Newton method, the modified
Picard algorithm [18], the L-scheme [67], and a new multilevel Picard iteration [39].
However, we will use the classical Picard linearization algorithm as in [73].
We are given an initial pair p0 = (p1,0 , p2,0 ) ∈ V . With i, j = 1, 2 , on each continuum i , providing a fixed ui ∈ V , and letting u = (u1 , u2 ) ∈ V , we define the following bilinear

12

J.S.R. Park, S.W. Cheung, T. Mai

forms: for p, φ ∈ V and p = (p1 , p2 ) ∈ V ,
Z
(4.1)
ai (p, φ; ui ) =
κi (ui )∇p · ∇φ dx ,
Ω
XZ
bi (p, φ; u) =
(4.2)
(bij (u1 , u2 ) · ∇pj )φ dx , j

(4.3)

qi (p, φ; u) =

Ω

XZ
j

ci (u1 , u2 )(pi − pj )φ dx .

Ω

Recall that (·, ·) denotes the standard L2 (Ω) inner product. The variational form of (3.15)
is as follows: for i = 1, 2, find p = (p1 , p2 ) ∈ V such that


∂pi
, φi + ai (pi , φi ; pi ) + bi (p, φi ; p) + qi (p, φi ; p) = (fi , φi ) ,
(4.4)
∂t with all φ = (φ1 , φ2 ) ∈ V , for a.e. t ∈ (0, T ) . For simplicity, we can drop the subscript i while keeping the meaning that each equation (4.4) corresponds to one of the dual continua.
To reach the first goal regarding time discretization (see [71, 81], for instance) of (4.4), we will apply the following standard backward Euler finite-difference scheme: find p =
(p1 , p2 ) ∈ V such that for any φ = (φ1 , φ2 ) ∈ V ,


pi,s+1 − pi,s
, φi + ai (pi,s+1 , φi ; pi,s+1 ) + bi (ps+1 , φi ; ps+1 ) + qi (ps+1 , φi ; ps+1 )
(4.5)
τ
= (fi,s+1 , φi ) , where the time range [0, T ] is divided into S equal intervals, with the time step size
τ = T /S > 0, and the subscript s denotes the evaluation of a function at the time instant ts = sτ (for s = 0, 1, · · · , S).
Next, the nonlinearity in space will be handled through linearization using Picard iteration (see [71, 81, 52], for instance) as follows. At the (s + 1)th time step, we guess p0s+1 ∈ V . For n = 0, 1, 2, · · · , given pns+1 ∈ V , we find pn+1
s+1 ∈ V such that for any
φ = (φ1 , φ2 ) ∈ V ,
!
pn+1
i,s+1 − pi,s n
n+1
n n+1
n
(4.6)
, φi + ai (pn+1
i,s+1 , φi ; pi,s+1 ) + bi (ps+1 , φi ; ps+1 ) + qi (ps+1 , φi ; ps+1 )
τ
= (fi,s+1 , φi ) .
The Picard iteration process converges to a limit as n → ∞ (see a theoretical proof in
Appendix D). In practice, we terminate this process at an αth iteration when it meets a certain stopping criterion, and let
(4.7)

ps+1 = pαs+1

be the previous time data to proceed to the next temporal step in (4.5). Throughout this work, we propose a stopping criterion using the relative successive difference, that is,

Multiscale simulations for multi-continuum Richards equations

13

given a user-defined tolerance δ0 > 0, if
(4.8)

n kpn+1
i,s+1 − pi,s+1 kL2 (Ω)
≤ δ0 , kpni,s+1 kL2 (Ω)

for both i = 1, 2 , then we terminate the iteration procedure.
Now, we discuss the fine-grid notation. Toward discretizing the variational problem
(4.4), we first let Th be a fine grid with grid size h. Here, h is assumed to be significantly small so that the fine-scale solution is close enough to the exact solution. Second, we let
Vh be the conforming piecewise bilinear finite element basis space with reference to the rectangular fine grid Th , that is,
(4.9)

Vh := {u ∈ V : u|K ∈ Q1 (K) ∀K ∈ Th } ,

where Q1 (K) is the space of all bilinear (or multilinear if d > 2) elements on K. We let
Vh = Vh × Vh .
In Th , the fully Picard discrete scheme reads: starting with an initial ph,0 ∈ Vh , at the
(s + 1)th time step, we guess p0h,s+1 ∈ Vh , and iterate in Vh from (4.6):
(4.10)
!
pn+1
−
p i,h,s i,h,s+1
n+1
n+1
n n
n
, φi + ai (pn+1
i,h,s+1 , φi ; pi,h,s+1 ) + bi (ph,s+1 , φi ; ph,s+1 ) + qi (ph,s+1 , φi ; ph,s+1 )
τ
= (fi,s+1 , φi ) , with any φ = (φ1 , φ2 ) ∈ Vh , for n = 0, 1, 2, · · · , until meeting (4.8) at some αth Picard step. We use (4.7) for setting the previous time data ph,s+1 = pαh,s+1 to go ahead to the next time step in (4.5).
5. GMsFEM for coupled dual-continuum nonlinear equations
5.1. Overview. The purpose of this section is to build multiscale spaces (in the pressure head computation) for the coupled nonlinear system (4.4). Toward establishing an appropriate generalized multiscale finite element method (GMsFEM, [43]), from the linearized system (4.6) in Section 4, the nonlinearity may be treated as constant at each Picard iteration step (after temporal discretization) so that multiscale spaces are able to be built respecting this nonlinearity.
First, we discuss the coarse-grid notation. Let T H be a coarse grid that has Th as a refinement. We denote by H the coarse-mesh size (with h  H). Each element of T H is named a coarse-grid block (or patch or element). We call N the total number of coarse v
blocks (elements) and Nv the total number of interior vertices of T H . Let {xj }N
j=1 be the collection of vertices (nodes) in T H . The jth coarse neighborhood of the coarse node xj is defined by the union of all coarse elements Km ∈ T H possessing such xj , as follows:
[
(5.1)
ωj = {Km ∈ T H : xj ∈ Km } .
Our primary goal is using the GMsFEM to seek a multiscale solution pms (which is a good approximation of the fine-scale solution ph ). In order to do so, first, we utilize (on coarse grid) the GMsFEM [43], where we solve local problems (to be specified then) in each coarse neighborhood, to systematically construct multiscale basis functions (degrees of

14

J.S.R. Park, S.W. Cheung, T. Mai

freedom for the solution) that still contain fine-scale information. The resulting multiscale space is called the global offline space Vms , comprising multiscale basis functions. Last, we find the multiscale solution pms in Vms . For the GMsFEM, in this section, we note that the system (3.15) possesses multiscale high-contrast coefficients κi , bij , ci (depending on x) for i, j = 1, 2 .
We refer the readers to [43] for the GMsFEM’s details and to [74, 51, 28] for its overview.
In the following subsections, we will present two different methods (mainly based on [74])
for constructing uncoupled multiscale basis functions (uncoupled GMsFEM) and coupled multiscale basis functions (coupled GMsFEM). For each type of these methods, using the GMsFEM’s framework [43] as above, we build a local snapshot space for each coarse neighborhood ωj then solve a relevant local spectral problem (defined on the snapshot space), to generate a multiscale (offline) space Vms .
More specifically, in the next Subsections 5.2 and 5.3, provided pms,s (at time step sth)
n and pnms,s+1 (at time step (s + 1)th and Picard iteration nth), we will build Vms = Vms,s+1
.
0
In practice (Subsection 5.4), given pms,0 and a starting guess pms,1 , we will need to establish
0
. Here, the snapshot functions as well as the basis functions are timeonly one Vms = Vms,1
independent.
5.2. Uncoupled GMsFEM. This terminology means that multiscale basis functions in i
are constructed for the solutions pn+1
Vms i,ms,s+1 , separately, with i, j = 1, 2 .
We let the fine-scale approximation (FEM) space for the ith continuum be Vhi (ωj ) =
Vh (ωj ), which is the conforming space Vhi = Vh restricted to the coarse neighborhood ωj .
The notation Jh (ωj ) stands for the set of all nodes of the fine grid Th locating on ∂ωj . The cardinality of Jh (ωj ) is abbreviated by NJj , and we let the index k varies 1 ≤ k ≤ NJj .
We will construct multiscale basis functions for each ith continuum distinctly by excluding the transfer functions and taking into consideration only the conductivity κi .
In particular, for each ith continuum, on every coarse neighborhood ωj , we first solve
(j),snap the following local snapshot problem: find the kth snapshot function φk,i
∈ Vh (ωj )
satisfying
(j),snap

(5.2)

− div(κi (x, pni,ms,s+1 )∇φk,i

) = 0 in ωj ,

(j),snap

φk,i

= δk,i

on ∂ωj ,

where δk,i is a function defined as
(
1 m = k,
δk,i (xjm ) =
0 m=
6 k, for all xjm in Jh (ωj ) , 1 ≤ k ≤ NJj . Hence, for the ith continuum, we obtain the jth local snapshot space
(j),snap

i
Vsnap
(ωj ) = span{φk,i

1 ≤ k ≤ NJj } .

Now, let Ni be the number of interior vertices of T H on the ith continuum. The jth local multiscale basis functions are built on ωj with respect to the ith continuum, by
(j)
i solving the local spectral problems: find the kth eigenfunction ψk,i ∈ Vsnap
(ωj ) and its

Multiscale simulations for multi-continuum Richards equations

15

(j)

i corresponding real eigenvalue λk,i such that for all ξi in Vsnap
(ωj ) ,
(j)

(5.3)

(j)

(j) (j)

(j)

ai (ψk,i , ξi ) = λk,i si (ψk,i , ξi ) .

i
For any φ, ψ ∈ Vsnap
(ωj ) , these operators are defined as follows [51, 28, 74]:
Z
(j)
ai (φ, ψ) =
κi (x, pni,ms,s+1 )∇φ · ∇ψdx,
ωj

(5.4)
(j)

Z

si (φ, ψ) =

κi (x, pni,ms,s+1 )

ωj

Ni
X

!
|∇χl,i |2

φψ dx ,

l=1

where each χl,i is a standard multiscale finite element basis function in the ith continuum, for the coarse node xl (that is, with linear boundary conditions for cell problems [44]).
H
i
We note that {χl,i }N
l=1 is a set of partition of unity functions (for T ) supported in the ith continuum.
(j)
After arranging the eigenvalues λk,i from (5.3) in ascending order, we take the first
(j)

(j)

Lωj eigenfunctions, and they are still denoted as ψ1,i , · · · , ψLω ,i , . Last, we define the kth j
multiscale basis function for the ith continuum on ωj by
(j),ms

ψk,i

(j)

= χj,i ψk,i ,

where 1 ≤ k ≤ Lωj .
Within the ith continuum, the local auxiliary offline multiscale space is defined by n
o
(j),ms i
Vms
(ωj ) = span ψk,i
1 ≤ k ≤ Lωj .
We then define the global offline space i
Vms
=

Ni
X

o n
(j),ms i
1 ≤ j ≤ Ni , 1 ≤ k ≤ Lωj .
Vms
(ωj ) = span ψk,i

j=1

Finally, the multiscale space as the global offline space is defined by
1
2
Vms = Vms
× Vms
,

which will be employed to find solution at the next (n + 1)th Picard iteration.
5.3. Coupled GMsFEM. In this section, we construct coupled multiscale basis funcn+1
n+1
tions in Vms for the solution pn+1
ms,s+1 = (p1,ms,s+1 , p2,ms,s+1 ) .
Here, regarding the coupled local snapshot problems, we will take into account the interaction terms c1 and c2 from (3.14). For the GMsFEM analysis, the operators of eigenvalue problems should be symmetric. Thus, we only choose the symmetric part of c1 and c2 , namely, cs = (c1 + c2 )/2.
More specifically, for i, j, r = 1, 2 , on each coarse neighborhood
ωj , we solve the local


(j),snap

snapshot problem: find the snapshot functions φk,r

(j),snap

= φk,1,r

(j),snap

, φk,2,r

in Vh (ωj ) =

16

J.S.R. Park, S.W. Cheung, T. Mai

Vh (ωj ) × Vh (ωj ) that satisfy




(j),snap
(j),snap
(j),snap n
n
− div κ1 (x, p1,ms,s+1 )∇φk,1,r
+ cs (x, pms,s+1 ) φk,1,r − φk,2,r
= 0 in ωj ,




(j),snap
(j),snap
(5.5) − div κ2 (x, pn2,ms,s+1 )∇φ(j),snap
+ cs (x, pnms,s+1 ) φk,2,r − φk,1,r
= 0 in ωj , k,2,r
(j),snap

φk,r

= δk,r

on ∂ωj ,

for 1 ≤ k ≤ NJj , where each δk,r is specified as
δk,r (xm ) = δk (xm )er , r = 1, 2 ,

(5.6)

with all xm in Jh (ωj ) and {er | r = 1, 2} as a standard basis in R2 . The local snapshot space is then of the form o
n
(j),snap
1 ≤ k ≤ NJj , r = 1, 2 .
(5.7)
Vsnap (ωj ) = span φk,r


(j)
(j)
(j)
Now, we solve the local eigenvalue problems: find the kth eigenfunction ψk = ψk,1 , ψk,2 ∈
(j)

Vsnap (ωj ) and its corresponding real eigenvalue λk such that for all ξ ∈ Vsnap (ωj ) ,




(j)
(j) (j)
(j)
(j)
(5.8)
ac s ψ k , ξ = λ k s
ψk , ξ ,
Such operators are defined as follows [51, 28, 74]:
2 Z
X
(j)
κi (x, pni,ms,s+1 )∇φi · ∇ψi dx , acs (φ, ψ) =
(5.9)

i=1

ωj

2 Z
2
Nv
X
X
X
(j)
(j)
κi (x, pni,ms,s+1 )
|∇χl,i |2
s (φ, ψ) =
si (φi , ψi ) =
i=1

i=1

ωj

!
φi ψi dx ,

l=1

for any φ = (φ1 , φ2 ), ψ = (ψ1 , ψ2 ) ∈ Vsnap (ωj ) . Here, each χl,i is defined as in (5.4).
(j)
After sorting the eigenvalues λk from (5.8) in rising up order, we choose the first
(j)
(j)
smallest Lωj eigenfunctions and still call them ψ1 , · · · , ψLw . Now, the kth multiscale j
basis functions for ωj are defined by
(j),ms

ψk

(j)

(j)

= (χj,1 ψk,1 , χj,2 ψk,2 ) ,

where each χj,i is defined as in (5.4) and 1 ≤ k ≤ Lwj . These functions form the local offline multiscale space n
o
(j),ms
Vms (ωj ) = span ψk
1 ≤ k ≤ Lwj .
Last, we obtain the multiscale space (as the global offline space)
Vms =

Nv
X

n o
(j),ms
Vms (ωj ) = span ψk
1 ≤ j ≤ Nv , 1 ≤ k ≤ Lwj ,

j=1

where we seek solution for the later (n + 1)th Picard iteration.

Multiscale simulations for multi-continuum Richards equations

17

Remark 5.1. Note that the local spectral problems, (5.3) and (5.8), are constructed taking into account the convergence analysis. The convergence rate of the proposed methods is proportional to 1/Λ, where Λ represents the minimum among all eigenvalues that correspond to eigenfunctions that are not included in the global offline space. This suggests that we include optimal number of multiscale basis functions associated with the smallest eigenvalues (see [22], for instance).
5.4. GMsFEM for coupled nonlinear system of equations. At the fixed time step
(s + 1)th, our tactic (as in [51, 52]) is using either the uncoupled or coupled GMsFEM
(in Subsections 5.2 or 5.3) and the corresponding constructed offline multiscale space
0
(introduced at the end of Subsection 5.1) to solve the problem (3.15) with
Vms = Vms,1
equivalent variational form (4.4) via linearization based on Picard iteration. During the online stage, the model reduction scheme reads: starting with pms,0 ∈ Vms , at the (s+1)th time step, we guess p0ms,s+1 ∈ Vms and iterate in Vms from (4.6):
!
pn+1
−
p i,ms,s i,ms,s+1
n n+1
n
, φi + ai (pn+1
i,ms,s+1 , φi ; pi,ms,s+1 ) + bi (pms,s+1 , φi ; pms,s+1 )
τ
(5.10)

n
+ qi (pn+1
ms,s+1 , φi ; pms,s+1 ) = (fi,s+1 , φi ) ,

with φ = (φ1 , φ2 ) ∈ Vms , for n = 0, 1, 2, · · · , until reaching (4.8) at some αth Picard step.
We employ (4.7) for choosing the previous time data pms,s+1 = pαms,s+1 to advance on the next temporal step in (4.5).
6. Numerical results
In this section, we will show some numerical experiments to demonstrate the performance of our strategy, for both the uncoupled and coupled GMsFEM, with linearization by Picard iteration procedure. In the simulations, we consider the high-contrast coefficients ai (x) (for i = 1, 2), which are presented in Fig. 1, having Ω = [0, 1]2 . The blue regions are for a1 (x) = 10, a2 (x) = 1, whereas the yellow regions (channels) represent a1 (x) = 105 , a2 (x) = 10 . Each coefficient ai (x) is defined on 128 × 128 fine grid, while the coarse-grid size is H = 1/16 . We choose the terminal time T = Sτ = 2, and the time step size is taken as τ = 1/10. Each initial pressure pi is zero. The Picard iterative termination criterion is δ0 = 10−5 , which ensures the linearization process’s convergence.
We assume that each medium is isotropic (and the anisotropic case is handled similarly).
Then, each hydraulic conductivity tensor becomes κi multiplying with the identity matrix, for i = 1, 2 .
Example 1. We consider the following problem as a special case of (3.14) (using the provided conditions there): find p = (p1 , p2 ) ∈ V in the domain (0, T ) × Ω such that


∂p1
a1 (x)
105
− div
∇p1 + 30b11 · ∇p1 + 30b12 · ∇p2 +
(p1 − p2 ) = 1 ,
∂t
1 + |p1 |
1 + |p1 |


(6.1)
∂p2
a2 (x)
105
− div
∇p2 + 30b21 · ∇p1 + 30b22 · ∇p2 +
(p2 − p1 ) = 1 ,
∂t
1 + |p2 |
1 + |p2 |
with the initial condition pi (0, x) = 0 in Ω , and with the Dirichlet boundary condition pi (t, x) = 0 on (0, T ) × ∂Ω . The vector fields are given by b11 = b21 = (p1 , p1 ), b12 =

18

J.S.R. Park, S.W. Cheung, T. Mai

b22 = (−p2 , −p2 ), and the functions a1 (x), a2 (x) are high-contrast (multiscale) coefficients shown in Fig. 1.
This example has a physical meaning. That is, for the given dual-continuum porous media, the first equation can describe a flow (the homogenized first Richards equation) in small highly connected fracture network as the first continuum, and the second equation can represent a flow (the homogenized second Richards equation) in the matrix as the second continuum [81]. Our model problem (6.1) can also be attained from upscaling of some highly heterogeneous media using Representative Volume Element (RVE) Approach
[24]. In that paper, the authors employ sub RVE scale to form a multi-continuum model
(at fine-grid level), that is upscaled further. In our context, we assume that

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
