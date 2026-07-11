---
source: "https://arxiv.org/abs/2505.06206v1"
title: "Constructing All Birthday 3 Games as Digraphs"
author: "Alexander Clow, Alfie Davies, Neil Anderson McKay"
year: "2025"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/2505.06206v1"
pdf: "https://arxiv.org/pdf/2505.06206v1"
captured_at: "2026-06-24T11:11:04Z"
updated_at: "2026-06-24T11:11:04Z"
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

# Constructing All Birthday 3 Games as Digraphs

- 著者: Alexander Clow, Alfie Davies, Neil Anderson McKay
- 年: 2025
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/2505.06206v1)
- ダウンロード: https://arxiv.org/pdf/2505.06206v1
- PDF: https://arxiv.org/pdf/2505.06206v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Recently, Clow and McKay proved that the Digraph Placement ruleset is universal for normal play: for all normal play combinatorial games $X$, there is a Digraph Placement game $G$ with $G=X$. Clow and McKay also showed that the 22 game values born by day 2 correspond to Digraph Placement games with at most 4 vertices. This bound is best possible. We extend this work using a combination of exhaustive and random searches to demonstrate all 1474 values born by day 3 correspond to Digraph Placement games on at most 8 vertices. We provide a combinatorial proof that this bound is best possible. We conclude by giving improved bounds on the number of vertices required to construct all game values born by days 4 and 5.

## PDF Text

arXiv:2505.06206v1 [math.CO] 9 May 2025

Constructing All Birthday 3 Games as Digraphs
1
Alexander Clow
, Alfie Davies
Neil Anderson McKay 3

2

, and

1

2

Department of Mathematics, Simon Fraser University
Department of Mathematics and Statistics, Memorial University of Newfoundland
3
Department of Mathematics and Statistics, University of New Brunswick

Abstract
Recently, Clow and McKay proved that the digraph placement ruleset is universal for normal play: for all normal play combinatorial games X, there is a digraph placement game G with G = X.
Clow and McKay also showed that the 22 game values born by day
2 correspond to digraph placement games with at most 4 vertices.
This bound is best possible. We extend this work using a combination of exhaustive and random searches to demonstrate all 1474 values born by day 3 correspond to digraph placement games on at most 8
vertices. We provide a combinatorial proof that this bound is best possible. We conclude by giving improved bounds on the number of vertices required to construct all game values born by days 4 and 5.

1

Introduction

Let G = (V, E) be a directed graph and let ϕ : V → {blue, red} be a (not necessarily proper) 2-colouring of G. The digraph placement game played on (G, ϕ), which we write simply as G when ϕ is clear from context, is described as follows: there are two players who play alternately, Left (blue)
and Right (red); on her turn, Left chooses a blue vertex u and deletes N + [u]
from G; on his turn, Right chooses a red vertex v and deletes N + [v] from
G. The digraph placement game resulting from a player choosing the vertex v (i.e. from the deletion of a vertex v and its out-neighbours) in a digraph placement game G is denoted G/v. When a sequence of vertices v1 , . . . , vk and their out-neighbours are deleted, we write the resulting game
G/[v1 , . . . , vk ]. If, at the beginning of a player’s turn, there are no vertices of
1

that player’s colour remaining, then that player is unable to move and thus loses (as is the normal play convention). The set of all digraph placement games is the ruleset digraph placement.

Figure 1: A digraph placement game equal to ↑ (left) and another equal to {2 | −2} (right). Blue vertices are circles, red vertices are squares, and pairs of arcs (u, v), (v, u) are shown with a bolded edge.
digraph placement is an example of a (normal play) combinatorial game theory ruleset as introduced by Conway in [7]. A normal play ruleset is one where there are two players who alternate turns, both players have perfect information about the game state and moves (the game involves no randomness), and if on a player’s turn they cannot move then they lose. For the rest of the paper we take game to mean an instance of a normal play ruleset. All games considered have finite numbers of options and end in finitely many turns, regardless of how the players move.
It is customary to call the players of a game Left and Right. A game X can be written as X = {X1 , . . . , Xk | Y1 , . . . , Yt } where L(X) := {X1 , . . . , Xk }
is the set of all games Left can move to from X (i.e. Left’s options), and
R(X) := {Y1 , . . . , Yt } is the analogous set for Right (i.e. Right’s options). It is a fundamental theorem of CGT that every game X has one of the following four outcomes: Left wins (playing first or second), Right wins (playing first or second), the first player wins, or the second player wins.
digraph placement was introduced in 2024 by Clow and McKay [6]
as a means of generalising some well-studied combinatorial game rulesets.
Rulesets that were generalised include, but are not limited to, nim, col, snort, poset games, node kayles, arc kayles, domineering, and teetering towers. For more on these rulesets, see [2, 3, 5, 13]. Like all of these rulesets, digraph placement is a placement game. Placement games are a special class of rulesets introduced by Brown et al. [4]. The digraph
2

