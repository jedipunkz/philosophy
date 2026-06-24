---
source: "https://arxiv.org/abs/1503.05934v2"
title: "Plane partitions in the work of Richard Stanley and his school"
author: "C. Krattenthaler"
year: "2015"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1503.05934v2"
pdf: "https://arxiv.org/pdf/1503.05934v2"
captured_at: "2026-06-24T11:10:29Z"
updated_at: "2026-06-24T11:10:29Z"
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

# Plane partitions in the work of Richard Stanley and his school

- 著者: C. Krattenthaler
- 年: 2015
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1503.05934v2)
- ダウンロード: https://arxiv.org/pdf/1503.05934v2
- PDF: https://arxiv.org/pdf/1503.05934v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

These notes provide a survey of the theory of plane partitions, seen through the glasses of the work of Richard Stanley and his school.

## PDF Text

arXiv:1503.05934v2 [math.CO] 15 May 2015

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND
HIS SCHOOL
C. KRATTENTHALER
Abstract. These notes provide a survey of the theory of plane partitions, seen through the glasses of the work of Richard Stanley and his school.

1. Introduction
Plane partitions were introduced to (combinatorial) mathematics by Major Percy Alexander MacMahon [71] around 1900. What he had in mind was a planar analogue of a(n integer) partition.1
In order to explain this, let us start with the definition of a(n integer) partition. A
partition of a positive integer n is a way to represent n as a sum of positive integers, where the order of the summands does not play any role. So, a partition of n is n = λ1 + λ2 + · · · + λk ,

(1.1)

for some k, where all summands λi are positive integers and, since the order of summands is irrelevant, we may assume without loss of generality that λ1 ≥ λ2 ≥ · · · ≥ λk > 0. For example, there are 7 partitions of 5, namely
5 = 4 + 1 = 3 + 2 = 3 + 1 + 1 = 2 + 2 + 1 = 2 + 1 + 1 + 1 = 1 + 1 + 1 + 1 + 1.
If we want to represent a partition as in (1.1) very succinctly, then we just write
λ1 λ2 . . . λk .
A plane partition is a planar analogue of this. A plane partition of a positive integer n is a planar array π of non-negative integers of the form
π1,1
π2,1
..
.

. . . . . . . . . . . . . . . . . . . . π1,λ1
. . . . . . . . . . . . . π2,λ2
.
..

πr,1 . . .

(1.2)

πr,λr

such that entries along rows and along columns are weakly decreasing, i.e.,
πi,j ≥ πi,j+1

and πi,j ≥ πi+1,j ,

P

(1.3)

and such that the sum
πi,j of all entries πi,j equals n. Here, the sequence of row-lengths,
(λ1 , λ2 , . . . , λr ), is assumed to form a partition, i.e., λ1 ≥ λ2 ≥ · · · ≥ λr > 0, and this partition is called the shape of the plane partition π. The individual entries πi,j are called
1

In particular, a plane partition is not a partition of the plane.
1

2

C. KRATTENTHALER

P
parts of π. The sum
πi,j of all parts of the plane partition π — so-to-speak the “size”
of π — i.e., the number that π partitions, will be denoted by |π|. For example, Figure 1
shows a plane partition of 24. That is, if π0 denotes this plane partition, then |π0 | = 24, and its shape is (4, 3, 2).
5 3 3 2
5 1 1
3 1
Figure 1. A plane partition

Figure 2. A pile of cubes
There is a nice way to represent a plane partition as a three-dimensional object: this is done by replacing each part k of the plane partition by a stack of k unit cubes. Thus we obtain a pile of unit cubes. The pile of cubes corresponding to the plane partition in
Figure 1 is shown in Figure 2. It is placed into a coordinate system (the coordinate axes being indicated by arrows in the figure) so that its “back corner” resides in the origin of the coordinate system and it aligns with the coordinate axes as shown in the figure. These piles are not arbitrary piles; clearly they inherit the monotonicity condition on the parts of the original plane partition.

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

3




















If we forget that Figure 2 is meant to show a 3-dimensional object, but rather consider it as graphic object in the plane (which it really is in this printed form), then we realise that this object consists entirely of rhombi with side length 1 (say) and angles of 60◦ and
120◦ , respectively, which fit together perfectly, in the sense that there are no “holes” in the figure which are not such rhombi. By adding a few more rhombi, we may enlarge this figure to a hexagon, see Figure 3.b. (At this point, the thin lines should be ignored).
This hexagon has the property that opposite sides have the same length, and its angles are all 120◦. The upshot of all this is that plane partitions with a rows and b columns, and with parts bounded above by c (recall that parts are non-negative integers) are in bijection with tilings of a hexagon with side lengths a, b, c, a, b, c into unit rhombi, cf. again
Figures 1–3, with Figure 3.a showing the assignments of side lengths of the hexagon. This is an important point of view which has only been observed relatively recently [22]. In particular, it was not known at the time when Richard Stanley studied plane partitions.
Nevertheless, I shall use this point of view in the exposition here, since it facilitates many explanations and arguments enormously. I would even say that, had it been known earlier, much of the plane partition literature could have been presented much more elegantly . . .

b

c





















a

























































































b

c

















a

a. A hexagon with sides a, b, c, a, b, c, where a = 3, b = 4, c = 5

Figure 3.

b. A rhombus tiling of a hexagon with sides a, b, c, a, b, c

4

C. KRATTENTHALER

2. Percy Alexander MacMahon
The problem that MacMahon posed to himself was:
Given positive integers a, b, c, compute the generating function

X

q |π| ,

π

