---
source: "https://archive.org/details/truthmakers_for_1st_order_sentences_-a_p"
title: "truthmakers for 1st order sentences - a proposal"
author: "Friedrich Wilhelm Grafe"
year: "2020"
captured_at: "2026-07-02T14:25:31Z"
updated_at: "2026-07-02T14:25:31Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "スピノザ"
query: "Ethica ordine geometrico demonstrata"
plain_text_url: "https://archive.org/download/truthmakers_for_1st_order_sentences_-a_p/truthmakers_for_1st_order_sentences_-a_p_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "汎神論"
  - "倫理学"
status: raw
---

# truthmakers for 1st order sentences - a proposal

- 著者: Friedrich Wilhelm Grafe
- 初版: 2020
- 情報源: [archive](https://archive.org/details/truthmakers_for_1st_order_sentences_-a_p)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[スピノザ]]
- 研究動向: [[スピノザ-現代研究動向]]

## Full Text

tmthmakers for 1st order sentences - a proposal

Friedrich Wilhelm Grafe
2020 - 07-22

Abstract

The purpose of this paper is to communicate - as a proposal - a general
method of assigning a ’truthmaker’ to any 1st order sentence in each of
its models. The respective construct is derived from the standard model
theoretic (recursive) satisfaction definition for 1st order languages and is
a conservative extension thereof.

The heuristics of the proposal (which has been somewhat idiosyncratic
from the current point of view) and some more technical detail of the
construction may be found in my article on part I of Spinoza’s ‘ethica,
ordine geometrico demonstrata’ 1 , which is the context within which I
elaborated the assignment. But this context need not be repeated here,
the presentation of the truthmaker assignment will be comprehensible to
anybody with solid basic knowledge in 1st order model theory, anything
used is standard and no advanced techniques are required.

Contents

1 introductory remarks 2

2 truthmaker assignment 3

2.1 preliminaries . 3

2.2 truthmaker assignment step by step . 3

3 truthmaking relation 5

3.1 on whether a truthmaker for a sentence A can be said in some

sense to logically imply A . 5

3.2 truthmaking relation in detail. 6

3.3 truthmaking selection apology. 6

4 first conclusions 7

4.1 concerning the correspondence theory of truth: Occam’s razor

cuts Plato’s beard ?. 7

4.2 concerning Quine’s ’’standard” for judging ontological commit¬
ment of a theory. 8

5 ’true to the facts’ is true to the facts 9

1, a formal analogy to elements of ”de deo” ’[7]

1

1 introductory remarks

It lias become popular, to frame discussions on truth (at least in the context
of a correspondence theory of truth) in terms of ‘truthbearer’ (e.g. ‘sentence’,
‘proposition’, ‘judgement’, ‘belief’, ... or the like), ‘truthmaker’ (‘facts’, ‘events’,
... or other ) and the relation between truthmaker and truthbearer, called
sometimes the ‘truthmaking relation’. 2

From the abundant literature in the held I mention, for sake of shortness, in
this context only these two: first, a Tarski-related discussion by Stephen Yablo
in chapter 4 ‘A Semantic Conception of Truthmaking’ of his book ‘aboutness’
[13] , and secondly Kit Fine’s ’Truthmaker Semantics’ [6], which contains both,
discussion and a general survey of research on ’semantic truthmakers’.

When the discussion turns on truthmakers for quantified sentences, Kit Fine
(op.cit., chapter 1 §7 Quantifiers) reflects unwelcome restrictions of considered
truthmaker assignments 3 , and Stephen Yablo, looking for a Tarski style general
solution, raises an objection to Donald Davidsons ’True to the Facts’ plea for
Tarski semantics as the better correspondence theory of truth:

”4.2 RECURSIVE TRUTHMAKERS Tarski gave us the semantic concep¬
tion of truth. Might there be room in his system also for truthmakers? David¬
son considers this question in his 1969 article “True to the Facts” (Davidson
[1969]). He believes there is not only room for truthmakers in Tarski, they are
in some sense already there. He defines truth, recall, in terms of satisfaction.
Sentences are true because of what they are true of : certain sequences of ob¬
jects. Sequences are what facts become in Tarski’s system. Satisfaction is all
that remains of correspondence. Davidson’s idea here is puzzling, for truths are
satisfied by all sequences. ...”

and concerning Davidson’s supposed way out

” ... The problem is not entirely solved, however, for distinct truths may
agree too in their derivational history. Take for instance two universal gen¬
eralizations, Vx Fx and Vx Gx, understood both to be true, and where the
predicates are atomic. Fx and Gx are both satisfied by all sequences, and there
is no more to the story than that. This is the wrong result; true generalizations
are not all true for the same reason. ...”[13] p.44 f.

The passage of Davidson’s article, which evocates Yablo’s ”... is not entirely
solved ...” criticism, reads in full

’’Since different assignments of entities to variables satisfy different open
sentences and since closed sentences are constructed from open, truth is reached,
in the semantic approach, by different routes for different sentences. All true
sentences end up in the same place, but there are different stories about how
they got there; a semantic theory of truth tells the story for a particular sentence
by running through the steps of the recursive account of satisfaction appropriate
to the sentence.” ([5], p.759 )

2 a concise survey of this terminology is contained e.g. in the article ‘correspondence theory’
of the Stanford Encyclopedia of Philosophy, §2. Truthbearers, Truthmakers, Truth[4]

3 e.g. ’’One problem with these clauses is that they presuppose a fixed domain of indi¬
viduals.For suppose that the actual individuals are ai,a 2 ,... and that a is a merely possible
individual (distinct from each of ai,a 2 ,... )• Then in a possible world in which a exists, the
truth of the instances </>(ai), <t>( a 2 ) , ... is not sufficient to guarantee the truth of V x <px and
hence the fusion of verifiers for </>(ai), <f>(a, 2 ), ... need not be a verifier for Va; <f>x (this is a
familiar problem, going back to the early days of logical atomism).” [6], p. 12

2

In what follows, it will appear, that the fact that ” ... truths are satisfied by
all sequences ... ” is not a problem at all in assigning a truthmaker to 1st order
sentences via the Tarski-type recursive satisfaction definition, though Yablo is of
course right in this insistence, that in general the truthmakers for two different
true sentences should differ 4 . But Davidson’s statement cited above, as far as it
goes, seems to me perfectly alright, and Yablo’s general request and Davidsons
statement will be seen to be compatible. Details will become obvious, as we
proceed.

2 truthmaker assignment

2.1 preliminaries

In this communication discussion is restricted to 1st order sentences, i.e., closed
formulae in 1st order language, as truth bearers, and the assignment of truth-
makers to them. The nature of the respective truthmaking relation will be clear
from this assignment. The proposed method of assignment uses and is restricted
to the set theoretic semantics of 1st order languages (theory of models) in the
tradition originating with Alfred Tarski(1935)[ll],(1954)[12].

By this methodology to any 1st order sentence (closed formula) of a 1st order
language L 5 , in any relational structure 1ZS for L, a truth value ( ’true’ or ’false’
) is assigned.This assignment of a truth value is supplied by the corresponding
recursive definition of satisfaction for any formula A of L in any relational
structure 1ZS for L. 6

2.2 truthmaker assignment step by step

The assignment of a truthmaker for each sentence (closed formula) A of a 1st
order language L in a relational structure 1ZS for L, in which A is true, runs
thus:

(I) we use the fact from 1st order model theory, that for each formula (in¬
cluding closed formulae) A of L there exists a logically equivalent prenex normal
form pnf (A) = QB , where Q is some finite sequence of universal and existen¬
tial quantifiers , and B is a quantifier free formula; B is called the matrix of
prenex normal form QB .

4 whereas in the special case mentioned [Vx Fx and Vx Gx] , in relational structures, in
which both sentences are true, Fx and Gx need have the same extension and thus presum¬
ably Vx Fx and Vx Gx will have the same truthmaker in the framework of any extensional
semantics

5 1st order languages L, as we presuppose here, may or may not, in addition to n-ary
predicate letters, contain n-ary function symbols and/or identity.

6 For technical details especially concerning the recursive satisfation definition I refer
in the main to Kreisel-Krivine(1967)[9] chapter 2 ‘Predicate Calculus’, for the use of the
term ’relational structure’ instead of the Kreisel-Krivine term ’realization’ see e.g. Bell &
Slomson(1974)[l]. The reason for my preference of Kreisel-Krivine is, that in this text the
recursive definition of satisfaction is remarkably concise formulated; but in principle any other
serious introductory textbook of model theory will do as well as a reference.

3

(II) the set of all formulae of L valid in 1ZS , including the set of all closed
formulae(sentences) of L true in 1ZS , is called the theory of 1ZS , in symbols
TH (HS) As we proceed, we define ‘truthmaker’ for all sentences A e TH (HS)

(III) The recursive satisfaction definition assigns to any formula A of L in
a relational structure 1ZS for L a value A, viz. a set of variable assignments
(a variable assignment being a map, assigning to each individual variable of L
an element of the universe of discourse U of the relational structure 1ZS). And
thus is assigned the value B to the matrix B of pnf (A) = QB.

Now we restrict (without loss of generality) the value B of B in 1ZS, by
restricting any map in B to the individual variables occurring (free) in B, and
get the (finite or infinite) set /B/ as a truthmaker for QB. We may view /B/ as
(finite or infinite) set of finite sequences of objects from U, the domain (’universe
of discourse’ ) of 1ZS (, each object in each sequence ‘indexed’ by the variable
it is assigned to). In what follows, it is of utmost importance, to keep in mind,
that this truthmaker assignment is relative to the relational structure 1ZS.

As, due to logical equivalence transformations, the prenex normal form of
a given formula need not be unique, and as not every sentence A e TH(1ZS)
is in prenex normal form, the truthmaker definition extends the truthmaker
assignment for sentences AcTH (HS) hr prenex normal form to all sentences
A e TH ( US ) by

(IV) If /B/ is a truthmaker for 1st order sentence A in the relational struc¬
ture 1ZS, /B/ is also a truthmaker in 1ZS for any 1st order sentence A' with
II- A ^ A'

The Definition is summarized by:

(V) A truthmaker /B/ for 1st order sentence A in the relational structure 1ZS
is either constructed according to (I)-(III) or inherited via logical equivalence
by (IV)

Corollary

In each relational structure 1ZS the set of all truthmakers of a sentence
(closed formula) A true in 1ZS is the smallest set, that contains the truthmaker
of some prenex normal form pnf (A) = QB and in addition the truthmakers of
all sentences A' 6 TH (HS) with II- A <-> A'.

’having the same set of truthmakers with respect to relational structure 1ZS'
is of course an equivalence relation, and thus splits the set of all sentences
A e TH(1ZS) into equivalence classes, but the respective equivalence relation
is not simply induced by logical equivalence. Again the dependence of the
truthmakers on 1ZS plays a crucial role: Simple example is Yablos A x Fx. In
relational structures, in which A x Fx is true, V x Fx is also true and has the
same truthmaker(s).

4

3 truthmaking relation

The dependency of such truthmakers on the respective relational structure was
already tacitly used in my note on Yablo’s supposed counterexample to David¬
sons statement. Really one get’s into a logical mine held, if the dependence of
the truthmaker assignment on the respective relational structure is ignored or
violated; let’s have a look on the following ‘trap’:

3.1 on whether a truthmaker for a sentence A can be said
in some sense to logically imply A

Let A = QB a A e TTi (TZS) with truthmaker /B/ :

We adopt proper names for all the objects from the universe of discourse
contained in /B/ , and from the matrix formula B by substituting proper names
corresponding to truthmaker sequences {</> i} for the variables in B we get a
finite or infinite set {B su b</>i} of quantiherfree sentences true in the relational
structure. Now the trap is the conjecture, that this set of ‘matrix sentences’
does logically imply our 1st order sentence A, in symbols

(_?_) {B sub (/>i} lh A

This were of course a complete misunderstanding of the situation. The truth
of the sentences in {B su b0i} relates to the current relational structure 1ZS , and
the truth of the sentences in {B su b^i} necessitates the truth of A in the general
case only with respect to 1ZS. All this is far from logical implication, as logical
implication is the situation that in any relational structure, in which all the
premises are true, the conclusion is also true.

And there is another argument to the effect, that the conjecture ( ? ) is
not valid, which, while it does not cover the general case, is still instructive in
another way. Consider the following counterexample:

Suppose, our relational structure were the set IN of natural numbers, with
the successor-function s: IN »-»• IN , and A be the sentence stating that 0 is not
successor of any number and that the successor function is injective (1-1), this
sentence already in prenex normal form. The set of all true ‘matrix sentences’

{Bgub^i} ia this intended model is of course infinite.

Now, by indirect proof: suppose

{b s (x) = 0 a [s (x) = S (y) X = y]]sub<M IH A x A y bs (x) = 0 a [s (x) = s (y) ->■ x = y]]

then, by compactness,

A x A y hs (x) = 0 a [s (x) = s (y) -» x = y]]

