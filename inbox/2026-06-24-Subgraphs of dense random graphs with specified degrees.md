---
source: "https://arxiv.org/abs/1002.3018v2"
title: "Subgraphs of dense random graphs with specified degrees"
author: "Brendan D McKay"
year: "2010"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1002.3018v2"
pdf: "https://arxiv.org/pdf/1002.3018v2"
captured_at: "2026-06-24T11:11:00Z"
updated_at: "2026-06-24T11:11:00Z"
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

# Subgraphs of dense random graphs with specified degrees

- 著者: Brendan D McKay
- 年: 2010
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1002.3018v2)
- ダウンロード: https://arxiv.org/pdf/1002.3018v2
- PDF: https://arxiv.org/pdf/1002.3018v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Let d = (d1, d2, ..., dn) be a vector of non-negative integers with even sum. We prove some basic facts about the structure of a random graph with degree sequence d, including the probability of a given subgraph or induced subgraph. Although there are many results of this kind, they are restricted to the sparse case with only a few exceptions. Our focus is instead on the case where the average degree is approximately a constant fraction of n. Our approach is the multidimensional saddle-point method. This extends the enumerative work of McKay and Wormald (1990) and is analogous to the theory developed for bipartite graphs by Greenhill and McKay (arXiv:math/0701600, 2009).

## PDF Text

arXiv:1002.3018v2 [math.CO] 27 Nov 2010

Subgraphs of dense random graphs with specified degrees
Brendan D. McKay∗
School of Computer Science
Australian National University
Canberra ACT 0200, Australia bdm@cs.anu.edu.au

Abstract
Let d = (d1 , d2 , . . . , dn ) be a vector of non-negative integers with even sum.
We prove some basic facts about the structure of a random graph with degree sequence d, including the probability of a given subgraph or induced subgraph.
Although there are many results of this kind, they are restricted to the sparse case with only a few exceptions. Our focus is instead on the case where the average degree is approximately a constant fraction of n.
Our approach is the multidimensional saddle-point method. This extends the enumerative work of McKay and Wormald (1990) and is analogous to the theory developed for bipartite graphs by Greenhill and McKay (2009).

1

Introduction

Let d = (d1 , d2 , . . . , dn ) be a vector of non-negative integers with even sum. Let X = (xjk )
be a symmetric n × n matrix over {0, 1} with zero diagonal. Define G(d, X) to be the number of n × n symmetric matrices A = (ajk ) over {0, 1} with zero diagonal, such that
(i) row j sums to dj , for 1 ≤ j ≤ n;
(ii) ajk = 0 whenever xjk = 1, for 1 ≤ j, k ≤ n.

Equivalently, G(d, X) is the number of labelled simple graphs with n vertices of degree d1 , d2 , . . . , dn , having no edges in common with the simple graph X. The special case
∗

Research supported by the Australian Research Council.

1

where X is the zero matrix 0 will also be denoted G(d). Define x = (x1 , x2 , . . . , xn ), where xj is the sum of the jth row of X.
One motive for interest in G(d, X) is that the ratio G(d, X)/G(d) is the probability that a random simple graph with degree sequence d has no edge in common with X.
Similarly, G(d − x, X)/G(d) is the probability that X appears as a subgraph. In these cases, and throughout the paper, probability spaces have the uniform distribution.
Define the matrix X = (x̄jk ) over {0, 1} with x̄jk = 1 iff j 6= k and xjk = 0. For
P
convenience we will adopt the convention that jk∈X means the sum over all {j, k} such
P
that xjk = 1, and similarly jk∈X means the sum over all {j, k} such that x̄jk = 1. Note that the equal sets {j, k} and {k, j} do not appear as separate terms in these sums.
Define the following key parameters.
E = 12

n
X

dj

(the number of edges)

j=1

2E
n d
λ=
n−1
d=

(the average degree)
(the density ignoring the diagonal)

A = 12 λ(1−λ)
X = 12

n
X

xj

(the number of edges of X)

j=1

δj = dj − d + λxj

(1 ≤ j ≤ n)

Direct asymptotic estimation of G(d, X) for nonzero X has been previously restricted to the sparse range. For representative results with bounded or very slowly growing degrees, see Bollobás and McKay [3] and Wormald [21]. For somewhat higher degrees we have the following. Let dmax = maxj dj , xmax = maxj xj and ∆ = dmax (dmax + xmax ).
Theorem 1 ([15]). Suppose dmax ≥ 1 and ∆ = o(E). Then, as n → ∞,
2
Pn
 Pn
(2E)!
j=1 dj (dj −1)
j=1 dj (dj −1)
G(d, X) =
exp −
−
Q
4E
16E 2
E! 2E nj=1 dj !
P

jk∈X dj dk
2
+ O(∆ /E) .
−
2E
The error term in Theorem 1 is o(1) only under the stronger condition that ∆2 = o(E), which implies that the graphs are quite sparse. The special case G(d) was determined by
McKay and Wormald [20] under the weaker condition d3max = o(E).
2

The probability G(d, X)/G(d) of being edge-disjoint from X, and the probability
G(d − x, X)/G(d) of containing X as a subgraph are easily deduced from Theorem 1.
They can also be found directly over a sometimes wider range of d values. Let (a)b denote the falling factorial. The following is a consequence of Theorems 2.9 and 2.10 of
McKay [14].
Theorem 2 ([14]). If ∆ + X = o(E) then, as n → ∞,
Qn

G(d − x, X)
j=1 (dj )xj exp O(∆X/E) .
= X
G(d)
2 (E)X
In the case of dense matrices, the first asymptotically precise enumeration result was that of McKay and Wormald [19], who proved Theorem 3 (below) in the case of X = 0
with a slightly weaker error term. This has been extended by Barvinok and Hartigan [1].
They identify a symmetric matrix (λjk ) over {0, 1} introduced in [19] (and used in greater generality in Section 3) as the matrix that maximises a certain entropy function. They then express the asymptotic value of G(d) as an effectively computable function of (λjk )
provided the values {λjk }j6=k are uniformly bounded away from 0 and 1. This forces the average degree to be Θ(n) but allows a much greater variation of degrees than we allow.
They also show that (λjk ) matches a typical graph with degree sequence d in a sense that we will describe in Section 2. This theme is explored in a somewhat different way by
Chatterjee, Diaconis and Sly [5].
Despite the absence of precise enumerative results for densities between o(n−1/2 ) and c/ log n, Krivelevich, Sudakov, Vu and Wormald [12] determined several almost-sure properties of random regular graphs over various ranges of density using a combination of switchings and analysis. Other such properties were determined by Boldi and Vigna [2], and Cooper, Frieze, Reed and Riordan [6, 7].
More recently, Krivelevich, Sudakov and Wormald [11] determined the probability of small induced subgraphs in random regular graphs of degree (n − 1)/2 under some conditions on the order and degree sequence of the subgraph.
The corresponding problems for bipartite graphs and digraphs were studied by Greenhill and McKay [8]; see that paper for a bibliography. The proof method in [8] is quite similar to that here.
We will also have need for the following additional parameters, for 1 ≤ j ≤ n and

