---
source: "https://arxiv.org/abs/1909.06252v1"
title: "Friedrichs inequality in irregular domains"
author: "S. Creo, M. R. Lancia"
year: "2019"
publication: "arXiv preprint / math.AP"
download: "https://arxiv.org/pdf/1909.06252v1"
pdf: "https://arxiv.org/pdf/1909.06252v1"
captured_at: "2026-05-09T12:08:03Z"
updated_at: "2026-05-09T12:08:03Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Friedrich Nietzsche"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Friedrichs inequality in irregular domains

- 著者: S. Creo, M. R. Lancia
- 年: 2019
- 掲載情報: arXiv preprint / math.AP
- 情報源: [arxiv](https://arxiv.org/abs/1909.06252v1)
- ダウンロード: https://arxiv.org/pdf/1909.06252v1
- PDF: https://arxiv.org/pdf/1909.06252v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

We prove a generalized version of Friedrichs and Gaffney inequalities for a bounded $(\varepsilon,δ)$ domain $Ω\subset\mathbb{R}^n$, $n=2,3$, by adapting the methods of Jones to our framework.

## PDF Text

arXiv:1909.06252v1 [math.AP] 13 Sep 2019
FriedrichsinequalityinirregulardomainsSimoneCreo,MariaRosariaLanciaDipartimentodiScienzediBaseeApplicateperl’Ingegneria,Universit`adeglistudidiRomaSapienza,ViaA.Scarpa16,00161Roma,Italy.AbstractWeproveageneralizedversionofFriedrichsandGaﬀneyinequalitiesforabounded(ε,δ)domainΩ⊂Rn,n=2,3,byadaptingthemethodsofJonestoourframework.Keywords:Friedrichsinequality,Gaﬀneyinequality,(ε,δ)domains,Whitneydecom-position,coercivityestimates.
2010MathematicsSubjectClassiﬁcation:Primary:35A23.Secondary:35Q61,78A25.
1IntroductionΩTheaimofthispaperistoproveFriedrichsinequalityfor(ε,δ)domains.Thisinequal-ityhasbeenintroducedindiﬀerentframeworksbyFriedrichs[
6
]andGaﬀney[
7
],andintheliteratureitisknownwithdiﬀerentnamesaccordingtothesettingwhereitisused.InthestudyofMaxwellproblemsorNavier-Stokesequations,thisinequalityisakeytooltoprovethecoercivityoftheassociatedenergyforms.Fromthepointofviewofapplications,itisinterestingtostudyvectorBVPsinirregulardomains(seee.g.[
3
,
13
])andtheirnumericalapproximation,henceitiscrucialtoextendtheseinequalitiestothecaseofsuitableirregularsets.Fromthisperspective,weconﬁneourselvestotwoorthree-dimensionaldomains.ΩGaﬀneyinequalitycanbededucedfromtheFriedrichsinequality.Toourknowledge,1
suchinequalitiesholdforconvexandLipschitzdomains;amongtheothers,wereferto[
2
,
16
,
15
,
1
],seealso[
4
]andthereferenceslistedin.Inthispaper,weﬁrstproveFriedrichsinequalityfor(ε,δ)domains,andthenproveGaﬀneyinequalitybyadaptingthemethodsof[
5
](developedforKorninequality)tothisframework.Theclassof(ε,δ)domainshasbeenintroducedbyJones[
8
],anditisquitegeneral,sincetheboundaryofan(ε,δ)domaincanbehighlynon-rectiﬁable,e.g.fractalorad-set(seeDeﬁnitions
2.1
and
2.2
).Intheliterature,forΩ⊂Rn(n=2,3)suﬃcientlysmooth,theFriedrichsinequalityreadsasfollows:ifv∈W1,p(Ω)n,thereexistsapositiveconstantC,dependingonΩ,nandp,suchthatkvkW1,p(Ω)n≤C(kvkLp(Ω)n+kdivvkLp(Ω)+kcurlvkLp(Ω)n).(1.1)GaﬀneyinequalityisadirectconsequenceofFriedrichsinequality(
1.1
)whenconsid-eringboundaryconditions.Weintroducethefollowingspaces:Wp(div,Ω):={u∈Lp(Ω)n:divu∈Lp(Ω)},Wp0(div,Ω):={u∈Wp(div,Ω):ν·u=0on∂Ω},Wp(curl,Ω):={u∈Lp(Ω)n:curlu∈Lp(Ω)n},Wp0(curl,Ω):={u∈Wp(curl,Ω):ν×u=0on∂Ω},where·and×denoterespectivelytheusualscalarandcrossproductsbetweenvectorsinRn.Theboundaryconditionshavetobeinterpretedinasuitableweaksense(seee.g.[
18
]).Whenv∈Wp(div,Ω)∩Wp0(curl,Ω)orv∈Wp(curl,Ω)∩Wp0(div,Ω),Gaﬀneyinequalitytakesthefollowingform:k∇vkLp(Ω)n×n≤C(kdivvkLp(Ω)+kcurlvkLp(Ω)n).(1.2)OuraimistoextendGaﬀneyinequalitytothose(ε,δ)domainsforwhichitispossibletogiveaninterpretationoftheboundaryconditions.Inparticular,weconsider(ε,δ)domainsΩinRnwhoseboundariesared-setsorarbitraryclosedsetsinthesenseofJonsson[
9
].Inthesecases,itcanbeprovedthatthespacesWp0(div,Ω)andWp0(curl,Ω)arewelldeﬁnedbecausegeneralizedGreenandStokesformulashold.ThisimpliesthatthenormalandtangentialtracesarewelldeﬁnedaselementsofthedualsofsuitabletraceBesovspacesontheboundary(see[
12
]and[
3
]).Weextend(
1.1
)and(
1.2
)to(ε,δ)domainsΩ⊂Rnforeitherv∈Wp(curl,Ω)∩Wp0(div,Ω)(seeSection
3.1
)orv∈Wp(div,Ω)∩Wp0(curl,Ω)(seeSection
3.2
),accord-ingtotheboundaryconditionsunderconsideration.Themainresultsofthispaper2
areTheorems
3.6
,
3.7
,
3.9
and
3.10
.TheproofofourresultsdeeplyreliesontheassumptionsonΩ.SinceΩisan(ε,δ)domain,foreachv∈Wp(curl,Ω)∩Wp0(div,Ω)(foreachv∈Wp(div,Ω)∩Wp0(curl,Ω)respectively)weconstructasuitableextensionEvbyadaptingJones’approach[
8
].Moreprecisely,weconsideraWhitneydecompositionofΩandweconstructanexten-sionoperatorintermsofsuitablelinearpolynomialswhichsatisﬁesthecrucialestimates(
3.31
)and(
3.32
)((
3.37
)and(
3.38
)respectively).Thethesisisthenachievedbyden-sityarguments.ΩThroughoutthepaper,Cwilldenotediﬀerentpositiveconstant.Sometimes,weindi-catethedependenceoftheseconstantsonsomeparticularparametersinparentheses.2(ε,δ)domainsandtraceresultsWerecallthedeﬁnitionof(ε,δ)(orJones)domain.Deﬁnition2.1.LetF⊂RnbeopenandconnectedandFc:=Rn\F.Forx∈F,letd(x):=infy∈Fc|x−y|.WesaysaythatFisan(ε,δ)domainif,wheneverx,y∈Fwith|x−y|<δ,thereexistsarectiﬁablearcγ∈Fjoiningxtoysuchthat`(γ)≤1
ε|x−y|andd(z)≥ε|x−z||y−z|
|x−y|foreveryz∈γ.AspointedoutintheIntroduction,weconsidertwoparticularclassesof(ε,δ)domainsΩ⊂Rn:i)(ε,δ)domainshavingasboundaryad-set;ii)arbitraryclosed(ε,δ)domainsinthesenseof[
9
].Forthesakeofcompleteness,werecallthedeﬁnitionofd-setgivenin[
10
].Deﬁnition2.2.AclosednonemptysetM⊂Rnisad-set(for0<d≤n)ifthereexistaBorelmeasureµwithsuppµ=Mandtwopositiveconstantsc1andc2suchthatc1rd≤µ(B(P,r)∩M)≤c2rd∀P∈M.(2.1)Themeasureµiscalledd-measure.Inboththecasesi)andii),wecanprovetracetheorems,i.e.GreenandStokesformulas.Forthesakeofsimplicity,werestrictourselvestothecaseinwhich∂Ωisad-set.WerecallthedeﬁnitionofBesovspacespecializedtoourcase.ForgeneralitiesonBesovspaces,wereferto[
10
].3
Deﬁnition2.3.LetGbead-setwithrespecttoad-measureµandα=1−n−d p.Bp,pα(G)isthespaceoffunctionsforwhichthefollowingnormisﬁnite:kukBp,pα(G)=kukLp(G)+