placement ruleset is a natural instance of a placement game played on digraphs, and it is interesting because, as Clow and McKay [6] demonstrated, it is the first known universal ruleset that is a placement game (ruleset).
That is, for all normal play combinatorial games X, there is a digraph placement game G such that G = X.
Two games X and Y are equal, written X = Y , if for all games Z, the outcome of X + Z is the same as the outcome of Y + Z. The equivalence classes under this notion of equality are called values. Addition here is the disjunctive sum of games (see [14]).
The birthday of a game X, denoted by b(X), is the least (non-negative)
integer strictly greater than the birthday of all the options of X. As we consider finite games, birthdays are well defined. For more on the basic theory of disjunctive sums and equality we refer the reader to Siegel’s textbook [14].
The birthday of a value is the minimum birthday of all games with that value. Note, b(X) is sometime called the formal birthday of X to distinguish the birthday of X from the birthday of the value of X.
We consider the following problem: given a game X, what is the least integer f (X) such that there exists a digraph placement game G = X
on f (X) vertices. We focus on bounding the function F (b), which is the maximum of f (X) over all games X born by day b. Clow and McKay [6]
proved F (2) = 4 and posed determining F (3) as an open problem. The main contribution of this paper is proving F (3) = 8.
Theorem 1.1. For all games X born by day 3, there exists a digraph placement game G with at most 8 vertices such that G = X. Furthermore, there exists a game X with birthday 3 such that H ̸= X for all digraph placement games H with at most 7 vertices.
The paper is structured as follows. In Section 2 we explain our proof that
F (3) ≤ 8 which uses both exhaustive and random searches. The relevant output of these searches can be found in the ancillary files and acts as a certificate that F (3) ≤ 8. In Section 3, we prove F (3) ≥ 8 via a combinatorial argument, which demonstrates an explicit construction of a game requiring
8 vertices. In Section 4 we provide improved bounds on F (4) and F (5).

2

Constructing all Day 3 Games

Among its many uses, the nauty package [11] provides a suite of programs called gtools for generating non-isomorphic graphs. OEIS sequence A00059
[12] reports that there are 112 282 908 928 non-isomorphic digraphs on 7
3

unlabelled vertices with loops allowed but not multiple arcs. Using gtools, we generated all of these 7-vertex digraphs up to isomorphism, as well as all of those on at most 6 vertices.
There is a bijection between the set of digraphs on n unlabelled vertices with loops allowed but not multiple arcs, and the set of (not necessarily properly) 2-coloured digraphs on n unlabelled vertices with no loops and no multiple arcs. Such a bijection can be constructed simply by mapping a vertex with a loop to a blue vertex, and a vertex without a loop to a red vertex.
As such, with gtools, we generated all digraph placement positions that use at most 7 vertices. To exhaustively check through all of these positions, we wrote a program in Rust that converted each digraph into a combinatorial game and kept a hash table of all of the game values found so far. We used the gemau library [8] to work with the combinatorial game representations of the digraph placement positions.
We leveraged Rust’s concurrency model to improve performance by using multiple threads. This efficiency was very beneficial given the scale of our computations.
The exhaustive search of all diplace graphs using at most 7 vertices yielded 1455 out of the 1474 values born by day 3. To find the remaining values, we performed a targeted random sampling of graphs. In particular, we generated random digraphs on 8 vertices using the Erdős–Rényi–Gilbert model with various edge probabilities, focusing primarily around p = 0.5
(with colours chosen uniformly at random), and for each we computed its value. We generated at least 20 billion graphs in this way, but did not keep detailed logs of these searches.
Blue-red digraphs searched

New values found

Instances checked

At most 6 vertices
7-vertex
Random 8-vertex
Colour Isomorphic 8-vertex

1089
366
18
1

97 224 121
112 282 908 928
20 000 000 000
14 286 848

Table 1: Summary of the search for representatives of values born by day 3.
Each row describes one aspect of our search, with the first rows having been conducted first, and the last rows having been conducted last. All values considered are day 3.

4

Sampling random 8-vertex digraphs yielded 18 of 19 missing day 3 values.
A full list of these 19 values is given in the appendix. A significant portion of these values were found early in the random search, with a majority of the search being spent trying to find a small number of the missing values.
In particular, we were unable to find the value
{↑ + ∗, ↑, {1 | ∗, 0} | {0, ∗ | −1}, ↓, ↓ + ∗}
via our random search, which we call Z; in Section 3, we prove that Z
requires at least 8 vertices to represent.
In order to construct a digraph placement game on at most 8-vertices equal to Z, we performed an exhaustive search over those 8-vertex digraphs exhibiting an automorphism switching blue and red vertices. We call such graphs colour isomorphic. We generated this list of digraphs as follows: first, we constructed all 218 isomorphism classes for monochromatic digraphs on
4-vertices with gtools; then, for each of these 218 isomorphism classes, we take a representative H and add arcs in all possible ways from a blue copy of
H to a red copy of H (and vice-versa) such that there is an automorphism which switches the blue and red copies of H.
We treat the copies of H as labelled, and thus it is clear that there are
216 possible ways to configure the arcs joining the blue and red copies of
H. This of course overcounts the number of 8-vertex colour isomorphic digraph, but not by a significant margin. Finally, we checked the value of the 218 · 216 = 14 286 848 8-vertex colour isomorphic digraphs we generated.
For one example of a digraph equal to Z, see Figure 2.
The source code for both the exhaustive and random searches are available from the authors upon request. Thankfully, one does not have to perform the search again to validate the results, as we have an explicit list of digraphs that have the 1474 values born by day 3. Given one of these digraphs, it is trivial to manually check that it has the asserted value.
The digraphs we generated are given in the ancillary graphs.d6 file; their corresponding values are given (in the same order) in the games.txt file.
Digraphs are stored in the digraph6 format where a loop corresponds to a blue vertex and the absence of a loop corresponds to a red vertex. We provide a short script utilising gemau to automate the checking that each of the digraphs provided has the game value asserted, which the reader can find in src/main.rs (with instructions for running it in README.md).

