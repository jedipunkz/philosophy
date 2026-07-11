---
source: "https://arxiv.org/abs/1604.08305v4"
title: "Complex martingales and asymptotic enumeration"
author: "Mikhail Isaev, Brendan D. McKay"
year: "2016"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1604.08305v4"
pdf: "https://arxiv.org/pdf/1604.08305v4"
captured_at: "2026-06-24T11:11:13Z"
updated_at: "2026-06-24T11:11:13Z"
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

# Complex martingales and asymptotic enumeration

- 著者: Mikhail Isaev, Brendan D. McKay
- 年: 2016
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1604.08305v4)
- ダウンロード: https://arxiv.org/pdf/1604.08305v4
- PDF: https://arxiv.org/pdf/1604.08305v4

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Many enumeration problems in combinatorics, including such fundamental questions as the number of regular graphs, can be expressed as high-dimensional complex integrals. Motivated by the need for a systematic study of the asymptotic behaviour of such integrals, we establish explicit bounds on the exponentials of complex martingales. Those bounds applied to the case of truncated normal distributions are precise enough to include and extend many enumerative results of Barvinok, Canfield, Gao, Greenhill, Hartigan, Isaev, McKay, Wang, Wormald, and others. Our method applies to sums as well as integrals. As a first illustration of the power of our theory, we considerably strengthen existing results on the relationship between random graphs or bipartite graphs with specified degrees and the so-called $β$-model of random graphs with independent edges, which is equivalent to the Rasch model in the bipartite case.

## PDF Text

Complex martingales and asymptotic enumeration arXiv:1604.08305v4 [math.CO] 27 Dec 2017

Mikhail Isaev∗† and Brendan D. McKay∗
∗

†

Research School of Computer Science
Australian National University
Canberra ACT 2601, Australia

Moscow Institute of Physics and Technology
Dolgoprudny, 141700, Russia

isaev.m.i@gmail.com, brendan.mckay@anu.edu.au

Abstract
Many enumeration problems in combinatorics, including such fundamental questions as the number of regular graphs, can be expressed as high-dimensional complex integrals. Motivated by the need for a systematic study of the asymptotic behaviour of such integrals, we establish explicit bounds on the exponentials of complex martingales. Those bounds applied to the case of truncated normal distributions are precise enough to include and extend many enumerative results of Barvinok, Canfield, Gao, Greenhill, Hartigan, Isaev, McKay, Wang, Wormald, and others. Our method applies to sums as well as integrals.
As a first illustration of the power of our theory, we considerably strengthen existing results on the relationship between random graphs or bipartite graphs with specified degrees and the so-called β-model of random graphs with independent edges, which is equivalent to the Rasch model in the bipartite case.
∗

Research supported by the Australian Research Council.

1

1

Introduction

A large number of combinatorial enumeration problems can be expressed in terms of highdimensional integrals, often, but not always, resulting from Fourier inversion applied to a multivariable generating function.
To illustrate what we mean, here are two examples. The number of undirected simple graphs with degrees d1 , . . . , dn is given by
I
I Q
1
1≤j<k≤n (1 + zj zk )
·
·
·
dz1 · · · dzn ,
(1.1)
(2πi)n z1d1 +1 · · · zndn +1
while the number of m × n nonnegative integer matrices (contingency tables) with row sums r1 , . . . , rm and column sums c1 , . . . , cn is given by
I
I Q
−1
1
1≤j≤m,1≤k≤n (1 − wj zk )
···
dw1 · · · dwm dz1 · · · dzn ,
(1.2)
r +1
rn +1 c1 +1
(2πi)m+n w11 · · · wm z1 · · · zncn +1
where each contour encloses the origin once anticlockwise. Although explicit evaluation of such integrals is rarely possible, under some circumstances asymptotic estimation is tractable. This was first achieved by McKay and Wormald in 1990, for (1.1) in the case of degree sequences not far from regular [39] and some classes of digraphs that include regular tournaments [35].

Since then, many other examples have appeared that include classes of 0-1 matrices [3,
5,7,8,16,34,41]; directed graphs by degree sequence [15,16,35,38,44,45]; eulerian digraphs
[23, 46]; eulerian circuits [21, 22, 24, 37]; types of integer matrices [4, 9, 36]; and multiple other problems [6, 14, 30, 40]. The method often gives a surprisingly good approximation even for structures of moderate size [8, 9, 17, 25, 36, 37].
Estimation of integrals like (1.1) and (1.2) involves several steps, none of them trivial.
(a) Choose as contours circles rj eiθj whose radii are chosen so that they pass together through the saddle-point (or close enough to it). This involves solving nonlinear equations or maximizing an entropy function.
(b) Identify one or more small regions (in {θj }-space) in which the value of the integral is concentrated. This might be small boxes enclosing two points (as in (1.1)) or the neighbourhood of a low-dimensional subspace (as in (1.2)).
(c) Within those small regions, approximate the integrand by a more tractable function and estimate its integral.

2

The present paper is motivated by step (c). The integrals that occur are typically of the form
Z

I=
exp −xTA x + f (x) dx,
B

where B is a region containing the origin, A is a positive-semidefinite real matrix, and f (x)
is a function well-approximated by a truncated Taylor series with complex coefficients.
The matrix A might not be of full rank.
Now let X be a random variable whose distribution is given by the gaussian density
C exp(−xTAx) truncated to domain B, where C is the normalising constant. Then, by the definition of expectation, we have
R

exp −xTA x + f (x) dx
B
R

= C −1 E ef (X ) ,
I=
T
C B exp −x A x dx

so the problem is reduced to estimating E ef (X) . Our main aim is to make estimation of such integrals more systematic by providing some general theory about E ef (X) .
We will give explicit bounds on E ef (X) that are general and precise enough to cover and generalize the steps corresponding to (c) in all of the examples listed above and many more similar examples. In fact, we will not restrict ourselves to truncated gaussian measures or to functions f that are approximated by polynomials. Furthermore, both our measure and our functions f can be either smooth or discrete, allowing for sums as well as integrals.

1.1

Summary of the paper

Section 2 gives our main theorem in its most general form, providing explicit bounds on E eZn when Z0 , . . . , Zn is a complex martingale, based on properties of the martingale differences. Section 3 applies the martingale theorems to functions of independent random variables, via the Doob martingale. We also show how to bound the necessary parameters for smooth functions and how to handle vector measures whose components are independent only when the measure is rotated.
Section 4 considers the case of gaussian measures which are truncated to a finite region
(usually a cuboid, perhaps intersected with a linear subspace). These are the theorems which can be applied directly to the enumeration problems we have surveyed. The cases of full-rank and non-full-rank gaussians are somewhat different. Finally in that section we give some lemmas useful for managing the quadratic forms which occur.
In Section 5 we demonstrate the power of our theorems using the example of graphs or
3

bipartite graphs with given degrees. In each case, we allow degree sequences as general as those allowed by Barvinok and Hartigan [5], but we also allow a moderate number of forced and forbidden edges. This permits us to prove, in Section 5.3, that the corresponding βmodels are closer than previously known to the uniform model of random graphs with given degrees.
The Appendix collects some technical lemmas we need in the proofs.

1 Introduction
1.1 Summary of the paper . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
2 The exponential of a complex martingale
2.1 The diameter of a complex random variable . . . . . . . . . . . . . . . . .
2.2 First order approach . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
2.3 Second order approach . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
3 Functions of independent random variables
3.1 Estimating the exponential . . . . . . . . . . . . . . . . . . . . . . . . . . .
3.2 Smooth and transformed functions . . . . . . . . . . . . . . . . . . . . . .
4 Truncated gaussian measures
4.1 Truncated gaussian measures of full rank . . . . . . . . . . . . . . . . . . .
4.2 Truncated gaussian measures of less than full rank . . . . . . . . . . . . . .
4.3 Example: regular tournaments . . . . . . . . . . . . . . . . . . . . . . . . .
4.4 The case of weakly dependent components . . . . . . . . . . . . . . . . . .
5 Graphs with given degrees
5.1 General graphs . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
5.2 Bipartite graphs . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
5.3 Concentration near the β-model . . . . . . . . . . . . . . . . . . . . . . . .
6 Appendix

2

2
3
4
5
10
12
15
17
20
22
25
28
30
32
36
38
42
47
50

The exponential of a complex martingale

In this section we state and prove our theorems in their most general forms.
Let P = (Ω, F , P ) be a probability space. We are interested in estimates for the expected value of eZ , where Z is a complex-valued random variable on P . Such estimates for the case of real Z are commonplace as intermediate steps towards concentration inequalities, such as in the classical works of Hoeffding and McDiarmid [19, 33]. However,

