---
source: "https://arxiv.org/abs/0808.1551v2"
title: "On SYZ mirror transformations"
author: "Kwokwai Chan, Naichung Conan Leung"
year: "2008"
publication: "arXiv preprint / math.SG"
download: "https://arxiv.org/pdf/0808.1551v2"
pdf: "https://arxiv.org/pdf/0808.1551v2"
captured_at: "2026-06-24T11:11:30Z"
updated_at: "2026-06-24T11:11:30Z"
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

# On SYZ mirror transformations

- 著者: Kwokwai Chan, Naichung Conan Leung
- 年: 2008
- 掲載情報: arXiv preprint / math.SG
- 情報源: [arxiv](https://arxiv.org/abs/0808.1551v2)
- ダウンロード: https://arxiv.org/pdf/0808.1551v2
- PDF: https://arxiv.org/pdf/0808.1551v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

In this expository paper, we discuss how Fourier-Mukai-type transformations, which we call SYZ mirror transformations, can be applied to provide a geometric understanding of the mirror symmetry phenomena for semi-flat Calabi-Yau manifolds and toric Fano manifolds. We also speculate the possible applications of these transformations to other more general settings.

## PDF Text

arXiv:0808.1551v2 [math.SG] 30 Jun 2009

ON SYZ MIRROR TRANSFORMATIONS
KWOKWAI CHAN AND NAICHUNG CONAN LEUNG
Abstract. In this expository paper, we discuss how Fourier-Mukai-type transformations, which we call SYZ mirror transformations, can be applied to provide a geometric understanding of the mirror symmetry phenomena for semi-flat
Calabi-Yau manifolds and toric Fano manifolds. We also speculate the possible applications of these transformations to other more general settings.

Contents
1. Introduction
2. SYZ mirror transformations without corrections
2.1. Semi-flat SYZ mirror transformations
2.2. Transformations of branes
3. SYZ mirror transformations with corrections
3.1. Mirror symmetry for toric Fano manifolds
3.2. SYZ transformations for toric Fano manifolds
3.3. Transformation of branes
4. Further questions
4.1. Toric Fano manifolds
4.2. Toric non-Fano or non-toric Fano manifolds
4.3. Calabi-Yau manifolds
References

1
4
4
8
9
10
12
18
19
19
20
21
21

1. Introduction
In 1996, Strominger, Yau and Zaslow suggested, in their ground-breaking work
[40], a geometric approach to the mirror symmetry for Calabi-Yau manifolds.
Roughly speaking, the Strominger-Yau-Zaslow (SYZ) Conjecture asserts that any
Calabi-Yau manifold X should admit a fibration by special Lagrangian tori and the mirror of X, which is another Calabi-Yau manifold Y, can be obtained by Tduality, i.e. dualizing the special Lagrangian torus fibration of X. Moreover, the symplectic geometry (A-model) of X should be interchanged with the complex geometry (B-model) of Y, and vice versa, through fiberwise Fourier-Mukai-type transformations, suitably modified by quantum corrections. These transformations are called SYZ mirror transformations and they will be the theme in this article.
Much work has been done on the SYZ Conjecture. Following the work of
Hitchin [24], Leung-Yau-Zaslow [32] and Leung [31] explained successfully and neatly the mirror symmetry for semi-flat Calabi-Yau manifolds by using semi-flat
1

2

K.-W. CHAN AND N.-C. LEUNG

SYZ mirror transformations. These are honest fiberwise real Fourier-Mukai transformations. The advantage in this case is the absence of quantum corrections by holomorphic curves and discs. This is due to the fact that the special Lagrangian torus fibrations on semi-flat Calabi-Yau manifolds do not admit singularities, and, accordingly, the bases are smooth affine manifolds.
To deal with general compact Calabi-Yau manifolds, however, one cannot avoid singularities in Lagrangian torus fibrations, and hence singularities in the base affine manifolds. Consequently, quantum corrections will come into play.
This necessitates the study of moduli spaces of special Lagrangian submanifolds and affine manifolds with singularities, which makes the subject much more sophisticated and difficult. Nevertheless, the recent progress made by Gross and
Siebert [21], after earlier works of Fukaya [13] and Kontsevich-Soibelman [30], was doubtlessly a significant step towards establishing the SYZ Conjecture for general compact Calabi-Yau manifolds.1
On the other hand, mirror symmetry phenomena have also been observed for
Fano manifolds (and other classes of manifolds or orbifolds as well). The mirror of a Fano manifold X̄ is predicted by Physicists to be given by a Landau-Ginzburg model, which is a pair (Y, W ), consisting of a non-compact Kähler manifold Y and a holomorphic function W : Y → C called the superpotential. A very important class of examples is provided by toric Fano manifolds. In this case, the mirror manifold Y is biholomorphic to (a bounded domain of) (C ∗ )n and the superpotential W is a Laurent polynomial which can be written down explicitly. Ample evidences have been found in this toric Fano case; in particular, Cho and Oh
[9] proved that the superpotential can be computed in terms of the counting of
Maslov index two holomorphic discs in X̄ with boundary in Lagrangian torus fibers. In [4], Auroux applied the SYZ philosophy to the study of the mirror symmetry for a compact Kähler manifold equipped with an anticanonical divisor. This is a generalization of the mirror symmetry for Fano manifolds, and, again, the mirror is given by a Landau-Ginzburg model. Auroux also made an attempt to compute the superpotential in terms of the counting of holomorphic discs, and analyzed the resulting wall-crossing phenomena. In [7], we studied the mirror symmetry for toric Fano manifolds, again through the SYZ approach, and we constructed and applied SYZ mirror transformations for toric Fano manifolds to explain various geometric results implied by mirror symmetry.
A brief explanation of the results in [7] is now in order; for more details, see
Section 3. Let X̄ be a toric Fano manifold, i.e. a smooth projective toric variety such that the anticanonical line bundle K X̄ is ample. Let ω X̄ be a toric Kähler structure on X̄. The moment map µ X̄ : X̄ → P̄ of the Hamiltonian T n -action on ( X̄, ω X̄ ) is a natural Lagrangian torus fibration. Here P̄ ⊂ R n is a polytope defining ( X̄, ω X̄ ). The restriction of the moment map to the open dense T n -orbit
X ∼
= (C ∗ )n ⊂ X̄ is a Lagrangian torus bundle µX = µ X̄ | X : X → P, where P
denotes the interior of the polytope P̄. Our first result in [7] showed that the mirror manifold Y is nothing but the SYZ mirror manifold of X, i.e. the total space

1We should mention that the Gross-Siebert program is expected to work for non-Calabi-Yau manifolds (e.g. Fano manifolds) as well.

SYZ TRANSFORMATIONS

3

2
of the torus bundle dual to µ X : X → P (see Proposition 3.1).√
Furthermore, sf the semi-flat SYZ transformation F takes the exponential of ( −1 times) the symplectic structure ω X = ω X̄ | X on X to the holomorphic volume form ΩY on
Y.3 Note that ΩY determines a complex structure on Y by declaring that a 1-form
α is a (1, 0)-form if and only if αyΩY = 0. This part of the mirror symmetry does not involve quantum corrections.
To get the superpotential W, however, we need to take into account the quantum corrections due to the anticanonical toric divisor D∞ = X̄ \ X, which we have ignored above. Before doing that, we first take a digression to a well-known construction. For a simply connected symplectic manifold ( M, ω ), let LM be the free loop space, i.e. the space of smooth maps γ : S1 → M. The symplectic structure on M induces a symplectic structure on LM which will also be denoted by ω.
The action functional defined by

1
H (γ) :=
2π

Z

Dγ

ω,

where Dγ is a disk contracting γ, becomes a well-defined function on the univerg of the free loop space LM. The group of deck transformations sal covering LM
is H2 ( M, Z ). It is not hard to see that H is the moment map for the built-in g and the gradient flow lines of H are (pseudo-)holomorphic
S1 -action on LM, cylinders if we fix a compatible (almost) complex structure on M. Tentatively, the quantum cohomology (or Floer cohomology) is the S1 -equivariant Morse-Witten g However, the fact that LM
g is infinite cohomology of the moment map H on LM.
dimensional poses severe difficulties in implementing this idea.
One of our discoveries in [7] was that a finite dimensional subspace of LM is enough to capture the quantum corrections and recover the quantum cohomology, in the case when M = X̄ is a toric Fano manifold. Consider the subspace LX
of LX̄ consisting of those loops which are geodesic in the Lagrangian torus fibers
(with respect to the flat metrics) of the moment map µ X̄ : X̄ → P̄. We consider the function Ψ on LX defined by Ψ(γ) = exp(− H (γ)) if γ bounds a Maslov index two holomorphic disc and Ψ(γ) = 0 otherwise. The function Ψ : LX → C, as an object in the A-model of X̄, turns out to be the mirror of the superpotential W. In
[7], we constructed the SYZ mirror transformation F for the toric Fano manifold
X̄, and showed that the SYZ mirror transformation of Ψ is precisely the B-model superpotential W. Moreover, by incorporating the symplectic structure ω X and the holomorphic volume form ΩY , we proved that
√

F (e −1ω X +Ψ ) = eW ΩY ,
√

F −1 (eW ΩY ) = e −1ω X +Ψ ,

where F −1 is the inverse SYZ mirror transformation (see Theorem 3.1). Hence, the corrected symplectic structure on X and the complex structure on (Y, W ) are interchanged by the SYZ mirror transformation. On the other hand, we identified the small quantum cohomology ring QH ∗ ( X̄ ) of X̄ with an algebra of functions
2More precisely, the SYZ mirror manifold is a bounded domain in the mirror manifold Y predicted by Physicists.
3Throughout this paper, we assume that the B-field is zero.

4

K.-W. CHAN AND N.-C. LEUNG

on LX, and realized the quantum product as a convolution product (see Proposition 3.2). Then, we showed that the SYZ mirror transformation F exhibits a natural isomorphism between QH ∗ ( X̄ ) and the Jacobian ring Jac(W ) of the superpotential W, which takes the quantum product (now as a convolution product) to the ordinary product of Laurent polynomials, just as what classical Fourier series do (see
Theorem 3.2). We conclude that the mirror symmetry for toric Fano manifolds is nothing but a Fourier transformation!
The main goal of this article is to popularize the use of SYZ mirror transformations in exploring mirror symmetry phenomena. In Section 2, we review the use of semi-flat SYZ mirror transformations in the study of the mirror symmetry for semi-flat Calabi-Yau manifolds, where quantum corrections are absent. This is the toy case which lays the basis for subsequent development in the investigation of the SYZ Conjecture. Section 3 discusses the mirror symmetry for toric
Fano manifolds, where quantum corrections arise due to the anticanonical toric divisor. Following [7], we demonstrate how to construct and apply SYZ mirror transformations in this case. The final section contains a brief discussion of possible generalizations.
Acknowledgments. The authors are grateful to the organizers of the conference
"New developments in Algebraic Geometry, Integrable Systems and Mirror Symmetry" held in Kyoto University in January 2008 for giving them an opportunity to participate in such a stimulating and fruitful event. Thanks are also due to
Hiroshi Iritani and Cheol-Hyun Cho for many useful discussions. Finally, we thank the referee for several helpful comments. K.-W. C. was partially supported by Harvard University and the Croucher Foundation Fellowship. N.-C. L. was partially supported by RGC grants from the Hong Kong Government.
2. SYZ mirror transformations without corrections
In this section, we review the construction of SYZ mirror transformations for semi-flat Calabi-Yau manifolds and see how they were applied in the study of semi-flat mirror symmetry.
2.1. Semi-flat SYZ mirror transformations. Denote by N ∼
= Z n a rank-n lattice and M = Hom( N, Z ) the dual lattice. Let D ⊂ M√R = M ⊗Z R be a convex domain. 4 Then the tangent bundle TD =√D × −1MR is naturally a complex manifold with complex coordinates x j + −1y j , j = 1, . . . , n, where x1 , . . . , x n ∈ R and y1 , . . . , y n ∈ R are respectively the base coordinates on D
and fiber coordinates on MR . We have
√
√ the standard holomorphic volume form
Ω TD = d( x1 + −1y1 ) ∧ . . . ∧ d( xn + −1yn ) on TD. By taking fiberwise quotient by the lattice M ⊂ MR , we can compactify the fiber directions to give the complex manifold
√
Y = TD/M = D × −1TM , where TM denotes the torus
√ MR /M. The complex coordinates on Y are naturally given by z j = exp( x j + −1y j ), j = 1, . . . , n, where y1 , . . . , y n ∈ R/2πZ are now coordinates on TM . Note that Y is biholomorphic to an open part of (C ∗ )n =
4More generally, instead of a convex domain, one may consider a smooth affine manifold.

SYZ TRANSFORMATIONS

5

TMR /M. The projection to D is a torus bundle, which we denote by νY : Y → D.
The holomorphic n-form Ω TD descends to give the holomorphic volume form
ΩY =

dzn dz1
∧...∧
z1
zn

on Y. As mentioned in the introduction, ΩY in turn determines the complex structure on Y: a 1-form α is of (1, 0)-type if and only if αyΩY = 0. Further, if φ
is an elliptic solution of the real Monge-Ampère equation det then the Kähler form
ωY : =

√

 ∂2 φ 
= const,
∂x j ∂xk

¯ = ∑ φjk dx j ∧ dyk ,
−1∂∂φ
j,k

∂2 φ

with φjk denoting ∂x ∂x , gives a Calabi-Yau metric on Y, and j

k

νY : Y → D
becomes a special Lagrangian torus bundle (SYZ fibration). In summary, we have the following structures on the complex n-dimensional semi-flat Calabi-Yau manifold Y:
Riemannian metric gY = ∑ j,k φjk (dx j ⊗ dxk + dy j ⊗ dyk )
√
V
Holomorphic volume form ΩY = nj=1 (dx j + −1dy j )
Symplectic form
ωY = ∑ j,k φjk dx j ∧ dyk
SYZ fibration
νY : Y → D
As suggested in the monumental work Strominger-Yau-Zaslow [40], the mirror of Y, which is another Calabi-Yau manifold we denote by X, should be given by the moduli space of pairs ( L, ∇), where L is a special Lagrangian torus fiber in
Y, and ∇ is a flat U (1)-connection on the trivial complex line bundle L√× C → L.
This is nothing but the total space of the torus fibration µ X : X = D × −1TN →
D, where TN = NR /N = ( TM )∨ and NR = N ⊗Z R, which is dual to νY : Y → D.
This is called T-duality in physics. Furthermore, X can naturally be viewed as
√
the fiberwise quotient of the cotangent bundle T ∗ D = D × −1NR by the lattice
N ⊂ NR . In particular, the standard symplectic form ω T ∗ D = ∑nj=1 dx j ∧ du j descends to give a symplectic form n

ω X = ∑ dx j ∧ du j j =1

on X = T ∗ D/N, where u1 , . . . , un ∈ R/2πZ are coordinates on TN . Through the metric

gX = ∑(φjk dx j ⊗ dxk + φ jk du j ⊗ duk ), j,k

where (φ jk ) is the inverse matrix of (φ

jk ), we obtain a complex structure on X
√
with complex coordinates given by d log(w j ) = ∑nk=1 φjk dxk + −1du j . There is

6

K.-W. CHAN AND N.-C. LEUNG

a corresponding holomorphic volume form which can be written as
ΩX =

n n
√
^
dwn dw1
( ∑ φjk dxk + −1du j ).
∧...∧
=
w1
wn j =1 k =1

The projection map
µX : X → D
now naturally becomes a special Lagrangian torus fibration. In summary, we have the following structures on X:
Riemannian metric
Holomorphic volume form
Symplectic form
SYZ fibration

gX = ∑ j,k (φjk dx j ⊗ dxk + φ jk du j ⊗ duk )
√
V
Ω X = nj=1 (∑nk=1 φjk dxk + −1du j )
ω X = ∑nj=1 dx j ∧ du j
µX : X → D

We remark that both Y and X admit natural Hamiltonian T n -actions, but while
µ : X → D is a moment map for the TN -action on X, ν : Y → D is not a moment map for the TM -action on Y. In fact, a moment map µY : Y → NR for the TM action on Y is given by
µY = Lφ ◦ νY , where Lφ : D → NR is the Legendre transform of φ defined by
 ∂φ
∂φ 
Lφ ( x1 , . . . , x n ) = dφx =
.
,...,
∂x1
∂xn

Since φ is convex, the image D ∗ = Lφ ( D ) is an open convex subset of ( MR )∗ =
NR . (For this and other properties of the Legendre transform, see the book of
Guillemin [22], Appendix 1.) In the action coordinates x1 , . . . , x n of D ∗ , which are
∂x j
= φjk , the various structures on Y can be rewritten as: given by ∂x k

Riemannian metric
Holomorphic volume form
Symplectic form
SYZ fibration

gY = ∑ j,k (φ jk dx j ⊗ dx k + φjk dy j ⊗ dyk )
√
V
ΩY = nj=1 (∑nk=1 φ jk dx k + −1dy j )
ωY = ∑nj=1 dx j ∧ dy j
µY : Y → D ∗

We call X the SYZ mirror manifold of Y (and vice versa) since the symplectic
(resp. complex) geometry of X and the complex (resp. symplectic) geometry of Y
are interchanged under the semi-flat SYZ mirror transformation, which is described as follows.
First recall that the dual torus TM = ( TN )∨ can be interpreted as the moduli space of flat U (1)-connections on the trivial complex line bundle over TN . More precisely, given y = (y1 , . . . , y n ) ∈ MR ∼
= R n , we have a flat U (1)-connection
√
−1 n
∇y = d +
y j du j
2 j∑
=1
on TN × C → C. The holonomy of ∇y is given by the map
√

hol∇y : N → U (1), v 7→ e− −1h y,vi .

Hence, ∇y is gauge equivalent to the trivial connection if and only if y ∈ M ∼
=
(2πZ )n . Moreover this construction gives all flat U (1)-connections on the trivial

SYZ TRANSFORMATIONS

7

complex line bundle over TN up to unitary gauge transformations. The universal U (1)-bundle, i.e. the Poincaré line bundle P , is given by the trivial complex√line bundle ( TN × TM ) × C → TN × TM equipped with the connection
−1 n
2 ∑ j =1 ( y j du j − u j dy j ). The curvature of this connection is the two form

d+

F=

√

n

−1 ∑ dy j ∧ du j .
j =1

√
Now consider the relative version of this picture. Let X × D Y = D × −1( TN ×
TM ) be the fiber product of the dual torus bundles
µ : X → D and ν : Y → D. By
√
abuse of notations, we still use P and F = −1 ∑nj=1 dy j ∧ du j ∈ Ω2 ( X × D Y ) to denote the fiberwise universal line bundle and curvature two form respectively.
Definition 2.1. The semi-flat SYZ mirror transformation

F sf : Ω∗ ( X ) → Ω∗ (Y )
is defined by

F sf (α) =
=

√
1
∗
√
(α) ∧ e −1F )
πY,∗ (π X
n
(2π −1)
Z
√
1
∗
√
(
α) ∧ e −1F ,
πX
(2π −1)n TN

where π X : X × D Y → X and πY : X × D Y → Y are the two projections.
What is crucial is that this Fourier-Mukai-type transformation transforms the symplectic structure on X to the complex structure on Y in the sense of the following two propositions. These already appeared in [7], Proposition 3.2. We include their proofs, which are somewhat interesting, here for completeness.
Proposition 2.1.

√

F sf (e −1ω X ) = ΩY .
Proof.
√

F sf (e −1ω X ) =
=
=

=

where we have

R

Z

√
√
1
∗
√
πX
(e −1ω X ) ∧ e −1F
n
(2π −1) TN
Z
√
√
1
−1 ∑nj=1 ( dx j + −1dy j )∧ du j
√
e
(2π −1)n TN
Z
n
√
^
√

1
√
1 + −1(dx j + −1dy j ) ∧ du j n
(2π −1) TN j=1
!
Z
n
√
^
1
(dx + −1dy j ) ∧ du1 ∧ . . . ∧ dun
(2π )n TN j=1 j

= ΩY ,
TN du1 ∧ . . . ∧ dun = (2π )

n in the final step.



As a mirror transformation, F sf should have the inversion property. This is the following proposition.

8

K.-W. CHAN AND N.-C. LEUNG

Proposition 2.2. If we define the inverse transform (F sf )−1 : Ω∗ (Y ) → Ω∗ ( X ) by
√
1
√
π X,∗ (πY∗ (α) ∧ e− −1F )
(2π −1)n
Z
√
1
√
πY∗ (α) ∧ e− −1F ,
(2π −1)n TM

(F sf )−1 (α) =
=
then we have

√

(F sf )−1 (ΩY ) = e −1ω X .
Proof.

(F sf )−1 (ΩY ) =
=
=
=
=
=
=

Z

1
√
(2π −1)n

Z

1
√
(2π −1)n

Z

1
√
(2π −1)n
1
√
(2π −1)n
1
(2π )n

Z

1
(2π )n

Z

1
(2π )n
√

Z

TM

TM

TM

Z

TM

πY∗ (ΩY ) ∧ e−
n
^

TM

TM

TM

(dx j +

√

√

j =1
n
^

(dx j +

n
^

dx j +

√

j =1

j =1

n
^

(1 +

n
^

e

j =1

√

√

√

−1F

!

−1dy j ) ∧ edy j ∧du j

j =1
√
−1 ∑nj=1 dx j ∧ du j

e



−1dy j + dx j ∧ dy j ∧ du j

−1dx j ∧ du j ) ∧ dy j

−1dx j ∧ du j

n

−1dy j ) ∧ e∑ j=1 dy j ∧du j

∧ dy j







∧ dy1 ∧ . . . ∧ dyn

= e −1ω X .

By exactly the same arguments, one can also show that
√

√

F sf (Ω X ) = e −1ωY , (F sf )−1 (e −1ωY ) = Ω X .
If we take into account the B-fields, then the semi-flat SYZ transformation will give an identification between the moduli space of complexified Kähler structures on
X with the moduli space of complex structures on Y, and vice versa. For this and transformations of other geometric structures, we refer the reader to Leung [31].
2.2. Transformations of branes. Lying at the heart of the SYZ
√ Conjecture is the basic but important observation that a point z = exp( x + −1y) ∈ Y defines a flat U (1)-connection ∇y on the trivial complex line bundle over the special La1
grangian torus fiber L x = µ−
X ( x ). Now, the point z ∈ Y together with its structure sheaf Oz can be considered as a B-brane on Y; while the pair ( L x , L y ), where L y denotes the flat U (1)-bundle ( L x × C, ∇y ), gives an A-brane on X. This implements the simplest case of correspondence between branes on mirror manifolds

SYZ TRANSFORMATIONS

9

via SYZ transformations:

( L x , L y ) ←→ (z, Oz ).
The space of infinitestimal deformations of the A-brane ( L x , L y ), which is given
√
by H 1 ( L x , R ) × H 1 ( L x , −1R ) = H 1 ( L x , C ), is canonically identified with the tangent space Tz Y, the space of infinitestimal deformations of the sheaf Oz .
On the other hand, consider a section L = {( x, u( x )) ∈ X : x ∈ D } of µ X :
X → D. The submanifold L is Lagrangian if and only if (locally) there exists
∂f a function f such that u j = ∂x . By the above observation (now used in the j

opposite way), a point ( x, u( x )) ∈ L determines a flat U (1)-connection ∇u( x ) on the trivial complex line bundle over the fiber ( L x )∨ = νY−1 ( x ). The family of points {( x, u( x )) : x ∈ D } thus patch together to give the U (1)-connection
√
−1 n u j ( x )dy j
∇ L = dY −
2 j∑
=1
on a certain complex line bundle over Y; its curvature two form is given by
√
 √ −1 n

∂u j
−1
FL = dY −
dx ∧ dy j , u j ( x )dy j = −
∑
∑
2 j =1
2 j,k ∂xk k and, in particular,
FL2,0 =

 ∂u
∂u  dz j dzk
1
j
− k
∧
.
∑
8 j<k ∂xk
∂x j z j zk

We conclude that ∇ L is integrable, i.e. FL2,0 = 0, if and only if L is Lagrangian.
More generally, we can equip L with a flat U (1)-bundle L = ( L × C, d L + α), where α ∈ Ω1 ( L, R ) is a closed (and hence exact) one-form. The A-brane ( L, L )
is then transformed to the U (1)-connection

∇ L,L = ∇ L + α,

which again is integrable if and only if L is Lagrangian. Furthermore, one can prove that ∇ L,L satisfies the deformed Hermitian-Yang-Mills equations if and only if L is special Lagrangian (see Leung-Yau-Zaslow [32] and Leung [31] for the detailed proofs). ∇ L,L is a connection on the holomorphic line bundle over Y
given by the semi-flat SYZ transformation of L:
∗
L L,L = πY,∗ (π X
( ι ∗ L ) ⊗ P ),

where ι : L ֒→ X is the inclusion map. In conclusion, the A-brane ( L, L ) is transformed to the B-brane (Y, L L,L ) through semi-flat SYZ transformations:

( L, L ) ←→ (Y, L L,L ).

3. SYZ mirror transformations with corrections
In the previous section, we see that T-duality and SYZ mirror transformations can be applied successfully to give a geometric understanding of the mirror symmetry for semi-flat Calabi-Yau manifolds. However, no quantum corrections were involved in this case due to the absence of holomorphic curves and discs. The existence of quantum corrections is also closely related to the singularities of the
Lagrangian torus fibrations, which again are not present in the semi-flat case. In

10

K.-W. CHAN AND N.-C. LEUNG

this section, following [7], we are going to discuss how SYZ mirror transformations can be applied to a case where quantum corrections do exist, namely, the mirror symmetry for toric Fano manifolds.
3.1. Mirror symmetry for toric Fano manifolds. We begin with a more detailed description of the mirror picture for toric Fano manifolds [17], [29], [27]. Let
P̄ ⊂ MR be a smooth reflexive polytope given by the inequalities

h x, vi i ≥ λi ,

i = 1, . . . , d,

where v1 , . . . , v d ∈ N are primitive vectors and h·, ·i : MR × NR → R is the dual pairing. This determines a toric Fano manifold X̄, together with a Kähler structure ω X̄ . Unlike the case of Calabi-Yau manifolds, the mirror of X̄ is not another compact Kähler manifold, but a Landau-Ginzburg model: a pair (Y, W )
consisting of a noncompact Kähler manifold Y, which (as a complex manifold) is biholomorphic to (a bounded domain of) (C ∗ )n , and the Laurent polynomial
W = eλ1 zv1 + . . . + eλd zvd : Y → C,

v

v

which is called the superpotential. Here zvi denotes the monomial z1 i1 . . . znin in the coordinates z1 , . . . , zn of Y. For example, if P = {( x1 , x2 , x3 ) ∈ R3 : x1 ≥
0, x2 ≥ 0, x1 + x2 ≤ t}, then X̄ = CP2 and the mirror Landau-Ginzburg model is
−t given by the Laurent polynomial W (z1 , z2 ) = z1 + z2 + ze z2 on Y = (C ∗ )2 .
1
Among the many mirror symmetry predictions are the following conjectures:
Conjecture 3.1.
1. The small quantum cohomology ring QH ∗ ( X̄ ) of X̄ is isomorphic to the Jacobian ring Jac(W ) of W, where
1
Jac(W ) = C [z1±1 , . . . , z±
n ] /h ∂1 W, . . . , ∂n W i,

and ∂ j denotes z j ∂z∂ .
j

2. (Homological mirror symmetry, see [29], [39], [37]) There are equivalences of triangulated categories

∼
=
π
∼
D Fuk( X̄ ) =
D b Coh( X̄ )

D π Fuk(Y, W )
DSing (Y, W )

where D π Fuk(Y, W ) is (a suitably defined version of) the derived Fukaya category of the Landau-Ginzurg model (Y, W ) and DSing (Y, W ) is the category of singularities of (Y, W ).
Substantial evidences [19], [25], [39], [41], [5], [6], [1], [2], [9], [8] have been found for these conjectures, while evidence in the Calabi-Yau and other non-toric cases is much rarer. This is partly due to the fact that geometric structures on toric varieties are highly computable and explicit, making them an exceptionally fertile testing ground for techniques and conjectures.
One of these explicit structures: the Lagrangian torus fibration on X̄ given by the moment map µ X̄ : X̄ → P̄ of the Hamiltonian TN -action on ( X̄, ω X̄ ), is particularly important in the SYZ spproach and in the constructions of SYZ
mirror transformations. Let
µX : X → P

SYZ TRANSFORMATIONS

11

be the restriction of the moment map to the open dense TN -orbit X = X̄ \ D∞ ,
S
where D∞ = di=1 Di is the anticanonical toric divisor, and P is the interior of P̄.
In the symplectic (or action-angle) coordinates,
√
X = T ∗ P/N = P × −1TN
and the restriction of ω X̄ to X is nothing but the standard symplectic structure n

ω X = ∑ dx j ∧ du j , j =1

where x1 , . . . , x n ∈ R and u1 , . . . , un ∈ R/2πZ are respectively the base coordinates on P and fiber coordinates on TN (see Abreu [3]). Now we are in exactly the same situation as in the previous section and it is tempting to assert that the mirror manifold Y predicted
√ by Physicists is given by the SYZ mirror manifold of
X, which is TP/M = P × −1TM . This is indeed nearly the case.
∗ n
Proposition 3.1 (Proposition 3.1 in [7]). The mirror manifold
√ Y = (C ) predicted by
Physicists contains the SYZ mirror manifold TP/M = P × −1TM of X = X̄ \ D∞ as a bounded domain

{(z1 , . . . , zn ) ∈ Y : |eλi zvi | < 1 for i = 1, . . . , d}.

Equivalently, the SYZ mirror manifold is given by the preimage of P ⊂ MR = R n under the Log map
Log : (C ∗ )n → R n , (z1 , . . . , zn ) 7→ (log |z1 |, . . . , log |zn |).
The same result also appeared in Auroux’s paper [4] (Proposition 4.2). Also included in his paper was a discussion of the issue that the SYZ mirror manifold
(a bounded domain in (C ∗ )n ) is "smaller" than Hori-Vafa’s mirror manifold (the whole (C ∗ )n ). There is evidence (say, in Abouzaid’s works [1], [2]) showing that one should work with the SYZ mirror manifold, instead of the whole (C ∗ )n , in studying mirror symmetry. In any case, we will use and work with the SYZ mirror manifold, i.e. the bounded domain in (C ∗ )n , and
√ denote it by Y henceforth.
In terms of the coordinates z
=
exp
(−
x
−
−1y1 ), . . . , zn = exp(− xn −
1
1
√
−1yn ) ∈ C ∗ of Y ⊂ (C ∗ )n , the holomorphic volume form is given by the standard one on (C ∗ )n : dzn dz
ΩY = 1 ∧ . . . ∧
z1
zn and the torus fibration νY : Y → P is the restriction of the Log map. We remark that metrically we are not considering X = X̄ \ D∞ as a Calabi-Yau manifold; instead of the semi-flat Calabi-Yau metric, we use the TN -invariant Kähler metric on X̄ (and the corresponding dual metric on Y). These are defined (cf. Guillemin
[22] and Abreu [3]) using the strictly convex function φP : P → R given by
φP ( x ) =

1 d li ( x ) log li ( x ),
2 i∑
=1

where li ( x ) = h x, vi i − λi for i = 1, . . . , d, instead of a solution of the real MongeAmpère equation. For example, this gives the standard Fubini-Study metric on
X̄ = CPn . Using these metrics and the corresponding holomorphic volume forms,
X and Y are almost Calabi-Yau manifolds and the torus fibers of µ X and νY are special Lagrangian submanifolds (also see Section 2 in Auroux [4]).

12

K.-W. CHAN AND N.-C. LEUNG

3.2. SYZ transformations for toric Fano manifolds. By applying the semi-flat
SYZ mirror transformation or T-duality, we can obtain the mirror manifold Y.
But where comes the superpotential W : Y → C? Recall that, in applying Tduality, we have completely ignored the compactification of X, which is given
S
by adding the anticanonical toric divisor D∞ = di=1 Di . As suggested in the foundational work of Fukaya-Oh-Ohta-Ono [14], this has tremendous effect on the Floer theory of the Lagrangian torus fibers of µ X : X → P, and this is indeed where quantum corrections by holomorphic discs come into play.
As have been discussed in the introduction, motivated by the idea of using
Morse theory on the free loop space LX̄ to construct the quantum cohomology
QH ∗ ( X̄ ), we introduce the subspace LX ⊂ LX̄ consisting of those loops which are geodesic in the Lagrangian torus fibers of the moment map µ X : X → P, i.e.
1
LX = {γ ∈ LX̄ : γ is a geodesic in L x = µ−
X ( x ) for some x ∈ P}.

Concretely, we have
LX = X × N = P ×

√

−1TN × N,

and we consider it as a (trivial) Z n -cover of X, π : LX → X. Notice that, for each
Lagrangian torus fiber L x , x ∈ P, we have a canonical identification π1 ( L x ) ∼
= N.
We are going to define a function Ψ on LX in terms of the counting of holomorphic discs in X̄ of minimal Maslov index. This will recapture the information of the compactification of X by D∞ , which we have ignored previously, and Ψ
serves as the object in the A-model of X̄ mirror to the superpotential W. To do this, let’s first recall the fundamental results of Cho-Oh [9] on the classification of holomorphic discs in X̄ with boundary in Lagrangian torus fibers of µ X : X → P.
1
Let L x = µ−
X ( x ) be the Lagrangian torus fiber in X over a point x ∈ P. Then the relative homotopy group π2 ( X̄, L x ) is generated by the Maslox index two classes β 1 , . . . , β d , which are represented by holomorphic discs in ( X̄, L x ). Note that we have, ∂β i = vi , for i = 1, . . . , d, where ∂ : π2 ( X̄, L x ) → π1 ( L x ) ∼
= N is the natural boundary map. In [9], Cho and Oh proved that, for i = 1, . . . , d and for each point p ∈ L x , there is a unique (up to automorphism of the domain)
Maslov index two J-holomorphic disc ϕi : ( D2, ∂D2 ) → ( X̄, L x ) in the class β i which passes through p and intersects the toric divisor Di at an interior point.5
Here J is the complex structure on X̄ determined by the fan Σ dual to P̄.
Definition 3.1. For i = 1, . . . , d, define Ψi : LX → R by
(
R
1
ni ( p) exp(− 2π
β i ω X̄ ) if v = v i
Ψi ( p, v) =
0
if v 6= v i , for ( p, v) ∈ LX = X × N, where ni ( p) is the algebraic number of Maslov index two
J-holomorphic discs in ( X̄, Lµ X ( p) ) in the class β i which pass through p. Then set
Ψ = Ψ1 + . . . + Ψd : LX → R.
5Another way to state this result is the following. Let M ( β ) be the moduli space of J-holomorphic
1
i discs ϕ : ( D 2 , ∂D 2 ) → ( X̄, L x ) in the class β i with 1 boundary marked point. Let ev : M1 ( β i ) → L x be the evaluation map at the boundary marked point. Then the result of Cho and Oh says that ev∗ [M1 ( β i )] = [ L x ] as n-cycles in L x . See also Sections 3.1 and 4 in Auroux [4].

SYZ TRANSFORMATIONS

13

By their definitons, the TN -invariant functions Ψ1 , . . . , Ψd carry enumerative meaning, although by Cho and Oh’s result, we always have ni ( p) = 1, for all i and any p. One may think of the TN -invariant function Ψ as recording which cycle v ∈ N = π1 ( L x ) collapses to a point as one goes towards the anticanonical toric divisor D∞ , or equivalently, which geodesic loop γ ∈ LX bounds a holomorphic disc of Maslov index two.
Remark 3.1. Before showing how to transform Ψ to get the superpotential W, we remark that the TN -invariant function Φ : LX → R introduced in [7], Definition 2.1, is nothing but the "exponential" of Ψ, i.e.
Φ = Exp Ψ,
1
where Exp Ψ is defined as ∑∞
Ψ ⋆ .{z
. . ⋆ Ψ} in which ⋆ denotes the convolution product k =0 k! |
k times

of a certain class of functions on LX with respect to the lattice N. Now each point q = (q1 , . . . , ql ) (l = d − n) in the Kähler cone K( X̄ ) ⊂ H 2 ( X̄, R ) determines a symplectic structure ω X̄ on X̄ and we can choose the polytope P̄ = { x ∈ MR : h x, vi i ≥
λi , i = 1, . . . , d} such that v1 = e1 , . . . , v n = en is the standard basis of N = Z n
, λ1 = . . . = λn = 0 and λn+ a = log q a for a = 1, . . . , l. We thus get two families of functions {Ψq }q∈K and {Φq }q∈K . By the symplectic area formula of Cho-Oh ([9],
Theorem 8.1), we have
Z

D2

ϕ∗i ω X̄ =

Z

βi

ω X̄ = 2π (h x, vi i − λi ),

for i = 1, . . . , d. Hence, for any ( p, v) ∈ LX,
 −h x,v i i
e
Ψi ( p, v) =
0
for i = 1, . . . , n, and

Ψn+ a ( p, v) =



q a e−h x,vn+ ai
0

for a = 1, . . . , l, where x = µ X ( p). It follows that qa

if v = vi if v 6= vi , if v = v n+1
if v 6= v n+ a ,

∂Φq
= Φq ⋆ Ψ n+ a
∂q a

for a = 1, . . . , l, which is the first part of Proposition 1.1 in [7].
On the other hand, the functions Ψ1 , . . . , Ψd are intimately related to the small quantum cohomology QH ∗ ( X̄ ) of X̄, as was shown in the following
Proposition 3.2 (Second part of Proposition 1.1 in [7]). Assume that X̄ is a product of projective spaces. Then we have a natural isomorphism of C-algebras
1
QH ∗ ( X̄ ) ∼
= C [Ψ1±1 , . . . , Ψ±
n ] /L

±1
1
±1
where C [Ψ1±1 , . . . , Ψ±
n ] is the polynomial algebra generated by Ψ1 , . . . , Ψ n with respect to the convolution product ⋆, and L is the ideal generated by linear relations:
∑di=1 ai Ψi ∼ ∑di=1 bi Ψi if and only if the corresponding divisors ∑di=1 ai Di and ∑di=1 bi Di are linearly equivalent.

Remark 3.2. By employing Givental’s mirror theorem [19], one can in fact show that the proposition holds for all toric Fano manifolds. See Remark 2.3 in [7] for details.

14

K.-W. CHAN AND N.-C. LEUNG

We need the assumption that X̄ is a product of projective spaces as we are intended for a geometric understanding of the isomorphism in Proposition 3.2
by using tropical geometry. This is briefly described as follows (see Subsection
∗ ( X̄ ) of the small
2.2 in [7] for details). One first defines a tropical version QHtrop quantum cohomology ring of X̄. Since X̄ is a product of projective spaces, we have a one-to-one correspondences between the J-holomorphic curves in X̄ which have contribution to the quantum product in QH ∗ ( X̄ ) and those tropical curves
∗ ( X̄ ), by in NR which have contribution to the tropical quantum product in QHtrop the correspondence theorem of Mikhalkin [33] and Nishinou-Siebert [36]. From this follows the canonical isomorphism
QH ∗ ( X̄ ) ∼
= QH ∗ ( X̄ ).
trop

Then comes a simple but important observation: Each tropical curve which has
∗ ( X̄ ) is obtained by gluing tropical contribution to the tropical quantum product in QHtrop discs in NR .6 On the other hand, these tropical discs are exactly corresponding to the families of Maslov index two J-holomorphic discs in X̄ with boundary in Lagrangian torus fibers, which were used to define the functions Ψ1 , . . . , Ψd .
Hence, we naturally have another canonical isomorphism
1
∗
QHtrop
( X̄ ) ∼
= C [Ψ1±1 , . . . , Ψ±
n ] /L.

For example, let us take a look at the case of X̄ = CP2 . See Figure 3.1 below.
✻
t.................

D1

....
....
....
....
....
....
....
....
....
....
....
....
.... 3
....
....
....
....
....
....
R
....
....
....
....
....
....
....
....
..

D

P̄ ⊂ M

...
...
...
...
... 2
...
..
...
...
...
...
...
...
...
1
...............................................................................
.
.
...
.
.
...
.
.
...
...
....
....
....
.
.
.
.
....
3
...
...
R
....
....

v

ξ •

v

v
Γ⊂N
✲
0
t
D2
Figure 3.1
Denote by {e1 , e2 } the standard basis of N = Z2 . We have v1 = (1, 0), v2 =
(0, 1), v3 = (−1, −1), and the polytope P̄ ⊂ MR ∼
= R2 is defined by the inequalities x1 ≥ 0, x2 ≥ 0, x1 + x2 ≤ t,

where t > 0. There are three toric divisors D1 , D2, D3 corresponding to three functions Ψ1 , Ψ2 , Ψ3 ∈ C ∞ ( LX ) defined by
 −x e 1 if v = (1, 0)
Ψ1 ( p, v) =
0
otherwise,
 −x
2
e if v = (0, 1)
Ψ2 ( p, v) =
0
otherwise,

−(
t
−
x
1 − x2 )
e if v = (−1, −1)
Ψ3 ( p, v) =
0
otherwise,
6This idea was recently generalized by Gross [20] to understand tropically the big quantum cohomology and mirror symmetry of CP2 .

SYZ TRANSFORMATIONS

15

for ( p, v) ∈ LX and ( x1 , x2 ) = µ X ( p) ∈ P, respectively. The small quantum cohomology ring is given by

QH ∗ (CP2 ) = C [ D1 , D2, D3 ] D1 − D3 , D2 − D3 , D1 ∗ D2 ∗ D3 − q

= C[ H ] H3 − q , where we have, by abuse of notations, also use Di ∈ H 2 (CP2, C ) to denote the cohomology class Poincaré dual to Di , H ∈ H 2 (CP2, C ) is the hyperplane class and q = e−t . Fix any point p ∈ CP2 \ D∞ , then the quantum corrections, which appear in the relation
D1 ∗ D2 ∗ D3 = H 3 = q,

is due to the unique holomorphic curve ϕ : (P1 ; x1 , x2 , x3 , x4 ) → CP2 of degree
1 (i.e. a line) with 4 marked points such that ϕ( x4 ) = p and ϕ( x i ) ∈ Di , for
1
i = 1, 2, 3. Let x = µ X ( p) ∈ P and L x = µ−
X ( x ) be the Lagrangian torus fiber containing p. Using tropical geometry, one sees that there is a tropical curve
Γ in NR with three unbounded edges in the directions v1 , v2, v3 and the vertex mapped to ξ = Log( p) ∈ NR , which is corresponding to this holomorphic curve
(see Figure 3.1 above). Here, we identify X with (C ∗ )2 , and Log : X = (C ∗ )2 →
NR = R2 is the Log map we defined in Proposition 3.1. It is obvious that Γ can
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
..
...............................................................................
.
.
..
.
.
.
..
.
.
.
....
...
....
....
...
.
.
.
.
....
...
....
....
....

ξ •

...
...
...
...
...
...
...
...
..
...
...
...
...
...
.

• .................................................................................
••
v

glued from

Γ

Figure 3.2

v2

.
....
....
....
...
.
.
...
...
....
....
...
.
.
.
....
...
....
....
3
...

1

v

be obtained by gluing the three half lines emanating from the point ξ ∈ NR in the directions v1 , v2 , v3 . See Figure 3.2. These half lines are the tropical discs which are corresponding to the three families of Maslov index two holomorphic discs ϕ1 , ϕ2 , ϕ3 respectively. We see that the above quantum relation corresponds exactly to the equation
Ψ1 ⋆ Ψ2 ⋆ Ψ3 = q in C [Ψ1±1 , Ψ2±1 ].
Without the assumption that X̄ is a product of projective spaces, the tropical interpretation will break down. This is because for general toric Fano manifolds, the holomorphic curves which contribute to the small quantum product may have components mapped into the anticanonical toric divisor D∞ . An example is provided by the exceptional curve in the blowup of CP2 at one TN -invariant point
(see Example 3 in Section 4 in [7]). Now the problem is that tropical geometry cannot be used to count these holomorphic curves. In other words, there are no tropical curves corresponding to such holomorphic curves (cf. Rau [38]).
Now it’s time to return to the main theme of this section, namely, we can construct and apply SYZ mirror transformations to the study of mirror symmetry for toric Fano manifolds. First we equip LX = X × N with the symplectic

16

K.-W. CHAN AND N.-C. LEUNG

structure π ∗ (ω X ), which we denote again by ω X . Also let µ LX : LX → P be the composition map µ X ◦ π. √Analog to the semi-flat case, we consider the fiber product LX × P Y = P × N × −1( TN × TM ) of the fibrations µ LX : LX → P and
νY : Y → P. Note that we have a covering √
map LX × P Y → X × P Y. Pulling back the universal curvature two-form F = −1 ∑nj=1 dy j ∧ du j ∈ Ω2 ( X × P Y ), we get a two-form on LX × P Y, which we again denote by F. We further define the holonomy function hol : LX × P Y → C by
√

hol( p, v, z) = hol∇y (v) = e− −1h y,vi
√
for ( p, v) ∈ LX, z = exp(− x − −1y) ∈ Y such that µ X ( p) = νY (z) = x. The SYZ
mirror transformation for toric Fano manifolds is constructed as a combination of the semi-flat SYZ transformation F sf and fiberwise Fourier series.

Definition 3.2. The SYZ mirror transformation F : Ω∗ ( LX ) → Ω∗ (Y ) for X̄ is defined by
√
√
F (α) = (−2π −1)−n πY,∗ (π ∗LX (α) ∧ e −1F hol)
Z
√
√
= (−2π −1)−n
π ∗LX (α) ∧ e −1F hol,
N × TN

where π LX : LX × P Y → LX and πY : LX × P Y → Y are the two natural projections.
The basic properties of F are similar to those of other Fourier-type transformations, and in particular, it satisfies the inversion property with the inverse SYZ
mirror transformation F −1 : Ω∗ (Y ) → Ω∗ ( LX ) defined by
√
√
F −1 (α) = (−2π −1)−n π LX,∗ (πY∗ (α) ∧ e− −1F hol−1 )
Z
√
√
= (−2π −1)−n
πY∗ (α) ∧ e− −1F hol−1 .
TM

In [7], the SYZ mirror transformation was, for the first time, used to study the appearance of the superpotential W as quantum corrections. More precisely, we showed that
Theorem 3.1 (First part of Theorem 1.1 in [7]). The SYZ mirror transformation (or fiberwise Fourier series) of the function Ψ, defined in terms of the counting of Maslov index two J-holomorphic discs in the toric Fano manifold X̄ with boundary in Lagrangian torus fibers, gives the superpotential W : Y → C on the mirror manifold:

F (Ψ) = W.

Furthermore, we can incorporate the symplectic structure ω X to give the holomorphic volume form of the Landau-Ginzburg model (Y, W ) in the sense that
√

Conversely, we have

F (e −1ω X +Ψ ) = eW ΩY .
√

F −1 (W ) = Ψ, F −1 (eW ΩY ) = e −1ω X +Ψ .
Remark 3.3.
1. We shall mention that the fact that the superpotential W can be computed in terms of the counting of Maslov index two holomorphic discs in X̄ with boundary in Lagrangian torus fibers was originally due to Cho and Oh [9]. The key

SYZ TRANSFORMATIONS

17

point of our result is that there is an explicit Fourier-Mukai-type transformation, namely, the SYZ mirror transformation F , that gives the superpotential W by transforming an object (the function Ψ) in the A-model of X̄.
2. Apparently, the statements written here are slightly different from those in Theorem 1.1 in [7], but realizing that Φ = Exp Ψ, it is easy to see that they are in fact the same statements.
3. The complex oscillatory integrals
Z

Γ

eW Ω Y

of the n-form eW ΩY over Lefschetz thimbles Γ ⊂ Y (defined by the singularities of W : Y → C), which satisfy certain Picard-Fuchs equations, play the role of periods for Calabi-Yau manifolds. This is why we call eW ΩY the holomorphic volume form of the Landau-Ginzburg model (Y, W ).
On the other hand, we also showed that the SYZ mirror transformation (which, in this case, is fiberwise Fourier series) F (Ψi ) of the function Ψi is nothing but the monomial eλi zvi on Y, for i = 1, . . . , d. Since the Jacobian ring Jac(W ) of the superpotential W is generated by the monomials eλ1 zv1 , . . . , eλd zvd , by Proposition 3.2, the SYZ mirror transformation realizes a natural isomorphism between the small quantum cohomology QH ∗ ( X̄ ) and the Jacobian ring Jac(W ).
Theorem 3.2 (Second part of Theorem 1.1 in [7]). The SYZ mirror transformation F
induces a natural isomorphism of C-algebras
∼
=

F : QH ∗ ( X̄ ) −→ Jac(W ), which takes the quantum product, now realized as a convolution product, to the ordinary product of Laurent polynomials, provided that X̄ is a product of projective spaces.
In the example of X̄ = CP2 , the superpotential is the Laurent polynomial q
W (z1 , z2 ) = z1 + z2 + z z2 on Y = (C ∗ )2 , where q = e−t . Its logarithmic partial
1
derivatives are given by q
q
∂1 W = z 1 −
, ∂2 W = z 2 −
, z1 z2
z1 z2
so that the Jacobian ring is given by

q q
, z2 −
Jac(W ) = C [z1±1 , z2±1 ] z1 −
z1 z2
z1 z2

= C [ Z1 , Z2 , Z3 ] Z1 − Z3 , Z2 − Z3 , Z1 Z2 Z3 − q , q

where the monomials Z1 = z1 , Z2 = z2 and Z3 = z1 z2 are the SYZ mirror transformations (i.e. fiberwise Fourier series) of the functions Ψ1 , Ψ2 and Ψ3 respectively.
Remark 3.4.
1. In [10], Coates, Corti, Iritani and Tseng formulated the mirror symmetry conjecture for toric manifolds (and orbifolds) as an isomorphism of graded ∞
2 VHS be∞
tween the A-model ∞
VHS
associated to a
toric manifold and the
B-model
2
2 VHS
associated to the mirror Landau-Ginzburg model (see also Iritani [28]). It is desirable to have this isomorphism, which contains more information than the isomorphism in the above theorem, realized by SYZ mirror transformations.

18

K.-W. CHAN AND N.-C. LEUNG

2. In [15] (and also [16]), Fukaya-Oh-Ohta-Ono applied the machinery developed in [14] to the case of toric manifolds. They considered Floer cohomology with coefficients in the Novikov ring, instead of C used here and in Auroux’s paper
[4]. They have results on the superpotential even in the non-Fano toric case. The isomorphism QH ∗ ( X̄ ) ∼
= Jac(W ) (over the Novikov ring) was also discussed and proved in their work (Theorem 1.9 in [15]). Their proof is combinatorial, using Batyrev’s presentation of the small quantum cohomology ring for toric
Fano manifolds, the validity of which in turn relies on Givental’s mirror theorem.
They claimed that a more conceptual and geometric proof for toric, not necessarily
Fano, manifolds will appear in a sequel to their paper.
3.3. Transformation of branes. This subsection is an attempt to understand the correspondence between A-branes of the toric Fano manifold X̄ and B-branes of the mirror Landau-Ginzburg model (Y, W ) via SYZ mirror transformations.
1
We will deal with the simplest case of the correspondence. So let L x = µ−
X (x)
be the Lagrangian torus fiber of X̄ over a point x ∈ P. We equip L x with a flat U (1)-bundle L y = ( L x × C, ∇y ), where ∇y is the flat U (1)-connection corresponding to y ∈ ( L x )∨ . The mirror of the A-brane ( L x , L y ) is given, according to
√
SYZ, by the B-brane (z = exp(− x − −1y) ∈ Y, Oz ). In other words, the correspondence on the level of objects is the same as in the semi-flat Calabi-Yau case.
Quantum corrections will emerge and make a difference when we consider their endomorphisms.
According to Hori (see [26], Chapter 39), the endomorphism algebra End(z, Oz )
of the B-brane (z, Oz ), as a C-vector space, is given by the cohomology of the complex

(

^∗

Tz Y, δ = ι ∂W ( z)),

where ι ∂W ( z) is contraction with the vector ∂W (z) = ∑nj=1 ∂ j W (z)(∂ j )z and here again ∂ j denotes z j ∂z∂ . The following elementary proposition shows that the inj

troduction of the superpotential W "localizes" the category B-branes to the critical points of W.
Proposition 3.3. The endomorphism End(z, Oz ) is nontrivial if and only if z ∈ Y is a critical pointV of the superpotential W : Y → C, and in which case, End(z, Oz ) is isomorphic to ∗ Tz Y as C-vector spaces.

On the other hand, the endomorphism algebra of the A-brane ( L x , L y ) in the (derived) Fukaya category is given by the Floer cohomology ring HF ( L x , L y ),7
which in turn, as a C-vector space, is given by the cohomology of the Floer complex

( C ∗ ( L x , C ), δ = m1 )
where m1 = m1 ( L x , L y ) denotes the Floer differential. In [9], [8], Cho and Oh explicitly computed the Floer differential m1 . Recall that H 1 ( L x , C ), viewed as the space of infinitestimal deformations of the pair ( L x , L y ), is canonically isomorphic to Tz Y. Let C1 , . . . , Cn be the basis of H 1 ( L x , C ) corresponding to (∂1 )z , . . . , (∂n )z .
7We use C as the coefficient ring, instead of the Novikov ring.

SYZ TRANSFORMATIONS

19

j

Then the results of Cho and Oh stated that m1,β i (Cj ) = Cj · ∂β i = v i and d

m1 ( Cj )

=

1

i =1
d

=

Z

∑ m1,βi (Cj ) exp(− 2π β ωX )hol∇y (∂β i )
i

j

∑ v i z v i = ∂ j W ( z ).

i =1

This shows that m1 = ι ∂W ( z) on H 1 ( L x , C ) = Tz Y, and m1 = 0 on H 1 ( L x , C ) if and only if z is a critical point of W. The following result proved by Cho-Oh in [9] is parallel to the above proposition.
Theorem 3.3 (Cho-Oh [9]). The Floer cohomology HF ( L x , L y ) is nontrivial and isomorphic to H ∗ ( L x , C ) if and only if m1 = 0 on H 1 ( L x , C ).
We conclude that
Theorem 3.4. The Floer cohomology HF ( L x , L y ) of the A-brane ( L x , L y ) is isomorphic to the endomorphism algebra End(z, Oz ) of the mirror B-brane (z, Oz ) as C-vector spaces.

It is intriguing to see whether this isomorphism can be realized by explicit SYZ
mirror transformations.

Remark 3.5. In [8], Cho proved that the Floer cohomology ring HF ( L x , L y ), equipped with the product structure given by m2 = m2 ( L x , L y ), is a Clifford algebra generated by
H 1 ( L x , C ) with the bilinear form given by the Hessian of W: Q(Cj , Ck ) = ∂ j ∂k W (z).
This implies that the isomorphism in Theorem 3.4 is in fact an isomorphism of C-algebras.
This confirms a prediction by Physicists. See the paper of Cho [8] for details.
4. Further questions
The results described in this article represent the first step in our program which is aimed at exploring mirror symmetry via SYZ mirror transformations.
In particular, they showed that these transformations can be applied successfully to explain the mirror symmetry for toric Fano manifolds, a case where quantum corrections do exist. However, we shall emphasize that the quantum corrections in the toric Fano case, which are due to the anticanonical toric divisor, are much simpler than those in the general case (Gross-Siebert [21], Auroux [4]), where quantum corrections may arise due to contributions from the proper singular Lagrangian fibers of the Lagrangian torus fibrations and complicated wall-crossing phenomena start to interfere. In terms of affine geometry, this means that the bases of the Lagranigan torus fibrations in the toric case are affine manifolds with boundary but without singularities, while in the general case, the bases are affine manifolds with both boundary and singularities (and in the semi-flat case, the bases are affine manifolds without boundary and singularities). Certainly much more work remains to be done in the future. In this final section, we will comment on several possible future research directions. The discussion is going to be rather speculative.
4.1. Toric Fano manifolds. We have seen that the simplest correspondence between A-branes on a toric Fano manifold X̄ and B-branes on the mirror LandauGinzburg model (Y, W ), namely

( L x , L y ) ←→ (z, Oz ),

20

K.-W. CHAN AND N.-C. LEUNG

is compatible with the SYZ philosophy. It is desirable to see how other A-branes on X are transformed to the corresponding mirror B-branes on (Y, W ). An interesting and important example would be the Lagrangian submanifold RPn ⊂ CPn for odd n, which can be viewed as a multi-section of the moment map of CPn .

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
