---
source: "https://archive.org/details/arxiv-1111.6013"
title: "Direct embeddings of relatively hyperbolic groups with optimal $\\ell^p$ compression exponent"
author: "David Hume"
year: "2012"
captured_at: "2026-08-10T19:21:55Z"
updated_at: "2026-08-10T19:21:55Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "デイヴィッド・ヒューム"
query: "David Hume"
plain_text_url: "https://archive.org/download/arxiv-1111.6013/1111.6013_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "経験論"
  - "懐疑主義"
status: raw
---

# Direct embeddings of relatively hyperbolic groups with optimal $\ell^p$ compression exponent

- 著者: David Hume
- 初版: 2012
- 情報源: [archive](https://archive.org/details/arxiv-1111.6013)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[デイヴィッド・ヒューム]]
- 研究動向: [[デイヴィッド・ヒューム-現代研究動向]]

## Full Text

Direct embeddings of relatively hyperbolic groups
with optimal £ p compression exponent

David Hume*

Abstract

We prove that for all p > 1, every relatively hyperbolic group has l p com-
pression exponent equal to the minimum of the exponents of its maximal
parabolic subgroups.

1 Introduction and statement of results

Coarse embeddings of discrete metric spaces into Banach spaces are an impor-
tant tool in geometric group theory, combinatorics and if -theory. They have
a strong connection to expander graphs, |LLR95j . |HLW06] . while e mbedding s
into low dimensional spaces are of much interest in computer science, BDG + 05] .

Gromov suggested |Gro93] and later Yu proved [YuOOj that any group admit-
ting a coarse embedding into a Hilbert space satisfies the Novikov conjecture.
Kasparov and Yu |KY06| then extended this result to any group admitting a
coarse embedding into any uniformly convex Banach space. If, in addition to
admitting a uniform embedding into Hilbert space, a discrete metric space has
bounded geometry, then it also satisfies the coarse Baum-Connes conjecture,
[YuOOj . An introduction to these conjectures can be found in [BCH94 .

As a first step, we recall the definition of a coarse embedding.
Definition 1.1. Coarse embeddings

Let X be a discrete metric space and let Y be any metric space.

A map (j) : X —> Y is called a coarse embedding if there are two functions

p± :M> ^M> ,

such that p- (r) — > oo as r — > oo and

p- (d x {xi,x 2 )) < d Y ((f>(xi),(f)(x2)) < p + (dx(xi,X 2 )) ■

Often such embeddings are called 'uniform', but this term means different things
to different areas so we settle instead on the term 'coarse'. By a 'Lipschitz coarse
embedding' we mean a coarse embedding with p+(n) < Cn and often abbre-
viate the compression function p_ simply as p. Gromov, GroOO , proved the
existence of a finitely generated group which does not admit a coarse embedding

"University of Oxford, Mathematical Institute, 24-29 St Giles, Oxford OX1 3LB, United
Kingdom. Email: hume@maths.ox.ac.uk

1

into any Hilbert space, for more details see |AD) . It is unknown whether such
a group can be coarsely embedded into some uniformly convex Banach space.

More recently, Guentner and Kaminker GKQ4] introduced compression expo-
nents to distinguish between groups admitting coarse embeddings, dependent
on a form of Holder equivalence. To clarify notation, given two functions /
and g, we will write / ^ g to mean there exists some constant C such that
f(n) < Cg(n) + C, and write / xj when f < g and g -< f.

Definition 1.2. (Equivariant) compression exponent

Let X be a discrete metric space and let Y be a Banach space. The compression
exponent, a Y (X) is defined to be the supremum of those values a 6 [0,1] for
which there exists some Lipschitz coarse embedding <f> : X — > Y with

p{n) h n a .

In the case where the metric space is the Cayley graph of a finitely generated
group G, the equivariant compression exponent cty(G) is the same supremum
but where each coarse embedding must also be equivariant, so there exists some
isometric action of G onY such that <j){g) = g(0y).

WhenY = £P(N), a* Y {X) is denoted by a*{X) anda*{X) is denoted by a*(X).

We note that a* is a quasi-isometry invariant and therefore is a quasi-isometry
invariant of the class of groups. It is not known whether a& is a quasi-isometry
invariant, however it is invariant up to choice of finite generating set so the
equivariant compression exponent of a finitely generated group is well-defined.

In the case where X is the Cayley graph of a finitely generated group, these
values are closely related to forms of amenability. In fact, for amenable groups G,
a^G) = af(G), this was first proved in the abelian case by Aharoni, Mityagin
and Murray, (AMM85] . and then generalised by Gromov. For a full proof, see
jdCTV071 Proposition 4.4]. By contrast, any hyperbolic group with Kazhdan's
property (T) has af(G) = and a*(G) = 1.

Any finitely generated group with af(G) > is a-T-menable, moreover, if
af(G) > \ then G is amenable |GK04| . It was conjectured - with strong
support from the work linking equivariant compression exponents with the speed
of random walks |NP08| . [LP] - that every amenable group G has af(G) > L
However, Austin |Ausj answered this in the negative by displaying many finitely
generated solvable groups with af(G) = 0. Currently there are no known
examples of amenable groups G with af(G) <E (0, \).

In the non-equivariant setting a* 2 (G) > h implies that G is exact and hence has
Yu's property (A) ( GK04 . YuOOj ) and there exists a family of finitely generated
groups {G a } ae \o ^ of asymptotic dimension at most 2 with ay(G a ) — a for all
uniformly convex Banach spaces Y }ADS09| .