4

we seek E eZ itself and few such intermediate results carry over unchanged to the complex case, perhaps fundamentally due to the non-convexity of the exponential function in the complex plane.
As our primary measure of spread of a complex random variable we use the diameter of its essential support. This choice was inspired by its effective use (in the real case) by
McDiarmid [33, Theorem 3.1] in analysing the concentration of functions of independent random variables. A bound on |f (x′ ) − f (x)| whenever x, x′ differ only in the k-th position is, roughly speaking, the same as a bound on the diameter of the random variable f (x1 , . . . , xk−1 , Xk , xk+1 , . . . , xn ) for constant x1 , . . . , xk−1 , xk+1 , . . . , xn .
Note that having diameter α is weaker than being confined to a disk of diameter α.
So while we could alternatively have generalized real intervals into complex disks, doing so would weaken our theorems.
In the next subsection we define the diameter formally, including a conditional version, and prove some properties that we will need. Then, in two further subsections, we use the diameter to bound the exponential of a complex martingale.
Recall that for complex random variables Z there are two types of squared variation commonly defined. The variance is
Var Z = E |Z − E Z|2 = E |Z|2 − |E Z|2 = Var ℜZ + Var ℑZ, while the pseudovariance is
VZ = E (Z − E Z)2 = E Z 2 − (E Z)2 = Var ℜZ − Var ℑZ + 2i Cov(ℜZ, ℑZ).
Of course, these are equal for real random variables.

2.1

The diameter of a complex random variable

Let X be an a.s. bounded real random variable on P = (Ω, F , P ). As usual, define the essential supremum of X as

ess sup X = sup x ∈ R P (X > x) > 0bigr}.
1/r
If |X| ≤ c a.s., it is well-known that ess sup X = −c + limr→∞ E((X + c)r )
. If Z is an a.s. bounded complex random variable on P then we define the diameter of Z to be diam Z = ess sup |Z − Z ′ |,

where Z ′ is an independent copy of Z.

(2.1)

The probability in (2.1) is interpreted in the product space P ⊗P in the standard fashion.
We will also use an equivalent definition that does not use the product space. Given an
5

angle θ, the extent of Z in the θ direction (i.e., the inner product of Z with the unit vector in the θ direction), is ℜ(e−iθ Z), so we can alternatively define

diam Z = sup ess sup(ℜ(e−iθ Z)) + ess sup(−ℜ(e−iθ Z)) .
(2.2)
θ∈(−π,π]

Remark 2.1. To see that (2.1) and (2.2) are equivalent, suppose first that diam Z > d+ε
according to (2.1), for some ε > 0. Assuming that |Z| ≤ c a.s., cover the disk {z | |z| ≤ c}
by finitely many open disks of radius ε/4. If for each pair D, D ′ of such disks whose centres are at least d + ε/2 apart we have P (Z ∈ D, Z ′ ∈ D ′ ) = 0, then ess sup |Z − Z ′ | ≤ d + ε, a contradiction. So choose two of the disks, D, D ′ , with centres at least d + ε/2 apart, such that P (Z ∈ D, Z ′ ∈ D ′ ) = P (Z ∈ D) P (Z ′ ∈ D ′ ) > 0. Taking θ to be the direction from the centre of D to the centre of D ′ , we find that diam Z ≥ d according to (2.2).
Conversely, if there is θ such that the argument of the supθ in (2.2) is greater than d, there are half-planes more than d apart in each of which Z has nonzero probability, proving that diam Z > d according to (2.1).
The basic properties of the diameter of a complex random variable are given by the following lemma.
Lemma 2.2. Let Z be an a.s. bounded complex random variable on P . Then,
(a) diam Z = 0 iff Z is a.s. constant.
(b) diam (aZ + b) = |a| diam Z for any a, b ∈ C.
(c) diam (Z + W ) ≤ diam Z + diam W for any a.s. bounded complex random variable
W on P .
(d) diam ℜZ ≤ diam Z ≤ 2 ess sup |Z|.

(e) diam (ℜZ)2 ≤ diam Z 2 ≤ 2 ess sup |Z| · diam Z.
(f ) |Z − E Z| ≤ diam Z a.s.
(g) There exists a ∈ C such that |Z − a| ≤ √13 diam Z a.s.

(h) Var Z = E |Z − E Z|2 ≤ 31 (diam Z)2 and |VZ| = |E (Z − E Z)2 | ≤ 41 (diam Z)2 .
Proof. Claims (a),(b) follow immediately from Definition (2.1). We get claim (c) from
Definition (2.2) and the fact that ess sup(X + Y ) ≤ ess sup X + ess sup Y for any a.s.
bounded real random variables X, Y on P .

6

Let Z ′ be an independent copy of Z. We note then (almost surely) that
|ℜZ − ℜZ ′ | ≤ |Z − Z ′ | ≤ ess sup |Z − Z ′ | = diam Z,

|(ℜZ)2 − (ℜZ ′ )2 | = |ℜ(Z − Z ′ )| · |ℜ(Z + Z ′ )| ≤ |Z − Z ′ | · |Z + Z ′ |

≤ ess sup |Z 2 − (Z ′ )2 | = diam Z 2 .

|Z − Z ′ | ≤ ess sup |Z| + ess sup |Z ′ | = 2 ess sup |Z|.

Z 2 − (Z ′ )2 ≤ ess sup |Z + Z ′ | · ess sup |Z − Z ′ | ≤ 2 ess sup |Z| · diam Z.
Due to Definition (2.1), claims (d) and (e) follow.
Using Definition (2.2), the fact that |X − EX| ≤ ess sup(X) − ess sup(−X) a.s. for any a.s. bounded real random variable X on P and the equation
|Z − E Z| = sup

θ∈(−π,π]

ℜ(e−iθ (Z − E Z)) ,

we obtain claim (f).
Claim (g) follows from a standard result on convex sets, see [32, Thm. 12.3] for example. An equilateral triangle shows that the constant cannot be reduced. To prove the first part of claim (h), note that Var Z = Var(Z − a) ≤ ess sup |Z − a|2 .
However, for any a.s. bounded real random variable X on P

X − 21 (ess sup X + ess sup(−X)) ≤ 12 (ess sup X − ess sup(−X)) a.s., which implies Var X ≤ 41 (diam X)2 . To prove the second part of claim (h), note that

|E(Z − E Z)2 | ≤ E ℜ(e−iθ (Z − E Z))2 = Var ℜ(e−iθ Z),

where eiθ = E(Z − E Z)2 /|E(Z − E Z)2 |, and diam ℜ(e−iθ Z) ≤ diam Z on account of claims (b) and (d).
We will also use a conditional version of the diameter. Let G ⊆ F be a σ-field. For a real random variable X on P = (Ω, F , P ) such that |X| ≤ c a.s., we can define the conditional essential supremum of X to be the G-measurable function
1/r ess sup (X | G) = −c + lim E((X + c)r G)
.
(2.3)
r→∞

Alternative equivalent definitions and many properties of the conditional essential supremum are given in [1]. Informally, ess sup (X | G) is the least G-measurable function
G : Ω → R such that X ≤ G a.s. Now we can extend (2.2) to define the conditional diameter :

(2.4)
diam(Z | G) = sup ess sup(ℜ(e−iθ Z) | G) + ess sup(−ℜ(e−iθ Z) | G) .
θ∈(−π,π]

7

Note that diam(Z | G) is a function from Ω to R+ . For any ω ∈ Ω, the argument of the supθ in (2.4) is a continuous function of θ (since Z is a.s. bounded), so the supremum over θ is the same if restricted to a dense countable subset of (−π, π]. This proves that diam(Z | G) is G-measurable.
If Z is real, we can restrict (2.4) to θ = 0 and then diam(Z | G) is the same as the conditional range defined by McDiarmid [33, Sec. 3.4].

Now let PZ|G : B(C) × Ω → [0, 1] be a regular conditional distribution for Z given G, where B(C) is the Borel field of C. That is, for each ω ∈ Ω, PZ|G (·, ω) is a probability measure on B(C), and for each A ∈ B(C), PZ|G (A, ·) is G-measurable and PZ|G (A, ·) =
P (Z −1 (A) | G) a.s. For the existence of PZ|G and basic theory, see [29, Chap. 6].

