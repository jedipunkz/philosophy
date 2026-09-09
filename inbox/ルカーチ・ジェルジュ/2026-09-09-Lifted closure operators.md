---
source: "https://arxiv.org/abs/math/0502410v2"
title: "Lifted closure operators"
author: "Gábor Lukács"
year: "2005"
publication: "arXiv preprint / math.CT"
download: "https://arxiv.org/pdf/math/0502410v2"
pdf: "https://arxiv.org/pdf/math/0502410v2"
captured_at: "2026-09-09T04:08:09Z"
updated_at: "2026-09-09T04:08:09Z"
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

# Lifted closure operators

- 著者: Gábor Lukács
- 年: 2005
- 掲載情報: arXiv preprint / math.CT
- 情報源: [arxiv](https://arxiv.org/abs/math/0502410v2)
- ダウンロード: https://arxiv.org/pdf/math/0502410v2
- PDF: https://arxiv.org/pdf/math/0502410v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

In this paper, we study the properties of closure operators obtained as initial lifts along a reflector, and compactness with respect to them in particular. Applications in the areas of topology, topological groups and topological *-algebras are presented.

## PDF Text

arXiv:math/0502410v2 [math.CT] 28 Feb 2005

Lifted closure operators ∗
Gábor Lukács †
November 21, 2018
Abstract
In this paper, we study the properties of closure operators obtained as initial lifts along a reflector, and compactness with respect to them in particular. Applications in the areas of topology, topological groups and topological ∗-algebras are presented.
1.

INTRODUCTION

Throughout this paper, X and A are finitely-complete categories with proper (E, M) and (F , N )
factorization systems, respectively (cf. [17, 7.2]). For X ∈ X, the class M/X of subobjects of X
is preordered by the relation m ≤ n ⇐⇒ (∃j) m = nj, and denoted by sub X.
A closure operator with respect to M is a family c = (cX )X∈X of maps cX : sub X → sub X
such that m ≤ cX (m), cX (m) ≤ cX (n) if m ≤ n, and f (cX (m)) ≤ cY (f (m)) for all f : X → Y
in X and m, n ∈ sub X (cf. [5]). The term “closure operator” has been used for almost a century in various meanings in the context of categories of topological spaces and lattices. However, this
(general) categorical notion of closure operators was invented by Dikranjan and Giuli [6], and was further developed by Dikranjan, Giuli and Tholen in [8]. Following [7], an object X ∈ X is said to be c-compact if for every Y ∈ X, the projection πY : X × Y → Y is c-preserving, or in other words, if πY (cX (m)) = cY (πY (m)) for every m ∈ sub(X × Y ).
If F ⊣ U : A −→ X is an adjunction with unit η, and c is a closure operator of A with respect to N , then one defines the F -initial lift of c as
−1
cηX (m) := ηX
(UcF X (F m))

(1)

for every X ∈ X, m ∈ sub(X) (cf. [5, 5.13]). The morphism F m need not belong to N , so in order to ensure that (1) is defined, one extends the notion of closure as cF X (F m) = cF X ((F m)(1F M ))
(cf. [5, 5.7]).
In this paper, we study the special case of this construction where A is a reflective subcategory of X with reflector F = R and reflection η = ρ. We provide a characterization of cρ -compact objects in X under certain conditions (Theorem 2.7), and applications in the areas of topology, topological groups and topological ∗-algebras are presented We also investigate the categorical properties of c that are inherited by cρ .
∗

2000 Mathematics Subject Classification: 18A32 54B30 (18B30 22A05 46K10)
I gratefully acknowledge the generous financial support received from the Alexander von Humboldt Foundation, the Killam Trusts, and Dalhousie University that enabled me to do this research.
†

2

G. Lukács / Lifted closure operators

The paper is structured as follows: We start off with an investigation of the properties of lifted closure operators in section 2, with an emphasis on compactness with respect to them. In section 3, we show that the notion of w-compactness (introduced by T. Ishii in [23]) is equivalent to being compact with respect to the initial lift of the Kuratowski closure operator on Tych (Tychonoff spaces and their continuous maps) along the Tychonoff functor τ : Top −→ Tych. In section 4, the Bohr-closure cb for topological groups is presented as the initial lift of the Kuratowski closure along the Bohr-compactification, and Theorem 2.7 is used to characterize cb -compactness. The
Bohr-closure also turns out to be a convenient example to show that certain properties are not preserved by the initial lift along a reflection. Finally, in section 5, we prove that the ∗-representation topology on topological ∗-algebras is an initial lift of the Kuratowski closure along a suitable reflection.
2.

PROPERTIES OF LIFTED CLOSURE OPERATORS

In this section, we present results of a positive nature, in other words, properties that carry over from c to cρ . Counterexamples, which show that certain properties are not inherited by cρ , appear later on.
As stated in the Introduction, A is a reflective subcategory of X with reflector R and reflection
ρ, and thus (1) becomes cρX (m) := ρ−1
(2)
X (cRX (Rm)).
It can be shown that N ⊆ M and RE ⊆ F (cf. [5, 5.13]), so in order to avoid confusion, we put subA A for the class of subobjects of A ∈ A in A (i.e., N /A) when necessary.
It was observed by Baron [1] that every reflection decomposes into two epireflections. Using the same idea, we factorize R as an E-reflector followed by an epireflector: For every X ∈ X,
E
E
M
E
one can factorize the reflection as ρX = ρM
X ρX , where ρX ∈ E and ρX ∈ M, and put R X
for the codomain of ρEX . Obviously, RE X belongs to B = {S ∈ X | ρS ∈ M}, and B is the
E-reflective hull of A in X (naturally, with reflector RE and reflection ρE ). Furthermore, ρM is an epimorphism in B, so A is an epireflective subcategory of B with reflector RM = R|| , and
B

R = RM RE . Since B is E-reflective, it inherits from X a proper (EB, MB)-factorization system, where EB = E ∩ mor B and MB = M ∩ mor B.
Remark 2.1. One says that m ∈ sub X is c-dense if cX (m) = 1X . Every S ∈ B is cρ -dense in its reflection RS, because cρRS (ρS ) = ρ−1
RS (cRS (RρS )) = cRS (1RS ) = 1RS . Thus, for every X ∈ X,
ρ
ρX has a c -dense image in RX.
Proposition 2.2. Let c be a closure operator on A. For every X ∈VX and m ∈ sub X, the class
−1
C := {ρ−1
C = cρX (m).
X (cRX (n)) | m ≤ ρX (n), n ∈ subA} has an infimum and
P ROOF. Clearly, cρX (m) ∈ C, because ρX (m) ≤ (Rm)(1RM ). On the other hand, if m ≤ ρ−1
X (n)
for n ∈ subA, then ρX (m) ≤ n, and thus (Rm)(1RM ) ≤ n; in particular, cRX (Rm) ≤ cRX (n).
Therefore, cρX (m) is a lower bound to C.
A closure operator c is idempotent if cA (cA (m)) = cA (m) for every A ∈ A and m ∈ subA A.

G. Lukács / Lifted closure operators

3

Proposition 2.3. Let c be a closure operator of A. If c is idempotent, then so is cρ .
P ROOF. Let X ∈ X and m ∈ sub X. Clearly, (Rρ−1
X (k))(1R(K×RX X) ) ≤ k for every
−1
k ∈ subA RX, so for k = cRX (Rm), the image of RρX (cRX (Rm)) is contained in cRX (Rm).
Since c is idempotent, we obtain:
−1
cρX (cρX (m)) = ρ−1
X (cRX (RρX (cRX (Rm))))
ρ
−1
≤ ρ−1
X (cRX (cRX (Rm))) = ρX (cRX (Rm))) = cX (m).

