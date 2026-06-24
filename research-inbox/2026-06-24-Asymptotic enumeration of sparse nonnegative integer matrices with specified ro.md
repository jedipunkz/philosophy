---
source: "https://arxiv.org/abs/0707.0340v3"
title: "Asymptotic enumeration of sparse nonnegative integer matrices with specified row and column sums"
author: "Catherine Greenhill, Brendan D. McKay"
year: "2007"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/0707.0340v3"
pdf: "https://arxiv.org/pdf/0707.0340v3"
captured_at: "2026-06-24T11:11:06Z"
updated_at: "2026-06-24T11:11:06Z"
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

# Asymptotic enumeration of sparse nonnegative integer matrices with specified row and column sums

- 著者: Catherine Greenhill, Brendan D. McKay
- 年: 2007
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/0707.0340v3)
- ダウンロード: https://arxiv.org/pdf/0707.0340v3
- PDF: https://arxiv.org/pdf/0707.0340v3

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Let \svec = (s_1,...,s_m) and \tvec = (t_1,...,t_n) be vectors of nonnegative integer-valued functions of m,n with equal sum S = sum_{i=1}^m s_i = sum_{j=1}^n t_j. Let M(\svec,\tvec) be the number of m*n matrices with nonnegative integer entries such that the i-th row has row sum s_i and the j-th column has column sum t_j for all i,j. Such matrices occur in many different settings, an important example being the contingency tables (also called frequency tables) important in statistics. Define s=max_i s_i and t=max_j t_j. Previous work has established the asymptotic value of M(\svec,\tvec) as m,n\to\infty with s and t bounded (various authors independently, 1971-1974), and when \svec,\tvec are constant vectors with m/n,n/m,s/n >= c/log n for sufficiently large (Canfield and McKay, 2007). In this paper we extend the sparse range to the case st=o(S^(2/3)). The proof in part follows a previous asymptotic enumeration of 0-1 matrices under the same conditions (Greenhill, McKay and Wang, 2006). We also generalise the enumeration to matrices over any subset of the nonnegative integers that includes 0 and 1.

## PDF Text

arXiv:0707.0340v3 [math.CO] 16 Apr 2012

Asymptotic enumeration of sparse nonnegative integer matrices with specified row and column sums
Catherine Greenhill

Brendan D. McKay

School of Mathematics and Statistics

Research School of Computer Science

The University of New South Wales

Australian National University

Sydney NSW 2052, Australia

Canberra, ACT 0200, Australia

csg@unsw.edu.au

bdm@cs.anu.edu.au

Keywords: asymptotic enumeration, non-negative integer matrices, contingency tables, switchings
MSC 2000: 05A16, 05C50, 62H17

Abstract
Let s = (s1 , . . . , sm ) and t = (t1 , . . . , tn ) be vectors of nonnegative integer-valued
P
Pn functions of m, n with equal sum S = m i=1 si =
j=1 tj . Let M (s, t) be the number of m × n matrices with nonnegative integer entries such that the ith row has row sum si and the jth column has column sum tj for all i, j. Such matrices occur in many different settings, an important example being the contingency tables (also called frequency tables) important in statistics. Define s = maxi si and t = maxj tj .
Previous work has established the asymptotic value of M (s, t) as m, n → ∞ with s and t bounded (various authors independently, 1971–1974), and when all entries of s equal s, all entries of t equal t, and m/n, n/m, s/n ≥ c/ log n for sufficiently large c
(Canfield and McKay, 2007). In this paper we extend the sparse range to the case st = o(S 2/3 ). The proof in part follows a previous asymptotic enumeration of 0-1
matrices under the same conditions (Greenhill, McKay and Wang, 2006). We also generalise the enumeration to matrices over any subset of the nonnegative integers that includes 0 and 1.
Note added in proof, 2011: This paper appeared in Advances in Applied Mathematics 41 (2008), 459–481. Here we fix a small gap in the proof of Lemma 5.1
and make some other minor corrections. We emphasise that the statements of our results have not changed.

1

1

Introduction

Let s = s(m, n) = (s1 , . . . , sm ) and t = t(m, n) = (t1 , . . . , tn ) be vectors of nonnegative
Pn
P
integers with equal sum S = m j=1 tj . Let M(s, t) be the set of all m × n i=1 si =
matrices with nonnegative integer entries such that the ith row has row sum si and the jth column has column sum tj for each i, j. Then define M(s, t) = |M(s, t)| to be the number of such matrices.
Our task in this paper is to determine the asymptotic value of M(s, t) as m, n → ∞
under suitable conditions on s and t.
The matrices M(s, t) appear in many combinatorial contexts; see Stanley [16, Chapter 1] for a brief history. A large body of statistical literature is devoted to them under the name of contingency tables or frequency tables; see [6, 7] for a partial survey. In theoretical computer science there has been interest in efficient algorithms for the problem of generating contingency tables with prescribed margins at random, and for approximately counting these tables. See for example [1, 8, 14].
The history of the enumeration problem for nonnegative integer matrices is surveyed in [5], while a history for the corresponding problem for 0-1 matrices is given in [12]. Here we recall only the few previous exact results on asymptotic enumeration for nonnegative integer matrices. Define s = maxi si and t = maxj tj . The first non-trivial case s1 = · · · =
sm = t1 = · · · = tn = 3 was solved by Read [15] in 1958. During the period 1971–74, this was generalised to bounded s, t by three independent groups: Békéssy, Békéssy and
Komlós [2], Bender [3], and Everett and Stein [9], under slightly different conditions.
In the case of denser matrices, the only precise asymptotics were found by Canfield and McKay [5] in the case that the row sums are all the same and the column sums are all the same. Let M(m, s; n, t) = M((s, s, . . . , s), (t, t, . . . , t)), where the vectors have length m, n, respectively, and ms = nt.
Theorem 1.1 ([5, Theorem 1]). Let s = s(m, n), t = t(m, n) be positive integers satisfying ms = nt. Define λ = s/n = t/m. Let a, b > 0 be constants such that a + b < 21 . Suppose that m, n → ∞ in such a way that


5m
5n
(1 + 2λ)2
1+
≤ a log n.
(1.1)
+
4λ(1 + λ)
6n
6m
Define ∆(m, s; n, t) by
n
m 

m+t−1
n+s−1
t s


M(m, s; n, t) =
mn + λmn − 1
λmn
 m + 1 (m−1)/2  n + 1 (n−1)/2

∆(m, s; n, t) 
1
×
.
exp − 2 +
m n
m+n
2

Then ∆(m, s; n, t) = O(n−b)(m + n) as m, n → ∞.
Canfield and McKay conjectured that in fact 0 < ∆(m, s; n, t) < 2 for all s, t ≥ 1.
The results in the present paper establish that conjecture for sufficiently large m, n in the

case st = o (mn)1/5 . (See Corollary 4.2.)

The main result in this paper is the asymptotic value of M(s, t) for st = o(S 2/3 ). Our proof uses the method of switchings in a number of different ways. In several aspects our approach is parallel to that which provided our previous asymptotic estimate of N(s, t), the number of 0-1 matrices in the class M(s, t). We now restate that result for convenience. For any x, define [x]0 = 1 and for a positive integer k, [x]k = x(x−1) · · · (x−k+1).
Also define n
m
X
X
[si ]k ,
Tk =
[tj ]k
Sk =
j=1

i=1

for k ≥ 1. Note that S1 = T1 = S.
Theorem 1.2 ([12, Corollary 5.1]). Let s = s(m, n) = (s1 , . . . , sm ) and t = t(m, n) =
Pn
Pm
(t1 , . . . , tn ) be vectors of nonnegative integers with equal sum S =
j=1 tj .
i=1 si =
2/3
Suppose that m, n → ∞, S → ∞ and 1 ≤ st = o(S ). Then

S!
S2 T2 S2 T2 S3 T3 S2 T2 (S2 + T2 )
Qn
N(s, t) = Qm exp − 2 −
+
−
2S
2S 3
3S 3
4S 4
j=1 tj !
i=1 si !
 3 3 