3

ℓ, m ≥ 1.
Bj =

X

δk ,

R=

j=1

k|jk∈X
n
X
Rℓ =
δjℓ , j=1

Xℓ =

D=

H=

X

L=

jk∈X

K=

X

jk∈X

n
X

(dj − d)2 , xℓj ,

j=1

δj δk ,

jk∈X

X

n
X

X

xj xk ,

jk∈X
n
X
δjℓ xm
Cℓ,m =
j ,

(δj − xj )(δk − xk ),

j=1

(dj − d)(dk − d).

To calibrate and motivate our main enumeration result, we first develop a naı̈ve estimate of G(d, X) by extending an idea introduced in [19]. Generate a random graph by independently creating an edge jk with probability λ for each jk ∈ X. Each graph with n
E edges (none in common with X) appears with probability λE (1−λ)( 2 )−X−E . Moreover,
 d j
the event Ej that vertex j has degree dj has probability n−1−x
λ j (1 − λ)n−1−xj −dj for dj each j. If we (incorrectly) assume that the events E1 , . . . , En are independent, we obtain a guess for G(d, X) as follows:

n 
(n2 ) Y
n−1−x j
−X
λ
1−λ
b X) = (1 − λ)
.
(1)
G(d,
λ (1 − λ)
d j
j=1

In [19] it was proved that

G(d) =

√

b 0) exp
2 G(d,




R2
1
−
+ o(1) .
4 16A2 n4

under certain conditions on d. Our main result extends this to nonzero X.
Theorem 3. Let a, b > 0 be constants such that a + b < 21 . Then there is a constant
ε = ε(a, b) > 0 such that the following holds. Suppose that dj −d, xj = O(n1/2+ε ) uniformly for 1 ≤ j ≤ n, that X = O(n1+2ε ), and that min{d, n − d − 1} ≥

n
.
3a log n

for sufficiently large n. Then, as n → ∞,


√
R2
λX 2
D
1
−b b
−
+
−
+ O(n ) .
G(d, X) = 2 G(d, X) exp
4 16A2 n4 (1−λ)n2 2An2
4

Proof. The proof of this theorem is the main task of the paper. Here we will summarize the main phases and draw their conclusions together. The basic idea is to identify G(d, X)
as a coefficient in a multivariable generating function and to extract that coefficient using the saddle-point method. In Section 3, we write G(d, X) = P (d, X)I(d, X), where
P (d, X) is a rational expression and I(d, X) is an integral in n complex dimensions.
Both depend on the location of the saddle point, which is the solution of some nonlinear equations. Those equations are solved in Section 3.1, and this leads to the value of
P (d, X) in (26). In Sections 3.3–3.5, the integral I(d, X) is estimated in a superset of a small region R enclosing the origin (defined in (28)) and an equivalent small region R′
enclosing (π, . . . , π). The result is given by Lemma 12. Finally, in Section 3.6, we note that the integral restricted to the exterior of R ∪ R′ is negligible. The present theorem thus follows from (4), (26) and Lemmas 12 and 14.
The proof in various places requires ε to be sufficiently small, but there are only a finite number of such places so we can choose ε(a, b) to satisfy all of them at once. Note that the theorem remains true if ε(a, b) is decreased, since the conditions become stronger.
Throughout the paper, the asymptotic notation O(f (n)) refers to the passage of n e (n)), which is to be taken as a shorthand for any to ∞. We also use a modified notation O(f expression of the form O(f (n)ncε ) with c a numerical constant (perhaps a different constant at each occurrence). Under the assumptions of Theorem 3, we have λ−1 , (1−λ)−1 =
e c3 ).
O(log n). This implies, if c1 , c2 , c3 , c4 are constants, that λc1 (1 − λ)c2 nc3 +c4 ε = O(n

2

Subgraph probabilities

Define functions miss(d, X) and hit(d, X) as follows. The probability that a random simple graph with degrees d has no edges in common with X is
(1 − λ)X miss(d, X), and the probability that it includes X as a subgraph is
λX hit(d, X).
In this section, we apply Theorem 3 to estimate these probabilities. To avoid unnecessary messiness regarding the value of ε, we will suppose ε(a, b) in Theorem 3 is chosen to be small enough to satisfy the finite number of places in this section where a statement is only true if ε is small enough.
5

Theorem 4. Under the conditions of Theorem 3, we have

λX2
λ(1−2λ)X3
λX
λX 2
D
+
+
miss(d, X) = exp
+
2 2
2 −
(1−λ)n 2(1−λ)n
6(1−λ) n
(1−λ)n
λ(1−λ)n2

C2,1
(1−2λ)C1,2
C1,1
−b
−
+ O(n )
−
−
(1 − λ)n
2(1−λ)2 n2
2(1−λ)2 n2
and


(1−λ)X (1+λ)X2 (1+λ)(1+2λ)X3 (1−λ)X 2
hit(d, X) = exp
−
−
+
λn
2λn
6λ2 n2
λn2

C1,1 (1+2λ)C1,2
C2,1
L
−b
+
+
− 2 2 + O(n ) .
−
λn
λ(1−λ)n2
2λ2 n2
2λ n
Proof. Since (1 − λ)X miss(d, X) = G(d, X)/G(d), the first part can be obtained from
Theorem 3. The second part can be found in similar fashion, or by noting that the probability of a random graph avoiding X is the probability of the complement of the graph having X as a subgraph.
If X is not too dense, the probabilities in Theorem 4 asymptotically match those for an ordinary random graph with edge probability λ. Sufficient conditions are that miss(d, X) = 1 + o(1) if

λX2 + X maxj |dj − d| = o (1 − λ)n , and hit(d, X) = 1 + o(1) if


(1 − λ)X2 + X maxj |dj − d| = o λn .

Both these sufficient conditions hold, for example, if X = O(n1/2−2ε ), or if dj −d and xj are uniformly O(nε ) for 1 ≤ j ≤ n and X = O(n1−2ε ).

Since Theorem 4 is rather complex, we give some special cases to facilitate its application. We also give the value of num(d, X), which is the exponential factor in Theorem 3.

Corollary 5. Suppose the conditions of Theorem 3 hold, and in addition assume that d1 = · · · = dn = d. Then


1 λ(X 2 −H)
−b num(d, X) = exp
+ O(n ) ,
+
4
(1−λ)n2

λX
λX2
λ(2−λ)X3
miss(d, X) = exp
−
−
(1−λ)n 2(1−λ)n 6(1−λ)2 n2

λH
λX 2
−b
−
+ O(n ) ,
+
(1−λ)n2 (1−λ)n2
6

hit(d, X) = exp



(1−λ)X (1−λ)X2 (1−λ2 )X3
−
−
λn
2λn
6λ2 n2

(1−λ)X 2 (1−λ)H
−b
−
+ O(n ) .
+
λn2
λn2

Corollary 6. Suppose the conditions of Theorem 3 hold, and in addition assume that x1 = · · · = xn = x (which implies that x = O(n2ε )). Then


1
R2
λx2
K
−b num(d, X) = exp
−
+ O(n ) ,
+
−
4 4(1−λ) 2An2 16A2 n4


λx(x−2)
xR
K
−b miss(d, X) = exp −
−
−
+ O(n ) ,
4(1−λ)
2(1−λ)2 n2 2An2


(1−λ)x(x−2)
K
xR
−b hit(d, X) = exp −
+ O(n ) .
− 2 2−
4λ
2λ n
2An2
The two parts of Theorem 4 have a common generalization. Let Y be a supergraph of X. Then the probability that a random graph with degrees d has intersection with Y
equal to X is
G(d − x, Y )
.
G(d)
P
If the degrees of Y are y1 , . . . , yn , with yj = O(n1/2+ε ) uniformly over j, and nj=1 yj =
O(n1+2ε ), then this probability can be computed using two applications of Theorem 3.
The resulting general formula is rather complex, so we will be content with presenting the special case where Y consists of a single clique and otherwise isolated vertices. This is the important case of an induced subgraph.
Suppose that for some m, we have xm+1 = · · · = xn = 0. Let X [m] be the subgraph of
X induced by vertices 1, . . . , m (so X [m] has the same edges as X). For k, ℓ ≥ 0, define the quantity m
X
ωk,ℓ =
(dj − d)k (xj − λ(m − 1))ℓ .
j=1

Theorem 7. Assume the conditions of Theorem 3 and in addition that m = O(n1/2+ε )
and xm+1 = · · · = xn = 0. Then the probability that a random graph with degree sequence d has X [m] as an induced subgraph is m

λX (1 − λ)( 2 )−X

2
2
− 2ω1,0
2ω1,1 − ω0,2 m2 (1−2λ)ω0,1 4ω1,0ω0,1 − ω0,1
+
+
+
× exp
4An
2n
4An
8An2

(2ω1,1 − ω2,0 − ω0,2 )m (1−2λ)(ω0,3 + 3ω2,1 − 3ω1,2 )
−b
+
−
+ O(n ) .
4An2
24A2 n2
7

Note that, within the stated error term, the probability is independent of dm+1 , . . . , dn except inasmuch as they contribute to d and λ.
m
The factor λX (1 − λ)( 2 )−X in Theorem 7 is the probability for an ordinary ErdősRényi random graph with edge probability λ. A sufficient condition for the argument of the exponential to be o(1) is m2 (m + maxm j=1 |dj − d|) = o(An). Relaxing this condition by a factor of m allows us to see the leading terms of the deviation of behaviour from an ordinary random graph.
Corollary 8. Assume the conditions of Theorem 3 and also that m(m+maxm j=1 |dj −d|) =
o(An) and xm+1 = · · · = xn = 0. Then the probability that a random graph with degree sequence d has X [m] as an induced subgraph is


m
ω0,2
ω1,1
−X
X
)
(
2
−
+ o(1) .
exp
λ (1 − λ)
2An 4An
The term ω0,2 /(4An) was obtained in [11] in the case that d = ((n−1)/2, . . . , (n−1)/2),
√
m = o( n), provided x1 , . . . , xm don’t differ too much from λm. (In the regular case
ω1,1 = 0.)
It has been shown by Barvinok and Hartigan [1], see also [5], that an independent-edge model more accurately matching the d model has each edge jk chosen with probability λjk , where these constants were introduced in [19] and will appear generalised in the following section. We will call this the {λjk }-model. Under our strict constraints on d, we have
λjk = λ +

dj − d dk − d (1 − 2λ)(dj − d)(dk − d) e −3/2
+
+
+ O(n
).
n n
2An2

(2)

We can restate Theorem 7 with that model in mind.

Corollary 9. Assume the conditions of Theorem 3 and in addition that m = O(n1/2+ε )
and xm+1 = · · · = xn = 0. Then the probability that a random graph with degree sequence d has X [m] as an induced subgraph is
Y
Y
λjk
(1 − λjk )
jk∈X

jk ∈X
/


2
ω0,2
m2 (1−2λ)ω0,1 4ω1,0 ω0,1 − ω0,1
+
+
+
× exp −
4An
2n
4An
8An2

(2ω1,1 − ω0,2 )m (1−2λ)(ω0,3 − 3ω1,2 )
−b
−
+ O(n ) ,
+
4An2
24A2 n2

where the two products are restricted to 1 ≤ j < k ≤ m.
8

Note that the ω1,1 term of Corollary 8 has disappeared, but the ω0,2 term remains.
The exponential factor quantifies how much the {λjk }-model is in error. However, when
X is generated according to the {λjk }-model, the expectation of the argument of the e −1/2 ), so there is some sense in which we exponential is O(n−b ) and the variance is O(n can say that the {λjk }-model gives a very accurate estimate for typical subgraphs. The details of this remain to be worked out.
Let Y be any graph on n vertices and let Y be its number of edges. Define the random variable Ed,Y to be the number of edges that a random graph with degree sequence d has in common with Y . Barvinok and Hartigan [1] proved that Ed,Y is concentrated
P
close to jk∈Y λjk when Y = Ω(n2 ). In fact their estimate is explicit enough to infer this concentration with weaker bounds for Y = ω(n3/2 log n). Under our strict conditions on d, we can obtain such a result for all Y . We have not determined the best result that follows from Theorem 3, but will for this paper be content with the following weak corollary of Theorem 7.
Corollary 10. Assume the conditions of Theorem 3. If Y = O(n1/5 ), then
 

Y
Y
e −1/5 )
λk (1 − λ)( 2 )−k 1 + O(n
Prob(Ed,Y ) = k) =
k

