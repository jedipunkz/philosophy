---
source: "https://arxiv.org/abs/1608.07762v1"
title: "A class of Ramsey-extremal hypergraphs"
author: "Brendan D. McKay"
year: "2016"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1608.07762v1"
pdf: "https://arxiv.org/pdf/1608.07762v1"
captured_at: "2026-06-24T11:11:02Z"
updated_at: "2026-06-24T11:11:02Z"
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

# A class of Ramsey-extremal hypergraphs

- 著者: Brendan D. McKay
- 年: 2016
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1608.07762v1)
- ダウンロード: https://arxiv.org/pdf/1608.07762v1
- PDF: https://arxiv.org/pdf/1608.07762v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

In 1991, McKay and Radziszowski proved that, however each 3-subset of a 13-set is assigned one of two colours, there is some 4-subset whose four 3-subsets have the same colour. More than 25 years later, this remains the only non-trivial classical Ramsey number known for hypergraphs. In this article, we find all the extremal colourings of the 3-subsets of a 12-set and list some of their properties. Using the catalogue, we provide an answer to a question of Dudek, Fleur, Mubayi and Rödl about the size-Ramsey numbers of hypergraphs.

## PDF Text

arXiv:1608.07762v1 [math.CO] 28 Aug 2016

A class of Ramsey-extremal hypergraphs
Brendan D. McKay∗
Research School of Computer Science
Australian National University
Canberra ACT 2601, Australia brendan.mckay@anu.edu.au

Abstract
In 1991, McKay and Radziszowski proved that, however each 3-subset of a 13-set is assigned one of two colours, there is some 4-subset whose four 3-subsets have the same colour. More than 25 years later, this remains the only non-trivial classical Ramsey number known for hypergraphs. In this article, we find all the extremal colourings of the 3-subsets of a 12-set and list some of their properties. Using the catalogue, we provide an answer to a question of Dudek, Fleur, Mubayi and Rödl about the size-Ramsey numbers of hypergraphs.

1

Introduction

A colouring of all the s-subsets of an n-set with two colours is called R(j, k; s)-good if there is no j-subset (of the n-set) containing only s-subsets of the first colour, and no k-subset containing only s-subsets of the second colour. (Note that it is the s-subsets receiving colours, not the elements of the n-set.) The Ramsey number R(j, k; s) is defined to be the least n for which there is no R(j, k; s)-good colouring.
Although there are several known values of R(j, k; 2) [8], which is usually written as just
R(j, k), the only known non-trivial value of R(j, k; s) for s ≥ 3 is R(4, 4; 3) = 13. As a lower bound, a suitable colouring of the 3-subsets of a 12-set was presented by Isbell in 1969 [2], and this was proved best possible by the present author and Radziszowski in 1991 [6]. During that project we found more than 200,000 R(4, 4; 3)-good colourings for 12 points, but did not have the resources to compute them all. With the aid of an improved algorithm and the much greater computing resources available today, we can now show that the number of
∗

Research supported by the Australian Research Council.

1

R(4, 4; 3)-good colourings for 12 points is precisely 434,714. We hope that this compilation of data will assist further investigations.

2

Method

We prefer to use slightly different terminology for this description. Suppose we have an
R(4, 4; 3)-good colouring of the 3-subsets of an n-set V . We will call the 4-subsets of V
quadruples.
If we choose just the 3-subsets of V having the first colour, we obtain a 3-uniform hypergraph on V with the property that every quadruple contains 1, 2 or 3 edges (the other possibilities 0 and 4 being forbidden). We will call this a R(4, 4; 3)-good hypergraph. Note that we could have chosen the other colour instead and would have obtained the complementary hypergraph. We can obvious recover the colouring from the hypergraph, so we lose nothing by continuing with hypergraph terminology.
Denote by R(n) the set of R(4, 4; 3)-good hypergraphs with n points. If we wish to emphasize the point set V , we may write R(V ) instead. More generally, R(n, e) is the set of
R(4, 4; 3)-good hypergraphs with n points and e edges, and notations like R(V, ≤110) have their obvious meanings.
Our aim is to find R(12). By the remark just made, it will suffice to find R(12, ≤110),

since 110 = 12 12
and the rest are complements. Given G ∈ R(V ) and v ∈ V , define Gv
3
to be the hypergraph with point set V −v and all the edges of G that lie in V −v. Clearly
Gv ∈ R(V −v). Since the points of G ∈ R(n, e) lie on average in 3e/n edges, we find that for G ∈ R(12, ≤110) there is some v such that Gv ∈ R(11, ≤82). Continuing such logic we find a construction path
R(9, ≤41) → R(10, ≤59) → R(11, ≤82) → R(12, ≤110).

(2.1)