need follow already from a finite subset of {[— .s (x) = 0 A [s (x) = s (y) -»■ x = y]] su tbi}>

5

which is obviously false, because any such finite subset has a model in a
finite domain, viz. here in an initial segment of IN, whereas A clearly has not.'

3.2 truthmaking relation in detail

We conclude from this concerning the

(VI) truthmaking relation:

the truth of A is necessitated by its truthnraker /B/, but this necessitating
relation is given by the recursive satisfation definition working on the current re¬
lational structure, and not by logical implication from a set of matrix-instances. 8

At this stage one might perhaps be inclined to say something like

”ok, the truthnraker is the relational structure, the truthmaking relation
between the sentences of the 1st order language and the relational structure
given by the recursive satisfaction definition, so what? I knew that already ... ”

The reply is of course, that the assignment (with respect to the relational
structure) of a uniquely determined set of truthnrakers to any set of logically
equivalent sentences true in this relational structure, is far more specific.

As well much more specific is the truthmaking relation:

For sentences A already in prenex normal form A = QB its simply the work¬
ing of the satisfaction definition on the quantifier prefix Q, taking as a starting
point the value B of the matrix B 9 , 10 .

In case of sentences A not in prenex normal form, the truthmaking relation
is given by (IV), i.e. inherited from the above case.

