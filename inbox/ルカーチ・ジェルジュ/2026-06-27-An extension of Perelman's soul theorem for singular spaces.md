---
source: "https://arxiv.org/abs/0706.0565v6"
title: "An extension of Perelman's soul theorem for singular spaces"
author: "Jianguo Cao, Bo Dai, Jiaqiang Mei"
year: "2007"
publication: "arXiv preprint / math.DG"
download: "https://arxiv.org/pdf/0706.0565v6"
pdf: "https://arxiv.org/pdf/0706.0565v6"
captured_at: "2026-06-27T06:16:01Z"
updated_at: "2026-06-27T06:16:01Z"
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

# An extension of Perelman's soul theorem for singular spaces

- 著者: Jianguo Cao, Bo Dai, Jiaqiang Mei
- 年: 2007
- 掲載情報: arXiv preprint / math.DG
- 情報源: [arxiv](https://arxiv.org/abs/0706.0565v6)
- ダウンロード: https://arxiv.org/pdf/0706.0565v6
- PDF: https://arxiv.org/pdf/0706.0565v6

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

In this paper, we study open complete metric spaces with non-negative curvature. Among other things, we establish an extension of Perelman's soul theorem for possibly singular spaces: "Let X be a complete, non-compact, finite dimensional Alexandrov space with non-negative curvature. Suppose that X has no boundary and has positive curvature on a non-empty open subset. Then X must be a contractible space". The proof of this result uses the detailed analysis of concavity of distance functions and Busemann functions on singular spaces with non-negative curvature. We will introduce a family of angular excess functions to measure convexity and extrinsic curvature of convex hypersurfaces in singular spaces. We also derive a new comparison for trapezoids in non-negatively curved spaces, which led to desired convexity estimates for the proof of our new soul theorem.

## PDF Text

AN EXTENSION OF PERELMAN’S SOUL

arXiv:0706.0565v6 [math.DG] 9 Jun 2010

THEOREM FOR SINGULAR SPACES

Jianguo Cao1 , Bo Dai2 and Jiaqiang Mei3
University of Notre Dame, Peking University and Nanjing University
In memory of Professor Xiao-Song Lin
Abstract. In this paper, we study open complete metric spaces with non-negative curvature. Among other things, we establish an extension of Perelman’s soul theorem for possibly singular spaces: “Let X be a complete, non-compact, finite dimensional
Alexandrov space with non-negative curvature. Suppose that X has no boundary and has positive curvature on a non-empty open subset. Then X must be a contractible space”. The proof of this result uses the detailed analysis of concavity of distance functions and Busemann functions on singular spaces with non-negative curvature.
We will introduce a family of angular excess functions to measure convexity and extrinsic curvature of convex hypersurfaces in singular spaces. We also derive a new comparison for trapezoids in non-negatively curved spaces, which led to desired convexity estimates for the proof of our new soul theorem.

§0. Introduction
In this paper, we will study open complete and finite dimensional metric spaces with non-negative curvature. A metric space (X, d) is called a length space if any pair of points {p, q} in X can be joined by a path of length equal to d(p, q). A
length-minimizing path of unit speed is called a geodesic. A metric space (X, d) is said to have non-negative curvature if any geodesic triangle △{l1 ,l2 ,l3 } of side lengths

{l1 , l2 , l3 } is “fatter” than the comparison triangle △∗{l1 ,l2 ,l3 } of the same lengths on
Key words and phrases. Singular spaces, Perelman’s soul theorem, Cheeger-Gromoll convex exhaustion, generalized soul theory, angular excess functions.
1 The first author is supported in part by an NSF grant and Changjiang Scholarship of China at Nanjing University.
2 The second author is supported in part by NSFC Grants 10621061, 10701003.
3 The third author is supported by Natural Science Foundation of China (grant 10401015).

1

the Euclidean plane. Similarly, we can define the notion of curvature ≥ k for metric spaces with any real number k ∈ R, see [BGP92], [BBI02].
It is known that if a metric space (X, d) has curvature bounded from below, then the topological dimension of X is equal to its Hausdorff dimension. Moreover, dim(X) must be an integer or infinity (cf. [BGP92]). Our main result of this paper is the following:
Theorem 0.1. Let X n be a complete and non-compact, n-dimensional metric space of non-negative curvature. Suppose that X n has no boundary and has positive curvature on a metric ball Bε0 (x0 ). Then X n must be contractible.
When X n is a smooth open Riemannian manifold of non-negative curvature, our
Theorem 0.1 above is related to the so-called Cheeger-Gromoll soul conjecture (cf.
[CG72]), which was successfully solved by Perelman (cf. [Per94a]).
However, there are examples of open metric spaces with positive curvature which are not topological manifolds. For instance, let
M n = {(x1 , . . . , xn , xn+1 ) ∈ Rn+1 | xn+1 = x21 + · · · + x2n }
and X n = M n /Z2 , where Z2 is a group generated by the involution
ϕ : Rn+1 → Rn+1
(x1 , . . . , xn , xn+1 ) 7→ (−x1 , . . . , −xn , xn+1 ).
It is known that X n is a space of positive curvature. The space Σn−1
(X n ) of unit
0
tangent directions of X n at the origin 0 is homeomorphic to the real projective space RP n−1 . Thus X n is not a manifold near the origin.
For smooth open Riemannian manifolds with non-negative curvature, Perelman
(cf. [Per94]) established a flat strip theorem and hence provided an affirmative solution to the Cheeger-Gromoll soul conjecture. The proof of Perelman’s flat strip theorem uses the fact that, for a smooth Riemannian manifold M n , its tangent
2

space Tx M n at x is always isometric to Rn . For an Alexandrov space X n with nonnegative curvature, its tangent cone Tx X n at a point x is not necessarily isometric to
Rn . Therefore, a different method is needed for the verification of our main theorem above. In next section, we sketch our new approach by the study of convexity of
Busemann functions on non-negatively curved singular spaces.
§1. Outline of the proof of main theorem
Our proof of Theorem 0.1 is inspired by H.Wu’s proof (cf. [Wu79], [Wu87]) of
Gromoll-Meyer theorem (cf. [GM69]).
Proposition 1.1. ([GM69], [Wu79], [Wu87]) Suppose that M n is a complete and non-compact smooth Riemannian manifold with positive sectional curvature, and suppose that h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

is a Busemann function, where Bt (x0 ) = {y ∈ M n |d(x0 , y) < t} is a metric ball of

radius t centered at x0 . Then the function (1 − e−h ) is a proper and strictly concave function with a unique maximum point p̂ ∈ M n . Consequently, M n is contractible to a point p̂ and hence M n is diffeomorphic to Rn .

Perelman (cf. [Per91]) was able to derive a similar result for singular spaces as well. To proceed, we need to recall the notion of λ-concavity introduced by
Perelman (cf. [Per94a]) for functions defined on singular spaces.
Definition 1.2. (λ-concave functions [Per94]) Let f : X n → R be a continuous function defined on an Alexandrov space with curvature ≥ −1. We say that f is λ-concave at p (or Hess(f ) p ≤ λ) in barrier sense if for all quasi-geodesics
σ : (−ε, ε) → X n of unit speed with σ(0) = p, the inequality d2
[f (σ(t))] t=0 ≤ λ
dt2
holds in barrier sense.
3

Proposition 1.3. (Perelman [Per91] Chapter 6) Let X n be a complete, non-compact and n-dimensional metric space of positive curvature. Suppose that X n has no boundary and h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞


is a Busemann function. Then the function f (x) = 1−e−h(x) is a strictly concave and a proper function with a unique maximum point p̂ ∈ X n . Consequently, X n is

contractible to the maximum point p̂ via the Sharafutdinov semi-flow
∇h d+ ϕ
=
.
dt
|∇h|2
We will review the Perelman-Sharafutdinov semi-gradient flow in upcoming sections.
We would like to say a few words about why we used f = 1 − e−h instead of the

Busemann function h. It is known (cf. [Wu79]) that, for any x∗ ∈ X n , there is a

geodesic ray σ : [0, +∞) → X n of unit speed such that σ(0) = x∗ and h(σ(t)) = h(x∗ ) − t.

It follows that h ◦ σ is a linear function and hence h can not be strictly concave along the geodesic ray σ. Hence, it is reasonable to consider f = 1 − e−h .

In our case, the singular space X n has positive curvature on a small ball Bε0 (x0 ).

The function f (x) = 1 − e−h(x) is only weakly concave on the whole space X n . The proof of Perelman’s result in Proposition 1.3 also implies the following.
Proposition 1.4. (Perelman [Per94a] ) Let X n be a complete, non-compact ndimensional metric space of non-negative curvature. Suppose that X n has no boundary and has positive curvature on a metric ball Bε0 (x0 ) and h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

as above. Then the function f (x) = 1 − e−h(x) is strictly concave on a small ball
Bε0 /4 (x0 ).
4

Inspired by Proposition 1.4, we will take a close look on the concavity of f (x) =
1−e−h(x) outside the small ball Bε0 /4 (x0 ). Following Cheeger-Gromoll (cf. [CG72]), we consider the convex sup-level sets:
ωc = h−1 ([c, +∞))
for all c ∈ R. it is known (cf. [Per94a] or [CMD09]) that Ωc is a totally convex subset

of X n . Moreover, its boundary ∂Ωc has strictly convex portion (∂Ωc ) ∩ Bε0 /4 (x0 )

when c ∈ [ĉ0 − ε40 , ĉ0 + ε40 ] and ĉ0 = h(x0 ). It is also known that each Ωc is compact.
Thus, we may assume c0 = maxn {h(x)} < +∞.
x∈X

(1.1)

Hence, Ωc = h−1 ([c, +∞)) = h−1 ([c, c0 ]) for c ≤ c0 .

If c0 = ĉ0 and f (x) = 1 − e−h(x) is strictly concave at x0 then x0 is the unique

maximum of f and h; hence X n is contractible. Thus, we may assume that ĉ0 =
h(x0 ) < c0 .
There are three possibilities for the maximum set Ωc0 = A0 of h:
Case 1. Ωc0 is a convex and compact subset without boundary and dim(Ωc0 ) ≥
1. In this case, Ωc0 remains an Alexandrov space of non-negative curvature.
Case 2. dim(Ωc0 ) = 0 and X n is contractible.
Case 3. dim(Ωc0 ) ≥ 1 but Ωc0 is a convex subset with non-empty boundary. In this case, we let
Ωc0 +s = {x ∈ Ωc0 | d(x, ∂Ωc0 )} ≥ x}
and consider the distance function r∂Ωc0 (x) = d(x, ∂Ωc0 )