S22 T3 + S3 T22 S22 T22
s t
−
.
+
+
O
2S 4
2S 5
S2
We now state our main result, which is the asymptotic value of M(s, t) for sufficiently sparse matrices. Note that the answer is obtained by multiplying the expression for
N(s, t) from Theorem 1.2 by a simple adjustment factor.
Theorem 1.3. Let s = s(m, n) = (s1 , . . . , sm ) and t = t(m, n) = (t1 , . . . , tn ) be vectors
Pn
P
of nonnegative integers with equal sum S = m j=1 tj . Suppose that m, n → ∞, i=1 si =
2/3
S → ∞ and 1 ≤ st = o(S ). Then

 3 3 
S2 T2 S2 T2
s t
M(s, t) = N(s, t) exp
+
+
O
S2
S3
S2

S!
S2 T2 S2 T2 S3 T3 S2 T2 (S2 + T2 )
Qn
= Qm
+
+
−
exp
2S 2
2S 3
3S 3
4S 4
j=1 tj !
i=1 si !
 3 3 
S22 T3 + S3 T22 S22 T22
st
.
−
+
+O
4
5
2S
2S
S2

3

Proof. The proof of this theorem is presented in Sections 2 and 3. First we show that the set of matrices in M(s, t) with an entry greater than 3 forms a vanishingly small proportion of M(s, t). We also show that it is very unusual for an element of M(s, t)
to have a “large” number of entries equal to 2 or a “large” number of entries equal to 3, where “largeness” is defined in Section 2. We establish these facts using switchings on the matrix entries. This allows us to concentrate on matrices in M(s, t) with no entries greater than 3 and not very many entries equal to 2 or 3.
We then proceed in Section 3 to compare the number of these matrices with the number
N(s, t) of {0, 1}-matrices with row sums s and column sums t. We do this by adapting the results from [12] used to prove Theorem 1.2. These calculations are carried out in the pairing model, which is described in Section 3. Our theorem follows on combining
Lemmas 3.1, 3.2 and Corollary 3.8.
In the semiregular case where si = s for 1 ≤ i ≤ m and tj = t for 1 ≤ j ≤ n,
Theorem 1.3 says the following.
Corollary 1.4. Suppose that m, n → ∞ and that sm = tn = S for nonnegative integer functions s = s(m, n), t = t(m, n) and S = S(m, n). If 1 ≤ st = o(S 2/3 ) then
M(m, s; n, t) =

 3 3 
(s − 1)(t − 1) (s − 1)(t − 1)(2st − s − t − 10)
st
S!
.
exp
−
+O
m n
(s!) (t!)
2
12S
S2
For some applications the statement of Theorem 1.3 is not very convenient. In Section 4 we will derive an alternative formulation, very similar to one given for N(s, t)
in [12]. For k = 2, 3, define m
X
mn
µ̂k =
(si − S/m)k
S(mn + S) i=1
n
X
mn
(tj − S/n)k .
ν̂k =
S(mn + S) j=1

To motivate the definitions, recall that S/m is the mean value of si and S/n is the mean value of tj , so these are scaled central moments. We will prove Corollary 4.1, stated in
Section 4, which has the following special case.
Corollary 1.5. Under the conditions of Theorem 1.3, if (1 + µ̂2 )(1 + ν̂2 ) = O(S 1/3 ) then

 n 
m 
Y
n+si −1 Y m+tj −1



tj si st j=1
i=1
1


exp 2 (1 − µ̂2 )(1 − ν̂2 ) + O
.
M(s, t) =
mn+S−1
S 2/3
S
4

Corollary 1.5 has an instructive interpretation. Following [5], we write M(s, t) =
MP1 P2 E, where




n 
m 
Y
Y
m+tj −1
n+si −1
mn+S−1
−1
−1
,
, P2 = M
, P1 = M
M=
t s
S
j i
j=1
i=1



st
E = exp 21 (1 − µ̂2 )(1 − ν̂2 ) + O
.
S 2/3
Clearly, M is the number of m × n nonnegative matrices whose entries sum to S. In the uniform probability space on these M matrices, P1 is the probability of the event that the row sums are given by s and P2 is the probability of the event that the column sums are given by t. The final quantity E is thus a correction to account for the non-independence of these two events.
Finally, in Section 5 we show how to generalise Theorem 1.3 to matrices whose entries are restricted to any subset of the natural numbers that includes 0 and 1.
A note on our usage of the O( ) notation in the following is in order. Given a fixed function f (S) = o(S 2/3 ), and any quantity φ that depends on any of our variables, O(φ)
denotes any quantity whose absolute value is bounded above by |cφ| for some constant c that depends on f and nothing else, provided that 1 ≤ st ≤ f (S).
Note added in proof, 2011: This version of the paper the same as the journal version [11], except as follows:
• Theorem 2.1, a statement of a special case of a more general result from [10], was previously incomplete. The first inequality in (2.1) need only hold if v is a sink, but this condition was absent in [11].
• A note has been added at the end of the proofs of Lemmas 3.5 and 3.7, clarifying why it is valid to apply [12, Lemma 4.6] and [12, Lemma 4.8] with a possibly larger value of N2 , N3 than used in [12].
• The proof of Lemma 5.1 has been changed to fix a small gap. The old proof did not guarantee that n1 (Q) = S − o(S) when Q ∈ M− \ M∗ . The definition of M− has changed and a new switching argument is given to correct this.
• We added a reference to the journal version [11] of this paper.
Note that none of the statements of our own results from [11] have changed.

2

Switchings on matrices

In this section we will show that the condition st = o(S 2/3 ) implies that most matrices have no entries greater than 3. We also find bounds on the number of entries equal to 2
5

or 3. Our tool will be the method of switchings, which we will analyse using results of
Fack and McKay [10] from which we will distill the following special case.
Theorem 2.1. Let G = (V, E) be a finite simple acyclic directed graph, with each v ∈ V
being associated with a finite set C(v), these sets being disjoint. Suppose that S is a multiset of ordered pairs such that for each (Q, R) ∈ S there is an edge vw ∈ E with
Q ∈ C(v) and R ∈ C(w). Further suppose that a, b : V → R are positive functions such that, for each v ∈ V ,
{(Q, R) ∈ S | Q ∈ C(v)} ≥ a(v) |C(v)| if v is not a sink,

(2.1)

{(Q, R) ∈ S | R ∈ C(v)} ≤ b(v) |C(v)| ,

where the left hand sides are multiset cardinalities. Let ∅ 6= Y ⊆ V . Then there is a directed path v1 , v2 , . . . , vk in G, where v1 ∈ Y and vk is a sink, such that
P
P
N(vi )
v∈Y |C(v)|
P
≤ P vi ∈Y
,
(2.2)
v∈V |C(v)|
1≤i≤k N(vi )

where N(vi ) is defined by

N(v1 ) = 1, a(v1 ) · · · a(vi−1 )
N(vi ) =
b(v2 ) · · · b(vi )

(2 ≤ i ≤ k).

Proof. This follows from Theorems 1 and 2 of [10].
For D ≥ 2, a D-switching is described by the sequence
Q; (i0 , j0 ), (i1 , j1 ), . . . , (iD , jD )



where Q is a matrix in M(s, t) and (i0 , j0 ), (i1 , j1 ), . . . , (iD , jD ) is a (D+1)-tuple of positions such that
• the rows i0 , . . . , iD are all distinct and the columns j0 , . . . , jD are all distinct;

• there is a D in position (i0 , j0 ) of Q;

• the entries in positions (iℓ , jℓ ) of Q are not equal to 0 or D + 1, for 1 ≤ ℓ ≤ D;
• there is a 0 in position (iℓ , j0 ) and position (i0 , jℓ ) of Q for 1 ≤ ℓ ≤ D.

This D-switching transforms Q into a matrix R ∈ M(s, t) by acting on the (D+1)×(D+1)
submatrix consisting of rows (i0 , . . . , iD ) and columns (j0 , . . . , jD ) as follows:




 D 0 0 ··· 0

 0 q1

0
q2
Q=

..
 ..
.
 .

0
qD





 7−→















6

0
1
1
···
1
1 q1 −1
1
q2 −1
..
..
.
.
1
qD −1





 = R.





Matrix entries not shown can have any values and are unchanged by the switching operation. Notice that the D-switching preserves all row and column sums and reduces the number of entries equal to D by at least 1 and at most D + 1. The number of entries greater than D is unchanged.
A reverse D-switching, which undoes a D-switching (and vice-versa), is described by