5

3

Some Day 3 Games Require 8 Vertices

In Section 3, we demonstrate that there exists a game X with birthday 3 and f (X) ≥ 8. That is, X requires at least 8 vertices to be realised as a value of a digraph placement position. Technically, this was already shown by our exhaustive search of digraph placement games with at most 7 vertices, as we summarised in Section 2. However, reproducing the search requires significant computer resources, hence this search is not a good certificate that F (3) ≥ 8. To this end, we provide a ‘human’ proof that F (3) ≥ 8 which does not use computer assistance.
Providing multiple avenues to assert F (3) ≥ 8 is beneficial because demonstrating good lower bounds on f (X) and F (b) has proven difficult. We feel that our combinatorial argument is easier to generalise to more vertices than is exhaustive search on at most 7 vertices.
We assume the reader is familiar with the basics of combinatorial game theory. This includes familiarity with the values ↑, numbers, and nimbers, as well as knowledge regarding ordinal sums, and reversibility. We recommend
Lessons in Play [1] as a reference for readers unfamiliar with these concepts.
In this section, one birthday 3 game has the focus of our attention. We denote this game by
Z = {↑ + ∗, ↑, {1 | ∗, 0} | {0, ∗ | −1}, ↓, ↓ + ∗}.
As we mentioned in Section 2, there are 19 values that are not realisable with at most 7 vertices. So presumably we could have picked any of those to show F (3) ≥ 8. The reason we are choosing Z here is because of some useful properties that it exhibits, which we explore in the subsequent lemmas. One particularly useful feature is that Z is a symmetric form (i.e. Z = − Z).
Now we consider digraph placement games with at most one blue vertex.
Lemma 3.1. If G is a digraph placement game with at least one red vertex and no blue vertex, then G = k for a negative integer k.
Proof. Let G be a digraph placement game with no blue vertex. If G has exactly one red vertex, then G = −1. Assume now that G has n > 1 vertices.
For all red vertices u in G, the game G/u is a digraph placement game with no blue vertices and strictly less than n red vertices. If G/u has no vertices, then G/u = 0. Otherwise, G/u is a negative integer by induction.
So, Left has no options in G, and Right’s options are to 0 and/or a negative integer. Thus, G is a negative integer.
6

Lemma 3.2. If G is a Left-win digraph placement game with exactly one blue vertex, then G is a positive number and G ≤ 1.
Proof. Suppose G is a Left-win digraph placement game with exactly one blue vertex: call this vertex v. As G is Left-win, there must exist a Left option GL of G with GL ≥ 0. It is clear that GL = G/v and GL cannot have any blue vertices. So Lemma 3.1 implies GL has no red vertices either. Thus,
GL is the digraph placement game with no vertices, yielding GL = 0.
This implies (v, u) is an arc in G for every red vertex u.
Since G is Left-win, Right has no winning option. This implies that, for all red vertices u in G, we must have that G/u is a win either for Left or the next player. So, for all red vertices u in G, Left has an option in G/u. This implies that, for all red vertices u, there is no (u, v) arc in G.
Let G′ be the digraph placement game formed by removing v (and only v) from G. Then, the reader can easily verify G = 1 : G′ .
If G′ has no vertices, then clearly G = 1. Otherwise, G′ has no blue vertex but at least one red vertex, and so Lemma 3.1 implies G′ = k for some negative integer k. In this case, the Colon Principle implies
G=1:k =

1
.
2|k|

Thus, G is equal to a number satisfying 0 < G ≤ 1, as desired.
Next, we consider a lemma that deals with reversibility in digraph placement games with at most 2 blue vertices.
Lemma 3.3. If G is a digraph placement game with at most two blue vertices and X is the canonical form of G, then for all Left options X L of X, either G/v = X L or else G/[v, u, w] = X L , where v and w are blue vertices, and u is a red vertex.
Proof. Let G be a digraph placement game with at most two blue vertices and let X be the canonical form of G. Let X L be a fixed but arbitrary Left option of X. If there exists a blue vertex v such that G/v = X L , then we are satisfied. Suppose then for all blue vertices v in G, G/v ̸= X L . It follows that X L is not a Left option of G.
Since X is the canonical form of G, the literal form of X can be obtained from the literal form of G by removing dominated options and reversing reversible options. In particular, by removing dominated Left options, and through reversing Left options, G is transformed into a position, say G′ , with a Left option equal to X L . Removing a dominated option never gives a new
7