(1.2)

for x ∈ Ωc0 . Since Ωc0 = A0 is compact, the distance function r∂Ωc0 from the boundary has a maximum value on A0 , say l1 = max{r∂Ωc0 (x) | x ∈ Ωc0 }.
5

−1
There are also three possibilities for the maximum subset A1 = Ωc0 +l1 = r∂Ω
(l1 )
c
0

as well.
Since we have dim(X n ) > dim(A0 ) > dim(A1 ) > · · · , repeating above operations at most n times, we end up either Case 1 or Case 2.
Under the assumption that X n has positive curvature on Bε0 (x0 ), we will make efforts to rule out Case 1 above.
Theorem 1.5. Let X n be a complete, non-compact n-dimensional metric space of non-negative curvature and h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

as above. Suppose that X n has no boundary and has positive curvature on a metric ball Bε0 (x0 ). Then the maximum subset A0 = h−1 (c0 ) of h must be either a point or A0 has a (strictly convex) boundary point y0 ∈ ∂A0 in the sense defined below.
The proof of Theorem 1.5 will be given in upcoming sections.
For a convex subset A ⊂ X n , there is sufficient condition for the subset A to have a boundary point.
Definition 1.6. Let X n be an Alexandrov space of curvature ≥ −1.
(1.6.l) If, for any quasi-geodesic σ : [a, b] → X n with ending points {σ(a), σ(b)} ⊂
A, quasi-geodesic segment σ([a, b]) ⊂ A, then A is called a totally convex subset of
X n.

(1.6.2) Suppose that A is a totally convex subset A of X n , p ∈ A and ~u ∈ Tp (X n )

is a unit tangent direction of X n at p; suppose that
∡p (~u, x) ≥

π
2

for any x ∈ A − {p}. Then ~u is called at least normal to A at p.
6

(1.3)

(1.6.3) If ~u ∈ Tp (X n ) and k~uk = 1 then the sub-space
Hu~+ = {y ∈ X n − {p} | ∡p (~u, y)} >

π
2

