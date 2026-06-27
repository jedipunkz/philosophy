---
source: "https://arxiv.org/abs/alg-geom/9612003v2"
title: "A correspondence dual to McKay's"
author: "Jean-Luc Brylinski"
year: "1996"
publication: "arXiv preprint / math.AG"
download: "https://arxiv.org/pdf/alg-geom/9612003v2"
pdf: "https://arxiv.org/pdf/alg-geom/9612003v2"
captured_at: "2026-06-24T11:10:56Z"
updated_at: "2026-06-24T11:10:56Z"
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

# A correspondence dual to McKay's

- 著者: Jean-Luc Brylinski
- 年: 1996
- 掲載情報: arXiv preprint / math.AG
- 情報源: [arxiv](https://arxiv.org/abs/alg-geom/9612003v2)
- ダウンロード: https://arxiv.org/pdf/alg-geom/9612003v2
- PDF: https://arxiv.org/pdf/alg-geom/9612003v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We give a dual to the McKay correspondence, involving conjugacy classes of subgroups of SU(2). We prove a determinantal formula involving both correspondences. We pose some questions concerning a non-commutative Fourier transform.

## PDF Text

January 30, 1997
A correspondence dual to McKay’s
Jean-Luc Brylinski*

arXiv:alg-geom/9612003v2 7 Feb 1997

1. Introduction
It is well-known from the work of DuVal [DuVal1] and M. Artin [A]that there is a one-to-one correspondence between finite subgroups G of SU (2) and Coxeter-Dynkin diagrams ∆ of type A, D, E. This involves a minimal resolution of singularities X̃ of the singular algebraic surface C2 /G. Around 1980 McKay found a deep correspondence between vertices of the affine Coxeter-Dynkin diagram and irreducible representations of the group [McKay1] [Mckay2]. Several systematic representation-theoretic proofs were given by Kostant [Ko], Steinberg [St2], Springer [Sp]. A geometric interpretation of the correspondence was given by Gonzalez-Sprinberg and Verdier [GS-V] and also by Knörrer
[Kn].
There is also a dual correspondence, this time between vertices of ∆ and non-trivial conjugacy classes of G. This dual correspondence was introduced by Ito and Reid [I-R]
in the more general context of a finite subgroup of SU (n). The construction is in fact very simple topologically: we interpret G as the fundamental group of the complement of the exceptional divisor in X̃. Then each vertex of ∆ corresponds to a component of this divisor, and there is an associated class of a small loop encircling said component.
The dual correspondence was studied by Ito and Reid from the point of view of valuations on function fields. From the description of the fundamental group due to Mumford [Mu]
one sees that this gives a one-to-one correspondence between vertices of the diagram and non-trivial conjugacy classes. This result amounts to the dimension 2 case of a more general theorem proved by Ito and Reid [I-R]. The dual correspondence has a number of interesting further properties, which are detailed in Theorem 4.1. These properties involve the three (or two) so-called special conjugacy classes corresponding to ends of the diagram, and the description of the edges of the diagram involve pairs of commuting elements x, y such that y is special; then the conjugacy classes of x and xy are joined by an edge.
There seem to be intriguing connections between the McKay correspondence and the dual correspondence. We prove a determinantal formula concerning an element gj of G
associated to a vertex vj of ∆ and the irreducible representation Ek associated to a vertex vk : det(gj , Ek ) = exp(−2πi(C −1 )jk ), where C −1 is the inverse of the Cartan matrix. One tool we use in proving this formula is the geometric description of the McKay correspondence in [GS-V] by means of the first
Chern class of the vector bundle associated to a representation of G. We also use the properties of vector bundles with integrable connections admitting logarithmic poles, in particular the computation of the first Chern class from the residues of the connection.
* This research was supported in part by NSF grant DMS-9504522.
1

The paper ends with some remarks on the matrix-valued Fourier transform which results from comparing the two correspondences.
This work was written in May-June 1996 as I was visiting Harvard University. I
thank David Kazhdan and the Harvard Mathematics Department for their hospitality. I
am grateful to Hélène Esnault for correspondence and information about her work [E-V]
and for her useful remarks on a first draft of this paper. I thank Victor Batyrev for pointing out to me the paper [I-R]. I am grateful to Igor Dolgachev and to John McKay for reading a first version of this paper and making many valuable comments.

2. The McKay correspondence
There is a by now classical correspondence between conjugacy classes of finite subgroups of SU (2) (or equivalently, of SL(2, C)) and simply-laced Coxeter-Dynkin diagrams
(thus of type An , Dn or En ). The correspondence was constructed by DuVal [DuVal1] using algebraic geometry. It may be phrased as follows. Given a finite subgroup G ⊂ SU (2)
one can construct the singular algebraic surface X = C2 /G, quotient of the affine plane by the action of G. Let f : C2 → X be the quotient map; the point o = f (0) is called the origin of X. The singular locus of X is reduced to o, unless G = 1, in which case X = C2
is smooth. X is a normal affine surface, whose algebra of regular functions is the algebra
C[x, y]G of invariants in the polynomial algebra on two generators. There is a minimal resolution of singularities p : X̃ → X, so
- X̃ is a smooth algebraic surface;
- p is a proper regular mapping, which induces an isomorphism over the open subset
U = X \ {o} of X;
- the minimality of the resolution means that X̃ does not contain any rational curve of self-intersection −1.
Then the reduced fiber D = p−1 (o)red is a curve, which is a union of smooth rational curves D1 , · · · , Dr . Any two curves Di and Dj intersect transversally at at most one point of X̃. One can then construct a graph ∆ such that the vertices v1 , · · · , vr of ∆ correspond to the curves Dj , and where one joins the vertices vi and vj by an edge whenever the divisors Di and Dj intersect.
One associates to the system of curves (Dj ) a square matrix A = A(G) of size r, called the intersection matrix, such that Aij is the intersection number Di · Dj on the smooth surface X̃. On the other hand the graph ∆ has an incidence matrix M . Since any Di has self-intersection −2 we have A = −2Id + M .
We then have
Theorem 2.1. (Du Val [DuVal1], M. Artin [A]) (1) For any finite subgroup G of
SU (2), the graph Γ is a simply-laced Coxeter-Dynkin diagram.
(2) The Cartan matrix C of the Coxeter-Dynkin diagram ∆ is the opposite of the intersection matrix A.
(3) This construction induces a one-to-one correspondence between conjugacy classes of finite subgroups of SL(2, C) and simply-laced Coxeter-Dynkin diagrams.
2

We give the table showing the simply-laced diagrams and the corresponding finite subgroups of SU (2). For a regular polyhedron, we have the corresponding symmetry group H ⊂ SO(3). Its inverse image G in the double cover SU (2) is the corresponding binary polyhedral group.
Table of subgroups of SU (2)

∆
An
Dn , n ≥ 3
E6
E7
E8

FINITE SUBGROUPS OF SU (2)
order of G
G
n+1
cyclic
4n − 8
binary dihedral
24
binary tetrahedral
48
binary octahedral
120
binary icosahedral

The McKay correspondence on the other hand involves simply-laced affine CoxeterDynkin diagrams [McKay1] [McKay2]. Given a simply-laced Coxeter-Dynkin diagram
˜ which is obtained by adding one vertex v0
∆, there is a corresponding affine diagram ∆, to ∆. We need to explain for which i ∈ {1, · · · , r} the vertex v0 and vi are linked by an edge. This requires introducing the root system R corresponding to ∆. This a finite subset of an euclidean vector space E of dimension r, consisting of vectors of length 2. The vertices v1 , · · · , vr correspond to the simple roots α1 , · · · , · · · , αr , with respect to system
R+ of positive roots. The Cartan matrix is given by Cij = (αi , αj ). There is a longest root θ (so that θ is a positive root, and θ + αi is not a root for j = 1, · · · , r). Then the
˜ corresponds to α0 := −θ. There is an edge in ∆
˜ between the vertices new vertex v0 of ∆
v0 and vi if and only if (α0 , αi ) 6= 0.
˜
Pr Each vertex vi of ∆ is labeled by a positive
Pr integer mi in such a way that m0 = 1 and m
α
=
0.
Equivalently we have
θ
=
i i j=10
j=1 mj αj .
We can now state the McKay correspondence.
˜
Theorem 2.2. (Mckay [McKay]) Let G be a finite subgroup of SL(2, C) and let ∆
be the corresponding affine Coxeter-Dynkin diagram. (1) There is a one-to-one correspon˜ and equivalence classes of irreducible representations dence i 7→ Ri between vertices of ∆
of G. The dimension of Ri is equal to mi .
(2) Let E be the two-dimensional representation of G in C2 . Then for any i ∈
{0, · · · , r} we have
Ri ⊗ E →
˜ ⊕j incident to i Rj

(2)

This correspondence was constructed empirically by McKay. Coherent proofs were given in [St2] [Sp]. The representation-theoretic and invariant-theoretic aspects of the
3

correspondence were developed further in [Ko]. A geometric construction was given by
Gonzalez-Sprinberg and Verdier [GS-V]; this will be used in §5.

3. The special conjugacy classes.
We will use a well-known topological interpretation of the group G ⊂ SL(2, C).
Lemma 3.1. We have
G→π
˜ 1 (X \ {o}) = π1 (X̃ \ D)

(3 − 1)

We did not specify a base point in Lemma 3.1. This is because we only need the isomorphism G→π
˜ 1 (X̃ \ D) up to conjugation.
Proof. The space X \ {o} is the quotient of the simply-connected space C2 \ {0} by the action of G. Because G ⊂ SL(2, C), the action of G on C2 \ {0} is fixed point free.
Thus C2 \ {0} → X \ {o} is a Galois covering with group G.
For any component Di of D (1 ≤ i ≤ r), there is a well-defined conjugacy class in
π1 (X̃ \ D) which corresponds to a small oriented loop γi around the divisor Di . Of course the precise construction of this loop depends on the base point but the conjugacy class is well-defined. Using the isomorphism 3-1, this defines a conjugacy class in G, which will be denoted by Oi .
Definition 3.2. The dual McKay correspondence is the map
{vertices of ∆} → G/conj which maps vi to Oi .
This is a more topological version of the construction of Ito and Reid [I-R], which is phrased in terms of valuations and is purely algebraic i.e., invariant under automorphisms of the field C.
The main properties of the dual correspondence will be given in §4. These will involve some special conjugacy classes in G, which are indexed by the ends of the graph ∆. There are two ends for the graph An and three ends for the graph Dn and En .
For this purpose we consider the induced action of G on the projective line CP1 of lines in C2 . This action factors through an effective action of the image H of G in the quotient group P GL(2, C) = SL(2, C)/ ± 1. The quotient CP1 /H is isomorphic to the projective line, so we have a ramified covering π : CP1 → CP1 /H, which is a Galois covering of group
H. We note the well-known lemma
Lemma 3.3. (1) If ∆ is of type An for n odd, we have: G→H.
˜
4

(2) In every other case we have an exact sequence
1 → ±{1} → G → H → 1

(3 − 2)

There are three ramification points of π in CP1 /H, except in case An , when there are only two. We will give a bijection between the ramification set of π and the set of ends of
∆.
For each ramification point q ∈ CP1 /H pick a point q̃ ∈ π −1 (q) which corresponds to a line l ⊂ C2 . We have a natural mapping l ֒→ C2 → C2 /G. The inclusion l \ {0} ֒→ C2 \ {0}
gives a regular mapping l \ {0} → X \ {0} = X̃ \ D.
(3 − 3)
Because the map X̃ → X is proper, it follows that we can extend the mapping (3-3) to a regular mapping φ : l → X̃. The point φ(0) of X̃ is independent of the choice of q̃ ∈ π −1 (q).
Thus to q̃ we have attached the point x = φ(0) of X̃. Let Cq be the image of φ, which only depends on q, not on the choice of q̃. The map l → Cq is a ramified Galois covering with Galois group equal to the stabilizer Gl of l in G.
Proposition 3.4. (1) For a ramification point q ∈ CP1 /H, the corresponding point x of X̃ belongs to only one divisor Dj . The curve Cq ⊂ X̃ is a smooth curve which meets
Dj transversally.
(2) The vertex vj is an end of the graph ∆.
(3) The map q 7→ vj gives a bijection between the set S of ramification points of
π : CP1 → CP1 /H and the set of ends of ∆.
Proof. In case of a cyclic group G of order n, the statement is easy to prove using the explicit description of X̃ given in [Br] or in [GS-V]. In that case one of the lines l is the line x = 0. There is a covering of X̃ by n − 1 affine open sets, each of them isomorphic to the affine plane C2 . Consider the first open set U1 with coordinates (u, v). These can be chosen so that x = v n un−1 and y = u. Then the point φ(0) is the point u = 0, v = 0
which verifies the statement, as the divisor v = 0 is the divisor Dj corresponding to an end of the graph. A similar argument can be applied to the line y = 0. For the other cases one can use the results of Brieskorn [Br] to deduce them from the cyclic case. First we make a preliminary observation concerning the natural action of C∗ on C2 by dilations, which induces an algebraic C∗ -action on X and on X̃. The punctured lines l \ {0} ⊂ C2 are
C∗ -orbits. Their images Cl \ {φ(0)} in X are therefore also C∗ -orbits, and this describes all the 1-dimensional orbits. Among the orbits of dimension 1, those corresponding to ramification points of π are characterized by the fact that the action of C∗ is not free (the m-th roots of unity act trivially, if m is the order of Gl ). Now let Y be the blow-up of o ∈ X. According to [Br] Y has only isolated singularities (as many as the ends of ∆)
which are rational surface singularities of type An . The isomorphism between the germ of Y at such a point q and the germ of C2 /µn+1 at o can be made C∗ -equivariant. Now for the line l corresponding to ramification point, the corresponding limiting point in Y is a fixed point of C∗ , so it is one of the singular. Furthermore, the corresponding germ of
5

orbit in the singular surface C2 /µn+1 is a special orbit on which C∗ does not act freely.
This special orbit itself corresponding to a line x = 0 or y = 0. Now the resolution of singularities X̃ is obtained from Y by minimally resolving each singular point. The effect of this operation is already understood.
Now for any ramification point q ∈ S ⊂ CP1 /H there is a well-defined conjugacy class
Vi in H, which is defined as follows. There is a group homomorphism f : π1 ([CP1 /H]\S) →
H, which is only well-defined up to conjugacy. Take a small loop in [CP1 /H] \ S encircling the point q, and let hi be its image in G. Then Vi is the conjugacy class of hi . We can now state
Lemma 3.5. The conjugacy class Vi ⊂ H is the image of the conjugacy class Oi ⊂ G
under the canonical homomorphism G → H.

Proof. Clearly a representative hi of Vi is the image in H of the generator of the
2πi stabilizer Gl of l which admits e s as an eigenvalue, where s is the order of Gl . This is the same as the image under the group homomorphism
π1 ([l \ {0}]/Gl ) → π1 ([C2 \ {0}]/G)→G
˜ →H
of a small loop in [l \ {0}]/Gl which turns once around the point 0. On the other hand the curve Cq = φ(l) ⊂ X̃ is the closure of [l \ {0}]/Gl ⊂ X̃ \ D. From Proposition 3.3 we see that the conjugacy class Oi ⊂ G is represented by a small loop inside this curve which turns once around the point x = φ(0).
In case there are three ramification points q1 , q2 , q3 we can choose representatives h1 , h2 , h3 of the three corresponding conjugacy classes in H such that h1 h2 h3 = 1 (indeed this relation holds among the conjugacy classes in π1 ([CP1 /H] \ S) corresponding to the three punctures). it is natural to ask what relation exists among the conjugacy classes in
G.
Lemma 3.6. (1) In cases Dn and En , the conjugacy classes C1 , C2 , C3 corresponding to the ends v1 , v2 , v3 of the graph ∆ have representatives g1 , g2 , g3 which satisfy g1 g2 g3 = −1

(3 − 4).

(2) In the case An the conjugacy classes g1 and g2 corresponding to the two ends of
∆ satisfy g1 g2 = 1
(3 − 5).

Proof. This is easily checked using the explicit description of the group G given for instance in [Coxeter] or in [DuVal2].
6

We state the following result only in the case of a graph with three ends (the case of two ends is simpler and is left to the reader).
Proposition 3.7. (Coxeter) (1) The group H is isomorphic to the abstract group with generators h1 ,h2 , h3 and defining relations m

hj j = 1, h1 h2 h3 = 1

(3 − 6)

where mj is the order of hj in H, which is also equal to the length from vj to the central vertex vcen .
(2) The group G is isomorphic to the abstract group with generators g1 ,g2 , g3 and defining relations
(3 − 7).
g1m1 = g2m2 = g3m3 = g1 g2 g3 , (g1 g2 g3 )2 = 1
Of course the central element c = g1 g2 g3 is of order 2 and corresponds to −1 ∈
SL(2, C).

4. The dual McKay correspondence.
The dual McKay correspondence was introduced in Definition 3.2. It associates to a vertex vi of ∆ a conjugacy class Oi in G, which is defined topologically as the class of a small loop in X̃ \ D encircling the divisor Di .
The main properties of this correspondence are given in Theorem 1 below. Except in case An there is a central vertex vcen of ∆, and there are three branches.
We will use the notion of canonical numbering of the vertices of the tree ∆. This means that the vertices are numbered 1, · · · , r and that for any 2 ≤ k ≤ r the corresponding vertices v1 , · · · vk are the vertices of a subtree. Any canonical numbering gives an ordering of the set of vertices. Such an ordering will be called canonical. This notion was used by vonRandow [vR].
Theorem 4.1. (1) The correspondence vi 7→ Ci gives a bijection between the set of vertices of ∆ and the set of non-trivial conjugacy classes of G.
(2) The ends of ∆ correspond to the special conjugacy classes of G.
(3) (cases Dn , En ) The conjugacy class corresponding to the central vertex vcen consists of the central element −1.
(4) A branch v1 , · · · , vm of ∆ corresponds to a “geometric progression”
g, g 2, · · · , g m .
(5) Two vertices vi and vj belong to the same branch if and only if the corresponding conjugacy classes Ci , Cj have representatives gi and gj which commute.
(6) Two vertices vi and vj are joined by an edge if and only there exists a representative gi of Ci and a representative u of some special conjugacy class such that u commutes with gi and such that ugi belongs to Cj .
7

(7) Pick a canonical ordering of the vertices of ∆. Let vi be any vertex of ∆, and let vj , vk , · · · be the ordered set of neighbors of vi . Then there are representatives gi , gj , · · · of the corresponding conjugacy classes such that gi2 = gj gk · · · .

(4 − 1)

(8) For any canonical ordering of the vertices of ∆, once can choose an element gi of each conjugacy class Ci such that (4-1) holds for any vertex, and such that gi and gj commute whenever the vertices vi and vj are joined by an edge.
(4-2)
Then G is described as an abstract group as the group generated by these elements gi subject to these two types of relations.

Proof. We have again G = π1 (X̃ \ D). The fundamental group π1 (X̃ \ D) has been described by Mumford [Mu] in terms of precisely chosen loops around Di with class gi ∈ G.
Mumford proved statement (8). Now by Proposition 3.6 an end vj of ∆ corresponds to the conjugacy class of some gj whose order is exactly the length mj of the branch which ends at vj . Index the vertices on the branch by 1, · · · , mj . Then applying (4-1) inductively, we see that the vertex labeled by k, 1 ≤ k ≤ mj corresponds to the conjugacy class of gjk . In case An−1 , G = µn , with the vertices labeled linearly by {1, · · · , n − 1}, we see that the
2πik vertex labeled by k corresponds to e n , and all statements of the theorem are clear. So we now assume that we are in the case Dn or En , which means there is a central vertex vcen . Now for k = mj , we have the other end of the branch, which is the central vertex; m
the corresponding conjugacy class is represented by gj j = −1 (cf. Proposition 3.6 (2)).
At this point we can describe graphically the dual correspondence.
y x

x

2

···

c

(Dn )
yx

where G is generated by x and y with defining relations: xn−1 = y n−1 = c, c2 = 1, yxy −1 = x−1

x

x2

c

y2

(4 − 3)

y
(E6 )

z where G is generated by x, y, z with defining relations x3 = y 3 = z 2 = c, c2 = 1, xyz = −1
8

(4 − 4)

x

x2

x3

y2

c

y
(E7 )

z where G is generated by x, y, z with defining relations x4 = y 3 = z 2 = c, c2 = 1, xyz = −1

x

x2

x3

x4

c

y2

(4 − 5)

y
(E8 )

z where G is generated by x, y, z with defining relations x5 = y 3 = z 2 = c, c2 = 1, xyz = −1

(4 − 6)

The remaining statements can then be checked directly. Since the number of conjugacy classes of G is equal to the number of vertices of ∆, (1) will follow if we show that distinct vertices vi , vj correspond to distinct conjugacy classes. This is easy to see if vi and vj are on the same branch. Then the trace of g i in the two-dimensional representation C2 of G is
πj
πi equal to 2cos( πi m ) and for i 6= j, 1 ≤ i, j ≤ m we have 2cos( m ) 6= 2cos( m ). Then there are some cases to be considered where vertices on different branches correspond to conjugacy classes of elements of the same order. In case Dn it is easy to see that y and yx are not conjugate. In case E6 one checks that x2 and y 2 are not conjugate (this is related to the fact that there are two distinct conjugacy classes of rotations of order 3 in the symmetry group of the tetrahedron). This implies that x and y are not conjugate. In case E7 one checks that x2 and y are not conjugate, as their images in the symmetry group of the cube are not conjugate: the first is a flip whose axis goes through the center of two faces, the second one a flip whose axis goes through the middle of two edges. Statement (5) is easy to check, and then (6) follows directly.
Part (1) of the theorem is due to Ito and Reid [I-R, Theorem 1.4]. In fact, for an arbitary finite subgroup G of SU (n), they establish a bijection between non-trivial conjugacy classes in G and so-called crepant discrete valuations of the quotient variety
Cn /G.
Statements (3) and (4) were observed by Steinberg [St2]. We note that there is an a priori proof by vonRandow [vR] that the group defined by generators and relations in statement (8) of Theorem 4.1 is indeed independent of the canonical ordering of the vertices of ∆. Equation (4-1) may be viewed as saying that the assignment vi 7→ gi is a
“non-commutative harmonic map”.

5. Relation with the McKay correspondence
9

Gonzalez-Sprinberg and Verdier [GS-V] gave a geometric construction of the McKay corerspondence in terms of the first Chern class of some vector bundles on X̃. Given any representation of G in a a finite-dimensional vector space E, there is a natural algebraic vector bundle E over X̃ \ D = X \ {o}. In terms of locally free sheaves, the locally free sheaf f∗ OC2 \{0} has an action of G hence it admits a decomposition into irreducible representations of G: f∗ OC2 \{0} = ⊕E∈Irr(G) E ′ ⊗ FE

(5 − 1)

where each FE is a locally free sheaf. Then E is the algebraic vector bundle corresponding to FE .
We then want to extend the vector bundle E to X̃. For this it is enough to extend the locally free sheaf FE to X̃. Giving such an extension for any E ∈ Irr(G) amounts to giving an extension of f∗ OC2 \{0} to a locally free sheaf over X̃. Let j : X̃ \ D ֒→ X̃
be the inclusion. The extension given in [GS-V] and [Kn] is the sheaf A of subalgebras of j∗ f∗ OC2 \{0} generated by f∗ OC2 and by OX̃ . It is proved in [GS-V] [Kn] that this is actually locally free. The McKay correspondence is essentially given by the first Chern class of the vector bundle E. We have the following
Proposition 5.1. (1) The group H2 (X̃, Z) is the free abelian group of rank r generated by the homology classes [Di ] of the divisors Di , for 1 ≤ i ≤ r.
(2) The Picard group P ic(X̃) is isomorphic to H 2 (X̃, Z), hence to the Z-dual of
H2 (X̃, Z) = Zr . The isomorphism P ic(X̃) → Zr maps a line bundle L to the vector
(deg(L/Di )).
We identify both H2 (X̃, Z) and H 2 (X̃, Z) with Zr . Let (e1 , · · · , er ) denote the standard basis of H 2 (X̃, Z) = Zr .
Let Λ be the image of the natural injection Since H2 (X̃, Z) identifies by Poincaré
duality with the cohomology group Hc2 (X̃, Z) with compact supports, there is a natural map κ : H2 (X̃, Z) ֒→ H 2 (X̃, Z).
Then we have
Lemma 5.2. (1) The matrix of κ is the opposite of the Cartan matrix C.
(2) Λ is a lattice in H 2 (X̃, Z). Its index is equal to the connection index of the root system.
For each divisor Di the class κ[Di ] ∈ H 2 (X̃, Z) will again be denoted by [Di ]. It is the first Chern class of the locally free sheaf O(Di ). It follows from Lemma 5.2 that the classes [Di ] form a basis of the Q-vector space H 2 (X, Q). and that we have ei = −

X

(C −1 )ij [Dj ]

j

The theorem of [GS-V] can be stated as follows:
10

(5 − 2)

Theorem 5.3. (Gonzalez-Sprinberg [GS-V], see also [Kn]) Let Ei be the irreducible representation of G corresponding to the vertex vi of ∆. Then the first Chern class of the vector bundle Ei is equal to the standard vector ei of Zr .
Theorem 5.3 indeed gives a geometric construction of the McKay correspondence. The proofs of the theorem in [GS-V] and [Kn] involve a case by case computation, but there is a uniform proof in [A-V], which furthermore applies in arbitrary characteristic.
Now we will relate the vector bundles E to vector bundles with integrable connection.
We recall the notion of a Deligne vector bundle with meromorphic integrable connection over an algebraic manifold Z with respect to divisor Y ⊂ Z with normal crossings. Let
V be a an algebraic vector bundle over Z. Assume we have an integrable meromorphic connection ∇ on Z, which is holomorphic over Z \ Y . Then we say that (V, ∇) is a Deligne vector bundle with connection if
(1) for any germ of holomorphic section s of V , ∇(s) is a holomorphic section of
Ω1Z (log Y ) ⊗ V , where Ω1Z (log Y ) is the sheaf of 1-forms with logarithmic poles (so ∇ has at worst logarithmic poles along Y );
(2) for any component Yj of Y , and any eigenvalue α of the residue of ∇ along Yj , we have:
0 ≤ Re(α) < 1
(5 − 3)
The first Chern class of the vector bundle V is computable in terms of the residues of the connection along the components Yj of Y . For each j, we have the cohomology class
[Yj ] ∈ H 2 (Z, C). The residue ResYj (∇) is a complex number. Then we have:
Proposition 5.4. (Esnault-Verdier, see Appendix B of [E-V]) Assume the vector bundle V with integrable meromorphic connection ∇ satisfies (1) above. Then we have: c1 (V ) = −

X

T r ResYj (∇) [Yj ] ∈ H 2 (Z, C)

(5 − 4)

j

In fact, the result is proved in [E-V] for a proper algebraic variety. However, consider an algebraic vector bundle V over Z with integrable meromorphic connection satisfying
(1). There exists a smooth compactification X̄ of X such that (X̄ \ X) ∪ Ȳ is a divisor with normal crossings. Then it follows from the theory of [De] that V can be extended to a vector bundle over X̄ satisfying (1) with respect to the divisor (X̄ \ X) ∪ Ȳ . The equality
(5-4) for V̄ implies the corresponding equality for V .
Here is an important class of examples of Deligne vector bundles with integrable connection.
Lemma 5.5. Let Y be a divisor with normal crossings in the smooth complex algebraic variety Z. Let S be a normal algebraic variety and let h : S → Z be a proper morphism such that
(1) h is an étale morphism over Z \ Y .
(2) h∗ OS is a locally free sheaf over Z.
11

Then the vector bundle associated to h∗ OS has the natural structure of a Deligne line bundle with integrable connection.

Proof. Over Z \ Y we have a unique connection ∇ on h∗ OS which is compatible with the algebra structure. For any section u of h∗ OS over an open subset of Z \ Y , we can find
Pn−1
a polynomial equation P (u) = 0, where P (x) = xn + i=0 ai xi is a monic polynomial with coefficients in OZ such that P ′ (u) is nowhere vanishing. Then we have:
1

∇(u) = α ⊗

P ′ (u)

,

P
where α = i dai xi . To prove properties (1)-(2) we work with holomorphic sheaves. Now near a point of Y where Y has local equation x1 · · · xl = 0, the locally free sheaf h∗ OS is a
1
1
direct factor of the sheaf of algebras OZ [x1m , · · · , xlm ] for some m. Indeed it is obtained as the subsheaf of invariants under some subgroup of (µm )l . It is therefore enough to treat
1

1

the case of the sheaf of algebras OZ [x1m , · · · , xlm ]. This has a basis consisting of functions q1

ql

of the type u = x1m · · · xlm for 0 ≤ qj ≤ m − 1. Then we have
∇u =

X qj dxj j

m xj

⊗ u,

which makes (1) and (2) apparent.
This however does not directly apply to our bundle of algebras over X̃, because the corresponding sheaf of algebras A is not integrally closed. We need to introduce the integral closure Ã. This satisfies all the assumptions of proposition 5.3, at least on X̃ \ T , where
T ⊂ X̃ is a finite set. Since we are interested in the first Chern class, we may neglect the effect of deleting this finite set. We then have an exact sequence of G-equivariant coherent sheaves on X̃:
0 → A → Ã → F → 0
(5 − 5)
for some sheaf F supported on Y . For the isotypic components associated to E ∈ Irr(G)
this yields an exact sequence
0 → E → Ẽ → H → 0
(5 − 6)
for some coherent sheaf H supported on Y . For the first Chern classes we obtain c1 (E) = c1 (Ẽ) −

r
X
j=1

where mj ≥ 0.
We can state
12

mj [Dj ]

(5 − 7)

Proposition 5.6. Let E be the vector bundle over X̃ associated to an irreducible representation E of G. We have c1 (E) = −

r
X

qj [Dj ] with qj ≥ 0

(5 − 8)

j=1

Proof. The vector bundle Ẽ has an integrable connection, and is a direct factor of the vector bundle with integrable connection A. We have from Proposition 5.4 the equality
X
c1 (Ẽ) = −
λj [Dj ], j

where λj is the trace of the residue of the integrable connection along Dj . By Lemma 5.5
all the eigenvalues of the residue have real part ≥ 0, so that Re(λj ) ≥ 0. Thus we get the equality
X
c1 (Ẽ) = −
(λj + qj )[Dj ], j

with Re(λj + qj ) ≥ 0. Since the [Dj ] are linearly independent, it follows that the λj + qj are positive rational numbers.
On the other hand we know from Theorem 5.3 that the qj are (up to sign) equal to the coefficients of the inverse of the Cartan matrix. Therefore we conclude
Corollary 5.7. The coefficients (C −1 )ij of the inverse C −1 of the Cartan matrix are all ≥ 0.
This can of course be easily read off the tables in Bourbaki [Bo]. Indeed (C −1 )jk is the coefficient of αj in the fundamental weight ωk . This was pointed out to me by Dolgachev.
We note the graph-theoretic interpretation of C −1 given by Lusztig and Tits [L-T].
It is actually easy to prove directly that all coefficients of C −1 are positive. Recall that
C = 2I − A where A is the matrix of ∆. Now the operator norm of A is the norm of the largest eigenvalue, which is well-known to be smaller than 2. Thus we have a convergent series:
1 X −n n
2 (A )ij
(5 − 9)
C −1 = (2I − A)−1 =
2
n≥0

The coefficient (An )ij is the number of paths of length n from vi to vj in ∆ . So it is always
≥ 0 and given (i, j) there exists n such that (An )ij > 0. In fact we have an estimate:
(C −1 )ij ≥

21−n
3

(5 − 10)

where n is the distance between the vertices vi and vj . This estimate is sharp for the A2
case.
13

Our final result involves both the McKay correspondence and the dual correspondence.
Let gj be a representative of the conjugacy class associated to the divisor Dj . Then we have:
Proposition 5.8. For the irreducible representation Ek of G associated to the vertex vk of ∆ we have det(gj , Ek ) = exp(−2πi(C −1 )jk )
(5 − 11)

Proof. The conjugacy class of the monodromy operator gj on Ek is represented by the operator exp(−2πiResDj ∇), where ∇ is the meromorphic connection on the Deligne vector bundle Ẽk ; this is a general property of the residue in the case of semisimple monodromy, proved in [De]. So we have: det(gj , Ek ) = exp(−2πiT r ResDj ∇).
Now by Proposition 5.4 the trace of the residue is the opposite of the coefficient of c1 (Lj )
in c1 (Ẽk ). This coefficient is congruent modulo Z to the similar coefficient for c1 (Ek ). The latter coefficient is the opposite of the coefficient (C −1 )jk of C −1 .
Recall that the connection index of ∆ is equal to the determinant of the Cartan matrix.
Corollary 5.9. The exponent of the abelianization Gab of G is equal to the connection index of ∆.
Proof. Indeed if n is the connection index, then all coefficients (C −1 )ij belong to n1 Z.
It then follows from (5-11) that any character G → C∗ takes values in the n-th roots of unity. So the exponent m of Gab divides n. Now let n = mq; i it follows from (5-11) that for each coefficient (C −1 )ij we have m(C −1 )ij ∈ Z. But it is easy to see that the g.c.d of the integers n(C −1 )ij is equal to 1. This implies q = 1 and m = n.
This result in fact follows easily from the presentation of the group G in terms of the
Cartan matrix given in [H-N-K].
For instance, the connection index of the diagram E8 is equal to 1, which implies that the binary icosahedral group is perfect (a well-known fact, of course).
This gives evidence for the idea that the matrix-valued Fourier transform obtained by combining the two types of correspondences should be very significant geometrically. One
−jk modulo Z. Therefore we get easily checks in the An case that (C −1 )jk is congruent to n+1
an automorphism F of the space of functions on G, whose matrix is
Fjk = exp(

−2πijk
)
n+1

(5 − 12).

This is of course the usual Fourier transform on the cyclic group µn+1 . For G non-abelian, we obtain a matrix-valued Fourier transform; by taking the trace of a representation, we
14

obtain an automorphism of the space of central functions on G; however, this is not of finite order, already in the case of the binary octahedral group (case D4 ).
This strongly suggests that the main object of interest should be the matrix-valued
Fourier transform, not just the scalar-valued Fourier transform.
REFERENCES
[A] M.Artin, On isolated rational singularities of surfaces, Amer. J. Math. 88 (1966),
129-136
[A-V] M. Artin and J-L. Verdier, Reflexive modules over rational double points, Math.
Ann. 270 (1985), 79-82
[Bo] N. Bourbaki, Lie Groups and Lie Algebras, Chapters 4,5,6
[Br] E. Brieskorn, Rationale singularitäten komplexer Flächen, Invent. Math. 4
(1968), 336-358
[Co] H. S. M. Coxeter, Regular Complex Polytopes, Cambridge Univ. Press (1974)
[De] P. Deligne, Equations Différentielles à Points Singuliers Réguliers, Lecture Notes on Math vol. 163 (1970), Springer Verlag
[Du Val1] P. Du Val, On isolated singularities which do not affect the condition of adjunction, Proc. Cambridge Phil. Soc. 30 (1934), 453-465
[Du Val2] P. Du Val, Homographies, Quaternions and Rotations, Clarendon Press
(1964)
[E-V] H. Esnault and E. Viewhweg, Logarithmic de Rham complexes and vanishing theorems, Invent. Math. 86 (1986), 161-194
[GS-V] G. Gonzalez-Sprinberg and J-L. Verdier, Construction géométrique de la correspondence de McKay, Ann. Sc. ec. Norm. Sup. 16 (1983), 409-449
[H1] F. Hirzebruch, Über vierdimensionale Riemmansche Flächen mehr-deutiger analytischer Funktionen von zwei komplexer Veränderlichen, Math. Ann. 126
(1953), 1-22
[H-N-K] F. Hirzebruch, W. D. Neumann and S. S. Koch, Differentiable manifolds and Quadratic Forms, Lecture Notes in Pure and Applied Math. vol 4 , Marcel Dekker
(1971)
[I-R] Y. Ito and M. Reid, The Mckay correspondence for finite subgroups of SL(3, C), preprint (1994), al-geom/9411010
[Kn] H. Knörrer, Group representations and the resolution of rational double points,
Contemp. Math. vol. 45 91985), 175-221
[Ko] B. Kostant, On finite subgroups of SU (2), simple Lie algebras and the McKay correspondence, Astérisque vol. Hors-Série (1985), 109-255
[L-T] G. Lusztig and J. Tits, The inverse of a Cartan matrix, Ann. Univ. Timişoara
Ser Ştiinţ. Mat. 30 (1992), no. 1, 17-23
[McKay1] J. McKay, Graphs, singularities and finite groups, Proc. Symp. Pure
Math. 37 (1980), 183-186
[McKay2] J. McKay, Cartan matrices, finite groups of quaternions, and kleinian singularities, Proc. Amer. Math. Soc. (1981), 153-154
[Mu] D. Mumford, The topology of normal singularities of an algebraic surface and a criterion for simplicity, Publ. Math. IHES 9 (1961), 23-64
15

[St1] R. Steinberg, Kleinian singularities and unipotent elements, Proc. Symp. Pure
Math. 37 (1980), 265-270
[St2] R. Steinberg, Subgroups of SU2 , Dynkin diagrams and affine Coxeter elements,
Pacific Jour. Math. 118 (1985), 587-598
[vR] R. von Randow, Zur Topologie von dreidimensionalen Baummannifal-tigkeiten, Bonner Math. Schriften 14 (1962)
Penn State University
Department of Mathematics
305 McAllister
University Park, PA. 16802
USA
e-mail:jlb@math.psu.edu

16

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
