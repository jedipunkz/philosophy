---
source: "https://arxiv.org/abs/2407.16203v2"
title: "Cutoff for Contingency Table and Torus Random Walks with Low Incremental Correlations"
author: "Zihao Fang, Andrew Heeszel"
year: "2024"
publication: "arXiv preprint / math.PR"
download: "https://arxiv.org/pdf/2407.16203v2"
pdf: "https://arxiv.org/pdf/2407.16203v2"
captured_at: "2026-06-24T11:12:06Z"
updated_at: "2026-06-24T11:12:06Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Rorty Contingency Irony and Solidarity"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Cutoff for Contingency Table and Torus Random Walks with Low Incremental Correlations

- 著者: Zihao Fang, Andrew Heeszel
- 年: 2024
- 掲載情報: arXiv preprint / math.PR
- 情報源: [arxiv](https://arxiv.org/abs/2407.16203v2)
- ダウンロード: https://arxiv.org/pdf/2407.16203v2
- PDF: https://arxiv.org/pdf/2407.16203v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We use the correlation matrix of the generating distribution to determine the mixing time for random walks on the torus $(\mathbb{Z}/q\mathbb{Z})^n$. We present our method in the context of the Diaconis-Gangolli random walk on both the $1 \times n$ and $m \times n$ contingency tables over $\mathbb{Z}/q\mathbb{Z}$. In the $1 \times n$ case, we prove that the random walk exhibits cutoff at time $\dfrac{n q^2 \log(n)}{8 π^2}$ when $q \gg n$; in the $m \times n$ case, where $m, n$ are of the same order, we establish cutoff for the random walk at time $\dfrac{mn q^2 \log(mn)}{16 π^2}$ when $q \gg n^2$. Our method reveals that a general class of random walks on the torus $(\mathbb{Z}/q\mathbb{Z})^n$ has cutoff. If each coordinate of the lifted random walk onto $\mathbb{Z}^n$ has variance $σ^2/n$ in each jump, and the between-coordinate correlations are sufficiently low, then cutoff occurs at time $\dfrac{nq^2 \log(n)}{4π^2 σ^2}$.

## PDF Text

Cutoff for Contingency Table and Torus Random Walks with Low Incremental Correlations arXiv:2407.16203v2 [math.PR] 12 Mar 2026

Zihao Fang ∗

Andrew Heeszel †

Abstract
We use the correlation matrix of the generating distribution to determine the mixing time for random walks on the torus (Z/qZ)n . We present our method in the context of the Diaconis-Gangolli random walk on both the 1 × n and m × n contingency tables over Z/qZ. In the 1 × n case, we prove that the random walk exhibits cutoff at time nq 2 log(n)
when q ≫ n; in the m × n case, where m, n are of the same order, we
8π 2
mnq 2 log(mn)
establish cutoff for the random walk at time when q ≫ n2 . Our method
16π 2
reveals that a general class of random walks on the torus (Z/qZ)n has cutoff. If each coordinate of the lifted random walk onto Zn has variance σ 2 /n in each jump, and the between-coordinate correlations are sufficiently low, then cutoff occurs at time nq 2 log(n)
.
4π 2 σ 2

Contents
1 Introduction
1.1 Motivation and Basic Definitions . . . . . . . . . . . . . . . . . . . . . . . .
1.2 Notations . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
1.3 Main Results . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
1.4 Background . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
1.5 Outline . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .

3
3
5
6
8
9

2 Key Lemmas
2.1 Distinguishing Statistics . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
2.2 Technical Lemmas . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .

9
9
9

∗
†

Department of Mathematics, The Ohio State University, USA, email: fang.735@osu.edu
Department of Statistics, The Ohio State University, USA, email: heeszel.1@osu.edu

1

3 Lower Bounds
3.1 Setup . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
3.2 Representations and Eigenvalues . . . . . . . . . . . . . . . . . . . . . . . . .
3.3 Mixing Time Lower Bound . . . . . . . . . . . . . . . . . . . . . . . . . . . .
3.3.1 Forming a Distinguishing Statistic . . . . . . . . . . . . . . . . . . . .
3.3.2 Mean and Variance under Ht (0, .) . . . . . . . . . . . . . . . . . . . .

13
13
13
14
14
15

4 1 × n Contingency Tables
4.1 Introduction . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
4.2 Eigenvalues for {Xt } and {Ys } . . . . . . . . . . . . . . . . . . . . . . . . . .
4.3 Strategy for Mixing Time Upper Bound . . . . . . . . . . . . . . . . . . . . .
4.4 Bounding Decay in Characteristic Function for Large θ . . . . . . . . . . . .
π
1
. . . . . . . . . . . . . . . . . . . . . . .
4.4.1 Regime 1: √ < ∥θ∥∞ ≤
q
2
π
≤ ∥θ∥∞ ≤ π . . . . . . . . . . . . . . . . . . . . . . . .
4.4.2 Regime 2:
2
4.4.3 Combining the Two Regimes of Large θ . . . . . . . . . . . . . . . .
4.5 Bounding Decay in Characteristic Function for Small θ . . . . . . . . . . . .
4.5.1 Decomposing the Characteristic Function in ℓ2 -Bound . . . . . . . . .
4.5.2 Simplifying the Quadratic Form y T Γy . . . . . . . . . . . . . . . . . .
4.5.3 Bounding the Sum for Small y . . . . . . . . . . . . . . . . . . . . . .
4.6 Mixing Time Upper Bound . . . . . . . . . . . . . . . . . . . . . . . . . . . .
4.7 Mixing Time Lower Bound . . . . . . . . . . . . . . . . . . . . . . . . . . . .

17
17
18
19
20

5 Generalization for Mixing Time Upper Bound
5.1 Setup . . . . . . . .√. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
q
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
5.2 Regime 1: ∥y∥∞ >
2π
√
q
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
5.3 Regime 2: ∥y∥∞ ≤
2π
5.4 Cutoff for a Family of Random Walks on Znq . . . . . . . . . . . . . . . . . .
5.5 Discussion of the Regime q → ∞ . . . . . . . . . . . . . . . . . . . . . . . .

26
27

6 m × n Contingency Tables
6.1 Inverting the Correlation Matrix . . . . . . . . . . . . . . . . . . . . . . . . .
6.2 Bounding Decay in Characteristic Function for Large θ . . . . . . . . . . . .
1
π
6.2.1 Regime 1: √ < ∥θ∥∞ ≤
. . . . . . . . . . . . . . . . . . . . . . .
q
4
π
6.2.2 Regime 2:
≤ ∥θ∥∞ ≤ π . . . . . . . . . . . . . . . . . . . . . . . .
4
6.2.3 Combining the Two Regimes of Large θ . . . . . . . . . . . . . . . .
6.3 Bounding Decay in Characteristic Function for Small θ . . . . . . . . . . . .
6.3.1 The Schur Complement of a Submatrix of Ψ . . . . . . . . . . . . . .
6.3.2 The Correlation Condition . . . . . . . . . . . . . . . . . . . . . . . .
6.4 Mixing Time Upper Bound . . . . . . . . . . . . . . . . . . . . . . . . . . . .

35
36
37

2

20
21
22
22
23
23
24
26
26

27
30
32
33

37
39
41
41
41
44
46

6.5

Mixing Time Lower Bound . . . . . . . . . . . . . . . . . . . . . . . . . . . .

7 Acknowledgments

47
47

1

Introduction

1.1

Motivation and Basic Definitions

We are interested in showing that a class of torus random walks has cutoff primarily based on the correlation matrix of the incremental distribution of the walk. We consider a continuoustime random walk {Xt }t≥0 on the torus (Z/qZ)n when q is large with respect to n, and the correlations between coordinates in the incremental distribution are sufficiently small. We provide sufficient conditions for cutoff based on the correlation matrix of Xt when lifted to Zn and the decay of its characteristic function. This setup covers various random walks currently studied in the literature, including the Diaconis-Gangolli random walk on contingency tables over Zq := Z/qZ. We will use the Diaconis-Gangolli random walk as a primary model to present our method.
A contingency table is a matrix of integer entries with prescribed row and column sums.
Enumerating the contingency tables with given marginals is closely related to counting a variety of other combinatorial objects, including permutations with descent restrictions, double cosets of the symmetric groups, certain quantities about the induced representations of the symmetric groups, and some special types of Young tableaux. The name “contingency table” is commonly used in the context of statistics and is linked to Fisher’s exact test of association and goodness of fit, along with the approximate χ2 test. Random sampling from the set of contingency tables with fixed margins can be used to form an approximate p-value for the test of association when the χ2 approximation is not well met.
Markov chain Monte Carlo (MCMC) methods can be used to randomly sample and approximately count contingency tables with fixed row and column sums. One of the famous
MCMC algorithms was described in [7] and [11], giving rise to the Diaconis-Gangolli random walk on contingency tables over Zq . Because the sums of rows and columns are fixed, we can view the dynamics as a random walk on the torus by ignoring the last row and column.
We will first use the Diaconis-Gangolli random walk on 1 × n contingency tables over Zq to outline how we can use incremental correlations to establish cutoff for torus random walks
(Theorem 1). Then we will present our method in full generality to prove cutoff for a general class of random walks on the torus Znq satisfying assumptions on the correlations between the coordinates of a single increment (Theorem 3). Finally, we apply our method to the
Diaconis-Gangolli random walks on the more complicated m × n contingency tables over Zq and illustrate the cutoff phenomenon in this model (Theorem 2).
In particular, the last part of this paper could be viewed as a complement to a previous work by Nestoridi and Nguyen [25]. They established cutoff for the discrete-time DiaconisGangolli random walk on n × n contingency tables over Zq , as n → ∞. When m = n, our
Theorem 2 gives cutoff for this model in continuous time when q is sufficiently large relative

3

!
p log(n)
to mn = n2 , while Nestoridi and Nguyen needed log(q) = o to achieve cutoff.
log log n
Note that we need q to grow faster than n2 to establish cutoff, whereas Nestoridi and Nguyen required q to be very small.
In addition, our general result Theorem 3 expands on examples of cutoff for non-local random walks on Znq . Previously, Nestoridi [24] found mixing times and regimes of cutoff for a non-local random walk on Zn2 = (Z/2Z)n such that k coordinates are randomly selected and flipped in a single iteration. In contrast, our results show cutoff for random walks on
Znq with non-nearest-neighbor steps when both n → ∞ and q = q(n) → ∞.
We now define the dynamics of the Diaconis-Gangolli random walk, and note that we will study this model in continuous time. We start with the 1 × n contingency tables.
Definition 1. A 1 × n contingency table over Zq is an n-dimensional vector with entries from Zq such that the sum of all the coordinates is fixed.
We will consider the following rate-1 continuous-time random walk defined on 1 × n contingency tables over Zq . Each jump of this random walk is governed by the following rule.
(i) Select two coordinates uniformly at random without replacement.
(ii) Flip a fair coin. If heads, increment the first coordinate selected by +1 and increment the second coordinate selected by −1; if tails, increment the first coordinate selected by −1 and increment the second coordinate selected by +1.
The analogous definition and dynamics for m × n contingency tables are the following.
Definition 2. An m × n contingency table over Zq is an m × n matrix with entries from Zq n
such that the row sums (r1 , . . . , rm ) ∈ Zm q and column sums (c1 , . . . , cn ) ∈ Zq are each fixed.
We will study the rate-1 continuous-time Diaconis-Gangolli random walk on m × n contingency tables over Zq , each of whose jumps is governed by the following rule:
(i) Select two distinct rows i < i′ and two distinct columns j < j ′ independently and uniformly at random.
(ii) Flip a fair coin, and consider the m × n matrix T = T (i, i′ , j, j ′ ) whose (i, j)-th and
(i′ , j ′ )-th entries are +1, (i, j ′ )-th and (i′ , j)-th entries are −1, and zero elsewhere,


1 −1
T =
.
−1 1
If heads, add T to the current matrix; if tails, add −T to the current matrix.

4

As mentioned above, we specifically consider the case when m has the same order n. More precisely, we assume that there are universal constants κ1 , κ2 such that 0 < κ1 ≤ 1 ≤ κ2 and
κ1 n ≤ m ≤ κ2 n.
For convenience, we will also assume that κ1 n, κ2 n ∈ Z, and κ1 (n − 1) ≤ m − 1 ≤ κ2 (n − 1).
Since we will take n → ∞, these assumptions will not affect our results.
Both of these two random walks converge to the uniform distribution on the respective state space with respect to the total variation distance. More specifically, if the random walk starts at x ∈ Ω, where Ω is the state space, then we can denote by Ht (x, y) the probability of the walk going from x to y ∈ Ω at time t ≥ 0. Let π be the uniform measure on Ω, and write d(t) := max ∥Ht (x, ·) − π∥TV
x∈Ω

where
∥µ − ν∥TV := max |µ(A) − ν(A)|.
A⊂Ω

is the total variation distance between two probability measures µ, ν on Ω. By the convergence theorem of Markov chains, d(t) → 0 as t → ∞. Our goal is to determine the rate of convergence, which could be quantified by the ε-mixing time tmix (ε) := inf{t ≥ 0 : d(t) ≤ ε}
for ε ∈ (0, 1). In particular, we show that both the walks on 1 × n and m × n contingency tables exhibit cutoff as n → ∞. We say that a sequence of Markov chains indexed by n = 1, 2, . . . has cutoff if for all ε ∈ (0, 1),
(n)

lim

tmix (ε)

n→∞ t(n) (1 − ε)
mix

= 1.

In other words, for large enough n, there will be a sharp transition for d(t) from almost 1 to almost 0.

1.2

Notations

Throughout the paper, we will use Zq to denote the cyclic group Z/qZ of q elements. We do not need q to be a prime number, but we intrinsically require q = q(n) to increase along with n. In the case of m × n contingency table, m = m(n) should also be considered as a function of n such that κ1 n ≤ m ≤ κ2 n.
We will use the standard asymptotic notation:
• f = O(g), or equivalently, f ≲ g, if lim sup f /g < ∞, and
• f = o(g), or equivalently, f ≪ g, if lim f /g = 0.
All the limits are taken as n → ∞ unless otherwise specified. Lastly, we will write N =
{0, 1, 2, 3, . . . } for the set of nonnegative integers.
5

1.3

Main Results

Our main results are the following.
Theorem 1. Let ε ∈ (0, 1) be arbitrary. For the rate-1 continuous-time Diaconis-Gangolli random walk on 1 × n contingency tables over Zq , we have the following.
(a) Let




4 + 4e2/e
2
√
−1 .
α := inf z ≥ 1 : 4ε ≥ exp z

−1
nq 2 log(αn)
1
When t =
1−
,
8π 2
3q d(t) ≤ ε + o(1)
as n → ∞ with q ≫ n.
p
(b) On the other hand, as n → ∞ with q ≫ log(n),



 −1

nq 2
1
ε −1
tmix (ε) ≥ 2 log(n) + log 1 −
+ log
.
8π
n
5

Theorem 1 implies that as n → ∞ with q ≫ n, this family of continuous-time random nq 2 log(n)
, with a cutoff walks on 1 × n contingency tables exhibits cutoff at the mixing time
8π 2
window of order O(nq 2 ).
Theorem 2. Let ε ∈ (0, 1) be arbitrary, and suppose there exist universal constants κ1 , κ2 >
0 so that 0 < κ1 ≤ 1 ≤ κ2 and κ1 n ≤ m ≤ κ2 n. For the rate-1 continuous-time DiaconisGangolli random walk on m × n contingency tables over Zq , we have the following.
(a) Let

 2


κ2 (6 + 12e4/e + 12e8/e + 6e32/e )
2
α := inf z ≥ 1 : 4ε ≥ exp
−1 .
κ1 z 1/4

−1
4
mnq 2 log(α(m − 1)(n − 1))
When t =
1−
,
16π 2
3q d(t) ≤ ε + o(1)
as n → ∞ with q ≫ mn.
p
(b) On the other hand, as n → ∞ and q → ∞ with q ≫ log(n),

 −1

mnq 2
ε −1
tmix (ε) ≥
log ((m − 1)(n − 1)) + log
.
16π 2
5
6

Theorem 2 implies that as n → ∞ with q ≫ mn, this family of continuous-time random mnq 2 log(mn)
walks on m × n contingency tables exhibits cutoff at the mixing time
, with a
16π 2
2
cutoff window of order O(mnq ).
As mentioned in Section 1.1, we also establish cutoff for a general class of random walks on the torus Znq . We will now give a shortened version of the generalization, and see Section 5.4
(n)
for the full statement of the theorem. In our setup, we assume that {Xt }t≥0 is an irreducible rate-1 continuous-time random walk on the torus Znq with symmetric generating distribution
µ(n) , and let Y (n) be a random vector defined on Zn with the law of µ(n) lifted to Zn . We
σ2
, with σ 2 = σ 2 (n), assume the coordinates of Y (n) have a marginal incremental variance n
and Y (n) has correlation matrix Γn . We will also assume that ∥Y (n) ∥1 ≤ r = r(n), the
(n) ∞
}n=1 meets sequence {Γn }∞
n=1 satisfies the Correlation Condition, and the sequence {Y
the Decay Condition in terms of its characteristic function (see Section 5 for the formal definitions of both conditions). In our regime, we assume that both n → ∞ and q → ∞.
Based on this setup, we can now state a shortened version of the general theorem.
Theorem 3. Let {Xt }(n) be an arbitrary family of irreducible symmetric continuous-time
Markov chains on the torus Znq with the setup above, assuming Y (n) has correlation matrix
Γn . Let ε ∈ (0, 1) be arbitrary. Then we have the following.
(a) Suppose that {Y (n) } meets the Decay Condition and {Γn }∞
n=1 meets the correlation condition. Then there exists α = α(ε) so that when
−1

nq 2 log(αn)
r2
t=
1−
4π 2 σ 2
12q
Then, d(n) (t) ≤ ε + o(1)
as n → ∞, q → ∞ with r2 ≪ q.
1
r2
1
, and 2 ≪
, log(n)
q log(n)

 −1

ε −1
nq 2
(n)
tmix (ε) ≥ 2 2 log(n) + log
4π σ
5

(b) On the other hand as n → ∞, q → ∞ with supk̸=ℓ |Γn,kℓ | ≪

(n)

We see that if the family of Markov chains {Xt }∞
n=1 meets the assumptions of Theorem
2
nq log(n)
nq 2
(n)
, with a
cutoff window of order
.
3, then {Xt }∞
has cutoff at the mixing time n=1
4π 2 σ 2
σ2
Here our primary regularity conditions rely on the correlations of the walk being sufficiently low, and the generating distribution of the walk having a fast enough rate of decay outside a neighborhood of 0. We will postpone a more precise description of the general result to the end of Section 5.
7

1.4

Background

This sampling method for contingency tables via the described Markov chain had been applied in practice, but it was first formally discussed by Diaconis and Gangolli [7] and
Diaconis and Saloff-Coste (page 373 in [11]), where the entries of the contingency tables were nonnegative integers. For m × n contingency tables over N with m, n fixed, Diaconis and showed that the mixing time of the Markov chain is of order N 2 , where N :=
P Saloff-Coste
P
ci =
ri is the sum of all the entries of the table. Hernek [16] specifically investigated the model on 2 × n contingency tables and proved that the mixing time is polynomial in n and N . A modified version of the chain on m × n contingency tables was considered by
Chung, Graham, and Yau [4], and they showed that the mixing time is polynomial in m, n, and N .
Through the adjacency matrix of a given graph, the dynamics of the Diaconis-Gangolli random walk is closely related to the simple switching operations of the edges. Namely, we pick two non-incident edges i ↔ j, i′ ↔ j ′ uniformly at random at each step, and then we delete the existing two edges and add the edges i ↔ j ′ , i′ ↔ j. The switch chain formed by these dynamics could be used to generate configuration model random graphs or to generate random graphs with prescribed degree requirements in random networks. Recently,
Tikhomirov and Youssef [27] showed that the switch chain on d-regular bipartite graphs on n vertices has a mixing time of order (nd)2 log(nd), under the assumption that the degree d ≥ 3 is at most polynomial in n. For references in the context of random networks, see, for example, Milo, Kashtan, Itzkovitz, Newman, and Alon [22], Rao, Jana, and Bandyopadhyay
[26], and the survey by Wormald [28].
Besides the Diaconis-Gangolli model, there have been other algorithms and Markov chains to sample contingency tables. Dyer, Kannan, and Mount used polytopes to establish an algorithm to sample m × n contingency tables within polynomial time with respect to m, n, and log N . Morris [23] later simplified and improved this algorithm. Dyer and Greenhill [13]
proved pre-cutoff for a Markov chain on 2 × n contingency tables, known as the heat-bath chain, via the method of path coupling. This result was later generalized by Matsui, Matsui, and Ono [21] to 2 × 2 × · · · × 2 × J contingency tables, and then by Cryan, Dyer, Goldberg,
Jerrum, and Martin [5] to contingency tables with a constant number of rows.
In some other Markov chain models, the cutoff phenomenon has also been widely observed and studied. Many of the first examples of cutoffs were found in the context of card shuffling, or random walks on the symmetric group. For example, the pioneering paper by Diaconis and Shahshahani [10] established cutoff for the random transposition walk on symmetric groups. Aldous and Diaconis [1] were the first to use the term “cut-off phenomenon”. Bayer and Diaconis [2] proved the cutoff for the riffle shuffle chain, a result later known as “seven shuffles suffice”.
Cutoffs on other models have also been investigated. For example, Ding, Lubetzky, and
Peres [12] gave a general cutoff result for birth-and-death chains. Diaconis and Saloff-Coste
[9] proved a similar result in terms of separation cutoff. Lubetzky and Sly [20] found cutoff for random walks on random regular graphs. Lubetzky and Peres [19] established cutoff for random walks on Ramanujan graphs. Ganguly, Lubetzky, and Martinelli [14] obtained
8

cutoff for the East model. Numerous other examples could be found in the book by Levin and Peres [18].

1.5

Outline

Section 2 contains a few lemmas that will be useful later. We prove a mixing time lower bound for more general random walks on the torus Znq in Section 3, and this result can be applied to both of the specific contingency table random walks. Section 4 gives detailed proof of the upper mixing time bound for the random walk on 1 × n contingency tables over
Zq . We use the same proof strategy in Section 5 to prove an upper mixing time bound for more general random walks on the torus Znq , and together with our lower bound from Section
3, we establish cutoff for a class of general random walks on the torus. Finally, in Section 6, we apply an analogous argument to show cutoff for the random walk on m × n contingency tables.

2

Key Lemmas

2.1

Distinguishing Statistics

First, we need the following lemma from [18] in order to prove the lower mixing time bounds.
Lemma 1 (Proposition 7.12, [18]). Let µ and ν be two probability distributions on Ω, and let f be a real-valued function on Ω. If
|Eµ (f ) − Eν (f )| ≥ bη
where η 2 :=

Varµ (f ) + Varν (f )
, then
2
∥µ − ν∥TV ≥ 1 −

2.2

4
.
4 + b2

Technical Lemmas

Lemma 2. For α ≥ 1 and n ≥ 1, n ∈ N, n
X
k
1 + e2/e
(αn)− k+1 ≤ √
.
α
k=1

Proof. We break the sum into two parts and bound them, respectively, as follows. We can write
√
⌊ n⌋
X
√
k
1
(αn)− k+1 ≤ n(αn)−1/2 = √
α
k=1
9

and

n
X

(αn)

k
− k+1

≤ (n −

√

√1

n)(αn)

1√
−1
1+ n

≤ n(αn)

1√
−1
1+ n

√
k=⌈ n⌉

n n e2/e
≤ √ ≤ √ .
α
α

√

For the last inequality, we used that for x > 0, x1/ x is maximized as e2/e at x = e2 by elementary calculus. Thus n
X

k
1
e2/e
1 + e2/e
.
(αn)− k+1 ≤ √ + √ ≤ √
α
α
α
k=1

Lemma 3. For all n ≥ 1, n ∈ N,
 ℓ k n X
n 
X
1 ℓ+1 k+1
≤ 1 + 2e4/e + 2e8/e + e32/e .
2
n
ℓ=1 k=1
If κ1 n ≤ m ≤ κ2 n with 0 < κ1 ≤ 1 ≤ κ2 and m ≥ 1, m ∈ N, then
 ℓ k n 
m X
X
1 ℓ+1 k+1
κ2 (1 + 2e4/e + 2e8/e + e32/e )
≤ 2
.
mn
κ1
ℓ=1 k=1
Proof. We will prove the first bound in the claim, and the second bound will follow as a straightforward corollary. For any n ∈ N, we can write k+ℓ+1
 ℓ k
 − (k+1)(ℓ+1)
n n
n X
n 
n n
X
2(k+ℓ+1)
1 X X (k+1)(ℓ+1)
1 ℓ+1 k+1 X X 1
1
=
=
n n2
n2 n2
n2 ℓ=1 k=1
ℓ=1 k=1
ℓ=1 k=1
2(k+ℓ+1)

Note that n (ℓ+1)(k+1) is decreasing in both ℓ and k. We break the sum into four parts. First, for ℓ ≤ n1/4 and k ≤ n1/4 ,
1/4

1/4

⌊n ⌋ ⌊n ⌋
2(k+ℓ+1)
1 X X (k+1)(ℓ+1)
1
n
· n1/4 · n1/4 · n3/2 = 1.
≤
n2 ℓ=1 k=1
n2

(1)

Next, for ℓ ≥ n1/4 and k ≥ n1/4 ,
1
n2

n
X

n
X

n

2(k+ℓ+1)
(k+1)(ℓ+1)

ℓ=⌊n1/4 ⌋ k=⌊n1/4 ⌋

1
≤ 2
n

n
X

n
X

n

4(k+ℓ)
ℓk

ℓ=⌊n1/4 ⌋ k=⌊n1/4 ⌋

≤

1
−1/4
· n · n · n8n
≤ e32/e
2
n

(2)

−1/4

as x8x is maximized as e32/e at x = e4 by elementary calculus. Now it suffices to bound the sum for ℓ ≤ n1/4 and k ≥ n1/4 , and the cases ℓ ≥ n1/4 and k ≤ n1/4 would follow in an analogous way. We consider the scenario of ℓ = 1 separately. When ℓ = 1, we have
1
n2

n
X
k=⌊n1/4 ⌋

n

k+2
k+1

1
=
n

n
X

1

n k+1 ≤

k=⌊n1/4 ⌋

10

1
−1/4
· n · nn
≤ e4/e .
n

(3)

If 2 ≤ ℓ ≤ n1/4 , then
⌊n1/4 ⌋

1 X
n2 ℓ=2

n
X

n

2(k+ℓ+1)
(k+1)(ℓ+1)

k=⌊n1/4 ⌋

2(n1/4 +3)
1
−1/4
1/4
≤ e8/e .
≤ 2 · n · n · n 3n1/4 = n−1/12 · n2n n

(4)

Combining (1), (2), (3), and (4) establishes the first bound in the claim.
Now for the second bound, we can write k
ℓ
 ℓ k
 ℓ k
 ℓ+1
κ2 n X
κ2 n  2
κ2 n X
κ2 n 
m X
n 
X
k+1
1 ℓ+1 k+1 X
κ2 n/m ℓ+1 k+1
κ22 n X
1
≤
.
≤
mn
(κ2 n)2
m ℓ=1 k=1 (κ2 n)2
ℓ=1 k=1
ℓ=1 k=1
Since m ≥ κ1 n,

κ22 n
κ2
≤ 2 and the result follows from the first bound.
m
κ1

We note that in Lemmas 2, 3, the constants e2/e , e4/e , etc. are chosen for convenience to give explicit uniform bounds.
P
2
Lemma 4. Let F (x) := N
j=−N exp(−c(j − x) ), where N ∈ N and c > 0 are constants.
Then there exists a constant c0 > 0 independent of N , such that whenever c > c0 , F (x)
attains its global maximum at x = 0.
Proof. First, we observe that for any x ∈ (−1/2, 1/2], we have
F (x) > F (x + 1) > F (x + 2) > F (x + 3) > · · ·
and
F (x) ≥ F (x − 1) > F (x − 2) > F (x − 3) > · · ·
by comparing the summands in F (x + k) and F (x + k + 1), k ∈ Z. Thus, it suffices to maximize F on the interval (−1/2, 1/2].
Let Bj (δ) := {x ∈ R : |j − x| ≤ δ} for each j. We look at the second derivative of F ,
′′

F (x) =

N
X

2c(2c(j − x)2 − 1) exp(−c(j − x)2 ).

j=−N

√
S
′′
We will only consider c > 0 that is not too small. If x ̸∈ N
j=−N Bj (1/ 2c), then F (x)
is necessarily positive, since all the summands above are positive. Combining this
√
√ with the previous observation, we only need to maximize
F
on the interval
[−1/
2c,
1/
2c].
√
′
We will show that F (x) ≤ 0 for x ∈ [0, 1/ 2c], which will imply the desired result since
F is an even function. We can write
F ′ (x) =

N
X

2c(j − x) exp(−c(j − x)2 )

j=−N
−cx2

= −2c xe

N
X
2
2
+
(j + x)e−c(j+x) − (j − x)e−c(j−x)
j=1

11

!

and we need to show
2

xe−cx +

N
X

2

2

(j + x)e−c(j+x) − (j − x)e−c(j−x) ≥ 0.

j=1

√
Note that for all c not too small and any fixed x ∈ [0, 1/ 2c], the function
2

(z + x)e−c(z+x) − (z − x)e−c(z−x)

2

is negative and increasing in z for z ≥ 1. To see that this function is negative, one can apply
2
the Mean Value Theorem to the function ze−cz of z; and to prove that it is increasing, one
2
2
d can apply the Mean Value Theorem to the derivative dz
(ze−cz ) = (1 − 2cz 2 )e−cz . It follows that
N
X
2
2
(j + x)e−c(j+x) − (j − x)e−c(j−x)
j=1

≥

∞
X

2

2

(j + x)e−c(j+x) − (j − x)e−c(j−x)

j=1

≥ (1 + x)e

−c(1+x)2

−c(1−x)2

− (1 − x)e

Z ∞
+

2

2

(z + x)e−c(z+x) − (z − x)e−c(z−x) dz

1

= (1 + x)e

−c(1+x)2

−c(1−x)2

− (1 − x)e

1
2
2
+ (e−c(1+x) − e−c(1−x) ).
2c

Then it suffices to show that




1
1
2
−c(1+x)2
−cx2
e
− 1−x+
e−c(1−x) ≥ 0.
xe
+ 1+x+
2c
2c
2

1
Note that for all c not too small, the function g(x) := (1 + x + 2c
)e−c(1+x) is convex on the
√
√
√
interval [−1/ 2c, 1/ 2c]. So for x ∈ [0, 1/ 2c],

x
′
−c(1−x)2
2
g(x) − g(−x) ≥ 2xg (−x) = 4cxe
.
x − 2x + 1 −
2c
Now, it remains to show that

x
2
2
e−cx + 4ce−c(1−x) x2 − 2x + 1 −
≥ 0.
2c
√
For x ∈ [0, 1/ 2c], we can write

x
2
2
e−cx + 4ce−c(1−x) x2 − 2x + 1 −
2c
 x
−cx2
−c(1−x)2
≥e
+ 4ce
−
2c
−cx2
−c(1−x)2
=e
− 2xe p
2
−cx2
≥e
− 2/c e−c(1−x) .

1 log(c/2)
The last expression above is strictly positive for all x < +
, and this completes
2
4c the proof.
12

3

Lower Bounds

3.1

Setup

We provide a general method for mixing time lower bounds for a family of random walks on the torus Znq . We parameterize the torus Znq by (Z ∩ [−q/2, q/2))n , i.e. these are the representatives we will use for elements in Znq .
Let {Ys }s∈N be an arbitrary irreducible symmetric discrete-time Markov chain on Znq with transition matrix P , such that P(Ys = ys |Ys−1 = ys−1 ) = µ(ys − ys−1 ). Suppose that
∥Y ∥1 ≤ r for some real number r in the L1 norm, and ∥Y ∥∞ < q/2 in the sup norm. We set Y0 = 0. Symmetry of {Ys } implies µ(y) = µ(−y) for all y ∈ Znq . Let Y ∈ Zn be a lattice-valued random vector with law µ mapped to (Z ∩ [−q/2, q/2))n . We will assume Y
σ2
σ2
has the same marginal variance for each coordinate. Write Var(Y ) =
Γ where Γ is n
n the correlation matrix of Y . Let {Xt }t≥0 be the rate-1 continuous-time Markov chain on Znq with rate matrix Q = P − I and transition kernel Ht = exp(tQ). Set X0 = 0.
In this section for lower bounds, we will further assume that as n → ∞, sup |Γij | ≪
i̸=j

1
log(n)

r2
1
≪
.
2
q log(n)

and

(5)

Symmetry and irreducibility of P also imply that Ht will converge to the uniform distribution on Znq . Additionally, if f is an eigenfunction for P with eigenvalue λ, then f is also an eigenfunction for Ht with eigenvalue exp(t(λ − 1)).

3.2

Representations and Eigenvalues

Definition 3. Let p be any probability measure on Znq , and let ρ be a representation of Znq .
The Fourier transform p̂ of p is defined as
X
p̂(ρ) :=
p(g)ρ(g).
g∈Zn q

The irreducible representations of the torus Znq are all 1-dimensional and of the form

ρy (g) = exp

2πi⟨y, g⟩
q


,

y ∈ Znq

P
for g ∈ Znq , where ⟨y, g⟩ = n−1
j=1 gi yi is the inner product of y and g (see, for example, [25]).
n
Let S ⊂ Zq be the symmetric set of generators for {Ym }. Since µ(g) = P(Y1 = g) for

13

g ∈ S and zero elsewhere in Znq , then we can write
µ̂(ρy ) =

X

µ(g)ρy (g)

g∈Zn q


 
2πy
,g
=
P(Y1 = g) exp i q
g∈S


2πy
= ΦY
q
X

where ΦY is the characteristic function of Y . Note that when put into ΦY , y ∈ Znq should be understood as y ∈ (Z ∩ [−q/2, q/2))n .
By Theorem 3 of [10] (a more detailed proof could be found in Theorem 6 of [6]), the
Fourier transforms {µ̂(ρy ) : y ∈ Znq } are precisely the eigenvalues of the discrete-time process
{Ys }, and the characters are the eigenfunctions of P . Thus the eigenvalues {λj : 1 ≤ j ≤ q n }
of the continuous-time process {Xt } are
{λj : 1 ≤ j ≤ q n } = {exp(ΦY (2πy/q) − 1) : y ∈ (Z ∩ [−q/2, q/2))n }.
Since µ is symmetric, ΦY is a real-valued function. Therefore, the real parts 
of the char
2π⟨y, g⟩
n acters are exactly the eigenfunctions of P . Hence, for each y ∈ Zq , fy (g) := cos q
is an eigenfunction of P with its corresponding eigenvalue ΦY (2πy/q).

3.3

Mixing Time Lower Bound

We now aim to apply Lemma 1 to obtain a mixing time lower bound.
3.3.1

Forming a Distinguishing Statistic

Let


ψi (x) := cos

2πxi q



for (Z ∩ [−q/2, q/2))n . We form a distinguishing statistic
ψ(x) :=

n
X

ψi (x) =

i=1

n
X
i=1


cos

2πxi q



for x ∈ (Z ∩ [−q/2, q/2))n . We will use ψ as the real-valued function in Lemma 1.
Note that each ψi is an eigenfunction of the transition matrix P . Since ψ is the sum of eigenfunctions whose corresponding eigenvalues are not 1, we know that
Eπ (ψ) = 0

14

(6)

(see, for example, Lemma 12.3 in [18]). Since π is uniform over Znq , this implies that each
ψi is independent of each ψj for i ̸= j under the stationary distribution. Also, for each i = 1, . . . , n,



1 1
4πxi
1
2
+ cos
= .
Eπ (ψi ) = Eπ
2 2
q
2
It follows that
Varπ (ψ) =
3.3.2

n
.
2

(7)

Mean and Variance under Ht (0, .)

Let δi be the vector whose i-th coordinate is 1 and zero elsewhere for notation. We know that each ψi is an eigenfunction of Ht with a corresponding eigenvalue of
  


2πδi exp t ΦY
−1
.
q
By the fourth-order Taylor expansion of the characteristic function ΦY , we obtain that for any i,




  
2π 2 tσ 2
2πδi
−1
.
(8)
exp −
≤ exp t ΦY
nq 2
q


2π(δi + δj )
The fourth-order Taylor expansion of ΦY
for i ̸= j gives q


  
2π(δi + δj )
−1
exp t ΦY
q
!!
2π(δi +δj )
4
E((
·
Y
)
)
2π 2 σ 2
q
≤ exp t −
(δi + δj )T Γ (δi + δj ) +
(9)
nq 2
24
!!
2π(δ +δ )
E(( iq j · Y )4 )
4π 2 σ 2
= exp t −
(1 + Γij ) +
nq 2
24
Recall the assumption ∥Y ∥1 ≤ r, so by Cauchy-Schwarz, the above is upper bounded by
!!
2π(δ +δ )
π 2 r2 E(( iq j · Y )2 )
4π 2 σ 2
≤ exp t −
(1 + Γij ) +
nq 2
6q 2
 


4π 2 σ 2
π 2 r2
= exp t −
(1 + Γij ) 1 −
nq 2
3q 2
  2 2


4π σ
π 2 r2
≤ exp −t
1 − sup |Γkℓ |
1−
.
nq 2
3q 2
k̸=ℓ

15

By an argument similar to (9),
  


2π(δi − δj )
exp t ΦY
−1
q


  2 2
π 2 r2
4π σ
1 − sup |Γkℓ |
1−
.
≤ exp −t nq 2
3q 2
k̸=ℓ
Recall that X0 = 0. Let νt be the law of Ht (0, ·). By (8), we have


2π 2 tσ 2
Eνt (ψ) ≥ n exp −
.
nq 2

(10)

(11)

nq 2 log(γn)
We now set our time to be t =
, where γ > 0 is independent of n. Plugging this
4π 2 σ 2
t into (11), we obtain
√
n
Eνt (ψ) ≥ √ .
(12)
γ
Next, we bound the variance of ψ under νt at the time set above. We write n
X






2πXt (i)
Varνt (ψ) =
Varνt cos q
i=1
 



n X
X
2πXt (i)
2πXt (j)
Covνt cos
+
, cos q
q i=1 j̸=i

(13)

where Xt (i) is the ith coordinate of Xt . Since cosine is bounded by 1 in absolute value, n
X

 

2πXt (i)
Varνt cos
≤ n.
q i=1

(14)

It remains to bound the covariance terms in (13). For any fixed i, j, we have by (9) and (10)
that



 
2πXt (j)
2πXt (i)
cos
Eνt cos q
q

 

 

1
2π(Xt (i) + Xt (j))
2π(Xt (i) − Xt (j))
=
Eνt cos
+ Eνt cos
2
q q




2 2
π r
≤ exp − log(γn) 1 − sup |Γkℓ |
1−
3q 2
k̸=ℓ


 (1−supk̸=ℓ |Γkℓ |) 1− π2 r22
3q
1
=
.
γn

16

It follows from (8) and the above that for any i,



 
X
2πXt (j)
2πXt (i)
, cos
Covνt cos q
q j̸=i




 (1−supk̸=ℓ |Γkℓ |) 1− π2 r22
3q
1
1 
≤ n
−
γn
γn


2 2
2 2
supk̸=ℓ |Γkℓ |+ π r2 −supk̸=ℓ |Γkℓ | π r2
−1
3q
3q
(γn)
−1
=γ
where the last expression vanishes to 0 as n → ∞ by the assumptions (5). Hence, for sufficiently large n,
 



X
2πXt (i)
2πXt (j)
Covνt cos
, cos
≤ 1.
(15)
q q
j̸=i
Putting (14) and (15) into (13), we obtain
Varνt (ψ) ≤ 2n.

(16)

2
Now set b := √ . With (6), (7), (12), and (16), Lemma 1 implies that for sufficiently large
5γ
n, d(t) ≥ 1 −
For any ε ∈ (0, 1), we can set γ =
we may conclude that

1
4
=
.
4 + b2
5γ + 1

ε−1 − 1
so that d(t) ≥ ε. With this choice of γ = γ(ε),
5

nq 2
tmix (ε) ≥ 2 2
4π σ




 −1
ε −1
log(n) + log
5

(17)

for n large enough.

4

1 × n Contingency Tables

4.1

Introduction

Recall from Section 1 the dynamics of the rate-1 continuous-time random walk over the
1 × n contingency tables with any fixed coordinate sum. We will only keep track of the first n − 1 coordinates so that the resulting process {Xt }t≥0 on Zn−1
will be aperiodic and q
irreducible. We will use (Z ∩ [−q/2, q/2))n−1 as representatives of Zn−1
. Let {Ys }s∈N be the q
embedded discrete-time chain of {Xt }. Since our processes are vertex-transitive, without loss
17

of generality, we may assume that X0 = 0 and Y0 = 0. Denote by P the transition matrix of {Ys }. Then the rate matrix Q for {Xt } is Q = P − I where I is the identity matrix, and the transition kernel Ht for {Xt } is Ht = exp(tQ). We write
Ht (x, y) := Px (Xt = y) = P(Xt = y|X0 = x).
Define
S := {aij ∈ Zn−1
|1 ≤ i, j ≤ n, i ̸= j}
q where aij ∈ Zn−1
is defined as follows: if 1 ≤ i, j ≤ n − 1, then aij has 1 on the i-th q
coordinate, −1 on the j-th coordinate, and zero elsewhere; if i = n, then aij has −1 on the j-th coordinate and zero elsewhere; if j = n, then aij has 1 on the i-th coordinate and with the symmetric set of zero elsewhere. Then {Ys } is a random walk on the group Zn−1
q generators S.
Let Y be a lattice-valued random vector with the same law as that of Y1 mapped to
2
(Z ∩ [−q/2, q/2))n−1 . Observe that Y has covariance matrix Γ, where n

1
i=j
.
Γij =
1
−
else n−1
Moreover, it is straightforward to check that

2(n − 1)




n
(Γ−1 )ij =



n − 1
n

i=j
.
else

Since P is aperiodic, irreducible, and symmetric, both {Xt } and {Ym } converge to the by the convergence theorem of Markov chains. We will give uniform distribution on Zn−1
q the mixing time of the continuous-time process {Xt }, and therefore the mixing time of the
1 × n contingency table random walk in continuous time.

4.2

Eigenvalues for {Xt } and {Ys }

By the same reasoning as in Section 3.2, the Fourier transforms {ρ̂y : y ∈ Zn−1
} =
q n−1
{ΦY (2πy/q) : y ∈ (Z ∩ [−q/2, q/2)) } are precisely the eigenvalues of the discrete-time process {Ys }, where ΦY is again the characteristic function of Y . And the eigenvalues
{λj : 1 ≤ j ≤ q n−1 } of the continuous-time process {Xt } are
{λj : 1 ≤ j ≤ q n−1 } = {exp(ΦY (2πy/q) − 1) : y ∈ (Z ∩ [−q/2, q/2))n−1 }.

18

By the continuous-time ℓ2 -bound, q n−1
2

4d(t) ≤

X

λ2t j

j=2



X

=



exp 2t ΦY

y∈(Z∩[−q/2,q/2))n−1



2πy q




−1

(18)

y̸=0

= −1 +

X



  
2πy
−1
.
exp 2t ΦY
q n−1

y∈(Z∩[−q/2,q/2))

One can find a proof of the discrete-time analog of the ℓ2 -bound in Lemma 12.18 of [18].

4.3

Strategy for Mixing Time Upper Bound

We will form the mixing time upper bound via the ℓ2 bound in (18). We use the link between the eigenvalues of Xt and the characteristic function of the lifted distribution to gain control on the rate of decay of the √
spectrum of Xt . This is done by Taylor expanding ΦY (2πy/q) in q
the regime where ∥y∥∞ ≤
. In this regime, the eigenvalues of Xt resemble the Gaussian
2π
characteristic function and depend on the quadratic form y T Γy, where Γ is the correlation matrix of the lifted distribution Y . In Section 4.5.2., we simplify the quadratic form y T Γy based on Schur complements of Γ−1 . We then use this in Section
4.5.3. to show the ℓ2 -bound
√
q
). In Section 4.5.1 we also show is sufficiently small in the high frequency regime (∥y∥∞ ≤
2π
√
q
2
all contributions to the ℓ -bound in the lower frequency regime (∥y∥∞ >
) by directly
2π
analyzing the decay of ΦY , the characteristic function of our lifted distribution, outside of a neighborhood of zero under the regime where q ≫ n.
√
q n−1
We will first show that the sum in (18) over y ∈ (Z ∩ [−q/2, q/2))
with ∥y∥∞ >
is
2π
small enough. For this
 to assume that q ≫ n. Specifically, we show that
 we need
√ to be true, q
1
2πy for y with ∥y∥∞ >
that ΦY
≤1−
. For the proof, we set the time to be
2π
q
25nq

−1
nq 2 log(αn)
3
t=
1−
(19)
8π 2
q where α ≥ 1 is independent of n. Then exp(2t(ΦY (2πy/q) −√1)) is of order n−O(q) at this q
time t, and therefore the sum in (18) over y with ∥y∥∞ >
vanishes as n → ∞ with
2π
q ≫ n.
√
q n−1
For y ∈ (Z ∩ [−q/2, q/2))
with ∥y∥∞ ≤
, we show that at time (19),
2π
  


2πy exp 2t ΦY
−1
≤ exp(−y T Γy log(αn)).
q
19

√
q
Finally, we will use this to conclude that at time t, the sum in (18) over y with ∥y∥∞ ≤
2π


4 + 4e2/e
√
is bounded above by exp for all n large enough.
α

4.4

Bounding Decay in Characteristic Function for Large θ

We will first show that the characteristic function ΦY (θ) has sufficient decay when ∥θ∥∞ >
1
2πy
. We can write the characteristic function as
√ , and then we will plug in θ =
q q
Pn−1
ΦY (θ) =

j=1 cos(θj ) +

Pn−1 Pj−1
cos(θj − θk )
j=2
 k=1
n

(20)

2

with θ = (θ1 , . . . , θn−1 )T .
4.4.1

1
π
Regime 1: √ < ∥θ∥∞ ≤
q
2

Recall that Y has covariance matrix expansion of ΦY (θ), we have

Pn−1 2 Pn−1 Pj−1
2
j=2
k=1 (θj − θk )
j=1 θj +

.
n

2θT Γθ
=
n
Since


cos(x) ≤ 1 −

2
Γ. By comparing the second-order terms in the Taylor n

(21)

2

1 π2
−
2 24



x2 ≤ 1 −

2 2
x
25

for |x| ≤ π,

(22)

it follows from (20) and (21) that
ΦY (θ) ≤ 1 −

4θT Γθ
25n

for ∥θ∥∞ ≤

π
.
2

(23)

1
Next, note that Γ has eigenvalue 1 +
with eigenspace span({1n−1 })⊥ and multiplicity n−1
1
n − 2, and eigenvalue with multiplicity 1 and eigenvector 1n−1 . Here, 1n−1 is the all-1
n−1
P
vector of dimension n − 1. Let θ := n−1
j=1 θj /(n − 1) be the average value of the coordinates
1
π
1
of θ. Suppose √ < ∥θ∥∞ ≤ . If |θ| ≤ √ , then q
2
2 q
θT Γθ ≥ ∥θ − θ1n−1 ∥22 ≥ (∥θ∥∞ − |θ|)2 ≥

20

1
.
4q

(24)

1
Otherwise if |θ| ≥ √ , then
2 q
θT Γθ ≥

1
1
∥θ̄1n−1 ∥22 ≥ .
n−1
4q

(25)

Combining (23),(24), and (25), we obtain
ΦY (θ) ≤ 1 −

1
25nq

(26)

1
π
whenever √ < ∥θ∥∞ ≤ .
q
2
4.4.2

Regime 2:

π
≤ ∥θ∥∞ ≤ π
2
π
. Also, we may assume that θ ∈
2

We will now show the decay of ΦY (θ) for ∥θ∥∞ ≥
[−π, π]n−1 . For θ = (θ1 , . . . , θn−1 ), let

A = A(θ) := {j = 1, . . . , n − 1 : |θj | ≥ 1}
with |A| being the cardinality of the set A.
n−1
. Since ∥θ∥∞ ≥ π/2, there exists j ∗ ∈ {1, . . . , n − 1} such that
Suppose |A| ≤
2
|θj ∗ | ≥ π/2. Then
!
P
j−1
n−1
n−1 X
X
1 X
j∈Ac 1 − cos(θj ∗ − θj )

ΦY (θ) = n
cos(θj ) +
cos(θj − θk ) ≤ 1 −
. (27)
n
2

j=1

2

j=2 k=1

Here, we replaced all the cosines in the first sum by 1, and in the second sum, we only considered the case where j = j ∗ and k ∈ Ac , and replaced all the other cosines by 1. Since
|θj ∗ | ≥ π/2 and θ ∈ [−π, π]n−1 ,
π/2 − 1 ≤ |θj ∗ − θj | ≤ π + 1
for every j ∈ Ac . So (27) is bounded above by n−1
(1 − cos(π/2 − 1))
1 − cos(π/2 − 1)

1− 2
=1−
.
n n
2

Now we suppose |A| ≥
1
ΦY (θ) = n
2

n−1
X
j=1

(28)

n−1
. Then
2

cos(θj ) +

j−1
n−1 X
X

!
cos(θj − θk )

j=2 k=1

21

P
≤1−

j∈A (1 − cos(θj ))

.
n
2

(29)

Here, we replaced all the cosines in the second sum by 1, and in the first sum, we only considered the case where j ∈ A and replaced all the other cosines by 1. Since θ ∈ [−π, π]n−1 ,
−1 ≤ cos(θj ) ≤ cos(1)
for every j ∈ A. So (29) is bounded above by
1−
4.4.3

n−1
(1 − cos(1))
1 − cos(1)
2

≤1−
.
n n
2

(30)

Combining the Two Regimes of Large θ

Now putting together (26), (28), and (30), we get
ΦY (θ) ≤ 1 −

1
25nq

(31)

√
q
1
whenever θ ∈ [−π, π]
and ∥θ∥∞ > √ . Thus for ∥y∥∞ >
, q
2π
  




2t
2πy exp 2t ΦY
−1
≤ exp −
.
q
25nq n−1

(32)

Plugging in the mixing time in (19) gives




exp 2t ΦY



2πy q




−1

q
≤ n 100π 2 .
−

(33)

√
√
q q
(n−1)
for ∥y∥∞ >
. Note that the number of y ∈ (Z∩[−q/2, q/2))
that satisfies ∥y∥∞ >
2π
2π
is
 √
n−1
q q n−1 − ⌊
⌋+1
= O(q n−1 ).
(34)
π
Thus if q ≫ n, then at the mixing time in (19), we conclude that
  


X
2πy exp 2t ΦY
−1
→0
q
√
q

(35)

∥y∥∞ > 2π

as n → ∞.

4.5

Bounding Decay in Characteristic Function for Small θ

By the previous section, we know that the contribution to the sum in (18) over y with
2πy
1
2πy
1
> √ will be small. We will therefore focus on the regime
≤ √ of y.
q ∞
q q ∞
q
22

4.5.1

Decomposing the Characteristic Function in ℓ2 -Bound

By the Taylor expansion of ΦY , we can write
"


T   

4 #

1 2πy
2
2πy
1
2πy
2πy
≤1−
Γ
+ E
·Y
.
ΦY
q
2
q n
q
24
q
√
q
2πy
1
Since
, it follows that
≤ √ , i.e. ∥y∥∞ ≤
q ∞
q
2π
n−1
2π∥y∥∞ X
2
2πy
·Y ≤
|Y (j)| ≤ √
q q
q j=1

(36)

(37)

as all the coordinates of Y = (Y (1), . . . , Y (n − 1))T are zero except at most two coordinates which take values ±1. Combining (36) and (37), we obtain
"


2 #
2πy
4π 2 T
1
2πy
ΦY
≤ 1 − 2 y Γy + E
·Y
q nq
6q q
(38)


4π 2 T
1
≤ 1 − 2 y Γy 1 −
.
nq
3q
Therefore, at the mixing time (19),
  


2πy exp 2t ΦY
−1
≤ exp(− log(αn)y T Γy).
q
4.5.2

Simplifying the Quadratic Form y T Γy

In this section, we will simplify the quadratic form y T Γy in the upper mixing time bound. For notations, we write Ik for the identity matrix and Jk for the all-1 matrix, both of dimension k × k. And recall that 1k is the all-1 column vector of dimension
 k. We will use the fact
B11 B12
written in block form, if that for any (n − 1) × (n − 1) symmetric matrix B =
B21 B22
B11 is invertible, then for x = (x1 , x2 )T ,
−1
−1
−1
−1
x1 ).
xT B −1 x = xT1 B11
x1 + (x2 − B21 B11
x1 )T (B22 − B21 B11
B12 )−1 (x2 − B21 B11

(39)

One may verify (39) using the decomposition




−1
Ik B11
B12
Ik
0
B11
0
.
B=
−1
−1
B12 0 In−k−1
B21 B11
In−k−1
0 B22 − B21 B11
n−1
(In−1 +Jn−1 ). Fix k = 2, . . . , n−1.
n
2
Let A(k) ∈ Rk be any principal submatrix of Γ−1 and write


n−1
A11 (k) A12 (k)
A(k) =
(Ik + Jk ) =
A21 (k) A22 (k)
n
Meanwhile, it is straightforward to check that Γ−1 =

23

with A22 (k) ∈ R. Since Γ is symmetric, so is Γ−1 , and thus A21 (k) = A12 (k)T . Note that
A11 (k) =

n−1
(Ik−1 + Jk−1 ), n

A12 (k) =

n−1
1k−1 , n

2(n − 1)
.
n

(40)

1 T
1 , k k−1

(41)

A22 (k) =

We then have
A11 (k)

−1

n
=
n−1



1
Ik−1 − Jk−1 , k

and
A21 (k)A11 (k)−1 A12 (k) =

A21 (k)A11 (k)−1 =
(n − 1)(k − 1)
.
nk

(42)

Thus by (40) and (42),
(n − 1)(k + 1)
.
(43)
nk
Now for any y = (y1 , . . . , yn−1 )T , we can inductively apply (39), (41), and (43) to obtain
!
n−1
X
n k
1
y T Γy =
(yk − µk (y))2
y2 +
(44)
n − 1 2 1 k=2
k+1
A22 (k) − A21 (k)A11 (k)−1 A12 (k) =

with
1
T
µk (y) := A21 (k)A−1
11 (k)(y1 , . . . , yk−1 ) =

k−1
X

k j=1

yj

for each k = 2, 3, . . . , n − 1.
Remark 1. As we will show in the next section, the actual value of µk (y) is not significant.
We will take all the µk (y)’s to be zero to obtain an upper bound which is enough for our conclusion.
4.5.3

Bounding the Sum for Small y
2πy
1
≤ √ is small at the mixing q ∞
q


1
2πy n−1
M := y ∈ (Z ∩ [−q/2, q/2))
:
≤√
q ∞
q

with
We now show that the sum in (18) over y ∈ Zn−1
q time. Let
√
and N := ⌊

q
⌋. For y = (y1 , . . . , yn−1 )T , we can write
2π
X
exp(− log(αn)y T Γy)
y∈M
n−1
1 2 X
k y1 +
(yk − µk (y))2
2
k+1
k=2
!!
n−1
1 2 X
k
(yk − µk (y))2
y1 +
.
2
k
+
1
k=2

n exp − log(αn)
=
n−1
y∈M
X

≤

X
y∈M

exp − log(αn)

24

!!

By regrouping the summation with the definition of M, the expression above can be written as
!

 X

N
N
X
1
2
···
exp − log(αn)j12
exp − log(αn)(j2 − µ2 (y))2
2
3
j1 =−N
j2 =−N


(45)
N
X
...
exp(− log(αn)(jn−1 − µn−1 (y))2 ) .
jn−1 =−N



k+1
2
log(αn)(jk − x) is maximized
Next, note that for all 2 ≤ k ≤ n − 1, j=−N exp −
k at x = 0 as long as n is large enough by Lemma 4. Thus we obtain
X
exp(− log(αn)y T Γy)
PN

y∈M



! n−1
!
N
Y
X
1
k exp − log(αn)j 2
≤
exp −
log(αn)j 2
2
k+1
j=−N
k=2 j=−N
!!


n−1
N
X
X
k
2
log(αn)j
.
= exp log exp −
k
+
1
j=−N
k=1
N
X

(46)

We look at the inner expression:
N
X





N
X
k k
2
2
exp −
log(αn)j = 1 + 2
log(αn)j exp −
k+1
k+1
j=1
j=−N
Bounding the sum by an integral, we have that the previous expression is at most




Z N
k k
2
log(αn) + 2
exp −
log(αn)z dz
1 + 2 exp −
k+1
k+1
1




Z ∞
k k
2
≤ 1 + 2 exp −
log(αn) + 2
z · exp −
log(αn)z dz k+1
k+1
1




k k+1
k log(αn) +
exp −
log(αn)
= 1 + 2 exp −
k+1
k log(αn)
k+1


k
≤ 1 + 4 exp −
log(αn) .
k+1

25

(47)

Using this bound, the fact that log(1 + x) ≤ x for all x > −1, and Lemma 2, we have
!


n−1
X
X
k
T
exp(− log(αn)y Γy) ≤ exp log(αn)
log 1 + 4 exp −
k+1
y∈M
k=1
!

n−1
X
k
(48)
log(αn)
≤ exp 4
exp −
k
+
1
k=1


4 + 4e2/e
√
≤ exp
.
α

4.6

Mixing Time Upper Bound

We are now ready to state the mixing time upper bound. For any ε ∈ (0, 1), set




4 + 4e2/e
2
√
α := inf z ≥ 1 : 4ε ≥ exp
−1 .
z

−1
nq 2 log(αn)
1
Putting together (18), (35), and (48), we conclude when t =
1−
,
8π 2
3q d(t) ≤ ε + o(1)
as n → ∞.

4.7

Mixing Time Lower Bound

We will now compute the mixing time lower bound for the 1 × n contingency table walk in continuous time. We showed that the discrete-time increment Y has equivariant coordinates,
2
with each coordinate having a marginal variance of . Recall that the off-diagonal entries of n
−1
the correlation matrix are equal to exactly
. Since at each jump the 1 × n contingency n−1
table random walk updates at most two coordinates, each by ±1, we have that ∥Y ∥1 ≤ 2.
2(n − 1)
And therefore, applying (17) and setting r = 2, σ 2 =
, we obtain n



 −1

nq 2
1
ε −1
tmix (ε) ≥ 2 log(n) + log 1 −
+ log
8π
n
5
p for all n large enough, assuming q ≫ log(n).
Now combining Sections 4.6 and 4.7 completes the proof of Theorem 1.

5

Generalization for Mixing Time Upper Bound

We now generalize our methods for mixing time upper bounds to a family of random walks on tori. The setup is similar to the one in Section 3 for the general mixing time lower bound.
26

5.1

Setup

Let {Ys }s∈N be an arbitrary irreducible symmetric discrete-time Markov chain on the torus
Znq with transition matrix P , such that
P(Ys = ys |Ys−1 = ys−1 ) = µ(ys − ys−1 ).
We set Y0 = 0. Let Y ∈ Zn be a lattice-valued random vector with law µ mapped to
(Z ∩ [−q/2, q/2))n . Suppose that ∥Y ∥1 ≤ r, ∥Y ∥∞ < q/2, and r2 ≪ q. Note that Y
σ2
for each coordinate. Write is symmetric. Assume Y has the same marginal variance n
2
σ
Var(Y ) =
Γ where Γ is the correlation matrix of Y . We will assume that Γ is positive n
definite. Let {Xt } be the rate-1 continuous-time Markov chain on Znq with rate matrix
Q = P − I and transition kernel Ht = exp(tQ).
By the same reasoning as in Section 3.2, the eigenvalues of {Xt } are

 



2πy n
exp ΦY
− 1 : y ∈ (Z ∩ [−q/2, q/2)) .
(49)
q
Recall from (18) that the ℓ2 -bound says


X

2

4d(t) ≤ −1 +



exp 2t ΦY

y∈(Z∩[−q/2,q/2))n



2πy q




−1

.

We will again upper bound the sum above in two regimes of y separately.
Our strategy in generalizing our upper bound method will be similar to that in Section 4
by analyzing the ℓ2 bound. We leverage the correspondence between eigenvalues of our walk
Xt and the characteristic function of its lift, ΦY . Similar to a local central limit theorem, we show our characteristic function ΦY is sufficiently decayed in the low frequency regime q
, and thus our ℓ2 sum is controlled. We define a Decay Condition that describes
∥y∥∞ ≥
2π
the amount of control needed, and√briefly discuss tools to show the condition holds. In q
, we Taylor expand ΦY and leverage the correlation the high frequency regime ∥y∥∞ ≤
2π
structure of the lifted distribution Y . We form a Correlation Condition based on Γ to gain control on the rate of decay of the spectrum of Xt .

√
5.2

Regime 1: ∥y∥∞ >

q
2π

When y is big, we need Y to satisfy the following decay condition, so that we can show that our ℓ2 sum is very small in the low frequency regime in a format analogous to a local central limit theorem. This condition is based on the rate of decay of the characteristic function ΦY
outside of a neighborhood of zero. For this condition to hold, we may need some additional asymptotic assumptions on q.
27

Definition 4. Let

L :=

2πy
1
y ∈ (Z ∩ [−q/2, q/2)) :
>√
q ∞
q n


.

and let Y be a symmetric random variable supported on (Z ∩ [−q/2, q/2))n with marginal
σ2
variance by coordinate. We say that Y satisfies the Decay Condition if for n
nq 2 log(n)
t≥
,
4π 2 σ 2
the sum



  
2πy
−1
→0
exp 2t ΦY
q y∈L

X

as n → ∞, where ΦY is the characteristic function of Y .
Note that we required q ≫ n in the 1 × n contingency table case for this condition to hold. We give a few other examples in which the Decay Condition is met.
Example 1 (Simple Random Walk on Znq ). If {Ys }s∈N is the simple random walk on Znq , then Y is the increment for the simple random walk on Zn . In other words, Y takes values in
{(0, . . . , 0, ±1, 0 . . . , 0) ∈ Zn },
1
. In this case, r = σ 2 = 1. The characteristic function ΦY (θ) =
each with probability
2n
1 Pn cos(θk ) for θ = (θ1 , . . . , θn ) ∈ Rn . Using the bound (22) for cosine, we have that for n k=1
nq 2 log(n)
y = (y1 , . . . , yn ) ∈ L and t ≥
,
4π 2
  



!!
n
1X
2πyk
2πy exp 2t ΦY
−1
≤ exp 2t −1 +
cos q
n k=1
q
!!
n
1X
2 4π 2 yk2
≤ exp 2t −1 +
1−
(50)
n k=1
25 q 2
!!
n
X
2 4π 2 yk2
= exp −2t
25n q 2
k=1
√
P
q
For y ∈ L, ∥y∥∞ >
. Since nk=1 yk2 ≥ ∥y∥2∞ , the expression above can be upper bounded
2π
by q
 

 

−
2 4π 2 ∥y∥2∞
2
≤ exp 2t −
≤ n 25π 2 .
exp 2t −
25n q2
25nq
Similar to (34), the cardinality of L is O(q n ). Thus the Decay Condition is met as long as q ≫ n.
28

Example 2 (Dirichlet Form Comparison). Methods of Dirichlet form comparison can be used to find regimes with the Decay Condition.
Definition 5. Let f be a real-valued function on a state space Ω, and let P be a reversible and irreducible transition matrix on Ω with stationary distribution π. We define the Dirichlet form,
E(f ) := ⟨(I − P )f, f ⟩π
where the inner product is defined as
⟨f, g⟩π :=

X

π(x)f (x)g(x)

x∈Ω

For any non-constant eigenfunction f of P with Varπ (f ) = 1, we have that
E(f ) = 1 − λ

(51)

with λ being the corresponding eigenvalue of f .
Using this, let P be the transition matrix for {Ys } and Pe be the transition matrix for the simple random walk on Znq . Suppose we have a constant B = B(n) so that for any real-valued function g on Ω we have e ≤ BE(g)

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