a sequence R; (i0 , j0 ), . . . , (iD , jD ) where R ∈ M(s, t) and (i0 , j0 ), (i1 , j1 ), . . . , (iD , jD )
is a (D+1)-tuple of positions such that
• the rows i0 , . . . , iD are all distinct and the columns j0 , . . . , jD are all distinct;
• there is a zero in position (i0 , j0 ) of R;
• the entries in positions (iℓ , jℓ ) of R are not equal to D, for 1 ≤ ℓ ≤ D;
• there is a 1 in position (iℓ , j0 ) and position (i0 , jℓ ) of R for 1 ≤ ℓ ≤ D.
Lemma 2.2. Let D ≥ 2 and let Q ∈ M(s, t) have at least K ≥ 2st non-zero entries that are not greater than D, and at least J entries equal to D. Then there are at least
J(K − 2st)D D-switchings and at most SD TD reverse D-switchings that apply to Q.
Proof. First consider D-switchings. We want a lower bound on the number of (D+1)tuples (i0 , j0 ), . . . , (iD , jD ) of indices where a D-switching may be performed. There are at least J ways to choose the position (i0 , j0 ). Then we can choose the remaining positions one at a time, avoiding choices which violate the rules. The choice of the last position
(iD , jD ) is the most restricted, so we bound that. By assumption, there are at least K
nonzero entries in Q that are not greater than D. Of these we must exclude the entry in position (i0 , j0 ) as well as entries in the same column as a nonzero entry in row i0 other than column j0 (at most (s − D)t positions), entries in the same row as a nonzero entry in column j0 other than row i0 (at most (t − D)s positions), and entries in row iℓ or column jℓ for 1 ≤ ℓ ≤ D − 1 (at most (D − 1)(s + t − 2) positions). Overall, we can choose position
(iD , jD ) in at least
K − 1 − (s − D)t − (t − D)s − (D − 1)(s + t − 2) ≥ K − 2st ways, and as we noted this also applies to each of the less restricted positions (iℓ , jℓ ), where 1 ≤ ℓ < D. Hence at most J(K − 2st)D D-switchings involve Q.
Next consider reverse D-switchings. An ordered sequence of D entries in the same row which equal 1 may be chosen in at most SD ways, and an ordered sequence of D entries in the same column which equal 1 may be chosen in at most TD ways. Some of these choices will not give a legal position for a reverse D-switching, but SD TD is certainly an upper bound.

7

Our first application of switchings will be to show that only a vanishing fraction of our matrices have any entries greater than 3. For j ≥ 0 and D ≥ 2, let MD (j) be the set of all matrices in M(s, t) with exactly j entries equal to D and none greater than D.
S
Define MD (>0) = j>0 MD (j), and note that MD+1 (0) = MD (0) ∪ MD (>0).

Lemma 2.3. Suppose that 1 ≤ st = o(S 2/3 ). Let U1 = U1 (s, t) be the set of all matrices in M(s, t) which contain an entry greater than 3. Then
|U1 |/M(s, t) = O(s3 t3 /S 2 ).
Proof. The largest possible entry of a matrix in M(s, t) is ∆ = min{s, t}. We will apply
Theorem 2.1 to successively bound the possibility that the maximum entry is D, for
D = ∆, ∆ − 1, . . . , 4.
Fix D with 4 ≤ D ≤ ∆. Define a directed graph G = (V, E) with vertex set V =
{v0 , v1 , v2 , . . . } and edge set E = {vj vi | j − D − 1 ≤ i ≤ j − 1}. Associate each vi with the set C(vi ) = MD (i). Define S to be the set of pairs (Q, R) related by a D-switching, where
Q ∈ vj , R ∈ vi for some vj vi ∈ E. Define Y = {v1 , v2 , . . . } ⊆ V . Note that SD TD > 0
since D ≤ ∆.
We can now use Theorem 2.1 to bound
P
|C(v)|
|MD (>0)|
= Pv∈Y
,
|MD+1(0)|
v∈V |C(v)|

once we have found positive functions a, b : V → R satisfying (2.1). These are provided by Lemma 2.2 with J = j and K = S/D, the latter being clear since there are no entries greater than D and the total of all the entries is S. We have S/D > 2st since
D ≤ ∆ ≤ (st)1/2 . Thus we can take a(vj ) = j(S/D − 2st)D and b(vj ) = SD TD .
Theorem 2.1 tells us that, unless MD (>0) = ∅, there is a directed path vt1 , vt2 , . . . , vtq , where q > 1 and t1 > t2 > · · · > tq = 0 (since v0 is the only sink) such that (2.2) holds.
Hence, using the values of N as given in Theorem 2.1 we have
N(vtq−1 ) + · · · + N(vt1 )
|MD (>0)|
≤
|MD+1(0)|
N(vtq ) + · · · + N(vt2 )
N(vti−1 )
≤ max
1≤i≤q N(vti )
b(MD (ti ))
= max
1≤i≤q a(MD (ti−1 ))
SD TD
≤
.
(S/D − 2st)D

Let ξD denote this upper bound: that is, ξD = SD TD /(S/D − 2st)D for 4 ≤ D ≤ ∆. Note
8

that ξ4 = O(s3t3 /S 2 ). For 4 ≤ D < ∆, we have ξD > 0 and


ξD+1
(D + 1)D+1
(S − 2stD)D
≤ st
ξD
DD
(S − 2st(D + 1))D+1

−D
2st
Dst
1−
= O(1)
S − 2st(D + 1)
S − 2stD
= o(1)
uniformly over D, where the last step uses the observation that ∆ ≤ (st)1/2 = o(S 1/3 ).
Since U1 = M4 (>0)∪M5 (>0)∪· · ·∪M∆ (>0) and MD+1 (0) ⊆ M(s, t) for 4 ≤ D ≤ ∆, we have |U1 |/M(s, t) ≤ ξ4 + ξ5 + · · · + ξ∆ = O(s3 t3 /S 2 ) as required.
We may therefore restrict our attention to matrices with no entry greater than 3.
Next we find upper bounds on the numbers of entries equal to 2 or 3 which hold with high probability.
Define

22
if S2 T2 < S 7/4 ,



1
N2 = ⌈log S⌉
if S 7/4 ≤ S2 T2 < 5600
S 2 log S,



1
S 2 log S ≤ S2 T2 ;
⌈5600S2 T2 /S 2 ⌉ if 5600

N3 = max ⌈log S⌉, ⌈230000S3T3 /S 3 ⌉ .
(Here and throughout the paper we have not attempted to optimise constants.)
We will use the following lemma.

Lemma 2.4. Let k be a positive integer and let q and n be positive real numbers such that n ≥ kq. Then n(n − q) · · · (n − (k − 1)q) ≥ (n/e)k .
Proof. Dividing the left side by nk gives, for n > kq, k−1
Y
i=0

(1 − iq/n) = exp

X
k−1

log(1 − iq/n)

Z k


log(1 − xq/n) dx

≥ exp

i=0

0



= exp −k − (n/q − k) log(1 − kq/n))

≥ exp(−k).

The second line holds because log(1 − xq/n) is a decreasing function for x ∈ [0, k]. The case n = kq follows by continuity.
9

Lemma 2.5. Let 1 ≤ st = o(S 2/3 ). Then, with probability 1 − O(s3t3 /S 2 ), a random element of M(s, t) has no entry greater than 3, at most N3 entries equal to 3, and at most N2 entries equal to 2.
Proof. In view of Lemma 2.3, we may restrict our attention to the set M4 (0) of all matrices in M(s, t) with maximum entry at most 3. We will start by applying 3-switchings as in
Lemma 2.3 but the analysis will be more delicate.
In applying Theorem 2.1 we have V = {v0 , v1 , . . . }, with vh associated with M3 (h), and Y = {vh | h > N3 }. For sufficiently large S, we have from Lemma 2.2 that we can
1
take a(vh ) = 28
hS 3 and b(vh ) = S3 T3 . If S3 T3 = 0 then entries equal to 3 are impossible, so we assume that S3 T3 > 0. Define ϕ = 28S3 T3 /S 3 .
According to Theorem 2.1, there is a sequence h1 > h2 > · · · > hq = 0, with h1 > N3 and hi−1 − 4 ≤ hi < hi−1 for all i, such that
|M3 (>N3 )|
|M3(>N3 )|
≤
|M(s, t)|
|M4 (0)|
N(hℓ ) + N(hℓ−1 ) + · · · + N(h1 )
,
≤
N(hq ) + N(hq−1 ) + · · · + N(h1 )
where ℓ is the largest index such that hℓ ≥ N3 + 1 and N(hi ) = h1 · · · hi−1 ϕ−i+1 for all i.
Define u = ⌊ 41 log S⌋. Since N3 ≥ ⌈log S⌉, we have ℓ + u ≤ q. Also, for 0 ≤ i ≤ ℓ − 1,
N(hℓ )
N(hℓ−i )
≤
.
N(hℓ+u−i )
N(hℓ+u )
Therefore,
|M3 (>N3 )|
N(hℓ ) + N(hℓ−1 ) + · · · + N(h1 )
≤
|M(s, t)|
N(hℓ+u ) + N(hℓ+u−1 ) + · · · N(hu+1 )
N(hℓ )
≤
N(hℓ+u )
ϕu
=
hℓ hℓ+1 · · · hℓ+u−1
ϕu
≤
.
(N3 + 1)(N3 − 3) · · · (N3 − 4u + 5)
Since N3 + 1 > 4u we can apply Lemma 2.4 to obtain the bound
u

