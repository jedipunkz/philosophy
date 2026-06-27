---
source: "https://arxiv.org/abs/1404.3373v3"
title: "Wilder McKay correspondences"
author: "Takehiko Yasuda"
year: "2014"
publication: "arXiv preprint / math.AG"
download: "https://arxiv.org/pdf/1404.3373v3"
pdf: "https://arxiv.org/pdf/1404.3373v3"
captured_at: "2026-06-24T11:11:08Z"
updated_at: "2026-06-24T11:11:08Z"
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

# Wilder McKay correspondences

- 著者: Takehiko Yasuda
- 年: 2014
- 掲載情報: arXiv preprint / math.AG
- 情報源: [arxiv](https://arxiv.org/abs/1404.3373v3)
- ダウンロード: https://arxiv.org/pdf/1404.3373v3
- PDF: https://arxiv.org/pdf/1404.3373v3

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

A conjectural generalization of the McKay correspondence in terms of stringy invariants to arbitrary characteristic, including the wild case, was recently formulated by the author in the case where the given finite group linearly acts on an affine space. In cases of very special groups and representations, the conjecture has been verified and related stringy invariants have been explicitly computed. In this paper, we try to generalize the conjecture and computations to more complicated situations such as non-linear actions on possibly singular spaces and non-permutation representations of non-abelian groups.

## PDF Text

WILDER MCKAY CORRESPONDENCES

arXiv:1404.3373v3 [math.AG] 5 Jun 2014

TAKEHIKO YASUDA
Abstract. A conjectural generalization of the McKay correspondence in terms of stringy invariants to arbitrary characteristic, including the wild case, was recently formulated by the author in the case where the given finite group linearly acts on an affine space. In cases of very special groups and representations, the conjecture has been verified and related stringy invariants have been explicitly computed. In this paper, we try to generalize the conjecture and computations to more complicated situations such as non-linear actions on possibly singular spaces and non-permutation representations of non-abelian groups.

Contents
1. Introduction
2. Motivic integration and stringy motifs
3. G-arcs
4. The untwisting technique revisited
5. The change of variables formula
6. The McKay correspondence for linear actions
7. The McKay correspondence for non-linear actions
8. Computing boundaries of untwisting varieties
9. A tame singular example
10. A wild non-linear example
11. Stable hyperplanes in permutation representations
12. Some S4 -masses in characteristic two
13. Concluding remarks
References

1
5
10
13
17
20
23
25
26
28
31
35
38
40

1. Introduction
The McKay correspondence in terms of stringy invariants was first studied by Batyrev and Dais [BD96], and Batyrev [Bat99]. Denef and Loeser [DL02] later took a more conceptual approach, where the McKay correspondence directly follows from the theory of motivic integration suitably generalized to a situation involving finite group actions.
These works were confined to characteristic zero. Except some works for the tame case (the finite group has order coprime to the characteristic), an attempt of generalization to arbitrary characteristics, including the wild (non-tame) case, was only recently
2010 Mathematics Subject Classification. 14E18, 14E16, 11S15.
This work was partially supported by Grants-in-Aid for Scientific Research (22740020).
1

WILDER MCKAY CORRESPONDENCES

2

started in [Yasa, Yasb]. There it was formulated a conjectural generalization of results in characteristic zero. Subsequently it turned out in [WYa] that the conjecture is closely related to the number theory, in particular, the problem of counting local Galois representations. In these papers, however, only linear actions on affine spaces have been discussed. In characteristic zero, since every finite group action on a smooth variety is locally linearizable, many studies can be reduced to the linear case. This is no longer true in positive or mixed characteristic. The conjecture has been verified in very special cases, by computing stringy invariants explicitly. The aims of this paper are firstly to generalize the conjecture to non-linear actions on a (possibly singular) affine variety, and secondly to make it possible to compute stringy invariants in more complicated examples.
We now recall the conjecture from [Yasb]. Set the base scheme to be D = Spec OD
with OD a complete discrete valuation ring and suppose that its residue field, denoted by k, is algebraically closed.
Remark 1.1. Working over a discrete valuation ring rather than a field is natural in our arguments. We can easily switch from a field to a discrete valuation ring by the base change by Spec k[[t]] → Spec k.
We consider a linear action of a finite group G on the affine d-space V = AdD over D
and the associated quotient scheme X := V /G.
Conjecture 1.2 (The wild McKay correspondence conjecture [Yasb]). Let o ∈ X(k)
denote the image of the origin and Mst (X)o the stringy motif of X at o. Suppose that the quotient morphism V → X is étale in codimension one. Then
ˆ
Mst (X)o =
Lw dτ.
G-Cov(D)

Here G-Cov(D) is the (conjectural) moduli space of G-covers of D, w is the weight function on G-Cov(D) associated to the G-representation V and τ is the tautological motivic measure on G-Cov(D).
In [Yasa], Conjecture 1.2 was verified, when OD = k[[t]] with k of characteristic p > 0
and G is the cyclic group of order p. In [WYa], a variant conjecture was verified when the symmetric group Sn acts on A2n
D by two copies of the standard representation. In the same paper, a generalization to the case where k is only perfect was formulated by modifying the function w.
Roughly the conjecture was derived as follows: we first express Mst (X)o as a motivic integral over the space of arcs of X, that is, D-morphisms D → X. We then transform the motivic integral to a motivic integral over the space of G-arcs of V , that is, G-equivariant D-morphisms E → V for G-covers E → D. We can see that the contribution of each G-cover E → D to Mst (X)o is Lw(E) , and hence the conjecture.
At this last point, we used the technique of untwisting, which enables us to reduce the study of G-arcs to the one of ordinary arcs. A prototype of untwisting was introduced by Denef and Loeser [DL02]. In [Yasb], the author developed it so that we can use it even in the wild case. In this paper, we will refine the technique slightly more. For each
G-cover E of D with a connected component F , we can construct another affine space

WILDER MCKAY CORRESPONDENCES

3

V |F | ∼
= AdD and a morphism V |F | → X such that there is a correspondence between
G-arcs of V and ordinary arcs of V |F | . Through the correspondence, we can represent the contribution of E to Mst (X)0 as a motivic integral over the ordinary arcs D → V |F | .
Our strategy of generalization to the non-linear case is quite simple: given an affine variety v with a G-action, we equivariantly embed v into an affine space V with a linear G-action. For each G-cover E → D with a connected component F , we take the subvariety v|F | ⊂ V |F | corresponding to v ⊂ V , which plays the same role as V |F | in the linear case.
We also need an idea from the minimal model program, that is, working with varieties endowed with divisors rather than varieties themselves. Encapsulating one more information, we will introduce the notion of centered log structures or centered log Dvarieties, which are just triples X = (X, ∆, W ) of a normal D-variety X, a Q-divisor
∆ and a closed subset W of X ⊗OD k with KX/D + ∆ Q-Cartier. It is straightforward to generalize the stringy motif to centered log D-varieties. We write it as Mst (X). For instance, the stringy motif Mst (X)o mentioned above is the same as Mst ((X, 0, {o})).
Returning to the equivariant immersion v ֒→ V , if v is given a centered log structure v = (v, δ, w), then there exist unique centered log structures x on x := v/G and v|F |,ν on the normalization v|F |,ν of v|F | so that all the morphisms connecting them are crepant
(see Section 2.2 for details). If H ⊂ G is the stabilizer of the component F ⊂ E, then the centralizer CG (H) of H acts on v|F |,ν and on its arc space J∞ v|F |,ν . We will define
Mst,CG (H) (v|F |,ν ) in the same way as defining the ordinary stringy motif except that we will use the quotient space (J∞ v|F |,ν )/CG (H) rather than the arc space J∞ v|F |,ν itself.
We formulate the following conjecture which generalizes Conjecture 1.2.
Conjecture 1.3 (Conjecture 7.3). We have
ˆ
Mst (x) =
Mst,CG (H) (v|F |,ν ) dτ.
G-Cov(D)

We will verify Conjecture 1.3 in two examples from the simplest ones, computing both sides of the equality independently. One example is a tame action on a singular variety, the other a wild non-linear action on a smooth variety.
Keeping arguments above in mind, let us return to the linear case. One difficulty in computing the right side of the equality in Conjecture 1.2 is computing weights w(E)
explicitly, and another is computing the moduli space G-Cov(D). Taking an equivariant immersion v ֒→ V is useful also in solving the former difficulty for some linear actions.
With E, F and H as before, if V = AdD has a linear G-action and if V0 := V ⊗OD k, then the weight of E with respect to V is, by definition, wV (E) = codim(V0H , V0 ) − vV (E)
with another function vV on G-Cov(D) and V0H the H-fixed point locus in V0 . The first term, codim(V0H , V0 ), is easy to compute, while the second generally not. However, if
G acts on V by permutations of coordinates, then vV (E) is represented in terms of the discriminant [WYa]: in this situation, we can associate a degree d cover C → D to a
G-cover E → D and dC/D
vV (E) =
2

WILDER MCKAY CORRESPONDENCES

4

with dC/D the discriminant exponent of the cover C → D. We generalize this equality to hyperplanes in V defined by a G-invariant linear form. For simplicity, we consider the case where v ⊂ V is defined by x1 + · · · + xd = 0
with x1 , . . . , xd coordinates of V .
Proposition 1.4 (See Corollary 11.4 for a slightly more general result). Let C =
Fl j=1 Cj be the decomposition of C into the connected components. Then



dCj /D
dC/D
vv (E) =
|1≤j≤l
− min
2
[Cj : D]

with [Cj : D] the degree of Cj → D.

Using this and assuming a motivic version of Krasner’s formula [Kra66] for counting local field extensions, we explicitly compute
ˆ
ˆ
−v3v
L
dτ and
Lw3v dτ,
G-Cov(D)

G-Cov(D)

when G = S4 acts on V ∼
= A4D by the standard representation, v ⊂ V is given by x1 +x2 +
x3 + x4 = 0 and 3v is the direct sum of three copies of v. We find that the two integrals are dual to each other. The same kind of duality was observed in [WYa] and will be discussed in [WYb] in more details. Obtained formulas for these integrals are thought of as (motivic) variants of mass formulas for local Galois representations [Bha07, Ked07,
Woo08] with respect to weights coming from a non-permutation representation.
From Section 2 to 6, we review the theory in the linear case and finally formulate the
McKay correspondence for linear actions. Most materials here are not new and found for instance in [Yasb], although arguments are refined and adjusted to our purpose. In
Section 7, we formulate the McKay correspondence for non-linear actions. In Section
8, we study how to determine the centered log structure v|F |,ν under some assumptions.
In Sections 9 and 10, we will compute non-linear examples. In Sections 11 and 12, we treat hyperplanes in permutation representations. We will end the paper with concluding remarks in Section 13.
1.1. Acknowledgments. I wish to thank Yusuke Nakamura and Shuji Saito for useful discussions directing my interests to non-linear actions, Johannes Nicaise and Julien
Sebag for their kind answers to my questions on motivic integration over formal schemes, and Melanie Wood for stimulating discussions in our joint work.
1.2. Convention and notation. If X is an affine scheme, OX denotes its coordinate ring. By the same symbol OX , we sometimes denote also the structure sheaf on a scheme
X. This abuse of notation would not cause any problem. When a group G acts on X
from left, then we suppose that G acts on OX from right: for g ∈ G, if φg : X → X is the g-action on X, then g acts on OX by the pull-back of functions by φg . Throughout the paper, we fix an affine scheme D with OD a complete discrete valuation field. We denote the residue field of OD by k and suppose that k is algebraically closed. For an integral scheme X, we denote by K(X) its function field. If X is affine, then K(X) is

WILDER MCKAY CORRESPONDENCES

5

the fraction (quotient) field of the ring OX . Again, by abuse of notation, K(X) also denotes the constant sheaf on X associated to the function field. For a D-scheme X, we denote by X0 the special fiber with the reduced structure: X0 := (X ×D Spec k)red .
2. Motivic integration and stringy motifs
In this section, we review the theories of motivic integration over ordinary (untwisted)
arcs and stringy invariants, mainly developed in [Kon95, DL99, Bat98, Bat99, DL02,
Seb04].
2.1. Centered log varieties. We call an integral D-scheme X a D-variety if X is flat, separated and of finite type over D and X is smooth over D at the generic point of X.
For a D-variety X, we denote the smooth locus of X by Xsm and the regular locus by
Xreg .
Let X be a normal D-variety. We can define the canonical sheaf ωX = ωX/D of X
V
over D as in [Kol13, page 8]. On Xsm , the canonical sheaf is isomorphic to d ΩX/D
withVd the relative dimension of X over D. Therefore we can think of ωX as a subsheaf of ( d ΩX/D ) ⊗ K(X). We define the canonical divisor of X, denoted by KX = KX/D , to be the linear equivalence class of Weil divisors corresponding to ωX .
A log D-variety is a pair (X, ∆) of a normal D-variety X and a Weil Q-divisor ∆ such that KX + ∆ is Q-Cartier. We call ∆ the boundary of the log variety. The canonical divisor of a log D-variety (X, ∆) is K(X,∆) := KX + ∆.
A centered log D-variety is a triple X = (X, ∆, W ) such that (X, ∆) is a log D-variety and W is a closed subset of X0 , where X0 denotes the special fiber of the structure morphism X → D with the reduced structure. We call W the center of X. We also say that X is a centered log structure on X. For a centered log variety X = (X, ∆, W ), we define a canonical divisor of X as the one of (X, ∆):
KX := K(X,∆) = KX + ∆.
sometimes, we identify a normal Q-Gorenstein (KX is Q-Cartier) D-variety X with the log D-variety (X, 0), and identify a log D-variety (X, ∆) with the centered log
D-variety (X, ∆, X0 ):
normal Q-Gorenstein  
/ {log D-varieties}  
/ {centered log D-varieties}
(2.1)
D-varieties
X✤

/ (X, 0)

(X, ∆) ✤

/ (X, ∆, X0 )

2.2. Crepant morphisms. For centered log D-varieties X = (X, ∆, W ) and X′ =
(X ′ , ∆′ , W ′ ), a morphism f : X → X′ is just a morphism f : X → X ′ of the underlying varieties with f (W ) ⊂ W ′ . We say that a morphism X → X′ is proper or birational if it

WILDER MCKAY CORRESPONDENCES

6

is so as the morphism f : X → X ′ of the underlying varieties. We say that a morphism f : X → X′ is crepant if f −1 (W ′ ) = W and KX = f ∗ KX′ .
The right equality should be understood that for r ∈ Z>0 such that r(KX + ∆) and r(KX ′ + ∆′ ) are Cartier, we have a natural isomorphism
[r]
[r]
ωX ′ (r∆′ ) ∼
= f ∗ ωX (r∆).
[r]

⊗r
Here ωX (r∆) is the invertible sheaf which is identical to ωX
(r∆) on Xreg . We adopt this convention throughout the paper.
Given a generically étale morphism f : X → X ′ of normal D-varieties, a centered log structure X′ on X ′ induces a unique centered log structure X on X such that the morphism f : X → X′ is crepant. Conversely, if f : X → X ′ is additionally proper, then for each centered log structure X on X, there exists at most one centered log structure on X′ such that f : X → X′ is crepant.

Remark 2.1. For our purpose, we may slightly weaken the assumptions in the definition of crepant morphisms. For instance, concerning the equality f −1 (W ′ ) = W , we only need this equality outside Xsm \ Xreg . This is because the locus Xsm \ Xreg does not contribute to stringy motifs at all, which will be defined below. However, for simplicity, we will cling to our definition as above.
2.3. Motivic integration. Let X = (X, ∆, W ) be a centered log D-variety. An arc of
X is a D-morphism D → X sending the closed point of D into W . The arc space of X, denoted J∞ X, is a k-scheme parameterizing the arcs of X. We put Dn := Spec OD /mn+1
D
with mD the maximal ideal of OD . An n-jet of X is a D-morphism Dn → X sending the unique point of Dn into W . For each n, there exists a k-scheme Jn X parametrizing n-jets of X. For n ≥ m, we have natural morphisms Jn X → Jm X and the arc space
J∞ X is identified with the projective limit of Jn X, n ≥ 0 with respect to these maps.
We have the induce maps
πn : J∞ X → Jn X.
For n < ∞, Jn X are of finite type over k. For a morphism f : Y → X and each n ∈ Z≥0 ∪ {∞}, there exists a natural map fn : Jn Y → Jn X.
The arc space J∞ X has the so-called motivic measure, denoted by µJ∞ X . The measure takes values in some (semi-)ring, say R, which is often a suitable modification of the
Grothendieck (semi-)ring of k-varieties. In this paper, we will fix R satisfying the following properties: denoting by [T ] the class of a k-variety T in R, we have
• for a bijective morphism S → T , we have [S] = [T ] in R,
• putting L := [A1k ],Pwe have all fractional powers La , a ∈ Q in R, ai
• an infinite series ∞
with limi→∞ dim Ti + ai = −∞ converges, i=1 [Ti ]L
• for a morphism f : S → T and for n ∈ Z≥0 , if every fiber of f admits a homeomorphism from or to the quotient Ank /G for some linear action of a finite group G on Ank , then [S] = [T ]Ln .

WILDER MCKAY CORRESPONDENCES

7

One possible choice is the field of Puiseux series in t−1 ,
R :=

∞
[

Z((t−1/r )),

r=1

where we put [T ] to be the Poincaré polynomial as in [Nic11].
A subset A ⊂ J∞ X is called stable if there exists n ∈ Z≥0 such that πn (A) ⊂ Jn X is a constructible subset and A = πn−1 πn (A) and for every m ≥ n, every fiber of the map
πn+1 (A) → πn (A) is homeomorphic to Ank . The measure of a stable subset A is given by
µJ∞ X (A) := [πn (A)]L−nd (n ≫ 0).
More generally, we can define the measure for measurable subsets, which are roughly the limits of stable subsets.
Let Φ : C → R ∪ {∞} be a measurable function on a subset C ⊂ J∞ X, that is, the image of Φ is countable, all fibers Φ−1 (a) are measurable and µJ∞ X (Φ−1 (∞)) = 0. We define
ˆ
X
Φ µJ∞ X :=
µJ∞ X (Φ−1 (a)) · a ∈ R ∪ {∞}.
C

a∈R

2.4. Stringy invariants. We still suppose that X = (X, ∆, W ) is a centered log Dvariety.
Definition 2.2. To a coherent ideal sheaf I 6= 0 on X defining a closed subscheme
Z ( X, we associate the order function, ord I = ord Z : J∞ X → Z≥0 ∪ {∞}, as follows: for an arc γ : D → X, the pull-back γ −1 I of I is an ideal of OD and of the form mlD for some l ∈ Z≥0 ∪ {∞}, where we put (0) := m∞
D by convention. For a fractional ideal I (that is, a coherent OX -submodule of K(X)), if we write I = I+ · I−−1
for ideal sheaves I+ and I− with I− locally principal, then we put ord I := ord I+ − ord I− .
Here we put ord IP
= ∞ if either ord I+ = ∞ or ord I− = ∞. Similarly, for a Q-linear combination Z = ni=1 ai Zi of closed subschemes Zi ( X, we define ord Z :=

n
X

ai · ord Zi ,

i=1

taking values in Q ∪ {∞}.

Remark 2.3. For a closed subscheme Z ( X, we expect that (ord Z)−1 (∞) has measure zero. The author does not know if this has been proved, but this follows from the change of variables formula, if there exists a resolution of singularities f : X̃ → X so that X̃ is regular and X̃0 ∪f −1 (Z) is a simple normal crossing divisor. If the expectation is actually true, then order functions for fractional ideals and Q-linear combination of closed subschemes are well-defined modulo measure zero subsets.

WILDER MCKAY CORRESPONDENCES

8
[r]

Let r ∈ Z>0 be such that rKX is Cartier. Since the sheaf OX (rKX ) = ωX (r∆) is
V
invertible and thought of as a subsheaf of the constant sheaf ( d ΩX/D )⊗r ⊗ K(X), we
V
can define a fractional ideal sheaf IXr by the equality of subsheaves of ( d ΩX/D )⊗r ⊗
K(X),
!⊗r d
^
ΩX/D
/tors = IXr · OX (rKX ).
We then put a function fX on J∞ X by
1
fX := ord IXr .
r
Since (IXr )n = IXr·n , the function fX is independent of the choice of r. If X is smooth, then we simply have fX = ord ∆.
Definition 2.4. The stringy motif of X is defined to be
ˆ
Mst (X) :=
LfX dµJ∞ X .
J∞ X

We also write Mst (X) = Mst (X, ∆)W and sometimes omit ∆ if ∆ = 0, and W if
W = X0 . When the integral above converges, we call X stringily log terminal. When diverges, we put Mst (X) := ∞.
Conjecture 2.5. If a morphism f : X → X′ of centered log D-varieties is proper, birational and crepant, then
Mst (X) = Mst (X′ ).
Proposition 2.6. Conjecture 2.5 holds if there exists a proper birational morphism
Y → X of D-varieties such that Y ⊗OD K(D) is smooth over K(D). In particular,
Conjecture 2.5 holds if K(D) has characteristic zero.
Proof. Let Xη be the generic fiber of X → D. From the Hironaka theorem, there exists a coherent ideal sheaf Iη ⊂ OXη such that the blowup of Xη along Iη is smooth over
K(D). Let I ⊂ OX be an coherent ideal sheaf such that I|Xη = Iη . The blowup of X
along I has a smooth generic fiber. Hence the second assertion of the lemma follows from the first one.
To prove the first one, we can apply the version of the change of variables formula proved by Sebag [Seb04, Th. 8.0.5] (see also [NS11]). If Y is the centered log structure on Y such that the induced morphism f : Y → X is crepant, then the change of variables formula shows that
ˆ
ˆ
fX
L dµJ∞ X =
LfX ◦f∞ −ord jacf dµJ∞ Y .
J∞ X

J∞ Y

WILDER MCKAY CORRESPONDENCES

9

Here ord jacf is the function of Jacobian orders as defined in [Seb04, page 29]. For r ∈ Z>0 such that rKX and rKY are Cartier, we have

!⊗r 
d
^
f ∗
 /tors = f −1 IXr · OY (rKY ) and
ΩX/D
d
^

This shows that

ΩY /D

!⊗r

/tors = IYr · OY (rKY ).

fX ◦ f∞ − ord jacf = fY .
We obtain Mst (X) = Mst (Y), and similarly Mst (X′ ) = Mst (Y). We have proved the proposition.

Proposition 2.7. Let X = (X, ∆, W ) be a centered log D-variety and write
∆=

l
X

ah Ah +

m
X

bi Bi +

cj C j

(ah , bi ∈ Q, cj ∈ Q \ {0})

j=1

i=1

h=1

n
X

such that Ah are the irreducible components of the closure of X0 ∩ Xsm , Bi are the irreducible components of X0 \ Xsm and Cj are prime divisors not contained in X0 . Let


A◦h := Ah ∩ Xsm = Ah \ X0 \ Ah and
\
[
CJ◦ :=
Cj \
Cj , j∈J

j∈{1,..,n}\J

S
with X0 \ Ah the closure of X0 \ Ah . We suppose that X is regular and that nj=1 Cj is simple
T normal crossing, that is, for any J ⊂ {1, . . . , n}, the scheme-theoretic intersection j∈J Cj is smooth over D. Then X is stringily log terminal if and only if cj < 1
for every j with Cj ∩ W ∩ Xsm 6= ∅. Moreover, if it is the case,
Mst (X) =

l
X
h=1

Lah

X

[W ∩ A◦h ∩ CJ◦ ]

Y
j∈J

J⊂{1,...,n}

L−1
L1−cj − 1

.

Proof. We first note that the locus X0 \Xsm does not have any arc, hence not contribute to Mst (X). Since l
G
X0 ∩ Xsm =
A◦h , h=1

we can decompose Mst (X) into the sum of components corresponding to A◦h , h = 1, .., l.
The divisor ah Ah contributes to the component corresponding to A◦h by the multiplication with Lah . From all these arguments, the proposition has been reduced to the formula
X
Y L−1
Mst (X) =
[W ∩ CJ◦ ]
L1−cj − 1
j∈J
J⊂{1,...,n}

WILDER MCKAY CORRESPONDENCES

in the case where X is smooth and ∆ =
(see for instance [Bat98]).

P

10

j cj Cj . This is the standard explicit formula



2.5. Group actions.
Definition 2.8. A centered log G-D-variety is a centered log D-variety X = (X, ∆, W )
endowed with a faithful G-action on X such that ∆ and W are stable under the Gaction. Given a variety X with a faithful G-action, we say that a centered log structure
X on X is G-equivariant if it is a centered log G-D-variety.
For a centered log G-D-variety X, the arc space J∞ X has a natural G-action. We define a motivic measure on (J∞ X)/G, denoted by µ(J∞ X)/G , in the same way as defining the motivic measure on J∞ X except that in the definition of stable subsets, say A, fibers of πn+1 (A) → πn (A) are only assumed to be homeomorphic the quotient Adk /H for some linear action of a finite group H on Adk .
The function fX on J∞ X is G-invariant and gives a function on (J∞ X)/G, which we again denote by fX . We define
ˆ
Mst,G (X) :=
LfX dµ(J∞ X)/G .
(J∞ X)/G

The reader should not confuse Mst,G (X) with the orbifold stringy motif MstG (X), defined later.
P
Let us define a G-prime divisor as a divisor of the form li=1 Di , where Di are prime divisors permuted transversally by the G-actions,

Proposition 2.9. Let X = (X, ∆, W ) be a centered log G-D-variety and write
∆=

l
X
h=1

ah Ah +

m
X
i=1

bi Bi +

n
X

cj C j

(ah , bi ∈ Q, cj ∈ Q \ {0})

j=1

S
such that Ah are the distinct G-prime divisors such thatS Ah is equal to the closure of X0 ∩ Xsm , Bi are the distinct G-prime divisors with i Bi = X0 \ Xsm and Cj are
G-prime divisors not contained in X0 . We suppose
• X is regular,
T
• for any J ⊂ {1, . . . , n}, the scheme-theoretic intersection j∈J Cj is smooth over
D, and
• for every j with Cj ∩ W ∩ Xsm 6= ∅, cj < 1.
With the same notation as in Proposition 2.7, we have l
X
X  W ∩ A◦ ∩ C ◦  Y L − 1
h
J
ah
Mst,G (X) =
L
.
1−cj − 1
G
L
h=1
j∈J
J⊂{1,...,n}

3. G-arcs
In the last subsection, we considered motivic integration over varieties endowed with finite group actions. However we considered only ordinary (untwisted) arcs, which are not general enough to apply to the McKay correspondence. Suitably generalized arcs were introduced by Denef and Loeser [DL02] in characteristic zero. The author

WILDER MCKAY CORRESPONDENCES

11

[Yasb] further generalized them to arbitrary characteristics. We may use generalized arcs of orbifolds or Deligne-Mumford stacks as in [LP04, Yas04, Yas06, Yasb] so that we can treat general orbifolds, having group actions only locally. We do not pursue generalization in this direction, however.
From now on, we fix a finite group G.
Definition 3.1. By a G-cover of D, we mean a D-scheme E endowed with a left
G-action such that E ⊗OD K(D) is an étale G-torsor over Spec K(D) and E is the normalization of D in OE⊗OD K(D) . We denote by G-Cov(D) the set of G-covers of D
up to isomorphism.
Remark 3.2. In the tame case, there is a one-to-one correspondence between the points of G-Cov(D) and the conjugacy classes in G. In the wild case, however, G-Cov(D)
is expected to be an infinite dimensional space having a countable stratification with finite-dimensional strata.
We now fix the following notation: E is a G-cover of D, F is a connected component of E with a stabilizer H so that F is an H-cover of D.
E ❅o

conn. comp.

❅❅
❅❅
❅❅
G-cover ❅

D

? _F
⑦
⑦⑦
⑦⑦
⑦
~⑦⑦ H-cover

Lemma 3.3. Let Aut(E) be the automorphism group of E as a G-cover of D, that is, it consists of G-equivariant D-automorphisms of E. We have a natural isomorphism
Aut(E) ∼
= CG (H)op , where the right side is the opposite group of the centralizer of H in G.
Proof. If E is the trivial G-cover D × G of D, then its automorphisms are nothing but the right G-action on G. Therefore Aut(E) = Gop .
For the general case, let EF be the normalization of the fiber product E ×D F . This is a trivial G-cover of F and we have a natural injection
Aut(E) → Aut(EF ) = Gop .
Its image is the automorphisms of EF compatible with the action of Gal(F/D) = H.
This shows the lemma.

Let V be a D-variety endowed with a faithful left G-action.
Definition 3.4. We define an E-twisted G-arc of V as a G-equivariant D-morphism
E → V and a G-arc of V as an E-twisted G-arc for some E. Two G-arcs E → V and
E ′ → V are said to be isomorphic if there exists a G-equivariant D-isomorphism E → E ′
G,E
compatible with the morphisms to V . We denote by J∞
V the set of isomorphism
G
classes of E-twisted G-arcs of V and by J∞ V the set of isomorphism classes of G-arcs of V .

WILDER MCKAY CORRESPONDENCES

12

Obviously,
G

G
J∞
V =

G,E
J∞
V.

E∈G-Cov(D)

Let HomG
D (E, V ) be the space of G-equivariant D-morphisms E → V . We define a left op action of CG (H) = Aut(E)op on HomG
D (E, V ) as follows: for a ∈ CG (H) = Aut(E)
G
and f ∈ HomD (E, V ),
(3.1)

f

f

a

f

a

a · (E ←
− V ) := (V ←
−E←
− E) = (V ←
−V ←
− E).

By definition, we have
G,E
J∞
V = HomG
D (E, V )/CG (H).

(3.2)

S
For n ∈ Z≥0 , we put Fn := F/mFn·h+1 with h := ♯H and define En := g∈G g(Fn ).
In particular, F0 ∼
= Spec k and E0 consists of the closed points of E with the reduced scheme structure.
Definition 3.5. We define an E-twisted G-n-jet of V as a G-equivariant D-morphism
En → V and put
JnG,E V := HomG
D (En , V )/CG (H) and
G
JnG V =
JnG,E V.
E∈G-Cov(D)

Here the CG (H)-action on HomG
D (En , V ) is similarly defined as (3.1).
Note that if E 6∼
= E ′ , then E-twisted and E ′ -twisted G-n-jets never give the same point of JnG V . For each n ∈ Z≥0 and E ∈ G-Cov(D), we have natural maps
G,E
G
J∞
V → JnG,E V and J∞
V → JnG V,

both of which we will denote by πn . We have obtained the following commutative diagram:
G,E
J∞
V

πn+1

/ J G,E V
n

/ {E}




/ J GV


/ G-Cov(D)

n+1



G
J∞
V

/ J G,E V

πn+1

/ JG

n+1 V

n

Remark 3.6. In [Yasb], the author conjectured that the sets G-Cov(D), JnG,E V , JnG V
(0 ≤ n < ∞) are realized as k-schemes admitting stratifications with at most countably many finite-dimensional strata, which will be necessary below to define the motivic measure.
Let X be the quotient scheme V /G, writing the quotient morphism as p : V → X.
Given a G-arc E → V , we get an arc D → X by taking the G-quotients of the source and the target. This gives a natural map
G
p∞ : J∞
V → J∞ X.

WILDER MCKAY CORRESPONDENCES

13

Let T ⊂ V be the ramification locus of π say with the reduced scheme structure and
T̄ ⊂ X its image. The map p∞ restricts to the bijection
G
G
J∞
V \ J∞
T → J∞ X \ J∞ T̄ .

For n < ∞, we have a natural map
G
pn : πn (J∞
V ) → Jn X,
G
where πn denotes the natural map J∞
V → JnG V .
For a centered log G-D-variety V and n ∈ Z≥0 ∪ {∞}, we define JnG V and JnG,E V as the subsets of JnG V and JnG,E V consisting of the morphisms En → V sending the closed points of En into the center of V.

4. The untwisting technique revisited
In this section, we revisit the technique of untwisting, which was first used by Denef and Loeser [DL02] in characteristic zero and generalized to arbitrary characteristics by the author [Yasb]. Our constructions below are slightly different and refined from the ones in [Yasb].
Let us now turn to the case where V is an affine space over D and the given G-action is linear. We keep to fix a G-cover E of D and a connected component F of E with stabilizer H.
•
For a free OD -module M of rank d, let OV := SO
M be its symmetric algebra and
D
put
V = Spec OV = AdD .
We suppose that the module M and hence the OD -algebra OV have faithful right Gactions. Then V has the induced left G-action. The set HomG
D (E, V ) can be identified with the OD -module
G
ΞF := HomH
OD (M, OF ) = HomOD (M, OE ).

We call ΞF the tuning module.
Remark 4.1. If we fix a basis of M, then the module HomOD (M, OE ) is identified with
OE⊕d . This module OE⊕d has two G-actions: one is the diagonal G-action induced from the given G-action on OE and the other is the one induced from the G-action on M.
For an element of OE⊕d corresponding to a G-equivariant map M → OE , two actions must coincide. We thus can identify ΞF with the locus in OE⊕d where the two actions coincide. This was how the module ΞF was presented in previous papers [Yasb, WYa].
Lemma 4.2 ([Yasb, WYa]). The module ΞF is a free OD -module of rank d. Moreover it is a saturated OD -submodule of HomOD (M, OF ) and of HomOD (M, OE ): for a ∈ OD
and f ∈ HomOD (M, OE ), if af ∈ ΞF , then f ∈ ΞF .
From (3.2),
G,E
J∞
V = ΞF /CG (H)

Note that the CG (H)-action on ΞF is OD -linear.

WILDER MCKAY CORRESPONDENCES

14

Lemma 4.3. The maps
G
G
πn+1 (J∞
V ) → πn (J∞
V)

have fibers homeomorphic to the quotient of Adk by a linear action of a finite group.
Proof. If we denote the map ΞF → HomH
D (Fn , V ) again by πn , the image πn (ΞF ) is n+1 ⊕d isomorphic to (OD /mD ) . This shows that the fibers of
πn+1 (ΞF ) → πn (ΞF )
are isomorphic to Adk , proving the lemma.



G
Definition 4.4. We define a motivic measure µJ∞
G V on J
∞ V in the same way as the ones on J∞ V and (J∞ V )/G. If V is a G-equivariant centered log structure on V , we
G
define the measure µJ∞
G V on J
GV .
∞ V as the restriction of µJ∞

Remark 4.5. For the definition above making sense, we need the conjecture that moduli spaces G-Cov(D) and JnG V exist and have some finiteness (see Remark 3.6).
Definition 4.6. We define modules,
M |F | := HomOD (ΞF , OD ) and
M hF i := M |F | ⊗OD OF = HomOD (ΞF , OF ), which are free modules of rank d over OD and OF respectively. We define an OD -linear map u∗ = u∗F : M → M hF i m 7→ (Ξ ∋ f 7→ f (m) ∈ OF ),
G
identifying ΞF with HomH
OD (M, OF ) rather than HomOD (M, OE ).

Lemma 4.7. We suppose that H and CG (H) acts on M by restricting the given Gaction.
(1) With respect to the H-action on M hF i induced from the H-action on OF , the map u∗ is H-equivariant.
(2) With respect to the CG (H)-action on M hF i induced from the (left) CG (H)-action on ΞF , the map u∗ is CG (H)-equivariant.
Proof.

(1) For h ∈ H and m ∈ M, we have u∗ (mh) = (f 7→ f (mh))
= (f 7→ f (m)h)
= (f 7→ f (m))h,

since f ∈ Ξ are H-equivariant. This shows the assertion.
(2) Let M hEi := HomOD (ΞF , OE ) and consider the natural map u∗E : M → M hEi , m 7→ (f 7→ f (m)),

WILDER MCKAY CORRESPONDENCES

15

now identifying ΞF with HomG
OD (M, OE ). This map is CG (H)-equivariant. Indeed, for g ∈ CG (H) and m ∈ M, from (3.1), we have u∗E (mg) = (f 7→ f (mg))
= (f 7→ f (m)g)
= (f 7→ (gf )(m)).
The map u∗E factors as u∗

F
M −→
M hF i ֒→ M hEi .

Since the inclusion M hF i ֒→ M hEi is also CG (H)-equivariant, so does u∗F .

Note that the H- and CG (H)- actions above on M hF i commute.
Definition 4.8. We define the untwisting variety (resp. pre-untwisting) variety of V
with respect to F as
•
•
V |F | := Spec SO
M |F | = AdD (resp. V hF i := Spec SO
M hF i = AdF ).
D
F

We denote the projection V hF i → V |F | by r = rF , r standing for the restriction of scalars (see diagram (4.1) below). The map u∗ defines a D-morphism u : V hF i → V, which is both H- and CG (H)-equivariant. We call the pair of r and u the untwisting correspondence of V with respect to F .
Let X := V /G and identify OX with (OV )G . Since the H-invariant subring of OV hF i is
(OV hF i )H = OV |F | , we have u∗ (OX ) ⊂ OV |F | .
We denote the induced morphism V |F | → X by p|F | . We have the following commutative diagram:
V hF i = AdF

(4.1)

qq

u qqq

qqq x qq q

V = AdD

▼▼▼
▼▼▼
▼
p ▼▼▼▼
&

◆◆◆
◆◆◆r
◆◆◆
◆◆'

V |F | = AdD

♦
♦♦♦
♦
♦
♦♦
♦
w ♦♦ p|F |

X = V /G

Lemma 4.9.

(1) The map hF i
HomH
) → ΞF = HomH
F (F, V
D (F, V )
γ 7→ u ◦ γ

is bijective.

WILDER MCKAY CORRESPONDENCES

16

(2) The map hF i
HomH
) → J∞ V |F | = HomD (D, V |F | )
F (F, V
sending a morphism F → V hF i to the induced one of quotients,

D = F/H → V |F | = V hF i /H, is bijective.
Proof.

(1) With the identification, hF i
HomH
) = HomH
F (F, V
OF (HomOD (ΞF , OF ), OF ).

the map of the assertion is identified with the map a : HomH
OF (HomOD (ΞF , OF ), OF ) → ΞF
φ 7→ (m 7→ φ((f 7→ f (m)))), where m ∈ M and f ∈ ΞF . Let us consider the map b : ΞF → HomH
OF (HomOD (ΞF , OF ), OF )
f 7→ (z 7→ z(f )), where z ∈ HomOD (ΞF , OF ). The composition a ◦ b sends f ∈ ΞF to
(m 7→ (z 7→ z(f )) (h 7→ h(m)))
= (m 7→ f (m))
= f, and hence is the identity map. It follows that a is surjective. Now the assertion follows from the fact that the source and target of a are free OD -modules of the same rank and a is a homomorphism of OD -modules.
(2) We can give the converse by the base change associated to F → D.

In summary, we have a one-to-one correspondence between ΞF and J∞ V |F | , induced from the untwisting correspondence. From Lemma 4.7, the correspondence is compatible with the CG (H)-actions on both sides. Therefore it descends to a one-to-one
G,E
correspondence between J∞
V and (J∞ V |F | )/CG (H). We obtain the following commutative diagram: hF i
HomH
)
F (F, V

(4.2)

♦♦7

1-to-1♦♦♦

ΞF o

♦
♦♦♦
♦
♦
w♦♦



G,E
J∞
V o

❖❖❖
❖❖❖
❖❖
p∞ ❖❖❖
❖'

i❙❙❙
❙❙❙1-to-1
❙❙❙
❙❙❙
❙❙❙
)
1-to-1
/ J∞ V |F |

/ (J∞ V |F | )/CG (H)
❦❦
❦❦❦❦
❦
❦
❦❦
❦❦❦
u ❦❦
❦

1-to-1

J∞ X

WILDER MCKAY CORRESPONDENCES

17

For n < ∞, we have a similar diagram: hF i
πn (HomH
))
F (F, V

(4.3)

i❙❙❙❙
❙❙❙1-to-1
❙❙❙❙
❙❙❙❙
❙❙)

❧

β ❧❧❧❧❧

πn (ΞF )

❧❧
❧❧❧
v v ❧❧
❧

Jn V |F |




(Jn V |F | )/CG (H)

G,E
πn (J∞
V)

❘❘❘
❘❘❘
❘❘❘
❘❘❘
❘❘❘
)

❥❥
❥❥❥❥
❥
❥
❥
❥
❥❥❥❥
u❥❥❥❥

Jn X

Note that the arrow β has no longer bijective. When n = 0, the diagram is represented as: hF i

(4.4)
rr

V0

β rrr

(V0 )H

H

rrr yry rr

f▲▲▲
▲▲▲1-to-1
▲▲▲
▲▲▲
&

|F |

V0




|F |
(V0 )/CG (H)

(V0 ) /CG (H)

▼▼▼
▼▼▼
▼▼▼
▼▼▼
&

X0

qq qqq q
q qq x qq q

Here (V0 )H is the fixed-point locus of the H-action on V0 .
5. The change of variables formula
The untwisting technique, discussed in the last section, enables us to deduce a conG
jectural change of variables formula for the map p∞ : J∞
V → J∞ X. In turn, it will derive the McKay correspondence for linear actions in the next section.
We keep the notation from the last section.
Definition 5.1. Let f : T → S be a morphism of D-varieties which is generically étale.
The Jacobian ideal (sheaf )
Jacf = JacT /S ⊂ OT
is defined as the 0-th Fitting ideal (sheaf) of ΩT /S , the sheaf of Kähler differentials. We
G
denote by jf the order function of Jacf on J∞ T , (J∞ T )/G or J∞
T if T has a faithful action of a finite group G. The ambiguity of the domain will not cause a confusion.
Remark 5.2. When T is smooth, the function jf on J∞ T coincides with the Jacobian order function, denoted by ord jacf , in [Seb04] and mentioned in the proof of Proposition
2.6.

WILDER MCKAY CORRESPONDENCES

18

Conjecture 5.3. Let the assumption be as in Section 4. Let Φ : J∞ X ⊃ A → R ∪
|F |
G,E
{∞} be a measurable function with A ⊂ p∞ (J∞
V ) and let p(∞) be the natural map
(J∞ V |F | )/CG (H) → J∞ X. We have
ˆ
ˆ
−j
|F |
Φ dµJ∞ X =
(Φ ◦ p(∞) )L p|F | dµ(J∞ V |F | )/CG (H) .
|F |

(p(∞) )−1 (A)

A

This conjecture would be proved by using existing techniques and arguments from
[DL02] and [Seb04].
Definition 5.4 ([Yasb].). For E ∈ G-Cov(D) with a connected component F , we define the weights of E and F with respect to V as wV (E) = wV (F ) := codim((V0 )H , V0 ) − vV (E)
with
1
HomOD (M, OE )
1
HomOD (M, OF )
· length
=
· length
.
♯G
OE · ΞF
♯H
OF · ΞF
For the generalization to the case where k is only perfect, see [WYa].
vV (E) = vV (F ) :=

The definition above gives the weight function, wV : G-Cov(D) →

1
Z.
♯G

We will denote the composition
G
J∞
V → G-Cov(D) →

1
Z
♯G

again by wV .
Definition 5.5. For an ideal I ⊂ OV stable under the G-action and a G-arc γ : E → V , we define a function
1
G
ord I : J∞
V →
Z ∪ {∞}
♯G
by
OE
1
OF
1
length −1 =
length
.
(ord I)(γ) :=
♯G
γ I
♯H
(γ|F )−1 I
We then extend this to G-stable fractional ideals and G-stable Q-linear combinations of closed subschemes as in Definition 2.2.
The conjectural change of variables formula is stated as follows:
Conjecture 5.6 ([Yasb]). For a measurable function Φ : J∞ X ⊃ C → R ∪ {∞}, we have
ˆ
ˆ
Φ dµJ∞ X =
(Φ ◦ p∞ )L−jp +wV dµJ∞
GV .
C

p−1
∞ (C)

To explain where the formula comes from, we first show a lemma:
Lemma 5.7. We have
♯H·vV (F )

JacV hF i /V ×D F = mF

OV hF i .

WILDER MCKAY CORRESPONDENCES

19

Proof. Let u′ : V hF i → V ×D F be the natural map. We have the standard exact sequence
(u′ )∗ ΩV ×D F/F → ΩV hF i /F → ΩV hF i /V ×D F → 0.
The left map is identical to the map
M ⊗OD OV hF i → M hF i ⊗OF OV hF i .
Since the Fitting ideal is compatible with base changes (for instance, see [Eis95, Cor.
20.5]), if I denotes the 0th Fitting ideal of

coker M ⊗OD OF → M hF i ,
♯H·v (F )

we have JacV hF i /V ×D F = I · OV hF i . It is now easy to see that I = mF V , for instance, by considering a triangular matrix representing the map M ⊗OD OF → M hF i for suitable bases.

Conjecture 5.6 can be guessed from the following conjecture:
G
Conjecture 5.8. For γ ∈ J∞
V and n ≫ 0, the fiber of the map
G
pn : πn (J∞
V ) → Jn X

over the image of γ is homeomorphic to a quotient of the affine space
(j −wV )(γ)

Ak p by a linear finite group action.

To see this, we first note that since two G-arcs E → V and E ′ → V with E ∼
6= E ′
G,E
have distinct images in Jn X for n ≫ 0, we can focus on J∞ V for fixed E. Fixing a
G-arc γ : E → V , we consider the map
(Jn V |F |)/CG (H) → Jn X.
The fiber of this map over the image of γ should be homeomorphic to j |F | (γ ′ )

Akp

/A,

where γ ′ is an arc of V |F | corresponding to γ and A is a certain subgroup of CG (H)
acting linearly on the affine space. This fact would be proved in the course of proving
Conjecture 5.3. On the other hand, the map

hF i
G
πn HomH
) /CG (H) → πn (J∞
V)
F (F, V
induced by u has fibers homeomorphic to

codim((V0 )H ,V0 )

Ak

/B

for some finite group B, which can be seen by looking at diagrams (4.2)-(4.4). From
Lemma 5.7, jp|F | − codim((V0 )H , V0 )
= jV hF i /X×D F − codim((V0 )H , V0 )
= (jV ×D F/X×D F + jV hF i /V ×D F ) − codim((V0 )H , V0 )
= jp − wV ,

WILDER MCKAY CORRESPONDENCES

20

concluding Conjecture 5.8.
6. The McKay correspondence for linear actions
To state the McKay correspondence conjecture for linear actions, we first define the notion of orbifold stringy motifs. Keeping the notation from the last section, let X, V,
VhF i and V|F | be centered log structures on X, V , V hF i and V |F | respectively so that the following morphisms are all crepant:
VhF i❋

④
④④
④④
④
④
}④
④
V❉
❉❉
❉❉
❉❉
❉❉
"

❋❋
❋❋
❋❋
❋#

X

V|F |

✇
✇✇
✇✇
✇
✇✇
{✇
✇

Since X is Q-factorial, either X or V determines the other centered log structures. The centered log structure V is G-equivariant and V|F | CG (H)-equivariant.
Definition 6.1. We define the orbifold stringy motif of the centered log G-D-variety
V to be
ˆ
MstG (V) :=

GV
J∞

LfV +wV dµG
V.

Note that since V is smooth over D, we have fV = ord ∆ for the boundary ∆ of V.
Arguments as in the proof of Proposition 2.6 deduce the following conjecture from
Conjecture 5.6:
Conjecture 6.2 (The motivic McKay correspondence for linear actions I). We have
Mst (X) = MstG (V).
We will next formulate a conjecture presented in a slightly different way so that we will be able to generalize it to the non-linear case easily.
Definition 6.3. For E ∈ G-Cov(D), we define the E-parts of MstG (V) and Mst (X)
respectively by
ˆ
G,E
Mst (V) : =
LfV +wV dµJ∞
G V and
G,E
J∞ V
ˆ
E
Mst (X) :=
LfX dµJ∞ X .
G,E
p∞ (J∞
V)

By the same reasoning as the one for the last conjecture, we would have
(6.1)

MstG,E (V) = MstE (X).

On the other hand, from Conjecture 5.3, we would have
(6.2)

MstE (X) = Mst,CG (H) (V|F | ).

WILDER MCKAY CORRESPONDENCES

21

F
Let G-Cov(D) = ∞
i=0 Ai be a conjectural stratification with finite dimensional strata
Ai (see Remark 3.6). The author [Yasb] conjectures also that each stratum Ai may not be of finite type over k, but the limit of a family f1

f2

X1 −
→ X2 −
→ ···
such that Xj are of finite type and fi are homeomorphisms.
Fn We then define a constructible subset of G-Cov(D) as a constructible subset of i=0 Ai for some n < ∞, which would be well-defined thanks to this conjecture. For a constructible subset C
of G-Cov(D), its class [C] in R is well-defined. Let τ denote the tautological motivic measure on G-Cov(D) given by τ (C) := [C] for a constructible subset C. If a function
Φ : G-Cov(D) → R ∪ {∞} is constructible, that is,´its image is countable and all fibers
Φ−1 (a), a ∈ R are constructible, then the integral G-Cov(D) Φ dτ is defined by
ˆ
X
Φ dτ :=
τ (Φ−1 (a)) · a ∈ R ∪ {∞}.
G-Cov(D)

a∈R

From Conjecture 6.2 and conjectural equations (6.1) and (6.2), it seems natural to expect
ˆ
G
Mst (V) =
Mst,CG (H) (V|F |) dτ
G-Cov(D)

and hence:
Conjecture 6.4 (The motivic McKay correspondence for linear actions II). We have
ˆ
Mst (X) =
Mst,CG (H) (V|F | ) dτ.
G-Cov(D)

This formulation of the McKay correspondence is what we will generalize to the non-linear case.
To make this conjecture more meaningful, it would be nice if we can compute
Mst,CG (H) (V|F | ) explicitly. For this purpose, next we see how to determine the centered log structures VhF i and V|F | from V. Let us write V = (V, ∆, W ), VhF i =
(V, ∆hF i , W hF i ) and V|F | = (V, ∆|F | , W |F |). The centers W hF i and W |F | are simply determined by
W hF i = u−1 (W ) and W |F | = r(W hF i ).
The boundaries ∆hF i and ∆|F | are determined as follows: hF i

Lemma 6.5. Regarding V0

|F |

and V0

prime divisors on V hF i and V |F | , we have hF i

∆hF i = u∗∆ − (♯H · vV (E) + dF/D ) · V0
1
|F |
· r∗ u∗ ∆ − vV (E) · V0 .
∆|F | =
♯H

d

F /D
Here dF/D is the different exponent of F/D, characterized by ΩF/D ∼
= OF /mF .

WILDER MCKAY CORRESPONDENCES

22

Proof. For the first equality, we have u∗ (KV + ∆)
= KV hF i − KV hF i /V + u∗ ∆
= KV hF i − KV hF i /V ×D F − (u′ )∗ KV ×D F/V + u∗ ∆.
Here KV hF i is the canonical divisor of V hF i as a D-variety rather than a F -variety and u′ denotes the natural morphism V hF i → V ×D F . From Lemma 5.7, hF i

KV hF i /V ×D F = ♯H · vV (E) · V0

.

Since (u′ )∗ KV ×D F/V is the pull-back of KF/D , we have hF i

(u′ )∗ KV ×D F/V = dF/D · V0

.

These equalities show the first equality of the lemma.
The second one follows from
1
|F |
· r∗ u∗ ∆ − v(E) · V0 )
r ∗ (KV |F | +
♯H
hF i

= KV hF i − KV hF i /V |F | + u∗ ∆ − ♯H · vV (E) · V0
hF i

= KV hF i + u∗ ∆ − (♯H · vV (E) + dF/E ) · V0
= KV hF i + ∆hF i .


Example 6.6. Suppose that ∆ = 0 and W = {o} with o ∈ V0 the origin. Then
|F |
codim((V0 )H ,V0 )
∆|F | = −vV (E) · V0 and W |F | ∼
. Hence
= Ak
MstG,E (V) = Mst,CG (H) (V|F | ) = LwV (E) .
Conjecture 6.4 is reduced to the form,
(6.3)

Mst (X) =

ˆ

LwV dτ.

G-Cov(D)

If p : V → X is étale in codimension one and if we denote p(o) again by o, then
Mst (X) = Mst (X)o and the last equality is exactly what was conjectured in [Yasb].
Remark 6.7. If ♯G is prime to the characteristic of k, then G-Cov(D) is identified with the set of conjugacy classes of G, denoted by Conj(G). Equality (6.3) in the last example is then written as
X
Mst (X) =
LwV (g) .
[g]∈Conj(G)

Expressing the weights wV (g) in terms of eigenvalues, we recover results by Batyrev
[Bat99], and Denef and Loeser [DL02].

WILDER MCKAY CORRESPONDENCES

23

7. The McKay correspondence for non-linear actions
In this section, we generalize Conjecture 6.4 to the non-linear case. It is rather easy, once we have formulated the conjecture as it is.
Let us consider an affine D-variety v = Spec Ov endowed with a faithful G-action.
We fix a G-equivarint (locally closed) immersion v ֒→ V
into an affine space V ∼
= AdD endowed with a linear G-action. Identifying G-arcs of v
G
G
with those of V factoring through v, we regard J∞
v as a subset of J∞
V.
Remark 7.1. Such an immersion
S always exists. Indeed, let f1 , . . . , fn be generators of
Ov as an OD -algebra, let A := i fi G, the union of their orbits, and let OD [xf | f ∈ A]
be the polynomial ring with variables xf , f ∈ A over OD . The ring has a natural
G-action by permutations of variables. The OD -algebra homomorphism
OD [xf | f ∈ A] → Ov , xf 7→ f defines a desired immersion. Moreover this construction gives a closed immersion into
V on which G acts by permutations. In this case, our weigh function wV is closely related to the Artin and Swan conductors [WYa], although we do not use this fact in this paper.
Definition 7.2. For E ∈ G-Cov(D) with a connected component F , we define the preuntwisting variety of v, denoted by vhF i , as the irreducible component of r −1 (v) ⊂ V hF i which dominates v. We then define the untwisting variety, denoted by v|F | , as the image of vhF i in V |F |. We also define the normalized pre-untwisting vhF i,ν and untwisting varieties v|F |,ν to be the normalizations of vhF i and v|F | respectively.
Let x := v/G. The following diagram shows natural morphisms of relevant varieties and symbols t, s and q denote morphisms as indicated:
(7.1)

vhF i,ν❍

❍❍
✌
❍❍ s
✌✌
❍❍
✌
❍❍
✌

#
✌
✌
t ✌✌ vhF i v|F |,ν
✌ ④ ❍❍
✌
❍❍
✌ ④④
❍❍
✌✌ ④④④
❍❍
✌
④
❍# 
✌
✌}④④
v ❊❊
v|F |
✉
❊❊
✉
✉✉
❊❊
✉✉
q ❊❊
❊" z✉✉✉✉

x

The one-to-one correspondence obtained in the last section
G,E
J∞
V ↔ (J∞ V |F | )/CG (H)

induces a one-to-one correspondence
G,E
J∞
v ↔ (J∞ v|F | )/CG (H).

WILDER MCKAY CORRESPONDENCES

24

We obtain the following diagram:
(J∞ v|F |,ν )/CG (H)

/ (J∞ v|F | )/CG (H)
♦♦
♦♦♦
♦
♦
♦♦
w♦♦♦

1-to-1

G,E o
J∞
v
●

●●
●●
●●
●●
#

J∞ x

E
G,E
E
If we put J∞
x to be the image of J∞
v in J∞ x, then we can naturally expect that J∞
x
|F |
|F |,ν
coincides with the images of J∞ v and J∞ v modulo measure zero subsets.
From now on, we suppose that v is normal. Let v, vhF i,ν , v|F |,ν and x be centered log structures on v, vhF i,ν and v|F |,ν respectively such that the morphisms

vhF i,ν❍

④
④④
④
④
④④
}④
④
v ❉❉
❉❉
❉❉
❉❉
❉❉
!

❍❍
❍❍
❍❍
❍❍
#

x

v|F |,ν

✈
✈✈
✈✈
✈
✈
✈✈
z✈✈

are all crepant. The centered log D-varieties v and v|F |,ν are G- and CG (H)-equivariant respectively. If we define the E-part MstE (x) of Mst (x), we can expect
MstE (x) = Mst,CG (H) (v|F |,ν )
similarly to the linear case. For the equality is a slight generalization of Conjecture
2.5 and would follow from the change of variables formula generalized along the line of
[DL02], applied to the almost bijection
E
J∞ v|F |,ν → J∞
x.

It is then natural to expect:
Conjecture 7.3 (The McKay correspondence for non-linear actions). We have
ˆ
Mst (x) =
Mst,CG (H) (v|F |,ν ) dτ.
G-Cov(D)

Definition 7.4. We define the E-part of the orbifold stringy motif of v as
MstG,E (v) := Mst,CG (H) (v|F |,ν )
and the orbifold stringy motif of v as
MstG (v) :=

ˆ

MstG,E (v) dτ.
G-Cov(D)

With this definition, the last conjecture simply says
Mst (x) = MstG (v).

WILDER MCKAY CORRESPONDENCES

25

Remark 7.5. The reader may wonder why we do not define MstG (v) as a motivic integral
G
on J∞
v, which appears more natural. It is because the author does not know whether
G
one can define a motivic measure on J∞
v. For, he does not know how to compute dimensions of fibers of hF i,ν
G
πn (HomH
))/CG (H) → πn (J∞
v).
F (F, v

Knowing it was, in the linear case, a key in formulating the change of variables formula
(Conjecture 5.6) and determining the integrand LfV +wV in the definition of MstG (V).
8. Computing boundaries of untwisting varieties
To compute examples of the wild McKay correspondence, we need to determine centered log varieties v|F |,ν . It is easy to determine the center. In this section, supposing v and v|F | are both normal and complete intersections in V and V |F | respectively, we compute the boundary of v|F | .
Let t : vhF i → v and s : vhF i → v|F |
be the natural morphisms, although they are different from the morphisms denoted by the same symbols in diagram (7.1) unless vhF i is also normal. The subvariety vhF i ⊂ V hF i is a complete intersection. To see this, that if s−1 (v|F | ) denotes the scheme firsthFnote
−1 |F |
i theoretic preimage, then s (v ) red = v . The subscheme s−1 (v|F | ) ⊂ V hF i is a complete intersection, hence Cohen-Macaulay, and generically reduced. From [Eis95,
THeorem 18.15], s−1 (v|F |) is actually reduced and s−1 (v|F | ) = vhF i .
I

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