option to Left; it follows that the option equal to X L in L(G′ ) is obtained through reversing a Left option of G one or more times.
Thus, there exists a sequence G = G0 , G1 , G2 , . . . , G2k−1 = X L where
Gi ∈ R(Gi−1 ) for odd i, and Gj ∈ L(Gj−1 ) for even j ≥ 2. By our assumption that G contains at most 2 blue vertices, if G = G0 , G1 , G2 , . . . , G2k−1 = X L , then k ≤ 2. This is because for all even j ≥ 2, Gj is obtained from Gj−1 by deleting a blue vertex. Here every Gt is a subgraph of Gt−1 . So all Gt are subgraphs of G = G0 .
Therefore, either G has a Left option to X L which can be written G/v =
L
X for a blue vertex v, or there exists a sequence G = G0 , G1 , G2 , G3 = X L
where G1 ∈ L(G0 ), G2 ∈ R(G1 ), G3 ∈ L(G2 ). Let G = G0 , G1 , G2 , G3 = X L
be such a sequence, then there exist blue vertices v and w, and a red vertex u such that G1 = G/v, G2 = G/[v, u], and X L = G3 = G/[v, u, w]. This concludes the proof.
We now study the Left options of Z.
Lemma 3.4. If G is a digraph placement game and G = {1 | 0, ∗}, then
G contains at least two blue vertices.
Proof. Let G be a digraph placement game equal to {1 | 0, ∗}. If G has an option to 1, call it G/v, then G/v contains a blue vertex by Lemma 3.1.
If G/v = 1 is a Left option of G, then v is a blue vertex that is not in G/v.
Hence, if G has a Left option to 1, then G has at least two blue vertices.
Otherwise, G has no option to 1. Then Lemma 3.3 implies there exist blue vertices v and w, and a red vertex u, such that G/[v, u, w] = 1. Again, this implies G has at least two blue vertices. Therefore, if G is a digraph placement game equal to {1 | 0, ∗}, then G contains at least two blue vertices.
Lemma 3.5. If G is a digraph placement game and G = ↑, then G
contains at least two blue vertices.
Proof. Let G be a digraph placement game equal to ↑. Then, G is Leftwin and G is not a number. As G is Left-win, Lemma 3.1 implies G requires at least one blue vertex. As G is Left-win and not a number, Lemma 3.2
implies G does not have exactly one blue vertex. Thus, G has at least two blue vertices.
Lemma 3.6. If G is a digraph placement game and G = ↑ + ∗, then
G contains at least two blue vertices. If G additionally contains exactly two blue vertices, then the blue vertices of G do not form an independent set.
8

Proof. Let G be a digraph placement game equal ↑ + ∗. Note that
{0, ∗ | 0} is the canonical form of ↑ + ∗. We begin by proving G has at least
2 blue vertices.
If G has a Left option to ∗, call it G/v, then G/v contains a blue vertex by Lemma 3.1. If G/v = ∗ is a Left option of G, then v is a blue vertex that is not in G/v. Hence, if G has an option to ∗, then G has at least two blue vertices.
Otherwise, G has no option to equal to ∗. Then Lemma 3.3 implies there exists a blue vertices v and w, and a red vertex u, such that G/[v, u, w] = ∗.
Since ∗ is not a negative integer Lemma 3.1 implies G/[v, u, w] has a blue vertex. Thus, if Left has no option equal to ∗ in G, then G has at least three blue vertices. Therefore, if G is a digraph placement game equal to ↑ + ∗, then G contains at least two blue vertices.
Now suppose H is a digraph placement game such that H contains exactly two blue vertices and H = ↑ + ∗. Then, from the previous paragraph, there exists a blue vertex v in H such that H/v = ∗. Since H = ↑ + ∗,
{0, ∗ | 0} is the canonical form of H. Suppose for the sake of contradiction that the blue vertices of H form an independent set. We consider the following two cases:
1. There exists a blue vertex w in H such that H/w = 0.
Since the blue vertices of G form an independent set, v is a vertex in H/w. Hence, H/w has exactly one blue vertex. As H/w = 0 Left has no winning move in H/w. Then, there exists a red vertex u in
H/w such that (v, u) is not an arc in H/w. Otherwise, H/[w, v] is a zero vertex graph, implying that Left has a winning option in H/w.
Furthermore, since u is a vertex in H/w, (w, u) is not an arc in H.
Thus, (v, u) and (w, u) are not arcs in H.
Now consider H/v which equals ∗. Since (v, u) is not an arc in H, u is a red vertex in H/v. Given (w, u) is not an arc in H, we note that (w, u) is not an arc in H/v. This implies u is a vertex in H/[v, w].
Hence, as u is a red vertex in H/[v, w], while H/[v, w] contains no blue vertices by our assumption H contains exactly two blue vertices.
Since H/[v, w] contains at least one red vertex, u, and no blue vertices,
Lemma 3.1 implies H/[v, w] < 0.
But this is a contradiction, since H/v = ∗, while H/v has exactly one blue vertex w implies that H/[v, w] = 0. So we have shown
0 = H/[v, w] < 0 which is absurd.
It follows that if Left has an option to 0 in H, then we contradict that
9

