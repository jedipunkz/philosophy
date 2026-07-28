---
source: "https://arxiv.org/abs/1903.02782v1"
title: "Family size decomposition of genealogical trees"
author: "Max Grieshammer"
year: "2019"
publication: "arXiv preprint / math.PR"
download: "https://arxiv.org/pdf/1903.02782v1"
pdf: "https://arxiv.org/pdf/1903.02782v1"
captured_at: "2026-07-28T08:44:35Z"
updated_at: "2026-07-28T08:44:35Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche genealogy of morals"
tags:
  - "近代哲学"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Family size decomposition of genealogical trees

- 著者: Max Grieshammer
- 年: 2019
- 掲載情報: arXiv preprint / math.PR
- 情報源: [arxiv](https://arxiv.org/abs/1903.02782v1)
- ダウンロード: https://arxiv.org/pdf/1903.02782v1
- PDF: https://arxiv.org/pdf/1903.02782v1

## Obsidian Links

- 研究動向: [[ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代哲学]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代哲学 #実存主義 #ニヒリズム

## Abstract

We study the path of family size decompositions of varying depth of genealogical trees. We prove that this decomposition as a function on (equivalence classes of) ultra-metric measure spaces to the Skorohod space describing the family sizes at different depths is perfect onto its image, i.e. there is a suitable topology such that this map is continuous closed surjective and pre-images of compact sets are compact. We also specify a (dense) subset so that the restriction of the function to this subspace is a homeomorphism. This property allows us to argue that the whole genealogy of a Fleming-Viot process with mutation and selection as well as the genealogy in a Feller branching population can be reconstructed by the genealogical distance of two randomly chosen individuals.

## PDF Text

arXiv:1903.02782v1 [math.PR] 7 Mar 2019

Family size decomposition of genealogical trees
Max Grieshammer∗
March 8, 2019

Abstract
We study the path of family size decompositions of varying depth of genealogical trees. We prove that this decomposition as a function on
(equivalence classes of) ultra-metric measure spaces to the Skorohod space describing the family sizes at different depths is perfect onto its image, i.e. there is a suitable topology such that this map is continuous closed surjective and pre-images of compact sets are compact. We also specify a (dense) subset so that the restriction of the function to this subspace is a homeomorphism. This property allows us to argue that the whole genealogy of a Fleming-Viot process with mutation and selection as well as the genealogy in a Feller branching population can be reconstructed by the genealogical distance of two randomly chosen individuals.

Keywords: Genealogical distance, (ultra-)metric measure spaces, mass coalescent, family size decomposition.
AMS 2010 Subject Classification: Primary: 60G07, 54C10; Secondary:
60J25, 60G12, 92D15;

∗

Institute for Mathematics, Friedrich-Alexander Universität Erlangen-Nürnberg, Germany; max.grieshammer@math.uni-erlangen.de, MG was supported by DFG-Grant GR876-17.1 of A. Greven

Contents
1 Introduction

3

2 Metric measure spaces and the Gromov weak atomic topology
8
3 Family size decomposition of ultra-metric measure spaces 12
3.1 Definitions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 12
3.2 Results . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 14
4 Application to the tree-valued Fleming-Viot process
18
4.1 Definition of the neutral model . . . . . . . . . . . . . . . . . 18
4.2 Results for the tree-valued Fleming-Viot process . . . . . . . 20
5 Preparations for the proofs
21
5.1 Bounds for the Gromov-Prohorov metric and the function Φ . 21
5.2 Concatenation of trees . . . . . . . . . . . . . . . . . . . . . . 24
6 A short note on the weak atomic topology

26

7 Proof of Theorem 2.12 (a), (b)

28

8 Proofs for Section 3
29
8.1 Proof of Lemma 3.1 . . . . . . . . . . . . . . . . . . . . . . . 29
8.2 A result on relative compactness and proof of Lemma 3.8 and
Proposition 3.13 . . . . . . . . . . . . . . . . . . . . . . . . . 30
8.3 Proof of Theorem 3.11 and Theorem 3.17 . . . . . . . . . . . 39
8.4 Proof of Proposition 3.16 . . . . . . . . . . . . . . . . . . . . 48
9 Proof of Theorem 2.12 (c)

49

10 Proofs for Section 4.2
50
10.1 Proof of Proposition 4.1 . . . . . . . . . . . . . . . . . . . . . 50
10.1.1 Connection to the Kingman-coalescent . . . . . . . . . 51
10.1.2 Proof of Lemma 10.1 . . . . . . . . . . . . . . . . . . . 52
10.2 Proof of Theorem 4.3 . . . . . . . . . . . . . . . . . . . . . . . 54
A Skorohod topology

54

3

1

Introduction

There are several approaches to study genealogical properties of a WrightFisher population. For example, one can use the Kingman coalescent (see
[Kin82]) which naturally generates the genealogical tree of a neutral Moran model or a neutral Wright-Fisher population at a fixed time. When selection is present, things get harder. This is because in contrast to the neutral model, the tree of the current population now depends on the whole type evolution until the present time (see Theorem 2 in [DGP12]). Nevertheless, one can use, for example, the ancestral selection graph introduced by Krone and Neuhauser (see [KN97] and [NK97]) or the lookdown construction introduced by Donelly and Kurtz (see [DK96] and [DK99]), to construct or read off genealogical properties. But still, it is quite hard to get explicit results on the genealogy.
Depperschmidt, Greven, Pfaffelhuber, Winter [GPW09], (compare also
[GPW13]; and [DGP12] for a survey) followed a different approach. They constructed the genealogy dynamically as a tree-valued process (we explain the notion “tree” in more details below). This has the advantage, that one can use the generator of this Markov process to get, for example, recurrence relations of the genealogical distance of two randomly chosen individuals in equilibrium.
Here we try to connect the “classical” approach of coalescent models and the setting in Depperschmidt, Greven, Pfaffelhuber and Winter. Namely we ask for a quantity that both, (exchangeable) coalescents and (genealogical)
trees, have in common.
Recall, that one can use the Kingman coalescent to construct the genealogy in the neutral case and that the law of the Kingman coalescent is determined by the law of its block frequencies via Kingman’s paint box construction (see for example [Ber06]). The process associated with the evolution of the block frequencies is usually called mass coalescent. We will show that a tree naturally contains the concept of “mass coalescent”, which we call path of family size decompositions (or family size decomposition for simplicity) in this context. Our goal in this paper is to study properties of this decomposition and how to apply these results to (large) Wright-Fisher populations and Feller branching populations to gain information about genealogy.
To get a bit more precise (see Section 2 and Section 3 for all details), we consider the space U of (equivalence classes of) ultra-metric measure spaces

4

(X, r, µ), where we interpret X as a set of individuals, r as a genealogical distance and µ as a sampling measure. We assume that (X, r) is complete and separable and note that one can map (X, r) isometrically to the leaves of a rooted R-tree (see Remark 2.7 in [DGP12] and Remark 2.2 in [DGP11]), which justifies the name tree for an ultra-metric measure space.
Now, we decompose X into balls B̄h (x) := {y ∈ X : r(x, y) ≤ h} for some h > 0 and note that B̄h (x) ∩ B̄h (y) = ∅ for r(x, y) > h, since r is an ultra-metric, and the number of balls needed to cover X is countable, since X
separable. We can interpret B̄h (x) as a family descending from an ancestor who lived at the time h (measured backwards) (see Figure 1). When we now calculate the sizes, µ(Bh (x)), of the different families and denote the size ordered vector by a(h) := (a1 (h), a2 (h), . . .), i.e. ak (h) ≥ ak+1 (h) for all k, then we finally get the notion of family size decomposition of trees:
We call the function F : U → D((0, ∞), S ↓ ) that maps an ultra-metric measure space to the (cadlag) function h 7→ a(h) family size decomposition.
Here,
(
S ↓ :=

(x1 , x2 , . . .) ∈ [0, ∞)N :

)
X

xi < ∞, x1 ≥ x2 ≥ . . .

(1.1)

i∈N

h

Figure 1: On the left side we draw an ultra-metric measure space (X, r, µ), where |X| = 7
and µ({x}) = 1 for all x ∈ X. We can decompose this tree into four disjoint (closed) balls of radius h > 0 (drawn on the right side).

Our first goal in this paper is to study properties of this map and we can summarize our results as follows:
The function F is continuous, closed and preimages of compact sets are compact, i.e. F is perfect.

5

Of course, we have to define suitable topologies on the respective spaces
U and D((0, ∞), S ↓ ), so that the result is valid. While we equip the space of cadlag functions h 7→ a(h) with the Skorohod topology, we need to specify the topology on U. Typically, one would equip U with the so called Gromovweak topology. Convergence in this topology is equivalent to convergence of the corresponding distance matrix distributions ν m,(X,r,µ) , where
ν m,(X,r,µ) (·) := µ⊗m ({x1 , . . . , xm : (r(xi , xj ))1≤i<j≤m ∈ ·})

(1.2)

(see Section 2 for details) but we point out that the function F is not continuous in this topology. This is the reason why we need to introduce a finer topology which we call Gromov-weak atomic topology.
Convergence of a sequence un ∈ U, n ∈ N to a limit object u ∈ U in this new topology is equivalent to convergence of the distance matrix distribution, i.e. convergence in the Gromov-weak topology, and convergence of the following quantities
P
(a) (ν 2,un )∗ := h≥0 ν 2,un ({h})2 δh ⇒ (ν 2,u )∗ as n → ∞, where “⇒” denotes the convergence in the weak topology on finite measures.
(b) ν 2,un ({0}) → ν 2,u ({0}) as n → ∞.
The following example shows the differences between the convergence in the Gromov-weak and Gromov-weak atomic topology (see Section 2 for all details).
Example 1.1. (Convergence in the Gromov-weak atomic topology) We consider the sequence ({x1 , x2 , x3 }, rn , δx1 + δx2 + δx3 ), with rn (x1 , x2 ) = 1, rn (x1 , x3 ) = rn (x2 , x3 ) = 1 +

1
, n

(1.3)

rn (xi , xi ) = 0, i = 1, 2, 3
for n ≥ 1 and the ultra-metric measure space ({x1 , x2 , x3 }, r, δx1 + δx2 + δx3 ), where r(xi , xj ) = 1, i 6= j, r(xi , xj ) = 0, i = j.
(1.4)
then un → u in the Gromov-weak topology (see Figure 2) but un 6→ u not in the Gromov-weak atomic topology.

♣

6

n→∞



1
n

Figure 2: Convergence in the Gromov-weak but not in the Gromov-weak atomic topology.

In view of the above example we can interpret convergence in the Gromovweak atomic topology as convergence in the Gromov-weak topology plus some additional conditions on the convergence of the “branching points of the tree”.
Even though an ultra-metric measure space (X, r, µ) can not be reconstructed by the value F((X, r, µ)) in general (since F is not injective), we may hope that, as in the case of the Kingman coalescent, we can find a
“nice” subspace on which a reconstruction is possible. Indeed we prove that there is a dense Gδ subset of U such that F restricted to this subspace is a homeomorphism (onto its image)
and call elements u of this subspace identifiable by family sizes.
Next we want to apply our results on the tree-valued Feller diffusion and the tree-valued Fleming-Viot process and note that even though both processes are different, from a genealogical perspective they look quite similar: One can show that conditioned on the total mass, the genealogy in a
Feller diffusion is a time inhomogeneous tree-valued Fleming-Viot process.
We do not want to go into detail but refer to [GGR16], [Glö13], Chapter 5
or[GD18]. The important observation is, that it is enough to consider the tree-valued Fleming-Viot process to gain genealogical information for the tree-valued Feller diffusion and vice versa. Having that in mind, we get the following result:
The genealogy, denoted by Ut (as an U-valued random variable), in a
(large) Wright-Fisher population (with or without selection) at time t > 0 is completely determined by the distribution of the distance of two randomly chosen individuals, i.e. L(ν 2,Ut ).

7

Note that due to the above observation, this result stays valid in the situation of a tree-valued Feller diffusion.
In fact, we think that in the ”most” infinite population models the genealogy at a given time t is identifiable by family sizes, where we have processes in mind that arise as large population limits of graphically constructed finite population models equipped with the uniform distribution on the set of individuals. The reason is Proposition 3.16 that mainly says that, whenever the vector of family sizes at some depth is absolutely continuous to the (product-)Lebesgue measure, we can reconstruct the whole genealogy by the path of family size decompositions of varying depth.
We also point out that the above result is not a probabilistic result in the sense that we use probabilistic arguments to reconstruct the law of the genealogy, but rather a result about the “states” of the genealogy, i.e. the genealogy realizes its values (with probability one) in a subspace of U on which the function ν 2,· or F is a homeomorphism (onto its image).
Although we think that the function F (and the concept of family size decompositions) is an interesting object itself it can also be used to construct compact subsets of U and therefore gives a tool to prove compact containment conditions for evolving genealogies (see Corollary 3.12). To be a bit more precise we have the following application in mind: (1) Construct a tree-valued process U N via a graphical construction. (2) Show that it has a large population limit, i.e. U N ⇒ U, when U is equipped with the
Gromov-weak atomic topology. (3) Use the continuous mapping theorem to deduce that F(U N ) ⇒ F(U) (see also Corollary 3.12). (4) Assume that the process U is again indexed, i.e. U = U M , for some M = 1, 2, . . . and we want to study the behavior of U M for M → ∞. The key idea in doing that is to observe that for an evolving genealogy U = (Ut )t≥0 we can not only consider F(Ut ) = (f(Ut , h))h≥0 for fixed t (which corresponds to a backward in time picture), but also (f(Ut , t + h))t≥0 for fixed h (which corresponds to a forward in time picture). It will turn out that a combination of both a backward as well as a forward in time picture can be used to get tightness of the evolving genealogies in terms of simpler processes. Roughly speaking, it is enough to prove convergence of ((f(UtM , t + hi ))t≥0 )i=1,...,M , M ∈ N to get tightness of U M . The above is not quite rigorous (and will be part of an upcoming paper) but should give a hint that it is important to understand the function F in order to get results on U-valued processes.
At this point we should note that the results presented here are generalizations and extensions of results in [Gri17], but for the sake of completeness we include all proofs needed for the results in this paper.

8

2

Metric measure spaces and the Gromov weak atomic topology

Here we give the definition and basic properties of (ultra-)metric measure spaces, the subspaces we are interested in and the Gromov-weak (see
[DGP11] and [GPW09]) and Gromov-weak atomic topology.
Recall that the support of a finite Borel measure µ, denoted by supp(µ), on some separable metric space (X, d) is defined as the smallest closed set
C with µ(X\C) = 0. Note that supp(µ) is also given as supp(µ) = {x ∈ X ∀ε > 0 : µ(Bε (x)) > 0},

(2.1)

where Bε (x) is the open ball of radius ε around x.
Definition 2.1. (Metric measure spaces) We call the triple (X, r, µ)
1. a metric measure space, short mm-space, if
(a) (X, r) is a complete separable metric space, where we assume that
X ⊂ R (one needs this to avoid set theoretic pathologies).
(b) µ ∈ Mf (X), i.e. µ is a finite measure on the Borel sets generated by r.
2. ultra-metric, if r(x1 , x2 ) ≤ r(x1 , x3 )∨r(x3 , x2 ) for µ-almost all x1 , x2 , x3 ,
3. compact, if supp(µ) is compact,
P
P
4. purely atomic, if x∈X µ({x}) = µ(X), and non atomic if x∈X µ̃({x}) =
0.
5. identifiable (by family sizes), if it is ultra-metric and
X
X
µ(B̄h (x)) 6=
µ(B̄h (x)), x∈A1h

(2.2)

x∈A2h

for all h > 0 and all measurable subsets A1h , A2h ⊂ supp(µ) with
(A1h )h := {y ∈ supp(µ) : ∃x ∈ A1h , r(x, y) ≤ h} =
6 (A2h )h

(2.3)

and x, y ∈ Aih , x 6= y implies r(x, y) > h, i = 1, 2.
6. (non simultaneous) binary if r(x1 , x2 ) = r(x3 , x4 ) implies either x1 =
x3 and x2 = x4 or x1 = x4 and x2 = x3 for µ almost all x1 , x2 , x3 , x4 .

9

We say that two mm-spaces (X, rX , µX ) and (Y, rY , µY ) are equivalent if there is a measure-preserving isometry between this spaces, i.e. a map ϕ : supp(µX ) → supp(µY ) with rX (x, y) = rY (ϕ(x), ϕ(y)), x, y ∈ supp(µX )
and µY = µX ◦ ϕ−1 . This property defines an equivalence relation, and we denote by [X, r, µ] the equivalence class of a mm-space (X, r, µ).
We define the following sets:
M := {[X, r, µ] : (X, r, µ) is a metric measure space} ,

(2.4)

U := {[X, r, µ] ∈ M : (X, r, µ) is an ultra-metric measure space} , (2.5)
Ua := {[X, r, µ] ∈ U : (X, r, µ) is purely atomic} , c

(2.6)

U := {[X, r, µ] ∈ U : (X, r, µ) is non-atomic} ,

(2.7)

Uc := {[X, r, µ] ∈ U : (X, r, µ) is compact} ,

(2.8)

I := {[X, r, µ] ∈ U : (X, r, µ) is identifiable} ,

(2.9)

B := {[X, r, µ] ∈ U : (X, r, µ) is binary} .

(2.10)

We also use combinations of the above spaces such as Uac := Ua ∩ Uc etc.
♦
We will typically use m and u for elements of M and U.
Remark 2.2. Clearly, the property of being measure preserving isometric is reflexiv and transitiv. To see that it is symmetric one can first show that the image ϕ(supp(µX )) is dense in supp(µY ) and then extend the inverse to a measure-preserving isometry supp(µY ) → supp(µX ). We may therefore assume w.l.o.g. that the measure-preserving isometries are surjective.
♣
Example 2.3. (Identifiable elements) Let (Ui )i∈N be independent R+ -valued random variables,Pwhich are all absolutely continuous to the Lebesgue measure and satisfy i Ui < ∞ almost surely. Let (N, r) be a complete ultrametric space, then
"
#
X
N, r,
Ui δi ∈ I, almost surely.
(2.11)
i∈N

♣
Definition 2.4. (Distance matrix distribution) Let k ∈ N≥2 , m = [X, r, µ] ∈
M and set
(
(k )
k,(X,r)
X k → R+2 ,
R
:
(2.12)
(xi )1≤i≤k 7→ (r(xi , xj ))1≤i<j≤k .

10

We define the distance matrix distribution of order k by:
 k 
()
k,m k,(X,r)
⊗k
ν
:= (R
)∗ µ ∈ Mf R+2 ,

(2.13)

(k )
where R+2 is equipped with the product topology. For k = 1 we define
ν 1,m := m := µ(X).

(2.14)
♦

Remark 2.5. Note that ν k,m in the above definition does not depend on the representative (X, r, µ) of m. In particular ν k,m is well defined for all k ∈ N.
♣
Definition 2.6. (Gromov-weak topology) Let m, m1 , m2 , . . . ∈ M. We say mn → m for n → ∞ in the Gromov-weak topology, if n→∞

ν k,mn =⇒ ν k,m
 k 
()
in the weak topology on Mf R+2
for all k ∈ N.

Remark 2.7. Since m =
Gromov-weak topology.

(2.15)
♦

p
ν 2,m (R+ ), we have m 7→ m is continuous in the
♣

For our results it will be necessary to introduce a finer topology:
Definition 2.8. (Gromov-weak atomic topology) Let u, u1 , u2 , . . . ∈ U. We say un → u for n → ∞ in the Gromov-weak atomic topology, if un → u for n → ∞ in the Gromov-weak topology and
P
a) (ν 2,un )∗ ⇒ (ν 2,u )∗ , where (ν 2,u )∗ = h≥0 ν 2,u ({h})2 δh .
b) ν 2,un ({0}) → ν 2,u ({0}).
♦

11

Remark 2.9. This topology is related to the so called weak atomic topology on finite measures, introduced by [EK94], where one says that a sequence µn ∈ Mf (X), n ∈ N of finite Borel-measures converges to a finite
Borel-measure µ ∈ Mf (X) in the weak atomicPtopology, when µn ⇒ µ (i.e.
convergence in the weak topology) and µ∗n := x∈X µn ({x})2 δx ⇒ µ∗ .
This explains the origin of the name “Gromov-weak atomic”.
♣
Example 2.10. (Convergence in the Gromov-weak atomic topology - Example 1.1 continued) Assume we are in the situation of Example 1.1, then un → u in the Gromov-weak topology. Note that
(ν 2,un )∗ = 32 δ0 + 22 δ1 + 42 δ1+ 1

n

(ν 2,u )∗ = 32 δ0 + 62 δ1 .

(2.16)

Hence
(ν 2,un )∗ ⇒ 32 δ0 + (22 + 42 )δ1 6= (ν 2,u )∗ .

(2.17)

This means un 6→ u in the Gromov-weak atomic topology.

♣

Now recall the definition of the Prohorov distance of two finite measures
µ1 and µ2 on a metric space (E, r) with Borel σ-field B(E)
n dPr (µ1 , µ2 ) := inf ε > 0 : µ1 (A) ≤ µ2 (Aε ) + ε, o (2.18)
µ2 (A) ≤ µ1 (Aε ) + ε for all A closed , where

n o
Aε := x ∈ E : r(x, x0 ) < ε, for some x0 ∈ A .

(2.19)

The next proposition summarizes some important facts about the Gromovweak topology (see [DGP11] and [LVW15] section 2.1; compare also [GPW09]).
Proposition 2.11. (Properties of the Gromov-weak topology) (a) M equipped with the Gromov-weak topology is Polish and the subspace U ⊂ M is closed.
(b) An example for a complete metric on M (respectively U) is the
Gromov-Prohorov metric dGPr , where for two mm-spaces [X, rX , µX ] and
[Y, rY , µY ]
dGPr ([X, rX , µX ], [Y, rY , µY ]) :=

inf
(ϕX ,ϕY ,Z)


(Z,r )
−1
dPr Z µX ◦ ϕ−1
X , µY ◦ ϕY ,
(2.20)

12

where the infimum is taken over all isometric embeddings ϕX and ϕY from supp(µX ) and supp(µY ) into some complete separable metric space (Z, rZ )
(Z,r )
and dPr Z denotes the Prohorov distance on Mf (Z).
We close this section with some properties of the Gromov-weak atomic topology:
Theorem 2.12. (Properties of the Gromov-weak atomic topology) If we equip U with the Gromov-weak atomic topology then the following holds (recall (2.9)):
(a) U is a Polish space,
(b) I ⊂ U is dense.
(c) Let Uc be equipped with the subspace topology, then B ⊂ Uc is closed.
Remark 2.13. (1) As in [EK94] (see the discussion after (2.3)), the Borel sets generated by the Gromov-weak topology coincide with the Borel-sets generated by the Gromov-weak atomic topology.
(2) Uc ⊂ U is measurable in the Gromov-weak topology and therefore measurable in the Gromov-weak atomic topology (see Remark 2.8 and Corollary 3.6 in [ALW16]).
♣

3

Family size decomposition of ultra-metric measure spaces

We will now introduce the function F that gives the size of the different families of an ultra-metric measure space u.

3.1

Definitions

We start with the following Lemma, that gives us the existence of an “almost surely” disjoint decomposition of an ultra-metric measure space into closed balls.
Lemma 3.1. Let 0 < h, u = [X, r, µ] ∈ U and B̄h (x) be the closed ball of radius ≤ h around x ∈ X. Then there is a n(h) ∈ N ∪ {∞} and a family
{rhi : i ∈ {1, 2, . . . , n(h)}} of elements of supp(µ) with

µ B̄(rhi , h) ∩ B̄(rhj , h) = 0,
(3.1)

13

for i 6= j and n(h)

µ(X) =

X


µ B̄(rhi , h) .

(3.2)

i=1

Moreover, if 0 < δ ≤ h, then there is a partition {Ii }i∈1,...,n(h) of {1, . . . , n(δ)}
such that
X
µ(B̄(rhi , h)) =
µ(B̄(rδj , δ)),
∀i = 1, . . . , n(h).
(3.3)
j∈Ii

Remark 3.2. (i) By the definition of the support we get µ(B̄(rhi , h)) > 0
for all i ∈ {1, . . . , n(h)}.
(ii) The analogue of Lemma 3.1 holds if we replace ≤ h by < h.
♣
Remark 3.3. Another important observation is the following: Given a finite ultra-metric space ({1, . . . , n(0)}, r) = (X, r), then the path of partitions h 7→ ({Iih }i=1,...,n(h) ) of {1, . . . , n(0)} contains all information of the metric r, i.e. the function that maps an ultra-metric r on X to the path of partitions is an injection and given an element (h 7→ π h ) contained in the range of this map, the corresponding metric r can be reconstructed by r(k, l) := inf{h > 0 : k, l ∈ πih for some i = 1, . . . , n(h)},

(3.4)

for k, l ∈ {1, . . . , n(0)}.

♣

Let C ≥ 0 and set
(
SC↓ :=

)

(x1 , x2 , . . .) ∈ [0, ∞)N :

X

xi ≤ C, x1 ≥ x2 ≥ . . . ,

(3.5)

i∈N

(
↓

S :=

)
N

(x1 , x2 , . . .) ∈ [0, ∞) :

X

xi < ∞, x1 ≥ x2 ≥ . . . .

(3.6)

i∈N

We consider the following two distances on SC↓ and S ↓ : d1 (x, y) =

∞
X

|xi − yi | = x − y 1

(3.7)

i=1

and d∞ (x, y) = max |xi − yi |.
i∈N

(3.8)

We note that SC↓ and S ↓ are typically equipped with the `1 -distance, d1 .

14

Definition 3.4. (Definition of f) Let u ∈ U. We define the map f(u, ·) :
(0, ∞) → S ↓ , f(u, h) = (a1 (h), a2 (h), . . .),
(3.9)
where the ak (h) are given by


n(h)


X
ak (h) = max c ≥ 0 :
1(µ(B̄(rhi , h)) ≥ c) ≥ k ,



k = 1, 2, . . . , n(h),

i=1

ak (h) = 0,

for k > n(h).
(3.10)

Note that ak (h) ≥ ak+1 (h) is the non-increasing reordering of the sequence
(µ(B̄(rhi , h)))i=1,...,n(h) .
♦
Remark 3.5.
(i) Let (X, rX , µX ) and (Y, rY , µY ) be two equivalent ultra-metric measure spaces and let ϕ : supp(µX ) → supp(µY ) be a measure preserving isometry. Then {rhi : i ∈ {1, 2, . . . , n(h)}} ⊂ supp(µX ) satisfies the conditions in Lemma 3.1 if and only if {ϕ(rhi ) : i ∈ {1, 2, . . . , n(h)}} ⊂
supp(µX ) satisfies the conditions.
(ii) If x ∈ supp(µX ) and h > 0, then there is exactly one i ∈ {1, . . . , n(h)}
with
µX (B̄(rhi , h)) = µX (B̄(rhi , h) ∩ B̄(x, h)) = µX (B̄(x, h)).

(3.11)

As a consequence, the definition of f does not depend on the representatives.
♣
Note that the domain of f(u, ·) is (0, ∞). In some cases it is also possible to add 0 to the domain and we close this section with the following remark:
Remark 3.6. In the case, where u ∈ Ua is purely atomic we can extend the function f(u, ·) to a function f̂(u, ·) : [0, ∞) → S ↓ .
♣

3.2

Results

We start with the following definition:

15

Definition 3.7. (Definition of F) We define
F : U → (S ↓ )(0,∞) ,

u 7→ f(u, ·).

(3.12)
♦

The first observation is, that F maps ultra-metric measure spaces to cadlag (i.e. right continuous with left limits) functions:
Lemma 3.8. F := F(U) ⊂ D((0, ∞), S ↓ ), where S ↓ is equipped with d1 .
In the following
D((0, ∞), S ↓ ) is always equipped with the Skorohod topology, given in Appendix A.
Now the question is whether F is continuous when U is equipped with the Gromov-weak topology.
Example 3.9. Assume we are in the situation of Example 1.1. Observe that if we take for example tn = 1 + n1 → 1 then f(un , tn ) ≡ (2, 1, 0, . . .)
6∈ {(1, 1, 1, 0 . . .) , (3, 0, 0, . . .)} = {f(u, 1), f(u, 1−)} ,

(3.13)

i.e. f(un , ·) 6→ f(u, ·) in the Skorohod topology (see Proposition 3.6.5 in
[EK86]).
♣
In other words we can not expect F to be continuous, when U is equipped with the Gromov-weak topology. But as we have seen in Example 1.1, the sequence un does not converge in the Gromov-weak atomic topology and in fact, this is the reason why we introduced this new topology.
Recall that a function f : X → Y between two topological spaces is called perfect, if it is continuous, surjective, closed (i.e. maps closed sets to closed sets) and f −1 ({y}) is compact in X for all y ∈ Y . We remark the following:
Remark 3.10. If X is a topological space and Y is a compactly generated
Hausdorff space (for example a metric space) and f : X → Y is surjective, then the following is equivalent (see for example [Pal70]):
(i) f is perfect,

16

(ii) f is continuous and proper, i.e. f −1 (K) is compact in X for all compact sets K ⊂ Y .
Note that a perfect map is also a quotient map, i.e. surjective and f −1 (U )
is open in X iff U is open in Y .
♣
Theorem 3.11. (Properties of F) Let U be equipped with the Gromov-weak atomic topology, then F : U → F has the following properties: i) F is perfect.
ii) The restriction F|I of F to I (see Definition 2.1) is a homeomorphism onto its image.
Recall that a collection of cadlag process {X n : n ∈ N} with values in some Polish space E satisfies a compact containment condition if for all
ε > 0 and T > 0 there is a compact set K ⊂ E such that inf P (X n (t) ∈ K ∀t ∈ [0, T ]) ≥ 1 − ε.

n∈N

(3.14)

Corollary 3.12. Let Un be a sequence in U and U be equipped with the
Gromov-weak atomic topology. Then
(i) (L(Un ))n∈N is tight if and only if (L(F(Un )))n∈N is tight.
(ii) Un ⇒ U for some U-valued random variable U implies F(Un ) ⇒ F(U).
Moreover, the map D([0, ∞), U) → D([0, ∞), F(U)), (ut )t≥0 7→ (F(ut ))t≥0 is continuous and a collection of U-valued cadlag processes {(Utn )t≥0 : n ∈ N}
satisfies a compact containment condition if and only if {(F(Utn ))t≥0 : n ∈
N} satisfies a compact containment condition.
Proof. This is a direct consequence of Theorem 3.11, the continuous mapping theorem, Prohorov’s theorem and the fact that continuous images of compact sets are compact. See Problem 3.13 in [EK86].
Another interesting observation is, that even though, the above result is related to the Gromov-weak atomic topology, we also get a result for the
Gromov-weak topology:
Proposition 3.13. Let U n , n = 1, 2, . . . be a sequence of U-valued random variables and let U be equipped with the Gromov-weak topology. Assume that for all δ > 0 and all ε > 0

17

(i) there is a compact set Γ ⊂ S ↓ such that lim sup P (f(Un , δ) ∈ Γc ) ≤ ε,

(3.15)

n→∞

(ii) there is an H ≥ 0 such that


!2
∞
∞
X
X
lim sup P 
f(U n , H)i −
f(U n , H)2i ≥ ε ≤ ε
n→∞

i=1

(3.16)

i=1
n

and that the total mass (ν 1,U )n∈N is tight. Then, (U n )n∈N is tight.
Remark 3.14. Note that SC↓ equipped with d∞ is a compact space (this follows analogue to Proposition 2.1. in [Ber06]). It is not hard to see that
Γ ⊂ SC↓ , equipped with d1 , is compact, if for all ε > 0 there is a M ∈ N such that
X
sup fi ≤ ε.
(3.17)
f ∈Γ i≥M

We will discuss this property in more detail in Section 8.2.

♣

Even though the function F is not injective on the whole space it is at least injective on a dense subset (see Theorem 2.12).
Remark 3.15. By Lavrentiev’s Theorem (see Section 35.II in [Kur14]) there are two Gδ sets I ⊂ I∗ ⊂ U and F(I) ⊂ I∗F ⊂ F(U) and a homeomorphism
F∗ : I∗ → I∗F extending F. In addition I∗ is dense in U since I is dense in U.
♣
Next, we give a criterion when a U-valued random variable takes values in I:
Proposition 3.16. Assume that U is an U-valued random variable. Let
Nh ∈ N ∪ {∞} be the number of non-zero entries of f(U, h). If L(f(U, h)) 
λ⊗Nh conditioned on Nh for all h > 0, where λ denotes the Lebesgue measure, then U ∈ I almost surely.
We close this section with a result that is even stronger than the above result, when one considers the subspace B ∩ I (see Definition 2.1):
Theorem 3.17. (Properties of B∩I) Let B∩I be equipped with the Gromovweak atomic topology, then we have V : B∩I → V(B∩I) ⊂ Mf (R+ ), u 7→ ν 2,u is a homeomorphism.

18

Remark 3.18. As in Remark 3.15 we can extend the homeomorphism to
Gδ subsets.
♣

4

Application to the tree-valued Fleming-Viot process

In this section we give a short introduction to tree-valued Fleming-Viot processes and show that these processes live in the subspace B ∩ I (see
Theorem 3.17). For simplicity, we will only introduce the neutral model and refer to [DGP12] for the general case.
In section 4.1 we define the neutral tree-valued Moran model of a given size N (the population size). This model was defined by [GPW13] and extended by [DGP12] to include selection and mutation. In section 4.2 we consider the large population limit (i.e. N → ∞) of the tree-valued Moran models, the so called tree-valued Fleming-Viot process, and give our main result for this process.

4.1

Definition of the neutral model

We want to describe the genealogy of a population, consisting of N ∈ N
individuals, that evolves according to the following dynamic:
Resampling: Every pair i 6= j of individuals is replaced with rate one. If such an event occurs, i is replaced by an offspring of j with probability 12 , or j is replaced by an offspring of i with probability 21 .
In order to describe the evolution of this process formally, let IN :=
[N ] := {1, . . . , N }, N ∈ N and
 i,j
η : i, j ∈ IN , i 6= j
(4.1)
be a realization of a family of independent rate 1 Poisson point processes.
For i, i0 ∈ IN , 0 ≤ h < t < ∞ we say that there is a path from
(i, h) to (i0 , t) if there is an n ∈ N, h ≤ t1 < t2 < · · · < tn ≤ t and j1 , . . . , jn ∈ IN such that for all k ∈ {1, . . . , n + 1} (j0 := i, jn+1 := i0 )
η jk−1 ,jk {tk } = 1, η x,jk−1 ((tk−1 , tk )) = 0 for all x ∈ IN .
Note that for all i ∈ IN and 0 ≤ h ≤ t there exists an unique element
Ah (i, t) ∈ IN

(4.2)

19

time t

h
1

2

3

4

Figure 3:

On the left side we see the graphical construction of the Moran model;
→ indicates a resampling event. On the right side we see the genealogical tree of the population at time t. In this case the ancestor of all individuals at time h would be individual 4, i.e. Ah (i, t) = 4 for all i = 1, . . . , 4

with the property that there is a path from (Ah (i, t), h) to (i, t). We call
Ah (i, t) the ancestor of (i, t) at time h (see Figure 3).
Let r0 be an ultra-metric on IN and i, j ∈ IN . Then we define the following (pseudo) ultra-metric on IN :
(
rt (i, j) :=

t − sup{h ∈ [0, t] : Ah (i, t) = Ah (j, t)},

if A0 (i, t) = A0 (j, t),

t + r0 (A0 (i, t), A0 (j, t)),

if A0 (i, t) 6= A0 (j, t).
(4.3)

Since rt , is only a pseudo-metric, we consider the following equivalence t := I /≈ the relation ≈t on IN : x ≈t y ⇔ rt (x, y) = 0. We denote by I˜N
t
N
set of equivalence classes and note that we can find a set of representatives t such that I¯t → I˜t , x → [x]
I¯N
≈t is a bijection.
N
N
N
Let µ ∈ M1 (IN ) be the uniform distribution on IN , i.e.
µN =

1 X
δk
N

(4.4)

k∈IN

and define r̄t (ī, j̄) = rt (ī, j̄),

N
µ̄N
t ({ī} × ·) = µt ([ī]≈t × ·),

t ī, j̄ ∈ I¯N
.

(4.5)

20

Then the tree-valued Moran model of size N is defined as t
UtN := [I¯N
, r̄t , µ̄N
t ],

4.2

(4.6)

Results for the tree-valued Fleming-Viot process

Assume that L(U0N ) ⇒ µ ∈ M1 (U), where U is equipped with the Gromovweak topology. Then
N →∞

(UtN )t≥0 ⇒ (Ut )t≥0

(4.7)

weakly in the Skorohod topology on D([0, ∞), U), where L(U0 ) = µ and
(Ut )t≥0 is the solution of a well-posed martingale problem (see Theorem 2
in [GPW13]). We call the process U = (Ut )t≥0 tree-valued Fleming-Viot process.
Proposition 4.1. (Convergence of the tree-valued Moran models) Let U be equipped with the Gromov-weak atomic topology and let (Vi )i∈N be a sequence of independent [0, 1]-uniformly distributed random variables. We assume that U0N = [[0, 1], r0 ,P
µN ], where ([0, 1], r0 ) is a compact binary ultra-metric space, and µN = N1 N
i=1 δVi . Then
UtN ⇒ Ut

for all t ≥ 0.

(4.8)

Remark 4.2. Even though we choose a special initial condition, the proof for general initial conditions should be similar but more technical (one needs to use Lemma 5.8 in [GPW09] for example). We also note that the initial condition does not really matter when one wants to study genealogical properties that are generated by an evolving population.
♣
We are now ready for our main result:
Theorem 4.3. (State space of tree-valued FV-processes) Recall Remark 2.13
and assume that P (U0 ∈ B ∩ Uc ) = 1, then
P (Ut ∈ B ∩ I ∩ Uc ) = 1,

∀t > 0.

(4.9)

Remark 4.4. Even though, it is not hard to see that B ∩ I is measurable we can also apply Remark 3.18 and replace B ∩ I in the above theorem by a suitable Gδ -set.
♣

21

Since we did not define the model with selection we need to refer all interested readers to [DGP12]. But, as a direct consequence of the Girsanov transform - Theorem 2 in this paper, one can prove the following.
Corollary 4.5. If we denote by U α the tree-valued Fleming-Viot process with mutation and selection parameter α ≥ 0, defined in [DGP12], with
P (U0α ∈ B ∩ Uc ) = 1, then
P (Utα ∈ B ∩ I ∩ Uc ) = 1,

5

∀t > 0.

(4.10)

Preparations for the proofs

We start with some preparations needed for the proofs of our results. In section 5.1 we prove some bounds for the Gromov-Prohorov metric and in section 5.2 we introduce the notion of concatenation of trees, which will be useful in order to prove the continuity of F.

5.1

Bounds for the Gromov-Prohorov metric and the function Φ

We start with the following observation.
Remark 5.1. Let (X, r, µ) and (X̃, r̃, µ̃) be two equivalent ultra-metric measure spaces. If we denote by {rhi : i = 1, . . . , n(h)} and {r̃hi : i =
1, . . . , ñ(h)} two families of representatives in the sense of Lemma 3.1, then it is not hard to see (see also Remark 3.5) that


X
{rhi : i ∈ {1, . . . , n(h)}}, r,
µ(B̄ r (rhi , h))δrh 
i

i∈{1,...,n(h)}


= {r̃hi : i ∈ {1, . . . , ñ(h)}}, r̃,


X

µ̃(B̄ r̃ (r̃hi , h))δr̃h 
i

i∈{1,...,ñ(h)}

(5.1)
and it is possible to define for h > 0
h i
Φ̂h (u) = {rhi : i ∈ {1, . . . , n(h)}}, r, µh

(5.2)

and i
h
Φh (u) = {rhi : i ∈ {1, . . . , n(h)}}, r − h · 1(rhi 6= rhj ), µh ,

(5.3)

22

where
µh :=

X



µ B̄(rhi , h) δrh .
i

(5.4)

i∈{1,...,n(h)}

♣
These functions will appear in several proofs. The reason is the following
Lemma:
Lemma 5.2. Let 0 < h and u = [X, r, µ] ∈ U.
(i) If A ⊂ X is measurable, and µA (·) := µ(· ∩ A) then dGPr ([A, r, µA ], [X, r, µ]) ≤ µ(X\A).

(5.5)

(ii) If u0 = [X, r, µ0 ] ∈ U, then dGPr (u, u0 ) ≤ dPr (µ, µ0 ),

(5.6)

where the Prohorov distance is taken on the set of Borel-measures on
X (see (2.18)).
(iii) Let Φh and Φ̂h be the functions from Remark 5.1. Then dGPr (u, Φ̂h (u)) ≤ h,

dGPr (Φh (u), Φ̂h (u)) ≤ h.

(5.7)

(iv) The functions h 7→ Φh (u) and h 7→ Φ̂h (u) as functions from (0, ∞) →
U are both cadlag.
Proof. (i) Note that the identity id : X → X is an isometric embedding from A to X. Using the definition of the Gromov-Prohorov metric from
Proposition 2.11, it is enough to bound (note that µA ≤ µ):

dP r (µA , µ) = inf  > 0 : µ(B) ≤µA (B ε ) + ε,
(5.8)
∀B ⊂ X Borel-measurable , where
B ε = {x ∈ X : ∃x0 ∈ B, r(x, x0 ) < ε}.

(5.9)

Note that if µ(X\A) = 0 then dP r (µA , µ) = 0 and if µ(X\A) > 0 we can take  = µ(X\A) and the result follows.
(ii) As in (i) one can use the identity as isometric embedding.

23

(iii) We use the notation of Remark 5.1 and note that id is an isometric embedding of {rhi , i ∈ N} in X. Define the measure µ̄ on X × X by
X
µ̄(A1 × A2 ) :=
µ(A1 ∩ B̄(rhi , h))δrh (A2 ).
(5.10)
i

i∈N

for all measurable sets A1 , A2 ⊂ X and observe that µ̄ is a coupling of
µ and µh . Since µh ({rhi , i ∈ N}) = µ(X) and by the definition of the
Gromov-Prohorov metric from Proposition 2.11 together with Theorem 3.1.2
in [EK86] (with the obvious extension to couplings of finite measures with the same mass), we get dGPr (u,Φ̂h (u)) ≤
n o

inf inf  > 0 : ν {(x, x0 ) ∈ X × X : r(x, x0 ) ≥ ε} ≤ ε ,

(5.11)

ν

where the infimum is taken over all couplings ν of µ and µh . It follows that dGPr (u, Φ̂h (u)) ≤ inf{ > 0 : µ̄({(x, x0 ) ∈ X × X : r(x, x0 ) ≥ ε}) ≤ ε}
(5.12)
and if we choose ε > h then


µ̄ (x, x0 ) ∈ X × X : r(x, x0 ) ≥ ε
X
≤
µ(B̄(rhi , h) ∩ B̄(rhi , h))δrh (B̄(rhj , h)) = 0.

(5.13)

i

i,j∈N, i6=j

For the second part, we use the same argument as in section 3 in [Loe13]: Let
Y := {rhi : i ∈ {1, . . . , n(h)}, r1 = r, r2 = r − h1(x 6= y) and µ1 = µ2 = µh .
We denote by Y ] Y the disjoint union of Y and Y and let ϕi : Y → Y ] Y
be the canonical embeddings, i = 1, 2. Define the metric d on Y ] Y by d(ϕ1 (x), ϕ1 (y)) = r1 (x, y),

(5.14)

2

d(ϕ2 (x), ϕ2 (y)) = r (x, y),
1

(5.15)
2

d(ϕ1 (x), ϕ2 (y)) = inf (r (x, z) + r (y, z)) + h, z∈Y

(5.16)

where x, y ∈ Y . Then, as in [Loe13] it is easy to see that this is a metric on
Y ] Y that extends the metrics r1 and r2 (i.e. ϕi is an isometry for i = 1, 2)
and we have h0
ϕ2 (ϕ−1
:= {x ∈ Y ] Y : ∃x0 ∈ F s.t. d(x, x0 ) < h0 },
1 (F )) ⊂ F

(5.17)

24

for all h0 > h. Since µ1 = µ2 this gives:
−1
−1
−1
−1
2
2
2
h0
µ1 ◦ ϕ−1
1 (F ) = µ ◦ ϕ1 (F ) ≤ µ ◦ ϕ2 (ϕ2 (ϕ1 (F )) ≤ µ ◦ ϕ2 (F ) + h0 ,
(5.18)
for all h0 > h and the result follows.

(iv) A similar argument as in (iii) shows that Φ̂h0 (u) → Φ̂h (u) for h0 ↓ h and by definition we have Φh+δ (u) = Φδ (Φh (u)) and hence by (iii) Φh0 (u) →
Φh (u) for h0 ↓ h. This shows the right continuity. For the existence of the left limits set i
h
(5.19)
uh := {rhi : i ∈ {1, . . . , n(h)}}, r − h · 1(τi 6= τj ), µ◦h , i
h
(5.20)
ûh := {rhi : i ∈ {1, . . . , n(h)}}, r, µ◦h , where
µ◦h (A) :=



µ B(rhi , h) δrh (A)

X

i

(5.21)

i∈{1,...,n(h)}

is given in terms of open balls with radius < h (instead of ≤ h - see Remark
3.2 (ii)). By a similar argument as in (iii) together with the fact that


h h 0
lim
µ
B(r
, h)
\
B̄(r
, h
)
= 0,
(5.22)
i i
0
h ↑h

we get dGPr (Φh0 (u), uh ) ∨ dGPr (Φ̂h0 (u), ûh ) → 0,

5.2

h0 ↑ h.

(5.23)

Concatenation of trees

We summarize some properties of the concatenation of trees given in [GGR16]
(see also [EM17]).
Definition 5.3. (Concatenation of trees) LetPh > 0 and ui = [Xi , ri , µi ], i ∈ I (I ⊂ N ∪ {∞}) be a sequence in U with i∈I µi (Xi ) ≤ C < ∞, ui := µi (Xi ) =

p
ν 2,ui [0, ∞)) > 0

(5.24)

and
ν 2,ui (h, ∞) = 0.

(5.25)

25

We define the concatenation: h
G

"
ui := ui1 th ui2 th . . . :=

U

h

X

µi ,

(5.26)

for x, y ∈ Xi , for x ∈ Xi , y ∈ Xj , i 6= j.

(5.27)

Xi , r ,

i∈I

i∈I

where

#
]

i∈I

i∈I Xi is the disjoint union of the Xi and h



r (x, y) =

ri (x, y), h,

♦
Definition 5.4. (h-top) In the sense of Lemma 3.1 we define for 0 < h and u = [X, r, µ] ∈ U the h-top Ψh (u) ∈ U as
Ψh (u) :=

h
G

h

B̄(τih , h), r, µ|B̄(τ h ,h)
i

i

(5.28)

i∈{1,...,n(h)}

where µ|B̄(τ h ,h) (·) = µ(· ∩ B̄(τih , h)). By Remark 3.5 this definition is indei pendent of the representative (X, r, µ).
♦
Remark 5.5. Note that Ψh ([X, r, µ]) = [X, rh , µ] with

r(x, y), if r(x, y) ≤ h, h
r (x, y) =
h, otherwise.

(5.29)
♣

Remark 5.6. (i) Let h > 0 and ui := [B̄(τih , h), r, µ|B̄(τ h ,h) ] ∈ U, then i
ν 2,ui ((h, ∞) = 0 for all i = 1, 2, . . . , n(h).
(ii) Let u ∈ U and ui given as in (i) for some h > 0. If x ∈ supp(µ), then there is a i, j ∈ {1, . . . , n(h)} such that
µ(B̄(x, h)) = ui = f (h, ui ) = f (h, u)j .

(5.30)

(iii) As in (ii) we get for h > 0: n(h)

u = µ(X) =

X

n(h)

f (h, u)i =

i=1

Note that u 7→ u is continuous, since u =

X

ui .

(5.31)

p
ν 2,u ([0, ∞)).

♣

i=1

26

Definition 5.7. (Concatenation as partial order) Define for 0 < h the
0
relation ≤h on U by saying u ≤h v if there is a u0 ∈ U with ν 2,u (h, ∞) = 0
such that Ψh (v) = Ψh (u) th u0 .
♦
Lemma 5.8. Let 0 < h, u, v ∈ U, (un ), (vn ) be two sequences in U and U
be equipped with the Gromov-weak topology.
(i) Suppose that Ψh (un ) → u and χn := Ψh (un ) th Ψh (vn ) → χ, for n → ∞. Then u ≤h χ.
(ii) If dGPr (un , u) → 0, then Ψh (un ) → Ψh (u) for all h > 0, i.e. u 7→ Ψh (u)
is continuous. Moreover, if hn → h, then Ψhn (u) → Ψh (u).
(iii) If un → u, for n → ∞, and vn ≤h un , for all n ∈ N, then {Ψh (vn ) : n ∈ N} is compact.
(iv) Assume we are in the situation of Remark 5.6 for some h > 0, then
ν 2,ui ≤ ν 2,u for all i = 1, . . . , n(h).
(v) Let un , vn ∈ U such that un → u ∈ U and vn → v ∈ U and assume that un , u, vn , v satisfy (5.25), then un th vn → u th v.
Proof. This is a summary of Proposition 2.17, Lemma 3.3 and Lemma 3.5
in [GGR16].

6

A short note on the weak atomic topology

Here we give a short introduction to the weak atomic topology (see [EK94])
and prove a Proposition that gives a characterization of convergence in this topology in terms of cumulative distribution functions.
Definition 6.1. (Weak atomic topology) Let (E, r) be a complete separable metric space and µ1 , µ2 , . . . ∈ Mf (E) (space of finite Borel-measures on E).
We say that µn → µ in the weak-atomic topology if
• µn ⇒ µ in the weak topology and
• µ∗n ⇒ µ∗ in the weak topology, where µ∗ :=

2
x∈E µ({x}) δx .

P

♦

27

Proposition 6.2. Assume that E = R, and let µ, µ1 , µ2 , . . . ∈ Mf (E) be finite measures with µn (E) → µ(E). Then µn → µ in the weak atomic topology if and only if Fn → F in the Skorohod topology on D(R, R+ ), where
F (t) := µ((−∞, t]), F1 (t) := µ1 ((−∞, t]), F2 (t) := µ2 ((−∞, t]), . . ., t ≥ 0.
Proof. First observe that a classical result says that µn ⇒ µ is equivalent to
Fn (t) → F (t) and µn (E) → µ(E) for all continuity points t of F .
“⇒” If µ({t}) > 0 for some t ∈ R, then, according to Lemma 2.5 in
[EK94], there is an unique sequence (tn )n∈N with (µ({tn }), tn ) → (µ({t}), t)
and all other sequences sn → t with sn 6= tn satisfy µ({sn })) → 0. Moreover, a simple application of the Portmanteau Theorem gives: For all ε > 0 there is a δ̄ > 0 such that µn ((t − δ, tn ) ∪ (tn , t + δ)) < ε for all n large enough and
µ((t − δ, t) ∪ (t, t + δ)) < ε for all δ < δ̄. If we now choose δ > 0 in such a way that t − δ is a continuity point of F , then lim |Fn (tn ) − F (t)| = lim |µn ((−∞, tn ]) − µ((−∞, t])|

n→∞

n→∞

≤ lim |Fn (t − δ) − F (t − δ)| + lim |µn ({tn }) − µ({t})|
n→∞

n→∞

+ lim sup |µn ((t − δ, tn )) − µ((t − δ, t))|

(6.1)

n→∞

≤ε
and hence Fn (tn ) → F (t). A similar argument shows that the conditions of Proposition 3.6.5 in [EK86] are satisfied and therefore Fn → F in the
Skorohod topology.
“⇐” Now let Fn → F in the Skorohod topology. Then for all discontinuity points t of F there is one sequence (tn )n∈N such that F (tn ) → F (t) and
F (tn −) → F (t−) (see (6.20) in the proof of Proposition 3.6.5 in [EK86]).
Since µ({t}) = F (t) − F (t−) this gives lim µn ({tn }) = lim (Fn (tn ) − Fn (tn −)) = F (t) − F (t−) = µ({t}) > 0.

n→∞

n→∞

(6.2)
Moreover, all other sequences (sn )n∈N with sn < tn and sn → t satisfy
|Fn (sn ) − F (t−)| → 0 and hence lim µn ({sn }) = lim (Fn (sn ) − Fn (sn −)) = 0

n→∞

n→∞

(6.3)

and the analogue holds for sequences sn > tn and sn → t. Hence we can apply Lemma 2.5 in [EK94] and get the result.

28

7

Proof of Theorem 2.12 (a), (b)

(a) First of all observe that
0

0

dGPa (u, u0 ) := dGPr (u, u0 ) + ν 2,u ({0}) − ν 2,u ({0}) + ρa (ν 2,u , ν 2,u )

(7.1)

is a metric on U, where ρa is given in [EK94]. Now, the properties follow analogue to Lemma 2.3 (combined with Lemma 2.5) in [EK94] and Proposition 5.6 in [GPW09].
(b) Recall the notation in Remark 5.1 and note that for u := [X, r, µ] ∈ U,
µh is purely atomic, h > 0. Let Ah be a finite subset of {x ∈ X : µh ({x}) >
0} with the property, that
µh (X \ Ah ) < h

(7.2)

and let µ̄h (·) := µh (· ∩ Ah ), then, by Lemma 5.2
[X, r, µ̄h ] → u.

(7.3)

In addition, note that
ν 2,u ({h0 }) = ν 2,Φ̂h (u) ({h0 })

for all 0 < h < h0

(7.4)

and
ν 2,[X,r,µh ] ({h0 }) − ν 2,[X,r,µ̄h ] ({h0 }) ≤ 2µh (X) · µh (X \ Ah ).

(7.5)

This shows [X, r, µ̄h ] → u in the Gromov-weak atomic topology (see again
Section 6 or [EK94]).
Let a ∈ Rn+ . Using an induction argument and the fact
S n ∈PN and k
that k∈N { i∈I ai : I ⊂ {1, . . . , n}} is countable, where ak ∈ Rn+ , k ∈ N, it is straight forward to see that one can approximate a (pointwise) by a sequences ak with
X
X
∀I, J ⊂ {1, . . . , n}, I ∩ J = ∅ : aki 6=
akj .
(7.6)
i∈I

Pn

j∈J

When we now take [X, r, µ] ∈ U with µ = i=1 aiP
δxi for xi ∈ X, i = 1, . . . , n, k
this shows, that the sequence of measures µ = ni=1 aki δxi satisfy µk ⇒ µ.
Using a similar argument as above together with Lemma 7.2 below, finally gives the result (compare also Proposition 5.6 in [GPW09]).

29

Assume Ah = {x1 , . . . , xn }. Then, another another way of proving this result is to disturb the measure a bit,P
i.e. to add a realization of independent, positive, variables U1 , . . . , Un with i Ui = 1/n to µh ({x1 }), . . . , µ({xn })
(compare Example 2.3).
In order to prove (c) of this theorem, we need some more results on the function F. Therefore, we skip the proof at this point and refer to Section
9.
Remark 7.1. Note that the above argument can be modified to prove that
F(I) is dense in F(U).
♣
Lemma 7.2. Let u = [{x1 , . . . , xn }, r, µ] ∈ U, n ∈ N ∪ {∞} and ai =
µ({xi }), i ∈ N. Then u ∈ I if and only if
X
X
ai 6=
ai ,
∀ I, J ⊂ {1, . . . , n}, I 6= J.
(7.7)
i∈I

i∈J

Proof. This follows directly by definition, since for
P all h ≥ 0 and x ∈
{x1 , . . . , xn } there is a set I such that µ(B̄h (x)) = i∈I µ({xi }).

8

Proofs for Section 3

Here we give the proofs of our results on the function F.

8.1

Proof of Lemma 3.1

For the first part we observe that since (supp(µ), r) is separable there is a countable set J ⊂ supp(µ), such that
[
B̄(x, h).
(8.1)
supp(µ) ⊂
x∈J

We define the set


I := I ⊂ J : µ B̄(x, h) ∩ B̄(y, h) = 0, ∀x, y ∈ I, x 6= y .

(8.2)

Note that ⊂ defines
S a partial order on I. If we take a totallySordered subset T ⊂ I, then A∈T A ∈ I (for two different elements x, y ∈ A∈T A, there is a set A0 ∈ T , since T is totally ordered, such that x, y ∈ A0 ) is an upper bound for T . By Zorn’s lemma, we can find a maximal set I ∈ I.

30

It remains to proof that
µ(X) = µ(supp(µ)) =

X

µ(B̄(x, h)).

(8.3)

x∈I

Note that for x, y ∈ supp(µ), since r is an ultra-metric µ almost surely, we either have

µ B̄(x, h) ∩ B̄(y, h) = 0,
(8.4)
or


µ B̄(x, h) ∩ B̄(y, h) = µ̃ B̄(y, h) .

(8.5)

By (8.1), we have
!
µ(X) = µ

[

B̄(x, h) .

(8.6)

x∈J

If we would assume that µ(X) >
would find a x̃ ∈ J such that

P

x∈I µ(B̄(x, h)), then, since I ⊂ J, we


µ B̄(x̃, h) ∩ B̄(x, h) = 0,

∀x ∈ I.

(8.7)

This is a contradiction, since I is a maximal element of I.
For the second part we set n
o
Ii := j ∈ {1, . . . , n(δ)} : µ(B̄(τjδ , δ) ∩ B̄(τih , h)) > 0 .

(8.8)

Since r is an ultra-metric µ-almost surely, we get µ(B̄(τjδ , δ) ∩ B̄(τih , h)) =
µ(B̄(τjδ , δ)) for all j ∈ Ii . This together with the first part implies ≥.
S
Let A := B̄(τih , h)\ j∈Ii B̄(τjδ , δ). If we assume that µ(A) > 0, then we can take a x ∈ A ∩ supp(µ) and, by Remark 3.5, we find a j such that
µ(B̄(x, δ) ∩ B̄(τjδ , δ)) = µ(B̄(x, δ)). It follows that µ(B̄(τih , h) ∩ B̄(τjδ , δ)) =
µ(B̄(τjδ , δ)) > 0 and hence j ∈ Ii . A contradiction and therefore µ(A) = 0.
To see that {Ii }i=1,...,n(h) forms a partition follows by similar arguments.

8.2

A result on relative compactness and proof of Lemma
3.8 and Proposition 3.13

We have the following result on relative compactness:
Proposition 8.1. (Relative compactness and further properties) Let (un )n∈N
be a sequence in U and u ∈ U.

31

(i) If un → u in the Gromov-weak topology and (hn )n∈N is a sequence in
(0, ∞) with hn → h ∈ (0, ∞), then {f(un , hn ) : n ∈ N} is relatively compact in S ↓ , equipped with d1 .
(ii) F(U) ⊂ D((0, ∞), S ↓ ) and lim max |f(u, h)i − f(u, 0)i | = lim d∞ (f(u, h), f(u, 0)) = 0, h↓0

i

h↓0

(8.9)

where f([X, r, µ], 0) := (µ({x})x∈X )↓ , i.e. the decreasing rearrangement of the atoms of µ.
(iii) If d1 (f(un , h), f(u, h)) → 0 for all continuity points h of f(u, ·), then
{un : n ∈ N} is relatively compact with respect to the Gromov-weak topology.
♣

Remark 8.2. Note that (ii) is Lemma 3.8.

Before we start we need the following result on monotonicity, which is a direct consequence of Lemma 3.1:
Lemma 8.3. Let 0 < δ ≤ h, u ∈ U and assume we are in the situation of
Lemma 3.1. Then n(h) ≤ n(δ). Moreover, for M ≤ n(h):
M
X

f(u, h)i ≥

i=1

M
X

f(u, δ)i .

(8.10)

i=1

We are now able to prove Proposition 8.1 (i).
Proof. Proposition 8.1 - (i) Note that if we equip SC↓ , C > 0 with d∞ (x, y) := max{|xi − yi | : i ∈ N},

x, y ∈ SC↓ ,

(8.11)

then (SC↓ , d∞ ) is a compact space (this follows analogue to Proposition
2.1. in [Ber06]).
p
Note that un → u implies un → u, where u = ν 2,u ([0, ∞)) and hence we find a constant C > 0 such that
X
sup un = sup f(un , hn )i ≤ C.
(8.12)
n∈N

n∈N i∈N

It follows that {f(un , hn ) : n ∈ N} is relatively compact in (SC↓ , d∞ ).
Hence, there is a x ∈ SC↓ such that d∞ (f(unk , hnk ), x) → 0 along some

32

subsequence and we have to show that d1 (f(unk , hnk ), x) → 0. We suppress the dependence on the subsequence and set
δ := inf hn > 0.
n∈N

(8.13)

Next we prove that for all 0 < ε ≤ δ there is a M ∈ N such that sup

∞
X

f(un , hn )i < ε.

(8.14)

n∈N i=M +1

We assume the converse, i.e. assume there is an ε > 0 with ε ≤ δ such that for all M ∈ N there is a n ∈ N with
∞
X

f(un , hn )i ≥ ε.

(8.15)

i=M +1
C
Note that f(un , ε̄)i ≤ M
for all i ≥ M , M ∈ N, ε̄ > 0. Moreover, note that when ε̄ ≤ δ, Lemma 8.3 implies
∞
X

f(un , ε̄)i ≥ ε.

(8.16)

i=M +1

Since [Xn , rn , µn ] = un → u, we have (see Proposition 7.1 in [GPW09] and
Proposition B.2 in [GGR16]):
0 = lim sup ν C (un )
M →∞ n∈N M




C
rn
= lim sup inf ε̄ > 0 : µn
≤ ε̄
x ∈ Xn : µn (B (x, ε̄)) ≤
M →∞ n∈N
M




C
rn
≥ lim sup inf ε̄ > 0 : µn x ∈ Xn : µn (B̄

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
