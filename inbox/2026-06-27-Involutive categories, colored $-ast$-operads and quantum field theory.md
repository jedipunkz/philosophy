---
source: "https://arxiv.org/abs/1802.09555v2"
title: "Involutive categories, colored $\\ast$-operads and quantum field theory"
author: "Marco Benini, Alexander Schenkel, Lukas Woike"
year: "2018"
publication: "arXiv preprint / math.CT"
download: "https://arxiv.org/pdf/1802.09555v2"
pdf: "https://arxiv.org/pdf/1802.09555v2"
captured_at: "2026-06-27T06:15:36Z"
updated_at: "2026-06-27T06:15:36Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács Theory of the Novel"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# Involutive categories, colored $\ast$-operads and quantum field theory

- 著者: Marco Benini, Alexander Schenkel, Lukas Woike
- 年: 2018
- 掲載情報: arXiv preprint / math.CT
- 情報源: [arxiv](https://arxiv.org/abs/1802.09555v2)
- ダウンロード: https://arxiv.org/pdf/1802.09555v2
- PDF: https://arxiv.org/pdf/1802.09555v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

Involutive category theory provides a flexible framework to describe involutive structures on algebraic objects, such as anti-linear involutions on complex vector spaces. Motivated by the prominent role of involutions in quantum (field) theory, we develop the involutive analogs of colored operads and their algebras, named colored $\ast$-operads and $\ast$-algebras. Central to the definition of colored $\ast$-operads is the involutive monoidal category of symmetric sequences, which we obtain from a general product-exponential $2$-adjunction whose right adjoint forms involutive functor categories. For $\ast$-algebras over $\ast$-operads we obtain involutive analogs of the usual change of color and operad adjunctions. As an application, we turn the colored operads for algebraic quantum field theory into colored $\ast$-operads. The simplest instance is the associative $\ast$-operad, whose $\ast$-algebras are unital and associative $\ast$-algebras.

## PDF Text

Involutive categories, colored ∗-operads and quantum field theory
Marco Benini1,a , Alexander Schenkel2,b and Lukas Woike1,c
1

arXiv:1802.09555v2 [math.CT] 11 Feb 2019

2

Fachbereich Mathematik, Universität Hamburg,
Bundesstr. 55, 20146 Hamburg, Germany.

School of Mathematical Sciences, University of Nottingham,
University Park, Nottingham NG7 2RD, United Kingdom.
Email:

a

marco.benini@uni-hamburg.de alexander.schenkel@nottingham.ac.uk c
lukas.jannik.woike@uni-hamburg.de b

January 2019

Abstract
Involutive category theory provides a flexible framework to describe involutive structures on algebraic objects, such as anti-linear involutions on complex vector spaces. Motivated by the prominent role of involutions in quantum (field) theory, we develop the involutive analogs of colored operads and their algebras, named colored ∗-operads and ∗-algebras. Central to the definition of colored ∗-operads is the involutive monoidal category of symmetric sequences, which we obtain from a general product-exponential 2-adjunction whose right adjoint forms involutive functor categories. For ∗-algebras over ∗-operads we obtain involutive analogs of the usual change of color and operad adjunctions. As an application, we turn the colored operads for algebraic quantum field theory into colored ∗-operads. The simplest instance is the associative ∗-operad, whose ∗-algebras are unital and associative ∗-algebras.

Report no.:

ZMP-HH/18-6, Hamburger Beiträge zur Mathematik Nr. 725

Keywords: involutive categories, involutive monoidal categories, ∗-monoids, colored operads,
∗-algebras, algebraic quantum field theory
MSC 2010:

18Dxx, 81Txx

Contents
1 Introduction and summary

2

2 Involutive categories
2.1 Basic definitions and properties . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
2.2 ∗-objects . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .

4
4
6

3 Involutive structures on monoidal categories
9
3.1 (Symmetric) monoidal categories and monoids . . . . . . . . . . . . . . . . . . . . 9
3.2 Involutive (symmetric) monoidal categories . . . . . . . . . . . . . . . . . . . . . . 10
3.3 ∗-monoids . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 14
1

4 Involutive structures on colored symmetric sequences
4.1 Product-exponential 2-adjunction . . . . . . . . . . . . . . . . . . . . . . . . . . . .
4.2 Involutive colored symmetric sequences . . . . . . . . . . . . . . . . . . . . . . . . .
4.3 ∗-objects . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .

16
17
19
21

5 Colored ∗-operads

22

6 ∗-algebras over colored ∗-operads

24

7 Algebraic quantum field theory ∗-operads

28

1

Introduction and summary

In ordinary category theory, an involution on an object c ∈ C of a category C is an endomorphism i : c → c that squares to the identity, i.e. i2 = idc . Unfortunately, this concept is too rigid to describe many examples of interest. For instance, given an associative and unital ∗-algebra A
over C, e.g. the algebra of observables of a quantum system, the involution ∗ : A → A on its underlying vector space is not an endomorphism in the category of complex vector spaces, but rather a complex anti-linear map.
Involutive categories [BM09, Egg11, Jac12] were developed in order to introduce the flexibility required to resolve this insufficiency. Their definition is a particular instance of the “microcosm principle” of Baez and Dolan [BD98], which states that certain algebraic structures can be defined in any category equipped with a categorified version of the same structure. Hence, an involutive category is a category C equipped with an endofunctor J : C → C that squares to the identity endofunctor IdC , up to a given natural isomorphism j : IdC → J 2 which has to satisfy certain coherence conditions (cf. Definition 2.1). In an involutive category (C, J, j), one can introduce a more flexible concept of involution on an object c ∈ C, which is given by a C-morphism
∗ : c → Jc satisfying (J∗) ∗ = jc as morphisms from c to J 2 c (cf. Definition 2.14). Such objects
(homotopy fixed points, as a matter of fact) are called self-conjugates in [Jac12], involutive objects in [Egg11] and ∗-objects in [BM09]. We shall follow the latter terminology because it seems the most natural one to us. If a category is equipped with its trivial involutive structure J = IdC and j = idIdC (cf. Example 2.2), then ∗-objects are just endomorphisms squaring to the identity, i.e.
the ordinary involutions mentioned above. This framework, however, becomes much richer and flexible by allowing for non-trivial involutive structures: For example, endowing the category of complex vector spaces VecC with the involutive structure given by the endofunctor that assigns to a complex vector space V its complex conjugate vector space V , the complex anti-linear map underlying a ∗-algebra may be regarded as a ∗-object ∗ : A → A in this involutive category (cf.
Examples 2.3 and 2.17).
The observables of a quantum system form a unital and associative ∗-algebra over C. This shows the relevance of involutive categories for general quantum theory, quantum field theory and also noncommutative geometry. Our main motivation for this paper stems precisely from these areas and more specifically from our recent operadic approach to algebraic quantum field theory
[BSW17]. There the axioms of algebraic quantum field theory [HK64, BFV03] are encoded in a colored operad and generalized to richer target categories, such as chain complexes and other symmetric monoidal categories, which are central in modern approaches to quantum gauge theories [CG17, BSS15, BS17, BSW17, BSW18, Yau18]. For their physical interpretation, however, it is essential that quantum systems such as quantum field theories come equipped with involutions.
These enable us to perform the GNS construction and recover the usual probabilistic interpretation of quantum theory. We refer to [Jac12] for a generalization of the GNS construction to involutive symmetric monoidal categories.
The purpose of this paper is to combine the theory of colored operads and that of involutive
2

categories, resulting in what we shall call colored ∗-operads. Despite of our quite concrete motivation, we believe that working out the theory of colored ∗-operads in full generality provides an interesting and valuable addition to the largely unexplored field of involutive category theory. On the one hand, our constructions naturally lead to interesting new structures such as involutive functor categories, which have not been discussed in the literature. On the other hand, our study of involutive structures on the category of symmetric sequences, which is a monoidal category that does not admit a braiding, provides an interesting example of an involutive monoidal category in the sense of [Jac12], but not in the sense of [BM09, Egg11], see Remark 4.6 for details.
This shows that Jacobs’ definition of involutive monoidal categories is the one suitable to develop the theory of colored ∗-operads, consequently we shall use this one in our paper.
The outline of the paper is as follows: Sections 2 and 3 contain a brief review of involutive categories and involutive (symmetric) monoidal categories following mostly [Jac12]. We shall in particular emphasize and further develop the 2-categorical aspects of this theory, including the
2-functorial behavior of the assignments of the categories of ∗-objects and ∗-monoids. For the sake of concreteness, we also describe the most relevant constructions and definitions arising this way in fully explicit terms. Theorems 2.23 and 3.17 establish simple criteria that are useful to detect whether an involutive ((symmetric) monoidal) category is isomorphic to one with a trivial involutive structure. In Section 4 we show that the category of colored symmetric sequences, which underlies colored operad theory, carries a canonical involutive monoidal structure in the sense of [Jac12], but not in the sense of [BM09, Egg11]. The relevant involutive structure is obtained by employing a general construction, namely exponentiation of involutive categories, which results in involutive structures on functor categories. Colored ∗-operads with values in any cocomplete involutive closed symmetric monoidal category (M, J, j) are defined in Section 5 as
∗-monoids in our involutive monoidal category of colored symmetric sequences. In Proposition
5.4 we shall prove that the resulting category is isomorphic to the category of ordinary colored operads with values in the category of ∗-objects in (M, J, j), which provides an alternative point of view on colored ∗-operads. The possibility to switch between these equivalent perspectives is useful for concrete applications and also to import techniques from ordinary operad theory to the involutive setting. In Section 6 we introduce and study the category of ∗-algebras over colored
∗-operads. In particular, we prove that a change of colored ∗-operad induces an adjunction between the associated categories of ∗-algebras, which generalizes the corresponding crucial and widely used result from ordinary to involutive category theory. Finally, in Section 7 we endow the algebraic quantum field theory operads constructed in [BSW17] with a canonical order-reversing structure of colored ∗-operads and provide a characterization of the corresponding categories of
∗-algebras. As a simple example, we obtain a ∗-operad structure on the associative operad and show that its ∗-algebras behave like ∗-algebras over C in the sense that the involution reverses the order of multiplication (a b)∗ = b∗ a∗ . It is essential to emphasize that this order-reversal is encoded in our ∗-operad structure. This is radically different from the approach of [BM09, Egg11], whose definition of an involutive monoidal category prescribes that the endofunctor J reverses the monoidal structure up to natural isomorphism, thus recovering unital and associative ∗-algebras over C directly as ∗-monoids in VecC .
Notations: We denote categories by boldface letters like C, D and E. Objects in categories are indicated by c ∈ C and we write C(c, c′ ) for the set of morphisms from c to c′ in C. Functors are denoted by capital letters like F : C → C′ or X : D → C, and so are the identity functors
IdC : C → C. Natural transformations are denoted by Greek letters like ζ : F → G or α : X → Y .
Given functors K : D′ → D, X : D → C and J : C → C′ , we denote their composition simply by juxtaposition JXK : D′ → C′ . Given also a natural transformation α : X → Y of functors
X, Y : D → C, we denote by
JαK : JXK −→ JY K

3

(1.1a)

the whiskering of J, α and K. Explicitly, JαK is the natural transformation with components
(JαK)d′ = JαKd′ : JXKd′ −→ JY Kd′

,

for all d′ ∈ D′ . For β : Y → Z another natural transformation, one easily confirms that

(JβK) (JαK) = J βα K : JXK −→ JZK ,

(1.1b)

(1.2)

where (vertical) composition of natural transformations is also denoted by juxtaposition. We shall need some basic elements of (strict) 2-category theory, for which we refer to [KS74].

2

Involutive categories

This section contains a brief review of involutive categories. We shall mostly follow the definitions and conventions of Jacobs [Jac12] and refer to this paper for more details and some of the proofs.
We strongly emphasize and also develop further the 2-categorical aspects of involutive category theory established in [Jac12], which will be relevant for the development of our present paper.
When it comes to notations and terminology, we sometimes prefer the work of Beggs and Majid
[BM09] and the one of Egger [Egg11].

2.1

Basic definitions and properties

Definition 2.1. An involutive category is a triple (C, J, j) consisting of a category C, an endofunctor J : C → C and a natural isomorphism j : IdC → J 2 satisfying jJ = Jj : J −→ J 3

.

(2.1)

Example 2.2. For any category C, the triple (C, IdC , idIdC ) defines an involutive category. We call this the trivial involutive category over C.
▽
Example 2.3. Let VecC be the category of complex vector spaces. Consider the endofunctor
(−) : VecC → VecC that assigns to any V ∈ VecC its complex conjugate vector space V ∈ VecC
and to any C-linear map f : V → W the canonically induced C-linear map f : V → W . Notice that (−) = IdVecC , hence the triple (VecC , (−), idIdVecC ) is an involutive category.
▽
Example 2.4. Let C be any non-empty set and ΣC the associated groupoid of C-profiles. The objects of ΣC are finite sequences c = (c1 , . . . , cn ) of elements in C, including also the empty sequence ∅ ∈ ΣC . We denote by |c| = n the length of the sequence. The morphisms of ΣC are right permutations σ : c → cσ := (cσ(1) , . . . , cσ(n) ), with σ ∈ Σ|c| in the symmetric group on |c|
letters. We define an endofunctor Rev : ΣC → ΣC as follows: To an object c = (c1 , . . . , cn ) ∈ ΣC
it assigns the reversed sequence
Rev(c) := c ρ|c| := (cn , . . . , c1 ) ,

(2.2a)

where ρ|c| ∈ Σ|c| denotes the order-reversal permutation. To a ΣC -morphism σ : c → cσ it assigns the right permutation
Rev(σ) := ρ|c| σ ρ|c| : Rev(c) −→ Rev(cσ) ,

(2.2b)

where we also used that |cσ| = |c|. Notice that Rev2 = IdΣC , hence the triple (ΣC , Rev, idIdΣC ) is an involutive category.
▽
The following very useful result appears in [Jac12, Lemma 1].
Lemma 2.5. For every involutive category (C, J, j), the endofunctor J : C → C is self-adjoint, i.e. J ⊣ J. As a consequence, J preserves all limits and colimits that exist in C.
4

Definition 2.6. An involutive functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ) consists of a functor F :
C → C′ and a natural transformation ν : F J → J ′ F satisfying
F

