---
source: "https://arxiv.org/abs/1407.0255v2"
title: "Stanley's Major Contributions to Ehrhart Theory"
author: "Matthias Beck"
year: "2014"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1407.0255v2"
pdf: "https://arxiv.org/pdf/1407.0255v2"
captured_at: "2026-06-24T11:10:51Z"
updated_at: "2026-06-24T11:10:51Z"
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

# Stanley's Major Contributions to Ehrhart Theory

- 著者: Matthias Beck
- 年: 2014
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1407.0255v2)
- ダウンロード: https://arxiv.org/pdf/1407.0255v2
- PDF: https://arxiv.org/pdf/1407.0255v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

This expository paper features a few highlights of Richard Stanley's extensive work in Ehrhart theory, the study of integer-point enumeration in rational polyhedra. We include results from the recent literature building on Stanley's work, as well as several open problems.

## PDF Text

STANLEY’S MAJOR CONTRIBUTIONS TO EHRHART THEORY

arXiv:1407.0255v2 [math.CO] 8 Oct 2015

MATTHIAS BECK
Dedicated to Richard Stanley on the occasion of his 70th birthday
A BSTRACT. This expository paper features a few highlights of Richard Stanley’s extensive work in Ehrhart theory, the study of integer-point enumeration in rational polyhedra. We include results from the recent literature building on Stanley’s work, as well as several open problems.

1. I NTRODUCTION
This expository paper features a few highlights of Richard Stanley’s extensive work in Ehrhart theory, with some pointers to the recent literature and open problems. Pre-Stanley times saw two major results in this area: In 1962, Eugène Ehrhart established the following fundamental theorem for a lattice polytope, i.e., the convex hull of finitely many integer points in Rd .
Theorem 1 (Ehrhart [18]). If P ⊂ Rd is a lattice polytope and n ∈ Z>0 then


ehrP (n) := # nP ∩ Zd

evaluates to a polynomial in n (the Ehrhart polynomial of P). Equivalently, the accompanying generating function (the Ehrhart series of P) evaluates to a rational function:
EhrP (x) := 1 + ∑ ehrP (n) xn =
n>0

h∗P (x)
(1 − x)dim(P)+1

for some polynomial h∗P (x) of degree at most dim(P), the Ehrhart h-vector of P.1
We remark that the step from an Ehrhart polynomial to its rational generating function is a mere of
 change

n+k
,
.
.
.
, variables: the coefficients of h∗P (x) express ehrP (n) in the binomial-coefficient basis nk , n+1
k k , where k = dim(P).
In 1971, I.G. Macdonald proved the following reciprocity theorem, which had been conjectured (and proved for several special cases) by Ehrhart.
Theorem 2 (Ehrhart–Macdonald [27]). The evaluation of the Ehrhart polynomial of P at a negative integer yields ehrP (−n) = (−1)dim(P) ehrP ◦ (n) ,
Date: 13 September 2015.
2010 Mathematics Subject Classification. Primary 05A15; Secondary 05A20, 05E40, 52B20.
Key words and phrases. Richard Stanley, Ehrhart polynomial, Ehrhart series, lattice polytope, rational polyhedron, rational generating function, combinatorial reciprocity theorem.
We thank Ben Braun, Martin Henk, Richard Stanley, and Alan Stapledon for helpful discussions and suggestions. This work was partially supported by the U. S. National Science Foundation (DMS-1162638).
1The Ehrhart h-vector is also known by the names of h∗ -vector/polynomial and δ -vector/polynomial.
1

2

MATTHIAS BECK

where P ◦ denotes the (relative) interior of P. Equivalently, we have the following identity of rational functions:
EhrP ( 1x ) = (−1)dim(P)+1 EhrP ◦ (x) , where EhrP ◦ (x) := ∑n>0 ehrP ◦ (n) xn .
Stanley made several fundamental contributions to Ehrhart theory, starting in the 1970s. This paper attempts to highlight some of them, roughly in historical order. The starting point, in Section 2, is Stanley’s proof of the Anand–Dumir–Gupta conjecture. Section 3 features a reciprocity theorem of Stanley that generalizes Theorem 2, and Section 4 contains Stanley’s inequalities on Ehrhart h-vectors. Throughout the paper we mention open problems, and Section 5 gives some recent results in Ehrhart theory building on
Stanley’s work.
2. “I AM GRATEFUL TO K. BACLAWSKI FOR CALLING MY ATTENTION TO THE WORK OF E UG ÈNE
E HRHART...”
The starting point for Stanley’s work in Ehrhart theory is arguably his proof of the Anand–Dumir–Gupta conjecture [1]; this proof appeared in a 1973 paper from which the quote of the section header is taken [35, p. 631]. (One could, in fact, argue that P-partitions, which were introduced in Stanley’s Ph.D. thesis [34] and are featured in Gessel’s article in this volume, already show an Ehrhart-theoretic flavor, but this geometric realization came slightly later in Stanley’s work.) The conjecture concerns the counting function Hn (r), the number of (n × n)-matrices with nonnegative integer entries that sum to r in every row and column; these matrices are often referred to as semimagic squares.
The function Hn (r) goes back to MacMahon [28], who

