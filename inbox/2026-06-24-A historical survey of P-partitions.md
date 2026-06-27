---
source: "https://arxiv.org/abs/1506.03508v1"
title: "A historical survey of P-partitions"
author: "Ira M. Gessel"
year: "2015"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1506.03508v1"
pdf: "https://arxiv.org/pdf/1506.03508v1"
captured_at: "2026-06-24T11:10:39Z"
updated_at: "2026-06-24T11:10:39Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Richard Rorty"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# A historical survey of P-partitions

- 著者: Ira M. Gessel
- 年: 2015
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1506.03508v1)
- ダウンロード: https://arxiv.org/pdf/1506.03508v1
- PDF: https://arxiv.org/pdf/1506.03508v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We give a historical survey of the theory P-partitions, starting with MacMahon's work, describing Richard Stanley's contributions and his related work, and continuing with more recent developments.

## PDF Text

arXiv:1506.03508v1 [math.CO] 10 Jun 2015

A Historical Survey of P -Partitions
Ira M. Gessel
Abstract. We give a historical survey of the theory P -partitions, starting with MacMahon’s work, describing Richard Stanley’s contributions and his related work, and continuing with more recent developments.

Dedicated to Richard Stanley on the occasion of his seventieth birthday.
1. Introduction
Richard Stanley’s 1971 Harvard Ph.D. thesis [64] studied two related topics, plane partitions and P -partitions. A short article on his work on P -partitions appeared in 1970 [63], and a detailed exposition of this thesis work appeared as an
American Mathematical Society Memoir [66] in 1972. In this paper, I will describe some of the historical background of the theory of P -partitions, sketch Stanley’s contribution to the theory, and mention some more recent developments.
The basic idea of the theory of P -partitions was discovered by P. A. MacMahon, generalized by Knuth, and independently rediscovered by Kreweras. In order to understand the different notations and approaches that these authors used, it will be helpful to start with a short exposition of Stanley’s approach.
2. Stanley’s theory of P -partitions
I will usually follow Stanley’s notation, but with some minor modifications.
Stanley has given a more recent account of the theory of P -partitions in Enumerative Combinatorics, Vol. 1 [74], Sections 1.4 and 3.15.
Let P be a finite partially ordered set (poset) with p elements and let ω be a bijection from P to [p] = {1, 2, . . . , p}, called a labeling of P . We use the symbols
 and ≺ for the partial order relation of P . Then a (P, ω)-partition is a map σ
from P to the set N of nonnegative integers satisfying the conditions
(i) If X ≺ Y then σ(X) ≥ σ(Y ).
(ii) If X ≺ Y and ω(X) > ω(Y ) then σ(X) > σ(Y ).
P
If X∈P σ(X) = n then we call σ a P -partition of n. For simplicity, we assume that the elements of P are 1, 2, . . . , p. We denote by A (P, ω) the set of all (P, ω)partitions and by A (P, ω; m) the set of all (P, ω)-partitions with largest part at most m. Thus for the labeled poset P of Figure 1, where the elements of P are identified with their labels, σ : {1, 2, 3} → N is a P -partition if and only if σ(2) > σ(1) and
σ(2) ≥ σ(3).
2010 Mathematics Subject Classification. Primary 05A15. Secondary 05A17.
1

2

IRA M. GESSEL

3

1

2

Figure 1. A poset with three elements
If ω(X) < ω(Y ) whenever X ≺ Y , then ω is called a natural labeling, and if
ω(X) > ω(Y ) whenever X ≺ Y , then ω is called a strict labeling. If ω is a natural labeling, then all of the inequalities in the definition of a P -partition are weak, and if ω is a strict labeling then all the inequalities are strict.
A linear extension of P is a total order (or chain) on P that contains P .
Every linear extension of P inherits the labeling ω of P . We denote by L (P, ω)
the set of linear extensions of P . We may identify a linear extension of P with the permutation of [p] obtained by listing the labels of its elements in increasing order. Thus for P in Figure 1, L (P, ω) consists of the two permutations 213 and
231. Then the “fundamental theorem of P -partitions” (stated somewhat differently, but equivalently, by Stanley [66, Lemma 6.1]) asserts that the set A (P, ω) of P partitions is the disjoint union of A (π, ω) over all elements π of L (P, ω).
Thus for the poset P of Figure 1, the fundamental theorem says that the set of P -partitions is the disjoint union of the solutions of σ(2) > σ(1) ≥ σ(3) and the solutions of σ(2) ≥ σ(3) > σ(1).
We sketch here Stanley’s proof. Given a P -partition σ, we can arrange its values in weakly decreasing order and thus find a permutation π of P such that
σ(π(1)) ≥ σ(π(2)) ≥ · · · ≥ σ(π(p)).
There may be many ways to do this, but we get a unique permutation if we require that the labels of equal values be arranged in increasing order; i.e., if
σ(π(i)) = σ(π(i + 1)) then ω(π(i)) < ω(π(i + 1)). But the contrapositive of this property is (since the labels are all distinct and σ(π(i)) ≥ σ(π(i + 1))) that if
ω(π(i)) > ω(π(i + 1)) then σ(π(i)) > σ(π(i + 1)). This shows that every P partition is in A (π, ω) for a unique π ∈ L (P, ω), and it is clear from the definitions that A (π, ω) ⊆ A (P, ω) for each π ∈ L (P, ω).
The fundamental theorem has enumerative consequences, since (P, ω)-partitions are easy to count when P is a total order. In particular, Stanley defines
X
Um (P, ω) =
q σ(1)+···+σ(p) ,
σ∈A (P,ω;m)

U (P, ω) = lim Um (P, ω) =

X

m→∞

q σ(1)+···+σ(p) ,

σ∈A (P,ω)

Ω(P, ω; m) = Um−1 (P, ω)|q=1 ;
Ω(P, ω; m) is called the order polynomial of (P, ω). (In section 7.1 we will write
Um (P, ω; q) and U (P, ω; q) when we need to show dependence on q.) To evaluate these sums we apply the fundamental theorem, which reduces the problem to the case in which P is a total order. If π is a permutation of [p], we define the descent set S (π) to be the set { j | π(j) > π(j + 1) }, and we denote by des(π) the number

P -PARTITIONS

3

of elements of S (π) and by maj(π) (the major index 1 of π) the sum of the elements of S (π). Then the basic result on these quantities is
P
∞
maj(π) des(π)
X
t
π∈L (P,ω) q m
(2.1)
.
Um (P, ω)t =
(1 − t)(1 − tq) · · · (1 − tq p )
m=0
Equation (2.1) is easily proved directly if P is a total order (chain), and the general case follows from the fundamental theorem. As consequences of (2.1) we have
P
maj(π)
π∈L (P,ω) q
U (p, ω) =
(2.2)
(1 − q) · · · (1 − q p )
and
(2.3)

∞
X

1+des(π)
π∈L (P,ω) t
.
(1 − t)p+1