In section 4 ’first conclusions’, I’ll try to show, that this is not only a for¬
mal exercise, but gives philosophically relevant structural information on ’’the
concept of truth in the formalized languages”.

3.3 truthmaking selection apology

Another objection might be raised with respect to my selecting the matrix of
prenex normal form as the anchor point of the truthmaker definition. Ok, in
some sense, for every term frequently used but without sharp contours of usage,

7 and generally, in any case, where the considered sentence A is satisfiable only in an infinite
domain, this reasoning by compactness applies

8 both counter arguments stated above do not only serve as refutations of (-?-), but relate
- in some way and to some extent - as well to Kit Fine’s dealing with quantification within
the ’state space’ semantics sketched in [6], Part I, §§ 3-4 and 7.

"using e.g. projection along x; for any existential quantifier V Xi i and complement for any
negation sign - for details see e.g. Kreisel-Krivine(1967) [9],chapter 2, p.17

10 Thus, the set of matrix sentences {B sl ,i,rq} may be generated from the truthmaker B
and the matrix formula B, but the working input in ‘truthmaking’ is not this set of matrix
sentences, but the set of variable assignments satifying the matrix in the current relational
structure.

6

the selection of a definiens has a shadow of arbitraryness. In this case, in order
to define truthmakers for 1st order sentences, this arbitrariness to me seems
little if not marginal. The selection of the matrix of the prenex normal form
has two indisputable advantages. First, that it serves our ubiquitous intuitions
in the use of matrix sentences in order to cope with quantification in defining
truthmakers, and let’s us understand, why and in what respect, these intuitions
fail. Secondly, by the metatheorem on the existence of a logically equivalent
prenex normal form for any formula A e L, we have the means to spread the
truthmakers derived from sentences in prenex normal form within the same
relational structure 1ZS to any sentence 1ZS II- A. Thus, if someone still insisted,
that this selection is arbitrary, I would concede, that it is, but, in my view, much
less so than anything else, I met in the field.