F

(2.3)
j′F

Fj



F J2

νJ

/ J ′F J

J ′ν


/ J ′2 F

An involutive natural transformation ζ : (F, ν) → (G, χ) between involutive functors (F, ν), (G, χ) :
(C, J, j) → (C′ , J ′ , j ′ ) is a natural transformation ζ : F → G satisfying
FJ

ζJ

/ GJ

(2.4)

χ

ν





J ′F

/ J ′G

J ′ζ

Proposition 2.7. Involutive categories, involutive functors and involutive natural transformations form a 2-category ICat.
Remark 2.8. Let us describe the 2-category structure on ICat explicitly.
(i) For any involutive category (C, J, j), the identity involutive functor is given by Id(C,J,j) :=
(IdC , idJ ) : (C, J, j) → (C, J, j).
(ii) Given two involutive functors (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ) and (F ′ , ν ′ ) : (C′ , J ′ , j ′ ) →
(C′′ , J ′′ , j ′′ ), their composition is given by

(F ′ , ν ′ ) (F, ν) := F ′ F, (ν ′ F ) (F ′ ν) : (C, J, j) −→ (C′′ , J ′′ , j ′′ ) .
(2.5)
(iii) Vertical/horizontal composition of involutive natural transformations is given by vertical/horizontal composition of their underlying natural transformations. (It is easy to verify that the latter compositions define involutive natural transformations.)
△
The following technical lemma is proven in [Jac12, Lemma 2].
Lemma 2.9. For every involutive functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ), the natural transformation ν : F J → J ′ F is a natural isomorphism.
As in any 2-category, there exists the concept of adjunctions in the 2-category ICat.
Definition 2.10. An involutive adjunction
/