For each ω ∈ Ω, let Kω (Z | G) be the class of random variables from Ω to C that induce the distribution PZ|G (·, ω) on B(C). The most important property of Kω (Z | G) is:
Lemma 2.3. Let G ⊆ F be a σ-field and Z be an a.s. bounded complex random variable on P . Let Zω be an arbitrary member of Kω (Z | G) for each ω ∈ Ω. Let W be a Gmeasurable random variable on P , and let φ : C × W (Ω) → C be a measurable function such that E |φ(Z, W )| < ∞. Then, for almost all ω ∈ Ω,
E(φ(Z, W ) | G)(ω) = E φ(Zω , W ),

(2.5)

and moreover φ(Zω , W ) ∈ Kω (φ(Z, W ) | G). Also the random variable ω 7→ E φ(Zω , W )
is G-measurable. Consequently, for almost all ω ∈ Ω, ess sup (|φ(Z, W )| | G)(ω) = ess sup |φ(Zω , W )|, diam(φ(Z, W ) | G)(ω) = diam φ(Zω , W ).

(2.6)

Proof. Equation (2.5) is Theorem 6.4 in [29]. By applying it to functions of the form
1A (φ(·, ·)) for each A ∈ B(C), we find that φ(Zω , W ) ∈ Kω (φ(Z, W ) | G). The Gmeasurability of ω 7→ E φ(Zω ) follows from the G-measurability of the left side of (2.5).
Equation (2.6) follows from (2.5) on account of (2.3) and (2.4). Note also that the Gmeasurability of the conditional essential supremum and the conditional diameter is just a special case of this.
We now list a number of properties of the conditional diameter that we will need.
Lemma 2.4. Let G ⊆ F be a σ-field and Z be an a.s. bounded complex random variable on P . Then,
(a) diam(ℜZ | G) ≤ diam(Z | G) a.s.
8


(b) diam Z | G ≤ 2 ess sup(|Z| | G) a.s.

(c) diam(Z 2 | G) ≤ 2 ess sup(|Z| | G) diam(Z | G) a.s.

(d) |Z − E(Z | G)| ≤ diam(Z | G) a.s.
(e) If the σ-field H ⊆ F is independent of G, then diam E(Z | H) ≤ E diam(Z | G) a.s.

(f ) E (Z − E(Z | G))(W − E(W | G)) G ≤ 31 diam(Z | G) · diam(W | G) a.s. for any a.s. bounded complex random variable W on P .
(g) If U and W are G-measurable, then diam(UZ + W | G) = |U| diam(Z | G).
Proof. Let Zω be an arbitrary member of Kω (Z | G) for each ω ∈ Ω. By Lemma 2.3, we have that, for almost all ω ∈ Ω,
E(Z | G)(ω) = E Zω ,

ess sup(|Z| | G)(ω) = ess sup |Zω |, diam(ℜZ | G)(ω) = diam ℜZω , diam(Z | G)(ω) = diam Zω ,

diam(Z 2 | G)(ω) = diam Zω2 ,

ess sup( Z − E(Z | G) | G)(ω) = ess sup |Zω − E Zω |,

2
E Z − E(Z | G) | G (ω) = E |Zω − E Zω |2 = Var Zω .

Due to Lemma 2.2(d, e), claims (a)–(c) follow.

In order to prove claims (d) and (e), recall from [1, Prop. 2.6] that for a bounded real random variable X,
X ≤ ess sup(X | G) a.s.

(2.7)

Therefore,
|Z − E(Z | G)| ≤ ess sup(|Z − E(Z | G)| | G) a.s.
and we get claim (d) from Lemma 2.2(f).
Using (2.7) and the independence of G and H,
E(X | H) ≤ E(ess sup(X | G) | H) = E ess sup(X | G) a.s.
for a bounded real random variable X. Applying this to the Definition (2.4) with X =
ℜ(e−iθ Z) and X = −ℜ(e−iθ Z), claim (e) follows.

9

Claim (f) is due to Lemma 2.2(h) and the conditional Cauchy-Schwartz inequality

E (Z − E(Z | G))(W − E(W | G)) | G
q
q

2
2
≤ E Z − E(Z | G) | G
E W − E(W | G) | G a.s.

To prove claim (g), note that the properties of the conditional essential supremum imply


ess sup ℜ(e−iθ (U + W Z)) | G = ℜ(e−iθ U) + |W | ess sup ℜ(e−iθ+i arg(W ) Z) | G ,

and apply this to the definition of conditional diameter.

In [26] we proved the following generalization of a bound of Hoeffding [19].
Lemma 2.5. If Z is an a.s. bounded complex random variable on P , then
2

1

E eZ−E Z − 1 ≤ e 8 diam(Z) − 1.
Corollary 2.6. Let Z be an a.s. bounded complex random variable on P and let G ⊆ F
be a σ-field. Then we have
1

2

E(eZ−E(Z | G) | G) − 1 ≤ e 8 diam(Z | G) − 1 a.s.
Proof. It suffices to apply the lemma to arbitrary random variables Zω ∈ Kω (Z | G), with the help of (2.6).

2.2

First order approach

A sequence F = F0 , . . . , Fn of σ-subfields of F is a filter if F0 ⊆ · · · ⊆ Fn . A sequence
Z0 , . . . , Zn of random variables on P = (Ω, F , P ) is a martingale with respect to F if
(i) Zj is Fj -measurable and has finite expectation, for 0 ≤ j ≤ n;
(ii) E(Zj | Fj−1) = Zj−1 for 1 ≤ j ≤ n.
Note that, up to almost-sure equality, the martingale is determined by Zn and F, namely
Zj = E(Zn | Fj ) a.s. for each j.

If Z is a random variable on P and 0 ≤ j ≤ n, we use the following notations for statistics conditional on Fj :
Ej Z = E(Z | Fj ),


Vj Z = E (Z − Ej (Z))2 Fj = Ej Z 2 − (Ej Z)2 ,

diamj Z = diam(Z | Fj ).

10

If F0 = {∅, Ω}, which we not assume unless it is stated explicitly, E0 Z, V0 Z and diam0 Z
equal the unconditional versions E Z, VZ and diam Z, respectively.
An extremely large literature concerns concentration of martingales derived from restrictions on the differences Zj − Zj−1 , but most of it considers only real martingales and can’t be assumed to hold for complex martingales. The fact that the real and imaginary parts of a complex martingale are real martingales can often be applied, but at the cost of weaker bounds. In any case, our aim is for estimates of the exponential rather than for concentration. Here again, the non-convexity of the exponential function in the complex plane often means that theorems and proofs for the real case do not carry over immediately to the complex case.
Theorem 2.7. Let Z = Z0 , Z1 , . . . , Zn be an a.s. bounded complex-valued martingale with respect to a filter F0 , . . . , Fn . Define
Rk = diamk−1 Zk

(2.8)

for 1 ≤ k ≤ n. Then
E0 eZn = eZ0 (1 + K(Z)), where K(Z) is an F0 -measurable random variable with
2
1 Pn k=1 Rk


F0 − 1 a.s.

|K(Z)| ≤ ess sup e 8

Proof. Since Ek−1 Zk = Zk−1, we have for 1 ≤ k ≤ n that
Ek−1 eZk = Ek−1 eZk−1 +(Zk −Zk−1 )
= eZk−1 1 + Ek−1 (eZk −Zk−1 − 1)
= eZk−1 + Uk eZk−1



(2.9)

2

for some Fk−1 -measurable Uk such that |Uk | ≤ eRk /8 − 1 a.s., by Corollary 2.6.

Now recall that |ez | = eℜz for all z and note that ℜZ0 , . . . , ℜZn is also a martingale satisfying the conditions of the theorem on account of Lemma 2.4(a). Therefore we similarly have that
Ek−1 |eZk | = |eZk−1 | + Uk′ |eZk−1 | = (1 + Uk′ ) |eZk−1 |
2
Rk /8

for some Fk−1 -measurable Uk′ such that |Uk′ | ≤ e backwards induction on k that for 0 ≤ k ≤ n,

Ek eZn = eZk + Wk eZk ,

11

(2.10)

− 1 a.s. Now we can prove by
(2.11)

2

1 Pn where Wk is Fk -measurable and |Wk | ≤ ess sup e 8 j=k+1 Rj Fk −1 a.s. Obviously (2.11)
is true for k = n. Now observe from (2.9) and (2.11) that Ek−1 Zn = eZk−1 + Uk eZk−1 +

Ek−1 (Wk eZk ) and note that Ek−1(Wk eZk ) ≤ ess sup |Wk | Fk−1 Ek−1 |eZk | a.s. Applying (2.10) to the last term and combining the error terms using Lemma 6.2, we obtain (2.11) for k − 1. The case k = 0 gives the theorem.