4 first conclusions

4.1 concerning the correspondence theory of truth: Oc¬
cam’s razor cuts Plato’s beard ?

In an obvious sense, the proposed truthmaker construction is in accordance with
Tarski type recursive satisfaction definition, first, in that it renounces the in¬
troduction of additional abstract entities duplicating the linguistic entities(e.g.
propositions assigned to sentences, properties or the like assigned to unary pred¬
icate letters, ... ), and secondly, in that it does not support ‘falsemakers’ (un¬
willing to serve Plato’s beard 11 ).

Concerning the first point, not so much can be said in favour of ‘intensions’
and the like in the current context. The model theoretic recursive satisfaction
definition proceeds with reference to extensions (set theoretic entities, maybe
including ’atoms’ 12 ) only and does so successfully. The whole spectrum of dis¬
cussion concerning say ’sense’ as opposed to ’reference’, intension, propositions,
properties, tropes, ... etc. is seemingly bypassed here without any loss with
respect to truth value distribution or to the proposed truthmaker definition.

Concerning the second point: An adequacy condition on any theory of truth
is, that besides providing an account of ’truth’, to provide also an account of
’falsehood’. And the discussion concerning Plato’s beard concerns the question
of whether the provided account of falsehood will be in some sense symmetric
or asymmetric to the given account of truth. 13