(L, λ) : (C, J, j) o

(D, K, k) : (R, ρ)

(2.6)

consists of two involutive functors (L, λ) : (C, J, j) → (D, K, k) and (R, ρ) : (D, K, k) → (C, J, j)
together with two involutive natural transformations η : Id(C,J,j) → (R, ρ) (L, λ) (called unit) and
ǫ : (L, λ) (R, ρ) → Id(D,K,k) (called counit) that satisfy the triangle identities
(R, ρ) ❚❚

η (R,ρ)

/ (R, ρ) (L, λ) (R, ρ)
❚❚❚❚
❚❚❚❚
❚❚
(R,ρ) ǫ
id(R,ρ) ❚❚❚❚❚

❚)

(L, λ) ❚❚

(L,λ) η

/ (L, λ) (R, ρ) (L, λ)
❚❚❚❚
❚❚❚❚
❚❚
ǫ (L,λ)
id(L,λ) ❚❚❚❚❚

❚)

(R, ρ)

(L, λ)

We also denote involutive adjunctions simply by (L, λ) ⊣ (R, ρ).
5

(2.7)

Remark 2.11. Applying the forgetful 2-functor ICat → Cat, every involutive adjunction
(L, λ) ⊣ (R, ρ) defines an ordinary adjunction L ⊣ R in the 2-category of categories Cat. Notice that an involutive adjunction is the same thing as an ordinary adjunction L ⊣ R (between categories equipped with an involutive structure) whose functors L and R are equipped with involutive structures that are compatible with the unit and counit in the sense that the latter become of involutive natural transformations. This alternative point of view will be useful in
Corollary 4.7 and Theorem 6.6 below, where we make use of the construction in the following proposition.
△
Proposition 2.12. Let (R, ρ) : (D, K, k) → (C, J, j) be an involutive functor and suppose that
L : C → D is a left adjoint to the functor R : D → C. Define a natural transformation λ by
λ

LJ

/ KL
O

LJη

(2.8)

ǫKL



LJRL

Lρ−1 L

/ LRKL

where η : IdC → RL and ǫ : LR → IdD are the unit and counit of the adjunction L ⊣ R. Then
(L, λ) ⊣ (R, ρ) is an involutive adjunction.
Proof. The above diagram defines a natural transformation λ because ρ is a natural isomorphism, cf. Lemma 2.9. A slightly lengthy diagram chase shows that (L, λ) : (C, J, j) → (D, K, k) is an involutive functor. Furthermore, by the definition of λ, the natural transformations η and ǫ are involutive natural transformations.
Remark 2.13. Even though we will not need it in the following, let us briefly mention that the dual of Proposition 2.12 also holds true: Let (L, λ) : (C, J, j) → (D, K, k) be an involutive functor and suppose that R : D → C is a right adjoint to the functor L : C → D. Then (L, λ) ⊣ (R, ρ)
is an involutive adjunction for ρ defined by
ρ−1

JR

/ RK
O

ηJR

RKǫ



RLJR

RλR

/ RKLR

where η : IdC → RL and ǫ : LR → IdD are the unit and counit of the adjunction L ⊣ R.

2.2

(2.9)

△

∗-objects

Definition 2.14. A ∗-object in an involutive category (C, J, j) is a C-morphism ∗ : c → Jc satisfying c ◆◆◆

∗

(2.10)

/ Jc

◆◆◆
◆◆◆
jc ◆◆◆◆
&

J∗



J 2c

A ∗-morphism f : (∗ : c → Jc) → (∗′ : c′ → Jc′ ) is a C-morphism f : c → c′ satisfying c

f

/ c′
∗′

∗





Jc

Jf

/ Jc′

We denote the category of ∗-objects in (C, J, j) by ∗-Obj(C, J, j).
6

(2.11)

Remark 2.15. For any ∗-object (∗ : c → Jc) ∈ ∗-Obj(C, J, j), the C-morphism ∗ : c → Jc is an isomorphism with inverse given by jc−1 J∗ : Jc → c.
△
Example 2.16. Consider the trivial involutive category (C, IdC , idIdC ) from Example 2.2. A ∗object consists of an object c ∈ C equipped with a C-endomorphism ∗ : c → c satisfying ∗2 = idc , i.e. an object equipped with an involution.
▽
Example 2.17. Consider the involutive category (VecC , (−), idIdVecC ) from Example 2.3. A ∗object consists of a complex vector space V equipped with a complex anti-linear map ∗ : V → V
satisfying ∗2 = idV .
▽
Example 2.18. Consider the involutive category (ΣC , Rev, idIdΣC ) from Example 2.4. A ∗-object consists of a C-profile c = (c1 , . . . , cn ) equipped with a right permutation ∗ : c → Rev(c) = c ρ|c|
satisfying ∗ρ|c| ∗ ρ|c| = e ∈ Σ|c| , where e denotes the identity permutation. In particular, any object c ∈ ΣC carries a canonical ∗-object structure given by ρ|c| : c → c ρ|c| . The assignment c 7→ (ρ|c| : c → c ρ|c| ) defines a functor ρ : ΣC → ∗-Obj(ΣC , Rev, idIdΣC ) that is a section of the
▽
forgetful functor U : ∗-Obj(ΣC , Rev, idIdΣC ) → ΣC .
For any involutive category (C, J, j), there exists a forgetful functor U : ∗-Obj(C, J, j) → C
specified by (∗ : c → Jc) 7→ c. If the category C has coproducts, we can define for any object c ∈ C a morphism
F (c) :=



id⊔jc
/ Jc ⊔ J 2 c ∼
c ⊔ Jc ∼
= J(c ⊔ Jc)
= Jc ⊔ c



(2.12)