2.3

Second order approach

In the following we need two technical bounds that are in the Appendix, Lemma 6.1. We also use the following elementary lemma.
Lemma 2.8. Let Z = Z0 , Z1 , . . . , Zn be a bounded complex-valued martingale with respect to a filter F0 , . . . , Fn . Then
2

Ek (Zn − Zk ) =

n
X

j=k+1

Ek (Zj − Zj−1)2

for 0 ≤ k ≤ n.
Proof. For 0 ≤ k ≤ j ≤ ℓ ≤ n,
Ek (Zℓ − Zj )2 = Ek (Ej (Zℓ − Zj )2 ) = Ek (Ej Zℓ2 − Zj2 ) = Ek (Zℓ2 − Zj2 ).
Therefore,
Ek (Zn − Zk )2 = Ek (Zn2 − Zk2 )

2
2
2
2
= Ek (Zn2 − Zn−1
) + (Zn−1
− Zn−2
) + · · · + (Zk+1
− Zk2 )
n
X
Ek (Zj − Zj−1 )2 .
=



j=k+1

Theorem 2.9. Let Z = Z0 , Z1 , . . . , Zn be a bounded complex-valued martingale with respect to a filter F0 , . . . , Fn . For 1 ≤ k ≤ n, define

Then

Rk = diamk−1 Zk ,

Qk = max diamk−1 Ek (Zn − Zk )2 , diamk−1 Ek (ℜZn − ℜZk )2 .
1

1

E0 eZn = eZ0 + 2 V0 Zn + L(Z)eℜZ0 + 2 V0 (ℜZn )

1
1
= eZ0 + 2 V0 Zn 1 + L′ (Z)e 2 V0 (ℑZn ) ,
12

(2.12a)
(2.12b)

where L(Z), L′ (Z) are F0 -measurable random variables with n


X

2
′
1
5 4
5
1 3
R + 6 Rk Qk + 8 Rk + 32 Qk
F0 − 1 a.s.
|L(Z)| = |L (Z)| ≤ ess sup exp
6 k k=1

Proof. All equalities and inequalities in the proof should be taken “almost surely”. For
1 ≤ k ≤ n we have, using Lemma 2.8,
1

1

Ek−1 eZk + 2 Vk Zn = eZk−1 + 2 Vk−1 Zn

2

1

2

+ eZk−1 + 2 Ek−1 (Zn −Zk ) Ek−1 (eAk − eAk /2 − Ak )eBk
2

1
+ eZk−1 + 2 Ek−1 (Zn −Zk ) Ek−1 Ak (eBk − Bk − 1)



2

1

+ eZk−1 + 2 Ek−1 (Zn −Zk ) Ek−1 (Ak + Ak Bk )
1

+ eZk−1 + 2 Vk−1 Zn Ek−1 (eCk − 1),

where

Ak = Zk − Zk−1 ,
Bk = 12 Ek (Zn − Zk )2 − 21 Ek−1(Zn − Zk )2 = 21 Vk Zn − 12 Ek−1 Vk Zn ,
Ck = 21 (Zk − Zk−1)2 + 21 Vk Zn − 12 Vk−1 Zn
= Bk + 21 A2k − 12 Ek−1 A2k .
Note that
Ek−1 Ak = Ek−1 Bk = Ek−1 Ck = 0.
Therefore, by the conditions of the theorem and Lemma 2.4(b,d),
|Ak | ≤ Rk ,

|Bk | ≤ diamk−1 Bk = 21 diamk−1 Ek (Zn − Zk )2 ≤ 12 Qk and

|Ck | ≤ diamk−1 Ck ≤ 21 diamk−1 (Zk − Zk−1 )2 + 12 diamk−1 Ek (Zn − Zk )2

≤ ess sup |Zk − Zk−1 |2 Fk−1 + 21 Qk ≤ Rk2 + 21 Qk

By Corollary 2.6,

1

2

2

1

4

1

2

Ek−1(eCk − 1) ≤ e 8 (Rk +Qk /2) − 1 ≤ e 4 Rk + 16 Qk − 1.
Using Lemma 6.1 and Corollary 2.6 with the triangle inequality, we get that
2

1 3 1 4
Ek−1 (eAk − eAk /2 − Ak )eBk ≤ (e 6 Rk + 8 Rk − 1) Ek−1 (|eBk |)
2
2

1 3 1 4
1 3 1 4
1
1
= (e 6 Rk + 8 Rk − 1) Ek−1 (eℜBk − 1) + 1 ≤ e 6 Rk + 8 Rk + 32 Qk − e 32 Qk ,
1 2
1 2 1 4
1

Ek−1 Ak (eBk − Bk − 1) ≤ e 32 Qk + e 6 Rk Qk + 16 Qk + 4 Rk − 16 Rk Qk − 2.
13

(2.13)

By Lemma 2.4(f), we have
|Ek−1 Ak Bk | ≤ 31 diamk−1 Ak · diamk−1 Bk ≤ 16 Rk Qk .
Therefore, for each k, formula (2.13) gives
1

1

1

Ek−1 eZk + 2 Vk Zn = eZk−1 + 2 Vk−1 Zn + Lk eZk−1 + 2 Vk−1 Zn
2

1

+L′k eZk−1 + 2 Ek−1 (Zn −Zk )

(2.14)

for some Fk−1 -measurable random variables Lk and L′k with
1

4

1

1

3

3

2

|Lk | ≤ e 4 Rk + 16 Qk − 1,
4

1

3

2

|L′k | ≤ e 6 Rk + 8 Rk + 6 Rk Qk + 32 Qk − 1.

Now consider the martingale X0 , . . . , Xn of the real parts of Z0 , . . . , Zn . In order to bound the second and third terms of (2.14) we consider the absolute value
1

1

1

1

|eZk−1 + 2 Vk−1 Zn | = eℜZk−1 + 2 Vk−1 ℜZn − 2 Vk−1 ℑZn ≤ eXk−1 + 2 Vk−1 Xn ,
2

1

2

1

2

1

|eZk−1 + 2 Ek−1 (Zn −Zk ) | = eℜZk−1 + 2 Ek−1 (ℜZn −ℜZk ) − 2 Ek−1 (ℑZn −ℑZk )
2

1

≤ eXk−1 + 2 Ek−1 (Xn −Xk )

(2.15)

2
Xk−1 + 21 Vk−1 Xn − 21 Ek−1 (Xk −Xk−1 )

=e

1

≤ eXk−1 + 2 Vk−1 Xn .
Due to Lemma 2.4(a), this martingale also satisfies conditions (2.12). Therefore, by the same reasoning as before and using the inequality
2

1

1

eXk−1 + 2 Ek−1 (Xn −Xk ) ≤ eXk−1 + 2 Vk−1 Xn , we get that
1

1

1

Ek−1 eXk + 2 Vk Xn = eXk−1 + 2 Vk−1 Xn + L′′k eXk−1 + 2 Vk−1 Xn , where |L′′k | ≤ e

2
1 4
R +1Q
4 k 16 k

−1+e

2
1 3 3 4 1
R + R + R Q +3Q
6 k 8 k 6 k k 32 k

−1≤e

2
1 3 5 4 1
R + R + R Q +5Q
6 k 8 k 6 k k 32 k

Using (2.14) and (2.16), we now prove by backwards induction on k that
1

1

Ek eZn = eZk + 2 Vk Zn + Mk eXk + 2 Vk Xn ,

(2.16)
− 1.
(2.17)

where
Pn

|Mk | ≤ ess sup e

2
1 3 5 4 1
5
j=k+1 ( 6 Rj + 8 Rj + 6 Rj Qj + 32 Qj )


Fk − 1.

The claim is obviously true for k = n. To perform the induction step, take the expectation of (2.17) with respect to Fk−1 , using (2.14) and (2.15) for the first term on the right, and (2.16) to bound the second term on the right. Using the bound
14

1

3

5

4

1

5

2

e 6 Rk + 8 Rk + 6 Rk Qk + 32 Qk − 1 for both |Lk | + |L′k | and |L′′k |, we obtain (2.17) for Ek−1 eZn on combining the error terms using Lemma 6.2. After n steps we reach the expression for
E0 eZn stated in the theorem.
Remark 2.10. Although the two options on the right side of (2.12b) are the same for real martingales, either one of them can be the largest for complex martingales. The case of a purely imaginary martingale shows that the first can be larger. To show that the second can be larger, consider independent variables X, Y , where X ∈ {1, e2πi/3 , e−2πi/3 } with equal probabilities, and Y ∈ {0, 1} with equal probabilities. Now consider the martingale
Z0 , Z1 , Z2 where Z2 = XY , Z1 = E(Z2 | F1 ) = 0 and Z0 = E(Z1 | F0 ) = 0, where
F0 = {∅, Ω} and F1 = σ(Y ). We find that E1 (Z2 − Z1 )2 = 0 and E1 (ℜZ2 − ℜZ1 )2 ∈