(3)

uniformly over k = 0, . . . , Y . If Y = Ω(n1/5 ), then

Ed,Y = λY (1 + O(n−1/10+δ ))
1
).
with probability 1 − O(exp(n−δ )) for any δ ∈ (0, 10

Proof. If Y = O(n1/5 ), there are O(n1/5 ) vertices incident with edges of Y , so Theorem 7
shows that Z has a binomial distribution with the precision given by (3). If Y = Ω(n1/5 ), divide the edges of Y into subsets of size Θ(n1/5 ) and apply (3) and Chernoff’s Inequality to bound the number of edges in each subset that are in common with the random graph.
(Clearly the constants in this corollary can be tuned in various ways.)
The theorems above should be enough to allow transfer of quite a lot of the theory of ordinary random graphs to dense random graphs with given degrees. However, our purpose in this paper is to develop the tools rather than to explore the applications in detail. We will be content with some simple illustrations.
Theorem 11. Let d = (d, d, . . . , d) satisfy the conditions of Theorem 3. Then for a random d-regular graph, we have the following. (Note that in each case the quantity in front of the exponential is the expectation for ordinary random graphs with edge probability λ.)
9

(a) If n is even, the expected number of perfect matchings is


1−λ
λn/2 n!
−b exp
+ O(n ) .
4λ
2n/2 (n/2)!
(b) If q = q(n) is a integer function such that 3 ≤ q ≤ n, then the expected number of q-cycles is


(1−λ)q(n−q)
λq n!
−b
+ O(n ) .
exp −
2q(n−q)!
λn2
(c) The expected number of spanning trees is


7(1−λ)
n−2 n−1
−b n λ
exp
+ O(n ) .
2λ
Proof. Parts (a) and (b) follow immediately from Corollaries 6 and 5, respectively.
Part (c) is not so simple since trees have various degree sequences and those with maximum degree greater than n1/2+ε do not satisfy the requirements of Theorem 3. Let
T be the set of all labelled trees with n vertices. If x denotes the degree sequence of a member of T , let T1 be the subset of T with xmax ≤ nε and let T2 = T \ T1 .
For sufficiently small ε > 0, the following are true.

(i) The probability in T that the maximum degree exceeds k is at most 2n/k! for any integer k ≥ 0.
(ii) X = n − 1.

ε

(iii) In T1 we have that X2 = 5n + O(n1/2+3ε ) with probability 1 − O(e−n ).
e e
(iv) In T1 we have X3 = O(n)
and H = O(n).
(v) The sum of λ−

Pn

j=1 max{0,xj −n

ε

}

e nn−1 /(nε )! .
over T2 is O(1)

Facts (i) and (iv) follow from the well-known generating function for labelled trees by degree sequence, which is z1 z2 · · · zn (z1 + z2 + · · · + zn )n−2 .
To obtain (iii), note that the same probabilities occur if we take xj = 1 + Yj for
1 ≤ j ≤ n, where Y1 , . . . , Yn are independent Poisson variates with mean 1 truncated at nε , subject to having sum n−2. (This is a standard property of multinomial distributions; any mean will do.) Now we can write

P
2
1/2+3ε

P
Prob
Y
−
2n
≥
n j
P
P

Yj = n − 2 ≤
Prob
Yj2 − 2n ≥ n1/2+3ε
.
Prob
Yj = n − 2
10

The expectation of Yj2 is 2 + O(n−1). Now bound the numerator by applying a concenP
tration inequality like Hoefffing’s [10] and the denominator by noting that
Yj would be a Poisson distribution with mean n except for the truncation. Item (iii) follows.
P
To obtain (v), note that nj=1 max{0, xj − nε } ≥ ∆ − nε for trees with maximum degree ∆, and the number of such trees is bounded by (i) with k = ∆ − 1. Summing over
∆ > nε gives the desired bound.
Now we bound the expected number of trees in T2 that appear in a random d-regular graph. For a tree T ∈ T2 , let F (T ) be any forest obtained by deleting all but ⌈nε ⌉ edges from each vertex that has degree greater than nε . Then the probability that T appears is bounded by the probability that F (T ) appears. Moreover, F (T ) satisfies the requirements
P
of Corollary 5 and has at least n − 1 − nj=1 max{0, xj − nε } edges. Corollary 5 gives
ε

hit(d, F (T )) = eO(n ) . Applying fact (v), we find that the expected number of these trees easily falls within the error term of part (c) of the theorem.

Finally, the trees in T1 all satisfy the conditions of Corollary 5 and have hit(d, X) =

O(n3 ). Those with X2 = 5n + O(n1/2+3ε ) have hit(d, X) = exp 7(1−λ)/(2λ) + O(n−b) .
Part (c) of the theorem now follows from items (i) and (iii).
The average number of spanning trees in random regular graphs of bounded degree was studied in [13].

3

Proof of Theorem 3

In this section we express G(d, X) as a contour integral in n-dimensional complex space, then estimate its value using the saddle-point method.
We will use a shorthand notation for summation over doubly subscripted variables.
From the matrix X = (xjk ), define sets
X(j) = { k : 1 ≤ k ≤ n, xjk = 1 },

X(j) = { k : 1 ≤ k ≤ n, xjk = 0, k 6= j }

for 1 ≤ j ≤ n. Note that j ∈
/ X(j), X(j), also that |X(j)| = xj and |X(j)| = n − 1 − xj .
If zjk is a symmetric variable for 1 ≤ j, k ≤ n, we define zj∗ =

n
X

zjk ,

z∗∗ =

n X
n
X
j=1 k=1

k=1

11

zjk ,

zj∗|X =

X

zjk ,

z∗∗|X =

X

zjk ,

jk∈X

k∈X(j)

zj∗|X =

X

z∗∗|X =

zjk ,

k∈X(j)

X

zjk .

jk∈X

There is some slight lack of symmetry in the definitions. To clarify, we note that
X
X
X
zj∗ = z∗∗ .
zj∗|X = 2z∗∗|X and zj∗|X = 2z∗∗|X , but j

j

j

d

d

Firstly, notice that G(d, X) is the coefficient of z1 1 z2 2 · · · zndn in the function
Y
(1 + zj zk ).
jk∈X

By Cauchy’s theorem this equals
1
G(d, X) =
(2πi)n

I

···

I Q

jk∈X (1 + zj zk )
dz1 · · · dzn , z1d1 +1 · · · zndn +1

where each integral is along a simple closed contour enclosing the origin anticlockwise.
It will suffice to take each contour to be a circle; specifically, we will write zj = rj eiθj for 1 ≤ j ≤ n. Also define for 1 ≤ j, k ≤ n. Then
Q

jk∈X (1 + rj rk )
G(d, X) =
Q
d
(2π)n nj=1 rj j

λjk =

rj rk
1 + rj rk

Z π

Z π Q

···

−π

−π


i(θj +θk )
1
+
λ
(e
−
1)
jk jk∈X
P
dθ, exp(i nj=1 dj θj )

(4)

where θ = (θ1 , . . . , θn ). Write G(d, X) = P (d, X)I(d, X) where P (d, X) denotes the factor in front of the integral in (4) and I(d, X) denotes the integral. We will choose the radii rj so that there is no linear term in the logarithm of the integrand of I(d, X) when expanded for small θ. The linear term is
X

jk∈X

λjk (θj + θk ) −

n
X

dj θj .

j=1

For this to vanish for all θ, we require
λj∗|X = dj

(1 ≤ j ≤ n).
12

(5)

Although it is not hard to show that (5) has an exact solution, we can get by with a nearsolution since (4) is valid for all positive radii. In Section 3.1 we find such a near-solution and determine to sufficient accuracy the various functions of the radii, such as P (d, X), that we require. In Section 3.3 we evaluate the integral I(d, X) within a certain region R
defined in (28). Section 3.6 notes that the contribution to the integral from the region outside of R and its translate R + (π, . . . , π) is minor in comparison.

3.1

Locating the saddle-point

In this section we derive a near-solution of (5) and record some of the consequences. As with the whole paper, we work under the assumptions of Theorem 3.
Change variables to {aj }nj=1 as follows: rj = r where r=

1 + aj

r

From (6) we find that

,

1 − r 2 aj
λ
.
1−λ

λjk /λ = 1 + aj + ak + Zjk , where
Zjk =

(6)

aj ak (1 − r 2 − r 2 aj − r 2 ak )
,
1 + r 2 aj ak

(7)

(8)

and that equation (5) can be rewritten as
X
δj ak + Zj∗|X .
= (n − 1)aj − aj xj +
λ

(9)

k∈X(j)

Summing (9) over all j, we find that
X=

n
X
j=1

Replace the term

P


(n − 1)aj − aj xj + Z∗∗|X .

k∈X(j) ak in (9) by n
X

n

Pn

k=1 ak −

P

k∈X(j) ak −aj , and substitute the value

 X
1
1X
ak + ak xk +
− Z∗∗|X
ak =
n k=1
n n
k=1
13

(10)

implied by (10). After some rearrangement, we find that aj = Aj (a1 , . . . , an ) for each j, where n
2aj + aj xj
δj
X
1 X
+
− 2− 2
ak + ak xk )
Aj (a1 , . . . , an ) =
λn n
n n k=1
1 X
1
1
+
ak − Zj∗|X + 2 Z∗∗|X .
n n
n k∈X(j)