|M3(>N3 )|
ϕe
≤
.
|M(s, t)|
N3 + 1
10

Now N3 ≥ 230000S3T3 /S 3 ≥ 8214 ϕ, and u ≥ 14 log S − 1, so this upper bound is at most
 e 1 log S−1
1
4
= O(1)S 4 log(e/8214) = O(S −2 ).
8214
This shows that with probability O(s3 t3 /S 2 ) there are at most N3 entries equal to 3, as required.
To bound the number of entries equal to 2, we proceed in the same manner using
2-switchings, working under the assumption that there are at most N3 entries equal to 3
and none greater than 3. In applying Lemma 2.2, we can take K = 21 (S − 3N3 ), so that
(K − 2st)2 ≥ 51 S 2 for sufficiently large S. Define ψ = 5S2 T2 /S 2 . Arguing as above we find a sequence d1 > d2 > · · · > dr = 0, with the following properties: (i) d1 > N2 and di−1 − 3 ≤ di < di−1 for all i, and (ii) if p is the greatest integer such that dp > N2 then, for any w with 0 < w ≤ r − p, the probability that there are more than N2 entries equal to 2, subject to there being at most N3 equal to 3, is bounded above by
ψw
.
(2.3)
dp dp+1 · · · dp+w−1
First suppose that S2 T2 < S 7/4 , so that N2 = 22 and ψ < 5S −1/4 . Since dp ≥ N2 + 1 =
23, it follows that r − p ≥ 8. Taking w = 8 in (2.3) gives
|Y |
ψ8
≤
= O(S −2).
|M(s, t)|
dp dp+1 · · · dp+7
Now suppose that S2 T2 ≥ S 7/4 . Then N2 ≥ ⌈log S⌉ so we can take w = ⌊ 31 log S⌋.
Arguing as above by applying Lemma 2.4 to (2.3), we obtain the bound O(S −2 ) again.
This completes the proof.
From now on we proceed in two cases, as in [12]. Say that the pair (S2 , T2 ) is substantial if the following conditions hold:
• 1 ≤ st = o(S 2/3 ),

• S2 ≥ s log2 S and T2 ≥ t log2 S,

• S2 T2 ≥ (st)3/2 S.

Lemma 2.6. If 1 ≤ st = o(S 2/3 ) and (S2 , T2 ) is insubstantial, then with probability
1 − O(s3 t3 /S 2 ), a random element of M(s, t) has no entry greater than 3, at most one entry equal to 3 and at most two entries equal to 2.

11

Proof. The absence of entries greater than 3 follows from Lemma 2.3. We can also, by
Lemma 2.5, assume that the number of entries equal to 2 or 3 is o(S). Therefore, most of the matrix entries are 0 or 1. Let N be the set of all matrices in M(s, t) with no entries greater than 3, at most N2 entries equal to 2 and at most N3 entries equal to 3.
To bound the number of entries equal to 2 or 3 even more tightly, as this lemma requires, we employ D-switchings (D = 2, 3) with the additional restriction that q1 =
· · · = qD = 1. This ensures that these restricted D-switchings reduce the number of entries equal to D by exactly one and do not create any new entries equal to 2 or 3.
Let N ′′ (h) be the number of matrices in N with h entries equal to 3. If Q is such a matrix then the number of restricted 3-switchings applicable to Q is hS 3 (1 + o(1)) and the number of reverse restricted 3-switchings is at most S3 T3 . (This follows using arguments similar to those in Lemma 2.2, since there are S − o(S) entries equal to 1.) Therefore, if the denominator is nonzero,
N ′′ (h)
S3T3
= O(1) 3 .
(2.4)
′′
N (h−1)
hS
We can now easily check that each of the three causes of insubstantiality (namely, S2 <
s log2 S, T2 < t log2 S, and S2 T2 < (st)3/2 S) imply that
S3 T3
= O(s3/2 t3/2 /S) = o(1).
S3
Hence (2.4) implies that
P

′′
h≥2 N (h)
= O(s3 t3 /S 2 ).
N ′′ (0)

In precisely the same way, using restricted 2-switchings, we find that
P
′
d≥3 N (d)
= O(s3t3 /S 2 ),
′
N (0)
where N ′ (d) is the number of matrices in N with d entries equal to 2 and at most one entry equal to 3. The lemma follows.

3

From pairings to matrices

The remainder of the paper will involve calculations in the pairing model, which we now describe. (This model is standard for working with random bipartite graphs of fixed degrees: see for example [13].) Consider a set of S points arranged in cells x1 , x2 , . . . , xm , where cell xi has size si for 1 ≤ i ≤ m, and another set of S points arranged in cells y1 , y2 , . . . , yn where cell yj has size tj for 1 ≤ j ≤ n. Take a partition P (called a pairing)
of the 2S points into S pairs with each pair having the form (x, y) where x ∈ xi and y ∈ yj for some i, j. The set of all such pairings, of which there are S!, will be denoted by P(s, t). We work in the uniform probability space on P(s, t).
12

Two pairs are called parallel if they involve the same two cells. A parallel class is a maximal set of mutually parallel pairs. The multiplicity of a parallel class (and of the pairs in the class) is the cardinality of the class. As important special cases, a simple pair is a parallel class of multiplicity one, a double pair is a parallel class of multiplicity two, while a triple pair is a parallel class of multiplicity three.
Each pairing P ∈ P(s, t) gives rise to a matrix in M(s, t) by letting the (i, j)-th entry of the matrix equal the multiplicity of the parallel class from xi to yj in P .
In [12] we noted that the number of pairings which gives rise to each 0-1 matrix in
M(s, t) depends only on s and t and is independent of the structure of the matrix. Hence the task of counting such matrices reduces to finding the fraction of pairings that have no multiplicities greater than 1.
More generally, matrices in M(s, t) correspond to different numbers of pairings. For a pairing P ∈ P(s, t), define the multiplicity vector of P to be a(P ) = (a2 , a3 , . . . ) where ar is the number of parallel classes of multiplicity r. Also define the weight of P as w(P ) = (2!)a2 (3!)a3 (4!)a4 · · ·
For Q ∈ M(s, t), define w(Q) and a(Q) to be the common weight and multiplicity vectors of the pairings that yield Q.
By elementary counting, a matrix Q ∈ M(s, t) corresponds to exactly m
n
Y
1 Y
si !
tj !
w(Q) i=1
j=1

pairings in P(s, t). Therefore, if A is a set of multiplicity vectors, PA = {P ∈ P(s, t) |
a(P ) ∈ A}, and MA = {Q ∈ M(s, t) | a(Q) ∈ A}, then
P
P ∈PA w(P )
Qn
|MA | = Qm
.
(3.1)
j=1 tj !
i=1 si !
This holds in particular if A is the set of all nonnegative integer sequences, in which case
PA = P(s, t) and MA = M(s, t).
We first prove Theorem 1.3 in the case that (S2 , T2 ) is insubstantial.
Lemma 3.1. If 1 ≤ st = o(S 2/3 ) and (S2 , T2 ) is insubstantial then Theorem 1.3 holds.
Proof. Similarly to [12, Lemma 2.2], define a doublet to be to be an unordered set of 2
parallel pairs. A double pair provides one doublet, while a triple pair provides 3 doublets.
For the uniform probability space over P(s, t), let br be the expectation of the number of

13

