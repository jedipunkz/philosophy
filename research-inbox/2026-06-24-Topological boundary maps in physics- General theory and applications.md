---
source: "https://arxiv.org/abs/math-ph/0605048v1"
title: "Topological boundary maps in physics: General theory and applications"
author: "Johannes Kellendonk, Serge Richard"
year: "2006"
publication: "arXiv preprint / math-ph"
download: "https://arxiv.org/pdf/math-ph/0605048v1"
pdf: "https://arxiv.org/pdf/math-ph/0605048v1"
captured_at: "2026-06-24T11:10:45Z"
updated_at: "2026-06-24T11:10:45Z"
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

# Topological boundary maps in physics: General theory and applications

- 著者: Johannes Kellendonk, Serge Richard
- 年: 2006
- 掲載情報: arXiv preprint / math-ph
- 情報源: [arxiv](https://arxiv.org/abs/math-ph/0605048v1)
- ダウンロード: https://arxiv.org/pdf/math-ph/0605048v1
- PDF: https://arxiv.org/pdf/math-ph/0605048v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

The material presented here covers two talks given by the authors at the conference Operator Algebras and Mathematical Physics organised in Bucharest in August 2005. The first one was a review given by J. Kellendonk on the relation between bulk and boundary topological invariants in physical systems. In the second talk S. Richard described an application of these ideas to scattering theory. It leads to a topological version of the so-called Levinson's theorem.

## PDF Text

arXiv:math-ph/0605048v1 16 May 2006

Topological boundary maps in physics:
General theory and applications
Johannes Kellendonk and Serge Richard

Institut Camille Jordan, Bâtiment Braconnier, Université Claude Bernard Lyon 1,
43 avenue du 11 novembre 1918, 69622 Villeurbanne cedex, France
E-mails: kellendonk@math.univ-lyon1.fr and srichard@math.univ-lyon1.fr
May 2006

Abstract
The material presented here covers two talks given by the authors at the conference Operator Algebras and Mathematical Physics organised in Bucharest in August 2005. The first one was a review given by J. Kellendonk on the relation between bulk and boundary topological invariants in physical systems. In the second talk S. Richard described an application of these ideas to scattering theory. It leads to a topological version of the so-called Levinson’s theorem.

Introduction
The natural language for quantum physics is linear operators on Hilbert spaces and underlying operator algebras. These algebras are fundamentally non-commutative. Topological properties of quantum systems should hence be connected with the topology of these algebras, which is what one calls non-commutative topology. An important first question to be answered is therefore: what is the correct operator algebra related to a physical system? Since we are looking for topological effects this algebra should be a separable C ∗ -algebra and a good starting point is to look for the C ∗ -version of the observables algebra. Once the question about the right algebra is settled we are interested in studying its invariants asking above all: which of them have a physical interpretation? Finally, when we have identified the invariants, we want to derive relations between them, typically equations between topological quantised transport coefficients or, as in Levinson’s theorem, between invariants of the bounded part and the scattering part of the physical system. Such relations can be obtained from topological boundary maps which do not exist on the algebraic level.
The purpose of this paper is twofold: Explain with more details the general theory outlined in the previous paragraph, and show its relevance in various applications in mathematical physics.
1

The first section is devoted to a brief introduction to the natural framework of topological boundary maps and to the description of the the general theory. The second section contains examples of applications to solid states physics, while the third one is entirely dedicated to an application to potential scattering. Since crossed product C ∗ -algebras and their twisted versions play an important rôle in the applications, we have decided to incorporate an appendix on these algebras. Let us finally mention that Proposition 4.1 on the decomposition of magnetic twisted crossed product C ∗ -algebras as iterated twisted crossed products is of independent interest.

1

General theory of topological boundary maps in physics

A C ∗ -algebra is a special kind of Banach algebra. For our purposes, the fact that its norm satisfies the so-called C ∗ -condition does not play an important rôle, but we wish it to be separable, i.e. to contain a countable dense set. Infinite dimensional von Neumann algebras are non-separable
C ∗ -algebras and therefore not suited. K-groups are topological invariants of C ∗ -algebras: they are abelian groups, which are countable for separable C ∗ -algebras. They are isomorphic for isomorphic algebras and one may think of them as simpler objects which might tell apart C ∗ algebras. A very concise formulation of our philosophy is the following: If the C ∗ -algebra is somewhat naturally assigned to a physical system, then the elements of its K-groups are to be understood as topological invariants of that system.

1.1

K-groups and n-traces

The K0 -group of a unital C ∗ -algebra C is constructed from the homotopy classes of projections in the set of square matrices with entries in C . Its addition is induced from the addition of two orthogonal projections: if p and q are orthogonal projections, i.e. pq = 0, then also p + q is a projection. Thus, the sum of two homotopy classes [p]0 + [q]0 is defined as the class of the sum of the block matrices [p ⊕ q]0 on the diagonal. This new class does not depend on the choice of the representatives p and q. K0 (C ) is defined as the Grothendieck group of this set of homotopy classes of projections endowed with the mentioned addition. In other words, the elements of the
K0 -group are given by formal differences: [p]0 − [q]0 is identified with [p′ ]0 − [q ′ ]0 if there exists a projection r such that [p]0 + [q ′ ]0 + [r]0 = [p′ ]0 + [q]0 + [r]0 . In the general non-unital case the construction is a little bit more subtle.
The K1 -group of a C ∗ -algebra C is constructed from the homotopy classes of unitaries in the set of square matrices with entries in the unitisation of C . Its addition is again defined by:
[u]1 +[v]1 = [u⊕v]1 as a block matrix on the diagonal. The homotopy class of the added identity is the neutral element.
For our purpose, higher traces will always be constructed from ordinary traces and derivations, which might both be unbounded. More precisely, an n-trace on a C ∗ -algebra C is determined by the data (T ; δ1 , . . . , δn ) where T is a trace on C and {δj }nj=1 are n commuting derivations on C which leave the trace invariant: T ◦ δj = 0. More important than the fact that their characters define unbounded cyclic cocycles is for us that they define additive functionals on the K-groups by Connes’ pairing: Extending T and δj to matrices with entries in C in the
2

canonical way one has, up to constants, for even n a functional h(T ; δ1 , . . . , δn ), ·i : K0 (C ) → C
defined by
X

sgn(π) T pδπ(1) (p) · · · δπ(n) (p) , h(T ; δ1 , . . . , δn ), [p]i =
π∈Sn

and for odd n a functional h(T ; δ1 , . . . , δn ), ·i : K1 (C ) → C defined by
X

sgn(π) T (u∗ − 1)δπ(1) (u)δπ(2) (u∗ ) · · · δπ(n) (u) .
h(T ; δ1 , . . . , δn ), [u]i =
π∈Sn

Here, Sn is the group of permutations of n elements.

1.2

The general theory

Let us consider a quantum system described by a linear operator in a Hilbert space H, and let C be a C ∗ -subalgebra of B(H) that is related with this system. Here and in the sequel,
B(H) denotes the C ∗ -algebra of all bounded operators in H and K(H) the ideal of compact operators in H. For instance, the system is described by a self-adjoint operator H in H and C
contains the C0 -functional calculus of H. Suppose now that we can identify certain elements of the K-groups of C with physically meaningful quantities. For example, the spectral projection
P(−∞,c) (H) of H would give rise to an element of K0 (C ), provided H is bounded from below and the value c lies in a gap of the spectrum of H. Since the elements of the K-groups exhibit some homotopy invariance, we expect that they will be stable under certain perturbations of the system. Suppose moreover that we have a higher trace such that its pairing with the K-groups describes a physically significant quantity. Then this quantity is topologically quantised, i.e. it takes values in a countable subgroup of the real numbers.
Now assume that we have two quantum systems, the first one related with a C ∗ -algebra
J and the second one with a C ∗ -algebra C . Assume moreover that these are related via an extension, i.e. there exists a third algebra E such that J is an ideal of E and C is isomorphic to the quotient E /J . Another way of saying this is that J and C are the left and right part of an exact sequence of C ∗ -algebras q

i

0 → J → E → C → 0, i being an injective morphism and q a surjective morphism satisfying ker q = im i. There might not be any reasonable algebra morphism between J and C but algebraic topology provides us with homomorphisms between their K-groups: ind : K1 (C ) → K0 (J ) and exp : K0 (C ) →
K1 (J ), the index map and the exponential map. These maps, which are also referred to as boundary maps, allow us to relate topological invariants of the two systems. Furthermore, with a little luck we also obtain dual maps on the functionals defined by higher traces and therefore equations between numerical topological invariants. Therefore, the procedure goes as follows:
1. Find a suitable C ∗ -algebra related with a given quantum system.
2. Identify K-elements and higher traces whose pairings allow for a physical interpretation.
3

3. Construct extensions to the C ∗ -algebras related with two different systems and compute the boundary maps to obtain relations between topological quantised quantities.
The reader may wonder at this point that operators don’t seem to come up at all in the picture.
This is missleading. An important part of the first and third step is actually to prove that the operators describing the system are related in some sense to the algebra. This affiliation is often a difficult analytical problem! Furthermore, the physical interpretation of pairings involves as well the operators.

2

Applications to solid states physics

One of the most common realisations of the ideas presented above is furnished by a quantum system described by a self-adjoint operator H in a Hilbert space H and a norm closed subalgebra
C of B(H) that can be considered as the algebra of observables. In particular, C is expected to contain the energy observables η(H), obtained from functions η which belong to C0 (R), the algebra of continuous functions on R that vanish at infinity. This section is devoted to a presentation of such realisations in the context of solid states physics.
Since crossed product C ∗ -algebras and their twisted versions are discussed in the appendix, we shall not recall their definitions in this section.

2.1

Algebras of energy observables derived from the set of atomic postions

An important class of C ∗ -algebras of observables can be obtained from the geometry of the set
P of equilibrium atomic positions in a solid. P is a discrete subset of Rn which we suppose to be of finite local complexity, i.e. for each r > 0, there exists only finitely many so-called r-patches
(P − x) ∩ Br , with x varying in P. Br denotes the closed ball centered at 0 and of radius r.
A continuous function f : Rn → C is called P-equivariant with range r if Br ∩ (P − x) =
Br ∩ (P − y) implies f (x) = f (y), x, y ∈ Rn . The sup-norm closure of all P-equivariant functions with arbitrary range is called the C ∗ -algebra CP (Rn ) of P-equivariant P
functions. A typical example of a P-equivariant function is a potential V defined by V (x) := y∈P v(x − y) where v is a short range atomic potential, i.e. a function which decays sufficiently fast for the sum to be finite. The unital algebra CP (Rn ) carries the continuous Rn -action α by translation.
Then the algebra of the aperiodic structure described by P is the corresponding crossed product
C ∗ -algebra CP := CP (Rn ) ⋊α Rn .
Suppose now that the system is in an exterior constant magnetic field whose components we

denote by {Bjk }nj,k=1 with Bjk ∈ R. For any x, y ∈ Rn , let ω B (x, y) := exp − iΓB h0, x, x + yi , where ΓB h0, x, x + yi is the flux of the magnetic field though the triangle defined by the points 0, x and x + y. We refer to the appendix for the general construction in the case of a non-constant magnetic field with components in CP (Rn ). One may thus form the magnetic twisted crossed n
B
product C ∗ -algebra CP (Rn ) ⋊B
α R associated with the twisted actions (α, ω ) . This algebra is
B
simply denoted by CP , and if the magnetic field vanishes, then this algebra corresponds to CP .

4

The important fact is the following: Let A := {Aj }nj=1 with Aj : Rn → R be a continuous vector potential for the magnetic field. Let H = (P − A)2 + V be the Landau operator perturbed by a potential V which is a P-equivariant real function. This magnetic Schrödinger operator, which is a self-adjoint operator in H := L2 (Rn ), is affiliated to CP in the following sense. There exists a faithfull representation π : CPB → B(H) such that for any η ∈ C0 (R), η(H) ∈ π(CPB ).
Note that the representation is constructed with the help of the vector potential A. It can be argued that CPB is the algebra of observables for the system describing the motion of an electric particle in the aperiodic solid described by P under the influence of the constant magnetic field
B.
Let us discuss now some examples or related constructions.
Finite systems without magnetic field. Suppose that we have a finite system like an atom or a molecule and no external magnetic field. In that case P would simply be a finite set contained in some ball Br′ , say. Therefore a P-equivariant function with range r would be constant outside
Br+r′ . Hence CP (Rn ) is the unitisation of C0 (Rn ) and CP = C0 (Rn ) ⋊α Rn + C ⋊ Rn , which is commonly
 called the two-body algebra. Let us note that the first summand is isomorphic to
K L2 (Rn ) and hence an ideal.
Crystals and quasi-crystals without magnetic field. Before the discovery of quasi-crystals a crystal was considered to be a periodic arrangement of atoms, the set P being therefore a regular lattice. In that case, CP (Rn ) is simply the algebra of continuous P-periodic functions on
Rn . The corresponding crossed product algebra CP can be seen as the C ∗ -algebra of observables associated with this periodic crystal.
Idealised quasi-crystals are often described by quasi-periodic sets P. For example such a set can be obtained from a cut & project scheme. More generally it has been proposed to describe aperiodic ordered systems by repetitive Delone sets P of finite local complexity with uniform existence of patch frequencies. In this case CP (Rn ) is a lot more complicated. Its spectrum ΩP is a foliated space which is transversally totally disconnected. Repetivity corresponds to simplicity of the algebra CP , i.e. absense of non-trivial closed ideals, and uniform existence of patch frequencies to the fact that CP (Rn ) carries a unique invariant normalised trace τ . We get then
 a
0-trace on CP by defining T (F ) = τ (F (0)) on any continuous element F ∈ L1 Rn , CP (Rn ) of the crossed product.
We add the remark that the unique invariant trace on CP (Rn ) corresponds to a unique invariant ergodic probability measure on its spectrum ΩP . The condition of P having uniform existence of patch frequencies can be relaxed leading to the freedom of choice for the trace which corresponds to a choice of ergodic probability measure on ΩP and may be interpreted as a choice of physical phase.
Solids with boundary in a constant magnetic field. We consider a solid described by
P restricted to the half-space Rn−1 × (−∞, s], and therefore with a boundary at x⊥ ≡ xn = s.
The new variable s describes the relative position between the P-equivariant potential and
5

the boundary. We let it vary over R ∪ {+∞}, s = +∞ corresponding to the system without boundary.
In order to construct a suitable C ∗ -algebra for that system, we rewrite CPB as a crossed product by R,

n ∼
n
B
n−1
CP (Rn ) ⋊B
⋊β R =: C .
(1)
α R = CP (R ) ⋊αk R

n−1 is obtained by restricting the action to translations which are parallel to
Here CP (Rn ) ⋊B
αk R

the boundary and the twisting cocycle to ωkB (x, y) := ω B (x, 0), (y, 0) for x, y ∈ Rn−1 . In the case of constant magnetic field this could be understood as a choice of gauge, but a more elegant approach which uses only gauge invariant quantities and generalises to variable magnetic field is presented in the appendix, Proposition 4.1. The action β contains, of course, the part coming from translations perpendicular to the boundary, see the appendix. Now, for the system with boundary it is natural to consider the algebra


n−1
⋊γ⊗β R ,
(2)
E := C0 (R ∪ {+∞}) ⊗ CP (Rn ) ⋊B
R
αk

which is the Wiener-Hopf extension of (1). γ is the translation action on C0 (R ∪ {+∞}) which has +∞ as fixed point. The evaluation at +∞ defines a surjective morphism from E onto C .
The important fact, proved for n = 2 in [19], is the following: Let b be the component of the magnetic field pointing in the direction perpendicular to the plane, and let Hs be the restriction of H = (P1 − bQ2 )2 + P22 + V to the half space R × (−∞, s] with Dirichlet boundary conditions.
The family {Hs }s∈R∪{+∞} is affiliated to E in the following sense:
 For any η ∈ C0 (R), there exists F ∈ E such that η(Hs ) = πs (F ), where πs : E → B L2 (R2 ) is a representation induced by the evaluation map at (s, 0), evs,0 : C0 (R ∪ {+∞}) ⊗ CP (R2 ) → C. The possibility of letting s tend to infinity allows to relate continuously η(H) = π+∞ (F ) with η(Hs ) = πs (F ) for any s ∈ R. Whereas the individual representations πs are not faithful, their direct sum is faithful.
From this, it can be argued that E is the algebra of observables for the family over s of systems describing the motion of an electric particle, in the aperiodic solid described by P
and under the influence of the constant magnetic field B, that is confined to the half space
R × (−∞, s].
The edge algebra. Let us describe the ideal J which is the kernel of the surjection E → C .
It is


n−1
⋊γ⊗β R ,
(3)
R
J := C0 (R) ⊗ CP (Rn ) ⋊B
αk


n−1 . Its elements are thus limits of which is isomorphic to C0 (R) ⋊γ R ⊗ CP (Rn ) ⋊B
αk R
elementary tensors F⊥ ⊗ Fk where F⊥ is a compact operator, and F⊥ is an element of the aln−1 . J is therefore the algebra of observables which are localised near gebra CP (Rn ) ⋊B
αk R
the boundary (or edge) in the loose sense of being compact in the perpendicular direction. We call it the edge algebra. As for CP we can construct a traceon J starting from the trace τ
on CP (Rn ), namely we define T̂ (F⊥ ⊗ Fk ) := Tr(F⊥ )τ Fk (0) , with Tr the standard trace on compact operators, onany trace class element F⊥ of C0 (R) ⋊γ R and any continuous element
Fk ∈ L1 Rn−1 , CP (Rn ) of the crossed product.
6

2.2

Examples of pairings with physical interpretation

Let us present some systems in which the pairing of a K-element with a higher trace has a physical interpretation.
The simplest construction consists in using the K0 -elements of an algebra of observables CPB
defined by the spectral projections and a trace on the algebra to pair with them. For example, consider the K0 -elements defined by the projection P(−∞,EF ) (H) of the Hamiltonian H to the energies below the Fermi energy EF , provided this value lies in a gap of spectrum of H. Pairing it with a suitable trace yields IDS(EF ), the integrated density of states at the Fermi level.
In the absence of magnetic field, we mention that also the pairing of the full K0 -group with the 0-trace T , the so-called gap-labelling group has been of great interest. It describes the set of possible gap labels of a physical system. One can construct an element in the K0 -group for each r-patch of P. Pairing this element with the 0-trace T yields the frequency of the patch.
Note that the notion of frequency depends on the choice of ergodic measure on ΩP . As has been proved relatively recently, these frequencies generate the gap-labelling group [5, 6, 16].
The most famous example is related to the topological quantisation of the Integer Quantum
Hall Effect. One typically finds two models used to describe the quantisation. In the bulk model,
2
the sample is modelled by a Hamiltonian H on L2 (R2 ) which is affiliated to CP (R2 ) ⋊B
α R .
Here the Hall-conductivity σH , the transverse component of the conductivity tensor, is up to a universal constant the pairing between the K0 -element determined by the spectral projection
P(−∞,EF ) (H), provided the Fermi energy EF lies in a gap
 of its spectrum, and the 2-trace
(T ; δ1 , δ2 ). In the representation of the algebra in B L2 (R2 ) discussed above, T is the trace per unit volume and δj = i[Qj , ·], the commutator with the j-component of the position operator.
We refer to [4] for a discussion of the tight binding case where it is also explained that the condition that EF belongs to a gap can be relaxed to EF belongs to a mobility gap.
Examples of pairings related to edge states. The edge algebra J is the algebra in which we expect to find the operators describing the physics on the boundary. We first construct an element of its K1 -group.
Let ∆ be an interval contained in a gap of the spectrum of the Hamiltonian H. Then P∆ (Hs )
will not be 0 for s < ∞, but rather the projection onto the edge states with energy in ∆. Now, consider the bounded continuous function u : R → C defined for all t ∈ R by



(t
−
inf
∆)
−
1
, u(t) = 1 + χ∆ (t) exp −2πi
|∆|

where χ∆ is the characteristic function on the interval ∆. Thus, there exists an element U −1 ∈ E
such that πs (U − 1) = u(Hs ) − 1. However, since u(H∞ ) − 1 ≡ u(H) − 1 = 0, it follows that U
belongs to J ⊂ E . Therefore U defines an element of K1 (J ).
To construct odd higher traces we make use of the trace constructed on J and consider the derivations δj for j 6= n as above and ∂⊥ := i[P⊥ , ·] ≡ i[Pn , ·], the commutator with the infinitesimal generator of translation in position space perpendicular to the boundary.
The pairings of [U ] with the 1-traces (T̂ , δj ) for j 6= n, and (T̂ , ∂⊥ ) have physical interpretation. First of all, the trace T̂ of an element F⊥ ⊗ Fk may be interpreted as an average, namely
7

the average over the position s of the boundary of the usual trace on L2 (R) times the trace
1
h(Tˆ , δj ), [U ]i is the average of the operator per unit volume on L2 (Rn−1 ) of πs (F⊥ ⊗ Fk ). 2π
1
1
|∆| P∆ (Hs )[Qj , Hs ]P∆ (Hs ). Since this operator is |∆| times the j-component of the current op1
erator restricted to the edge states, σj := 2π
h(Tˆ , δj ), [U ]i is the j-direction of the conductivity along the boundary provided the Fermi energy lies in ∆.
1
1
∂V
P∆ (Hs ). Since the operator h(Tˆ , ∂⊥ ), [U ]i is the average of |∆|
Similarily Π := 2π
P∆ (Hs ) ∂x n

∂V
P∆ (Hs ) is the perpendicular component of the gradient force restricted to the edge
P∆ (Hs ) ∂x n
states, Π can be understood as the gradient pressure per unit energy on the boundary of the system again supposing that the Fermi energy lies in ∆.

2.3

Relating two systems

Now we consider two quantum systems, one with an algebra of observables J the other with an algebra of observables C , which are related via an extension E . It turns out that in all our examples the extensions are Wiener-Hopf extensions determined by a continuous action β of R
on some auxiliary C ∗ -algebra B. These are abstractly defined as follows: Given a C ∗ -algebra
B with a continuous R-action β, the Wiener-Hopf extension of B ⋊β R is the crossed product
C ∗ -algebra E := C0 (R ∪ {+∞}) ⊗ B ⋊γ⊗β R. The evaluation at +∞ for C0 (R ∪ {+∞}) gives a surjective morphism onto B ⋊β R, which we assumed to be equal to C . The kernel of this morphism is supposed to be equal to J . One good feature of the Wiener-Hopf extension is that the K-groups of E always vanish making the boundary maps ind and exp isomorphisms. They are in fact the inverses of the Connes-Thom isomorphism. Another advantage is that the dual maps on functionals defined by higher traces are simple and explicit.
Solids with boundary. In the context of aperiodic ordered solids with boundary we have already seen that the algebra E is the Wiener-Hopf extension of C with ideal J , all these algebras being defined in equations (2), (1) and (3). So let us consider the boundary map on
K-theory exp : K0 (C ) → K1 (J ). Under the assumption that the Fermi energy EF belongs to an interval ∆ which does not overlap with the ”bulk” spectrum one can show that the image of the K0 -class defined by the projection P(−∞,EF ) (H) under this map is the K1 -class defined by the unitary U above.
To describe the dual map we consider for simplicity the case n = 2. In this case the map identifies the functional defined by the 2-trace (T , δ1 , δ2 ) with the functional defined by the
1-trace (Tˆ , δ1 ). This leads to the relation
σ1 = σH
which expresses the fact that the Hall conductivity defined as the transverse component of the conductivity tensor in the bulk equals the conductivity σ1 of the current along the edge [18, 20].
The dual map identifies furthermore the functional defined by the 0-trace T with that defined by the 1-trace (T̂ , ∂s ) where ∂s is the infinitesimal generator of translation of the boundary. The

8

latter functional is a linear combination of the functionals defined by (Tˆ , δ1 ) and (T̂ , ∂⊥ ). As a consequence we get the relation
IDS = Π + Bσ1 .
This relation is valid at the Fermi energy provided it belongs to a gap of the spectrum. We note that IDS is a bulk quantity which cannot be obtained from a measurement of the density of states near the Fermi energy, since IDS(EF ) depends on the density of states at all energies below EF . By constrast, Π and Bσ1 need only to be measured in an arbitrarily small interval containing EF . Note that they are a priori introduced as quantities depending on an interval containing EF but turn out to be largely independent of that choice. We mention that derivating
∂
IDS = σH [27].
the above relation w.r.t. the magnetic field strength B yields Streda’s formula ∂B

3

Application to scattering theory

Let us start by recalling very heuristically the main idea of potential scattering. We consider a wave packet that is prepared in the far past far enough from a probe. Since we assume that the probe is of finite size, this initial wave packet is presumably asymptotically free. It is then supposed to evolve in time under the influence of the potential describing the probe to then move far away from the probe so that it can again be considered asymptotically free in the far future. It is commonly expected that all the observable information of the scattering process is contained in the so-called S-operator, an operator that relates the initial wave packet with the final wave packet. Under some weak hypotheses, this operator is unitary. On the other hand, the probe can possibly bind some states. In that situation, the projection on these states is... a projection! Thus, we face a situation in which there exist a unitary operator and a projection that are related with two connected systems: The system of a scattering process by a probe and the system consisting of the bound states of that probe. Having in mind the general theory presented in Section 1.2, one is naturally led to consider an algebraic framework that can link these two objects. This section is devoted to such a construction in the case of a two-body
Schrödinger operator. Other applications and extensions are in preparation [17].

3.1

The framework

Let us consider the self-adjoint operators H0 := −∆ and H := H0 + V in the Hilbert space
H := L2 (Rn ), where |V (x)| ≤ c (1 + |x|)−β with β > 1. It is well known that for such short range potentials V , the wave operators
Ω± := s − lim eitH e−itH0
t→±∞

(4)

exist and have same range. The complement of this range is spanned by the eigenvectors of H, we let P denote the projection on this subspace. The scattering operator S for this system is defined by the product Ω∗+ Ω− , where Ω∗+ is the adjoint of Ω+ .
Levinson’s theorem establishes a relation between an expression in terms of the unitary operator S and an expression depending on the projection P . There exist many presentations of
9

this theorem, but we recall only the one of [22]. We refer to [8], [12] and [23] for other versions of a similar result.

Let U : H → L2 R+ ; L2 (Sn−1 ) be the unitary transformation that diagonalizes H0 , i.e. that satisfies [U H0 f ](λ, ω) = λ[U f ](λ, ω), with f in the domain of H0 , λ ∈ R+ and ω ∈ Sn−1 . Since the operator S commutes with H0 , there exists a family {S(λ)}λ∈R+ of unitary operators in
L2 (Sn−1 ) satisfying U S U ∗ = {S(λ)} almost everywhere in λ [3, Chapter 5.7]. Under suitable hypotheses on V [22] and in the case n = 3, Levinson’s theorem takes the form
Z ∞

√ν
(5)
dλ tr[iS(λ)∗ dS
dλ (λ)] − λ = 2π Tr[P ],
0

R
where tr is the trace on L2 (Sn−1 ), Tr the trace on H and ν = (4π)−1 R3 dx V (x). Clearly the r.h.s. of this equality is invariant under variations of V that do not change the number of bound states of H. But it is not at all clear how this stability comes about in the l.h.s.
In the sequel we propose a modification of the l.h.s. of (5) in order to restore the topological nature of this equality. The idea is very natural from the point of view presented in Section 1.2: we rewrite the l.h.s. of (5) as the result of a pairing between an element of K1 and a 1-trace.
Beyond formula (5), we show that the unitary S is related to the projection P at the level of
K-theory by the index map, cf. Theorem 3.2. Let us point out that the wave operators play a key role in this work. Sufficient conditions on Ω− imply that H has only a finite set of bound states, but also give information on the behaviour of S(·) at the origin.

3.2

A suitable short exact sequence and its representation

In this section we construct a short exact sequence, i.e. an extension of two algebras. One algebra is associated with the scattering system and the other with the bound state system. We permit ourselves to do that twice, first in a heuristic way similar to Section 2.1 and then again more rigorously, shifting attention to the wave operator. This will lead us to natural hypotheses under which we obtain a relation between the scattering operator S and the projection P on the bound states via a boundary map of K-theory.
The size of the probe being finite, it could be described by finite set P in the spirit of
Section 2.1. For such a system we obtained the two body algebra CP = C0 (Rn ) ⋊α Rn + C ⋊ Rn.
The first summand is isomorphic to K L2 (Rn ) , and we expect the projection onto the bound states to be compact, supposing that there are only finitely many. So the algebra of the bound state system should be that ideal of CP . Although CP is actually an extension of C ⋊ Rn by the ideal, it is not this extension which will be used.
The algebra describing the scattering part should contain all possible S-operators. Writing an S-operator as a unitary
 operator valued function of energy as above, it is therefore contained in L∞ R+ ; B L2 (Sn−1 ) . We now make assumptions which allow us to obtain topological infor
mation: (1) the map λ 7→ S(λ) is continuous w.r.t. norm topology, (2) S(λ) − 1 ∈ K L2 (Sn−1 ) ,
(3) S(0) = S(∞) = 1. These will be a consequenceof our hypotheses
 below. They allow us
2
n−1
2
n−1
∼
to regard S − 1 as an element of C0 R+ ; K L (S
) = K L (S
) ⋊ R which we therefore
10

consider as the algebra of the scattering system. The action in this crossed product is trivial and the isomorphism is given by Fourier transformation.
The extension we will use is


ev
(6)
0 → C0 R; K ⋊γ R → C0 R ∪ {+∞}; K ⋊γ R →∞ K ⋊R → 0, where K is the algebra of compact operators in some Hilbert space. The sequence (6) is the
Wiener-Hopf extension of the crossed product K ⋊R with trivial R-action on K ; γ is the action

on C0 R ∪ {+∞} by translation, leaving the point {+∞} invariant, and the surjection ev∞
is induced at {+∞}. Note that setting K = K L2 (Sn−1 ) we have naturally
 by evaluation
2
∼
C0 R; K ⋊γ R = K L (Rn ) and so we expect P to be in the left algebra and S − 1 to be in the right algebra. But instead of verifying that directly we change perspective and concentrate on the middle algebra with the goal to identify the wave operator as an element of it. To do so we represent the above short exact sequence in the physical Hilbert space H.
Following the developments of [13] we first consider the case K = C and let A, B be selfadjoint operators in H both with purely absolutely continuous spectrum equal to R and commutator given formally by [iA, B] = −1. We can then represent C0 R∪{+∞}; K ⋊γR faithfully as the norm closure E ′ in B(H)
 of the set of finite sums of the form ϕ1 (A) η1 (B)+. . .+ϕm (A) ηm (B)
where ϕi ∈ C0 R ∪ {+∞} and ηi ∈ C0 (R). We denote by J ′ the ideal obtained by choosing functions ϕi that vanish at {+∞}. Furthermore, we can represent K ⋊R faithfully in B(H) by elements of the form η(B) with η ∈ C0 (R). This algebra is denoted by C ′ .
In [13] position and momentum operators were chosen for A and B but we take A :=
i
− 2 (Q · ∇ + ∇ · Q) and B := 21 ln H0 . We refer to [14] for a thorough description of A in various
′
representations. Let
 us notice that a typical element of E is of the form ϕ(A) η(H0 ) with
ϕ ∈ C0 R ∪ {+∞} and η ∈ C0 (R+ ), the algebra of continuous functions on R+ that vanish at

2 (Sn−1 ) from the decomposition the origin and at infinity.
We shall now consider
K
=
K
L

H ∼
= L2 R+ ; L2 (Sn−1 ) in spherical coordinates. Since A and H0 are rotation invariant the presence of a larger K does not interfere with the above argument. Thus we set E := E ′ ⊗ K ,
J := E ′ ⊗ K and C := C ′ ⊗ K . These algebras are all represented in B(H), although C is a quotient of E . The surjection ev∞ becomes the map P∞ , where P∞ [T ] := T∞ , with T∞ uniquely
∗ )k → 0 as t → +∞, defined by the conditions kχ(A ≥ t) (T − T∞ )k → 0 and kχ(A ≥ t) (T ∗ − T∞
χ denoting the characteristic function. We easily observe that P∞ [ϕ(A) η(H0 )] = ϕ(+∞) η(H0 )

for any ϕ ∈ C0 R ∪ {+∞} and η ∈ C0 R+ ; K , where ϕ(+∞) is simply the value of the function ϕ at the point {+∞}. Let us summarise our findings:
Lemma 3.1. All three algebras of (6) are represented faithfully in H by J , E and C . In B(H)
the surjection ev∞ becomes P∞ .
Note that J is equal to the set of compact operators in H. For suitable potentials V , the operator S − 1 belongs to C [14, 15] and P is a compact operator. The key ingredient below is the use of Ω− to make the link between the K1 -class [S]1 of S and the K0 -class [P ]0 of P .

11

Theorem 3.2. Assume that Ω− − 1 belongs to E . Then S − 1 is an element of C , P belongs to
J and one has at the level of K-theory: ind[S]1 = − [P ]0 .

(7)

Proof. Let T ∈ E . Then T∞ = P∞ (T ) ∈ C satisfies kχ(A ≥ t)(T − T∞ )k → 0 as t → +∞.
Equivalently, kχ(A ≥ 0) [U (t)T U (t)∗ − T∞ ]k → 0 as t → +∞, since T∞ commutes with U (t) :=
i e 2 t ln H0 for all t ∈ R. It is then easily observed that s − limt→+∞ U (t) T U (t)∗ = T∞ . Now, if T
is replaced by Ω− − 1, the operator T∞ has to be equal to S − 1, since s − limt→+∞ U (t) Ω− U (t)∗
is equal to S. Indeed, this result directly follows from the intertwining relation of Ω− and the invariance principle [1, Theorem 7.1.4].
We thus have shown that Ω− − 1 is a preimage of S − 1 in E . It is well known that
Ω− Ω∗− = 1 − P and Ω∗− Ω− = 1. In particular Ω− is a partial isometry so that ind[S]1 =
[Ω− Ω∗− ]0 − [Ω∗− Ω− ]0 = −[P ]0 , see e.g. [26, Proposition 9.2.2].
Remark 3.3. It seems interesting that the condition Ω− − 1 ∈ E implies the finiteness of the set of eigenvalues of H. Another consequence of this hypothesis is that S(0) = 1, a result which is also not obvious. See [15, Section 5] for a detailed analysis of the behaviour of S(·) near the origin.
It is important to express the above condition on Ω− in a more traditional way, i.e. in terms of scattering conditions. The following lemma is based on an alternative description of the C ∗ algebra E . Its easy proof can be obtained by mimicking some developments given in Section 3.5
of [13]. We also use the convention of that reference, that is: if a symbol like T (∗) appears in a relation, it means that this relation has to hold for T and for its adjoint T ∗ .

Lemma 3.4. The operator Ω− − 1 belongs to E if and only if S(·) − 1 belongs to C0 R+ ; K
and the following conditions are satisfied:
(i) limε→0 kχ(H0 ≤ ε) (Ω− − 1)(∗) k = 0, and limε→+∞ kχ(H0 ≥ ε) (Ω− − 1)(∗) k = 0,
(ii) limt→−∞ kχ(A ≤ t) (Ω− − 1)(∗) k = 0, and limt→+∞ kχ(A ≥ t) (Ω− − S)(∗) k = 0.
Let us note that conditions (ii) can be rewritten as lim kχ(A ≤ 0) U (t) (Ω− − 1)(∗) U (t)∗ k = 0

t→−∞

and lim kχ(A ≥ 0) U (t) (Ω− − S)(∗) U (t)∗ k = 0.

t→+∞

3.3

The topological version of Levinson’s theorem


In the next statement, it is required that the map R+ ∋ λ 7→ S(λ) ∈ B L2 (Sn−1 ) is differentiable. We refer for example to [14, Theorem 3.6] for sufficient conditions on V for that purpose.
Trace class conditions on S(λ) − 1 for all λ ∈ R+ are common requirements [11]. Unfortunately, similar conditions on S ′ (λ) were much less studied in the literature. However, let us already mention that these technical conditions are going to be weaken in [17].
12

Theorem 3.5.
Let Ω− − 1 belong to E . Assume furthermore that the map

 R+ ∋ λ 7→ S(λ) ∈
B L2 (Sn−1 ) is differentiable, and that λ 7→ tr[S ′ (λ)] belongs to L1 R+ , dλ . Then the following equality holds:
Z ∞


dλ tr i(S(λ) − 1)∗ S ′ (λ) = 2π Tr[P ].
(8)
0

Proof. The boundary maps in K-theory of the exact sequence (6) are the inverses of the ConnesThom isomorphism (which here specialises to the Bott-isomorphism as the action in the quotient is trivial) and have a dual in cyclic cohomology [9], or rather on higher traces [10, 20], which gives rise to an equality between pairings which we first recall: Tr is a 0-trace on the ideal

2 (R) ⊗ K L2 (Sn−1 ) which we factor Tr = Tr′ ⊗ tr. Then tr
ˆ : K ⋊R → C,
K
L
C0 R; K ⋊γ R ∼
=
ˆ
ˆ
tr[a] = tr[a(0)] is a trace on the crossed product and (a, b) 7→ tr[aδ(b)] a 1-trace where [δ(b)](t) =
itb(t). With these ingredients
ˆ
tr[i(u
− 1)∗ δ(u)] = −2πTr[p] if

ind[u]1 = [p]0 ,

(9)

provided u is a representative of its K1 -class [u]1 on which the 1-trace can be evaluated. This
ˆ
is for instance the case if δ(u) is tr-traceclass.
To apply this to our situation, in which u is the unitary represented by the scattering operator and p is represented by the projection onto the bound states, we express δ and t̂r on U C U R∗ where U is the unitary from Section 3.1 diagonalising dλ
d
ˆ becomes and tr hypothesis implies the neccessary trace
H0 . Then δ becomes λ dλ
R+ λ tr. Our
R∞


class property so that the l.h.s. of (9) corresponds to 0 dλ tr i(S(λ)−1)∗ S ′ (λ) and the r.h.s. to
2π Tr[P ].
Remark 3.6. Expressions very similar to (8) already appeared in [8] and [12]. However, it seems that they did not attract the attention of the respective authors and that a formulation closer to (5) was preferred. One reason is that the operator {S(λ)∗ S ′ (λ)}λ∈R+ has a physical meaning: it represents the time delay of the system under consideration. We refer to [2] for more explanations and results on this operator.
Remark 3.7. At present our approach does not allow to say anything about a half-bound state but this will be remedied in [17]. We refer to [15], [23] or [24] for explanations on that concept and to [23] or [24] for corrections of Levinson’s theorem in the presence of such a 0-energy resonance.

3.4

Further prospects

We outline several improvements or extensions that ought to be carried out or seem natural in view of this note. We hope to express some of these in [17].
1. Our main hypothesis of Theorem 3.5, that Ω− − 1 belongs to the C ∗ -algebra E , is crucial and we have provided estimates in Lemma 3.4 which would guarantee it. Such estimates are rather difficult to obtain and we were not able to locate similar conditions in the literature. They clearly need to be addressed.
13

2. Similar results should hold for a more general operator H0 with absolutely continuous spectrum. In that case, the role of A would be played by an operator conjugate to H0 .
We refer to [1, Proposition 7.2.14] for the construction of such an operator in a general framework.
3. More general short range potentials or trace class perturbations can also be treated in a very similar way. By our initial hypothesis on V we have purposely eliminated positive eigenvalues of H, but it would be interesting to have a better understanding of their role with respect to Theorems 3.2 and 3.5.
4. In principle, Theorem 3.2 is stronger than Levinson’s theorem and one could therefore expect new topological relations from pairings with other cyclic cocycles. In the present setting these do not yet show up as the ranks of the K-groups are too small. But in more complicated scattering processes this could well be the case.
5. In the literature one finds also the so-called higher-order Levinson’s Theorems [7]. In the case n = 3 and under suitable hypotheses they take the form [7, equation 3.28]
Z ∞
X
 

dλ λN tr iS(λ)∗ S ′ (λ) − CN (λ) = 2π
eNj ,
0

j

where N is any natural number, CN are correction terms, and {ej } is the set of eigenvalues of H with multiplicities counted. The correction terms can be explicitly computed in terms of H0 and V [7] and we expect that they can be absorbed in a similar manner into the
S-operator as above.

4

Appendix on twisted crossed products

In this section, we start by recalling the definition of a twisted crossed product C ∗ -algebra borrowed from [21, Section 2]. We refer to the references quoted in that paper for a more general definition. Then, we consider the particular situation of the 2-cocycle defined by a magnetic field. Finally, a decomposition of the magnetic twisted crossed product as an iterated twisted crossed product is proved.
Let X be an abelian, second countable, locally compact group, and let C be an abelian C ∗ algebra, with its norm denoted by kf k and its involution by f ∗ , for any f ∈ C. Assume that there exists a group morphism α : X → Aut(C) from X to the group of automorphisms of C such that the map X ∋ x 7→ αx [f ] ∈ C is continuous for all f ∈ C. Assume also that there exists a strictly continuous normalized 2-cocycle ω on X with values in the unitary group of the multiplier algebra M(C) of C. We refer to [28] for the definition of the multiplier algebra, but recall that if
C ⊂ B(H), for some Hilbert space H, then M(C) = {a ∈ B(H) | af, f a ∈ C, ∀f ∈ C}. Since C is abelian, so is M(C). In other words, ω : X × X → M(C) satisfies the following conditions: For any x, y, z ∈ X and f ∈ C, (1) ω(x, y)∗ ω(x, y) = 1, (2) the map X × X ∋ (x, y) 7→ ω(x, y)f ∈ C

14

is continuous, (3) ω(x, 0) = ω(0, x) = 1, and (4) the 2-cocycle relation
ω(x, y) ω(x + y, z) = αx [ω(y, z)] ω(x, y + z) .

(10)

We have used in this relation that any automorphism of C extends uniquely to an automorphism of M(C). The quadruple (C, α, ω, X) is usually called an abelian twisted C ∗ -dynamical system.
Finally, we shall also assume the additional condition ω(x, −x) = 1, which holds in all the applications we have in mind.
Now, let κ ∈ [0, 1]; this additional parameter is convenient in order to relate our expressions with earlier results found in the literature. The special cases κ = 0 and κ = 1 are related with the right and the left quantisation respectively. Most of the time only the case κ = 0 is presented, but κ = 1/2 is prefered in quantum mechanics because of some additional symmetry properties. However, let us already mention that the following structures are isomorphic for different κ.
R
We consider the set L1 (X; C) endowed with the norm kF k1 := X kF (x)kdx for any F ∈
L1 (X; C), the multiplication
Z
ακ(y−x) [F (y)] α(1−κ)y [G(x − y)] α−κx [ω(y, x − y)] dy ,
(F ⋄ G)(x) :=
X

and the involution F ⋄ (x) := α(1−2κ)x [F (−x)∗ ].

Then the envelopping C ∗ -algebra of L1 (X; C)
endowed with these operations is called the twisted crossed product of C by X associated with the twisted actions (α, ω). We shall denote it by C ⋊ωα X, or simply C ⋊α X if ω ≡ 1.
Let us now assume that C is a C ∗ -algebra of bounded and uniformly continuous functions on
Rn , stable under translations. We also fix X := Rn and the action α of Rn on C is simply given by translations. Moreover, suppose that a continuous magnetic field on Rn is also present. Its components are denoted by {Bjk }nj,k=1 , and for any q, x, y ∈ Rn we define

ω B (q; x, y) := exp − iΓB hq, q + x, q + x + yi ,

where ΓB hq, q + q, q + x + yi is the flux of the magnetic field through the triangle defined by the points q, q + x and q + x + y. If all Bjk belong to C, then the map ω B : Rn × Rn ∋ (x, y) 7→
ω B (·; x, y) ∈ M(C) satisfies all the above conditions imposed on ω. Furthermore, the additional property ω B (x, −x) = 1 is always fulfilled. One may thus form the magnetic twisted crossed product C ∗ -algebra associated with the magnetic twisted actions (α, ω B ) and denote it simply n
by C ⋊B
α R .
n
We now show how the magnetic twisted crossed product C ⋊B
α R can be decomposed as an iterated twisted crossed product. The strategy is inspired from [25, Theorem 4.1] which deals with more general algebras C, 2-cocycles ω and groups X, but only in the special case
κ = 0. Let us first observe that if (C, α, ω, Rn ) is an abelian twisted C ∗ -dynamical system, then
(C, αk , ωk , Rn−1 ), with ωk (x, y) := ω (x, 0), (y, 0) for all x, y ∈ Rn−1 and αk the restriction of the action to Rn−1 , is also an abelian twisted C ∗ -dynamical system. For simplicity, we shall keep writing α and ω for αk and ωk , and ⋄, ⋄ for the multiplication and the involution in L1 (Rn−1 ; C).
Furthermore, we omit the superscript B in ω B in the following statement and in its proof.
15

Proposition 4.1. For any magnetic abelian twisted C ∗ -dynamical system (C, α, ω, Rn ) and any
κ ∈ [0, 1], there exits a continuous group morphism β from R to the group of automorphisms of
C ⋊ωα Rn−1 such that

C ⋊ωα Rn ∼
(11)
= C ⋊ωα Rn−1 ⋊β R .
Proof. In this proof, we consider the elements xk , yk of Rn−1 and the elements x⊥ , y⊥ of R. For
F ∈ L1 (Rn−1 ; C), let us set
βx⊥ [F ](xk ) := α−κ(xk ,0⊥ ) [✷(x⊥ , xk )] α(0k ,x⊥ ) [F (xk )]

∗
that for each q ∈ Rn with ✷(x⊥ , xk ) := ω (0k , x⊥ ), (xk , 0⊥ ) ω (xk, 0⊥ ), (0k , x⊥ ) . We may observe

B
the expression ✷(q; x⊥ , xk ) is equal to exp − iΓ✷ q; (0k , x⊥ ), (xk , 0⊥ ) , where the exponent is the flux of the magnetic field trough the square defined by the points q, q + (0k , x⊥ ), q + (xk , x⊥ )
and q + (xk , 0⊥ ).
(i) Let us first prove that β defines a continuous group morphism from R to Aut(C ⋊ωαRn−1 ).
It is obvious that βx⊥ [F ] belongs to L1 (Rn−1 ; C). By taking into account the relation (10) and n
the special property
ω(x,
2-cocycles, one also easily

 tx) = 1 ∀x ∈ R and t ∈ R of magnetic obtains that βx⊥ βy⊥ [F ] = βx⊥ +y⊥ [F ]. Furthermore, for G ∈ L1 (Rn−1 ; C) one has βx⊥ [F ⋄ G] =
βx⊥ [F ] ⋄ βx⊥ [G] if for all yk ∈ Rn−1 the following equality holds:
✷(x⊥ , xk ) = ✷(x⊥ , yk ) α(yk ,0⊥ ) [✷(x⊥ , xk − yk )] ω(yk , xk − yk ) α(0k ,x⊥ ) [ω(yk , xk − yk )]∗ .
But again, this can be verified with the help of relation (10). The same relation also leads to the equality βx⊥ [F ⋄ ] = (βx⊥ [F ])⋄ . Finally, the continuity of the map R ∋ x⊥ 7→ βx⊥ [F ] ∈ L1 (Rn−1 ; C)
can be proved by taking into account the strict continuity of ω and the continuity of the map
α : Rn → Aut(C). By a density argument, one completes the proof of the assertion. 
(ii) Let us now define the bijective map: Cc (Rn ; C) ∋ F 7→ F ′ ∈ Cc R; Cc (Rn−1 ; C) given by


∗ 
F ′ (xk ; x⊥ ) := α−κx ✷(κx⊥ , xk )∗ ω (xk , 0⊥ ), (0k , x⊥ ) F (xk , x⊥ ).

The multiplication in Cc R; Cc (Rn−1 ; C) is defined by
Z
′
′
dy⊥ βκ(y⊥ −x⊥ ) [F ′ (·; y⊥ )] ⋄ β(1−κ)y⊥ [G′ (·; x⊥ − y⊥ )]
(F ⋄β G )(·; x⊥ ) :=
R


⋄ 
and the involution is given by (F ′ )⋄β (·; x⊥ ) = β(1−2κ)x⊥ F ′ (·; −x⊥ ) , i.e. :


(F ′ )⋄β (xk ; x⊥ ) = α−κ(xk ,0⊥ ) ✷ (1 − 2κ)x⊥ , xk α(1−2κ)x [F ′ (−xk ; −x⊥ )∗ ].
The final step consists in verifying that F ′ ⋄β G′ is equal to (F ⋄ G)′ , and that [(F ′ )⋄β ] is equal to (F ⋄ )′ . These equalities can be checked without difficulty by taking into account the relation
(10) and the already mentioned property of magnetic 2-cocycles. A density argument completes the proof.

16

Acknowledgements
Serge Richard thanks the Swiss National Science Foundation and the european network: Quantum Spaces - Noncommutative Geometry for financial support.

References
[1] W.O. Amrein, A. Boutet de Monvel, V. Georgescu, C0 -groups, commutator methods and spectral theory of N-body Hamiltonians, Progress in Math. 135, Birkhäuser, 1996.
[2] W.O. Amrein, M.B. Cibils, Global and Eisenbud-Wigner time delay in scattering theory,
Helv. Phys. Acta 60 (1987), 481–500.
[3] W.O. Amrein, J.M. Jauch, K.B. Sinha, Scattering theory in quantum mecanics, W.A. Benjamin, 1977.
[4] J. Bellissard, A. van Elst, H. Schulz-Baldes, The noncommutative geometry of the quantum
Hall effect, J. Math. Phys. 35 (1994), 5373–5451.
[5] J. Bellissard, R. Benedetti, J-M. Gambaudo, Spaces of Tilings, Finite Telescopic Approximations and Gap-Labeling, Commun. Math. Phys. 261 (2006), 1–41.
[6] M. Benameur, H. Oyono-Oyono, Gap-labelling for quasi-crystals (proving a conjecture by
J. Bellissard), in: Operator Algebras and Mathematical Physics, Conference Proceedings:
Constanţa (Romania) July 2001, 11–22, Theta Foundation, 2003.
[7] D. Bollé, Higher-order Levinson’s theorems and the high-temperature expansion of the partition function, Ann. Phys. 121 (1979), 131–146.
[8] D. Bollé, T.A. Osborn, An extended Levinson’s theorem, J. Math. Phys. 18 (1977), 432–440.
[9] A. Connes, An analogue of the Thom isomorphism for crossed products of a C ∗ -algebra by an action of R, Adv. in Math. 39 (1981), 31–55.
[10] A. Connes, Cyclic cohomology and the transverse fundamental class of a foliation, in: Geometric methods in operator algebras, Kyoto, 1983, pp. 52–144, Pitman Res. Notes in Math.,
Longman, Harlow, 1986.
[11] E.B. Davies, Energy dependence of the scattering operator, Adv. App. Math. 1 (1980),
300–323.
[12] T. Dreyfus, The determinant of the scattering matrix and its relation to the number of eigenvalues, J. Math. Anal. Appl. 64 (1978), 114–134, and The number of states bound by non-central potentials, Helv. Phys. Acta 51 (1978), 321–329.

17

[13] V. Georgescu, A. Iftimovici, C ∗ -algebras of quantum Hamiltonians, in: Operator Algebras and Mathematical Physics, Conference Proceedings: Constanţa (Romania) July 2001, 123–
167, Theta Foundation, 2003.
[14] A. Jensen, Time-delay in potential scattering theory, Commun. Math. Phys. 82 (1981),
435–456.
[15] A. Jensen, T. Kato, Spectral properties of Schrödinger operators and time-decay of the wave functions, Du

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