Each step in (2.1) involves adding one point and some edges that include the new point.
Moreover, we can assume that the new point is in at least as many edges as any of the old points (after the new edges are added).
The programs developed for [6] are fast enough to find R(9, ≤41) in a few hours. There are exactly 3,030,480,232 such hypergraphs and these form our starting point. It would be convenient to perform each of the three steps of (2.1) separately, but it would be quite expensive. The number of hypergraphs in R(10) and R(11) is greater than 1011 and even the task of extending one hypergraph by one point requires solution of a large set of integer inequalities. We need a better way.
If S is a set and B ⊆ T ⊆ S, then the interval [B, T ] is {X ⊆ S | B ⊆ X ⊆ T }. The use of intervals for solving sets of inequalities efficiently was introduced in [7].
2

Define V9 = {0, 1, . . . , 8} and V10 = V9 ∪ {a}. Consider extending G9 ∈ R(V9 ) to all possible G10 ∈ R(V10 ) by adding the point a and some edges that include a. The possible edges all have the form {i, j, a} where i, j ∈ V9 ; number these e0 , e1 , . . . , e35 in some order.
Each solution for G10 corresponds to a subset of {e0 , e1 , . . . , e35 }.
Now consider the constraints required for G10 to be R(4, 4; 3)-good. The quadruples within V9 are fine already, since we are not adding any further edges inside V9 . So consider a quadruple {i, j, k, a}, where i, j, k ∈ V9 . If {i, j, k} is an edge of G9 , we need that at least one of the edges {i, j, a}, {i, k, a}, {j, k, a} is not selected, while if {i, j, k} is not an edge of
G9 , at least one of those three edges must be selected.
Now we can describe how intervals are used to process many cases simultaneously. Consider one interval [B, T ] ⊆ {e0 , e1 , . . . , e35 } and one quadruple {i, j, k, a}. Define X =
{er , es , et }, where er = {i, j, a}, es = {i, k, a} and et = {j, k, a}. Now we apply the following collapsing rules:
{i, j, k} ∈ G ? B ∩ X
NO
6= ∅
∅
∅
∅
∅

T ∩X
any
∅
{i}
{i, j}
{i, j, k}

replace [B, T ] by
[B, T ]
nothing
[B+i, T ]
[B+i, T ], [B+j, T −i]
[B+i, T ], [B+j, T −i], [B+k, T −i−j]

{i, j, k} ∈ G ?
YES

B̄ ∩ X
any
∅
{i}
{i, j}
{i, j, k}

replace [B, T ] by
[B, T ]
nothing
[B, T −i]
[B, T −i], [B+i, T −j]
[B, T −i], [B+i, T −j], [B+i+j, T −k]

T̄ ∩ X
6= ∅
∅
∅
∅
∅

Figure 1: Collapsing rules for an interval [B, T ] based on quadruple {i, j, k, a}.
By considering each case, we find that the effect of the collapsing rules is to replace
[B, T ] by a set of disjoint intervals whose union is the set of all sets in [B, T ] that satisfy the quadruple {i, j, k, a}. For best practical performance, subsets of {e0 , e1 , . . . , e35 } can be represented by the bits in a single machine word, then the collapsing rules can be implemented in a few machine instructions each.
Starting with the interval [∅, {e0 , e1 , . . . , e35 }] we apply the collapsing rules for each quadruple {i, j, k, a}. The result is a set of disjoint intervals (typically a few hundred)
whose union gives exactly the set of all extensions of G9 to R(10). The efficiency depends a lot on the order in which quadruples are processed; we found a good order by experiment.
Now consider further extension to R(4, 4; 3)-good hypergraphs on V11 = {0, . . . , 8, a, b}.
3

