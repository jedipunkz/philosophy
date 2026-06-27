---
source: "https://arxiv.org/abs/math/0606496v2"
title: "Asymptotic enumeration of dense 0-1 matrices with specified line sums"
author: "E. Rodney Canfield, Catherine Greenhill, Brendan D. McKay"
year: "2006"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/math/0606496v2"
pdf: "https://arxiv.org/pdf/math/0606496v2"
captured_at: "2026-06-24T11:11:15Z"
updated_at: "2026-06-24T11:11:15Z"
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

# Asymptotic enumeration of dense 0-1 matrices with specified line sums

- 著者: E. Rodney Canfield, Catherine Greenhill, Brendan D. McKay
- 年: 2006
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/math/0606496v2)
- ダウンロード: https://arxiv.org/pdf/math/0606496v2
- PDF: https://arxiv.org/pdf/math/0606496v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Let S=(s_1,s_2,..., s_m) and T = (t_1,t_2,..., t_n) be vectors of non-negative integers with sum_{i=1}^{m} s_i = sum_{j=1}^n t_j. Let B(S,T) be the number of m*n matrices over {0,1} with j-th row sum equal to s_j for 1 <= j <= m and k-th column sum equal to t_k for 1 <= k <= n. Equivalently, B(S,T) is the number of bipartite graphs with m vertices in one part with degrees given by S, and n vertices in the other part with degrees given by T. Most research on the asymptotics of B(S,T) has focused on the sparse case, where the best result is that of Greenhill, McKay and Wang (2006). In the case of dense matrices, the only precise result is for the case of equal row sums and equal column sums (Canfield and McKay, 2005). This paper extends the analytic methods used by the latter paper to the case where the row and column sums can vary within certain limits. Interestingly, the result can be expressed by the same formula which holds in the sparse case.

## PDF Text

arXiv:math/0606496v2 [math.CO] 21 Jan 2007

Asymptotic enumeration of dense 0-1 matrices with specified line sums
E. Rodney Canfield∗

Catherine Greenhill†

Department of Computer Science
University of Georgia
Athens, GA 30602, USA

School of Mathematics and Statistics
University of New South Wales
Sydney, Australia 2052

ercanfie@uga.edu

csg@unsw.edu.au

Brendan D. McKay‡
Department of Computer Science
Australian National University
Canberra ACT 0200, Australia bdm@cs.anu.edu.au

Abstract
Let s = (s1 , s2 , . . . , sm ) and t = (t1 , t2 , . . . , tn ) be vectors of non-negative integers
P
Pn with m i=1 si =
j=1 tj . Let B(s, t) be the number of m × n matrices over {0, 1}
with j-th row sum equal to sj for 1 ≤ j ≤ m and k-th column sum equal to tk for
1 ≤ k ≤ n. Equivalently, B(s, t) is the number of bipartite graphs with m vertices in one part with degrees given by s, and n vertices in the other part with degrees given by t. Most research on the asymptotics of B(s, t) has focused on the sparse case, where the best result is that of Greenhill, McKay and Wang (2006). In the case of dense matrices, the only precise result is for the case of equal row sums and equal column sums (Canfield and McKay, 2005). This paper extends the analytic methods used by the latter paper to the case where the row and column sums can vary within certain limits. Interestingly, the result can be expressed by the same formula which holds in the sparse case.
∗

Research supported by the NSA Mathematical Sciences Program.
Research supported by the UNSW Faculty Research Grants Scheme.
‡
Research supported by the Australian Research Council.

†

1

1

Introduction

Let s = (s1 , s2 , . . . , sm ) and t = (t1 , t2 , . . . , tn ) be vectors of positive integers with
Pm
Pn i=1 si =
j=1 tj . Let B(s, t) be the number of m × n matrices over {0, 1} with jth row sum equal to sj for 1 ≤ j ≤ m and k-th column sum equal to tk for 1 ≤ k ≤ n.
Equivalently, B(s, t) is the number of labelled bipartite graphs with m vertices in one part of the bipartition with degrees given by s, and n vertices in the other part of the bipartition with degrees given by t. Let s be the average value of s1 , s2 , . . . , sm and let t be the average value of t1 , t2 , . . . , tn . Define the density λ = s/n = t/m, which is the fraction of entries in the matrix which equal 1.
The asymptotic value of B(s, t) has been much studied, especially since the celebrated
Gale-Ryser Theorem [10] that characterises (s, t) such that B(s, t) > 0. Various authors have considered the semiregular case, where sj = s for 1 ≤ j ≤ m and tk = t for
1 ≤ k ≤ n. Write B(m, s; n, t) for B(s, t) in this case. For the sparse (low-λ) semiregular case, the best result is by McKay and Wang [7] who gave an asymptotic expression for

B(m, s; n, t) which holds when st = o (mn)1/2 . In the dense (λ not close to 0 or 1)
semiregular case, Canfield and McKay [1] used analytic methods to obtain an asymptotic expression for B(m, s; n, t) in two ranges: in the first, the matrix is relatively square and the density is not too close to 0 or 1, while in the second, the matrix is much wider than high (or vice-versa) but the density is arbitrary. For the sparse irregular case, the best result is that of Greenhill, McKay and Wang [2], who gave an asymptotic expression for

B(s, t) which holds when max{sj } max{tk } = o (λmn)2/3 .
See [1], [2] and [7] for a more extensive historical survey.

The contribution of this paper is to adapt the approach of [1] to the dense irregular case when the matrix is relatively square and the density is not too close to 0 or 1. See
McKay and Wormald [8] for the corresponding calculation for symmetric matrices.
In keeping with these earlier papers, the asymptotic value of B(s, t) can be expressed by a very nice formula involving binomial coefficients. We now state our theorem.
Theorem 1. Let s = s(m, n) = (s1 , s2 , . . . , sm ) and t = t(m, n) = (t1 , t2 , . . . , tn ) be vecP
P
P
tors of positive integers such that m sj = nk=1 tk for all m, n. Define s = m−1 m j=1
j=1 sj ,
−1 Pn
1
t = n k=1 tk , λ = s/n = t/m and A = 2 λ(1 − λ). For some ε > 0, suppose that
|sj − s| = O(n1/2+ε ) uniformly for 1 ≤ j ≤ m, and |tk − t| = O(m1/2+ε ) uniformly for
Pn
P
2
2
1 ≤ k ≤ n. Define R = m k=1 (tk − t) . Let a, b > 0 be constants j=1 (sj − s) and C =
such that a + b < 21 . Suppose that m, n → ∞ with n = o(A2 m1+ε ), m = o(A2 n1+ε ) and


(1 − 2λ)2
5m
5n
1+
≤ a log n.
+
8A
6n
6m

2

Then, provided ε > 0 is small enough, we have
B(s, t) =




 
−1 Y
n  
m  Y
R 
C 
1
m n
mn
−b
1−
+ O(n ) .
exp − 1 −
2
2Amn
2Amn t
s
λmn k
j j=1
k=1

Proof. The proof of this theorem is the topic of the paper; here we will summarize the main phases and draw their conclusions together. The basic idea is to identify B(s, t) as a coefficient in a multivariable generating function and to extract that coefficient using the saddle-point method. In Section 2, equation (1), we write B(s, t) = P (s, t)I(s, t), where
P (s, t) is a rational expression and I(s, t) is an integral in m + n complex dimensions.
Both depend on the location of the saddle point, which is the solution of some nonlinear equations. Those equations are solved in Section 3, and this leads to the value of P (s, t)
in (20). In Section 4, the integral I(s, t) is estimated in a small region R′ defined in (33).
The result is given by Theorem 2 together with (24). Finally, in Section 5, it is shown that the integral I(s, t) restricted to the exterior of R′ is negligible. The present theorem thus follows from (1), (20), Theorems 2–3 and (24).
Note that the error term in the above slightly improves the error term for the semiregular case proved in [1].
Thereom 1 has an instructive interpretation. Write it as B(s, t) = NP1 P2 E, where

n  
m  
Y
Y
m n
mn
−1
−1
,
, P2 = N
, P1 = N
N=
tk sj
λmn j=1
k=1
 

1
R 
C 
−b
E = exp − 1 −
1−
+ O(n ) .
2
2Amn
2Amn


Clearly, N is the number of m × n binary matrices with λmn ones. P1 is the probability that a matrix randomly chosen from this class has row sums s, while P2 is the probability of the similar event of having column sums t. If these two events were independent, we would have E = 1, so E can be taken as a measure of their non-independence. For the case when s and t are vectors of constants, that is, R = C = 0, Ordentlich and Roth [9]
proved that E ≤ 1.

It is proved in [2] that the same formula for B(s, t) modulo the error term also holds in the sparse case. Specifically, it holds with a different vanishing error term whenever



max{sj } max{tk } = o (λmn)2/3 , R + C = O (λmn)4/3 and RC = O (λmn)7/3 . In [1], evidence is presented that the formula is universal in the semiregular case (R = C = 0)
and it is tempting to conjecture that the same is true in the irregular case for a wide range of R, C values.
We will use a shorthand notation for summation over doubly subscripted variables. If
3

xjk is a variable for 1 ≤ j ≤ m and 1 ≤ k ≤ n, then xj• =

n
X

xjk ,

x•k =

xj∗ =

xjk ,

x•• =

j=1

k=1

n−1
X

m
X

xjk ,

x∗k =

m−1
X

xjk ,

j=1 k=1

xjk ,

x∗∗ =

j=1

k=1

m X
n
X

m−1
n−1
XX

xjk ,

j=1 k=1

for 1 ≤ j ≤ m and 1 ≤ k ≤ n.