in C, where in the last step we used that J preserves coproducts because of Lemma 2.5. One can easily check that (2.12) defines a ∗-object in (C, J, j), i.e. F (c) ∈ ∗-Obj(C, J, j). Another direct computation shows
Proposition 2.19. Let (C, J, j) be an involutive category that admits coproducts. The assignment c 7→ F (c) given by (2.12) naturally extends to a functor F : C → ∗-Obj(C, J, j), which is a left adjoint of the forgetful functor U : ∗-Obj(C, J, j) → C.
Remark 2.20. [Jac12, Lemma 5] shows that ∗-Obj(C, J, j) inherits all limits and colimits that exist in C. These are preserved by the forgetful functor U : ∗-Obj(C, J, j) → C.
△
As noted in [Jac12, Lemma 6], the assignment of the categories of ∗-objects extends to a
2-functor
∗-Obj : ICat −→ Cat .

(2.13)

Concretely, this 2-functor is given by the following assignment:
• an involutive category (C, J, j) is mapped to its category of ∗-objects ∗-Obj(C, J, j);
• an involutive functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ) is mapped to the functor ∗-Obj(F, ν) :
∗-Obj(C, J, j) → ∗-Obj(C′ , J ′ , j ′ ) that acts on objects as

∗-Obj(F, ν) ∗ : c → Jc :=

Fc

F∗ /

F Jc

νc /

J ′F c



(2.14)

and on morphisms as F ;
• an involutive natural transformation ζ : (F, ν) → (G, χ) is mapped to the natural transformation ∗-Obj(ζ) : ∗-Obj(F, ν) → ∗-Obj(G, χ) with components ∗-Obj(ζ)(∗:c→Jc) := ζc , for all (∗ : c → Jc) ∈ ∗-Obj(C, J, j).

7

Recalling the trivial involutive categories from Example 2.2, we obtain another 2-functor triv : Cat −→ ICat .

(2.15)

Concretely, this 2-functor assigns to a category C the trivial involutive category (C, IdC , idIdC ), to a functor F : C → C′ the involutive functor (F, idF ) : (C, IdC , idIdC ) → (C′ , IdC′ , idIdC′ ), and to a natural transformation ζ : F → G the involutive natural transformation ζ : (F, idF ) → (G, idG ).
Theorem 2.21. The 2-functors (2.13) and (2.15) form a 2-adjunction triv : Cat o

/

ICat : ∗-Obj

.

(2.16)

The unit η : IdCat → ∗-Obj triv and counit ǫ : triv ∗-Obj → IdICat 2-natural transformations are stated explicitly in the proof below.
Proof. The component at C ∈ Cat of the 2-natural transformation η is the functor

ηC : C −→ ∗-Obj triv(C)

(2.17)

that equips objects with their identity involution (cf. Example 2.16), i.e. c 7→ (idc : c → c). The component at (C, J, j) ∈ ICat of the 2-natural transformation ǫ is the involutive functor

ǫ(C,J,j) = (U, ν) : triv ∗-Obj(C, J, j) −→ (C, J, j) ,
(2.18)
where U : ∗-Obj(C, J, j) → C is the forgetful functor (∗ : c → Jc) 7→ c and its involutive structure
ν : U → JU is the natural transformation defined by the components ν(∗:c→Jc) = ∗ : c → Jc, for all (∗ : c → Jc) ∈ ∗-Obj(C, J, j). An elementary check shows that η and ǫ are indeed 2-natural transformations that satisfy the triangle identities, hence (2.16) is a 2-adjunction with unit η and counit ǫ.

Remark 2.22. Notice that both Cat and ICat carry a Cartesian monoidal structure, which is concretely given by the product categories C × D in Cat and the product involutive categories
(C, J, j) × (D, K, k) = (C × D, J × K, j × k) in ICat. Because ∗-Obj is a right adjoint functor, it follows that there are canonical isomorphisms

∗-Obj (C, J, j) × (D, K, k) ∼
(2.19)
= ∗-Obj(C, J, j) × ∗-Obj(D, K, k) , for all involutive categories (C, J, j) and (D, K, k).

△

We conclude this section with a useful result that allows us to detect involutive categories carrying a trivial involutive structure.
Theorem 2.23. Let (C, J, j) be an involutive category. Any section ∗ : C → ∗-Obj(C, J, j)
of the forgetful functor U : ∗-Obj(C, J, j) → C canonically determines an ICat-isomorphism between (C, J, j) and the trivial involutive category (C, IdC , idIdC ). In particular, if a section of
U exists, then the involutive categories (C, J, j) and (C, IdC , idIdC ) are isomorphic.
Proof. A section ∗ : C → ∗-Obj(C, J, j) of U assigns to each c ∈ C a ∗-object ∗c : c → Jc and to each C-morphism f : c → c′ a ∗-morphism c

f

/ c′

(2.20)

∗c′

∗c





Jc

Jf

/ Jc′

Notice that this diagram implies that ∗c are the components of a natural transformation ∗ :
IdC → J. It is straightforward to check that (IdC , ∗) : (C, IdC , idIdC ) → (C, J, j) is an involutive functor, which is invertible via the involutive functor (IdC , ∗−1 ) : (C, J, j) → (C, IdC , idIdC ).
Corollary 2.24. The involutive category (ΣC , Rev, idIdΣC ) of C-profiles equipped with reversal as involutive structure (cf. Examples 2.4 and 2.18) is isomorphic to the trivial involutive category
(ΣC , IdΣC , idIdΣC ).
8

3

Involutive structures on monoidal categories

In this section we review involutive (symmetric) monoidal categories and ∗-monoids therein. We again shall follow mostly the definitions and conventions of Jacobs [Jac12]. Our main goal is to clarify and work out the 2-functorial behavior of the assignment of the categories of ∗-objects and monoids to involutive (symmetric) monoidal categories. To fix our notations, we start with a brief review of some basic aspects of (symmetric) monoidal categories and monoids therein.

3.1

(Symmetric) monoidal categories and monoids

Recall that a monoidal category (C, ⊗, I, α, λ, ρ) consists of a category C, a functor ⊗ : C × C →
C, an object I ∈ C and three natural isomorphisms
α : ⊗ (⊗ × IdC ) −→ ⊗ (IdC × ⊗) ,

(3.1a)

λ : I ⊗ (−) −→ IdC

,

(3.1b)

ρ : (−) ⊗ I −→ IdC

,

(3.1c)

which satisfy the pentagon and triangle identities. We follow the usual abuse of notation and often denote a monoidal category by its underlying category C. The associator α and the unitors λ and
ρ will always be suppressed. Given two monoidal categories C and C′ , a (lax) monoidal functor from C to C′ is a triple (F, F2 , F0 ) consisting of a functor F : C → C′ , a natural transformation
F2 : ⊗′ (F × F ) −→ F ⊗

,

(3.2a)

and a C′ -morphism
F0 : I ′ −→ F I

,

(3.2b)

which are required to satisfy the usual coherence conditions involving the associators and unitors.
We often denote a monoidal functor by its underlying functor F : C → C′ . A monoidal natural transformation ζ : F → G between monoidal functors F = (F, F2 , F0 ) and G = (G, G2 , G0 ) is a natural transformation ζ : F → G satisfying
⊗′ (F × F )

⊗′ (ζ×ζ)

/ ⊗′ (G × G)
G2

F2





F⊗

/ G⊗

ζ⊗

FI