Seemingly innocent but nevertheless a trap concerning ‘the riddle of nonbe¬
ing’ is:

11 for Quine’s ’On what there is’ epitheton ’’Plato’s beard” ([10], pp.1-2) see e.g. Sophie-
Grace Chappell, Plato on Knowledge in the Theaetetus, 7.1 The Puzzle of Misidentification:
187e5-188c8 [3]

12 ’atoms’ in set theory are members of sets, which are not itself sets; see e.g. Ketland[8]

13 in the beginnings of the 20th century, Bertrand Russell and G.E. Moore switched a number
of times between different symmetric and asymmetric variants

7

... the fact, that exists, if p is true = the fact, that not exists, if p is false ...

[ the problem here, of course, is the right hand side of the equation ].

Whatever in this respect, it is clear, that the truthmaker proposal, as here
presented, does not supply a ’falsemaker’ for false sentences in a symmetric man¬
ner. This might perhaps be misunderstood. Of course, a sentence A false in the
relational structure 1ZS, is as well assigned the value of the matrix of any of its
prenex normal forms, and such a value may be considered as a falsemaker. The
asymmetry then lies in this: consider the respective matrix sentences, generated
from this kind of ’falsemaker’. What they are about, is, how facts are in this
relational structure 1ZS, and these matrix sentences do not refer to something
would-be, if A were true.

This asymmetry is of course grounded in the fact, that the (recursive) gen¬
eration process for the value of a formula A of L in 1ZS starts from the values of
those atomic formulae, which are part of A, each value of each atomic formula
corresponding to a set of ’positive’ sentences true in 1ZS . The recursive satis¬
faction definition assigns to 1st order sentences the truth value ‘true’ or ‘false’,
according to the facts in the relational structure, but while it does not supply
a symmetric falsemaker for a 1st order sentence A false in this relational struc¬
ture, it does of course supply a truthmaker for ->A in any relational structure,
in which -iA is true.

4.2 concerning Quine’s ’’standard” forjudging ontological
commitment of a theory

Lastly, let’s consider Willard van Orman Quine’s often cited ”To be ..., is ... to
be ...the value of a variable” 14 , more explicitly ”...: a theory is committed to
those and only those entities, to which the bound variables of the theory must
be capable of referring in order that the affirmations made in the theory be
true” 15

A modification of Quine’s standard might be suggested by the above truth-
maker definition, tentatively: the ontology, which a say finitely axiomatized 1st
order theory T is committed to, with respect to its intended model TZStj is given
not by the (complete) domain of the intended model but by the truthmakers of
its theorems (special sets of sequences of objects over that domain). Thus the
suggestion is, to take as ontologically relevant for e.g. a theory in physics neither
say real numbers in isolation nor Active physical objects (e.g. ‘masspoints’), but
only complex arrangements of them as the ’’objects, the theory ‘deals with’, or
is about’ ”. But this is a great topic of its own, it is reflected in a suitable frame
in Jeffrey Ketland’s paper ’Foundations of Applied Mathematics I’ under the
topic of ’mixed mathematical objects’ 16 .

14 ’On what there is’[10], cf. pp. 13-18

15 ibid., while Quine immediately relativizes the importance of this ’’semantical formula”
with respect to conventionalist or phenomenalist preferences of theory selection, he keeps the
’’formula” still running as valid standard.

16 [8], sections 1.2-1.3, pp.3-7

5 ’true to the facts’ is true to the facts