sets of r doublets, for r ≥ 0. In [12, Lemma 2.2] it is shown that b0 = 1,
S2 T2
, b1 =
2[S]2
S3 T3
(S 2 − 4S3 − 2S2 )(T22 − 4T3 − 2T2 )
b2 =
+ 2
,
2[S]3
8[S]4
S3 T3
+ O(s3 t3 /S 2 ), b3 =
6[S]3
b4 = O(s3t3 /S 2 ).
Let pk denote the probability that a randomly chosen pairing contains exactly k doublets, for k ≥ 0. Then
 
X
r+k r br pk =
(−1)
k r≥k and the partial sums of this series alternate above and below pk (see for example [4,
Theorem 1.10]). Applying this, we find that
S2 T2
S3 T3
(S2 2 − 4S3 − 2S2 )(T2 2 − 4T3 − 2T2 )
+
+
+ O(s3t3 /S 2 ),
2[S]2 3[S]3
8[S]4
S3 T3
(S 2 − 4S3 − 2S2 )(T22 − 4T3 − 2T2 )
S2 T2
−
− 2
+ O(s3t3 /S 2 ), p1 =
2[S]2 2[S]3
4[S]4
(S22 − 4S3 − 2S2 )(T22 − 4T3 − 2T2 )
p2 =
+ O(s3 t3 /S 2 ),
8[S]4
S3 T3
+ O(s3t3 /S 2 ).
p3 =
6[S]3

p0 = 1 −

(The expression for p0 was also derived in [12, Lemma 2.2].) The configurations defining these cases are, respectively, no parallel pairs, one double pair, two double pairs, and one triple pair.
Applying Lemma 2.6 and (3.1),

S!
Qn p0 + 2p1 + 4p2 + 6p3
j=1 tj !
i=1 si !

p0 + 2p1 + 4p2 + 6p3 + O(s3t3 /S 2 )


M(s, t) = 1 + O(s3t3 /S 2 ) Qm

S!
Qn j=1 tj !
i=1 si !
S!
Qn
= Qm j=1 tj !
i=1 si !


S3 T3
(S22 − 4S3 − 2S2 )(T22 − 4T3 − 2T2 )
S2 T2
3 3
2
+
+
+ O(s t /S ) ,
× 1+
2[S]2 3[S]3
8[S]4
= Qm

14

where we have used the fact that p0 + 2p1 + 4p2 + 6p3 = 1 + o(1) in the insubstantial case to get the second line.
This expression is equal to the expression in Theorem 1.3 under our present assumptions. (Note that since (S2 , T2 ) is insubstantial, the term S22 T22 /2S 5 which appears in the statement of Theorem 1.3 is absorbed into the error term.)
For nonnegative integers d, h, define Cd,h = Cd,h (s, t) to be the set of all pairings in
P(s, t) with exactly d double pairs and h triple pairs, but no parallel classes of multiplicity greater than 3. Also define
X
w(P ) = 2d 6h |Cd,h |.
w(Cd,h ) =
P ∈Cd,h

A special case of (3.1), used in [12], is that the number of 0-1 matrices in P(s, t) is
|C0,0 |
Qn
.
j=1 tj !
i=1 si !

N(s, t) = Qm

We will proceed by writing M(s, t) in terms of N(s, t), as follows.
Lemma 3.2. If (S2 , T2 ) is substantial then
M(s, t) = N(s, t)

N2 X
N3
X
w(Cd,h )
d=0 h=0

w(C0,0 )


1 + O(s3 t3 /S 2 ) .

Proof. By Lemma 2.5 and (3.1),
N

N

2 X
3
X

1
Qn
M(s, t) = Qm w(Cd,h ) 1 + O(s3 t3 /S 2 )
j=1 tj !
i=1 si !

d=0 h=0

= N(s, t)

N2 X
N3
X
d=0 h=0


w(Cd,h )
1 + O(s3t3 /S 2 ) .
w(C0,0 )

We will evaluate the sum in Lemma 3.2 using two summation lemmas proved in [12]
and restated below.
Lemma 3.3 ([12, Corollary 4.3]). Let 0 ≤ A1 ≤ A2 and B1 ≤ B2 be real numbers.
Suppose that there exist integers N, K with N ≥ 2 and 0 ≤ K ≤ N, and a real number c > 2e such that 0 ≤ Ac < N − K + 1 and |BN| < 1 for all A ∈ [A1 , A2 ] and B ∈ [B1 , B2 ].
Further suppose that there are real numbers δi , for 1 ≤ i ≤ N, and γi ≥ 0, for 0 ≤ i ≤ K,
PK
Pi
1
such that j=0 γj [i]j < 5 for 1 ≤ i ≤ N.
j=1 |δj | ≤
Given A(1), . . . , A(N) ∈ [A1 , A2 ] and B(1), . . . , B(N) ∈ [B1 , B2 ], define n0 , n1 , . . . , nN by n0 = 1 and

ni
A(i)
1 − (i − 1)B(i) 1 + δi )
=
ni−1
i
15

for 1 ≤ i ≤ N, with the following interpretation: if A(i) = 0 then nj = 0 for i ≤ j ≤ N.
Then
N
X
Σ1 ≤
ni ≤ Σ2 , i=0

where
K
X


γj (3A1 )j − 14 (2e/c)N ,

Σ1 = exp



A1 − 21 A21 B2 − 4

Σ2 = exp



A2 − 21 A22 B1 + 12 A32 B12 + 4

j=0

K
X

γj (3A2 )

j

j=0



+ 41 (2e/c)N .

Lemma 3.4 ([12, Corollary 4.5]). Let N ≥ 2 be an integer and, for 1 ≤ i ≤ N, let real numbers A(i), B(i) be given such that A(i) ≥ 0 and 1 − (i − 1)B(i) ≥ 0. Define
N
N
N
A1 = minN
i=1 A(i), A2 = maxi=1 A(i), C1 = mini=1 A(i)B(i) and C2 = maxi=1 A(i)B(i).
Suppose that there exists a real number ĉ with 0 < ĉ < 31 such that max{A/N, |C|} ≤ ĉ
for all A ∈ [A1 , A2 ], C ∈ [C1 , C2 ]. Define n0 , . . . , nN by n0 = 1 and

A(i)
ni
1 − (i − 1)B(i)
=
ni−1
i for 1 ≤ i ≤ N, with the following interpretation: if A(i) = 0 or 1 − (i − 1)B(i) = 0, then nj = 0 for i ≤ j ≤ N. Then
N
X
Σ1 ≤
ni ≤ Σ2
i=0

where

Σ1 = exp A1 − 21 A1 C2 − (2eĉ)N ,

Σ2 = exp A2 − 21 A2 C1 + 12 A2 C12 + (2eĉ)N .

We obtain bounds on the ratios we require by applying results from [12]. To begin with we focus on the effect of changing the number of triple pairs while keeping the number of double pairs fixed.
Lemma 3.5. Suppose 0 ≤ d ≤ N2 and 1 < h ≤ N3 , with Cd,h 6= ∅. If (S2 , T2 ) is substantial then
S3 T3 + O(s2 t2 (st + d + h)S)
w(Cd,h )
=
.
w(Cd,h−1 )
hS 3
Proof. This follows from [12, Lemma 4.6] since, for h ≥ 1, w(Cd,h )
6 |Cd,h |
=
.
w(Cd,h−1)
|Cd,h−1 |
16

Note that the values of N2 , N3 used in this paper are no smaller than, and are at most a constant factor larger than, the values used in [12]. For example, we have


N3 = max ⌈log S⌉, ⌈230000S3 T3 /S 3 ⌉ , while in [12] the value max ⌈log S⌉, ⌈7S3 T3 /S 3 ⌉
was used. Examination of the proof of [12, Lemma 4.6] shows that the bound given there for |Cd,h |/|Cd,h−1| also holds for all 0 ≤ d ≤ N2 and 1 ≤ h ≤ N3 .
Next, adapting the proof of [12, Corollary 4.7] gives:
Corollary 3.6. Suppose 0 ≤ d ≤ N2 with Cd,0 6= ∅. Further suppose that (S2 , T2 ) is substantial. Then


N3
X

w(Cd,h )
S3 T3
2 2
2
= exp
+ O s t (st + d)/S
.
w(Cd,0 )
S3
h=0