{ 21 , 1} with probabilities 32 , 13 respectively. Therefore diam0 E1 (Z2 − Z1 )2 = 0 < 21 =

diam0 E1 (ℜZ2 − ℜZ1 )2 .

3

Functions of independent random variables

In this section we apply our martingale theorems to the case of functions of independent random variables.
An important example of a martingale is made by the so-called Doob martingale process. Suppose X1 , X2 , . . . , Xn are random variables and f (X1 , X2 , . . . , Xn ) is a random variable of bounded expectation. Then we have the martingale {Zj } with respect to
{Fj }, where for each j, Fj = σ(X1 , . . . , Xj ) and Zj = E(f (X1 , X2 , . . . , Xn ) | Fj ). In particular, F0 = {∅, Ω} and Z0 = E f (X1 , X2 , . . . , Xn ). We will also use the σ-fields
F (j) = σ(X1 , . . . , Xj−1 , Xj+1, . . . , Xn ) for 1 ≤ j ≤ n.

In this section we use the Doob martingale to find bounds on E ef . We first need some preliminary lemmas in order to show that all assumptions of Theorems 2.7 and 2.9 are satisfied.

Lemma 3.1. Suppose that X, Y are independent random variables on P , and that g is a complex-valued function such that g(X, Y ) is bounded and measurable. Then,
(a) diam(g(X, Y ) | σ(X))(ω) = diam g(X, Y (ω)) for almost all ω ∈ Ω.

(b) diam(g(X, Y ) | σ(X)) ≤ supx∈X(Ω), y,y′ ∈Y (Ω) |g(x, y) − g(x, y ′)| a.s.

(c) diam g(X, Y ) − E(g(X, Y ) | σ(Y )) σ(X)
≤ supx,x′ ∈X(Ω), y,y′ ∈Y (Ω) |g(x, y) − g(x′ , y) − g(x, y ′) + g(x′ , y ′)| a.s.
15

Proof. Since Y is by definition σ(Y )-measurable, Lemma 2.3 tells us that g(X(ω), Y ) ∈
Kω (g(X, Y ) | σ(X)). Claim (a) is thus just the definition of the conditional diameter.
Similarly, g(X(ω), Y ) ∈ Kω (g(X, Y ) | σ(X)), which gives claim (b) when we apply the definition of diam(g(X(ω), Y )).
For claim (c), note that for almost all ω ∈ Ω, E(g(X, Y ) | σ(Y ))(ω) = E g(X, Y (ω)).
Therefore, applying claim (b),

diam g(X, Y ) − E(g(X, Y ) | σ(Y )) σ(X)
≤

≤

sup

′

x∈X(Ω), y,y ∈Y (Ω)

|g(x, y) − E(X, y) − g(x, y ′) + E(X, y ′ )|

sup
′

′

x,x ∈X(Ω), y,y ∈Y (Ω)

|g(x, y) − g(x′ , y) − g(x, y ′) + g(x′ , y ′)|

since |E U| ≤ sup |U| for any complex random variable.
We will deal with functions with additional arguments. For these purposes we state the following corollary of Lemma 2.4 and Lemma 3.1.
Corollary 3.2. Suppose that W, X, Y are independent random variables on P , and that h is a complex-valued function such that h(W, X, Y ) is bounded and measurable. Then, a.s.,


(a) diam E(h(W, X, Y ) | σ(W, X)) σ(W ) ≤ E diam(h(W, X, Y ) | σ(W, Y )) σ(W ) .

(b) diam h(W, X, Y ) − E(h(W, X, Y ) | σ(W, Y )) σ(W, X)
≤ supw∈W (Ω),x,x′∈X(Ω), y,y′ ∈Y (Ω) |h(w, x, y) − h(w, x′ , y) − h(w, x, y ′) + h(w, x′ , y ′)|.
Proof. Using Lemma 3.1(a), we note that for almost all ω ∈ Ω,
E(h(W, X, Y ) | σ(W, X))(ω) = E(h(W (ω), X(ω), Y ))

= E(h(W (ω), X, Y ) | σ(X))(ω),

diam(h(W, X, Y ) | σ(W, X))(ω) = diam(h(W (ω), X(ω), Y ))

= diam(h(W (ω), X, Y ) | σ(X))(ω),

and the same with X and Y interchanged, so we can apply Lemma 2.4(e) and Lemma 3.1(c)
to random variables given by the two-argument functions gω (X, Y ) = h(W (ω), X, Y ) to obtain both claims.

16

3.1

Estimating the exponential

In this section we state our main results when applied to the case of complex functions of independent random variables. Let P = (Ω, F , P ) be a probability space. Let S =
S1 × · · · × Sn be any n-dimensional domain and consider a function F : S → C. For
1 ≤ k ≤ n, define
αk (F, S) = sup |F (xk ) − F (x)|,

(3.1)

where the supremum is over x, xk ∈ S that differ only in the k-th coordinate. Similarly, for j 6= k, define
∆jk (F, S) = sup |F (x) − F (xj ) − F (xk ) + F (xjk )|,

(3.2)

where the supremum is over x, xk , xj , xjk ∈ S such that x, xk differ only in the k-th component, x, xj differ only in the j-th component, xj , xjk only in the k-th component, and xk , xjk only in the j-th component. We also define the column vector α(F, S) =

(α1 (F, S), . . . , αn (F, S))T and the matrix ∆(F, S) = ∆jk (F, S) with zero diagonal.

Theorem 3.3. Let X = (X1 , . . . , Xn ) be a random vector on P with independent components, and let f : X(Ω) → C be a measurable function. Let α = α(f, X(Ω)) and
∆ = ∆(f, X(Ω)).
(a) We have
E ef (X) = eE f (X) (1 + K),

(3.3)

where K = K(f (X)) is a complex constant with |K| ≤ e

1 T
α α
8

− 1.

(b) We have
1

1

E ef (X ) = eE f (X)+ 2 Vf (X) + L eE ℜf (X)+ 2 Var ℜf (X)
1

1

= eE f (X)+ 2 Vf (X) 1 + L′ e 2 Var ℑf (X) ),

(3.4)

where L = L(f (X)) and L′ = L′ (f (X)) are complex constants with
 X

n n
X
′
3
4
T
T 2
1
1
5
5
αk + 6 α ∆α + 8
αk + 16 α ∆ α − 1.
|L| = |L | ≤ exp 6
k=1

k=1

Proof. Consider the martingale {Zk } with respect to {Fk } obtained by the Doob martin-

17

gale process. For 1 ≤ k ≤ n we have diamk−1 Zk = diam E(f (X) | Fk ) Fk−1




≤ E diam(f (X) | F (k) ) Fk−1 , by Corollary 3.2(a)

≤ ess sup diam(f (X) F (k) )
≤ sup |f (x) − f (xk )|, by Lemma 3.1(b)
k

x,x

≤ αk , by assumption.
Now formula (3.3) follows from Theorem 2.7.
We next consider Ek (Zn −Zk )2 , which by Lemma 2.8 is equal to diamk−1 Ek (Zj − Zj−1 )2

= diam E (Zj − Zj−1)2 | Fk

Pn

2
j=k+1 Ek (Zj −Zj−1 ) .




Fk−1 ,


≤ E diam (Zj − Zj−1 )2 | F (k) ∩ Fj Fk−1 by Corollary 3.2(a).

(3.5)

≤ 2 ess sup |Zj − Zj−1| · ess sup diam(Zj − Zj−1 | F (k) ∩ Fj ),

by Lemma 2.4(c).

By Lemma 2.4(d),
|Zj − Zj−1| ≤ diamj−1 Zj ≤ αj .

(3.6)

Using Corollary 3.2(a,b), we find that

diam(Zj − Zj−1 | F (k) ∩ Fj ) = diam Ej (f (X) − E(f (X) | F (j) )) | F (k) ∩ Fj


≤ E diam f (X) − E(f (X) | F (j) ) | F (k) | F (k) ∩ Fj

≤ ess sup diam f (X) − E(f (X) | F (j) ) F (k)
≤ sup |f (x) − f (xj ) − f (xk ) + f (xjk )|
≤ ∆jk , by assumption,

