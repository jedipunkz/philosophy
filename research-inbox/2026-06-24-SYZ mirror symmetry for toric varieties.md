---
source: "https://arxiv.org/abs/1412.7231v3"
title: "SYZ mirror symmetry for toric varieties"
author: "Kwokwai Chan"
year: "2014"
publication: "arXiv preprint / math.SG"
download: "https://arxiv.org/pdf/1412.7231v3"
pdf: "https://arxiv.org/pdf/1412.7231v3"
captured_at: "2026-06-24T11:11:32Z"
updated_at: "2026-06-24T11:11:32Z"
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

# SYZ mirror symmetry for toric varieties

- 著者: Kwokwai Chan
- 年: 2014
- 掲載情報: arXiv preprint / math.SG
- 情報源: [arxiv](https://arxiv.org/abs/1412.7231v3)
- ダウンロード: https://arxiv.org/pdf/1412.7231v3
- PDF: https://arxiv.org/pdf/1412.7231v3

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We survey recent developments in the study of SYZ mirror symmetry for compact toric and toric Calabi-Yau varieties, with a special emphasis on works of the author and his collaborators.

## PDF Text

SYZ MIRROR SYMMETRY FOR TORIC VARIETIES

arXiv:1412.7231v3 [math.SG] 28 Jan 2019

KWOKWAI CHAN

Abstract. We survey recent developments in the study of SYZ mirror symmetry for compact toric and toric Calabi-Yau varieties, with a special emphasis on works of the author and his collaborators.

Contents
1. The SYZ proposal
1.1. Semi-flat SYZ: a toy example
2. A quick review on toric varieties
3. Mirror symmetry for compact toric varieties
3.1. Landau-Ginzburg models as mirrors
3.2. The SYZ construction
3.3. Open mirror theorems
4. Mirror symmetry for toric Calabi-Yau varieties
4.1. Physicists’ mirrors
4.2. The SYZ construction
4.3. Period integrals and disk counting – a conjecture of Gross-Siebert
Acknowledgment
References

1
2
5
6
6
7
11
15
15
17
21
23
23

1. The SYZ proposal
In 1996, drawing on the new idea of D-branes from string theory, Strominger,
Yau and Zaslow [76] made a ground-breaking proposal to explain mirror symmetry geometrically as a Fourier–type transform, known as T-duality.
Conjecture 1.1 (The SYZ conjecture [76]). If X and X̌ are Calabi-Yau manifolds mirror to each other, then
(i) both X and X̌ admit special Lagrangian torus fibrations with sections ρ :
X → B and ρ̌ : X̌ → B over the same base which are fiberwise dual to each other in the sense that regular fibers ρ−1 (b) ⊂ X and ρ̌−1 (b) ⊂ X̌ over the same point b ∈ B are dual tori, and
(ii) there exist Fourier–type transforms which exchange symplectic-geometric
(resp. complex-geometric) data on X with complex-geometric (resp. symplecticgeometric) data on X̌.
This conjecture has since become one of the two major mathematical approaches in the study of mirror symmetry (the other being Kontsevich’s homological mirror
Date: January 29, 2019.
1

2

K. CHAN

symmetry conjecture [54]); it provides a concrete and beautiful picture depicting the geometry underlying mirror symmetry.
The SYZ conjecture, in particular, suggests a geometric construction of mirrors, namely, by taking the fiberwise dual of a special Lagrangian torus fibration on a given Calabi-Yau manifold. It turns out that such a construction has to be modified by quantum (or instanton) corrections except in the semi-flat case (which will be reviewed below).
By SYZ mirror symmetry we mean the investigation of mirror symmetry along the lines of the SYZ proposal. The aim of this note is to review recent developments in SYZ mirror symmetry for toric varieties, including SYZ constructions of their mirrors and various open mirror theorems which lead to an enumerating meaning of (inverse of) mirror maps.
1.1. Semi-flat SYZ: a toy example. In the so-called semi-flat case where special
Lagrangian torus fibrations do not admit any singular fibers, the SYZ picture is particularly appealing. To describe this, we need the following
Definition 1.2. An affine structure on an n-dimensional smooth manifold is an atlas of charts such that the transition maps lie in the affine linear group Aff(Rn ) =
Rn oGLn (R). The affine structure is called tropical (resp. integral) if the transition maps lie in Rn o GLn (Z) (resp. Aff(Zn ) = Zn o GLn (Z)).
Given a special Lagrangian torus fibration ρ : X → B without singular fibers,1
classic results of McLean [71] give us two naturally defined tropical affine structures on the base B: the symplectic and complex affine structures. In these terms, mirror symmetry can be explained neatly as an interchange of these structures.
To describe these structures, let Lb := ρ−1 (b) be a fiber of ρ. A tangent vector v ∈ Tb B can be lifted to a vector field (also denoted by v by abuse of notations)
normal to the fiber Lb , and this determines the differential forms
α := −ιv ω ∈ Ω1 (Lb ; R),
β := ιv Im Ω ∈ Ωn−1 (Lb ; R), where ω and Ω denote the Kähler form and holomorphic volume form on the CalabiYau manifold X respectively.
McLean [71] proved that the corresponding deformation of Lb gives special Lagrangian submanifolds if and only if both α and β are closed. By identifying the tangent space Tb B with the cohomology group H 1 (Lb ; R) by v 7→ −ιv ω, we obtain the symplectic affine structure on B; similarly, identifying Tb B with the cohomology group H n−1 (Lb ; R) by v 7→ ιv Im Ω produces the complex affine structure on
B. We also have the McLean metric on B defined by
Z
g(v1 , v2 ) := −
ιv1 ω ∧ ιv2 Im Ω.
Lb

In a very interesting paper [46], Hitchin explains how these tropical affine structures are related in an elegant way by the Legendre transform. Let x1 , . . . , xn be local affine coordinates on B with respect to the symplectic affine structure, then
1In this note, a Lagrangian submanifold L ⊂ X is called special if Im Ω| = 0, where Ω is the
L
holomorphic volume form on X.

SYZ FOR TORIC VARIETIES

3

the McLean metric can be written as the Hessian of a convex function φ on B, i.e.


∂
∂
∂2φ
g
,
=
.
∂xj ∂xk
∂xj ∂xk
Hitchin observed that setting x̌j := ∂φ/∂xj (j = 1, . . . , n) gives precisely the local affine coordinates on B with respect to the complex affine structure. Furthermore, let n
X
φ̌ :=
x̌j xj − φ(x1 , . . . , xn )
j=1

be the Legendre transform of φ. Then we have xj = ∂ φ̌/∂ x̌j and


∂
∂ 2 φ̌
∂
,
=
.
g
∂ x̌j ∂ x̌k
∂ x̌j ∂ x̌k
Now suppose that the fibration ρ : X → B admits a Lagrangian section. Then a theorem of Duistermaat [22] gives global action-angle coordinates and hence a symplectomorphism
X∼
= T ∗ B/Λ∨ , where Λ∨ ⊂ T ∗ B denotes the lattice locally generated by dx1 , . . . , dxn , and the quotient T ∗ B/Λ∨ is equipped with the canonical symplectic structure
ωcan :=

n
X

dxj ∧ dy j ,

j=1

where y , . . . , y are fiber coordinates on T ∗ B.
The SYZ proposal suggests that the mirror of X is the total space of the fiberwise dual of ρ : X → B, which is nothing but the quotient
1

n

X̌ := T B/Λ, where Λ ⊂ T B is the dual lattice locally generated by ∂/∂x1 , . . . , ∂/∂xn . Geometrically, X̌ should be viewed as the moduli space of pairs (L, ∇) where L is a fiber of ρ and ∇ is a flat U (1)-connection over L; this is in fact the key idea lying at the heart of the SYZ proposal.
The quotient X̌ = T B/Λ is naturally a complex manifold with holomorphic coordinates given by zj := exp(xj + iyj ), where y1 , . . . , yn are fiber coordinates on T B dual to y 1 , . . . , y n . So the SYZ approach constructs the mirror of X as a complex manifold equipped with a nowhere vanishing holomorphic volume form
Ω̌ := d log z1 ∧ · · · ∧ d log zn and a dual torus fibration ρ̌ : X̌ → B.
In this case, one can write down an explicit fiberwise Fourier-type transform, which we call the semi-flat SYZ transform F semi-flat , that carries eiω to Ω̌. Indeed, by straightforward computations, we have
Z
P
j
Ω̌ =
eiω ∧ ei j dyj ∧dy ,
ρ−1 (b)
Z
P
j eiω =
Ω̌ ∧ e−i j dyj ∧dy .
ρ̌−1 (b)

4

K. CHAN

These can naturallyPbe interpreted as Fourier-Mukai–type transforms because the j
differential form ei j dyj ∧dy is the Chern character for a universal connection on the fiberwise Poincaré bundle P on X ×B X̌:
Proposition 1.3. We have

Ω̌ = F semi-flat eiω ,
−1 
eiω = F semi-flat
Ω̌ , where the semi-flat SYZ transforms are defined as
F semi-flat (−) := π2∗ (π1∗ (−) ∧ ch(P)) ,
−1

F semi-flat
(−) := π1∗ π2∗ (−) ∧ ch(P)−1 ; here πj denotes projection to the j th -factor.
This explains one-half of mirror symmetry in the semi-flat case, namely, the correspondence between the A-model on X and B-model on X̌. If we now switch to the complex affine structure on B, then the symplectic structure
X
¯ =
ω̌ := i∂ ∂φ
φjk dxj ∧ dyk , j,k
2

φ
, on X̌ is compatible with its complex structure, and thus X̌
where φjk = ∂x∂j ∂x k
becomes a Kähler manifold.
On the other side, we obtain a complex
Pnstructure on X with holomorphic coordinates wi , . . . , wn defined by d log wj = k=1 φjk dxk + idy j , so X also becomes a
Kähler manifold, and is equipped with the nowhere vanishing holomorphic n-form

Ω := d log w1 ∧ · · · ∧ d log wn .
The maps ρ : X → B and ρ̌ : X̌ → B now give fiberwise dual special Lagrangian torus fibrations.
Remark 1.4. When the function φ above satisfies the real Monge-Ampère equation


∂2φ
det
= constant,
∂xj ∂xk we obtain T n -invariant Ricci-flat metrics on both X and its mirror X̌. In this case, the McLean metric on B is called a Monge-Ampère metric and B will be called a
Monge-Ampère manifold. This explains why SYZ mirror symmetry is intimately related to the study of real Monge-Ampère equations and affine Kähler geometry.
We refer the interested readers to the papers [56, 69, 70] and references therein for the detail story.
The SYZ conjecture therefore paints a beautiful picture for mirror symmetry in the semi-flat case; see [64, 62] for other details on semi-flat SYZ mirror symmetry
(cf. also [16, Section 2] and [2, Chapter 6]). Unfortunately, this nice picture remains valid only at the large complex structure/volume limits where all quantum corrections are suppressed. Away from these limits, special Lagrangian torus fibrations will have singular fibers and the SYZ mirror construction must be modified by instanton corrections which, as suggested again by SYZ [76], should come from holomorphic disks bounded by the Lagrangian torus fibers.

SYZ FOR TORIC VARIETIES

5

2. A quick review on toric varieties
Before discussing mirror symmetry for toric varieties, let us set up some notations and review certain basic facts in toric geometry which we need later. We will follow standard references such as the books [30, 19].
Let N ∼
= Zn be a rank n lattice. We denote by M := Hom(N, Z) the dual of N
and by h·, ·i : M × N → Z the natural pairing. For any Z-module R, we also let
NR := N ⊗Z R and MR := M ⊗Z R.
Let X = XΣ be an n-dimensional toric variety defined by a fan Σ supported in the real vector space NR = N ⊗Z R. The variety X admits an action by the algebraic torus TNC := N ⊗Z C× ∼
= (C× )n , whence its name “toric variety”, and it contains a Zariski open dense orbit U0 ⊂ X on which TNC acts freely.
Denote by v1 , . . . , vm ∈ N the primitive generators of rays in Σ, and by
D1 , . . . , D m ⊂ X
the corresponding toric prime divisors. We have the fan sequence associated to X:
φ

0 −→ K := Ker(φ) −→ Zm −→ N −→ 0, where the fan map φ : Zd → N is defined by φ(ei ) = vi , and {e1 , . . . , em } is the standard basis of Zd . Applying Hom(−, Z) to above gives the divisor sequence:
(2.1)

0 −→ M −→

m
M

ZDi −→ K∨ −→ 0.

i=1

In this note, we will always assume that X = XΣ is smooth2 and the support
|Σ| of the fan Σ is convex and of full dimension in NR . In this case, we have
K∼
= H2 (X; Z), K∨ ∼
= Pic(X) = H 2 (X; Z), and both are of rank r := m − n. We let {p1 , . . . , pr } be a nef basis of H 2 (X; Z)
and let {γ1 , . . . , γr } ⊂ H2 (X; Z) be the dual basis.
Definition 2.1. The complex moduli of the mirror of X = XΣ is defined as
M̌B := K∨ ⊗Z C× .
We will denote by t = (t1 , . . . , tr ) the set of coordinates on M̌B ∼
= (C× )r with
2
respect to the nef basis {p1 , . . . , pr } ⊂ H (X; Z), so that t = 0 is a large complex structure limit of the mirror complex moduli.
Suppose further that X = XΣ is Kähler. Recall that the Kähler cone of a Kähler manifold X is the open cone of Kähler classes in H 2 (X; R).
Definition 2.2. The complexified Kähler moduli of X = XΣ is defined as

MA := KX ⊕ 2πi H 2 (X; R)/H 2 (X; Z) ⊂ K∨ ⊗Z C× , where KX ⊂ H 2 (X; R) denotes the Kähler cone of X.
We will denote by q = (q1 , . . . , qr ) ∈ MA the complexified Kähler parameters defined by
 Z

qa = exp −

ωC ,
γa

2Most of the results described here have natural generalization to toric orbifolds, i.e. toric varieties with at most quotient singularities. We will make comments where appropriate.

6

K. CHAN

where ωC = ω +2πiβ ∈ H 2 (X; C) is a complexified Kähler class; here the imaginary part β ∈ H 2 (X; R) is the so-called B-field.
A variety X is called Calabi-Yau if it is Gorenstein and its canonical line bundle
KX is trivial. By this definition, a toric variety X = XΣ is Calabi-Yau if and only if there exists a lattice point u ∈ M such that hu, vi i = 1 for i = 1, . . . , m [19]. This is in turn equivalent to the existence of u ∈ M such that the corresponding character
C
χu ∈ Hom(TM
, C× ) defines a holomorphic function on X with simple zeros along each toric prime divisor Di and which is non-vanishing elsewhere. In such a case,
X is necessarily noncompact.
3. Mirror symmetry for compact toric varieties
3.1. Landau-Ginzburg models as mirrors. Through the works of Batyrev [5],
Givental [31, 32, 33], Kontsevich [55], Hori-Vafa [48] and many others, mirror symmetry has been extended beyond the Calabi-Yau setting. In that situation, the mirror geometry is believed to be given by a Landau-Ginzburg model which is a pair
(X̌, W ) consisting of a Kähler manifold X̌ and a holomorphic function W : X̌ → C
called the superpotential of the model [78, 79]. Note that when W is nonconstant, the mirror manifold X̌ is necessarily noncompact.
Many interesting mathematical consequences can be deduced from this mirror symmetry. For instance, Kontsevich has extended his homological mirror symmetry conjecture to this setting [55]. On the other hand, it is also predicted that genus 0
closed Gromov-Witten invariants of X are encoded in the deformation or unfolding of the superpotential W ; in particular, there should be an isomorphism of Frobenius algebras
(3.1)

QH ∗ (X) ∼
= Jac(W )

between the small quantum cohomology ring QH ∗ (X) of X and the Jacobian ring
Jac(W ) := O(X̌)/JW of W , where JW denotes the Jacobian ideal.
Let X = XΣ be an n-dimensional compact toric Kähler manifold defined by a fan Σ in NR ∼
= Rn . The mirror of X is conjecturally given by a Landau-Ginzburg
C
:= M ⊗Z C× ∼
model (X̌, W ), where X̌ is simply the algebraic torus TM
= (C× )n so that the Jacobian ring is simply
(3.2)

C[z1±1 , . . . , zn±1 ]
E.
Jac(W ) = D
∂W
z1 ∂W
,
.
.
.
, z
n ∂zn
∂z1

In a more canonical way, tensoring the divisor sequence (2.1) with C× gives the exact sequence
C
0 −→ TM
−→ (C× )m −→ M̌B −→ 0,

and the mirror is given by the family of algebraic tori (C× )m → M̌B .
In [48], Hori and Vafa argued using physical arguments that the superpotential should be given by a family of Laurent polynomials WtHV ∈ C[z1±1 , . . . , zn±1 ] over the mirror complex moduli M̌B whose Newton polytope is the convex hull of v1 , . . . , vm .
More concretely, the superpotential can be explicitly written as
WtHV =

m
X
i=1

Či z vi ,

SYZ FOR TORIC VARIETIES

7

1
n where z v denotes the monomial z1v · · · znv for v = (v 1 , . . . , v n ) ∈ N ∼
= Zn , and the coefficients Či ∈ C are constants subject to the constraints

ta =

m
Y

ČiDi ·γa ,

a = 1, . . . , r;

i=1

here recall that t = (t1 , . . . , tr ) are complex coordinates on M̌B ∼
= (C× )r .3 We call
HV
the Hori-Vafa mirror of X.
X̌, Wt
Example 3.1. The fan P
of Pn has rays generated by the lattice points vi = ei for n
i = 1, . . . , n and v0 = − i=1 ei in N , where {e1 , . . . , en } is the standard basis of
N = Zn . So its Hori-Vafa mirror is the family of Laurent polynomials t
WtHV = z1 + · · · + zn +
, z1 · · · zn where (z1 , . . . , zn ) are coordinates on X̌ ∼
= (C× )n .
Example 3.2. For another example, let us consider the Hirzebruch surface Fk =
P(OP1 (k) ⊕ OP1 ) (k ∈ Z≥0 ) whose fan has rays generated by the lattice points v1 = (1, 0), v2 = (0, 1), v3 = (−1, −k), v4 = (0, −1) ∈ N = Z2 .
Its Hori-Vafa mirror is the family of Laurent polynomials
(3.3)

WtHV = z1 + z2 +

t2
t1 tk2
+ , z2
z1 z2k

where (z1 , z2 ) are coordinates on X̌ ∼
= (C× )2 .
In terms of the Hori-Vafa mirror, the isomorphism (3.1) can be established as a consequence of the mirror theorem of Givental [33] and Lian-Liu-Yau [66] in the case when X is semi-Fano, i.e. when the anti-canonical divisor −KX is nef.
3.2. The SYZ construction. It is natural to ask whether the SYZ proposal can be generalized to non–Calabi-Yau cases as well. Auroux’s pioneering work [3] gave a satisfactory answer to this question. He considered pairs (X, D) consisting of a compact Kähler manifold X together with a simple normal crossing anti-canonical divisor D ⊂ X. A defining section of D gives a meromorphic top form on X with only simple poles along D and non-vanishing elsewhere, so it defines a holomorphic volume form on the complement X \ D (making it Calabi-Yau). Hence it makes sense to speak about special Lagrangian submanifolds in X \ D. If there is a special
Lagrangian torus fibration ρ : X \ D → B, we can then try to run the SYZ program to construct a mirror X̌.
When X = XΣ is a compact toric Kähler manifold, a canonical choice of D is the union of all the toric prime divisors D1 , . . . , Dm . By equipping X with a toric
Kähler structure ω, we have an action on (X, ω) by the torus TN := N ⊗Z U (1) ∼
= Tn which is Hamiltonian. The associated moment map
ρ : X → ∆, where ∆ ⊂ MR denotes the moment polytope, is then a nice Lagrangian torus fibration. We view the polytope base ∆ as a tropical affine manifold with boundary
3Scaling of the coordinates z , . . . , z n gives the same superpotential, so the constants
1
Č1 , . . . , Čm are effectively determined by the parameters t1 , . . . , tr .

8

K. CHAN

∂∆. Note that the restriction of ρ to the complement U0 = X \ D is a Lagrangian
T n -fibration without singular fibers. Indeed,
ρ|U0 : U0 → ∆◦ , where ∆◦ denotes the interior of ∆, is a trivial torus bundle and hence admits a section.
As in the semi-flat case, there is then a symplectomorphism
U0 = X \ D ∼
= T ∗ ∆◦ /Λ∨ , where Λ∨ is the trivial lattice bundle ∆◦ × N . The SYZ mirror X̌ is therefore simply given by taking the fiberwise dual:
X̌ = T ∆◦ /Λ, where Λ is the dual lattice bundle ∆◦ × M .
We remark that the SYZ mirror manifold X̌ is a bounded domain inside the
C ∼
algebraic torus T MR /Λ = TM
= (C× )n instead of the whole space [3, 15]. While this is natural from the point of view of SYZ, Hori and Vafa [48] have proposed a renormalization procedure to enlarge the mirror and recover the whole algebraic torus; see e.g. [3, Section 4.2] for a discussion.
How about the superpotential W ? The original SYZ proposal [76] suggests that instanton corrections should stem from holomorphic disks in X with boundaries lying on Lagrangian torus fibers of ρ. In view of this, W should be closely related to counting of such holomorphic disks. This turns out to be a natural guess also from the Floer-theoretic viewpoint. Indeed, the superpotential W is precisely the mirror of Fukaya-Oh-Ohta-Ono’s obstruction chain m0 in the Floer complex of a
Lagrangian torus fiber of the moment map ρ [25, 26, 27, 29], which can be expressed in terms of disk counting invariants or genus 0 open Gromov-Witten invariants, which we now review.
Let L ⊂ X be a Lagrangian torus fiber of the moment map ρ. We fix a relative homotopy class β ∈ π2 (X, L) and consider the moduli space Mk (L; β) of stable maps from genus 0 bordered Riemann surfaces with connected boundary and k boundary marked points respecting the cyclic order which represent the class β.
This moduli space is a stable map compactification of the space of holomorphic disks, i.e. holomorphic embeddings (D2 , ∂D2 ) ,→ (X, L). We call elements in
Mk (L; β) stable disks, the domains of which are in general configurations of sphere and disk bubbles.
Fukaya, Oh, Ohta and Ono [26] showed that the moduli space Mk (L; β) admits a Kuranishi structure with virtual dimension vir. dimR Mk (L; β) = n + µ(β) + k − 3, where µ(β) is the Maslov index of β. They also constructed, using perturbations by torus-equivariant multi-sections, a virtual fundamental chain [Mk (L; β)]vir [26], which becomes a cycle in the case k = 1 and µ(β) = 2 (roughly speaking, this is because µ(β) = 2 is the minimal Maslov index and there cannot be disk bubbles in the domains of stable disks).
Definition 3.3. The genus 0 open Gromov-Witten invariant nβ is defined as nβ := ev∗ ([M1 (L; β)]vir ) ∈ Hn (L; Q) ∼
= Q, where ev : M1 (L; β) → L is the evaluation at the (unique) boundary marked point.

SYZ FOR TORIC VARIETIES

9

It was shown in [26] that the numbers nβ are independent of the choice of perturbations, so they are invariants for the pair (X, L) (hence the name “open
Gromov-Witten invariants”). In general, these invariants are very difficult to compute because the moduli problem can be highly obstructed. We will come back to this in the next section.
In terms of the open Gromov-Witten invariants nβ , we have the following
Definition 3.4. The Lagrangian Floer superpotential is defined as
X
(3.4)
WqLF :=
nβ Zβ .
β∈π2 (X,L)
µ(β)=2

Here, we regard X̌ as the moduli space of pairs (L, ∇) where L is a Lagrangian torus fiber of the moment map ρ and ∇ is a flat U (1)-connection over L, and Zβ
is the function (in fact a monomial) explicitly defined by
 Z

Zβ (L, ∇) = exp − ωq hol∇ (∂β),
β

where hol∇ (∂β) denotes the holonomy of the connection ∇ around the loop ∂β and q = (q1 , . . . , qr ) ∈ MA are complexified Kähler parameters.
Remark 3.5. In general, the sum (3.4) involves infinitely many terms. In order to make sense of it, one needs to regard the qa ’s as formal Novikov variables and regard
(3.4) as a formal power series over the Novikov ring. However, for simplicity of exposition, we will always assume that the power series in this paper all converge so that we can work over C. (We will see later that these series indeed converge in the semi-Fano case.) Then we can regard WqLF as a family of holomorphic functions parameterized by q = (q1 , . . . , qr ) ∈ MA .
In [15], mirror symmetry for toric Fano manifolds was used as a testing ground to see how useful fiberwise Fourier–type transforms, or SYZ transforms, could be in the investigation of the geometry of mirror symmetry.
Before we proceed, let us look at Fourier transforms on a single torus T = V /Λ
more closely. Recall that the semi-flat SYZ transform F semi-flat : Ω0 (T ) → Ωn (T ∗ )
is given by
Z
 P
j semi-flat
F
(φ) =
φ y 1 , · · · , y n ei j dyj ∧dy ,
T

while the usual Fourier transform (or series) gives a function φ̂ : Λ∨ → C on the dual lattice Λ∨ by
Z
P
j
φ̂ (m1 , ..., mn ) =
φ (y1 , · · · , yn ) ei j mj y dy1 ∧ · · · ∧ dyn .
T

It is natural to combine these two transforms to define a transformation which transforms differential forms on T × Λ to those on T ∗ × Λ∨ .
Note that T × Λ is the space of geodesic (or affine) loops in T sitting inside the loop space of T , i.e.
T × Λ = Lgeo T ⊂ LT, and the loop space certainly plays an important role in string theory. We are going to describe how such a transformation F SYZ , called the SYZ transform, interchanges symplectic and complex geometries, with quantum corrections included, when X is a compact toric Fano manifold.

10

K. CHAN

Let us consider the open dense subset U0 := X \ D, which is the union of all the
Lagrangian torus fibers of the moment map. Recall that we have an identification
U0 ∼
= T ∗ ∆◦ /Λ∨ as symplectic manifolds. Next we consider the space
X̃ := U0 × Λ ⊂ LX
of fiberwise geodesic/affine loops in U0 . On X̃, we have the quantum-corrected symplectic structure
ω̃ = ω + Φ, where the values of Φ are the Fourier modes of the Lagrangian Floer superpotential
WqLF , so they are generating functions of the genus 0 open Gromov-Witten invariants nβ in Definition 3.3 which count (virtually) holomorphic disks bounded by the
Lagrangian torus fibers of the moment map ρ.
An explicit SYZ transform F SYZ can then be constructed by combining the semiflat SYZ transform F semi-flat with fiberwise Fourier series, which can be applied to understand the geometry of mirror symmetry:
Theorem 3.6 ([15], Theorem 1.1).
(1) The SYZ transforms interchange the quantum-corrected symplectic structure on X̃ with the holomorphic volume form of the pair (X̌, W ), i.e. we have


F SYZ ei(ω+Φ) = eW Ω̌,
−1 W 
F SYZ
e Ω̌ = ei(ω+Φ) .
(2) The SYZ transform induces an isomorphism
∼
=

F SYZ : QH ∗ (X, ωq ) → Jac WqLF



between the small quantum cohomology ring of X and the Jacobian ring of the Lagrangian Floer superpotential WqLF .4
The second part of this theorem was proved using the idea that holomorphic spheres can be described as suitable gluing of holomorphic disks. More precisely, we pass to the tropical limit (see Mikhalkin [72] for a nice introduction to tropical geometry) and observe that a tropical curve can be obtained as a gluing of tropical disks. For example, a line P1 in P2 is obtained as a gluing of three disks, and this can be made precise in the tropical limit (see Figure 1). This observation was later generalized and used by Gross [38] in his study of mirror symmetry for the big quantum cohomology of P2 via tropical geometry; see also [39].
The first part of the above theorem can in fact be generalized to any compact toric manifold by appropriately defining the function Φ as a power series over certain Novikov rings; the second part was proved for semi-Fano toric surfaces by direct computations in [11], and in general it follows from the main theorem in
Fukaya-Oh-Ohta-Ono [29].
4Note that here W LF is defined only on a bounded domain inside T C ∼ (C× )n but we still
=

q

use the formula (3.2) to define the Jacobian ring Jac WqLF .

M

SYZ FOR TORIC VARIETIES

11
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...

...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
.............................................................................
.
.
....
.
.
.
.
....
.
.
.
.
...
.
.
.
.
...
.....
.....
.....
.....
.....
.
.
.
.
..
.....
.....
.....
.....
.....

•

glued from

• .................................................................................
••

...
.....
.....
.....
.....
.
.
.
.
..
.....
.....
.....
.....
.....
.
.
.
.
.....
.....
.....
.....
.....

Figure 1. A tropical curve as a gluing of tropical disks.
3.3. Open mirror theorems. For semi-Fano toric manifolds, the mirror theorems of Givental [33] and Lian-Liu-Yau [66] imply that there is an isomorphism


HV
QH ∗ (X, ωq ) ∼
,
= Jac Wt(q)
where t(q) = ψ −1 (q) is the inverse of the toric mirror map
ψ : M̌B → MA , t = (t1 , . . . , tr ) 7→ ψ(t) = (q1 (t), . . . , qr (t)), which, by definition, is given by the 1/z-coefficient of the combinatorially defined, cohomology-valued I-function: m Q0
X
Y
Pr k=−∞ (Di + kz)
IX (t, z) = e a=1 pa log ta /z td
,
QDi ·d
(D
+
kz)
i eff k=−∞
i=1
d∈H (X;Z)
2

Qr

where td = a=1 tapa ·d and Di is identified with its cohomology class in H 2 (X). The toric mirror map ψ, which gives a local isomorphism between the mirror complex moduli M̌B and the complexified Kähler moduli MA of X near the large complex structure and volume limits respectively, can also be obtained by solving a system of PDEs known as Picard-Fuchs equations.
On the other hand, instanton corrections can be realized using the genus 0 open
Gromov-Witten invariants nβ . In terms of the Lagrangian Floer superpotential
(Definition 3.4)
X
WqLF =
nβ Zβ ,
β∈π2 (X,L)
µ(β)=2

Fukaya, Oh, Ohta and Ono proved in [29] that there is an isomorphism of Frobenius algebras5

QH ∗ (X, ωq ) ∼
= Jac WqLF .
While the superpotentials W HV and W LF originate from different approaches, they give rise to essentially the same mirror symmetry statement. This leads us to the following conjecture:
5In fact what Fukaya, Oh, Ohta and Ono obtained in [29] is a ring isomorphism between the big quantum cohomology ring of X and the Jacobian ring of the so-called bulk-deformed superpotential
[27]. For the purpose of this note, we restrict ourselves to the small case.

12

K. CHAN

Conjecture 3.7. Let X be a semi-Fano toric manifold, i.e. its anti-canonical divisor −KX is nef. Let W HV and W LF be the Hori-Vafa and Lagrangian Floer superpotentials of X respectively. Then we have the equality
HV
Wt(q)
= WqLF

(3.5)

via the inverse mirror map t(q) = ψ −1 (q).
In [18], Cho and Oh classified all embedded holomorphic disks in X bounded by a fixed Lagrangian torus fiber L of the moment map. Applying this classification, they computed all the invariants nβ in the Fano case:

1 if β = βi for some i = 1, . . . , m, nβ =
0 otherwise.
Here, the βi ’s are the so-called basic disk classes: for each i ∈ {1, . . . , m}, the relative homotopy class βi is represented by an embedded Maslov index 2 holomorphic disk in X with boundary on L which intersects the toric divisor Di at exactly one point in its interior; it can be shown that {β1 , . . . , βm } form a Z-basis of the relative homotopy group π2 (X, L) [18].
The results of Cho and Oh give an explicit formula for W LF and also a direct verification of the above conjecture in the Fano case. Note that in this case the mirror map is trivial. In fact, Cho and Oh proved that nβi = 1 for all i for any compact toric manifold. However, for a general homotopy class β ∈ π2 (X, L) which maybe represented by a stable map with disk and/or sphere bubbles, the invariants nβ are usually very difficult to compute.
Nevertheless, in the semi-Fano case, since the first Chern number of any effective curve is nonnegative, only sphere bubbles can appear in Maslov index 2 stable disks.
Therefore, nβ is nonzero only when β is of the form
β = βi + d, where βi is a basic disk class and d ∈ H2eff (X; Z) is an effective curve class with first Chern number c1 (d) := c1 (X) · d = 0.
In particular, the Lagrangian Floer superpotential is a Laurent polynomial which can be expressed nicely as: m
X
WqLF =
Ci (1 + δi (q))z vi , i=1

where
1 + δi (q) :=

X

nβi +α q α

α∈H2eff (X;Z); c1 (α)=0

is a generating function of genus 0 open Gromov-Witten invariants, and the coefficients Ci ∈ C are constants subject to the constraints m
Y
qa =
CiDi ·γa , a = 1, . . . , r, i=1

where q = (q1 , . . . , qr ) ∈ MA are the complexified Kähler parameters.
This suggests an equivalent formulation of Conjecture 3.7 in terms of the socalled SYZ map:

SYZ FOR TORIC VARIETIES

13

Definition 3.8. The SYZ map is the map
φ : MA → M̌B , q = (q1 , . . . , qr ) 7→ φ(q) = (t1 (q), . . . , tr (q))
defined by ta (q) = qa

m
Y

D ·γa

(1 + δi (q)) i

i=1

for a = 1, . . . , r.
Then Conjecture 3.7 is equivalent to the following
Conjecture 3.9. The SYZ map φ is inverse to the mirror map ψ.
In particular this gives an enumerative meaning to inverses of mirror maps, and explains the mysterious integrality property of mirror maps observed earlier in
[67, 68, 58, 80]
To prove the above conjectures, we need to compute the invariants nβ . The first known non-Fano example is the Hirzebruch surface F2 , whose Lagrangian Floer superpotential was first computed by Auroux [4] using wall-crossing and degeneration techniques (which work for F3 as well) and then by Fukaya-Oh-Ohta-Ono in [28] by employing their big machinery [25]. The result is given by the Laurent polynomial
(3.6)

WqLF = z1 + z2 +

q1 q22
q2 (1 + q1 )
+
.
z1 z22
z2

Comparing (3.3) (for k = 2) and (3.6), we see that, besides the basic disks, there is one extra stable disk whose domain contains a sphere bubble which is a copy of the (−2)-curve in F2 ; see Figure 2 below.

D4
D3

D1
D2

Figure 2. The stable disk in F2 as a union of a disk and the sphere D4 .
In [7], the invariants nβ were computed for any toric P1 -bundle of the form
X = P(KY ⊕ OY ), where Y is a compact toric Fano manifold and KY denotes
(by abuse of notations) the canonical line bundle over Y ; note that this class of examples includes the Hirzebruch surface F2 = P(KP1 ⊕ OP1 ). The computation is done by establishing the following equality [7, Theorem 1.1]:
(3.7)

nβ = h[pt]iX
0,1,β̄ ,

between open and closed Gromov-Witten invariants. The right-hand side of the equality is the genus 0 closed Gromov-Witten invariant h[pt]iX
, where [pt] ∈
0,1,β̄

14

K. CHAN

H 2n (X; Q) is the class of a point and β̄ is the class of the closed stable map obtained by capping off the disk component of a stable disk representing β. This equality, together with the observation that the invariants h[pt]iX
appear in a certain
0,1,β̄
coefficient of Givental’s J-function and also a computation of the mirror map [14], prove both Conjectures 3.7 and 3.9 for the class of examples X = P(KY ⊕ OY ).
In [11], the open Gromov-Witten invariants nβ and hence the Lagrangian Floer superpotential W LF were computed for all semi-Fano toric surfaces. For these surfaces, every toric divisor (curve) has self-intersection number at least −2 (indeed, this is equivalent to the surface being semi-Fano). Stable disks with nontrivial sphere bubbles can only appear along a chain of (−2)-toric curves. The above open/closed equality then allows us to identify the invariant nβ with a certain closed local Gromov-Witten invariant, which was previously computed by Bryan and Leung [6]. We thus obtain explicit formulas for both nβ and W LF ; see [11,
Theorem 1.2 and Appendix A]. As an application, an explicit presentation for the small quantum cohomology ring of such a surface was also obtained; see [11, Section
4] for details.
As an example, consider the semi-Fano toric surface X = XΣ defined by the fan and polytope shown in Figure 3 below.
-2
-2

D5

D4

-1

D6

D3
0

-2 D1

D2
1

Figure 3. The fan Σ (left) and the polytope ∆ (right) defining
X; the number besides each edge is the self-intersection number of the corresponding toric divisor.
Then the superpotential W LF is given by
W LF =

(1 + q1 )z1 + z2 +

q1 q2 q32 q43
q q q2
+ (1 + q2 + q2 q3 ) 1 z32 4
z1 z2

+(1 + q3 + q2 q3 ) q1 zq42 z1 +

q1 z12
z2 ,

where ql = exp(−tl ), l = 1, . . . , 4, are complexified Kähler parameters on MA .
For general semi-Fano toric manifolds, the major difficulty is to find a suitable space to replace X on the right-hand side of the open/closed equality (3.7) so that it remains valid. In [13], it was discovered that a natural choice was given by the so-called Seidel space Ei , an X-bundle over P1 corresponding to the C× -action defined by the divisor Di which has been used by Seidel [74] to construct the Seidel representation – a nontrivial action of π1 (Ham(X, ω)) on QH ∗ (X, ω).
By establishing an open/closed equality similar to (3.7) but with X replaced by
Ei on its right-hand side, and applying various techniques such as certain degeneration formulas, the following theorem was proved in [13].
Theorem 3.10 ([13], Theorem 1.2 and Corollary 6.6). Both Conjectures 3.7 and
3.9 are true for any compact semi-Fano toric manifold.

SYZ FOR TORIC VARIETIES

15

We call this the open mirror theorem. When combined with the closed mirror theorem of Givental [33] and Lian-Liu-Yau [66], this gives the isomorphism
QH ∗ (X, ωq ) ∼
= Jac(W LF ), q

which recovers the aforementioned result of Fukaya, Oh, Ohta and Ono [29].
More importantly, the open mirror theorem gives us explicit formulas which can effectively compute all the genus 0 open Gromov-Witten invariants nβ in the semi-Fano case:
Theorem 3.11 ([13], Theorem 1.1). Let X be a compact semi-Fano toric manifold.
Then we have the following formula for the generating function of genus 0 open
Gromov-Witten invariants:
1 + δi (q) = exp gi (t(q)),
−1

where t(q) = ψ (q) is the inverse of the mirror map q = ψ(t) and gi is a hypergeometric function explicitly defined by
X (−1)(Di ·d) (−(Di · d) − 1)!
Q
td ; gi (t) :=
p6=i (Dp · d)!
d

here the summation is over all effective curve classes d ∈ H2eff (X; Z) satisfying
−KX · d = 0, Di · d < 0 and Dp · d ≥ 0 for all p 6= i.
As another immediate corollary, we also have the following convergence result.
Theorem 3.12 ([13], Theorem 1.3). For a compact semi-Fano toric manifold X, the generating functions 1 + δi (q) are convergent power series, and hence the Lagrangian Floer superpotential WqLF is a family of holomorphic functions parametrized by the complexified Kähler parameters q = (q1 , . . . , qr ) ∈ MA .
We refer the reader to the paper [13] for more details. We remark that an open mirror theorem in the orbifold setting was also proved in [9] for weighted projective spaces of the form P(1, . . . , 1, n) using similar techniques; in fact, Theorem 3.11 can be extended to all semi-Fano toric orbifolds as well [8].
4. Mirror symmetry for toric Calabi-Yau varieties
4.1. Physicists’ mirrors. Let X = XΣ be a toric Calabi-Yau variety of complex dimension n. By choosing a suitable basis of N = Zn , we may write vi = (wi , 1) ∈ N = Zn−1 ⊕ Z, where wi ∈ Zn−1 . Recall that we have assumed that X is smooth and the fan Σ
has convex support. In this case, the Calabi-Yau manifold X is a crepant resolution of an affine toric variety (defined by the cone |Σ|) with Gorenstein canonical singularities, which is equivalent to saying that X is semi-projective [19, p.332].
An important class of examples of toric Calabi-Yau manifolds is given by total spaces of the canonical line bundles KY over compact toric manifolds Y . For example, the total space of KP2 = OP2 (−3) is a toric Calabi-Yau 3-fold whose fan
Σ has rays generated by the lattice points v0 = (0, 0, 1), v1 = (1, 0, 1), v2 = (0, 1, 1), v3 = (−1, −1, 1) ∈ N = Z3 .
Mirror symmetry in this setting is called local mirror symmetry because it originated from an application of mirror symmetry techniques to Fano surfaces (e.g P2 )

16

K. CHAN

contained inside compact Calabi-Yau manifolds and could be derived using physical arguments from mirror symmetry for compact Calabi-Yau hypersurfaces in toric varieties by taking certain limits in the complexified Kähler and complex moduli spaces [51].
Given a toric Calabi-Yau manifold X as above, its mirror is predicted to be a family of affine hypersurfaces in C2 × (C× )n−1 [63, 17, 47]
(
)
m
X
2
× n−1
wi
(4.1)
X̌t = (u, v, z1 , . . . , zn−1 ) ∈ C × (C )
: uv =
,
Či z i=1

where the coefficients Či ∈ C are constants subject to the constraints m
Y
ta =
ČiDi ·γa , a = 1, . . . , r; i=1

here, again, t = (t1 , . . . , tr ) are coordinates on the mirror complex moduli M̌B :=
K ∨ ⊗ Z C× ∼
= (C× )r . X̌t are noncompact Calabi-Yau manifolds since


du ∧ dv ∧ d log z1 ∧ · · · ∧ d log zn−1
Ω̌t := Res
Pm uv − i=1 Či z wi is a nowhere vanishing holomorphic volume form on X̌t .
Example 4.1. When X = KP2 , its mirror is predicted to be

X̌t = (u, v, z1 , z2 ) ∈ C2 × (C× )2 : uv = 1 + z1 + z2 +

t z1 z2


,

where t is a coordinate on the mirror complex moduli M̌B ∼
= C× .
This mirror symmetry has been a rich source of interesting examples and has numerous applications. This has resulted in a flurry of research works by both physicists and mathematicians [63, 17, 47, 36, 37, 77, 52, 35, 49, 50, 23, 24, 53,
75].6 In particular, just like the usual case, it has been applied to make powerful enumerative predictions.
For instance, in the 3-fold case, mirror symmetry was employed to compute the local Gromov-Witten invariants
Z
Ng,d (X) :=
1,
[Mg,0 (X,d)]vir

where Mg,0 (X, d) is the moduli space of genus g, 0-pointed closed stable maps into
X with class d ∈ H2eff (X; Z). These invariants are in general hard to compute due to nontrivial obstructions, but mirror symmetry can be used to compute (at least)
the genus 0 invariants all at once.
To carry out the computations, one needs a mirror map
ψ : M̌B → MA , which is a local isomorphism from the complex moduli M̌B of the mirror X̌ to the complexified Kähler moduli MA of X. Traditionally, mirror maps are defined as quotients of period integrals, which are integrals of the form
Z
Ω̌,
Γ
6This list is certainly not meant to be exhaustive.

SYZ FOR TORIC VARIETIES

17

where Γ ∈ H3 (X̌; Z) is a middle-dimensional integral cycle. These integrals satisfy a system of Picard-Fuchs PDEs [5], whose solutions are explicit hypergeometric–type functions, and indeed this is usually how mirror maps are computed.
For the example of X = KP2 , the Picard-Fuchs system is just a single ODE
 3

Θ + 3tΘ(3Θ + 1)(3Θ + 2) Φ = 0, where Θ = td/dt. A basis of solutions to this equation is explicitly given by
Φ0 (t) = 1,
Φ1 (t) = log t +

∞
X
(−1)k (3k)!
k=1

k

(k!)3

tk ,

2

Φ2 (t) = (log t) + · · · .
Note that constants are the only holomorphic solutions of the Picards-Fuchs equation, and this is true for general toric Calabi-Yau manifolds.
The mirror map is then given by q = ψ(t) := exp (Φ1 (t)/Φ0 (t)) = t(1 + · · · ), and genus 0 local Gromov-Witten invariants N0,d (X) are predicted to appear as coefficients of the Taylor series expansion of Φ2 (ψ −1 (q)) around q = 0.
In [35], it was shown that, via the inverse mirror map q 7→ ψ −1 (q) and a coordinate change, the mirror of KP2 can be rewritten as the family


q
X̌q = (u, v, z1 , z2 ) ∈ C2 × (C× )2 : uv = f (q) + z1 + z2 +
, z1 z2
where
(4.2)

1

f (q) = (ψ(t)/t) 3 = 1 − 2q + 5q 2 − 32q 3 + 286q 4 − 3038q 5 + · · · .

You may wonder why the coefficients of f (q) are all integers and whether they are counting something. We are going to see that these questions can all be answered by applying the SYZ construction, and that the coefficients of f (q) are (virtual)
counts of holomorphic disks.
4.2. The SYZ construction. To carry out the SYZ mirror construction for a toric Calabi-Yau manifold (or more generally an orbifold) X, we need a special
Lagrangian torus fibration on X. This is provided by a construction due independently to Goldstein [34] and Gross [36].
To begin with, recall that the lattice point u ∈ M , which defines the hyperplane containing all the generators vi ’s, corresponds to a holomorphic function χu : X →
C with simple zeros along each toric prime divisor Di ⊂ X. We equip X with a toric
Kähler structure ω and consider the action by the subtorus T n−1 ⊂ TN ∼
= T n which u
preserves χ , or equivalently, the subtorus whose action preserves the canonical holomorphic volume form Ω on X. Let ρ0 : X → Rn−1 be the corresponding moment map which is given by composing the TN -moment map with the projection along the ray in MR spanned by u.
Proposition 4.2 (Goldstein [34], Gross [36]). For any nonzero constant  ∈ C× , the map defined by
ρ := (ρ0 , |χu − |) : X → B := Rn−1 × R≥0 ,

18

K. CHAN

B+
||

Γ

B−
∂B

Figure 4. The base of the Gross fibration on KP2 , which is an upper half space in R3 .
is a special Lagrangian torus fibration, where the fibers are special with respect to the meromorphic volume form
Ω :=

Ω
.
χu − 

We call ρ the Gross fibration, which is non-toric in the sense that it is not the
TN -moment map. Its discriminant locus is easy to describe, namely, a fiber of ρ is non-regular if and only if either
• when it intersects nontrivially with (and hence is contained inside) the hypersurface D ⊂ X defined by χu = , in which case the fiber is mapped to a point on the boundary ∂B = Rn−1 × {0}; or
• when it contains a point where the T n−1 -action is not free, i.e. when at least two of the homogeneous coordinates on X vanish, in which case the fiber is mapped to the image Γ of the codimension 2 subvariety
[
Di ∩ Dj i6=j

under ρ.
We regard B as a tropical affine manifold with boundary ∂B and singularities Γ.
Note that Γ has real codimension 2 in B.
For example, the base B of the Gross fibration on X = KP2 is an upper half space in R3 , and the discriminant locus is a graph Γ which is contained in a hyperplane
H parallel to the boundary ∂B, as shown in Figure 4.
Starting with a (special) Lagrangian torus fibration ρ : X → B on a Calabi-Yau manifold, the SYZ mirror construction proceeds in several steps:
Step 1 Over the smooth locus B0 := B \ (∂B ∪ Γ), the pre-image X0 := ρ−1 (B0 )
can be identified with the quotient T ∗ B0 /Λ∨ by Duistermaat’s action-angle coordinates.
Step 2 Define the semi-flat mirror X̌0 as T B0 /Λ, which is not quite the correct mirror because the complex structure on X̌0 cannot be extended further to any (partial) compactification as the monodromy of the tropical affine structure around the discriminant locus Γ is nontrivial.
Step 3 To obtain the correct and (partially) compactified mirror X̌, we need to modify the complex structure on X̌0 by instanton corrections, which should

SYZ FOR TORIC VARIETIES

19

come from holomorphic disks in X bounded by Lagrangian torus fibers of the fibration ρ.
Such a procedure was pioneered and first put into practice by Auroux in [3, 4].
Among the many interesting and motivating examples studied by him there was
X = Cn , which can be regarded as the simplest example of a toric Calabi-Yau manifold. The work [12] gives a generalization of Auroux’s construction to all toric
Calabi-Yau manifolds, which we now review.
First of all, by definition, the wall(s) inside the base of a Lagrangian torus fibration is the loci of regular fibers which bound nonconstant Maslov index 0
holomorphic or stable disks in X. For the Gross fibration on a toric Calabi-Yau manifold, there is only one wall given by the hyperplane
H := Rn−1 × {||} ⊂ B
parallel to the boundary ∂B. The wall H contains the discriminant locus Γ as a tropical hypersurface, and it divides the base B into two chambers:
B+ := Rn−1 × (||, +∞),
B− := Rn−1 × (0, ||)
over which the Lagrangian torus fibers behave differently in a Floer-theoretic sense.
To explain what is going on, let us consider the genus 0 open Gromov-Witten invariants defined as in Definition 3.3, or in other words, the virtual counts of
Maslov index 2 stable disks in X bounded by fibers of ρ. Here, Maslov index 2
means, geometrically, that the stable disks intersect with the hypersurface D at only one point with multiplicity one. As one moves from one chamber to the other by crossing the wall H, the virtual counts of Maslov index 2 disks bounded by the corresponding Lagrangian torus fiber would jump, exhibiting a wall-crossing phenomenon.
It was proved in [12] that fibers over the chamber B− bound only one (family)
of Maslov index 2 disks, so that the corresponding generating function of open
Gromov-Witten invariants has just one term with coefficient 1; while fibers over
B+ bound as many Maslov index 2 disks as a moment map Lagrangian torus fiber, so that the corresponding generating function of open Gromov-Witten invariants has (possibly infinitely) many terms. The resulting wall-crossing factor is exactly what we need in order to get the correct mirror.
More explicitly, the Lagrangian Floer superpotential WqLF is given by

Pm z0 i=1 (1 + δi (q))Ci z wi over B+ ,
LF
Wq =
u over B− , where
1 + δi (q) =

X

nβi +d q d

α∈H2eff (X;Z)

is a generating function of genus 0 open Gromov-Witten invariants. Here, both z0
and u denote the coordinate associated to the lattice point (0, . . . , 0, 1) ∈ Zn−1 ⊕ Z
1
wn−1
for w =
(over B+ and B− respectively), z w denotes the monomial z1w . . . zn−1
R
(w1 , . . . , wn−1 ) ∈ Zn−1 , q d denotes exp − d ωC which can be expressed in terms of the complexified Kähler parameters q1 , . . . , qr , and β1 , . . . , βm ∈ π2 (X, L) are the basic disk classes as before.

20

K. CHAN

The semi-flat mirror X̌0 = T B0 /Λ in this case can be viewed as a gluing of the two complex charts
X̌+ = T B+ /Λ, X̌− = T B− /Λ
associated to the chambers B+ and B− respectively. As aforementioned, nontrivial monodromy of the tropical affine structure on B0 around the discriminant locus Γ
prevents the complex structure on X̌ from extending any further.
Here comes an important idea of Auroux [3, 4]: Floer theory says that WqLF is an analytic function, so one should modify the gluing between the complex charts
X̌+ , X̌− exactly by the wall-crossing factor: m
X

(1 + δi (q))Ci z wi .

i=1

Such a modification would cancel the nontrivial monodromy of the complex structure around Γ ⊂ B and gives an analytic superpotential on the corrected mirror X̌.
We therefore arrive at the following theorem.7
Theorem 4.3 ([12], Theorem 4.37). The SYZ mirror for the toric Calabi-Yau manifold X is given by the family of affine hypersurfaces8
)
(
m
X
wi
2
× n−1
(1 + δi (q))Ci z
.
(4.3) X̌q = (u, v, z1 , . . . , zn−1 ) ∈ C × (C )
: uv =
i=1

Notice that the SYZ mirror family (4.3) is entirely written in terms of complexified Kähler parameters and disk counting invariants of X, and it agrees, up to a mirror map, with the prediction (4.1) using physical arguments [63, 17, 47].
Example 4.4. The SYZ mirror of X = KP2 is given by

(4.4)
X̌ = (u, v, z1 , z2 ) ∈ C2 × (C× )2 : uv = 1 + δ0 (q) + z1 + z2 +

q z1 z2


,

where q is the Kähler parameter which measures the symplectic area of a projective line contained inside the zero section of KP2 over P2 , and
(4.5)

1 + δ0 (q) =

∞
X

nβ0 +kl q k

k=0

is a generating function of genus 0 open Gromov-Witten invariants, where l ∈
H2 (KP2 ; Z) = H2 (P2 ; Z) is the hyperplane class in P2 , q = exp(−t) and β0 is the basic disk class corresponding to the zero section P2 ⊂ KP2 , which is the only compact toric divisor in KP2 .
By applying the open/closed formula in [7] and a flop formula for closed GromovWitten invariants [65], the genus 0 open Gromov-Witten invariants nβ0 +kl can be expressed in terms of local BPS invariants of KF1 , where F1 is the blowup of P2 at one point [60]. More specifically, if e, f ∈ H2 (KF1 ; Z) = H2 (F1 ; Z)
7We remark that there is no scattering phenomenon in this case because there is only one wall, so no further corrections are needed.
8Strictly speaking, this is the SYZ mirror for the complement X \ D only; the mirror of X

itself is given by the Landau-Ginzburg model (X̌, u).

SYZ FOR TORIC VARIETIES

21

are the classes represented by the exceptional divisor and fiber of the blowup F1 → P2
respectively, then
KF1
nβ0 +kl = N0,0,kf
+(k−1)e , where the right-hand side is the so-called local BPS invariant for the class kf + (k −
1)e ∈ H2 (KF1 , Z). The latter invariants have been computed by Chiang, Klemm,
Yau and Zaslow and listed in the “sup-diagonal” of Table 10 on p. 56 in [17]: nβ0 +l = −2, nβ0 +2l = 5, nβ0 +3l = −32, nβ0 +4l = 286, nβ0 +5l = −3038, nβ0 +6l = 35870,
..
.
Hence, the first few terms of the generating function 1 + δ0 (q) is given by
1 + δ0 (q) = 1 − 2q + 5q 2 − 32q 3 + 286q 4 − 3038q 5 + · · · , which, when compared with the power series (4.2), suggests that the coefficients of f (q) are exactly the virtual counts nβ0 +kl of holomorphic disks. This is indeed the case as we will see in the next section.
We remark that the SYZ construction can be carried out also in the reverse

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