(3)
(4)

Therefore, cρX (m) = cρX (cρX (m)), because cρX is extensive.
One says that a closure operator c is (finitely) productive
Q if forQevery (finite) family {Ai }i∈I of
Q
cAi (mi ).
objects in A and their subobjects mi ∈ subA Ai , c Ai ( mi ) =
i∈I

i∈I

Proposition 2.4. Suppose that F is (finite) productive, and let c be a closure operator of A. If c is
(finitely) productive and R preserves (finite) products, then cρ is (finitely) productive.
P ROOF. Let {XQ
family of objects in X, and let mi ∈ sub Xi be their i }i∈I be a (finite)Q
Q subobjects. Put X =
Xi and m =
mi . Since R preserves (finite) products, ρX =
ρXi and i∈I
i∈I
i∈I
Q
Q
Rm =
Rmi , and one has (Rm)(1RM ) = (Rmi )(1RMi ), because F is (finitely) productive.
i∈I

i∈I

Therefore, we obtain

Q
Q
Q
def cρX (m) = ρ−1
Rmi )) = ( ρXi )−1 ( cRXi (Rmi ))
X (cRX (
i∈I
i∈I
i∈I
Q −1
Q ρ
=
ρXi (cRXi (Rmi )) =
cXi (mi ), i∈I

(5)
(6)

i∈I

because c is (finitely) productive, and limits interchange with each other.
For A ∈ A, a subobject m ∈ subA A is c-closed if cA (m) = m. One says that A is c-separated
(or c-Hausdorff) if the diagonal δA : A → A × A is c-closed in A (cf. [3, 4.1]).
Proposition 2.5. Let c be a a closure operator of A. For X ∈ X, suppose that (a) RX is cseparated and (b) ρX×X = ρX × ρX . Then X is cρ -separated if and only if ρX is a monomorphism.
P ROOF. Using the assumptions, one can easily see that def

(a)

(b)

−1
−1
cρX×X (δX ) = ρ−1
X×X (cR(X×X) (RδX )) = (ρX × ρX ) (cRX×RX (δRX )) = (ρX × ρX ) (δRX ). (7)

This simple computation shows that cρX×X (δX ) is precisely the kernelpair of ρX .
We turn to characterizing the cρ -compact objects, and to that end, until the end of this section,
E is assumed to be pullback stable.
Proposition 2.6. Let c be a closure operator on A. If X ∈ X is cρ -compact, then:
(a) RE X is cρ -compact;

4

G. Lukács / Lifted closure operators

(b) if RX is c-separated, then ρX ∈ E and RX is c-compact.
P ROOF. Since E is pullback stable, the image of any cρ -compact object under a morphism is again cρ -compact (cf. [3, 5.2(3)]), and thus (a) follows. In order to show (b), notice that RE X is cρ -closed in RX (cf. [3, 5.2(2)]), because RX is c-separated. On the other hand, RE X is cρ -dense in RX (Remark 2.1), and therefore RE X = RX.
For X ∈ X, we say that R preserves products with X if ρX×Y = ρX × ρY for every Y ∈ X.
Theorem 2.7. Suppose that F ⊂ E. Let c be a closure operator of A, and let X ∈ X. If R
preserves products with X, ρX ∈ E, and RX is c-compact, then X is cρ -compact.
Remark 2.8. The condition of F ⊂ E implies that F ⊆ RE, and therefore F = RE (cf. [5, 5.13]).
Its role in the theorem is to ensure that the factorization of a morphism in A coincides with its factorization in X.
In order to prove Theorem 2.7, a technical lemma is required. It is an easy consequence of the
Beck-Chevalley Property (cf. [26, 3.8]), and thus its straightforward proof is omitted:
Lemma 2.9. Let X1 , X2 , Y1 , Y2 ∈ X, and let e : X1 → X2 and f : Y1 → Y2 be morphisms. Put g = e×f : X1 ×Y1 → X2 ×Y2 . If e ∈ E (and E is pullback stable), then for every n ∈ sub(X2 ×Y2 ),
πY1 (g −1 (n)) = f −1 (πY2 (n)).
P ROOF OF T HEOREM 2.7. Let Y ∈ X. We show that πY : X × Y → Y is cρ -preserving. To that end, let m ∈ sub(X × Y ). First, note that
πRY ((Rm)(1RM )) = R(πY (m))(1RN ),

(8)

where πY (m) : N → Y . Indeed, πY m = πY (m)e, where e ∈ E, so R(πY m) = R(πY (m))Re.
Thus, πRY ((Rm)(1RM )) = R(πY m)(1RM ) = R(πY (m))(Re(1RM )) = R(πY (m))(1RN ), as
Re : RM → RN belongs to RE = F and R preserves products with X (so RπY = πRY ).
−1
By Lemma 2.9, πY (ρ−1
X×Y (n)) = ρY (πRY (n)) for every n ∈ sub(RX × RY ), because ρX ∈ E
and ρX×Y = ρX × ρY . Thus, putting n = cRX×RY (Rm) yields: def

−1
πY (cρX×Y (m)) = πY (ρ−1
X×Y (cRX×RY (Rm))) = ρY (πRY (cRX×RY (Rm))).

Since RX is c-compact, one has cRY (πRY (k)) = πRY (cRX×RY (k)) for every k ∈ sub(RX ×RY ).
By letting k = (Rm)(1RM ) and then applying ρ−1
Y to both sides, one obtains
−1
ρ−1
Y (πRY (cRX×RY (Rm))) = ρY (cRY (πRY ((Rm)(1RM ))))
(8)

def

ρ
= ρ−1
Y (cRY (R(πY (m))(1RN ))) = cY (πY (m)).

Therefore, cρY (πY (m)) = πY (cρX×Y (m)), as desired.
Corollary 2.10. Suppose that N = M ∩ mor A and that A is E-reflective in X. Let c be a closure operator of A, and let X ∈ X. If R preserves products with X, then X is cρ -compact if and only if
RX is c-compact.

G. Lukács / Lifted closure operators

5

P ROOF. Since A is E-reflective in X, the (E, M)-factorization system in X restricts to a proper
(EA, MA)-factorization system in A, where EA = E ∩ mor A and MA = M ∩ mor A. Because of the “orthogonality” relation, N and F determine each other, and thus N = MA implies F = EA.
So, the conditions of Theorem 2.7 are fulfilled, and therefore if RX is c-compact, then X is cρ compact. On the other hand, by Proposition 2.6(a), if X is cρ -compact, then so is RE X = RX; hence, RX is c-compact, as desired.
A combination of Theorem 2.7 and Proposition 2.6 yields:
Corollary 2.11. Let c be a closure operator of A. Suppose that F ⊂ E and every object in A is c-separated. If X ∈ X is so that R preserves products with X, then X is cρ -compact if and only if
ρX ∈ E and RX is c-compact.
The discrete closure operator s is defined as sA (m) = m for every m ∈ subA A and A ∈ A.

Corollary 2.12. Suppose that F ⊂ E, and let X ∈ X. If R preserves products with X, then X is sρ -compact if and only if ρX ∈ E.
3.

APPLICATION I: THE TYCHONOFF FUNCTOR

Put I for the closed unit interval. For any X ∈ Top (the category of topological spaces and their continuous maps), the evaluation map ΦX : X → IC(X,I) , mapping x ∈ X to (f (x))f ∈C(X,I) , is continuous. One sets τ X = ΦX (X), and τ is called the Tychonoff functor : It is the reflector of
Tych (the full subcategory formed by the Tychonoff spaces) in Top.
Put X = Top, A = Tych, R = τ , and ρ = θ. The category Top admits a proper (Onto, Embed)factorization system, and the reflection θX : X → τ X is surjective (by definition) for every
X ∈ Top. Thus, by putting N = MA, we arrive at the setting described in Corollary 2.10, because E = Onto is obviously pullback stable. Therefore, for any closure operator c of Tych, if
τ preserves products with X ∈ Top, then X is cθ -compact if and only if τ X is c-compact. Here, we are interested in the case where c = K, the Kuratowski-closure. Since the K-compact spaces in Tych are precisely the compact ones (by the classical Kuratowski-Mrówka Theorem, cf. [13,
3.1.16]), the problem of characterizing K θ -compact spaces in Top boils down to a question concerning preservation of products by τ . Hence, we use the terminology and the results of T. Ishii, who studied very carefully and extensively the Tychonoff functor in [24] and [23], and in [25] gave a survey of the results concerning it.
Let X ∈ Top; F ⊂ X is a zero-set of X if there exists a continuous map f : X → I such that
F = {x ∈ X | f (x) = 0}, and G ⊂ X is a cozero-set of X if there exists a continuous map g : X → I such that G = {x ∈ X | g(x) > 0}. A subset A of X is τ -open if A is a union of cozero-sets of X. (The complement of a τ -open set is called a τ -closed set.) Following T. Ishii, a space X isTw-compact if for any family {Pλ } of τ -open subsets of X with the finite intersection property, P̄λ 6= ∅ holds; or equivalently, if for any collection {Fα } of closed sets of X that is closed
T under finite intersections and such that each Fα contains a non-empty cozero-set of X, one has Fα 6= ∅ (cf. [23, 1.4]). Although τ X is compact for every w-compact space X (cf. [23,
2.1]), there exists a non-w-compact, regular Hausdorff space X such that τ X is compact (cf. [23,
2.3]).

6

G. Lukács / Lifted closure operators

Theorem 3.1. A space X ∈ Top is K θ -compact if and only if it is w-compact.
In order to prove Theorem 3.1, we need two results of T. Ishii. Recall that a continuous map f : X → Y of topological spaces is a Z-map if f (Z) is closed in Y for every zero-set Z of X.
Fact A. ([23, 2.7], [25, 3.11]) For a space X ∈ Top, the following properties are equivalent:
(i) X is w-compact;
(ii) The projection πY : X × Y → Y is a Z -map for every Y ∈ Top.

Fact B. ([23, 1.5], [25, 4.1]) For a space X ∈ Top, the following statements are equivalent:
(i) for every x ∈ X there exists a cozero-set W containing x such that W̄ is w-compact;
(ii) τ (X × Y ) = τ (X) × τ (Y ) for every Y ∈ Top.

P ROOF OF T HEOREM 3.1. Suppose that X is K θ -compact, and let Y ∈ Top. In order to show that πY : X × Y → Y is a Z-map, let Z ⊂ X × Y be a zero-set in X × Y . In particular, Z
is K θ -closed in X × Y , and thus πY (Z) is K θ -closed in Y , because X is K θ -compact and K θ
idempotent (Proposition 2.3). Therefore, πY (Z) is closed in Y , which shows that πY is a Z-map.
This holds for every Y ∈ Top, and hence X is w-compact (Fact A).
Conversely, suppose that X is w-compact. Then τ preserves products with X (Fact B, with
W = X in (i)). Obviously, E = Onto is pullback stable, and thus, by Corollary 2.10, X is K θ compact if and only if τ X is K-compact. The latter is equivalent to τ X being compact, by the classical Kuratowski-Mrówka Theorem (cf. [13, 3.1.16]). Since X is w-compact, τ X is compact
(cf. [23, 2.1]), and therefore X is K θ -compact.
We note that if τ preserves products with X (which can be guaranteed by the condition in
Fact B), then τ X being compact implies that X is K θ -compact. Therefore, Theorem 3.1 has the following corollary.
Corollary 3.2. Let X ∈ Top be such that for every x ∈ X there exists a cozero-set W containing x such that W̄ is w-compact. Then the following statements are equivalent:
(i) X is K θ -compact;
(ii) X is w-compact;
(iii) τ X is compact.
4.

APPLICATION II: THE BOHR-CLOSURE

Similarly to the relationship between Top and HComp (compact Hausdorff spaces), the full subcategory Grp(HComp) of compact Hausdorff groups is reflective in Grp(Top) (topological groups and their continuous homomorphisms). The reflection is called the Bohr-compactification, and is denoted by ρG : G → bG; ρG has a dense image, but it need not be onto. The kernel of ρG is denoted by n(G) and it is called the von Neumann kernel of G. There are many groups that do not admit any non-trivial continuous homomorphism into a compact group at all. Such groups are called minimally almost periodic, and their Bohr-compactification is trivial (cf. [28]). Therefore,

G. Lukács / Lifted closure operators

7

one should not expect ρG to be an injection. Those groups G for which n(G) = {e} are called maximally almost periodic.
A group G is precompact if for every neighborhood U of e ∈ G, there exists a finite subset
F ⊂ G such that G = F U. It is interesting to note that while topological spaces X whose StoneČech reflection X → βX is an embedding are exactly the Tychonoff spaces, groups G such that
ρG is an embedding are characterized by being Hausdorff (and thus Tychonoff) and precompact.
Remark 4.1. It follows from the famous Peter-Weyl Theorem that finite-dimensional irreducible unitary representations of a compact Hausdorff group separate its points (cf. [33, Thm. 33]). Every such representation of an abelian group is one dimensional. Thus, for every compact Hausdorff abelian group K and k ∈ K such that k 6= 0, there is χ ∈ K̂ such that χ(k) 6= 0, where
K̂ = H (K, T) is the group of continuous characters of K (i.e., continuous homomorphisms
χ : K → T, where T = R/Z). If G is an abelian group, then so is its Bohr-compactification bG (because ρG (G) is an abelian dense subgroup of bG). Therefore, g ∈ ker ρG if and only if c By the universal property of bG, there is a one-to-one corresponχ(ρG (g)) = 0 for every χ ∈ bG.
c because T ∈ Grp(HComp) - in other words, Ĝ = bG
c as sets. Hence, if dence between Ĝ and bG,
G is an abelian group, then
\
n(G) =
ker χ.
(9)
χ∈Ĝ

The Bohr-closure cb of a subgroup S ⊆ G is defined as cb (S) = ρ−1
G (ρG (S)). We equip
Grp(Top) with the (Onto, Embed)-factorization system, while Grp(HComp) is provided with the
(Onto, ClEmb)-factorization system (ClEmb stands for the closed embeddings). In this setting, it follows from Proposition 2.2 that cb is precisely the b-initial lift sρ of the discrete closure operator s on Grp(HComp).
Proposition 4.2.
(a) cb is idempotent;
(b) cb is productive;
(c) G ∈ Grp(Top) is cb -separated if and only if ρG is injective (i.e., G is maximally almost periodic);
(d) G ∈ Grp(Top) is cb -compact if and only if ρG is surjective.
P ROOF. (a) follows from Proposition 2.3. (b) follows from Proposition 2.4, because the Bohrcompactification preserves products, and surjective maps in Grp(HComp) are productive. (c) is an immediate consequence of Proposition 2.5, because b is productive and the class of monomorphisms coincides with the injective maps in Grp(Top). Finally, (d) is a special case of Corollary 2.12.
Let G be a locally compact Hausdorff group. One says that G is central if G/Z(G) is compact
(cf. [14], [15]); G is a Moore group if every irreducible unitary representation of G is finitedimensional (cf. [27], [36]). Moore groups are automatically maximally almost periodic, because by the Gelfand-Raı̆kov Theorem ([18, 22.12]), irreducible representations of a locally compact group separate its points.

8

G. Lukács / Lifted closure operators

Theorem 4.3. Let G be a locally compact Hausdorff cb -compact group.
(a) If G is Moore, then it is compact.
(b) If G is central, then it is compact.
(c) If G has nilpotency class ≤ 2, then it is compact.

(d) The quotient G/[[G, G], G] is compact; in particular, n(G) ⊆ [[G, G], G].

(e) Every continuous homomorphism α : G → hR, +i is trivial, and thus G is unimodular.

P ROOF. (a) Hughes [19, Theorem 1] & [20, Theorem 1] showed that for a locally compact Hausdorff group G, a subspace K ⊂ G is compact if and only if K is compact in Gw , the group G
equipped with the pointwise topology induced by its irreducible unitary representations. Remus and Trigos-Arrieta [35, Theorem 1] showed that for Moore groups, Gw = ρG (G). Since G is cb -compact, ρG (G) = bG, and therefore Gw is compact. Hence, by Hughes’ result, G is compact.
(b) follows from (a), because every central group is Moore (cf. [15, Theorem. 2.1]).
(c) The locally compact abelian group Gab := G/[G, G], being the continuous image of G, is cb -compact. It is also Moore, because every irreducible unitary representation of an abelian group is one dimensional. Thus, by (a), Gab is compact. Since G has nilpotency class 2, one has
[[G, G], G] = 1, and so [G, G] ⊆ Z(G). Therefore, G/Z(G) is a quotient of Gab = G/[G, G], and hence G is central. The statement now follows by (b).
(d) follows from (c).
(e) α factors through Gab , which is compact (by (d)), but the only compact subgroup of R is the trivial one. Therefore, α must be trivial.
A group is balanced (or equivalently, admits an invariant basis) if its left and right uniformities coincide. A Hausdorff locally compact connected group is balanced if and only if it is central (cf.
[14, Theorem 4.3]), which yields:
Corollary 4.4. Every Hausdorff locally compact connected balanced cb -compact group is compact.
One might hope that cb -compactness “almost” implies compactness, but this is not the case.
For instance, every minimally almost periodic group has trivial Bohr-compactification, and in particular, they are cb -compact. The examples below show that all the conditions in Corollary 4.4 are indeed necessary.
Example 4.5. The special linear group SL2 (R) is a locally compact, connected, and cb -compact group, which is not compact. Indeed, SL2 (R) has the property that its continuous surjective homomorphisms are open - in other words, it is totally minimal (cf. [10, 7.4], [34]). In particular, every image of SL2 (R) is complete, and thus cb -compact; in particular, ρSL2 (R) (SL2 (R)) is a complete dense subgroup of b(SL2 (R)), and so ρSL2 (R) is onto. Therefore, b(SL2 (R)) is a quotient of
SL2 (R). Theonly non-trivial quotient of SL2 (R) is P SL2 (R), which is not compact, because the

1 1
matrix A =
generates a discrete infinite subgroup. Hence, SL2 (R) is minimally almost
0 1
periodic.

G. Lukács / Lifted closure operators

9

Example 4.6. Nienhuys [29] showed that there exists a coarser metrizable group topology T on R
such that G := (R, T ) is minimally almost periodic. Thus, G is cb -compact, and it is also balanced, because G is abelian. Therefore, G is a balanced, connected, and cb -compact group, which is not compact.
Van der Waerden proved that every (algebraic) homomorphism from a compact connected semisimple Lie group into a compact group is continuous (cf. [38]). Groups K with this property satisfy K = b(Kd ), where Kd is the group K equipped with the discrete topology, and they are called van der Waerden groups, or briefly, vdW-groups (cf. [16]). For D = Kd , one obtains a discrete group that is cb -compact.
Example 4.7. SO3 (R) is a vdW-group, so SO3 (R)d is discrete, infinite, and cb -compact. Thus,
SO3 (R)d is a balanced, locally compact, and cb -compact group, which is not compact.
A closure operator c is hereditary if for every m : M → A in subA A and n : N → N in subA N, cM (n) = m−1 (cA (n)); c is weakly-hereditary if for every m : M → A in subA A, m is c-dense in the domain of cA (m).
Theorem 4.8. The lift cρ need not be weakly-hereditary when c is hereditary, even if A is Ereflective in X and F = E ∩ mor A.
P ROOF. As we mentioned at the beginning of this section, groups whose Bohr-compactification is an embedding are precisely the precompact ones, and for them cb coincides with the Kuratowski closure. Hausdorff precompact groups form an Onto-reflective subcategory of Grp(Top), and the reflection G → b+ G is given by the image of G in bG. Therefore, cb can also be seen as the b+ -initial lift of the hereditary Kuratowski closure operator of the precompact Hausdorff groups.
Hence, it suffices to show that cb is not weakly-hereditary.
The Bohr-closure of the trivial subgroup of a topological group G is simply its von Neumann kernel, n(G). Had cb been weakly-hereditary, then n(n(G)) = (cb )n(G) ({e}) = (cb )G ({e}) = n(G)

(10)

would have been true for every topological group G, but this is not the case, as the Example 4.9
shows.
The example below was suggested by Dikran Dikranjan.
Example 4.9. For a prime number p 6= 2, let A = Z(p∞ ), the Prüfer group. It can be seen as the subgroup of Q/Z generated by the elements of p-power order, or the group formed by all pn th roots of unity. Since every proper subgroup of A is finite, if A is equipped with a Hausdorff group topology in which it is not minimally almost periodic (in other words, n(A) 6= A), then n(A) is a finite subgroup with the discrete topology, and thus n(n(A)) is trivial. Therefore, if τ
is a Hausdorff group topology on A such that (A, τ ) is neither minimally nor maximally almost periodic, then n(n(A, τ )) 6= n(A, τ ).
Following Zelenyuk and Protasov [40], a sequence {an } ⊆ G in a (discrete) abelian group G
τ
is said to be a T -sequence if there is a Hausdorff group topology τ on A such that an −→ 0. It

10

G. Lukács / Lifted closure operators

follows from the proof of [40, Theorem 1] that if {an } is a T -sequence and τ is the finest group
τ
topology such that an −→ 0, then (G, τ ) is universal in the following sense: a homomorphism
ϕ : G{an } → H into a topological group H is continuous if and only if ϕ(an ) → 0.
Let cn = p1n in A, put
1
1
1
1
1
bn = −c1 +cn3 −n2 +· · ·+cn3 −2n +cn3 −n +cn3 = − + n3 −n2 +· · ·+ n3 −2n + n3 −n + n3 , (11)
p p p
p p
and consider the sequence an defined as b1 , c1 , b2 , c2 , b3 , c3 , . . .. One can show that {an } is a T sequence in A; we chose to omit the proof of this statement because of its technical nature and length. We set τ to be the finest Hausdorff group topology on A such that an → 0 in τ , and apply
(9) from Remark 4.1 to show that n(A, τ ) = h p1 i.
\
To that end, let χ ∈ (A,
τ ); then χ(an ) → 0, and in particular, χ(cn ) → 0. By [40, Example 6]
and [9, 3.3], χ(cn ) → 0 is equivalent to χ = mχ1 , where m ∈ Z and χ1 is the natural embedding of Z(p∞ ) into Q/Z ⊆ R/Z. It is easily seen that χ1 (bn ) = − 1p . Thus, if χ(bn ) → 0, then −m p1 = 0
in R/Z, and so p | m. On the other hand, it follows from pχ1 (bn ) → −p p1 = 0 that pχ1 (an ) → 0,
\
so pχ1 is continuous with respect to τ . Therefore, (A,
τ ) = {mχ1 | m ∈ pZ}. Hence, h 1 i ⊆ ker χ
p

\
for every χ ∈ (A,
τ ). Since ker pχ1 = h 1p i, this shows that n(A, τ ) = h p1 i, as desired.
5.

APPLICATION III: THE ∗-REPRESENTATION TOPOLOGY

A ∗-algebra A is an algebra over C with an involution ∗ : A → A such that (a + λb)∗ = a∗ + λ̄b∗
and (ab)∗ = b∗ a∗ for every a, b ∈ A and λ ∈ C. A topological ∗-algebra is a ∗-algebra A and a topology on A making the operations (addition, multiplication, additive inverse, involution) jointly continuous. The category of topological ∗-algebras and their continuous ∗-homomorphisms is denoted by T∗ A. A C ∗ -[semi]norm on a ∗-algebra A is a [semi]norm p that is submultiplicative and satisfies the C ∗ -identity - in other words, p(ab) ≤ p(a)p(b) and p(a∗ a) = p(a)2 for every a, b ∈ A. We put N (A) for the set of continuous C ∗ -seminorms on a topological ∗-algebra A. For every p, q ∈ N (A), one has max{p, q} ∈ N (A) (max{p, q} is defined pointwise), which turns
N (A) into a directed set.
A C ∗ -algebra is a complete Hausdorff topological algebra whose topology is given by a single C ∗ -norm. The full subcategory of T∗ A formed by the C ∗ -algebras is denoted by C∗ A. For
A ∈ T∗ A, a ∗-representation of A is a continuous ∗-homomorphism π : A → B(H) of A into the
C ∗ -algebra of bounded operators on some Hilbert space H (i.e., a morphism in T∗ A). The class of
∗-representations of A is denoted by R(A).
Let p ∈ N (A); ker p = {a ∈ A | p(a) = 0} is a ∗-ideal in A, and p induces a C ∗ -norm on the quotient A/ ker p, so the completion Ap of this quotient with respect to p is a C ∗ -algebra.
By the celebrated Gelfand-Naı̆mark-Segal theorem, every C ∗ -algebra is ∗-isomorphic (and thus isometric) to a closed subalgebra of B(H) for a large enough Hilbert space H (cf. [11, 2.6.1]).
Thus, we obtain a ∗-representation πp : A → A/ ker p → Ap → B(H) such that p(x) = kπp (x)k.
Conversely, each π ∈ R(A) gives rise to a C ∗ -seminorm pπ (x) = kπ(x)k. Therefore, the initial topology TA induced by the class R(A) coincides with the one induces by the family of C ∗ seminorms N (A). The topology TA is called the ∗-representation topology, and it defines a closure

G. Lukács / Lifted closure operators

11

operator c∗ on subalgebras in T∗ A. The closure c∗A ({0}) of the trivial subalgebra is a closed ∗-ideal of A; it is called the reducing ideal of A, and denoted by AR (cf. [30, 9.7]). Notice that c∗A ({0}) = AR =

\

ker p =

p∈N (A)

\

ker π,

(12)

π∈R(A)

so TA is Hausdorff (i.e., AR = 0) if and only if the ∗-representations of A separate the points of A.
A topological ∗-algebra A is said to be a pro-C ∗ -algebra if it is Hausdorff, complete, and its topology is generated by a family of C ∗ -seminorms - in other words, the topology of A coincides with TA , it is complete, and AR = 0. One can show that A is a pro-C ∗ -algebra if and only if
A is the limit in T∗ A of C ∗ -algebras (cf. [32, 1.1.1]). The full subcategory of T∗ A consisting of pro-C ∗ -algebras is denoted by P∗ A. These algebras were studied under various names (locally
C ∗ -algebras, LMC*-algebras, etc.) by numerous authors; for instance, in section 3 of [12], Dubuc and Porta essentially characterized commutative pro-C ∗ -algebras (up to a k-ification). Alluding only to a few more highlights, we mention the works of Inoue [21], Schmüdgen [37], Phillips [31]
& [32], Bhatt and Karia [2], and Inoue and Kürsten [22].
Proposition 5.1. P∗ A is a reflective subcategory of T∗ A.
P ROOF. Let A ∈ T∗ A. For each p ∈ N (A), we put Ap to be the completion of A/ ker p with respect to the C ∗ -norm that p induces on it; obviously, Ap is a C ∗ -algebra. For every q ≥ p in the directed set N (A), there is a surjective morphism Aq → Ap . This defines a functor
A∗ : N (A) −→ C∗ A ⊂ T∗ A
p 7−→ Ap ,

(13)
(14)

and so we define the reflector
P C ∗ (A) = lim A∗ ,

(15)

where the limit is taken in T∗ A. It follows from the definition that P C ∗(A) is a pro-C ∗ -algebra.
To show that P C ∗ (−) : T∗ A −→ P∗ A is a functor, let ϕ : A → B be a morphism in T∗ A. Then
N (ϕ) : N (B) → N (A) defined by r 7→ rϕ is an order-preserving map, and ϕr : Arϕ → Br is a natural transformation from A∗ N (ϕ) to B∗ . Thus, there is a morphism lim A∗ N (ϕ) → lim B∗ , and therefore one sets P C ∗(ϕ) = (lim A∗ → lim A∗ N (ϕ) → lim B∗ ).
To complete the proof, one defines ηA : A → P C ∗ (A) by setting ηA (x) = (x + ker p)p∈N (A) .
Obviously, ηA is a morphism (because each A → Ap is so), and it is a natural transformation. Let x = (xp )p∈N (A) ∈ P C ∗ (A), and pick p1 , . . . , pk ∈ N (A). Then q = max{p1 , . . . , pk } ∈ N (A), and Aq → Api is surjective. The image of A under the canonical morphism A → Aq is dense, so for every ε > 0, there is a ∈ A such that q(āq − xq ) < ε, and thus pi (āpi − xpi ) < ε (where āp stands for the image of a in Ap ). Therefore, the image ηA (A) is dense in P C ∗ (A). If B ∈ P∗ A, then ηB = idB . Hence, if ϕ : A → B is a morphism from A ∈ T∗ A into B ∈ P∗ A, then
ϕ = ηB ϕ = P C ∗ (ϕ)ηA . This decomposition is unique, because ηA (A) is dense in P C ∗(A).

12

G. Lukács / Lifted closure operators

We turn now to fitting all these into the setting of section 2. Equip X = T∗ A with the usual
(Onto, Embed)-factorization system, and provide A = P∗ A with the (Dense, ClEmb)-factorization system (Dense and ClEmb stand for the maps with a dense image and closed embeddings, respectively). The Onto-reflective hull of A in X is the category B = P∗ A consisting of the algebras
A ∈ T∗ A such that AR = 0 and TA coincides with the topology of A; the reflector P C ∗ is given by
A 7−→ ηA (A), and the reflection αA : A → P C ∗ (A) is the same as η, but with different codomain.
The category B inherits the (Onto, Embed)-factorization system of A.
Theorem 5.2.
(a) c∗ = K α = K η , in other words, the ∗-representation topology is the P C ∗ -initial [P C ∗ -initial]
lift of the Kuratowski closure K on P∗ A [P∗ A].
(b) c∗ is finitely productive.
(c) An algebra A ∈ T∗ A is c∗ -separated if and only if αA is injective, or equivalently, AR = 0; in particular, each algebra in P∗ A is c∗ -separated.
(d) An algebra A ∈ T∗ A is c∗ -compact in T∗A if and only if P C ∗ (A) is K-compact in P∗ A; in this case, ηA is surjective.
P ROOF. (a) It follows from the definition that TA is the initial topology on A induced by the family of *-homomorphisms (A → ApQ
)p∈N (A) . Equivalently, TA is the initial topology induced by the single *-homomorphism A →
Ap , whose image is contained in P C ∗ (A) ⊆ P C ∗ (A).
p∈N (A)

Therefore, the statement follows from Proposition 2.2.
(b) Since P C ∗ (−) is additive, it preserves finite products in T∗ A. Therefore, the statement follows from Proposition 2.4, as K is productive, and surjective morphisms in P∗ A are productive.
(c) follows from Proposition 2.5, because P C ∗ (−) preserves finite products, and monomorphism in T∗ A are injections.
(d) It would be tempting to apply Theorem 2.7 to A, but unfortunately, the condition of F ⊂ E
fails: in our case, F = Dense in P∗ A while E = Onto. Instead, we apply Corollary 2.10 to
B = P∗ A in order to obtain the first statement, while the second one is a consequence of Proposition 2.6(b).
To conclude, we investigate the class of topological ∗-algebras whose ∗-representation topology is generated by a single continuous seminorm. As the next lemma (which is modeled on [30,
10.1.7]) reveals, these are precisely the algebras for which P C ∗ (A) is a C ∗ -algebra.
Lemma 5.3. Let A ∈ T∗ A. The following statements are equivalent:
(i) For every x ∈ A,
γ(x) = sup{p(x) | p ∈ N (A)} = {kπ(x)k | π ∈ R(A)} < ∞, and γ is continuous on A;
(ii) N (A) has a largest element;

(iii) P C ∗ (A) is a C ∗ -algebra;

(iv) TA is defined by a single continuous linear space seminorm on A.

(16)

G. Lukács / Lifted closure operators

13

Remark 5.4. The seminorm in (iv) is not assumed a priori to be submultiplicative or a C ∗ seminorm.
P ROOF. (i) ⇒ (ii) Since each p ∈ N (A) is a C ∗ -seminorms, so is γ, which (being continuous)
belongs to N (A). Thus, γ is the largest element of N (A).
(ii) ⇒ (iii) Let r ∈ N (A) be the largest element. Then Ar is a C ∗ -algebra, and for each p ∈ N (A) there is a morphism Ar → Ap . Thus, P C ∗ (A) = lim A∗ = Ar is a C ∗ -algebra.
(iii) ⇒ (iv) Let k · k be the norm of the C ∗ -algebra P C ∗ (A). It is certainly a C ∗ -seminorm on
A, and by definition, it defines TA . In particular, k · k is continuous.
(iv) ⇒ (i) Let σ be a seminorm that defines TA . Since multiplication is jointly continuous in P C ∗ (A) (which carries the quotient topology TA /AR ), there is a constant B > 0 such that
σ(xy) ≤ Bσ(x)σ(y) for every x, y ∈ A. Thus, by replacing σ with Bσ if necessary, we may assume that B = 1 and σ is submultiplicative. Let π : A → B(H) ∈ R(A). Pick a ∈ A. Since
B(H) is a C ∗ -algebra,
1

kπ(a)k2 =kπ(a∗ a)k = rB(H) (π(a∗ a)) = lim k(π(a∗ a))n k n .
n→∞

(17)

By the continuity of π with respect to TA , there is a constant Cπ such that kπ(x)k ≤ Cπ σ(x) for every x ∈ A, so p
p
1
1
(18)
k(π(a∗ a))n k n ≤ n Cπ σ((a∗ a)n ) n ≤ n Cπ σ(a∗ a).

Therefore,

p
1
kπ(a)k2 = lim k(π(a∗ a))n k n ≤ lim n Cπ σ(a∗ a) = σ(a∗ a) ≤ Dσ(a)2 ,
(19)
n→∞
n→∞
√
where D is a constant such that σ(x∗ ) ≤ Aσ(x) for every x ∈ A. Hence, γ(a) ≤ Dσ(a), which proves the second statement too, because σ is assumed to be continuous on A.
Palmer [30, 10.1] investigated discrete ∗-algebras that satisfied the equivalent conditions of
Lemma 5.3, and named them G∗ -algebras. Thus, we extend this terminology to topological ∗algebras too, and denote by G∗ A the full subcategory of T∗ A formed by the (generalized) G∗ algebras. For A ∈ G∗ A, we put C ∗ (A) for P C ∗ (A), and call it the enveloping C ∗ -algebra of A.
The next Proposition is a consequence of Proposition 5.1 and Lemma 5.3; its restricted version, for discrete G∗ -algebras, appears in [30, 10.1.11].
Proposition 5.5. The category C∗ A is a full reflective subcategory of G∗ A, and its reflector is given by A 7−→ C ∗ (A).

We once again return to the setting of section 2. Equip X = G∗ A with the (Onto, Embed)factorization system, and provide A = C∗ A with the (Onto, Inj)-factorization system (Inj stands for injections). Notice that every injection of C ∗ -algebras is actually an embedding (cf. [11, 1.8.1]), so the Kuratowski closure on C ∗ -subalgebras of C ∗ -algebras is just the discrete closure s. (This is not too surprising, though, because van Osdol [39] showed that the category of unital C ∗ -algebras and their unital ∗-homomorphisms is monadic over Set.) Therefore, by Corollary 2.12, we get:
Theorem 5.6. An algebra A ∈ G∗ A is c∗ -compact if and only if αA : A → C ∗ (A) is surjective.

14

G. Lukács / Lifted closure operators
ACKNOWLEDGMENTS

I am deeply indebted to Professor Walter Tholen, my PhD supervisor, who introduced me to the concept of closure operators, for his attention to this work and his helpful suggestions.
I am grateful to Professor Dikran Dikranjan, whose talk at CITA-2003 was the main source of inspiration for this research, for his sharp and constructive criticism of my early notes for the paper, and for numerous valuable suggestions at the later stages.
I wish to thank Professor Salvador Hernández, Professor Maria Manuel Clementino, and Professor Bob Paré for the valuable discussions that were of great assistance in writing this paper.
REFERENCES

[1] S. Baron. Reflectors as compositions of epi-reflectors. Trans. Amer. Math. Soc., 136:499–
508, 1969.
[2] Subhash J. Bhatt and Dinesh J. Karia. An intrinsic characterization of Pro-C ∗ -algebras and its applications. J. Math. Anal. Appl., 175(1):68–80, 1993.
[3] M. M. Clementino, E. Giuli, and W. Tholen. Topology in a category: compactness. Portugal.
Math., 53(4):397–433, 1996.
[4] M. M. Clementino, E. Giuli, and W. Tholen. A functional approach to general topology.
In M. C. Pedicchio and W. Tholen, editors, Categorical Foundations. Cambridge University
Press, 2003.
[5] D. Dikranjan and W. Tholen. Categorical structure of closure operators, volume 346 of
Mathematics and its Applications. Kluwer Academic Publishers Group, Dordrecht, 1995.
With applications to topology, algebra and discrete mathematics.
[6] Dikran Dikranjan and Eraldo Giuli. Closure operators. I. In Proceedings of the 8th international conference on categorical topology (L’Aquila, 1986), volume 27, pages 129–143,
1987.
[7] Dikran Dikranjan and Eraldo Giuli. Compactness, minimality and closedness with respect to a closure operator. In Categorical topology and its relation to analysis, algebra and combinatorics (Prague, 1988), pages 284–296. World Sci. Publishing, Teaneck, NJ, 1989.
[8] Dikran Dikranjan, Eraldo Giuli, and Walter Tholen. Closure operators. II. In Categorical topology and its relation to analysis, algebra and combinatorics (Prague, 1988), pages 297–
335. World Sci. Publishing, Teaneck, NJ, 1989.
[9] Dikran Dikranjan, Chiara Milan, and Albert Tonolo. A characterization of the maximally almost periodic abelian groups. J. Pure Appl. Algebra, to appear.
[10] Dikran N. Dikranjan, Ivan R. Prodanov, and Luchezar N. Stoyanov. Topological groups, volume 130 of Monographs and Textbooks in Pure and Applied Mathematics. Marcel Dekker
Inc., New York, 1990. Characters, dualities and minimal group topologies.

G. Lukács / Lifted closure operators

15

[11] Jacques Dixmier. C ∗ -algebras. North-Holland Publishing Co., Amsterdam, 1977. Translated from the French by Francis Jellett, North-Holland Mathematical Library, Vol. 15.
[12] Eduardo J. Dubuc and Horacio Porta. Convenient categories of topological algebras, and their duality theory. J. Pure Appl. Algebra, 1(3):281–316, 1971.
[13] Ryszard Engelking. General topology, volume 6 of Sigma Series in Pure Mathematics. Heldermann Verlag, Berlin, second edition, 1989. Translated from the Polish by the author.
[14] Siegfried Grosser and Martin Moskowitz. On central topological groups. Trans. Amer. Math.
Soc., 127:317–340, 1967.
[15] Siegfried Grosser and Martin Moskowitz. Representation theory of central topological groups. Trans. Amer. Math. Soc., 129:361–390, 1967.
[16] Joan E. Hart and Kenneth Kunen. Bohr Compactifications of Non-Abeilan Groups. Topology
Proc., 26(2):593–626, 2001/02.
[17] Horst Herrlich. Topologische Reflexionen und Coreflexionen. Lecture Notes in Mathematics,
No. 78. Springer-Verlag, Berlin, 1968.
[18] Edwin Hewitt and Kenneth A. Ross. Abstract harmonic analysis. Academic Press Inc.,
Publishers, New York, 1963.
[19] Ralph Hughes. Weak compactnes and topological groups. PhD thesis, University of North
Carolina at Chapel Hill, 1972.
[20] Ralph Hughes. Compactness in locally compact groups. Bull. Amer. Math. Soc., 79:122–123,
1973.
[21] Atsushi Inoue. Locally C ∗ -algebra. Mem. Fac. Sci. Kyushu Univ. Ser. A, 25:197–235, 1971.
[22] Atsushi Inoue and Klaus-Detlef Kürsten. On C ∗ -like locally convex ∗-algebras. Math. Nachr.,
235:51–58, 2002.
[23] Tadashi Ishii. On the Tychonoff functor and w-compactness. Topology Appl., 11(2):173–187,
1980.
[24] Tadashi Ishii. On the Tychonoff functor and k-spaces. Questions Answers Gen. Topology,
7(2):129–133, 1989.
[25] Tadashi Ishii. The Tychonoff functor and related topics. In Topics in general topology, pages
203–243. North-Holland, Amsterdam, 1989.
[26] George Janelidze and Walter Tholen. Facets of descent. I. Appl. Categ. Structures, 2(3):245–
281, 1994.
[27] Calvin C. Moore. Groups with finite dimensional irreducible representations. Trans. Amer.
Math. Soc., 166:401–410, 1972.

16

G. Lukács / Lifted closure operators

[28] J. Neumann and E. P. Wigner. Minimally almost periodic groups. Ann. of Math. (2), 41:746–
750, 1940.
[29] J. W. Nienhuys. A solenoidal and monothetic minimally almost periodic group. Fund. Math.,
73(2):167–169, 1971/72.
[30] Theodore W. Palmer. Banach algebras and the general theory of ∗-algebras. Vol. 2, volume 79 of Encyclopedia of Mathematics and its Applications. Cambridge University Press,
Cambridge, 2001. ∗-algebras.
[31] N. Christopher Phillips. Inverse limits of C ∗ -algebras. J. Operator Theory, 19(1):159–195,
1988.
[32] N. Christopher Phillips. Inverse limits of C ∗ -algebras and applications. In Operator algebras and applications, Vol. 1, volume 135 of London Math. Soc. Lecture Note Ser., pages 127–185.
Cambridge Univ. Press, Cambridge, 1988.
[33] L. S. Pontryagin. Selected works. Vol. 2. Gordon & Breach Science Publishers, New York, third edition, 1986. Topological groups, Edited and with a preface by R. V. Gamkrelidze,
Translated from the Russian and with a preface by Arlen Brown, With additional material translated by P. S. V. Naidu.
[34] Dieter Remus and Luchezar Stojanov. Complete minimal and totally minimal groups. Topology Appl., 42(1):57–69, 1991.
[35] Dieter Remus and F. Javier Trigos-Arrieta. The Bohr topology of Moore groups. Topology
Appl., 97(1-2):85–98, 1999. Special issue in honor of W. W. Comfort (Curacao, 1996).
[36] Lewis C. Robertson. A note on the structure of Moore groups. Bull. Amer. Math. Soc.,
75:594–599, 1969.
[37] Konrad Schmüdgen. Über LMC-Algebren. Math. Nachr., 68:167–182, 1975.
[38] B.L. van der Waerden. Stetigkeitsstze fr halbeinfache Liesche Gruppen. Math. Z., 36:780–
786, 1933.
[39] D. H. Van Osdol. C ∗ -algebras and cohomology. In Categorical topology (Toledo, Ohio,
1983), volume 5 of Sigma Ser. Pure Math., pages 582–587. Heldermann, Berlin, 1984.
[40] E. G. Zelenyuk and I. V. Protasov. Topologies on abelian groups. Izv. Akad. Nauk SSSR Ser.
Mat., 54(5):1090–1107, 1990.

FB 3 - Mathematik und Informatik
Universität Bremen
Bibliothekstrasse 1
28359 Bremen
Germany

G. Lukács / Lifted closure operators

Current address:
Department of Mathematics and Statistics
Dalhousie University
Halifax, B3H 3J5, Nova Scotia
Canada e-mail: lukacs@mathstat.yorku.ca

17

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