⑦
F0 ⑦⑦
⑦
⑦
⑦
~⑦
⑦

I′ ❅

ζI

❅❅
❅❅G0
❅❅
❅
/ GI

(3.3)

Proposition 3.1. Monoidal categories, (lax) monoidal functors and monoidal natural transformations form a 2-category MCat.
A symmetric monoidal category is a monoidal category C together with a natural isomorphism called braiding
τ : ⊗ −→ ⊗op := ⊗ σ

(3.4)

from the tensor product to the opposite tensor product, where σ : C × C → C × C is the flip functor (c1 , c2 ) 7→ (c2 , c1 ), which satisfies the hexagon identities and the symmetry constraint
⊗❇

id⊗

❇❇
❇❇
τ ❇❇❇
!

⊗σ

/ ⊗ = ⊗ σ2
9
sss s
ss ss τ σ
ss

9

(3.5)

We often denote a symmetric monoidal category by its underling category C. A symmetric monoidal functor is a monoidal functor F : C → C′ that preserves the braidings, i.e.
⊗′ (F × F )

τ ′ (F ×F )

/ ⊗′ σ(F × F ) = ⊗′ (F × F )σ

F2

(3.6)

F2 σ





F⊗

/F ⊗σ

Fτ

commutes. A symmetric monoidal natural transformation is just a monoidal natural transformation between symmetric monoidal functors.
Proposition 3.2. Symmetric monoidal categories, symmetric monoidal functors and symmetric monoidal natural transformations form a 2-category SMCat.
Definition 3.3. A monoid in a (symmetric) monoidal category C is a triple (M, µ, η) consisting of an object M ∈ C and two C-morphisms µ : M ⊗ M → M (called multiplication) and η :
I → M (called unit) satisfying the associativity and unitality axioms. A monoid morphism f : (M, µ, η) → (M ′ , µ′ , η ′ ) is a C-morphism f : M → M ′ preserving multiplications and units.
We denote the category of monoids in C by Mon(C).
The assignment of the categories of monoids extends to a 2-functor
Mon : (S)MCat −→ Cat .

(3.7)

Concretely, this 2-functor is given by the following assignment:
• a (symmetric) monoidal category C is mapped to its category of monoids Mon(C);
• a (symmetric) monoidal functor F : C → C′ is mapped to the functor Mon(F ) : Mon(C) →
Mon(C′ ) that acts on objects as


F2 M,M
/ F (M ⊗ M )
Mon(F ) M, µ, η := F M, F M ⊗′ F M

Fµ

/ F M, I ′ F0 / F I F η / F M



(3.8)

and on morphisms as F ;
• a (symmetric) monoidal natural transformation ζ : F → G is mapped to the natural transformation Mon(ζ) : Mon(F ) → Mon(G) with components Mon(ζ)(M,µ,η) := ζM , for all (M, µ, η) ∈ Mon(C).

3.2

Involutive (symmetric) monoidal categories

The following definition of an involutive (symmetric) monoidal category is due to [Jac12]. We prefer this definition over the one in [Egg11, BM09] as it has the advantage that the category of ∗-objects inherits a monoidal structure (cf. [Jac12, Proposition 1] and Proposition 3.15 in the present paper). This has interesting consequences for the theory of involutive monads in [Jac12]
and the developments in our present paper.
Definition 3.4. An involutive (symmetric) monoidal category is a triple (C, J, j) consisting of a
(symmetric) monoidal category C, a (symmetric) monoidal endofunctor J = (J, J2 , J0 ) : C → C
and a (symmetric) monoidal natural isomorphism j : IdC → J 2 satisfying jJ = Jj : J −→ J 3
The following statement is proven in [Jac12, Lemma 7].
10

.

(3.9)

Lemma 3.5. For any involutive (symmetric) monoidal category, the (symmetric) monoidal endofunctor J = (J, J2 , J0 ) : C → C is strong, i.e. J2 : ⊗ (J × J) → J ⊗ and J0 : I → JI are isomorphisms.
Remark 3.6. Let us emphasize again and more clearly that our Definition 3.4 of involutive
(symmetric) monoidal categories agrees with the one of Jacobs [Jac12]. The definitions in [BM09]
and [Egg11] are different because their analog of J2 is order-reversing, i.e. a natural isomorphism
⊗op (J × J) → J ⊗. The reason why we consider order-preserving J2 as in [Jac12] is that this is better suited for our development of involutive operad theory, cf. Remark 4.6 below.
△
Remark 3.7. The condition for j : IdC → J 2 to be a (symmetric) monoidal natural transformation explicitly means that the diagrams
⊗

⊗ (j×j)

/ ⊗(J 2 × J 2 )
J2 (J×J)



J⊗(J × J)

id⊗

I

⊗

j⊗

JI
JJ0

idI

J J2





(3.10)

I
✄ ❄❄❄
❄❄J0
❄❄
✄
✄
❄
✄✄

idI ✄✄✄





/ J 2⊗

I

/ J 2I

jI

commute. One may reinterpret these diagrams as follows: The left diagram states that (⊗, J2 ) :
(C, J, j) × (C, J, j) → (C, J, j) is an involutive functor on the product involutive category
(C, J, j) × (C, J, j) = (C × C, J × J, j × j), see also Remark 2.22. The right diagram states that (J0 : I → JI) ∈ ∗-Obj(C, J, j) is a ∗-object in (C, J, j). These two structures allow us to endow the functor I ⊗ (−) : C → C with an involutive structure I ⊗ J(−) → J(I ⊗ (−)) defined by the components
I ⊗ Jc

J0 ⊗id

/ JI ⊗ Jc

J2 I,c

/ J(I ⊗ c)

,

(3.11)

for all c ∈ C. An analogous statement holds true for the functor (−) ⊗ I : C → C. The axioms for the (symmetric) monoidal structure on J can then be reinterpreted as the equivalent property that the associator and unitors (as well as the braiding in the symmetric case) are involutive natural transformations.
Summing up, we obtain an equivalent description of an involutive (symmetric) monoidal category in terms of the following data: An involutive category (C, J, j), an involutive functor
(⊗, J2 ) : (C, J, j) × (C, J, j) → (C, J, j), a ∗-object (J0 : I → JI) ∈ ∗-Obj(C, J, j) and involutive natural transformations for the associator and unitors (as well as the braiding in the symmetric case), which satisfy analogous axioms as those for (symmetric) monoidal categories. This alternative point of view is useful for (3.16) and (3.17) below.
△
Example 3.8. For any (symmetric) monoidal category C, the triple (C, IdC , idIdC ), with IdC
the identity (symmetric) monoidal functor and idIdC the identity (symmetric) monoidal natural transformation, defines an involutive (symmetric) monoidal category. We call this the trivial involutive (symmetric) monoidal category over C.
▽
Example 3.9. Let us equip the category of complex vector spaces VecC with its standard symmetric monoidal structure where ⊗ is the usual tensor product, I = C is the ground field and τ is given by the flip maps τV,W : V ⊗ W → W ⊗ V , v ⊗ w 7→ w ⊗ v. The endofunctor
(−) : VecC → VecC from Example 2.3 can be promoted to a symmetric monoidal functor by using the canonical maps (−)2V,W : V ⊗ W → V ⊗ W and complex conjugation (−)0 : C → C.
The resulting triple (VecC , (−), idIdVecC ) is an involutive symmetric monoidal category.
11

