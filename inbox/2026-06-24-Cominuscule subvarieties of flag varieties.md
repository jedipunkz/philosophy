---
source: "https://arxiv.org/abs/2002.08193v8"
title: "Cominuscule subvarieties of flag varieties"
author: "Benjamin McKay"
year: "2020"
publication: "arXiv preprint / math.AG"
download: "https://arxiv.org/pdf/2002.08193v8"
pdf: "https://arxiv.org/pdf/2002.08193v8"
captured_at: "2026-06-24T11:11:11Z"
updated_at: "2026-06-24T11:11:11Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Richard McKay Rorty"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Cominuscule subvarieties of flag varieties

- 著者: Benjamin McKay
- 年: 2020
- 掲載情報: arXiv preprint / math.AG
- 情報源: [arxiv](https://arxiv.org/abs/2002.08193v8)
- ダウンロード: https://arxiv.org/pdf/2002.08193v8
- PDF: https://arxiv.org/pdf/2002.08193v8

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We show that every flag variety contains a naturally defined homogeneous cominuscule subvariety. From the Dynkin diagram of the flag variety, we compute the Dynkin diagram of that subvariety.

## PDF Text

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES
BENJAMIN McKAY

Abstract. We show that every flag variety contains a naturally defined homogeneous cominuscule subvariety. From the Dynkin diagram of the flag variety, we compute the Dynkin diagram of that subvariety. We also investigate the tangent bundles of flag varieties.

Creative Commons Licence, Minneapolis College of Art and Design
Attribution 2.0 Generic (CC BY 2.0)

Date: March 30, 2026.
Key words and phrases. flag variety, Hermitian symmetric space.
This research was supported in part by the International Centre for Theoretical Sciences
(ICTS) during a visit for participating in the program Analytic and Algebraic Geometry (Code:
ICTS/aag2018/03). I am grateful for the hospitality of the University of Catania, where this paper was mostly written, and in particular to Francesco Russo. Thanks to Indranil Biswas, Anca
Mustaţă and Andrei Mustaţă for help with algebraic geometry. This article/publication is based upon work from COST Action CaLISTA, CA21109, supported by COST (European Cooperation in Science and Technology), www.cost.eu.
1

2

BENJAMIN McKAY

Contents
1. Introduction, 3
1.1. The algorithm in examples, 3
1.2. Why the associated cominuscule subvariety matters, 4
2. Review of flag varieties, 6
2.1. Flag varieties, 6
2.2. Reducible flag varieties, 6
2.3. Cominuscule varieties, 6
2.4. Structure of linear algebraic groups, 7
2.5. Parabolic subgroups, 7
2.6. Opposite parabolic subgroups, 7
3. Definition of the associated cominuscule subvariety, 8
3.1. The definition, 8
3.2. Example: the general linear flag variety, 8
4. Dynkin diagrams, 10
4.1. Compact roots, 10
4.2. The Dynkin diagram of a flag variety, 10
4.3. Affine extensions, 10
4.4. The extended Dynkin diagram of a flag variety, 11
5. Reducing to root systems, 11
5.1. Gradings, 11
5.2. The box and the Lie algebra, 11
5.3. Associated cominuscule subvarieties in rank 2, 12
6. The associated cominuscule subvariety is cominuscule, 14
6.1. Roots and opposite parabolics, 16
6.2. Proof of the cominuscule property, 16
6.3. Automorphisms of the associated cominuscule subvariety, 18
6.4. Computing the automorphism Lie algebra, 19
7. The Hasse diagram of a root system, 21
8. Homogeneous vector bundles, 23
8.1. Filtration and grading, 23
8.2. The Hasse diagram of a flag variety, 23
8.3. The Hasse diagram of a cominuscule variety, 25
9. The algorithm is correct, 27
9.1. Reducing to the irreducible, 27
9.2. Finding the box, 27
9.3. The lowest box root, 29
9.4. The Levi Dynkin diagram, 30
9.5. Edge multiplicities, 31
9.6. Symmetry of the box, 32
9.7. Outer automorphisms, 33
9.8. Using the symmetry, 39
10. Freedom, 39
10.1. Defining freedom, 39
10.2. Free morphisms, 40
10.3. Automorphisms, 40
10.4. Root systems, 40
10.5. Chevalley bases, 41
10.6. Dual bases, 41
10.7. Proving the freedom theorems, 42
11. Conclusion, 46
Appendix A. Refining the filtrations and gradings, 47

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

3

A.1. The problem, 47
A.2. Graded vector spaces, 47
A.3. Grading by groups, 47
A.4. Graded Lie algebras, 47
A.5. Graded modules, 47
A.6. Filtered vector spaces, 48
A.7. Changing partial ordering, 48
A.8. Ordered groups, 49
A.9. Filtering by ordered groups, 49
A.10. Filtered Lie algebras, 49
A.11. Filtered modules, 50
A.12. Augmentation, 50
A.13. Underlying integer filtration, 50
A.14. Underlying integer grading, 51
A.15. Flag varieties and filtrations, 51
A.16. Hasse diagrams of homogeneous vector bundles, 52
A.17. The associated graded as a module, 52
A.18. The Hasse diagram and irreducibles, 53
Appendix B. The Dynkin diagrams of flag varieties and their associated cominuscule subvarieties, 56
Appendix C. The Hasse diagrams of the exceptional simple Lie groups, 60
Appendix D. The Hasse diagrams of the cominuscule varieties, 62
References, 65

1. Introduction
We cannot draw a flag variety, or even its associated root system (except in low dimensions), but we can always draw its Hasse diagram. The Hasse diagrams as drawn by Claus Ringel [50] are very clear, so we will follow his conventions. Our eyes immediately spot in that Hasse diagram its uppermost component. This component is always the Hasse diagram of a unique cominuscule subvariety. We correctly predict that each flag variety contains an associated homogeneous cominuscule subvariety.
We see the root system of the subvariety inside the root system of the flag variety.
Theorem 1. Take any flag variety. Draw its extended Dynkin diagram (defined in §4.4 on page 11). Erase every crossed node and its adjacent edges. Erase any component not connected to an affine node. Redraw every affine node as a crossed node. This yields the Dynkin diagram of the associated cominuscule subvariety
(defined in §3 on page 8).
To clarify the theorem, in table 6 on page 56, we draw the Dynkin diagram of every irreducible flag variety and the Dynkin diagram of its associated cominuscule subvariety.
1.1. The algorithm in examples.
Example 1. The extended Dynkin diagram of E8 is

Take the E8 -flag variety with Dynkin diagram

4

BENJAMIN McKAY

Its extended Dynkin diagram is

Let us apply the algorithm of our main theorem, theorem 1 on the previous page.
Remove the crossed node, giving D8 but with one root still drawn as a hollow :

Change the hollow dot

to a cross :

As we will prove, this is the Dynkin diagram of the associated cominuscule subvariety.
Example 2. Instead take the C7 -flag variety with Dynkin diagram

Its extended Dynkin diagram is

Let us apply the algorithm of our main theorem, theorem 1 on the preceding page.
Remove the crossed node and its adjacent edges:

Drop all components not connected to a .

Change the hollow dot

to a cross :

As we will prove, this is the Dynkin diagram of the associated cominuscule subvariety.
1.2. Why the associated cominuscule subvariety matters. Flag varieties have few regular maps between them [2, 36, 46, 47, 55, 56, 57]. So they have few flag subvarieties and these subvarieties are surprising. Cominuscule varieties are simpler than other flag varieties in many ways; we hope that these cominuscule subvarieties will shine light on their ambient flag varieties.
We give some evidence for their significance in section 10 on page 39. We will see that the associated cominuscule subvariety of an irreducible flag variety pX, Gq satisfies an open condition on its tangent spaces, which we call freedom. We will see that the associated cominuscule subvariety of any irreducible flag variety is the unique submanifold of maximal symmetry group among all free subvarieties. We will see that all free smooth subvarieties are homogeneous. This supports the conjecture that the associated cominuscule subvariety is the unique free smooth subvariety.
Example 3. As in the image of Aten’s rays, pick a point p0 and a line ℓ0 in the projective plane P2 , with p0 not lying on ℓ0 .

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

5

Each point p of ℓ0 has an associated pointed line: the pair pp, pp0 q.

These pointed lines form a rational curve in the variety of pointed lines (not in P2 ).

This rational curve is homogeneous under the projective transformations fixing p0
and ℓ0 ; it is the associated cominuscule subvariety of the variety of pointed lines.
Each Cartan subgroup of the projective transformations of the plane consists of the transformations which preserve three points in general position.

Hence the associated cominuscule subvariety is invariant under a Cartan subgroup.
Conversely there are three associated cominuscule subvarieties invariant under any given Cartan subgroup. The projective transformations preserving the point p0 and line ℓ0 act transitively on the associated cominuscule subvariety, moving the points p of the line ℓ0 .
Let us apply the algorithm of theorem 1 on page 3. The Dynkin diagram (defined in §4.2 on page 10) of the variety of pointed lines is
. The extended Dynkin diagram is
.
Cut out the crosses and their edges, leaving only the affine node . Turn the affine node into a cross , giving us . This is the Dynkin diagram of the projective line, our rational curve.
Example 4. Take a vector space V and write it as the direct sum of linear subspaces
Vi Ď V , say of dimension ni , i “ 1, 2, . . . , k. Let G :“ SLV . Let P Ă G be the subgroup of linear transformations preserving the successive sums
V1 , V1 ‘ V2 , . . . , V1 ‘ ¨ ¨ ¨ ‘ Vk “ V.
So X :“ G{P is the set of partial flags of dimensions
0, n1 , n1 ` n2 , . . . , n1 ` ¨ ¨ ¨ ` nk “ n.
Let Ğ Ă G be the subgroup preserving V1 ‘ Vk and acting as the identity on every Vi , i “ 2, . . . , k ´ 1. Let P̆ Ď Ğ be the subgroup preserving V1 . Then every element of P̆
preserves V1 , V1 ‘ V2 , . . .. So P̆ Ď P . So X̆ “ Ğ{P̆ Ď X “ G{P is the Grassmannian inside the partial flag variety X. The points of X̆ are precisely the partial flags
0 “ W0 Ă W1 Ă ¨ ¨ ¨ Ă Wk “ V

6

BENJAMIN McKAY

obtained by taking any linear subspace W1 Ď V1 ‘ Vk of the same dimension as V1 , and then taking
W2 :“ W1 ‘ V2 , W3 :“ W1 ‘ V2 ‘ V3 , . . . , Wk :“ W1 ‘ V2 ‘ ¨ ¨ ¨ ‘ Vk .
In other words, pV1 ‘ Vk q X W1 “ W1 , V2 X W1 “ 0, V3 X W1 “ 0, . . . , Vk´1 X W1 “ 0.
Moreover
V2 Ď W2 , V2 ‘ V3 Ď W3 , . . . , V2 ‘ ¨ ¨ ¨ ‘ Vk´1 Ď Wk´1 .
We can phrase this as dim W1 “ dimppV1 ‘ Vk q X W1 q, 0 “ dimpV2 X W1 q “ ¨ ¨ ¨ “ dimpVk´1 X W1 q, and dimpV2 X W2 q ě n2 , . . . , dimppV2 ‘ ¨ ¨ ¨ ‘ Vk´1 q X Wk´1 q ě n1 ` ¨ ¨ ¨ ` nk´1 .
In particular, X̆ Ď X is an intersection of Schubert cells.
2. Review of flag varieties
2.1. Flag varieties. A flag variety pX, Gq is a complex projective variety X acted on transitively and holomorphically by a connected complex semisimple Lie group G
[24] pp. 134–135. Some authors prefer the term generalized flag variety or rational homogeneous variety, We will need to make use of ineffective flag varieties, i.e. G
might not act faithfully on X. It is traditional to denote the stabilizer Gx0 of a point x0 P X as P ; the subgroup P Ď G is a connected complex linear algebraic subgroup. A subgroup of G is parabolic if it is the stabilizer of a point of a flag variety pX, Gq of G.
2.2. Reducible flag varieties. The center of any complex semisimple Lie group
G lies in every maximal torus, so in every Cartan subgroup [7] p. 220, so in every parabolic subgroup, so acts trivially on every flag variety. An irreducible flag variety is a flag variety pX, Gq with G a simple Lie group, and with only the center of G
acting trivially. The product of flag varieties pX1 , G1 q, . . . , pXs , Gs q is
X “ X1 ˆ X2 ˆ ¨ ¨ ¨ ˆ Xs ,
G “ G1 ˆ G2 ˆ ¨ ¨ ¨ ˆ Gs .
Every flag variety pX, Gq, up to finite central extension of G, admits a factorization into a product of (i) irreducible flag varieties and (ii) at most one factor pX0 , G0 q consisting of a point X0 “ t x0 u. This factorisation is unique up to permutation of the pXi , Gi q for i ą 0 and isomorphism. The flag variety pX, Gq is effective if and only if all its factors are effective, i.e. if and only if G0 “ t 1 u is trivial and
G1 , . . . , Gs are in adjoint form, and then pX, Gq is their product (not just up to finite central extension) [1] p. 74.
2.3. Cominuscule varieties. A flag variety is cominuscule if g{p “ Tx0 X is a sum of irreducible complex algebraic P -modules. This occurs just when there is a compact subgroup K Ď G so that pX, Kq is a compact Hermitian symmetric space [35] p. 379 Proposition 8.2, [3] p. 26. Some authors prefer the term compact
Hermitian symmetric space, cominuscule Grassmannian, or generalized Grassmannian to cominuscule variety. Every effective cominuscule variety is a product of the following irreducible effective cominuscule varieties [3] p. 26, Example 3.1.10, [34]
theorem 1 p. 401:

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

G
Ar

dim

G{P

k

kpr ` 1 ´ kq

7

description
Grassmannian of k-planes in Cr`1

Br

2r ´ 1

quadric hypersurface in P2r

Cr

rpr`1q
2

space of Lagrangian r-planes in C2r

Dr

2r ´ 2

quadric hypersurface in P2r´1

Dr

rpr´1q
2

component of space of null r-planes in C2r

Dr

rpr´1q
2

component of space of null r-planes in C2r

E6

16

complexified octave projective plane

E7

27

space of null octave 3-planes in octave 6-space

2.4. Structure of linear algebraic groups. A complex linear map is unipotent if its only eigenvalue is 1. A subgroup of a linear algebraic group is unipotent if it consists of unipotent linear maps. Every complex linear algebraic group G has a unipotent radical, the unique maximal unipotent normal subgroup, which is a closed complex linear algebraic subgroup [7] p. 85 Theorem 4.5, p. 86 Theorem 4.7, p. 157
11.21.
A complex linear algebraic group is reductive if it contains a Zariski dense compact subgroup [61] p. 142 Theorem 2.7, Theorem 2.8. A reductive Levi factor of a complex linear algebraic group is a maximally reductive complex linear algebraic subgroup.
Every complex linear algebraic group G has a Levi factor, unique up to conjugacy.
The group G is a semidirect product of its reductive Levi factor and its unipotent radical [7] p. 158 11.22. Any Cartan subgroup is therefore a subgroup of the reductive
Levi factor, after perhaps a conjugacy. Every compact subgroup lies in a maximal compact subgroup, which is unique up to conjugacy, so lies in the reductive Levi factor up to conjugacy [22] p. 531 Theorem 14.1.3. After perhaps extending by some finite group of order a power of 2, the Weyl group embeds in G as a finite subgroup
[60]. So this group lies in the reductive Levi factor up to conjugacy.
2.5. Parabolic subgroups. A Zariski closed subgroup P Ď G of a connected linear algebraic group G is parabolic if X :“ G{P is a projective variety, and this occurs just when pX, Lq is a flag variety for a semisimple Levi factor L Ď G, and this occurs just when P contains a Borel subgroup (i.e. a maximal connected solvable subgroup)
[7] p. 148. Every parabolic subgroup is connected [7] p. 197 Proposition 14.18. The unipotent radical of P is denoted G` Ď P , and a reductive Levi factor is denoted
G0 Ď P , so P “ G0 ˙ G` . (This notation is potentially confusing. The reader might expect to write these as P` and P0 since they lie in P . But our notation is standard
[16] p. 293 theorem 3.2.1. It is due to the presence of the grading of the Lie algebra of G which we will define.) A flag variety is cominuscule just when G` is abelian
[16] p. 296 §3.2.3. Denote the center of the unipotent radical by Z :“ ZG` .
2.6. Opposite parabolic subgroups. Two parabolic subgroups P, P op Ď G of a complex semisimple Lie group are opposite if P X P op is a reductive Levi factor of both P and P op . Every parabolic subgroup P has an opposite, unique up to conjugation by an element of G` [7] p. 198 proposition 14.21.

8

BENJAMIN McKAY

3. Definition of the associated cominuscule subvariety
3.1. The definition. Take a flag variety pX, Gq and opposite parabolic subgroups
P, P op Ď G, so that P is the stabilizer of x0 P X. As above, take their unipotent op radicals G` , Gop of these. Let Ğ :“ ⟨Z, Z op ⟩ Ď G be the
` and the centers Z, Z
op subgroup generated by Z Y Z , P̆ :“ Ğ X P , X̆ :“ Ğ{P̆ . Then pX̆, Ğq is the associated cominuscule subvariety passing through the point x0 P X.
Lemma 1. The associated cominuscule subvariety of a flag variety pX, Gq is uniquely determined up to G-conjugacy.
Proof. Since G acts transitively on X, it suffices to prove the conjugacy of any two associated cominuscule flag varieties through the same point. Take a point x0 P X
and consider two associated cominuscule subvarieties passing through x0 . Thus they are stabilized by the same parabolic subgroup P :“ Gx0 . So it suffices to prove that any two associated cominuscule subvarieties passing through x0 are P -conjugate. By definition, each associated cominuscule subvariety with a given parabolic subgroup P
is determined by the group generated by Z, Z op . But Z is the center of the unipotent radical of P , so determined by P . So each associated cominuscule subvariety through the given point x0 is determined by the choice of Z op , the center of the unipotent radical of the choice of opposite parabolic subgroup P op , so determined by P op .
But the opposite parabolic subgroup is uniquely determined up to G` -conjugacy
[7] p. 198 proposition 14.21, hence up to P -conjugacy.
□
The tangent space to X̆ at x0 is zop . Intuitively, the associated cominuscule variety arises by integrating G-translates of that subspace. Those G-translates are linear subspaces transverse to the largest P -invariant holomorphic distribution on the flag variety pX, Gq. Thus X̆ is an integral manifold of those G-translates; it is not a priori obvious that the G-translates have integral manifolds. So our interest in associated cominuscule subvarieties fits into a larger story about tangent geometry of flag varieties [26, 27, 32].
3.2. Example: the general linear flag variety. We return to the study of the flag varieties of An´1 “ PSLn ; see example 4 on page 5. We took a vector space V
of dimension n. We let G :“ PSLV and X the set of partial flags of dimensions
0, n1 , n1 ` n2 , . . . , n1 ` ¨ ¨ ¨ ` nk “ n.
We wrote V as the direct sum of linear subspaces Vi Ď V , say of dimension ni , i “ 1, 2, . . . , k. Supposing that V “ Cn , our automorphism becomes transpose inverse. For concreteness take k “ 4. Let G1 Ď G be the subgroup preserving
X̆ Ď X and set P 1 :“ G1 X P . In our table, each line is a pair Γ M , of a group
Γ and a matrix M with some unspecified entries. In each pair, Γ is the group of unimodular matrices of the given form M , modulo rescaling by the matrices λI of that form where λ is an nth root of unity.
¨
˛
˚ ˚ ˚ ˚
˚˚ ˚ ˚ ˚‹
‹
G ˚
˝˚ ˚ ˚ ˚‚
˚ ˚ ˚ ˚
˛
˚ ˚ ˚ ˚
˚0 ˚ ˚ ˚‹
˚
‹
˝0 0 ˚ ˚‚
0 0 0 ˚
¨
P

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

G`

Z

¨

I
˚0
˚
˝0
0

˚
I
0
0

˚
˚
I
0

˛
˚
˚‹
‹
˚‚
I

¨

0
I
0
0

0
0
I
0

˛
˚
0‹
‹
0‚
I

I
˚0
˚
˝0
0

9

¨
G0

Ğ

P̆

˛
˚ 0 0 0
˚0 ˚ 0 0‹
˚
‹
˝0 0 ˚ 0‚
0 0 0 ˚
¨

˚
˚0
˚
˝0
˚

0
I
0
0

0
0
I
0

˛
˚
0‹
‹
0‚
˚

¨

0
I
0
0

0
0
I
0

˛
˚
0‹
‹
0‚
˚

˚
˚0
˚
˝0
0
¨

G1

˛
˚ 0 0 ˚
˚0 ˚ 0 0‹
˚
‹
˝0 0 ˚ 0‚
˚ 0 0 ˚
¨

P1

˛
˚ 0 0 ˚
˚0 ˚ 0 0‹
˚
‹
˝0 0 ˚ 0‚
0 0 0 ˚

To be more precise, G` consists of the matrices
¨
λI
˚0
˚
˝0
0

˚
λI
0
0

˚
˚
λI
0

˛
˚
˚‹
‹
˚‚
λI

with determinant 1, up to rescaling by nth roots of unity. But for each such matrix equivalence class, pick any representative and we can pick that root of unity uniquely to write it as an actual matrix with λ “ 1, and similarly for Z. Clearly Ğ has Lie algebra containing all root vectors of all P -maximal and P -minimal roots, and is generated by the 1-parameter subgroups of those root vectors. So it contains the root vectors of the root system generated by these, and the Cartan subgroup of that root system, hence a semisimple group. Note that P 1 preserves V1 , V2 , V3 , and
V1 ‘ V4 , whereas G1 preserves V2 , V3 and V1 ‘ V4 . Hence X̆ “ Ğ{P̆ “ G1 {P 1 is the
Grassmannian of linear subspaces of dimension n1 inside V1 ‘ V4 .

10

BENJAMIN McKAY

4. Dynkin diagrams
4.1. Compact roots. Denote the Lie algebras of P Ď G by p Ď g. One can select a Cartan subgroup of G lying inside P , whose positive root spaces all lie in p. A
simple root α is P -compact (compact if P is understood) if the root space of ´α
belongs to the Lie algebra of P . This occurs just when the sl2 generated by the root vectors of α, ´α lies inside the Lie algebra p of P , hence inside a semisimple Levi factor. So we can assume that all P -compact roots are roots of a semisimple Levi factor lying inside G0 .
4.2. The Dynkin diagram of a flag variety. The Dynkin diagram of a flag variety pX, Gq is the Dynkin diagram of G decorated with on each compact simple root and on each noncompact simple root [7] p. 197 Proposition 14.18, [24] p. 197
Theorem1 . Each flag variety is determined uniquely, up to finite central extension of
G and up to isomorphism, by its Dynkin diagram.
4.3. Affine extensions. Every Dynkin diagram sits inside its extended Dynkin diagram [9] p. 198, [15] p. 262, [31] lecture 17, [48] p. 164, 5o . The extended Dynkin diagram has one node added to each component of the original Dynkin diagram. We call this node the affine node, and draw it as a hollow dot , with edges attached as:

G

Dynkin

extended

A1
Arě2
B2 “ C2
Brě3
Crě2
Drě4

E6
E7
E8
F4
G2

Recall that we add this node to the original Dynkin diagram to represent the negative α0 :“ ´θ of the highest root θ. We then compute the Cartan matrix as usual. Consider how we draw edges connecting that new node to each old node according to the usual Dynkin diagram rules. For any finite or extended root system, denote the Cartan matrix entry in row i column j as nij . If i ‰ j and |nij | ě |nji |, we draw |nij | edges from vertex j to vertex i. If |nij | ą 1 we draw an arrow pointing to node i.

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

11

4.4. Rank 1. For A1 , the Cartan matrix of the extended root system is
ˆ
˙
2 ´2
´2 2
so both off-diagonal Cartan matrix entries are ´2. Hence we draw the extended
Dynkin diagram of A1 as with a double edge with arrows pointing in both directions [30] p. 44 table Aff 1. The total number of edges of the extended Ar
Dynkin diagram is r ` 1, for any integer r ě 1.
There is a different rule [48] p. 164 for forming Dynkin diagrams: we could instead make the edge between node i and node j have multiplicity nij nji , and make an arrow from node j to node i just when |nij | ą |nji |. This rule leads to exactly the same Dynkin diagrams and extended Dynkin diagrams, except precisely for the extended Dynkin diagram of A1 , yielding a quadruple edge [61] p. 228 table 3, [48]
p. 166
.
The choice of convention for drawing the extended Dynkin diagram of A1 does not affect our algorithm for computing the associated cominuscule subvariety: either convention yields the same result.
4.5. The extended Dynkin diagram of a flag variety. Take the Dynkin diagram of a flag variety pX, Gq. The extended Dynkin diagram of pX, Gq is the associated extended Dynkin diagram, with nodes drawn as or as for the flag variety, keeping the affine node drawn as . So the flag variety has extended
Dynkin diagram

5. Reducing to root systems
5.1. Gradings.
A root system
ř
ř with a basis of simple roots α1 , . . . , αr is graded: each root ni αi has grade i ni . For a flag variety X “ G{P , the root system is
ř
also P -graded by ni , but summing only over the noncompact (crossed) simple roots [16] p. 292. The P -grade is also called the P -height [16] p. 292. A root is
P -maximal, also called a box root, if it has maximal P -grade in its irreducible factor.
The box is the set of box roots, terminology which roughly follows what [14], [37]
p. 57 might perhaps call the maximal box, by analogy with Young tableaux. We will see in lemma 5 on page 18 that the box generates the root system of Ğ.
5.2. The box and the Lie algebra. The unipotent radical G` Ď P has Lie algebra g` Ď p the sum of the root vectors of the positively graded roots [7] p. 197
Proposition 14.18, [33] p. 482. The zero graded roots are invariant under reflection in one another. The sum of the Cartan subalgebra with the sum of the root spaces of the zero graded roots is the semisimple factor of the reductive Levi factor g0 of p [7] p. 197 Proposition 14.18. Denote by z the Lie algebra of Z. For any root α, denote by gα the α-root space of g. We will prove:
Lemma 2. The abelian Lie algebra z is the span of the root vectors of the roots of the box:
à
z“
gα .
αPbox

Take the box roots (i.e. roots of the box) as vertices of a graph. If the difference α ´ β
of two box roots is a P -compact simple root, corresponding to node ℓ of the Dynkin diagram of pX, Gq, draw an edge from β to α labelled ℓ. Thus the box becomes a graph. As a graph, the connected components of the box are precisely the boxes of the irreducible factors of the flag variety. Each component contains the highest root of the factor. In particular, the box of an irreducible flag variety is connected and contains the highest root.

12

BENJAMIN McKAY

5.3. Associated cominuscule subvarieties in rank 2.
Example 5. Here we will explain how to read our drawings. We pick a basis of simple roots in each root system. We draw the simple roots as empty circles:

Start with the roots of G2 :

Each parabolic subgroup has Lie algebra consisting of the sum of the Cartan subalgebra and the root spaces of certain roots. To be precise, these roots are the ones which lie on or on one side of a hyperplane. The compact roots are those on the hyperplane. Conversely, draw any hyperplane and it produces a parabolic subgroup. For example, here is the hyperplane of some parabolic subgroup.

Drawing both the hyperplane and roots together:

We can always pick a basis of simple roots so that every simple root lies on the hyperplane, or on the chosen side of the hyperplane. It is compact (so uncrossed)
just when it is on the hyperplane:

(When working with G2 , we will always pick our hyperplane to allow for the chosen bases of simple roots shown above.) We pick the hyperplane to be the zero locus of the real linear function vanishing on uncrossed simple roots and taking value one on crossed simple roots [16] p. 239 Proposition 3.1.2. Grade the roots by P -height, i.e. by sum of coefficients of crossed simple roots. Our hyperplane is thus always chosen so that we can see the heights, i.e. roots of a given height lie on a parallel hyperplane:

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

13

By definition of Ğ, its root system is the root system generated by the box, i.e. by the maximal graded roots:

So the Ğ-roots form the smallest root subsystem containing these. Finally, draw dark lines through the Ğ-roots:

The roots of P̆ “ P X Ğ are the roots of Ğ lying on or to the indicated side of the hyperplane: 4 of them in our picture:

Hence X̆ “ Ğ{P̆ has dimension equal to the number of Ğ-roots not lying in P̆ , i.e.
the dimension of X̆ is the number of roots in the box. We can see in the picture that the root system of Ğ in this example is that of A2 “ PSL3 , and that X̆ has dimension 2, so must be pX̆, Ğq “ pP2 , A2 q. The elements of G preserving X̆ form a subgroup of G which maps to Ğ, which we call the automorphism group. We will see that the automorphism group is generated by the flow through 1 P G of various root vectors of G, associated to various roots in the root system of G. Among these roots are the roots of Ğ. Below we will colour these in; for our example, the roots of the automorphism group are precisely those of Ğ.
Example 6. We draw the gradings of the positive roots of the parabolic subgroups of the rank 2 simple groups. Under the heading ğ, we draw the roots of the symmetry
Lie algebra of the associated cominuscule subvariety, and under the heading g1 the roots of the automorphism Lie algebra.
Grading

A2

A2

B2
continued. . .

ğ

g1

14

BENJAMIN McKAY

continued. . .
Grading

ğ

g1

B2

B2

G2

G2

G2

Example 7. Any maximal irreducible flag variety X “ G{B has associated cominuscule subvariety “ pP1 , PSL2 q, reducing maximally. This occurs because G has a unique highest root, whose root space generates Z.
Example 8. Similarly, the associated cominuscule subvariety of any adjoint variety is “ pP1 , PSL2 q. The adjoint varieties are precisely those with invariant contact structures. An invariant contact structure arises from a hyperplane in each tangent space, invariant under the parabolic. By invariance, it is a sum of root spaces. So a single root, negative for the parabolic, doesn’t have its root vector in this hyperplane, i.e. a single root is at higher weight. So the associated cominuscule subvariety is complementary to the hyperplane, one dimensional, hence a projective line.
6. The associated cominuscule subvariety is cominuscule
Henceforth suppose that G is a connected complex semisimple Lie group. Take a factorization
X “ X0 ˆ X1 ˆ X2 ˆ ¨ ¨ ¨ ˆ Xs ,
G “ G0 ˆ G1 ˆ G2 ˆ ¨ ¨ ¨ ˆ Gs , into irreducible flag varieties pXi , Gi q, i ą 0, and a point X0 “ t x0 u. Then
P “ G0 ˆ P1 ˆ P2 ˆ ¨ ¨ ¨ ˆ Ps , with Pi :“ P X Gi . The unipotent radical of P is obviously the product of the unipotent radicals of the Pi :
G` “ t 1 u ˆ G1` ˆ ¨ ¨ ¨ ˆ Gs` ,

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

15

where Gi` is the unipotent radical of Pi . Therefore
Z “ t 1 u ˆ Z1 ˆ ¨ ¨ ¨ ˆ Zs where Zi is the center of Gi` . We prove lemma 2 on page 11.
Proof. Let z1 be the direct sum of the root spaces of the box roots, i.e. the P maximal roots. Recall that z is the center of the nilpotent radical g` of p. We have to prove that z1 “ z. Note that in our decomposition above, z1 “ t 0 u ‘ z11 ‘ ¨ ¨ ¨ ‘ z1s .
So it suffices to prove the result for an irreducible flag variety pX, Gq. By invariance under the Cartan subalgebra, z is a sum of root spaces. So we let B 1 be the box and B the set of roots whose roots spaces belong to z. To prove z “ z1 is precisely to prove that B “ B 1 .
Recall that the bracket of root vectors of roots α, β is either zero or is a root vector of the root α ` β. If α P B 1 , i.e. α is P -maximal, its associated root vector brackets to zero with the root vector of any root β of positive P -grade. Hence the root vector lies inside z, i.e. z1 Ď z, i.e. B 1 Ď B.
Recall that every root system has a unique highest root, i.e. the highest weight of the adjoint representation [53] p. 61 Theorem 3. Picture a sequence of steps from root to root, by adding only simple roots. Every root can be brought to the highest root by such a sequence [20] pp. 330–331 §21.3, [25] p. 56, [23], [53] p. 46
2.12, [53] p. 58 Theorem 3. Adding P -compact simple roots preserves P -height.
Adding P -noncompact simple roots increases P -height. When this sequence reaches a P -maximal root, it can no longer increase the P -height. So we can only add
P -compact roots, moving through P -maximal roots, so moving along B 1 . So the highest root is P -maximal, so in B 1 , and is obtained from any root of B 1 by adding only P -compact simple roots. Every B 1 -root can thus be brought to the highest root by adding only P -compact simple roots. So the box becomes a connected graph, with the highest root as one vertex. The vector space z1 is therefore an irreducible
G0 -module [33] p. 332 proposition 5.105, and a G0 -submodule of z.
Pick a root α P B. It suffices to prove that α P B 1 , i.e. P -maximal. Pick a root vector eα P z for root α. Suppose that α ` β is a root, for some root β of positive
P -height. Then 0 ‰ reα , eβ s P gα`β for some root vector eβ of β [53] p. 46 2.12.
But then eα is not in z, i.e. does not centralize g` . But the highest root is α or is of the form α ` β for some positive root β. Hence β is P -compact, i.e. α is P -maximal.
So α is a box root: α P B 1 .
Similarly, if α ` β is a root, for some root β of zero P -height, i.e. a compact root, then 0 ‰ reα , eβ s P gα`β , so α ` β also belongs to the box, and its root space gα`β
also belongs to z. So B is invariant under stepping through roots by compact roots.
So B “ B 1 and z “ z1 .
□
Lemma 3. The group Ğ is connected.
Proof. Every parabolic subgroup of a connected reductive Lie group is connected, so G and P are connected [7] p. 154 theorem 11.16. By Langlands decomposition
[33] p. 482, G, P, X, G` , G0 are connected. The subgroup Z “ ZG` is a Zariski closed subgroup of a unipotent linear algebraic group, so connected and isomorphic as an affine variety to complex Euclidean space [17] p. 36 corollary 15.1.11, [19]
p. 499, Corollaire 4. The exponential map on the Lie algebra of a unipotent linear algebraic group is a biholomorphism and a regular algebraic morphism mapping subalgebras to subgroups, so taking z Ñ Z, an isomorphism of schemes [43] p. 289
Proposition 14.32. By definition of Ğ, Ğ :“ ⟨Z, Z op ⟩ is a path connected complex
Lie group, since Z is path connected. Moreover, Ğ is connected in the algebraic

16

BENJAMIN McKAY

Zariski topology by Chevalley’s theorem on constructible sets [17] p. 38 Proposition
16.2.1.
□
6.1. Roots and opposite parabolics. All Borel subgroups of G are conjugate [7]
p. 147 chapter IV 11.1, each containing a Cartan subgroup, hence the Lie algebra of P is the sum of (1) the Cartan subalgebra with (2) various root spaces. Every automorphism of a root system arises from an automorphism of the associated semisimple Lie group [20] p. 498 Proposition D.40, [25], p. 87, §16.5. Hence a
there is an automorphism G Ý
Ñ G of G which yields α ÞÑ ´α in the root system.
(We can define such an automorphism explicitly as eα ÞÑ ´e´α on root vectors in a Chevalley basis; see §10.5 on page 41.) Our automorphism sends P to an opposite parabolic subgroup P op :“ aP with P X P op “ G0 . Letting G´ :“ aG` ,
G` X G´ “ t 1 u. An open subset of G consists of elements uniquely expressed as a product p, q P P ˆ G´ ÞÑ pq P G [16] p. 294, [7] p. 198 Proposition 14.21. Every root system also has an automorphism, traditionally called w0 , which belongs to the
Weyl group and which interchanges the positive and negative roots of a root system
[16] pp. 323–324; it is the unique element of the Weyl group of maximum length.
Note that w0 might not reverse the signs of simple roots [16] p. 324. The Weyl group lifts to a group of automorphisms of the Lie group G, after perhaps extension by some finite group of order a power of 2. We can use such an extension of w0 in place of a throughout this paper, as we will only need that a is an automorphism of a given root system which extends to an automorphism of G taking a given parabolic subgroup to an opposite.
6.2. Proof of the cominuscule property. A Lie group G acts almost effectively on a manifold X if the elements of G fixing every point of X form a finite subgroup.
Lemma 4. The complex homogeneous space pX̆, Ğq is a positive dimensional homogeneously embedded cominuscule subvariety of pX, Gq. If pX, Gq is almost effective then so is pX̆, Ğq. The Dynkin diagram of pX̆, Ğq has one connected component for each connected component of the Dynkin diagram of pX, Gq.
Proof. Take notation as above for a flag variety X “ G{P . It suffices to assume that pX, Gq is effective. It also suffices to assume that pX, Gq is an irreducible flag variety, as otherwise it is a product of irreducibles and the subgroup Ğ is the product of the associated subgroups. The Lie algebra z of the center Z of the unipotent radical G`
of P is the sum of the root spaces of the box, i.e. of the P -maximal roots. Let a be an automorphism of G which changes the sign of the P -grading of the roots; see
§2.6 on page 7 where we constructed one such automorphism. So az is the sum of the root spaces of the roots of minimal P -grade, the opposite box. Under bracket, these root spaces generate only root vectors and coroots, up to scaling. Hence the
Lie algebra ğ of Ğ is the sum of some such. A coroot arises only when we bracket the root vectors of the associated root and its negative. So ğ contains all of the root vectors of the root system generated by the box. It also contains their brackets.
Hence it contains the complex semisimple subalgebra with that root system. That subalgebra is generated in the same way, by the box root vectors. Every complex semisimple Lie group is generated from the root vectors of a basis of roots [53]
p. 43 Theorem 1. So ğ is that subalgebra, so is complex semisimple. By lemma 3
on the previous page, Ğ is a complex semisimple Lie group. Since its root system lies inside that of G, its Cartan subgroup is a subgroup of the Cartan subgroup of
G, generated by its coroots. Note that P contains the Cartan subgroup of G, hence that of Ğ, and contains Z, so contains the parabolic subgroup of Ğ generated by the box. So this parabolic subgroup fixes the same point x0 P X so lies in P̆ . The

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

17

Ğ-orbit X̆ of that point x0 is a flag variety pX̆, Ğq, with stabilizer P̆ “ P X Ğ a parabolic subgroup, so connected.
Since ğ is a complex subgroup of g, Ğ is a complex subgroup of G, and so X̆ Ď X
is a compact complex submanifold, and X is a smooth projective variety, so X̆ Ď X
is a smooth subvariety.
The vector spaces z and az are irreducible G0 -modules [33] p. 332 proposition
5.105. Hence ğ is a G0 -module. As G0 is reductive, ğ is a direct sum of irreducible
G0 -modules: ğ “ z ‘ ğ0 ‘ az.
Suppose that α is a P -compact root, i.e. a root of G0 . Reflection in α is carried out by an element of the Weyl group of G0 , and so preserves the P -grading. So if
β is any P -maximal root, then so is its α-reflection. In other words, reflection in
α preserves the box. Recall that an α-root string is a directed path of roots, each given by adding α to the previous. Reflection in α moves β along an α-root string.
If that root string has more than one root in it, then α is a difference of P -maximal roots. Their root vectors lie in p̆. So the difference of their root vectors lies in ğ, since it is semisimple. So α is a Ğ-root, and a difference of P̆ -maximal roots.
Suppose that every such root string has a single root in it. So it is a root string of length 1. Reflection in α fixes all roots in the box, i.e. all P -maximal roots. Hence
˘ to be the root system generated by the box.
α is perpendicular to the box. Take ∆
˘ Hence α is perpendicular to every
Reflection in α therefore fixes every root in ∆.
˘
root in ∆.
Any two root subsystems which are mutually perpendicular arise from factors of
˘
the associated semisimple Lie group. Hence the P -compact roots divide into (1) ∆
˘
˘
and (2) those perpendicular to ∆. The roots of ∆ are those which are differences
˘ form a of P -maximal roots, i.e. P̆ -compact roots. The roots perpendicular to ∆
root subsystem of the P -compact roots giving a complex semisimple subgroup of
G0 . This subgroup acts trivially on ğ because it is generated by root vectors in a
˘ is graded into the P -minimal perpendicular root subsystem. The root system ∆
roots (grade ´1), differences of P -maximal roots (grade 0) and P -maximal roots
˘
(grade 1). The Lie algebra ğ consists of the sum of the root vectors of the ∆-spaces, and their coroots (grade 0). The subalgebra p̆ :“ p X ğ consists precisely of the 0
and 1 grades. Note that p̆ acts on z as ğ0 :“ g0 X ğ, i.e. as a sum of irreducible p̆-modules, so X̆ “ Ğ{P̆ is cominuscule.
Recall that z is an irreducible G0 -module. If we start at the highest root, we can pass from it via root strings to get to any P -maximal root, i.e. anywhere in the box.
As we do this, we repeatedly pass between P -maximal roots α, β by subtracting a P̆ -compact positive root α ´ β. Bracketing a root vector eα´β of root α ´ β
takes eα to a nonzero multiple of eβ . Hence z is an irreducible ğ0 -module, hence an irreducible p̆-module. As we have seen on page 6, the number of irreducible modules of the parabolic subgroup is the number of irreducible factors of its flag variety. Hence X̆ is an irreducible flag variety.
We next prove that pX̆, Ğq is almost effective. An element g P Ğ acts trivially on X̆ just when gg1 P̆ “ g1 P̆ for all g1 P Ğ, i.e. just when g lies in all Ğ-conjugates of P̆ . Since Ğ is complex semisimple, it admits an automorphism ă interchanging
P̆ and P̆ op . This automorphism can be arranged to be conjugation by an element op of Ğ [60]. So P̆ op is a conjugate of P̆ . But Ğ` X Ğop
` “ 1 and Ğ0 “ Ğ0 . So an element g P Ğ acting trivially on X̆ lies in Ğ0 , the maximal reductive subgroup of
P̆ . Acting trivially on Tx0 X̆, g acts trivially on z. Reversing, it acts trivially on zop , so acts trivially on ğ, so lies in the center of Ğ: g P ZĞ . The center of a complex semisimple Lie group is finite [53] p. 66, so Ğ is almost effective.
□

18

BENJAMIN McKAY

We restate part of the proof of the previous result, as a separate lemma:
Lemma 5. For any flag variety pX, Gq, the box generates the Ğ-root system. The
P -compact roots divide into (1) the P̆ -compact roots, which are precisely those roots which are differences of roots of the box and (2) those perpendicular to the Ğroots, forming a root subsystem of the P -compact roots giving a complex semisimple subgroup of G0 acting trivially on ğ.
Lemma 6. The associated cominuscule subvariety pX̆, Ğq of a flag variety pX, Gq has the same box as pX, Gq.
Proof. As above, we can assume that pX, Gq is an irreducible flag variety. The definition of associated cominuscule subvariety pX̆, Ğq of a flag variety pX, Gq has
Ğ the group generated by Z, Z op , so Ğ has Lie algebra containing the Lie algebra z of Z and P̆ is defined to be Ğ X P , so the Lie algebra of P̆ contains z. The roots of
G whose root spaces sum up to z are precisely the box roots by lemma 2 on page 11.
By lemma 4 on page 16, P̆ Ă Ğ is a parabolic subgroup. By lemma 5, pX̆, Ğq has root system, marked into compact and noncompact, a subsystem of that of pX, Gq.
Moreover, as a graded Lie algebra, ğ “ z ‘ ğ0 ‘ zop is graded into ´1, 0, 1 grades.
So ğ` “ z is abelian, so its own center: z̆ “ ğ` “ z.
□
6.3. Automorphisms of the associated cominuscule subvariety.
Example 9. The flag variety pX, Gq “ pB2 {P, B2 q, where P Ď B2 is the Borel subgroup:

has associated cominuscule subvariety

This is a rational curve pX̆, Ğq “ pP1 , PSL2 q. The rational curve invariant under not only its automorphism group Ğ “ A1 “ PSL2 , which we see in our diagram, but also, as we will see, under rescaling by this root:

The root is perpendicular to the Ğ root system, so commutes with the Ğ root vectors. The vector field on X associated to the root vector of that root vanishes at our chosen point x0 P X stabilized by P , since the root lies in the Lie algebra of P . Since the vector field commutes with those of Ğ, it is Ğ-invariant. Since Ğ
acts transitively on X̆, our vector field vanishes at all points of X̆. Hence, in this example, the automorphism group G1 of X̆ as a subvariety of X is larger than the automorphism group Ğ of X̆ as a flag variety.
Conversely, consider the root

It doesn’t belong to the Lie algebra of P , so doesn’t vanish at x0 . Commuting with the root vectors of Ğ, it is Ğ-invariant, so it doesn’t vanish at any point of X̆. If it were tangent to X̆ at some point of X̆, Ğ-invariance would force it to be tangent at

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

19

every point. But every holomorphic vector field on the projective line X̆ “ P1 has a zero. Indeed, every holomorphic vector field on any flag variety has a zero, since every automorphism has a fixed point [59]. Hence this root vector is a vector field on X, nowhere tangent to X̆.
6.4. Computing the automorphism Lie algebra. We return our thoughts to the general flag variety pX, Gq. By definition of pX̆, Ğq, X̆ Ď X is a smooth subvariety.
So the subgroup G1 Ď G leaving X̆ Ď X invariant is a Zariski closed subgroup of G, hence linear algebraic. (As we have noted, we will find that, while G1 contains Ğ, it is often larger than Ğ and is not always semisimple.) Let P 1 :“ G1 X P , so
X̆ “ Ğ{P̆ “ G1 {P 1 .
Clearly X̆, P̆ , Ğ are connected, so G1 is connected just when P 1 is connected.
Theorem 2. The automorphism group G1 of the associated cominuscule subvariety
X̆ Ă X of a flag variety is a complex linear algebraic group with Lie algebra precisely the sum of
‚ the P -maximal and P -minimal root spaces,
‚ the maximal reductive g0 , and
‚ the root spaces of all positive roots perpendicular to all P -maximal roots.
Example 10. In example 9 on the facing page, we saw that the third case above is not redundant: there is a positive root perpendicular to all P -maximal roots in that example, and this root is not a root of the maximal reductive g0 , since in that example, g0 is abelian.
Proof. Clearly G1 Ď G is a complex linear algebraic subgroup containing Ğ, hence acts transitively on X̆. So we only need to prove that the Lie algebra g1 of G1 is the sum of the appropriate root spaces. To see that g1 is a sum of root spaces, we need to prove that it contains the Cartan subalgebra of g.
Since X̆ is a projective variety and G1 is a linear algebraic group acting transitively on X̆, the stabilizer P 1 :“ P X G1 of the point x0 P X is a parabolic subgroup of G1
[7] p. 148, so contains a Borel subgroup of G1 .
Let G10 Ă G1 be the connected component of the identity. Since X̆ is a projective variety and G10 is a connected linear algebraic group acting transitively on X̆, the stabilizer P01 :“ P 1 X G10 of the point x0 P X is a parabolic subgroup of G10 [7] p. 148, hence is connected. Since X̆ “ G1 {P 1 is connected, P 1 intersects every component of G1 .
Since Ğ acts transitively on X̆, every element of G1 is a product of an element of
Ğ and an element of P .
Claim: P 1 is the normalizer of ĞP in P . Proof: by definition p P P 1 if and only if p P P and pX̆ “ X̆. The points of X̆ have the form gx0 for g P Ğ, unique up to multiplying by an element of P , i.e. pgP “ g 1 P , for some g 1 P Ğ. So p P P 1 if and only if p P P and pĞP Ď ĞP . Since P 1 is a group, p P P 1 just when p P P and pĞP “ ĞP , or equivalently, pĞP p´1 “ ĞP.
1
Thus P contains all elements of P which normalize Ğ. It contains P̆ . Hence it contains Z “ ZG` Ď P̆ . The maximal reductive G0 Ď P normalizes P hence G`
hence Z, and is invariant under the automorphism a. Thus G0 Ď P 1 . Hence P 1
contains the Cartan subgroup. So the Lie algebra p1 of P 1 is a sum of the Cartan subalgebra together with various root spaces. Since P 1 lies inside G1 , the Lie algebra of G1 is also a sum of the Cartan subalgebra together with various root spaces. We will say that a root α of G is a G1 -root if its root space lies in g1 , a P 1 -root if its root space lies in p1 .

20

BENJAMIN McKAY

The remainder of the proof consists in identifying the G1 -roots. First, we find some roots which are obviously G1 -roots. Take a root α perpendicular to all roots in the box, i.e. to all P -maximal roots. It is then also perpendicular to their negatives, and so its associated root vector eα brackets to zero with all of the root vectors eβ
of all of those roots, and hence of all roots β of Ğ. As we argued in §6.3 on page 18,
α is P -nonnegative if and only if eα is tangent everywhere to X̆, hence generates an automorphism of X̆. So the Lie algebra g1 of G1 contains the root vector eα of a root α perpendicular to the box roots just when α is P -nonnegative, so either
P -compact or P -positive. If α is P -compact, α is a root of g0 , so we may assume that α is P -positive, hence positive.
So far, we have affirmed that the roots which belong to the box or its negative or to the maximal reductive or are positive and perpendicular to the box are G1 -roots.
We have also affirmed that the roots which are negative and perpendicular to the box have root spaces disjoint from g1 , so are not G1 -roots.
We are left to consider a root α which is not in the box, and ´α is not in the box, and α is not a root of the maximal reductive, and α is not perpendicular to the box. We need precisely to prove that α is not a G1 -root, i.e. the root vector eα
of α is not in g1 .
By definition P 1 preserves ğ ` p. Pick a root β ` in the box, not perpendicular to our P -submaximal P -positive root α. Note that if two roots have more than a right angle between them, then their sum is a root and their root vectors have nonzero
Lie bracket [53] p. 29. The root system generated by α and β ´ :“ ´β ` is of rank
2. Our P -grading grades that rank 2 root system. Our pictures above classify every parabolic subalgebra of every rank 2 simple Lie algebra. In those pictures, we see that the root vectors of α and β ´ have more than a right angle between them and that their sum is a root. Their P -grading is at least that of β ´ , P -negative. Hence the Lie bracket of their root vectors is not in p, but also is not P -minimal, so not in ğ. Hence their root vectors bracket out of ğ ` p. Hence the Lie algebra g1 of G1
does not contain the root space of α. In other words, α is not a G1 -root.
On the other hand, if α is P -negative, the same argument applies with β ´
replaced by β ` . So neither α nor ´α are G1 -roots.
Note that the P 1 -roots are precisely the P -maximal roots and the P -positive roots perpendicular to them. Thus p1 is the sum of the Cartan subalgebra with the roots spaces of these roots.
□
Lemma 7. For any flag variety pX, Gq, the group G1 of automorphisms of the associated cominuscule subvariety X̆ Ď X is connected. The subgroup P 1 Ď G1 fixing a point of X̆ is a connected parabolic subgroup.
Proof. Since Ğ Ď G1 , G1 acts transitively on the connected variety X̆. We saw in the proof of theorem 2 on the preceding page that the stabilizer P 1 Ă G1 of any point x0 P X̆ intersects every component of G1 . Flag varieties are connected and simply connected [6]. By exact sequence in homotopy, inclusion P 1 Ñ G1 is an isomorphism on π0 and π1 . It suffices to prove that P 1 is connected. The subgroup
G1` :“ G` X P 1 Ă G1 is unipotent. It is therefore connected and isomorphic as an affine variety to complex Euclidean space [17] p. 36 corollary 15.1.11, [19] p. 499,
Corollaire 4.
Since P “ G0 ˙ G` , we can write each element p P P as a product of elements of
G0 and G` . Since G0 Ď P 1 , we can write each element p P P 1 as product of elements of G0 and G1` so P 1 “ G0 ˙ G1` so P 1 is connected. In the proof of theorem 2 on the previous page, we saw that P 1 is therefore a parabolic subgroup.
□
Theorem 3. For any flag variety pX, Gq, the group G1 of automorphisms of the associated cominuscule subvariety X̆ Ď X acts on X̆ as precisely the biholomorphisms

COMINUSCULE SUBVARIETIES OF FLAG VARIETIES

21

of X̆ arising from elements of Ğ: the morphisms
Ğ Ñ AutX̆ , G1 Ñ AutX̆
have the same image. The first morphism has kernel a finite central subgroup of the reductive Levi factor Ğ0 Ď P̆ , central in Ğ. The second morphism has kernel the torus generated by the roots perpendicular to the box.
Proof. We can assume that G acts almost effectively on X. The auto

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
