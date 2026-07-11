---
source: "https://arxiv.org/abs/hep-th/9707234v2"
title: "Variational Approach to Quantum Field Theory: Gaussian Approximation and the Perturbative Expansion around It"
author: "Jae Hyung Yee"
year: "1997"
publication: "arXiv preprint / hep-th"
download: "https://arxiv.org/pdf/hep-th/9707234v2"
pdf: "https://arxiv.org/pdf/hep-th/9707234v2"
captured_at: "2026-06-27T06:15:14Z"
updated_at: "2026-06-27T06:15:14Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács Theory of the Novel"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# Variational Approach to Quantum Field Theory: Gaussian Approximation and the Perturbative Expansion around It

- 著者: Jae Hyung Yee
- 年: 1997
- 掲載情報: arXiv preprint / hep-th
- 情報源: [arxiv](https://arxiv.org/abs/hep-th/9707234v2)
- ダウンロード: https://arxiv.org/pdf/hep-th/9707234v2
- PDF: https://arxiv.org/pdf/hep-th/9707234v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

The functional Schrodinger picture formulation of quantum field theory and the variational Gaussian approximation method based on the formulation are briefly reviewed. After presenting recent attempts to improve the variational approximation, we introduce a new systematic method based on the background field method, which enables one to compute the order-by-order correction terms to the Gaussian approximation of the effective action.

## PDF Text

arXiv:hep-th/9707234v2 21 Aug 1997

Variational Approach to Quantum Field Theory: Gaussian
Approximation and the Perturbative Expansion around Ita
Jae Hyung Yee
Department of Physics and Institute for Mathematical Sciences
Yonsei University, Seoul 120-749, Korea
E-mail: jhyee@phya.yonsei.ac.kr
The functional Schrödinger picture formulation of quantum field theory and the variational Gaussian approximation method based on the formulation are briefly reviewed.
After presenting recent attempts to improve the variational approximation, we introduce a new systematic method based on the background field method, which enables one to compute the order-by-order correction terms to the Gaussian approximation of the effective action.

a To appear in the proceedings of APCTP-ICTP Joint International Conference ’97 on Recent

Developments in Nonperturbative Quantum Field Theory, May 26-30, 1997, Seoul

1

1

Introduction

The variational approximation method has been proved to be a convenient nonperturbative approximation method for both the quantum mechanical and quantum field theoretic systems. The method, which was originally developed for a non-perturbative evaluation of quantum mechanical energy eigenvalues, has been extended to an action principle by Jackiw and Kerman1 , which enables one to compute the effective actions for time-dependent systems and also for quantum field theories. The usefulness of the method for quantum field theories, however, is limited by the fact that only the Gaussian functional and it’s obvious generalizations can be used as trial wave functionals, which is the reason why it is often called the Gaussian approximation. Attempts to improve the Gaussian approximation by using more general non-Gaussian trial wave functionals have not been as fruitful as expected, due to the difficulties in solving the variational gap equations2 . To overcome such limitations the authors of Ref.[3] has developed the so-called post-Gaussian expansion method, which enables one to compute the order-by-order correction terms to the Gaussian effective potentials for the time-independent systems.
It is the purpose of this note to give a brief review of the variational approximation method in quantum field theory, and to present a new systematic method of perturbative expansion around the Gaussian effective action based on the background field method, which is also applicable to the time-dependent systems.
The variational method in quantum mechanics is based on the fact that the ground state energy is the minimum of the expectation values of the Hamiltonian of the system. This can be written as a variational principle,
δΨ < Ψ|Ĥ(Q̂, P̂ ) − λ|Ψ >= 0,

(1)

where Ĥ is the Hamiltonian of the system and |Ψ > is an arbitrary state. If one starts from a trial wave function with variational parameters, then Eq.(1) provides a variational approximation of the energy eigenstates.
This simple variational method can be generalized to include the information on the time-evolution of the system, by choosing the state to satisfy the constraints,
< Ψ|Ψ > = 1
< Ψ|Q̂|Ψ > = q(t)

(2)

< Ψ|P̂ |Ψ > = p(t).
Minimizing the expectation value of the Hamiltonian with respect to such states,
δΨ < Ψ|Ĥ(Q̂, P̂ )|Ψ >= 0,

(3)

we obtain the effective Hamiltonian,
Hef f (q, p) =< Ψ0 |Ĥ(Q̂, P̂ )|Ψ0 >,

(4)

where |Ψ0 > is the solution of Eq.(3). This effective Hamiltonian provides many useful informations on the quantum system. If one minimizes Hef f with respect to
2

the variations of q(t) and p(t), one obtains the ground state energy eigenvalue of the system. If one solves the Hamilton’s equation, on the other hand,
∂Hef f
∂p
∂Hef f ṗ(t) = −
,
∂q q̇(t) =

(5)

the solution q and p give the quantum mechanical informations on < Ψ0 |Q̂|Ψ0 >
and < Ψ0 |P̂ |Ψ0 >, respectively.
If we set q̇ = 0 = p in Eq.(2), we obtain the effective potential,
Vef f (q) = Hef f (q, p)|p=0 .

(6)

If one starts from a trial wave function which satisfies Eq.(2), then one obtains the variational approximation of the effective Hamiltonian or the effective potential.
This simple argument is in fact a part of the quantum action principle in the
Schrödinger picture called the action principle of Dirac1,4 . This formulation can be generalized to the case of quantum field theory, which is the functional Schrödinger picture formulation of quantum field theory4 .
2

Functional Schrödinger picture representation of quantum field theory and the variational method.

2.1

Schrödinger picture representation of scalar field theory

The scalar field theory is defined by the Lagrangian density,
L=

1
(π φ̇ − φπ̇) − H[φ, π],
2

(7)

and the equal-time commutation relations
[φ(x), π(x′ )]x0 =x′0 = iδ(~x − ~x′ )
[φ(x), φ(x′ )]x0 =x′0 = 0 = [π(x), π(x′ )]x0 =x′0 ,

(8)

where H[φ, π] = 12 π 2 + V (φ) is the Hamiltonian density.
As in the quantum mechanical case, we choose the basis vectors for the Schrödinger representation to be the eigenstates of the field operator φ(x),
φ|ϕ(~x) >= ϕ(~x)|ϕ(~x) >,

(9)

which satisfy the orthogonality relation and the completeness relation,
< ϕ|ϕ′ >= δ[ϕ(~x) − ϕ′ (~x′ )]
Z
Dϕ|ϕ >< ϕ| = I.
3

(10)

Then the state of the physical system is represented by the component, Ψ[ϕ, t] =
< ϕ|Ψ >, of the expansion,
Z
|Ψ >= Dϕ|ϕ >< ϕ|Ψ >,
(11)
and the operators are represented by the matrix elements,
< ϕ|φ|ϕ′ > = ϕ(~x)δ[ϕ − ϕ′ ]
δ
< ϕ|π|ϕ′ > = −i
δ[ϕ − ϕ′ ],
δϕ(~x)

(12)

which are obtained from the matrix elements of Eq.(8). The time-evolution of the system is then determined by the functional Schrödinger equation, i

d
Ψ[ϕ, t] = < ϕ|H[φ, π]|ϕ′ >< ϕ′ |Ψ >
dt
δ
= H[ϕ, −i ]Ψ[ϕ, t].
δϕ

(13)

This is a simple generalization of the Schrödinger picture representation of quantum mechanics to the case of the systems with infinite degrees of freedom.
2.2

Free Scalar field theory

To obtain physical information in the Schrödinger picture formulation one has to solve the functional Schrödinger equation (13). As an example of solving the equation of motion, we consider the free scalar field theory described by the Hamiltonian density,
1
1
1
(14)
H = π 2 + (∇φ)2 + m2 φ2 .
2
2
2
By using the representation (12), the Hamiltonian of the system can be written as
Z
1
δ
δ
H=
− ϕ(~x)(∇2 − m2 )ϕ(~x)].
(15)
d3 x[−
2
δϕ(~x) δϕ(~x)
Since the Hamiltonian has no explicit time-dependence, there exist stationary states satisfying the energy eigenvalue equation,
H[ϕ, −i

δ
]Ψ[ϕ] = EΨ[ϕ],
δϕ

(16)

where E is the energy eigenvalue. Since the Hamiltonian (15) is quadratic in ϕ(~x), the Gaussian functional
R 3 3 ′
′
′
1
Ψ[ϕ] = e− 2 d xd x ϕ(~x)G(~x,~x )ϕ(~x ) ,
(17)

is a solution of (16). Substituting (17) into (16), we obtain.

1
[G(~x, ~x′ ) − ϕ(~x)G2 (~x, ~x′ )ϕ(~x′ ) + ϕ(~x)(−∇2 + m2 )ϕ(~x)]Ψ[ϕ] = EΨ[ϕ],
2
4

(18)

where the repeated arguments imply the integration over the variables. For this equation to be true for all eigenvalue ϕ(~x), the following relations must be satisfied,
Z
1
1
E = T rG =
d3 xG(~x, ~x)
(19)
2
2
G2 (~x, ~x′ ) = (−∇2 + m2 )δ(~x − ~x′ ).
Solving the second equation of (19) via Fourier transformation, one finds the kernel of the Gaussian functional (17) is given by
Z
′
1
1
G−1 (~x, ~x′ ) =
d3 pei~p·(~x−~x ) p
,
(20)
3
(2π)
p~2 + m2

which is the propagator of the free scalar field theory at equal-time. From this the eigenvalue E is found to be
Z
Z
p
1 1
3
d x
d3 p p~2 + m2 ,
(21)
E=
3
2 (2π)

which is the well-known zero-point energy of the free scalar field theory.
We thus find that the ground state of the free scalar field theory is described by the Gaussian wave functional with the kernel given by the inverse propagator at equal-time. One can construct all the eigenstates of the Hamiltonian by considering the creation and annihilation operators. They are given by the Hermite polynomial functionals multiplied by the ground(vacuum) state5 .
For the interacting theory, however, we cannot solve the functional Schrödinger equation (13) in closed form. One thus has to use an approximation method to find solutions to the equation of motion. A natural way to evaluate an approximate solution to (13) in the Schrödinger picture is to use the variational method, which is the subject of the next subsection.
2.3

Gaussian approximation of φ4 theory

To illustrate the variational approximation method in quantum field theory, we now consider the scalar φ4 theory described by the Hamiltonian density,
H=

1 2 1 ~ 2 1 2 2 λ 4
π + (∇φ) + µ φ + φ ,
2
2
2
4!

(22)

and compute the Gaussian approximation of the effective potential6 . To do this we take as a trial wave functional the solution of a free field theory discussed in the last subsection, namely, the Gaussian wave functional,
1

Ψ[ϕ, t] = N e− 4 [ϕ(~x)−φ̂]G

−1

(~
x,~
x′ ;t)[ϕ(~
x′ )−φ̂]+ h̄i π̂(~
x,t)[ϕ(~
x)−φ̂]

(23)

which satisfies the constraints,
<Ψ|Ψ> = 1

< Ψ | φ | Ψ > = φ̂(~x, t)

< Ψ | π | Ψ > = π̂(~x, t).
5

(24)

Since the effective potential is the effective Hamiltonian with vanishing momenta, we set π̂ equal to zero and consider the space-time independent φ̂. Then the expectation value of the Hamiltonian with respect to the trial state(23) is given by

Z
1
1
3
<H> =
d x V (φ̂) + h̄[ G−1 (~x, ~x) − ∇2 G(~x, ~x′ )|~x=~x′
8
2

2
h̄
+ V (2) (φ̂)G(~x, ~x)] + V (4) (φ̂)(G(~x, ~x))2 ,
(25)
8
2

λ 4
φ̂ is the classical potential and V (n) (φ̂) = ( ddφ̂ )n V (φ̂).
where V (φ̂) = µ2 φ̂2 + 4!
The variational equation,
δG < H >= 0,
(26)

yields the variational gap equation,
1 −2
h̄
G (~x, ~x′ ) = [−∇2 + V (2) (φ̂) + V (4) (φ̂)G(~x, ~x)]δ(~x − ~x′ ).
4
2

(27)

Substituting(27) into (25) one obtains the effective potential,
Vef f (φ̂) =

h̄2
1 2 2 λ 4 h̄ −1
µ φ̂ + φ̂ + G (~x, ~x) − G(~x, ~x)G(~x, ~x).
2
4!
4
8

(28)

The first three terms Eq.(28) constitute the one-loop effective potential, and the last term contains the higher-loop contribution, which shows the non-perturbative nature of the approximation.
Using the Fourier transformation method we can solve the gap equation(27):
Z
d3 p i~p·(~x−~x′ )
1
p
G(~x − ~x′ ) =
,
(29)
3e
2
(2π)
p~ + m2

where the effective mass m is defined by m2 = µ2 +

λ 2 λ
φ̂ + h̄
2
2

Z

d3 p

1
p
.
2
(2π)
p~ + m2
3

Substituting these results into (28) yields the effective potential:

2
µ2 2 λ 4
λ 2
1
m2
2
2
Vef f (φ̂) =
m − µ − φ̂
φ̂ + φ̂ −
+
I1
4
4!
2λ
2
2


m2
1
m4
m4
, ln
−
I2 (M ) +
−
2
2
4
64π
M
2

(30)

(31)

where

λ
m2
λ
λm2
λ
ln 2 ,
I1 + φ̂2 − m2 I2 (M ) +
2
2
2
2
32π
M
and, I1 and I2 (M ) are given by m2 = µ2 +

I1 =

Z

d3 p

1
, p|
(2π)3 2|~

I2 (M ) =

Z

6

d3 p
(2π)3

!
1
1
.
− p
2|~
p| 2 p~2 + µ2

(32)

(33)

The effective potential (31) has two divergences, I1 and I2 (M ), and thus we need to renormalize the theory. The renormalization conditions can be written as
[6], dVef f
µ2R
d2 Vef f
1
=
,
,
(34)
=
2
2
dm2 m0
λR
3λ
R
d(m )
m0

which, in our case, yields
µ2R
µ2
1
=
+ I1
λR
λ
2
1 1
1
=
+ I2 (M ).
λR
λ 2

(35)

Substituting (35) into (31), we finally obtain the renormalized effective potential,


m4
m2
m2 2 µ2R 2
m4
1
Vef f (φ̂) = −
,
(36)
ln 2 −
+
φ̂ +
m +
2λR
2
λR
64π 2
M
2
where

λR 2
λR
m2
φ̂ +
ln 2 .
(37)
2
2
32π
M
This result is qualitatively equivalent to the large-N approximation (where N is the number of the scalar field components)7 , and shows the non-perturbative nature of the approximation. The result (36) can be used to study the quantum nature of the theory, such as the symmetry breaking phenomena6 . The method can be generalized to the case of the field theories in curved space8 , to the finite temperature quantum field theory9 , and also to the systems out of thermal equilibrium10 .
m2 = µ2R +

2.4

Fermionic theory

Finding the Schrödinger picture representation of fermion fields is non-trivial since the equal-time anti-commutation relation, in Hermitian representation,
{ψα (x), ψβ (x′ )}x0 =x0 ′ = δαβ δ(~x − ~x′ ),

(38)

implies that the field operators ψ’s are their own momentum conjugates. In the non-Hermitian representation, on the other hand, the anti-commutation relation is given by
(39)
{ψα (x), ψβ† (x′ )}x0 =x0 ′ = δαβ δ(~x − ~x′ ), but now the eigenstates of field operators do not form an orthonormal basis. One can show, however, that there exists a one parameter family of equivalent representations of fermion fields4 . Floreanini and Jackiw11 , in particular, have found that the representation,
1
δ
ψ = √ [θ(~x) + †
]
δθ (~x)
2
δ
1
],
ψ † = √ [θ† (~x) +
δθ(~x)
2
7

(40)

where θ(~x) is a complex Grassmann variable, satisfies the anti-commutation relation
(39).
Using this representation one can follow the same procedure as in the case of the scalar field theory to evaluate the Gaussian approximation of the Fermionic field theories.
This method has been applied to the cases of (1+1)-dimensional and (2+1)dimensional Gross-Neveu models12 , and found that the results are qualitatively equivalent to those of the large-N approximation. It has also been applied to the
(2+1)-dimensional Thirring model, and in this case the Gaussian result has been found to be qualitatively better than the large-N approximation13 . The Gaussian effective potential shows the existence of the symmetry breaking in the theory, while the large-N does not. The method has also been found to be effective in determining the existence of the bound states in the theory14 .
3

Beyond Gaussian approximation?

We have shown that the Gaussian approximation based on the functional Schrödinger picture formulation is a simple generalization of the quantum mechanical variational method, and provides a convenient non-perturbative approximation method for various quantum field theories including the systems out of thermal equilibrium.
However, the method has a serious drawback in that, unlike the quantum mechanical cases, one can compute the expectation value of the quantum field theoretic
Hamiltonian only for the Gaussian wave functional and its simple generalizations.
If one could develop a systematic method to compute the correction terms to this approximation, therefore, the variational method would become a very powerful technique in obtaining physical information from quantum field theories. In fact there have been some attempts in this direction recently.
One of such attempts, known as the post-Gaussian approximation method or the variational perturbation theory3 , is based on the observation that the Gaussian wave functional
1

Ψ0 = e− 4 ϕ̃(~x)G

−1

(~
x−~
x′ )ϕ̃(~
x′ )

,

ϕ̃ = ϕ − φ̂

is a vacuum state for a free theory with mass given by
Z
d3 p
λ 2 λ
2
2
m = µ + φ̂ +
g(p),
2
2
(2π)3

(41)

(42)

where g(p) is determined by the variational gap equation (29) and (30).
One can construct all the eigenstates of this free theory:
Ψn [ϕ̃] = Hn [ϕ̃]Ψ0 [ϕ̃] ,

n = 0, 1, 2, · · ·

(43)

where Hn [ϕ̃] is the Hermite polynomial functional. One splits the Hamiltonian as
H = H0 + H I

(44)

where H0 is the free Hamiltonian with the mass m given by Eq.(42) and HI is the remaining part of H, and then applies the well known perturbation theory with the states (43) taken to be the basis for the perturbation.
8

The effective potential is the expectation value of the Hamiltonian with the constraint
< Ψ0 |φ|Ψ0 >= φ̂,
(45)
where |Ψ0 > is the ground state of the system.
Taking the eigenstates (43) to be the basis for perturbative expansion, the effective potential can be written as

n
∞
X
1
Vef f (φ̂) = E0 (φ̂) +
< Ψ0 |ΨI
|Ψ0 >,
(46)
HI
E0 − H0
n=0
where the first term is the Gaussian effective potential and the second term represents the perturbative correction terms3 . This method has been used to compute the first non-trivial contribution to the Gaussian effective potential for the (0+1)dimensional φ4 theory by Jeon15 and the (1+1)- and (2+1)-dimensional φ4 theories by Cea and Tedesco3 .
This method provides a conceptually simple procedure to compute the orderby-order correction terms to the Gaussian effective potential for time-independent systems. For the time-dependent systems such as those out of thermal equilibrium10 , however, we need a method that can deal with the effective action, which contains information on the time-evolution of the system. It is the subject of the next section to introduce a new systematic expansion method for effective action, which was developed by Geon Hyoung Lee16 .
4

The Background field method

We now introduce a systematic method which enables one to compute the order by order correction terms to the Gaussian approximation of the effective action, that contains the information on the time-evolution of the system. This new method is based on the background field method. So we will start to give a brief review on the background field method17 .
4.1

Background field method

The generating functional for the connected Green’s functions, W [J], is defined by the path integral,
Z
eiW [J] = N

DφeiS[φ]+iJφ ,

(47)

R
where S[φ] is the classical action, and the integral convention, Jφ ≡ dn xJ(x)φ(x)
in n-dimensional space-time, is used in the exponent. The effective action Γ[φ] is then defined by the Legendre transformation,
Γ[ϕ] = W [J] − Jϕ,

(48)

where ϕ(x) is the vacuum expectation value of the field operator,
ϕ(x) =< φ(x) >J =
9

δ
W [J].
δJ(x)

(49)

The background field method is one of the most effective ways of computing the effective action in the loop expansion, and is obtained by introducing the background field B(x) to the classical action:
S[φ] → S[φ + B].

(50)

This defines a new generating functional,
Z
iW̃ [J,B]
e
= N DφeiS[φ+B]+iJφ
Z
′
′
= N Dφ′ eiS[φ ]+iJ[φ −B] .

(51)

Eg.(51) implies that this new generating functional is related to the old one by
W̃ [J, B] = W [J] − JB,

(52)

and the vacuum expectation value of the field operator in the presence of the background field is given by,
δ W̃
ϕ̃ =
= ϕ − B.
(53)
δJ
Now consider the case when the vacuum expectation value of the field operator in the presense of the background field, ϕ̃, vanishes:
δ W̃
= ϕ̃ = 0.
δJ

(54)

This determines the configuration of J which extremizes the generating functional
W̃ [J, B]:
J = J[B].
(55)
Then one can easily show from Eq.(52) that the extremum value of the generating functional W̃ becomes the effective action:
W̃ [J, B]|ϕ̃=0 = W [J(B)] − J(B)B = Γ[B].

(56)

The simpler way to compute the effective action Γ[B] is to consider the Legendre transformation of W̃ [J, B]:
Γ̃[ϕ̃, B] = W̃ [J, B] − J ϕ̃,

(57)

which generates the one ϕ̃-particle irreducible diagrams in the presence of the background field. If we set ϕ̃ = 0 in Eq.(57), we find that
Γ̃[ϕ̃, B]

ϕ̃=0

= W̃ [J(B), B] − J ϕ̃

ϕ̃=0

= Γ[B].

(58)

This implies that the effective action Γ[B] consists of one ϕ̃-particle irreducible bubble diagrams with no external ϕ̃-lines. Thus one can write the effective action as
Z
iΓ[B]
iW̃ [J,B]
e
=e
= N DφeiS[φ+B]+iJφ
.
(59)
ϕ̃=0

ϕ̃=0

10

To illustrate how the method facilitates the computation of the effective action, we now consider the scalar φ4 theory as an example. The Lagrangian density of φ4
theory can be written as
1
λ
1
λ
L(φ) = − φ(∂ 2 + µ2 )φ − φ4 ≡ − φ(iD−1 )φ − φ4 ,
2
4!
2
4!

(60)

where D is the free propagator defined by
D ≡ i(∂ 2 + µ2 )−1 .

(61)

The shifted Lagrangian is then given by
1
λ
λ
1
λ
λ
L(φ + B) = −i BD−1 B − B 4 − i(BD−1 − i B 3 )φ − φiK −1 φ − Bφ3 − φ4 ,
2
4!
3!
2
3!
4!
(62)
where K is the new propagator,
K = i(∂ 2 + µ2 +

λ 2 −1
B ) ,
2

(63)

which includes a part of the interaction effect through the term proportional to
B 2 . The generating functional in the presence of the background field can then be written as,
1

1

eiW̃ [J,B] = N det(−K) 2 e 2 BD

−1

4
3
3
4
δ
iλ
δ
iλ
δ
B− iλ
(BD −1 − iλ
4! B
3! B ) iδJ − 3! B( iδJ ) − 4! ( iδJ )

e

1

e 2 JKJ .
(64)

The effective action for the scalar φ4 theory now becomes eiΓ[B] = eiW̃ [J,B]

1

ϕ̃=0

= N e 2 BD

−1

4
B− iλ
4! B

1

det(−K) 2 I0 [B],

(65)

where
I0 [B] = e(BD

−1

3
δ
− iλ
3! B ) iδJ

exp[−

δ 3 iλ δ 4 1 JKJ
iλ
B(
) − (
) ]e 2
.
3!
iδJ
4! iδJ
ϕ̃=0

(66)

The coefficient of I0 [B] in Eq.(65) represents the one-loop effective action, and
I0 [B] represents the higher loop contributions to the effective action. The first factor of Eq.(66) constitutes a shift operator, which does not change the extremum value of W̃ [J, B]. Thus this factor, which generates the tadpole diagrams, does not contribute to the effective action , and we can neglect it . One can therefore write
I0 [B] as
δ 3 iλ δ 4 21 JKJ
iλ
(67)
) − (
) ]e
I0 [B] = exp[− B(
3! iδJ
4! iδJ
ϕ̃=0
which implies that the higher loop contributions to the effective action consist of one ϕ̃-irreducible bubble diagrams with no external ϕ̃-lines. One can compute the higher loop contributions by the Feynman rules:

11

= K = ∂ 2 +µ2i+ λ B 2

: propagator

q

= −iλB

: 3-vertex

q

= −iλ

: 4-vertex

2

One can easily show that the two-loop contribution to the effective action consists of two bubble diagrams with no external lines17 :

-i λ

Γ(2) [B] =
K

-i λB
K

-i λ B .

(68)

K

Comparing this with the conventional method of computing the effective action, one can easily convince oneself that the background field method provides a very simple and effective way to compute the effective action in the loop expansion.
4.2

Expansion around the Gaussian effective action

The reason why the procedure of computing the effective action in the background field method is simplified is that the propagator used in the Feynman rules, Eq.(63), includes a part of the interaction effect through the background field B(x), and the generating functional has been rearranged in such a way that the zeroth order term in the expansion, the coefficient of I0 [B] in Eq.(65), is the one-loop effective action.
We now want to proceed one step further and rearrange the effective action in such a way that the propagator used in the expansion would include the nonperturbative effect and the zeroth order contribution would be the result of the non-perturbative approximation, such as the result of the Gaussian approximation.
Note that the Gaussian approximation consists of the cactus type diagrams where an internal line comes out of a point in space-time and goes back to the same point, such as the first term of Eq.(68). So we need to extract such diagrams out of the I0 [B] part of the effective action. To do this we examine the functional
1
derivatives of the generating functional e 2 JGJ with respect to J(x) :
1
1
δ 1 Jy Gyz Jz
= Gxy Jy e 2 Jy Gyz Jz ≡ (GJ)x e 2 JGJ
e2
δJx
1
δ 2 21 Jy Gyz Jz
= ((GJ)2x + Gxx )e 2 JGJ
e
2
δJx

12

(69)
(70)

1
δ 3 21 Jy Gyz Jz
= ((GJ)3x + 3Gxx (GJ)x )e 2 JGJ
e
δJx3
1
δ 4 1 Jy Gyz Jz
= ((GJ)4x + 6Gxx (GJ)2x + 3G2xx )e 2 JGJ , e2
δJx4

(71)
(72)

which would appear in the expansion of I0 [B]. Note that these derivatives, which generates 2-loop and higher order contributions to Γ[B], generate the cactus type diagrams as can be seen in Eqs.(71) and (72).
To extract such diagrams out of I0 [B], we define new functional derivatives so that the following relations satisfy:
1
δ 3 ′ 1 JGJ
= (GJ)3x e 2 JGJ
) e2
δJx3
1
1
δ4
( 4 )′ e 2 JGJ = (GJ)4x e 2 JGJ .
δJx

(

(73)
(74)

Comparing Eqs.(71)-(72) and Eqs.(73)-(74), we find that the primed derivatives must be defined by
δ
δ3
δ3 ′
)
≡
− 3Gxx
3
3
δJx
δJx
δJx
δ4 ′
δ2
δ4
( 4) ≡
−
6G
+ 3G2xx .
xx
δJx
δJx3
δJx2

(75)

(

(76)
1

Note that the primed derivatives operated on the functional e 2 JGJ do not generate the cactus type diagrams that contribute to the Gaussian effective action. This implies that by using the primed derivatives one may in fact rewrite I0 [B] in such a way that all the cactus type diagrams would be included in the zeroth order term in the expansion.
To do this we need to find a new Green’s function G such that
λB

δ

3

iλ

δ

4

1

δ

λB

δ3

′

iλ

δ4

′

1

e 3! ( δJ ) − 4! ( δJ ) e 2 JKJ = N ′ eA δJ e 3! ( δJ 3 ) − 4! ( δJ 4 ) e 2 JGJ ,

(77)

′

where a normalization constant N and a function A are to be determined. Using the definitions (75) and (76), one easily finds G, N ′ and A that satisfy Eq.(77):
1

N ′ = det(−G) 2 det(−G−1 +

iλ 2
1
iλ
Gxx ) 2 e 8 Gxx ,
2

(78)

λ
BGxx ,
(79)
2
iλ
−1
Gxx δxy
G−1
xy = Kxy +
2
iλ
iλ
−1
= Dxy
− Bx2 δxy + Gxx δxy .
(80)
2
2
Using these results we can write the generating functional in the presence of the background field as
A =

1

1

eiW̃ [J,B] = det(−G) 2 e 2 BD
× e(−iBD

−1

−1

4
iλ 2
B− iλ
4! B + 8 Gxx

λ
δ
− 3!
B3 + λ
2 BGxx ) δJ

13

λB

e 3!

3

4

δJx

δJx

′
δ
( δ 3 )′ − iλ
4)
4! (

1

e 2 JGJ .

(81)

The effective action can thus be written as eiΓ[B] = eW̃ [J,B]

ϕ̃=0
1

1

= det(−G) 2 e 2 BD

−1

4
iλ 2
B− iλ
4! B + 8 Gxx

I[B],

(82)

where


λ
λ
δ
I[B] = exp (−iBD−1 − B 3 + BGxx )
3!
2
δJ
3
4
λB δ ′ iλ δ ′ 1 JGJ
(
) − ( 4 ) ]e 2
× exp[
3! δJx3
4! δJx
ϕ̃=0
= exp[

λB δ 3 ′ iλ δ 4 ′ 1 JGJ
.
(
) − ( 4 ) ]e 2
3! δJx3
4! δJx
ϕ̃=0

(83)

In Eq.(83) we have used the fact that the shift operator, the tadpole part of the first line, does not affect the extremum value. Note that Eq.(80), which determines the Green’s function Gxy , is exactly the variational gap equation written in the
Minkowski space notation, and therefore the coefficient of I[B] in Eq.(82) gives rise to the Gaussian effective action as we have alluded to above. Since I[B] can be expanded as a power series in λ as is done in the conventional background field method, we have the perturbative expansion of the effective action around the
Gaussian approximation.
The structure of I[B] is the same as that of I0 [B] in Eq.(67), except that the propagator is now replaced by Gxy and the functional derivatives are replaced by the primed derivatives. Thus the I[B] contribution to the effective action consists of one ϕ̃-particle irreducible bubble diagrams with no external ϕ̃-lines, and without the cactus type diagrams. The contributions from I[B] can therefore be evaluated by the Feynman rules:

r

r

= K = ∂ 2 +µ2 + λ iB 2 − λ G
2

2

xx

: propagator

s

= −iλB

: 3-vertex

s

= −iλ

: 4-vertex

14

We have thus established the systematic method of computing the order-byorder correction terms to the Gaussian effective action.
One can easily show that the λ2 order contribution of I[B], for example, consists of the diagrams,
G

G

I (2) [B] = -i λΒ

-i λ

-i λB

-i λ .

(84)

Note that the first diagram of Eq.(68) is a cactus type diagram and does not appear in the expansion of I[B].
4.3

Effective potential

If we consider the space-time independent background field in the above formulation, we obtain the effective potential defined by
Γ[B]
Vef f (B) ≡ − R n .
d x

(85)

Since Gxy is a function of (x − y) for a constant B, one can solve Eq.(80) by Fourier transformation,
Z
dn p g(p)eip(x−y) .
(86)
G(x − y) ≡
(2π)n
The Green’s function Gxy in the momentum space can then be written as g(p) =
where m2G ≡ µ2 +

1
, i(p2 − m2G )

λ 2 λ
B −
2
2

Z

dn p g(p).
(2π)n

(87)

(88)

Note that Eq.(87) is the variational gap equation of Gaussian approximation written in the Minkowski notation.
Up to the λ2 order contribution of I[B], the effective potential is therefore given by
Vef f (B) = VG + VP ,
(89)
where
VG ≡

√
µ2 2 λ 4
1
λ
B + B −
(m2 − µ2 − B 2 )2 + iTr log −G,
2
4!
2λ
2

and

(90)

λ2 2 3
λ2
B G − i G4 .
(91)
12
48
VG of (90) is the Gaussian effective potential and VP of (91) is the first non-trivial contribution from I[B], which is represented by the Feynman diagrams shown in
VP ≡ i

15

(84). This result is the same as that of Cea and Tedesco3 computed from the variational perturbation theory mentioned earlier.
To illustrate the effectiveness of the perturbative correction method, we now consider the (0 + 1)-dimensional case, which is the quantum mechanical anharmonic oscillator. It is convenient to write the Green’s function in Euclidean space which is defined by p2 = −p2E , dn p = idn pE .
(92)
Then G(x, x) in Eq.(88) can be written as
Z

dp
1
=−
2π i(p2 − m2 )

Z

dpE
1
1 1
= −√ √ ,
2π p2E + m2
2 m2

(93)

λ 2
λ
B + √ .
2
4 m2

(94)

and the gap equation becomes m2 = µ2 +

The effective potential up to the λ2 order of I[B] becomes
Vef f (B) = VG + VP ,

(95)

where
µ2 2 λ 4
1
λ
VG =
B + B −
(m2 − µ2 − B 2 )2 +
2
4!
2λ
2
and
VP = −

√
m2
,
2

λ2
λ2
2
√
.
B
−
144m4
1536( m2 )5

(96)

(97)

We can compare this result with other approximation methods. For the comparison of the effectiveness of various approximation methods it is convenient to compare the ground state energy eigenvalue of the anharmonic oscillator, which is defined by the minimum value of Vef f . The table below shows the errors of various approximation methods compared to the numerical results18 . This shows that our
Table 1: Errors of E0 (µ2 = 1) compared to the numerical result

λ
0.24
24

1-loop
1.4%
38%

2-loop
0.05%
56%

Gaussian Approx.
0.006%
1.1%

Improved Gaussian
0.0003%
0.37%

method greatly improves the Gaussian approximation even at the first non-trivial correction level, for all values of the coupling constant. Note that, for strong coupling, the Gaussian and the improved Gaussian approximation display much better results than the loop expansion methods, which shows the non-perturbative nature of the approximation methods.
16

5

Conclusion

We have shown that the functional Schrödinger picture formulation of quantum field theory provides a conceptually simple non-perturbative approximation method, namely the Gaussian approximation. It has been proved to be a reliable method to obtain non-perturbative informations on the symmetry breaking phenomena, existence of bound states and time-evolution of the time-dependent systems with relatively little computational difficulties. The method, however, has a serious drawback in that it does not provide a systematic procedure to improve the approximation.
We have shown that the background field method provides a solution to this limitation of the variational method. We were able to rearrange the effective action in such a way that the zeroth order contribution in the expansion of the effective action consists of the Gaussian effective action. This provides a new systematic expansion method which enables one to compute the order-by-order correction terms to the Gaussian effective action.
The effective potentials for the (0 + 1), (1 + 1) and (2 + 1) dimensional φ4
theory computed by the new method agree with those of the variational perturbation theory. It has been shown that the perturbative correction greatly improves the
Gaussian approximation even at the first non-trivial correction level. Our method can also be applied to the time-dependent systems and can easily be extended to the cases of fermionic and gauge field theories.
Although the effective action expanded around the Gaussian approximation for the lower-dimensional φ4 theories can easily be shown to be renormalizable, one has to study the renormalization procedure systematically for the (3 + 1) dimensional case.

17

Acknowledgements
This work was supported in part by the Korea Science and Engineering Foundation under Grant No.95-0701-04-01-3 and 065-0200-001-2, the Center for Theoretical
Physics (SNU), and the Basic Science Research Institute Program, Ministy of Education, under project No. BSRI-96-2425.

18

References
1. R. Jackiw and A. Kerman, Phys. Lett. A71, 158 (1979).
2. L. Polley and D. E. L. Pottinger, eds., Variational calculations in quantum field theory, World Scientific, Singapore (1987); L. Polley and U. Ritschel,
Phys. Lett. B221, 44 (1989); U. Ritschel, Phys. Lett. B227, 251 (1989); Z.
Phys. 47, 457 (1990); H. Verschelde and M. Coppens, Phys. Lett. B287, 133
(1992).
3. A.Okopińska, Phys. Rev. D 35, 1835 (1987); Phys. Rev. D 38, 2507 (1988);
Ann. Phys.(N.Y.) 228, 19 (1993). I. Stancu and P. M. Stevenson, Phys. Rev.
D 42, 2710 (1990); P. Cea, Phys. Lett. B236, 191 (1990); I. Stancu, Phys.
Rev. D 43, 1283 (1991); P. Cea and L. Tedesco, Phys. Lett. B335, 423
(1994);Phys. Rev. D 55, 4967 (1997).
4. K. Symanzik, Nucl. Phys. B190, 1 (1981); M. Lüscher, Nucl. Phys. B254,
52 (1985); J. H. Yee, in Recent Developments in Field Theory, Proc. 10th symposium on Theoretical Physics, pp210, ed. by J. E. Kim, Min Eum
Sa, Seoul(1992); R. Jackiw, Diverse Topics in Theoretical and Mathematical Physics, World Scientific, Singapore (1995) and references therein.
5. D. V. Long and G. M. Shore, Swansea preprint, hep-th/9605004.
6. F. Cooper and E. Mottola, Phys. Rev. D 36, 3114 (1987); S. -Y. Pi and M.
Samiullah, Phys. Rev. D 36, 3128 (1987); P. M. Stevenson, Phys. Rev. D 30,
1714 (1984); Phys. Rev. D 32, 1389 (1985); T. Barnes and G. I. Ghandour,
Phys. Rev. D 22, 924 (1980).
7. S. Coleman, R. Jackiw and D. Politzer, Phys. Rev. D 10, 2491 (1974);
8. S. K. Kim, W. Namgung, K. S. Soh and J. H. Yee, Phys. Rev. D 38, 1853
(1988); J. Guven, G. Lieberman and C. T. Hill, Phys. Rev. D 39, 438 (1989);
O. Éboli, S. -Y. Pi, and M. Samiullah, Ann. Phys. 193, 102 (1989).
9. H. -J. Lee, K. Na and J. H. Yee, Phys. Rev. D 51, 3125 (1995).
10. O. Éboli, R. Jackiw and S. -Y. Pi, Phys. Rev. D 37, 3557 (1988).
11. R. Floreanini and R. Jackiw, Phys. Rev. D 37, 2206 (1988).
12. S. K. Kim, J. Yang, K. S. Soh and J. H. Yee, Phys. Rev. D 40, 2647 (1989);
S. K. Kim, K. S. Soh and J. H. Yee, Phys. Rev. D 41, 1345 (1990).
13. S. Hyun, G. H. Lee and J. H. Yee, Phys. Rev. D 50, 6542 (1994).
14. S. K. Kim, W. Namgung, K. S. Soh and J. H. Yee, Ann. Phys. 214, 142
(1992).
15. K. J. Jeon, Master’s Thesis, Yonsei University (1996).
16. G. H. Lee, Ph.D. Thesis, Yonsei University (1997); G. H. Lee and J. H. Yee , preprint(1997), hep-th/9707035.
17. B. S. DeWitt, Phys. Rev. 162, 1195 (1967); G. ’t Hooft, in Acta Universitatis
Wratislavensis no. 38, 12th Winter School of Theoretical physics in Karpacz;
Functional and probabilistic methods in quantum field theory vol. I (1975);
L. F. Abbott, Nucl. Phys. B189, 185 (1981); C. Lee, Nucl. Phys. B207,
157 (1982); in Selected Topics in Theoretical Physics, Proc. Fifth Symposium on Theoretical Physics, Kyohak Yunkusa, Seoul (1986); W. Ditrich and M.
Reuter, Selected Topics in Gauge Theories, Springer-Verlag (1985).
19

18. P. M. Stevenson, Phys. Rev. D 30, 1712 (1984).

20

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