Throughout the paper, the asymptotic notation O(f (m, n)) refers to the passage of m e (m, n)), which is to be taken as a and n to ∞. We also use a modified notation O(f

O(1)ε
shorthand for O f (m, n)n
. In this case it is important that the O(1) factor is
−1

uniform over ε provided ε is small enough; for example we cannot write f (m, n)n(ε )ε
e (m, n)) even though ε−1 = O(1) (ε being defined as a constant). Under the as O(f e
e assumptions of Theorem 1, we have m = O(n)
and n = O(m).
We also have that
−1
−1
c1 c2 +c3 ε c4 +c5 ε
e e c2 +c4 ) if
8 ≤ A ≤ O(log n), so A = O(1). More generally, A m n
= O(n c1 , c2 , c3 , c4 , c5 are constants.

2

Expressing the desired quantity as an integral

In this section we express B(s, t) as a contour integral in (m + n)-dimensional complex space, then begin to estimate its value using the saddle-point method.
Firstly, notice that B(s, t) is the coefficient of xs11 · · · xsmm y1t1 · · · yntn in the function m Y
n
Y

(1 + xj yk ).

j=1 k=1

By Cauchy’s coefficient theorem this equals
Qm Qn
I
I
1
k=1 (1 + xj yk )
j=1
B(s, t) =
···
dx1 · · · dxm dy1 · · · dyn , m+n s1 +1
(2πi)
x1 · · · xsmm +1 y1t1 +1 · · · yntn +1
where each integral is along a simple closed contour enclosing the origin anticlockwise.
It will suffice to take each contour to be a circle; specifically, we will write xj = qj eiθj

and

for 1 ≤ j ≤ m and 1 ≤ k ≤ n. Also define
λjk =

yk = rk eiφk

qj rk
1 + qj rk
4


for 1 ≤ j ≤ m and 1 ≤ k ≤ n. Then 1 + xj yk = (1 + qj rk ) 1 + λjk (ei(θj +φk ) − 1) , so
Qm Qn k=1 (1 + qj rk )
j=1
B(s, t) =
sj Qn tk m+n Qm
(2π)
k=1 rk j=1 qj
(1)
Z π
Z π Qm Qn i(θj +φk )
(1
+
λ
(e
−
1))
jk k=1
j=1
Pn
P
dθdφ,
×
···
m s
θ
exp(i
−π
−π
k=1 tk φk )
j=1 j j + i

where θ = (θ1 , . . . , θm ) and φ = (φ1 , . . . , φn ). Write B(s, t) = P (s, t)I(s, t) where P (s, t)
denotes the factor in front of the integral in (1) and I(s, t) denotes the integral. We will choose the radii qj , rk so that there is no linear term in the logarithm of the integrand of
I(s, t) when expanded for small θ, φ. This gives the equation m X
n
X
j=1 k=1

λjk (θj + φk ) −

m
X
j=1

sj θj −

n
X

tk φk = 0.

k=1

For this to hold for all θ, φ, we require
λj• = sj

(1 ≤ j ≤ m),

λ•k = tk

(1 ≤ k ≤ n).

(2)

In Section 3 we show that (2) has a solution, and determine to sufficient accuracy the various functions of the radii, such as P (s, t), that we require. In Section 4 we evaluate the integral I(s, t) within a certain region R defined in (22). Section 5 contains the proof that the integral is concentrated within the region R.

3

Locating the saddle-point

In this section we solve (2) and derive some of the consequences of the solution. As with the whole paper, we work under the assumptions of Theorem 1.
n
Change variables to {aj }m j=1 , {bk }k=1 as follows:

qj = r

1 + aj
2

,

rk = r

r

λ
.
1−λ

1 − r aj

where r=

1 + bk
,
1 − r 2 bk

(3)

Equation (2) is slightly underdetermined, which we will exploit to impose an additional condition. If {qj }, {rk } satisfy (2) and c > 0 is a constant, then {cqj }, {rk /c} also satisfy (2). From this we can see that, if there is a solution to (2) at all, there is one
5

P
P
P
for which m aj < 0 and nk=1 bk > 0, and also a solution for which m j=1
j=1 aj > 0 and
Pn b
<
0.
It follows from the
Intermediate
Value
Theorem that there is a solution for k=1 k which m
n
X
X
n aj = m bk ,
(4)
j=1

k=1

so we will seek a common solution to (2) and (4).
From (3) we find that
λjk /λ = 1 + aj + bk + Zjk , where
Zjk =

aj bk (1 − r 2 − r 2 aj − r 2 bk )
,
1 + r 2 aj bk

(5)

(6)

and that equations (2) can be rewritten as n
Zj•
sj − s 1 X
bk −
−
aj =
λn n k=1
n

m
Z
1 X
t −t aj − •k
−
bk = k
λm m j=1
m

(1 ≤ j ≤ m),
(7)
(1 ≤ k ≤ n).

Summing (7) over j, k, we find in both cases that n

m
X

aj + m

j=1

n
X
k=1

bk = −Z•• .

(8)

Equations (4) and (8) together imply that n

m
X

aj = m

j=1

n
X
k=1

bk = − 21 Z•• .

Substituting back into (7), we obtain aj = Aj (a1 , . . . , am , b1 , . . . , bn ), bk = Bk (a1 , . . . , am , b1 , . . . , bn ), for 1 ≤ j ≤ m, 1 ≤ k ≤ n, where sj − s Zj•
Z
−
+ •• ,
λn n
2mn
Z
t − t Z•k
−
+ •• .
Bk (a1 , . . . , am , b1 , . . . , bn ) = k
λm m
2mn
Aj (a1 , . . . , am , b1 , . . . , bn ) =

6

(9)

(0)

(0)

Equation (9) suggests an iteration. Start with aj = bk = 0 for all j, k, and, for each
ℓ ≥ 0, define
(ℓ+1)

(ℓ)
= Aj (a1 , . . . , a(ℓ)
m , b1 , . . . , bn ),

(ℓ+1)

(ℓ)
= Bk (a1 , . . . , a(ℓ)
m , b1 , . . . , bn ),

aj

bk
(ℓ)

(ℓ)

(ℓ)

(ℓ)

(ℓ)

(ℓ)

(10)

(ℓ)

(ℓ)

(ℓ)
where Zj• = Zj• = Zj• (a1 , . . . , am
, b1 , . . . , b(ℓ)
n ) and similarly for Z•k = Z•k and Z•• =
(ℓ)
Z•• . We will show that this iteration converges to a solution of (9) using a standard contraction-mapping argument.

Recall that A−1 = O(log n) under the assumptions of Theorem 1 (which we are adopting throughout). This implies that r 2 = O(log n). Therefore, within the region A defined by |aj |, |bk | ≤ n−1/3 for all j, k, we have that
∂Zj,k
= o(m−1/4 ) and
∂aj

∂Zj,k
= o(n−1/4 ),
∂bk

which imply that, in the same region, we have
(
o(m−1/4 ) (j ′ = j)
∂Aj
=
∂aj ′
o(m−5/4 ) (j ′ 6= j),
∂Bk
= o(m−5/4 ),
∂aj

∂Aj
= o(n−5/4 ),
∂bk
(
o(n−1/4 ) (k ′ = k)
∂Bk
=
∂bk′
o(n−5/4 ) (k ′ =
6 k).

Therefore, by the mean value theorem, we have for ℓ ≥ 1 that
(ℓ+1)

max |aj j

(ℓ)

(ℓ+1)

− aj | + max |bk

(ℓ)

(ℓ)

(ℓ−1)

− bk | = o(m−1/4 ) max |aj − aj j

k

(ℓ)

|

(ℓ−1)

+ o(n−1/4 ) max |bk − bk k

(ℓ−1)

provided {aj

(ℓ−1)

} ∪ {bk

(ℓ)

|,

(ℓ)

} ∪ {aj } ∪ {bk } ⊆ A.

Applying the iteration once, we have
(1)

(1)

aj = (sj − s)/(λn) and bk = (tk − t)/(λm).
(0)

(0)

(1)

(1)

(ℓ)

(ℓ)

(ℓ)

(ℓ−1)

Since {aj }, {bk } and {aj }, {bk } lie inside 21 A, we find by induction that {aj }, {bk }
ℓ
lie in ℓ+1
A for all ℓ. Moreover, the iteration is Cauchy-convergent in the maximum norm,
(ℓ)

(ℓ)

(ℓ)

(ℓ−1)

and the error in stopping at {aj }, {bk } is at most maxj |aj −aj

|+maxk |bk −bk
(l)

(l)

|.

When we carry out this iteration, we find that all the encountered aj and bk values e −1/2 ). It helps to know that the following approximation holds in that case: are O(n e −5/2 ).
Zjk = (1 − r 2 )aj bk − r 2 aj b2k − r 2 a2j bk − r 2 (1 − r 2 )a2j b2k + O(n
7

Using the fact that

Pm

Pn (1)
(1)
j=1 aj = 0 and k=1 bk = 0, we find that
(1)
(1)
Zj• = −r 2 aj
(1)

(1)

Z•k = −r 2 bk
(1)
e
Z••
= O(1).

Therefore,

n
X

k=1
m
X
j=1

(1) 2

bk

(1) 2

aj

e −1 ),
+ O(n

e −1 ),
+ O(n

(sj − s)C
sj − s e −2
+ 2
2 2 + O(n ) (1 ≤ j ≤ m),
λn
λ (1 − λ)m n t −t
(t − t)R
(2)
e −2
bk = k
+ 2 k
2 2 + O(n ) (1 ≤ k ≤ n).
λm
λ (1 − λ)m n
(2)

aj =

Similarly,
(2)

(2)

Zj• = −r 2 aj

n
X
k=1

(2) 2

bk

(2) 2

− r 2 (1 − r 2 ) aj

n
X
k=1

(2) 2

bk

e −3/2 ),
+ O(n

m m
X
X
(2) 2
(2)
(2) 2
(2) 2
2 (2)
2
2
e −3/2 ), aj
Z•k = −r bk
− r (1 − r ) bk aj
+ O(n j=1

(2)
= −r 2 (1 − r 2 )
Z••

j=1

m
X
j=1

(2) 2

aj

n
X
k=1

which gives

(2) 2

bk

e −1/2 ),
+ O(n

(sj − s)C
sj − s
(1 − 2λ)(sj − s)2 C
+ 2
+
λn
λ (1 − λ)m2 n2
λ3 (1 − λ)2 m2 n3
(1 − 2λ)RC
e −5/2 )
(1 ≤ j ≤ m),
− 3
2 3 3 + O(n
2λ (1 − λ) m n
(t − t)R
t −t
(1 − 2λ)(tk − t)2 R
(3)
+ 2 k bk = k
+
λm
λ (1 − λ)m2 n2
λ3 (1 − λ)2 m3 n2
(1 − 2λ)RC
e −5/2 )
(1 ≤ k ≤ n).
− 3
2 3 3 + O(n
2λ (1 − λ) m n
(3)

aj =

(11)

(3)
e −5/2 )
Further iterations make no change to this accuracy, so we have that aj = aj +O(n

8

(3)

e −5/2 ). We also have that and bk = bk + O(n
Zjk =

(sj − s)(tk − t)2

(1 − 2λ)(sj − s)(tk − t)

(sj − s)2 (tk − t)

− 2
− 2
λ2 (1 − λ)mn
λ (1 − λ)m2 n
λ (1 − λ)mn2
(1 − 2λ)(sj − s)2 (tk − t)2 (1 − 2λ)(sj − s)(tk − t)R
−
+
λ3 (1 − λ)2 m2 n2
λ3 (1 − λ)2 m2 n3
(1 − 2λ)(sj − s)(tk − t)C
e −5/2 ).
+
+ O(n
λ3 (1 − λ)2 m3 n2

(12)

A sufficient approximation of λjk is given by substituting (11) and (12) into (5).
In evaluating the integral I(s, t), the following approximations will be required:
(1 − 2λ)(sj − s) (1 − 2λ)(tk − t) (sj − s)2
+
−
n m
n2
2
(t − t)2 (1 − 6λ + 6λ )(sj − s)(tk − t) e −3/2
− k 2 +
+ O(n
),
λ(1 − λ)mn m

λjk (1 − λjk ) = λ(1 − λ) +

(1 − 6λ + 6λ2 )(sj − s)
λjk (1 − λjk )(1 − 2λjk ) = λ(1 − λ)(1 − 2λ) +
n
(1 − 6λ + 6λ2 )(tk − t)
e −1 ),
+
+ O(n m
e −1/2 ).
λjk (1 − λjk )(1 − 6λjk + 6λ2jk ) = λ(1 − λ)(1 − 6λ + 6λ2 ) + O(n

3.1

Estimating the factor P (s, t)

Let
Λ=

m Y
n
Y

j=1 k=1

Then
Λ

−1

=
=

λ

λjkjk (1 − λjk )1−λjk .


m Y
n 
Y
1 + qj rk λjk j=1 k=1
m Y
n
Y

qj rk

(1 + qj rk )

(1 + qj rk )1−λjk

Y
m

λ
qj j•

−1
n
Y
λ•k rk

j=1
k=1
m n
Y −s Y −t
=
(1 + qj rk )
qj j rk k j=1 k=1
j=1
k=1
j=1 k=1
m Y
n
Y

using (2). Therefore the factor P (s, t) in front of the integral in (1) is given by
P (s, t) = (2π)−(m+n) Λ−1 .
9

(13)

(14)

(15)

We proceed to estimate Λ. Writing λjk = λ(1 + xjk ), we have log

 λjk

λjk (1 − λjk )1−λjk
λλ (1 − λ)1−λ



λ
= λxjk log
1−λ





x5jk
λ
λ(1 − 2λ) 3
λ(1 − 3λ + 3λ2 ) 4
2
+
x −
x +
xjk + O
.
2(1 − λ) jk 6(1 − λ)2 jk
12(1 − λ)3
(1 − λ)4

(16)

We know from (2) that λ•• = mnλ, which implies that x•• = 0, hence the first term on the right side of (16) does not contribute to Λ. Now using (5) we can write xjk = aj + bk + Zjk and apply the estimates in (11) and (12) to obtain


C
RC
(1 − 2λ)R3 (1 − 2λ)C3
R
λ
1−λ mn
+
+
−
Λ = λ (1 − λ)
exp
2 2 2 −
4An 4Am 8A m n
24A2 n2
24A2 m2
 (17)
(1 − 3λ + 3λ2 )R4 (1 − 3λ + 3λ2 )C4
−1/2
e
+
+ O(n
) ,
+
96A3 n3
96A3 m3
P
Pn
ℓ
ℓ
where Rℓ = m j=1 (sj − s) and Cℓ =
k=1 (t − tk ) for any ℓ. Note that R2 = R and
C2 = C.
To match the formula from the sparse case solved in [2], we will write (17) in terms of binomial coefficients. First, by Stirling’s expansion of the logarithm of the gamma function, we have that


(xx+d (1 − x)1−x−d )−N
N
√
=
(x + d)N
2 πXN

1 − 2X d2 N
(1 − 2x)d (1 − 4X)d2
(18)
× exp −
−
−
+
24XN
4X
4X
16X 2

 d5 N
1 
(1 − 6X)d4 N
d
(1 − 2x)d3 N
−
+O
+ 2 + 3 3
+
24X 2
96X 3
X4
X N
X N
as N → ∞, provided x = x(N), X = X(N) = 21 x(1 − x) and d = d(N) are such that
0 < x < 1, 0 < x + d < 1 and provided that the error term in the above is o(1). From this we infer that
−1 Y

n  
m  Y
(λλ (1 − λ)1−λ )−mn m
n mn
=
sj k=1 tk
λmn
(4πA)(m+n−1)/2 m(n−1)/2 n(m−1)/2
j=1

C 
R
C
1 − 2A  m n  1 − 4A  R
(19)
+
× exp −
+
−
−
+
4An 4Am
24A
n m
16A2 n2 m2

1 − 2λ  R3
C3  1 − 6A  R4
C4  e −1/2
+
+ 2 −
+ 3 + O(n
) .
24A2 n2
m
96A3 n3
m
10

Putting (17) and (19) together, we find that
P (s, t) = Λ−1 (2π)−(m+n)

−1 Y
n  
m  Y
m n
=
(20)
tk sj
2π (m+n+1)/2
j=1
k=1


1 − 2A  m
RC
n
1 − 4A  R
C  e −1/2
× exp
+ O(n
) .
−
+
−
+
24A
n m
8A2 m2 n2
16A2 n2 m2
A(m+n−1)/2 m(n−1)/2 n(m−1)/2

4



mn
λmn

Evaluating the integral

Our next task is to evaluate the integral I(s, t) given by

Z π Qm Qn i(θj +φk )
− 1)
k=1 (1 + λj,k (e j=1
P
Pn
I(s, t) =
···
dθdφ.
exp(i m
−π
−π
j=1 sj θj + i k=1 tk φk )
Z π

(21)

It is convenient to think of θj , φk as points on the unit circle. We wish to define
“averages” of the angles θj , φk . To do this cleanly we make the following definitions, as in [1]. Let C be the ring of real numbers modulo 2π, which we can interpret as points on a circle in the usual way. Let z be the canonical mapping from C to the real interval
(−π, π]. An open half-circle is Ct = (t − π/2, t + π/2) ⊆ C for some t. Now define
[ N
Ĉ N =
Ĉt = { x = (x1 , . . . , xN ) ∈ C N | x1 , . . . , xN ∈ Ct for some t ∈ R }.
t

If x = (x1 , . . . , xN ) ∈ C0N then define x̄ = z

−1




N
1 X
z(xj ) .
N j=1

More generally, if x ∈ CtN then define x̄ = t + (x1 − t, . . . , xN − t). The function x → x̄
is well-defined and continuous for x ∈ Ĉ N .
Let R denote the set of vector pairs (θ, φ) ∈ Ĉ m × Ĉ n such that
|θ̄ + φ̄| ≤ (mn)−1/2+2ε ,
|θ̂j | ≤ n−1/2+ε

|φ̂k | ≤ m−1/2+ε

(1 ≤ j ≤ m),

(22)

(1 ≤ k ≤ n),

where θ̂j = θj − θ̄ and φ̂k = φk − φ̄. In this definition, values are considered in C. The constant ε is the sufficiently-small value required by Theorem 1.
11

Let IR′′ (s, t) denote the integral I(s, t) restricted to any region R′′ . In this section, we estimate IR′ (s, t) in a certain region R′ ⊇ R. In Section 5 we will show that the remaining parts of I(s, t) are negligible. We begin by analysing the integrand in R, but for future use when we expand the region to R′ (to be defined in (33)), note that all the approximations we establish for the integrand in R also hold in the superset of R′
defined by
|θ̄ + φ̄| ≤ 3(mn)−1/2+2ε ,
|θ̂j | ≤ 3n−1/2+ε

(1 ≤ j ≤ m − 1),

|φ̂k | ≤ 3m−1/2+ε

(1 ≤ k ≤ n − 1),

|θ̂m | ≤ 2n−1/2+3ε ,
|φ̂n | ≤ 2m−1/2+3ε .

(23)

Define θ̂ = (θ̂1 , . . . , θ̂m−1 ) and φ̂ = (φ̂1 , . . . , φ̂n−1). Let T1 be the transformation
T1 (θ̂, φ̂, ν, δ) = (θ, φ) defined by
ν = θ̄ + φ̄,

δ = θ̄ − φ̄,

together with θ̂j = θj − θ̄ (1 ≤ j ≤ m − 1) and φ̂k = φk − φ̄ (1 ≤ k ≤ n − 1). We also define the 1-many transformation T1∗ by
[
T1 (θ̂, φ̂, ν, δ).
T1∗ (θ̂, φ̂, ν) =
δ

After applying the transformation T1 to IR (s, t), the new integrand is easily seen to be independent of δ, so we can multiply by the range of δ and remove it as an independent variable. Therefore, we can continue with an (m + n − 1)-dimensional integral over S such that R = T1∗ (S). More generally, if S ′′ ⊆ (− 12 π, 21 π)m+n−2 × (−2π, 2π] and R′′ = T1∗ (S ′′ ), we have
Z
IR′′ (s, t) = 2πmn
G(θ̂, φ̂, ν) dθ̂dφ̂dν,
(24)


S

′′

where G(θ̂, φ̂, ν) = F T1 (θ̂, φ̂, ν, 0) with F (θ, φ) being the integrand of (21). The factor
2πmn combines the range of δ, which is 4π, and the Jacobian of T1 , which is mn/2.
Note that S is defined by the same inequalities (22) as define R. The first inequality is now |ν| ≤ (mn)−1/2+2ε and the bounds on
θ̂m = −

m−1
X
j=1

θ̂j and φ̂m = −

n−1
X

φ̂k

k=1

still apply even though these are no longer variables of integration.
Our main result in this section is the following.
12

Theorem 2. Under the conditions of Theorem 1, there is a region S ′ ⊇ S such that
Z
 π 1/2  π (m−1)/2  π (n−1)/2
G(θ̂, φ̂, ν) dθ̂dφ̂dν = (mn)−1/2
′
Amn
An
Am
S





1 1 − 2A m
1 1
n
1 R C
× exp − −
+
+
+
+
2
24A
n m
4A m n n m

C 
1 − 8A  R
−b
+ O(n ) .
+
+
16A2 n2 m2
In the region S, the integrand of (24) can be expanded as

 X
m X
n m X
n
X
(A3 + βjk )(ν + θ̂j + φ̂k )3
(A + αjk )(ν + θ̂j + φ̂k )2 − i
G(θ̂, φ̂, ν) = exp −
j=1 k=1
m X
n
X

+

j=1 k=1

j=1 k=1

m X
n

 X
5
.
(A4 + γjk )(ν + θ̂j + φ̂k ) + O A
|ν + θ̂j + φ̂k |
4

j=1 k=1

Here αjk , βjk , and γjk are defined by
1 λ (1 − λ ) = A + α , jk jk
2 jk

1 λ (1 − λ )(1 − 2λ ) = A + β , jk jk
3
jk
6 jk

(25)

1 λ (1 − λ )(1 − 6λ + 6λ2 ) = A + γ , jk jk jk
4
jk
24 jk

where
1 λ(1 − λ)(1 − 6λ + 6λ2 ).
A = 12 λ(1 − λ), A3 = 16 λ(1 − λ)(1 − 2λ), and A4 = 24

Approximations for αjk , βjk , γjk were given in (13)–(15).

4.1

Another change of variables

We now make a second change of variables (θ̂, φ̂, ν) = T2 (ζ, ξ, ν), where ζ = (ζ1 , . . . , ζm−1 )
and ξ = (ξ1 , . . . , ξn−1), whose purpose is to almost diagonalize the quadratic part of G.
The diagonalization will be completed in the next subsection. The transformation T2 is defined as follows. For 1 ≤ j ≤ m − 1 and 1 ≤ k ≤ n − 1 let
θ̂j = ζj + cπ1 , where c=−

1
1/2

m+m

φ̂k = ξk + dρ1 ,

and d = −
13

1
n + n1/2

and, for 1 ≤ h ≤ 4,
πh =

m−1
X

n−1
X
ρh =
ξkh .

ζjh ,

j=1

k=1

The Jacobian of the transformation is (mn)−1/2 . In [1], this transformation was seen to exactly diagonalize the quadratic part of the integrand in the semiregular case. In the present irregular case, the diagonalization is no longer exact but still provides useful progress.
By summing the equations θ̂j = ζj + cπ1 and φ̂k = ξk + dρ1 , we find that
1/2

π1 = m

m−1
X

θ̂j ,

j=1

n−1
X
φ̂k ,
ρ1 = n1/2
k=1

|π1 | ≤ m1/2 n−1/2+ε ,
(26)
|ρ1 | ≤ n1/2 m−1/2+ε ,

where the right sides come from the bounds on θ̂m and φ̂n . This implies that e −1 ) (1 ≤ j ≤ m − 1),
ζj = θ̂j + O(n e −1 ) (1 ≤ k ≤ n − 1).
ξk = φ̂k + O(n

(27)

The transformed region of integration is T2−1 (S), but for convenience we will expand it a little to be the region defined by the inequalities
|ζj | ≤ 32 n−1/2+ε

|ξk | ≤ 32 m−1/2+ε

(1 ≤ j ≤ m − 1),

(1 ≤ k ≤ n − 1),

|π1 | ≤ m1/2 n−1/2+ε ,

(28)

|ρ1 | ≤ n1/2 m−1/2+ε ,

|ν| ≤ (mn)−1/2+2ε .

We now consider the new integrand E1 = exp(L1 ) = G ◦ T2 . As in [1], the semiregular parts of the integrand (those not involving αjk , βjk or γjk ) transform to
− Amnν 2 − Anπ2 − Amρ2 − 3iA3 nνπ2 − 3iA3 mνρ2 + 6A4 π2 ρ2

e −1/2 ).
− iA3 nπ3 − iA3 nρ3 − 3iA3 cnπ1 π2 − 3iA3 dmρ1 ρ2 + A4 nπ4 + A4 mρ4 + O(n

(29)

To see the effect of the transformation on the irregular parts of the integrand, write e −1/2 ) and
ζm = θ̂m − cπ1 and ξn = θ̂n − dρ1 . From (26) we can see that ζm = O(n e −1/2 ). Thus we have, for all 1 ≤ j ≤ m and 1 ≤ k ≤ n, ζj + ξk = O(n e −1/2 ) and
ξn = O(n
14

e −1 ). Recalling also that αjk , βjk , γjk = O(n e −1/2 ), we have cπ1 + dρ1 + ν = O(n m X
n
X

αjk (ν + θ̂j + φ̂k )2

j=1 k=1

m X
n
X

=

j=1 k=1


e −1/2 ),
αjk (ζj + ξk )2 + 2(ζj + ξk )(ν + cπ1 + dρ1 ) + O(n

m X
n
X

βjk (ν + θ̂j + φ̂k ) =

j=1 k=1

e −1/2 ).
γjk (ν + θ̂j + φ̂k )4 = O(n

j=1 k=1
n m X
X

3

m X
n
X
j=1 k=1

e −1/2 ),
βjk (ζj + ξk )3 + O(n

Moreover, the terms on the right sides of the above that involve ζm or ξn contribute only e −1/2 ) in total, so we can drop them. Combining this with (29), we have
O(n
L1 = −Amnν 2 − Anπ2 − Amρ2 − 3iA3 nνπ2 − 3iA3 mνρ2 + 6A4 π2 ρ2

− iA3 nπ3 − iA3 nρ3 − 3iA3 cnπ1 π2 − 3iA3 dmρ1 ρ2 + A4 nπ4 + A4 mρ4
−

m−1
n−1
XX
j=1 k=1

−i

4.2

αjk (ζj + ξk )2 + 2(ζj + ξk )(ν + cπ1 + dρ1 )

m−1
n−1
XX
j=1 k=1



(30)

e −1/2 ).
βjk (ζj + ξk )3 + O(n

Completing the diagonalization

The quadratic form in E1 is the following function of the m + n − 1 variables ζ, ξ, ν:
Q = −Amnν 2 − Anπ2 − Amρ2
−

m−1
n−1
XX
j=1 k=1


αjk (ζj + ξk )2 + 2(ζj + ξk )(ν + cπ1 + dρ1 ) .

(31)

We will make a third change of variables, (ζ, ξ, ν) = T3 (σ, τ , µ), that diagonalizes this quadratic form, where σ = (σ1 , . . . , σm−1 ) and τ = (τ1 , . . . , τn−1 ). This is achieved using a slight extension of [6, Lemma 3.2].
Lemma 1. Let X and Y be square matrices of the same order, such that X −1 exists and all the eigenvalues of X −1 Y are less than 1 in absolute value. Then
(I + Y X −1 )−1/2 (X + Y ) (I + X −1 Y )−1/2 = X, where the fractional powers are defined by the binomial expansion.
15

Note that X −1 Y and Y X −1 have the same eigenvalues, so the eigenvalue condition on
X −1 Y applies equally to Y X −1 . If we also have that both X and Y are symmetric, then
X − 1 
X − 1  −1 T T X − 1  −1
2 (Y X −1 )T =
2 (X ) Y =
2 X Y
r r
r r≥0
r≥0
r≥0
so (I + Y X −1 )−1/2 is the transpose of (I + X −1 Y )−1/2 . Let V be the symmetric matrix associated with the quadratic form Q. Write V = Vd + Vnd where Vd has all off-diagonal entries equal to zero and matches V on the diagonal entries, and Vnd has all diagonal entries zero and matches V on the off-diagonal entries. We will apply Lemma 1 with X = Vd and Y = Vnd . Note that Vd is invertible and that both Vd and Vnd are symmetric. Let
T3 be the transformation given by T3 (σ, τ , µ)T = (ζ, ξ, ν)T = (I + Vd−1 Vnd )−1/2 (σ, τ , µ)T .
If the eigenvalue condition of Lemma 1 is satisfied then this transformation diagonalizes the quadratic form Q, keeping the diagonal entries unchanged.
From the formula for Q we extract the following coefficients, which tell us the diagonal and off-diagonal entries of V :
[ζj2] Q = −An − (1 + 2c)αj∗ ,

[ξk2 ] Q = −Am − (1 + 2d)α∗k ,

[ν 2 ] Q = −Amn,

[ζj1 ζj2 ] Q = −2c(αj1 ∗ + αj2 ∗ )

(j1 6= j2 ),

[ξk1 ξk2 ] Q = −2d(α∗k1 + α∗k2 )

(k1 6= k2 ),

[ζj ξk ] Q = −2αjk − 2dαj∗ − 2cα∗k ,
[ζj ν] Q = −2αj∗ ,

[ξk ν] Q = −2α∗k .
e −3/2 ), except
Using these equations we find that all off-diagonal entries of Vd−1 Vnd are O(n e −1/2 ). Simifor the column corresponding to ν which has off-diagonal entries of size O(n e −3/2 ), except for the row corresponding larly, the off-diagonal entries of Vnd Vd−1 are all O(n e −1/2 ). To see that these conditions imply to ν, which has off-diagonal entries of size O(n that the eigenvalues of Vd−1 Vnd are less than one, recall that the value of any matrix norm is greater than or equal to the greatest absolute value of an eigenvalue. The ∞-norm e −1/2 ), so the eigenvalues are
(maximum row sum of absolute values) of Vd−1 Vnd is O(n e −1/2 ).
all O(n
We also need to know the Jacobian of the transformation T3 .

e −1/2 ).
Lemma 2. Let M be a matrix of order O(m+n) with all eigenvalues uniformly O(n
Then

e −1/2 ) .
det(I + M) = exp tr M − 21 tr M 2 + O(n
16

Proof. The eigenvalue condition ensures that the Taylor series for log(I + M) converges and that

det(I + M) = exp tr log(I + M) .

e −(r−2)/2 ) for r ≥ 3 gives the
Expanding the logarithm and noting that |tr M r | = O(n result.
e −1/2 ) so Lemma 2
Let M = Vd−1 Vnd . As noted before, the eigenvalues of M are all O(n e −1 ), we conclude that applies. Noting that tr(M) = 0 and calculating that tr(M 2 ) = O(n the Jacobian of T3 is

−1/2
e −1/2 ).
det (I + M)−1/2 = det(I + M)
= 1 + O(n

To derive T3 explicitly, we can expand (I + Vd−1 Vnd )−1/2 while noting that αj∗ =
O(n1/2+ε ) for all j, α∗k = O(m1/2+ε ) for all k, α∗∗ = O(mn2ε + nm2ε ), R ≤ mn1+2ε and
C ≤ nm1+2ε .
This gives

m−1
X

n−1 


X
c(αj∗ + αj ′∗ )
αjk + dαj∗ + cα∗k
−2
−2
e e
′
σj = ζj +
+ O(n ) ζj +
+ O(n ) ξk
2An
2An
′
k=1
j =1
α

j∗
−1
e e −2 ),
+
+ O(n ) ν + O(n
2An n−1 
m−1

X
X  αjk + dαj∗ + cα∗k d(α∗k + α∗k′ ) e −2 
−2
e
+ O(n ) ζj +
+ O(n ) ξk′
τk = ξk +
2Am
2Am
′
j=1
k =1
 α

∗k e −1 ) ν + O(n e −2 ),
+
+ O(n
2Am n−1 
m−1


X
X  αj∗
α∗k e −2 ) ζj +
e −2 ) ξk + O(n e −1 )ν,
+ O(n
+ O(n
µ=ν+
2Amn
2Amn j=1
k=1

for 1 ≤ j ≤ m − 1, 1 ≤ k ≤ n − 1.

The transformation T3−1 perturbs the region of integration in an irregular fashion that we must bound. From the explicit form of T3 above, we have
σj = ζj +

m−1
X
′

j =1

τk = ξk +

m−1
X
j=1

e −3/2 )ζ ′ +
O(n j
e −3/2 )ζj +
O(n

n−1
X
k=1

n−1
X
′

k =1

e −3/2 )ξk + O(n e −1/2 )ν + O(n e −2 ) = ζj + O(n e −1 ),
O(n e −3/2 )ξ ′ + O(n e −1/2 )ν + O(n e −2 ) = ξk + O(n e −1 )
O(n k

for 1 ≤ j ≤ m − 1, 1 ≤ k ≤ n − 1, so σ, τ are only slightly different from ζ, ξ.
17

For µ versus ν we have
µ = ν + O(n−1+2ε /A) + O(m−1+2ε /A)

= ν + o (mn)−1/2+2ε ,

where the second step requires our assumptions m = o(A2 n1+ε ) and n = o(A2 m1+ε ). This shows that the bound |ν| ≤ (mn)−1/2+2ε is adequately covered by |µ| ≤ 2(mn)−1/2+2ε .
For 1 ≤ h ≤ 4, define

µh =

m−1
X

h

σj ,

νh =

j=1

n−1
X

τk h .

k=1

From (28), we see that |π1 | ≤ m1/2 n−1/2+ε and |ρ1 | ≤ m−1/2+ε n1/2 are the remaining constraints that define the region of integration. We next apply these constraints to bound µ1 and ν1 . From the explicit form of T3 , we have
µ1 = π1 +

m−1
X m−1
X

c(αj∗ + αj ′∗ )
2An

j=1 j ′ =1

+

m−1
n−1 
XX
j=1 k=1


e −2 ) ζ ′
+ O(n j

m−1


X  αj∗
αjk + dαj∗ + cα∗k
−2
e e −1 ) ν + O(n e −1 )
+ O(n ) ξk +
+ O(n
2An
2An j=1

cα
dα
α
= π1 + ∗∗ m1/2 n−1/2+ε + ∗∗ m−1/2+ε n1/2 + ∗∗ ν
2An
2An
2An m−1
n−1
 X α∗k c(m − 1) X
e −1/2 )
αj ′ ∗ ζj ′ + O(n
ξk +
+ 1 + c(m − 1)
2An
2An
′
k=1
j =1

= π1 +

m−1
X

c(m − 1)
2An
′

j =1

−1

e −1/2 )
αj ′ ∗ ζj ′ + O(n

(32)

= π1 + O(A mn−1+2ε )

= π1 + o(m1/2 n−1/2+5ε/2 ).
To derive the above we have used 1 + c(m−1) = m1/2 and the bounds we have established on the various variables. For the last step, we need the assumption m = o(A2 n1+ε ), which implies that A−1 mn−1+2ε = o(m1/2 n−1/2+5ε/2 ).
Since our region of integration has |π1 | ≤ m1/2 n−1/2+ε , we see that this implies the bound |µ1 | ≤ m1/2 n−1/2+3ε . By a parallel argument, we have
ν1 = ρ1 + o(m−1/2+5ε/2 n1/2 ), which implies |ν1 | ≤ n1/2 m−1/2+3ε . Putting together all the bounds we have derived, we see that
T3−1 (T2−1 (S)) ⊆ Q ∩ M,
18

where
Q = { |σj | ≤ 2n−1/2+ε , j = 1, . . . , m − 1 } ∩ { |τk | ≤ 2m−1/2+ε , k = 1, . . . , n − 1 }
∩ {|µ| ≤ 2(mn)−1/2+2ε },

M = { |µ1| ≤ m1/2 n−1/2+3ε } ∩ { |ν1 | ≤ n1/2 m−1/2+3ε }.
Now define
S ′ = T2 (T3 (Q ∩ M)),

(33)

R′ = T1∗ (S ′ ).

We have proved that S ′ ⊇ S, so it is valid to take S ′ to be the region required by
Theorem 2. Also notice that R′ is contained in the region defined by the inequalities (23).
As we forecast at that time, our estimates of the integrand have been valid inside this expanded region. It remains to apply the transformation T3−1 to the integrand (30) so that we have it in terms of (σ, τ , µ). The explicit form of T3−1 is similar to the explicit form for T3 , namely: n−1 


X
αjk + dαj∗ + cα∗k e −2 ) σ ′ −
e −2 ) τk
+ O(n
+
O(n j
2An
2An
′
k=1
j =1
α

j∗
e −1 ) µ + O(n e −2 ),
−
+ O(n
2An m−1
n−1 



X αjk − dαj∗ + cα∗k
X
d(α∗k + α∗k′ )
e −2 ) σj −
e −2 ) τ ′
ξk = τk −
+ O(n
+ O(n k
2Am
2Am
′
j=1
k =1
 α

∗k
−1
−2
e e
−
+ O(n ) µ + O(n ),
2Am m−1
n−1 


X  αj∗
X
α∗k e −2 ) σj −
e −2 ) τk + O(n e −1 )µ,
ν =µ−
+ O(n
+ O(n
2Amn
2Amn j=1
k=1

ζj = σj −

m−1
X

c(αj∗ + αj ′ ∗ )

for 1 ≤ j ≤ m−1, 1 ≤ k ≤ n−1. In addition to the relationships between the old and new e −1/2 ), ρ2 = ν2 + O(n e −1/2 ), variables that we proved before, we can note that π2 = µ2 + O(n e −1 ), ρ3 = ν3 + O(n e −1 ), π4 = µ4 + O(n e −3/2 ), and ρ4 = ν4 + O(n e −3/2 ).
π3 = µ3 + O(n

The quadratic part of L1 , which we called Q in (31), loses its off-diagonal parts according to our design of T3 . Thus, what remains is
−Amnµ2 −

m−1
X
j=1

n−1
X


An + (1 + 2c)αj∗ σj2 −
Am + (1 + 2d)α∗k τk2
k=1

2

= −Amnµ − Anµ2 − Amν2 −
19

m−1
X
j=1

αj∗ σj2 −

n−1
X
k=1

e −1/2 ).
α∗k τk2 + O(n

Next consider the cubic terms of L1 . These are
− 3iA3 nνπ2 − 3iA3 mνρ2 − iA3 nπ3 − iA3 nρ3
− 3iA3 cnπ1 π2 − 3iA3 dnρ1 ρ2 − i

m−1
n−1
XX

βjk (ζj + ξk )3 .

j=1 k=1

We calculate the following in Q ∩ M:

 m−1

n−1
X
3iA3 µ2 X
e −1/2 ),
−3iA3 nνπ2 = −3iA3 nµµ2 +
αj∗ σj +
α∗k τk + O(n
2Am j=1
k=1
 m−1
3iA3 X
c(αj∗ + αj ′ ∗ )σj2 σj ′ ,
−iA3 nπ3 = −iA3 nµ3 +
2A
′
j,j =1

+

m−1
n−1
XX

(αjk + dαj∗ + cα∗k )σj2 τk

j=1 k=1



m−1

3iA3 c mµ2 X
e −1/2 ),
−3iA3 cnπ1 π2 = −3iA3 cnµ1 µ2 +
αj∗σj + O(n
2A
j=1
2

−i

n−1
m−1
XX
j=1 k=1

3

βjk (ζj + ξk ) = −i

n−1
m−1
XX
j=1 k=1

e −1/2 ),
+ O(n

e −1/2 ),
βjk (σj + τk )3 + O(n

(34)

(35)

and the remaining cubic terms are each parallel to one of those. The proof of (34) is similar to the proof of (32).
Finally we come to the quartic part of E1 , which is e −1/2 ).
6A4 π2 ρ2 + A4 nπ4 + A4 mρ4 = 6A4 µ2 ν2 + A4 nµ4 + A4 mν4 + O(n


e −1/2 ) ,
In summary, the value of the integrand for (σ, τ , µ) ∈ Q ∩ M is exp L2 + O(n where
2

L2 = −Amnµ − Anµ2 − Amν2 −

m−1
X

αj∗ σj2 −

m−1
X

βj∗ σj3 − i

j=1

n−1
X

α∗k τk2 + 6A4 µ2 ν2

k=1

+ A4 nµ4 + A4 mν4 − iA3 nµ3 − iA3 mν3 − 3iA3 cnµ1 µ2 − 3iA3 dmν1 ν2
− 3iA3 nµµ2 − 3iA3 mµν2 − i
+i

m−1
X
′

j,j =1

gjj ′ σj σj2′ + i

n−1
X

j=1

hkk′ τk τk2′ + i

n−1
X
k=1

m−1
n−1
XX
j=1 k=1

′

k,k =1

20

β∗k τk3

ujk σj τk2 + vjk σj2 τk ,

(36)

with

3A3
(1 + cm + c2 m2 )αj∗ + cmαj ′ ∗ = O(n−1/2+ε ),
2Am

3A3
(1 + dn + d2 n2 )α∗k + dnα∗k′ = O(m−1/2+ε ), hkk′ =
2An

3A3
nαjk + (1 + dn)αj∗ + cnα∗k − 3βjk = O(m−1/2+2ε + n−1/2+2ε ), ujk =
2An

3A3
mαjk + (1 + cm)α∗k + dmαj∗ − 3βjk = O(m−1/2+2ε + n−1/2+2ε ).
vjk =
2Am gjj ′ =

Note that the O( ) estimates in the last four lines are uniform over j, j ′ , k, k ′ .

4.3

Estimating the main part of the integral

Define E2 = exp(L2 ). We have shown that the value of the integrand in Q ∩ M is

e −1/2 ) . Denote the complement of the region M by Mc . We can
E1 = E2 1 + O(n approximate our integral as follows:
Z
Z
Z
−1/2
e
E1 =
E2 + O(n
)
|E2 |
Q∩M
Q∩M
Q∩M
Z
Z
−1/2
e
=
E2 + O(n
)
|E2 |
Q∩M
Q
Z
Z
Z
−1/2
e
=
E2 + O(1)
|E2 | + O(n
)
|E2 |.
(37)
Q

Q∩M

c

Q

It suffices to estimate the value of each integral in (37).

We first compute the integral of E2 over Q. We proceed in three stages, starting with integration with respect to µ. For the latter, we can use the formula
−1/2+2ε

(mn)Z



 π 1/2

β2
−1
exp −Amnµ − iβµ dµ =
+ O(n ) , exp −
Amn
4Amn
2

−1/2+2ε

−(mn)

provided β = o(A(mn)1/2+2ε ). In our case, β = 3A3 (nµ2 + mν2 ), which is small enough because of the assumptions m = o(A2 n1+ε ) and n = o(A2 m1+ε ). Therefore, integration over µ contributes


 π 1/2
−9A23 (nµ2 + mν2 )2
−1
+ O(n ) .
(38)
exp
Amn
4Amn

21

The second step is to integrate with respect to σ the integrand


exp − Anµ2 −
−i

m−1
X

m−1
X
j=1

9A23 n 2
2
µ2 − iA3 nµ3 − 3iA3 cnµ1 µ2
αj∗ σj −
4Am

βj∗ σj3 + i

j=1

m−1
X

gjj ′ σj σj2′ + i

j,j =1

9A23 

+ 6A4 −

2A

ujk σj τk2 + vjk σj2 τk

j=1 k=1

′



m−1
n−1
XX


µ2 ν2 + A4 nµ4 + O(n ) .



(39)

−1

This is accomplished by an appeal to Theorem 4, presented in the Appendix. In the terminology of that theorem, we have N = m − 1, δ(N) = O(n−1 ), ε′ = 32 ε, ε′′ = 53 ε,
ε′′′ = 3ε, ε̄ = 6ε, and ε̂(N) = ε + o(1) is defined by 2n−1/2+ε = N −1/2+ε̂ . Furthermore, n−1
X
9A23 
ν +i vjk τk , âj = −αj∗ + 6A4 −
2A 2
k=1



An
Â =
, m−1

iA3 n i
−
β , m − 1 m − 1 j∗
A4 n
Êj =
, m−1
n−1
X
ujk τk2 .
Jˆj = i

B̂j = −

Ĉjj ′ = −3iA3 cn + igjj ′ ,
9A23 n
F̂jj ′ = −
,
4Am

k=1

We can take ∆ = 43 , and calculate that
3

N
X

4Â2 N j=1
15

N
X

16Â3 N j=1

B̂j2 +

Êj +
3

8Â3 N 2

1
4Â2 N 2
N
X
′

j,j =1

N
X
′

j,j =1

B̂j Ĉjj ′ +



9A23
m 3A4
e −1 ),
+ O(n
F̂jj ′ =
2 −
3
n 4A
16A
1

16Â3 N 3

N
X
′ ′′

j,j ,j =1

3A23 m e −1 ),
Ĉjj ′ Ĉjj ′′ = − 3 + O(n
8A n

m−1
N
X
1 X
1
2
(αj∗ )2
âj +
α∗∗ +
âj = −
2 2
2 2
2An
4A n j=1
2ÂN j=1
4Â N j=1
(40)


n−1
i X
m 3A4 9A23
e −1/2 ),
−
v∗k τk + O(n
+
2 ν2 +
n
A
2An
4A
k=1
 2



3A3 m
(1 − 2λ)2 m
−1
e
Ẑ = Z1 = exp
.
+ O(n ) = O(1) exp
24An
8A3 n

1

N
X

1

e −1/2 ), and so integration with respect to σ
Applying Theorem 4, we see that Θ2 = O(n
22

contributes a τ -free factor
 
 π (m−1)/2
1
m 3A4 15A23 
exp
α
2 −
3 −
An n 4A
2An ∗∗
16A

m−1
1 X
2
−1/2
−3/4
e
+
(αj∗) + O(n
) + O(n
Z1 ) .
4A2 n2 j=1

(41)

e −1/2 ) + O(n−3/4 Z1 ) = O(n e −1/2 ) = o(1)
By the conditions of Theorem 1, Z1 ≤ n1/5 , so O(n as required by Theorem 4.

Finally, we need to integrate over τ . Collecting the remaining terms from (36), and the terms involving τ from (38) and (40), we have an integrand equal to

 3A m 9A2 m 
9A23 m 2
4
3
ν
−
exp −Amν2 +
−
ν + A4 mν4 − iA3 mν3 − 3iA3 dmν2 ν1
2
An
4An 2
4A2 n

n−1
n−1
n−1
n−1
X
X
X
i X
2
−1/2
2
3
e v∗k τk + i hkk′ τk τk′ + O(n
−
α∗k τk − i
β∗k τk +
) .
2An
′
k=1
k=1
k=1
k,k =1

e −1/2 ), ε′ = 3 ε, ε′′ = 5 ε,
In the terminology of Theorem 4, N = n − 1, δ(N) = O(n
2
3
ε′′′ = 3ε, ε̄ = 4ε, and ε̂(N) = ε + o(1) is defined by 2m−1/2+ε = N −1/2+ε̂ . Furthermore,

Am
3A4 m 9A23 m
− α∗k ,
Â =
, âk =
−
n−1
An
4A2 n iA m i
B̂k = − 3 −
β ,
Ĉkk′ = −3iA3 dm + ihkk′ , n − 1 n − 1 ∗k
Am
9A2 m
Êk = 4 ,
F̂kk′ = − 3 , n−1
4An i
v .
Jˆk =
2An ∗k
We can take ∆ = 43 again and calculate that


N
N
X
1
9A23
3 X
n 3A4
e −1 ),
Êk +
+ O(n
F̂kk′ =
2 −
3
2
2 2
m 4A
16A
4Â N k=1
4Â N
′
k,k =1

15

N
X

16Â3 N k=1

B̂k2 +
1

3

8Â3 N 2
N
X

2ÂN k=1

N
X

B̂j Ĉkk +
′

′

k,k =1

âk +

1

N
X

4Â2 N 2 k=1

1

16Â3 N 3

â2k = −

N
X
′

′′

k,k ,k =1

Ĉkk Ĉkk
′

′′

3A23 n e −1 ),
= − 3 + O(n
8A m

n−1

1
1 X
(α∗k )2
α∗∗ +
2 2
2Am
4A m k=1

9A23 3A4
e −1/2 ),
−
3 +
2 + O(n
2A 
8A
 2
2 
(1
−
2λ)
n
3A3 n
−1
e
.
+ O(n
) = O(1) exp
Ẑ = Z2 = exp
3
24Am
8A m
23

e −1/2 ). Including the contributions from (38) and (41),
We again find that Θ2 = O(n we obtain
Z
 π 1/2  π (m−1)/2  π (n−1)/2
E2 =
Amn
An
Am
Q

2
9A
n  3A4 15A23 
3A4  m
× exp − 33 +
+
+
−
n m 4A2 16A3
8A
2A2
n−1
 1
(42)
1 X
1 
2
−
α∗∗ +
+
(α
)
∗k
2Am 2An
4A2 m2 k=1

m−1
1 X
2
−1/2
e
(αj∗) + O(n
)Z2 .
+
4A2 n2 j=1
Using (13) and the conditions of Theorem 1, we calculate that
α∗∗ = −
m−1
X
j=1

n−1
X
k=1

1  R C  e 1/2
+ O(n ),
+
2 n m

e 3/2 ),
(αj∗ )2 = 41 (1 − 2λ)2 R + O(n

e 3/2 ),
(α∗k )2 = 14 (1 − 2λ)2 C + O(n

e −1/2 )Z2 = O(n e −1/2 )n2a/5 = O(n−b ).
O(n

Substituting these values into (42) together with the actual values of A, A3 , A4 , we conclude that
Z
 π 1/2  π (m−1)/2  π (n−1)/2
E2 =
Amn
An
Am
Q


1 1
n
1  R C 
1 1 − 2A m
+
+
+
+
× exp − −
(43)
2
24A
n m
4A m n n m

1 − 8A  R
C 
−b
+
+
+ O(n ) .
16A2 n2 m2
R
We next infer a estimate of Q |E2 |. The calculation that lead to (42) remains valid if we set all the values A3 , βjk , gjj ′ , hkk′ , ujk and vjk to zero, which is the same as replacing

24


L2 by its real part. Since |E2 | = exp Re(L2 ) , this gives


Z
9A23 15A23  m n
+
|E2 | = exp
+ o(1)
E2
+
m
8A3 16A3 n
Q
Q

Z
5n
(1 − 2λ)2 
5m 
= exp
1+
+ o(1)
E2
+
8A
6m
6n
Q
Z
a
= O(n ) E2

Z

(44)

Q

under the assumptions of Theorem 1. The third term of (37) can now be identified:
Z
Z
Z
−b
−1/2
−1/2
a e
e
E2 = O(n ) E2 ,
(45)
O(n
)
|E2 | = O(n
)n
Q

Q

Q

where, as always, we suppose that ε is sufficiently small.
Finally, we consider the second term of (37), namely
Z
|E2 |,
Q∩M

c

R
which we will bound as a fraction of Q |E2 | using a statistical technique. The following is a well-known result of Hoeffding [3].
Lemma 3. Let X1 , X2 , . . . , XN be independent random variables such that E Xi = 0 and
|Xi | ≤ M for all i. Then, for any t ≥ 0,
Prob

N
X
i=1



Xi ≥ t ≤ exp −


t2
.
2NM 2


Now consider |E2 | = exp Re(L2 ) . Write M = M1 ∩ M2 , where M1 = { |µ1| ≤
m1/2 n−1/2+3ε } and M2 = { |ν1 | ≤ n1/2 m−1/2+3ε }. For fixed values of µ and σ, Re(L2 ) separates over τ1 , τ2 , . . . , τn−1 and therefore, apart from normalization, it is the joint density of independent random variables X1 , X2 , . . . , Xn−1 which satisfy E Xk = 0 (by symmetry)
and |Xk | ≤ 2m−1/2+ε (by the definition of Q). By Lemma 3, the fraction of the integral over τ (for fixed µ, σ) that has ν1 ≥ n1/2 m−1/2+3ε is at most exp(−m4ε /2). By symmetry, the same bound holds for ν1 ≤ −n1/2 m−1/2+3ε . Since these bounds are independent of µ
and σ, we have
Z
Z
c
Q∩M2

|E2 | ≤ 2 exp(−m4ε /2)

Q

|E2 |.

Q

|E2 |.

By the same argument,

Z

4ε

c

Q∩M1

|E2 | ≤ 2 exp(−n /2)
25

Z

Therefore we have in total that
Z
Z

4ε
4ε
|E2 |
|E2 | ≤ 2 exp(−m /2) + exp(−n /2)
c
Q
Q∩M
Z
−b
≤ O(n ) E2 ,

(46)

Q

R
as for (45). Applying (37) with (43), (45) and (46), we find that Q∩M E1 is given by (43).
Multiplying by the Jacobians of the transformations T2 and T3 , we find that Theorem 2
is proved for S ′ given by (33).

5

Bounding the remainder of the integral

In the previous section, we estimated the value of the integral IR′ (s, t), which is the same as I(s, t) except that it is restricted to a certain region R′ ⊇ R (see (21–23)) In this section, we extend this to an estimate of I(s, t) by showing that the remainder of the region of integration contributes negligibly.
Precisely, we show the following.
Theorem 3. Let F (θ, φ) be the integrand of I(s, t) as defined in (21). Then, under the conditions of Theorem 1,
Z
Z
−1
|F (θ, φ)| dθdφ = O(n )
F (θ, φ) dθdφ.
c

′

R

R

For 1 ≤ j ≤ m, 1 ≤ k ≤ n, let Ajk = A + αjk = 21 λjk (1 − λjk ) (recall (25)), and define e −1/2 ). We begin with two technical lemmas whose proofs are
Amin = minjk Ajk = A + O(n omitted.
Lemma 4.

|F (θ, φ)| =
where fjk (z) =
Moreover, for all real z,

p

m Y
n
Y

fjk (θj + φk ),

j=1 k=1

1 − 4Ajk (1 − cos z) .

1 A z 4 .
0 ≤ fjk (z) ≤ exp −Ajk z 2 + 12
jk

Lemma 5. For all c > 0,
Z 8π/75
p

exp c(−x2 + 37 x4 ) dx ≤ π/c exp(3/c).
−8π/75

26

R
Proof of Theorem 3. Our approach will be to bound |F (θ, φ)| over a variety of regions
R
whose union covers Rc . To make the comparison of these bounds with R′ F (θ, φ) easier, we note that
Z


F (θ, φ) dθdφ = exp O(mε + nε ) I0 = exp O(m3ε + n3ε ) I1 ,
(47)
′

R

where

m 
n
 π 1/2 Y
π 1/2 Y π 1/2
I0 =
,
A••
A
A
j•
•k j=1
k=1
 π m/2  π n/2
.
I1 =
An
Am

To see this, expand




2
αj•
αj•
Aj• = An + αj• = An exp
−
+··· ,
An 2A2 n2
and similarly for A•k , and compare the result to Theorem 2 using the assumptions of
Theorem 1. It may help to recall the calculation following (42).
Take κ = π/300 and define x0 , x1 , . . . , x299 by xℓ = 2ℓκ. For any ℓ, let S1 (ℓ) be the set of (θ, φ) such that θj ∈ [xℓ − κ, xℓ + κ] for at least κm/π values of j and φk ∈
/
ε
[−xℓ − 2κ, −xℓ + 2κ] for at least n values of k. For (θ, φ) ∈ S1 (ℓ), θj + φk ∈
/ [−κ, κ] for at
ε
ε
least κmn /π pairs (j, k) so, by Lemma 4, |F (θ, φ)| ≤ exp(−c1 Amin mn ) for some c1 > 0
which is independent of ℓ.
Next define S2 (ℓ) to be the set of (θ, φ) such that θj ∈ [xℓ −κ, xℓ +κ] for at least κm/π
values of j, φk ∈ [−xℓ −2κ, −xℓ +2κ] for at least n−nε values of k and θj ∈
/ [xℓ −3κ, xℓ +3κ]
ε
for at least m values of j. By the same argument with the roles of θ and φ reversed,
|F (θ, φ)| ≤ exp(−c2 Amin mε n) for some c2 > 0 independent of ℓ when (θ, φ) ∈ S2 (ℓ).

Now define R1 (ℓ) to be the set of pairs (θ, φ) such that θj ∈ [xℓ − 3κ, xℓ + 3κ] for at least m − mε values of j, and φk ∈ [−xℓ − 3κ, −xℓ + 3κ] for at least n − nε values of k.
By the pigeonhole principle, for any θ there is some ℓ such that [xℓ − κ, xℓ + κ] contains at least κm/π values of θj . Therefore,
299
[

ℓ=0

S

299
c [

S1 (ℓ) ∪ S2 (ℓ) .
R1 (ℓ) ⊆

c

ℓ=0

Since the total volume of is at most (2m)m+n , we find that for some c3 > 0,
ℓ R1 (ℓ)
Z
|F (θ, φ)| dθdφ
S
(

c
ℓ R1 (ℓ))

≤ (2π)m+n exp(−c3 Amin mnε ) + exp(−c3 Amin mε n)

≤ e−n I1 .

27



(48)

S
We are left with (θ, φ) ∈ ℓ R1 (ℓ). If we subtract xℓ from each θj and add xℓ to each
φk the integrand F (θ, φ) is unchanged, so we can assume for convenience that ℓ = 0 and that (θ, φ) ∈ R1 = R1 (0). The bounds we obtain on parts of the integral we seek to reject will be at least 1/300 of the total and thus be of the right order of magnitude. We will not mention this point again.
For a given θ, partition {1, 2, . . . , m} into sets J0 = J0 (θ), J1 = J1 (θ) and J2 = J2 (θ), containing the indices j such that |θj | ≤ 3κ, 3κ < |θj | ≤ 15κ and |θj | > 15κ, respectively.
Similarly partition {1, 2, . . . , n} into K0 = K0 (φ), K1 = K1 (φ) and K2 = K2 (φ). The value of |F (θ, φ)| can now be bounded using fjk (θj + φk )


1

exp −Amin (θj + φk )2 + 12
Amin (θj + φk )4 if (j, k) ∈ (J0 ∪ J1 ) × (K0 ∪ K1 ),


p
≤
1 − 4Amin(1 − cos(12κ)) ≤ e−Amin /64
if (j, k) ∈ (J0 × K2 ) ∪ (J2 × K0 ),



1
otherwise.

R
Let I2 (m2 , n2 ) be the contribution to R |F (θ, φ)| of those (θ, φ) with |J2 | = m2 and
1
|K2 | = n2 . Recall that |J0 | > m − mε and |K0 | > n − nε . We have
  
n m
(2π)m2 +n2
I2 (m2 , n2 ) ≤
n2
m2
(49)

ε
ε
′
1 A (n − n )m − 1 A (m − m )n I (m , n ),
× exp − 64
min
2
2 2
2
2
64 min

where

I2′ (m2 , n2 ) =

Z 15κ Z 15κ 

X′
X′
4
1A
···
exp −Amin
(θj + φk )2 + 12
dθ ′ dφ′ ,
(θ
+
φ
)
min j
k
−15κ

−15κ

jk

jk

and the primes denote restriction to j ∈ J0 ∪ J1 and k ∈ K0 ∪ K1 . Write m′ = m − m2 and
P
P
′
′
′
n′ = n − n2 and define θ̄ = (m′ )−1 ′j θj , θ̆j = θj − θ̄ for j ∈ J0 ∪ J1 , φ̄ = (n′ )−1 ′k φk ,
′
′
′
′
′
φ̆k = φk − φ̄ for k ∈ K0 ∪ K1 , ν ′ = φ̄ + θ̄ and δ ′ = θ̄ − φ̄ . Change variables from (θ ′ , φ′ )
to {θ̆j | j ∈ J3 } ∪ {φ̆k | k ∈ K3 } ∪ {ν ′ , δ ′ }, where J3 is some subset of m′ − 1 elements of
J0 ∪ J1 and K3 is some subset of n′ − 1 elements of K0 ∪ K1 . From Section 4 we know that the Jacobian of this transformation is m′ n′ /2. The integrand of I2′ can now be bounded using
X′ 2
X′ 2
X′
φ̆k + m′ n′ ν ′2
θ̆j + m′
(θj + φk )2 = n′
j

jk

and

X′

jk

(θj + φk )4 ≤ 27n′

X′

j

θ̆j4 + 27m′

28

k

X′

k

φ̆4k + 27m′ n′ ν ′4 .

The latter follows from the inequality (x + y + z)4 ≤ 27(x4 + y 4 + z 4 ) valid for all x, y, z.
Therefore,
Z
Z
Z 30κ

X′
X′
O(1) 30κ 30κ
′
I2 (m2 , n2 ) ≤ ′ ′
g(φ̆k )
g(θ̆j ) + Aminm′
···
exp Amin n′
k j
m n −30κ −30κ
−30κ

+ Amin m′ n′ g(ν ′ ) dθ̆j∈J3 dφ̆k∈K3 dν ′ ,

where g(z) = −z 2 + 94 z 4 . Since g(z) ≤ 0 for |z| ≤ 30κ, and we only need an upper bound, we can restrict the summations in the integrand to j ∈ J3 and k ∈ K3 . The integral now separates into m′ + n′ − 1 one-dimensional integrals and Lemma 5 (by monotonicity) gives that
′

I2′ (m2 , n2 ) = O(1)

′

π (m +n )/2
′

′

(m +n −1)/2

Amin

′

′

(m′ )n /2−1 (n′ )m /2−1

Applying (47) and (49), we find that
ε

ε

m n
X
X

m2 =0 n2 =0
m2 +n2 ≥1


exp O(m′ /(Amin n′ ) + n′ /(Amin m′ )) .


I2 (m2 , n2 ) = O e−c4 Am + e−c4 An I1

(50)

for some c4 > 0.
We have now bounded contributions to the integral of |F (θ, φ)| from everywhere outside the region

X = (θ, φ) |θj |, |φk | ≤ 15κ for 1 ≤ j ≤ m, 1 ≤ k ≤ n .
By Lemma 4, we have for (θ, φ) ∈ Ĉ m+n (which includes X ) that

m X
n m X
n

 X
X
2
1
Ajk (θ̂j + φ̂k + ν) + 12
Ajk (θ̂j + φ̂k + ν)4 ,
|F (θ, φ)| ≤ exp −
j=1 k=1

j=1 k=1

where θ̂j = θj − θ̄, φ̂k = φk − φ̄ and ν = θ̄ + φ̄. As before, the integrand is independent of δ = θ̄ − φ̄ and our notation will tend to ignore δ for that reason; for our bounds it will suffice to remember that δ has a bounded range.
We proceed by exactly diagonalizing the (m+n+1)-dimensional quadratic form. Since
Pm
Pn j=1 θ̂j =
k=1 φ̂k = 0, we have m X
n
X
j=1 k=1

2

Ajk (θ̂j + φ̂k + ν) =

m
X

Aj• θ̂j2 +

j=1

+2

n
X

A•k φ̂2k + A•• ν 2

k=1

m X
n
X

αjk θ̂j φ̂k + 2ν

j=1 k=1

29

m
X
j=1

αj• θ̂j + 2ν

n
X
k=1

α•k φ̂k .

e −1/2 ), and we can correct it with the slight
This is almost diagonal, because αjk = O(n additional transformation (I + X −1 Y )−1/2 described by Lemma 1, where X is a diagonal matrix with diagonal entries Aj•, A•k and A•• . The matrix Y has zero diagonal and e −1/2 ) apart from the row and column indexed by ν, which other entries of magnitude O(n e 1/2 ). By the same argument as used in Section 4.2, all have entries of magnitude O(n e −1/2 ), so the transformation is well-defined.
eigenvalues of X −1 Y have magnitude O(n
The new variables {ϑ̂j }, {ϕ̂k } and ν̇ are related to the old by
(θ̂1 , . . . , θ̂m , φ̂1 , . . . , φ̂n , ν)T = (I + X −1 Y )−1/2 (ϑ̂1 , . . . , ϑ̂m , ϕ̂1 , . . . , ϕ̂n , ν̇)T .

We will keep the variable δ as a variable of integration but, as noted before, our notation will generally ignore it.
e −3/2 ), we have uniformly over
More explicitly, for some d1 , . . . , dm , d′1 , . . . , d′n = O(n j = 1, . . . , m, k = 1, . . . , n that
θ̂j = ϑ̂j +

m
X

φ̂k = ϕ̂k +
ν = ν̇ +

q=1
m
X

j=1
m
X

e −2 )ϑ̂q +
O(n

n
X

e −3/2 )ϕ̂k + O(n e −1/2 )ν̇,
O(n

k=1
n
X

e −3/2 )ϑ̂j +
O(n

q=1

e −2 )ϕ̂q + O(n e −1/2 )ν̇,
O(n

(51)

n
X
e −1 )ν̇.
d′k ϕ̂k + O(n dj ϑ̂j +

j=1

k=1

Note that the expressions O( ) in (51) represent values that depend on m, n, s, t but not on {ϑ̂j }, {ϕ̂k }, ν̇.

The region of integration X is (m+n)-dimensional. In place of the variables (θ, φ)
Pn−1
P
we can use (θ̂, φ̂, ν, δ) by applying the identities θ̂m = − m−1
k=1 φ̂k .
j=1 θ̂j and φ̂n = −
(Recall that θ̂ and φ̂ don’t include θ̂m and φ̂n .) The additional transformation (51)
maps the two just-mentioned identities into identities that define ϑ̂m and ϕ̂n in terms of
(ϑ̂, ϕ̂, ν̇), where ϑ̂ = (ϑ̂1 , . . . , ϑ̂m−1 ) and ϕ̂ = (ϕ̂1 , . . . , ϕ̂n−1 ). These have the form
ϑ̂m = −
ϕ̂n =

m−1
X
j=1

m−1
X
j=1

n−1
X

−1
e −1/2 )ϕ̂k + O(n e 1/2 )ν̇, e
O(n
1 + O(n ) ϑ̂j +
k=1

e
O(n

−1/2

)ϑ̂j −

n−1
X
k=1


e e 1/2 )ν̇.
1 + O(n
) ϕ̂k + O(n

(52)

−1

Therefore, we can now integrate over (ϑ̂, ϕ̂, ν̇, δ). The Jacobian of the transformation from (θ, φ) to (θ̂, φ̂, ν, δ) is mn/2, as in Section 4. The Jacobian of the transformation e −1/2 ) by Lemma 2, using
T4 (ϑ̂, ϕ̂, ν̇) = (θ̂, φ̂, ν) def

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