P
Ω(P, ω; m)t

m

=

m=0

3. MacMahon
The story of P -partitions begins with Percy A. MacMahon’s work on plane partitions [36] (see also [39, Vol. 2, Section X, Chapter 3]). The problem that
MacMahon considers is that of counting plane partitions of a given shape; that is, arrangements of nonnegative integers with a given sum in a “lattice” such as
4
4
2

4
3
1

2
2

1

in which the entries are weakly decreasing in each row and column. MacMahon gives a simple example to illustrate his idea. We want to count arrays of nonnegative integers p q r s satisfying p ≥ q ≥ s and p ≥ r ≥ s, and we assign to such an array the weight xp+q+r+s . The set of solutions of these inequalities is the disjoint union of the solution sets of the inequalities
(i) p ≥ q ≥ r ≥ s and

(ii) p ≥ r > q ≥ s.

Setting r = s + A, q = s + A + B, and p = s + A + BP
+ C, where A, B, and C are arbitrary nonnegative integers, we see that the sum xp+q+r+s over solutions of p ≥ q ≥ r ≥ s is equal to
X
1
xC+2B+3A+4s =
,
(1)(2)(3)(4)
A,B,C,s≥0

n

where (n) = (1 − x ). Similarly, the generating function for solutions of p ≥ r >
q ≥ s is x2 /(1)(2)(3)(4), so the generating function for all of the arrays is
(3.1)

1 + x2
.
(1)(2)(3)(4)

MacMahon explains (but does not prove) that a similar decomposition exists for counting plane partitions of any shape, and moreover, the terms that appear in the
1Stanley used the term “index”.

4

IRA M. GESSEL

numerator have combinatorial interpretations. They correspond to what MacMahon calls lattice arrangements, which are essentially what we now call standard
Young tableaux. In the example under discussion there are two lattice arrangements,
4 3
4 2
and
.
2 1
3 1
They are the plane partitions of the shape under consideration in which the entries are 1, 2, . . . , n, where n is the number of positions in the shape. To each lattice arrangement MacMahon associates a lattice permutation: the ith entry in the lattice permutation corresponding to an arrangement is the row of the arrangement in which n + 1 − i appears, where the rows are represented by the Greek letters α, β,
. . . . So the lattice permutation associated to the first arrangement is ααββ and to the second is αβαβ. (A sequence of Greek letters is called a lattice permutation if any initial segment contains at least as many αs as βs, at least as many βs as
γs, and so on.) To each lattice permutation, MacMahon associates an inequality relating p, q, r, and s; the αs are replaced, in left-to-right order with the first-row variables p and q, and the βs are replaced with the second-row variables r and s.
A greater than or equals sign is inserted between two Greek letters that are in alphabetical order and a greater than sign is inserted between two Greek letters that are out of alphabetical order. So the lattice permutation ααββ gives the inequalities p ≥ q ≥ r ≥ s and the lattice permutation αβαβ give the inequalities p ≥ r > q ≥ s. Each lattice permutation contributes one term to the numerator of
(3.1), and the power of x in such a term is the sum of the positions of the Greek letters that are followed by a smaller Greek letter. MacMahon then describes the variation in which a restriction on part sizes is imposed. The decomposition into disjoint inequalities works exactly as in the unrestricted case, and reduces the problem to counting partitions with a given number of parts and a bound on the largest part. Most of the rest of the paper is devoted to applications of this idea to the enumeration of plane partitions. In a postscript, MacMahon considers the analogous situation in which only decreases in the rows are required, not in the columns. The enumeration of such arrays is not of much interest in itself, since the generating function for an array with p1 , p2 , . . . , pn nodes in its n rows is clearly
1
.
(1) · · · (p1 )(1) · · · (p2 ) · · · · · · (1) · · · (pn )
However the same decomposition that is used in the case of plane partitions yields interesting results about permutations.
In a follow-up paper [37] (see also [39, Vol. 2, Section IX, Chapter 3]), MacMahon elaborates on this idea. Given a sequence of elements of a totally ordered set,
MacMahon defines a major contact 2 to be a pair of consecutive entries in which the first is greater than the second, and he defines the greater index 3 to be the sum of the positions of the first elements of the major contacts. Thus the greater index of βαααγγβαγ, where the letters are ordered alphabetically, is 1 + 6 + 7 = 14.
(He similarly defines the “equal index” and “lesser index” but these do not play
2Now usually called a descent.
3MacMahon’s greater index is now usually called the major index. The term “major index”
was introduced by Foata and Schützenberger [19, 22] in reference to MacMahon’s military rank.
Curiously, MacMahon used the term “major index” for a related concept that does not seem to have been further studied.

P -PARTITIONS

5

muchP
of a role in what follows.) MacMahon’s main result in the paper is that the sum xp , where p is the greater index, over all “permutations of the assemblage i j k
α β γ · · · ” is
(1)(2) · · · (i + j + k + · · · )
(1)(2) · · · (i) · (1)(2) · · · (j) · (1)(2) · · · (k) · · ·
As in the previous paper MacMahon illustrates with an example, but does not give a formal proof, nor even an informal explanation of why the procedure works. We consider the sum of xa1 +a2 +a3 +b1 +b2 over all inequalities a1 ≥ a2 ≥ a3 , b1 ≥ b2 . We see directly that the sum is
1
.
(1)(2)(3) · (1)(2)
MacMahon breaks up these inequalities just as before into subsets corresponding to all the permutations of α3 β 2 ; for example, to the permutation αβαβα corresponds the inequalities a1 ≥ b1 > a2 ≥ b2 > a3 , where the strict inequalities correspond to the major contacts, which in this example are all of the form βα. The generating function for the solutions of these inequalities is x6
;
(1)(2)(3)(4)(5)
here 6 = 2 + 4 is the the greater index of the permutation αβαβα. Summing the contributions from all ten permutations of α3 β 2 gives
P p x
1
=
.
(1)(2)(3)(4)(5)
(1)(2)(3) · (1)(2)
In his book Combinatory Analysis [39, Vol. 2, Section IX] MacMahon discusses the analogous result when a bound is imposed on the part sizes. The sum of xa1 +···+ap over all solutions of n ≥ a1 ≥ · · · ≥ ap is (n + 1) · · · (n + p)/(1) · · · (p), and MacMahon derives in Art. 462 an important, though not well-known, formula that he writes as
(3.2)
∞
X
(n + 1) · · · (n + p1 ) · (n + 1) · · · (n + p2 ) · · · · · · (n + 1) · · · (n + pm )
gn
(1)(2) · · · (p1 ) · (1)(2) · · · (p2 ) · · · · · · (1)(2) · · · (pm )
n=0
=

1 + g PF1 +g 2 PF2 + · · · + g ν PFν
.
(1 − g)(1 − gx)(1 − gx2 ) · · · (1 − gxp1 +···+pν )