(3.7)

where the last supremum is over x, xk , xj , xjk ∈ X(Ω) such that x, xk differ only in the k-th component, x, xj differ only in the j-th component, xj , xjk only in the k-th component, and xk , xjk only in the j-th component. Combining (3.5)–(3.7), we obtain that diamk−1 Ek (Zj − Zj−1)2 ≤ 2αj ∆jk .
The same bound holds for diamk−1 Ek (ℜZn − ℜZk )2 , since the Doob martingale of
ℜf (X) also satisfies conditions (a) and (b) of the theorem.

18

Now we can apply Theorem 2.9 to obtain (3.3) with
 X
n n
n n
n n  X
2 
X
X
X
X
3
4
1
5
1
5
αj ∆jk
− 1.
|L(f, X)| ≤ exp 6
αk + 8
αk + 3
αj αk ∆jk + 8
k=1

k=1

k=1 j=k+1

k=1 j=k+1

Since the matrix ∆ is symmetric, the third term in the summation equals 61 αT∆α.
2
P
Pn
The term A = nk=1
depends on the order that the arguments of j=k+1 αj ∆jk f are listed, but we can define the martingale using any order we wish. If we write
P
A = j>k,ℓ>k αj ∆jk ∆kℓ αℓ , then the version from the reverse order of the arguments is
P
A′ = j<k,ℓ<k αj ∆jk ∆kℓ αℓ . Since A and A′ provide disjoint sets of terms of αT∆2 α =
P
1 T 2
j,k,ℓ αj ∆jk ∆kℓ αℓ , at least one of them is bounded by 2 α ∆ α. This completes the proof.
Remark 3.4. A result similar to Theorem 3.3 was proved by Catoni [10, 11] when the function f is real, and used to obtain concentration bounds of the form



t2
 ,
P f (X) ≥ E f (X) + t ≤ exp −
2 Var f (X) + ηt/ Var f (X)

where η is a certain constant depending on α and ∆. We won’t pursue that direction here since we are interested in the complex case which is required for our applications.
The complex case has the added advantage that we can use it to estimate characteristic functions and not just Laplace transforms, with interesting consequences that include
Berry–Esseen-type inequalities which we will explore in a further paper.

Another point to mention in comparison with Catoni’s theorems is that he doesn’t
P
have fourth-order terms such as the term 58 nk=1 αk4 in Theorem 3.3(b). Although those terms make the bound much larger for very large {αj }, in such extreme cases part (a) of the theorem generally gives a better result anyway. We have included fourth order terms in order to allow better constants on the third order terms.
1

Remark 3.5. The factor e 2 Var ℑf (X) appearing in the error term of (3.4) is of course redundant in the case that f is real. The following example shows that some such multiplier is required in the general complex case. Suppose that the components of
X = (X1 , . . . , Xn ) are iid random variables with mass 21 at each of ±n−1/2+ε . Define
P
X = nj=1 Xj and f (X) = iX + n1 e−iX . We obviously have E X = 0 and E X 2 = n2ε .
For c = ±1, we have
−1/2+ε
−1/2+ε n
n
E eicX = E eicX1 = 21 e−in
+ 21 ein
2ε
−1/2+3ε
n
)
.
= 1 − 12 n−1+2ε + O(n−3/2+3ε ) = e−n /2+O(n
19

Using ef (X) = eiX + n1 +O( 2n1 2 ) we have E ef (X ) = n1 +O( n12 ). Now let us apply Theorem 3.3.
2ε

1

We have E f (X) = n1 e−n /2+o(1)) and Vf (X) = −n2ε + o(1). Therefore eE f + 2 Vf =
1

2ε

e− 2 n +o(1) . In the error term of (3.4) we have αk = O(n−1/2+ε ) and ∆jk = O(n−2+2ε ). So
1
eE f + 2 Vf is very much smaller than E ef (X) even though L′ = o(1). In a later paper we will investigate a wide class of complex functions for which a theorem similar to Theorem 3.3
1
is true without the factor e 2 Var ℑf (X) .

3.2

Smooth and transformed functions

In the case of smooth functions, the parameters αj and ∆jk can be bounded in terms of derivatives or other measures such as Lipschitz constants. For our applications in
Section 4, it will suffice to have continuous differentiability.
If f is a function one of whose arguments is x, then fx is the partial derivative ∂f /∂x, and similarly for notations like fxy . If the arguments are a subscripted list, like x1 , . . . , xn , we will further abbreviate fxj to fj and fxj xk to fjk . The notations k·k1 , k·k2 and k·k∞
have their usual meanings as vector norms and the corresponding induced matrix norms.
For a matrix A = (ajk ) we will also use kAkmax = maxjk |ajk | but note that it is not submultiplicative.
Lemma 3.6.
(a) Let L be the closed line segment [x1 , x2 ] ⊆ R and let S be its interior minus a countable set of points. Suppose that the function f : L → C is continuous, and that fx exists and is bounded in S. Then
|f (x2 ) − f (x1 )| ≤ |x2 − x1 | sup |fx (x)|.
x∈S

(b) Let R be the closed rectangle [x1 , x2 ] × [y1 , y2] ⊆ R2 and let S be its interior minus a countable set of lines. Suppose that the function f : R → C is continuous and fx exists and is continuous. Moreover assume that fxy exists and is bounded in S.
Then f (x1 , y1) − f (x1 , y2 ) − f (x2 , y1 ) + f (x2 , y2) ≤ |x2 − x1 | |y2 − y1 | sup |fxy (x, y)|.
(x,y)∈S

Proof. The conditions we have given are sufficient to imply that
Z (HK)
f (x2 ) − f (x1 ) =
fx (x) dx
L

20

in case (a) and f (x1 , y1) − f (x1 , y2) − f (x2 , y1) + f (x2 , y2 ) =

Z (HK)  Z (HK)
[y1 ,y2 ]

[x1 ,x2 ]


fxy (x, y) dx dy,

in case (b), where we have used the Henstock–Kurzweil (gauge) integral [2, Thm. 4.7].
The claims now follow readily.
Note that in part (a) we did not require that fx is continuous, and in part (b) we did not require that fy or fxy are continuous. The lemma is not true if “countable” is replaced by “measure zero”. In the following we will adopt more stringent conditions on derivatives than Lemma 3.6 allows, leaving the generalizations to future applications.
Corollary 3.7. Let B = [a1 , b1 ] × · · · × [an , bn ] ⊆ Rn . Suppose that f : B → C is continuous. Then, provided the suprema exist,
(a) If f is continuously differentiable in the interior int B of B, then, for 1 ≤ j ≤ n,
αj (f, B) ≤ (bj − aj ) sup |fj (x)|.
x∈int B

(b) If f is twice continuously differentiable in int B, then, for 1 ≤ j < k ≤ n,
∆jk (f, B) ≤ (bj − aj )(bk − ak ) sup |fjk (x)|.
x∈int B

Proof. This follows immediately from Lemma 3.6, noting that line segments or rectangles in the boundary of B are limits of line segments or rectangles not in the boundary.
In the case of a transformed cuboid, it is convenient to be able to bound kαk∞ , αT∆α
and αT∆2 α in terms of the derivatives in the original coordinates. We will only treat the case of uniformly bounded derivatives. For ρ > 0, define
Un (ρ) = {x ∈ Rn | |xj | ≤ ρ for 1 ≤ j ≤ n}.
If B ⊆ Rn is some set and f : B → C is twice differentiable in some open set containing B, define the matrix H(f, B) = (hjk ), where, provided the suprema exist, hjk = supy∈B |fjk (y)|.
Lemma 3.8. For some ρ > 0, let B = Un (ρ). Suppose that T : Rn → Rm is a differentiable transformation and let JT denote its Jacobian matrix. Let S ⊆ Rm be an open set that contains T (int B). Suppose f : T (B) ∪ S → C is continuous, and define f˜ : B → C
by f˜(x) = f (T (x)). Write α = α(f˜, B) and ∆ = (∆jk ) = ∆(f˜, B), Then

21

(a) Suppose that f is continuously differentiable in S with |fj (y)| ≤ m1 for y ∈ T (int B)
and 1 ≤ j ≤ n. Then kαk∞ ≤ 2ρ m1 sup kJT (x)k1 .
x∈int B

(b) Suppose that f is twice continuously differentiable in S with kH(f, T (int B))k∞ ≤
m2 . Then

