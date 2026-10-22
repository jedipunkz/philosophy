---
source: "https://archive.org/details/arxiv-1107.0434"
title: "Abstraction Super-structuring Normal Forms: Towards a Theory of Structural Induction"
author: "Adrian Silvescu, Vasant Honavar"
year: "2011"
captured_at: "2026-10-22T04:22:49Z"
updated_at: "2026-10-22T04:22:49Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "デイヴィッド・ヒューム"
query: "David Hume"
plain_text_url: "https://archive.org/download/arxiv-1107.0434/1107.0434_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "経験論"
  - "懐疑主義"
status: raw
---

# Abstraction Super-structuring Normal Forms: Towards a Theory of Structural Induction

- 著者: Adrian Silvescu, Vasant Honavar
- 初版: 2011
- 情報源: [archive](https://archive.org/details/arxiv-1107.0434)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[デイヴィッド・ヒューム]]
- 研究動向: [[デイヴィッド・ヒューム-現代研究動向]]

## Full Text

Abstraction Super-structuring Normal Forms:
Towards a Tlieory of Structural Induction

Adrian Silvescu and Vasant Honavar

^D Technical report

1— I Department of Computer Science, Iowa State University, Ames, lA, USA

1—5

m

cd

Abstract

Induction is the process by which we obtain predictive laws or theories

^ or models of the world. We consider the structural aspect of induction. We

^~j answer the question as to whether we can find a finite and minmalistic set

■^[^ of operations on structural elements in terms of which any theory can be

r/^ expressed. We identify abstraction (grouping similar entities) and super-

O structuring (combining topologically e.g., spatio-temporally close entities)

as the essential structural operations in the induction process. We show

. that only two more structural operations, namely, reverse abstraction and

►^ reverse super-structuring (the duals of abstraction and super-structuring

^^ respectively) suffice in order to exploit the full power of Turing-equivalent

^^ generative grammars in induction. We explore the implications of this

^^ theorem with respect to the nature of hidden variables, radical positivism

C*^ and the 2-century old claim of David Hume about the principles of con-

r ^ nexion among ideas.

o

,— I 1 Introduction

. !_^ The logic of induction, the process by which we obtain predictive laws, theo-

S^ ries, or models of the world, has been a long standing concern of philosophy,

^ science, statistics and artifical intelligence. Theories typically have two aspects:

structural or qualitative (corresponding to concepts or variables and their rela-
tionships, or, in philosophical parlance, ontology) and numeric or quantitative
(corresponding to parameters e.g., probabilities). Once the qualitative aspect of
a certain law is fixed, the quantitative aspect becomes the subject of experimen-
tal science and statistics. Induction is the process of inferring predictive laws,
theories, or models of the world from a stream of observations. In general, the
observations may be passive, or may be the outcomes of interventions by the
learning agent. Here, we limit ourselves to induction from passive observation
alone.

Under the computationalistic assumption (i.e., the Church- Turing thesis,
which asserts that any theory can be described by a Turing Machine fTSl), one

way to solve the induction problem is to enumerate all the Turing machines
(dovetailing in order to cope with the countably infinite number of them) and
pick one that strikes a good balance between the predictability (of the finite
experience stream) and size (complexity) jTB], |T7], [TS] or within a Bayesian
setting, using a weighted vote among the predictions of the various models |2
(See|2] and references therein). In the general setting, a priori the number of
types of possible structural laws that can be postulated is infinite. This makes
it difficult to design general purpose induction strategy. We ask whether a finite
and minimalistic set of fundamental structural operations suffice to construct
any set of laws. If such a set will render induction more tractable because at any
step the learner will have to pick from a small finite set of possible operations
as opposed to an infinite one.

Because Turing machines are rather opaque from a structural standpoint, we
use the alternative, yet equivalent, mechanism of generative grammar^ This al-
lows us to work with theories that can be built recursively by applying structural
operations drawn from a finite set. The intuition behind this approach is that
induction involves incrementally constructing complex structures using simpler
structures (e.g., using super-structuring, also called chunking), and simplifying
complex structures when possible (e.g., using abstraction). Such a compositional
approach to induction offers the advantage of increased transparency over the
enumerate-and-select approach pioneered by Solomonoff [TB], [TT]. It also offers
the possibility of reusing intermediate structures as opposed to starting afresh
with a new Turing machine at each iteration, thereby replacing enumeration by
a process akin to dynamic programming or its heuristic variants such as the A*
algorithm.

We seek laws or patterns that explain a stream of observations through suc-
cessive applications of operations drawn from a small finite set. The induced
patterns are not necessarily described solely in terms of the input observations,
but may also use (a finite number of) additional internal or hidden (i.e., not
directly observable) entities. The role of these internal variables is to simplify
explanation. The introduction of internal variables to aid the explanation pro-
cess is not without perils [1 2 F] One way to preclude the introduction of internal
variables is to apply the following demarcation criterion: If the agent cannot
distinguish possible streams of observations based on the values of an inter-
nal variable, then the variable is non-sensical (i.e., independent of the data or
"senses") F] The direct connection requirement restricts the no-nonsense theo-
ries to those formed out empirical laws [T] (i.e, laws that relate only measurable
quantities). However several scientists, including Albert Einstein, while being
sympathetic to the positivists ideas, have successfully used in their theories,
hidden variables that have at best indirect connection to observables. This has

^See llOl for a similarly motivated attempt using Lambda calculus

■^Consider for example, a hidden variable which stands for the truth value of the sentence:
"In heaven, if it rains, do the angels get wet or not?"

^This is a radical interpretation of an idea that shows up in the history of Philosophy from
Positivism through the empiricists and scholastics down to Aristotle's "Nihil est in intellectu
quod non prius fuerit in sensu".

led to a series of revisions of the positivists doctrine culminating in Carnap's
attempt to accommodate hidden variables in scientific explanations [3 . The ob-
servables and the internal variables in terms of which the explanation is offered
can be seen as the ontologjj^]- i.e., the set of concepts and their interrelation-
ships found useful by the agent in theorizing about its experience. In this setting,
structural induction is tantamount to ontology construction.

The rest of the paper is organized as follows: Section 2 introduces Ab-
straction Super-structuring Normal Forms that correspond to a general class
of Turing-equivalent generative grammars that can be used to express theories
about the world; and shows that: abstraction (grouping similar entities) and
super-structuring (combining topologically e.g., spatio-temporally close entities)
as the essential structural operations in the induction process; Only two more
structural operations, namely, reverse abstraction and reverse super- structuring
(the duals of abstraction and super-structuring respectively, suffice in order to
exploit the full power of Turing-equivalent generative grammars in induction.
Section 3 interprets the theoretical results in a larger context the nature of hid-
den variables, radical positivism and the 2-century old claim of David Hume
about the principles of connexion among ideas. Section 4 concludes with a
summary.

2 Abstraction Super- Structuring Normal Forms

We start by recapitulating the definitions and notations for generative grammars
and the theorem that claims the equivalence between Generative Grammars
and Turing Machines. We then draw the connections between the process of
induction and the formalism of generative grammars and motivate the quest
for a minimalistic set of fundamental structural operations. We then get to the
main results of the paper: a series of characterization theorems of two important
classes of Generative Grammars: Context-Free and General Grammars, in terms
of a small set of fundamental structural operations.

2.1 Generative Grammars and Turing Machines

Definitions (Grammar) A (generative) grammar is a quadruple {N,T, S,R)
where N and T are disjoint finite sets called NonTerminals and Terminals,
respectively, S* is a distinguished element from N called the start symbol and
i? is a set of rewrite rules (a.k.a. production rules) of the form {I — >■ r) where
I € {NU T)*N{N U T)* and r e (NUT)*. Additionally, we call I the left
hand side (Ihs) and r the right hand side (rhs) of the rule (/ — s- r). The language
generated by a grammar is defined by L{G) — {w E T*\S ^^ w} where — )• stands
for the refiexive transitive closure of the rules from R. Furthermore — >■ stands
for the transitive (but not reflexive) closure of the rules from R. We say that
two grammars G,G' are equivalent if L{G) = L{G'). The steps contained in a

*The ontology in this case is not universal as it is often the case in philosophy; it is just a
set of concepts and interrelations among them that afford the expression of theories.

set of transitions a — ^ /3 is called a derivation. If we want to distinguish between
derivations in different grammars we will write a -^g P or mention it explicitly.
We denote by e the empty string in the language. We will sometimes use the
shorthand notation / -^ ?'i|''2|---kn to stand for the set of rules {I -^ r^lj^i „.
See e.g., \Y6\ for more details and examples.

Definition (Grammar Types) Let G — {N, T, S, R) be a grammar. Then

