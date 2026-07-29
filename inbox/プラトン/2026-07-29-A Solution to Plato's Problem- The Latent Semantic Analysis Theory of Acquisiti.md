---
source: "https://www.semanticscholar.org/paper/68dd4b89ce1407372a29d05ca9e4e1a2e0513617"
title: "A Solution to Plato's Problem: The Latent Semantic Analysis Theory of Acquisition, Induction, and Representation of Knowledge."
author: "T. Landauer, S. Dumais"
year: "1997"
publication: ""
download: "http://www.stat.cmu.edu/%7Ecshalizi/350/2008/readings/Landauer-Dumais.pdf"
pdf: "http://www.stat.cmu.edu/%7Ecshalizi/350/2008/readings/Landauer-Dumais.pdf"
captured_at: "2026-07-29T08:55:44Z"
updated_at: "2026-07-29T08:55:44Z"
capture_tool: "scrapem"
source_name: "semanticscholar"
keyword: "プラトン"
query: "Plato"
tags:
  - "古代哲学"
  - "イデア論"
  - "倫理学"
status: raw
---

# A Solution to Plato's Problem: The Latent Semantic Analysis Theory of Acquisition, Induction, and Representation of Knowledge.

- 著者: T. Landauer, S. Dumais
- 年: 1997
- 情報源: [semanticscholar](https://www.semanticscholar.org/paper/68dd4b89ce1407372a29d05ca9e4e1a2e0513617)
- ダウンロード: http://www.stat.cmu.edu/%7Ecshalizi/350/2008/readings/Landauer-Dumais.pdf
- PDF: http://www.stat.cmu.edu/%7Ecshalizi/350/2008/readings/Landauer-Dumais.pdf

## Obsidian Links

- 研究動向: [[プラトン-現代研究動向]]
- キーワード: [[プラトン]]
- 関連分野: [[古代哲学]]
- 関連分野: [[イデア論]]
- 関連分野: [[倫理学]]
- 関連タグ: #古代哲学 #イデア論 #倫理学

## Abstract

取得したページに要旨は含まれていない。

## Citation

DOI: 10.1037/0033-295X.104.2.211

## PDF Text

Copyright 1997 by the American Psychological Association, Inc.
0033-295X/97/J3.00

Psychological Review
1997. Vol. 1M. No. 2, 211-240

A Solution to Plato's Problem: The Latent Semantic Analysis Theory of Acquisition, Induction, and Representation of Knowledge
Susan T. Dutnais

Thomas K Landauer
University of Colorado at Boulder

Bellcore

How do people know as much as they do with as little information as they get? The problem takes many forms; learning vocabulary from text is an especially dramatic and convenient case for research.
A new general theory of acquired similarity and knowledge representation, latent semantic analysis
(LSA), is presented and used to successfully simulate such learning and several other psycholinguistic phenomena. By inducing global knowledge indirectly from local co-occurrence data in a large body of representative text, LSA acquired knowledge about the full vocabulary of English at a comparable rate to schoolchildren. LSA uses no prior linguistic or perceptual similarity knowledge; it is based solely on a general mathematical learning method that achieves powerful inductive effects by extracting the right number of dimensions (e.g., 300) to represent objects and contexts. Relations to other theories, phenomena, and problems are sketched.

24 centuries ago, the fact that people have much more knowledge than appears to be present in the information to which they have been exposed. Plato's solution, of course, was that people must come equipped with most of their knowledge and need only hints and contemplation to complete it.
In this article we suggest a very different hypothesis to explain the mystery of excessive learning. It rests on the simple notion that some domains of knowledge contain vast numbers of weak interrelations that, if properly exploited, can greatly amplify learning by a process of inference. We have discovered that a very simple mechanism of induction, the choice of the correct dimensionality in which to represent similarity between objects and events, can sometimes, in particular in learning about the similarity of the meanings of words, produce sufficient enhancement of knowledge to bridge the gap between the information available in local contiguity and what people know after large amounts of experience.

Prologue
"How much do we know at any time? Much more, or so I believe, than we know we know!"
—Agatha Christie, The Moving Finger

A typical American seventh grader knows the meaning of
10-15 words today that she did not know yesterday. She must have acquired most of them as a result of reading because (a)
the majority of English words are used only in print, (b) she already knew well almost all the words she would have encountered in speech, and (c) she learned less than one word by direct instruction. Studies of children reading grade-school text find that about one word in every 20 paragraphs goes from wrong to right on a vocabulary test. The typical seventh grader would have read less than 50 paragraphs since yesterday, from which she should have learned less than three new words. Apparently, she mastered the meanings of many words that she did not encounter. Evidence for all these assertions is given in detail later.
This phenomenon offers an ideal case in which to study a problem that has plagued philosophy and science since Plato

Overview
In this article we report the results of using latent semantic analysis (LSA), a high-dimensional linear associative model that embodies no human knowledge beyond its general learning mechanism, to analyze a large corpus of natural text and generate a representation that captures the similarity of words and text passages. The model's resulting knowledge was tested with a standard multiple-choice synonym test, and its learning power compared to the rate at which school-aged children improve their performance on similar tests as a result of reading. The model's improvement per paragraph of encountered text approximated the natural rate for schoolchildren, and most of its acquired knowledge was attributable to indirect inference rather than direct co-occurrence relations. This result can be interpreted in at least two ways. The more conservative interpretation is that it shows that with the right analysis a substantial portion of the information needed to answer common vocabulary test questions can be inferred from the contextual statistics of usage alone. This is not a trivial conclusion. As we alluded to earlier

Thomas K Landauer, Institute of Cognitive Science, University of
Colorado at Boulder; Susan T. Dumais, Information Science Research
Department, Bellcore, Morristown, New Jersey.
We thank Karen Lochbaum for valuable help in analysis; George
Furnas for early ideas and inspiration; Peter Foltz, Walter Kintsch, and
Ernie Mross for unpublished data; and for helpful comments on the ideas and drafts, we thank, in alphabetic order, Richard Anderson, Doug
Carroll, Peter Fbltz, George Pumas, Walter Kintsch, Lise Menn, and
Lynn Streeter.
Correspondence concerning this article should be addressed to
Thomas K Landauer, Campus Box 345, University of Colorado, Boulder,
Colorado 80309. Electronic mail may be sent via Internet to landauer
@psych.colorado.edu.

211

212

LANDAUER AND DUMAIS

and elaborate later, much theory in philosophy, linguistics, artificial intelligence research, and psychology has supposed that acquiring human knowledge, especially knowledge of language, requires more specialized primitive structures and processes, ones that presume the prior existence of special foundational knowledge rather than just a general purpose analytic device.
This result questions the scope and necessity of such assumptions. Moreover, no previous model has been applied to simulate the acquisition of any large body of knowledge from the same kind of experience used by a human learner.
The other, more radical, interpretation of this result takes the mechanism of the model seriously as a possible theory about all human knowledge acquisition, as a homologue of an important underlying mechanism of human cognition in general. In particular, the model employs a means of induction—dimension optimization—that greatly amplifies its learning ability, allowing it to correctly infer indirect similarity relations only implicit in the temporal correlations of experience. The model exhibits humanlike generalization that is based on learning and that does not rely on primitive perceptual or conceptual relations or representations. Similar induction processes are inherent in the mechanisms of certain other theories (e.g., some associative, semantic, and neural network models). However, as we show later, substantial effects arise only if the body of knowledge to be learned contains appropriate structure and only when a sufficient—possibly quite large—quantity of it has been learned.
As a result, the posited induction mechanism has not previously been credited with the significance it deserves or exploited to explain the many poorly understood psychological phenomena to which it may be germane. The mechanism lends itself, among other things, to a deep reformulation of associational learning theory that appears to offer explanations and modeling directions for a wide variety of cognitive phenomena. One set of phenomena that we discuss later in detail, along with some auxiliary data and simulation results, is contextual disambiguation of words and passages in text comprehension.
Because readers with different theoretical interests may find these two interpretations differentially attractive, we have followed a slightly unorthodox manner of exposition. Although we later present a general theory, or at least the outline of one, that incorporates and fleshes out the implications of the inductive mechanism of the formal model, we have tried to keep this development somewhat independent of the report of our simulation studies. That is, we eschew the conventional stance that the theory is primary and the simulation studies are tests of it.
Indeed, the historical fact is that the mathematical text analysis technique came first, as a practical expedient for automatic information retrieval, the vocabulary acquisition simulations came next, and the theory arose last, as a result of observed empirical successes and discovery of the unsuspectedly important effects of the model's implicit inferential operations.

The Problem of Induction
One of the deepest, most persistent mysteries of cognition is how people acquire as much knowledge as they do on the basis of as little information as they get. Sometimes called "Plato's problem'' o r ' 'the poverty of the stimulus,'' the question is how observing a relatively small set of events results in beliefs that

are usually correct or behaviors that are usually adaptive in a large, potentially infinite variety of situations. Following Plato, philosophers (e.g., Goodman, 1972; Quine, 1960), psychologists (e.g., Shepard, 1987; Vygotsky, 1968), linguists (e.g.,
Chomsky, 1991; Jackendoff, 1992; Pinker, 1990), computation scientists (e.g., Angluin & Smith, 1983; Michaelski, 1983) and combinations thereof (Holland, Holyoak, Nisbett, & Thagard,
1986) have wrestled with the problem in many guises. Quine
(1960), following a tortured history of philosophical analysis of scientific truth, has called the problem ' 'the scandal of induction," essentially concluding that purely experience-based objective truth cannot exist. Shepard (1987) has placed the problem at the heart of psychology, maintaining that a general theory of generalization and similarity is as necessary to psychology as
Newton's laws are to physics. Perhaps the most well-advertised examples of the mystery lie in the acquisition of language.
Chomsky (e.g., Chomsky, 1991) and followers assert that a child's exposure to adult language provides inadequate evidence from which to learn either grammar or lexicon. Gold, Osherson,
Feldman, and others (see Osherson, Weinstein, & Stob, 1986)
have formalized this argument, showing mathematically that certain kinds of languages cannot be learned to certain criteria on the basis of finite data. The puzzle presents itself with quantitative clarity in the learning of vocabulary during the school years, the particular case that we address most fully in this article. Schoolchildren learn to understand words at a rate that appears grossly inconsistent with the information about each word provided by the individual language samples to which they are exposed and much faster than they can be made to by explicit tuition.
Recently Pinker (1994) has summarized the broad spectrum of evidence on the origins of language—in evolution, history, anatomy, physiology, and development. In accord with Chomsky's dictum, he concludes that language learning must be based on a very strong and specific innate foundation, a set of general rules and predilections that need parameter setting and filling in, but not acquisition as such, from experience. Although this
"language instinct" position is debatable as stated, it rests on an idea that is surely correct, that some powerful mechanism exists in the minds of children that can use the finite information they receive to turn them into competent users of human language. What we want to know, of course, is what this mechanism is, what it does, how it works. Unfortunately the rest of the instinctivist answers are as yet of limited help. The fact that the mechanism is given by biology or that it exists as an autonomous mental or physical "module" (if it does), tells us next to nothing about how the mind solves the basic inductive problem.
Shepard's (1987) answer to the induction problem in stimulus generalization is equally dependent on biological givens, but offers a more precise description of some parts of the proposed mechanism. He has posited that the nervous system has evolved general functional relations between monotone transductions of perceptual values and the similarity of central interpretive processes. On average, he has maintained, the similarities generated by these functions are adaptive because they predict in what situations—consequential regions in his terminology—the same behavioral cause-effect relations are likely to hold. Shepard's mathematical laws for stimulus generalization are empiriTHE LATENT SEMANTIC ANALYSIS THEORY OF KNOWLEDGE

cally correct or nearly so for a considerable range of low-dimensional perceptual continua and for certain functions computed on behaviorally measured relations such as choices between stimuli or judgments of similarity or inequality on some experiential dimension. However, his laws fall short of being able to predict whether cheetahs are considered more similar to zebras or tigers, whether friendship is thought to be more similar to love or hate, and are mute, or at least very incomplete, on the similarity of the meanings of the words cheetah, zebra, tiger, love, hate, andpode. Indeed, it is the generation of psychological similarity relations based solely on experience and the achievement of bridging inferences from experience about cheetahs and friendship to behavior about tigers and love and from hearing conversations about one to knowledge about the other that pose the most difficult and tantalizing puzzle.
Often the cognitive aspect of the induction puzzle is cast as the problem of categorization, of finding a mechanism by which a set of stimuli, words, or concepts (cheetahs, tigers) come to be treated as the same for some purposes (running away from, or using metaphorically to describe a friend or enemy). The most common attacks on this problem invoke similarity as the underlying relation among stimuli, concepts, or features (e.g.,
Rosch, 1978; Smith & Medin, 1981; Vygotsky, 1968). But as
Goodman (1972) has trenchantly remarked, "similarity is an impostor," at least for the solution of the fundamental problem of induction. For example, the categorical status of a concept is often assumed to be determined by similarity to a prototype, or to some set of exemplars (e.g., Rosch, 1978; Smith & Medin,
1981). Similarity is either taken as primitive (e.g., Posner &
Keele, 1968; Rosch, 1978) or as dependent on shared component features (e.g.. Smith & Medin, 1981; Tversky, 1977; Tversky &
Gati, 1978). But this throws us into an unpleasant regress:
When is a feature a feature? Do bats have wings? When is a wing a wing? Apparently, the concept wing is also a category dependent on the similarity of features. Presumably, the regress ends when it grounds out in the primitive perceptual relations assumed, for example, by Shepard's theory. But only some basic perceptual similarities are relevant to any feature or category, others are not; a wing can be almost any color. The combining of disparate things into a common feature identity or into a common category must very often depend on experience. How does that work? Crisp categories, logically defined on rules about feature combinations, such as those often used in category learning, probability estimation, choice and judgment experiments, lend themselves to acquisition by logical rule-induction processes, although whether such processes are what humans always or usually use is questionable (Holland, Holyoak, Nisbett, & Thagard, 1986; Medin, Goldstone, & Centner, 1993;
Murphy & Medin, 1985; Smith & Medin, 1981). Sorely, the natural acquisition of fuzzy or probabilistic features or categories relies on some other underlying process, some mechanism by which experience with examples can lead to treating new instances more or less equivalently, some mechanism by which common significance, common fate, or common context of encounter can generate acquired similarity. We seek a mechanism by which the experienced and functional similarity of concepts—especially complex, largely arbitrary ones, such as the meaning of concept, component, or feature, or, perhaps, the component features of which concepts might consist—are cre-

