---
source: "https://arxiv.org/abs/2310.03369v3"
title: "The Injectivity Radius of Souls of Alexandrov Spaces"
author: "Elena Mäder-Baumdicker, Jona Seidel"
year: "2023"
publication: "arXiv preprint / math.DG"
download: "https://arxiv.org/pdf/2310.03369v3"
pdf: "https://arxiv.org/pdf/2310.03369v3"
captured_at: "2026-06-27T06:15:54Z"
updated_at: "2026-06-27T06:15:54Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács Soul and Form"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# The Injectivity Radius of Souls of Alexandrov Spaces

- 著者: Elena Mäder-Baumdicker, Jona Seidel
- 年: 2023
- 掲載情報: arXiv preprint / math.DG
- 情報源: [arxiv](https://arxiv.org/abs/2310.03369v3)
- ダウンロード: https://arxiv.org/pdf/2310.03369v3
- PDF: https://arxiv.org/pdf/2310.03369v3

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

A sharp lower bound for the injectivity radius in noncompact nonnegatively curved Riemannian manifolds involving their soul goes back to Šarafutdinov. We generalize this bound to the setting of Alexandrov spaces. Our main theorem reads as follows. If the injectivity radius of an Alexandrov space of nonnegative curvature does not coincide with the one of its souls, then it is at least $ πK^{-1/2} $, where $ K $ is an upper curvature bound. We introduce the soul of Alexandrov spaces in some detail and compare two notions of injectivity radii.

## PDF Text

arXiv:2310.03369v3 [math.DG] 3 Jan 2025

The Injectivity Radius of Souls of Alexandrov
Spaces
Elena Mäder-Baumdicker∗, Jona Seidel†
January 7, 2025

Abstract
A sharp lower bound for the injectivity radius in noncompact nonnegatively curved Riemannian manifolds involving their soul goes back to Šarafutdinov. We generalize this bound to the setting of Alexandrov spaces. Our main theorem reads as follows. If the injectivity radius of an Alexandrov space of nonnegative curvature does not coincide with the one of its souls, then it is at least πK −1/2 , where K is an upper curvature bound. We introduce the soul of Alexandrov spaces in some detail and compare two notions of injectivity radii. 1

∗

maeder-baumdicker@mathematik.tu-darmstadt.de,
ORCID: 0000-0001-8125-8799,
Technical University Darmstadt, Department of Mathematics, Schlossgartenstr. 7, 64289
Darmstadt, Germany
†
seidel@mathematik.tu-darmstadt.de,
ORCID: 0009-0000-1282-8209,
Technical University Darmstadt, Department of Mathematics, Schlossgartenstr. 7, 64289
Darmstadt, Germany
1
Mathematics Subject Classification 53C20; 53C23, 53C45,
Keywords: Alexandrov Geometry, Soul, Injectivity Radius, Conjugate Points

1

Acknowledgments
We are thankful to Alexander Lytchak for extensive discussions which were essential to our understanding and improved the exposition of this paper at various places. We are also thankful to Karsten Große-Brauckmann for his helpful comments and commitment to our fruitful discussions. Special thanks should also go to Fernando Galaz-Garcı́a and Diego Corro who together with the first author suggested the formulation of Theorem 1 and shared helpful comments. The first author is partially supported by the
Deutsche Forschungsgemeinschaft DFG (MA 7559/1-2). This work is part of the second author’s PhD thesis.

2

Contents
1 Introduction

3

2 Preliminaries

6

3 The Cheeger-Gromoll Soul
12
3.1 The Soul Construction . . . . . . . . . . . . . . . . . . . . . . 12
3.2 Šarafutdinov’s Retraction . . . . . . . . . . . . . . . . . . . . 22
4 The Injectivity Radius in Length Spaces
26
4.1 Two Injectivity Radii . . . . . . . . . . . . . . . . . . . . . . . 26
4.2 Conjugate Points . . . . . . . . . . . . . . . . . . . . . . . . . 27
4.3 The Injectivity Radius Before Conjugate Points . . . . . . . . 30
4.4 Conjugate Points and an Upper Curvature Bound . . . . . . 32
5 The Injectivity Radius of a Soul

35

A Appendix
40
A.1 Hausdorff Distance . . . . . . . . . . . . . . . . . . . . . . . . 40
A.2 Real Functions with One-Sided Derivative . . . . . . . . . . . 41
A.3 Klingenberg’s Injectivity Radius Estimate . . . . . . . . . . . 42
References

1

43

Introduction

An Alexandrov space is a complete length space, where curvature bounds are introduced by triangle comparison. This defines a wide class of spaces including complete Riemannian manifolds and their limits [11, 5, 10]. The metric toolbox of Alexandrov geometry is different to the usual apparatus of smooth differentiable structures and builds upon basic descriptions of geometry like distance and angle. Many results from Riemannian geometry have been generalized to Alexandrov spaces. In fact, some questions appear to be more naturally described in the metric vocabulary and lead to new results even for Riemannian manifolds, e.g. see [4, 2].
The concept of the soul was a critical contribution to the study of Riemannian manifolds of nonnegative curvature by Cheeger and Gromoll [12].
It is a compact submanifold without boundary that contains any (local)
geodesic joining two of its points. The Soul Theorem [12, Theorem 2.2]
essentially reduces the study of complete nonnegatively curved Riemannian
3

manifolds to the compact case. It states that these manifolds are diffeomorphic to the normal bundle over their soul. The construction of the soul in [12] was generalized to Alexandrov spaces by Perelman [25]. We call a soul which can be obtained by this construction a Cheeger-Gromoll soul.
It was also shown in [25] that the full Soul Theorem does not generalize to Alexandrov spaces (see Example 3.17). Still, the soul serves as a key tool to investigate the structure of nonnegatively curved Alexandrov spaces
[41, 20, 15, 37, 14]. For the convenience of the reader we include a section summarizing the construction of the soul in [25]. Moreover, we state a proof for the uniqueness of Cheeger-Gromoll souls up to isometry.
We show that, while the Soul Theorem does not generalize to Alexandrov spaces, a theorem by Šarafutdinov [39] does. It establishes a lower bound for the injectivity radius inj(M ) of a Riemannian manifold M in terms of an upper curvature bound and the injectivity radius of its soul. In Riemannian manifolds, each pair of points at a distance less than inj(M ) are uniquely joined by a geodesic, and each geodesic of length less than inj(M )
is minimizing. These properties provide two natural notions of an injectivity radius in Alexandrov spaces, cf. Section 4. Lower bounds on the injectivity radius are significant as they control the volume growth of balls [13], retain curvature bounds under limits [5, Prop. 11.3], prevent collapsing or yield topological implications such as sphere theorems [7, 6.5].
An alternative proof to Šarafutdinov’s theorem was given by Croke and
Karcher [13, Theorem I]. We adapt it to the setting of Alexandrov spaces and obtain the following as our main result.
Theorem 1. Let K > 0 and X be a finite-dimensional Alexandrov space of nonnegative curvature bounded above by K and S ⊆ X be a Cheeger-Gromoll soul. If inj(X) < √πK , then inj(X) = inj(S).
The assumption inj(X) < √πK in Theorem 1 is necessary. The injectivity radius of the paraboloid is finite. However, for its soul, a singleton, it is infinite.
Remark 1.1. Note that, as a consequence of Theorem 1 we obtain the lower bound on the injectivity radius


π
inj(X) ≥ min √ , inj(S) .
K
In this form the theorem is stated in [39, 13]. However, the slightly stronger formulation in Theorem 1 also follows from their proofs.
4

Notably, Theorem 1 can reduce the study of the injectivity radius to compact spaces without boundary, a case well studied in the Riemannian setting, see [7, 6.5.2]. In addition, further information on the soul can give a lower bound independent of quantities such as dimension, volume or diameter. The following corollary is an example of that. It is known that in the case of positive curvature (see Definition 2.4), the soul is a single point unless it coincides with the entire space (see Proposition 3.16). A space consisting of a single point has infinite injectivity radius, and so we conclude:
Corollary 1.2. Let K > 0 and X be a finite-dimensional Alexandrov space of positive curvature bounded above by K. If X is noncompact or has nonempty boundary, then
π
inj(X) ≥ √ .
K
This contrasts to the case of compact spaces without boundary where curvature bounds are not sufficient to establish a lower bound for the injectivity radius [7, 6.5.2].
Remark 1.3. In Riemannian manifolds a single point with positive sectional curvatures suffices to guarantee a singleton soul. This is the Soul Conjecture which was proposed in [12] and famously proved in [26]. In Alexandrov spaces this is still open, with some recent progress [41, 20, 33]. In Alexandrov spaces where the Soul Conjecture continues to hold, it allows for the weaker curvature assumption in Corollary 1.2.
Remark 1.4. Although we work in the setting of Alexandrov spaces, the assumptions of Theorem 1 yield more regularity than one might expect.
Two-sided curvature bounds imply that X and S are topological manifolds, possibly with boundary [31]. Moreover, S has empty Alexandrov boundary
(see Definition 2.10). Therefore, S has empty boundary in the manifold sense (Corollary 2.12) and hence satisfies local extendability of geodesics [9, ch. II, 5.12]. This qualifies S for the theory of [24, 5] which yields that S
is in fact a Riemannian manifold, although not necessarily smooth [5, 14.1].
Furthermore, it can be approximated by smooth Riemannian manifolds [5,
15.1]. However, in the present paper we do not use the Riemannian metric or any of these results. Instead, the metric tools of Alexandrov spaces suffice to conclude Theorem 1.
The role of the curvature bounds in Theorem 1 can be seen as follows.
Nonnegative curvature provides the existence of the soul and thereby controls the length of closed geodesics. On the other hand, the upper curvature
5

bound controls the distance to conjugate points. These two quantities determine the injectivity radius, see Lemma 4.10.
The paper is organized as follows. After introducing basic facts in
Alexandrov geometry in Section 2, we present in Section 3 Perelman’s existence proof of the soul and gather some properties of the soul construction.
In Section 4, we introduce two definitions of an injectivity radius in Alexandrov spaces and treat conjugate points. These are the tools necessary to prove Theorem 1 in Section 5. Finally, Appendix A collects elementary proofs and an addendum to [34, Theorem 8.3].

2

Preliminaries

In this section we fix notation and recall basic properties of Alexandrov spaces. We mainly refer to [9, 10, 2]. Other excellent introductions include
[11, 36, 32].
In the following κ, K ∈ R denote real numbers and n a natural number.
If not specified otherwise, d denotes the metric of a metric space X. We denote an open ball of radius r ∈ R at p ∈ X by Br (p). By reversal of a curve c : [a, b] → X, we mean the curve −c : [a, b] → X, t 7→ c(a + b − t).
Definition 2.1. Let X be a metric space.
(i) A length space is a metric space where the distance between two points is given by the infimum over lengths of all rectifiable curves joining those two points.
(ii) A curve c : [a, b] → X is called a (local) geodesic if it is a linear reparametrization of a (local) isometry. If it is a (local) isometry, then it is called a unit-speed (local) geodesic. A curve c : [0, ∞) → X is called a (unit-speed) ray if each restriction to a compact interval is a nontrivial (unit-speed) geodesic.
(iii) X is called geodesic if for each pair of points there exists a geodesic joining them.
(iv) A subset C ⊆ X is called convex if any two points can be joint by a geodesic that lies entirely in C.
Note that our local geodesics correspond to some authors notion of
“geodesics” and our geodesics to “minimizing geodesics”. While others require (local) geodesics to be parametrized by arc length (unit-speed), the above definition allows for parametrizations proportional to arc length.
6

Curvature bounds can be introduced in a length space X via comparison of triangles in X to triangles of the same side lengths in the simply connected 2-dimensional Riemannian manifolds MK of constant curvature
K. We denote its distance metric by dK and set DK to be its diameter, namely
(
∞
if K ≤ 0,
DK :=
π
√
if K > 0.
K
Triangles in MK are well understood and described by the law of cosines
[2, Section 1.A]. For a choice of three pairwise distinct points p, q, r ∈ X
satisfying d(p, q) + d(q, r) + d(r, p) < 2DK we denote by p̃, q̃, r̃ ∈ MK a choice of points such that each pair is the same distance apart as their
˜ K we denote the triangle with vertices p̃, q̃ and r̃
counterparts in X. By ∆
K
˜
˜ K in p̃.
and ∠p (q, r) denotes the angle of ∆
Definition 2.2. Let X be a geodesic space. X is a CAT(K) space (resp.
CBB(K) space) if it satisfies the following for all pairwise distinct points p, x, y ∈ X with d(p, x) + d(x, y) + d(y, p) < 2DK .
• Let z be a point on a geodesic joining x to y. The point z̃ ∈ MK with d˜K (z̃, x̃) = d(z, x) and d˜K (z̃, ỹ) = d(z, y) satisfies d(p, z) ≤ d˜K (p̃, z̃)

(resp. d(p, z) ≥ d˜K (p̃, z̃)).

We say that a metric space is of curvature ≤ K (resp. ≥ K) if each point admits a neighborhood which is a CAT(K) space (resp. CBB(K) space).
See [2] for equivalent and more general definitions.
Definition 2.3. An Alexandrov space is a complete length space that admits a curvature bound in the sense of Definition 2.2.
The notion of an Alexandrov space is not consistent within literature and the different definitions differ in whether they require completeness, local compactness, or a lower curvature bound. Note that by the Globalization
Theorem [2, 8.31] a geodesic Alexandrov space of curvature ≥ K is also a
CBB(K) space. That is, the local triangle comparison holds in fact globally.
However, an analogous statement for upper curvature bounds is not true.
Alexandrov spaces of curvature ≥ K (resp. ≤ K) generalize smooth complete connected Riemannian manifolds without boundary with sectional curvatures ≥ K (resp. ≤ K) (cf. [28, Theorem 12.2.2 + Exercise 12.8.4]
and [9, ch. II, 1A.6], respectively). When referring to Riemannian manifolds
7

in the following, we always mean smooth complete connected Riemannian manifolds without boundary. A natural generalization of strict inequalities for sectional curvatures in Riemannian manifolds is the following.
Definition 2.4. Let κ ∈ R and X be a metric space such that for each point p there exists a constant κp > κ such that a neighborhood at p is of curvature ≥ κp . Then we call X of curvature > κ.
A proper Alexandrov space X of curvature ≥ κ > 0 is compact if it is not isometric to the real (half-)line (Lemma 2.6 (ix)+(x)). Therefore, except for the real (half-)line, a locally compact Alexandrov space of positive curvature is of curvature ≥ κ > 0 if and only if it is compact.
Definition 2.5. Given two nontrivial geodesics γ, σ in a metric space X
issuing in a common point p ∈ X, we define the angle
˜ 0 (γ(s), σ(t))
∠p (γ, σ) := lim ∠
p s,t↘0

if the limit exists. We define Σ′p to be the set of all nontrivial unit-speed geodesics issuing in p, pairwise identified if they have an angle of zero. The angle ∠p defines a metric on Σ′p . The space of directions Σp is the completion of Σ′p .
We gather standard facts about Alexandrov spaces.
Lemma 2.6. A complete, locally compact length space X has the following properties.
(i) [9, ch. I, 3.7]. X is proper and geodesic.
(ii) [9, ch. I, 3.11]. A sequence of geodesics cn : [0, 1] → X contained in a bounded ball admits a convergent subsequence that converges uniformly to a (possibly constant) geodesic c : [0, 1] → X.
(iii) If X is of curvature ≥ κ (resp. ≤ K) and A ⊆ X is a closed and convex subset, then A is an Alexandrov space of curvature ≥ κ (resp. ≤ K).
(iv) [2, 8.14+9.14]. If X is of curvature ≥ κ (resp. ≤ κ) and γ, σ are non˜ κ (γ(s), σ(t))
trivial geodesics issuing in p ∈ X, then the map (s, t) 7→ ∠
p is nonincreasing (resp. nondecreasing) in both arguments. In particular, all angles exist.
(v) [2, 9.45]. If X is of curvature ≤ K and (γn : [0, 1] → X)n∈N is a sequence of local geodesics which converge uniformly to some curve
γ : [0, 1] → X, then γ is a local geodesic and lim L(γn ) = L(γ).
n→∞

8

(vi) [9, ch. II, 1.4]. If X is CAT(K), then any local geodesic of length
< DK is geodesic and each pair of points x, y ∈ X with d(x, y) < DK
is joined by a unique geodesic. This geodesic varies continuously with its endpoints.
(vii) [2, 8.39]. If X is of curvature bounded below, c : [0, 1] → X is a geodesic joining distinct points x, y ∈ X and γ is a geodesic joining an inner point p ∈ c((0, 1)) to some point q ∈ X \ {p}, then
∠p (γ ′ , x′ ) + ∠p (γ ′ , y ′ ) = π, where x′ , y ′ , γ ′ ∈ Σp are the directions corresponding to −c, c and γ, respectively.
(viii) [2, 8.37]. If X is of curvature bounded below and two geodesics issuing in a common point coincide on an open interval, then one is an extension of the other.
(ix) [2, 8.44]. If X is of curvature ≥ κ and not one-dimensional, then diam X ≤ Dκ .
(x) [2, 15.18]. If X is one-dimensional and of curvature bounded below, then it is isometric to a connected complete one-dimensional Riemannian manifold with possibly nonempty boundary.
Remark 2.7. Note that, other than for spaces of curvature ≤ K, for CAT(K)
spaces Corollary 1.2 even holds without a lower curvature bound and is an immediate consequence of the CAT(K) property Lemma 2.6 (vi).
The last statements refer to the dimension of an Alexandrov space which we introduce now.
Definition 2.8. The dimension of an Alexandrov space is defined to be its
Hausdorff dimension.
Notably, every finite-dimensional Alexandrov space is locally compact and the statements in Lemma 2.6 apply. The dimension of an Alexandrov space is an integer (or infinite). Furthermore, it agrees with the so-called strainer number which gives the construction of the distance charts, cf. [10, ch. 10]. We note a rather specific but simple observation of this theory for later use.
Remark 2.9. Let X be an n-dimensional Alexandrov space of curvature ≥ κ
and A ⊆ X be a closed and convex subset of dimension n. Then there exists a point p ∈ A and a neighborhood in X of p which is contained in A.
9

This is a consequence of the distance coordinates in the following way.
Since A is n-dimensional, it contains an n-strained point p which is clearly also an n-strained point in X. There exists the distance coordinate map f which is a homeomorphism from an open neighborhood U ⊆ X of p to some subset of Rn (see [10, 10.8.18]). To prove that the distance coordinates map surjectively onto an open subset (see [10, 10.8.15]) it suffices to consider points along geodesics from p to the strainer points (ai , bi ) and their limits.
Since A is convex and closed, they are contained in A and hence f |U ∩A
is already surjective onto f (U ). But since f is injective, this implies that
U ∩ A = U , that is, U is contained in A.
Definition 2.10. The boundary of an n-dimensional Alexandrov space of curvature bounded below is inductively defined to be the set of points p ∈ X where the space of directions Σp , which is an Alexandrov space of curvature ≥ 1 and dimension n − 1 ([11, 7.10+7.11]), has nonempty boundary. One-dimensional Alexandrov spaces are one-dimensional manifolds (Lemma 2.6 (x)) and their boundary is defined to be their topological boundary. Zero-dimensional Alexandrov spaces are singleton sets and have, by definition, empty boundary.
The following statement is extracted from [22] and reveals that the
Alexandrov boundary behaves similarly to the manifold boundary.
Lemma 2.11. Let X be an n-dimensional Alexandrov space of curvature bounded below and U ⊆ X be a nonempty, open and connected subset.
Then the n-th singular homology group Hn (U ; Z/2Z) in Z/2Z coordinates is nontrivial if and only if X is compact with empty boundary and U = X.
Proof. First, note that if U ∩ ∂X = ∅, then U is an “NB-space” as defined in [22] which allows us to apply their results. Let U = X be compact and
∂X = ∅. In [22] a notion of “orientability” is defined. They show that if X
is “orientable”, then Hn (X; Z) = Z ([22, Theorem 1.8 (A)]) and if it is not, then still Hn (X; Z/2Z) = Z/2Z ([22, Corollary 5.7 (i)]). We have in both cases Hn (X; Z/2Z) ̸= {0} since Hn (X; Z) ̸= {0} implies Hn (X; Z/2Z) ̸=
{0} which is immediate by the definition of singular homology or a trivial application of the universal coefficient theorem for homology. To show the converse, assume U is a proper subset of X or X is not compact or X has nonempty boundary. If U ∩ ∂X is nonempty, then, by [22, Corollary 1.14], we obtain Hn (U ; Z/2Z) = {0}. If U ∩ ∂X is empty, then either U = X and by assumption X is not compact or U ̸= X in which case U is a nonempty proper open subset of the connected space X and thus not compact either.
10

Hence, U is not compact but satisfies U ∩ ∂X = ∅ which allows us to apply
[22, Theorem 1.8 (C)].
As a consequence, the boundary of Alexandrov spaces generalizes the manifold boundary in the following sense.
Corollary 2.12. If φ : X → M is a homeomorphism between a finitedimensional Alexandrov space X of curvature bounded below and a topological manifold M , then it satisfies φ(∂X) = ∂M .
Proof. Let n := dim X and note that M must be n-dimensional too. The relative homology groups of the space of directions at a point p ∈ X indicates if it has nonempty (Hn−1 (Σp ; Z/2Z) = {0}) or empty (Hn−1 (Σp ; Z/2Z) ̸=
{0}) boundary (see Lemma 2.11). And by [22, Theorem 6.4] a sufficiently small ball B at p ∈ X is pointed homeomorphic to the tangent cone K(Σp )
with the cone-tip 0p . The tangent cone is contractible but with the tip removed it is homotopy equivalent to Σp . The long exact sequence for relative homology groups yields (all homology groups in Z/2Z coefficients)
Hn (B, B \ {p}) ≈ Hn (K(Σp ), K(Σp ) \ {0p }) ≈ Hn−1 (Σp ).
We can find an open neighborhood C ⊆ φ(B) of φ(p) which is homeomorphic to Rn or its closed half space. By excision, we conclude
Hn−1 (Σp ) ≈ Hn (B, B \ {p}) ≈ Hn (C, C \ {φ(p)})
(
{0}
if φ(p) is a manifold boundary point,
≈
Z/2Z
if φ(p) is a manifold inner point.
Therefore, p is a boundary point if and only if φ(p) is a (manifold) boundary point.
A common construction that produces spaces with empty boundary is the following.
Definition 2.13. Let X be a finite-dimensional Alexandrov space of curvature ≥ κ with nonempty boundary. The doubling X of X is defined to be the metric space obtained from gluing together two copies of X along its boundary equipped with the induced length metric.
The doubling is again an Alexandrov space of curvature ≥ κ and has empty boundary (cf. [25, 5.2] and [29, 2.1]). In particular, this implies the following.
11

Lemma 2.14. Let X be an Alexandrov space of curvature ≥ κ > 0 and finite dimension ≥ 2 with nonempty boundary. Then the distance to the boundary dist∂X is bounded by Dκ /2.
Proof. Consider the doubling X of X which satisfies diam X ≤ Dκ by
Lemma 2.6 (ix). Fix p ∈ X and assume dist∂X (p) > 0. Then the distance of p to its reflection point p̄ in the other copy of X satisfies d(p, p̄) =
2 inf q∈∂X d(p, q). Therefore,
1
Dκ
dist∂X (p) = d(p, p̄) ≤
.
2
2

3

The Cheeger-Gromoll Soul

In this section we summarize the results of [25] that generalize the soul to
Alexandrov spaces. For convenience of the reader we do so in a detailed manner.

3.1

The Soul Construction

To define the soul let us first recall the definition of totally convex subsets.
Definition 3.1. Let X be a metric space. A subset A ⊆ X is said to be totally convex if every pair of points in A can be joined by a local geodesic in X and the image of every such local geodesic is contained in A.
Total convexity is far more restricting than convexity. For instance, the only totally convex subset of the sphere Sn is Sn itself. In Riemannian manifolds the soul of a manifold is a compact, totally geodesic and totally convex submanifold (without boundary). This leads to the following generalization to Alexandrov spaces.
Definition 3.2. Let X be a finite-dimensional Alexandrov space of curvature bounded below. A nonempty, totally convex, compact subset of X
without boundary is called a soul of X.
Note that, as a closed and convex subset, a soul is a finite-dimensional
Alexandrov space with lower curvature bound itself and the boundary is defined. Since a closed and convex subset of a Riemannian manifold is a submanifold possibly with (non-smooth) boundary ([12, 1.6]), the preceding definition of a soul in Alexandrov space generalizes the Riemannian soul.
A compact space X without boundary always serves as its own soul and is thus rather uninteresting. But for other spaces the soul can serve as a
12

compact and boundaryless reduction of the whole space, inheriting some of its topological and geometric properties. A natural question is if such a soul always exists. Cheeger and Gromoll gave a constructive proof for
Riemannian manifolds with nonnegative sectional curvature in [12] which
Perelman generalized to Alexandrov spaces in [25, 6]. To obtain totally convex compact subsets of a noncompact space, the so-called Busemann functions will turn out to be useful.
Definition 3.3. Let X be a metric space and c : [0, ∞) → X be a unit-speed ray. Then the Busemann function bc is given by bc : X → R,

q 7→ lim d(q, c(t)) − t t→∞

Remark 3.4. By the standard and reverse triangle inequality, the map t 7→ d(q, c(t)) − t is nonincreasing and bounded from below. Hence, the
Busemann function is well-defined. Furthermore, standard arguments show that it is continuous. Superlevel sets can be understood as half spaces, e.g.
X \ b−1
c ([0, ∞)) = ∪r>0 Br (c(r)).
The following establishes the existence of rays in noncompact spaces to enable the use of Busemann functions. This is a well-known and often used fact and mainly due to the compactness of the space of directions which was proved in [11, 7.3]. We include a proof for the convenience of the reader.
Lemma 3.5. Let X be a noncompact finite-dimensional Alexandrov space of curvature bounded below. Then for every point p ∈ X, there exists a ray issuing in p.
Proof. Since X is proper and not compact it must be unbounded. Therefore, there exists a sequence of points (qi )i∈N with d(p, qi ) → ∞. For i ∈ N, fix the directions ξi ∈ Σp corresponding to a geodesic from p to qi . The space of directions Σp is compact and hence there is a convergent subsequence of directions (ξin )n∈N which correspond to unit-speed geodesics (cn )n∈N from p to qin . Let ξ ∈ Σp be the limit of that sequence and rename ξn := ξin .
Fix t ∈ [0, ∞) and consider the closed ball of radius t at p which contains the restrictions (cn |[0,t] )n∈N for sufficiently large n. By Lemma 2.6 (ii), there exists a subsequence (cnk |[0,t] )k∈N which converges to a unit-speed geodesic rt of length t. By the law of cosines, it has direction ξ. We show that the whole sequence cn |[0,t] converges to rt . Consider s ∈ [0, t] and the comparison
˜ K (p, cn (s), rt (s)) with the angle triangle ∆
˜ K (cn (s), rt (s)) ≤ ∠p (ξn , ξ)
α̃n (s) := ∠
p
13

This gives a bound independent of s that tends to zero. Hence, so does the
˜ K (p, cn (s), rt (s))
length d(cn (s), rt (s)) of the side in the isosceles triangle ∆
opposite to the angle α̃n (s). Consequently, the geodesics cn |[0,t] converge uniformly to rt . This allows us to define, for all t ∈ [0, ∞), the pointwise limit r(t) := limn→∞ cn (t) = rt (t). To show that r is indeed a ray, let t ∈ [0, ∞). Since (cn |[0,t] )n∈N converges uniformly to rt , we have r(s) = rt (s)
for all s ∈ [0, t] and therefore r|[0,t] = rt is a geodesic. Therefore, r is a ray.
We now give a short informal description of the soul construction. To find a soul of a finite-dimensional Alexandrov space X of nonnegative curvature, proceed as follows. If X is not compact, choose an arbitrary point p0 ∈ X.
Obtain a compact totally convex subset as the intersection of superlevel sets of all Busemann functions corresponding to rays emanating from p. As long as this set has nonempty boundary, obtain the next set as the subset of maximum distance to the boundary. In each step the dimension is reduced so that we arrive at a subset of dimension zero or a subset without boundary after finitely many steps. This is a soul (see Figure 1).

Figure 1: The soul construction on the cylinder terminating in a soul S. C
is a superlevel set of the minimum of the Busemann functions corresponding to the two rays issuing in p0
We need to prove that the construction outlined above is well-defined, of which one part is that it yields a totally convex subset in each step. For this, the concavity of the Busemann function and the boundary distance function is crucial in the following way.
Definition 3.6. Let X be a metric space. A function f : X → R is said to be (strictly) concave if, for every nontrivial geodesic c : [0, L] → X, the real map f ◦ c is (strictly) concave.
Note that some authors use the notion of “convexity” of functions (e.g.
[12, 1.9],[35, 4.2],[25, 6.1]) which corresponds to our definition of concav14

ity. Also the meaning of “strictly concave” varies in the literature (see
Remark 3.11). Note that even the composition of a concave function with a local geodesic is concave since concavity of real continuous functions is a local property (e.g. see [23, Theorem 1.1.3]). An immediate consequence is the following.
Lemma 3.7. Let X be a metric space and f : X → R be a concave function.
Then for each a ∈ R the superlevel set f −1 ([a, ∞)) is totally convex.
Consider a geodesic segment c on the southern hemisphere of the twodimensional sphere of curvature κ > 0 with north pole N . Then the distance function t 7→ d(c(t), N ) is concave. By triangle comparison (Definition 2.2)
one can translate this to Alexandrov spaces of curvature ≥ κ > 0, i.e. the distance function is concave above D2κ , half the diameter of the sphere. In the case of nonnegative curvature where Dκ = ∞ the distance function indeed becomes concave “at infinity”, that is, the Busemann function is concave.
This is well known and a proof can be found for instance in [35, 4.2] or [2,
Exercise 8.25].
Lemma 3.8. Let X be an Alexandrov space of nonnegative curvature and c : [0, ∞) → X be a unit-speed ray. Then the Busemann function bc is concave.
The concavity of the boundary distance function dist∂X : X → [0, ∞), p 7→ d(p, ∂X).
is less obvious. It was proved in [25] and later put into a more general context in [1]. In [30, 3.3.1] a proof can be found that simplifies the application of some of the structure theory in [25] to the use of the gradient exponential map. To treat the case of a positive lower curvature bound, we introduce a generalized notion of concavity.
Definition 3.9. Let c ∈ R, I ⊆ R an interval, t0 ∈ I and h : I ⊆ R → R be a continuous function. We say that h satisfies h′′ (t0 ) ≤ c in the barrier sense if there exists a smooth function g : I → R such that h ≤ g, h(t0 ) = g(t0 )
and g ′′ (t0 ) ≤ c. For a continuous function φ : R → R, we say that h′′ ≤ φ ◦ h in the barrier sense (or h is (φ ◦ h)-concave) if for all t0 ∈ I, we have h′′ (t0 ) ≤ φ(h(t0 )) in the barrier sense. A function f : X → R on a metric space X is called (φ ◦ f )-concave if h := f ◦ c is (φ ◦ h)-concave for each unit-speed geodesic c in X.

15

Note that there are other common definitions of differential inequalities, some of which are equivalent to the above. We are especially interested in the case φ(f ) = λ − κf with λ, κ ∈ R for which equivalent definitions can be found in [2, 3.14].
Proposition 3.10. Let X be a finite-dimensional Alexandrov space of curvature ≥ κ with nonempty boundary ∂X.
(i) [25, 6.1]. If κ = 0, then dist∂X is concave.
√
(ii) [1, 1.1 (5B)]. If κ > 0, then f := sin ( κ dist∂X ) is (−κf )-concave.
√
Remark 3.11. For κ > 0 the function sin( κ dist∂X ) is, in particular, strictly concave on X \ ∂X but dist∂X itself is not. Along any geodesic c that minimizes the distance of some point p ∈ X \ ∂X to the boundary, the composition dist∂X ◦ c is linear. For illustration, take a geodesic connecting the north pole to the equator on the closed northern hemisphere of S2 . Note that in [25, 6.1] the boundary distance function dist∂X is called “strictly convex” (or “strictly concave” when cited in [1] or [41]) which, hence, does not correspond to our notion of a strictly concave function.
We retain the following as a direct consequence of Proposition 3.10.
Corollary 3.12. Let X be a finite-dimensional Alexandrov space of curvature ≥ κ > 0 with nonempty boundary and c : [0, L] → X \ ∂X a nontrivial geodesic, then dist∂X ◦ c attains a unique maximum.
Proof. This is obvious for zero- and one-dimensional Alexandrov spaces (see
Lemma 2.6 (x)). Assume dim X ≥ 2 and set h := dist∂X ◦ c and f :=
√
sin( κ h). The distance h to the closed set ∂X is positive and is bounded by Dκ /2 (see Lemma 2.14). Hence, we have m := min f > 0 and, by
Proposition 3.10 (ii), f ′′ ≤ −κf ≤ −κm

(in the barrier sense)

(1)

with −κm < 0. Now, assume h attains its maximum at two distinct points.
Then by concavity of h (Proposition 3.10 (i)), it is constant between those two points. But then f is also constant on that subinterval which is a contradiction to (1).
We summarize the properties of the superlevel sets of the soul construction (see also Figure 2).

16

Lemma 3.13. Let X be an n-dimensional Alexandrov space of nonnegative curvature. If X is not compact (resp. compact with nonempty boundary), then there exists a ≤ 0 and an interval I = [a, ∞) (resp. I = [a, 0]) and a family {Ct : t ∈ I} of nonempty, compact and totally convex subsets Ct ⊆ X
which satisfy:
(i) If t1 , t2 ∈ I with t1 ≤ t2 , then
Ct1 = {p ∈ Ct2 : d(p, ∂Ct2 ) ≥ t2 − t1 }, in particular Ct1 ⊆ Ct2 .
S
(ii)
Ct = X.
t∈I

(iii) The map t 7→ Ct is continuous w.r.t. the Hausdorff distance.
(iv) dim Ca < dim X.

Figure 2: The soul S and the sets Ct from Lemma 3.13 for a closed half-plane
X when starting in p0 . Some directions of rays issuing in p0 are indicated
Proof. The following proofs are largely analogous to the Riemannian case and of elementary character. The description in (i) was already formulated in [12]. Continuity of the map t 7→ Ct in (iii) was treated in [38, §2].
We assume that X is not compact. The compact case follows mostly analogously and is sometimes simpler with respective differences addressed in brackets. Fix some point p0 ∈ X. Then there exists at least one ray

17

emanating from p0 by Lemma 3.5. We consider the set Rp0 of all such unitspeed rays and define the infimum over all the corresponding Busemannfunctions b : X → R, p 7→ inf bc (p) = inf

lim d(p, c(s)) − s.

c∈Rp0 s→∞

c∈Rp0

(2)

Estimating b(p) ≥ −d(p, p0 ), we see that b is well-defined. Now set f := b
(resp. f := dist∂X ) and define for t ∈ R the superlevel set
Ct := f −1 ([−t, ∞)).
By Proposition 3.8 (resp. Proposition 3.10), the function f is an infimum over concave functions and thus itself concave. Fix t ∈ R. By concavity of f , the superlevel set Ct is totally convex. If X is compact, then the closed subset Ct is clearly compact itself. Otherwise, assume Ct would not be compact for some t ∈ R. Then Lemma 3.5 gives a unit-speed ray c emanating from p0 with image in Ct . But b(c(t + 1)) ≤ bc (c(t + 1)) = −t − 1 < −t and hence c(t + 1) cannot be contained in Ct , a contradiction. We conclude that Ct is compact. The set Ct could be empty but f is continuous and has compact superlevel sets and therefore assumes a maximal value −a. This allows us to set I := [a, ∞) (resp. I := [a, 0]) and obtain the family (Ct )t∈I
of nonempty, compact and totally convex subsets.
(i) We present the proof for noncompact X. The compact case follows similarly where instead of points c(s) on rays we directly consider closest points on the boundary. We first show f −1 ({−t}) = ∂Ct .

(3)

By continuity of f , we have ∂Ct ⊆ f −1 ({−t}). To show the converse inclusion let p ∈ f −1 ({−t}). Assume p is an inner point of Ct , then there exists
ε > 0 such that Bε (p) ⊆ Ct . The equality f (p) = −t allows us to choose a unit-speed ray c issuing in p0 and s > 0 such that d(p, c(s)) − s ≤ −t + 4ε .
Then there is a geodesic from p to c(s) and we consider the point q on this geodesic with distance 2ε to p. Recall that d(q, c(s)) − s is nonincreasing in s and thus not smaller than bc (q). We then have f (q) ≤ bc (q) ≤ d(q, c(s)) − s
ε
= d(p, c(s)) − s −
2
ε
≤ −t − < −t,
4
18

a contradiction to q ∈ Ct . Therefore, (3) follows.
Now let t1 , t2 ∈ I with t1 ≤ t2 and p ∈ Ct1 . We show that d(p, ∂Ct2 ) ≥
t2 −t1 . The reverse triangle inequality yields for all r ∈ ∂Ct2 , all rays c ∈ Rp0
and all s > 0 the inequality d(p, r) ≥ d(p, c(s)) − d(c(s), r)
= d(p, c(s)) − s − (d(c(s), r) − s)
≥ bc (p) − (d(c(s), r) − s)
≥ −t1 − (d(c(s), r) − s), and hence d(p, r) ≥ −t1 − f (r) = t2 − t1 . Since this holds for all r ∈ ∂Ct2 , we obtain d(p, ∂Ct2 ) ≥ t2 − t1 . Let conversely p ∈ Ct2 with d(p, ∂Ct2 ) ≥ t2 − t1 .
Let c ∈ Rp0 , s ≥ max{0, −t2 } and consider a geodesic from p to c(s + t2 ).
Since b(c(s + t2 )) ≤ bc (c(s + t2 )) = −s − t2 ≤ −t2 , this geodesic must cross
∂Ct2 at some point. Let rs denote that crossing point. For all sufficiently large s, d(p, c(s + t2 )) − (s + t2 ) = d(p, rs ) + d(rs , c(s + t2 )) − (s + t2 )
≥ d(p, ∂Ct2 ) + bc (rs )
≥ t2 − t1 + f (rs )
= −t1 , from which we obtain bc (p) ≥ −t1 for all c ∈ Rp0 . Consequently, f (p) ≥ −t1
and in particular p ∈ Ct1 .
S
Ct = f −1 (R) = X.
(ii) This follows immediately from t∈I

(iii) Fix t ∈ I and ε > 0. Let us denote the ε-neighborhood of a subset
A ⊆ X by Uε (A) := {p ∈ X : d(p, A) < ε}. We have to show that there exists
δ > 0 such that Ct+δ ⊆ Uε (Ct ) (if I = [0, ∞) or t < 0) and Ct ⊆ Uε (Ct−δ )
(if t > a).
The former relation follows just from continuity of f . To see this, assume the contrary. Then there exists a sequence (pn )n∈N such that pn ∈ Ct+1/n and d(pn , Ct ) ≥ ε (if I = [a, 0], we may start the sequence at n large enough such that t + n1 < 0). By compactness of Ct+1 (resp. X), there exists a limit p of a convergent subsequence (pnk )k∈N . We have f (pn ) ≥ −t − n1 and, by continuity, f (p) ≥ −t, i.e. p ∈ Ct . A contradiction to d(pnk , Ct ) ≥ ε.
To prove the second relation, assume on the contrary that for all δ > 0
the subset Ct \ Uε (Ct−δ ) is nonempty. Let p be a limit point of a sequence of points pn ∈ Ct \ Uε (Ct− 1 ) (start sequence at sufficiently large n such that n

t − n1 > a). Continuity of f and the relation d(p, Ct−δ ) ≥ ε for all δ > 0
19

imply p ∈ ∂Ct . Consider a unit-speed geodesic c : [0, L] → X joining p to some point in Ca . Using f (c(0)) = −t < −a = f (c(L))
and concavity of f ◦ c, we obtain for all s ∈ (0, L] that f (c(s)) > −t, that is, δs := d(c(s), ∂Ct )) > 0 and c(s) ∈ Ct−δs . For s ∈ (0, ε) this yields a contradiction with s = d(p, c(s)) ≥ d(p, Ct−δs ) ≥ ε.
(iv) Assume that X is compact. As a closed and convex subset of X, the set Ca is again a finite-dimensional Alexandrov space of nonnegative curvature. Assume that Ca has dimension n := dim X. Then by Remark 2.9, there exists an open neighborhood U in X of some point p ∈ Ca which is contained in Ca . Since ∂X is closed ([25, 4.6]), we can consider the shortest geodesic c from p to ∂X. Starting at the inner point p, this geodesic must pass through some ball Bε (p) ⊆ U ⊆ Ca with ε > 0. Hence, the distance function t 7→ d(c(t), ∂X) is constant with value −a on [0, ε) but it should decrease by the choice of c, a contradiction. We conclude that dim Ca < dim X. If X is not compact, let t > a. Setting X ′ := Ct , we apply the compact version of this Proposition. Then (i) implies that the obtained maximum level set Ca′ ′ coincides with Ca and the previous argument ensures dim Ca < dim Ct ≤ dim X.
Proposition 3.14. ([25, 6.2]). Every finite-dimensional Alexandrov space of nonnegative curvature contains a soul.
Proof. Let X be a finite-dimensional Alexandrov space of nonnegative curvature. If X is not compact, we apply Lemma 3.13 to get a nonempty, compact and totally convex subset X1 := Ca , otherwise we set X1 := X. Iteratively, we set Xn+1 to be the set Ca obtained from applying Lemma 3.13
to the compact space Xn as long as Xn has boundary and dim Xn ≥ 1.
The obtained set Xn+1 is again a nonempty, compact and totally convex subset of Xn . This is again an Alexandrov space of nonnegative curvature and qualifies for further application of Lemma 3.13. By transitivity of total convexity, Xn+1 is a totally convex subset not only of Xn but also of X.
Furthermore, we have dim Xn+1 < dim Xn and thus obtain after a finite number of steps either a totally convex singleton or a nonempty, compact and totally convex subset without boundary. In both cases we obtained a soul of X.
Definition 3.15. We call a soul obtained by the preceding construction a
Cheeger-Gromoll soul.
20

The soul of a positively curved space is given by the following.
Proposition 3.16 ([25, 6.2]). Let X be a finite-dimensional Alexandrov space of positive curvature which is noncompact or has nonempty boundary.
Then each Cheeger-Gromoll soul of X consists of a single point.
Proof. Let S be a Cheeger-Gromoll soul and let (Ct )t∈I be the superlevel sets from Lemma 3.13 that led to S. By property (i), we can assume that
X is compact (if not, set X := Ct with t > a) and thus satisfies a positive lower curvature bound. Suppose the set Ca of maximum distance to the boundary contains two distinct points. By convexity, it must contain a geodesic joining those two points. But, by Corollary 3.12, dist∂X assumes a unique maximum on this geodesic, a contradiction. Therefore, S = Ca consists of a single point.
Recall that the Soul Theorem [12, Theorem 2.2] states that a complete noncompact Riemannian manifold of nonnegative curvature is diffeomorphic to the normal bundle of its soul. This does not generalize to Alexandrov spaces in the following sense.
Example 3.17 ([25, 6.4]). There is a finitedimensional Alexandrov space of nonnegative curvature without boundary which is not homeomorphic to the total space of a locally trivial fiber bundle over its soul. We present the example [25,
6.4] with real instead of complex projective spaces.
Consider X := R3 /∼ where x ∼ ±x with induced length metric. X is isometric to the metric cone K(RP 2 ) over the real projective plane. Set
Figure 3: The soul S
Y := {(x, y, z) ∈ X : x2 + y 2 ≤ 1} and consider in Example 3.17
the doubling Y (see Definition 2.13). Starting the soul construction in the origin yields the Cheeger-Gromoll soul S that is the doubling of {(x, y, z) ∈ X : x2 + y 2 ≤ 1, z = 0} ⊆ X and hence homeomorphic to a sphere S2 (in fact isometric to the doubling of the closed unit ball in the cone K(RP 1 ), see Figure 3). Suppose there exists a projection π : Y → S defining a locally trivial fiber bundle. Then there exist disjoint open balls U, V ⊆ S with π(0) ∈ U such that neither U or V contain the image π(0) of the reflection 0 of the origin and the open subsets
π −1 (U ), π −1 (V ) ⊆ X are homeomorphic to the product of some fiber with U
(resp. V ). Therefore, there exists a homeomorphism φ : π −1 (U ) → π −1 (V ).
It induces an isomorphism between the relative singular homology groups
21

H3 (π −1 (U ), π −1 (U ) \ {0}) and H3 (π −1 (V ), π −1 (V ) \ {φ(0)}) (with Z coefficients). The following computation shows that there are indeed not isomorphic, yielding the desired contradiction. By excision, we can replace the open subsets by open balls BU ⊆ π −1 (U ) and BV ⊆ π −1 (V ) at 0 and φ(0), respectively. Note that BU is contractible and BU \ {0} is homotopy equivalent to ∂BU ≈ RP 2 . Similarly, BV is contractible, too, but BV \ {φ(0)} is homotopy equivalent to ∂BV ≈ S2 . The corresponding long exact homology sequences yield the remaining isomorphisms
H3 (π −1 (U ), π −1 (U ) \ {0}) ≈ H3 (BU , BU \ {0}) ≈ H2 (RP 2 ) ≈ Z/2Z,
H3 (π −1 (V ), π −1 (V ) \ {φ(0)}) ≈ H3 (BV , BV \ {φ(0)}) ≈ H2 (S2 ) ≈ Z.
Another formulation of the above is the following. Unlike for Riemannian manifolds, the (local) topology of an Alexandrov space is not determined by its soul. The manifold S × R has the same (isometric) Cheeger-Gromoll soul
S as the “non-manifold” Y . That Y is not a manifold can be proved by a computation of relative homology groups analogous to the above.
Nonetheless, a weaker version of the Soul Theorem generalizes to Alexandrov spaces. Namely, the soul is a strong deformation retract of the whole space. This is described in the next section.

3.2

Šarafutdinov’s Retraction

In [38] a non-expanding retraction of a Riemannian manifold onto its soul was found. It was generalized to Alexandrov spaces in [25, 6.3] and sparked a subsequent study of gradient flows in Alexandrov spaces which give a modern formulation of this retraction. We refer to [27, 30] and [2, ch. 16].
Proposition 3.18 ([25, 6.3] Šarafutdinov’s Retraction). Let X be a finitedimensional Alexandrov space of nonnegative curvature and S be a CheegerGromoll soul of X. Then there exists a continuous map F : X × [0, 1] → X
that satisfies for all t ∈ [0, 1], x ∈ X, s ∈ S,
F (x, 0) = x,

F (s, t) = s,

F (x, 1) ∈ S,

that is, S is a strong deformation retract of X. Furthermore, F is nonexpanding, that is, it satisfies for all p, q ∈ X and 0 ≤ t1 ≤ t2 ≤ 1 the inequality d(F (p, t2 ), F (q, t2 )) ≤ d(F (p, t1 ), F (q, t1 )).

22

Proof. We show that in each step of the soul construction such a map exists.
By concatenation one obtains the deformation onto the soul. Let X be noncompact (resp. compact with nonempty boundary). Let f denote the
Busemann function (resp. boundary distance function) from the proof of
Lemma 3.13 which assumes its maximum on a set Ca . By [21, 1.8], the gradient flow Φt : X → X associated to the concave function f is 1-Lipschitz and converges to a 1-Lipschitz function Φ∞ (p) = limt→∞ Φt (p) with image in Ca . This allows us to define the continuous map
(
Φtan(tπ/2) (p)
if t ∈ [0, 1),
F : X × [0, 1] → X, (p, t) 7→
∞
Φ (p)
if t = 1, which satisfies the desired properties of a deformation. Furthermore, it is non-expanding by the relation Φt2 = Φt2 −t1 ◦ Φt1 for 0 ≤ tt ≤ t2 < 1 and the flow being 1-Lipschitz.
Remark 3.19. In [25, 6.3] another flow is constructed (not using the modern notions of gradient curves and flows), namely a flow Ψt that satisfies f (Ψt (x)) = f (x) + t (if f (x) + t < a) and consists of curves β with
β + (t) = ∇β(t) f /|∇β(t) f |2 in contrast to the usual gradient curves α with
α+ (t) = ∇α(t) f that make up the gradient flow Φt . A modified retraction F̃ using Ψt instead of Φt is defined and yields the superlevel sets
Ct = F̃ (X × {t}) from Lemma 3.13.
A natural question about the soul is that of uniqueness and was studied in [40] for Riemannian manifolds. In a compact space, the soul construction yields unique Cheeger-Gromoll souls. But in noncompact spaces, the obtained Cheeger-Gromoll soul depends on the point chosen to start the construction and may yield a different soul for each starting point. An application of Šarafutdinov’s retraction yields the following.
Proposition 3.20. In a finite-dimensional Alexandrov space of nonnegative curvature any two Cheeger-Gromoll souls are isometric.
The proof of [40] carries over to Alexandrov spaces in a straightforward way. It requires the following two lemmata of which only the latter, originally formulated for compact manifolds, needs some further arguments to generalize to Alexandrov spaces.
Lemma 3.21 ([10, 1.6.15]). Let X be a compact metric space and f : X →
X be a surjective function. If f is 1-Lipschitz, then f is an isometry.

23

Lemma 3.22. Let X be a compact n-dimensional Alexandrov space of curvature bounded below without boundary. If f : X → X is homotopic to the identity map, then f is surjective.
Proof. If n = 0, then X is a single point and the statement is trivially true.
Assume n ≥ 1. The proof of [40] for manifolds only relies on the two facts about the singular homology groups: Hn (X) ̸= {0} and Hn (X \ {x0 }) = {0}
for all x0 ∈ X (and some fixed choice of coefficients). Note that X \ {x0 } is not compact since there are geodesics in X issuing in x0 . Hence, if X \{x0 } is connected, we obtain the desired homology groups from Lemma 2.11. Now, if f is not surjective, it can be written as a composition
X → X \ {x0 } → X
for some x0 ∈ X \ f (X). The induced homomorphisms
Hn (X) → Hn (X \ {x0 }) → Hn (X)
then yield with Hn (X \ {x0 }) = {0} that the induced homomorphism f∗ by f is zero. But since the identity id on X induces the identity isomorphism on Hn (X), the assumption Hn (X) ̸= {0} yields id∗ ̸= f∗ , a contradiction to the homotopy invariance of homology (cf. [16]).
It remains to show that X \ {x0 } is connected. Suppose it were not, then all points of a path-component are joined to some fixed point in a distinct path-component by a geodesic in X which passes through x0 . Since geodesics do not branch (Lemma 2.6 (viii)), this implies that each pathcomponent united with x0 is the image of a geodesic (or ray) emanating in x0 . Since X \ {x0 } is open in X and hence n-dimensional ([10, 10.6.1]), there exists an open neighborhood homeomorphic to some open subset Rn .
Therefore, n = 1. But in this case, by Lemma 2.6 (x) and X being compact without boundary, X must be homeomorphic to the circle S1 and removing a point still yields a connected subset. A contradiction.
The remaining proof works exactly as given in [40] and is illustrated here for the convenience of the reader.
Proof of Proposition 3.20. Let S1 , S2 ⊆ X denote two Cheeger-Gromoll souls of a finite-dimensional Alexandrov space X of nonnegative curvature. Let
F1 , F2 : X × [0, 1] → X be the deformations obtained from Proposition 3.18
corresponding to the soul constructions of S1 and S2 , respectively. For t ∈ [0, 1], we define the map ht : S1 → S1 ,

p 7→ F1 (F2 (p, t), 1).
24

Then h0 is the identity on S1 and the maps (ht )t∈[0,1] define a homotopy to the map h1 which is 1-Lipschitz by the non-expanding properties of F1 and
F2 (see Figure 4 for illustration). Applying the two previous Lemmata 3.21
and 3.22 yields that h1 is a surjective isometry on S1 . Clearly, the maps f1 : S2 → S1 ,

p 7→ F1 (p, 1),

f2 : S1 → S2 ,

p 7→ F2 (p, 1),

are non-expanding. Suppose there exists points p, q ∈ S1 that satisfy d(f2 (p), f2 (q)) < d(p, q). This would imply d(f1 (f2 (p)), f1 (f2 (q))) ≤ d(f2 (p), f2 (q)) < d(p, q), which contradicts h1 = f1 ◦ f2 being an isometry. Therefore, f2 is an isometry. Analogously, we obtain that f1 is an isometry. It remains to show that

Figure 4: Two souls on a cylinder with the isometries f1 , f2 and the homotopy (t, x) 7→ ht (x)
f2 is surjective. Suppose there exists a point p ∈ S2 \ f2 (S1 ). Since f1 is injective, it would follow that f1 (p) is not in the image of h1 = f1 ◦ f2 which is a surjective map, a contradiction. Hence, f1 : S1 → S2 defines a surjective isometry and the souls are isometric.
Note that souls that are not necessarily Cheeger-Gromoll souls can be much more abundant. For instance, in a compact convex subset of R2 the
Cheeger-Gromoll soul is a point and uniquely defined, but any other point is a soul, too. Uniqueness up to isometry (or even up to homeomorphism)
for these souls seems to be a more difficult question already in Riemannian manifolds (see [12, 2.1]).

25

4

The Injectivity Radius in Length Spaces

The usual definition of the

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