Proof. We will apply Lemma 3.4. Let h′ be the first value of h ≤ N3 for which Cd,h = ∅, or h′ = N3 + 1 if there is no such value. Define αh , 1 ≤ h < h′ , by

S3 T3 − αh s2 t2 (st + d + (h − 1)S)
|Cd,h |
=
.
(3.2)
|Cd,h−1 |
hS 3
Lemma 3.5 says that αh is bounded independently of h, d and S.
For 1 ≤ h < h′ , define
A(h) =

S3 T3 − αh (s2 t2 (st + d)S)
,
S3

C(h) =

αh s2 t2
.
S2

If αh ≤ 0 then by definition A(h) ≥ S3 T3 /S 3 , and S3 T3 > 0 since h < h′ . Therefore
A(h) > 0 in this case. If αh > 0 then C(h) > 0, which implies that A(h) > 0 since the right side of (3.2) has the same sign as A(h) − (h − 1)C(h). Therefore A(h) > 0 whenever h < h′ . Define B(h) = C(h)/A(h) for 1 ≤ h < h′ . Also define A(h) = B(h) = 0 for h′ ≤ h ≤ N3 .
Define A1 , A2 , C1 , C2 by taking the minimum and maximum of the A(h) and C(h) over
1
1 ≤ h ≤ N3 , as in Lemma 3.4. Let A ∈ [A1 , A2 ] and C ∈ [C1 , C2 ], and set ĉ = 41
. Since
3
A = S3 T3 /S + o(1) and C = o(1), we have that max{A/N3 , |C|} < ĉ for S sufficiently large, by the definition of N3 .
Therefore Lemma 3.4 applies and says that
N3
X
|Cd,h |






S3 T3
2 2
2
= exp
+ O s t (st + d)/S
+ O (2e/41)N3 .
3
|Cd,0 |
S
h=0

Finally, (2e/41)N3 ≤ (2e/41)log S ≤ S −2 . Since the sum we are estimating is at least equal to one, this additive error term is covered by the error terms inside the exponential. This completes the proof.
17

Now we must sum over pairings with no triple pairs.
Lemma 3.7. Suppose that (S2 , T2 ) is substantial and that 1 ≤ d ≤ N2 with Cd,0 6= ∅.
Then

2A(d)
w(Cd,0 )
1 − (d − 1)B (1 + δd )
=
w(Cd−1,0 )
d where


 3 3
S3 T3
2S2 T2
S2 T2
S2
T2
1 2S3 T2 2S2 T3
s t
−
−
A(d) =
1+ 2 + 2 + +
+ 2
+O
,
2
2
3
2S
S
S
S
S2 S
S T2
SS2 T2
S
S2
2
2
4T3 4S3
4
B=
+
+ 2 + 2 − ,
S2 T2
T2
S2
S


2 2
2 2
(d − 1) s
(d − 1) t dst(d + st)
δd = O
+
+
.
S22
T22
S2 T2
Proof. This follows from [12, Lemma 4.8] since, for d ≥ 1,
2 |Cd,0 |
w(Cd,0 )
=
.
w(Cd−1,0 )
|Cd−1,0 |
As in Lemma 3.5, our value of N2 is no smaller than, and is at most a constant factor larger than, the value used in [12]. Examination of the proof of [12, Lemma 4.8] shows that the expression given there for |Cd,0 |/|Cd−1,0 | also holds for 1 ≤ d ≤ N2 .
Adapting the proof of [12, Corollary 4.9] gives the following:
Corollary 3.8. If (S2 , T2 ) is substantial then
N2 X
N3
X
w(Cd,h )

 3 3 
S2 T2 S2 T2
st
.
= exp
+ 3 +O
2
2
w(C
S
S
S
0,0 )
d=0 h=0


Proof. We need to apply Lemma 3.3 to the result of Lemma 3.7, and take into account the terms coming from the triple pairs (as given by Corollary 3.6).
Let d′ be the first value of d ≤ N2 for which |Cd,0 | = 0, or d′ = N2 + 1 if no such value of d exists. Define m0 , m1 , . . . , mN2 by
N

md =

3
w(Cd,h )
w(Cd,0 ) X
w(C0,0 )
w(Cd,0)

h=0

for 0 ≤ d < d′ , and md = 0 for d′ ≤ d ≤ N2 . Then clearly
N3
N2 X
X
w(Cd,h )
d=0 h=0

w(C0,0 )
18

=

N2
X
d=0

md .

Corollary 3.6 tells us that for d < d′ we have


w(Cd,0 )
S3 T3
3 3
2
2 2
2
md =
exp
+ O(s t /S ) + ξd s t /S
w(C0,0)
S3

(3.3)

where ξ0 = 0 and in general ξd = O(d). (Note that (3.3) is also true for d′ ≤ d ≤ N2 , since both sides equal zero.) If α is a constant such that |ξd| ≤ αd for 0 ≤ d ≤ d′ , then
X
X

N2
N2
N2
X
S3 T3
S3 T3
3 3
2
3 3
2
+
O(s t
/S
)
+
O(s t
/S
)
n
(−1)
≤
m
≤
exp nd (1)
exp d
d
3
S3
S
d=0
d=0
d=0
(3.4)
where

w(Cd,0 )
nd (x) =
exp xαds2 t2 /S 2 .
w(C0,0)


Next we note that, for x ∈ {−1, 1}, n0 (x) = 1, and for 1 ≤ d ≤ d′ ,


nd (x)
= 2A(d) 1 − (d − 1)B 1 + δd nd−1 (x)

with A(d), B, and δd satisfying the expressions given in the statement of Lemma 3.7.
This follows since the factor exp(xαs2 t2 /S 2 ) is covered by the error term on A(d). For d′ ≤ d ≤ N2 define A(d) = 0.
Now let A1 = A1 (x) = mind 2A(d), A2 = A2 (x) = maxd 2A(d), where the maximum and minimum are taken over 1 ≤ d ≤ N2 . Also let B1 = B2 = B, and K = 3, and define c = S 1/4 if S2 T2 < S 7/4 and c = 41 otherwise. The conditions of Lemma 3.3 now hold as we will show. Let A ∈ [A1 , A2 ] be arbitrary.
Clearly c > 2e. If S2 T2 < S 7/4 then N2 = 22. Using the condition S2 T2 ≥ (st)3/2 S
implied by the substantiality of (S2 , T2 ), we find that Ac = 1 + o(1). For S2 T2 ≥ S 7/4 ,
Ac = 41S2 T2 /S 2 (1 + o(1)). It is also easy to check that BN2 = o(1). Thus, in all cases we have that Ac < N2 − 2 and |BN2 | < 1 for sufficiently large S.
If d = O(S2T2 /S 2 ) then
N2
X

 2

 3 3
s S2 T23 t2 S23 T2 stS22 T22 s2 t2 S2 T2
s t
=O
= o(1),
|δd | = O
+
+
+
6
6
6
4
S
S
S
S
S2
d=1

while if d ≤ ⌈log S⌉ then
N2
X


 2 3
s log S t2 log3 S st log3 S s2 t2 log2 S
+
+
= o(1).
+
|δd | = O
2
2
S
T
S
S
2 T2
2 T2
2
2
d=1

19

Finally, for 1 ≤ k ≤ N2 , we have

X

X

k k
d2 st ds2 t2
t2 
+O
|δd | = O
+O
(d − 1)
+
S22 T22
S2 T2
S2 T2
d=1
d=1
d=1
d=1


 s2
t2  k(k + 1)(2k + 1)st k(k + 1)s2 t2
+
= O k(k − 1)(2k − 1) 2 + 2 +
S2
T2
S2 T2
S2 T2
K
X
≤
γj [k]j ,

k
X

X
k

2

 s2

j=0

where

 2

 2

s2 t2
t2
s2 t2
t2
st s
s
γ0 = 0, γ1 = O
, γ2 = O
+
+
, γ3 = O
+
+
.
S2 T2
S22 T22 S2 T2
S22 T22 S2 T2


Since N23 (s2 /S22 + t2 /T22 + st/S2 T2 ) = o(1), which is easily checked, it follows that
PK
j=0 γj [k]j < 1/5 for 1 ≤ k ≤ N2 , when S is large enough.
Therefore the conditions of Lemma 3.3 hold, and we conclude that each of the bounds
P 2
given by that lemma for N
d=0 nd (x) has the form

3


X