213

ated from an interaction of experience with the logical (or mathematical or neural) machinery of mind.
In attempting to explain the astonishing rate of vocabulary learning—some 7-10 words per day—in children during the early years of preliterate language growth, theorists such as
Carey (1985), Clark (1987), Keil (1989), and Markman
(1994) have hypothesized constraints on the assignment of meanings to words. For example it has been proposed that early learners assume that most words are names for perceptually coherent objects, that any two words usually have two distinct meanings, that words containing common sounds have related meanings, that an unknown speech sound probably refers to something for which the child does not yet have a word, and that children obey certain strictures on the structure of relations among concept classes. Some theorists have supposed that the proposed constraints are biological givens, some have supposed that they derive from progressive logical derivation during development, some have allowed that constraints may have prior bases in experience. Many have hedged on the issue of origins, which is probably not a bad thing, given our state of knowledge.
For the most part, proposed constraints on lexicon learning have also been described in qualitative mentalistic terminology that fails to provide entirely satisfying causal explanations; Exactly how, for example does a child apply the idea that a new word has a new meaning?
What all modern theories of knowledge acquisition (as well as Plato's) have in common is the postulation of constraints that greatly (in fact, infinitely) narrow the solution space of the problem that is to be solved by induction, that is, by learning.
This is the obvious, indeed the only, escape from the inductive paradox. The fundamental notion is to replace an intractably large or infinite set of possible solutions with a problem that is soluble on the data available. So, for example, if biology specifies a function on wavelength of light that is assumed to map the difference between two objects that differ only in color onto the probability that doing the same thing with them will have the same consequences, then a bear need sample only one color of a certain type of berry before knowing which others to pick.
There are several problematical aspects to constraint-based resolutions of the induction paradox. One is whether a particular constraint exists as supposed. For example, is it true that young children assume that the same object is given only one name, and if so is the assumption correct about the language to which they are exposed? (It is not in adult English usage; ask 100
people what to title a recipe or name a computer command, and you will get almost 30 different answers on average—see Furnas, Landauer, Gomez, & Dumais, 1983, 1987). These are empirical questions, and ones to which most of the research in early lexical acquisition has been addressed. One can also wonder about the origin of a particular constraint and whether it is plausible to regard it as a primitive process with an evolutionary basis. For example, most of the constraints proposed for language learning are very specific and relevant only to human language, making their postulation consistent with a very strong instinctive and modular view of mental processes.
The existence and origin of particular constraints is only one part of the problem. The existence of some set of constraints is a logical necessity, so that showing that some exist is good but not nearly enough. We also need to know whether a particular

214

LANDAUER AND DUMA1S

set of constraints is logically and pragmatically sufficient, that is, whether the problem space remaining after applying them is soluble. For example, suppose that young children do, in fact, assume that there are no synonyms. How much could that help them in learning the lexicon from the language to which they are exposed? Enough? Indeed, that particular constraint leaves the mapping problem potentially infinite; it could even exacerbate the problem by tempting the child to assign too much or the wrong difference to our dog, the collie, and Fido. Add in the rest of the constraints that have been proposed: Enough now?
How can one determine whether a specified combination of constraints would solve the problem, or perhaps better, determine how much of the problem it would solve? We believe that the best available strategy is to specify a concrete computational model embodying the proposed constraints and to simulate as realistically as possible its application to the acquisition of some measurable and interesting properties of human knowledge. In particular, with respect to constraints supposed to allow the learning of language and other large bodies of complexly structured knowledge, domains in which there are very many facts each weakly related to very many others, effective simulation may require data sets of the same size and content as those encountered by human learners. formally, that is because weak local constraints can combine to produce strong inductive effects in aggregate. A simple analog is the familiar example of a diagonal brace to produce rigidity in a structure made of three beams. Each connection between three beams can be a single bolt. Two such connections exert no constraint at all on the angle between the beams. However, when all three beams are so connected, all three angles are completely specified. In structures consisting of thousands of elements weakly connected
(i.e., constrained) in hundreds of different ways (i.e., in hundreds of dimensions instead of two), the effects of constraints may emerge only in very large, naturally generated ensembles.
In other words, experiments with miniature or concocted subsets of language experience may not be sufficient to reveal or assess the forces that hold conceptual knowledge together. The relevant quantitative effects of such phenomena may only be ascertainable from experiments or simulations based on the same masses of input data encountered by people.
Moreover, even if a model could solve the same difficult problem that a human does given the same data it would not prove that the model solves the problem in the same way. What to do? Apparently, one necessary test is to require a conjunction of both kinds of evidence—observational or experimental evidence, that learners are exposed to and influenced by a certain set of constraints, and evidence that the same constraints approximate natural human learning and performance when embedded in a simulation model running over a natural body of data.
However, in the case of effective but locally weak constraints, the first part of this two-pronged test—experimental or observational demonstration of their human use—might well fail. Such constraints might not be detectable by isolating experiments or in small samples of behavior. Thus, although an experiment or series of observational studies could prove that a particular constraint is used by people, it could not prove that it is not. A
useful strategy for such a situation is to look for additional effects predicted by the postulated constraint system in other

phenomena exhibited by learners after exposure to large amounts of data.

The Latent Semantic Analysis Model
The model we have used for simulation is a purely mathematical analysis technique. However, we want to interpret the model in a broader and more psychological manner. In doing so, we hope to show that the fundamental features of the theory that we later describe are plausible, to reduce the otherwise magical appearance of its performance, and to suggest a variety of relations to psychological phenomena other than the ones to which we have as yet applied it.
We explicate all of this in a somewhat spiral fashion. First, we try to explain the underlying inductive mechanism of dimensionality optimization upon which the model's power hinges.
We then sketch how the model's mathematical machinery operates and how it has been applied to data and prediction. Next, we offer a psychological process interpretation of the model that shows how it maps onto but goes beyond familiar theoretical ideas, empirical principles, findings, and conjectures. We finally return to a more detailed and rigorous presentation of the model and its applications.

An Informal Explanation of the Inductive Value of Dimensionality Optimization
Suppose that Jack and Jill can only communicate by telephone. Jack, sitting high on a hill and looking down at the terrain below estimates the distances separating three houses:
A, B, and C. He says that House A is 5 units from both House
B and House C, and that Houses B and C are separated by 8
units. Jill uses these estimates to plot the position of the three houses, as shown in the top portion of Figure I. But then Jack says, "By the way, they are all on the same straight, flat road."
Now Jill knows that Jack's estimates must have contained errors and revises her own in a way that uses all three together to improve each one, to 4.5, 4.5, and 9.0, as shown in the bottom portion of Figure 1.
Three distances among three objects are always consistent in

B

B

A

C

Figure 1. An illustration of the advantage of assuming the correct dimensionality when estimating a set of interpoint distances. Given noisy estimates of AB, AC, and CB, the top portion would be the best guess unless the data source was known to be one-dimensional, in which case the bottom construction would recover the true line lengths more accurately.

THE LATENT SEMANTIC ANALYSIS THEORY OF KNOWLEDGE

two dimensions so long as they obey the triangle inequality (the longest distance must be less than or equal to the sum of the other two). But, knowing that all three distances must be accommodated in one dimension strengthens the constraint (the longest must be exactly equal to the sum of the other two). If the dimensional constraint is not met, the apparent errors in the estimates must be resolved. One compromise is to adjust each distance by the same proportion so as to make two of the lengths add up to the third. The important point is that knowing the dimensionality improves the estimates. Of course, this works the other way around as well. Had the distances been generated from a two- or three-dimensional array (e.g., the road was curved or hilly), accommodating the estimates on a straight line would have distorted their original relations and added error rather than reducing it.
Sometimes researchers have considered dimensionality reduction as a method to reduce computational complexity or for smoothing, that is for simplifying the description of data or interpolating intermediate points (e.g., Church & Hanks, 1990;
Grefenstette, 1994; Schutze, 1992a, 1992b). However, as we will see later, choosing the optimum dimensionality, when appropriate, can have a much more dramatic effect than these interpretations would seem to suggest.
Let us now construe the semantic similarity between two words in terms of distance in semantic space: The smaller the distance, the greater the similarity. Suppose we also assume that two words that appear in the same window of discourse—a phrase, a sentence, a paragraph, or what have you—tend to come from nearby locations in semantic space.1 We could then obtain an initial estimate of the relative similarity of any pair of words by observing the relative frequency of their joint occurrence in such windows.
Given a finite sample of language, such estimates would be quite noisy. Moreover, because of the huge number of words relative to received discourse, many pairwise frequencies would be zero. But two words could also fail to co-occur for a variety of reasons other than thin sampling statistics, with different implications for their semantic similarity. The words might be truly unrelated (e.g., semantic and carburetor). On the other hand, they might be near-perfect synonyms of which people usually use only one in a given utterance (e.g., overweight or corpulent), have somewhat different but systematically related meanings (e.g., purple and lavender), or be relevant to different aspects of the same object (e.g., gears and brakes) and therefore tend not to occur together (just as only one view of the same object may be present in a given scene). To estimate similarity in this situation, more complex, indirect relations (for example, that both gears and brakes co-occur with cars, but semantic and carburetor have no common bridge) must somehow be used.
One way of doing this is to take all of the local estimates of distance into account at once. This is exactly analogous to our houses example, and, as in that example, the choice of dimensionality in which to accommodate the pairwise estimates determines how well their mutual constraints combine to give the right results. That is, we suppose that word meanings are represented as points (or vectors; later we use angles rather than distances) in k dimensional space, and we conjecture that it is possible to materially improve estimates of pairwise meaning

215

similarities, and to accurately estimate the similarities among related pairs never observed together, by fitting them simultaneously into a space of the same (k) dimensionality.
This idea is closely related to familiar uses of factor analysis and multi-dimensional scaling, and to unfolding, (J. D. Carroll & Arabie, in press; Coombs, 1964), but using a particular kind of data and writ very large. Charles Osgood (1971) seems to have anticipated such a theoretical development when computational power eventually rose to the task, as it now has. How much improvement results from optimal dimensionality choice depends on empirical issues, the distribution of interword distances, the frequency and composition of their contexts in natural discourse, the detailed structure of distances among words estimated with varying precision, and so forth.
The scheme just outlined would make it possible to build a communication system in which two parties could come to agree on the usage of elementary components (e.g., words, at least up to the relative similarity among pairs of words). The same process would presumably be used to reach agreement on similarities between words and perceptual inputs and between perceptual inputs and each other, but for clarity and simplicity and because the word domain is where we have data and have simulated the process, we concentrate here on word-word relations. Suppose that a communicator possesses a representation of a large number of words as points in a high dimensional space. In generating strings of words, the sender tends to choose words located near each other. Over short time spans, contiguities among output words would reflect closeness in the sender's semantic space. A receiver could make first-order estimates of the distance between pairs by their relative frequency of occurrence in the same temporal contexts (e.g., a paragraph). If the receiver then sets out to represent the results of its statistical knowledge as points in a space of the same or nearly the same dimensionality as that from which it was generated, it may be able to do better, especially, perhaps, in estimating the similarities of words that never or rarely occur together. How much better depends, as we have already said, on matters that can only be settled by observation.
Except for some technical matters, our model works exactly as if the assumption of such a communicative process characterizes natural language (and, possibly, other domains of natural knowledge). In essence, and in detail, it assumes that the psychological similarity between any two words is reflected in the way they co-occur in small subsamples of language, that the source of language samples produces words in a way that ensures a mostly orderly stochastic mapping between semantic similarity and output distance. It then fits all of the pairwise similarities into a common space of high but not unlimited dimensionality. Because, as we see later, the model predicts what words should occur in the same contexts, an organism using such a mechanism could, either by evolution or learning,

1

For simplicity of exposition, we are intentionally imprecise here in

the use of the terms distance and similarity. In the actual modeling, similarity was measured as the cosine of the angle between two vectors in hyperspace. Note that this measure is directly related to the distance between two points described by the projection of the vectors onto the surface of the hypersphere in which they are contained; thus at a qualitative level the two vocabularies for describing the relations are equivalent.

216

LANDAUER AND DUMAIS

adaptively adjust the number of dimensions on the basis of trial and error. By the same token, not knowing this dimensionality a priori, in our studies we have varied the dimensionality of the simulation model to determine what produces the best results.2
More conceptually or cognitively elaborate mechanisms for the representation of meaning also might generate dimensional constraints and might correspond more closely to the mentalistic hypotheses of current linguistic and psycho-linguistic theories.
For example, theories that postulate meaningful semantic features could be effectively isomorphic to LSA given the identification of a sufficient number of sufficiently independent features and their accurate quantitative assignment to all the words of a large vocabulary. But suppose that it is not necessary to add such subjective interpretations or elaborations for the model to work. Then LSA could be a direct expression of the fundamental principles on which semantic similarity (as well as other perceptual and memorial relations) are built rather than being a reflection of some other system. It is too early to tell whether the model is merely a mathematical convenience that approximates the effects of true cognitive features and processes or corresponds directly to the actual underlying mechanism of which more qualitative theories now current are themselves but partial approximations. The model we propose is at the computational level described by Marr (1982; see also Anderson, 1990), that is, it specifies the natural problem that must be solved and an abstract computational method for its solution.

A Psychological Description of LSA as a Theory of Learning, Memory, and Knowledge
We give a more complete description of LSA as a mathematical model later when we use it to simulate lexical acquisition.
However, an overall outline is necessary to understand a roughly equivalent psychological theory we wish to present first. The input to LSA is a matrix consisting of rows representing unitary event types by columns representing contexts in which instances of the event types appear. One example is a matrix of unique word types by many individual paragraphs in which the words are encountered, where a cell contains the number of times that a particular word type, say model, appears in a particular paragraph, say this one. After an initial transformation of the cell entries, this matrix is analyzed by a statistical technique called singular value decomposition (SVD) closely akin to factor analysis, which allows event types and individual contexts to be re-represented as points or vectors in a high dimensional abstract space (Golub, Luk, & Overton, 1981). The final output is a representation from which one can calculate similarity measures between all pairs consisting of either event types or contexts (e.g., word-word, word-paragraph, or paragraph-paragraph similarities).
Psychologically, the data that the model starts with are raw, first-order co-occurrence relations between stimuli and the local contexts or episodes in which they occur. The stimuli or event types may be thought of as unitary chunks of perception or memory. The first-order process by which initial pairwise associations are entered and transformed in LSA resembles classical conditioning in that it depends on contiguity or co-occurrence, but weights the result first nonlinearly with local occurrence frequency, then inversely with a function of the nu mber of different contexts in which the particular component is encountered overall and the extent to which its occurrences are spread evenly over contexts. However, there are possibly important differences in the details as currently implemented; in particular, LSA associations are symmetrical; a context is associated with the individual events it contains by the same cell entry as the events are associated with the context. This would not be a necessary feature of the model; it would be possible to make the initial matrix asymmetrical, with a cell indicating the co-occurrence relation, for example, between a word and closely following words. Indeed, Lund and Burgess (in press; Lund, Burgess, &
Atchley, 1995). and SchUtze (1992a, 1992b), have explored related models in which such data are the input.
The first step of the LSA analysis is to transform each cell entry from the number of times that a word appeared in a particular context to the log of that frequency. This approximates the standard empirical growth functions of simple learning. The fact that this compressive function begins anew with each context also yields a kind of spacing effect; the association of A
and B is greater if both appear in two different contexts than if they each appear twice in one context. In a second transformation, all cell entries for a given word are divided by the entropy for that word, -S p log p over all its contexts. Roughly speaking, this step accomplishes much the same thing as conditioning rules such as those described by Rescorla & Wagner (1972), in that it makes the primary association better represent the informative relation between the entities rather than the mere fact that they occurred together. Somewhat more formally, the inverse entropy measure estimates the degree to which observing the occurrence of a component specifies what context it is in; the larger the entropy of, say, a word, the less information its observation transmits about the places it has occurred, so the less usage-defined meaning it acquires, and conversely, the less the meaning of a particular context is determined by containing the word.
It is interesting to note that automatic information retrieval methods (including LSA when used for the purpose) are greatly improved by transformations of this general form, the present one usually appearing to be the best (Harman, 1986). It does not seem far-fetched to believe that the necessary transform for good information retrieval, retrieval that brings back text corresponding to what a person has in mind when the person offers one or more query words, corresponds to the functional relations in basic associative processes. Anderson (1990) has drawn attention to the analogy between information retrieval in external systems and those in the human mind. It is not clear which way the relationship goes. Does information retrieval in automatic systems work best when it mimics the circumstances that make people think two things are related, or is there a general logic that tends to make them have similar forms? In automatic information retrieval the logic is usually assumed to be that idealized searchers have in mind exactly the same text as they would like the system to find and draw the words in
2

Although this exploratory process takes some advantage of chance,

there is no reason why any number of dimensions should be much better than any other unless some mechanism like the one proposed is at work.
In all cases, the model's remaining parameters were filled only to its input (training) data and not to the criterion (generalization) test.

THE LATENT SEMANTIC ANALYSIS THEORY OF KNOWLEDGE

their queries from that text (see Bookstein & Swanson, 1974).
Then the system's challenge is to estimate the probability that each text in its store is the one that the searcher was thinking about. This characterization, then, comes full circle to the kind of communicative agreement model we outlined above: The sender issues a word chosen to express a meaning he or she has in mind, and the receiver tries to estimate the probability of each of the sender's possible messages.
Gallistel (1990), has argued persuasively for the need to separate local conditioning or associative processes from global representation of knowledge. The LSA model expresses such a separation in a very clear and precise way. The initial matrix after transformation to log frequency divided by entropy represents the product of the local or pairwise processes.3 The subsequent analysis and dimensionality reduction takes all of the previously acquired local information and turns it into a unified representation of knowledge.
Thus, the first processing step of the model, modulo its associational symmetry, is a rough approximation to conditioning or associative processes. However, the model's next steps, the singular value decomposition and dimensionality optimization, are not contained as such in any extant psychological theory of learning, although something of the kind may be hinted at in some modern discussions of conditioning and, on a smaller scale and differently interpreted, is often implicit and sometimes explicit in many neural net and spreading-activation architectures.
This step converts the transformed associative data into a condensed representation. The condensed representation can be seen as achieving several things, although they are at heart the result of only one mechanism. First, the re-representation captures indirect, higher-order associations. That is, if a particular stimulus, X, (e.g., a word) has been associated with some other stimulus, Y, by being frequently found in joint context (i.e., contiguity), and Y is associated with Z, then the condensation can cause X and Z to have similar representations. However, the strength of the indirect XZ association depends on much more than a combination of the strengths of XY and YZ. This is because the relation between X and Z also depends, in a wellspecified manner, on the relation of each of the stimuli, X, Y, and Z, to every other entity in the space. In the past, attempts to predict indirect associations by stepwise chaining rules have not been notably successful (see, e.g., Pollio, 1968; \bung,
1968). If associations correspond to distances in space, as supposed by LSA, stepwise chaining rules would not be expected to work well; if X is two units from Y and Y is two units from
Z, all we know about the distance from X to Z is that it must be between zero and four. But with data about the distances between X, Y, Z, and other points, the estimate of XZ may be greatly improved by also knowing XY and YZ.
An alternative view of LSA's effects is the one given earlier, the induction of a latent higher order similarity structure (thus its name) among representations of a large collection of events.
Imagine, for example, that every time a stimulus (e.g., a word)
is encountered, the distance between its representation and that of every other stimulus that occurs in close proximity to it is adjusted to be slightly smaller. The adjustment is then allowed to percolate through the whole previously constructed structure of relations, each point pulling on its neighbors until all settle into a compromise configuration (physical objects, weather sys-

217

tems, and Hopfield nets do this too; Hopfield, 1982). It is easy to see that the resulting relation between any two representations depends not only on direct experience with them but with everything else ever experienced. Although the current mathematical implementation of LSA does not work in this incremental way, its effects are much the same. The question, then, is whether such a mechanism, when combined with the statistics of experience, produces a faithful reflection of human knowledge.
Finally, to anticipate what is developed later, the computational scheme used by LSA for combining and condensing local information into a common representation captures multivariate correlational contingencies among all the events about which it has local knowledge. In a mathematically well-defined sense it optimizes the prediction of the presence of all other events from those currently identified in a given context and does so using all relevant information it has experienced.
Having thus cloaked the model in traditional memory and learning vestments, we next reveal it as a bare mathematical formalism.

A Neural Net Analog of LSA
We describe the matrix-mathematics of singular value decomposition used in LSA more fully, but still informally, next and in somewhat greater detail in the Appendix. But first, for those more familiar with neural net models, we offer a rough equivalent in that terminology. Conceptually, the LSA model can be viewed as a simple but rather large three-layered neural net. It has a Layer 1 node for every word type (event type), a
Layer 3 node for every text window (context or episode) ever encountered, several hundred Layer 2 nodes—the choice of number is presumed to be important—and complete connectivity between Layers 1 and 2 and between Layers 2 and 3. (Obviously, one could substitute other identifications of the elements and episodes). The network is symmetrical; it can be run in either direction. One finds an optimal number of middle-layer nodes, then maximizes the accuracy (in a least-squares sense)
with which activating any Layer 3 node activates the Layer 1
nodes that are its elementary contents, and, simultaneously, vice versa. The conceptual representation of either kind of event, a unitary episode or a word, for example, is a pattern of activation across Layer 2 nodes. All activations and summations are linear.
Note that the vector multiplication needed to generate the middle-layer activations from Layer 3 values is, in general, different from that to generate them from Layer 1 values. Thus a different computation is required to assess the similarity between two episqdes, two event types, or an event type and an episode, even though both kinds of entities can be represented as values in the same middle-layer space. Moreover, an event type or a set of event types could also be compared with another of the same or with an episode or combination of episodes by computing their activations on Layer 3. Thus the network can

3

Strictly speaking, the entropy operation is global, added up over all

occurrences of the event type (conditioned stimulus; CS), but it is here represented as a local consequence, as might be the case, for example, if the presentation of a CS on many occasions in the absence of the unconditioned stimulus (US) has its effect by appropriately weakening the local representation of the CS-US connection.

218

LANDAUER AND DUMAIS

create artificial or "imaginary" episodes, and, by the inverse operations, episodes can generate "utterances" to represent themselves as patterns of event types with appropriately varying strengths. The same things are true in the equivalent singularvalue-decomposition matrix model of LSA.

The Singular Value Decomposition (SVD)
Realization of LSA
The principal virtues of SVD for this research are that it embodies the kind of inductive mechanisms that we want to explore, that it provides a convenient way to vary dimensionality, and that it can fairly easily be applied to data of the amount and kind that a human learner encounters over many years of experience. Realized as a mathematical data-analysis technique, however, the particular model studied should be considered only one case of a class of potential models that one would eventually wish to explore, a case that uses a very simplified parsing and representation of input and makes use only of linear relations.
In possible elaborations one might want to add features that make it more closely resemble what we know or think we know about the basic processes of perception, learning, and memory.
It is plausible that complicating the model appropriately might allow it to simulate phenomena to which it has not been applied and to which it currently seems unlikely to give a good account, for example certain aspects of grammar and syntax that involve ordered and hierarchical relations rather than unsigned similarities. However, what is most interesting at this point is how much it does in its present form.

Singular Value Decomposition (SVD)
SVD is the general method for linear decomposition of a matrix into independent principal components of which factor analysis is the special case for square matrices with the same entities as columns and rows. Factor analysis finds a parsimonious representation of all the intercorrelations between a set of variables in terms of a new set of abstract variables, each of which is unrelated to any other but which can be combined to regenerate the original data. SVD does the same thing for an arbitrarily shaped rectangular matrix in which the columns and rows stand for different things, as in the present case one stands for words, the other for contexts in which the words appear. (For those with yet other vocabularies, SVD is a form of eigenvalueeigenvector analysis or principal components decomposition and, in a more general sense, of two-way, two-mode multidimensional scaling (see J. D. Carroll & Arabic, in press).
To implement the model concretely and simulate human word learning, SVD was used to analyze 4.6 million words of text taken from an electronic version of Grolier's Academic American Encyclopedia, a work intended for young students. This encyclopedia has 30,473 articles. From each article we took a sample consisting of (usually) the whole text, or its first 2,000
characters, whichever was less, for a mean text sample length of 151 words, roughly the size of a rather lo

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