▽

Example 3.10. Recall the groupoid of C-profiles ΣC from Example 2.4. The category ΣC may be equipped with the symmetric monoidal structure given by concatenation of C-profiles, i.e.
c ⊗ d = (c1 , . . . , cn , d1 , . . . , dm ), I = ∅ is the empty C-profile and τc,d := τ h|c|, |d|i : c ⊗ d → d ⊗ c is the block transposition. The reversal endofunctor Rev : ΣC → ΣC can be promoted to a symmetric monoidal functor by using
Rev2c,d := τ h|c|, |d|i : Rev(c) ⊗ Rev(d) −→ Rev(c ⊗ d)

(3.12)

and Rev0 := id∅ : ∅ → Rev(∅) = ∅. The resulting triple (ΣC , Rev, idIdΣC ) is an involutive symmetric monoidal category.
▽
Definition 3.11. An involutive (symmetric) monoidal functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ )
consists of a (symmetric) monoidal functor F = (F, F2 , F0 ) : C → C′ and a (symmetric) monoidal natural transformation ν : F J → J ′ F satisfying the analog of diagram (2.3) in Definition 2.6.
An involutive (symmetric) monoidal natural transformation ζ : (F, ν) → (G, χ) between involutive (symmetric) monoidal functors (F, ν), (G, χ) : (C, J, j) → (C′ , J ′ , j ′ ) is a natural transformation ζ : F → G that is both involutive and (symmetric) monoidal.
Proposition 3.12. Involutive (symmetric) monoidal categories, involutive (symmetric) monoidal functors and involutive (symmetric) monoidal natural transformations form a 2-category I(S)MCat.
Remark 3.13. The condition for the natural transformation ν : F J → J ′ F to be monoidal explicitly means that the diagrams
⊗′ (F J × F J)

⊗′ (ν×ν)

/ ⊗′ (J ′ F × J ′ F )
J2′ (F ×F )

F2 (J×J)





J ′ ⊗′ (F × F )

F ⊗ (J × J)

FI

J ′ F2

F J2
ν⊗

I′ ❈

(3.13)

❈❈ J ′
❈❈ 0
❈❈
❈!

J ′I ′
J ′ F0

F J0





F J⊗

⑤
F0 ⑤⑤⑤
⑤
⑤⑤
~⑤⑤





/ J ′F ⊗

F JI

νI

/ J ′F I

commute. From the perspective established in Remark 3.7, one may reinterpret these diagrams as follows: The left diagram states that F2 is an involutive natural transformation

F2 : (⊗′ , J2′ ) (F, ν) × (F, ν) −→ (F, ν) (⊗, J2 )
(3.14)

of involutive functors from (C, J, j) × (C, J, j) to (C′ , J ′ , j ′ ). The right diagram states that F0
defines a morphism


F0 : J0′ : I ′ → J ′ I ′ −→ ∗-Obj(F, ν) J0 : I → JI
(3.15)

in the category ∗-Obj(C′ , J ′ , j ′ ) of ∗-objects in (C′ , J ′ , j ′ ).

Summing up, we obtain an equivalent description of an involutive (symmetric) monoidal functor in terms of the following data: An involutive functor (F, ν) : (C, J, j) → (C′ , J ′ , j), an involutive natural transformation F2 as in (3.14) and a ∗-morphism F0 as in (3.15), which satisfy axioms analogous to those for a (symmetric) monoidal functor. This alternative point of view is useful for (3.20) below.
△
Remark 3.14. Let us summarize Remarks 3.7 and 3.13 by one slogan: Involutive (symmetric)
monoidal categories are the same things as (symmetric) monoidal involutive categories.
△
Let (C, J, j) be an involutive (symmetric) monoidal category and consider its category of
∗-objects ∗-Obj(C, J, j). Making use of the 2-functor ∗-Obj : ICat → Cat given in (2.13), we
12

may equip the category ∗-Obj(C, J, j) with a (symmetric) monoidal structure. Concretely, the tensor product functor is given by
⊗

/ ∗-Obj(C, J, j)
3
❣
❣
❣
❣
❣❣
❣
❣
❣
❣
❣❣❣❣❣∗-Obj(⊗,J2 )
❣❣❣❣❣

∗-Obj(C, J, j) × ∗-Obj(C, J, j)
∼
=



(3.16)

∗-Obj (C, J, j) × (C, J, j)

where the vertical isomorphism was explained in Remark 2.22 and the involutive functor (⊗, J2 )
in Remark 3.7. The unit object

J0 : I → JI ∈ ∗-Obj(C, J, j)
(3.17)
is the ∗-object constructed in Remark 3.7. The associator and unitors (as well as the braiding in the symmetric case) are obtained by applying the 2-functor ∗-Obj to the associator and unitors
(as well as the braiding in the symmetric case) of (C, J, j), which makes sense because Remark
3.7 shows that these are involutive natural transformations. Let us also mention that the tensor product of two ∗-objects (∗ : c → Jc), (∗′ : c′ → Jc′ ) ∈ ∗-Obj(C, J, j) explicitly reads as
(∗ : c → Jc) ⊗ (∗′ : c′ → Jc′ ) =



c ⊗ c′

∗⊗∗′ /

Jc ⊗ Jc′

J2 c,c′

/ J(c ⊗ c′ )



.

(3.18)

Summing up, we have proven
Proposition 3.15. Let (C, J, j) be an involutive (symmetric) monoidal category. Then the category of ∗-objects ∗-Obj(C, J, j) is a (symmetric) monoidal category with tensor product (3.16)
and unit object (3.17). Moreover, if (C, J, j) is also closed, i.e. it has internal homs, then
∗-Obj(C, J, j) is closed too (cf. [Jac12, Proposition 1]).
The assignment of the (symmetric) monoidal categories of ∗-objects extends to a 2-functor
∗-Obj : I(S)MCat −→ (S)MCat

,

(3.19)

which we shall denote with an abuse of notation by the same symbol as the 2-functor in (2.13).
Concretely, this 2-functor is given by the following assignment:
• an involutive (symmetric) monoidal category (C, J, j) is mapped to the (symmetric) monoidal category ∗-Obj(C, J, j) given in Proposition 3.15;
• an involutive (symmetric) monoidal functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ) is mapped to the
(symmetric) monoidal functor
∗-Obj(F, ν) : ∗-Obj(C, J, j) −→ ∗-Obj(C′ , J ′ , j ′ )

(3.20a)

with underlying functor as in (2.13) and (symmetric) monoidal structure given by
∗-Obj(F )2 := ∗-Obj(F2 )

,

∗-Obj(F )0 := F0

,

(3.20b)

where F2 and F0 should be interpreted according to Remark 3.13;
• an involutive (symmetric) monoidal natural transformation ζ : (F, ν) → (G, χ) is mapped to the (symmetric) monoidal natural transformation determined by (2.13).
Remark 3.16. Notice that the 2-functor ∗-Obj : I(S)MCat → (S)MCat given in (3.19) is a lift of the 2-functor ∗-Obj : ICat → Cat given in (2.13) along the forgetful 2-functors forget⊗ :
I(S)MCat → ICat and forget⊗ : (S)MCat → Cat that forget the (symmetric) monoidal
13

structures. More precisely, using the explicit descriptions of our 2-functors, one easily confirms that the diagram
I(S)MCat