r+2
computed the first nontrivial case, H3 (r) = r+5
−
5
5 . Anand, Dumir, and Gupta conjectured that
• Hn (r) is a polynomial in r for any fixed n,
• this polynomial has roots at −1, −2, . . . , −n + 1, and
• it satisfies the symmetry relation Hn (−r) = (−1)n−1 Hn (r − n).
Stanley showed that one could use what is now called the Elliott–MacMahon algorithm (going back to [20]
and [28]; see also [2] and its dozen follow-up papers) in connection with the Hilbert syzygy theorem (see, e.g., [19]) to prove the Anand–Dumir–Gupta conjecture.2 Stanley realized that Hn (r) is closely related to the geometry of the Birkhoff–von Neumann polytope Bn , the set of all doubly-stochastic (n × n)-matrices.
In modern language, Hn (r) equals the Ehrhart polynomial of Bn , and because Bn is integrally closed (essentially by the Birkhoff–von Neumann theorem that the extreme points of Bn are precisely the permutation matrices; see, e.g., [13] for more on integrally closed polytopes), this Ehrhart series equals the Hilbert series of the semigroup algebra generated by the integer point in Bn × {1}, graded by the last coordinate.
Stanley proved that this Hilbert series has the following properties, which are a direct translation of the
Anand–Dumir–Gupta conjecture into the language of generating functions:
Theorem 3 (Stanley [35]). For each n there exists a palindromic polynomial hn (x) of degree n2 − 3n + 2
such that hn (x)
.
1 + ∑ Hn (r) xr =
2
(1 − x)n −2n+2
r>0
One could skip the detour through commutative algebra and directly realize this rational generating function as the Ehrhart series of Bn , with the needed properties to affirm the Anand–Dumir–Gupta conjecture.
Nevertheless, the algebraic detour is worth taking; aside from its inherent elegance, it allowed Stanley to
2Stanley’s comment in [47, p. 6] is amusing in this context: “I had taken a course in graduate school on commutative algebra

that I did not find very interesting. It did not cover the Hilbert syzygy theorem. I had to learn quite a bit of commutative algebra from scratch in order to understand the work of Hilbert.”

STANLEY’S MAJOR CONTRIBUTIONS TO EHRHART THEORY
m2

3

2

n
1 m2
n realize that the monomials zm
1 z2 · · · zn2 , where (m1 , m2 , . . . , mn2 ) ∈ Z≥0 ranges over all semimagic squares, generate a Cohen–Macaulay algebra, and this implies, following Hochster’s work [25]:

