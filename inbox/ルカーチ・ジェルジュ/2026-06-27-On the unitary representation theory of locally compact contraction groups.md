---
source: "https://arxiv.org/abs/2311.03041v2"
title: "On the unitary representation theory of locally compact contraction groups"
author: "Max Carter"
year: "2023"
publication: "arXiv preprint / math.GR"
download: "https://arxiv.org/pdf/2311.03041v2"
pdf: "https://arxiv.org/pdf/2311.03041v2"
captured_at: "2026-06-27T06:15:32Z"
updated_at: "2026-06-27T06:15:32Z"
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

# On the unitary representation theory of locally compact contraction groups

- 著者: Max Carter
- 年: 2023
- 掲載情報: arXiv preprint / math.GR
- 情報源: [arxiv](https://arxiv.org/abs/2311.03041v2)
- ダウンロード: https://arxiv.org/pdf/2311.03041v2
- PDF: https://arxiv.org/pdf/2311.03041v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

The unitary representation theory of locally compact contraction groups and their semi-direct products with $\mathbb{Z}$ is studied. We put forward the problem of completely characterising such groups which are type I or CCR and this article provides a stepping stone towards a solution to this problem. In particular, we determine new examples of type I and non-type-I groups in this class, and we completely classify the irreducible unitary representations of the torsion-free groups, which are shown to be type I. When these groups are totally disconnected, they admit a faithful action by automorphisms on an infinite locally-finite regular tree; this work thus provides new examples of automorphism groups of regular trees with interesting representation theory, adding to recent work on this topic.

## PDF Text

arXiv:2311.03041v2 [math.GR] 14 Oct 2024

ON THE UNITARY REPRESENTATION THEORY OF LOCALLY
COMPACT CONTRACTION GROUPS
MAX CARTER
Abstract. The unitary representation theory of locally compact contraction groups and their semi-direct products with Z is studied. We put forward the problem of completely characterising such groups which are type I or CCR
and this article provides a stepping stone towards a solution to this problem.
In particular, we determine new examples of type I and non-type-I groups in this class, and we completely classify the irreducible unitary representations of the torsion-free groups, which are shown to be type I. When these groups are totally disconnected, they admit a faithful action by automorphisms on an infinite locally-finite regular tree; this work thus provides new examples of automorphism groups of regular trees with interesting representation theory, adding to recent work on this topic.

1. Introduction
The unitary representation theory of locally compact groups enjoys an extensive history with strong connections to many areas of mathematics including physics and number theory (c.f. [30, 41, 62]). The subject dates back to the early 1900’s, with the pioneering work of Peter-Weyl on the representation theory of compact groups [51], and the work of Pontryagin on the representation theory of locally compact abelian groups [52]. Over the intervening years, a substantial theory of unitary representations for general locally compact groups has been developed, amongst which the work of Mackey is most notable; see [44] for an overview of his work. The strong interplay between unitary representations and operator algebras is critical to the theory and a fruitful area of study in its own right [47, 19, 20, 15, 5].
In the unitary representation theory of locally compact groups, an important question is, when given a locally compact group G, can we determine all the irreducible unitary representations of G up to unitary equivalence. The collection of all equivalences classes of such representations forms a set, called the unitary dual b The set G
b is equipped with a topology, called the Fell of G, and is denoted by G.
Date: October 16, 2024.
2020 Mathematics Subject Classification. 20C25, 22D10, 22D12, 22D25, 20G05, 43A65.
Key words and phrases. Unitary representation, type I group, CCR group, scale group, contraction group, unipotent linear algebraic group, amenable group, groups acting on trees.
The author acknowledges support from an Australian Government Research Training Program
(RTP) scholarship and from the FWO and F.R.S.-FNRS under the Excellence of Science (EOS)
program (project ID 40007542).
1

2

MAX CARTER

topology, and a Borel structure, called the Mackey-Borel structure. It is a consequence of an important result of Glimm [25], referred to as Glimm’s theorem, that b if and only if one of the following equivalent conditions essentially we can classify G
is satisfied: G is type I ; the group C∗ -algebra C ∗ (G) is GCR; the Fell topology on b is T0 ; the Mackey-Borel structure on G
b is countably separated [22, Theorem 7.6].
G
A locally compact group G is called type I if every unitary representation of G
generates a type I von Neumann algebra. There is a strong dichotomy between type I groups and those which are not type I. The unitary representation theory of type I locally compact groups is quite tame: every unitary representation of a type
I locally compact group decomposes uniquely into a direct integral of irreducible unitary representations [15, §8.6.6], so the classification of unitary representations of a type I locally compact group reduces to classifying the irreducible ones. Moreover, as already alluded to, type I groups are, in some sense, those groups whose irreducible unitary representations can be classified in a ‘measurable’ way. For locally compact groups which are not type I, however, pathological things occur: the prior mentioned direct integral decomposition is non-unique in a strong sense (c.f.
[22, Theorem 7.41] and [15, §8.7]), and consequently, the study of general unitary representations can no longer be reduced to studying irreducible ones. Also, the complexity of classifying all irreducible unitary representations of a group which is not type I is extremely high, and in most cases leaves the problem of classifying all irreducible unitary representations of the group completely out of reach (see
[59, 61] for example, also refer to Glimm’s theorem again).
It is thus a fundamental problem in unitary representation theory to determine which groups are type I. A substantial amount of work has been done on this problem over many decades and a comprehensive but non-exhaustive list of type I
locally compact groups can be found in [5, Thm 6.E.19, Thm 6.E.20, Thm 7.D.1].
In recent years, the problem of determining which groups of automorphisms of regular trees are type I has been an important problem in the representation theory of totally disconnected locally compact (tdlc) groups [49, 1, 12, 33, 11, 55, 56].
Automorphism groups of regular trees are often described by the experts as a
‘microcosm’ for the theory of (simple) tdlc groups [10] which partially elucidates the importance of understanding the representation theory of such groups. For groups of automorphisms of regular trees that are unimodular and act transitively on the vertices of the tree, it was conjectured by Nebbia in [49] that the group C∗ -algebra is CCR (see [22, Page 207] for definition), which is stronger than the property of the group being type I, if and only if the group acts transitively on the boundary of the tree. Nebbia resolved the ‘only if’ direction in the same paper, however, the ‘if’ direction still remains open. Recent work by Semal in [55] suggests that a stronger version of the ‘if’ direction may even hold. More recently than Nebbia’s work, it was conjectured in [33] that for non-amenable automorphism groups of regular trees that act ‘minimally’ on the tree, the group is type I if and only if it acts transitively on the boundary of the tree. Parts of the conjecture were resolved in [33] and improved upon in [11]. The conjecture, however, also still remains open.

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

3

The research described in the above paragraph mostly concerns automorphism groups of trees that are non-amenable, act (highly) transitively on the boundary of the tree, and in some cases satisfy Tits’ independence property or one of its variants
[60, 3]. At the time of writing this article, the (non-compact) amenable automorphism groups of trees have not received as much attention as their non-amenable counterparts, at least in terms of the study of their unitary representations. The only article that the author is aware of that studies unitary representations of noncompact amenable automorphism groups of regular trees is a short note by Nebbia
[48].
The present article arose out of a desire to understand, on one hand, more about the unitary representation theory of amenable automorphism groups of regular trees, and on the other hand, begin research on the unitary representation theory of the class of scale groups [65] which have recently risen to importance in the theory of tdlc groups. A scale group is a closed subgroup of the automorphism group of a regular tree that fixes a boundary point and acts transitively on the vertices of the tree; such groups are necessarily amenable [21, Chp I, Thm 8.1].
It is an important fact that every tdlc group with non-trivial scale function has a subquotient isomorphic to a scale group [4], so scale groups have connections with the broader theory of tdlc groups, and they also have many connections with the theory of self-replicating groups [32, 65].
At the present time, completely characterising the type I scale groups and their irreducible unitary representations seems to be out of reach, so in this article we investigate the unitary representation theory of a particular subclass of scale groups which we call contractive scale groups. Every scale group is a semi-direct product of the form N ⋊ hti, where N is the normal subgroup of automorphisms that fix at least one vertex, and t is an automorphism of the tree that translates towards the fixed boundary point [65, Proposition 2.2]. A scale group N ⋊ hti is contractive if for all x ∈ N , tn xt−n → id as n → ∞.
We abstract the above setup as follows so that the action on the tree is no longer present; we do this for simplification as our methods in this article do not require the tree-action and they rely soley on the already known structure theory of these groups. A locally compact contraction group is a pair (N, α), where N is a locally compact group, and α ∈ Aut(N ) is a contractive automorphism; see [26, 27, 28] for more details on the theory of contraction groups and their connections to the theory of (td)lc groups. Given any locally compact contraction group (N, α), we can look at the corresponding semi-direct product N ⋊α Z. If N is totally disconnected, then it turns out that the group N ⋊α Z admits a faithful action by automorphisms on a regular tree with respect to which it is a contractive scale group; this is a consequence of the tree representation theorem [4, Theorem 4.1]. We thus get a correspondence between tdlc contraction groups (N, α) and contractive scale groups
N ⋊α Z.
The problem that this article works towards solving is the following.
Problem 1.1. Let (N, α) be a locally compact contraction group. Determine when
N , or its semi-direct product N ⋊α Z, is type I or CCR.

4

MAX CARTER

The main result of this article is the following theorem which characterises a large class of type I contractive scale groups and shows that they are never CCR
groups. This is a new situation in the context of groups acting on trees, as for many of the known type I automorphism groups of regular trees, they are shown to be type I by showing that they satisfy the stronger CCR property. This is, however, expected, as these scale groups are, in some sense, totally disconnected analogues of the real ax + b group and the representation theories of these groups are analogous in many ways. We also remark that the result [17, pg. 240, Cor. 2] implies that the left-regular representation of the groups in the following theorem are type I when
N is CCR, however, this does not imply that the groups themselves are type I.
Theorem 1.2. Let (N, α) be a locally compact contraction group and G := N ⋊α Z.
The group G is not CCR. Furthermore, if N is assumed to be CCR, then the following hold:
(i) G is type I; b . Then
(ii) Let X be a cross-section of the non-trivial orbits of the action Z y N
G
b b
G = {indN π : π ∈ X} ∪ Z.
It is shown in [26] that every locally compact contraction group (N, α) is of the form
N∼
= N0 × Np1 × · · · × Npn × tor(N )

where N0 is the connected component of the identity in N , Npi is a nilpotent pi -adic
Lie group for each i where p1 , . . . , pn are distinct primes, and tor(N ) the subgroup of torsion elements of N . Each of the factors in this decomposition are invariant under the action of α. Furthermore, it is known that N0 is a connected simplyconnected nilpotent real Lie group [57], each of the Npi are the Qpi -rational points of a unipotent linear algebraic group over Qpi [63], and tor(N ) is totally disconnected
[26]. The Kirillov orbit method [37, 38, 39] provides a classification of the irreducible unitary representations of N0 and shows that this group is CCR. Also, an extension of the Kirillov orbit method to unipotent linear algebraic groups over Qp by Moore
[46] gives a classification of the irreducible unitary representations of each of the
Npi and shows that they are CCR groups too. It thus follows that all torsion-free locally compact contraction groups are CCR and hence Theorem 1.2 applies to these groups. In Theorem 4.6 of this article, we provide an analogue of the Kirillov orbit method for scale groups of the form N ⋊α Z, where (N, α) is a torsion-free locally compact contraction group, and completely classify the irreducible unitary representations of these semi-direct products in terms of ‘coadjoint orbits’.
Much less is known about both the group-theoretic and the representationtheoretic properties of torsion locally compact contraction groups in comparison to the torsion-free case. It is shown in [27] that locally pro-p torsion locally compact contraction groups are nilpotent and their composition series, which are known to have finite length, are described in [26] (we note that these result hold in the torsionfree case too). Also, in the torsion case we see non-type-I phenomena occurring: it is shown in the proof of [11, Proposition 5.8] that if F is a finite
Q
L simple non-abelian group, then the torsion locally compact contraction group ( Z<0 F ) × ( Z⩾0 F )

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

5

equipped with the right-shift automorphism is a non-type-I group. This, unfortunately, is about the extent of our general knowledge on torsion locally compact contraction groups and their representation theory at the current time. Also, as far as the author is aware, prior to the writing of this article, there are no other known classes of non-type-I torsion locally compact contraction groups.
The final section of this article initiates research on understanding the unitary representation theory of torsion locally compact contraction groups. We use unique and distinct methods to [11] to show that a countable subset of an uncountable set of torsion locally compact contraction groups is non-type-I, providing a second distinct class of examples of non-type-I torsion locally compact contraction groups
(see Theorem 5.9). We pose the question as to whether all uncountably many of these groups are non-type-I too. Some difficulties prevented us from solving this problem using the techniques contained in this article, but a follow up article with a second author will treat this problem using more algebraic techniques. We finish the article with some brief discussion regarding the representation theory of unipotent linear algebraic groups over Fp ((t)). A natural question leading on from this project is to determine which unipotent linear algebraic groups over Fp ((t))
are type I or CCR. This question has been investigated multiple times in the past by prominent authors [34, 35, 18], however, the question remains unresolved. We give an elementary proof that the n-dimensional Heisenberg groups over Fp ((t)) are
CCR groups. This also shows, in particular, that there exist CCR locally compact contraction groups with non-abelian torsion subgroups.
Our arguments in this article rely heavily on the Mackey little group method
(sometime referred to as the Mackey machine) and its generalisations to C∗ -dynamical systems and crossed-product C∗ -algebras [43, 44, 20, 19, 53, 64, 36, 14]. Giving a very simplified description, the Mackey little group method allows one to do the following: given a separable locally compact group G and a closed normal type I
subgroup N of G, provided certain technical assumption hold, one can construct all the factor unitary representations (and in particular irreducible unitary representations) of G from factor unitary representations of certain subgroups of G/N
referred to as the ‘Little groups’. This process of constructing representations of
G from representations of subgroups of G/N preserves irreducibility and the type of the representations and this is critical to our arguments. The method also relies b (or of the related C∗ -dynamical heavily on analysing the orbits of the action G y N
∗
b /G.
system G y C (N )) and topological properties of the quotient space N
We give a detailed summary of the Mackey little group method along with other related representation and operator-algebra theoretic concepts in the preliminaries section of the article. We provide an extensive preliminaries section as we anticipate that there will be a broad range of mathematicians interested in this work, some coming from group theory, while others coming from representation theory and/or operator algebras.

6

MAX CARTER

2. Preliminaries
In this section we collect the notation and results that will be most critical to understanding the article. It is assumed that the reader has some familiarity with the standard results of functional analysis [54, 13], topological group theory [31, 22],
C∗ -algebras and their representations [15], and von Neumann algebras [16].
2.1. Graphs and groups. Let V Γ be a set and EΓ ⊆ {{u, v} : u, v ∈ V Γ}. The pair Γ = (V Γ, EΓ) is a graph. The elements of V Γ are the vertices and elements of EΓ the edges. For v ∈ V Γ, the cardinality of the set E(v) := {e ∈ EΓ : v ∈ e}
is the degree of v, denoted deg(v), and we call Γ locally-finite if deg(v) is finite for all v ∈ V Γ. A graph is infinite if V Γ has infinite cardinality.
Let I be a countable indexing set and (vi )i∈I a sequence in V Γ. The sequence is called a path if vi 6= vi+1 and {vi , vi+1 } ∈ EΓ for all i ∈ I. If (vi )i∈I is a path and I has finite cardinality, then we define its length to be |I| − 1, otherwise, we say the path is a ray if I = N or bi-infinite if I = Z.
If I := {0, 1, . . . , n} and (vi )i∈I is a path, the vertices v0 and vn are the endpoints of the path, and we say that the path is a cycle if v0 = vn . For v ∈ V Γ, an edge of the form {v, v} is called a loop. The graph Γ is connected if for any two vertices u, v ∈ V Γ, there is a path in Γ with endpoints u and v. A connected graph with no loops or cycles is a tree. The infinite tree in which every vertex has degree d, called the d-regular tree, or a regular tree for short, is denoted by Td .
Fix an infinite tree T = (V T, ET ). Two rays in T are said to be equivalent if their intersection is also a ray. It can be checked that this is an equivalence relation on the set of all rays in T . The set of equivalence classes of rays under this equivalence relation is called the boundary of T and denoted by ∂T . The elements of ∂T are the ends (or boundary points) of T .
If Γ and Λ are graphs, a homomorphism φ : Γ → Λ is a function φ : V Γ → V Λ
such that for any {u, v} ∈ EΓ, {φ(u), φ(v)} ∈ EΛ. An automorphism of Γ is a bijective graph homomorphism of Γ to itself. The group of automorphisms of a graph Γ is denoted by Aut(Γ). Throughout this paper, it will be assumed that
Aut(Γ) has the permutation topology: the topology whose base neighbourhoods are the collection of sets of the form U(g0 , F ) := {g ∈ Aut(Γ) : g(x) = g0 (x) for all x ∈
F }, with g0 ranging over all elements of the group Aut(Γ) and F ranging over all finite subsets of V Γ. It is assumed that any subgroup of Aut(Γ) has the induced topology. If Γ is locally-finite i.e. every vertex in Γ has finite degree, then Aut(Γ) is a totally disconnected locally compact second countable group when endowed with the permutation topology.
Given a group G acting on a set X, for any subset Y ⊆ X, define GY := {g ∈
G | g(Y ) = Y }, the stabiliser of Y under the action of G. If Y = {y} then we will write Gy instead of G{y} . Similarly, FixG (Y ) := {g ∈ G | g(y) = y ∀y ∈ Y } and is called the fixator of Y in G. For x ∈ X, G(x) := {g(x) | g ∈ G} denotes the orbit of x under G.
2.2. Scale groups. We begin by defining scale groups. The primary reference on scale groups is [65].

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

7

Definition 2.1 (Scale group). Let d ∈ N>2 and ω ∈ ∂Td . A subgroup G ⩽ Aut(Td )
which is closed, acts transitively on the vertices of Td and fixes ω is called a scale group.
Let N be the subgroup of the scale group G consisting of the elements of G that fix at least one vertex of Td .The subgroup N is normal in G. The orbits of the action N y V Td are called the horocycles of the end ω. The subgroup N will from now on be referred to as the horocycle stabiliser subgroup of G.
Proposition 2.2. [65, Proposition 2.2] Let G ⩽ Aut(Td ) be a scale group fixing the end ω ∈ ∂Td and N the subgroup of G that stabilises the horocycles of ω. There exists a translation t ∈ G, translating towards ω, such that G = N ⋊ hti.
Definition 2.3. The scale group G = N ⋊ hti is contractive if tn nt−n → idN as n → ∞ for all n ∈ N .
2.3. Contraction groups. The primary references on contraction groups are [26,
27, 28].
Definition 2.4 (Contraction group). Let N be a topological group and α an
(bicontinuous) automorphism of N . The pair (N, α) is called a contraction group if for all g ∈ N , αn (g) → idN as n → ∞. If N satisfies a property P (e.g. P =
locally compact, totally disconnected, torsion etc.) then (N, α) will be called a P
contraction group.
We now list some important examples of contraction groups that will be used in this article and set some corresponding notation.
Example 2.5. (i) The additive group (Qp , +) with automorphism multiplication by p.
(ii) The additive group (Fpn ((t)), +) of formal Laurent series with coefficients in the field Fpn is an abelian contraction group with respect to the automorphism
αt : Fpn ((t)) → Fpn ((t)), x 7→ tx.
(iii) Generalising the previous example, let F be a finite group. The group F ((t)) :=
Q
L
( Z<0 F )×( Z⩾0 F ) equipped with the automorphism αrs : F ((t)) → F ((t)), (xi )i∈Z 7→
(xi+1 )i∈Z is a locally compact contraction group. A special case of this that will be used in the next section is when F = Cpn , the cyclic group of order pn where p is a prime.
(iv) Following the notation and results in [27], let P denote the set of all primes and Q∞ := R. For p ∈ P ∪ {∞}, let Ωp ⊆ Qp [X] be the set of all monic irreducible polynomials whose roots are all non-zero and have absolute value
< 1 in some algebraic closure Qp of Qp . For f ∈ Ωp and n ∈ N, define Ef n :=
Qp [X]/f nQp [X]. The group Ef n is a locally compact contraction group when equipped with the automorphism αf n : Ef n → Ef n , g + f n Qp [X] 7→ Xg +
f n Qp [X]. Note that Ef n is isomorphic to a direct sum of finitely many copies of Qp .
(v) The group Un (Qp ) of n-dimensional unipotent matrices over Qp equipped with the automorphism which is conjugation by the dialgonal matrix diag(p, p2 , . . . , pn ).

8

MAX CARTER

The following is an important structure theorem about locally compact contraction groups that will be used throughout the article. Note that if (N, α) is a contraction group and M ⊆ N , then we call M α-stable if α(M ) = M . Similarly,
M is called fully invariant (resp. topologically fully invariant ) if f (M ) ⊆ M for every set map f : N → N (resp. every continuous set map f : N → N ).
Theorem 2.6. [27, Theorem A] For each locally compact contraction group (N, α), we have:
(i) Let N0 denote the connected component of the identity in N . Then N has a unique closed normal subgroup Ntd such that N = N0 × Ntd internally as a topological group. The subgroup Ntd is totally disconnected, α-stable, and topologically fully invariant.
(ii) The set tor(N ) of torsion elements of N is a closed subgroup of N and totally disconnected. There is a unique n ∈ N ∪ {0}, unique prime numbers p1 <
· · · < pn and unique p-adic Lie groups Np 6= {id} for p ∈ {p1 , . . . , pn } which are closed normal subgroups of N such that
N = N0 × Np1 × · · · × Npn × tor(N )
internally as a topological group. Each Np is topologically fully invariant in
N and hence α-stable.
(iii) If tor(N ) is locally pro-nilpotent, then, for each prime number p, the set torp (N ) of p-torsion elements of N is a fully invariant closed subgroup of
N which is locally pro-p. Moreover, torp (N ) 6= {idN } for only finitely many p, say for p among the prime numbers q1 < q2 < · · · < qm , and tor(N ) = torq1 (N ) × · · · × torqm (N )
internally as a topological group.
Remark 2.7. Recall, as mentioned in the introduction, the group N0 is always a connected simply-connected nilpotent real Lie group and the Npi are unipotent linear algebraic groups over Qp . See [27] for more details.
The following result classifies all abelian locally compact contraction groups up to isomorphism. Consult Example 2.5 for the definition of the notation.
Theorem 2.8. [27, Theorem E] Let A be a locally compact abelian group and suppose that α is an automorphism of A such that (A, α) is a contraction group.
Then (A, α) is isomorphic to
M M M
MM
(Ef n , αf n )µ(p,f,n) ⊕
(Cpn ((t)), αrs )ν(p,n)
p∈P∪{∞} f ∈Ωp n∈N

p∈P n∈N

as a contraction group, for uniquely determined µ(p, f, n) ∈ N ∪ {0} which are non-zero for only finitely many (p, f, n), and uniquely determined ν(p, n) ∈ N ∪ {0}
which are non-zero for only finitely many (p, n). Conversely, all groups of the above form are locally compact abelian contraction groups.

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

9

It is shown in [26] that for any locally compact contraction group (N, α), there exists a composition series
1 = N0 ⊳ · · · ⊳ Nm = N
of α-stable closed subgroups of N . The quotients Ni+1 /Ni are all simple contraction groups in the following sense.
Definition 2.9. A locally compact contraction group (N, α) is called a simple contraction group if it has no α-stable closed normal subgroups.
Glöckner and Willis have classified all of the simple locally compact contraction groups. Again, the reader should consult Example 2.5 for the definition of the notation in the following theorem.
Theorem 2.10. [26, 27] Let (N, α) be a simple locally compact contraction group.
The following hold:
(i) N is either connected or totally disconnected. If it is connected, then it is torsion-free, and if it is totally disconnected, then it is either torsion-free or torsion;
(ii) If N is connected, then it is abelian, and (N, α) is isomorphic to (Ef n , αf n )
for some f ∈ Ω∞ and n ∈ N;
(iii) If N is totally-disconnected and torsion-free, then it is abelian, and (N, α) is isomorphic to (Ef n , αf n ) for some f ∈ Ωp , n ∈ N and p ∈ P;
L
(iv) If N is totally-disconnected and torsion then (N, α) is isomorphic to (( Z<0 F )×
Q
( Z⩾0 F ), αrs ) for some finite simple group F .
The link between contraction groups, scale groups and groups acting on trees is the following theorem. Its proof is given by the tree representation theorem in
[4] and extensively uses the scale theory of totally disconnected locally compact groups.

Proposition 2.11. [4, Theorem 4.1] Let (N, α) be a totally disconnected locally compact contraction group. The group N ⋊α Z admits a faithful action by automorphisms on a regular tree Td such that it is a scale group.
It is clear by definition that any scale group of the form N ⋊α Z, where (N, α)
is a totally disconnected locally compact contraction group, is a contractive scale group.
2.4. Multipliers and representations. In this subsection we follow the notation and terminology used in [2, 40]. The reader can consult these reference for more details on the following.
Let G be a locally compact second countable group. A multiplier ω on G is a measurable function ω : G × G → T that satisfies the following properties:
(M1) ω(x, y)ω(xy, z) = ω(x, yz)ω(y, z), for all x, y, z ∈ G;
(M2) ω(x, e) = ω(e, x) = 1, for all x ∈ G.
Two multipliers ω1 and ω2 of G are similar if there exists a measurable function
ξ : G → T such that ω1 (x, y) = ξ(x)ξ(y)ξ(xy)−1 ω2 (x, y) for all x, y ∈ G. The

10

MAX CARTER

multiplier ω is normalized if ω(x, x−1 ) = 1 for all x ∈ G. Every multiplier is similar to a normalized multiplier. For a multiplier ω on G, define ω (2) (x, y) :=
ω(x, y)ω(y, x)−1 and Sω := {x ∈ G : ω (2) (x, y) = 1 ∀y ∈ G}. Note that Sω is a closed subgroup of G. The multiplier ω is called totally skew if Sω is trivial.
Fix a multiplier ω of G. Given a Hilbert space H, U(H) will denote the multiplicative group of unitary operators on H equipped with the strong operator topology. An ω-unitary representation of G, or ω-representation for short, is a pair
(π, Hπ ), where Hπ is a Hilbert space of arbitrary dimension and π : G → U(Hπ ) a function that satisfies:
(MR1) π(x)π(y) = ω(x, y)π(xy), for all x, y ∈ G;
(MR2) π : G → U(Hπ ) is measurable.
If ω is the trivial multiplier, then it turns out that π is continuous [40, Theorem
1], and in this case (π, Hπ ) is a unitary representation of G. The ω-representation
π is called irreducible if there does not exists any proper non-trivial closed invariant subspaces K ⊆ Hπ . Given two ω-representations (π, Hπ ) and (σ, Hσ ), a bounded operator T : B(Hπ ) → B(Hσ ) intertwines π and σ if T π(g) = σ(g)T for all g ∈
G. The vector space of all operators that intertwine π and σ will be denoted by Hom(π, σ). If π = σ, then Hom(π, σ) is a von Neumann algebra. The ωrepresentations (π, Hπ ) and (σ, Hσ ) are unitary equivalent if there exists a unitary operator U ∈ Hom(π, σ), in which case, we write π ≃ σ. The set of all irreducible bω (or
ω-representations of the group G up to unitary equivalence is denoted by G
b if ω is trivial). If P is a property of von Neumann algebras (e.g. type I, by G
type II, type III, factor etc.) then (π, Hπ ) is called a P ω-representation if the von Neumann algebra generated by π(G) is a P von Neumann algebra. If all the
ω-representations of G are type I, then we call ω a type I multiplier.
2.5. Group C∗ -algebras and the Fell topology. We now discuss group C∗ algebras and the Fell topology. It is assumed throughout this subsection that G is a locally compact group equipped with a fixed Haar measure µ.
Given a unitary representation (π, Hπ ) of G, we can extend π to Ra ∗-representation of L1 (G) as follows: for f ∈ L1 (G) and ξ ∈ Hπ , we define π(f )ξ := G f (g)π(g)ξ dµ(g), where we interpret the integral as a Bochner integral [22, Appendix 3]. This sets up a correspondence between unitary representations of G and non-degenerate ∗representations of L1 (G) (here, a ∗-representation π of L1 (G) is non-degenerate if the only vector ξ ∈ Hπ satisfying π(L1 (G))ξ = {0} is the trivial vector).
Theorem 2.12. [15, §13.3] Let G be a locally compact group. There is a one-toone correspondence between the unitary representations of G and non-degenerate ∗representations of L1 (G). This correspondence preserves irreducibility and unitary equivalence.
Define a norm k · kmax on L1 (G), which we call the maximal C∗ -norm, by b for f ∈ L1 (G). It can be checked that this kf kmax := sup{kπ(f )kop : π ∈ G},
1
is indeed a norm on L (G) and the completion of L1 (G) with respect to this norm is a C∗ -algebra. The completion is the group C∗ -algebra of G and is denoted by
C ∗ (G).

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

11

Since L1 (G) is dense in C ∗ (G) and every non-degenerate ∗-representation of
L (G) is continuous with respect to the maximal C∗ -norm, every non-degenerate
∗-representation of L1 (G) extends uniquely to a non-degenerate ∗-representation of C ∗ (G), and by the prior discussion, the unitary representations of G are thus in bijective correspondence with the non-degenerate ∗-representations of C ∗ (G).
Two ∗-representations (π, Hπ ) and (σ, Hσ ) of a C∗ -algebra A are called unitary equivalent if there exists a unitary operator U : Hπ → Hσ such that U π(x) = σ(x)U
for all x ∈ A, identical to the definition for groups. The collection of equivalence classes of irreducible ∗-representations of a C∗ -algebra A up to unitary equivalence b and is called the unitary dual of A.
forms a set, denoted A,
1

Definition 2.13. Let A be a C∗ -algebra and I an ideal of A. The ideal I is called primitive if there exists an irreducible ∗-representation (π, Hπ ) of A such that I = ker(π). The collection of all primitive ideals of A is called the primitive ideal space and is denoted by Prim(A).

b → Prim(A), π 7→ ker(π). The set
We denote by κ the canonical map κ : A
Prim(A) also has a topology defined on it as follows. First we define two operations hull and ker: if X ⊆ 2Prim(A) , then ker(X) is the intersection of all ideals in X, and if J ⊆ A is an ideal, hull(J) is the set of all primitive ideals in A containing J.
Definition 2.14. A set in X ⊆ Prim(A) is defined as closed if X = hull(ker(X)).
The topology generated by these closed sets is called the hull-kernel topology on
Prim(A).
Also, note that by taking the pullback of the hull-kernel topology on Prim(A)
b which is called the Fell topology.
along κ, we also get a topology on the dual space A,
The Fell topology can also be characterised in terms of matrix coefficients of irreducible unitary representations, which we now describe. Given a representation
(π, Hπ ) of a C∗ -algebra A, a positive linear functional associated with π is any positive linear functional on A of the form f (x) := hπ(x)ξ, ξi for ξ ∈ Hπ . These linear functionals are also called diagonal matrix coefficients of the representation
π. A linear functional is called a state if it is normalised i.e. has norm 1.
b and X ⊆ A.
b
Proposition 2.15. [15, Theorem 3.4.10] Let A be a C∗ -algebra, π ∈ A
The following are equivalent:
(i) π ∈ X;
(ii) At least one of the non-zero positive linear functionals associated with π is a weak∗ -limit of positive linear functionals associated with representations in
X;
(iii) Every state associated with π is a weak∗ -limit of states associated with representations in X.

We now describe the Fell topology on the group level and state some results about the Fell topology that will be used in this article. We start with the notion of a function of positive type on a group G. These are the analogues of positive linear functionals in the C∗ -algebra context.

12

MAX CARTER

Definition 2.16. Let G be a locally compact group. A function of positive type on G is a continuous function ϕ : G → C such that for all g1 , . . . , gn ∈ G and c1 , . . . , cn ∈ C
n X
n
X
ci cj ϕ(gi gj−1 ) ⩾ 0.
i=1 j=1

It can be checked that every diagonal matrix coefficient corresponding to a unitary representation on a locally compact group is a function of positive type on that group. We thus make the following definition.
Definition 2.17. Let G be a locally compact group and (π, Hπ ) a unitary representation of G. For ξ ∈ Hπ , define ϕπ,ξ : G → C, g 7→ hπ(g)ξ, ξi. The functions ϕπ,ξ
with ξ ranging over all ξ ∈ Hπ are called the functions of positive type associated with π.
It should be noted that, by a variation of the GNS construction for groups, one can show that every function of positive type on a locally compact group is a diagonal matrix coefficient of a cyclic unitary representation on that group [5,
Proposition 1.B.10].
b of a locally compact
We now define the Fell topology on the unitary dual G
b group G and then connect convergence in G with the notion of weak containment of representations.
b is
Definition 2.18. Let G be a locally compact group. The Fell topology on G
b the topology generated by the basis of open sets W (π, ϕ1 , . . . , ϕn , ǫ), where π ∈ G,
K ⊆ G compact, ϕ1 , . . . , ϕn functions of positive type associated to π and ǫ > 0.
The set W (π, K, ϕ1 , . . . , ϕn , ǫ) consists of precisely all of those representations σ ∈
b such that for each ϕi , there exists a function ψ which is a finite sum of functions
G
of positive type associated to σ such that
|ϕi (x) − ψ(x)| < ǫ

for all x ∈ K.
b is endowed with this topology, the bijection
It can be shown that when G
\
∗
∗ (G)
b
G → C (G) is a homeomorphism [5, Corollary 8.B.5]. The Fell topology on C\
can also be defined in terms of the open sets W (π, K, ϕ1 , . . . , ϕn , ǫ) and the corresponding matrix coefficients on the C∗ -algebra level, instead of the definition given earlier in this subsection.
We now briefly discuss weak containment of unitary representations which can be used to characterise convergence in the Fell topology. The notion of weak containment of representations is a weakening of the notion of containment of unitary representations.
Definition 2.19. Let G be a locally compact group and (π, Hπ ) and (σ, Hσ ) two unitary representations of G. We say the π is weakly contained in σ, denoted
π  σ, if for every ǫ > 0, every compact K ⊆ G and every ξ ∈ Hπ , there exists

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

13

ζ1 , . . . , ζn ∈ Hσ such that
|hπ(x)ξ, ξi −

n
X

hσ(x)ζi , ζi i| < ǫ

i=1

for all x ∈ K.

It is an important fact that convergence in the Fell topolgy can be characterised in terms of weak containment.
Proposition 2.20. [5, Proposition 1.C.8] Let G be a locally compact group and b Then the net (πi )i∈I converges to π ∈ G
b if and only if π 
(πi )i∈I a net in G.
L
π
for every subnet
(π
)
of
(π
)
.
j j∈J
i i∈I
j∈J j

Weak containment of representations can also be characterised in terms of kernels of representations on the C∗ -algebra level. Given a unitary representation π of a locally compact group G, then recall that π extends to a ∗-representation of C ∗ (G), also denoted by π. Let C ∗ ker(π) denote the kernel of π when considered as a b → Prim(C ∗ (G)), π 7→
representation of C ∗ (G). It is clear that there is a map G
∗
C ker(π). Weak containment can be characterised as follows.

Proposition 2.21. [5, Proposition 8.B.4] Let G be a locally compact group and
(π, Hπ ) and (σ, Hσ ) two unitary representations of G. Then π  σ if and only if
C ∗ ker(σ) ⊆ C ∗ ker(π).
To finish this subsection on the Fell topology, we state a result about the Fell topology for CCR groups that will be used later. Recall the definition of a CCR
C∗ -algebra in [15, §4.2] or [22, pg. 270].
Definition 2.22. Let G be a locally compact group. The group G is called a CCR
group, or CCR for short, if C ∗ (G) is CCR.

Since every CCR C∗ -algebra is type I, it follows that every CCR group is type
I. We state the following result which provides equivalent characterisations of the
CCR property for groups. This is again a consequence of a result of Glimm: see
[25] and [15, §4.2].
Proposition 2.23. Let G be a locally compact second countable group. The following are equivalent:
(i) The group G is CCR;
(ii) For every irreducible unitary representation π of G and every f ∈ L1 (G),
π(f ) is a compact operator;
∗ (G)) is T .
b (equiv. C\
(iii) The Fell topology on G
1

2.6. Induced representations. Here we give a quick summary of the definition of an induced representation, in-part to set some notation concerning induced representations, and then we prove a lemma that will be required later in the article.
The reader should refer to the books [44, 36, 23] for further information and other constructions of induced representations.

14

MAX CARTER

Let G be an arbitrary locally compact group and H a closed subgroup of G. In the following, q : G → G/H will denote the quotient map and supp(f ) will denote the support of a given function f . The modular function on G is the homomorphism
∆G : Aut(G) → R>0 in which ∆G (α) (α ∈ Aut(G)) is defined to be the real number such that ∆G (α)µG (B) = µG (α−1 (B)) for all Borel sets B ⊆ G (see [31, §15.26]
for more details on the function ∆G ). If α is the inner-automorphism defined by an element g ∈ G, then we write ∆G (g) instead of ∆G (α). The modular function on H is similar defined and denoted by ∆H .
Let (π, Hπ ) be a unitary representation of the subgroup H. Define Cc (G, π) to be the space of all functions ξ : G → Hπ satisfying the following properties:
(i) ξ is continuous with respect to the norm topology on Hπ ;
(ii) q(supp(ξ)) is compact in G/H;
(iii) For all x ∈ G and h ∈ H, ξ(xh) = ∆H (h)1/2 ∆G (h)−1/2 π(h−1 )ξ(x).
We will now discuss how to define an inner-product on Cc (G, π). The completion of Cc (G, π) with respect to the inner-product that we will define will be the representation space of the induced representation of π from H to G.
For a fixed ϕ ∈ Cc (G), define a function on G by
Z
ϕ(xh) dµH (h), (x ∈ G),
ϕH (x) =
H

where µH denotes the Haar measure on the subgroup H. Since ϕ is uniformly continuous, the function ϕH is continuous on G. Also ϕ(xh) = ϕ(x) for all x ∈
G and h ∈ H since µH is left invariant. It follows that there is a well defined continuous function ϕG/H ∈ Cc (G/H) given by ϕG/H (xH) = ϕH (x). The next proposition is important in what follows.
Proposition 2.24. [36, Propositon 1.9] For every ψ ∈ Cc (G/H), there exists a
ϕ ∈ Cc (G) such that ϕG/H = ψ.

Now, for ξ ∈ Cc (G, π), define a function λξ : Cc (G/H) → C as follows: given
ψ ∈ Cc (G/H), choose ϕ ∈ Cc (G) such that ϕG/H = ψ. Then define
Z
2
ϕ(x) kξ(x)k dµG (x)
λξ (ψ) :=
G

where dµG denotes the Haar measure on G. It can be shown that λξ is a well-defined positive linear functional on Cc (G/H). Consequently, there exists a positive Radon measure µξ on G/H such that
Z
Z
2
ϕ(x) kξ(x)k dµG (x) =
ϕG/H (xH) dµξ (xH).
G

G/H

Then, for any ξ, ζ ∈ Cc (G, π), one can define a complex Radon measure on G/H
by

1
(µξ+ζ − µξ−ζ + iµξ+iζ − iµξ−iζ ).
4
It can be checked that the map Cc (G, π) × Cc (G, π) → C, (ξ, ζ) 7→ µξ,ζ (G/H) is a non-degenerate bilinear form on Cc (G, π) which makes Cc (G, π) into a pre-Hilbert space. The Hilbert space completion of this space is denoted by HindG
.
Hπ
µξ,ζ :=

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

15

The induced representation of π from H to G is denoted by indG
H π and acts by left translation on functions in Cc (G, π) i.e. for x, y ∈ G and ξ ∈ Cc (G, π),
−1
indG
y). It can be checked that this representation acts continuH π(x)ξ(y) = ξ(x ously on Cc (G, π) and extends to a unitary representations on HindG
. See [36,
H (π)
§2.3] for more information on the above construction.
We now prove the lemma that will be used later in the article. This result is probably already known to experts in the area, but we could not find a proof of this result in the literature in the generality given here.
Lemma 2.25. Let G be a locally compact group, H a closed subgroup of G and π
G
a unitary representation of H. For any α ∈ Aut(G), α · indG
H π ≃ indα−1 (H) α · π,
−1
where α · π denotes the unitary representation of α (H) defined by α · π := π ◦ α.
Proof. We construct an isometric isomorphism T : HindG−1
α

(H)

which
α·π → HindG
Hπ

G
intertwines α·indG
H π and indα−1 (H) α·π. To do this, define a map T : Cc (G, α·π) →
Cc (G, π) by T ξ(x) = ∆G (α)1/2 ξ(α−1 (x)) (ξ ∈ Cc (G, α · π), x ∈ G). Clearly T is a

linear mapping.
The proof now proceeds in two steps. First, we show that T is an isometric isomorphism of the pre-Hilbert space Cc (G, α · π) onto Cc (G, π) and hence extends
. Then, we show that T
to an isometric isomorphism T : HindG−1 α·π → HindG
Hπ
α

(H)

G
intertwines α · indG
H π and indα−1 (H) α · π.
We may assume that the Haar measures on H and α−1 (H) are related by the equation
µH (B) = µα−1 (H) (α−1 (B))
for any Borel set B ⊆ H. Also, we note that ∆H (h) = ∆α−1 (H) (α−1 (h)) for any h ∈ H.
Let ξ ∈ Cc (G, α · π). Then, since ξ is compactly supported modulo α−1 (H), there exists a compact subset K ⊆ G such that supp(ξ) ⊆ Kα−1 (H). It then follows that supp(T ξ) ⊆ α(K)H, and since α(K) is compact by continuity of α, it is thus the case that T ξ is compactly supported modulo H.
For x ∈ G and h ∈ H, we also have that

T ξ(xh) = ∆G (α)1/2 ξ(α−1 (xh)) = ∆G (α)1/2 ξ(α−1 (x)α−1 (h))
= ∆G (α)1/2 ∆α−1 (H) (α−1 (h))1/2 ∆G (α−1 (h))−1/2 α · π(α−1 (h−1 ))ξ(α−1 (x))
= ∆H (h)1/2 ∆G (h)−1/2 ∆G (α)1/2 π(h−1 )ξ(α−1 (x))
= ∆H (h)1/2 ∆G (h)−1/2 π(h−1 )T ξ(x).
It is clear that T ξ is continuous since α and ξ are continuous. It thus follows by the above arguments that T ξ ∈ Cc (G, π) for any ξ ∈ Cc (G, α · π).
To see that T is surjective, suppose that η ∈ Cc (G, π). Then, define a function
ξ on G by ξ(x) := ∆G (α)−1/2 η(α(x)) for x ∈ G. By replacing α with α−1 in the above arguments, one checks that ξ ∈ Cc (G, α · π), and it is clear that T ξ = η, which proves surjectivity of T .

16

MAX CARTER

We now show that T is norm preserving. Let ξ ∈ Cc (G, α·π) and let K ⊆ G be a compact set such that supp(ξ) is contained in Kα−1 (H). Note that, as seen above, this also implies that supp(TRξ) is contained in α(K)H. Then, choose ψ ∈ Cc (G)
such that for all x ∈ α(K), H ψ(xt) dµH (t) = 1. We then have, by definition of the inner-product on HindG
,
Hπ
Z
2
2
ψ(x) kT ξ(x)k dµG (x)
kT ξk =
ZG
2
ψ(x) ∆G (α)1/2 ξ(α−1 (x)) dµG (x)
=
ZG
2
=
ψ(x) ξ(α−1 (x)) ∆G (α) dµG (x)
ZG
2
=
ψ(α(x)) kξ(x)k dµG (x).
G

Note that theR function α · ψ defined by α · ψ(x) := ψ(α(x)) (x ∈ G) is in Cc (G)
and satisfies α−1 (H) α · ψ(xt) dµα−1 (H) (t) = 1 for all x ∈ K. It then follows by the definition of the inner-product on HindG−1 α·π that
α
(H)
Z
2
2
ψ(α(x)) kξ(x)k dµG (x) = kξk
G

which implies kT ξk = kξk for all ξ ∈ Cc (G, α · π). Since Cc (G, α · π) is dense in HindG−1 α·π and Cc (G, π) is dense in HindG
, it thus follows from the prior
Hπ
α

(H)

arguments that T extends to an isometric isomorphism T : HindG−1
α

(H)

.
α·π → HindG
Hπ

G
We now show that T intertwines α·indG
H π and indα−1 (H) α·π. For ξ ∈ Cc (G, α·π)
and x, y ∈ G, we have
G
(((α · indG
H π)(x))(T ξ))(y) = ((indH π(α(x))(T ξ))(y)

= T ξ(α(x)−1 y)
= T ξ(α(x−1 )y)
= ∆(α)1/2 ξ(α−1 (α(x−1 )y))
= ∆(α)1/2 ξ(x−1 α−1 (y))
−1
= ∆(α)1/2 ((indG
(y)))
α−1 (H) α · π)(x))(ξ(α

= T (((indG
α−1 (H) α · π)(x))ξ)(y).
Then, again, since Cc (G, α · π) is dense in HindG−1
α

(H)

α·π , Cc (G, π) is dense in

HindG
, and T is an isometric isomorphism, it follows from the above calculation
Hπ
G
G
that T intertwines α · indG
H π and indα−1 (H) α · π. Thus we have that α · indH π ≃
indG

α−1 (H) α · π.

ON THE UNITARY REPRESENTATION THEORY OF CONTRACTION GROUPS

17

2.7. Mackey little group method. The main mathematical tool used in our proofs is the Mackey little group method and its generalisations to crossed product
C∗ -algebras. In this subsection we collect the information that we will need regarding the Mackey little group method on the group level and in the next subsection we collect the information concerning the Mackey little group method in the context of crossed-product C∗ -algebras. We refer the reader to [53] for a historical overview of the Mackey little group method with details concerning its relation to
C∗ -algebra theory.
A Borel space X is called countably separated if there exists a countable collection
{Ui }∞
i=1 of Borel subsets of X such that for any x ∈ X, the intersection of all the
Ui containing x is precisely the set {x}.
b as
Whenever we have a group G and a normal subgroup N of G, G acts on N
follows. An element g ∈ G acts on a unitary representation π of N by g · π(n) =
b by g · [π] = [g · π]
π(gng −1 ). This action then gives rise to an action of G on N
for π an irreducible unitary representation of N and [π] the corresponding unitary equivalence class.
The subgroup of all elements in G that stabilise either a representation π of N
b are both denoted by Gπ . Note that Gπ always or an equivalence class [π] ∈ N
contains N .
Definition 2.26. Let G be a locally compact group and N a normal subgroup of b /G
G. The group N is said to be regularly embedded in G if the quotient space N
is a countably separated Borel space.
We give a simple condition on the groups G and N which implies that N is regularly embedded.

Proposition 2.27. [24, Theorem 1] Let G be a locally compact second countable group and N a type I closed normal subgroup of G. Then N is regularly embedded b are locally closed.
in G if and only if the orbits of the action G y N

For use throughout this subsection, we state the definition of quasi-equivalence of unitary representations.

Definition 2.28. Let (π, Hπ ) and (σ, Hσ ) be unitary representations of a locally compact group G. The representations (π, Hπ ) and (σ, Hσ ) are quasi-equivalent if every non-trivial subrepresentation π ′ of π is not disjoint from σ, and every non-trivial subrepresentation σ ′ of σ is not disjoint from π.
Remark 2.29. (i) Two unitary representations π1 and π2 are quasi-equivalent if and only if there exists a cardinal n such that nπ1 and nπ2 are unitary equivalent [5, Proposition 6.A.4].
(ii) Two quasi-equivalent unitary representations generate isomorphic von Neumann algebras [15, §5.3.1].
We now move onto describing the Mackey little group method. First we describe the Mackey little group method for ordinary unitary representations, and then later we describe it for when multiplier representations are necessary.

18

MAX CARTER

In Theorem 2.31, we use the term quasi-orbit of a representation, and the notion of a quasi-orbit being concentrated in an orbit. It is not necessary for the purpose of this article to understand these terms, however, one can consult their definitions in
[44, pg. 186]. The important fact that one must know is that if G is a locally compact separable group and N a closed normal type I regularly embedded subgroup of G, then the quasi-orbit of every factor representation of G is concentrated in a single orbit. Thus, in this case, every factor representation of G, and in particular, every irreducible representation of G, can be constructed from a representation of b ) via the method described in Theorem 2.31. We state one of the groups Gπ (π ∈ N
this fact as a proposition below due to its importance.

Proposition 2.30. [44, pg. 186] If N is a type I regularly embedded subgroup of G, then the quasi-orbit of every factor representation is concentrated in a single orbit b.
of the action G y N

Theorem 2.31. [44, Theorem 3.11] Let G be a separable locally compact group b
and N a type I closed normal subgroup of G. Fix an orbit O of the action G y N
′′
π
=:
σ
and let π be a representative of this orbit. Then the map π ′′ → indG
Gπ
establishes, up to unitary equivalence, a one-to-one correspondence between:
(a) the set of all factor representations π ′′ of Gπ such that π ′′ |N is quasi-equivalent to π; and,
(b) the set of all factor representations σ of G whose quasi-orbits are concentrated in O.
∼ Hom(σ, σ), in particular,
Furthermore, under this correspondence, Hom(π ′′ , π ′′ ) =
the irreducible unitary representations in (a) are in one-to-one correspondence with the irreducible unitary representations in (b).
Remark 2.32. Note that since Hom(π ′′ , π ′′ ) ∼
= Hom(σ, σ), it follows that the von
Neumann algebras generated by π ′′ and σ are isomor

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