Again, searching for a more sensitive invariant, we can ask for bounds (up to
multiplicative and additive constants) on the functions p that it is possible to
achieve as lower bounds for Lipschitz coarse embeddings. Though all finitely
generated polycyclic groups G have compression exponent 1, [Tesllj . no such G

2

of exponential growth admits a quasi-isometric embedding into a Hilbert space,
[dCT08j .

In [Gro87j . Gromov introduced relatively hyperbolic groups as a generalisation
of hyperbolic groups. These groups have many different characterisations: in
terms of group actions Bow99 , group-theoretic structure |Far98] . [Dah03b .
|Osi06j . topology |Yam04| and metric geometry [DS05] . This paper focuses
its attention on metric relative hyperbolicity, i.e. asymptotically tree-graded
spaces.

The class of relatively hyperbolic groups includes:

(i) hyperbolic groups,

(ii) amalgamated products and HNN-extensions over finite subgroups,

(iii) limit groups (fully residually free groups),

(iv) geometrically finite Kleinian groups,

(v) fundamental groups of finite volume hyperbolic manifolds,

(vi) fundamental groups of non-geometric closed 3-manifolds with at least one
hyperbolic component.

We mention also that the Teichmuller space (equipped with the Weil-Petersson
metric) corresponding to surface £ = with (g,n) £ {(2,0), (1,3), (0,6)} is
asymptotically tree-graded, }BM08j .

To apply a little perspective we present an overview of results concerning these
three invariants for relatively hyperbolic groups (and sufficiently similar metric
spaces).

Simplicial trees of uniformly bounded valency have compression exponent 1,
[GK041 Proposition 4.2], a bounding function

n

p(n) h 1 ■

log2(n + 2)loglog(n + 2)

was obtained in [BS08] , before Tessera |Tesllj proved that given any countable
simplicial tree there is a Lipschitz coarse embedding into a £ p space (p 6 [1, oo))
with p y f for any increasing function / satisfying property (C p ).

In particular, for all e > and all p,

(log(n + 2) loglog 1+£ (n + 2))p

has property (C p ).

Gal Gal08 proved this independently and extended it to all spaces with finite
Assouad-Nagata dimension.

Building on the work of Bourgain, [ Bou85j . both authors prove that this tree

3

embedding result is optimal when p > 2. It is certainly not for p = 1 as any
countable metric tree (any weighting of the edges of a countable simplicial tree
by finite values) can be isometrically embedded into ^(E) where E is the set
of edges of the underlying simplicial tree.

Moving on to hyperbolic groups, any finitely generated hyperbolic group G
quasi- isometrically embeds into a finite product of binary trees BDS07]. so
such groups have Hilbcrt compression exponent o^G) = 1. Previous to this,
Bonk and Schramm [BSOQ] Theorem 1.1] found quasi-isometric embeddings of
any hyperbolic metric space X with bounded growth at some scale into some
hyperbolic space H™. (We recall that bounded growth at some scale means
there exist constants R > r > and meN such that any ball of radius R in X
can be covered by m balls of radius r.) As we may quasi-isometrically embed
hyperbolic space into a finite I 1 product of (possibly infinite valence) simplicial
trees BS05, Theorem 1.1] for each increasing function / : N — » K>o satisfying
(C p ) a Lipschitz coarse embedding of X into an HP space with ph f [Tesllj .

Considering relatively hyperbolic groups, coarse embeddability results were first
obtained for free products of coarsely embeddable groups, CDGY03 , free prod-
ucts of coarsely embeddable groups with non-trivial amalgamation IDG03j and
then for all relatively hyperbolic groups whose maximal parabolic subgroups
admit coarse embeddings [DG07J. Dreesen, [DrelOj . proves that given finitely
generated groups A, B and G, where G is a finite subgroup of A and B,

Within this collection, however, Brodskiy and Higes, [BH , prove that if two
finitely generated groups A and B have finite Assouad-Nagata dimension, then
so does A* B. Hence A* B has Lipschitz coarse embeddings with p >z f for any
increasing function / with property (G p ) |Gal08] .

Staying within this framework of spaces with a tree-like structure every finite
dimensional CAT(O) cube complex has Hilbert compression exponent 1, |CN05| .
generalising the result for finite products of simplicial trees. Coxeter groups
embed isometrically (and equivariantly) into an i 1 product of finitely many
trees |Uj99) , and as each right-angled Artin group is commensurable with some
Coxeter group |DJ00] , any Coxeter or right-angled Artin group admits Lipschitz
coarse embeddings into some l v space with p >: / for any increasing function /
with property (C p ). Additionally, by [Wis] Theorem 17.4], limit groups (finitely
generated fully residually free groups) are virtually undistorted subgroups of
right-angled Artin groups, so the result holds for these groups as well. Moreover,
such groups admit quasi-isometric embeddings into some I 1 space. Limit groups
are relatively hyperbolic with respect to abelian groups [AliQSl , so this is another
collection of relatively hyperbolic groups for which the embedding question is
answered.

Before presenting the results of this paper we make a useful definition which will
characterise the functions we obtain as lower bounds. It is essentially Tessera's
property, but with some local assumption which is necessary for our direct
methods.

muxia* 2 {A),a^(B),-

< a* 2 (A* c B) < v3xa{al(A),al(B)} and

4

Definition 1.3. Concave functions and property (Cp)

We will call a function f : N —t R>o concave if f is non- decreasing and for all
m, n S N with n > m:

f(n + m) - f(n) < f(n) - f[n - m).

This is modelled on the usual concavity condition f" < given for smooth func-
tions.

A concave function satisfying (C p ) is said to satisfy (Cp) if ^ is non-
decreasing for all n sufficiently large.

We observe here that for all e > and all p > 1

f(n) = r

(log 2 (n + 2)(log 2 log 2 )i+ e (n + 2))?

has property (Cp).

Bearing in mind [Teslll Proposition 7.5], the following question is clearly rele-
vant.

Question 1.4. Is there a sub-linear function g such that for all f with property
(Cp) there exists a constant C > such that f(n) < Cg(n) + C?

The obvious candidate would be g(n) = n(log(n + 2))~ p .
The main results of this paper is the following.
Theorem 1. (cf. EM

Let X be the -skeleton of an asymptotically tree-graded simplicial graph with
bounded vertex degree. Let the collection of pieces be A — {Ai \ i £ I } and let
p > 1. Suppose we have a concave function p' : M>o — > K>o and 1-Lipschitz
embeddings ipi : Ai — > £ p (Xi) such that for all x, y € IY,

p'(d x (x,y)) < \\Mx)-My)\\ P -

Then for every function f : N — > M>o with property (C^), there is a Lipschitz
coarse embedding (j) from X to an £ p space such that p(n) >z min{/(n), p'(n)}.

This yields the result promised by the title.

Corollary 2. (cf. EM

Let G be a finitely generated group which is hyperbolic relative to a collection of
subgroups {Hi}. Given any p > 1, any collection of Lipschitz coarse embeddings
of the Hi into £ p spaces with associated concave compression functions pi and
any function f with property (Cp) there is a Lipschitz coarse embedding of G
with

p(n) y mm {p t (n),f(n)} .
So, for all p > 1, a*(G) — mm {a*(7Ji)} .

Proof: Every relatively hyperbolic group is the 0-skeleton an asymptotically
tree-graded simplicial graph DS05 , where each piece is quasi-isometric to one
of the Hi, which are undistorted. Then the result follows from theorem [1] □

In particular, we obtain an embedding result for all closed 3-mamfolds.

5

Corollary 3. Let M be a closed 3-manifold, then for all p > 1 and all f

satisfying property (C"), there exists a Lipschitz coarse embedding ttx(M) into
some i p space, such that

Phf-

Proof: Consider first the geometric manifolds. The fundamental groups of
these are quasi-isometric to one of eight Thurston geometries, of which only Nil
and Sol need worrying about. Both are polycyclic, locally compact and com-
pactly generated, so we are done by Tessera's result, [Tesllj , theorem 1.

In the non-geometric case, we decompose the manifold along tori using the
Geometrisation Theorem, (p5rQ2], |Per03j . |CZ06aj . |CZ06bj . [KL08] . |MT07| .
MT08]). If M has no hyperbolic part then it is a graph manifold and Smirnov,
jSmilOj . proves this has finite Assouad-Nagata dimension. Applying Gal's re-
sult |Gal08j gives for each / with Tessera's property (C p ), a Lipschitz coarse
embedding of such a space with

Phf-

Finally, if it has a hyperbolic part, then tti(M) is hyperbolic relative to the
fundamental groups of a finite collection of graph manifolds, tori and Klein
bottles, [Dah03a|. The fundamental group of the Klein bottle is polycyclic, so
using Tessera's result again and applying corollary [5] completes the result. □

By way of complete contrast, Sapir |Sapll| proves the existence of a closed
aspherical 4-manifold M where tti(M) coarsely contains expanders and hence
admits no coarse embedding into any Hilbert space. This uses Gromov's proof
[GroOO] of the existence of a group coarsely containing expanders.

We also obtain an estimate for LP compression.

Corollary 4. Let X be an asymptotically tree-graded simplicial graph of
bounded degree and let {Ai \ i E 1} be a suitable choice of pieces. Suppose we
are given a collection of Lipschitz coarse embeddings ipi : Ai — ► L p ([0, 1]) and a
concave function p' such that for all i E L and all x, y G A4,

p'{d(x,y)) > ||^(a:)-Vi(l/)||p.

For each function f satisfying property (C£) where q — max{p, 2} there exists a
Lipschitz coarse embedding cj> of X into L P ([0,1]) with

P(d(x,y)) t ™n{f(d(x,y)),p'(d(x,y))}.

Proof: For any p 6 [l,oo) and any countable set Y we may embed £ P (Y)
into L p ([0, 1]) as follows. Enumerate Y — {yo, yi, 2/2, • ■ • } and define ipy on the
basis of Dirac unit vectors {e y } as follows.

2 n+1 if z e (2-",2-(" +1 )]
otherwise.

The remainder then follows by recalling the fact that L 2 ([0, 1]) isometrically
embeds into £ p ([0, 1]) when p £ [1, 2], |Woj91| and applying theorem[TJ □

As a useful introduction to the techniques required in this paper, we first present
proofs in the same guise for uniformly discrete hyperbolic spaces and tree-graded
spaces, before embarking on the proof of theorem [I] which is more technical.

6

2 Hyperbolic metric spaces

Here we will provide a short self-contained method of embedding uniformly
discrete hyperbolic metric spaces into l v spaces, which is close to optimal when
P > 2.

As an aid to notation in this section we will define [a;, yj to be the set of all
geodesies from x to y in a given metric space.

We first require the following lemma.

Lemma 2.1. Let X be the 0-skeleton of a 6 -hyperbolic graph and let e £ X .
Then there exists some constant K = K(8), such that for all n > 3(5, for all
x, y £ X with d(x, e) > n and d(x, y) < j, for all geodesies go £ [x, ej, g £ [y, e]
and p £ g([n 7 2n]),

d(p,go([0,3n])) <K.

Proof: We use the Rips definition of hyperbolicity, so in a geodesic triangle
any edge is contained in the union of the 5 neighbourhoods of the other two.
Select p £ g([n, 2n]), if p lies within the S neighbourhood of go then we are done
as a sufficiently close point must lie within the required range.

Alternately, if p does not lie within the 6 neighbourhood of go, then it must lie
within the 6 neighbourhood of any geodesic in [x, y].

Let z be a point on some geodesic in [x, y\ with d(p, z) < 5, then

n

d{x,p) < d(x, z) + d(z,p) < — + S.

However,

n

d{x,p) > d{y,p) - d{x,y) > n - -,

which is a contradiction as n > 35. □

Theorem 2.2. Let X be a countable uniformly discrete 8-hyperbolic metric
space with bounded geometry. Then for each concave function f with property
(C p ) there exists a Lipschitz coarse embedding <f> : X — > ©„ gN ^ p (^0 with

Phf-C.

Proof: We can reduce our problem to the case where X is the 0-skeleton
of a connected simplicial graph, using [TuOll Lemmas 4.1 and 7.3]. As X has
bounded geometry we can define N(k) to be a bound on the cardinality of any
ball of radius k.

Fix a basepoint e £ X. We define the set of restricted geodesies

G x , k ,n = |J {£([«, 2n])}.
d(x, y) < k
9 G [y, e]

We define F x ^, n to be the rest of all points in X lying on some g £ G x ^k,n but
not in B 3 s{e) and set F(x, k 7 n) to be the characteristic function of F x ^, n - We
use the bounded geometry of X to ensure F(x, k, n) £ l p (X).

7

Next define

H(x,n) — — y F(x,k,n).
n

k<§

The following three lemmas provide bounds on the p-norms of these functions.

Lemma 2.3. There exists some constant C such that for all x £ X , k < j
and n £ N \ {0} with d(x, e) > 2n,

n-36 < \\F(x, k,n)\\l < Cn.

Proof: Notice ||F( )|r = 11^(35, k, n)\\ 1 as this is a characteristic function.

The first inequality is obvious as l-F^k,™! > n — 38. For the second we use lemma
12.11 and the bounded geometry of X ,

||F(a;,*,n)lli < (n + l)N(K(S))

< 2N(K(S))n < Cn.

This upper bound does not rely on the fact that d(x, e) > 2n.

Lemma 2.4. If d(x,y) < R then \\H(x,n) - H(y,n)\\ < 2C(R + l)^^ .

Proof: Choose x,y £ X with d(x,y) < R. Then as F x ,k, n C F y ^+R,n,

^2 F(x,k,n) < F(x,k,n)+ ^ F(y,k + R,n)

0<fc<f f-i?<fc<f 0<fc<f-i?,

< Yj F(x,k,n)+ F (v> k > n )-

f-i?<fc<f 0<fc<f

Switching x and y in the above argument we conclude that

\\H(x,n)-H(y,n)\\ p p < ^ £ \\F(x,k,n)\f p + \\F(y,k,n)\f p

< \{2C{R + l))n< (2C(R + l))n-(P-^

as required. Notice we have made no assumption that Hi(x, n), Hi(y, n) ^ 0.
Lemma 2.5. \\H(x, n) \\ p p x n, whenever d(x, e) > 2n.

Proof: For the lower bound on n)\\ p we notice that given any fixed

geodesic g £ {x, e],

g([n, 2n]) C F(x, k, n) for all k.

Hence, the function H(x,n) has at least n — 36 points on which it takes value
at least j, so the lower bound is justified.

As an upper bound,

\\H(x,n)\\ p < n- 1 Y \\F(x,k,n)\\ p < n' 1 Q + l) (Cn)* * nk

8

With these three lemmata we are now in a position to define our embedding

4>{x) = Y J f -^p-H{x,2 n ).

n>l

To show <p is Lipschitz, consider with d{x, y) < R. Then, using lemma

\2~H

/(2T

n=l
oo

* E

/(2T P

2"

But then we notice that as / is concave f{Tf < 2 P+1 ^ -J{i) p

n=l

2 »+i

i=2 n +l

So

00 / t(o n \ \ p 00 2 " +1 1

U{x) - 4>{y)\\ p P < E Rrr 1 ^E 2 ^ E 7/« p

n=l ^ ' n=l i=2" + l

* E- Or 2 ^>

2 ,

i=l x '

as / has property (C p ).

For the lower bound on </>, consider two distinct points X, with at least

one of d(x, e), d(y, e) > 6(5 and compute their Gromov product (x,y) e . Set
Mr = [iog 2 (d(x,e) - {x,y) e )\ and fc a = [\og 2 (d{y, e) - (x,j/) e )J.

Doing this we obtain a constant c = c(<5) such that

c^d(x, y) < 2 k * + 2 k y < cd(x, y).

Thus we can see that d(x,y) x max{2 fcx , 2 ky } uniformly. Without loss of gen-
erality, we assume k x >k y .

Then, by lemma [2"31

U{x)-4>{y)\\ p P h £^~l|ff(*,2 n )ll£

n=l
K

n=l

h f(2 k *) p hf(d(x,y)y.
The final step here is due to the concavity of /. □

3 Amalgamated products and HNN extensions
over finite groups

In this section, we prove that the compression exponent of amalgamated prod-
ucts and HNN extensions over finite groups are exactly as expected, depending
only on the compression of the initial groups.

9

Definition 3.1. Let V be the Q-skeleton of a geodesic metric space with
bounded vertex degree. We define a tree-grading of T {in the sense of Drufu
and Sapir, [DS0 5 to be a collection of non-empty subsets

V := {r\W

with the following properties:

(i) every vertex and every simple loop ofT is contained in some Fj,

(ii) ifi^j then T, t % Tj and \V(Ti) n V( r j)l < 1.

In particular, we may consider the Cayley graph of a free product of two finitely
generated groups.

At this point we fix some notation for paths in T. By \p\ we mean the edge
length of the path p. The initial vertex of p is denoted l(j>) and the terminal
vertex r(p). Define to be the unique closest point to e in IV Notice that if
r 4 D Tj ^ then {e 4 } = {e 3 } = T l D Tj.

For each such tree-graded metric space T with a basepoint e we define a quasi-
isometric tree-graded simplicial graph T' with set of pieces {1^} and basepoint
e' satisfying the additional property that any two pieces are disjoint.

To achieve this we build T' from the disjoint union of spaces [Jil^i} U W} a s
follows. Given two pieces r^r^ , we attach an edge between the unique pair of
points at distance precisely d(Ti , Tj ) if and only if this value is at most 1 in T
and d(e, ej) ^ d(e, ey). Finally we add a new vertex e' and add the edge ee' and
define the {e'} to be a element of {1^}.

This transformation can be achieved by a quasi-isometry.

The following picture illuminates both definition 13.11 and this process. Each
loop and each black dot not lying on a loop represents a piece.

In return for doing this additional messing around we obtain the following struc-
ture property.

For each vertex y we fix a geodesic p(y) between e and y and can write it as

p(y) = ITn, 7777//.., ■ ■ • yj^yjn

with the following restrictions:

(i) L(Vh) = e ' and r (%>.) = y>

10

(ii) for each I, the path yj l lies entirely in one piece Tj t , (note \yj t | may be 0),

(iii) y]~ is a path of length 1 with endpoints in two distinct pieces.

Notice that due to adding a new basepoint e' in the construction of V we ensure
the third condition is satisfied by yj[ .

In what follows we will assume any tree-graded graph takes this form. Before
stating the result, we will require the following two lemmas.

Lemma 3.2. Let M be a finite subset ofN such that M — {mi, m 2 , ■ ■ ■ , rn 2 k\
with rrii < m^+i and mi > 1.

Let p > 1 and let f : N — >• IR>o be a concave function such that is non-

decreasing. Then

m 2

m 2 , - m 2 i-i

Proof: For ease of notation we set m = m 2 i — m^i-i-

i=i

As is non-decreasing,

> [m 2i - m 2i -i) > } .

*-f m 2l n

2—1 n=l

The result then follows from the method in [Tesllj . theorem 7.3. We advertised
a self-contained document so relate the details here.

We rewrite as follows

^ n

n=l

E— > E 1 /(K2]) p

n=l n=m/2

> \f([m/2]r>^ +P f(m).

Lemma 3.3. Let M be a finite subset ofN such that M — {mi, m 2 , . . . , m 2 k\
with mi < m.j + i and mi > 1.

Let p > 1 and let f : N — ¥ R>o be a concave function with property (C p ). Then
there exists some uniform constant C such that

/ f(m 2 j) \ p m 2l - mgt-i < ^
^ V m 2l J m 2l

Moreover, if for each i, m 2 i < 2m 2 i—i, then

\p / /(m 2j _i) \ p m 2i - m 2t _i < 2P+ i c
t-r 1 \ m 2i -x J m 2i -i

11

Proof: As / is concave is non-increasing. Hence

v-/7(™2i)V, , ^ ff(n)

M^r) (--—)< £ '

n=m2i— i+l

Therefore,

f{m 2 i)\ P m 2 i - m 2i -i ^ 1 //(n)

^ ' m 2 , / m 2i ~ n

' n—l

which is uniformly bounded as / has property (C p )

n>2i-m2i-i ^ r,m 2 i~ m 2 i-i

;unu jjaru jusu nonce unau

decreasing

For the second part just notice that m2 L ™ 2 ' 1 < 2 m2 ' ™ 2 ' 1 and as f is non-

V m 2i )

□

With these two lemmata complete, we state our next theorem.

Theorem 3.4. Let V be the 0-skeleton of a simplicial graph. Suppose T admits
a tree-grading V = {Ti}i e j and fix a basepoint e and some p S [1, oo). Suppose
also that we are provided a concave function p' : — > Kq" and a collection
of 1-Lipschitz coarse embeddings of pieces ipi : Ti — > £ p {Xi) such that for all

//(dr(a;,i/)) < Ui{x) - ^{y)\\ p .
If p = 1 t/iere is o Lipschitz coarse embedding <j> of X into an £ p space with

Php'-

For p > 1, given any function f : N — > M> satisfying (C p ) there is a Lipschitz
coarse embedding <f> of X into an l v space with

p{n) h mm {p'(n),f(n}} .

Proof: As stated above, we may quasi-isometrically transform T so that it
is in the desired form. Without loss of generality we may translate each tpi so
that Vi( e i) = 0.

With ease of notation in mind we denote by y Ji the path yjly 3l ■ ■ ■ yj^yj n , we
write

HVji ) = iPii ( t (Vji )) - ^ji ( L (yji )) = i'ji ( T (Vn))
(the last equality following from the fact that L(yj t ) — e^) and finally set

\y h \ + 1 ifp=l,

wti(y) := <

f(\n-"\) I -

\y

We are now ready to define the embedding <f> : X -> £ P (X) © ieJ ^ P (AT,).
For leTwe set

wt ; (y)(5(x) if x = t(y jl ),
otherwise.

12

While for i 6 I we set

(mm

otherwise.

Here S(x) is just the unit vector in the £ P (T) corresponding to the characteristic
function of {x}. This embedding emphasises the tree-like structure of T, but in
such a way that pieces are not additionally distorted.

In order to complete the proof, we now prove that this embedding satisfies the
following inequalities:

p(d T (y,z))±\\<t>(y)-<l>(z)\\ p ±dr(y,z).

Given two vertices y and z, with corresponding chosen shortest paths

P(y) = WiVnVn ■ ■ ■ yJ^Din and P( z ) = ~ 1 z kl ~ 1 . . . z km _ 1 z km

we set i = i(y, z) to be such that r(j/j ! ) = T(z kl ) for all I < i but the next part
of the path differs, i.e. r(yj i+1 ) ^ T(z ki+1 ). Notice that by construction yj^, z kl
are equal as paths.

In doing so, we may split <f)(y) — <p(z) in to three parts, which we will call the
root R, the split S and the tail T.

(j){y) - 4>{z) = R + S + T, where
R = J2(^i(y)~wt l (z))S(e ll ),

l<i

S = (wt i+1 (y)S(e ji+1 ) - wt i+1 (z)8{e ki+1 )) + {ip(y ji+1 ) - ip(z ki+1 ))

and

T= [wt ; (2/)^)+Vfe)]- E [wti(^(e fcl )+^ fcl )].

l>i+2 l>i+2

Note ||^)-^)||^ = ||^ + ||^ + ||T||^.
Step 1: \\<t>(y)-<t>(z)\\ p ±dr(y,z).

Suppose d{x,y) < 1. Then T = 0, moreover, as / is concave

< i + (f(\y ji+1 \ i) - f(\y ji+1 \)) < i + /(i).

Therefore we need only provide a uniform bound on ||i?|| .
If p = 1, then R = 0, so we are left with the case p > 1.

Since j s a non-decreasing function for n sufficiently large,

« , E (/(|y \^!|+ 1 /(| ^ |))P (i^i + 1 )

13

By concavity, |/(n) — /(n+ 1)| < < 2 ^fc 1 ^ , justifying the penultimate
inequality. The final inequality comes from lemma [3731

Step 2: Here we show \\<f>{y) — 0(z)|| p — p(dr(v,z)) for V i= z.
Firstly, notice that

IIS'Hp > \\^{Vm+i) -V'^fci+JHp ^ p'( d r(T(j/j 1+1 ),r(z fei+1 ))).

So if <ir(i"(2/j i+1 ), r(zfc i+1 )) > ^c?r(y, z) we are done by the concavity of p' . If
not, then

wnih E wt K2/) p + E wt <w p

Z>i+2 i>z+2

which completes the proof in the case p — 1, as the sets {ji+2 , ji+3 , • • • } and
{fci+2, fci+3, ■ • ■ } are disjoint.

When p > 1, we use lemma |3~21 to deduce,

\\T\f p > E + E wt fe (^b/(k' 4+2 |F + /(k fc4+2 |f.

2>i+2 fc>i+2

Without loss of generality we may assume jy^ -1 " 2 ! > |z fci+2 |, so as / is concave,
U(x) - m\\ P P > h f(d{x,y)/A) h f(d(x,y)).

□

Corollary 3.5. Let G and H be finitely generated groups and let F be a finite
subgroup of G and H . Then for allp > I, all concave functions pa and pn which
come from Lipschitz coarse embeddings of G and H into l v spaces respectively
and all functions f with property (G^) there exist Lipschitz coarse embeddings
of

(i) G*fH with p(n) y min {pc(n), pnin), f(n)} and

(ii) HNN(G, F) with p(n) h min {p G (n), f(ri)}.
When p = 1 the above equations hold with f(n) = n.

Proof: It is an obvious corollary of the previous theorem that (i) holds
whenever F is the trivial group.

If G and H are both finite, then G*f H and HNN(G, F) are both hyperbolic by
[Gro87j . so the result holds. If G is infinite then HNN(G, F) is quasi-isometric to
G*Z, similarly, if at least one of G or H is infinite then G*fH is quasi-isometric
to G * H. Both these results are due to PW02]. Z isometrically embeds into
any £ p space, so we may choose pi{n) = n, completing the result. □

4 Relatively hyperbolic groups

Let X be the 0-skeleton of a simplicial graph with uniformly bounded valency,
let {Ai | i e /} be a countable collection of subsets of X. Fixing a basepoint
e G X, we introduce the following notation.

14

(i) C(Ai) is the set of all points lying on any geodesic with both endpoints in
A. '

(ii) G x .k is the set of all geodesies in \y, e], with d(x, y) < k. To ease notation
we write Gi :— G mx.aa .

(iii) Given a geodesic g, the i-domain of g, g\i is the restriction of g to the
convex hull (within g) of g n M- Note g\i C C(-Aj). We set g| + to be the
initial point of g\i and to be the terminal point, with respect to the
orientation inherited from g.

(iv) The length of a geodesic g in Ai is Z, (g) = d <?li~) + 1-

(v) Given a collection of geodesies G, and a constant if > we define df(G)
to be the set of g\i + which satisfy the following condition.

For all g[_ G G and any j G I with ^ (5^) > 5K,

d(e,g\ + ) i [d(e,l\j) + 2K,d(e,l\f)-2K]

(vi) We define the set I X (K) to be the set of i G J such that 9j (^x) 7^ $ an d
d(e 7 d l (gQ) > 3iT

The subset I' X {K) consists of those i G I X (K) such that x £ Ai.

We make a particular point of emphasising here that by (v), for all K and all
pairs of points x, y with d(x, y) < R,

d?( G x , k ) C df( G y ^ +R ).

Definition 4.1. The collection X

The 0-skeleton X of some simplicial graph of uniformly bounded valency equipped
with a basepoint e lies in the collection X if there is a covering of X by sub-
sets A = {Ai I i E 1} C V(X) and a constant K = K{X) with the following
properties:

(i) 1 < \{i I x G Ai}\ < K for all x G X and \C{Ai) D C{Aj)\ < K for all
i,j el with i ^ j.

(ii) For each t G N ; \{i G 4 (A") | d(a;, Aj) = i}| < if.

(iii) Given any two geodesies g x G \x, e] and g y G [y, e] which both intersect a
subset Ai G .4, i/ien d(<7 x |^, g^l^ - ) < K.

(iv) Le£ x,y E X be such that there exists some i G I' x (K)nI' y (K) with d(x, y) <
max ^ i|. Given any two geodesies g x G [a;, ej and g y G [y,e],
^(5£l^flj/L + ) < K. Moreover, if g^f) A { = 0, then h (5^) < if.

15

Note that (i) implies / is countable. From parts (iii) and (iv) it follows imme-
diately that both the values \d(x,Ai) — d(y, Ai)\ and h(g x ) — h{9y) are also

uniformly bounded under the hypotheses of (iv) so we increase K, if necessary,
to be at least as large as these bounds.

Of course, we can think of X as the closure of the set of spaces satisfying
conditions (i)-(iv) up to quasi-isometry, as we are looking only to calculate
quasi-isometry invariants.

Question 4.2. Is X already closed up to quasi-isometry within the collection
of O-skeletons of simplicial graphs!

Theorem 4.3. (cf. [TJ)

Let X £ X and let a collection of subsets {Ai}j £ j and constant K — K(X)
be suitable for the definition with respect to some basepoint e. Suppose we are
provided 1-Lipschitz embeddings ipi : A4 — > £ p (Xi) and a concave function p' :
Rq — > Rq such that for all x,y G Ai,

p'(d x (x,y)) < \\ipi(x) -ipi{y)\\ p

then, for all p > 1 and all functions f : N — > K>o with property (C£) there
exists a Lipschitz coarse embedding <fi from X to some £ p space such that for all
x,y e X,

p(n) y mm{p'(n), f(n)}.

Proof: We define to be some closest point of Ai to e, by condition (iii)
of definition 14. 1[ the diameter of the set of possible choices for is at most K.
Without loss of generality we may assume ipi( e i) = for each i <E I. As the
constant K is now fixed we will write I x and I' x in place of I X {K) and I' X (K)
respectively. Similarly, we drop the K in the notation .

For each i e I' x , we define functions Fi(x, k) £ £ P (X) as follows:
Fi(x,k)(y) -

mm{d(x,Ai),d x (y,ei) + 1}* if y G d t {G Xtk )
otherwise.

As a useful shorthand we set d x ^{y) = mm{d(x, Ai), dx(y, &i) + 1}- We then
define

The following three lemmas (mirroring lemmas 12.31 12.41 and I2.5[) provide useful
information on these new objects.

Lemma 4.4. For all x 6 X, g G [x, e], i G I' x with g\f G di(G Xt o), and

k < d(x,Aj)

d x Mt) < \\W*>k)\\ P P < d i( G *,k) (d x Ag\t) + K) ± d x M + ).

16

Proof: By assumption, g\f G di(G x< o) C di(G X) k) for all k < ffedii^ so t ne
first bound is satisfied.

For the second bound,

\\Fi(x,kX< <i(y)<\di(G x< k)\(d x< i(9\t) + K )-

The final bound holds as X has uniformly bounded valency and the diameter
of di{G X: k) is at most K, I4.1l iv).

Lemma 4.5. For all x G X , i € 7^, and <? G [x, e], wzi/i G di(G X: o),
\d x Mj) < WHtWWZ < (d x 4g\ + ) + K) .

Proof: The upper bound follows from lemma POI as g\f G di(G x _k) for all

d(x, Aij s

k < so

m -( x )\\ < 1 + 1

For the lower bound, we evaluate the contribution to ||-ffi(aO||p coming from

nudity.

Lemma 4.6. There exists some constant C > such that for all x. y G X

with d(x, y) < 1, all g G [at, e] and aZ/ i E I' x U I' y

\mx)-H t {y)\\l<C

dxAg\i

p- d(x,A 2 )P

Proof: We first bound the absolute value of Hi(x) — Hi{y) at some point
aed^GVjUd^Gl).

As d(x,y) < 1, di(G x>k ) Q d^Gy^+i), so

1(^(1)- Aid/)) (o) |

n Xt i(a) , . i n x i {a) i

where

s,i(a) := |n

<

a G dj( G z ,fc ) |

Also, dj(G Xk ) C di( G Vik +i ), so |n X)i (a) - < l+ldfoAj) - ^ 2 -

If i G ZJ, \ ly then Hi(y) — 0, and n Xi i(a) < 2. Again, we use the fact that
di(G x ) is uniformly bounded bv !4.1f iv) and the uniformly bounded valency of
X so we are done. The case i € I' y \ I' x is treated in the same way.

Suppose now that i G I^f] I' y , so Aj), d(y, Ai) > 1. Notice that d x ,i(a)

1 7

d Vt i(a) unless one (or both) are equal to d(x,Ai) (respectively d(y,Ai)). There-
fore

i. k p— i p— i

d x ,i{a)p - dy :i (a)p < mm{d(x,Ai),d(y, Ai)} v <2d[x > A i ) p .

Finally, combining these observations we have

d(y, Ai)n X! i(a)d Xi i(a)p - d(x, Ai)n y ^(a)d yti (a)p

\{Hi(x) - Hi(y)) (a) |

d(x,Ai)d(y,Ai)
By the triangle inequality, we can bound this from above by

n Xt jd X! i(a)p \d(x,Aj) -d{y,Aj)\ d{x,Aj)d x ^{a)p \n xA (a) - n Vti (a)\
d(x,Ai)d(y,Ai) d(x, A t )d{y, Ai)

d{x,Ai)n y j(a) d y>i (a)p - d x ,i(ap
+

d(x,Ai)d(y,Ai)

Applying all the previous deductions and noticing that n Xl j(a) < d{x,A{) we
obtain a uniform constant C such that

\(Hi(x) - Hi(y)) (a)\ <C

, d x< i(a)v
d(x,Ai)

Finally, we use definition I4.1f iv) to deduce that di(G l x ) U di(G y )
bounded and the lemma follows.

We are now ready to define the first part of our embedding:

m := E m^ Hi{ x).

ie/ , d(x,Ai)p

Lemma 4.7. (jf : X — > £ P (X) is Lipschitz for all p > 1.
Proof: Consider two points x, y G X with d(x, y) < 1.
Firstly, suppose i G I x \ I' y . Then by lemma 14.61

■*<*>«; **3^

for any geodesic 5 G [x, e], but bv l4.1f iv). £(<?|j) < 2if, so

The case i E I y \I' x is treated similarly and as Ai) — d(y, Ai)\ < 1,

» h *>ii* s raw

is uniformly
□

18

By the triangle inequality, the contribution made to ||^ s (x) — by those

i G I' x n I' y is at most

+ E (/(^_/«E^y l|JJ(Wlc . (t2)

ie/^n/; V ^i)* A i) p J

As ^5^- is non-decreasing we may use the same argument as in the tree-graded

nP

case to deduce that (+2) is bounded from above (up to some uniform multiplica-
tive constant) by

min{diam(^(g,)) + l,d(x ; A I )} / f(d(x, A,)) \ p
ie |^, d(x,Ai) + l { d(x,Ai) J

Also, by lemma S3] and the fact that / is concave, (fi) is bounded from above
(up to some uniform multiplicative constant) by

mm{d iam(ft(G%)) + l,d(x,Aj)} ( f(d{x,Ai))\ p

d(x,Ai) + l V d(x,Ai)

Hence,

\WM MAU*^ \- ^diam^G^ + MQr,^)} f f(d(x,A l )) \ p

d(a;,j4 4 ) + 1 V d(a;,A)

which is uniformly bounded. (The additional +1 in the denominator here is just
to cover the situation where d(x, Ai) = for some i G I' y .) To see this recall that
there is a uniform bound on the number of pieces any point can lie in 14. If i) and
a uniform bound on the cardinality of the subset of I x with d(x, A{) = n. l4.1f ii).
Combining these we can partition I' x U I' y into 2K subsets so that the above sum
restricted to any such subset satisfies the hypotheses of lemma 13.31 □

For the second part of the embedding we make a complementary construction,
using the existing embeddings of pieces (ipi)i e i.

We set a x ,i = ^ \dj{ G x . k )

and define

k=0

i ■ f d(x,j
k x .i = mm <^ a x<i , 1 H —

Recall that we made the convention ipi(ei) = for each j 6 7.

We then proceed towards the definition of the second part of the embedding.

F!(x,k)= Yl ^W-

19

Following the usual averaging procedure, we now set

The second part of the embedding cj> 1 : X — » (D i6J £ p (Xi) is defined as
Lemma 4.8. For all p > 1, <j> is Lipschitz.

Proof: Let x,y G X \ B^ie) with d(x,y) < 1. We show that for each

i G 4U I y ,

for some C > not depending on i. This suffices bv l4.1f i).(n).

Initially, suppose k Xt i — a Xt i and k y j = a y ^. Then notice that the function

- E x(9i(G^))

fc<;

is non-negative and has I 1 norm exactly 1, as k x ,i = ^(G^.fe)

Moreover,

-r^ E x(3»(G*,fc)) - t~~ E X(^(G^)) (fa)

has ^ 1 norm at most d ^ xA .^ ^ or some uniform constant C, and the sum of its
entries is 0.

The second of these claims follows from the fact that this is a difference of non-
negative functions of I 1 norm 1. For the first we use the same trick as in the
hyperbolic proof to prove \k x ,i — k Vi \ is uniformly bounded.

As di{G x ^k) C di(G y fc+i), the contribution to \k x> i — k Vi \ made by any point

a G di(G x )Udi(G y ) is at most 1. Moreover, the set di(G x )Udi(G y ) has uniformly

bounded cardinality, so \k x ,i — ky,i\ is uniformly bounded by a constant we will
label C" during the remainder of the proof of this lemma.

Next, fix any point a G di(G x ) U di(G y ). Again, as di{G Xt k) Q di(G Vt k+i), the
contribution to (fa) coming from a is at most

n Xi j(a) _ n yii (a)

k ■ ~ h

™ x ,i y,i

20

In particular, \n x ^(a) — n y ^{a)\ < 1 and n y ,i(a) < k y> i, so

n Xi i(a) %,i(a)

<

<

<

n x .i{a) n y .i(a)

+

n Vti {a) n Vl i(a)

\n Xl j(a) - n yA (a)\ n y ^(a) \k Vi

kx,i

C"

h h

<

c

d(x,A l ) + 1'

with the final step coming from the fact that k Xi i > 1+ ffedii Now we return
our attention to H'^x) — H'^y), which we deduce from our previous arguments
can be written in the following way:

H'i{x) - H-(y) = ^2,^i{b n ),

where each b n € di(G x ) U di(G y ) and /x„ is the value of the function (+3) at b n .
From the above argument we know that ^2 l^n = and X) I I — d( x A )+i •

But for any two points a, b G di(G x ) U di(G y ), \\ipi(a) — < 2i4T, by defini-

tion HTTJiii) and (iv) and the fact that each ipi is 1-Lipschitz. Therefore,

2KC

\H'M-Hl

<

" d{x,Ai) + l

Instead, assume without loss of generality that k x ,i > a x .i, then di{G X fi) =
and by definition 14. If iv) the length of any z-domain of any g e di(G y ) U di(G x )
is bounded from above by 2K .

Hence, using the fact that \k x ,i — k y ^\ < K (see the comment following definition
I4.1j) . we deduce in the same way as above that (+3) has I 1 norm bounded by
d(x ^ 0r some uniform constant C .

Again writing

we see that as each ^ is 1-Lipschitz,

\\H[{x)~H[{y)\\ p <Y,\^\\\Ub n )\\ p <

2KC

□

d(x, Ai) + 1'

completing the lemma.
Now we are ready to prove the theorem using the embedding

^:I^P(X 2 )©0f(Ii) given by 4>{x) = 4> s {x) + <t>\x).
iei

This is Lipschitz by lemmas 14.71 and 14.81

Consider x,y £ X with d(x,y) > CK (C is chosen such that p'(CK) > 35K
and C > 35).

Fix geodesies g x £ [x, e] and g y £ {y, e].

21

Set x y to be the closest point p XiV on g x to e such that d(p XiV ,g y ) > 5-fsT and
define similarly. Notice that if x y ,y x € C(Ai) for some i, then that i is
unique, as any g £ lx y ,y x } is entirely contained in C(A{) and intersections of
these sets have cardinality at most K, by definition l4.1f iV

Let J x = {j e I x | 9x\i x , Xy \ n A 3 ■, } and J' x = J x n 4- We define J y and
similarly.

J x PI Jy has cardinality at most 1, by definition 14 . 1 f in) .

Suppose | J K n J y \ = 1, label that index i and suppose d(g x \f , g y \f) > y),
then

W«)-0(y)ll?>ll^)-^(»)llp.

We notice that the sets di(G x ) and di{G x ) are disjoint as di(G x ) has diameter
at most if, so the function defined in the proof of lemma |4~51 (^3) has t 1 norm
2 in this case. Therefore, we can write

Hi(x) - Hi(y) =J2»nHl(x)(b n ) -J2^mHi(y)(b m )

n m

with n(a) being the value of (fa) evaluated at a, ^ m jU m = — J^n Mn = 1 and
the sets {b m } and {6 n } disjoint. Pairing up the fx n and fi m and applying 14.1 f ry)
we see that

\\H((x) - H'(y)\\ p > p'(d{ g x \i , g y \t )) — AK > ±p>(d(x,y)),

where the last step comes from the concavity of p' and the upper bound on
d(x,y).

If J x n J y — {i} we now set x y = g x \f and y x = g y \f ■ Otherwise we leave x y
and y x as before.

Suppose now, without loss of generality, that d(x,x y ) > d(y,y x ), so d(x,x y ) >
jd(x,y). If there exists some j S J x \ J y with lj(g x ) > ^d(x,y), then

\\<Kx)-<Kv)\\ P > \\H' j ( x )- H M\ P =\\ H 'M\ P
If this does not happen, then

As every point p lying on y x at distance between ^d(x, x y ) and d(x,x y ) from x
lies in some C(Aj) with j e J x \ J y , we use lemma ET21 and the lower bound of
lemma S3] to deduce

\\m-m\\ihf(~d(x,x y )-i\ p .

Finally, / is concave, so this gives

H(x) - 0(y)|| p h f (yd(x,y) -ljh f(d(x,y).

□

22

Corollary 4.9. Let G be a finitely generated group which is hyperbolic rel-
ative to a collection of subgroups {Hi}. Given any p > 1, any collection of
Lipschitz coarse embeddings of Hi into £ p spaces with associated concave bound-
ing functions pi and any function f with property (Cp) there is a Lipschitz
coarse embedding of G with

p(n) y min{ j o 4 (n), f(n)} .

Proof: (References in this corollary are to DS05 unless otherwise stated).

It suffices to show G G X. Appendix A proves that G is asymptotically tree-
graded with respect to a collection of cosets, which we will label {I\ | i G /}.
To satisfy property (i) we put each point not lying in such a coset into its own
piece. Then A is the collection of all M-neighbourhoods of these pieces, where
M is the constant obtained in the proof of the Rips' hyperbolicity of saturations,
4.27.

Property (i) then follows from theorem 4.1(ai), as C(A) is contained in the Mr
neighbourhood of A for each A £ A, with r > 2 the constant obtained in lemma
8.10.

Property (iii) is the conclusion of corollary 8.14. Finally (iv) follows from the
Rips' hyperbolicity of saturations and the argument in lemma |2~T1 Every A G A
is quasi-isometric to either a point, or one of the Hi, so we do not need to worry
about the ipi being 1-Lipschitz, as there are only finitely many of them. □

This corollary proves that X contains all asymptotically tree-graded uniformly
discrete metric spaces with bounded geometry.

Question 4.10. Is every X G X asymptotically tree-graded"!

References

[AD] Goulnara Arzhantseva and Thomas Delzant. Examples of random groups. Avail-

able from www.unige.ch/math/folks/arjantsc/Abs/random.pdf.

[ADS09] Goulnara Arzhantseva, Cornelia Drut.u, and Mark Sapir. Compression functions
of uniform embeddings of groups into Hilbert and Banach spaces. J. Reine Angew.
Math., 633:213-235, 2009.

[Ali05] Emina Alibcgovic. A combination theorem for relatively hyperbolic groups. Bull.
London Math. Soc, 37(3):459-466, 2005.

[AMM85] Israel Aharoni, Bernard Maurey, and Boris S. Mityagin. Uniform embeddings of
metric spaces and of Banach spaces into Hilbert spaces. Israel J. Math., 52(3):251-
265, 1985.

[Aus] Tim Austin. A finitely-generated amenable group with very poor compression into

Lebesgue spaces, (preprint). arXiv:0909.2047v2.

[BCH94] Paul Baum, Alain Connes, and Nigel Higson. Classifying space for proper actions
and X-theory of group C*-algebras. 167:240-291, 1994.

[BDG+05] Mihai Badoiu, Kedar Dhamdhere, Anupam Gupta, Yuri Rabinovich, Harald
Racke, R. Ravi, and Anastasios Sidiropoulos. Approximation algorithms for low-
distortion embeddings into low-dimensional spaces. In Proceedings of the Six-
teenth Annual ACM-SIAM Symposium on Discrete Algorithms, pages 119-128
(electronic), New York, 2005. ACM.

[BDS07] Sergei Buyalo, Alexander Dranishnikov, and Viktor Schroeder. Embedding of
hyperbolic groups into products of binary trees. Invent. Math., 169(1):153-192,
2007.

23

[BH] Nikolay Brodskiy and Jose Higes. Assouad-nagata dimension of tree-graded spaces.

arXiv:0910.2378vl.

[BM08] Jeffrey Brock and Howard Masur. Coarse and synthetic Weil-Petersson geometry:
quasi-fiats, geodesies and relative hyperbolicity. Geom. Topol, 12(4):2453-2495,
2008.

[Bou85] Jean Bourgain. On Lipschitz embedding of finite metric spaces in Hilbert space.
Israel J. Math., 52(l-2):46-52, 1985.

[Bow99] Brian Bowditch. Relatively hyperbolic groups. Digital copy available from
http : //eprints . soton. ac. uk/29769 / 1 /bhb-relhyp. pdf , 1999 .

[BS00] Mario Bonk and Odcd Schramm. Embcddings of Gromov hyperbolic spaces.
Geom. Fund. Anal, 10(2):266-306, 2000.

[BS05] Sergei Buyalo and Viktor Schroeder. Embedding of hyperbolic spaces in the prod-
uct of trees. Geom. Dedicata, 113:75-93, 2005.

[BS08] Nikolay Brodskiy and Dmitry Sonkin. Compression of uniform embcddings into
Hilbert space. Topology Appl, 155(7):725-732, 2008.

[CDGY03] Xiaoman Chen, Marius Dadarlat, Erik Gucntner, and Guoliang Yu. Uniform
embcddability and exactness of free products. J. Funct. Anal, 205(1):168-179,
2003.

[CN05] Sarah Campbell and Graham A. Niblo. Hilbert space compression and exactness
of discrete groups. J. Funct. Anal, 222(2):292-305, 2005.

[CZ06a] Huai-Dong Cao and Xi-Ping Zhu. A complete proof of the Poincare and ge-
ometrization conjectures — application of the Hamilton-Pcrelman theory of the
Ricci flow. Asian J. Math., 10(2):165-492, 2006.

[CZ06b] Huai-Dong Cao and Xi-Ping Zhu. Erratum to: "A complete proof of the Poincare
and geomctrization conjectures — application of the Hamilton-Perelman theory of
the Ricci flow" [Asian J. Math. 10 (2006), no. 2, 165-492]. Asian J. Math.,
10(4):663, 2006.

[Dah03a] Francois Dahmani. Combination of convergence groups. Geom. Topol, 7:933-963
(electronic), 2003.

[Dah03b] Francois Dahmani. Les groupes relativement hyperboliques et leurs bords.

Prepublication de l'lnstitut de Recherche Mathematique Avancee [Prepublication
of the Institute of Advanced Mathematical Research], 2003/13. Universite Louis
Pasteur Departement de Mathematique Institut de Recherche Mathematique
Avancee, Strasbourg, 2003. These, l'Universite Louis Pasteur (Strasbourg I),
Strasbourg, 2003.

[dCT08] Yves de Cornulier and Romain Tessera. Quasi-isomctrically embedded free sub-
semigroups. Geom. Topol, 12(l):461-473, 2008.

[dCTV07] Yves de Cornulier, Romain Tessera, and Alain Valette. Isometric group actions
on Hilbert spaces: growth of cocycles. Geom. Funct. Anal, 17(3):770-792, 2007.

[DG03] Marius Dadarlat and Erik Guentncr. Constructions preserving Hilbert space uni-
form embcddability of discrete groups. Trans. Amer. Math. Soc, 355(8):3253-
3275 (electronic), 2003.

[DG07] Marius Dadarlat and Erik Gucntner. Uniform embcddability of relatively hyper-
bolic groups. J. Reine Angew. Math., 612:1-15, 2007.

[DJ99] Alexander Dranishnikov and Tadeusz Januszkiewicz. Every Coxeter group acts
amenably on a compact space. Topology Proc, 24(Spring):135— 141, 1999.

[DJ00] Michael W. Davis and Tadeusz Januszkiewicz. Right-angled Artin groups are com-
mensurable with right-angled Coxeter groups. J. Pure Appl Algebra, 153(3):229-
235, 2000.

[DrclO] Dennis Drccscn. Hilbert space compression for free products and HNN-cxtcnsions.
arXiv:1002.3879v3, 2010.

[DS05] Cornelia Drul^u and Mark Sapir. Tree-graded spaces and asymptotic cones of
groups. Topology, 44(5):959-1058, 2005. With an appendix by Denis Osin and
Mark Sapir.

24

[Far98] Benson Farb. Relatively hyperbolic groups. Geom. Fund. Anal, 8(5):810-840,
1998.

[Gal08] Swiatoslaw R. Gal. Asymptotic dimension and uniform embeddings. Groups
Geom. Dyn., 2(l):63-84, 2008.

[GK04] Erik Guentner and Jerome Kaminkcr. Exactness and uniform embcddability of
discrete groups. J. London Math. Soc. (2), 70(3):703-718, 2004.

[Gro87] Misha Gromov. Hyperbolic groups. In Essays in group theory, volume 8 of Math.
Sci. Res. Inst. PubL, pages 75-263. Springer, New York, 1987.

[Gro93] Misha Gromov. Asymptotic invariants of infinite groups. In Geometric group
theory, Vol. 2 (Sussex, 1991), volume 182 of London Math. Soc. Lecture Note
Ser., pages 1-295. Cambridge Univ. Press, Cambridge, 1993.

[GroOO] Misha Gromov. Spaces and questions. Geom. Fund. Anal, (Special Volume, Part
I):118-161, 2000. GAFA 2000 (Tel Aviv, 1999).

[HLW06] Shlomo Hoory, Nathan Linial, and Avi Wigderson. Expander graphs and their
applications. 2006.

[KL08] Bruce Kleiner and John Lott. Notes on Perelman's papers. Geom. Topol,
12(5):2587-2855, 2008.

[KY06] Gennadi Kasparov and Guoliang Yu. The coarse geometric Novikov conjecture
and uniform convexity. Adv. Math., 206(l):l-56, 2006.

[LLR95] Nathan Linial, Eran London, and Yuri Rabinovich. The geometry of graphs and
some of its algorithmic applications. Combinatorica, 15(2):215-245, 1995.

[LP] James Lee and Yuval Peres. Harmonic maps on amenable groups and a diffusive

lower bound for random walks, (preprint). arXiv:0911.0274v4.

[MT07] John Morgan and Gang Tian. Ricci flow and the Poincare conjecture, volume 3
of Clay Mathematics Monographs. American Mathematical Society, Providence,
RI, 2007.

[MT08] John Morgan and Gang Tian. Completion of the proof of the geometrization
conjecture. arXiv:0809.4040vl, 2008.

[NP08] Assaf Naor and Yuval Peres. Embeddings of discrete groups and the speed of
random walks. Int. Math. Res. Not. IMRN, pages Art. ID rnn 076, 34, 2008.

[Osi06] Denis V. Osin. Relatively hyperbolic groups: intrinsic geometry, algebraic proper-
ties, and algorithmic problems. Mem. Amer. Math. Soc, 179(843):vi+100, 2006.

[Per02] Grisha Perclman. The entropy formula for the ricci flow and its geometric appli-
cations. arXiv:math/0211159vl, 2002.

[Per03] Grisha Perelman. Ricci flow with surgery on three-manifolds.

arXiv:math/0303109vl, 2003.

[PW02] Panos Papasoglu and Kevin Whyte. Quasi-isomctries between groups with in-
finitely many ends. Comment. Math. Helv., 77(1): 133-144, 2002.

[Sapll] Mark Sapir. Asphcrical groups and manifolds with extreme properties.
arXiv:1103.3873v3, 2011.

[SmilO] Alexander Smirnov. The linearly controlled asymptotic dimension of the funda-
mental group of a graph manifold. Algebra i Analiz, 22(2):185-203, 2010.

[Tesll] Romain Tessera. Asymptotic isopcrimctry on groups and uniform embeddings
into Banach spaces. Comment. Math. Helv., 86(3):499-535, 2011.

[TuOl] Jean-Louis Tu. Remarks on Yu's "property A" for discrete metric spaces and
groups. Bull. Soc. Math. France, 129(1):115-139, 2001.

[Wis] Daniel Wise. The structure of groups with a quasiconvex hierarchy.

[Woj91] Przcmyslaw Wojtaszczyk. Banach spaces for analysts, volume 25 of Cambridge
Studies in Advanced Mathematics. Cambridge University Press, Cambridge, 1991.

[Yam04] Asli Yaman. A topological characterisation of relatively hyperbolic groups. J.
Reine Angew. Math., 566:41-89, 2004.

[YuOO] Guoliang Yu. The coarse Baum-Connes conjecture for spaces which admit a uni-
form embedding into Hilbert space. Invent. Math., 139(l):201-240, 2000.

25

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