∗-Obj

/ (S)MCat

forget⊗

(3.21)

forget⊗



ICat



/ Cat

∗-Obj

of 2-categories and 2-functors commutes (on the nose).

△

We conclude this section with a useful result that generalizes Theorem 2.23 to the (symmetric)
monoidal setting. Let us first notice that the forgetful functor U : ∗-Obj(C, J, j) → C satisfies
⊗(U × U ) = U ⊗ and U (J0 : I → JI) = I, hence it can be promoted to a (symmetric) monoidal functor via the trivial (symmetric) monoidal structure U2 = idU ⊗ and U0 = idI .
Theorem 3.17. Let (C, J, j) be an involutive (symmetric) monoidal category. Any (symmetric)
monoidal section ∗ : C → ∗-Obj(C, J, j) of the forgetful (symmetric) monoidal functor U :
∗-Obj(C, J, j) → C canonically determines an I(S)MCat-isomorphism between (C, J, j) and the trivial involutive (symmetric) monoidal category (C, IdC , idIdC ). In particular, if such a section of U exists, then the involutive (symmetric) monoidal categories (C, J, j) and (C, IdC , idIdC ) are isomorphic.
Proof. Using that the (symmetric) monoidal structure on U is trivial, i.e. U2 = idU ⊗ and U0 = idI , and also that U is a faithful functor, one observes that the (symmetric) monoidal structure on the (symmetric) monoidal section ∗ : C → ∗-Obj(C, J, j) is necessarily trivial. The proof then proceeds analogously to the one of Theorem 2.23.
Corollary 3.18. The involutive symmetric monoidal category (ΣC , Rev, idIdΣC ) of C-profiles equipped with reversal as involutive structure (cf. Example 3.10) is isomorphic to the trivial involutive symmetric monoidal category (ΣC , IdΣC , idIdΣC ).
Proof. By Theorem 3.17, it is sufficient to construct a symmetric monoidal section ρ = (ρ, ρ2 , ρ0 ) :
ΣC → ∗-Obj(ΣC , Rev, idIdΣC ) of the forgetful symmetric monoidal functor U . Taking the underlying functor as in Example 2.18, i.e. ρ : c 7→ (ρ|c| : c → cρ|c| ) with the order-reversal permutations
ρ|c| ∈ Σ|c| , one easily checks that ⊗(ρ × ρ) = ρ⊗ and ρ(∅) = (id∅ : ∅ → ∅) = (Rev0 : ∅ → Rev(∅)).
We choose the trivial symmetric monoidal structure ρ2 = idρ⊗ and ρ0 = id∅ .

3.3

∗-monoids

Let us recall the 2-functors Mon : (S)MCat → Cat given in (3.7), ∗-Obj : ICat → Cat given in (2.13) and its lift ∗-Obj : I(S)MCat → S(M)Cat given in (3.19). The aim of this subsection is to describe a 2-functor Mon : I(S)MCat → ICat that lifts Mon : (S)MCat → Cat to the involutive setting, such that the diagram
I(S)MCat

∗-Obj

✤

/ (S)MCat

✤

(3.22)

Mon

Mon ✤





ICat

∗-Obj

/ Cat

of 2-categories and 2-functors commutes (on the nose). We then define ∗-monoids in terms of the diagonal 2-functor ∗-Mon : I(S)MCat → Cat in this square.
Let us start with describing the 2-functor
Mon : I(S)MCat −→ ICat that lifts (3.7) to the involutive setting in some detail:
14

(3.23)

• an involutive (symmetric) monoidal category (C, J, j) is mapped to the involutive category

Mon(C, J, j) := Mon(C), Mon(J), Mon(j) ∈ ICat
(3.24)

given by evaluating the 2-functor (3.7) on the (symmetric) monoidal category C, on the
(symmetric) monoidal endofunctor J : C → C and on the (symmetric) monoidal natural isomorphism j : IdC → J 2 ;

• an involutive (symmetric) monoidal functor (F, ν) : (C, J, j) → (C′ , J ′ , j ′ ) is mapped to the involutive functor

Mon(F, ν) := Mon(F ), Mon(ν) : Mon(C, J, j) −→ Mon(C′ , J ′ , j ′ )
(3.25)
given by evaluating the 2-functor (3.7) on the (symmetric) monoidal functor F : C → C′
and on the (symmetric) monoidal natural transformation ν : F J → J ′ F ;

• an involutive (symmetric) monoidal natural transformation ζ : (F, ν) → (G, χ) is mapped to the involutive natural transformation
Mon(ζ) : Mon(F, ν) −→ Mon(G, χ)

(3.26)

given by evaluating the 2-functor (3.7) on ζ.
Lemma 3.19. The diagram (3.22) of 2-categories and 2-functors commutes (on the nose).
Proof. This is an elementary check using the explicit definitions of the 2-functors given in (3.7),
(2.13), (3.19) and (3.23).
Definition 3.20. The 2-functor ∗-Mon : I(S)MCat → Cat is defined as the diagonal 2-functor in the commutative square (3.22), i.e.
∗-Obj
/ (S)MCat
❚❚❚❚
❚❚❚❚∗-Mon
❚❚❚❚
Mon
❚❚❚❚ 
)/

(3.27)

I(S)MCat
Mon



ICat

∗-Obj

Cat

For an involutive (symmetric) monoidal category (C, J, j), we call ∗-Mon(C, J, j) the category of ∗-monoids in (C, J, j).
Remark 3.21. Let (C, J, j) be an involutive (symmetric) monoidal category. We provide an explicit description of the objects and morphisms in the associated category of ∗-monoids
∗-Mon(C, J, j), which we shall call ∗-monoids and ∗-monoid morphisms. Unpacking Definition
3.20, one obtains that a ∗-monoid is a quadruple (M, µ, η, ∗) ∈ ∗-Mon(C, J, j) consisting of an object M ∈ C and three C-morphisms µ : M ⊗ M → M , η : I → M and ∗ : M → JM , which satisfy the following conditions:
(1) (M, µ, η) is a monoid in the (symmetric) monoidal category C;
(2) ∗ : M → JM is a ∗-object in the involutive category (C, J, j);
(3) these two structures are compatible in the sense that the diagrams
I

η

/M
∗

J0



JI



Jη

/ JM

M ⊗M

∗⊗∗

/ JM ⊗ JM

J2M,M

/ J(M ⊗ M )

µ

Jµ





M

∗

in C commute.
15

/ JM

(3.28)