The edges we need to add in total to G9 either have the form {i, j, a} (already added in making G10 ), {i, j, b}, or {i, a, b}, where in each case i, j, k ∈ V9 . Here we can make an observation that is key to the whole computation: The sets of edges {i, j, b} which satisfy quadruples of the form {i, j, k, b} are the same as the sets of edges {i, j, a} which satisfy quadruples of the form {i, j, k, b}, except that a is replaced by b.
Given this observation, we make the possibilities for G11 as follows, given G9 , a set I
of intervals describing the extensions of G9 to R(10), and a particular extension G10 . The possible new edges are numbered e0 , . . . , e44 , where e0 , . . . , e35 are edges of the form {i, j, b}
numbered in the same order as we numbered the edges {i, j, a} in the previous step, and e36 , . . . , e44 are the edges of the form {i, a, b} in any order. To find all solutions, instead of starting with the single interval [∅, {e0 , e1 , . . . , e44 }] as in the previous step, we start with the set of intervals [B, T ∪ {e36 , . . . , e44 }] for [B, T ] ∈ I. Then we avoid collapsing rules which are unnecessary for the stated reasons. This results in a massive speedup.
To complete the process by extending from 11 to 12 points, we use the same idea to begin with intervals obtained during the extension to 11 points. This phase is very fast as most intervals are destroyed very quickly and only a comparatively small number of solutions are found.
It would be possible to apply the general method of [4] to perform exhaustive isomorph reduction at each step in the computation, but the large number of intermediate hypergraphs makes that unwise. Instead, we applied a weaker filter. For a hypergraph with points V and
P
point v ∈ V , define dv to be the number of edges that include v. Also define fv = e dv dw dx , where the sum is over all edges e = {v, w, x} that include v. Suppose we make G ∈ G(V ) by extending a smaller hypergraph, and that v ∈ V is the last point added. The construction path (2.1) assumed that dv ≥ dw for all w ∈ V , so that is the first filter applied. If that doesn’t eliminate G, we also require that fv be maximum out of all w ∈ V with maximum dw . These rules eliminate most isomorphs and are fast to apply. When we finally have a collection of
R(4, 4; 3)-good hypergraphs on 12 points, we perform complete isomorphism reduction using nauty [5].

3

Results

There are about 8.4 × 1011 R(4, 4; 3)-good hypergraphs altogether, including 434,714 with 12
points. Table 1 details the numbers of R(4, 4; 3)-good hypergraphs for each number of points and edges. For 10 and 11 points we only did incomplete isomorph reduction, as explained above; hence the totals for those sizes are estimates.
The automorphism group Aut(G) of a hypergraph G ∈ R(V ) is the set of permutations of V which preserve the edge set. As detailed in Table 2, most hypergraphs in R(12) have a trivial group and none have a transitive group. The unique hypergraph with |Aut(G)| = 60,
4

n
3

e
0
total

4

1
2
total

5

3
4
5
total

6

6
7
8
9
10
total
7
12
13
14
15
16
17
total
8
21
22
23
24
25
26
27
28
total

count
1
2
1
1
3
1
3
4
12
1
5
22
50
70
226
1
26
338
1793
5055
8317
31060
1
278
9763
107241
573596
1764747
3380337
4182459
15854385

n
9

e
33
34
35
36
37
38
39
40
41
42
total
10
50
51
52
···
total
11
73
74
75
···
total
12
104
105
106
107
108
109
110
total

count
2
204
22616
774043
10877731
79336073
341024774
928650036
1669794753
2025923846
8086884310
13
1810
121356
≈ 6.2×1011
36
4725
246299
≈ 2.1×1011
4
123
1465
10235
41939
98235
130712
434714

Table 1: The numbers of R(4, 4; 3)-good hypergraphs with n points and e edges. The totals include complements.

5

which has two orbits of size 6, is presented in Figure 2 using letters for elements of V . This hypergraph is one of the 1306 in R(12) that are self-complementary and is isomorphic to the one found by Isbell [2].
|Aut(G)| orbits count
1
12
432300
2
6
18
7
112
8
1669
4
529
3
4
6
32
2
20
6
4
17
10
4
1
2
15
12
60
2
1
Table 2: Counts of R(12) by automophism group.
None of the hypergraphs in R(12) extend to a hypergraph in R(13), consistently with the finding of [6] that R(13) = ∅. This raises the question of how close we can get to a hypergraph
(3)
in R(13); specifically, how many edges of the complete hypergraph K13 can we colour without
(3)
obtaining a monochromatic induced K4 ? The generation method described in the previous section can be easily adapted to ignore particular quadruples. If we ignore the constraints normally attributed to the quadruples {i, j, k, a} which contain a specified {i, j, a}, then we are colouring the edges of the complete hypergraph except for one uncoloured edge. Using
(3)
this method we found that K13 minus one edge cannot be coloured with two colours without
(3)
creating a monochromatic K4 .
(3)

On the other hand, if we omit two edges of K13 , a colouring without a monochromatic
(3)
K4 may be possible. In Figure 3 we give examples where the two omitted edges overlap in one or two points. We did not find any examples with the omitted two edges being disjoint, but our search in that case was not exhaustive. We can report these partial results: there is
(3)
no good colouring of K13 minus the edges {1, 2, 3} and {4, 5, 6} such that a good colouring of
(3)
K12 can be obtained either by deleting vertex 1 and colouring edge {4, 5, 6}, or by deleting vertex 7 and colouring both edges {1, 2, 3} and {4, 5, 6}. We propose the remaining cases of two disjoint edges as a challenge for the reader.
If H is a 3-uniform hypergraph, the size-Ramsey number R̂(3) (H) is the least number m such that for some 3-uniform hypergraph G with m edges, every colouring of the
(3)
edges of G with two colours includes a monochromatic copy of H. If H = K4 , then the

