---
source: "https://arxiv.org/abs/1304.4278v1"
title: "Dendroidal weak 2-categories"
author: "Andor Lukacs"
year: "2013"
publication: "arXiv preprint / math.AT"
download: "https://arxiv.org/pdf/1304.4278v1"
pdf: "https://arxiv.org/pdf/1304.4278v1"
captured_at: "2026-09-19T20:23:40Z"
updated_at: "2026-09-19T20:23:40Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "György Lukács"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# Dendroidal weak 2-categories

- 著者: Andor Lukacs
- 年: 2013
- 掲載情報: arXiv preprint / math.AT
- 情報源: [arxiv](https://arxiv.org/abs/1304.4278v1)
- ダウンロード: https://arxiv.org/pdf/1304.4278v1
- PDF: https://arxiv.org/pdf/1304.4278v1

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

We discuss the dendroidal notion of weak higher categories introduced by Moerdijk and Weiss and we prove that dendroidal weak 2-categories are equivalent to bicategories.

## PDF Text

arXiv:1304.4278v1 [math.AT] 15 Apr 2013

Dendroidal weak 2-categories
ANDOR LUKÁCS
We discuss the dendroidal notion of weak higher categories introduced by Moerdijk and Weiss in [11] and we prove that dendroidal weak 2-categories are equivalent to bicategories.
55U10, 55P48, 55U40, 18G30, 18D50, 18D10;

1 Introduction
Weakened notions of categories are central in many branches of mathematics. One would often like to form certain “categories” where the composition of arrows is not strictly associative, but only up to some coherent higher cells that should be part of the structure. Important examples of such structures in the literature are homotopy ntypes. Roughly speaking, a homotopy n-type in topological spaces is the equivalence class of a space X such that all the homotopy groups πk (X) are trivial when k > n.
These classes are taken with respect to weak homotopy. Describing algebraic models for homotopy n-types is a classical problem in algebraic topology. For n = 2, 3
the first such models were given by Whitehead and Mac Lane [15, 10]. Following the influence of Grothendieck and R. Brown, who emphasized that groupoids should provide the natural framework for homotopy types, higher categorical algebraic models of homotopy 3-types were given and studied by Leroy [9], Joyal and Tierney [7], Berger
[2].
If we work out the topological intuitions coming from the interplay between spaces, maps and homotopies between maps, we arrive to the abstract notion of bicategories, first defined by Bénabou in [1]. Bicategories are structures that consist of 0-cells (objects), 1-cells (arrows) and 2-cells; 1-cells are composable but not strictly associatively, the failure of this is measured by some natural 2-cells. We can iterate the process to arrive to a definition of tricategories and so on, the n-th step of this process would give us a notion of weak n-categories. The problem we encounter is that each step we take for defining a one-level higher notion increases radically the complexity of the necessary coherence conditions between the higher dimensional cells. These steps also diminish the intuition on the nature of the coherence axioms. As a result, there exist a

2

Andor Lukács

plethora of different definitions of weak n-categories in the literature. Comparing the different notions of weak n-categories is one of the main problems in higher category theory. One of the issues can be formulated as follows. Intuitively, the right place to compare two notions of weak n-categories would be inside a weak n + 1-category, but how do we decide which notion of weak n + 1-categories to use for this comparison?
To deal with these problems, one can consider stricter- or non-iterative approaches to define weak n-categories. The idea is that the resulting stricter notions should be enough to deal with the applications on the one hand, and the slogan is that “weak ncategories are strictifiable up to some extent” on the other. Examples of this approach include Baez and Dolan’s notion of (∞, n)-categories, Tamsamani categories, etc. We would like to mention explicitly one such example, originating in the observation that the category of categories embeds to simplicial sets, via the nerve functor:
N : Cat −→ sSets.
Certain simplicial sets which are not in the image of the nerve functor behave much like categories, except that “composition” of arrows is well defined only up to some higher degree terms in that simplicial set. A. Joyal in [6] called these simplicial sets quasicategories, although the notion was already introduced by Boardman and Vogt under the name of restricted Kan complex in [4]. A quasi-category is an (∞, 1)-category in Baez and Dolan’s sense, i.e. all the degree 2- or higher cells are invertible. An important fact about quasi-categories is that they are exactly the fibrant objects in the
Joyal model structure for the category of simplicial sets.
The starting point of our investigation is that, there exist a generalisation of simplicial sets, that is suitable to study operads in the context of homotopy theory. On the one hand, operads (or rather coloured operads) can be viewed as generalisations of categories, where we consider arrows that can have multiple inputs as opposed to one.
It is natural to ask wether there exists a presheaf category that extends the category of operads in the same way as simplicial sets extend the category of categories via the nerve functor. The question was studied in the papers of Moerdijk and Weiss [11, 12]: the category of dendroidal sets satisfies the requirements and fits in a commutative diagram of categories
Cat


N

/ sSets




Nd


/ dSets

Op

Since dendroidal sets are an extension of simplicial sets, suitable for studying the homotopy theory of operads, the theory of dendroidal sets inherits a lot of questions

Dendroidal weak 2-categories

3

from the theory of simplicial sets. For example, this extension allows us to consider quasi-operads in the category of dendroidal sets, i.e. analogs of Joyal’s quasi-categories.
One can then ask whether the Joyal model structure on the category of simplicial sets extends to that of dendroidal sets in such a way, that the fibrant objects of this model category are exactly the quasi-operads. Cisinski and Moerdijk in [5] gave a positive answer to this question. One nice feature of dendroidal sets, observed in [11], is that they contribute to the theory of higher categories with a new compact definition of weak n-categories.
The aim of this paper is to study the dendroidal definition of weak n-categories mentioned above in low degree. We restrict ourselves to the cases n = 1 and n = 2.
In the case of degree 2, the corresponding classical notion is bicategories. We prove that dendroidal weak 2-categories are equivalent to bicategories (even more, they are isomorphic as categories).
Our paper is organised as follows:
In Section 2 we introduce dendroidal sets, with emphasis on the necessary notions and terminology that will be used in the next Sections. The symmetric monoidal structure on the category of dendroidal sets that makes the definition of dendroidal weak ncategories possible is induced by the dendroidal nerve functor and the Boardman-Vogt tensor product of coloured operads. The dendroidal Grothendieck construction gives us a way to systematically glue together dendroidal sets, and is an important ingredient that will allow us to consider later dendroidal weak n-categories with any set of objects.
The weakening of the higher dimensional cells in the dendroidal setting is done with the homotopy coherent dendroidal nerve functor, that is also introduced in this Section.
In Section 3 we define dendroidal weak n-categories and prove that they are 3coskeletal for every n. This is an important property of dendroidal weak n-categories, since it implies that degree 0, 1 and 2 completely determines dendroidal weak ncategories.
In Section 4 first we describe dendroidal weak 1-categories. The iterative definition of dendroidal weak n-categories makes it then possible to discuss dendroidal weak 2categories. We prove that the quasi-category of dendroidal weak 2-categories (denoted by i∗ (wCat2 )) has equivalent homotopy category to the category of bicategories:
Theorem 4.5 The category of unbiased bicategories ubiCtg and ho(i∗ (wCat2 )) are isomorphic. Hence the category of classical bicategories is equivalent to the category of dendroidal weak 2-categories.
The definition and basic properties of classical and unbiased bicategories are recalled in the Appendix.

4

Andor Lukács

2 Dendroidal sets
2.1 Terminology and basic facts about dendroidal sets
Dendroidal sets generalise simplicial sets in a suitable way for studying the homotopy theory of (coloured) operads and their algebras. They were introduced in the papers of I. Moerdijk and I. Weiss [11, 12]. The idea behind the notion of dendroidal sets is that in the same way as simplicial sets help us understanding categories via the nerve functor, there should be an analogous notion for studying coloured operads as a generalisation of categories. Our goal here is to write a self-contained introduction to dendroidal sets, including all the terminology necessary for the next Sections.
Let us start with the notion of trees. A tree is a finite non-planar contractible graph with a distinguished leaf called root. A tree thus has many planar representatives, when we draw a picture of it we actually pick one. We will use the following terminology on trees: Corn denotes the n-corolla, i.e. the tree with one vertex an n + 1 leaves (one of these leaves is the root), Vert(T) denotes the set of vertices of the tree T , Edg(T) denotes the set of edges of the tree T and InEdg(T) denotes the set of internal edges of the tree
T . We will say that a vertex v ∈ Vert(T) of a tree is an outer vertex if v is adjacent to at most one inner edge of T . For example, on the following picture of a (planar representative of a) tree T we have Vert(T) = {u, v, w}, Edg(T) = {a, b, c, d, e, f },
InEdg(T) = {c, b}. The vertices u and w are outer vertices.
❄❄
❄
⑧⑧
d ❄❄ e ⑧⑧f
❄❄ ⑧⑧
⑧❄w
•❄
⑧• u
❄❄
⑧⑧
b ❄❄ ⑧⑧c
❄⑧⑧
•v a

We will make frequent use of symmetric coloured operads (both in Sets and enriched in a monoidal category E ) in the sense of [8] and we will refer to them as operads from now on. Recall that if P is an operad in Sets, then it comes equipped with a set of objects or colours ob(P) and for each ordered sequence of objects σ = (e1 , . . . , en ; e), a set of operations P(e1 , . . . , en ; e) = P(σ). We will use the ◦i -definition for the composition of operations, i.e. if σ is an ordered sequence or a signature as before and
ρ = (f1 , . . . , fm ; ei ) for a fixed 1 ≤ i ≤ n then
σ ◦i ρ = (e1 , . . . , ei−1 , f1 , . . . , fm , ei+1 , . . . , en ; e)
and there is a given composition map
◦i : P(σ) × P(ρ) −→ P(σ ◦i ρ).

Dendroidal weak 2-categories

5

The category of operads in Sets will be denoted by Op, and the category of planaror non symmetric operads in Sets by Opπ . Sometimes it will be useful to construct operads from planar ones, via the free symmetrization functor Symm : Opπ −→ Op, the left adjoint to the forgetful functor U : Op −→ Opπ . This feature already appears in the definition of dendroidal sets.
The category Ωπ consists of planar trees as objects and planar operad maps as arrows.
To be more precise, any planar tree T gives rise to a planar operad Ωπ (T). The objects of this non symmetric operad are the edges of T , and the operations are freely generated by the vertices of T , i.e. if σ = (e1 , e2 , . . . , en ; e) is an ordered sequence of edges of
T and there is a vertex v with incoming edges e1 , . . . , en in this order and outgoing edge e, then Ωπ (T)(σ) = {v}. One can then “compose” vertices, indicated by the tree
T . Hence a map R −→ T in Ωπ is simply a planar operad map Ωπ (R) −→ Ωπ (T).
We observe that if f : R −→ T is an isomorphism, then the planar operad structures imply that R and T have the same planar shape and they differ only on the names of their edges and vertices. To avoid dealing with these irrelevant isomorphisms, further on we will replace Ωπ by a skeleton of it, and call this new category Ωπ as well. With this new convention, we observe that all the maps of Ωπ are generated by two types, faces and degeneracies. These types of maps generalise the corresponding notions in the category ∆ defining simplicial sets, in the following way. Let Ln denote the linear tree with n vertices, n ≥ 0:
•
•

..
.
•

If we consider the categorical definition of ∆, we observe that the category
[n] = 0 o

1o

2o

· · ·o

n

is in fact [n] = Ω(Ln ), hence ∆ is fully faithfully embedded into Ωπ by [n] 7→ Ln .
The face maps in Ωπ are all those monic operad maps ∂ : Ωπ (R) −→ Ωπ (T) which increase the number of vertices by one (i.e. | Vert(T)| = | Vert(R)| + 1) and the degeneracies are all those epic operad maps σ : Ωπ (T) −→ Ωπ (R) which decrease the number of vertices by one.
It follows that face maps can be of the following types:

6

Andor Lukács

(a) the following picture is an example of an outer face
❄❄
⑧• u
❄❄
⑧⑧
b ❄❄ ⑧⑧c
❄⑧⑧
•v

❄❄
❄
⑧⑧
d ❄❄ e ⑧⑧f
❄❄ ⑧⑧
⑧❄w
•❄
⑧• u
❄❄
⑧⑧
b ❄❄ ⑧⑧c
❄⑧⑧
/
•v

∂w

a

a

which is just an inclusion of operads;
(b) the following picture is an example of an inner face
✵✵
✍
❖❖❖e ✵✵ ✍✍f ♦♦
✵
✍
❖
❖
❖❖✵♦
d
✍ ♦♦♦c
•✍ ♦u

❄❄
❄
⑧⑧
d ❄❄ e ⑧⑧f
❄❄ ⑧⑧
⑧❄w
•❄
⑧
❄❄
⑧⑧
b ❄❄ ⑧⑧c
❄⑧⑧
/
•v

∂b

a

a

where ∂b : Ωπ (R) −→ Ωπ (T) is the identity on the objects (edges), and sends the operation u ∈ Ωπ (R)(d, e, f , c; a) to the composite operation v ◦1 w ∈
Ωπ (T)(d, e, f , c; a), which we can denote without ambiguity by v ◦b w.
Note that the seemingly special cases of face maps into the corolla Corn , n ≥ 2 fall under case (a): these face maps are all the n + 1 possible edge inclusions of the trivial tree | to Corn .
On the other hand, a degeneracy always looks like
❄❄
❄❄ ⑧⑧⑧
❄❄
⑧
❄❄
•⑧❄❄
❄❄ ⑧⑧⑧
e ❄❄
❄⑧⑧
❄
v
• ❄❄
⑧•
⑧
f ❄❄ ⑧⑧
⑧
•

σv

❄❄
❄❄
❄❄ ⑧⑧⑧
❄❄ ⑧⑧⑧
❄⑧
❄⑧
⑧
❏
• ❏❏❏
t•⑧
t t
❏
e ❏❏ ttt
/
•t

where both of the objects e, f are sent to e, the operation v to the identity operation ide and σv is the identity elsewhere.
We will use the following terminology with respect to faces and degeneracies:
• If e is an inner edge of a tree T , then T/e denotes the tree resulting from T

by contracting e. The inner face corresponding to this contraction is usually denoted by ∂e : T/e −→ T .
• If v is an outer vertex of a tree T (that is, it has exactly one inner edge adjacent to it), then T/v denotes the tree resulting from T by removing v and all the external edges adjacent to it (with the obvious choice for the root of T/v when

Dendroidal weak 2-categories

7

one of these external edges happens to be the root of T ). We call this procedure
“cutting vertex v”. The outer face correspondig to cutting v is usually denoted by ∂v : T/v −→ T .
• If v is a vertex of valence one of a tree T then T\v denotes the tree resulting from T by removing v. The degeneracy corresponding to this removal is usually denoted by σv : T −→ T\v.
The category Ω is obtained from Ωπ via the functor Symm. The objects of Ω are (non planar) trees and the arrows R −→ T are operad maps Symm(Ω(R̄)) −→ Symm(Ω(T̄)), where T̄ denotes a planar representative of T . One can check that the resulting operad does not depend on the chosen representatives, hence the definition makes sense. Later on we will use this independence from chosen representatives: we often describe the operad Ω(T) by picking a representative T̄ and giving only the description of the planar operad Ωπ (T̄).
The definition given above implies that there is an extra type of generator for the maps in Ω, namely the isomorphisms.
The category of dendroidal sets is the presheaf category on Ω: op

dSets := SetsΩ = Funct(Ωop , Sets).
If X is a dendroidal set, the elements of XT are called dendrices of shape T . The representable dendroidal set associated to a tree T is the functor
Ω[T] := Ω(−, T) : Ωop −→ Sets.
By the Yoneda lemma, a dendrex t ∈ XT is the same thing as a map of dendroidal sets Ω[T] −→ X . The Yoneda lemma in general is a very useful tool in the theory of simplicial- and dendroidal sets, allowing us to swap between maps and dendrices whenever needed. We are going to exploit this property in the following without mentioning it. A first application of the Yoneda lemma in the dendroidal setting proves that every dendroidal set is a colimit of representable ones, a property generalising the well known fact for simplicial sets.
For any given tree T one can define some dendroidal subsets of the representable Ω[T], like the boundary ∂Ω[T] or the inner horn Λe [T] with respect to the inner edge e.
Dendroidal sets are analogous to simplicial sets in many ways. For example, inner horns can be used to define inner Kan complexes in the category of dendroidal sets: we say that a dendroidal set X satisfies the inner Kan condition if for any inner horn h : Λe [T] −→ X there exists a dendrex t : Ω[T] −→ X such that the following diagram

8

Andor Lukács

commutes:
Λe [T]


h

/X
③=
③
③③
③③t
③
 ③③

Ω[T]

In this case X is called an inner Kan complex or a quasi-operad, analogously to the simplicial case where an inner Kan complex was called by A. Joyal a quasi-category.
Another notion that generalises from simplicial sets and categories to dendroidal sets and operads is the nerve functor. The dendroidal nerve Nd : Op −→ dSets can be defined by setting for any operad P

Nd (P) T := Op(Ω(T), P).

In the next few lines we introduce the notions of k-skeleton and k-coskeleton of a dendroidal set. For every k ∈ N let Ωk denote the full subcategory of Ω consisting of trees with at most k vertices. The presheaf category on Ωk is called the category of k-truncated dendroidal sets and is denoted by dSetsk . The inclusion ik : Ωk −→ Ω
induces the truncation functor i∗k : dSets −→ dSetsk which has a left adjoint (ik )! and a right adjoint (ik )∗ . It follows that the composites
(ik )! i∗k , (ik )∗ i∗k : dSets −→ dSets form an adjoint pair of endofunctors. The left adjoint (ik )! i∗k is usually denoted by Skk and is called the k-skeleton functor. The right adjoint is denoted by coSkk and is called the k-coskeleton functor.
A dendroidal set X is said to be k-coskeletal if the canonical map X −→ coSkk (X) is an isomorphism. Another way to state this is that for every dendroidal set Y and every map of dendroidal sets φ : Skk (Y) −→ X there exists a unique extension
Skk (Y)


φ

/X
<

∃!

Y
Since any Y ∈ dSets is a colimit of representables, we can infer that if the previous statement holds for all Y = Ω[T] where T is any tree with k + 1 vertices, then it holds in general. In this case Skk (Y) = Skk (Ω[T]) = ∂Ω[T]. Note that in view of the
Yoneda lemma we can think of the composite
Skk (Ω[T]) /

/ Ω[T]

t

/X

Dendroidal weak 2-categories

9

as the boundary- or k-skeleton of the dendrex t. To emphasize this point of view, sometimes we will denote this composite by Skk (t).
One can define the dual notion of k-skeletal dendroidal sets similarly.
Remark 2.1 Note that the dendroidal definition of the functors Skk and coSkk uses a filtration of the objects of Ω by the number of the vertices of trees. Later on, we will use the term degree to refer to this natural number.

2.2 A closed symmetric monoidal category structure on dendroidal sets
Since dSets is a presheaf category, it can be endowed with the usual cartesian closed category structure present in any presheaf category. There is another interesting symmetric monoidal structure on dSets that will prove to be useful in the definition of dendroidal weak n-categories of Section 3. Our goal is to recall this monoidal structure in the current section, together with those properties that will be used. For more details on this subject one can consult [11, 12, 14].
One way to define the mentioned monoidal structure on dSets is by transferring the
Boardman-Vogt monoidal structure on Op, via the dendroidal nerve functor. We adopt this road, and we start by recalling the Boardman-Vogt tensor product for symmetric operads (a generalisation of the B-V tensor product for classical operads in [4]).
Let P, Q ∈ Op. We define a new operad, P ⊗ Q, as follows. The set of objects is ob(P ⊗ Q) := ob(P) × ob(Q) and we denote the elements of this set by a ⊗ x := (a, x).
We describe the operations of P ⊗ Q in terms of generators and relations. There are two types of generators:
(a) For any p ∈ P(a1 , . . . , an ; a) and any x ∈ ob(Q), p ⊗ x ∈ P ⊗ Q(a1 ⊗ x, . . . , an ⊗ x; a ⊗ x).
(b) For any a ∈ ob(P) and any q ∈ Q(x1 , . . . , xm ; x), a ⊗ q ∈ P ⊗ Q(a ⊗ x1 , . . . , a ⊗ xm ; a ⊗ x).
The relations also are of two types:
(a) Relations that imply precisely that the obvious maps
P

id ⊗x /

P ⊗ Q and Q

a⊗id /

P⊗Q

are maps of operads for any fixed x ∈ ob(P), a ∈ ob(Q).

10

Andor Lukács

(b) For any p ∈ P(a1 , . . . , an ; a) and q ∈ Q(x1 , . . . , xm ; x) the following two operations are the same in P ⊗ Q
❄❄
❄

❄❄
❄❄
⑧⑧
❄❄. . .⑧⑧⑧an ⊗xm
❄⑧
an ⊗qtt•⑧
❏❏
t
. . . ttttan ⊗x a1 ⊗x ❏❏❏
◦t

❄❄
❄

⑧⑧

. . .⑧⑧⑧a1 ⊗xm a1 ⊗x1❄❄❄
❄❏
⑧
•⑧❏a❏1 ⊗q p⊗x

❄❄
❄❄
⑧⑧
❄.❄. .⑧⑧⑧an ⊗xm p⊗xmt◦⑧
❏❏
tt
❏
a⊗x1 ❏❏. . . ttta⊗xm
❏tt
σn,m a⊗q•
⑧⑧

a1 ⊗x1❄❄.❄. .⑧⑧⑧an ⊗x1
◦⑧❏p⊗x
❏ 1

=

a⊗x

a⊗x

where σn,m ∈ Σn·m denotes the permutation that makes the order of the inputs on the right-hand side of the equation the same as the order of the inputs on the left-hand side.
The tensor product we defined is a bifunctor − ⊗ − : Op× Op −→ Op and it induces a symmetric closed monoidal category structure on Op. The right adjoint of any functor
− ⊗ Q is denoted by Op(Q, −) : Op −→ Op. In particular, Op(Q, Sets) is the operad of Q-algebras. (For the definition of Op(Q, −) see [14].)
We can now make use of the functor Nd : Op −→ dSets to transfer the Boardman-Vogt tensor product to dendroidal sets:
– For any two representable dendroidal sets Ω[T] and Ω[R], define
Ω[T] ⊗ Ω[R] := Nd (Ω(T) ⊗ Ω(R)).
– Extend the definition cocontinuously, i.e. for any X, Y ∈ dSets write X =
colimT Ω[T], Y = colimR Ω[R] as colimits of representables and define
X ⊗ Y := colimT,R Ω[T] ⊗ Ω[R].
The bifunctor − ⊗ − : dSets × dSets −→ dSets induces a symmetric closed monoidal structure on dSets, the right adjoint of − ⊗ Y is the functor dSets(Y, −) : dSets −→ dSets, given on objects (by Yoneda lemma) by dSets(Y, Z)T = dSets(Y ⊗ Ω[T], Z).
The following properties will prove to be useful in Section 4:
Proposition 2.2 (Lemma 4.3.3 in [14]) For any operad P ∈ Op and for any tree
T ∈Ω
Nd (P) ⊗ Ω[T] ≃ Nd (P ⊗ Ω(T)).
Proposition 2.3 (Corollary 9.3 in [12]) For all operads P, Q ∈ Op dSets(Nd (P), Nd (Q)) ≃ Nd (Op(P, Q)).

Dendroidal weak 2-categories

11

2.3 The dendroidal Grothendieck construction
The aim of this Section is to provide an ingredient we are going to use in the description of dendroidal weak higher categories. The data we start with is a functor X : Sop −→
dSets
R where S is a cartesian category, and we are going to assign to X a new dendroidal set S X , called the Grothendieck construction of X .

To achieve this goal, we need some preliminary definitions. Since S is cartesian it is an operad, hence it makes sense to talk about the dendroidal set Nd (S). Suppose that for a fixed tree T , t ∈ Nd (S) is a dendrex of shape T . That is, t intuitively looks like the tree T decorated with objects and operations of the operad S:
❄❄
❄❄
❄
❄
⑧⑧
⑧⑧
s7 ❄❄❄ ⑧⑧⑧s8
s4❄❄❄ s5⑧⑧⑧s6
❄⑧
❄❏⑧
❏❏ u3 •
t•⑧u4
u2 •⑧
t
❏❏❏
t s1 ❏ s2 ttts3
❏❏ tt
•tu1
s0

where the si are objects of S, and – for example – u1 : s1 × s2 × s3 −→ s0 is a map in
S. To such a t we can assign an object of S, called in(t), which is the cartesian product of the objects labeling the leaves of T : since t ∈ Op(Ω(T), S),
Y
in(t) :=
t(l).
l∈Leaves(T)

Furthermore, we can assign to a t ∈ Nd (S)T and a map α : R −→ T of Ω a map in S
in(α) : in(t) −→ in(α∗ t)
by first composing the maps of S, indicated by t and α, and then taking the product.
For example, if α : R −→ T is the inclusion to the root vertex (in this case a composite of three outer faces)
❏❏❏
❏❏❏ tttt
❏ttt
•

α
/

❄❄
❄❄
❄❄ ⑧⑧⑧
❄❄ ⑧⑧⑧
❄⑧⑧
❄❏
⑧
⑧
• ❏❏ •
❏❏❏ tttt•
❏ttt
•

and t is as above, then α∗ t is
❏❏❏
❏
ttt s1 ❏❏❏ s2 ttts3
❏❏ tt
•tu1
s0

12

Andor Lukács

and in(α) = u2 × u3 × u4 . In particular, if α is an inner face or a degeneracy then
β

α

/R
/ T are maps of Ω
in(α) is the identity map of in(t) = in(α∗ t), and if R′
then in(αβ) = in(β) in(α).
R
R 
In view of the definitions above we can define S X as follows. The set S X T consists of pairs (t, x) where t ∈ Nd (S)T and a
x : Ω[T] −→
X(s)
s∈S

is a degree preserving map such that x(r) ∈ X(in(r∗ t)) for any r ∈ Ω[T]R . There is one more condition on x: it has to be compatible with the dendroidal structure of the r /
α /
R
T in various dendroidal sets involved. Explicitly, for a chain of arrows R′
∗
Ω we have r ∈ Ω[T]R and α r = rα ∈ Ω[T]R′ , hence


x(r) ∈ X in(r∗ t) R
and x(α∗ r) ∈ X in((rα)∗ t) R′ .
The data above also induces two maps

X in(r∗ t) R
◆◆◆
◆◆◆
◆◆
α∗ ◆◆◆'


X in((rα)∗ t) R′

♥♥
♥♥♥
♥
♥
♥
♥♥♥ X(in α)
 ♥w

X in(r∗ t) R′

We require



α∗ x(r) = X(in α) x(α∗ r) .
R
The dendroidal structureRon  S X is defined as follows. Suppose that δ : R −→ T is a map in Ω and (t, x) ∈ S X T a dendrex of shape T . The map δ induces the map of dendroidal sets Ω[δ] : Ω[R] −→ Ω[T]. We define

(2–2)
δ∗ (t, x) := δ∗ t, x ◦ Ω[δ] .
(2–1)

R
One can check that with this structure S X is indeed a dendroidal set. The following theorem and proposition collect two important properties of the dendroidal Grothendieck construction.

Theorem 2.4 ( [11, 14]) Let X : Sop −→ dSets be a diagram
R of dendroidal sets. If for all s ∈ S every X(s) is an inner Kan complex then so is S X .

Proposition 2.5 Let X : Sop −→ dSets be a diagram of dendroidal
R sets and k ≥ 2 a natural number. If X(s) is k -coskeletal for every s ∈ S then so is S X .

Dendroidal weak 2-categories

13

Proof Let us start with the remark that k ≥ 2 is needed because dendroidal nerves of operads are 2-coskeletal (a generalisation of the well known fact for nerves of categories, proven in [11, 14]).
Our task is to proveRthat, for any tree T with k + 1 vertices, every map of dendroidal sets φ : ∂Ω[T] −→ S X extends uniquely as
∂Ω[T]



φ

/

;
∃!

R

SX

Ω[T]
R 
We suppose existence and prove uniqueness first. Let (t1 , x1 ), (t2 , x2 ) ∈ S X T be two dendrices filling the boundary φ. The dendroidal set Nd (S) is k-coskeletal since k ≥ 2. Hence by equation (2–2) we can infer that t1 = t2 = t. Let u : R −→ T be a face. Since u∗ (t, x1 ) = u∗ (t, x2 ), we obtain x1 ◦ Ω[u] = x2 ◦ Ω[u]. On the other hand,

xi (u) = xi ◦ Ω[u] (idR )

for i = 1, 2, implying x1 (u) = x2 (u). We can use now equation (2–1) for r = idT and
α = u to conclude that u∗ (x1 (idT )) = u∗ (x2 (idT )) as dendrices of shape R in X(in(t)).
Since this is true for any face u : R −→ T and X(in(t)) is k-coskeletal, it follows that also x1 (idT ) = x2 (idT ). We can infer that x1 = x2 , thus the filler is unique.
The argument above also contains the information how to construct a filler (t, x) of φ, giving a proof of the existence of such an extension.
Remark 2.6 If we restrict our attention to dendroidal sets where the only nontrivial dendrices are of linear shapes, Proposition 2.5 implies that the same property is true for simplicial sets and the simplicial Grothendieck construction.

2.4 The homotopy coherent dendroidal nerve of an operad
Let E be a symmetric monoidal model category with an interval H : that is, an object
H of E , together with two points 0 : I −→ H and 1 : I −→ H , an augmentation
ǫ : H −→ I and an associative binary operation ∨ : H ⊗ H −→ H for which 0 is unital and 1 is absorbing, satisfying ǫ1 = ǫ0 = id. In this case one can modify the nerve construction for operads enriched in E in such a way that the resulting dendroidal set encodes also the homotopies in the operad.

14

Andor Lukács

An interesting example of such a situation is when E is the category of categories with the usual cartesian product, the folk model structure and the interval H is the category
0

≃

/1

with two objects and one isomorphism between them. The required structure on H is the obvious one: 0 is the neutral element, 1 is the absorbing one, and the rest of the interval structure on H is completely determined by the previous choices. Indeed, since the unit of the monoidal structure is the terminal object in E (the category ∗ with one object and no other morphisms than the identity), the counit ǫ : H −→ ∗ is obvious.
The various compatibility conditions imply that the monoid structure ∨ : H × H −→ H
is given by “the maximum operation”: on the objects, i ∨ j = max{i, j}.
Since we are interested only in this example, from now on E denotes the category of categories with the structure mentioned above, although everything can be carried out similarly in the general case.
Let us denote the category of operads enriched in E by OpE . The functor hcNd : OpE −→ dSets is defined by hcNd (P)T := OpE (W(Ω(T)), P)
where W is the W -construction for coloured operads (see [3] for details) and Ω(T) is the discrete version in E of the operad induced by the tree T . We will need later an explicit description of W(Ω(T)), hence we give it here.
Recall that for a tree T
Ω(T) = Symm(Ωπ (T̄))
where Symm : OpπE −→ OpE is the E -enriched version of the symmetrization functor from non symmetric operads to operads, and T̄ is any planar representative of T .
Moreover, the W -construction commutes with Symm, thus

WΩ(T) = Symm WΩπ (T̄) .

This property allows us to describe WΩ(T) by using an arbitrary planar representative of T . The objects of WΩπ (T̄) are the edges of T . Suppose that σ = (e1 , e2 , . . . , en ; e)
is an ordered sequence of objects. We can distinguish two cases for the category of operations corresponding to σ :
(1) If Ωπ (T̄)(σ) = ∅ – the empty category – then also WΩπ (T̄)(σ) = ∅.

Dendroidal weak 2-categories

15

(2) If Ωπ (T̄)(σ) 6= ∅, it follows that T̄ has a subtree T̄σ with leaves e1 , . . . , en and root e. The set of internal edges of T̄σ is denoted by InEdg(T̄σ ). From the
W -construction it follows then that
Y
WΩπ (T̄)(σ) =
H, f ∈InEdg(T̄σ )

where in case the product is empty the result is the unit object of the monoidal structure, which is the category ∗ with one object and no other morphism than the identity.

We still need to define the composition maps in the operad WΩπ (T̄). Suppose that
σ = (e1 , . . . , en ; e) and ρ = (f1 , . . . , fm ; ei ) are ordered sequences of edges of T , such that neither Ωπ (T̄)(σ), nor Ωπ (T̄)(ρ) is the empty category. It follows that T̄
has subtrees T̄σ and T̄ρ , the sets of internal edges of these trees are disjoint and the tree T̄σ◦i ρ obtained by grafting along the edge ei has one more internal edge than the previous two together. Let us denote these sets of internal edges by int(σ), int(ρ) and int(σ ◦i ρ) respectively. The composition map
◦i : WΩπ (T̄)(σ) × WΩπ (T̄)(ρ) −→ WΩπ (T̄)(σ ◦i ρ)
is given by
Q

int(σ)

H

!

×

Q

int(ρ)

H

!

≃

Q

int(σ)∪int(ρ)

H

!

×∗

id ×1 /

where the functor 1 : ∗ −→ H is the absorbing element of H .

Q

H,

int(σ◦i ρ)

This concludes the description of the operad WΩ(T). Note that we still need to mention how the dendroidal structure on hcNd (P) is defined. If δ : R −→ T is a face map in
Ω then it induces a map of operads δ : WΩ(R) −→ WΩ(T) via the neutral element functor 0 : ∗ −→ H . In case δ is a degeneracy, the induced functor is obtained by the monoid structure ∨ : H × H −→ H . These definitions are functorial, hence they induce a dendroidal structure on hcNd (P).

3 Dendroidal weak n-categories
For any set A there exists a planar operad AsπA whose algebras are small categories with set of objects A. The objects of AsπA are ordered pairs (a1 , a2 ) ∈ A × A, and the sets of operations are defined by

AsπA
; (a, a) = ∗,

AsπA (a1 , a2 ), (a2 , a3 ), . . . , (an−1 , an ); (a1 , an ) = ∗

16

Andor Lukács

and in all the other cases the set of operations is empty (those ordered sequences
σ = (c1 , c2 , . . . , cn ; c) of objects of A × A for which AsπA (σ) is not empty will be called admissible).
Let α : AsπA −→ Sets be a map of operads. The data-part of such an α determines for any (a1 , a2 ) ∈ A × A a set A(a1 , a2 ) and for any admissible signature σ =
(a1 , a2 ), (a2 , a3 ), . . . , (an−1 , an ); (a1 , an ) a function

compσ : A(a1 , a2 ) × A(a2 , a3 ) × · · · × A(an−1 , an ) −→ A(a1 , an )

; (a, a) is a function ∗ −→ A(a, a). The which in the particular case of σ =
compatibility-part of such an α ensures that the various functions compσ fit nicely to define units and compositions of arrows in a category A with object set A. Indeed, we arrive to the conclusion that the relevant ordered sequences are of the type

(a1 , a2 ), (a2 , a3 ); (a1 , a3 ) and
; (a, a) , etc.

Since the forgetful functor U : Op −→ Opπ is right adjoint to the symmetrization functor, we infer that the algebras of the operad AsA := Symm(AsπA ) are categories with set of objects A as well.

Remark 3.1 Note that in the description of AsπA -algebras given above we used the unconventional “left-to-right” composition order for arrows, i.e.
A(a1 , a2 ) × A(a2 , a3 )

/ A(a1 , a3 )

instead of the conventional
A(a2 , a3 ) × A(a1 , a2 )

/ A(a1 , a3 )

To avoid unnecessary complications in the future, arising only from notation, whenever we need to give such a composition map associated to some signature σ , we will always stick to the order determined by σ , thus the unconventional order. However, when it is required to give extra details with explicit composites of maps, we will use the conventional “right-to-left” order.
Let X be a dendroidal set and define the functor
Cat(X)− : Setsop −→ dSets,

Cat(X)A := dSets(Nd (AsA ), X).

The dendroidal set of categories enriched in X (see [11]) is by definition the Grothendieck construction of Cat(X)− . We denote it by
Z
Cat(X)− .
Cat(X) :=
Sets

Dendroidal weak 2-categories

17

One can iterate the process above to obtain a definition of the dendroidal set of ncategories enriched in X :
Cat0 (X) := X,
Catn (X) := Cat(Catn−1 )(X).
To see why this definition is plausible, one can try particular choices of X . For example if X = Nd (Sets), we can prove inductively that Catn (X) is the dendroidal nerve of strict n-categories with the classical definition (see also Example 4.5.5 in [14]). Indeed, for n=1
Z
dSets(Nd (AsA ), Nd (Sets))
Cat(Nd (Sets)) =
A∈Sets

≃

Z

Nd (Op(AsA , Sets))

Z

Nd (CategA )

A∈Sets

≃

A∈Sets

≃ Nd (Categ), where Categ denotes the usual monoidal category of small categories, viewed as an operad. The second part of the inductive proof is similar (one uses that for any monoidal category M, Op(AsA , M) ≃ CategA (M), where the right-hand side denotes the monoidal category of categories enriched in M, with set of object A).
We are interested here in another choice for X , which yields the dendroidal definition of weak n-categories: it is plausible to define X := hcNd (Ctg) where Ctg is the category of small categories enriched in E . (Recall from Section 2.4 that E denotes the symmetric monoidal model category of categories, together with the interval H .
Hence the set of functors between two fixed categories is a category with natural transformations as maps.)
Definition 3.2 ([11]) The dendroidal set of weak n-categories is defined as follows: wCat0 := Nd (Sets), wCatn := Catn−1 (hcNd (Ctg))

for n > 0.

The rest of this Section is dedicated to the study of coskeletality of weak n-categories.
The first result is
Lemma 3.3 If T ∈ Ω is a tree with 3 vertices and t, s ∈ wCatT1 satisfy Sk2 (t) = Sk2 (s)
then t = s .

18

Andor Lukács

Proof To illustrate our argument better, we will work with a chosen tree, the general case can be carried out in the same way. So let T ∈ Ω be the tree with a planar representative as below.
❄❄
⑧
❄❄
⑧⑧
d ❄❄❄ ⑧⑧⑧e
⑧❄v
•❄
⑧• w
❄❄
⑧⑧
❄
⑧
b ❄❄ ⑧⑧ c
•⑧u a

Let x ∈ wCatT1 be a dendrex of shape T , that is a map of operads enriched in E , x : WΩ(T) −→ Ctg. Let us adopt the notations of Section 2.4. It follows that x consists of compatible functors xσ : H int(σ) −→ Ctg(σ) and the only functor we have to describe in terms of the 2-skeleton of x is the one corresponding to σ = (d, e; a)
(the other functors lie in the image of Sk2 (x)). In this case the domain of xσ is the groupoid H 2 = Hc × Hb , represented as the square
(id0 ,b)

(0, 0)

/ (0, 1)

(c,id0 )

(c,id1 )



(id1 ,b)

(1, 0)


/ (1, 1)

where we think of the copy of H corresponding to an internal edge f as the groupoid f

Hf = 0 −→ 1. Since the domain is a groupoid, we observe that if xσ is already defined on a “connected part” of the “square” Hc × Hb , then it is defined on the “convex hull”
of that component. We conclude that in order to know xσ , it is enough to know its image on the sets of arrows Opd = {(c, id1 ), (id1 , b)} and Face = {(c, id0 ), (id0 , b)}.
To conclude the proof, first we show that since x is a map of operads, xσ (Opd) is determined by Sk2 (x). Indeed, the commutative square
Hc × {v}

x×x /

Ctg(b; a) × Ctg(d, e; b)
◦b

◦b



Hc × Hb

xσ



/ Ctg(d, e; a)

implies that we know xσ on the arrow (c, id1 ), and a similar square gives the image of
(id1 , b).
Second, we show that xσ (Face) is in the image of Sk2 (x). We observe that the inner faces ∂b : R −→ T and ∂c : R′ −→ T , according to the definition of the dendroidal structure on wCat1 , induce enriched operad maps WΩ(R) −→ WΩ(T)

Dendroidal weak 2-categories

19

and WΩ(R′ ) −→ WΩ(T) respectively. Each of these maps has in its image the corresponding element of Dend, hence x(Face) is in the image of Sk2 (x).
Proposition 3.4 Let T ∈ Ω be a tree such that | Vert(T)| ≥ 3 . If t, s ∈ wCatT1 satisfy
Sk2 (t) = Sk2 (s) then t = s .
Proof We proceed by induction on n = | Vert(T)|, the case n = 3 is covered in
Lemma 3.3. Suppose that x : WΩ(T) −→ Ctg is a dendrex of shape T . First we notice that we only need to describe the functor xσ : H int(σ) −→ Ctg(σ) in terms of the Skn−1 (x) where σ is the ordered sequence of colours (Leaves(T̄); root(T̄)) for a chosen planar representative T̄ . (The other components of x are already contained in the image of Skn−1 (x).)
The domain of xσ is a groupoid with the shape of an n-cube, having in its vertices the trivial categories (ǫ1 , . . . , ǫn ), ǫi ∈ {0, 1}. Denote by H̄k the full subcategory of
H int(σ) spanned by the categories (ǫ1 , . . . , ǫk−1 , 1, ǫk+1 , . . . , ǫn ), ǫi ∈ {0, 1} (one of the hyperfaces of the n-cube, containing the vertex (1,1,. . . ,1)). Denote by φk the arrow
(0, 0, . . . , 0) −→ (0, . . . , 0, 1, 0, . . . , 0) of H int(σ) (one of the edges of the n-cube, starting in (0, 0, . . . , 0)). Define the sets
Opd := {H̄k |k = 1, 2, . . . n} and

Face := {φk |k = 1, 2, . . . n}.

Since the “convex hull” of Opd ∪ Face is the whole domain of xσ , it is enough to prove that xσ (Opd) is completely determined by Skn−1 (x) and xσ (Face) is in the image of
Skk (x). Both of these assertions are true, by similar arguments to the ones in the proof of Lemma 3.3.
The following propositions of [14] and [11] helps us in proving that wCat1 is 3coskeletal.
Proposition 3.5 (Proposition 3.2.5 in [14]) Let X be a dendroidal set and k ≥ 2 an integer. If X satisfies the strict inner Kan condition for all trees T of degree at least k , then X is k -coskeletal.
Proposition 3.6 (Proposition 7.2 in [11]) Let P be a locally fibrant operad in E (that is, for any ordered sequence of objects σ = (c1 , . . . , cn ; c) the category P(σ) is fibrant with respect to the folk model structure). Then hcNd (P) is an inner Kan complex.
Corollary 3.7 (Lemma 4.6.3 in [14]) The dendroidal set wCat1 is 3 -coskeletal.

20

Andor Lukács

Proof In view of Proposition 3.5 it is enough to prove that wCat1 satisfies the strict inner Kan condition for all trees with | Vert(T)| ≥ 3. Let T be such a tree. Theorem 3.6
implies that wCat1 is an inner Kan complex, hence every inner horn Λe [T] −→ wCat1
has at least one filler t. Suppose that s is an other filler of the same horn. Since
| Vert(T)| ≥ 3, it follows that Sk2 (t) = Sk2 (s). We conclude thus by Proposition 3.4
that t = s.
Corollary 3.7 implies that wCatn is 3-coskeletal for every n ≥ 1. (Note that wCat0 is already 2-coskeletal.) To prove this, the following lemma is needed.
Lemma 3.8 If X is a k -coskeletal dendroidal set and Z is an arbitrary dendroidal set then dSets(Z, X) is k -coskeletal.
Proof The goal is to see that for any dendroidal set Y there exists a natural bijection dSets(Y, dSets(Z, X)) ≃ dSets(Skk Y, dSets(Z, X)).
Indeed, once one observes that Skk (Y ⊗ Z) ⊆ (Skk Y) ⊗ Z , one can conclude that there are natural one-to-one correspondences between the following Hom sets: dSets(Y, dSets(Z, X)) ≃ dSets(Y ⊗ Z, X)
≃ dSets(Y ⊗ Z, coSkk X)
≃ dSets(Skk (Y ⊗ Z), X)
≃ dSets((Skk Y) ⊗ Z, X)
≃ dSets(Skk Y, dSets(Z, X)).

Theorem 3.9 For every n ≥ 1 the dendroidal set wCatn is 3 -coskeletal.
Proof We proceed by induction on n. It was proven in Corollary 3.7 that wCat1 is
3-coskeletal. Suppose that wCatn is 3-coskeletal. It follows from Lemma 3.8 that for any set A, the dendroidal set Cat(wCatn )A = dSets(Nd (AsA ), wCatn ) is 3-coskeletal.
Hence Proposition 2.5 implies that
Z
n+1
Cat(wCatn )−
wCat
=
Sets

is 3-coskeletal.

Dendroidal weak 2-categories

21

4 Dendroidal weak 1- and 2-categories
4.1 Weak 1-categories
We start the Section with the description of the dendroidal set wCat1 = hcNd (Ctg).
We can use Corollary 3.7 to come to the conclusion that it is enough to describe the sets (wCat1 )T = OpE (WΩ(T), Ctg) for trees T with at most 3 vertices. Before we start with the description, let us make a useful notational convention: from now on, given n categories X1 , . . . , Xn and integers 1 ≤ i ≤ j ≤ n, (X)ji will denote the category
Xi × · · · × Xj .
(1) The first choice of T is the tree |. In this case WΩ(T) = Ω(|) is the operad on one object and only the identity operation, hence an element of (wCat1 )| is the same as the choice of a category.
(2) Let T = Corn , the n-corolla. In this case still WΩ(Corn ) = Ω(Corn ), hence an element of (wCat1 )Corn is the same as the choice of n + 1 categories X1 , . . . , Xn and X , together with a functor F : (X)n1 −→ X . Note that in case n = 0, the
E -enriched operad structure on Ctg implies that (X)n1 has to be considered the unit of the E -enriched monoidal category Ctg. This unit is the category ∗ on one object and no other arrows than the identity. Hence we infer that a dendrex of shape Cor0 amounts to the choice of a category X , together with an object of it.
(3) Let T = Corn ◦i Corm . Let us give a detailed description of maps of operads
α : WΩ(T) −→ Ctg since this is the first time when the interval H plays a role in the definition of the operad WΩ(T). So far it is clear that, as in cases (1) and
(2), such an α determines
(3a) a choice of n + 1 categories X1 , . . . , Xn , X together with a functor F1 :
(X)n1 −→ X ;
(3b) a choice of m categories Y1 , . . . , Ym and a functor F2 : (Y)m
1 −→ Xi .
There is one more building part of such an α, which is a functor

m n
H −→ Ctg (X)i−1
1 × (Y)1 × (X)i+1 , X .

But such a functor contains exactly the same data as the choice of two functors m
n
′
G, G′ : (X)i−1
1 × (Y)1 × (X)i+1 −→ X and a natural isomorphism φ : G −→ G .
The only thing we have not covered yet with the investigation of such a dendrex

22

Andor Lukács

is that α is a map of operads, which means that the diagram of categories


α×α /
∗×∗
Ctg (X)n1 , X × Ctg (Y)m
1 , Xi
◦i

◦i



H

α


/ Ctg (X)i−1 × (Y)m × (X)n , X
1
i+1
1



is commutative. One can spell out that this yields to G′ = F1 ◦i F2 . We can conclude thus that the last bit of information α provides is n
× (Y)m
(3c) a choice of a functor G : (X)i−1
1 × (X)i+1 −→ X and a natural
1
isomorphism φ : G −→ F1 ◦i F2 .

For the remaining choices of the tree T we give only the result.
(4) Let T = Corn ◦i (Corm ◦j Cork ). A map of operads WΩ(T) −→ Ctg is the same as
(4a) a choice of n + 1 categories X1 , . . . , Xn , X together with a functor F1 :
(X)n1 −→ X ;
(4b) a choice of m categories Y1 , . . . , Ym and a functor F2 : (Y)m
1 −→ Xi ; k
(4c) a choice of k categories Z1 , . . . , Zk and a functor F3 : (Z)1 −→ Yj ; n
× (Y)m
(4d) a choice of a functor G1 : (X)i−1
1 × (X)i+1 −→ X and a natural
1
isomorphism φ1 : G1 −→ F1 ◦i F2 ;
(4e) a choice of a functor G2 : (Y)j−1
× (Z)k1 × (Y)m j+1 −→ Xi and a natural
1
isomorphism φ2 : G2 −→ F2 ◦j F3 ; j−1
k m
n
(4f) a choice of a functor K : (X)i−1
1 × (Y)1 × (Z)1 × (Y)j+1 × (X)i+1 −→ X
and two natural isomorphisms
ψ1 : K −→ F1 ◦i G2 ,

ψ2 : K −→ G1 ◦j̃ F3

where j̃ = i+j−1, such that the following diagram of natural isomorphisms is commutative:
K

ψ1

/ F1 ◦i G2
F1 ◦i φ2

ψ2



G1 ◦ĩ F3

φ1 ◦j̃ F3



/ F1 ◦i F2 ◦j F3

(5) Let T = Corn ◦i,j (Corm , Cork ) for some 1 ≤ i < j ≤ n. A map of operads
WΩ(T) −→ Ctg is the same as
(5a) a choice of n + 1 categories X1 , . . . , Xn , X together with a functor F1 :
(X)n1 −→ X ;

Dendroidal weak 2-categories

23

(5b) a choice of m categories Y1 , . . . , Ym and a functor F2 : (Y)m
1 −→ Xi ; k
(5c) a choice of k categories Z1 , . . . , Zk and a functor F3 : (Z)1 −→ Xj ; n
× (Y)m
(5d) a choice of a functor G1 : (X)i−1
1 × (X)i+1 −→ X and a natural
1
isomorphism φ1 : G1 −→ F1 ◦i F2 ;
(5e) a choice of a functor G2 : (X)j−1
× (Z)k1 × (X)nj+1 −→ X and a natural
1
isomorphism φ2 : G2 −→ F2 ◦j F3 ; j−1
k n
× (Y)m
(5f) a choice of a functor K : (X)i−1
1 × (X)i+1 × (Z)1 × (X)j+1 −→ X
1
and two natural isomorphisms
ψ1 : K −→ F1 ◦i G2 ,

ψ2 : K −→ G1 ◦j̃ F3

where j̃ = i + j − 1 (we suppose j > i), such that the following diagram of natural isomorphisms is commutative:
K

ψ1

/ F1 ◦i G2
F1 ◦i φ2

ψ2



G1 ◦ĩ F3

φ1 ◦j̃ F3



/ F1 ◦i,j (F2 , F3 )

We are going to illustrate with some examples the dendroidal structure of wCat1 in the context described above. Let T = Corn ◦i Corm , hence a dendrex α of shape T is the same thing as the data described in (3) above. If ∂ : Corn −→ T is the obvious outer face of T then ∂ ∗ (α) corresponds to the choice of the categories X1 , . . . , Xn , X and the functor F1 : (X)n1 −→ X. If ∂ : Corn+m−1 −→ T is the inner face of T then ∂ ∗ (α)
corresponds to the choice of the categories X1 , . . . , Xi−1 , Y1 , . . . , Ym , Xi+1 , . . . , Xn , X
m n
and the functor G : (X)i−1
1 × (Y)1 × (X)i+1 −→ X . (The choice of G instead of F1 ◦i F2
follows from the definition of the map of operads ∂ : WΩ(Corn+m−1 ) −→ WΩ(T).)
One can similarly decipher what a degeneracy looks like. A simple case of such occurs when R = Corn ◦i Cor1 , T = Corn and σ : R −→ T is the degeneracy in question. If
β is a dendrex of shape T , that is a choice of categories X1 , . . . , Xn , X and a functor
F1 : (X)n1 −→ X , then σ ∗ (β) adds to the information contained in β the identity functor id : Xi −→ Xi , and the identity natural transformation F1 −→ F1 ◦i id.

4.2 Weak 2-categories
We turn our attention now to the dendroidal set wCat2 . Our goal is to unpack the definition and compare the result with bicategories. It will become apparent later that the right notion to compare the data of wCat2 contained in lower degrees is unbiased

24

Andor Lukács

bicategories and their homomorphisms. These notions were defined by Tom Leinster in [8]. We recall them in the Appendix.
2
The Section is organized as follows: First we analyse the sets wCat|2 , wCatCor and
1
their relations to unbiased bicategories, and we prove that the category of unbiased bicategories is isomorphic to the homotopy category of dendroidal weak 2-categories.
Then e conclude the Section by a conjecture that predicts a stronger relation between bicategories and dendroidal weak 2-categories.

4.2.1

Dendroidal weak 2-categories

In this Subsection we analyse those components of the dendroidal set wCat2 which will correspond to bicategories and homomorphisms of bicategories.

4.2.2

Dendrices of shape |

Since
(wCat2 )| =

Z

Cat(hcNd (Ctg))−

Sets



,
|

the definition of the Grothendieck construction implies that an element of (wCat2 )|
is a pair (A, x), where A is a set and x is a dendrex of shape | in the dendroidal set
Cat(hcNd (Ctg))A . Hence x ∈ dSets(Nd (AsA ) ⊗ Ω[|], hcNd (Ctg)) = dSets(Nd (AsA ), hcNd (Ctg)).
Since hcNd (Ctg) is 3-coskeletal, it is enough to look at the degree 0, 1, 2 and 3
components of x.
(0) The degree 0 component of x is the map of sets x| : Nd (AsA )| −→ hcNd (Ctg)| .
Since Nd (AsA )| consists of the objects of the operad AsA and hcNd (Ctg)| consists of categories, it follows that x| is the same thing as the choice of a category
A(a1 , a2 ) for each ordered pair (a1 , a2 ) ∈ A × A.
(1) Let us look at the xCorn component, n ∈ N. There are three cases to distinguish.
First, an element in Nd (AsA )Cor0 consists of a pair (a, a) where a ∈ A, and the
; (a, a) . We have seen in Subsection 4.1 that an element operation ∗ ∈ AsA
of hcNd (Ctg)Cor 0 is a category together with an object of it. Since x has to be

Dendroidal weak 2-categories

25

compatible with the face map | −→ Cor0 , it follows that xCor0 picks for each a ∈ A a functor Ψa : ∗ −→ A(a, a).
Second, an element in Nd (AsA )Cor1 consists of a pair (a1 , a2 ) ∈ A2 and the operation ∗ ∈ AsA (a1 , a2 ); (a1 , a2 ) , which is also the corresponding unit operation in the operad AsA . An element of hcNd (Ctg)Cor 1 is a functor between two chosen categories. Again, since x has to be compatible with the various face and degeneracy maps, it follows that xCor1 amounts to choosing the identity functor on every already chosen category A(a1 , a2 ). Hence xCor1 does not contribute with any new information.
Third, for n ≥ 2 xCorn picks for each admissible ordered sequence

σ = (a1 , a2 ), (a2 , a3 ), . . . , (an−1 , an ); (a1 , an )
of n + 1 objects of AsA a functor

Ψσ : A(a1 , a2 ) × · · · × A(an−1 , an ) −→ A(a1 , an ).
We can include the cases n = 0, 1 in the third one in the obvious way.
(2) The degree 2 component of x consists of xT where T = Corn ◦i Corm for the various n, m, i ∈ N, n 6= 0. There are face maps into the tree T from the m, n and m + n − 1 corollas, and in case n = 1 or m = 1 there are also degeneracy maps with T as the domain. Since x has to be compatible with these faces and degeneracies, we can conclude that xT provides the following bit of extra data:
For any pair of admissible ordered sequences

σ = (a1 , a2 ), . . . , (an−1 , an ); (a1 , an ) ,

ρ = (ai , b2 ), (b2 , b3 ), . . . , (bm−1 , ai+1 ); (ai , ai+1 )
and any 1 ≤ i ≤ n a natural isomorphism

φσ,ρ,i : Ψσ◦i ρ −→ Ψσ ◦i Ψρ ,
There is one condition on these natural isomorphisms: in case n = 1 or m = 1, the corresponding natural isomorphism has to be the identity (it follows from the compatibility with degeneracies again).
(3) The degree 3 components of x do not give rise to any extra data, but the dendroidal identities with face maps in

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