αT∆α ≤ 4ρ2 nm2 kαk2∞ sup kJT (x)k1 kJT (x)k∞ and x∈int B
2
T 2
4
2
2
α ∆ α ≤ 16ρ nm2 kαk∞ sup kJT (x)k1 kJT (x)k∞ .
x∈int B

Proof. Suppose JT = (tjk (x)). Observe that for x ∈ int B
f˜j (x) =

m
X

f˜jk (x) =

m X
n
X

trj (x)fr (T (x)),

r=1

trj (x)frs (T (x))tsk (x).

r=1 s=1

From Corollary 3.7 for function f˜, we get
αj ≤ 2ρm1 sup

m
X

x∈int B r=1

|trj (x)|,

which is equivalent to part (a), and
2

∆jk ≤ 4ρ

sup x∈int B

m X
n
X
r=1 s=1

|trj (x)| hrs |tsk (x)|,

where H(f, T (int B)) = (hrs ). Note that the expression

m P
n
P

|trj (x)| hrs |tsk (x)| of the

r=1 s=1

right hand side is the (j, k) element of (JˆT )T H(f, T (int B))JˆT , where JˆT = (|trs |). Claim
(b) now follows on recalling that the ∞-norm is submultiplicative.

4

Truncated gaussian measures

In this section explore the application of Theorem 3.3 to the case where the distribution of X is a truncated gaussian. This is the case that has occurred in the most applications so far.
It will often be convenient to approximate the expectation and pseudovariance of a complex function of a truncated gaussian by their values for the unrestricted gaussian.
The following gives a general principle.
22

Lemma 4.1. Let A be an n × n symmetric positive-definite real matrix. Let f : Rn → C
be a measurable function satisfying b

T

|f (x)| ≤ e n x Ax

(4.1)

for all x ∈ Rn and some b ≥ 0. Let X : R → R be a random variable with density
T

π −n/2 |A|1/2 e−x Ax .
Suppose Ω is a measurable subset of Rn and define p = Prob(X ∈
/ Ω). Then, if p ≤ 43
and n ≥ b + b2 , we have
E(f (X) | X ∈ Ω) − E f (X) ≤ 15 eb/2 p1−b/n .
Moreover, for p ≤ 43 and n ≥ 2b + 4b2 , we have
V(f (X) | X ∈ Ω) − Vf (X) ≤ 112 eb p1−2b/n .
b

T

Proof. By linear transform we can assume that A = 12 I and that |f (x)| ≤ e 2n x x . Let
T

1

µ denote the measure with density π −n/2 e− 2 x x , which is the gaussian measure defined by X after transformation. From the definition of expectation,
 Z

Z
−1
p f (x) dµ −
f (x) dµ .
E(f (X) | X ∈ Ω) − E f (X) = (1 − p)
R

n

n

R −Ω

For any r > 0, since Prob(X ∈
/ Ω ∧ |X| ≤ r) ≤ p, we can bound
Z
Z
|f (x)| dµ ≤ p sup |f (x)| +
|f (x)| dµ.
n

R −Ω

|x|≤r

|x|≥r

Consequently we have

Z
b 2
r
2n pe
E(f (X) | X ∈ Ω) − E f (X) ≤ (1 − p)
+
−1

e

T
b x x
2n

dµ + p

|x|≥r

Z

R

n

e

T
b x x
2n


dµ .

The second integral is easily calculated to be (1 − b/n)−n/2 , provided n > b. The first integral has no closed form; it is (1 − b/n)−n/2 Fn ((1 − b/n)r 2 ), where Fn (u) denotes the upper tail of the χ2 -distribution with n degrees of freedom. From [31, (4.3)] we have that Fn (n + 2u1/2 n1/2 + 2u) ≤ e−u for any u ≥ 0. Consequently, if we put r 2 =
(n + 2u1/2 n1/2 + 2u)/(1 − b/n), we find for any u ≥ 0 and n > b that
E(f (X) | X ∈ Ω) − E f (X)

1/2

≤ (1 − p)−1 peb(1/2+(u/n)

+u/n)/(1−b/n)


+ (1 − b/n)−n/2 (p + e−u ) .

To obtain the version in the theorem statement, use q
b 4(n − b) ln(1/p) − b2
u = (1 − b/n) ln(1/p) +
,
2n
23

which satisfies the equation b(1/2 + (u/n)1/2 + u/n)/(1 − b/n) = −u + ln(1/p). The conditions p ≤ 43 and n ≥ b + b2 imply that the argument of the square root is positive.
Now note that for n ≥ b + b2 , b ≥ 0 the function (1 − b/n)−n/2 e−b/2 < e1/4 is increasing in b and nonincreasing in n, so (1 − b/n)−n/2 e−b/2 < e1/4 . Applying this bound and also u ≥ (1 − b/n) ln(1/p) completes the proof of the first part.
For the second part, we have

V f (X) | X ∈ Ω − Vf (X) = E(f (X)2 | X ∈ Ω) − E f (X)2


− E(f (X) | X ∈ Ω) − E f (X) E f (X) + E(f (X) | X ∈ Ω) .

Note from above that |E f (X)| ≤ E |f (X)| ≤ (1 − b/n)−n/2 . Using the definition of p, we
R
have E(f (X) | X ∈ Ω) ≤ (1 − p)−1 Ω |f (X)| dµ ≤ 4(1 − b/n)−n/2 . Now apply the first part of the lemma to f (X) and f (X)2 , as well as the bound (1 − b/n)−n/2 e−b/2 < e1/4
used earlier. This completes the proof.
Lemma 4.1 is not useful for exponential functions on account of condition (4.1). However, since (4.1) is satisfied by all polynomials (after scaling), the lemma becomes useful in conjunction with Theorem 3.3 for estimating E ef when f has polynomial growth. For convenience, we give the theorem of Isserlis [28] (see [20] for a treatment in modern notation) that tells us how to compute the expectations of polynomials with respect to a multivariate normal distribution.
Theorem 4.2. Let A be a positive-definite real symmetric matrix of order n and let
T
X = (X1 , . . . , Xn ) be a random variable with the normal density π −n/2 |A|1/2 e−x Ax . Let
Σ = (σjk ) = (2A)−1 be the corresponding covariance matrix. Consider a product Z =
Xj1 Xj2 · · · Xjk , where the subscripts do not need to be distinct. If k is odd, then E Z = 0.
If k is even, then
X
EZ =
σji ji · · · σji ji ,
1

2

k−1

k

(i1 ,i2 ),(i2 ,i3 ),...,(ik−1 ,ik )

where the sum is over all unordered partitions of {1, 2, . . . , k} into k/2 disjoint unordered pairs. The number of terms in the sum is (k − 1)(k − 3) · · · 3 · 1.
The following are examples of Theorem 4.2.
E X12 = σ11

2
E X14 = 3σ11

2
E X12 X22 = σ11 σ22 + 2σ12

E X12 X2 X3 = σ11 σ23 + 2σ12 σ13

E X1 X2 X3 X4 = σ12 σ34 + σ13 σ24 + σ14 σ23

24

3
E X16 = 15σ11

4.1

Truncated gaussian measures of full rank

Theorem 4.3. Let c1 , c2 , c3 , ε, ρ1 , ρ2 , φ1 , φ2 be nonnegative real constants with c1 , ε > 0.
Let A be an n × n positive-definite symmetric real matrix and let T be a real matrix such that T TAT = I. Let Ω be a measurable set such that Un (ρ1 ) ⊆ T −1 (Ω) ⊆ Un (ρ2 ), and let f : Rn → C, g : Rn → R and h : Ω → C be measurable functions. We make the following assumptions.
(a) c1 (log n)1/2+ε ≤ ρ1 ≤ ρ2 .

(b) For x ∈ T (Un (ρ1 )), 2ρ1 kT k1 |fj (x)| ≤ φ1 n−1/2 for each j.

(c) For x ∈ Ω, ℜf (x) ≤ g(x). For x ∈ T (Un (ρ2 )), 2ρ2 kT k1 |gj (x)| ≤ φ2 n−1/2 for each j.
T

(d) |f (x)|, |g(x)| ≤ nc3 ec2 x Ax/n for x ∈ Rn .
T

Let X be a random variable with the normal density π −n/2 |A|1/2 e−x Ax . Then, provided
E f (X) and E g(X) are finite and h is bounded in Ω,
Z
T
e−x Ax+f (x)+h(x) dx = (1 + K)π n/2 |A|−1/2 eE f (X) ,
Ω

where, for some constant C depending only on c1 , c2 , c3 , ε,
1

2

|K| ≤ C e 8 φ1 +e

2
−ρ1 /2

1