(3)
value R(4, 4; 3) = 13 implies that R̂(3) (H) ≤ 13
= 286 since we can take G = K13 .
3
6

Dudek, La Fleur, Mubayi and Rödl [1, Question 2.2] ask whether this bound is sharp. Since
(3)
(3)
K13 minus one edge cannot be coloured without creating a monochromatic K4 we have
R̂(3) (H) ≤ 285, which answers Dudek et al.’s question in the negative.
The extremal R(4, 4; 3)-good hypergraphs are available online [3]. Finally, we thank
Staszek Radziszowski for many useful comments.
Let Γ = h(cd)(ef)(CD)(EF), (bc)(de)(BC)(DE), (ab)(ef)(AB)(EF)i be a permutation acting on the points abcdefABCDEF. It is isomorphic to the alternating group A5 and acts 2-transitively on each of its orbits {a, . . . , f}
and {A, . . . , F}. Now construct a hypergraph by applying Γ to each of the starting edges {abe, ABE, abC, aAB, cAB}. These provide 10, 10, 30, 30 and
30 edges, respectively. The hypergraph induced by each orbit is the same
2-(6,3,2) design. The relabelling (aD)(bC)(cB)(dA)(eF)(fE) takes the hypergraph onto its complement.

Figure 2: The unique hypergraph in R(12) with automorphism group of order 60.

acd bcd abe ace bce cde adf cdf def adg aeg beg ceg deg afg bfg efg ach bch adh bdh aeh beh deh afh efh bgh dgh fgh bdi cdi bei cei afi bfi dfi efi bgi cgi dgi ahi chi dhi ehi abj acj bcj cdj aej dej bfj cfj agj bgj cgj dgj bhj ehj fhj aij gij abk bck bdk cek afk bfk cfk efk cgk fgk dhk ghk aik bik eik hik bjk djk gjk hjk ijk abl acl bcl adl bel afl bfl cfl bgl dgl chl fhl ghl eil fil gil hil ajl djl ejl fjl ijl akl ckl dkl ekl hkl abm adm bdm aem dem bfm cfm dfm cgm egm ahm bhm chm ghm aim cim fim ejm fjm hjm ijm akm dkm ekm fkm gkm jkm blm clm dlm elm glm ilm
Omitted edges: abc ade

bcd cde acf bcf aef def adg bdg cdg aeg beg deg bfg efg abh ach bch adh bdh beh cfh egh fgh aci aei bei cei dei afi dfi agi cgi fgi bhi dhi fhi adj bdj aej cej bfj cfj dfj agj bgj ahj bhj chj ehj ghj bij dij eij abk ack bck cdk aek bek cek cfk dfk efk bgk cgk fgk ahk dhk fhk aik bik gik hik ejk fjk gjk hjk ijk abl acl bcl adl bdl afl bfl dfl efl cgl dgl fgl chl dhl ehl ghl bil cil eil fil hil ajl cjl ejl fjl gjl bkl dkl gkl abm adm bdm cdm bem cem afm bfm cfm dfm efm agm bgm ahm ehm fhm ghm cim fim gim cjm djm gjm ijm dkm ekm hkm jkm blm elm ilm jlm klm
Omitted edges: abc abd
(3)

Figure 3: Two R(4, 4; 3)-good colourings of the complete hypergraph K13 minus two edges.
Edges not mentioned have the second colour.

7

References
[1] A. Dudek, S. La Fleur, D. Mubayi and V. Rödl, On the size-Ramsey number of hypergraphs, preprint 2015, arXiv:1503.06304.
[2] J. R. Isbell, R(4, 4; 3) ≥ 13, J. Combin. Th., 6 (1969) 210.
[3] B. D. McKay, Combinatorial data. Online at http://users.cecs.anu.edu.au/∼bdm/data/ramsey.html.
[4] B. D. McKay, Isomorph-free exhaustive generation, J. Algorithms, 26 (1998) 306–324.
[5] B. D. McKay and A. Piperno, Practical Graph Isomorphism, II. J. Symbolic Comput.,
60 (2014) 94–112.
[6] B. D. McKay and S. P. Radziszowski, The first classical Ramsey number for hypergraphs is computed, Proceedings of the Second Annual ACM-SIAM Symposium on Discrete
Algorithms, SODA’91, San Francisco, (1991) 304–308.
[7] B. D. McKay and S. P. Radziszowski, R(4, 5) = 25, J. Graph Theory, 19 (1995) 309–322.
[8] S. P. Radziszowski, Small Ramsey Numbers, Electron. J. Combin., Dynamic Survey
DS1, 1994–2014.

8

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