ZZ|P−P0|<1|u(P)−u(P0)|p
|P−P0|d+pαdµ(P)dµ(P0)

1
p.Throughoutthepaper,p0willdenotetheH¨olderconjugateexponentofp.Inthefollowing,wedenotethedualoftheBesovspaceonad-setGwith(Bp,pα(G))0;thisspacecoincideswiththespaceBp0,p0−α(G)(see[
11
]).Theorem2.4(Stokesformula).Letu∈Wp(curl,Ω).Thereexistsalinearandcon-tinuousoperatorlτ(u)=u×νfromWp(curl,Ω)to((Bp0,p0α(∂Ω))0)3.ThefollowinggeneralizedStokesformulaholdsforeveryv∈W1,p0(Ω)n:hu×ν,vi((Bp0,p0α(∂Ω))0)3,Bp0,p0α(∂Ω)3=ZΩu·curlvdx+ZΩv·curludx.(2.2)Moreover,theoperatoru7→lτ(u)=u×νislinearandcontinuousonBp0,p0α(∂Ω)3.Theorem2.5(Greenformula).Letu∈Wp(div,Ω).Thereexistsalinearandcon-tinuousoperatorlν(u)=u·νfromWp(div,Ω)to(Bp0,p0α(∂Ω))0.ThefollowinggeneralizedGreenformulaholdsforeveryv∈W1,p0(Ω):hu·ν,vi(Bp0,p0α(∂Ω))0,Bp0,p0α(∂Ω)=ZΩu·∇vdx+ZΩvdivudx.(2.3)Moreover,theoperatoru7→lν(u)=u·νislinearandcontinuousonBp0,p0α(∂Ω).Fortheproofswereferthereaderto[
12
]and[
3
]withsmallsuitablechanges.ExamplesofdomainsforwhichTheorems
2.4
and
2.5
holdare2Dor3DKoch-typedomains.For-mulas(
2.2
)and(
2.3
)givearigorousmeaningoftheboundaryconditionsinWp0(curl,Ω)andWp0(div,Ω)respectivelyintermsofthedualofsuitableBesovspaces.3FriedrichsandGaﬀneyinequalitiesΩFromnowon,letΩ⊂Rnbeabounded(ε,δ)domain,forn=2,3,havingasboundary∂Ωad-set.4
3.1Thecasev∈Wp(curl,Ω)∩Wp0(div,Ω)Weﬁrstconsiderthecasev∈Wp(curl,Ω)∩Wp0(div,Ω).Wepointoutthat,sinceν·v=0on∂Ω,wehavethatZΩdivvdx=0.(3.1)LetS⊂RnbeameasurablesubsetofRn;wedenoteby¯xitsbarycenter.WeconstructtheaﬃnevectorﬁeldPS(u)associatedtoSandu∈Wp(curl,S)∩Wp0(div,S)inthefollowingway:PS(u)(x)=a+B(x−¯x),(3.2)wherea∈RnandBisan×nmatrixwithentriesbijdeﬁnedasa=1
|S|ZSudxandbij=1
2|S|ZS∂ui
∂xj+∂uj
∂xidx.(3.3)Wepointoutthat,fromthedeﬁnition,Bisasymmetricmatrix.Moreover,bycalcu-lationitfollowsthatcurl(PS(u))=0,div(PS(u))=1
|S|ZSdivudx(3.4)andZS(u−PS(u))dx=0.(3.5)Bydirectcomputation,itholdsthatk∇(u−PS(u))kLp(S)n×n≤Ck∇ukLp(S)n×n,(3.6)whereCdependsonlyon|S|.Letusnowsupposethat(
1.2
)holdsinSforu∈Wp(curl,S)∩Wp0(div,S).(
3.6
)infersthatk∇(u−PS(u))kLp(S)n×n≤C kcurlukLp(S)n+kdivukLp(S).(3.7)Sinceu−PS(u)hasvanishingmeanvalueonSfrom(
3.5
),fromPoincar´e-Wirtingerinequalityand(
3.7
)wehaveku−PS(u)kLp(S)n≤Cdiam(S) kcurlukLp(S)n+kdivukLp(S),(3.8)wherediam(S)isthediameterofS.Now,onecaneasilyseethatk∇PS(u)kL∞(S)n×n≤k∇ukL∞(S)n×n;(3.9)5
hence,byusingagainPoincar´e-Wirtingerinequality(withp=∞),triangleinequalityand(
3.9
)wegetku−PS(u)kL∞(S)n≤Cdiam(S)k∇(u−PS(u))kL∞(S)n×n≤2Cdiam(S)k∇ukL∞(S)n×n.(3.10)Fromnowon,wechoosev∈W1,∞(Ω)n.Thethesiswillthenfollowbydensityargu-ments.WeconstructtheextensionEvfollowingtheapproachofJones[
8
]byusingthelinearpolynomialsPS(v).LetusrecallthatanyopensetΩ⊂Rnadmitsaso-calledWhitneydecomposition(see[
19
],[
17
])intodyadiccubesSk,i.e.Ω=[kSk.Thisdecompositionissuchthat1≤dist(Sk,∂Ω)
`(Sk)≤4√
n∀k,(3.11)S0j∩S0k=∅ifj6=k,(3.12)1
4≤`(Sj)
`(Sk)≤4ifSj∩Sk6=∅,(3.13)whereS0denotestheinteriorofSand`(S)istheedgelengthofacubeS.LetnowW1={Sk}beaWhitneydecompositionofΩandW2={Qj}beaWhitneydecompositionof(Ωc)0.WesetW3=Qj∈W2:`(Qj)≤εδ
16n.Inhispaper,Joneshasshownthat,foreveryQj∈W3,onecanchoosea“reﬂected”cubeQ∗
j=Sk∈W1suchthatΩ1≤`(Sk)
`(Qj)≤4anddist(Qj,Sk)≤C`(Qj),(3.14)seeLemma2.4andLemma2.8in[
8
].Moreover,ifQj,Qk∈W3havenon-emptyintersection,thereexistsachainFj,k={Q∗
j=S1,S2,...,Sm=Q∗
k}ofcubesinW1whichconnectsQ∗
jandQ∗
ksuchthatSi∩Si+16=∅andm≤C(ε,δ).From[
17
],[
19
]itfollowsthatthereexistsapartitionofunity{φj},associatedwiththeWhitneydecomposition,suchthatφj∈C∞(Rn),suppφj⊂17
16Qj,0≤φj≤1,PQj∈W3φj=1onSQj∈W3Qjand|∇φj|≤C`(Qj)−1∀j.6
Forv∈W1,∞(Ω)n,letPj:=PQ∗
j(v)bedeﬁnedasin(
3.2
)and(
3.3
).WenowdeﬁnetheextensionEvofvtoRninthefollowingway:Ev=



XQj∈W3Pjφjin(Ωc)0,vinΩ.Wepointoutthat,sincetheboundaryofan(ε,δ)domainhaszeromeasure(seeLemma2.3in[
8
]),itfollowsthatEvisdeﬁneda.e.inRn.Fromnowon,ifnototherwisespeciﬁed,inthissubsectionweassumethatv∈Wp(curl,Ω)∩Wp0(div,Ω)∩Wp0(div,S)foreveryS∈W1.Wenowprovesomepre-liminarylemmas.Forthesakeofcompleteness,werecallLemma2.1in[
8
].Lemma3.1.LetQbeacubeandletF,G⊂Qbetwomeasurablesubsetssuchthat|F|,|G|≥γ|Q|forsomeγ>0.IfPisapolynomialofdegree1,thenkPkLp(F)≤C(γ)kPkLp(G).Lemma3.2.LetF={S1,...,Sm}beachainofcubesinW1.ThenkPS1(v)−PSm(v)kLp(S1)n≤C(m)`(S1) kcurlvkLp(∪jSj)n+kdivvkLp(∪jSj)(3.15)andkPS1(v)−PSm(v)kL∞(S1)n≤C(m)`(S1)k∇vkL∞(∪jSj)n×n.(3.16)Proof.Wewilluse(
3.8
),whereSisacubeoraunionoftwoneighboringcubes.From(
3.13
),itfollowsthatthenumberofpossiblegeometriesofSisﬁnite;hence,wecanﬁndauniformconstantin(
3.8
).ByusingLemma
3.1
,wegetkPS1(v)−PSm(v)kLp(S1)n≤m−1Xr=1kPSr(v)−PSr+1(v)kLp(S1)n≤c(m)m−1Xr=1kPSr(v)−PSr+1(v)kLp(Sr)n≤c(m)m−1Xr=1kPSr(v)−PSr∪Sr+1(v)kLp(Sr)n+kPSr∪Sr+1(v)−PSr+1(v)kLp(Sr+1)n ≤c(m)m−1Xr=1kPSr(v)−vkLp(Sr)n+kPSr+1(v)−vkLp(Sr+1)n+kPSr∪Sr+1(v)−vkLp(Sr∪Sr+1)n ≤Cc(m)`(S1) kcurlvkLp(∪jSj)n+kdivvkLp(∪jSj),7
whereweusedthefactthatFisachain,integralpropertiesandﬁnally(
3.8
).Theproofof(
3.16
)followsanalogouslybyusing(
3.10
).
ForeveryQj,Qk∈W3withnon-emptyintersection,wenowchooseachainFj,kwhichconnectsQ∗
jandQ∗
kandsuchthatm≤C(ε,δ).WedeﬁneF(Qj)=[Qk∈W3,Qj∩Qk6=∅Fj,k;hence

XQk:Qj∩Qk6=∅χ∪Fj,k

L∞(Rn)≤C∀Qj∈W3.(3.17)WenowprovetwolemmaswhichallowustocontrolthenormsofEv,div(Ev),curl(Ev)and∇(Ev)in(Ωc)0.Lemma3.3.LetQ0∈W3.Wehavethat:kEvkLp(Q0)n≤C kvkLp(Q∗
0)n+`(Q0)(kcurlvkLp(F(Q0))n+kdivvkLp(F(Q0))),(3.18)kcurl(Ev)kLp(Q0)n+kdiv(Ev)kLp(Q0)≤C kcurlvkLp(F(Q0))n+kdivvkLp(F(Q0)),(3.19)kEvkL∞(Q0)n≤C kvkL∞(Q∗
0)n+`(Q0)k∇vkL∞(F(Q0))n×n,(3.20)k∇(Ev)kL∞(Q0)n×n≤Ck∇vkL∞(F(Q0))n×n.(3.21)Proof.Werecallthat,fromthedeﬁnitionofEv,onQ0wehavethatEv=XQj∈W3Pjφj.Moreover,sinceXQj∈W3φj≡1on[Qj∈W3Qj,weget

XQj∈W3Pjφj

Lp(Q0)n≤kP0kLp(Q0)n+

XQj∈W3(Pj−P0)φj

Lp(Q0)n:=A+B.WenowestimateAandBseparately.AstoA,fromLemma
3.1
and(
3.8
),wegetA=kP0kLp(Q0)n≤CkP0kLp(Q∗
0)n≤C(kP0−vkLp(Q∗
0)n+kvkLp(Q∗
0)n)≤C(`(Q0)(kcurlvkLp(Q∗
0)n+kdivvkLp(Q∗
0))+kvkLp(Q∗
0)n),(3.22)whereweestimated`(Q∗Ω0)with`(Q0)using(
3.14
),sinceQ0∈W3.Wepointoutthat,thanksto(
3.17
),thenormsintheright-handsideof(
3.22
)canbeestimatedintermsoftheLp(∪jF0,j)-norms.Hence,wegetthefollowing:A≤C(`(Q0)(kcurlvkLp(∪jF0,j)n+kdivvkLp(∪jF0,j))+kvkLp(∪jF0,j)n).(3.23)8
AstoB,fromthepropertiesofφjitissuﬃcienttoboundkPj−P0kLp(Q0)n.ByusingagainLemma
3.1
,(
3.15
)andproceedingasabove,wegetB≤kPj−P0kLp(Q0)n≤CkPj−P0kLp(Q∗
0)n≤C`(Q0)(kcurlvkLp(∪jF0,j)n+kdivvkLp(∪jF0,j)).(3.24)Hencefrom(
3.23
)and(
3.24
)weget(
3.18
).Estimate(
3.20
)followssimilarlybyusing(
3.10
)and(
3.16
).Wenowremarkthat,onQ0,wehavethatEv=XQj∈W3Pjφj=P0XQj∈W3φj+XQj∈W3(Pj−P0)φj=P0+XQj∈W3(Pj−P0)φj.Therefore,sincecurl(P0)=0,wehavethatcurl(Ev)=XQj∈W3curl((Pj−P0)φj).Moreover,from(
3.1
)and(
3.4
)itfollowsthatdiv(Ev)=XQj∈W3div((Pj−P0)φj).SincethereisaﬁnitenumberofcubesQjsuchthatφj6=0inQ0andhavingnon-emptyintersectionwithQ0,from(
3.13
)wehavethat`(Qj)≥1
4`(Q0).Fromthepropertiesofφj,thisimpliesthat|∇φj|≤C
4`(Q0)−1.Byusingvectoridentities,Lemma
3.1
and(
3.15
),wehavethatkcurl((Pj−P0)φj)kLp(Q0)n=k(Pj−P0)×∇φjkLp(Q0)n≤CkPj−P0kLp(Q0)nk∇φjkLp(Q0)n≤C`(Q0)−1kPj−P0kLp(Q0)n≤C`(Q0)−1kPj−P0kLp(Q∗
0)n≤C(kcurlvkLp(∪jF0,j)n+kdivvkLp(∪jF0,j)).Astodivergenceterm,similarlyasabovewegetkdiv((Pj−P0)φj)kLp(Q0)=k(Pj−P0)·∇φjkLp(Q0)≤kPj−P0kLp(Q0)nk∇φjkLp(Q0)n≤C(kcurlvkLp(∪jF0,j)n+kdivvkLp(∪jF0,j)).Summingupinjwegetkcurl(Ev)kLp(Q0)n+kdiv(Ev)kLp(Q0)≤C(kcurlvkLp(F(Q0))n+kdivvkLp(F(Q0))),i.e.(
3.19
).Wearelefttoprove(
3.21
).Similarlyasabove,wehavethat∇(Ev)=∇P0+XQj∈W3∇((Pj−P0)φj).9
FromLemma
3.1
and(
3.9
),wegetk∇P0kL∞(Q0)n×n≤Ck∇P0kL∞(Q∗
0)n×n≤Ck∇vkL∞(Q∗
0)n×n≤Ck∇vkL∞(∪jF0,j)n×n.Asabove,itfollowsthatk∇(Pj−P0)kL∞(Q0)n×n≤Ck∇(Pj−P0)kL∞(Q∗
0)n×n≤Ck∇(Pj−P0)kL∞(Q∗
0∪Q∗
j)n×n≤Ck∇vkL∞(Q∗
0∪Q∗
j)n×n≤Ck∇vkL∞(∪jF0,j)n×nFromtheseinequalities,(
3.21
)followsandtheproofiscomplete.
WenowprovearesultsimilartoLemma
3.3
,whichrelatestothecubesof(Ωc)0notbelongingtoW3.Lemma3.4.LetQ0∈W2\W3.Wehavethat:kEvkLp(Q0)n≤CXQj∈W3:Qj∩Q06=∅kvkLp(Q∗
j)n+kcurlvkLp(Q∗
j)n+kdivvkLp(Q∗
j),(3.25)kcurl(Ev)kLp(Q0)n+kdiv(Ev)kLp(Q0)≤CXQj∈W3:Qj∩Q06=∅kvkLp(Q∗
j)n+kcurlvkLp(Q∗
j)n+kdivvkLp(Q∗
j),(3.26)kEvkL∞(Q0)n≤CXQj∈W3:Qj∩Q06=∅kvkL∞(Q∗
j)n+k∇vkL∞(Q∗
j)n×n,(3.27)k∇(Ev)kL∞(Q0)n×n≤CXQj∈W3:Qj∩Q06=∅kvkL∞(Q∗
j)n+k∇vkL∞(Q∗
j)n×n.(3.28)Proof.Westartbypointingoutthat,ifφj6=0onQ0,wehaveQj∩Q06=∅(sincesuppφj⊂17
16Qj).Therefore,sinceQ0∈W2\W3,wehave`(Qj)≥1
4`(Q0)≥εδ
64n.(3.29)OnQ0wehavethat|Ev|=

XQj∈W3:Qj∩Q06=∅Pjφj

≤XQj∈W3:Qj∩Q06=∅|Pj|.FromLemma
3.1
andtriangleinequality,wegetkPjkLp(Q0)n≤CkPjkLp(Qj)n≤CkPjkLp(Q∗
j)n≤C(kPj−vkLp(Q∗
j)n+kvkLp(Q∗
j)n).(3.30)10
From(
3.30
)and(
3.8
),itfollowsthatkEvkLp(Q0)n≤XQj∈W3:Qj∩Q06=∅kPjkLp(Q0)n≤CXQj∈W3:Qj∩Q06=∅kvkLp(Q∗
j)n+diam(Q∗
j)(kcurlvkLp(Q∗
j)n+kdivvkLp(Q∗
j)).SineΩisbounded,wecanestimatediam(Q∗
j)withaconstantdependingondiam(Ω),thusproving(
3.25
).Wecometo(
3.26
).ByproceedingasintheproofofLemma
3.3
andbyusing(
3.29
),thefollowingestimateholds:kcurl(Ev)kLp(Q0)n+kdiv(Ev)kLp(Q0)=

XQj∈W3:Qj∩Q06=∅curl(Pjφj)

Lp(Q0)n+

XQj∈W3:Qj∩Q06=∅div(Pjφj)

Lp(Q0)≤CXQj∈W3:Qj∩Q06=∅kPjkLp(Q0)nk∇φjkLp(Q0)n×n≤C`(Q0)−1XQj∈W3:Qj∩Q06=∅kPjkLp(Q0)n≤Cεδ
64n−1XQj∈W3:Qj∩Q06=∅kPjkLp(Q∗
0)n.Byproceedingasabove,weget(
3.26
).Estimates(
3.27
)and(
3.28
)followinasimilarwaybyusing(
3.10
)and(
3.9
).
Fromtheabovelemmasweobtainthefollowingresult.
Proposition3.5.Foreveryv∈W1,∞(Ω)nsuchthatv∈Wp(curl,Ω)∩Wp0(div,Ω)wehavekEvkLp((Ωc)0)n+kdiv(Ev)kLp((Ωc)0)+kcurl(Ev)kLp((Ωc)0)n≤C kvkLp(Ω)n+kdivvkLp(Ω)+kcurlvkLp(Ω)n(3.31)andkEvkW1,∞((Ωc)0)n≤CkvkW1,∞(Ω)n.(3.32)Proof.BysummingupovereveryQ0∈W2,thethesisfollowsasadirectconsequenceofLemma
3.3
andLemma
3.4
.Inparticular,(
3.31
)followsfrom(
3.18
),(
3.19
),(
3.25
)and(
3.26
),while(
3.32
)followsfrom(
3.20
),(
3.21
),(
3.27
)and(
3.28
).
Wenowprovetheﬁrstmainresultofthispaper,whichfollowsfromtheabovelemmas.11
Theorem3.6(Friedrichsinequality).LetΩ⊂Rnbeabounded(ε,δ)domainwith∂Ωad-set.ThereexistsaconstantC=C(ε,δ,n,p,Ω)>0suchthat,foreveryv∈W1,p(Ω)nsuchthatv∈Wp(curl,Ω)∩Wp0(div,Ω),kvkW1,p(Ω)n≤C kvkLp(Ω)n+kcurlvkLp(Ω)n+kdivvkLp(Ω).(3.33)Proof.Itissuﬃcienttoprove(
3.33
)forv∈W1,∞(Ω)n;thethesiswillthenfollowbydensity.WerecallthattheextensionEvisdeﬁneda.e.onRnsince|∂Ω|=0.Moreover,fromthedeﬁnitionofEvwecansupposethatsuppEviscontainedinaballB.SinceEv∈W1,p(B)n,from(
3.31
)wehavethatkEvkLp(B)n+kcurl(Ev)kLp(B)n+kdiv(Ev)kLp(B)≤C kvkLp(Ω)n+kcurlvkLp(Ω)n+kdivvkLp(Ω).Hence,fromFriedrichsinequalityforsmoothdomainsandtheaboveinequality,wegetkvkW1,p(Ω)n=kEvkW1,p(Ω)n≤kEvkW1,p(B)n≤C(kEvkLp(B)n+kcurl(Ev)kLp(B)n+kdiv(Ev)kLp(B))≤C kvkLp(Ω)n+kcurlvkLp(Ω)n+kdivvkLp(Ω),i.e.thethesis.
WeconcludethissectionbyprovingGaﬀneyinequalityasadirectconsequenceofTheorem
3.6
.Theorem3.7(Gaﬀneyinequality).LetΩ⊂Rnbeaboundedsimplyconnected(ε,δ)domainwith∂Ωad-set.Letv∈W1,p(Ω)nbesuchthatv∈Wp(curl,Ω)∩Wp0(div,Ω).ThenthereexistsC=C(ε,δ,n,p,Ω)>0suchthatkvkW1,p(Ω)n≤C kcurlvkLp(Ω)n+kdivvkLp(Ω).(3.34)Proof.Wearguebycontradiction.Letussupposethat(
3.34
)doesnothold;hence,thereexistsasequenceofvectors{vk}⊂W1,p(Ω)n∩Wp(curl,Ω)∩Wp0(div,Ω)suchthatkvkkW1,p(Ω)n=1andkcurlvkkLp(Ω)n+kdivvkkLp(Ω)−−−−→k→+∞0.SincekvkkW1,p(Ω)n=1,thereexistsasubsequenceof{vk}(whichwestilldenotebyvk)suchthatvk*vinW1,p(Ω)nandvk→vinLp(Ω)n.12
Sincethedistributionallimitscoincidewiththeweaklimits,itimmediatelyfollowsthatdivv=0andcurlv=0.Wenowprovethat{vk}isaCauchysequenceinW1,p(Ω)n.FromFriedrichsinequality(
3.33
),foreveryk,j∈Nonehaskvk−vjkW1,p(Ω)n≤C kvk−vjkLp(Ω)n+kdiv(vk−vj)kLp(Ω)+kcurl(vk−vj)kLp(Ω)n.(3.35)FromthestrongconvergenceofvkinLp(Ω)n,theﬁrsttermontheright-handsideof(
3.35
)vanishes.Astotheothertwoterms,theyalsovanishsincecurlvkanddivvkbothtendto0inLpask→+∞.HencevkisaCauchysequenceinW1,p(Ω)n,andvk→vstronglyinW1,p(Ω)n.Werecallthatifcurlv=0inΩandΩissimplyconnected,thereexistsafunctionΦ∈W1,p(Ω)suchthatv=∇Φ.Thisinturnimpliesthat∆Φ=div∇Φ=divv=0inΩ.Moreover,sincev∈Wp0(div),wealsohavethat∂Φ
∂ν=ν·∇Φ=ν·v=0on∂Ω.HenceΦ∈W1,p(Ω)istheuniqueweaksolutionofthefollowingproblem



∆Φ=0inΩ,∂Φ
∂ν=0on∂Ω.(3.36)ThisimpliesthatΦisconstant,andsov=∇Φ=0onΩ.Wereachedacontradiction,since1=kvkkW1,p(Ω)n−−−−→k→+∞kvkW1,p(Ω)n=0.
3.2Thecasev∈Wp(div,Ω)∩Wp0(curl,Ω)Wenowconsiderthecasev∈Wp(div,Ω)∩Wp0(curl,Ω).Werecallthatthisimpliesν×v=0on∂ΩinthedualofBp0,p0α(∂Ω).Weapproximatev∈Wp(div,Ω)∩Wp0(curl,Ω)∩W1,∞(Ω)nbymeansofthepolynomialsPjasintheprevioussection.WeremarkthatinthiscasedivPj=1
|Q∗
j|ZQ∗
jdivvdx6=0.Asintheprevioussubsection,estimates(
3.7
)and(
3.8
)hold,aswellaslemmas
3.2
,
3.3
and
3.4
,underthehypothesisthatv∈Wp(div,Ω)∩Wp0(curl,Ω)∩Wp0(curl,S)foreveryS∈W1.13
Forthesakeofclarity,westatetheanalogousofProposition
3.5
andTheorem
3.6
inthiscase.
Proposition3.8.Foreveryv∈W1,∞(Ω)nsuchthatv∈Wp(div,Ω)∩Wp0(curl,Ω)wehavekEvkLp((Ωc)0)n+kdiv(Ev)kLp((Ωc)0)+kcurl(Ev)kLp((Ωc)0)n≤C kvkLp(Ω)n+kdivvkLp(Ω)+kcurlvkLp(Ω)n(3.37)andkEvkW1,∞((Ωc)0)n≤CkvkW1,∞(Ω)n.(3.38)Theorem3.9(Friedrichsinequality).LetΩ⊂Rnbeabounded(ε,δ)domainwith∂Ωad-set.ThereexistsaconstantC=C(ε,δ,n,p,Ω)>0suchthat,foreveryv∈W1,p(Ω)nsuchthatv∈Wp(div,Ω)∩Wp0(curl,Ω),kvkW1,p(Ω)n≤C kvkLp(Ω)n+kcurlvkLp(Ω)n+kdivvkLp(Ω).(3.39)WeconcludebyprovingGaﬀneyinequality.
Theorem3.10(Gaﬀneyinequality).LetΩ⊂Rnbeaboundedsimplyconnected(ε,δ)domainwith∂Ωad-set.Letv∈W1,p(Ω)nbesuchthatv∈Wp(div,Ω)∩Wp0(curl,Ω).ThenthereexistsC=C(ε,δ,n,p,Ω)>0suchthatkvkW1,p(Ω)n≤C kcurlvkLp(Ω)n+kdivvkLp(Ω).(3.40)Proof.Weproceedasintheproofof[
14
,Corollary3.51];wearguebycontradictionandwesupposethat(
3.40
)doesnothold.AsinTheorem
3.7
,thismeansthatthereexistsasequenceofvectors{vk}⊂W1,p(Ω)n∩Wp(div,Ω)∩Wp0(curl,Ω)suchthatkvkkW1,p(Ω)n=1andkcurlvkkLp(Ω)n+kdivvkkLp(Ω)−−−−→k→+∞0.Thisimpliesthatvk*vinW1,p(Ω)nandvk→vinLp(Ω)n,withdivv=0andcurlv=0.From(
3.39
),foreveryk,j∈Nwehavethatkvk−vjkW1,p(Ω)n≤C kvk−vjkLp(Ω)n+kdiv(vk−vj)kLp(Ω)+kcurl(vk−vj)kLp(Ω)n.(3.41)14
AsintheproofofTheorem
3.7
,allthetermsontheright-handsideof(
3.41
)vanishwhenk,j→+∞,hence{vk}isaCauchysequenceinW1,p(Ω)n.Asinthecasev∈Wp(curl,Ω)∩Wp0(div,Ω),thereexistsafunctionΦ∈W1,p(Ω)suchthatv=∇Φand∆Φ=0inΩ.Sinceinthiscasev∈Wp0(curl,Ω),wealsohavethatν×∇Φ=ν×v=0on∂Ω.UptoshiftingΦbyaconstant,thisimpliesthatΦ=0on∂Ωinthetracesense.HenceΦ∈W1,p(Ω)istheuniqueweaksolutionofthefollowingproblem

∆Φ=0inΩ,Φ=0on∂Ω.(3.42)ThisimpliesthatΦ=0,thereforev=0onΩandwereachthecontradiction.
Acknowledgements.TheauthorshavebeensupportedbytheGruppoNazionaleperl’AnalisiMatematica,laProbabilit`aeleloroApplicazioni(GNAMPA)oftheIstitutoNazionalediAltaMatematica(INdAM).
References[1]C.Amrouche,C.Bernardi,M.Dauge,V.Girault,Vectorpotentialsinthree-dimensionalnon-smoothdomains,Math.MethodsAppl.Sci.,21(1998),823–864.[2]S.Bauer,D.Pauly,OnKorn’sﬁrstinequalityfortangentialornormalboundaryconditionswithexplicitconstants,Math.MethodsAppl.Sci.,39(2016),5695–5704.[3]S.Creo,M.Hinz,M.R.Lancia,A.Teplyaev,P.Vernole,Magnetostaticproblemsinfractaldomains,acceptedforpublicationonFractalsandDynamicsinMath-ematics,ScienceandtheArtspublishedbyWorldScientiﬁc.AvailableonarXiv:https://arxiv.org/abs/1805.08262[4]G.Csato,B.Dacorogna,S.Sil,OnthebestconstantinGaﬀneyinequality,J.Funct.Anal.,274(2018),461–503.[5]R.Dur´an,M.A.Muschietti,TheKorninequalityforJonesdomains,Electron.J.DiﬀerentialEquations,127(2004),10pp.[6]K.O.Friedrichs,DiﬀerentialformsonRiemannianmanifolds,Comm.PureAppl.Math.,8(1955),551–590.15
[7]M.P.Gaﬀney,Hilbertspacemethodsinthetheoryofharmonicintegrals,Trans.Amer.Math.Soc.,78(1955),426–444.[8]P.W.Jones,QuasiconformalmappingandextendabilityoffunctionsinSobolevspaces,ActaMath.,147(1981),71–88.[9]A.Jonsson,BesovspacesonclosedsubsetsofRn,Trans.Amer.Math.Soc.,341(1994),355–370.[10]A.Jonsson,H.Wallin,FunctionSpacesonSubsetsofRn,Part1,Math.Reports,vol.2,HarwoodAcad.Publ.,London,1984.[11]A.Jonsson,H.Wallin,ThedualofBesovspacesonfractals,StudiaMath.,112(1995),285–300.[12]M.R.Lancia,P.Vernole,Semilinearfractalproblems:approximationandregu-larityresults,NonlinearAnal.,80(2013),216–232.[13]M.R.Lancia,P.Vernole,TheStokesproblemsinfractaldomains:asymptoticbehaviorofthesolutions,acceptedforpublicationonDiscreteContin.Dyn.Syst.Ser.S.[14]P.Monk,FiniteElementMethodsforMaxwell’sEquations,OxfordUniversityPress,NewYork,2003.[15]P.Neﬀ,D.Pauly,K.J.Witsch,Poincar´emeetsKornviaMaxwell:ExtendingKorn’sﬁrstinequalitytoincompatibletensorﬁelds,J.Diﬀ.Eq.,258(2015),1267–1302.[16]B.Schweizer,OnFriedrichsinequality,Helmholtzdecomposition,vectorpotentials,andthediv-curllemma,preprint(2016),TUDortmund.[17]E.M.Stein,SingularIntegralsandDiﬀerentiabilityPropertiesofFunctions,PrincetonUniversityPress,Princeton,NewJersey,1970.[18]R.Temam,Navier-StokesEquations.TheoryandNumericalAnalysis,StudiesinMathematicsanditsApplications,2,North-HollandPublishingCo.,Amsterdam-NewYork,1979.[19]H.Whitney,Analyticextensionsofdiﬀerentiablefunctionsdeﬁnedinclosedsets,Trans.Amer.Math.Soc.,36(1934),63–89.16

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
