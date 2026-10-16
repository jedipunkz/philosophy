---
source: "https://arxiv.org/abs/0807.4229v3"
title: "Positive circuits and maximal number of fixed points in discrete dynamical systems"
author: "Adrien Richard"
year: "2008"
publication: "arXiv preprint / cs.DM"
download: "https://arxiv.org/pdf/0807.4229v3"
pdf: "https://arxiv.org/pdf/0807.4229v3"
captured_at: "2026-10-16T04:03:31Z"
updated_at: "2026-10-16T04:03:31Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Richard Rorty"
tags:
  - "現代哲学"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Positive circuits and maximal number of fixed points in discrete dynamical systems

- 著者: Adrien Richard
- 年: 2008
- 掲載情報: arXiv preprint / cs.DM
- 情報源: [arxiv](https://arxiv.org/abs/0807.4229v3)
- ダウンロード: https://arxiv.org/pdf/0807.4229v3
- PDF: https://arxiv.org/pdf/0807.4229v3

## Obsidian Links

- 研究動向: [[リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代哲学 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

We consider the Cartesian product X of n finite intervals of integers and a map F from X to itself. As main result, we establish an upper bound on the number of fixed points for F which only depends on X and on the topology of the positive circuits of the interaction graph associated with F. The proof uses and strongly generalizes a theorem of Richard and Comet which corresponds to a discrete version of the Thomas' conjecture: if the interaction graph associated with F has no positive circuit, then F has at most one fixed point. The obtained upper bound on the number of fixed points also strongly generalizes the one established by Aracena et al for a particular class of Boolean networks.

## PDF Text

Positive circuits and maximal number of fixed points in discrete dynamical systems

arXiv:0807.4229v3 [cs.DM] 1 Dec 2008

Adrien Richard
Laboratoire I3S, UMR 6070 CNRS & Université de Nice-Sophia Antipolis,
2000 route des Lucioles, 06903 Sophia Antipolis, France.
Email: richard@i3s.unice.fr

Abstract
We consider the Cartesian product X of n finite intervals of integers and a map F
from X to itself. As main result, we establish an upper bound on the number of fixed points for F which only depends on X and on the topology of the positive circuits of the interaction graph associated with F . The proof uses and strongly generalizes a theorem of Richard and Comet which corresponds to a discrete version of the
Thomas’ conjecture: if the interaction graph associated with F has no positive circuit, then F has at most one fixed point. The obtained upper bound on the number of fixed points also strongly generalizes the one established by Aracena et al for a particular class of Boolean networks.
Key words: Discrete dynamical system, Discrete Jacobian matrix, Interaction graph, Positive circuit, Fixed point.

1

Introduction

We are interested by the number of fixed points for maps that operate on the
Cartesian product of n finite intervals of integers (when this product is {0, 1}n , such maps are usually called Boolean network ). Our motivation comes from biology, where these maps are extensively used to describe the behavior of gene networks. The context is then the following.
When studying gene networks, biologists often illustrate their results by interaction graphs. These are directed graphs where vertices correspond to genes and where edges are labelled with a sign: a positive (resp. negative) edge from j to i means that the protein encoded by gene j activates (resp. represses) the synthesis of the protein encoded by gene i. These graphs are then used as basis
Preprint submitted to Elsevier

3 June 2018

to generate dynamical models describing the temporal evolution of the concentration of the encoded proteins (see [1] for a literature review). Unfortunately, these models require, in most cases, unavailable information on the strength of the interactions. One is thus faced with the following difficult question: Which dynamical properties of a gene network can be inferred from its interaction graph (in the absence of information on the strength of the interactions)?
In this paper, we focus on this question in a general discrete modeling framework. The set of states of a network of n genes is represented by the product
Q
X = ni=1 Xi of n finite intervals of integers. Each interval Xi then corresponds to the set of possible concentration levels for the protein encoded by gene i.
On one hand, the dynamics of the network is described by the successive iterations of a map F from X to itself whose fixed points correspond to the stable states of the network. At this stage, it is worth noting that the number of stable states is a key feature of gene networks dynamics: according to an idea of Delbrück [2], the presence of multiple stable states is one possible mechanism for biological differentiation. One the other hand, the interaction graph of the network is deduced from F in two steps. First, to each state x ∈ X
and to each directional vector v ∈ {−1, 1}n such that x + v ∈ X is associated a local interaction graph GF (x, v) which contains a positive (resp. negative)
edge from j to i if fi (x1 , . . . , xj + vj , . . . , xn ) − fi (x)
vj is positive (resp. negative) (fi denotes the ith component of F ). Then, the global interaction graph G(F ) of the network is defined to be the union of all the local interaction graphs. Note that each local interaction graph is a subgraph of the global one, and that the global interaction graph can have both a positive and a negative edge from one vertex to another.
In this setting, Richard and Comet [3] partialy answer the previous question by proving a well known conjecture of René Thomas relating the stable states of the network to the positive circuits of its local interaction graphs (a circuit is positive if it has an even number of negative edges). A weak form of their result follows (the original statement needs additional definitions and is given latter in the paper):
Theorem 1 [3] Let X be a product of n finite interval of integers, and let
F be a map from X to itself. If all the local interaction graphs GF (x, v) are without positive circuit, then F has at most one fixed point.
Aracena, Demongeot and Goles [4,5] proved another theorem relating stable states to positive circuits. They establish, in a boolean context, an upper bound on the number of stable states which only depends on the positive circuits of the global interaction graph of the network:
2

Theorem 2 [4,5] Let F be a map from {0, 1}n to itself such that G(F ) has no both a positive and a negative edge from one vertex to another. If I is a subset of {1, . . . , n} such that each positive circuit of G(F ) has at least one vertex in I, then the number of fixed points for F is less than or equal to 2|I| .
The main result of this paper is a significative generalization of both Theorem 1
and Theorem 2. A weak form of this result is:
Theorem 3 Let X = ni=1 Xi be a product of n finite interval of integers, and let F be a map from X to itself. If I is a subset of {1, . . . , n} such that each positive circuit of each local interaction graph GF (x, v) has at least one vertex
Q
in I, then the number of fixed points for F is less than or equal to i∈I |Xi |.
Q

Theorem 3 implies Theorem 1, since if all the local interaction graphs GF (x, v)
are without positive circuit then I = ∅ satisfies the conditions of Theorem 3.
So the corresponding bound is 1 and Theorem 1 is recovered. Theorem 3 also implies Theorem 2. Indeed, let F be a map from X to itself and suppose I to be such that each positive circuit of G(F ) has at least one vertex in I. Then
I satisfies the conditions of Theorem 3 because each local interaction graphs
Q
GF (x, v) is a subgraph of G(F ). So the corresponding bound is i∈I |Xi | and it equals 2|I| in the particular case where X is the n-cube {0, 1}n . We thus recover the conclusion of Theorem 2 (even if G(F ) has both a positive and a negative edge from one vertex to another). The proof of Theorem 3, which is done by induction on I with Theorem 1 as base case, is independent of the proof of
Theorem 2 given in [4,5]. Note also that Theorem 2 does not imply Theorem 1
even if this latter is stated for maps F from {0, 1}n to itself such that G(F )
has no both a positive and a negative edge from one vertex to another.
The paper is organized as follow. In Section 2, in order to obtain a bound stronger than the one mentioned above and more relevant from a biological point of view, we focus on the asynchronous iterations of F that Thomas use to describe the dynamics of gene networks [6,7,8,9]. First, we represent these iterations under the form of a directed graph Γ(F ) on X usually called asynchronous state transition graph. Then, we define the attractors of Γ(F ) to be the smallest subsets of states without output edges in Γ(F ). The fixed points of F then correspond to particular attractors. In Section 3, we characterize a subgraph GF (x, v) of GF (x, v) which only depends on Γ(F ) and which is, for this reason, well suited to the study of Γ(F ). In Section 4 we establish our main result: an upper bound on the number of attractors in Γ(F ) which only depends on the map GF and which has Theorem 3 as immediate consequence.
Final comments are given in Section 5. These are about the influence of connections between positive circuits and the interest of the established bound in the context of the so called Thomas’ logical method [6,7,8,9] which is, in practice, one of the most usual discrete modeling method of gene networks.
3

2

Asynchronous state transition graph and attractors

Let X = ni=1 Xi be the product of n finite intervals of integers of cardinality strictly greater than 1, and consider a map F from X to itself,
Q

x = (x1 , . . . , xn ) ∈ X 7→ F (x) = (f1 (x), . . . , fn (x)) ∈ X.
In the following definition, we attach to F a directed graph on X called asynchronous state transition graph. According to Thomas [6,7,8,9], this state graph can be seen as a model for the dynamics of a network of n genes: the set of vertices X is the set of possible states for the network (each interval Xi corresponds to the possible concentration of the protein encoded by gene i), and each path corresponds to a possible evolution of the system. [Asynchronous state transition graphs can also be seen as discretizations of piecewise-linear differential systems, see [10,11] for instance.]
Definition 1 The asynchronous state transition graph of F is the directed graph Γ(F ) whose set of vertices is X and which contains an edge from x to y if there exists i ∈ {1, . . . , n} such that fi (x) 6= xi

and

y = x + sign(fi (x) − xi ) · ei ,

where ei denotes the n-tuple whose ith component is 1 and whose other components are 0, and where sign(a) = a/|a| for all integer a 6= 0.
[Following this description of the dynamics, fi (x) can be seen as the value toward which the concentration xi of the protein encoded by gene i evolves: at state x, there exists a state transition allowing the ith component of the system to increase (resp. decrease) if and only if xi < fi (x) (resp. xi > fi (x)).]
The fixed points of F have no successor in Γ(F ) and naturally correspond to the stable states of the system. In the next definition, we introduce the notion of attractor which extends, in a natural way, the notion of stable state.
Definition 2 A trap domain of Γ(F ) is a non-empty subset A of X such that, for all edges (x, y) of Γ(F ), if x ∈ A then y ∈ A. An attractor of Γ(F ) is a smallest trap domain with respect to the inclusion relation.
In other words, the attractors of Γ(F ) are the smallest set of states that we cannot leave. They extend the notion of stable state in the sense that x is a fixed points of F if and only if {x} is an attractor of Γ(F ). Note also that there always exists at least one attractor (since X is a trap domain). Other easy observations follow: (1) From each state, there is a path which leads to an attractor (this is why one can say that attractors perform, in weak sense, an
4

attraction); (2) Attractors are strongly connected components; (3) Attractors are mutually disjointed (this point used in the proof of our main result).

3

Discrete Jacobian matrix and interaction graph

In this section, we introduce a notion of local interaction graph well suited to the study of Γ(F ). We proceed as in [3] by first introducing a discrete Jacobian matrix for F based on a notion of discrete directional derivative.
Let X ′ be the set of couples (x, v) such that x ∈ X, v ∈ {−1, 1}n and x+v ∈ X.
Definition 3 For all (x, v) ∈ X ′ , we call Jacobian matrix of F evaluated at x along the directional vector v the n × n matrix F ′ (x, v) = (fij (x, v)) defined by fij (x, v) =

fi (x + vj ej ) − fi (x)
vj

(i, j = 1, . . . , n).

[If vj is positive (resp. negative), then fij (x, v) may be seen as the right (resp.
left) partial derivative of fi with respect to the jth variable evaluated at x. In both cases, fij (x, v) is a natural discrete analogue of (∂fi /∂xj )(x).]
An interaction graph is here a directed graph whose set of vertices is {1, . . . , n}
and where each edge is provided with a sign. More formally, each edge is characterized by a triple (j, s, i) where j (resp. i) is the initial (resp. final)
vertex and where s ∈ {−1, 1} is the sign of the edge. The set of edges of an interaction graph G is denoted E(G). An interaction graph G is a subgraph of an interaction graph G ′ if E(G) ⊆ E(G ′ ).
Definition 4 We call interaction graph of F evaluated at (x, v) ∈ X ′ , and we denoted by GF (x, v), the interaction graph which contains a positive (resp.
negative) edge from j to i if fij (x, v) is positive (resp. negative).
[To illustrate this definition, assume that fij (x, v) is positive and that vj = 1.
Then, fi (x) < fi (x + ej ) so we can say that, at state x, an increase of xj induces an increase of fi , that is, an increase of the value toward which the ith component of the system evolves. In other words, j acts as an activator of i, and we have a positive edge from j to i in GF (x, v).]
In our context, the obvious fact that GF (x, v) does not only depend on Γ(F )
is not satisfactory since it is commonly accepted that the interaction graph of a network only depends on its dynamics, which is here characterized by Γ(F ).
This lead us, as in [3], to slightly modify the definition of GF (x, v) in order to obtain an interaction graph GF (x, v) which only depends Γ(F ).
5

Definition 5 We call interaction graph of F evaluated at (x, v) ∈ X ′ with thresholds, and we denote by GF (x, v), the interaction graph which contains a positive (resp. negative) edge from j to i if fij (x, v) is positive (resp. negative)
and if fi (x) and fi (x + vj ej ) are on both sides of (the threshold) xi + vi /2.
[a and b are on both sides of c if a < c < b or b < c < a.]
Remark 1 GF (x, v) is a subgraph of GF (x, v) (often strict since the additional condition “on both sides of the threshold” is rather strong).
Remark 2 The introduction of GF (x, v) has been motivated by arguments coming from the modeling context. Another relevant argument is the following: because GF (x, v) is a subgraph of GF (x, v), all the incoming results remains valid but becomes less strong when stated with GF (x, v) instead of GF (x, v).
Remark 3 In the Boolean case, i.e. when X = {0, 1}n , GF (x, v) = GF (x, v).
Definition 6 We call global interaction graph of F , and we denote by G(F ),
S
the interaction graph whose set of edges is (x,v)∈X ′ E(GF (x, v)).
Obviously, G(F ) only depends on Γ(F ) and can thus be seen as the global interaction graph of the network of dynamics Γ(F ). Note that G(F ) can have both a positive and a negative edge from one vertex to another.
Now, we recall the notion of positive circuit and the notion of positive feedback vertex set. This has been introduced by Aracena et al [4,5] to study the fixed points of Boolean networks.
Definition 7 A positive circuit in an interaction graph G is a non-empty sequence of edges, say
(j1 , s1 , i1 ), (j2 , s2 , i2 ), . . . , (jr , sr , ir ), such that: ik = jk+1 for 1 ≤ k < r (the sequence is a path); ir = j1 (the path is a circuit); the vertices jk are mutually distinct (the circuit is elementary); the product of the signs sk is positive (even number of negative edges).
Definition 8 [4] A positive feedback vertex set of an interaction graph G is a subset I ⊆ {1, . . . , n} such that each positive circuit of G has a vertex in I.
One can remark that: (1) The set of vertices of G is always a positive feedback vertex set of G; (2) The empty set is a positive feedback vertex set of G if and only if G has no positive circuit; (3) If G ′ is a subgraph of G then all the positive feedback vertex sets of G are positive feedback vertex sets of G ′ .
6

4

Positive circuits and attractors

As previously, let X = ni=1 Xi be the product of n finite intervals of integers of cardinality strictly greater than 1, and let F be a map from X to itself.
Q

We are interested by the relations between the map GF (defined on X ′ ) and the number of attractors in Γ(F ). The following theorem, presented in [3] as solution of a discrete version of the Thomas’ conjecture, gives such a relation.
Theorem 4 [3] If GF (x, v) has no positive circuit for all (x, v) ∈ X ′ , then
Γ(F ) has a unique attractor.
The following theorem extends the previous one by providing, without any condition on the map GF , an upper bound on the number of attractors in
Γ(F ) which only depends on GF .
Theorem 5 (main result) For each i ∈ {1, . . . , n}, let Ti (GF ) be the set of real numbers t for which there exists (x, v) ∈ X ′ such that t = xi + vi /2 and such that i belongs to a positive circuit of GF (x, v). Suppose I to be, for all
(x, v) ∈ X ′ , a positive feedback vertex set of GF (x, v). Then, the number of attractors in Γ(F ) is less than
Yh

i

|Ti (GF )| + 1 .

i∈I

Proof − We reason by induction on I. Suppose I to be, for any (x, v) ∈ X ′ , a positive feedback vertex set of GF (x, v).
Base case. If I = ∅ it means that there is no (x, v) ∈ X ′ such that GF (x, v) has a positive circuit. So, following Theorem 4, Γ(F ) has at most one attractors and the theorem holds.
Induction step. Suppose that I 6= ∅. The induction hypothesis is the following:
Induction hypothesis: Let F̃ be a map from X to itself. If I˜ is, for all (x, v) ∈
X ′ , a positive feedback vertex set of GF̃ (x, v), and if I˜ is strictly included
Q
in I, then Γ(F̃ ) has at most i∈I˜ |Ti (GF̃ )| + 1 attractors.
Without loss of generality, suppose that 1 ∈ I. Let P be the partition of
X1 whose elements Y are the maximal intervals of X1 (with respect to the inclusion relation) verifying
∀t ∈ Ti (GF ),

t < min(Y ) or
7

max(Y ) < t.

(1)

Remark that, by definition,
|P| = |T1 (GF )| + 1.

(2)

Let Y be any interval of P, and consider the map F̃ = (f˜1 , . . . , f˜n ) : X → X
defined by f˜i = fi for i > 1 and by

f˜1 (x) =

∀x ∈ X,





min(Y ) if f1 (x) < min(Y )




f1 (x)







if f1 (x) ∈ Y

max(Y ) if f1 (x) > max(Y ).

Then, for all x, y ∈ X, f˜i (x) < f˜i (y) ⇒ fi (x) ≤ f˜i (x) < f˜i (y) ≤ fi (y)

(i = 1, . . . , n).

(3)

Indeed, this is obvious for i > 1, and for i = 1 it is sufficient to remark that f˜1 (x) < f˜1 (y) ⇒ f˜1 (x) < max(Y ) ⇒ f1 (x) ≤ f˜1 (x), and that

f˜1 (x) < f˜1 (y) ⇒ min(Y ) < f˜1 (y) ⇒ f˜1 (y) ≤ f1 (y).

Now, we prove that, for all (x, v) ∈ X ′ ,
GF̃ (x, v) is a subgraph of GF (x, v).

(4)

Let (x, v) ∈ X ′ and suppose (j, s, i) to be an edge of GF̃ (x, v). According to (3), f˜ij (x, v) and fij (x, v) have the same sign (here s), and fi (x) and fi (x+vj ej ) are on both sides of xi +vi /2 since f˜i (x) and f˜i (x+vj ej ) are. In other words, (j, s, i)
is an edge of GF (x, v). So (4) is proved and, as an immediate consequence,
Ti (GF̃ ) ⊆ Ti (GF )

(i = 1, . . . , n).

(5)

Then, for all (x, v) ∈ X ′ , we have the following:
Vertex 1 belongs to none positive circuit of GF̃ (x, v).

(6)

Indeed, suppose, by contradiction, that vertex 1 belongs to a positive circuit of GF̃ (x, v). Let j be the predecessor of 1 in this circuit, and let t = x1 + v1 /2.
By definition, t ∈ T1 (GF̃ ) and from (5) it comes that t ∈ T1 (GF ). We then deduce, from (1) and the fact that the images of f˜1 are in Y , that f˜1 (x) and f˜1 (x + vj ej ) are not on both sides of t. In other words, there is no edge from j to 1 in GF̃ (x, v), a contradiction.
8

Let Ã be the set of attractors of Γ(F̃ ) and let
I˜ = I \ {1}.

(7)

Let (x, v) be any element of X ′ . Since I is a positive feedback vertex set of
GF (x, v) and since GF̃ (x, v) is a subgraph of GF (x, v), I is also a positive feedback vertex set of GF̃ (x, v). We then deduce from (6) that I˜ is a positive feedback vertex set of GF̃ (x, v). Since this holds for all (x, v) ∈ X ′ , by induction hypothesis,
Y
|Ã| ≤
|Ti (GF̃ )| + 1, i∈I˜

and from (5) we obtain:
|Ã| ≤

Y

|Ti (GF )| + 1.

(8)

i∈I˜

Now, let A be the set of attractors of Γ(F ), and let AY be the set of A ∈ A
containing a point x such that x1 ∈ Y . We claim that:
∀A ∈ AY , there exists Ã ∈ Ã such that Ã ⊆ A.

(9)

So let A ∈ AY , and consider the set Ā of x ∈ A such that x1 ∈ Y . We prove that Ā is a trap domain of Γ(F̃ ). Suppose (x, y) to be an edge of Γ(F̃ )
such that x ∈ Ā. By definition, there exists index i such that f˜i (x) 6= xi and y = x + sign(f˜i (x) − xi )ei . We consider two cases:
(1) Case i > 1. Then, y1 = x1 ∈ Y . In addition, f˜i (x) = fi (x) so (x, y) is an edge of Γ(F ). Hence y ∈ A (since x ∈ A) and we deduce that y ∈ Ā.
(2) Case i = 1. Suppose that x1 < f˜1 (x) (the proof is similar if x1 > f˜1 (x)).
Then, x1 < y1 ≤ f˜1 (x) and since x1 and f˜1 (x) are in Y we have y1 ∈ Y .
In addition, min(Y ) ≤ x1 < f˜1 (x) so x1 < f˜1 (x) ≤ f1 (x). Thus (x, y) is an egde of Γ(F ). Hence y ∈ A (since x ∈ A) and we deduce that y ∈ Ā.
Since y ∈ Ā in both cases, Ā is trap domain of Γ(F̃ ). Thus there exists at least one attractor Ã ∈ Ã such that Ã ⊆ Ā, and (9) holds since Ā ⊆ A.
Following (9), there exists a map H : AY → Ã such that H(A) ⊆ A for all
A ∈ AY . Since the attractors of Γ(F ) are mutually disjointed, the elements of AY are mutually disjointed, and we deduces that the images of H are also mutually disjointed. Consequently, H is an injection. So |AY | ≤ |Ã| and we deduce from (8) that
Y
|AY | ≤
|Ti (GF )| + 1.
i∈I˜

9

Since this inequality holds for all Y ∈ P, and since A = ∪Y ∈P AY , we have:
|A| ≤

X

|AY | ≤

Y ∈P

X hY

Y ∈P

i

|Ti (GF )| + 1 = |P|

i∈I˜

Y

|Ti (GF )| + 1.

i∈I˜

Using (2) and (7) we conclude: h

|A| ≤ |T1 (GF )| + 1

i Y

|Ti (GF )| + 1 =

i∈I˜

Y

|Ti (GF )| + 1.

i∈I


Corollary 1 If I is a positive feedback vertex set of G(F ), then the number of attractors in Γ(F ) and, in particular, the number of fixed points for F are
Q
less than i∈I |Xi |.
Proof − It is sufficient to apply Theorem 5 by noting that: (1) each GF (x, v)
is a subgraph of G(F ); (2) |Ti (GF )| + 1 ≤ |Xi |; (3) The number of fixed points for F is less than the number of attractors in Γ(F ).

Remark 4 The bound on the number of fixed points for F given Corollary 1
has been proved by Aracena et al [4,5] in the Boolean case and under the strong hypothesis that G(F ) does not contain both a positive and a negative edge from one vertex to another (that is, the entries of the Jacobian matrix of F are everywhere ≥ 0 or everywhere ≤ 0); see the Theorem 2 stated in the introduction.
Remark 5 Theorems 1 and 3 stated in the introduction are obtained from
Theorems 4 and 5 by noting that GF (x, v) is a subgraph of GF (x, v) and by using the points (2) and (3) in the proof of Corollary 1.

5

Comments

5.1 Influence of connections between positive circuits
Corollary 1 is sufficient to highlight the fact that: “A high level of connection between positive circuits leads to a small number of fixed points”. Suppose, for sake of simplicity, that all the intervals Xi are of cardinality q, and let r be the smallest number of vertices that a positive feedback vertex set of G(F )
can contain. Then, the smallest upper bound for the number of fixed points for F given by Corollary 1 is q r , and the more the positive circuits of G(F )
are connected, the more r is small. Indeed, let us say that a vertex represents a circuit when it belongs to this circuit. Then, r corresponds to the smallest number of vertices allowing the representation of each positive circuit. So,
10

the more the positive circuit are connected, the more it is possible to choose vertices representing a number of positive circuits, and the more r is small. For instance, r is always ≤ to the number p of positive circuits that G(F ) contains, but r < p whenever G(F ) has connected positive circuits, and in the extremal case where all the positive circuits of G(F ) share a same vertex, r = 1.
5.2 Thomas’ logical method
In practice, the dynamics of a gene network is often modeled from its interaction graph G, typically by using the well known Thomas’ logical method [7,8,9].
In few words, Thomas associates to G a finite state space X and describes the behavior of the interactions of G by logical parameters. Then, he deduces from the value of these parameters a map F from X to itself whose asynchronous state transition graph describes a possible dynamics for the network; see [12]
for a formal presentation.
This modeling method is coherent with our notion of interaction graph in the sense that, for all parameters values, the resulting map F has the property to be such that G(F ) is a subgraph of G [13]. So, thanks to Corollary 1, one can say, in the total absence of information on the value of the parameters, that following Thomas’ logical method, the number of attractors in the dynamics of the network is less than
µ(G, X) = min

I∈I(G)

Y

|Xi |,

i∈I

where I(G) is the set of smallest positive feedback vertex sets of G (with respect to the inclusion relation). This result is of practical interest since the value of the parameters is most often unknown and difficult to estimate, and since the number of attractors is an important feature of the dynamics of the network. For instance, if the network is known to control a differentiation process into k cell types, one often considers that the dynamics of the network has to contain at least k attractors. The bound µ(G, X) can then been used in order to check if the data of G and X is consistent with the presence of k attractors (there is inconsistence whenever µ(G, X) < k).
5.3 Feedback circuit functionality
Finally, Theorem 5 is related to one of the main concept raised by the Thomas’
logical method: the concept of feedback circuit functionality [11,8,9,14]. Roughly speaking, it has been observed that some inequality constraints on the logical parameters describing the behavior of the interactions of a positive (resp. negative) circuit of G often lead to a dynamics containing several attractors (resp.
11

describing oscillations). For that reason, when these constraints are satisfied, the corresponding circuit is said functional. Even if this notion is not well understand and often informally stated, it is often used in practice to establish the value of the logical parameters, see [15,16,17,18,19,20] for instance.
A natural formalization of the notion of functional circuit, also proposed in
[13,21], is the following: given a map F from X to itself whose interaction graph G(F ) is a subgraph of G, a circuit C of G is functional at (x, v) ∈ X ′
if C is a circuit of GF (x, v). It is then easy to see that the upper bound for the number of attractors given by Theorem 5 only depends on the localization
(inside X ′ ) and on the connections of the functional positive circuits of the system. In our knowledge, this is one of the first mathematical result relating the functional circuits of the system to its global dynamical properties (for relations between functional circuits and local dynamical properties, see the recent parer [21]).

Acknowledgement
I wish to thank Christophe Soulé for his precious suggestions.

References
[1] H. de Jong, Modeling and simulation of genetic regulatory systems: a literature review, Journal of Computational Biology, 9 (2002) 67-103.
[2] O. Delbrück, Discussion, in Unités Biologiques Douées de Continuités
Génétiques, Volume 33, Edition CNRS, Lyon, 1949.
[3] A. Richard, J.-P. Comet, Necessary conditions for multistationarity in discrete dynamical systems, Discrete Applied Mathematics, 155 (2007) 2403-2413.
[4] J. Aracena, J. Demongeot, E. Goles, Positive and negative circuits in discrete neural networks, IEEE Transactions of Neural Networks, 15 (2004) 77-83.
[5] J. Aracena, Maximum number of fixed points in regulatory boolean networks,
Bulletin of Mathematical Biology, 70 (2008) 1398-1409.
[6] R. Thomas, Boolean formalization of genetic control circuits, Journal of
Theoretical Biology, 42 (1973) 563-585.
[7] R. Thomas, R. d’Ari, Biological Feedback, CRC Press, 1990.
[8] R. Thomas, Regulatory Networks Seen as Asynchronous Automata : A logical
Description, Journal of Theoretical Biology, 153 (1991) 1-23.

12

[9] R. Thomas, M. Kaufman, Multistationarity, the basis of cell differentiation and memory. I. & II., Chaos, 11 (2001) 170-195.
[10] E.H. Snoussi, Qualitative dynamics of a piecewise-linear differential equations : a discrete mapping approach, Dynamics and stability of Systems, 4 (1989) 189207.
[11] E.H. Snoussi, R. Thomas, Logical identification of all steady states : the concept of feedback loop caracteristic states, Bulletin of Mathematical Biology, 55 (1993)
973-991.
[12] G. Bernot, J.-P. Comet, A. Richard, J. Guespin, Application of formal methods to biological regulatory networks: extending Thomas’ asynchronous logical approach with temporal logic, Journal of Theoretical Biology, 229 (2004) 339347.
[13] A. Richard, Modèle formel pour les réseaux de régulation génétique et influence des circuits de rétroaction, Ph.D. Thesis, University of Evry Val d’Essonne,
France, 2006.
[14] R. Thomas, D. Thieffry, M. Kaufman, Dynamical behaviour of biological regulatory networks - I Biological role of feedback loops an practical use of the concept of the loop-characteristic state, Bulletin Mathemathical Biology, 57
(1995) 247-276.
[15] D. Thieffry, R. Thomas, Dynamical behaviour of biological regulatory networks
- II Immunity control in bacteriophage lambda, Bulletin of Mathematical Biology,
57 (1995) 277-297.
[16] E. Muraille, D. Thieffry, O. Leo, M. Kaufman, Toxicity and neuroendocrine regulation of the immune response: a model analysis, Journal of Theoretical
Biology, 183 (1996) 285-305.
[17] L. Sánchez, J. van Helden, D. Thieffry, Establishement of the dorso-ventral pattern during embryonic development of drosophila melanogaster: a logical analysis, Journal of Theoretical Biology, 189 (1997) 377-389.
[18] L. Mendoza, D. Thieffry, E.R. Alvarez-Buylla, Genetic control of flower morphogenesis in Arabidopsis thaliana: a logical analysis, Bioinformatics, 15
(1999) 593-606.
[19] L. Sánchez, D. Thieffry, A logical analysis of the Drosophila gap-gene system,
Journal of Theoretical Biology, 211 (2001) 115-141.
[20] L. Sánchez, D. Thieffry, Segmenting the fly embryo: a logical analysis of the pair-rule cross-regulatory module, Journal of Theoretical Biology, 224 (2003)
517-537.
[21] E. Remy, P. Ruet, From elementary signed circuits to the dynamics of Boolean regulatory networks, Bioinformatics, 24 (2008) 220-226.

13

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