2

− 1 + (2e 8 φ2 +e

2
−ρ1 /2


− 2 + sup |eh(x) − 1|) eE(g(X)−ℜf (X)) .
x∈Ω

In particular, if n ≥ (1 + c2 )2 and ρ21 ≥ 7 + 2c2 + (3 + 4c3 ) log n, we can take C = 1.
Proof. We will use Lemma 6.2 repeatedly to combine error terms. Change variables by x = T y. Since |T | = |A|−1/2 , we have
Z
Z
T
T
−x Ax+f (x)+h(x)
−1/2
e−y y+f (T y)+h(T y) dy.
e dx = |A|
T

Ω

−1

(Ω)

Suppose ρ ≥ c1 (log n)1/2+ε and let F : Un (ρ) → C be measurable and such that
T
|F (x)| ≤ nc3 ec2 x x/n for x ∈ Rn and kα(F, Un (ρ))k∞ ≤ φn−1/2 . By Theorem 3.3(a),
Z
Z
T
T
−y y+F (y)
′
E(F (Y )|Y ∈Un (ρ))
e−y y dy, e
dy = (1 + K ) e
Un (ρ)

Un (ρ)

2

T

1

where |K ′ | ≤ e 8 φ − 1 and Y has the normal density π −n/2 e−y y .

Define p = Prob(Y ∈
/ Un (ρ)). By standard bounds on the tail of the normal distribu2
−ρ
tion, p ≤ ne /(1 + ρ). Under our assumptions, there is n0 = n0 (c1 , c2 , c3 , ε) such that for
25

2

n ≥ n0 , we have p ≤ 43 , n ≥ c2 + c22 and 15nc3 ec2 /2 p1−c2 /n ≤ e−ρ /2 . Those three conditions are enough that we can apply Lemma 4.1 to the function n−c3 F (y) to conclude that
Z
T
e−y y+F (y) dy = (1 + K ′′ ) π n/2 eE F (Y ) ,
(4.2)
Un (ρ)

1

2

where |K ′′ | ≤ e 8 φ +e

2

−ρ /2

− 1.

We can finish the proof by applying (4.2) to each of the functions f (T y) and g(T y).
For n < n0 we can increase C to make the theorem hold, so assume n ≥ n0 . By Lemma 3.8
we have kα(f (T y), Un (ρ1 ))k∞ ≤ n−1/2 φ1 and kα(g(T y), Un (ρ2 ))k∞ ≤ n−1/2 φ2 . Now we have
Z
Z
T
T
f (T y)+h(T y)−y y e
dy =
ef (T y)−y y dy
Ω
Un (ρ1 )
Z
Z
T
T
f (T y)−y y
+
e dy + (eh(y) − 1)ef (T y)−y y dy
Ω\Un (ρ1 )
Ω

Z
Z
Z
T
T
f (T y)−y y
−
eg(T y)−y y dy
=
e dy + A
Un (ρ1 )
Un (ρ2 )
Un (ρ1 )
Z
T
+ A′ sup |eh(x) − 1|
eg(T y)−y y dy for |A|, |A′| ≤ 1 (4.3)
x∈Ω

R

n

= π n/2 eE f (X) (1 + K1 ) + K2 π n/2 eE g(X) + K3 π n/2 eE g(X) ,
1

2

2
−ρ1 /2

1

2

2
−ρ1 /2

−1, |K2 | ≤ 2e 8 φ2 +e
−2 and |K3 | ≤ supx∈Ω |eh(x) −1|.
where we have |K1 | ≤ e 8 φ1 +e
Finally note that |eE f (X) | = eE ℜf (X) ; the theorem follows.

To establish the final claim, it will suffice to show that for ρ2 ≥ 7 + 2c2 + (3 + 4c3 ) log n we can prove (4.2) with n0 = (1 + c2 )2 . Obviously n ≥ (1 + c2 )2 implies that n ≥ c2 + c22 ,
2
and it also implies that 1 − c2 /n ≥ 34 . The bounds ρ2 ≥ 7 + 3 log n and p ≤ ne−ρ /(1 + ρ)
√
2
imply that p ≤ 43 and also that p ≤ e−ρ n/(1 + 7). Combining these bounds produces
2

the third required inequality 15nc3 ec2 /2 p1−c2 /n ≤ e−ρ /2 , completing the proof.

Theorem 4.4. Let c1 , c2 , c3 , ε, ρ1 , ρ2 , φ1 , φ2 be nonnegative real constants with c1 , ε > 0.
Let A be an n × n positive-definite symmetric real matrix and let T be a real matrix such that T TAT = I. Let Ω be a measurable set such that Un (ρ1 ) ⊆ T −1 (Ω) ⊆ Un (ρ2 ), and let f : Rn → C, g : Rn → R and h : Ω → C be measurable functions. We make the following assumptions.
(a) c1 (log n)1/2+ε ≤ ρ1 ≤ ρ2 .

(b) For x ∈ T (Un (ρ1 )), 2ρ1 kT k1 |fj (x)| ≤ φ1 n−1/3 ≤ 32 for 1 ≤ j ≤ n and
4ρ21 kT k1 kT k∞ kH(f, T (Un (ρ1 )))k∞ ≤ φ1 n−1/3 .
26

(c) For x ∈ Ω, ℜf (x) ≤ g(x). For x ∈ T (Un (ρ2 )), either
(i) 2ρ2 kT k1 |gj (x)| ≤ (2φ2 )3/2 n−1/2 for 1 ≤ j ≤ n, or
(ii) 2ρ2 kT k1 |gj (x)| ≤ φ2 n−1/3 for 1 ≤ j ≤ n and
4ρ22 kT k1 kT k∞ kH(g, T (Un (ρ2 )))k∞ ≤ φ2 n−1/3 .
T

(d) |f (x)|, |g(x)| ≤ nc3 ec2 x Ax/n for x ∈ Rn .

T

Let X be a random variable with the normal density π −n/2 |A|1/2 e−x Ax . Then, provided
Vf (X) and Vg(X) are finite and h is bounded in Ω,
Z
T
1
e−x Ax+f (x)+h(x) dx = (1 + K)π n/2 |A|−1/2 eE f (X)+ 2 Vf (X) ,
Ω

where, for some constant C depending only on c1 , c2 , c3 , ε,
 3 −ρ21 /2
1
−1
|K| ≤ Ce 2 Var ℑf (X) eφ1 +e
3

+ 2eφ2 +e

2
−ρ1 /2



1
− 2 + sup |eh(x) − 1| eE(g(X)−ℜf (X ))+ 2 (Var g(X)−Var ℜf (X)) .
x∈Ω

In particular, if n ≥ (1 + 2c2 )2 and ρ21 ≥ 15 + 4c2 + (3 + 8c3 ) log n, we can take C = 1.
Proof. We will divide the integral in the same fashion as (4.3), and will use estimate (4.2)
again. We also need a similar estimate using Theorem 3.3(b). Lemma 6.2 will be used to combine error terms.
Suppose ρ ≥ c1 (log n)1/2+ε and let F : Un (ρ) → C be measurable and such that
T
|F (x)| ≤ nc3 ec2 x x for x ∈ Rn , and for x ∈ T (Un (ρ)), kα(F, Un (ρ))k∞ ≤ φn−1/3 ≤ 32 and k∆(F, Un (ρ))k∞ ≤ φn−1/3 . By Theorem 3.3(b),
Z
T

1
e−y y+F (y) dy = 1 + K ′ e 2 Var(ℑF (Y )|Y ∈Un (ρ))
Un (ρ)
Z
T
E(F (Y )|Y ∈Un (ρ))+ 12 V(F (Y )|Y ∈Un (ρ))
e−y y dy,
×e
Un (ρ)

3

T

where |K ′ | ≤ eφ − 1 and Y has the normal density π −n/2 e−y y . Similarly to the proof of Theorem 4.3, we can apply Lemma 4.1 to conclude that there is a constant n0 =
n0 (c1 , c2 , c3 , ε) such that for n ≥ n0 ,
Z
T
1
1
e−y y+F (y) dy = (1 + K ′′ e 2 Var ℑF (Y ) ) π n/2 eE F (Y )+ 2 VF (Y ) ,
(4.4)
Un (ρ)
3

where |K ′′ | ≤ eφ +e

2
−ρ /2

− 1.

27

Now consider the expansion given by (4.3). By condition (b) and Lemma 3.8, we have kα(f (T y), Un (ρ1 ))k∞ ≤ n−1/3 ρ1 ≤ 32 , and k∆(f (T y), Un

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