(11)


In the vicinity of a = (0, 0, . . . , 0), the iteration a := A1 (a), . . . , An (a) is a contraction mapping that converges to a solution of (5), as can be proved using the method demonstrated in [4]. However, as noted above, we do not need to solve (5) exactly but will work with an approximate solution. Hopefully without confusing the reader, from now on we will use a to denote the result of four iterations starting at a = (0, 0, . . . , 0). We will also write Zjk and λjk to mean the values implied by (7) and (8) for our chosen a.
Applying (11) four times, we find aj =

δj
δj xj
Bj
X
e −5/2 ),
+
+ · · · + O(n
2 − 2 +
λn λn n
λn2

(12)

e −3/2 ). Most of the terms involve where the ellipsis conceals about 60 terms of order O(n counts of subgraphs of X up to order 5, with the vertices weighted by powers of the numbers {δj }. This implies an expansion
Zjk =

δj δk (1 − 2λ)
e −5/2 ).
+ · · · + O(n
2λAn2

The value of λjk is given by substituting estimate (12) into (7). In particular, uniformly over j, e −3/2 ).
λj∗|X = dj + O(n
(13)

Define αjk , βjk , γjk by αjk = βjk = γjk = 0 if j = k and

1 λ (1 − λ ) = A + α , jk jk
2 jk
1 λ (1 − λ )(1 − 2λ ) = A + β , jk jk
3
jk
6 jk
1 λ (1 − λ )(1 − 6λ + 6λ2 ) = A + γ , jk jk jk
4
jk
24 jk

if j 6= k, where
1 λ(1 − λ)(1 − 6λ + 6λ2 ).
A = 21 λ(1 − λ), A3 = 16 λ(1 − λ)(1 − 2λ), and A4 = 24

14

(14)

In evaluating the integral I(d, X), the following approximations of αjk , βjk , and γjk for j 6= k will be required:
αjk =

(1 − 2λ)(δj + δk ) δj2 + δk2 (1 − 12A)δj δk (1 − 2λ)(Bj + Bk )
+
+
−
2n
2n2
4An2
2n2
λ(1 − 2λ)X (1 − 2λ)(δj xj + δk xk ) e −3/2
−
+
+ O(n
), n2
2n2
βjk =

(1 − 12A)(δj + δk ) e −1
+ O(n ),
6n e −1/2 ).
γjk = O(n

(15)

(16)
(17)

We will also need the following summations.
αj∗ = 21 (1 − 2λ)δj −

δj2
(1 − 2λ)(δj xj + Bj ) e −1/2
R
− 22 +
+ O(n
)
2n 2n
2n

R2
e 1/2 )
+ λ(1 − 2λ)X + O(n n
1
e
βj∗ = 6 (1 − 12A)δj + O(1)

α∗∗ = −

(19)
(20)

e
β∗∗ = O(n)

3.2

(18)

(21)

Estimating the factor P (d, X)

Let
Λ=

Y

jk∈X

λ

λjkjk (1 − λjk )1−λjk .

Then
Λ

−1

=

Y  1 + rj rk λjk

jk∈X

=

Y

jk∈X

=

Y

jk∈X

rj rk

(1 + rj rk )

(1 + rj rk )

1−λjk



n
Y
−λ
rj j∗|X
j=1

n
Y
e −3/2 )
−d +O(n
(1 + rj rk )
rj j j=1

using (13). Therefore the factor P (d, X) in front of the integral in (4) is given by

e −1/2 ) .
P (d, X) = (2π)−n Λ−1 exp O(n
15

(22)

We proceed to estimate Λ. Writing λjk = λ(1 + zjk ), we have log

 λjk

λjk (1 − λjk )1−λjk



λ
= λzjk log
1−λ



λλ (1 − λ)1−λ
λ
λ(1 − 3λ + 3λ2 ) 4
λ(1 − 2λ) 3
2
e −5/2 ).
+
z
+
zjk + O(n zjk
−
jk
2(1 − λ)
6(1 − λ)2
12(1 − λ)3

(23)

e −1/2 ), which implies that z e −1/2 ),
We know from (13) that λ∗∗|X = E + O(n
∗∗|X = X + O(n

e −1/2 )
hence the first term on the right side of (23) contributes λλX (1 − λ)−λX exp O(n

to Λ. Now using (7), and recalling that |X| = n2 − X, we can write zjk = aj + ak + Zjk and apply the estimates in the previous subsection to obtain
(n)
Λ = λλ (1 − λ)1−λ 2 (1 − λ)−X

(n + 2)R2 (1 − 2λ)R3 (1 − 6A)R4
× exp
−
+
4An2
24A2 n2
96A3 n3

C2,1 + 2D λ2 X 2
R22
−1/2
e
−
+
+ O(n
) .
+
4An2
2An2 16A2 n4

(24)

As in [19], our answer will be simpler when written in terms of binomial coefficients.
Using Stirling’s formula or otherwise we find that

n 
Y
n−xj −1
j=1

dj

= (2πn)−n/2 λ−n/2−λn(n−1) (1 − λ)−n/2−(1−λ)n(n−1)+2X

R2
(1 − 2λ)2 R2 (1 − 2λ)R3
1 − 14A
−
+
+
× exp −
24A
4An
16A2 n2
24A2 n2

C2,1
(1 − 6A)R4
λX
−1/2
e
−
−
+
+ O(n
) .
96A3 n3
4An2 (1 − λ)n

(25)

Combining (22), (24) and (25), we find that
P (d, X) =

−n/2

n 
Y
π
n−xj −1
j=1

λλ (1 − λ)1−λ

(n2 )

An dj

R22
R2
1 − 14A
−
−
× exp
24A
16A2 n2 16A2 n4

λX
λX 2
D
−b
−
+
−
+ O(n ) .
(1−λ)n (1−λ)n2 2An2

16

(26)

3.3

Estimating the main part of the integral

Our next task is to evaluate the main part of the integral I(d, X) given by

Z π
Z π Q
i(θj +θk )
− 1)
jk∈X 1 + λjk (e
P
I(d, X) =
···
dθ .
exp(i nj=1 dj θj )
−π
−π

(27)

It will be established in this section and the next that the value of the integral is concentrated near the places where the integrand has the largest absolute value. This happens at the two points θ = (0, 0, . . . , 0) and θ = (π, π, . . . , π). These two points are equivalent, since the integrand is unchanged under the mapping θ 7→ θ + (π, π, . . . , π).
P
(This requires the fact that dj is even; otherwise the mapping changes the sign of the integrand and the integral is zero as it should be.) Consequently, in this section we will focus on a neighbourhood of (0, 0, . . . , 0), specifically the hypercube R defined by

R = θ : |θj | ≤ n−1/2+ε , 1 ≤ j ≤ n .
(28)
Let F (θ) be the integrand of (27). We are going to establish the following.

Lemma 12. Under the conditions of Theorem 3, there is a region S with R ⊆ S ⊆ 4R
such that


n/2

Z
π
λX
R2
1 − 20A
−1/2
−b
F (θ) dθ = 2
+ O(n ) . (29)
+
+
exp −
An
24A
(1 − λ)n 16A2 n2
S
In a region O(1)R, we can expand
 X
X
(A + αjk )(θj + θk )2 − i
(A3 + βjk )(θj + θk )3
F (θ) = exp −
jk∈X

+


X

jk∈X

= exp −

X

jk∈X



X
−1/2
5
e n
.
(A4 + γjk )(θj + θk ) + O
+A
|θj + θk |
4

jk∈X

(A + αjk )(θj + θk )2 − i

1≤j<k≤n

+

X

4

A4 (θj + θk ) +

X

jk∈X

1≤j<k≤n

X

(A3 + βjk )(θj + θk )3

1≤j<k≤n


−1/2
e
A (θj + θk ) + O(n
) ,
2

where A, A3 , A4 , αjk , βjk , and γjk were defined in (14). Approximations for αjk , βjk , γjk e −1/2 ) uniformly over j, k.
were given in (15)–(17). Note that αjk , βjk , γjk = O(n

We will transform the integral to diagonalize the quadratic terms, proceeding in two
P
steps. The first step will diagonalize the quadratic form 1≤j<k≤n (θj + θk )2 , and the second will complete the diagonalization.
17

3.4

First change of variables

We first adopt from [19] a linear transformation that diagonalizes the quadratic form
P
2
1≤j<k≤n (θj + θk ) . Define c and y = (y1 , y2 , . . . , yn ) by s
n−2
c=1−
= 1 − 2−1/2 + O(n−1 )
(30)
2(n − 1)
n

cX
yk
θj = yj −
n k=1

(1 ≤ j ≤ n).

(31)

The transformation θ = T1 (y) defined by (31) has determinant 1 − c. Also
(1 + c)R ⊆ T1−1 R ⊆ (1 − c)−1 R.

P
For ℓ ≥ 1, define µℓ = nj=1 yjℓ . We find the following translations.
X
θj = (1 − c)µ1
X

j

(θj + θk )2 = (n − 2)µ2

1≤j<k≤n

X


(θj + θk )3 = (n − 4)µ3 + 3(1 − 2c) + 12c/n µ1 µ2


+ (−6c + 12c2 − 4c3 )/n − 4c2 (3 − c)/n2 µ31
X

(θj + θk )4 = (n − 8)µ4 + 3µ22 + 4(1 − 2c) + 32c/n µ1 µ3

1≤j<k≤n
− 24c(1 − c)/n + 48c2 /n2 µ21 µ2

+ 8c2 (1 − c)(3 − c)/n2 + 8c3 (4 − c)/n3 µ41
X
X

αjk (θj + θk )2 =
(1 − 4c/n)αj∗ + 2c2 α∗∗ /n2 yj2
1≤j<k≤n

j

1≤j<k≤n

+
X

X′


αjk − 4cαj∗/n + 2c2 α∗∗ /n2 yj yk

X′

(3 − 12c/n)βjk − 6c(1 − 4c/n)βk∗ /n

j,k

βjk (θj + θk )3 =

X
j

1≤j<k≤n

+


1 − 6c/n + 12c2 /n2 )βj∗ − 4c3 β∗∗ /n3 yj3


+ 12c2 βj∗ /n2 − 12c3 β∗∗ /n3 yj yk2
X′

−6cβjk /n + 12c2 βj∗ /n2 − 4c3 β∗∗ /n3 yj yk yℓ
+
j,k

j,k,ℓ

X

X

4c X
4c2 2
(θj + θk ) =
(yj + yk ) − µ1
xj yj + 2 µ1 X .
n n
j jk∈X
jk∈X
2

2

18

(32)

In the above, and following, a summation is over 1, 2, . . . , n for each index unless otherwise
P
specified. Moreover, a prime on the summation symbol (as ′ ) means that only terms where the summation indices have distinct values are included. For example,
X
X′
.
means
1≤j≤n,1≤k≤n j6=k

j,k

Using the size of the hypercube R together with the bounds (18)–(21), we find that whenever θ ∈ O(1)R, we have F (θ) = G(y), where
X

G(y) = −
(n − 2)A + αj∗ − Axj yj2
j

+

X′

−αjk + 2cαj∗ /n + 2cαk∗/n − 2c2 α∗∗ /n2


+ Axjk − 2Acxj /n − 2Acxk /n + 4Ac2 X/n2 yj yk
X

−i nA3 + βj∗ yj3
j,k

j

−i
−i

X′


A3 (3 − 6c) + 3βjk − 6cβk∗ /n yj yk2

j,k

X′


A3 (−6c + 12c2 − 4c3 )/n − 6cβjk /n + 12c2 βj∗ /n2 yj yk yℓ

j,k,ℓ

+ nA4

X′ 2 2
X 4
yj + 3A4
yj yk j

+ 4A4 (1 − 2c)

X′
j,k

j,k

yj yk3 − 24A4 c(1 − c)/n

X′

(33)

yj yk yℓ2

j,k,ℓ

X′
e −1/2 ).
+ 8c2 (1 − c)(3 − c)A4 /n2
yj yk yℓ ym + O(n j,k,ℓ,m

3.5

Completing the diagonalization

We now make a second change of variables, y = T2 (z), that diagonalizes the quadratic part of G(y), where z = (z1 , . . . , zn ). We will use the method from [8] that is a slight extension of [18, Lemma 3.2].
Lemma 13. Let U and Y be square matrices of the same order, such that U −1 exists and all the eigenvalues of U −1 Y are less than 1 in absolute value. Then
(I + Y U −1 )−1/2 (U + Y ) (I + U −1 Y )−1/2 = U , where the fractional powers are defined by the binomial expansion.
19

If we also have that both U and Y are symmetric, then (I + Y U −1 )−1/2 is the transpose of (I + U −1 Y )−1/2 , as proved in [4].
Let V = (vjk ) be the symmetric matrix such that the quadratic terms of (33) are
−yV y T . We have for all j 6= k that e 1/2 ), vjj = An + O(n e −1/2 ).
vjk = Axjk + O(n

Apply Lemma 13 with V = U +Y where U is the diagonal matrix with the same diagonal e −3/2 ). Therefore, since entries as V . The matrix U −1 Y has jk-entry equal to n−1 xjk + O(n e −1/2 ), the eigenvalues the ∞-norm (maximum row sum of absolute values) of U −1 Y is O(n e −1/2 ).
of U −1 Y are all O(n

Let T2 be the transformation given by T2 (y) = z, where z T = (I + U −1 Y )−1/2 y T .
e −1/2 ). Expanding (I + U −1 Y )−1/2 we find
By [4, Lemma 2], the Jacobian of T2 is 1 + O(n that for y ∈ O(1)R, yj = zj +

n
X
k=1


e −3/2 )zk + O(n e −1 )xjk zk ,
O(n

(34)

for each j, where the coefficients are uniform and independent of z. An expression of identical form writes z in terms of y. For y ∈ O(1)R, we find that G(y) = H(z) where
X
X


e 1/2 ) zj3
H(z) = −
(n − 2)A + αj∗ − Axj zj2 − i nA3 + O(n j

−i

X′
j,k

+ nA4
+

e
3A3 (1 − 2c) + O(n

−1/2

X′
j,k

X
j

j



) zj zk2 − i

X′
j,k,ℓ

e −1 )zj zk zℓ
O(n

X′


e −1 ) zj2 zk2
e −1 ) zj4 + 3A4
1 + O(n
1 + O(n

3
e
O(1)z j zk −

X′
j,k,ℓ

j,k

e −1 )zj zk zℓ2 +
O(n

X′

j,k,ℓ,m

(35)

e −2 )zj zk zℓ zm + O(n e −1/2 ),
O(n

e ) being a function of z.
with only the final expression of the form O(

Now define S = T1−1 (T2−1 (2R)). By (32) and (34), R ⊆ S ⊆ 4R. Consequently the conditions for our approximations are satisfied and (35) is valid for z ∈ 2R.
We can now apply Theorem 15 (see Appendix) to estimate the integral of H(z)
over 2R. We list the coefficients required.
e
D̂jkℓ = O(1)

Â = A
N =n
Jˆj = 0
20

e −1 )
Êj = A4 + O(n e −1 )
F̂jk = 3A4 + O(n

e −1/2 )
Ĝjk = O(n e −1/2 )
Ĥjkℓ = O(n

âj = (2A − αj∗ + Axj )n−1/2
e −1/2 )
B̂j = −iA3 + O(n
√
e −1/2 )
Ĉjk = −3(1 − 2)iA3 + O(n

e −1/2 )
Iˆjkℓm = O(n

e −1/2 ). Applying Theorem 15, we find that
We can take ∆ = 43 and have δ(z) = O(n
Z
 π n/2
 1 − 20A

λX
R2
e −1/2 )Ẑ ,
H(z) dz =
+
+
exp −
+
O(n
(36)
An
24A
(1 − λ)n 16A2 n2
2R
where

Ẑ = exp

 (1 − 2λ)2 
6A

.

e −1/2 )Ẑ = O(n−b ). Lemma 12 now
From the conditions of Theorem 3, we have that O(n
√
e −1/2 ) and follows on recalling that the Jacobian determinants of T1 and T2 are 2 + O(n e −1/2 ), respectively.
1 + O(n

3.6

Bounding the remainder of the integral

In the previous section, we estimated the value of the integral I(d, X) restricted to a small region S ⊇ R. As mentioned earlier, the integral over S + (π, . . . , π) is the same.
It remains to bound the integral over the remaining parts of [−π, π]n . Define Rc =

[−π, π]n \ R ∪ (R + (π, . . . , π) . By employing the same technique as in [8], but with the dissection of the region utilised in [19], we can establish the following. We will omit the proof since no new techniques are required, but we note for convenience that
Y
fjk (A + αjk , θj + θk )
|F (θ)| =
jk∈X

where fjk (q, z) =

p

1 qz 4 .
1 − 4q(1 − cos z) ≤ exp −qz 2 + 12

Lemma 14. Under the conditions of Theorem 3,
Z
Z
−1
|F (θ)| dθ = O(n ) F (θ) dθ.
c

R

S

21

Appendix: The value of an integral
In this appendix we give the value of a certain multi-dimensional integral. A similar integral appeared in [19] and variations of it appeared in [9, 16–18].
This appendix is notationally independent of the rest of the paper. Summations without explicit limits are over 1, 2, . . . , N for each of the summation indices. A prime
P
on the summation symbol (as ′ ) indicates that only terms with distinct values of the summation indices are included.
Theorem 15. Let ε′ , ε′′ , ε′′′ , ε̄, ε̌, ∆ be constants such that 0 < ε′ < ε′′ < ε′′′ , ε̌ > 0, ε̄ ≥ 0, and 0 < ∆ < 1. The following is true if ε′′′ and ε̄ are sufficiently small.
′

Let Â = Â(N) be a real-valued function such that Â(N) = Ω(N −ε ). For 1 ≤ j, k, ℓ, m, let âj , B̂j , Ĉjk , D̂jkℓ , Êj , F̂jk , Ĝjk , Ĥjkℓ , Iˆjkℓm, and Jˆj be complex-valued functions of N
such that âj , B̂j , . . . , Jˆj = O(N ε̄ ) uniformly over j, k, ℓ, m. Suppose that

X 2 X
X
X
X′
Ĉjk zj zk2
f (z) = exp −ÂN
zj +
Jˆj zj + N 1/2
âj zj2 + N
B̂j zj3 +
j

+N

X′
−1

j

D̂jkℓ zj zk zℓ + N

j

Êj zj4 +

j

j,k,ℓ

+N

X

X′
−1/2

Ĥjkℓ zj zk zℓ2 + N

X′

j

j,k

F̂jk zj2 zk2 + N 1/2

Iˆjkℓm zj zk zℓ zm + δ(z)

j,k,ℓ,m

j,k,ℓ

Ĝjk zj zk3

j,k

j,k

X′
−3/2

X′


is integrable for z = (z1 , z2 , . . . , zN ) ∈ UN and δ(N) = maxz∈UN |δ(z)| = o(1), where

UN = z ⊆ RN : |zj | ≤ N −1/2+ε̂ for 1 ≤ j ≤ N ,

where ε̂ = ε̂(N) satisfies ε′′ ≤ 2ε̂ ≤ ε′′′ . Then, provided the O( ) term in the following converges to zero,

N/2
Z

π
f (z) dz =
exp Θ1 + O N −1/2+ε̌ + (N −∆ + δ(N))Ẑ ,
ÂN
UN

where

Θ1 =

1
2ÂN 1/2
+

1

X

âj +

j

X′

1
4Â2 N

X 2
âj +
j

3

15

16Â3 N
X
Êj +

X

Ĉjk Ĉjℓ +
16Â3 N 3 j,k,ℓ
4Â2 N j
3 X
1
4 X ˆ2
ˆj +
Jj +
B̂
J
+
j
ÂN j
4Â2 N j
4Â2 N 2
22

B̂j2 +

3

8Â3 N 2
j
1 X′
F̂jk
4Â2 N 2 j,k
X′
Ĉj,k Jˆk j,k

X′
j,k

B̂j Ĉjk

Ẑ = exp



1
4Â2 N
+

X

Im(âj )2 +

j

1

3

3

X′

15
16Â3 N

X

Im(B̂j )2 +

j

Im(Ĉjk ) Im(Ĉjℓ ) +

1

3

X′

8Â3 N 2 j,k
X
Im(Jˆj )2

Im(B̂j ) Im(Ĉjk )

16Â N j,k,ℓ
4ÂN j

X′
3 X
1
+
Im(B̂j ) Im(Jˆj ) +
Im(B̂jk ) Im(Jˆk ) .
4Â2 N j
4Â2 N 2 j,k

Proof. The method of proof is the same as in [4], with extra terms added. To simplify the process, we did not explicitly compute the lower order terms which are presented as Θ2
in [4]. The details will be omitted.

References
[1] A. Barvinok and J. A. Hartigan, The number of graphs and a random graph with a given degree sequence, preprint (2010). Available at arxiv.org/abs/1003.0356.
[2] P. Boldi and S. Vigna, Lower bounds for sense of direction in regular graphs, Distrib.
Comput., 16 (2003) 279–286.
[3] B. Bollobás and B. D. McKay, The number of matchings in random regular graphs and bipartite graphs, J. Combin. Th. Ser. B, 41 (1986) 80–91.
[4] E. R. Canfield, C. Greenhill and B. D. McKay, Asymptotic enumeration of dense 0-1
matrices with specified line sums, J. Combin. Th. Ser. A, 115 (2008) 32–66.
[5] S. Chatterjee, P. Diaconis and A. Sly, Random graphs with a given degree sequence, preprint (2010). Available at arxiv.org/abs/1005.1136.
[6] C. Cooper, A. Frieze and B. Reed, Random regular graphs of non-constant degree: connectivity and hamiltonicity, Combin. Prob. Comput., 11 (2002) 249–261.
[7] C. Cooper, A. Frieze, B. Reed and O. Riordan, Random regular graphs of nonconstant degree: independence and chromatic number, Combin. Prob. Comput., 11
(2002) 323–341.
[8] C. Greenhill and B. D. McKay, Random dense bipartite graphs and directed graphs with specified degrees, Random Struct. Alg., 35 (2009) 222–249.

23

[9] C. Greenhill, B. D. McKay and X. Wang, Asymptotic enumeration of sparse irregular bipartite graphs, J. Combin. Th. Ser. A, 113 (2006) 291–324.
[10] W. Hoeffding, Probability inequalities for sums of bounded random variables, J.
Amer. Statist. Assoc., 58 (1963) 13–30.
[11] M. Krivelevich, B. Sudakov and N. C. Wormald, Regular induced subgraphs of a random graph, Random Struct. Alg., to appear.
[12] M. Krivelevich, B. Sudakov, V. Vu and N. C. Wormald, Random regular graphs of high degree, Europ. J. Combin., 18 (2001) 346–363.
[13] B. D. McKay, Spanning trees in random regular graphs, Third Caribbean Conference on Combinatorics and Computing, (University of West Indies, 1981) 139–143.
[14] B. D. McKay, Subgraphs of random graphs with specified degrees, Congr. Numer.,
33 (1981) 213–223.
[15] B. D. McKay, Asymptotics for symmetric 0-1 matrices with prescribed row sums, Ars
Combin., 19A (1985) 15–26.
[16] B. D. McKay, The asymptotic numbers of regular tournaments, eulerian digraphs and eulerian oriented graphs, Combinatorica, 10 (1990) 367–377.
[17] B. D. McKay and R. W. Robinson, Asymptotic enumeration of Eulerian circuits in the complete graph, Combin. Prob. Comput., 7 (1998) 437–449.
[18] B. D. McKay and X. Wang, Asymptotic enumeration of tournaments with a given score sequence, J. Combin. Theory Ser. A, 73 (1996) 77–90.
[19] B. D. McKay and N. C. Wormald, Asymptotic enumeration by degree sequence of graphs of high degree, European J. Combin., 11 (1990) 565–580.
[20] B. D. McKay and N. C. Wormald, Asymptotic enumeration by degree sequence of graphs with degrees o(n1/2 ), Combinatorica, 11 (1991) 369–382.
[21] N. C. Wormald, Models of random regular graphs, in Surveys in Combinatorics, 1999
(eds. J. D. Lamb and D. A. Preece), Cambridge University Press, 1999, 239–298.

24

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