2
3 2
j
1
+ O (2e/c)N2 ,
γj (3A)
exp A − 2 A B + O A B +
j=0

where A is either A1 or A2 . A somewhat tedious check shows that
O(A3 B 2 ) +

3
X

γj (3A)j = O(s3t3 /S 2 ).

j=0


Next consider the error term O (2e/c)N2 . If N2 = 22 then (2e/c)N2 = (2eS −1/4 )22 =
O(S −2), while in the other cases we have (2e/c)N2 = (2e/41)N2 ≤ (2e/41)log S = O(S −2).
Since n0 = 1, this additive error term is covered by a relative error of the same form.
P 2
Therefore, each of the bounds on N
d=0 nd (x) has the form


 3 3 
 s3 t3 
S2 T2 S2 T2 S3 T3
s t
2
1
exp A − 2 A B + O
= exp
+ 3 − 3 +O
.
S2
S2
S
S
S2
Modulo the given error terms, the final expression does not depend on x, nor on whether we are taking a lower bound or upper bound in Lemma 3.3. To complete the proof, just apply (3.4).
Corollary 3.8 and Lemma 3.2 together prove Theorem 1.3 in the substantial case. The insubstantial case was already proved in Lemma 3.1.
20

4

Alternative formulation

We now derive an alternative formulation of Theorem 1.3. Recall the definition of µ̂k and
ν̂k given in the Introduction.
Corollary 4.1. Under the conditions of Theorem 1.3,

 n 
m 
Y
n+si −1 Y m+tj −1
tj si j=1
i=1


M(s, t) =
mn+S−1
S



1 3 − µ̂2 ν̂2
+
× exp (1 − µ̂2 )(1 − ν̂2 )
2
4S
(1 − µ̂2 )(3 + µ̂2 − 2µ̂2ν̂2 ) (1 − ν̂2 )(3 + ν̂2 − 2µ̂2 ν̂2 )
−
4n
4m
 3 3 
2
2
s t
(1 − 3µ̂2 + 2µ̂3 )(1 − 3ν̂2 + 2ν̂3 )
.
+O
+
12S
S2

−

Proof. By Stirling’s formula or otherwise,




[x]3
[x]2
[x]2
Nx
N+x−1
4
3
exp
−
−
+ O(x /N )
=
x!
2N
6N 2 4N 2
x as N → ∞, provided that the error term is bounded. This gives us the approximations

 3 3 

m 
Y
nS
n+si −1
S2
S3
st
S2
=Q
− 2 − 2 +O
exp si
2n 4n
6n
S2
i si !
i=1
 3 3 


n 
Y
T2
T3
s t
T2
mS
m+tj −1
−
−
+O
exp
=Q
2
2
2m 4m
6m
S2
tj j tj !
j=1


 2
 3 3 
(mn)S
mn+S−1
S
S3
S
st
=
.
exp
−
−
+O
2
2
S
S!
2mn 2mn 6m n
S2
Substitute these expressions into Theorem 1.3 and replace S2 , S3 , T2 , T3 by their equivalents in terms of µ̂2 , µ̂3 , ν̂2 , ν̂3 . The desired result is obtained.
As noted in the Introduction, Theorem 1.3 establishes the conjecture recalled after
Theorem 1.1 in some cases. Using Corollary 4.1, the following is easily seen. (Note that
µ̂2 = µ̂3 = ν̂2 = ν̂3 = 0 in the semiregular case.)

Corollary 4.2. If s = s(m, n) and t = t(m, n) satisfy ms = nt and st = o (mn)1/5 , then

5(s + t)
1 + o(1) .
∆(m, s; n, t) =
6st
21

Most of the terms inside the exponential of Corollary 4.1 are tiny unless at least one of µ̂2 , ν̂2 is quite large (that is, the graph is very far from semiregular). In particular we can now prove Corollary 1.5 which was stated in the Introduction.
Proof of Corollary 1.5. It is only necessary to check that the additional terms in Corollary 4.1 have the required size. It helps to realise that µ̂2 ≤ s, |µ̂3 | ≤ sµ̂2 , ν̂2 ≤ t and
|ν̂3 | ≤ tν̂2 .
A random nonnegative m × n matrix with entries summing to S is just a random composition of S into mn parts. (A composition is an ordered sum of nonnegative numbers.)
In particular, for 1 ≤ i ≤ m the row sum si satisfies




S + mn − 1
k + n − 1 S − k + (m − 1)n − 1
(0 ≤ k ≤ S).
Pr(si = k) =
S
S−k k
From this we can compute the following expected values.
n(m − 1)
, mn + 1
n(m − 1)(m − 2)(mn + 2S)
E µ̂3 =
, m(mn + 1)(mn + 2)

m(n − 1)
, mn + 1
m(n − 1)(n − 2)(mn + 2S)
E ν̂3 =
.
n(mn + 1)(mn + 2)

E µ̂2 =

E ν̂2 =

The first two expectations suggest that the argument of the exponential in Corollary 1.5
is close to 0 with high probability for such a random matrix. We will prove this in a future paper, and note that the result gives a model for the row and column sums of random matrices.

5

Restricted sets of allowed entries

Given a subset J of the nonnegative integers, let M(s, t, J ) denote the set of matrices in
M(s, t) with all entries in the set J . Let M(s, t, J ) = |M(s, t, J )|. By generalising the techniques of the preceding sections, we can find an asymptotic expression for M(s, t, J )
whenever 0, 1 ∈ J .
Lemma 5.1. Let J ⊆ N with 0, 1 ∈ J . Define χ2 = 0 if 2 ∈
/ J , χ2 = 1 if 2 ∈ J , and similarly χ3 . Then

 3 3 
S2 T2
S2 T2
S3 T3
st
M(s, t, J ) = N(s, t) exp χ2 2 + χ2 3 + (χ3 − χ2 ) 3 + O
S
S
S
S2

S2 T2
S2 T2
S3 T3
S!
Qn exp (χ2 − 21 ) 2 + (χ2 − 12 ) 3 + (χ3 − χ2 + 13 ) 3
= Qm
S
S
S
i=1 si !
j=1 tj !
 3 3 
2
2
2 2
S2 T2 (S2 + T2 ) S2 T3 + S3 T2
S T
s t
−
−
+ 2 52 + O
.
4
4
4S
2S
2S
S2
22

Proof. Our general approach will be similar to that we used for Theorem 1.3, but the methods of Section 2 will need significant modification. The source of the problem is that a D-switching may introduce an entry that is not in J .
For Q ∈ M(s, t) and i ≥ 1, let ni (Q) be the number of entries of Q equal to i. Also
P
let n≥5 (Q) = i≥5 ni (Q). Define N2 and N3 as in Section 2 when (S2 , T2 ) is substantial, and N2 = 2 and N3 = 1 when (S2 , T2 ) is insubstantial. For Q ∈ M+ , let
⌈(st)1/4 ⌉
+

E (Q) =

X

−

nD (Q),

E (Q) =

X

nD (Q).

D=5

D>⌈(st)1/4 ⌉

Consider the following subsets of M(s, t):
M+ = M(s, t, J ∪ {4, 5, 6, . . . }),

M = M(s, t, J ),

M− = Q ∈ M+ n2 (Q), n3 (Q), n4 (Q) ≤ S 5/6 ,

E − (Q) ≤ ⌈2(stS)1/2 ⌉, E + (Q) ≤ ⌈2(st)1/4 S 1/2 ⌉},


M∗ = Q ∈ M(s, t, J ∩ {0, 1, 2, 3}) n2 (Q) ≤ N2 , n3 (Q) ≤ N3 .

Also define the cardinalities M + , M, M − , M ∗ , respectively. By monotonicity, we have
M ∗ ≤ M ≤ M + and M ∗ ≤ M − ≤ M + .
We now employ switchings to establish that M + −M − < M − −M ∗ and M − −M ∗ =

O(s3 t3 /S 2 )M ∗ , from which it follows that M = 1 + O(s3 t3 /S 2 ) M ∗ . (OK?)
We start with the claim that M + −M − < M − −M ∗ . Let Q ∈ M+ − M− such that
E + (Q) > ⌈2(st)1/4 S 1/2 ⌉. We will use the following switching, illustrated by this operation performed on submatrices:




D1 − 1
1
D1 0
(5.1)
7→
1
D2 − 1
0 D2
where D1 , D2 ≥ (st)1/4 . The number of forward switchings is bounded below by

