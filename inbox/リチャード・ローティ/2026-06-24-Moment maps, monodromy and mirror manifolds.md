---
source: "https://arxiv.org/abs/math/0104196v1"
title: "Moment maps, monodromy and mirror manifolds"
author: "R. P. Thomas"
year: "2001"
publication: "arXiv preprint / math.DG"
download: "https://arxiv.org/pdf/math/0104196v1"
pdf: "https://arxiv.org/pdf/math/0104196v1"
captured_at: "2026-06-24T11:11:35Z"
updated_at: "2026-06-24T11:11:35Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Rorty Philosophy and the Mirror of Nature"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Moment maps, monodromy and mirror manifolds

- 著者: R. P. Thomas
- 年: 2001
- 掲載情報: arXiv preprint / math.DG
- 情報源: [arxiv](https://arxiv.org/abs/math/0104196v1)
- ダウンロード: https://arxiv.org/pdf/math/0104196v1
- PDF: https://arxiv.org/pdf/math/0104196v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Via considerations of symplectic reduction, monodromy, mirror symmetry and Chern-Simons functionals, a conjecture is proposed on the existence of special Lagrangians in the hamiltonian deformation class of a given Lagrangian submanifold of a Calabi-Yau manifold. It involves a stability condition for graded Lagrangians, and can be proved for the simple case of $T^2$.

## PDF Text

arXiv:math/0104196v1 [math.DG] 19 Apr 2001

Moment maps, monodromy and mirror manifolds
R. P. Thomas

Abstract
Via considerations of symplectic reduction, monodromy, mirror symmetry and
Chern-Simons functionals, a conjecture is proposed on the existence of special
Lagrangians in the hamiltonian deformation class of a given Lagrangian submanifold of a Calabi-Yau manifold. It involves a stability condition for graded
Lagrangians, and can be proved for the simple case of T 2 .

1

Introduction

Just as explicit solutions of the Einstein and Hermitian-Yang-Mills equations exist only on spaces that are either low dimensional, non-compact and/or highly symmetric, so the equations for special Lagrangian (SLag) cycles, also important in physics, have the same properties. Physically there are also similarities in that we have two first order supersymmetric minimal energy equations (HYMs and SLag) implying the more standard second order equations (YMs and minimal volume equations).
There are powerful existence results of Calabi and Yau (and more recently Tian,
Donaldson and others) for the Einstein equations, and of Donaldson-Uhlenbeck-Yau for the HYM equations, so long as we are on a Kähler (or projective) manifold; this often reduces an infinite dimensional problem in PDEs to a finite dimensional problem in linear algebra. Producing many Kähler-Einstein (e.g. Calabi-Yau) manifolds becomes trivial, and dealing with Hermitian-Yang-Mills connections requires only algebraic computations; in both cases the complicated role of the Kähler form and/or metric is almost removed. This can be thought of as possible because of the existence of some infinite dimensional geometry recasting the equations in terms of moment maps and symplectic reduction. A similar situation for SLags would therefore be highly desirable. In particular it might give a way of studying SLags using only Lagrangians and symplectic geometry, much as HYM connections are studied via stable bundles and algebraic geometry.
This paper explores the mirror symmetry of holomorphic bundles (on a CalabiYau 3-fold M , often referred to here as ‘the complex side’) and Lagrangians (on the
1

mirror Calabi-Yau 3-fold W , ’the symplectic side’, known as the Kähler side in the physics literature). Many people have worked and are still working on proving some kind of direct correspondence between such objects given an SYZ torus fibration
[SYZ]; see for example [AP], [BBHM], [Ch], [Fu1], [Gr], [LYZ], [PZ], [Ty], and see
[MMM] for a review of this and many many more issues in mirror symmetry. Here, however, we work purely formally without reference to a particular pair of mirror manifolds, without worrying about what mirror symmetry might rigorously mean, and we will not try to give any explicit correspondence. Using mirror symmetry merely as motivation, we point out some similar structures on both sides of the mirror map. Under some conditions (in some ‘large complex structure’ or ‘semiclassical’ or somesuch limit) these structures might be genuinely dual; again it does not matter if they are not in general. For instance, physics [MMMS], [DFR] predicts that one should consider not the HYM equations and slope but some perturbation of them away from the large complex structure limit; however these equations also come from a moment map and, conjecturally, a stability condition (for a discussion of such matters see [Le] or [T3]). So while the slope and phase of Lagrangians discussed below might not be exactly mirror to slope of bundles, it should be mirror to something with analogous properties and significance.
Loosely, we would like to think of submanifolds in a fixed homology class as mirror to connections on a fixed topological complex bundle (with Chern classes mirror to the homology class); then Lagrangians should correspond to holomorphic connections (i.e. integrable connections; those with no (0, 2)-curvature) and special
Lagrangians to those with HYM curvature. These last two conditions should be stability conditions for the group actions of hamiltonian deformations and complex gauge transformations, respectively. The full picture is much more complicated, involving triangulated categories and so forth, as envisaged some six years ago in the seminal conjecture of Kontsevich [K]; we can ignore this in only using mirror symmetry as motivation. It could be noted, however, that the functionals defined below are additive under exact sequences of holomorphic vector bundles and sums of Lagrangians, so should extend to the derived category of coherent sheaves and the derived Fukaya category of Lagrangians respectively.
First note that while the connections side has a complex structure and a complex gauge group involved, the Lagrangian side needs complexifying. So motivated by
Kontsevich [K] and by physics (e.g. [SYZ]) we add in connections on the submanifolds (which will later reduce to flat connections on Lagrangians). The dictionary we are aiming towards, much of which is already standard, is the following in the
3-dimensional case; all the terms used will be defined in due course.

2

Complex side M

Symplectic side W

Ω = ΩM ∈ H 3,0

ω = ωW ∈ H 1,1

H ev

H3

Connections A on a fixed C ∞ complex Submanifolds/cycles L in a fixed class
√
bundle E; v(E) = ch(E) Td X ∈ H ev [L] ∈ H 3 , with a connection on C × L
RL
fC (A, L) = L0 (F + ω)2
RL
RL
= L0 (F 2 + ω 2 ) + 2 L0 ω ∧ F

CSC (A = A0 + a) =

R
1
tr ∂¯A a ∧ a + 2 a∧3 ∧ Ω
2
4π

M

3

0

Critical points : FA0,2 = 0
holomorphic bundles

Critical points : ω|L = 0, FA = 0
Lagrangians + flat line bundles

Holomorphic Casson invariant [T1]

Counting SLags [J]

Gauge group

U (1) gauge group on L

Complexified gauge group

Hamiltonian deformations

ω = ωM ∈ H 1,1

Ω = ΩW ∈ H 3,0

Moment map FA ∧ ω n−1

Moment map Im Ω|L

Stability, slope µ = rk1E

R

Stability, slope µ = vol1(L)

tr FA ∧ ω n−1

R

L Im Ω

(In the fourth line, v(E) is the Mukai vector of E; in the last line, vol (L) is the cohomological volume measured with respect to Re Ω.) A SLag cycle (of phase φ) in a Calabi-Yau is a Lagrangian with Im (e−iφ Ω)|L ≡ 0 [HL]; then Re (e−iφ Ω)|L is the
Riemannian volume form on L induced by the Ricci-flat metric on the Calabi-Yau.
Obviously, rotating Ω by e−iφ gives SLags in the more traditional sense of phase zero. The part of the theory to do with SLags will apply in all dimensions; it is only the functionals that are special to Calabi-Yau 3-folds.
We will partially justify the above table, though the symplectic structure and moment map give problems that will appear in due course. However we can derive enough to arrive at a conjecture about Lagrangians and SLags for which evidence
3

will be given by using monodromy and mirror symmetry from [ST] to interpret an example of Lawlor and Joyce [J].
Acknowledgements. The debt of any ex-student of Simon Donaldson writing a paper on moment maps should be clear. This work is also more immediately influenced by the papers [D], [J], [K]. In particular I was surprised to see the Lagrangian condition coming from a moment map in [D], [H], which does not fit into the scheme
I always supposed was true. So the purpose of this paper, apart from trying to set a record for the number of m’s in a title, is to expand on that scheme and to try to get the special condition from a moment map instead. This paper was finished in the summer of 2000 and reported on in [T3]; since then exciting new ideas have appeared in physics [Do] and mathematics [KS] better explaining mirror symmetry.
I would like to thank S.-T. Yau, C. H. Taubes and Harvard University for support, and Yi Hu, Albrecht Klemm, Ivan Smith and Xiao Wei Wang for useful conversations. Communications with Mike Douglas, Dominic Joyce, Paul Seidel, S.-T. Yau and Eric Zaslow have been extremely influential.

2

Chern-Simons-type functionals and critical points

Consider the space A of (0, 1)-connections A on a fixed complex bundle on a CalabiYau 3-fold M . This infinite dimensional space has a natural complex structure, with respect to which it admits a holomorphic functional, Witten’s holomorphic
Chern-Simons functional [W1], [DT],


Z
2
1
¯
tr ∂A0 a ∧ a + a ∧ a ∧ a ∧ Ω,
CSC (A = A0 + a) = 2
4π M
3
where Ω is the holomorphic (3,0)-form. It is infinitesimally gauge-invariant (gauge transformations not homotopic to the identity can give periods to CSC ) and its gradient is FA0,2 , with zeros the integrable connections. That is, after dividing by gauge equivalence (under which grad CSC is invariant), the critical points of CSC
form the space of holomorphic bundles of the same topological type. As critical points of a functional, moduli of holomorphic bundles have virtual dimension zero, and one might try to make sense of counting them – a holomorphic Casson invariant
[T1]. This is independent of deformations of the complex structure, but can have wall-crossing changes as the Kähler form varies. (This is because we count only stable bundles, and the notion of stability depends on a Kähler form.)
On the other hand, on a different Calabi-Yau 3-fold W (for instance the mirror, in some situation where this makes sense), Lagrangians are the critical points of a

4

functional too, on the space of all 3-dimensional submanifolds (or cycles): fR (L) =

Z L
L0

ω ∧ ω,

where ω is the symplectic form on W . Here L0 is a fixed cycle in the same homology class, and we integrate over a 4-chain with boundary L − L0 ; the functional fR is invariant under the choice of different, homologous, 4-chains (picking non-homologous
4-chains can give periods to fR ). It is invariant under deformations of L pulled back from hamiltonian deformations of W (deformations generated by vector fields vR on W whose contraction with ω is exact v p ω = dh at each point in time) as
L ω ∧ dh = 0, and its gradient is ω|L . Thus its critical points are Lagrangian submanifolds. We would like to think of fR as mirror to CSC , but to do so we must complexify it.
Thus we work on the space A of submanifolds L of W with U (1) connections A
on the trivial bundle C × L on L. Notice these submanifolds are not parameterised by a map of a real 3-manifold into W ; we are only interested in the image L. From now on we shall restrict attention to smooth Lagrangian submanifolds. Formally, we consider the tangent space to A at a point (A, L ⊂ W ) to be
Ω1 (L; R) ⊕ Ω1 (L; R),

(2.1)

at least for those L with no J-invariant subspaces of its tangent spaces (J is the complex structure on W , and this is reasonable since we are looking for Lagrangian submanifolds after all). The first factor is the obvious tangent space to the connections on L; the second gives deformations of L via the vector fields produced by contracting with the Kähler form ω on W . That is, we use the metric on W to map
Ω1 (L) to Ω1 (W )|L , then use the isomorphism provided by ω to get a vector field along L. Equivalently, using the metric on W , we may think of one-forms on L as tangent vectors to L, then apply the complex structure J on W to give W -vector fields on L. We denote this map from one-forms to normal vector fields by
Ω1 (L) → T W |L ,

σ 7→ σ p ω −1 .

(2.2)

Connections on L are carried along by the vector field to connections on nearby cycles, and we are identifying the space of U (1) connections with iΩ1 (L; R).
There is a natural almost complex structure on A, acting as


0 1
J=
,
−1 0

5

with respect to the splitting (2.1) of the tangent spaces. With respect to this we claim to have the following holomorphic functional
Z L
Z L
Z L
2
2
2
fC (A, L) =
(F + ω) =
(F + ω ) + 2
ω ∧ F.
L0

L0

L0

Here we have extended A to a connection on the trivial bundle on the whole of W
(restricting to a fixed connection A0 on L0 , and to A on L) and taken its curvature form F . We have again picked a 4-cycle bounding L − L0 ; because F and ω are closed the resulting functional is independent of different homologous choices of the
4-cycle, and in general well defined up to the addition of some discrete periods.
R L It2
is also (again) independent of hamiltonian isotopies of L. Notice that the L0 F
term is just the real Chern-Simons functional CSR of the connection A on L, whose critical points are well known to be flat connections. As pointed out to me by Eric
Zaslow, the real and complex Chern-Simons functionals already appear in [W1] and
[Va] as possible mirror partners (this is partially justified in [LYZ]), but without the terms in the symplectic form (and including instanton corrections from holomorphic discs which we are ignoring for our rough analogy). Asking for a real function to be equal to a complex one is possible when one restricts attention to a real slice such as the space of Lagrangian submanifolds in A; deforming within this space the imaginary part of fC remains constant and it reduces to CSR . But allowing the imaginary counterparts to these real deformations the right functional to consider is fC . Notice also that if ω/2π is integral, so that we can pick a connection B
with curvature −iω, then the action functional can be written in the more familiar looking Chern-Simons form
Z
Z
CdC
fC (A, L) = (B + iA) ∧ d(B + iA) =
L

L

for the ‘complexified connection’ C = B + iA (a C × -connection, instead of a U (1)connection.) This makes more contact with the physics literature and allows one to extend the identification of CSR and CSC in [LYZ] to non Lagrangian sections, giving complex values. Tian has informed me that he and Chen have also considered the functional fC (A, L) [Ch].
Mirror symmetry should relate Lagrangians not just to bundles but the whole derived category. For Riemann surfaces C ⊂ M , for instance, there is a functional in [DT], [W2] rather like fR above:
Z C
Ω
C0

is formally holomorphic and has as critical points the holomorphic curves C. Similarly for four-manifolds S ⊂ M with connections on them the following functional
6

(formally similar to fC )

Z S
S0

tr F ∧ Ω

has critical points the holomorphic surfaces with flat connection on them. Alternatively, as CSC is additive under extensions of bundles it does extend to the derived category. (Whether these two approaches are compatible; i.e. whether or not the functional associated to a curve or surface is the same as CSC applied to a locally free resolution of its structure sheaf, up to a constant, seems to not have been worked out.)
That fC is holomorphic follows from the computation that the derivative
R of fC
1
down a ∈ Ω (L) ⊕ 0 (that only changes the connection A 7→ A + δa) is L 2F ∧
ia + 2ω ∧ ia, while the derivative down −Ja ∈ 0 ⊕ Ω1 (L), i.e. down the vector field
R
a p ω −1 , is L 2ω ∧ a + 2a ∧ F . The second expression is −i times the first, so the derivative is complex linear and f is holomorphic. Equivalently we are saying that dfC is the 2-form
2i (F + ω) ⊕ 2 (F + ω),

which pairs with the tangent space (2.1) by integration over L to give a form of type
(1,0) on (2.1).
Thus critical points of the functional are Lagrangian cycles with flat line bundles on them: exactly the basic building blocks of the objects proposed in [K] to be mirror dual to the holomorphic bundles that are the critical points of CS. So this ties in three well known moduli problems of virtual dimension zero (i.e. with deformation theories whose Euler characteristic vanishes) – flat bundles on 3-manifolds, holomorphic bundles on Calabi-Yau 3-folds, and Lagrangians (up to hamiltonian deformation) in symplectic 6-manifolds.
So as mirror to [T1] one would like to count Lagrangians (up to hamiltonian deformations) plus flat line bundles on them, and this is what Joyce’s work [J]
has begun to tackle (in the rigid case of L being a homology sphere). Mirroring precisely the behaviour of the holomorphic Casson invariant this count appears to be independent of deformations of the Kähler form and to have wall-crossing changes as the complex structure varies.

3

Gauge equivalence and moment maps

In fact what Joyce is proposing to count is special Lagrangian spheres with flat line bundles on them (hence the otherwise anomalous dependence on the complex structure), while [T1] counts stable bundles (i.e. by Donaldson-Uhlenbeck-Yau, modulo the technicalities of polystable and non-locally-free sheaves, we count HermitianYang-Mills connections; hence the dependence on the Kähler form). (Tyurin [Ty]
7

was perhaps the first to suggest that the holomorphic Casson invariant should be related by mirror symmetry to the real Casson invariant (here the U (1) Casson invariant) of SLag submanifolds.)
The link should be, of course, that we want to consider holomorphic connections on one side, up to complex gauge equivalence, and Lagrangians on the other side, up to hamiltonian isotopy, and in both cases we try to do this by picking distinguished representatives of equivalence classes by the usual method of symplectic reduction.
Bringing in a Kähler structure on the complex side, we get a moment map for the gauge group action, whose zeros give the HYM equations. Dually, we would like to bring in the holomorphic 3-form on the symplectic (Kähler) side, and get a complex group to act. So again complexify by adding flat line bundles: consider the critical points of the functional f of the last section, i.e. the space
Z = {(L, A) : L ⊂ W is Lagrangian, A is a flat connection on L }
(not up to gauge equivalence). In fact consider this space on a Calabi-Yau manifold
W of any dimension n. It has tangent space
T(L,A) Z = Z 1 (L) ⊕ Z 1 (L)
(Z 1 (L) denotes closed real one-forms on L), the first being tangent to the space of flat connections, the second giving normal vector fields (by contracting with ω −1 )
preserving the Lagrangian condition. We have an obvious almost complex structure


0 1
J=
.
(3.1)
−1 0
Then the real group C ∞ (L; R)/R acts as the Lie algebra to the group of gauge transformations on the flat line bundles (taking d and adding to the connection)
whose complexification C ∞ (L; C )/C acts complex linearly: the imaginary part
C ∞ (L; R)/R acts by hamiltonian deformations through the normal vector field h 7→ dh p ω −1 .
Unfortunately, without using a metric this vector field is only defined up to the addition of tangent vector fields to L; the map (2.2) is really a map to (T W |L )/T L
which we have lifted to T W |L using the metric. (Equivalently we can extend h to a first formal neighbourhood of L in different ways to get a different vector field.)
How we pick this alters how we carry the flat connection along with L, and how the almost complex structure (3.1) acts. For instance suppose we are in the rather artificial case of L being transverse to an SYZ T n -fibration. Then we can carry L and the flat connection up the fibres and identify the functions C ∞ (L) from Lagrangian
8

to Lagrangian using the projection. Thus the group remains constant as L moves
(effectively what we are doing is extending functions from L to a neighbourhood of
L in W by pulling up along the SYZ fibres). This does not work when L branches over the base of such a fibration. One can instead use the metric to define normal vector fields, but then identifying the Lie algebra C ∞ (L) with a fixed C ∞ (L0 ) for all L becomes difficult.
This problem is perhaps not so surprising – the moment the Lagrangian has branching over the base of an SYZ fibration simple explicit correspondences between
Lagrangians and vector bundles (such as [LYZ]) also break down due to our ignoring important holomorphic disc instanton corrections that appear in the physics. For instance recent work of Fukaya, Oh, Ohta and Ono [FO3], surveyed in [Fu1], show these provide the obstructions mirror to those of deformations of holomorphic bundles [T2] – one should not in general consider all (S)Lags (which are unobstructed)
as mirror to holomorphic bundles, but only those whose Floer cohomology (whose definition involves holomorphic discs) is well defined.
However, what is clear is the totality of the group action, even if identifying individual elements causes problems, and this is all we really need. For instance in the K3 (or T 4 ) case one can get the same total group orbit, with a genuine fixed group acting, by hyperkähler rotating a construction due to Donaldson [D]. The end result is that one considers parametrised Lagrangian embeddings f from a Riemann surface L into the K3 such that the pullback of Re Ω is a fixed symplectic form on
L. Then the group of exact symplectomorphisms of
(L, f ∗ Re Ω)
provide a symmetry group of the space of maps f , which also carries a natural Kähler structure. Complexified orbits give hamiltonian deformations, and the moment map is m(f ) = f ∗ Im Ω. The connection with our construction is that after fixing a line bundle η and connection with curvature f ∗ Re Ω, an infinitesimal symplectomorphism
φ induces a flat connection, via parallel transport and pull back, on the bundle
η ⊗ φ∗ η ∗ . Globally the action is different (this action has non-zero Lie bracket, for instance, and a fixed group) but the total group orbit and the moment map (see below) are the same.
In general it is clear that the problem of identifying the group for different embeddings of L should be resolved by working with the space of maps from a fixed
L0 to W , and enlarging the group by including diffeomorphisms of L0 , giving a semi-direct product of Diff(L0 ) and U (1) gauge transformations on L0 . Then the moment map for the diffeomorphism part of the total group would be the Lagrangian condition as in [D], and the problems we are encountering would come from the fact that the group is a semi-direct product and not a product, so that we cannot separate
9

the two out and divide by them separately, as in effect we have been trying to do.
Unfortunately, I have not found the correct formulation of the problem, but it is not so important for follows.
So we shall not worry too much about whether the complex structure defined above is integrable, the group is fixed, or the symplectic structure below is closed.
In 1 complex dimension it is trivial, in 2 we can use Donaldson’s picture, and in
3 dimensions we could either try to use an abstract SYZ fibration to deform and identify Lagrangians transverse to it, or take everything in this section as motivation for finding the stability condition for Lagrangians of the next section.
Fix a homology
R class of Lagrangians and multiply Ω by a unit norm complex number so that L Im Ω = 0. We induce a symplectic structure on Z from J and the following metric on the tangent space
Z
a ∧ ((b p ω −1 ) p Im Ω), ha, bi =
L

for a, b closed 1-forms. A computation in local coordinates shows this is symmetric in a and b; in fact it can be written as
Z
Z
e cos θ (a ∧ ∗b),
(3.2)
a ∧ (b p Re Ω) =
L

L

where e is the isomorphism T ∗ L → T L set up by the induced metric on L, Ω|L =
eiθ volL , and volL the Riemannian volume form on L induced by the Ricci-flat metric.
Thus for Lagrangians with θ ∈ (−π/2, π/2), i.e. those for which Re Ω restricts to a nowhere vanishing volume form on L and so are not too far from being SLag (θ ≡ 0), this gives a non-degenerate metric.
The symplectic form is invariant under the group action, and formally the moment map is indeed m(L, A) = Im Ω in the dual Ωn (L)0 of the Lie algebra (i.e.
n-forms on L with integral zero). This follows from the computation
Z
Z
Z
dh ∧ ((σ p ω −1 ) p Im Ω), h d (X p Im Ω) =
h Im Ω =
X
L

L

L

where X = σ p ω −1 is a normal vector field
R to the Lagrangian L down which we compute the derivative of the hamiltonian L h Im Ω = hm(L, A), hi for the infinitesimal action of h. Here have extended h to a first-order neighbourhood of L ⊂ W so that it is constant in the direction of X = σ p ω −1 . Then the right hand side of the above equation is the pairing using the symplectic form of dh and σ, as required.
Infinitesimally we can see the moment map interpretation very easily, and fitting naturally with the mirror bundle point of view. Deformations of holomorphic connections A modulo complex gauge equivalence are given by a ker ∂¯A /im ∂¯A
10

∗ of the HYM equafirst cohomology group, related to deformations ker ∂¯A ∩ ker ∂¯A
tions (modulo unitary gauge transformations) via Hodge theory, with the moment map equation providing the d∗ = 0 slice to the imaginary part of the linearised group action. Similarly, deformations of Lagrangians are given by closed 1-forms ker d : Ω1 (L; R) → Ω2 (L; R), so that dividing by hamiltonian deformations we get

H 1 (L) = ker d/im d.
If instead of dividing we impose the special condition, we get a ker d∗ slice
H 1 (L) = ker d ∩ ker d∗ , to the (imaginary) deformations (real deformations are given by changing the flat
U (1) connection that can be incorporated into this).

A symplectic example
To motivate a guess at the correct definition of stability for Lagrangians, we expand on an example of Lawlor and Joyce ([J] Sections 6 and 7, building on work of [Ha],
[L]; see also a similar example in [SV] that is studied in [TY]), explaining its relevance to mirror symmetry, and giving a simple example in algebraic geometry that mirrors it.
First define the pointwise phase θ of a submanifold L: we may write
Ω|L = eiθ vol where vol is the Riemannian volume form on L induced by Yau’s Ricci-flat metric
[Y] on W . Thus vol provides a (local) orientation for L, and reversing its sign alters the phase θ by π. A SLag is a Lagrangian with constant phase θ.
At first sight θ is multiply-valued; we always choose it to be a fixed singlevalued function to R, lifting eiθ : L → S 1 and thus providing the Lagrangian with a grading as introduced by Kontsevich [K], [S2]. Thus we only consider Lagrangians of vanishing Maslov class – for a Calabi-Yau this is the winding class π1 (L) → π1 (S 1 )
of the phase map eiθ

L −→ S 1 , which of course vanishes for a SLag. (The definition of grading in [K], [S2] is topological and uses the universal Z-cover of the bundle of Lagrangian Grassmannians; here we first pass to the Z/2 orientation cover of the Grassmannian, choosing an orientation of our Lagrangians, and then use a complex structure to pass to the universal Z-cover of this. The two definitions are of course equivalent.)

11

Similarly we can define a kind of average phase φ = φ(L) of a submanifold (or homology class) L ⊂ W by
Z
Ω = A eiφ(L) ,

L

for some real number A; we then use Re (e−iφ(L) Ω|L ) to orient L. Reversing the sign of A alters the phase by π and reverses the orientation. Again for a graded
Lagrangian L = (L, θ), and we will always implicitly assume a grading, φ(L) is canonically a real number (rather than S 1 -valued). Shifting the grading [ 2n ] : θ 7→
θ + 2nπ gives a similar shift to the phase φ(L).
The terminology comes from the fact that if there is a submanifold in the same homology class as L that is SLag with respect to some rotation of Ω, then it is with respect to e−iφ(L) Ω. Slope, which we define as
Z
1
Im Ω,
µ(L) := tan(φ(L)) = R
L Re Ω L

is defined independently of grading, is monotonic in φ in the range (−π/2, π/2), and is invariant under change of orientation φ 7→ φ ± π. This agrees with the slope of a straight line SLag in the case of T 2 , as featured in [PZ], and we think of it as mirror to the slope of a mirror sheaf, as is shown for tori in [PZ] (see [DFR] for corrections in higher dimensions away from the large complex structure limit).
Joyce describes examples of SLags which we interpret as follows. We have a family of Calabi-Yau 3-folds W t as t ranges through (a small open subset of) the moduli space of complex structures on W with fixed symplectic structure. That is, the holomorphic 3-form Ωt varies with t, but the Kähler form ω is fixed. We also have a family of SLag homology 3-spheres Lt1 , Lt2 ⊂ W t such that Lt1 and L2t intersect at a point. If we choose a rotation of Ωt such that Lt2 always has phase φt2 ≡ 0 (this is possible locally at least; in the family described later it will have to be modified slightly), then we are interested as t varies only in the complex number
Z
t
Ωt = Rt eφ1
Lt1

and its polar phase φ = φt1 ; we plot this (i.e. the projection from the complex structure moduli space to C via this map) in Figure 1.
Then in Joyce’s example, for φ < 0 (and Rt > 0) there is a SLag Lt (of some phase φt ) in the homology class [Lt ] = [Lt1 ]+[Lt2 ], such that as φ ↑ 0, this degenerates to a singular union of SLags of the same phase Lt = Lt1 ∪ Lt2 and then disappears for φ > 0.
Most importantly, where Lt exists as a smooth SLag (φ < 0) we have the slope
(and phase) inequality
µt1 < µt2 ,

i.e. φ = φt1 < φt2 ≡ 0;
12

(3.3)

at t = 0, Lt becomes the singular union of Lt1 and Lt2 , with
µt1 = µt2 (φt1 = φt2 ); then there is no SLag in L’s homology class for
µt1 > µt2 (φt1 > φt2 ), though there is a Lagrangian, of course – the symplectic structure has not changed.
Though we have been using slope µ in order to strengthen the analogy with the mirror (bundle) situation, from now on we shall use only the phase (lifted to R using the grading). While each is monotonic in the other for small phase (as tan φ = µ), slope does not see orientation as phase does; reversing orientation adds ±π to the phase but leaves µ unchanged. This is related to the fact that we should really be working with complexes and so forth on the mirror side (the bundle analogy is too narrow) and changing orientation has no mirror analogue in terms of only stable bundles; it corresponds to shifting (complexes of) bundles by one place in the derived category. While slopes of bundles cannot go past infinity (without moving degree in the derived category at least), for Lagrangians they certainly can, and phase φ continues monotonically upwards as its slope tan φ becomes singular and then negative.
Importantly, we can think of the various SLags as independent of time when thought of as Lagrangians in the fixed symplectic manifold W t :
Lemma 3.4 For t > 0 the SLags Lt are all in the same hamiltonian deformation class. Similarly for Lt1 , Lt2 , and for t < 0.
Proof Now choosing the phase of Ωt such that φ(Lt ) ≡ 0,
Z
Z
d t
Im Ω̇t = 0.
( Im Ω ) =
L
L dt

(3.5)

To show this deformation preserves the hamiltonian class of L, we need to find a corresponding first order hamiltonian deformation dh p ω −1 under which the change in Im Ωt ,
Ldh pω−1 ( Im Ωt )|L = d((dh p ω −1 ) p Im Ωt )|L ,

is −Im Ω̇t |L . But as Re Ωt |L is the induced Riemannian volume form volt on L, this means we want to solve f p volt ) = d(∗dh) = ∆(∗h),
−Im Ω̇t |L = d(J(dh p ω) p Re Ωt |L ) = d(dh

where J is the complex structure and e is the isomorphism T ∗ L → T L set up by the induced metric on L. So the equation has a solution by the Fredholm alternative
13

and (3.5).



Thus for φ > 0 we consider the Lt s as the same as Lagrangian submanifolds
(up to hamiltonian deformation) in the fixed symplectic manifold W t ; it is only the
SLag representative that changes as Ωt varies. We think of this as mirror to a fixed holomorphic bundle in a fixed complex structure, with varying HYM connection as the mirror Kähler form changes.
Lemma 3.6 In the analogous 2-dimensional situation of SLags in a K3 or abelian surface, the obstruction does not occur.
Proof Choose a real path of complex structures W t , t ∈ (−ǫ, ǫ) in complex structure moduli space such that there is a nodal SLag L0 = L01 ∪ L02 in W 0 . Without loss of generality we can choose the phase of Ωt so that both ω and Im Ωt pair to zero on the homology class of L0 . Now hyperkähler rotate the complex structures so that instead the new Re Ωt and Im Ωt pair to zero on the homology class of L0 for all t. L0 is now a nodal holomorphic curve C in the central K3. We can understand deformations of C via deformations of the ideal sheaf JC , with obstructions in
Ext2 (JC , JC ) → H 0,2 (W ) ∼
= C, where the arrow is the trace map and is an isomorphism by Serre duality. Standard deformation theory shows the obstruction is purely cohomological – it is the derivative of the H 0,2 -component of the class
[C] ∈ H 2 (W ) ∼
= H 2,0 (W ) ⊕ H 1,1 (W ) ⊕ H 0,2 (W ).
But we have fixed this to remain zero by the phase condition, so the curve deforms to all t (really we should assume the family is analytic in t here and extend to t ∈ C , or just work with first order deformations). Hyperkähler unrotating gives back a family of SLags.

There is a notion of connect summing Lagrangian submanifolds intersecting in a single point (probably due to Polterovich) – see for instance Appendix A of [S1]
– which we claim gives the smoothings Lt of the singular L0 = L1 ∪ L2 . This follows by comparing the local models [J], [S1] for the Lagrangians; see [TY] where it is studied in more detail for a related purpose, and our conventions (slightly different from those of [S2]) are described. While topologically we are just connect summing L1 and L2 by removing a small 3-ball containing the intersection point from each and gluing the resulting boundary S 3 s (there are two ways, depending on orientation), symplectically the construction does not explicitly use orientations of
14

the submanifolds. (Effectively we are using their relative orientation – the canonical orientation of the sum of the tangent spaces of L1 , L2 at the intersection point given by the symplectic structure.)
Giving L1 and L2 in that order produces a Lagrangian, well defined up to hamiltonian isotopy (this will be shown in Section 4 in more generality; see (4.1)),
L1 #L2 , with the singular union L1 ∪ L2 a limit point in the hamiltonian isotopy class, which is not itself hamiltonian isotopic to L1 #L2 (we have seen a family of hamiltonian deformations which has limit L1 ∪ L2 , but the deformations are singular at this limit).
There is also an obvious notion of graded connect sum, which is in fact what we shall always mean by #. There is a unique grading on L1 compatible with a fixed grading on L2 such that we can give a (continuous) grading to the smoothing
L1 #L2 . In the case of connect summing at multiple intersection points (Section 4)
there is at most one such grading; in general the graded connect sum may not exist.
In n dimensions, if L1 and L2 are graded such that L1 #L2 exists, then on reversing the order of the Li , the graded sum that exists is
L2 #(L1 [2 − n]) in the homology class

[L2 ] + (−1)n [L1 ].

(3.7)

Here L[m] means the graded Lagrangian L with its grading changed by adding mπ
to θ, and the homology class of L1 #L2 is [L1 ] + [L2 ] using the orientations on the
Li s induced by the gradings.
This is closely related, as we shall see, to Joyce’s obstruction, and the lack of it in dimension 2 (Lemma 3.6). In 2 dimensions, L1 #L2 and L2 #L1 are in the same homology class, though by a result of Seidel [S1] not in general in the same hamiltonian isotopy class,
L1 #L2 6≈ L2 #L1 , importantly (we use ≈ to denote equivalence up to hamiltonian deformations). For t > 0 in the above family Lt is in the constant hamiltonian deformation class of
L1 #L2 , for t < 0 it is in the different class of L2 #L1 , and at t = 0 it is L1 ∪ L2 –
in neither class but in the closure of both. (For complex t the symplectic structure is no longer constant like it is for t ∈ R, as one can see by following through the hyperkähler rotation; thus we do not get a contradiction to the above statement by going round t = 0 in C .) In 3 dimensions, however, the corresponding obvious choice for a SLag on the other side of the π1t = 0 wall, L2 #L1 [−1], is in the wrong homology class.

15

In the case that the Li are Lagrangian spheres we can see this by going round the wall
φ(Lt1 ) = 0 ≃ φ(Lt2 ),
(3.8)

and using monodromy. In the 2-dimensional K3 or T 4 case this works as follows.

C =

nR

Ω = Reiφ(L1 )
L1

o
φ(L2 ) < φ(L1 )

2
(L1 #L2 ) ≈ L2 #L1
TL
1

L2 #L1 SLag
φ(L ) = φ(L )
0000000000000
1111111111111
1
0
1

2

L1 #L2 SLag
φ(L1 ) < φ(L2 )

Figure 1:

R


Ω
-space, as Ω on K3 varies, with polar coordinates (R, φ(L1 ))
L1

Consider a disc in complex structure moduli space over which the family of
Kähler K3 surfaces (with constant Kähler form) degenerates at the origin to a K3
with an ordinary double point (ODP) with the Lagrangian L1 ∼
= S 2 as vanishing cycle. A local model is the standard Kähler structure on x2 + y 2 + z 2 = u, over the parameter u in the unit disc in C . Now base-changing by pulling back to the double cover in u, u 7→ u2 , we get the 3-fold x2 + y 2 + z 2 = u2 ,

with a 3-fold ODP which has a small resolution at the origin putting in a holomorphic sphere resolving the central K3 fibre u = 0. Choosing a nowhere-zero holomorphic section Ωu of the fibrewise (2, 0)-forms (using the fact that the relative canonical bundle of either family is trivial), this restricts to zero on the exceptional P1 (which is homologous to the vanishing cycle L1 ). Therefore the function
Z
Ωu
(3.9)
L1

has a simple zero at u = 0, i.e. it vanishes to order 1 in u. (The same expression
√
vanished only as u in the original family with the singular fibre, and as such its
16

sign was not well defined; this is because the class [L1 ] was defined globally only up to the monodromy TL1 [L1 ] = −[L1 ], i.e. up to a sign. In our new family the monodromy action TL21 is trivial on homology so it makes sense to talk about the homology class [L1 ] in any fibre, and (3.9) is single valued.)
Then our loop of complex structures is given by taking the loop u = eit and setting Ωt = Ωeit . Pulling back the Kähler form from the original family, we get a locally trivial fibre bundle of symplectic manifolds over the circle whose monodromy is the Dehn twist TL21 (since the monodromy round the un-base-changed loop is TL1
[S1]). As the family no longer has a singular fibre this monodromy is trivial as a diffeomorphism, but it is a result of [S1], [S2] that as a symplectic automorphism it is non-trivial. This is possible because although the family is a locally trivial bundle of symplectic manifolds over the punctured disc, over u = 0 the symplectic form becomes degenerate since it was pulled back via the resolution map.
Measuring [L1 ] against Ωu as in (3.9) we see a principle familiar in physics
(in issues of ‘marginal stability’, and taught to me by Eric RZaslow) – we detect a monodromy, like the degree 1 map S 1 → C × given by t 7→ L1 Ωt in this example,
R
by counting wall crossing where a certain real part (here L1 Im Ωt , or the phase φt1 )
hits 0 ≃ φt2 .
(Here we can no longer choose the phase of Ω such that φt2 = φ(Lt2 ) ≡ 0 in the whole family, as the homology class of L2 is not preserved in the family:
[TL21 L2 ] = [L2 ] + 2[L1 ].

R
However, for a sufficiently small loop about the ODP, i.e. for | L1 Ω| sufficiently small, this will not affect us much and we can write φt2 ≃ 0: we are only interested in topological information like winding numbers and φt1 crossing the wall at φt2 ≃ 0, which are unaffected by small perturbations.)
So instead of going through the φ(Lt1 ) = φ(Lt2 ) ≃ 0 wall we can go round it. If the loop is sufficiently small we do not encounter any more walls where the homology class [L1 ]+[L2 ] can be split into classes of the same phase to possibly make the SLag a singular union of distinct SLags of equal phase. For instance the wall at phase 0
does not extend past u = 0 to phase φt1 = π (even though there µt1 = 0) – the phase of L1 is not zero but π, and is only zero for L1 with the opposite orientation, so it does not exist as a SLag (e.g. in the hyperkähler rotated situation, we are saying there is no complex curve in L1 ’s homology class to possibly make L the nodal union of L1 and something else, there is only an anti-complex curve). So we really can go round the wall; it ends at u = 0.
So this monodromy description shows that on the other t ↑ 2π side of the wall the SLag deforming L2 ∪ L1 is in the hamiltonian deformation class
TL21 L = TL21 (L1 #L2 ) = TL21 (TL−1
L2 ) ≈ TL1 L2 ≈ L2 #L1 ,
1
17

(3.10)

as claimed (for the above equalities see [S1], [S2]).
Notice that the alternative connect sum description of the above Lagrangian
L2 #L1 = TL21 (L1 #L2 ) ≈ TL21 (L1 )#TL21 (L2 ) ≈ L1 [−2]#TL21 (L2 ),

(3.11)

does not violate the phase inequality to (3.3), as
−2π + ǫ ≈ φ(L1 [−2]) < φ(TL21 (L2 )) ≃ 0.
This is why it is important here to keep track of gradings – assigning the phase ǫ to
φ(TL21 (L1 )) would give the opposite inequality, but one would not be able to form the above graded connect sum without also shifting the phase of TL21 (L2 ) by −2π.

C =

nR

L1

iφ(L1 )

Ω = Re

o
φ(L2 ) < φ(L1 )

TL1 (L1 #L2 ) ≈ L2

L2 SLag
φ(L ) = φ(L )
0000000000000
1111111111111
1
0
1

2

L1 #L2 SLag
φ(L1 ) < φ(L2 )

Figure 2:

R


Ω
-space, as Ω on a 3-fold varies, with polar coordinates (R, φ(L1 ))
L1

The 3-fold case (which Dominic Joyce has also, independently, considered) is slightly different; we need only take a single Dehn twist TL1 corresponding to the local family x2 + y 2 + z 2 = u, over u ∈ C to get a winding number one loop in the phase of L1 . This is because
TL1 L1 ≈ L1 [1 − n]
in dimension n, so in 3 dimensions the homology class [L1 ] is preserved instead of being reversed. The corresponding picture is displayed in Figure 2.

18

Again there is a SLag on the other side of the φ = 0 wall, but it is in the wrong homology class [L2 ]:
TL1 L ≈ L2 .
(3.12)
Analogously to (3.11) this has a number of decompositions as connect sums induced by monodromy,
TL1 (L1 #L2 ) ≈ L1 [−2]#(L2 #(L1 [−1])) ≈ L2 ≈ (L1 #L2 )#(L1 [ 1 ]), none of which violate the phase inequality (3.3). The only other obvious choice for a (S)Lag on the other side of the φ = 0 wall (given the K3 result) is TL21 (L1 #L2 ) ≈
L2 #(L1 [−1]); this however is also in the wrong homology class, and in any case does violate (3.3) and so, by Joyce’s analysis, should not be represented by a SLag.
Thinking of TL21 as rotating through −4π in Figure 2, it is at roughly −3π that the phase inequality (3.3) gets violated, and the −π rotation of L2 splits as a SLag into the union of the −π rotations of (L1 #L2 ) and L1 [ 1 ]: these both have phase approximately zero.

A holomorphic bundle example
These phenomena are similar to wall-crossing in bundle theory on the complex side
– in a real one-parameter family of Kähler forms, for fixed complex structure, stable holomorphic bundles for t > 0 can become semistable at t = 0 and unstable for t < 0.
An example that mirrors Joyce’s is the following. Suppose we have two stable bundles (or coherent sheaves) E1 and E2 with
Ext1 (E2 , E1 ) ∼
= C.
This is H 1 (E1 ⊗E2∗ ) in the case of bundles and is the mirror [K] of the one dimensional
Floer cohomology HF ∗ (L2 , L1 ) ∼
= C that is defined by the single intersection point of L1 and L2 (see Section 4 for more details of this, and an explanation of why we are dealing with Ext1 and HF 1 here). We then form E from this extension class
0 → E1 → E → E2 → 0.

(3.13)

Take a family of Kähler forms ω t such that µt (E2 )−µt(E1 ) is the same sign as t (here
µt (F ) = c1 (F ) . (ω t )n−1 / rk (F ) is the slope of F with respect to ω t ). Supposing that the Ei are stable for all t ∈ (−ǫ, ǫ), we claim that E is stable for sufficiently small t > 0, while it is destabilised by E1 for t ≤ 0. Without loss of generality take
µt (E2 ) = µ fixed, and µt (E1 ) = µ − t. As E2 is stable, for t sufficiently small there
19

are no subsheaves of E2 of slope greater than µ − t, so for any stable destabilising subsheaf F of E, the composition
F ֒→ E → E2
cannot be an injection (unless it is an isomorphism, but (3.13) does not split. So
F ∩ E1 6= 0, and the quotient Q = F/(F ∩ E1 ) has slope µ(Q) > µ(F ) > µ − t by the stability of F and instability of E. But Q injects into E2 , which we know is impossible.
In the 2-dimensional case, by Serre duality Ext1 (E1 , E2 ) ∼
= C on
= Ext1 (E2 , E1 )∗ ∼
4
K3 or T , so for t < 0 we can instead form an extension
0 → E2 → E ′ → E1 → 0,

(3.14)

to give a new bundle E ′ which is also stable, and has the same Mukai vector v(E ′ ) = v(E1 ) + v(E2 ); compare (3.7). At t = 0 we take the (polystable) bundle
E1 ⊕ E2 .
This is because the semistable extension (3.13) no longer admits a Hermitian-YangMills metric, but E1 ⊕ E2 does. Also, the algebraic geometry of the moduli problem shows that while a semistable bundle gets identified in the moduli space with the other (“S-equivalent”) sheaves in the closure of its gauge group orbit, there is a distinguished representative of its equivalence class – the polystable direct sum (of the Jordan-Hölder filtration, which here is E1 ⊕ E2 ).
Thus, while the HYM connections vary, the bundle has only 3 different holomorphic structures – for t > 0, t = 0, and t < 0. Put another way (to spell out the analogy with the Lagrangians Lt , L1 , L2 ) as ωt varies with t > 0 we take different points in a fixed complexified gauge group orbit, and at t = 0 we take as limit point something in a different orbit that is nonetheless in the closure of the t > 0 (and t < 0) orbit. The stable deformations of the polystable E1 ⊕ E2 (which we are thinking of as the mirror of the singular union L1 ∪ L2 , of course) are precisely (3.13) for t > 0 and (3.14) for t < 0.
In the 3-fold case, however, Serre duality gives Ext2 (E1 , E2 ) ∼
=
= Ext1 (E2 , E1 )∗ ∼
C instead, and so no stable extension (3.14). In fact one would expect there to be no stable bundle with the right Chern classes; instead the one dimensional Ext2
gives us a complex E ′ in the derived category D b (M ) fitting into an exact sequence of complexes
0 → E2 → E ′ → E1 [−1] → 0,
20

where E1 [−1] is E1 shifted in degree by one place to the right as a complex. This has Mukai vector v(E ′ ) = v(E2 ) − v(E1 ), compare (3.7). Thus, just as in the case of SLags, as we pass through t = 0 there is no natural stable object on the other side in the same homology class in 3 dimensions
(though there is in 2 dimensions) and so an element of the appropriate moduli space disappears.
In fact, as in the Lagrangian example, the natural stable object on the other side of the wall is E2 if we consider monodromy. The mirror of the symplectic Dehn twists of above are described in [ST] (in the case that the bundles Ei are spherical in the sense of [ST]: Extk (Ei , Ei ) ∼
= H k (S n ; C ); this is the natural mirror analogue of the Li s being spheres). These are the twists TE1 of [ST] on the derived category of the Calabi-Yau that act on the extension bundle E of (3.13) to give precisely the extension (3.14),
TE2 1 E = E ′
(compare (3.10)), as a short calculation using [ST] shows. Similarly
TE1 E = E2 , the analogue of (3.12). (In both of these calculations it is important to calculate this monodromy in the derived category; in the K3 case the action of TE2 1 is trivial on K-theory and cohomology, and we cannot distinguish between (3.13) and (3.14), but they are very different as holomorphic bundles and as elements of the derived category.)
The mirror wall crossing, with a SLag splitting into two and then disappearing, is interpreted in [DFR] (and in [SV] in a different case) as the state it represents decaying as we reach a point of ‘marginal stability’. Despite this dealing with only
SLags (and so with only a priori stable Lagrangians in our mathematical sense of stability), this suggestive language does in fact have something to say about the stability, in our sense of group actions, of (non-special) Lagrangians, by considering the nodal limit L1 ∪ L2 to be a semistable Lagrangian.
Thus the Lagrangian L1 #L2 (which always exists as a Lagrangian as the complex structure varies with fixed Kähler form) becomes semistable at t = 0 and is represented by something in a different orbit of the hamiltonian deformation symmetry group (but in the closure of the original orbit), and is unstable for t < 0 so exists there only as a Lagrangian and not as a SLag. This, and the bundle analogue described above, leads us to think of the Lagrangian L1 as destabilising L = L1 #L2
when φ(L1 ) ≥ φ(L2 ). This motivates the now obvious definition of stability in
Section 5; first we explain more about the connections to mirror symmetry, and generalisations to connect sums at more intersection points.
21

4

Relationship to Kontsevich’s mirror conjecture

The inspiration behind most of this paper is of course Kontsevich’s mirror conjecture
[K]. In particular, Kontsevich proposes that the graded vector spaces Ext∗ and HF ∗
should be isomorphic for mirror choices of bundles Ei and graded Lagrangians Li
(or more exotic objects in their derived categories)
HF ∗ (L2 , L1 ) ∼
= Ext∗ (E2 , E1 ); this corresponds to the equality of (graded) morphisms on both sides. Here HF ∗
is Floer cohomology [Fl] – a symplectic refinement of the intersection number of L1
and L2 – which can be Z-graded for graded Lagrangians [S2], whenever it is defined
[FO3], [Fu1]. (More precisely it is the cohomology of a chain complex built out of the free vector space generated by the intersection points, with the differential defined by counting holomorphic discs with boundary in the Lagrangians running from one intersection point to another.) In mirror symmetry, and so in this paper, one should only really consider those Lagrangians whose Floer cohomology is well defined [Fu1].
Thus the point of intersection of the L1 and L2 of the last section define the
Floer cohomology HF ∗ (L2 , L1 ) ∼
= C , and the grading of [S2] is designed specifically so that L1 #L2 can be graded precisely when the relative gradings of the Li force the
Floer cohomology to be concentrated in degree 1; HF ∗ (L2 , L1 ) = HF 1 (L2 , L1 ). We then think of the connect sum L1 #L2 as being mirror to the extension (3.13) defined by Ext1 (E2 , E1 ) ∼
= C . Fukaya, Seidel, and perhaps others have also proposed that
Lagrangian connect sum should be mirror to extensions [Fu2], [S3].
We also consider connect sums of Lagrangians intersecting at n points pi . Then the connect sum is not unique up to hamiltonian deformation: H 1 is added to the
Lagrangian as loops between the intersection points, giving addition

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