H = ↑ + ∗, or we contradict that H has exactly two blue vertices, or we contradict that the blue vertices of H form an independent set.
2. No Left option of H is equal to 0.
Then Lemma 3.3 implies H must have 0 as a reversible Left option. Let w be the second blue vertex in H, recalling we have assumed H/v = ∗.
As 0 is not an option of any Right option of ∗ and H/v = ∗, there exists a red vertex u ∈
/ N + (w) such that 0 is a Left option of H/[w, u]
and H/[w, u] ≤ H. Let u be such a vertex.
Since 0 is a Left option of H/[w, u], the game H/[w, u] must have a blue vertex. This blue vertex must be v. Hence, H/[w, u, v] = 0. Then
Lemma 3.1 and the fact that H/[w, u] has only one blue vertex, implies that for every red vertex z in H/[w, u], (v, z) is an arc in H/[w, u].
Since H/[w, u] ≤ H = ↑ + ∗, where ↑ + ∗ is next player win, Right has a winning option in H/[w, u]. If x is a red vertex such that
H/[w, u, x] ≤ 0, then (x, v) is an arc in H/[w, u]. Otherwise, the vertex v exists in H/[w, u, x] implying Left can win moving first in
H/[w, u, x] by delete v and all remaining red vertices, contradicting that H/[w, u, x] ≤ 0. Since Right has a winning move in H/[w, u] there exists a red vertex x such that H/[w, u, x] ≤ 0. Let x be such a red vertex.
Then, H/[w, x] has no blue vertices, since v is the only blue vertex in
H/w and (x, v) is an arc in H. Hence, Lemma 3.1 implies H/[w, x]
is equal to 0 or a negative integer. It follows that Right-wins H/w moving first.
But this implies Left’s only options in H are to H/v = ∗ and to H/w where Right-wins moving first in H/w. Thus, Left loses moving first in H. This is a contradiction as we have assumed H = ↑ + ∗, which is next player win.
Therefore, we have shown that if H is a digraph placement game satisfying that H contains exactly two blue vertices, H = ↑ + ∗, and the blue vertices of H form an independent set, then this leads to a contradiction. It follows no such H exists. This concludes the proof.
Theorem 3.7. There exists a game X with birthday 3 such that X = −X
and f (X) ≥ 8.
Proof. Recall that Z is {↑ + ∗, ↑, {1 | ∗, 0} | {0, ∗ | −1}, ↓, ↓ + ∗}. It is easy to verify that Z has birthday 3, is in canonical form, and is a symmetric form (i.e. Z = − Z). We show that f (Z) ≥ 8.
10

Suppose for the sake of contradiction f (Z) < 8. Then, there exists a digraph placement game H = Z with at most 7 vertices. Since H has at most 7 vertices, there cannot be at least 4 blue vertices and at least 4
red vertices in H. If there are at most 3 blue vertices in H, let G ∼
= H = Z.
If there are at most 3 red vertices in H, let G ∼
= −H = − Z = Z. Thus, digraph placement game G has at most 3 blue vertices.
If G has at most 2 blue vertices, then Left cannot move in G to ↑ by
Lemma 3.5. Moreover, ↑ is not be a Left option of G that can be obtained by reversibility since any game G/[v, u, w] where v, w are blue vertices and u is red, would have less than 2 blue vertices. As this contradicts G = Z, we suppose G has exactly 3 blue vertices.
Every Left option of G obtained from reversibility is of the form G′ =
G/[v, u, w] where v and w are blue vertices and u is a red vertex, or of the form G′ = G/[v, u, w, x, y] where v, w, y are blue vertices and u, x are red vertices. Since G has exactly three blue vertices, any such G′ has at most one blue vertex. So by Lemma 3.4 G cannot have {1 | 0, ∗} as a Left option from reversibility. Similarly, by Lemma 3.5 G cannot have ↑ as a Left option from reversibility, and by Lemma 3.6 G cannot have ↑ + ∗ as a Left option from reversibility.
Then, G cannot obtain any of the Left options of Z via reversibility. It follows every Left option of G is equal to a Left option of Z. Hence, there are blue vertices v1 , v2 , v3 in G where
• G/v1 = {1 | 0, ∗}, and
• G/v2 = ↑, and
• G/v3 = ↑ + ∗.
Since each of {1 | 0, ∗}, ↑ + ∗, and ↑ require at least 2 blue vertices to be realized as digraph placement games, G having exactly 3 blue vertices implies {v1 , v2 , v3 } forms an independent set.
Observe that G/v3 is a subgraph of G with exactly two blue vertices.
Furthermore, we have assumed without loss of generality that G/v3 = ↑ + ∗.
Since {v1 , v2 , v3 } is an independent set in G, {v1 , v3 } is an independent set in G/v3 . But this implies G/v3 is a digraph placement game equal to
↑ + ∗, with exactly two blue vertices, where the blue vertices in G/v3 form an independent set. This is a contradiction given that Lemma 3.6 states no such digraph placement game exists.
Therefore, we have contradicted that there exists a digraph placement game equal to Z on at most 7 vertices. It follows that f (Z) ≥ 8 as desired.
11

