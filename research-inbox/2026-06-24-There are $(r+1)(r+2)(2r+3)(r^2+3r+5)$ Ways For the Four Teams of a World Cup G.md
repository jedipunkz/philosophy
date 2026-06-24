---
source: "https://arxiv.org/abs/1407.1919v1"
title: "There are $(r+1)(r+2)(2r+3)(r^2+3r+5)$ Ways For the Four Teams of a World Cup Group to Each Have $r$ Goals For and $r$ Goals Against [Thanks to the Soccer Analog of Prop. 4.6.19 of Richard Stanley's (Classic!) EC1]"
author: "Shalosh B. Ekhad, Doron Zeilberger"
year: "2014"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/1407.1919v1"
pdf: "https://arxiv.org/pdf/1407.1919v1"
captured_at: "2026-06-24T11:10:47Z"
updated_at: "2026-06-24T11:10:47Z"
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

# There are $(r+1)(r+2)(2r+3)(r^2+3r+5)$ Ways For the Four Teams of a World Cup Group to Each Have $r$ Goals For and $r$ Goals Against [Thanks to the Soccer Analog of Prop. 4.6.19 of Richard Stanley's (Classic!) EC1]

- 著者: Shalosh B. Ekhad, Doron Zeilberger
- 年: 2014
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/1407.1919v1)
- ダウンロード: https://arxiv.org/pdf/1407.1919v1
- PDF: https://arxiv.org/pdf/1407.1919v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

This short tribute to the guru of Enumerative and Algebraic Combinatorics started out when one the authors(DZ) attended the Stanely@70 conference, that took place at the same time as the preliminary stage of the 2014 World Cup. It states a surprising application of an analog of Richard Stanley's famous theorem about the enumeration of magic squares to the enumeration of possible outcomes in a World Cup Group.

## PDF Text

1
There are 30
(r + 1)(r + 2)(2r + 3)(r2 + 3r + 5) Ways For the Four Teams of a World Cup Group to Each Have r Goals For and r Goals Against
[Thanks to the Soccer Analog of Prop. 4.6.19 of Richard Stanley’s (Classic!) EC1]

arXiv:1407.1919v1 [math.CO] 8 Jul 2014

By Shalosh B. EKHAD and Doron ZEILBERGER

Dedicated to Richard Peter Stanley who just turned “number of ways for a simple
1D Drunkard to return home after 8 steps”-years-old

Proof of the Statement in the Title: Calling this quantity S4 (r), and using the case n = 4 of the Soccer analog of Prop. 4.6.19 of [EC1] (see Comment 1 below), we see that
• S4 (r) is a polynomial of degree 5
• S4 (−1) = 0

,

S4 (−2) = 0

• S4 (−3 − r) = −S4 (r)

.

.

.

Hence, it suffices to check the statement at the two values r = 0 and r = 1. Obviously S4 (0) = 1
(there is just one all-zero 4 × 4 matrix), and almost obviously, S4 (1) = 9 (the number of derangements of length 4 equals 9). QED!
Comments
1. The same method of proof that guru Richard Stanley used to prove Prop. 4.6.19 of [EC1] (first edition; it is Prop. 4.6.2 in the second edition.) [the first edition is missing (−1)n−1 in front of Hn (r)]
yields:
The Soccer analog of Prop. 4.6.19 of EC1
Let Sn (r) be the number of ways n Soccer teams, playing a round-robin tournament, each scored a total of r “Goals For” (GF), and a total of r “Goals Against” (GA), (in other words the number of n × n magic squares whose diagonal entries are all 0 (no team plays against itself!)).
For each fixed integer n ≥ 3, the function Sn (r) is a polynomial in r of degree n2 − 3n + 1. Since it is a polynomial, it can be evaluated at any (not necessarily positive) integer, and we have
Sn (−1) = Sn (−2) = . . . = Sn (−n + 2) = 0 ,
Sn (−(n − 1) − r) = −Sn (r) .
2. We got the idea for this tribute to our beloved guru when we attended the Stanley@70 conference, held at MIT, June 23-27, 2014, at the same time as the preliminary Group stage of the World Cup
2014.
1

3. While the proof of the statement of the title did not require any computer, the analogous result for n = 5 is already beyond the scope of mere humans (or the human would have to be very stupid to spend time on it). We have
S5 (r) =

1
(r + 1) (r + 2) (r + 3) ·
241920