where the sum is over all plane partitions π contained in an a × b × c box!
As explained in the previous section, “contained in an a×b×c box” has to be understood in the sense of the “pile of cubes”-interpretation (see Figure 2) of plane partitions. In terms of rhombus tilings, we are considering rhombus tilings of a hexagon with side lengths a, b, c, a, b, c (see Figure 3), while in terms of the original definition (1.2)/(1.3), we are considering plane partitions of shape (b, b, . . . , b) (with a occurrences of b) with parts at most c.
Why were plane partitions so fascinating for MacMahon, and for legions of followers?
From his writings, it is clear that MacMahon did not have any external motivation to consider these objects, nor did he have any second thoughts. For him it was obvious that these plane partitions are very natural, as two-dimensional analogues of (linear) partitions
(for which at the time already a well established theory was available), and as such of intrinsic interest. Moreover, this intuition was “confirmed” by the extremely elegant product formula in Theorem 1 below. He himself — conjecturally — found another intriguing product formula for so-called “symmetric” plane partitions contained in a given box (see
(6.2)). Later many more such formulae were found (again, first conjecturally, and some of them still quite mysterious); see Section 6 below. Moreover, over time it turned out that plane partitions (and rhombus tilings) are related to many other areas of mathematics, most notably to the theory of symmetric functions and representation theory of classical groups (as Richard Stanley pointed out; see Sections 4 and 7), representation theory of quantum groups (cf. [64]), enumeration of integer points in polytopes and commutative algebra (cf. [12, 79, 17] and the references therein), enumeration of matchings in graphs
(cf. [62, 66]), and to statistical physics (cf. [39, 61]). In brief, the theory of plane partitions offered challenging conjectures (most of them solved now, but see Section 8), connections to many other areas, and therefore many different views and approaches to attack problems are possible, making this a very rich research field.
Returning to MacMahon, the surprisingly simple formula [73, Sec. 429] he found for the generating function for all plane partitions contained in a given box was the following.
P
Theorem 1. The generating function π q |π| for all plane partitions π contained in an a × b × c box is given by2
a Y
b Y
c
Y
1 − q i+j+k−1

1 − q i+j+k−2
i=1 j=1 k=1

.

(2.1)

2MacMahon did not write the result in this form, for which there are many ways to express it.

particular product (2.1) is due to Macdonald [70, Eq. (2) on p. 81].

The

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