E + (Q)2 − O stE + (Q) = E + (Q)(1 − o(1)),

and the number of reverse switchings is bounded above by
√
2stS
st S.
=
2
(st)1/2

Hence the number of reverse switchings divided by the number of forward switchings is bounded above by
√
2(1 + o(1)) stS
1 + o(1)
≤
< 23 ,
+
2
E (Q)
2
23

using the assumed lower bound on E + (Q). After repeatedly applying this switching, we reach a matrix Q which satisfies
E + (Q) ≤ ⌈2(st)1/4 S 1/2 ⌉.

(5.2)

The next switching is applied to matrices Q ∈ M+ for which (5.2) holds but E − (Q) >
⌈2(stS)1/2 ⌉. The switching that we used is the same as that shown in (5.1) except that now D1 , D2 ∈ {5, . . . , ⌈(st)1/4 ⌉}. The number of forward switchings is bounded below by

E − (Q)2 − O stE − (Q) = E − (Q)2 (1 − o(1)), and the number of reverse switchings is bounded above by
2stS.
Hence the number of reverse switchings divided by the number of forward switchings is bounded above by
2(1 + o(1))stS
1 + o(1)
≤
< 23 .
−
2
E (Q)
2
We apply this switching until we reach a matrix Q which satisfies both (5.2) and
E − (Q) ≤ ⌈2(stS)1/2 ⌉.

(5.3)

To analyse these two switchings using Theorem 2.1, we can define the sets
X
C(i) = {Q ∈ M+ − M− |
DnD (Q) = i}.
D≥2

If Q ∈ C(i) and R can be obtained from Q using one of the switchings described above, then R ∈ C(i − 2). This leads to an acyclic directed graph in each case, and we have shown above that all the ratios b(vi )/a(vi−1 ) in Theorem 2.1 are at most 2/3.
Next we need to reduce each of n2 (Q), n3 (Q), n4 (Q) to below S 5/6 . We achieve this using a succession of three types of switchings, illustrated by the following operations on submatrices: for example, the switching




1 1 1 1
4 0 0 0




 1 1 1 1
0 4 0 0

 7→ 

 1 1 1 1
0 0 4 0
1 1 1 1
0 0 0 4
will be used to reduce n4 (Q) (with analogous operations for D = 2, 3). First we apply the switching for D = 4 until n4 (Q) ≤ S 5/6 , then the switching for D = 3 until n3 (Q) ≤ S 5/6 , finally applying the switching for D = 2 until n2 (Q) ≤ S 5/6 . As a representative example,
24

take the switching for D = 4. By counting similarly to Lemma 2.2, this switching can be applied to Q in at least (n4 (Q) − O(st))4 ways, and the inverse can be applied in at most Ss3 t3 ways. For n4 (Q) > S 5/6 , the condition s3 t3 = o(S 2 ) implies that Ss3 t3 =

o (n4 (Q) − O(st))4 , so the ratios denoted by b(vi )/a(vi−1 ) in Theorem 2.1 are all o(1).
Since none of the switchings can undo the work of a previous switching, the end result is a matrix M− \ M∗ . (Note that in the resulting matrix R, at least one of E + (R),
E − (R), n2 (R), n3 (R) or n4 (R) will be just under the threshold value. This implies that
R 6∈ M∗ .) This establishes the bound M + −M − < M − −M ∗ .
For any matrix Q ∈ M− \ M∗ we have
X
DnD (Q) ≤ min{s, t}E + (Q) ≤ 3(st)3/4 S 1/2 ,
D≥⌈(st)1/4 ⌉

since using (5.2) and since min{s, t} ≤ (st)1/2 . Similarly, (5.3) implies that
⌈(st)1/4 ⌉

X

DnD (Q) ≤ ⌈(st)1/4 ⌉E − (Q) ≤ 3(st)3/4 S 1/2 ,

X

DnD (Q) ≤ 6(st)3/4 S 1/2 + 3S 5/6 = o(S).

D=5

which leads to
D≥2

Hence when Q ∈ M− \ M∗ , we know that n1 (Q) = S − o(S). We can now continue precisely as in Lemmas 2.3, 2.5, using D-switchings restricted to q1 = · · · = qD = 1.
This restriction ensures that D-switchings only create entries with value equal to 0 or
1. The various switching counts can be taken as essentially the same as before, since all but a vanishing fraction of the non-zero entries are 1. We conclude that M − −M ∗ =

O(s3 t3 /S 2 )M ∗ which, as noted above, implies that M = 1 + O(s3 t3 /S 2 ) M ∗ .
Having now reduced the task to evaluation of M ∗ , we can complete the proof following
Lemma 3.1 in the insubstantial case, and Section 4 in the substantial case. In Lemma 3.1
the only modification is to replace the expression p0 + 2p1 + 4p2 + 6p3 by p0 + 2χ2 p1 +
4χ2 p2 + 6χ3 p3 .
Now suppose that (S2 , T2 ) is substantial. If χ2 = χ3 then the result is given by either
Theorem 1.2 or Theorem 1.3. If χ2 = 0 and χ3 = 1 then the result follows from applying
Corollary 3.6 with d = 0, since arguing as in Lemma 3.2 shows that
M(s, t, J ) = N(s, t)

N3
X
w(C0,h )
h=0

w(C0,0 )

1 + O(s3t3 /S 2 )



in this case. Finally, if χ2 = 1 and χ3 = 0 then
M(s, t, J ) = N(s, t)

N2
X
w(Cd,0 )
d=0

w(C0,0 )

25

1 + O(s3t3 /S 2 )



so in place of (3.3) we simply have md = w(Cd,0 )/w(C0,0 ) = nd (0) for 0 ≤ d ≤ N2 .
The remainder of the proof is identical except that there is no need to apply (3.4) at the end.

References
[1] A. Barvinok, A. Samorodnitsky and A. Yong, Counting magic squares in quasipolynomial time, preprint (2007); http://www.arxiv.org/abs/math/0703227.
[2] A. Békéssy, P. Békéssy and J. Komlós, Asymptotic enumeration of regular matrices,
Studia Sci. Math. Hungar., 7 (1972) 343–353.
[3] E. A. Bender, The asymptotic number of nonnegative integer matrices with given row and column sums, Discrete Math., 10 (1974) 345–353.
[4] B. Bollobás, Random Graphs (2nd edn.), Cambridge University Press, Cambridge,
2001.
[5] E. R. Canfield and B. D. McKay, Asymptotic enumeration of integer matrices with large equal row and column sums, Combinatorica, 30 (2010) 655–680.
[6] P. Diaconis and B. Efron, Testing for independence in a two-way table: new interpretations of the chi-square statistic (with discussion), Ann. Statist., 13 (1995)
845–913.
[7] P. Diaconis and A. Gangolli, Rectangular arrays with fixed margins, in: Discrete
Probability and Algorithms, IMA Volumes on Mathematics and its Applications, vol. 72, Springer, New York, (1995), pp. 15–41.
[8] M. Dyer, R. Kannan and J. Mount, Sampling contingency tables, Random Structures Algorithms, 10 (1997), 487–506.
[9] C. J. Everett, Jr., and P. R. Stein, The asymptotic number of integer stochastic matrices, Discrete Math., 1 (1971) 33–72.
[10] V. Fack and B. D. McKay, A generalized switching method for combinatorial estimation, Australas. J. Combin., 39 (2007), 141–154.
[11] C. Greenhill and B. D. McKay, Asymptotic enumeration of sparse nonnegative integer matrices with specified row and column sums, Advances in Applied Mathematics, 41 (2008), 459–481.

26

[12] C. Greenhill, B. D. McKay and X. Wang, Asymptotic enumeration of sparse 0-1
matrices with irregular row and column sums, J. Combin. Theory Ser. A, 113 (2006)
291–324.
[13] B. D. McKay, Asymptotics for 0-1 matrices with prescribed line sums, in: Enumeration and Design, Academic Press, Toronto, 1984, pp. 225–238.
[14] B. Morris, Improved bounds for sampling contingency tables, Random Structures
Algorithms, 21 (2002) 135–146.
[15] R. C. Read, Some enumeration problems in graph theory, Doctoral Thesis, London
University, (1958).
[16] R. P. Stanley, Combinatorics and Commutative Algebra, Progress in Mathematics, vol. 41, Birkhäuser, Boston, 1983.

27

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