Here PFs is the generating function, by greater index, of permutations of the aspm semblage α1p1 α2p2 · · · αm with s major contacts. This result is worth restating it in more modern terminology: Let
X
Ap1 ,...,pm (t, q) =
tdes(π) q maj(π) ,
π

where the sum is over all permutations π of the multiset {1p1 , 2p2 , . . . , mpm }, and if π = a1 · · · ap , where p = p1 + · · · pm then des(π) is the number of descents of π, that is, the number of indices i for which ai > ai+1 , and maj(π) is the sum of the n−1
descents of π. Let (a; q)m be the q-rising factorial
), let
m(1 − a)(1 − aq) · · · (1 − aq n
(q)n denote (q; q)n = (1−q) · · · (1−q ) and let n denote the q-binomial coefficient
(q)m
.
(q)n (q)m−n

6

IRA M. GESSEL

Then
(3.3)

∞
X
n=0

tn



n + p1
p1


 

n + p2
n + pm
Ap1 ,...,pm (t, q)
···
=
.
p2
pm
(t; q)p+1

Several specializations of (3.3) are worth mentioning. For q = 1 the polynomials
Ap1 ,...,pm (t, 1) solve Simon Newcomb’s problem, the problem of counting permutations of a multiset by descents.4 (MacMahon had solved Simon Newcomb’s problem earlier by a different method [35]; however, he does not note here the connection with Simon Newcomb’s problem.) In the case q = 1, p1 = · · · = pm = 1, the polynomials A1m (t, 1) are the Eulerian polynomials5, satisfying
∞
X

tn (n + 1)m =

n=0

A1m (t, 1)
.
(1 − t)m+1

For p1 = · · · = pm = 1, (3.3) becomes
∞
X

tn (1 + q + · · · + q n )m =

n=0

A1m (t, q)
,
(t; q)m+1

a result often attributed to Carlitz [14] (though Carlitz stated an equivalent result much earlier [13], attributing it to John Riordan).
4. Kreweras
In 1967, Germain Kreweras [33] (see also [32] for a briefer account) used an approach similar to MacMahon’s, though stated very differently, to solve a common generalization of Simon Newcomb’s problem6 and what he calls “Young’s problem”. In Young’s problem, we are given two weakly decreasing sequences
Y = (y1 , . . . , yh ) and Y 0 = (y10 , . . . , yh0 ) with yi ≥ yi0 for each i and we ask how many “Young chains” there are from Y 0 to Y , which are sequences of partitions
(weakly decreasing sequences of integers) starting with Y 0 and ending with Y in which each partition is obtained from the previous one by increasing one part by 1.
In modern terminology, these are standard tableaux of shape Y /Y 0 ; that is, fillings of a Young diagram of shape Y with the squares of a Young diagram of shape
Y 0 removed from it, with the integers 1, 2, . . . , m (where m is the total number of squares), so that the entries are increasing in every row and column. For example, if Y 0 = (2, 1, 0) and Y = (3, 2, 2) then one of the Young chains from Y 0 to Y is
Y 0 = (2, 1, 0), (3, 1, 0), (3, 1, 1), (3, 2, 1), (3, 2, 2) = Y . This corresponds to the skew
Young tableau
1
3
2 4
in which the entry i occurs in row j if the ith step in the chain is an increase by 1
in the jth position.
4In [35] MacMahon describes Simon Newcomb’s problem as the equivalent problem of counting permutations of a multiset by consecutive pairs which are not descents, though in his book
[39, volume 1, pp. 187] he counts by descents.
5The Eulerian polynomials are often defined to be our tA m (t, 1), making the generating
1
P
n m function the nicer ∞
n=0 t n .
6Kreweras’s seems to be unaware of MacMahon’s work on Simon Newcomb’s problem and refers only to Riordan [56, 216–219] as a reference.

P -PARTITIONS

7

Kreweras defines a “return” (retour en arrière) of a Young chain to consist of three consecutive partitions U V W such that the entry augmented in passing from V to W has an index that is strictly less than which is augmented in passing from U to V . In terms of Young tableau, a return corresponds to an entry i which is in a higher row than i + 1. (In our example, 1 and 3 correspond to returns.) In MacMahon’s approach, the returns correspond to major contacts of lattice permutations. Kreweras writes θr (Y, Y 0 ) for the number of Young chains from Y 0 to Y with r returns.
He observes that Simon Newcomb’s problem is equivalent to a special case of computing θr (Y, Y 0 ); thus the number of permutations of the multiset {13 , 2, 32 }
with r descents is equal to the number of skew Young tableaux of shape

with r returns. He then gives the solution to this problem in the form
P
0 r
X
r≥0 θr (Y, Y )t
=
wr tr .
0 +1
η−η
(1 − t)
r≥0

Here η is the sum of the entries of Y , η 0 is the sum of the entries of Y 0 , and wr is the number of chains
(4.1)

Y 0 ≤ Z1 ≤ · · · ≤ Zr ≤ Y ;

in an earlier work [31], Kreweras had given the formula


yi − yj0 + r wr = det
, i−j+r i,j=1,...,h where Y = (y1 , . . . , yh ) and Y 0 = (y10 , . . . , yh0 ). In the case of Simon Newcomb’s problem, the determinant is upper triangular, and is therefore a product of binomial coefficients (as can also be seen directly).
Kreweras’s method of proof is ultimately equivalent to MacMahon’s approach, though described very differently: he associates to every
Pchain (4.1) a Young chain from Y 0 to Y in such a way that the contribution to r wr tr corresponding to a
0
given Young chain with r returns is tr /(1 − t)η−η +1 .
In a later paper [34], Kreweras studied what is in Stanley’s terminology the order polynomial of a naturally labeled poset. Although published in 1981, long after Stanley’s memoir [66], Kreweras stated that Stanley’s work was unknown to him when the paper was written.
5. Knuth
In 1970, Donald E. Knuth [30] used MacMahon’s approach to study solid
(i.e., three-dimensional) partitions. MacMahon had conjectured that the generi+1
Q∞
ating function for solid partitions was i=1 (1 − z i )−( 2 ) . This conjecture had been disproved earlier [2], but Knuth wanted to compute the number c(n) of solid partitions of n for larger values of n in an (unsuccessful) attempt to find patterns.
Knuth realized that MacMahon’s approach would work for arbitrary partially ordered sets, not just those corresponding to plane partitions. Knuth takes a set P
partially ordered by the relation ≺ and well-ordered by the total order <, where

8

IRA M. GESSEL

x ≺ y implies x < y. He defines a P -partition of N to be a function n from P to the set of nonnegative integers satisfying
P(i) x ≺ y implies n(x) ≥ n(y), (ii) only finitely many x have n(x) > 0, and (iii) x∈P n(x) = N . Knuth proves that there is a bijection from P -partitions of N to pairs of sequences n1 ≥ n2 ≥ · · · ≥ nm x1 , x2 , . . . , xm where m ≥ 0, the ni are positive integers with sum N , and the xi are distinct elements of P satisfying
(S1) For 1 ≤ j ≤ m and x ∈ P , x ≺ xi implies x = xi for some i < j.
(S2) xi > xi+1 implies ni > ni+1 for 1 ≤ i < m.
Knuth is interested primarily in the case in which P is countably infinite, for which he uses a modification of this bijection to prove that if P is an infinite poset and s(n) is the number of P -partitions of n then
1 + s(1)z + s(2)z 2 + · · · = (1 + t(1)z + t(2)z 2 + · · · )/(1 − z)(1 − z 2 )(1 − z 3 ) · · ·
where t(k) is the number of linear extensions of finite order ideals of P with “index”
k; Knuth’s index is a variant of MacMahon’s greater index.
6. Thomas
Glânffrwd Thomas’s 1977 paper [78], based on his 1974 Ph.D. thesis [79] appeared after Stanley’s memoir, but it was written without knowledge of Stanley’s work (but with knowledge of MacMahon’s). Thomas’s motivation was the combinatorial definition of Schur functions. If λ is a partition, then a Young tableau of shape λ is a filling of the Young diagram of λ that is weakly increasing in rows and strictly increasing in columns. For example, if λ is the partition (4, 2, 1) then a Young tableau of shape λ is
(6.1)

4 4 1 1
2 1
1

The Schur function sλ is the sum of the weights of all Young tableaux of shape
λ, where the weight of a Young tableau is the product of xi over all of its entries i. (So the weight of the tableau (6.1) is x41 x2 x24 .) Schur functions are symmetric in the variables xi and have important applications in enumeration and in the representation theory of symmetric and general linear groups.
Thomas considers a more general situation, in which he allows as shapes (which he calls “frames”) any subset of Z × Z and he defines a numbering of a frame to be filling with positive integers that is weakly increasing in rows and strictly increasing in columns. For example,
(6.2)

1 2
4
2
4

is a numbering. To any numbering he associates an index numbering by replacing its entries in increasing order with 1, 2, . . . , m, where m is the number of entries,

P -PARTITIONS

9

and ties are broken from bottom to top and then left to right. Thus the index numbering corresponding to (6.2) is
1 3
5
2
4
Thomas calls two numberings equivalent if they have the same indexQnumbering.
He defines the monomial of a numbering of a frame to be the product xi over all the entries i of the frame. (So the sum of the monomials of all the frames of a Young diagram is a Schur function.) His goal is determine the sum of the monomials of an equivalence class of numberings.
The numberings of frames are a particular case of P -partitions corresponding to subposets of Z × Z, and except for the case of skew Schur functions, which are symmetric, one might just as well study general P -partitions with his weighting.
The interesting aspect of his work is that it seems to be the earliest appearance (after a brief mention by Stanley [66, p. 81]) of what are now called quasi-symmetric generating functions for P -partitions, which we will discuss in more detail in Section8.5.
As we have seen, the study of P -partitions leads to inequalities like j1 ≥ j2 >
j3 ≥ j4 , or equivalently (following Thomas), i1 ≤ i2 < i3 ≤ i4 .
P i1 +···+i4
While MacMahon wanted to compute x
, Thomas was interested in the more refined multivariable generating function
X
x i1 x i2 x i3 x i4 .
i1 ≤i2 <i3 ≤i4

This is a fundamental quasi-symmetric function; these form a basis for the algebra of quasi-symmetric functions, which will be discussed in Section 8.5.
Thomas applies Baxter operators to construct quasi-symmetric functions. A
Baxter operator on a commutative algebra A over a field K is a linear operator
B : A → A such that for some fixed nonzero θ ∈ K,
B(aB(b)) + B(bB(a)) = B(a)B(b) + B(θab)
for all a, b ∈ A. Now let A be the algebra of infinite sequences (a1 , a2 , . . . ) with entries in a field, with componentwise operations.
We define two maps A → A; first a map introduced by Rota and Smith [59]


r−1
X
S(a1 , a2 , . . . ) = 0, a1 , a1 + a2 , . . . , ai , . . . , i=1



Pr−1
which we write as . . . , i=1 ai , . . . , and a variant
P (a1 , a2 , . . . ) =



r
X
a1 , a1 + a2 , . . . , ai , . . .
i=1



r
X
= ..., ai , . . . .
i=1

Then S and P are Baxter operators.

10

IRA M. GESSEL

We show by an example the connection between these operators and quasisymmetric functions: Let x = (x1 , x2 , x3 , . . . ). Then


X
xS(xS(xP (x))) = . . . , xi xj xk xr , . . . .
1≤i≤j<k<r

Thus the fundamental quasi-symmetric function
X
xi xj xk xl , . . .
i≤j<k<l

is obtained by adding all the entries of xS(xS(xP (x))).
Thomas’s approach has been further developed by Rudolf Winkel [80].
7. Stanley
In this section we discuss a few highlights of Stanley’s memoir [66].
7.1. Reciprocity theorems. If ω is a labeling of a poset P of size p, we define the complementary labeling ω̄ of P by ω̄(i) = p + 1 − ω(i). Thus when we change the labeling of P from ω to ω̄, the strict and weak inequalities in the definition of a P -partition are switched. If the permutation π in L (P, ω) corresponds to the permutation π̄ in L (P, ω̄) then the descent sets S (π) and S (π̄) are complementary subsets of [p − 1], and thus by (2.1)–(2.3), Um (P, ω̄; q), U (P, ω̄; q), and Ω(P, ω̄; m)
are determined by Um (P, ω; q), U (P, ω; q), and Ω(P, ω; m). The formulas expressing these relations are surprisingly simple. We first note that if P is a chain and

 (P, ω)
.
corresponds to a permutation π with s descents then Um (P, ω; q) = q maj(π) p+m−s p
It follows that in general, Um (P, ω; q) is a polynomial in q m and Ω(P, ω; m) is a polynomial in m, so Um (P, ω; q) and Ω(P, ω; m) can be extended in a natural way to negative values of m. Moreover, U (P, ω; q) is a rational function of q, so
U (P, ω; q −1 ) is well-defined as a rational function of q. Then we have the following reciprocity formulas q p Um (P, ω̄; q) = (−1)p U−(m+2) (P, ω, q −1 ), with U−1 (P, ω) = 0
q p U (P, ω̄; q) = (−1)q U (P, ω; q −1 )
(7.1)

Ω(P, ω̄; m) = (−1)p Ω(P, ω; −m).

When P satisfies certain “chain conditions” there is an additional relation between the enumerative quantities associated with (P, ω) and (P, ω̄) [66, 18–19]. We state here one of these results [66, Proposition 19.3]: Suppose that (P, ω) is naturally labeled and that every maximal chain in P has length l. Then for all m,
Ω(P, ω; m) = (−1)p Ω(P, ω; −l − m) = Ω(P, ω̄; l + m)
and the number of permutations in L (P, ω) with s descents is equal to the number of permutations in L (P, ω) with n − l − s descents.
A nice application of the reciprocity theorem for order polynomials (7.1) is
Stanley’s result [67] that if χ(λ) is the chromatic polynomial of a graph G with p vertices then (−1)p χ(−1) is the number of acyclic orientations of G. Any proper coloring of G with the integers 1, 2, . . . , λ yields an acyclic orientation of G in which edges are directed from the lower color to the higher. The number of proper colorings of G in λ colors associated to an acyclic orientation O is Ω(PO , ω; λ)
where P is the poset associated to O and ω is a strict labeling. Then by (7.1),

P -PARTITIONS

11

(−1)p Ω(PO , ω; −1) = Ω(PO , ω̄; 1) = 1, since ω̄ is a natural labeling, and thus each acyclic orientation of G contributes exactly 1 to (−1)p χ(−1).
7.2. Disjoint unions. We may allow the labeling ω of a poset P to be an arbitrary function from P to the positive integers as long as incomparable elements of P have distinct labels. Then if (P, ω1 ) and (Q, ω2 ) are labeled posets in which the images of ω1 and of ω2 are disjoint, we can construct the disjoint union labeled poset (P + Q, ω1 + ω2 ) where the labeling ω1 + ω2 is ω1 on P and is ω2 on Q. It is clear that, with the notation of Section 2,
Um (P + Q, ω1 + ω2 ) = Um (P, ω1 )Um (Q, ω2 )
and similarly for a disjoint union of more than two posets. Moreover, there is a simple description of L (P + Q, ω1 + ω2 ): it is the set of “shuffles” of L (P, ω1 ) and
L (Q, ω2 ). Thus to obtain MacMahon’s formula (3.3) from (2.1) we take (P, ω) =
(P1 + · · · + Pr , ω1 + · · · + ω
Pi is a chain of size pi with every label equal
 r ) where

i to i, so that Um (Pi , ωi ) = m+p
.
pi
Now let
X
Ws (P, ω) =
q maj(π) ,
π

where the sum is over all permutations π ∈ L (P, ω) with s descents. Stanley [66,
Prop. 12.6] proves the formula
Ws (P + Q, ω1 + ω2 )
|P |−1 |Q|−1

=

X X
i=0

j=0

q (s−i)(s−j)




|P | + j − i |Q| + i − j
Wi (P, ω1 )Wj (Q, ω2 )
s−i s−j

which is especially interesting in (and equivalent to) the case in which P and Q
are chains, where it describes the enumeration by descents and major index of the shuffles of two permutations. Bijective proofs of this formula were later found by
Goulden [29] and by Stadler [61].
7.3. α and β. For any subset S of [p − 1], let α(P, ω; S) be the number of permutations in L (P, ω) with descent set contained in S and let β(P, ω; S) be the number of permutations in L (P, ω) with descent set equal to S. Then α(P, ω; S) =
P
β(P,
ω; T ), so by inclusion-exclusion,
T ⊆S
X
β(P, ω; S) =
(−1)|S|−|T | α(P, ω; T ).
T ⊆S

We can give another interpretation to α(P, ω; S). An order ideal of P is a subset I
of P such that if X ∈ I and Y ≺ X then Y ∈ I. A chain of order ideals in P
∅ = I0 ⊂ I1 ⊂ · · · ⊂ Ik = P
is called ω-compatible if the restriction of ω to each Ii+1 − Ii is order-preserving.
(If ω is natural, then any chain of order ideals is ω-compatible.) It is not hard to see that if the elements of S are m1 < m2 < · · · < ms then α(P, ω; S) is the number of ω-compatible chains ∅ = I0 ⊂ I1 ⊂ · · · ⊂ Is+1 = P in which |Ii | = mi for i = 1, . . . , s: given such a chain we associate to it the permutation consisting of the labels of I1 in increasing order, followed by the labels of I2 − I1 in increasing order, and so on.

12

IRA M. GESSEL

We call two labelings ω1 and ω2 of P equivalent if A (P, ω1 ) = A (P, ω2 ).
Alternatively, ω1 and ω2 are equivalent if whenever Y covers X in P , ω1 (X) > ω1 (Y )
if and only if ω2 (X) > ω2 (Y ). If ω1 and ω2 are equivalent labelings, then for every
S ⊆ [p − 1], we have α(P, ω1 ; S) = α(P, ω2 ; S), and there is a simple bijection between the permutations counted by α(P, ω1 ; S) and those counted by α(P, ω2 ; S).
It follows that β(P, ω1 ; S) = β(P, ω2 ; S). This fact has interesting consequences; for example [73, Exercise 7.95], it can be used to show the existence of Solomon’s descent algebra [60] for the symmetric group.

8. Further Developments
8.1. Posets. As Stanley noted in [66, Section 4], P -partitions are closely related to the distributive lattice J(P ) of order ideals of P , ordered by inclusion. In particular, if ω is a natural labeling then the order polynomial Ω(P, ω; m) is the number of chains of order ideals
∅ = I0 ⊆ I1 ⊆ · · · ⊆ Im = P, and as described in Section 7.3, α(P, ω; S), and thus β(P, ω; S), can be defined in terms of chains in J(P ). These counts of chains make sense in any graded poset with a unique minimal and maximal element, and in this context the analogue of the order polynomial is called the zeta polynomial and α and β are called the flag f -vector and flag h-vector (or rank-selected Möbius invariant). An account of their basic properties can be found in [74, Sections 3.12 and 3.13]. Stanley studied aspects of these concepts in [62], [65], [68], and [69]. Without further conditions the numbers β(S) need not be nonnegative, but if the edges of the Hasse diagram of the poset can be labeled with integers so that whenever s ≤ t there is a unique saturated chain from s to t with nondecreasing labels (an “R-labeling”) then β(S)
has a combinatorial interpretation completely analogous to the P -partition case.
(See [74, Section 3.14].)
Cohen-Macaulay posets [7, 8] are another important class of posets for which
β(S) can be shown to be nonnegative by algebraic or topological methods, but for which β(S) does not in general have a combinatorial interpretation.
8.2. Counting lattice points. If P is a lattice polytope in Rp (the convex hull of a set of lattice points) then the number of lattice points in mP (P dilated by a factor of m) is a polynomial in m, called the Ehrhart polynomial of P. If (P, ω) is naturally labeled then Ω(P, ω; m+1) is the Ehrhart polynomial of the order polytope of P which is the set of all (x1 , . . . , xp ) in Rp satisfying xi ≥ xj whenever i ≺ j, and 0 ≤ xi ≤ 1 for all i. Some of the properties of order polynomials generalize to Ehrhart polynomials, including the reciprocity theorem for order polynomials, equation (7.1). Stanley has made important contributions to the theory of Ehrhart polynomials and their generalizations, as described in Matthias Beck’s paper [3] in this volume; see also [74, sections 4.5–4.6] and [4].
In [71], Stanley defines the chain polytope of the poset P to be the set of points
(x1 , . . . , xp ) in Rp satisfying xi ≥ 0 for all i and xi1 + · · · + xik ≤ 1 for every chain i1 ≺ · · · ≺ ik in P , and proves that this polytope has the same Ehrhart polynomial as the order polytope of P .

P -PARTITIONS

13

8.3. Root systems. Gessel [26, p. 300] suggested that the inequalities that define P -partitions could be generalized to the inequalities determined by the reflecting hyperplanes of a Coxeter group.
Victor Reiner [53, 55] observed that the definition of a P -partition can be restated in terms of the root system of type Ap−1 , which is the set of vectors in
Rp of the form ei − ej , with i 6= j, where ei is ith standard basis vector in Rp .
The positive roots are the roots ei − ej where i < j and the negative roots are the negatives of these. Given a set R of roots, we can consider the set of vectors
σ = (σ1 , . . . , σp ) ∈ Np satisfying hα, σi ≥ 0, for all roots α in R, hα, σi > 0, for all negative roots α in R, where h · , · i is the usual inner product on Rp . Then if this system of inequalities is consistent, the set of solutions will be the set of (P, ω)-partitions for some partial order on P = [p] with the labeling ω(i) = i for all i ∈ [p]. For example, the poset of Figure 1 corresponds the the set of roots {e2 − e1 , e2 − e3 }; the negative root e2 − e1 gives the inequality σ2 − σ1 > 0 and the (positive) root e2 − e3 gives the inequality σ2 − σ3 ≥ 0.
Reiner [54, 55] generalized this idea to arbitrary root systems, which are sets of vectors in Rp satisfying certain properties; each root system consists of a set of positive roots and their negatives, the negative roots. To each root system is associated its Weyl group, which is the finite group of isometries of Rp generated by the reflections in the roots. (So for the root system of type Ap−1 , the Weyl group is the symmetric group Sp .) Reiner shows that there is a generalization of the fundamental theorem of P -partitions to root systems, in which the role of the symmetric group is replaced by the corresponding Weyl group.
He then studies in particular the case of the root system of type Bp , in which the positive roots are ei for 1 ≤ i ≤ p and ei +ej and ei −ej for 1 ≤ i < j ≤ n. Thus, for example, if we take the roots e1 , −e2 − e3 , and e3 − e1 , then the corresponding inequalities are σ1 ≥ 0, −σ2 − σ3 > 0 and σ3 − σ1 < 0. (Unlike the case of ordinary
P -partitions, here we allow the σi to take on arbitrary integer values.) The elements of the Weyl group of type Bp , the hyperoctahedral group, may be viewed as signed permutations, which are permutations π of the set ±[p] = {−p, · · · , −1, 1, · · · p}, such that π(−i) = −π(i) for all i ∈ ±[p]; a signed permutation is determined by its values on [p]. The general definition of descent set for a Weyl group reduces in this case to
S (π) = { i ∈ [p] | π(i) > π(i + 1) }, where we take π(p + 1) = p + 1 and use the order7 1 < 2 < · · · < p + 1 < −p <
· · · < −2 < −1. Reiner then obtains for signed permutations analogues of all the basic P -partition for ordinary permutations.
Chak-On Chow [15, Section 2] studied “P -partitions of type B” with a closely related, but somewhat different approach, using “type B posets” which are partial orders ≺ on the set {−p, · · · , −1, 0, 1, · · · , p} with the property that i ≺ j if and only if −j ≺ −i.
Further work on root system analogues of P -partitions was undertaken by John
Stembridge in his study of Coxeter cones [77].
7Different choices for the roots would allow similar results with the usual order on ±[p].

14

IRA M. GESSEL

8.4. Lexicographic inequalities. MacMahon, in [38] and other papers (see
[40, Chapter 8]), studied multipartite partitions which are expressions of “multipartite numbers” as sums of multipartite numbers, where a multipartite number is a tuple of nonnegative integers. The number of partitions of the multipartite number
(n1 , . . . , ns ) with p parts, where (0, . . . , 0) is allowed as a part, is the coefficient of q1n1 · · · qsns in φr (x), where
∞
X

φp (q1 , . . . , qs )z p =

p=0

∞
Y

1

1 − q1k1 · · · qsks z k1 ,...,ks =0

It is not hard to show that there exist polynomials Λp (q1 , . . . , qs ) such that
φp (q1 , . . . , qs ) =

Λp (q1 , . . . , qs )
(q1 ; q1 )p · · · (qs ; qs )p

where (q; q)s = (1 − q) · · · (1 − q s ). E. M. Wright [81] conjectured that the polynomials Λp (q1 , . . . , qs ) have nonnegative coefficients, and his conjecture was proved by Basil Gordon [28] in 1963. We illustrate Gordon’s approach with the simplest example, when p = s = 2. An unordered pair (a1 , b1 ), (a2 , b2 ) of bipartite numbers may be arranged in decreasing lexicographic order, so counting such pairs is equivalent to counting solutions of the lexicographic inequality (a1 , b1 ) ≥ (a2 , b2 ). The set of solutions of this lexicographic inequality is the disjoint union of the solutions of a1 ≥ a2 , b1 ≥ b2
and a1 > a2 , b1 < b2 .
The solutions of the first inequalities contribute 1/(q1 ; q1 )2 (q2 ; q2 )2 to Λ2 (q1 , q2 )
and the solutions of the second inequalities contribute q1 q2 /(q1 ; q1 )2 (q2 ; q2 )2 , so
Λ2 (q1 , q2 ) = 1 + q1 q2 .
Gordon proved Wright’s conjecture by showing that the lexicographic inequalities specifying the terms in φp (q1 , . . . , qs ) can in general be decomposed in a similar way. D. P. Roselle [58], using essentially the same approach, gave a simple combinatorial interpretation to the coefficient of q1i1 q2i2 in Λr (q1 , q2 ); it is the number of permutations π of [p] such that maj(π) = i1 and maj(π −1 ) = i2 . Garsia and Gessel [23] showed that, more generally, the coefficient of q1i1 · · · qsis in
Λp (q1 , · · · , qs ) is the number of s-tuples (π1 , . . . , πs ) of permutations of [p] whose product π1 · · · πs is the identity permutation such that maj(πj ) = ij for each j.
They also showed that by considering multipartite partitions with bounded part sizes, one can count these s-tuples of permutations by major index and number of descents. Further results along these lines have been found by a number of authors
[5, 20, 21, 45, 49, 53, 54].
Gessel [26] studied “multipartite P -partitions” in which the inequalities defining (P, ω)-partitions are applied to multipartite numbers ordered lexicographically; his results enumerate s-tuples of permutations whose product is in L (P, ω) by their descent sets.
Another application of lexicographic inequalities related to P -partitions was given by Gessel and Reutenauer [27]. A Lyndon word is a sequence of nonnegative integers that is lexicographically strictly less than all of its cyclic permutations.

P -PARTITIONS

15

Thus the word (a1 , a2 , a3 ) is a Lyndon word if and only if (a1 , a2 , a3 ) < (a2 , a3 , a1 )
and (a1 , a2 , a3 ) < (a3 , a1 , a2 ). The set of solutions of these lexicographic inequalities is the disjoint union of the solutions of a1 ≤ a2 < a3
and a1 < a3 ≤ a2 .
Gessel and Reutenauer showed that a similar decomposition exists for Lyndon words of any length, and more generally, for multisets of Lyndon words, and this allowed them to count permutations of a given cycle type by their descent sets. A generalization to hyperoctahedral groups was given by Stéphane Poirier [51].
8.5. Quasi-symmetric functions. In his memoir [66], Stanley considered the generating function for (P, ω)-partitions
X
σ(1) σ(2)
F (P, ω) =
x1 x2 · · · xσ(p)
p
σ∈A (P,ω)

in which different (P, ω)-partitions contribute different terms. The theory of Schur functions (see, for example, [73, Section 7.10]) suggests looking at the less refined generating function
X
(8.1)
Γ(P, ω) =
xσ(1) xσ(2) · · · xσ(p)
σ∈A (P,ω)

discussed briefly by Stanley [66, p. 81] and in more detail by Gessel8 [26]. (See also
[73, Section 7.19].) By the fundamental theorem of P -partitions,
X
Γ(P, ω) =
Γ(π, ω).
π∈L (P,ω)

The fundamental quasi-symmetric functions, denoted Fα or Lα , are indexed by compositions (sequences of positive integers) and defined as follows: if α = (α1 , . . . ak )
is a composition of p then
X
Fα =
xi1 xi2 · · · xip where the sum is over all i1 ≤ i2 ≤ · · · ≤ ip satisfying ij < ij+1 if j ∈ {α1 , α1 +
α2 , . . . , α1 +· · ·+αk−1 }. It is not hard to show that the Fα are linearly independent and generate a ring, called the ring of quasi-symmetric functions, that contains the ring of symmetric functions [73, Chapter 7]. Thus the information contained in
Γ(P, ω) is precisely the multiset of descent sets of the permutations in L (P, ω); i.e., numbers β(P, ω; S). The quasi-symmetric generating function (8.1) extends to an encoding of the flag h-vector of a graded poset, as studied by Richard Ehrenborg
[16].
The theory of quasi-symmetric functions has proven useful in a number of enumeration problems. For example, Stanley [70] used them to define what are now called “Stanley symmetric functions” in the study of reduced decompositions in symmetric groups.
8Gessel took P -partitions to be order-preserving, rather than order-reversing, so his Γ(P, ω)
is slightly different from that defined here.

16

IRA M. GESSEL

There is a comultiplication on the ring of quasi-symmetric functions that makes it into a Hopf algebra (and an “internal” comultiplication that makes it a bialgebra). The dual Hopf algebra is the algebra of non-commutative symmetric functions that has been studied extensively by Jean-Yves Thibon and others in a series of papers beginning with [24]; see also Malvenuto and Reutenauer [42]. Type B
quasi-symmetric functions and noncommutative symmetric functions were studied by Chow [15]. Another related algebra with enumerative applications is the
Malvenuto-Reutenauer algebra [1, 41, 42].
A different encoding of the flag h-vector of a poset is the ab-index, which is especially useful in studying Eulerian posets [17, 72].
8.6. Enriched P -partitions. John Stembridge [75] introduced a generalization of (P, ω)-partitions that interpolates between (P, ω)-partitions and (P, ω̄)partitions. We introduce the following total ordering on the set P0 of nonzero integers:
−1 < +1 < −2 < +2 < −3 < +3 < · · · ; for k ∈ P0 the notations k > 0 and |k| retain their usual meanings. Then an enriched
(P, ω)-partition is a map σ : P → P0 such that for all X ≺ Y in P we have9
(i) σ(X) ≥ σ(Y )
(ii) If σ(X) = σ(Y ) > 0 then ω(X) < ω(Y )
(iii) If σ(X) = σ(Y ) < 0 then ω(X) > ω(Y ).
Note that if the image of σ lies in {+1, +2, . . . } then (iii) is vacuous and (ii) is equivalent to the condition that if ω(X) > ω(Y ) then σ(X) > σ(Y ), so σ is an ordinary
(P, ω)-partition, and if the image of σ lies in {−1, −2, · · · } then (ii) is vacuous and
(iii) is equivalent to the condition that if ω(X) < ω(Y ) then σ(X) > σ(Y ), so σ
is essentially a (P, ω̄)-partition. Stembridge proves a version of the fundamental theorem for enriched (P, ω)-partitions, and defines the quasi-symmetric generating function
X Y
∆(P, ω) =
x|σ(X)| ,
σ X∈P

so by the fundamental theorem,
∆(P, ω) =

X

∆(π, ω).

π∈L (P,ω)

It is a remarkable fact that ∆(π, ω) depends only on the peak set of π, that is, the set { i | π(i − 1) < π(i) > π(i + 1) }. The distinct ∆(π, ω) form a basis for a subalgebra of the algebra of quasi-symmetric functions.
Stembridge also discusses the analogue of the order polynomial for enriched
P -partitions and studies cases in which ∆(P, ω) is symmetric, which are related to
Schur’s Q-functions.
Kathryn Nyman [47] used enriched P -partitions to prove the existence of the
“peak algebra” of the symmetric group. T. Kyle Petersen [50] studied type B
enriched P -partitions and applied them to type B peak algebras. Enriched P partitions have also been applied to the study of chains in Eulerian posets [6].
9Stembridge defined enriched P -partitions to be order-preserving; for consistency we define them here to be order-reversing.

P -PARTITIONS

17

8.7. Additional applications and developments. SeungKyung Park [48]
studied naturally labeled posets P whose order polynomial ΩP (n) is the Stirling number of the second kind S(k + n, n), thereby giving a new combinatorial interpretation and q-analogue to the polynomials Bk (t) defined by
∞
X
n=0

S(k + n, n)tn =

Bk (t)
.
(1 − t)2k+1

Combinatorial interpretations for these polynomials had been given earlier by John
Riordan [57] and by Gessel and Stanley [25].
Sangwook Ree [52] applied P -partitions to count restricted lattice paths in the plane by left turns, obtaining generalizations of q-Narayana numbers. Petter
Brändén [10], taking a similar approach, gave several interpretations to q-Narayana numbers in counting Dyck paths.
Joseph Neggers [46]
P conjectured in 1978 that for any naturally labeled poset
(P, ω) the polynomial π∈L (P,ω) tdes(π) has all real roots. In 1986, Stanley conjectured that this holds more generally for any labeled poset (P, ω) (see [11]). Stanley’s conjecture was disproved by Petter Brändén [9] in 2004 and Neggers’s conjecture was disproved by John Stembridge [76] in 2007.
Peter McNamara and Christophe Reutenauer [43] used P -partitions to study idempotents in the group algebra of the symmetric group.
McNamara and Ryan Ward [44] studied the question of when two different labeled posets have the same generating function (8.1).
Valentin Féray and Victor Reiner [18] explored connections between P -partitions and commutative algebra,P
and in particular described a class of naturally labeled posets for which the sum π∈L (P ) q maj(π) factors nicely.
Loı̈c Foissy and Claudia Malvenuto [12] reinterpreted the fundamental theorem of P -partitions as an injective Hopf algebra morphism and generalized it to preorders, leading to a Hopf algebra on finite topologies.
Acknowledgments. I would like to thank Christian Krattenthaler, Claudia
Malvenuto, T. Kyle Petersen, Victor Reiner, and John Stembridge for their helpful suggestions, and Richard Stanley for his development of the theory of P -partitions.
References
[1] Marcelo Aguiar and Frank Sottile, Structure of the Malvenuto-Reutenauer Hopf algebra of permutations, Adv. Math. 191 (2005), no. 2, 225–275.
[2] A. O. L. Atkin, P. Bratley, I. G. Macdonald, and J. K. S. McKay, Some computations for m-dimensional partitions, Proc. Cambridge Philos. Soc. 63 (1967), 1097–1100.
[3] Matthias Beck, Stanley’s major contributions to Ehrhart theory, arXiv:1407.0255 [math.CO].
[4] Matthias Beck, Combinatorial reciprocity theorems, Jahresber. Dtsch. Math.-Ver. 114 (2012), no. 1, 3–22.
[5] Riccardo Biagioli and Jiang Zeng, Enumerating wreath products via Garsia-Gessel bijections,
European J. Combin. 32 (2011), no. 4, 538–553.
[6] Louis J. Billera, Samuel K. Hsiao, and Stephanie van Willigenburg, Peak quasisymmetric functions and Eulerian enumeration, Adv. Math. 176 (2003), no. 2, 248–276.
[7] A. Björner, A. M. Garsia, and R. P. Stanley, An introduction to Cohen-Macaulay partially ordered sets, Ordered sets (Banff, Alta., 1981), NATO Adv. Study Inst. Ser. C: Math. Phys.
Sci., vol. 83, Reidel, Dordrecht-Boston, Mass., 1982, pp. 583–615.
[8] Anders Björner, Shellable and Cohen-Macaulay partially ordered sets, Trans. Amer. Math.
Soc. 260 (1980), no. 1, 159–183.

18

IRA M. GESSEL

[9] Petter Brändén, Counterexamples to the Neggers-Stanley conjecture, Electron. Res. Announc.
Amer. Math. Soc. 10 (2004), 155–158 (electronic).
[10]
, q-Narayana numbers and the flag h-vector of J(2 × n), Discrete Math. 281 (2004), no. 1-3, 67–81.
[11] Francesco Brenti, Unimodal, log-concave and Pólya frequency sequences in combinatorics,
Mem. Amer. Math. Soc. 81 (1989), no. 413, viii+106.
[12] Loı̈c Foissy and Claudia Malvenuto, The Hopf algebra of finite topologies and T-partitions, preprint, 2014, http://arxiv.org/abs/1407.0476v2.
[13] L. Carlitz, q-Bernoulli and Eulerian numbers, Trans. Amer. Math. Soc. 76 (1954), 332–350.
[14]
, A combinatorial property of q-Eulerian numbers, Amer. Math. Monthly 82 (1975),
51–54.
[15] Chak-On Chow, Noncommutative Symmetric Functions of Type B, Ph.D. thesis, Massachusetts Institute of Technology, 2001.
[16] Richard Ehrenborg, On posets and Hopf algebras, Adv. Math. 119 (1996), no. 1, 1–25.
[17] Richard Ehrenborg and Margaret Readdy, Coproducts and the cd-index, J. Algebraic Combin.
8 (1998), no. 3, 273–299.
[18] Valentin Féray and Victor Reiner, P -partitions revisited, J. Commut. Algebra 4 (2012), no. 1,
101–152.
[19] Dominique Foata, Distributions eulériennes et mahoniennes sur le groupe des permutations,
Higher combinatorics (Proc. NATO Advanced Study Inst., Berlin, 1976), NATO Adv. Study
Inst. Ser., Ser. C: Math. Phys. Sci., vol. 31, Reidel, Dordrecht-Boston, Mass., 1977, With a comment by Richard P. Stanley, pp. 27–49.
[20] Dominique Foata and Guo-Niu Han, Signed words and permutations. II. The Euler-Mahonian polynomials, Electron. J. Combin. 11 (2004/06), no. 2, Research Paper 22, 18.
[21]
, Signed words and permutations. III. The MacMahon Verfahren, Sém. Lothar. Combin. 54 (2005/07), Art. B54a, 20.
[22] Dominique Foata and Marcel-Paul Schützenberger, Major index and inversion number of permutations, Math. Nachr. 83 (1978), 143–159.
[23] A. M. Garsia and I. Gessel, Permutation statistics and partitions, Adv. in Math. 31 (1979), no. 3, 288–305.
[24] Israel M. Gelfand, Daniel Krob, Alain Lascoux, Bernard Leclerc, Vladimir S. Retakh, and
Jean-Yves Thibon, Noncommutative symmetric functions, Adv. Math. 112 (1995), no. 2,
218–348.
[25] Ira Gessel and Richard P. Stanley, Stirling polynomials, J. Combinatorial Theory Ser. A 24
(1978), no. 1, 24–33.
[26] Ira M. Gessel, Multipartite P -partitions and inner products of skew Schur functions, Combinatorics and algebra (Boulder, Colo., 1983), Contemp. Math., vol. 34, Amer. Math. Soc.,
Providence, RI, 1984, pp. 289–317.
[27] Ira M. Gessel and Christophe Reutenauer, Counting permutations with given cycle structure and descent set, J. Combin. Theory Ser. A 64 (1993), no. 2, 189–215.
[28] B. Gordon, Two theorems on multipartite partitions, J. London Math. Soc. 38 (1963), 459–
464.
[29] I. P. Goulden, A bijective proof of Stanley’s shuffling theorem, Trans. Amer. Math. Soc. 288
(1985), no. 1, 147–160.
[30] Donald E. Knuth, A note on solid partitions, Math. Comp. 24 (1970), 955–961.
[31] G. Kreweras, Sur une classe de problèmes de dénombrement liés au treillis des partitions des entiers, Cahier de B.U.R.O. (1965), no. 6, 15–67.
[32]
, Sur une extension du problème dit “de Simon Newcomb”, C. R. Acad. Sci. Paris
Sér. A-B 263 (1966), A43–A45.
[33]
, Traitement simultané du “Problème de Young” et du “Problème de Simon Newcomb”, Cahiers de B.U.R.O. (1967), no. 10, 23–31.
[34]
, Polynômes de Stanley et extensions linéaires d’un ordre partiel, M

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