We note that, if we let a, b, c → ∞, then we obtain
X
Y
1
q |π| =
, i )i
(1
−
q
π
i≥1

5

(2.2)

an elegant product formula for the generating function for all plane partitions.
MacMahon developed two very interesting methods to prove (2.2) and Theorem 1. First,
MacMahon developed a whole theory, which he called “partition analysis”, and which today runs under the name of “omega calculus.” However, he realised finally that it would not do what it was supposed to do (namely prove (2.2)). So he abandoned this approach. It is interesting to note though that recently there has been a revival of MacMahon’s partition analysis (see [12]), and Andrews and Paule [6] finally “made MacMahon’s dream true.”
The second method consists in translating the problem from enumerating plane partitions to the enumeration of — what MacMahon called — lattice permutations. This idea led him finally to a proof of Theorem 1 in [73, Sec. 494]. The “correct” general framework for this second method was brought to light by Richard Stanley in his thesis (published in revised form as [88]), by developing his theory of poset partitions. See the next section, and the article by Ira Gessel in this volume for an extensive account.
3. The revival of plane partitions and Richard Stanley
As mentioned in the previous section, MacMahon left behind a very intriguing conjecture on “symmetric” plane partitions (these are plane partitions which are invariant under reflection in the main diagonal; see Section 6, Class 2). However, it seems that, at the time, there were not many others who shared MacMahon’s excitement for plane partitions.
(In particular, nobody seemed to care about his conjecture — or perhaps it was too difficult at the time . . . ) In any case, after MacMahon plane partitions were more or less forgotten, except that Wright [100] calculated the asymptotics of the number of plane partitions of n as n tends to infinity. It was Leonard Carlitz [15, 16], and Basil Gordon and Lorne Houten
[35, 36] who, in the 1960s, relaunched the interest in plane partitions.
However, the “rebirth” of intensive investigation of plane partitions was instigated by
Stanley’s two-part survey article [89] “Theory and applications of plane partitions”, together with Bender and Knuth’s article [10]. Already in his thesis (published in revised form as [88]) Stanley had dealt with plane partitions. More precisely, he introduced a vast generalisation of the notion of plane partitions, poset partitions, and built a sophisticated theory around it. In particular, this theory generalised the earlier mentioned key idea of
MacMahon in his proof of Theorem 1 to a correspondence between poset partitions and linear extensions of the underlying poset (again, see the article by Ira Gessel). In [89], he laid out the state of affairs in the theory of plane partitions as of 1971, and he presented his view of the subject matter. This meant, among others, to emphasise the intimate connection with the theory of symmetric functions. Most importantly, the article [89] offered a truly fascinating reading, in particular pointing out that there were several intriguing open problems, some of them waiting for a solution already for a very long time. I am saying all this frankly admitting that it was this article of Stanley which “sparked the fire”

6

C. KRATTENTHALER

for plane partitions in me (which is still “burning” in the form of my fascination for the enumeration of rhombus tilings).
One of the original contributions one finds in [89] is an extremely elegant bijective proof of MacMahon’s formula (2.2), reported from [87] (but which was published later than
[89]). It is an adaption of an idea of Bender and Knuth [10, proof of Theorem 2]. It had the advantage of allowing for an additional parameter to be put in, the trace of a plane partition, which by definition is the sum of the parts along the diagonal of the plane partition. More on this is to be found in Section 5.
With the increased interest in the enumeration of plane partition, several new authors entered the subject, the most prominent being George Andrews and Ian Macdonald. This led on the one hand to a proof of MacMahon’s conjecture by Andrews [1], a proof of a conjecture of Bender and Knuth in the aforementioned article [10, p. 50] by Gordon
[37] and later by Andrews [2], and alternative proofs of both by Macdonald [70, Ex. 16,
17, 19, pp. 83–86]. Moreover, Macdonald introduced another symmetry operation on plane partitions, cyclic symmetry (see Section 6, Class 3), and observed that, again, the generating function for plane partitions contained in a given box which are invariant under this operation seems to be given by an elegant closed form product formula (see (6.4)).

0
1

0

0
0
0


0
1
0 0 0
0 −1
1 0 0

0
1 −1 0 1

1 −1
1 0 0
0
1 −1 1 0
0
0
1 0 0

Figure 4. An alternating sign matrix
Then, William Mills, David Robbins and Howard Rumsey “entered the scene,” and
Stanley had again a part in this development. Robbins and Rumsey [83] had played with a parametric generalisation of the determinant which they called λ-determinant,3 and in that context were led to new combinatorial objects which they called alternating sign matrices.
An alternating sign matrix is a square matrix consisting of 0’s, 1’s and (−1)’s such that, ignoring 0’s, along each row and each column one reads 1, −1, 1, . . . , −1, 1 (that is, 1’s and
(−1)’s alternate, and at the beginning and at the end there stands a 1). Figure 4 shows a
6 × 6 alternating sign matrix. At some point, they became interested in how many such matrices there were. More precisely, they asked the question: how many n × n alternating sign matrices are there? By computer experiments combined with human insight, they observed that, apparently, that number was given by the amazingly simple product formula in (6.12). As Robbins writes in [84], after having in vain tried for some time to prove their
3It is interesting to note that this — at the time, isolated, and maybe somewhat obscure — object has

now become part of a fascinating theory, namely the theory of discrete integral systems (cf. e.g. [23]), and is maybe the first (non-trivial) example of the so-called Laurent phenomenon of Fomin and Zelevinsky [29].

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

7

conjecture, they began to suspect that this may have something to do with the theory of plane partitions. So they called the expert on plane partitions, Richard Stanley. After a few days, Stanley replied that, while he was not able to prove the conjecture, he had seen these numbers before in [3], namely as the numbers of “size n” descending plane partitions, a certain variation of cyclically symmetric plane partitions, which Andrews had introduced in his attempts to prove Macdonald’s conjectured formula for the latter plane partitions.
(Andrews “only” succeeded to prove the q = 1 case of Macdonald’s conjecture.)
This led to further remarkable developments in the theory of plane partitions. Andrews and Robbins independently came up with a conjectured formula for a certain generating function for plane partitions which are invariant under both reflective and cyclic symmetry
(see Section 6, Class 4). Moreover, Mills, Robbins and Rumsey now started to look at alternating sign matrices and plane partitions in parallel, trying to see connections between these seemingly so different objects. One of the directions they were investigating was the study of operations on these objects and whether there existed “analogous” operations on the “other side.” One of the outcomes was the discovery of a new (but, in retrospective, “obvious”) symmetry operation on plane partitions: complementation (see the more detailed account of this discovery in [84]). If combined with the other two symmetry operations — reflective symmetry, cyclic symmetry —, this gave rise to six further symmetry classes, beyond the already “existing” four (see Section 6, Classes 5–10). Amazingly, also for the six new classes, it seemed that the number of plane partitions in each of these six symmetry classes, the plane partitions being contained in a given box, was again given by a nice closed form product formula. Mills, Robbins and Rumsey found striking, much finer enumerative coincidences between plane partitions and alternating sign matrices (see
Section 8), and they initiated a programme of enumerating symmetry classes of alternating sign matrices (in analogy with the programme for plane partitions summarised here in
Section 6). Ironically, they never succeeded to prove any of their conjectures on alternating sign matrices, but their investigations did enable them to provide a full proof [74] of Macdonald’s conjecture on cyclically symmetric plane partitions and a proof of the conjectured formula for another symmetry class of plane partitions (see Section 6, Class 8).
In 1986, Stanley summarised the statements, conjectures, and results (at the time)
for symmetry classes of plane partitions in [90]. In another survey article [91] he also included the statements, conjectures, and results for symmetry classes of alternating sign matrices. So-to-speak, he “used the opportunity” in [90] to resolve one of the (then)
conjectures for symmetry classes of plane partitions, namely the conjecture on the number of self-complementary plane partitions contained in a given box (see Section 6, Class 6).
Stanley’s insight that led him to the solution was that the problem could be formulated as the problem of calculating a certain complicated sum, which he identified as a sum of specialised Schur functions. As Stanley demonstrates, “by coincidence,” this sum of Schur functions is precisely the expansion of a product of two special Schur functions, as an application of the Littlewood–Richardson rule shows; see Section 7.
So, in brief, here is a summary of Stanley’s contributions to the theory of plane partitions:
• poset partitions [88];

8

C. KRATTENTHALER

• trace generating functions for plane partitions [87];
• survey of the theory of plane partitions as of 1971, [89];
• application of symmetric functions to the theory of plane partitions [89, 90];
• surveys of the state of the art concerning symmetry classes of plane partitions and alternating sign matrices as of 1986 [90, 91];
• proof of the conjecture on self-complementary plane partitions [90].
However, Stanley’s contributions extend beyond this through the work of his students, particularly the work of Emden Gansner (who continued Stanley’s work on traces of plane partitions), Ira Gessel (who, together with Viennot, developed the powerful method of non-intersecting lattice paths), Bob Proctor (who proved the enumeration formulae for
Classes 6 and 7 and provided new proofs of several others), and John Stembridge (who, among other things, paved the way for Andrews’ proof of the enumeration formula for
Class 10, and who proved the q = 1 case of the enumeration formula for Class 4). Their work will also be discussed below.
In the next section, I give a quick survey of methods that have been used to enumerate plane partitions, covering in particular the non-intersecting path method. Then, in
Section 5, I shall describe Stanley’s beautiful bijective proof of the formula (2.2) for the generating function for all plane partitions, leading to a refined formula which features the statistic “trace.” This section also contains a brief exposition of the subsequent work of
Gansner on trace generating functions. Section 6 presents the project of enumeration of symmetry classes of plane partitions, listing all cases and those who proved the corresponding formulae. Pointers to further work are also included. This is followed by a detailed discussion of one of the classes in Section 7, namely the class of self-complementary plane partitions. The central piece of that section is Stanley’s proof of the corresponding enumeration formulae. The final section addresses the role of plane partitions in research work of today.
4. Methods for the enumeration of plane partitions
It was said in Section 2 that plane partitions can be approached from many different angles. Which are the methods which have (so far) been applied for the enumeration of plane partitions?
As already explained, MacMahon himself already introduced two methods: first, his partition analysis (nowadays often called “omega calculus”; cf. [12]), which is a generating function method that, in its utmost generality, can be applied for the enumeration of integer points in d-dimensional space obeying linear inequalities and equalities.
Second, MacMahon introduced a translation of counting problems for plane partitions to counting problems of so-called “lattice permutations,” which Stanley [88] generalised to his theory of poset partitions (see again Gessel’s article in the same volume).
The standard method nowadays for the enumeration of plane partitions is to use nonintersecting lattice paths to “reduce” the counting problem to the problem of evaluating

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

9

a certain determinant.4 I want to illustrate this by explaining the non-intersecting paths approach to MacMahon’s formula in Theorem 1. For the sake of simplicity, I concentrate on the q = 1 case, that is, on the plain enumeration of plane partitions contained in a given box. The case of generic q is by no means more complicated or difficult, it would only require more notation, which I want to avoid here.
Consider the plane partition in Figure 3, which is copied in Figure 5.a. We mark the midpoints of the edges along the south-west side of the hexagon and we start paths there, where the individual steps of the paths always connect midpoints of opposite sides of rhombi. See Figure 5.a,b for the result in our running example. We obtain a collection of paths which connect the midpoints of the edges along the south-west side with the midpoints of edges along the north-east side. Clearly, the paths are non-intersecting, meaning that no two paths have any points in common. By slightly deforming the obtained paths, we may place them into the plane integer lattice, see Figure 5.c. What we have obtained is a bijection between plane partitions in an a × b × c box and families (P1 , P2 , . . . , Pb ) of non-intersecting lattice paths consisting of unit horizontal and vertical steps in the positive direction, where Pi connects (−i, i) with (a − i, c + i), i = 1, 2, . . . , b.
The (first) main theorem on non-intersecting lattice paths implies that the number of the families of non-intersecting lattice paths that we have obtained above from our plane partitions can be written in terms of a determinant, see (4.2) below.
Theorem 2. Let G by a directed, acyclic graph, and let A1 , A2 , . . . , An and E1 , E2 , . . . , En be vertices in G with the property that any path from Ai to El and any path from Aj to Ek with i < j and k < l have a common vertex. Then the number of all families
(P1 , P2 , . . . , Pn ) of non-intersecting paths, where Pi runs from Ai to Ei , i = 1, 2, . . . , n, is given by

det LG (Aj → Ei ) ,
(4.1)
1≤i,j≤n

where LG (A → E) denotes the number of all paths starting in A and ending in E.

This theorem was discovered and independently rediscovered several times. It is originally due to Lindström [69, Lemma 1] (who, in fact, proved a more general theorem, and which also includes weights in a straightforward way; for a proof of Theorem 1 with generic q, we would need the weighted version), who discovered and used it in the context of matroid representations. It was Gessel and Viennot [32, 33] who realised its enormous significance for the enumeration of plane partitions.5
4Here, “reduce” is in quotes since, in the harder cases, the most difficult part then is the evaluation of

the determinant.
5Lindström’s theorem was rediscovered (not always in its most general form) in the 1980s at about the same time in three different communities, not knowing of each other at that time: in enumerative combinatorics by Gessel and Viennot [32, 33] in order to count tableaux and plane partitions, in statistical physics by Fisher [25, Sec. 5.3] in order to apply it to the analysis of vicious walkers as a model of wetting and melting, and in combinatorial chemistry by John and Sachs [45] and Gronau, Just, Schade, Scheffler and Wojciechowski [38] in order to compute Pauling’s bond order in benzenoid hydrocarbon molecules.
It must however be mentioned that in fact the same idea appeared even earlier in work by Karlin and
McGregor [47, 46] in a probabilistic framework, as well as that the so-called “Slater determinant” in

10

C. KRATTENTHALER

a. A plane partition

b. Non-intersecting lattice paths

•

•

•

•

•

•

•

•

•

• • • •
..
.
• •.. • • •
•. • •
..
...
. • •
•......•. • • •.. •
..
. ..
.•. • • •.....•... •... •
.
..
..
.. •
..
..
.. ...
•.. • • •.. •......•. •..
.
..
.. ..
.. •..........
.. •.. • •...
•
•
•
..
..
• ...
..
..
.
.
.
.....
.....
• •
• • •. • •. •.
..
..
.
.
• • •
•.....•. • •.. •
..
.
• • • •
•..........
• •. •

•

•

•
•
•
•
•
•
•
•

•

•...........
• •

•

•

•

•

•

•

•
•
•
•
•
•
•
•
•
•
•

c. Non-intersecting lattice paths made orthogonal
Figure 5.

quantum mechanics (cf. [85] and [86, Ch. 11]) may qualify as an “ancestor” of the determinantal formula of Lindström.

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

11

If we apply Theorem 2 with the graph consisting of the integer points in the plane as vertices and edges being horizontal and vertical unit vectors in the positive direction, with
Ai = (−i, i) and Ei = (a − i, c + i), then we obtain that the number of plane partitions contained in an a × b × c box equals the determinant


a+c
.
(4.2)
det
1≤i,j≤b a−i+j
There are many ways to evaluate this determinant, see e.g. [60, Secs. 2.2, 2.3, 2.5], and the result can be written in the form b Y
c a Y
Y
i+j+k−1
i=1 j=1 k=1

i+j+k−2

,

(4.3)

which is the q → 1 limit of (2.1).
Theorem 2 solves the problem of counting non-intersecting lattice paths for which starting and end points are fixed. If one tries the non-intersecting lattice path approach for symmetry classes of plane partitions, then one often encounters the problem that starting or/and end points are not fixed, but rather may vary in given sets. Inspired by work of
Okada [78] on summations of minors of given matrices (which was later generalised to the minor summation theorem of Ishikawa and Wakayama [43]), Stembridge developed the corresponding theory, which provides Pfaffian formulae (and thus again determinantal formulae, given the fact that the Pfaffian is the square root of a determinant) for the number of non-intersecting lattice paths where either starting or end points, or both, vary in given sets. With the exception of only very few, it is the non-intersecting lattice path method which constitutes the first step in proofs of enumeration formulae for symmetry classes of plane partitions, may it be explicitly or implicitly. We refer the reader to Section 7 for an
“implicit” application of non-intersecting lattice paths (in the sense that Stanley’s original proof did not mention non-intersecting lattice paths; they are still there).
Symmetric functions and representation theory come into play because plane partitions are very close to the tableau-like objects that index representations of classical groups. As said earlier, Stanley [89] was the first to emphasise and exploit this connection. It was picked up and deepened by Macdonald [70, Ch. I, Sec. 5, Ex. 13–19], Proctor [80, 81, 82],
Stembridge [93, 94, 96], Okounkov, Reshetikhin, and Vuletić (see [99] and the references therein), and Kuperberg [64]. To give a flavour, let us again consider the enumeration of all plane partitions in an a × b × c box. In terms of the original definition (1.2), these are plane partitions of shape (b, b, . . . , b) (with a occurrences of b) with entries at most c. An example with a = 3, b = 4, c = 6 is shown in Figure 4.a. (The corresponding “graphical”
representations are the ones in Figures 2 and 3.) We rotate the plane partition by 180◦
and add i to each entry in row i, for i = 1, 2, . . . , c. Thereby we obtain an array of integers of the same shape as the original plane partition, however with the property that entries along rows are weakly increasing and entries along columns are strictly increasing, all of them between 1 and a + c. Arrays satisfying the above two monotonicity properties are

12

C. KRATTENTHALER

called semistandard tableaux. See Figure 4.b for the semistandard tableau which arises in this way from the plane partition in Figure 4.a.
5 3 3 2
5 1 1 0
3 1 0 0

1 1 2 4
2 3 3 7
5 6 6 8

a. A plane partition

b. The corresponding semistandard tableau
Figure 6.

It is a central fact of the representation theory of SLn (C) that semistandard tableaux of shape λ (where shape is defined in the same way as for a plane partition) with entries between 1 and n index a basis of an irreducible representation of SLn (C). On the character level, this is reflected by the Schur function sλ (x1 , x2 , . . . , xn ) (cf. [70]), which is defined by sλ (x1 , x2 , . . . , xn ) =

n
XY
T

xi#entries i in T ,

(4.4)

i=1

where the sum is over all semistandard tableaux T of shape λ. Thus, we see that MacMahon’s generating function for plane partitions essentially equals a specialised Schur function, namely we have
X
a+1
(4.5)
q |π| = q −b( 2 ) s(b,b,...,b) (q, q 2 , . . . , q a+c ),
π

the sum on the left-hand side being over all plane partitions π contained in an a × b × c box. If this “principal specialisation” of the variables, i.e., xi = q i , i = 1, 2, . . . , a + c, is done in the Weyl character formula for Schur functions (cf. [70, p. 40, Eq. (3.1)]),
λ +n−j

det1≤i,j≤n (xi j
)
sλ (x1 , x2 , . . . , xn ) =
, n−j det1≤i,j≤n (xi )
then both determinants are Vandermonde determinants and can therefore be evaluated, thus establishing (2.1).
The next important method to enumerate plane partitions is actually a more general method to enumerate perfect matchings of bipartite graphs. A perfect matching of a graph is a collection of pairwise non-incident edges which cover all vertices of the graph. To see how plane partitions can indeed be seen as perfect matchings of certain graphs, in the triangular grid inside the bounding hexagon (see Figure 3.a) place a vertex into the centre of each triangle and connect vertices in adjacent triangles by an edge. In this way, we obtain a hexagonal graph; see Figure 7.a. A unit rhombus is the union of two adjacent triangles. Consequently, a rhombus tiling of a given hexagon corresponds to a collection of edges in the hexagonal graph which are pairwise non-incident and cover all vertices of the graph; in other words: a perfect matching of the hexagonal graph. See Figure 7.b,c for an example.

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

a. The triangular grid with its “dual” graph

b. A rhombus tiling and corresponding edges in the “dual” graph

c. The corresponding perfect matching of the “dual” graph

Figure 7.

13

14

C. KRATTENTHALER

Let G be a given graph. If we consider the defining expansion of the Pfaffian (cf. [94, p. 102]) of the adjacency matrix of G, then each term in this expansion corresponds to a perfect matching of G. Kasteleyn [48, 49] (see [98] for an excellent exposition) showed that for planar graphs there is a consistent way to introduce signs to the 1’s in the adjacency matrix so that the Pfaffian of this modified adjacency matrix (often called Kasteleyn–Percus matrix) has the property that all terms in its expansion have the same sign, and hence gives the number of perfect matchings of G. Since in the case of a bipartite graph, this
Pfaffian reduces to a determinant (namely the determinant of the suitably modified bipartite adjacency matrix of the graph), this approach has been called “permanent-determinant method” by Kuperberg [62, 66]. He has been the first and only one to effectively apply this approach for exact enumeration of plane partitions, providing in particular the first proof for the enumeration of cyclically symmetric self-complementary plane partitions (Class 9 in
Section 6). As a matter of fact, since the permanent-determinant method produces determinants (and permanents and Pfaffians) of very large matrices, this approach is best suited for conceptual insight in the enumeration of plane partitions (and, more generally, perfect matchings of bipartite graphs; see e.g. the last paragraph in Section 8), whereas less so for proving explicit exact enumeration formulae. It is interesting to note, though, that Kuperberg has shown that, for planar graphs, the non-intersecting lattice path method and the permanent-determinant method are equivalent, in a sense explained in [67, Sec. 3.3]. This does not include the situations where the starting or/and end points of the non-intersecting lattice paths to be counted vary in given sets (cf. the paragraph below (4.3)); Kuperberg has however also presented workarounds in [62] to simulate certain such situations within the permanent-determinant method, and he speculates in [67, end of Sec. 3.3] that this is always possible (in the planar case).
Extremely useful tools for the enumeration of plane partitions (and, more generally, perfect matchings) are moreover Kuo’s condensation method [54], which allows for inductive proofs of (conjectured) enumeration formulae, Ciucu’s matchings factorisation theorem [18]
for enumerating perfect matchings of graphs with a reflective symmetry, and Jockusch’s theorem [44] for graphs with a rotational symmetry. As Fulmek [28] (for the former)
and Kuperberg [66] (for the latter two) have shown, all three are consequences of the permanent-determinant method.
Finally, among the purely combinatorially methods that have been applied to the enumeration of plane partition, Robinson–Schensted–Knuth-like correspondences must be mentioned; see [30, 31, 87] and the next section.
5. Trace generating functions
In this section, I present one of the main results from Stanley’s first article [87] on plane partitions.6 There, Stanley introduces the notion of “trace” of a plane partition.
6It

is indeed the chronologically first article, see
Stanley’s publication list http://www-math.mit.edu/~rstan/pubs/ on his website. The order of his papers in the References section at the end of this article follows this order.

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

15

As already said above, the trace ofPa plane partition π as in (1.2) is the sum of its parts along the main diagonal, that is, i≥1 πi,i . He is led to this notion by observing that a combination of a construction of Frobenius [27, p. 523] for partitions with a variation of the
Robinson–Schensted–Knuth correspondence [52] leads to an elegant and effortless proof of
(2.2), and that in this proof a certain statistic, namely the trace, is preserved, leading to a refinement of (2.2) (see (5.5) below).
More precisely, let us consider the plane partition
4
4
2
1
1

3 2 2
3 1 1
2 1 1
1
1

(5.1)

Each row by itself is a(n ordinary) partition. For example, the first row represents the partition 4 + 3 + 2 + 2. A partition λ1 + λ2 + · · · + λk may be represented as a Ferrers diagram, that is, as a left-justified diagram of dots, with λi dots in the i-th row. Thus, the
Ferrers diagram of the above partition is
•
•
•
•

• • •
• •
•
•

(5.2)

The partition conjugate to λ is the partition λ′ corresponding to the transpose of the Ferrers diagram of λ. Thus, the partition conjugate to 4 + 3 + 2 + 2 is 4 + 4 + 2 + 1.
By conjugating each row of (5.1), we obtain
4 4 2 1
4 2 2 1
4 2
2
2

(5.3)

It should be noted that the trace of the original plane partition — in our example in (5.1)
this is 4 + 3 + 1 = 8 — equals the number of parts πi,j with πi,j ≥ i in the plane partition one obtains under this mapping. Moreover, the total sum of the parts of the plane partition is preserved under this operation.
Obviously, each column of the plane partition obtained is a(n ordinary) partition. We now write each column in (essentially) its Frobenius notation. More precisely, out of a partition λ1 + λ2 + · · · + λk , we form the pair
(λ1 , λ2 − 1, λ3 − 2, . . . | λ′1 , λ′2 − 1, λ′3 − 2, . . . ), with λ′ denoting the conjugate of λ, and only recording positive numbers. (In the Ferrers diagram picture, this corresponds to splitting the partition λ along its main diagonal,
“counting” the main diagonal “twice”. In Frobenius’ original notation, the main diagonal is not accounted for.) For example, from the partition 4 + 3 + 2 + 2 in our running example

16

C. KRATTENTHALER

we obtain the pair (4, 2 | 4, 3). We apply this modified Frobenius notation to each column of the plane partition we have obtained so far, and we collect the first components and the second components separately. Thus, for our running example (5.3), we obtain
4 4 2 1
3 1 1
2

5 3 2 2
4 2 1
1

(5.4)

(Here, the pair (4, 3, 2 | 5, 4, 1) corresponds to the first column in (5.3), etc.) It is easy to see that in this way we obtain a pair (C1 , C2 ) of so-called column-strict plane partitions, which are plane partitions with the additional condition that parts along columns are strictly decreasing. Moreover, C1 and C2 have the same shape. It should be observed that the trace of the original plane partition equals the number of parts of C1 (or of C2 ). It is not true that the total sum of the parts of the original plane partition equals the sum of all the parts of C1 and C2 . However, what is true is that the total sum of the original plane partition equals the sum of all parts of C1 and C2 minus the number of parts of C1 .
By a variant of the Robinson–Schensted–Knuth correspondence (see [92, p. 368]), the pairs (C1 , C2 ) correspond to matrices (mi,j )i,j≥1 with a finite number of positive integer entries, all other entries
P being zero. In this correspondence, the total sum of the parts of C1 and C2 equals i,j≥1 (i + j)mi,j (actually, something much finer is true), and the
P
number of parts of C1 equals i,j≥1 mi,j . If one puts everything together, this shows (see
[87, Sec. 3 in combination with Theorem 2.2])
X
π

ttrace(π) q |π| =

1
.
1 − tq i+j−1
i,j≥1
Y

(5.5)

In [87], we also find the notion of “conjugate trace”, together with further generating function results for plane partitions featuring that variation of “trace.”
Almost ten years later, Stanley’s first student, Emden Gansner, showed that there is much more to Stanley’s concept of “trace.” He generalised it and defined the i-trace of a plane
P partition π, denoted by ti (π), as the sum of the parts of the i-diagonal of π, that is, as ℓ≥1 πℓ,ℓ+i . Thus, in this more general context, Stanley’s original trace is the 0-trace.
Gansner’s result on trace generating functions in [31] actually concerns reverse plane partitions, which are arrays of non-negative integers as in (1.2) with the property that entries along rows and along columns are weakly increasing. For the formulation of the result, we need to define the monomial x(ρ; λ) for a “cell” ρ of a partition λ (a cell corresponds to a bullet in the Ferrers diagram of λ; see (5.2)). If ρ is located in the i-th row and the j-th column of λ, then x(ρ) := xj−λ′j · · · xλi −i−1 xλi −i ,

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

17

where λ′ is again the partition conjugate to λ (see the paragraph after (5.2)). Using this notation, Gansner’s elegant result [31, Theorem 5.1] is
1 −1
X λY

π

i=1−ℓ(λ)

t (π)

xii

=

1
,
1 − x(ρ)
ρ∈λ

Y

(5.6)

where the sum is over all reverse plane partitions π of shape λ. The proof is bijective; it exploits hidden properties of a correspondence of Hillman and Grassl [42] that Gansner lays open. He continued that work in [31], where he considers trace generating functions for
(ordinary) plane partitions. As it turns out, the trace generating function (with all traces present) for plane partitions of a given shape does not evaluate to an equally elegant product formula as for reverse plane partitions (except for rectangular shapes, which however follows trivially from the result (5.6) for reverse plane partitions by a 180◦ degree rotation), one has to be content with formulae given in terms of certain sums. In order to derive these results, Gansner uses Robinson–Schensted–Knuth-like correspondences due to Burge [13].
It is shown in [55, Theorems 7.6, 7.7] that, for generating functions for plane partitions which only take into account the size and 0-trace of plane partitions, there are closed form products available for certain shapes.
Finally, we mention that Proctor [82] has derived several results for alternating trace generating functions by relating these to characters of symplectic groups.
6. The 10 symmetry classes of plane partitions
The purpose of this section is to summarise the programme of enumeration of symmetry classes of plane partitions, the history of which was reviewed in Section 3. I start by defining the symmetry operations which give rise to 10 symmetry classes of plane partitions. Then, for each symmetry class, I state the corresponding enumeration formula, and report who had proved that formula, sometimes with pointers to further work.
The first operation is reflection. To reflect a plane partition π as in (1.2) means to reflect it along the main diagonal, that is, to map π = (πi,j )i,j to π = (πj,i )i,j . In the representation of a plane partition as a rhombus tiling as in Figure 3.b, reflection means reflection along the vertical symmetry axis of the hexagon (if there is one).
Rotation is the second operation. Given a plane partition π, viewed as a pile of cubes as in
Figure 2, to rotate it means to rotate it by 120◦ with rotation axis {(t, t, t) : −∞ < t < ∞}
(this is the rotation which leaves the origin of the coordinate system invariant and maps the coordinate axes to each other; for our purposes it is irrelevant whether we consider a
“left” or “right” rotation). In the representation of a plane partition as a rhombus tiling as in Figure 3.b, this rotation corresponds to a rotation of the hexagon (together with the tiling) by 120◦ .
Intuitively, the complement of a plane partition in a given a × b × c box is what the name says, namely the pile of cubes which remains (suitably rotated and reflected) after the plane partition is removed from the box. To be precise, if the cubes of a plane partition π
contained in an a×b×c box are coordinatised in the obvious way by (i, j, k) with 1 ≤ i ≤ a,

18

C. KRATTENTHALER

✧❜
✧
❜
✧❜
✧❜
✧
❜
✧
❜
✧
❜
✧
✧ ❜❜
✧
✧ ❜
❜
✧
❜
❜✧ ❜✧ ❜✧
✧❜
✧❜ ✧
✧❜
❜✧✧❜❜
✧
✧
❜ ❜
✧
❜
❜✧ ❜❜✧ ❜✧✧ ❜✧
❜✧
✧❜
✧❜ ✧
✧❜
✧
❜✧✧❜❜
✧
❜ ❜
✧
❜
❜✧
❜✧
❜✧
✧❜
✧
❜
✧
❜
✧
✧❜
✧❜
✧❜
❜ ✧❜ ✧
❜✧✧❜❜
✧
✧ ❜✧ ❜
✧
❜ ❜
✧
❜
✧ ❜✧✧ ❜✧
❜✧ ❜❜✧ ❜✧✧❜❜
❜✧
❜✧
✧❜
✧❜
✧❜ ✧
❜ ✧
❜
❜ ✧
❜
✧ ❜✧
✧
❜✧
❜✧ ❜
❜✧ ❜✧
✧❜
✧
❜
✧
✧
❜
❜
✧
❜ ✧
✧❜ ✧❜ ✧❜ ✧❜
❜
✧
✧
✧
✧ ✧❜
❜
❜
❜
❜
✧❜
✧ ❜
❜
✧
❜ ❜✧ ❜✧ ❜✧ ❜✧ ✧
❜ ✧❜ ✧❜ ✧❜ ✧
❜✧ ❜✧ ❜✧ ❜✧

✧❜
❜
✧
✧❜
✧❜
❜✧
❜
✧
✧❜ ✧
❜
✧ ❜❜
❜
✧ ✧
✧
✧❜
❜
✧ ❜✧ ❜✧ ❜✧ ❜
✧
❜
✧❜
✧❜
✧❜❜✧
✧
❜
❜✧✧❜❜
❜
✧
✧
❜
✧
❜
✧
✧❜
✧❜
✧ ❜✧ ❜✧ ❜✧ ❜✧ ❜
❜✧
✧
❜
✧❜
✧❜ ✧
✧❜
❜✧✧❜❜
✧
❜ ❜
✧
❜
✧
✧
❜
❜✧
❜✧ ❜
✧ ❜✧
✧❜
✧❜
❜
✧
✧
❜
❜
✧
✧
✧
❜
❜
✧
✧❜
❜
✧ ❜✧
❜
✧
❜✧ ❜
✧ ❜✧ ❜
❜
✧
✧ ❜✧✧ ❜✧
❜✧ ❜❜✧ ❜✧✧❜❜
❜✧
❜✧
❜ ✧
✧❜ ✧
✧❜
✧❜
❜
❜ ✧
✧ ❜✧
❜
✧
❜✧ ❜✧
❜✧ ❜
❜✧
❜
✧
✧❜
✧
✧
❜
❜
❜ ✧
✧❜ ✧❜ ✧❜
❜
✧
❜✧✧
✧
✧
✧
❜
❜
❜
❜
❜
✧
❜
✧
❜✧ ❜✧ ❜✧ ❜✧
❜
✧
❜ ✧❜ ✧❜ ✧❜ ✧
❜✧ ❜✧ ❜✧ ❜✧
✧❜ ✧❜
✧
❜ ✧
❜❜ ❜
✧
❜ ✧✧
✧
❜
✧
❜
❜✧

Figure 8. A totally symmetric self-complementary plane partition
— as a pile of cubes, and as a rhombus tiling of a hexagon
1 ≤ j ≤ b, and 1 ≤ k ≤ c, then the complement π c of π is defined by
π c = {(a + 1 − i, b + 1 − j, c + 1 − k) : (i, j, k) ∈
/ π}.
In the representation of a plane partition as a rhombus tiling as in Figure 3.b, to take the complement of a plane partition means to rotate the hexagon (together with the tiling) by
180◦ .
If one combines the three operations — reflection, rotation, complementation — in all possible ways, then there result ten symmetry classes of plane partitions. Remarkably, in all ten cases, there exist closed form product formulae for the number of plane partitions contained in a given box, respectively for certain generating functions. The first complete presentation of the corresponding results and conjectures was given by Stanley in [90], together with the state of the art at the time, and the proof of the conjectured formulae for Class 5 (see also Section 7).
While reading the descriptions of the various symmetry classes, Figure 8 may be helpful.
It shows a plane partition which has all possible symmetries. Thus, it belongs to all ten classes.
In order to have a convenient notation to write down some of the formulae, we use the symbol Nd (a, b, c) for the number of all plane partitions in symmetry class d which are contained in an a × b × c box.
Class 1: Unrestricted Plane Partitions. As we already said, MacMahon proved (see
Theorem 1) that the generating function for all plane partitions contained in an a × b × c

PLANE PARTITIONS IN THE WORK OF RICHARD STANLEY AND HIS SCHOOL

19

box is given by
X

q

|π|

π

=

a Y
b Y
c
Y
1 − q i+j+k−1

1 − q i+j+k−2
i=1 j=1 k=1

,

(6.1)

where the sum is over all plane partitions π contained in an a × b × c box.
Class 2: Symmetric Plane Partitions. A plane partition π as in (1.2) is called symmetric if it is invariant under reflection along the main diagonal. In the representation of a plane partition as a rhombus tiling as in Figure 3.b, to be symmetric means to be invariant under reflection along the vertical symmetry axis of the hexagon (two of the side lengths of the hexagon must be equal).
For this class, two different weights lead to generating functions which have closed form product formulae. The first weight is the usual “size” of a plane partition, and it leads to
MacMahon’s conjecture [72] that
X
π

q

|π|

=

a
Y
1 − q c+2i−1
i=1

1 − q 2i−1

1 − q 2(c+i+j−1)
,
1 − q 2(i+j−1)
1≤i<j≤a
Y

(6.2)

where the sum is over all symmetric plane partitions π that are contained in an a × a × c box. MacMahon’s conjecture was first proved by Andrews [1] and Macdonald [70, Ex. 16
and 17, pp. 83–85]. Since then, several other proofs and refinements were given, see [81,
Prop. 7.3], [82, Theorem 1, Case BYI], [63, Sec. 5], [56, Cor. 12], [57, Theorem 5].
The second weight is roughly “half” the size of a plane partition, and it leads to a conjecture of Bender and Knuth [10, P
p. 50]. Given a plane partition π as in (1.2), the weight |π|0 of π is defined as the sum 1≤j≤i πi,j . Then
X

q |π|0 =

π

1 − q c+i+j−1
, i+j−1
1
−
q
1≤i≤j≤a
Y

(6.3)

where the sum is over all symmetric plane partitions π that are contained in an a × a × c box. Bender and Knuth’s conjecture was first proved by Gordon [37] (as reported in [89,
Prop. 16.1]), but published only much later. The first published proofs are due to Andrews
[2] and Macdonald [70, Ex. 19, p. 86]. Since then, several other proofs and refinements were given, see [81, Prop. 7.2], [82, Theorem 1, Case BYH], [56, Theorem 21], [57, Theorem 6],
[24].
Class 3: Cyclically Symmetric Plane Partitions. A plane partition π is called cyclically symmetric if, viewed as a pile of cubes as in Figure 2, it is invariant under rotation by 120◦
as described above. In the representation of a plane partition as a rhombus tiling as in
Figure 3.b, to be cyclically symmetric means to be invariant under rotation of the hexagon
(together with the tiling) by 120◦ (all sides of the hexagon must have the same length).
For this class, Macdonald [70, Ex. 18, p.85] conjectured
X
π

q

|π|

=

a
Y
1 − q 3i−1

1 − q 3(2i+j−1) Y 1 − q 3(i+j+k−1)
,
1 − q 3i−2 1≤i<j≤a 1 − q 3(2i+j−2) 1≤i<j,k≤a 1 − q 3(i+j+k−2)
i=1
Y

(6.4)

20

C. KRATTENTHALER

where the sum is over all cyclically symmetric plane partitions π that are contained in an a × a × a box.
Andrews [3, Theorem 4] had found a determinant for the generating function in question; in retrospective this determinant can be explained by non-intersecting lattice paths (which have been discussed in Section 4). The conjecture was proved by Mills, Robbins and
Rumsey [74] by evaluating this determinant using some beautiful linear algebra, and it is still the only proof.
Class 4: Totally Symmetric Plane Partitions. A plane partition is called totally symmetric if it is at the same time symmetric and cyclically symmetric.
For this class, Andrews and Robbins (see [4]) conjectured
X
Y
1 − q i+j+k−1
q |π|0 =
,
(6.5)
i+j+k−2
1
−
q
π
1≤i≤j≤k≤a

where the sum is over all totally symmetric plane partitions π that are contained in an a × a × a box.
Okada [78, Theorem 5] found a determinant for the generating function in question.
However, for a long time nobody knew how to evaluate this determinant. In the q = 1
special case, Stembridge [97] was able to relate a slightly different determinant to the enumeration of cyclically symmetric plane partitions, thereby solving the problem of plain enumeration of totally symmetric plane partitions. Unfortunately, it seems that this very conceptual approach does not extend to the q-case. This was the longest standing conjecture of all plane partition conjectures. Finally, Koutschan, Kauers and Zeilberger [53]
succeeded to evaluate Okada’s determinant, by a (heavily) computer-assisted argument.
Class 5: Self-Complementary Plane Partitions. A plane partition is called self-complementary if it is equal to its complement. In the representation of a plane partition as a rhombus tiling as in Figure 3.b, to be self-complementary means to be invariant under rotation of the hexagon (together with the tiling) by 180◦ . In other words, to be selfcomplementary in the rhombus tiling picture means to be centrally symmetric.
For this class, Robbins and Stanley independently conjectured that
N5 (2a, 2b, 2c) = N1 (a, b, c)2 ,

(6.6a)

N5 (2a + 1, 2b, 2c) = N1 (a, b, c)N1 (a + 1, b, c),

(6.6b)

N5 (2a + 1, 2b + 1, 2c) = N1 (a + 1, b, c)N1 (a, b + 1, c).

(6.6c)

As mentioned above, Stanley proved this conjecture in [90, Sec. 3], see Section 7. For a different proof see [63, S

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