Figure 2: An 8-vertex digraph placement game equal to Z.
To see a digraph that shows f (Z) = 8 we provide the reader Figure 2.
We additionally observe that there is an automorphism of the underlying digraph of G that switches red and blue vertices. That is, G and −G are isomorphic as 2-coloured digraphs. It is immediate that this is only possible if G = −G.

4

Games Born Day 4 and Day 5

Now that we have proven F (3) = 8 we turn our attention to the values of
F (4) and F (5). Unfortunately, we are not able to prove an exact value for either constant. However, we are able to improve the bounds previously given in [6]. To see the bounds from [6] consider Table 2.
b

≤ F (b)

F (b) ≤

0
1
2
3
4
5

0
2
4
5
6
7

0
2
4
74
13 439
1.08 · 10189

Table 2: Bounds on F (b) for b ≤ 5 from [6].
We improve the bounds in Table 2 by leveraging a number of existing bounds in the literature, along with the fact that we have shown F (3) = 8.
12

Our upper bounds rely on a bound by Suetsugu [15] on the number of game values with birthday at most 4, and a recursive bound on F (b) given in
[6]. Meanwhile, our lower bounds combine the same work by Suetsugu [15], earlier lower bounds on the number of game values with birthday at most b by Wolfe and Fraser [16], and counting the number of isomorphism classes of
2-coloured digraphs on n vertices.
We let g(b) denote the number of game values, that is equivalence classes of games, that exist with birthday at most b. The currently best bounds on g(4) are included below.
Lemma 4.1 ([15, Corollary 2, Theorem 2]). 294 < g(4) ≤ 4 · 10184 .
Next, we give a useful lower bound on g(b + 1) in terms of g(b) and g(b − 1).
Lemma 4.2 ([16, Theorem 8]). For all b ≥ 0,


g(b + 1) ≥ 8g(b − 1) − 4



g(b)−2

2 2g(b−1)−1

−1



−1 .

Combining these results, and recalling that g(3) = 1474, we give a lower bound for g(5).
24 )

Lemma 4.3. g(5) > 26.7(10

.

Proof. Recall that g(3) = 1474 and by Lemma 4.1 that g(4) > 294 . Thus,
Lemma 4.2 implies


g(5) ≥ 8g(3) − 4



 294 −2

g(4)−2

2 2g(3)−1

> 11788 2 2947 −1 − 1
24 )

≥ 26.7(10

−1



−1



.

This concludes the proof.
For a non-negative integer n let D(n) be the number of 2-coloured digraphs on n unlabelled vertices with no more than one arc with the same start and end vertex. As noted in Section 2, there is a bijection between this set of n-vertex digraphs and the set of n vertex digraph placement games.
This means that D(n) counts the number of n-vertex digraph placement games, up to graph isomorphism. The list of small values for D(n) is OEIS
sequence A000595. See Table 3.
13

n

D(n)

0
1
2
3
4
5
6
7
8
9
10
11
12
13
14

1
2
10
104
3044
291 968
96 928 992
112 282 908 928
458 297 100 061 728
6 666 621 572 153 927 936
349 390 545 493 499 839 161 856
66 603 421 985 078 180 758 538 636 288
46 557 456 482 586 989 066 031 126 651 104 256
120 168 591 267 113 007 604 119 117 625 289 606 148 096
1 152 050 155 760 474 157 553 893 461 743 236 772 303 142 428 672
Table 3: A list of small values of D(n).

Additionally, an exact formula for D(n) was given by Davis [9]. Using
2
this formula, McIlroy [10] showed that that the ratio of D(n) and 2n /n!
tends to 1 as n tends to infinity. When we consider day 5 games, taking the following upper bound on D(n) is sufficient for our purposes.
2

Lemma 4.4. For all n ≥ 0, D(n) ≤ 2n .
Proof. Trivially, for all n ≥ 0, D(n) is bounded above by the number of labelled 2-vertex coloured digraphs with no more than one arc with the same
2
start and end vertex. We claim that for all n ≥ 0, there are exactly 2n digraphs of this type.
Notice there are n(n − 1) ordered pairs (u, v) for u, v ∈ {1, . . . , n}. Each ordered pair (u, v) corresponds to a possible arc. Hence, there are 2n(n−1)
labelled 1-vertex colour digraphs with no more than one arc with the same start and end vertex. From here it is easy to see that there are 2n(n−1) · 2n =
2
2n digraphs with no more than one arc with the same start and end vertex
2
whose vertices are 2-coloured. Therefore, D(n) ≤ 2n as desired.
For games X and Y , we say X ≤ Y if Right wins X − Y moving second.
Under this partial order, X ≤ Y and Y ≤ X if and only if X = Y . For more on the partial order of games, see [14].
14

Next, we recall an upper bound on F (b + 1) which depends on the size of a largest anti-chain of games with birthday at most b, and F (b). This upper bound is taken from [6].
Lemma 4.5 ([6, Lemma 5.1]). Let b ≥ 0. Then
F (b + 1) ≤ 2a(b)(F (b) + b + 1) + 5b + 8
where a(b) is the order of a largest anti-chain of games born by day b.
In order to apply the bound in Lemma 4.5 we require information about a(b). The size of a largest anti-chin of games born by 3 was proven to be 86
in [15]. We take g(4) as a trivial upper bound for a(4).
Lemma 4.6 ([15, Lemma 3, Theorem 2]). Letting a(b) be the order of a largest anti-chain of games born by day b, a(3) = 86 and a(4) ≤ 10184 .
We are now prepared to prove the main result of this section.
Theorem 4.7. The function F satisfies
• 11 ≤ F (4) ≤ 2087, and
• 2.58(1012 ) < F (5) < 10187.63 .
Proof. We begin by proving the stated lower bounds. Since each 2-coloured digraph corresponds to exactly one digraph placement game,
F (b)

X

D(n) ≥ g(b)

n=0

with equality if and only if each every digraph placement game on at most n vertices has a distinct value, which is false for digraph placement games on two vertices. Hence, Lemma 4.1 and Lemma 4.3 imply
F (4)

X

F (5)

D(n) > 2

94

≥ 10

28.2

and

n=0

X

24 )