Theorem 4 (Stanley [38]). The polynomial hn (x) in Theorem 3 has nonnegative coefficients.
Stanley conjectured that the coefficients of hn (x) in Theorem 3 are also unimodal (the coefficients increase up to some point and then decrease), which was proved by Athanasiadis some three decades later [4]; we will say more about this in Section 5.
Another problem connected to Theorem 3 (mentioned by Stanley but certainly older than his work) is still wide open, namely, the quest for the volume of Bn , which equals hn (1) after a suitable normalization.
This volume is known only for n ≤ 10, though there has been recent progress, e.g., in terms of asymptotic and combinatorial formulas [6, 15, 16].
It is worth noting that Stanley proved versions of Theorems 3 and 4 for the more general magic labellings of graphs (and semimagic squares correspond to such labellings for a complete bipartite graph Knn ). In this more general context, the associated counting functions become quasipolynomials of period 2 (see, e.g., [46, Chapter 4] for more about quasipolynomials), foreshadowing in some sense Zaslavsky’s work on enumerative properties of signed graphs: Stanley’s magic labellings are essentially flows on all-negative signed graphs [9, 52].
Stanley’s work on the Anand–Dumir–Gupta conjecture was not just the starting point of his contributions to Ehrhart theory. As he mentions in [47], it opened the door to what are now called Stanley–Reisner rings and the use of Cohen–Macaulay algebras in geometric combinatorics, famously leading to Stanley’s proof of the upper bound conjecture for spheres [37]. Stanley’s appreciation for the polynomials Hn (r) is also evident in his writings: they are prominently featured in his influential books [45, 46]. We close this section by mentioning that a highly readable account of commutative-algebra concepts behind the Anand–Dumir–
Gupta conjecture can be found in [12].
3. R ECIPROCITY
Theorem 2 is an example of a combinatorial reciprocity theorem: we get interesting information out of a counting function when we evaluate it at a negative integer (and so, a priori the counting function does not make sense at this number). We remark that the last part of the Anand–Dumir–Gupta conjecture follows from Theorem 2 applied to the Birkhoff–von Neumann polytope Bn , as one can easily show that ehrBn◦ (r) = Hn (r − n).
Stanley had discovered reciprocity theorems for P-partitions and order polynomials in his thesis [34], and so it was natural for
 him to realize Theorem 2 as a special case of a wider phenomenon. A rational cone is a set of the form x ∈ Rd : A x ≤ 0 for some integer matrix A. It is not hard to see that the multivariate generating function md
1 m2
zm
σK (z1 , z2 , . . . , zd ) :=
∑
1 z2 · · · zd
(m1 ,m2 ,...,md )∈K ∩Zd

evaluates to a rational function in z1 , z2 , . . . , zd if K is a rational cone (see, e.g., [7, Chapter 3]). Stanley proved that this rational function satisfies a reciprocity theorem.
Theorem 5 (Stanley [36]). If K is a rational cone then


σK z11 , z12 , . . . , z1d = (−1)dim(K ) σK ◦ (z1 , z2 , . . . , zd ) .

A simple proof of Theorem 5, based on the ideas of [8], can be found in [7, Chapter 4]. True to the theme of this survey, our description of Theorem 5 is geometric, while Stanley’s viewpoint presented in [36]
is based on integral solutions of a system of integral linear equations; in this language, the reciprocity is between nonnegative and positive solutions. (It is not hard to see that both viewpoints are equivalent.) This

4

MATTHIAS BECK

language also connects once more to the Elliott–MacMahon algorithm mentioned in Section 2, and Stanley uses this connection in [36] to present his monster reciprocity theorem, which (a bit oversimplified) can be thought of an affine version of Theorem 5. It has been recently revitalized by Xin [51].
We finish this section with a sketch how Theorem 2 follows as a corollary of Theorem 5. Given a lattice polytope P ⊂ Rd , we consider its homogenization cone(P) :=

∑

R≥0 (v, 1)

v vertex of P

by lifting the vertices of P into Rd+1 to the hyperplane xd+1 = 1 and taking the nonnegative span of this lifted version of P. Thus (by the Minkowski–Weyl theorem—see, e.g., [53, Lecture 1]) cone(P) is a rational cone, and because cone(P) ∩ {x ∈ Rd : xd+1 = n} is identical to nP (embedded in {x ∈ Rd : xd+1 = n}),
EhrP (x) = σcone(P) (1, 1, . . . , 1, x) .
Theorem 2 follows now by specializing all but one variable in Theorem 5.
4. E HRHART INEQUALITIES
Theorem 4 generalizes to all lattice polytopes, and this is arguably Stanley’s most important contribution to the intrinsic study of Ehrhart polynomials.
Theorem 6 (Stanley [40]). For any lattice polytope P, the Ehrhart h-vector h∗P (x) has nonnegative coefficients.
Even though Stanley explicitly states this result first in [40] (and gives a proof using a shelling triangulation argument), he attributes Theorem 6 to [35, Proposition 4.5]—the magic-graph-labelling version of
Theorem 4. Theorem 6 can be viewed as a starting point for the problem of classifying Ehrhart polynomials/Ehrhart h-vectors. This problem is wide open, already in dimension 3. (In dimension 2, it is essentially solved by Pick’s theorem [30] and an inequality of Scott [33].)
As part of his study of monotonicity of h-vectors of Cohen–Macaulay complexes, Stanley deduced the following refinement of Theorem 6.
Theorem 7 (Stanley [44]). If P ⊆ Q are lattice polytopes then h∗P (x) ≤ h∗Q (x) (component-wise).
Theorem 6 can be realized as a corollary of Theorem 7 by choosing P to be a unimodular simplex (i.e.,
1
a d-dimensional lattice polytope of volume d!
, which comes with the Ehrhart h-vector h∗P (x) = 1).
One can show that Theorems 6 and 7 also hold for rational d-polytopes P and Q (i.e., polytopes whose vertices have rational coordinates) if their Ehrhart series are written in the form h∗P/Q (x)
(1 − x p )d+1
for some p ∈ Z>0 such that pP and pQ are lattice polytopes. (The accompanying Ehrhart counting functions for P and Q are then quasipolynomials, and p is a period of them.) The arguably simplest proofs of
(rational versions of) Theorems 6 and 7 are given in [8].
A natural question is whether there are any natural upper bounds complementing Theorem 6. Of course, the volume (and therefore h∗P (1), the sum of the Ehrhart-h coefficients) of a lattice polytope P can be arbitrarily large, but one can ask for upper bounds given certain data.
Theorem 8 (Haase–Nill–Payne [21]). The volume of a lattice polytope P (and therefore also the coefficients of h∗P (x)) is bounded by a number that depends only on the degree and the leading coefficient of h∗P (x).

