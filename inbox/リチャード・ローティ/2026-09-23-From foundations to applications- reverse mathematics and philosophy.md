---
source: "https://arxiv.org/abs/2609.25183v1"
title: "From foundations to applications: reverse mathematics and philosophy"
author: "Benedict Eastaugh"
year: "2026"
publication: "arXiv preprint / math.HO"
download: "https://arxiv.org/pdf/2609.25183v1"
pdf: "https://arxiv.org/pdf/2609.25183v1"
captured_at: "2026-09-23T04:17:17Z"
updated_at: "2026-09-23T04:17:17Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Rorty Philosophy and the Mirror of Nature"
tags:
  - "現代哲学"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# From foundations to applications: reverse mathematics and philosophy

- 著者: Benedict Eastaugh
- 年: 2026
- 掲載情報: arXiv preprint / math.HO
- 情報源: [arxiv](https://arxiv.org/abs/2609.25183v1)
- ダウンロード: https://arxiv.org/pdf/2609.25183v1
- PDF: https://arxiv.org/pdf/2609.25183v1

## Obsidian Links

- 研究動向: [[リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代哲学 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Reverse mathematics is a branch of mathematical logic dedicated to determining the minimal set existence principles necessary and sufficient to derive ordinary mathematical theorems about concrete structures like the real line. Since the mid-1970s, reverse mathematics has developed a systematic classification of the strength of theorems in areas of mathematics ranging from real and complex analysis to infinitary combinatorics. This essay will place reverse mathematics in its historical and philosophical context, and reveal its relevance to central issues in the philosophy of mathematics, from the foundational programmes of Hilbert and Brouwer to contemporary debates about realism, determinacy, and applicability of mathematics. In doing so, it will discuss the role of computability theory in measuring the strength of set existence principles, as well as related questions about idealisation when these principles are applied in the physical sciences and in philosophy.

## PDF Text

From foundations to applications: reverse mathematics and philosophy
Benedict Eastaugh∗

arXiv:2609.25183v1 [math.HO] 21 Sep 2026

July 7, 2026

Reverse mathematics is a branch of mathematical logic dedicated to determining the minimal set existence principles necessary and sufficient to derive ordinary mathematical theorems about concrete structures like the real line. Since the mid-1970s, reverse mathematics has developed a systematic classification of the strength of theorems in areas of mathematics ranging from real and complex analysis to infinitary combinatorics. This essay will place reverse mathematics in its historical and philosophical context, and reveal its relevance to central issues in the philosophy of mathematics, from the foundational programmes of Hilbert and Brouwer to contemporary debates about realism, determinacy, and applicability of mathematics. In doing so, it will discuss the role of computability theory in measuring the strength of set existence principles, as well as related questions about idealisation when these principles are applied in the physical sciences and in philosophy.

1

From theorems to axioms

Establishing equivalences between axioms and theorems is not a new activity. We can trace it back to the study of Euclid’s parallel postulate and the proposals for substitute principles such as Playfair’s axiom, as already noted in late antiquity by Proclus.1 The development of mathematical logic in the late 19th and early 20th centuries made it possible to understand such equivalences as provable biconditionals of the form
B⊢φ↔ψ
where φ and ψ are sentences in the language of the base theory B and ⊢ denotes the provability relation of a logic such as classical first-order logic. The non-trivial cases of such equivalences are those in which φ (and hence also ψ) are independent of B, i.e. B ̸⊢ φ
and B ̸⊢ ¬φ. Many readers will be familiar with one such equivalence, between the axiom of choice and the well-ordering principle.2
Reverse mathematics is a subfield of mathematical logic which investigates these equivalences within the formal setting of second-order arithmetic. Second-order arithmetic is a quantified two-sorted language with the non-logical symbols of first-order arithmetic plus the membership symbol ∈. This links the two sorts by allowing us to express that a certain number (given by a term of the first sort) is a member of a certain set (given by a term of the second sort). The proof theory and semantics of second-order arithmetic are those
∗

University of Warwick, benedict@eastaugh.net. Thanks to Marianna Antonutti Marfori, Walter Dean, and Alex Paseau for their helpful comments. This work was supported by the Arts & Humanities Research
Council [grant number UKRI3466].
1
Heath (1908), Lewis (1920).
2
Zermelo (1904, 1908).

1

of two-sorted first-order logic, i.e. Henkin semantics. The main objects of study in reverse mathematics are subsystems of second-order arithmetic. These are axiom systems formulated in L2 all of whose axioms are theorems of the theory Z2 of second-order arithmetic, a second-order formalisation of the Dedekind–Peano axioms including the comprehension scheme for all formulas in the second-order language.
Much research in reverse mathematics has been dominated by the ‘Big Five’ subsystems of second-order arithmetic: RCA0 , WKL0 , ACA0 , ATR0 , and Π11 -CA0 . These systems were introduced by Friedman (1975, 1976), who also proved a number of equivalences between theorems of analysis and combinatorics, and the characteristic axioms of these subsystems.3 Although its roots lie deeper, in the arithmetisation of analysis by Dedekind and Cantor, second-order arithmetic as a formal system and its use for formalising analysis first emerge in the second volume of Hilbert and Bernays’s Grundlagen der Mathematik
(1939).4
Second-order arithmetic is thus, even in its historical inception, closely connected not only with the broader issue of foundations of mathematics, but with the Hilbert programme and finitism in particular. It is a highly expressive language and can formalise much of countable and countably representable mathematics, including many classical theorems of real, complex, and functional analysis; theorems from algebra such as the theories of countable commutative rings, abelian groups, and ordered fields; and countably infinite combinatorics such as Ramsey’s theorem and König’s infinity lemma. Reverse mathematics studies the equivalences between these theorems and axioms asserting the existence of different classes of infinite, non-computable sets, such as the axiom schemes of arithmetical comprehension or arithmetical transfinite recursion.
Results in reverse mathematics are thus closely connected to computability theory and the hierarchies of sets it studies, such as the arithmetical and hyperarithmetical hierarchies.
The base theory RCA0 , in which the equivalences of reverse mathematics are typically proved, also has a relationship to computability, as its characteristic axiom asserts the existence of all computable sets. Computability provides a connection to constructive mathematics, as the mathematical theory of computation has been used to make precise the notion of construction. This idea also provides a way of demonstrating that a classical mathematical theorem is not constructively provable, by exhibiting what is known as a recursive counterexample. This is a standard technique in computable and constructive analysis, and because recursive counterexamples are computable objects, their existence can be proved in the base theory RCA0 , forming the basis of many reversals from theorems to axioms.
Recursive counterexamples were developed from the 1940s onwards, while the proof theory of second-order arithmetic and its subsystems developed in parallel. The two areas were brought together decisively to form the new field of reverse mathematics in the mid-1970s. Harvey Friedman introduced the idea of reverse mathematics in a talk at the
International Congress of Mathematicians in 1974 (Friedman 1975). It developed into a substantial research area in mathematical logic in the 1980s and 1990s, particularly through the work of Stephen Simpson and his students and collaborators. In recent years the field has become more diverse in its methodology and its outlook, encompassing work on computability-theoretic reducibility notions and higher-order reverse mathematics.5
Because it has received the most attention from philosophers, we will concentrate on ‘clas3

The historical development of the subsystems is related by Dean and Walsh (2017).
Supplement IV, ‘Formalismen zur deduktiven Entwicklung der Analysis’, pp. 467–512 of (Hilbert and
Bernays 1970).
5
These developments are surveyed in Dzhafarov and Mummert (2022) and Eastaugh (2024).
4

2

sical’ reverse mathematics, meaning provable equivalences between axioms and theorems over the base theory RCA0 .

2

Subsystems of second-order arithmetic

Second-order arithmetic L2 is a two-sorted quantificational language. Variables x0 , x1 , . . .
of the first sort are called number variables and, in the intended interpretation, range over natural numbers. Variables X0 , X1 , . . . of the second sort are called set variables since in their intended interpretation they range over sets of natural numbers. The non-logical vocabulary consists of the symbols familiar from first-order Peano arithmetic: constants
0 and 1, function symbols + and ×, and the less-than relation symbol <, plus the membership relation symbol ∈. The numerical terms consist of the number variables, the constants 0 and 1, and any term of the form t1 + t2 or t1 × t2 where t1 and t2 are numerical terms. The atomic formulas of L2 are all expressions of the form t1 = t2 , t1 < t2 , and t1 ∈ X, where t1 and t2 are numerical terms and X is a set variable. Notice that there are no atomic formulas corresponding to identity for set variables. This is instead defined in terms of co-extensionality,
(=1 )

X = Y ⇔ ∀x(x ∈ X ↔ x ∈ Y ).

The formulas of L2 are obtained by closing the atomic formulas under propositional connectives and universal and existential quantifiers. When a quantifier binds a number variable we call it a number quantifier, while when it binds a set variable we call it a set quantifier. Formulas of second-order arithmetic in prenex normal form are stratified into a hierarchy. A formula is Σ00 if it contains no set quantifiers and all occurrences of number quantifiers are bounded, i.e. have the form ∃x(x < t ∧ φ(x)) or ∀x(x < t → φ(x)). A
prenex formula is Σ0n+1 if it has the form ∃xφ(x) where φ is Π0n , while it is Π0n+1 if it has the form ∀xψ(x) where ψ is Σ0n . Σ0n and Π0n formulas are called arithmetical since they do not involve quantification over sets of numbers, although they may include free set variables.
The semantics of second-order arithmetic are first-order, in the following sense. An
L2 -structure has the form
M = ⟨M, S M , 0M , 1M , +M , ×M , <M ⟩
where M is a non-empty set over which the number variables range and S M ⊆ P(M ) is a non-empty set of subsets of M over which the set variables range.6
The base theory in which equivalences between (formalisations of) mathematical theorems and axioms are proved plays a key role in reverse mathematics. It must be weak enough that the theorems and axioms in question are not already provable, but strong enough to carry out the proof of the equivalence. The standard base theory used in reverse mathematical practice is called RCA0 . Its axioms include, as do all the systems discussed here, the basic arithmetical axioms of PA− , i.e. those of Peano arithmetic minus the firstorder induction scheme.7 RCA0 also includes the Σ01 induction scheme, i.e. all universal generalisations of formulas of the form
(Σ01 -IND)

φ(0) ∧ ∀n(φ(n) → φ(n + 1)) → ∀nφ(n)

6
In other words, general or Henkin semantics rather than the standard semantics for second-order logic, in which second-order variables always range over the entire powerset P(M ) of the first-order domain
(Shapiro 1991, §§4.2–4.3).
7
See Simpson (2009, p. 4).

3

where φ(n) is a Σ01 formula, possibly with additional free number and set variables. Finally,
RCA0 includes the recursive comprehension axiom scheme which gives the system its name, i.e. all universal generalisations of formulas of the form
(∆01 -CA)

∀n(φ(n) ↔ ψ(n)) → ∃X∀n(n ∈ X ↔ φ(n))

where φ(n) is a Σ01 formula and ψ(n) is a Π01 formula, which like the formulas in the induction scheme may have free formula and set variables, except that the set variable X
may not occur freely in them.
The method of arithmetisation, familiar from its roots in the work of Cantor and
Dedekind, makes clear how to formalise statements about natural, rational, and real numbers in the language of second-order arithmetic.8 An arithmetically definable pairing function such as (m, n) = (m + n)2 + m allows pairs of numbers to be coded by single numbers. This simple operation is the foundation of all else. Integers can be represented as pairs of natural numbers, together with operations of integer addition, subtraction, and multiplication, which treat each pair (m, n) as the sum m − n. Rational numbers are pairs of integers (n, d) with d ̸= 0, interpreted as the quotient nd . Real numbers are coded as infinite sequences of rational numbers obeying a version of the Cauchy convergence criterion, where an infinite sequence is a (code for) a function f : N → Q. Functions are represented by sets of (codes of) pairs (m, n) such that for every m there exists exactly one n such that (m, n) ∈ f .
Using these representations, many basic facts about the ring of integers and the rational and real number fields can be proved, for example that Q is an ordered field, or that R
is archimedean. More substantial facts can also be proved in RCA0 , such as a form of completeness of the real numbers called nested interval completeness, or a form of the
Baire category theorem for Rn . The uncountability of the real numbers, in the form of the statement “For every countably infinite sequence of real numbers, there exists a real number which does not occur in the sequence”, is provable in RCA0 by formalising the standard diagonalisation, since the diagonal construction itself is computable relative to the initial sequence of reals.
RCA0 is closely linked to computable mathematics. Its characteristic axiom, the axiom scheme of recursive comprehension, is so called because of Post’s theorem that the sets of natural numbers which are ∆01 definable (meaning those definable by both a Σ01 and a Π01 formula) are exactly the computable (historically called recursive) sets of natural numbers. The axioms of RCA0 are computably true, meaning that they are true when the set quantifiers in the axioms are interpreted as ranging over the computable sets, with the number quantifiers ranging over the standard natural numbers ω = {0, 1, 2, . . . } and the symbols 0, 1, +, ×, and < having their standard meanings. Structures of this sort, with a standard first-order part, are called ω-models. Since they are distinguished from one another entirely by their second-order parts, they are usually referred to only in terms of that second-order part: the ω-model REC whose second-order part is the class REC
of computable sets, the ω-model ARITH whose second-order part is the class ARITH of arithmetically definable sets, and so on.
The system WKL0 is obtained by adding to the axioms of RCA0 an additional axiom known as weak König’s lemma (WKL). This axiom is a restriction to countably infinite trees formed of sequences of 1s and 0s of König’s lemma, the combinatorial principle
8
Cantor (1872), Dedekind (1872). Dauben (1979, p. 37 ff.) gives a condensed presentation of Cantor’s theory of irrational numbers, while Hallett (1984, p. 29 ff.) provides a critical assessment of arithmetisation.
For more precise details on the coding operations described here, see §§II.2–II.4 of Simpson (2009), or
§§3.2–3.3 of Eastaugh (2024).

4

familiar from graph theory and set theory which states that every finitely branching infinite tree has an infinite path through it. To see how to state this in the language of secondorder arithmetic, we start by noting that in RCA0 one can code finite sequences of natural numbers by individual numbers.9 2<N denotes the set of all finite sequences of 0s and
1s. A set T ⊆ 2<N is a tree if for all t ∈ T , if s is an initial subsequence of t, then s ∈ T . A function f : N → 2<N is a path through T . Weak König’s lemma or WKL is the
L2 -statement that every infinite tree T ⊆ 2<N has a path.
WKL0 can prove versions of the Heine–Borel theorem for sequential covers, the Hahn–
Banach theorem for separable spaces, as well as key theorems from mathematical logic such as Gödel’s completeness theorem and the compactness theorem for propositional and first-order logic. All of these theorems are equivalent to weak König’s lemma over
RCA0 , showing that they are not derivable in RCA0 . This is because there exist models of
RCA0 that are not models of WKL0 , most notably the ω-model REC. This follows from a construction due to Kleene (1952) of what has come to be called the Kleene tree, a computable tree T ⊆ 2<N with no computable path.10 WKL0 includes recursive comprehension, so it can prove that Kleene trees exist, but since the paths through such trees are not computable, weak König’s lemma is false in REC.
The arithmetical comprehension axiom scheme consists of the universal closures of all formulas of the form
(ACA)

∃X∀n(n ∈ X ↔ φ(n))

where φ is an arithmetical formula and X is not free in φ. ACA0 is the subsystem of secondorder arithmetic whose axioms are those of RCA0 plus all instances of the arithmetical comprehension axiom scheme. This system is substantially stronger than both RCA0 and
WKL0 . It can prove the sequential completeness and compactness of the real numbers: the Bolzano–Weierstraß theorem, the monotone convergence theorem, and generalisations of these theorems to arbitrary complete separable metric spaces.
ACA0 can also prove König’s infinity lemma, the statement that any finitely branching infinite tree on N has an infinite path. ‘Full’ König’s lemma is strictly stronger than weak
König’s lemma. This is because there are computable, finitely branching infinite trees T
on N such that every path p through T computes the halting problem
K = {e : Φe (e)↓ } , i.e. the set of indexes of Turing machines which halt on every input.11 Since there are ωmodels of WKL0 which do not contain K, WKL0 does not prove König’s infinity lemma.12
On the other hand, arithmetical comprehension is equivalent over RCA0 to König’s lemma.
As a consequence, every ω-model of ACA0 must contain K and indeed all of the finite iterations of the Turing jump operation

X ′ = e : ΦX
e (e)↓ ,
9
This can be done by, for example, using Gödel’s β-function which may be familiar from proofs of the incompleteness theorems.
10
For the standard construction of a Kleene tree, see theorem 9.3.2 of Soare (2016). It relies on the existence of computably inseparable sets, as introduced by Kleene (1950).
11
Φe denotes the function computed by the Turing machine with index e. The notation Φe (n)↓ means that the Turing machine with index e halts when given the input n.
12
The existence of such ω-models of WKL0 follows from the low basis theorem of Jockusch and Soare
(1972). For textbook presentations see Soare (2016, §3.7.2 and chapter 9), Hirschfeldt (2014, p. 59), and
Simpson (2009, §VIII.2).

5

the relativisation of the halting problem to an oracle X. The minimal ω-model of ACA0
has the class of arithmetically definable sets ARITH as its second-order part; by Post’s theorem these are exactly the sets which are computable relative to some finite number of iterations of the Turing jump to the empty set.
We will not discuss the fourth and fifth systems of the Big Five in great detail, since many of the philosophical issues connected to reverse mathematics already appear in relation to the first three systems, but for the sake of completeness and the few times we will need to discuss them, we briefly introduce them now. The fourth system of the Big
Five is ATR0 . Its characteristic axiom is the scheme of arithmetical transfinite recursion.
This has a slightly technical definition, but roughly speaking it is the principle that if a set
X codes a well-ordering, then one can iterate any arithmetical operation along that wellordering.13 ATR0 is thus strictly stronger than ACA0 , since not only do its axioms extend arithmetical definability into the transfinite, it can prove the existence of the minimal
ω-model ARITH of ACA0 .
The theorems provable in ATR0 extend beyond real and complex analysis and into the lower reaches of descriptive set theory. They include Lusin’s separation theorem, that any two disjoint analytic sets can be separated by a Borel set; the perfect set theorem, that every uncountable closed set has a perfect subset; and the fact that any two countable wellorderings are comparable. Π11 -CA0 is the final member of the Big Five. Its characteristic axiom is the Π11 comprehension axiom. This asserts the existence of sets definable by Π11
formulas, those of the form ∀Xφ(X, n) where X is a set variable and φ is an arithmetical formula. Π11 -CA0 is strictly stronger than ATR0 . The theorems it can prove include more results from descriptive set theory, such as the Cantor–Bendixson theorem.

3

Constructivity and computability

Much of the early philosophical discussion of reverse mathematics revolved around its relevance for the foundations of mathematics, and in particular its connections to the foundational programmes of intuitionism, predicativism, and finitism developed in the first decades of the 20th century. Despite the substantial intellectual differences between these programmes, they nevertheless share one important characteristic, namely an adherence to a spirit of what can broadly be called constructivism: the view that the only mathematical objects which exist are those which can be constructed. There is at least a surface-level similarity with the programme of reverse mathematics, since the set existence principles that characterise different subsystems of second-order arithmetic are computability-theoretic or definability-theoretic in nature. Perhaps unsurprisingly, there are in fact deeper connections between foundational perspectives and particular systems.
However, these relationships are not straightforward: one cannot simply read off the mathematical resources of a given foundational standpoint by looking at the theorems provable in a potentially associated subsystem of second-order arithmetic.
The first connection is that between computable functions and the various forms of constructivism, in the more restricted sense particular to Brouwer’s intuitionism, the Russian school of constructive mathematics following Markov, and Bishop’s constructive mathematics. From the 1940s onwards, researchers in recursive function theory, as computability theory was then called, attempted to use results concerning computable functions in order to better understand the limits of constructivism as a foundation.
One difficulty with the identification of constructive with computable is that constructive mathematicians do not accept the law of the excluded middle, and so not everything
13

For a precise definition, see Simpson (2009, §V.2)

6

that is computably true (that is, true in the ω-model REC) is constructively provable.
A standard example of this is the intermediate value theorem, which is true in REC
and provable in RCA0 ,14 but which is nonetheless non-constructive, since its proof uses a non-constructive case distinction in an essential way. Despite the seeming naturalness of identifying constructivity with computability, examples like this one show that computable truth is not sufficient for constructive provability.
Depending on the version of constructivism in question, it may not be necessary either. Brouwer’s conception of construction included not just lawlike potentially infinite sequences or functions given by rules, but also absolutely free or lawless choice sequences in which the intuitionistic mathematician can freely choose a value at every step.15 When interpreted in a classical model, such sequences can be non-computable, as seen in Kleene’s
(1952) work on Brouwer’s fan theorem. The contrapositive of the fan theorem is weak
König’s lemma, and thus when the fan theorem is included in a classical system, it implies the existence of non-computable sets. Other forms of constructivism aim to be compatible with constructive Church’s thesis (CT), the claim that every function f : Nk → N is computable. As well as the recursive constructive mathematics of Markov, this includes
Bishop’s constructive mathematics, despite Bishop’s opposition to the identification of constructive mathematics with computable mathematics.16
Cardinality considerations show us that most real numbers are non-computable, since there are uncountably many reals and only countably many computable reals.17 In other words, there are many ‘gaps’ in the computable real numbers RREC . One way in which this manifests itself is in the fact that many theorems of classical analysis are false when the range of the quantifiers in the statements of these theorems are restricted to the computable reals. Consider a theorem θ of the form
∀X(φ(X) → ∃Y ψ(X, Y )).
A recursive counterexample to θ is a computable set which satisfies its antecedent but not its consequent. In other words, it is a computable set X ⊆ ω such that φ(X) holds but there is no computable Y ⊆ ω such that ψ(X, Y ) holds. The existence of a recursive counterexample to a theorem shows that it is computably false.
For constructivists who take CT to be consistent, non-computability implies nonconstructivity. To show that a statement is constructively unprovable, it therefore suffices to show that there exists a recursive counterexample to it. An early example was
Specker’s (1949) construction of a computable, bounded, monotone sequence of rational numbers whose limit is non-computable. The existence of Specker sequences, as they are now known, shows that the monotone convergence theorem—together with other theorems which express the sequential completeness of the real line—is computably false, and thus constructively unprovable. In reverse mathematics, recursive counterexamples also play another role: reversals are often mediated by the construction of a recursive counterexample in the base theory. For example, one can construct a Specker sequence using recursive comprehension, apply the monotone convergence theorem to obtain its limit, which must compute the set K of solutions to the halting problem, thereby implying Σ01
comprehension and hence arithmetical comprehension.18
14

Simpson (2009, theorem II.6.6).
For an introduction to free choice sequences, see chapter 3 of (van Atten 2004).
16
Representative remarks appear in (Bishop 1967, pp. 6, 74) and (Bishop 1975, p. 514).
17
A computable real is, roughly speaking, a real number whose value can be calculated by a computer program to any desired degree of precision. For details see Pour-El and Richards (1989, p. 14).
18
For details of these arguments, see lemma III.1.3 and theorem III.2.2 in Simpson (2009, pp. 105–107).
15

7

Another example is a version of the Heine–Borel theorem (HB) stating that every countable open cover of the closed unit interval has a finite subcover. A recursive counterexample to HB is known as a singular cover, and is constructed by forming a computable sequence of open intervals Cn = (ln , rn ) which covers every computable real number, but which has classical measure < 1. There will then be reals x ∈ [0, 1] not covered by any Cn , and which therefore must be non-computable (Kreisel and Lacombe 1957, Le Roux and
Ziegler 2008). The construction of this recursive counterexample can be formalised within
RCA0 , and used to prove that HB implies weak König’s lemma (Simpson 2009, §IV.1).

4

Finitism and conservativity

As far as purely number-theoretic statements are concerned, WKL0 and RCA0 prove exactly the same ones: if φ is an arithmetical sentence such that WKL0 ⊢ φ, then RCA0 ⊢ φ
(Simpson 2009, §IX.2). Moreover, by a theorem of Parsons, WKL0 and RCA0 are both conservative for Π02 sentences over the theory PRA of primitive recursive arithmetic (Parsons 1970, Friedman 1976). These facts led Simpson (1985, 1988) to suggest that the mathematics provable in WKL0 , and other systems which are conservative over PRA for
Π01 sentences, is finitistically reducible, and that this constitutes a partial realisation of
Hilbert’s programme.
To understand this claim we first need to briefly reexamine Hilbert’s programme.
Hilbert took finitary mathematics—the part of mathematics with finitary content—to include number-theoretic equations and inequalities, but also universal generalisations of such formulas, i.e. Π01 sentences. Hilbert aimed to justify infinitary mathematics in a finitary way. Initially the intent of his programme was to provide finitary proofs of the consistency of infinitary theories. In the second half of the 1920s its aim shifted somewhat, towards proving the conservativity of infinitary theories over finitary theories. This new version of the programme can be understood in terms of an analogy with physics suggested by Weyl (1925). The finitary or ‘real’ part of a mathematical theory is accorded a status analogous to that of the observation sentences of a physical theory, the part of the theory that describe the outcomes of empirical observations. If the infinitary or ‘ideal’ part of the theory, analogous to the theoretical sentences of a physical theory, imply a finitary statement, then that finitary statement must already be provable using finitary means alone, just as observation sentences implied by a physical theory must be (in principle)
verifiable by experiment. In other words, a mathematical theory must be conservative over its finitary part.
In an influential article on Hilbertian finitism, Tait (1981) argues for two theses. Firstly, the finitist functions f : Nk → N are precisely the primitive recursive functions. Secondly, the finistically provable part of mathematics consists of those Π01 sentences which are provable in the formal system PRA of primitive recursive arithmetic.19 Building on Tait’s theses, Simpson claims that the Π01 conservativity of RCA0 and WKL0 over PRA shows that Hilbert’s programme can succeed, at least in part, because the infinitary mathematics provable in these systems is so substantial. This argument appears to be bolstered by the fact that the conservativity of these systems over PRA is finitistically provable, in the following sense. Sieg (1985) constructed a primitive recursive function g : N → N which,
19

Tait’s theses have gained broad but not universal acceptance. Kreisel (1958, 1970) considers them too restrictive, and argues that the finistically provable sentences include those provable in the stronger system
PA of first-order Peano arithmetic. Ganea (2010) instead argues that Tait’s thesis is too permissive, and that finitistic provability includes no more than sentences provable from the equational theory of Kalmár elementary functions.

8

given as input a proof p of a Π02 sentence φ in WKL0 , produces as output a proof g(p)
of φ in PRA. Because the conservativity theorem is, per Tait’s thesis, finitistically provable, Simpson argues that Hilbert’s programme is successfully realised for the infinitary mathematics provable in WKL0 .
As Burgess (2010) points out, a closer reading of the situation suggests a difficulty with
Simpson’s argument. Suppose that φ is a Π01 sentence, hence finitistically meaningful, with a proof p in WKL0 . Then g(p) is a finitist proof per Tait’s thesis: all the finitist has to do is check that all of the axioms used in the proof are finitistically acceptable, which they are, since they are axioms of PRA. The problem is that this is not sufficient for the finitist to have confidence in the infinitary theory WKL0 as a whole. For that, they would need not only to know (in a finitistically provable way, i.e. in PRA) that WKL0 is Π01 conservative over PRA, but that any Π01 consequence of WKL0 is finitistically true. In other words, they would need to know the Π01 reflection principle for WKL0 , the scheme
ProvWKL0 (⌜φ⌝) → φ
for all sentences φ ∈ Π01 . The conservativity theorem only guarantees, for the finitist, that
ProvWKL0 (⌜φ⌝) → ProvPRA (⌜φ⌝), since to establish φ as finitistically provable we need to invoke Tait’s thesis, the claim that
ProvPRA (⌜φ⌝) → φ
for any Π01 sentence φ. This scheme is not provable in PRA, since an instance of the scheme is
ProvPRA (⌜0 = 1⌝) → 0 = 1, which is equivalent to the consistency statement for PRA, and hence unprovable in PRA
by Gödel’s second incompleteness theorem.20
The problem is that Tait’s thesis is a claim made from outside the finitary standpoint, not from within it (Tait 1981, p. 527). Finitists can know of individual infinitary proofs that they establish finitistically true statements, via a finitarily provable conservativity theorem, since such conservativity theorems deliver a finite proof object that can be verified by the finitist to only use finitistically acceptable axioms. But they cannot know, at least for systems such as RCA0 or WKL0 , that these systems only prove finitistically true Π01
sentences, since this amounts to knowing the consistency of PRA, something which is unavailable to them.

5

Predicativity and definability

In the wake of the class-theoretic paradoxes, Russell realised that propositional functions do not always define a corresponding class. He called predicative those propositional functions φ(x) such that the corresponding class {x : φ(x)} exists. Propositional functions like x ̸∈ x which do not, on pain of contradiction, define a class, Russell called impredicative
(Russell 1907). This distinction gave way to one framed in terms of the theory of types, and of the ‘vicious circle principle’: “no totality can contain members defined in terms of itself” (Russell 1908, p. 237). On this account, a propositional function φ(x) has type n+1
just in case its arguments and values have type n, and all variables bound by quantifiers in
20

See Giaquinto (1983) and Dean (2015) for related discussion.

9

φ(x) have type n or lower. A predicative function is one which appears in this hierarchy, i.e. has type n + 1 for some n, where 0 is the type of individuals.21
Poincaré (1908) objected to the details of Russell’s attempts to spell out a solution, but agreed that the source of the paradoxes lay in definitions which contained vicious circles, in the sense that if one defines two concepts C and C ′ , their definitions are impredicative if the concept C occurs in the definition of C ′ , and conversely. However, Poincaré also calls definitions predicative if they are “not changed by the introduction of new elements”
(Poincaré 1910). Weyl (1918) drew on the ideas of Russell and Poincaré in developing his own view, one in which vicious circles were to be avoided by restricting comprehension to predicates which quantified only over natural numbers. Despite the fact that Weyl swiftly moved on from predicativism, adopting Brouwer’s intuitionism, his approach exerted a significant influence on later developments in predicative analysis and on reverse mathematics. Dedekind had already shown that fundamental theorems of 19th century analysis were equivalent to one another and to the least upper bound principle (Dedekind 1872,
§VII). Weyl now showed that versions of these theorems could be proved in a predicative framework.
After a fallow period, predicativity underwent a revival in the mid-1950s, becoming connected to the emerging subject of definability theory. Grzegorczyk (1955), Kondô
(1958), and Mostowski (1959) all suggested that predicative analysis in the sense of Weyl could be identified with elementary analysis, in which every set of natural numbers is given by an elementary definition, i.e. one in the first-order language of arithmetic. This amounts to studying analysis in the ω-model ARITH = {X ⊆ ω : X is arithmetically definable}.
Grzegorczyk and Mostowski both provided axiomatisations of elementary analysis, which can be seen as ancestors of the theory ACA0 . The axioms of this system can all be justified on predicative grounds, since it postulates the existence only of sets which are definable in terms of objects that are already ‘given’, namely the natural numbers. Theorems provable in ACA0 are thus predicatively provable.
The fundamental idea of predicative definability, in which sets can be defined in terms of previously given or already existing objects, is prima facie more general than arithmetical definability, since it seems predicatively acceptable to then iterate the procedure, taking the arithmetical sets as given and defining new sets on that basis. One way of doing this is by iterating arithmetical definability into the transfinite, along well-orderings which are predicatively acceptable, and Wang (1954) and Lorenzen (1955) made early forays into this approach. This idea underlaid the proposal of Kreisel (1960) to identify the predicatively definable sets of natural numbers with those occurring in the hyperarithmetical hierarchy.22 On this basis, Kreisel proposed several axiom systems true in the ω-model
HYP = {X ⊆ ω : X is hyperarithmetical}, including ∆11 -CA0 and Σ11 -AC0 . These systems of hyperarithmetical analysis did not prove substantially more mathematically fruitful than elementary analysis, with Kreisel writing that “in the portions of analysis developed by working mathematicians, a theorem is either derivable by means of the arithmetic comprehension axiom or else not predicative at all” (Kreisel 1962, p. 316). Hyperarithmetical analysis has only recently started to yield theorems which are not arithmetically true, such as statements concerning indecomposable linear orders (Montalbán 2006), and versions of
Halin’s infinite ray theorem (Barnes, Goh, and Shore 2022).
The two remaining Big Five systems, ATR0 and Π11 -CA0 , both have a level of stability and mathematical fruitfulness not achieved by theories of hyperarithmetical analysis.
21

The history of Russell’s thought is traced in Irvine and Deutsch (2020).
The hyperarithmetical sets are those which are computable relative to the αth iteration of the Turing jump operator starting from the empty set, where α is a computable ordinal (Sacks 1990, Part A).
22

10

However, on the widely accepted thesis that the ordinal Γ0 marks the outer limit of predicativity, both of these theories are impredicative. This is most obvious in the case of Π11 -CA0 , since its proof-theoretic ordinal is substantially larger than Γ0 , and its characteristic axiom is prima facie impredicative.23 Kreisel (1959) constructed a recursive counterexample to the Cantor–Bendixson theorem, which states that every closed set is the union of a perfect set and a countable set. Kreisel’s construction is of a computable (code for a) closed set
C such that, by the Cantor–Bendixson theorem, C = P ∪ S with P a perfect set and S a countable set, but such that both P and S are coanalytic (Π11 ) but not analytic (Σ11 ). It follows that neither are hyperarithmetical, and hence not predicative by Kreisel’s analysis of predicative definability. Kreisel’s argument can be internalised within ACA0 to show that the Cantor–Bendixson theorem implies Π11 comprehension (Simpson 2009, §VI.1).
The abundance of reversals to Π11 -CA0 and ATR0 suggests that these systems form natural accumulation points in the hierarchy of subsystems of second-order arithmetic, even in the absence of justifications for these systems in terms of historical foundational programmes.

6

The intrinsic significance of reversals

Philosophy of mathematics is broader than foundations of mathematics, especially when one considers this in the narrow sense of the early 20th century foundational programmes.
One might reasonably wonder what reverse mathematics can contribute to philosophy of mathematics more broadly, beyond measuring the limits of the mathematics that can be developed from within a given foundational perspective, especially from a more broadly realist point of view on which the axioms of subsystems of second-order arithmetic are viewed as true, but inadequate to formalise all of mathematics. We can start to understand what is at stake here by asking: What do we learn from reversals?
The received view, as presented most influentially by Simpson (2009, pp. 1–2), is that reversals show which set existence principles are needed to prove the theorem in question.
The use of this term emphasises issues of mathematical ontology, and might lead one to infer that the practice of reverse mathematics involves some underlying scepticism about the reliability of set-theoretic methods, or the ontological picture of the cumulative hierarchy of sets which is conventionally used to justify them. This is far from being the case: most work in reverse mathematics is neither motivated by a particular foundational programme, nor constrained by a desire to do away with tools and concepts that go beyond those formalisable in second-order arithmetic. It does, however, leave open how best to answer the question of what we learn from reversals, as a general issue in the philosophy of mathematics.
Hirschfeldt (2014) has argued that reversals reveal the combinatorial core of a theorem, the combinatorial principle that is essential to or underpins any proof of that theorem, regardless of whether it makes an explicit appearance in that proof. For example, the combinatorial core of Lindenbaum’s lemma is weak König’s lemma (Hirschfeldt 2014, p. 9). The use of WKL is clear in the standard proof of Lindenbaum’s lemma. Fix an
23

This analysis of the limits of predicativity is surveyed by Feferman (2005). The proof-theoretic ordinal of a system S is usually understood as the least ordinal α such that S cannot prove that α is wellordered, modulo a reasonable computable presentation (an ordinal notation system) of α and ordinals
β < α. Systems such as ATR0 which are not predicative themselves, but have proof-theoretic ordinal Γ0 , are sometimes called predicatively reducible. Simpson (1985) argues that ATR0 and other predicatively reducible systems can be used instrumentally by the predicativist. However, this analysis seems vulnerable to an argument parallel to that deployed against finististic reducibility, as suggested by Burgess (2010, p. 140) and discussed by (Eastaugh 2024, §5.5). The Γ0 analysis of the limits of predicative provability is not universally accepted, with Weaver (2009) a notable dissenter.

11

enumeration ⟨φn : n ∈ N⟩ of the sentences of a countable language L, and for any φi in the enumeration let φ1i ≡ φ and φ0i ≡ ¬φ. Given an L-theory S and a finite sequence
σ ∈ 2<N , set n
o
σ(i)
Sσ = S ∪ φi : i < |σ| .
We form a binary tree TS by putting σ ∈ TS if there is no proof of a contradiction of length ≤ |σ| from Sσ . This is a computable process, so TS can be proved to exist in RCA0 , relative to S. TS is infinite if and only if S is consistent, so by WKL there is an infinite path P : N → 2 through TS . The set S ∗ = {φi : P (i) = 1} is a maximal consistent set extending S. Weak König’s lemma is essential here, since there are computable consistent sets T (such as PA) which cannot be extended to a computable maximal consistent set
S ∗ . This points to a crucial aspect of reverse mathematics, namely the fact that the characteristic axioms of the Big Five have a computability-theoretic character.
This is most easily grasped by examining the ω-models of the Big Five. We have already discussed some examples: the computable sets REC form an ω-model of RCA0 , the arithmetical sets ARITH form an ω-model of ACA0 , and so on. These examples are quite representative, and reveal important structural properties common between all ωmodels of those systems. For example, every ω-model X ⊆ P(ω) of RCA0 is downwards closed under Turing reducibility, meaning that if X ∈ X and Y is computable by a Turing machine with an oracle for X, then Y ∈ X . Similarly every ω-model X ⊆ P(ω) of ACA0 is closed under arithmetical reducibility: if X ∈ X and Y is computable from X via a finite number of Turing jumps, then Y ∈ X .
Eastaugh (2019), following other authors such as Kreisel (1968), Shore (2010), and
Chong, Feng, Slaman, and Woodin (2014), foregrounds the computability-theoretic nature of the characteristic axioms of the Big Five, and argues that they express closure conditions on the second-order part of the model. This suggests the following unified way of looking at what we learn from reversals. Stronger axioms express stronger closure conditions, which guarantee that more sets exist (the existential aspect), that stronger combinatorial properties hold (the combinatorial aspect), and that models of the stronger axioms are closed under stronger computability-theoretic reducibility relations (the computational aspect). They are thus intertwined facets of the same basic phenomenon.
Reversals thus play an explanatory role, revealing the logico-combinatorial-computational structure that supports (and is supported by) a particular theorem—or, considered slightly differently, as revealing the logico-combinatorial-computational content of that theorem.24
For example, a theory adequate to the analysis of Weierstraß and Dedekind must include the Bolzano–Weierstraß theorem and the Cauchy convergence theorem. The fact that we can reason backwards from those theorems to the axiom scheme of arithmetical comprehension, as well as forwards from ACA0 to those theorems, is revealing of both an important logical and combinatorial structure in the theory of analysis, and the computabilitytheoretic content of those theorems. This latter fact was not evident from the practice of real analysis alone, and needed to be shown by work in mathematical logic. It is equally revealing that although the sequential Heine–Borel theorem is derivable in this system, one cannot reverse this implication, and the models of WKL0 (the weaker system necessary
24

Although this explanatory role can be partially understood in semantic terms, it seems implausible that reversals can give a satisfying theory of mathematical content broadly understood. Such a view might identify the content of a theorem φ with its set of models, or perhaps with the equivalence class of statements which are provably equivalent to φ in RCA0 . But as we have seen, these equivalence classes can be very large, and it seems implausible to say that for the practising analyst the Bolzano–Weierstraß theorem has the same content as König’s lemma. See Arana and Mancosu (2012) for some related considerations in the context of planar and solid geometry.

12

and sufficient to prove the sequential Heine–Borel theorem) have quite different properties to those of ACA0 . Compactness principles like Heine–Borel thus belong to a quite different species than completeness principles like the monotone convergence or Bolzano–Weierstraß
theorems. This is so even though the adoption of completeness principles was historically key to proving the Heine–Borel theorem.25

7

Reverse mathematics and nominalism

The indispensability argument seeks to justify realism about mathematics on the basis of mathematics’ role in our best scientific theories, where those theories are construed as true or at least truth-apt claims about the nature of the world. Mathematical entities, on Quine’s view, are epistemically on a par with other theoretical entities held to exist by scientific theories, since they are indispensable to our best science. A natural question which emerges in this context is how much mathematics is indispensable to science. Quine took even the irrational numbers to stand in need of justification, in terms of the ways in which admitting them into our theories simplifies computations and generalisations—in other words, theoretical virtues that go beyond the mere prediction of observed values in possible experiments. Set theory beyond the real numbers is admissible only because it completes and systematises our best theories of applied mathematics (Quine 1998, p. 400).
Putnam, another exponent of the indispensability argument, held that physics could survive—if not thrive—using only the mathematical resources of predicative set theory.
On the other hand, the mathematical usefulness of impredicative set theory for the mathematics ultimately applied in physics constitutes an argument for its truth, albeit not as strong as that for the indispensable predicative theory (Putnam 1971, p. 55–56). The first part of Putnam’s view has been embraced by those with nominalist or constructivist sympathies, since if the indispensability argument does not justify all or even much of classical set theory, then the door might be open to reconstructing the scientifically applicable part of mathematics in a nominalist or broadly constructive way. It offers an obvious role for reverse mathematics, in characterising the mathematical axioms necessary for deriving theorems which are indispensably applied in physics, or biology, or even in social sciences such as economics.
Feferman (1988, 1992) has defended a view along these lines, holding that the mathematics that is indispensable to our best science is predicative. The body of results proved in reverse mathematics endows Feferman’s claim with a certain plausibility, because a substantial fragment of 19th and 20th century analysis can be recovered in predicative systems such as ACA0 . This includes many theorems which are widely used in scientific applications. A counterexample to this claim would have to have the form of a mathematical theorem which is indispensable to some scientific application but unprovable in ACA0
(or some conservative extension thereof). One suggestion has been that non-separable spaces are used in nontrivial ways in applied mathematics, especially physics. Feferman’s response (Feferman 1998, p. 281) is that although such spaces are widely used, it is not clear that the non-separability of the spaces really makes a difference to the applications in question. This remains an area where careful, detailed work in logic and the foundations of physics remains necessary.
More strictly nominalist authors including Hellman (1989, 1999) and Bueno (2001)
have sought to use reverse mathematical results to defuse the indispensability argument.
Their general strategy is to use reversals to determine the scope of scientifically applicable mathematics in terms of subsystems of second-order arithmetic, and then argue that those
25

See Andre et al. (2013) and the extensive discussion by Hallett 1979, pp. 20–25.

13

subsystems can be reinterpreted in a nominalistically acceptable fashion. Hellman (1999), for example, argues that one can develop a nominalistic account of the natural numbers, and thus an indispensability argument is not needed in order to justify predicative theories such as ACA0 . This is because for such theories, membership ascriptions of the form n ∈ X can be reinterpreted as ascriptions of the form φ(n) where φ is an arithmetical formula defining the set X. These arithmetical ascriptions are taken to be ontologically non-committing due to the claimed availability of a nominalisation of the natural numbers. Where quantification over sets is needed, it can be replaced by quantification over arithmetical formulas. Note however that the acceptability of the nominalisation strategy rests on the complexity of the sets which the subsystem of second-order arithmetic proves to e

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