As a consequence of Lemma 3.19, these conditions have two equivalent interpretations which correspond to the counterclockwise and clockwise paths in the commutative diagram (3.27): The first option is to regard ∗ : (M, µ, η) → Mon(J)(M, µ, η) as a ∗-object in the involutive category
Mon(C, J, j) ∈ ICat. The second option is to regard η : (J0 : I → JI) → (∗ : M → JM ) and
µ : (∗ : M → JM ) ⊗ (∗ : M → JM ) → (∗ : M → JM ) as the structure maps of a monoid in the
(symmetric) monoidal category ∗-Obj(C, J, j) ∈ (S)MCat.
A ∗-monoid morphism f : (M, µ, η, ∗) → (M ′ , µ′ , η ′ , ∗′ ) is a C-morphism f : M → M ′ that preserves both the monoid structures and ∗-involutions.
△
Example 3.22. Let us consider a ∗-monoid (A, µ, η, ∗) in the involutive symmetric monoidal category (VecC , (−), idIdVecC ) from Example 3.9. In particular, the triple (A, µ, η) is an associative and unital algebra over C with multiplication a b = µ(a ⊗ b) and unit 1 = η(1). By Example
2.17, ∗ is a complex anti-linear automorphism of A that squares to the identity, i.e. a∗∗ = a.
The compatibility conditions in (3.28) state that 1∗ = 1 and (a b)∗ = a∗ b∗ . We would like to emphasize that the latter condition is not the usual axiom for associative and unital ∗-algebras over C, which is given by order-reversal (a b)∗ = b∗ a∗ . As a consequence, our concept of ∗monoids given in Definition 3.20 does not include the usual associative and unital ∗-algebras over
C as examples. We will show later in Example 7.9 that the usual associative and unital ∗-algebras over C are recovered as ∗-algebras over a suitable ∗-operad, which provides a sufficiently flexible framework to implement order-reversal (a b)∗ = b∗ a∗ .
▽

4

Involutive structures on colored symmetric sequences

Colored operads can be defined as monoids in the monoidal category of colored symmetric sequences, see e.g. [Yau16, WY18, Yau18, GJ17] and below for a brief review. Let C ∈ Set be any non-empty set and M any cocomplete closed symmetric monoidal category. (We denote the monoidal structure on M by ⊗ and I, and the internal hom by [−, −] : Mop × M → M.) The category of C-colored symmetric sequences with values in M is defined as the functor category
SymSeqC (M) := MΣC ×C

,

(4.1)

where ΣC is the groupoid of C-profiles defined in Example 2.4 and the set C is regarded as a discrete category. Given X ∈ SymSeqC (M), we write

(4.2a)
X ct ∈ M
for the evaluation of this functor on objects (c, t) ∈ ΣC × C and


t
X(σ) : X ct −→ X cσ

(4.2b)

for its evaluation on morphisms σ : (c, t) → (cσ, t) in ΣC × C.

The category SymSeqC (M) can be equipped with the following monoidal structure: The tensor product is given by the circle product ◦ : SymSeqC (M)×SymSeqC (M) → SymSeqC (M).
Concretely, the circle product of X, Y ∈ SymSeqC (M) is defined by the coend
Z a Z (b ,...,b )
1
m



am 
a 
t
ΣC b1 ⊗ · · · ⊗ bm , c ⊗ X at ⊗ Y b11 ⊗ · · · ⊗ Y bm
(X ◦ Y ) c :=
,
(4.3)
for all (c, t) ∈ ΣC × C. Two remarks are in order: (1) This expression makes use of the symmetric monoidal structure on ΣC that we
 described in Example 3.10.t  (2) The tensor product between the Hom-set ΣC b1 ⊗ · · · ⊗ bm , c ∈
` Set and the object X a ∈ M is given by the canonical
Set-tensoring of M, i.e. S ⊗ m := s∈S m for any S ∈ Set and m ∈ M. The circle unit is the object I◦ ∈ SymSeqC (M) defined by

(4.4)
I◦ ct := ΣC (t, c) ⊗ I , for all (c, t) ∈ ΣC × C.

16

Proposition 4.1. (SymSeqC (M), ◦, I◦ ) is a right closed monoidal category.
The aim of this section is to transfer these structures and results to the setting of involutive categories.

4.1

Product-exponential 2-adjunction

Because the category of symmetric sequences (4.1) is defined as a functor category, we shall start with developing a notion of functor categories in the involutive setting. For this we will first recall the relevant structures for ordinary category theory from a perspective that easily generalizes to involutive category theory.
e Cat the 2-category with objects given by pairs (C, D) of categories,
Let us denote by Cat ×
morphisms given by pairs (F, G) of functors and 2-morphisms given by pairs (ζ, ξ) of natural e to denote transformations, and all compositions given component-wise. (We use the symbol ×
the above product 2-category because we reserve the symbol × for the 2-functors defined below.)
Notice that taking products of categories, functors and natural transformations defines a 2-functor e Cat −→ Cat .
× : Cat ×

(4.5)

e Cat −→ Cat
(−)(−) : Catop ×

(4.6)

Let us denote by Catop the opposite 2-category, i.e. morphisms C → D are functors F : D → C
going in the opposite direction and 2-morphisms are not reversed. We define the exponential
2-functor

as follows:

• a pair (D, C) of categories is mapped to the functor category CD ;
′

• a pair (G : D′ → D, F : C → C′ ) of functors is mapped to the functor F G : CD → C′ D
that acts on objects and morphisms as

F G X : D → C := (F XG : D′ → C′ ) ,
(4.7a)

G
F α : X → Y := (F αG : F XG → F Y G) ;
(4.7b)
• a pair (ξ : G → G′ , ζ : F → F ′ ) of natural transformations is mapped to the natural
′
transformation ζ ξ : F G → F ′ G with components given by any of the two compositions in the commutative square
F XG❘
F Xξ



F XG′

ζXG

❘ ❘ (ζ ξ )
❘ ❘X
❘
ζXG′

/ F ′ XG

(4.8)

F ′ Xξ

❘(

/ F ′ XG′

for all X ∈ CD .
The two 2-functors × and (−)(−) are related by a family of 2-adjunctions.
Proposition 4.2. For every D ∈ Cat, there is a 2-adjunction
/

(−) × D : Cat o

17

Cat : (−)D

.

(4.9)

Proof. The component at C ∈ Cat of the unit 2-natural transformation η : IdCat → ((−) × D)D
is given by the functor
ηC : C −→ (C × D)D

(4.10)

that assigns to c ∈ C the inclusion functor ηC (c) : D → C × D specified by d 7→ (c, d). The component at C ∈ Cat of the counit 2-natural transformations ǫ : (−)D × D → IdCat is given by the evaluation functor
ǫC : CD × D −→ C ,

(4.11)

that assigns to (X, d) ∈ CD × D the object Xd ∈ C. The triangle identities are a straightforward check.
Because of their 2-functoriality, our constructions above can be immediately extended to involutive category theory. Concretely, using the 2-functor (4.5), we define the product 2-functor e ICat −→ ICat
× : ICat ×

(4.12)

in the involutive setting as follows:

• a pair of involutive categories is mapped to the involutive category
(C, J, j) × (D, K, k) := (C × D, J × K, j × k)

;

(4.13)

• a pair of involutive functors is mapped to the involutive functor
(F, ν) × (G, χ) := (F × G, ν × χ)

;

(4.14)

• a pair of involutive natural transformations is mapped to the involutive natural transformation ζ × ξ.
Similarly, using the 2-functor (4.6), we define the exponential 2-functor e ICat −→ ICat
(−)(−) : ICatop ×

(4.15)

in the involutive setting as follows:

• a pair of involutive categories is mapped to the involutive category

(C, J, j)(D,K,k) := CD , J K , j k
;

(4.16)

• a pair of involutive functors is mapped to the involutive functor
−1 
(F, ν)(G,χ) := F G , ν χ
;

(4.17)

• a pair of involutive natural transformations is mapped to the involutive natural transformation ζ ξ .
Analogously to Proposition 4.2, one can prove
Proposition 4.3. For every (D, K, k) ∈ ICat

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