1. G is a regular grammar (REG) if all the rules {I ^ r) £ R have the
property that I G N and r G (T* U T*N).

2. G is context-free grammar (CFG) if all the rules (/ ^- r) G i? have
the property that I £ N.

3. G is context-sensitive grammar (CSG) if all the rules {I ^ r) E R
have the property that they are of the form aA/3 — >■ ajP where A G N
and a, (3, 'J G {N U T)* and j ^ e. Furthermore if e is an element of
the language one rule of the form S" — >■ e is allowed and furthermore the
restriction that S does not appear in the right hand side of any rule is
imposed. We will call such a sentence an e — Amendment.

4. G is general grammar (GG) if all the rules (/ — > r) G i? have no
additional restrictions.

Theorem 1. The set of General Grammars are equivalent in power with the
set of Turing Machines. That is, for every Turing Machine T there exists a
General Grammar G such that L{G) = L[T) and vice versa.

Proof. This theorem is a well known result. See for example |13| for a prooPl

D

2.2 Structural Induction, Generative Grammars and Mo-
tivation

Before proceeding with the main results of the paper we examine the connec-
tions between the setting of generative grammars and the problem of structural
induction. The terminals in the grammar formalism denote the set of observ-
ables in our induction problem. The NonTerminals stand for internal variables
in terms of which the observations (terminals) are explained. The "explanation"
is given by a derivation of the stream of observations from the initial symbol
S ^ w. The NonTerminals that appear in the derivation are the internal vari-
ables in terms of which the surface structure given by the stream of observations
w is explained. Given this correspondence, structural induction aims to find an
appropriate set of NonTerminals N and a set of rewrite rules R that will al-
low us to derive (explain) the input stream of observations w from the initial
symbol S. The process of Structural Induction may invent a new rewrite rule

^Similar results of equivalence exist for transductive versions of Turing machines and gram-
mars as opposed to the recognition versions given here (See e.g., [2] and references therein).
Without loss of generality, we will assume the recognition as opposed to the transductive
setting.

Structural Induction

Generative Grammar

Observables

Terminals T

Internal Variables

NonTerminals N

Law / Theory

production rule(s) I -^ r

Ontology

Grammar G

Observations Stream

word w

Explanation

Derivation S -^ w

Partial Explanation

Derivation a A w

Table 1: Correspondence between Structural Induction and Generative Gram-
mars

I — ?> r under certain conditions and this new rule may contain in turn new
NonTerminals (internal variables) which are added to the already existing ones.
The common intuition is that / is a simpler version of r, as the final goal is to
reduce w to S. The terminals constitute the input symbols (standing for ob-
servables) , the NonTerminals constitute whatever additional "internal" variables
that are needed, the rewrite rules describe their interrelationship and altogether
they constitute the ontology. The correspondence between the terms used in
structural induction and generative grammars is summarized in Table [T]

Thus, in general, structural induction may invent any rewrite rule of the
form I — >■ r, potentially introducing new NonTerminals, the problem is that
there are infinitely many such rules that we could invent at any point in time.
In order to make the process more well defined we ask whether it is possible to
find a set of fundamental structural operations which is finite and minimalistic,
such that all the rules (or more precisely sets of rules) can be expressed in terms
of these operations. This would establish a normal form in terms of a finite
set of operations and then the problem of generating laws will be reduced to
making appropriate choices from this set without sacrificing completeness. In
the next subsection we will attempt to decompose the rules I — > r into a small
finite set of fundamental structural elements which will allow us to design better
structure search mechanisms.

2.3 ASNF Theorems

Issue (e — Construction). In the rest of the paper we will prove some theorems
that impose various sets of conditions on a grammar G in order for the grammar
to be considered in a certain Normal Form. If e £ L{G) however, we will allow
two specific rules of the grammar G to be exempted from these constraints and
still consider the grammar in the Normal Form. More exactly if e G L{G) and
given a grammar G" such that L(G") = L{G\{e}) and G" = {N',T,S',R') is in
a certain Normal Form then the grammar G = {N U {S}, T,S,R^ R' U{S ^
e,S — )■ S'}) where S ^ N' will also be considered in that certain Normal Form
despite the fact that the two productions {S —>■ e, S —>■ S'} may violate the
conditions of the Normal Form. These are the only productions that will be

allowed to violate the Normal Form conditions. Note that 5 is a brand new
NonTerminal and does not appear in any other productions aside from these two.
Without loss of generality we will assume in the rest of the paper that e ^ L{G).
This is because if e e L{G) we can always produce using the above-mentioned
construction a grammar G" that is in a certain Normal Form and L{G") = L{G')
from a grammar G" that is in that Normal Form and satisfies L{G') = L(G'\{e}).
We will call the procedure just outlined the e — Construction. We will call the
following statement the e — Amendment: Let G = {N, T, S, R) be a grammar,
if e is an element of the language L{G) one rule of the form 5 — > e is allowed
and furthermore the restriction that S does not appear in the right hand side
of any rule is imposed.

First we state a weak form of the Abstraction SuperStructuring Normal
Form for Context Free Grammars.

Theorem 2 (Weak-CFG-ASNF). Let G = {N,T,S,R), e ^ L{G) be a
Context Free Grammar. Then there exists a Context Free Grammar G" such
that L{G) — L{G') and G' contains only rules of the following type:

1. A^ B

2. A^ BC

3. A^ a

Proof . Since G is a CFG it can be written in the Chomsky Normal Form [13] .
That is, such that it contains only productions of the forms 2 and 3. If e e L{G)
a rule of the form 5 — > e is allowed and S does not appear in the rhs of any
other rule (e — Amendment). Since we have assumed that e ^ L{G) we do not
need to deal with e — Amendment and hence the proof.

D

Remarks.

1. We will call the rules of type 1 Renaniings (REN).

2. We will call the rules of type 2 Superstructures (SS) or compositions.

3. The rules of the type 3 are just convenience renaniings of observables into
internal variables in order to uniformize the notation and we will call them
Terminal (TERMINAL).

We are now ready to state the the Weak ASNF theorem for the general case.

Theorem 3 (Weak-GEN-ASNF). Let G ^ {N,T,S,R), e ^ L(G) be a
General (unrestricted) Grammar. Then there exists a grammar G' such that
L{G) — L{G') and G' contains only rules of the following type:

1. A^ B

2. A^ BC
2.. A^ a

4. AB^C

Proof . See Appendix.

Remark. We will call the rules of type 4 Reverse Super-Structuring (RSS).

In the next theorem we will strengthen our results by allowing only the re-
namings (REN) to be non unique. First we define what we mean by uniqueness
and then we proceed to state and prove a lemma that will allow us to strengthen
the Weak-GEN-ASNF by imposing uniqueness on all the productions safe re-
namings.

Definition (strong-uniqueness). We will say that a production a —>■ (3
respects strong-uniqueness if this is the only production that has the property
that it has a in the Ihs and also this is the only production that has /3 on the
rhs.

Lemma 2. Let G = (N, T, S", R), e ^ G a grammar such that all its produc-
tions are of the form:

1. A^ B

2. A^C,C^N

3. C-^ B ,C(^N

Modify the the grammar G to obtain G' = {N' , T, S' , R') as follows:

1. Introduce a new start symbol S' and the production S' — > S.

2. For each C ^ N that appears in the rhs of one production in G let {Ai — )■
C}i=i,n 0,^^ ^^6 ^^6 productions that contain ^ in the rhs of a production.
Introduce a new NonTerminal X(^ and the productions X(^ — ^ ^ and {Ai — >
X(}i=i,n and eliminate the old productions {Ai -^ C}i=i,n-

3. For each Q ^ N that appears in the Ihs of one production in G let {Q — ^
Sj}j=i,m all the the productions that contain C, the Ihs of a production.
Introduce a new NonTerminal Yq and the productions C, -^ Yq and {Yq — >
Bj}j=i,)n and eliminate the old productions {C — > Bj{.j^i„i.

Then the new grammar G" generates the same language as the initial grammar
G and all the productions of the form A -^ C, and C, ^ B , C, ^ N respect strong-
uniqueness. Furthermore, if the initial grammar has some restrictions on the
composition of the C, ^ N that appears in the productions of type 2 and 3, they
are respected since C, is left unchanged in the productions of the new grammar
and the only other types of productions introduced are renamings that are of
neither type 2 nor type 3.

Proof. See Appendix |19| .

By applying Lemma 2 to the previous two Weak-ASNF theorems we obtain
strong versions of these theorems which enforce strong-uniqueness in all the
productions safe the renamings.

Theorem 4 (Strong-CFG-ASNF). Let G = {N,T,S,R), e ^ L{G) be a
Context Free Grammar. Then there exists a Context Free Grammar G" such
that L{G) = L{G') and G' contains only rules of the following type:

1. A^ B

2. A — > BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs (strong-uniqueness) .

3. A — )■ a - and this is the only rule that has a in the rhs and this is the only
rule that has A in the Ihs (strong-uniqueness) .

Proof. Apply Lemma 2 to the grammar converted into Weak-CFG-ASNFD

Theorem 5 (Strong-GEN-ASNF). Let G ^ {N,T,S,R), e ^ L{G) be
a general (unrestricted) grammar. Then there exists a grammar G" such that
L{G) = L{G') and G' contains only rules of the following type:

1. A^ B

2. A — > BC - and this is the only rule that has BG in the rhs and this is the
only rule that has A in the Ihs (strong-uniqueness) .

3. A — > a - and this is the only rule that has a in the rhs and this is the only
rule that has A in the Ihs (strong-uniqueness) .

4. AB — > C - and this is the only rule that has C in the rhs and this is the
only rule that has AB in the Ihs (strong-uniqueness) .

Proof. Apply Lemma 2 to the grammar converted into Weak-GEN-ASNFD

Remark. After enforcing strong uniqueness the only productions that con-
tain choice are those of type 1 - renamings (REN) .

In the light of this theorem we proceed to introduce the concept of abstrac-
tion and prove some additional results.

2.4 Abstractions And Reverse Abstractions

Definitions (Abstractions Graph). Given a grammar G — {N,T, S,R)
which is in an ASNF from any of the Theorems 1 - 4 we call an Abstractions
Graph of the grammar G and denote it by AG{G) a Directed Graph G = {N, E)
whose nodes are the NonTerminals of the grammar G and whose edges are
constructed as follows: we put a directed edge starting from A and ending in B iff
A — >■ B is a production that occurs in the grammar. Without loss of generality,
we can assume that the graph has no self loops, i.e., edges of the form A^ A;\i
such self-loops exist, the corresponding productions can be eliminated from the
grammar without altering the language. In such a directed graph a node A has a
set of outgoing edges and a set of incoming edges which we refer to as out-edges
and in-edges respectively. We will call a node A along with its out-edges the
Abstraction at A and denote it ABS{A) = {A,OEa = {{A,B)\{A,B) e E}}.
Similarly, we will call a node A along with its in-edges the Reverse Abstraction
at A and denote it RABS{A) = {A,IEa = {{B,A)\{B,A) € E}}.

2.5 Grow Shrink Theorem

Theorem 6. Let G — {N,T, S, R), e ^ L{G) be a General Grammar. Then we
can convert such a grammar into the Strong-GEN-ASNF i.e., such that all the
productions are of the following form:

1. A^ B

2. A — > BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs. (strong-uniqueness)

3. v4 — > a - and this is the only rule that has A on the Ihs and there is no
other rule that has a on the rhs. (strong uniqueness)

4. AB — )• C - and this is the only rule that has C in the rhs and this is the
only rule that has AB in the Ihs. (strong-uniqueness)

And furthermore for any derivation w such that 7 — > w , in G, ^ ^ N^ there
exists a derivation 7 — > /i — > i^ — > w such that /i G N^ , v £ N* and 7 -^ /i
contains only rules of type 1 and 2 (REN, SS), /i ^- a contains only rules of the
type 1, more particularly only Reverse Abstractions and type 4 (REN(RABS),
RSS) and v ^ w contains only rules of type 3 (TERMINAL).

Proof. See Appendix.

We have therefore proved that for each General Grammar G we can trans-
form it in a Strong-GEN-ASNF such that the derivation (explanation in struc-
tural induction terminology) of any terminal string w can be organized in three
phases such that: Phase 1 uses only productions that grow (or leave unchanged)
the size of the intermediate string; Phase 2 uses only productions that shrink
(or leave unchanged) the size of the intermediate string; and Phase 3 uses only
TERMINAL productionfj^ In the case of grammars that are not in the nor-
mal form as defined above, the situation is a little more complicated because
of successive applications of grow and shrink phases. However, we have shown
that we can always transform an arbitrary grammar into one that in the normal
form. Note further that the grow phase in both theorems use only context free
productions.

We now proceed to examine the implications of the preceeding results in the
larger context including the nature of hidden variables, radical positivism and
the David Hume's principles of connexion among ideas.

3 The Fundamental Operations of Structural In-
duction

Recall that our notion of structural induction entails: Given a sequence of ob-
servations w we attempt to find a theory (grammar) that explains w and si-
multaneously also the explanation (derivation) S ^ w. In a local way we may

^At first sight, it may seem that this construction offers a way to solve the halting problem.
However, this is not the case, since we do not answer the question of deciding when to stop
expanding the current string and start shrinking which is key to solving the halting problem.

think that whenever we have a production rule I — )■ r that I explains r. In a
bottom up - data driven way we may proceed as follows: First introduce for
every observable a a production A — > a. The role of these productions is sim-
ply to bring the observables into the realm of internal variables. The resulting
association is between the observables and the corresponding internal variables
unique (one to one and onto) and hence, once this association is established, we
can forget about the existence of bbservables (Terminals). Since establishing
these associations is the only role of the TERMINAL productions, they are not
true structural operations. With this in mind, if we are to construct a theory
in the GEN-ASNF we can postulate laws of the following form:

1. yl — > BC - Super-structuring (SS) which takes two internal variables B
and C that occur within proximity of each other (adjacent) and labels the
compound. Henceforth, the shorter name A can be used instead for BC .
This is the sole role of super-structuring - to give a name to a composite
structure to facilitate shorter explanations at latter stages.

2. A ^ B\C - Abstraction (ABS). Introduces a name for the occurrence of
either of the variables B or C. This allows for compactly representing two
productions that are identical except that one uses B and the uses C by
a single production using A. The role of Abstraction is to give a name
to a group of entities (we have chosen two only for simplicity) in order
to facilitate more general explanations at latter stages which in turn will
produce more compact theories.

3. AB — > C - Reverse Super-structuring (RSS) which introduces up to two
existing or new internal variables that are close to each other (with respect
to a specified topology) that together "explain" the internal variable C

4. A — !> C, i? — > C - Reverse Abstraction (RABS) which uses existing or
new internal variables A and B as alternative explanations of the internal
variable C (we have chosen two variables only for simplicity).

3.1 Reasons for Postulating Hidden Variables

Recall that are at least two types of reasons for creating Hidden Variables:

1. (OR type) - [multiple alternative hidden causes] The OR type corresponds
to the case when some visible effect can have multiple hidden causes
HI — >■ Effect, H2 — ?> Effect . In our setting, this case corresponds
to Reverse Abstraction. One typical example of this is: The grass is wet,
and hence either it rained last night or the sprinkler was on. In the sta-
tistical and machine learning literature the models that use this type of
hidden variables are called mixture models [9_.

2. (T-AND type) - [multiple concurrent hidden causes] The T-AND type,
i.e., topological AND type, of which the AND is a sepcial case. This cor-
responds to the case when one visible effect has two hidden causes both of

10

which have to occur within proximity of each other (with respect to a spec-
ified topology) in order to produce the visible effect. H1H2 — >■ Effect.
In our setting, this corresponds Reverse Super-structuring. In the Sta-
tistical / Graphical Models literature the particular case of AND hidden
explanations is the one that introduces edges between hidden variables in
the depedence graph IH], [S], [TT] .

The preceeding discussion shows that we can associate with two possible reasons
for creating hidden variables, the structural operations of Reverse Abstraction
and Reverse Super Structuring respectively. Because these are the only two
types of productions that introduce hidden variables in the GEN-ASNF this
provides a characterization of the rationales for introducing hidden variables.

3.2 Radical Positivism

If we rule out the use of RSS and RABS, the only operations that involve the pos-
tulation of hidden variables, we are left with only SS and ABS which corresponds
to the radical positivist [T] stance under the computationalist assumption. An
explanation of a stream of observations w in the Radical Positivist theory of the
world is mainly a theory of how the observables in the world are grouped into
classes (Abstractions) and how smaller chunks of observations are tied together
into bigger ones (Super-Structures). The laws of the radical positivist theory
are truly empirical laws as they only address relations among observations.

However, structural induction, if it is constrained to using only ABS and
SS, the class of theories that can be induced is necessarily a subset of the set of
theories that can be described by Turing Machines. More precisely, the resulting
grammars will be a strict subset of Context Free Grammars, (since CFG contain
SS, and REN(ABS+RABS)). Next we will examine how any theory of the world
may look like from the most general perspective when we do allow Hidden
Variables.

3.3 General theories of the world

If structural induction is allowed to take advantage of RSS and RABS in addition
to SS and ABS, the resulting theories can make use of hidden variables. Obser-
vations are a derivative byproduct obtained from a richer hidden variable state
description by a reduction: either of size - performed by Reverse SuperStruc-
turing or of information - performed by Reverse Abstraction. Note that, while
in general, structural induction can alternate several times between REN+SS
and RABS+RSS, we have shown that three phases suffice: a growth phase
(REN+SS); a shrink phase (RABS+RSS); and a Terminal phase. Whether we
can push all the RABS from the first phase into the second phase and make
the first phase look like the one in the radical positivist stance (only ABS+SS)
remains an open question (See Appendix for a Conjecture to this effect).

3.4 Hume's principles of connexion among ideas

We now examine, against the backdrop of GEN-ASNF theorem, a statement
made by philosopher David Hume more that 2 centuries ago: "/ do not find

11

that any philosopher has attempted to enumerate or class all the principles of
association [of ideas]. ... To me, there appear to be only three principles of
connexion among ideas, namely. Resemblance, Contiguity in time or place, and
Cause and Effect" JUj. If we substitute Resemblance with Abstraction (since
abstraction is triggered by resemblance or similarity) , Contiguity in time or place
with Super-Structuring (since proximity, e.g., spatio-temporal proximity drives
Super-Structuring) and Cause and Effect with the two types of explanations that
utilize hidden variables, it is easy to see that the GEN-ASNF theorem is simply
a precise restatement of Hume's claim under the computationalist assumption.

4 Summary

We have shown that abstraction (grouping similar entities) and super-structuring
(combining topologically e.g., spatio-temporally close entities) as the essential
structural operations in the induction process. A structural induction process
that relies only on abstraction and super-structuring corresponds to the radical
positivist stance. We have shown that only two more structural operations,
namely, reverse abstraction and reverse super- structuring (the duals of abstrac-
tion and super-structuring respectively) (a) suffice in order to exploit the full
power of Turing-equivalent generative grammars in induction; and (b) opera-
tionalize two rationales for the introduction of hidden variables into theories of
the world. The GEN-ASNF theorem can be seen as simply a restatement, under
the computationalist assumption, of Hume's 2-century old claim regarding the
principles of connexion among ideas.

References

[1] A.J. Ayer, Language, Truth, and Logic. London: GoUancz. (2nd edition,
1946), 1936.

[2] M. Burgin. Super-Recursive Algorithms. Springer, 2005.

[3] R. Carnap. An introduction to the Philosophy of Science. 1966.

[4] N. Chomsky. Syntactic Structures. The Hague: Mouton. 1957.

[5] G. Elidan and N. Friedman. Learning Hidden Variable Networks: The
Information Bottleneck Approach. Journal of Machine Learning Research
(JMLR), 6:81-127, 2005.

[6] D. Hume. An Enquiry Concerning Human Understanding . Hackett Publ
Co. 1993.

[7] M. Hutter. Universal Artificial Intelligence: Sequential Decisions based on
Algorithmic Probability. EATCS, Springer, 2005.

[8] S.-Y. Kuroda. Classes of languages and linear-bounded automata. Infor-
mation and Control, 7(2): 207-223, 1964.

12

[9] S. L. Lauritzen. Graphical Models. Clarendon Press, Oxford, 1996.

[10] T. Oates, T. Armstrong, J. Harris and M. Nejnian. On the Relationship
Between Lexical Semantics and Syntax for the Inference of Context-Free
Grammars. Proceedings of the 19th National Conference on Artificial In-
telligence ({AAAI}), 431-436, 2004.

[11] J. Pearl. Probabilistic Reasoning in Intelligent Systems. Morgan Kaufmann
PubUshers, 1988.

[12] K. R. Popper. The Logic of Scientific Discovery, Basic Books (English ed.
1959), 1934.

[13] A. Salomaa. Computation and Automata. Cambridge University Press,
1985.

[14] W. Savitch. How to make arbitrary grammars look like context-free gram-
mars. SIAM Journal on Computing, 2:174-182, 1973.

[15] J. Schmidhuber, J. Zhao and M. Wiering Shifting Bias with Success Story
Algorithm. Machine Learning, 28:105-130, 1997.

[16] R. SolomonofF. A Formal Theory of Inductive Inference, Part I. Information
and Control, 7(l):l-22, 1964.

[17] R. SolomonofF. A Formal Theory of Inductive Inference, Part II. Informa-
tion and Control, 7(2):224-254, 1964.

[18] A. Turing. On computable numbers with an application to the
Entscheuidungs-problem, Proc. Lond. Math. Soc, Ser.2, 42:230-265, 1936.

[19] A. Silvescu and V. Honavar. Abstraction Super-structuring Normal Forms:
Towards a Computationalist Theory of Structural Induction. Technical Re-
port. Department of Computer Science. Iowa State University. 2011.

13

Appendix

Theorem 3 (Weak-GEN-ASNF). Let G = {N,T,S,R), e ^ L{G) be a
General (unrestricted) Grammar. Then there exists a grammar G' such that
L{G) = L{G') and G' contains only rules of the following type:

1. A^ B

2. A^ BG

3. A^ a

4. AB ^G

Proof . Every general grammar can be written in the Generalized Kuroda
Normal Form(GKNF) [13_. That is, it contains only productions of the form:

1. A^ e

2. AB -^ GD

3. A^ BG

4. A^ a

Assume that we have akeady converted our grammar in the GKNF. We will
prove our claim in 3 steps.

Step 1 . For each {AB — > GD} we introduce a new NonTerminal Xab.cd
and the following production rules {AB —>■ Xab,cd} and {Xab.cd —^ GD}
and eliminate the old {AB — > GD} production rules. In this way we ensure
that all the rules of type 2 in GKNF have been rewritten into rules of types 2
and 4 in the GEN-ASNF.

The new grammar generates the same language. To see this let oldG =
(NoidG, T, S, RoMg) denote the old grammar and newG = [NnewG, T, S, Rnewc)
denote the new grammar. Let 7 G -^J^^g and 7 — > w be a derivation in oldG.
In this derivation we can replace all the uses of the production AB -^ GD
with the sequence AB — > Xab,cd — ^ GD and get a derivation that is valid in
newG, since all the other productions are common between the two grammars.
Thus we have proved that for all 7 e ^omg ^^ *^^^ convert a valid derivation
from oldG, 7 — ^ w into a valid derivation in newG and in particular this is
true also for S -^ w. Therefore L{oldG) C L{newG). Conversely, let 7 €
■^oidG ^^'^ 7 — ^ w be a valid derivation in newG then whenever we use the
rule AB -> Xab.gd in this derivation 7 — > aABjS — > uXab.gdP — >■ w let
7 — >• aABfi -^ aXABfioP -^ SXab,cd'>] -^ SGDrj -^ w he the place where
the Xab,cd that occurs between a and /3 is rewritten (used in the Ihs of a
production, even if it rewrites to the same symbol) for the first time. Then
necessarily the Xab,cd -^ GD rule is the one that applies since it is the only
one that has Xab.cd in the Ihs. Furthermore, as a consequence of Lemma 1 we
have that a — > (5 and (3 ^ rj are valid derivations in newG. Therefore we can

14

bring the production Xab.cd -^ CD right before the use of AB -^ Xab,cd as
follows 7 A aABl3 -> aXAB,CDl3 -> aCD/S A SCD/S A SCDrj A w and still
have a valid derivation in newG. We can repeat this procedure for all places
where rules of the form AB — > Xab.cd appear. In this modified derivation
we can replace all the uses of the sequence AB — > Xab.cd — > CD with the
production AB — >■ CD and obtain a derivation that is valid in oldC since all
the other productions are common between the two grammars. Thus we have
proved that for all 7 € N^^^^ we can convert a valid derivation from newG
,7 — )■ w into a valid derivation in oldG and in particular this is true also for
S ^ w since S G N^^^q, and therefore L{newG) C L{oldC). Therefore we have
proved that the grammars are equivalent, i.e., L{oldG) — L{newG).

Step 2 . Returning to the main argument, so far our grammar has only rules
from Weak-GEN-ASNF safe for the e-productions A^f e. Next we will eliminate
e-productions ^ — > e in two steps. First for each ^ — )■ e we introduce a new
NonTerminal Xa.e and the productions Xa.e — > D and E ^ e and eliminate
A — > e, where _E is a distinguished new NonTerminal (that will basically stand
for e internally). This insures that we have only one e-production, namely i? — >■ e
and E does not appear on the Ihs of any other production and also that all the
rules that rewrite to E are of the form A ^^ E.

The new grammar generates the same language. We will use a similar proof
technique as in the previous step. Let oldC = {NoidcT^S^Roidc) denote the
old grammar and newG ~ {NnewG, T, S, Rnewo) denote the new grammar. Let
7 € N^ij^^ and 7 — ^ w be a derivation in oldC. In this derivation we can replace
all the uses of the production ^ — > e with the sequence Xa.e —>■ E —>■ e and get
a derivation that is valid in newC, since all the other productions are common
between the two grammars. Thus we have proved that for all 7 € N^i^q we can
convert a valid derivation from oldG ,7 — >■ w into a valid derivation in newG
and in particular this is true also for S ~> w, therefore L{oldG) C L{newG).
Conversely, let 7 G ^tidc ^-'^^ 7 — > w be a valid derivation in newG then
whenever we use the rule Xa,e -^ E in this derivation 7 -^ uXa.eP -^ aE/S — >
w let 7 — > uXa.eP —>■ aE/3 — > SErj -^ Sr] —>■ w he the place where the E that
occurs between a and /3 is rewritten (used in the Ihs of a production, even if it
rewrites to the same symbol) for the first time. Then necessarily the E -^ e rule
is the one that applies since it is the only one that has E in the Ihs. Furthermore,
as consequence of Lemma 1 we have that a ^^ 5 and /3 ^ 77 are valid derivations
in newG. Therefore we can bring the production _E — > e right before the use of
Xa^e -^ E as follows 7 — >■ aXA,Ef3 —>■ aE/3 — >■ a/3 -^ 5/3 ^^ Srj ^ w and still
have a valid derivations in newG. We can repeat this procedure for all places
where rules of the form Xa,e -^ E appear. In this modified derivation we can
replace all the uses of the sequence Xa.e — >■ i? — >■ e with the production A — > e
and obtain a derivation that is valid in oldG since all the other productions are
common between the two grammars. Thus we have proved that for all 7 £ ^oidc
we can convert a valid derivation from newG, 7 — > w into a valid derivation
in oldG and in particular this is true also for S ^ w since S G ^oZdc ^^'^

15

therefore L{newG) C L{oldG). Therefore we have proved that the grammars
are equivalent, i.e., L{oldG) — L{newG).

Step 3 . To summarize: the new grammar has only rules of the Weak-GEN-
ASNF type, safe for the production E ^ e which is the only production that has
E in the Ihs and there is no other rule that has e on the rhs (strong-uniqueness)
and furthermore the only rules that contain E in the rhs are of the form A ^>- E
{only renamings to E).

We will eliminate the e-production E -^ e as follows: Let {Ai -^ i?}i=i,n be
all the productions that have E on the rhs. For all NonTerminals X G iV U T
introduce productions {XA^ — > X}j=i „ and {AiX — > Xj^^i „ and eliminate
{Ai — )• -E}i=i,„, furthermore we also eliminate i? — > e.

The new grammar generates the same language. We will use a similar proof
technique as in the previous step. Let oldG =■ {NoidG,T, S, Roidc) denote the
old grammar and newG = {NnewGiT, S, Rnewc) denote the new grammar. Let
7 e N^ij^fj and 7 — > w be a derivation in oldG. Then whenever we use a rule of
the form A^ — > i? in this derivation 7 — > aAij3 — > aEfS — > ui let 7 — > aAij3 — >
aEj3 — > 5Eri — > J?] — )■ w be the place where the E that occurs between a
and P is rewritten(used in the Ihs of a production, even if it rewrites to the
same symbol) for the first time. Then necessarily the _E — > e rule is the one
that applies since it is the only one that has E in the Ihs. Furthermore, as a
consequence of Lemma 1 we have that a — >■ (5 and /3 ^ rj are valid derivations
in oldG. Therefore we can bring the production _E — > e right before the use of
Ai —>■ E 'AS follows 7 -^ aAiP — )■ aE/3 — )■ a/3 -^ 5/3 -^ dr] —>■ w and still have
a valid derivations in oldG. We can repeat this procedure for all places where
rules of the form A^ — > i? appear. In this modified derivation we can replace
all the uses of the sequence Ai ^S' E ^ e with the production XAi -^ X ii
a y^ e and XAi -^ X if (3 ^ e and obtain a derivation that is valid in newG
since all the other productions are common between the two grammars, (e.g.,
if a 7^ e, a = aiX replace aAi/3 — aiXAijS -> aiXE/S — >■ aiXj3 = aj3
which is valid in oldG with aAiP = oiiXAiP — > aiXj3 = a/3 which is valid in
newG and similarly for /3 7^ e). In this way since all the other productions are
common between the two grammars we can convert a derivation 7 -^oidG ^ into
a derivation 7 -^newG w. Note that it is not possible for both a and /3 to be
equal to e because this will imply that the derivation 7^wis7^^j^i?^e
which contradicts the hypothesis that w € T~^ . Thus we have proved that for
all 7 S N^i^Q we can convert a valid derivation from oldG, j —>■ w into a valid
derivation in newG and in particular this is true also for S ^ w since S G ^oidG^
and therefore L{oldG) C L{newG).

Conversely, let 7 € ^oidG ^^'^ 7 — > w be a derivation in newG . In this
derivation we can replace all the uses of the production XAi -^ X with the
sequence Ai ^ E —> e and get a derivation that is valid in oldG since all the other
productions are common between the two grammars (e.g., we replace aAi/3 =
aiXAiP — )■ aiX/3 = a/3 which is valid in newG with aAi(3 = aiXAi/3 — >■
aiX Ej3 — > aiX/3 = a/3 which is valid in oldG). We proceed similarly with
the productions of the form AiX -^ X and replace them with the sequence

16

Ai ^^ E ^^ e and get a derivation that is valid in oldG. Thus we have proved
that for all 7 S ^oidG ^^ '^^^ convert a valid derivation from newG, 7 — >■ w into
a valid derivation in oldG and in particular this is true also for S -^ w, therefore
L{newG) C L(oldG). Hence we have proved that the grammars are equivalent,
i.e., L{oldG) — L{newG).

D

Lemma 1. Let G ~ {N, T, S, R) be a grammar and let a^j3 — > 5^iri, fi G
(NUT)~^ a valid derivation in G that does not rewrites the fi (uses productions
whose Ihs m,atch any part of fj, even if it rewrites to itself), occurring between a
and /3, then a ^^ 5, j3 ^^ rj are valid derivations in G.

Proof. Because a^ji — > S^rj does not rewrites any part of /i (even to itself) it
follows that the Ihs of any production rule that is used in this derivation either
matches a string to the left of /i or to the right of /i. If we start from a and
use the productions that match to the left in the same order as in aXj3 — >■ 5Xrj
then we necessarily get a — > (5, a valid derivation in G. Similarly, by using the
productions that match to the right of /i we get (i ^ rj valid in G.

D

Lemma 2. Let G ~ (N, T, S, R), e ^ G a grammar such that all its produc-
tions are of the form:

1. A^ B

2. A^CCiN

Modify the the grammar into a new grammar G' = {N' ,T, S' , R') obtained as
follows:

1. Introduce a new start symbol S' and the production S' — > S.

2. For each C, ^ N that appears in the rhs of one production in G let {Ai — >
C}i=i,n '^^^ ^^6 ^^6 productions that contain C, in the rhs of a production.
Introduce a new NonTerminal X(^ and the productions X,^ — > ( and {Ai — >
X(}i=i,n o,nd eliminate the old productions {Ai — >■ C}i=i^„.

3. For each Q ^ N that appears in the Ihs of one production in G let {C, — >
Bj}j=i^m 0,11 the the productions that contain (. the Ihs of a production.
Introduce a new NonTerminal Y"^ and the productions C — > Y^ and {Y"^ — >
Bj}j^i„i and eliminate the old productions {(^ — > Bj}j^i„i-

Then the new grammar G" generates the same language as the initial grammar
G and all the productions of the form A ^>- ( and C ^^ B , C, ^ N respect
strong-uniqueness. Furthermore if the initial grammar has some restrictions on
the composition of the (^ ^ N that appears in the productions of type 2 and
3, they are still maintained since ( is left unchanged in the productions of the
new grammar and the only other types of productions introduced are Renamings
which do not belong to either type 2 or 3.

17

Proof. To show that the two grammars are equivalent we will use a similar
proof technique as in the previous theorem. Let 7 £ N~^ and 7 — >■ w be a
derivation in G. In this derivation we can replace all the uses of the production
Ai ^ C, with the sequence Ai -^ X^^ -^ C and the uses of the production ( -^ Bj
with the sequence C — )■ F^ — t- Bj and get a derivation that is valid in G, since all
the other productions are common between the two grammars. Thus we have
proved that for all 7 G 7V+ we can convert a valid derivation from G, 7 — > w
into a valid derivation in G' and in particular this is true also for S —^g w,
which can be converted into S — >g' w and which furthermore can be converted
into S' ^ S Ac w and therefore i(G) C L(G').

Conversely, let 7 G TV"*" and 7 — )■ w be a valid derivation in G' then whenever
we use the rule Ai -^ Xq in this derivation 7 — > aAil3 -^ aX^jS — )■ w let
7 -^- aAiP -^ aX^P -^ SX(^rj -^ SQ-q ^ w be the place where the X(^ that occurs
between a and fi is rewritten (used in the Ihs of a production, even if it rewrites
to the same symbol) for the first time. Then necessarily the Xq — >■ C, rule is the
one that applies since it is the only one that has Xq in the Ihs. Furthermore, as
consequence of Lemma 1 we have that a — )■ J and j3 ^^ rj are valid derivations
in G'. Therefore we can bring the production X(^ — > C, right before the use of
Ai -^ Xi^ as follows 7 — > aAij3 — > aX^/3 — > aC,j3 — > 5C,(i -^ 6(1] -^ w and still
have a valid derivation in G'. We can repeat this procedure for all places where
rules of the form Ai — )■ X^ appear. Similarly for the uses of the productions of
the type C — >■ y^ in a derivation 7 — > w we can bring the production that rewrites
Xq {Yq -^ Bj) right after as follows: change 7 -^ a(f3 -^ aY^l3 -^ SYi^rj -^
SBjrj — > w into 7 — > a(/3 — > aY^fi — > aBj/3 — > SBj/3 — > SBjrj — > w, because
from Lemma 1 we have that a — > i5 and 13 ^^ rj. We can repeat this procedure
for all places where rules of the form C — > Xq appear. In the new modified
derivation we can replace all the uses of the sequence Ai — > Xq — > C, with the
production ^j — > ^ and the sequence C — > F^ — > Bj with the production C, -^ Bj
and obtain a derivation that is valid in G since all the other productions are
common between the two grammars. Thus, we have proved that for all 7 G A^^
we can convert a valid derivation from G' ,7 — >■ w into a valid derivation in G.
For a derivation and 5" -^q' w we have that necessarily S' ^ S -^g' w since
S" ^ 5 is the only production that rewrites S' but since 5* G iV+, it follows
that we can use the previous procedure in order to convert S — ^g' w into a
derivation S -^g w which proves that L{G') C L{G). Hence, we have proved
that the grammars are equivalent, i.e., L(G) = L{G').

It is obvious that all the productions of the form A — > C and C, ^^ B ,
C ^ N respect strong-uniqueness. Furthermore if the initial grammar has some
restrictions on the composition of the ( ^ N that appear in the productions of
type 2 and 3 they are still maintained since ( is left unchanged in the productions
of the new grammar and the only other types of productions introduced are
Renamings which do not belong to either type 2 or 3.

D

18

Minimality

We can ask the question whether we can find even simpler types of structural
elements that can generate the full power of Turing Machines. Two of our
operators, RSS and SS require that the size of the production {\lhs\ + \rhs\) is
3. We can ask whether we can do it with size 2. This is answered negatively by
the following proposition.

Proposition (Minimality) If we impose the restriction that \lhs\ + \rhs\ <
2 for all the productions then we can only generate languages that are finite sets
Terminals, with the possible addition of the empty string e.

Proof . The only possible productions under this constraint are:

1. A^ a

2. A^ e

3. A^ B

4. AB ^ e

5. Aa-^ e

All the derivations that start from S, either keep on rewriting the current Non-
Terminal, because there is no production which increases the size, or rewrite
the current NonTerminal into a Terminal or e.

D

Another possible question to ask is whether we can find a smaller set of
structural elements set or at least equally minimal. The following theorem is a
known result by Savitch [14 .

Theorem 7 (Strong-CFG-ASNF-Savitch). (Savitch 1973 ml) Let G =
{N, T, S, R), e ^ L{G) be a General Grammar then there exists a grammar G'
such that L{G) = L{G') and G" contains only rules of the following type:

1. AB ^e

2. A^ BC

i. A^ a

The Savitch theorem has three types of rules TERMINAL (type 3), SS (type 2)
and ANNIHILATE2 (type 1). However no strong-uniqueness has been enforced;
It it were it to be enforced, then one more type of rule, REN will be needed.
Furthermore if we would not insist on uniqueness all the Renamings in the ASNF
could be eliminated too (Renamings elimination is a well known technique in
the theory of formal languages |13j ) . However this will make it very difficult
to make explicit the Abstractions. Therefore it can be argued that ASNF and
the Savitch Normal Form have the same number of fundamental structural
operations with the crucial difference between the two being the replacement
of the RSS with ANNIHILATE2. The Reverse SuperStructuring seems a more
intuitive structural operation however, at least from the point of view of this

19

paper.. Next we present for completeness the version of Savitch theorem where
strong-uniqueness is enforced for rules of type 2 and 3.

Theorem 8 (Strong-GEN-ASNF-Savitch). Let G = {N,T,S,R), e (^
L{G) he a General Grammar then there exists a grammar G' such that L{G) =
L{G') and G' contains only rules of the following type:

1. A^ B

2. yl — > BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs. (strong-uniqueness)

3. A — > a - and this is the only rule that has a in the rhs and this is the only
rule that has A in the Ihs. (strong-uniqueness)

4. AB — > e - and this is the only rule that has AB on the Ihs. (uniqueness)

Proof. By the original Savitch theorem [14] any grammar can we written such
that it only has rules of the type

1. AB ^ e

2. A^ BC

3. A^ a

The rules of the form AB — ;> e observe uniqueness by default since the only rules
that have more than one symbol in the Ihs are of the form AB — >■ e.

Then we can use the constructions in Lemma 2 on this grammar in order
to enforce strong-uniqueness for SuperStructuring (SS) A — >■ BG and Termi-
nals (TERMINAL) A — ?► a at the potential expense of introducing Renamings
(REN).

D

Grow &; Shrink theorems

Theorem 9. Let G = {N,T, S,R), e ^ L{G) be a General Grammar in the
Strong-GEN-ASNF-Savitch, i.e., all the productions are of the following form:

1. A^ B

2. A -^ BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs. (strong-uniqueness)

3. A — > a - and this is the only rule that has A on the Ihs and there is no
other rule that has a on the rhs. (strong uniqueness)

4. AB — ^ e - and this is the only rule that has AB on the Ihs. (uniqueness)

20

Then for any derivation w such that 7 — > w , in G, ^ £ N~^ there exists a
derivation 7 — > ^ ^ j^ ^ w such that fj. G N^ , v e N* and 7 — > /u contains
only rules of type 1 and 2 (REN, SS), /i — > a contains only rules of the type 4
(ANNIHILATE2) and v ^ w contains only rules of type 3 (TERMINAL).

Proof. Based on Lemma 3 (presented next) we can change the derivation
7 — > w into a derivation 7 — > j^ — >■ w such that the segment i/ — > w contains
only rules of type 3 (TERMINAL) and 7^1^ contains only rules of type 1,
2 or 4 (REN, SS, ANNIHILATE2) . Therefore the only thing we still have to
prove is that we can rearrange 7 — >■ z^ into 7 — > /i — >■ z^ such that 7 — >■ /i contains
only rules of type 1 or 2 (REN, SS) and /i — > i^ contains only rules of the type
4 (ANNIHILATE2).

Let 7 € -/V"*" , and w such that 7 — > i^ — > w is a derivation in G , the segment
V ^ w contains only rules of type 3 (TERMINAL) and 7 — > a contains only
rules of type 1, 2 or 4 (REN, SS, ANNIHILATE2) . If the derivation already
satisfies the condition of the lemma then we are done. Otherwise examine in
order the productions in 7 — > i^ from end to the beginning until we encounter the
first rule of type 4: AB — )■ e that violates the condition required by the theorem
and at least one production of the type 1 or 2 has been used after it (otherwise
we have no violation) . More exactly, prior to it only rules of type 1 or 2 were

used and at least one such rule was used. That is 7 — > aAB/3 — > a/S — > /i — > z^
and only rules of the type 1 or 2 have been used in a/3 — >■ /x, and only rules of
the type 4 have been used in y^' — > i^. Because only rules of the type 1 or 2
(which never have more than one symbol in the Ihs) have been used in af3 — > fj,
it follows that there exists /xi,/i2 G Af*such that a — > /ii and /3 — > /i2 and
H = /xi/^2- Therefore we can rearrange the rewriting 7 — > aABfi — > aj3 — >■ /x
into 7 — >■ aAB/3 — > ^iAB/3 — > ^iAB^2 ^^ /^iM2 = /i . In this way we have
obtained a derivation for ji m G that violates the conclusion of the lemma in
one place less than the initial derivation. Since there is a finite number of steps
in a derivation and therefore a finite number of places where the constraints can
be violated it can be inferred that after a finite number of applications of the
above-described "swapping" procedure we will obtain a derivation which satisfies
the rules of the theorem.

D

Lemma 3. Let G = {N, T, S, R) he a general (unrestricted) grammar that
contains only rules of the form:

1. a^/3, ae iV+,/3e N*

2. A^ a

Then for any derivation w such that j —>■ w , in G, j E N^ there exists a
derivation 7 — > j/ — > w such that v G N^ and 7 — > i^ contains only rules of type
1 and J/ — > ui contains only rules of type 2 (TERMINAL).

Proof. Let w such that "f ^ w in G, j G N^. If the derivation already
satisfies the condition of the lemma then we are done. Otherwise examine in

21

order the productions 7 ^ w until we encounter a rule of type , say it is A -^- a
, such that there are still rules of type 1 used after it 7 — >■ aA/3 — ?• aa/3 — t- w.
Because none of the rules in the grammar contain terminals in their Ihs it
follows that there exists Wi,W2 G T* such that a ~^ wi and /3 — > W2 and
w = Wiaw2- Therefore we can rearrange the rewriting 7 — > aA^ — > aa/3 — >■ w
into 7 — > aAf3 — > wiA/3 -^ wiAw2 — >■ 'Wiaw2 — w. In this way we have obtained
a derivation for w in G that violates the conclusion of the lemma in one place less
than the initial derivation. Since there is a finite number of steps in a derivation
and therefore a finite number of places where the constraints can be violated it
can be inferred that after finite number of application of the above-described
"swapping" procedure we will obtain a derivation which satisfies the rules of the
lemma.

D

Theorem 6. Let G = {N,T,S,R), e ^ L{G) be a General Grammar. Then
we can convert such a grammar into the Strong-GEN-ASNF i.e., such that all
the productions are of the following form:

1. A^ B

2. A — > BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs. (strong-uniqueness)

3. A — > a - and this is the only rule that has A on the Ihs and there is no
other rule that has a on the rhs. (strong uniqueness)

4. AB — > C - and this is the only rule that has C in the rhs and this is the
only rule that has AB in the Ihs. (strong-uniqueness)

And furthermore for any derivation w such that 7 — > w , m G, 7 G A^+ there
exists a derivation 7 — > /i — > z^ — > w such that 11 G N^ , v G N* and 7 — > /i
contains only rules of type 1 and 2 (REN, SS), /i — > a contains only rules of the
type 1, more particularly only Reverse Abstractions and type 4 (REN(RABS),
RSS) and i^ — > w contains only rules of type 3 (TERMINAL).

Proof Sketch. By Theorem 6 we can convert the grammar G into a grammar
G' in Strong-GEN-ASNF-Savitch. Then we can convert such a grammar into a
grammar G" in Strong-GEN-ASNF as follows: for all {AiBi -^ e}i=i_„ introduce
new NonTerminals {XAiBi}i=i,n and add AiBi -^ ^AiBt and XAiBi -^ E
and eliminate the original productions {AiBi — )■ e\i=i^n- Furthermore, for all
NonTerminals X ^ E in the new grammar, add new NonTerminals Xxe, Xex
and Xee and the production rules XE — >• Xxe, EX — >• Xex, Xxe -^ X,
Xex — > X, EE — )• Xee and Xee — >■ E. We can easily show using techniques
already developed that the new grammar will generate the same language as
the previous one and that it respects the strong-uniqueness for rules of type 2,
3 and 4.

Furthermore if we take a derivation for a string w in the grammar G' in the
Strong-GEN-ASNF-Savitch S -^c' "^ then from Theorem 7 we know that we
can convert it into a derivation that uses first only REN and SS in the phase

22

1, then only ANNIHILATE2 in phase 2 and finally only TERMINAL in phase
3. We can take such a derivation and replace the usage of the ANNIHILATE2
productions in phase 2 with productions as above in order to get a derivation
in G" . Note that the productions introduced above are meant to transform
the ANNIHILATE2 into rules that follow the Strong-GEN-ASNF, and the only
types of rules that we have introduced are RSS with strong-uniqueness holding
and REN of the RABS type. This proves the theorem.
D

Further Discussions and a Conjecture

We have therefore proved that for each General Grammar G we can transform it
both in a Strong-GEN-ASNF-Savitch and a Strong-GEN-ASNF such that the
derivation (explanation in Structural Induction terminology) of any terminal
string w can be organized in three phases such that: In Phase 1 we use only
productions that grow (or leave the same) the size of the intermediate string;
In Phase 2 we use only productions that shrink (or leave the same) the size of
the intermediate string and in Phase 3 we use only TERMINAL productions.

Naively, at first sight, it may seem that this is a way to solve the halting
problem and therefore maybe some mistake has been made in the argument.
However this is not the case, as the key question is when to stop expanding the
current string and start shrinking and this problem still remains. In a certain
sense these two theorems are a way to give a clear characterization of the issue
associated with solving the halting problem: namely that of knowing when to
stop expanding. In the case of Grammars in arbitrary forms the issue is a little
bit more muddled as we can have a succession of grow and shrink phases but
we have shown that if the form is constrained in a certain ways then we only
need one grow and one shrink. Note also that during the grow phase in both
theorems we are only using context free productions.

Structural Induction and the fundamental structural ele-
ments

In this section we will review the Structural Induction process in the light of the
concepts and results obtained for Generative Grammars and discuss the role of
each of the operators. Then we move on to make some more connections with
already existing concepts in Statistics and Philosophy and show how the ASNF
affords for very precise characterizations of these concepts.

In the context of Generative Grammars Structural Induction is concerned
with the following question: Given a sequence of observations w we attempt to
find a theory (grammar) that explains w and simultaneously also the explanation
(derivation) S ^>- w. In a local way we may think that whenever we have a
production rule I — >■ r that I explains r. In a bottom up - data driven way we
may proceed as follows: First introduce for every Observable a a production
j4 — > a. These are just convenience productions that bring the observables into
the realm of internal variables in order to make everything more uniform. The

23

association is unique (one to one and onto) and once we have done it we can
forget about the existence of Observables (Terminals). This is only the role of
the the TERMINAL productions, and for this reason we will not mention them
in future discussions as they are not true structural operations. With this in
mind if we are to construct a theory in the GEN-ASNF we can postulate any
of the following laws:

1. SS - A — > BC - SuperStructuring. Takes two internal variables B and
C that occur within proximity of each other (adjacent) and labels the
compound. From now on the shorter name A can be used instead on the
compound. This is the sole role of SuperStructuring - to give a name
to a bigger compound in order to facilitate shorter explanations at latter
stages.

2. ABS - j4 — > B\C - Abstraction. Makes a name for the occurrence of either
of the variables _Bor C . This may allow for the potential bundling of two
productions, the first one involving B and the other one involving C while
otherwise being the same, into a production involving only A. The role of
Abstraction is to give a name to a group of entities (we have chosen two
for simplicity) in order to facilitate more general explanations at latter
stages which in turn will produce more compact theories.

3. RSS - AB — > C - Reverse SuperStructuring - invent up to two new internal
variables (may also use already existing ones) which if they occur within
proximity of each other together they "explain" the internal variable C.

4. RABS -A^C, S^C- Reverse Abstraction - invent new internal
variables (may also use already existing ones) such that either of them
can "explain" the internal variable C (we have chosen two variables for
simplicity) .

In the next subsection we review two possible reasons for creating hidden vari-
ables and we identify them with Reverse Abstraction and Reverse SuperStruc-
turing. Because in our GEN-ASNF we do not have other types of reasons for
creating Hidden Variables as all the other production introduce convenience re-
namings only we can infer under the Computationalistic Assumption that these
two are the essential reasons for postulating Hidden Variables and any other
reasons must be derivative. This produces a definite structural characterization
of the rationales behind inventing Hidden Variables.

Reasons for postulating Hidden Variables

There are at least two types of reasons for creating Hidden Variables:

1. (OR type) - [multiple alternative hidden causes] The OR type corresponds
to the case when some visible effect can have multiple hidden causes HI — >
Ef feet, H2 — > Effect . This case corresponds in our formalism to the
notion of Reverse Abstraction. One typical example of this is: The grass

24

Figure 1: Feynman diagram for the electron positron Annihilation into a photon
followed by photon Disintegration into an electron positron pair again

is wet, hence either it rained last night or the sprinkler was on. In the
Statistical literature the models that use these types of hidden variables
are known as mixture models |5].

2. (T-AND type) - [multiple concurrent hidden causes] The T-AND type,
i.e., topological AND type, of which the AND is a particular case. This
corresponds to the case when one visible effect has two hidden causes that
both have to occur within proximity (hence the Topological prefix) of each
other in order for the visible effect to be produced. HIH2 — >■ Effect. This
corresponds in our formalism to the case of Reverse SuperStructuring. One
example of this case is the Annihilation of an electron and a positron into
a photon. e"'"e~ — > 7 as illustrated by the following Feynman diagram. In
the diagram the Annihilation is followed by a Disintegration which in our
formalism will be represented by a Superstructure 7 — >■ e"'"e~. Note that
the electron positron annihilation is a RSS and not an ANNIHILATE2
as in the Savitch type of theorems. In a certain sense one of the argu-
ments for using RSS versus ANNIHILATE2 is also the fact that physical
processes always have a byproduct despite carrying potentially misleading
names such as annihilation, or the byproduct being energetic rather than
material. Nevertheless, since our reasons for preferring GEN-ASNF versus
GEN-ASNF-Savitch alternative are at best aesthetic / intuitive reasons it
should always be kept in mind as a possible alternative. In the Statistical
/ Graphical Models literature the particular case of AND hidden expla-
nations is the one that introduces edges between hidden variables in the
depedence graph [5], [S], [TT] .

Our analysis of the Turing equivalent formalism of Generative Grammars
written in the ASNF has evidenced them as the only needed types and we
can infer under the Computationalistic Assumption that these are the only two

25

essential reasons for postulating Hidden Variables and any other reasons must
be derivative.

Radical Positivism

Since RSS and RABS involve the postulation of hidden variables, and we have
discussed the perils associated with it, one alternative is to choose to use only
Abstraction and SuperStructuring. We propose that this in fact this character-
izes the radical positivist [T] stance which allows for empirical laws only. After
we rule out RSS and RABS since they may introduce Hidden Variables we
are left with ABS and SS and this produces our proposed characterization of
the radical positivist position and what we mean by empirical laws under the
Computationalistic Assumption. The internal variables created by Abstraction
and SuperStructuring are going to be just convenient notations for aggregates
of input data but nothing else: SuperStructuring is just convenient labeling of
already existing structures for the sake of brevity and Abstraction on the other
hand aggregates a group of variables into a more general type so that we can
produce more encompassing laws but with coarser granularity. An explanation
of a stream of observations w in the Radical Positivist theory of the world (or
least a piece of it) will look like the picture in Figure l2]- top. In this picture the
atoms are the observations and the theory of the universe is mainly a theory
of how the observables are grouped into classes (Abstractions) and how smaller
chunks of observations are tied together into bigger ones (Superstructures). The
laws of the radical positivist theory are truly empirical laws as they only address
relations among observations.

However using only these two operators we cannot attain the power of Turing
Machines. More precisely, the resulting types of grammars will be a strict subset
of Context Free Grammars, (since CFG contain SS, and REN(ABS+RABS)).
Next we will examine how any theory of the world may look like from the most
general perspective when we do allow Hidden Variables.

General theories of the world

An explanation of a stream of observations S* — > w in the more general hidden
variable theory of the world is illustrated in Figure [2] - bottom. The atoms are
the hidden variables and their organization is again in turn addressed by the
Abstraction and SuperStructuring but also Reverse Abstraction. This part is
basically a context free part since all the productions are context free. Obser-
vations are a derivative byproduct obtained from a richer hidden variable state
description by a reduction: either of size - performed by Reverse SuperStruc-
turing or of information - performed by Reverse Abstraction.

The hidden variables theory of the world picture is an oversimplification of
the true story, in general we may have a set of alternations of REN+SS and
RABS+RSS rather than just one. However as we have shown in Theorem 8
we can turn any grammar into a grammar in Strong-GEN-ASNF such that any
explanation done in three phases only (as illustrated in Figure l2| :

26

1. a growth phase where we use only REN+SS

2. a shrink phase where we use only RABS+RSS

3. a phase where we only rewrite into Terminals.

Whether additional separations can be made, e.g., if we can push all the RABS
from the first phase into the second phase and make the first phase look like the
one in the radical positivist story (i.e., using only ABS+SS) is a topic of further
investigation. We take this opportunity to proposed it as a conjecture.

Conjecture. Let G = {N,T,S,R), e ^ L{G) be a General Grammar. Then
we can convert such a grammar into the Strong-GEN-ASNF i.e., such that all
the productions are of the following form:

1. A^ B

2. A — > BC - and this is the only rule that has BC in the rhs and this is the
only rule that has A in the Ihs (strong-uniqueness) .

3. A — > a - and this is the only rule that has A on the Ihs and there is no
other rule that has a on the rhs. (strong uniqueness)

4. AB — > C - and this is the only rule that has C in the rhs and this is the
only rule that has AB in the Ihs (strong-uniqueness) .

And furthermore for any derivation w such that 7 — > w , in G, ^ £ N~^ there
exists a derivation 7 — >■ /j. — >■ z^ — > w such that ji G N^ ,v G N* and 7 — >■ /z
contains only rules of type 1, more particularly only Abstractions, and type 2
(ABS,SS), /i — > a contains only rules of the type 1, more particularly only
Reverse Abstractions, and type 4 (RABS, ESS) and v -^ w contains only rules
of type 3 (TERMINAL).

In a certain sense the conjecture can be seen as way to try to salvage as much
as we can from the Radical Positivist position by saying that in principle it is
right with the caveat that the atoms should be internal(hidden) variables rather
than observations. If we were God (here used in the scientifico-philosophical
sense - i.e., somebody which knows the laws of the universe and has access to
the full hidden state of it) then we would be able to stop our explanation of
the current state after phase 1, which would use only Abstractions and Super-
Structures (the Radical Positivist position). However since we are mere mortals
and all we are able to perceive are Observables which are just a simplified small
refiection of the current hidden state of the world there is a the need for re-
duction operations: reduction in size - performed by Reverse SuperStructuring
and reduction of information - performed by Reverse Abstraction. This is then
followed by the one to one correspondence between some internal variable and
Observables (Terminals in grammar terminlogy).

Our main claim so far is that we can rewrite any General Grammar in GEN-
ASNF, that is using only ABS, SS, RABS and RSS. In the next section we will
examine a statement that was made by the philosopher David Hume more that

27

w

Figure 2: Theory of the world: (top) - Radical Positivist, (bottom) - with
Hidden Variables

28

200 years ago in the light of the GEN-ASNF theorem and propose that they
are more or less equivalent under the Computationalistic Assumption. We then
examine the developments that occured in the meantime in order to facilitate
our proof of Hume's claim.

Hume principles of connexion among ideas

"/ do not find that any philosopher has attempted to enumerate or class all the
principles of association [of ideas]. ... To me, there appear to be only three
principles of connexion among ideas, namely. Resemblance , Contiguity in
time or place, and Cause and Effect" - David Hume, Enquiry concerning
Human Understanding, HI (19), 174-8. JSjj

If we are to substitute Resemblance for Abstraction (as resemblance/similarity
is the main criterion for abstraction) , Contiguity in time or place for SuperStruc-
turing (as the requirement for SuperStructuring is proximity - in particular spa-
tial or temporal) and Cause and Effect for the two types of hidden variable
explanations then the GEN-ASNF theorem is the proof of a more the two hun-
dred years old claim. The main developments that have facilitated this result
are:

1. The Church- Turing thesis -1936 fW (i.e., the Computationalistic Assump-
tion) which allowed us to characterize what a theory of the world is, i.e.,
an entity expressed in a Turing equivalent formalism.

2. The Generative Grammars - 1957 [4J the development of the Composi-
tional formalism of Generative Grammars which is Turing equivalent.

3. Developments in understanding the structure of Generative Grammars -
The Kuroda Normal Form 1964 [5] and General Kuroda Normal Form

m-

4. The GEN-ASNF theorems proved in this paper.

Furthermore the elucidation of the two types of causal explanation (alternative
- Reverse Abstraction (RABS) and topological conjunctive - Reverse Super-
Structuring (RSS)) is an additional merit of GEN-ASNF theorem. It should be
mentioned also that we became aware of Hume's claim after we have already
stated the GEN-ASNF theorem but prior to it's proof in the full generality. We
were led to it in a certain sense by similar intuitions and investigations into the
nature of the Structural Induction process as David Hume's. Once aware of it,
this became an additional supporting argument for the fact that the proof may
be a feasible enterprise.

29

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