is called an open half sub-space relative to ~u.
(1.6.4) Suppose that ~u is an at least normal vector to a totally convex subset Ω
at p, and suppose that
Ω
θp,~
u, x)|x ∈ Ω, d(x, p) ≥ r} −
u (r) = inf{∡p (~

π
> 0,
2

(1.4)

for all sufficiently small r > 0. Then the point p is called a strictly convex boundary point of Ω.
When X n has positive curvature on Bε0 (x0 ), we already point out that the function f (x) = 1 −e−h(x) is strictly concave on a smaller ball Bε0 /4 (x0 ). Moreover, we have the following refined estimate.
Proposition 1.7. Let X n , h(x) and Ωc = h−1 ([−c, c0 ]) be as above. If X n has positive curvature ≥ k0 > 0 on Bε0 (x0 ), then for each p ∈ Bε0 /4 (x0 ) there exists a > 0 such that
Ωc
θp,~
u (r) ≥ ar > 0

(1.5)

for some ~u ∈ Tp (X n ) and sufficiently small r > 0, where c = h(p).
Let us now return to the maximum subset A0 = Ωc0 = h−1 (c0 ). For interior points of A0 , we have the following observation.
Ω
Proposition 1.8. (Compare [CG09] §2.2 ) Let X n , h, {Ωc } and θp,~
u (r) be as above.

Suppose that dim(A0 ) ≥ 1 and p is an interior point of A0 . Then
∡p (~u, y) =

π
2

(1.6)

for any at least normal direction ~u to A0 at p and any y ∈ A − {p}. Consequently,
A0
θp,~
u ≡0
7

(1.7)

for all r > 0.

∂Ωc

r

r
θ

p
~u

Figure 1. Strictly convexity and angular excess
In order to establish Theorem 1.5 and hence Theorem 0.1, by (1.7) it is more desirable to show that the inequality (1.5)
Ωc
θp,~
u (r) ≥ ar > 0

for some points p ∈ A0 and Ωc = A0 .
Thus, we are led to study the inequality (1.5) along a Perelman-Sharafutdinov curve

d+ ϕ
∇h
=
.
dt
|∇h|2

It was shown by H.Wu (cf. [Wu79]) that our Busemann function h is indeed a distance from an appropriate subset. In fact, one can show that h(x) = c + d(c, ∂Ωc) = c + r∂Ωc (x),

(1.8)

for all x ∈ Ωc = h−1 ([c, +∞)). Therefore, the semi-gradient semi flow of h is actually semi-gradient flow for distance functions. The evolution induced by PerelmanSharafutdinov semi-flow
∇h
∇r∂Ωc d+ ϕ
=
=
2
dt
|∇h|
|∇r∂Ωc |2

(1.9)

is indeed an one-sided equidistance evolution. To our surprise, the angular estimate of
Ωc
θp,~
u (r) ≥ ar > 0
8

is preserved by Perelman-Sharafutdinov equi-distance evolution {∂Ωc } in a space
X n of non-negative curvature.

Theorem 1.9. (Angular estimates under parallel translation ) Let X n , h(x), Ωc =
Ωc n
h−1 ([c, +∞)) and θp,~
u (r) be as above. Suppose that X has non-negative curvature,

ϕ : [a, b] → X n is a Perelman-Sharafutdinov curve d+ ϕ
∇h
=
dt
|∇h|2
with ci = h(ϕ(ti )), t1 ≤ t2 , ~u is the left-derivative of ϕ at t, and suppose that
Ω

h(ϕ(t))
θt (r) = θϕ(t),~
u(t) (r). Then

for any t1 ≤ t2 . Equivalently,

θt2 (r) ≥ θt1 (r)

(1.10)

∂θt (r)
≥0
∂t

(1.11)

holds.

r

r
ϕ(t2 )

≤ t2 − t1

≤ t2 − t1
θ1
≤ t2 − t1

∂Ωc1

≤ π2 −θ1
ϕ(t1 )

Figure 2. Angular estimates under equidistance evolution
If X n is a smooth Riemannian manifold of non-negative curvature and if ∂Ωci is a smooth convex hypersurface then the inequality
∂θt (r)
≥0
∂t
9

is related to the classical Riccati equation of the smooth second fundamental form of II∂Ωh(ϕ(t)) = IIt :
II′t + II2t + R = 0,

(1.12)

where R(V, Y )Z = −∇V ∇Y Z = ∇Y ∇V Z + ∇[V,Y ] Z is the curvature tensor of the smooth Riemannian manifold X n .

For instance, let Ω ⊂ R2 be a domain given by
Ω = {(x1 , x2 ) | x2 < −x21 }.
−1
Ω
If we choose p = (0, 0) and ~u = (0, 1), then θp,~
r ∼ r for all sufficiently u (r) ∼ tan

small r.

x2

θ

x1
(r, −r2 )

Figure 3.
Proposition 1.10. Let Ω ⊂ M n be a totally convex subdomain with smooth boundary ∂Ω in a smooth Riemannian manifold M n . Then the second fundamental form
II∂Ω of ∂Ω with respect to the outward unit normal vector ~u at p satisfies
IIu~∂Ω (w,
~ w)
~ = −h∇w~ w,
~ ~ui ≥ 2λ1 kwk
~ 2

(1.13)

for some λ1 > 0 if and only if lim

r→0

Ω
θp,~
u (r)

r

for some λ2 > 0.
10

≥ λ2 > 0

(1.14)

In a recent paper [AB09], Alexander and Bishop used the length-excess function for chords in Ω to define extrinsic curvature of ∂Ω in Alexandrov spaces. We will discuss the relation between our angular excess functions lim inf r→0

Ω
θp,~
u (r)
r

and

Alexander-Bishop’s extrinsic curvature in upcoming sections.
Among other things, we will use the following trapezoid comparison theorem to verify Theorem 1.9 above.
q2

σs (t)

r

q1

≤π − β
s sin β

ϕ2 (t)
β
p̂

ϕ1 (t)

α2

≤ s = t2 − t 1

α3
ϕ3 (t)

q3

Figure 4. Trapezoid comparison
Theorem 1.11. (Trapezoid comparison theorem ) Let X n be a complete Alexandrov space of non-negative curvature as above, and let {ϕ1 , ϕ2 , ϕ3 } be three geodesic segments with the same initial point p̂. Suppose that the three initial directions
{ϕ′1 (0), ϕ′2 (0), ϕ′3 (0)} are co-planar in the sense
∡p̂ (ϕ′2 (0), ϕ′1 (0)) + ∡p̂ (ϕ′1 (0), ϕ′3 (0))
= α2 + α3 = β
= ∡p̂ (ϕ′2 (0), ϕ′3 (0)),

(1.15)

l = d(p̂, q1 ), ϕ3 is a possible quasi-geodesic, d(p̂, q2 ) = sins β and σ3 : [0, r] → X n is a geodesic segment from q2 to q1 with
∡q2 σs′ (0),

d− ϕ s 
(
) ≤π−β
dt sin β

(1.16)

as in Figure 4. Then the inequality d(σs (r), ϕ3 (R)) ≤ d(σs (r), ϕ3 (l cos α3 )) ≤ s
11

(1.17)

holds.
We provide the detailed proofs of results stated above in upcoming sections.
§2. Non-negative curvature and weak concavity of Busemann functions
In this section, we will prove Proposition 1.8 which is related to the weak concavity of Busemann functions defined on open complete spaces of non-negative curvature.
We will use the same notation as in Section 1. Let h : X n → R be a Busemann function given by h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

and let Ωc = h−1 ([c, +∞)). In order to derive the desired angular estimates, we recall the totally convex subsets so that we can construct conic-like barrier hypersurfaces.
Proposition 2.1. Let X n be an open, complete and n-dimensional Alexandrov space with non-negative curvature and let h(x) and Ωc = h−1 ([c, +∞)) be as above.
Suppose that σ : [a, b] → X n is a quasi-geodesic of unit speed with {σ(a), σ(b)} ⊂ Ωc .
Then
σ([a, b]) ⊂ Ωc

(2.1)

and Ωc is a totally convex subset of X n .
Proof. The proof of Proposition 2.1 for smooth Riemannian manifold was proved by
Cheeger-Gromoll (cf. [CG72]). For singular spaces, Perelman (cf. [Per91] Chapter
6) also established the similar result (see also Section 2 of [CaG10]). For convenience of readers, we present another simple proof here. By an equivalent definition of nonnegative curvature, for any distance function r(x) = d(x, ŷ) from a given point ŷ, we have
1
Hess( r 2 ) ≤ I
2
12

(2.2)

which means that, for any quasi-geodesic σ of unit speed, we have

d2 1
2
[r(σ(s))]
≤ 1.
ds2 2

(2.3)

It follows that
Hess(r) ≤

1
I.
r

(2.4)

Recall that h(x) = limt→∞ [d(x, ∂Bt(x̂)) − t]. By (2.4) we have
1
1
] = 0.
Hess(h(x)) = lim [Hess( [d(x, ∂Bt(x̂))]2 )] ≤ lim [
t→∞ d(x, ∂Bt(x̂))
t→∞
2
Hence h is a concave function. It follows that h(σ(t)) ≥ min{h(σ(a)), h(σ(b))}
for all t ∈ [a, b]. Therefore, the sup-level set Ωc = h−1 ([c, +∞)) is a totally convex subset of X n . 

We also need to show that the Busemann function h is proper.
Proposition 2.2. (Compare [Wu79] ) Let X n be an open, connected and complete n-dimensional space of non-negative curvature and h(x) = lim [d(x, ∂Bt(x̂)) − t] as t→∞

−1

above. The for each c ∈ R, the sup-level set Ωc = h

([c, +∞)) is compact.

Proof. Let ĉ = min{h(x̂), c} = min{0, c}. It is sufficient to verify that Ωĉ is compact. Suppose to the contrary, Ωĉ were non-compact. There would be un-bounded sequence {qi } ⊂ Ωĉ such that lj = d(x̂, qj ) → +∞

(2.5)

as j → +∞. Let σj : [0, lj ] → X n be a length-minimizing geodesic from x̂ to qj . By passing to a subsequence {yjk } of {yj } if necessary, we may assume that

σj′ (0) → ~v∞ and σj → σ∞ as j → +∞, where σ∞ : [0, +∞) → X n is a geodesic ray from x̂. H.Wu (cf. [Wu79]) observed that h(σ∞ (t)) = h(x̂) − t = −t → −∞ < ĉ
13

(2.6)

as t → +∞. However, we already proved that Ωĉ is totally convex, and hence
σ∞ ([0, +∞)) ⊂ Ωĉ ,

(2.7)

h(σ∞ (t)) ≥ ĉ,

(2.8)

which implies

a contradiction to (2.6). 
By Proposition 2.2, we see that our Busemann function h has a finite maximum value a0 where a0 = maxn {h(x)}.

(2.9)

x∈X

We consider the maximum set A0 = h−1 (a0 ) of h. When dim(A0 ) ≥ 1, we would like to address the “weak concavity property” of A0 = h−1 (a0 ) at its interior points.

Proposition 2.3. (Compare Theorem 2.11(2) in [CaG10] ) Let X n be an open complete n-dimensional space of non-negative curvature, h(x) = lim [d(x, ∂Bt(x0 ))−
t→+∞

t] and a0 = max{h(x)|x ∈ X n }. Suppose that A0 = h−1 (a0 ) has positive dimension dim(A0 ) ≥ 1, x ∈ int(A) is an interior point of A0 , and ~v is a unit vector at least normal to A0 at x. Then
∡x (~v, w)
~ =

π
2

(2.10)

holds for all w
~ ∈ Tx (A).
Proof. Case 1. dim(A0 ) = 1. Since A0 is a totally convex subset of X n , A0 is either a closed geodesic or a length-minimizing geodesic segment. By our assumption, x is an interior point of A0 , there is a geodesic σ : (−ε, ε) → A0 ⊂ X n such that

σ(0) = x. It is known (cf. [BGP92]) that the tangent space Tx (X n ) has an isometric splitting
Tx (X n ) = Y n−1 × Tx (A) = Y n−1 × R.

(2.11)

It follows that, for any ~v ∈ Σx (X n ), we have
π
min{∡x (~v , σ ′ (0)), ∡x(~v , −σ ′ (0))} ≤ .
2
14

(2.12)

If ~v is at least normal to A0 at x, then by (2.12), we must have
∡x (~v , ±σ ′ (0)) =

π
.
2

and ~v ∈ Y n−1 ⊥ Tx (A0 ).
Case 2. When dim(A0 ) > 1, our proof becomes more involved. If Y is an
Alexandrov space with curvature ≥ c and y ∈ Y , then we let Σy (Y ) denote the space of unit tangent directions of Y at y. It is well-known that Σy (Y ) has curvature ≥ 1.

Let A′0 = Σx (A0 ), where A0 = {y ∈ X | h(y) = maxz∈X {h(z)}} is the maximum set of Busemann function h defined on the open space with non-negative curvature.

Using triangle comparison theorem for spaces Σ with curvature ≥ 1 and a result of
Perelman-Petrunin on quasi-geodesics, we will show that if x is an interior point of
A0 , then max{∡x (w,
~ A′0 ) w
~ ∈ Σx (X), ∡x (w,
~ A′0 ) = dΣ (x, A′0 )} ≤

π
.
2

(2.13)

Our strategy to establish (2.13) goes as follows (compare with the proof of Theorem
2.11(2) in [CaG10]).
Claim A. Let A0 be a totally convex and the maximum set of a Busemann function h as above. Suppose that x is an interior point of A0 , A′0 = Σx (A0 ) and dim(A0 ) =
k ≥ 2. Then the following is true:
(A.1) For each ~u ∈ Σξ (A′0 ), there is a quasi-geodesic σv : (−ε, ε) → A′0 such that

σv (0) = ξ~ and σv′ (0) = ~u.

(A.2) dΣ (w,
~ A′0 ) ≤ π2 for all w
~ ∈ Σx (X).
We first verify Assertion (A.1). Recall that A0 is totally convex in X. Therefore,
A′0 = Σx (A0 ) is totally convex in Σx (X). Since x is an interior point of A0 , the subspace A′0 = Σx (A0 ) has no boundary in Σx (X) with dim(A′0 ) = k − 1 ≥ 1.

Hence, A′0 is a compact, totally convex subspace without boundary in Σx (X). For each ξ ∈ A′0 , we let
ρξ : Σx (X) −→ R
w
~ −→ ∡x (ξ, w) = dΣ (ξ, w).
15

Since A′0 is a totally convex subspace without boundary in Σx (X), the semi-gradient
+

curves ϕ : [0, l) → Σx (X) of ddtϕ = ∇ρξ with ϕ(0) = ξ and ϕ′ (0) = ~u will remain

in A′0 . Moreover, the gradient exponential map at ξ, gexpξ : Tξ (Σx (X)) → Σx (X), has the property gexpξ (Tξ (A′0 )) ⊂ A′0 , because A′0 is totally convex in Σx (X) and because (Tξ (Σx (X)), oξ ) = lim (λΣx (X), ξ). It follows from a theorem of Perelmanλ→∞

Petrunin (cf. [Petr07]) that quasi-geodesics can be approximated by broken gradient exponential curves, (cf. Appendix of [Petr07]). Thus, by the construction of quasigeodesics given by Perelman-Petrunin (cf. Appendix of [Petr07]), we see that any quasi-geodesic
σ : [0, l] → Σx (X)
with σ(0) = ξ and σ ′ (0) = ~u ∈ Σξ (A′0 ) will stay in A′0 ⊂ Σx (X). This completes the proof of our Assertion (A.1).
For Assertion (A.2), we use a triangle comparison theorem for the space Σx (X).
Let ϕ : [0, l] → Σx (X ′ ) be a length-minimizing geodesic segment from A′0 to w
~ ∈

Σx (X) of unit speed. Suppose that ϕ(0) = ξ ∈ A′0 , ϕ(l) = w and dΣ (A′0 , w) = d(ξ, w) = l.
We will use comparison theorem to show that dΣ (A′0 , w) = d(ξ, w) = l ≤

π
.
2

(2.14)

Because dim(A′0 ) = k − 1 ≥ 1, we can choose a quasi-geodesic σ : (−ε, ε) → A′0 ⊂

′
Σx (X) such that σ(0) = ξ. Let σ±
(0) = ~u± be the left (or right) derivative of σ at

s = 0. Perelman and Petrunin (cf. [Petr07]) showed that h~u+ , ϕ′ (0)i + h~u− , ϕ′ (0)i ≥ 0.
It follows that
π
α = min{∡ξ (~u+ , ϕ′ (0)), ∡ξ (~u− , ϕ′ (0))} ≤ .
2
16

We may assume that ~u+ = σ ′ (0) has the property
′
α = ∡ξ (~u+ , ϕ′ (0)) = ∡ξ (σ+
(0), ϕ′ (0)) ≤

π
.
2

If l ≤ π2 , we are done. If l > π2 , then we will get a contradiction as follows. For geodesic hinge {σ, ϕ} with the vertex ξ and angle α ≤ π2 , since Σx (X) has curvature

≥ 1, we have

π
π
dΣ (ϕ( ), σ(ε)) ≤ .
2
2

(2.15)

Recall that ϕ : [0, l] → Σx (X) is a length-minimizing geodesic from A′0 , we also have
π
π
π
dΣ (ϕ( ), σ(ε)) ≥ dΣ (ϕ( ), A′0 ) ≥ .
2
2
2
Combing with (2.15), we see that
π
π
dΣ (ϕ( ), σ(ε)) = .
2
2
Therefore, we have two distinct points {ξ, σ(ε)} ⊂ A′0 such that
π
π
π
π
d(ξ, ϕ( )) = d(σ(ε), ϕ( )) = d(A′0 , ϕ( )) = .
2
2
2
2
Thus, ϕ|[0, π2 +ε′ ] is no longer a length-minimizing geodesic from A′0 for sufficiently small ε′ where 0 < ε′ ≤ [l − π2 ], a contradiction. Thus we established l = d(A′0 , w) ≤

π
2

for all w ∈ Σx (X). This completes the proof of our Assertion (A.2) as well as
Proposition 2.3. 
In next section, we discuss the relation between strictly concavity of h and positive curvature on the small ball Bε0 (x̂).
17

§3. Strictly positive curvature and strong concavity of Busemann function on a small region
By our discussions in §2 above, we see that if our Busemann function h is partially

strong concave on a portion of the maximum set A0 = h−1 (a0 ), then either A0 is a single point set or A0 has non-empty boundary.

In order to establish the partial strong concavity of h on a portion of the maximum set A0 , we begin with a small ball Bε0 /4 (x) where the curvature of X n is strictly positive.
Proposition 3.1. Let X n be an open complete Alexandrov space of non-negative curvature. Suppose that X n has curvature ≥ 1 on B2ε0 (p0 ), h(x) = lim [d(x, ∂Bt(x̂)) − t]
t→∞

and c0 = h(x̂). Then, for any p ∈ Bε0 (x̂), there exists ~u with
Ω

θp,~h(p)
u (r) ≥

ε0
r
8

(3.1)

for sufficiently small 0 < r ≤ r0 .
Proof. We will use a result of Petrunin ([Petr07]) to derive the desired estimate.
By our discussion, h(x) = (c0 − 2ε0 ) + d(x, ∂Ωc0−2ε0 )
∂Ω

for x ∈ B2ε0 (p0 ). We choose ~v ∈⇑p0 c0 −2ε0 and σ : [0, 2ε0 ] → X n is a lengthminimizing geodesic from p0 to ∂Ωc0 −2ε0 with σ ′ (0) = ~v . Let q0 = σ(ε0 ), w
~ ∈

Σp0 (X n ) such that ∡p0 (~v, w)
~ = π2 , and ψ : [0, r0 ] → X n be a quasi-geodesic from p0

with ψ ′ (0) = w.
~ Petrunin ([Petr07]) proved that d(ψ(r), ∂Ωc0−ε ) is bounded above by the 2-dimensional model case as follows. Let
S 2 = {(x1 , x2 , x3 ) ∈ R3 |x21 + x22 + x23 = 1}, S 1 = {(x1 , x2 , 0)|x21 + x22 = 1},
18

p∗0 = (cos ε0 , 0, sin ε0 ) and q0∗ = (1, 0, 0). The spherical geodesic

ψ ∗ (r) = (cos r)p0 + (sin r)z0 = (cos ε0 ) cos r, sin r, (sin ε0 ) cos r .
A calculation shows that

d(ψ ∗ (r), S 1 ) = tan−1 q

(sin ε0 ) cos r
(cos ε0 )2 (cos r)2 + sin2 r

≤ ε0 cos r ≤ ε0 (1 −



r4
r2
+ ).
2
24

(3.2)

Petrunin ([Petr07]) showed that d(ψ(r), ∂Ωc0−ε0 ) ≤ dS 2 (ψ ∗ (r), S 1 ) + o(r 2 ) ≤ ε0 (1 −

r2
r4
+ ) + o(r 2 ),
2
24

(3.3)

2

)
where lim o(r r2 = 0. Therefore, for sufficiently small r < r0 , we have r→0

d(ψ(r), ∂Ωc0−ε0 ) ≤ ε0 (1 −
It follows that h(ψ(r)) ≤ c0 −

r2
).
4

(3.4)

ε0 r 2
.
4
2

Thus for each y ∈ [Ωc0 − Br (p0 )] with d(y, p0 ) = r, we have d(y, ψ(r)) ≥ ε04r . By comparison theorem for spaces with curvature ≥ 0, we see that
∡p0 (y, ψ ′ (0)) ≥

ε0 r 2
ε0
≥ r.
4r
8

(3.5)

ε0
r
8

(3.6)

It follows that
Ω

θp0c,~0v (r) ≥
for sufficiently small r < r0 . 

In next section, we will discuss a sufficient condition so that the inequality
Ω

θp,~h(p)
u (r) > c0 r holds for some points p ∈ A0 = h−1 (a0 ).
19

§4. Preserving concavity of hypersurfaces under equi-distance evolution
In previous section, we showed that if X n has positive curvature on B2ε0 (x̂)
then Bε0 (x̂) ∩ ∂Ωh(x) is strictly concave. Our goal of this section is to show that the strictly concavity property of {∂Ωh(ϕ(t)) } is preserved along equi-distance evolution.
More precisely, if ϕ : [a, b] → X n is a Perelman-Sharafutdinov curve for a distance

function (or a Busemann function) and if ∂Ωh(ϕ(0)) is strictly concave at ϕ(0) then
∂Ωh(ϕ(t)) remains strictly concave at ϕ(t) for all t ≥ 0. Let us discuss several examples to motivate Theorem 1.9.
Example 4.1 (a) We consider the following domain
Ω0 = {(x1 , x2 ) ∈ R2 | |x1 | ≤ 4, |x2 | ≤ 1, or |x1 | ≥ 4,

1
(x1 ± 4)2 + x22 ≤ 1},
4

see Figure 5.
x2

x1
A0

Figure 5. two-step soul
We observe that l0 = max{d(p, ∂Ω0 ) | p ∈ Ω0 } = 1.
Let Ωs = {q ∈ Ω0 | d(q, ∂Ω0) ≥ s} and r(q) = d(q, ∂Ω0 ). The maximum subset

A0 = r −1 (1) is an interval, i.e., A0 = {(x1 , 0) | |x1| ≤ 4}. The soul of Ω0 is a single point set A1 = {(0, 0)}. The level set ∂Ωs has strict convex points in (∂Ωs )∩[R×{0}]
for 0 ≤ s < 1.

20

(b) We consider a family of rectangles {Ωs }, where

Ωs = {(x1 , x2 ) | |x1 | ≤ 2 − s, |x2 | ≤ 1 − s}

for 0 ≤ s ≤ 1.
ϕ(0)
ϕ(s)
Ωs

Ω0

Figure 6. Corner points and equi-distance evolution
The corner point qs = (2 −s, 1 −s) is a strictly convex point of ∂Ωs for 0 ≤ s ≤ 1.
It is clear that corner points are preserved by Perelman-Sharafutdinov semi-flow of distance functions.
(c) (Fermi coordinates and Riccati equation) Suppose that Ω0 is a convex domain with a smooth boundary ∂Ω0 in a smooth Riemannian manifold M n with nonnegative curvature. Suppose that ϕ : [0, l] → Ω0 is a length-minimizing geodesic segment of unit speed from ∂Ω0 with

d(ϕ(s), ∂Ω0 ) = s

for 0 ≤ s ≤ l.

21

(4.2)

ϕ(s)
∂Ωs
ρ

∂Ω0

ρ

Figure 7. fermi coordinates around a curve
~ j (t)}n be a parallel orthonormal frame along the geodesic segment ϕ
Let {E
j=1

~ n (t) = ϕ′ (t). If ∂Ωt is convex and if ϕ′ (t) is an inward unit normal such that E
n−1
] where vector, then ∂Ω0 lies at one side of [M n − Hϕ(t)


n−1
~
~
| E(t)
⊥ ϕ′ (t)}.
= {Expϕ(t) ρE(t)
Hϕ(t)
We consider the corresponding second fundamental form
II∂Ωt (X, Y ) = −hϕ′ (t), ∇X Y i for {X, Y } ∈ Tϕ(t) (∂Ωt ). If
~ i (t), E
~ j (t))
uij (t) = II∂Ωt (E

and if u(t) = uij (t) (n−1)×(n−1) is a matrix-valued representation of II∂Ωt then the
Riccati equation

u′ + u2 + R = 0
holds, where
~ i (t))ϕ′ (t), E
~ j (t)i,
Rij (t) = hR(ϕ′ (t), E

R(t) = Rij (t) (n−1)×(n−1) is the curvature matrix function and
R(X, Y )Z = −∇X ∇Y Z + ∇Y ∇X Z + ∇[X,Y ] Z.
22

(4.2)

When M n has non-negative sectional curvature, the matrix R(t) is positive semidefinitive:
R(t) ≥ 0.

(4.4)

u′ (t) = −u2 (t) − R(t) ≤ 0

(4.5)

It follows that

and hence the second fundamental form
II∂Ωt (E(t), E(t)) = hu(t)E(t), E(t)i is a monotone function of t for any parallel vector field {E(t)} along ϕ. This gives rise to a version of Theorem 1.9 for this special case.



If ϕ : [a, b] → X n is a Perelman-Sharafutdinov curve of a distance function (not

necessarily a length-minimizing geodesic) in a singular space X n of non-negative curvature, there is no Riccati equation available, we will use Trapezoid Comparison
Theorem (cf. Theorem 1.11 above) instead.
§4.1. Angular estimates under equi-distance evolution in dimension 2
In this sub-section, we provide a proof of Theorem 1.9 for the case when X n has dimension 2. We also derive some preliminary results for all dimensions.
Let us first recall the definition of barrier functions and barrier hypersurfaces, so that we can estimates the concavity of Busemann function h and its level sets
∂Ωc = h−1 (c).
Definition 4.2. (Calabi [Ca57] ) Let h : X n → Rn be a continuous function. We

say that Hess(h)(p) ≤ λ if for any quasi-geodesic σ : (−r, r) → X n with σ(0) = p there exists an upper barrier function ĥ such that (ĥ ◦ σ)(t) is smooth in t, h(σ(0)) = ĥ(σ(0)), h(σ(t)) ≤ ĥ(σ(t))
and

d2 [ĥ(σ(t))]
≤λ
dt2
t=0
23

hold.
We remark that, in above definition, we need to choose two-sided barrier functions instead of one-sided barrier functions to estimate the second derivative.
For example, if f (t) is a smooth function of t, then we have the Taylor expansion
1
f (t) = f (0) + f ′ (0)t + f ′′ (0)t2 + o(t2 )
2
2

where lim o(tt2 ) = 0. It follows that t→0

f (t) + f (−t) − 2f (0)
t→0
t2
(−t)
f (t)−f (0)
− f (0)−f t
t
= lim
.
t→0
t

f ′′ (0) = lim

The second derivative f ′′ (0) measures the rate of change for slopes
 f (t) − f (0) f (0) − f (−t) 
∼ λt
−
t t
up to the first order.
The geodesic curvature of a curve ψ : [−r, r] → ∂Ω0 in a smooth hypersurface
∂Ω0 of a smooth Riemannian manifold is given by
~i
II∂Ω0 (ψ ′ (0), ψ ′ (0)) = −h∇ψ ′ ψ ′ , ϕ′ (0)i = −hψ ′ , ∇ψ ′ N
t=0
~ is a smooth unit normal vector of ∂Ω0 . The geometric quantity −hψ ′ , ∇ψ ′ N
~i where N
measure changes of slope as well.
We consider another example to indicate why we need to choose two sided barrier function. Let us consider the function
1
h : R2 → R2 , (x1 , x2 ) 7→ √ (x2 − |x1 |).
2
24

x2

h−1 (0)

45◦

x1

Figure 8. non-smooth Busemann function
Clearly, h is a piecewise linear function. Hence, h(x1 , x2 ) has Hess(h)|p = 0 if p = (x1 , x2 ) satisfies x1 6= 0. However, at p = (0, 0), one can verify that
Hess(h)|(0,0) = −∞
in barrier sense.
By our definition, p = (0, 0) is a strictly convex point of h−1 (0), where Ω0 =
h−1 ([0, +∞)). In this example, x1 -axis lies below the level curve h−1 (0). We need to derive some preliminary results, in order to construct “supporting cones” of h−1 (c) outside a metric ball Br (p).
Proposition 4.3. (Compare with [CG72] ) Let X n , h(x) and {Ωs } be as above.

Suppose that σ : (−∞, +∞) → X n is a quasi-geodesic of unit speed with {σ(a), σ(b)} ⊂

∂Ωs = h−1 (s) and a < b. Then the following hold.

(4.3.a) If max {h(σ(t))} > s, then h(σ(t)) < s for any t ∈
/ [a, b]; a≤t≤b

(4.3.b) If max {h(σ(t))} = s, then h(σ(t)) ≤ s for all t.
a≤t≤b

Consequently, if σ̃ : [0, +∞) → Ωc is a quasi-geodesic segment with {σ̃(0), σ̃(r)} ⊂
∂Ωc and r > 0, then
σ̃([r, +∞)) ⊂ [X n − Int(Ωc )] = h−1 ((−∞, c]).
Proof. Let η(t) = h(σ(t)). by our assumption, we know that η(t) is a concave function. Our conclusion are direct consequences of concavity of η(t). 
25

η

t a

t0

b

Figure 9
On a singular space X with curvature ≥ −1, the extension of quasi-geodesic

segment σ : [a, b] → X n to a “longer” quasi-geodesic σ̃ : R → X n is not necessarily

unique. Thus, it is necessary for us to recall the notion of the gradient exponential map (cf. [Petr07]).
Such a gradient exponential map Expp : Tp X → X is related to the gradient flow of the distance function r(x) = d(p, x) which we now describe.
If X is an Alexandrov space with metric d,then we denote by λX the space
(X, λd). Let iλ : λX → X be the canonical map. The Gromov-Hausdorff limit of pointed spaces {(λX, x)} as λ → +∞ is the tangent cone of (Tx (X), ox) of X
at the point x, (see §7.8.1 of [BGP92]. For any λ-concave function, the function dx f : Tx (X) → R defined by

f ◦ iλ − f (x)
λ→+∞
1/λ

dx f = lim is called the differential of f at x.

The gradient vector ∇f of f at x is related to the inner product of Tx (X). For any pair of vectors ~u and ~v in Tx (X), we define
1
(|~u|2 + |~v |2 − |~u~v |2 ) = |~u||~v | cos θ
2
where θ is the angle between ~u and ~v , |~u~v| = dTx (X) (~u, ~v), |~u| = dTx (X) (~u, o), and o h~u, ~v i =

denotes the origin of the tangent cone. It follows that h~u, ~vi cos θ =
.
|~u||~v|
26

Definition 4.4. ([Per91], [petr07] ) For any given semi-concave function f on X, a vector ~η ∈ Tx (X) is called a gradient of f at x (η = ∇f in short) if the following holds:
(i) dx f (~v) ≤ h~η, ~v i for any ~v ∈ Tx (X);
(ii) dx f (~η ) = |~η |2 .

It is known that any semi-concave function has a uniquely defined gradient vector field. Moreover, if dx f (~v) ≤ 0 for all ~v ∈ Tx (X), then ∇f |x = 0. In this case, x is called a critical point of f . Otherwise, we set
~ ξ~
∇f = (dx f )(ξ)
where ξ~ is the (necessarily unique) unit vector for which dx f attains its positive maximum on Σx (X) and Σx (X) is the space of directions of X at x.
When X n has curvature ≥ 0, its energy function f (x) = 21 [d(x, q)]2 is 1-concave.
We consider the semi-flow
Φtf (q) = αq (t)
where αq (0) = q and

d+ αq (t)
= ∇f α (t) .
q dt

Recall that lim (λX, q) = (Tq (X), oq ).

λ→+∞

We define gexpq : Tq (X) → X by gexpq = lim Φtf ◦ iet , t→+∞

where iλ : λX → X is the canonical map. In fact, for each unit direction ξ~ ∈ Σp (X),

~ satisfies the equation the radial curve αξ~ : t 7→ gexpp (tξ)
d+ αξ~
dt

=

r̂p (αξ~(t))
t

where r̂p (x) = d(x, p).
27

∇r̂p

Theorem 4.5. ([PP94], [Petr07] p.152 ) Let X, r̂p and gexpp be as above. If
ξ~ ∈ Tp (X) and if h is a concave function, then
~ ≤ h(p) + t(dp h)(ξ)
~
h(gexpp (tξ))
for all t ≥ 0.
~ Thus, we need to recall
In our Theorem 4.5, there is a first order term (dp h)(ξ).
the first variational formula.
Theorem 4.6. ([BBI01], p.125 ) Let X n be a complete Alexandrov space of nonnegative curvature and fA (x) = d(x, A), where A is closed subset of X. Suppose that x ∈
/ A, ⇑A
u| = 1 and x is the set of minimal directions from x to A, |~
θ = min{∡x (~u, w)
~ |w
~ ∈⇑A
x }.
Then dfA (~u) = − cos θ.
Among other things, we need to use Theorem 1.11 to verify Theorem 1.9.
Proof of Trapezoid Comparison Theorem 1.11. Let us make some observations on
Figure 3. When curvature is non-negative, the (quasi)-geodesic triangle with vertices {q̂, p2 , p3 } satisfies


d(p2 , ϕ3 (l cos α2 ))

2

≤ [ℓ2 + (ℓ cos α2 )2 − 2ℓ2 (cos α2 )2 ] = (l sin α2 )2 .

In order to complete our proof, it is sufficient to estimate both l and α2 from above.
Since X n has non-negative curvature, by the triangle comparison theorem we have l2 = [d(q̂, p2 )]2 ≤
=

s2
(sin β̂)2

s 2

sin β̂

+ r2 − 2

rs sin β̂

cos(π − β̂)

+ r 2 + 2rs cot β̂

= ˆl2 .

(4.6)
28

Using Euclidean trigonometry, for the comparison triangle △q̂p̂2 p̂3 with vertices

{q̂, p̂2 , p̂3 } in R2 (see Figure 10 below), we have the inequality


2
d(p̂2 , ϕ̂3 (ˆl cos α̂2 )) = ˆl2 + (ˆl cos α̂2 )2 − 2ˆl2 (cos α̂2 )2
= (ˆl sin α̂2 )2 ≤ s2 ,

(4.7)

α̂2 = β̂ − α̂1 .

(4.8)

where

p̂1

r

p̂2

π − β̂
s sin β̂

ˆl

≤s

ϕ̂3 (ˆl cos α̂2 )

p̂3

α̂1
α̂2
q̂

Figure 10
Let us now consider the Euclidean triangle △q̂p̂1 p̂2 of the side lengths {ˆl, r, sins β̂ }.

Since X n has non-negative curvature, △q̂p1 p2 is fatter than △q̂p̂1 p̂2 . It follows that
α1 ≥ α̂1 and α2 = β̂ − α1 ≤ β̂ − α̂1 = α̂2 .

(4.9)

Therefore, by (4.6) − (4.9) we conclude that l ≤ ˆl, α2 ≤ α̂2 and


d(p2 , ϕ3 (l cos α2 ))

2

≤ (l sin α2 )2 ≤ (ˆl sin α̂2 )2 ≤ s2 .

This completes the proof of Theorem 1.11. 
We conclude this subsection by a special case of Theorem 1.9 when dim(X n ) = 2.
Theorem 4.7. Let X 2 be an open Alexandrov surface with non-negative curvature, h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

29

and Ωs = h−1 ([s, +∞)). Suppose that ϕ : [0, l] → X 2 is a Perelman-Sharafutdinov curve for h defined by

∇h d+ ϕ
,
=
dt
|∇h|2 ϕ

Ω

−

h(ϕ(s))
and η(s, r) = θϕ(s),~
u(s) = ddsϕ is the normalized left derivative of ϕ.
u(s) (r) where ~

Then
∂η
(s, r) ≥ 0.
∂s

(4.10)

Proof. If Ωc is a totally convex subset of X 2 and if dim(Ωc ) = 2, then Perelman
(cf. [Per91], Chapter 6) calculated ∇h as follows.
∂Ωc+s

r
≤s

z
≤s

xs
≤ π−β0

∂Ωc

z∗

β0
x∗ π +θ
2

0

Figure 11
If x ∈ ∂Ωc and β = 12 diam(Σx (Ωc )), then Perelman (cf. [Per91] page 33, line 1)
observed that
|∇h(x)| = sin[β(x)].

(4.11)

Ω

h(ϕ(s))
For fixed r, we would like to show η(s, r) = θϕ(s)
(r) is an increasing function in

s.
Perelman (cf. [Per91], Chapter 6) already proved that the Sharafutdinov retraction is distance non-increasing. Hence the map
π : ∂Ωc → ∂Ωc+ε
30

is a distance non-increasing map.
If x, y ∈ ∂Ωc+ε , x∗ ∈ π −1 (x) and y ∗ ∈ π −1 (y), then d(x∗ , y ∗ ) ≥ d(x, y).
Since X 2 has non-negative curvature, it is known that diam[Σ1x (X 2 )] ≤ π.
Ω

∗

If x∗ = ϕ(0) and θ0 (r) = θx∗h(x ) (r), then
β0 (r) ≤ π − (

π
π
+ θ0 ) = − θ0 .
2
2

(4.12)

Let xs = ϕ(s). We consider the left derivative ~vs of ϕ at xs . Since xs is the nearest point on ∂Ωh(xs ) to x0 , the vector ~vs is at least normal to ∂Ωh(xs ) at xs . By our assumption, we have
∡x∗ (xs , [∂Ωh(x∗ ) − Br (x∗ )]) ≤

π
∗
Ω
− θx∗h(x ) (r) + o(s),
2

(4.13)

where lim o(s)
s = 0.
s→0

We now consider any given quasi-geodesic σ0 : [0, +∞) → X 2 such that σ0 (0) =

x∗ ∈ ∂Ωc and σ0 (l) ∈ ∂Ωc with l > 0,

d(x∗ , σ0 (l)) ≤ r.

(4.14)

It follows from Proposition 4.3 that
σ0 ([l, +∞)) ⊂ h−1 ((−∞, c∗ ]) = [X 2 − int(Ωc∗ )].
By our discussion above, we have
β0∗ = ∡x∗ (xs , σ0′ (0)) ≤

π
− θxΩ∗c∗ (r) + o(s).
2

(4.15)

Our technical goal is the following
Claim 4.7a. Let X 2 , h, {Ωc }, θxΩ (x), σ0 and β0∗ be as above (see Figure 11).

r
Suppose that 0 < s ≤ 32
and that σs : [0, +∞) → X 2 is a quasi-geodesic with

σs (0) = xs , σs (l) ∈ [Ωc∗ − Br (xs )] ∩ U r8 (σ0 ([l, +∞))) and
π
α = ∡xs (x∗ , σs′ (0)) ≤ π − β0∗ = + θs∗ .
2
31

(4.16)

Then
σs ([l, +∞)) ⊂ Us (σ0 ([l, +∞))) ⊂ h−1 ((−∞, c∗ + s]).

(4.17)

If Claim 4.7a is true, then by (4.16)-(4.17), whenever y ∈ [Ωc∗ +s − Br (xs )] ⊂

h−1 ((−∞, c∗ + s]), we must have
Ω ∗

θxsc +s (r) ≥ π − β0∗ =

π
+ θx∗ + o(s),
2

(4.18)

Ω ∗

c +s where lims→0 o(s)
(r), then s = 0. It follows that if η(s, r) = θxs

∂η(s, r)
≥0
∂s

(4.19)

and Theorem 4.7 holds.
It remains to verify Claim 4.7a. Inspired by Perelman’s formula, |∇h|(x∗ ) ∼
sin(β̂x∗ ), where
~ Tx (∂Ωc∗ ))}.
β̂x∗ = max{∡x∗ (w,
For a fixed β0 (r), we choose a geodesic segment ψ : [0, ls ] → Ωc∗ of unit speed with
ψ(0) = x∗ , ∡x∗ (σ0′ (0), ψ ′ (0)) = β0 (r), ls = sinsβ0 and xs = ψ(ls ).

Recall that X 2 has non-negative curvature, using triangle comparison theorem, we have s2
1
+ s2 (cot β0 )2 − 2s2 cot β0
cos β0
2
(sin β0 )
sin β0
1 2
= s2 [(
) − (cot β0 )2 ]
sin β0
= s2 .

[d(xs , σ0 (s cot β0 ))]2 ≤

It follows that d(xs , σ0 (R)) ≤ s.
Let us choose y ∗ ∈ σ0 (R) such that d(xs , y ∗ ) = d(xs , σ0 (R)).
32

(4.20)

Since X 2 has non-negative curvature, the sum of interior angles of any geodesic triangle is greater than or equal to π. It follows that
∡xs (x∗ , y ∗ ) ≥ π − (

π
π
+ β0 ) = − β0 .
2
2

For each z ∈ [Ωc∗ +s − Br (xs )] ∩ U r8 (σ0 ([ls , +∞))), there are two cases.
If ∡xs (z, y ∗ ) ≥ π2 , then
∡xs (z, x∗ ) ≥ π − β0 =

π
+ θ∗,
2

and we are done.
If ∡xs (z, y ∗ ) < π2 , we get a contradiction as follows. Let z ∗ ∈ σ0 ((0, +∞))

satisfy d(z, z ∗ ) = d(z, σ0 ((0, +∞))). Since the Sharafutdinov flow is a distance non-increasing map, we must have d(x∗ , z ∗ ) ≥ d(xs , z) ≥ r.
By Proposition 4.3, we have z ∗ 6∈ Int(Ωc∗ ). It follows that h(z ∗ ) ≤ c∗ .

(4.21)

When ∡xs (z, y ∗ ) < π2 , by Theorem 1.11, we would have d(z, z ∗ ) < s.

(4.22)

This together (4.21) would imply that h(z) < c∗ + s

(4.23)

which contradicts to z ∈ ∂Ωc∗ +s . This completes the proof of Claim 4.7a and hence
Theorem 4.7. 
33

§4.2. Parallel supporting cones outside a
ρ-tube and generalized Fermi coordinates
In this subsection, we elaborate the proof of Theorem 4.7 and extend it to the higher dimensional case with necessary modifications.
∂Ωt2

u
≤ t2 −t1

σ3

p1 = ϕ(t2 ) ≤ π− β̂
u∗

r

p2
≤ t2 −t1
∂Ωt1
p3

β̂ α1α

2

q̂ = ϕ(t1 )

Figure 12. Parallel supporting cones outside r-tubes
Let us first make sure hat for each p ∈ X n , there exists r0 such that any quasigeodesic segments σ : [0, +∞) → X n starting from p will leave Br0 (p), i.e., σ(t) ∈
/
Br0 (p) for relative large t.
Proposition 4.8. (Perelman [Per94b] ) Let X n be a complete Alexandrov space with curvature ≥ −1. For each x̂ ∈ X n , there exist positive numbers λ, δ and a strictly concave function f : Bδ (p) → (−∞, 0] such that
(i) f (p) = 0;
δ
(ii) B λε (p) ⊂ f −1 ([−ε, 0]) ⊂ Bλε (p) for ε < 4λ
, where λ and δ depend on p.

Corollary 4.9. Let X n be a complete Alexandrov space with curvature ≥ −1. Suppose that p ∈ X n , λ, δ and f be as in Proposition 4.8 above. Then any quasi-geodesic
ψ : [0, +∞) → X n with ψ(0) = p must leave B δ (p) when t is sufficiently large.
λ

Proof. Since f is strictly concave on Bδ (p) and p is a unique maximum point of f ,
34

there exists ε0 > 0 such that
−ε1 = f (ψ(ε0 )) < f (ψ(0)) = 0.
It follows from the concavity of f that for t > ε0 > 0, we have f (ψ(t)) − f (ψ(ε0 ))
f (ψ(ε0 )) − 0
≤
t − ε0
ε0 − 0
and hence f (ψ(t)) ≤

ε1
f (ψ(ε0 ))
(t − ε0 ) = − (t − ε0 ).
ε0
ε0

Therefore, for sufficiently large t, we have
ψ(t) ∈
/ f −1 ([−ε0 , 0]).
Recall that B ελ0 (p) ⊂ f −1 ([−ε0 , 0]). Hence we have ψ(t) ∈
/ B ελ0 (p) for sufficiently large t. 
Ωc
π
We now discuss the supporting hypersurface when r = 0 and θp,~
u (0) = 2 .

Theorem 4.10. ([Per91] Chapter 6, [Petr07] p.156, [CDM09] p.40 ) Let X n be an open complete Alexandrov space with non-negative curvature, h(x) = lim [d(x, ∂Bt(x0 )) − t]
t→+∞

∂Ω

and Ωc = h−1 ([c, +∞)). Suppose that x̂ ∈ Ωĉ , ~v ∈⇑x̂ ĉ−δ , and ε is given by Proposin−1
tion 4.8. We consider Hx̂,ε
= {y ∈ Bε (x̂)|σ : [0, l] → Bε (x̂) is a quasi-geodesic, σ(0) =
n−1
x̂, σ ′ (0) ⊥ ~v , l < ε, y = σ(l)}. Then ∂Ωĉ ∩ Bε (x̂) lies above Hx̂,ε
(i.e., h(y) ≤ ĉ
n−1
for y ∈ Hx̂,ε
.)

Proof. It is known (cf. [Wu79])that h(y) = d(y, ∂Ωĉ−δ ) + ĉ for y ∈ Ωĉ−δ .
35

ϕ
∂Ωc
ϕ(t)

r

Hx̂,ε

~v

Figure 13. Supporting hypersurfaces with r = 0

Perelman also showed that if N = ∂Ωĉ−δ , then rN (y) = d(y, N ) is concave for y ∈ Ωĉ−δ . For any quasi-geodesic ψ : [0, +∞) → X n with ψ(0) = x̂, ψ ′ (0) = w,
~ and

∡x (~v, w)
~ ≤

π
,
2

(4.24)

we have d[h(ψ(t))]
≤ 0.
dt t=0
By an equivalent definition of quasi-geodesics, t 7→ h(ψ(t)) is a concave function.
Therefore, we have h(ψ(t)) ≤ h(ψ(0)) = ĉ. 

We now consider supporting cones of Ωc outside the metric ball Br (p) related to
Ωc angular excess function θp,~
u (r) > 0 for r > 0.
36

∂Ωc+ε

ε
β r

r p

∂Ωc

σ([r, +∞))

~u
Figure 14. Supporting cones outside the ball Br (p)
In a smooth Riemannian manifold M n , its tangent cone Tp (M n ) is isometric to the Euclidean space Rn . Hence, any unit vector ~u has its anti-podal vector −~u.

For singular space X n with curvature ≥ −1, we need to recall the notion of polar vectors, so that we can discuss relations between inner angles and outer angles.
Definition 4.11. (i) Two vectors ~u, ~v ∈ Tp (X) (not necessarily the same length)
are called polar if for any vector w
~ ∈ Tp (X), the inequality h~u, wi
~ + h~v , wi
~ ≥0
holds.
(ii) A vector w
~ ∈ Tp (X n ) is called a supporting vector of a semi-concave function f at p if
~ ≤ −hξ,
~ wi
(dp f (ξ))
~
holds for any ξ~ ∈ Tp (X n ).
It is known that the set of supporting vectors of f at p form a non-empty convex subset of Tp (X n ), see [Petr07] p143. In fact, if r(p) = d(p, A) with p ∈
/ A and if w
~ ∈⇑A
~ is a supporting vector of rA at p, by the first variational formula.
p , then w

Using the existence of quasi-geodesics, Petrunin (cf. [Petr07] p194) showed that for unit vector ~u ∈ Σp (X n ), there exists a polar unit vector ~u∗ ∈ Σp (X n ) such that h~u, wi
~ + h~u∗ , wi
~ ≥0
37

for any w
~ ∈ Tp (X n ). Consequently
∡p (~u, w)
~ + ∡p (~u∗ , w)
~ ≤π

(4.25)

for any w.
~
We now ready to construct the supporting cone of Ωc outside Br (p). Recall that if p ∈ ∂Ωc and ~u is a unit vector at least normal to Ωc at p then
Ωc
θp,~
u, x)|x ∈ [Ωc − Br (p)]} −
u (r) = inf{∡p (~

π
.
2

Ωc
Proposition 4.12. Let X n , h, Ωc = h−1 ([c, +∞)), p ∈ Ωc and θp,~
u (r) be as above.

Ωc
π
n
Suppose that r > 0, θ = θp,~
u (r) > 0, β = 2 − θ and σ : [0, l] → X is geodesic

segment with
∡p (~u∗ , σ ′ (0)) ≥ β

(4.26)

where {~u, ~u∗ } are polar vectors in Tp (X n ). Then
σ([r, +∞)) ⊂ [X n − Int(Ωc )] = h−1 ((−∞, c]).

(4.27)

Consequently, σ([r, +∞)) lies below ∂Ωc relative to the vector ~u∗ .
Proof. This is a direct consequence of concavity for Busemann function h.

∂Ωc+ε

≤s

ϕ(t2 ) ≤π−β ≤ s
~u∗ β
ϕ(t1 )

∂Ωc

σ0 ([r∗ , +∞))

~u

Figure 15. Parallel cones along a curve
38



Proof of Theorem 1.9. Let ϕ̃ : [0, l] → X n be a gradient exponential radius curve given by
ϕ̃(s) = gexpp (

s ∗
u ), sin β

Ωc
π
where β = π2 − θp,~
u (r) < 2 .

Let us consider a conic hypersurface
Hϕ(0),β (r) = {σ0 (t) | d(p, σ0 (t)) ≥ r, σ0 is a quasi-geodesic with σ0 (0) = p and ∡p (~u∗ , σ0′ (0)) = β}.
It follows from Proposition 4.12 that the conic hypersurface Hϕ(0),β (r) lies below
∂Ωc relative direction ~u∗ . More precisely,
Hϕ(0),β (r) ⊂ h−1 ((−∞, c]).

(4.28)

By our assumption ϕ̃(s) = gexpp ( sins β ~u∗ ), we have d(ϕ̃(s), σ0 (s cot β)) ≤ s + o(s)

(4.29)

where lim o(s)
s = 0.
s→0

We now consider a “parallel” transport of Hϕ(0),~u∗ (0) along the quasi-geodesic
ϕ in the sense of Petrunin (cf. [Petr98])
H̃ϕ̃(s),β (r) = {x|∡ϕ̃(s) (ϕ(0), x) = π − β, ∃σ0 such that
σ0 (t) ∈ Hϕ(0),β (r), ∡ϕ̃(0) (~u∗ , x) + ∡ϕ(0) (x, σ0′ (0)) = β}.
Recall that h is a Busemann function h(x) = (c − 100) + d(x, ∂Ωc−100)
which is also distance function. Moreover, Perelman-Sharafutdinov semi-flow related to h is a distance non-increasing semi-flow. Using Theorem 1.11 (Trapezoid
Comparison Theorem), we see that, for each x ∈ H̃ϕ(s),β , we have d(x, Hϕ(0),β (r)) ≤ s + o(s).
39

(4.30)

If follows from (4.28) and (4.30) that
H̃ϕ̃(s),β (r) ⊂ h((−∞, c + s + o(s))).

(4.31)

Thus, H̃ϕ̃(s),β (r) lies below Ωc+s+o(s) and outside Br (ϕ̃(s)).
If ~u(s) is the left derivative of ϕ̃ at s, then
∡ϕ̃(s) (~u(s), x)∡ϕ̃(s) (ϕ̃(0), x) = β
π
Ωc
≥ π − [ − θϕ̃(0),~
u (r)]
2
π
Ωc
(r)
= + θϕ̃(0)
2
Ω

h(ϕ(s))
Let θs (r) = θϕ(s),~
u(s) (r).

(4.32)

Recall that if X n has non-negative curvature, then

gexpq : Tq (X) → X n is a distance non-increasing map. With extra efforts and using discussions above, we have
θs (r) ≥ θ0 (r) + o(s)
where lim o(s)
s = 0. It follows that s→0

∂θs (r)
(0) ≥ 0. 
∂s
§5. Proof of Main theorem
Let a0 = max{h(x) | x ∈ X n } and A0 = h−1 (a0 ). Suppose that X n has positive curvature on B2ε0 (x̂). Then using Proposition 1.8, we see that
Ωc
θx̂,~
u (x) ≥ ar > 0

(5.1)

for some a > 0 and all sufficiently small r > 0. Suppose that ϕ : [0, l] → X n is a
Perelman-Sharafutdinov curve for h such that
 d+ ϕ
∇h

 dt = |∇h|2 ,
ϕ(0) = x̂,


ϕ(l) ∈ A0
40

by Theorem 1.9 and (5.1), we see that
Ω

a0
θϕ(l),~
u(l) (r) > ar > 0

(5.2)

for all 0 < r ≤ r0 . It follows from Proposition 1.8 and (5.2) that either A0 is a point set or A0 has non-empty boundary with a “strictly convex” boundary point
ϕ(0) ∈ ∂A0 .
When ∂A0 6= ∅, we let Ωa0 +s = {z ∈ A0 |d(z, ∂A0 ) ≥ s}, li = max{d(z, ∂A0 )|z ∈
A0 } and A1 = Ωa0 +l1 . Using the same argument as above, we can find z1 ∈ ∂A1
and ~u1 ∈ Tz1 (X) such that

θzA11,~u1 (r) ≥ a1 r > 0

for r > 0 unless dim(A1 ) = 0. Observe that dim(X n ) > dim(A0 ) > dim(A1 ) > · · · .
Repeating above procedure m ≤ n times, we conclude that Am is of dimension

zero. Therefore, X n can be contractible to a point zm = Am via multiple step
Perelman-Sharafutdinov retraction. 
Finally, we remark that our angular excess estimates are related to the lengthexcess function of Alexander-Bishop (cf. [AB09]). To see this, we use the following example. Let ψ : [−ρ, ρ] → R2 be a convex curve given by ψ(t) = (t, λ2 t2 ). The

chord joining ψ(−ρ) to ψ(ρ) has length 2ρ. The length-excess of Alexander-Bishop is given by
Z ρp
Z ρ p
1 + λ2 t2 dt − 2ρ =
[ 1 + λ2 t2 − 1] dt
−ρ

−ρ

=

Z ρ

−ρ

√

2
λ2 t2
dt ∼ λ2 ρ3 .
3
1 + λ2 t2 + 1

However, using our angular excess, we have
θ(ρ) ∼ λ tan ρ ∼ λρ
41

for sufficiently small ρ.
Acknowledgement: Authors are grateful to Professor Stephanie Alexander (University of Illinois at Urbana-Champaign), Professor Vitali Kapovitch (Toronto) and
Xueping Li (Capital Normal University, Beijing) for their criticisms on an earlier version of this paper.
References
[Al49]
[AB65]
[AB03]
[AB09]

[AKP07]

[BBI01]
[BGP92]
[Ca57]
[CDM09]

[CaG10]

[CE08]

[CG71]
[CG72]
[GM69]
[Gro80]

Aleksandrov, A. D., Quasigeodesics. (Russian), Doklady Akad. Nauk SSSR (N.S.)
69 (1949), 717-720.
Aleksandrov, A. D.; Burago, Ju. D., Quasigeodesics lines. (Russian), Trudy Mat.
Inst. Steklov. 76 (1965), 49-63.
Alexander, S. and Bishop, R., FK-convex functions on metric spaces, Manuscripta
Math. 110 (2003), 115-133.
Alexander, S. and Bishop, R., Extrinsic curvature of semiconvex subspaces in Alexandrov geometry, see http://www.math.uiuc.edu/ sba/semi.pdf, to appear in Annals
Global Anal

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