In my introductory remarks I stated, that the passage of Davidson’s ’True to
the facts’, quoted by Stephen Yablo, up to my understanding, seems perfectly
alright. This passage is taken from a context, in which Davidson compares
Tarki’s account of truth of sentences (via the recursive satisfaction definition)
with versions of the ‘correspondence theory of truth’, which rely on the relation
of true sentences to the facts, they express.

The summary of this comparison in Davidson’s article follows shortly after
the already cited passage. And now, having stated my truthmaker proposal,
and argued to some extent for it, I would like to conclude this paper by quoting
Davidson again with this summary:

’’Seen in retrospect, the failure of correspondence theories of truth based on
the notion of fact traces back to a common source: the desire to include in the
entity to which a true sentence corresponds not only the objects the sentence is
’’about” (another idea full of trouble) but also whatever it is the sentence says
about them. One well-explored consequence is that it becomes difficult to de¬
scribe the fact that verifies a sentence except by using that sentence itself. The
other consequence is that the relation of correspondence (or ’’picturing”) seems
to have direct application to only the simplest sentences (’Dolores loves Dag-
mar’). This prompts fact-theorists to try to explain the truth of all sentences in
terms of the truth of the simplest and hence in particular to interpret quantifi¬
cation as mere shorthand for conjunctions or alternations (perhaps infinite in
length) of the simplest sentences. The irony is that, insofar as we can see quan¬
tification in this light, there is no real need for anything like correspondence.
It is only when we are forced to take generality as an essential addition to the
conceptual resources of predication and the compounding of sentences, and not
reducible to them, that we appreciate the uses of a sophisticated correspondence
theory” [5], p.579

References

[1] John L. Bell, Alan B. Slomson. Models and Ultraproducts , North Hol¬
land Publishing, Amsterdam - Oxford, 1969, 3 1974.

[2] hrsg. Karel Berka, Lotliar Kreiser. Logik-Texte - Kommentierte
Auswahl zur Geschichte der modernen Logik. Akademie-Verlag,
Berlin, 1971.

[3] Sophie-Grace Chappell. Plato on Knowledge in the Theaetetus - in
The Stanford Encyclopedia of Philosophy, Winter 2019 Edition,

open access at https://plato.stanford.edu/archives/win2019/entries/plato-
theaetetus/, 2016.

[4] Marian David. The Correspondence Theory of Truth - in The Stan¬
ford Encyclopedia of Philosophy, Metaphysics Research Lab - Stanford
University,

open access at https://plato.stanford.edu/archives/fall2016/entries/truth-
correspondence/, 2016.

9

[5] Donald Davidson. True to the Facts , The Journal of Philosophy , vol.
66(1969), pp. 748-764 .

[6] Kit Fine. Truthmaker Semantics (Chapter for the Black-
well Philosophy of Language Handbook), open access at
https://www.researchgate.net/publication/313824698_Truthmaker_Semantics,
2017.

[7] Friedrich Wilhelm Grafe.

a formal analogy to elements of ’de deo ’, open access at
https://wilhehngrafe.academia.edu/research/, academia.edu, 2020.

[8] Jeffrey Ketland. Foundations of Applied Mathematics 7, Metaphysics
Research Lab - Stanford University,

draft paper, open access at

https: / / www.academia.edu/42107610/Foundations_of_Applied_Mathematics_I,
2020 .

[9] Georg Kreisel, Jean-Louis Krivine. Elements of Mathematical Logic
(model theory), North Holland Publishing, Amsterdam, 1967.

[10] Willard Van Orman Quine. On what there is, From a logical point of
view (©1953, 1961 by the President and Fellows of Harvard College),
Harper and Row, New York and Evanston, 3 1963, pp.1-19.

[11] Alfred Tarski. Der Wahrheitsbegrijf in den formalisierten Sprachen, Studia
Philosophia - Commentarii Societatis philosphicae Polonorum,
vol. 1 (1935), pp. 261-405, reprint in [2].

[12] Alfred Tarski. Contributions to the Theory of Models I,II,III, Inda-
gationes Mathematicae - Nederlandse Akademie van Weten-
schapen, vol. 16(1954), pp. 572-588, vol. 17(1955), pp. 56-64 .

[13] Stephen Yablo. Aboutness, Princeton University Press, Princeton - Ox¬
ford, 2014.

10

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。