STANLEY’S MAJOR CONTRIBUTIONS TO EHRHART THEORY

5

This result was conjectured by Batyrev [5] and improves a classic theorem of Lagarias and Ziegler [26]
that the volume of a lattice polytope is bounded by a number depending only on its dimension and the number of its interior lattice points, if the latter is positive.
Theorem 6 can be extended in a direction different from that of Theorem 7, namely, one can establish inequalities among the coefficients of an Ehrhart h-vector. Stanley derived one set of such inequalities as a corollary of his study of Hilbert functions of semistandard graded Cohen–Macaulay domains [42]; thus the following result holds in a more general situation. The degree of a lattice polytope P is the degree of its
Ehrhart h-vector h∗P (x).
Theorem 9 (Stanley [42]). If P is a d-dimensional lattice polytope of degree s then its Ehrhart h-vector h∗P (x) = h∗s xs + h∗s−1 xs−1 + · · · + h∗0 satisfies h∗0 + h∗1 + · · · + h∗j ≤ h∗s + h∗s−1 + · · · + h∗s− j

for

0 ≤ j ≤ d.

(In the above theorem and below, we define h∗j = 0 whenever j < 0 or j is larger than the degree of P.)
Theorem 9 complements inequalities discovered by Hibi around the same time [22, 24]; they were more recently improved by Stapledon. Together with Theorem 9 and the trivial inequality h∗1 ≥ h∗d (which follows from the facts h∗1 = #(P ∩ Zd ) − d − 1 and h∗d = #(P ◦ ∩ Zd )), the following result gives the state of the art in terms of linear constraints for the Ehrhart coefficients that can be easily written down in general.
Theorem 10 (Stapledon [49]). If P is a d-dimensional lattice polytope of degree s and codegree l :=
d + 1 − s, then its Ehrhart h-vector h∗P (x) = h∗s xs + h∗s−1 xs−1 + · · · + h∗0 satisfies h∗2 + h∗3 + · · · + h∗j+1 ≥ h∗d−1 + h∗d−2 + · · · + h∗d− j

for

0 ≤ j ≤ ⌊ d2 ⌋ − 1,

h∗2−l + h∗3−l + · · · + h∗1 ≤ h∗j + h∗j−1 + · · · + h∗j−l+1

for

2 ≤ j ≤ d − 1.

We will say more about this theorem in Section 5 and finish this section with the remark that a unimodular d-simplex satisfies each of the inequalities of Theorems 9 and 10 with equality.
5. S TANLEY & BEYOND
The Ehrhart h-vector is philosophically close to the h-vector of a simplicial complex. This statement can be made much more precise, as we will show in this final section, in which we give a flavor of recent results that build on Stanley’s work in Ehrhart theory.
The starting point can once more be found in Stanley’s papers; the following result follows essentially from the definition d

hT (z) := ∑ fk zk+1 (1 − z)d−k k=−1

of the h-vector of a given triangulation T of a d-dimensional polytope (here fk denotes the number of ksimplices in T ) and the fact that a unimodular simplex has a trivial Ehrhart h-vector (a triangulation is unimodular if all of its simplices are). Nevertheless, the following identity is an important base case for structural properties of Ehrhart h-vectors to come.
Theorem 11 (Stanley [40]). If P is a lattice polytope that admits a unimodular triangulation T then hT (z)
.
EhrP (z) =
(1 − z)dim(P)+1
In words, the Ehrhart h∗ -vector of P is given by the h-vector of T .
The hope is now to use properties of the h-vector of a triangulation to say something about Ehrhart h-vectors; for example, if T is the cone over a boundary triangulation of P then hT (z) satisfies the Dehn–
Sommerville equations (see, e.g., [17] for more about triangulations). Unfortunately, not all lattice polytopes

6

MATTHIAS BECK

admit unimodular triangulations in dimension ≥ 3—in fact, most do not—and so Theorem 11 needs some tweaking before we can apply it to general lattice polytopes. This tweaking, due to Betke and McMullen
[10], has two main ingredients: the link of a simplex ∆ in a triangulation T
link(∆) := {Ω ∈ T : Ω ∩ ∆ = ∅, Ω ⊆ Φ for some Φ ∈ T with ∆ ⊆ Φ} , and its box polynomial

∑

B∆ (x) :=

xheight(m)

m∈Π(∆)∩Zd+1

where we define
Π(∆) :=

(

∑

)

λv (v, 1) : 0 < λv < 1

v vertex of ∆

and height(m) denotes the last coordinate of m. (Geometrically, Π(∆) is the open fundamental parallelepiped of cone(∆).) For the empty simplex ∅ of a triangulation, we set B∅ (x) = 1.
Theorem 12 (Betke–McMullen [10]). Fix a triangulation T of the lattice polytope P. Then h∗P (x) = ∑ hlink(∆) (x) B∆ (x) .
∆∈T

If a simplex ∆ ∈ T is unimodular then B∆ (x) = 0, unless ∆ = ∅. Thus, if T is unimodular then the sum in Theorem 12 collapses to hlink(∅) (x) B∅ (x) = hT (x), and so Theorem 11 is a corollary to Theorem 12.
(This argument also shows, in general, that h∗P (x) ≥ hT (x) component-wise.) Furthermore, since all ingredients for the sum in Theorem 12 are nonnegative, this gives another (and the first combinatorial) proof of
Theorem 6.
Theorem 12 was greatly extended by Stanley in (and served as some motivation to) his work on local h-vectors of subdivisions [43]; see Athanasiadis’ contribution to this volume. Payne gave a different, multivariate generalization of Theorem 12 in [29].
Theorem 12 has a powerful consequence when P has an interior lattice point; this consequence was fully realized only by Stapledon [49] who extended it to general lattice polytopes—Theorem 14 below. Namely, if a lattice polytope P has an interior lattice point, it admits a regular triangulation that is a cone (at this point)
over a boundary triangulation. This has the charming effect that each hlink(∆) (x) appearing in Theorem 12
is palindromic (due to the afore-mentioned Dehn–Sommerville equations). Since the box polynomials are palindromic and both kinds of polynomials have nonnegative coefficients, a little massaging of the identity in Theorem 12 gives:
Corollary 13. Suppose P is a d-dimensional lattice polytope that contains an interior lattice point. Then there exist unique polynomials a(x) and b(x) with nonnegative coefficients such that h∗P (x) = a(x) + x b(x) , a(x) = xd a( 1x ), and b(x) = xd−1 b( 1x ).
The identities for a(x) and b(x) say that a(x) and b(x) are palindromic polynomials; the degree of a(x)
is necessarily d, while the degree of b(x) is d − 1 or smaller; in fact, b(x) can be zero—this happens if and only if P is the translate of a reflexive polytope (i.e., a lattice polytope whose dual is also a lattice polytope), due to Theorem 15 below. Stapledon recently introduced a weighted variant of the Ehrhart h∗ -vector which is always palindromic, motivated by motivic integration and the cohomology of certain toric varieties [48].
One can easily recover h∗P (x) from this weighted Ehrhart h∗ -vector, but one can also deduce the palindromy of both a(x) and b(x) as coming from the same source (and this perspective has some serious geometric applications).

STANLEY’S MAJOR CONTRIBUTIONS TO EHRHART THEORY

7

The statements that a(x) and b(x) in Corollary 13 have nonnegative coefficients are straightforward translations of Hibi’s and Stanley’s inequalities on Ehrhart h-vectors mentioned above (right after and in Theorem 9), in the case that the dimension and the degree of P are equal (which is equivalent to P containing an interior lattice point). The full generality of Theorem 9 as well as Theorem 10 follow from the following generalization of Corollary 13.
Theorem 14 (Stapledon [49]). Suppose P is a d-dimensional lattice polytope of degree s and codegree l = d + 1 − s. Then there exist unique polynomials a(x) and b(x) with nonnegative coefficients such that


1 + x + · · · + xl−1 h∗P (x) = a(x) + xl b(x) , a(x) = xd a( 1x ), b(x) = xd−l b( 1x ), and, writing a(x) = ad xd + ad−1 xd−1 + · · · + a0 ,
1 = a0 ≤ a1 ≤ a j

for

2 ≤ j ≤ d − 1.

Stapledon has recently improved this theorem further, giving infinitely many classes of linear inequalities among Ehrhart-h coefficients [50]. This exciting new line of research involves additional techniques from additive number theory.
One can, on the other hand, ask which classes of polytopes satisfy more special sets of equalities or inequalities. Arguably the most natural such equalities/inequalities are those expressing palindromy and unimodality.
Lattice polytopes with palindromic Ehrhart h-vectors are completely classified by the following theorem, which first explicitly surfaced in Hibi’s work on reflexive polytopes but can be traced back to Stanley’s work on Hilbert functions of Gorenstein rings.
Theorem 15 (Hibi–Stanley [23, 39]). If P is a lattice polytope of degree s and codegree l = d + 1 − s, then its Ehrhart h-vector is palindromic if and only if lP is a translate of a reflexive polytope.
The following result is a start towards a unimodality classification; it was proved by Athanasiadis [3] and independently by Hibi and Stanley (unpublished).
Theorem 16 (Athanasiadis–Hibi–Stanley [3]). If the d-dimensional lattice polytope P admits a regular unimodular triangulation, then h∗⌊ d+1 ⌋ ≥ · · · ≥ h∗d−1 ≥ h∗d
2

and

 ∗

h1 + j − 1
∗
hj ≤
j

for

0 ≤ j ≤ d.

Stapledon’s work in [49] implies further that if the boundary of P admits a regular unimodular triangulation, then h∗j+1 ≥ h∗d− j for 0 ≤ j ≤ ⌊ d2 ⌋ − 1
(which was also proved in [3] under the stronger assumption that P admits a regular unimodular triangulation) and
 ∗

h − h∗d + j + 1
h∗0 + · · · + h∗j+1 ≤ h∗d + · · · + h∗d− j + 1
for 0 ≤ j ≤ ⌊ d2 ⌋ − 1.
j+1
Naturally, if in addition to the conditions in Theorem 16, P has degree d and h∗P (x) is palindromic, then h∗P (x) is unimodal. The proof of Theorem 16 starts with Theorem 11 and then shows that the h-vector of the unimodular triangulation satisfies the stated inequalities. Athanasiadis’ approach was inspired by work of Reiner and Welker on order polytopes of graded posets and a connection between the Charney–
Davis and Neggers–Stanley conjectures [31] and can be taken further: Athanasiadis used similar methods to prove Stanley’s conjecture mentioned in Section 2 that the Ehrhart h-vector of the Birkhoff–von Neumann

8

MATTHIAS BECK

polytope is unimodal [4]. Bruns and Römer generalized this to any Gorenstein polytope that admits a regular unimodular triangulation [14].
Going into a somewhat different direction, Schepers and van Langenhoven recently proved that lattice parallelepipeds have a unimodal Ehrhart h-vector [32]. The conjecture that any integrally closed polytope
(of which both parallelepipeds and the Birkhoff–von Neumann polytope are examples) has a unimodal
Ehrhart h-vector remains open, even for integrally closed reflexive polytopes (though recent work of Braun and Davis give some pointers of what could be tried here [11]); this is closely related to a conjecture of
Stanley that every standard graded Cohen–Macaulay domain has a unimodal h-vector [41].
R EFERENCES
1. Harsh Anand, Vishwa Chander Dumir, and Hansraj Gupta, A combinatorial distribution problem, Duke Math. J. 33 (1966),
757–769.
2. George E. Andrews, MacMahon’s partition analysis. I. The lecture hall partition theorem, Mathematical essays in honor of
Gian-Carlo Rota (Cambridge, MA, 1996), Progr. Math., vol. 161, Birkhäuser Boston, Boston, MA, 1998, pp. 1–22.
3. Christos A. Athanasiadis, h∗ -vectors, Eulerian polynomials and stable polytopes of graphs, Electron. J. Combin. 11 (2004/06), no. 2, Research Paper 6, 13 pp. (electronic).
, Ehrhart polynomials, simplicial polytopes, magic squares and a conjecture of Stanley, J. Reine Angew. Math. 583
4.
(2005), 163–174, arXiv:math/0312031.
5. Victor V. Batyrev, Lattice polytopes with a given h∗ -polynomial, Algebraic and geometric combinatorics, Contemp. Math., vol.
423, Amer. Math. Soc., Providence, RI, 2006, pp. 1–10, arXiv:math.CO/0602593.
6. Matthias Beck and Dennis Pixton, The Ehrhart polynomial of the Birkhoff polytope, Discrete Comput. Geom. 30 (2003), no. 4,
623–637, arXiv:math.CO/0202267.
7. Matthias Beck and Sinai Robins, Computing the Continuous Discretely: Integer-point Enumeration in Polyhedra, Undergraduate Texts in Mathematics, Springer, New York, 2007, electronically available at http://math.sfsu.edu/beck/ccd.html.
8. Matthias Beck and Frank Sottile, Irrational proofs for three theorems of Stanley, European J. Combin. 28 (2007), no. 1, 403–
409, arXiv:math.CO/0506315.
9. Matthias Beck and Thomas Zaslavsky, The number of nowhere-zero flows on graphs and signed graphs, J. Combin. Theory
Ser. B 96 (2006), no. 6, 901–918, arXiv:math.CO/0309331.
10. Ulrich Betke and Peter McMullen, Lattice points in lattice polytopes, Monatsh. Math. 99 (1985), no. 4, 253–265.
11. Benjamin Braun and Robert Davis, Ehrhart series, unimodality, and integrally closed reflexive polytopes, Preprint
(arXiv:math/1403.5378), 2014.
12. Winfried Bruns, Commutative algebra arising from the Anand–Dumir–Gupta conjectures, Commutative algebra and combinatorics, Ramanujan Math. Soc. Lect. Notes Ser., vol. 4, Ramanujan Math. Soc., Mysore, 2007, pp. 1–38.
13. Winfried Bruns and Joseph Gubeladze, Polytopes, rings, and K-theory, Springer Monographs in Mathematics, Springer, Dordrecht, 2009.
14. Winfried Bruns and Tim Römer, h-vectors of Gorenstein polytopes, J. Combin. Theory Ser. A 114 (2007), no. 1, 65–76.
15. E. Rodney Canfield and Brendan D. McKay, The asymptotic volume of the Birkhoff polytope, Online J. Anal. Comb. (2009), no. 4, 4 pages, arXiv:math.0705.2422.
16. Jesús A. De Loera, Fu Liu, and Ruriko Yoshida, A generating function for all semi-magic squares and the volume of the Birkhoff polytope, J. Algebraic Combin. 30 (2009), no. 1, 113–139.
17. Jesús A. De Loera, Jörg Rambau, and Francisco Santos, Triangulations, Algorithms and Computation in Mathematics, vol. 25,
Springer-Verlag, Berlin, 2010.
18. Eugène Ehrhart, Sur les polyèdres rationnels homothétiques à n dimensions, C. R. Acad. Sci. Paris 254 (1962), 616–618.
19. David Eisenbud, Commutative Algebra With a View Toward Algebraic Geometry, Graduate Texts in Mathematics, vol. 150,
Springer-Verlag, New York, 1995.
20. E. B. Elliott, On linear homogeneous diophantine equations, Quartely J. Pure Appl. Math. 34 (1903), 348–377.
21. Christian Haase, Benjamin Nill, and Sam Payne, Cayley decompositions of lattice polytopes and upper bounds for h∗ polynomials, J. Reine Angew. Math. 637 (2009), 207–216, arXiv:math/0804.3667.
22. Takayuki Hibi, Some results on Ehrhart polynomials of convex polytopes, Discrete Math. 83 (1990), no. 1, 119–121.
23.
, Dual polytopes of rational convex polytopes, Combinatorica 12 (1992), no. 2, 237–240.
, A lower bound theorem for Ehrhart polynomials of convex polytopes, Adv. Math. 105 (1994), no. 2, 162–165.
24.
25. Melvin Hochster, Rings of invariants of tori, Cohen–Macaulay rings generated by monomials, and polytopes, Ann. of Math.
(2) 96 (1972), 318–337.

STANLEY’S MAJOR CONTRIBUTIONS TO EHRHART THEORY

9

26. Jeffrey C. Lagarias and Günter M. Ziegler, Bounds for lattice polytopes containing a fixed number of interior points in a sublattice, Canad. J. Math. 43 (1991), no. 5, 1022–1035.
27. Ian G. Macdonald, Polynomials associated with finite cell-complexes, J. London Math. Soc. (2) 4 (1971), 181–192.
28. Percy A. MacMahon, Combinatory Analysis, Chelsea Publishing Co., New York, 1960, reprint of the 1915 original.
29. Sam Payne, Ehrhart series and lattice triangulations, Discrete Comput. Geom. 40 (2008), no. 3, 365–376, arXiv:math/0702052.
30. Georg Alexander Pick, Geometrisches zur Zahlenlehre, Sitzenber. Lotos (Prague) 19 (1899), 311–319.
31. Victor Reiner and Volkmar Welker, On the Charney–Davis and Neggers–Stanley conjectures, J. Combin. Theory Ser. A 109
(2005), no. 2, 247–280.
32. Jan Schepers and Leen Van Langenhoven, Unimodality questions for integrally closed lattice polytopes, Ann. Comb. 17 (2013), no. 3, 571–589, arXiv:math/1110.3724.
33. Paul R. Scott, On convex lattice polygons, Bull. Austral. Math. Soc. 15 (1976), no. 3, 395–399.
34. Richard P. Stanley, Ordered structures and partitions, American Mathematical Society, Providence, R.I., 1972, Memoirs of the
American Mathematical Society, No. 119.
35.
, Linear homogeneous Diophantine equations and magic labelings of graphs, Duke Math. J. 40 (1973), 607–632.
, Combinatorial reciprocity theorems, Advances in Math. 14 (1974), 194–253.
36.
37.
, The upper bound conjecture and Cohen–Macaulay rings, Studies in Appl. Math. 54 (1975), no. 2, 135–142.
38.
, Magic labelings of graphs, symmetric magic squares, systems of parameters, and Cohen–Macaulay rings, Duke Math.
J. 43 (1976), no. 3, 511–531.
, Hilbert functions of graded algebras, Advances in Math. 28 (1978), no. 1, 57–83.
39.
40.
, Decompositions of rational convex polytopes, Ann. Discrete Math. 6 (1980), 333–342.
41.
, Log-concave and unimodal sequences in algebra, combinatorics, and geometry, Graph theory and its applications:
East and West (Jinan, 1986), Ann. New York Acad. Sci., vol. 576, New York Acad. Sci., New York, 1989, pp. 500–535.
, On the Hilbert function of a graded Cohen–Macaulay domain, J. Pure Appl. Algebra 73 (1991), no. 3, 307–314.
42.
43.
, Subdivisions and local h-vectors, J. Amer. Math. Soc. 5 (1992), no. 4, 805–851.
, A monotonicity property of h-vectors and h∗ -vectors, European J. Combin. 14 (1993), no. 3, 251–258.
44.
45.
, Combinatorics and Commutative Algebra, second ed., Progress in Mathematics, vol. 41, Birkhäuser Boston Inc.,
Boston, MA, 1996.
46.
, Enumerative Combinatorics. Volume 1, second ed., Cambridge Studies in Advanced Mathematics, vol. 49, Cambridge
University Press, Cambridge, 2012.
, How the upper bound conjecture was proved, http://www-math.mit.edu/∼rstan/papers.html, 2013.
47.
48. Alan Stapledon, Weighted Ehrhart theory and orbifold cohomology, Adv. Math. 219 (2008), no. 1, 63–88, arXiv:math/0711.4382.
, Inequalities and Ehrhart δ -vectors, Trans. Amer. Math. Soc. 361 (2009), no. 10, 5615–5626,
49.
arXiv:math/0801.0873.
50.
, Additive number theory and inequalities in Ehrhart theory, Preprint (arXiv:0904.3035v2), 2010.
51. Guoce Xin, Generalization of Stanley’s monster reciprocity theorem, J. Combin. Theory Ser. A 114 (2007), no. 8, 1526–1544, arXiv:math/0504425.
52. Thomas Zaslavsky, Orientation of signed graphs, European J. Combin. 12 (1991), no. 4, 361–375.
53. Günter M. Ziegler, Lectures on Polytopes, Springer-Verlag, New York, 1995.
D EPARTMENT OF M ATHEMATICS , S AN F RANCISCO S TATE U NIVERSITY , S AN F RANCISCO , CA 94132, U.S.A.
E-mail address: mattbeck@sfsu.edu
URL: http://math.sfsu.edu/beck

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