43 r 8 + 688 r 7 + 4934 r 6 + 20680 r 5 + 55907 r 4 + 101272 r 3 + 123436 r 2 + 96240 r + 40320

.

Of course S3 (r) = r + 1 (why?).
To see S6 (r), please go to: http://www.math.rutgers.edu/~zeilberg/tokhniot/oGOALS1 .
One can get this, and (potentially!) infinitely more results, using the Maple package GOALS
available directly from: http://www.math.rutgers.edu/~zeilberg/tokhniot/GOALS , or via the front of this article: http://www.math.rutgers.edu/~zeilberg/mamarim/mamarimhtml/worldcup.html .
In particular, procedure MagicPolAG(n,r,Sr,Sc,B,A) can find polynomial expressions for the number of n × n matrices with non-negative integer entries, whose row-sums are
(r + Sr[1], . . . , r + Sr[n]) , and whose column sums are
(r + Sc[1], . . . , r + Sc[n])

,

for any fixed numerical vectors Sr, Sc, of length n (of course they have to add-up to the same number), and any assignment A where some entries are fixed beforehand, where B denotes “wildcard”. See the on-line help there. See the above-mentioned front of this article for numerous sample input and output files.
4. The sequence for S4 (r) is already in Neil Sloane’s OEIS, (see [OEIS1]), but for different reasons!
Can you find a bijection? On the other hand, the sequence for S5 (r) is not (yet!) there:
1, 44, 870, 9480, 68290, 365936, 1573374, 5709120, 18107760, 51488800, 133748186, 321979164, . . .

,

but we are sure that very soon it will!
5. Another, even more useful, Maple package accompanying this article is WorldCup, one of whose many procedures is ‘Ptor’, that finds all the possible scenarios that lead to a given score board, consisting of the “Goals For”, “Goals Against”, and “PTS”, vectors. (Recall that a win yields 3
points, a draw, 1 point, and a loss, 0 points.).
The number of possible scenarios for groups A-H were as follows
2014: 8, 32, 7, 3, 13, 3, 12, 3

.

2010: 2, 3, 2, 2, 3, 3, 1, 2 (in particular, the score-board for Group G uniquely determined the individual game scores).
2

2006: 11, 1, 5, 3, 2, 6, 1, 8 (in particular, the score-boards for Groups B and G uniquely determined the individual game scores).
2002: 1, 12, 5, 5, 1, 2, 3, 1 (in particular, the score-boards for Groups A, E, H uniquely determined the individual game scores).
1998: 9, 6, 4, 3, 3, 3, 1, 1 (in particular, the score-boards for Groups G, H uniquely determined the individual game scores).
6. The Israeli daily Yedioth Ahronoth published, at the start of the 2014 World Cup, a “soccer puzzle” where the above-mentioned vectors, GF, GA, and PTS, are given, and the solver has to reconstruct the scores of the individual matches. The above-mentioned Maple package WorldCup has a procedure, Khida, that makes up such puzzles, and another procedure, Sefer1, that creates challenging puzzle books. For some samples, see http://www.math.rutgers.edu/~zeilberg/tokhniot/oWorldCupi

,

for 1 ≤ i ≤ 4. Enjoy!

7. Happy birthday Richard, and keep up the good work! (and continue to practice what Phil
Hanlon calls the “the three H’s”).
References
[OEIS1] Neil J.A. Sloane, Sequence A061927, http://oeis.org/A061927.
[EC1] Richard P. Stanley, “Enumerative Combinatorics, volume I”, first edition: Wadsworth &
Brooks/Cole, 1986; second edition: Cambridge University Press, 2011.
Available on-line (viewed July 7, 2014): http://math.mit.edu/~rstan/ec/ec1.pdf .

Shalosh B. Ekhad, c/o D. Zeilberger, Department of Mathematics, Rutgers University (New Brunswick),
Hill Center-Busch Campus, 110 Frelinghuysen Rd., Piscataway, NJ 08854-8019, USA.
Doron Zeilberger, Department of Mathematics, Rutgers University (New Brunswick), Hill CenterBusch Campus, 110 Frelinghuysen Rd., Piscataway, NJ 08854-8019, USA.
url: http://www.math.rutgers.edu/~zeilberg/ .
Email: zeilberg at math dot rutgers dot edu .

EXCLUSIVELY PUBLISHED IN THE PERSONAL JOURNAL OF SHALOSH B.
EKHAD and DORON ZEILBERGER http://www.math.rutgers.edu/~zeilberg/pj.html and arxiv.org .

July 7, 2014
3

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