D(n) > 26.7(10

i=0

respectively. Given Table 3 the reader can easily verify 11 is the smallest integer such that
11
X

D(n) > 1028.2 .

n=0

15

2

Thus, F (4) ≥ 11 as required. By Lemma 4.4, for all n ≥ 0, D(n) ≤ 2n .
Hence,
F (5)

X

2

24 )

2n > 26.7(10

.

n=0

Inspecting this inequality and taking logarithms gives
F (5)

&

F (5)2 + 1 =

log2

X

n2

!'

2

i=0

> 6.7(1024 ).
Hence, F (5) ≥ 2.58(1012 ) as required.
Next we derive upper bounds on F (4) and F (5). To do this we rely on
Lemma 4.5 and Lemma 4.6. Applying both of these Lemmas, in addition to
F (3) = 8 per Theorem 1.1, gives
F (4) ≤ 2(86)(8 + 3 + 1) + 5(3) + 8
= 2087
as desired. Similarity, by Lemma 4.5 and Lemma 4.6,
F (5) ≤ 2(10184 )(2087 + 4 + 1) + 5(4) + 8
= 4184(10184 ) + 28
< 4185(10184 )
< 10187.63 .
This completes the proof.

5

Final remarks

Here we discuss some questions for future work. Of course, the most obvious problem is, can we improve the bounds on F (4) and F (5)? Given F (5) > 1012
we are particularity interested in whether F (4) is small, say at most 100, as it would be interesting for F (b) to jump from having two digits to at least
12 in only one step. We pose the following easier problem.
Open Problem 5.1. Determine if F (4) is larger or smaller than 100.

16

Next we recall Problem 6.3 from [6], which asks to find a family of games
S, such that there exists a constant c > 0, where if Sb is the set of all games in S born by day b, then min f (X) ≥ cF (b).

X∈Sb

Given the properties of the game Z used in the proof of Theorem 3.7, we propose that some subset of the games X such that X = −X could be such a family. More strongly we conjecture the following.
Conjecture 5.2. For all integers b ≥ 0, there exists a game X with birthday b such that X = −X and f (X) = F (b).
This conjecture is true for all b ≤ 3. To see this, take the games
0, ∗, {1 | −1}, and Z, for the cases b = 0, 1, 2, and 3 respectively.
Next, we reflect on the fact that to prove Theorem 3.7 we relied on showing certain values do not exist in digraph placement if there is exactly one blue vertex. Could this class be characterised?
Open Problem 5.3. Characterise the values found in digraph placement with exactly one blue vertex.
One might also ask, what values do not appear in digraph placement games with at most k blue vertices? Is such a result useful in proving a better lower bound for F (4)?

References
[1] Michael Albert, Richard Nowakowski, and David Wolfe. Lessons in play.
Taylor & Francis, London, England, 2 edition, January 2023.
[2] Elwyn R. Berlekamp, John H. Conway, and Richard K. Guy. Winning ways for your mathematical plays. Vol. 1. A K Peters, Ltd., Natick, MA, second edition, 2001.
[3] Hans L. Bodlaender and Dieter Kratsch. Kayles and nimbers. J.
Algorithms, 43(1):106–119, 2002. doi:10.1006/jagm.2002.1215.
[4] Jason I. Brown, Danielle Cox, Andrew H. Hoefel, Neil McKay, Rebecca
Milley, Richard J. Nowakowski, and Angela A. Siegel. A note on polynomial profiles of placement games. In Games of no chance 5, volume 70 of Math. Sci. Res. Inst. Publ., pages 243–258. Cambridge
Univ. Press, Cambridge, 2019.
17

[5] Alexander Clow and Neil McKay. Ordinal sums of numbers, 2023. URL: https://arxiv.org/abs/2305.16516, arXiv:2305.16516.
[6] Alexander Clow and Neil A McKay. Digraph placement games, 2025.
URL: https://arxiv.org/abs/2407.12219, arXiv:2407.12219.
[7] J. H. Conway. On numbers and games. A K Peters, Ltd., Natick, MA, second edition, 2001.
[8] Alfie Davies. gemau, September 2024. URL: https://github.com/
alfiemd/gemau.
[9] Robert L. Davis. The number of structures of finite relations. Proc.
Amer. Math. Soc., 4:486–495, 1953. doi:10.2307/2032159.
[10] M. D. McIlroy. Calculation of numbers of structures of relations on finite sets. Technical Report 17, Massachusetts Institute of Technology,
1955.
[11] Brendan D. McKay and Adolfo Piperno. Practical graph isomorphism, ii.
Journal of Symbolic Computation, 60:94–112, 2014. URL: https://
www.sciencedirect.com/science/article/pii/S0747717113001193, doi:
10.1016/j.jsc.2013.09.003.
[12] OEIS Foundation Inc. The On-Line Encyclopedia of Integer Sequences,
2025. Published electronically at http://oeis.org.
[13] Thomas J. Schaefer. On the complexity of some two-person perfectinformation games. J. Comput. System Sci., 16(2):185–225, 1978. doi:
10.1016/0022-0000(78)90045-4.
[14] Aaron N. Siegel. Combinatorial game theory, volume 146 of Graduate
Studies in Mathematics. American Mathematical Society, Providence,
RI, 2013. doi:10.1090/gsm/146.
[15] Koki Suetsugu. Improving upper and lower bounds of the number of games born by day 4, 2024. To appear in Games of no chance 6, volume 71 of Math. Sci. Res. Inst. Publ., Cambridge Univ. Press. URL: https://arxiv.org/abs/2208.13403v2, arXiv:2208.13403.
[16] David Wolfe and William Fraser. Counting the number of games. Theoretical Computer Science, 313(3):527–532, 2004. Algorithmic Combinatorial Game Theory. URL: https://www.sciencedirect.com/science/
article/pii/S0304397503005991, doi:10.1016/j.tcs.2003.05.001.
18

Day 3 Values not found on at most 7 vertices
The following are the 19 values born on day 3 we did not find in our search of 7-vertex digraphs.
1. {↓, ↓ + ∗, ±1 | −2}
2. {2 | ±1, ↑ + ∗, ↑}
3. {↓, ↓ + ∗, ±1 | −1, −1 + ∗}
4. {1 + ∗, 1 | ±1, ↑ + ∗, ↑}
5. {∗, ∗2, 0 | {∗, 0 | −1}, ↓, ↓ + ∗}
6. {↑ + ∗, ↑, {1 | 0, ∗} | 0, ∗2, ∗}
7. {0, ↑ + ∗, {1 | ∗, 0} | {∗, 0 | −1}, ↓, ↓ + ∗}
8. {↑ + ∗, ↑, {1 | 0, ∗} | {0, ∗ | −1}, ↓ + ∗, 0}
9. {0, ↑ + ∗, {1 | ∗, 0} | ∗, {∗, 0 | −1}, ↓}
10. {↑, {1 | 0, ∗}, ∗ | {0, ∗ | −1}, ↓ + ∗, 0}
11. {∗, ↑, {1 | ∗, 0} | {∗, 0 | −1}, ↓, ↓ + ∗}
12. {↑ + ∗, ↑, {1 | 0, ∗} | {0, ∗ | −1}, ↓, ∗}
13. {↑ + ∗, ↑, {1 | ∗, 0} | ↓, {0 | −1}}
14. {{1 | 0}, ↑ | {0, ∗ | −1}, ↓, ↓ + ∗}
15. {↑ + ∗, ↑, {1 | ∗, 0} | {∗ | −1}, {−1 | 0}, {0 | −1}}
16. {{1 | 0}, {0 | 1}, {1 | ∗} | {0, ∗ | −1}, ↓, ↓ + ∗}
17. {∗, ↑, {1 | ∗, 0} | {0, ∗ | −1}, ↓, ∗}
18. {0, ↑ + ∗, {1 | ∗, 0} | {0, ∗ | −1}, ↓ + ∗, 0}
19. {↑ + ∗, ↑, {1 | ∗, 0} | {0, ∗ | −1}, ↓, ↓ + ∗}

19

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
